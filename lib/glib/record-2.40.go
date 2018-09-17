// This is a generated file - DO NOT EDIT
// +build glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "runtime"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Variantdict is a wrapper around the C record GVariantDict.
type Variantdict struct {
	native *C.GVariantDict
}

func variantdictNewFromC(c *C.GVariantDict) *Variantdict {
	if c == nil {
		return nil
	}

	g := &Variantdict{native: c}
	runtime.SetFinalizer(g, func(obj interface{}) {
		C.g_free(obj)
	})

	return g
}
