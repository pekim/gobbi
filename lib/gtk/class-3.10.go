// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gio "github.com/pekim/gobbi/lib/gio"
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

	void listbox_rowActivatedHandler(GObject *, GtkListBoxRow *, gpointer);

	static gulong ListBox_signal_connect_row_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-activated", G_CALLBACK(listbox_rowActivatedHandler), data);
	}

*/
/*

	void listbox_rowSelectedHandler(GObject *, GtkListBoxRow *, gpointer);

	static gulong ListBox_signal_connect_row_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-selected", G_CALLBACK(listbox_rowSelectedHandler), data);
	}

*/
/*

	void listboxrow_activateHandler(GObject *, gpointer);

	static gulong ListBoxRow_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(listboxrow_activateHandler), data);
	}

*/
/*

	gint placessidebar_dragActionAskHandler(GObject *, gint, gpointer);

	static gulong PlacesSidebar_signal_connect_drag_action_ask(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-action-ask", G_CALLBACK(placessidebar_dragActionAskHandler), data);
	}

*/
/*

	void placessidebar_openLocationHandler(GObject *, GFile *, GtkPlacesOpenFlags, gpointer);

	static gulong PlacesSidebar_signal_connect_open_location(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "open-location", G_CALLBACK(placessidebar_openLocationHandler), data);
	}

*/
/*

	void placessidebar_populatePopupHandler(GObject *, GtkWidget *, GFile *, GVolume *, gpointer);

	static gulong PlacesSidebar_signal_connect_populate_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "populate-popup", G_CALLBACK(placessidebar_populatePopupHandler), data);
	}

*/
/*

	void placessidebar_showErrorMessageHandler(GObject *, gchar*, gchar*, gpointer);

	static gulong PlacesSidebar_signal_connect_show_error_message(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-error-message", G_CALLBACK(placessidebar_showErrorMessageHandler), data);
	}

*/
/*

	void searchentry_searchChangedHandler(GObject *, gpointer);

	static gulong SearchEntry_signal_connect_search_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "search-changed", G_CALLBACK(searchentry_searchChangedHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_box_get_baseline_position

// Blacklisted : gtk_box_set_baseline_position

// Blacklisted : gtk_builder_new_from_file

// Blacklisted : gtk_builder_new_from_resource

// Blacklisted : gtk_builder_new_from_string

// Unsupported : gtk_builder_add_callback_symbol : unsupported parameter callback_symbol : no type generator for GObject.Callback (GCallback) for param callback_symbol

// Unsupported : gtk_builder_add_callback_symbols : unsupported parameter first_callback_symbol : no type generator for GObject.Callback (GCallback) for param first_callback_symbol

// Blacklisted : gtk_builder_get_application

// Unsupported : gtk_builder_lookup_callback_symbol : no return generator

// Blacklisted : gtk_builder_set_application

// Blacklisted : gtk_button_new_from_icon_name

// Blacklisted : gtk_entry_get_tabs

// Blacklisted : gtk_entry_set_tabs

// Blacklisted : gtk_grid_get_baseline_row

// Blacklisted : gtk_grid_get_row_baseline_position

// Blacklisted : gtk_grid_remove_column

// Blacklisted : gtk_grid_remove_row

// Blacklisted : gtk_grid_set_baseline_row

// Blacklisted : gtk_grid_set_row_baseline_position

// Blacklisted : gtk_header_bar_new

// Blacklisted : gtk_header_bar_get_custom_title

// Blacklisted : gtk_header_bar_get_show_close_button

// Blacklisted : gtk_header_bar_get_subtitle

// GetTitle is a wrapper around the C function gtk_header_bar_get_title.
func (recv *HeaderBar) GetTitle() string {
	retC := C.gtk_header_bar_get_title((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : gtk_header_bar_pack_end

// Blacklisted : gtk_header_bar_pack_start

// Blacklisted : gtk_header_bar_set_custom_title

// Blacklisted : gtk_header_bar_set_show_close_button

// Blacklisted : gtk_header_bar_set_subtitle

// SetTitle is a wrapper around the C function gtk_header_bar_set_title.
func (recv *HeaderBar) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_header_bar_set_title((*C.GtkHeaderBar)(recv.native), c_title)

	return
}

// Blacklisted : gtk_icon_info_get_base_scale

// Blacklisted : gtk_icon_info_load_surface

// Blacklisted : gtk_icon_theme_choose_icon_for_scale

// Blacklisted : gtk_icon_theme_load_icon_for_scale

// Blacklisted : gtk_icon_theme_load_surface

// Blacklisted : gtk_icon_theme_lookup_by_gicon_for_scale

// Blacklisted : gtk_icon_theme_lookup_icon_for_scale

// Blacklisted : gtk_image_new_from_surface

// Blacklisted : gtk_image_set_from_surface

// Blacklisted : gtk_info_bar_get_show_close_button

// Blacklisted : gtk_info_bar_set_show_close_button

// Blacklisted : gtk_label_get_lines

// Blacklisted : gtk_label_set_lines

type signalListBoxRowActivatedDetail struct {
	callback  ListBoxSignalRowActivatedCallback
	handlerID C.gulong
}

var signalListBoxRowActivatedId int
var signalListBoxRowActivatedMap = make(map[int]signalListBoxRowActivatedDetail)
var signalListBoxRowActivatedLock sync.RWMutex

// ListBoxSignalRowActivatedCallback is a callback function for a 'row-activated' signal emitted from a ListBox.
type ListBoxSignalRowActivatedCallback func(row *ListBoxRow)

/*
ConnectRowActivated connects the callback to the 'row-activated' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectRowActivated to remove it.
*/
func (recv *ListBox) ConnectRowActivated(callback ListBoxSignalRowActivatedCallback) int {
	signalListBoxRowActivatedLock.Lock()
	defer signalListBoxRowActivatedLock.Unlock()

	signalListBoxRowActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_row_activated(instance, C.gpointer(uintptr(signalListBoxRowActivatedId)))

	detail := signalListBoxRowActivatedDetail{callback, handlerID}
	signalListBoxRowActivatedMap[signalListBoxRowActivatedId] = detail

	return signalListBoxRowActivatedId
}

