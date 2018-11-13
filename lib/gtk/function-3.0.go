// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Parses a signal description from @signal_desc and incorporates
// it into @binding_set.
//
// Signal descriptions may either bind a key combination to
// one or more signals:
// |[
// bind "key" {
// "signalname" (param, ...)
// ...
// }
// ]|
//
// Or they may also unbind a key combination:
// |[
// unbind "key"
// ]|
//
// Key combinations must be in a format that can be parsed by
// gtk_accelerator_parse().
/*

C function

gtk_binding_entry_add_signal_from_string
*/
func BindingEntryAddSignalFromString(bindingSet *BindingSet, signalDesc string) glib.TokenType {
	c_binding_set := (*C.GtkBindingSet)(C.NULL)
	if bindingSet != nil {
		c_binding_set = (*C.GtkBindingSet)(bindingSet.ToC())
	}

	c_signal_desc := C.CString(signalDesc)
	defer C.free(unsafe.Pointer(c_signal_desc))

	retC := C.gtk_binding_entry_add_signal_from_string(c_binding_set, c_signal_desc)
	retGo := (glib.TokenType)(retC)

	return retGo
}

// This function is supposed to be called in #GtkWidget::draw
// implementations for widgets that support multiple windows.
// @cr must be untransformed from invoking of the draw function.
// This function will return %TRUE if the contents of the given
// @window are supposed to be drawn and %FALSE otherwise. Note
// that when the drawing was not initiated by the windowing
// system this function will return %TRUE for all windows, so
// you need to draw the bottommost window first. Also, do not
// use “else if” statements to check which window should be drawn.
/*

C function

gtk_cairo_should_draw_window
*/
func CairoShouldDrawWindow(cr *cairo.Context, window *gdk.Window) bool {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gtk_cairo_should_draw_window(c_cr, c_window)
	retGo := retC == C.TRUE

	return retGo
}

// Transforms the given cairo context @cr that from @widget-relative
// coordinates to @window-relative coordinates.
// If the @widget’s window is not an ancestor of @window, no
// modification will be applied.
//
// This is the inverse to the transformation GTK applies when
// preparing an expose event to be emitted with the #GtkWidget::draw
// signal. It is intended to help porting multiwindow widgets from
// GTK+ 2 to the rendering architecture of GTK+ 3.
/*

C function

gtk_cairo_transform_to_window
*/
func CairoTransformToWindow(cr *cairo.Context, widget *Widget, window *gdk.Window) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gtk_cairo_transform_to_window(c_cr, c_widget, c_window)

	return
}

// Adds a GTK+ grab on @device, so all the events on @device and its
// associated pointer or keyboard (if any) are delivered to @widget.
// If the @block_others parameter is %TRUE, any other devices will be
// unable to interact with @widget during the grab.
/*

C function

gtk_device_grab_add
*/
func DeviceGrabAdd(widget *Widget, device *gdk.Device, blockOthers bool) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_block_others :=
		boolToGboolean(blockOthers)

	C.gtk_device_grab_add(c_widget, c_device, c_block_others)

	return
}

// Removes a device grab from the given widget.
//
// You have to pair calls to gtk_device_grab_add() and
// gtk_device_grab_remove().
/*

C function

gtk_device_grab_remove
*/
func DeviceGrabRemove(widget *Widget, device *gdk.Device) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	C.gtk_device_grab_remove(c_widget, c_device)

	return
}

// Unsupported : gtk_draw_insertion_cursor : unsupported parameter location : Blacklisted record : GdkRectangle

// Returns the binary age as passed to `libtool`
// when building the GTK+ library the process is running against.
// If `libtool` means nothing to you, don't
// worry about it.
/*

C function

gtk_get_binary_age
*/
func GetBinaryAge() uint32 {
	retC := C.gtk_get_binary_age()
	retGo := (uint32)(retC)

	return retGo
}

// Returns the interface age as passed to `libtool`
// when building the GTK+ library the process is running against.
// If `libtool` means nothing to you, don't
// worry about it.
/*

C function

gtk_get_interface_age
*/
func GetInterfaceAge() uint32 {
	retC := C.gtk_get_interface_age()
	retGo := (uint32)(retC)

	return retGo
}

// Returns the major version number of the GTK+ library.
// (e.g. in GTK+ version 3.1.5 this is 3.)
//
// This function is in the library, so it represents the GTK+ library
// your code is running against. Contrast with the #GTK_MAJOR_VERSION
// macro, which represents the major version of the GTK+ headers you
// have included when compiling your code.
/*

C function

gtk_get_major_version
*/
func GetMajorVersion() uint32 {
	retC := C.gtk_get_major_version()
	retGo := (uint32)(retC)

	return retGo
}

// Returns the micro version number of the GTK+ library.
// (e.g. in GTK+ version 3.1.5 this is 5.)
//
// This function is in the library, so it represents the GTK+ library
// your code is are running against. Contrast with the
// #GTK_MICRO_VERSION macro, which represents the micro version of the
// GTK+ headers you have included when compiling your code.
/*

C function

gtk_get_micro_version
*/
func GetMicroVersion() uint32 {
	retC := C.gtk_get_micro_version()
	retGo := (uint32)(retC)

	return retGo
}

// Returns the minor version number of the GTK+ library.
// (e.g. in GTK+ version 3.1.5 this is 1.)
//
// This function is in the library, so it represents the GTK+ library
// your code is are running against. Contrast with the
// #GTK_MINOR_VERSION macro, which represents the minor version of the
// GTK+ headers you have included when compiling your code.
/*

C function

gtk_get_minor_version
*/
func GetMinorVersion() uint32 {
	retC := C.gtk_get_minor_version()
	retGo := (uint32)(retC)

	return retGo
}

