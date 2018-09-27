// This is a generated file - DO NOT EDIT

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	glib "github.com/pekim/gobbi/lib/glib"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// AppLaunchContext is a wrapper around the C record GdkAppLaunchContext.
type AppLaunchContext struct {
	native *C.GdkAppLaunchContext
}

func AppLaunchContextNewFromC(u unsafe.Pointer) *AppLaunchContext {
	c := (*C.GdkAppLaunchContext)(u)
	if c == nil {
		return nil
	}

	g := &AppLaunchContext{native: c}

	return g
}

func (recv *AppLaunchContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_app_launch_context_set_desktop : no return generator

// Unsupported : gdk_app_launch_context_set_display : no return generator

// Unsupported : gdk_app_launch_context_set_icon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gdk_app_launch_context_set_icon_name : no return generator

// Unsupported : gdk_app_launch_context_set_screen : no return generator

// Unsupported : gdk_app_launch_context_set_timestamp : no return generator

// Cursor is a wrapper around the C record GdkCursor.
type Cursor struct {
	native *C.GdkCursor
}

func CursorNewFromC(u unsafe.Pointer) *Cursor {
	c := (*C.GdkCursor)(u)
	if c == nil {
		return nil
	}

	g := &Cursor{native: c}

	return g
}

func (recv *Cursor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CursorNew is a wrapper around the C function gdk_cursor_new.
func CursorNew(cursorType CursorType) *Cursor {
	c_cursor_type := (C.GdkCursorType)(cursorType)

	retC := C.gdk_cursor_new(c_cursor_type)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_cursor_get_surface : unsupported parameter x_hot : no type generator for gdouble, gdouble*

// Ref is a wrapper around the C function gdk_cursor_ref.
func (recv *Cursor) Ref() *Cursor {
	retC := C.gdk_cursor_ref((*C.GdkCursor)(recv.native))
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_cursor_unref : no return generator

// Device is a wrapper around the C record GdkDevice.
type Device struct {
	native *C.GdkDevice
}

func DeviceNewFromC(u unsafe.Pointer) *Device {
	c := (*C.GdkDevice)(u)
	if c == nil {
		return nil
	}

	g := &Device{native: c}

	return g
}

func (recv *Device) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// ListSlaveDevices is a wrapper around the C function gdk_device_list_slave_devices.
func (recv *Device) ListSlaveDevices() *glib.List {
	retC := C.gdk_device_list_slave_devices((*C.GdkDevice)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_device_set_axis_use : no return generator

// Unsupported : gdk_device_set_key : no return generator

// SetMode is a wrapper around the C function gdk_device_set_mode.
func (recv *Device) SetMode(mode InputMode) bool {
	c_mode := (C.GdkInputMode)(mode)

	retC := C.gdk_device_set_mode((*C.GdkDevice)(recv.native), c_mode)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_device_ungrab : no return generator

// Unsupported : gdk_device_warp : no return generator

// DeviceManager is a wrapper around the C record GdkDeviceManager.
type DeviceManager struct {
	native *C.GdkDeviceManager
}

func DeviceManagerNewFromC(u unsafe.Pointer) *DeviceManager {
	c := (*C.GdkDeviceManager)(u)
	if c == nil {
		return nil
	}

	g := &DeviceManager{native: c}

	return g
}

func (recv *DeviceManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

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

// Display is a wrapper around the C record GdkDisplay.
type Display struct {
	native *C.GdkDisplay
}

func DisplayNewFromC(u unsafe.Pointer) *Display {
	c := (*C.GdkDisplay)(u)
	if c == nil {
		return nil
	}

	g := &Display{native: c}

	return g
}

func (recv *Display) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_display_beep : no return generator

// Unsupported : gdk_display_close : no return generator

// DeviceIsGrabbed is a wrapper around the C function gdk_display_device_is_grabbed.
func (recv *Display) DeviceIsGrabbed(device *Device) bool {
	c_device := (*C.GdkDevice)(device.ToC())

	retC := C.gdk_display_device_is_grabbed((*C.GdkDisplay)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_display_flush : no return generator

// Unsupported : gdk_display_get_event : no return generator

// Unsupported : gdk_display_get_maximal_cursor_size : unsupported parameter width : no type generator for guint, guint*

// Unsupported : gdk_display_get_pointer : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_display_get_window_at_pointer : unsupported parameter win_x : no type generator for gint, gint*

// Unsupported : gdk_display_keyboard_ungrab : no return generator

// Unsupported : gdk_display_notify_startup_complete : no return generator

// Unsupported : gdk_display_peek_event : no return generator

// Unsupported : gdk_display_pointer_ungrab : no return generator

// Unsupported : gdk_display_put_event : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_display_request_selection_notification : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_display_set_double_click_distance : no return generator

// Unsupported : gdk_display_set_double_click_time : no return generator

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets : no param type

// Unsupported : gdk_display_sync : no return generator

// Unsupported : gdk_display_warp_pointer : no return generator

// DisplayManager is a wrapper around the C record GdkDisplayManager.
type DisplayManager struct {
	native *C.GdkDisplayManager
}

func DisplayManagerNewFromC(u unsafe.Pointer) *DisplayManager {
	c := (*C.GdkDisplayManager)(u)
	if c == nil {
		return nil
	}

	g := &DisplayManager{native: c}

	return g
}

func (recv *DisplayManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_display_manager_set_default_display : no return generator

// DragContext is a wrapper around the C record GdkDragContext.
type DragContext struct {
	native *C.GdkDragContext
}

func DragContextNewFromC(u unsafe.Pointer) *DragContext {
	c := (*C.GdkDragContext)(u)
	if c == nil {
		return nil
	}

	g := &DragContext{native: c}

	return g
}

func (recv *DragContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetDevice is a wrapper around the C function gdk_drag_context_get_device.
func (recv *DragContext) GetDevice() *Device {
	retC := C.gdk_drag_context_get_device((*C.GdkDragContext)(recv.native))
	retGo := DeviceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_drag_context_set_device : no return generator

// Unsupported : gdk_drag_context_set_hotspot : no return generator

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

// FrameClock is a wrapper around the C record GdkFrameClock.
type FrameClock struct {
	native *C.GdkFrameClock
}

func FrameClockNewFromC(u unsafe.Pointer) *FrameClock {
	c := (*C.GdkFrameClock)(u)
	if c == nil {
		return nil
	}

	g := &FrameClock{native: c}

	return g
}

func (recv *FrameClock) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_frame_clock_begin_updating : no return generator

// Unsupported : gdk_frame_clock_end_updating : no return generator

// Unsupported : gdk_frame_clock_get_refresh_info : unsupported parameter refresh_interval_return : no type generator for gint64, gint64*

// Unsupported : gdk_frame_clock_request_phase : no return generator

// GLContext is a wrapper around the C record GdkGLContext.
type GLContext struct {
	native *C.GdkGLContext
}

func GLContextNewFromC(u unsafe.Pointer) *GLContext {
	c := (*C.GdkGLContext)(u)
	if c == nil {
		return nil
	}

	g := &GLContext{native: c}

	return g
}

func (recv *GLContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_gl_context_get_required_version : unsupported parameter major : no type generator for gint, int*

// Unsupported : gdk_gl_context_get_version : unsupported parameter major : no type generator for gint, int*

// Unsupported : gdk_gl_context_make_current : no return generator

// Unsupported : gdk_gl_context_set_debug_enabled : no return generator

// Unsupported : gdk_gl_context_set_forward_compatible : no return generator

// Unsupported : gdk_gl_context_set_required_version : no return generator

// Unsupported : gdk_gl_context_set_use_es : no return generator

// Keymap is a wrapper around the C record GdkKeymap.
type Keymap struct {
	native *C.GdkKeymap
}

func KeymapNewFromC(u unsafe.Pointer) *Keymap {
	c := (*C.GdkKeymap)(u)
	if c == nil {
		return nil
	}

	g := &Keymap{native: c}

	return g
}

func (recv *Keymap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_keymap_add_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// GetDirection is a wrapper around the C function gdk_keymap_get_direction.
func (recv *Keymap) GetDirection() pango.Direction {
	retC := C.gdk_keymap_get_direction((*C.GdkKeymap)(recv.native))
	retGo := (pango.Direction)(retC)

	return retGo
}

// Unsupported : gdk_keymap_get_entries_for_keycode : unsupported parameter keys : no param type

// Unsupported : gdk_keymap_get_entries_for_keyval : unsupported parameter keys : no param type

// LookupKey is a wrapper around the C function gdk_keymap_lookup_key.
func (recv *Keymap) LookupKey(key *KeymapKey) uint32 {
	c_key := (*C.GdkKeymapKey)(key.ToC())

	retC := C.gdk_keymap_lookup_key((*C.GdkKeymap)(recv.native), c_key)
	retGo := (uint32)(retC)

	return retGo
}

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

// Unsupported : gdk_monitor_get_geometry : no return generator

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

// Unsupported : gdk_monitor_get_workarea : no return generator

// Screen is a wrapper around the C record GdkScreen.
type Screen struct {
	native *C.GdkScreen
}

func ScreenNewFromC(u unsafe.Pointer) *Screen {
	c := (*C.GdkScreen)(u)
	if c == nil {
		return nil
	}

	g := &Screen{native: c}

	return g
}

func (recv *Screen) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_screen_get_monitor_geometry : no return generator

// Unsupported : gdk_screen_get_monitor_workarea : no return generator

// Unsupported : gdk_screen_set_font_options : no return generator

// Unsupported : gdk_screen_set_resolution : no return generator

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

// Unsupported : gdk_seat_ungrab : no return generator

// Visual is a wrapper around the C record GdkVisual.
type Visual struct {
	native *C.GdkVisual
}

func VisualNewFromC(u unsafe.Pointer) *Visual {
	c := (*C.GdkVisual)(u)
	if c == nil {
		return nil
	}

	g := &Visual{native: c}

	return g
}

func (recv *Visual) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gdk_visual_get_blue_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_green_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_red_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Window is a wrapper around the C record GdkWindow.
type Window struct {
	native *C.GdkWindow
}

func WindowNewFromC(u unsafe.Pointer) *Window {
	c := (*C.GdkWindow)(u)
	if c == nil {
		return nil
	}

	g := &Window{native: c}

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowNew is a wrapper around the C function gdk_window_new.
func WindowNew(parent *Window, attributes *WindowAttr, attributesMask int32) *Window {
	c_parent := (*C.GdkWindow)(parent.ToC())

	c_attributes := (*C.GdkWindowAttr)(attributes.ToC())

	c_attributes_mask := (C.gint)(attributesMask)

	retC := C.gdk_window_new(c_parent, c_attributes, c_attributes_mask)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_beep : no return generator

// Unsupported : gdk_window_begin_move_drag : no return generator

// Unsupported : gdk_window_begin_move_drag_for_device : no return generator

// Unsupported : gdk_window_begin_paint_rect : no return generator

// Unsupported : gdk_window_begin_paint_region : no return generator

// Unsupported : gdk_window_begin_resize_drag : no return generator

// Unsupported : gdk_window_begin_resize_drag_for_device : no return generator

// Unsupported : gdk_window_configure_finished : no return generator

// Unsupported : gdk_window_coords_from_parent : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_coords_to_parent : unsupported parameter parent_x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_create_similar_image_surface : unsupported parameter format : no type generator for gint, cairo_format_t

// Unsupported : gdk_window_deiconify : no return generator

// Unsupported : gdk_window_destroy : no return generator

// Unsupported : gdk_window_destroy_notify : no return generator

// Unsupported : gdk_window_enable_synchronized_configure : no return generator

// Unsupported : gdk_window_end_draw_frame : no return generator

// Unsupported : gdk_window_end_paint : no return generator

// Unsupported : gdk_window_flush : no return generator

// Unsupported : gdk_window_focus : no return generator

// Unsupported : gdk_window_freeze_toplevel_updates_libgtk_only : no return generator

// Unsupported : gdk_window_freeze_updates : no return generator

// Unsupported : gdk_window_fullscreen : no return generator

// Unsupported : gdk_window_fullscreen_on_monitor : no return generator

// Unsupported : gdk_window_geometry_changed : no return generator

// GetChildren is a wrapper around the C function gdk_window_get_children.
func (recv *Window) GetChildren() *glib.List {
	retC := C.gdk_window_get_children((*C.GdkWindow)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetClipRegion is a wrapper around the C function gdk_window_get_clip_region.
func (recv *Window) GetClipRegion() *cairo.Region {
	retC := C.gdk_window_get_clip_region((*C.GdkWindow)(recv.native))
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_get_decorations : unsupported parameter decorations : GdkWMDecoration* with indirection level of 1

// Unsupported : gdk_window_get_device_position : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_device_position_double : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_get_drag_protocol : unsupported parameter target : record with indirection level of 2

// GetEvents is a wrapper around the C function gdk_window_get_events.
func (recv *Window) GetEvents() EventMask {
	retC := C.gdk_window_get_events((*C.GdkWindow)(recv.native))
	retGo := (EventMask)(retC)

	return retGo
}

// Unsupported : gdk_window_get_frame_extents : no return generator

// Unsupported : gdk_window_get_geometry : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_origin : unsupported parameter x : no type generator for gint, gint*

// GetParent is a wrapper around the C function gdk_window_get_parent.
func (recv *Window) GetParent() *Window {
	retC := C.gdk_window_get_parent((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_get_pointer : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_position : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_root_coords : unsupported parameter root_x : no type generator for gint, gint*

// Unsupported : gdk_window_get_root_origin : unsupported parameter x : no type generator for gint, gint*

// GetSourceEvents is a wrapper around the C function gdk_window_get_source_events.
func (recv *Window) GetSourceEvents(source InputSource) EventMask {
	c_source := (C.GdkInputSource)(source)

	retC := C.gdk_window_get_source_events((*C.GdkWindow)(recv.native), c_source)
	retGo := (EventMask)(retC)

	return retGo
}

// GetState is a wrapper around the C function gdk_window_get_state.
func (recv *Window) GetState() WindowState {
	retC := C.gdk_window_get_state((*C.GdkWindow)(recv.native))
	retGo := (WindowState)(retC)

	return retGo
}

// GetToplevel is a wrapper around the C function gdk_window_get_toplevel.
func (recv *Window) GetToplevel() *Window {
	retC := C.gdk_window_get_toplevel((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUpdateArea is a wrapper around the C function gdk_window_get_update_area.
func (recv *Window) GetUpdateArea() *cairo.Region {
	retC := C.gdk_window_get_update_area((*C.GdkWindow)(recv.native))
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_get_user_data : unsupported parameter data : no type generator for gpointer, gpointer*

// GetVisibleRegion is a wrapper around the C function gdk_window_get_visible_region.
func (recv *Window) GetVisibleRegion() *cairo.Region {
	retC := C.gdk_window_get_visible_region((*C.GdkWindow)(recv.native))
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWindowType is a wrapper around the C function gdk_window_get_window_type.
func (recv *Window) GetWindowType() WindowType {
	retC := C.gdk_window_get_window_type((*C.GdkWindow)(recv.native))
	retGo := (WindowType)(retC)

	return retGo
}

// Unsupported : gdk_window_hide : no return generator

// Unsupported : gdk_window_iconify : no return generator

// Unsupported : gdk_window_input_shape_combine_region : no return generator

// Unsupported : gdk_window_invalidate_maybe_recurse : unsupported parameter child_func : no type generator for WindowChildFunc, GdkWindowChildFunc

// Unsupported : gdk_window_invalidate_rect : no return generator

// Unsupported : gdk_window_invalidate_region : no return generator

// IsViewable is a wrapper around the C function gdk_window_is_viewable.
func (recv *Window) IsViewable() bool {
	retC := C.gdk_window_is_viewable((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsVisible is a wrapper around the C function gdk_window_is_visible.
func (recv *Window) IsVisible() bool {
	retC := C.gdk_window_is_visible((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_window_lower : no return generator

// Unsupported : gdk_window_mark_paint_from_clip : no return generator

// Unsupported : gdk_window_maximize : no return generator

// Unsupported : gdk_window_merge_child_input_shapes : no return generator

// Unsupported : gdk_window_merge_child_shapes : no return generator

// Unsupported : gdk_window_move : no return generator

// Unsupported : gdk_window_move_region : no return generator

// Unsupported : gdk_window_move_resize : no return generator

// PeekChildren is a wrapper around the C function gdk_window_peek_children.
func (recv *Window) PeekChildren() *glib.List {
	retC := C.gdk_window_peek_children((*C.GdkWindow)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_process_updates : no return generator

// Unsupported : gdk_window_raise : no return generator

// Unsupported : gdk_window_register_dnd : no return generator

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_reparent : no return generator

// Unsupported : gdk_window_resize : no return generator

// Unsupported : gdk_window_restack : no return generator

// Unsupported : gdk_window_scroll : no return generator

// Unsupported : gdk_window_set_accept_focus : no return generator

// Unsupported : gdk_window_set_background : no return generator

// Unsupported : gdk_window_set_background_pattern : no return generator

// Unsupported : gdk_window_set_background_rgba : no return generator

// Unsupported : gdk_window_set_child_input_shapes : no return generator

// Unsupported : gdk_window_set_child_shapes : no return generator

// Unsupported : gdk_window_set_composited : no return generator

// Unsupported : gdk_window_set_cursor : no return generator

// Unsupported : gdk_window_set_decorations : no return generator

// Unsupported : gdk_window_set_device_cursor : no return generator

// Unsupported : gdk_window_set_device_events : no return generator

// Unsupported : gdk_window_set_event_compression : no return generator

// Unsupported : gdk_window_set_events : no return generator

// Unsupported : gdk_window_set_focus_on_map : no return generator

// Unsupported : gdk_window_set_fullscreen_mode : no return generator

// Unsupported : gdk_window_set_functions : no return generator

// Unsupported : gdk_window_set_geometry_hints : no return generator

// Unsupported : gdk_window_set_group : no return generator

// Unsupported : gdk_window_set_icon_list : no return generator

// Unsupported : gdk_window_set_icon_name : no return generator

// Unsupported : gdk_window_set_invalidate_handler : unsupported parameter handler : no type generator for WindowInvalidateHandlerFunc, GdkWindowInvalidateHandlerFunc

// Unsupported : gdk_window_set_keep_above : no return generator

// Unsupported : gdk_window_set_keep_below : no return generator

// Unsupported : gdk_window_set_modal_hint : no return generator

// Unsupported : gdk_window_set_opacity : no return generator

// Unsupported : gdk_window_set_opaque_region : no return generator

// Unsupported : gdk_window_set_override_redirect : no return generator

// Unsupported : gdk_window_set_pass_through : no return generator

// Unsupported : gdk_window_set_role : no return generator

// Unsupported : gdk_window_set_shadow_width : no return generator

// Unsupported : gdk_window_set_skip_pager_hint : no return generator

// Unsupported : gdk_window_set_skip_taskbar_hint : no return generator

// Unsupported : gdk_window_set_source_events : no return generator

// Unsupported : gdk_window_set_startup_id : no return generator

// SetStaticGravities is a wrapper around the C function gdk_window_set_static_gravities.
func (recv *Window) SetStaticGravities(useStatic bool) bool {
	c_use_static :=
		boolToGboolean(useStatic)

	retC := C.gdk_window_set_static_gravities((*C.GdkWindow)(recv.native), c_use_static)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_window_set_support_multidevice : no return generator

// Unsupported : gdk_window_set_title : no return generator

// Unsupported : gdk_window_set_transient_for : no return generator

// Unsupported : gdk_window_set_type_hint : no return generator

// Unsupported : gdk_window_set_urgency_hint : no return generator

// Unsupported : gdk_window_set_user_data : no return generator

// Unsupported : gdk_window_shape_combine_region : no return generator

// Unsupported : gdk_window_show : no return generator

// Unsupported : gdk_window_show_unraised : no return generator

// Unsupported : gdk_window_show_window_menu : unsupported parameter event : no type generator for Event, GdkEvent*

// Unsupported : gdk_window_stick : no return generator

// Unsupported : gdk_window_thaw_toplevel_updates_libgtk_only : no return generator

// Unsupported : gdk_window_thaw_updates : no return generator

// Unsupported : gdk_window_unfullscreen : no return generator

// Unsupported : gdk_window_unmaximize : no return generator

// Unsupported : gdk_window_unstick : no return generator

// Unsupported : gdk_window_withdraw : no return generator
