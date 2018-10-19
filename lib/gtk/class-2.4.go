// This is a generated file - DO NOT EDIT
// +build gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void Action_activateHandler();

	static gulong Action_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", Action_activateHandler, data);
	}

*/
/*

	void ActionGroup_connectProxyHandler();

	static gulong ActionGroup_signal_connect_connect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "connect-proxy", ActionGroup_connectProxyHandler, data);
	}

*/
/*

	void ActionGroup_disconnectProxyHandler();

	static gulong ActionGroup_signal_connect_disconnect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "disconnect-proxy", ActionGroup_disconnectProxyHandler, data);
	}

*/
/*

	void ActionGroup_postActivateHandler();

	static gulong ActionGroup_signal_connect_post_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "post-activate", ActionGroup_postActivateHandler, data);
	}

*/
/*

	void ActionGroup_preActivateHandler();

	static gulong ActionGroup_signal_connect_pre_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pre-activate", ActionGroup_preActivateHandler, data);
	}

*/
/*

	void CellRenderer_editingCanceledHandler();

	static gulong CellRenderer_signal_connect_editing_canceled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "editing-canceled", CellRenderer_editingCanceledHandler, data);
	}

*/
/*

	void ColorButton_colorSetHandler();

	static gulong ColorButton_signal_connect_color_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "color-set", ColorButton_colorSetHandler, data);
	}

*/
/*

	void ComboBox_changedHandler();

	static gulong ComboBox_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", ComboBox_changedHandler, data);
	}

*/
/*

	void FontButton_fontSetHandler();

	static gulong FontButton_signal_connect_font_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "font-set", FontButton_fontSetHandler, data);
	}

*/
/*

	void RadioAction_changedHandler();

	static gulong RadioAction_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", RadioAction_changedHandler, data);
	}

*/
/*

	void RadioButton_groupChangedHandler();

	static gulong RadioButton_signal_connect_group_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "group-changed", RadioButton_groupChangedHandler, data);
	}

*/
/*

	void Style_realizeHandler();

	static gulong Style_signal_connect_realize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "realize", Style_realizeHandler, data);
	}

*/
/*

	void Style_unrealizeHandler();

	static gulong Style_signal_connect_unrealize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unrealize", Style_unrealizeHandler, data);
	}

*/
/*

	void UIManager_actionsChangedHandler();

	static gulong UIManager_signal_connect_actions_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "actions-changed", UIManager_actionsChangedHandler, data);
	}

*/
/*

	void UIManager_addWidgetHandler();

	static gulong UIManager_signal_connect_add_widget(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "add-widget", UIManager_addWidgetHandler, data);
	}

*/
/*

	void UIManager_connectProxyHandler();

	static gulong UIManager_signal_connect_connect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "connect-proxy", UIManager_connectProxyHandler, data);
	}

*/
/*

	void UIManager_disconnectProxyHandler();

	static gulong UIManager_signal_connect_disconnect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "disconnect-proxy", UIManager_disconnectProxyHandler, data);
	}

*/
/*

	void UIManager_postActivateHandler();

	static gulong UIManager_signal_connect_post_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "post-activate", UIManager_postActivateHandler, data);
	}

*/
/*

	void UIManager_preActivateHandler();

	static gulong UIManager_signal_connect_pre_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pre-activate", UIManager_preActivateHandler, data);
	}

*/
import "C"

// Unsupported signal 'changed' for AccelMap : unsupported parameter accel_path : type utf8 :

var signalActionActivateId int
var signalActionActivateMap = make(map[int]ActionSignalActivateCallback)
var signalActionActivateLock sync.Mutex

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
	signalActionActivateMap[signalActionActivateId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Action_signal_connect_activate(instance, C.gpointer(uintptr(signalActionActivateId)))
	return int(retC)
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the Action.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *Action) DisconnectActivate(connectionID int) {
	signalActionActivateLock.Lock()
	defer signalActionActivateLock.Unlock()

	_, exists := signalActionActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalActionActivateMap, connectionID)
}

//export Action_activateHandler
func Action_activateHandler() {}

