// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// AppInfoMonitor is a wrapper around the C record GAppInfoMonitor.
type AppInfoMonitor struct {
	native *C.GAppInfoMonitor
}

func AppInfoMonitorNewFromC(u unsafe.Pointer) *AppInfoMonitor {
	c := (*C.GAppInfoMonitor)(u)
	if c == nil {
		return nil
	}

	g := &AppInfoMonitor{native: c}

	return g
}

func (recv *AppInfoMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Notification is a wrapper around the C record GNotification.
type Notification struct {
	native *C.GNotification
}

func NotificationNewFromC(u unsafe.Pointer) *Notification {
	c := (*C.GNotification)(u)
	if c == nil {
		return nil
	}

	g := &Notification{native: c}

	return g
}

func (recv *Notification) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotificationNew is a wrapper around the C function g_notification_new.
func NotificationNew(title string) *Notification {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	retC := C.g_notification_new(c_title)
	retGo := NotificationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_notification_add_button : no return generator

// Unsupported : g_notification_add_button_with_target : unsupported parameter ... : varargs

// Unsupported : g_notification_add_button_with_target_value : unsupported parameter target : Blacklisted record : GVariant

// Unsupported : g_notification_set_body : no return generator

// Unsupported : g_notification_set_default_action : no return generator

// Unsupported : g_notification_set_default_action_and_target : unsupported parameter ... : varargs

// Unsupported : g_notification_set_default_action_and_target_value : unsupported parameter target : Blacklisted record : GVariant

// Unsupported : g_notification_set_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_notification_set_priority : no return generator

// Unsupported : g_notification_set_title : no return generator

// Unsupported : g_notification_set_urgent : no return generator

// Subprocess is a wrapper around the C record GSubprocess.
type Subprocess struct {
	native *C.GSubprocess
}

func SubprocessNewFromC(u unsafe.Pointer) *Subprocess {
	c := (*C.GSubprocess)(u)
	if c == nil {
		return nil
	}

	g := &Subprocess{native: c}

	return g
}

func (recv *Subprocess) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_subprocess_new : unsupported parameter error : in string with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no param type

// Communicate is a wrapper around the C function g_subprocess_communicate.
func (recv *Subprocess) Communicate(stdinBuf *glib.Bytes, cancellable *Cancellable) (bool, **glib.Bytes, **glib.Bytes, error) {
	c_stdin_buf := (*C.GBytes)(stdinBuf.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var c_stdout_buf *C.GBytes

	var c_stderr_buf *C.GBytes

	var cThrowableError *C.GError

	retC := C.g_subprocess_communicate((*C.GSubprocess)(recv.native), c_stdin_buf, c_cancellable, &c_stdout_buf, &c_stderr_buf, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	stdoutBuf := glib.BytesNewFromC(unsafe.Pointer(c_stdout_buf))

	stderrBuf := glib.BytesNewFromC(unsafe.Pointer(c_stderr_buf))

	return retGo, stdoutBuf, stderrBuf, goThrowableError
}

// Unsupported : g_subprocess_communicate_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_subprocess_communicate_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// CommunicateUtf8 is a wrapper around the C function g_subprocess_communicate_utf8.
func (recv *Subprocess) CommunicateUtf8(stdinBuf string, cancellable *Cancellable) (bool, string, string, error) {
	c_stdin_buf := C.CString(stdinBuf)
	defer C.free(unsafe.Pointer(c_stdin_buf))

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var c_stdout_buf *C.char

	var c_stderr_buf *C.char

	var cThrowableError *C.GError

	retC := C.g_subprocess_communicate_utf8((*C.GSubprocess)(recv.native), c_stdin_buf, c_cancellable, &c_stdout_buf, &c_stderr_buf, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	stdoutBuf := C.GoString(c_stdout_buf)
	defer C.free(unsafe.Pointer(c_stdout_buf))

	stderrBuf := C.GoString(c_stderr_buf)
	defer C.free(unsafe.Pointer(c_stderr_buf))

	return retGo, stdoutBuf, stderrBuf, goThrowableError
}

// Unsupported : g_subprocess_communicate_utf8_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_subprocess_communicate_utf8_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_subprocess_force_exit : no return generator

// GetExitStatus is a wrapper around the C function g_subprocess_get_exit_status.
func (recv *Subprocess) GetExitStatus() int32 {
	retC := C.g_subprocess_get_exit_status((*C.GSubprocess)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetIdentifier is a wrapper around the C function g_subprocess_get_identifier.
func (recv *Subprocess) GetIdentifier() string {
	retC := C.g_subprocess_get_identifier((*C.GSubprocess)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetIfExited is a wrapper around the C function g_subprocess_get_if_exited.
func (recv *Subprocess) GetIfExited() bool {
	retC := C.g_subprocess_get_if_exited((*C.GSubprocess)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIfSignaled is a wrapper around the C function g_subprocess_get_if_signaled.
func (recv *Subprocess) GetIfSignaled() bool {
	retC := C.g_subprocess_get_if_signaled((*C.GSubprocess)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetStatus is a wrapper around the C function g_subprocess_get_status.
func (recv *Subprocess) GetStatus() int32 {
	retC := C.g_subprocess_get_status((*C.GSubprocess)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetStderrPipe is a wrapper around the C function g_subprocess_get_stderr_pipe.
func (recv *Subprocess) GetStderrPipe() *InputStream {
	retC := C.g_subprocess_get_stderr_pipe((*C.GSubprocess)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStdinPipe is a wrapper around the C function g_subprocess_get_stdin_pipe.
func (recv *Subprocess) GetStdinPipe() *OutputStream {
	retC := C.g_subprocess_get_stdin_pipe((*C.GSubprocess)(recv.native))
	retGo := OutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStdoutPipe is a wrapper around the C function g_subprocess_get_stdout_pipe.
func (recv *Subprocess) GetStdoutPipe() *InputStream {
	retC := C.g_subprocess_get_stdout_pipe((*C.GSubprocess)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSuccessful is a wrapper around the C function g_subprocess_get_successful.
func (recv *Subprocess) GetSuccessful() bool {
	retC := C.g_subprocess_get_successful((*C.GSubprocess)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTermSig is a wrapper around the C function g_subprocess_get_term_sig.
func (recv *Subprocess) GetTermSig() int32 {
	retC := C.g_subprocess_get_term_sig((*C.GSubprocess)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_subprocess_send_signal : no return generator

// Wait is a wrapper around the C function g_subprocess_wait.
func (recv *Subprocess) Wait(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_subprocess_wait((*C.GSubprocess)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_subprocess_wait_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// WaitCheck is a wrapper around the C function g_subprocess_wait_check.
func (recv *Subprocess) WaitCheck(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_subprocess_wait_check((*C.GSubprocess)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_subprocess_wait_check_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_subprocess_wait_check_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_subprocess_wait_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// SubprocessLauncher is a wrapper around the C record GSubprocessLauncher.
type SubprocessLauncher struct {
	native *C.GSubprocessLauncher
}

func SubprocessLauncherNewFromC(u unsafe.Pointer) *SubprocessLauncher {
	c := (*C.GSubprocessLauncher)(u)
	if c == nil {
		return nil
	}

	g := &SubprocessLauncher{native: c}

	return g
}

func (recv *SubprocessLauncher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SubprocessLauncherNew is a wrapper around the C function g_subprocess_launcher_new.
func SubprocessLauncherNew(flags SubprocessFlags) *SubprocessLauncher {
	c_flags := (C.GSubprocessFlags)(flags)

	retC := C.g_subprocess_launcher_new(c_flags)
	retGo := SubprocessLauncherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Getenv is a wrapper around the C function g_subprocess_launcher_getenv.
func (recv *SubprocessLauncher) Getenv(variable string) string {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	retC := C.g_subprocess_launcher_getenv((*C.GSubprocessLauncher)(recv.native), c_variable)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_subprocess_launcher_set_child_setup : unsupported parameter child_setup : no type generator for GLib.SpawnChildSetupFunc, GSpawnChildSetupFunc

// Unsupported : g_subprocess_launcher_set_cwd : no return generator

// Unsupported : g_subprocess_launcher_set_environ : unsupported parameter env : no param type

// Unsupported : g_subprocess_launcher_set_flags : no return generator

// Unsupported : g_subprocess_launcher_set_stderr_file_path : no return generator

// Unsupported : g_subprocess_launcher_set_stdin_file_path : no return generator

// Unsupported : g_subprocess_launcher_set_stdout_file_path : no return generator

// Unsupported : g_subprocess_launcher_setenv : no return generator

// Unsupported : g_subprocess_launcher_spawn : unsupported parameter error : in string with indirection level of 2

// Unsupported : g_subprocess_launcher_spawnv : unsupported parameter argv : no param type

// Unsupported : g_subprocess_launcher_take_fd : no return generator

// Unsupported : g_subprocess_launcher_take_stderr_fd : no return generator

// Unsupported : g_subprocess_launcher_take_stdin_fd : no return generator

// Unsupported : g_subprocess_launcher_take_stdout_fd : no return generator

// Unsupported : g_subprocess_launcher_unsetenv : no return generator