/*
DisconnectRowActivated disconnects a callback from the 'row-activated' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectRowActivated.
*/
func (recv *ListBox) DisconnectRowActivated(connectionID int) {
	signalListBoxRowActivatedLock.Lock()
	defer signalListBoxRowActivatedLock.Unlock()

	detail, exists := signalListBoxRowActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowActivatedMap, connectionID)
}

//export listbox_rowActivatedHandler
func listbox_rowActivatedHandler(_ *C.GObject, c_row *C.GtkListBoxRow, data C.gpointer) {
	signalListBoxRowActivatedLock.RLock()
	defer signalListBoxRowActivatedLock.RUnlock()

	row := ListBoxRowNewFromC(unsafe.Pointer(c_row))

	index := int(uintptr(data))
	callback := signalListBoxRowActivatedMap[index].callback
	callback(row)
}

type signalListBoxRowSelectedDetail struct {
	callback  ListBoxSignalRowSelectedCallback
	handlerID C.gulong
}

var signalListBoxRowSelectedId int
var signalListBoxRowSelectedMap = make(map[int]signalListBoxRowSelectedDetail)
var signalListBoxRowSelectedLock sync.RWMutex

// ListBoxSignalRowSelectedCallback is a callback function for a 'row-selected' signal emitted from a ListBox.
type ListBoxSignalRowSelectedCallback func(row *ListBoxRow)

/*
ConnectRowSelected connects the callback to the 'row-selected' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectRowSelected to remove it.
*/
func (recv *ListBox) ConnectRowSelected(callback ListBoxSignalRowSelectedCallback) int {
	signalListBoxRowSelectedLock.Lock()
	defer signalListBoxRowSelectedLock.Unlock()

	signalListBoxRowSelectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_row_selected(instance, C.gpointer(uintptr(signalListBoxRowSelectedId)))

	detail := signalListBoxRowSelectedDetail{callback, handlerID}
	signalListBoxRowSelectedMap[signalListBoxRowSelectedId] = detail

	return signalListBoxRowSelectedId
}

