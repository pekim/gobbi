// This is a generated file - DO NOT EDIT

package gdk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// Blacklisted : gdk_cursor_new

// Blacklisted : gdk_cursor_ref

// Blacklisted : gdk_cursor_unref

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
// Blacklisted : gdk_device_grab_info_libgtk_only

// Blacklisted : gdk_device_get_axis

// Unsupported : gdk_device_get_history : unsupported parameter events : output array param events

// Unsupported : gdk_device_get_state : unsupported parameter mask : GdkModifierType* with indirection level of 1

// Blacklisted : gdk_device_list_slave_devices

// Blacklisted : gdk_device_set_axis_use

// Blacklisted : gdk_device_set_key

// Blacklisted : gdk_device_set_mode

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

// Blacklisted : gdk_display_open_default_libgtk_only

// Blacklisted : gdk_display_device_is_grabbed

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

// Blacklisted : gdk_drag_context_get_device

// Blacklisted : gdk_drag_context_set_device

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

// Blacklisted : gdk_keymap_get_default

// Blacklisted : gdk_keymap_get_direction

// Unsupported : gdk_keymap_get_entries_for_keycode : unsupported parameter keys : output array param keys

// Unsupported : gdk_keymap_get_entries_for_keyval : unsupported parameter keys : output array param keys

// Blacklisted : gdk_keymap_lookup_key

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

// Blacklisted : gdk_screen_height

// Blacklisted : gdk_screen_height_mm

// Blacklisted : gdk_screen_width

// Blacklisted : gdk_screen_width_mm

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

// Blacklisted : gdk_visual_get_best

// Blacklisted : gdk_visual_get_best_depth

// Blacklisted : gdk_visual_get_best_type

// Blacklisted : gdk_visual_get_best_with_both

// Blacklisted : gdk_visual_get_best_with_depth

// Blacklisted : gdk_visual_get_best_with_type

// Blacklisted : gdk_visual_get_system

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

// Blacklisted : gdk_window_at_pointer

// Blacklisted : gdk_window_constrain_size

// Blacklisted : gdk_window_process_all_updates

// Blacklisted : gdk_window_set_debug_updates

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc (GdkFilterFunc) for param function

// Blacklisted : gdk_window_begin_move_drag

// Blacklisted : gdk_window_begin_paint_rect

// Blacklisted : gdk_window_begin_paint_region

// Blacklisted : gdk_window_begin_resize_drag

// Blacklisted : gdk_window_deiconify

// Destroy is a wrapper around the C function gdk_window_destroy.
func (recv *Window) Destroy() {
	C.gdk_window_destroy((*C.GdkWindow)(recv.native))

	return
}

// Blacklisted : gdk_window_destroy_notify

// Blacklisted : gdk_window_end_paint

// Blacklisted : gdk_window_focus

// Blacklisted : gdk_window_freeze_toplevel_updates_libgtk_only

// Blacklisted : gdk_window_freeze_updates

// Blacklisted : gdk_window_fullscreen_on_monitor

// Blacklisted : gdk_window_get_children

// Blacklisted : gdk_window_get_clip_region

// Unsupported : gdk_window_get_decorations : unsupported parameter decorations : GdkWMDecoration* with indirection level of 1

// Blacklisted : gdk_window_get_events

// Blacklisted : gdk_window_get_frame_extents

// Blacklisted : gdk_window_get_geometry

// Blacklisted : gdk_window_get_origin

// Blacklisted : gdk_window_get_parent

// Unsupported : gdk_window_get_pointer : unsupported parameter mask : GdkModifierType* with indirection level of 1

// Blacklisted : gdk_window_get_position

// Blacklisted : gdk_window_get_root_origin

// Blacklisted : gdk_window_get_source_events

// Blacklisted : gdk_window_get_state

// Blacklisted : gdk_window_get_toplevel

// Blacklisted : gdk_window_get_update_area

// Unsupported : gdk_window_get_user_data : unsupported parameter data : no type generator for gpointer (gpointer*) for param data

// Blacklisted : gdk_window_get_visible_region

// Blacklisted : gdk_window_get_window_type

// Blacklisted : gdk_window_hide

// Blacklisted : gdk_window_iconify

// Unsupported : gdk_window_invalidate_maybe_recurse : unsupported parameter child_func : no type generator for WindowChildFunc (GdkWindowChildFunc) for param child_func

// Blacklisted : gdk_window_invalidate_rect

// Blacklisted : gdk_window_invalidate_region

// Blacklisted : gdk_window_is_viewable

// IsVisible is a wrapper around the C function gdk_window_is_visible.
func (recv *Window) IsVisible() bool {
	retC := C.gdk_window_is_visible((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : gdk_window_lower

// Blacklisted : gdk_window_maximize

// Blacklisted : gdk_window_merge_child_shapes

// Blacklisted : gdk_window_move

// Blacklisted : gdk_window_move_resize

// Blacklisted : gdk_window_peek_children

// Blacklisted : gdk_window_process_updates

// Blacklisted : gdk_window_raise

// Blacklisted : gdk_window_register_dnd

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc (GdkFilterFunc) for param function

// Blacklisted : gdk_window_reparent

// Blacklisted : gdk_window_resize

// Blacklisted : gdk_window_scroll

// Blacklisted : gdk_window_set_background

// Blacklisted : gdk_window_set_background_pattern

// Blacklisted : gdk_window_set_background_rgba

// Blacklisted : gdk_window_set_child_shapes

// Blacklisted : gdk_window_set_cursor

// Blacklisted : gdk_window_set_decorations

// Blacklisted : gdk_window_set_events

// Blacklisted : gdk_window_set_functions

// Blacklisted : gdk_window_set_geometry_hints

// Blacklisted : gdk_window_set_group

// Blacklisted : gdk_window_set_icon_list

// Blacklisted : gdk_window_set_icon_name

// Blacklisted : gdk_window_set_modal_hint

// Blacklisted : gdk_window_set_override_redirect

// Blacklisted : gdk_window_set_role

// Blacklisted : gdk_window_set_static_gravities

// SetTitle is a wrapper around the C function gdk_window_set_title.
func (recv *Window) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gdk_window_set_title((*C.GdkWindow)(recv.native), c_title)

	return
}

// Blacklisted : gdk_window_set_transient_for

// Blacklisted : gdk_window_set_type_hint

// Blacklisted : gdk_window_set_user_data

// Blacklisted : gdk_window_shape_combine_region

// Blacklisted : gdk_window_show

// Blacklisted : gdk_window_show_unraised

// Blacklisted : gdk_window_stick

// Blacklisted : gdk_window_thaw_toplevel_updates_libgtk_only

// Blacklisted : gdk_window_thaw_updates

// Blacklisted : gdk_window_unmaximize

// Blacklisted : gdk_window_unstick

// Blacklisted : gdk_window_withdraw
