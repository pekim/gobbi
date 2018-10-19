// This is a generated file - DO NOT EDIT
// +build gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"fmt"
	"sync"
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

	void MountOperation_abortedHandler();

	static gulong MountOperation_signal_connect_aborted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "aborted", MountOperation_abortedHandler, data);
	}

*/
import "C"

// Unsupported signal 'launched' for AppLaunchContext : unsupported parameter info : no type generator for AppInfo,

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_data_input_stream_read_line_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_data_input_stream_read_line_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_data_input_stream_read_until_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_data_input_stream_read_until_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported signal 'changed' for FileMonitor : unsupported parameter file : no type generator for File,

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

var signalMountOperationAbortedId int
var signalMountOperationAbortedMap = make(map[int]MountOperationSignalAbortedCallback)
var signalMountOperationAbortedLock sync.Mutex

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
	signalMountOperationAbortedMap[signalMountOperationAbortedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.MountOperation_signal_connect_aborted(instance, C.gpointer(uintptr(signalMountOperationAbortedId)))
	return int(retC)
}

/*
DisconnectAborted disconnects a callback from the 'aborted' signal for the MountOperation.

The connectionID should be a value returned from a call to ConnectAborted.
*/
func (recv *MountOperation) DisconnectAborted(connectionID int) {
	signalMountOperationAbortedLock.Lock()
	defer signalMountOperationAbortedLock.Unlock()

	_, exists := signalMountOperationAbortedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalMountOperationAbortedMap, connectionID)
}

//export MountOperation_abortedHandler
func MountOperation_abortedHandler() {
	fmt.Println("cb")
}

// Unsupported signal 'ask-question' for MountOperation : unsupported parameter choices : no param type

// Unsupported signal 'show-processes' for MountOperation : unsupported parameter processes : no param type

// Unsupported signal 'change-event' for Settings : unsupported parameter keys : no param type

// Unsupported signal 'activate' for SimpleAction : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported signal 'change-state' for SimpleAction : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

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

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type

// Unsupported signal 'drive-changed' for VolumeMonitor : unsupported parameter drive : no type generator for Drive,

// Unsupported signal 'drive-connected' for VolumeMonitor : unsupported parameter drive : no type generator for Drive,

// Unsupported signal 'drive-disconnected' for VolumeMonitor : unsupported parameter drive : no type generator for Drive,

// Unsupported signal 'drive-eject-button' for VolumeMonitor : unsupported parameter drive : no type generator for Drive,

// Unsupported signal 'drive-stop-button' for VolumeMonitor : unsupported parameter drive : no type generator for Drive,

// Unsupported signal 'mount-added' for VolumeMonitor : unsupported parameter mount : no type generator for Mount,

// Unsupported signal 'mount-changed' for VolumeMonitor : unsupported parameter mount : no type generator for Mount,

// Unsupported signal 'mount-pre-unmount' for VolumeMonitor : unsupported parameter mount : no type generator for Mount,

// Unsupported signal 'mount-removed' for VolumeMonitor : unsupported parameter mount : no type generator for Mount,

// Unsupported signal 'volume-added' for VolumeMonitor : unsupported parameter volume : no type generator for Volume,

// Unsupported signal 'volume-changed' for VolumeMonitor : unsupported parameter volume : no type generator for Volume,

// Unsupported signal 'volume-removed' for VolumeMonitor : unsupported parameter volume : no type generator for Volume,
