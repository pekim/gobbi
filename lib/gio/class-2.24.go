// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// Unsupported signal : unsupported parameter info : no type generator for AppInfo,

// CharsetConverterNew is a wrapper around the C function g_charset_converter_new.
func CharsetConverterNew(toCharset string, fromCharset string) (*CharsetConverter, error) {
	c_to_charset := C.CString(toCharset)
	defer C.free(unsafe.Pointer(c_to_charset))

	c_from_charset := C.CString(fromCharset)
	defer C.free(unsafe.Pointer(c_from_charset))

	var cThrowableError *C.GError

	retC := C.g_charset_converter_new(c_to_charset, c_from_charset, &cThrowableError)
	retGo := CharsetConverterNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
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

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_input_stream_get_converter : no return generator

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_get_converter : no return generator

// Unsupported : g_data_input_stream_read_upto_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// GetFilename is a wrapper around the C function g_desktop_app_info_get_filename.
func (recv *DesktopAppInfo) GetFilename() string {
	retC := C.g_desktop_app_info_get_filename((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported signal : unsupported parameter file : no type generator for File,

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// Unsupported signal : unsupported parameter choices : no param type

// Unsupported signal : unsupported parameter processes : no param type

// IsClosing is a wrapper around the C function g_output_stream_is_closing.
func (recv *OutputStream) IsClosing() bool {
	retC := C.g_output_stream_is_closing((*C.GOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported signal : unsupported parameter keys : no param type

// Unsupported signal : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported signal : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported signal : unsupported parameter connectable : no type generator for SocketConnectable,

// AddAnyInetPort is a wrapper around the C function g_socket_listener_add_any_inet_port.
func (recv *SocketListener) AddAnyInetPort(sourceObject *gobject.Object) (uint16, error) {
	c_source_object := (*C.GObject)(sourceObject.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_any_inet_port((*C.GSocketListener)(recv.native), c_source_object, &cThrowableError)
	retGo := (uint16)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// UnixFDListNew is a wrapper around the C function g_unix_fd_list_new.
func UnixFDListNew() *UnixFDList {
	retC := C.g_unix_fd_list_new()
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// Append is a wrapper around the C function g_unix_fd_list_append.
func (recv *UnixFDList) Append(fd int32) (int32, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_unix_fd_list_append((*C.GUnixFDList)(recv.native), c_fd, &cThrowableError)
	retGo := (int32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Get is a wrapper around the C function g_unix_fd_list_get.
func (recv *UnixFDList) Get(index int32) (int32, error) {
	c_index_ := (C.gint)(index)

	var cThrowableError *C.GError

	retC := C.g_unix_fd_list_get((*C.GUnixFDList)(recv.native), c_index_, &cThrowableError)
	retGo := (int32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetLength is a wrapper around the C function g_unix_fd_list_get_length.
func (recv *UnixFDList) GetLength() int32 {
	retC := C.g_unix_fd_list_get_length((*C.GUnixFDList)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_unix_fd_list_peek_fds : no return type

// Unsupported : g_unix_fd_list_steal_fds : no return type

// UnixFDMessageNewWithFdList is a wrapper around the C function g_unix_fd_message_new_with_fd_list.
func UnixFDMessageNewWithFdList(fdList *UnixFDList) *UnixFDMessage {
	c_fd_list := (*C.GUnixFDList)(fdList.ToC())

	retC := C.g_unix_fd_message_new_with_fd_list(c_fd_list)
	retGo := UnixFDMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFdList is a wrapper around the C function g_unix_fd_message_get_fd_list.
func (recv *UnixFDMessage) GetFdList() *UnixFDList {
	retC := C.g_unix_fd_message_get_fd_list((*C.GUnixFDMessage)(recv.native))
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter volume : no type generator for Volume,

// Unsupported signal : unsupported parameter volume : no type generator for Volume,

// Unsupported signal : unsupported parameter volume : no type generator for Volume,

// ZlibCompressorNew is a wrapper around the C function g_zlib_compressor_new.
func ZlibCompressorNew(format ZlibCompressorFormat, level int32) *ZlibCompressor {
	c_format := (C.GZlibCompressorFormat)(format)

	c_level := (C.int)(level)

	retC := C.g_zlib_compressor_new(c_format, c_level)
	retGo := ZlibCompressorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ZlibDecompressorNew is a wrapper around the C function g_zlib_decompressor_new.
func ZlibDecompressorNew(format ZlibCompressorFormat) *ZlibDecompressor {
	c_format := (C.GZlibCompressorFormat)(format)

	retC := C.g_zlib_decompressor_new(c_format)
	retGo := ZlibDecompressorNewFromC(unsafe.Pointer(retC))

	return retGo
}
