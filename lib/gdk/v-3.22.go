// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
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

type AnchorHints C.GdkAnchorHints

const (
	GDK_ANCHOR_FLIP_X   AnchorHints = 1
	GDK_ANCHOR_FLIP_Y   AnchorHints = 2
	GDK_ANCHOR_SLIDE_X  AnchorHints = 4
	GDK_ANCHOR_SLIDE_Y  AnchorHints = 8
	GDK_ANCHOR_RESIZE_X AnchorHints = 16
	GDK_ANCHOR_RESIZE_Y AnchorHints = 32
	GDK_ANCHOR_FLIP     AnchorHints = 3
	GDK_ANCHOR_SLIDE    AnchorHints = 12
	GDK_ANCHOR_RESIZE   AnchorHints = 48
)

type AxisFlags C.GdkAxisFlags

const (
	GDK_AXIS_FLAG_X        AxisFlags = 2
	GDK_AXIS_FLAG_Y        AxisFlags = 4
	GDK_AXIS_FLAG_PRESSURE AxisFlags = 8
	GDK_AXIS_FLAG_XTILT    AxisFlags = 16
	GDK_AXIS_FLAG_YTILT    AxisFlags = 32
	GDK_AXIS_FLAG_WHEEL    AxisFlags = 64
	GDK_AXIS_FLAG_DISTANCE AxisFlags = 128
	GDK_AXIS_FLAG_ROTATION AxisFlags = 256
	GDK_AXIS_FLAG_SLIDER   AxisFlags = 512
)

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

