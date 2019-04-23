// This is a generated file - DO NOT EDIT
// +build gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
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

	void accelmap_changedHandler(GObject *, gchar*, guint, guint, gpointer);

	static gulong AccelMap_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(accelmap_changedHandler), data);
	}

*/
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

	void cellrenderer_editingCanceledHandler(GObject *, gpointer);

	static gulong CellRenderer_signal_connect_editing_canceled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "editing-canceled", G_CALLBACK(cellrenderer_editingCanceledHandler), data);
	}

*/
/*

	void colorbutton_colorSetHandler(GObject *, gpointer);

	static gulong ColorButton_signal_connect_color_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "color-set", G_CALLBACK(colorbutton_colorSetHandler), data);
	}

*/
/*

	void combobox_changedHandler(GObject *, gpointer);

	static gulong ComboBox_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(combobox_changedHandler), data);
	}

*/
/*

	void entrycompletion_actionActivatedHandler(GObject *, gint, gpointer);

	static gulong EntryCompletion_signal_connect_action_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-activated", G_CALLBACK(entrycompletion_actionActivatedHandler), data);
	}

*/
/*

	gboolean entrycompletion_matchSelectedHandler(GObject *, GtkTreeModel *, GtkTreeIter *, gpointer);

	static gulong EntryCompletion_signal_connect_match_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "match-selected", G_CALLBACK(entrycompletion_matchSelectedHandler), data);
	}

*/
/*

	void fontbutton_fontSetHandler(GObject *, gpointer);

	static gulong FontButton_signal_connect_font_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "font-set", G_CALLBACK(fontbutton_fontSetHandler), data);
	}

*/
/*

	void radioaction_changedHandler(GObject *, GtkRadioAction *, gpointer);

	static gulong RadioAction_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(radioaction_changedHandler), data);
	}

*/
/*

	void radiobutton_groupChangedHandler(GObject *, gpointer);

	static gulong RadioButton_signal_connect_group_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "group-changed", G_CALLBACK(radiobutton_groupChangedHandler), data);
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
/*

	void uimanager_actionsChangedHandler(GObject *, gpointer);

	static gulong UIManager_signal_connect_actions_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "actions-changed", G_CALLBACK(uimanager_actionsChangedHandler), data);
	}

*/
/*

	void uimanager_addWidgetHandler(GObject *, GtkWidget *, gpointer);

	static gulong UIManager_signal_connect_add_widget(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "add-widget", G_CALLBACK(uimanager_addWidgetHandler), data);
	}

*/
/*

	void uimanager_connectProxyHandler(GObject *, GtkAction *, GtkWidget *, gpointer);

	static gulong UIManager_signal_connect_connect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "connect-proxy", G_CALLBACK(uimanager_connectProxyHandler), data);
	}

*/
/*

	void uimanager_disconnectProxyHandler(GObject *, GtkAction *, GtkWidget *, gpointer);

	static gulong UIManager_signal_connect_disconnect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "disconnect-proxy", G_CALLBACK(uimanager_disconnectProxyHandler), data);
	}

*/
/*

	void uimanager_postActivateHandler(GObject *, GtkAction *, gpointer);

	static gulong UIManager_signal_connect_post_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "post-activate", G_CALLBACK(uimanager_postActivateHandler), data);
	}

*/
/*

	void uimanager_preActivateHandler(GObject *, GtkAction *, gpointer);

	static gulong UIManager_signal_connect_pre_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pre-activate", G_CALLBACK(uimanager_preActivateHandler), data);
	}

*/
import "C"

type signalAccelMapChangedDetail struct {
	callback  AccelMapSignalChangedCallback
	handlerID C.gulong
}

var signalAccelMapChangedId int
var signalAccelMapChangedMap = make(map[int]signalAccelMapChangedDetail)
var signalAccelMapChangedLock sync.RWMutex

// AccelMapSignalChangedCallback is a callback function for a 'changed' signal emitted from a AccelMap.
type AccelMapSignalChangedCallback func(accelPath string, accelKey uint32, accelMods gdk.ModifierType)

/*
ConnectChanged connects the callback to the 'changed' signal for the AccelMap.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *AccelMap) ConnectChanged(callback AccelMapSignalChangedCallback) int {
	signalAccelMapChangedLock.Lock()
	defer signalAccelMapChangedLock.Unlock()

	signalAccelMapChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.AccelMap_signal_connect_changed(instance, C.gpointer(uintptr(signalAccelMapChangedId)))

	detail := signalAccelMapChangedDetail{callback, handlerID}
	signalAccelMapChangedMap[signalAccelMapChangedId] = detail

	return signalAccelMapChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the AccelMap.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *AccelMap) DisconnectChanged(connectionID int) {
	signalAccelMapChangedLock.Lock()
	defer signalAccelMapChangedLock.Unlock()

	detail, exists := signalAccelMapChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAccelMapChangedMap, connectionID)
}

//export accelmap_changedHandler
func accelmap_changedHandler(_ *C.GObject, c_accel_path *C.gchar, c_accel_key C.guint, c_accel_mods C.guint, data C.gpointer) {
	signalAccelMapChangedLock.RLock()
	defer signalAccelMapChangedLock.RUnlock()

	accelPath := C.GoString(c_accel_path)

	accelKey := uint32(c_accel_key)

	accelMods := gdk.ModifierType(c_accel_mods)

	index := int(uintptr(data))
	callback := signalAccelMapChangedMap[index].callback
	callback(accelPath, accelKey, accelMods)
}

// Blacklisted : gtk_accel_map_get

// Blacklisted : gtk_accel_map_lock_path

// Blacklisted : gtk_accel_map_unlock_path

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

// Blacklisted : gtk_alignment_get_padding

// Blacklisted : gtk_alignment_set_padding

// Blacklisted : gtk_button_get_alignment

// Blacklisted : gtk_button_get_focus_on_click

// Blacklisted : gtk_button_set_alignment

// Blacklisted : gtk_button_set_focus_on_click

// Blacklisted : gtk_button_box_get_child_secondary

// Blacklisted : gtk_calendar_get_display_options

// Blacklisted : gtk_calendar_set_display_options

type signalCellRendererEditingCanceledDetail struct {
	callback  CellRendererSignalEditingCanceledCallback
	handlerID C.gulong
}

var signalCellRendererEditingCanceledId int
var signalCellRendererEditingCanceledMap = make(map[int]signalCellRendererEditingCanceledDetail)
var signalCellRendererEditingCanceledLock sync.RWMutex

// CellRendererSignalEditingCanceledCallback is a callback function for a 'editing-canceled' signal emitted from a CellRenderer.
type CellRendererSignalEditingCanceledCallback func()

/*
ConnectEditingCanceled connects the callback to the 'editing-canceled' signal for the CellRenderer.

The returned value represents the connection, and may be passed to DisconnectEditingCanceled to remove it.
*/
func (recv *CellRenderer) ConnectEditingCanceled(callback CellRendererSignalEditingCanceledCallback) int {
	signalCellRendererEditingCanceledLock.Lock()
	defer signalCellRendererEditingCanceledLock.Unlock()

	signalCellRendererEditingCanceledId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellRenderer_signal_connect_editing_canceled(instance, C.gpointer(uintptr(signalCellRendererEditingCanceledId)))

	detail := signalCellRendererEditingCanceledDetail{callback, handlerID}
	signalCellRendererEditingCanceledMap[signalCellRendererEditingCanceledId] = detail

	return signalCellRendererEditingCanceledId
}

