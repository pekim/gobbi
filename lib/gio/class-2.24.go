// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// GetNumFallbacks is a wrapper around the C function g_charset_converter_get_num_fallbacks.
func (recv *CharsetConverter) GetNumFallbacks() uint32 {
	retC := C.g_charset_converter_get_num_fallbacks((*C.GCharsetConverter)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetUseFallback is a wrapper around the C function g_charset_converter_get_use_fallback.
func (recv *CharsetConverter) GetUseFallback() bool {
	retC := C.g_charset_converter_get_use_fallback((*C.GCharsetConverter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetUseFallback is a wrapper around the C function g_charset_converter_set_use_fallback.
func (recv *CharsetConverter) SetUseFallback(useFallback bool) {
	c_use_fallback :=
		boolToGboolean(useFallback)

	C.g_charset_converter_set_use_fallback((*C.GCharsetConverter)(recv.native), c_use_fallback)

	return
}

// GetConverter is a wrapper around the C function g_converter_input_stream_get_converter.
func (recv *ConverterInputStream) GetConverter() *Converter {
	retC := C.g_converter_input_stream_get_converter((*C.GConverterInputStream)(recv.native))
	retGo := ConverterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetConverter is a wrapper around the C function g_converter_output_stream_get_converter.
func (recv *ConverterOutputStream) GetConverter() *Converter {
	retC := C.g_converter_output_stream_get_converter((*C.GConverterOutputStream)(recv.native))
	retGo := ConverterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ReadUptoFinish is a wrapper around the C function g_data_input_stream_read_upto_finish.
func (recv *DataInputStream) ReadUptoFinish(result *AsyncResult) (string, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_upto_finish((*C.GDataInputStream)(recv.native), c_result, &c_length, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetFilename is a wrapper around the C function g_desktop_app_info_get_filename.
func (recv *DesktopAppInfo) GetFilename() string {
	retC := C.g_desktop_app_info_get_filename((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : g_io_module_query

// IsClosing is a wrapper around the C function g_output_stream_is_closing.
func (recv *OutputStream) IsClosing() bool {
	retC := C.g_output_stream_is_closing((*C.GOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// AddAnyInetPort is a wrapper around the C function g_socket_listener_add_any_inet_port.
func (recv *SocketListener) AddAnyInetPort(sourceObject *gobject.Object) (uint16, error) {
	c_source_object := (*C.GObject)(C.NULL)
	if sourceObject != nil {
		c_source_object = (*C.GObject)(sourceObject.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_any_inet_port((*C.GSocketListener)(recv.native), c_source_object, &cThrowableError)
	retGo := (uint16)(retC)

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
	c_fds_array := make([]C.gint, len(fds), len(fds))
	c_fds := &c_fds_array[0]

	c_n_fds := (C.gint)(len(fds))

	retC := C.g_unix_fd_list_new_from_array(c_fds, c_n_fds)
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Append is a wrapper around the C function g_unix_fd_list_append.
func (recv *UnixFDList) Append(fd int32) (int32, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_unix_fd_list_append((*C.GUnixFDList)(recv.native), c_fd, &cThrowableError)
	retGo := (int32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Get is a wrapper around the C function g_unix_fd_list_get.
func (recv *UnixFDList) Get(index int32) (int32, error) {
	c_index_ := (C.gint)(index)

	var cThrowableError *C.GError

	retC := C.g_unix_fd_list_get((*C.GUnixFDList)(recv.native), c_index_, &cThrowableError)
	retGo := (int32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetLength is a wrapper around the C function g_unix_fd_list_get_length.
func (recv *UnixFDList) GetLength() int32 {
	retC := C.g_unix_fd_list_get_length((*C.GUnixFDList)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_unix_fd_list_peek_fds : array return type :

// Unsupported : g_unix_fd_list_steal_fds : array return type :

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

// GetFdList is a wrapper around the C function g_unix_fd_message_get_fd_list.
func (recv *UnixFDMessage) GetFdList() *UnixFDList {
	retC := C.g_unix_fd_message_get_fd_list((*C.GUnixFDMessage)(recv.native))
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

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