/*
DisconnectRowSelected disconnects a callback from the 'row-selected' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectRowSelected.
*/
func (recv *ListBox) DisconnectRowSelected(connectionID int) {
	signalListBoxRowSelectedLock.Lock()
	defer signalListBoxRowSelectedLock.Unlock()

	detail, exists := signalListBoxRowSelectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowSelectedMap, connectionID)
}

//export listbox_rowSelectedHandler
func listbox_rowSelectedHandler(_ *C.GObject, c_row *C.GtkListBoxRow, data C.gpointer) {
	signalListBoxRowSelectedLock.RLock()
	defer signalListBoxRowSelectedLock.RUnlock()

	row := ListBoxRowNewFromC(unsafe.Pointer(c_row))

	index := int(uintptr(data))
	callback := signalListBoxRowSelectedMap[index].callback
	callback(row)
}

// Blacklisted : gtk_list_box_new

// Blacklisted : gtk_list_box_drag_highlight_row

// Blacklisted : gtk_list_box_drag_unhighlight_row

// Blacklisted : gtk_list_box_get_activate_on_single_click

// Blacklisted : gtk_list_box_get_adjustment

// Blacklisted : gtk_list_box_get_row_at_index

// Blacklisted : gtk_list_box_get_row_at_y

// Blacklisted : gtk_list_box_get_selected_row

// Blacklisted : gtk_list_box_get_selection_mode

// Blacklisted : gtk_list_box_insert

// Blacklisted : gtk_list_box_invalidate_filter

// Blacklisted : gtk_list_box_invalidate_headers

// Blacklisted : gtk_list_box_invalidate_sort

// Blacklisted : gtk_list_box_prepend

// Blacklisted : gtk_list_box_select_row

// Blacklisted : gtk_list_box_set_activate_on_single_click

// Blacklisted : gtk_list_box_set_adjustment

// Unsupported : gtk_list_box_set_filter_func : unsupported parameter filter_func : no type generator for ListBoxFilterFunc (GtkListBoxFilterFunc) for param filter_func

// Unsupported : gtk_list_box_set_header_func : unsupported parameter update_header : no type generator for ListBoxUpdateHeaderFunc (GtkListBoxUpdateHeaderFunc) for param update_header

// Blacklisted : gtk_list_box_set_placeholder

// Blacklisted : gtk_list_box_set_selection_mode

// Unsupported : gtk_list_box_set_sort_func : unsupported parameter sort_func : no type generator for ListBoxSortFunc (GtkListBoxSortFunc) for param sort_func

type signalListBoxRowActivateDetail struct {
	callback  ListBoxRowSignalActivateCallback
	handlerID C.gulong
}

var signalListBoxRowActivateId int
var signalListBoxRowActivateMap = make(map[int]signalListBoxRowActivateDetail)
var signalListBoxRowActivateLock sync.RWMutex

// ListBoxRowSignalActivateCallback is a callback function for a 'activate' signal emitted from a ListBoxRow.
type ListBoxRowSignalActivateCallback func()

/*
ConnectActivate connects the callback to the 'activate' signal for the ListBoxRow.

The returned value represents the connection, and may be passed to DisconnectActivate to remove it.
*/
func (recv *ListBoxRow) ConnectActivate(callback ListBoxRowSignalActivateCallback) int {
	signalListBoxRowActivateLock.Lock()
	defer signalListBoxRowActivateLock.Unlock()

	signalListBoxRowActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBoxRow_signal_connect_activate(instance, C.gpointer(uintptr(signalListBoxRowActivateId)))

	detail := signalListBoxRowActivateDetail{callback, handlerID}
	signalListBoxRowActivateMap[signalListBoxRowActivateId] = detail

	return signalListBoxRowActivateId
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the ListBoxRow.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *ListBoxRow) DisconnectActivate(connectionID int) {
	signalListBoxRowActivateLock.Lock()
	defer signalListBoxRowActivateLock.Unlock()

	detail, exists := signalListBoxRowActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowActivateMap, connectionID)
}

