// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import (
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

	void device_toolChangedHandler(GObject *, GdkDeviceTool *, gpointer);

	static gulong Device_signal_connect_tool_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "tool-changed", G_CALLBACK(device_toolChangedHandler), data);
	}

*/
/*

	void display_monitorAddedHandler(GObject *, GdkMonitor *, gpointer);

	static gulong Display_signal_connect_monitor_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "monitor-added", G_CALLBACK(display_monitorAddedHandler), data);
	}

*/
/*

	void display_monitorRemovedHandler(GObject *, GdkMonitor *, gpointer);

	static gulong Display_signal_connect_monitor_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "monitor-removed", G_CALLBACK(display_monitorRemovedHandler), data);
	}

*/
/*

	void monitor_invalidateHandler(GObject *, gpointer);

	static gulong Monitor_signal_connect_invalidate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "invalidate", G_CALLBACK(monitor_invalidateHandler), data);
	}

*/
/*

	void seat_toolAddedHandler(GObject *, GdkDeviceTool *, gpointer);

	static gulong Seat_signal_connect_tool_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "tool-added", G_CALLBACK(seat_toolAddedHandler), data);
	}

*/
/*

	void seat_toolRemovedHandler(GObject *, GdkDeviceTool *, gpointer);

	static gulong Seat_signal_connect_tool_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "tool-removed", G_CALLBACK(seat_toolRemovedHandler), data);
	}

*/
import "C"

type signalDeviceToolChangedDetail struct {
	callback  DeviceSignalToolChangedCallback
	handlerID C.gulong
}

var signalDeviceToolChangedId int
var signalDeviceToolChangedMap = make(map[int]signalDeviceToolChangedDetail)
var signalDeviceToolChangedLock sync.RWMutex

// DeviceSignalToolChangedCallback is a callback function for a 'tool-changed' signal emitted from a Device.
type DeviceSignalToolChangedCallback func(tool *DeviceTool)

/*
ConnectToolChanged connects the callback to the 'tool-changed' signal for the Device.

The returned value represents the connection, and may be passed to DisconnectToolChanged to remove it.
*/
func (recv *Device) ConnectToolChanged(callback DeviceSignalToolChangedCallback) int {
	signalDeviceToolChangedLock.Lock()
	defer signalDeviceToolChangedLock.Unlock()

	signalDeviceToolChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Device_signal_connect_tool_changed(instance, C.gpointer(uintptr(signalDeviceToolChangedId)))

	detail := signalDeviceToolChangedDetail{callback, handlerID}
	signalDeviceToolChangedMap[signalDeviceToolChangedId] = detail

	return signalDeviceToolChangedId
}

/*
DisconnectToolChanged disconnects a callback from the 'tool-changed' signal for the Device.

The connectionID should be a value returned from a call to ConnectToolChanged.
*/
func (recv *Device) DisconnectToolChanged(connectionID int) {
	signalDeviceToolChangedLock.Lock()
	defer signalDeviceToolChangedLock.Unlock()

	detail, exists := signalDeviceToolChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDeviceToolChangedMap, connectionID)
}

//export device_toolChangedHandler
func device_toolChangedHandler(_ *C.GObject, c_tool *C.GdkDeviceTool, data C.gpointer) {
	signalDeviceToolChangedLock.RLock()
	defer signalDeviceToolChangedLock.RUnlock()

	tool := DeviceToolNewFromC(unsafe.Pointer(c_tool))

	index := int(uintptr(data))
	callback := signalDeviceToolChangedMap[index].callback
	callback(tool)
}

// Blacklisted : gdk_device_get_axes

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DeviceTool) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DeviceTool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DeviceTool with another DeviceTool, and returns true if they represent the same GObject.
func (recv *DeviceTool) Equals(other *DeviceTool) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DeviceTool) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DeviceTool.
// Exercise care, as this is a potentially dangerous function if the Object is not a DeviceTool.
func CastToDeviceTool(object *gobject.Object) *DeviceTool {
	return DeviceToolNewFromC(object.ToC())
}

// Blacklisted : gdk_device_tool_get_hardware_id

// Blacklisted : gdk_device_tool_get_serial

// Blacklisted : gdk_device_tool_get_tool_type

type signalDisplayMonitorAddedDetail struct {
	callback  DisplaySignalMonitorAddedCallback
	handlerID C.gulong
}

