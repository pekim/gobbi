// This is a generated file - DO NOT EDIT
// +build gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_application_add_main_option : unsupported parameter short_name : no type generator for gchar (char) for param short_name

// Gets the resource base path of @application.
//
// See g_application_set_resource_base_path() for more information.
/*

C function

g_application_get_resource_base_path
*/
func (recv *Application) GetResourceBasePath() string {
	retC := C.g_application_get_resource_base_path((*C.GApplication)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets (or unsets) the base resource path of @application.
//
// The path is used to automatically load various [application
// resources][gresource] such as menu layouts and action descriptions.
// The various types of resources will be found at fixed names relative
// to the given base path.
//
// By default, the resource base path is determined from the application
// ID by prefixing '/' and replacing each '.' with '/'.  This is done at
// the time that the #GApplication object is constructed.  Changes to
// the application ID after that point will not have an impact on the
// resource base path.
//
// As an example, if the application has an ID of "org.example.app" then
// the default resource base path will be "/org/example/app".  If this
// is a #GtkApplication (and you have not manually changed the path)
// then Gtk will then search for the menus of the application at
// "/org/example/app/gtk/menus.ui".
//
// See #GResource for more information about adding resources to your
// application.
//
// You can disable automatic resource loading functionality by setting
// the path to %NULL.
//
// Changing the resource base path once the application is running is
// not recommended.  The point at which the resource path is consulted
// for forming paths for various purposes is unspecified.  When writing
// a sub-class of #GApplication you should either set the
// #GApplication:resource-base-path property at construction time, or call
// this function during the instance initialization. Alternatively, you
// can call this function in the #GApplicationClass.startup virtual function,
// before chaining up to the parent implementation.
/*

C function

g_application_set_resource_base_path
*/
func (recv *Application) SetResourceBasePath(resourcePath string) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	C.g_application_set_resource_base_path((*C.GApplication)(recv.native), c_resource_path)

	return
}
