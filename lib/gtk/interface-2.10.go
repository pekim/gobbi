// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void recentchooser_itemActivatedHandler(GObject *, gpointer);

	static gulong RecentChooser_signal_connect_item_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "item-activated", G_CALLBACK(recentchooser_itemActivatedHandler), data);
	}

*/
/*

	void recentchooser_selectionChangedHandler(GObject *, gpointer);

	static gulong RecentChooser_signal_connect_selection_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-changed", G_CALLBACK(recentchooser_selectionChangedHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_print_operation_preview_end_preview

// Blacklisted : gtk_print_operation_preview_is_selected

// Blacklisted : gtk_print_operation_preview_render_page

type signalRecentChooserItemActivatedDetail struct {
	callback  RecentChooserSignalItemActivatedCallback
	handlerID C.gulong
}

var signalRecentChooserItemActivatedId int
var signalRecentChooserItemActivatedMap = make(map[int]signalRecentChooserItemActivatedDetail)
var signalRecentChooserItemActivatedLock sync.RWMutex

// RecentChooserSignalItemActivatedCallback is a callback function for a 'item-activated' signal emitted from a RecentChooser.
type RecentChooserSignalItemActivatedCallback func()

/*
ConnectItemActivated connects the callback to the 'item-activated' signal for the RecentChooser.

The returned value represents the connection, and may be passed to DisconnectItemActivated to remove it.
*/
func (recv *RecentChooser) ConnectItemActivated(callback RecentChooserSignalItemActivatedCallback) int {
	signalRecentChooserItemActivatedLock.Lock()
	defer signalRecentChooserItemActivatedLock.Unlock()

	signalRecentChooserItemActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.RecentChooser_signal_connect_item_activated(instance, C.gpointer(uintptr(signalRecentChooserItemActivatedId)))

	detail := signalRecentChooserItemActivatedDetail{callback, handlerID}
	signalRecentChooserItemActivatedMap[signalRecentChooserItemActivatedId] = detail

	return signalRecentChooserItemActivatedId
}

/*
DisconnectItemActivated disconnects a callback from the 'item-activated' signal for the RecentChooser.

The connectionID should be a value returned from a call to ConnectItemActivated.
*/
func (recv *RecentChooser) DisconnectItemActivated(connectionID int) {
	signalRecentChooserItemActivatedLock.Lock()
	defer signalRecentChooserItemActivatedLock.Unlock()

	detail, exists := signalRecentChooserItemActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalRecentChooserItemActivatedMap, connectionID)
}

//export recentchooser_itemActivatedHandler
func recentchooser_itemActivatedHandler(_ *C.GObject, data C.gpointer) {
	signalRecentChooserItemActivatedLock.RLock()
	defer signalRecentChooserItemActivatedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalRecentChooserItemActivatedMap[index].callback
	callback()
}

type signalRecentChooserSelectionChangedDetail struct {
	callback  RecentChooserSignalSelectionChangedCallback
	handlerID C.gulong
}

var signalRecentChooserSelectionChangedId int
var signalRecentChooserSelectionChangedMap = make(map[int]signalRecentChooserSelectionChangedDetail)
var signalRecentChooserSelectionChangedLock sync.RWMutex

// RecentChooserSignalSelectionChangedCallback is a callback function for a 'selection-changed' signal emitted from a RecentChooser.
type RecentChooserSignalSelectionChangedCallback func()

/*
ConnectSelectionChanged connects the callback to the 'selection-changed' signal for the RecentChooser.

The returned value represents the connection, and may be passed to DisconnectSelectionChanged to remove it.
*/
func (recv *RecentChooser) ConnectSelectionChanged(callback RecentChooserSignalSelectionChangedCallback) int {
	signalRecentChooserSelectionChangedLock.Lock()
	defer signalRecentChooserSelectionChangedLock.Unlock()

	signalRecentChooserSelectionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.RecentChooser_signal_connect_selection_changed(instance, C.gpointer(uintptr(signalRecentChooserSelectionChangedId)))

	detail := signalRecentChooserSelectionChangedDetail{callback, handlerID}
	signalRecentChooserSelectionChangedMap[signalRecentChooserSelectionChangedId] = detail

	return signalRecentChooserSelectionChangedId
}

/*
DisconnectSelectionChanged disconnects a callback from the 'selection-changed' signal for the RecentChooser.

The connectionID should be a value returned from a call to ConnectSelectionChanged.
*/
func (recv *RecentChooser) DisconnectSelectionChanged(connectionID int) {
	signalRecentChooserSelectionChangedLock.Lock()
	defer signalRecentChooserSelectionChangedLock.Unlock()

	detail, exists := signalRecentChooserSelectionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalRecentChooserSelectionChangedMap, connectionID)
}

//export recentchooser_selectionChangedHandler
func recentchooser_selectionChangedHandler(_ *C.GObject, data C.gpointer) {
	signalRecentChooserSelectionChangedLock.RLock()
	defer signalRecentChooserSelectionChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalRecentChooserSelectionChangedMap[index].callback
	callback()
}

// Blacklisted : gtk_recent_chooser_add_filter

// Blacklisted : gtk_recent_chooser_get_current_item

// Blacklisted : gtk_recent_chooser_get_current_uri

// Blacklisted : gtk_recent_chooser_get_filter

// Blacklisted : gtk_recent_chooser_get_items

// Blacklisted : gtk_recent_chooser_get_limit

// Blacklisted : gtk_recent_chooser_get_local_only

// Blacklisted : gtk_recent_chooser_get_select_multiple

// Blacklisted : gtk_recent_chooser_get_show_icons

// Blacklisted : gtk_recent_chooser_get_show_not_found

// Blacklisted : gtk_recent_chooser_get_show_private

// Blacklisted : gtk_recent_chooser_get_show_tips

// Blacklisted : gtk_recent_chooser_get_sort_type

// Blacklisted : gtk_recent_chooser_get_uris

// Blacklisted : gtk_recent_chooser_list_filters

// Blacklisted : gtk_recent_chooser_remove_filter

// Blacklisted : gtk_recent_chooser_select_all

// Blacklisted : gtk_recent_chooser_select_uri

// Blacklisted : gtk_recent_chooser_set_current_uri

// Blacklisted : gtk_recent_chooser_set_filter

// Blacklisted : gtk_recent_chooser_set_limit

// Blacklisted : gtk_recent_chooser_set_local_only

// Blacklisted : gtk_recent_chooser_set_select_multiple

// Blacklisted : gtk_recent_chooser_set_show_icons

// Blacklisted : gtk_recent_chooser_set_show_not_found

// Blacklisted : gtk_recent_chooser_set_show_private

// Blacklisted : gtk_recent_chooser_set_show_tips

// Unsupported : gtk_recent_chooser_set_sort_func : unsupported parameter sort_func : no type generator for RecentSortFunc (GtkRecentSortFunc) for param sort_func

// Blacklisted : gtk_recent_chooser_set_sort_type

// Blacklisted : gtk_recent_chooser_unselect_all

// Blacklisted : gtk_recent_chooser_unselect_uri