// ActionNew is a wrapper around the C function gtk_action_new.
func ActionNew(name string, label string, tooltip string, stockId string) *Action {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_action_new(c_name, c_label, c_tooltip, c_stock_id)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Activate is a wrapper around the C function gtk_action_activate.
func (recv *Action) Activate() {
	C.gtk_action_activate((*C.GtkAction)(recv.native))

	return
}

// ConnectAccelerator is a wrapper around the C function gtk_action_connect_accelerator.
func (recv *Action) ConnectAccelerator() {
	C.gtk_action_connect_accelerator((*C.GtkAction)(recv.native))

	return
}

// Unsupported : gtk_action_create_icon : unsupported parameter icon_size : no type generator for gint, GtkIconSize

// CreateMenuItem is a wrapper around the C function gtk_action_create_menu_item.
func (recv *Action) CreateMenuItem() *Widget {
	retC := C.gtk_action_create_menu_item((*C.GtkAction)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateToolItem is a wrapper around the C function gtk_action_create_tool_item.
func (recv *Action) CreateToolItem() *Widget {
	retC := C.gtk_action_create_tool_item((*C.GtkAction)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DisconnectAccelerator is a wrapper around the C function gtk_action_disconnect_accelerator.
func (recv *Action) DisconnectAccelerator() {
	C.gtk_action_disconnect_accelerator((*C.GtkAction)(recv.native))

	return
}

// GetName is a wrapper around the C function gtk_action_get_name.
func (recv *Action) GetName() string {
	retC := C.gtk_action_get_name((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetProxies is a wrapper around the C function gtk_action_get_proxies.
func (recv *Action) GetProxies() *glib.SList {
	retC := C.gtk_action_get_proxies((*C.GtkAction)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSensitive is a wrapper around the C function gtk_action_get_sensitive.
func (recv *Action) GetSensitive() bool {
	retC := C.gtk_action_get_sensitive((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisible is a wrapper around the C function gtk_action_get_visible.
func (recv *Action) GetVisible() bool {
	retC := C.gtk_action_get_visible((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsSensitive is a wrapper around the C function gtk_action_is_sensitive.
func (recv *Action) IsSensitive() bool {
	retC := C.gtk_action_is_sensitive((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsVisible is a wrapper around the C function gtk_action_is_visible.
func (recv *Action) IsVisible() bool {
	retC := C.gtk_action_is_visible((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAccelGroup is a wrapper around the C function gtk_action_set_accel_group.
func (recv *Action) SetAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	C.gtk_action_set_accel_group((*C.GtkAction)(recv.native), c_accel_group)

	return
}

// SetAccelPath is a wrapper around the C function gtk_action_set_accel_path.
func (recv *Action) SetAccelPath(accelPath string) {
	c_accel_path := C.CString(accelPath)
	defer C.free(unsafe.Pointer(c_accel_path))

	C.gtk_action_set_accel_path((*C.GtkAction)(recv.native), c_accel_path)

	return
}

var signalActionGroupConnectProxyId int
var signalActionGroupConnectProxyMap = make(map[int]ActionGroupSignalConnectProxyCallback)
var signalActionGroupConnectProxyLock sync.Mutex

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
	signalActionGroupConnectProxyMap[signalActionGroupConnectProxyId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ActionGroup_signal_connect_connect_proxy(instance, C.gpointer(uintptr(signalActionGroupConnectProxyId)))
	return int(retC)
}

/*
DisconnectConnectProxy disconnects a callback from the 'connect-proxy' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectConnectProxy.
*/
func (recv *ActionGroup) DisconnectConnectProxy(connectionID int) {
	signalActionGroupConnectProxyLock.Lock()
	defer signalActionGroupConnectProxyLock.Unlock()

	_, exists := signalActionGroupConnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalActionGroupConnectProxyMap, connectionID)
}

//export ActionGroup_connectProxyHandler
func ActionGroup_connectProxyHandler(c_action *C.GtkAction, c_proxy *C.GtkWidget) {
	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

}

var signalActionGroupDisconnectProxyId int
var signalActionGroupDisconnectProxyMap = make(map[int]ActionGroupSignalDisconnectProxyCallback)
var signalActionGroupDisconnectProxyLock sync.Mutex

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
	signalActionGroupDisconnectProxyMap[signalActionGroupDisconnectProxyId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ActionGroup_signal_connect_disconnect_proxy(instance, C.gpointer(uintptr(signalActionGroupDisconnectProxyId)))
	return int(retC)
}

/*
DisconnectDisconnectProxy disconnects a callback from the 'disconnect-proxy' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectDisconnectProxy.
*/
func (recv *ActionGroup) DisconnectDisconnectProxy(connectionID int) {
	signalActionGroupDisconnectProxyLock.Lock()
	defer signalActionGroupDisconnectProxyLock.Unlock()

	_, exists := signalActionGroupDisconnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalActionGroupDisconnectProxyMap, connectionID)
}

//export ActionGroup_disconnectProxyHandler
func ActionGroup_disconnectProxyHandler(c_action *C.GtkAction, c_proxy *C.GtkWidget) {
	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

}

var signalActionGroupPostActivateId int
var signalActionGroupPostActivateMap = make(map[int]ActionGroupSignalPostActivateCallback)
var signalActionGroupPostActivateLock sync.Mutex

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
	signalActionGroupPostActivateMap[signalActionGroupPostActivateId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ActionGroup_signal_connect_post_activate(instance, C.gpointer(uintptr(signalActionGroupPostActivateId)))
	return int(retC)
}

/*
DisconnectPostActivate disconnects a callback from the 'post-activate' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectPostActivate.
*/
func (recv *ActionGroup) DisconnectPostActivate(connectionID int) {
	signalActionGroupPostActivateLock.Lock()
	defer signalActionGroupPostActivateLock.Unlock()

	_, exists := signalActionGroupPostActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalActionGroupPostActivateMap, connectionID)
}

//export ActionGroup_postActivateHandler
func ActionGroup_postActivateHandler(c_action *C.GtkAction) {
	action := ActionNewFromC(unsafe.Pointer(c_action))

}

var signalActionGroupPreActivateId int
var signalActionGroupPreActivateMap = make(map[int]ActionGroupSignalPreActivateCallback)
var signalActionGroupPreActivateLock sync.Mutex

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
	signalActionGroupPreActivateMap[signalActionGroupPreActivateId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ActionGroup_signal_connect_pre_activate(instance, C.gpointer(uintptr(signalActionGroupPreActivateId)))
	return int(retC)
}

/*
DisconnectPreActivate disconnects a callback from the 'pre-activate' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectPreActivate.
*/
func (recv *ActionGroup) DisconnectPreActivate(connectionID int) {
	signalActionGroupPreActivateLock.Lock()
	defer signalActionGroupPreActivateLock.Unlock()

	_, exists := signalActionGroupPreActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalActionGroupPreActivateMap, connectionID)
}

//export ActionGroup_preActivateHandler
func ActionGroup_preActivateHandler(c_action *C.GtkAction) {
	action := ActionNewFromC(unsafe.Pointer(c_action))

}

// ActionGroupNew is a wrapper around the C function gtk_action_group_new.
func ActionGroupNew(name string) *ActionGroup {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_action_group_new(c_name)
	retGo := ActionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddAction is a wrapper around the C function gtk_action_group_add_action.
func (recv *ActionGroup) AddAction(action *Action) {
	c_action := (*C.GtkAction)(action.ToC())

	C.gtk_action_group_add_action((*C.GtkActionGroup)(recv.native), c_action)

	return
}

// AddActionWithAccel is a wrapper around the C function gtk_action_group_add_action_with_accel.
func (recv *ActionGroup) AddActionWithAccel(action *Action, accelerator string) {
	c_action := (*C.GtkAction)(action.ToC())

	c_accelerator := C.CString(accelerator)
	defer C.free(unsafe.Pointer(c_accelerator))

	C.gtk_action_group_add_action_with_accel((*C.GtkActionGroup)(recv.native), c_action, c_accelerator)

	return
}

// Unsupported : gtk_action_group_add_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_radio_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_radio_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_toggle_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_toggle_actions_full : unsupported parameter entries : no param type

// GetAction is a wrapper around the C function gtk_action_group_get_action.
func (recv *ActionGroup) GetAction(actionName string) *Action {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.gtk_action_group_get_action((*C.GtkActionGroup)(recv.native), c_action_name)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function gtk_action_group_get_name.
func (recv *ActionGroup) GetName() string {
	retC := C.gtk_action_group_get_name((*C.GtkActionGroup)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSensitive is a wrapper around the C function gtk_action_group_get_sensitive.
func (recv *ActionGroup) GetSensitive() bool {
	retC := C.gtk_action_group_get_sensitive((*C.GtkActionGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisible is a wrapper around the C function gtk_action_group_get_visible.
func (recv *ActionGroup) GetVisible() bool {
	retC := C.gtk_action_group_get_visible((*C.GtkActionGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ListActions is a wrapper around the C function gtk_action_group_list_actions.
func (recv *ActionGroup) ListActions() *glib.List {
	retC := C.gtk_action_group_list_actions((*C.GtkActionGroup)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveAction is a wrapper around the C function gtk_action_group_remove_action.
func (recv *ActionGroup) RemoveAction(action *Action) {
	c_action := (*C.GtkAction)(action.ToC())

	C.gtk_action_group_remove_action((*C.GtkActionGroup)(recv.native), c_action)

	return
}

// SetSensitive is a wrapper around the C function gtk_action_group_set_sensitive.
func (recv *ActionGroup) SetSensitive(sensitive bool) {
	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_action_group_set_sensitive((*C.GtkActionGroup)(recv.native), c_sensitive)

	return
}

// Unsupported : gtk_action_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GtkTranslateFunc

// SetTranslationDomain is a wrapper around the C function gtk_action_group_set_translation_domain.
func (recv *ActionGroup) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.gtk_action_group_set_translation_domain((*C.GtkActionGroup)(recv.native), c_domain)

	return
}

// SetVisible is a wrapper around the C function gtk_action_group_set_visible.
func (recv *ActionGroup) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_action_group_set_visible((*C.GtkActionGroup)(recv.native), c_visible)

	return
}

// GetPadding is a wrapper around the C function gtk_alignment_get_padding.
func (recv *Alignment) GetPadding() (uint32, uint32, uint32, uint32) {
	var c_padding_top C.guint

	var c_padding_bottom C.guint

	var c_padding_left C.guint

	var c_padding_right C.guint

	C.gtk_alignment_get_padding((*C.GtkAlignment)(recv.native), &c_padding_top, &c_padding_bottom, &c_padding_left, &c_padding_right)

	paddingTop := (uint32)(c_padding_top)

	paddingBottom := (uint32)(c_padding_bottom)

	paddingLeft := (uint32)(c_padding_left)

	paddingRight := (uint32)(c_padding_right)

	return paddingTop, paddingBottom, paddingLeft, paddingRight
}

// SetPadding is a wrapper around the C function gtk_alignment_set_padding.
func (recv *Alignment) SetPadding(paddingTop uint32, paddingBottom uint32, paddingLeft uint32, paddingRight uint32) {
	c_padding_top := (C.guint)(paddingTop)

	c_padding_bottom := (C.guint)(paddingBottom)

	c_padding_left := (C.guint)(paddingLeft)

	c_padding_right := (C.guint)(paddingRight)

	C.gtk_alignment_set_padding((*C.GtkAlignment)(recv.native), c_padding_top, c_padding_bottom, c_padding_left, c_padding_right)

	return
}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// GetAlignment is a wrapper around the C function gtk_button_get_alignment.
func (recv *Button) GetAlignment() (float32, float32) {
	var c_xalign C.gfloat

	var c_yalign C.gfloat

	C.gtk_button_get_alignment((*C.GtkButton)(recv.native), &c_xalign, &c_yalign)

	xalign := (float32)(c_xalign)

	yalign := (float32)(c_yalign)

	return xalign, yalign
}

// GetFocusOnClick is a wrapper around the C function gtk_button_get_focus_on_click.
func (recv *Button) GetFocusOnClick() bool {
	retC := C.gtk_button_get_focus_on_click((*C.GtkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAlignment is a wrapper around the C function gtk_button_set_alignment.
func (recv *Button) SetAlignment(xalign float32, yalign float32) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_button_set_alignment((*C.GtkButton)(recv.native), c_xalign, c_yalign)

	return
}

// SetFocusOnClick is a wrapper around the C function gtk_button_set_focus_on_click.
func (recv *Button) SetFocusOnClick(focusOnClick bool) {
	c_focus_on_click :=
		boolToGboolean(focusOnClick)

	C.gtk_button_set_focus_on_click((*C.GtkButton)(recv.native), c_focus_on_click)

	return
}

// GetChildSecondary is a wrapper around the C function gtk_button_box_get_child_secondary.
func (recv *ButtonBox) GetChildSecondary(child *Widget) bool {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_button_box_get_child_secondary((*C.GtkButtonBox)(recv.native), c_child)
	retGo := retC == C.TRUE

	return retGo
}

// GetDisplayOptions is a wrapper around the C function gtk_calendar_get_display_options.
func (recv *Calendar) GetDisplayOptions() CalendarDisplayOptions {
	retC := C.gtk_calendar_get_display_options((*C.GtkCalendar)(recv.native))
	retGo := (CalendarDisplayOptions)(retC)

	return retGo
}

// SetDisplayOptions is a wrapper around the C function gtk_calendar_set_display_options.
func (recv *Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	c_flags := (C.GtkCalendarDisplayOptions)(flags)

	C.gtk_calendar_set_display_options((*C.GtkCalendar)(recv.native), c_flags)

	return
}

var signalCellRendererEditingCanceledId int
var signalCellRendererEditingCanceledMap = make(map[int]CellRendererSignalEditingCanceledCallback)
var signalCellRendererEditingCanceledLock sync.Mutex

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
	signalCellRendererEditingCanceledMap[signalCellRendererEditingCanceledId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.CellRenderer_signal_connect_editing_canceled(instance, C.gpointer(uintptr(signalCellRendererEditingCanceledId)))
	return int(retC)
}

/*
DisconnectEditingCanceled disconnects a callback from the 'editing-canceled' signal for the CellRenderer.

The connectionID should be a value returned from a call to ConnectEditingCanceled.
*/
func (recv *CellRenderer) DisconnectEditingCanceled(connectionID int) {
	signalCellRendererEditingCanceledLock.Lock()
	defer signalCellRendererEditingCanceledLock.Unlock()

	_, exists := signalCellRendererEditingCanceledMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalCellRendererEditingCanceledMap, connectionID)
}

//export CellRenderer_editingCanceledHandler
func CellRenderer_editingCanceledHandler() {}

// GetDrawAsRadio is a wrapper around the C function gtk_check_menu_item_get_draw_as_radio.
func (recv *CheckMenuItem) GetDrawAsRadio() bool {
	retC := C.gtk_check_menu_item_get_draw_as_radio((*C.GtkCheckMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDrawAsRadio is a wrapper around the C function gtk_check_menu_item_set_draw_as_radio.
func (recv *CheckMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	c_draw_as_radio :=
		boolToGboolean(drawAsRadio)

	C.gtk_check_menu_item_set_draw_as_radio((*C.GtkCheckMenuItem)(recv.native), c_draw_as_radio)

	return
}

// Unsupported : gtk_clipboard_request_targets : unsupported parameter callback : no type generator for ClipboardTargetsReceivedFunc, GtkClipboardTargetsReceivedFunc

// Unsupported : gtk_clipboard_wait_for_targets : unsupported parameter targets : no param type

var signalColorButtonColorSetId int
var signalColorButtonColorSetMap = make(map[int]ColorButtonSignalColorSetCallback)
var signalColorButtonColorSetLock sync.Mutex

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
	signalColorButtonColorSetMap[signalColorButtonColorSetId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ColorButton_signal_connect_color_set(instance, C.gpointer(uintptr(signalColorButtonColorSetId)))
	return int(retC)
}

/*
DisconnectColorSet disconnects a callback from the 'color-set' signal for the ColorButton.

The connectionID should be a value returned from a call to ConnectColorSet.
*/
func (recv *ColorButton) DisconnectColorSet(connectionID int) {
	signalColorButtonColorSetLock.Lock()
	defer signalColorButtonColorSetLock.Unlock()

	_, exists := signalColorButtonColorSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalColorButtonColorSetMap, connectionID)
}

//export ColorButton_colorSetHandler
func ColorButton_colorSetHandler() {}

// ColorButtonNew is a wrapper around the C function gtk_color_button_new.
func ColorButtonNew() *ColorButton {
	retC := C.gtk_color_button_new()
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ColorButtonNewWithColor is a wrapper around the C function gtk_color_button_new_with_color.
func ColorButtonNewWithColor(color *gdk.Color) *ColorButton {
	c_color := (*C.GdkColor)(color.ToC())

	retC := C.gtk_color_button_new_with_color(c_color)
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAlpha is a wrapper around the C function gtk_color_button_get_alpha.
func (recv *ColorButton) GetAlpha() uint16 {
	retC := C.gtk_color_button_get_alpha((*C.GtkColorButton)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetColor is a wrapper around the C function gtk_color_button_get_color.
func (recv *ColorButton) GetColor() *gdk.Color {
	var c_color C.GdkColor

	C.gtk_color_button_get_color((*C.GtkColorButton)(recv.native), &c_color)

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetTitle is a wrapper around the C function gtk_color_button_get_title.
func (recv *ColorButton) GetTitle() string {
	retC := C.gtk_color_button_get_title((*C.GtkColorButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUseAlpha is a wrapper around the C function gtk_color_button_get_use_alpha.
func (recv *ColorButton) GetUseAlpha() bool {
	retC := C.gtk_color_button_get_use_alpha((*C.GtkColorButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAlpha is a wrapper around the C function gtk_color_button_set_alpha.
func (recv *ColorButton) SetAlpha(alpha uint16) {
	c_alpha := (C.guint16)(alpha)

	C.gtk_color_button_set_alpha((*C.GtkColorButton)(recv.native), c_alpha)

	return
}

// SetColor is a wrapper around the C function gtk_color_button_set_color.
func (recv *ColorButton) SetColor(color *gdk.Color) {
	c_color := (*C.GdkColor)(color.ToC())

	C.gtk_color_button_set_color((*C.GtkColorButton)(recv.native), c_color)

	return
}

// SetTitle is a wrapper around the C function gtk_color_button_set_title.
func (recv *ColorButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_color_button_set_title((*C.GtkColorButton)(recv.native), c_title)

	return
}

// SetUseAlpha is a wrapper around the C function gtk_color_button_set_use_alpha.
func (recv *ColorButton) SetUseAlpha(useAlpha bool) {
	c_use_alpha :=
		boolToGboolean(useAlpha)

	C.gtk_color_button_set_use_alpha((*C.GtkColorButton)(recv.native), c_use_alpha)

	return
}

var signalComboBoxChangedId int
var signalComboBoxChangedMap = make(map[int]ComboBoxSignalChangedCallback)
var signalComboBoxChangedLock sync.Mutex

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
	signalComboBoxChangedMap[signalComboBoxChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ComboBox_signal_connect_changed(instance, C.gpointer(uintptr(signalComboBoxChangedId)))
	return int(retC)
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *ComboBox) DisconnectChanged(connectionID int) {
	signalComboBoxChangedLock.Lock()
	defer signalComboBoxChangedLock.Unlock()

	_, exists := signalComboBoxChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalComboBoxChangedMap, connectionID)
}

//export ComboBox_changedHandler
func ComboBox_changedHandler() {}

// ComboBoxNew is a wrapper around the C function gtk_combo_box_new.
func ComboBoxNew() *ComboBox {
	retC := C.gtk_combo_box_new()
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// GetActive is a wrapper around the C function gtk_combo_box_get_active.
func (recv *ComboBox) GetActive() int32 {
	retC := C.gtk_combo_box_get_active((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetActiveIter is a wrapper around the C function gtk_combo_box_get_active_iter.
func (recv *ComboBox) GetActiveIter() (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	retC := C.gtk_combo_box_get_active_iter((*C.GtkComboBox)(recv.native), &c_iter)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// Unsupported : gtk_combo_box_get_model : no return generator

// Popdown is a wrapper around the C function gtk_combo_box_popdown.
func (recv *ComboBox) Popdown() {
	C.gtk_combo_box_popdown((*C.GtkComboBox)(recv.native))

	return
}

// Popup is a wrapper around the C function gtk_combo_box_popup.
func (recv *ComboBox) Popup() {
	C.gtk_combo_box_popup((*C.GtkComboBox)(recv.native))

	return
}

// SetActive is a wrapper around the C function gtk_combo_box_set_active.
func (recv *ComboBox) SetActive(index int32) {
	c_index_ := (C.gint)(index)

	C.gtk_combo_box_set_active((*C.GtkComboBox)(recv.native), c_index_)

	return
}

// SetActiveIter is a wrapper around the C function gtk_combo_box_set_active_iter.
func (recv *ComboBox) SetActiveIter(iter *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	C.gtk_combo_box_set_active_iter((*C.GtkComboBox)(recv.native), c_iter)

	return
}

// SetColumnSpanColumn is a wrapper around the C function gtk_combo_box_set_column_span_column.
func (recv *ComboBox) SetColumnSpanColumn(columnSpan int32) {
	c_column_span := (C.gint)(columnSpan)

	C.gtk_combo_box_set_column_span_column((*C.GtkComboBox)(recv.native), c_column_span)

	return
}

// Unsupported : gtk_combo_box_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// SetRowSpanColumn is a wrapper around the C function gtk_combo_box_set_row_span_column.
func (recv *ComboBox) SetRowSpanColumn(rowSpan int32) {
	c_row_span := (C.gint)(rowSpan)

	C.gtk_combo_box_set_row_span_column((*C.GtkComboBox)(recv.native), c_row_span)

	return
}

// SetWrapWidth is a wrapper around the C function gtk_combo_box_set_wrap_width.
func (recv *ComboBox) SetWrapWidth(width int32) {
	c_width := (C.gint)(width)

	C.gtk_combo_box_set_wrap_width((*C.GtkComboBox)(recv.native), c_width)

	return
}

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// GetAlignment is a wrapper around the C function gtk_entry_get_alignment.
func (recv *Entry) GetAlignment() float32 {
	retC := C.gtk_entry_get_alignment((*C.GtkEntry)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// GetCompletion is a wrapper around the C function gtk_entry_get_completion.
func (recv *Entry) GetCompletion() *EntryCompletion {
	retC := C.gtk_entry_get_completion((*C.GtkEntry)(recv.native))
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAlignment is a wrapper around the C function gtk_entry_set_alignment.
func (recv *Entry) SetAlignment(xalign float32) {
	c_xalign := (C.gfloat)(xalign)

	C.gtk_entry_set_alignment((*C.GtkEntry)(recv.native), c_xalign)

	return
}

// SetCompletion is a wrapper around the C function gtk_entry_set_completion.
func (recv *Entry) SetCompletion(completion *EntryCompletion) {
	c_completion := (*C.GtkEntryCompletion)(completion.ToC())

	C.gtk_entry_set_completion((*C.GtkEntry)(recv.native), c_completion)

	return
}

// Unsupported signal 'action-activated' for EntryCompletion : unsupported parameter index : type gint :

// Unsupported signal 'match-selected' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// EntryCompletionNew is a wrapper around the C function gtk_entry_completion_new.
func EntryCompletionNew() *EntryCompletion {
	retC := C.gtk_entry_completion_new()
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Complete is a wrapper around the C function gtk_entry_completion_complete.
func (recv *EntryCompletion) Complete() {
	C.gtk_entry_completion_complete((*C.GtkEntryCompletion)(recv.native))

	return
}

// DeleteAction is a wrapper around the C function gtk_entry_completion_delete_action.
func (recv *EntryCompletion) DeleteAction(index int32) {
	c_index_ := (C.gint)(index)

	C.gtk_entry_completion_delete_action((*C.GtkEntryCompletion)(recv.native), c_index_)

	return
}

// GetEntry is a wrapper around the C function gtk_entry_completion_get_entry.
func (recv *EntryCompletion) GetEntry() *Widget {
	retC := C.gtk_entry_completion_get_entry((*C.GtkEntryCompletion)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMinimumKeyLength is a wrapper around the C function gtk_entry_completion_get_minimum_key_length.
func (recv *EntryCompletion) GetMinimumKeyLength() int32 {
	retC := C.gtk_entry_completion_get_minimum_key_length((*C.GtkEntryCompletion)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_entry_completion_get_model : no return generator

// InsertActionMarkup is a wrapper around the C function gtk_entry_completion_insert_action_markup.
func (recv *EntryCompletion) InsertActionMarkup(index int32, markup string) {
	c_index_ := (C.gint)(index)

	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_entry_completion_insert_action_markup((*C.GtkEntryCompletion)(recv.native), c_index_, c_markup)

	return
}

// InsertActionText is a wrapper around the C function gtk_entry_completion_insert_action_text.
func (recv *EntryCompletion) InsertActionText(index int32, text string) {
	c_index_ := (C.gint)(index)

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_entry_completion_insert_action_text((*C.GtkEntryCompletion)(recv.native), c_index_, c_text)

	return
}

// Unsupported : gtk_entry_completion_set_match_func : unsupported parameter func : no type generator for EntryCompletionMatchFunc, GtkEntryCompletionMatchFunc

// SetMinimumKeyLength is a wrapper around the C function gtk_entry_completion_set_minimum_key_length.
func (recv *EntryCompletion) SetMinimumKeyLength(length int32) {
	c_length := (C.gint)(length)

	C.gtk_entry_completion_set_minimum_key_length((*C.GtkEntryCompletion)(recv.native), c_length)

	return
}

// Unsupported : gtk_entry_completion_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// SetTextColumn is a wrapper around the C function gtk_entry_completion_set_text_column.
func (recv *EntryCompletion) SetTextColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_entry_completion_set_text_column((*C.GtkEntryCompletion)(recv.native), c_column)

	return
}

// Unsupported : EntryIconAccessible : no CType

// GetAboveChild is a wrapper around the C function gtk_event_box_get_above_child.
func (recv *EventBox) GetAboveChild() bool {
	retC := C.gtk_event_box_get_above_child((*C.GtkEventBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisibleWindow is a wrapper around the C function gtk_event_box_get_visible_window.
func (recv *EventBox) GetVisibleWindow() bool {
	retC := C.gtk_event_box_get_visible_window((*C.GtkEventBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAboveChild is a wrapper around the C function gtk_event_box_set_above_child.
func (recv *EventBox) SetAboveChild(aboveChild bool) {
	c_above_child :=
		boolToGboolean(aboveChild)

	C.gtk_event_box_set_above_child((*C.GtkEventBox)(recv.native), c_above_child)

	return
}

// SetVisibleWindow is a wrapper around the C function gtk_event_box_set_visible_window.
func (recv *EventBox) SetVisibleWindow(visibleWindow bool) {
	c_visible_window :=
		boolToGboolean(visibleWindow)

	C.gtk_event_box_set_visible_window((*C.GtkEventBox)(recv.native), c_visible_window)

	return
}

// ExpanderNew is a wrapper around the C function gtk_expander_new.
func ExpanderNew(label string) *Expander {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_expander_new(c_label)
	retGo := ExpanderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ExpanderNewWithMnemonic is a wrapper around the C function gtk_expander_new_with_mnemonic.
func ExpanderNewWithMnemonic(label string) *Expander {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_expander_new_with_mnemonic(c_label)
	retGo := ExpanderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetExpanded is a wrapper around the C function gtk_expander_get_expanded.
func (recv *Expander) GetExpanded() bool {
	retC := C.gtk_expander_get_expanded((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetLabel is a wrapper around the C function gtk_expander_get_label.
func (recv *Expander) GetLabel() string {
	retC := C.gtk_expander_get_label((*C.GtkExpander)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLabelWidget is a wrapper around the C function gtk_expander_get_label_widget.
func (recv *Expander) GetLabelWidget() *Widget {
	retC := C.gtk_expander_get_label_widget((*C.GtkExpander)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSpacing is a wrapper around the C function gtk_expander_get_spacing.
func (recv *Expander) GetSpacing() int32 {
	retC := C.gtk_expander_get_spacing((*C.GtkExpander)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetUseMarkup is a wrapper around the C function gtk_expander_get_use_markup.
func (recv *Expander) GetUseMarkup() bool {
	retC := C.gtk_expander_get_use_markup((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseUnderline is a wrapper around the C function gtk_expander_get_use_underline.
func (recv *Expander) GetUseUnderline() bool {
	retC := C.gtk_expander_get_use_underline((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetExpanded is a wrapper around the C function gtk_expander_set_expanded.
func (recv *Expander) SetExpanded(expanded bool) {
	c_expanded :=
		boolToGboolean(expanded)

	C.gtk_expander_set_expanded((*C.GtkExpander)(recv.native), c_expanded)

	return
}

// SetLabel is a wrapper around the C function gtk_expander_set_label.
func (recv *Expander) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_expander_set_label((*C.GtkExpander)(recv.native), c_label)

	return
}

// SetLabelWidget is a wrapper around the C function gtk_expander_set_label_widget.
func (recv *Expander) SetLabelWidget(labelWidget *Widget) {
	c_label_widget := (*C.GtkWidget)(labelWidget.ToC())

	C.gtk_expander_set_label_widget((*C.GtkExpander)(recv.native), c_label_widget)

	return
}

// SetSpacing is a wrapper around the C function gtk_expander_set_spacing.
func (recv *Expander) SetSpacing(spacing int32) {
	c_spacing := (C.gint)(spacing)

	C.gtk_expander_set_spacing((*C.GtkExpander)(recv.native), c_spacing)

	return
}

// SetUseMarkup is a wrapper around the C function gtk_expander_set_use_markup.
func (recv *Expander) SetUseMarkup(useMarkup bool) {
	c_use_markup :=
		boolToGboolean(useMarkup)

	C.gtk_expander_set_use_markup((*C.GtkExpander)(recv.native), c_use_markup)

	return
}

// SetUseUnderline is a wrapper around the C function gtk_expander_set_use_underline.
func (recv *Expander) SetUseUnderline(useUnderline bool) {
	c_use_underline :=
		boolToGboolean(useUnderline)

	C.gtk_expander_set_use_underline((*C.GtkExpander)(recv.native), c_use_underline)

	return
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// FileChooserWidgetNew is a wrapper around the C function gtk_file_chooser_widget_new.
func FileChooserWidgetNew(action FileChooserAction) *FileChooserWidget {
	c_action := (C.GtkFileChooserAction)(action)

	retC := C.gtk_file_chooser_widget_new(c_action)
	retGo := FileChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileFilterNew is a wrapper around the C function gtk_file_filter_new.
func FileFilterNew() *FileFilter {
	retC := C.gtk_file_filter_new()
	retGo := FileFilterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_file_filter_add_custom : unsupported parameter func : no type generator for FileFilterFunc, GtkFileFilterFunc

// AddMimeType is a wrapper around the C function gtk_file_filter_add_mime_type.
func (recv *FileFilter) AddMimeType(mimeType string) {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	C.gtk_file_filter_add_mime_type((*C.GtkFileFilter)(recv.native), c_mime_type)

	return
}

// AddPattern is a wrapper around the C function gtk_file_filter_add_pattern.
func (recv *FileFilter) AddPattern(pattern string) {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	C.gtk_file_filter_add_pattern((*C.GtkFileFilter)(recv.native), c_pattern)

	return
}

// Filter is a wrapper around the C function gtk_file_filter_filter.
func (recv *FileFilter) Filter(filterInfo *FileFilterInfo) bool {
	c_filter_info := (*C.GtkFileFilterInfo)(filterInfo.ToC())

	retC := C.gtk_file_filter_filter((*C.GtkFileFilter)(recv.native), c_filter_info)
	retGo := retC == C.TRUE

	return retGo
}

// GetName is a wrapper around the C function gtk_file_filter_get_name.
func (recv *FileFilter) GetName() string {
	retC := C.gtk_file_filter_get_name((*C.GtkFileFilter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNeeded is a wrapper around the C function gtk_file_filter_get_needed.
func (recv *FileFilter) GetNeeded() FileFilterFlags {
	retC := C.gtk_file_filter_get_needed((*C.GtkFileFilter)(recv.native))
	retGo := (FileFilterFlags)(retC)

	return retGo
}

// SetName is a wrapper around the C function gtk_file_filter_set_name.
func (recv *FileFilter) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_file_filter_set_name((*C.GtkFileFilter)(recv.native), c_name)

	return
}

var signalFontButtonFontSetId int
var signalFontButtonFontSetMap = make(map[int]FontButtonSignalFontSetCallback)
var signalFontButtonFontSetLock sync.Mutex

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
	signalFontButtonFontSetMap[signalFontButtonFontSetId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.FontButton_signal_connect_font_set(instance, C.gpointer(uintptr(signalFontButtonFontSetId)))
	return int(retC)
}

/*
DisconnectFontSet disconnects a callback from the 'font-set' signal for the FontButton.

The connectionID should be a value returned from a call to ConnectFontSet.
*/
func (recv *FontButton) DisconnectFontSet(connectionID int) {
	signalFontButtonFontSetLock.Lock()
	defer signalFontButtonFontSetLock.Unlock()

	_, exists := signalFontButtonFontSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalFontButtonFontSetMap, connectionID)
}

//export FontButton_fontSetHandler
func FontButton_fontSetHandler() {}

// FontButtonNew is a wrapper around the C function gtk_font_button_new.
func FontButtonNew() *FontButton {
	retC := C.gtk_font_button_new()
	retGo := FontButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontButtonNewWithFont is a wrapper around the C function gtk_font_button_new_with_font.
func FontButtonNewWithFont(fontname string) *FontButton {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_button_new_with_font(c_fontname)
	retGo := FontButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFontName is a wrapper around the C function gtk_font_button_get_font_name.
func (recv *FontButton) GetFontName() string {
	retC := C.gtk_font_button_get_font_name((*C.GtkFontButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetShowSize is a wrapper around the C function gtk_font_button_get_show_size.
func (recv *FontButton) GetShowSize() bool {
	retC := C.gtk_font_button_get_show_size((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowStyle is a wrapper around the C function gtk_font_button_get_show_style.
func (recv *FontButton) GetShowStyle() bool {
	retC := C.gtk_font_button_get_show_style((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTitle is a wrapper around the C function gtk_font_button_get_title.
func (recv *FontButton) GetTitle() string {
	retC := C.gtk_font_button_get_title((*C.GtkFontButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUseFont is a wrapper around the C function gtk_font_button_get_use_font.
func (recv *FontButton) GetUseFont() bool {
	retC := C.gtk_font_button_get_use_font((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseSize is a wrapper around the C function gtk_font_button_get_use_size.
func (recv *FontButton) GetUseSize() bool {
	retC := C.gtk_font_button_get_use_size((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFontName is a wrapper around the C function gtk_font_button_set_font_name.
func (recv *FontButton) SetFontName(fontname string) bool {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_button_set_font_name((*C.GtkFontButton)(recv.native), c_fontname)
	retGo := retC == C.TRUE

	return retGo
}

// SetShowSize is a wrapper around the C function gtk_font_button_set_show_size.
func (recv *FontButton) SetShowSize(showSize bool) {
	c_show_size :=
		boolToGboolean(showSize)

	C.gtk_font_button_set_show_size((*C.GtkFontButton)(recv.native), c_show_size)

	return
}

// SetShowStyle is a wrapper around the C function gtk_font_button_set_show_style.
func (recv *FontButton) SetShowStyle(showStyle bool) {
	c_show_style :=
		boolToGboolean(showStyle)

	C.gtk_font_button_set_show_style((*C.GtkFontButton)(recv.native), c_show_style)

	return
}

// SetTitle is a wrapper around the C function gtk_font_button_set_title.
func (recv *FontButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_font_button_set_title((*C.GtkFontButton)(recv.native), c_title)

	return
}

// SetUseFont is a wrapper around the C function gtk_font_button_set_use_font.
func (recv *FontButton) SetUseFont(useFont bool) {
	c_use_font :=
		boolToGboolean(useFont)

	C.gtk_font_button_set_use_font((*C.GtkFontButton)(recv.native), c_use_font)

	return
}

// SetUseSize is a wrapper around the C function gtk_font_button_set_use_size.
func (recv *FontButton) SetUseSize(useSize bool) {
	c_use_size :=
		boolToGboolean(useSize)

	C.gtk_font_button_set_use_size((*C.GtkFontButton)(recv.native), c_use_size)

	return
}

// Copy is a wrapper around the C function gtk_icon_info_copy.
func (recv *IconInfo) Copy() *IconInfo {
	retC := C.gtk_icon_info_copy((*C.GtkIconInfo)(recv.native))
	retGo := IconInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_icon_info_free.
func (recv *IconInfo) Free() {
	C.gtk_icon_info_free((*C.GtkIconInfo)(recv.native))

	return
}

// Unsupported : gtk_icon_info_get_attach_points : unsupported parameter points : no param type

// GetBaseSize is a wrapper around the C function gtk_icon_info_get_base_size.
func (recv *IconInfo) GetBaseSize() int32 {
	retC := C.gtk_icon_info_get_base_size((*C.GtkIconInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetBuiltinPixbuf is a wrapper around the C function gtk_icon_info_get_builtin_pixbuf.
func (recv *IconInfo) GetBuiltinPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_icon_info_get_builtin_pixbuf((*C.GtkIconInfo)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplayName is a wrapper around the C function gtk_icon_info_get_display_name.
func (recv *IconInfo) GetDisplayName() string {
	retC := C.gtk_icon_info_get_display_name((*C.GtkIconInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_icon_info_get_embedded_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// GetFilename is a wrapper around the C function gtk_icon_info_get_filename.
func (recv *IconInfo) GetFilename() string {
	retC := C.gtk_icon_info_get_filename((*C.GtkIconInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// LoadIcon is a wrapper around the C function gtk_icon_info_load_icon.
func (recv *IconInfo) LoadIcon() (*gdkpixbuf.Pixbuf, error) {
	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_icon((*C.GtkIconInfo)(recv.native), &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetRawCoordinates is a wrapper around the C function gtk_icon_info_set_raw_coordinates.
func (recv *IconInfo) SetRawCoordinates(rawCoordinates bool) {
	c_raw_coordinates :=
		boolToGboolean(rawCoordinates)

	C.gtk_icon_info_set_raw_coordinates((*C.GtkIconInfo)(recv.native), c_raw_coordinates)

	return
}

// IconThemeNew is a wrapper around the C function gtk_icon_theme_new.
func IconThemeNew() *IconTheme {
	retC := C.gtk_icon_theme_new()
	retGo := IconThemeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendSearchPath is a wrapper around the C function gtk_icon_theme_append_search_path.
func (recv *IconTheme) AppendSearchPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.gtk_icon_theme_append_search_path((*C.GtkIconTheme)(recv.native), c_path)

	return
}

// GetExampleIconName is a wrapper around the C function gtk_icon_theme_get_example_icon_name.
func (recv *IconTheme) GetExampleIconName() string {
	retC := C.gtk_icon_theme_get_example_icon_name((*C.GtkIconTheme)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_theme_get_search_path : unsupported parameter path : no param type

// HasIcon is a wrapper around the C function gtk_icon_theme_has_icon.
func (recv *IconTheme) HasIcon(iconName string) bool {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	retC := C.gtk_icon_theme_has_icon((*C.GtkIconTheme)(recv.native), c_icon_name)
	retGo := retC == C.TRUE

	return retGo
}

// ListIcons is a wrapper around the C function gtk_icon_theme_list_icons.
func (recv *IconTheme) ListIcons(context string) *glib.List {
	c_context := C.CString(context)
	defer C.free(unsafe.Pointer(c_context))

	retC := C.gtk_icon_theme_list_icons((*C.GtkIconTheme)(recv.native), c_context)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LoadIcon is a wrapper around the C function gtk_icon_theme_load_icon.
func (recv *IconTheme) LoadIcon(iconName string, size int32, flags IconLookupFlags) (*gdkpixbuf.Pixbuf, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_icon((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_flags, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LookupIcon is a wrapper around the C function gtk_icon_theme_lookup_icon.
func (recv *IconTheme) LookupIcon(iconName string, size int32, flags IconLookupFlags) *IconInfo {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_icon((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_flags)
	retGo := IconInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PrependSearchPath is a wrapper around the C function gtk_icon_theme_prepend_search_path.
func (recv *IconTheme) PrependSearchPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.gtk_icon_theme_prepend_search_path((*C.GtkIconTheme)(recv.native), c_path)

	return
}

// RescanIfNeeded is a wrapper around the C function gtk_icon_theme_rescan_if_needed.
func (recv *IconTheme) RescanIfNeeded() bool {
	retC := C.gtk_icon_theme_rescan_if_needed((*C.GtkIconTheme)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetCustomTheme is a wrapper around the C function gtk_icon_theme_set_custom_theme.
func (recv *IconTheme) SetCustomTheme(themeName string) {
	c_theme_name := C.CString(themeName)
	defer C.free(unsafe.Pointer(c_theme_name))

	C.gtk_icon_theme_set_custom_theme((*C.GtkIconTheme)(recv.native), c_theme_name)

	return
}

// SetScreen is a wrapper around the C function gtk_icon_theme_set_screen.
func (recv *IconTheme) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gtk_icon_theme_set_screen((*C.GtkIconTheme)(recv.native), c_screen)

	return
}

// Unsupported : gtk_icon_theme_set_search_path : unsupported parameter path : no param type

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Attach is a wrapper around the C function gtk_menu_attach.
func (recv *Menu) Attach(child *Widget, leftAttach uint32, rightAttach uint32, topAttach uint32, bottomAttach uint32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_left_attach := (C.guint)(leftAttach)

	c_right_attach := (C.guint)(rightAttach)

	c_top_attach := (C.guint)(topAttach)

	c_bottom_attach := (C.guint)(bottomAttach)

	C.gtk_menu_attach((*C.GtkMenu)(recv.native), c_child, c_left_attach, c_right_attach, c_top_attach, c_bottom_attach)

	return
}

// SetMonitor is a wrapper around the C function gtk_menu_set_monitor.
func (recv *Menu) SetMonitor(monitorNum int32) {
	c_monitor_num := (C.gint)(monitorNum)

	C.gtk_menu_set_monitor((*C.GtkMenu)(recv.native), c_monitor_num)

	return
}

// Cancel is a wrapper around the C function gtk_menu_shell_cancel.
func (recv *MenuShell) Cancel() {
	C.gtk_menu_shell_cancel((*C.GtkMenuShell)(recv.native))

	return
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// SetMarkup is a wrapper around the C function gtk_message_dialog_set_markup.
func (recv *MessageDialog) SetMarkup(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_message_dialog_set_markup((*C.GtkMessageDialog)(recv.native), c_str)

	return
}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// GetChild1 is a wrapper around the C function gtk_paned_get_child1.
func (recv *Paned) GetChild1() *Widget {
	retC := C.gtk_paned_get_child1((*C.GtkPaned)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetChild2 is a wrapper around the C function gtk_paned_get_child2.
func (recv *Paned) GetChild2() *Widget {
	retC := C.gtk_paned_get_child2((*C.GtkPaned)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

var signalRadioActionChangedId int
var signalRadioActionChangedMap = make(map[int]RadioActionSignalChangedCallback)
var signalRadioActionChangedLock sync.Mutex

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
	signalRadioActionChangedMap[signalRadioActionChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.RadioAction_signal_connect_changed(instance, C.gpointer(uintptr(signalRadioActionChangedId)))
	return int(retC)
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the RadioAction.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *RadioAction) DisconnectChanged(connectionID int) {
	signalRadioActionChangedLock.Lock()
	defer signalRadioActionChangedLock.Unlock()

	_, exists := signalRadioActionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalRadioActionChangedMap, connectionID)
}

//export RadioAction_changedHandler
func RadioAction_changedHandler(c_current *C.GtkRadioAction) {
	current := RadioActionNewFromC(unsafe.Pointer(c_current))

}

// RadioActionNew is a wrapper around the C function gtk_radio_action_new.
func RadioActionNew(name string, label string, tooltip string, stockId string, value int32) *RadioAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_value := (C.gint)(value)

	retC := C.gtk_radio_action_new(c_name, c_label, c_tooltip, c_stock_id, c_value)
	retGo := RadioActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentValue is a wrapper around the C function gtk_radio_action_get_current_value.
func (recv *RadioAction) GetCurrentValue() int32 {
	retC := C.gtk_radio_action_get_current_value((*C.GtkRadioAction)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetGroup is a wrapper around the C function gtk_radio_action_get_group.
func (recv *RadioAction) GetGroup() *glib.SList {
	retC := C.gtk_radio_action_get_group((*C.GtkRadioAction)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetGroup is a wrapper around the C function gtk_radio_action_set_group.
func (recv *RadioAction) SetGroup(group *glib.SList) {
	c_group := (*C.GSList)(group.ToC())

	C.gtk_radio_action_set_group((*C.GtkRadioAction)(recv.native), c_group)

	return
}

var signalRadioButtonGroupChangedId int
var signalRadioButtonGroupChangedMap = make(map[int]RadioButtonSignalGroupChangedCallback)
var signalRadioButtonGroupChangedLock sync.Mutex

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
	signalRadioButtonGroupChangedMap[signalRadioButtonGroupChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.RadioButton_signal_connect_group_changed(instance, C.gpointer(uintptr(signalRadioButtonGroupChangedId)))
	return int(retC)
}

/*
DisconnectGroupChanged disconnects a callback from the 'group-changed' signal for the RadioButton.

The connectionID should be a value returned from a call to ConnectGroupChanged.
*/
func (recv *RadioButton) DisconnectGroupChanged(connectionID int) {
	signalRadioButtonGroupChangedLock.Lock()
	defer signalRadioButtonGroupChangedLock.Unlock()

	_, exists := signalRadioButtonGroupChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalRadioButtonGroupChangedMap, connectionID)
}

//export RadioButton_groupChangedHandler
func RadioButton_groupChangedHandler() {}

// RadioMenuItemNewFromWidget is a wrapper around the C function gtk_radio_menu_item_new_from_widget.
func RadioMenuItemNewFromWidget(group *RadioMenuItem) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(group.ToC())

	retC := C.gtk_radio_menu_item_new_from_widget(c_group)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithLabelFromWidget is a wrapper around the C function gtk_radio_menu_item_new_with_label_from_widget.
func RadioMenuItemNewWithLabelFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(group.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_label_from_widget(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithMnemonicFromWidget is a wrapper around the C function gtk_radio_menu_item_new_with_mnemonic_from_widget.
func RadioMenuItemNewWithMnemonicFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(group.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_mnemonic_from_widget(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNew is a wrapper around the C function gtk_radio_tool_button_new.
func RadioToolButtonNew(group *glib.SList) *RadioToolButton {
	c_group := (*C.GSList)(group.ToC())

	retC := C.gtk_radio_tool_button_new(c_group)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNewFromStock is a wrapper around the C function gtk_radio_tool_button_new_from_stock.
func RadioToolButtonNewFromStock(group *glib.SList, stockId string) *RadioToolButton {
	c_group := (*C.GSList)(group.ToC())

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_radio_tool_button_new_from_stock(c_group, c_stock_id)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNewFromWidget is a wrapper around the C function gtk_radio_tool_button_new_from_widget.
func RadioToolButtonNewFromWidget(group *RadioToolButton) *RadioToolButton {
	c_group := (*C.GtkRadioToolButton)(group.ToC())

	retC := C.gtk_radio_tool_button_new_from_widget(c_group)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNewWithStockFromWidget is a wrapper around the C function gtk_radio_tool_button_new_with_stock_from_widget.
func RadioToolButtonNewWithStockFromWidget(group *RadioToolButton, stockId string) *RadioToolButton {
	c_group := (*C.GtkRadioToolButton)(group.ToC())

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_radio_tool_button_new_with_stock_from_widget(c_group, c_stock_id)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGroup is a wrapper around the C function gtk_radio_tool_button_get_group.
func (recv *RadioToolButton) GetGroup() *glib.SList {
	retC := C.gtk_radio_tool_button_get_group((*C.GtkRadioToolButton)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetGroup is a wrapper around the C function gtk_radio_tool_button_set_group.
func (recv *RadioToolButton) SetGroup(group *glib.SList) {
	c_group := (*C.GSList)(group.ToC())

	C.gtk_radio_tool_button_set_group((*C.GtkRadioToolButton)(recv.native), c_group)

	return
}

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// GetLayout is a wrapper around the C function gtk_scale_get_layout.
func (recv *Scale) GetLayout() *pango.Layout {
	retC := C.gtk_scale_get_layout((*C.GtkScale)(recv.native))
	retGo := pango.LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLayoutOffsets is a wrapper around the C function gtk_scale_get_layout_offsets.
func (recv *Scale) GetLayoutOffsets() (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	C.gtk_scale_get_layout_offsets((*C.GtkScale)(recv.native), &c_x, &c_y)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// SeparatorToolItemNew is a wrapper around the C function gtk_separator_tool_item_new.
func SeparatorToolItemNew() *SeparatorToolItem {
	retC := C.gtk_separator_tool_item_new()
	retGo := SeparatorToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDraw is a wrapper around the C function gtk_separator_tool_item_get_draw.
func (recv *SeparatorToolItem) GetDraw() bool {
	retC := C.gtk_separator_tool_item_get_draw((*C.GtkSeparatorToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDraw is a wrapper around the C function gtk_separator_tool_item_set_draw.
func (recv *SeparatorToolItem) SetDraw(draw bool) {
	c_draw :=
		boolToGboolean(draw)

	C.gtk_separator_tool_item_set_draw((*C.GtkSeparatorToolItem)(recv.native), c_draw)

	return
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

var signalStyleRealizeId int
var signalStyleRealizeMap = make(map[int]StyleSignalRealizeCallback)
var signalStyleRealizeLock sync.Mutex

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
	signalStyleRealizeMap[signalStyleRealizeId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Style_signal_connect_realize(instance, C.gpointer(uintptr(signalStyleRealizeId)))
	return int(retC)
}

/*
DisconnectRealize disconnects a callback from the 'realize' signal for the Style.

The connectionID should be a value returned from a call to ConnectRealize.
*/
func (recv *Style) DisconnectRealize(connectionID int) {
	signalStyleRealizeLock.Lock()
	defer signalStyleRealizeLock.Unlock()

	_, exists := signalStyleRealizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalStyleRealizeMap, connectionID)
}

//export Style_realizeHandler
func Style_realizeHandler() {}

var signalStyleUnrealizeId int
var signalStyleUnrealizeMap = make(map[int]StyleSignalUnrealizeCallback)
var signalStyleUnrealizeLock sync.Mutex

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
	signalStyleUnrealizeMap[signalStyleUnrealizeId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Style_signal_connect_unrealize(instance, C.gpointer(uintptr(signalStyleUnrealizeId)))
	return int(retC)
}

/*
DisconnectUnrealize disconnects a callback from the 'unrealize' signal for the Style.

The connectionID should be a value returned from a call to ConnectUnrealize.
*/
func (recv *Style) DisconnectUnrealize(connectionID int) {
	signalStyleUnrealizeLock.Lock()
	defer signalStyleUnrealizeLock.Unlock()

	_, exists := signalStyleUnrealizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalStyleUnrealizeMap, connectionID)
}

//export Style_unrealizeHandler
func Style_unrealizeHandler() {}

// SelectRange is a wrapper around the C function gtk_text_buffer_select_range.
func (recv *TextBuffer) SelectRange(ins *TextIter, bound *TextIter) {
	c_ins := (*C.GtkTextIter)(ins.ToC())

	c_bound := (*C.GtkTextIter)(bound.ToC())

	C.gtk_text_buffer_select_range((*C.GtkTextBuffer)(recv.native), c_ins, c_bound)

	return
}

// GetAcceptsTab is a wrapper around the C function gtk_text_view_get_accepts_tab.
func (recv *TextView) GetAcceptsTab() bool {
	retC := C.gtk_text_view_get_accepts_tab((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetOverwrite is a wrapper around the C function gtk_text_view_get_overwrite.
func (recv *TextView) GetOverwrite() bool {
	retC := C.gtk_text_view_get_overwrite((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAcceptsTab is a wrapper around the C function gtk_text_view_set_accepts_tab.
func (recv *TextView) SetAcceptsTab(acceptsTab bool) {
	c_accepts_tab :=
		boolToGboolean(acceptsTab)

	C.gtk_text_view_set_accepts_tab((*C.GtkTextView)(recv.native), c_accepts_tab)

	return
}

// SetOverwrite is a wrapper around the C function gtk_text_view_set_overwrite.
func (recv *TextView) SetOverwrite(overwrite bool) {
	c_overwrite :=
		boolToGboolean(overwrite)

	C.gtk_text_view_set_overwrite((*C.GtkTextView)(recv.native), c_overwrite)

	return
}

// ToggleActionNew is a wrapper around the C function gtk_toggle_action_new.
func ToggleActionNew(name string, label string, tooltip string, stockId string) *ToggleAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_toggle_action_new(c_name, c_label, c_tooltip, c_stock_id)
	retGo := ToggleActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_toggle_action_get_active.
func (recv *ToggleAction) GetActive() bool {
	retC := C.gtk_toggle_action_get_active((*C.GtkToggleAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDrawAsRadio is a wrapper around the C function gtk_toggle_action_get_draw_as_radio.
func (recv *ToggleAction) GetDrawAsRadio() bool {
	retC := C.gtk_toggle_action_get_draw_as_radio((*C.GtkToggleAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActive is a wrapper around the C function gtk_toggle_action_set_active.
func (recv *ToggleAction) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_toggle_action_set_active((*C.GtkToggleAction)(recv.native), c_is_active)

	return
}

// SetDrawAsRadio is a wrapper around the C function gtk_toggle_action_set_draw_as_radio.
func (recv *ToggleAction) SetDrawAsRadio(drawAsRadio bool) {
	c_draw_as_radio :=
		boolToGboolean(drawAsRadio)

	C.gtk_toggle_action_set_draw_as_radio((*C.GtkToggleAction)(recv.native), c_draw_as_radio)

	return
}

// Toggled is a wrapper around the C function gtk_toggle_action_toggled.
func (recv *ToggleAction) Toggled() {
	C.gtk_toggle_action_toggled((*C.GtkToggleAction)(recv.native))

	return
}

// ToggleToolButtonNew is a wrapper around the C function gtk_toggle_tool_button_new.
func ToggleToolButtonNew() *ToggleToolButton {
	retC := C.gtk_toggle_tool_button_new()
	retGo := ToggleToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleToolButtonNewFromStock is a wrapper around the C function gtk_toggle_tool_button_new_from_stock.
func ToggleToolButtonNewFromStock(stockId string) *ToggleToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_toggle_tool_button_new_from_stock(c_stock_id)
	retGo := ToggleToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_toggle_tool_button_get_active.
func (recv *ToggleToolButton) GetActive() bool {
	retC := C.gtk_toggle_tool_button_get_active((*C.GtkToggleToolButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActive is a wrapper around the C function gtk_toggle_tool_button_set_active.
func (recv *ToggleToolButton) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_toggle_tool_button_set_active((*C.GtkToggleToolButton)(recv.native), c_is_active)

	return
}

// ToolButtonNew is a wrapper around the C function gtk_tool_button_new.
func ToolButtonNew(iconWidget *Widget, label string) *ToolButton {
	c_icon_widget := (*C.GtkWidget)(iconWidget.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_tool_button_new(c_icon_widget, c_label)
	retGo := ToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToolButtonNewFromStock is a wrapper around the C function gtk_tool_button_new_from_stock.
func ToolButtonNewFromStock(stockId string) *ToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_tool_button_new_from_stock(c_stock_id)
	retGo := ToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIconWidget is a wrapper around the C function gtk_tool_button_get_icon_widget.
func (recv *ToolButton) GetIconWidget() *Widget {
	retC := C.gtk_tool_button_get_icon_widget((*C.GtkToolButton)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLabel is a wrapper around the C function gtk_tool_button_get_label.
func (recv *ToolButton) GetLabel() string {
	retC := C.gtk_tool_button_get_label((*C.GtkToolButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLabelWidget is a wrapper around the C function gtk_tool_button_get_label_widget.
func (recv *ToolButton) GetLabelWidget() *Widget {
	retC := C.gtk_tool_button_get_label_widget((*C.GtkToolButton)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStockId is a wrapper around the C function gtk_tool_button_get_stock_id.
func (recv *ToolButton) GetStockId() string {
	retC := C.gtk_tool_button_get_stock_id((*C.GtkToolButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUseUnderline is a wrapper around the C function gtk_tool_button_get_use_underline.
func (recv *ToolButton) GetUseUnderline() bool {
	retC := C.gtk_tool_button_get_use_underline((*C.GtkToolButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetIconWidget is a wrapper around the C function gtk_tool_button_set_icon_widget.
func (recv *ToolButton) SetIconWidget(iconWidget *Widget) {
	c_icon_widget := (*C.GtkWidget)(iconWidget.ToC())

	C.gtk_tool_button_set_icon_widget((*C.GtkToolButton)(recv.native), c_icon_widget)

	return
}

// SetLabel is a wrapper around the C function gtk_tool_button_set_label.
func (recv *ToolButton) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_tool_button_set_label((*C.GtkToolButton)(recv.native), c_label)

	return
}

// SetLabelWidget is a wrapper around the C function gtk_tool_button_set_label_widget.
func (recv *ToolButton) SetLabelWidget(labelWidget *Widget) {
	c_label_widget := (*C.GtkWidget)(labelWidget.ToC())

	C.gtk_tool_button_set_label_widget((*C.GtkToolButton)(recv.native), c_label_widget)

	return
}

// SetStockId is a wrapper around the C function gtk_tool_button_set_stock_id.
func (recv *ToolButton) SetStockId(stockId string) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_tool_button_set_stock_id((*C.GtkToolButton)(recv.native), c_stock_id)

	return
}

// SetUseUnderline is a wrapper around the C function gtk_tool_button_set_use_underline.
func (recv *ToolButton) SetUseUnderline(useUnderline bool) {
	c_use_underline :=
		boolToGboolean(useUnderline)

	C.gtk_tool_button_set_use_underline((*C.GtkToolButton)(recv.native), c_use_underline)

	return
}

// ToolItemNew is a wrapper around the C function gtk_tool_item_new.
func ToolItemNew() *ToolItem {
	retC := C.gtk_tool_item_new()
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetExpand is a wrapper around the C function gtk_tool_item_get_expand.
func (recv *ToolItem) GetExpand() bool {
	retC := C.gtk_tool_item_get_expand((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHomogeneous is a wrapper around the C function gtk_tool_item_get_homogeneous.
func (recv *ToolItem) GetHomogeneous() bool {
	retC := C.gtk_tool_item_get_homogeneous((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tool_item_get_icon_size : no return generator

// GetIsImportant is a wrapper around the C function gtk_tool_item_get_is_important.
func (recv *ToolItem) GetIsImportant() bool {
	retC := C.gtk_tool_item_get_is_important((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetOrientation is a wrapper around the C function gtk_tool_item_get_orientation.
func (recv *ToolItem) GetOrientation() Orientation {
	retC := C.gtk_tool_item_get_orientation((*C.GtkToolItem)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// GetProxyMenuItem is a wrapper around the C function gtk_tool_item_get_proxy_menu_item.
func (recv *ToolItem) GetProxyMenuItem(menuItemId string) *Widget {
	c_menu_item_id := C.CString(menuItemId)
	defer C.free(unsafe.Pointer(c_menu_item_id))

	retC := C.gtk_tool_item_get_proxy_menu_item((*C.GtkToolItem)(recv.native), c_menu_item_id)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetReliefStyle is a wrapper around the C function gtk_tool_item_get_relief_style.
func (recv *ToolItem) GetReliefStyle() ReliefStyle {
	retC := C.gtk_tool_item_get_relief_style((*C.GtkToolItem)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// GetToolbarStyle is a wrapper around the C function gtk_tool_item_get_toolbar_style.
func (recv *ToolItem) GetToolbarStyle() ToolbarStyle {
	retC := C.gtk_tool_item_get_toolbar_style((*C.GtkToolItem)(recv.native))
	retGo := (ToolbarStyle)(retC)

	return retGo
}

// GetUseDragWindow is a wrapper around the C function gtk_tool_item_get_use_drag_window.
func (recv *ToolItem) GetUseDragWindow() bool {
	retC := C.gtk_tool_item_get_use_drag_window((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisibleHorizontal is a wrapper around the C function gtk_tool_item_get_visible_horizontal.
func (recv *ToolItem) GetVisibleHorizontal() bool {
	retC := C.gtk_tool_item_get_visible_horizontal((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisibleVertical is a wrapper around the C function gtk_tool_item_get_visible_vertical.
func (recv *ToolItem) GetVisibleVertical() bool {
	retC := C.gtk_tool_item_get_visible_vertical((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// RetrieveProxyMenuItem is a wrapper around the C function gtk_tool_item_retrieve_proxy_menu_item.
func (recv *ToolItem) RetrieveProxyMenuItem() *Widget {
	retC := C.gtk_tool_item_retrieve_proxy_menu_item((*C.GtkToolItem)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetExpand is a wrapper around the C function gtk_tool_item_set_expand.
func (recv *ToolItem) SetExpand(expand bool) {
	c_expand :=
		boolToGboolean(expand)

	C.gtk_tool_item_set_expand((*C.GtkToolItem)(recv.native), c_expand)

	return
}

// SetHomogeneous is a wrapper around the C function gtk_tool_item_set_homogeneous.
func (recv *ToolItem) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_tool_item_set_homogeneous((*C.GtkToolItem)(recv.native), c_homogeneous)

	return
}

// SetIsImportant is a wrapper around the C function gtk_tool_item_set_is_important.
func (recv *ToolItem) SetIsImportant(isImportant bool) {
	c_is_important :=
		boolToGboolean(isImportant)

	C.gtk_tool_item_set_is_important((*C.GtkToolItem)(recv.native), c_is_important)

	return
}

// SetProxyMenuItem is a wrapper around the C function gtk_tool_item_set_proxy_menu_item.
func (recv *ToolItem) SetProxyMenuItem(menuItemId string, menuItem *Widget) {
	c_menu_item_id := C.CString(menuItemId)
	defer C.free(unsafe.Pointer(c_menu_item_id))

	c_menu_item := (*C.GtkWidget)(menuItem.ToC())

	C.gtk_tool_item_set_proxy_menu_item((*C.GtkToolItem)(recv.native), c_menu_item_id, c_menu_item)

	return
}

// SetUseDragWindow is a wrapper around the C function gtk_tool_item_set_use_drag_window.
func (recv *ToolItem) SetUseDragWindow(useDragWindow bool) {
	c_use_drag_window :=
		boolToGboolean(useDragWindow)

	C.gtk_tool_item_set_use_drag_window((*C.GtkToolItem)(recv.native), c_use_drag_window)

	return
}

// SetVisibleHorizontal is a wrapper around the C function gtk_tool_item_set_visible_horizontal.
func (recv *ToolItem) SetVisibleHorizontal(visibleHorizontal bool) {
	c_visible_horizontal :=
		boolToGboolean(visibleHorizontal)

	C.gtk_tool_item_set_visible_horizontal((*C.GtkToolItem)(recv.native), c_visible_horizontal)

	return
}

// SetVisibleVertical is a wrapper around the C function gtk_tool_item_set_visible_vertical.
func (recv *ToolItem) SetVisibleVertical(visibleVertical bool) {
	c_visible_vertical :=
		boolToGboolean(visibleVertical)

	C.gtk_tool_item_set_visible_vertical((*C.GtkToolItem)(recv.native), c_visible_vertical)

	return
}

// GetDropIndex is a wrapper around the C function gtk_toolbar_get_drop_index.
func (recv *Toolbar) GetDropIndex(x int32, y int32) int32 {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_toolbar_get_drop_index((*C.GtkToolbar)(recv.native), c_x, c_y)
	retGo := (int32)(retC)

	return retGo
}

// GetItemIndex is a wrapper around the C function gtk_toolbar_get_item_index.
func (recv *Toolbar) GetItemIndex(item *ToolItem) int32 {
	c_item := (*C.GtkToolItem)(item.ToC())

	retC := C.gtk_toolbar_get_item_index((*C.GtkToolbar)(recv.native), c_item)
	retGo := (int32)(retC)

	return retGo
}

// GetNItems is a wrapper around the C function gtk_toolbar_get_n_items.
func (recv *Toolbar) GetNItems() int32 {
	retC := C.gtk_toolbar_get_n_items((*C.GtkToolbar)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNthItem is a wrapper around the C function gtk_toolbar_get_nth_item.
func (recv *Toolbar) GetNthItem(n int32) *ToolItem {
	c_n := (C.gint)(n)

	retC := C.gtk_toolbar_get_nth_item((*C.GtkToolbar)(recv.native), c_n)
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetReliefStyle is a wrapper around the C function gtk_toolbar_get_relief_style.
func (recv *Toolbar) GetReliefStyle() ReliefStyle {
	retC := C.gtk_toolbar_get_relief_style((*C.GtkToolbar)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// GetShowArrow is a wrapper around the C function gtk_toolbar_get_show_arrow.
func (recv *Toolbar) GetShowArrow() bool {
	retC := C.gtk_toolbar_get_show_arrow((*C.GtkToolbar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Insert is a wrapper around the C function gtk_toolbar_insert.
func (recv *Toolbar) Insert(item *ToolItem, pos int32) {
	c_item := (*C.GtkToolItem)(item.ToC())

	c_pos := (C.gint)(pos)

	C.gtk_toolbar_insert((*C.GtkToolbar)(recv.native), c_item, c_pos)

	return
}

// SetDropHighlightItem is a wrapper around the C function gtk_toolbar_set_drop_highlight_item.
func (recv *Toolbar) SetDropHighlightItem(toolItem *ToolItem, index int32) {
	c_tool_item := (*C.GtkToolItem)(toolItem.ToC())

	c_index_ := (C.gint)(index)

	C.gtk_toolbar_set_drop_highlight_item((*C.GtkToolbar)(recv.native), c_tool_item, c_index_)

	return
}

// SetShowArrow is a wrapper around the C function gtk_toolbar_set_show_arrow.
func (recv *Toolbar) SetShowArrow(showArrow bool) {
	c_show_arrow :=
		boolToGboolean(showArrow)

	C.gtk_toolbar_set_show_arrow((*C.GtkToolbar)(recv.native), c_show_arrow)

	return
}

// ClearCache is a wrapper around the C function gtk_tree_model_filter_clear_cache.
func (recv *TreeModelFilter) ClearCache() {
	C.gtk_tree_model_filter_clear_cache((*C.GtkTreeModelFilter)(recv.native))

	return
}

// ConvertChildIterToIter is a wrapper around the C function gtk_tree_model_filter_convert_child_iter_to_iter.
func (recv *TreeModelFilter) ConvertChildIterToIter(childIter *TreeIter) (bool, *TreeIter) {
	var c_filter_iter C.GtkTreeIter

	c_child_iter := (*C.GtkTreeIter)(childIter.ToC())

	retC := C.gtk_tree_model_filter_convert_child_iter_to_iter((*C.GtkTreeModelFilter)(recv.native), &c_filter_iter, c_child_iter)
	retGo := retC == C.TRUE

	filterIter := TreeIterNewFromC(unsafe.Pointer(&c_filter_iter))

	return retGo, filterIter
}

// ConvertChildPathToPath is a wrapper around the C function gtk_tree_model_filter_convert_child_path_to_path.
func (recv *TreeModelFilter) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	c_child_path := (*C.GtkTreePath)(childPath.ToC())

	retC := C.gtk_tree_model_filter_convert_child_path_to_path((*C.GtkTreeModelFilter)(recv.native), c_child_path)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ConvertIterToChildIter is a wrapper around the C function gtk_tree_model_filter_convert_iter_to_child_iter.
func (recv *TreeModelFilter) ConvertIterToChildIter(filterIter *TreeIter) *TreeIter {
	var c_child_iter C.GtkTreeIter

	c_filter_iter := (*C.GtkTreeIter)(filterIter.ToC())

	C.gtk_tree_model_filter_convert_iter_to_child_iter((*C.GtkTreeModelFilter)(recv.native), &c_child_iter, c_filter_iter)

	childIter := TreeIterNewFromC(unsafe.Pointer(&c_child_iter))

	return childIter
}

// ConvertPathToChildPath is a wrapper around the C function gtk_tree_model_filter_convert_path_to_child_path.
func (recv *TreeModelFilter) ConvertPathToChildPath(filterPath *TreePath) *TreePath {
	c_filter_path := (*C.GtkTreePath)(filterPath.ToC())

	retC := C.gtk_tree_model_filter_convert_path_to_child_path((*C.GtkTreeModelFilter)(recv.native), c_filter_path)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_model_filter_get_model : no return generator

// Refilter is a wrapper around the C function gtk_tree_model_filter_refilter.
func (recv *TreeModelFilter) Refilter() {
	C.gtk_tree_model_filter_refilter((*C.GtkTreeModelFilter)(recv.native))

	return
}

// Unsupported : gtk_tree_model_filter_set_modify_func : unsupported parameter types : no param type

// SetVisibleColumn is a wrapper around the C function gtk_tree_model_filter_set_visible_column.
func (recv *TreeModelFilter) SetVisibleColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_tree_model_filter_set_visible_column((*C.GtkTreeModelFilter)(recv.native), c_column)

	return
}

// Unsupported : gtk_tree_model_filter_set_visible_func : unsupported parameter func : no type generator for TreeModelFilterVisibleFunc, GtkTreeModelFilterVisibleFunc

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// GetExpand is a wrapper around the C function gtk_tree_view_column_get_expand.
func (recv *TreeViewColumn) GetExpand() bool {
	retC := C.gtk_tree_view_column_get_expand((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetExpand is a wrapper around the C function gtk_tree_view_column_set_expand.
func (recv *TreeViewColumn) SetExpand(expand bool) {
	c_expand :=
		boolToGboolean(expand)

	C.gtk_tree_view_column_set_expand((*C.GtkTreeViewColumn)(recv.native), c_expand)

	return
}

var signalUIManagerActionsChangedId int
var signalUIManagerActionsChangedMap = make(map[int]UIManagerSignalActionsChangedCallback)
var signalUIManagerActionsChangedLock sync.Mutex

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
	signalUIManagerActionsChangedMap[signalUIManagerActionsChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.UIManager_signal_connect_actions_changed(instance, C.gpointer(uintptr(signalUIManagerActionsChangedId)))
	return int(retC)
}

/*
DisconnectActionsChanged disconnects a callback from the 'actions-changed' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectActionsChanged.
*/
func (recv *UIManager) DisconnectActionsChanged(connectionID int) {
	signalUIManagerActionsChangedLock.Lock()
	defer signalUIManagerActionsChangedLock.Unlock()

	_, exists := signalUIManagerActionsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalUIManagerActionsChangedMap, connectionID)
}

//export UIManager_actionsChangedHandler
func UIManager_actionsChangedHandler() {}

var signalUIManagerAddWidgetId int
var signalUIManagerAddWidgetMap = make(map[int]UIManagerSignalAddWidgetCallback)
var signalUIManagerAddWidgetLock sync.Mutex

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
	signalUIManagerAddWidgetMap[signalUIManagerAddWidgetId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.UIManager_signal_connect_add_widget(instance, C.gpointer(uintptr(signalUIManagerAddWidgetId)))
	return int(retC)
}

/*
DisconnectAddWidget disconnects a callback from the 'add-widget' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectAddWidget.
*/
func (recv *UIManager) DisconnectAddWidget(connectionID int) {
	signalUIManagerAddWidgetLock.Lock()
	defer signalUIManagerAddWidgetLock.Unlock()

	_, exists := signalUIManagerAddWidgetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalUIManagerAddWidgetMap, connectionID)
}

//export UIManager_addWidgetHandler
func UIManager_addWidgetHandler(c_widget *C.GtkWidget) {
	widget := WidgetNewFromC(unsafe.Pointer(c_widget))

}

var signalUIManagerConnectProxyId int
var signalUIManagerConnectProxyMap = make(map[int]UIManagerSignalConnectProxyCallback)
var signalUIManagerConnectProxyLock sync.Mutex

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
	signalUIManagerConnectProxyMap[signalUIManagerConnectProxyId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.UIManager_signal_connect_connect_proxy(instance, C.gpointer(uintptr(signalUIManagerConnectProxyId)))
	return int(retC)
}

/*
DisconnectConnectProxy disconnects a callback from the 'connect-proxy' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectConnectProxy.
*/
func (recv *UIManager) DisconnectConnectProxy(connectionID int) {
	signalUIManagerConnectProxyLock.Lock()
	defer signalUIManagerConnectProxyLock.Unlock()

	_, exists := signalUIManagerConnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalUIManagerConnectProxyMap, connectionID)
}

//export UIManager_connectProxyHandler
func UIManager_connectProxyHandler(c_action *C.GtkAction, c_proxy *C.GtkWidget) {
	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

}

var signalUIManagerDisconnectProxyId int
var signalUIManagerDisconnectProxyMap = make(map[int]UIManagerSignalDisconnectProxyCallback)
var signalUIManagerDisconnectProxyLock sync.Mutex

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
	signalUIManagerDisconnectProxyMap[signalUIManagerDisconnectProxyId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.UIManager_signal_connect_disconnect_proxy(instance, C.gpointer(uintptr(signalUIManagerDisconnectProxyId)))
	return int(retC)
}

/*
DisconnectDisconnectProxy disconnects a callback from the 'disconnect-proxy' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectDisconnectProxy.
*/
func (recv *UIManager) DisconnectDisconnectProxy(connectionID int) {
	signalUIManagerDisconnectProxyLock.Lock()
	defer signalUIManagerDisconnectProxyLock.Unlock()

	_, exists := signalUIManagerDisconnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalUIManagerDisconnectProxyMap, connectionID)
}

//export UIManager_disconnectProxyHandler
func UIManager_disconnectProxyHandler(c_action *C.GtkAction, c_proxy *C.GtkWidget) {
	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

}

var signalUIManagerPostActivateId int
var signalUIManagerPostActivateMap = make(map[int]UIManagerSignalPostActivateCallback)
var signalUIManagerPostActivateLock sync.Mutex

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
	signalUIManagerPostActivateMap[signalUIManagerPostActivateId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.UIManager_signal_connect_post_activate(instance, C.gpointer(uintptr(signalUIManagerPostActivateId)))
	return int(retC)
}

/*
DisconnectPostActivate disconnects a callback from the 'post-activate' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectPostActivate.
*/
func (recv *UIManager) DisconnectPostActivate(connectionID int) {
	signalUIManagerPostActivateLock.Lock()
	defer signalUIManagerPostActivateLock.Unlock()

	_, exists := signalUIManagerPostActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalUIManagerPostActivateMap, connectionID)
}

//export UIManager_postActivateHandler
func UIManager_postActivateHandler(c_action *C.GtkAction) {
	action := ActionNewFromC(unsafe.Pointer(c_action))

}

var signalUIManagerPreActivateId int
var signalUIManagerPreActivateMap = make(map[int]UIManagerSignalPreActivateCallback)
var signalUIManagerPreActivateLock sync.Mutex

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
	signalUIManagerPreActivateMap[signalUIManagerPreActivateId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.UIManager_signal_connect_pre_activate(instance, C.gpointer(uintptr(signalUIManagerPreActivateId)))
	return int(retC)
}

/*
DisconnectPreActivate disconnects a callback from the 'pre-activate' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectPreActivate.
*/
func (recv *UIManager) DisconnectPreActivate(connectionID int) {
	signalUIManagerPreActivateLock.Lock()
	defer signalUIManagerPreActivateLock.Unlock()

	_, exists := signalUIManagerPreActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalUIManagerPreActivateMap, connectionID)
}

//export UIManager_preActivateHandler
func UIManager_preActivateHandler(c_action *C.GtkAction) {
	action := ActionNewFromC(unsafe.Pointer(c_action))

}

// UIManagerNew is a wrapper around the C function gtk_ui_manager_new.
func UIManagerNew() *UIManager {
	retC := C.gtk_ui_manager_new()
	retGo := UIManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddUi is a wrapper around the C function gtk_ui_manager_add_ui.
func (recv *UIManager) AddUi(mergeId uint32, path string, name string, action string, type_ UIManagerItemType, top bool) {
	c_merge_id := (C.guint)(mergeId)

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_action := C.CString(action)
	defer C.free(unsafe.Pointer(c_action))

	c_type := (C.GtkUIManagerItemType)(type_)

	c_top :=
		boolToGboolean(top)

	C.gtk_ui_manager_add_ui((*C.GtkUIManager)(recv.native), c_merge_id, c_path, c_name, c_action, c_type, c_top)

	return
}

// AddUiFromFile is a wrapper around the C function gtk_ui_manager_add_ui_from_file.
func (recv *UIManager) AddUiFromFile(filename string) (uint32, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gtk_ui_manager_add_ui_from_file((*C.GtkUIManager)(recv.native), c_filename, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// AddUiFromString is a wrapper around the C function gtk_ui_manager_add_ui_from_string.
func (recv *UIManager) AddUiFromString(buffer string, length int64) (uint32, error) {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_length := (C.gssize)(length)

	var cThrowableError *C.GError

	retC := C.gtk_ui_manager_add_ui_from_string((*C.GtkUIManager)(recv.native), c_buffer, c_length, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// EnsureUpdate is a wrapper around the C function gtk_ui_manager_ensure_update.
func (recv *UIManager) EnsureUpdate() {
	C.gtk_ui_manager_ensure_update((*C.GtkUIManager)(recv.native))

	return
}

// GetAccelGroup is a wrapper around the C function gtk_ui_manager_get_accel_group.
func (recv *UIManager) GetAccelGroup() *AccelGroup {
	retC := C.gtk_ui_manager_get_accel_group((*C.GtkUIManager)(recv.native))
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAction is a wrapper around the C function gtk_ui_manager_get_action.
func (recv *UIManager) GetAction(path string) *Action {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.gtk_ui_manager_get_action((*C.GtkUIManager)(recv.native), c_path)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActionGroups is a wrapper around the C function gtk_ui_manager_get_action_groups.
func (recv *UIManager) GetActionGroups() *glib.List {
	retC := C.gtk_ui_manager_get_action_groups((*C.GtkUIManager)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAddTearoffs is a wrapper around the C function gtk_ui_manager_get_add_tearoffs.
func (recv *UIManager) GetAddTearoffs() bool {
	retC := C.gtk_ui_manager_get_add_tearoffs((*C.GtkUIManager)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetToplevels is a wrapper around the C function gtk_ui_manager_get_toplevels.
func (recv *UIManager) GetToplevels(types UIManagerItemType) *glib.SList {
	c_types := (C.GtkUIManagerItemType)(types)

	retC := C.gtk_ui_manager_get_toplevels((*C.GtkUIManager)(recv.native), c_types)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUi is a wrapper around the C function gtk_ui_manager_get_ui.
func (recv *UIManager) GetUi() string {
	retC := C.gtk_ui_manager_get_ui((*C.GtkUIManager)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetWidget is a wrapper around the C function gtk_ui_manager_get_widget.
func (recv *UIManager) GetWidget(path string) *Widget {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.gtk_ui_manager_get_widget((*C.GtkUIManager)(recv.native), c_path)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertActionGroup is a wrapper around the C function gtk_ui_manager_insert_action_group.
func (recv *UIManager) InsertActionGroup(actionGroup *ActionGroup, pos int32) {
	c_action_group := (*C.GtkActionGroup)(actionGroup.ToC())

	c_pos := (C.gint)(pos)

	C.gtk_ui_manager_insert_action_group((*C.GtkUIManager)(recv.native), c_action_group, c_pos)

	return
}

// NewMergeId is a wrapper around the C function gtk_ui_manager_new_merge_id.
func (recv *UIManager) NewMergeId() uint32 {
	retC := C.gtk_ui_manager_new_merge_id((*C.GtkUIManager)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// RemoveActionGroup is a wrapper around the C function gtk_ui_manager_remove_action_group.
func (recv *UIManager) RemoveActionGroup(actionGroup *ActionGroup) {
	c_action_group := (*C.GtkActionGroup)(actionGroup.ToC())

	C.gtk_ui_manager_remove_action_group((*C.GtkUIManager)(recv.native), c_action_group)

	return
}

// RemoveUi is a wrapper around the C function gtk_ui_manager_remove_ui.
func (recv *UIManager) RemoveUi(mergeId uint32) {
	c_merge_id := (C.guint)(mergeId)

	C.gtk_ui_manager_remove_ui((*C.GtkUIManager)(recv.native), c_merge_id)

	return
}

// SetAddTearoffs is a wrapper around the C function gtk_ui_manager_set_add_tearoffs.
func (recv *UIManager) SetAddTearoffs(addTearoffs bool) {
	c_add_tearoffs :=
		boolToGboolean(addTearoffs)

	C.gtk_ui_manager_set_add_tearoffs((*C.GtkUIManager)(recv.native), c_add_tearoffs)

	return
}

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// AddMnemonicLabel is a wrapper around the C function gtk_widget_add_mnemonic_label.
func (recv *Widget) AddMnemonicLabel(label *Widget) {
	c_label := (*C.GtkWidget)(label.ToC())

	C.gtk_widget_add_mnemonic_label((*C.GtkWidget)(recv.native), c_label)

	return
}

// CanActivateAccel is a wrapper around the C function gtk_widget_can_activate_accel.
func (recv *Widget) CanActivateAccel(signalId uint32) bool {
	c_signal_id := (C.guint)(signalId)

	retC := C.gtk_widget_can_activate_accel((*C.GtkWidget)(recv.native), c_signal_id)
	retGo := retC == C.TRUE

	return retGo
}

// DragSourceGetTargetList is a wrapper around the C function gtk_drag_source_get_target_list.
func (recv *Widget) DragSourceGetTargetList() *TargetList {
	retC := C.gtk_drag_source_get_target_list((*C.GtkWidget)(recv.native))
	retGo := TargetListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DragSourceSetTargetList is a wrapper around the C function gtk_drag_source_set_target_list.
func (recv *Widget) DragSourceSetTargetList(targetList *TargetList) {
	c_target_list := (*C.GtkTargetList)(targetList.ToC())

	C.gtk_drag_source_set_target_list((*C.GtkWidget)(recv.native), c_target_list)

	return
}

// GetNoShowAll is a wrapper around the C function gtk_widget_get_no_show_all.
func (recv *Widget) GetNoShowAll() bool {
	retC := C.gtk_widget_get_no_show_all((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ListMnemonicLabels is a wrapper around the C function gtk_widget_list_mnemonic_labels.
func (recv *Widget) ListMnemonicLabels() *glib.List {
	retC := C.gtk_widget_list_mnemonic_labels((*C.GtkWidget)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// QueueResizeNoRedraw is a wrapper around the C function gtk_widget_queue_resize_no_redraw.
func (recv *Widget) QueueResizeNoRedraw() {
	C.gtk_widget_queue_resize_no_redraw((*C.GtkWidget)(recv.native))

	return
}

// RemoveMnemonicLabel is a wrapper around the C function gtk_widget_remove_mnemonic_label.
func (recv *Widget) RemoveMnemonicLabel(label *Widget) {
	c_label := (*C.GtkWidget)(label.ToC())

	C.gtk_widget_remove_mnemonic_label((*C.GtkWidget)(recv.native), c_label)

	return
}

// SetNoShowAll is a wrapper around the C function gtk_widget_set_no_show_all.
func (recv *Widget) SetNoShowAll(noShowAll bool) {
	c_no_show_all :=
		boolToGboolean(noShowAll)

	C.gtk_widget_set_no_show_all((*C.GtkWidget)(recv.native), c_no_show_all)

	return
}

// ActivateKey is a wrapper around the C function gtk_window_activate_key.
func (recv *Window) ActivateKey(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(event.ToC())

	retC := C.gtk_window_activate_key((*C.GtkWindow)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// GetAcceptFocus is a wrapper around the C function gtk_window_get_accept_focus.
func (recv *Window) GetAcceptFocus() bool {
	retC := C.gtk_window_get_accept_focus((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasToplevelFocus is a wrapper around the C function gtk_window_has_toplevel_focus.
func (recv *Window) HasToplevelFocus() bool {
	retC := C.gtk_window_has_toplevel_focus((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsActive is a wrapper around the C function gtk_window_is_active.
func (recv *Window) IsActive() bool {
	retC := C.gtk_window_is_active((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PropagateKeyEvent is a wrapper around the C function gtk_window_propagate_key_event.
func (recv *Window) PropagateKeyEvent(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(event.ToC())

	retC := C.gtk_window_propagate_key_event((*C.GtkWindow)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// SetAcceptFocus is a wrapper around the C function gtk_window_set_accept_focus.
func (recv *Window) SetAcceptFocus(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_accept_focus((*C.GtkWindow)(recv.native), c_setting)

	return
}

// SetKeepAbove is a wrapper around the C function gtk_window_set_keep_above.
func (recv *Window) SetKeepAbove(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_keep_above((*C.GtkWindow)(recv.native), c_setting)

	return
}

// SetKeepBelow is a wrapper around the C function gtk_window_set_keep_below.
func (recv *Window) SetKeepBelow(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_keep_below((*C.GtkWindow)(recv.native), c_setting)

	return
}
