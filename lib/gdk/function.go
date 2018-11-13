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

// Appends gdk option entries to the passed in option group. This is
// not public API and must not be used by applications.
/*

C function : gdk_add_option_entries_libgtk_only
*/
func AddOptionEntriesLibgtkOnly(group *glib.OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.gdk_add_option_entries_libgtk_only(c_group)

	return
}

// Unsupported : gdk_atom_intern : return type : Blacklisted record : GdkAtom

// Emits a short beep on the default display.
/*

C function : gdk_beep
*/
func Beep() {
	C.gdk_beep()

	return
}

// Unsupported : gdk_cairo_get_clip_rectangle : unsupported parameter rect : Blacklisted record : GdkRectangle

// Creates region that describes covers the area where the given
// @surface is more than 50% opaque.
//
// This function takes into account device offsets that might be
// set with cairo_surface_set_device_offset().
/*

C function : gdk_cairo_region_create_from_surface
*/
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	retC := C.gdk_cairo_region_create_from_surface(c_surface)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : gdk_color_parse

// Aborts a drag without dropping.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop
// operations. See gdk_drag_context_manage_dnd() for more information.
/*

C function : gdk_drag_abort
*/
func DragAbort(context *DragContext, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_time_ := (C.guint32)(time)

	C.gdk_drag_abort(c_context, c_time_)

	return
}

