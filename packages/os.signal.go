package packages

import (
	"os/signal"
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
			"Notify": reflect.ValueOf(signal.Notify),
			"Stop":   reflect.ValueOf(signal.Stop),
		}
		storeFuncs("os/signal", funcs)
	}()
}
