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

// GetDisplayName is a wrapper around the C function g_app_info_get_display_name.
func (recv *AppInfo) GetDisplayName() string {
	retC := C.g_app_info_get_display_name((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Converter is a wrapper around the C record GConverter.
type Converter struct {
	native *C.GConverter
}

func ConverterNewFromC(u unsafe.Pointer) *Converter {
	c := (*C.GConverter)(u)
	if c == nil {
		return nil
	}

	g := &Converter{native: c}

	return g
}

func (recv *Converter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_converter_convert : unsupported parameter inbuf : no param type

// Reset is a wrapper around the C function g_converter_reset.
func (recv *Converter) Reset() {
	C.g_converter_reset((*C.GConverter)(recv.native))

	return
}

// Unsupported : g_file_has_parent : unsupported parameter parent : no type generator for File, GFile*

// GetFd is a wrapper around the C function g_file_descriptor_based_get_fd.
func (recv *FileDescriptorBased) GetFd() int32 {
	retC := C.g_file_descriptor_based_get_fd((*C.GFileDescriptorBased)(recv.native))
	retGo := (int32)(retC)

	return retGo
}
