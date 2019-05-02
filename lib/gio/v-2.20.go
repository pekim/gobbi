// Code generated - DO NOT EDIT.
// +build gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
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
/*

	void mountoperation_abortedHandler(GObject *, gpointer);

	static gulong MountOperation_signal_connect_aborted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "aborted", G_CALLBACK(mountoperation_abortedHandler), data);
	}

*/
import "C"

// Unsupported : g_data_input_stream_read_line_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_data_input_stream_read_line_finish : array return type :

// Unsupported : g_data_input_stream_read_until_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReadUntilFinish is a wrapper around the C function g_data_input_stream_read_until_finish.
func (recv *DataInputStream) ReadUntilFinish(result *AsyncResult) (string, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_until_finish((*C.GDataInputStream)(recv.native), c_result, &c_length, &cThrowableError)
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

type signalMountOperationAbortedDetail struct {
	callback  MountOperationSignalAbortedCallback
	handlerID C.gulong
}

var signalMountOperationAbortedId int
var signalMountOperationAbortedMap = make(map[int]signalMountOperationAbortedDetail)
var signalMountOperationAbortedLock sync.RWMutex

// MountOperationSignalAbortedCallback is a callback function for a 'aborted' signal emitted from a MountOperation.
type MountOperationSignalAbortedCallback func()

/*
ConnectAborted connects the callback to the 'aborted' signal for the MountOperation.

The returned value represents the connection, and may be passed to DisconnectAborted to remove it.
*/
func (recv *MountOperation) ConnectAborted(callback MountOperationSignalAbortedCallback) int {
	signalMountOperationAbortedLock.Lock()
	defer signalMountOperationAbortedLock.Unlock()

	signalMountOperationAbortedId++
	instance := C.gpointer(recv.native)
	handlerID := C.MountOperation_signal_connect_aborted(instance, C.gpointer(uintptr(signalMountOperationAbortedId)))

	detail := signalMountOperationAbortedDetail{callback, handlerID}
	signalMountOperationAbortedMap[signalMountOperationAbortedId] = detail

	return signalMountOperationAbortedId
}

/*
DisconnectAborted disconnects a callback from the 'aborted' signal for the MountOperation.

The connectionID should be a value returned from a call to ConnectAborted.
*/
func (recv *MountOperation) DisconnectAborted(connectionID int) {
	signalMountOperationAbortedLock.Lock()
	defer signalMountOperationAbortedLock.Unlock()

	detail, exists := signalMountOperationAbortedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountOperationAbortedMap, connectionID)
}

//export mountoperation_abortedHandler
func mountoperation_abortedHandler(_ *C.GObject, data C.gpointer) {
	signalMountOperationAbortedLock.RLock()
	defer signalMountOperationAbortedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalMountOperationAbortedMap[index].callback
	callback()
}

// g_simple_async_result_is_valid : unsupported parameter source_tag : no type generator for gpointer (gpointer) for param source_tag
// GetCloseFd is a wrapper around the C function g_unix_input_stream_get_close_fd.
func (recv *UnixInputStream) GetCloseFd() bool {
	retC := C.g_unix_input_stream_get_close_fd((*C.GUnixInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFd is a wrapper around the C function g_unix_input_stream_get_fd.
func (recv *UnixInputStream) GetFd() int32 {
	retC := C.g_unix_input_stream_get_fd((*C.GUnixInputStream)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetCloseFd is a wrapper around the C function g_unix_input_stream_set_close_fd.
func (recv *UnixInputStream) SetCloseFd(closeFd bool) {
	c_close_fd :=
		boolToGboolean(closeFd)

	C.g_unix_input_stream_set_close_fd((*C.GUnixInputStream)(recv.native), c_close_fd)

	return
}

// GetCloseFd is a wrapper around the C function g_unix_output_stream_get_close_fd.
func (recv *UnixOutputStream) GetCloseFd() bool {
	retC := C.g_unix_output_stream_get_close_fd((*C.GUnixOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFd is a wrapper around the C function g_unix_output_stream_get_fd.
func (recv *UnixOutputStream) GetFd() int32 {
	retC := C.g_unix_output_stream_get_fd((*C.GUnixOutputStream)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetCloseFd is a wrapper around the C function g_unix_output_stream_set_close_fd.
func (recv *UnixOutputStream) SetCloseFd(closeFd bool) {
	c_close_fd :=
		boolToGboolean(closeFd)

	C.g_unix_output_stream_set_close_fd((*C.GUnixOutputStream)(recv.native), c_close_fd)

	return
}

const FILE_ATTRIBUTE_PREVIEW_ICON string = C.G_FILE_ATTRIBUTE_PREVIEW_ICON
const FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE string = C.G_FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE

// AppInfoResetTypeAssociations is a wrapper around the C function g_app_info_reset_type_associations.
func AppInfoResetTypeAssociations(contentType string) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	C.g_app_info_reset_type_associations(c_content_type)

	return
}

// CanDelete is a wrapper around the C function g_app_info_can_delete.
func (recv *AppInfo) CanDelete() bool {
	retC := C.g_app_info_can_delete((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Delete is a wrapper around the C function g_app_info_delete.
func (recv *AppInfo) Delete() bool {
	retC := C.g_app_info_delete((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCommandline is a wrapper around the C function g_app_info_get_commandline.
func (recv *AppInfo) GetCommandline() string {
	retC := C.g_app_info_get_commandline((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// IconNewForString is a wrapper around the C function g_icon_new_for_string.
func IconNewForString(str string) (*Icon, error) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	var cThrowableError *C.GError

	retC := C.g_icon_new_for_string(c_str, &cThrowableError)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ToString is a wrapper around the C function g_icon_to_string.
func (recv *Icon) ToString() string {
	retC := C.g_icon_to_string((*C.GIcon)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// IsShadowed is a wrapper around the C function g_mount_is_shadowed.
func (recv *Mount) IsShadowed() bool {
	retC := C.g_mount_is_shadowed((*C.GMount)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Shadow is a wrapper around the C function g_mount_shadow.
func (recv *Mount) Shadow() {
	C.g_mount_shadow((*C.GMount)(recv.native))

	return
}

// Unshadow is a wrapper around the C function g_mount_unshadow.
func (recv *Mount) Unshadow() {
	C.g_mount_unshadow((*C.GMount)(recv.native))

	return
}