/*
DisconnectEditingCanceled disconnects a callback from the 'editing-canceled' signal for the CellRenderer.

The connectionID should be a value returned from a call to ConnectEditingCanceled.
*/
func (recv *CellRenderer) DisconnectEditingCanceled(connectionID int) {
	signalCellRendererEditingCanceledLock.Lock()
	defer signalCellRendererEditingCanceledLock.Unlock()

	detail, exists := signalCellRendererEditingCanceledMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellRendererEditingCanceledMap, connectionID)
}

//export cellrenderer_editingCanceledHandler
func cellrenderer_editingCanceledHandler(_ *C.GObject, data C.gpointer) {
	signalCellRendererEditingCanceledLock.RLock()
	defer signalCellRendererEditingCanceledLock.RUnlock()

	index := int(uintptr(data))
	callback := signalCellRendererEditingCanceledMap[index].callback
	callback()
}

// Blacklisted : gtk_check_menu_item_get_draw_as_radio

// Blacklisted : gtk_check_menu_item_set_draw_as_radio

// Unsupported : gtk_clipboard_request_targets : unsupported parameter callback : no type generator for ClipboardTargetsReceivedFunc (GtkClipboardTargetsReceivedFunc) for param callback

// Unsupported : gtk_clipboard_wait_for_targets : unsupported parameter targets : output array param targets

type signalColorButtonColorSetDetail struct {
	callback  ColorButtonSignalColorSetCallback
	handlerID C.gulong
}

var signalColorButtonColorSetId int
var signalColorButtonColorSetMap = make(map[int]signalColorButtonColorSetDetail)
var signalColorButtonColorSetLock sync.RWMutex

// ColorButtonSignalColorSetCallback is a callback function for a 'color-set' signal emitted from a ColorButton.
type ColorButtonSignalColorSetCallback func()

/*
ConnectColorSet connects the callback to the 'color-set' signal for the ColorButton.

The returned value represents the connection, and may be passed to DisconnectColorSet to remove it.
*/
func (recv *ColorButton) ConnectColorSet(callback ColorButtonSignalColorSetCallback) int {
	signalColorButtonColorSetLock.Lock()
	defer signalColorButtonColorSetLock.Unlock()

	signalColorButtonColorSetId++
	instance := C.gpointer(recv.native)
	handlerID := C.ColorButton_signal_connect_color_set(instance, C.gpointer(uintptr(signalColorButtonColorSetId)))

	detail := signalColorButtonColorSetDetail{callback, handlerID}
	signalColorButtonColorSetMap[signalColorButtonColorSetId] = detail

	return signalColorButtonColorSetId
}

/*
DisconnectColorSet disconnects a callback from the 'color-set' signal for the ColorButton.

The connectionID should be a value returned from a call to ConnectColorSet.
*/
func (recv *ColorButton) DisconnectColorSet(connectionID int) {
	signalColorButtonColorSetLock.Lock()
	defer signalColorButtonColorSetLock.Unlock()

	detail, exists := signalColorButtonColorSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalColorButtonColorSetMap, connectionID)
}

//export colorbutton_colorSetHandler
func colorbutton_colorSetHandler(_ *C.GObject, data C.gpointer) {
	signalColorButtonColorSetLock.RLock()
	defer signalColorButtonColorSetLock.RUnlock()

	index := int(uintptr(data))
	callback := signalColorButtonColorSetMap[index].callback
	callback()
}

// Blacklisted : gtk_color_button_new

// Blacklisted : gtk_color_button_new_with_color

// Blacklisted : gtk_color_button_get_alpha

// Blacklisted : gtk_color_button_get_color

