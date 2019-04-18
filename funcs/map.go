package funcs

import (
	"sort"
)

func keys(v map[string]interface{}) (l []string) {
	for k := range v {
		l = append(l, k)
	}
	sort.Strings(l)
	return l
}

func init() {
	Register("keys", keys)
}
