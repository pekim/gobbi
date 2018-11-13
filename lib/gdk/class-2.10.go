// This is a generated file - DO NOT EDIT
// +build gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void screen_compositedChangedHandler(GObject *, gpointer);

	static gulong Screen_signal_connect_composited_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "composited-changed", G_CALLBACK(screen_compositedChangedHandler), data);
	}

*/
import "C"

// Returns %TRUE if gdk_window_input_shape_combine_mask() can
// be used to modify the input shape of windows on @display.
/*

C function

gdk_display_supports_input_shapes
*/
func (recv *Display) SupportsInputShapes() bool {
	retC := C.gdk_display_supports_input_shapes((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if gdk_window_shape_combine_mask() can
// be used to create shaped windows on @display.
/*

C function

gdk_display_supports_shapes
*/
func (recv *Display) SupportsShapes() bool {
	retC := C.gdk_display_supports_shapes((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

type signalScreenCompositedChangedDetail struct {
	callback  ScreenSignalCompositedChangedCallback
	handlerID C.gulong
}

var signalScreenCompositedChangedId int
var signalScreenCompositedChangedMap = make(map[int]signalScreenCompositedChangedDetail)
var signalScreenCompositedChangedLock sync.Mutex

// ScreenSignalCompositedChangedCallback is a callback function for a 'composited-changed' signal emitted from a Screen.
type ScreenSignalCompositedChangedCallback func()

/*
ConnectCompositedChanged connects the callback to the 'composited-changed' signal for the Screen.

The returned value represents the connection, and may be passed to DisconnectCompositedChanged to remove it.
*/
func (recv *Screen) ConnectCompositedChanged(callback ScreenSignalCompositedChangedCallback) int {
	signalScreenCompositedChangedLock.Lock()
	defer signalScreenCompositedChangedLock.Unlock()

	signalScreenCompositedChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Screen_signal_connect_composited_changed(instance, C.gpointer(uintptr(signalScreenCompositedChangedId)))

	detail := signalScreenCompositedChangedDetail{callback, handlerID}
	signalScreenCompositedChangedMap[signalScreenCompositedChangedId] = detail

	return signalScreenCompositedChangedId
}

/*
DisconnectCompositedChanged disconnects a callback from the 'composited-changed' signal for the Screen.

The connectionID should be a value returned from a call to ConnectCompositedChanged.
*/
func (recv *Screen) DisconnectCompositedChanged(connectionID int) {
	signalScreenCompositedChangedLock.Lock()
	defer signalScreenCompositedChangedLock.Unlock()

	detail, exists := signalScreenCompositedChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScreenCompositedChangedMap, connectionID)
}

//export screen_compositedChangedHandler
func screen_compositedChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalScreenCompositedChangedMap[index].callback
	callback()
}

// Returns the screen’s currently active window.
//
// On X11, this is done by inspecting the _NET_ACTIVE_WINDOW property
// on the root window, as described in the
// [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec).
// If there is no currently currently active
// window, or the window manager does not support the
// _NET_ACTIVE_WINDOW hint, this function returns %NULL.
//
// On other platforms, this function may return %NULL, depending on whether
// it is implementable on that platform.
//
// The returned window should be unrefed using g_object_unref() when
// no longer needed.
/*

C function

gdk_screen_get_active_window
*/
func (recv *Screen) GetActiveWindow() *Window {
	retC := C.gdk_screen_get_active_window((*C.GdkScreen)(recv.native))
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets any options previously set with gdk_screen_set_font_options().
/*

C function

gdk_screen_get_font_options
*/
func (recv *Screen) GetFontOptions() *cairo.FontOptions {
	retC := C.gdk_screen_get_font_options((*C.GdkScreen)(recv.native))
	var retGo (*cairo.FontOptions)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.FontOptionsNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the resolution for font handling on the screen; see
// gdk_screen_set_resolution() for full details.
/*

C function

gdk_screen_get_resolution
*/
func (recv *Screen) GetResolution() float64 {
	retC := C.gdk_screen_get_resolution((*C.GdkScreen)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Returns a #GList of #GdkWindows representing the current
// window stack.
//
// On X11, this is done by inspecting the _NET_CLIENT_LIST_STACKING
// property on the root window, as described in the
// [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec).
// If the window manager does not support the
// _NET_CLIENT_LIST_STACKING hint, this function returns %NULL.
//
// On other platforms, this function may return %NULL, depending on whether
// it is implementable on that platform.
//
// The returned list is newly allocated and owns references to the
// windows it contains, so it should be freed using g_list_free() and
// its windows unrefed using g_object_unref() when no longer needed.
/*

C function

gdk_screen_get_window_stack
*/
func (recv *Screen) GetWindowStack() *glib.List {
	retC := C.gdk_screen_get_window_stack((*C.GdkScreen)(recv.native))
	var retGo (*glib.List)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.ListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns whether windows with an RGBA visual can reasonably
// be expected to have their alpha channel drawn correctly on
// the screen.
//
// On X11 this function returns whether a compositing manager is
// compositing @screen.
/*

C function

gdk_screen_is_composited
*/
func (recv *Screen) IsComposited() bool {
	retC := C.gdk_screen_is_composited((*C.GdkScreen)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the default font options for the screen. These
// options will be set on any #PangoContext’s newly created
// with gdk_pango_context_get_for_screen(). Changing the
// default set of font options does not affect contexts that
// have already been created.
/*

C function

gdk_screen_set_font_options
*/
func (recv *Screen) SetFontOptions(options *cairo.FontOptions) {
	c_options := (*C.cairo_font_options_t)(C.NULL)
	if options != nil {
		c_options = (*C.cairo_font_options_t)(options.ToC())
	}

	C.gdk_screen_set_font_options((*C.GdkScreen)(recv.native), c_options)

	return
}

// Sets the resolution for font handling on the screen. This is a
// scale factor between points specified in a #PangoFontDescription
// and cairo units. The default value is 96, meaning that a 10 point
// font will be 13 units high. (10 * 96. / 72. = 13.3).
/*

C function

gdk_screen_set_resolution
*/
func (recv *Screen) SetResolution(dpi float64) {
	c_dpi := (C.gdouble)(dpi)

	C.gdk_screen_set_resolution((*C.GdkScreen)(recv.native), c_dpi)

	return
}

// This function returns the type hint set for a window.
/*

C function

gdk_window_get_type_hint
*/
func (recv *Window) GetTypeHint() WindowTypeHint {
	retC := C.gdk_window_get_type_hint((*C.GdkWindow)(recv.native))
	retGo := (WindowTypeHint)(retC)

	return retGo
}

// Like gdk_window_shape_combine_region(), but the shape applies
// only to event handling. Mouse events which happen while
// the pointer position corresponds to an unset bit in the
// mask will be passed on the window below @window.
//
// An input shape is typically used with RGBA windows.
// The alpha channel of the window defines which pixels are
// invisible and allows for nicely antialiased borders,
// and the input shape controls where the window is
// “clickable”.
//
// On the X11 platform, this requires version 1.1 of the
// shape extension.
//
// On the Win32 platform, this functionality is not present and the
// function does nothing.
/*

C function

gdk_window_input_shape_combine_region
*/
func (recv *Window) InputShapeCombineRegion(shapeRegion *cairo.Region, offsetX int32, offsetY int32) {
	c_shape_region := (*C.cairo_region_t)(C.NULL)
	if shapeRegion != nil {
		c_shape_region = (*C.cairo_region_t)(shapeRegion.ToC())
	}

	c_offset_x := (C.gint)(offsetX)

	c_offset_y := (C.gint)(offsetY)

	C.gdk_window_input_shape_combine_region((*C.GdkWindow)(recv.native), c_shape_region, c_offset_x, c_offset_y)

	return
}

// Merges the input shape masks for any child windows into the
// input shape mask for @window. i.e. the union of all input masks
// for @window and its children will become the new input mask
// for @window. See gdk_window_input_shape_combine_region().
//
// This function is distinct from gdk_window_set_child_input_shapes()
// because it includes @window’s input shape mask in the set of
// shapes to be merged.
/*

C function

gdk_window_merge_child_input_shapes
*/
func (recv *Window) MergeChildInputShapes() {
	C.gdk_window_merge_child_input_shapes((*C.GdkWindow)(recv.native))

	return
}

// Sets the input shape mask of @window to the union of input shape masks
// for all children of @window, ignoring the input shape mask of @window
// itself. Contrast with gdk_window_merge_child_input_shapes() which includes
// the input shape mask of @window in the masks to be merged.
/*

C function

gdk_window_set_child_input_shapes
*/
func (recv *Window) SetChildInputShapes() {
	C.gdk_window_set_child_input_shapes((*C.GdkWindow)(recv.native))

	return
}
