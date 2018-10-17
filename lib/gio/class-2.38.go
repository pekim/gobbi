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

// Unsupported signal : unsupported parameter info : no type generator for AppInfo,

// Unsupported signal : unsupported parameter files : no param type

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

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no param type

// GetPropertyInfo is a wrapper around the C function g_dbus_method_invocation_get_property_info.
func (recv *DBusMethodInvocation) GetPropertyInfo() *DBusPropertyInfo {
	retC := C.g_dbus_method_invocation_get_property_info((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusPropertyInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal : unsupported parameter changed_properties : Blacklisted record : GVariant

// Unsupported signal : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported signal : unsupported parameter changed_properties : Blacklisted record : GVariant

// Unsupported signal : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

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

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported signal : unsupported parameter file : no type generator for File,

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// RemoveAll is a wrapper around the C function g_menu_remove_all.
func (recv *Menu) RemoveAll() {
	C.g_menu_remove_all((*C.GMenu)(recv.native))

	return
}

// Unsupported : g_menu_item_set_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported signal : unsupported parameter choices : no param type

// Unsupported signal : unsupported parameter processes : no param type

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

// Unsupported signal : unsupported parameter keys : no param type

// Unsupported signal : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported signal : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported signal : unsupported parameter connectable : no type generator for SocketConnectable,

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no param type

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter volume : no type generator for Volume,

// Unsupported signal : unsupported parameter volume : no type generator for Volume,

// Unsupported signal : unsupported parameter volume : no type generator for Volume,
