// This is a generated file - DO NOT EDIT
// +build gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_atom_intern : return type : Blacklisted record : GdkAtom

// Unsupported : gdk_atom_intern_static_string : return type : Blacklisted record : GdkAtom

// Unsupported : gdk_cairo_get_clip_rectangle : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gdk_cairo_rectangle : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// CairoSurfaceCreateFromPixbuf is a wrapper around the C function gdk_cairo_surface_create_from_pixbuf.
func CairoSurfaceCreateFromPixbuf(pixbuf *gdkpixbuf.Pixbuf, scale int32, forWindow *Window) *cairo.Surface {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	c_scale := (C.int)(scale)

	c_for_window := (*C.GdkWindow)(forWindow.ToC())

	retC := C.gdk_cairo_surface_create_from_pixbuf(c_pixbuf, c_scale, c_for_window)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_drag_find_window_for_screen : unsupported parameter dest_window : record with indirection level of 2

// Unsupported : gdk_drag_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gdk_event_get : no return generator

// Unsupported : gdk_event_handler_set : unsupported parameter func : no type generator for EventFunc (GdkEventFunc) for param func

// Unsupported : gdk_event_peek : no return generator

// Unsupported : gdk_events_get_angle : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1

// Unsupported : gdk_events_get_center : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1

// Unsupported : gdk_events_get_distance : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1

// Unsupported : gdk_init : unsupported parameter argv : no type generator for utf8 (gchar**) for array param argv

// Unsupported : gdk_init_check : unsupported parameter argv : no type generator for utf8 (gchar**) for array param argv

// Unsupported : gdk_pango_layout_line_get_clip_region : unsupported parameter index_ranges : no type generator for gint (gint) for array param index_ranges

// Unsupported : gdk_parse_args : unsupported parameter argv : no type generator for utf8 (gchar**) for array param argv

// Unsupported : gdk_property_change : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_property_delete : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_property_get : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_query_depths : unsupported parameter depths : no type generator for gint (gint*) for array param depths

// Unsupported : gdk_query_visual_types : unsupported parameter visual_types : no type generator for VisualType (GdkVisualType*) for array param visual_types

// Unsupported : gdk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_get : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_get_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_set_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_property_get : unsupported parameter data : guchar** with indirection level of 2

// Unsupported : gdk_selection_send_notify : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_send_notify_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// SetAllowedBackends is a wrapper around the C function gdk_set_allowed_backends.
func SetAllowedBackends(backends string) {
	c_backends := C.CString(backends)
	defer C.free(unsafe.Pointer(c_backends))

	C.gdk_set_allowed_backends(c_backends)

	return
}

// Unsupported : gdk_text_property_to_utf8_list_for_display : unsupported parameter encoding : Blacklisted record : GdkAtom

// Unsupported : gdk_threads_add_idle : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_add_idle_full : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_add_timeout : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_add_timeout_full : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_add_timeout_seconds : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_add_timeout_seconds_full : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_set_lock_functions : unsupported parameter enter_fn : no type generator for GObject.Callback (GCallback) for param enter_fn
