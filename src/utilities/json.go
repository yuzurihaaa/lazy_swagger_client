package utilities

import "encoding/json"

func JsonUnmarshal[T any](in []byte) T {
	t := new(T)
	err := json.Unmarshal(in, t)
	if err != nil {
		return *t
	}
	return *t
}
