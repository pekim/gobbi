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

// Increases the busy count of @application.
//
// Use this function to indicate that the application is busy, for instance
// while a long running operation is pending.
//
// The busy state will be exposed to other processes, so a session shell will
// use that information to indicate the state to the user (e.g. with a
// spinner).
//
// To cancel the busy indication, use g_application_unmark_busy().
/*

C function

g_application_mark_busy
*/
func (recv *Application) MarkBusy() {
	C.g_application_mark_busy((*C.GApplication)(recv.native))

	return
}

// Decreases the busy count of @application.
//
// When the busy count reaches zero, the new state will be propagated
// to other processes.
//
// This function must only be called to cancel the effect of a previous
// call to g_application_mark_busy().
/*

C function

g_application_unmark_busy
*/
func (recv *Application) UnmarkBusy() {
	C.g_application_unmark_busy((*C.GApplication)(recv.native))

	return
}

// Creates a new icon for a bytes.
/*

C function

g_bytes_icon_new
*/
func BytesIconNew(bytes *glib.Bytes) *BytesIcon {
	c_bytes := (*C.GBytes)(C.NULL)
	if bytes != nil {
		c_bytes = (*C.GBytes)(bytes.ToC())
	}

	retC := C.g_bytes_icon_new(c_bytes)
	retGo := BytesIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GBytes associated with the given @icon.
/*

C function

g_bytes_icon_get_bytes
*/
func (recv *BytesIcon) GetBytes() *glib.Bytes {
	retC := C.g_bytes_icon_get_bytes((*C.GBytesIcon)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets information about the property that this method call is for, if
// any.
//
// This will only be set in the case of an invocation in response to a
// property Get or Set call that has been directed to the method call
// handler for an object on account of its property_get() or
// property_set() vtable pointers being unset.
//
// See #GDBusInterfaceVTable for more information.
//
// If the call was GetAll, %NULL will be returned.
/*

C function

g_dbus_method_invocation_get_property_info
*/
func (recv *DBusMethodInvocation) GetPropertyInfo() *DBusPropertyInfo {
	retC := C.g_dbus_method_invocation_get_property_info((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusPropertyInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the user-visible display name of the "additional application
// action" specified by @action_name.
//
// This corresponds to the "Name" key within the keyfile group for the
// action.
/*

C function

g_desktop_app_info_get_action_name
*/
func (recv *DesktopAppInfo) GetActionName(actionName string) string {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_desktop_app_info_get_action_name((*C.GDesktopAppInfo)(recv.native), c_action_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Activates the named application action.
//
// You may only call this function on action names that were
// returned from g_desktop_app_info_list_actions().
//
// Note that if the main entry of the desktop file indicates that the
// application supports startup notification, and @launch_context is
// non-%NULL, then startup notification will be used when activating the
// action (and as such, invocation of the action on the receiving side
// must signal the end of startup notification when it is completed).
// This is the expected behaviour of applications declaring additional
// actions, as per the desktop file specification.
//
// As with g_app_info_launch() there is no way to detect failures that
// occur while using this function.
/*

C function

g_desktop_app_info_launch_action
*/
func (recv *DesktopAppInfo) LaunchAction(actionName string, launchContext *AppLaunchContext) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_launch_context := (*C.GAppLaunchContext)(C.NULL)
	if launchContext != nil {
		c_launch_context = (*C.GAppLaunchContext)(launchContext.ToC())
	}

	C.g_desktop_app_info_launch_action((*C.GDesktopAppInfo)(recv.native), c_action_name, c_launch_context)

	return
}

// Unsupported : g_desktop_app_info_list_actions : no return type

// Removes all items in the menu.
/*

C function

g_menu_remove_all
*/
func (recv *Menu) RemoveAll() {
	C.g_menu_remove_all((*C.GMenu)(recv.native))

	return
}

// Sets (or unsets) the icon on @menu_item.
//
// This call is the same as calling g_icon_serialize() and using the
// result as the value to g_menu_item_set_attribute_value() for
// %G_MENU_ATTRIBUTE_ICON.
//
// This API is only intended for use with "noun" menu items; things like
// bookmarks or applications in an "Open With" menu.  Don't use it on
// menu items corresponding to verbs (eg: stock icons for 'Save' or
// 'Quit').
//
// If @icon is %NULL then the icon is unset.
/*

C function

g_menu_item_set_icon
*/
func (recv *MenuItem) SetIcon(icon *Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.g_menu_item_set_icon((*C.GMenuItem)(recv.native), c_icon)

	return
}

// A #GPropertyAction is a way to get a #GAction with a state value
// reflecting and controlling the value of a #GObject property.
//
// The state of the action will correspond to the value of the property.
// Changing it will change the property (assuming the requested value
// matches the requirements as specified in the #GParamSpec).
//
// Only the most common types are presently supported.  Booleans are
// mapped to booleans, strings to strings, signed/unsigned integers to
// int32/uint32 and floats and doubles to doubles.
//
// If the property is an enum then the state will be string-typed and
// conversion will automatically be performed between the enum value and
// "nick" string as per the #GEnumValue table.
//
// Flags types are not currently supported.
//
// Properties of object types, boxed types and pointer types are not
// supported and probably never will be.
//
// Properties of #GVariant types are not currently supported.
//
// If the property is boolean-valued then the action will have a NULL
// parameter type, and activating the action (with no parameter) will
// toggle the value of the property.
//
// In all other cases, the parameter type will correspond to the type of
// the property.
//
// The general idea here is to reduce the number of locations where a
// particular piece of state is kept (and therefore has to be synchronised
// between). #GPropertyAction does not have a separate state that is kept
// in sync with the property value -- its state is the property value.
//
// For example, it might be useful to create a #GAction corresponding to
// the "visible-child-name" property of a #GtkStack so that the current
// page can be switched from a menu.  The active radio indication in the
// menu is then directly determined from the active page of the
// #GtkStack.
//
// An anti-example would be binding the "active-id" property on a
// #GtkComboBox.  This is because the state of the combobox itself is
// probably uninteresting and is actually being used to control
// something else.
//
// Another anti-example would be to bind to the "visible-child-name"
// property of a #GtkStack if this value is actually stored in
// #GSettings.  In that case, the real source of the value is
// #GSettings.  If you want a #GAction to control a setting stored in
// #GSettings, see g_settings_create_action() instead, and possibly
// combine its use with g_settings_bind().
/*

C type

GPropertyAction
*/
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

// Creates a #GAction corresponding to the value of property
// @property_name on @object.
//
// The property must be existent and readable and writable (and not
// construct-only).
//
// This function takes a reference on @object and doesn't release it
// until the action is destroyed.
/*

C function

g_property_action_new
*/
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
