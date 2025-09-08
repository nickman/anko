package packages

import (
	"reflect"
	"regexp"
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
			"Match":            reflect.ValueOf(regexp.Match),
			"MatchReader":      reflect.ValueOf(regexp.MatchReader),
			"MatchString":      reflect.ValueOf(regexp.MatchString),
			"QuoteMeta":        reflect.ValueOf(regexp.QuoteMeta),
			"Compile":          reflect.ValueOf(regexp.Compile),
			"CompilePOSIX":     reflect.ValueOf(regexp.CompilePOSIX),
			"MustCompile":      reflect.ValueOf(regexp.MustCompile),
			"MustCompilePOSIX": reflect.ValueOf(regexp.MustCompilePOSIX),
		}
		storeFuncs("regexp", funcs)
	}()
}
