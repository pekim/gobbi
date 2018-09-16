// This is a generated file - DO NOT EDIT
// +build glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Hmac is a wrapper around the C record GHmac.
type Hmac struct {
	native *C.GHmac
}

func hmacNewFromC(c *C.GHmac) *Hmac {
	if c == nil {
		return nil
	}

	r := &Hmac{}
	return r
}
