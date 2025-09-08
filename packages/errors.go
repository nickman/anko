package packages

import (
	"errors"
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

		errFuncs := map[string]reflect.Value{
			"New": reflect.ValueOf(errors.New),
		}
		storeFuncs("errors", errFuncs)
	}()
}