// Renders an activity indicator (such as in #GtkSpinner).
// The state %GTK_STATE_FLAG_CHECKED determines whether there is
// activity going on.
/*

C function

gtk_render_activity
*/
func RenderActivity(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_activity(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// Renders an arrow pointing to @angle.
//
// Typical arrow rendering at 0, 1⁄2 π;, π; and 3⁄2 π:
//
// ![](arrows.png)
/*

C function

gtk_render_arrow
*/
func RenderArrow(context *StyleContext, cr *cairo.Context, angle float64, x float64, y float64, size float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_angle := (C.gdouble)(angle)

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_size := (C.gdouble)(size)

	C.gtk_render_arrow(c_context, c_cr, c_angle, c_x, c_y, c_size)

	return
}

// Renders a checkmark (as in a #GtkCheckButton).
//
// The %GTK_STATE_FLAG_CHECKED state determines whether the check is
// on or off, and %GTK_STATE_FLAG_INCONSISTENT determines whether it
// should be marked as undefined.
//
// Typical checkmark rendering:
//
// ![](checks.png)
/*

C function

gtk_render_check
*/
func RenderCheck(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_check(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// Renders an expander (as used in #GtkTreeView and #GtkExpander) in the area
// defined by @x, @y, @width, @height. The state %GTK_STATE_FLAG_CHECKED
// determines whether the expander is collapsed or expanded.
//
// Typical expander rendering:
//
// ![](expanders.png)
/*

C function

gtk_render_expander
*/
func RenderExpander(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_expander(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// Renders a extension (as in a #GtkNotebook tab) in the rectangle
// defined by @x, @y, @width, @height. The side where the extension
// connects to is defined by @gap_side.
//
// Typical extension rendering:
//
// ![](extensions.png)
/*

C function

gtk_render_extension
*/
func RenderExtension(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	C.gtk_render_extension(c_context, c_cr, c_x, c_y, c_width, c_height, c_gap_side)

	return
}

// Renders a focus indicator on the rectangle determined by @x, @y, @width, @height.
//
// Typical focus rendering:
//
// ![](focus.png)
/*

C function

gtk_render_focus
*/
func RenderFocus(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_focus(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// Renders a frame around the rectangle defined by @x, @y, @width, @height.
//
// Examples of frame rendering, showing the effect of `border-image`,
// `border-color`, `border-width`, `border-radius` and junctions:
//
// ![](frames.png)
/*

C function

gtk_render_frame
*/
func RenderFrame(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_frame(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// Renders a frame around the rectangle defined by (@x, @y, @width, @height),
// leaving a gap on one side. @xy0_gap and @xy1_gap will mean X coordinates
// for %GTK_POS_TOP and %GTK_POS_BOTTOM gap sides, and Y coordinates for
// %GTK_POS_LEFT and %GTK_POS_RIGHT.
//
// Typical rendering of a frame with a gap:
//
// ![](frame-gap.png)
/*

C function

gtk_render_frame_gap
*/
func RenderFrameGap(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType, xy0Gap float64, xy1Gap float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	c_xy0_gap := (C.gdouble)(xy0Gap)

	c_xy1_gap := (C.gdouble)(xy1Gap)

	C.gtk_render_frame_gap(c_context, c_cr, c_x, c_y, c_width, c_height, c_gap_side, c_xy0_gap, c_xy1_gap)

	return
}

// Renders a handle (as in #GtkHandleBox, #GtkPaned and
// #GtkWindow’s resize grip), in the rectangle
// determined by @x, @y, @width, @height.
//
// Handles rendered for the paned and grip classes:
//
// ![](handles.png)
/*

C function

gtk_render_handle
*/
func RenderHandle(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_handle(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// Renders the icon specified by @source at the given @size, returning the result
// in a pixbuf.
/*

C function

gtk_render_icon_pixbuf
*/
func RenderIconPixbuf(context *StyleContext, source *IconSource, size IconSize) *gdkpixbuf.Pixbuf {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_source := (*C.GtkIconSource)(C.NULL)
	if source != nil {
		c_source = (*C.GtkIconSource)(source.ToC())
	}

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_render_icon_pixbuf(c_context, c_source, c_size)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Renders @layout on the coordinates @x, @y
/*

C function

gtk_render_layout
*/
func RenderLayout(context *StyleContext, cr *cairo.Context, x float64, y float64, layout *pango.Layout) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	C.gtk_render_layout(c_context, c_cr, c_x, c_y, c_layout)

	return
}

// Renders a line from (x0, y0) to (x1, y1).
/*

C function

gtk_render_line
*/
func RenderLine(context *StyleContext, cr *cairo.Context, x0 float64, y0 float64, x1 float64, y1 float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	C.gtk_render_line(c_context, c_cr, c_x0, c_y0, c_x1, c_y1)

	return
}

// Renders an option mark (as in a #GtkRadioButton), the %GTK_STATE_FLAG_CHECKED
// state will determine whether the option is on or off, and
// %GTK_STATE_FLAG_INCONSISTENT whether it should be marked as undefined.
//
// Typical option mark rendering:
//
// ![](options.png)
/*

C function

gtk_render_option
*/
func RenderOption(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_option(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// Renders a slider (as in #GtkScale) in the rectangle defined by @x, @y,
// @width, @height. @orientation defines whether the slider is vertical
// or horizontal.
//
// Typical slider rendering:
//
// ![](sliders.png)
/*

C function

gtk_render_slider
*/
func RenderSlider(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, orientation Orientation) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_render_slider(c_context, c_cr, c_x, c_y, c_width, c_height, c_orientation)

	return
}
