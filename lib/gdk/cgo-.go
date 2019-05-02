// This is a generated file - DO NOT EDIT

package gdk

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : alias has no type generator for none, void

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

// Blacklisted : gdk_atom_intern

// Blacklisted : gdk_cairo_get_clip_rectangle

// Blacklisted : gdk_color_parse

// Unsupported : gdk_event_get : no return generator

// Unsupported : gdk_event_handler_set : unsupported parameter func : no type generator for EventFunc (GdkEventFunc) for param func

// Unsupported : gdk_event_peek : no return generator

// Unsupported : gdk_property_get : unsupported parameter actual_length : array length param actual_length is pointer (gint*)

// Unsupported : gdk_query_depths : unsupported parameter depths : output array param depths

// Unsupported : gdk_query_visual_types : unsupported parameter visual_types : output array param visual_types

// Unsupported : gdk_selection_property_get : unsupported parameter data : guchar** with indirection level of 2

// Blacklisted : gdk_synthesize_window_state

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
