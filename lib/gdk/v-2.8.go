// Code generated - DO NOT EDIT.
// +build gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22 gdk_3.24

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	"unsafe"
)

// UNSUPPORTED : XEvent : blacklisted

// CairoCreate wraps the C function gdk_cairo_create.
//
// since 2.8
func CairoCreate(window *Window) *cairo.Context {
	sys_window := window.ToC()
	retSys := gdk.Fn_gdk_cairo_create(sys_window)
	ret := cairo.ContextNewFromC(retSys)

	return ret
}

// UNSUPPORTED : gdk_cairo_get_clip_rectangle : has [in]out param, rect

// CairoRectangle wraps the C function gdk_cairo_rectangle.
//
// since 2.8
func CairoRectangle(cr *cairo.Context, rectangle *Rectangle) {
	sys_cr := cr.ToC()
	sys_rectangle := rectangle.ToC()
	gdk.Fn_gdk_cairo_rectangle(sys_cr, sys_rectangle)
}

// CairoRegion wraps the C function gdk_cairo_region.
//
// since 2.8
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	sys_cr := cr.ToC()
	sys_region := region.ToC()
	gdk.Fn_gdk_cairo_region(sys_cr, sys_region)
}

// CairoSetSourceColor wraps the C function gdk_cairo_set_source_color.
//
// since 2.8
func CairoSetSourceColor(cr *cairo.Context, color *Color) {
	sys_cr := cr.ToC()
	sys_color := color.ToC()
	gdk.Fn_gdk_cairo_set_source_color(sys_cr, sys_color)
}

// CairoSetSourcePixbuf wraps the C function gdk_cairo_set_source_pixbuf.
//
// since 2.8
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, pixbufX float64, pixbufY float64) {
	sys_cr := cr.ToC()
	sys_pixbuf := pixbuf.ToC()
	sys_pixbufX := pixbufX
	sys_pixbufY := pixbufY
	gdk.Fn_gdk_cairo_set_source_pixbuf(sys_cr, sys_pixbuf, sys_pixbufX, sys_pixbufY)
}

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

// EventGrabBroken is a representation of the C record GdkEventGrabBroken.
//
// since 2.8
type EventGrabBroken struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventGrabBroken that represents the EventGrabBroken.
func (recv *EventGrabBroken) ToC() unsafe.Pointer {
	return recv.native
}

// EventGrabBrokenNewFromC creates a new EventGrabBroken from a pointer to the C GdkEventGrabBroken that represents the EventGrabBroken.
func EventGrabBrokenNewFromC(native unsafe.Pointer) *EventGrabBroken {
	return &EventGrabBroken{native: native}
}
