// Code generated - DO NOT EDIT.
// +build gdk_2.14

package gdk

import (
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	"unsafe"
)

// UNSUPPORTED : XEvent : blacklisted

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

// TestRenderSync wraps the C function gdk_test_render_sync.
//
// since 2.14
func TestRenderSync(window *Window) {
	sys_window := window.ToC()
	gdk.Fn_gdk_test_render_sync(sys_window)
}

// TestSimulateButton wraps the C function gdk_test_simulate_button.
//
// since 2.14
func TestSimulateButton(window *Window, x int, y int, button uint, modifiers ModifierType, buttonPressrelease EventType) bool {
	sys_window := window.ToC()
	sys_x := x
	sys_y := y
	sys_button := button
	sys_modifiers := (int)(modifiers)
	sys_buttonPressrelease := (int)(buttonPressrelease)
	retSys := gdk.Fn_gdk_test_simulate_button(sys_window, sys_x, sys_y, sys_button, sys_modifiers, sys_buttonPressrelease)
	ret := retSys

	return ret
}

// TestSimulateKey wraps the C function gdk_test_simulate_key.
//
// since 2.14
func TestSimulateKey(window *Window, x int, y int, keyval uint, modifiers ModifierType, keyPressrelease EventType) bool {
	sys_window := window.ToC()
	sys_x := x
	sys_y := y
	sys_keyval := keyval
	sys_modifiers := (int)(modifiers)
	sys_keyPressrelease := (int)(keyPressrelease)
	retSys := gdk.Fn_gdk_test_simulate_key(sys_window, sys_x, sys_y, sys_keyval, sys_modifiers, sys_keyPressrelease)
	ret := retSys

	return ret
}

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
