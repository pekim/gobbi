// This is a generated file - DO NOT EDIT
// +build gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File (GFile*) for param file

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no type generator for GType (GType) for array param types

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup (GActionGroup*) for param group

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported signal 'show-starred-location' for PlacesSidebar : unsupported parameter object : type PlacesOpenFlags :

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

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no type generator for GType (GType) for array param types

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_widget_new : unsupported parameter ... : varargs
