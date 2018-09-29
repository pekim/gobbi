// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_app_launch_context_set_icon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gdk_cursor_get_surface : unsupported parameter x_hot : no type generator for gdouble, gdouble*

// GetAxes is a wrapper around the C function gdk_device_get_axes.
func (recv *Device) GetAxes() AxisFlags {
	retC := C.gdk_device_get_axes((*C.GdkDevice)(recv.native))
	retGo := (AxisFlags)(retC)

	return retGo
}

// Unsupported : gdk_device_get_axis : unsupported parameter axes : no param type

// Unsupported : gdk_device_get_axis_value : unsupported parameter axes : no param type

// Unsupported : gdk_device_get_history : unsupported parameter events : no param type

// Unsupported : gdk_device_get_key : unsupported parameter keyval : no type generator for guint, guint*

// Unsupported : gdk_device_get_position : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_device_get_position_double : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_device_get_state : unsupported parameter axes : no param type

// Unsupported : gdk_device_get_window_at_position : unsupported parameter win_x : no type generator for gint, gint*

// Unsupported : gdk_device_get_window_at_position_double : unsupported parameter win_x : no type generator for gdouble, gdouble*

// DeviceTool is a wrapper around the C record GdkDeviceTool.
type DeviceTool struct {
	native *C.GdkDeviceTool
}

func DeviceToolNewFromC(u unsafe.Pointer) *DeviceTool {
	c := (*C.GdkDeviceTool)(u)
	if c == nil {
		return nil
	}

	g := &DeviceTool{native: c}

	return g
}