//export listboxrow_activateHandler
func listboxrow_activateHandler(_ *C.GObject, data C.gpointer) {
	signalListBoxRowActivateLock.RLock()
	defer signalListBoxRowActivateLock.RUnlock()

	index := int(uintptr(data))
	callback := signalListBoxRowActivateMap[index].callback
	callback()
}

// Blacklisted : gtk_list_box_row_new

// Blacklisted : gtk_list_box_row_changed

// Blacklisted : gtk_list_box_row_get_header

// Blacklisted : gtk_list_box_row_get_index

// Blacklisted : gtk_list_box_row_set_header

type signalPlacesSidebarDragActionAskDetail struct {
	callback  PlacesSidebarSignalDragActionAskCallback
	handlerID C.gulong
}

var signalPlacesSidebarDragActionAskId int
var signalPlacesSidebarDragActionAskMap = make(map[int]signalPlacesSidebarDragActionAskDetail)
var signalPlacesSidebarDragActionAskLock sync.RWMutex

// PlacesSidebarSignalDragActionAskCallback is a callback function for a 'drag-action-ask' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalDragActionAskCallback func(actions int32) int32

/*
ConnectDragActionAsk connects the callback to the 'drag-action-ask' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectDragActionAsk to remove it.
*/
func (recv *PlacesSidebar) ConnectDragActionAsk(callback PlacesSidebarSignalDragActionAskCallback) int {
	signalPlacesSidebarDragActionAskLock.Lock()
	defer signalPlacesSidebarDragActionAskLock.Unlock()

	signalPlacesSidebarDragActionAskId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_drag_action_ask(instance, C.gpointer(uintptr(signalPlacesSidebarDragActionAskId)))

	detail := signalPlacesSidebarDragActionAskDetail{callback, handlerID}
	signalPlacesSidebarDragActionAskMap[signalPlacesSidebarDragActionAskId] = detail

	return signalPlacesSidebarDragActionAskId
}

/*
DisconnectDragActionAsk disconnects a callback from the 'drag-action-ask' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectDragActionAsk.
*/
func (recv *PlacesSidebar) DisconnectDragActionAsk(connectionID int) {
	signalPlacesSidebarDragActionAskLock.Lock()
	defer signalPlacesSidebarDragActionAskLock.Unlock()

	detail, exists := signalPlacesSidebarDragActionAskMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarDragActionAskMap, connectionID)
}

//export placessidebar_dragActionAskHandler
func placessidebar_dragActionAskHandler(_ *C.GObject, c_actions C.gint, data C.gpointer) C.gint {
	signalPlacesSidebarDragActionAskLock.RLock()
	defer signalPlacesSidebarDragActionAskLock.RUnlock()

	actions := int32(c_actions)

	index := int(uintptr(data))
	callback := signalPlacesSidebarDragActionAskMap[index].callback
	retGo := callback(actions)
	retC :=
		(C.gint)(retGo)
	return retC
}

// Unsupported signal 'drag-action-requested' for PlacesSidebar : param source_file_list : gpointer

// Unsupported signal 'drag-perform-drop' for PlacesSidebar : param source_file_list : gpointer

type signalPlacesSidebarOpenLocationDetail struct {
	callback  PlacesSidebarSignalOpenLocationCallback
	handlerID C.gulong
}

var signalPlacesSidebarOpenLocationId int
var signalPlacesSidebarOpenLocationMap = make(map[int]signalPlacesSidebarOpenLocationDetail)
var signalPlacesSidebarOpenLocationLock sync.RWMutex

// PlacesSidebarSignalOpenLocationCallback is a callback function for a 'open-location' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalOpenLocationCallback func(location *gio.File, openFlags PlacesOpenFlags)