// GetTitle is a wrapper around the C function gtk_color_button_get_title.
func (recv *ColorButton) GetTitle() string {
	retC := C.gtk_color_button_get_title((*C.GtkColorButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : gtk_color_button_get_use_alpha

// Blacklisted : gtk_color_button_set_alpha

// Blacklisted : gtk_color_button_set_color

// SetTitle is a wrapper around the C function gtk_color_button_set_title.
func (recv *ColorButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_color_button_set_title((*C.GtkColorButton)(recv.native), c_title)

	return
}

// Blacklisted : gtk_color_button_set_use_alpha

type signalComboBoxChangedDetail struct {
	callback  ComboBoxSignalChangedCallback
	handlerID C.gulong
}

var signalComboBoxChangedId int
var signalComboBoxChangedMap = make(map[int]signalComboBoxChangedDetail)
var signalComboBoxChangedLock sync.RWMutex

// ComboBoxSignalChangedCallback is a callback function for a 'changed' signal emitted from a ComboBox.
type ComboBoxSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the ComboBox.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *ComboBox) ConnectChanged(callback ComboBoxSignalChangedCallback) int {
	signalComboBoxChangedLock.Lock()
	defer signalComboBoxChangedLock.Unlock()

	signalComboBoxChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ComboBox_signal_connect_changed(instance, C.gpointer(uintptr(signalComboBoxChangedId)))

	detail := signalComboBoxChangedDetail{callback, handlerID}
	signalComboBoxChangedMap[signalComboBoxChangedId] = detail

	return signalComboBoxChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *ComboBox) DisconnectChanged(connectionID int) {
	signalComboBoxChangedLock.Lock()
	defer signalComboBoxChangedLock.Unlock()

	detail, exists := signalComboBoxChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComboBoxChangedMap, connectionID)
}

//export combobox_changedHandler
func combobox_changedHandler(_ *C.GObject, data C.gpointer) {
	signalComboBoxChangedLock.RLock()
	defer signalComboBoxChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalComboBoxChangedMap[index].callback
	callback()
}

// Blacklisted : gtk_combo_box_new

// Blacklisted : gtk_combo_box_new_with_model

// Blacklisted : gtk_combo_box_get_active

// Blacklisted : gtk_combo_box_get_active_iter

// Blacklisted : gtk_combo_box_get_model

// Blacklisted : gtk_combo_box_popdown

// Blacklisted : gtk_combo_box_popup

// Blacklisted : gtk_combo_box_set_active

// Blacklisted : gtk_combo_box_set_active_iter

// Blacklisted : gtk_combo_box_set_column_span_column

// Blacklisted : gtk_combo_box_set_model

// Blacklisted : gtk_combo_box_set_row_span_column

// Blacklisted : gtk_combo_box_set_wrap_width

// Blacklisted : gtk_entry_get_alignment

// Blacklisted : gtk_entry_get_completion

// Blacklisted : gtk_entry_set_alignment

// Blacklisted : gtk_entry_set_completion

type signalEntryCompletionActionActivatedDetail struct {
	callback  EntryCompletionSignalActionActivatedCallback
	handlerID C.gulong
}

var signalEntryCompletionActionActivatedId int
var signalEntryCompletionActionActivatedMap = make(map[int]signalEntryCompletionActionActivatedDetail)
var signalEntryCompletionActionActivatedLock sync.RWMutex

// EntryCompletionSignalActionActivatedCallback is a callback function for a 'action-activated' signal emitted from a EntryCompletion.
type EntryCompletionSignalActionActivatedCallback func(Index int32)

/*
ConnectActionActivated connects the callback to the 'action-activated' signal for the EntryCompletion.

The returned value represents the connection, and may be passed to DisconnectActionActivated to remove it.
*/
func (recv *EntryCompletion) ConnectActionActivated(callback EntryCompletionSignalActionActivatedCallback) int {
	signalEntryCompletionActionActivatedLock.Lock()
	defer signalEntryCompletionActionActivatedLock.Unlock()

	signalEntryCompletionActionActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryCompletion_signal_connect_action_activated(instance, C.gpointer(uintptr(signalEntryCompletionActionActivatedId)))

	detail := signalEntryCompletionActionActivatedDetail{callback, handlerID}
	signalEntryCompletionActionActivatedMap[signalEntryCompletionActionActivatedId] = detail

	return signalEntryCompletionActionActivatedId
}

/*
DisconnectActionActivated disconnects a callback from the 'action-activated' signal for the EntryCompletion.

The connectionID should be a value returned from a call to ConnectActionActivated.
*/
func (recv *EntryCompletion) DisconnectActionActivated(connectionID int) {
	signalEntryCompletionActionActivatedLock.Lock()
	defer signalEntryCompletionActionActivatedLock.Unlock()

	detail, exists := signalEntryCompletionActionActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryCompletionActionActivatedMap, connectionID)
}

//export entrycompletion_actionActivatedHandler
func entrycompletion_actionActivatedHandler(_ *C.GObject, c__index C.gint, data C.gpointer) {
	signalEntryCompletionActionActivatedLock.RLock()
	defer signalEntryCompletionActionActivatedLock.RUnlock()

	Index := int32(c__index)

	index := int(uintptr(data))
	callback := signalEntryCompletionActionActivatedMap[index].callback
	callback(Index)
}

type signalEntryCompletionMatchSelectedDetail struct {
	callback  EntryCompletionSignalMatchSelectedCallback
	handlerID C.gulong
}

var signalEntryCompletionMatchSelectedId int
var signalEntryCompletionMatchSelectedMap = make(map[int]signalEntryCompletionMatchSelectedDetail)
var signalEntryCompletionMatchSelectedLock sync.RWMutex

// EntryCompletionSignalMatchSelectedCallback is a callback function for a 'match-selected' signal emitted from a EntryCompletion.
type EntryCompletionSignalMatchSelectedCallback func(model *TreeModel, iter *TreeIter) bool

/*
ConnectMatchSelected connects the callback to the 'match-selected' signal for the EntryCompletion.

The returned value represents the connection, and may be passed to DisconnectMatchSelected to remove it.
*/
func (recv *EntryCompletion) ConnectMatchSelected(callback EntryCompletionSignalMatchSelectedCallback) int {
	signalEntryCompletionMatchSelectedLock.Lock()
	defer signalEntryCompletionMatchSelectedLock.Unlock()

	signalEntryCompletionMatchSelectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryCompletion_signal_connect_match_selected(instance, C.gpointer(uintptr(signalEntryCompletionMatchSelectedId)))

	detail := signalEntryCompletionMatchSelectedDetail{callback, handlerID}
	signalEntryCompletionMatchSelectedMap[signalEntryCompletionMatchSelectedId] = detail

	return signalEntryCompletionMatchSelectedId
}

/*
DisconnectMatchSelected disconnects a callback from the 'match-selected' signal for the EntryCompletion.

The connectionID should be a value returned from a call to ConnectMatchSelected.
*/
func (recv *EntryCompletion) DisconnectMatchSelected(connectionID int) {
	signalEntryCompletionMatchSelectedLock.Lock()
	defer signalEntryCompletionMatchSelectedLock.Unlock()

	detail, exists := signalEntryCompletionMatchSelectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryCompletionMatchSelectedMap, connectionID)
}

