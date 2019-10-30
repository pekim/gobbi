// Code generated - DO NOT EDIT.
// +build gdk_2.10

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"reflect"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <gdk_event.h>
// #include <stdlib.h>
/*

	void device_changedHandler(GObject *, gpointer);

	static gulong Device_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(device_changedHandler), data);
	}

*/
/*

	void devicemanager_deviceAddedHandler(GObject *, GdkDevice *, gpointer);

	static gulong DeviceManager_signal_connect_device_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "device-added", G_CALLBACK(devicemanager_deviceAddedHandler), data);
	}

*/
/*

	void devicemanager_deviceChangedHandler(GObject *, GdkDevice *, gpointer);

	static gulong DeviceManager_signal_connect_device_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "device-changed", G_CALLBACK(devicemanager_deviceChangedHandler), data);
	}

*/
/*

	void devicemanager_deviceRemovedHandler(GObject *, GdkDevice *, gpointer);

	static gulong DeviceManager_signal_connect_device_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "device-removed", G_CALLBACK(devicemanager_deviceRemovedHandler), data);
	}

*/
/*

	void display_closedHandler(GObject *, gboolean, gpointer);

	static gulong Display_signal_connect_closed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "closed", G_CALLBACK(display_closedHandler), data);
	}

*/
/*

	void display_openedHandler(GObject *, gpointer);

	static gulong Display_signal_connect_opened(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "opened", G_CALLBACK(display_openedHandler), data);
	}

*/
/*

	void displaymanager_displayOpenedHandler(GObject *, GdkDisplay *, gpointer);

	static gulong DisplayManager_signal_connect_display_opened(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "display-opened", G_CALLBACK(displaymanager_displayOpenedHandler), data);
	}

*/
/*

	void frameclock_afterPaintHandler(GObject *, gpointer);

	static gulong FrameClock_signal_connect_after_paint(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "after-paint", G_CALLBACK(frameclock_afterPaintHandler), data);
	}

*/
/*

	void frameclock_beforePaintHandler(GObject *, gpointer);

	static gulong FrameClock_signal_connect_before_paint(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "before-paint", G_CALLBACK(frameclock_beforePaintHandler), data);
	}

*/
/*

	void frameclock_flushEventsHandler(GObject *, gpointer);

	static gulong FrameClock_signal_connect_flush_events(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "flush-events", G_CALLBACK(frameclock_flushEventsHandler), data);
	}

*/
/*

	void frameclock_layoutHandler(GObject *, gpointer);

	static gulong FrameClock_signal_connect_layout(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "layout", G_CALLBACK(frameclock_layoutHandler), data);
	}

*/
/*

	void frameclock_paintHandler(GObject *, gpointer);

	static gulong FrameClock_signal_connect_paint(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "paint", G_CALLBACK(frameclock_paintHandler), data);
	}

*/
/*

	void frameclock_resumeEventsHandler(GObject *, gpointer);

	static gulong FrameClock_signal_connect_resume_events(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "resume-events", G_CALLBACK(frameclock_resumeEventsHandler), data);
	}

*/
/*

	void frameclock_updateHandler(GObject *, gpointer);

	static gulong FrameClock_signal_connect_update(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "update", G_CALLBACK(frameclock_updateHandler), data);
	}

*/
/*

	void keymap_directionChangedHandler(GObject *, gpointer);

	static gulong Keymap_signal_connect_direction_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "direction-changed", G_CALLBACK(keymap_directionChangedHandler), data);
	}

*/
/*

	void keymap_keysChangedHandler(GObject *, gpointer);

	static gulong Keymap_signal_connect_keys_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "keys-changed", G_CALLBACK(keymap_keysChangedHandler), data);
	}

*/
/*

	void screen_compositedChangedHandler(GObject *, gpointer);

	static gulong Screen_signal_connect_composited_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "composited-changed", G_CALLBACK(screen_compositedChangedHandler), data);
	}

*/
/*

	void screen_sizeChangedHandler(GObject *, gpointer);

	static gulong Screen_signal_connect_size_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "size-changed", G_CALLBACK(screen_sizeChangedHandler), data);
	}

*/
import "C"

var gobjectClassGoTypeMap = make(map[string]reflect.Type)

// Unsupported : alias has no type generator for none, void

type DragAction C.GdkDragAction

const (
	GDK_ACTION_DEFAULT DragAction = 1
	GDK_ACTION_COPY    DragAction = 2
	GDK_ACTION_MOVE    DragAction = 4
	GDK_ACTION_LINK    DragAction = 8
	GDK_ACTION_PRIVATE DragAction = 16
	GDK_ACTION_ASK     DragAction = 32
)

type EventMask C.GdkEventMask

const (
	GDK_EXPOSURE_MASK            EventMask = 2
	GDK_POINTER_MOTION_MASK      EventMask = 4
	GDK_POINTER_MOTION_HINT_MASK EventMask = 8
	GDK_BUTTON_MOTION_MASK       EventMask = 16
	GDK_BUTTON1_MOTION_MASK      EventMask = 32
	GDK_BUTTON2_MOTION_MASK      EventMask = 64
	GDK_BUTTON3_MOTION_MASK      EventMask = 128
	GDK_BUTTON_PRESS_MASK        EventMask = 256
	GDK_BUTTON_RELEASE_MASK      EventMask = 512
	GDK_KEY_PRESS_MASK           EventMask = 1024
	GDK_KEY_RELEASE_MASK         EventMask = 2048
	GDK_ENTER_NOTIFY_MASK        EventMask = 4096
	GDK_LEAVE_NOTIFY_MASK        EventMask = 8192
	GDK_FOCUS_CHANGE_MASK        EventMask = 16384
	GDK_STRUCTURE_MASK           EventMask = 32768
	GDK_PROPERTY_CHANGE_MASK     EventMask = 65536
	GDK_VISIBILITY_NOTIFY_MASK   EventMask = 131072
	GDK_PROXIMITY_IN_MASK        EventMask = 262144
	GDK_PROXIMITY_OUT_MASK       EventMask = 524288
	GDK_SUBSTRUCTURE_MASK        EventMask = 1048576
	GDK_SCROLL_MASK              EventMask = 2097152
	GDK_TOUCH_MASK               EventMask = 4194304
	GDK_SMOOTH_SCROLL_MASK       EventMask = 8388608
	GDK_TOUCHPAD_GESTURE_MASK    EventMask = 16777216
	GDK_TABLET_PAD_MASK          EventMask = 33554432
	GDK_ALL_EVENTS_MASK          EventMask = 67108862
)

type ModifierType C.guint

const (
	GDK_SHIFT_MASK                ModifierType = 1
	GDK_LOCK_MASK                 ModifierType = 2
	GDK_CONTROL_MASK              ModifierType = 4
	GDK_MOD1_MASK                 ModifierType = 8
	GDK_MOD2_MASK                 ModifierType = 16
	GDK_MOD3_MASK                 ModifierType = 32
	GDK_MOD4_MASK                 ModifierType = 64
	GDK_MOD5_MASK                 ModifierType = 128
	GDK_BUTTON1_MASK              ModifierType = 256
	GDK_BUTTON2_MASK              ModifierType = 512
	GDK_BUTTON3_MASK              ModifierType = 1024
	GDK_BUTTON4_MASK              ModifierType = 2048
	GDK_BUTTON5_MASK              ModifierType = 4096
	GDK_MODIFIER_RESERVED_13_MASK ModifierType = 8192
	GDK_MODIFIER_RESERVED_14_MASK ModifierType = 16384
	GDK_MODIFIER_RESERVED_15_MASK ModifierType = 32768
	GDK_MODIFIER_RESERVED_16_MASK ModifierType = 65536
	GDK_MODIFIER_RESERVED_17_MASK ModifierType = 131072
	GDK_MODIFIER_RESERVED_18_MASK ModifierType = 262144
	GDK_MODIFIER_RESERVED_19_MASK ModifierType = 524288
	GDK_MODIFIER_RESERVED_20_MASK ModifierType = 1048576
	GDK_MODIFIER_RESERVED_21_MASK ModifierType = 2097152
	GDK_MODIFIER_RESERVED_22_MASK ModifierType = 4194304
	GDK_MODIFIER_RESERVED_23_MASK ModifierType = 8388608
	GDK_MODIFIER_RESERVED_24_MASK ModifierType = 16777216
	GDK_MODIFIER_RESERVED_25_MASK ModifierType = 33554432
	GDK_SUPER_MASK                ModifierType = 67108864
	GDK_HYPER_MASK                ModifierType = 134217728
	GDK_META_MASK                 ModifierType = 268435456
	GDK_MODIFIER_RESERVED_29_MASK ModifierType = 536870912
	GDK_RELEASE_MASK              ModifierType = 1073741824
	GDK_MODIFIER_MASK             ModifierType = 1543512063
)

type WMDecoration C.GdkWMDecoration

const (
	GDK_DECOR_ALL      WMDecoration = 1
	GDK_DECOR_BORDER   WMDecoration = 2
	GDK_DECOR_RESIZEH  WMDecoration = 4
	GDK_DECOR_TITLE    WMDecoration = 8
	GDK_DECOR_MENU     WMDecoration = 16
	GDK_DECOR_MINIMIZE WMDecoration = 32
	GDK_DECOR_MAXIMIZE WMDecoration = 64
)

type WMFunction C.GdkWMFunction

const (
	GDK_FUNC_ALL      WMFunction = 1
	GDK_FUNC_RESIZE   WMFunction = 2
	GDK_FUNC_MOVE     WMFunction = 4
	GDK_FUNC_MINIMIZE WMFunction = 8
	GDK_FUNC_MAXIMIZE WMFunction = 16
	GDK_FUNC_CLOSE    WMFunction = 32
)

type WindowAttributesType C.GdkWindowAttributesType

const (
	GDK_WA_TITLE     WindowAttributesType = 2
	GDK_WA_X         WindowAttributesType = 4
	GDK_WA_Y         WindowAttributesType = 8
	GDK_WA_CURSOR    WindowAttributesType = 16
	GDK_WA_VISUAL    WindowAttributesType = 32
	GDK_WA_WMCLASS   WindowAttributesType = 64
	GDK_WA_NOREDIR   WindowAttributesType = 128
	GDK_WA_TYPE_HINT WindowAttributesType = 256
)

type WindowHints C.GdkWindowHints

const (
	GDK_HINT_POS         WindowHints = 1
	GDK_HINT_MIN_SIZE    WindowHints = 2
	GDK_HINT_MAX_SIZE    WindowHints = 4
	GDK_HINT_BASE_SIZE   WindowHints = 8
	GDK_HINT_ASPECT      WindowHints = 16
	GDK_HINT_RESIZE_INC  WindowHints = 32
	GDK_HINT_WIN_GRAVITY WindowHints = 64
	GDK_HINT_USER_POS    WindowHints = 128
	GDK_HINT_USER_SIZE   WindowHints = 256
)

type WindowState C.GdkWindowState

const (
	GDK_WINDOW_STATE_WITHDRAWN        WindowState = 1
	GDK_WINDOW_STATE_ICONIFIED        WindowState = 2
	GDK_WINDOW_STATE_MAXIMIZED        WindowState = 4
	GDK_WINDOW_STATE_STICKY           WindowState = 8
	GDK_WINDOW_STATE_FULLSCREEN       WindowState = 16
	GDK_WINDOW_STATE_ABOVE            WindowState = 32
	GDK_WINDOW_STATE_BELOW            WindowState = 64
	GDK_WINDOW_STATE_FOCUSED          WindowState = 128
	GDK_WINDOW_STATE_TILED            WindowState = 256
	GDK_WINDOW_STATE_TOP_TILED        WindowState = 512
	GDK_WINDOW_STATE_TOP_RESIZABLE    WindowState = 1024
	GDK_WINDOW_STATE_RIGHT_TILED      WindowState = 2048
	GDK_WINDOW_STATE_RIGHT_RESIZABLE  WindowState = 4096
	GDK_WINDOW_STATE_BOTTOM_TILED     WindowState = 8192
	GDK_WINDOW_STATE_BOTTOM_RESIZABLE WindowState = 16384
	GDK_WINDOW_STATE_LEFT_TILED       WindowState = 32768
	GDK_WINDOW_STATE_LEFT_RESIZABLE   WindowState = 65536
)

// AddToGobjectClassGoTypeMap : GdkAppLaunchContext

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AppLaunchContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AppLaunchContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AppLaunchContext with another AppLaunchContext, and returns true if they represent the same GObject.
func (recv *AppLaunchContext) Equals(other *AppLaunchContext) bool {
	return other.ToC() == recv.ToC()
}

// AppLaunchContext upcasts to *AppLaunchContext
func (recv *AppLaunchContext) AppLaunchContext() *gio.AppLaunchContext {
	return gio.AppLaunchContextNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *AppLaunchContext) Object() *gobject.Object {
	return recv.AppLaunchContext().Object()
}

// CastToWidget down casts any arbitrary Object to AppLaunchContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a AppLaunchContext.
func CastToAppLaunchContext(object *gobject.Object) *AppLaunchContext {
	return AppLaunchContextNewFromC(object.ToC())
}

// AddToGobjectClassGoTypeMap : GdkCursor

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Cursor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Cursor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Cursor with another Cursor, and returns true if they represent the same GObject.
func (recv *Cursor) Equals(other *Cursor) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Cursor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Cursor.
// Exercise care, as this is a potentially dangerous function if the Object is not a Cursor.
func CastToCursor(object *gobject.Object) *Cursor {
	return CursorNewFromC(object.ToC())
}

// CursorNew is a wrapper around the C function gdk_cursor_new.
func CursorNew(cursorType CursorType) *Cursor {
	c_cursor_type := (C.GdkCursorType)(cursorType)

	retC := C.gdk_cursor_new(c_cursor_type)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// CursorNewForDisplay is a wrapper around the C function gdk_cursor_new_for_display.
func CursorNewForDisplay(display *Display, cursorType CursorType) *Cursor {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_cursor_type := (C.GdkCursorType)(cursorType)

	retC := C.gdk_cursor_new_for_display(c_display, c_cursor_type)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// CursorNewFromName is a wrapper around the C function gdk_cursor_new_from_name.
func CursorNewFromName(display *Display, name string) *Cursor {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gdk_cursor_new_from_name(c_display, c_name)
	var retGo (*Cursor)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CursorNewFromC(unsafe.Pointer(retC))
	}

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// CursorNewFromPixbuf is a wrapper around the C function gdk_cursor_new_from_pixbuf.
func CursorNewFromPixbuf(display *Display, pixbuf *gdkpixbuf.Pixbuf, x int32, y int32) *Cursor {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gdk_cursor_new_from_pixbuf(c_display, c_pixbuf, c_x, c_y)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_cursor_get_display.
func (recv *Cursor) GetDisplay() *Display {
	retC := C.gdk_cursor_get_display((*C.GdkCursor)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetImage is a wrapper around the C function gdk_cursor_get_image.
func (recv *Cursor) GetImage() *gdkpixbuf.Pixbuf {
	retC := C.gdk_cursor_get_image((*C.GdkCursor)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

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

// AddToGobjectClassGoTypeMap : GdkDevice

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Device) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Device) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Device with another Device, and returns true if they represent the same GObject.
func (recv *Device) Equals(other *Device) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Device) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Device.
// Exercise care, as this is a potentially dangerous function if the Object is not a Device.
func CastToDevice(object *gobject.Object) *Device {
	return DeviceNewFromC(object.ToC())
}

type signalDeviceChangedDetail struct {
	callback  DeviceSignalChangedCallback
	handlerID C.gulong
}

var signalDeviceChangedId int
var signalDeviceChangedMap = make(map[int]signalDeviceChangedDetail)
var signalDeviceChangedLock sync.RWMutex

// DeviceSignalChangedCallback is a callback function for a 'changed' signal emitted from a Device.
type DeviceSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Device.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Device) ConnectChanged(callback DeviceSignalChangedCallback) int {
	signalDeviceChangedLock.Lock()
	defer signalDeviceChangedLock.Unlock()

	signalDeviceChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Device_signal_connect_changed(instance, C.gpointer(uintptr(signalDeviceChangedId)))

	detail := signalDeviceChangedDetail{callback, handlerID}
	signalDeviceChangedMap[signalDeviceChangedId] = detail

	return signalDeviceChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Device.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Device) DisconnectChanged(connectionID int) {
	signalDeviceChangedLock.Lock()
	defer signalDeviceChangedLock.Unlock()

	detail, exists := signalDeviceChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDeviceChangedMap, connectionID)
}

//export device_changedHandler
func device_changedHandler(_ *C.GObject, data C.gpointer) {
	signalDeviceChangedLock.RLock()
	defer signalDeviceChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDeviceChangedMap[index].callback
	callback()
}

// gdk_device_free_history : unsupported parameter events :
// DeviceGrabInfoLibgtkOnly is a wrapper around the C function gdk_device_grab_info_libgtk_only.
func DeviceGrabInfoLibgtkOnly(display *Display, device *Device) (bool, *Window, bool) {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	var c_grab_window *C.GdkWindow

	var c_owner_events C.gboolean

	retC := C.gdk_device_grab_info_libgtk_only(c_display, c_device, &c_grab_window, &c_owner_events)
	retGo := retC == C.TRUE

	grabWindow := WindowNewFromC(unsafe.Pointer(c_grab_window))

	ownerEvents := c_owner_events == C.TRUE

	return retGo, grabWindow, ownerEvents
}

// GetAxis is a wrapper around the C function gdk_device_get_axis.
func (recv *Device) GetAxis(axes []float64, use AxisUse) (bool, float64) {
	c_axes_array := make([]C.gdouble, len(axes)+1, len(axes)+1)
	for i, item := range axes {
		c := (C.gdouble)(item)
		c_axes_array[i] = c
	}
	c_axes_array[len(axes)] = 0
	c_axes_arrayPtr := &c_axes_array[0]
	c_axes := (*C.gdouble)(unsafe.Pointer(c_axes_arrayPtr))

	c_use := (C.GdkAxisUse)(use)

	var c_value C.gdouble

	retC := C.gdk_device_get_axis((*C.GdkDevice)(recv.native), c_axes, c_use, &c_value)
	retGo := retC == C.TRUE

	value := (float64)(c_value)

	return retGo, value
}

// Unsupported : gdk_device_get_history : unsupported parameter events : output array param events

// GetState is a wrapper around the C function gdk_device_get_state.
func (recv *Device) GetState(window *Window, axes []float64) ModifierType {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_axes_array := make([]C.gdouble, len(axes)+1, len(axes)+1)
	for i, item := range axes {
		c := (C.gdouble)(item)
		c_axes_array[i] = c
	}
	c_axes_array[len(axes)] = 0
	c_axes_arrayPtr := &c_axes_array[0]
	c_axes := (*C.gdouble)(unsafe.Pointer(c_axes_arrayPtr))

	var c_mask C.GdkModifierType

	C.gdk_device_get_state((*C.GdkDevice)(recv.native), c_window, c_axes, &c_mask)

	mask := (ModifierType)(c_mask)

	return mask
}

// ListSlaveDevices is a wrapper around the C function gdk_device_list_slave_devices.
func (recv *Device) ListSlaveDevices() *glib.List {
	retC := C.gdk_device_list_slave_devices((*C.GdkDevice)(recv.native))
	var retGo (*glib.List)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.ListNewFromC(unsafe.Pointer(retC))
	}

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

// AddToGobjectClassGoTypeMap : GdkDeviceManager

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DeviceManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DeviceManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DeviceManager with another DeviceManager, and returns true if they represent the same GObject.
func (recv *DeviceManager) Equals(other *DeviceManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DeviceManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DeviceManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a DeviceManager.
func CastToDeviceManager(object *gobject.Object) *DeviceManager {
	return DeviceManagerNewFromC(object.ToC())
}

type signalDeviceManagerDeviceAddedDetail struct {
	callback  DeviceManagerSignalDeviceAddedCallback
	handlerID C.gulong
}

var signalDeviceManagerDeviceAddedId int
var signalDeviceManagerDeviceAddedMap = make(map[int]signalDeviceManagerDeviceAddedDetail)
var signalDeviceManagerDeviceAddedLock sync.RWMutex

// DeviceManagerSignalDeviceAddedCallback is a callback function for a 'device-added' signal emitted from a DeviceManager.
type DeviceManagerSignalDeviceAddedCallback func(device *Device)

/*
ConnectDeviceAdded connects the callback to the 'device-added' signal for the DeviceManager.

The returned value represents the connection, and may be passed to DisconnectDeviceAdded to remove it.
*/
func (recv *DeviceManager) ConnectDeviceAdded(callback DeviceManagerSignalDeviceAddedCallback) int {
	signalDeviceManagerDeviceAddedLock.Lock()
	defer signalDeviceManagerDeviceAddedLock.Unlock()

	signalDeviceManagerDeviceAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DeviceManager_signal_connect_device_added(instance, C.gpointer(uintptr(signalDeviceManagerDeviceAddedId)))

	detail := signalDeviceManagerDeviceAddedDetail{callback, handlerID}
	signalDeviceManagerDeviceAddedMap[signalDeviceManagerDeviceAddedId] = detail

	return signalDeviceManagerDeviceAddedId
}

/*
DisconnectDeviceAdded disconnects a callback from the 'device-added' signal for the DeviceManager.

The connectionID should be a value returned from a call to ConnectDeviceAdded.
*/
func (recv *DeviceManager) DisconnectDeviceAdded(connectionID int) {
	signalDeviceManagerDeviceAddedLock.Lock()
	defer signalDeviceManagerDeviceAddedLock.Unlock()

	detail, exists := signalDeviceManagerDeviceAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDeviceManagerDeviceAddedMap, connectionID)
}

//export devicemanager_deviceAddedHandler
func devicemanager_deviceAddedHandler(_ *C.GObject, c_device *C.GdkDevice, data C.gpointer) {
	signalDeviceManagerDeviceAddedLock.RLock()
	defer signalDeviceManagerDeviceAddedLock.RUnlock()

	device := DeviceNewFromC(unsafe.Pointer(c_device))

	index := int(uintptr(data))
	callback := signalDeviceManagerDeviceAddedMap[index].callback
	callback(device)
}

type signalDeviceManagerDeviceChangedDetail struct {
	callback  DeviceManagerSignalDeviceChangedCallback
	handlerID C.gulong
}

var signalDeviceManagerDeviceChangedId int
var signalDeviceManagerDeviceChangedMap = make(map[int]signalDeviceManagerDeviceChangedDetail)
var signalDeviceManagerDeviceChangedLock sync.RWMutex

// DeviceManagerSignalDeviceChangedCallback is a callback function for a 'device-changed' signal emitted from a DeviceManager.
type DeviceManagerSignalDeviceChangedCallback func(device *Device)

/*
ConnectDeviceChanged connects the callback to the 'device-changed' signal for the DeviceManager.

The returned value represents the connection, and may be passed to DisconnectDeviceChanged to remove it.
*/
func (recv *DeviceManager) ConnectDeviceChanged(callback DeviceManagerSignalDeviceChangedCallback) int {
	signalDeviceManagerDeviceChangedLock.Lock()
	defer signalDeviceManagerDeviceChangedLock.Unlock()

	signalDeviceManagerDeviceChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DeviceManager_signal_connect_device_changed(instance, C.gpointer(uintptr(signalDeviceManagerDeviceChangedId)))

	detail := signalDeviceManagerDeviceChangedDetail{callback, handlerID}
	signalDeviceManagerDeviceChangedMap[signalDeviceManagerDeviceChangedId] = detail

	return signalDeviceManagerDeviceChangedId
}

/*
DisconnectDeviceChanged disconnects a callback from the 'device-changed' signal for the DeviceManager.

The connectionID should be a value returned from a call to ConnectDeviceChanged.
*/
func (recv *DeviceManager) DisconnectDeviceChanged(connectionID int) {
	signalDeviceManagerDeviceChangedLock.Lock()
	defer signalDeviceManagerDeviceChangedLock.Unlock()

	detail, exists := signalDeviceManagerDeviceChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDeviceManagerDeviceChangedMap, connectionID)
}

//export devicemanager_deviceChangedHandler
func devicemanager_deviceChangedHandler(_ *C.GObject, c_device *C.GdkDevice, data C.gpointer) {
	signalDeviceManagerDeviceChangedLock.RLock()
	defer signalDeviceManagerDeviceChangedLock.RUnlock()

	device := DeviceNewFromC(unsafe.Pointer(c_device))

	index := int(uintptr(data))
	callback := signalDeviceManagerDeviceChangedMap[index].callback
	callback(device)
}

type signalDeviceManagerDeviceRemovedDetail struct {
	callback  DeviceManagerSignalDeviceRemovedCallback
	handlerID C.gulong
}

var signalDeviceManagerDeviceRemovedId int
var signalDeviceManagerDeviceRemovedMap = make(map[int]signalDeviceManagerDeviceRemovedDetail)
var signalDeviceManagerDeviceRemovedLock sync.RWMutex

// DeviceManagerSignalDeviceRemovedCallback is a callback function for a 'device-removed' signal emitted from a DeviceManager.
type DeviceManagerSignalDeviceRemovedCallback func(device *Device)

/*
ConnectDeviceRemoved connects the callback to the 'device-removed' signal for the DeviceManager.

The returned value represents the connection, and may be passed to DisconnectDeviceRemoved to remove it.
*/
func (recv *DeviceManager) ConnectDeviceRemoved(callback DeviceManagerSignalDeviceRemovedCallback) int {
	signalDeviceManagerDeviceRemovedLock.Lock()
	defer signalDeviceManagerDeviceRemovedLock.Unlock()

	signalDeviceManagerDeviceRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DeviceManager_signal_connect_device_removed(instance, C.gpointer(uintptr(signalDeviceManagerDeviceRemovedId)))

	detail := signalDeviceManagerDeviceRemovedDetail{callback, handlerID}
	signalDeviceManagerDeviceRemovedMap[signalDeviceManagerDeviceRemovedId] = detail

	return signalDeviceManagerDeviceRemovedId
}

/*
DisconnectDeviceRemoved disconnects a callback from the 'device-removed' signal for the DeviceManager.

The connectionID should be a value returned from a call to ConnectDeviceRemoved.
*/
func (recv *DeviceManager) DisconnectDeviceRemoved(connectionID int) {
	signalDeviceManagerDeviceRemovedLock.Lock()
	defer signalDeviceManagerDeviceRemovedLock.Unlock()

	detail, exists := signalDeviceManagerDeviceRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDeviceManagerDeviceRemovedMap, connectionID)
}

//export devicemanager_deviceRemovedHandler
func devicemanager_deviceRemovedHandler(_ *C.GObject, c_device *C.GdkDevice, data C.gpointer) {
	signalDeviceManagerDeviceRemovedLock.RLock()
	defer signalDeviceManagerDeviceRemovedLock.RUnlock()

	device := DeviceNewFromC(unsafe.Pointer(c_device))

	index := int(uintptr(data))
	callback := signalDeviceManagerDeviceRemovedMap[index].callback
	callback(device)
}

// AddToGobjectClassGoTypeMap : GdkDisplay

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Display) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Display) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Display with another Display, and returns true if they represent the same GObject.
func (recv *Display) Equals(other *Display) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Display) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Display.
// Exercise care, as this is a potentially dangerous function if the Object is not a Display.
func CastToDisplay(object *gobject.Object) *Display {
	return DisplayNewFromC(object.ToC())
}

type signalDisplayClosedDetail struct {
	callback  DisplaySignalClosedCallback
	handlerID C.gulong
}

var signalDisplayClosedId int
var signalDisplayClosedMap = make(map[int]signalDisplayClosedDetail)
var signalDisplayClosedLock sync.RWMutex

// DisplaySignalClosedCallback is a callback function for a 'closed' signal emitted from a Display.
type DisplaySignalClosedCallback func(isError bool)

/*
ConnectClosed connects the callback to the 'closed' signal for the Display.

The returned value represents the connection, and may be passed to DisconnectClosed to remove it.
*/
func (recv *Display) ConnectClosed(callback DisplaySignalClosedCallback) int {
	signalDisplayClosedLock.Lock()
	defer signalDisplayClosedLock.Unlock()

	signalDisplayClosedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Display_signal_connect_closed(instance, C.gpointer(uintptr(signalDisplayClosedId)))

	detail := signalDisplayClosedDetail{callback, handlerID}
	signalDisplayClosedMap[signalDisplayClosedId] = detail

	return signalDisplayClosedId
}

/*
DisconnectClosed disconnects a callback from the 'closed' signal for the Display.

The connectionID should be a value returned from a call to ConnectClosed.
*/
func (recv *Display) DisconnectClosed(connectionID int) {
	signalDisplayClosedLock.Lock()
	defer signalDisplayClosedLock.Unlock()

	detail, exists := signalDisplayClosedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDisplayClosedMap, connectionID)
}

//export display_closedHandler
func display_closedHandler(_ *C.GObject, c_is_error C.gboolean, data C.gpointer) {
	signalDisplayClosedLock.RLock()
	defer signalDisplayClosedLock.RUnlock()

	isError := c_is_error == C.TRUE

	index := int(uintptr(data))
	callback := signalDisplayClosedMap[index].callback
	callback(isError)
}

type signalDisplayOpenedDetail struct {
	callback  DisplaySignalOpenedCallback
	handlerID C.gulong
}

var signalDisplayOpenedId int
var signalDisplayOpenedMap = make(map[int]signalDisplayOpenedDetail)
var signalDisplayOpenedLock sync.RWMutex

// DisplaySignalOpenedCallback is a callback function for a 'opened' signal emitted from a Display.
type DisplaySignalOpenedCallback func()

/*
ConnectOpened connects the callback to the 'opened' signal for the Display.

The returned value represents the connection, and may be passed to DisconnectOpened to remove it.
*/
func (recv *Display) ConnectOpened(callback DisplaySignalOpenedCallback) int {
	signalDisplayOpenedLock.Lock()
	defer signalDisplayOpenedLock.Unlock()

	signalDisplayOpenedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Display_signal_connect_opened(instance, C.gpointer(uintptr(signalDisplayOpenedId)))

	detail := signalDisplayOpenedDetail{callback, handlerID}
	signalDisplayOpenedMap[signalDisplayOpenedId] = detail

	return signalDisplayOpenedId
}

/*
DisconnectOpened disconnects a callback from the 'opened' signal for the Display.

The connectionID should be a value returned from a call to ConnectOpened.
*/
func (recv *Display) DisconnectOpened(connectionID int) {
	signalDisplayOpenedLock.Lock()
	defer signalDisplayOpenedLock.Unlock()

	detail, exists := signalDisplayOpenedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDisplayOpenedMap, connectionID)
}

//export display_openedHandler
func display_openedHandler(_ *C.GObject, data C.gpointer) {
	signalDisplayOpenedLock.RLock()
	defer signalDisplayOpenedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDisplayOpenedMap[index].callback
	callback()
}

// DisplayGetDefault is a wrapper around the C function gdk_display_get_default.
func DisplayGetDefault() *Display {
	retC := C.gdk_display_get_default()
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// DisplayOpen is a wrapper around the C function gdk_display_open.
func DisplayOpen(displayName string) *Display {
	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	retC := C.gdk_display_open(c_display_name)
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// DisplayOpenDefaultLibgtkOnly is a wrapper around the C function gdk_display_open_default_libgtk_only.
func DisplayOpenDefaultLibgtkOnly() *Display {
	retC := C.gdk_display_open_default_libgtk_only()
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Beep is a wrapper around the C function gdk_display_beep.
func (recv *Display) Beep() {
	C.gdk_display_beep((*C.GdkDisplay)(recv.native))

	return
}

// Close is a wrapper around the C function gdk_display_close.
func (recv *Display) Close() {
	C.gdk_display_close((*C.GdkDisplay)(recv.native))

	return
}

// DeviceIsGrabbed is a wrapper around the C function gdk_display_device_is_grabbed.
func (recv *Display) DeviceIsGrabbed(device *Device) bool {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gdk_display_device_is_grabbed((*C.GdkDisplay)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// Flush is a wrapper around the C function gdk_display_flush.
func (recv *Display) Flush() {
	C.gdk_display_flush((*C.GdkDisplay)(recv.native))

	return
}

// GetDefaultCursorSize is a wrapper around the C function gdk_display_get_default_cursor_size.
func (recv *Display) GetDefaultCursorSize() uint32 {
	retC := C.gdk_display_get_default_cursor_size((*C.GdkDisplay)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetDefaultGroup is a wrapper around the C function gdk_display_get_default_group.
func (recv *Display) GetDefaultGroup() *Window {
	retC := C.gdk_display_get_default_group((*C.GdkDisplay)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDefaultScreen is a wrapper around the C function gdk_display_get_default_screen.
func (recv *Display) GetDefaultScreen() *Screen {
	retC := C.gdk_display_get_default_screen((*C.GdkDisplay)(recv.native))
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetEvent is a wrapper around the C function gdk_display_get_event.
func (recv *Display) GetEvent() *Event {
	retC := C.gdk_display_get_event((*C.GdkDisplay)(recv.native))
	var retGo (*Event)
	if retC == nil {
		retGo = nil
	} else {
		retGo = EventNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetMaximalCursorSize is a wrapper around the C function gdk_display_get_maximal_cursor_size.
func (recv *Display) GetMaximalCursorSize() (uint32, uint32) {
	var c_width C.guint

	var c_height C.guint

	C.gdk_display_get_maximal_cursor_size((*C.GdkDisplay)(recv.native), &c_width, &c_height)

	width := (uint32)(c_width)

	height := (uint32)(c_height)

	return width, height
}

// GetNScreens is a wrapper around the C function gdk_display_get_n_screens.
func (recv *Display) GetNScreens() int32 {
	retC := C.gdk_display_get_n_screens((*C.GdkDisplay)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetName is a wrapper around the C function gdk_display_get_name.
func (recv *Display) GetName() string {
	retC := C.gdk_display_get_name((*C.GdkDisplay)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPointer is a wrapper around the C function gdk_display_get_pointer.
func (recv *Display) GetPointer() (*Screen, int32, int32, ModifierType) {
	var c_screen *C.GdkScreen

	var c_x C.gint

	var c_y C.gint

	var c_mask C.GdkModifierType

	C.gdk_display_get_pointer((*C.GdkDisplay)(recv.native), &c_screen, &c_x, &c_y, &c_mask)

	screen := ScreenNewFromC(unsafe.Pointer(c_screen))

	x := (int32)(c_x)

	y := (int32)(c_y)

	mask := (ModifierType)(c_mask)

	return screen, x, y, mask
}

// GetScreen is a wrapper around the C function gdk_display_get_screen.
func (recv *Display) GetScreen(screenNum int32) *Screen {
	c_screen_num := (C.gint)(screenNum)

	retC := C.gdk_display_get_screen((*C.GdkDisplay)(recv.native), c_screen_num)
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWindowAtPointer is a wrapper around the C function gdk_display_get_window_at_pointer.
func (recv *Display) GetWindowAtPointer() (*Window, int32, int32) {
	var c_win_x C.gint

	var c_win_y C.gint

	retC := C.gdk_display_get_window_at_pointer((*C.GdkDisplay)(recv.native), &c_win_x, &c_win_y)
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

// KeyboardUngrab is a wrapper around the C function gdk_display_keyboard_ungrab.
func (recv *Display) KeyboardUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_display_keyboard_ungrab((*C.GdkDisplay)(recv.native), c_time_)

	return
}

// ListDevices is a wrapper around the C function gdk_display_list_devices.
func (recv *Display) ListDevices() *glib.List {
	retC := C.gdk_display_list_devices((*C.GdkDisplay)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PeekEvent is a wrapper around the C function gdk_display_peek_event.
func (recv *Display) PeekEvent() *Event {
	retC := C.gdk_display_peek_event((*C.GdkDisplay)(recv.native))
	var retGo (*Event)
	if retC == nil {
		retGo = nil
	} else {
		retGo = EventNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PointerIsGrabbed is a wrapper around the C function gdk_display_pointer_is_grabbed.
func (recv *Display) PointerIsGrabbed() bool {
	retC := C.gdk_display_pointer_is_grabbed((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PointerUngrab is a wrapper around the C function gdk_display_pointer_ungrab.
func (recv *Display) PointerUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_display_pointer_ungrab((*C.GdkDisplay)(recv.native), c_time_)

	return
}

// PutEvent is a wrapper around the C function gdk_display_put_event.
func (recv *Display) PutEvent(event *Event) {
	c_event := (*C.GdkEvent)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEvent)(event.ToC())
	}

	C.gdk_display_put_event((*C.GdkDisplay)(recv.native), c_event)

	return
}

// RequestSelectionNotification is a wrapper around the C function gdk_display_request_selection_notification.
func (recv *Display) RequestSelectionNotification(selection *Atom) bool {
	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	retC := C.gdk_display_request_selection_notification((*C.GdkDisplay)(recv.native), c_selection)
	retGo := retC == C.TRUE

	return retGo
}

// SetDoubleClickDistance is a wrapper around the C function gdk_display_set_double_click_distance.
func (recv *Display) SetDoubleClickDistance(distance uint32) {
	c_distance := (C.guint)(distance)

	C.gdk_display_set_double_click_distance((*C.GdkDisplay)(recv.native), c_distance)

	return
}

// SetDoubleClickTime is a wrapper around the C function gdk_display_set_double_click_time.
func (recv *Display) SetDoubleClickTime(msec uint32) {
	c_msec := (C.guint)(msec)

	C.gdk_display_set_double_click_time((*C.GdkDisplay)(recv.native), c_msec)

	return
}

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets :

// SupportsClipboardPersistence is a wrapper around the C function gdk_display_supports_clipboard_persistence.
func (recv *Display) SupportsClipboardPersistence() bool {
	retC := C.gdk_display_supports_clipboard_persistence((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsCursorAlpha is a wrapper around the C function gdk_display_supports_cursor_alpha.
func (recv *Display) SupportsCursorAlpha() bool {
	retC := C.gdk_display_supports_cursor_alpha((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsCursorColor is a wrapper around the C function gdk_display_supports_cursor_color.
func (recv *Display) SupportsCursorColor() bool {
	retC := C.gdk_display_supports_cursor_color((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsInputShapes is a wrapper around the C function gdk_display_supports_input_shapes.
func (recv *Display) SupportsInputShapes() bool {
	retC := C.gdk_display_supports_input_shapes((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsSelectionNotification is a wrapper around the C function gdk_display_supports_selection_notification.
func (recv *Display) SupportsSelectionNotification() bool {
	retC := C.gdk_display_supports_selection_notification((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsShapes is a wrapper around the C function gdk_display_supports_shapes.
func (recv *Display) SupportsShapes() bool {
	retC := C.gdk_display_supports_shapes((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sync is a wrapper around the C function gdk_display_sync.
func (recv *Display) Sync() {
	C.gdk_display_sync((*C.GdkDisplay)(recv.native))

	return
}

// WarpPointer is a wrapper around the C function gdk_display_warp_pointer.
func (recv *Display) WarpPointer(screen *Screen, x int32, y int32) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_display_warp_pointer((*C.GdkDisplay)(recv.native), c_screen, c_x, c_y)

	return
}

// AddToGobjectClassGoTypeMap : GdkDisplayManager

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DisplayManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DisplayManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DisplayManager with another DisplayManager, and returns true if they represent the same GObject.
func (recv *DisplayManager) Equals(other *DisplayManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DisplayManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DisplayManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a DisplayManager.
func CastToDisplayManager(object *gobject.Object) *DisplayManager {
	return DisplayManagerNewFromC(object.ToC())
}

type signalDisplayManagerDisplayOpenedDetail struct {
	callback  DisplayManagerSignalDisplayOpenedCallback
	handlerID C.gulong
}

var signalDisplayManagerDisplayOpenedId int
var signalDisplayManagerDisplayOpenedMap = make(map[int]signalDisplayManagerDisplayOpenedDetail)
var signalDisplayManagerDisplayOpenedLock sync.RWMutex

// DisplayManagerSignalDisplayOpenedCallback is a callback function for a 'display-opened' signal emitted from a DisplayManager.
type DisplayManagerSignalDisplayOpenedCallback func(display *Display)

/*
ConnectDisplayOpened connects the callback to the 'display-opened' signal for the DisplayManager.

The returned value represents the connection, and may be passed to DisconnectDisplayOpened to remove it.
*/
func (recv *DisplayManager) ConnectDisplayOpened(callback DisplayManagerSignalDisplayOpenedCallback) int {
	signalDisplayManagerDisplayOpenedLock.Lock()
	defer signalDisplayManagerDisplayOpenedLock.Unlock()

	signalDisplayManagerDisplayOpenedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DisplayManager_signal_connect_display_opened(instance, C.gpointer(uintptr(signalDisplayManagerDisplayOpenedId)))

	detail := signalDisplayManagerDisplayOpenedDetail{callback, handlerID}
	signalDisplayManagerDisplayOpenedMap[signalDisplayManagerDisplayOpenedId] = detail

	return signalDisplayManagerDisplayOpenedId
}

/*
DisconnectDisplayOpened disconnects a callback from the 'display-opened' signal for the DisplayManager.

The connectionID should be a value returned from a call to ConnectDisplayOpened.
*/
func (recv *DisplayManager) DisconnectDisplayOpened(connectionID int) {
	signalDisplayManagerDisplayOpenedLock.Lock()
	defer signalDisplayManagerDisplayOpenedLock.Unlock()

	detail, exists := signalDisplayManagerDisplayOpenedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDisplayManagerDisplayOpenedMap, connectionID)
}

//export displaymanager_displayOpenedHandler
func displaymanager_displayOpenedHandler(_ *C.GObject, c_display *C.GdkDisplay, data C.gpointer) {
	signalDisplayManagerDisplayOpenedLock.RLock()
	defer signalDisplayManagerDisplayOpenedLock.RUnlock()

	display := DisplayNewFromC(unsafe.Pointer(c_display))

	index := int(uintptr(data))
	callback := signalDisplayManagerDisplayOpenedMap[index].callback
	callback(display)
}

// DisplayManagerGet is a wrapper around the C function gdk_display_manager_get.
func DisplayManagerGet() *DisplayManager {
	retC := C.gdk_display_manager_get()
	retGo := DisplayManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDefaultDisplay is a wrapper around the C function gdk_display_manager_get_default_display.
func (recv *DisplayManager) GetDefaultDisplay() *Display {
	retC := C.gdk_display_manager_get_default_display((*C.GdkDisplayManager)(recv.native))
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ListDisplays is a wrapper around the C function gdk_display_manager_list_displays.
func (recv *DisplayManager) ListDisplays() *glib.SList {
	retC := C.gdk_display_manager_list_displays((*C.GdkDisplayManager)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetDefaultDisplay is a wrapper around the C function gdk_display_manager_set_default_display.
func (recv *DisplayManager) SetDefaultDisplay(display *Display) {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	C.gdk_display_manager_set_default_display((*C.GdkDisplayManager)(recv.native), c_display)

	return
}

// AddToGobjectClassGoTypeMap : GdkDragContext

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DragContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DragContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DragContext with another DragContext, and returns true if they represent the same GObject.
func (recv *DragContext) Equals(other *DragContext) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DragContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DragContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a DragContext.
func CastToDragContext(object *gobject.Object) *DragContext {
	return DragContextNewFromC(object.ToC())
}

// GetDevice is a wrapper around the C function gdk_drag_context_get_device.
func (recv *DragContext) GetDevice() *Device {
	retC := C.gdk_drag_context_get_device((*C.GdkDragContext)(recv.native))
	retGo := DeviceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetDevice is a wrapper around the C function gdk_drag_context_set_device.
func (recv *DragContext) SetDevice(device *Device) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	C.gdk_drag_context_set_device((*C.GdkDragContext)(recv.native), c_device)

	return
}

// AddToGobjectClassGoTypeMap : GdkFrameClock

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FrameClock) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FrameClock) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FrameClock with another FrameClock, and returns true if they represent the same GObject.
func (recv *FrameClock) Equals(other *FrameClock) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FrameClock) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FrameClock.
// Exercise care, as this is a potentially dangerous function if the Object is not a FrameClock.
func CastToFrameClock(object *gobject.Object) *FrameClock {
	return FrameClockNewFromC(object.ToC())
}

type signalFrameClockAfterPaintDetail struct {
	callback  FrameClockSignalAfterPaintCallback
	handlerID C.gulong
}

var signalFrameClockAfterPaintId int
var signalFrameClockAfterPaintMap = make(map[int]signalFrameClockAfterPaintDetail)
var signalFrameClockAfterPaintLock sync.RWMutex

// FrameClockSignalAfterPaintCallback is a callback function for a 'after-paint' signal emitted from a FrameClock.
type FrameClockSignalAfterPaintCallback func()

/*
ConnectAfterPaint connects the callback to the 'after-paint' signal for the FrameClock.

The returned value represents the connection, and may be passed to DisconnectAfterPaint to remove it.
*/
func (recv *FrameClock) ConnectAfterPaint(callback FrameClockSignalAfterPaintCallback) int {
	signalFrameClockAfterPaintLock.Lock()
	defer signalFrameClockAfterPaintLock.Unlock()

	signalFrameClockAfterPaintId++
	instance := C.gpointer(recv.native)
	handlerID := C.FrameClock_signal_connect_after_paint(instance, C.gpointer(uintptr(signalFrameClockAfterPaintId)))

	detail := signalFrameClockAfterPaintDetail{callback, handlerID}
	signalFrameClockAfterPaintMap[signalFrameClockAfterPaintId] = detail

	return signalFrameClockAfterPaintId
}

/*
DisconnectAfterPaint disconnects a callback from the 'after-paint' signal for the FrameClock.

The connectionID should be a value returned from a call to ConnectAfterPaint.
*/
func (recv *FrameClock) DisconnectAfterPaint(connectionID int) {
	signalFrameClockAfterPaintLock.Lock()
	defer signalFrameClockAfterPaintLock.Unlock()

	detail, exists := signalFrameClockAfterPaintMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFrameClockAfterPaintMap, connectionID)
}

//export frameclock_afterPaintHandler
func frameclock_afterPaintHandler(_ *C.GObject, data C.gpointer) {
	signalFrameClockAfterPaintLock.RLock()
	defer signalFrameClockAfterPaintLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFrameClockAfterPaintMap[index].callback
	callback()
}

type signalFrameClockBeforePaintDetail struct {
	callback  FrameClockSignalBeforePaintCallback
	handlerID C.gulong
}

var signalFrameClockBeforePaintId int
var signalFrameClockBeforePaintMap = make(map[int]signalFrameClockBeforePaintDetail)
var signalFrameClockBeforePaintLock sync.RWMutex

// FrameClockSignalBeforePaintCallback is a callback function for a 'before-paint' signal emitted from a FrameClock.
type FrameClockSignalBeforePaintCallback func()

/*
ConnectBeforePaint connects the callback to the 'before-paint' signal for the FrameClock.

The returned value represents the connection, and may be passed to DisconnectBeforePaint to remove it.
*/
func (recv *FrameClock) ConnectBeforePaint(callback FrameClockSignalBeforePaintCallback) int {
	signalFrameClockBeforePaintLock.Lock()
	defer signalFrameClockBeforePaintLock.Unlock()

	signalFrameClockBeforePaintId++
	instance := C.gpointer(recv.native)
	handlerID := C.FrameClock_signal_connect_before_paint(instance, C.gpointer(uintptr(signalFrameClockBeforePaintId)))

	detail := signalFrameClockBeforePaintDetail{callback, handlerID}
	signalFrameClockBeforePaintMap[signalFrameClockBeforePaintId] = detail

	return signalFrameClockBeforePaintId
}

/*
DisconnectBeforePaint disconnects a callback from the 'before-paint' signal for the FrameClock.

The connectionID should be a value returned from a call to ConnectBeforePaint.
*/
func (recv *FrameClock) DisconnectBeforePaint(connectionID int) {
	signalFrameClockBeforePaintLock.Lock()
	defer signalFrameClockBeforePaintLock.Unlock()

	detail, exists := signalFrameClockBeforePaintMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFrameClockBeforePaintMap, connectionID)
}

//export frameclock_beforePaintHandler
func frameclock_beforePaintHandler(_ *C.GObject, data C.gpointer) {
	signalFrameClockBeforePaintLock.RLock()
	defer signalFrameClockBeforePaintLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFrameClockBeforePaintMap[index].callback
	callback()
}

type signalFrameClockFlushEventsDetail struct {
	callback  FrameClockSignalFlushEventsCallback
	handlerID C.gulong
}

var signalFrameClockFlushEventsId int
var signalFrameClockFlushEventsMap = make(map[int]signalFrameClockFlushEventsDetail)
var signalFrameClockFlushEventsLock sync.RWMutex

// FrameClockSignalFlushEventsCallback is a callback function for a 'flush-events' signal emitted from a FrameClock.
type FrameClockSignalFlushEventsCallback func()

/*
ConnectFlushEvents connects the callback to the 'flush-events' signal for the FrameClock.

The returned value represents the connection, and may be passed to DisconnectFlushEvents to remove it.
*/
func (recv *FrameClock) ConnectFlushEvents(callback FrameClockSignalFlushEventsCallback) int {
	signalFrameClockFlushEventsLock.Lock()
	defer signalFrameClockFlushEventsLock.Unlock()

	signalFrameClockFlushEventsId++
	instance := C.gpointer(recv.native)
	handlerID := C.FrameClock_signal_connect_flush_events(instance, C.gpointer(uintptr(signalFrameClockFlushEventsId)))

	detail := signalFrameClockFlushEventsDetail{callback, handlerID}
	signalFrameClockFlushEventsMap[signalFrameClockFlushEventsId] = detail

	return signalFrameClockFlushEventsId
}

/*
DisconnectFlushEvents disconnects a callback from the 'flush-events' signal for the FrameClock.

The connectionID should be a value returned from a call to ConnectFlushEvents.
*/
func (recv *FrameClock) DisconnectFlushEvents(connectionID int) {
	signalFrameClockFlushEventsLock.Lock()
	defer signalFrameClockFlushEventsLock.Unlock()

	detail, exists := signalFrameClockFlushEventsMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFrameClockFlushEventsMap, connectionID)
}

//export frameclock_flushEventsHandler
func frameclock_flushEventsHandler(_ *C.GObject, data C.gpointer) {
	signalFrameClockFlushEventsLock.RLock()
	defer signalFrameClockFlushEventsLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFrameClockFlushEventsMap[index].callback
	callback()
}

type signalFrameClockLayoutDetail struct {
	callback  FrameClockSignalLayoutCallback
	handlerID C.gulong
}

var signalFrameClockLayoutId int
var signalFrameClockLayoutMap = make(map[int]signalFrameClockLayoutDetail)
var signalFrameClockLayoutLock sync.RWMutex

// FrameClockSignalLayoutCallback is a callback function for a 'layout' signal emitted from a FrameClock.
type FrameClockSignalLayoutCallback func()

/*
ConnectLayout connects the callback to the 'layout' signal for the FrameClock.

The returned value represents the connection, and may be passed to DisconnectLayout to remove it.
*/
func (recv *FrameClock) ConnectLayout(callback FrameClockSignalLayoutCallback) int {
	signalFrameClockLayoutLock.Lock()
	defer signalFrameClockLayoutLock.Unlock()

	signalFrameClockLayoutId++
	instance := C.gpointer(recv.native)
	handlerID := C.FrameClock_signal_connect_layout(instance, C.gpointer(uintptr(signalFrameClockLayoutId)))

	detail := signalFrameClockLayoutDetail{callback, handlerID}
	signalFrameClockLayoutMap[signalFrameClockLayoutId] = detail

	return signalFrameClockLayoutId
}

/*
DisconnectLayout disconnects a callback from the 'layout' signal for the FrameClock.

The connectionID should be a value returned from a call to ConnectLayout.
*/
func (recv *FrameClock) DisconnectLayout(connectionID int) {
	signalFrameClockLayoutLock.Lock()
	defer signalFrameClockLayoutLock.Unlock()

	detail, exists := signalFrameClockLayoutMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFrameClockLayoutMap, connectionID)
}

//export frameclock_layoutHandler
func frameclock_layoutHandler(_ *C.GObject, data C.gpointer) {
	signalFrameClockLayoutLock.RLock()
	defer signalFrameClockLayoutLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFrameClockLayoutMap[index].callback
	callback()
}

type signalFrameClockPaintDetail struct {
	callback  FrameClockSignalPaintCallback
	handlerID C.gulong
}

var signalFrameClockPaintId int
var signalFrameClockPaintMap = make(map[int]signalFrameClockPaintDetail)
var signalFrameClockPaintLock sync.RWMutex

// FrameClockSignalPaintCallback is a callback function for a 'paint' signal emitted from a FrameClock.
type FrameClockSignalPaintCallback func()

/*
ConnectPaint connects the callback to the 'paint' signal for the FrameClock.

The returned value represents the connection, and may be passed to DisconnectPaint to remove it.
*/
func (recv *FrameClock) ConnectPaint(callback FrameClockSignalPaintCallback) int {
	signalFrameClockPaintLock.Lock()
	defer signalFrameClockPaintLock.Unlock()

	signalFrameClockPaintId++
	instance := C.gpointer(recv.native)
	handlerID := C.FrameClock_signal_connect_paint(instance, C.gpointer(uintptr(signalFrameClockPaintId)))

	detail := signalFrameClockPaintDetail{callback, handlerID}
	signalFrameClockPaintMap[signalFrameClockPaintId] = detail

	return signalFrameClockPaintId
}

/*
DisconnectPaint disconnects a callback from the 'paint' signal for the FrameClock.

The connectionID should be a value returned from a call to ConnectPaint.
*/
func (recv *FrameClock) DisconnectPaint(connectionID int) {
	signalFrameClockPaintLock.Lock()
	defer signalFrameClockPaintLock.Unlock()

	detail, exists := signalFrameClockPaintMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFrameClockPaintMap, connectionID)
}

//export frameclock_paintHandler
func frameclock_paintHandler(_ *C.GObject, data C.gpointer) {
	signalFrameClockPaintLock.RLock()
	defer signalFrameClockPaintLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFrameClockPaintMap[index].callback
	callback()
}

type signalFrameClockResumeEventsDetail struct {
	callback  FrameClockSignalResumeEventsCallback
	handlerID C.gulong
}

var signalFrameClockResumeEventsId int
var signalFrameClockResumeEventsMap = make(map[int]signalFrameClockResumeEventsDetail)
var signalFrameClockResumeEventsLock sync.RWMutex

// FrameClockSignalResumeEventsCallback is a callback function for a 'resume-events' signal emitted from a FrameClock.
type FrameClockSignalResumeEventsCallback func()

/*
ConnectResumeEvents connects the callback to the 'resume-events' signal for the FrameClock.

The returned value represents the connection, and may be passed to DisconnectResumeEvents to remove it.
*/
func (recv *FrameClock) ConnectResumeEvents(callback FrameClockSignalResumeEventsCallback) int {
	signalFrameClockResumeEventsLock.Lock()
	defer signalFrameClockResumeEventsLock.Unlock()

	signalFrameClockResumeEventsId++
	instance := C.gpointer(recv.native)
	handlerID := C.FrameClock_signal_connect_resume_events(instance, C.gpointer(uintptr(signalFrameClockResumeEventsId)))

	detail := signalFrameClockResumeEventsDetail{callback, handlerID}
	signalFrameClockResumeEventsMap[signalFrameClockResumeEventsId] = detail

	return signalFrameClockResumeEventsId
}

/*
DisconnectResumeEvents disconnects a callback from the 'resume-events' signal for the FrameClock.

The connectionID should be a value returned from a call to ConnectResumeEvents.
*/
func (recv *FrameClock) DisconnectResumeEvents(connectionID int) {
	signalFrameClockResumeEventsLock.Lock()
	defer signalFrameClockResumeEventsLock.Unlock()

	detail, exists := signalFrameClockResumeEventsMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFrameClockResumeEventsMap, connectionID)
}

//export frameclock_resumeEventsHandler
func frameclock_resumeEventsHandler(_ *C.GObject, data C.gpointer) {
	signalFrameClockResumeEventsLock.RLock()
	defer signalFrameClockResumeEventsLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFrameClockResumeEventsMap[index].callback
	callback()
}

type signalFrameClockUpdateDetail struct {
	callback  FrameClockSignalUpdateCallback
	handlerID C.gulong
}

var signalFrameClockUpdateId int
var signalFrameClockUpdateMap = make(map[int]signalFrameClockUpdateDetail)
var signalFrameClockUpdateLock sync.RWMutex

// FrameClockSignalUpdateCallback is a callback function for a 'update' signal emitted from a FrameClock.
type FrameClockSignalUpdateCallback func()

/*
ConnectUpdate connects the callback to the 'update' signal for the FrameClock.

The returned value represents the connection, and may be passed to DisconnectUpdate to remove it.
*/
func (recv *FrameClock) ConnectUpdate(callback FrameClockSignalUpdateCallback) int {
	signalFrameClockUpdateLock.Lock()
	defer signalFrameClockUpdateLock.Unlock()

	signalFrameClockUpdateId++
	instance := C.gpointer(recv.native)
	handlerID := C.FrameClock_signal_connect_update(instance, C.gpointer(uintptr(signalFrameClockUpdateId)))

	detail := signalFrameClockUpdateDetail{callback, handlerID}
	signalFrameClockUpdateMap[signalFrameClockUpdateId] = detail

	return signalFrameClockUpdateId
}

/*
DisconnectUpdate disconnects a callback from the 'update' signal for the FrameClock.

The connectionID should be a value returned from a call to ConnectUpdate.
*/
func (recv *FrameClock) DisconnectUpdate(connectionID int) {
	signalFrameClockUpdateLock.Lock()
	defer signalFrameClockUpdateLock.Unlock()

	detail, exists := signalFrameClockUpdateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFrameClockUpdateMap, connectionID)
}

//export frameclock_updateHandler
func frameclock_updateHandler(_ *C.GObject, data C.gpointer) {
	signalFrameClockUpdateLock.RLock()
	defer signalFrameClockUpdateLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFrameClockUpdateMap[index].callback
	callback()
}

// AddToGobjectClassGoTypeMap : GdkGLContext

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GLContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GLContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GLContext with another GLContext, and returns true if they represent the same GObject.
func (recv *GLContext) Equals(other *GLContext) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *GLContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to GLContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a GLContext.
func CastToGLContext(object *gobject.Object) *GLContext {
	return GLContextNewFromC(object.ToC())
}

// AddToGobjectClassGoTypeMap : GdkKeymap

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Keymap) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Keymap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Keymap with another Keymap, and returns true if they represent the same GObject.
func (recv *Keymap) Equals(other *Keymap) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Keymap) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Keymap.
// Exercise care, as this is a potentially dangerous function if the Object is not a Keymap.
func CastToKeymap(object *gobject.Object) *Keymap {
	return KeymapNewFromC(object.ToC())
}

type signalKeymapDirectionChangedDetail struct {
	callback  KeymapSignalDirectionChangedCallback
	handlerID C.gulong
}

var signalKeymapDirectionChangedId int
var signalKeymapDirectionChangedMap = make(map[int]signalKeymapDirectionChangedDetail)
var signalKeymapDirectionChangedLock sync.RWMutex

// KeymapSignalDirectionChangedCallback is a callback function for a 'direction-changed' signal emitted from a Keymap.
type KeymapSignalDirectionChangedCallback func()

/*
ConnectDirectionChanged connects the callback to the 'direction-changed' signal for the Keymap.

The returned value represents the connection, and may be passed to DisconnectDirectionChanged to remove it.
*/
func (recv *Keymap) ConnectDirectionChanged(callback KeymapSignalDirectionChangedCallback) int {
	signalKeymapDirectionChangedLock.Lock()
	defer signalKeymapDirectionChangedLock.Unlock()

	signalKeymapDirectionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Keymap_signal_connect_direction_changed(instance, C.gpointer(uintptr(signalKeymapDirectionChangedId)))

	detail := signalKeymapDirectionChangedDetail{callback, handlerID}
	signalKeymapDirectionChangedMap[signalKeymapDirectionChangedId] = detail

	return signalKeymapDirectionChangedId
}

/*
DisconnectDirectionChanged disconnects a callback from the 'direction-changed' signal for the Keymap.

The connectionID should be a value returned from a call to ConnectDirectionChanged.
*/
func (recv *Keymap) DisconnectDirectionChanged(connectionID int) {
	signalKeymapDirectionChangedLock.Lock()
	defer signalKeymapDirectionChangedLock.Unlock()

	detail, exists := signalKeymapDirectionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalKeymapDirectionChangedMap, connectionID)
}

//export keymap_directionChangedHandler
func keymap_directionChangedHandler(_ *C.GObject, data C.gpointer) {
	signalKeymapDirectionChangedLock.RLock()
	defer signalKeymapDirectionChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalKeymapDirectionChangedMap[index].callback
	callback()
}

type signalKeymapKeysChangedDetail struct {
	callback  KeymapSignalKeysChangedCallback
	handlerID C.gulong
}

var signalKeymapKeysChangedId int
var signalKeymapKeysChangedMap = make(map[int]signalKeymapKeysChangedDetail)
var signalKeymapKeysChangedLock sync.RWMutex

// KeymapSignalKeysChangedCallback is a callback function for a 'keys-changed' signal emitted from a Keymap.
type KeymapSignalKeysChangedCallback func()

/*
ConnectKeysChanged connects the callback to the 'keys-changed' signal for the Keymap.

The returned value represents the connection, and may be passed to DisconnectKeysChanged to remove it.
*/
func (recv *Keymap) ConnectKeysChanged(callback KeymapSignalKeysChangedCallback) int {
	signalKeymapKeysChangedLock.Lock()
	defer signalKeymapKeysChangedLock.Unlock()

	signalKeymapKeysChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Keymap_signal_connect_keys_changed(instance, C.gpointer(uintptr(signalKeymapKeysChangedId)))

	detail := signalKeymapKeysChangedDetail{callback, handlerID}
	signalKeymapKeysChangedMap[signalKeymapKeysChangedId] = detail

	return signalKeymapKeysChangedId
}

/*
DisconnectKeysChanged disconnects a callback from the 'keys-changed' signal for the Keymap.

The connectionID should be a value returned from a call to ConnectKeysChanged.
*/
func (recv *Keymap) DisconnectKeysChanged(connectionID int) {
	signalKeymapKeysChangedLock.Lock()
	defer signalKeymapKeysChangedLock.Unlock()

	detail, exists := signalKeymapKeysChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalKeymapKeysChangedMap, connectionID)
}

//export keymap_keysChangedHandler
func keymap_keysChangedHandler(_ *C.GObject, data C.gpointer) {
	signalKeymapKeysChangedLock.RLock()
	defer signalKeymapKeysChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalKeymapKeysChangedMap[index].callback
	callback()
}

// KeymapGetDefault is a wrapper around the C function gdk_keymap_get_default.
func KeymapGetDefault() *Keymap {
	retC := C.gdk_keymap_get_default()
	retGo := KeymapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// KeymapGetForDisplay is a wrapper around the C function gdk_keymap_get_for_display.
func KeymapGetForDisplay(display *Display) *Keymap {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	retC := C.gdk_keymap_get_for_display(c_display)
	retGo := KeymapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDirection is a wrapper around the C function gdk_keymap_get_direction.
func (recv *Keymap) GetDirection() pango.Direction {
	retC := C.gdk_keymap_get_direction((*C.GdkKeymap)(recv.native))
	retGo := (pango.Direction)(retC)

	return retGo
}

// Unsupported : gdk_keymap_get_entries_for_keycode : unsupported parameter keys : output array param keys

// Unsupported : gdk_keymap_get_entries_for_keyval : unsupported parameter keys : output array param keys

// LookupKey is a wrapper around the C function gdk_keymap_lookup_key.
func (recv *Keymap) LookupKey(key *KeymapKey) uint32 {
	c_key := (*C.GdkKeymapKey)(C.NULL)
	if key != nil {
		c_key = (*C.GdkKeymapKey)(key.ToC())
	}

	retC := C.gdk_keymap_lookup_key((*C.GdkKeymap)(recv.native), c_key)
	retGo := (uint32)(retC)

	return retGo
}

// TranslateKeyboardState is a wrapper around the C function gdk_keymap_translate_keyboard_state.
func (recv *Keymap) TranslateKeyboardState(hardwareKeycode uint32, state ModifierType, group int32) (bool, uint32, int32, int32, ModifierType) {
	c_hardware_keycode := (C.guint)(hardwareKeycode)

	c_state := (C.GdkModifierType)(state)

	c_group := (C.gint)(group)

	var c_keyval C.guint

	var c_effective_group C.gint

	var c_level C.gint

	var c_consumed_modifiers C.GdkModifierType

	retC := C.gdk_keymap_translate_keyboard_state((*C.GdkKeymap)(recv.native), c_hardware_keycode, c_state, c_group, &c_keyval, &c_effective_group, &c_level, &c_consumed_modifiers)
	retGo := retC == C.TRUE

	keyval := (uint32)(c_keyval)

	effectiveGroup := (int32)(c_effective_group)

	level := (int32)(c_level)

	consumedModifiers := (ModifierType)(c_consumed_modifiers)

	return retGo, keyval, effectiveGroup, level, consumedModifiers
}

// AddToGobjectClassGoTypeMap : GdkScreen

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Screen) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Screen) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Screen with another Screen, and returns true if they represent the same GObject.
func (recv *Screen) Equals(other *Screen) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Screen) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Screen.
// Exercise care, as this is a potentially dangerous function if the Object is not a Screen.
func CastToScreen(object *gobject.Object) *Screen {
	return ScreenNewFromC(object.ToC())
}

type signalScreenCompositedChangedDetail struct {
	callback  ScreenSignalCompositedChangedCallback
	handlerID C.gulong
}

var signalScreenCompositedChangedId int
var signalScreenCompositedChangedMap = make(map[int]signalScreenCompositedChangedDetail)
var signalScreenCompositedChangedLock sync.RWMutex

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
	signalScreenCompositedChangedLock.RLock()
	defer signalScreenCompositedChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalScreenCompositedChangedMap[index].callback
	callback()
}

type signalScreenSizeChangedDetail struct {
	callback  ScreenSignalSizeChangedCallback
	handlerID C.gulong
}

var signalScreenSizeChangedId int
var signalScreenSizeChangedMap = make(map[int]signalScreenSizeChangedDetail)
var signalScreenSizeChangedLock sync.RWMutex

// ScreenSignalSizeChangedCallback is a callback function for a 'size-changed' signal emitted from a Screen.
type ScreenSignalSizeChangedCallback func()

/*
ConnectSizeChanged connects the callback to the 'size-changed' signal for the Screen.

The returned value represents the connection, and may be passed to DisconnectSizeChanged to remove it.
*/
func (recv *Screen) ConnectSizeChanged(callback ScreenSignalSizeChangedCallback) int {
	signalScreenSizeChangedLock.Lock()
	defer signalScreenSizeChangedLock.Unlock()

	signalScreenSizeChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Screen_signal_connect_size_changed(instance, C.gpointer(uintptr(signalScreenSizeChangedId)))

	detail := signalScreenSizeChangedDetail{callback, handlerID}
	signalScreenSizeChangedMap[signalScreenSizeChangedId] = detail

	return signalScreenSizeChangedId
}

/*
DisconnectSizeChanged disconnects a callback from the 'size-changed' signal for the Screen.

The connectionID should be a value returned from a call to ConnectSizeChanged.
*/
func (recv *Screen) DisconnectSizeChanged(connectionID int) {
	signalScreenSizeChangedLock.Lock()
	defer signalScreenSizeChangedLock.Unlock()

	detail, exists := signalScreenSizeChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScreenSizeChangedMap, connectionID)
}

//export screen_sizeChangedHandler
func screen_sizeChangedHandler(_ *C.GObject, data C.gpointer) {
	signalScreenSizeChangedLock.RLock()
	defer signalScreenSizeChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalScreenSizeChangedMap[index].callback
	callback()
}

// ScreenGetDefault is a wrapper around the C function gdk_screen_get_default.
func ScreenGetDefault() *Screen {
	retC := C.gdk_screen_get_default()
	var retGo (*Screen)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ScreenNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ScreenHeight is a wrapper around the C function gdk_screen_height.
func ScreenHeight() int32 {
	retC := C.gdk_screen_height()
	retGo := (int32)(retC)

	return retGo
}

// ScreenHeightMm is a wrapper around the C function gdk_screen_height_mm.
func ScreenHeightMm() int32 {
	retC := C.gdk_screen_height_mm()
	retGo := (int32)(retC)

	return retGo
}

// ScreenWidth is a wrapper around the C function gdk_screen_width.
func ScreenWidth() int32 {
	retC := C.gdk_screen_width()
	retGo := (int32)(retC)

	return retGo
}

// ScreenWidthMm is a wrapper around the C function gdk_screen_width_mm.
func ScreenWidthMm() int32 {
	retC := C.gdk_screen_width_mm()
	retGo := (int32)(retC)

	return retGo
}

// GetActiveWindow is a wrapper around the C function gdk_screen_get_active_window.
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

// GetDisplay is a wrapper around the C function gdk_screen_get_display.
func (recv *Screen) GetDisplay() *Display {
	retC := C.gdk_screen_get_display((*C.GdkScreen)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFontOptions is a wrapper around the C function gdk_screen_get_font_options.
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

// GetHeight is a wrapper around the C function gdk_screen_get_height.
func (recv *Screen) GetHeight() int32 {
	retC := C.gdk_screen_get_height((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetHeightMm is a wrapper around the C function gdk_screen_get_height_mm.
func (recv *Screen) GetHeightMm() int32 {
	retC := C.gdk_screen_get_height_mm((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMonitorAtPoint is a wrapper around the C function gdk_screen_get_monitor_at_point.
func (recv *Screen) GetMonitorAtPoint(x int32, y int32) int32 {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gdk_screen_get_monitor_at_point((*C.GdkScreen)(recv.native), c_x, c_y)
	retGo := (int32)(retC)

	return retGo
}

// GetMonitorAtWindow is a wrapper around the C function gdk_screen_get_monitor_at_window.
func (recv *Screen) GetMonitorAtWindow(window *Window) int32 {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gdk_screen_get_monitor_at_window((*C.GdkScreen)(recv.native), c_window)
	retGo := (int32)(retC)

	return retGo
}

// GetMonitorGeometry is a wrapper around the C function gdk_screen_get_monitor_geometry.
func (recv *Screen) GetMonitorGeometry(monitorNum int32) *Rectangle {
	c_monitor_num := (C.gint)(monitorNum)

	var c_dest C.GdkRectangle

	C.gdk_screen_get_monitor_geometry((*C.GdkScreen)(recv.native), c_monitor_num, &c_dest)

	dest := RectangleNewFromC(unsafe.Pointer(&c_dest))

	return dest
}

// GetNMonitors is a wrapper around the C function gdk_screen_get_n_monitors.
func (recv *Screen) GetNMonitors() int32 {
	retC := C.gdk_screen_get_n_monitors((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNumber is a wrapper around the C function gdk_screen_get_number.
func (recv *Screen) GetNumber() int32 {
	retC := C.gdk_screen_get_number((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetResolution is a wrapper around the C function gdk_screen_get_resolution.
func (recv *Screen) GetResolution() float64 {
	retC := C.gdk_screen_get_resolution((*C.GdkScreen)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetRgbaVisual is a wrapper around the C function gdk_screen_get_rgba_visual.
func (recv *Screen) GetRgbaVisual() *Visual {
	retC := C.gdk_screen_get_rgba_visual((*C.GdkScreen)(recv.native))
	var retGo (*Visual)
	if retC == nil {
		retGo = nil
	} else {
		retGo = VisualNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetRootWindow is a wrapper around the C function gdk_screen_get_root_window.
func (recv *Screen) GetRootWindow() *Window {
	retC := C.gdk_screen_get_root_window((*C.GdkScreen)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSetting is a wrapper around the C function gdk_screen_get_setting.
func (recv *Screen) GetSetting(name string, value *gobject.Value) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.gdk_screen_get_setting((*C.GdkScreen)(recv.native), c_name, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// GetSystemVisual is a wrapper around the C function gdk_screen_get_system_visual.
func (recv *Screen) GetSystemVisual() *Visual {
	retC := C.gdk_screen_get_system_visual((*C.GdkScreen)(recv.native))
	retGo := VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetToplevelWindows is a wrapper around the C function gdk_screen_get_toplevel_windows.
func (recv *Screen) GetToplevelWindows() *glib.List {
	retC := C.gdk_screen_get_toplevel_windows((*C.GdkScreen)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidth is a wrapper around the C function gdk_screen_get_width.
func (recv *Screen) GetWidth() int32 {
	retC := C.gdk_screen_get_width((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWidthMm is a wrapper around the C function gdk_screen_get_width_mm.
func (recv *Screen) GetWidthMm() int32 {
	retC := C.gdk_screen_get_width_mm((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWindowStack is a wrapper around the C function gdk_screen_get_window_stack.
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

// IsComposited is a wrapper around the C function gdk_screen_is_composited.
func (recv *Screen) IsComposited() bool {
	retC := C.gdk_screen_is_composited((*C.GdkScreen)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ListVisuals is a wrapper around the C function gdk_screen_list_visuals.
func (recv *Screen) ListVisuals() *glib.List {
	retC := C.gdk_screen_list_visuals((*C.GdkScreen)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MakeDisplayName is a wrapper around the C function gdk_screen_make_display_name.
func (recv *Screen) MakeDisplayName() string {
	retC := C.gdk_screen_make_display_name((*C.GdkScreen)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// SetFontOptions is a wrapper around the C function gdk_screen_set_font_options.
func (recv *Screen) SetFontOptions(options *cairo.FontOptions) {
	c_options := (*C.cairo_font_options_t)(C.NULL)
	if options != nil {
		c_options = (*C.cairo_font_options_t)(options.ToC())
	}

	C.gdk_screen_set_font_options((*C.GdkScreen)(recv.native), c_options)

	return
}

// SetResolution is a wrapper around the C function gdk_screen_set_resolution.
func (recv *Screen) SetResolution(dpi float64) {
	c_dpi := (C.gdouble)(dpi)

	C.gdk_screen_set_resolution((*C.GdkScreen)(recv.native), c_dpi)

	return
}

// AddToGobjectClassGoTypeMap : GdkVisual

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Visual) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Visual) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Visual with another Visual, and returns true if they represent the same GObject.
func (recv *Visual) Equals(other *Visual) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Visual) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Visual.
// Exercise care, as this is a potentially dangerous function if the Object is not a Visual.
func CastToVisual(object *gobject.Object) *Visual {
	return VisualNewFromC(object.ToC())
}

// VisualGetBest is a wrapper around the C function gdk_visual_get_best.
func VisualGetBest() *Visual {
	retC := C.gdk_visual_get_best()
	retGo := VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VisualGetBestDepth is a wrapper around the C function gdk_visual_get_best_depth.
func VisualGetBestDepth() int32 {
	retC := C.gdk_visual_get_best_depth()
	retGo := (int32)(retC)

	return retGo
}

// VisualGetBestType is a wrapper around the C function gdk_visual_get_best_type.
func VisualGetBestType() VisualType {
	retC := C.gdk_visual_get_best_type()
	retGo := (VisualType)(retC)

	return retGo
}

// VisualGetBestWithBoth is a wrapper around the C function gdk_visual_get_best_with_both.
func VisualGetBestWithBoth(depth int32, visualType VisualType) *Visual {
	c_depth := (C.gint)(depth)

	c_visual_type := (C.GdkVisualType)(visualType)

	retC := C.gdk_visual_get_best_with_both(c_depth, c_visual_type)
	var retGo (*Visual)
	if retC == nil {
		retGo = nil
	} else {
		retGo = VisualNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// VisualGetBestWithDepth is a wrapper around the C function gdk_visual_get_best_with_depth.
func VisualGetBestWithDepth(depth int32) *Visual {
	c_depth := (C.gint)(depth)

	retC := C.gdk_visual_get_best_with_depth(c_depth)
	retGo := VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VisualGetBestWithType is a wrapper around the C function gdk_visual_get_best_with_type.
func VisualGetBestWithType(visualType VisualType) *Visual {
	c_visual_type := (C.GdkVisualType)(visualType)

	retC := C.gdk_visual_get_best_with_type(c_visual_type)
	retGo := VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VisualGetSystem is a wrapper around the C function gdk_visual_get_system.
func VisualGetSystem() *Visual {
	retC := C.gdk_visual_get_system()
	retGo := VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gdk_visual_get_screen.
func (recv *Visual) GetScreen() *Screen {
	retC := C.gdk_visual_get_screen((*C.GdkVisual)(recv.native))
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddToGobjectClassGoTypeMap : GdkWindow

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Window) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Window with another Window, and returns true if they represent the same GObject.
func (recv *Window) Equals(other *Window) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Window) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Window.
// Exercise care, as this is a potentially dangerous function if the Object is not a Window.
func CastToWindow(object *gobject.Object) *Window {
	return WindowNewFromC(object.ToC())
}

// Unsupported signal 'from-embedder' for Window : unsupported parameter offscreen_x : direction is 'out'

// Unsupported signal 'moved-to-rect' for Window : param flipped_rect : gpointer

// Unsupported signal 'to-embedder' for Window : unsupported parameter embedder_x : direction is 'out'

// WindowNew is a wrapper around the C function gdk_window_new.
func WindowNew(parent *Window, attributes *WindowAttr, attributesMask WindowAttributesType) *Window {
	c_parent := (*C.GdkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GdkWindow)(parent.ToC())
	}

	c_attributes := (*C.GdkWindowAttr)(C.NULL)
	if attributes != nil {
		c_attributes = (*C.GdkWindowAttr)(attributes.ToC())
	}

	c_attributes_mask := (C.gint)(attributesMask)

	retC := C.gdk_window_new(c_parent, c_attributes, c_attributes_mask)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// WindowAtPointer is a wrapper around the C function gdk_window_at_pointer.
func WindowAtPointer() (*Window, int32, int32) {
	var c_win_x C.gint

	var c_win_y C.gint

	retC := C.gdk_window_at_pointer(&c_win_x, &c_win_y)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	winX := (int32)(c_win_x)

	winY := (int32)(c_win_y)

	return retGo, winX, winY
}

// WindowConstrainSize is a wrapper around the C function gdk_window_constrain_size.
func WindowConstrainSize(geometry *Geometry, flags WindowHints, width int32, height int32) (int32, int32) {
	c_geometry := (*C.GdkGeometry)(C.NULL)
	if geometry != nil {
		c_geometry = (*C.GdkGeometry)(geometry.ToC())
	}

	c_flags := (C.GdkWindowHints)(flags)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	var c_new_width C.gint

	var c_new_height C.gint

	C.gdk_window_constrain_size(c_geometry, c_flags, c_width, c_height, &c_new_width, &c_new_height)

	newWidth := (int32)(c_new_width)

	newHeight := (int32)(c_new_height)

	return newWidth, newHeight
}

// WindowProcessAllUpdates is a wrapper around the C function gdk_window_process_all_updates.
func WindowProcessAllUpdates() {
	C.gdk_window_process_all_updates()

	return
}

// WindowSetDebugUpdates is a wrapper around the C function gdk_window_set_debug_updates.
func WindowSetDebugUpdates(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gdk_window_set_debug_updates(c_setting)

	return
}

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc (GdkFilterFunc) for param function

// BeginMoveDrag is a wrapper around the C function gdk_window_begin_move_drag.
func (recv *Window) BeginMoveDrag(button int32, rootX int32, rootY int32, timestamp uint32) {
	c_button := (C.gint)(button)

	c_root_x := (C.gint)(rootX)

	c_root_y := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_begin_move_drag((*C.GdkWindow)(recv.native), c_button, c_root_x, c_root_y, c_timestamp)

	return
}

// BeginPaintRect is a wrapper around the C function gdk_window_begin_paint_rect.
func (recv *Window) BeginPaintRect(rectangle *Rectangle) {
	c_rectangle := (*C.GdkRectangle)(C.NULL)
	if rectangle != nil {
		c_rectangle = (*C.GdkRectangle)(rectangle.ToC())
	}

	C.gdk_window_begin_paint_rect((*C.GdkWindow)(recv.native), c_rectangle)

	return
}

// BeginPaintRegion is a wrapper around the C function gdk_window_begin_paint_region.
func (recv *Window) BeginPaintRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

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

// ConfigureFinished is a wrapper around the C function gdk_window_configure_finished.
func (recv *Window) ConfigureFinished() {
	C.gdk_window_configure_finished((*C.GdkWindow)(recv.native))

	return
}

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

// Blacklisted : gdk_window_destroy_notify

// EnableSynchronizedConfigure is a wrapper around the C function gdk_window_enable_synchronized_configure.
func (recv *Window) EnableSynchronizedConfigure() {
	C.gdk_window_enable_synchronized_configure((*C.GdkWindow)(recv.native))

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

// Fullscreen is a wrapper around the C function gdk_window_fullscreen.
func (recv *Window) Fullscreen() {
	C.gdk_window_fullscreen((*C.GdkWindow)(recv.native))

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

// GetDecorations is a wrapper around the C function gdk_window_get_decorations.
func (recv *Window) GetDecorations() (bool, WMDecoration) {
	var c_decorations C.GdkWMDecoration

	retC := C.gdk_window_get_decorations((*C.GdkWindow)(recv.native), &c_decorations)
	retGo := retC == C.TRUE

	decorations := (WMDecoration)(c_decorations)

	return retGo, decorations
}

// GetEvents is a wrapper around the C function gdk_window_get_events.
func (recv *Window) GetEvents() EventMask {
	retC := C.gdk_window_get_events((*C.GdkWindow)(recv.native))
	retGo := (EventMask)(retC)

	return retGo
}

// GetFrameExtents is a wrapper around the C function gdk_window_get_frame_extents.
func (recv *Window) GetFrameExtents() *Rectangle {
	var c_rect C.GdkRectangle

	C.gdk_window_get_frame_extents((*C.GdkWindow)(recv.native), &c_rect)

	rect := RectangleNewFromC(unsafe.Pointer(&c_rect))

	return rect
}

// GetGeometry is a wrapper around the C function gdk_window_get_geometry.
func (recv *Window) GetGeometry() (int32, int32, int32, int32) {
	var c_x C.gint

	var c_y C.gint

	var c_width C.gint

	var c_height C.gint

	C.gdk_window_get_geometry((*C.GdkWindow)(recv.native), &c_x, &c_y, &c_width, &c_height)

	x := (int32)(c_x)

	y := (int32)(c_y)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return x, y, width, height
}

// GetGroup is a wrapper around the C function gdk_window_get_group.
func (recv *Window) GetGroup() *Window {
	retC := C.gdk_window_get_group((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOrigin is a wrapper around the C function gdk_window_get_origin.
func (recv *Window) GetOrigin() (int32, int32, int32) {
	var c_x C.gint

	var c_y C.gint

	retC := C.gdk_window_get_origin((*C.GdkWindow)(recv.native), &c_x, &c_y)
	retGo := (int32)(retC)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return retGo, x, y
}

// GetParent is a wrapper around the C function gdk_window_get_parent.
func (recv *Window) GetParent() *Window {
	retC := C.gdk_window_get_parent((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPointer is a wrapper around the C function gdk_window_get_pointer.
func (recv *Window) GetPointer() (*Window, int32, int32, ModifierType) {
	var c_x C.gint

	var c_y C.gint

	var c_mask C.GdkModifierType

	retC := C.gdk_window_get_pointer((*C.GdkWindow)(recv.native), &c_x, &c_y, &c_mask)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	x := (int32)(c_x)

	y := (int32)(c_y)

	mask := (ModifierType)(c_mask)

	return retGo, x, y, mask
}

// GetPosition is a wrapper around the C function gdk_window_get_position.
func (recv *Window) GetPosition() (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	C.gdk_window_get_position((*C.GdkWindow)(recv.native), &c_x, &c_y)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// GetRootOrigin is a wrapper around the C function gdk_window_get_root_origin.
func (recv *Window) GetRootOrigin() (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	C.gdk_window_get_root_origin((*C.GdkWindow)(recv.native), &c_x, &c_y)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

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

// GetTypeHint is a wrapper around the C function gdk_window_get_type_hint.
func (recv *Window) GetTypeHint() WindowTypeHint {
	retC := C.gdk_window_get_type_hint((*C.GdkWindow)(recv.native))
	retGo := (WindowTypeHint)(retC)

	return retGo
}

// GetUpdateArea is a wrapper around the C function gdk_window_get_update_area.
func (recv *Window) GetUpdateArea() *cairo.Region {
	retC := C.gdk_window_get_update_area((*C.GdkWindow)(recv.native))
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUserData is a wrapper around the C function gdk_window_get_user_data.
func (recv *Window) GetUserData() uintptr {
	var c_data C.gpointer

	C.gdk_window_get_user_data((*C.GdkWindow)(recv.native), &c_data)

	data := (uintptr)(unsafe.Pointer(&c_data))

	return data
}

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

// InputShapeCombineRegion is a wrapper around the C function gdk_window_input_shape_combine_region.
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

// Unsupported : gdk_window_invalidate_maybe_recurse : unsupported parameter child_func : no type generator for WindowChildFunc (GdkWindowChildFunc) for param child_func

// InvalidateRect is a wrapper around the C function gdk_window_invalidate_rect.
func (recv *Window) InvalidateRect(rect *Rectangle, invalidateChildren bool) {
	c_rect := (*C.GdkRectangle)(C.NULL)
	if rect != nil {
		c_rect = (*C.GdkRectangle)(rect.ToC())
	}

	c_invalidate_children :=
		boolToGboolean(invalidateChildren)

	C.gdk_window_invalidate_rect((*C.GdkWindow)(recv.native), c_rect, c_invalidate_children)

	return
}

// InvalidateRegion is a wrapper around the C function gdk_window_invalidate_region.
func (recv *Window) InvalidateRegion(region *cairo.Region, invalidateChildren bool) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

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

// MergeChildInputShapes is a wrapper around the C function gdk_window_merge_child_input_shapes.
func (recv *Window) MergeChildInputShapes() {
	C.gdk_window_merge_child_input_shapes((*C.GdkWindow)(recv.native))

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

// MoveRegion is a wrapper around the C function gdk_window_move_region.
func (recv *Window) MoveRegion(region *cairo.Region, dx int32, dy int32) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	c_dx := (C.gint)(dx)

	c_dy := (C.gint)(dy)

	C.gdk_window_move_region((*C.GdkWindow)(recv.native), c_region, c_dx, c_dy)

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

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc (GdkFilterFunc) for param function

// Reparent is a wrapper around the C function gdk_window_reparent.
func (recv *Window) Reparent(newParent *Window, x int32, y int32) {
	c_new_parent := (*C.GdkWindow)(C.NULL)
	if newParent != nil {
		c_new_parent = (*C.GdkWindow)(newParent.ToC())
	}

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

// SetAcceptFocus is a wrapper around the C function gdk_window_set_accept_focus.
func (recv *Window) SetAcceptFocus(acceptFocus bool) {
	c_accept_focus :=
		boolToGboolean(acceptFocus)

	C.gdk_window_set_accept_focus((*C.GdkWindow)(recv.native), c_accept_focus)

	return
}

// SetBackground is a wrapper around the C function gdk_window_set_background.
func (recv *Window) SetBackground(color *Color) {
	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	C.gdk_window_set_background((*C.GdkWindow)(recv.native), c_color)

	return
}

// SetBackgroundPattern is a wrapper around the C function gdk_window_set_background_pattern.
func (recv *Window) SetBackgroundPattern(pattern *cairo.Pattern) {
	c_pattern := (*C.cairo_pattern_t)(C.NULL)
	if pattern != nil {
		c_pattern = (*C.cairo_pattern_t)(pattern.ToC())
	}

	C.gdk_window_set_background_pattern((*C.GdkWindow)(recv.native), c_pattern)

	return
}

// SetBackgroundRgba is a wrapper around the C function gdk_window_set_background_rgba.
func (recv *Window) SetBackgroundRgba(rgba *RGBA) {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gdk_window_set_background_rgba((*C.GdkWindow)(recv.native), c_rgba)

	return
}

// SetChildInputShapes is a wrapper around the C function gdk_window_set_child_input_shapes.
func (recv *Window) SetChildInputShapes() {
	C.gdk_window_set_child_input_shapes((*C.GdkWindow)(recv.native))

	return
}

// SetChildShapes is a wrapper around the C function gdk_window_set_child_shapes.
func (recv *Window) SetChildShapes() {
	C.gdk_window_set_child_shapes((*C.GdkWindow)(recv.native))

	return
}

// SetCursor is a wrapper around the C function gdk_window_set_cursor.
func (recv *Window) SetCursor(cursor *Cursor) {
	c_cursor := (*C.GdkCursor)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkCursor)(cursor.ToC())
	}

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

// SetFocusOnMap is a wrapper around the C function gdk_window_set_focus_on_map.
func (recv *Window) SetFocusOnMap(focusOnMap bool) {
	c_focus_on_map :=
		boolToGboolean(focusOnMap)

	C.gdk_window_set_focus_on_map((*C.GdkWindow)(recv.native), c_focus_on_map)

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
	c_geometry := (*C.GdkGeometry)(C.NULL)
	if geometry != nil {
		c_geometry = (*C.GdkGeometry)(geometry.ToC())
	}

	c_geom_mask := (C.GdkWindowHints)(geomMask)

	C.gdk_window_set_geometry_hints((*C.GdkWindow)(recv.native), c_geometry, c_geom_mask)

	return
}

// SetGroup is a wrapper around the C function gdk_window_set_group.
func (recv *Window) SetGroup(leader *Window) {
	c_leader := (*C.GdkWindow)(C.NULL)
	if leader != nil {
		c_leader = (*C.GdkWindow)(leader.ToC())
	}

	C.gdk_window_set_group((*C.GdkWindow)(recv.native), c_leader)

	return
}

// SetIconList is a wrapper around the C function gdk_window_set_icon_list.
func (recv *Window) SetIconList(pixbufs *glib.List) {
	c_pixbufs := (*C.GList)(C.NULL)
	if pixbufs != nil {
		c_pixbufs = (*C.GList)(pixbufs.ToC())
	}

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

// SetKeepAbove is a wrapper around the C function gdk_window_set_keep_above.
func (recv *Window) SetKeepAbove(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gdk_window_set_keep_above((*C.GdkWindow)(recv.native), c_setting)

	return
}

// SetKeepBelow is a wrapper around the C function gdk_window_set_keep_below.
func (recv *Window) SetKeepBelow(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gdk_window_set_keep_below((*C.GdkWindow)(recv.native), c_setting)

	return
}

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

// SetSkipPagerHint is a wrapper around the C function gdk_window_set_skip_pager_hint.
func (recv *Window) SetSkipPagerHint(skipsPager bool) {
	c_skips_pager :=
		boolToGboolean(skipsPager)

	C.gdk_window_set_skip_pager_hint((*C.GdkWindow)(recv.native), c_skips_pager)

	return
}

// SetSkipTaskbarHint is a wrapper around the C function gdk_window_set_skip_taskbar_hint.
func (recv *Window) SetSkipTaskbarHint(skipsTaskbar bool) {
	c_skips_taskbar :=
		boolToGboolean(skipsTaskbar)

	C.gdk_window_set_skip_taskbar_hint((*C.GdkWindow)(recv.native), c_skips_taskbar)

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
	c_parent := (*C.GdkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GdkWindow)(parent.ToC())
	}

	C.gdk_window_set_transient_for((*C.GdkWindow)(recv.native), c_parent)

	return
}

// SetTypeHint is a wrapper around the C function gdk_window_set_type_hint.
func (recv *Window) SetTypeHint(hint WindowTypeHint) {
	c_hint := (C.GdkWindowTypeHint)(hint)

	C.gdk_window_set_type_hint((*C.GdkWindow)(recv.native), c_hint)

	return
}

// SetUrgencyHint is a wrapper around the C function gdk_window_set_urgency_hint.
func (recv *Window) SetUrgencyHint(urgent bool) {
	c_urgent :=
		boolToGboolean(urgent)

	C.gdk_window_set_urgency_hint((*C.GdkWindow)(recv.native), c_urgent)

	return
}

// SetUserData is a wrapper around the C function gdk_window_set_user_data.
func (recv *Window) SetUserData(userData *gobject.Object) {
	c_user_data := (C.gpointer)(C.NULL)
	if userData != nil {
		c_user_data = (C.gpointer)(userData.ToC())
	}

	C.gdk_window_set_user_data((*C.GdkWindow)(recv.native), c_user_data)

	return
}

// ShapeCombineRegion is a wrapper around the C function gdk_window_shape_combine_region.
func (recv *Window) ShapeCombineRegion(shapeRegion *cairo.Region, offsetX int32, offsetY int32) {
	c_shape_region := (*C.cairo_region_t)(C.NULL)
	if shapeRegion != nil {
		c_shape_region = (*C.cairo_region_t)(shapeRegion.ToC())
	}

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

// Unfullscreen is a wrapper around the C function gdk_window_unfullscreen.
func (recv *Window) Unfullscreen() {
	C.gdk_window_unfullscreen((*C.GdkWindow)(recv.native))

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

const CURRENT_TIME int32 = C.GDK_CURRENT_TIME
const KEY_0 int32 = C.GDK_KEY_0
const KEY_1 int32 = C.GDK_KEY_1
const KEY_2 int32 = C.GDK_KEY_2
const KEY_3 int32 = C.GDK_KEY_3
const KEY_3270_AltCursor int32 = C.GDK_KEY_3270_AltCursor
const KEY_3270_Attn int32 = C.GDK_KEY_3270_Attn
const KEY_3270_BackTab int32 = C.GDK_KEY_3270_BackTab
const KEY_3270_ChangeScreen int32 = C.GDK_KEY_3270_ChangeScreen
const KEY_3270_Copy int32 = C.GDK_KEY_3270_Copy
const KEY_3270_CursorBlink int32 = C.GDK_KEY_3270_CursorBlink
const KEY_3270_CursorSelect int32 = C.GDK_KEY_3270_CursorSelect
const KEY_3270_DeleteWord int32 = C.GDK_KEY_3270_DeleteWord
const KEY_3270_Duplicate int32 = C.GDK_KEY_3270_Duplicate
const KEY_3270_Enter int32 = C.GDK_KEY_3270_Enter
const KEY_3270_EraseEOF int32 = C.GDK_KEY_3270_EraseEOF
const KEY_3270_EraseInput int32 = C.GDK_KEY_3270_EraseInput
const KEY_3270_ExSelect int32 = C.GDK_KEY_3270_ExSelect
const KEY_3270_FieldMark int32 = C.GDK_KEY_3270_FieldMark
const KEY_3270_Ident int32 = C.GDK_KEY_3270_Ident
const KEY_3270_Jump int32 = C.GDK_KEY_3270_Jump
const KEY_3270_KeyClick int32 = C.GDK_KEY_3270_KeyClick
const KEY_3270_Left2 int32 = C.GDK_KEY_3270_Left2
const KEY_3270_PA1 int32 = C.GDK_KEY_3270_PA1
const KEY_3270_PA2 int32 = C.GDK_KEY_3270_PA2
const KEY_3270_PA3 int32 = C.GDK_KEY_3270_PA3
const KEY_3270_Play int32 = C.GDK_KEY_3270_Play
const KEY_3270_PrintScreen int32 = C.GDK_KEY_3270_PrintScreen
const KEY_3270_Quit int32 = C.GDK_KEY_3270_Quit
const KEY_3270_Record int32 = C.GDK_KEY_3270_Record
const KEY_3270_Reset int32 = C.GDK_KEY_3270_Reset
const KEY_3270_Right2 int32 = C.GDK_KEY_3270_Right2
const KEY_3270_Rule int32 = C.GDK_KEY_3270_Rule
const KEY_3270_Setup int32 = C.GDK_KEY_3270_Setup
const KEY_3270_Test int32 = C.GDK_KEY_3270_Test
const KEY_4 int32 = C.GDK_KEY_4
const KEY_5 int32 = C.GDK_KEY_5
const KEY_6 int32 = C.GDK_KEY_6
const KEY_7 int32 = C.GDK_KEY_7
const KEY_8 int32 = C.GDK_KEY_8
const KEY_9 int32 = C.GDK_KEY_9
const KEY_A int32 = C.GDK_KEY_A
const KEY_AE int32 = C.GDK_KEY_AE
const KEY_Aacute int32 = C.GDK_KEY_Aacute
const KEY_Abelowdot int32 = C.GDK_KEY_Abelowdot
const KEY_Abreve int32 = C.GDK_KEY_Abreve
const KEY_Abreveacute int32 = C.GDK_KEY_Abreveacute
const KEY_Abrevebelowdot int32 = C.GDK_KEY_Abrevebelowdot
const KEY_Abrevegrave int32 = C.GDK_KEY_Abrevegrave
const KEY_Abrevehook int32 = C.GDK_KEY_Abrevehook
const KEY_Abrevetilde int32 = C.GDK_KEY_Abrevetilde
const KEY_AccessX_Enable int32 = C.GDK_KEY_AccessX_Enable
const KEY_AccessX_Feedback_Enable int32 = C.GDK_KEY_AccessX_Feedback_Enable
const KEY_Acircumflex int32 = C.GDK_KEY_Acircumflex
const KEY_Acircumflexacute int32 = C.GDK_KEY_Acircumflexacute
const KEY_Acircumflexbelowdot int32 = C.GDK_KEY_Acircumflexbelowdot
const KEY_Acircumflexgrave int32 = C.GDK_KEY_Acircumflexgrave
const KEY_Acircumflexhook int32 = C.GDK_KEY_Acircumflexhook
const KEY_Acircumflextilde int32 = C.GDK_KEY_Acircumflextilde
const KEY_AddFavorite int32 = C.GDK_KEY_AddFavorite
const KEY_Adiaeresis int32 = C.GDK_KEY_Adiaeresis
const KEY_Agrave int32 = C.GDK_KEY_Agrave
const KEY_Ahook int32 = C.GDK_KEY_Ahook
const KEY_Alt_L int32 = C.GDK_KEY_Alt_L
const KEY_Alt_R int32 = C.GDK_KEY_Alt_R
const KEY_Amacron int32 = C.GDK_KEY_Amacron
const KEY_Aogonek int32 = C.GDK_KEY_Aogonek
const KEY_ApplicationLeft int32 = C.GDK_KEY_ApplicationLeft
const KEY_ApplicationRight int32 = C.GDK_KEY_ApplicationRight
const KEY_Arabic_0 int32 = C.GDK_KEY_Arabic_0
const KEY_Arabic_1 int32 = C.GDK_KEY_Arabic_1
const KEY_Arabic_2 int32 = C.GDK_KEY_Arabic_2
const KEY_Arabic_3 int32 = C.GDK_KEY_Arabic_3
const KEY_Arabic_4 int32 = C.GDK_KEY_Arabic_4
const KEY_Arabic_5 int32 = C.GDK_KEY_Arabic_5
const KEY_Arabic_6 int32 = C.GDK_KEY_Arabic_6
const KEY_Arabic_7 int32 = C.GDK_KEY_Arabic_7
const KEY_Arabic_8 int32 = C.GDK_KEY_Arabic_8
const KEY_Arabic_9 int32 = C.GDK_KEY_Arabic_9
const KEY_Arabic_ain int32 = C.GDK_KEY_Arabic_ain
const KEY_Arabic_alef int32 = C.GDK_KEY_Arabic_alef
const KEY_Arabic_alefmaksura int32 = C.GDK_KEY_Arabic_alefmaksura
const KEY_Arabic_beh int32 = C.GDK_KEY_Arabic_beh
const KEY_Arabic_comma int32 = C.GDK_KEY_Arabic_comma
const KEY_Arabic_dad int32 = C.GDK_KEY_Arabic_dad
const KEY_Arabic_dal int32 = C.GDK_KEY_Arabic_dal
const KEY_Arabic_damma int32 = C.GDK_KEY_Arabic_damma
const KEY_Arabic_dammatan int32 = C.GDK_KEY_Arabic_dammatan
const KEY_Arabic_ddal int32 = C.GDK_KEY_Arabic_ddal
const KEY_Arabic_farsi_yeh int32 = C.GDK_KEY_Arabic_farsi_yeh
const KEY_Arabic_fatha int32 = C.GDK_KEY_Arabic_fatha
const KEY_Arabic_fathatan int32 = C.GDK_KEY_Arabic_fathatan
const KEY_Arabic_feh int32 = C.GDK_KEY_Arabic_feh
const KEY_Arabic_fullstop int32 = C.GDK_KEY_Arabic_fullstop
const KEY_Arabic_gaf int32 = C.GDK_KEY_Arabic_gaf
const KEY_Arabic_ghain int32 = C.GDK_KEY_Arabic_ghain
const KEY_Arabic_ha int32 = C.GDK_KEY_Arabic_ha
const KEY_Arabic_hah int32 = C.GDK_KEY_Arabic_hah
const KEY_Arabic_hamza int32 = C.GDK_KEY_Arabic_hamza
const KEY_Arabic_hamza_above int32 = C.GDK_KEY_Arabic_hamza_above
const KEY_Arabic_hamza_below int32 = C.GDK_KEY_Arabic_hamza_below
const KEY_Arabic_hamzaonalef int32 = C.GDK_KEY_Arabic_hamzaonalef
const KEY_Arabic_hamzaonwaw int32 = C.GDK_KEY_Arabic_hamzaonwaw
const KEY_Arabic_hamzaonyeh int32 = C.GDK_KEY_Arabic_hamzaonyeh
const KEY_Arabic_hamzaunderalef int32 = C.GDK_KEY_Arabic_hamzaunderalef
const KEY_Arabic_heh int32 = C.GDK_KEY_Arabic_heh
const KEY_Arabic_heh_doachashmee int32 = C.GDK_KEY_Arabic_heh_doachashmee
const KEY_Arabic_heh_goal int32 = C.GDK_KEY_Arabic_heh_goal
const KEY_Arabic_jeem int32 = C.GDK_KEY_Arabic_jeem
const KEY_Arabic_jeh int32 = C.GDK_KEY_Arabic_jeh
const KEY_Arabic_kaf int32 = C.GDK_KEY_Arabic_kaf
const KEY_Arabic_kasra int32 = C.GDK_KEY_Arabic_kasra
const KEY_Arabic_kasratan int32 = C.GDK_KEY_Arabic_kasratan
const KEY_Arabic_keheh int32 = C.GDK_KEY_Arabic_keheh
const KEY_Arabic_khah int32 = C.GDK_KEY_Arabic_khah
const KEY_Arabic_lam int32 = C.GDK_KEY_Arabic_lam
const KEY_Arabic_madda_above int32 = C.GDK_KEY_Arabic_madda_above
const KEY_Arabic_maddaonalef int32 = C.GDK_KEY_Arabic_maddaonalef
const KEY_Arabic_meem int32 = C.GDK_KEY_Arabic_meem
const KEY_Arabic_noon int32 = C.GDK_KEY_Arabic_noon
const KEY_Arabic_noon_ghunna int32 = C.GDK_KEY_Arabic_noon_ghunna
const KEY_Arabic_peh int32 = C.GDK_KEY_Arabic_peh
const KEY_Arabic_percent int32 = C.GDK_KEY_Arabic_percent
const KEY_Arabic_qaf int32 = C.GDK_KEY_Arabic_qaf
const KEY_Arabic_question_mark int32 = C.GDK_KEY_Arabic_question_mark
const KEY_Arabic_ra int32 = C.GDK_KEY_Arabic_ra
const KEY_Arabic_rreh int32 = C.GDK_KEY_Arabic_rreh
const KEY_Arabic_sad int32 = C.GDK_KEY_Arabic_sad
const KEY_Arabic_seen int32 = C.GDK_KEY_Arabic_seen
const KEY_Arabic_semicolon int32 = C.GDK_KEY_Arabic_semicolon
const KEY_Arabic_shadda int32 = C.GDK_KEY_Arabic_shadda
const KEY_Arabic_sheen int32 = C.GDK_KEY_Arabic_sheen
const KEY_Arabic_sukun int32 = C.GDK_KEY_Arabic_sukun
const KEY_Arabic_superscript_alef int32 = C.GDK_KEY_Arabic_superscript_alef
const KEY_Arabic_switch int32 = C.GDK_KEY_Arabic_switch
const KEY_Arabic_tah int32 = C.GDK_KEY_Arabic_tah
const KEY_Arabic_tatweel int32 = C.GDK_KEY_Arabic_tatweel
const KEY_Arabic_tcheh int32 = C.GDK_KEY_Arabic_tcheh
const KEY_Arabic_teh int32 = C.GDK_KEY_Arabic_teh
const KEY_Arabic_tehmarbuta int32 = C.GDK_KEY_Arabic_tehmarbuta
const KEY_Arabic_thal int32 = C.GDK_KEY_Arabic_thal
const KEY_Arabic_theh int32 = C.GDK_KEY_Arabic_theh
const KEY_Arabic_tteh int32 = C.GDK_KEY_Arabic_tteh
const KEY_Arabic_veh int32 = C.GDK_KEY_Arabic_veh
const KEY_Arabic_waw int32 = C.GDK_KEY_Arabic_waw
const KEY_Arabic_yeh int32 = C.GDK_KEY_Arabic_yeh
const KEY_Arabic_yeh_baree int32 = C.GDK_KEY_Arabic_yeh_baree
const KEY_Arabic_zah int32 = C.GDK_KEY_Arabic_zah
const KEY_Arabic_zain int32 = C.GDK_KEY_Arabic_zain
const KEY_Aring int32 = C.GDK_KEY_Aring
const KEY_Armenian_AT int32 = C.GDK_KEY_Armenian_AT
const KEY_Armenian_AYB int32 = C.GDK_KEY_Armenian_AYB
const KEY_Armenian_BEN int32 = C.GDK_KEY_Armenian_BEN
const KEY_Armenian_CHA int32 = C.GDK_KEY_Armenian_CHA
const KEY_Armenian_DA int32 = C.GDK_KEY_Armenian_DA
const KEY_Armenian_DZA int32 = C.GDK_KEY_Armenian_DZA
const KEY_Armenian_E int32 = C.GDK_KEY_Armenian_E
const KEY_Armenian_FE int32 = C.GDK_KEY_Armenian_FE
const KEY_Armenian_GHAT int32 = C.GDK_KEY_Armenian_GHAT
const KEY_Armenian_GIM int32 = C.GDK_KEY_Armenian_GIM
const KEY_Armenian_HI int32 = C.GDK_KEY_Armenian_HI
const KEY_Armenian_HO int32 = C.GDK_KEY_Armenian_HO
const KEY_Armenian_INI int32 = C.GDK_KEY_Armenian_INI
const KEY_Armenian_JE int32 = C.GDK_KEY_Armenian_JE
const KEY_Armenian_KE int32 = C.GDK_KEY_Armenian_KE
const KEY_Armenian_KEN int32 = C.GDK_KEY_Armenian_KEN
const KEY_Armenian_KHE int32 = C.GDK_KEY_Armenian_KHE
const KEY_Armenian_LYUN int32 = C.GDK_KEY_Armenian_LYUN
const KEY_Armenian_MEN int32 = C.GDK_KEY_Armenian_MEN
const KEY_Armenian_NU int32 = C.GDK_KEY_Armenian_NU
const KEY_Armenian_O int32 = C.GDK_KEY_Armenian_O
const KEY_Armenian_PE int32 = C.GDK_KEY_Armenian_PE
const KEY_Armenian_PYUR int32 = C.GDK_KEY_Armenian_PYUR
const KEY_Armenian_RA int32 = C.GDK_KEY_Armenian_RA
const KEY_Armenian_RE int32 = C.GDK_KEY_Armenian_RE
const KEY_Armenian_SE int32 = C.GDK_KEY_Armenian_SE
const KEY_Armenian_SHA int32 = C.GDK_KEY_Armenian_SHA
const KEY_Armenian_TCHE int32 = C.GDK_KEY_Armenian_TCHE
const KEY_Armenian_TO int32 = C.GDK_KEY_Armenian_TO
const KEY_Armenian_TSA int32 = C.GDK_KEY_Armenian_TSA
const KEY_Armenian_TSO int32 = C.GDK_KEY_Armenian_TSO
const KEY_Armenian_TYUN int32 = C.GDK_KEY_Armenian_TYUN
const KEY_Armenian_VEV int32 = C.GDK_KEY_Armenian_VEV
const KEY_Armenian_VO int32 = C.GDK_KEY_Armenian_VO
const KEY_Armenian_VYUN int32 = C.GDK_KEY_Armenian_VYUN
const KEY_Armenian_YECH int32 = C.GDK_KEY_Armenian_YECH
const KEY_Armenian_ZA int32 = C.GDK_KEY_Armenian_ZA
const KEY_Armenian_ZHE int32 = C.GDK_KEY_Armenian_ZHE
const KEY_Armenian_accent int32 = C.GDK_KEY_Armenian_accent
const KEY_Armenian_amanak int32 = C.GDK_KEY_Armenian_amanak
const KEY_Armenian_apostrophe int32 = C.GDK_KEY_Armenian_apostrophe
const KEY_Armenian_at int32 = C.GDK_KEY_Armenian_at
const KEY_Armenian_ayb int32 = C.GDK_KEY_Armenian_ayb
const KEY_Armenian_ben int32 = C.GDK_KEY_Armenian_ben
const KEY_Armenian_but int32 = C.GDK_KEY_Armenian_but
const KEY_Armenian_cha int32 = C.GDK_KEY_Armenian_cha
const KEY_Armenian_da int32 = C.GDK_KEY_Armenian_da
const KEY_Armenian_dza int32 = C.GDK_KEY_Armenian_dza
const KEY_Armenian_e int32 = C.GDK_KEY_Armenian_e
const KEY_Armenian_exclam int32 = C.GDK_KEY_Armenian_exclam
const KEY_Armenian_fe int32 = C.GDK_KEY_Armenian_fe
const KEY_Armenian_full_stop int32 = C.GDK_KEY_Armenian_full_stop
const KEY_Armenian_ghat int32 = C.GDK_KEY_Armenian_ghat
const KEY_Armenian_gim int32 = C.GDK_KEY_Armenian_gim
const KEY_Armenian_hi int32 = C.GDK_KEY_Armenian_hi
const KEY_Armenian_ho int32 = C.GDK_KEY_Armenian_ho
const KEY_Armenian_hyphen int32 = C.GDK_KEY_Armenian_hyphen
const KEY_Armenian_ini int32 = C.GDK_KEY_Armenian_ini
const KEY_Armenian_je int32 = C.GDK_KEY_Armenian_je
const KEY_Armenian_ke int32 = C.GDK_KEY_Armenian_ke
const KEY_Armenian_ken int32 = C.GDK_KEY_Armenian_ken
const KEY_Armenian_khe int32 = C.GDK_KEY_Armenian_khe
const KEY_Armenian_ligature_ew int32 = C.GDK_KEY_Armenian_ligature_ew
const KEY_Armenian_lyun int32 = C.GDK_KEY_Armenian_lyun
const KEY_Armenian_men int32 = C.GDK_KEY_Armenian_men
const KEY_Armenian_nu int32 = C.GDK_KEY_Armenian_nu
const KEY_Armenian_o int32 = C.GDK_KEY_Armenian_o
const KEY_Armenian_paruyk int32 = C.GDK_KEY_Armenian_paruyk
const KEY_Armenian_pe int32 = C.GDK_KEY_Armenian_pe
const KEY_Armenian_pyur int32 = C.GDK_KEY_Armenian_pyur
const KEY_Armenian_question int32 = C.GDK_KEY_Armenian_question
const KEY_Armenian_ra int32 = C.GDK_KEY_Armenian_ra
const KEY_Armenian_re int32 = C.GDK_KEY_Armenian_re
const KEY_Armenian_se int32 = C.GDK_KEY_Armenian_se
const KEY_Armenian_separation_mark int32 = C.GDK_KEY_Armenian_separation_mark
const KEY_Armenian_sha int32 = C.GDK_KEY_Armenian_sha
const KEY_Armenian_shesht int32 = C.GDK_KEY_Armenian_shesht
const KEY_Armenian_tche int32 = C.GDK_KEY_Armenian_tche
const KEY_Armenian_to int32 = C.GDK_KEY_Armenian_to
const KEY_Armenian_tsa int32 = C.GDK_KEY_Armenian_tsa
const KEY_Armenian_tso int32 = C.GDK_KEY_Armenian_tso
const KEY_Armenian_tyun int32 = C.GDK_KEY_Armenian_tyun
const KEY_Armenian_verjaket int32 = C.GDK_KEY_Armenian_verjaket
const KEY_Armenian_vev int32 = C.GDK_KEY_Armenian_vev
const KEY_Armenian_vo int32 = C.GDK_KEY_Armenian_vo
const KEY_Armenian_vyun int32 = C.GDK_KEY_Armenian_vyun
const KEY_Armenian_yech int32 = C.GDK_KEY_Armenian_yech
const KEY_Armenian_yentamna int32 = C.GDK_KEY_Armenian_yentamna
const KEY_Armenian_za int32 = C.GDK_KEY_Armenian_za
const KEY_Armenian_zhe int32 = C.GDK_KEY_Armenian_zhe
const KEY_Atilde int32 = C.GDK_KEY_Atilde
const KEY_AudibleBell_Enable int32 = C.GDK_KEY_AudibleBell_Enable
const KEY_AudioCycleTrack int32 = C.GDK_KEY_AudioCycleTrack
const KEY_AudioForward int32 = C.GDK_KEY_AudioForward
const KEY_AudioLowerVolume int32 = C.GDK_KEY_AudioLowerVolume
const KEY_AudioMedia int32 = C.GDK_KEY_AudioMedia
const KEY_AudioMicMute int32 = C.GDK_KEY_AudioMicMute
const KEY_AudioMute int32 = C.GDK_KEY_AudioMute
const KEY_AudioNext int32 = C.GDK_KEY_AudioNext
const KEY_AudioPause int32 = C.GDK_KEY_AudioPause
const KEY_AudioPlay int32 = C.GDK_KEY_AudioPlay
const KEY_AudioPrev int32 = C.GDK_KEY_AudioPrev
const KEY_AudioRaiseVolume int32 = C.GDK_KEY_AudioRaiseVolume
const KEY_AudioRandomPlay int32 = C.GDK_KEY_AudioRandomPlay
const KEY_AudioRecord int32 = C.GDK_KEY_AudioRecord
const KEY_AudioRepeat int32 = C.GDK_KEY_AudioRepeat
const KEY_AudioRewind int32 = C.GDK_KEY_AudioRewind
const KEY_AudioStop int32 = C.GDK_KEY_AudioStop
const KEY_Away int32 = C.GDK_KEY_Away
const KEY_B int32 = C.GDK_KEY_B
const KEY_Babovedot int32 = C.GDK_KEY_Babovedot
const KEY_Back int32 = C.GDK_KEY_Back
const KEY_BackForward int32 = C.GDK_KEY_BackForward
const KEY_BackSpace int32 = C.GDK_KEY_BackSpace
const KEY_Battery int32 = C.GDK_KEY_Battery
const KEY_Begin int32 = C.GDK_KEY_Begin
const KEY_Blue int32 = C.GDK_KEY_Blue
const KEY_Bluetooth int32 = C.GDK_KEY_Bluetooth
const KEY_Book int32 = C.GDK_KEY_Book
const KEY_BounceKeys_Enable int32 = C.GDK_KEY_BounceKeys_Enable
const KEY_Break int32 = C.GDK_KEY_Break
const KEY_BrightnessAdjust int32 = C.GDK_KEY_BrightnessAdjust
const KEY_Byelorussian_SHORTU int32 = C.GDK_KEY_Byelorussian_SHORTU
const KEY_Byelorussian_shortu int32 = C.GDK_KEY_Byelorussian_shortu
const KEY_C int32 = C.GDK_KEY_C
const KEY_CD int32 = C.GDK_KEY_CD
const KEY_CH int32 = C.GDK_KEY_CH
const KEY_C_H int32 = C.GDK_KEY_C_H
const KEY_C_h int32 = C.GDK_KEY_C_h
const KEY_Cabovedot int32 = C.GDK_KEY_Cabovedot
const KEY_Cacute int32 = C.GDK_KEY_Cacute
const KEY_Calculator int32 = C.GDK_KEY_Calculator
const KEY_Calendar int32 = C.GDK_KEY_Calendar
const KEY_Cancel int32 = C.GDK_KEY_Cancel
const KEY_Caps_Lock int32 = C.GDK_KEY_Caps_Lock
const KEY_Ccaron int32 = C.GDK_KEY_Ccaron
const KEY_Ccedilla int32 = C.GDK_KEY_Ccedilla
const KEY_Ccircumflex int32 = C.GDK_KEY_Ccircumflex
const KEY_Ch int32 = C.GDK_KEY_Ch
const KEY_Clear int32 = C.GDK_KEY_Clear
const KEY_ClearGrab int32 = C.GDK_KEY_ClearGrab
const KEY_Close int32 = C.GDK_KEY_Close
const KEY_Codeinput int32 = C.GDK_KEY_Codeinput
const KEY_ColonSign int32 = C.GDK_KEY_ColonSign
const KEY_Community int32 = C.GDK_KEY_Community
const KEY_ContrastAdjust int32 = C.GDK_KEY_ContrastAdjust
const KEY_Control_L int32 = C.GDK_KEY_Control_L
const KEY_Control_R int32 = C.GDK_KEY_Control_R
const KEY_Copy int32 = C.GDK_KEY_Copy
const KEY_CruzeiroSign int32 = C.GDK_KEY_CruzeiroSign
const KEY_Cut int32 = C.GDK_KEY_Cut
const KEY_CycleAngle int32 = C.GDK_KEY_CycleAngle
const KEY_Cyrillic_A int32 = C.GDK_KEY_Cyrillic_A
const KEY_Cyrillic_BE int32 = C.GDK_KEY_Cyrillic_BE
const KEY_Cyrillic_CHE int32 = C.GDK_KEY_Cyrillic_CHE
const KEY_Cyrillic_CHE_descender int32 = C.GDK_KEY_Cyrillic_CHE_descender
const KEY_Cyrillic_CHE_vertstroke int32 = C.GDK_KEY_Cyrillic_CHE_vertstroke
const KEY_Cyrillic_DE int32 = C.GDK_KEY_Cyrillic_DE
const KEY_Cyrillic_DZHE int32 = C.GDK_KEY_Cyrillic_DZHE
const KEY_Cyrillic_E int32 = C.GDK_KEY_Cyrillic_E
const KEY_Cyrillic_EF int32 = C.GDK_KEY_Cyrillic_EF
const KEY_Cyrillic_EL int32 = C.GDK_KEY_Cyrillic_EL
const KEY_Cyrillic_EM int32 = C.GDK_KEY_Cyrillic_EM
const KEY_Cyrillic_EN int32 = C.GDK_KEY_Cyrillic_EN
const KEY_Cyrillic_EN_descender int32 = C.GDK_KEY_Cyrillic_EN_descender
const KEY_Cyrillic_ER int32 = C.GDK_KEY_Cyrillic_ER
const KEY_Cyrillic_ES int32 = C.GDK_KEY_Cyrillic_ES
const KEY_Cyrillic_GHE int32 = C.GDK_KEY_Cyrillic_GHE
const KEY_Cyrillic_GHE_bar int32 = C.GDK_KEY_Cyrillic_GHE_bar
const KEY_Cyrillic_HA int32 = C.GDK_KEY_Cyrillic_HA
const KEY_Cyrillic_HARDSIGN int32 = C.GDK_KEY_Cyrillic_HARDSIGN
const KEY_Cyrillic_HA_descender int32 = C.GDK_KEY_Cyrillic_HA_descender
const KEY_Cyrillic_I int32 = C.GDK_KEY_Cyrillic_I
const KEY_Cyrillic_IE int32 = C.GDK_KEY_Cyrillic_IE
const KEY_Cyrillic_IO int32 = C.GDK_KEY_Cyrillic_IO
const KEY_Cyrillic_I_macron int32 = C.GDK_KEY_Cyrillic_I_macron
const KEY_Cyrillic_JE int32 = C.GDK_KEY_Cyrillic_JE
const KEY_Cyrillic_KA int32 = C.GDK_KEY_Cyrillic_KA
const KEY_Cyrillic_KA_descender int32 = C.GDK_KEY_Cyrillic_KA_descender
const KEY_Cyrillic_KA_vertstroke int32 = C.GDK_KEY_Cyrillic_KA_vertstroke
const KEY_Cyrillic_LJE int32 = C.GDK_KEY_Cyrillic_LJE
const KEY_Cyrillic_NJE int32 = C.GDK_KEY_Cyrillic_NJE
const KEY_Cyrillic_O int32 = C.GDK_KEY_Cyrillic_O
const KEY_Cyrillic_O_bar int32 = C.GDK_KEY_Cyrillic_O_bar
const KEY_Cyrillic_PE int32 = C.GDK_KEY_Cyrillic_PE
const KEY_Cyrillic_SCHWA int32 = C.GDK_KEY_Cyrillic_SCHWA
const KEY_Cyrillic_SHA int32 = C.GDK_KEY_Cyrillic_SHA
const KEY_Cyrillic_SHCHA int32 = C.GDK_KEY_Cyrillic_SHCHA
const KEY_Cyrillic_SHHA int32 = C.GDK_KEY_Cyrillic_SHHA
const KEY_Cyrillic_SHORTI int32 = C.GDK_KEY_Cyrillic_SHORTI
const KEY_Cyrillic_SOFTSIGN int32 = C.GDK_KEY_Cyrillic_SOFTSIGN
const KEY_Cyrillic_TE int32 = C.GDK_KEY_Cyrillic_TE
const KEY_Cyrillic_TSE int32 = C.GDK_KEY_Cyrillic_TSE
const KEY_Cyrillic_U int32 = C.GDK_KEY_Cyrillic_U
const KEY_Cyrillic_U_macron int32 = C.GDK_KEY_Cyrillic_U_macron
const KEY_Cyrillic_U_straight int32 = C.GDK_KEY_Cyrillic_U_straight
const KEY_Cyrillic_U_straight_bar int32 = C.GDK_KEY_Cyrillic_U_straight_bar
const KEY_Cyrillic_VE int32 = C.GDK_KEY_Cyrillic_VE
const KEY_Cyrillic_YA int32 = C.GDK_KEY_Cyrillic_YA
const KEY_Cyrillic_YERU int32 = C.GDK_KEY_Cyrillic_YERU
const KEY_Cyrillic_YU int32 = C.GDK_KEY_Cyrillic_YU
const KEY_Cyrillic_ZE int32 = C.GDK_KEY_Cyrillic_ZE
const KEY_Cyrillic_ZHE int32 = C.GDK_KEY_Cyrillic_ZHE
const KEY_Cyrillic_ZHE_descender int32 = C.GDK_KEY_Cyrillic_ZHE_descender
const KEY_Cyrillic_a int32 = C.GDK_KEY_Cyrillic_a
const KEY_Cyrillic_be int32 = C.GDK_KEY_Cyrillic_be
const KEY_Cyrillic_che int32 = C.GDK_KEY_Cyrillic_che
const KEY_Cyrillic_che_descender int32 = C.GDK_KEY_Cyrillic_che_descender
const KEY_Cyrillic_che_vertstroke int32 = C.GDK_KEY_Cyrillic_che_vertstroke
const KEY_Cyrillic_de int32 = C.GDK_KEY_Cyrillic_de
const KEY_Cyrillic_dzhe int32 = C.GDK_KEY_Cyrillic_dzhe
const KEY_Cyrillic_e int32 = C.GDK_KEY_Cyrillic_e
const KEY_Cyrillic_ef int32 = C.GDK_KEY_Cyrillic_ef
const KEY_Cyrillic_el int32 = C.GDK_KEY_Cyrillic_el
const KEY_Cyrillic_em int32 = C.GDK_KEY_Cyrillic_em
const KEY_Cyrillic_en int32 = C.GDK_KEY_Cyrillic_en
const KEY_Cyrillic_en_descender int32 = C.GDK_KEY_Cyrillic_en_descender
const KEY_Cyrillic_er int32 = C.GDK_KEY_Cyrillic_er
const KEY_Cyrillic_es int32 = C.GDK_KEY_Cyrillic_es
const KEY_Cyrillic_ghe int32 = C.GDK_KEY_Cyrillic_ghe
const KEY_Cyrillic_ghe_bar int32 = C.GDK_KEY_Cyrillic_ghe_bar
const KEY_Cyrillic_ha int32 = C.GDK_KEY_Cyrillic_ha
const KEY_Cyrillic_ha_descender int32 = C.GDK_KEY_Cyrillic_ha_descender
const KEY_Cyrillic_hardsign int32 = C.GDK_KEY_Cyrillic_hardsign
const KEY_Cyrillic_i int32 = C.GDK_KEY_Cyrillic_i
const KEY_Cyrillic_i_macron int32 = C.GDK_KEY_Cyrillic_i_macron
const KEY_Cyrillic_ie int32 = C.GDK_KEY_Cyrillic_ie
const KEY_Cyrillic_io int32 = C.GDK_KEY_Cyrillic_io
const KEY_Cyrillic_je int32 = C.GDK_KEY_Cyrillic_je
const KEY_Cyrillic_ka int32 = C.GDK_KEY_Cyrillic_ka
const KEY_Cyrillic_ka_descender int32 = C.GDK_KEY_Cyrillic_ka_descender
const KEY_Cyrillic_ka_vertstroke int32 = C.GDK_KEY_Cyrillic_ka_vertstroke
const KEY_Cyrillic_lje int32 = C.GDK_KEY_Cyrillic_lje
const KEY_Cyrillic_nje int32 = C.GDK_KEY_Cyrillic_nje
const KEY_Cyrillic_o int32 = C.GDK_KEY_Cyrillic_o
const KEY_Cyrillic_o_bar int32 = C.GDK_KEY_Cyrillic_o_bar
const KEY_Cyrillic_pe int32 = C.GDK_KEY_Cyrillic_pe
const KEY_Cyrillic_schwa int32 = C.GDK_KEY_Cyrillic_schwa
const KEY_Cyrillic_sha int32 = C.GDK_KEY_Cyrillic_sha
const KEY_Cyrillic_shcha int32 = C.GDK_KEY_Cyrillic_shcha
const KEY_Cyrillic_shha int32 = C.GDK_KEY_Cyrillic_shha
const KEY_Cyrillic_shorti int32 = C.GDK_KEY_Cyrillic_shorti
const KEY_Cyrillic_softsign int32 = C.GDK_KEY_Cyrillic_softsign
const KEY_Cyrillic_te int32 = C.GDK_KEY_Cyrillic_te
const KEY_Cyrillic_tse int32 = C.GDK_KEY_Cyrillic_tse
const KEY_Cyrillic_u int32 = C.GDK_KEY_Cyrillic_u
const KEY_Cyrillic_u_macron int32 = C.GDK_KEY_Cyrillic_u_macron
const KEY_Cyrillic_u_straight int32 = C.GDK_KEY_Cyrillic_u_straight
const KEY_Cyrillic_u_straight_bar int32 = C.GDK_KEY_Cyrillic_u_straight_bar
const KEY_Cyrillic_ve int32 = C.GDK_KEY_Cyrillic_ve
const KEY_Cyrillic_ya int32 = C.GDK_KEY_Cyrillic_ya
const KEY_Cyrillic_yeru int32 = C.GDK_KEY_Cyrillic_yeru
const KEY_Cyrillic_yu int32 = C.GDK_KEY_Cyrillic_yu
const KEY_Cyrillic_ze int32 = C.GDK_KEY_Cyrillic_ze
const KEY_Cyrillic_zhe int32 = C.GDK_KEY_Cyrillic_zhe
const KEY_Cyrillic_zhe_descender int32 = C.GDK_KEY_Cyrillic_zhe_descender
const KEY_D int32 = C.GDK_KEY_D
const KEY_DOS int32 = C.GDK_KEY_DOS
const KEY_Dabovedot int32 = C.GDK_KEY_Dabovedot
const KEY_Dcaron int32 = C.GDK_KEY_Dcaron
const KEY_Delete int32 = C.GDK_KEY_Delete
const KEY_Display int32 = C.GDK_KEY_Display
const KEY_Documents int32 = C.GDK_KEY_Documents
const KEY_DongSign int32 = C.GDK_KEY_DongSign
const KEY_Down int32 = C.GDK_KEY_Down
const KEY_Dstroke int32 = C.GDK_KEY_Dstroke
const KEY_E int32 = C.GDK_KEY_E
const KEY_ENG int32 = C.GDK_KEY_ENG
const KEY_ETH int32 = C.GDK_KEY_ETH
const KEY_EZH int32 = C.GDK_KEY_EZH
const KEY_Eabovedot int32 = C.GDK_KEY_Eabovedot
const KEY_Eacute int32 = C.GDK_KEY_Eacute
const KEY_Ebelowdot int32 = C.GDK_KEY_Ebelowdot
const KEY_Ecaron int32 = C.GDK_KEY_Ecaron
const KEY_Ecircumflex int32 = C.GDK_KEY_Ecircumflex
const KEY_Ecircumflexacute int32 = C.GDK_KEY_Ecircumflexacute
const KEY_Ecircumflexbelowdot int32 = C.GDK_KEY_Ecircumflexbelowdot
const KEY_Ecircumflexgrave int32 = C.GDK_KEY_Ecircumflexgrave
const KEY_Ecircumflexhook int32 = C.GDK_KEY_Ecircumflexhook
const KEY_Ecircumflextilde int32 = C.GDK_KEY_Ecircumflextilde
const KEY_EcuSign int32 = C.GDK_KEY_EcuSign
const KEY_Ediaeresis int32 = C.GDK_KEY_Ediaeresis
const KEY_Egrave int32 = C.GDK_KEY_Egrave
const KEY_Ehook int32 = C.GDK_KEY_Ehook
const KEY_Eisu_Shift int32 = C.GDK_KEY_Eisu_Shift
const KEY_Eisu_toggle int32 = C.GDK_KEY_Eisu_toggle
const KEY_Eject int32 = C.GDK_KEY_Eject
const KEY_Emacron int32 = C.GDK_KEY_Emacron
const KEY_End int32 = C.GDK_KEY_End
const KEY_Eogonek int32 = C.GDK_KEY_Eogonek
const KEY_Escape int32 = C.GDK_KEY_Escape
const KEY_Eth int32 = C.GDK_KEY_Eth
const KEY_Etilde int32 = C.GDK_KEY_Etilde
const KEY_EuroSign int32 = C.GDK_KEY_EuroSign
const KEY_Excel int32 = C.GDK_KEY_Excel
const KEY_Execute int32 = C.GDK_KEY_Execute
const KEY_Explorer int32 = C.GDK_KEY_Explorer
const KEY_F int32 = C.GDK_KEY_F
const KEY_F1 int32 = C.GDK_KEY_F1
const KEY_F10 int32 = C.GDK_KEY_F10
const KEY_F11 int32 = C.GDK_KEY_F11
const KEY_F12 int32 = C.GDK_KEY_F12
const KEY_F13 int32 = C.GDK_KEY_F13
const KEY_F14 int32 = C.GDK_KEY_F14
const KEY_F15 int32 = C.GDK_KEY_F15
const KEY_F16 int32 = C.GDK_KEY_F16
const KEY_F17 int32 = C.GDK_KEY_F17
const KEY_F18 int32 = C.GDK_KEY_F18
const KEY_F19 int32 = C.GDK_KEY_F19
const KEY_F2 int32 = C.GDK_KEY_F2
const KEY_F20 int32 = C.GDK_KEY_F20
const KEY_F21 int32 = C.GDK_KEY_F21
const KEY_F22 int32 = C.GDK_KEY_F22
const KEY_F23 int32 = C.GDK_KEY_F23
const KEY_F24 int32 = C.GDK_KEY_F24
const KEY_F25 int32 = C.GDK_KEY_F25
const KEY_F26 int32 = C.GDK_KEY_F26
const KEY_F27 int32 = C.GDK_KEY_F27
const KEY_F28 int32 = C.GDK_KEY_F28
const KEY_F29 int32 = C.GDK_KEY_F29
const KEY_F3 int32 = C.GDK_KEY_F3
const KEY_F30 int32 = C.GDK_KEY_F30
const KEY_F31 int32 = C.GDK_KEY_F31
const KEY_F32 int32 = C.GDK_KEY_F32
const KEY_F33 int32 = C.GDK_KEY_F33
const KEY_F34 int32 = C.GDK_KEY_F34
const KEY_F35 int32 = C.GDK_KEY_F35
const KEY_F4 int32 = C.GDK_KEY_F4
const KEY_F5 int32 = C.GDK_KEY_F5
const KEY_F6 int32 = C.GDK_KEY_F6
const KEY_F7 int32 = C.GDK_KEY_F7
const KEY_F8 int32 = C.GDK_KEY_F8
const KEY_F9 int32 = C.GDK_KEY_F9
const KEY_FFrancSign int32 = C.GDK_KEY_FFrancSign
const KEY_Fabovedot int32 = C.GDK_KEY_Fabovedot
const KEY_Farsi_0 int32 = C.GDK_KEY_Farsi_0
const KEY_Farsi_1 int32 = C.GDK_KEY_Farsi_1
const KEY_Farsi_2 int32 = C.GDK_KEY_Farsi_2
const KEY_Farsi_3 int32 = C.GDK_KEY_Farsi_3
const KEY_Farsi_4 int32 = C.GDK_KEY_Farsi_4
const KEY_Farsi_5 int32 = C.GDK_KEY_Farsi_5
const KEY_Farsi_6 int32 = C.GDK_KEY_Farsi_6
const KEY_Farsi_7 int32 = C.GDK_KEY_Farsi_7
const KEY_Farsi_8 int32 = C.GDK_KEY_Farsi_8
const KEY_Farsi_9 int32 = C.GDK_KEY_Farsi_9
const KEY_Farsi_yeh int32 = C.GDK_KEY_Farsi_yeh
const KEY_Favorites int32 = C.GDK_KEY_Favorites
const KEY_Finance int32 = C.GDK_KEY_Finance
const KEY_Find int32 = C.GDK_KEY_Find
const KEY_First_Virtual_Screen int32 = C.GDK_KEY_First_Virtual_Screen
const KEY_Forward int32 = C.GDK_KEY_Forward
const KEY_FrameBack int32 = C.GDK_KEY_FrameBack
const KEY_FrameForward int32 = C.GDK_KEY_FrameForward
const KEY_G int32 = C.GDK_KEY_G
const KEY_Gabovedot int32 = C.GDK_KEY_Gabovedot
const KEY_Game int32 = C.GDK_KEY_Game
const KEY_Gbreve int32 = C.GDK_KEY_Gbreve
const KEY_Gcaron int32 = C.GDK_KEY_Gcaron
const KEY_Gcedilla int32 = C.GDK_KEY_Gcedilla
const KEY_Gcircumflex int32 = C.GDK_KEY_Gcircumflex
const KEY_Georgian_an int32 = C.GDK_KEY_Georgian_an
const KEY_Georgian_ban int32 = C.GDK_KEY_Georgian_ban
const KEY_Georgian_can int32 = C.GDK_KEY_Georgian_can
const KEY_Georgian_char int32 = C.GDK_KEY_Georgian_char
const KEY_Georgian_chin int32 = C.GDK_KEY_Georgian_chin
const KEY_Georgian_cil int32 = C.GDK_KEY_Georgian_cil
const KEY_Georgian_don int32 = C.GDK_KEY_Georgian_don
const KEY_Georgian_en int32 = C.GDK_KEY_Georgian_en
const KEY_Georgian_fi int32 = C.GDK_KEY_Georgian_fi
const KEY_Georgian_gan int32 = C.GDK_KEY_Georgian_gan
const KEY_Georgian_ghan int32 = C.GDK_KEY_Georgian_ghan
const KEY_Georgian_hae int32 = C.GDK_KEY_Georgian_hae
const KEY_Georgian_har int32 = C.GDK_KEY_Georgian_har
const KEY_Georgian_he int32 = C.GDK_KEY_Georgian_he
const KEY_Georgian_hie int32 = C.GDK_KEY_Georgian_hie
const KEY_Georgian_hoe int32 = C.GDK_KEY_Georgian_hoe
const KEY_Georgian_in int32 = C.GDK_KEY_Georgian_in
const KEY_Georgian_jhan int32 = C.GDK_KEY_Georgian_jhan
const KEY_Georgian_jil int32 = C.GDK_KEY_Georgian_jil
const KEY_Georgian_kan int32 = C.GDK_KEY_Georgian_kan
const KEY_Georgian_khar int32 = C.GDK_KEY_Georgian_khar
const KEY_Georgian_las int32 = C.GDK_KEY_Georgian_las
const KEY_Georgian_man int32 = C.GDK_KEY_Georgian_man
const KEY_Georgian_nar int32 = C.GDK_KEY_Georgian_nar
const KEY_Georgian_on int32 = C.GDK_KEY_Georgian_on
const KEY_Georgian_par int32 = C.GDK_KEY_Georgian_par
const KEY_Georgian_phar int32 = C.GDK_KEY_Georgian_phar
const KEY_Georgian_qar int32 = C.GDK_KEY_Georgian_qar
const KEY_Georgian_rae int32 = C.GDK_KEY_Georgian_rae
const KEY_Georgian_san int32 = C.GDK_KEY_Georgian_san
const KEY_Georgian_shin int32 = C.GDK_KEY_Georgian_shin
const KEY_Georgian_tan int32 = C.GDK_KEY_Georgian_tan
const KEY_Georgian_tar int32 = C.GDK_KEY_Georgian_tar
const KEY_Georgian_un int32 = C.GDK_KEY_Georgian_un
const KEY_Georgian_vin int32 = C.GDK_KEY_Georgian_vin
const KEY_Georgian_we int32 = C.GDK_KEY_Georgian_we
const KEY_Georgian_xan int32 = C.GDK_KEY_Georgian_xan
const KEY_Georgian_zen int32 = C.GDK_KEY_Georgian_zen
const KEY_Georgian_zhar int32 = C.GDK_KEY_Georgian_zhar
const KEY_Go int32 = C.GDK_KEY_Go
const KEY_Greek_ALPHA int32 = C.GDK_KEY_Greek_ALPHA
const KEY_Greek_ALPHAaccent int32 = C.GDK_KEY_Greek_ALPHAaccent
const KEY_Greek_BETA int32 = C.GDK_KEY_Greek_BETA
const KEY_Greek_CHI int32 = C.GDK_KEY_Greek_CHI
const KEY_Greek_DELTA int32 = C.GDK_KEY_Greek_DELTA
const KEY_Greek_EPSILON int32 = C.GDK_KEY_Greek_EPSILON
const KEY_Greek_EPSILONaccent int32 = C.GDK_KEY_Greek_EPSILONaccent
const KEY_Greek_ETA int32 = C.GDK_KEY_Greek_ETA
const KEY_Greek_ETAaccent int32 = C.GDK_KEY_Greek_ETAaccent
const KEY_Greek_GAMMA int32 = C.GDK_KEY_Greek_GAMMA
const KEY_Greek_IOTA int32 = C.GDK_KEY_Greek_IOTA
const KEY_Greek_IOTAaccent int32 = C.GDK_KEY_Greek_IOTAaccent
const KEY_Greek_IOTAdiaeresis int32 = C.GDK_KEY_Greek_IOTAdiaeresis
const KEY_Greek_IOTAdieresis int32 = C.GDK_KEY_Greek_IOTAdieresis
const KEY_Greek_KAPPA int32 = C.GDK_KEY_Greek_KAPPA
const KEY_Greek_LAMBDA int32 = C.GDK_KEY_Greek_LAMBDA
const KEY_Greek_LAMDA int32 = C.GDK_KEY_Greek_LAMDA
const KEY_Greek_MU int32 = C.GDK_KEY_Greek_MU
const KEY_Greek_NU int32 = C.GDK_KEY_Greek_NU
const KEY_Greek_OMEGA int32 = C.GDK_KEY_Greek_OMEGA
const KEY_Greek_OMEGAaccent int32 = C.GDK_KEY_Greek_OMEGAaccent
const KEY_Greek_OMICRON int32 = C.GDK_KEY_Greek_OMICRON
const KEY_Greek_OMICRONaccent int32 = C.GDK_KEY_Greek_OMICRONaccent
const KEY_Greek_PHI int32 = C.GDK_KEY_Greek_PHI
const KEY_Greek_PI int32 = C.GDK_KEY_Greek_PI
const KEY_Greek_PSI int32 = C.GDK_KEY_Greek_PSI
const KEY_Greek_RHO int32 = C.GDK_KEY_Greek_RHO
const KEY_Greek_SIGMA int32 = C.GDK_KEY_Greek_SIGMA
const KEY_Greek_TAU int32 = C.GDK_KEY_Greek_TAU
const KEY_Greek_THETA int32 = C.GDK_KEY_Greek_THETA
const KEY_Greek_UPSILON int32 = C.GDK_KEY_Greek_UPSILON
const KEY_Greek_UPSILONaccent int32 = C.GDK_KEY_Greek_UPSILONaccent
const KEY_Greek_UPSILONdieresis int32 = C.GDK_KEY_Greek_UPSILONdieresis
const KEY_Greek_XI int32 = C.GDK_KEY_Greek_XI
const KEY_Greek_ZETA int32 = C.GDK_KEY_Greek_ZETA
const KEY_Greek_accentdieresis int32 = C.GDK_KEY_Greek_accentdieresis
const KEY_Greek_alpha int32 = C.GDK_KEY_Greek_alpha
const KEY_Greek_alphaaccent int32 = C.GDK_KEY_Greek_alphaaccent
const KEY_Greek_beta int32 = C.GDK_KEY_Greek_beta
const KEY_Greek_chi int32 = C.GDK_KEY_Greek_chi
const KEY_Greek_delta int32 = C.GDK_KEY_Greek_delta
const KEY_Greek_epsilon int32 = C.GDK_KEY_Greek_epsilon
const KEY_Greek_epsilonaccent int32 = C.GDK_KEY_Greek_epsilonaccent
const KEY_Greek_eta int32 = C.GDK_KEY_Greek_eta
const KEY_Greek_etaaccent int32 = C.GDK_KEY_Greek_etaaccent
const KEY_Greek_finalsmallsigma int32 = C.GDK_KEY_Greek_finalsmallsigma
const KEY_Greek_gamma int32 = C.GDK_KEY_Greek_gamma
const KEY_Greek_horizbar int32 = C.GDK_KEY_Greek_horizbar
const KEY_Greek_iota int32 = C.GDK_KEY_Greek_iota
const KEY_Greek_iotaaccent int32 = C.GDK_KEY_Greek_iotaaccent
const KEY_Greek_iotaaccentdieresis int32 = C.GDK_KEY_Greek_iotaaccentdieresis
const KEY_Greek_iotadieresis int32 = C.GDK_KEY_Greek_iotadieresis
const KEY_Greek_kappa int32 = C.GDK_KEY_Greek_kappa
const KEY_Greek_lambda int32 = C.GDK_KEY_Greek_lambda
const KEY_Greek_lamda int32 = C.GDK_KEY_Greek_lamda
const KEY_Greek_mu int32 = C.GDK_KEY_Greek_mu
const KEY_Greek_nu int32 = C.GDK_KEY_Greek_nu
const KEY_Greek_omega int32 = C.GDK_KEY_Greek_omega
const KEY_Greek_omegaaccent int32 = C.GDK_KEY_Greek_omegaaccent
const KEY_Greek_omicron int32 = C.GDK_KEY_Greek_omicron
const KEY_Greek_omicronaccent int32 = C.GDK_KEY_Greek_omicronaccent
const KEY_Greek_phi int32 = C.GDK_KEY_Greek_phi
const KEY_Greek_pi int32 = C.GDK_KEY_Greek_pi
const KEY_Greek_psi int32 = C.GDK_KEY_Greek_psi
const KEY_Greek_rho int32 = C.GDK_KEY_Greek_rho
const KEY_Greek_sigma int32 = C.GDK_KEY_Greek_sigma
const KEY_Greek_switch int32 = C.GDK_KEY_Greek_switch
const KEY_Greek_tau int32 = C.GDK_KEY_Greek_tau
const KEY_Greek_theta int32 = C.GDK_KEY_Greek_theta
const KEY_Greek_upsilon int32 = C.GDK_KEY_Greek_upsilon
const KEY_Greek_upsilonaccent int32 = C.GDK_KEY_Greek_upsilonaccent
const KEY_Greek_upsilonaccentdieresis int32 = C.GDK_KEY_Greek_upsilonaccentdieresis
const KEY_Greek_upsilondieresis int32 = C.GDK_KEY_Greek_upsilondieresis
const KEY_Greek_xi int32 = C.GDK_KEY_Greek_xi
const KEY_Greek_zeta int32 = C.GDK_KEY_Greek_zeta
const KEY_Green int32 = C.GDK_KEY_Green
const KEY_H int32 = C.GDK_KEY_H
const KEY_Hangul int32 = C.GDK_KEY_Hangul
const KEY_Hangul_A int32 = C.GDK_KEY_Hangul_A
const KEY_Hangul_AE int32 = C.GDK_KEY_Hangul_AE
const KEY_Hangul_AraeA int32 = C.GDK_KEY_Hangul_AraeA
const KEY_Hangul_AraeAE int32 = C.GDK_KEY_Hangul_AraeAE
const KEY_Hangul_Banja int32 = C.GDK_KEY_Hangul_Banja
const KEY_Hangul_Cieuc int32 = C.GDK_KEY_Hangul_Cieuc
const KEY_Hangul_Codeinput int32 = C.GDK_KEY_Hangul_Codeinput
const KEY_Hangul_Dikeud int32 = C.GDK_KEY_Hangul_Dikeud
const KEY_Hangul_E int32 = C.GDK_KEY_Hangul_E
const KEY_Hangul_EO int32 = C.GDK_KEY_Hangul_EO
const KEY_Hangul_EU int32 = C.GDK_KEY_Hangul_EU
const KEY_Hangul_End int32 = C.GDK_KEY_Hangul_End
const KEY_Hangul_Hanja int32 = C.GDK_KEY_Hangul_Hanja
const KEY_Hangul_Hieuh int32 = C.GDK_KEY_Hangul_Hieuh
const KEY_Hangul_I int32 = C.GDK_KEY_Hangul_I
const KEY_Hangul_Ieung int32 = C.GDK_KEY_Hangul_Ieung
const KEY_Hangul_J_Cieuc int32 = C.GDK_KEY_Hangul_J_Cieuc
const KEY_Hangul_J_Dikeud int32 = C.GDK_KEY_Hangul_J_Dikeud
const KEY_Hangul_J_Hieuh int32 = C.GDK_KEY_Hangul_J_Hieuh
const KEY_Hangul_J_Ieung int32 = C.GDK_KEY_Hangul_J_Ieung
const KEY_Hangul_J_Jieuj int32 = C.GDK_KEY_Hangul_J_Jieuj
const KEY_Hangul_J_Khieuq int32 = C.GDK_KEY_Hangul_J_Khieuq
const KEY_Hangul_J_Kiyeog int32 = C.GDK_KEY_Hangul_J_Kiyeog
const KEY_Hangul_J_KiyeogSios int32 = C.GDK_KEY_Hangul_J_KiyeogSios
const KEY_Hangul_J_KkogjiDalrinIeung int32 = C.GDK_KEY_Hangul_J_KkogjiDalrinIeung
const KEY_Hangul_J_Mieum int32 = C.GDK_KEY_Hangul_J_Mieum
const KEY_Hangul_J_Nieun int32 = C.GDK_KEY_Hangul_J_Nieun
const KEY_Hangul_J_NieunHieuh int32 = C.GDK_KEY_Hangul_J_NieunHieuh
const KEY_Hangul_J_NieunJieuj int32 = C.GDK_KEY_Hangul_J_NieunJieuj
const KEY_Hangul_J_PanSios int32 = C.GDK_KEY_Hangul_J_PanSios
const KEY_Hangul_J_Phieuf int32 = C.GDK_KEY_Hangul_J_Phieuf
const KEY_Hangul_J_Pieub int32 = C.GDK_KEY_Hangul_J_Pieub
const KEY_Hangul_J_PieubSios int32 = C.GDK_KEY_Hangul_J_PieubSios
const KEY_Hangul_J_Rieul int32 = C.GDK_KEY_Hangul_J_Rieul
const KEY_Hangul_J_RieulHieuh int32 = C.GDK_KEY_Hangul_J_RieulHieuh
const KEY_Hangul_J_RieulKiyeog int32 = C.GDK_KEY_Hangul_J_RieulKiyeog
const KEY_Hangul_J_RieulMieum int32 = C.GDK_KEY_Hangul_J_RieulMieum
const KEY_Hangul_J_RieulPhieuf int32 = C.GDK_KEY_Hangul_J_RieulPhieuf
const KEY_Hangul_J_RieulPieub int32 = C.GDK_KEY_Hangul_J_RieulPieub
const KEY_Hangul_J_RieulSios int32 = C.GDK_KEY_Hangul_J_RieulSios
const KEY_Hangul_J_RieulTieut int32 = C.GDK_KEY_Hangul_J_RieulTieut
const KEY_Hangul_J_Sios int32 = C.GDK_KEY_Hangul_J_Sios
const KEY_Hangul_J_SsangKiyeog int32 = C.GDK_KEY_Hangul_J_SsangKiyeog
const KEY_Hangul_J_SsangSios int32 = C.GDK_KEY_Hangul_J_SsangSios
const KEY_Hangul_J_Tieut int32 = C.GDK_KEY_Hangul_J_Tieut
const KEY_Hangul_J_YeorinHieuh int32 = C.GDK_KEY_Hangul_J_YeorinHieuh
const KEY_Hangul_Jamo int32 = C.GDK_KEY_Hangul_Jamo
const KEY_Hangul_Jeonja int32 = C.GDK_KEY_Hangul_Jeonja
const KEY_Hangul_Jieuj int32 = C.GDK_KEY_Hangul_Jieuj
const KEY_Hangul_Khieuq int32 = C.GDK_KEY_Hangul_Khieuq
const KEY_Hangul_Kiyeog int32 = C.GDK_KEY_Hangul_Kiyeog
const KEY_Hangul_KiyeogSios int32 = C.GDK_KEY_Hangul_KiyeogSios
const KEY_Hangul_KkogjiDalrinIeung int32 = C.GDK_KEY_Hangul_KkogjiDalrinIeung
const KEY_Hangul_Mieum int32 = C.GDK_KEY_Hangul_Mieum
const KEY_Hangul_MultipleCandidate int32 = C.GDK_KEY_Hangul_MultipleCandidate
const KEY_Hangul_Nieun int32 = C.GDK_KEY_Hangul_Nieun
const KEY_Hangul_NieunHieuh int32 = C.GDK_KEY_Hangul_NieunHieuh
const KEY_Hangul_NieunJieuj int32 = C.GDK_KEY_Hangul_NieunJieuj
const KEY_Hangul_O int32 = C.GDK_KEY_Hangul_O
const KEY_Hangul_OE int32 = C.GDK_KEY_Hangul_OE
const KEY_Hangul_PanSios int32 = C.GDK_KEY_Hangul_PanSios
const KEY_Hangul_Phieuf int32 = C.GDK_KEY_Hangul_Phieuf
const KEY_Hangul_Pieub int32 = C.GDK_KEY_Hangul_Pieub
const KEY_Hangul_PieubSios int32 = C.GDK_KEY_Hangul_PieubSios
const KEY_Hangul_PostHanja int32 = C.GDK_KEY_Hangul_PostHanja
const KEY_Hangul_PreHanja int32 = C.GDK_KEY_Hangul_PreHanja
const KEY_Hangul_PreviousCandidate int32 = C.GDK_KEY_Hangul_PreviousCandidate
const KEY_Hangul_Rieul int32 = C.GDK_KEY_Hangul_Rieul
const KEY_Hangul_RieulHieuh int32 = C.GDK_KEY_Hangul_RieulHieuh
const KEY_Hangul_RieulKiyeog int32 = C.GDK_KEY_Hangul_RieulKiyeog
const KEY_Hangul_RieulMieum int32 = C.GDK_KEY_Hangul_RieulMieum
const KEY_Hangul_RieulPhieuf int32 = C.GDK_KEY_Hangul_RieulPhieuf
const KEY_Hangul_RieulPieub int32 = C.GDK_KEY_Hangul_RieulPieub
const KEY_Hangul_RieulSios int32 = C.GDK_KEY_Hangul_RieulSios
const KEY_Hangul_RieulTieut int32 = C.GDK_KEY_Hangul_RieulTieut
const KEY_Hangul_RieulYeorinHieuh int32 = C.GDK_KEY_Hangul_RieulYeorinHieuh
const KEY_Hangul_Romaja int32 = C.GDK_KEY_Hangul_Romaja
const KEY_Hangul_SingleCandidate int32 = C.GDK_KEY_Hangul_SingleCandidate
const KEY_Hangul_Sios int32 = C.GDK_KEY_Hangul_Sios
const KEY_Hangul_Special int32 = C.GDK_KEY_Hangul_Special
const KEY_Hangul_SsangDikeud int32 = C.GDK_KEY_Hangul_SsangDikeud
const KEY_Hangul_SsangJieuj int32 = C.GDK_KEY_Hangul_SsangJieuj
const KEY_Hangul_SsangKiyeog int32 = C.GDK_KEY_Hangul_SsangKiyeog
const KEY_Hangul_SsangPieub int32 = C.GDK_KEY_Hangul_SsangPieub
const KEY_Hangul_SsangSios int32 = C.GDK_KEY_Hangul_SsangSios
const KEY_Hangul_Start int32 = C.GDK_KEY_Hangul_Start
const KEY_Hangul_SunkyeongeumMieum int32 = C.GDK_KEY_Hangul_SunkyeongeumMieum
const KEY_Hangul_SunkyeongeumPhieuf int32 = C.GDK_KEY_Hangul_SunkyeongeumPhieuf
const KEY_Hangul_SunkyeongeumPieub int32 = C.GDK_KEY_Hangul_SunkyeongeumPieub
const KEY_Hangul_Tieut int32 = C.GDK_KEY_Hangul_Tieut
const KEY_Hangul_U int32 = C.GDK_KEY_Hangul_U
const KEY_Hangul_WA int32 = C.GDK_KEY_Hangul_WA
const KEY_Hangul_WAE int32 = C.GDK_KEY_Hangul_WAE
const KEY_Hangul_WE int32 = C.GDK_KEY_Hangul_WE
const KEY_Hangul_WEO int32 = C.GDK_KEY_Hangul_WEO
const KEY_Hangul_WI int32 = C.GDK_KEY_Hangul_WI
const KEY_Hangul_YA int32 = C.GDK_KEY_Hangul_YA
const KEY_Hangul_YAE int32 = C.GDK_KEY_Hangul_YAE
const KEY_Hangul_YE int32 = C.GDK_KEY_Hangul_YE
const KEY_Hangul_YEO int32 = C.GDK_KEY_Hangul_YEO
const KEY_Hangul_YI int32 = C.GDK_KEY_Hangul_YI
const KEY_Hangul_YO int32 = C.GDK_KEY_Hangul_YO
const KEY_Hangul_YU int32 = C.GDK_KEY_Hangul_YU
const KEY_Hangul_YeorinHieuh int32 = C.GDK_KEY_Hangul_YeorinHieuh
const KEY_Hangul_switch int32 = C.GDK_KEY_Hangul_switch
const KEY_Hankaku int32 = C.GDK_KEY_Hankaku
const KEY_Hcircumflex int32 = C.GDK_KEY_Hcircumflex
const KEY_Hebrew_switch int32 = C.GDK_KEY_Hebrew_switch
const KEY_Help int32 = C.GDK_KEY_Help
const KEY_Henkan int32 = C.GDK_KEY_Henkan
const KEY_Henkan_Mode int32 = C.GDK_KEY_Henkan_Mode
const KEY_Hibernate int32 = C.GDK_KEY_Hibernate
const KEY_Hiragana int32 = C.GDK_KEY_Hiragana
const KEY_Hiragana_Katakana int32 = C.GDK_KEY_Hiragana_Katakana
const KEY_History int32 = C.GDK_KEY_History
const KEY_Home int32 = C.GDK_KEY_Home
const KEY_HomePage int32 = C.GDK_KEY_HomePage
const KEY_HotLinks int32 = C.GDK_KEY_HotLinks
const KEY_Hstroke int32 = C.GDK_KEY_Hstroke
const KEY_Hyper_L int32 = C.GDK_KEY_Hyper_L
const KEY_Hyper_R int32 = C.GDK_KEY_Hyper_R
const KEY_I int32 = C.GDK_KEY_I
const KEY_ISO_Center_Object int32 = C.GDK_KEY_ISO_Center_Object
const KEY_ISO_Continuous_Underline int32 = C.GDK_KEY_ISO_Continuous_Underline
const KEY_ISO_Discontinuous_Underline int32 = C.GDK_KEY_ISO_Discontinuous_Underline
const KEY_ISO_Emphasize int32 = C.GDK_KEY_ISO_Emphasize
const KEY_ISO_Enter int32 = C.GDK_KEY_ISO_Enter
const KEY_ISO_Fast_Cursor_Down int32 = C.GDK_KEY_ISO_Fast_Cursor_Down
const KEY_ISO_Fast_Cursor_Left int32 = C.GDK_KEY_ISO_Fast_Cursor_Left
const KEY_ISO_Fast_Cursor_Right int32 = C.GDK_KEY_ISO_Fast_Cursor_Right
const KEY_ISO_Fast_Cursor_Up int32 = C.GDK_KEY_ISO_Fast_Cursor_Up
const KEY_ISO_First_Group int32 = C.GDK_KEY_ISO_First_Group
const KEY_ISO_First_Group_Lock int32 = C.GDK_KEY_ISO_First_Group_Lock
const KEY_ISO_Group_Latch int32 = C.GDK_KEY_ISO_Group_Latch
const KEY_ISO_Group_Lock int32 = C.GDK_KEY_ISO_Group_Lock
const KEY_ISO_Group_Shift int32 = C.GDK_KEY_ISO_Group_Shift
const KEY_ISO_Last_Group int32 = C.GDK_KEY_ISO_Last_Group
const KEY_ISO_Last_Group_Lock int32 = C.GDK_KEY_ISO_Last_Group_Lock
const KEY_ISO_Left_Tab int32 = C.GDK_KEY_ISO_Left_Tab
const KEY_ISO_Level2_Latch int32 = C.GDK_KEY_ISO_Level2_Latch
const KEY_ISO_Level3_Latch int32 = C.GDK_KEY_ISO_Level3_Latch
const KEY_ISO_Level3_Lock int32 = C.GDK_KEY_ISO_Level3_Lock
const KEY_ISO_Level3_Shift int32 = C.GDK_KEY_ISO_Level3_Shift
const KEY_ISO_Level5_Latch int32 = C.GDK_KEY_ISO_Level5_Latch
const KEY_ISO_Level5_Lock int32 = C.GDK_KEY_ISO_Level5_Lock
const KEY_ISO_Level5_Shift int32 = C.GDK_KEY_ISO_Level5_Shift
const KEY_ISO_Lock int32 = C.GDK_KEY_ISO_Lock
const KEY_ISO_Move_Line_Down int32 = C.GDK_KEY_ISO_Move_Line_Down
const KEY_ISO_Move_Line_Up int32 = C.GDK_KEY_ISO_Move_Line_Up
const KEY_ISO_Next_Group int32 = C.GDK_KEY_ISO_Next_Group
const KEY_ISO_Next_Group_Lock int32 = C.GDK_KEY_ISO_Next_Group_Lock
const KEY_ISO_Partial_Line_Down int32 = C.GDK_KEY_ISO_Partial_Line_Down
const KEY_ISO_Partial_Line_Up int32 = C.GDK_KEY_ISO_Partial_Line_Up
const KEY_ISO_Partial_Space_Left int32 = C.GDK_KEY_ISO_Partial_Space_Left
const KEY_ISO_Partial_Space_Right int32 = C.GDK_KEY_ISO_Partial_Space_Right
const KEY_ISO_Prev_Group int32 = C.GDK_KEY_ISO_Prev_Group
const KEY_ISO_Prev_Group_Lock int32 = C.GDK_KEY_ISO_Prev_Group_Lock
const KEY_ISO_Release_Both_Margins int32 = C.GDK_KEY_ISO_Release_Both_Margins
const KEY_ISO_Release_Margin_Left int32 = C.GDK_KEY_ISO_Release_Margin_Left
const KEY_ISO_Release_Margin_Right int32 = C.GDK_KEY_ISO_Release_Margin_Right
const KEY_ISO_Set_Margin_Left int32 = C.GDK_KEY_ISO_Set_Margin_Left
const KEY_ISO_Set_Margin_Right int32 = C.GDK_KEY_ISO_Set_Margin_Right
const KEY_Iabovedot int32 = C.GDK_KEY_Iabovedot
const KEY_Iacute int32 = C.GDK_KEY_Iacute
const KEY_Ibelowdot int32 = C.GDK_KEY_Ibelowdot
const KEY_Ibreve int32 = C.GDK_KEY_Ibreve
const KEY_Icircumflex int32 = C.GDK_KEY_Icircumflex
const KEY_Idiaeresis int32 = C.GDK_KEY_Idiaeresis
const KEY_Igrave int32 = C.GDK_KEY_Igrave
const KEY_Ihook int32 = C.GDK_KEY_Ihook
const KEY_Imacron int32 = C.GDK_KEY_Imacron
const KEY_Insert int32 = C.GDK_KEY_Insert
const KEY_Iogonek int32 = C.GDK_KEY_Iogonek
const KEY_Itilde int32 = C.GDK_KEY_Itilde
const KEY_J int32 = C.GDK_KEY_J
const KEY_Jcircumflex int32 = C.GDK_KEY_Jcircumflex
const KEY_K int32 = C.GDK_KEY_K
const KEY_KP_0 int32 = C.GDK_KEY_KP_0
const KEY_KP_1 int32 = C.GDK_KEY_KP_1
const KEY_KP_2 int32 = C.GDK_KEY_KP_2
const KEY_KP_3 int32 = C.GDK_KEY_KP_3
const KEY_KP_4 int32 = C.GDK_KEY_KP_4
const KEY_KP_5 int32 = C.GDK_KEY_KP_5
const KEY_KP_6 int32 = C.GDK_KEY_KP_6
const KEY_KP_7 int32 = C.GDK_KEY_KP_7
const KEY_KP_8 int32 = C.GDK_KEY_KP_8
const KEY_KP_9 int32 = C.GDK_KEY_KP_9
const KEY_KP_Add int32 = C.GDK_KEY_KP_Add
const KEY_KP_Begin int32 = C.GDK_KEY_KP_Begin
const KEY_KP_Decimal int32 = C.GDK_KEY_KP_Decimal
const KEY_KP_Delete int32 = C.GDK_KEY_KP_Delete
const KEY_KP_Divide int32 = C.GDK_KEY_KP_Divide
const KEY_KP_Down int32 = C.GDK_KEY_KP_Down
const KEY_KP_End int32 = C.GDK_KEY_KP_End
const KEY_KP_Enter int32 = C.GDK_KEY_KP_Enter
const KEY_KP_Equal int32 = C.GDK_KEY_KP_Equal
const KEY_KP_F1 int32 = C.GDK_KEY_KP_F1
const KEY_KP_F2 int32 = C.GDK_KEY_KP_F2
const KEY_KP_F3 int32 = C.GDK_KEY_KP_F3
const KEY_KP_F4 int32 = C.GDK_KEY_KP_F4
const KEY_KP_Home int32 = C.GDK_KEY_KP_Home
const KEY_KP_Insert int32 = C.GDK_KEY_KP_Insert
const KEY_KP_Left int32 = C.GDK_KEY_KP_Left
const KEY_KP_Multiply int32 = C.GDK_KEY_KP_Multiply
const KEY_KP_Next int32 = C.GDK_KEY_KP_Next
const KEY_KP_Page_Down int32 = C.GDK_KEY_KP_Page_Down
const KEY_KP_Page_Up int32 = C.GDK_KEY_KP_Page_Up
const KEY_KP_Prior int32 = C.GDK_KEY_KP_Prior
const KEY_KP_Right int32 = C.GDK_KEY_KP_Right
const KEY_KP_Separator int32 = C.GDK_KEY_KP_Separator
const KEY_KP_Space int32 = C.GDK_KEY_KP_Space
const KEY_KP_Subtract int32 = C.GDK_KEY_KP_Subtract
const KEY_KP_Tab int32 = C.GDK_KEY_KP_Tab
const KEY_KP_Up int32 = C.GDK_KEY_KP_Up
const KEY_Kana_Lock int32 = C.GDK_KEY_Kana_Lock
const KEY_Kana_Shift int32 = C.GDK_KEY_Kana_Shift
const KEY_Kanji int32 = C.GDK_KEY_Kanji
const KEY_Kanji_Bangou int32 = C.GDK_KEY_Kanji_Bangou
const KEY_Katakana int32 = C.GDK_KEY_Katakana
const KEY_KbdBrightnessDown int32 = C.GDK_KEY_KbdBrightnessDown
const KEY_KbdBrightnessUp int32 = C.GDK_KEY_KbdBrightnessUp
const KEY_KbdLightOnOff int32 = C.GDK_KEY_KbdLightOnOff
const KEY_Kcedilla int32 = C.GDK_KEY_Kcedilla
const KEY_Korean_Won int32 = C.GDK_KEY_Korean_Won
const KEY_L int32 = C.GDK_KEY_L
const KEY_L1 int32 = C.GDK_KEY_L1
const KEY_L10 int32 = C.GDK_KEY_L10
const KEY_L2 int32 = C.GDK_KEY_L2
const KEY_L3 int32 = C.GDK_KEY_L3
const KEY_L4 int32 = C.GDK_KEY_L4
const KEY_L5 int32 = C.GDK_KEY_L5
const KEY_L6 int32 = C.GDK_KEY_L6
const KEY_L7 int32 = C.GDK_KEY_L7
const KEY_L8 int32 = C.GDK_KEY_L8
const KEY_L9 int32 = C.GDK_KEY_L9
const KEY_Lacute int32 = C.GDK_KEY_Lacute
const KEY_Last_Virtual_Screen int32 = C.GDK_KEY_Last_Virtual_Screen
const KEY_Launch0 int32 = C.GDK_KEY_Launch0
const KEY_Launch1 int32 = C.GDK_KEY_Launch1
const KEY_Launch2 int32 = C.GDK_KEY_Launch2
const KEY_Launch3 int32 = C.GDK_KEY_Launch3
const KEY_Launch4 int32 = C.GDK_KEY_Launch4
const KEY_Launch5 int32 = C.GDK_KEY_Launch5
const KEY_Launch6 int32 = C.GDK_KEY_Launch6
const KEY_Launch7 int32 = C.GDK_KEY_Launch7
const KEY_Launch8 int32 = C.GDK_KEY_Launch8
const KEY_Launch9 int32 = C.GDK_KEY_Launch9
const KEY_LaunchA int32 = C.GDK_KEY_LaunchA
const KEY_LaunchB int32 = C.GDK_KEY_LaunchB
const KEY_LaunchC int32 = C.GDK_KEY_LaunchC
const KEY_LaunchD int32 = C.GDK_KEY_LaunchD
const KEY_LaunchE int32 = C.GDK_KEY_LaunchE
const KEY_LaunchF int32 = C.GDK_KEY_LaunchF
const KEY_Lbelowdot int32 = C.GDK_KEY_Lbelowdot
const KEY_Lcaron int32 = C.GDK_KEY_Lcaron
const KEY_Lcedilla int32 = C.GDK_KEY_Lcedilla
const KEY_Left int32 = C.GDK_KEY_Left
const KEY_LightBulb int32 = C.GDK_KEY_LightBulb
const KEY_Linefeed int32 = C.GDK_KEY_Linefeed
const KEY_LiraSign int32 = C.GDK_KEY_LiraSign
const KEY_LogGrabInfo int32 = C.GDK_KEY_LogGrabInfo
const KEY_LogOff int32 = C.GDK_KEY_LogOff
const KEY_LogWindowTree int32 = C.GDK_KEY_LogWindowTree
const KEY_Lstroke int32 = C.GDK_KEY_Lstroke
const KEY_M int32 = C.GDK_KEY_M
const KEY_Mabovedot int32 = C.GDK_KEY_Mabovedot
const KEY_Macedonia_DSE int32 = C.GDK_KEY_Macedonia_DSE
const KEY_Macedonia_GJE int32 = C.GDK_KEY_Macedonia_GJE
const KEY_Macedonia_KJE int32 = C.GDK_KEY_Macedonia_KJE
const KEY_Macedonia_dse int32 = C.GDK_KEY_Macedonia_dse
const KEY_Macedonia_gje int32 = C.GDK_KEY_Macedonia_gje
const KEY_Macedonia_kje int32 = C.GDK_KEY_Macedonia_kje
const KEY_Mae_Koho int32 = C.GDK_KEY_Mae_Koho
const KEY_Mail int32 = C.GDK_KEY_Mail
const KEY_MailForward int32 = C.GDK_KEY_MailForward
const KEY_Market int32 = C.GDK_KEY_Market
const KEY_Massyo int32 = C.GDK_KEY_Massyo
const KEY_Meeting int32 = C.GDK_KEY_Meeting
const KEY_Memo int32 = C.GDK_KEY_Memo
const KEY_Menu int32 = C.GDK_KEY_Menu
const KEY_MenuKB int32 = C.GDK_KEY_MenuKB
const KEY_MenuPB int32 = C.GDK_KEY_MenuPB
const KEY_Messenger int32 = C.GDK_KEY_Messenger
const KEY_Meta_L int32 = C.GDK_KEY_Meta_L
const KEY_Meta_R int32 = C.GDK_KEY_Meta_R
const KEY_MillSign int32 = C.GDK_KEY_MillSign
const KEY_ModeLock int32 = C.GDK_KEY_ModeLock
const KEY_Mode_switch int32 = C.GDK_KEY_Mode_switch
const KEY_MonBrightnessDown int32 = C.GDK_KEY_MonBrightnessDown
const KEY_MonBrightnessUp int32 = C.GDK_KEY_MonBrightnessUp
const KEY_MouseKeys_Accel_Enable int32 = C.GDK_KEY_MouseKeys_Accel_Enable
const KEY_MouseKeys_Enable int32 = C.GDK_KEY_MouseKeys_Enable
const KEY_Muhenkan int32 = C.GDK_KEY_Muhenkan
const KEY_Multi_key int32 = C.GDK_KEY_Multi_key
const KEY_MultipleCandidate int32 = C.GDK_KEY_MultipleCandidate
const KEY_Music int32 = C.GDK_KEY_Music
const KEY_MyComputer int32 = C.GDK_KEY_MyComputer
const KEY_MySites int32 = C.GDK_KEY_MySites
const KEY_N int32 = C.GDK_KEY_N
const KEY_Nacute int32 = C.GDK_KEY_Nacute
const KEY_NairaSign int32 = C.GDK_KEY_NairaSign
const KEY_Ncaron int32 = C.GDK_KEY_Ncaron
const KEY_Ncedilla int32 = C.GDK_KEY_Ncedilla
const KEY_New int32 = C.GDK_KEY_New
const KEY_NewSheqelSign int32 = C.GDK_KEY_NewSheqelSign
const KEY_News int32 = C.GDK_KEY_News
const KEY_Next int32 = C.GDK_KEY_Next
const KEY_Next_VMode int32 = C.GDK_KEY_Next_VMode
const KEY_Next_Virtual_Screen int32 = C.GDK_KEY_Next_Virtual_Screen
const KEY_Ntilde int32 = C.GDK_KEY_Ntilde
const KEY_Num_Lock int32 = C.GDK_KEY_Num_Lock
const KEY_O int32 = C.GDK_KEY_O
const KEY_OE int32 = C.GDK_KEY_OE
const KEY_Oacute int32 = C.GDK_KEY_Oacute
const KEY_Obarred int32 = C.GDK_KEY_Obarred
const KEY_Obelowdot int32 = C.GDK_KEY_Obelowdot
const KEY_Ocaron int32 = C.GDK_KEY_Ocaron
const KEY_Ocircumflex int32 = C.GDK_KEY_Ocircumflex
const KEY_Ocircumflexacute int32 = C.GDK_KEY_Ocircumflexacute
const KEY_Ocircumflexbelowdot int32 = C.GDK_KEY_Ocircumflexbelowdot
const KEY_Ocircumflexgrave int32 = C.GDK_KEY_Ocircumflexgrave
const KEY_Ocircumflexhook int32 = C.GDK_KEY_Ocircumflexhook
const KEY_Ocircumflextilde int32 = C.GDK_KEY_Ocircumflextilde
const KEY_Odiaeresis int32 = C.GDK_KEY_Odiaeresis
const KEY_Odoubleacute int32 = C.GDK_KEY_Odoubleacute
const KEY_OfficeHome int32 = C.GDK_KEY_OfficeHome
const KEY_Ograve int32 = C.GDK_KEY_Ograve
const KEY_Ohook int32 = C.GDK_KEY_Ohook
const KEY_Ohorn int32 = C.GDK_KEY_Ohorn
const KEY_Ohornacute int32 = C.GDK_KEY_Ohornacute
const KEY_Ohornbelowdot int32 = C.GDK_KEY_Ohornbelowdot
const KEY_Ohorngrave int32 = C.GDK_KEY_Ohorngrave
const KEY_Ohornhook int32 = C.GDK_KEY_Ohornhook
const KEY_Ohorntilde int32 = C.GDK_KEY_Ohorntilde
const KEY_Omacron int32 = C.GDK_KEY_Omacron
const KEY_Ooblique int32 = C.GDK_KEY_Ooblique
const KEY_Open int32 = C.GDK_KEY_Open
const KEY_OpenURL int32 = C.GDK_KEY_OpenURL
const KEY_Option int32 = C.GDK_KEY_Option
const KEY_Oslash int32 = C.GDK_KEY_Oslash
const KEY_Otilde int32 = C.GDK_KEY_Otilde
const KEY_Overlay1_Enable int32 = C.GDK_KEY_Overlay1_Enable
const KEY_Overlay2_Enable int32 = C.GDK_KEY_Overlay2_Enable
const KEY_P int32 = C.GDK_KEY_P
const KEY_Pabovedot int32 = C.GDK_KEY_Pabovedot
const KEY_Page_Down int32 = C.GDK_KEY_Page_Down
const KEY_Page_Up int32 = C.GDK_KEY_Page_Up
const KEY_Paste int32 = C.GDK_KEY_Paste
const KEY_Pause int32 = C.GDK_KEY_Pause
const KEY_PesetaSign int32 = C.GDK_KEY_PesetaSign
const KEY_Phone int32 = C.GDK_KEY_Phone
const KEY_Pictures int32 = C.GDK_KEY_Pictures
const KEY_Pointer_Accelerate int32 = C.GDK_KEY_Pointer_Accelerate
const KEY_Pointer_Button1 int32 = C.GDK_KEY_Pointer_Button1
const KEY_Pointer_Button2 int32 = C.GDK_KEY_Pointer_Button2
const KEY_Pointer_Button3 int32 = C.GDK_KEY_Pointer_Button3
const KEY_Pointer_Button4 int32 = C.GDK_KEY_Pointer_Button4
const KEY_Pointer_Button5 int32 = C.GDK_KEY_Pointer_Button5
const KEY_Pointer_Button_Dflt int32 = C.GDK_KEY_Pointer_Button_Dflt
const KEY_Pointer_DblClick1 int32 = C.GDK_KEY_Pointer_DblClick1
const KEY_Pointer_DblClick2 int32 = C.GDK_KEY_Pointer_DblClick2
const KEY_Pointer_DblClick3 int32 = C.GDK_KEY_Pointer_DblClick3
const KEY_Pointer_DblClick4 int32 = C.GDK_KEY_Pointer_DblClick4
const KEY_Pointer_DblClick5 int32 = C.GDK_KEY_Pointer_DblClick5
const KEY_Pointer_DblClick_Dflt int32 = C.GDK_KEY_Pointer_DblClick_Dflt
const KEY_Pointer_DfltBtnNext int32 = C.GDK_KEY_Pointer_DfltBtnNext
const KEY_Pointer_DfltBtnPrev int32 = C.GDK_KEY_Pointer_DfltBtnPrev
const KEY_Pointer_Down int32 = C.GDK_KEY_Pointer_Down
const KEY_Pointer_DownLeft int32 = C.GDK_KEY_Pointer_DownLeft
const KEY_Pointer_DownRight int32 = C.GDK_KEY_Pointer_DownRight
const KEY_Pointer_Drag1 int32 = C.GDK_KEY_Pointer_Drag1
const KEY_Pointer_Drag2 int32 = C.GDK_KEY_Pointer_Drag2
const KEY_Pointer_Drag3 int32 = C.GDK_KEY_Pointer_Drag3
const KEY_Pointer_Drag4 int32 = C.GDK_KEY_Pointer_Drag4
const KEY_Pointer_Drag5 int32 = C.GDK_KEY_Pointer_Drag5
const KEY_Pointer_Drag_Dflt int32 = C.GDK_KEY_Pointer_Drag_Dflt
const KEY_Pointer_EnableKeys int32 = C.GDK_KEY_Pointer_EnableKeys
const KEY_Pointer_Left int32 = C.GDK_KEY_Pointer_Left
const KEY_Pointer_Right int32 = C.GDK_KEY_Pointer_Right
const KEY_Pointer_Up int32 = C.GDK_KEY_Pointer_Up
const KEY_Pointer_UpLeft int32 = C.GDK_KEY_Pointer_UpLeft
const KEY_Pointer_UpRight int32 = C.GDK_KEY_Pointer_UpRight
const KEY_PowerDown int32 = C.GDK_KEY_PowerDown
const KEY_PowerOff int32 = C.GDK_KEY_PowerOff
const KEY_Prev_VMode int32 = C.GDK_KEY_Prev_VMode
const KEY_Prev_Virtual_Screen int32 = C.GDK_KEY_Prev_Virtual_Screen
const KEY_PreviousCandidate int32 = C.GDK_KEY_PreviousCandidate
const KEY_Print int32 = C.GDK_KEY_Print
const KEY_Prior int32 = C.GDK_KEY_Prior
const KEY_Q int32 = C.GDK_KEY_Q
const KEY_R int32 = C.GDK_KEY_R
const KEY_R1 int32 = C.GDK_KEY_R1
const KEY_R10 int32 = C.GDK_KEY_R10
const KEY_R11 int32 = C.GDK_KEY_R11
const KEY_R12 int32 = C.GDK_KEY_R12
const KEY_R13 int32 = C.GDK_KEY_R13
const KEY_R14 int32 = C.GDK_KEY_R14
const KEY_R15 int32 = C.GDK_KEY_R15
const KEY_R2 int32 = C.GDK_KEY_R2
const KEY_R3 int32 = C.GDK_KEY_R3
const KEY_R4 int32 = C.GDK_KEY_R4
const KEY_R5 int32 = C.GDK_KEY_R5
const KEY_R6 int32 = C.GDK_KEY_R6
const KEY_R7 int32 = C.GDK_KEY_R7
const KEY_R8 int32 = C.GDK_KEY_R8
const KEY_R9 int32 = C.GDK_KEY_R9
const KEY_Racute int32 = C.GDK_KEY_Racute
const KEY_Rcaron int32 = C.GDK_KEY_Rcaron
const KEY_Rcedilla int32 = C.GDK_KEY_Rcedilla
const KEY_Red int32 = C.GDK_KEY_Red
const KEY_Redo int32 = C.GDK_KEY_Redo
const KEY_Refresh int32 = C.GDK_KEY_Refresh
const KEY_Reload int32 = C.GDK_KEY_Reload
const KEY_RepeatKeys_Enable int32 = C.GDK_KEY_RepeatKeys_Enable
const KEY_Reply int32 = C.GDK_KEY_Reply
const KEY_Return int32 = C.GDK_KEY_Return
const KEY_Right int32 = C.GDK_KEY_Right
const KEY_RockerDown int32 = C.GDK_KEY_RockerDown
const KEY_RockerEnter int32 = C.GDK_KEY_RockerEnter
const KEY_RockerUp int32 = C.GDK_KEY_RockerUp
const KEY_Romaji int32 = C.GDK_KEY_Romaji
const KEY_RotateWindows int32 = C.GDK_KEY_RotateWindows
const KEY_RotationKB int32 = C.GDK_KEY_RotationKB
const KEY_RotationPB int32 = C.GDK_KEY_RotationPB
const KEY_RupeeSign int32 = C.GDK_KEY_RupeeSign
const KEY_S int32 = C.GDK_KEY_S
const KEY_SCHWA int32 = C.GDK_KEY_SCHWA
const KEY_Sabovedot int32 = C.GDK_KEY_Sabovedot
const KEY_Sacute int32 = C.GDK_KEY_Sacute
const KEY_Save int32 = C.GDK_KEY_Save
const KEY_Scaron int32 = C.GDK_KEY_Scaron
const KEY_Scedilla int32 = C.GDK_KEY_Scedilla
const KEY_Scircumflex int32 = C.GDK_KEY_Scircumflex
const KEY_ScreenSaver int32 = C.GDK_KEY_ScreenSaver
const KEY_ScrollClick int32 = C.GDK_KEY_ScrollClick
const KEY_ScrollDown int32 = C.GDK_KEY_ScrollDown
const KEY_ScrollUp int32 = C.GDK_KEY_ScrollUp
const KEY_Scroll_Lock int32 = C.GDK_KEY_Scroll_Lock
const KEY_Search int32 = C.GDK_KEY_Search
const KEY_Select int32 = C.GDK_KEY_Select
const KEY_SelectButton int32 = C.GDK_KEY_SelectButton
const KEY_Send int32 = C.GDK_KEY_Send
const KEY_Serbian_DJE int32 = C.GDK_KEY_Serbian_DJE
const KEY_Serbian_DZE int32 = C.GDK_KEY_Serbian_DZE
const KEY_Serbian_JE int32 = C.GDK_KEY_Serbian_JE
const KEY_Serbian_LJE int32 = C.GDK_KEY_Serbian_LJE
const KEY_Serbian_NJE int32 = C.GDK_KEY_Serbian_NJE
const KEY_Serbian_TSHE int32 = C.GDK_KEY_Serbian_TSHE
const KEY_Serbian_dje int32 = C.GDK_KEY_Serbian_dje
const KEY_Serbian_dze int32 = C.GDK_KEY_Serbian_dze
const KEY_Serbian_je int32 = C.GDK_KEY_Serbian_je
const KEY_Serbian_lje int32 = C.GDK_KEY_Serbian_lje
const KEY_Serbian_nje int32 = C.GDK_KEY_Serbian_nje
const KEY_Serbian_tshe int32 = C.GDK_KEY_Serbian_tshe
const KEY_Shift_L int32 = C.GDK_KEY_Shift_L
const KEY_Shift_Lock int32 = C.GDK_KEY_Shift_Lock
const KEY_Shift_R int32 = C.GDK_KEY_Shift_R
const KEY_Shop int32 = C.GDK_KEY_Shop
const KEY_SingleCandidate int32 = C.GDK_KEY_SingleCandidate
const KEY_Sinh_a int32 = C.GDK_KEY_Sinh_a
const KEY_Sinh_aa int32 = C.GDK_KEY_Sinh_aa
const KEY_Sinh_aa2 int32 = C.GDK_KEY_Sinh_aa2
const KEY_Sinh_ae int32 = C.GDK_KEY_Sinh_ae
const KEY_Sinh_ae2 int32 = C.GDK_KEY_Sinh_ae2
const KEY_Sinh_aee int32 = C.GDK_KEY_Sinh_aee
const KEY_Sinh_aee2 int32 = C.GDK_KEY_Sinh_aee2
const KEY_Sinh_ai int32 = C.GDK_KEY_Sinh_ai
const KEY_Sinh_ai2 int32 = C.GDK_KEY_Sinh_ai2
const KEY_Sinh_al int32 = C.GDK_KEY_Sinh_al
const KEY_Sinh_au int32 = C.GDK_KEY_Sinh_au
const KEY_Sinh_au2 int32 = C.GDK_KEY_Sinh_au2
const KEY_Sinh_ba int32 = C.GDK_KEY_Sinh_ba
const KEY_Sinh_bha int32 = C.GDK_KEY_Sinh_bha
const KEY_Sinh_ca int32 = C.GDK_KEY_Sinh_ca
const KEY_Sinh_cha int32 = C.GDK_KEY_Sinh_cha
const KEY_Sinh_dda int32 = C.GDK_KEY_Sinh_dda
const KEY_Sinh_ddha int32 = C.GDK_KEY_Sinh_ddha
const KEY_Sinh_dha int32 = C.GDK_KEY_Sinh_dha
const KEY_Sinh_dhha int32 = C.GDK_KEY_Sinh_dhha
const KEY_Sinh_e int32 = C.GDK_KEY_Sinh_e
const KEY_Sinh_e2 int32 = C.GDK_KEY_Sinh_e2
const KEY_Sinh_ee int32 = C.GDK_KEY_Sinh_ee
const KEY_Sinh_ee2 int32 = C.GDK_KEY_Sinh_ee2
const KEY_Sinh_fa int32 = C.GDK_KEY_Sinh_fa
const KEY_Sinh_ga int32 = C.GDK_KEY_Sinh_ga
const KEY_Sinh_gha int32 = C.GDK_KEY_Sinh_gha
const KEY_Sinh_h2 int32 = C.GDK_KEY_Sinh_h2
const KEY_Sinh_ha int32 = C.GDK_KEY_Sinh_ha
const KEY_Sinh_i int32 = C.GDK_KEY_Sinh_i
const KEY_Sinh_i2 int32 = C.GDK_KEY_Sinh_i2
const KEY_Sinh_ii int32 = C.GDK_KEY_Sinh_ii
const KEY_Sinh_ii2 int32 = C.GDK_KEY_Sinh_ii2
const KEY_Sinh_ja int32 = C.GDK_KEY_Sinh_ja
const KEY_Sinh_jha int32 = C.GDK_KEY_Sinh_jha
const KEY_Sinh_jnya int32 = C.GDK_KEY_Sinh_jnya
const KEY_Sinh_ka int32 = C.GDK_KEY_Sinh_ka
const KEY_Sinh_kha int32 = C.GDK_KEY_Sinh_kha
const KEY_Sinh_kunddaliya int32 = C.GDK_KEY_Sinh_kunddaliya
const KEY_Sinh_la int32 = C.GDK_KEY_Sinh_la
const KEY_Sinh_lla int32 = C.GDK_KEY_Sinh_lla
const KEY_Sinh_lu int32 = C.GDK_KEY_Sinh_lu
const KEY_Sinh_lu2 int32 = C.GDK_KEY_Sinh_lu2
const KEY_Sinh_luu int32 = C.GDK_KEY_Sinh_luu
const KEY_Sinh_luu2 int32 = C.GDK_KEY_Sinh_luu2
const KEY_Sinh_ma int32 = C.GDK_KEY_Sinh_ma
const KEY_Sinh_mba int32 = C.GDK_KEY_Sinh_mba
const KEY_Sinh_na int32 = C.GDK_KEY_Sinh_na
const KEY_Sinh_ndda int32 = C.GDK_KEY_Sinh_ndda
const KEY_Sinh_ndha int32 = C.GDK_KEY_Sinh_ndha
const KEY_Sinh_ng int32 = C.GDK_KEY_Sinh_ng
const KEY_Sinh_ng2 int32 = C.GDK_KEY_Sinh_ng2
const KEY_Sinh_nga int32 = C.GDK_KEY_Sinh_nga
const KEY_Sinh_nja int32 = C.GDK_KEY_Sinh_nja
const KEY_Sinh_nna int32 = C.GDK_KEY_Sinh_nna
const KEY_Sinh_nya int32 = C.GDK_KEY_Sinh_nya
const KEY_Sinh_o int32 = C.GDK_KEY_Sinh_o
const KEY_Sinh_o2 int32 = C.GDK_KEY_Sinh_o2
const KEY_Sinh_oo int32 = C.GDK_KEY_Sinh_oo
const KEY_Sinh_oo2 int32 = C.GDK_KEY_Sinh_oo2
const KEY_Sinh_pa int32 = C.GDK_KEY_Sinh_pa
const KEY_Sinh_pha int32 = C.GDK_KEY_Sinh_pha
const KEY_Sinh_ra int32 = C.GDK_KEY_Sinh_ra
const KEY_Sinh_ri int32 = C.GDK_KEY_Sinh_ri
const KEY_Sinh_rii int32 = C.GDK_KEY_Sinh_rii
const KEY_Sinh_ru2 int32 = C.GDK_KEY_Sinh_ru2
const KEY_Sinh_ruu2 int32 = C.GDK_KEY_Sinh_ruu2
const KEY_Sinh_sa int32 = C.GDK_KEY_Sinh_sa
const KEY_Sinh_sha int32 = C.GDK_KEY_Sinh_sha
const KEY_Sinh_ssha int32 = C.GDK_KEY_Sinh_ssha
const KEY_Sinh_tha int32 = C.GDK_KEY_Sinh_tha
const KEY_Sinh_thha int32 = C.GDK_KEY_Sinh_thha
const KEY_Sinh_tta int32 = C.GDK_KEY_Sinh_tta
const KEY_Sinh_ttha int32 = C.GDK_KEY_Sinh_ttha
const KEY_Sinh_u int32 = C.GDK_KEY_Sinh_u
const KEY_Sinh_u2 int32 = C.GDK_KEY_Sinh_u2
const KEY_Sinh_uu int32 = C.GDK_KEY_Sinh_uu
const KEY_Sinh_uu2 int32 = C.GDK_KEY_Sinh_uu2
const KEY_Sinh_va int32 = C.GDK_KEY_Sinh_va
const KEY_Sinh_ya int32 = C.GDK_KEY_Sinh_ya
const KEY_Sleep int32 = C.GDK_KEY_Sleep
const KEY_SlowKeys_Enable int32 = C.GDK_KEY_SlowKeys_Enable
const KEY_Spell int32 = C.GDK_KEY_Spell
const KEY_SplitScreen int32 = C.GDK_KEY_SplitScreen
const KEY_Standby int32 = C.GDK_KEY_Standby
const KEY_Start int32 = C.GDK_KEY_Start
const KEY_StickyKeys_Enable int32 = C.GDK_KEY_StickyKeys_Enable
const KEY_Stop int32 = C.GDK_KEY_Stop
const KEY_Subtitle int32 = C.GDK_KEY_Subtitle
const KEY_Super_L int32 = C.GDK_KEY_Super_L
const KEY_Super_R int32 = C.GDK_KEY_Super_R
const KEY_Support int32 = C.GDK_KEY_Support
const KEY_Suspend int32 = C.GDK_KEY_Suspend
const KEY_Switch_VT_1 int32 = C.GDK_KEY_Switch_VT_1
const KEY_Switch_VT_10 int32 = C.GDK_KEY_Switch_VT_10
const KEY_Switch_VT_11 int32 = C.GDK_KEY_Switch_VT_11
const KEY_Switch_VT_12 int32 = C.GDK_KEY_Switch_VT_12
const KEY_Switch_VT_2 int32 = C.GDK_KEY_Switch_VT_2
const KEY_Switch_VT_3 int32 = C.GDK_KEY_Switch_VT_3
const KEY_Switch_VT_4 int32 = C.GDK_KEY_Switch_VT_4
const KEY_Switch_VT_5 int32 = C.GDK_KEY_Switch_VT_5
const KEY_Switch_VT_6 int32 = C.GDK_KEY_Switch_VT_6
const KEY_Switch_VT_7 int32 = C.GDK_KEY_Switch_VT_7
const KEY_Switch_VT_8 int32 = C.GDK_KEY_Switch_VT_8
const KEY_Switch_VT_9 int32 = C.GDK_KEY_Switch_VT_9
const KEY_Sys_Req int32 = C.GDK_KEY_Sys_Req
const KEY_T int32 = C.GDK_KEY_T
const KEY_THORN int32 = C.GDK_KEY_THORN
const KEY_Tab int32 = C.GDK_KEY_Tab
const KEY_Tabovedot int32 = C.GDK_KEY_Tabovedot
const KEY_TaskPane int32 = C.GDK_KEY_TaskPane
const KEY_Tcaron int32 = C.GDK_KEY_Tcaron
const KEY_Tcedilla int32 = C.GDK_KEY_Tcedilla
const KEY_Terminal int32 = C.GDK_KEY_Terminal
const KEY_Terminate_Server int32 = C.GDK_KEY_Terminate_Server
const KEY_Thai_baht int32 = C.GDK_KEY_Thai_baht
const KEY_Thai_bobaimai int32 = C.GDK_KEY_Thai_bobaimai
const KEY_Thai_chochan int32 = C.GDK_KEY_Thai_chochan
const KEY_Thai_chochang int32 = C.GDK_KEY_Thai_chochang
const KEY_Thai_choching int32 = C.GDK_KEY_Thai_choching
const KEY_Thai_chochoe int32 = C.GDK_KEY_Thai_chochoe
const KEY_Thai_dochada int32 = C.GDK_KEY_Thai_dochada
const KEY_Thai_dodek int32 = C.GDK_KEY_Thai_dodek
const KEY_Thai_fofa int32 = C.GDK_KEY_Thai_fofa
const KEY_Thai_fofan int32 = C.GDK_KEY_Thai_fofan
const KEY_Thai_hohip int32 = C.GDK_KEY_Thai_hohip
const KEY_Thai_honokhuk int32 = C.GDK_KEY_Thai_honokhuk
const KEY_Thai_khokhai int32 = C.GDK_KEY_Thai_khokhai
const KEY_Thai_khokhon int32 = C.GDK_KEY_Thai_khokhon
const KEY_Thai_khokhuat int32 = C.GDK_KEY_Thai_khokhuat
const KEY_Thai_khokhwai int32 = C.GDK_KEY_Thai_khokhwai
const KEY_Thai_khorakhang int32 = C.GDK_KEY_Thai_khorakhang
const KEY_Thai_kokai int32 = C.GDK_KEY_Thai_kokai
const KEY_Thai_lakkhangyao int32 = C.GDK_KEY_Thai_lakkhangyao
const KEY_Thai_lekchet int32 = C.GDK_KEY_Thai_lekchet
const KEY_Thai_lekha int32 = C.GDK_KEY_Thai_lekha
const KEY_Thai_lekhok int32 = C.GDK_KEY_Thai_lekhok
const KEY_Thai_lekkao int32 = C.GDK_KEY_Thai_lekkao
const KEY_Thai_leknung int32 = C.GDK_KEY_Thai_leknung
const KEY_Thai_lekpaet int32 = C.GDK_KEY_Thai_lekpaet
const KEY_Thai_leksam int32 = C.GDK_KEY_Thai_leksam
const KEY_Thai_leksi int32 = C.GDK_KEY_Thai_leksi
const KEY_Thai_leksong int32 = C.GDK_KEY_Thai_leksong
const KEY_Thai_leksun int32 = C.GDK_KEY_Thai_leksun
const KEY_Thai_lochula int32 = C.GDK_KEY_Thai_lochula
const KEY_Thai_loling int32 = C.GDK_KEY_Thai_loling
const KEY_Thai_lu int32 = C.GDK_KEY_Thai_lu
const KEY_Thai_maichattawa int32 = C.GDK_KEY_Thai_maichattawa
const KEY_Thai_maiek int32 = C.GDK_KEY_Thai_maiek
const KEY_Thai_maihanakat int32 = C.GDK_KEY_Thai_maihanakat
const KEY_Thai_maihanakat_maitho int32 = C.GDK_KEY_Thai_maihanakat_maitho
const KEY_Thai_maitaikhu int32 = C.GDK_KEY_Thai_maitaikhu
const KEY_Thai_maitho int32 = C.GDK_KEY_Thai_maitho
const KEY_Thai_maitri int32 = C.GDK_KEY_Thai_maitri
const KEY_Thai_maiyamok int32 = C.GDK_KEY_Thai_maiyamok
const KEY_Thai_moma int32 = C.GDK_KEY_Thai_moma
const KEY_Thai_ngongu int32 = C.GDK_KEY_Thai_ngongu
const KEY_Thai_nikhahit int32 = C.GDK_KEY_Thai_nikhahit
const KEY_Thai_nonen int32 = C.GDK_KEY_Thai_nonen
const KEY_Thai_nonu int32 = C.GDK_KEY_Thai_nonu
const KEY_Thai_oang int32 = C.GDK_KEY_Thai_oang
const KEY_Thai_paiyannoi int32 = C.GDK_KEY_Thai_paiyannoi
const KEY_Thai_phinthu int32 = C.GDK_KEY_Thai_phinthu
const KEY_Thai_phophan int32 = C.GDK_KEY_Thai_phophan
const KEY_Thai_phophung int32 = C.GDK_KEY_Thai_phophung
const KEY_Thai_phosamphao int32 = C.GDK_KEY_Thai_phosamphao
const KEY_Thai_popla int32 = C.GDK_KEY_Thai_popla
const KEY_Thai_rorua int32 = C.GDK_KEY_Thai_rorua
const KEY_Thai_ru int32 = C.GDK_KEY_Thai_ru
const KEY_Thai_saraa int32 = C.GDK_KEY_Thai_saraa
const KEY_Thai_saraaa int32 = C.GDK_KEY_Thai_saraaa
const KEY_Thai_saraae int32 = C.GDK_KEY_Thai_saraae
const KEY_Thai_saraaimaimalai int32 = C.GDK_KEY_Thai_saraaimaimalai
const KEY_Thai_saraaimaimuan int32 = C.GDK_KEY_Thai_saraaimaimuan
const KEY_Thai_saraam int32 = C.GDK_KEY_Thai_saraam
const KEY_Thai_sarae int32 = C.GDK_KEY_Thai_sarae
const KEY_Thai_sarai int32 = C.GDK_KEY_Thai_sarai
const KEY_Thai_saraii int32 = C.GDK_KEY_Thai_saraii
const KEY_Thai_sarao int32 = C.GDK_KEY_Thai_sarao
const KEY_Thai_sarau int32 = C.GDK_KEY_Thai_sarau
const KEY_Thai_saraue int32 = C.GDK_KEY_Thai_saraue
const KEY_Thai_sarauee int32 = C.GDK_KEY_Thai_sarauee
const KEY_Thai_sarauu int32 = C.GDK_KEY_Thai_sarauu
const KEY_Thai_sorusi int32 = C.GDK_KEY_Thai_sorusi
const KEY_Thai_sosala int32 = C.GDK_KEY_Thai_sosala
const KEY_Thai_soso int32 = C.GDK_KEY_Thai_soso
const KEY_Thai_sosua int32 = C.GDK_KEY_Thai_sosua
const KEY_Thai_thanthakhat int32 = C.GDK_KEY_Thai_thanthakhat
const KEY_Thai_thonangmontho int32 = C.GDK_KEY_Thai_thonangmontho
const KEY_Thai_thophuthao int32 = C.GDK_KEY_Thai_thophuthao
const KEY_Thai_thothahan int32 = C.GDK_KEY_Thai_thothahan
const KEY_Thai_thothan int32 = C.GDK_KEY_Thai_thothan
const KEY_Thai_thothong int32 = C.GDK_KEY_Thai_thothong
const KEY_Thai_thothung int32 = C.GDK_KEY_Thai_thothung
const KEY_Thai_topatak int32 = C.GDK_KEY_Thai_topatak
const KEY_Thai_totao int32 = C.GDK_KEY_Thai_totao
const KEY_Thai_wowaen int32 = C.GDK_KEY_Thai_wowaen
const KEY_Thai_yoyak int32 = C.GDK_KEY_Thai_yoyak
const KEY_Thai_yoying int32 = C.GDK_KEY_Thai_yoying
const KEY_Thorn int32 = C.GDK_KEY_Thorn
const KEY_Time int32 = C.GDK_KEY_Time
const KEY_ToDoList int32 = C.GDK_KEY_ToDoList
const KEY_Tools int32 = C.GDK_KEY_Tools
const KEY_TopMenu int32 = C.GDK_KEY_TopMenu
const KEY_TouchpadOff int32 = C.GDK_KEY_TouchpadOff
const KEY_TouchpadOn int32 = C.GDK_KEY_TouchpadOn
const KEY_TouchpadToggle int32 = C.GDK_KEY_TouchpadToggle
const KEY_Touroku int32 = C.GDK_KEY_Touroku
const KEY_Travel int32 = C.GDK_KEY_Travel
const KEY_Tslash int32 = C.GDK_KEY_Tslash
const KEY_U int32 = C.GDK_KEY_U
const KEY_UWB int32 = C.GDK_KEY_UWB
const KEY_Uacute int32 = C.GDK_KEY_Uacute
const KEY_Ubelowdot int32 = C.GDK_KEY_Ubelowdot
const KEY_Ubreve int32 = C.GDK_KEY_Ubreve
const KEY_Ucircumflex int32 = C.GDK_KEY_Ucircumflex
const KEY_Udiaeresis int32 = C.GDK_KEY_Udiaeresis
const KEY_Udoubleacute int32 = C.GDK_KEY_Udoubleacute
const KEY_Ugrave int32 = C.GDK_KEY_Ugrave
const KEY_Uhook int32 = C.GDK_KEY_Uhook
const KEY_Uhorn int32 = C.GDK_KEY_Uhorn
const KEY_Uhornacute int32 = C.GDK_KEY_Uhornacute
const KEY_Uhornbelowdot int32 = C.GDK_KEY_Uhornbelowdot
const KEY_Uhorngrave int32 = C.GDK_KEY_Uhorngrave
const KEY_Uhornhook int32 = C.GDK_KEY_Uhornhook
const KEY_Uhorntilde int32 = C.GDK_KEY_Uhorntilde
const KEY_Ukrainian_GHE_WITH_UPTURN int32 = C.GDK_KEY_Ukrainian_GHE_WITH_UPTURN
const KEY_Ukrainian_I int32 = C.GDK_KEY_Ukrainian_I
const KEY_Ukrainian_IE int32 = C.GDK_KEY_Ukrainian_IE
const KEY_Ukrainian_YI int32 = C.GDK_KEY_Ukrainian_YI
const KEY_Ukrainian_ghe_with_upturn int32 = C.GDK_KEY_Ukrainian_ghe_with_upturn
const KEY_Ukrainian_i int32 = C.GDK_KEY_Ukrainian_i
const KEY_Ukrainian_ie int32 = C.GDK_KEY_Ukrainian_ie
const KEY_Ukrainian_yi int32 = C.GDK_KEY_Ukrainian_yi
const KEY_Ukranian_I int32 = C.GDK_KEY_Ukranian_I
const KEY_Ukranian_JE int32 = C.GDK_KEY_Ukranian_JE
const KEY_Ukranian_YI int32 = C.GDK_KEY_Ukranian_YI
const KEY_Ukranian_i int32 = C.GDK_KEY_Ukranian_i
const KEY_Ukranian_je int32 = C.GDK_KEY_Ukranian_je
const KEY_Ukranian_yi int32 = C.GDK_KEY_Ukranian_yi
const KEY_Umacron int32 = C.GDK_KEY_Umacron
const KEY_Undo int32 = C.GDK_KEY_Undo
const KEY_Ungrab int32 = C.GDK_KEY_Ungrab
const KEY_Uogonek int32 = C.GDK_KEY_Uogonek
const KEY_Up int32 = C.GDK_KEY_Up
const KEY_Uring int32 = C.GDK_KEY_Uring
const KEY_User1KB int32 = C.GDK_KEY_User1KB
const KEY_User2KB int32 = C.GDK_KEY_User2KB
const KEY_UserPB int32 = C.GDK_KEY_UserPB
const KEY_Utilde int32 = C.GDK_KEY_Utilde
const KEY_V int32 = C.GDK_KEY_V
const KEY_VendorHome int32 = C.GDK_KEY_VendorHome
const KEY_Video int32 = C.GDK_KEY_Video
const KEY_View int32 = C.GDK_KEY_View
const KEY_VoidSymbol int32 = C.GDK_KEY_VoidSymbol
const KEY_W int32 = C.GDK_KEY_W
const KEY_WLAN int32 = C.GDK_KEY_WLAN
const KEY_WWW int32 = C.GDK_KEY_WWW
const KEY_Wacute int32 = C.GDK_KEY_Wacute
const KEY_WakeUp int32 = C.GDK_KEY_WakeUp
const KEY_Wcircumflex int32 = C.GDK_KEY_Wcircumflex
const KEY_Wdiaeresis int32 = C.GDK_KEY_Wdiaeresis
const KEY_WebCam int32 = C.GDK_KEY_WebCam
const KEY_Wgrave int32 = C.GDK_KEY_Wgrave
const KEY_WheelButton int32 = C.GDK_KEY_WheelButton
const KEY_WindowClear int32 = C.GDK_KEY_WindowClear
const KEY_WonSign int32 = C.GDK_KEY_WonSign
const KEY_Word int32 = C.GDK_KEY_Word
const KEY_X int32 = C.GDK_KEY_X
const KEY_Xabovedot int32 = C.GDK_KEY_Xabovedot
const KEY_Xfer int32 = C.GDK_KEY_Xfer
const KEY_Y int32 = C.GDK_KEY_Y
const KEY_Yacute int32 = C.GDK_KEY_Yacute
const KEY_Ybelowdot int32 = C.GDK_KEY_Ybelowdot
const KEY_Ycircumflex int32 = C.GDK_KEY_Ycircumflex
const KEY_Ydiaeresis int32 = C.GDK_KEY_Ydiaeresis
const KEY_Yellow int32 = C.GDK_KEY_Yellow
const KEY_Ygrave int32 = C.GDK_KEY_Ygrave
const KEY_Yhook int32 = C.GDK_KEY_Yhook
const KEY_Ytilde int32 = C.GDK_KEY_Ytilde
const KEY_Z int32 = C.GDK_KEY_Z
const KEY_Zabovedot int32 = C.GDK_KEY_Zabovedot
const KEY_Zacute int32 = C.GDK_KEY_Zacute
const KEY_Zcaron int32 = C.GDK_KEY_Zcaron
const KEY_Zen_Koho int32 = C.GDK_KEY_Zen_Koho
const KEY_Zenkaku int32 = C.GDK_KEY_Zenkaku
const KEY_Zenkaku_Hankaku int32 = C.GDK_KEY_Zenkaku_Hankaku
const KEY_ZoomIn int32 = C.GDK_KEY_ZoomIn
const KEY_ZoomOut int32 = C.GDK_KEY_ZoomOut
const KEY_Zstroke int32 = C.GDK_KEY_Zstroke
const KEY_a int32 = C.GDK_KEY_a
const KEY_aacute int32 = C.GDK_KEY_aacute
const KEY_abelowdot int32 = C.GDK_KEY_abelowdot
const KEY_abovedot int32 = C.GDK_KEY_abovedot
const KEY_abreve int32 = C.GDK_KEY_abreve
const KEY_abreveacute int32 = C.GDK_KEY_abreveacute
const KEY_abrevebelowdot int32 = C.GDK_KEY_abrevebelowdot
const KEY_abrevegrave int32 = C.GDK_KEY_abrevegrave
const KEY_abrevehook int32 = C.GDK_KEY_abrevehook
const KEY_abrevetilde int32 = C.GDK_KEY_abrevetilde
const KEY_acircumflex int32 = C.GDK_KEY_acircumflex
const KEY_acircumflexacute int32 = C.GDK_KEY_acircumflexacute
const KEY_acircumflexbelowdot int32 = C.GDK_KEY_acircumflexbelowdot
const KEY_acircumflexgrave int32 = C.GDK_KEY_acircumflexgrave
const KEY_acircumflexhook int32 = C.GDK_KEY_acircumflexhook
const KEY_acircumflextilde int32 = C.GDK_KEY_acircumflextilde
const KEY_acute int32 = C.GDK_KEY_acute
const KEY_adiaeresis int32 = C.GDK_KEY_adiaeresis
const KEY_ae int32 = C.GDK_KEY_ae
const KEY_agrave int32 = C.GDK_KEY_agrave
const KEY_ahook int32 = C.GDK_KEY_ahook
const KEY_amacron int32 = C.GDK_KEY_amacron
const KEY_ampersand int32 = C.GDK_KEY_ampersand
const KEY_aogonek int32 = C.GDK_KEY_aogonek
const KEY_apostrophe int32 = C.GDK_KEY_apostrophe
const KEY_approxeq int32 = C.GDK_KEY_approxeq
const KEY_approximate int32 = C.GDK_KEY_approximate
const KEY_aring int32 = C.GDK_KEY_aring
const KEY_asciicircum int32 = C.GDK_KEY_asciicircum
const KEY_asciitilde int32 = C.GDK_KEY_asciitilde
const KEY_asterisk int32 = C.GDK_KEY_asterisk
const KEY_at int32 = C.GDK_KEY_at
const KEY_atilde int32 = C.GDK_KEY_atilde
const KEY_b int32 = C.GDK_KEY_b
const KEY_babovedot int32 = C.GDK_KEY_babovedot
const KEY_backslash int32 = C.GDK_KEY_backslash
const KEY_ballotcross int32 = C.GDK_KEY_ballotcross
const KEY_bar int32 = C.GDK_KEY_bar
const KEY_because int32 = C.GDK_KEY_because
const KEY_blank int32 = C.GDK_KEY_blank
const KEY_botintegral int32 = C.GDK_KEY_botintegral
const KEY_botleftparens int32 = C.GDK_KEY_botleftparens
const KEY_botleftsqbracket int32 = C.GDK_KEY_botleftsqbracket
const KEY_botleftsummation int32 = C.GDK_KEY_botleftsummation
const KEY_botrightparens int32 = C.GDK_KEY_botrightparens
const KEY_botrightsqbracket int32 = C.GDK_KEY_botrightsqbracket
const KEY_botrightsummation int32 = C.GDK_KEY_botrightsummation
const KEY_bott int32 = C.GDK_KEY_bott
const KEY_botvertsummationconnector int32 = C.GDK_KEY_botvertsummationconnector
const KEY_braceleft int32 = C.GDK_KEY_braceleft
const KEY_braceright int32 = C.GDK_KEY_braceright
const KEY_bracketleft int32 = C.GDK_KEY_bracketleft
const KEY_bracketright int32 = C.GDK_KEY_bracketright
const KEY_braille_blank int32 = C.GDK_KEY_braille_blank
const KEY_braille_dot_1 int32 = C.GDK_KEY_braille_dot_1
const KEY_braille_dot_10 int32 = C.GDK_KEY_braille_dot_10
const KEY_braille_dot_2 int32 = C.GDK_KEY_braille_dot_2
const KEY_braille_dot_3 int32 = C.GDK_KEY_braille_dot_3
const KEY_braille_dot_4 int32 = C.GDK_KEY_braille_dot_4
const KEY_braille_dot_5 int32 = C.GDK_KEY_braille_dot_5
const KEY_braille_dot_6 int32 = C.GDK_KEY_braille_dot_6
const KEY_braille_dot_7 int32 = C.GDK_KEY_braille_dot_7
const KEY_braille_dot_8 int32 = C.GDK_KEY_braille_dot_8
const KEY_braille_dot_9 int32 = C.GDK_KEY_braille_dot_9
const KEY_braille_dots_1 int32 = C.GDK_KEY_braille_dots_1
const KEY_braille_dots_12 int32 = C.GDK_KEY_braille_dots_12
const KEY_braille_dots_123 int32 = C.GDK_KEY_braille_dots_123
const KEY_braille_dots_1234 int32 = C.GDK_KEY_braille_dots_1234
const KEY_braille_dots_12345 int32 = C.GDK_KEY_braille_dots_12345
const KEY_braille_dots_123456 int32 = C.GDK_KEY_braille_dots_123456
const KEY_braille_dots_1234567 int32 = C.GDK_KEY_braille_dots_1234567
const KEY_braille_dots_12345678 int32 = C.GDK_KEY_braille_dots_12345678
const KEY_braille_dots_1234568 int32 = C.GDK_KEY_braille_dots_1234568
const KEY_braille_dots_123457 int32 = C.GDK_KEY_braille_dots_123457
const KEY_braille_dots_1234578 int32 = C.GDK_KEY_braille_dots_1234578
const KEY_braille_dots_123458 int32 = C.GDK_KEY_braille_dots_123458
const KEY_braille_dots_12346 int32 = C.GDK_KEY_braille_dots_12346
const KEY_braille_dots_123467 int32 = C.GDK_KEY_braille_dots_123467
const KEY_braille_dots_1234678 int32 = C.GDK_KEY_braille_dots_1234678
const KEY_braille_dots_123468 int32 = C.GDK_KEY_braille_dots_123468
const KEY_braille_dots_12347 int32 = C.GDK_KEY_braille_dots_12347
const KEY_braille_dots_123478 int32 = C.GDK_KEY_braille_dots_123478
const KEY_braille_dots_12348 int32 = C.GDK_KEY_braille_dots_12348
const KEY_braille_dots_1235 int32 = C.GDK_KEY_braille_dots_1235
const KEY_braille_dots_12356 int32 = C.GDK_KEY_braille_dots_12356
const KEY_braille_dots_123567 int32 = C.GDK_KEY_braille_dots_123567
const KEY_braille_dots_1235678 int32 = C.GDK_KEY_braille_dots_1235678
const KEY_braille_dots_123568 int32 = C.GDK_KEY_braille_dots_123568
const KEY_braille_dots_12357 int32 = C.GDK_KEY_braille_dots_12357
const KEY_braille_dots_123578 int32 = C.GDK_KEY_braille_dots_123578
const KEY_braille_dots_12358 int32 = C.GDK_KEY_braille_dots_12358
const KEY_braille_dots_1236 int32 = C.GDK_KEY_braille_dots_1236
const KEY_braille_dots_12367 int32 = C.GDK_KEY_braille_dots_12367
const KEY_braille_dots_123678 int32 = C.GDK_KEY_braille_dots_123678
const KEY_braille_dots_12368 int32 = C.GDK_KEY_braille_dots_12368
const KEY_braille_dots_1237 int32 = C.GDK_KEY_braille_dots_1237
const KEY_braille_dots_12378 int32 = C.GDK_KEY_braille_dots_12378
const KEY_braille_dots_1238 int32 = C.GDK_KEY_braille_dots_1238
const KEY_braille_dots_124 int32 = C.GDK_KEY_braille_dots_124
const KEY_braille_dots_1245 int32 = C.GDK_KEY_braille_dots_1245
const KEY_braille_dots_12456 int32 = C.GDK_KEY_braille_dots_12456
const KEY_braille_dots_124567 int32 = C.GDK_KEY_braille_dots_124567
const KEY_braille_dots_1245678 int32 = C.GDK_KEY_braille_dots_1245678
const KEY_braille_dots_124568 int32 = C.GDK_KEY_braille_dots_124568
const KEY_braille_dots_12457 int32 = C.GDK_KEY_braille_dots_12457
const KEY_braille_dots_124578 int32 = C.GDK_KEY_braille_dots_124578
const KEY_braille_dots_12458 int32 = C.GDK_KEY_braille_dots_12458
const KEY_braille_dots_1246 int32 = C.GDK_KEY_braille_dots_1246
const KEY_braille_dots_12467 int32 = C.GDK_KEY_braille_dots_12467
const KEY_braille_dots_124678 int32 = C.GDK_KEY_braille_dots_124678
const KEY_braille_dots_12468 int32 = C.GDK_KEY_braille_dots_12468
const KEY_braille_dots_1247 int32 = C.GDK_KEY_braille_dots_1247
const KEY_braille_dots_12478 int32 = C.GDK_KEY_braille_dots_12478
const KEY_braille_dots_1248 int32 = C.GDK_KEY_braille_dots_1248
const KEY_braille_dots_125 int32 = C.GDK_KEY_braille_dots_125
const KEY_braille_dots_1256 int32 = C.GDK_KEY_braille_dots_1256
const KEY_braille_dots_12567 int32 = C.GDK_KEY_braille_dots_12567
const KEY_braille_dots_125678 int32 = C.GDK_KEY_braille_dots_125678
const KEY_braille_dots_12568 int32 = C.GDK_KEY_braille_dots_12568
const KEY_braille_dots_1257 int32 = C.GDK_KEY_braille_dots_1257
const KEY_braille_dots_12578 int32 = C.GDK_KEY_braille_dots_12578
const KEY_braille_dots_1258 int32 = C.GDK_KEY_braille_dots_1258
const KEY_braille_dots_126 int32 = C.GDK_KEY_braille_dots_126
const KEY_braille_dots_1267 int32 = C.GDK_KEY_braille_dots_1267
const KEY_braille_dots_12678 int32 = C.GDK_KEY_braille_dots_12678
const KEY_braille_dots_1268 int32 = C.GDK_KEY_braille_dots_1268
const KEY_braille_dots_127 int32 = C.GDK_KEY_braille_dots_127
const KEY_braille_dots_1278 int32 = C.GDK_KEY_braille_dots_1278
const KEY_braille_dots_128 int32 = C.GDK_KEY_braille_dots_128
const KEY_braille_dots_13 int32 = C.GDK_KEY_braille_dots_13
const KEY_braille_dots_134 int32 = C.GDK_KEY_braille_dots_134
const KEY_braille_dots_1345 int32 = C.GDK_KEY_braille_dots_1345
const KEY_braille_dots_13456 int32 = C.GDK_KEY_braille_dots_13456
const KEY_braille_dots_134567 int32 = C.GDK_KEY_braille_dots_134567
const KEY_braille_dots_1345678 int32 = C.GDK_KEY_braille_dots_1345678
const KEY_braille_dots_134568 int32 = C.GDK_KEY_braille_dots_134568
const KEY_braille_dots_13457 int32 = C.GDK_KEY_braille_dots_13457
const KEY_braille_dots_134578 int32 = C.GDK_KEY_braille_dots_134578
const KEY_braille_dots_13458 int32 = C.GDK_KEY_braille_dots_13458
const KEY_braille_dots_1346 int32 = C.GDK_KEY_braille_dots_1346
const KEY_braille_dots_13467 int32 = C.GDK_KEY_braille_dots_13467
const KEY_braille_dots_134678 int32 = C.GDK_KEY_braille_dots_134678
const KEY_braille_dots_13468 int32 = C.GDK_KEY_braille_dots_13468
const KEY_braille_dots_1347 int32 = C.GDK_KEY_braille_dots_1347
const KEY_braille_dots_13478 int32 = C.GDK_KEY_braille_dots_13478
const KEY_braille_dots_1348 int32 = C.GDK_KEY_braille_dots_1348
const KEY_braille_dots_135 int32 = C.GDK_KEY_braille_dots_135
const KEY_braille_dots_1356 int32 = C.GDK_KEY_braille_dots_1356
const KEY_braille_dots_13567 int32 = C.GDK_KEY_braille_dots_13567
const KEY_braille_dots_135678 int32 = C.GDK_KEY_braille_dots_135678
const KEY_braille_dots_13568 int32 = C.GDK_KEY_braille_dots_13568
const KEY_braille_dots_1357 int32 = C.GDK_KEY_braille_dots_1357
const KEY_braille_dots_13578 int32 = C.GDK_KEY_braille_dots_13578
const KEY_braille_dots_1358 int32 = C.GDK_KEY_braille_dots_1358
const KEY_braille_dots_136 int32 = C.GDK_KEY_braille_dots_136
const KEY_braille_dots_1367 int32 = C.GDK_KEY_braille_dots_1367
const KEY_braille_dots_13678 int32 = C.GDK_KEY_braille_dots_13678
const KEY_braille_dots_1368 int32 = C.GDK_KEY_braille_dots_1368
const KEY_braille_dots_137 int32 = C.GDK_KEY_braille_dots_137
const KEY_braille_dots_1378 int32 = C.GDK_KEY_braille_dots_1378
const KEY_braille_dots_138 int32 = C.GDK_KEY_braille_dots_138
const KEY_braille_dots_14 int32 = C.GDK_KEY_braille_dots_14
const KEY_braille_dots_145 int32 = C.GDK_KEY_braille_dots_145
const KEY_braille_dots_1456 int32 = C.GDK_KEY_braille_dots_1456
const KEY_braille_dots_14567 int32 = C.GDK_KEY_braille_dots_14567
const KEY_braille_dots_145678 int32 = C.GDK_KEY_braille_dots_145678
const KEY_braille_dots_14568 int32 = C.GDK_KEY_braille_dots_14568
const KEY_braille_dots_1457 int32 = C.GDK_KEY_braille_dots_1457
const KEY_braille_dots_14578 int32 = C.GDK_KEY_braille_dots_14578
const KEY_braille_dots_1458 int32 = C.GDK_KEY_braille_dots_1458
const KEY_braille_dots_146 int32 = C.GDK_KEY_braille_dots_146
const KEY_braille_dots_1467 int32 = C.GDK_KEY_braille_dots_1467
const KEY_braille_dots_14678 int32 = C.GDK_KEY_braille_dots_14678
const KEY_braille_dots_1468 int32 = C.GDK_KEY_braille_dots_1468
const KEY_braille_dots_147 int32 = C.GDK_KEY_braille_dots_147
const KEY_braille_dots_1478 int32 = C.GDK_KEY_braille_dots_1478
const KEY_braille_dots_148 int32 = C.GDK_KEY_braille_dots_148
const KEY_braille_dots_15 int32 = C.GDK_KEY_braille_dots_15
const KEY_braille_dots_156 int32 = C.GDK_KEY_braille_dots_156
const KEY_braille_dots_1567 int32 = C.GDK_KEY_braille_dots_1567
const KEY_braille_dots_15678 int32 = C.GDK_KEY_braille_dots_15678
const KEY_braille_dots_1568 int32 = C.GDK_KEY_braille_dots_1568
const KEY_braille_dots_157 int32 = C.GDK_KEY_braille_dots_157
const KEY_braille_dots_1578 int32 = C.GDK_KEY_braille_dots_1578
const KEY_braille_dots_158 int32 = C.GDK_KEY_braille_dots_158
const KEY_braille_dots_16 int32 = C.GDK_KEY_braille_dots_16
const KEY_braille_dots_167 int32 = C.GDK_KEY_braille_dots_167
const KEY_braille_dots_1678 int32 = C.GDK_KEY_braille_dots_1678
const KEY_braille_dots_168 int32 = C.GDK_KEY_braille_dots_168
const KEY_braille_dots_17 int32 = C.GDK_KEY_braille_dots_17
const KEY_braille_dots_178 int32 = C.GDK_KEY_braille_dots_178
const KEY_braille_dots_18 int32 = C.GDK_KEY_braille_dots_18
const KEY_braille_dots_2 int32 = C.GDK_KEY_braille_dots_2
const KEY_braille_dots_23 int32 = C.GDK_KEY_braille_dots_23
const KEY_braille_dots_234 int32 = C.GDK_KEY_braille_dots_234
const KEY_braille_dots_2345 int32 = C.GDK_KEY_braille_dots_2345
const KEY_braille_dots_23456 int32 = C.GDK_KEY_braille_dots_23456
const KEY_braille_dots_234567 int32 = C.GDK_KEY_braille_dots_234567
const KEY_braille_dots_2345678 int32 = C.GDK_KEY_braille_dots_2345678
const KEY_braille_dots_234568 int32 = C.GDK_KEY_braille_dots_234568
const KEY_braille_dots_23457 int32 = C.GDK_KEY_braille_dots_23457
const KEY_braille_dots_234578 int32 = C.GDK_KEY_braille_dots_234578
const KEY_braille_dots_23458 int32 = C.GDK_KEY_braille_dots_23458
const KEY_braille_dots_2346 int32 = C.GDK_KEY_braille_dots_2346
const KEY_braille_dots_23467 int32 = C.GDK_KEY_braille_dots_23467
const KEY_braille_dots_234678 int32 = C.GDK_KEY_braille_dots_234678
const KEY_braille_dots_23468 int32 = C.GDK_KEY_braille_dots_23468
const KEY_braille_dots_2347 int32 = C.GDK_KEY_braille_dots_2347
const KEY_braille_dots_23478 int32 = C.GDK_KEY_braille_dots_23478
const KEY_braille_dots_2348 int32 = C.GDK_KEY_braille_dots_2348
const KEY_braille_dots_235 int32 = C.GDK_KEY_braille_dots_235
const KEY_braille_dots_2356 int32 = C.GDK_KEY_braille_dots_2356
const KEY_braille_dots_23567 int32 = C.GDK_KEY_braille_dots_23567
const KEY_braille_dots_235678 int32 = C.GDK_KEY_braille_dots_235678
const KEY_braille_dots_23568 int32 = C.GDK_KEY_braille_dots_23568
const KEY_braille_dots_2357 int32 = C.GDK_KEY_braille_dots_2357
const KEY_braille_dots_23578 int32 = C.GDK_KEY_braille_dots_23578
const KEY_braille_dots_2358 int32 = C.GDK_KEY_braille_dots_2358
const KEY_braille_dots_236 int32 = C.GDK_KEY_braille_dots_236
const KEY_braille_dots_2367 int32 = C.GDK_KEY_braille_dots_2367
const KEY_braille_dots_23678 int32 = C.GDK_KEY_braille_dots_23678
const KEY_braille_dots_2368 int32 = C.GDK_KEY_braille_dots_2368
const KEY_braille_dots_237 int32 = C.GDK_KEY_braille_dots_237
const KEY_braille_dots_2378 int32 = C.GDK_KEY_braille_dots_2378
const KEY_braille_dots_238 int32 = C.GDK_KEY_braille_dots_238
const KEY_braille_dots_24 int32 = C.GDK_KEY_braille_dots_24
const KEY_braille_dots_245 int32 = C.GDK_KEY_braille_dots_245
const KEY_braille_dots_2456 int32 = C.GDK_KEY_braille_dots_2456
const KEY_braille_dots_24567 int32 = C.GDK_KEY_braille_dots_24567
const KEY_braille_dots_245678 int32 = C.GDK_KEY_braille_dots_245678
const KEY_braille_dots_24568 int32 = C.GDK_KEY_braille_dots_24568
const KEY_braille_dots_2457 int32 = C.GDK_KEY_braille_dots_2457
const KEY_braille_dots_24578 int32 = C.GDK_KEY_braille_dots_24578
const KEY_braille_dots_2458 int32 = C.GDK_KEY_braille_dots_2458
const KEY_braille_dots_246 int32 = C.GDK_KEY_braille_dots_246
const KEY_braille_dots_2467 int32 = C.GDK_KEY_braille_dots_2467
const KEY_braille_dots_24678 int32 = C.GDK_KEY_braille_dots_24678
const KEY_braille_dots_2468 int32 = C.GDK_KEY_braille_dots_2468
const KEY_braille_dots_247 int32 = C.GDK_KEY_braille_dots_247
const KEY_braille_dots_2478 int32 = C.GDK_KEY_braille_dots_2478
const KEY_braille_dots_248 int32 = C.GDK_KEY_braille_dots_248
const KEY_braille_dots_25 int32 = C.GDK_KEY_braille_dots_25
const KEY_braille_dots_256 int32 = C.GDK_KEY_braille_dots_256
const KEY_braille_dots_2567 int32 = C.GDK_KEY_braille_dots_2567
const KEY_braille_dots_25678 int32 = C.GDK_KEY_braille_dots_25678
const KEY_braille_dots_2568 int32 = C.GDK_KEY_braille_dots_2568
const KEY_braille_dots_257 int32 = C.GDK_KEY_braille_dots_257
const KEY_braille_dots_2578 int32 = C.GDK_KEY_braille_dots_2578
const KEY_braille_dots_258 int32 = C.GDK_KEY_braille_dots_258
const KEY_braille_dots_26 int32 = C.GDK_KEY_braille_dots_26
const KEY_braille_dots_267 int32 = C.GDK_KEY_braille_dots_267
const KEY_braille_dots_2678 int32 = C.GDK_KEY_braille_dots_2678
const KEY_braille_dots_268 int32 = C.GDK_KEY_braille_dots_268
const KEY_braille_dots_27 int32 = C.GDK_KEY_braille_dots_27
const KEY_braille_dots_278 int32 = C.GDK_KEY_braille_dots_278
const KEY_braille_dots_28 int32 = C.GDK_KEY_braille_dots_28
const KEY_braille_dots_3 int32 = C.GDK_KEY_braille_dots_3
const KEY_braille_dots_34 int32 = C.GDK_KEY_braille_dots_34
const KEY_braille_dots_345 int32 = C.GDK_KEY_braille_dots_345
const KEY_braille_dots_3456 int32 = C.GDK_KEY_braille_dots_3456
const KEY_braille_dots_34567 int32 = C.GDK_KEY_braille_dots_34567
const KEY_braille_dots_345678 int32 = C.GDK_KEY_braille_dots_345678
const KEY_braille_dots_34568 int32 = C.GDK_KEY_braille_dots_34568
const KEY_braille_dots_3457 int32 = C.GDK_KEY_braille_dots_3457
const KEY_braille_dots_34578 int32 = C.GDK_KEY_braille_dots_34578
const KEY_braille_dots_3458 int32 = C.GDK_KEY_braille_dots_3458
const KEY_braille_dots_346 int32 = C.GDK_KEY_braille_dots_346
const KEY_braille_dots_3467 int32 = C.GDK_KEY_braille_dots_3467
const KEY_braille_dots_34678 int32 = C.GDK_KEY_braille_dots_34678
const KEY_braille_dots_3468 int32 = C.GDK_KEY_braille_dots_3468
const KEY_braille_dots_347 int32 = C.GDK_KEY_braille_dots_347
const KEY_braille_dots_3478 int32 = C.GDK_KEY_braille_dots_3478
const KEY_braille_dots_348 int32 = C.GDK_KEY_braille_dots_348
const KEY_braille_dots_35 int32 = C.GDK_KEY_braille_dots_35
const KEY_braille_dots_356 int32 = C.GDK_KEY_braille_dots_356
const KEY_braille_dots_3567 int32 = C.GDK_KEY_braille_dots_3567
const KEY_braille_dots_35678 int32 = C.GDK_KEY_braille_dots_35678
const KEY_braille_dots_3568 int32 = C.GDK_KEY_braille_dots_3568
const KEY_braille_dots_357 int32 = C.GDK_KEY_braille_dots_357
const KEY_braille_dots_3578 int32 = C.GDK_KEY_braille_dots_3578
const KEY_braille_dots_358 int32 = C.GDK_KEY_braille_dots_358
const KEY_braille_dots_36 int32 = C.GDK_KEY_braille_dots_36
const KEY_braille_dots_367 int32 = C.GDK_KEY_braille_dots_367
const KEY_braille_dots_3678 int32 = C.GDK_KEY_braille_dots_3678
const KEY_braille_dots_368 int32 = C.GDK_KEY_braille_dots_368
const KEY_braille_dots_37 int32 = C.GDK_KEY_braille_dots_37
const KEY_braille_dots_378 int32 = C.GDK_KEY_braille_dots_378
const KEY_braille_dots_38 int32 = C.GDK_KEY_braille_dots_38
const KEY_braille_dots_4 int32 = C.GDK_KEY_braille_dots_4
const KEY_braille_dots_45 int32 = C.GDK_KEY_braille_dots_45
const KEY_braille_dots_456 int32 = C.GDK_KEY_braille_dots_456
const KEY_braille_dots_4567 int32 = C.GDK_KEY_braille_dots_4567
const KEY_braille_dots_45678 int32 = C.GDK_KEY_braille_dots_45678
const KEY_braille_dots_4568 int32 = C.GDK_KEY_braille_dots_4568
const KEY_braille_dots_457 int32 = C.GDK_KEY_braille_dots_457
const KEY_braille_dots_4578 int32 = C.GDK_KEY_braille_dots_4578
const KEY_braille_dots_458 int32 = C.GDK_KEY_braille_dots_458
const KEY_braille_dots_46 int32 = C.GDK_KEY_braille_dots_46
const KEY_braille_dots_467 int32 = C.GDK_KEY_braille_dots_467
const KEY_braille_dots_4678 int32 = C.GDK_KEY_braille_dots_4678
const KEY_braille_dots_468 int32 = C.GDK_KEY_braille_dots_468
const KEY_braille_dots_47 int32 = C.GDK_KEY_braille_dots_47
const KEY_braille_dots_478 int32 = C.GDK_KEY_braille_dots_478
const KEY_braille_dots_48 int32 = C.GDK_KEY_braille_dots_48
const KEY_braille_dots_5 int32 = C.GDK_KEY_braille_dots_5
const KEY_braille_dots_56 int32 = C.GDK_KEY_braille_dots_56
const KEY_braille_dots_567 int32 = C.GDK_KEY_braille_dots_567
const KEY_braille_dots_5678 int32 = C.GDK_KEY_braille_dots_5678
const KEY_braille_dots_568 int32 = C.GDK_KEY_braille_dots_568
const KEY_braille_dots_57 int32 = C.GDK_KEY_braille_dots_57
const KEY_braille_dots_578 int32 = C.GDK_KEY_braille_dots_578
const KEY_braille_dots_58 int32 = C.GDK_KEY_braille_dots_58
const KEY_braille_dots_6 int32 = C.GDK_KEY_braille_dots_6
const KEY_braille_dots_67 int32 = C.GDK_KEY_braille_dots_67
const KEY_braille_dots_678 int32 = C.GDK_KEY_braille_dots_678
const KEY_braille_dots_68 int32 = C.GDK_KEY_braille_dots_68
const KEY_braille_dots_7 int32 = C.GDK_KEY_braille_dots_7
const KEY_braille_dots_78 int32 = C.GDK_KEY_braille_dots_78
const KEY_braille_dots_8 int32 = C.GDK_KEY_braille_dots_8
const KEY_breve int32 = C.GDK_KEY_breve
const KEY_brokenbar int32 = C.GDK_KEY_brokenbar
const KEY_c int32 = C.GDK_KEY_c
const KEY_c_h int32 = C.GDK_KEY_c_h
const KEY_cabovedot int32 = C.GDK_KEY_cabovedot
const KEY_cacute int32 = C.GDK_KEY_cacute
const KEY_careof int32 = C.GDK_KEY_careof
const KEY_caret int32 = C.GDK_KEY_caret
const KEY_caron int32 = C.GDK_KEY_caron
const KEY_ccaron int32 = C.GDK_KEY_ccaron
const KEY_ccedilla int32 = C.GDK_KEY_ccedilla
const KEY_ccircumflex int32 = C.GDK_KEY_ccircumflex
const KEY_cedilla int32 = C.GDK_KEY_cedilla
const KEY_cent int32 = C.GDK_KEY_cent
const KEY_ch int32 = C.GDK_KEY_ch
const KEY_checkerboard int32 = C.GDK_KEY_checkerboard
const KEY_checkmark int32 = C.GDK_KEY_checkmark
const KEY_circle int32 = C.GDK_KEY_circle
const KEY_club int32 = C.GDK_KEY_club
const KEY_colon int32 = C.GDK_KEY_colon
const KEY_comma int32 = C.GDK_KEY_comma
const KEY_containsas int32 = C.GDK_KEY_containsas
const KEY_copyright int32 = C.GDK_KEY_copyright
const KEY_cr int32 = C.GDK_KEY_cr
const KEY_crossinglines int32 = C.GDK_KEY_crossinglines
const KEY_cuberoot int32 = C.GDK_KEY_cuberoot
const KEY_currency int32 = C.GDK_KEY_currency
const KEY_cursor int32 = C.GDK_KEY_cursor
const KEY_d int32 = C.GDK_KEY_d
const KEY_dabovedot int32 = C.GDK_KEY_dabovedot
const KEY_dagger int32 = C.GDK_KEY_dagger
const KEY_dcaron int32 = C.GDK_KEY_dcaron
const KEY_dead_A int32 = C.GDK_KEY_dead_A
const KEY_dead_E int32 = C.GDK_KEY_dead_E
const KEY_dead_I int32 = C.GDK_KEY_dead_I
const KEY_dead_O int32 = C.GDK_KEY_dead_O
const KEY_dead_U int32 = C.GDK_KEY_dead_U
const KEY_dead_a int32 = C.GDK_KEY_dead_a
const KEY_dead_abovecomma int32 = C.GDK_KEY_dead_abovecomma
const KEY_dead_abovedot int32 = C.GDK_KEY_dead_abovedot
const KEY_dead_abovereversedcomma int32 = C.GDK_KEY_dead_abovereversedcomma
const KEY_dead_abovering int32 = C.GDK_KEY_dead_abovering
const KEY_dead_acute int32 = C.GDK_KEY_dead_acute
const KEY_dead_belowbreve int32 = C.GDK_KEY_dead_belowbreve
const KEY_dead_belowcircumflex int32 = C.GDK_KEY_dead_belowcircumflex
const KEY_dead_belowcomma int32 = C.GDK_KEY_dead_belowcomma
const KEY_dead_belowdiaeresis int32 = C.GDK_KEY_dead_belowdiaeresis
const KEY_dead_belowdot int32 = C.GDK_KEY_dead_belowdot
const KEY_dead_belowmacron int32 = C.GDK_KEY_dead_belowmacron
const KEY_dead_belowring int32 = C.GDK_KEY_dead_belowring
const KEY_dead_belowtilde int32 = C.GDK_KEY_dead_belowtilde
const KEY_dead_breve int32 = C.GDK_KEY_dead_breve
const KEY_dead_capital_schwa int32 = C.GDK_KEY_dead_capital_schwa
const KEY_dead_caron int32 = C.GDK_KEY_dead_caron
const KEY_dead_cedilla int32 = C.GDK_KEY_dead_cedilla
const KEY_dead_circumflex int32 = C.GDK_KEY_dead_circumflex
const KEY_dead_currency int32 = C.GDK_KEY_dead_currency
const KEY_dead_dasia int32 = C.GDK_KEY_dead_dasia
const KEY_dead_diaeresis int32 = C.GDK_KEY_dead_diaeresis
const KEY_dead_doubleacute int32 = C.GDK_KEY_dead_doubleacute
const KEY_dead_doublegrave int32 = C.GDK_KEY_dead_doublegrave
const KEY_dead_e int32 = C.GDK_KEY_dead_e
const KEY_dead_grave int32 = C.GDK_KEY_dead_grave
const KEY_dead_greek int32 = C.GDK_KEY_dead_greek
const KEY_dead_hook int32 = C.GDK_KEY_dead_hook
const KEY_dead_horn int32 = C.GDK_KEY_dead_horn
const KEY_dead_i int32 = C.GDK_KEY_dead_i
const KEY_dead_invertedbreve int32 = C.GDK_KEY_dead_invertedbreve
const KEY_dead_iota int32 = C.GDK_KEY_dead_iota
const KEY_dead_macron int32 = C.GDK_KEY_dead_macron
const KEY_dead_o int32 = C.GDK_KEY_dead_o
const KEY_dead_ogonek int32 = C.GDK_KEY_dead_ogonek
const KEY_dead_perispomeni int32 = C.GDK_KEY_dead_perispomeni
const KEY_dead_psili int32 = C.GDK_KEY_dead_psili
const KEY_dead_semivoiced_sound int32 = C.GDK_KEY_dead_semivoiced_sound
const KEY_dead_small_schwa int32 = C.GDK_KEY_dead_small_schwa
const KEY_dead_stroke int32 = C.GDK_KEY_dead_stroke
const KEY_dead_tilde int32 = C.GDK_KEY_dead_tilde
const KEY_dead_u int32 = C.GDK_KEY_dead_u
const KEY_dead_voiced_sound int32 = C.GDK_KEY_dead_voiced_sound
const KEY_decimalpoint int32 = C.GDK_KEY_decimalpoint
const KEY_degree int32 = C.GDK_KEY_degree
const KEY_diaeresis int32 = C.GDK_KEY_diaeresis
const KEY_diamond int32 = C.GDK_KEY_diamond
const KEY_digitspace int32 = C.GDK_KEY_digitspace
const KEY_dintegral int32 = C.GDK_KEY_dintegral
const KEY_division int32 = C.GDK_KEY_division
const KEY_dollar int32 = C.GDK_KEY_dollar
const KEY_doubbaselinedot int32 = C.GDK_KEY_doubbaselinedot
const KEY_doubleacute int32 = C.GDK_KEY_doubleacute
const KEY_doubledagger int32 = C.GDK_KEY_doubledagger
const KEY_doublelowquotemark int32 = C.GDK_KEY_doublelowquotemark
const KEY_downarrow int32 = C.GDK_KEY_downarrow
const KEY_downcaret int32 = C.GDK_KEY_downcaret
const KEY_downshoe int32 = C.GDK_KEY_downshoe
const KEY_downstile int32 = C.GDK_KEY_downstile
const KEY_downtack int32 = C.GDK_KEY_downtack
const KEY_dstroke int32 = C.GDK_KEY_dstroke
const KEY_e int32 = C.GDK_KEY_e
const KEY_eabovedot int32 = C.GDK_KEY_eabovedot
const KEY_eacute int32 = C.GDK_KEY_eacute
const KEY_ebelowdot int32 = C.GDK_KEY_ebelowdot
const KEY_ecaron int32 = C.GDK_KEY_ecaron
const KEY_ecircumflex int32 = C.GDK_KEY_ecircumflex
const KEY_ecircumflexacute int32 = C.GDK_KEY_ecircumflexacute
const KEY_ecircumflexbelowdot int32 = C.GDK_KEY_ecircumflexbelowdot
const KEY_ecircumflexgrave int32 = C.GDK_KEY_ecircumflexgrave
const KEY_ecircumflexhook int32 = C.GDK_KEY_ecircumflexhook
const KEY_ecircumflextilde int32 = C.GDK_KEY_ecircumflextilde
const KEY_ediaeresis int32 = C.GDK_KEY_ediaeresis
const KEY_egrave int32 = C.GDK_KEY_egrave
const KEY_ehook int32 = C.GDK_KEY_ehook
const KEY_eightsubscript int32 = C.GDK_KEY_eightsubscript
const KEY_eightsuperior int32 = C.GDK_KEY_eightsuperior
const KEY_elementof int32 = C.GDK_KEY_elementof
const KEY_ellipsis int32 = C.GDK_KEY_ellipsis
const KEY_em3space int32 = C.GDK_KEY_em3space
const KEY_em4space int32 = C.GDK_KEY_em4space
const KEY_emacron int32 = C.GDK_KEY_emacron
const KEY_emdash int32 = C.GDK_KEY_emdash
const KEY_emfilledcircle int32 = C.GDK_KEY_emfilledcircle
const KEY_emfilledrect int32 = C.GDK_KEY_emfilledrect
const KEY_emopencircle int32 = C.GDK_KEY_emopencircle
const KEY_emopenrectangle int32 = C.GDK_KEY_emopenrectangle
const KEY_emptyset int32 = C.GDK_KEY_emptyset
const KEY_emspace int32 = C.GDK_KEY_emspace
const KEY_endash int32 = C.GDK_KEY_endash
const KEY_enfilledcircbullet int32 = C.GDK_KEY_enfilledcircbullet
const KEY_enfilledsqbullet int32 = C.GDK_KEY_enfilledsqbullet
const KEY_eng int32 = C.GDK_KEY_eng
const KEY_enopencircbullet int32 = C.GDK_KEY_enopencircbullet
const KEY_enopensquarebullet int32 = C.GDK_KEY_enopensquarebullet
const KEY_enspace int32 = C.GDK_KEY_enspace
const KEY_eogonek int32 = C.GDK_KEY_eogonek
const KEY_equal int32 = C.GDK_KEY_equal
const KEY_eth int32 = C.GDK_KEY_eth
const KEY_etilde int32 = C.GDK_KEY_etilde
const KEY_exclam int32 = C.GDK_KEY_exclam
const KEY_exclamdown int32 = C.GDK_KEY_exclamdown
const KEY_ezh int32 = C.GDK_KEY_ezh
const KEY_f int32 = C.GDK_KEY_f
const KEY_fabovedot int32 = C.GDK_KEY_fabovedot
const KEY_femalesymbol int32 = C.GDK_KEY_femalesymbol
const KEY_ff int32 = C.GDK_KEY_ff
const KEY_figdash int32 = C.GDK_KEY_figdash
const KEY_filledlefttribullet int32 = C.GDK_KEY_filledlefttribullet
const KEY_filledrectbullet int32 = C.GDK_KEY_filledrectbullet
const KEY_filledrighttribullet int32 = C.GDK_KEY_filledrighttribullet
const KEY_filledtribulletdown int32 = C.GDK_KEY_filledtribulletdown
const KEY_filledtribulletup int32 = C.GDK_KEY_filledtribulletup
const KEY_fiveeighths int32 = C.GDK_KEY_fiveeighths
const KEY_fivesixths int32 = C.GDK_KEY_fivesixths
const KEY_fivesubscript int32 = C.GDK_KEY_fivesubscript
const KEY_fivesuperior int32 = C.GDK_KEY_fivesuperior
const KEY_fourfifths int32 = C.GDK_KEY_fourfifths
const KEY_foursubscript int32 = C.GDK_KEY_foursubscript
const KEY_foursuperior int32 = C.GDK_KEY_foursuperior
const KEY_fourthroot int32 = C.GDK_KEY_fourthroot
const KEY_function int32 = C.GDK_KEY_function
const KEY_g int32 = C.GDK_KEY_g
const KEY_gabovedot int32 = C.GDK_KEY_gabovedot
const KEY_gbreve int32 = C.GDK_KEY_gbreve
const KEY_gcaron int32 = C.GDK_KEY_gcaron
const KEY_gcedilla int32 = C.GDK_KEY_gcedilla
const KEY_gcircumflex int32 = C.GDK_KEY_gcircumflex
const KEY_grave int32 = C.GDK_KEY_grave
const KEY_greater int32 = C.GDK_KEY_greater
const KEY_greaterthanequal int32 = C.GDK_KEY_greaterthanequal
const KEY_guillemotleft int32 = C.GDK_KEY_guillemotleft
const KEY_guillemotright int32 = C.GDK_KEY_guillemotright
const KEY_h int32 = C.GDK_KEY_h
const KEY_hairspace int32 = C.GDK_KEY_hairspace
const KEY_hcircumflex int32 = C.GDK_KEY_hcircumflex
const KEY_heart int32 = C.GDK_KEY_heart
const KEY_hebrew_aleph int32 = C.GDK_KEY_hebrew_aleph
const KEY_hebrew_ayin int32 = C.GDK_KEY_hebrew_ayin
const KEY_hebrew_bet int32 = C.GDK_KEY_hebrew_bet
const KEY_hebrew_beth int32 = C.GDK_KEY_hebrew_beth
const KEY_hebrew_chet int32 = C.GDK_KEY_hebrew_chet
const KEY_hebrew_dalet int32 = C.GDK_KEY_hebrew_dalet
const KEY_hebrew_daleth int32 = C.GDK_KEY_hebrew_daleth
const KEY_hebrew_doublelowline int32 = C.GDK_KEY_hebrew_doublelowline
const KEY_hebrew_finalkaph int32 = C.GDK_KEY_hebrew_finalkaph
const KEY_hebrew_finalmem int32 = C.GDK_KEY_hebrew_finalmem
const KEY_hebrew_finalnun int32 = C.GDK_KEY_hebrew_finalnun
const KEY_hebrew_finalpe int32 = C.GDK_KEY_hebrew_finalpe
const KEY_hebrew_finalzade int32 = C.GDK_KEY_hebrew_finalzade
const KEY_hebrew_finalzadi int32 = C.GDK_KEY_hebrew_finalzadi
const KEY_hebrew_gimel int32 = C.GDK_KEY_hebrew_gimel
const KEY_hebrew_gimmel int32 = C.GDK_KEY_hebrew_gimmel
const KEY_hebrew_he int32 = C.GDK_KEY_hebrew_he
const KEY_hebrew_het int32 = C.GDK_KEY_hebrew_het
const KEY_hebrew_kaph int32 = C.GDK_KEY_hebrew_kaph
const KEY_hebrew_kuf int32 = C.GDK_KEY_hebrew_kuf
const KEY_hebrew_lamed int32 = C.GDK_KEY_hebrew_lamed
const KEY_hebrew_mem int32 = C.GDK_KEY_hebrew_mem
const KEY_hebrew_nun int32 = C.GDK_KEY_hebrew_nun
const KEY_hebrew_pe int32 = C.GDK_KEY_hebrew_pe
const KEY_hebrew_qoph int32 = C.GDK_KEY_hebrew_qoph
const KEY_hebrew_resh int32 = C.GDK_KEY_hebrew_resh
const KEY_hebrew_samech int32 = C.GDK_KEY_hebrew_samech
const KEY_hebrew_samekh int32 = C.GDK_KEY_hebrew_samekh
const KEY_hebrew_shin int32 = C.GDK_KEY_hebrew_shin
const KEY_hebrew_taf int32 = C.GDK_KEY_hebrew_taf
const KEY_hebrew_taw int32 = C.GDK_KEY_hebrew_taw
const KEY_hebrew_tet int32 = C.GDK_KEY_hebrew_tet
const KEY_hebrew_teth int32 = C.GDK_KEY_hebrew_teth
const KEY_hebrew_waw int32 = C.GDK_KEY_hebrew_waw
const KEY_hebrew_yod int32 = C.GDK_KEY_hebrew_yod
const KEY_hebrew_zade int32 = C.GDK_KEY_hebrew_zade
const KEY_hebrew_zadi int32 = C.GDK_KEY_hebrew_zadi
const KEY_hebrew_zain int32 = C.GDK_KEY_hebrew_zain
const KEY_hebrew_zayin int32 = C.GDK_KEY_hebrew_zayin
const KEY_hexagram int32 = C.GDK_KEY_hexagram
const KEY_horizconnector int32 = C.GDK_KEY_horizconnector
const KEY_horizlinescan1 int32 = C.GDK_KEY_horizlinescan1
const KEY_horizlinescan3 int32 = C.GDK_KEY_horizlinescan3
const KEY_horizlinescan5 int32 = C.GDK_KEY_horizlinescan5
const KEY_horizlinescan7 int32 = C.GDK_KEY_horizlinescan7
const KEY_horizlinescan9 int32 = C.GDK_KEY_horizlinescan9
const KEY_hstroke int32 = C.GDK_KEY_hstroke
const KEY_ht int32 = C.GDK_KEY_ht
const KEY_hyphen int32 = C.GDK_KEY_hyphen
const KEY_i int32 = C.GDK_KEY_i
const KEY_iTouch int32 = C.GDK_KEY_iTouch
const KEY_iacute int32 = C.GDK_KEY_iacute
const KEY_ibelowdot int32 = C.GDK_KEY_ibelowdot
const KEY_ibreve int32 = C.GDK_KEY_ibreve
const KEY_icircumflex int32 = C.GDK_KEY_icircumflex
const KEY_identical int32 = C.GDK_KEY_identical
const KEY_idiaeresis int32 = C.GDK_KEY_idiaeresis
const KEY_idotless int32 = C.GDK_KEY_idotless
const KEY_ifonlyif int32 = C.GDK_KEY_ifonlyif
const KEY_igrave int32 = C.GDK_KEY_igrave
const KEY_ihook int32 = C.GDK_KEY_ihook
const KEY_imacron int32 = C.GDK_KEY_imacron
const KEY_implies int32 = C.GDK_KEY_implies
const KEY_includedin int32 = C.GDK_KEY_includedin
const KEY_includes int32 = C.GDK_KEY_includes
const KEY_infinity int32 = C.GDK_KEY_infinity
const KEY_integral int32 = C.GDK_KEY_integral
const KEY_intersection int32 = C.GDK_KEY_intersection
const KEY_iogonek int32 = C.GDK_KEY_iogonek
const KEY_itilde int32 = C.GDK_KEY_itilde
const KEY_j int32 = C.GDK_KEY_j
const KEY_jcircumflex int32 = C.GDK_KEY_jcircumflex
const KEY_jot int32 = C.GDK_KEY_jot
const KEY_k int32 = C.GDK_KEY_k
const KEY_kana_A int32 = C.GDK_KEY_kana_A
const KEY_kana_CHI int32 = C.GDK_KEY_kana_CHI
const KEY_kana_E int32 = C.GDK_KEY_kana_E
const KEY_kana_FU int32 = C.GDK_KEY_kana_FU
const KEY_kana_HA int32 = C.GDK_KEY_kana_HA
const KEY_kana_HE int32 = C.GDK_KEY_kana_HE
const KEY_kana_HI int32 = C.GDK_KEY_kana_HI
const KEY_kana_HO int32 = C.GDK_KEY_kana_HO
const KEY_kana_HU int32 = C.GDK_KEY_kana_HU
const KEY_kana_I int32 = C.GDK_KEY_kana_I
const KEY_kana_KA int32 = C.GDK_KEY_kana_KA
const KEY_kana_KE int32 = C.GDK_KEY_kana_KE
const KEY_kana_KI int32 = C.GDK_KEY_kana_KI
const KEY_kana_KO int32 = C.GDK_KEY_kana_KO
const KEY_kana_KU int32 = C.GDK_KEY_kana_KU
const KEY_kana_MA int32 = C.GDK_KEY_kana_MA
const KEY_kana_ME int32 = C.GDK_KEY_kana_ME
const KEY_kana_MI int32 = C.GDK_KEY_kana_MI
const KEY_kana_MO int32 = C.GDK_KEY_kana_MO
const KEY_kana_MU int32 = C.GDK_KEY_kana_MU
const KEY_kana_N int32 = C.GDK_KEY_kana_N
const KEY_kana_NA int32 = C.GDK_KEY_kana_NA
const KEY_kana_NE int32 = C.GDK_KEY_kana_NE
const KEY_kana_NI int32 = C.GDK_KEY_kana_NI
const KEY_kana_NO int32 = C.GDK_KEY_kana_NO
const KEY_kana_NU int32 = C.GDK_KEY_kana_NU
const KEY_kana_O int32 = C.GDK_KEY_kana_O
const KEY_kana_RA int32 = C.GDK_KEY_kana_RA
const KEY_kana_RE int32 = C.GDK_KEY_kana_RE
const KEY_kana_RI int32 = C.GDK_KEY_kana_RI
const KEY_kana_RO int32 = C.GDK_KEY_kana_RO
const KEY_kana_RU int32 = C.GDK_KEY_kana_RU
const KEY_kana_SA int32 = C.GDK_KEY_kana_SA
const KEY_kana_SE int32 = C.GDK_KEY_kana_SE
const KEY_kana_SHI int32 = C.GDK_KEY_kana_SHI
const KEY_kana_SO int32 = C.GDK_KEY_kana_SO
const KEY_kana_SU int32 = C.GDK_KEY_kana_SU
const KEY_kana_TA int32 = C.GDK_KEY_kana_TA
const KEY_kana_TE int32 = C.GDK_KEY_kana_TE
const KEY_kana_TI int32 = C.GDK_KEY_kana_TI
const KEY_kana_TO int32 = C.GDK_KEY_kana_TO
const KEY_kana_TSU int32 = C.GDK_KEY_kana_TSU
const KEY_kana_TU int32 = C.GDK_KEY_kana_TU
const KEY_kana_U int32 = C.GDK_KEY_kana_U
const KEY_kana_WA int32 = C.GDK_KEY_kana_WA
const KEY_kana_WO int32 = C.GDK_KEY_kana_WO
const KEY_kana_YA int32 = C.GDK_KEY_kana_YA
const KEY_kana_YO int32 = C.GDK_KEY_kana_YO
const KEY_kana_YU int32 = C.GDK_KEY_kana_YU
const KEY_kana_a int32 = C.GDK_KEY_kana_a
const KEY_kana_closingbracket int32 = C.GDK_KEY_kana_closingbracket
const KEY_kana_comma int32 = C.GDK_KEY_kana_comma
const KEY_kana_conjunctive int32 = C.GDK_KEY_kana_conjunctive
const KEY_kana_e int32 = C.GDK_KEY_kana_e
const KEY_kana_fullstop int32 = C.GDK_KEY_kana_fullstop
const KEY_kana_i int32 = C.GDK_KEY_kana_i
const KEY_kana_middledot int32 = C.GDK_KEY_kana_middledot
const KEY_kana_o int32 = C.GDK_KEY_kana_o
const KEY_kana_openingbracket int32 = C.GDK_KEY_kana_openingbracket
const KEY_kana_switch int32 = C.GDK_KEY_kana_switch
const KEY_kana_tsu int32 = C.GDK_KEY_kana_tsu
const KEY_kana_tu int32 = C.GDK_KEY_kana_tu
const KEY_kana_u int32 = C.GDK_KEY_kana_u
const KEY_kana_ya int32 = C.GDK_KEY_kana_ya
const KEY_kana_yo int32 = C.GDK_KEY_kana_yo
const KEY_kana_yu int32 = C.GDK_KEY_kana_yu
const KEY_kappa int32 = C.GDK_KEY_kappa
const KEY_kcedilla int32 = C.GDK_KEY_kcedilla
const KEY_kra int32 = C.GDK_KEY_kra
const KEY_l int32 = C.GDK_KEY_l
const KEY_lacute int32 = C.GDK_KEY_lacute
const KEY_latincross int32 = C.GDK_KEY_latincross
const KEY_lbelowdot int32 = C.GDK_KEY_lbelowdot
const KEY_lcaron int32 = C.GDK_KEY_lcaron
const KEY_lcedilla int32 = C.GDK_KEY_lcedilla
const KEY_leftanglebracket int32 = C.GDK_KEY_leftanglebracket
const KEY_leftarrow int32 = C.GDK_KEY_leftarrow
const KEY_leftcaret int32 = C.GDK_KEY_leftcaret
const KEY_leftdoublequotemark int32 = C.GDK_KEY_leftdoublequotemark
const KEY_leftmiddlecurlybrace int32 = C.GDK_KEY_leftmiddlecurlybrace
const KEY_leftopentriangle int32 = C.GDK_KEY_leftopentriangle
const KEY_leftpointer int32 = C.GDK_KEY_leftpointer
const KEY_leftradical int32 = C.GDK_KEY_leftradical
const KEY_leftshoe int32 = C.GDK_KEY_leftshoe
const KEY_leftsinglequotemark int32 = C.GDK_KEY_leftsinglequotemark
const KEY_leftt int32 = C.GDK_KEY_leftt
const KEY_lefttack int32 = C.GDK_KEY_lefttack
const KEY_less int32 = C.GDK_KEY_less
const KEY_lessthanequal int32 = C.GDK_KEY_lessthanequal
const KEY_lf int32 = C.GDK_KEY_lf
const KEY_logicaland int32 = C.GDK_KEY_logicaland
const KEY_logicalor int32 = C.GDK_KEY_logicalor
const KEY_lowleftcorner int32 = C.GDK_KEY_lowleftcorner
const KEY_lowrightcorner int32 = C.GDK_KEY_lowrightcorner
const KEY_lstroke int32 = C.GDK_KEY_lstroke
const KEY_m int32 = C.GDK_KEY_m
const KEY_mabovedot int32 = C.GDK_KEY_mabovedot
const KEY_macron int32 = C.GDK_KEY_macron
const KEY_malesymbol int32 = C.GDK_KEY_malesymbol
const KEY_maltesecross int32 = C.GDK_KEY_maltesecross
const KEY_marker int32 = C.GDK_KEY_marker
const KEY_masculine int32 = C.GDK_KEY_masculine
const KEY_minus int32 = C.GDK_KEY_minus
const KEY_minutes int32 = C.GDK_KEY_minutes
const KEY_mu int32 = C.GDK_KEY_mu
const KEY_multiply int32 = C.GDK_KEY_multiply
const KEY_musicalflat int32 = C.GDK_KEY_musicalflat
const KEY_musicalsharp int32 = C.GDK_KEY_musicalsharp
const KEY_n int32 = C.GDK_KEY_n
const KEY_nabla int32 = C.GDK_KEY_nabla
const KEY_nacute int32 = C.GDK_KEY_nacute
const KEY_ncaron int32 = C.GDK_KEY_ncaron
const KEY_ncedilla int32 = C.GDK_KEY_ncedilla
const KEY_ninesubscript int32 = C.GDK_KEY_ninesubscript
const KEY_ninesuperior int32 = C.GDK_KEY_ninesuperior
const KEY_nl int32 = C.GDK_KEY_nl
const KEY_nobreakspace int32 = C.GDK_KEY_nobreakspace
const KEY_notapproxeq int32 = C.GDK_KEY_notapproxeq
const KEY_notelementof int32 = C.GDK_KEY_notelementof
const KEY_notequal int32 = C.GDK_KEY_notequal
const KEY_notidentical int32 = C.GDK_KEY_notidentical
const KEY_notsign int32 = C.GDK_KEY_notsign
const KEY_ntilde int32 = C.GDK_KEY_ntilde
const KEY_numbersign int32 = C.GDK_KEY_numbersign
const KEY_numerosign int32 = C.GDK_KEY_numerosign
const KEY_o int32 = C.GDK_KEY_o
const KEY_oacute int32 = C.GDK_KEY_oacute
const KEY_obarred int32 = C.GDK_KEY_obarred
const KEY_obelowdot int32 = C.GDK_KEY_obelowdot
const KEY_ocaron int32 = C.GDK_KEY_ocaron
const KEY_ocircumflex int32 = C.GDK_KEY_ocircumflex
const KEY_ocircumflexacute int32 = C.GDK_KEY_ocircumflexacute
const KEY_ocircumflexbelowdot int32 = C.GDK_KEY_ocircumflexbelowdot
const KEY_ocircumflexgrave int32 = C.GDK_KEY_ocircumflexgrave
const KEY_ocircumflexhook int32 = C.GDK_KEY_ocircumflexhook
const KEY_ocircumflextilde int32 = C.GDK_KEY_ocircumflextilde
const KEY_odiaeresis int32 = C.GDK_KEY_odiaeresis
const KEY_odoubleacute int32 = C.GDK_KEY_odoubleacute
const KEY_oe int32 = C.GDK_KEY_oe
const KEY_ogonek int32 = C.GDK_KEY_ogonek
const KEY_ograve int32 = C.GDK_KEY_ograve
const KEY_ohook int32 = C.GDK_KEY_ohook
const KEY_ohorn int32 = C.GDK_KEY_ohorn
const KEY_ohornacute int32 = C.GDK_KEY_ohornacute
const KEY_ohornbelowdot int32 = C.GDK_KEY_ohornbelowdot
const KEY_ohorngrave int32 = C.GDK_KEY_ohorngrave
const KEY_ohornhook int32 = C.GDK_KEY_ohornhook
const KEY_ohorntilde int32 = C.GDK_KEY_ohorntilde
const KEY_omacron int32 = C.GDK_KEY_omacron
const KEY_oneeighth int32 = C.GDK_KEY_oneeighth
const KEY_onefifth int32 = C.GDK_KEY_onefifth
const KEY_onehalf int32 = C.GDK_KEY_onehalf
const KEY_onequarter int32 = C.GDK_KEY_onequarter
const KEY_onesixth int32 = C.GDK_KEY_onesixth
const KEY_onesubscript int32 = C.GDK_KEY_onesubscript
const KEY_onesuperior int32 = C.GDK_KEY_onesuperior
const KEY_onethird int32 = C.GDK_KEY_onethird
const KEY_ooblique int32 = C.GDK_KEY_ooblique
const KEY_openrectbullet int32 = C.GDK_KEY_openrectbullet
const KEY_openstar int32 = C.GDK_KEY_openstar
const KEY_opentribulletdown int32 = C.GDK_KEY_opentribulletdown
const KEY_opentribulletup int32 = C.GDK_KEY_opentribulletup
const KEY_ordfeminine int32 = C.GDK_KEY_ordfeminine
const KEY_oslash int32 = C.GDK_KEY_oslash
const KEY_otilde int32 = C.GDK_KEY_otilde
const KEY_overbar int32 = C.GDK_KEY_overbar
const KEY_overline int32 = C.GDK_KEY_overline
const KEY_p int32 = C.GDK_KEY_p
const KEY_pabovedot int32 = C.GDK_KEY_pabovedot
const KEY_paragraph int32 = C.GDK_KEY_paragraph
const KEY_parenleft int32 = C.GDK_KEY_parenleft
const KEY_parenright int32 = C.GDK_KEY_parenright
const KEY_partdifferential int32 = C.GDK_KEY_partdifferential
const KEY_partialderivative int32 = C.GDK_KEY_partialderivative
const KEY_percent int32 = C.GDK_KEY_percent
const KEY_period int32 = C.GDK_KEY_period
const KEY_periodcentered int32 = C.GDK_KEY_periodcentered
const KEY_permille int32 = C.GDK_KEY_permille
const KEY_phonographcopyright int32 = C.GDK_KEY_phonographcopyright
const KEY_plus int32 = C.GDK_KEY_plus
const KEY_plusminus int32 = C.GDK_KEY_plusminus
const KEY_prescription int32 = C.GDK_KEY_prescription
const KEY_prolongedsound int32 = C.GDK_KEY_prolongedsound
const KEY_punctspace int32 = C.GDK_KEY_punctspace
const KEY_q int32 = C.GDK_KEY_q
const KEY_quad int32 = C.GDK_KEY_quad
const KEY_question int32 = C.GDK_KEY_question
const KEY_questiondown int32 = C.GDK_KEY_questiondown
const KEY_quotedbl int32 = C.GDK_KEY_quotedbl
const KEY_quoteleft int32 = C.GDK_KEY_quoteleft
const KEY_quoteright int32 = C.GDK_KEY_quoteright
const KEY_r int32 = C.GDK_KEY_r
const KEY_racute int32 = C.GDK_KEY_racute
const KEY_radical int32 = C.GDK_KEY_radical
const KEY_rcaron int32 = C.GDK_KEY_rcaron
const KEY_rcedilla int32 = C.GDK_KEY_rcedilla
const KEY_registered int32 = C.GDK_KEY_registered
const KEY_rightanglebracket int32 = C.GDK_KEY_rightanglebracket
const KEY_rightarrow int32 = C.GDK_KEY_rightarrow
const KEY_rightcaret int32 = C.GDK_KEY_rightcaret
const KEY_rightdoublequotemark int32 = C.GDK_KEY_rightdoublequotemark
const KEY_rightmiddlecurlybrace int32 = C.GDK_KEY_rightmiddlecurlybrace
const KEY_rightmiddlesummation int32 = C.GDK_KEY_rightmiddlesummation
const KEY_rightopentriangle int32 = C.GDK_KEY_rightopentriangle
const KEY_rightpointer int32 = C.GDK_KEY_rightpointer
const KEY_rightshoe int32 = C.GDK_KEY_rightshoe
const KEY_rightsinglequotemark int32 = C.GDK_KEY_rightsinglequotemark
const KEY_rightt int32 = C.GDK_KEY_rightt
const KEY_righttack int32 = C.GDK_KEY_righttack
const KEY_s int32 = C.GDK_KEY_s
const KEY_sabovedot int32 = C.GDK_KEY_sabovedot
const KEY_sacute int32 = C.GDK_KEY_sacute
const KEY_scaron int32 = C.GDK_KEY_scaron
const KEY_scedilla int32 = C.GDK_KEY_scedilla
const KEY_schwa int32 = C.GDK_KEY_schwa
const KEY_scircumflex int32 = C.GDK_KEY_scircumflex
const KEY_script_switch int32 = C.GDK_KEY_script_switch
const KEY_seconds int32 = C.GDK_KEY_seconds
const KEY_section int32 = C.GDK_KEY_section
const KEY_semicolon int32 = C.GDK_KEY_semicolon
const KEY_semivoicedsound int32 = C.GDK_KEY_semivoicedsound
const KEY_seveneighths int32 = C.GDK_KEY_seveneighths
const KEY_sevensubscript int32 = C.GDK_KEY_sevensubscript
const KEY_sevensuperior int32 = C.GDK_KEY_sevensuperior
const KEY_signaturemark int32 = C.GDK_KEY_signaturemark
const KEY_signifblank int32 = C.GDK_KEY_signifblank
const KEY_similarequal int32 = C.GDK_KEY_similarequal
const KEY_singlelowquotemark int32 = C.GDK_KEY_singlelowquotemark
const KEY_sixsubscript int32 = C.GDK_KEY_sixsubscript
const KEY_sixsuperior int32 = C.GDK_KEY_sixsuperior
const KEY_slash int32 = C.GDK_KEY_slash
const KEY_soliddiamond int32 = C.GDK_KEY_soliddiamond
const KEY_space int32 = C.GDK_KEY_space
const KEY_squareroot int32 = C.GDK_KEY_squareroot
const KEY_ssharp int32 = C.GDK_KEY_ssharp
const KEY_sterling int32 = C.GDK_KEY_sterling
const KEY_stricteq int32 = C.GDK_KEY_stricteq
const KEY_t int32 = C.GDK_KEY_t
const KEY_tabovedot int32 = C.GDK_KEY_tabovedot
const KEY_tcaron int32 = C.GDK_KEY_tcaron
const KEY_tcedilla int32 = C.GDK_KEY_tcedilla
const KEY_telephone int32 = C.GDK_KEY_telephone
const KEY_telephonerecorder int32 = C.GDK_KEY_telephonerecorder
const KEY_therefore int32 = C.GDK_KEY_therefore
const KEY_thinspace int32 = C.GDK_KEY_thinspace
const KEY_thorn int32 = C.GDK_KEY_thorn
const KEY_threeeighths int32 = C.GDK_KEY_threeeighths
const KEY_threefifths int32 = C.GDK_KEY_threefifths
const KEY_threequarters int32 = C.GDK_KEY_threequarters
const KEY_threesubscript int32 = C.GDK_KEY_threesubscript
const KEY_threesuperior int32 = C.GDK_KEY_threesuperior
const KEY_tintegral int32 = C.GDK_KEY_tintegral
const KEY_topintegral int32 = C.GDK_KEY_topintegral
const KEY_topleftparens int32 = C.GDK_KEY_topleftparens
const KEY_topleftradical int32 = C.GDK_KEY_topleftradical
const KEY_topleftsqbracket int32 = C.GDK_KEY_topleftsqbracket
const KEY_topleftsummation int32 = C.GDK_KEY_topleftsummation
const KEY_toprightparens int32 = C.GDK_KEY_toprightparens
const KEY_toprightsqbracket int32 = C.GDK_KEY_toprightsqbracket
const KEY_toprightsummation int32 = C.GDK_KEY_toprightsummation
const KEY_topt int32 = C.GDK_KEY_topt
const KEY_topvertsummationconnector int32 = C.GDK_KEY_topvertsummationconnector
const KEY_trademark int32 = C.GDK_KEY_trademark
const KEY_trademarkincircle int32 = C.GDK_KEY_trademarkincircle
const KEY_tslash int32 = C.GDK_KEY_tslash
const KEY_twofifths int32 = C.GDK_KEY_twofifths
const KEY_twosubscript int32 = C.GDK_KEY_twosubscript
const KEY_twosuperior int32 = C.GDK_KEY_twosuperior
const KEY_twothirds int32 = C.GDK_KEY_twothirds
const KEY_u int32 = C.GDK_KEY_u
const KEY_uacute int32 = C.GDK_KEY_uacute
const KEY_ubelowdot int32 = C.GDK_KEY_ubelowdot
const KEY_ubreve int32 = C.GDK_KEY_ubreve
const KEY_ucircumflex int32 = C.GDK_KEY_ucircumflex
const KEY_udiaeresis int32 = C.GDK_KEY_udiaeresis
const KEY_udoubleacute int32 = C.GDK_KEY_udoubleacute
const KEY_ugrave int32 = C.GDK_KEY_ugrave
const KEY_uhook int32 = C.GDK_KEY_uhook
const KEY_uhorn int32 = C.GDK_KEY_uhorn
const KEY_uhornacute int32 = C.GDK_KEY_uhornacute
const KEY_uhornbelowdot int32 = C.GDK_KEY_uhornbelowdot
const KEY_uhorngrave int32 = C.GDK_KEY_uhorngrave
const KEY_uhornhook int32 = C.GDK_KEY_uhornhook
const KEY_uhorntilde int32 = C.GDK_KEY_uhorntilde
const KEY_umacron int32 = C.GDK_KEY_umacron
const KEY_underbar int32 = C.GDK_KEY_underbar
const KEY_underscore int32 = C.GDK_KEY_underscore
const KEY_union int32 = C.GDK_KEY_union
const KEY_uogonek int32 = C.GDK_KEY_uogonek
const KEY_uparrow int32 = C.GDK_KEY_uparrow
const KEY_upcaret int32 = C.GDK_KEY_upcaret
const KEY_upleftcorner int32 = C.GDK_KEY_upleftcorner
const KEY_uprightcorner int32 = C.GDK_KEY_uprightcorner
const KEY_upshoe int32 = C.GDK_KEY_upshoe
const KEY_upstile int32 = C.GDK_KEY_upstile
const KEY_uptack int32 = C.GDK_KEY_uptack
const KEY_uring int32 = C.GDK_KEY_uring
const KEY_utilde int32 = C.GDK_KEY_utilde
const KEY_v int32 = C.GDK_KEY_v
const KEY_variation int32 = C.GDK_KEY_variation
const KEY_vertbar int32 = C.GDK_KEY_vertbar
const KEY_vertconnector int32 = C.GDK_KEY_vertconnector
const KEY_voicedsound int32 = C.GDK_KEY_voicedsound
const KEY_vt int32 = C.GDK_KEY_vt
const KEY_w int32 = C.GDK_KEY_w
const KEY_wacute int32 = C.GDK_KEY_wacute
const KEY_wcircumflex int32 = C.GDK_KEY_wcircumflex
const KEY_wdiaeresis int32 = C.GDK_KEY_wdiaeresis
const KEY_wgrave int32 = C.GDK_KEY_wgrave
const KEY_x int32 = C.GDK_KEY_x
const KEY_xabovedot int32 = C.GDK_KEY_xabovedot
const KEY_y int32 = C.GDK_KEY_y
const KEY_yacute int32 = C.GDK_KEY_yacute
const KEY_ybelowdot int32 = C.GDK_KEY_ybelowdot
const KEY_ycircumflex int32 = C.GDK_KEY_ycircumflex
const KEY_ydiaeresis int32 = C.GDK_KEY_ydiaeresis
const KEY_yen int32 = C.GDK_KEY_yen
const KEY_ygrave int32 = C.GDK_KEY_ygrave
const KEY_yhook int32 = C.GDK_KEY_yhook
const KEY_ytilde int32 = C.GDK_KEY_ytilde
const KEY_z int32 = C.GDK_KEY_z
const KEY_zabovedot int32 = C.GDK_KEY_zabovedot
const KEY_zacute int32 = C.GDK_KEY_zacute
const KEY_zcaron int32 = C.GDK_KEY_zcaron
const KEY_zerosubscript int32 = C.GDK_KEY_zerosubscript
const KEY_zerosuperior int32 = C.GDK_KEY_zerosuperior
const KEY_zstroke int32 = C.GDK_KEY_zstroke
const MAX_TIMECOORD_AXES int32 = C.GDK_MAX_TIMECOORD_AXES
const PARENT_RELATIVE int32 = C.GDK_PARENT_RELATIVE
const PRIORITY_REDRAW int32 = C.GDK_PRIORITY_REDRAW

type AxisUse C.GdkAxisUse

const (
	GDK_AXIS_IGNORE   AxisUse = 0
	GDK_AXIS_X        AxisUse = 1
	GDK_AXIS_Y        AxisUse = 2
	GDK_AXIS_PRESSURE AxisUse = 3
	GDK_AXIS_XTILT    AxisUse = 4
	GDK_AXIS_YTILT    AxisUse = 5
	GDK_AXIS_WHEEL    AxisUse = 6
	GDK_AXIS_DISTANCE AxisUse = 7
	GDK_AXIS_ROTATION AxisUse = 8
	GDK_AXIS_SLIDER   AxisUse = 9
	GDK_AXIS_LAST     AxisUse = 10
)

type ByteOrder C.GdkByteOrder

const (
	GDK_LSB_FIRST ByteOrder = 0
	GDK_MSB_FIRST ByteOrder = 1
)

type CrossingMode C.GdkCrossingMode

const (
	GDK_CROSSING_NORMAL        CrossingMode = 0
	GDK_CROSSING_GRAB          CrossingMode = 1
	GDK_CROSSING_UNGRAB        CrossingMode = 2
	GDK_CROSSING_GTK_GRAB      CrossingMode = 3
	GDK_CROSSING_GTK_UNGRAB    CrossingMode = 4
	GDK_CROSSING_STATE_CHANGED CrossingMode = 5
	GDK_CROSSING_TOUCH_BEGIN   CrossingMode = 6
	GDK_CROSSING_TOUCH_END     CrossingMode = 7
	GDK_CROSSING_DEVICE_SWITCH CrossingMode = 8
)

type CursorType C.GdkCursorType

const (
	GDK_X_CURSOR            CursorType = 0
	GDK_ARROW               CursorType = 2
	GDK_BASED_ARROW_DOWN    CursorType = 4
	GDK_BASED_ARROW_UP      CursorType = 6
	GDK_BOAT                CursorType = 8
	GDK_BOGOSITY            CursorType = 10
	GDK_BOTTOM_LEFT_CORNER  CursorType = 12
	GDK_BOTTOM_RIGHT_CORNER CursorType = 14
	GDK_BOTTOM_SIDE         CursorType = 16
	GDK_BOTTOM_TEE          CursorType = 18
	GDK_BOX_SPIRAL          CursorType = 20
	GDK_CENTER_PTR          CursorType = 22
	GDK_CIRCLE              CursorType = 24
	GDK_CLOCK               CursorType = 26
	GDK_COFFEE_MUG          CursorType = 28
	GDK_CROSS               CursorType = 30
	GDK_CROSS_REVERSE       CursorType = 32
	GDK_CROSSHAIR           CursorType = 34
	GDK_DIAMOND_CROSS       CursorType = 36
	GDK_DOT                 CursorType = 38
	GDK_DOTBOX              CursorType = 40
	GDK_DOUBLE_ARROW        CursorType = 42
	GDK_DRAFT_LARGE         CursorType = 44
	GDK_DRAFT_SMALL         CursorType = 46
	GDK_DRAPED_BOX          CursorType = 48
	GDK_EXCHANGE            CursorType = 50
	GDK_FLEUR               CursorType = 52
	GDK_GOBBLER             CursorType = 54
	GDK_GUMBY               CursorType = 56
	GDK_HAND1               CursorType = 58
	GDK_HAND2               CursorType = 60
	GDK_HEART               CursorType = 62
	GDK_ICON                CursorType = 64
	GDK_IRON_CROSS          CursorType = 66
	GDK_LEFT_PTR            CursorType = 68
	GDK_LEFT_SIDE           CursorType = 70
	GDK_LEFT_TEE            CursorType = 72
	GDK_LEFTBUTTON          CursorType = 74
	GDK_LL_ANGLE            CursorType = 76
	GDK_LR_ANGLE            CursorType = 78
	GDK_MAN                 CursorType = 80
	GDK_MIDDLEBUTTON        CursorType = 82
	GDK_MOUSE               CursorType = 84
	GDK_PENCIL              CursorType = 86
	GDK_PIRATE              CursorType = 88
	GDK_PLUS                CursorType = 90
	GDK_QUESTION_ARROW      CursorType = 92
	GDK_RIGHT_PTR           CursorType = 94
	GDK_RIGHT_SIDE          CursorType = 96
	GDK_RIGHT_TEE           CursorType = 98
	GDK_RIGHTBUTTON         CursorType = 100
	GDK_RTL_LOGO            CursorType = 102
	GDK_SAILBOAT            CursorType = 104
	GDK_SB_DOWN_ARROW       CursorType = 106
	GDK_SB_H_DOUBLE_ARROW   CursorType = 108
	GDK_SB_LEFT_ARROW       CursorType = 110
	GDK_SB_RIGHT_ARROW      CursorType = 112
	GDK_SB_UP_ARROW         CursorType = 114
	GDK_SB_V_DOUBLE_ARROW   CursorType = 116
	GDK_SHUTTLE             CursorType = 118
	GDK_SIZING              CursorType = 120
	GDK_SPIDER              CursorType = 122
	GDK_SPRAYCAN            CursorType = 124
	GDK_STAR                CursorType = 126
	GDK_TARGET              CursorType = 128
	GDK_TCROSS              CursorType = 130
	GDK_TOP_LEFT_ARROW      CursorType = 132
	GDK_TOP_LEFT_CORNER     CursorType = 134
	GDK_TOP_RIGHT_CORNER    CursorType = 136
	GDK_TOP_SIDE            CursorType = 138
	GDK_TOP_TEE             CursorType = 140
	GDK_TREK                CursorType = 142
	GDK_UL_ANGLE            CursorType = 144
	GDK_UMBRELLA            CursorType = 146
	GDK_UR_ANGLE            CursorType = 148
	GDK_WATCH               CursorType = 150
	GDK_XTERM               CursorType = 152
	GDK_LAST_CURSOR         CursorType = 153
	GDK_BLANK_CURSOR        CursorType = -2
	GDK_CURSOR_IS_PIXMAP    CursorType = -1
)

type DeviceType C.GdkDeviceType

const (
	GDK_DEVICE_TYPE_MASTER   DeviceType = 0
	GDK_DEVICE_TYPE_SLAVE    DeviceType = 1
	GDK_DEVICE_TYPE_FLOATING DeviceType = 2
)

type DragProtocol C.GdkDragProtocol

const (
	GDK_DRAG_PROTO_NONE            DragProtocol = 0
	GDK_DRAG_PROTO_MOTIF           DragProtocol = 1
	GDK_DRAG_PROTO_XDND            DragProtocol = 2
	GDK_DRAG_PROTO_ROOTWIN         DragProtocol = 3
	GDK_DRAG_PROTO_WIN32_DROPFILES DragProtocol = 4
	GDK_DRAG_PROTO_OLE2            DragProtocol = 5
	GDK_DRAG_PROTO_LOCAL           DragProtocol = 6
	GDK_DRAG_PROTO_WAYLAND         DragProtocol = 7
)

type EventType C.GdkEventType

const (
	GDK_NOTHING             EventType = -1
	GDK_DELETE              EventType = 0
	GDK_DESTROY             EventType = 1
	GDK_EXPOSE              EventType = 2
	GDK_MOTION_NOTIFY       EventType = 3
	GDK_BUTTON_PRESS        EventType = 4
	GDK_2BUTTON_PRESS       EventType = 5
	GDK_DOUBLE_BUTTON_PRESS EventType = 5
	GDK_3BUTTON_PRESS       EventType = 6
	GDK_TRIPLE_BUTTON_PRESS EventType = 6
	GDK_BUTTON_RELEASE      EventType = 7
	GDK_KEY_PRESS           EventType = 8
	GDK_KEY_RELEASE         EventType = 9
	GDK_ENTER_NOTIFY        EventType = 10
	GDK_LEAVE_NOTIFY        EventType = 11
	GDK_FOCUS_CHANGE        EventType = 12
	GDK_CONFIGURE           EventType = 13
	GDK_MAP                 EventType = 14
	GDK_UNMAP               EventType = 15
	GDK_PROPERTY_NOTIFY     EventType = 16
	GDK_SELECTION_CLEAR     EventType = 17
	GDK_SELECTION_REQUEST   EventType = 18
	GDK_SELECTION_NOTIFY    EventType = 19
	GDK_PROXIMITY_IN        EventType = 20
	GDK_PROXIMITY_OUT       EventType = 21
	GDK_DRAG_ENTER          EventType = 22
	GDK_DRAG_LEAVE          EventType = 23
	GDK_DRAG_MOTION         EventType = 24
	GDK_DRAG_STATUS         EventType = 25
	GDK_DROP_START          EventType = 26
	GDK_DROP_FINISHED       EventType = 27
	GDK_CLIENT_EVENT        EventType = 28
	GDK_VISIBILITY_NOTIFY   EventType = 29
	GDK_SCROLL              EventType = 31
	GDK_WINDOW_STATE        EventType = 32
	GDK_SETTING             EventType = 33
	GDK_OWNER_CHANGE        EventType = 34
	GDK_GRAB_BROKEN         EventType = 35
	GDK_DAMAGE              EventType = 36
	GDK_TOUCH_BEGIN         EventType = 37
	GDK_TOUCH_UPDATE        EventType = 38
	GDK_TOUCH_END           EventType = 39
	GDK_TOUCH_CANCEL        EventType = 40
	GDK_TOUCHPAD_SWIPE      EventType = 41
	GDK_TOUCHPAD_PINCH      EventType = 42
	GDK_PAD_BUTTON_PRESS    EventType = 43
	GDK_PAD_BUTTON_RELEASE  EventType = 44
	GDK_PAD_RING            EventType = 45
	GDK_PAD_STRIP           EventType = 46
	GDK_PAD_GROUP_MODE      EventType = 47
	GDK_EVENT_LAST          EventType = 48
)

type FilterReturn C.GdkFilterReturn

const (
	GDK_FILTER_CONTINUE  FilterReturn = 0
	GDK_FILTER_TRANSLATE FilterReturn = 1
	GDK_FILTER_REMOVE    FilterReturn = 2
)

type GrabOwnership C.GdkGrabOwnership

const (
	GDK_OWNERSHIP_NONE        GrabOwnership = 0
	GDK_OWNERSHIP_WINDOW      GrabOwnership = 1
	GDK_OWNERSHIP_APPLICATION GrabOwnership = 2
)

type GrabStatus C.GdkGrabStatus

const (
	GDK_GRAB_SUCCESS         GrabStatus = 0
	GDK_GRAB_ALREADY_GRABBED GrabStatus = 1
	GDK_GRAB_INVALID_TIME    GrabStatus = 2
	GDK_GRAB_NOT_VIEWABLE    GrabStatus = 3
	GDK_GRAB_FROZEN          GrabStatus = 4
	GDK_GRAB_FAILED          GrabStatus = 5
)

type Gravity C.GdkGravity

const (
	GDK_GRAVITY_NORTH_WEST Gravity = 1
	GDK_GRAVITY_NORTH      Gravity = 2
	GDK_GRAVITY_NORTH_EAST Gravity = 3
	GDK_GRAVITY_WEST       Gravity = 4
	GDK_GRAVITY_CENTER     Gravity = 5
	GDK_GRAVITY_EAST       Gravity = 6
	GDK_GRAVITY_SOUTH_WEST Gravity = 7
	GDK_GRAVITY_SOUTH      Gravity = 8
	GDK_GRAVITY_SOUTH_EAST Gravity = 9
	GDK_GRAVITY_STATIC     Gravity = 10
)

type InputMode C.GdkInputMode

const (
	GDK_MODE_DISABLED InputMode = 0
	GDK_MODE_SCREEN   InputMode = 1
	GDK_MODE_WINDOW   InputMode = 2
)

type InputSource C.GdkInputSource

const (
	GDK_SOURCE_MOUSE       InputSource = 0
	GDK_SOURCE_PEN         InputSource = 1
	GDK_SOURCE_ERASER      InputSource = 2
	GDK_SOURCE_CURSOR      InputSource = 3
	GDK_SOURCE_KEYBOARD    InputSource = 4
	GDK_SOURCE_TOUCHSCREEN InputSource = 5
	GDK_SOURCE_TOUCHPAD    InputSource = 6
	GDK_SOURCE_TRACKPOINT  InputSource = 7
	GDK_SOURCE_TABLET_PAD  InputSource = 8
)

type NotifyType C.GdkNotifyType

const (
	GDK_NOTIFY_ANCESTOR          NotifyType = 0
	GDK_NOTIFY_VIRTUAL           NotifyType = 1
	GDK_NOTIFY_INFERIOR          NotifyType = 2
	GDK_NOTIFY_NONLINEAR         NotifyType = 3
	GDK_NOTIFY_NONLINEAR_VIRTUAL NotifyType = 4
	GDK_NOTIFY_UNKNOWN           NotifyType = 5
)

type OwnerChange C.GdkOwnerChange

const (
	GDK_OWNER_CHANGE_NEW_OWNER OwnerChange = 0
	GDK_OWNER_CHANGE_DESTROY   OwnerChange = 1
	GDK_OWNER_CHANGE_CLOSE     OwnerChange = 2
)

type PropMode C.GdkPropMode

const (
	GDK_PROP_MODE_REPLACE PropMode = 0
	GDK_PROP_MODE_PREPEND PropMode = 1
	GDK_PROP_MODE_APPEND  PropMode = 2
)

type PropertyState C.guint

const (
	GDK_PROPERTY_NEW_VALUE PropertyState = 0
	GDK_PROPERTY_DELETE    PropertyState = 1
)

type ScrollDirection C.GdkScrollDirection

const (
	GDK_SCROLL_UP     ScrollDirection = 0
	GDK_SCROLL_DOWN   ScrollDirection = 1
	GDK_SCROLL_LEFT   ScrollDirection = 2
	GDK_SCROLL_RIGHT  ScrollDirection = 3
	GDK_SCROLL_SMOOTH ScrollDirection = 4
)

type SettingAction C.GdkSettingAction

const (
	GDK_SETTING_ACTION_NEW     SettingAction = 0
	GDK_SETTING_ACTION_CHANGED SettingAction = 1
	GDK_SETTING_ACTION_DELETED SettingAction = 2
)

type Status C.GdkStatus

const (
	GDK_OK          Status = 0
	GDK_ERROR       Status = -1
	GDK_ERROR_PARAM Status = -2
	GDK_ERROR_FILE  Status = -3
	GDK_ERROR_MEM   Status = -4
)

type TouchpadGesturePhase C.GdkTouchpadGesturePhase

const (
	GDK_TOUCHPAD_GESTURE_PHASE_BEGIN  TouchpadGesturePhase = 0
	GDK_TOUCHPAD_GESTURE_PHASE_UPDATE TouchpadGesturePhase = 1
	GDK_TOUCHPAD_GESTURE_PHASE_END    TouchpadGesturePhase = 2
	GDK_TOUCHPAD_GESTURE_PHASE_CANCEL TouchpadGesturePhase = 3
)

type VisibilityState C.GdkVisibilityState

const (
	GDK_VISIBILITY_UNOBSCURED     VisibilityState = 0
	GDK_VISIBILITY_PARTIAL        VisibilityState = 1
	GDK_VISIBILITY_FULLY_OBSCURED VisibilityState = 2
)

type VisualType C.GdkVisualType

const (
	GDK_VISUAL_STATIC_GRAY  VisualType = 0
	GDK_VISUAL_GRAYSCALE    VisualType = 1
	GDK_VISUAL_STATIC_COLOR VisualType = 2
	GDK_VISUAL_PSEUDO_COLOR VisualType = 3
	GDK_VISUAL_TRUE_COLOR   VisualType = 4
	GDK_VISUAL_DIRECT_COLOR VisualType = 5
)

type WindowEdge C.GdkWindowEdge

const (
	GDK_WINDOW_EDGE_NORTH_WEST WindowEdge = 0
	GDK_WINDOW_EDGE_NORTH      WindowEdge = 1
	GDK_WINDOW_EDGE_NORTH_EAST WindowEdge = 2
	GDK_WINDOW_EDGE_WEST       WindowEdge = 3
	GDK_WINDOW_EDGE_EAST       WindowEdge = 4
	GDK_WINDOW_EDGE_SOUTH_WEST WindowEdge = 5
	GDK_WINDOW_EDGE_SOUTH      WindowEdge = 6
	GDK_WINDOW_EDGE_SOUTH_EAST WindowEdge = 7
)

type WindowType C.GdkWindowType

const (
	GDK_WINDOW_ROOT       WindowType = 0
	GDK_WINDOW_TOPLEVEL   WindowType = 1
	GDK_WINDOW_CHILD      WindowType = 2
	GDK_WINDOW_TEMP       WindowType = 3
	GDK_WINDOW_FOREIGN    WindowType = 4
	GDK_WINDOW_OFFSCREEN  WindowType = 5
	GDK_WINDOW_SUBSURFACE WindowType = 6
)

type WindowTypeHint C.GdkWindowTypeHint

const (
	GDK_WINDOW_TYPE_HINT_NORMAL        WindowTypeHint = 0
	GDK_WINDOW_TYPE_HINT_DIALOG        WindowTypeHint = 1
	GDK_WINDOW_TYPE_HINT_MENU          WindowTypeHint = 2
	GDK_WINDOW_TYPE_HINT_TOOLBAR       WindowTypeHint = 3
	GDK_WINDOW_TYPE_HINT_SPLASHSCREEN  WindowTypeHint = 4
	GDK_WINDOW_TYPE_HINT_UTILITY       WindowTypeHint = 5
	GDK_WINDOW_TYPE_HINT_DOCK          WindowTypeHint = 6
	GDK_WINDOW_TYPE_HINT_DESKTOP       WindowTypeHint = 7
	GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU WindowTypeHint = 8
	GDK_WINDOW_TYPE_HINT_POPUP_MENU    WindowTypeHint = 9
	GDK_WINDOW_TYPE_HINT_TOOLTIP       WindowTypeHint = 10
	GDK_WINDOW_TYPE_HINT_NOTIFICATION  WindowTypeHint = 11
	GDK_WINDOW_TYPE_HINT_COMBO         WindowTypeHint = 12
	GDK_WINDOW_TYPE_HINT_DND           WindowTypeHint = 13
)

type WindowWindowClass C.GdkWindowWindowClass

const (
	GDK_INPUT_OUTPUT WindowWindowClass = 0
	GDK_INPUT_ONLY   WindowWindowClass = 1
)

// AddOptionEntriesLibgtkOnly is a wrapper around the C function gdk_add_option_entries_libgtk_only.
func AddOptionEntriesLibgtkOnly(group *glib.OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.gdk_add_option_entries_libgtk_only(c_group)

	return
}

// Blacklisted : gdk_atom_intern

// Beep is a wrapper around the C function gdk_beep.
func Beep() {
	C.gdk_beep()

	return
}

// CairoCreate is a wrapper around the C function gdk_cairo_create.
func CairoCreate(window *Window) *cairo.Context {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gdk_cairo_create(c_window)
	retGo := cairo.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : gdk_cairo_get_clip_rectangle

// CairoRectangle is a wrapper around the C function gdk_cairo_rectangle.
func CairoRectangle(cr *cairo.Context, rectangle *Rectangle) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_rectangle := (*C.GdkRectangle)(C.NULL)
	if rectangle != nil {
		c_rectangle = (*C.GdkRectangle)(rectangle.ToC())
	}

	C.gdk_cairo_rectangle(c_cr, c_rectangle)

	return
}

// CairoRegion is a wrapper around the C function gdk_cairo_region.
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gdk_cairo_region(c_cr, c_region)

	return
}

// CairoRegionCreateFromSurface is a wrapper around the C function gdk_cairo_region_create_from_surface.
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	retC := C.gdk_cairo_region_create_from_surface(c_surface)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CairoSetSourceColor is a wrapper around the C function gdk_cairo_set_source_color.
func CairoSetSourceColor(cr *cairo.Context, color *Color) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	C.gdk_cairo_set_source_color(c_cr, c_color)

	return
}

// CairoSetSourcePixbuf is a wrapper around the C function gdk_cairo_set_source_pixbuf.
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, pixbufX float64, pixbufY float64) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	c_pixbuf_x := (C.gdouble)(pixbufX)

	c_pixbuf_y := (C.gdouble)(pixbufY)

	C.gdk_cairo_set_source_pixbuf(c_cr, c_pixbuf, c_pixbuf_x, c_pixbuf_y)

	return
}

// Blacklisted : gdk_color_parse

// DragAbort is a wrapper around the C function gdk_drag_abort.
func DragAbort(context *DragContext, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_time_ := (C.guint32)(time)

	C.gdk_drag_abort(c_context, c_time_)

	return
}

// DragBegin is a wrapper around the C function gdk_drag_begin.
func DragBegin(window *Window, targets *glib.List) *DragContext {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	retC := C.gdk_drag_begin(c_window, c_targets)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DragBeginForDevice is a wrapper around the C function gdk_drag_begin_for_device.
func DragBeginForDevice(window *Window, device *Device, targets *glib.List) *DragContext {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	retC := C.gdk_drag_begin_for_device(c_window, c_device, c_targets)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DragDrop is a wrapper around the C function gdk_drag_drop.
func DragDrop(context *DragContext, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_time_ := (C.guint32)(time)

	C.gdk_drag_drop(c_context, c_time_)

	return
}

// DragDropSucceeded is a wrapper around the C function gdk_drag_drop_succeeded.
func DragDropSucceeded(context *DragContext) bool {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	retC := C.gdk_drag_drop_succeeded(c_context)
	retGo := retC == C.TRUE

	return retGo
}

// DragFindWindowForScreen is a wrapper around the C function gdk_drag_find_window_for_screen.
func DragFindWindowForScreen(context *DragContext, dragWindow *Window, screen *Screen, xRoot int32, yRoot int32) (*Window, DragProtocol) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_drag_window := (*C.GdkWindow)(C.NULL)
	if dragWindow != nil {
		c_drag_window = (*C.GdkWindow)(dragWindow.ToC())
	}

	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	c_x_root := (C.gint)(xRoot)

	c_y_root := (C.gint)(yRoot)

	var c_dest_window *C.GdkWindow

	var c_protocol C.GdkDragProtocol

	C.gdk_drag_find_window_for_screen(c_context, c_drag_window, c_screen, c_x_root, c_y_root, &c_dest_window, &c_protocol)

	destWindow := WindowNewFromC(unsafe.Pointer(c_dest_window))

	protocol := (DragProtocol)(c_protocol)

	return destWindow, protocol
}

// DragGetSelection is a wrapper around the C function gdk_drag_get_selection.
func DragGetSelection(context *DragContext) Atom {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	retC := C.gdk_drag_get_selection(c_context)
	retGo := *AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DragMotion is a wrapper around the C function gdk_drag_motion.
func DragMotion(context *DragContext, destWindow *Window, protocol DragProtocol, xRoot int32, yRoot int32, suggestedAction DragAction, possibleActions DragAction, time uint32) bool {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_dest_window := (*C.GdkWindow)(C.NULL)
	if destWindow != nil {
		c_dest_window = (*C.GdkWindow)(destWindow.ToC())
	}

	c_protocol := (C.GdkDragProtocol)(protocol)

	c_x_root := (C.gint)(xRoot)

	c_y_root := (C.gint)(yRoot)

	c_suggested_action := (C.GdkDragAction)(suggestedAction)

	c_possible_actions := (C.GdkDragAction)(possibleActions)

	c_time_ := (C.guint32)(time)

	retC := C.gdk_drag_motion(c_context, c_dest_window, c_protocol, c_x_root, c_y_root, c_suggested_action, c_possible_actions, c_time_)
	retGo := retC == C.TRUE

	return retGo
}

// DragStatus is a wrapper around the C function gdk_drag_status.
func DragStatus(context *DragContext, action DragAction, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_action := (C.GdkDragAction)(action)

	c_time_ := (C.guint32)(time)

	C.gdk_drag_status(c_context, c_action, c_time_)

	return
}

// DropFinish is a wrapper around the C function gdk_drop_finish.
func DropFinish(context *DragContext, success bool, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_success :=
		boolToGboolean(success)

	c_time_ := (C.guint32)(time)

	C.gdk_drop_finish(c_context, c_success, c_time_)

	return
}

// DropReply is a wrapper around the C function gdk_drop_reply.
func DropReply(context *DragContext, accepted bool, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_accepted :=
		boolToGboolean(accepted)

	c_time_ := (C.guint32)(time)

	C.gdk_drop_reply(c_context, c_accepted, c_time_)

	return
}

// ErrorTrapPop is a wrapper around the C function gdk_error_trap_pop.
func ErrorTrapPop() int32 {
	retC := C.gdk_error_trap_pop()
	retGo := (int32)(retC)

	return retGo
}

// ErrorTrapPush is a wrapper around the C function gdk_error_trap_push.
func ErrorTrapPush() {
	C.gdk_error_trap_push()

	return
}

// Unsupported : gdk_event_handler_set : unsupported parameter func : no type generator for EventFunc (GdkEventFunc) for param func

// EventsPending is a wrapper around the C function gdk_events_pending.
func EventsPending() bool {
	retC := C.gdk_events_pending()
	retGo := retC == C.TRUE

	return retGo
}

// Flush is a wrapper around the C function gdk_flush.
func Flush() {
	C.gdk_flush()

	return
}

// GetDefaultRootWindow is a wrapper around the C function gdk_get_default_root_window.
func GetDefaultRootWindow() *Window {
	retC := C.gdk_get_default_root_window()
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_get_display.
func GetDisplay() string {
	retC := C.gdk_get_display()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetDisplayArgName is a wrapper around the C function gdk_get_display_arg_name.
func GetDisplayArgName() string {
	retC := C.gdk_get_display_arg_name()
	retGo := C.GoString(retC)

	return retGo
}

// GetProgramClass is a wrapper around the C function gdk_get_program_class.
func GetProgramClass() string {
	retC := C.gdk_get_program_class()
	retGo := C.GoString(retC)

	return retGo
}

// GetShowEvents is a wrapper around the C function gdk_get_show_events.
func GetShowEvents() bool {
	retC := C.gdk_get_show_events()
	retGo := retC == C.TRUE

	return retGo
}

// Init is a wrapper around the C function gdk_init.
func Init(args []string) []string {
	cArgc, cArgv := argsIn(args)

	C.gdk_init(&cArgc, &cArgv)

	args = argsOut(cArgc, cArgv)

	return args
}

// InitCheck is a wrapper around the C function gdk_init_check.
func InitCheck(args []string) (bool, []string) {
	cArgc, cArgv := argsIn(args)

	retC := C.gdk_init_check(&cArgc, &cArgv)
	retGo := retC == C.TRUE

	args = argsOut(cArgc, cArgv)

	return retGo, args
}

// KeyboardGrab is a wrapper around the C function gdk_keyboard_grab.
func KeyboardGrab(window *Window, ownerEvents bool, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_time_ := (C.guint32)(time)

	retC := C.gdk_keyboard_grab(c_window, c_owner_events, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// KeyboardUngrab is a wrapper around the C function gdk_keyboard_ungrab.
func KeyboardUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_keyboard_ungrab(c_time_)

	return
}

// KeyvalConvertCase is a wrapper around the C function gdk_keyval_convert_case.
func KeyvalConvertCase(symbol uint32) (uint32, uint32) {
	c_symbol := (C.guint)(symbol)

	var c_lower C.guint

	var c_upper C.guint

	C.gdk_keyval_convert_case(c_symbol, &c_lower, &c_upper)

	lower := (uint32)(c_lower)

	upper := (uint32)(c_upper)

	return lower, upper
}

// KeyvalFromName is a wrapper around the C function gdk_keyval_from_name.
func KeyvalFromName(keyvalName string) uint32 {
	c_keyval_name := C.CString(keyvalName)
	defer C.free(unsafe.Pointer(c_keyval_name))

	retC := C.gdk_keyval_from_name(c_keyval_name)
	retGo := (uint32)(retC)

	return retGo
}

// KeyvalIsLower is a wrapper around the C function gdk_keyval_is_lower.
func KeyvalIsLower(keyval uint32) bool {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_is_lower(c_keyval)
	retGo := retC == C.TRUE

	return retGo
}

// KeyvalIsUpper is a wrapper around the C function gdk_keyval_is_upper.
func KeyvalIsUpper(keyval uint32) bool {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_is_upper(c_keyval)
	retGo := retC == C.TRUE

	return retGo
}

// KeyvalName is a wrapper around the C function gdk_keyval_name.
func KeyvalName(keyval uint32) string {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_name(c_keyval)
	retGo := C.GoString(retC)

	return retGo
}

// KeyvalToLower is a wrapper around the C function gdk_keyval_to_lower.
func KeyvalToLower(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_lower(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// KeyvalToUnicode is a wrapper around the C function gdk_keyval_to_unicode.
func KeyvalToUnicode(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_unicode(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// KeyvalToUpper is a wrapper around the C function gdk_keyval_to_upper.
func KeyvalToUpper(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_upper(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// ListVisuals is a wrapper around the C function gdk_list_visuals.
func ListVisuals() *glib.List {
	retC := C.gdk_list_visuals()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// NotifyStartupComplete is a wrapper around the C function gdk_notify_startup_complete.
func NotifyStartupComplete() {
	C.gdk_notify_startup_complete()

	return
}

// OffscreenWindowGetSurface is a wrapper around the C function gdk_offscreen_window_get_surface.
func OffscreenWindowGetSurface(window *Window) *cairo.Surface {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gdk_offscreen_window_get_surface(c_window)
	var retGo (*cairo.Surface)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.SurfaceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PangoContextGet is a wrapper around the C function gdk_pango_context_get.
func PangoContextGet() *pango.Context {
	retC := C.gdk_pango_context_get()
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PangoContextGetForScreen is a wrapper around the C function gdk_pango_context_get_for_screen.
func PangoContextGetForScreen(screen *Screen) *pango.Context {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	retC := C.gdk_pango_context_get_for_screen(c_screen)
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PangoLayoutGetClipRegion is a wrapper around the C function gdk_pango_layout_get_clip_region.
func PangoLayoutGetClipRegion(layout *pango.Layout, xOrigin int32, yOrigin int32, indexRanges int32, nRanges int32) *cairo.Region {
	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	c_x_origin := (C.gint)(xOrigin)

	c_y_origin := (C.gint)(yOrigin)

	c_index_ranges := (C.gint)(indexRanges)

	c_n_ranges := (C.gint)(nRanges)

	retC := C.gdk_pango_layout_get_clip_region(c_layout, c_x_origin, c_y_origin, &c_index_ranges, c_n_ranges)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PangoLayoutLineGetClipRegion is a wrapper around the C function gdk_pango_layout_line_get_clip_region.
func PangoLayoutLineGetClipRegion(line *pango.LayoutLine, xOrigin int32, yOrigin int32, indexRanges []int32, nRanges int32) *cairo.Region {
	c_line := (*C.PangoLayoutLine)(C.NULL)
	if line != nil {
		c_line = (*C.PangoLayoutLine)(line.ToC())
	}

	c_x_origin := (C.gint)(xOrigin)

	c_y_origin := (C.gint)(yOrigin)

	c_index_ranges_array := make([]C.gint, len(indexRanges)+1, len(indexRanges)+1)
	for i, item := range indexRanges {
		c := (C.gint)(item)
		c_index_ranges_array[i] = c
	}
	c_index_ranges_array[len(indexRanges)] = 0
	c_index_ranges_arrayPtr := &c_index_ranges_array[0]
	c_index_ranges := (*C.gint)(unsafe.Pointer(c_index_ranges_arrayPtr))

	c_n_ranges := (C.gint)(nRanges)

	retC := C.gdk_pango_layout_line_get_clip_region(c_line, c_x_origin, c_y_origin, c_index_ranges, c_n_ranges)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParseArgs is a wrapper around the C function gdk_parse_args.
func ParseArgs(args []string) []string {
	cArgc, cArgv := argsIn(args)

	C.gdk_parse_args(&cArgc, &cArgv)

	args = argsOut(cArgc, cArgv)

	return args
}

// PixbufGetFromSurface is a wrapper around the C function gdk_pixbuf_get_from_surface.
func PixbufGetFromSurface(surface *cairo.Surface, srcX int32, srcY int32, width int32, height int32) *gdkpixbuf.Pixbuf {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	c_src_x := (C.gint)(srcX)

	c_src_y := (C.gint)(srcY)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.gdk_pixbuf_get_from_surface(c_surface, c_src_x, c_src_y, c_width, c_height)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PixbufGetFromWindow is a wrapper around the C function gdk_pixbuf_get_from_window.
func PixbufGetFromWindow(window *Window, srcX int32, srcY int32, width int32, height int32) *gdkpixbuf.Pixbuf {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_src_x := (C.gint)(srcX)

	c_src_y := (C.gint)(srcY)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.gdk_pixbuf_get_from_window(c_window, c_src_x, c_src_y, c_width, c_height)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PointerGrab is a wrapper around the C function gdk_pointer_grab.
func PointerGrab(window *Window, ownerEvents bool, eventMask EventMask, confineTo *Window, cursor *Cursor, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_event_mask := (C.GdkEventMask)(eventMask)

	c_confine_to := (*C.GdkWindow)(C.NULL)
	if confineTo != nil {
		c_confine_to = (*C.GdkWindow)(confineTo.ToC())
	}

	c_cursor := (*C.GdkCursor)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkCursor)(cursor.ToC())
	}

	c_time_ := (C.guint32)(time)

	retC := C.gdk_pointer_grab(c_window, c_owner_events, c_event_mask, c_confine_to, c_cursor, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// PointerIsGrabbed is a wrapper around the C function gdk_pointer_is_grabbed.
func PointerIsGrabbed() bool {
	retC := C.gdk_pointer_is_grabbed()
	retGo := retC == C.TRUE

	return retGo
}

// PointerUngrab is a wrapper around the C function gdk_pointer_ungrab.
func PointerUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_pointer_ungrab(c_time_)

	return
}

// PreParseLibgtkOnly is a wrapper around the C function gdk_pre_parse_libgtk_only.
func PreParseLibgtkOnly() {
	C.gdk_pre_parse_libgtk_only()

	return
}

// PropertyChange is a wrapper around the C function gdk_property_change.
func PropertyChange(window *Window, property *Atom, type_ *Atom, format int32, mode PropMode, data uint8, nelements int32) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_property := (C.GdkAtom)(C.NULL)
	if property != nil {
		c_property = (C.GdkAtom)(property.ToC())
	}

	c_type := (C.GdkAtom)(C.NULL)
	if type_ != nil {
		c_type = (C.GdkAtom)(type_.ToC())
	}

	c_format := (C.gint)(format)

	c_mode := (C.GdkPropMode)(mode)

	c_data := (C.guchar)(data)

	c_nelements := (C.gint)(nelements)

	C.gdk_property_change(c_window, c_property, c_type, c_format, c_mode, &c_data, c_nelements)

	return
}

// PropertyDelete is a wrapper around the C function gdk_property_delete.
func PropertyDelete(window *Window, property *Atom) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_property := (C.GdkAtom)(C.NULL)
	if property != nil {
		c_property = (C.GdkAtom)(property.ToC())
	}

	C.gdk_property_delete(c_window, c_property)

	return
}

// Unsupported : gdk_property_get : unsupported parameter actual_length : array length param actual_length is pointer (gint*)

// Unsupported : gdk_query_depths : unsupported parameter depths : output array param depths

// Unsupported : gdk_query_visual_types : unsupported parameter visual_types : output array param visual_types

// SelectionConvert is a wrapper around the C function gdk_selection_convert.
func SelectionConvert(requestor *Window, selection *Atom, target *Atom, time uint32) {
	c_requestor := (*C.GdkWindow)(C.NULL)
	if requestor != nil {
		c_requestor = (*C.GdkWindow)(requestor.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	c_target := (C.GdkAtom)(C.NULL)
	if target != nil {
		c_target = (C.GdkAtom)(target.ToC())
	}

	c_time_ := (C.guint32)(time)

	C.gdk_selection_convert(c_requestor, c_selection, c_target, c_time_)

	return
}

// SelectionOwnerGet is a wrapper around the C function gdk_selection_owner_get.
func SelectionOwnerGet(selection *Atom) *Window {
	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	retC := C.gdk_selection_owner_get(c_selection)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SelectionOwnerGetForDisplay is a wrapper around the C function gdk_selection_owner_get_for_display.
func SelectionOwnerGetForDisplay(display *Display, selection *Atom) *Window {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	retC := C.gdk_selection_owner_get_for_display(c_display, c_selection)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SelectionOwnerSet is a wrapper around the C function gdk_selection_owner_set.
func SelectionOwnerSet(owner *Window, selection *Atom, time uint32, sendEvent bool) bool {
	c_owner := (*C.GdkWindow)(C.NULL)
	if owner != nil {
		c_owner = (*C.GdkWindow)(owner.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	c_time_ := (C.guint32)(time)

	c_send_event :=
		boolToGboolean(sendEvent)

	retC := C.gdk_selection_owner_set(c_owner, c_selection, c_time_, c_send_event)
	retGo := retC == C.TRUE

	return retGo
}

// SelectionOwnerSetForDisplay is a wrapper around the C function gdk_selection_owner_set_for_display.
func SelectionOwnerSetForDisplay(display *Display, owner *Window, selection *Atom, time uint32, sendEvent bool) bool {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_owner := (*C.GdkWindow)(C.NULL)
	if owner != nil {
		c_owner = (*C.GdkWindow)(owner.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	c_time_ := (C.guint32)(time)

	c_send_event :=
		boolToGboolean(sendEvent)

	retC := C.gdk_selection_owner_set_for_display(c_display, c_owner, c_selection, c_time_, c_send_event)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_selection_property_get : unsupported parameter data : guchar** with indirection level of 2

// SelectionSendNotify is a wrapper around the C function gdk_selection_send_notify.
func SelectionSendNotify(requestor *Window, selection *Atom, target *Atom, property *Atom, time uint32) {
	c_requestor := (*C.GdkWindow)(C.NULL)
	if requestor != nil {
		c_requestor = (*C.GdkWindow)(requestor.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	c_target := (C.GdkAtom)(C.NULL)
	if target != nil {
		c_target = (C.GdkAtom)(target.ToC())
	}

	c_property := (C.GdkAtom)(C.NULL)
	if property != nil {
		c_property = (C.GdkAtom)(property.ToC())
	}

	c_time_ := (C.guint32)(time)

	C.gdk_selection_send_notify(c_requestor, c_selection, c_target, c_property, c_time_)

	return
}

// SelectionSendNotifyForDisplay is a wrapper around the C function gdk_selection_send_notify_for_display.
func SelectionSendNotifyForDisplay(display *Display, requestor *Window, selection *Atom, target *Atom, property *Atom, time uint32) {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_requestor := (*C.GdkWindow)(C.NULL)
	if requestor != nil {
		c_requestor = (*C.GdkWindow)(requestor.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	c_target := (C.GdkAtom)(C.NULL)
	if target != nil {
		c_target = (C.GdkAtom)(target.ToC())
	}

	c_property := (C.GdkAtom)(C.NULL)
	if property != nil {
		c_property = (C.GdkAtom)(property.ToC())
	}

	c_time_ := (C.guint32)(time)

	C.gdk_selection_send_notify_for_display(c_display, c_requestor, c_selection, c_target, c_property, c_time_)

	return
}

// SetDoubleClickTime is a wrapper around the C function gdk_set_double_click_time.
func SetDoubleClickTime(msec uint32) {
	c_msec := (C.guint)(msec)

	C.gdk_set_double_click_time(c_msec)

	return
}

// SetProgramClass is a wrapper around the C function gdk_set_program_class.
func SetProgramClass(programClass string) {
	c_program_class := C.CString(programClass)
	defer C.free(unsafe.Pointer(c_program_class))

	C.gdk_set_program_class(c_program_class)

	return
}

// SetShowEvents is a wrapper around the C function gdk_set_show_events.
func SetShowEvents(showEvents bool) {
	c_show_events :=
		boolToGboolean(showEvents)

	C.gdk_set_show_events(c_show_events)

	return
}

// SettingGet is a wrapper around the C function gdk_setting_get.
func SettingGet(name string, value *gobject.Value) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.gdk_setting_get(c_name, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : gdk_synthesize_window_state

// Unsupported : gdk_text_property_to_utf8_list_for_display : unsupported parameter list : output array param list

// ThreadsEnter is a wrapper around the C function gdk_threads_enter.
func ThreadsEnter() {
	C.gdk_threads_enter()

	return
}

// ThreadsInit is a wrapper around the C function gdk_threads_init.
func ThreadsInit() {
	C.gdk_threads_init()

	return
}

// ThreadsLeave is a wrapper around the C function gdk_threads_leave.
func ThreadsLeave() {
	C.gdk_threads_leave()

	return
}

// Unsupported : gdk_threads_set_lock_functions : unsupported parameter enter_fn : no type generator for GObject.Callback (GCallback) for param enter_fn

// UnicodeToKeyval is a wrapper around the C function gdk_unicode_to_keyval.
func UnicodeToKeyval(wc uint32) uint32 {
	c_wc := (C.guint32)(wc)

	retC := C.gdk_unicode_to_keyval(c_wc)
	retGo := (uint32)(retC)

	return retGo
}

// Utf8ToStringTarget is a wrapper around the C function gdk_utf8_to_string_target.
func Utf8ToStringTarget(str string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.gdk_utf8_to_string_target(c_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Atom is a wrapper around the C record GdkAtom.
type Atom struct {
	native *C.GdkAtom
}

func AtomNewFromC(u unsafe.Pointer) *Atom {
	c := (*C.GdkAtom)(u)
	if c == nil {
		return nil
	}

	g := &Atom{native: c}

	return g
}

func (recv *Atom) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Atom with another Atom, and returns true if they represent the same GObject.
func (recv *Atom) Equals(other *Atom) bool {
	return other.ToC() == recv.ToC()
}

// AtomIntern is a wrapper around the C function gdk_atom_intern.
func AtomIntern(atomName string, onlyIfExists bool) Atom {
	c_atom_name := C.CString(atomName)
	defer C.free(unsafe.Pointer(c_atom_name))

	c_only_if_exists :=
		boolToGboolean(onlyIfExists)

	retC := C.gdk_atom_intern(c_atom_name, c_only_if_exists)
	retGo := *AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AtomInternStaticString is a wrapper around the C function gdk_atom_intern_static_string.
func AtomInternStaticString(atomName string) Atom {
	c_atom_name := C.CString(atomName)
	defer C.free(unsafe.Pointer(c_atom_name))

	retC := C.gdk_atom_intern_static_string(c_atom_name)
	retGo := *AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Name is a wrapper around the C function gdk_atom_name.
func (recv *Atom) Name() string {
	retC := C.gdk_atom_name((C.GdkAtom)(*recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Color is a wrapper around the C record GdkColor.
type Color struct {
	native *C.GdkColor
	Pixel  uint32
	Red    uint16
	Green  uint16
	Blue   uint16
}

func ColorNewFromC(u unsafe.Pointer) *Color {
	c := (*C.GdkColor)(u)
	if c == nil {
		return nil
	}

	g := &Color{
		Blue:   (uint16)(c.blue),
		Green:  (uint16)(c.green),
		Pixel:  (uint32)(c.pixel),
		Red:    (uint16)(c.red),
		native: c,
	}

	return g
}

func (recv *Color) ToC() unsafe.Pointer {
	recv.native.pixel =
		(C.guint32)(recv.Pixel)
	recv.native.red =
		(C.guint16)(recv.Red)
	recv.native.green =
		(C.guint16)(recv.Green)
	recv.native.blue =
		(C.guint16)(recv.Blue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Color with another Color, and returns true if they represent the same GObject.
func (recv *Color) Equals(other *Color) bool {
	return other.ToC() == recv.ToC()
}

// ColorParse is a wrapper around the C function gdk_color_parse.
func ColorParse(spec string) (bool, *Color) {
	c_spec := C.CString(spec)
	defer C.free(unsafe.Pointer(c_spec))

	var c_color C.GdkColor

	retC := C.gdk_color_parse(c_spec, &c_color)
	retGo := retC == C.TRUE

	color := ColorNewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// Copy is a wrapper around the C function gdk_color_copy.
func (recv *Color) Copy() *Color {
	retC := C.gdk_color_copy((*C.GdkColor)(recv.native))
	retGo := ColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function gdk_color_equal.
func (recv *Color) Equal(colorb *Color) bool {
	c_colorb := (*C.GdkColor)(C.NULL)
	if colorb != nil {
		c_colorb = (*C.GdkColor)(colorb.ToC())
	}

	retC := C.gdk_color_equal((*C.GdkColor)(recv.native), c_colorb)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function gdk_color_free.
func (recv *Color) Free() {
	C.gdk_color_free((*C.GdkColor)(recv.native))

	return
}

// Hash is a wrapper around the C function gdk_color_hash.
func (recv *Color) Hash() uint32 {
	retC := C.gdk_color_hash((*C.GdkColor)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// EventAny is a wrapper around the C record GdkEventAny.
type EventAny struct {
	native *C.GdkEventAny
	Type   EventType
	// window : record
	SendEvent int8
}

func EventAnyNewFromC(u unsafe.Pointer) *EventAny {
	c := (*C.GdkEventAny)(u)
	if c == nil {
		return nil
	}

	g := &EventAny{
		SendEvent: (int8)(c.send_event),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventAny) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventAny with another EventAny, and returns true if they represent the same GObject.
func (recv *EventAny) Equals(other *EventAny) bool {
	return other.ToC() == recv.ToC()
}

// EventButton is a wrapper around the C record GdkEventButton.
type EventButton struct {
	native *C.GdkEventButton
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	X         float64
	Y         float64
	// axes : gdouble* with indirection level of 1
	State  ModifierType
	Button uint32
	// device : record
	XRoot float64
	YRoot float64
}

func EventButtonNewFromC(u unsafe.Pointer) *EventButton {
	c := (*C.GdkEventButton)(u)
	if c == nil {
		return nil
	}

	g := &EventButton{
		Button:    (uint32)(c.button),
		SendEvent: (int8)(c.send_event),
		State:     (ModifierType)(c.state),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		X:         (float64)(c.x),
		XRoot:     (float64)(c.x_root),
		Y:         (float64)(c.y),
		YRoot:     (float64)(c.y_root),
		native:    c,
	}

	return g
}

func (recv *EventButton) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.x =
		(C.gdouble)(recv.X)
	recv.native.y =
		(C.gdouble)(recv.Y)
	recv.native.state =
		(C.guint)(recv.State)
	recv.native.button =
		(C.guint)(recv.Button)
	recv.native.x_root =
		(C.gdouble)(recv.XRoot)
	recv.native.y_root =
		(C.gdouble)(recv.YRoot)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventButton with another EventButton, and returns true if they represent the same GObject.
func (recv *EventButton) Equals(other *EventButton) bool {
	return other.ToC() == recv.ToC()
}

// EventConfigure is a wrapper around the C record GdkEventConfigure.
type EventConfigure struct {
	native *C.GdkEventConfigure
	Type   EventType
	// window : record
	SendEvent int8
	X         int32
	Y         int32
	Width     int32
	Height    int32
}

func EventConfigureNewFromC(u unsafe.Pointer) *EventConfigure {
	c := (*C.GdkEventConfigure)(u)
	if c == nil {
		return nil
	}

	g := &EventConfigure{
		Height:    (int32)(c.height),
		SendEvent: (int8)(c.send_event),
		Type:      (EventType)(c._type),
		Width:     (int32)(c.width),
		X:         (int32)(c.x),
		Y:         (int32)(c.y),
		native:    c,
	}

	return g
}

func (recv *EventConfigure) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)
	recv.native.width =
		(C.gint)(recv.Width)
	recv.native.height =
		(C.gint)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventConfigure with another EventConfigure, and returns true if they represent the same GObject.
func (recv *EventConfigure) Equals(other *EventConfigure) bool {
	return other.ToC() == recv.ToC()
}

// EventCrossing is a wrapper around the C record GdkEventCrossing.
type EventCrossing struct {
	native *C.GdkEventCrossing
	Type   EventType
	// window : record
	SendEvent int8
	// subwindow : record
	Time   uint32
	X      float64
	Y      float64
	XRoot  float64
	YRoot  float64
	Mode   CrossingMode
	Detail NotifyType
	Focus  bool
	State  ModifierType
}

func EventCrossingNewFromC(u unsafe.Pointer) *EventCrossing {
	c := (*C.GdkEventCrossing)(u)
	if c == nil {
		return nil
	}

	g := &EventCrossing{
		Detail:    (NotifyType)(c.detail),
		Focus:     c.focus == C.TRUE,
		Mode:      (CrossingMode)(c.mode),
		SendEvent: (int8)(c.send_event),
		State:     (ModifierType)(c.state),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		X:         (float64)(c.x),
		XRoot:     (float64)(c.x_root),
		Y:         (float64)(c.y),
		YRoot:     (float64)(c.y_root),
		native:    c,
	}

	return g
}

func (recv *EventCrossing) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.x =
		(C.gdouble)(recv.X)
	recv.native.y =
		(C.gdouble)(recv.Y)
	recv.native.x_root =
		(C.gdouble)(recv.XRoot)
	recv.native.y_root =
		(C.gdouble)(recv.YRoot)
	recv.native.mode =
		(C.GdkCrossingMode)(recv.Mode)
	recv.native.detail =
		(C.GdkNotifyType)(recv.Detail)
	recv.native.focus =
		boolToGboolean(recv.Focus)
	recv.native.state =
		(C.guint)(recv.State)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventCrossing with another EventCrossing, and returns true if they represent the same GObject.
func (recv *EventCrossing) Equals(other *EventCrossing) bool {
	return other.ToC() == recv.ToC()
}

// EventDND is a wrapper around the C record GdkEventDND.
type EventDND struct {
	native *C.GdkEventDND
	Type   EventType
	// window : record
	SendEvent int8
	// context : record
	Time  uint32
	XRoot int16
	YRoot int16
}

func EventDNDNewFromC(u unsafe.Pointer) *EventDND {
	c := (*C.GdkEventDND)(u)
	if c == nil {
		return nil
	}

	g := &EventDND{
		SendEvent: (int8)(c.send_event),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		XRoot:     (int16)(c.x_root),
		YRoot:     (int16)(c.y_root),
		native:    c,
	}

	return g
}

func (recv *EventDND) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.x_root =
		(C.gshort)(recv.XRoot)
	recv.native.y_root =
		(C.gshort)(recv.YRoot)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventDND with another EventDND, and returns true if they represent the same GObject.
func (recv *EventDND) Equals(other *EventDND) bool {
	return other.ToC() == recv.ToC()
}

// EventExpose is a wrapper around the C record GdkEventExpose.
type EventExpose struct {
	native *C.GdkEventExpose
	Type   EventType
	// window : record
	SendEvent int8
	// area : record
	// region : record
	Count int32
}

func EventExposeNewFromC(u unsafe.Pointer) *EventExpose {
	c := (*C.GdkEventExpose)(u)
	if c == nil {
		return nil
	}

	g := &EventExpose{
		Count:     (int32)(c.count),
		SendEvent: (int8)(c.send_event),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventExpose) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.count =
		(C.gint)(recv.Count)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventExpose with another EventExpose, and returns true if they represent the same GObject.
func (recv *EventExpose) Equals(other *EventExpose) bool {
	return other.ToC() == recv.ToC()
}

// EventFocus is a wrapper around the C record GdkEventFocus.
type EventFocus struct {
	native *C.GdkEventFocus
	Type   EventType
	// window : record
	SendEvent int8
	In        int16
}

func EventFocusNewFromC(u unsafe.Pointer) *EventFocus {
	c := (*C.GdkEventFocus)(u)
	if c == nil {
		return nil
	}

	g := &EventFocus{
		In:        (int16)(c.in),
		SendEvent: (int8)(c.send_event),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventFocus) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.in =
		(C.gint16)(recv.In)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventFocus with another EventFocus, and returns true if they represent the same GObject.
func (recv *EventFocus) Equals(other *EventFocus) bool {
	return other.ToC() == recv.ToC()
}

// EventGrabBroken is a wrapper around the C record GdkEventGrabBroken.
type EventGrabBroken struct {
	native *C.GdkEventGrabBroken
	Type   EventType
	// window : record
	SendEvent int8
	Keyboard  bool
	Implicit  bool
	// grab_window : record
}

func EventGrabBrokenNewFromC(u unsafe.Pointer) *EventGrabBroken {
	c := (*C.GdkEventGrabBroken)(u)
	if c == nil {
		return nil
	}

	g := &EventGrabBroken{
		Implicit:  c.implicit == C.TRUE,
		Keyboard:  c.keyboard == C.TRUE,
		SendEvent: (int8)(c.send_event),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventGrabBroken) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.keyboard =
		boolToGboolean(recv.Keyboard)
	recv.native.implicit =
		boolToGboolean(recv.Implicit)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventGrabBroken with another EventGrabBroken, and returns true if they represent the same GObject.
func (recv *EventGrabBroken) Equals(other *EventGrabBroken) bool {
	return other.ToC() == recv.ToC()
}

// EventKey is a wrapper around the C record GdkEventKey.
type EventKey struct {
	native *C.GdkEventKey
	Type   EventType
	// window : record
	SendEvent       int8
	Time            uint32
	State           ModifierType
	Keyval          uint32
	Length          int32
	String          string
	HardwareKeycode uint16
	Group           uint8
	// Bitfield not supported :  1 is_modifier
}

func EventKeyNewFromC(u unsafe.Pointer) *EventKey {
	c := (*C.GdkEventKey)(u)
	if c == nil {
		return nil
	}

	g := &EventKey{
		Group:           (uint8)(c.group),
		HardwareKeycode: (uint16)(c.hardware_keycode),
		Keyval:          (uint32)(c.keyval),
		Length:          (int32)(c.length),
		SendEvent:       (int8)(c.send_event),
		State:           (ModifierType)(c.state),
		String:          C.GoString(c.string),
		Time:            (uint32)(c.time),
		Type:            (EventType)(c._type),
		native:          c,
	}

	return g
}

func (recv *EventKey) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.state =
		(C.guint)(recv.State)
	recv.native.keyval =
		(C.guint)(recv.Keyval)
	recv.native.length =
		(C.gint)(recv.Length)
	recv.native.string =
		C.CString(recv.String)
	recv.native.hardware_keycode =
		(C.guint16)(recv.HardwareKeycode)
	recv.native.group =
		(C.guint8)(recv.Group)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventKey with another EventKey, and returns true if they represent the same GObject.
func (recv *EventKey) Equals(other *EventKey) bool {
	return other.ToC() == recv.ToC()
}

// EventMotion is a wrapper around the C record GdkEventMotion.
type EventMotion struct {
	native *C.GdkEventMotion
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	X         float64
	Y         float64
	// axes : gdouble* with indirection level of 1
	State  ModifierType
	IsHint int16
	// device : record
	XRoot float64
	YRoot float64
}

func EventMotionNewFromC(u unsafe.Pointer) *EventMotion {
	c := (*C.GdkEventMotion)(u)
	if c == nil {
		return nil
	}

	g := &EventMotion{
		IsHint:    (int16)(c.is_hint),
		SendEvent: (int8)(c.send_event),
		State:     (ModifierType)(c.state),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		X:         (float64)(c.x),
		XRoot:     (float64)(c.x_root),
		Y:         (float64)(c.y),
		YRoot:     (float64)(c.y_root),
		native:    c,
	}

	return g
}

func (recv *EventMotion) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.x =
		(C.gdouble)(recv.X)
	recv.native.y =
		(C.gdouble)(recv.Y)
	recv.native.state =
		(C.guint)(recv.State)
	recv.native.is_hint =
		(C.gint16)(recv.IsHint)
	recv.native.x_root =
		(C.gdouble)(recv.XRoot)
	recv.native.y_root =
		(C.gdouble)(recv.YRoot)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventMotion with another EventMotion, and returns true if they represent the same GObject.
func (recv *EventMotion) Equals(other *EventMotion) bool {
	return other.ToC() == recv.ToC()
}

// EventOwnerChange is a wrapper around the C record GdkEventOwnerChange.
type EventOwnerChange struct {
	native *C.GdkEventOwnerChange
	Type   EventType
	// window : record
	SendEvent int8
	// owner : record
	Reason OwnerChange
	// selection : record
	Time          uint32
	SelectionTime uint32
}

func EventOwnerChangeNewFromC(u unsafe.Pointer) *EventOwnerChange {
	c := (*C.GdkEventOwnerChange)(u)
	if c == nil {
		return nil
	}

	g := &EventOwnerChange{
		Reason:        (OwnerChange)(c.reason),
		SelectionTime: (uint32)(c.selection_time),
		SendEvent:     (int8)(c.send_event),
		Time:          (uint32)(c.time),
		Type:          (EventType)(c._type),
		native:        c,
	}

	return g
}

func (recv *EventOwnerChange) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.reason =
		(C.GdkOwnerChange)(recv.Reason)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.selection_time =
		(C.guint32)(recv.SelectionTime)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventOwnerChange with another EventOwnerChange, and returns true if they represent the same GObject.
func (recv *EventOwnerChange) Equals(other *EventOwnerChange) bool {
	return other.ToC() == recv.ToC()
}

// EventProperty is a wrapper around the C record GdkEventProperty.
type EventProperty struct {
	native *C.GdkEventProperty
	Type   EventType
	// window : record
	SendEvent int8
	// atom : record
	Time  uint32
	State PropertyState
}

func EventPropertyNewFromC(u unsafe.Pointer) *EventProperty {
	c := (*C.GdkEventProperty)(u)
	if c == nil {
		return nil
	}

	g := &EventProperty{
		SendEvent: (int8)(c.send_event),
		State:     (PropertyState)(c.state),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventProperty) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.state =
		(C.guint)(recv.State)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventProperty with another EventProperty, and returns true if they represent the same GObject.
func (recv *EventProperty) Equals(other *EventProperty) bool {
	return other.ToC() == recv.ToC()
}

// EventProximity is a wrapper around the C record GdkEventProximity.
type EventProximity struct {
	native *C.GdkEventProximity
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	// device : record
}

func EventProximityNewFromC(u unsafe.Pointer) *EventProximity {
	c := (*C.GdkEventProximity)(u)
	if c == nil {
		return nil
	}

	g := &EventProximity{
		SendEvent: (int8)(c.send_event),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventProximity) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventProximity with another EventProximity, and returns true if they represent the same GObject.
func (recv *EventProximity) Equals(other *EventProximity) bool {
	return other.ToC() == recv.ToC()
}

// EventScroll is a wrapper around the C record GdkEventScroll.
type EventScroll struct {
	native *C.GdkEventScroll
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	X         float64
	Y         float64
	State     ModifierType
	Direction ScrollDirection
	// device : record
	XRoot  float64
	YRoot  float64
	DeltaX float64
	DeltaY float64
	// Bitfield not supported :  1 is_stop
}

func EventScrollNewFromC(u unsafe.Pointer) *EventScroll {
	c := (*C.GdkEventScroll)(u)
	if c == nil {
		return nil
	}

	g := &EventScroll{
		DeltaX:    (float64)(c.delta_x),
		DeltaY:    (float64)(c.delta_y),
		Direction: (ScrollDirection)(c.direction),
		SendEvent: (int8)(c.send_event),
		State:     (ModifierType)(c.state),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		X:         (float64)(c.x),
		XRoot:     (float64)(c.x_root),
		Y:         (float64)(c.y),
		YRoot:     (float64)(c.y_root),
		native:    c,
	}

	return g
}

func (recv *EventScroll) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.x =
		(C.gdouble)(recv.X)
	recv.native.y =
		(C.gdouble)(recv.Y)
	recv.native.state =
		(C.guint)(recv.State)
	recv.native.direction =
		(C.GdkScrollDirection)(recv.Direction)
	recv.native.x_root =
		(C.gdouble)(recv.XRoot)
	recv.native.y_root =
		(C.gdouble)(recv.YRoot)
	recv.native.delta_x =
		(C.gdouble)(recv.DeltaX)
	recv.native.delta_y =
		(C.gdouble)(recv.DeltaY)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventScroll with another EventScroll, and returns true if they represent the same GObject.
func (recv *EventScroll) Equals(other *EventScroll) bool {
	return other.ToC() == recv.ToC()
}

// EventSelection is a wrapper around the C record GdkEventSelection.
type EventSelection struct {
	native *C.GdkEventSelection
	Type   EventType
	// window : record
	SendEvent int8
	// selection : record
	// target : record
	// property : record
	Time uint32
	// requestor : record
}

func EventSelectionNewFromC(u unsafe.Pointer) *EventSelection {
	c := (*C.GdkEventSelection)(u)
	if c == nil {
		return nil
	}

	g := &EventSelection{
		SendEvent: (int8)(c.send_event),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventSelection) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventSelection with another EventSelection, and returns true if they represent the same GObject.
func (recv *EventSelection) Equals(other *EventSelection) bool {
	return other.ToC() == recv.ToC()
}

// EventSequence is a wrapper around the C record GdkEventSequence.
type EventSequence struct {
	native *C.GdkEventSequence
}

func EventSequenceNewFromC(u unsafe.Pointer) *EventSequence {
	c := (*C.GdkEventSequence)(u)
	if c == nil {
		return nil
	}

	g := &EventSequence{native: c}

	return g
}

func (recv *EventSequence) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventSequence with another EventSequence, and returns true if they represent the same GObject.
func (recv *EventSequence) Equals(other *EventSequence) bool {
	return other.ToC() == recv.ToC()
}

// EventSetting is a wrapper around the C record GdkEventSetting.
type EventSetting struct {
	native *C.GdkEventSetting
	Type   EventType
	// window : record
	SendEvent int8
	Action    SettingAction
	Name      string
}

func EventSettingNewFromC(u unsafe.Pointer) *EventSetting {
	c := (*C.GdkEventSetting)(u)
	if c == nil {
		return nil
	}

	g := &EventSetting{
		Action:    (SettingAction)(c.action),
		Name:      C.GoString(c.name),
		SendEvent: (int8)(c.send_event),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventSetting) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.action =
		(C.GdkSettingAction)(recv.Action)
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventSetting with another EventSetting, and returns true if they represent the same GObject.
func (recv *EventSetting) Equals(other *EventSetting) bool {
	return other.ToC() == recv.ToC()
}

// EventTouch is a wrapper around the C record GdkEventTouch.
type EventTouch struct {
	native *C.GdkEventTouch
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	X         float64
	Y         float64
	// axes : gdouble* with indirection level of 1
	State ModifierType
	// sequence : record
	EmulatingPointer bool
	// device : record
	XRoot float64
	YRoot float64
}

func EventTouchNewFromC(u unsafe.Pointer) *EventTouch {
	c := (*C.GdkEventTouch)(u)
	if c == nil {
		return nil
	}

	g := &EventTouch{
		EmulatingPointer: c.emulating_pointer == C.TRUE,
		SendEvent:        (int8)(c.send_event),
		State:            (ModifierType)(c.state),
		Time:             (uint32)(c.time),
		Type:             (EventType)(c._type),
		X:                (float64)(c.x),
		XRoot:            (float64)(c.x_root),
		Y:                (float64)(c.y),
		YRoot:            (float64)(c.y_root),
		native:           c,
	}

	return g
}

func (recv *EventTouch) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.x =
		(C.gdouble)(recv.X)
	recv.native.y =
		(C.gdouble)(recv.Y)
	recv.native.state =
		(C.guint)(recv.State)
	recv.native.emulating_pointer =
		boolToGboolean(recv.EmulatingPointer)
	recv.native.x_root =
		(C.gdouble)(recv.XRoot)
	recv.native.y_root =
		(C.gdouble)(recv.YRoot)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventTouch with another EventTouch, and returns true if they represent the same GObject.
func (recv *EventTouch) Equals(other *EventTouch) bool {
	return other.ToC() == recv.ToC()
}

// EventTouchpadPinch is a wrapper around the C record GdkEventTouchpadPinch.
type EventTouchpadPinch struct {
	native *C.GdkEventTouchpadPinch
	Type   EventType
	// window : record
	SendEvent  int8
	Phase      int8
	NFingers   int8
	Time       uint32
	X          float64
	Y          float64
	Dx         float64
	Dy         float64
	AngleDelta float64
	Scale      float64
	XRoot      float64
	YRoot      float64
	State      ModifierType
}

func EventTouchpadPinchNewFromC(u unsafe.Pointer) *EventTouchpadPinch {
	c := (*C.GdkEventTouchpadPinch)(u)
	if c == nil {
		return nil
	}

	g := &EventTouchpadPinch{
		AngleDelta: (float64)(c.angle_delta),
		Dx:         (float64)(c.dx),
		Dy:         (float64)(c.dy),
		NFingers:   (int8)(c.n_fingers),
		Phase:      (int8)(c.phase),
		Scale:      (float64)(c.scale),
		SendEvent:  (int8)(c.send_event),
		State:      (ModifierType)(c.state),
		Time:       (uint32)(c.time),
		Type:       (EventType)(c._type),
		X:          (float64)(c.x),
		XRoot:      (float64)(c.x_root),
		Y:          (float64)(c.y),
		YRoot:      (float64)(c.y_root),
		native:     c,
	}

	return g
}

func (recv *EventTouchpadPinch) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.phase =
		(C.gint8)(recv.Phase)
	recv.native.n_fingers =
		(C.gint8)(recv.NFingers)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.x =
		(C.gdouble)(recv.X)
	recv.native.y =
		(C.gdouble)(recv.Y)
	recv.native.dx =
		(C.gdouble)(recv.Dx)
	recv.native.dy =
		(C.gdouble)(recv.Dy)
	recv.native.angle_delta =
		(C.gdouble)(recv.AngleDelta)
	recv.native.scale =
		(C.gdouble)(recv.Scale)
	recv.native.x_root =
		(C.gdouble)(recv.XRoot)
	recv.native.y_root =
		(C.gdouble)(recv.YRoot)
	recv.native.state =
		(C.guint)(recv.State)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventTouchpadPinch with another EventTouchpadPinch, and returns true if they represent the same GObject.
func (recv *EventTouchpadPinch) Equals(other *EventTouchpadPinch) bool {
	return other.ToC() == recv.ToC()
}

// EventTouchpadSwipe is a wrapper around the C record GdkEventTouchpadSwipe.
type EventTouchpadSwipe struct {
	native *C.GdkEventTouchpadSwipe
	Type   EventType
	// window : record
	SendEvent int8
	Phase     int8
	NFingers  int8
	Time      uint32
	X         float64
	Y         float64
	Dx        float64
	Dy        float64
	XRoot     float64
	YRoot     float64
	State     ModifierType
}

func EventTouchpadSwipeNewFromC(u unsafe.Pointer) *EventTouchpadSwipe {
	c := (*C.GdkEventTouchpadSwipe)(u)
	if c == nil {
		return nil
	}

	g := &EventTouchpadSwipe{
		Dx:        (float64)(c.dx),
		Dy:        (float64)(c.dy),
		NFingers:  (int8)(c.n_fingers),
		Phase:     (int8)(c.phase),
		SendEvent: (int8)(c.send_event),
		State:     (ModifierType)(c.state),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		X:         (float64)(c.x),
		XRoot:     (float64)(c.x_root),
		Y:         (float64)(c.y),
		YRoot:     (float64)(c.y_root),
		native:    c,
	}

	return g
}

func (recv *EventTouchpadSwipe) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.phase =
		(C.gint8)(recv.Phase)
	recv.native.n_fingers =
		(C.gint8)(recv.NFingers)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.x =
		(C.gdouble)(recv.X)
	recv.native.y =
		(C.gdouble)(recv.Y)
	recv.native.dx =
		(C.gdouble)(recv.Dx)
	recv.native.dy =
		(C.gdouble)(recv.Dy)
	recv.native.x_root =
		(C.gdouble)(recv.XRoot)
	recv.native.y_root =
		(C.gdouble)(recv.YRoot)
	recv.native.state =
		(C.guint)(recv.State)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventTouchpadSwipe with another EventTouchpadSwipe, and returns true if they represent the same GObject.
func (recv *EventTouchpadSwipe) Equals(other *EventTouchpadSwipe) bool {
	return other.ToC() == recv.ToC()
}

// EventVisibility is a wrapper around the C record GdkEventVisibility.
type EventVisibility struct {
	native *C.GdkEventVisibility
	Type   EventType
	// window : record
	SendEvent int8
	State     VisibilityState
}

func EventVisibilityNewFromC(u unsafe.Pointer) *EventVisibility {
	c := (*C.GdkEventVisibility)(u)
	if c == nil {
		return nil
	}

	g := &EventVisibility{
		SendEvent: (int8)(c.send_event),
		State:     (VisibilityState)(c.state),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventVisibility) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.state =
		(C.GdkVisibilityState)(recv.State)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventVisibility with another EventVisibility, and returns true if they represent the same GObject.
func (recv *EventVisibility) Equals(other *EventVisibility) bool {
	return other.ToC() == recv.ToC()
}

// EventWindowState is a wrapper around the C record GdkEventWindowState.
type EventWindowState struct {
	native *C.GdkEventWindowState
	Type   EventType
	// window : record
	SendEvent      int8
	ChangedMask    WindowState
	NewWindowState WindowState
}

func EventWindowStateNewFromC(u unsafe.Pointer) *EventWindowState {
	c := (*C.GdkEventWindowState)(u)
	if c == nil {
		return nil
	}

	g := &EventWindowState{
		ChangedMask:    (WindowState)(c.changed_mask),
		NewWindowState: (WindowState)(c.new_window_state),
		SendEvent:      (int8)(c.send_event),
		Type:           (EventType)(c._type),
		native:         c,
	}

	return g
}

func (recv *EventWindowState) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.changed_mask =
		(C.GdkWindowState)(recv.ChangedMask)
	recv.native.new_window_state =
		(C.GdkWindowState)(recv.NewWindowState)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventWindowState with another EventWindowState, and returns true if they represent the same GObject.
func (recv *EventWindowState) Equals(other *EventWindowState) bool {
	return other.ToC() == recv.ToC()
}

// FrameClockClass is a wrapper around the C record GdkFrameClockClass.
type FrameClockClass struct {
	native *C.GdkFrameClockClass
}

func FrameClockClassNewFromC(u unsafe.Pointer) *FrameClockClass {
	c := (*C.GdkFrameClockClass)(u)
	if c == nil {
		return nil
	}

	g := &FrameClockClass{native: c}

	return g
}

func (recv *FrameClockClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FrameClockClass with another FrameClockClass, and returns true if they represent the same GObject.
func (recv *FrameClockClass) Equals(other *FrameClockClass) bool {
	return other.ToC() == recv.ToC()
}

// FrameClockPrivate is a wrapper around the C record GdkFrameClockPrivate.
type FrameClockPrivate struct {
	native *C.GdkFrameClockPrivate
}

func FrameClockPrivateNewFromC(u unsafe.Pointer) *FrameClockPrivate {
	c := (*C.GdkFrameClockPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FrameClockPrivate{native: c}

	return g
}

func (recv *FrameClockPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FrameClockPrivate with another FrameClockPrivate, and returns true if they represent the same GObject.
func (recv *FrameClockPrivate) Equals(other *FrameClockPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FrameTimings is a wrapper around the C record GdkFrameTimings.
type FrameTimings struct {
	native *C.GdkFrameTimings
}

func FrameTimingsNewFromC(u unsafe.Pointer) *FrameTimings {
	c := (*C.GdkFrameTimings)(u)
	if c == nil {
		return nil
	}

	g := &FrameTimings{native: c}

	return g
}

func (recv *FrameTimings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FrameTimings with another FrameTimings, and returns true if they represent the same GObject.
func (recv *FrameTimings) Equals(other *FrameTimings) bool {
	return other.ToC() == recv.ToC()
}

// GetFrameTime is a wrapper around the C function gdk_frame_timings_get_frame_time.
func (recv *FrameTimings) GetFrameTime() int64 {
	retC := C.gdk_frame_timings_get_frame_time((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Geometry is a wrapper around the C record GdkGeometry.
type Geometry struct {
	native     *C.GdkGeometry
	MinWidth   int32
	MinHeight  int32
	MaxWidth   int32
	MaxHeight  int32
	BaseWidth  int32
	BaseHeight int32
	WidthInc   int32
	HeightInc  int32
	MinAspect  float64
	MaxAspect  float64
	WinGravity Gravity
}

func GeometryNewFromC(u unsafe.Pointer) *Geometry {
	c := (*C.GdkGeometry)(u)
	if c == nil {
		return nil
	}

	g := &Geometry{
		BaseHeight: (int32)(c.base_height),
		BaseWidth:  (int32)(c.base_width),
		HeightInc:  (int32)(c.height_inc),
		MaxAspect:  (float64)(c.max_aspect),
		MaxHeight:  (int32)(c.max_height),
		MaxWidth:   (int32)(c.max_width),
		MinAspect:  (float64)(c.min_aspect),
		MinHeight:  (int32)(c.min_height),
		MinWidth:   (int32)(c.min_width),
		WidthInc:   (int32)(c.width_inc),
		WinGravity: (Gravity)(c.win_gravity),
		native:     c,
	}

	return g
}

func (recv *Geometry) ToC() unsafe.Pointer {
	recv.native.min_width =
		(C.gint)(recv.MinWidth)
	recv.native.min_height =
		(C.gint)(recv.MinHeight)
	recv.native.max_width =
		(C.gint)(recv.MaxWidth)
	recv.native.max_height =
		(C.gint)(recv.MaxHeight)
	recv.native.base_width =
		(C.gint)(recv.BaseWidth)
	recv.native.base_height =
		(C.gint)(recv.BaseHeight)
	recv.native.width_inc =
		(C.gint)(recv.WidthInc)
	recv.native.height_inc =
		(C.gint)(recv.HeightInc)
	recv.native.min_aspect =
		(C.gdouble)(recv.MinAspect)
	recv.native.max_aspect =
		(C.gdouble)(recv.MaxAspect)
	recv.native.win_gravity =
		(C.GdkGravity)(recv.WinGravity)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Geometry with another Geometry, and returns true if they represent the same GObject.
func (recv *Geometry) Equals(other *Geometry) bool {
	return other.ToC() == recv.ToC()
}

// KeymapKey is a wrapper around the C record GdkKeymapKey.
type KeymapKey struct {
	native  *C.GdkKeymapKey
	Keycode uint32
	Group   int32
	Level   int32
}

func KeymapKeyNewFromC(u unsafe.Pointer) *KeymapKey {
	c := (*C.GdkKeymapKey)(u)
	if c == nil {
		return nil
	}

	g := &KeymapKey{
		Group:   (int32)(c.group),
		Keycode: (uint32)(c.keycode),
		Level:   (int32)(c.level),
		native:  c,
	}

	return g
}

func (recv *KeymapKey) ToC() unsafe.Pointer {
	recv.native.keycode =
		(C.guint)(recv.Keycode)
	recv.native.group =
		(C.gint)(recv.Group)
	recv.native.level =
		(C.gint)(recv.Level)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this KeymapKey with another KeymapKey, and returns true if they represent the same GObject.
func (recv *KeymapKey) Equals(other *KeymapKey) bool {
	return other.ToC() == recv.ToC()
}

// Point is a wrapper around the C record GdkPoint.
type Point struct {
	native *C.GdkPoint
	X      int32
	Y      int32
}

func PointNewFromC(u unsafe.Pointer) *Point {
	c := (*C.GdkPoint)(u)
	if c == nil {
		return nil
	}

	g := &Point{
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *Point) ToC() unsafe.Pointer {
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Point with another Point, and returns true if they represent the same GObject.
func (recv *Point) Equals(other *Point) bool {
	return other.ToC() == recv.ToC()
}

// RGBA is a wrapper around the C record GdkRGBA.
type RGBA struct {
	native *C.GdkRGBA
	Red    float64
	Green  float64
	Blue   float64
	Alpha  float64
}

func RGBANewFromC(u unsafe.Pointer) *RGBA {
	c := (*C.GdkRGBA)(u)
	if c == nil {
		return nil
	}

	g := &RGBA{
		Alpha:  (float64)(c.alpha),
		Blue:   (float64)(c.blue),
		Green:  (float64)(c.green),
		Red:    (float64)(c.red),
		native: c,
	}

	return g
}

func (recv *RGBA) ToC() unsafe.Pointer {
	recv.native.red =
		(C.gdouble)(recv.Red)
	recv.native.green =
		(C.gdouble)(recv.Green)
	recv.native.blue =
		(C.gdouble)(recv.Blue)
	recv.native.alpha =
		(C.gdouble)(recv.Alpha)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RGBA with another RGBA, and returns true if they represent the same GObject.
func (recv *RGBA) Equals(other *RGBA) bool {
	return other.ToC() == recv.ToC()
}

// Rectangle is a wrapper around the C record GdkRectangle.
type Rectangle struct {
	native *C.GdkRectangle
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func RectangleNewFromC(u unsafe.Pointer) *Rectangle {
	c := (*C.GdkRectangle)(u)
	if c == nil {
		return nil
	}

	g := &Rectangle{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *Rectangle) ToC() unsafe.Pointer {
	recv.native.x =
		(C.int)(recv.X)
	recv.native.y =
		(C.int)(recv.Y)
	recv.native.width =
		(C.int)(recv.Width)
	recv.native.height =
		(C.int)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Rectangle with another Rectangle, and returns true if they represent the same GObject.
func (recv *Rectangle) Equals(other *Rectangle) bool {
	return other.ToC() == recv.ToC()
}

// Intersect is a wrapper around the C function gdk_rectangle_intersect.
func (recv *Rectangle) Intersect(src2 *Rectangle) (bool, *Rectangle) {
	c_src2 := (*C.GdkRectangle)(C.NULL)
	if src2 != nil {
		c_src2 = (*C.GdkRectangle)(src2.ToC())
	}

	var c_dest C.GdkRectangle

	retC := C.gdk_rectangle_intersect((*C.GdkRectangle)(recv.native), c_src2, &c_dest)
	retGo := retC == C.TRUE

	dest := RectangleNewFromC(unsafe.Pointer(&c_dest))

	return retGo, dest
}

// Union is a wrapper around the C function gdk_rectangle_union.
func (recv *Rectangle) Union(src2 *Rectangle) *Rectangle {
	c_src2 := (*C.GdkRectangle)(C.NULL)
	if src2 != nil {
		c_src2 = (*C.GdkRectangle)(src2.ToC())
	}

	var c_dest C.GdkRectangle

	C.gdk_rectangle_union((*C.GdkRectangle)(recv.native), c_src2, &c_dest)

	dest := RectangleNewFromC(unsafe.Pointer(&c_dest))

	return dest
}

// TimeCoord is a wrapper around the C record GdkTimeCoord.
type TimeCoord struct {
	native *C.GdkTimeCoord
	Time   uint32
	// no type for axes
}

func TimeCoordNewFromC(u unsafe.Pointer) *TimeCoord {
	c := (*C.GdkTimeCoord)(u)
	if c == nil {
		return nil
	}

	g := &TimeCoord{
		Time:   (uint32)(c.time),
		native: c,
	}

	return g
}

func (recv *TimeCoord) ToC() unsafe.Pointer {
	recv.native.time =
		(C.guint32)(recv.Time)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TimeCoord with another TimeCoord, and returns true if they represent the same GObject.
func (recv *TimeCoord) Equals(other *TimeCoord) bool {
	return other.ToC() == recv.ToC()
}

// WindowAttr is a wrapper around the C record GdkWindowAttr.
type WindowAttr struct {
	native    *C.GdkWindowAttr
	Title     string
	EventMask int32
	X         int32
	Y         int32
	Width     int32
	Height    int32
	Wclass    WindowWindowClass
	// visual : record
	WindowType WindowType
	// cursor : record
	WmclassName      string
	WmclassClass     string
	OverrideRedirect bool
	TypeHint         WindowTypeHint
}

func WindowAttrNewFromC(u unsafe.Pointer) *WindowAttr {
	c := (*C.GdkWindowAttr)(u)
	if c == nil {
		return nil
	}

	g := &WindowAttr{
		EventMask:        (int32)(c.event_mask),
		Height:           (int32)(c.height),
		OverrideRedirect: c.override_redirect == C.TRUE,
		Title:            C.GoString(c.title),
		TypeHint:         (WindowTypeHint)(c.type_hint),
		Wclass:           (WindowWindowClass)(c.wclass),
		Width:            (int32)(c.width),
		WindowType:       (WindowType)(c.window_type),
		WmclassClass:     C.GoString(c.wmclass_class),
		WmclassName:      C.GoString(c.wmclass_name),
		X:                (int32)(c.x),
		Y:                (int32)(c.y),
		native:           c,
	}

	return g
}

func (recv *WindowAttr) ToC() unsafe.Pointer {
	recv.native.title =
		C.CString(recv.Title)
	recv.native.event_mask =
		(C.gint)(recv.EventMask)
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)
	recv.native.width =
		(C.gint)(recv.Width)
	recv.native.height =
		(C.gint)(recv.Height)
	recv.native.wclass =
		(C.GdkWindowWindowClass)(recv.Wclass)
	recv.native.window_type =
		(C.GdkWindowType)(recv.WindowType)
	recv.native.wmclass_name =
		C.CString(recv.WmclassName)
	recv.native.wmclass_class =
		C.CString(recv.WmclassClass)
	recv.native.override_redirect =
		boolToGboolean(recv.OverrideRedirect)
	recv.native.type_hint =
		(C.GdkWindowTypeHint)(recv.TypeHint)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WindowAttr with another WindowAttr, and returns true if they represent the same GObject.
func (recv *WindowAttr) Equals(other *WindowAttr) bool {
	return other.ToC() == recv.ToC()
}

// WindowClass is a wrapper around the C record GdkWindowClass.
type WindowClass struct {
	native *C.GdkWindowClass
	// parent_class : record
	// no type for pick_embedded_child
	// no type for to_embedder
	// no type for from_embedder
	// no type for create_surface
	// no type for _gdk_reserved1
	// no type for _gdk_reserved2
	// no type for _gdk_reserved3
	// no type for _gdk_reserved4
	// no type for _gdk_reserved5
	// no type for _gdk_reserved6
	// no type for _gdk_reserved7
	// no type for _gdk_reserved8
}

func WindowClassNewFromC(u unsafe.Pointer) *WindowClass {
	c := (*C.GdkWindowClass)(u)
	if c == nil {
		return nil
	}

	g := &WindowClass{native: c}

	return g
}

func (recv *WindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WindowClass with another WindowClass, and returns true if they represent the same GObject.
func (recv *WindowClass) Equals(other *WindowClass) bool {
	return other.ToC() == recv.ToC()
}

// WindowRedirect is a wrapper around the C record GdkWindowRedirect.
type WindowRedirect struct {
	native *C.GdkWindowRedirect
}

func WindowRedirectNewFromC(u unsafe.Pointer) *WindowRedirect {
	c := (*C.GdkWindowRedirect)(u)
	if c == nil {
		return nil
	}

	g := &WindowRedirect{native: c}

	return g
}

func (recv *WindowRedirect) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WindowRedirect with another WindowRedirect, and returns true if they represent the same GObject.
func (recv *WindowRedirect) Equals(other *WindowRedirect) bool {
	return other.ToC() == recv.ToC()
}

// Event is a wrapper around the C record GdkEvent_.
type Event struct {
	native *C.GdkEvent_
	Type   EventType
}

func EventNewFromC(u unsafe.Pointer) *Event {
	c := (*C.GdkEvent_)(u)
	if c == nil {
		return nil
	}

	g := &Event{
		Type:   (EventType)(c._type),
		native: c,
	}

	return g
}

func (recv *Event) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Event with another Event, and returns true if they represent the same GObject.
func (recv *Event) Equals(other *Event) bool {
	return other.ToC() == recv.ToC()
}
