// This is a generated file - DO NOT EDIT
// +build gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void action_activateHandler(GObject *, gpointer);

	static gulong Action_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(action_activateHandler), data);
	}

*/
/*

	void actiongroup_connectProxyHandler(GObject *, GtkAction *, GtkWidget *, gpointer);

	static gulong ActionGroup_signal_connect_connect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "connect-proxy", G_CALLBACK(actiongroup_connectProxyHandler), data);
	}

*/
/*

	void actiongroup_disconnectProxyHandler(GObject *, GtkAction *, GtkWidget *, gpointer);

	static gulong ActionGroup_signal_connect_disconnect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "disconnect-proxy", G_CALLBACK(actiongroup_disconnectProxyHandler), data);
	}

*/
/*

	void actiongroup_postActivateHandler(GObject *, GtkAction *, gpointer);

	static gulong ActionGroup_signal_connect_post_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "post-activate", G_CALLBACK(actiongroup_postActivateHandler), data);
	}

*/
/*

	void actiongroup_preActivateHandler(GObject *, GtkAction *, gpointer);

	static gulong ActionGroup_signal_connect_pre_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pre-activate", G_CALLBACK(actiongroup_preActivateHandler), data);
	}

*/
/*

	void style_realizeHandler(GObject *, gpointer);

	static gulong Style_signal_connect_realize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "realize", G_CALLBACK(style_realizeHandler), data);
	}

*/
/*

	void style_unrealizeHandler(GObject *, gpointer);

	static gulong Style_signal_connect_unrealize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unrealize", G_CALLBACK(style_unrealizeHandler), data);
	}

*/
import "C"

type signalActionActivateDetail struct {
	callback  ActionSignalActivateCallback
	handlerID C.gulong
}

var signalActionActivateId int
var signalActionActivateMap = make(map[int]signalActionActivateDetail)
var signalActionActivateLock sync.RWMutex

// ActionSignalActivateCallback is a callback function for a 'activate' signal emitted from a Action.
type ActionSignalActivateCallback func()

/*
ConnectActivate connects the callback to the 'activate' signal for the Action.

The returned value represents the connection, and may be passed to DisconnectActivate to remove it.
*/
func (recv *Action) ConnectActivate(callback ActionSignalActivateCallback) int {
	signalActionActivateLock.Lock()
	defer signalActionActivateLock.Unlock()

	signalActionActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.Action_signal_connect_activate(instance, C.gpointer(uintptr(signalActionActivateId)))

	detail := signalActionActivateDetail{callback, handlerID}
	signalActionActivateMap[signalActionActivateId] = detail

	return signalActionActivateId
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the Action.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *Action) DisconnectActivate(connectionID int) {
	signalActionActivateLock.Lock()
	defer signalActionActivateLock.Unlock()

	detail, exists := signalActionActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionActivateMap, connectionID)
}

//export action_activateHandler
func action_activateHandler(_ *C.GObject, data C.gpointer) {
	signalActionActivateLock.RLock()
	defer signalActionActivateLock.RUnlock()

	index := int(uintptr(data))
	callback := signalActionActivateMap[index].callback
	callback()
}

// Blacklisted : gtk_action_new

// Blacklisted : gtk_action_activate

// Blacklisted : gtk_action_connect_accelerator

// Blacklisted : gtk_action_create_icon

// Blacklisted : gtk_action_create_menu_item

// Blacklisted : gtk_action_create_tool_item

// Blacklisted : gtk_action_disconnect_accelerator

// Blacklisted : gtk_action_get_name

// Blacklisted : gtk_action_get_proxies

// Blacklisted : gtk_action_get_sensitive

// Blacklisted : gtk_action_get_visible

// Blacklisted : gtk_action_is_sensitive