//export entrycompletion_matchSelectedHandler
func entrycompletion_matchSelectedHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, data C.gpointer) C.gboolean {
	signalEntryCompletionMatchSelectedLock.RLock()
	defer signalEntryCompletionMatchSelectedLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalEntryCompletionMatchSelectedMap[index].callback
	retGo := callback(model, iter)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_entry_completion_new

// Blacklisted : gtk_entry_completion_complete

// Blacklisted : gtk_entry_completion_delete_action

// Blacklisted : gtk_entry_completion_get_entry

// Blacklisted : gtk_entry_completion_get_minimum_key_length

// Blacklisted : gtk_entry_completion_get_model

// Blacklisted : gtk_entry_completion_insert_action_markup

// Blacklisted : gtk_entry_completion_insert_action_text

// Unsupported : gtk_entry_completion_set_match_func : unsupported parameter func : no type generator for EntryCompletionMatchFunc (GtkEntryCompletionMatchFunc) for param func

// Blacklisted : gtk_entry_completion_set_minimum_key_length

// Blacklisted : gtk_entry_completion_set_model

// Blacklisted : gtk_entry_completion_set_text_column

// Blacklisted : gtk_event_box_get_above_child

// Blacklisted : gtk_event_box_get_visible_window

// Blacklisted : gtk_event_box_set_above_child

// Blacklisted : gtk_event_box_set_visible_window

// Blacklisted : gtk_expander_new

// Blacklisted : gtk_expander_new_with_mnemonic

// Blacklisted : gtk_expander_get_expanded

// Blacklisted : gtk_expander_get_label

// Blacklisted : gtk_expander_get_label_widget

// Blacklisted : gtk_expander_get_spacing

// Blacklisted : gtk_expander_get_use_markup

// Blacklisted : gtk_expander_get_use_underline

// Blacklisted : gtk_expander_set_expanded

// Blacklisted : gtk_expander_set_label

// Blacklisted : gtk_expander_set_label_widget

// Blacklisted : gtk_expander_set_spacing

// Blacklisted : gtk_expander_set_use_markup

// Blacklisted : gtk_expander_set_use_underline

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Blacklisted : gtk_file_chooser_widget_new

// Blacklisted : gtk_file_filter_new

// Unsupported : gtk_file_filter_add_custom : unsupported parameter func : no type generator for FileFilterFunc (GtkFileFilterFunc) for param func

// Blacklisted : gtk_file_filter_add_mime_type

// Blacklisted : gtk_file_filter_add_pattern

// Blacklisted : gtk_file_filter_filter

// Blacklisted : gtk_file_filter_get_name

// Blacklisted : gtk_file_filter_get_needed

// Blacklisted : gtk_file_filter_set_name

type signalFontButtonFontSetDetail struct {
	callback  FontButtonSignalFontSetCallback
	handlerID C.gulong
}

var signalFontButtonFontSetId int
var signalFontButtonFontSetMap = make(map[int]signalFontButtonFontSetDetail)
var signalFontButtonFontSetLock sync.RWMutex

// FontButtonSignalFontSetCallback is a callback function for a 'font-set' signal emitted from a FontButton.
type FontButtonSignalFontSetCallback func()

/*
ConnectFontSet connects the callback to the 'font-set' signal for the FontButton.

The returned value represents the connection, and may be passed to DisconnectFontSet to remove it.
*/
func (recv *FontButton) ConnectFontSet(callback FontButtonSignalFontSetCallback) int {
	signalFontButtonFontSetLock.Lock()
	defer signalFontButtonFontSetLock.Unlock()

	signalFontButtonFontSetId++
	instance := C.gpointer(recv.native)
	handlerID := C.FontButton_signal_connect_font_set(instance, C.gpointer(uintptr(signalFontButtonFontSetId)))

	detail := signalFontButtonFontSetDetail{callback, handlerID}
	signalFontButtonFontSetMap[signalFontButtonFontSetId] = detail

	return signalFontButtonFontSetId
}

/*
DisconnectFontSet disconnects a callback from the 'font-set' signal for the FontButton.

The connectionID should be a value returned from a call to ConnectFontSet.
*/
func (recv *FontButton) DisconnectFontSet(connectionID int) {
	signalFontButtonFontSetLock.Lock()
	defer signalFontButtonFontSetLock.Unlock()

	detail, exists := signalFontButtonFontSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFontButtonFontSetMap, connectionID)
}

//export fontbutton_fontSetHandler
func fontbutton_fontSetHandler(_ *C.GObject, data C.gpointer) {
	signalFontButtonFontSetLock.RLock()
	defer signalFontButtonFontSetLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFontButtonFontSetMap[index].callback
	callback()
}

// Blacklisted : gtk_font_button_new

// Blacklisted : gtk_font_button_new_with_font

// Blacklisted : gtk_font_button_get_font_name

// Blacklisted : gtk_font_button_get_show_size

// Blacklisted : gtk_font_button_get_show_style

