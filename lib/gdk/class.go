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

// Unsupported : gdk_app_launch_context_set_icon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

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

// Unref is a wrapper around the C function gdk_cursor_unref.
func (recv *Cursor) Unref() {
	C.gdk_cursor_unref((*C.GdkCursor)(recv.native))

	return
}

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

// SetAxisUse is a wrapper around the C function gdk_device_set_axis_use.
func (recv *Device) SetAxisUse(index uint32, use AxisUse) {
	c_index_ := (C.guint)(index)

	c_use := (C.GdkAxisUse)(use)

	C.gdk_device_set_axis_use((*C.GdkDevice)(recv.native), c_index_, c_use)

	return
}

// SetKey is a wrapper around the C function gdk_device_set_key.
func (recv *Device) SetKey(index uint32, keyval uint32, modifiers ModifierType) {
	c_index_ := (C.guint)(index)

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gdk_device_set_key((*C.GdkDevice)(recv.native), c_index_, c_keyval, c_modifiers)

	return
}

// SetMode is a wrapper around the C function gdk_device_set_mode.
func (recv *Device) SetMode(mode InputMode) bool {
	c_mode := (C.GdkInputMode)(mode)

	retC := C.gdk_device_set_mode((*C.GdkDevice)(recv.native), c_mode)
	retGo := retC == C.TRUE

	return retGo
}

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

