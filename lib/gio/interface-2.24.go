// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// Converter is a wrapper around the C record GConverter.
type Converter struct {
	native unsafe.Pointer
}

func ConverterNewFromC(u unsafe.Pointer) *Converter {
	if u == nil {
		return nil
	}

	g := &Converter{native: u}

	return g
}

func (recv *Converter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
