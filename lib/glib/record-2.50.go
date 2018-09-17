// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "runtime"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Logfield is a wrapper around the C record GLogField.
type Logfield struct {
	native *C.GLogField
	Key    string
	Value  uintptr
	Length int64
}

func logfieldNewFromC(c *C.GLogField, finalizeFree bool) *Logfield {
	if c == nil {
		return nil
	}

	g := &Logfield{
		Key:    C.GoString(c.key),
		Length: (int64)(c.length),
		Value:  (uintptr)(c.value),
		native: c,
	}

	if finalizeFree {
		runtime.SetFinalizer(g, func(obj interface{}) {
			C.g_free((C.gpointer)(c))
		})
	}

	return g
}
