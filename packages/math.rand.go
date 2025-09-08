package packages

import (
	"math/rand"
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
			"ExpFloat64":  reflect.ValueOf(rand.ExpFloat64),
			"Float32":     reflect.ValueOf(rand.Float32),
			"Float64":     reflect.ValueOf(rand.Float64),
			"Int":         reflect.ValueOf(rand.Int),
			"Int31":       reflect.ValueOf(rand.Int31),
			"Int31n":      reflect.ValueOf(rand.Int31n),
			"Int63":       reflect.ValueOf(rand.Int63),
			"Int63n":      reflect.ValueOf(rand.Int63n),
			"Intn":        reflect.ValueOf(rand.Intn),
			"NormFloat64": reflect.ValueOf(rand.NormFloat64),
			"Perm":        reflect.ValueOf(rand.Perm),
			"Seed":        reflect.ValueOf(rand.Seed),
			"Uint32":      reflect.ValueOf(rand.Uint32),
		}
		storeFuncs("math/rand", funcs)
	}()
}
