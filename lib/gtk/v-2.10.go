// Code generated - DO NOT EDIT.
// +build gtk_2.10

package gtk

import (
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
	"unsafe"
)

// STOCK_ORIENTATION_LANDSCAPE is a representation of the C constant GTK_STOCK_ORIENTATION_LANDSCAPE.
//
// since 2.10
const STOCK_ORIENTATION_LANDSCAPE = "gtk-orientation-landscape"

// STOCK_ORIENTATION_PORTRAIT is a representation of the C constant GTK_STOCK_ORIENTATION_PORTRAIT.
//
// since 2.10
const STOCK_ORIENTATION_PORTRAIT = "gtk-orientation-portrait"

// STOCK_ORIENTATION_REVERSE_LANDSCAPE is a representation of the C constant GTK_STOCK_ORIENTATION_REVERSE_LANDSCAPE.
//
// since 2.10
const STOCK_ORIENTATION_REVERSE_LANDSCAPE = "gtk-orientation-reverse-landscape"

// STOCK_ORIENTATION_REVERSE_PORTRAIT is a representation of the C constant GTK_STOCK_ORIENTATION_REVERSE_PORTRAIT.
//
// since 2.10
const STOCK_ORIENTATION_REVERSE_PORTRAIT = "gtk-orientation-reverse-portrait"

// STOCK_SELECT_ALL is a representation of the C constant GTK_STOCK_SELECT_ALL.
//
// since 2.10
const STOCK_SELECT_ALL = "gtk-select-all"

// RecentChooserError is a representation of the C enumeration GtkRecentChooserError.
type RecentChooserError int

// RecentChooserError_not_found is a representation of the C enumeration member GTK_RECENT_CHOOSER_ERROR_NOT_FOUND.
const RecentChooserError_not_found = RecentChooserError(0)

// RecentChooserError_invalid_uri is a representation of the C enumeration member GTK_RECENT_CHOOSER_ERROR_INVALID_URI.
const RecentChooserError_invalid_uri = RecentChooserError(1)

// RecentManagerError is a representation of the C enumeration GtkRecentManagerError.
type RecentManagerError int

// RecentManagerError_not_found is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_NOT_FOUND.
const RecentManagerError_not_found = RecentManagerError(0)

// RecentManagerError_invalid_uri is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_INVALID_URI.
const RecentManagerError_invalid_uri = RecentManagerError(1)

// RecentManagerError_invalid_encoding is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING.
const RecentManagerError_invalid_encoding = RecentManagerError(2)

// RecentManagerError_not_registered is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED.
const RecentManagerError_not_registered = RecentManagerError(3)

// RecentManagerError_read is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_READ.
const RecentManagerError_read = RecentManagerError(4)

// RecentManagerError_write is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_WRITE.
const RecentManagerError_write = RecentManagerError(5)

// RecentManagerError_unknown is a representation of the C enumeration member GTK_RECENT_MANAGER_ERROR_UNKNOWN.
const RecentManagerError_unknown = RecentManagerError(6)

// RecentSortType is a representation of the C enumeration GtkRecentSortType.
type RecentSortType int

// RecentSortType_none is a representation of the C enumeration member GTK_RECENT_SORT_NONE.
const RecentSortType_none = RecentSortType(0)

// RecentSortType_mru is a representation of the C enumeration member GTK_RECENT_SORT_MRU.
const RecentSortType_mru = RecentSortType(1)

// RecentSortType_lru is a representation of the C enumeration member GTK_RECENT_SORT_LRU.
const RecentSortType_lru = RecentSortType(2)

// RecentSortType_custom is a representation of the C enumeration member GTK_RECENT_SORT_CUSTOM.
const RecentSortType_custom = RecentSortType(3)

// UNSUPPORTED : gtk_accelerator_parse : has [in]out param, accelerator_key

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// UNSUPPORTED : gtk_get_current_event_state : has [in]out param, state

// UNSUPPORTED : gtk_icon_size_lookup : has [in]out param, width

