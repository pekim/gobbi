// This is a generated file - DO NOT EDIT
// +build gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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

// Escape @string so it can appear in a D-Bus address as the value
// part of a key-value pair.
//
// For instance, if @string is `/run/bus-for-:0`,
// this function would return `/run/bus-for-%3A0`,
// which could be used in a D-Bus address like
// `unix:nonce-tcp:host=127.0.0.1,port=42,noncefile=/run/bus-for-%3A0`.
/*

C function : g_dbus_address_escape_value
*/
func DbusAddressEscapeValue(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_address_escape_value(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GFile with the given argument from the command line.
//
// This function is similar to g_file_new_for_commandline_arg() except
// that it allows for passing the current working directory as an
// argument instead of using the current working directory of the
// process.
//
// This is useful if the commandline argument was given in a context
// other than the invocation of the current process.
//
// See also g_application_command_line_create_file_for_arg().
/*

C function : g_file_new_for_commandline_arg_and_cwd
*/
func FileNewForCommandlineArgAndCwd(arg string, cwd string) *File {
	c_arg := C.CString(arg)
	defer C.free(unsafe.Pointer(c_arg))

	c_cwd := C.CString(cwd)
	defer C.free(unsafe.Pointer(c_cwd))

	retC := C.g_file_new_for_commandline_arg_and_cwd(c_arg, c_cwd)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Initializes the platform networking libraries (eg, on Windows, this
// calls WSAStartup()). GLib will call this itself if it is needed, so
// you only need to call it if you directly call system networking
// functions (without calling any GLib networking functions first).
/*

C function : g_networking_init
*/
func NetworkingInit() {
	C.g_networking_init()

	return
}
