// This is a generated file - DO NOT EDIT
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// Unsupported : gtk_accelerator_parse_with_keycode : unsupported parameter accelerator_codes : no type generator for guint (guint*) for array param accelerator_codes

// BindingEntrySkip is a wrapper around the C function gtk_binding_entry_skip.
func BindingEntrySkip(bindingSet *BindingSet, keyval uint32, modifiers gdk.ModifierType) {
	c_binding_set := (*C.GtkBindingSet)(bindingSet.ToC())

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gtk_binding_entry_skip(c_binding_set, c_keyval, c_modifiers)

	return
}

// Unsupported : gtk_drag_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// Unsupported : gtk_draw_insertion_cursor : unsupported parameter location : Blacklisted record : GdkRectangle

// Unsupported : gtk_get_current_event : no return generator

// Unsupported : gtk_get_current_event_state : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_icon_size_from_name : no return generator

// Unsupported : gtk_icon_size_get_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_icon_size_lookup : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_icon_size_lookup_for_settings : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_icon_size_register : no return generator

// Unsupported : gtk_icon_size_register_alias : unsupported parameter target : no type generator for gint (GtkIconSize) for param target

// Unsupported : gtk_init : unsupported parameter argv : no type generator for utf8 (char**) for array param argv

// Unsupported : gtk_init_check : unsupported parameter argv : no type generator for utf8 (char**) for array param argv

// Unsupported : gtk_init_with_args : unsupported parameter argv : no type generator for utf8 (gchar**) for array param argv

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc (GtkKeySnoopFunc) for param snooper

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// PaperSizeGetPaperSizes is a wrapper around the C function gtk_paper_size_get_paper_sizes.
func PaperSizeGetPaperSizes(includeCustom bool) *glib.List {
	c_include_custom :=
		boolToGboolean(includeCustom)

	retC := C.gtk_paper_size_get_paper_sizes(c_include_custom)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_parse_args : unsupported parameter argv : no type generator for utf8 (char**) for array param argv

// Unsupported : gtk_print_run_page_setup_dialog_async : unsupported parameter done_cb : no type generator for PageSetupDoneFunc (GtkPageSetupDoneFunc) for param done_cb

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_rc_get_default_files : no return type

// RcParseColorFull is a wrapper around the C function gtk_rc_parse_color_full.
func RcParseColorFull(scanner *glib.Scanner, style *RcStyle) (uint32, *gdk.Color) {
	c_scanner := (*C.GScanner)(scanner.ToC())

	c_style := (*C.GtkRcStyle)(style.ToC())

	var c_color C.GdkColor

	retC := C.gtk_rc_parse_color_full(c_scanner, c_style, &c_color)
	retGo := (uint32)(retC)

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// Unsupported : gtk_rc_parse_priority : unsupported parameter priority : GtkPathPriorityType* with indirection level of 1

// Unsupported : gtk_rc_parse_state : unsupported parameter state : GtkStateType* with indirection level of 1

// Unsupported : gtk_rc_property_parse_border : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_color : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_enum : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_flags : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_requisition : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_set_default_files : unsupported parameter filenames : no type generator for filename () for array param filenames

// Unsupported : gtk_render_background_get_clip : unsupported parameter out_clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_render_icon_pixbuf : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_selection_add_target : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_add_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_clear_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_show_about_dialog : unsupported parameter ... : varargs

// Unsupported : gtk_stock_add : unsupported parameter items : no type generator for StockItem (GtkStockItem) for array param items

// Unsupported : gtk_stock_add_static : unsupported parameter items : no type generator for StockItem (GtkStockItem) for array param items

// Unsupported : gtk_stock_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GtkTranslateFunc) for param func

// Unsupported : gtk_target_table_free : unsupported parameter targets : no type generator for TargetEntry (GtkTargetEntry) for array param targets

// Unsupported : gtk_target_table_new_from_list : no return type

// Unsupported : gtk_targets_include_image : unsupported parameter targets : no type generator for Gdk.Atom (GdkAtom) for array param targets

// Unsupported : gtk_targets_include_rich_text : unsupported parameter targets : no type generator for Gdk.Atom (GdkAtom) for array param targets

// Unsupported : gtk_targets_include_text : unsupported parameter targets : no type generator for Gdk.Atom (GdkAtom) for array param targets

// Unsupported : gtk_targets_include_uri : unsupported parameter targets : no type generator for Gdk.Atom (GdkAtom) for array param targets

// Unsupported : gtk_test_create_widget : unsupported parameter ... : varargs

// Unsupported : gtk_test_display_button_window : unsupported parameter ... : varargs

// Unsupported : gtk_test_init : unsupported parameter argvp : no type generator for utf8 (char**) for array param argvp

// Unsupported : gtk_test_list_all_types : no return type

// Unsupported : gtk_tree_get_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel (GtkTreeModel**) for param tree_model

// Unsupported : gtk_tree_row_reference_reordered : unsupported parameter new_order : no type generator for gint (gint) for array param new_order

// Unsupported : gtk_tree_set_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel (GtkTreeModel*) for param tree_model