// GetTitle is a wrapper around the C function gtk_font_button_get_title.
func (recv *FontButton) GetTitle() string {
	retC := C.gtk_font_button_get_title((*C.GtkFontButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : gtk_font_button_get_use_font

// Blacklisted : gtk_font_button_get_use_size

// Blacklisted : gtk_font_button_set_font_name

// Blacklisted : gtk_font_button_set_show_size

// Blacklisted : gtk_font_button_set_show_style

// SetTitle is a wrapper around the C function gtk_font_button_set_title.
func (recv *FontButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_font_button_set_title((*C.GtkFontButton)(recv.native), c_title)

	return
}

// Blacklisted : gtk_font_button_set_use_font

// Blacklisted : gtk_font_button_set_use_size

// Blacklisted : gtk_icon_info_copy

// Blacklisted : gtk_icon_info_free

// Unsupported : gtk_icon_info_get_attach_points : unsupported parameter points : output array param points

// Blacklisted : gtk_icon_info_get_base_size

// Blacklisted : gtk_icon_info_get_builtin_pixbuf

// Blacklisted : gtk_icon_info_get_display_name

// Blacklisted : gtk_icon_info_get_embedded_rect

// Blacklisted : gtk_icon_info_get_filename

// Blacklisted : gtk_icon_info_load_icon

// Blacklisted : gtk_icon_info_set_raw_coordinates

// Blacklisted : gtk_icon_theme_new

// Blacklisted : gtk_icon_theme_add_builtin_icon

// Blacklisted : gtk_icon_theme_get_default

// Blacklisted : gtk_icon_theme_get_for_screen

// Blacklisted : gtk_icon_theme_append_search_path

// Blacklisted : gtk_icon_theme_get_example_icon_name

// Unsupported : gtk_icon_theme_get_search_path : unsupported parameter path : output array param path

// Blacklisted : gtk_icon_theme_has_icon

// Blacklisted : gtk_icon_theme_list_icons

// Blacklisted : gtk_icon_theme_load_icon

// Blacklisted : gtk_icon_theme_lookup_icon

// Blacklisted : gtk_icon_theme_prepend_search_path

// Blacklisted : gtk_icon_theme_rescan_if_needed

// Blacklisted : gtk_icon_theme_set_custom_theme

// Blacklisted : gtk_icon_theme_set_screen

// Blacklisted : gtk_icon_theme_set_search_path

// Blacklisted : gtk_menu_attach

// Blacklisted : gtk_menu_set_monitor

// Blacklisted : gtk_menu_shell_cancel

// Blacklisted : gtk_message_dialog_new_with_markup

// Blacklisted : gtk_message_dialog_set_markup

// Blacklisted : gtk_paned_get_child1

// Blacklisted : gtk_paned_get_child2

type signalRadioActionChangedDetail struct {
	callback  RadioActionSignalChangedCallback
	handlerID C.gulong
}

var signalRadioActionChangedId int
var signalRadioActionChangedMap = make(map[int]signalRadioActionChangedDetail)
var signalRadioActionChangedLock sync.RWMutex

// RadioActionSignalChangedCallback is a callback function for a 'changed' signal emitted from a RadioAction.
type RadioActionSignalChangedCallback func(current *RadioAction)

/*
ConnectChanged connects the callback to the 'changed' signal for the RadioAction.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *RadioAction) ConnectChanged(callback RadioActionSignalChangedCallback) int {
	signalRadioActionChangedLock.Lock()
	defer signalRadioActionChangedLock.Unlock()

	signalRadioActionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.RadioAction_signal_connect_changed(instance, C.gpointer(uintptr(signalRadioActionChangedId)))

	detail := signalRadioActionChangedDetail{callback, handlerID}
	signalRadioActionChangedMap[signalRadioActionChangedId] = detail

	return signalRadioActionChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the RadioAction.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *RadioAction) DisconnectChanged(connectionID int) {
	signalRadioActionChangedLock.Lock()
	defer signalRadioActionChangedLock.Unlock()

	detail, exists := signalRadioActionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalRadioActionChangedMap, connectionID)
}

//export radioaction_changedHandler
func radioaction_changedHandler(_ *C.GObject, c_current *C.GtkRadioAction, data C.gpointer) {
	signalRadioActionChangedLock.RLock()
	defer signalRadioActionChangedLock.RUnlock()

	current := RadioActionNewFromC(unsafe.Pointer(c_current))

	index := int(uintptr(data))
	callback := signalRadioActionChangedMap[index].callback
	callback(current)
}

// Blacklisted : gtk_radio_action_new

// Blacklisted : gtk_radio_action_get_current_value

// Blacklisted : gtk_radio_action_get_group

// Blacklisted : gtk_radio_action_set_group

type signalRadioButtonGroupChangedDetail struct {
	callback  RadioButtonSignalGroupChangedCallback
	handlerID C.gulong
}

var signalRadioButtonGroupChangedId int
var signalRadioButtonGroupChangedMap = make(map[int]signalRadioButtonGroupChangedDetail)
var signalRadioButtonGroupChangedLock sync.RWMutex

// RadioButtonSignalGroupChangedCallback is a callback function for a 'group-changed' signal emitted from a RadioButton.
type RadioButtonSignalGroupChangedCallback func()

/*
ConnectGroupChanged connects the callback to the 'group-changed' signal for the RadioButton.

The returned value represents the connection, and may be passed to DisconnectGroupChanged to remove it.
*/
func (recv *RadioButton) ConnectGroupChanged(callback RadioButtonSignalGroupChangedCallback) int {
	signalRadioButtonGroupChangedLock.Lock()
	defer signalRadioButtonGroupChangedLock.Unlock()

	signalRadioButtonGroupChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.RadioButton_signal_connect_group_changed(instance, C.gpointer(uintptr(signalRadioButtonGroupChangedId)))

	detail := signalRadioButtonGroupChangedDetail{callback, handlerID}
	signalRadioButtonGroupChangedMap[signalRadioButtonGroupChangedId] = detail

	return signalRadioButtonGroupChangedId
}

/*
DisconnectGroupChanged disconnects a callback from the 'group-changed' signal for the RadioButton.

The connectionID should be a value returned from a call to ConnectGroupChanged.
*/
func (recv *RadioButton) DisconnectGroupChanged(connectionID int) {
	signalRadioButtonGroupChangedLock.Lock()
	defer signalRadioButtonGroupChangedLock.Unlock()

	detail, exists := signalRadioButtonGroupChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalRadioButtonGroupChangedMap, connectionID)
}

//export radiobutton_groupChangedHandler
func radiobutton_groupChangedHandler(_ *C.GObject, data C.gpointer) {
	signalRadioButtonGroupChangedLock.RLock()
	defer signalRadioButtonGroupChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalRadioButtonGroupChangedMap[index].callback
	callback()
}

// Blacklisted : gtk_radio_menu_item_new_from_widget

// Blacklisted : gtk_radio_menu_item_new_with_label_from_widget

// Blacklisted : gtk_radio_menu_item_new_with_mnemonic_from_widget

// Blacklisted : gtk_radio_tool_button_new

// Blacklisted : gtk_radio_tool_button_new_from_stock

// Blacklisted : gtk_radio_tool_button_new_from_widget

// Blacklisted : gtk_radio_tool_button_new_with_stock_from_widget

// Blacklisted : gtk_radio_tool_button_get_group

// Blacklisted : gtk_radio_tool_button_set_group

// Blacklisted : gtk_scale_get_layout

// Blacklisted : gtk_scale_get_layout_offsets

// Blacklisted : gtk_separator_tool_item_new

// Blacklisted : gtk_separator_tool_item_get_draw

// Blacklisted : gtk_separator_tool_item_set_draw

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

// Blacklisted : gtk_text_buffer_select_range

// Blacklisted : gtk_text_view_get_accepts_tab

// Blacklisted : gtk_text_view_get_overwrite

// Blacklisted : gtk_text_view_set_accepts_tab

// Blacklisted : gtk_text_view_set_overwrite

// Blacklisted : gtk_toggle_action_new

// Blacklisted : gtk_toggle_action_get_active

// Blacklisted : gtk_toggle_action_get_draw_as_radio

// Blacklisted : gtk_toggle_action_set_active

// Blacklisted : gtk_toggle_action_set_draw_as_radio

// Blacklisted : gtk_toggle_action_toggled

// Blacklisted : gtk_toggle_tool_button_new

// Blacklisted : gtk_toggle_tool_button_new_from_stock

// Blacklisted : gtk_toggle_tool_button_get_active

// Blacklisted : gtk_toggle_tool_button_set_active

// Blacklisted : gtk_tool_button_new

// Blacklisted : gtk_tool_button_new_from_stock

// Blacklisted : gtk_tool_button_get_icon_widget

// Blacklisted : gtk_tool_button_get_label

// Blacklisted : gtk_tool_button_get_label_widget

// Blacklisted : gtk_tool_button_get_stock_id

// Blacklisted : gtk_tool_button_get_use_underline

// Blacklisted : gtk_tool_button_set_icon_widget

// Blacklisted : gtk_tool_button_set_label

// Blacklisted : gtk_tool_button_set_label_widget

// Blacklisted : gtk_tool_button_set_stock_id

// Blacklisted : gtk_tool_button_set_use_underline

// Blacklisted : gtk_tool_item_new

// Blacklisted : gtk_tool_item_get_expand

// Blacklisted : gtk_tool_item_get_homogeneous

// Blacklisted : gtk_tool_item_get_icon_size

// Blacklisted : gtk_tool_item_get_is_important

// Blacklisted : gtk_tool_item_get_orientation

// Blacklisted : gtk_tool_item_get_proxy_menu_item

// Blacklisted : gtk_tool_item_get_relief_style

// Blacklisted : gtk_tool_item_get_toolbar_style

// Blacklisted : gtk_tool_item_get_use_drag_window

// Blacklisted : gtk_tool_item_get_visible_horizontal

// Blacklisted : gtk_tool_item_get_visible_vertical

// Blacklisted : gtk_tool_item_retrieve_proxy_menu_item

// Blacklisted : gtk_tool_item_set_expand

// Blacklisted : gtk_tool_item_set_homogeneous

// Blacklisted : gtk_tool_item_set_is_important

// Blacklisted : gtk_tool_item_set_proxy_menu_item

// Blacklisted : gtk_tool_item_set_use_drag_window

// Blacklisted : gtk_tool_item_set_visible_horizontal

// Blacklisted : gtk_tool_item_set_visible_vertical

// Blacklisted : gtk_toolbar_get_drop_index

// Blacklisted : gtk_toolbar_get_item_index

// Blacklisted : gtk_toolbar_get_n_items

// Blacklisted : gtk_toolbar_get_nth_item

// Blacklisted : gtk_toolbar_get_relief_style

// Blacklisted : gtk_toolbar_get_show_arrow

// Blacklisted : gtk_toolbar_insert

// Blacklisted : gtk_toolbar_set_drop_highlight_item

// Blacklisted : gtk_toolbar_set_show_arrow

// Blacklisted : gtk_tree_model_filter_clear_cache

// Blacklisted : gtk_tree_model_filter_convert_child_iter_to_iter

// Blacklisted : gtk_tree_model_filter_convert_child_path_to_path

// Blacklisted : gtk_tree_model_filter_convert_iter_to_child_iter

// Blacklisted : gtk_tree_model_filter_convert_path_to_child_path

// Blacklisted : gtk_tree_model_filter_get_model

// Blacklisted : gtk_tree_model_filter_refilter

// Unsupported : gtk_tree_model_filter_set_modify_func : unsupported parameter func : no type generator for TreeModelFilterModifyFunc (GtkTreeModelFilterModifyFunc) for param func

// Blacklisted : gtk_tree_model_filter_set_visible_column

// Unsupported : gtk_tree_model_filter_set_visible_func : unsupported parameter func : no type generator for TreeModelFilterVisibleFunc (GtkTreeModelFilterVisibleFunc) for param func

// Blacklisted : gtk_tree_view_column_get_expand

// Blacklisted : gtk_tree_view_column_set_expand

type signalUIManagerActionsChangedDetail struct {
	callback  UIManagerSignalActionsChangedCallback
	handlerID C.gulong
}

var signalUIManagerActionsChangedId int
var signalUIManagerActionsChangedMap = make(map[int]signalUIManagerActionsChangedDetail)
var signalUIManagerActionsChangedLock sync.RWMutex

// UIManagerSignalActionsChangedCallback is a callback function for a 'actions-changed' signal emitted from a UIManager.
type UIManagerSignalActionsChangedCallback func()

/*
ConnectActionsChanged connects the callback to the 'actions-changed' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectActionsChanged to remove it.
*/
func (recv *UIManager) ConnectActionsChanged(callback UIManagerSignalActionsChangedCallback) int {
	signalUIManagerActionsChangedLock.Lock()
	defer signalUIManagerActionsChangedLock.Unlock()

	signalUIManagerActionsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_actions_changed(instance, C.gpointer(uintptr(signalUIManagerActionsChangedId)))

	detail := signalUIManagerActionsChangedDetail{callback, handlerID}
	signalUIManagerActionsChangedMap[signalUIManagerActionsChangedId] = detail

	return signalUIManagerActionsChangedId
}

/*
DisconnectActionsChanged disconnects a callback from the 'actions-changed' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectActionsChanged.
*/
func (recv *UIManager) DisconnectActionsChanged(connectionID int) {
	signalUIManagerActionsChangedLock.Lock()
	defer signalUIManagerActionsChangedLock.Unlock()

	detail, exists := signalUIManagerActionsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerActionsChangedMap, connectionID)
}

//export uimanager_actionsChangedHandler
func uimanager_actionsChangedHandler(_ *C.GObject, data C.gpointer) {
	signalUIManagerActionsChangedLock.RLock()
	defer signalUIManagerActionsChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalUIManagerActionsChangedMap[index].callback
	callback()
}

type signalUIManagerAddWidgetDetail struct {
	callback  UIManagerSignalAddWidgetCallback
	handlerID C.gulong
}

var signalUIManagerAddWidgetId int
var signalUIManagerAddWidgetMap = make(map[int]signalUIManagerAddWidgetDetail)
var signalUIManagerAddWidgetLock sync.RWMutex

// UIManagerSignalAddWidgetCallback is a callback function for a 'add-widget' signal emitted from a UIManager.
type UIManagerSignalAddWidgetCallback func(widget *Widget)

/*
ConnectAddWidget connects the callback to the 'add-widget' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectAddWidget to remove it.
*/
func (recv *UIManager) ConnectAddWidget(callback UIManagerSignalAddWidgetCallback) int {
	signalUIManagerAddWidgetLock.Lock()
	defer signalUIManagerAddWidgetLock.Unlock()

	signalUIManagerAddWidgetId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_add_widget(instance, C.gpointer(uintptr(signalUIManagerAddWidgetId)))

	detail := signalUIManagerAddWidgetDetail{callback, handlerID}
	signalUIManagerAddWidgetMap[signalUIManagerAddWidgetId] = detail

	return signalUIManagerAddWidgetId
}

/*
DisconnectAddWidget disconnects a callback from the 'add-widget' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectAddWidget.
*/
func (recv *UIManager) DisconnectAddWidget(connectionID int) {
	signalUIManagerAddWidgetLock.Lock()
	defer signalUIManagerAddWidgetLock.Unlock()

	detail, exists := signalUIManagerAddWidgetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerAddWidgetMap, connectionID)
}

