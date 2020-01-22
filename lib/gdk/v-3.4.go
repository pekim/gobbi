// Code generated - DO NOT EDIT.
// +build gdk_3.4

package gdk

import "unsafe"

// UNSUPPORTED : XEvent : blacklisted

// BUTTON_MIDDLE is a representation of the C constant GDK_BUTTON_MIDDLE.
//
// since 3.4
const BUTTON_MIDDLE = 2

// BUTTON_PRIMARY is a representation of the C constant GDK_BUTTON_PRIMARY.
//
// since 3.4
const BUTTON_PRIMARY = 1

// BUTTON_SECONDARY is a representation of the C constant GDK_BUTTON_SECONDARY.
//
// since 3.4
const BUTTON_SECONDARY = 3

// EVENT_PROPAGATE is a representation of the C constant GDK_EVENT_PROPAGATE.
//
// since 3.4
const EVENT_PROPAGATE = false

// EVENT_STOP is a representation of the C constant GDK_EVENT_STOP.
//
// since 3.4
const EVENT_STOP = true

// ModifierIntent is a representation of the C enumeration GdkModifierIntent.
type ModifierIntent int

// ModifierIntent_primary_accelerator is a representation of the C enumeration member GDK_MODIFIER_INTENT_PRIMARY_ACCELERATOR.
const ModifierIntent_primary_accelerator = ModifierIntent(0)

// ModifierIntent_context_menu is a representation of the C enumeration member GDK_MODIFIER_INTENT_CONTEXT_MENU.
const ModifierIntent_context_menu = ModifierIntent(1)

// ModifierIntent_extend_selection is a representation of the C enumeration member GDK_MODIFIER_INTENT_EXTEND_SELECTION.
const ModifierIntent_extend_selection = ModifierIntent(2)

// ModifierIntent_modify_selection is a representation of the C enumeration member GDK_MODIFIER_INTENT_MODIFY_SELECTION.
const ModifierIntent_modify_selection = ModifierIntent(3)

// ModifierIntent_no_text_input is a representation of the C enumeration member GDK_MODIFIER_INTENT_NO_TEXT_INPUT.
const ModifierIntent_no_text_input = ModifierIntent(4)

// ModifierIntent_shift_group is a representation of the C enumeration member GDK_MODIFIER_INTENT_SHIFT_GROUP.
const ModifierIntent_shift_group = ModifierIntent(5)

// ModifierIntent_default_mod_mask is a representation of the C enumeration member GDK_MODIFIER_INTENT_DEFAULT_MOD_MASK.
const ModifierIntent_default_mod_mask = ModifierIntent(6)

// UNSUPPORTED : gdk_cairo_get_clip_rectangle : has [in]out param, rect

// UNSUPPORTED : gdk_color_parse : has [in]out param, color

// UNSUPPORTED : gdk_drag_find_window_for_screen : has [in]out param, dest_window

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

// UNSUPPORTED : gdk_events_get_angle : has [in]out param, angle

// UNSUPPORTED : gdk_events_get_center : has [in]out param, x

// UNSUPPORTED : gdk_events_get_distance : has [in]out param, distance

// UNSUPPORTED : gdk_init : has array param, argv

// UNSUPPORTED : gdk_init_check : has array param, argv

// UNSUPPORTED : gdk_keyval_convert_case : has [in]out param, lower

// UNSUPPORTED : gdk_pango_layout_line_get_clip_region : parameter 'index_ranges' is array parameter without length parameter

// UNSUPPORTED : gdk_parse_args : has array param, argv

// UNSUPPORTED : gdk_property_get : has [in]out param, actual_property_type

// UNSUPPORTED : gdk_query_depths : has array param, depths

// UNSUPPORTED : gdk_query_visual_types : has array param, visual_types

// UNSUPPORTED : gdk_synthesize_window_state : blacklisted

// UNSUPPORTED : gdk_text_property_to_utf8_list_for_display : parameter 'list' is array parameter without length parameter

// UNSUPPORTED : gdk_threads_add_idle : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_idle_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_set_lock_functions : parameter 'enter_fn' is callback

// Event is a representation of the C union GdkEvent.
type Event struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEvent that represents the Event.
func (recv *Event) ToC() unsafe.Pointer {
	return recv.native
}

// EventNewFromC creates a new Event from a pointer to the C GdkEvent that represents the Event.
func EventNewFromC(native unsafe.Pointer) *Event {
	return &Event{native: native}
}
