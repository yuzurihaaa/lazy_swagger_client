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

func MergeMaps[M ~map[K]V, K comparable, V any](src ...M) M {
	merged := make(M)
	for _, m := range src {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}
