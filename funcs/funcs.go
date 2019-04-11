package funcs

import (
	"reflect"
	"sync"
)

const (
	pkgName = "gobst/funcs"
)

var (
	funcs = map[string]interface{}{}

	funcsMu sync.Mutex
)

func Funcs() map[string]interface{} {
	funcsMu.Lock()
	defer funcsMu.Unlock()

	out := make(map[string]interface{}, len(funcs))
	for k, v := range funcs {
		out[k] = v
	}
	return out
}

func Register(name string, fn interface{}) {
	rt := reflect.TypeOf(fn)
	if rt.Kind() != reflect.Func {
		panic(pkgName + ": fn is not a function")
	}

	funcsMu.Lock()
	defer funcsMu.Unlock()

	if _, dup := funcs[name]; dup {
		panic(pkgName + `: duplicate func name "` + name + `"`)
	}
	funcs[name] = fn
}
