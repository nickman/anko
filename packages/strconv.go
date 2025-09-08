package packages

import (
	"reflect"
	"strconv"
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
			"FormatBool":  reflect.ValueOf(strconv.FormatBool),
			"FormatFloat": reflect.ValueOf(strconv.FormatFloat),
			"FormatInt":   reflect.ValueOf(strconv.FormatInt),
			"FormatUint":  reflect.ValueOf(strconv.FormatUint),
			"ParseBool":   reflect.ValueOf(strconv.ParseBool),
			"ParseFloat":  reflect.ValueOf(strconv.ParseFloat),
			"ParseInt":    reflect.ValueOf(strconv.ParseInt),
			"ParseUint":   reflect.ValueOf(strconv.ParseUint),
			"Atoi":        reflect.ValueOf(strconv.Atoi),
			"Itoa":        reflect.ValueOf(strconv.Itoa),
		}
		storeFuncs("strconv", funcs)
	}()

}
