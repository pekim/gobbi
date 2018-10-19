// This is a generated file - DO NOT EDIT
// +build gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"fmt"
	"sync"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void PlacesSidebar_showStarredLocationHandler();

	static gulong PlacesSidebar_signal_connect_show_starred_location(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-starred-location", PlacesSidebar_showStarredLocationHandler, data);
	}

*/
import "C"

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported signal 'application-activated' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'application-selected' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'populate-popup' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported signal 'add-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'apply-attributes' for CellArea : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'remove-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'editing-started' for CellRenderer : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'format-entry-text' for ComboBox : return value utf8 :

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported signal 'cursor-on-match' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'match-selected' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported signal 'get-child-position' for Overlay : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported signal 'drag-action-ask' for PlacesSidebar : return value gint :

// Unsupported signal 'drag-action-requested' for PlacesSidebar : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal 'drag-perform-drop' for PlacesSidebar : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal 'open-location' for PlacesSidebar : unsupported parameter location : no type generator for Gio.File,

// Unsupported signal 'populate-popup' for PlacesSidebar : unsupported parameter selected_item : no type generator for Gio.File,

var signalPlacesSidebarShowStarredLocationId int
var signalPlacesSidebarShowStarredLocationMap = make(map[int]PlacesSidebarSignalShowStarredLocationCallback)
var signalPlacesSidebarShowStarredLocationLock sync.Mutex

// PlacesSidebarSignalShowStarredLocationCallback is a callback function for a 'show-starred-location' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowStarredLocationCallback func(object PlacesOpenFlags)

/*
ConnectShowStarredLocation connects the callback to the 'show-starred-location' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectShowStarredLocation to remove it.
*/
func (recv *PlacesSidebar) ConnectShowStarredLocation(callback PlacesSidebarSignalShowStarredLocationCallback) int {
	signalPlacesSidebarShowStarredLocationLock.Lock()
	defer signalPlacesSidebarShowStarredLocationLock.Unlock()

	signalPlacesSidebarShowStarredLocationId++
	signalPlacesSidebarShowStarredLocationMap[signalPlacesSidebarShowStarredLocationId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.PlacesSidebar_signal_connect_show_starred_location(instance, C.gpointer(uintptr(signalPlacesSidebarShowStarredLocationId)))
	return int(retC)
}

/*
DisconnectShowStarredLocation disconnects a callback from the 'show-starred-location' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectShowStarredLocation.
*/
func (recv *PlacesSidebar) DisconnectShowStarredLocation(connectionID int) {
	signalPlacesSidebarShowStarredLocationLock.Lock()
	defer signalPlacesSidebarShowStarredLocationLock.Unlock()

	_, exists := signalPlacesSidebarShowStarredLocationMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalPlacesSidebarShowStarredLocationMap, connectionID)
}

//export PlacesSidebar_showStarredLocationHandler
func PlacesSidebar_showStarredLocationHandler() {
	fmt.Println("cb")
}

// GetShowStarredLocation is a wrapper around the C function gtk_places_sidebar_get_show_starred_location.
func (recv *PlacesSidebar) GetShowStarredLocation() bool {
	retC := C.gtk_places_sidebar_get_show_starred_location((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowStarredLocation is a wrapper around the C function gtk_places_sidebar_set_show_starred_location.
func (recv *PlacesSidebar) SetShowStarredLocation(showStarredLocation bool) {
	c_show_starred_location :=
		boolToGboolean(showStarredLocation)

	C.gtk_places_sidebar_set_show_starred_location((*C.GtkPlacesSidebar)(recv.native), c_show_starred_location)

	return
}

// Unsupported signal 'preview' for PrintOperation : unsupported parameter preview : no type generator for PrintOperationPreview,

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported signal 'format-value' for Scale : return value utf8 :

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported signal 'input' for SpinButton : return value gint :

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported signal 'event' for TextTag : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported signal 'child-notify' for Widget : unsupported parameter child_property : Blacklisted record : GParamSpec

// Unsupported signal 'delete-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'destroy-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event-after' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'size-allocate' for Widget : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported signal 'touch-event' for Widget : unsupported parameter object : no type generator for Gdk.Event,

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType
