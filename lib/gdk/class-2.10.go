// This is a generated file - DO NOT EDIT
// +build gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func (recv *AppLaunchContext) AppLaunchContext() *gio.AppLaunchContext {}

func (recv *Cursor) Object() *gobject.Object {}

func (recv *Device) Object() *gobject.Object {}

func (recv *DeviceManager) Object() *gobject.Object {}

func (recv *DeviceTool) Object() *gobject.Object {}

// SupportsInputShapes is a wrapper around the C function gdk_display_supports_input_shapes.
func (recv *Display) SupportsInputShapes() bool {
	retC := C.gdk_display_supports_input_shapes((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsShapes is a wrapper around the C function gdk_display_supports_shapes.
func (recv *Display) SupportsShapes() bool {
	retC := C.gdk_display_supports_shapes((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

func (recv *Display) Object() *gobject.Object {}

func (recv *DisplayManager) Object() *gobject.Object {}

func (recv *DragContext) Object() *gobject.Object {}

func (recv *DrawingContext) Object() *gobject.Object {}

func (recv *FrameClock) Object() *gobject.Object {}

func (recv *GLContext) Object() *gobject.Object {}

func (recv *Keymap) Object() *gobject.Object {}

func (recv *Monitor) Object() *gobject.Object {}

// GetActiveWindow is a wrapper around the C function gdk_screen_get_active_window.
func (recv *Screen) GetActiveWindow() *Window {
	retC := C.gdk_screen_get_active_window((*C.GdkScreen)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFontOptions is a wrapper around the C function gdk_screen_get_font_options.
func (recv *Screen) GetFontOptions() *cairo.FontOptions {
	retC := C.gdk_screen_get_font_options((*C.GdkScreen)(recv.native))
	retGo := cairo.FontOptionsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetResolution is a wrapper around the C function gdk_screen_get_resolution.
func (recv *Screen) GetResolution() float64 {
	retC := C.gdk_screen_get_resolution((*C.GdkScreen)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetWindowStack is a wrapper around the C function gdk_screen_get_window_stack.
func (recv *Screen) GetWindowStack() *glib.List {
	retC := C.gdk_screen_get_window_stack((*C.GdkScreen)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsComposited is a wrapper around the C function gdk_screen_is_composited.
func (recv *Screen) IsComposited() bool {
	retC := C.gdk_screen_is_composited((*C.GdkScreen)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFontOptions is a wrapper around the C function gdk_screen_set_font_options.
func (recv *Screen) SetFontOptions(options *cairo.FontOptions) {
	c_options := (*C.cairo_font_options_t)(options.ToC())

	C.gdk_screen_set_font_options((*C.GdkScreen)(recv.native), c_options)

	return
}

// SetResolution is a wrapper around the C function gdk_screen_set_resolution.
func (recv *Screen) SetResolution(dpi float64) {
	c_dpi := (C.gdouble)(dpi)

	C.gdk_screen_set_resolution((*C.GdkScreen)(recv.native), c_dpi)

	return
}

func (recv *Screen) Object() *gobject.Object {}

func (recv *Seat) Object() *gobject.Object {}

func (recv *Visual) Object() *gobject.Object {}

// GetTypeHint is a wrapper around the C function gdk_window_get_type_hint.
func (recv *Window) GetTypeHint() WindowTypeHint {
	retC := C.gdk_window_get_type_hint((*C.GdkWindow)(recv.native))
	retGo := (WindowTypeHint)(retC)

	return retGo
}

// InputShapeCombineRegion is a wrapper around the C function gdk_window_input_shape_combine_region.
func (recv *Window) InputShapeCombineRegion(shapeRegion *cairo.Region, offsetX int32, offsetY int32) {
	c_shape_region := (*C.cairo_region_t)(shapeRegion.ToC())

	c_offset_x := (C.gint)(offsetX)

	c_offset_y := (C.gint)(offsetY)

	C.gdk_window_input_shape_combine_region((*C.GdkWindow)(recv.native), c_shape_region, c_offset_x, c_offset_y)

	return
}

// MergeChildInputShapes is a wrapper around the C function gdk_window_merge_child_input_shapes.
func (recv *Window) MergeChildInputShapes() {
	C.gdk_window_merge_child_input_shapes((*C.GdkWindow)(recv.native))

	return
}

// SetChildInputShapes is a wrapper around the C function gdk_window_set_child_input_shapes.
func (recv *Window) SetChildInputShapes() {
	C.gdk_window_set_child_input_shapes((*C.GdkWindow)(recv.native))

	return
}

func (recv *Window) Object() *gobject.Object {}
