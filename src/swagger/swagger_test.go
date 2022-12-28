package swagger

import (
	"reflect"
	"testing"
)

func Test_buildUrl(t *testing.T) {
	type args struct {
		url string
		arg map[string]any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test url conversion",
			args: args{
				url: "/{book}/user/{id}",
				arg: map[string]any{
					"book": "My lovely Book",
					"id":   20,
				},
			},
			want: "/My lovely Book/user/20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildUrl(tt.args.url, tt.args.arg); got != tt.want {
				t.Errorf("buildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSwaggerF(t *testing.T) {
	tests := []struct {
		name string
		f    string
		want *Swagger
	}{
		{
			name: "Test read sample file",
			f:    "./test.openapi.json",
			want: &Swagger{Cache: map[string]Out{
				"listVersionsv2": {
					Url:    "/",
					Method: "GET",
				},
				"getVersionDetailsv2": {
					Url:    "/v2",
					Method: "GET",
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSwaggerF(tt.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSwaggerF() = %v, want %v", got, tt.want)
			}
		})
	}
}
