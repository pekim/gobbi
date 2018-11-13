// This is a generated file - DO NOT EDIT
// +build gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Returns the cursor type for this cursor.
/*

C function : gdk_cursor_get_cursor_type
*/
func (recv *Cursor) GetCursorType() CursorType {
	retC := C.gdk_cursor_get_cursor_type((*C.GdkCursor)(recv.native))
	retGo := (CursorType)(retC)

	return retGo
}

// Finds out if the display has been closed.
/*

C function : gdk_display_is_closed
*/
func (recv *Display) IsClosed() bool {
	retC := C.gdk_display_is_closed((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines the bitmask of actions proposed by the source if
// gdk_drag_context_get_suggested_action() returns %GDK_ACTION_ASK.
/*

C function : gdk_drag_context_get_actions
*/
func (recv *DragContext) GetActions() DragAction {
	retC := C.gdk_drag_context_get_actions((*C.GdkDragContext)(recv.native))
	retGo := (DragAction)(retC)

	return retGo
}

// Determines the action chosen by the drag destination.
/*

C function : gdk_drag_context_get_selected_action
*/
func (recv *DragContext) GetSelectedAction() DragAction {
	retC := C.gdk_drag_context_get_selected_action((*C.GdkDragContext)(recv.native))
	retGo := (DragAction)(retC)

	return retGo
}

// Returns the #GdkWindow where the DND operation started.
/*

C function : gdk_drag_context_get_source_window
*/
func (recv *DragContext) GetSourceWindow() *Window {
	retC := C.gdk_drag_context_get_source_window((*C.GdkDragContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Determines the suggested drag action of the context.
/*

C function : gdk_drag_context_get_suggested_action
*/
func (recv *DragContext) GetSuggestedAction() DragAction {
	retC := C.gdk_drag_context_get_suggested_action((*C.GdkDragContext)(recv.native))
	retGo := (DragAction)(retC)

	return retGo
}

// Retrieves the list of targets of the context.
/*

C function : gdk_drag_context_list_targets
*/
func (recv *DragContext) ListTargets() *glib.List {
	retC := C.gdk_drag_context_list_targets((*C.GdkDragContext)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the number of significant bits per red, green and blue value.
//
// Not all GDK backend provide a meaningful value for this function.
/*

C function : gdk_visual_get_bits_per_rgb
*/
func (recv *Visual) GetBitsPerRgb() int32 {
	retC := C.gdk_visual_get_bits_per_rgb((*C.GdkVisual)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Obtains values that are needed to calculate blue pixel values in TrueColor
// and DirectColor. The “mask” is the significant bits within the pixel.
// The “shift” is the number of bits left we must shift a primary for it
// to be in position (according to the "mask"). Finally, "precision" refers
// to how much precision the pixel value contains for a particular primary.
/*

C function : gdk_visual_get_blue_pixel_details
*/
func (recv *Visual) GetBluePixelDetails() (uint32, int32, int32) {
	var c_mask C.guint32

	var c_shift C.gint

	var c_precision C.gint

	C.gdk_visual_get_blue_pixel_details((*C.GdkVisual)(recv.native), &c_mask, &c_shift, &c_precision)

	mask := (uint32)(c_mask)

	shift := (int32)(c_shift)

	precision := (int32)(c_precision)

	return mask, shift, precision
}

// Returns the byte order of this visual.
//
// The information returned by this function is only relevant
// when working with XImages, and not all backends return
// meaningful information for this.
/*

C function : gdk_visual_get_byte_order
*/
func (recv *Visual) GetByteOrder() ByteOrder {
	retC := C.gdk_visual_get_byte_order((*C.GdkVisual)(recv.native))
	retGo := (ByteOrder)(retC)

	return retGo
}

// Returns the size of a colormap for this visual.
//
// You have to use platform-specific APIs to manipulate colormaps.
/*

C function : gdk_visual_get_colormap_size
*/
func (recv *Visual) GetColormapSize() int32 {
	retC := C.gdk_visual_get_colormap_size((*C.GdkVisual)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the bit depth of this visual.
/*

C function : gdk_visual_get_depth
*/
func (recv *Visual) GetDepth() int32 {
	retC := C.gdk_visual_get_depth((*C.GdkVisual)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Obtains values that are needed to calculate green pixel values in TrueColor
// and DirectColor. The “mask” is the significant bits within the pixel.
// The “shift” is the number of bits left we must shift a primary for it
// to be in position (according to the "mask"). Finally, "precision" refers
// to how much precision the pixel value contains for a particular primary.
/*

C function : gdk_visual_get_green_pixel_details
*/
func (recv *Visual) GetGreenPixelDetails() (uint32, int32, int32) {
	var c_mask C.guint32

	var c_shift C.gint

	var c_precision C.gint

	C.gdk_visual_get_green_pixel_details((*C.GdkVisual)(recv.native), &c_mask, &c_shift, &c_precision)

	mask := (uint32)(c_mask)

	shift := (int32)(c_shift)

	precision := (int32)(c_precision)

	return mask, shift, precision
}

// Obtains values that are needed to calculate red pixel values in TrueColor
// and DirectColor. The “mask” is the significant bits within the pixel.
// The “shift” is the number of bits left we must shift a primary for it
// to be in position (according to the "mask"). Finally, "precision" refers
// to how much precision the pixel value contains for a particular primary.
/*

C function : gdk_visual_get_red_pixel_details
*/
func (recv *Visual) GetRedPixelDetails() (uint32, int32, int32) {
	var c_mask C.guint32

	var c_shift C.gint

	var c_precision C.gint

	C.gdk_visual_get_red_pixel_details((*C.GdkVisual)(recv.native), &c_mask, &c_shift, &c_precision)

	mask := (uint32)(c_mask)

	shift := (int32)(c_shift)

	precision := (int32)(c_precision)

	return mask, shift, precision
}

// Returns the type of visual this is (PseudoColor, TrueColor, etc).
/*

C function : gdk_visual_get_visual_type
*/
func (recv *Visual) GetVisualType() VisualType {
	retC := C.gdk_visual_get_visual_type((*C.GdkVisual)(recv.native))
	retGo := (VisualType)(retC)

	return retGo
}

// Transforms window coordinates from a parent window to a child
// window, where the parent window is the normal parent as returned by
// gdk_window_get_parent() for normal windows, and the window's
// embedder as returned by gdk_offscreen_window_get_embedder() for
// offscreen windows.
//
// For normal windows, calling this function is equivalent to subtracting
// the return values of gdk_window_get_position() from the parent coordinates.
// For offscreen windows however (which can be arbitrarily transformed),
// this function calls the GdkWindow::from-embedder: signal to translate
// the coordinates.
//
// You should always use this function when writing generic code that
// walks down a window hierarchy.
//
// See also: gdk_window_coords_to_parent()
/*

C function : gdk_window_coords_from_parent
*/
func (recv *Window) CoordsFromParent(parentX float64, parentY float64) (float64, float64) {
	c_parent_x := (C.gdouble)(parentX)

	c_parent_y := (C.gdouble)(parentY)

	var c_x C.gdouble

	var c_y C.gdouble

	C.gdk_window_coords_from_parent((*C.GdkWindow)(recv.native), c_parent_x, c_parent_y, &c_x, &c_y)

	x := (float64)(c_x)

	y := (float64)(c_y)

	return x, y
}

// Transforms window coordinates from a child window to its parent
// window, where the parent window is the normal parent as returned by
// gdk_window_get_parent() for normal windows, and the window's
// embedder as returned by gdk_offscreen_window_get_embedder() for
// offscreen windows.
//
// For normal windows, calling this function is equivalent to adding
// the return values of gdk_window_get_position() to the child coordinates.
// For offscreen windows however (which can be arbitrarily transformed),
// this function calls the GdkWindow::to-embedder: signal to translate
// the coordinates.
//
// You should always use this function when writing generic code that
// walks up a window hierarchy.
//
// See also: gdk_window_coords_from_parent()
/*

C function : gdk_window_coords_to_parent
*/
func (recv *Window) CoordsToParent(x float64, y float64) (float64, float64) {
	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	var c_parent_x C.gdouble

	var c_parent_y C.gdouble

	C.gdk_window_coords_to_parent((*C.GdkWindow)(recv.native), c_x, c_y, &c_parent_x, &c_parent_y)

	parentX := (float64)(c_parent_x)

	parentY := (float64)(c_parent_y)

	return parentX, parentY
}

// Create a new surface that is as compatible as possible with the
// given @window. For example the new surface will have the same
// fallback resolution and font options as @window. Generally, the new
// surface will also use the same backend as @window, unless that is
// not possible for some reason. The type of the returned surface may
// be examined with cairo_surface_get_type().
//
// Initially the surface contents are all 0 (transparent if contents
// have transparency, black otherwise.)
/*

C function : gdk_window_create_similar_surface
*/
func (recv *Window) CreateSimilarSurface(content cairo.Content, width int32, height int32) *cairo.Surface {
	c_content := (C.cairo_content_t)(content)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	retC := C.gdk_window_create_similar_surface((*C.GdkWindow)(recv.native), c_content, c_width, c_height)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Determines whether or not the desktop environment shuld be hinted that
// the window does not want to receive input focus.
/*

C function : gdk_window_get_accept_focus
*/
func (recv *Window) GetAcceptFocus() bool {
	retC := C.gdk_window_get_accept_focus((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the pattern used to clear the background on @window. If @window
// does not have its own background and reuses the parent's, %NULL is
// returned and you’ll have to query it yourself.
/*

C function : gdk_window_get_background_pattern
*/
func (recv *Window) GetBackgroundPattern() *cairo.Pattern {
	retC := C.gdk_window_get_background_pattern((*C.GdkWindow)(recv.native))
	var retGo (*cairo.Pattern)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.PatternNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Determines whether @window is composited.
//
// See gdk_window_set_composited().
/*

C function : gdk_window_get_composited
*/
func (recv *Window) GetComposited() bool {
	retC := C.gdk_window_get_composited((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Obtains the parent of @window, as known to GDK. Works like
// gdk_window_get_parent() for normal windows, but returns the
// window’s embedder for offscreen windows.
//
// See also: gdk_offscreen_window_get_embedder()
/*

C function : gdk_window_get_effective_parent
*/
func (recv *Window) GetEffectiveParent() *Window {
	retC := C.gdk_window_get_effective_parent((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the toplevel window that’s an ancestor of @window.
//
// Works like gdk_window_get_toplevel(), but treats an offscreen window's
// embedder as its parent, using gdk_window_get_effective_parent().
//
// See also: gdk_offscreen_window_get_embedder()
/*

C function : gdk_window_get_effective_toplevel
*/
func (recv *Window) GetEffectiveToplevel() *Window {
	retC := C.gdk_window_get_effective_toplevel((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Determines whether or not the desktop environment should be hinted that the
// window does not want to receive input focus when it is mapped.
/*

C function : gdk_window_get_focus_on_map
*/
func (recv *Window) GetFocusOnMap() bool {
	retC := C.gdk_window_get_focus_on_map((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether or not the window manager is hinted that @window
// has modal behaviour.
/*

C function : gdk_window_get_modal_hint
*/
func (recv *Window) GetModalHint() bool {
	retC := C.gdk_window_get_modal_hint((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks whether the window has a native window or not. Note that
// you can use gdk_window_ensure_native() if a native window is needed.
/*

C function : gdk_window_has_native
*/
func (recv *Window) HasNative() bool {
	retC := C.gdk_window_has_native((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether or not the window is an input only window.
/*

C function : gdk_window_is_input_only
*/
func (recv *Window) IsInputOnly() bool {
	retC := C.gdk_window_is_input_only((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether or not the window is shaped.
/*

C function : gdk_window_is_shaped
*/
func (recv *Window) IsShaped() bool {
	retC := C.gdk_window_is_shaped((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