/*
ConnectOpenLocation connects the callback to the 'open-location' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectOpenLocation to remove it.
*/
func (recv *PlacesSidebar) ConnectOpenLocation(callback PlacesSidebarSignalOpenLocationCallback) int {
	signalPlacesSidebarOpenLocationLock.Lock()
	defer signalPlacesSidebarOpenLocationLock.Unlock()

	signalPlacesSidebarOpenLocationId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_open_location(instance, C.gpointer(uintptr(signalPlacesSidebarOpenLocationId)))

	detail := signalPlacesSidebarOpenLocationDetail{callback, handlerID}
	signalPlacesSidebarOpenLocationMap[signalPlacesSidebarOpenLocationId] = detail

	return signalPlacesSidebarOpenLocationId
}

/*
DisconnectOpenLocation disconnects a callback from the 'open-location' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectOpenLocation.
*/
func (recv *PlacesSidebar) DisconnectOpenLocation(connectionID int) {
	signalPlacesSidebarOpenLocationLock.Lock()
	defer signalPlacesSidebarOpenLocationLock.Unlock()

	detail, exists := signalPlacesSidebarOpenLocationMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarOpenLocationMap, connectionID)
}

//export placessidebar_openLocationHandler
func placessidebar_openLocationHandler(_ *C.GObject, c_location *C.GFile, c_open_flags C.GtkPlacesOpenFlags, data C.gpointer) {
	signalPlacesSidebarOpenLocationLock.RLock()
	defer signalPlacesSidebarOpenLocationLock.RUnlock()

	location := gio.FileNewFromC(unsafe.Pointer(c_location))

	openFlags := PlacesOpenFlags(c_open_flags)

	index := int(uintptr(data))
	callback := signalPlacesSidebarOpenLocationMap[index].callback
	callback(location, openFlags)
}

type signalPlacesSidebarPopulatePopupDetail struct {
	callback  PlacesSidebarSignalPopulatePopupCallback
	handlerID C.gulong
}

var signalPlacesSidebarPopulatePopupId int
var signalPlacesSidebarPopulatePopupMap = make(map[int]signalPlacesSidebarPopulatePopupDetail)
var signalPlacesSidebarPopulatePopupLock sync.RWMutex

// PlacesSidebarSignalPopulatePopupCallback is a callback function for a 'populate-popup' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalPopulatePopupCallback func(container *Widget, selectedItem *gio.File, selectedVolume *gio.Volume)

/*
ConnectPopulatePopup connects the callback to the 'populate-popup' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectPopulatePopup to remove it.
*/
func (recv *PlacesSidebar) ConnectPopulatePopup(callback PlacesSidebarSignalPopulatePopupCallback) int {
	signalPlacesSidebarPopulatePopupLock.Lock()
	defer signalPlacesSidebarPopulatePopupLock.Unlock()

	signalPlacesSidebarPopulatePopupId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_populate_popup(instance, C.gpointer(uintptr(signalPlacesSidebarPopulatePopupId)))

	detail := signalPlacesSidebarPopulatePopupDetail{callback, handlerID}
	signalPlacesSidebarPopulatePopupMap[signalPlacesSidebarPopulatePopupId] = detail

	return signalPlacesSidebarPopulatePopupId
}

/*
DisconnectPopulatePopup disconnects a callback from the 'populate-popup' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectPopulatePopup.
*/
func (recv *PlacesSidebar) DisconnectPopulatePopup(connectionID int) {
	signalPlacesSidebarPopulatePopupLock.Lock()
	defer signalPlacesSidebarPopulatePopupLock.Unlock()

	detail, exists := signalPlacesSidebarPopulatePopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarPopulatePopupMap, connectionID)
}

//export placessidebar_populatePopupHandler
func placessidebar_populatePopupHandler(_ *C.GObject, c_container *C.GtkWidget, c_selected_item *C.GFile, c_selected_volume *C.GVolume, data C.gpointer) {
	signalPlacesSidebarPopulatePopupLock.RLock()
	defer signalPlacesSidebarPopulatePopupLock.RUnlock()

	container := WidgetNewFromC(unsafe.Pointer(c_container))

	selectedItem := gio.FileNewFromC(unsafe.Pointer(c_selected_item))

	selectedVolume := gio.VolumeNewFromC(unsafe.Pointer(c_selected_volume))

	index := int(uintptr(data))
	callback := signalPlacesSidebarPopulatePopupMap[index].callback
	callback(container, selectedItem, selectedVolume)
}