//export uimanager_addWidgetHandler
func uimanager_addWidgetHandler(_ *C.GObject, c_widget *C.GtkWidget, data C.gpointer) {
	signalUIManagerAddWidgetLock.RLock()
	defer signalUIManagerAddWidgetLock.RUnlock()

	widget := WidgetNewFromC(unsafe.Pointer(c_widget))

	index := int(uintptr(data))
	callback := signalUIManagerAddWidgetMap[index].callback
	callback(widget)
}

type signalUIManagerConnectProxyDetail struct {
	callback  UIManagerSignalConnectProxyCallback
	handlerID C.gulong
}

var signalUIManagerConnectProxyId int
var signalUIManagerConnectProxyMap = make(map[int]signalUIManagerConnectProxyDetail)
var signalUIManagerConnectProxyLock sync.RWMutex

// UIManagerSignalConnectProxyCallback is a callback function for a 'connect-proxy' signal emitted from a UIManager.
type UIManagerSignalConnectProxyCallback func(action *Action, proxy *Widget)

/*
ConnectConnectProxy connects the callback to the 'connect-proxy' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectConnectProxy to remove it.
*/
func (recv *UIManager) ConnectConnectProxy(callback UIManagerSignalConnectProxyCallback) int {
	signalUIManagerConnectProxyLock.Lock()
	defer signalUIManagerConnectProxyLock.Unlock()

	signalUIManagerConnectProxyId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_connect_proxy(instance, C.gpointer(uintptr(signalUIManagerConnectProxyId)))

	detail := signalUIManagerConnectProxyDetail{callback, handlerID}
	signalUIManagerConnectProxyMap[signalUIManagerConnectProxyId] = detail

	return signalUIManagerConnectProxyId
}

