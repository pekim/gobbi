// This is a generated file - DO NOT EDIT

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
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

	void display_openedHandler(GObject *, gpointer);

	static gulong Display_signal_connect_opened(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "opened", G_CALLBACK(display_openedHandler), data);
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
	c_axes := &axes[0]

	c_use := (C.GdkAxisUse)(use)

	var c_value C.gdouble

	retC := C.gdk_device_get_axis((*C.GdkDevice)(recv.native), (*C.gdouble)(unsafe.Pointer(c_axes)), c_use, &c_value)
	retGo := retC == C.TRUE

	value := (float64)(c_value)

	return retGo, value
}

// Unsupported : gdk_device_get_history : unsupported parameter events : output array param events

// Unsupported : gdk_device_get_state : unsupported parameter mask : GdkModifierType* with indirection level of 1

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

// KeymapGetDefault is a wrapper around the C function gdk_keymap_get_default.
func KeymapGetDefault() *Keymap {
	retC := C.gdk_keymap_get_default()
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

// Unsupported : gdk_keymap_translate_keyboard_state : unsupported parameter consumed_modifiers : GdkModifierType* with indirection level of 1

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

// WindowNew is a wrapper around the C function gdk_window_new.
func WindowNew(parent *Window, attributes *WindowAttr, attributesMask int32) *Window {
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

// Unsupported : gdk_window_get_pointer : unsupported parameter mask : GdkModifierType* with indirection level of 1

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

// SetUserData is a wrapper around the C function gdk_window_set_user_data.
func (recv *Window) SetUserData(userData uintptr) {
	c_user_data := (C.gpointer)(userData)

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