type signalPlacesSidebarShowErrorMessageDetail struct {
	callback  PlacesSidebarSignalShowErrorMessageCallback
	handlerID C.gulong
}

var signalPlacesSidebarShowErrorMessageId int
var signalPlacesSidebarShowErrorMessageMap = make(map[int]signalPlacesSidebarShowErrorMessageDetail)
var signalPlacesSidebarShowErrorMessageLock sync.RWMutex

// PlacesSidebarSignalShowErrorMessageCallback is a callback function for a 'show-error-message' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowErrorMessageCallback func(primary string, secondary string)

/*
ConnectShowErrorMessage connects the callback to the 'show-error-message' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectShowErrorMessage to remove it.
*/
func (recv *PlacesSidebar) ConnectShowErrorMessage(callback PlacesSidebarSignalShowErrorMessageCallback) int {
	signalPlacesSidebarShowErrorMessageLock.Lock()
	defer signalPlacesSidebarShowErrorMessageLock.Unlock()

	signalPlacesSidebarShowErrorMessageId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_show_error_message(instance, C.gpointer(uintptr(signalPlacesSidebarShowErrorMessageId)))

	detail := signalPlacesSidebarShowErrorMessageDetail{callback, handlerID}
	signalPlacesSidebarShowErrorMessageMap[signalPlacesSidebarShowErrorMessageId] = detail

	return signalPlacesSidebarShowErrorMessageId
}

/*
DisconnectShowErrorMessage disconnects a callback from the 'show-error-message' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectShowErrorMessage.
*/
func (recv *PlacesSidebar) DisconnectShowErrorMessage(connectionID int) {
	signalPlacesSidebarShowErrorMessageLock.Lock()
	defer signalPlacesSidebarShowErrorMessageLock.Unlock()

	detail, exists := signalPlacesSidebarShowErrorMessageMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarShowErrorMessageMap, connectionID)
}

//export placessidebar_showErrorMessageHandler
func placessidebar_showErrorMessageHandler(_ *C.GObject, c_primary *C.gchar, c_secondary *C.gchar, data C.gpointer) {
	signalPlacesSidebarShowErrorMessageLock.RLock()
	defer signalPlacesSidebarShowErrorMessageLock.RUnlock()

	primary := C.GoString(c_primary)

	secondary := C.GoString(c_secondary)

	index := int(uintptr(data))
	callback := signalPlacesSidebarShowErrorMessageMap[index].callback
	callback(primary, secondary)
}

// Blacklisted : gtk_places_sidebar_new

// Blacklisted : gtk_places_sidebar_add_shortcut

// Blacklisted : gtk_places_sidebar_get_location

// Blacklisted : gtk_places_sidebar_get_nth_bookmark

// Blacklisted : gtk_places_sidebar_get_open_flags

// Blacklisted : gtk_places_sidebar_get_show_desktop

// Blacklisted : gtk_places_sidebar_list_shortcuts

// Blacklisted : gtk_places_sidebar_remove_shortcut

// Blacklisted : gtk_places_sidebar_set_location

// Blacklisted : gtk_places_sidebar_set_open_flags

// Blacklisted : gtk_places_sidebar_set_show_connect_to_server

// Blacklisted : gtk_places_sidebar_set_show_desktop

// Blacklisted : gtk_revealer_new

// Blacklisted : gtk_revealer_get_child_revealed

// Blacklisted : gtk_revealer_get_reveal_child

// Blacklisted : gtk_revealer_get_transition_duration

// Blacklisted : gtk_revealer_get_transition_type

// Blacklisted : gtk_revealer_set_reveal_child

// Blacklisted : gtk_revealer_set_transition_duration

// Blacklisted : gtk_revealer_set_transition_type

