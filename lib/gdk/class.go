// This is a generated file - DO NOT EDIT

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
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

// GdkAppLaunchContext is an implementation of #GAppLaunchContext that
// handles launching an application in a graphical context. It provides
// startup notification and allows to launch applications on a specific
// screen or workspace.
//
// Launching an application
//
// |[<!-- language="C" -->
// GdkAppLaunchContext *context;
//
// context = gdk_display_get_app_launch_context (display);
//
// gdk_app_launch_context_set_screen (screen);
// gdk_app_launch_context_set_timestamp (event->time);
//
// if (!g_app_info_launch_default_for_uri ("http://www.gtk.org", context, &error))
// g_warning ("Launching failed: %s\n", error->message);
//
// g_object_unref (context);
// ]|
/*

C type

GdkAppLaunchContext
*/
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

// AppLaunchContext upcasts to *AppLaunchContext
func (recv *AppLaunchContext) AppLaunchContext() *gio.AppLaunchContext {
	return gio.AppLaunchContextNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *AppLaunchContext) Object() *gobject.Object {
	return recv.AppLaunchContext().Object()
}

// CastToWidget down casts any arbitary Object to AppLaunchContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a AppLaunchContext.
func CastToAppLaunchContext(object *gobject.Object) *AppLaunchContext {
	return AppLaunchContextNewFromC(object.ToC())
}

// A #GdkCursor represents a cursor. Its contents are private.
/*

C type

GdkCursor
*/
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

