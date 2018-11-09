// This is a generated file - DO NOT EDIT
// +build gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported signal 'launch-failed' for AppLaunchContext : unsupported parameter startup_notify_id : type utf8 :

// Unsupported signal 'launched' for AppLaunchContext : unsupported parameter platform_data : type GLib.Variant : Blacklisted record : GVariant

// CreateFileForArg is a wrapper around the C function g_application_command_line_create_file_for_arg.
func (recv *ApplicationCommandLine) CreateFileForArg(arg string) *File {
	c_arg := C.CString(arg)
	defer C.free(unsafe.Pointer(c_arg))

	retC := C.g_application_command_line_create_file_for_arg((*C.GApplicationCommandLine)(recv.native), c_arg)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_credentials_get_unix_pid : no return generator

// GetBoolean is a wrapper around the C function g_desktop_app_info_get_boolean.
func (recv *DesktopAppInfo) GetBoolean(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_get_boolean((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// GetString is a wrapper around the C function g_desktop_app_info_get_string.
func (recv *DesktopAppInfo) GetString(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_get_string((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// HasKey is a wrapper around the C function g_desktop_app_info_has_key.
func (recv *DesktopAppInfo) HasKey(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_has_key((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// GetChild is a wrapper around the C function g_file_enumerator_get_child.
func (recv *FileEnumerator) GetChild(info *FileInfo) *File {
	c_info := (*C.GFileInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GFileInfo)(info.ToC())
	}

	retC := C.g_file_enumerator_get_child((*C.GFileEnumerator)(recv.native), c_info)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDeletionDate is a wrapper around the C function g_file_info_get_deletion_date.
func (recv *FileInfo) GetDeletionDate() *glib.DateTime {
	retC := C.g_file_info_get_deletion_date((*C.GFileInfo)(recv.native))
	retGo := glib.DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MemoryOutputStreamNewResizable is a wrapper around the C function g_memory_output_stream_new_resizable.
func MemoryOutputStreamNewResizable() *MemoryOutputStream {
	retC := C.g_memory_output_stream_new_resizable()
	retGo := MemoryOutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetDefaultProxy is a wrapper around the C function g_simple_proxy_resolver_set_default_proxy.
func (recv *SimpleProxyResolver) SetDefaultProxy(defaultProxy string) {
	c_default_proxy := C.CString(defaultProxy)
	defer C.free(unsafe.Pointer(c_default_proxy))

	C.g_simple_proxy_resolver_set_default_proxy((*C.GSimpleProxyResolver)(recv.native), c_default_proxy)

	return
}

// Unsupported : g_simple_proxy_resolver_set_ignore_hosts : unsupported parameter ignore_hosts : in string with indirection level of 2

// SetUriProxy is a wrapper around the C function g_simple_proxy_resolver_set_uri_proxy.
func (recv *SimpleProxyResolver) SetUriProxy(uriScheme string, proxy string) {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	c_proxy := C.CString(proxy)
	defer C.free(unsafe.Pointer(c_proxy))

	C.g_simple_proxy_resolver_set_uri_proxy((*C.GSimpleProxyResolver)(recv.native), c_uri_scheme, c_proxy)

	return
}

// GetOption is a wrapper around the C function g_socket_get_option.
func (recv *Socket) GetOption(level int32, optname int32) (bool, int32, error) {
	c_level := (C.gint)(level)

	c_optname := (C.gint)(optname)

	var c_value C.gint

	var cThrowableError *C.GError

	retC := C.g_socket_get_option((*C.GSocket)(recv.native), c_level, c_optname, &c_value, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	value := (int32)(c_value)

	return retGo, value, goThrowableError
}

// SetOption is a wrapper around the C function g_socket_set_option.
func (recv *Socket) SetOption(level int32, optname int32, value int32) (bool, error) {
	c_level := (C.gint)(level)

	c_optname := (C.gint)(optname)

	c_value := (C.gint)(value)

	var cThrowableError *C.GError

	retC := C.g_socket_set_option((*C.GSocket)(recv.native), c_level, c_optname, c_value, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetProxyResolver is a wrapper around the C function g_socket_client_get_proxy_resolver.
func (recv *SocketClient) GetProxyResolver() *ProxyResolver {
	retC := C.g_socket_client_get_proxy_resolver((*C.GSocketClient)(recv.native))
	retGo := ProxyResolverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetProxyResolver is a wrapper around the C function g_socket_client_set_proxy_resolver.
func (recv *SocketClient) SetProxyResolver(proxyResolver *ProxyResolver) {
	c_proxy_resolver := (*C.GProxyResolver)(proxyResolver.ToC())

	C.g_socket_client_set_proxy_resolver((*C.GSocketClient)(recv.native), c_proxy_resolver)

	return
}

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_task_attach_source : unsupported parameter callback : no type generator for GLib.SourceFunc (GSourceFunc) for param callback

// GetCancellable is a wrapper around the C function g_task_get_cancellable.
func (recv *Task) GetCancellable() *Cancellable {
	retC := C.g_task_get_cancellable((*C.GTask)(recv.native))
	retGo := CancellableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCheckCancellable is a wrapper around the C function g_task_get_check_cancellable.
func (recv *Task) GetCheckCancellable() bool {
	retC := C.g_task_get_check_cancellable((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetContext is a wrapper around the C function g_task_get_context.
func (recv *Task) GetContext() *glib.MainContext {
	retC := C.g_task_get_context((*C.GTask)(recv.native))
	retGo := glib.MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPriority is a wrapper around the C function g_task_get_priority.
func (recv *Task) GetPriority() int32 {
	retC := C.g_task_get_priority((*C.GTask)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetReturnOnCancel is a wrapper around the C function g_task_get_return_on_cancel.
func (recv *Task) GetReturnOnCancel() bool {
	retC := C.g_task_get_return_on_cancel((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSourceObject is a wrapper around the C function g_task_get_source_object.
func (recv *Task) GetSourceObject() uintptr {
	retC := C.g_task_get_source_object((*C.GTask)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// GetSourceTag is a wrapper around the C function g_task_get_source_tag.
func (recv *Task) GetSourceTag() uintptr {
	retC := C.g_task_get_source_tag((*C.GTask)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetTaskData is a wrapper around the C function g_task_get_task_data.
func (recv *Task) GetTaskData() uintptr {
	retC := C.g_task_get_task_data((*C.GTask)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// HadError is a wrapper around the C function g_task_had_error.
func (recv *Task) HadError() bool {
	retC := C.g_task_had_error((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PropagateBoolean is a wrapper around the C function g_task_propagate_boolean.
func (recv *Task) PropagateBoolean() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_task_propagate_boolean((*C.GTask)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PropagateInt is a wrapper around the C function g_task_propagate_int.
func (recv *Task) PropagateInt() (int64, error) {
	var cThrowableError *C.GError

	retC := C.g_task_propagate_int((*C.GTask)(recv.native), &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PropagatePointer is a wrapper around the C function g_task_propagate_pointer.
func (recv *Task) PropagatePointer() (uintptr, error) {
	var cThrowableError *C.GError

	retC := C.g_task_propagate_pointer((*C.GTask)(recv.native), &cThrowableError)
	retGo := (uintptr)(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ReturnBoolean is a wrapper around the C function g_task_return_boolean.
func (recv *Task) ReturnBoolean(result bool) {
	c_result :=
		boolToGboolean(result)

	C.g_task_return_boolean((*C.GTask)(recv.native), c_result)

	return
}

// ReturnError is a wrapper around the C function g_task_return_error.
func (recv *Task) ReturnError(error *glib.Error) {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	C.g_task_return_error((*C.GTask)(recv.native), c_error)

	return
}

// ReturnErrorIfCancelled is a wrapper around the C function g_task_return_error_if_cancelled.
func (recv *Task) ReturnErrorIfCancelled() bool {
	retC := C.g_task_return_error_if_cancelled((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ReturnInt is a wrapper around the C function g_task_return_int.
func (recv *Task) ReturnInt(result int64) {
	c_result := (C.gssize)(result)

	C.g_task_return_int((*C.GTask)(recv.native), c_result)

	return
}

// Unsupported : g_task_return_new_error : unsupported parameter ... : varargs

// Unsupported : g_task_return_pointer : unsupported parameter result_destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param result_destroy

// Unsupported : g_task_run_in_thread : unsupported parameter task_func : no type generator for TaskThreadFunc (GTaskThreadFunc) for param task_func

// Unsupported : g_task_run_in_thread_sync : unsupported parameter task_func : no type generator for TaskThreadFunc (GTaskThreadFunc) for param task_func

// SetCheckCancellable is a wrapper around the C function g_task_set_check_cancellable.
func (recv *Task) SetCheckCancellable(checkCancellable bool) {
	c_check_cancellable :=
		boolToGboolean(checkCancellable)

	C.g_task_set_check_cancellable((*C.GTask)(recv.native), c_check_cancellable)

	return
}

// SetPriority is a wrapper around the C function g_task_set_priority.
func (recv *Task) SetPriority(priority int32) {
	c_priority := (C.gint)(priority)

	C.g_task_set_priority((*C.GTask)(recv.native), c_priority)

	return
}

// SetReturnOnCancel is a wrapper around the C function g_task_set_return_on_cancel.
func (recv *Task) SetReturnOnCancel(returnOnCancel bool) bool {
	c_return_on_cancel :=
		boolToGboolean(returnOnCancel)

	retC := C.g_task_set_return_on_cancel((*C.GTask)(recv.native), c_return_on_cancel)
	retGo := retC == C.TRUE

	return retGo
}

// SetSourceTag is a wrapper around the C function g_task_set_source_tag.
func (recv *Task) SetSourceTag(sourceTag uintptr) {
	c_source_tag := (C.gpointer)(sourceTag)

	C.g_task_set_source_tag((*C.GTask)(recv.native), c_source_tag)

	return
}

// Unsupported : g_task_set_task_data : unsupported parameter task_data_destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param task_data_destroy
