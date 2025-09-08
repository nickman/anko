package packages

import (
	"path/filepath"
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
			"Abs":          reflect.ValueOf(filepath.Abs),
			"Base":         reflect.ValueOf(filepath.Base),
			"Clean":        reflect.ValueOf(filepath.Clean),
			"Dir":          reflect.ValueOf(filepath.Dir),
			"EvalSymlinks": reflect.ValueOf(filepath.EvalSymlinks),
			"Ext":          reflect.ValueOf(filepath.Ext),
			"FromSlash":    reflect.ValueOf(filepath.FromSlash),
			"Glob":         reflect.ValueOf(filepath.Glob),
			"HasPrefix":    reflect.ValueOf(filepath.HasPrefix),
			"IsAbs":        reflect.ValueOf(filepath.IsAbs),
			"Join":         reflect.ValueOf(filepath.Join),
			"Match":        reflect.ValueOf(filepath.Match),
			"Rel":          reflect.ValueOf(filepath.Rel),
			"Split":        reflect.ValueOf(filepath.Split),
			"SplitList":    reflect.ValueOf(filepath.SplitList),
			"ToSlash":      reflect.ValueOf(filepath.ToSlash),
			"VolumeName":   reflect.ValueOf(filepath.VolumeName),
		}
		storeFuncs("path/filepath", funcs)
	}()
}
