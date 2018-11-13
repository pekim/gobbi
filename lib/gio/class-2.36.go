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

// Creates a #GFile corresponding to a filename that was given as part
// of the invocation of @cmdline.
//
// This differs from g_file_new_for_commandline_arg() in that it
// resolves relative pathnames using the current working directory of
// the invoking process rather than the local process.
/*

C function : g_application_command_line_create_file_for_arg
*/
func (recv *ApplicationCommandLine) CreateFileForArg(arg string) *File {
	c_arg := C.CString(arg)
	defer C.free(unsafe.Pointer(c_arg))

	retC := C.g_application_command_line_create_file_for_arg((*C.GApplicationCommandLine)(recv.native), c_arg)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_credentials_get_unix_pid : no return generator

// Looks up a boolean value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
/*

C function : g_desktop_app_info_get_boolean
*/
func (recv *DesktopAppInfo) GetBoolean(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_get_boolean((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Looks up a string value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
/*

C function : g_desktop_app_info_get_string
*/
func (recv *DesktopAppInfo) GetString(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_get_string((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns whether @key exists in the "Desktop Entry" group
// of the keyfile backing @info.
/*

C function : g_desktop_app_info_has_key
*/
func (recv *DesktopAppInfo) HasKey(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_has_key((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Return a new #GFile which refers to the file named by @info in the source
// directory of @enumerator.  This function is primarily intended to be used
// inside loops with g_file_enumerator_next_file().
//
// This is a convenience method that's equivalent to:
// |[<!-- language="C" -->
// gchar *name = g_file_info_get_name (info);
// GFile *child = g_file_get_child (g_file_enumerator_get_container (enumr),
// name);
// ]|
/*

C function : g_file_enumerator_get_child
*/
func (recv *FileEnumerator) GetChild(info *FileInfo) *File {
	c_info := (*C.GFileInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GFileInfo)(info.ToC())
	}

	retC := C.g_file_enumerator_get_child((*C.GFileEnumerator)(recv.native), c_info)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the #GDateTime representing the deletion date of the file, as
// available in G_FILE_ATTRIBUTE_TRASH_DELETION_DATE. If the
// G_FILE_ATTRIBUTE_TRASH_DELETION_DATE attribute is unset, %NULL is returned.
/*

C function : g_file_info_get_deletion_date
*/
func (recv *FileInfo) GetDeletionDate() *glib.DateTime {
	retC := C.g_file_info_get_deletion_date((*C.GFileInfo)(recv.native))
	retGo := glib.DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GMemoryOutputStream, using g_realloc() and g_free()
// for memory allocation.
/*

C function : g_memory_output_stream_new_resizable
*/
func MemoryOutputStreamNewResizable() *MemoryOutputStream {
	retC := C.g_memory_output_stream_new_resizable()
	retGo := MemoryOutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the default proxy on @resolver, to be used for any URIs that
// don't match #GSimpleProxyResolver:ignore-hosts or a proxy set
// via g_simple_proxy_resolver_set_uri_proxy().
//
// If @default_proxy starts with "socks://",
// #GSimpleProxyResolver will treat it as referring to all three of
// the socks5, socks4a, and socks4 proxy types.
/*

C function : g_simple_proxy_resolver_set_default_proxy
*/
func (recv *SimpleProxyResolver) SetDefaultProxy(defaultProxy string) {
	c_default_proxy := C.CString(defaultProxy)
	defer C.free(unsafe.Pointer(c_default_proxy))

	C.g_simple_proxy_resolver_set_default_proxy((*C.GSimpleProxyResolver)(recv.native), c_default_proxy)

	return
}

// Unsupported : g_simple_proxy_resolver_set_ignore_hosts : unsupported parameter ignore_hosts : in string with indirection level of 2

// Adds a URI-scheme-specific proxy to @resolver; URIs whose scheme
// matches @uri_scheme (and which don't match
// #GSimpleProxyResolver:ignore-hosts) will be proxied via @proxy.
//
// As with #GSimpleProxyResolver:default-proxy, if @proxy starts with
// "socks://", #GSimpleProxyResolver will treat it
// as referring to all three of the socks5, socks4a, and socks4 proxy
// types.
/*

C function : g_simple_proxy_resolver_set_uri_proxy
*/
func (recv *SimpleProxyResolver) SetUriProxy(uriScheme string, proxy string) {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	c_proxy := C.CString(proxy)
	defer C.free(unsafe.Pointer(c_proxy))

	C.g_simple_proxy_resolver_set_uri_proxy((*C.GSimpleProxyResolver)(recv.native), c_uri_scheme, c_proxy)

	return
}

// Gets the value of an integer-valued option on @socket, as with
// getsockopt(). (If you need to fetch a  non-integer-valued option,
// you will need to call getsockopt() directly.)
//
// The [<gio/gnetworking.h>][gio-gnetworking.h]
// header pulls in system headers that will define most of the
// standard/portable socket options. For unusual socket protocols or
// platform-dependent options, you may need to include additional
// headers.
//
// Note that even for socket options that are a single byte in size,
// @value is still a pointer to a #gint variable, not a #guchar;
// g_socket_get_option() will handle the conversion internally.
/*

C function : g_socket_get_option
*/
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

// Sets the value of an integer-valued option on @socket, as with
// setsockopt(). (If you need to set a non-integer-valued option,
// you will need to call setsockopt() directly.)
//
// The [<gio/gnetworking.h>][gio-gnetworking.h]
// header pulls in system headers that will define most of the
// standard/portable socket options. For unusual socket protocols or
// platform-dependent options, you may need to include additional
// headers.
/*

C function : g_socket_set_option
*/
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

// Gets the #GProxyResolver being used by @client. Normally, this will
// be the resolver returned by g_proxy_resolver_get_default(), but you
// can override it with g_socket_client_set_proxy_resolver().
/*

C function : g_socket_client_get_proxy_resolver
*/
func (recv *SocketClient) GetProxyResolver() *ProxyResolver {
	retC := C.g_socket_client_get_proxy_resolver((*C.GSocketClient)(recv.native))
	retGo := ProxyResolverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Overrides the #GProxyResolver used by @client. You can call this if
// you want to use specific proxies, rather than using the system
// default proxy settings.
//
// Note that whether or not the proxy resolver is actually used
// depends on the setting of #GSocketClient:enable-proxy, which is not
// changed by this function (but which is %TRUE by default)
/*

C function : g_socket_client_set_proxy_resolver
*/
func (recv *SocketClient) SetProxyResolver(proxyResolver *ProxyResolver) {
	c_proxy_resolver := (*C.GProxyResolver)(proxyResolver.ToC())

	C.g_socket_client_set_proxy_resolver((*C.GSocketClient)(recv.native), c_proxy_resolver)

	return
}

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_task_attach_source : unsupported parameter callback : no type generator for GLib.SourceFunc (GSourceFunc) for param callback

// Gets @task's #GCancellable
/*

C function : g_task_get_cancellable
*/
func (recv *Task) GetCancellable() *Cancellable {
	retC := C.g_task_get_cancellable((*C.GTask)(recv.native))
	retGo := CancellableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @task's check-cancellable flag. See
// g_task_set_check_cancellable() for more details.
/*

C function : g_task_get_check_cancellable
*/
func (recv *Task) GetCheckCancellable() bool {
	retC := C.g_task_get_check_cancellable((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the #GMainContext that @task will return its result in (that
// is, the context that was the
// [thread-default main context][g-main-context-push-thread-default]
// at the point when @task was created).
//
// This will always return a non-%NULL value, even if the task's
// context is the default #GMainContext.
/*

C function : g_task_get_context
*/
func (recv *Task) GetContext() *glib.MainContext {
	retC := C.g_task_get_context((*C.GTask)(recv.native))
	retGo := glib.MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @task's priority
/*

C function : g_task_get_priority
*/
func (recv *Task) GetPriority() int32 {
	retC := C.g_task_get_priority((*C.GTask)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets @task's return-on-cancel flag. See
// g_task_set_return_on_cancel() for more details.
/*

C function : g_task_get_return_on_cancel
*/
func (recv *Task) GetReturnOnCancel() bool {
	retC := C.g_task_get_return_on_cancel((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the source object from @task. Like
// g_async_result_get_source_object(), but does not ref the object.
/*

C function : g_task_get_source_object
*/
func (recv *Task) GetSourceObject() uintptr {
	retC := C.g_task_get_source_object((*C.GTask)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Gets @task's source tag. See g_task_set_source_tag().
/*

C function : g_task_get_source_tag
*/
func (recv *Task) GetSourceTag() uintptr {
	retC := C.g_task_get_source_tag((*C.GTask)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Gets @task's `task_data`.
/*

C function : g_task_get_task_data
*/
func (recv *Task) GetTaskData() uintptr {
	retC := C.g_task_get_task_data((*C.GTask)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Tests if @task resulted in an error.
/*

C function : g_task_had_error
*/
func (recv *Task) HadError() bool {
	retC := C.g_task_had_error((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the result of @task as a #gboolean.
//
// If the task resulted in an error, or was cancelled, then this will
// instead return %FALSE and set @error.
//
// Since this method transfers ownership of the return value (or
// error) to the caller, you may only call it once.
/*

C function : g_task_propagate_boolean
*/
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

// Gets the result of @task as an integer (#gssize).
//
// If the task resulted in an error, or was cancelled, then this will
// instead return -1 and set @error.
//
// Since this method transfers ownership of the return value (or
// error) to the caller, you may only call it once.
/*

C function : g_task_propagate_int
*/
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

// Gets the result of @task as a pointer, and transfers ownership
// of that value to the caller.
//
// If the task resulted in an error, or was cancelled, then this will
// instead return %NULL and set @error.
//
// Since this method transfers ownership of the return value (or
// error) to the caller, you may only call it once.
/*

C function : g_task_propagate_pointer
*/
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

// Sets @task's result to @result and completes the task (see
// g_task_return_pointer() for more discussion of exactly what this
// means).
/*

C function : g_task_return_boolean
*/
func (recv *Task) ReturnBoolean(result bool) {
	c_result :=
		boolToGboolean(result)

	C.g_task_return_boolean((*C.GTask)(recv.native), c_result)

	return
}

// Sets @task's result to @error (which @task assumes ownership of)
// and completes the task (see g_task_return_pointer() for more
// discussion of exactly what this means).
//
// Note that since the task takes ownership of @error, and since the
// task may be completed before returning from g_task_return_error(),
// you cannot assume that @error is still valid after calling this.
// Call g_error_copy() on the error if you need to keep a local copy
// as well.
//
// See also g_task_return_new_error().
/*

C function : g_task_return_error
*/
func (recv *Task) ReturnError(error *glib.Error) {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	C.g_task_return_error((*C.GTask)(recv.native), c_error)

	return
}

// Checks if @task's #GCancellable has been cancelled, and if so, sets
// @task's error accordingly and completes the task (see
// g_task_return_pointer() for more discussion of exactly what this
// means).
/*

C function : g_task_return_error_if_cancelled
*/
func (recv *Task) ReturnErrorIfCancelled() bool {
	retC := C.g_task_return_error_if_cancelled((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets @task's result to @result and completes the task (see
// g_task_return_pointer() for more discussion of exactly what this
// means).
/*

C function : g_task_return_int
*/
func (recv *Task) ReturnInt(result int64) {
	c_result := (C.gssize)(result)

	C.g_task_return_int((*C.GTask)(recv.native), c_result)

	return
}

// Unsupported : g_task_return_new_error : unsupported parameter ... : varargs

// Unsupported : g_task_return_pointer : unsupported parameter result_destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param result_destroy

// Unsupported : g_task_run_in_thread : unsupported parameter task_func : no type generator for TaskThreadFunc (GTaskThreadFunc) for param task_func

// Unsupported : g_task_run_in_thread_sync : unsupported parameter task_func : no type generator for TaskThreadFunc (GTaskThreadFunc) for param task_func

// Sets or clears @task's check-cancellable flag. If this is %TRUE
// (the default), then g_task_propagate_pointer(), etc, and
// g_task_had_error() will check the task's #GCancellable first, and
// if it has been cancelled, then they will consider the task to have
// returned an "Operation was cancelled" error
// (%G_IO_ERROR_CANCELLED), regardless of any other error or return
// value the task may have had.
//
// If @check_cancellable is %FALSE, then the #GTask will not check the
// cancellable itself, and it is up to @task's owner to do this (eg,
// via g_task_return_error_if_cancelled()).
//
// If you are using g_task_set_return_on_cancel() as well, then
// you must leave check-cancellable set %TRUE.
/*

C function : g_task_set_check_cancellable
*/
func (recv *Task) SetCheckCancellable(checkCancellable bool) {
	c_check_cancellable :=
		boolToGboolean(checkCancellable)

	C.g_task_set_check_cancellable((*C.GTask)(recv.native), c_check_cancellable)

	return
}

// Sets @task's priority. If you do not call this, it will default to
// %G_PRIORITY_DEFAULT.
//
// This will affect the priority of #GSources created with
// g_task_attach_source() and the scheduling of tasks run in threads,
// and can also be explicitly retrieved later via
// g_task_get_priority().
/*

C function : g_task_set_priority
*/
func (recv *Task) SetPriority(priority int32) {
	c_priority := (C.gint)(priority)

	C.g_task_set_priority((*C.GTask)(recv.native), c_priority)

	return
}

// Sets or clears @task's return-on-cancel flag. This is only
// meaningful for tasks run via g_task_run_in_thread() or
// g_task_run_in_thread_sync().
//
// If @return_on_cancel is %TRUE, then cancelling @task's
// #GCancellable will immediately cause it to return, as though the
// task's #GTaskThreadFunc had called
// g_task_return_error_if_cancelled() and then returned.
//
// This allows you to create a cancellable wrapper around an
// uninterruptable function. The #GTaskThreadFunc just needs to be
// careful that it does not modify any externally-visible state after
// it has been cancelled. To do that, the thread should call
// g_task_set_return_on_cancel() again to (atomically) set
// return-on-cancel %FALSE before making externally-visible changes;
// if the task gets cancelled before the return-on-cancel flag could
// be changed, g_task_set_return_on_cancel() will indicate this by
// returning %FALSE.
//
// You can disable and re-enable this flag multiple times if you wish.
// If the task's #GCancellable is cancelled while return-on-cancel is
// %FALSE, then calling g_task_set_return_on_cancel() to set it %TRUE
// again will cause the task to be cancelled at that point.
//
// If the task's #GCancellable is already cancelled before you call
// g_task_run_in_thread()/g_task_run_in_thread_sync(), then the
// #GTaskThreadFunc will still be run (for consistency), but the task
// will also be completed right away.
/*

C function : g_task_set_return_on_cancel
*/
func (recv *Task) SetReturnOnCancel(returnOnCancel bool) bool {
	c_return_on_cancel :=
		boolToGboolean(returnOnCancel)

	retC := C.g_task_set_return_on_cancel((*C.GTask)(recv.native), c_return_on_cancel)
	retGo := retC == C.TRUE

	return retGo
}

// Sets @task's source tag. You can use this to tag a task return
// value with a particular pointer (usually a pointer to the function
// doing the tagging) and then later check it using
// g_task_get_source_tag() (or g_async_result_is_tagged()) in the
// task's "finish" function, to figure out if the response came from a
// particular place.
/*

C function : g_task_set_source_tag
*/
func (recv *Task) SetSourceTag(sourceTag uintptr) {
	c_source_tag := (C.gpointer)(sourceTag)

	C.g_task_set_source_tag((*C.GTask)(recv.native), c_source_tag)

	return
}

// Unsupported : g_task_set_task_data : unsupported parameter task_data_destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param task_data_destroy
