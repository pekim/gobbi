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

// Unsupported : g_async_initable_init_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes asynchronous initialization and returns the result.
// See g_async_initable_init_async().
/*

C function : g_async_initable_init_finish
*/
func (recv *AsyncInitable) InitFinish(res *AsyncResult) (bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_async_initable_init_finish((*C.GAsyncInitable)(recv.native), c_res, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Finishes the async construction for the various g_async_initable_new
// calls, returning the created object or %NULL on error.
/*

C function : g_async_initable_new_finish
*/
func (recv *AsyncInitable) NewFinish(res *AsyncResult) (*gobject.Object, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_async_initable_new_finish((*C.GAsyncInitable)(recv.native), c_res, &cThrowableError)
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

type signalDriveStopButtonDetail struct {
	callback  DriveSignalStopButtonCallback
	handlerID C.gulong
}

var signalDriveStopButtonId int
var signalDriveStopButtonMap = make(map[int]signalDriveStopButtonDetail)
var signalDriveStopButtonLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalDriveStopButtonMap[index].callback
	callback()
}

// Checks if a drive can be started.
/*

C function : g_drive_can_start
*/
func (recv *Drive) CanStart() bool {
	retC := C.g_drive_can_start((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if a drive can be started degraded.
/*

C function : g_drive_can_start_degraded
*/
func (recv *Drive) CanStartDegraded() bool {
	retC := C.g_drive_can_start_degraded((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if a drive can be stopped.
/*

C function : g_drive_can_stop
*/
func (recv *Drive) CanStop() bool {
	retC := C.g_drive_can_stop((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_drive_eject_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes ejecting a drive. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
/*

C function : g_drive_eject_with_operation_finish
*/
func (recv *Drive) EjectWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_eject_with_operation_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets a hint about how a drive can be started/stopped.
/*

C function : g_drive_get_start_stop_type
*/
func (recv *Drive) GetStartStopType() DriveStartStopType {
	retC := C.g_drive_get_start_stop_type((*C.GDrive)(recv.native))
	retGo := (DriveStartStopType)(retC)

	return retGo
}

// Unsupported : g_drive_start : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes starting a drive.
/*

C function : g_drive_start_finish
*/
func (recv *Drive) StartFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_start_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_drive_stop : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes stopping a drive.
/*

C function : g_drive_stop_finish
*/
func (recv *Drive) StopFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_stop_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a new file and returns a stream for reading and
// writing to it. The file must not already exist.
//
// By default files created are generally readable by everyone,
// but if you pass #G_FILE_CREATE_PRIVATE in @flags the file
// will be made readable only to the current user, to the level
// that is supported on the target filesystem.
//
// If @cancellable is not %NULL, then the operation can be cancelled
// by triggering the cancellable object from another thread. If the
// operation was cancelled, the error %G_IO_ERROR_CANCELLED will be
// returned.
//
// If a file or directory with this name already exists, the
// %G_IO_ERROR_EXISTS error will be returned. Some file systems don't
// allow all file names, and may return an %G_IO_ERROR_INVALID_FILENAME
// error, and if the name is too long, %G_IO_ERROR_FILENAME_TOO_LONG
// will be returned. Other errors are possible too, and depend on what
// kind of filesystem the file is on.
//
// Note that in many non-local file cases read and write streams are
// not supported, so make sure you really need to do read and write
// streaming, rather than just opening for reading or writing.
/*

C function : g_file_create_readwrite
*/
func (recv *File) CreateReadwrite(flags FileCreateFlags, cancellable *Cancellable) (*FileIOStream, error) {
	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_create_readwrite((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_create_readwrite_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous file create operation started with
// g_file_create_readwrite_async().
/*

C function : g_file_create_readwrite_finish
*/
func (recv *File) CreateReadwriteFinish(res *AsyncResult) (*FileIOStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_create_readwrite_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_eject_mountable_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous eject operation started by
// g_file_eject_mountable_with_operation().
/*

C function : g_file_eject_mountable_with_operation_finish
*/
func (recv *File) EjectMountableWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_eject_mountable_with_operation_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Opens an existing file for reading and writing. The result is
// a #GFileIOStream that can be used to read and write the contents
// of the file.
//
// If @cancellable is not %NULL, then the operation can be cancelled
// by triggering the cancellable object from another thread. If the
// operation was cancelled, the error %G_IO_ERROR_CANCELLED will be
// returned.
//
// If the file does not exist, the %G_IO_ERROR_NOT_FOUND error will
// be returned. If the file is a directory, the %G_IO_ERROR_IS_DIRECTORY
// error will be returned. Other errors are possible too, and depend on
// what kind of filesystem the file is on. Note that in many non-local
// file cases read and write streams are not supported, so make sure you
// really need to do read and write streaming, rather than just opening
// for reading or writing.
/*

C function : g_file_open_readwrite
*/
func (recv *File) OpenReadwrite(cancellable *Cancellable) (*FileIOStream, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_open_readwrite((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_open_readwrite_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous file read operation started with
// g_file_open_readwrite_async().
/*

C function : g_file_open_readwrite_finish
*/
func (recv *File) OpenReadwriteFinish(res *AsyncResult) (*FileIOStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_open_readwrite_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_poll_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes a poll operation. See g_file_poll_mountable() for details.
//
// Finish an asynchronous poll operation that was polled
// with g_file_poll_mountable().
/*

C function : g_file_poll_mountable_finish
*/
func (recv *File) PollMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_poll_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Returns an output stream for overwriting the file in readwrite mode,
// possibly creating a backup copy of the file first. If the file doesn't
// exist, it will be created.
//
// For details about the behaviour, see g_file_replace() which does the
// same thing but returns an output stream only.
//
// Note that in many non-local file cases read and write streams are not
// supported, so make sure you really need to do read and write streaming,
// rather than just opening for reading or writing.
/*

C function : g_file_replace_readwrite
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_replace_readwrite_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous file replace operation started with
// g_file_replace_readwrite_async().
/*

C function : g_file_replace_readwrite_finish
*/
func (recv *File) ReplaceReadwriteFinish(res *AsyncResult) (*FileIOStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_replace_readwrite_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_start_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes a start operation. See g_file_start_mountable() for details.
//
// Finish an asynchronous start operation that was started
// with g_file_start_mountable().
/*

C function : g_file_start_mountable_finish
*/
func (recv *File) StartMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_start_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_stop_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an stop operation, see g_file_stop_mountable() for details.
//
// Finish an asynchronous stop operation that was started
// with g_file_stop_mountable().
/*

C function : g_file_stop_mountable_finish
*/
func (recv *File) StopMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_stop_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if @file supports
// [thread-default contexts][g-main-context-push-thread-default-context].
// If this returns %FALSE, you cannot perform asynchronous operations on
// @file in a thread that has a thread-default context.
/*

C function : g_file_supports_thread_contexts
*/
func (recv *File) SupportsThreadContexts() bool {
	retC := C.g_file_supports_thread_contexts((*C.GFile)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_file_unmount_mountable_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an unmount operation,
// see g_file_unmount_mountable_with_operation() for details.
//
// Finish an asynchronous unmount operation that was started
// with g_file_unmount_mountable_with_operation().
/*

C function : g_file_unmount_mountable_with_operation_finish
*/
func (recv *File) UnmountMountableWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_unmount_mountable_with_operation_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
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

// Initializes the object implementing the interface.
//
// This method is intended for language bindings. If writing in C,
// g_initable_new() should typically be used instead.
//
// The object must be initialized before any real use after initial
// construction, either with this function or g_async_initable_init_async().
//
// Implementations may also support cancellation. If @cancellable is not %NULL,
// then initialization can be cancelled by triggering the cancellable object
// from another thread. If the operation was cancelled, the error
// %G_IO_ERROR_CANCELLED will be returned. If @cancellable is not %NULL and
// the object doesn't support cancellable initialization the error
// %G_IO_ERROR_NOT_SUPPORTED will be returned.
//
// If the object is not initialized, or initialization returns with an
// error, then all operations on the object except g_object_ref() and
// g_object_unref() are considered to be invalid, and have undefined
// behaviour. See the [introduction][ginitable] for more details.
//
// Callers should not assume that a class which implements #GInitable can be
// initialized multiple times, unless the class explicitly documents itself as
// supporting this. Generally, a class’ implementation of init() can assume
// (and assert) that it will only be called once. Previously, this documentation
// recommended all #GInitable implementations should be idempotent; that
// recommendation was relaxed in GLib 2.54.
//
// If a class explicitly supports being initialized multiple times, it is
// recommended that the method is idempotent: multiple calls with the same
// arguments should return the same results. Only the first call initializes
// the object; further calls return the result of the first call.
//
// One reason why a class might need to support idempotent initialization is if
// it is designed to be used via the singleton pattern, with a
// #GObjectClass.constructor that sometimes returns an existing instance.
// In this pattern, a caller would expect to be able to call g_initable_init()
// on the result of g_object_new(), regardless of whether it is in fact a new
// instance.
/*

C function : g_initable_init
*/
func (recv *Initable) Init(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_initable_init((*C.GInitable)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

type signalMountPreUnmountDetail struct {
	callback  MountSignalPreUnmountCallback
	handlerID C.gulong
}

var signalMountPreUnmountId int
var signalMountPreUnmountMap = make(map[int]signalMountPreUnmountDetail)
var signalMountPreUnmountLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalMountPreUnmountMap[index].callback
	callback()
}

// Unsupported : g_mount_eject_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes ejecting a mount. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
/*

C function : g_mount_eject_with_operation_finish
*/
func (recv *Mount) EjectWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_eject_with_operation_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_mount_unmount_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes unmounting a mount. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
/*

C function : g_mount_unmount_with_operation_finish
*/
func (recv *Mount) UnmountWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_unmount_with_operation_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a #GSocketAddressEnumerator for @connectable.
/*

C function : g_socket_connectable_enumerate
*/
func (recv *SocketConnectable) Enumerate() *SocketAddressEnumerator {
	retC := C.g_socket_connectable_enumerate((*C.GSocketConnectable)(recv.native))
	retGo := SocketAddressEnumeratorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_volume_eject_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes ejecting a volume. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
/*

C function : g_volume_eject_with_operation_finish
*/
func (recv *Volume) EjectWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_volume_eject_with_operation_finish((*C.GVolume)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
