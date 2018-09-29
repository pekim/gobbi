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

// Unsupported : g_application_add_main_option_entries : unsupported parameter entries : no param type

// AddOptionGroup is a wrapper around the C function g_application_add_option_group.
func (recv *Application) AddOptionGroup(group *glib.OptionGroup) {
	c_group := (*C.GOptionGroup)(group.ToC())

	C.g_application_add_option_group((*C.GApplication)(recv.native), c_group)

	return
}

// SendNotification is a wrapper around the C function g_application_send_notification.
func (recv *Application) SendNotification(id string, notification *Notification) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_notification := (*C.GNotification)(notification.ToC())

	C.g_application_send_notification((*C.GApplication)(recv.native), c_id, c_notification)

	return
}

// WithdrawNotification is a wrapper around the C function g_application_withdraw_notification.
func (recv *Application) WithdrawNotification(id string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	C.g_application_withdraw_notification((*C.GApplication)(recv.native), c_id)

	return
}

// GetOptionsDict is a wrapper around the C function g_application_command_line_get_options_dict.
func (recv *ApplicationCommandLine) GetOptionsDict() *glib.VariantDict {
	retC := C.g_application_command_line_get_options_dict((*C.GApplicationCommandLine)(recv.native))
	retGo := glib.VariantDictNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no param type

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// InetSocketAddressNewFromString is a wrapper around the C function g_inet_socket_address_new_from_string.
func InetSocketAddressNewFromString(address string, port uint32) *SocketAddress {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_port := (C.guint)(port)

	retC := C.g_inet_socket_address_new_from_string(c_address, c_port)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

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

// AddButton is a wrapper around the C function g_notification_add_button.
func (recv *Notification) AddButton(label string, detailedAction string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_notification_add_button((*C.GNotification)(recv.native), c_label, c_detailed_action)

	return
}

// Unsupported : g_notification_add_button_with_target : unsupported parameter ... : varargs

// Unsupported : g_notification_add_button_with_target_value : unsupported parameter target : Blacklisted record : GVariant

// SetBody is a wrapper around the C function g_notification_set_body.
func (recv *Notification) SetBody(body string) {
	c_body := C.CString(body)
	defer C.free(unsafe.Pointer(c_body))

	C.g_notification_set_body((*C.GNotification)(recv.native), c_body)

	return
}

// SetDefaultAction is a wrapper around the C function g_notification_set_default_action.
func (recv *Notification) SetDefaultAction(detailedAction string) {
	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_notification_set_default_action((*C.GNotification)(recv.native), c_detailed_action)

	return
}

// Unsupported : g_notification_set_default_action_and_target : unsupported parameter ... : varargs

// Unsupported : g_notification_set_default_action_and_target_value : unsupported parameter target : Blacklisted record : GVariant

// Unsupported : g_notification_set_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// SetPriority is a wrapper around the C function g_notification_set_priority.
func (recv *Notification) SetPriority(priority NotificationPriority) {
	c_priority := (C.GNotificationPriority)(priority)

	C.g_notification_set_priority((*C.GNotification)(recv.native), c_priority)

	return
}

// SetTitle is a wrapper around the C function g_notification_set_title.
func (recv *Notification) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.g_notification_set_title((*C.GNotification)(recv.native), c_title)

	return
}

// SetUrgent is a wrapper around the C function g_notification_set_urgent.
func (recv *Notification) SetUrgent(urgent bool) {
	c_urgent :=
		boolToGboolean(urgent)

	C.g_notification_set_urgent((*C.GNotification)(recv.native), c_urgent)

	return
}

// Unsupported : g_output_stream_printf : unsupported parameter bytes_written : no type generator for gsize, gsize*

// Unsupported : g_output_stream_vprintf : unsupported parameter bytes_written : no type generator for gsize, gsize*

// Unsupported : g_settings_get_default_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_get_user_value : return type : Blacklisted record : GVariant

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

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

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no param type

// Unsupported : g_subprocess_communicate : unsupported parameter stdout_buf : record with indirection level of 2

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

// ForceExit is a wrapper around the C function g_subprocess_force_exit.
func (recv *Subprocess) ForceExit() {
	C.g_subprocess_force_exit((*C.GSubprocess)(recv.native))

	return
}

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

// SendSignal is a wrapper around the C function g_subprocess_send_signal.
func (recv *Subprocess) SendSignal(signalNum int32) {
	c_signal_num := (C.gint)(signalNum)

	C.g_subprocess_send_signal((*C.GSubprocess)(recv.native), c_signal_num)

	return
}

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

// SetCwd is a wrapper around the C function g_subprocess_launcher_set_cwd.
func (recv *SubprocessLauncher) SetCwd(cwd string) {
	c_cwd := C.CString(cwd)
	defer C.free(unsafe.Pointer(c_cwd))

	C.g_subprocess_launcher_set_cwd((*C.GSubprocessLauncher)(recv.native), c_cwd)

	return
}

// Unsupported : g_subprocess_launcher_set_environ : unsupported parameter env : no param type

// SetFlags is a wrapper around the C function g_subprocess_launcher_set_flags.
func (recv *SubprocessLauncher) SetFlags(flags SubprocessFlags) {
	c_flags := (C.GSubprocessFlags)(flags)

	C.g_subprocess_launcher_set_flags((*C.GSubprocessLauncher)(recv.native), c_flags)

	return
}

// SetStderrFilePath is a wrapper around the C function g_subprocess_launcher_set_stderr_file_path.
func (recv *SubprocessLauncher) SetStderrFilePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_subprocess_launcher_set_stderr_file_path((*C.GSubprocessLauncher)(recv.native), c_path)

	return
}

