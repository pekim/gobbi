// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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
/*

	void menumodel_itemsChangedHandler(GObject *, gint, gint, gint, gpointer);

	static gulong MenuModel_signal_connect_items_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "items-changed", G_CALLBACK(menumodel_itemsChangedHandler), data);
	}

*/
/*

	void socketclient_eventHandler(GObject *, GSocketClientEvent, GSocketConnectable *, GIOStream *, gpointer);

	static gulong SocketClient_signal_connect_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "event", G_CALLBACK(socketclient_eventHandler), data);
	}

*/
import "C"

// Blacklisted : g_app_launch_context_get_environment

// Blacklisted : g_app_launch_context_setenv

// Blacklisted : g_app_launch_context_unsetenv

// Blacklisted : g_application_get_default

// Blacklisted : g_application_quit

// Blacklisted : g_application_set_default

// Blacklisted : g_dbus_action_group_get

// Blacklisted : g_dbus_connection_export_action_group

// Blacklisted : g_dbus_connection_export_menu_model

// Blacklisted : g_dbus_connection_unexport_action_group

// Blacklisted : g_dbus_connection_unexport_menu_model

// Blacklisted : g_dbus_interface_skeleton_get_connections

// Blacklisted : g_dbus_interface_skeleton_has_connection

// Blacklisted : g_dbus_interface_skeleton_unexport_from_connection

// Blacklisted : g_dbus_menu_model_get

// Blacklisted : g_desktop_app_info_get_keywords

// InetAddressMask is a wrapper around the C record GInetAddressMask.
type InetAddressMask struct {
	native *C.GInetAddressMask
	// parent_instance : record
	// Private : priv
}

func InetAddressMaskNewFromC(u unsafe.Pointer) *InetAddressMask {
	c := (*C.GInetAddressMask)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressMask{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InetAddressMask) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InetAddressMask) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InetAddressMask with another InetAddressMask, and returns true if they represent the same GObject.
func (recv *InetAddressMask) Equals(other *InetAddressMask) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *InetAddressMask) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to InetAddressMask.
// Exercise care, as this is a potentially dangerous function if the Object is not a InetAddressMask.
func CastToInetAddressMask(object *gobject.Object) *InetAddressMask {
	return InetAddressMaskNewFromC(object.ToC())
}

// Blacklisted : g_inet_address_mask_new

// Blacklisted : g_inet_address_mask_new_from_string

// Blacklisted : g_inet_address_mask_equal

// Blacklisted : g_inet_address_mask_get_address

// Blacklisted : g_inet_address_mask_get_family

// Blacklisted : g_inet_address_mask_get_length

// Blacklisted : g_inet_address_mask_matches

// Blacklisted : g_inet_address_mask_to_string

// Blacklisted : g_inet_socket_address_get_flowinfo

// Blacklisted : g_inet_socket_address_get_scope_id

// Menu is a wrapper around the C record GMenu.
type Menu struct {
	native *C.GMenu
}

