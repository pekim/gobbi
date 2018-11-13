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

// Returns the associated device to @device, if @device is of type
// %GDK_DEVICE_TYPE_MASTER, it will return the paired pointer or
// keyboard.
//
// If @device is of type %GDK_DEVICE_TYPE_SLAVE, it will return
// the master device to which @device is attached to.
//
// If @device is of type %GDK_DEVICE_TYPE_FLOATING, %NULL will be
// returned, as there is no associated device.
/*

C function : gdk_device_get_associated_device
*/
func (recv *Device) GetAssociatedDevice() *Device {
	retC := C.gdk_device_get_associated_device((*C.GdkDevice)(recv.native))
	var retGo (*Device)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DeviceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gdk_device_get_axis_value : unsupported parameter axis_label : Blacklisted record : GdkAtom

// Returns the device type for @device.
/*

C function : gdk_device_get_device_type
*/
func (recv *Device) GetDeviceType() DeviceType {
	retC := C.gdk_device_get_device_type((*C.GdkDevice)(recv.native))
	retGo := (DeviceType)(retC)

	return retGo
}

// Returns the #GdkDisplay to which @device pertains.
/*

C function : gdk_device_get_display
*/
func (recv *Device) GetDisplay() *Display {
	retC := C.gdk_device_get_display((*C.GdkDevice)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the number of axes the device currently has.
/*

C function : gdk_device_get_n_axes
*/
func (recv *Device) GetNAxes() int32 {
	retC := C.gdk_device_get_n_axes((*C.GdkDevice)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the current location of @device. As a slave device
// coordinates are those of its master pointer, This function
// may not be called on devices of type %GDK_DEVICE_TYPE_SLAVE,
// unless there is an ongoing grab on them, see gdk_device_grab().
/*

C function : gdk_device_get_position
*/
func (recv *Device) GetPosition() (*Screen, int32, int32) {
	var c_screen *C.GdkScreen

	var c_x C.gint

	var c_y C.gint

	C.gdk_device_get_position((*C.GdkDevice)(recv.native), &c_screen, &c_x, &c_y)

	screen := ScreenNewFromC(unsafe.Pointer(c_screen))

	x := (int32)(c_x)

	y := (int32)(c_y)

	return screen, x, y
}

// Obtains the window underneath @device, returning the location of the device in @win_x and @win_y. Returns
// %NULL if the window tree under @device is not known to GDK (for example, belongs to another application).
//
// As a slave device coordinates are those of its master pointer, This
// function may not be called on devices of type %GDK_DEVICE_TYPE_SLAVE,
// unless there is an ongoing grab on them, see gdk_device_grab().
/*

C function : gdk_device_get_window_at_position
*/
func (recv *Device) GetWindowAtPosition() (*Window, int32, int32) {
	var c_win_x C.gint

	var c_win_y C.gint

	retC := C.gdk_device_get_window_at_position((*C.GdkDevice)(recv.native), &c_win_x, &c_win_y)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	winX := (int32)(c_win_x)

	winY := (int32)(c_win_y)

	return retGo, winX, winY
}

// Obtains the window underneath @device, returning the location of the device in @win_x and @win_y in
// double precision. Returns %NULL if the window tree under @device is not known to GDK (for example,
// belongs to another application).
//
// As a slave device coordinates are those of its master pointer, This
// function may not be called on devices of type %GDK_DEVICE_TYPE_SLAVE,
// unless there is an ongoing grab on them, see gdk_device_grab().
/*

C function : gdk_device_get_window_at_position_double
*/
func (recv *Device) GetWindowAtPositionDouble() (*Window, float64, float64) {
	var c_win_x C.gdouble

	var c_win_y C.gdouble

	retC := C.gdk_device_get_window_at_position_double((*C.GdkDevice)(recv.native), &c_win_x, &c_win_y)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	winX := (float64)(c_win_x)

	winY := (float64)(c_win_y)

	return retGo, winX, winY
}

// Grabs the device so that all events coming from this device are passed to
// this application until the device is ungrabbed with gdk_device_ungrab(),
// or the window becomes unviewable. This overrides any previous grab on the device
// by this client.
//
// Note that @device and @window need to be on the same display.
//
// Device grabs are used for operations which need complete control over the
// given device events (either pointer or keyboard). For example in GTK+ this
// is used for Drag and Drop operations, popup menus and such.
//
// Note that if the event mask of an X window has selected both button press
// and button release events, then a button press event will cause an automatic
// pointer grab until the button is released. X does this automatically since
// most applications expect to receive button press and release events in pairs.
// It is equivalent to a pointer grab on the window with @owner_events set to
// %TRUE.
//
// If you set up anything at the time you take the grab that needs to be
// cleaned up when the grab ends, you should handle the #GdkEventGrabBroken
// events that are emitted when the grab ends unvoluntarily.
/*

C function : gdk_device_grab
*/
func (recv *Device) Grab(window *Window, grabOwnership GrabOwnership, ownerEvents bool, eventMask EventMask, cursor *Cursor, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_grab_ownership := (C.GdkGrabOwnership)(grabOwnership)

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_event_mask := (C.GdkEventMask)(eventMask)

	c_cursor := (*C.GdkCursor)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkCursor)(cursor.ToC())
	}

	c_time_ := (C.guint32)(time)

	retC := C.gdk_device_grab((*C.GdkDevice)(recv.native), c_window, c_grab_ownership, c_owner_events, c_event_mask, c_cursor, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// Returns a #GList of #GdkAtoms, containing the labels for
// the axes that @device currently has.
/*

C function : gdk_device_list_axes
*/
func (recv *Device) ListAxes() *glib.List {
	retC := C.gdk_device_list_axes((*C.GdkDevice)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Release any grab on @device.
/*

C function : gdk_device_ungrab
*/
func (recv *Device) Ungrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_device_ungrab((*C.GdkDevice)(recv.native), c_time_)

	return
}

// Warps @device in @display to the point @x,@y on
// the screen @screen, unless the device is confined
// to a window by a grab, in which case it will be moved
// as far as allowed by the grab. Warping the pointer
// creates events as if the user had moved the mouse
// instantaneously to the destination.
//
// Note that the pointer should normally be under the
// control of the user. This function was added to cover
// some rare use cases like keyboard navigation support
// for the color picker in the #GtkColorSelectionDialog.
/*

C function : gdk_device_warp
*/
func (recv *Device) Warp(screen *Screen, x int32, y int32) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_device_warp((*C.GdkDevice)(recv.native), c_screen, c_x, c_y)

	return
}

// Returns the client pointer, that is, the master pointer that acts as the core pointer
// for this application. In X11, window managers may change this depending on the interaction
// pattern under the presence of several pointers.
//
// You should use this function seldomly, only in code that isn’t triggered by a #GdkEvent
// and there aren’t other means to get a meaningful #GdkDevice to operate on.
/*

C function : gdk_device_manager_get_client_pointer
*/
func (recv *DeviceManager) GetClientPointer() *Device {
	retC := C.gdk_device_manager_get_client_pointer((*C.GdkDeviceManager)(recv.native))
	retGo := DeviceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GdkDisplay associated to @device_manager.
/*

C function : gdk_device_manager_get_display
*/
func (recv *DeviceManager) GetDisplay() *Display {
	retC := C.gdk_device_manager_get_display((*C.GdkDeviceManager)(recv.native))
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the list of devices of type @type currently attached to
// @device_manager.
/*

C function : gdk_device_manager_list_devices
*/
func (recv *DeviceManager) ListDevices(type_ DeviceType) *glib.List {
	c_type := (C.GdkDeviceType)(type_)

	retC := C.gdk_device_manager_list_devices((*C.GdkDeviceManager)(recv.native), c_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns a #GdkAppLaunchContext suitable for launching
// applications on the given display.
/*

C function : gdk_display_get_app_launch_context
*/
func (recv *Display) GetAppLaunchContext() *AppLaunchContext {
	retC := C.gdk_display_get_app_launch_context((*C.GdkDisplay)(recv.native))
	retGo := AppLaunchContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the #GdkDeviceManager associated to @display.
/*

C function : gdk_display_get_device_manager
*/
func (recv *Display) GetDeviceManager() *DeviceManager {
	retC := C.gdk_display_get_device_manager((*C.GdkDisplay)(recv.native))
	var retGo (*DeviceManager)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DeviceManagerNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns whether the display has events that are waiting
// to be processed.
/*

C function : gdk_display_has_pending
*/
func (recv *Display) HasPending() bool {
	retC := C.gdk_display_has_pending((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Indicates to the GUI environment that the application has
// finished loading, using a given identifier.
//
// GTK+ will call this function automatically for #GtkWindow
// with custom startup-notification identifier unless
// gtk_window_set_auto_startup_notification() is called to
// disable that feature.
/*

C function : gdk_display_notify_startup_complete
*/
func (recv *Display) NotifyStartupComplete(startupId string) {
	c_startup_id := C.CString(startupId)
	defer C.free(unsafe.Pointer(c_startup_id))

	C.gdk_display_notify_startup_complete((*C.GdkDisplay)(recv.native), c_startup_id)

	return
}

// Opens a display.
/*

C function : gdk_display_manager_open_display
*/
func (recv *DisplayManager) OpenDisplay(name string) *Display {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gdk_display_manager_open_display((*C.GdkDisplayManager)(recv.native), c_name)
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the destination windw for the DND operation.
/*

C function : gdk_drag_context_get_dest_window
*/
func (recv *DragContext) GetDestWindow() *Window {
	retC := C.gdk_drag_context_get_dest_window((*C.GdkDragContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the drag protocol thats used by this context.
/*

C function : gdk_drag_context_get_protocol
*/
func (recv *DragContext) GetProtocol() DragProtocol {
	retC := C.gdk_drag_context_get_protocol((*C.GdkDragContext)(recv.native))
	retGo := (DragProtocol)(retC)

	return retGo
}

// Returns whether the Num Lock modifer is locked.
/*

C function : gdk_keymap_get_num_lock_state
*/
func (recv *Keymap) GetNumLockState() bool {
	retC := C.gdk_keymap_get_num_lock_state((*C.GdkKeymap)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported signal 'create-surface' for Window : unsupported parameter width : type gint :

// Retrieves a #GdkCursor pointer for the @device currently set on the
// specified #GdkWindow, or %NULL.  If the return value is %NULL then
// there is no custom cursor set on the specified window, and it is
// using the cursor for its parent window.
/*

C function : gdk_window_get_device_cursor
*/
func (recv *Window) GetDeviceCursor(device *Device) *Cursor {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gdk_window_get_device_cursor((*C.GdkWindow)(recv.native), c_device)
	var retGo (*Cursor)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CursorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the event mask for @window corresponding to an specific device.
/*

C function : gdk_window_get_device_events
*/
func (recv *Window) GetDeviceEvents(device *Device) EventMask {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gdk_window_get_device_events((*C.GdkWindow)(recv.native), c_device)
	retGo := (EventMask)(retC)

	return retGo
}

// Unsupported : gdk_window_get_device_position : unsupported parameter mask : GdkModifierType* with indirection level of 1

// Finds out the DND protocol supported by a window.
/*

C function : gdk_window_get_drag_protocol
*/
func (recv *Window) GetDragProtocol() (DragProtocol, *Window) {
	var c_target *C.GdkWindow

	retC := C.gdk_window_get_drag_protocol((*C.GdkWindow)(recv.native), &c_target)
	retGo := (DragProtocol)(retC)

	target := WindowNewFromC(unsafe.Pointer(c_target))

	return retGo, target
}

// Returns %TRUE if the window is aware of the existence of multiple
// devices.
/*

C function : gdk_window_get_support_multidevice
*/
func (recv *Window) GetSupportMultidevice() bool {
	retC := C.gdk_window_get_support_multidevice((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets a specific #GdkCursor for a given device when it gets inside @window.
// Use gdk_cursor_new_for_display() or gdk_cursor_new_from_pixbuf() to create
// the cursor. To make the cursor invisible, use %GDK_BLANK_CURSOR. Passing
// %NULL for the @cursor argument to gdk_window_set_cursor() means that
// @window will use the cursor of its parent window. Most windows should
// use this default.
/*

C function : gdk_window_set_device_cursor
*/
func (recv *Window) SetDeviceCursor(device *Device, cursor *Cursor) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_cursor := (*C.GdkCursor)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkCursor)(cursor.ToC())
	}

	C.gdk_window_set_device_cursor((*C.GdkWindow)(recv.native), c_device, c_cursor)

	return
}

// Sets the event mask for a given device (Normally a floating device, not
// attached to any visible pointer) to @window. For example, an event mask
// including #GDK_BUTTON_PRESS_MASK means the window should report button
// press events. The event mask is the bitwise OR of values from the
// #GdkEventMask enumeration.
//
// See the [input handling overview][event-masks] for details.
/*

C function : gdk_window_set_device_events
*/
func (recv *Window) SetDeviceEvents(device *Device, eventMask EventMask) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_event_mask := (C.GdkEventMask)(eventMask)

	C.gdk_window_set_device_events((*C.GdkWindow)(recv.native), c_device, c_event_mask)

	return
}

// Sets the event mask for any floating device (i.e. not attached to any
// visible pointer) that has the source defined as @source. This event
// mask will be applied both to currently existing, newly added devices
// after this call, and devices being attached/detached.
/*

C function : gdk_window_set_source_events
*/
func (recv *Window) SetSourceEvents(source InputSource, eventMask EventMask) {
	c_source := (C.GdkInputSource)(source)

	c_event_mask := (C.GdkEventMask)(eventMask)

	C.gdk_window_set_source_events((*C.GdkWindow)(recv.native), c_source, c_event_mask)

	return
}

// This function will enable multidevice features in @window.
//
// Multidevice aware windows will need to handle properly multiple,
// per device enter/leave events, device grabs and grab ownerships.
/*

C function : gdk_window_set_support_multidevice
*/
func (recv *Window) SetSupportMultidevice(supportMultidevice bool) {
	c_support_multidevice :=
		boolToGboolean(supportMultidevice)

	C.gdk_window_set_support_multidevice((*C.GdkWindow)(recv.native), c_support_multidevice)

	return
}
