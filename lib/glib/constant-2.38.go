// This is a generated file - DO NOT EDIT
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

/*
A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string list
giving the available application actions.
*/
const KEY_FILE_DESKTOP_KEY_ACTIONS string = C.G_KEY_FILE_DESKTOP_KEY_ACTIONS

/*
A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean set to true
if the application is D-Bus activatable.
*/
const KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE string = C.G_KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE
