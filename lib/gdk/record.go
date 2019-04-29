// This is a generated file - DO NOT EDIT

package gdk

import "unsafe"

// Atom is a wrapper around the C record GdkAtom.
type Atom struct {
	native unsafe.Pointer
}

// Color is a wrapper around the C record GdkColor.
type Color struct {
	native unsafe.Pointer
	Pixel  uint32
	Red    uint16
	Green  uint16
	Blue   uint16
}

// EventAny is a wrapper around the C record GdkEventAny.
type EventAny struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
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

// EventFocus is a wrapper around the C record GdkEventFocus.
type EventFocus struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	In        int16
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

// EventProximity is a wrapper around the C record GdkEventProximity.
type EventProximity struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	// device : record
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

// EventSequence is a wrapper around the C record GdkEventSequence.
type EventSequence struct {
	native unsafe.Pointer
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

// EventVisibility is a wrapper around the C record GdkEventVisibility.
type EventVisibility struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	State     VisibilityState
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

// FrameClockClass is a wrapper around the C record GdkFrameClockClass.
type FrameClockClass struct {
	native unsafe.Pointer
}

// FrameClockPrivate is a wrapper around the C record GdkFrameClockPrivate.
type FrameClockPrivate struct {
	native unsafe.Pointer
}

// FrameTimings is a wrapper around the C record GdkFrameTimings.
type FrameTimings struct {
	native unsafe.Pointer
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

// KeymapKey is a wrapper around the C record GdkKeymapKey.
type KeymapKey struct {
	native  unsafe.Pointer
	Keycode uint32
	Group   int32
	Level   int32
}

// Point is a wrapper around the C record GdkPoint.
type Point struct {
	native unsafe.Pointer
	X      int32
	Y      int32
}

// RGBA is a wrapper around the C record GdkRGBA.
type RGBA struct {
	native unsafe.Pointer
	Red    float64
	Green  float64
	Blue   float64
	Alpha  float64
}

// Rectangle is a wrapper around the C record GdkRectangle.
type Rectangle struct {
	native unsafe.Pointer
	X      int32
	Y      int32
	Width  int32
	Height int32
}

// TimeCoord is a wrapper around the C record GdkTimeCoord.
type TimeCoord struct {
	native unsafe.Pointer
	Time   uint32
	// no type for axes
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

// WindowRedirect is a wrapper around the C record GdkWindowRedirect.
type WindowRedirect struct {
	native unsafe.Pointer
}