var signalDisplayMonitorAddedId int
var signalDisplayMonitorAddedMap = make(map[int]signalDisplayMonitorAddedDetail)
var signalDisplayMonitorAddedLock sync.RWMutex

// DisplaySignalMonitorAddedCallback is a callback function for a 'monitor-added' signal emitted from a Display.
type DisplaySignalMonitorAddedCallback func(monitor *Monitor)

/*
ConnectMonitorAdded connects the callback to the 'monitor-added' signal for the Display.

The returned value represents the connection, and may be passed to DisconnectMonitorAdded to remove it.
*/
func (recv *Display) ConnectMonitorAdded(callback DisplaySignalMonitorAddedCallback) int {
	signalDisplayMonitorAddedLock.Lock()
	defer signalDisplayMonitorAddedLock.Unlock()

	signalDisplayMonitorAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Display_signal_connect_monitor_added(instance, C.gpointer(uintptr(signalDisplayMonitorAddedId)))

	detail := signalDisplayMonitorAddedDetail{callback, handlerID}
	signalDisplayMonitorAddedMap[signalDisplayMonitorAddedId] = detail

	return signalDisplayMonitorAddedId
}

/*
DisconnectMonitorAdded disconnects a callback from the 'monitor-added' signal for the Display.

The connectionID should be a value returned from a call to ConnectMonitorAdded.
*/
func (recv *Display) DisconnectMonitorAdded(connectionID int) {
	signalDisplayMonitorAddedLock.Lock()
	defer signalDisplayMonitorAddedLock.Unlock()

	detail, exists := signalDisplayMonitorAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDisplayMonitorAddedMap, connectionID)
}

//export display_monitorAddedHandler
func display_monitorAddedHandler(_ *C.GObject, c_monitor *C.GdkMonitor, data C.gpointer) {
	signalDisplayMonitorAddedLock.RLock()
	defer signalDisplayMonitorAddedLock.RUnlock()

	monitor := MonitorNewFromC(unsafe.Pointer(c_monitor))

	index := int(uintptr(data))
	callback := signalDisplayMonitorAddedMap[index].callback
	callback(monitor)
}

type signalDisplayMonitorRemovedDetail struct {
	callback  DisplaySignalMonitorRemovedCallback
	handlerID C.gulong
}

var signalDisplayMonitorRemovedId int
var signalDisplayMonitorRemovedMap = make(map[int]signalDisplayMonitorRemovedDetail)
var signalDisplayMonitorRemovedLock sync.RWMutex

// DisplaySignalMonitorRemovedCallback is a callback function for a 'monitor-removed' signal emitted from a Display.
type DisplaySignalMonitorRemovedCallback func(monitor *Monitor)

/*
ConnectMonitorRemoved connects the callback to the 'monitor-removed' signal for the Display.

The returned value represents the connection, and may be passed to DisconnectMonitorRemoved to remove it.
*/
func (recv *Display) ConnectMonitorRemoved(callback DisplaySignalMonitorRemovedCallback) int {
	signalDisplayMonitorRemovedLock.Lock()
	defer signalDisplayMonitorRemovedLock.Unlock()

	signalDisplayMonitorRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Display_signal_connect_monitor_removed(instance, C.gpointer(uintptr(signalDisplayMonitorRemovedId)))

	detail := signalDisplayMonitorRemovedDetail{callback, handlerID}
	signalDisplayMonitorRemovedMap[signalDisplayMonitorRemovedId] = detail

	return signalDisplayMonitorRemovedId
}

/*
DisconnectMonitorRemoved disconnects a callback from the 'monitor-removed' signal for the Display.

The connectionID should be a value returned from a call to ConnectMonitorRemoved.
*/
func (recv *Display) DisconnectMonitorRemoved(connectionID int) {
	signalDisplayMonitorRemovedLock.Lock()
	defer signalDisplayMonitorRemovedLock.Unlock()

	detail, exists := signalDisplayMonitorRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDisplayMonitorRemovedMap, connectionID)
}

//export display_monitorRemovedHandler
func display_monitorRemovedHandler(_ *C.GObject, c_monitor *C.GdkMonitor, data C.gpointer) {
	signalDisplayMonitorRemovedLock.RLock()
	defer signalDisplayMonitorRemovedLock.RUnlock()

	monitor := MonitorNewFromC(unsafe.Pointer(c_monitor))

	index := int(uintptr(data))
	callback := signalDisplayMonitorRemovedMap[index].callback
	callback(monitor)
}

// Blacklisted : gdk_display_get_monitor

