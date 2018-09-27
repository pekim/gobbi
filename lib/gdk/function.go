// This is a generated file - DO NOT EDIT

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_add_option_entries_libgtk_only : no return generator

// Unsupported : gdk_atom_intern : return type : Blacklisted record : GdkAtom

// Unsupported : gdk_beep : no return generator

// Unsupported : gdk_cairo_get_clip_rectangle : unsupported parameter rect : Blacklisted record : GdkRectangle

// CairoRegionCreateFromSurface is a wrapper around the C function gdk_cairo_region_create_from_surface.
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.gdk_cairo_region_create_from_surface(c_surface)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : gdk_color_parse

// Unsupported : gdk_drag_abort : no return generator

// DragBegin is a wrapper around the C function gdk_drag_begin.
func DragBegin(window *Window, targets *glib.List) *DragContext {
	c_window := (*C.GdkWindow)(window.ToC())

	c_targets := (*C.GList)(targets.ToC())

	retC := C.gdk_drag_begin(c_window, c_targets)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DragBeginForDevice is a wrapper around the C function gdk_drag_begin_for_device.
func DragBeginForDevice(window *Window, device *Device, targets *glib.List) *DragContext {
	c_window := (*C.GdkWindow)(window.ToC())

	c_device := (*C.GdkDevice)(device.ToC())

	c_targets := (*C.GList)(targets.ToC())

	retC := C.gdk_drag_begin_for_device(c_window, c_device, c_targets)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_drag_drop : no return generator

// Unsupported : gdk_drag_get_selection : return type : Blacklisted record : GdkAtom

// DragMotion is a wrapper around the C function gdk_drag_motion.
func DragMotion(context *DragContext, destWindow *Window, protocol DragProtocol, xRoot int32, yRoot int32, suggestedAction DragAction, possibleActions DragAction, time uint32) bool {
	c_context := (*C.GdkDragContext)(context.ToC())

	c_dest_window := (*C.GdkWindow)(destWindow.ToC())

	c_protocol := (C.GdkDragProtocol)(protocol)

	c_x_root := (C.gint)(xRoot)

	c_y_root := (C.gint)(yRoot)

	c_suggested_action := (C.GdkDragAction)(suggestedAction)

	c_possible_actions := (C.GdkDragAction)(possibleActions)

	c_time_ := (C.guint32)(time)

	retC := C.gdk_drag_motion(c_context, c_dest_window, c_protocol, c_x_root, c_y_root, c_suggested_action, c_possible_actions, c_time_)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_drag_status : no return generator

// Unsupported : gdk_drop_finish : no return generator

// Unsupported : gdk_drop_reply : no return generator

// ErrorTrapPop is a wrapper around the C function gdk_error_trap_pop.
func ErrorTrapPop() int32 {
	retC := C.gdk_error_trap_pop()
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_error_trap_push : no return generator

// Unsupported : gdk_event_get : no return generator

// Unsupported : gdk_event_handler_set : unsupported parameter func : no type generator for EventFunc, GdkEventFunc

// Unsupported : gdk_event_peek : no return generator

// EventsPending is a wrapper around the C function gdk_events_pending.
func EventsPending() bool {
	retC := C.gdk_events_pending()
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_flush : no return generator

// GetDefaultRootWindow is a wrapper around the C function gdk_get_default_root_window.
func GetDefaultRootWindow() *Window {
	retC := C.gdk_get_default_root_window()
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_get_display.
func GetDisplay() string {
	retC := C.gdk_get_display()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetProgramClass is a wrapper around the C function gdk_get_program_class.
func GetProgramClass() string {
	retC := C.gdk_get_program_class()
	retGo := C.GoString(retC)

	return retGo
}

// GetShowEvents is a wrapper around the C function gdk_get_show_events.
func GetShowEvents() bool {
	retC := C.gdk_get_show_events()
	retGo := retC == C.TRUE

	return retGo
}

// GlErrorQuark is a wrapper around the C function gdk_gl_error_quark.
func GlErrorQuark() glib.Quark {
	retC := C.gdk_gl_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : gdk_init : unsupported parameter argc : no type generator for gint, gint*

// Unsupported : gdk_init_check : unsupported parameter argc : no type generator for gint, gint*

// KeyboardGrab is a wrapper around the C function gdk_keyboard_grab.
func KeyboardGrab(window *Window, ownerEvents bool, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(window.ToC())

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_time_ := (C.guint32)(time)

	retC := C.gdk_keyboard_grab(c_window, c_owner_events, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// Unsupported : gdk_keyboard_ungrab : no return generator

// Unsupported : gdk_keyval_convert_case : unsupported parameter lower : no type generator for guint, guint*

// KeyvalFromName is a wrapper around the C function gdk_keyval_from_name.
func KeyvalFromName(keyvalName string) uint32 {
	c_keyval_name := C.CString(keyvalName)
	defer C.free(unsafe.Pointer(c_keyval_name))

	retC := C.gdk_keyval_from_name(c_keyval_name)
	retGo := (uint32)(retC)

	return retGo
}

// KeyvalIsLower is a wrapper around the C function gdk_keyval_is_lower.
func KeyvalIsLower(keyval uint32) bool {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_is_lower(c_keyval)
	retGo := retC == C.TRUE

	return retGo
}

// KeyvalIsUpper is a wrapper around the C function gdk_keyval_is_upper.
func KeyvalIsUpper(keyval uint32) bool {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_is_upper(c_keyval)
	retGo := retC == C.TRUE

	return retGo
}

// KeyvalName is a wrapper around the C function gdk_keyval_name.
func KeyvalName(keyval uint32) string {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_name(c_keyval)
	retGo := C.GoString(retC)

	return retGo
}

// KeyvalToLower is a wrapper around the C function gdk_keyval_to_lower.
func KeyvalToLower(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_lower(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// KeyvalToUnicode is a wrapper around the C function gdk_keyval_to_unicode.
func KeyvalToUnicode(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_unicode(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// KeyvalToUpper is a wrapper around the C function gdk_keyval_to_upper.
func KeyvalToUpper(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_upper(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// ListVisuals is a wrapper around the C function gdk_list_visuals.
func ListVisuals() *glib.List {
	retC := C.gdk_list_visuals()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// OffscreenWindowGetSurface is a wrapper around the C function gdk_offscreen_window_get_surface.
func OffscreenWindowGetSurface(window *Window) *cairo.Surface {
	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gdk_offscreen_window_get_surface(c_window)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PangoContextGet is a wrapper around the C function gdk_pango_context_get.
func PangoContextGet() *pango.Context {
	retC := C.gdk_pango_context_get()
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_pango_layout_get_clip_region : unsupported parameter index_ranges : no type generator for gint, const gint*

// Unsupported : gdk_pango_layout_line_get_clip_region : unsupported parameter index_ranges : no param type

// PixbufGetFromSurface is a wrapper around the C function gdk_pixbuf_get_from_surface.
func PixbufGetFromSurface(surface *cairo.Surface, srcX int32, srcY int32, width int32, height int32) *gdkpixbuf.Pixbuf {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	c_src_x := (C.gint)(srcX)

	c_src_y := (C.gint)(srcY)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.gdk_pixbuf_get_from_surface(c_surface, c_src_x, c_src_y, c_width, c_height)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PixbufGetFromWindow is a wrapper around the C function gdk_pixbuf_get_from_window.
func PixbufGetFromWindow(window *Window, srcX int32, srcY int32, width int32, height int32) *gdkpixbuf.Pixbuf {
	c_window := (*C.GdkWindow)(window.ToC())

	c_src_x := (C.gint)(srcX)

	c_src_y := (C.gint)(srcY)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.gdk_pixbuf_get_from_window(c_window, c_src_x, c_src_y, c_width, c_height)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PointerGrab is a wrapper around the C function gdk_pointer_grab.
func PointerGrab(window *Window, ownerEvents bool, eventMask EventMask, confineTo *Window, cursor *Cursor, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(window.ToC())

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_event_mask := (C.GdkEventMask)(eventMask)

	c_confine_to := (*C.GdkWindow)(confineTo.ToC())

	c_cursor := (*C.GdkCursor)(cursor.ToC())

	c_time_ := (C.guint32)(time)

	retC := C.gdk_pointer_grab(c_window, c_owner_events, c_event_mask, c_confine_to, c_cursor, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// PointerIsGrabbed is a wrapper around the C function gdk_pointer_is_grabbed.
func PointerIsGrabbed() bool {
	retC := C.gdk_pointer_is_grabbed()
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_pointer_ungrab : no return generator

// Unsupported : gdk_pre_parse_libgtk_only : no return generator

// Unsupported : gdk_property_change : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_property_delete : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_property_get : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_query_depths : unsupported parameter depths : no param type

// Unsupported : gdk_query_visual_types : unsupported parameter visual_types : no param type

// Unsupported : gdk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_get : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_property_get : unsupported parameter data : no type generator for guint8, guchar**

// Unsupported : gdk_selection_send_notify : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_set_double_click_time : no return generator

// Unsupported : gdk_set_program_class : no return generator

// Unsupported : gdk_set_show_events : no return generator

// SettingGet is a wrapper around the C function gdk_setting_get.
func SettingGet(name string, value *gobject.Value) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.GValue)(value.ToC())

	retC := C.gdk_setting_get(c_name, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_synthesize_window_state : no return generator

// Unsupported : gdk_threads_enter : no return generator

// Unsupported : gdk_threads_init : no return generator

// Unsupported : gdk_threads_leave : no return generator

// UnicodeToKeyval is a wrapper around the C function gdk_unicode_to_keyval.
func UnicodeToKeyval(wc uint32) uint32 {
	c_wc := (C.guint32)(wc)

	retC := C.gdk_unicode_to_keyval(c_wc)
	retGo := (uint32)(retC)

	return retGo
}

// Utf8ToStringTarget is a wrapper around the C function gdk_utf8_to_string_target.
func Utf8ToStringTarget(str string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.gdk_utf8_to_string_target(c_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
