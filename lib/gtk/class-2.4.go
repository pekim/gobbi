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

// Unsupported signal 'changed' for AccelMap : unsupported parameter accel_path : type utf8 :

type signalActionActivateDetail struct {
	callback  ActionSignalActivateCallback
	handlerID C.gulong
}

var signalActionActivateId int
var signalActionActivateMap = make(map[int]signalActionActivateDetail)
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
	index := int(uintptr(data))
	callback := signalActionActivateMap[index].callback
	callback()
}

// Creates a new #GtkAction object. To add the action to a
// #GtkActionGroup and set the accelerator for the action,
// call gtk_action_group_add_action_with_accel().
// See the [UI Definition section][XML-UI] for information on allowed action
// names.
/*

C function : gtk_action_new
*/
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

// Emits the “activate” signal on the specified action, if it isn't
// insensitive. This gets called by the proxy widgets when they get
// activated.
//
// It can also be used to manually activate an action.
/*

C function : gtk_action_activate
*/
func (recv *Action) Activate() {
	C.gtk_action_activate((*C.GtkAction)(recv.native))

	return
}

// Installs the accelerator for @action if @action has an
// accel path and group. See gtk_action_set_accel_path() and
// gtk_action_set_accel_group()
//
// Since multiple proxies may independently trigger the installation
// of the accelerator, the @action counts the number of times this
// function has been called and doesn’t remove the accelerator until
// gtk_action_disconnect_accelerator() has been called as many times.
/*

C function : gtk_action_connect_accelerator
*/
func (recv *Action) ConnectAccelerator() {
	C.gtk_action_connect_accelerator((*C.GtkAction)(recv.native))

	return
}

