// This is a generated file - DO NOT EDIT
// +build gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Checks if @action_name is valid.
//
// @action_name is valid if it consists only of alphanumeric characters,
// plus '-' and '.'.  The empty string is not a valid action name.
//
// It is an error to call this function with a non-utf8 @action_name.
// @action_name must not be %NULL.
/*

C function : g_action_name_is_valid
*/
func ActionNameIsValid(actionName string) bool {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_name_is_valid(c_action_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_action_parse_detailed_name : unsupported parameter target_value : Blacklisted record : GVariant

// Unsupported : g_action_print_detailed_name : unsupported parameter target_value : Blacklisted record : GVariant

// Unsupported : g_icon_deserialize : unsupported parameter value : Blacklisted record : GVariant
