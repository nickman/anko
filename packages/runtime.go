package packages

import (
	"reflect"
	"runtime"
	"sync/atomic"
)

func init() {
	defer func() {
		atomic.AddInt32(initCount, 1)
	}()

	go func() {
		completion.Add(1)
		defer completion.Done()

		funcs := map[string]reflect.Value{
			"GC":         reflect.ValueOf(runtime.GC),
			"GOARCH":     reflect.ValueOf(runtime.GOARCH),
			"GOMAXPROCS": reflect.ValueOf(runtime.GOMAXPROCS),
			"GOOS":       reflect.ValueOf(runtime.GOOS),
			"GOROOT":     reflect.ValueOf(runtime.GOROOT),
			"Version":    reflect.ValueOf(runtime.Version),
		}
		storeFuncs("runtime", funcs)
	}()
}
