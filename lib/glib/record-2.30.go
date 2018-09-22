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

	g := &Hmac{native: c}

	return g
}

func (recv *Hmac) toC() *C.GHmac {

	return recv.native
}

// Copy is a wrapper around the C function g_hmac_copy.
func (recv *Hmac) Copy() *Hmac {
	retC := C.g_hmac_copy((*C.GHmac)(recv.native))
	retGo := hmacNewFromC(retC)

	return retGo
}

// Unsupported : g_hmac_get_digest : unsupported parameter buffer : no type generator for guint8, guint8*

// GetString is a wrapper around the C function g_hmac_get_string.
func (recv *Hmac) GetString() string {
	retC := C.g_hmac_get_string((*C.GHmac)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Ref is a wrapper around the C function g_hmac_ref.
func (recv *Hmac) Ref() *Hmac {
	retC := C.g_hmac_ref((*C.GHmac)(recv.native))
	retGo := hmacNewFromC(retC)

	return retGo
}

// Unsupported : g_hmac_unref : no return generator

// Unsupported : g_hmac_update : unsupported parameter data : no param type
