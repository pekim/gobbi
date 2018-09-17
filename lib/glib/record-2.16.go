// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "runtime"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Checksum is a wrapper around the C record GChecksum.
type Checksum struct {
	native *C.GChecksum
}

func checksumNewFromC(c *C.GChecksum, finalizeFree bool) *Checksum {
	if c == nil {
		return nil
	}

	g := &Checksum{native: c}

	if finalizeFree {
		runtime.SetFinalizer(g, func(obj interface{}) {
			C.g_free((C.gpointer)(c))
		})
	}

	return g
}
