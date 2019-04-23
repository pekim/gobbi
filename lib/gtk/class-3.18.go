// This is a generated file - DO NOT EDIT
// +build gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

	void placessidebar_showOtherLocationsHandler(GObject *, gpointer);

	static gulong PlacesSidebar_signal_connect_show_other_locations(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-other-locations", G_CALLBACK(placessidebar_showOtherLocationsHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_assistant_get_page_has_padding

// Blacklisted : gtk_assistant_set_page_has_padding

// Blacklisted : gtk_container_child_notify_by_pspec

// Unsupported : gtk_flow_box_bind_model : unsupported parameter create_widget_func : no type generator for FlowBoxCreateWidgetFunc (GtkFlowBoxCreateWidgetFunc) for param create_widget_func

// Blacklisted : gtk_overlay_get_overlay_pass_through

// Blacklisted : gtk_overlay_reorder_overlay

// Blacklisted : gtk_overlay_set_overlay_pass_through

type signalPlacesSidebarShowOtherLocationsDetail struct {
	callback  PlacesSidebarSignalShowOtherLocationsCallback
	handlerID C.gulong
}

var signalPlacesSidebarShowOtherLocationsId int
var signalPlacesSidebarShowOtherLocationsMap = make(map[int]signalPlacesSidebarShowOtherLocationsDetail)
var signalPlacesSidebarShowOtherLocationsLock sync.RWMutex

// PlacesSidebarSignalShowOtherLocationsCallback is a callback function for a 'show-other-locations' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowOtherLocationsCallback func()

/*
ConnectShowOtherLocations connects the callback to the 'show-other-locations' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectShowOtherLocations to remove it.
*/
func (recv *PlacesSidebar) ConnectShowOtherLocations(callback PlacesSidebarSignalShowOtherLocationsCallback) int {
	signalPlacesSidebarShowOtherLocationsLock.Lock()
	defer signalPlacesSidebarShowOtherLocationsLock.Unlock()

	signalPlacesSidebarShowOtherLocationsId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_show_other_locations(instance, C.gpointer(uintptr(signalPlacesSidebarShowOtherLocationsId)))

	detail := signalPlacesSidebarShowOtherLocationsDetail{callback, handlerID}
	signalPlacesSidebarShowOtherLocationsMap[signalPlacesSidebarShowOtherLocationsId] = detail

	return signalPlacesSidebarShowOtherLocationsId
}

/*
DisconnectShowOtherLocations disconnects a callback from the 'show-other-locations' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectShowOtherLocations.
*/
func (recv *PlacesSidebar) DisconnectShowOtherLocations(connectionID int) {
	signalPlacesSidebarShowOtherLocationsLock.Lock()
	defer signalPlacesSidebarShowOtherLocationsLock.Unlock()

	detail, exists := signalPlacesSidebarShowOtherLocationsMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarShowOtherLocationsMap, connectionID)
}

//export placessidebar_showOtherLocationsHandler
func placessidebar_showOtherLocationsHandler(_ *C.GObject, data C.gpointer) {
	signalPlacesSidebarShowOtherLocationsLock.RLock()
	defer signalPlacesSidebarShowOtherLocationsLock.RUnlock()

	index := int(uintptr(data))
	callback := signalPlacesSidebarShowOtherLocationsMap[index].callback
	callback()
}

// Blacklisted : gtk_places_sidebar_get_show_other_locations

// Blacklisted : gtk_places_sidebar_get_show_recent

// Blacklisted : gtk_places_sidebar_get_show_trash

// Blacklisted : gtk_places_sidebar_set_drop_targets_visible

// Blacklisted : gtk_places_sidebar_set_show_other_locations

// Blacklisted : gtk_places_sidebar_set_show_recent

// Blacklisted : gtk_places_sidebar_set_show_trash

// Blacklisted : gtk_popover_get_default_widget

// Blacklisted : gtk_popover_set_default_widget

// Blacklisted : gtk_radio_menu_item_join_group

// Blacklisted : gtk_stack_get_interpolate_size

// Blacklisted : gtk_stack_set_interpolate_size

// Blacklisted : gtk_text_view_get_bottom_margin

// Blacklisted : gtk_text_view_get_top_margin

// Blacklisted : gtk_text_view_set_bottom_margin

// Blacklisted : gtk_text_view_set_top_margin

// Blacklisted : gtk_widget_get_font_map

// Blacklisted : gtk_widget_get_font_options

// Blacklisted : gtk_widget_set_font_map

// Blacklisted : gtk_widget_set_font_options

// Blacklisted : gtk_window_fullscreen_on_monitor
