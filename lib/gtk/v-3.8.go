// Code generated - DO NOT EDIT.
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
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// ExposeObject is a wrapper around the C function gtk_builder_expose_object.
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

// LoadIconFinish is a wrapper around the C function gtk_icon_info_load_icon_finish.
func (recv *IconInfo) LoadIconFinish(res *gio.AsyncResult) (*gdkpixbuf.Pixbuf, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_icon_finish((*C.GtkIconInfo)(recv.native), c_res, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : gtk_icon_info_load_symbolic_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LoadSymbolicFinish is a wrapper around the C function gtk_icon_info_load_symbolic_finish.
func (recv *IconInfo) LoadSymbolicFinish(res *gio.AsyncResult) (*gdkpixbuf.Pixbuf, bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic_finish((*C.GtkIconInfo)(recv.native), c_res, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goError
}

// Unsupported : gtk_icon_info_load_symbolic_for_context_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LoadSymbolicForContextFinish is a wrapper around the C function gtk_icon_info_load_symbolic_for_context_finish.
func (recv *IconInfo) LoadSymbolicForContextFinish(res *gio.AsyncResult) (*gdkpixbuf.Pixbuf, bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic_for_context_finish((*C.GtkIconInfo)(recv.native), c_res, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goError
}

// GetActivateOnSingleClick is a wrapper around the C function gtk_icon_view_get_activate_on_single_click.
func (recv *IconView) GetActivateOnSingleClick() bool {
	retC := C.gtk_icon_view_get_activate_on_single_click((*C.GtkIconView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActivateOnSingleClick is a wrapper around the C function gtk_icon_view_set_activate_on_single_click.
func (recv *IconView) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_icon_view_set_activate_on_single_click((*C.GtkIconView)(recv.native), c_single)

	return
}

// GetInverted is a wrapper around the C function gtk_level_bar_get_inverted.
func (recv *LevelBar) GetInverted() bool {
	retC := C.gtk_level_bar_get_inverted((*C.GtkLevelBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetInverted is a wrapper around the C function gtk_level_bar_set_inverted.
func (recv *LevelBar) SetInverted(inverted bool) {
	c_inverted :=
		boolToGboolean(inverted)

	C.gtk_level_bar_set_inverted((*C.GtkLevelBar)(recv.native), c_inverted)

	return
}

// GetFrameClock is a wrapper around the C function gtk_style_context_get_frame_clock.
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

// SetFrameClock is a wrapper around the C function gtk_style_context_set_frame_clock.
func (recv *StyleContext) SetFrameClock(frameClock *gdk.FrameClock) {
	c_frame_clock := (*C.GdkFrameClock)(C.NULL)
	if frameClock != nil {
		c_frame_clock = (*C.GdkFrameClock)(frameClock.ToC())
	}

	C.gtk_style_context_set_frame_clock((*C.GtkStyleContext)(recv.native), c_frame_clock)

	return
}

// GetActivateOnSingleClick is a wrapper around the C function gtk_tree_view_get_activate_on_single_click.
func (recv *TreeView) GetActivateOnSingleClick() bool {
	retC := C.gtk_tree_view_get_activate_on_single_click((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActivateOnSingleClick is a wrapper around the C function gtk_tree_view_set_activate_on_single_click.
func (recv *TreeView) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_tree_view_set_activate_on_single_click((*C.GtkTreeView)(recv.native), c_single)

	return
}

// Unsupported : gtk_widget_add_tick_callback : unsupported parameter callback : no type generator for TickCallback (GtkTickCallback) for param callback

// GetFrameClock is a wrapper around the C function gtk_widget_get_frame_clock.
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

// GetOpacity is a wrapper around the C function gtk_widget_get_opacity.
func (recv *Widget) GetOpacity() float64 {
	retC := C.gtk_widget_get_opacity((*C.GtkWidget)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// IsVisible is a wrapper around the C function gtk_widget_is_visible.
func (recv *Widget) IsVisible() bool {
	retC := C.gtk_widget_is_visible((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// RegisterWindow is a wrapper around the C function gtk_widget_register_window.
func (recv *Widget) RegisterWindow(window *gdk.Window) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gtk_widget_register_window((*C.GtkWidget)(recv.native), c_window)

	return
}

// RemoveTickCallback is a wrapper around the C function gtk_widget_remove_tick_callback.
func (recv *Widget) RemoveTickCallback(id uint32) {
	c_id := (C.guint)(id)

	C.gtk_widget_remove_tick_callback((*C.GtkWidget)(recv.native), c_id)

	return
}

// SetOpacity is a wrapper around the C function gtk_widget_set_opacity.
func (recv *Widget) SetOpacity(opacity float64) {
	c_opacity := (C.double)(opacity)

	C.gtk_widget_set_opacity((*C.GtkWidget)(recv.native), c_opacity)

	return
}

// UnregisterWindow is a wrapper around the C function gtk_widget_unregister_window.
func (recv *Widget) UnregisterWindow(window *gdk.Window) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gtk_widget_unregister_window((*C.GtkWidget)(recv.native), c_window)

	return
}