// IsVisible is a wrapper around the C function gtk_action_is_visible.
func (recv *Action) IsVisible() bool {
	retC := C.gtk_action_is_visible((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : gtk_action_set_accel_group

// Blacklisted : gtk_action_set_accel_path

type signalActionGroupConnectProxyDetail struct {
	callback  ActionGroupSignalConnectProxyCallback
	handlerID C.gulong
}

var signalActionGroupConnectProxyId int
var signalActionGroupConnectProxyMap = make(map[int]signalActionGroupConnectProxyDetail)
var signalActionGroupConnectProxyLock sync.RWMutex

// ActionGroupSignalConnectProxyCallback is a callback function for a 'connect-proxy' signal emitted from a ActionGroup.
type ActionGroupSignalConnectProxyCallback func(action *Action, proxy *Widget)

/*
ConnectConnectProxy connects the callback to the 'connect-proxy' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectConnectProxy to remove it.
*/
func (recv *ActionGroup) ConnectConnectProxy(callback ActionGroupSignalConnectProxyCallback) int {
	signalActionGroupConnectProxyLock.Lock()
	defer signalActionGroupConnectProxyLock.Unlock()

	signalActionGroupConnectProxyId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_connect_proxy(instance, C.gpointer(uintptr(signalActionGroupConnectProxyId)))

	detail := signalActionGroupConnectProxyDetail{callback, handlerID}
	signalActionGroupConnectProxyMap[signalActionGroupConnectProxyId] = detail

	return signalActionGroupConnectProxyId
}

/*
DisconnectConnectProxy disconnects a callback from the 'connect-proxy' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectConnectProxy.
*/
func (recv *ActionGroup) DisconnectConnectProxy(connectionID int) {
	signalActionGroupConnectProxyLock.Lock()
	defer signalActionGroupConnectProxyLock.Unlock()

	detail, exists := signalActionGroupConnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupConnectProxyMap, connectionID)
}

//export actiongroup_connectProxyHandler
func actiongroup_connectProxyHandler(_ *C.GObject, c_action *C.GtkAction, c_proxy *C.GtkWidget, data C.gpointer) {
	signalActionGroupConnectProxyLock.RLock()
	defer signalActionGroupConnectProxyLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

	index := int(uintptr(data))
	callback := signalActionGroupConnectProxyMap[index].callback
	callback(action, proxy)
}

type signalActionGroupDisconnectProxyDetail struct {
	callback  ActionGroupSignalDisconnectProxyCallback
	handlerID C.gulong
}

var signalActionGroupDisconnectProxyId int
var signalActionGroupDisconnectProxyMap = make(map[int]signalActionGroupDisconnectProxyDetail)
var signalActionGroupDisconnectProxyLock sync.RWMutex

// ActionGroupSignalDisconnectProxyCallback is a callback function for a 'disconnect-proxy' signal emitted from a ActionGroup.
type ActionGroupSignalDisconnectProxyCallback func(action *Action, proxy *Widget)

/*
ConnectDisconnectProxy connects the callback to the 'disconnect-proxy' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectDisconnectProxy to remove it.
*/
func (recv *ActionGroup) ConnectDisconnectProxy(callback ActionGroupSignalDisconnectProxyCallback) int {
	signalActionGroupDisconnectProxyLock.Lock()
	defer signalActionGroupDisconnectProxyLock.Unlock()

	signalActionGroupDisconnectProxyId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_disconnect_proxy(instance, C.gpointer(uintptr(signalActionGroupDisconnectProxyId)))

	detail := signalActionGroupDisconnectProxyDetail{callback, handlerID}
	signalActionGroupDisconnectProxyMap[signalActionGroupDisconnectProxyId] = detail

	return signalActionGroupDisconnectProxyId
}

/*
DisconnectDisconnectProxy disconnects a callback from the 'disconnect-proxy' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectDisconnectProxy.
*/
func (recv *ActionGroup) DisconnectDisconnectProxy(connectionID int) {
	signalActionGroupDisconnectProxyLock.Lock()
	defer signalActionGroupDisconnectProxyLock.Unlock()

	detail, exists := signalActionGroupDisconnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupDisconnectProxyMap, connectionID)
}

//export actiongroup_disconnectProxyHandler
func actiongroup_disconnectProxyHandler(_ *C.GObject, c_action *C.GtkAction, c_proxy *C.GtkWidget, data C.gpointer) {
	signalActionGroupDisconnectProxyLock.RLock()
	defer signalActionGroupDisconnectProxyLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

	index := int(uintptr(data))
	callback := signalActionGroupDisconnectProxyMap[index].callback
	callback(action, proxy)
}

type signalActionGroupPostActivateDetail struct {
	callback  ActionGroupSignalPostActivateCallback
	handlerID C.gulong
}

var signalActionGroupPostActivateId int
var signalActionGroupPostActivateMap = make(map[int]signalActionGroupPostActivateDetail)
var signalActionGroupPostActivateLock sync.RWMutex

// ActionGroupSignalPostActivateCallback is a callback function for a 'post-activate' signal emitted from a ActionGroup.
type ActionGroupSignalPostActivateCallback func(action *Action)

/*
ConnectPostActivate connects the callback to the 'post-activate' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectPostActivate to remove it.
*/
func (recv *ActionGroup) ConnectPostActivate(callback ActionGroupSignalPostActivateCallback) int {
	signalActionGroupPostActivateLock.Lock()
	defer signalActionGroupPostActivateLock.Unlock()

	signalActionGroupPostActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_post_activate(instance, C.gpointer(uintptr(signalActionGroupPostActivateId)))

	detail := signalActionGroupPostActivateDetail{callback, handlerID}
	signalActionGroupPostActivateMap[signalActionGroupPostActivateId] = detail

	return signalActionGroupPostActivateId
}

/*
DisconnectPostActivate disconnects a callback from the 'post-activate' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectPostActivate.
*/
func (recv *ActionGroup) DisconnectPostActivate(connectionID int) {
	signalActionGroupPostActivateLock.Lock()
	defer signalActionGroupPostActivateLock.Unlock()

	detail, exists := signalActionGroupPostActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupPostActivateMap, connectionID)
}

