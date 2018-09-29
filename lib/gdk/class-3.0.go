// This is a generated file - DO NOT EDIT
// +build gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_app_launch_context_set_icon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gdk_cursor_get_surface : unsupported parameter x_hot : no type generator for gdouble, gdouble*

// GetAssociatedDevice is a wrapper around the C function gdk_device_get_associated_device.
func (recv *Device) GetAssociatedDevice() *Device {
	retC := C.gdk_device_get_associated_device((*C.GdkDevice)(recv.native))
	retGo := DeviceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_device_get_axis : unsupported parameter axes : no param type

// Unsupported : gdk_device_get_axis_value : unsupported parameter axes : no param type

// GetDeviceType is a wrapper around the C function gdk_device_get_device_type.
func (recv *Device) GetDeviceType() DeviceType {
	retC := C.gdk_device_get_device_type((*C.GdkDevice)(recv.native))
	retGo := (DeviceType)(retC)

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_device_get_display.
func (recv *Device) GetDisplay() *Display {
	retC := C.gdk_device_get_display((*C.GdkDevice)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_device_get_history : unsupported parameter events : no param type

// Unsupported : gdk_device_get_key : unsupported parameter keyval : no type generator for guint, guint*

// GetNAxes is a wrapper around the C function gdk_device_get_n_axes.
func (recv *Device) GetNAxes() int32 {
	retC := C.gdk_device_get_n_axes((*C.GdkDevice)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_device_get_position : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_device_get_position_double : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_device_get_state : unsupported parameter axes : no param type

// Unsupported : gdk_device_get_window_at_position : unsupported parameter win_x : no type generator for gint, gint*

// Unsupported : gdk_device_get_window_at_position_double : unsupported parameter win_x : no type generator for gdouble, gdouble*

// Grab is a wrapper around the C function gdk_device_grab.
func (recv *Device) Grab(window *Window, grabOwnership GrabOwnership, ownerEvents bool, eventMask EventMask, cursor *Cursor, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(window.ToC())

	c_grab_ownership := (C.GdkGrabOwnership)(grabOwnership)

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_event_mask := (C.GdkEventMask)(eventMask)

	c_cursor := (*C.GdkCursor)(cursor.ToC())

	c_time_ := (C.guint32)(time)

	retC := C.gdk_device_grab((*C.GdkDevice)(recv.native), c_window, c_grab_ownership, c_owner_events, c_event_mask, c_cursor, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// ListAxes is a wrapper around the C function gdk_device_list_axes.
func (recv *Device) ListAxes() *glib.List {
	retC := C.gdk_device_list_axes((*C.GdkDevice)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ungrab is a wrapper around the C function gdk_device_ungrab.
func (recv *Device) Ungrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_device_ungrab((*C.GdkDevice)(recv.native), c_time_)

	return
}

// Warp is a wrapper around the C function gdk_device_warp.
func (recv *Device) Warp(screen *Screen, x int32, y int32) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_device_warp((*C.GdkDevice)(recv.native), c_screen, c_x, c_y)

	return
}

// GetClientPointer is a wrapper around the C function gdk_device_manager_get_client_pointer.
func (recv *DeviceManager) GetClientPointer() *Device {
	retC := C.gdk_device_manager_get_client_pointer((*C.GdkDeviceManager)(recv.native))
	retGo := DeviceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_device_manager_get_display.
func (recv *DeviceManager) GetDisplay() *Display {
	retC := C.gdk_device_manager_get_display((*C.GdkDeviceManager)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListDevices is a wrapper around the C function gdk_device_manager_list_devices.
func (recv *DeviceManager) ListDevices(type_ DeviceType) *glib.List {
	c_type := (C.GdkDeviceType)(type_)

	retC := C.gdk_device_manager_list_devices((*C.GdkDeviceManager)(recv.native), c_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAppLaunchContext is a wrapper around the C function gdk_display_get_app_launch_context.
func (recv *Display) GetAppLaunchContext() *AppLaunchContext {
	retC := C.gdk_display_get_app_launch_context((*C.GdkDisplay)(recv.native))
	retGo := AppLaunchContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDeviceManager is a wrapper around the C function gdk_display_get_device_manager.
func (recv *Display) GetDeviceManager() *DeviceManager {
	retC := C.gdk_display_get_device_manager((*C.GdkDisplay)(recv.native))
	retGo := DeviceManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_display_get_event : no return generator

// Unsupported : gdk_display_get_maximal_cursor_size : unsupported parameter width : no type generator for guint, guint*

// Unsupported : gdk_display_get_pointer : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_display_get_window_at_pointer : unsupported parameter win_x : no type generator for gint, gint*

// HasPending is a wrapper around the C function gdk_display_has_pending.
func (recv *Display) HasPending() bool {
	retC := C.gdk_display_has_pending((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// NotifyStartupComplete is a wrapper around the C function gdk_display_notify_startup_complete.
func (recv *Display) NotifyStartupComplete(startupId string) {
	c_startup_id := C.CString(startupId)
	defer C.free(unsafe.Pointer(c_startup_id))

	C.gdk_display_notify_startup_complete((*C.GdkDisplay)(recv.native), c_startup_id)

	return
}

// Unsupported : gdk_display_peek_event : no return generator

// Unsupported : gdk_display_put_event : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_display_request_selection_notification : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets : no param type

// OpenDisplay is a wrapper around the C function gdk_display_manager_open_display.
func (recv *DisplayManager) OpenDisplay(name string) *Display {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gdk_display_manager_open_display((*C.GdkDisplayManager)(recv.native), c_name)
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDestWindow is a wrapper around the C function gdk_drag_context_get_dest_window.
func (recv *DragContext) GetDestWindow() *Window {
	retC := C.gdk_drag_context_get_dest_window((*C.GdkDragContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetProtocol is a wrapper around the C function gdk_drag_context_get_protocol.
func (recv *DragContext) GetProtocol() DragProtocol {
	retC := C.gdk_drag_context_get_protocol((*C.GdkDragContext)(recv.native))
	retGo := (DragProtocol)(retC)

	return retGo
}

// Unsupported : gdk_frame_clock_get_refresh_info : unsupported parameter refresh_interval_return : no type generator for gint64, gint64*

// Unsupported : gdk_gl_context_get_required_version : unsupported parameter major : no type generator for gint, int*

// Unsupported : gdk_gl_context_get_version : unsupported parameter major : no type generator for gint, int*

// Unsupported : gdk_keymap_add_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gdk_keymap_get_entries_for_keycode : unsupported parameter keys : no param type

// Unsupported : gdk_keymap_get_entries_for_keyval : unsupported parameter keys : no param type

// GetNumLockState is a wrapper around the C function gdk_keymap_get_num_lock_state.
func (recv *Keymap) GetNumLockState() bool {
	retC := C.gdk_keymap_get_num_lock_state((*C.GdkKeymap)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_keymap_map_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gdk_keymap_translate_keyboard_state : unsupported parameter keyval : no type generator for guint, guint*

// Unsupported : gdk_monitor_get_geometry : unsupported parameter geometry : Blacklisted record : GdkRectangle

// Unsupported : gdk_monitor_get_workarea : unsupported parameter workarea : Blacklisted record : GdkRectangle

// Unsupported : gdk_screen_get_monitor_geometry : unsupported parameter dest : Blacklisted record : GdkRectangle

// Unsupported : gdk_screen_get_monitor_workarea : unsupported parameter dest : Blacklisted record : GdkRectangle

// Unsupported : gdk_seat_grab : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_visual_get_blue_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_green_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_red_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_begin_paint_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_coords_from_parent : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_coords_to_parent : unsupported parameter parent_x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_create_similar_image_surface : unsupported parameter format : no type generator for gint, cairo_format_t

// Unsupported : gdk_window_get_decorations : unsupported parameter decorations : GdkWMDecoration* with indirection level of 1

// GetDeviceCursor is a wrapper around the C function gdk_window_get_device_cursor.
func (recv *Window) GetDeviceCursor(device *Device) *Cursor {
	c_device := (*C.GdkDevice)(device.ToC())

	retC := C.gdk_window_get_device_cursor((*C.GdkWindow)(recv.native), c_device)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDeviceEvents is a wrapper around the C function gdk_window_get_device_events.
func (recv *Window) GetDeviceEvents(device *Device) EventMask {
	c_device := (*C.GdkDevice)(device.ToC())

	retC := C.gdk_window_get_device_events((*C.GdkWindow)(recv.native), c_device)
	retGo := (EventMask)(retC)

	return retGo
}

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

// GetSupportMultidevice is a wrapper around the C function gdk_window_get_support_multidevice.
func (recv *Window) GetSupportMultidevice() bool {
	retC := C.gdk_window_get_support_multidevice((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_window_get_user_data : unsupported parameter data : no type generator for gpointer, gpointer*

// Unsupported : gdk_window_invalidate_maybe_recurse : unsupported parameter child_func : no type generator for WindowChildFunc, GdkWindowChildFunc

// Unsupported : gdk_window_invalidate_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// SetDeviceCursor is a wrapper around the C function gdk_window_set_device_cursor.
func (recv *Window) SetDeviceCursor(device *Device, cursor *Cursor) {
	c_device := (*C.GdkDevice)(device.ToC())

	c_cursor := (*C.GdkCursor)(cursor.ToC())

	C.gdk_window_set_device_cursor((*C.GdkWindow)(recv.native), c_device, c_cursor)

	return
}

// SetDeviceEvents is a wrapper around the C function gdk_window_set_device_events.
func (recv *Window) SetDeviceEvents(device *Device, eventMask EventMask) {
	c_device := (*C.GdkDevice)(device.ToC())

	c_event_mask := (C.GdkEventMask)(eventMask)

	C.gdk_window_set_device_events((*C.GdkWindow)(recv.native), c_device, c_event_mask)

	return
}

// Unsupported : gdk_window_set_invalidate_handler : unsupported parameter handler : no type generator for WindowInvalidateHandlerFunc, GdkWindowInvalidateHandlerFunc

// SetSourceEvents is a wrapper around the C function gdk_window_set_source_events.
func (recv *Window) SetSourceEvents(source InputSource, eventMask EventMask) {
	c_source := (C.GdkInputSource)(source)

	c_event_mask := (C.GdkEventMask)(eventMask)

	C.gdk_window_set_source_events((*C.GdkWindow)(recv.native), c_source, c_event_mask)

	return
}

// SetSupportMultidevice is a wrapper around the C function gdk_window_set_support_multidevice.
func (recv *Window) SetSupportMultidevice(supportMultidevice bool) {
	c_support_multidevice :=
		boolToGboolean(supportMultidevice)

	C.gdk_window_set_support_multidevice((*C.GdkWindow)(recv.native), c_support_multidevice)

	return
}

// Unsupported : gdk_window_show_window_menu : unsupported parameter event : no type generator for Event, GdkEvent*
