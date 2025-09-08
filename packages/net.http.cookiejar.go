package packages

import (
	"net/http/cookiejar"
	"reflect"
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
			"New": reflect.ValueOf(cookiejar.New),
		}
		storeFuncs("net/http/cookiejar", funcs)
		//env.PackageTypes["net/http/cookiejar"]
		types := map[string]reflect.Type{
			"Options": reflect.TypeOf(cookiejar.Options{}),
		}
		storeTypes("net/http/cookiejar", types)
	}()
}