// Starts a drag and creates a new drag context for it.
// This function assumes that the drag is controlled by the
// client pointer device, use gdk_drag_begin_for_device() to
// begin a drag with a different device.
//
// This function is called by the drag source.
/*

C function : gdk_drag_begin
*/
func DragBegin(window *Window, targets *glib.List) *DragContext {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	retC := C.gdk_drag_begin(c_window, c_targets)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Starts a drag and creates a new drag context for it.
//
// This function is called by the drag source.
/*

C function : gdk_drag_begin_for_device
*/
func DragBeginForDevice(window *Window, device *Device, targets *glib.List) *DragContext {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	retC := C.gdk_drag_begin_for_device(c_window, c_device, c_targets)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Drops on the current destination.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop
// operations. See gdk_drag_context_manage_dnd() for more information.
/*

C function : gdk_drag_drop
*/
func DragDrop(context *DragContext, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_time_ := (C.guint32)(time)

	C.gdk_drag_drop(c_context, c_time_)

	return
}

// Unsupported : gdk_drag_get_selection : return type : Blacklisted record : GdkAtom

// Updates the drag context when the pointer moves or the
// set of actions changes.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop
// operations. See gdk_drag_context_manage_dnd() for more information.
/*

C function : gdk_drag_motion
*/
func DragMotion(context *DragContext, destWindow *Window, protocol DragProtocol, xRoot int32, yRoot int32, suggestedAction DragAction, possibleActions DragAction, time uint32) bool {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_dest_window := (*C.GdkWindow)(C.NULL)
	if destWindow != nil {
		c_dest_window = (*C.GdkWindow)(destWindow.ToC())
	}

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

// Selects one of the actions offered by the drag source.
//
// This function is called by the drag destination in response to
// gdk_drag_motion() called by the drag source.
/*

C function : gdk_drag_status
*/
func DragStatus(context *DragContext, action DragAction, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_action := (C.GdkDragAction)(action)

	c_time_ := (C.guint32)(time)

	C.gdk_drag_status(c_context, c_action, c_time_)

	return
}

// Ends the drag operation after a drop.
//
// This function is called by the drag destination.
/*

C function : gdk_drop_finish
*/
func DropFinish(context *DragContext, success bool, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_success :=
		boolToGboolean(success)

	c_time_ := (C.guint32)(time)

	C.gdk_drop_finish(c_context, c_success, c_time_)

	return
}

// Accepts or rejects a drop.
//
// This function is called by the drag destination in response
// to a drop initiated by the drag source.
/*

C function : gdk_drop_reply
*/
func DropReply(context *DragContext, accepted bool, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_accepted :=
		boolToGboolean(accepted)

	c_time_ := (C.guint32)(time)

	C.gdk_drop_reply(c_context, c_accepted, c_time_)

	return
}

// Removes an error trap pushed with gdk_error_trap_push().
// May block until an error has been definitively received
// or not received from the X server. gdk_error_trap_pop_ignored()
// is preferred if you don’t need to know whether an error
// occurred, because it never has to block. If you don't
// need the return value of gdk_error_trap_pop(), use
// gdk_error_trap_pop_ignored().
//
// Prior to GDK 3.0, this function would not automatically
// sync for you, so you had to gdk_flush() if your last
// call to Xlib was not a blocking round trip.
/*

C function : gdk_error_trap_pop
*/
func ErrorTrapPop() int32 {
	retC := C.gdk_error_trap_pop()
	retGo := (int32)(retC)

	return retGo
}

// This function allows X errors to be trapped instead of the normal
// behavior of exiting the application. It should only be used if it
// is not possible to avoid the X error in any other way. Errors are
// ignored on all #GdkDisplay currently known to the
// #GdkDisplayManager. If you don’t care which error happens and just
// want to ignore everything, pop with gdk_error_trap_pop_ignored().
// If you need the error code, use gdk_error_trap_pop() which may have
// to block and wait for the error to arrive from the X server.
//
// This API exists on all platforms but only does anything on X.
//
// You can use gdk_x11_display_error_trap_push() to ignore errors
// on only a single display.
//
// ## Trapping an X error
//
// |[<!-- language="C" -->
// gdk_error_trap_push ();
//
// ... Call the X function which may cause an error here ...
//
//
// if (gdk_error_trap_pop ())
// {
// ... Handle the error here ...
// }
// ]|
/*

C function : gdk_error_trap_push
*/
func ErrorTrapPush() {
	C.gdk_error_trap_push()

	return
}

// Unsupported : gdk_event_get : no return generator

// Unsupported : gdk_event_handler_set : unsupported parameter func : no type generator for EventFunc (GdkEventFunc) for param func

// Unsupported : gdk_event_peek : no return generator

// Checks if any events are ready to be processed for any display.
/*

C function : gdk_events_pending
*/
func EventsPending() bool {
	retC := C.gdk_events_pending()
	retGo := retC == C.TRUE

	return retGo
}

// Flushes the output buffers of all display connections and waits
// until all requests have been processed.
// This is rarely needed by applications.
/*

C function : gdk_flush
*/
func Flush() {
	C.gdk_flush()

	return
}

// Obtains the root window (parent all other windows are inside)
// for the default display and screen.
/*

C function : gdk_get_default_root_window
*/
func GetDefaultRootWindow() *Window {
	retC := C.gdk_get_default_root_window()
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the name of the display, which usually comes from the
// `DISPLAY` environment variable or the
// `--display` command line option.
/*

C function : gdk_get_display
*/
func GetDisplay() string {
	retC := C.gdk_get_display()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the program class. Unless the program class has explicitly
// been set with gdk_set_program_class() or with the `--class`
// commandline option, the default value is the program name (determined
// with g_get_prgname()) with the first character converted to uppercase.
/*

C function : gdk_get_program_class
*/
func GetProgramClass() string {
	retC := C.gdk_get_program_class()
	retGo := C.GoString(retC)

	return retGo
}

// Gets whether event debugging output is enabled.
/*

C function : gdk_get_show_events
*/
func GetShowEvents() bool {
	retC := C.gdk_get_show_events()
	retGo := retC == C.TRUE

	return retGo
}

/*

C function : gdk_gl_error_quark
*/
func GlErrorQuark() glib.Quark {
	retC := C.gdk_gl_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Initializes the GDK library and connects to the windowing system.
// If initialization fails, a warning message is output and the application
// terminates with a call to `exit(1)`.
//
// Any arguments used by GDK are removed from the array and @argc and @argv
// are updated accordingly.
//
// GTK+ initializes GDK in gtk_init() and so this function is not usually
// needed by GTK+ applications.
/*

C function : gdk_init
*/
func Init(args []string) []string {
	cArgc, cArgv := argsIn(args)

	C.gdk_init(&cArgc, &cArgv)

	args = argsOut(cArgc, cArgv)

	return args
}

// Initializes the GDK library and connects to the windowing system,
// returning %TRUE on success.
//
// Any arguments used by GDK are removed from the array and @argc and @argv
// are updated accordingly.
//
// GTK+ initializes GDK in gtk_init() and so this function is not usually
// needed by GTK+ applications.
/*

C function : gdk_init_check
*/
func InitCheck(args []string) (bool, []string) {
	cArgc, cArgv := argsIn(args)

	retC := C.gdk_init_check(&cArgc, &cArgv)
	retGo := retC == C.TRUE

	args = argsOut(cArgc, cArgv)

	return retGo, args
}

// Grabs the keyboard so that all events are passed to this
// application until the keyboard is ungrabbed with gdk_keyboard_ungrab().
// This overrides any previous keyboard grab by this client.
//
// If you set up anything at the time you take the grab that needs to be cleaned
// up when the grab ends, you should handle the #GdkEventGrabBroken events that
// are emitted when the grab ends unvoluntarily.
/*

C function : gdk_keyboard_grab
*/
func KeyboardGrab(window *Window, ownerEvents bool, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_time_ := (C.guint32)(time)

	retC := C.gdk_keyboard_grab(c_window, c_owner_events, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// Ungrabs the keyboard on the default display, if it is grabbed by this
// application.
/*

C function : gdk_keyboard_ungrab
*/
func KeyboardUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_keyboard_ungrab(c_time_)

	return
}

// Obtains the upper- and lower-case versions of the keyval @symbol.
// Examples of keyvals are #GDK_KEY_a, #GDK_KEY_Enter, #GDK_KEY_F1, etc.
/*

C function : gdk_keyval_convert_case
*/
func KeyvalConvertCase(symbol uint32) (uint32, uint32) {
	c_symbol := (C.guint)(symbol)

	var c_lower C.guint

	var c_upper C.guint

	C.gdk_keyval_convert_case(c_symbol, &c_lower, &c_upper)

	lower := (uint32)(c_lower)

	upper := (uint32)(c_upper)

	return lower, upper
}

// Converts a key name to a key value.
//
// The names are the same as those in the
// `gdk/gdkkeysyms.h` header file
// but without the leading “GDK_KEY_”.
/*

C function : gdk_keyval_from_name
*/
func KeyvalFromName(keyvalName string) uint32 {
	c_keyval_name := C.CString(keyvalName)
	defer C.free(unsafe.Pointer(c_keyval_name))

	retC := C.gdk_keyval_from_name(c_keyval_name)
	retGo := (uint32)(retC)

	return retGo
}

// Returns %TRUE if the given key value is in lower case.
/*

C function : gdk_keyval_is_lower
*/
func KeyvalIsLower(keyval uint32) bool {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_is_lower(c_keyval)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if the given key value is in upper case.
/*

C function : gdk_keyval_is_upper
*/
func KeyvalIsUpper(keyval uint32) bool {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_is_upper(c_keyval)
	retGo := retC == C.TRUE

	return retGo
}

// Converts a key value into a symbolic name.
//
// The names are the same as those in the
// `gdk/gdkkeysyms.h` header file
// but without the leading “GDK_KEY_”.
/*

C function : gdk_keyval_name
*/
func KeyvalName(keyval uint32) string {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_name(c_keyval)
	retGo := C.GoString(retC)

	return retGo
}

// Converts a key value to lower case, if applicable.
/*

C function : gdk_keyval_to_lower
*/
func KeyvalToLower(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_lower(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// Convert from a GDK key symbol to the corresponding ISO10646 (Unicode)
// character.
/*

C function : gdk_keyval_to_unicode
*/
func KeyvalToUnicode(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_unicode(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// Converts a key value to upper case, if applicable.
/*

C function : gdk_keyval_to_upper
*/
func KeyvalToUpper(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_upper(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// Lists the available visuals for the default screen.
// (See gdk_screen_list_visuals())
// A visual describes a hardware image data format.
// For example, a visual might support 24-bit color, or 8-bit color,
// and might expect pixels to be in a certain format.
//
// Call g_list_free() on the return value when you’re finished with it.
/*

C function : gdk_list_visuals
*/
func ListVisuals() *glib.List {
	retC := C.gdk_list_visuals()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the offscreen surface that an offscreen window renders into.
// If you need to keep this around over window resizes, you need to
// add a reference to it.
/*

C function : gdk_offscreen_window_get_surface
*/
func OffscreenWindowGetSurface(window *Window) *cairo.Surface {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gdk_offscreen_window_get_surface(c_window)
	var retGo (*cairo.Surface)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.SurfaceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Creates a #PangoContext for the default GDK screen.
//
// The context must be freed when you’re finished with it.
//
// When using GTK+, normally you should use gtk_widget_get_pango_context()
// instead of this function, to get the appropriate context for
// the widget you intend to render text onto.
//
// The newly created context will have the default font options (see
// #cairo_font_options_t) for the default screen; if these options
// change it will not be updated. Using gtk_widget_get_pango_context()
// is more convenient if you want to keep a context around and track
// changes to the screen’s font rendering settings.
/*

C function : gdk_pango_context_get
*/
func PangoContextGet() *pango.Context {
	retC := C.gdk_pango_context_get()
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Obtains a clip region which contains the areas where the given ranges
// of text would be drawn. @x_origin and @y_origin are the top left point
// to center the layout. @index_ranges should contain
// ranges of bytes in the layout’s text.
//
// Note that the regions returned correspond to logical extents of the text
// ranges, not ink extents. So the drawn layout may in fact touch areas out of
// the clip region.  The clip region is mainly useful for highlightling parts
// of text, such as when text is selected.
/*

C function : gdk_pango_layout_get_clip_region
*/
func PangoLayoutGetClipRegion(layout *pango.Layout, xOrigin int32, yOrigin int32, indexRanges int32, nRanges int32) *cairo.Region {
	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	c_x_origin := (C.gint)(xOrigin)

	c_y_origin := (C.gint)(yOrigin)

	c_index_ranges := (C.gint)(indexRanges)

	c_n_ranges := (C.gint)(nRanges)

	retC := C.gdk_pango_layout_get_clip_region(c_layout, c_x_origin, c_y_origin, &c_index_ranges, c_n_ranges)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Obtains a clip region which contains the areas where the given
// ranges of text would be drawn. @x_origin and @y_origin are the top left
// position of the layout. @index_ranges
// should contain ranges of bytes in the layout’s text. The clip
// region will include space to the left or right of the line (to the
// layout bounding box) if you have indexes above or below the indexes
// contained inside the line. This is to draw the selection all the way
// to the side of the layout. However, the clip region is in line coordinates,
// not layout coordinates.
//
// Note that the regions returned correspond to logical extents of the text
// ranges, not ink extents. So the drawn line may in fact touch areas out of
// the clip region.  The clip region is mainly useful for highlightling parts
// of text, such as when text is selected.
/*

C function : gdk_pango_layout_line_get_clip_region
*/
func PangoLayoutLineGetClipRegion(line *pango.LayoutLine, xOrigin int32, yOrigin int32, indexRanges []int32, nRanges int32) *cairo.Region {
	c_line := (*C.PangoLayoutLine)(C.NULL)
	if line != nil {
		c_line = (*C.PangoLayoutLine)(line.ToC())
	}

	c_x_origin := (C.gint)(xOrigin)

	c_y_origin := (C.gint)(yOrigin)

	c_index_ranges := &indexRanges[0]

	c_n_ranges := (C.gint)(nRanges)

	retC := C.gdk_pango_layout_line_get_clip_region(c_line, c_x_origin, c_y_origin, (*C.gint)(unsafe.Pointer(c_index_ranges)), c_n_ranges)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Transfers image data from a #cairo_surface_t and converts it to an RGB(A)
// representation inside a #GdkPixbuf. This allows you to efficiently read
// individual pixels from cairo surfaces. For #GdkWindows, use
// gdk_pixbuf_get_from_window() instead.
//
// This function will create an RGB pixbuf with 8 bits per channel.
// The pixbuf will contain an alpha channel if the @surface contains one.
/*

C function : gdk_pixbuf_get_from_surface
*/
func PixbufGetFromSurface(surface *cairo.Surface, srcX int32, srcY int32, width int32, height int32) *gdkpixbuf.Pixbuf {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	c_src_x := (C.gint)(srcX)

	c_src_y := (C.gint)(srcY)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.gdk_pixbuf_get_from_surface(c_surface, c_src_x, c_src_y, c_width, c_height)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Transfers image data from a #GdkWindow and converts it to an RGB(A)
// representation inside a #GdkPixbuf. In other words, copies
// image data from a server-side drawable to a client-side RGB(A) buffer.
// This allows you to efficiently read individual pixels on the client side.
//
// This function will create an RGB pixbuf with 8 bits per channel with
// the size specified by the @width and @height arguments scaled by the
// scale factor of @window. The pixbuf will contain an alpha channel if
// the @window contains one.
//
// If the window is off the screen, then there is no image data in the
// obscured/offscreen regions to be placed in the pixbuf. The contents of
// portions of the pixbuf corresponding to the offscreen region are undefined.
//
// If the window you’re obtaining data from is partially obscured by
// other windows, then the contents of the pixbuf areas corresponding
// to the obscured regions are undefined.
//
// If the window is not mapped (typically because it’s iconified/minimized
// or not on the current workspace), then %NULL will be returned.
//
// If memory can’t be allocated for the return value, %NULL will be returned
// instead.
//
// (In short, there are several ways this function can fail, and if it fails
// it returns %NULL; so check the return value.)
/*

C function : gdk_pixbuf_get_from_window
*/
func PixbufGetFromWindow(window *Window, srcX int32, srcY int32, width int32, height int32) *gdkpixbuf.Pixbuf {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_src_x := (C.gint)(srcX)

	c_src_y := (C.gint)(srcY)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.gdk_pixbuf_get_from_window(c_window, c_src_x, c_src_y, c_width, c_height)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Grabs the pointer (usually a mouse) so that all events are passed to this
// application until the pointer is ungrabbed with gdk_pointer_ungrab(), or
// the grab window becomes unviewable.
// This overrides any previous pointer grab by this client.
//
// Pointer grabs are used for operations which need complete control over mouse
// events, even if the mouse leaves the application.
// For example in GTK+ it is used for Drag and Drop, for dragging the handle in
// the #GtkHPaned and #GtkVPaned widgets.
//
// Note that if the event mask of an X window has selected both button press and
// button release events, then a button press event will cause an automatic
// pointer grab until the button is released.
// X does this automatically since most applications expect to receive button
// press and release events in pairs.
// It is equivalent to a pointer grab on the window with @owner_events set to
// %TRUE.
//
// If you set up anything at the time you take the grab that needs to be cleaned
// up when the grab ends, you should handle the #GdkEventGrabBroken events that
// are emitted when the grab ends unvoluntarily.
/*

C function : gdk_pointer_grab
*/
func PointerGrab(window *Window, ownerEvents bool, eventMask EventMask, confineTo *Window, cursor *Cursor, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_event_mask := (C.GdkEventMask)(eventMask)

	c_confine_to := (*C.GdkWindow)(C.NULL)
	if confineTo != nil {
		c_confine_to = (*C.GdkWindow)(confineTo.ToC())
	}

	c_cursor := (*C.GdkCursor)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkCursor)(cursor.ToC())
	}

	c_time_ := (C.guint32)(time)

	retC := C.gdk_pointer_grab(c_window, c_owner_events, c_event_mask, c_confine_to, c_cursor, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// Returns %TRUE if the pointer on the default display is currently
// grabbed by this application.
//
// Note that this does not take the inmplicit pointer grab on button
// presses into account.
/*

C function : gdk_pointer_is_grabbed
*/
func PointerIsGrabbed() bool {
	retC := C.gdk_pointer_is_grabbed()
	retGo := retC == C.TRUE

	return retGo
}

// Ungrabs the pointer on the default display, if it is grabbed by this
// application.
/*

C function : gdk_pointer_ungrab
*/
func PointerUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_pointer_ungrab(c_time_)

	return
}

// Prepare for parsing command line arguments for GDK. This is not
// public API and should not be used in application code.
/*

C function : gdk_pre_parse_libgtk_only
*/
func PreParseLibgtkOnly() {
	C.gdk_pre_parse_libgtk_only()

	return
}

// Unsupported : gdk_property_change : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_property_delete : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_property_get : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_query_depths : unsupported parameter depths : output array param depths

// Unsupported : gdk_query_visual_types : unsupported parameter visual_types : output array param visual_types

// Unsupported : gdk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_get : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_property_get : unsupported parameter data : guchar** with indirection level of 2

// Unsupported : gdk_selection_send_notify : unsupported parameter selection : Blacklisted record : GdkAtom

// Set the double click time for the default display. See
// gdk_display_set_double_click_time().
// See also gdk_display_set_double_click_distance().
// Applications should not set this, it is a
// global user-configured setting.
/*

C function : gdk_set_double_click_time
*/
func SetDoubleClickTime(msec uint32) {
	c_msec := (C.guint)(msec)

	C.gdk_set_double_click_time(c_msec)

	return
}

// Sets the program class. The X11 backend uses the program class to set
// the class name part of the `WM_CLASS` property on
// toplevel windows; see the ICCCM.
//
// The program class can still be overridden with the --class command
// line option.
/*

C function : gdk_set_program_class
*/
func SetProgramClass(programClass string) {
	c_program_class := C.CString(programClass)
	defer C.free(unsafe.Pointer(c_program_class))

	C.gdk_set_program_class(c_program_class)

	return
}

// Sets whether a trace of received events is output.
// Note that GTK+ must be compiled with debugging (that is,
// configured using the `--enable-debug` option)
// to use this option.
/*

C function : gdk_set_show_events
*/
func SetShowEvents(showEvents bool) {
	c_show_events :=
		boolToGboolean(showEvents)

	C.gdk_set_show_events(c_show_events)

	return
}

// Obtains a desktop-wide setting, such as the double-click time,
// for the default screen. See gdk_screen_get_setting().
/*

C function : gdk_setting_get
*/
func SettingGet(name string, value *gobject.Value) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.gdk_setting_get(c_name, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : gdk_synthesize_window_state

// This function marks the beginning of a critical section in which
// GDK and GTK+ functions can be called safely and without causing race
// conditions. Only one thread at a time can be in such a critial
// section.
/*

C function : gdk_threads_enter
*/
func ThreadsEnter() {
	C.gdk_threads_enter()

	return
}

// Initializes GDK so that it can be used from multiple threads
// in conjunction with gdk_threads_enter() and gdk_threads_leave().
//
// This call must be made before any use of the main loop from
// GTK+; to be safe, call it before gtk_init().
/*

C function : gdk_threads_init
*/
func ThreadsInit() {
	C.gdk_threads_init()

	return
}

// Leaves a critical region begun with gdk_threads_enter().
/*

C function : gdk_threads_leave
*/
func ThreadsLeave() {
	C.gdk_threads_leave()

	return
}

// Convert from a ISO10646 character to a key symbol.
/*

C function : gdk_unicode_to_keyval
*/
func UnicodeToKeyval(wc uint32) uint32 {
	c_wc := (C.guint32)(wc)

	retC := C.gdk_unicode_to_keyval(c_wc)
	retGo := (uint32)(retC)

	return retGo
}

// Converts an UTF-8 string into the best possible representation
// as a STRING. The representation of characters not in STRING
// is not specified; it may be as pseudo-escape sequences
// \x{ABCD}, or it may be in some other form of approximation.
/*

C function : gdk_utf8_to_string_target
*/
func Utf8ToStringTarget(str string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.gdk_utf8_to_string_target(c_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
