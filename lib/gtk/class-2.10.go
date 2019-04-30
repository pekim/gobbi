// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// Unsupported : gtk_assistant_new : return type :

// Unsupported : gtk_cell_renderer_accel_new : return type :

// Unsupported : gtk_cell_renderer_spin_new : return type :

// Unsupported : gtk_link_button_new : return type :

// Unsupported : gtk_link_button_new_with_label : return type :

// Unsupported : gtk_page_setup_new : return type :

// Unsupported : gtk_print_operation_new : return type :

// Unsupported : gtk_print_settings_new : return type :

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_menu_new : return type :

// Unsupported : gtk_recent_chooser_menu_new_for_manager : return type :

// Unsupported : gtk_recent_chooser_widget_new : return type :

// Unsupported : gtk_recent_chooser_widget_new_for_manager : return type :

// Unsupported : gtk_recent_filter_new : return type :

// RecentManager is a wrapper around the C record GtkRecentManager.
type RecentManager struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func RecentManagerNewFromC(u unsafe.Pointer) *RecentManager {
	if u == nil {
		return nil
	}

	g := &RecentManager{native: u}

	return g
}

func (recv *RecentManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_recent_manager_new : return type :

// Unsupported : gtk_status_icon_new : return type :

// Unsupported : gtk_status_icon_new_from_file : return type :

// Unsupported : gtk_status_icon_new_from_icon_name : return type :

// Unsupported : gtk_status_icon_new_from_pixbuf : return type :

// Unsupported : gtk_status_icon_new_from_stock : return type :
