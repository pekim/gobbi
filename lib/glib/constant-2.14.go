// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// The name of the main group of a desktop entry file, as defined in the
// [Desktop Entry Specification](http://freedesktop.org/Standards/desktop-entry-spec).
// Consult the specification for more
// details about the meanings of the keys below.
const KEY_FILE_DESKTOP_GROUP string = C.G_KEY_FILE_DESKTOP_GROUP

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list
// of strings giving the categories in which the desktop entry
// should be shown in a menu.
const KEY_FILE_DESKTOP_KEY_CATEGORIES string = C.G_KEY_FILE_DESKTOP_KEY_CATEGORIES

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
// string giving the tooltip for the desktop entry.
const KEY_FILE_DESKTOP_KEY_COMMENT string = C.G_KEY_FILE_DESKTOP_KEY_COMMENT

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
// giving the command line to execute. It is only valid for desktop
// entries with the `Application` type.
const KEY_FILE_DESKTOP_KEY_EXEC string = C.G_KEY_FILE_DESKTOP_KEY_EXEC

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
// string giving the generic name of the desktop entry.
const KEY_FILE_DESKTOP_KEY_GENERIC_NAME string = C.G_KEY_FILE_DESKTOP_KEY_GENERIC_NAME

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
// stating whether the desktop entry has been deleted by the user.
const KEY_FILE_DESKTOP_KEY_HIDDEN string = C.G_KEY_FILE_DESKTOP_KEY_HIDDEN

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
// string giving the name of the icon to be displayed for the desktop
// entry.
const KEY_FILE_DESKTOP_KEY_ICON string = C.G_KEY_FILE_DESKTOP_KEY_ICON

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list
// of strings giving the MIME types supported by this desktop entry.
const KEY_FILE_DESKTOP_KEY_MIME_TYPE string = C.G_KEY_FILE_DESKTOP_KEY_MIME_TYPE

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
// string giving the specific name of the desktop entry.
const KEY_FILE_DESKTOP_KEY_NAME string = C.G_KEY_FILE_DESKTOP_KEY_NAME

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list of
// strings identifying the environments that should not display the
// desktop entry.
const KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN string = C.G_KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
// stating whether the desktop entry should be shown in menus.
const KEY_FILE_DESKTOP_KEY_NO_DISPLAY string = C.G_KEY_FILE_DESKTOP_KEY_NO_DISPLAY

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list of
// strings identifying the environments that should display the
// desktop entry.
const KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN string = C.G_KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
// containing the working directory to run the program in. It is only
// valid for desktop entries with the `Application` type.
const KEY_FILE_DESKTOP_KEY_PATH string = C.G_KEY_FILE_DESKTOP_KEY_PATH

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
// stating whether the application supports the
// [Startup Notification Protocol Specification](http://www.freedesktop.org/Standards/startup-notification-spec).
const KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY string = C.G_KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is string
// identifying the WM class or name hint of a window that the application
// will create, which can be used to emulate Startup Notification with
// older applications.
const KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS string = C.G_KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
// stating whether the program should be run in a terminal window.
// It is only valid for desktop entries with the
// `Application` type.
const KEY_FILE_DESKTOP_KEY_TERMINAL string = C.G_KEY_FILE_DESKTOP_KEY_TERMINAL

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
// giving the file name of a binary on disk used to determine if the
// program is actually installed. It is only valid for desktop entries
// with the `Application` type.
const KEY_FILE_DESKTOP_KEY_TRY_EXEC string = C.G_KEY_FILE_DESKTOP_KEY_TRY_EXEC

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
// giving the type of the desktop entry. Usually
// #G_KEY_FILE_DESKTOP_TYPE_APPLICATION,
// #G_KEY_FILE_DESKTOP_TYPE_LINK, or
// #G_KEY_FILE_DESKTOP_TYPE_DIRECTORY.
const KEY_FILE_DESKTOP_KEY_TYPE string = C.G_KEY_FILE_DESKTOP_KEY_TYPE

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
// giving the URL to access. It is only valid for desktop entries
// with the `Link` type.
const KEY_FILE_DESKTOP_KEY_URL string = C.G_KEY_FILE_DESKTOP_KEY_URL

// A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
// giving the version of the Desktop Entry Specification used for
// the desktop entry file.
const KEY_FILE_DESKTOP_KEY_VERSION string = C.G_KEY_FILE_DESKTOP_KEY_VERSION

// The value of the #G_KEY_FILE_DESKTOP_KEY_TYPE, key for desktop
// entries representing applications.
const KEY_FILE_DESKTOP_TYPE_APPLICATION string = C.G_KEY_FILE_DESKTOP_TYPE_APPLICATION

// The value of the #G_KEY_FILE_DESKTOP_KEY_TYPE, key for desktop
// entries representing directories.
const KEY_FILE_DESKTOP_TYPE_DIRECTORY string = C.G_KEY_FILE_DESKTOP_TYPE_DIRECTORY

// The value of the #G_KEY_FILE_DESKTOP_KEY_TYPE, key for desktop
// entries representing links to documents.
const KEY_FILE_DESKTOP_TYPE_LINK string = C.G_KEY_FILE_DESKTOP_TYPE_LINK
