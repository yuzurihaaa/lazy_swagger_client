package swagger

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"swagger_client/src/utilities"
	"time"
)

type Swagger struct {
	Cache map[string]Out
	Config
}

type Config struct {
	Transport http.RoundTripper
	Scheme    string
	Host      string
}

type Out struct {
	Url    string
	Method string
}

// NewSwaggerF
// Constructor
// Read swagger files and map to Cache
func NewSwaggerF(f string) *Swagger {
	jsonFile, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	byteValue, _ := io.ReadAll(jsonFile)
	swaggerMap := utilities.JsonUnmarshal[map[string]interface{}](byteValue)
	return NewSwagger(swaggerMap)
}

// NewSwagger
// Constructor
// Read swagger json and map to Cache
func NewSwagger(in map[string]interface{}) *Swagger {
	paths, ok := in["paths"].(map[string]interface{})
	if !ok {
		return nil
	}
	cache := buildCache(paths)
	return &Swagger{
		Cache: cache,
	}
}

type Args struct {
	Body  io.Reader
	Path  map[string]any
	Query map[string]string
}

func (s *Swagger) UpdateConfig(c Config) {
	s.Config = c
}

func (s *Swagger) Execute(ctx context.Context, operationId string, arg Args) (*http.Response, error) {
	out := s.Cache[operationId]

	query := url.Values{}
	if len(arg.Query) > 0 {
		for q := range arg.Query {
			query.Set(q, arg.Query[q])
		}
	}

	outUrl := url.URL{
		Scheme:   s.Scheme,
		Host:     s.Host,
		Path:     buildUrl(out.Url, arg.Path),
		RawQuery: query.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, out.Method, outUrl.String(), arg.Body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: s.Transport,
		Timeout:   time.Second * 30,
	}

	return client.Do(req)
}

// buildUrl
// Converts path parameters define in swagger to value
// Example: /{arg1}/url/endpoints where arg1 is with value "TEST" it will be /TEST/url/endpoints
func buildUrl(url string, arg map[string]any) string {
	if len(arg) <= 0 {
		return url
	}
	out := url
	for p := range arg {
		out = strings.Replace(out, fmt.Sprintf("{%v}", p), fmt.Sprintf("%v", arg[p]), 1)
	}
	return out
}

func buildCache(in map[string]interface{}) map[string]Out {
	cache := make(map[string]Out)
	for kp, vp := range in {
		pathMethods, ok := vp.(map[string]interface{})
		if !ok {
			continue
		}
		methods := []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch}
		for _, method := range methods {
			op := pathMethods[strings.ToLower(method)]
			if op == nil {
				continue
			}
			config, ok := op.(map[string]interface{})
			if !ok {
				continue
			}

			cache[config["operationId"].(string)] = Out{
				Url:    kp,
				Method: method,
			}
		}

	}
	return cache
}