//export actiongroup_postActivateHandler
func actiongroup_postActivateHandler(_ *C.GObject, c_action *C.GtkAction, data C.gpointer) {
	signalActionGroupPostActivateLock.RLock()
	defer signalActionGroupPostActivateLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	index := int(uintptr(data))
	callback := signalActionGroupPostActivateMap[index].callback
	callback(action)
}

type signalActionGroupPreActivateDetail struct {
	callback  ActionGroupSignalPreActivateCallback
	handlerID C.gulong
}

var signalActionGroupPreActivateId int
var signalActionGroupPreActivateMap = make(map[int]signalActionGroupPreActivateDetail)
var signalActionGroupPreActivateLock sync.RWMutex

// ActionGroupSignalPreActivateCallback is a callback function for a 'pre-activate' signal emitted from a ActionGroup.
type ActionGroupSignalPreActivateCallback func(action *Action)

/*
ConnectPreActivate connects the callback to the 'pre-activate' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectPreActivate to remove it.
*/
func (recv *ActionGroup) ConnectPreActivate(callback ActionGroupSignalPreActivateCallback) int {
	signalActionGroupPreActivateLock.Lock()
	defer signalActionGroupPreActivateLock.Unlock()

	signalActionGroupPreActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_pre_activate(instance, C.gpointer(uintptr(signalActionGroupPreActivateId)))

	detail := signalActionGroupPreActivateDetail{callback, handlerID}
	signalActionGroupPreActivateMap[signalActionGroupPreActivateId] = detail

	return signalActionGroupPreActivateId
}

/*
DisconnectPreActivate disconnects a callback from the 'pre-activate' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectPreActivate.
*/
func (recv *ActionGroup) DisconnectPreActivate(connectionID int) {
	signalActionGroupPreActivateLock.Lock()
	defer signalActionGroupPreActivateLock.Unlock()

	detail, exists := signalActionGroupPreActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupPreActivateMap, connectionID)
}

//export actiongroup_preActivateHandler
func actiongroup_preActivateHandler(_ *C.GObject, c_action *C.GtkAction, data C.gpointer) {
	signalActionGroupPreActivateLock.RLock()
	defer signalActionGroupPreActivateLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	index := int(uintptr(data))
	callback := signalActionGroupPreActivateMap[index].callback
	callback(action)
}

// Blacklisted : gtk_action_group_new

// Blacklisted : gtk_action_group_add_action

// Blacklisted : gtk_action_group_add_action_with_accel

// Unsupported : gtk_action_group_add_actions : unsupported parameter entries :

// Unsupported : gtk_action_group_add_actions_full : unsupported parameter entries :

// Unsupported : gtk_action_group_add_radio_actions : unsupported parameter entries :

// Unsupported : gtk_action_group_add_radio_actions_full : unsupported parameter entries :

// Unsupported : gtk_action_group_add_toggle_actions : unsupported parameter entries :

// Unsupported : gtk_action_group_add_toggle_actions_full : unsupported parameter entries :

// Blacklisted : gtk_action_group_get_action

// Blacklisted : gtk_action_group_get_name

// Blacklisted : gtk_action_group_get_sensitive

// Blacklisted : gtk_action_group_get_visible

// Blacklisted : gtk_action_group_list_actions

// Blacklisted : gtk_action_group_remove_action

// Blacklisted : gtk_action_group_set_sensitive

// Unsupported : gtk_action_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GtkTranslateFunc) for param func

// Blacklisted : gtk_action_group_set_translation_domain

// Blacklisted : gtk_action_group_set_visible

type signalStyleRealizeDetail struct {
	callback  StyleSignalRealizeCallback
	handlerID C.gulong
}

