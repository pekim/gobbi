// This is a generated file - DO NOT EDIT
// +build glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// Hmac is a wrapper around the C record GHmac.
type Hmac struct {
	native unsafe.Pointer
}

func HmacNewFromC(u unsafe.Pointer) *Hmac {
	if u == nil {
		return nil
	}

	g := &Hmac{native: u}

	return g
}

func (recv *Hmac) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
