package packages

import (
	"os/exec"
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
			"ErrNotFound": reflect.ValueOf(exec.ErrNotFound),
			"LookPath":    reflect.ValueOf(exec.LookPath),
			"Command":     reflect.ValueOf(exec.Command),
		}
		storeFuncs("os/exec", funcs)
	}()
}