var signalStyleRealizeId int
var signalStyleRealizeMap = make(map[int]signalStyleRealizeDetail)
var signalStyleRealizeLock sync.RWMutex

// StyleSignalRealizeCallback is a callback function for a 'realize' signal emitted from a Style.
type StyleSignalRealizeCallback func()

/*
ConnectRealize connects the callback to the 'realize' signal for the Style.

The returned value represents the connection, and may be passed to DisconnectRealize to remove it.
*/
func (recv *Style) ConnectRealize(callback StyleSignalRealizeCallback) int {
	signalStyleRealizeLock.Lock()
	defer signalStyleRealizeLock.Unlock()

	signalStyleRealizeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Style_signal_connect_realize(instance, C.gpointer(uintptr(signalStyleRealizeId)))

	detail := signalStyleRealizeDetail{callback, handlerID}
	signalStyleRealizeMap[signalStyleRealizeId] = detail

	return signalStyleRealizeId
}

/*
DisconnectRealize disconnects a callback from the 'realize' signal for the Style.

The connectionID should be a value returned from a call to ConnectRealize.
*/
func (recv *Style) DisconnectRealize(connectionID int) {
	signalStyleRealizeLock.Lock()
	defer signalStyleRealizeLock.Unlock()

	detail, exists := signalStyleRealizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStyleRealizeMap, connectionID)
}

//export style_realizeHandler
func style_realizeHandler(_ *C.GObject, data C.gpointer) {
	signalStyleRealizeLock.RLock()
	defer signalStyleRealizeLock.RUnlock()

	index := int(uintptr(data))
	callback := signalStyleRealizeMap[index].callback
	callback()
}

type signalStyleUnrealizeDetail struct {
	callback  StyleSignalUnrealizeCallback
	handlerID C.gulong
}

var signalStyleUnrealizeId int
var signalStyleUnrealizeMap = make(map[int]signalStyleUnrealizeDetail)
var signalStyleUnrealizeLock sync.RWMutex

// StyleSignalUnrealizeCallback is a callback function for a 'unrealize' signal emitted from a Style.
type StyleSignalUnrealizeCallback func()

/*
ConnectUnrealize connects the callback to the 'unrealize' signal for the Style.

The returned value represents the connection, and may be passed to DisconnectUnrealize to remove it.
*/
func (recv *Style) ConnectUnrealize(callback StyleSignalUnrealizeCallback) int {
	signalStyleUnrealizeLock.Lock()
	defer signalStyleUnrealizeLock.Unlock()

	signalStyleUnrealizeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Style_signal_connect_unrealize(instance, C.gpointer(uintptr(signalStyleUnrealizeId)))

	detail := signalStyleUnrealizeDetail{callback, handlerID}
	signalStyleUnrealizeMap[signalStyleUnrealizeId] = detail

	return signalStyleUnrealizeId
}

/*
DisconnectUnrealize disconnects a callback from the 'unrealize' signal for the Style.

The connectionID should be a value returned from a call to ConnectUnrealize.
*/
func (recv *Style) DisconnectUnrealize(connectionID int) {
	signalStyleUnrealizeLock.Lock()
	defer signalStyleUnrealizeLock.Unlock()

	detail, exists := signalStyleUnrealizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStyleUnrealizeMap, connectionID)
}

//export style_unrealizeHandler
func style_unrealizeHandler(_ *C.GObject, data C.gpointer) {
	signalStyleUnrealizeLock.RLock()
	defer signalStyleUnrealizeLock.RUnlock()

	index := int(uintptr(data))
	callback := signalStyleUnrealizeMap[index].callback
	callback()
}

// Blacklisted : gtk_widget_add_mnemonic_label

// Blacklisted : gtk_widget_can_activate_accel

// Blacklisted : gtk_drag_source_get_target_list

// Blacklisted : gtk_drag_source_set_target_list

// Blacklisted : gtk_widget_get_no_show_all

// Blacklisted : gtk_widget_list_mnemonic_labels

// Blacklisted : gtk_widget_queue_resize_no_redraw

// Blacklisted : gtk_widget_remove_mnemonic_label

// Blacklisted : gtk_widget_set_no_show_all

// Blacklisted : gtk_window_set_default_icon

// Blacklisted : gtk_window_activate_key

// Blacklisted : gtk_window_get_accept_focus

// Blacklisted : gtk_window_has_toplevel_focus

// Blacklisted : gtk_window_is_active

// Blacklisted : gtk_window_propagate_key_event

// Blacklisted : gtk_window_set_accept_focus

// Blacklisted : gtk_window_set_keep_above

// Blacklisted : gtk_window_set_keep_below