// Blacklisted : gdk_display_get_monitor_at_point

// Blacklisted : gdk_display_get_monitor_at_window

// Blacklisted : gdk_display_get_n_monitors

// Blacklisted : gdk_display_get_primary_monitor

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DrawingContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DrawingContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DrawingContext with another DrawingContext, and returns true if they represent the same GObject.
func (recv *DrawingContext) Equals(other *DrawingContext) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DrawingContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DrawingContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a DrawingContext.
func CastToDrawingContext(object *gobject.Object) *DrawingContext {
	return DrawingContextNewFromC(object.ToC())
}

// Blacklisted : gdk_drawing_context_get_cairo_context

// Blacklisted : gdk_drawing_context_get_clip

// Blacklisted : gdk_drawing_context_get_window

// Blacklisted : gdk_drawing_context_is_valid

// Blacklisted : gdk_gl_context_get_use_es

// Blacklisted : gdk_gl_context_set_use_es

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Monitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Monitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Monitor with another Monitor, and returns true if they represent the same GObject.
func (recv *Monitor) Equals(other *Monitor) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Monitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Monitor.
// Exercise care, as this is a potentially dangerous function if the Object is not a Monitor.
func CastToMonitor(object *gobject.Object) *Monitor {
	return MonitorNewFromC(object.ToC())
}

type signalMonitorInvalidateDetail struct {
	callback  MonitorSignalInvalidateCallback
	handlerID C.gulong
}

var signalMonitorInvalidateId int
var signalMonitorInvalidateMap = make(map[int]signalMonitorInvalidateDetail)
var signalMonitorInvalidateLock sync.RWMutex

// MonitorSignalInvalidateCallback is a callback function for a 'invalidate' signal emitted from a Monitor.
type MonitorSignalInvalidateCallback func()

/*
ConnectInvalidate connects the callback to the 'invalidate' signal for the Monitor.

The returned value represents the connection, and may be passed to DisconnectInvalidate to remove it.
*/
func (recv *Monitor) ConnectInvalidate(callback MonitorSignalInvalidateCallback) int {
	signalMonitorInvalidateLock.Lock()
	defer signalMonitorInvalidateLock.Unlock()

	signalMonitorInvalidateId++
	instance := C.gpointer(recv.native)
	handlerID := C.Monitor_signal_connect_invalidate(instance, C.gpointer(uintptr(signalMonitorInvalidateId)))

	detail := signalMonitorInvalidateDetail{callback, handlerID}
	signalMonitorInvalidateMap[signalMonitorInvalidateId] = detail

	return signalMonitorInvalidateId
}

/*
DisconnectInvalidate disconnects a callback from the 'invalidate' signal for the Monitor.

The connectionID should be a value returned from a call to ConnectInvalidate.
*/
func (recv *Monitor) DisconnectInvalidate(connectionID int) {
	signalMonitorInvalidateLock.Lock()
	defer signalMonitorInvalidateLock.Unlock()

	detail, exists := signalMonitorInvalidateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMonitorInvalidateMap, connectionID)
}

//export monitor_invalidateHandler
func monitor_invalidateHandler(_ *C.GObject, data C.gpointer) {
	signalMonitorInvalidateLock.RLock()
	defer signalMonitorInvalidateLock.RUnlock()

	index := int(uintptr(data))
	callback := signalMonitorInvalidateMap[index].callback
	callback()
}

// Blacklisted : gdk_monitor_get_display

// Blacklisted : gdk_monitor_get_geometry

// Blacklisted : gdk_monitor_get_height_mm

// Blacklisted : gdk_monitor_get_manufacturer

// Blacklisted : gdk_monitor_get_model

// Blacklisted : gdk_monitor_get_refresh_rate

// Blacklisted : gdk_monitor_get_scale_factor

// Blacklisted : gdk_monitor_get_subpixel_layout

// Blacklisted : gdk_monitor_get_width_mm

// Blacklisted : gdk_monitor_get_workarea

// Blacklisted : gdk_monitor_is_primary

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Seat) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Seat) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Seat with another Seat, and returns true if they represent the same GObject.
func (recv *Seat) Equals(other *Seat) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Seat) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Seat.
// Exercise care, as this is a potentially dangerous function if the Object is not a Seat.
func CastToSeat(object *gobject.Object) *Seat {
	return SeatNewFromC(object.ToC())
}

type signalSeatToolAddedDetail struct {
	callback  SeatSignalToolAddedCallback
	handlerID C.gulong
}