// DeviceIsGrabbed is a wrapper around the C function gdk_display_device_is_grabbed.
func (recv *Display) DeviceIsGrabbed(device *Device) bool {
	c_device := (*C.GdkDevice)(device.ToC())

	retC := C.gdk_display_device_is_grabbed((*C.GdkDisplay)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_display_get_event : no return generator

// Unsupported : gdk_display_get_maximal_cursor_size : unsupported parameter width : no type generator for guint, guint*

// Unsupported : gdk_display_get_pointer : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_display_get_window_at_pointer : unsupported parameter win_x : no type generator for gint, gint*

// Unsupported : gdk_display_peek_event : no return generator

// Unsupported : gdk_display_put_event : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_display_request_selection_notification : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets : no param type

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

// SetDevice is a wrapper around the C function gdk_drag_context_set_device.
func (recv *DragContext) SetDevice(device *Device) {
	c_device := (*C.GdkDevice)(device.ToC())

	C.gdk_drag_context_set_device((*C.GdkDragContext)(recv.native), c_device)

	return
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

// Unsupported : gdk_frame_clock_get_refresh_info : unsupported parameter refresh_interval_return : no type generator for gint64, gint64*

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

// Unsupported : gdk_screen_get_monitor_geometry : unsupported parameter dest : Blacklisted record : GdkRectangle

// Unsupported : gdk_screen_get_monitor_workarea : unsupported parameter dest : Blacklisted record : GdkRectangle

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

// BeginMoveDrag is a wrapper around the C function gdk_window_begin_move_drag.
func (recv *Window) BeginMoveDrag(button int32, rootX int32, rootY int32, timestamp uint32) {
	c_button := (C.gint)(button)

	c_root_x := (C.gint)(rootX)

	c_root_y := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_begin_move_drag((*C.GdkWindow)(recv.native), c_button, c_root_x, c_root_y, c_timestamp)

	return
}

// Unsupported : gdk_window_begin_paint_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// BeginPaintRegion is a wrapper around the C function gdk_window_begin_paint_region.
func (recv *Window) BeginPaintRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(region.ToC())

	C.gdk_window_begin_paint_region((*C.GdkWindow)(recv.native), c_region)

	return
}

// BeginResizeDrag is a wrapper around the C function gdk_window_begin_resize_drag.
func (recv *Window) BeginResizeDrag(edge WindowEdge, button int32, rootX int32, rootY int32, timestamp uint32) {
	c_edge := (C.GdkWindowEdge)(edge)

	c_button := (C.gint)(button)

	c_root_x := (C.gint)(rootX)

	c_root_y := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_begin_resize_drag((*C.GdkWindow)(recv.native), c_edge, c_button, c_root_x, c_root_y, c_timestamp)

	return
}

// Unsupported : gdk_window_coords_from_parent : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_coords_to_parent : unsupported parameter parent_x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_create_similar_image_surface : unsupported parameter format : no type generator for gint, cairo_format_t

// Deiconify is a wrapper around the C function gdk_window_deiconify.
func (recv *Window) Deiconify() {
	C.gdk_window_deiconify((*C.GdkWindow)(recv.native))

	return
}

// Destroy is a wrapper around the C function gdk_window_destroy.
func (recv *Window) Destroy() {
	C.gdk_window_destroy((*C.GdkWindow)(recv.native))

	return
}

// DestroyNotify is a wrapper around the C function gdk_window_destroy_notify.
func (recv *Window) DestroyNotify() {
	C.gdk_window_destroy_notify((*C.GdkWindow)(recv.native))

	return
}

// EndPaint is a wrapper around the C function gdk_window_end_paint.
func (recv *Window) EndPaint() {
	C.gdk_window_end_paint((*C.GdkWindow)(recv.native))

	return
}

// Focus is a wrapper around the C function gdk_window_focus.
func (recv *Window) Focus(timestamp uint32) {
	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_focus((*C.GdkWindow)(recv.native), c_timestamp)

	return
}

// FreezeToplevelUpdatesLibgtkOnly is a wrapper around the C function gdk_window_freeze_toplevel_updates_libgtk_only.
func (recv *Window) FreezeToplevelUpdatesLibgtkOnly() {
	C.gdk_window_freeze_toplevel_updates_libgtk_only((*C.GdkWindow)(recv.native))

	return
}

// FreezeUpdates is a wrapper around the C function gdk_window_freeze_updates.
func (recv *Window) FreezeUpdates() {
	C.gdk_window_freeze_updates((*C.GdkWindow)(recv.native))

	return
}

// FullscreenOnMonitor is a wrapper around the C function gdk_window_fullscreen_on_monitor.
func (recv *Window) FullscreenOnMonitor(monitor int32) {
	c_monitor := (C.gint)(monitor)

	C.gdk_window_fullscreen_on_monitor((*C.GdkWindow)(recv.native), c_monitor)

	return
}

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

// Unsupported : gdk_window_get_frame_extents : unsupported parameter rect : Blacklisted record : GdkRectangle

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

// Hide is a wrapper around the C function gdk_window_hide.
func (recv *Window) Hide() {
	C.gdk_window_hide((*C.GdkWindow)(recv.native))

	return
}

// Iconify is a wrapper around the C function gdk_window_iconify.
func (recv *Window) Iconify() {
	C.gdk_window_iconify((*C.GdkWindow)(recv.native))

	return
}

// Unsupported : gdk_window_invalidate_maybe_recurse : unsupported parameter child_func : no type generator for WindowChildFunc, GdkWindowChildFunc

// Unsupported : gdk_window_invalidate_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// InvalidateRegion is a wrapper around the C function gdk_window_invalidate_region.
func (recv *Window) InvalidateRegion(region *cairo.Region, invalidateChildren bool) {
	c_region := (*C.cairo_region_t)(region.ToC())

	c_invalidate_children :=
		boolToGboolean(invalidateChildren)

	C.gdk_window_invalidate_region((*C.GdkWindow)(recv.native), c_region, c_invalidate_children)

	return
}

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

// Lower is a wrapper around the C function gdk_window_lower.
func (recv *Window) Lower() {
	C.gdk_window_lower((*C.GdkWindow)(recv.native))

	return
}

// Maximize is a wrapper around the C function gdk_window_maximize.
func (recv *Window) Maximize() {
	C.gdk_window_maximize((*C.GdkWindow)(recv.native))

	return
}

// MergeChildShapes is a wrapper around the C function gdk_window_merge_child_shapes.
func (recv *Window) MergeChildShapes() {
	C.gdk_window_merge_child_shapes((*C.GdkWindow)(recv.native))

	return
}

// Move is a wrapper around the C function gdk_window_move.
func (recv *Window) Move(x int32, y int32) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_window_move((*C.GdkWindow)(recv.native), c_x, c_y)

	return
}

// MoveResize is a wrapper around the C function gdk_window_move_resize.
func (recv *Window) MoveResize(x int32, y int32, width int32, height int32) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gdk_window_move_resize((*C.GdkWindow)(recv.native), c_x, c_y, c_width, c_height)

	return
}