func (recv *DeviceTool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetHardwareId is a wrapper around the C function gdk_device_tool_get_hardware_id.
func (recv *DeviceTool) GetHardwareId() uint64 {
	retC := C.gdk_device_tool_get_hardware_id((*C.GdkDeviceTool)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetSerial is a wrapper around the C function gdk_device_tool_get_serial.
func (recv *DeviceTool) GetSerial() uint64 {
	retC := C.gdk_device_tool_get_serial((*C.GdkDeviceTool)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetToolType is a wrapper around the C function gdk_device_tool_get_tool_type.
func (recv *DeviceTool) GetToolType() DeviceToolType {
	retC := C.gdk_device_tool_get_tool_type((*C.GdkDeviceTool)(recv.native))
	retGo := (DeviceToolType)(retC)

	return retGo
}

// Unsupported : gdk_display_get_event : no return generator

// Unsupported : gdk_display_get_maximal_cursor_size : unsupported parameter width : no type generator for guint, guint*

// GetMonitor is a wrapper around the C function gdk_display_get_monitor.
func (recv *Display) GetMonitor(monitorNum int32) *Monitor {
	c_monitor_num := (C.int)(monitorNum)

	retC := C.gdk_display_get_monitor((*C.GdkDisplay)(recv.native), c_monitor_num)
	retGo := MonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMonitorAtPoint is a wrapper around the C function gdk_display_get_monitor_at_point.
func (recv *Display) GetMonitorAtPoint(x int32, y int32) *Monitor {
	c_x := (C.int)(x)

	c_y := (C.int)(y)

	retC := C.gdk_display_get_monitor_at_point((*C.GdkDisplay)(recv.native), c_x, c_y)
	retGo := MonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMonitorAtWindow is a wrapper around the C function gdk_display_get_monitor_at_window.
func (recv *Display) GetMonitorAtWindow(window *Window) *Monitor {
	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gdk_display_get_monitor_at_window((*C.GdkDisplay)(recv.native), c_window)
	retGo := MonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetNMonitors is a wrapper around the C function gdk_display_get_n_monitors.
func (recv *Display) GetNMonitors() int32 {
	retC := C.gdk_display_get_n_monitors((*C.GdkDisplay)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_display_get_pointer : unsupported parameter screen : record with indirection level of 2

// GetPrimaryMonitor is a wrapper around the C function gdk_display_get_primary_monitor.
func (recv *Display) GetPrimaryMonitor() *Monitor {
	retC := C.gdk_display_get_primary_monitor((*C.GdkDisplay)(recv.native))
	retGo := MonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_display_get_window_at_pointer : unsupported parameter win_x : no type generator for gint, gint*

// Unsupported : gdk_display_peek_event : no return generator

// Unsupported : gdk_display_put_event : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_display_request_selection_notification : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets : no param type

// DrawingContext is a wrapper around the C record GdkDrawingContext.
type DrawingContext struct {
	native *C.GdkDrawingContext
}

func DrawingContextNewFromC(u unsafe.Pointer) *DrawingContext {
	c := (*C.GdkDrawingContext)(u)
	if c == nil {
		return nil
	}

	g := &DrawingContext{native: c}

	return g
}

func (recv *DrawingContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetCairoContext is a wrapper around the C function gdk_drawing_context_get_cairo_context.
func (recv *DrawingContext) GetCairoContext() *cairo.Context {
	retC := C.gdk_drawing_context_get_cairo_context((*C.GdkDrawingContext)(recv.native))
	retGo := cairo.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetClip is a wrapper around the C function gdk_drawing_context_get_clip.
func (recv *DrawingContext) GetClip() *cairo.Region {
	retC := C.gdk_drawing_context_get_clip((*C.GdkDrawingContext)(recv.native))
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWindow is a wrapper around the C function gdk_drawing_context_get_window.
func (recv *DrawingContext) GetWindow() *Window {
	retC := C.gdk_drawing_context_get_window((*C.GdkDrawingContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsValid is a wrapper around the C function gdk_drawing_context_is_valid.
func (recv *DrawingContext) IsValid() bool {
	retC := C.gdk_drawing_context_is_valid((*C.GdkDrawingContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_frame_clock_get_refresh_info : unsupported parameter refresh_interval_return : no type generator for gint64, gint64*

// Unsupported : gdk_gl_context_get_required_version : unsupported parameter major : no type generator for gint, int*

// GetUseEs is a wrapper around the C function gdk_gl_context_get_use_es.
func (recv *GLContext) GetUseEs() bool {
	retC := C.gdk_gl_context_get_use_es((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_gl_context_get_version : unsupported parameter major : no type generator for gint, int*

// SetUseEs is a wrapper around the C function gdk_gl_context_set_use_es.
func (recv *GLContext) SetUseEs(useEs int32) {
	c_use_es := (C.int)(useEs)

	C.gdk_gl_context_set_use_es((*C.GdkGLContext)(recv.native), c_use_es)

	return
}

// Unsupported : gdk_keymap_add_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gdk_keymap_get_entries_for_keycode : unsupported parameter keys : no param type

// Unsupported : gdk_keymap_get_entries_for_keyval : unsupported parameter keys : no param type

// Unsupported : gdk_keymap_map_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gdk_keymap_translate_keyboard_state : unsupported parameter keyval : no type generator for guint, guint*

// Monitor is a wrapper around the C record GdkMonitor.
type Monitor struct {
	native *C.GdkMonitor
}

func MonitorNewFromC(u unsafe.Pointer) *Monitor {
	c := (*C.GdkMonitor)(u)
	if c == nil {
		return nil
	}

	g := &Monitor{native: c}

	return g
}

func (recv *Monitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetDisplay is a wrapper around the C function gdk_monitor_get_display.
func (recv *Monitor) GetDisplay() *Display {
	retC := C.gdk_monitor_get_display((*C.GdkMonitor)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_monitor_get_geometry : unsupported parameter geometry : Blacklisted record : GdkRectangle

// GetHeightMm is a wrapper around the C function gdk_monitor_get_height_mm.
func (recv *Monitor) GetHeightMm() int32 {
	retC := C.gdk_monitor_get_height_mm((*C.GdkMonitor)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetManufacturer is a wrapper around the C function gdk_monitor_get_manufacturer.
func (recv *Monitor) GetManufacturer() string {
	retC := C.gdk_monitor_get_manufacturer((*C.GdkMonitor)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetModel is a wrapper around the C function gdk_monitor_get_model.
func (recv *Monitor) GetModel() string {
	retC := C.gdk_monitor_get_model((*C.GdkMonitor)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetRefreshRate is a wrapper around the C function gdk_monitor_get_refresh_rate.
func (recv *Monitor) GetRefreshRate() int32 {
	retC := C.gdk_monitor_get_refresh_rate((*C.GdkMonitor)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetScaleFactor is a wrapper around the C function gdk_monitor_get_scale_factor.
func (recv *Monitor) GetScaleFactor() int32 {
	retC := C.gdk_monitor_get_scale_factor((*C.GdkMonitor)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSubpixelLayout is a wrapper around the C function gdk_monitor_get_subpixel_layout.
func (recv *Monitor) GetSubpixelLayout() SubpixelLayout {
	retC := C.gdk_monitor_get_subpixel_layout((*C.GdkMonitor)(recv.native))
	retGo := (SubpixelLayout)(retC)

	return retGo
}

// GetWidthMm is a wrapper around the C function gdk_monitor_get_width_mm.
func (recv *Monitor) GetWidthMm() int32 {
	retC := C.gdk_monitor_get_width_mm((*C.GdkMonitor)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_monitor_get_workarea : unsupported parameter workarea : Blacklisted record : GdkRectangle

// IsPrimary is a wrapper around the C function gdk_monitor_is_primary.
func (recv *Monitor) IsPrimary() bool {
	retC := C.gdk_monitor_is_primary((*C.GdkMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_screen_get_monitor_geometry : unsupported parameter dest : Blacklisted record : GdkRectangle

// Unsupported : gdk_screen_get_monitor_workarea : unsupported parameter dest : Blacklisted record : GdkRectangle

// Seat is a wrapper around the C record GdkSeat.
type Seat struct {
	native *C.GdkSeat
	// parent_instance : record
}

func SeatNewFromC(u unsafe.Pointer) *Seat {
	c := (*C.GdkSeat)(u)
	if c == nil {
		return nil
	}

	g := &Seat{native: c}

	return g
}

func (recv *Seat) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetDisplay is a wrapper around the C function gdk_seat_get_display.
func (recv *Seat) GetDisplay() *Display {
	retC := C.gdk_seat_get_display((*C.GdkSeat)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_seat_grab : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_visual_get_blue_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_green_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_red_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// BeginDrawFrame is a wrapper around the C function gdk_window_begin_draw_frame.
func (recv *Window) BeginDrawFrame(region *cairo.Region) *DrawingContext {
	c_region := (*C.cairo_region_t)(region.ToC())

	retC := C.gdk_window_begin_draw_frame((*C.GdkWindow)(recv.native), c_region)
	retGo := DrawingContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_begin_paint_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_coords_from_parent : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_coords_to_parent : unsupported parameter parent_x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_create_similar_image_surface : unsupported parameter format : no type generator for gint, cairo_format_t

// EndDrawFrame is a wrapper around the C function gdk_window_end_draw_frame.
func (recv *Window) EndDrawFrame(context *DrawingContext) {
	c_context := (*C.GdkDrawingContext)(context.ToC())

	C.gdk_window_end_draw_frame((*C.GdkWindow)(recv.native), c_context)

	return
}

// Unsupported : gdk_window_get_decorations : unsupported parameter decorations : GdkWMDecoration* with indirection level of 1

// Unsupported : gdk_window_get_device_position : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_device_position_double : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_get_drag_protocol : unsupported parameter target : record with indirection level of 2

// Unsupported : gdk_window_get_frame_extents : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_get_geometry : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_origin : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_pointer : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_position : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_root_coords : unsupported parameter root_x : no type generator for gint, gint*

// Unsupported : gdk_window_get_root_origin : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_user_data : unsupported parameter data : no type generator for gpointer, gpointer*

// Unsupported : gdk_window_invalidate_maybe_recurse : unsupported parameter child_func : no type generator for WindowChildFunc, GdkWindowChildFunc

// Unsupported : gdk_window_invalidate_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_set_invalidate_handler : unsupported parameter handler : no type generator for WindowInvalidateHandlerFunc, GdkWindowInvalidateHandlerFunc

// Unsupported : gdk_window_show_window_menu : unsupported parameter event : no type generator for Event, GdkEvent*
