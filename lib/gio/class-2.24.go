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

// Creates a new #GCharsetConverter.
/*

C function

g_charset_converter_new
*/
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

// Gets the number of fallbacks that @converter has applied so far.
/*

C function

g_charset_converter_get_num_fallbacks
*/
func (recv *CharsetConverter) GetNumFallbacks() uint32 {
	retC := C.g_charset_converter_get_num_fallbacks((*C.GCharsetConverter)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the #GCharsetConverter:use-fallback property.
/*

C function

g_charset_converter_get_use_fallback
*/
func (recv *CharsetConverter) GetUseFallback() bool {
	retC := C.g_charset_converter_get_use_fallback((*C.GCharsetConverter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the #GCharsetConverter:use-fallback property.
/*

C function

g_charset_converter_set_use_fallback
*/
func (recv *CharsetConverter) SetUseFallback(useFallback bool) {
	c_use_fallback :=
		boolToGboolean(useFallback)

	C.g_charset_converter_set_use_fallback((*C.GCharsetConverter)(recv.native), c_use_fallback)

	return
}

// Gets the #GConverter that is used by @converter_stream.
/*

C function

g_converter_input_stream_get_converter
*/
func (recv *ConverterInputStream) GetConverter() *Converter {
	retC := C.g_converter_input_stream_get_converter((*C.GConverterInputStream)(recv.native))
	retGo := ConverterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GConverter that is used by @converter_stream.
/*

C function

g_converter_output_stream_get_converter
*/
func (recv *ConverterOutputStream) GetConverter() *Converter {
	retC := C.g_converter_output_stream_get_converter((*C.GConverterOutputStream)(recv.native))
	retGo := ConverterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Finish an asynchronous call started by
// g_data_input_stream_read_upto_async().
//
// Note that this function does not consume the stop character. You
// have to use g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto_async() again.
//
// The returned string will always be nul-terminated on success.
/*

C function

g_data_input_stream_read_upto_finish
*/
func (recv *DataInputStream) ReadUptoFinish(result *AsyncResult) (string, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_upto_finish((*C.GDataInputStream)(recv.native), c_result, &c_length, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goThrowableError
}

// When @info was created from a known filename, return it.  In some
// situations such as the #GDesktopAppInfo returned from
// g_desktop_app_info_new_from_keyfile(), this function will return %NULL.
/*

C function

g_desktop_app_info_get_filename
*/
func (recv *DesktopAppInfo) GetFilename() string {
	retC := C.g_desktop_app_info_get_filename((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Checks if an output stream is being closed. This can be
// used inside e.g. a flush implementation to see if the
// flush (or other i/o operation) is called from within
// the closing operation.
/*

C function

g_output_stream_is_closing
*/
func (recv *OutputStream) IsClosing() bool {
	retC := C.g_output_stream_is_closing((*C.GOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Listens for TCP connections on any available port number for both
// IPv6 and IPv4 (if each is available).
//
// This is useful if you need to have a socket for incoming connections
// but don't care about the specific port number.
//
// @source_object will be passed out in the various calls
// to accept to identify this particular source, which is
// useful if you're listening on multiple addresses and do
// different things depending on what address is connected to.
/*

C function

g_socket_listener_add_any_inet_port
*/
func (recv *SocketListener) AddAnyInetPort(sourceObject *gobject.Object) (uint16, error) {
	c_source_object := (*C.GObject)(C.NULL)
	if sourceObject != nil {
		c_source_object = (*C.GObject)(sourceObject.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_any_inet_port((*C.GSocketListener)(recv.native), c_source_object, &cThrowableError)
	retGo := (uint16)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a new #GUnixFDList containing no file descriptors.
/*

C function

g_unix_fd_list_new
*/
func UnixFDListNew() *UnixFDList {
	retC := C.g_unix_fd_list_new()
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GUnixFDList containing the file descriptors given in
// @fds.  The file descriptors become the property of the new list and
// may no longer be used by the caller.  The array itself is owned by
// the caller.
//
// Each file descriptor in the array should be set to close-on-exec.
//
// If @n_fds is -1 then @fds must be terminated with -1.
/*

C function

g_unix_fd_list_new_from_array
*/
func UnixFDListNewFromArray(fds []int32) *UnixFDList {
	c_fds := &fds[0]

	c_n_fds := (C.gint)(len(fds))

	retC := C.g_unix_fd_list_new_from_array((*C.gint)(unsafe.Pointer(c_fds)), c_n_fds)
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a file descriptor to @list.
//
// The file descriptor is duplicated using dup(). You keep your copy
// of the descriptor and the copy contained in @list will be closed
// when @list is finalized.
//
// A possible cause of failure is exceeding the per-process or
// system-wide file descriptor limit.
//
// The index of the file descriptor in the list is returned.  If you use
// this index with g_unix_fd_list_get() then you will receive back a
// duplicated copy of the same file descriptor.
/*

C function

g_unix_fd_list_append
*/
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

// Gets a file descriptor out of @list.
//
// @index_ specifies the index of the file descriptor to get.  It is a
// programmer error for @index_ to be out of range; see
// g_unix_fd_list_get_length().
//
// The file descriptor is duplicated using dup() and set as
// close-on-exec before being returned.  You must call close() on it
// when you are done.
//
// A possible cause of failure is exceeding the per-process or
// system-wide file descriptor limit.
/*

C function

g_unix_fd_list_get
*/
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

// Gets the length of @list (ie: the number of file descriptors
// contained within).
/*

C function

g_unix_fd_list_get_length
*/
func (recv *UnixFDList) GetLength() int32 {
	retC := C.g_unix_fd_list_get_length((*C.GUnixFDList)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_unix_fd_list_peek_fds : no return type

// Unsupported : g_unix_fd_list_steal_fds : no return type

// Creates a new #GUnixFDMessage containing @list.
/*

C function

g_unix_fd_message_new_with_fd_list
*/
func UnixFDMessageNewWithFdList(fdList *UnixFDList) *UnixFDMessage {
	c_fd_list := (*C.GUnixFDList)(C.NULL)
	if fdList != nil {
		c_fd_list = (*C.GUnixFDList)(fdList.ToC())
	}

	retC := C.g_unix_fd_message_new_with_fd_list(c_fd_list)
	retGo := UnixFDMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GUnixFDList contained in @message.  This function does not
// return a reference to the caller, but the returned list is valid for
// the lifetime of @message.
/*

C function

g_unix_fd_message_get_fd_list
*/
func (recv *UnixFDMessage) GetFdList() *UnixFDList {
	retC := C.g_unix_fd_message_get_fd_list((*C.GUnixFDMessage)(recv.native))
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GZlibCompressor.
/*

C function

g_zlib_compressor_new
*/
func ZlibCompressorNew(format ZlibCompressorFormat, level int32) *ZlibCompressor {
	c_format := (C.GZlibCompressorFormat)(format)

	c_level := (C.int)(level)

	retC := C.g_zlib_compressor_new(c_format, c_level)
	retGo := ZlibCompressorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GZlibDecompressor.
/*

C function

g_zlib_decompressor_new
*/
func ZlibDecompressorNew(format ZlibCompressorFormat) *ZlibDecompressor {
	c_format := (C.GZlibCompressorFormat)(format)

	retC := C.g_zlib_decompressor_new(c_format)
	retGo := ZlibDecompressorNewFromC(unsafe.Pointer(retC))

	return retGo
}