// PeekChildren is a wrapper around the C function gdk_window_peek_children.
func (recv *Window) PeekChildren() *glib.List {
	retC := C.gdk_window_peek_children((*C.GdkWindow)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ProcessUpdates is a wrapper around the C function gdk_window_process_updates.
func (recv *Window) ProcessUpdates(updateChildren bool) {
	c_update_children :=
		boolToGboolean(updateChildren)

	C.gdk_window_process_updates((*C.GdkWindow)(recv.native), c_update_children)

	return
}

// Raise is a wrapper around the C function gdk_window_raise.
func (recv *Window) Raise() {
	C.gdk_window_raise((*C.GdkWindow)(recv.native))

	return
}

// RegisterDnd is a wrapper around the C function gdk_window_register_dnd.
func (recv *Window) RegisterDnd() {
	C.gdk_window_register_dnd((*C.GdkWindow)(recv.native))

	return
}

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Reparent is a wrapper around the C function gdk_window_reparent.
func (recv *Window) Reparent(newParent *Window, x int32, y int32) {
	c_new_parent := (*C.GdkWindow)(newParent.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_window_reparent((*C.GdkWindow)(recv.native), c_new_parent, c_x, c_y)

	return
}

// Resize is a wrapper around the C function gdk_window_resize.
func (recv *Window) Resize(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gdk_window_resize((*C.GdkWindow)(recv.native), c_width, c_height)

	return
}

// Scroll is a wrapper around the C function gdk_window_scroll.
func (recv *Window) Scroll(dx int32, dy int32) {
	c_dx := (C.gint)(dx)

	c_dy := (C.gint)(dy)

	C.gdk_window_scroll((*C.GdkWindow)(recv.native), c_dx, c_dy)

	return
}

// SetBackground is a wrapper around the C function gdk_window_set_background.
func (recv *Window) SetBackground(color *Color) {
	c_color := (*C.GdkColor)(color.ToC())

	C.gdk_window_set_background((*C.GdkWindow)(recv.native), c_color)

	return
}

// SetBackgroundPattern is a wrapper around the C function gdk_window_set_background_pattern.
func (recv *Window) SetBackgroundPattern(pattern *cairo.Pattern) {
	c_pattern := (*C.cairo_pattern_t)(pattern.ToC())

	C.gdk_window_set_background_pattern((*C.GdkWindow)(recv.native), c_pattern)

	return
}

// SetBackgroundRgba is a wrapper around the C function gdk_window_set_background_rgba.
func (recv *Window) SetBackgroundRgba(rgba *RGBA) {
	c_rgba := (*C.GdkRGBA)(rgba.ToC())

	C.gdk_window_set_background_rgba((*C.GdkWindow)(recv.native), c_rgba)

	return
}

// SetChildShapes is a wrapper around the C function gdk_window_set_child_shapes.
func (recv *Window) SetChildShapes() {
	C.gdk_window_set_child_shapes((*C.GdkWindow)(recv.native))

	return
}

// SetCursor is a wrapper around the C function gdk_window_set_cursor.
func (recv *Window) SetCursor(cursor *Cursor) {
	c_cursor := (*C.GdkCursor)(cursor.ToC())

	C.gdk_window_set_cursor((*C.GdkWindow)(recv.native), c_cursor)

	return
}

// SetDecorations is a wrapper around the C function gdk_window_set_decorations.
func (recv *Window) SetDecorations(decorations WMDecoration) {
	c_decorations := (C.GdkWMDecoration)(decorations)

	C.gdk_window_set_decorations((*C.GdkWindow)(recv.native), c_decorations)

	return
}

// SetEvents is a wrapper around the C function gdk_window_set_events.
func (recv *Window) SetEvents(eventMask EventMask) {
	c_event_mask := (C.GdkEventMask)(eventMask)

	C.gdk_window_set_events((*C.GdkWindow)(recv.native), c_event_mask)

	return
}

// SetFunctions is a wrapper around the C function gdk_window_set_functions.
func (recv *Window) SetFunctions(functions WMFunction) {
	c_functions := (C.GdkWMFunction)(functions)

	C.gdk_window_set_functions((*C.GdkWindow)(recv.native), c_functions)

	return
}

// SetGeometryHints is a wrapper around the C function gdk_window_set_geometry_hints.
func (recv *Window) SetGeometryHints(geometry *Geometry, geomMask WindowHints) {
	c_geometry := (*C.GdkGeometry)(geometry.ToC())

	c_geom_mask := (C.GdkWindowHints)(geomMask)

	C.gdk_window_set_geometry_hints((*C.GdkWindow)(recv.native), c_geometry, c_geom_mask)

	return
}

// SetGroup is a wrapper around the C function gdk_window_set_group.
func (recv *Window) SetGroup(leader *Window) {
	c_leader := (*C.GdkWindow)(leader.ToC())

	C.gdk_window_set_group((*C.GdkWindow)(recv.native), c_leader)

	return
}

// SetIconList is a wrapper around the C function gdk_window_set_icon_list.
func (recv *Window) SetIconList(pixbufs *glib.List) {
	c_pixbufs := (*C.GList)(pixbufs.ToC())

	C.gdk_window_set_icon_list((*C.GdkWindow)(recv.native), c_pixbufs)

	return
}

// SetIconName is a wrapper around the C function gdk_window_set_icon_name.
func (recv *Window) SetIconName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gdk_window_set_icon_name((*C.GdkWindow)(recv.native), c_name)

	return
}

// Unsupported : gdk_window_set_invalidate_handler : unsupported parameter handler : no type generator for WindowInvalidateHandlerFunc, GdkWindowInvalidateHandlerFunc

// SetModalHint is a wrapper around the C function gdk_window_set_modal_hint.
func (recv *Window) SetModalHint(modal bool) {
	c_modal :=
		boolToGboolean(modal)

	C.gdk_window_set_modal_hint((*C.GdkWindow)(recv.native), c_modal)

	return
}

// SetOverrideRedirect is a wrapper around the C function gdk_window_set_override_redirect.
func (recv *Window) SetOverrideRedirect(overrideRedirect bool) {
	c_override_redirect :=
		boolToGboolean(overrideRedirect)

	C.gdk_window_set_override_redirect((*C.GdkWindow)(recv.native), c_override_redirect)

	return
}

// SetRole is a wrapper around the C function gdk_window_set_role.
func (recv *Window) SetRole(role string) {
	c_role := C.CString(role)
	defer C.free(unsafe.Pointer(c_role))

	C.gdk_window_set_role((*C.GdkWindow)(recv.native), c_role)

	return
}

// SetStaticGravities is a wrapper around the C function gdk_window_set_static_gravities.
func (recv *Window) SetStaticGravities(useStatic bool) bool {
	c_use_static :=
		boolToGboolean(useStatic)

	retC := C.gdk_window_set_static_gravities((*C.GdkWindow)(recv.native), c_use_static)
	retGo := retC == C.TRUE

	return retGo
}

// SetTitle is a wrapper around the C function gdk_window_set_title.
func (recv *Window) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gdk_window_set_title((*C.GdkWindow)(recv.native), c_title)

	return
}

// SetTransientFor is a wrapper around the C function gdk_window_set_transient_for.
func (recv *Window) SetTransientFor(parent *Window) {
	c_parent := (*C.GdkWindow)(parent.ToC())

	C.gdk_window_set_transient_for((*C.GdkWindow)(recv.native), c_parent)

	return
}

// SetTypeHint is a wrapper around the C function gdk_window_set_type_hint.
func (recv *Window) SetTypeHint(hint WindowTypeHint) {
	c_hint := (C.GdkWindowTypeHint)(hint)

	C.gdk_window_set_type_hint((*C.GdkWindow)(recv.native), c_hint)

	return
}

// SetUserData is a wrapper around the C function gdk_window_set_user_data.
func (recv *Window) SetUserData(userData uintptr) {
	c_user_data := (C.gpointer)(userData)

	C.gdk_window_set_user_data((*C.GdkWindow)(recv.native), c_user_data)

	return
}

// ShapeCombineRegion is a wrapper around the C function gdk_window_shape_combine_region.
func (recv *Window) ShapeCombineRegion(shapeRegion *cairo.Region, offsetX int32, offsetY int32) {
	c_shape_region := (*C.cairo_region_t)(shapeRegion.ToC())

	c_offset_x := (C.gint)(offsetX)

	c_offset_y := (C.gint)(offsetY)

	C.gdk_window_shape_combine_region((*C.GdkWindow)(recv.native), c_shape_region, c_offset_x, c_offset_y)

	return
}

// Show is a wrapper around the C function gdk_window_show.
func (recv *Window) Show() {
	C.gdk_window_show((*C.GdkWindow)(recv.native))

	return
}

// ShowUnraised is a wrapper around the C function gdk_window_show_unraised.
func (recv *Window) ShowUnraised() {
	C.gdk_window_show_unraised((*C.GdkWindow)(recv.native))

	return
}

// Unsupported : gdk_window_show_window_menu : unsupported parameter event : no type generator for Event, GdkEvent*

// Stick is a wrapper around the C function gdk_window_stick.
func (recv *Window) Stick() {
	C.gdk_window_stick((*C.GdkWindow)(recv.native))

	return
}

// ThawToplevelUpdatesLibgtkOnly is a wrapper around the C function gdk_window_thaw_toplevel_updates_libgtk_only.
func (recv *Window) ThawToplevelUpdatesLibgtkOnly() {
	C.gdk_window_thaw_toplevel_updates_libgtk_only((*C.GdkWindow)(recv.native))

	return
}

// ThawUpdates is a wrapper around the C function gdk_window_thaw_updates.
func (recv *Window) ThawUpdates() {
	C.gdk_window_thaw_updates((*C.GdkWindow)(recv.native))

	return
}

// Unmaximize is a wrapper around the C function gdk_window_unmaximize.
func (recv *Window) Unmaximize() {
	C.gdk_window_unmaximize((*C.GdkWindow)(recv.native))

	return
}

// Unstick is a wrapper around the C function gdk_window_unstick.
func (recv *Window) Unstick() {
	C.gdk_window_unstick((*C.GdkWindow)(recv.native))

	return
}

// Withdraw is a wrapper around the C function gdk_window_withdraw.
func (recv *Window) Withdraw() {
	C.gdk_window_withdraw((*C.GdkWindow)(recv.native))

	return
}
