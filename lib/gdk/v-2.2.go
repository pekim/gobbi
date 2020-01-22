// Code generated - DO NOT EDIT.
// +build gdk_2.2

package gdk

import (
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	pango "github.com/pekim/gobbi/lib/pango"
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

// GetDisplayArgName wraps the C function gdk_get_display_arg_name.
//
// since 2.2
func GetDisplayArgName() string {
	retSys := gdk.Fn_gdk_get_display_arg_name()
	ret := retSys

	return ret
}

// UNSUPPORTED : gdk_init : has array param, argv

// UNSUPPORTED : gdk_init_check : has array param, argv

// UNSUPPORTED : gdk_keyval_convert_case : has [in]out param, lower

// NotifyStartupComplete wraps the C function gdk_notify_startup_complete.
//
// since 2.2
func NotifyStartupComplete() {
	gdk.Fn_gdk_notify_startup_complete()
}

// PangoContextGetForScreen wraps the C function gdk_pango_context_get_for_screen.
//
// since 2.2
func PangoContextGetForScreen(screen *Screen) *pango.Context {
	sys_screen := screen.ToC()
	retSys := gdk.Fn_gdk_pango_context_get_for_screen(sys_screen)
	ret := pango.ContextNewFromC(retSys)

	return ret
}

// UNSUPPORTED : gdk_pango_layout_line_get_clip_region : parameter 'index_ranges' is array parameter without length parameter

// UNSUPPORTED : gdk_parse_args : has array param, argv

// UNSUPPORTED : gdk_property_get : has [in]out param, actual_property_type

// UNSUPPORTED : gdk_query_depths : has array param, depths

// UNSUPPORTED : gdk_query_visual_types : has array param, visual_types

// SelectionOwnerGetForDisplay wraps the C function gdk_selection_owner_get_for_display.
//
// since 2.2
func SelectionOwnerGetForDisplay(display *Display, selection Atom) *Window {
	sys_display := display.ToC()
	sys_selection := selection.ToC()
	retSys := gdk.Fn_gdk_selection_owner_get_for_display(sys_display, sys_selection)
	ret := WindowNewFromC(retSys)

	return ret
}

// SelectionOwnerSetForDisplay wraps the C function gdk_selection_owner_set_for_display.
//
// since 2.2
func SelectionOwnerSetForDisplay(display *Display, owner *Window, selection Atom, time uint32, sendEvent bool) bool {
	sys_display := display.ToC()
	sys_owner := owner.ToC()
	sys_selection := selection.ToC()
	sys_time := time
	sys_sendEvent := sendEvent
	retSys := gdk.Fn_gdk_selection_owner_set_for_display(sys_display, sys_owner, sys_selection, sys_time, sys_sendEvent)
	ret := retSys

	return ret
}

// SelectionSendNotifyForDisplay wraps the C function gdk_selection_send_notify_for_display.
//
// since 2.2
func SelectionSendNotifyForDisplay(display *Display, requestor *Window, selection Atom, target Atom, property Atom, time uint32) {
	sys_display := display.ToC()
	sys_requestor := requestor.ToC()
	sys_selection := selection.ToC()
	sys_target := target.ToC()
	sys_property := property.ToC()
	sys_time := time
	gdk.Fn_gdk_selection_send_notify_for_display(sys_display, sys_requestor, sys_selection, sys_target, sys_property, sys_time)
}

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
