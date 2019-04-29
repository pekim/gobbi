// This is a generated file - DO NOT EDIT

package gdk

import "unsafe"

// Atom is a wrapper around the C record GdkAtom.
type Atom struct {
	native unsafe.Pointer
}

func AtomNewFromC(u unsafe.Pointer) *Atom {
	if u == nil {
		return nil
	}

	g := &Atom{native: u}

	return g
}

func (recv *Atom) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Color is a wrapper around the C record GdkColor.
type Color struct {
	native unsafe.Pointer
	Pixel  uint32
	Red    uint16
	Green  uint16
	Blue   uint16
}

func ColorNewFromC(u unsafe.Pointer) *Color {
	if u == nil {
		return nil
	}

	g := &Color{native: u}

	return g
}

func (recv *Color) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventAny is a wrapper around the C record GdkEventAny.
type EventAny struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
}

func EventAnyNewFromC(u unsafe.Pointer) *EventAny {
	if u == nil {
		return nil
	}

	g := &EventAny{native: u}

	return g
}

func (recv *EventAny) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventButton is a wrapper around the C record GdkEventButton.
type EventButton struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventButton{native: u}

	return g
}

func (recv *EventButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventConfigure is a wrapper around the C record GdkEventConfigure.
type EventConfigure struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	X         int32
	Y         int32
	Width     int32
	Height    int32
}

func EventConfigureNewFromC(u unsafe.Pointer) *EventConfigure {
	if u == nil {
		return nil
	}

	g := &EventConfigure{native: u}

	return g
}

func (recv *EventConfigure) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventCrossing is a wrapper around the C record GdkEventCrossing.
type EventCrossing struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventCrossing{native: u}

	return g
}

func (recv *EventCrossing) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventDND is a wrapper around the C record GdkEventDND.
type EventDND struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	// context : record
	Time  uint32
	XRoot int16
	YRoot int16
}

func EventDNDNewFromC(u unsafe.Pointer) *EventDND {
	if u == nil {
		return nil
	}

	g := &EventDND{native: u}

	return g
}

func (recv *EventDND) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventExpose is a wrapper around the C record GdkEventExpose.
type EventExpose struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	// area : record
	// region : record
	Count int32
}

func EventExposeNewFromC(u unsafe.Pointer) *EventExpose {
	if u == nil {
		return nil
	}

	g := &EventExpose{native: u}

	return g
}

func (recv *EventExpose) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventFocus is a wrapper around the C record GdkEventFocus.
type EventFocus struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	In        int16
}

func EventFocusNewFromC(u unsafe.Pointer) *EventFocus {
	if u == nil {
		return nil
	}

	g := &EventFocus{native: u}

	return g
}

func (recv *EventFocus) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventKey is a wrapper around the C record GdkEventKey.
type EventKey struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventKey{native: u}

	return g
}

func (recv *EventKey) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventMotion is a wrapper around the C record GdkEventMotion.
type EventMotion struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventMotion{native: u}

	return g
}

func (recv *EventMotion) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventProperty is a wrapper around the C record GdkEventProperty.
type EventProperty struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	// atom : record
	Time  uint32
	State PropertyState
}

func EventPropertyNewFromC(u unsafe.Pointer) *EventProperty {
	if u == nil {
		return nil
	}

	g := &EventProperty{native: u}

	return g
}

func (recv *EventProperty) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventProximity is a wrapper around the C record GdkEventProximity.
type EventProximity struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	// device : record
}

func EventProximityNewFromC(u unsafe.Pointer) *EventProximity {
	if u == nil {
		return nil
	}

	g := &EventProximity{native: u}

	return g
}

func (recv *EventProximity) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventScroll is a wrapper around the C record GdkEventScroll.
type EventScroll struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventScroll{native: u}

	return g
}

func (recv *EventScroll) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventSelection is a wrapper around the C record GdkEventSelection.
type EventSelection struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventSelection{native: u}

	return g
}

func (recv *EventSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventSequence is a wrapper around the C record GdkEventSequence.
type EventSequence struct {
	native unsafe.Pointer
}

func EventSequenceNewFromC(u unsafe.Pointer) *EventSequence {
	if u == nil {
		return nil
	}

	g := &EventSequence{native: u}

	return g
}

func (recv *EventSequence) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventSetting is a wrapper around the C record GdkEventSetting.
type EventSetting struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	Action    SettingAction
	Name      string
}

func EventSettingNewFromC(u unsafe.Pointer) *EventSetting {
	if u == nil {
		return nil
	}

	g := &EventSetting{native: u}

	return g
}

func (recv *EventSetting) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventTouch is a wrapper around the C record GdkEventTouch.
type EventTouch struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventTouch{native: u}

	return g
}

