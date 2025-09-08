package packages

import (
	"encoding/json"
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

		jsonFuncs := map[string]reflect.Value{
			"Marshal":   reflect.ValueOf(json.Marshal),
			"Unmarshal": reflect.ValueOf(json.Unmarshal),
		}
		storeFuncs("encoding/json", jsonFuncs)
	}()
}
