// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void Device_toolChangedHandler();

	static gulong Device_signal_connect_tool_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "tool-changed", Device_toolChangedHandler, data);
	}

*/
/*

	void Display_monitorAddedHandler();

	static gulong Display_signal_connect_monitor_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "monitor-added", Display_monitorAddedHandler, data);
	}

*/
/*

	void Display_monitorRemovedHandler();

	static gulong Display_signal_connect_monitor_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "monitor-removed", Display_monitorRemovedHandler, data);
	}

*/
/*

	void Monitor_invalidateHandler();

	static gulong Monitor_signal_connect_invalidate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "invalidate", Monitor_invalidateHandler, data);
	}

*/
import "C"

var signalDeviceToolChangedId int
var signalDeviceToolChangedMap = make(map[int]DeviceSignalToolChangedCallback)
var signalDeviceToolChangedLock sync.Mutex

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
	signalDeviceToolChangedMap[signalDeviceToolChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Device_signal_connect_tool_changed(instance, C.gpointer(uintptr(signalDeviceToolChangedId)))
	return int(retC)
}

/*
DisconnectToolChanged disconnects a callback from the 'tool-changed' signal for the Device.

The connectionID should be a value returned from a call to ConnectToolChanged.
*/
func (recv *Device) DisconnectToolChanged(connectionID int) {
	signalDeviceToolChangedLock.Lock()
	defer signalDeviceToolChangedLock.Unlock()

	_, exists := signalDeviceToolChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalDeviceToolChangedMap, connectionID)
}

//export Device_toolChangedHandler
func Device_toolChangedHandler(c_tool *C.GdkDeviceTool) {

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

	return g
}

func (recv *DeviceTool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *DeviceTool) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DeviceTool.
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

var signalDisplayMonitorAddedId int
var signalDisplayMonitorAddedMap = make(map[int]DisplaySignalMonitorAddedCallback)
var signalDisplayMonitorAddedLock sync.Mutex

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
	signalDisplayMonitorAddedMap[signalDisplayMonitorAddedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Display_signal_connect_monitor_added(instance, C.gpointer(uintptr(signalDisplayMonitorAddedId)))
	return int(retC)
}

/*
DisconnectMonitorAdded disconnects a callback from the 'monitor-added' signal for the Display.

The connectionID should be a value returned from a call to ConnectMonitorAdded.
*/
func (recv *Display) DisconnectMonitorAdded(connectionID int) {
	signalDisplayMonitorAddedLock.Lock()
	defer signalDisplayMonitorAddedLock.Unlock()

	_, exists := signalDisplayMonitorAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalDisplayMonitorAddedMap, connectionID)
}

//export Display_monitorAddedHandler
func Display_monitorAddedHandler(c_monitor *C.GdkMonitor) {

}

var signalDisplayMonitorRemovedId int
var signalDisplayMonitorRemovedMap = make(map[int]DisplaySignalMonitorRemovedCallback)
var signalDisplayMonitorRemovedLock sync.Mutex

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
	signalDisplayMonitorRemovedMap[signalDisplayMonitorRemovedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Display_signal_connect_monitor_removed(instance, C.gpointer(uintptr(signalDisplayMonitorRemovedId)))
	return int(retC)
}

/*
DisconnectMonitorRemoved disconnects a callback from the 'monitor-removed' signal for the Display.

The connectionID should be a value returned from a call to ConnectMonitorRemoved.
*/
func (recv *Display) DisconnectMonitorRemoved(connectionID int) {
	signalDisplayMonitorRemovedLock.Lock()
	defer signalDisplayMonitorRemovedLock.Unlock()

	_, exists := signalDisplayMonitorRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalDisplayMonitorRemovedMap, connectionID)
}

//export Display_monitorRemovedHandler
func Display_monitorRemovedHandler(c_monitor *C.GdkMonitor) {

}

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

// GetPrimaryMonitor is a wrapper around the C function gdk_display_get_primary_monitor.
func (recv *Display) GetPrimaryMonitor() *Monitor {
	retC := C.gdk_display_get_primary_monitor((*C.GdkDisplay)(recv.native))
	retGo := MonitorNewFromC(unsafe.Pointer(retC))

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

	return g
}

func (recv *DrawingContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *DrawingContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DrawingContext.
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

	return g
}

func (recv *Monitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Monitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Monitor.
// Exercise care, as this is a potentially dangerous function if the Object is not a Monitor.
func CastToMonitor(object *gobject.Object) *Monitor {
	return MonitorNewFromC(object.ToC())
}

var signalMonitorInvalidateId int
var signalMonitorInvalidateMap = make(map[int]MonitorSignalInvalidateCallback)
var signalMonitorInvalidateLock sync.Mutex

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
	signalMonitorInvalidateMap[signalMonitorInvalidateId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Monitor_signal_connect_invalidate(instance, C.gpointer(uintptr(signalMonitorInvalidateId)))
	return int(retC)
}

/*
DisconnectInvalidate disconnects a callback from the 'invalidate' signal for the Monitor.

The connectionID should be a value returned from a call to ConnectInvalidate.
*/
func (recv *Monitor) DisconnectInvalidate(connectionID int) {
	signalMonitorInvalidateLock.Lock()
	defer signalMonitorInvalidateLock.Unlock()

	_, exists := signalMonitorInvalidateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalMonitorInvalidateMap, connectionID)
}

//export Monitor_invalidateHandler
func Monitor_invalidateHandler() {}

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

// Object upcasts to *Object
func (recv *Seat) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Seat.
// Exercise care, as this is a potentially dangerous function if the Object is not a Seat.
func CastToSeat(object *gobject.Object) *Seat {
	return SeatNewFromC(object.ToC())
}

// GetDisplay is a wrapper around the C function gdk_seat_get_display.
func (recv *Seat) GetDisplay() *Display {
	retC := C.gdk_seat_get_display((*C.GdkSeat)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'moved-to-rect' for Window : unsupported parameter flipped_rect : type gpointer :

// BeginDrawFrame is a wrapper around the C function gdk_window_begin_draw_frame.
func (recv *Window) BeginDrawFrame(region *cairo.Region) *DrawingContext {
	c_region := (*C.cairo_region_t)(region.ToC())

	retC := C.gdk_window_begin_draw_frame((*C.GdkWindow)(recv.native), c_region)
	retGo := DrawingContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EndDrawFrame is a wrapper around the C function gdk_window_end_draw_frame.
func (recv *Window) EndDrawFrame(context *DrawingContext) {
	c_context := (*C.GdkDrawingContext)(context.ToC())

	C.gdk_window_end_draw_frame((*C.GdkWindow)(recv.native), c_context)

	return
}