/*
DisconnectConnectProxy disconnects a callback from the 'connect-proxy' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectConnectProxy.
*/
func (recv *UIManager) DisconnectConnectProxy(connectionID int) {
	signalUIManagerConnectProxyLock.Lock()
	defer signalUIManagerConnectProxyLock.Unlock()

	detail, exists := signalUIManagerConnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerConnectProxyMap, connectionID)
}

//export uimanager_connectProxyHandler
func uimanager_connectProxyHandler(_ *C.GObject, c_action *C.GtkAction, c_proxy *C.GtkWidget, data C.gpointer) {
	signalUIManagerConnectProxyLock.RLock()
	defer signalUIManagerConnectProxyLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

	index := int(uintptr(data))
	callback := signalUIManagerConnectProxyMap[index].callback
	callback(action, proxy)
}

type signalUIManagerDisconnectProxyDetail struct {
	callback  UIManagerSignalDisconnectProxyCallback
	handlerID C.gulong
}

var signalUIManagerDisconnectProxyId int
var signalUIManagerDisconnectProxyMap = make(map[int]signalUIManagerDisconnectProxyDetail)
var signalUIManagerDisconnectProxyLock sync.RWMutex

// UIManagerSignalDisconnectProxyCallback is a callback function for a 'disconnect-proxy' signal emitted from a UIManager.
type UIManagerSignalDisconnectProxyCallback func(action *Action, proxy *Widget)

/*
ConnectDisconnectProxy connects the callback to the 'disconnect-proxy' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectDisconnectProxy to remove it.
*/
func (recv *UIManager) ConnectDisconnectProxy(callback UIManagerSignalDisconnectProxyCallback) int {
	signalUIManagerDisconnectProxyLock.Lock()
	defer signalUIManagerDisconnectProxyLock.Unlock()

	signalUIManagerDisconnectProxyId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_disconnect_proxy(instance, C.gpointer(uintptr(signalUIManagerDisconnectProxyId)))

	detail := signalUIManagerDisconnectProxyDetail{callback, handlerID}
	signalUIManagerDisconnectProxyMap[signalUIManagerDisconnectProxyId] = detail

	return signalUIManagerDisconnectProxyId
}

