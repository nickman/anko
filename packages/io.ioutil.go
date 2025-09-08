package packages

import (
	"io/ioutil"
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

		ioFuncs := map[string]reflect.Value{
			"ReadAll":   reflect.ValueOf(ioutil.ReadAll),
			"ReadDir":   reflect.ValueOf(ioutil.ReadDir),
			"ReadFile":  reflect.ValueOf(ioutil.ReadFile),
			"WriteFile": reflect.ValueOf(ioutil.WriteFile),
		}
		storeFuncs("io/ioutil", ioFuncs)
	}()
}
