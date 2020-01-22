// Code generated - DO NOT EDIT.
// +build gtk_3.24

package gtk

import "unsafe"

// EventControllerScrollFlags is a representation of the C bitfield GtkEventControllerScrollFlags.
type EventControllerScrollFlags int

// EventControllerScrollFlags_none is a representation of the C bitfield member GTK_EVENT_CONTROLLER_SCROLL_NONE.
const EventControllerScrollFlags_none = EventControllerScrollFlags(0)

// EventControllerScrollFlags_vertical is a representation of the C bitfield member GTK_EVENT_CONTROLLER_SCROLL_VERTICAL.
const EventControllerScrollFlags_vertical = EventControllerScrollFlags(1)

// EventControllerScrollFlags_horizontal is a representation of the C bitfield member GTK_EVENT_CONTROLLER_SCROLL_HORIZONTAL.
const EventControllerScrollFlags_horizontal = EventControllerScrollFlags(2)

// EventControllerScrollFlags_discrete is a representation of the C bitfield member GTK_EVENT_CONTROLLER_SCROLL_DISCRETE.
const EventControllerScrollFlags_discrete = EventControllerScrollFlags(4)

// EventControllerScrollFlags_kinetic is a representation of the C bitfield member GTK_EVENT_CONTROLLER_SCROLL_KINETIC.
const EventControllerScrollFlags_kinetic = EventControllerScrollFlags(8)

// EventControllerScrollFlags_both_axes is a representation of the C bitfield member GTK_EVENT_CONTROLLER_SCROLL_BOTH_AXES.
const EventControllerScrollFlags_both_axes = EventControllerScrollFlags(3)

// FontChooserLevel is a representation of the C bitfield GtkFontChooserLevel.
type FontChooserLevel int

// FontChooserLevel_family is a representation of the C bitfield member GTK_FONT_CHOOSER_LEVEL_FAMILY.
const FontChooserLevel_family = FontChooserLevel(0)

// FontChooserLevel_style is a representation of the C bitfield member GTK_FONT_CHOOSER_LEVEL_STYLE.
const FontChooserLevel_style = FontChooserLevel(1)

// FontChooserLevel_size is a representation of the C bitfield member GTK_FONT_CHOOSER_LEVEL_SIZE.
const FontChooserLevel_size = FontChooserLevel(2)

// FontChooserLevel_variations is a representation of the C bitfield member GTK_FONT_CHOOSER_LEVEL_VARIATIONS.
const FontChooserLevel_variations = FontChooserLevel(4)

// FontChooserLevel_features is a representation of the C bitfield member GTK_FONT_CHOOSER_LEVEL_FEATURES.
const FontChooserLevel_features = FontChooserLevel(8)

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

// EventControllerKeyClass is a representation of the C record GtkEventControllerKeyClass.
//
// since 3.24
type EventControllerKeyClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkEventControllerKeyClass that represents the EventControllerKeyClass.
func (recv *EventControllerKeyClass) ToC() unsafe.Pointer {
	return recv.native
}

// EventControllerKeyClassNewFromC creates a new EventControllerKeyClass from a pointer to the C GtkEventControllerKeyClass that represents the EventControllerKeyClass.
func EventControllerKeyClassNewFromC(native unsafe.Pointer) *EventControllerKeyClass {
	return &EventControllerKeyClass{native: native}
}

// UNSUPPORTED : EventControllerMotionClass : blacklisted

// UNSUPPORTED : EventControllerScrollClass : blacklisted

// UNSUPPORTED : GestureStylusClass : blacklisted

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted

// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted

// UNSUPPORTED : StackAccessibleClass : blacklisted

// UNSUPPORTED : EntryIconAccessible : blacklisted

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// UNSUPPORTED : GestureStylus : blacklisted

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// UNSUPPORTED : StackAccessible : blacklisted
