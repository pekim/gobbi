// This is a generated file - DO NOT EDIT
// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// Once is a wrapper around the C record GOnce.
type Once struct {
	native unsafe.Pointer
	Status OnceStatus
	// retval : no type generator for gpointer, volatile gpointer
}

func OnceNewFromC(u unsafe.Pointer) *Once {
	if u == nil {
		return nil
	}

	g := &Once{native: u}

	return g
}

func (recv *Once) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