// This function is intended for use by action implementations to
// create icons displayed in the proxy widgets.
/*

C function : gtk_action_create_icon
*/
func (recv *Action) CreateIcon(iconSize IconSize) *Widget {
	c_icon_size := (C.GtkIconSize)(iconSize)

	retC := C.gtk_action_create_icon((*C.GtkAction)(recv.native), c_icon_size)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a menu item widget that proxies for the given action.
/*

C function : gtk_action_create_menu_item
*/
func (recv *Action) CreateMenuItem() *Widget {
	retC := C.gtk_action_create_menu_item((*C.GtkAction)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a toolbar item widget that proxies for the given action.
/*

C function : gtk_action_create_tool_item
*/
func (recv *Action) CreateToolItem() *Widget {
	retC := C.gtk_action_create_tool_item((*C.GtkAction)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Undoes the effect of one call to gtk_action_connect_accelerator().
/*

C function : gtk_action_disconnect_accelerator
*/
func (recv *Action) DisconnectAccelerator() {
	C.gtk_action_disconnect_accelerator((*C.GtkAction)(recv.native))

	return
}

// Returns the name of the action.
/*

C function : gtk_action_get_name
*/
func (recv *Action) GetName() string {
	retC := C.gtk_action_get_name((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the proxy widgets for an action.
// See also gtk_activatable_get_related_action().
/*

C function : gtk_action_get_proxies
*/
func (recv *Action) GetProxies() *glib.SList {
	retC := C.gtk_action_get_proxies((*C.GtkAction)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether the action itself is sensitive. Note that this doesn’t
// necessarily mean effective sensitivity. See gtk_action_is_sensitive()
// for that.
/*

C function : gtk_action_get_sensitive
*/
func (recv *Action) GetSensitive() bool {
	retC := C.gtk_action_get_sensitive((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the action itself is visible. Note that this doesn’t
// necessarily mean effective visibility. See gtk_action_is_sensitive()
// for that.
/*

C function : gtk_action_get_visible
*/
func (recv *Action) GetVisible() bool {
	retC := C.gtk_action_get_visible((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the action is effectively sensitive.
/*

C function : gtk_action_is_sensitive
*/
func (recv *Action) IsSensitive() bool {
	retC := C.gtk_action_is_sensitive((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the action is effectively visible.
/*

C function : gtk_action_is_visible
*/
func (recv *Action) IsVisible() bool {
	retC := C.gtk_action_is_visible((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the #GtkAccelGroup in which the accelerator for this action
// will be installed.
/*

C function : gtk_action_set_accel_group
*/
func (recv *Action) SetAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(C.NULL)
	if accelGroup != nil {
		c_accel_group = (*C.GtkAccelGroup)(accelGroup.ToC())
	}

	C.gtk_action_set_accel_group((*C.GtkAction)(recv.native), c_accel_group)

	return
}

// Sets the accel path for this action.  All proxy widgets associated
// with the action will have this accel path, so that their
// accelerators are consistent.
//
// Note that @accel_path string will be stored in a #GQuark. Therefore, if you
// pass a static string, you can save some memory by interning it first with
// g_intern_static_string().
/*

C function : gtk_action_set_accel_path
*/
func (recv *Action) SetAccelPath(accelPath string) {
	c_accel_path := C.CString(accelPath)
	defer C.free(unsafe.Pointer(c_accel_path))

	C.gtk_action_set_accel_path((*C.GtkAction)(recv.native), c_accel_path)

	return
}

type signalActionGroupConnectProxyDetail struct {
	callback  ActionGroupSignalConnectProxyCallback
	handlerID C.gulong
}

var signalActionGroupConnectProxyId int
var signalActionGroupConnectProxyMap = make(map[int]signalActionGroupConnectProxyDetail)
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
	action := ActionNewFromC(unsafe.Pointer(c_action))

	index := int(uintptr(data))
	callback := signalActionGroupPreActivateMap[index].callback
	callback(action)
}

// Creates a new #GtkActionGroup object. The name of the action group
// is used when associating [keybindings][Action-Accel]
// with the actions.
/*

C function : gtk_action_group_new
*/
func ActionGroupNew(name string) *ActionGroup {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_action_group_new(c_name)
	retGo := ActionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds an action object to the action group. Note that this function
// does not set up the accel path of the action, which can lead to problems
// if a user tries to modify the accelerator of a menuitem associated with
// the action. Therefore you must either set the accel path yourself with
// gtk_action_set_accel_path(), or use
// `gtk_action_group_add_action_with_accel (..., NULL)`.
/*

C function : gtk_action_group_add_action
*/
func (recv *ActionGroup) AddAction(action *Action) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	C.gtk_action_group_add_action((*C.GtkActionGroup)(recv.native), c_action)

	return
}

// Adds an action object to the action group and sets up the accelerator.
//
// If @accelerator is %NULL, attempts to use the accelerator associated
// with the stock_id of the action.
//
// Accel paths are set to `<Actions>/group-name/action-name`.
/*

C function : gtk_action_group_add_action_with_accel
*/
func (recv *ActionGroup) AddActionWithAccel(action *Action, accelerator string) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	c_accelerator := C.CString(accelerator)
	defer C.free(unsafe.Pointer(c_accelerator))

	C.gtk_action_group_add_action_with_accel((*C.GtkActionGroup)(recv.native), c_action, c_accelerator)

	return
}

// Unsupported : gtk_action_group_add_actions : unsupported parameter entries :

// Unsupported : gtk_action_group_add_actions_full : unsupported parameter entries :

// Unsupported : gtk_action_group_add_radio_actions : unsupported parameter entries :

// Unsupported : gtk_action_group_add_radio_actions_full : unsupported parameter entries :

// Unsupported : gtk_action_group_add_toggle_actions : unsupported parameter entries :

// Unsupported : gtk_action_group_add_toggle_actions_full : unsupported parameter entries :

// Looks up an action in the action group by name.
/*

C function : gtk_action_group_get_action
*/
func (recv *ActionGroup) GetAction(actionName string) *Action {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.gtk_action_group_get_action((*C.GtkActionGroup)(recv.native), c_action_name)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the name of the action group.
/*

C function : gtk_action_group_get_name
*/
func (recv *ActionGroup) GetName() string {
	retC := C.gtk_action_group_get_name((*C.GtkActionGroup)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns %TRUE if the group is sensitive.  The constituent actions
// can only be logically sensitive (see gtk_action_is_sensitive()) if
// they are sensitive (see gtk_action_get_sensitive()) and their group
// is sensitive.
/*

C function : gtk_action_group_get_sensitive
*/
func (recv *ActionGroup) GetSensitive() bool {
	retC := C.gtk_action_group_get_sensitive((*C.GtkActionGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if the group is visible.  The constituent actions
// can only be logically visible (see gtk_action_is_visible()) if
// they are visible (see gtk_action_get_visible()) and their group
// is visible.
/*

C function : gtk_action_group_get_visible
*/
func (recv *ActionGroup) GetVisible() bool {
	retC := C.gtk_action_group_get_visible((*C.GtkActionGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Lists the actions in the action group.
/*

C function : gtk_action_group_list_actions
*/
func (recv *ActionGroup) ListActions() *glib.List {
	retC := C.gtk_action_group_list_actions((*C.GtkActionGroup)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes an action object from the action group.
/*

C function : gtk_action_group_remove_action
*/
func (recv *ActionGroup) RemoveAction(action *Action) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	C.gtk_action_group_remove_action((*C.GtkActionGroup)(recv.native), c_action)

	return
}

// Changes the sensitivity of @action_group
/*

C function : gtk_action_group_set_sensitive
*/
func (recv *ActionGroup) SetSensitive(sensitive bool) {
	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_action_group_set_sensitive((*C.GtkActionGroup)(recv.native), c_sensitive)

	return
}

// Unsupported : gtk_action_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GtkTranslateFunc) for param func

// Sets the translation domain and uses g_dgettext() for translating the
// @label and @tooltip of #GtkActionEntrys added by
// gtk_action_group_add_actions().
//
// If you’re not using gettext() for localization, see
// gtk_action_group_set_translate_func().
/*

C function : gtk_action_group_set_translation_domain
*/
func (recv *ActionGroup) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.gtk_action_group_set_translation_domain((*C.GtkActionGroup)(recv.native), c_domain)

	return
}

// Changes the visible of @action_group.
/*

C function : gtk_action_group_set_visible
*/
func (recv *ActionGroup) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_action_group_set_visible((*C.GtkActionGroup)(recv.native), c_visible)

	return
}

// Gets the padding on the different sides of the widget.
// See gtk_alignment_set_padding ().
/*

C function : gtk_alignment_get_padding
*/
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

// Sets the padding on the different sides of the widget.
// The padding adds blank space to the sides of the widget. For instance,
// this can be used to indent the child widget towards the right by adding
// padding on the left.
/*

C function : gtk_alignment_set_padding
*/
func (recv *Alignment) SetPadding(paddingTop uint32, paddingBottom uint32, paddingLeft uint32, paddingRight uint32) {
	c_padding_top := (C.guint)(paddingTop)

	c_padding_bottom := (C.guint)(paddingBottom)

	c_padding_left := (C.guint)(paddingLeft)

	c_padding_right := (C.guint)(paddingRight)

	C.gtk_alignment_set_padding((*C.GtkAlignment)(recv.native), c_padding_top, c_padding_bottom, c_padding_left, c_padding_right)

	return
}

// Gets the alignment of the child in the button.
/*

C function : gtk_button_get_alignment
*/
func (recv *Button) GetAlignment() (float32, float32) {
	var c_xalign C.gfloat

	var c_yalign C.gfloat

	C.gtk_button_get_alignment((*C.GtkButton)(recv.native), &c_xalign, &c_yalign)

	xalign := (float32)(c_xalign)

	yalign := (float32)(c_yalign)

	return xalign, yalign
}

// Returns whether the button grabs focus when it is clicked with the mouse.
// See gtk_button_set_focus_on_click().
/*

C function : gtk_button_get_focus_on_click
*/
func (recv *Button) GetFocusOnClick() bool {
	retC := C.gtk_button_get_focus_on_click((*C.GtkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the alignment of the child. This property has no effect unless
// the child is a #GtkMisc or a #GtkAlignment.
/*

C function : gtk_button_set_alignment
*/
func (recv *Button) SetAlignment(xalign float32, yalign float32) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_button_set_alignment((*C.GtkButton)(recv.native), c_xalign, c_yalign)

	return
}

// Sets whether the button will grab focus when it is clicked with the mouse.
// Making mouse clicks not grab focus is useful in places like toolbars where
// you don’t want the keyboard focus removed from the main area of the
// application.
/*

C function : gtk_button_set_focus_on_click
*/
func (recv *Button) SetFocusOnClick(focusOnClick bool) {
	c_focus_on_click :=
		boolToGboolean(focusOnClick)

	C.gtk_button_set_focus_on_click((*C.GtkButton)(recv.native), c_focus_on_click)

	return
}

// Returns whether @child should appear in a secondary group of children.
/*

C function : gtk_button_box_get_child_secondary
*/
func (recv *ButtonBox) GetChildSecondary(child *Widget) bool {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	retC := C.gtk_button_box_get_child_secondary((*C.GtkButtonBox)(recv.native), c_child)
	retGo := retC == C.TRUE

	return retGo
}

// Returns the current display options of @calendar.
/*

C function : gtk_calendar_get_display_options
*/
func (recv *Calendar) GetDisplayOptions() CalendarDisplayOptions {
	retC := C.gtk_calendar_get_display_options((*C.GtkCalendar)(recv.native))
	retGo := (CalendarDisplayOptions)(retC)

	return retGo
}

// Sets display options (whether to display the heading and the month
// headings).
/*

C function : gtk_calendar_set_display_options
*/
func (recv *Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	c_flags := (C.GtkCalendarDisplayOptions)(flags)

	C.gtk_calendar_set_display_options((*C.GtkCalendar)(recv.native), c_flags)

	return
}

type signalCellRendererEditingCanceledDetail struct {
	callback  CellRendererSignalEditingCanceledCallback
	handlerID C.gulong
}

var signalCellRendererEditingCanceledId int
var signalCellRendererEditingCanceledMap = make(map[int]signalCellRendererEditingCanceledDetail)
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
	index := int(uintptr(data))
	callback := signalCellRendererEditingCanceledMap[index].callback
	callback()
}

// Returns whether @check_menu_item looks like a #GtkRadioMenuItem
/*

C function : gtk_check_menu_item_get_draw_as_radio
*/
func (recv *CheckMenuItem) GetDrawAsRadio() bool {
	retC := C.gtk_check_menu_item_get_draw_as_radio((*C.GtkCheckMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether @check_menu_item is drawn like a #GtkRadioMenuItem
/*

C function : gtk_check_menu_item_set_draw_as_radio
*/
func (recv *CheckMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	c_draw_as_radio :=
		boolToGboolean(drawAsRadio)

	C.gtk_check_menu_item_set_draw_as_radio((*C.GtkCheckMenuItem)(recv.native), c_draw_as_radio)

	return
}

// Unsupported : gtk_clipboard_request_targets : unsupported parameter callback : no type generator for ClipboardTargetsReceivedFunc (GtkClipboardTargetsReceivedFunc) for param callback

// Unsupported : gtk_clipboard_wait_for_targets : unsupported parameter targets : output array param targets

type signalColorButtonColorSetDetail struct {
	callback  ColorButtonSignalColorSetCallback
	handlerID C.gulong
}

var signalColorButtonColorSetId int
var signalColorButtonColorSetMap = make(map[int]signalColorButtonColorSetDetail)
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
	index := int(uintptr(data))
	callback := signalColorButtonColorSetMap[index].callback
	callback()
}

// Creates a new color button.
//
// This returns a widget in the form of a small button containing
// a swatch representing the current selected color. When the button
// is clicked, a color-selection dialog will open, allowing the user
// to select a color. The swatch will be updated to reflect the new
// color when the user finishes.
/*

C function : gtk_color_button_new
*/
func ColorButtonNew() *ColorButton {
	retC := C.gtk_color_button_new()
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new color button.
/*

C function : gtk_color_button_new_with_color
*/
func ColorButtonNewWithColor(color *gdk.Color) *ColorButton {
	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	retC := C.gtk_color_button_new_with_color(c_color)
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the current alpha value.
/*

C function : gtk_color_button_get_alpha
*/
func (recv *ColorButton) GetAlpha() uint16 {
	retC := C.gtk_color_button_get_alpha((*C.GtkColorButton)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Sets @color to be the current color in the #GtkColorButton widget.
/*

C function : gtk_color_button_get_color
*/
func (recv *ColorButton) GetColor() *gdk.Color {
	var c_color C.GdkColor

	C.gtk_color_button_get_color((*C.GtkColorButton)(recv.native), &c_color)

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return color
}

// Gets the title of the color selection dialog.
/*

C function : gtk_color_button_get_title
*/
func (recv *ColorButton) GetTitle() string {
	retC := C.gtk_color_button_get_title((*C.GtkColorButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Does the color selection dialog use the alpha channel ?
/*

C function : gtk_color_button_get_use_alpha
*/
func (recv *ColorButton) GetUseAlpha() bool {
	retC := C.gtk_color_button_get_use_alpha((*C.GtkColorButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the current opacity to be @alpha.
/*

C function : gtk_color_button_set_alpha
*/
func (recv *ColorButton) SetAlpha(alpha uint16) {
	c_alpha := (C.guint16)(alpha)

	C.gtk_color_button_set_alpha((*C.GtkColorButton)(recv.native), c_alpha)

	return
}

// Sets the current color to be @color.
/*

C function : gtk_color_button_set_color
*/
func (recv *ColorButton) SetColor(color *gdk.Color) {
	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	C.gtk_color_button_set_color((*C.GtkColorButton)(recv.native), c_color)

	return
}

// Sets the title for the color selection dialog.
/*

C function : gtk_color_button_set_title
*/
func (recv *ColorButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_color_button_set_title((*C.GtkColorButton)(recv.native), c_title)

	return
}

// Sets whether or not the color button should use the alpha channel.
/*

C function : gtk_color_button_set_use_alpha
*/
func (recv *ColorButton) SetUseAlpha(useAlpha bool) {
	c_use_alpha :=
		boolToGboolean(useAlpha)

	C.gtk_color_button_set_use_alpha((*C.GtkColorButton)(recv.native), c_use_alpha)

	return
}

type signalComboBoxChangedDetail struct {
	callback  ComboBoxSignalChangedCallback
	handlerID C.gulong
}

var signalComboBoxChangedId int
var signalComboBoxChangedMap = make(map[int]signalComboBoxChangedDetail)
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
	index := int(uintptr(data))
	callback := signalComboBoxChangedMap[index].callback
	callback()
}

// Creates a new empty #GtkComboBox.
/*

C function : gtk_combo_box_new
*/
func ComboBoxNew() *ComboBox {
	retC := C.gtk_combo_box_new()
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkComboBox with the model initialized to @model.
/*

C function : gtk_combo_box_new_with_model
*/
func ComboBoxNewWithModel(model *TreeModel) *ComboBox {
	c_model := (*C.GtkTreeModel)(model.ToC())

	retC := C.gtk_combo_box_new_with_model(c_model)
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the index of the currently active item, or -1 if there’s no
// active item. If the model is a non-flat treemodel, and the active item
// is not an immediate child of the root of the tree, this function returns
// `gtk_tree_path_get_indices (path)[0]`, where
// `path` is the #GtkTreePath of the active item.
/*

C function : gtk_combo_box_get_active
*/
func (recv *ComboBox) GetActive() int32 {
	retC := C.gtk_combo_box_get_active((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets @iter to point to the currently active item, if any item is active.
// Otherwise, @iter is left unchanged.
/*

C function : gtk_combo_box_get_active_iter
*/
func (recv *ComboBox) GetActiveIter() (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	retC := C.gtk_combo_box_get_active_iter((*C.GtkComboBox)(recv.native), &c_iter)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// Returns the #GtkTreeModel which is acting as data source for @combo_box.
/*

C function : gtk_combo_box_get_model
*/
func (recv *ComboBox) GetModel() *TreeModel {
	retC := C.gtk_combo_box_get_model((*C.GtkComboBox)(recv.native))
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Hides the menu or dropdown list of @combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
/*

C function : gtk_combo_box_popdown
*/
func (recv *ComboBox) Popdown() {
	C.gtk_combo_box_popdown((*C.GtkComboBox)(recv.native))

	return
}

// Pops up the menu or dropdown list of @combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, @combo_box must be mapped, or nothing will happen.
/*

C function : gtk_combo_box_popup
*/
func (recv *ComboBox) Popup() {
	C.gtk_combo_box_popup((*C.GtkComboBox)(recv.native))

	return
}

// Sets the active item of @combo_box to be the item at @index.
/*

C function : gtk_combo_box_set_active
*/
func (recv *ComboBox) SetActive(index int32) {
	c_index_ := (C.gint)(index)

	C.gtk_combo_box_set_active((*C.GtkComboBox)(recv.native), c_index_)

	return
}

// Sets the current active item to be the one referenced by @iter, or
// unsets the active item if @iter is %NULL.
/*

C function : gtk_combo_box_set_active_iter
*/
func (recv *ComboBox) SetActiveIter(iter *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_combo_box_set_active_iter((*C.GtkComboBox)(recv.native), c_iter)

	return
}

// Sets the column with column span information for @combo_box to be
// @column_span. The column span column contains integers which indicate
// how many columns an item should span.
/*

C function : gtk_combo_box_set_column_span_column
*/
func (recv *ComboBox) SetColumnSpanColumn(columnSpan int32) {
	c_column_span := (C.gint)(columnSpan)

	C.gtk_combo_box_set_column_span_column((*C.GtkComboBox)(recv.native), c_column_span)

	return
}

// Sets the model used by @combo_box to be @model. Will unset a previously set
// model (if applicable). If model is %NULL, then it will unset the model.
//
// Note that this function does not clear the cell renderers, you have to
// call gtk_cell_layout_clear() yourself if you need to set up different
// cell renderers for the new model.
/*

C function : gtk_combo_box_set_model
*/
func (recv *ComboBox) SetModel(model *TreeModel) {
	c_model := (*C.GtkTreeModel)(model.ToC())

	C.gtk_combo_box_set_model((*C.GtkComboBox)(recv.native), c_model)

	return
}

// Sets the column with row span information for @combo_box to be @row_span.
// The row span column contains integers which indicate how many rows
// an item should span.
/*

C function : gtk_combo_box_set_row_span_column
*/
func (recv *ComboBox) SetRowSpanColumn(rowSpan int32) {
	c_row_span := (C.gint)(rowSpan)

	C.gtk_combo_box_set_row_span_column((*C.GtkComboBox)(recv.native), c_row_span)

	return
}

// Sets the wrap width of @combo_box to be @width. The wrap width is basically
// the preferred number of columns when you want the popup to be layed out
// in a table.
/*

C function : gtk_combo_box_set_wrap_width
*/
func (recv *ComboBox) SetWrapWidth(width int32) {
	c_width := (C.gint)(width)

	C.gtk_combo_box_set_wrap_width((*C.GtkComboBox)(recv.native), c_width)

	return
}

// Gets the value set by gtk_entry_set_alignment().
/*

C function : gtk_entry_get_alignment
*/
func (recv *Entry) GetAlignment() float32 {
	retC := C.gtk_entry_get_alignment((*C.GtkEntry)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// Returns the auxiliary completion object currently in use by @entry.
/*

C function : gtk_entry_get_completion
*/
func (recv *Entry) GetCompletion() *EntryCompletion {
	retC := C.gtk_entry_get_completion((*C.GtkEntry)(recv.native))
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the alignment for the contents of the entry. This controls
// the horizontal positioning of the contents when the displayed
// text is shorter than the width of the entry.
/*

C function : gtk_entry_set_alignment
*/
func (recv *Entry) SetAlignment(xalign float32) {
	c_xalign := (C.gfloat)(xalign)

	C.gtk_entry_set_alignment((*C.GtkEntry)(recv.native), c_xalign)

	return
}

// Sets @completion to be the auxiliary completion object to use with @entry.
// All further configuration of the completion mechanism is done on
// @completion using the #GtkEntryCompletion API. Completion is disabled if
// @completion is set to %NULL.
/*

C function : gtk_entry_set_completion
*/
func (recv *Entry) SetCompletion(completion *EntryCompletion) {
	c_completion := (*C.GtkEntryCompletion)(C.NULL)
	if completion != nil {
		c_completion = (*C.GtkEntryCompletion)(completion.ToC())
	}

	C.gtk_entry_set_completion((*C.GtkEntry)(recv.native), c_completion)

	return
}

// Unsupported signal 'action-activated' for EntryCompletion : unsupported parameter index : type gint :

type signalEntryCompletionMatchSelectedDetail struct {
	callback  EntryCompletionSignalMatchSelectedCallback
	handlerID C.gulong
}

var signalEntryCompletionMatchSelectedId int
var signalEntryCompletionMatchSelectedMap = make(map[int]signalEntryCompletionMatchSelectedDetail)
var signalEntryCompletionMatchSelectedLock sync.Mutex

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
	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalEntryCompletionMatchSelectedMap[index].callback
	retGo := callback(model, iter)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Creates a new #GtkEntryCompletion object.
/*

C function : gtk_entry_completion_new
*/
func EntryCompletionNew() *EntryCompletion {
	retC := C.gtk_entry_completion_new()
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Requests a completion operation, or in other words a refiltering of the
// current list with completions, using the current key. The completion list
// view will be updated accordingly.
/*

C function : gtk_entry_completion_complete
*/
func (recv *EntryCompletion) Complete() {
	C.gtk_entry_completion_complete((*C.GtkEntryCompletion)(recv.native))

	return
}

// Deletes the action at @index_ from @completion’s action list.
//
// Note that @index_ is a relative position and the position of an
// action may have changed since it was inserted.
/*

C function : gtk_entry_completion_delete_action
*/
func (recv *EntryCompletion) DeleteAction(index int32) {
	c_index_ := (C.gint)(index)

	C.gtk_entry_completion_delete_action((*C.GtkEntryCompletion)(recv.native), c_index_)

	return
}

// Gets the entry @completion has been attached to.
/*

C function : gtk_entry_completion_get_entry
*/
func (recv *EntryCompletion) GetEntry() *Widget {
	retC := C.gtk_entry_completion_get_entry((*C.GtkEntryCompletion)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the minimum key length as set for @completion.
/*

C function : gtk_entry_completion_get_minimum_key_length
*/
func (recv *EntryCompletion) GetMinimumKeyLength() int32 {
	retC := C.gtk_entry_completion_get_minimum_key_length((*C.GtkEntryCompletion)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the model the #GtkEntryCompletion is using as data source.
// Returns %NULL if the model is unset.
/*

C function : gtk_entry_completion_get_model
*/
func (recv *EntryCompletion) GetModel() *TreeModel {
	retC := C.gtk_entry_completion_get_model((*C.GtkEntryCompletion)(recv.native))
	var retGo (*TreeModel)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreeModelNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Inserts an action in @completion’s action item list at position @index_
// with markup @markup.
/*

C function : gtk_entry_completion_insert_action_markup
*/
func (recv *EntryCompletion) InsertActionMarkup(index int32, markup string) {
	c_index_ := (C.gint)(index)

	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_entry_completion_insert_action_markup((*C.GtkEntryCompletion)(recv.native), c_index_, c_markup)

	return
}

// Inserts an action in @completion’s action item list at position @index_
// with text @text. If you want the action item to have markup, use
// gtk_entry_completion_insert_action_markup().
//
// Note that @index_ is a relative position in the list of actions and
// the position of an action can change when deleting a different action.
/*

C function : gtk_entry_completion_insert_action_text
*/
func (recv *EntryCompletion) InsertActionText(index int32, text string) {
	c_index_ := (C.gint)(index)

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_entry_completion_insert_action_text((*C.GtkEntryCompletion)(recv.native), c_index_, c_text)

	return
}

// Unsupported : gtk_entry_completion_set_match_func : unsupported parameter func : no type generator for EntryCompletionMatchFunc (GtkEntryCompletionMatchFunc) for param func

// Requires the length of the search key for @completion to be at least
// @length. This is useful for long lists, where completing using a small
// key takes a lot of time and will come up with meaningless results anyway
// (ie, a too large dataset).
/*

C function : gtk_entry_completion_set_minimum_key_length
*/
func (recv *EntryCompletion) SetMinimumKeyLength(length int32) {
	c_length := (C.gint)(length)

	C.gtk_entry_completion_set_minimum_key_length((*C.GtkEntryCompletion)(recv.native), c_length)

	return
}

// Sets the model for a #GtkEntryCompletion. If @completion already has
// a model set, it will remove it before setting the new model.
// If model is %NULL, then it will unset the model.
/*

C function : gtk_entry_completion_set_model
*/
func (recv *EntryCompletion) SetModel(model *TreeModel) {
	c_model := (*C.GtkTreeModel)(model.ToC())

	C.gtk_entry_completion_set_model((*C.GtkEntryCompletion)(recv.native), c_model)

	return
}

// Convenience function for setting up the most used case of this code: a
// completion list with just strings. This function will set up @completion
// to have a list displaying all (and just) strings in the completion list,
// and to get those strings from @column in the model of @completion.
//
// This functions creates and adds a #GtkCellRendererText for the selected
// column. If you need to set the text column, but don't want the cell
// renderer, use g_object_set() to set the #GtkEntryCompletion:text-column
// property directly.
/*

C function : gtk_entry_completion_set_text_column
*/
func (recv *EntryCompletion) SetTextColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_entry_completion_set_text_column((*C.GtkEntryCompletion)(recv.native), c_column)

	return
}

// Returns whether the event box window is above or below the
// windows of its child. See gtk_event_box_set_above_child()
// for details.
/*

C function : gtk_event_box_get_above_child
*/
func (recv *EventBox) GetAboveChild() bool {
	retC := C.gtk_event_box_get_above_child((*C.GtkEventBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the event box has a visible window.
// See gtk_event_box_set_visible_window() for details.
/*

C function : gtk_event_box_get_visible_window
*/
func (recv *EventBox) GetVisibleWindow() bool {
	retC := C.gtk_event_box_get_visible_window((*C.GtkEventBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Set whether the event box window is positioned above the windows
// of its child, as opposed to below it. If the window is above, all
// events inside the event box will go to the event box. If the window
// is below, events in windows of child widgets will first got to that
// widget, and then to its parents.
//
// The default is to keep the window below the child.
/*

C function : gtk_event_box_set_above_child
*/
func (recv *EventBox) SetAboveChild(aboveChild bool) {
	c_above_child :=
		boolToGboolean(aboveChild)

	C.gtk_event_box_set_above_child((*C.GtkEventBox)(recv.native), c_above_child)

	return
}

// Set whether the event box uses a visible or invisible child
// window. The default is to use visible windows.
//
// In an invisible window event box, the window that the
// event box creates is a %GDK_INPUT_ONLY window, which
// means that it is invisible and only serves to receive
// events.
//
// A visible window event box creates a visible (%GDK_INPUT_OUTPUT)
// window that acts as the parent window for all the widgets
// contained in the event box.
//
// You should generally make your event box invisible if
// you just want to trap events. Creating a visible window
// may cause artifacts that are visible to the user, especially
// if the user is using a theme with gradients or pixmaps.
//
// The main reason to create a non input-only event box is if
// you want to set the background to a different color or
// draw on it.
//
// There is one unexpected issue for an invisible event box that has its
// window below the child. (See gtk_event_box_set_above_child().)
// Since the input-only window is not an ancestor window of any windows
// that descendent widgets of the event box create, events on these
// windows aren’t propagated up by the windowing system, but only by GTK+.
// The practical effect of this is if an event isn’t in the event
// mask for the descendant window (see gtk_widget_add_events()),
// it won’t be received by the event box.
//
// This problem doesn’t occur for visible event boxes, because in
// that case, the event box window is actually the ancestor of the
// descendant windows, not just at the same place on the screen.
/*

C function : gtk_event_box_set_visible_window
*/
func (recv *EventBox) SetVisibleWindow(visibleWindow bool) {
	c_visible_window :=
		boolToGboolean(visibleWindow)

	C.gtk_event_box_set_visible_window((*C.GtkEventBox)(recv.native), c_visible_window)

	return
}

// Creates a new expander using @label as the text of the label.
/*

C function : gtk_expander_new
*/
func ExpanderNew(label string) *Expander {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_expander_new(c_label)
	retGo := ExpanderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new expander using @label as the text of the label.
// If characters in @label are preceded by an underscore, they are underlined.
// If you need a literal underscore character in a label, use “__” (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic.
// Pressing Alt and that key activates the button.
/*

C function : gtk_expander_new_with_mnemonic
*/
func ExpanderNewWithMnemonic(label string) *Expander {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_expander_new_with_mnemonic(c_label)
	retGo := ExpanderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Queries a #GtkExpander and returns its current state. Returns %TRUE
// if the child widget is revealed.
//
// See gtk_expander_set_expanded().
/*

C function : gtk_expander_get_expanded
*/
func (recv *Expander) GetExpanded() bool {
	retC := C.gtk_expander_get_expanded((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Fetches the text from a label widget including any embedded
// underlines indicating mnemonics and Pango markup, as set by
// gtk_expander_set_label(). If the label text has not been set the
// return value will be %NULL. This will be the case if you create an
// empty button with gtk_button_new() to use as a container.
//
// Note that this function behaved differently in versions prior to
// 2.14 and used to return the label text stripped of embedded
// underlines indicating mnemonics and Pango markup. This problem can
// be avoided by fetching the label text directly from the label
// widget.
/*

C function : gtk_expander_get_label
*/
func (recv *Expander) GetLabel() string {
	retC := C.gtk_expander_get_label((*C.GtkExpander)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the label widget for the frame. See
// gtk_expander_set_label_widget().
/*

C function : gtk_expander_get_label_widget
*/
func (recv *Expander) GetLabelWidget() *Widget {
	retC := C.gtk_expander_get_label_widget((*C.GtkExpander)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the value set by gtk_expander_set_spacing().
/*

C function : gtk_expander_get_spacing
*/
func (recv *Expander) GetSpacing() int32 {
	retC := C.gtk_expander_get_spacing((*C.GtkExpander)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns whether the label’s text is interpreted as marked up with
// the [Pango text markup language][PangoMarkupFormat].
// See gtk_expander_set_use_markup().
/*

C function : gtk_expander_get_use_markup
*/
func (recv *Expander) GetUseMarkup() bool {
	retC := C.gtk_expander_get_use_markup((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether an embedded underline in the expander label
// indicates a mnemonic. See gtk_expander_set_use_underline().
/*

C function : gtk_expander_get_use_underline
*/
func (recv *Expander) GetUseUnderline() bool {
	retC := C.gtk_expander_get_use_underline((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the state of the expander. Set to %TRUE, if you want
// the child widget to be revealed, and %FALSE if you want the
// child widget to be hidden.
/*

C function : gtk_expander_set_expanded
*/
func (recv *Expander) SetExpanded(expanded bool) {
	c_expanded :=
		boolToGboolean(expanded)

	C.gtk_expander_set_expanded((*C.GtkExpander)(recv.native), c_expanded)

	return
}

// Sets the text of the label of the expander to @label.
//
// This will also clear any previously set labels.
/*

C function : gtk_expander_set_label
*/
func (recv *Expander) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_expander_set_label((*C.GtkExpander)(recv.native), c_label)

	return
}

// Set the label widget for the expander. This is the widget
// that will appear embedded alongside the expander arrow.
/*

C function : gtk_expander_set_label_widget
*/
func (recv *Expander) SetLabelWidget(labelWidget *Widget) {
	c_label_widget := (*C.GtkWidget)(C.NULL)
	if labelWidget != nil {
		c_label_widget = (*C.GtkWidget)(labelWidget.ToC())
	}

	C.gtk_expander_set_label_widget((*C.GtkExpander)(recv.native), c_label_widget)

	return
}

// Sets the spacing field of @expander, which is the number of
// pixels to place between expander and the child.
/*

C function : gtk_expander_set_spacing
*/
func (recv *Expander) SetSpacing(spacing int32) {
	c_spacing := (C.gint)(spacing)

	C.gtk_expander_set_spacing((*C.GtkExpander)(recv.native), c_spacing)

	return
}

// Sets whether the text of the label contains markup in
// [Pango’s text markup language][PangoMarkupFormat].
// See gtk_label_set_markup().
/*

C function : gtk_expander_set_use_markup
*/
func (recv *Expander) SetUseMarkup(useMarkup bool) {
	c_use_markup :=
		boolToGboolean(useMarkup)

	C.gtk_expander_set_use_markup((*C.GtkExpander)(recv.native), c_use_markup)

	return
}

// If true, an underline in the text of the expander label indicates
// the next character should be used for the mnemonic accelerator key.
/*

C function : gtk_expander_set_use_underline
*/
func (recv *Expander) SetUseUnderline(useUnderline bool) {
	c_use_underline :=
		boolToGboolean(useUnderline)

	C.gtk_expander_set_use_underline((*C.GtkExpander)(recv.native), c_use_underline)

	return
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Creates a new #GtkFileChooserWidget. This is a file chooser widget that can
// be embedded in custom windows, and it is the same widget that is used by
// #GtkFileChooserDialog.
/*

C function : gtk_file_chooser_widget_new
*/
func FileChooserWidgetNew(action FileChooserAction) *FileChooserWidget {
	c_action := (C.GtkFileChooserAction)(action)

	retC := C.gtk_file_chooser_widget_new(c_action)
	retGo := FileChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkFileFilter with no rules added to it.
// Such a filter doesn’t accept any files, so is not
// particularly useful until you add rules with
// gtk_file_filter_add_mime_type(), gtk_file_filter_add_pattern(),
// or gtk_file_filter_add_custom(). To create a filter
// that accepts any file, use:
// |[<!-- language="C" -->
// GtkFileFilter *filter = gtk_file_filter_new ();
// gtk_file_filter_add_pattern (filter, "*");
// ]|
/*

C function : gtk_file_filter_new
*/
func FileFilterNew() *FileFilter {
	retC := C.gtk_file_filter_new()
	retGo := FileFilterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_file_filter_add_custom : unsupported parameter func : no type generator for FileFilterFunc (GtkFileFilterFunc) for param func

// Adds a rule allowing a given mime type to @filter.
/*

C function : gtk_file_filter_add_mime_type
*/
func (recv *FileFilter) AddMimeType(mimeType string) {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	C.gtk_file_filter_add_mime_type((*C.GtkFileFilter)(recv.native), c_mime_type)

	return
}

// Adds a rule allowing a shell style glob to a filter.
/*

C function : gtk_file_filter_add_pattern
*/
func (recv *FileFilter) AddPattern(pattern string) {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	C.gtk_file_filter_add_pattern((*C.GtkFileFilter)(recv.native), c_pattern)

	return
}

// Tests whether a file should be displayed according to @filter.
// The #GtkFileFilterInfo @filter_info should include
// the fields returned from gtk_file_filter_get_needed().
//
// This function will not typically be used by applications; it
// is intended principally for use in the implementation of
// #GtkFileChooser.
/*

C function : gtk_file_filter_filter
*/
func (recv *FileFilter) Filter(filterInfo *FileFilterInfo) bool {
	c_filter_info := (*C.GtkFileFilterInfo)(C.NULL)
	if filterInfo != nil {
		c_filter_info = (*C.GtkFileFilterInfo)(filterInfo.ToC())
	}

	retC := C.gtk_file_filter_filter((*C.GtkFileFilter)(recv.native), c_filter_info)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the human-readable name for the filter. See gtk_file_filter_set_name().
/*

C function : gtk_file_filter_get_name
*/
func (recv *FileFilter) GetName() string {
	retC := C.gtk_file_filter_get_name((*C.GtkFileFilter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the fields that need to be filled in for the #GtkFileFilterInfo
// passed to gtk_file_filter_filter()
//
// This function will not typically be used by applications; it
// is intended principally for use in the implementation of
// #GtkFileChooser.
/*

C function : gtk_file_filter_get_needed
*/
func (recv *FileFilter) GetNeeded() FileFilterFlags {
	retC := C.gtk_file_filter_get_needed((*C.GtkFileFilter)(recv.native))
	retGo := (FileFilterFlags)(retC)

	return retGo
}

// Sets the human-readable name of the filter; this is the string
// that will be displayed in the file selector user interface if
// there is a selectable list of filters.
/*

C function : gtk_file_filter_set_name
*/
func (recv *FileFilter) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_file_filter_set_name((*C.GtkFileFilter)(recv.native), c_name)

	return
}

type signalFontButtonFontSetDetail struct {
	callback  FontButtonSignalFontSetCallback
	handlerID C.gulong
}

var signalFontButtonFontSetId int
var signalFontButtonFontSetMap = make(map[int]signalFontButtonFontSetDetail)
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
	index := int(uintptr(data))
	callback := signalFontButtonFontSetMap[index].callback
	callback()
}

// Creates a new font picker widget.
/*

C function : gtk_font_button_new
*/
func FontButtonNew() *FontButton {
	retC := C.gtk_font_button_new()
	retGo := FontButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new font picker widget.
/*

C function : gtk_font_button_new_with_font
*/
func FontButtonNewWithFont(fontname string) *FontButton {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_button_new_with_font(c_fontname)
	retGo := FontButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the name of the currently selected font. This name includes
// style and size information as well. If you want to render something
// with the font, use this string with pango_font_description_from_string() .
// If you’re interested in peeking certain values (family name,
// style, size, weight) just query these properties from the
// #PangoFontDescription object.
/*

C function : gtk_font_button_get_font_name
*/
func (recv *FontButton) GetFontName() string {
	retC := C.gtk_font_button_get_font_name((*C.GtkFontButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns whether the font size will be shown in the label.
/*

C function : gtk_font_button_get_show_size
*/
func (recv *FontButton) GetShowSize() bool {
	retC := C.gtk_font_button_get_show_size((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the name of the font style will be shown in the label.
/*

C function : gtk_font_button_get_show_style
*/
func (recv *FontButton) GetShowStyle() bool {
	retC := C.gtk_font_button_get_show_style((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the title of the font chooser dialog.
/*

C function : gtk_font_button_get_title
*/
func (recv *FontButton) GetTitle() string {
	retC := C.gtk_font_button_get_title((*C.GtkFontButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns whether the selected font is used in the label.
/*

C function : gtk_font_button_get_use_font
*/
func (recv *FontButton) GetUseFont() bool {
	retC := C.gtk_font_button_get_use_font((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the selected size is used in the label.
/*

C function : gtk_font_button_get_use_size
*/
func (recv *FontButton) GetUseSize() bool {
	retC := C.gtk_font_button_get_use_size((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets or updates the currently-displayed font in font picker dialog.
/*

C function : gtk_font_button_set_font_name
*/
func (recv *FontButton) SetFontName(fontname string) bool {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_button_set_font_name((*C.GtkFontButton)(recv.native), c_fontname)
	retGo := retC == C.TRUE

	return retGo
}

// If @show_size is %TRUE, the font size will be displayed along with the name of the selected font.
/*

C function : gtk_font_button_set_show_size
*/
func (recv *FontButton) SetShowSize(showSize bool) {
	c_show_size :=
		boolToGboolean(showSize)

	C.gtk_font_button_set_show_size((*C.GtkFontButton)(recv.native), c_show_size)

	return
}

// If @show_style is %TRUE, the font style will be displayed along with name of the selected font.
/*

C function : gtk_font_button_set_show_style
*/
func (recv *FontButton) SetShowStyle(showStyle bool) {
	c_show_style :=
		boolToGboolean(showStyle)

	C.gtk_font_button_set_show_style((*C.GtkFontButton)(recv.native), c_show_style)

	return
}

// Sets the title for the font chooser dialog.
/*

C function : gtk_font_button_set_title
*/
func (recv *FontButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_font_button_set_title((*C.GtkFontButton)(recv.native), c_title)

	return
}

// If @use_font is %TRUE, the font name will be written using the selected font.
/*

C function : gtk_font_button_set_use_font
*/
func (recv *FontButton) SetUseFont(useFont bool) {
	c_use_font :=
		boolToGboolean(useFont)

	C.gtk_font_button_set_use_font((*C.GtkFontButton)(recv.native), c_use_font)

	return
}

// If @use_size is %TRUE, the font name will be written using the selected size.
/*

C function : gtk_font_button_set_use_size
*/
func (recv *FontButton) SetUseSize(useSize bool) {
	c_use_size :=
		boolToGboolean(useSize)

	C.gtk_font_button_set_use_size((*C.GtkFontButton)(recv.native), c_use_size)

	return
}

// Make a copy of a #GtkIconInfo.
/*

C function : gtk_icon_info_copy
*/
func (recv *IconInfo) Copy() *IconInfo {
	retC := C.gtk_icon_info_copy((*C.GtkIconInfo)(recv.native))
	retGo := IconInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free a #GtkIconInfo and associated information
/*

C function : gtk_icon_info_free
*/
func (recv *IconInfo) Free() {
	C.gtk_icon_info_free((*C.GtkIconInfo)(recv.native))

	return
}

// Unsupported : gtk_icon_info_get_attach_points : unsupported parameter points : output array param points

// Gets the base size for the icon. The base size
// is a size for the icon that was specified by
// the icon theme creator. This may be different
// than the actual size of image; an example of
// this is small emblem icons that can be attached
// to a larger icon. These icons will be given
// the same base size as the larger icons to which
// they are attached.
//
// Note that for scaled icons the base size does
// not include the base scale.
/*

C function : gtk_icon_info_get_base_size
*/
func (recv *IconInfo) GetBaseSize() int32 {
	retC := C.gtk_icon_info_get_base_size((*C.GtkIconInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the built-in image for this icon, if any. To allow GTK+ to use
// built in icon images, you must pass the %GTK_ICON_LOOKUP_USE_BUILTIN
// to gtk_icon_theme_lookup_icon().
/*

C function : gtk_icon_info_get_builtin_pixbuf
*/
func (recv *IconInfo) GetBuiltinPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_icon_info_get_builtin_pixbuf((*C.GtkIconInfo)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// This function is deprecated and always returns %NULL.
/*

C function : gtk_icon_info_get_display_name
*/
func (recv *IconInfo) GetDisplayName() string {
	retC := C.gtk_icon_info_get_display_name((*C.GtkIconInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_icon_info_get_embedded_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Gets the filename for the icon. If the %GTK_ICON_LOOKUP_USE_BUILTIN
// flag was passed to gtk_icon_theme_lookup_icon(), there may be no
// filename if a builtin icon is returned; in this case, you should
// use gtk_icon_info_get_builtin_pixbuf().
/*

C function : gtk_icon_info_get_filename
*/
func (recv *IconInfo) GetFilename() string {
	retC := C.gtk_icon_info_get_filename((*C.GtkIconInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Renders an icon previously looked up in an icon theme using
// gtk_icon_theme_lookup_icon(); the size will be based on the size
// passed to gtk_icon_theme_lookup_icon(). Note that the resulting
// pixbuf may not be exactly this size; an icon theme may have icons
// that differ slightly from their nominal sizes, and in addition GTK+
// will avoid scaling icons that it considers sufficiently close to the
// requested size or for which the source image would have to be scaled
// up too far. (This maintains sharpness.). This behaviour can be changed
// by passing the %GTK_ICON_LOOKUP_FORCE_SIZE flag when obtaining
// the #GtkIconInfo. If this flag has been specified, the pixbuf
// returned by this function will be scaled to the exact size.
/*

C function : gtk_icon_info_load_icon
*/
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

// Sets whether the coordinates returned by gtk_icon_info_get_embedded_rect()
// and gtk_icon_info_get_attach_points() should be returned in their
// original form as specified in the icon theme, instead of scaled
// appropriately for the pixbuf returned by gtk_icon_info_load_icon().
//
// Raw coordinates are somewhat strange; they are specified to be with
// respect to the unscaled pixmap for PNG and XPM icons, but for SVG
// icons, they are in a 1000x1000 coordinate space that is scaled
// to the final size of the icon.  You can determine if the icon is an SVG
// icon by using gtk_icon_info_get_filename(), and seeing if it is non-%NULL
// and ends in “.svg”.
//
// This function is provided primarily to allow compatibility wrappers
// for older API's, and is not expected to be useful for applications.
/*

C function : gtk_icon_info_set_raw_coordinates
*/
func (recv *IconInfo) SetRawCoordinates(rawCoordinates bool) {
	c_raw_coordinates :=
		boolToGboolean(rawCoordinates)

	C.gtk_icon_info_set_raw_coordinates((*C.GtkIconInfo)(recv.native), c_raw_coordinates)

	return
}

// Creates a new icon theme object. Icon theme objects are used
// to lookup up an icon by name in a particular icon theme.
// Usually, you’ll want to use gtk_icon_theme_get_default()
// or gtk_icon_theme_get_for_screen() rather than creating
// a new icon theme object for scratch.
/*

C function : gtk_icon_theme_new
*/
func IconThemeNew() *IconTheme {
	retC := C.gtk_icon_theme_new()
	retGo := IconThemeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Appends a directory to the search path.
// See gtk_icon_theme_set_search_path().
/*

C function : gtk_icon_theme_append_search_path
*/
func (recv *IconTheme) AppendSearchPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.gtk_icon_theme_append_search_path((*C.GtkIconTheme)(recv.native), c_path)

	return
}

// Gets the name of an icon that is representative of the
// current theme (for instance, to use when presenting
// a list of themes to the user.)
/*

C function : gtk_icon_theme_get_example_icon_name
*/
func (recv *IconTheme) GetExampleIconName() string {
	retC := C.gtk_icon_theme_get_example_icon_name((*C.GtkIconTheme)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_theme_get_search_path : unsupported parameter path : output array param path

// Checks whether an icon theme includes an icon
// for a particular name.
/*

C function : gtk_icon_theme_has_icon
*/
func (recv *IconTheme) HasIcon(iconName string) bool {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	retC := C.gtk_icon_theme_has_icon((*C.GtkIconTheme)(recv.native), c_icon_name)
	retGo := retC == C.TRUE

	return retGo
}

// Lists the icons in the current icon theme. Only a subset
// of the icons can be listed by providing a context string.
// The set of values for the context string is system dependent,
// but will typically include such values as “Applications” and
// “MimeTypes”. Contexts are explained in the
// [Icon Theme Specification](http://www.freedesktop.org/wiki/Specifications/icon-theme-spec).
// The standard contexts are listed in the
// [Icon Naming Specification](http://www.freedesktop.org/wiki/Specifications/icon-naming-spec).
// Also see gtk_icon_theme_list_contexts().
/*

C function : gtk_icon_theme_list_icons
*/
func (recv *IconTheme) ListIcons(context string) *glib.List {
	c_context := C.CString(context)
	defer C.free(unsafe.Pointer(c_context))

	retC := C.gtk_icon_theme_list_icons((*C.GtkIconTheme)(recv.native), c_context)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Looks up an icon in an icon theme, scales it to the given size
// and renders it into a pixbuf. This is a convenience function;
// if more details about the icon are needed, use
// gtk_icon_theme_lookup_icon() followed by gtk_icon_info_load_icon().
//
// Note that you probably want to listen for icon theme changes and
// update the icon. This is usually done by connecting to the
// GtkWidget::style-set signal. If for some reason you do not want to
// update the icon when the icon theme changes, you should consider
// using gdk_pixbuf_copy() to make a private copy of the pixbuf
// returned by this function. Otherwise GTK+ may need to keep the old
// icon theme loaded, which would be a waste of memory.
/*

C function : gtk_icon_theme_load_icon
*/
func (recv *IconTheme) LoadIcon(iconName string, size int32, flags IconLookupFlags) (*gdkpixbuf.Pixbuf, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_icon((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_flags, &cThrowableError)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Looks up a named icon and returns a #GtkIconInfo containing
// information such as the filename of the icon. The icon
// can then be rendered into a pixbuf using
// gtk_icon_info_load_icon(). (gtk_icon_theme_load_icon()
// combines these two steps if all you need is the pixbuf.)
//
// When rendering on displays with high pixel densities you should not
// use a @size multiplied by the scaling factor returned by functions
// like gdk_window_get_scale_factor(). Instead, you should use
// gtk_icon_theme_lookup_icon_for_scale(), as the assets loaded
// for a given scaling factor may be different.
/*

C function : gtk_icon_theme_lookup_icon
*/
func (recv *IconTheme) LookupIcon(iconName string, size int32, flags IconLookupFlags) *IconInfo {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_icon((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_flags)
	var retGo (*IconInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Prepends a directory to the search path.
// See gtk_icon_theme_set_search_path().
/*

C function : gtk_icon_theme_prepend_search_path
*/
func (recv *IconTheme) PrependSearchPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.gtk_icon_theme_prepend_search_path((*C.GtkIconTheme)(recv.native), c_path)

	return
}

// Checks to see if the icon theme has changed; if it has, any
// currently cached information is discarded and will be reloaded
// next time @icon_theme is accessed.
/*

C function : gtk_icon_theme_rescan_if_needed
*/
func (recv *IconTheme) RescanIfNeeded() bool {
	retC := C.gtk_icon_theme_rescan_if_needed((*C.GtkIconTheme)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the name of the icon theme that the #GtkIconTheme object uses
// overriding system configuration. This function cannot be called
// on the icon theme objects returned from gtk_icon_theme_get_default()
// and gtk_icon_theme_get_for_screen().
/*

C function : gtk_icon_theme_set_custom_theme
*/
func (recv *IconTheme) SetCustomTheme(themeName string) {
	c_theme_name := C.CString(themeName)
	defer C.free(unsafe.Pointer(c_theme_name))

	C.gtk_icon_theme_set_custom_theme((*C.GtkIconTheme)(recv.native), c_theme_name)

	return
}

// Sets the screen for an icon theme; the screen is used
// to track the user’s currently configured icon theme,
// which might be different for different screens.
/*

C function : gtk_icon_theme_set_screen
*/
func (recv *IconTheme) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_icon_theme_set_screen((*C.GtkIconTheme)(recv.native), c_screen)

	return
}

// Unsupported : gtk_icon_theme_set_search_path : unsupported parameter path :

// Adds a new #GtkMenuItem to a (table) menu. The number of “cells” that
// an item will occupy is specified by @left_attach, @right_attach,
// @top_attach and @bottom_attach. These each represent the leftmost,
// rightmost, uppermost and lower column and row numbers of the table.
// (Columns and rows are indexed from zero).
//
// Note that this function is not related to gtk_menu_detach().
/*

C function : gtk_menu_attach
*/
func (recv *Menu) Attach(child *Widget, leftAttach uint32, rightAttach uint32, topAttach uint32, bottomAttach uint32) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_left_attach := (C.guint)(leftAttach)

	c_right_attach := (C.guint)(rightAttach)

	c_top_attach := (C.guint)(topAttach)

	c_bottom_attach := (C.guint)(bottomAttach)

	C.gtk_menu_attach((*C.GtkMenu)(recv.native), c_child, c_left_attach, c_right_attach, c_top_attach, c_bottom_attach)

	return
}

// Informs GTK+ on which monitor a menu should be popped up.
// See gdk_monitor_get_geometry().
//
// This function should be called from a #GtkMenuPositionFunc
// if the menu should not appear on the same monitor as the pointer.
// This information can’t be reliably inferred from the coordinates
// returned by a #GtkMenuPositionFunc, since, for very long menus,
// these coordinates may extend beyond the monitor boundaries or even
// the screen boundaries.
/*

C function : gtk_menu_set_monitor
*/
func (recv *Menu) SetMonitor(monitorNum int32) {
	c_monitor_num := (C.gint)(monitorNum)

	C.gtk_menu_set_monitor((*C.GtkMenu)(recv.native), c_monitor_num)

	return
}

// Cancels the selection within the menu shell.
/*

C function : gtk_menu_shell_cancel
*/
func (recv *MenuShell) Cancel() {
	C.gtk_menu_shell_cancel((*C.GtkMenuShell)(recv.native))

	return
}

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Sets the text of the message dialog to be @str, which is marked
// up with the [Pango text markup language][PangoMarkupFormat].
/*

C function : gtk_message_dialog_set_markup
*/
func (recv *MessageDialog) SetMarkup(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_message_dialog_set_markup((*C.GtkMessageDialog)(recv.native), c_str)

	return
}

// Obtains the first child of the paned widget.
/*

C function : gtk_paned_get_child1
*/
func (recv *Paned) GetChild1() *Widget {
	retC := C.gtk_paned_get_child1((*C.GtkPaned)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Obtains the second child of the paned widget.
/*

C function : gtk_paned_get_child2
*/
func (recv *Paned) GetChild2() *Widget {
	retC := C.gtk_paned_get_child2((*C.GtkPaned)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

type signalRadioActionChangedDetail struct {
	callback  RadioActionSignalChangedCallback
	handlerID C.gulong
}

var signalRadioActionChangedId int
var signalRadioActionChangedMap = make(map[int]signalRadioActionChangedDetail)
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
	current := RadioActionNewFromC(unsafe.Pointer(c_current))

	index := int(uintptr(data))
	callback := signalRadioActionChangedMap[index].callback
	callback(current)
}

// Creates a new #GtkRadioAction object. To add the action to
// a #GtkActionGroup and set the accelerator for the action,
// call gtk_action_group_add_action_with_accel().
/*

C function : gtk_radio_action_new
*/
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

// Obtains the value property of the currently active member of
// the group to which @action belongs.
/*

C function : gtk_radio_action_get_current_value
*/
func (recv *RadioAction) GetCurrentValue() int32 {
	retC := C.gtk_radio_action_get_current_value((*C.GtkRadioAction)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the list representing the radio group for this object.
// Note that the returned list is only valid until the next change
// to the group.
//
// A common way to set up a group of radio group is the following:
// |[<!-- language="C" -->
// GSList *group = NULL;
// GtkRadioAction *action;
//
// while ( ...more actions to add... /)
// {
// action = gtk_radio_action_new (...);
//
// gtk_radio_action_set_group (action, group);
// group = gtk_radio_action_get_group (action);
// }
// ]|
/*

C function : gtk_radio_action_get_group
*/
func (recv *RadioAction) GetGroup() *glib.SList {
	retC := C.gtk_radio_action_get_group((*C.GtkRadioAction)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the radio group for the radio action object.
/*

C function : gtk_radio_action_set_group
*/
func (recv *RadioAction) SetGroup(group *glib.SList) {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	C.gtk_radio_action_set_group((*C.GtkRadioAction)(recv.native), c_group)

	return
}

type signalRadioButtonGroupChangedDetail struct {
	callback  RadioButtonSignalGroupChangedCallback
	handlerID C.gulong
}

var signalRadioButtonGroupChangedId int
var signalRadioButtonGroupChangedMap = make(map[int]signalRadioButtonGroupChangedDetail)
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
	index := int(uintptr(data))
	callback := signalRadioButtonGroupChangedMap[index].callback
	callback()
}

// Creates a new #GtkRadioMenuItem adding it to the same group as @group.
/*

C function : gtk_radio_menu_item_new_from_widget
*/
func RadioMenuItemNewFromWidget(group *RadioMenuItem) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioMenuItem)(group.ToC())
	}

	retC := C.gtk_radio_menu_item_new_from_widget(c_group)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new GtkRadioMenuItem whose child is a simple GtkLabel.
// The new #GtkRadioMenuItem is added to the same group as @group.
/*

C function : gtk_radio_menu_item_new_with_label_from_widget
*/
func RadioMenuItemNewWithLabelFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioMenuItem)(group.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_label_from_widget(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new GtkRadioMenuItem containing a label. The label will be
// created using gtk_label_new_with_mnemonic(), so underscores in label
// indicate the mnemonic for the menu item.
//
// The new #GtkRadioMenuItem is added to the same group as @group.
/*

C function : gtk_radio_menu_item_new_with_mnemonic_from_widget
*/
func RadioMenuItemNewWithMnemonicFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioMenuItem)(group.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_mnemonic_from_widget(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkRadioToolButton, adding it to @group.
/*

C function : gtk_radio_tool_button_new
*/
func RadioToolButtonNew(group *glib.SList) *RadioToolButton {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	retC := C.gtk_radio_tool_button_new(c_group)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkRadioToolButton, adding it to @group.
// The new #GtkRadioToolButton will contain an icon and label from the
// stock item indicated by @stock_id.
/*

C function : gtk_radio_tool_button_new_from_stock
*/
func RadioToolButtonNewFromStock(group *glib.SList, stockId string) *RadioToolButton {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_radio_tool_button_new_from_stock(c_group, c_stock_id)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkRadioToolButton adding it to the same group as @gruup
/*

C function : gtk_radio_tool_button_new_from_widget
*/
func RadioToolButtonNewFromWidget(group *RadioToolButton) *RadioToolButton {
	c_group := (*C.GtkRadioToolButton)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioToolButton)(group.ToC())
	}

	retC := C.gtk_radio_tool_button_new_from_widget(c_group)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkRadioToolButton adding it to the same group as @group.
// The new #GtkRadioToolButton will contain an icon and label from the
// stock item indicated by @stock_id.
/*

C function : gtk_radio_tool_button_new_with_stock_from_widget
*/
func RadioToolButtonNewWithStockFromWidget(group *RadioToolButton, stockId string) *RadioToolButton {
	c_group := (*C.GtkRadioToolButton)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioToolButton)(group.ToC())
	}

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_radio_tool_button_new_with_stock_from_widget(c_group, c_stock_id)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the radio button group @button belongs to.
/*

C function : gtk_radio_tool_button_get_group
*/
func (recv *RadioToolButton) GetGroup() *glib.SList {
	retC := C.gtk_radio_tool_button_get_group((*C.GtkRadioToolButton)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds @button to @group, removing it from the group it belonged to before.
/*

C function : gtk_radio_tool_button_set_group
*/
func (recv *RadioToolButton) SetGroup(group *glib.SList) {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	C.gtk_radio_tool_button_set_group((*C.GtkRadioToolButton)(recv.native), c_group)

	return
}

// Gets the #PangoLayout used to display the scale. The returned
// object is owned by the scale so does not need to be freed by
// the caller.
/*

C function : gtk_scale_get_layout
*/
func (recv *Scale) GetLayout() *pango.Layout {
	retC := C.gtk_scale_get_layout((*C.GtkScale)(recv.native))
	var retGo (*pango.Layout)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.LayoutNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Obtains the coordinates where the scale will draw the
// #PangoLayout representing the text in the scale. Remember
// when using the #PangoLayout function you need to convert to
// and from pixels using PANGO_PIXELS() or #PANGO_SCALE.
//
// If the #GtkScale:draw-value property is %FALSE, the return
// values are undefined.
/*

C function : gtk_scale_get_layout_offsets
*/
func (recv *Scale) GetLayoutOffsets() (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	C.gtk_scale_get_layout_offsets((*C.GtkScale)(recv.native), &c_x, &c_y)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// Create a new #GtkSeparatorToolItem
/*

C function : gtk_separator_tool_item_new
*/
func SeparatorToolItemNew() *SeparatorToolItem {
	retC := C.gtk_separator_tool_item_new()
	retGo := SeparatorToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether @item is drawn as a line, or just blank.
// See gtk_separator_tool_item_set_draw().
/*

C function : gtk_separator_tool_item_get_draw
*/
func (recv *SeparatorToolItem) GetDraw() bool {
	retC := C.gtk_separator_tool_item_get_draw((*C.GtkSeparatorToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Whether @item is drawn as a vertical line, or just blank.
// Setting this to %FALSE along with gtk_tool_item_set_expand() is useful
// to create an item that forces following items to the end of the toolbar.
/*

C function : gtk_separator_tool_item_set_draw
*/
func (recv *SeparatorToolItem) SetDraw(draw bool) {
	c_draw :=
		boolToGboolean(draw)

	C.gtk_separator_tool_item_set_draw((*C.GtkSeparatorToolItem)(recv.native), c_draw)

	return
}

type signalStyleRealizeDetail struct {
	callback  StyleSignalRealizeCallback
	handlerID C.gulong
}

var signalStyleRealizeId int
var signalStyleRealizeMap = make(map[int]signalStyleRealizeDetail)
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
	index := int(uintptr(data))
	callback := signalStyleUnrealizeMap[index].callback
	callback()
}

// This function moves the “insert” and “selection_bound” marks
// simultaneously.  If you move them in two steps
// with gtk_text_buffer_move_mark(), you will temporarily select a
// region in between their old and new locations, which can be pretty
// inefficient since the temporarily-selected region will force stuff
// to be recalculated. This function moves them as a unit, which can
// be optimized.
/*

C function : gtk_text_buffer_select_range
*/
func (recv *TextBuffer) SelectRange(ins *TextIter, bound *TextIter) {
	c_ins := (*C.GtkTextIter)(C.NULL)
	if ins != nil {
		c_ins = (*C.GtkTextIter)(ins.ToC())
	}

	c_bound := (*C.GtkTextIter)(C.NULL)
	if bound != nil {
		c_bound = (*C.GtkTextIter)(bound.ToC())
	}

	C.gtk_text_buffer_select_range((*C.GtkTextBuffer)(recv.native), c_ins, c_bound)

	return
}

// Returns whether pressing the Tab key inserts a tab characters.
// gtk_text_view_set_accepts_tab().
/*

C function : gtk_text_view_get_accepts_tab
*/
func (recv *TextView) GetAcceptsTab() bool {
	retC := C.gtk_text_view_get_accepts_tab((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the #GtkTextView is in overwrite mode or not.
/*

C function : gtk_text_view_get_overwrite
*/
func (recv *TextView) GetOverwrite() bool {
	retC := C.gtk_text_view_get_overwrite((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the behavior of the text widget when the Tab key is pressed.
// If @accepts_tab is %TRUE, a tab character is inserted. If @accepts_tab
// is %FALSE the keyboard focus is moved to the next widget in the focus
// chain.
/*

C function : gtk_text_view_set_accepts_tab
*/
func (recv *TextView) SetAcceptsTab(acceptsTab bool) {
	c_accepts_tab :=
		boolToGboolean(acceptsTab)

	C.gtk_text_view_set_accepts_tab((*C.GtkTextView)(recv.native), c_accepts_tab)

	return
}

// Changes the #GtkTextView overwrite mode.
/*

C function : gtk_text_view_set_overwrite
*/
func (recv *TextView) SetOverwrite(overwrite bool) {
	c_overwrite :=
		boolToGboolean(overwrite)

	C.gtk_text_view_set_overwrite((*C.GtkTextView)(recv.native), c_overwrite)

	return
}

// Creates a new #GtkToggleAction object. To add the action to
// a #GtkActionGroup and set the accelerator for the action,
// call gtk_action_group_add_action_with_accel().
/*

C function : gtk_toggle_action_new
*/
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

// Returns the checked state of the toggle action.
/*

C function : gtk_toggle_action_get_active
*/
func (recv *ToggleAction) GetActive() bool {
	retC := C.gtk_toggle_action_get_active((*C.GtkToggleAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the action should have proxies like a radio action.
/*

C function : gtk_toggle_action_get_draw_as_radio
*/
func (recv *ToggleAction) GetDrawAsRadio() bool {
	retC := C.gtk_toggle_action_get_draw_as_radio((*C.GtkToggleAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the checked state on the toggle action.
/*

C function : gtk_toggle_action_set_active
*/
func (recv *ToggleAction) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_toggle_action_set_active((*C.GtkToggleAction)(recv.native), c_is_active)

	return
}

// Sets whether the action should have proxies like a radio action.
/*

C function : gtk_toggle_action_set_draw_as_radio
*/
func (recv *ToggleAction) SetDrawAsRadio(drawAsRadio bool) {
	c_draw_as_radio :=
		boolToGboolean(drawAsRadio)

	C.gtk_toggle_action_set_draw_as_radio((*C.GtkToggleAction)(recv.native), c_draw_as_radio)

	return
}

// Emits the “toggled” signal on the toggle action.
/*

C function : gtk_toggle_action_toggled
*/
func (recv *ToggleAction) Toggled() {
	C.gtk_toggle_action_toggled((*C.GtkToggleAction)(recv.native))

	return
}

// Returns a new #GtkToggleToolButton
/*

C function : gtk_toggle_tool_button_new
*/
func ToggleToolButtonNew() *ToggleToolButton {
	retC := C.gtk_toggle_tool_button_new()
	retGo := ToggleToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkToggleToolButton containing the image and text from a
// stock item. Some stock ids have preprocessor macros like #GTK_STOCK_OK
// and #GTK_STOCK_APPLY.
//
// It is an error if @stock_id is not a name of a stock item.
/*

C function : gtk_toggle_tool_button_new_from_stock
*/
func ToggleToolButtonNewFromStock(stockId string) *ToggleToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_toggle_tool_button_new_from_stock(c_stock_id)
	retGo := ToggleToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Queries a #GtkToggleToolButton and returns its current state.
// Returns %TRUE if the toggle button is pressed in and %FALSE if it is raised.
/*

C function : gtk_toggle_tool_button_get_active
*/
func (recv *ToggleToolButton) GetActive() bool {
	retC := C.gtk_toggle_tool_button_get_active((*C.GtkToggleToolButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the status of the toggle tool button. Set to %TRUE if you
// want the GtkToggleButton to be “pressed in”, and %FALSE to raise it.
// This action causes the toggled signal to be emitted.
/*

C function : gtk_toggle_tool_button_set_active
*/
func (recv *ToggleToolButton) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_toggle_tool_button_set_active((*C.GtkToggleToolButton)(recv.native), c_is_active)

	return
}

// Creates a new #GtkToolButton using @icon_widget as contents and @label as
// label.
/*

C function : gtk_tool_button_new
*/
func ToolButtonNew(iconWidget *Widget, label string) *ToolButton {
	c_icon_widget := (*C.GtkWidget)(C.NULL)
	if iconWidget != nil {
		c_icon_widget = (*C.GtkWidget)(iconWidget.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_tool_button_new(c_icon_widget, c_label)
	retGo := ToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkToolButton containing the image and text from a
// stock item. Some stock ids have preprocessor macros like #GTK_STOCK_OK
// and #GTK_STOCK_APPLY.
//
// It is an error if @stock_id is not a name of a stock item.
/*

C function : gtk_tool_button_new_from_stock
*/
func ToolButtonNewFromStock(stockId string) *ToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_tool_button_new_from_stock(c_stock_id)
	retGo := ToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Return the widget used as icon widget on @button.
// See gtk_tool_button_set_icon_widget().
/*

C function : gtk_tool_button_get_icon_widget
*/
func (recv *ToolButton) GetIconWidget() *Widget {
	retC := C.gtk_tool_button_get_icon_widget((*C.GtkToolButton)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the label used by the tool button, or %NULL if the tool button
// doesn’t have a label. or uses a the label from a stock item. The returned
// string is owned by GTK+, and must not be modified or freed.
/*

C function : gtk_tool_button_get_label
*/
func (recv *ToolButton) GetLabel() string {
	retC := C.gtk_tool_button_get_label((*C.GtkToolButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the widget used as label on @button.
// See gtk_tool_button_set_label_widget().
/*

C function : gtk_tool_button_get_label_widget
*/
func (recv *ToolButton) GetLabelWidget() *Widget {
	retC := C.gtk_tool_button_get_label_widget((*C.GtkToolButton)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the name of the stock item. See gtk_tool_button_set_stock_id().
// The returned string is owned by GTK+ and must not be freed or modifed.
/*

C function : gtk_tool_button_get_stock_id
*/
func (recv *ToolButton) GetStockId() string {
	retC := C.gtk_tool_button_get_stock_id((*C.GtkToolButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns whether underscores in the label property are used as mnemonics
// on menu items on the overflow menu. See gtk_tool_button_set_use_underline().
/*

C function : gtk_tool_button_get_use_underline
*/
func (recv *ToolButton) GetUseUnderline() bool {
	retC := C.gtk_tool_button_get_use_underline((*C.GtkToolButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets @icon as the widget used as icon on @button. If @icon_widget is
// %NULL the icon is determined by the #GtkToolButton:stock-id property. If the
// #GtkToolButton:stock-id property is also %NULL, @button will not have an icon.
/*

C function : gtk_tool_button_set_icon_widget
*/
func (recv *ToolButton) SetIconWidget(iconWidget *Widget) {
	c_icon_widget := (*C.GtkWidget)(C.NULL)
	if iconWidget != nil {
		c_icon_widget = (*C.GtkWidget)(iconWidget.ToC())
	}

	C.gtk_tool_button_set_icon_widget((*C.GtkToolButton)(recv.native), c_icon_widget)

	return
}

// Sets @label as the label used for the tool button. The #GtkToolButton:label
// property only has an effect if not overridden by a non-%NULL
// #GtkToolButton:label-widget property. If both the #GtkToolButton:label-widget
// and #GtkToolButton:label properties are %NULL, the label is determined by the
// #GtkToolButton:stock-id property. If the #GtkToolButton:stock-id property is
// also %NULL, @button will not have a label.
/*

C function : gtk_tool_button_set_label
*/
func (recv *ToolButton) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_tool_button_set_label((*C.GtkToolButton)(recv.native), c_label)

	return
}

// Sets @label_widget as the widget that will be used as the label
// for @button. If @label_widget is %NULL the #GtkToolButton:label property is used
// as label. If #GtkToolButton:label is also %NULL, the label in the stock item
// determined by the #GtkToolButton:stock-id property is used as label. If
// #GtkToolButton:stock-id is also %NULL, @button does not have a label.
/*

C function : gtk_tool_button_set_label_widget
*/
func (recv *ToolButton) SetLabelWidget(labelWidget *Widget) {
	c_label_widget := (*C.GtkWidget)(C.NULL)
	if labelWidget != nil {
		c_label_widget = (*C.GtkWidget)(labelWidget.ToC())
	}

	C.gtk_tool_button_set_label_widget((*C.GtkToolButton)(recv.native), c_label_widget)

	return
}

// Sets the name of the stock item. See gtk_tool_button_new_from_stock().
// The stock_id property only has an effect if not overridden by non-%NULL
// #GtkToolButton:label-widget and #GtkToolButton:icon-widget properties.
/*

C function : gtk_tool_button_set_stock_id
*/
func (recv *ToolButton) SetStockId(stockId string) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_tool_button_set_stock_id((*C.GtkToolButton)(recv.native), c_stock_id)

	return
}

// If set, an underline in the label property indicates that the next character
// should be used for the mnemonic accelerator key in the overflow menu. For
// example, if the label property is “_Open” and @use_underline is %TRUE,
// the label on the tool button will be “Open” and the item on the overflow
// menu will have an underlined “O”.
//
// Labels shown on tool buttons never have mnemonics on them; this property
// only affects the menu item on the overflow menu.
/*

C function : gtk_tool_button_set_use_underline
*/
func (recv *ToolButton) SetUseUnderline(useUnderline bool) {
	c_use_underline :=
		boolToGboolean(useUnderline)

	C.gtk_tool_button_set_use_underline((*C.GtkToolButton)(recv.native), c_use_underline)

	return
}

// Creates a new #GtkToolItem
/*

C function : gtk_tool_item_new
*/
func ToolItemNew() *ToolItem {
	retC := C.gtk_tool_item_new()
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether @tool_item is allocated extra space.
// See gtk_tool_item_set_expand().
/*

C function : gtk_tool_item_get_expand
*/
func (recv *ToolItem) GetExpand() bool {
	retC := C.gtk_tool_item_get_expand((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether @tool_item is the same size as other homogeneous
// items. See gtk_tool_item_set_homogeneous().
/*

C function : gtk_tool_item_get_homogeneous
*/
func (recv *ToolItem) GetHomogeneous() bool {
	retC := C.gtk_tool_item_get_homogeneous((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the icon size used for @tool_item. Custom subclasses of
// #GtkToolItem should call this function to find out what size icons
// they should use.
/*

C function : gtk_tool_item_get_icon_size
*/
func (recv *ToolItem) GetIconSize() IconSize {
	retC := C.gtk_tool_item_get_icon_size((*C.GtkToolItem)(recv.native))
	retGo := (IconSize)(retC)

	return retGo
}

// Returns whether @tool_item is considered important. See
// gtk_tool_item_set_is_important()
/*

C function : gtk_tool_item_get_is_important
*/
func (recv *ToolItem) GetIsImportant() bool {
	retC := C.gtk_tool_item_get_is_important((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the orientation used for @tool_item. Custom subclasses of
// #GtkToolItem should call this function to find out what size icons
// they should use.
/*

C function : gtk_tool_item_get_orientation
*/
func (recv *ToolItem) GetOrientation() Orientation {
	retC := C.gtk_tool_item_get_orientation((*C.GtkToolItem)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// If @menu_item_id matches the string passed to
// gtk_tool_item_set_proxy_menu_item() return the corresponding #GtkMenuItem.
//
// Custom subclasses of #GtkToolItem should use this function to
// update their menu item when the #GtkToolItem changes. That the
// @menu_item_ids must match ensures that a #GtkToolItem
// will not inadvertently change a menu item that they did not create.
/*

C function : gtk_tool_item_get_proxy_menu_item
*/
func (recv *ToolItem) GetProxyMenuItem(menuItemId string) *Widget {
	c_menu_item_id := C.CString(menuItemId)
	defer C.free(unsafe.Pointer(c_menu_item_id))

	retC := C.gtk_tool_item_get_proxy_menu_item((*C.GtkToolItem)(recv.native), c_menu_item_id)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the relief style of @tool_item. See gtk_button_set_relief().
// Custom subclasses of #GtkToolItem should call this function in the handler
// of the #GtkToolItem::toolbar_reconfigured signal to find out the
// relief style of buttons.
/*

C function : gtk_tool_item_get_relief_style
*/
func (recv *ToolItem) GetReliefStyle() ReliefStyle {
	retC := C.gtk_tool_item_get_relief_style((*C.GtkToolItem)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// Returns the toolbar style used for @tool_item. Custom subclasses of
// #GtkToolItem should call this function in the handler of the
// GtkToolItem::toolbar_reconfigured signal to find out in what style
// the toolbar is displayed and change themselves accordingly
//
// Possibilities are:
// - %GTK_TOOLBAR_BOTH, meaning the tool item should show
// both an icon and a label, stacked vertically
// - %GTK_TOOLBAR_ICONS, meaning the toolbar shows only icons
// - %GTK_TOOLBAR_TEXT, meaning the tool item should only show text
// - %GTK_TOOLBAR_BOTH_HORIZ, meaning the tool item should show
// both an icon and a label, arranged horizontally
/*

C function : gtk_tool_item_get_toolbar_style
*/
func (recv *ToolItem) GetToolbarStyle() ToolbarStyle {
	retC := C.gtk_tool_item_get_toolbar_style((*C.GtkToolItem)(recv.native))
	retGo := (ToolbarStyle)(retC)

	return retGo
}

// Returns whether @tool_item has a drag window. See
// gtk_tool_item_set_use_drag_window().
/*

C function : gtk_tool_item_get_use_drag_window
*/
func (recv *ToolItem) GetUseDragWindow() bool {
	retC := C.gtk_tool_item_get_use_drag_window((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the @tool_item is visible on toolbars that are
// docked horizontally.
/*

C function : gtk_tool_item_get_visible_horizontal
*/
func (recv *ToolItem) GetVisibleHorizontal() bool {
	retC := C.gtk_tool_item_get_visible_horizontal((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether @tool_item is visible when the toolbar is docked vertically.
// See gtk_tool_item_set_visible_vertical().
/*

C function : gtk_tool_item_get_visible_vertical
*/
func (recv *ToolItem) GetVisibleVertical() bool {
	retC := C.gtk_tool_item_get_visible_vertical((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the #GtkMenuItem that was last set by
// gtk_tool_item_set_proxy_menu_item(), ie. the #GtkMenuItem
// that is going to appear in the overflow menu.
/*

C function : gtk_tool_item_retrieve_proxy_menu_item
*/
func (recv *ToolItem) RetrieveProxyMenuItem() *Widget {
	retC := C.gtk_tool_item_retrieve_proxy_menu_item((*C.GtkToolItem)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets whether @tool_item is allocated extra space when there
// is more room on the toolbar then needed for the items. The
// effect is that the item gets bigger when the toolbar gets bigger
// and smaller when the toolbar gets smaller.
/*

C function : gtk_tool_item_set_expand
*/
func (recv *ToolItem) SetExpand(expand bool) {
	c_expand :=
		boolToGboolean(expand)

	C.gtk_tool_item_set_expand((*C.GtkToolItem)(recv.native), c_expand)

	return
}

// Sets whether @tool_item is to be allocated the same size as other
// homogeneous items. The effect is that all homogeneous items will have
// the same width as the widest of the items.
/*

C function : gtk_tool_item_set_homogeneous
*/
func (recv *ToolItem) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_tool_item_set_homogeneous((*C.GtkToolItem)(recv.native), c_homogeneous)

	return
}

// Sets whether @tool_item should be considered important. The #GtkToolButton
// class uses this property to determine whether to show or hide its label
// when the toolbar style is %GTK_TOOLBAR_BOTH_HORIZ. The result is that
// only tool buttons with the “is_important” property set have labels, an
// effect known as “priority text”
/*

C function : gtk_tool_item_set_is_important
*/
func (recv *ToolItem) SetIsImportant(isImportant bool) {
	c_is_important :=
		boolToGboolean(isImportant)

	C.gtk_tool_item_set_is_important((*C.GtkToolItem)(recv.native), c_is_important)

	return
}

// Sets the #GtkMenuItem used in the toolbar overflow menu. The
// @menu_item_id is used to identify the caller of this function and
// should also be used with gtk_tool_item_get_proxy_menu_item().
//
// See also #GtkToolItem::create-menu-proxy.
/*

C function : gtk_tool_item_set_proxy_menu_item
*/
func (recv *ToolItem) SetProxyMenuItem(menuItemId string, menuItem *Widget) {
	c_menu_item_id := C.CString(menuItemId)
	defer C.free(unsafe.Pointer(c_menu_item_id))

	c_menu_item := (*C.GtkWidget)(C.NULL)
	if menuItem != nil {
		c_menu_item = (*C.GtkWidget)(menuItem.ToC())
	}

	C.gtk_tool_item_set_proxy_menu_item((*C.GtkToolItem)(recv.native), c_menu_item_id, c_menu_item)

	return
}

// Sets whether @tool_item has a drag window. When %TRUE the
// toolitem can be used as a drag source through gtk_drag_source_set().
// When @tool_item has a drag window it will intercept all events,
// even those that would otherwise be sent to a child of @tool_item.
/*

C function : gtk_tool_item_set_use_drag_window
*/
func (recv *ToolItem) SetUseDragWindow(useDragWindow bool) {
	c_use_drag_window :=
		boolToGboolean(useDragWindow)

	C.gtk_tool_item_set_use_drag_window((*C.GtkToolItem)(recv.native), c_use_drag_window)

	return
}

// Sets whether @tool_item is visible when the toolbar is docked horizontally.
/*

C function : gtk_tool_item_set_visible_horizontal
*/
func (recv *ToolItem) SetVisibleHorizontal(visibleHorizontal bool) {
	c_visible_horizontal :=
		boolToGboolean(visibleHorizontal)

	C.gtk_tool_item_set_visible_horizontal((*C.GtkToolItem)(recv.native), c_visible_horizontal)

	return
}

// Sets whether @tool_item is visible when the toolbar is docked
// vertically. Some tool items, such as text entries, are too wide to be
// useful on a vertically docked toolbar. If @visible_vertical is %FALSE
// @tool_item will not appear on toolbars that are docked vertically.
/*

C function : gtk_tool_item_set_visible_vertical
*/
func (recv *ToolItem) SetVisibleVertical(visibleVertical bool) {
	c_visible_vertical :=
		boolToGboolean(visibleVertical)

	C.gtk_tool_item_set_visible_vertical((*C.GtkToolItem)(recv.native), c_visible_vertical)

	return
}

// Returns the position corresponding to the indicated point on
// @toolbar. This is useful when dragging items to the toolbar:
// this function returns the position a new item should be
// inserted.
//
// @x and @y are in @toolbar coordinates.
/*

C function : gtk_toolbar_get_drop_index
*/
func (recv *Toolbar) GetDropIndex(x int32, y int32) int32 {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_toolbar_get_drop_index((*C.GtkToolbar)(recv.native), c_x, c_y)
	retGo := (int32)(retC)

	return retGo
}

// Returns the position of @item on the toolbar, starting from 0.
// It is an error if @item is not a child of the toolbar.
/*

C function : gtk_toolbar_get_item_index
*/
func (recv *Toolbar) GetItemIndex(item *ToolItem) int32 {
	c_item := (*C.GtkToolItem)(C.NULL)
	if item != nil {
		c_item = (*C.GtkToolItem)(item.ToC())
	}

	retC := C.gtk_toolbar_get_item_index((*C.GtkToolbar)(recv.native), c_item)
	retGo := (int32)(retC)

	return retGo
}

// Returns the number of items on the toolbar.
/*

C function : gtk_toolbar_get_n_items
*/
func (recv *Toolbar) GetNItems() int32 {
	retC := C.gtk_toolbar_get_n_items((*C.GtkToolbar)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the @n'th item on @toolbar, or %NULL if the
// toolbar does not contain an @n'th item.
/*

C function : gtk_toolbar_get_nth_item
*/
func (recv *Toolbar) GetNthItem(n int32) *ToolItem {
	c_n := (C.gint)(n)

	retC := C.gtk_toolbar_get_nth_item((*C.GtkToolbar)(recv.native), c_n)
	var retGo (*ToolItem)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ToolItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the relief style of buttons on @toolbar. See
// gtk_button_set_relief().
/*

C function : gtk_toolbar_get_relief_style
*/
func (recv *Toolbar) GetReliefStyle() ReliefStyle {
	retC := C.gtk_toolbar_get_relief_style((*C.GtkToolbar)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// Returns whether the toolbar has an overflow menu.
// See gtk_toolbar_set_show_arrow().
/*

C function : gtk_toolbar_get_show_arrow
*/
func (recv *Toolbar) GetShowArrow() bool {
	retC := C.gtk_toolbar_get_show_arrow((*C.GtkToolbar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Insert a #GtkToolItem into the toolbar at position @pos. If @pos is
// 0 the item is prepended to the start of the toolbar. If @pos is
// negative, the item is appended to the end of the toolbar.
/*

C function : gtk_toolbar_insert
*/
func (recv *Toolbar) Insert(item *ToolItem, pos int32) {
	c_item := (*C.GtkToolItem)(C.NULL)
	if item != nil {
		c_item = (*C.GtkToolItem)(item.ToC())
	}

	c_pos := (C.gint)(pos)

	C.gtk_toolbar_insert((*C.GtkToolbar)(recv.native), c_item, c_pos)

	return
}

// Highlights @toolbar to give an idea of what it would look like
// if @item was added to @toolbar at the position indicated by @index_.
// If @item is %NULL, highlighting is turned off. In that case @index_
// is ignored.
//
// The @tool_item passed to this function must not be part of any widget
// hierarchy. When an item is set as drop highlight item it can not
// added to any widget hierarchy or used as highlight item for another
// toolbar.
/*

C function : gtk_toolbar_set_drop_highlight_item
*/
func (recv *Toolbar) SetDropHighlightItem(toolItem *ToolItem, index int32) {
	c_tool_item := (*C.GtkToolItem)(C.NULL)
	if toolItem != nil {
		c_tool_item = (*C.GtkToolItem)(toolItem.ToC())
	}

	c_index_ := (C.gint)(index)

	C.gtk_toolbar_set_drop_highlight_item((*C.GtkToolbar)(recv.native), c_tool_item, c_index_)

	return
}

// Sets whether to show an overflow menu when @toolbar isn’t allocated enough
// size to show all of its items. If %TRUE, items which can’t fit in @toolbar,
// and which have a proxy menu item set by gtk_tool_item_set_proxy_menu_item()
// or #GtkToolItem::create-menu-proxy, will be available in an overflow menu,
// which can be opened by an added arrow button. If %FALSE, @toolbar will
// request enough size to fit all of its child items without any overflow.
/*

C function : gtk_toolbar_set_show_arrow
*/
func (recv *Toolbar) SetShowArrow(showArrow bool) {
	c_show_arrow :=
		boolToGboolean(showArrow)

	C.gtk_toolbar_set_show_arrow((*C.GtkToolbar)(recv.native), c_show_arrow)

	return
}

// This function should almost never be called. It clears the @filter
// of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node(). This might be useful if the child model
// being filtered is static (and doesn’t change often) and there has been
// a lot of unreffed access to nodes. As a side effect of this function,
// all unreffed iters will be invalid.
/*

C function : gtk_tree_model_filter_clear_cache
*/
func (recv *TreeModelFilter) ClearCache() {
	C.gtk_tree_model_filter_clear_cache((*C.GtkTreeModelFilter)(recv.native))

	return
}

// Sets @filter_iter to point to the row in @filter that corresponds to the
// row pointed at by @child_iter.  If @filter_iter was not set, %FALSE is
// returned.
/*

C function : gtk_tree_model_filter_convert_child_iter_to_iter
*/
func (recv *TreeModelFilter) ConvertChildIterToIter(childIter *TreeIter) (bool, *TreeIter) {
	var c_filter_iter C.GtkTreeIter

	c_child_iter := (*C.GtkTreeIter)(C.NULL)
	if childIter != nil {
		c_child_iter = (*C.GtkTreeIter)(childIter.ToC())
	}

	retC := C.gtk_tree_model_filter_convert_child_iter_to_iter((*C.GtkTreeModelFilter)(recv.native), &c_filter_iter, c_child_iter)
	retGo := retC == C.TRUE

	filterIter := TreeIterNewFromC(unsafe.Pointer(&c_filter_iter))

	return retGo, filterIter
}

// Converts @child_path to a path relative to @filter. That is, @child_path
// points to a path in the child model. The rerturned path will point to the
// same row in the filtered model. If @child_path isn’t a valid path on the
// child model or points to a row which is not visible in @filter, then %NULL
// is returned.
/*

C function : gtk_tree_model_filter_convert_child_path_to_path
*/
func (recv *TreeModelFilter) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	c_child_path := (*C.GtkTreePath)(C.NULL)
	if childPath != nil {
		c_child_path = (*C.GtkTreePath)(childPath.ToC())
	}

	retC := C.gtk_tree_model_filter_convert_child_path_to_path((*C.GtkTreeModelFilter)(recv.native), c_child_path)
	var retGo (*TreePath)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreePathNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets @child_iter to point to the row pointed to by @filter_iter.
/*

C function : gtk_tree_model_filter_convert_iter_to_child_iter
*/
func (recv *TreeModelFilter) ConvertIterToChildIter(filterIter *TreeIter) *TreeIter {
	var c_child_iter C.GtkTreeIter

	c_filter_iter := (*C.GtkTreeIter)(C.NULL)
	if filterIter != nil {
		c_filter_iter = (*C.GtkTreeIter)(filterIter.ToC())
	}

	C.gtk_tree_model_filter_convert_iter_to_child_iter((*C.GtkTreeModelFilter)(recv.native), &c_child_iter, c_filter_iter)

	childIter := TreeIterNewFromC(unsafe.Pointer(&c_child_iter))

	return childIter
}

// Converts @filter_path to a path on the child model of @filter. That is,
// @filter_path points to a location in @filter. The returned path will
// point to the same location in the model not being filtered. If @filter_path
// does not point to a location in the child model, %NULL is returned.
/*

C function : gtk_tree_model_filter_convert_path_to_child_path
*/
func (recv *TreeModelFilter) ConvertPathToChildPath(filterPath *TreePath) *TreePath {
	c_filter_path := (*C.GtkTreePath)(C.NULL)
	if filterPath != nil {
		c_filter_path = (*C.GtkTreePath)(filterPath.ToC())
	}

	retC := C.gtk_tree_model_filter_convert_path_to_child_path((*C.GtkTreeModelFilter)(recv.native), c_filter_path)
	var retGo (*TreePath)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreePathNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns a pointer to the child model of @filter.
/*

C function : gtk_tree_model_filter_get_model
*/
func (recv *TreeModelFilter) GetModel() *TreeModel {
	retC := C.gtk_tree_model_filter_get_model((*C.GtkTreeModelFilter)(recv.native))
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Emits ::row_changed for each row in the child model, which causes
// the filter to re-evaluate whether a row is visible or not.
/*

C function : gtk_tree_model_filter_refilter
*/
func (recv *TreeModelFilter) Refilter() {
	C.gtk_tree_model_filter_refilter((*C.GtkTreeModelFilter)(recv.native))

	return
}

// Unsupported : gtk_tree_model_filter_set_modify_func : unsupported parameter func : no type generator for TreeModelFilterModifyFunc (GtkTreeModelFilterModifyFunc) for param func

// Sets @column of the child_model to be the column where @filter should
// look for visibility information. @columns should be a column of type
// %G_TYPE_BOOLEAN, where %TRUE means that a row is visible, and %FALSE
// if not.
//
// Note that gtk_tree_model_filter_set_visible_func() or
// gtk_tree_model_filter_set_visible_column() can only be called
// once for a given filter model.
/*

C function : gtk_tree_model_filter_set_visible_column
*/
func (recv *TreeModelFilter) SetVisibleColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_tree_model_filter_set_visible_column((*C.GtkTreeModelFilter)(recv.native), c_column)

	return
}

// Unsupported : gtk_tree_model_filter_set_visible_func : unsupported parameter func : no type generator for TreeModelFilterVisibleFunc (GtkTreeModelFilterVisibleFunc) for param func

// Returns %TRUE if the column expands to fill available space.
/*

C function : gtk_tree_view_column_get_expand
*/
func (recv *TreeViewColumn) GetExpand() bool {
	retC := C.gtk_tree_view_column_get_expand((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the column to take available extra space.  This space is shared equally
// amongst all columns that have the expand set to %TRUE.  If no column has this
// option set, then the last column gets all extra space.  By default, every
// column is created with this %FALSE.
//
// Along with “fixed-width”, the “expand” property changes when the column is
// resized by the user.
/*

C function : gtk_tree_view_column_set_expand
*/
func (recv *TreeViewColumn) SetExpand(expand bool) {
	c_expand :=
		boolToGboolean(expand)

	C.gtk_tree_view_column_set_expand((*C.GtkTreeViewColumn)(recv.native), c_expand)

	return
}

type signalUIManagerActionsChangedDetail struct {
	callback  UIManagerSignalActionsChangedCallback
	handlerID C.gulong
}

var signalUIManagerActionsChangedId int
var signalUIManagerActionsChangedMap = make(map[int]signalUIManagerActionsChangedDetail)
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
	action := ActionNewFromC(unsafe.Pointer(c_action))

	index := int(uintptr(data))
	callback := signalUIManagerPreActivateMap[index].callback
	callback(action)
}

// Creates a new ui manager object.
/*

C function : gtk_ui_manager_new
*/
func UIManagerNew() *UIManager {
	retC := C.gtk_ui_manager_new()
	retGo := UIManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a UI element to the current contents of @manager.
//
// If @type is %GTK_UI_MANAGER_AUTO, GTK+ inserts a menuitem, toolitem or
// separator if such an element can be inserted at the place determined by
// @path. Otherwise @type must indicate an element that can be inserted at
// the place determined by @path.
//
// If @path points to a menuitem or toolitem, the new element will be inserted
// before or after this item, depending on @top.
/*

C function : gtk_ui_manager_add_ui
*/
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

// Parses a file containing a [UI definition][XML-UI] and
// merges it with the current contents of @manager.
/*

C function : gtk_ui_manager_add_ui_from_file
*/
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

// Parses a string containing a [UI definition][XML-UI] and merges it with
// the current contents of @manager. An enclosing <ui> element is added if
// it is missing.
/*

C function : gtk_ui_manager_add_ui_from_string
*/
func (recv *UIManager) AddUiFromString(buffer string) (uint32, error) {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_length := (C.gssize)(len(buffer))

	var cThrowableError *C.GError

	retC := C.gtk_ui_manager_add_ui_from_string((*C.GtkUIManager)(recv.native), c_buffer, c_length, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Makes sure that all pending updates to the UI have been completed.
//
// This may occasionally be necessary, since #GtkUIManager updates the
// UI in an idle function. A typical example where this function is
// useful is to enforce that the menubar and toolbar have been added to
// the main window before showing it:
// |[<!-- language="C" -->
// gtk_container_add (GTK_CONTAINER (window), vbox);
// g_signal_connect (merge, "add-widget",
// G_CALLBACK (add_widget), vbox);
// gtk_ui_manager_add_ui_from_file (merge, "my-menus");
// gtk_ui_manager_add_ui_from_file (merge, "my-toolbars");
// gtk_ui_manager_ensure_update (merge);
// gtk_widget_show (window);
// ]|
/*

C function : gtk_ui_manager_ensure_update
*/
func (recv *UIManager) EnsureUpdate() {
	C.gtk_ui_manager_ensure_update((*C.GtkUIManager)(recv.native))

	return
}

// Returns the #GtkAccelGroup associated with @manager.
/*

C function : gtk_ui_manager_get_accel_group
*/
func (recv *UIManager) GetAccelGroup() *AccelGroup {
	retC := C.gtk_ui_manager_get_accel_group((*C.GtkUIManager)(recv.native))
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Looks up an action by following a path. See gtk_ui_manager_get_widget()
// for more information about paths.
/*

C function : gtk_ui_manager_get_action
*/
func (recv *UIManager) GetAction(path string) *Action {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.gtk_ui_manager_get_action((*C.GtkUIManager)(recv.native), c_path)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the list of action groups associated with @manager.
/*

C function : gtk_ui_manager_get_action_groups
*/
func (recv *UIManager) GetActionGroups() *glib.List {
	retC := C.gtk_ui_manager_get_action_groups((*C.GtkUIManager)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether menus generated by this #GtkUIManager
// will have tearoff menu items.
/*

C function : gtk_ui_manager_get_add_tearoffs
*/
func (recv *UIManager) GetAddTearoffs() bool {
	retC := C.gtk_ui_manager_get_add_tearoffs((*C.GtkUIManager)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Obtains a list of all toplevel widgets of the requested types.
/*

C function : gtk_ui_manager_get_toplevels
*/
func (recv *UIManager) GetToplevels(types UIManagerItemType) *glib.SList {
	c_types := (C.GtkUIManagerItemType)(types)

	retC := C.gtk_ui_manager_get_toplevels((*C.GtkUIManager)(recv.native), c_types)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a [UI definition][XML-UI] of the merged UI.
/*

C function : gtk_ui_manager_get_ui
*/
func (recv *UIManager) GetUi() string {
	retC := C.gtk_ui_manager_get_ui((*C.GtkUIManager)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Looks up a widget by following a path.
// The path consists of the names specified in the XML description of the UI.
// separated by “/”. Elements which don’t have a name or action attribute in
// the XML (e.g. <popup>) can be addressed by their XML element name
// (e.g. "popup"). The root element ("/ui") can be omitted in the path.
//
// Note that the widget found by following a path that ends in a <menu>;
// element is the menuitem to which the menu is attached, not the menu it
// manages.
//
// Also note that the widgets constructed by a ui manager are not tied to
// the lifecycle of the ui manager. If you add the widgets returned by this
// function to some container or explicitly ref them, they will survive the
// destruction of the ui manager.
/*

C function : gtk_ui_manager_get_widget
*/
func (recv *UIManager) GetWidget(path string) *Widget {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.gtk_ui_manager_get_widget((*C.GtkUIManager)(recv.native), c_path)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inserts an action group into the list of action groups associated
// with @manager. Actions in earlier groups hide actions with the same
// name in later groups.
//
// If @pos is larger than the number of action groups in @manager, or
// negative, @action_group will be inserted at the end of the internal
// list.
/*

C function : gtk_ui_manager_insert_action_group
*/
func (recv *UIManager) InsertActionGroup(actionGroup *ActionGroup, pos int32) {
	c_action_group := (*C.GtkActionGroup)(C.NULL)
	if actionGroup != nil {
		c_action_group = (*C.GtkActionGroup)(actionGroup.ToC())
	}

	c_pos := (C.gint)(pos)

	C.gtk_ui_manager_insert_action_group((*C.GtkUIManager)(recv.native), c_action_group, c_pos)

	return
}

// Returns an unused merge id, suitable for use with
// gtk_ui_manager_add_ui().
/*

C function : gtk_ui_manager_new_merge_id
*/
func (recv *UIManager) NewMergeId() uint32 {
	retC := C.gtk_ui_manager_new_merge_id((*C.GtkUIManager)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Removes an action group from the list of action groups associated
// with @manager.
/*

C function : gtk_ui_manager_remove_action_group
*/
func (recv *UIManager) RemoveActionGroup(actionGroup *ActionGroup) {
	c_action_group := (*C.GtkActionGroup)(C.NULL)
	if actionGroup != nil {
		c_action_group = (*C.GtkActionGroup)(actionGroup.ToC())
	}

	C.gtk_ui_manager_remove_action_group((*C.GtkUIManager)(recv.native), c_action_group)

	return
}

// Unmerges the part of @manager's content identified by @merge_id.
/*

C function : gtk_ui_manager_remove_ui
*/
func (recv *UIManager) RemoveUi(mergeId uint32) {
	c_merge_id := (C.guint)(mergeId)

	C.gtk_ui_manager_remove_ui((*C.GtkUIManager)(recv.native), c_merge_id)

	return
}

// Sets the “add_tearoffs” property, which controls whether menus
// generated by this #GtkUIManager will have tearoff menu items.
//
// Note that this only affects regular menus. Generated popup
// menus never have tearoff menu items.
/*

C function : gtk_ui_manager_set_add_tearoffs
*/
func (recv *UIManager) SetAddTearoffs(addTearoffs bool) {
	c_add_tearoffs :=
		boolToGboolean(addTearoffs)

	C.gtk_ui_manager_set_add_tearoffs((*C.GtkUIManager)(recv.native), c_add_tearoffs)

	return
}

// Adds a widget to the list of mnemonic labels for
// this widget. (See gtk_widget_list_mnemonic_labels()). Note the
// list of mnemonic labels for the widget is cleared when the
// widget is destroyed, so the caller must make sure to update
// its internal state at this point as well, by using a connection
// to the #GtkWidget::destroy signal or a weak notifier.
/*

C function : gtk_widget_add_mnemonic_label
*/
func (recv *Widget) AddMnemonicLabel(label *Widget) {
	c_label := (*C.GtkWidget)(C.NULL)
	if label != nil {
		c_label = (*C.GtkWidget)(label.ToC())
	}

	C.gtk_widget_add_mnemonic_label((*C.GtkWidget)(recv.native), c_label)

	return
}

// Determines whether an accelerator that activates the signal
// identified by @signal_id can currently be activated.
// This is done by emitting the #GtkWidget::can-activate-accel
// signal on @widget; if the signal isn’t overridden by a
// handler or in a derived widget, then the default check is
// that the widget must be sensitive, and the widget and all
// its ancestors mapped.
/*

C function : gtk_widget_can_activate_accel
*/
func (recv *Widget) CanActivateAccel(signalId uint32) bool {
	c_signal_id := (C.guint)(signalId)

	retC := C.gtk_widget_can_activate_accel((*C.GtkWidget)(recv.native), c_signal_id)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the list of targets this widget can provide for
// drag-and-drop.
/*

C function : gtk_drag_source_get_target_list
*/
func (recv *Widget) DragSourceGetTargetList() *TargetList {
	retC := C.gtk_drag_source_get_target_list((*C.GtkWidget)(recv.native))
	var retGo (*TargetList)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TargetListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Changes the target types that this widget offers for drag-and-drop.
// The widget must first be made into a drag source with
// gtk_drag_source_set().
/*

C function : gtk_drag_source_set_target_list
*/
func (recv *Widget) DragSourceSetTargetList(targetList *TargetList) {
	c_target_list := (*C.GtkTargetList)(C.NULL)
	if targetList != nil {
		c_target_list = (*C.GtkTargetList)(targetList.ToC())
	}

	C.gtk_drag_source_set_target_list((*C.GtkWidget)(recv.native), c_target_list)

	return
}

// Returns the current value of the #GtkWidget:no-show-all property,
// which determines whether calls to gtk_widget_show_all()
// will affect this widget.
/*

C function : gtk_widget_get_no_show_all
*/
func (recv *Widget) GetNoShowAll() bool {
	retC := C.gtk_widget_get_no_show_all((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns a newly allocated list of the widgets, normally labels, for
// which this widget is the target of a mnemonic (see for example,
// gtk_label_set_mnemonic_widget()).
//
// The widgets in the list are not individually referenced. If you
// want to iterate through the list and perform actions involving
// callbacks that might destroy the widgets, you
// must call `g_list_foreach (result,
// (GFunc)g_object_ref, NULL)` first, and then unref all the
// widgets afterwards.
/*

C function : gtk_widget_list_mnemonic_labels
*/
func (recv *Widget) ListMnemonicLabels() *glib.List {
	retC := C.gtk_widget_list_mnemonic_labels((*C.GtkWidget)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This function works like gtk_widget_queue_resize(),
// except that the widget is not invalidated.
/*

C function : gtk_widget_queue_resize_no_redraw
*/
func (recv *Widget) QueueResizeNoRedraw() {
	C.gtk_widget_queue_resize_no_redraw((*C.GtkWidget)(recv.native))

	return
}

// Removes a widget from the list of mnemonic labels for
// this widget. (See gtk_widget_list_mnemonic_labels()). The widget
// must have previously been added to the list with
// gtk_widget_add_mnemonic_label().
/*

C function : gtk_widget_remove_mnemonic_label
*/
func (recv *Widget) RemoveMnemonicLabel(label *Widget) {
	c_label := (*C.GtkWidget)(C.NULL)
	if label != nil {
		c_label = (*C.GtkWidget)(label.ToC())
	}

	C.gtk_widget_remove_mnemonic_label((*C.GtkWidget)(recv.native), c_label)

	return
}

// Sets the #GtkWidget:no-show-all property, which determines whether
// calls to gtk_widget_show_all() will affect this widget.
//
// This is mostly for use in constructing widget hierarchies with externally
// controlled visibility, see #GtkUIManager.
/*

C function : gtk_widget_set_no_show_all
*/
func (recv *Widget) SetNoShowAll(noShowAll bool) {
	c_no_show_all :=
		boolToGboolean(noShowAll)

	C.gtk_widget_set_no_show_all((*C.GtkWidget)(recv.native), c_no_show_all)

	return
}

// Activates mnemonics and accelerators for this #GtkWindow. This is normally
// called by the default ::key_press_event handler for toplevel windows,
// however in some cases it may be useful to call this directly when
// overriding the standard key handling for a toplevel window.
/*

C function : gtk_window_activate_key
*/
func (recv *Window) ActivateKey(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventKey)(event.ToC())
	}

	retC := C.gtk_window_activate_key((*C.GtkWindow)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value set by gtk_window_set_accept_focus().
/*

C function : gtk_window_get_accept_focus
*/
func (recv *Window) GetAcceptFocus() bool {
	retC := C.gtk_window_get_accept_focus((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the input focus is within this GtkWindow.
// For real toplevel windows, this is identical to gtk_window_is_active(),
// but for embedded windows, like #GtkPlug, the results will differ.
/*

C function : gtk_window_has_toplevel_focus
*/
func (recv *Window) HasToplevelFocus() bool {
	retC := C.gtk_window_has_toplevel_focus((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the window is part of the current active toplevel.
// (That is, the toplevel window receiving keystrokes.)
// The return value is %TRUE if the window is active toplevel
// itself, but also if it is, say, a #GtkPlug embedded in the active toplevel.
// You might use this function if you wanted to draw a widget
// differently in an active window from a widget in an inactive window.
// See gtk_window_has_toplevel_focus()
/*

C function : gtk_window_is_active
*/
func (recv *Window) IsActive() bool {
	retC := C.gtk_window_is_active((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Propagate a key press or release event to the focus widget and
// up the focus container chain until a widget handles @event.
// This is normally called by the default ::key_press_event and
// ::key_release_event handlers for toplevel windows,
// however in some cases it may be useful to call this directly when
// overriding the standard key handling for a toplevel window.
/*

C function : gtk_window_propagate_key_event
*/
func (recv *Window) PropagateKeyEvent(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventKey)(event.ToC())
	}

	retC := C.gtk_window_propagate_key_event((*C.GtkWindow)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// Windows may set a hint asking the desktop environment not to receive
// the input focus. This function sets this hint.
/*

C function : gtk_window_set_accept_focus
*/
func (recv *Window) SetAcceptFocus(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_accept_focus((*C.GtkWindow)(recv.native), c_setting)

	return
}

// Asks to keep @window above, so that it stays on top. Note that
// you shouldn’t assume the window is definitely above afterward,
// because other entities (e.g. the user or
// [window manager][gtk-X11-arch]) could not keep it above,
// and not all window managers support keeping windows above. But
// normally the window will end kept above. Just don’t write code
// that crashes if not.
//
// It’s permitted to call this function before showing a window,
// in which case the window will be kept above when it appears onscreen
// initially.
//
// You can track the above state via the “window-state-event” signal
// on #GtkWidget.
//
// Note that, according to the
// [Extended Window Manager Hints Specification](http://www.freedesktop.org/Standards/wm-spec),
// the above state is mainly meant for user preferences and should not
// be used by applications e.g. for drawing attention to their
// dialogs.
/*

C function : gtk_window_set_keep_above
*/
func (recv *Window) SetKeepAbove(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_keep_above((*C.GtkWindow)(recv.native), c_setting)

	return
}

// Asks to keep @window below, so that it stays in bottom. Note that
// you shouldn’t assume the window is definitely below afterward,
// because other entities (e.g. the user or
// [window manager][gtk-X11-arch]) could not keep it below,
// and not all window managers support putting windows below. But
// normally the window will be kept below. Just don’t write code
// that crashes if not.
//
// It’s permitted to call this function before showing a window,
// in which case the window will be kept below when it appears onscreen
// initially.
//
// You can track the below state via the “window-state-event” signal
// on #GtkWidget.
//
// Note that, according to the
// [Extended Window Manager Hints Specification](http://www.freedesktop.org/Standards/wm-spec),
// the above state is mainly meant for user preferences and should not
// be used by applications e.g. for drawing attention to their
// dialogs.
/*

C function : gtk_window_set_keep_below
*/
func (recv *Window) SetKeepBelow(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_keep_below((*C.GtkWindow)(recv.native), c_setting)

	return
}
