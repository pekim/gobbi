// This is a generated file - DO NOT EDIT
// +build gio_2.56

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

// Adds a description to the @application option context.
//
// See g_option_context_set_description() for more information.
/*

C function : g_application_set_option_context_description
*/
func (recv *Application) SetOptionContextDescription(description string) {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_application_set_option_context_description((*C.GApplication)(recv.native), c_description)

	return
}

// Sets the parameter string to be used by the commandline handling of @application.
//
// This function registers the argument to be passed to g_option_context_new()
// when the internal #GOptionContext of @application is created.
//
// See g_option_context_new() for more information about @parameter_string.
/*

C function : g_application_set_option_context_parameter_string
*/
func (recv *Application) SetOptionContextParameterString(parameterString string) {
	c_parameter_string := C.CString(parameterString)
	defer C.free(unsafe.Pointer(c_parameter_string))

	C.g_application_set_option_context_parameter_string((*C.GApplication)(recv.native), c_parameter_string)

	return
}

// Adds a summary to the @application option context.
//
// See g_option_context_set_summary() for more information.
/*

C function : g_application_set_option_context_summary
*/
func (recv *Application) SetOptionContextSummary(summary string) {
	c_summary := C.CString(summary)
	defer C.free(unsafe.Pointer(c_summary))

	C.g_application_set_option_context_summary((*C.GApplication)(recv.native), c_summary)

	return
}

// Looks up a localized string value in the keyfile backing @info
// translated to the current locale.
//
// The @key is looked up in the "Desktop Entry" group.
/*

C function : g_desktop_app_info_get_locale_string
*/
func (recv *DesktopAppInfo) GetLocaleString(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_get_locale_string((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Registers @socket to receive multicast messages sent to @group.
// @socket must be a %G_SOCKET_TYPE_DATAGRAM socket, and must have
// been bound to an appropriate interface and port with
// g_socket_bind().
//
// If @iface is %NULL, the system will automatically pick an interface
// to bind to based on @group.
//
// If @source_specific is not %NULL, use source-specific multicast as
// defined in RFC 4604. Note that on older platforms this may fail
// with a %G_IO_ERROR_NOT_SUPPORTED error.
//
// Note that this function can be called multiple times for the same
// @group with different @source_specific in order to receive multicast
// packets from more than one source.
/*

C function : g_socket_join_multicast_group_ssm
*/
func (recv *Socket) JoinMulticastGroupSsm(group *InetAddress, sourceSpecific *InetAddress, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(C.NULL)
	if group != nil {
		c_group = (*C.GInetAddress)(group.ToC())
	}

	c_source_specific := (*C.GInetAddress)(C.NULL)
	if sourceSpecific != nil {
		c_source_specific = (*C.GInetAddress)(sourceSpecific.ToC())
	}

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_join_multicast_group_ssm((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Removes @socket from the multicast group defined by @group, @iface,
// and @source_specific (which must all have the same values they had
// when you joined the group).
//
// @socket remains bound to its address and port, and can still receive
// unicast messages after calling this.
/*

C function : g_socket_leave_multicast_group_ssm
*/
func (recv *Socket) LeaveMulticastGroupSsm(group *InetAddress, sourceSpecific *InetAddress, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(C.NULL)
	if group != nil {
		c_group = (*C.GInetAddress)(group.ToC())
	}

	c_source_specific := (*C.GInetAddress)(C.NULL)
	if sourceSpecific != nil {
		c_source_specific = (*C.GInetAddress)(sourceSpecific.ToC())
	}

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_leave_multicast_group_ssm((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