// Blacklisted : gtk_search_bar_new

// Blacklisted : gtk_search_bar_connect_entry

// Blacklisted : gtk_search_bar_get_search_mode

// Blacklisted : gtk_search_bar_get_show_close_button

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_search_bar_set_search_mode

// Blacklisted : gtk_search_bar_set_show_close_button

type signalSearchEntrySearchChangedDetail struct {
	callback  SearchEntrySignalSearchChangedCallback
	handlerID C.gulong
}

var signalSearchEntrySearchChangedId int
var signalSearchEntrySearchChangedMap = make(map[int]signalSearchEntrySearchChangedDetail)
var signalSearchEntrySearchChangedLock sync.RWMutex

// SearchEntrySignalSearchChangedCallback is a callback function for a 'search-changed' signal emitted from a SearchEntry.
type SearchEntrySignalSearchChangedCallback func()

/*
ConnectSearchChanged connects the callback to the 'search-changed' signal for the SearchEntry.

The returned value represents the connection, and may be passed to DisconnectSearchChanged to remove it.
*/
func (recv *SearchEntry) ConnectSearchChanged(callback SearchEntrySignalSearchChangedCallback) int {
	signalSearchEntrySearchChangedLock.Lock()
	defer signalSearchEntrySearchChangedLock.Unlock()

	signalSearchEntrySearchChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.SearchEntry_signal_connect_search_changed(instance, C.gpointer(uintptr(signalSearchEntrySearchChangedId)))

	detail := signalSearchEntrySearchChangedDetail{callback, handlerID}
	signalSearchEntrySearchChangedMap[signalSearchEntrySearchChangedId] = detail

	return signalSearchEntrySearchChangedId
}

/*
DisconnectSearchChanged disconnects a callback from the 'search-changed' signal for the SearchEntry.

The connectionID should be a value returned from a call to ConnectSearchChanged.
*/
func (recv *SearchEntry) DisconnectSearchChanged(connectionID int) {
	signalSearchEntrySearchChangedLock.Lock()
	defer signalSearchEntrySearchChangedLock.Unlock()

	detail, exists := signalSearchEntrySearchChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSearchEntrySearchChangedMap, connectionID)
}

//export searchentry_searchChangedHandler
func searchentry_searchChangedHandler(_ *C.GObject, data C.gpointer) {
	signalSearchEntrySearchChangedLock.RLock()
	defer signalSearchEntrySearchChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalSearchEntrySearchChangedMap[index].callback
	callback()
}

// Blacklisted : gtk_stack_new

// Blacklisted : gtk_stack_add_named

// Blacklisted : gtk_stack_add_titled

// Blacklisted : gtk_stack_get_homogeneous

// Blacklisted : gtk_stack_get_transition_duration

// Blacklisted : gtk_stack_get_transition_type

// Blacklisted : gtk_stack_get_visible_child

// Blacklisted : gtk_stack_get_visible_child_name

// Blacklisted : gtk_stack_set_homogeneous

// Blacklisted : gtk_stack_set_transition_duration

// Blacklisted : gtk_stack_set_transition_type

// Blacklisted : gtk_stack_set_visible_child

// Blacklisted : gtk_stack_set_visible_child_full

// Blacklisted : gtk_stack_set_visible_child_name

// Blacklisted : gtk_stack_switcher_new

// Blacklisted : gtk_stack_switcher_get_stack

// Blacklisted : gtk_stack_switcher_set_stack

// Blacklisted : gtk_style_context_get_scale

// Blacklisted : gtk_style_context_set_scale

// Unsupported : gtk_drag_begin_with_coordinates : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_widget_get_allocated_baseline

// Blacklisted : gtk_widget_get_preferred_height_and_baseline_for_width

// Blacklisted : gtk_widget_get_scale_factor

// Blacklisted : gtk_widget_get_valign_with_baseline

// Blacklisted : gtk_widget_init_template

// Blacklisted : gtk_widget_size_allocate_with_baseline

// Blacklisted : gtk_window_close

// Blacklisted : gtk_window_set_titlebar