func (recv *EventTouch) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventTouchpadPinch is a wrapper around the C record GdkEventTouchpadPinch.
type EventTouchpadPinch struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventTouchpadPinch{native: u}

	return g
}

func (recv *EventTouchpadPinch) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventTouchpadSwipe is a wrapper around the C record GdkEventTouchpadSwipe.
type EventTouchpadSwipe struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventTouchpadSwipe{native: u}

	return g
}

func (recv *EventTouchpadSwipe) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventVisibility is a wrapper around the C record GdkEventVisibility.
type EventVisibility struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	State     VisibilityState
}

func EventVisibilityNewFromC(u unsafe.Pointer) *EventVisibility {
	if u == nil {
		return nil
	}

	g := &EventVisibility{native: u}

	return g
}

func (recv *EventVisibility) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventWindowState is a wrapper around the C record GdkEventWindowState.
type EventWindowState struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent      int8
	ChangedMask    WindowState
	NewWindowState WindowState
}

func EventWindowStateNewFromC(u unsafe.Pointer) *EventWindowState {
	if u == nil {
		return nil
	}

	g := &EventWindowState{native: u}

	return g
}

func (recv *EventWindowState) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FrameClockClass is a wrapper around the C record GdkFrameClockClass.
type FrameClockClass struct {
	native unsafe.Pointer
}

func FrameClockClassNewFromC(u unsafe.Pointer) *FrameClockClass {
	if u == nil {
		return nil
	}

	g := &FrameClockClass{native: u}

	return g
}

func (recv *FrameClockClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FrameClockPrivate is a wrapper around the C record GdkFrameClockPrivate.
type FrameClockPrivate struct {
	native unsafe.Pointer
}

func FrameClockPrivateNewFromC(u unsafe.Pointer) *FrameClockPrivate {
	if u == nil {
		return nil
	}

	g := &FrameClockPrivate{native: u}

	return g
}

func (recv *FrameClockPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FrameTimings is a wrapper around the C record GdkFrameTimings.
type FrameTimings struct {
	native unsafe.Pointer
}

func FrameTimingsNewFromC(u unsafe.Pointer) *FrameTimings {
	if u == nil {
		return nil
	}

	g := &FrameTimings{native: u}

	return g
}

func (recv *FrameTimings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Geometry is a wrapper around the C record GdkGeometry.
type Geometry struct {
	native     unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &Geometry{native: u}

	return g
}

func (recv *Geometry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// KeymapKey is a wrapper around the C record GdkKeymapKey.
type KeymapKey struct {
	native  unsafe.Pointer
	Keycode uint32
	Group   int32
	Level   int32
}

func KeymapKeyNewFromC(u unsafe.Pointer) *KeymapKey {
	if u == nil {
		return nil
	}

	g := &KeymapKey{native: u}

	return g
}

func (recv *KeymapKey) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Point is a wrapper around the C record GdkPoint.
type Point struct {
	native unsafe.Pointer
	X      int32
	Y      int32
}

func PointNewFromC(u unsafe.Pointer) *Point {
	if u == nil {
		return nil
	}

	g := &Point{native: u}

	return g
}

func (recv *Point) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RGBA is a wrapper around the C record GdkRGBA.
type RGBA struct {
	native unsafe.Pointer
	Red    float64
	Green  float64
	Blue   float64
	Alpha  float64
}

func RGBANewFromC(u unsafe.Pointer) *RGBA {
	if u == nil {
		return nil
	}

	g := &RGBA{native: u}

	return g
}

func (recv *RGBA) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Rectangle is a wrapper around the C record GdkRectangle.
type Rectangle struct {
	native unsafe.Pointer
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func RectangleNewFromC(u unsafe.Pointer) *Rectangle {
	if u == nil {
		return nil
	}

	g := &Rectangle{native: u}

	return g
}

func (recv *Rectangle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TimeCoord is a wrapper around the C record GdkTimeCoord.
type TimeCoord struct {
	native unsafe.Pointer
	Time   uint32
	// no type for axes
}

func TimeCoordNewFromC(u unsafe.Pointer) *TimeCoord {
	if u == nil {
		return nil
	}

	g := &TimeCoord{native: u}

	return g
}

func (recv *TimeCoord) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowAttr is a wrapper around the C record GdkWindowAttr.
type WindowAttr struct {
	native    unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &WindowAttr{native: u}

	return g
}

func (recv *WindowAttr) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowClass is a wrapper around the C record GdkWindowClass.
type WindowClass struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &WindowClass{native: u}

	return g
}

func (recv *WindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowRedirect is a wrapper around the C record GdkWindowRedirect.
type WindowRedirect struct {
	native unsafe.Pointer
}

func WindowRedirectNewFromC(u unsafe.Pointer) *WindowRedirect {
	if u == nil {
		return nil
	}

	g := &WindowRedirect{native: u}

	return g
}

func (recv *WindowRedirect) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
