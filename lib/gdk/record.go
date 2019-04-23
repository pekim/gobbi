// This is a generated file - DO NOT EDIT

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GdkAtom

// Blacklisted : GdkColor

// Blacklisted : GdkEventAny

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

// Blacklisted : GdkEventDND

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

// Blacklisted : GdkEventFocus

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

// Blacklisted : GdkEventMotion

// Blacklisted : GdkEventProperty

// Blacklisted : GdkEventProximity

// Blacklisted : GdkEventScroll

// Blacklisted : GdkEventSelection

// Blacklisted : GdkEventSequence

// Blacklisted : GdkEventSetting

// Blacklisted : GdkEventTouch

// Blacklisted : GdkEventTouchpadPinch

// Blacklisted : GdkEventTouchpadSwipe

// Blacklisted : GdkEventVisibility

// Blacklisted : GdkEventWindowState

// Blacklisted : GdkFrameClockClass

// Blacklisted : GdkFrameClockPrivate

// Blacklisted : GdkFrameTimings

// Blacklisted : GdkGeometry

// Blacklisted : GdkKeymapKey

// Blacklisted : GdkPoint

// Blacklisted : GdkRGBA

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

// Blacklisted : gdk_rectangle_intersect

// Blacklisted : gdk_rectangle_union

// Blacklisted : GdkTimeCoord

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

// Blacklisted : GdkWindowClass

// Blacklisted : GdkWindowRedirect