// UNSUPPORTED : gtk_icon_size_lookup_for_settings : has [in]out param, width

// UNSUPPORTED : gtk_init : has array param, argv

// UNSUPPORTED : gtk_init_check : has array param, argv

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_parse_args : has array param, argv

// PrintRunPageSetupDialog wraps the C function gtk_print_run_page_setup_dialog.
//
// since 2.10
func PrintRunPageSetupDialog(parent *Window, pageSetup *PageSetup, settings *PrintSettings) *PageSetup {
	sys_parent := parent.ToC()
	sys_pageSetup := pageSetup.ToC()
	sys_settings := settings.ToC()
	retSys := gtk.Fn_gtk_print_run_page_setup_dialog(sys_parent, sys_pageSetup, sys_settings)
	ret := PageSetupNewFromC(retSys)

	return ret
}

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_parse_color : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_color_full : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_state : has [in]out param, state

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

// UNSUPPORTED : gtk_render_background_get_clip : has [in]out param, out_clip

// UNSUPPORTED : gtk_rgb_to_hsv : has [in]out param, h

// UNSUPPORTED : gtk_selection_add_targets : has array param, targets

// UNSUPPORTED : gtk_show_uri : throws

// UNSUPPORTED : gtk_show_uri_on_window : throws

// UNSUPPORTED : gtk_stock_add : has array param, items

// UNSUPPORTED : gtk_stock_add_static : has array param, items

// UNSUPPORTED : gtk_stock_lookup : has [in]out param, item

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_target_table_free : has array param, targets

// UNSUPPORTED : gtk_target_table_new_from_list : has [in]out param, n_targets

// UNSUPPORTED : gtk_targets_include_image : has array param, targets

// UNSUPPORTED : gtk_targets_include_rich_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_uri : has array param, targets

// UNSUPPORTED : gtk_test_init : has array param, argvp

// UNSUPPORTED : gtk_test_list_all_types : has [in]out param, n_types

// UNSUPPORTED : gtk_tree_get_row_drag_data : has [in]out param, tree_model

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : EventControllerMotionClass : blacklisted

// UNSUPPORTED : EventControllerScrollClass : blacklisted

// UNSUPPORTED : GestureStylusClass : blacklisted

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted

// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted

// RecentInfo is a representation of the C record GtkRecentInfo.
//
// since 2.10
type RecentInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentInfo that represents the RecentInfo.
func (recv *RecentInfo) ToC() unsafe.Pointer {
	return recv.native
}

// RecentInfoNewFromC creates a new RecentInfo from a pointer to the C GtkRecentInfo that represents the RecentInfo.
func RecentInfoNewFromC(native unsafe.Pointer) *RecentInfo {
	return &RecentInfo{native: native}
}

// RecentManagerClass is a representation of the C record GtkRecentManagerClass.
//
// since 2.10
type RecentManagerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentManagerClass that represents the RecentManagerClass.
func (recv *RecentManagerClass) ToC() unsafe.Pointer {
	return recv.native
}

// RecentManagerClassNewFromC creates a new RecentManagerClass from a pointer to the C GtkRecentManagerClass that represents the RecentManagerClass.
func RecentManagerClassNewFromC(native unsafe.Pointer) *RecentManagerClass {
	return &RecentManagerClass{native: native}
}

// UNSUPPORTED : StackAccessibleClass : blacklisted

// UNSUPPORTED : EntryIconAccessible : blacklisted

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// UNSUPPORTED : GestureStylus : blacklisted

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// RecentManager is a representation of the C record GtkRecentManager.
//
// since 2.10
type RecentManager struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkRecentManager that represents the RecentManager.
func (recv *RecentManager) ToC() unsafe.Pointer {
	return recv.native
}

// RecentManagerNewFromC creates a new RecentManager from a pointer to the C GtkRecentManager that represents the RecentManager.
func RecentManagerNewFromC(native unsafe.Pointer) *RecentManager {
	return &RecentManager{native: native}
}

// UNSUPPORTED : StackAccessible : blacklisted