/*
DisconnectDisconnectProxy disconnects a callback from the 'disconnect-proxy' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectDisconnectProxy.
*/
func (recv *UIManager) DisconnectDisconnectProxy(connectionID int) {
	signalUIManagerDisconnectProxyLock.Lock()
	defer signalUIManagerDisconnectProxyLock.Unlock()

	detail, exists := signalUIManagerDisconnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerDisconnectProxyMap, connectionID)
}

//export uimanager_disconnectProxyHandler
func uimanager_disconnectProxyHandler(_ *C.GObject, c_action *C.GtkAction, c_proxy *C.GtkWidget, data C.gpointer) {
	signalUIManagerDisconnectProxyLock.RLock()
	defer signalUIManagerDisconnectProxyLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

	index := int(uintptr(data))
	callback := signalUIManagerDisconnectProxyMap[index].callback
	callback(action, proxy)
}

type signalUIManagerPostActivateDetail struct {
	callback  UIManagerSignalPostActivateCallback
	handlerID C.gulong
}

var signalUIManagerPostActivateId int
var signalUIManagerPostActivateMap = make(map[int]signalUIManagerPostActivateDetail)
var signalUIManagerPostActivateLock sync.RWMutex

// UIManagerSignalPostActivateCallback is a callback function for a 'post-activate' signal emitted from a UIManager.
type UIManagerSignalPostActivateCallback func(action *Action)

/*
ConnectPostActivate connects the callback to the 'post-activate' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectPostActivate to remove it.
*/
func (recv *UIManager) ConnectPostActivate(callback UIManagerSignalPostActivateCallback) int {
	signalUIManagerPostActivateLock.Lock()
	defer signalUIManagerPostActivateLock.Unlock()

	signalUIManagerPostActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_post_activate(instance, C.gpointer(uintptr(signalUIManagerPostActivateId)))

	detail := signalUIManagerPostActivateDetail{callback, handlerID}
	signalUIManagerPostActivateMap[signalUIManagerPostActivateId] = detail

	return signalUIManagerPostActivateId
}

/*
DisconnectPostActivate disconnects a callback from the 'post-activate' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectPostActivate.
*/
func (recv *UIManager) DisconnectPostActivate(connectionID int) {
	signalUIManagerPostActivateLock.Lock()
	defer signalUIManagerPostActivateLock.Unlock()

	detail, exists := signalUIManagerPostActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerPostActivateMap, connectionID)
}

//export uimanager_postActivateHandler
func uimanager_postActivateHandler(_ *C.GObject, c_action *C.GtkAction, data C.gpointer) {
	signalUIManagerPostActivateLock.RLock()
	defer signalUIManagerPostActivateLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	index := int(uintptr(data))
	callback := signalUIManagerPostActivateMap[index].callback
	callback(action)
}

type signalUIManagerPreActivateDetail struct {
	callback  UIManagerSignalPreActivateCallback
	handlerID C.gulong
}

var signalUIManagerPreActivateId int
var signalUIManagerPreActivateMap = make(map[int]signalUIManagerPreActivateDetail)
var signalUIManagerPreActivateLock sync.RWMutex

// UIManagerSignalPreActivateCallback is a callback function for a 'pre-activate' signal emitted from a UIManager.
type UIManagerSignalPreActivateCallback func(action *Action)

/*
ConnectPreActivate connects the callback to the 'pre-activate' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectPreActivate to remove it.
*/
func (recv *UIManager) ConnectPreActivate(callback UIManagerSignalPreActivateCallback) int {
	signalUIManagerPreActivateLock.Lock()
	defer signalUIManagerPreActivateLock.Unlock()

	signalUIManagerPreActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_pre_activate(instance, C.gpointer(uintptr(signalUIManagerPreActivateId)))

	detail := signalUIManagerPreActivateDetail{callback, handlerID}
	signalUIManagerPreActivateMap[signalUIManagerPreActivateId] = detail

	return signalUIManagerPreActivateId
}

/*
DisconnectPreActivate disconnects a callback from the 'pre-activate' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectPreActivate.
*/
func (recv *UIManager) DisconnectPreActivate(connectionID int) {
	signalUIManagerPreActivateLock.Lock()
	defer signalUIManagerPreActivateLock.Unlock()

	detail, exists := signalUIManagerPreActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerPreActivateMap, connectionID)
}

//export uimanager_preActivateHandler
func uimanager_preActivateHandler(_ *C.GObject, c_action *C.GtkAction, data C.gpointer) {
	signalUIManagerPreActivateLock.RLock()
	defer signalUIManagerPreActivateLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	index := int(uintptr(data))
	callback := signalUIManagerPreActivateMap[index].callback
	callback(action)
}

// Blacklisted : gtk_ui_manager_new

// Blacklisted : gtk_ui_manager_add_ui

// Blacklisted : gtk_ui_manager_add_ui_from_file

// Blacklisted : gtk_ui_manager_add_ui_from_string

// Blacklisted : gtk_ui_manager_ensure_update

// Blacklisted : gtk_ui_manager_get_accel_group

// Blacklisted : gtk_ui_manager_get_action

// Blacklisted : gtk_ui_manager_get_action_groups

// Blacklisted : gtk_ui_manager_get_add_tearoffs

// Blacklisted : gtk_ui_manager_get_toplevels

// Blacklisted : gtk_ui_manager_get_ui

// Blacklisted : gtk_ui_manager_get_widget

// Blacklisted : gtk_ui_manager_insert_action_group

// Blacklisted : gtk_ui_manager_new_merge_id

// Blacklisted : gtk_ui_manager_remove_action_group

// Blacklisted : gtk_ui_manager_remove_ui

// Blacklisted : gtk_ui_manager_set_add_tearoffs

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
