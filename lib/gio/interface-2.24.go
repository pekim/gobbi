// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

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

// Equals compares this Converter with another Converter, and returns true if they represent the same GObject.
func (recv *Converter) Equals(other *Converter) bool {
	return other.ToC() == recv.ToC()
}

// Convert is a wrapper around the C function g_converter_convert.
func (recv *Converter) Convert(inbuf []uint8, outbuf []uint8, flags ConverterFlags) (ConverterResult, uint64, uint64, error) {
	c_inbuf := &inbuf[0]

	c_inbuf_size := (C.gsize)(len(inbuf))

	c_outbuf := &outbuf[0]

	c_outbuf_size := (C.gsize)(len(outbuf))

	c_flags := (C.GConverterFlags)(flags)

	var c_bytes_read C.gsize

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_converter_convert((*C.GConverter)(recv.native), (unsafe.Pointer(c_inbuf)), c_inbuf_size, (unsafe.Pointer(c_outbuf)), c_outbuf_size, c_flags, &c_bytes_read, &c_bytes_written, &cThrowableError)
	retGo := (ConverterResult)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesRead, bytesWritten, goError
}

// Reset is a wrapper around the C function g_converter_reset.
func (recv *Converter) Reset() {
	C.g_converter_reset((*C.GConverter)(recv.native))

	return
}

// HasParent is a wrapper around the C function g_file_has_parent.
func (recv *File) HasParent(parent *File) bool {
	c_parent := (*C.GFile)(parent.ToC())

	retC := C.g_file_has_parent((*C.GFile)(recv.native), c_parent)
	retGo := retC == C.TRUE

	return retGo
}

// GetFd is a wrapper around the C function g_file_descriptor_based_get_fd.
func (recv *FileDescriptorBased) GetFd() int32 {
	retC := C.g_file_descriptor_based_get_fd((*C.GFileDescriptorBased)(recv.native))
	retGo := (int32)(retC)

	return retGo
}
