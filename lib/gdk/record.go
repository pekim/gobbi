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

// Blacklisted : GdkEventSequence

// Blacklisted : GdkEventSetting

// Blacklisted : GdkEventTouch

// Blacklisted : GdkEventTouchpadPinch

// Blacklisted : GdkEventTouchpadSwipe

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
