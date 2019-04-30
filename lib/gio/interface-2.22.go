// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// AsyncInitable is a wrapper around the C record GAsyncInitable.
type AsyncInitable struct {
	native unsafe.Pointer
}

func AsyncInitableNewFromC(u unsafe.Pointer) *AsyncInitable {
	if u == nil {
		return nil
	}

	g := &AsyncInitable{native: u}

	return g
}

func (recv *AsyncInitable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Initable is a wrapper around the C record GInitable.
type Initable struct {
	native unsafe.Pointer
}

func InitableNewFromC(u unsafe.Pointer) *Initable {
	if u == nil {
		return nil
	}

	g := &Initable{native: u}

	return g
}

func (recv *Initable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