func MenuNewFromC(u unsafe.Pointer) *Menu {
	c := (*C.GMenu)(u)
	if c == nil {
		return nil
	}

	g := &Menu{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Menu) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Menu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Menu with another Menu, and returns true if they represent the same GObject.
func (recv *Menu) Equals(other *Menu) bool {
	return other.ToC() == recv.ToC()
}

// MenuModel upcasts to *MenuModel
func (recv *Menu) MenuModel() *MenuModel {
	return MenuModelNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Menu) Object() *gobject.Object {
	return recv.MenuModel().Object()
}

// CastToWidget down casts any arbitrary Object to Menu.
// Exercise care, as this is a potentially dangerous function if the Object is not a Menu.
func CastToMenu(object *gobject.Object) *Menu {
	return MenuNewFromC(object.ToC())
}

// Blacklisted : g_menu_new

// Blacklisted : g_menu_append

// Blacklisted : g_menu_append_item

// Blacklisted : g_menu_append_section

// Blacklisted : g_menu_append_submenu

// Blacklisted : g_menu_freeze

// Blacklisted : g_menu_insert

// Blacklisted : g_menu_insert_item

// Blacklisted : g_menu_insert_section

// Blacklisted : g_menu_insert_submenu

// Blacklisted : g_menu_prepend

// Blacklisted : g_menu_prepend_item

// Blacklisted : g_menu_prepend_section

// Blacklisted : g_menu_prepend_submenu

// Blacklisted : g_menu_remove

// MenuAttributeIter is a wrapper around the C record GMenuAttributeIter.
type MenuAttributeIter struct {
	native *C.GMenuAttributeIter
	// parent_instance : record
	// priv : record
}

func MenuAttributeIterNewFromC(u unsafe.Pointer) *MenuAttributeIter {
	c := (*C.GMenuAttributeIter)(u)
	if c == nil {
		return nil
	}

	g := &MenuAttributeIter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuAttributeIter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuAttributeIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MenuAttributeIter with another MenuAttributeIter, and returns true if they represent the same GObject.
func (recv *MenuAttributeIter) Equals(other *MenuAttributeIter) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *MenuAttributeIter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to MenuAttributeIter.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuAttributeIter.
func CastToMenuAttributeIter(object *gobject.Object) *MenuAttributeIter {
	return MenuAttributeIterNewFromC(object.ToC())
}

// Blacklisted : g_menu_attribute_iter_get_name

// Blacklisted : g_menu_attribute_iter_get_next

// Blacklisted : g_menu_attribute_iter_get_value

// Blacklisted : g_menu_attribute_iter_next

// MenuItem is a wrapper around the C record GMenuItem.
type MenuItem struct {
	native *C.GMenuItem
}

func MenuItemNewFromC(u unsafe.Pointer) *MenuItem {
	c := (*C.GMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &MenuItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MenuItem with another MenuItem, and returns true if they represent the same GObject.
func (recv *MenuItem) Equals(other *MenuItem) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *MenuItem) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to MenuItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuItem.
func CastToMenuItem(object *gobject.Object) *MenuItem {
	return MenuItemNewFromC(object.ToC())
}

// Blacklisted : g_menu_item_new

// Blacklisted : g_menu_item_new_section

// Blacklisted : g_menu_item_new_submenu

// Blacklisted : g_menu_item_set_action_and_target

// Blacklisted : g_menu_item_set_action_and_target_value

// Blacklisted : g_menu_item_set_attribute

// Blacklisted : g_menu_item_set_attribute_value

// Blacklisted : g_menu_item_set_detailed_action

// Blacklisted : g_menu_item_set_label

// Blacklisted : g_menu_item_set_link

// Blacklisted : g_menu_item_set_section

// Blacklisted : g_menu_item_set_submenu

// MenuLinkIter is a wrapper around the C record GMenuLinkIter.
type MenuLinkIter struct {
	native *C.GMenuLinkIter
	// parent_instance : record
	// priv : record
}

func MenuLinkIterNewFromC(u unsafe.Pointer) *MenuLinkIter {
	c := (*C.GMenuLinkIter)(u)
	if c == nil {
		return nil
	}

	g := &MenuLinkIter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuLinkIter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuLinkIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MenuLinkIter with another MenuLinkIter, and returns true if they represent the same GObject.
func (recv *MenuLinkIter) Equals(other *MenuLinkIter) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *MenuLinkIter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to MenuLinkIter.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuLinkIter.
func CastToMenuLinkIter(object *gobject.Object) *MenuLinkIter {
	return MenuLinkIterNewFromC(object.ToC())
}

// Blacklisted : g_menu_link_iter_get_name

// Blacklisted : g_menu_link_iter_get_next

// Blacklisted : g_menu_link_iter_get_value

// Blacklisted : g_menu_link_iter_next

// MenuModel is a wrapper around the C record GMenuModel.
type MenuModel struct {
	native *C.GMenuModel
	// parent_instance : record
	// priv : record
}

func MenuModelNewFromC(u unsafe.Pointer) *MenuModel {
	c := (*C.GMenuModel)(u)
	if c == nil {
		return nil
	}

	g := &MenuModel{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuModel) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MenuModel with another MenuModel, and returns true if they represent the same GObject.
func (recv *MenuModel) Equals(other *MenuModel) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *MenuModel) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to MenuModel.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuModel.
func CastToMenuModel(object *gobject.Object) *MenuModel {
	return MenuModelNewFromC(object.ToC())
}

type signalMenuModelItemsChangedDetail struct {
	callback  MenuModelSignalItemsChangedCallback
	handlerID C.gulong
}

var signalMenuModelItemsChangedId int
var signalMenuModelItemsChangedMap = make(map[int]signalMenuModelItemsChangedDetail)
var signalMenuModelItemsChangedLock sync.RWMutex

// MenuModelSignalItemsChangedCallback is a callback function for a 'items-changed' signal emitted from a MenuModel.
type MenuModelSignalItemsChangedCallback func(position int32, removed int32, added int32)

/*
ConnectItemsChanged connects the callback to the 'items-changed' signal for the MenuModel.

The returned value represents the connection, and may be passed to DisconnectItemsChanged to remove it.
*/
func (recv *MenuModel) ConnectItemsChanged(callback MenuModelSignalItemsChangedCallback) int {
	signalMenuModelItemsChangedLock.Lock()
	defer signalMenuModelItemsChangedLock.Unlock()

	signalMenuModelItemsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.MenuModel_signal_connect_items_changed(instance, C.gpointer(uintptr(signalMenuModelItemsChangedId)))

	detail := signalMenuModelItemsChangedDetail{callback, handlerID}
	signalMenuModelItemsChangedMap[signalMenuModelItemsChangedId] = detail

	return signalMenuModelItemsChangedId
}

/*
DisconnectItemsChanged disconnects a callback from the 'items-changed' signal for the MenuModel.

The connectionID should be a value returned from a call to ConnectItemsChanged.
*/
func (recv *MenuModel) DisconnectItemsChanged(connectionID int) {
	signalMenuModelItemsChangedLock.Lock()
	defer signalMenuModelItemsChangedLock.Unlock()

	detail, exists := signalMenuModelItemsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMenuModelItemsChangedMap, connectionID)
}

//export menumodel_itemsChangedHandler
func menumodel_itemsChangedHandler(_ *C.GObject, c_position C.gint, c_removed C.gint, c_added C.gint, data C.gpointer) {
	signalMenuModelItemsChangedLock.RLock()
	defer signalMenuModelItemsChangedLock.RUnlock()

	position := int32(c_position)

	removed := int32(c_removed)

	added := int32(c_added)

	index := int(uintptr(data))
	callback := signalMenuModelItemsChangedMap[index].callback
	callback(position, removed, added)
}

// Blacklisted : g_menu_model_get_item_attribute

// Blacklisted : g_menu_model_get_item_attribute_value

// Blacklisted : g_menu_model_get_item_link

// Blacklisted : g_menu_model_get_n_items

// Blacklisted : g_menu_model_is_mutable

// Blacklisted : g_menu_model_items_changed

// Blacklisted : g_menu_model_iterate_item_attributes

// Blacklisted : g_menu_model_iterate_item_links

// Blacklisted : g_settings_new_full

// Blacklisted : g_settings_create_action

// Blacklisted : g_simple_async_result_set_check_cancellable

// Blacklisted : g_socket_condition_timed_wait

// Blacklisted : g_socket_get_available_bytes

// Blacklisted : g_socket_get_broadcast

// Blacklisted : g_socket_get_multicast_loopback

// Blacklisted : g_socket_get_multicast_ttl

// Blacklisted : g_socket_get_ttl

// Blacklisted : g_socket_join_multicast_group

// Blacklisted : g_socket_leave_multicast_group

// Blacklisted : g_socket_set_broadcast

// Blacklisted : g_socket_set_multicast_loopback

// Blacklisted : g_socket_set_multicast_ttl

// Blacklisted : g_socket_set_ttl

type signalSocketClientEventDetail struct {
	callback  SocketClientSignalEventCallback
	handlerID C.gulong
}

var signalSocketClientEventId int
var signalSocketClientEventMap = make(map[int]signalSocketClientEventDetail)
var signalSocketClientEventLock sync.RWMutex

// SocketClientSignalEventCallback is a callback function for a 'event' signal emitted from a SocketClient.
type SocketClientSignalEventCallback func(event SocketClientEvent, connectable *SocketConnectable, connection *IOStream)

/*
ConnectEvent connects the callback to the 'event' signal for the SocketClient.

The returned value represents the connection, and may be passed to DisconnectEvent to remove it.
*/
func (recv *SocketClient) ConnectEvent(callback SocketClientSignalEventCallback) int {
	signalSocketClientEventLock.Lock()
	defer signalSocketClientEventLock.Unlock()

	signalSocketClientEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.SocketClient_signal_connect_event(instance, C.gpointer(uintptr(signalSocketClientEventId)))

	detail := signalSocketClientEventDetail{callback, handlerID}
	signalSocketClientEventMap[signalSocketClientEventId] = detail

	return signalSocketClientEventId
}

/*
DisconnectEvent disconnects a callback from the 'event' signal for the SocketClient.

The connectionID should be a value returned from a call to ConnectEvent.
*/
func (recv *SocketClient) DisconnectEvent(connectionID int) {
	signalSocketClientEventLock.Lock()
	defer signalSocketClientEventLock.Unlock()

	detail, exists := signalSocketClientEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSocketClientEventMap, connectionID)
}

//export socketclient_eventHandler
func socketclient_eventHandler(_ *C.GObject, c_event C.GSocketClientEvent, c_connectable *C.GSocketConnectable, c_connection *C.GIOStream, data C.gpointer) {
	signalSocketClientEventLock.RLock()
	defer signalSocketClientEventLock.RUnlock()

	event := SocketClientEvent(c_event)

	connectable := SocketConnectableNewFromC(unsafe.Pointer(c_connectable))

	connection := IOStreamNewFromC(unsafe.Pointer(c_connection))

	index := int(uintptr(data))
	callback := signalSocketClientEventMap[index].callback
	callback(event, connectable, connection)
}

// Blacklisted : g_socket_connection_connect

// Unsupported : g_socket_connection_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_socket_connection_connect_finish

// Blacklisted : g_socket_connection_is_connected

// Unsupported : g_unix_connection_receive_credentials_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_unix_connection_receive_credentials_finish

// Unsupported : g_unix_connection_send_credentials_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_unix_connection_send_credentials_finish