// GetAxes is a wrapper around the C function gdk_device_get_axes.
func (recv *Device) GetAxes() AxisFlags {
	retC := C.gdk_device_get_axes((*C.GdkDevice)(recv.native))
	retGo := (AxisFlags)(retC)

	return retGo
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

// GetMonitor is a wrapper around the C function gdk_display_get_monitor.
func (recv *Display) GetMonitor(monitorNum int32) *Monitor {
	c_monitor_num := (C.int)(monitorNum)

	retC := C.gdk_display_get_monitor((*C.GdkDisplay)(recv.native), c_monitor_num)
	var retGo (*Monitor)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MonitorNewFromC(unsafe.Pointer(retC))
	}

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
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

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

// GetPrimaryMonitor is a wrapper around the C function gdk_display_get_primary_monitor.
func (recv *Display) GetPrimaryMonitor() *Monitor {
	retC := C.gdk_display_get_primary_monitor((*C.GdkDisplay)(recv.native))
	var retGo (*Monitor)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MonitorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

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

// GetCairoContext is a wrapper around the C function gdk_drawing_context_get_cairo_context.
func (recv *DrawingContext) GetCairoContext() *cairo.Context {
	retC := C.gdk_drawing_context_get_cairo_context((*C.GdkDrawingContext)(recv.native))
	retGo := cairo.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetClip is a wrapper around the C function gdk_drawing_context_get_clip.
func (recv *DrawingContext) GetClip() *cairo.Region {
	retC := C.gdk_drawing_context_get_clip((*C.GdkDrawingContext)(recv.native))
	var retGo (*cairo.Region)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.RegionNewFromC(unsafe.Pointer(retC))
	}

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

// GetUseEs is a wrapper around the C function gdk_gl_context_get_use_es.
func (recv *GLContext) GetUseEs() bool {
	retC := C.gdk_gl_context_get_use_es((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetUseEs is a wrapper around the C function gdk_gl_context_set_use_es.
func (recv *GLContext) SetUseEs(useEs int32) {
	c_use_es := (C.int)(useEs)

	C.gdk_gl_context_set_use_es((*C.GdkGLContext)(recv.native), c_use_es)

	return
}

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

// GetDisplay is a wrapper around the C function gdk_monitor_get_display.
func (recv *Monitor) GetDisplay() *Display {
	retC := C.gdk_monitor_get_display((*C.GdkMonitor)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGeometry is a wrapper around the C function gdk_monitor_get_geometry.
func (recv *Monitor) GetGeometry() *Rectangle {
	var c_geometry C.GdkRectangle

	C.gdk_monitor_get_geometry((*C.GdkMonitor)(recv.native), &c_geometry)

	geometry := RectangleNewFromC(unsafe.Pointer(&c_geometry))

	return geometry
}

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

// GetWorkarea is a wrapper around the C function gdk_monitor_get_workarea.
func (recv *Monitor) GetWorkarea() *Rectangle {
	var c_workarea C.GdkRectangle

	C.gdk_monitor_get_workarea((*C.GdkMonitor)(recv.native), &c_workarea)

	workarea := RectangleNewFromC(unsafe.Pointer(&c_workarea))

	return workarea
}

// IsPrimary is a wrapper around the C function gdk_monitor_is_primary.
func (recv *Monitor) IsPrimary() bool {
	retC := C.gdk_monitor_is_primary((*C.GdkMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

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

// GetDisplay is a wrapper around the C function gdk_seat_get_display.
func (recv *Seat) GetDisplay() *Display {
	retC := C.gdk_seat_get_display((*C.GdkSeat)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'moved-to-rect' for Window : unsupported parameter flipped_rect : no type generator for gpointer, gpointer

// BeginDrawFrame is a wrapper around the C function gdk_window_begin_draw_frame.
func (recv *Window) BeginDrawFrame(region *cairo.Region) *DrawingContext {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	retC := C.gdk_window_begin_draw_frame((*C.GdkWindow)(recv.native), c_region)
	retGo := DrawingContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EndDrawFrame is a wrapper around the C function gdk_window_end_draw_frame.
func (recv *Window) EndDrawFrame(context *DrawingContext) {
	c_context := (*C.GdkDrawingContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDrawingContext)(context.ToC())
	}

	C.gdk_window_end_draw_frame((*C.GdkWindow)(recv.native), c_context)

	return
}

type DevicePadFeature C.GdkDevicePadFeature

const (
	GDK_DEVICE_PAD_FEATURE_BUTTON DevicePadFeature = 0
	GDK_DEVICE_PAD_FEATURE_RING   DevicePadFeature = 1
	GDK_DEVICE_PAD_FEATURE_STRIP  DevicePadFeature = 2
)

type DeviceToolType C.GdkDeviceToolType

const (
	GDK_DEVICE_TOOL_TYPE_UNKNOWN  DeviceToolType = 0
	GDK_DEVICE_TOOL_TYPE_PEN      DeviceToolType = 1
	GDK_DEVICE_TOOL_TYPE_ERASER   DeviceToolType = 2
	GDK_DEVICE_TOOL_TYPE_BRUSH    DeviceToolType = 3
	GDK_DEVICE_TOOL_TYPE_PENCIL   DeviceToolType = 4
	GDK_DEVICE_TOOL_TYPE_AIRBRUSH DeviceToolType = 5
	GDK_DEVICE_TOOL_TYPE_MOUSE    DeviceToolType = 6
	GDK_DEVICE_TOOL_TYPE_LENS     DeviceToolType = 7
)

type SubpixelLayout C.GdkSubpixelLayout

const (
	GDK_SUBPIXEL_LAYOUT_UNKNOWN        SubpixelLayout = 0
	GDK_SUBPIXEL_LAYOUT_NONE           SubpixelLayout = 1
	GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB SubpixelLayout = 2
	GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR SubpixelLayout = 3
	GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB   SubpixelLayout = 4
	GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR   SubpixelLayout = 5
)

// CairoGetDrawingContext is a wrapper around the C function gdk_cairo_get_drawing_context.
func CairoGetDrawingContext(cr *cairo.Context) *DrawingContext {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	retC := C.gdk_cairo_get_drawing_context(c_cr)
	var retGo (*DrawingContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DrawingContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PangoContextGetForDisplay is a wrapper around the C function gdk_pango_context_get_for_display.
func PangoContextGetForDisplay(display *Display) *pango.Context {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	retC := C.gdk_pango_context_get_for_display(c_display)
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DevicePad is a wrapper around the C record GdkDevicePad.
type DevicePad struct {
	native *C.GdkDevicePad
}

func DevicePadNewFromC(u unsafe.Pointer) *DevicePad {
	c := (*C.GdkDevicePad)(u)
	if c == nil {
		return nil
	}

	g := &DevicePad{native: c}

	return g
}

func (recv *DevicePad) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DevicePad with another DevicePad, and returns true if they represent the same GObject.
func (recv *DevicePad) Equals(other *DevicePad) bool {
	return other.ToC() == recv.ToC()
}

// GetFeatureGroup is a wrapper around the C function gdk_device_pad_get_feature_group.
func (recv *DevicePad) GetFeatureGroup(feature DevicePadFeature, featureIdx int32) int32 {
	c_feature := (C.GdkDevicePadFeature)(feature)

	c_feature_idx := (C.gint)(featureIdx)

	retC := C.gdk_device_pad_get_feature_group((*C.GdkDevicePad)(recv.native), c_feature, c_feature_idx)
	retGo := (int32)(retC)

	return retGo
}

// GetGroupNModes is a wrapper around the C function gdk_device_pad_get_group_n_modes.
func (recv *DevicePad) GetGroupNModes(groupIdx int32) int32 {
	c_group_idx := (C.gint)(groupIdx)

	retC := C.gdk_device_pad_get_group_n_modes((*C.GdkDevicePad)(recv.native), c_group_idx)
	retGo := (int32)(retC)

	return retGo
}

// GetNFeatures is a wrapper around the C function gdk_device_pad_get_n_features.
func (recv *DevicePad) GetNFeatures(feature DevicePadFeature) int32 {
	c_feature := (C.GdkDevicePadFeature)(feature)

	retC := C.gdk_device_pad_get_n_features((*C.GdkDevicePad)(recv.native), c_feature)
	retGo := (int32)(retC)

	return retGo
}

// GetNGroups is a wrapper around the C function gdk_device_pad_get_n_groups.
func (recv *DevicePad) GetNGroups() int32 {
	retC := C.gdk_device_pad_get_n_groups((*C.GdkDevicePad)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// DevicePadInterface is a wrapper around the C record GdkDevicePadInterface.
type DevicePadInterface struct {
	native *C.GdkDevicePadInterface
}

func DevicePadInterfaceNewFromC(u unsafe.Pointer) *DevicePadInterface {
	c := (*C.GdkDevicePadInterface)(u)
	if c == nil {
		return nil
	}

	g := &DevicePadInterface{native: c}

	return g
}

func (recv *DevicePadInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DevicePadInterface with another DevicePadInterface, and returns true if they represent the same GObject.
func (recv *DevicePadInterface) Equals(other *DevicePadInterface) bool {
	return other.ToC() == recv.ToC()
}

// DrawingContextClass is a wrapper around the C record GdkDrawingContextClass.
type DrawingContextClass struct {
	native *C.GdkDrawingContextClass
}

func DrawingContextClassNewFromC(u unsafe.Pointer) *DrawingContextClass {
	c := (*C.GdkDrawingContextClass)(u)
	if c == nil {
		return nil
	}

	g := &DrawingContextClass{native: c}

	return g
}

func (recv *DrawingContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DrawingContextClass with another DrawingContextClass, and returns true if they represent the same GObject.
func (recv *DrawingContextClass) Equals(other *DrawingContextClass) bool {
	return other.ToC() == recv.ToC()
}

// EventPadAxis is a wrapper around the C record GdkEventPadAxis.
type EventPadAxis struct {
	native *C.GdkEventPadAxis
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	Group     uint32
	Index     uint32
	Mode      uint32
	Value     float64
}

func EventPadAxisNewFromC(u unsafe.Pointer) *EventPadAxis {
	c := (*C.GdkEventPadAxis)(u)
	if c == nil {
		return nil
	}

	g := &EventPadAxis{
		Group:     (uint32)(c.group),
		Index:     (uint32)(c.index),
		Mode:      (uint32)(c.mode),
		SendEvent: (int8)(c.send_event),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		Value:     (float64)(c.value),
		native:    c,
	}

	return g
}

func (recv *EventPadAxis) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.group =
		(C.guint)(recv.Group)
	recv.native.index =
		(C.guint)(recv.Index)
	recv.native.mode =
		(C.guint)(recv.Mode)
	recv.native.value =
		(C.gdouble)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventPadAxis with another EventPadAxis, and returns true if they represent the same GObject.
func (recv *EventPadAxis) Equals(other *EventPadAxis) bool {
	return other.ToC() == recv.ToC()
}

// EventPadButton is a wrapper around the C record GdkEventPadButton.
type EventPadButton struct {
	native *C.GdkEventPadButton
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	Group     uint32
	Button    uint32
	Mode      uint32
}

func EventPadButtonNewFromC(u unsafe.Pointer) *EventPadButton {
	c := (*C.GdkEventPadButton)(u)
	if c == nil {
		return nil
	}

	g := &EventPadButton{
		Button:    (uint32)(c.button),
		Group:     (uint32)(c.group),
		Mode:      (uint32)(c.mode),
		SendEvent: (int8)(c.send_event),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventPadButton) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.group =
		(C.guint)(recv.Group)
	recv.native.button =
		(C.guint)(recv.Button)
	recv.native.mode =
		(C.guint)(recv.Mode)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventPadButton with another EventPadButton, and returns true if they represent the same GObject.
func (recv *EventPadButton) Equals(other *EventPadButton) bool {
	return other.ToC() == recv.ToC()
}

// EventPadGroupMode is a wrapper around the C record GdkEventPadGroupMode.
type EventPadGroupMode struct {
	native *C.GdkEventPadGroupMode
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	Group     uint32
	Mode      uint32
}

func EventPadGroupModeNewFromC(u unsafe.Pointer) *EventPadGroupMode {
	c := (*C.GdkEventPadGroupMode)(u)
	if c == nil {
		return nil
	}

	g := &EventPadGroupMode{
		Group:     (uint32)(c.group),
		Mode:      (uint32)(c.mode),
		SendEvent: (int8)(c.send_event),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventPadGroupMode) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.group =
		(C.guint)(recv.Group)
	recv.native.mode =
		(C.guint)(recv.Mode)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventPadGroupMode with another EventPadGroupMode, and returns true if they represent the same GObject.
func (recv *EventPadGroupMode) Equals(other *EventPadGroupMode) bool {
	return other.ToC() == recv.ToC()
}

// MonitorClass is a wrapper around the C record GdkMonitorClass.
type MonitorClass struct {
	native *C.GdkMonitorClass
}

func MonitorClassNewFromC(u unsafe.Pointer) *MonitorClass {
	c := (*C.GdkMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &MonitorClass{native: c}

	return g
}

func (recv *MonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MonitorClass with another MonitorClass, and returns true if they represent the same GObject.
func (recv *MonitorClass) Equals(other *MonitorClass) bool {
	return other.ToC() == recv.ToC()
}
