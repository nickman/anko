package packages

import (
	"io"
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

		ioFuncs := map[string]reflect.Value{
			"Copy":             reflect.ValueOf(io.Copy),
			"CopyN":            reflect.ValueOf(io.CopyN),
			"EOF":              reflect.ValueOf(io.EOF),
			"ErrClosedPipe":    reflect.ValueOf(io.ErrClosedPipe),
			"ErrNoProgress":    reflect.ValueOf(io.ErrNoProgress),
			"ErrShortBuffer":   reflect.ValueOf(io.ErrShortBuffer),
			"ErrShortWrite":    reflect.ValueOf(io.ErrShortWrite),
			"ErrUnexpectedEOF": reflect.ValueOf(io.ErrUnexpectedEOF),
			"LimitReader":      reflect.ValueOf(io.LimitReader),
			"MultiReader":      reflect.ValueOf(io.MultiReader),
			"MultiWriter":      reflect.ValueOf(io.MultiWriter),
			"NewSectionReader": reflect.ValueOf(io.NewSectionReader),
			"Pipe":             reflect.ValueOf(io.Pipe),
			"ReadAtLeast":      reflect.ValueOf(io.ReadAtLeast),
			"ReadFull":         reflect.ValueOf(io.ReadFull),
			"TeeReader":        reflect.ValueOf(io.TeeReader),
			"WriteString":      reflect.ValueOf(io.WriteString),
		}
		storeFuncs("io", ioFuncs)
	}()
}