var signalSeatToolAddedId int
var signalSeatToolAddedMap = make(map[int]signalSeatToolAddedDetail)
var signalSeatToolAddedLock sync.RWMutex

// SeatSignalToolAddedCallback is a callback function for a 'tool-added' signal emitted from a Seat.
type SeatSignalToolAddedCallback func(tool *DeviceTool)

/*
ConnectToolAdded connects the callback to the 'tool-added' signal for the Seat.

The returned value represents the connection, and may be passed to DisconnectToolAdded to remove it.
*/
func (recv *Seat) ConnectToolAdded(callback SeatSignalToolAddedCallback) int {
	signalSeatToolAddedLock.Lock()
	defer signalSeatToolAddedLock.Unlock()

	signalSeatToolAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Seat_signal_connect_tool_added(instance, C.gpointer(uintptr(signalSeatToolAddedId)))

	detail := signalSeatToolAddedDetail{callback, handlerID}
	signalSeatToolAddedMap[signalSeatToolAddedId] = detail

	return signalSeatToolAddedId
}

/*
DisconnectToolAdded disconnects a callback from the 'tool-added' signal for the Seat.

The connectionID should be a value returned from a call to ConnectToolAdded.
*/
func (recv *Seat) DisconnectToolAdded(connectionID int) {
	signalSeatToolAddedLock.Lock()
	defer signalSeatToolAddedLock.Unlock()

	detail, exists := signalSeatToolAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSeatToolAddedMap, connectionID)
}

//export seat_toolAddedHandler
func seat_toolAddedHandler(_ *C.GObject, c_tool *C.GdkDeviceTool, data C.gpointer) {
	signalSeatToolAddedLock.RLock()
	defer signalSeatToolAddedLock.RUnlock()

	tool := DeviceToolNewFromC(unsafe.Pointer(c_tool))

	index := int(uintptr(data))
	callback := signalSeatToolAddedMap[index].callback
	callback(tool)
}

type signalSeatToolRemovedDetail struct {
	callback  SeatSignalToolRemovedCallback
	handlerID C.gulong
}

var signalSeatToolRemovedId int
var signalSeatToolRemovedMap = make(map[int]signalSeatToolRemovedDetail)
var signalSeatToolRemovedLock sync.RWMutex

// SeatSignalToolRemovedCallback is a callback function for a 'tool-removed' signal emitted from a Seat.
type SeatSignalToolRemovedCallback func(tool *DeviceTool)

/*
ConnectToolRemoved connects the callback to the 'tool-removed' signal for the Seat.

The returned value represents the connection, and may be passed to DisconnectToolRemoved to remove it.
*/
func (recv *Seat) ConnectToolRemoved(callback SeatSignalToolRemovedCallback) int {
	signalSeatToolRemovedLock.Lock()
	defer signalSeatToolRemovedLock.Unlock()

	signalSeatToolRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Seat_signal_connect_tool_removed(instance, C.gpointer(uintptr(signalSeatToolRemovedId)))

	detail := signalSeatToolRemovedDetail{callback, handlerID}
	signalSeatToolRemovedMap[signalSeatToolRemovedId] = detail

	return signalSeatToolRemovedId
}

/*
DisconnectToolRemoved disconnects a callback from the 'tool-removed' signal for the Seat.

The connectionID should be a value returned from a call to ConnectToolRemoved.
*/
func (recv *Seat) DisconnectToolRemoved(connectionID int) {
	signalSeatToolRemovedLock.Lock()
	defer signalSeatToolRemovedLock.Unlock()

	detail, exists := signalSeatToolRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSeatToolRemovedMap, connectionID)
}

//export seat_toolRemovedHandler
func seat_toolRemovedHandler(_ *C.GObject, c_tool *C.GdkDeviceTool, data C.gpointer) {
	signalSeatToolRemovedLock.RLock()
	defer signalSeatToolRemovedLock.RUnlock()

	tool := DeviceToolNewFromC(unsafe.Pointer(c_tool))

	index := int(uintptr(data))
	callback := signalSeatToolRemovedMap[index].callback
	callback(tool)
}

// Blacklisted : gdk_seat_get_display

// Unsupported signal 'moved-to-rect' for Window : unsupported parameter flipped_rect : no type generator for gpointer, gpointer

// Blacklisted : gdk_window_begin_draw_frame

// Blacklisted : gdk_window_end_draw_frame