// Object upcasts to *Object
func (recv *Cursor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Cursor.
// Exercise care, as this is a potentially dangerous function if the Object is not a Cursor.
func CastToCursor(object *gobject.Object) *Cursor {
	return CursorNewFromC(object.ToC())
}

// Creates a new cursor from the set of builtin cursors for the default display.
// See gdk_cursor_new_for_display().
//
// To make the cursor invisible, use %GDK_BLANK_CURSOR.
/*

C function

gdk_cursor_new
*/
func CursorNew(cursorType CursorType) *Cursor {
	c_cursor_type := (C.GdkCursorType)(cursorType)

	retC := C.gdk_cursor_new(c_cursor_type)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a reference to @cursor.
/*

C function

gdk_cursor_ref
*/
func (recv *Cursor) Ref() *Cursor {
	retC := C.gdk_cursor_ref((*C.GdkCursor)(recv.native))
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes a reference from @cursor, deallocating the cursor
// if no references remain.
/*

C function

gdk_cursor_unref
*/
func (recv *Cursor) Unref() {
	C.gdk_cursor_unref((*C.GdkCursor)(recv.native))

	return
}

// The #GdkDevice object represents a single input device, such
// as a keyboard, a mouse, a touchpad, etc.
//
// See the #GdkDeviceManager documentation for more information
// about the various kinds of master and slave devices, and their
// relationships.
/*

C type

GdkDevice
*/
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

// Object upcasts to *Object
func (recv *Device) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Device.
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
var signalDeviceChangedLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalDeviceChangedMap[index].callback
	callback()
}

// Interprets an array of double as axis values for a given device,
// and locates the value in the array for a given axis use.
/*

C function

gdk_device_get_axis
*/
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

// If the device if of type %GDK_DEVICE_TYPE_MASTER, it will return
// the list of slave devices attached to it, otherwise it will return
// %NULL
/*

C function

gdk_device_list_slave_devices
*/
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

// Specifies how an axis of a device is used.
/*

C function

gdk_device_set_axis_use
*/
func (recv *Device) SetAxisUse(index uint32, use AxisUse) {
	c_index_ := (C.guint)(index)

	c_use := (C.GdkAxisUse)(use)

	C.gdk_device_set_axis_use((*C.GdkDevice)(recv.native), c_index_, c_use)

	return
}

// Specifies the X key event to generate when a macro button of a device
// is pressed.
/*

C function

gdk_device_set_key
*/
func (recv *Device) SetKey(index uint32, keyval uint32, modifiers ModifierType) {
	c_index_ := (C.guint)(index)

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gdk_device_set_key((*C.GdkDevice)(recv.native), c_index_, c_keyval, c_modifiers)

	return
}

// Sets a the mode of an input device. The mode controls if the
// device is active and whether the device’s range is mapped to the
// entire screen or to a single window.
//
// Note: This is only meaningful for floating devices, master devices (and
// slaves connected to these) drive the pointer cursor, which is not limited
// by the input mode.
/*

C function

gdk_device_set_mode
*/
func (recv *Device) SetMode(mode InputMode) bool {
	c_mode := (C.GdkInputMode)(mode)

	retC := C.gdk_device_set_mode((*C.GdkDevice)(recv.native), c_mode)
	retGo := retC == C.TRUE

	return retGo
}

// In addition to a single pointer and keyboard for user interface input,
// GDK contains support for a variety of input devices, including graphics
// tablets, touchscreens and multiple pointers/keyboards interacting
// simultaneously with the user interface. Such input devices often have
// additional features, such as sub-pixel positioning information and
// additional device-dependent information.
//
// In order to query the device hierarchy and be aware of changes in the
// device hierarchy (such as virtual devices being created or removed, or
// physical devices being plugged or unplugged), GDK provides
// #GdkDeviceManager.
//
// By default, and if the platform supports it, GDK is aware of multiple
// keyboard/pointer pairs and multitouch devices. This behavior can be
// changed by calling gdk_disable_multidevice() before gdk_display_open().
// There should rarely be a need to do that though, since GDK defaults
// to a compatibility mode in which it will emit just one enter/leave
// event pair for all devices on a window. To enable per-device
// enter/leave events and other multi-pointer interaction features,
// gdk_window_set_support_multidevice() must be called on
// #GdkWindows (or gtk_widget_set_support_multidevice() on widgets).
// window. See the gdk_window_set_support_multidevice() documentation
// for more information.
//
// On X11, multi-device support is implemented through XInput 2.
// Unless gdk_disable_multidevice() is called, the XInput 2
// #GdkDeviceManager implementation will be used as the input source.
// Otherwise either the core or XInput 1 implementations will be used.
//
// For simple applications that don’t have any special interest in
// input devices, the so-called “client pointer”
// provides a reasonable approximation to a simple setup with a single
// pointer and keyboard. The device that has been set as the client
// pointer can be accessed via gdk_device_manager_get_client_pointer().
//
// Conceptually, in multidevice mode there are 2 device types. Virtual
// devices (or master devices) are represented by the pointer cursors
// and keyboard foci that are seen on the screen. Physical devices (or
// slave devices) represent the hardware that is controlling the virtual
// devices, and thus have no visible cursor on the screen.
//
// Virtual devices are always paired, so there is a keyboard device for every
// pointer device. Associations between devices may be inspected through
// gdk_device_get_associated_device().
//
// There may be several virtual devices, and several physical devices could
// be controlling each of these virtual devices. Physical devices may also
// be “floating”, which means they are not attached to any virtual device.
//
// # Master and slave devices
//
// |[
// carlos@sacarino:~$ xinput list
// ⎡ Virtual core pointer                          id=2    [master pointer  (3)]
// ⎜   ↳ Virtual core XTEST pointer                id=4    [slave  pointer  (2)]
// ⎜   ↳ Wacom ISDv4 E6 Pen stylus                 id=10   [slave  pointer  (2)]
// ⎜   ↳ Wacom ISDv4 E6 Finger touch               id=11   [slave  pointer  (2)]
// ⎜   ↳ SynPS/2 Synaptics TouchPad                id=13   [slave  pointer  (2)]
// ⎜   ↳ TPPS/2 IBM TrackPoint                     id=14   [slave  pointer  (2)]
// ⎜   ↳ Wacom ISDv4 E6 Pen eraser                 id=16   [slave  pointer  (2)]
// ⎣ Virtual core keyboard                         id=3    [master keyboard (2)]
// ↳ Virtual core XTEST keyboard               id=5    [slave  keyboard (3)]
// ↳ Power Button                              id=6    [slave  keyboard (3)]
// ↳ Video Bus                                 id=7    [slave  keyboard (3)]
// ↳ Sleep Button                              id=8    [slave  keyboard (3)]
// ↳ Integrated Camera                         id=9    [slave  keyboard (3)]
// ↳ AT Translated Set 2 keyboard              id=12   [slave  keyboard (3)]
// ↳ ThinkPad Extra Buttons                    id=15   [slave  keyboard (3)]
// ]|
//
// By default, GDK will automatically listen for events coming from all
// master devices, setting the #GdkDevice for all events coming from input
// devices. Events containing device information are #GDK_MOTION_NOTIFY,
// #GDK_BUTTON_PRESS, #GDK_2BUTTON_PRESS, #GDK_3BUTTON_PRESS,
// #GDK_BUTTON_RELEASE, #GDK_SCROLL, #GDK_KEY_PRESS, #GDK_KEY_RELEASE,
// #GDK_ENTER_NOTIFY, #GDK_LEAVE_NOTIFY, #GDK_FOCUS_CHANGE,
// #GDK_PROXIMITY_IN, #GDK_PROXIMITY_OUT, #GDK_DRAG_ENTER, #GDK_DRAG_LEAVE,
// #GDK_DRAG_MOTION, #GDK_DRAG_STATUS, #GDK_DROP_START, #GDK_DROP_FINISHED
// and #GDK_GRAB_BROKEN. When dealing with an event on a master device,
// it is possible to get the source (slave) device that the event originated
// from via gdk_event_get_source_device().
//
// On a standard session, all physical devices are connected by default to
// the "Virtual Core Pointer/Keyboard" master devices, hence routing all events
// through these. This behavior is only modified by device grabs, where the
// slave device is temporarily detached for as long as the grab is held, and
// more permanently by user modifications to the device hierarchy.
//
// On certain application specific setups, it may make sense
// to detach a physical device from its master pointer, and mapping it to
// an specific window. This can be achieved by the combination of
// gdk_device_grab() and gdk_device_set_mode().
//
// In order to listen for events coming from devices
// other than a virtual device, gdk_window_set_device_events() must be
// called. Generally, this function can be used to modify the event mask
// for any given device.
//
// Input devices may also provide additional information besides X/Y.
// For example, graphics tablets may also provide pressure and X/Y tilt
// information. This information is device-dependent, and may be
// queried through gdk_device_get_axis(). In multidevice mode, virtual
// devices will change axes in order to always represent the physical
// device that is routing events through it. Whenever the physical device
// changes, the #GdkDevice:n-axes property will be notified, and
// gdk_device_list_axes() will return the new device axes.
//
// Devices may also have associated “keys” or
// macro buttons. Such keys can be globally set to map into normal X
// keyboard events. The mapping is set using gdk_device_set_key().
//
// In GTK+ 3.20, a new #GdkSeat object has been introduced that
// supersedes #GdkDeviceManager and should be preferred in newly
// written code.
/*

C type

GdkDeviceManager
*/
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

// Object upcasts to *Object
func (recv *DeviceManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DeviceManager.
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
var signalDeviceManagerDeviceAddedLock sync.Mutex

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
var signalDeviceManagerDeviceChangedLock sync.Mutex

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
var signalDeviceManagerDeviceRemovedLock sync.Mutex

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
	device := DeviceNewFromC(unsafe.Pointer(c_device))

	index := int(uintptr(data))
	callback := signalDeviceManagerDeviceRemovedMap[index].callback
	callback(device)
}

// #GdkDisplay objects purpose are two fold:
//
// - To manage and provide information about input devices (pointers and keyboards)
//
// - To manage and provide information about the available #GdkScreens
//
// GdkDisplay objects are the GDK representation of an X Display,
// which can be described as a workstation consisting of
// a keyboard, a pointing device (such as a mouse) and one or more
// screens.
// It is used to open and keep track of various GdkScreen objects
// currently instantiated by the application. It is also used to
// access the keyboard(s) and mouse pointer(s) of the display.
//
// Most of the input device handling has been factored out into
// the separate #GdkDeviceManager object. Every display has a
// device manager, which you can obtain using
// gdk_display_get_device_manager().
/*

C type

GdkDisplay
*/
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

// Object upcasts to *Object
func (recv *Display) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Display.
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
var signalDisplayOpenedLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalDisplayOpenedMap[index].callback
	callback()
}

// Returns %TRUE if there is an ongoing grab on @device for @display.
/*

C function

gdk_display_device_is_grabbed
*/
func (recv *Display) DeviceIsGrabbed(device *Device) bool {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gdk_display_device_is_grabbed((*C.GdkDisplay)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// The purpose of the #GdkDisplayManager singleton object is to offer
// notification when displays appear or disappear or the default display
// changes.
//
// You can use gdk_display_manager_get() to obtain the #GdkDisplayManager
// singleton, but that should be rarely necessary. Typically, initializing
// GTK+ opens a display that you can work with without ever accessing the
// #GdkDisplayManager.
//
// The GDK library can be built with support for multiple backends.
// The #GdkDisplayManager object determines which backend is used
// at runtime.
//
// When writing backend-specific code that is supposed to work with
// multiple GDK backends, you have to consider both compile time and
// runtime. At compile time, use the #GDK_WINDOWING_X11, #GDK_WINDOWING_WIN32
// macros, etc. to find out which backends are present in the GDK library
// you are building your application against. At runtime, use type-check
// macros like GDK_IS_X11_DISPLAY() to find out which backend is in use:
//
/*
Backend-specific code

(see backend-specific)
*/
//
// |[<!-- language="C" -->
// #ifdef GDK_WINDOWING_X11
// if (GDK_IS_X11_DISPLAY (display))
// {
// make X11-specific calls here
// }
// else
// #endif
// #ifdef GDK_WINDOWING_QUARTZ
// if (GDK_IS_QUARTZ_DISPLAY (display))
// {
// make Quartz-specific calls here
// }
// else
// #endif
// g_error ("Unsupported GDK backend");
// ]|
/*

C type

GdkDisplayManager
*/
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

// Object upcasts to *Object
func (recv *DisplayManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DisplayManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a DisplayManager.
func CastToDisplayManager(object *gobject.Object) *DisplayManager {
	return DisplayManagerNewFromC(object.ToC())
}

/*

C type

GdkDragContext
*/
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

// Object upcasts to *Object
func (recv *DragContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DragContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a DragContext.
func CastToDragContext(object *gobject.Object) *DragContext {
	return DragContextNewFromC(object.ToC())
}

// Returns the #GdkDevice associated to the drag context.
/*

C function

gdk_drag_context_get_device
*/
func (recv *DragContext) GetDevice() *Device {
	retC := C.gdk_drag_context_get_device((*C.GdkDragContext)(recv.native))
	retGo := DeviceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Associates a #GdkDevice to @context, so all Drag and Drop events
// for @context are emitted as if they came from this device.
/*

C function

gdk_drag_context_set_device
*/
func (recv *DragContext) SetDevice(device *Device) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	C.gdk_drag_context_set_device((*C.GdkDragContext)(recv.native), c_device)

	return
}

// A #GdkFrameClock tells the application when to update and repaint a
// window. This may be synced to the vertical refresh rate of the
// monitor, for example. Even when the frame clock uses a simple timer
// rather than a hardware-based vertical sync, the frame clock helps
// because it ensures everything paints at the same time (reducing the
// total number of frames). The frame clock can also automatically
// stop painting when it knows the frames will not be visible, or
// scale back animation framerates.
//
// #GdkFrameClock is designed to be compatible with an OpenGL-based
// implementation or with mozRequestAnimationFrame in Firefox,
// for example.
//
// A frame clock is idle until someone requests a frame with
// gdk_frame_clock_request_phase(). At some later point that makes
// sense for the synchronization being implemented, the clock will
// process a frame and emit signals for each phase that has been
// requested. (See the signals of the #GdkFrameClock class for
// documentation of the phases. %GDK_FRAME_CLOCK_PHASE_UPDATE and the
// #GdkFrameClock::update signal are most interesting for application
// writers, and are used to update the animations, using the frame time
// given by gdk_frame_clock_get_frame_time().
//
// The frame time is reported in microseconds and generally in the same
// timescale as g_get_monotonic_time(), however, it is not the same
// as g_get_monotonic_time(). The frame time does not advance during
// the time a frame is being painted, and outside of a frame, an attempt
// is made so that all calls to gdk_frame_clock_get_frame_time() that
// are called at a “similar” time get the same value. This means that
// if different animations are timed by looking at the difference in
// time between an initial value from gdk_frame_clock_get_frame_time()
// and the value inside the #GdkFrameClock::update signal of the clock,
// they will stay exactly synchronized.
/*

C type

GdkFrameClock
*/
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

// Object upcasts to *Object
func (recv *FrameClock) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FrameClock.
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
var signalFrameClockAfterPaintLock sync.Mutex

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
var signalFrameClockBeforePaintLock sync.Mutex

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
var signalFrameClockFlushEventsLock sync.Mutex

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
var signalFrameClockLayoutLock sync.Mutex

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
var signalFrameClockPaintLock sync.Mutex

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
var signalFrameClockResumeEventsLock sync.Mutex

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
var signalFrameClockUpdateLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalFrameClockUpdateMap[index].callback
	callback()
}

// #GdkGLContext is an object representing the platform-specific
// OpenGL drawing context.
//
// #GdkGLContexts are created for a #GdkWindow using
// gdk_window_create_gl_context(), and the context will match
// the #GdkVisual of the window.
//
// A #GdkGLContext is not tied to any particular normal framebuffer.
// For instance, it cannot draw to the #GdkWindow back buffer. The GDK
// repaint system is in full control of the painting to that. Instead,
// you can create render buffers or textures and use gdk_cairo_draw_from_gl()
// in the draw function of your widget to draw them. Then GDK will handle
// the integration of your rendering with that of other widgets.
//
// Support for #GdkGLContext is platform-specific, context creation
// can fail, returning %NULL context.
//
// A #GdkGLContext has to be made "current" in order to start using
// it, otherwise any OpenGL call will be ignored.
//
// Creating a new OpenGL context ##
//
// In order to create a new #GdkGLContext instance you need a
// #GdkWindow, which you typically get during the realize call
// of a widget.
//
// A #GdkGLContext is not realized until either gdk_gl_context_make_current(),
// or until it is realized using gdk_gl_context_realize(). It is possible to
// specify details of the GL context like the OpenGL version to be used, or
// whether the GL context should have extra state validation enabled after
// calling gdk_window_create_gl_context() by calling gdk_gl_context_realize().
// If the realization fails you have the option to change the settings of the
// #GdkGLContext and try again.
//
// Using a GdkGLContext ##
//
// You will need to make the #GdkGLContext the current context
// before issuing OpenGL calls; the system sends OpenGL commands to
// whichever context is current. It is possible to have multiple
// contexts, so you always need to ensure that the one which you
// want to draw with is the current one before issuing commands:
//
// |[<!-- language="C" -->
// gdk_gl_context_make_current (context);
// ]|
//
// You can now perform your drawing using OpenGL commands.
//
// You can check which #GdkGLContext is the current one by using
// gdk_gl_context_get_current(); you can also unset any #GdkGLContext
// that is currently set by calling gdk_gl_context_clear_current().
/*

C type

GdkGLContext
*/
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

// Object upcasts to *Object
func (recv *GLContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to GLContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a GLContext.
func CastToGLContext(object *gobject.Object) *GLContext {
	return GLContextNewFromC(object.ToC())
}

// A #GdkKeymap defines the translation from keyboard state
// (including a hardware key, a modifier mask, and active keyboard group)
// to a keyval. This translation has two phases. The first phase is
// to determine the effective keyboard group and level for the keyboard
// state; the second phase is to look up the keycode/group/level triplet
// in the keymap and see what keyval it corresponds to.
/*

C type

GdkKeymap
*/
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

// Object upcasts to *Object
func (recv *Keymap) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Keymap.
// Exercise care, as this is a potentially dangerous function if the Object is not a Keymap.
func CastToKeymap(object *gobject.Object) *Keymap {
	return KeymapNewFromC(object.ToC())
}

// Returns the direction of effective layout of the keymap.
/*

C function

gdk_keymap_get_direction
*/
func (recv *Keymap) GetDirection() pango.Direction {
	retC := C.gdk_keymap_get_direction((*C.GdkKeymap)(recv.native))
	retGo := (pango.Direction)(retC)

	return retGo
}

// Unsupported : gdk_keymap_get_entries_for_keycode : unsupported parameter keys : output array param keys

// Unsupported : gdk_keymap_get_entries_for_keyval : unsupported parameter keys : output array param keys

// Looks up the keyval mapped to a keycode/group/level triplet.
// If no keyval is bound to @key, returns 0. For normal user input,
// you want to use gdk_keymap_translate_keyboard_state() instead of
// this function, since the effective group/level may not be
// the same as the current keyboard state.
/*

C function

gdk_keymap_lookup_key
*/
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

// #GdkScreen objects are the GDK representation of the screen on
// which windows can be displayed and on which the pointer moves.
// X originally identified screens with physical screens, but
// nowadays it is more common to have a single #GdkScreen which
// combines several physical monitors (see gdk_screen_get_n_monitors()).
//
// GdkScreen is used throughout GDK and GTK+ to specify which screen
// the top level windows are to be displayed on. it is also used to
// query the screen specification and default settings such as
// the default visual (gdk_screen_get_system_visual()), the dimensions
// of the physical monitors (gdk_screen_get_monitor_geometry()), etc.
/*

C type

GdkScreen
*/
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

// Object upcasts to *Object
func (recv *Screen) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Screen.
// Exercise care, as this is a potentially dangerous function if the Object is not a Screen.
func CastToScreen(object *gobject.Object) *Screen {
	return ScreenNewFromC(object.ToC())
}

// A #GdkVisual contains information about
// a particular visual.
/*

C type

GdkVisual
*/
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

// Object upcasts to *Object
func (recv *Visual) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Visual.
// Exercise care, as this is a potentially dangerous function if the Object is not a Visual.
func CastToVisual(object *gobject.Object) *Visual {
	return VisualNewFromC(object.ToC())
}

/*

C type

GdkWindow
*/
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

// Object upcasts to *Object
func (recv *Window) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Window.
// Exercise care, as this is a potentially dangerous function if the Object is not a Window.
func CastToWindow(object *gobject.Object) *Window {
	return WindowNewFromC(object.ToC())
}

// Creates a new #GdkWindow using the attributes from
// @attributes. See #GdkWindowAttr and #GdkWindowAttributesType for
// more details.  Note: to use this on displays other than the default
// display, @parent must be specified.
/*

C function

gdk_window_new
*/
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

	return retGo
}

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc (GdkFilterFunc) for param function

// Begins a window move operation (for a toplevel window).
//
// This function assumes that the drag is controlled by the
// client pointer device, use gdk_window_begin_move_drag_for_device()
// to begin a drag with a different device.
/*

C function

gdk_window_begin_move_drag
*/
func (recv *Window) BeginMoveDrag(button int32, rootX int32, rootY int32, timestamp uint32) {
	c_button := (C.gint)(button)

	c_root_x := (C.gint)(rootX)

	c_root_y := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_begin_move_drag((*C.GdkWindow)(recv.native), c_button, c_root_x, c_root_y, c_timestamp)

	return
}

// Unsupported : gdk_window_begin_paint_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Indicates that you are beginning the process of redrawing @region.
// A backing store (offscreen buffer) large enough to contain @region
// will be created. The backing store will be initialized with the
// background color or background surface for @window. Then, all
// drawing operations performed on @window will be diverted to the
// backing store.  When you call gdk_window_end_paint(), the backing
// store will be copied to @window, making it visible onscreen. Only
// the part of @window contained in @region will be modified; that is,
// drawing operations are clipped to @region.
//
// The net result of all this is to remove flicker, because the user
// sees the finished product appear all at once when you call
// gdk_window_end_paint(). If you draw to @window directly without
// calling gdk_window_begin_paint_region(), the user may see flicker
// as individual drawing operations are performed in sequence.  The
// clipping and background-initializing features of
// gdk_window_begin_paint_region() are conveniences for the
// programmer, so you can avoid doing that work yourself.
//
// When using GTK+, the widget system automatically places calls to
// gdk_window_begin_paint_region() and gdk_window_end_paint() around
// emissions of the expose_event signal. That is, if you’re writing an
// expose event handler, you can assume that the exposed area in
// #GdkEventExpose has already been cleared to the window background,
// is already set as the clip region, and already has a backing store.
// Therefore in most cases, application code need not call
// gdk_window_begin_paint_region(). (You can disable the automatic
// calls around expose events on a widget-by-widget basis by calling
// gtk_widget_set_double_buffered().)
//
// If you call this function multiple times before calling the
// matching gdk_window_end_paint(), the backing stores are pushed onto
// a stack. gdk_window_end_paint() copies the topmost backing store
// onscreen, subtracts the topmost region from all other regions in
// the stack, and pops the stack. All drawing operations affect only
// the topmost backing store in the stack. One matching call to
// gdk_window_end_paint() is required for each call to
// gdk_window_begin_paint_region().
/*

C function

gdk_window_begin_paint_region
*/
func (recv *Window) BeginPaintRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gdk_window_begin_paint_region((*C.GdkWindow)(recv.native), c_region)

	return
}

// Begins a window resize operation (for a toplevel window).
//
// This function assumes that the drag is controlled by the
// client pointer device, use gdk_window_begin_resize_drag_for_device()
// to begin a drag with a different device.
/*

C function

gdk_window_begin_resize_drag
*/
func (recv *Window) BeginResizeDrag(edge WindowEdge, button int32, rootX int32, rootY int32, timestamp uint32) {
	c_edge := (C.GdkWindowEdge)(edge)

	c_button := (C.gint)(button)

	c_root_x := (C.gint)(rootX)

	c_root_y := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_begin_resize_drag((*C.GdkWindow)(recv.native), c_edge, c_button, c_root_x, c_root_y, c_timestamp)

	return
}

// Attempt to deiconify (unminimize) @window. On X11 the window manager may
// choose to ignore the request to deiconify. When using GTK+,
// use gtk_window_deiconify() instead of the #GdkWindow variant. Or better yet,
// you probably want to use gtk_window_present(), which raises the window, focuses it,
// unminimizes it, and puts it on the current desktop.
/*

C function

gdk_window_deiconify
*/
func (recv *Window) Deiconify() {
	C.gdk_window_deiconify((*C.GdkWindow)(recv.native))

	return
}

// Destroys the window system resources associated with @window and decrements @window's
// reference count. The window system resources for all children of @window are also
// destroyed, but the children’s reference counts are not decremented.
//
// Note that a window will not be destroyed automatically when its reference count
// reaches zero. You must call this function yourself before that happens.
/*

C function

gdk_window_destroy
*/
func (recv *Window) Destroy() {
	C.gdk_window_destroy((*C.GdkWindow)(recv.native))

	return
}

// Blacklisted : gdk_window_destroy_notify

// Indicates that the backing store created by the most recent call
// to gdk_window_begin_paint_region() should be copied onscreen and
// deleted, leaving the next-most-recent backing store or no backing
// store at all as the active paint region. See
// gdk_window_begin_paint_region() for full details.
//
// It is an error to call this function without a matching
// gdk_window_begin_paint_region() first.
/*

C function

gdk_window_end_paint
*/
func (recv *Window) EndPaint() {
	C.gdk_window_end_paint((*C.GdkWindow)(recv.native))

	return
}

// Sets keyboard focus to @window. In most cases, gtk_window_present()
// should be used on a #GtkWindow, rather than calling this function.
/*

C function

gdk_window_focus
*/
func (recv *Window) Focus(timestamp uint32) {
	c_timestamp := (C.guint32)(timestamp)

	C.gdk_window_focus((*C.GdkWindow)(recv.native), c_timestamp)

	return
}

// Temporarily freezes a window and all its descendants such that it won't
// receive expose events.  The window will begin receiving expose events
// again when gdk_window_thaw_toplevel_updates_libgtk_only() is called. If
// gdk_window_freeze_toplevel_updates_libgtk_only()
// has been called more than once,
// gdk_window_thaw_toplevel_updates_libgtk_only() must be called
// an equal number of times to begin processing exposes.
//
// This function is not part of the GDK public API and is only
// for use by GTK+.
/*

C function

gdk_window_freeze_toplevel_updates_libgtk_only
*/
func (recv *Window) FreezeToplevelUpdatesLibgtkOnly() {
	C.gdk_window_freeze_toplevel_updates_libgtk_only((*C.GdkWindow)(recv.native))

	return
}

// Temporarily freezes a window such that it won’t receive expose
// events.  The window will begin receiving expose events again when
// gdk_window_thaw_updates() is called. If gdk_window_freeze_updates()
// has been called more than once, gdk_window_thaw_updates() must be called
// an equal number of times to begin processing exposes.
/*

C function

gdk_window_freeze_updates
*/
func (recv *Window) FreezeUpdates() {
	C.gdk_window_freeze_updates((*C.GdkWindow)(recv.native))

	return
}

// Moves the window into fullscreen mode on the given monitor. This means
// the window covers the entire screen and is above any panels or task bars.
//
// If the window was already fullscreen, then this function does nothing.
/*

C function

gdk_window_fullscreen_on_monitor
*/
func (recv *Window) FullscreenOnMonitor(monitor int32) {
	c_monitor := (C.gint)(monitor)

	C.gdk_window_fullscreen_on_monitor((*C.GdkWindow)(recv.native), c_monitor)

	return
}

// Gets the list of children of @window known to GDK.
// This function only returns children created via GDK,
// so for example it’s useless when used with the root window;
// it only returns windows an application created itself.
//
// The returned list must be freed, but the elements in the
// list need not be.
/*

C function

gdk_window_get_children
*/
func (recv *Window) GetChildren() *glib.List {
	retC := C.gdk_window_get_children((*C.GdkWindow)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Computes the region of a window that potentially can be written
// to by drawing primitives. This region may not take into account
// other factors such as if the window is obscured by other windows,
// but no area outside of this region will be affected by drawing
// primitives.
/*

C function

gdk_window_get_clip_region
*/
func (recv *Window) GetClipRegion() *cairo.Region {
	retC := C.gdk_window_get_clip_region((*C.GdkWindow)(recv.native))
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_get_decorations : unsupported parameter decorations : GdkWMDecoration* with indirection level of 1

// Gets the event mask for @window for all master input devices. See
// gdk_window_set_events().
/*

C function

gdk_window_get_events
*/
func (recv *Window) GetEvents() EventMask {
	retC := C.gdk_window_get_events((*C.GdkWindow)(recv.native))
	retGo := (EventMask)(retC)

	return retGo
}

// Unsupported : gdk_window_get_frame_extents : unsupported parameter rect : Blacklisted record : GdkRectangle

// Any of the return location arguments to this function may be %NULL,
// if you aren’t interested in getting the value of that field.
//
// The X and Y coordinates returned are relative to the parent window
// of @window, which for toplevels usually means relative to the
// window decorations (titlebar, etc.) rather than relative to the
// root window (screen-size background window).
//
// On the X11 platform, the geometry is obtained from the X server,
// so reflects the latest position of @window; this may be out-of-sync
// with the position of @window delivered in the most-recently-processed
// #GdkEventConfigure. gdk_window_get_position() in contrast gets the
// position from the most recent configure event.
//
// Note: If @window is not a toplevel, it is much better
// to call gdk_window_get_position(), gdk_window_get_width() and
// gdk_window_get_height() instead, because it avoids the roundtrip to
// the X server and because these functions support the full 32-bit
// coordinate space, whereas gdk_window_get_geometry() is restricted to
// the 16-bit coordinates of X11.
/*

C function

gdk_window_get_geometry
*/
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

// Obtains the position of a window in root window coordinates.
// (Compare with gdk_window_get_position() and
// gdk_window_get_geometry() which return the position of a window
// relative to its parent window.)
/*

C function

gdk_window_get_origin
*/
func (recv *Window) GetOrigin() (int32, int32, int32) {
	var c_x C.gint

	var c_y C.gint

	retC := C.gdk_window_get_origin((*C.GdkWindow)(recv.native), &c_x, &c_y)
	retGo := (int32)(retC)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return retGo, x, y
}

// Obtains the parent of @window, as known to GDK. Does not query the
// X server; thus this returns the parent as passed to gdk_window_new(),
// not the actual parent. This should never matter unless you’re using
// Xlib calls mixed with GDK calls on the X11 platform. It may also
// matter for toplevel windows, because the window manager may choose
// to reparent them.
//
// Note that you should use gdk_window_get_effective_parent() when
// writing generic code that walks up a window hierarchy, because
// gdk_window_get_parent() will most likely not do what you expect if
// there are offscreen windows in the hierarchy.
/*

C function

gdk_window_get_parent
*/
func (recv *Window) GetParent() *Window {
	retC := C.gdk_window_get_parent((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_get_pointer : unsupported parameter mask : GdkModifierType* with indirection level of 1

// Obtains the position of the window as reported in the
// most-recently-processed #GdkEventConfigure. Contrast with
// gdk_window_get_geometry() which queries the X server for the
// current window position, regardless of which events have been
// received or processed.
//
// The position coordinates are relative to the window’s parent window.
/*

C function

gdk_window_get_position
*/
func (recv *Window) GetPosition() (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	C.gdk_window_get_position((*C.GdkWindow)(recv.native), &c_x, &c_y)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// Obtains the top-left corner of the window manager frame in root
// window coordinates.
/*

C function

gdk_window_get_root_origin
*/
func (recv *Window) GetRootOrigin() (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	C.gdk_window_get_root_origin((*C.GdkWindow)(recv.native), &c_x, &c_y)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// Returns the event mask for @window corresponding to the device class specified
// by @source.
/*

C function

gdk_window_get_source_events
*/
func (recv *Window) GetSourceEvents(source InputSource) EventMask {
	c_source := (C.GdkInputSource)(source)

	retC := C.gdk_window_get_source_events((*C.GdkWindow)(recv.native), c_source)
	retGo := (EventMask)(retC)

	return retGo
}

// Gets the bitwise OR of the currently active window state flags,
// from the #GdkWindowState enumeration.
/*

C function

gdk_window_get_state
*/
func (recv *Window) GetState() WindowState {
	retC := C.gdk_window_get_state((*C.GdkWindow)(recv.native))
	retGo := (WindowState)(retC)

	return retGo
}

// Gets the toplevel window that’s an ancestor of @window.
//
// Any window type but %GDK_WINDOW_CHILD is considered a
// toplevel window, as is a %GDK_WINDOW_CHILD window that
// has a root window as parent.
//
// Note that you should use gdk_window_get_effective_toplevel() when
// you want to get to a window’s toplevel as seen on screen, because
// gdk_window_get_toplevel() will most likely not do what you expect
// if there are offscreen windows in the hierarchy.
/*

C function

gdk_window_get_toplevel
*/
func (recv *Window) GetToplevel() *Window {
	retC := C.gdk_window_get_toplevel((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Transfers ownership of the update area from @window to the caller
// of the function. That is, after calling this function, @window will
// no longer have an invalid/dirty region; the update area is removed
// from @window and handed to you. If a window has no update area,
// gdk_window_get_update_area() returns %NULL. You are responsible for
// calling cairo_region_destroy() on the returned region if it’s non-%NULL.
/*

C function

gdk_window_get_update_area
*/
func (recv *Window) GetUpdateArea() *cairo.Region {
	retC := C.gdk_window_get_update_area((*C.GdkWindow)(recv.native))
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the user data for @window, which is normally the widget
// that @window belongs to. See gdk_window_set_user_data().
/*

C function

gdk_window_get_user_data
*/
func (recv *Window) GetUserData() uintptr {
	var c_data C.gpointer

	C.gdk_window_get_user_data((*C.GdkWindow)(recv.native), &c_data)

	data := (uintptr)(unsafe.Pointer(&c_data))

	return data
}

// Computes the region of the @window that is potentially visible.
// This does not necessarily take into account if the window is
// obscured by other windows, but no area outside of this region
// is visible.
/*

C function

gdk_window_get_visible_region
*/
func (recv *Window) GetVisibleRegion() *cairo.Region {
	retC := C.gdk_window_get_visible_region((*C.GdkWindow)(recv.native))
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the type of the window. See #GdkWindowType.
/*

C function

gdk_window_get_window_type
*/
func (recv *Window) GetWindowType() WindowType {
	retC := C.gdk_window_get_window_type((*C.GdkWindow)(recv.native))
	retGo := (WindowType)(retC)

	return retGo
}

// For toplevel windows, withdraws them, so they will no longer be
// known to the window manager; for all windows, unmaps them, so
// they won’t be displayed. Normally done automatically as
// part of gtk_widget_hide().
/*

C function

gdk_window_hide
*/
func (recv *Window) Hide() {
	C.gdk_window_hide((*C.GdkWindow)(recv.native))

	return
}

// Asks to iconify (minimize) @window. The window manager may choose
// to ignore the request, but normally will honor it. Using
// gtk_window_iconify() is preferred, if you have a #GtkWindow widget.
//
// This function only makes sense when @window is a toplevel window.
/*

C function

gdk_window_iconify
*/
func (recv *Window) Iconify() {
	C.gdk_window_iconify((*C.GdkWindow)(recv.native))

	return
}

// Unsupported : gdk_window_invalidate_maybe_recurse : unsupported parameter child_func : no type generator for WindowChildFunc (GdkWindowChildFunc) for param child_func

// Unsupported : gdk_window_invalidate_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Adds @region to the update area for @window. The update area is the
// region that needs to be redrawn, or “dirty region.” The call
// gdk_window_process_updates() sends one or more expose events to the
// window, which together cover the entire update area. An
// application would normally redraw the contents of @window in
// response to those expose events.
//
// GDK will call gdk_window_process_all_updates() on your behalf
// whenever your program returns to the main loop and becomes idle, so
// normally there’s no need to do that manually, you just need to
// invalidate regions that you know should be redrawn.
//
// The @invalidate_children parameter controls whether the region of
// each child window that intersects @region will also be invalidated.
// If %FALSE, then the update area for child windows will remain
// unaffected. See gdk_window_invalidate_maybe_recurse if you need
// fine grained control over which children are invalidated.
/*

C function

gdk_window_invalidate_region
*/
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

// Check if the window and all ancestors of the window are
// mapped. (This is not necessarily "viewable" in the X sense, since
// we only check as far as we have GDK window parents, not to the root
// window.)
/*

C function

gdk_window_is_viewable
*/
func (recv *Window) IsViewable() bool {
	retC := C.gdk_window_is_viewable((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks whether the window has been mapped (with gdk_window_show() or
// gdk_window_show_unraised()).
/*

C function

gdk_window_is_visible
*/
func (recv *Window) IsVisible() bool {
	retC := C.gdk_window_is_visible((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Lowers @window to the bottom of the Z-order (stacking order), so that
// other windows with the same parent window appear above @window.
// This is true whether or not the other windows are visible.
//
// If @window is a toplevel, the window manager may choose to deny the
// request to move the window in the Z-order, gdk_window_lower() only
// requests the restack, does not guarantee it.
//
// Note that gdk_window_show() raises the window again, so don’t call this
// function before gdk_window_show(). (Try gdk_window_show_unraised().)
/*

C function

gdk_window_lower
*/
func (recv *Window) Lower() {
	C.gdk_window_lower((*C.GdkWindow)(recv.native))

	return
}

// Maximizes the window. If the window was already maximized, then
// this function does nothing.
//
// On X11, asks the window manager to maximize @window, if the window
// manager supports this operation. Not all window managers support
// this, and some deliberately ignore it or don’t have a concept of
// “maximized”; so you can’t rely on the maximization actually
// happening. But it will happen with most standard window managers,
// and GDK makes a best effort to get it to happen.
//
// On Windows, reliably maximizes the window.
/*

C function

gdk_window_maximize
*/
func (recv *Window) Maximize() {
	C.gdk_window_maximize((*C.GdkWindow)(recv.native))

	return
}

// Merges the shape masks for any child windows into the
// shape mask for @window. i.e. the union of all masks
// for @window and its children will become the new mask
// for @window. See gdk_window_shape_combine_region().
//
// This function is distinct from gdk_window_set_child_shapes()
// because it includes @window’s shape mask in the set of shapes to
// be merged.
/*

C function

gdk_window_merge_child_shapes
*/
func (recv *Window) MergeChildShapes() {
	C.gdk_window_merge_child_shapes((*C.GdkWindow)(recv.native))

	return
}

// Repositions a window relative to its parent window.
// For toplevel windows, window managers may ignore or modify the move;
// you should probably use gtk_window_move() on a #GtkWindow widget
// anyway, instead of using GDK functions. For child windows,
// the move will reliably succeed.
//
// If you’re also planning to resize the window, use gdk_window_move_resize()
// to both move and resize simultaneously, for a nicer visual effect.
/*

C function

gdk_window_move
*/
func (recv *Window) Move(x int32, y int32) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_window_move((*C.GdkWindow)(recv.native), c_x, c_y)

	return
}

// Equivalent to calling gdk_window_move() and gdk_window_resize(),
// except that both operations are performed at once, avoiding strange
// visual effects. (i.e. the user may be able to see the window first
// move, then resize, if you don’t use gdk_window_move_resize().)
/*

C function

gdk_window_move_resize
*/
func (recv *Window) MoveResize(x int32, y int32, width int32, height int32) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gdk_window_move_resize((*C.GdkWindow)(recv.native), c_x, c_y, c_width, c_height)

	return
}

// Like gdk_window_get_children(), but does not copy the list of
// children, so the list does not need to be freed.
/*

C function

gdk_window_peek_children
*/
func (recv *Window) PeekChildren() *glib.List {
	retC := C.gdk_window_peek_children((*C.GdkWindow)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sends one or more expose events to @window. The areas in each
// expose event will cover the entire update area for the window (see
// gdk_window_invalidate_region() for details). Normally GDK calls
// gdk_window_process_all_updates() on your behalf, so there’s no
// need to call this function unless you want to force expose events
// to be delivered immediately and synchronously (vs. the usual
// case, where GDK delivers them in an idle handler). Occasionally
// this is useful to produce nicer scrolling behavior, for example.
/*

C function

gdk_window_process_updates
*/
func (recv *Window) ProcessUpdates(updateChildren bool) {
	c_update_children :=
		boolToGboolean(updateChildren)

	C.gdk_window_process_updates((*C.GdkWindow)(recv.native), c_update_children)

	return
}

// Raises @window to the top of the Z-order (stacking order), so that
// other windows with the same parent window appear below @window.
// This is true whether or not the windows are visible.
//
// If @window is a toplevel, the window manager may choose to deny the
// request to move the window in the Z-order, gdk_window_raise() only
// requests the restack, does not guarantee it.
/*

C function

gdk_window_raise
*/
func (recv *Window) Raise() {
	C.gdk_window_raise((*C.GdkWindow)(recv.native))

	return
}

// Registers a window as a potential drop destination.
/*

C function

gdk_window_register_dnd
*/
func (recv *Window) RegisterDnd() {
	C.gdk_window_register_dnd((*C.GdkWindow)(recv.native))

	return
}

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc (GdkFilterFunc) for param function

// Reparents @window into the given @new_parent. The window being
// reparented will be unmapped as a side effect.
/*

C function

gdk_window_reparent
*/
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

// Resizes @window; for toplevel windows, asks the window manager to resize
// the window. The window manager may not allow the resize. When using GTK+,
// use gtk_window_resize() instead of this low-level GDK function.
//
// Windows may not be resized below 1x1.
//
// If you’re also planning to move the window, use gdk_window_move_resize()
// to both move and resize simultaneously, for a nicer visual effect.
/*

C function

gdk_window_resize
*/
func (recv *Window) Resize(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gdk_window_resize((*C.GdkWindow)(recv.native), c_width, c_height)

	return
}

// Scroll the contents of @window, both pixels and children, by the
// given amount. @window itself does not move. Portions of the window
// that the scroll operation brings in from offscreen areas are
// invalidated. The invalidated region may be bigger than what would
// strictly be necessary.
//
// For X11, a minimum area will be invalidated if the window has no
// subwindows, or if the edges of the window’s parent do not extend
// beyond the edges of the window. In other cases, a multi-step process
// is used to scroll the window which may produce temporary visual
// artifacts and unnecessary invalidations.
/*

C function

gdk_window_scroll
*/
func (recv *Window) Scroll(dx int32, dy int32) {
	c_dx := (C.gint)(dx)

	c_dy := (C.gint)(dy)

	C.gdk_window_scroll((*C.GdkWindow)(recv.native), c_dx, c_dy)

	return
}

// Sets the background color of @window.
//
// However, when using GTK+, influence the background of a widget
// using a style class or CSS — if you’re an application — or with
// gtk_style_context_set_background() — if you're implementing a
// custom widget.
/*

C function

gdk_window_set_background
*/
func (recv *Window) SetBackground(color *Color) {
	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	C.gdk_window_set_background((*C.GdkWindow)(recv.native), c_color)

	return
}

// Sets the background of @window.
//
// A background of %NULL means that the window will inherit its
// background from its parent window.
//
// The windowing system will normally fill a window with its background
// when the window is obscured then exposed.
/*

C function

gdk_window_set_background_pattern
*/
func (recv *Window) SetBackgroundPattern(pattern *cairo.Pattern) {
	c_pattern := (*C.cairo_pattern_t)(C.NULL)
	if pattern != nil {
		c_pattern = (*C.cairo_pattern_t)(pattern.ToC())
	}

	C.gdk_window_set_background_pattern((*C.GdkWindow)(recv.native), c_pattern)

	return
}

// Sets the background color of @window.
//
// See also gdk_window_set_background_pattern().
/*

C function

gdk_window_set_background_rgba
*/
func (recv *Window) SetBackgroundRgba(rgba *RGBA) {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gdk_window_set_background_rgba((*C.GdkWindow)(recv.native), c_rgba)

	return
}

// Sets the shape mask of @window to the union of shape masks
// for all children of @window, ignoring the shape mask of @window
// itself. Contrast with gdk_window_merge_child_shapes() which includes
// the shape mask of @window in the masks to be merged.
/*

C function

gdk_window_set_child_shapes
*/
func (recv *Window) SetChildShapes() {
	C.gdk_window_set_child_shapes((*C.GdkWindow)(recv.native))

	return
}

// Sets the default mouse pointer for a #GdkWindow.
//
// Note that @cursor must be for the same display as @window.
//
// Use gdk_cursor_new_for_display() or gdk_cursor_new_from_pixbuf() to
// create the cursor. To make the cursor invisible, use %GDK_BLANK_CURSOR.
// Passing %NULL for the @cursor argument to gdk_window_set_cursor() means
// that @window will use the cursor of its parent window. Most windows
// should use this default.
/*

C function

gdk_window_set_cursor
*/
func (recv *Window) SetCursor(cursor *Cursor) {
	c_cursor := (*C.GdkCursor)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkCursor)(cursor.ToC())
	}

	C.gdk_window_set_cursor((*C.GdkWindow)(recv.native), c_cursor)

	return
}

// “Decorations” are the features the window manager adds to a toplevel #GdkWindow.
// This function sets the traditional Motif window manager hints that tell the
// window manager which decorations you would like your window to have.
// Usually you should use gtk_window_set_decorated() on a #GtkWindow instead of
// using the GDK function directly.
//
// The @decorations argument is the logical OR of the fields in
// the #GdkWMDecoration enumeration. If #GDK_DECOR_ALL is included in the
// mask, the other bits indicate which decorations should be turned off.
// If #GDK_DECOR_ALL is not included, then the other bits indicate
// which decorations should be turned on.
//
// Most window managers honor a decorations hint of 0 to disable all decorations,
// but very few honor all possible combinations of bits.
/*

C function

gdk_window_set_decorations
*/
func (recv *Window) SetDecorations(decorations WMDecoration) {
	c_decorations := (C.GdkWMDecoration)(decorations)

	C.gdk_window_set_decorations((*C.GdkWindow)(recv.native), c_decorations)

	return
}

// The event mask for a window determines which events will be reported
// for that window from all master input devices. For example, an event mask
// including #GDK_BUTTON_PRESS_MASK means the window should report button
// press events. The event mask is the bitwise OR of values from the
// #GdkEventMask enumeration.
//
// See the [input handling overview][event-masks] for details.
/*

C function

gdk_window_set_events
*/
func (recv *Window) SetEvents(eventMask EventMask) {
	c_event_mask := (C.GdkEventMask)(eventMask)

	C.gdk_window_set_events((*C.GdkWindow)(recv.native), c_event_mask)

	return
}

// Sets hints about the window management functions to make available
// via buttons on the window frame.
//
// On the X backend, this function sets the traditional Motif window
// manager hint for this purpose. However, few window managers do
// anything reliable or interesting with this hint. Many ignore it
// entirely.
//
// The @functions argument is the logical OR of values from the
// #GdkWMFunction enumeration. If the bitmask includes #GDK_FUNC_ALL,
// then the other bits indicate which functions to disable; if
// it doesn’t include #GDK_FUNC_ALL, it indicates which functions to
// enable.
/*

C function

gdk_window_set_functions
*/
func (recv *Window) SetFunctions(functions WMFunction) {
	c_functions := (C.GdkWMFunction)(functions)

	C.gdk_window_set_functions((*C.GdkWindow)(recv.native), c_functions)

	return
}

// Sets the geometry hints for @window. Hints flagged in @geom_mask
// are set, hints not flagged in @geom_mask are unset.
// To unset all hints, use a @geom_mask of 0 and a @geometry of %NULL.
//
// This function provides hints to the windowing system about
// acceptable sizes for a toplevel window. The purpose of
// this is to constrain user resizing, but the windowing system
// will typically  (but is not required to) also constrain the
// current size of the window to the provided values and
// constrain programatic resizing via gdk_window_resize() or
// gdk_window_move_resize().
//
// Note that on X11, this effect has no effect on windows
// of type %GDK_WINDOW_TEMP or windows where override redirect
// has been turned on via gdk_window_set_override_redirect()
// since these windows are not resizable by the user.
//
// Since you can’t count on the windowing system doing the
// constraints for programmatic resizes, you should generally
// call gdk_window_constrain_size() yourself to determine
// appropriate sizes.
/*

C function

gdk_window_set_geometry_hints
*/
func (recv *Window) SetGeometryHints(geometry *Geometry, geomMask WindowHints) {
	c_geometry := (*C.GdkGeometry)(C.NULL)
	if geometry != nil {
		c_geometry = (*C.GdkGeometry)(geometry.ToC())
	}

	c_geom_mask := (C.GdkWindowHints)(geomMask)

	C.gdk_window_set_geometry_hints((*C.GdkWindow)(recv.native), c_geometry, c_geom_mask)

	return
}

// Sets the group leader window for @window. By default,
// GDK sets the group leader for all toplevel windows
// to a global window implicitly created by GDK. With this function
// you can override this default.
//
// The group leader window allows the window manager to distinguish
// all windows that belong to a single application. It may for example
// allow users to minimize/unminimize all windows belonging to an
// application at once. You should only set a non-default group window
// if your application pretends to be multiple applications.
/*

C function

gdk_window_set_group
*/
func (recv *Window) SetGroup(leader *Window) {
	c_leader := (*C.GdkWindow)(C.NULL)
	if leader != nil {
		c_leader = (*C.GdkWindow)(leader.ToC())
	}

	C.gdk_window_set_group((*C.GdkWindow)(recv.native), c_leader)

	return
}

// Sets a list of icons for the window. One of these will be used
// to represent the window when it has been iconified. The icon is
// usually shown in an icon box or some sort of task bar. Which icon
// size is shown depends on the window manager. The window manager
// can scale the icon  but setting several size icons can give better
// image quality since the window manager may only need to scale the
// icon by a small amount or not at all.
//
// Note that some platforms don't support window icons.
/*

C function

gdk_window_set_icon_list
*/
func (recv *Window) SetIconList(pixbufs *glib.List) {
	c_pixbufs := (*C.GList)(C.NULL)
	if pixbufs != nil {
		c_pixbufs = (*C.GList)(pixbufs.ToC())
	}

	C.gdk_window_set_icon_list((*C.GdkWindow)(recv.native), c_pixbufs)

	return
}

// Windows may have a name used while minimized, distinct from the
// name they display in their titlebar. Most of the time this is a bad
// idea from a user interface standpoint. But you can set such a name
// with this function, if you like.
//
// After calling this with a non-%NULL @name, calls to gdk_window_set_title()
// will not update the icon title.
//
// Using %NULL for @name unsets the icon title; further calls to
// gdk_window_set_title() will again update the icon title as well.
//
// Note that some platforms don't support window icons.
/*

C function

gdk_window_set_icon_name
*/
func (recv *Window) SetIconName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gdk_window_set_icon_name((*C.GdkWindow)(recv.native), c_name)

	return
}

// The application can use this hint to tell the window manager
// that a certain window has modal behaviour. The window manager
// can use this information to handle modal windows in a special
// way.
//
// You should only use this on windows for which you have
// previously called gdk_window_set_transient_for()
/*

C function

gdk_window_set_modal_hint
*/
func (recv *Window) SetModalHint(modal bool) {
	c_modal :=
		boolToGboolean(modal)

	C.gdk_window_set_modal_hint((*C.GdkWindow)(recv.native), c_modal)

	return
}

// An override redirect window is not under the control of the window manager.
// This means it won’t have a titlebar, won’t be minimizable, etc. - it will
// be entirely under the control of the application. The window manager
// can’t see the override redirect window at all.
//
// Override redirect should only be used for short-lived temporary
// windows, such as popup menus. #GtkMenu uses an override redirect
// window in its implementation, for example.
/*

C function

gdk_window_set_override_redirect
*/
func (recv *Window) SetOverrideRedirect(overrideRedirect bool) {
	c_override_redirect :=
		boolToGboolean(overrideRedirect)

	C.gdk_window_set_override_redirect((*C.GdkWindow)(recv.native), c_override_redirect)

	return
}

// When using GTK+, typically you should use gtk_window_set_role() instead
// of this low-level function.
//
// The window manager and session manager use a window’s role to
// distinguish it from other kinds of window in the same application.
// When an application is restarted after being saved in a previous
// session, all windows with the same title and role are treated as
// interchangeable.  So if you have two windows with the same title
// that should be distinguished for session management purposes, you
// should set the role on those windows. It doesn’t matter what string
// you use for the role, as long as you have a different role for each
// non-interchangeable kind of window.
/*

C function

gdk_window_set_role
*/
func (recv *Window) SetRole(role string) {
	c_role := C.CString(role)
	defer C.free(unsafe.Pointer(c_role))

	C.gdk_window_set_role((*C.GdkWindow)(recv.native), c_role)

	return
}

// Used to set the bit gravity of the given window to static, and flag
// it so all children get static subwindow gravity. This is used if you
// are implementing scary features that involve deep knowledge of the
// windowing system. Don’t worry about it.
/*

C function

gdk_window_set_static_gravities
*/
func (recv *Window) SetStaticGravities(useStatic bool) bool {
	c_use_static :=
		boolToGboolean(useStatic)

	retC := C.gdk_window_set_static_gravities((*C.GdkWindow)(recv.native), c_use_static)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the title of a toplevel window, to be displayed in the titlebar.
// If you haven’t explicitly set the icon name for the window
// (using gdk_window_set_icon_name()), the icon name will be set to
// @title as well. @title must be in UTF-8 encoding (as with all
// user-readable strings in GDK/GTK+). @title may not be %NULL.
/*

C function

gdk_window_set_title
*/
func (recv *Window) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gdk_window_set_title((*C.GdkWindow)(recv.native), c_title)

	return
}

// Indicates to the window manager that @window is a transient dialog
// associated with the application window @parent. This allows the
// window manager to do things like center @window on @parent and
// keep @window above @parent.
//
// See gtk_window_set_transient_for() if you’re using #GtkWindow or
// #GtkDialog.
/*

C function

gdk_window_set_transient_for
*/
func (recv *Window) SetTransientFor(parent *Window) {
	c_parent := (*C.GdkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GdkWindow)(parent.ToC())
	}

	C.gdk_window_set_transient_for((*C.GdkWindow)(recv.native), c_parent)

	return
}

// The application can use this call to provide a hint to the window
// manager about the functionality of a window. The window manager
// can use this information when determining the decoration and behaviour
// of the window.
//
// The hint must be set before the window is mapped.
/*

C function

gdk_window_set_type_hint
*/
func (recv *Window) SetTypeHint(hint WindowTypeHint) {
	c_hint := (C.GdkWindowTypeHint)(hint)

	C.gdk_window_set_type_hint((*C.GdkWindow)(recv.native), c_hint)

	return
}

// For most purposes this function is deprecated in favor of
// g_object_set_data(). However, for historical reasons GTK+ stores
// the #GtkWidget that owns a #GdkWindow as user data on the
// #GdkWindow. So, custom widget implementations should use
// this function for that. If GTK+ receives an event for a #GdkWindow,
// and the user data for the window is non-%NULL, GTK+ will assume the
// user data is a #GtkWidget, and forward the event to that widget.
/*

C function

gdk_window_set_user_data
*/
func (recv *Window) SetUserData(userData uintptr) {
	c_user_data := (C.gpointer)(userData)

	C.gdk_window_set_user_data((*C.GdkWindow)(recv.native), c_user_data)

	return
}

// Makes pixels in @window outside @shape_region be transparent,
// so that the window may be nonrectangular.
//
// If @shape_region is %NULL, the shape will be unset, so the whole
// window will be opaque again. @offset_x and @offset_y are ignored
// if @shape_region is %NULL.
//
// On the X11 platform, this uses an X server extension which is
// widely available on most common platforms, but not available on
// very old X servers, and occasionally the implementation will be
// buggy. On servers without the shape extension, this function
// will do nothing.
//
// This function works on both toplevel and child windows.
/*

C function

gdk_window_shape_combine_region
*/
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

// Like gdk_window_show_unraised(), but also raises the window to the
// top of the window stack (moves the window to the front of the
// Z-order).
//
// This function maps a window so it’s visible onscreen. Its opposite
// is gdk_window_hide().
//
// When implementing a #GtkWidget, you should call this function on the widget's
// #GdkWindow as part of the “map” method.
/*

C function

gdk_window_show
*/
func (recv *Window) Show() {
	C.gdk_window_show((*C.GdkWindow)(recv.native))

	return
}

// Shows a #GdkWindow onscreen, but does not modify its stacking
// order. In contrast, gdk_window_show() will raise the window
// to the top of the window stack.
//
// On the X11 platform, in Xlib terms, this function calls
// XMapWindow() (it also updates some internal GDK state, which means
// that you can’t really use XMapWindow() directly on a GDK window).
/*

C function

gdk_window_show_unraised
*/
func (recv *Window) ShowUnraised() {
	C.gdk_window_show_unraised((*C.GdkWindow)(recv.native))

	return
}

// “Pins” a window such that it’s on all workspaces and does not scroll
// with viewports, for window managers that have scrollable viewports.
// (When using #GtkWindow, gtk_window_stick() may be more useful.)
//
// On the X11 platform, this function depends on window manager
// support, so may have no effect with many window managers. However,
// GDK will do the best it can to convince the window manager to stick
// the window. For window managers that don’t support this operation,
// there’s nothing you can do to force it to happen.
/*

C function

gdk_window_stick
*/
func (recv *Window) Stick() {
	C.gdk_window_stick((*C.GdkWindow)(recv.native))

	return
}

// Thaws a window frozen with
// gdk_window_freeze_toplevel_updates_libgtk_only().
//
// This function is not part of the GDK public API and is only
// for use by GTK+.
/*

C function

gdk_window_thaw_toplevel_updates_libgtk_only
*/
func (recv *Window) ThawToplevelUpdatesLibgtkOnly() {
	C.gdk_window_thaw_toplevel_updates_libgtk_only((*C.GdkWindow)(recv.native))

	return
}

// Thaws a window frozen with gdk_window_freeze_updates().
/*

C function

gdk_window_thaw_updates
*/
func (recv *Window) ThawUpdates() {
	C.gdk_window_thaw_updates((*C.GdkWindow)(recv.native))

	return
}

// Unmaximizes the window. If the window wasn’t maximized, then this
// function does nothing.
//
// On X11, asks the window manager to unmaximize @window, if the
// window manager supports this operation. Not all window managers
// support this, and some deliberately ignore it or don’t have a
// concept of “maximized”; so you can’t rely on the unmaximization
// actually happening. But it will happen with most standard window
// managers, and GDK makes a best effort to get it to happen.
//
// On Windows, reliably unmaximizes the window.
/*

C function

gdk_window_unmaximize
*/
func (recv *Window) Unmaximize() {
	C.gdk_window_unmaximize((*C.GdkWindow)(recv.native))

	return
}

// Reverse operation for gdk_window_stick(); see gdk_window_stick(),
// and gtk_window_unstick().
/*

C function

gdk_window_unstick
*/
func (recv *Window) Unstick() {
	C.gdk_window_unstick((*C.GdkWindow)(recv.native))

	return
}

// Withdraws a window (unmaps it and asks the window manager to forget about it).
// This function is not really useful as gdk_window_hide() automatically
// withdraws toplevel windows before hiding them.
/*

C function

gdk_window_withdraw
*/
func (recv *Window) Withdraw() {
	C.gdk_window_withdraw((*C.GdkWindow)(recv.native))

	return
}
