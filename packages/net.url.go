//go:build !appengine
// +build !appengine

package packages

import (
	"net/url"
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
			"QueryEscape":     reflect.ValueOf(url.QueryEscape),
			"QueryUnescape":   reflect.ValueOf(url.QueryUnescape),
			"Parse":           reflect.ValueOf(url.Parse),
			"ParseRequestURI": reflect.ValueOf(url.ParseRequestURI),
			"User":            reflect.ValueOf(url.User),
			"UserPassword":    reflect.ValueOf(url.UserPassword),
			"ParseQuery":      reflect.ValueOf(url.ParseQuery),
		}
		storeFuncs("net/url", funcs)
		//env.PackageTypes["net/url"]
		types := map[string]reflect.Type{
			"Error":  reflect.TypeOf(url.Error{}),
			"URL":    reflect.TypeOf(url.URL{}),
			"Values": reflect.TypeOf(url.Values{}),
		}
		storeTypes("net/url", types)
	}()
}
