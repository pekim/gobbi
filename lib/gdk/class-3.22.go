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
import "C"

type signalDeviceToolChangedDetail struct {
	callback  DeviceSignalToolChangedCallback
	handlerID C.gulong
}

var signalDeviceToolChangedId int
var signalDeviceToolChangedMap = make(map[int]signalDeviceToolChangedDetail)
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
	tool := DeviceToolNewFromC(unsafe.Pointer(c_tool))

	index := int(uintptr(data))
	callback := signalDeviceToolChangedMap[index].callback
	callback(tool)
}

// Returns the axes currently available on the device.
/*

C function

gdk_device_get_axes
*/
func (recv *Device) GetAxes() AxisFlags {
	retC := C.gdk_device_get_axes((*C.GdkDevice)(recv.native))
	retGo := (AxisFlags)(retC)

	return retGo
}

/*

C type

GdkDeviceTool
*/
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

// Gets the hardware ID of this tool, or 0 if it's not known. When
// non-zero, the identificator is unique for the given tool model,
// meaning that two identical tools will share the same @hardware_id,
// but will have different serial numbers (see gdk_device_tool_get_serial()).
//
// This is a more concrete (and device specific) method to identify
// a #GdkDeviceTool than gdk_device_tool_get_tool_type(), as a tablet
// may support multiple devices with the same #GdkDeviceToolType,
// but having different hardware identificators.
/*

C function

gdk_device_tool_get_hardware_id
*/
func (recv *DeviceTool) GetHardwareId() uint64 {
	retC := C.gdk_device_tool_get_hardware_id((*C.GdkDeviceTool)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Gets the serial of this tool, this value can be used to identify a
// physical tool (eg. a tablet pen) across program executions.
/*

C function

gdk_device_tool_get_serial
*/
func (recv *DeviceTool) GetSerial() uint64 {
	retC := C.gdk_device_tool_get_serial((*C.GdkDeviceTool)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Gets the #GdkDeviceToolType of the tool.
/*

C function

gdk_device_tool_get_tool_type
*/
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
	monitor := MonitorNewFromC(unsafe.Pointer(c_monitor))

	index := int(uintptr(data))
	callback := signalDisplayMonitorRemovedMap[index].callback
	callback(monitor)
}

// Gets a monitor associated with this display.
/*

C function

gdk_display_get_monitor
*/
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

// Gets the monitor in which the point (@x, @y) is located,
// or a nearby monitor if the point is not in any monitor.
/*

C function

gdk_display_get_monitor_at_point
*/
func (recv *Display) GetMonitorAtPoint(x int32, y int32) *Monitor {
	c_x := (C.int)(x)

	c_y := (C.int)(y)

	retC := C.gdk_display_get_monitor_at_point((*C.GdkDisplay)(recv.native), c_x, c_y)
	retGo := MonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the monitor in which the largest area of @window
// resides, or a monitor close to @window if it is outside
// of all monitors.
/*

C function

gdk_display_get_monitor_at_window
*/
func (recv *Display) GetMonitorAtWindow(window *Window) *Monitor {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gdk_display_get_monitor_at_window((*C.GdkDisplay)(recv.native), c_window)
	retGo := MonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the number of monitors that belong to @display.
//
// The returned number is valid until the next emission of the
// #GdkDisplay::monitor-added or #GdkDisplay::monitor-removed signal.
/*

C function

gdk_display_get_n_monitors
*/
func (recv *Display) GetNMonitors() int32 {
	retC := C.gdk_display_get_n_monitors((*C.GdkDisplay)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the primary monitor for the display.
//
// The primary monitor is considered the monitor where the “main desktop”
// lives. While normal application windows typically allow the window
// manager to place the windows, specialized desktop applications
// such as panels should place themselves on the primary monitor.
/*

C function

gdk_display_get_primary_monitor
*/
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

// #GdkDrawingContext is an object that represents the current drawing
// state of a #GdkWindow.
//
// It's possible to use a #GdkDrawingContext to draw on a #GdkWindow
// via rendering API like Cairo or OpenGL.
//
// A #GdkDrawingContext can only be created by calling gdk_window_begin_draw_frame()
// and will be valid until a call to gdk_window_end_draw_frame().
//
// #GdkDrawingContext is available since GDK 3.22
/*

C type

GdkDrawingContext
*/
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

// Retrieves a Cairo context to be used to draw on the #GdkWindow
// that created the #GdkDrawingContext.
//
// The returned context is guaranteed to be valid as long as the
// #GdkDrawingContext is valid, that is between a call to
// gdk_window_begin_draw_frame() and gdk_window_end_draw_frame().
/*

C function

gdk_drawing_context_get_cairo_context
*/
func (recv *DrawingContext) GetCairoContext() *cairo.Context {
	retC := C.gdk_drawing_context_get_cairo_context((*C.GdkDrawingContext)(recv.native))
	retGo := cairo.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves a copy of the clip region used when creating the @context.
/*

C function

gdk_drawing_context_get_clip
*/
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

// Retrieves the window that created the drawing @context.
/*

C function

gdk_drawing_context_get_window
*/
func (recv *DrawingContext) GetWindow() *Window {
	retC := C.gdk_drawing_context_get_window((*C.GdkDrawingContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks whether the given #GdkDrawingContext is valid.
/*

C function

gdk_drawing_context_is_valid
*/
func (recv *DrawingContext) IsValid() bool {
	retC := C.gdk_drawing_context_is_valid((*C.GdkDrawingContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks whether the @context is using an OpenGL or OpenGL ES profile.
/*

C function

gdk_gl_context_get_use_es
*/
func (recv *GLContext) GetUseEs() bool {
	retC := C.gdk_gl_context_get_use_es((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Requests that GDK create a OpenGL ES context instead of an OpenGL one,
// if the platform and windowing system allows it.
//
// The @context must not have been realized.
//
// By default, GDK will attempt to automatically detect whether the
// underlying GL implementation is OpenGL or OpenGL ES once the @context
// is realized.
//
// You should check the return value of gdk_gl_context_get_use_es() after
// calling gdk_gl_context_realize() to decide whether to use the OpenGL or
// OpenGL ES API, extensions, or shaders.
/*

C function

gdk_gl_context_set_use_es
*/
func (recv *GLContext) SetUseEs(useEs int32) {
	c_use_es := (C.int)(useEs)

	C.gdk_gl_context_set_use_es((*C.GdkGLContext)(recv.native), c_use_es)

	return
}

// GdkMonitor objects represent the individual outputs that are
// associated with a #GdkDisplay. GdkDisplay has APIs to enumerate
// monitors with gdk_display_get_n_monitors() and gdk_display_get_monitor(), and
// to find particular monitors with gdk_display_get_primary_monitor() or
// gdk_display_get_monitor_at_window().
//
// GdkMonitor was introduced in GTK+ 3.22 and supersedes earlier
// APIs in GdkScreen to obtain monitor-related information.
/*

C type

GdkMonitor
*/
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

type signalMonitorInvalidateDetail struct {
	callback  MonitorSignalInvalidateCallback
	handlerID C.gulong
}

var signalMonitorInvalidateId int
var signalMonitorInvalidateMap = make(map[int]signalMonitorInvalidateDetail)
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
	index := int(uintptr(data))
	callback := signalMonitorInvalidateMap[index].callback
	callback()
}

// Gets the display that this monitor belongs to.
/*

C function

gdk_monitor_get_display
*/
func (recv *Monitor) GetDisplay() *Display {
	retC := C.gdk_monitor_get_display((*C.GdkMonitor)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_monitor_get_geometry : unsupported parameter geometry : Blacklisted record : GdkRectangle

// Gets the height in millimeters of the monitor.
/*

C function

gdk_monitor_get_height_mm
*/
func (recv *Monitor) GetHeightMm() int32 {
	retC := C.gdk_monitor_get_height_mm((*C.GdkMonitor)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the name of the monitor's manufacturer, if available.
/*

C function

gdk_monitor_get_manufacturer
*/
func (recv *Monitor) GetManufacturer() string {
	retC := C.gdk_monitor_get_manufacturer((*C.GdkMonitor)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the a string identifying the monitor model, if available.
/*

C function

gdk_monitor_get_model
*/
func (recv *Monitor) GetModel() string {
	retC := C.gdk_monitor_get_model((*C.GdkMonitor)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the refresh rate of the monitor, if available.
//
// The value is in milli-Hertz, so a refresh rate of 60Hz
// is returned as 60000.
/*

C function

gdk_monitor_get_refresh_rate
*/
func (recv *Monitor) GetRefreshRate() int32 {
	retC := C.gdk_monitor_get_refresh_rate((*C.GdkMonitor)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the internal scale factor that maps from monitor coordinates
// to the actual device pixels. On traditional systems this is 1, but
// on very high density outputs this can be a higher value (often 2).
//
// This can be used if you want to create pixel based data for a
// particular monitor, but most of the time you’re drawing to a window
// where it is better to use gdk_window_get_scale_factor() instead.
/*

C function

gdk_monitor_get_scale_factor
*/
func (recv *Monitor) GetScaleFactor() int32 {
	retC := C.gdk_monitor_get_scale_factor((*C.GdkMonitor)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets information about the layout of red, green and blue
// primaries for each pixel in this monitor, if available.
/*

C function

gdk_monitor_get_subpixel_layout
*/
func (recv *Monitor) GetSubpixelLayout() SubpixelLayout {
	retC := C.gdk_monitor_get_subpixel_layout((*C.GdkMonitor)(recv.native))
	retGo := (SubpixelLayout)(retC)

	return retGo
}

// Gets the width in millimeters of the monitor.
/*

C function

gdk_monitor_get_width_mm
*/
func (recv *Monitor) GetWidthMm() int32 {
	retC := C.gdk_monitor_get_width_mm((*C.GdkMonitor)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_monitor_get_workarea : unsupported parameter workarea : Blacklisted record : GdkRectangle

// Gets whether this monitor should be considered primary
// (see gdk_display_get_primary_monitor()).
/*

C function

gdk_monitor_is_primary
*/
func (recv *Monitor) IsPrimary() bool {
	retC := C.gdk_monitor_is_primary((*C.GdkMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// The #GdkSeat object represents a collection of input devices
// that belong to a user.
/*

C type

GdkSeat
*/
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

// Returns the #GdkDisplay this seat belongs to.
/*

C function

gdk_seat_get_display
*/
func (recv *Seat) GetDisplay() *Display {
	retC := C.gdk_seat_get_display((*C.GdkSeat)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'moved-to-rect' for Window : unsupported parameter flipped_rect : type gpointer :

// Indicates that you are beginning the process of redrawing @region
// on @window, and provides you with a #GdkDrawingContext.
//
// If @window is a top level #GdkWindow, backed by a native window
// implementation, a backing store (offscreen buffer) large enough to
// contain @region will be created. The backing store will be initialized
// with the background color or background surface for @window. Then, all
// drawing operations performed on @window will be diverted to the
// backing store. When you call gdk_window_end_frame(), the contents of
// the backing store will be copied to @window, making it visible
// on screen. Only the part of @window contained in @region will be
// modified; that is, drawing operations are clipped to @region.
//
// The net result of all this is to remove flicker, because the user
// sees the finished product appear all at once when you call
// gdk_window_end_draw_frame(). If you draw to @window directly without
// calling gdk_window_begin_draw_frame(), the user may see flicker
// as individual drawing operations are performed in sequence.
//
// When using GTK+, the widget system automatically places calls to
// gdk_window_begin_draw_frame() and gdk_window_end_draw_frame() around
// emissions of the `GtkWidget::draw` signal. That is, if you’re
// drawing the contents of the widget yourself, you can assume that the
// widget has a cleared background, is already set as the clip region,
// and already has a backing store. Therefore in most cases, application
// code in GTK does not need to call gdk_window_begin_draw_frame()
// explicitly.
/*

C function

gdk_window_begin_draw_frame
*/
func (recv *Window) BeginDrawFrame(region *cairo.Region) *DrawingContext {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	retC := C.gdk_window_begin_draw_frame((*C.GdkWindow)(recv.native), c_region)
	retGo := DrawingContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Indicates that the drawing of the contents of @window started with
// gdk_window_begin_frame() has been completed.
//
// This function will take care of destroying the #GdkDrawingContext.
//
// It is an error to call this function without a matching
// gdk_window_begin_frame() first.
/*

C function

gdk_window_end_draw_frame
*/
func (recv *Window) EndDrawFrame(context *DrawingContext) {
	c_context := (*C.GdkDrawingContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDrawingContext)(context.ToC())
	}

	C.gdk_window_end_draw_frame((*C.GdkWindow)(recv.native), c_context)

	return
}
