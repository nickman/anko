package packages

import (
	"reflect"
	"sync"
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
			"NewCond": reflect.ValueOf(sync.NewCond),
		}
		storeFuncs("sync", funcs)
		//env.PackageTypes["sync"]
		types := map[string]reflect.Type{
			"Cond":      reflect.TypeOf(sync.Cond{}),
			"Mutex":     reflect.TypeOf(sync.Mutex{}),
			"Once":      reflect.TypeOf(sync.Once{}),
			"Pool":      reflect.TypeOf(sync.Pool{}),
			"RWMutex":   reflect.TypeOf(sync.RWMutex{}),
			"WaitGroup": reflect.TypeOf(sync.WaitGroup{}),
			"Map":       reflect.TypeOf(sync.Map{}),
		}
		storeTypes("sync", types)
	}()

}
