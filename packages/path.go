package packages

import (
	"path"
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
			"Base":          reflect.ValueOf(path.Base),
			"Clean":         reflect.ValueOf(path.Clean),
			"Dir":           reflect.ValueOf(path.Dir),
			"ErrBadPattern": reflect.ValueOf(path.ErrBadPattern),
			"Ext":           reflect.ValueOf(path.Ext),
			"IsAbs":         reflect.ValueOf(path.IsAbs),
			"Join":          reflect.ValueOf(path.Join),
			"Match":         reflect.ValueOf(path.Match),
			"Split":         reflect.ValueOf(path.Split),
		}
		storeFuncs("path", funcs)
	}()
}