// SetStdinFilePath is a wrapper around the C function g_subprocess_launcher_set_stdin_file_path.
func (recv *SubprocessLauncher) SetStdinFilePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_subprocess_launcher_set_stdin_file_path((*C.GSubprocessLauncher)(recv.native), c_path)

	return
}

// SetStdoutFilePath is a wrapper around the C function g_subprocess_launcher_set_stdout_file_path.
func (recv *SubprocessLauncher) SetStdoutFilePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_subprocess_launcher_set_stdout_file_path((*C.GSubprocessLauncher)(recv.native), c_path)

	return
}

// Setenv is a wrapper around the C function g_subprocess_launcher_setenv.
func (recv *SubprocessLauncher) Setenv(variable string, value string, overwrite bool) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_overwrite :=
		boolToGboolean(overwrite)

	C.g_subprocess_launcher_setenv((*C.GSubprocessLauncher)(recv.native), c_variable, c_value, c_overwrite)

	return
}

// Unsupported : g_subprocess_launcher_spawn : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_launcher_spawnv : unsupported parameter argv : no param type

// TakeFd is a wrapper around the C function g_subprocess_launcher_take_fd.
func (recv *SubprocessLauncher) TakeFd(sourceFd int32, targetFd int32) {
	c_source_fd := (C.gint)(sourceFd)

	c_target_fd := (C.gint)(targetFd)

	C.g_subprocess_launcher_take_fd((*C.GSubprocessLauncher)(recv.native), c_source_fd, c_target_fd)

	return
}

// TakeStderrFd is a wrapper around the C function g_subprocess_launcher_take_stderr_fd.
func (recv *SubprocessLauncher) TakeStderrFd(fd int32) {
	c_fd := (C.gint)(fd)

	C.g_subprocess_launcher_take_stderr_fd((*C.GSubprocessLauncher)(recv.native), c_fd)

	return
}

// TakeStdinFd is a wrapper around the C function g_subprocess_launcher_take_stdin_fd.
func (recv *SubprocessLauncher) TakeStdinFd(fd int32) {
	c_fd := (C.gint)(fd)

	C.g_subprocess_launcher_take_stdin_fd((*C.GSubprocessLauncher)(recv.native), c_fd)

	return
}

// TakeStdoutFd is a wrapper around the C function g_subprocess_launcher_take_stdout_fd.
func (recv *SubprocessLauncher) TakeStdoutFd(fd int32) {
	c_fd := (C.gint)(fd)

	C.g_subprocess_launcher_take_stdout_fd((*C.GSubprocessLauncher)(recv.native), c_fd)

	return
}

// Unsetenv is a wrapper around the C function g_subprocess_launcher_unsetenv.
func (recv *SubprocessLauncher) Unsetenv(variable string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	C.g_subprocess_launcher_unsetenv((*C.GSubprocessLauncher)(recv.native), c_variable)

	return
}

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// InvokeRequestCertificate is a wrapper around the C function g_tls_interaction_invoke_request_certificate.
func (recv *TlsInteraction) InvokeRequestCertificate(connection *TlsConnection, flags TlsCertificateRequestFlags, cancellable *Cancellable) (TlsInteractionResult, error) {
	c_connection := (*C.GTlsConnection)(connection.ToC())

	c_flags := (C.GTlsCertificateRequestFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_invoke_request_certificate((*C.GTlsInteraction)(recv.native), c_connection, c_flags, c_cancellable, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// RequestCertificate is a wrapper around the C function g_tls_interaction_request_certificate.
func (recv *TlsInteraction) RequestCertificate(connection *TlsConnection, flags TlsCertificateRequestFlags, cancellable *Cancellable) (TlsInteractionResult, error) {
	c_connection := (*C.GTlsConnection)(connection.ToC())

	c_flags := (C.GTlsCertificateRequestFlags)(flags)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_request_certificate((*C.GTlsInteraction)(recv.native), c_connection, c_flags, c_cancellable, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_tls_interaction_request_certificate_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_tls_interaction_request_certificate_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type
