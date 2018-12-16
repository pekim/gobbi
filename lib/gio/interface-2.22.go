// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

	void drive_stopButtonHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_stop_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "stop-button", G_CALLBACK(drive_stopButtonHandler), data);
	}

*/
/*

	void mount_preUnmountHandler(GObject *, gpointer);

	static gulong Mount_signal_connect_pre_unmount(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pre-unmount", G_CALLBACK(mount_preUnmountHandler), data);
	}

*/
import "C"

// AsyncInitable is a wrapper around the C record GAsyncInitable.
type AsyncInitable struct {
	native *C.GAsyncInitable
}

func AsyncInitableNewFromC(u unsafe.Pointer) *AsyncInitable {
	c := (*C.GAsyncInitable)(u)
	if c == nil {
		return nil
	}

	g := &AsyncInitable{native: c}

	return g
}

func (recv *AsyncInitable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AsyncInitable with another AsyncInitable, and returns true if they represent the same GObject.
func (recv *AsyncInitable) Equals(other *AsyncInitable) bool {
	return other.ToC() == recv.ToC()
}

// g_async_initable_new_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// g_async_initable_new_valist_async : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args
// g_async_initable_newv_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// Unsupported : g_async_initable_init_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// InitFinish is a wrapper around the C function g_async_initable_init_finish.
func (recv *AsyncInitable) InitFinish(res *AsyncResult) (bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_async_initable_init_finish((*C.GAsyncInitable)(recv.native), c_res, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// NewFinish is a wrapper around the C function g_async_initable_new_finish.
func (recv *AsyncInitable) NewFinish(res *AsyncResult) (*gobject.Object, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_async_initable_new_finish((*C.GAsyncInitable)(recv.native), c_res, &cThrowableError)
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

type signalDriveStopButtonDetail struct {
	callback  DriveSignalStopButtonCallback
	handlerID C.gulong
}

var signalDriveStopButtonId int
var signalDriveStopButtonMap = make(map[int]signalDriveStopButtonDetail)
var signalDriveStopButtonLock sync.RWMutex

// DriveSignalStopButtonCallback is a callback function for a 'stop-button' signal emitted from a Drive.
type DriveSignalStopButtonCallback func()

/*
ConnectStopButton connects the callback to the 'stop-button' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectStopButton to remove it.
*/
func (recv *Drive) ConnectStopButton(callback DriveSignalStopButtonCallback) int {
	signalDriveStopButtonLock.Lock()
	defer signalDriveStopButtonLock.Unlock()

	signalDriveStopButtonId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_stop_button(instance, C.gpointer(uintptr(signalDriveStopButtonId)))

	detail := signalDriveStopButtonDetail{callback, handlerID}
	signalDriveStopButtonMap[signalDriveStopButtonId] = detail

	return signalDriveStopButtonId
}

/*
DisconnectStopButton disconnects a callback from the 'stop-button' signal for the Drive.

The connectionID should be a value returned from a call to ConnectStopButton.
*/
func (recv *Drive) DisconnectStopButton(connectionID int) {
	signalDriveStopButtonLock.Lock()
	defer signalDriveStopButtonLock.Unlock()

	detail, exists := signalDriveStopButtonMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveStopButtonMap, connectionID)
}

//export drive_stopButtonHandler
func drive_stopButtonHandler(_ *C.GObject, data C.gpointer) {
	signalDriveStopButtonLock.RLock()
	defer signalDriveStopButtonLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDriveStopButtonMap[index].callback
	callback()
}

// CanStart is a wrapper around the C function g_drive_can_start.
func (recv *Drive) CanStart() bool {
	retC := C.g_drive_can_start((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanStartDegraded is a wrapper around the C function g_drive_can_start_degraded.
func (recv *Drive) CanStartDegraded() bool {
	retC := C.g_drive_can_start_degraded((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanStop is a wrapper around the C function g_drive_can_stop.
func (recv *Drive) CanStop() bool {
	retC := C.g_drive_can_stop((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_drive_eject_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectWithOperationFinish is a wrapper around the C function g_drive_eject_with_operation_finish.
func (recv *Drive) EjectWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_eject_with_operation_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetStartStopType is a wrapper around the C function g_drive_get_start_stop_type.
func (recv *Drive) GetStartStopType() DriveStartStopType {
	retC := C.g_drive_get_start_stop_type((*C.GDrive)(recv.native))
	retGo := (DriveStartStopType)(retC)

	return retGo
}

// Unsupported : g_drive_start : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// StartFinish is a wrapper around the C function g_drive_start_finish.
func (recv *Drive) StartFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_start_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_drive_stop : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// StopFinish is a wrapper around the C function g_drive_stop_finish.
func (recv *Drive) StopFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_stop_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CreateReadwrite is a wrapper around the C function g_file_create_readwrite.
func (recv *File) CreateReadwrite(flags FileCreateFlags, cancellable *Cancellable) (*FileIOStream, error) {
	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_create_readwrite((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_create_readwrite_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CreateReadwriteFinish is a wrapper around the C function g_file_create_readwrite_finish.
func (recv *File) CreateReadwriteFinish(res *AsyncResult) (*FileIOStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_create_readwrite_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_eject_mountable_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectMountableWithOperationFinish is a wrapper around the C function g_file_eject_mountable_with_operation_finish.
func (recv *File) EjectMountableWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_eject_mountable_with_operation_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// OpenReadwrite is a wrapper around the C function g_file_open_readwrite.
func (recv *File) OpenReadwrite(cancellable *Cancellable) (*FileIOStream, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_open_readwrite((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_open_readwrite_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// OpenReadwriteFinish is a wrapper around the C function g_file_open_readwrite_finish.
func (recv *File) OpenReadwriteFinish(res *AsyncResult) (*FileIOStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_open_readwrite_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_poll_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// PollMountableFinish is a wrapper around the C function g_file_poll_mountable_finish.
func (recv *File) PollMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_poll_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReplaceReadwrite is a wrapper around the C function g_file_replace_readwrite.
func (recv *File) ReplaceReadwrite(etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable) (*FileIOStream, error) {
	c_etag := C.CString(etag)
	defer C.free(unsafe.Pointer(c_etag))

	c_make_backup :=
		boolToGboolean(makeBackup)

	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_replace_readwrite((*C.GFile)(recv.native), c_etag, c_make_backup, c_flags, c_cancellable, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_replace_readwrite_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReplaceReadwriteFinish is a wrapper around the C function g_file_replace_readwrite_finish.
func (recv *File) ReplaceReadwriteFinish(res *AsyncResult) (*FileIOStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_replace_readwrite_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_start_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// StartMountableFinish is a wrapper around the C function g_file_start_mountable_finish.
func (recv *File) StartMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_start_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_stop_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// StopMountableFinish is a wrapper around the C function g_file_stop_mountable_finish.
func (recv *File) StopMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_stop_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SupportsThreadContexts is a wrapper around the C function g_file_supports_thread_contexts.
func (recv *File) SupportsThreadContexts() bool {
	retC := C.g_file_supports_thread_contexts((*C.GFile)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_file_unmount_mountable_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// UnmountMountableWithOperationFinish is a wrapper around the C function g_file_unmount_mountable_with_operation_finish.
func (recv *File) UnmountMountableWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_unmount_mountable_with_operation_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Initable is a wrapper around the C record GInitable.
type Initable struct {
	native *C.GInitable
}

func InitableNewFromC(u unsafe.Pointer) *Initable {
	c := (*C.GInitable)(u)
	if c == nil {
		return nil
	}

	g := &Initable{native: c}

	return g
}

func (recv *Initable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Initable with another Initable, and returns true if they represent the same GObject.
func (recv *Initable) Equals(other *Initable) bool {
	return other.ToC() == recv.ToC()
}

// g_initable_new : unsupported parameter ... : varargs
// g_initable_new_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args
// g_initable_newv : unsupported parameter parameters :
// Init is a wrapper around the C function g_initable_init.
func (recv *Initable) Init(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_initable_init((*C.GInitable)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

type signalMountPreUnmountDetail struct {
	callback  MountSignalPreUnmountCallback
	handlerID C.gulong
}

var signalMountPreUnmountId int
var signalMountPreUnmountMap = make(map[int]signalMountPreUnmountDetail)
var signalMountPreUnmountLock sync.RWMutex

// MountSignalPreUnmountCallback is a callback function for a 'pre-unmount' signal emitted from a Mount.
type MountSignalPreUnmountCallback func()

/*
ConnectPreUnmount connects the callback to the 'pre-unmount' signal for the Mount.

The returned value represents the connection, and may be passed to DisconnectPreUnmount to remove it.
*/
func (recv *Mount) ConnectPreUnmount(callback MountSignalPreUnmountCallback) int {
	signalMountPreUnmountLock.Lock()
	defer signalMountPreUnmountLock.Unlock()

	signalMountPreUnmountId++
	instance := C.gpointer(recv.native)
	handlerID := C.Mount_signal_connect_pre_unmount(instance, C.gpointer(uintptr(signalMountPreUnmountId)))

	detail := signalMountPreUnmountDetail{callback, handlerID}
	signalMountPreUnmountMap[signalMountPreUnmountId] = detail

	return signalMountPreUnmountId
}

/*
DisconnectPreUnmount disconnects a callback from the 'pre-unmount' signal for the Mount.

The connectionID should be a value returned from a call to ConnectPreUnmount.
*/
func (recv *Mount) DisconnectPreUnmount(connectionID int) {
	signalMountPreUnmountLock.Lock()
	defer signalMountPreUnmountLock.Unlock()

	detail, exists := signalMountPreUnmountMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountPreUnmountMap, connectionID)
}

//export mount_preUnmountHandler
func mount_preUnmountHandler(_ *C.GObject, data C.gpointer) {
	signalMountPreUnmountLock.RLock()
	defer signalMountPreUnmountLock.RUnlock()

	index := int(uintptr(data))
	callback := signalMountPreUnmountMap[index].callback
	callback()
}

// Unsupported : g_mount_eject_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectWithOperationFinish is a wrapper around the C function g_mount_eject_with_operation_finish.
func (recv *Mount) EjectWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_eject_with_operation_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_mount_unmount_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// UnmountWithOperationFinish is a wrapper around the C function g_mount_unmount_with_operation_finish.
func (recv *Mount) UnmountWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_unmount_with_operation_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Enumerate is a wrapper around the C function g_socket_connectable_enumerate.
func (recv *SocketConnectable) Enumerate() *SocketAddressEnumerator {
	retC := C.g_socket_connectable_enumerate((*C.GSocketConnectable)(recv.native))
	retGo := SocketAddressEnumeratorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_volume_eject_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectWithOperationFinish is a wrapper around the C function g_volume_eject_with_operation_finish.
func (recv *Volume) EjectWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_volume_eject_with_operation_finish((*C.GVolume)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}
