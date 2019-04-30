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

func AsyncInitableIfaceNewFromC(u unsafe.Pointer) *AsyncInitableIface {
	if u == nil {
		return nil
	}

	g := &AsyncInitableIface{native: u}

	return g
}

func (recv *AsyncInitableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitableIface is a wrapper around the C record GInitableIface.
type InitableIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for init
}

func InitableIfaceNewFromC(u unsafe.Pointer) *InitableIface {
	if u == nil {
		return nil
	}

	g := &InitableIface{native: u}

	return g
}

func (recv *InitableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputVector is a wrapper around the C record GInputVector.
type InputVector struct {
	native unsafe.Pointer
	// buffer : no type generator for gpointer, gpointer
	Size uint64
}

func InputVectorNewFromC(u unsafe.Pointer) *InputVector {
	if u == nil {
		return nil
	}

	g := &InputVector{native: u}

	return g
}

func (recv *InputVector) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OutputVector is a wrapper around the C record GOutputVector.
type OutputVector struct {
	native unsafe.Pointer
	// buffer : no type generator for gpointer, gconstpointer
	Size uint64
}

func OutputVectorNewFromC(u unsafe.Pointer) *OutputVector {
	if u == nil {
		return nil
	}

	g := &OutputVector{native: u}

	return g
}

func (recv *OutputVector) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_srv_target_new : return type :
