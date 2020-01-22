// Code generated - DO NOT EDIT.
// +build gdk_3.16

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	"unsafe"
)

// UNSUPPORTED : XEvent : blacklisted

// GLError is a representation of the C enumeration GdkGLError.
type GLError int

// GLError_not_available is a representation of the C enumeration member GDK_GL_ERROR_NOT_AVAILABLE.
const GLError_not_available = GLError(0)

// GLError_unsupported_format is a representation of the C enumeration member GDK_GL_ERROR_UNSUPPORTED_FORMAT.
const GLError_unsupported_format = GLError(1)

// GLError_unsupported_profile is a representation of the C enumeration member GDK_GL_ERROR_UNSUPPORTED_PROFILE.
const GLError_unsupported_profile = GLError(2)

// CairoDrawFromGl wraps the C function gdk_cairo_draw_from_gl.
//
// since 3.16
func CairoDrawFromGl(cr *cairo.Context, window *Window, source int, sourceType int, bufferScale int, x int, y int, width int, height int) {
	sys_cr := cr.ToC()
	sys_window := window.ToC()
	sys_source := source
	sys_sourceType := sourceType
	sys_bufferScale := bufferScale
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gdk.Fn_gdk_cairo_draw_from_gl(sys_cr, sys_window, sys_source, sys_sourceType, sys_bufferScale, sys_x, sys_y, sys_width, sys_height)
}

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
