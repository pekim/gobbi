// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

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

// CharsetConverterNew is a wrapper around the C function g_charset_converter_new.
func CharsetConverterNew(toCharset string, fromCharset string) (*CharsetConverter, error) {
	c_to_charset := C.CString(toCharset)
	defer C.free(unsafe.Pointer(c_to_charset))

	c_from_charset := C.CString(fromCharset)
	defer C.free(unsafe.Pointer(c_from_charset))

	var cThrowableError *C.GError

	retC := C.g_charset_converter_new(c_to_charset, c_from_charset, &cThrowableError)
	retGo := CharsetConverterNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// UnixFDListNew is a wrapper around the C function g_unix_fd_list_new.
func UnixFDListNew() *UnixFDList {
	retC := C.g_unix_fd_list_new()
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixFDListNewFromArray is a wrapper around the C function g_unix_fd_list_new_from_array.
func UnixFDListNewFromArray(fds []int32) *UnixFDList {
	c_fds_array := make([]C.gint, len(fds)+1, len(fds)+1)
	for i, item := range fds {
		c := (C.gint)(item)
		c_fds_array[i] = c
	}
	c_fds_array[len(fds)] = 0
	c_fds_arrayPtr := &c_fds_array[0]
	c_fds := (*C.gint)(unsafe.Pointer(c_fds_arrayPtr))

	c_n_fds := (C.gint)(len(fds))

	retC := C.g_unix_fd_list_new_from_array(c_fds, c_n_fds)
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixFDMessageNewWithFdList is a wrapper around the C function g_unix_fd_message_new_with_fd_list.
func UnixFDMessageNewWithFdList(fdList *UnixFDList) *UnixFDMessage {
	c_fd_list := (*C.GUnixFDList)(C.NULL)
	if fdList != nil {
		c_fd_list = (*C.GUnixFDList)(fdList.ToC())
	}

	retC := C.g_unix_fd_message_new_with_fd_list(c_fd_list)
	retGo := UnixFDMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ZlibCompressorNew is a wrapper around the C function g_zlib_compressor_new.
func ZlibCompressorNew(format ZlibCompressorFormat, level int32) *ZlibCompressor {
	c_format := (C.GZlibCompressorFormat)(format)

	c_level := (C.int)(level)

	retC := C.g_zlib_compressor_new(c_format, c_level)
	retGo := ZlibCompressorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ZlibDecompressorNew is a wrapper around the C function g_zlib_decompressor_new.
func ZlibDecompressorNew(format ZlibCompressorFormat) *ZlibDecompressor {
	c_format := (C.GZlibCompressorFormat)(format)

	retC := C.g_zlib_decompressor_new(c_format)
	retGo := ZlibDecompressorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

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
