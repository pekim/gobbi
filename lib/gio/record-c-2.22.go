// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// AsyncInitableIface is a wrapper around the C record GAsyncInitableIface.
type AsyncInitableIface struct {
	native *C.GAsyncInitableIface
	// g_iface : record
	// no type for init_async
	// no type for init_finish
}

func AsyncInitableIfaceNewFromC(u unsafe.Pointer) *AsyncInitableIface {
	c := (*C.GAsyncInitableIface)(u)
	if c == nil {
		return nil
	}

	g := &AsyncInitableIface{native: c}

	return g
}

func (recv *AsyncInitableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitableIface is a wrapper around the C record GInitableIface.
type InitableIface struct {
	native *C.GInitableIface
	// g_iface : record
	// no type for init
}

func InitableIfaceNewFromC(u unsafe.Pointer) *InitableIface {
	c := (*C.GInitableIface)(u)
	if c == nil {
		return nil
	}

	g := &InitableIface{native: c}

	return g
}

func (recv *InitableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputVector is a wrapper around the C record GInputVector.
type InputVector struct {
	native *C.GInputVector
	// buffer : no type generator for gpointer, gpointer
	Size uint64
}

func InputVectorNewFromC(u unsafe.Pointer) *InputVector {
	c := (*C.GInputVector)(u)
	if c == nil {
		return nil
	}

	g := &InputVector{
		Size:   (uint64)(c.size),
		native: c,
	}

	return g
}

func (recv *InputVector) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OutputVector is a wrapper around the C record GOutputVector.
type OutputVector struct {
	native *C.GOutputVector
	// buffer : no type generator for gpointer, gconstpointer
	Size uint64
}

func OutputVectorNewFromC(u unsafe.Pointer) *OutputVector {
	c := (*C.GOutputVector)(u)
	if c == nil {
		return nil
	}

	g := &OutputVector{
		Size:   (uint64)(c.size),
		native: c,
	}

	return g
}

func (recv *OutputVector) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
