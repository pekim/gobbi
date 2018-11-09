// This is a generated file - DO NOT EDIT
// +build gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// MarkBusy is a wrapper around the C function g_application_mark_busy.
func (recv *Application) MarkBusy() {
	C.g_application_mark_busy((*C.GApplication)(recv.native))

	return
}

// UnmarkBusy is a wrapper around the C function g_application_unmark_busy.
func (recv *Application) UnmarkBusy() {
	C.g_application_unmark_busy((*C.GApplication)(recv.native))

	return
}

// BytesIconNew is a wrapper around the C function g_bytes_icon_new.
func BytesIconNew(bytes *glib.Bytes) *BytesIcon {
	c_bytes := (*C.GBytes)(bytes.ToC())

	retC := C.g_bytes_icon_new(c_bytes)
	retGo := BytesIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBytes is a wrapper around the C function g_bytes_icon_get_bytes.
func (recv *BytesIcon) GetBytes() *glib.Bytes {
	retC := C.g_bytes_icon_get_bytes((*C.GBytesIcon)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPropertyInfo is a wrapper around the C function g_dbus_method_invocation_get_property_info.
func (recv *DBusMethodInvocation) GetPropertyInfo() *DBusPropertyInfo {
	retC := C.g_dbus_method_invocation_get_property_info((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusPropertyInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActionName is a wrapper around the C function g_desktop_app_info_get_action_name.
func (recv *DesktopAppInfo) GetActionName(actionName string) string {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_desktop_app_info_get_action_name((*C.GDesktopAppInfo)(recv.native), c_action_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// LaunchAction is a wrapper around the C function g_desktop_app_info_launch_action.
func (recv *DesktopAppInfo) LaunchAction(actionName string, launchContext *AppLaunchContext) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_launch_context := (*C.GAppLaunchContext)(launchContext.ToC())

	C.g_desktop_app_info_launch_action((*C.GDesktopAppInfo)(recv.native), c_action_name, c_launch_context)

	return
}

// Unsupported : g_desktop_app_info_list_actions : no return type

// RemoveAll is a wrapper around the C function g_menu_remove_all.
func (recv *Menu) RemoveAll() {
	C.g_menu_remove_all((*C.GMenu)(recv.native))

	return
}

// Unsupported : g_menu_item_set_icon : unsupported parameter icon : no type generator for Icon (GIcon*) for param icon

// PropertyAction is a wrapper around the C record GPropertyAction.
type PropertyAction struct {
	native *C.GPropertyAction
}

func PropertyActionNewFromC(u unsafe.Pointer) *PropertyAction {
	c := (*C.GPropertyAction)(u)
	if c == nil {
		return nil
	}

	g := &PropertyAction{native: c}

	return g
}

func (recv *PropertyAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *PropertyAction) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to PropertyAction.
// Exercise care, as this is a potentially dangerous function if the Object is not a PropertyAction.
func CastToPropertyAction(object *gobject.Object) *PropertyAction {
	return PropertyActionNewFromC(object.ToC())
}

// PropertyActionNew is a wrapper around the C function g_property_action_new.
func PropertyActionNew(name string, object uintptr, propertyName string) *PropertyAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object := (C.gpointer)(object)

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.g_property_action_new(c_name, c_object, c_property_name)
	retGo := PropertyActionNewFromC(unsafe.Pointer(retC))

	return retGo
}
