// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// AsyncInitableIface is a wrapper around the C record GAsyncInitableIface.
type AsyncInitableIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for init_async
	// no type for init_finish
}

// InitableIface is a wrapper around the C record GInitableIface.
type InitableIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for init
}

// InputVector is a wrapper around the C record GInputVector.
type InputVector struct {
	native unsafe.Pointer
	// buffer : no type generator for gpointer, gpointer
	Size uint64
}

// OutputVector is a wrapper around the C record GOutputVector.
type OutputVector struct {
	native unsafe.Pointer
	// buffer : no type generator for gpointer, gconstpointer
	Size uint64
}
