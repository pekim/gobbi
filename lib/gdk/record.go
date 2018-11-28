// This is a generated file - DO NOT EDIT

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GdkAtom

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

// Equals compares this Color with another Color, and returns true if they represent the same GObject.
func (recv *Color) Equals(other *Color) bool {
	return other.ToC() == recv.ToC()
}

// Copy is a wrapper around the C function gdk_color_copy.
func (recv *Color) Copy() *Color {
	retC := C.gdk_color_copy((*C.GdkColor)(recv.native))
	retGo := ColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function gdk_color_equal.
func (recv *Color) Equal(colorb *Color) bool {
	c_colorb := (*C.GdkColor)(C.NULL)
	if colorb != nil {
		c_colorb = (*C.GdkColor)(colorb.ToC())
	}

	retC := C.gdk_color_equal((*C.GdkColor)(recv.native), c_colorb)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function gdk_color_free.
func (recv *Color) Free() {
	C.gdk_color_free((*C.GdkColor)(recv.native))

	return
}

// Hash is a wrapper around the C function gdk_color_hash.
func (recv *Color) Hash() uint32 {
	retC := C.gdk_color_hash((*C.GdkColor)(recv.native))
	retGo := (uint32)(retC)

	return retGo
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

// Equals compares this EventDND with another EventDND, and returns true if they represent the same GObject.
func (recv *EventDND) Equals(other *EventDND) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this EventSequence with another EventSequence, and returns true if they represent the same GObject.
func (recv *EventSequence) Equals(other *EventSequence) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this EventSetting with another EventSetting, and returns true if they represent the same GObject.
func (recv *EventSetting) Equals(other *EventSetting) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this EventTouch with another EventTouch, and returns true if they represent the same GObject.
func (recv *EventTouch) Equals(other *EventTouch) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this EventTouchpadPinch with another EventTouchpadPinch, and returns true if they represent the same GObject.
func (recv *EventTouchpadPinch) Equals(other *EventTouchpadPinch) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this EventTouchpadSwipe with another EventTouchpadSwipe, and returns true if they represent the same GObject.
func (recv *EventTouchpadSwipe) Equals(other *EventTouchpadSwipe) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this FrameClockClass with another FrameClockClass, and returns true if they represent the same GObject.
func (recv *FrameClockClass) Equals(other *FrameClockClass) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this FrameClockPrivate with another FrameClockPrivate, and returns true if they represent the same GObject.
func (recv *FrameClockPrivate) Equals(other *FrameClockPrivate) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this FrameTimings with another FrameTimings, and returns true if they represent the same GObject.
func (recv *FrameTimings) Equals(other *FrameTimings) bool {
	return other.ToC() == recv.ToC()
}

// GetFrameTime is a wrapper around the C function gdk_frame_timings_get_frame_time.
func (recv *FrameTimings) GetFrameTime() int64 {
	retC := C.gdk_frame_timings_get_frame_time((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
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

// Equals compares this Geometry with another Geometry, and returns true if they represent the same GObject.
func (recv *Geometry) Equals(other *Geometry) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this KeymapKey with another KeymapKey, and returns true if they represent the same GObject.
func (recv *KeymapKey) Equals(other *KeymapKey) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this Point with another Point, and returns true if they represent the same GObject.
func (recv *Point) Equals(other *Point) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this RGBA with another RGBA, and returns true if they represent the same GObject.
func (recv *RGBA) Equals(other *RGBA) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : GdkRectangle

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

// Equals compares this TimeCoord with another TimeCoord, and returns true if they represent the same GObject.
func (recv *TimeCoord) Equals(other *TimeCoord) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this WindowAttr with another WindowAttr, and returns true if they represent the same GObject.
func (recv *WindowAttr) Equals(other *WindowAttr) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this WindowClass with another WindowClass, and returns true if they represent the same GObject.
func (recv *WindowClass) Equals(other *WindowClass) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this WindowRedirect with another WindowRedirect, and returns true if they represent the same GObject.
func (recv *WindowRedirect) Equals(other *WindowRedirect) bool {
	return other.ToC() == recv.ToC()
}
