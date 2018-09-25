// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
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

// ConverterIface is a wrapper around the C record GConverterIface.
type ConverterIface struct {
	native *C.GConverterIface
	// g_iface : record
	// no type for convert
	// no type for reset
}

func ConverterIfaceNewFromC(u unsafe.Pointer) *ConverterIface {
	c := (*C.GConverterIface)(u)
	if c == nil {
		return nil
	}

	g := &ConverterIface{native: c}

	return g
}

func (recv *ConverterIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
