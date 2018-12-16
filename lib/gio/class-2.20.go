// This is a generated file - DO NOT EDIT
// +build gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
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
/*

	void mountoperation_abortedHandler(GObject *, gpointer);

	static gulong MountOperation_signal_connect_aborted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "aborted", G_CALLBACK(mountoperation_abortedHandler), data);
	}

*/
import "C"

// Unsupported : g_data_input_stream_read_line_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_data_input_stream_read_line_finish : no return type

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

// SimpleAsyncResultIsValid is a wrapper around the C function g_simple_async_result_is_valid.
func SimpleAsyncResultIsValid(result *AsyncResult, source *gobject.Object, sourceTag uintptr) bool {
	c_result := (*C.GAsyncResult)(result.ToC())

	c_source := (*C.GObject)(C.NULL)
	if source != nil {
		c_source = (*C.GObject)(source.ToC())
	}

	c_source_tag := (C.gpointer)(sourceTag)

	retC := C.g_simple_async_result_is_valid(c_result, c_source, c_source_tag)
	retGo := retC == C.TRUE

	return retGo
}

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
