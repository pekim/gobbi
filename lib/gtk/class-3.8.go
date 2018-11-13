// This is a generated file - DO NOT EDIT
// +build gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Add @object to the @builder object pool so it can be referenced just like any
// other object built by builder.
/*

C function

gtk_builder_expose_object
*/
func (recv *Builder) ExposeObject(name string, object *gobject.Object) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object := (*C.GObject)(C.NULL)
	if object != nil {
		c_object = (*C.GObject)(object.ToC())
	}

	C.gtk_builder_expose_object((*C.GtkBuilder)(recv.native), c_name, c_object)

	return
}

// Unsupported : gtk_icon_info_load_icon_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an async icon load, see gtk_icon_info_load_icon_async().
/*

C function

gtk_icon_info_load_icon_finish
*/
func (recv *IconInfo) LoadIconFinish(res *gio.AsyncResult) (*gdkpixbuf.Pixbuf, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_icon_finish((*C.GtkIconInfo)(recv.native), c_res, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_icon_info_load_symbolic_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an async icon load, see gtk_icon_info_load_symbolic_async().
/*

C function

gtk_icon_info_load_symbolic_finish
*/
func (recv *IconInfo) LoadSymbolicFinish(res *gio.AsyncResult) (*gdkpixbuf.Pixbuf, bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic_finish((*C.GtkIconInfo)(recv.native), c_res, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goThrowableError
}

// Unsupported : gtk_icon_info_load_symbolic_for_context_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an async icon load, see gtk_icon_info_load_symbolic_for_context_async().
/*

C function

gtk_icon_info_load_symbolic_for_context_finish
*/
func (recv *IconInfo) LoadSymbolicForContextFinish(res *gio.AsyncResult) (*gdkpixbuf.Pixbuf, bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic_for_context_finish((*C.GtkIconInfo)(recv.native), c_res, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goThrowableError
}

// Gets the setting set by gtk_icon_view_set_activate_on_single_click().
/*

C function

gtk_icon_view_get_activate_on_single_click
*/
func (recv *IconView) GetActivateOnSingleClick() bool {
	retC := C.gtk_icon_view_get_activate_on_single_click((*C.GtkIconView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Causes the #GtkIconView::item-activated signal to be emitted on
// a single click instead of a double click.
/*

C function

gtk_icon_view_set_activate_on_single_click
*/
func (recv *IconView) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_icon_view_set_activate_on_single_click((*C.GtkIconView)(recv.native), c_single)

	return
}

// Return the value of the #GtkLevelBar:inverted property.
/*

C function

gtk_level_bar_get_inverted
*/
func (recv *LevelBar) GetInverted() bool {
	retC := C.gtk_level_bar_get_inverted((*C.GtkLevelBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the value of the #GtkLevelBar:inverted property.
/*

C function

gtk_level_bar_set_inverted
*/
func (recv *LevelBar) SetInverted(inverted bool) {
	c_inverted :=
		boolToGboolean(inverted)

	C.gtk_level_bar_set_inverted((*C.GtkLevelBar)(recv.native), c_inverted)

	return
}

// Returns the #GdkFrameClock to which @context is attached.
/*

C function

gtk_style_context_get_frame_clock
*/
func (recv *StyleContext) GetFrameClock() *gdk.FrameClock {
	retC := C.gtk_style_context_get_frame_clock((*C.GtkStyleContext)(recv.native))
	var retGo (*gdk.FrameClock)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdk.FrameClockNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Attaches @context to the given frame clock.
//
// The frame clock is used for the timing of animations.
//
// If you are using a #GtkStyleContext returned from
// gtk_widget_get_style_context(), you do not need to
// call this yourself.
/*

C function

gtk_style_context_set_frame_clock
*/
func (recv *StyleContext) SetFrameClock(frameClock *gdk.FrameClock) {
	c_frame_clock := (*C.GdkFrameClock)(C.NULL)
	if frameClock != nil {
		c_frame_clock = (*C.GdkFrameClock)(frameClock.ToC())
	}

	C.gtk_style_context_set_frame_clock((*C.GtkStyleContext)(recv.native), c_frame_clock)

	return
}

// Gets the setting set by gtk_tree_view_set_activate_on_single_click().
/*

C function

gtk_tree_view_get_activate_on_single_click
*/
func (recv *TreeView) GetActivateOnSingleClick() bool {
	retC := C.gtk_tree_view_get_activate_on_single_click((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Cause the #GtkTreeView::row-activated signal to be emitted
// on a single click instead of a double click.
/*

C function

gtk_tree_view_set_activate_on_single_click
*/
func (recv *TreeView) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_tree_view_set_activate_on_single_click((*C.GtkTreeView)(recv.native), c_single)

	return
}

// Unsupported : gtk_widget_add_tick_callback : unsupported parameter callback : no type generator for TickCallback (GtkTickCallback) for param callback

// Obtains the frame clock for a widget. The frame clock is a global
// “ticker” that can be used to drive animations and repaints.  The
// most common reason to get the frame clock is to call
// gdk_frame_clock_get_frame_time(), in order to get a time to use for
// animating. For example you might record the start of the animation
// with an initial value from gdk_frame_clock_get_frame_time(), and
// then update the animation by calling
// gdk_frame_clock_get_frame_time() again during each repaint.
//
// gdk_frame_clock_request_phase() will result in a new frame on the
// clock, but won’t necessarily repaint any widgets. To repaint a
// widget, you have to use gtk_widget_queue_draw() which invalidates
// the widget (thus scheduling it to receive a draw on the next
// frame). gtk_widget_queue_draw() will also end up requesting a frame
// on the appropriate frame clock.
//
// A widget’s frame clock will not change while the widget is
// mapped. Reparenting a widget (which implies a temporary unmap) can
// change the widget’s frame clock.
//
// Unrealized widgets do not have a frame clock.
/*

C function

gtk_widget_get_frame_clock
*/
func (recv *Widget) GetFrameClock() *gdk.FrameClock {
	retC := C.gtk_widget_get_frame_clock((*C.GtkWidget)(recv.native))
	var retGo (*gdk.FrameClock)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdk.FrameClockNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Fetches the requested opacity for this widget.
// See gtk_widget_set_opacity().
/*

C function

gtk_widget_get_opacity
*/
func (recv *Widget) GetOpacity() float64 {
	retC := C.gtk_widget_get_opacity((*C.GtkWidget)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Determines whether the widget and all its parents are marked as
// visible.
//
// This function does not check if the widget is obscured in any way.
//
// See also gtk_widget_get_visible() and gtk_widget_set_visible()
/*

C function

gtk_widget_is_visible
*/
func (recv *Widget) IsVisible() bool {
	retC := C.gtk_widget_is_visible((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Registers a #GdkWindow with the widget and sets it up so that
// the widget receives events for it. Call gtk_widget_unregister_window()
// when destroying the window.
//
// Before 3.8 you needed to call gdk_window_set_user_data() directly to set
// this up. This is now deprecated and you should use gtk_widget_register_window()
// instead. Old code will keep working as is, although some new features like
// transparency might not work perfectly.
/*

C function

gtk_widget_register_window
*/
func (recv *Widget) RegisterWindow(window *gdk.Window) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gtk_widget_register_window((*C.GtkWidget)(recv.native), c_window)

	return
}

// Removes a tick callback previously registered with
// gtk_widget_add_tick_callback().
/*

C function

gtk_widget_remove_tick_callback
*/
func (recv *Widget) RemoveTickCallback(id uint32) {
	c_id := (C.guint)(id)

	C.gtk_widget_remove_tick_callback((*C.GtkWidget)(recv.native), c_id)

	return
}

// Request the @widget to be rendered partially transparent,
// with opacity 0 being fully transparent and 1 fully opaque. (Opacity values
// are clamped to the [0,1] range.).
// This works on both toplevel widget, and child widgets, although there
// are some limitations:
//
// For toplevel widgets this depends on the capabilities of the windowing
// system. On X11 this has any effect only on X screens with a compositing manager
// running. See gtk_widget_is_composited(). On Windows it should work
// always, although setting a window’s opacity after the window has been
// shown causes it to flicker once on Windows.
//
// For child widgets it doesn’t work if any affected widget has a native window, or
// disables double buffering.
/*

C function

gtk_widget_set_opacity
*/
func (recv *Widget) SetOpacity(opacity float64) {
	c_opacity := (C.double)(opacity)

	C.gtk_widget_set_opacity((*C.GtkWidget)(recv.native), c_opacity)

	return
}

// Unregisters a #GdkWindow from the widget that was previously set up with
// gtk_widget_register_window(). You need to call this when the window is
// no longer used by the widget, such as when you destroy it.
/*

C function

gtk_widget_unregister_window
*/
func (recv *Widget) UnregisterWindow(window *gdk.Window) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gtk_widget_unregister_window((*C.GtkWidget)(recv.native), c_window)

	return
}
