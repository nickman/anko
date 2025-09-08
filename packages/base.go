package packages

import (
	"reflect"
	"sync"
	"time"

	sm "github.com/chloyka/sync-map-generic"
	"github.com/mattn/anko/env"
)

var (
	completion = new(sync.WaitGroup)
	initCount  = new(int32)
	initStart  = time.Now()
)

func vof(i any) *reflect.Value {
	v := reflect.ValueOf(i)
	return &v
}

func tof(i any) *reflect.Type {
	t := reflect.TypeOf(i)
	return &t
}

func storeFuncs(name string, funcs map[string]reflect.Value) {
	methods, _ := env.Packages.LoadOrStore(name, new(sm.KVMap[string, reflect.Value]))
	for k, v := range funcs {
		methods.Store(k, vof(v))
	}
}

func storeTypes(name string, funcs map[string]reflect.Type) {
	types, _ := env.PackageTypes.LoadOrStore(name, new(sm.KVMap[string, reflect.Type]))
	for k, v := range funcs {
		types.Store(k, tof(v))
	}
}
