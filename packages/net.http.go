//go:build !appengine
// +build !appengine

package packages

import (
	"net/http"
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
			"DefaultClient":     reflect.ValueOf(http.DefaultClient),
			"DefaultServeMux":   reflect.ValueOf(http.DefaultServeMux),
			"DefaultTransport":  reflect.ValueOf(http.DefaultTransport),
			"Handle":            reflect.ValueOf(http.Handle),
			"HandleFunc":        reflect.ValueOf(http.HandleFunc),
			"ListenAndServe":    reflect.ValueOf(http.ListenAndServe),
			"ListenAndServeTLS": reflect.ValueOf(http.ListenAndServeTLS),
			"NewRequest":        reflect.ValueOf(http.NewRequest),
			"NewServeMux":       reflect.ValueOf(http.NewServeMux),
			"Serve":             reflect.ValueOf(http.Serve),
			"SetCookie":         reflect.ValueOf(http.SetCookie),
		}
		storeFuncs("net/http", funcs)
		//env.PackageTypes["net/http"]
		types := map[string]reflect.Type{
			"Client":   reflect.TypeOf(http.Client{}),
			"Cookie":   reflect.TypeOf(http.Cookie{}),
			"Request":  reflect.TypeOf(http.Request{}),
			"Response": reflect.TypeOf(http.Response{}),
		}
		storeTypes("net/http", types)
	}()
}
