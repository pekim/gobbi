// Code generated - DO NOT EDIT.

package gdk

import gi "github.com/pekim/gobbi/internal/gi"

type Atom struct {
	native uintptr
}

var nameAtomInvoker *gi.Function

// Name is a representation of the C type gdk_atom_name.
func (recv *Atom) Name() string {
	if nameAtomInvoker == nil {
		nameAtomInvoker = gi.StructFunctionInvokerNew("Gdk", "Atom", "name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nameAtomInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

type Color struct {
	native uintptr
	Pixel  uint32
	Red    uint16
	Green  uint16
	Blue   uint16
}

var copyColorInvoker *gi.Function

// Copy is a representation of the C type gdk_color_copy.
func (recv *Color) Copy() *Color {
	if copyColorInvoker == nil {
		copyColorInvoker = gi.StructFunctionInvokerNew("Gdk", "Color", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyColorInvoker.Invoke(inArgs[:], nil)

	retGo := &Color{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gdk_color_equal' : parameter 'colorb' of type 'Color' not supported

var freeColorInvoker *gi.Function

// Free is a representation of the C type gdk_color_free.
func (recv *Color) Free() {
	if freeColorInvoker == nil {
		freeColorInvoker = gi.StructFunctionInvokerNew("Gdk", "Color", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeColorInvoker.Invoke(inArgs[:], nil)

}

var hashColorInvoker *gi.Function

// Hash is a representation of the C type gdk_color_hash.
func (recv *Color) Hash() uint32 {
	if hashColorInvoker == nil {
		hashColorInvoker = gi.StructFunctionInvokerNew("Gdk", "Color", "hash")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hashColorInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var toStringColorInvoker *gi.Function

// ToString is a representation of the C type gdk_color_to_string.
func (recv *Color) ToString() string {
	if toStringColorInvoker == nil {
		toStringColorInvoker = gi.StructFunctionInvokerNew("Gdk", "Color", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringColorInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

type DevicePadInterface struct {
	native uintptr
}

type DrawingContextClass struct {
	native uintptr
}

type EventAny struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
}

type EventButton struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'axes' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'
	Button uint32
	// UNSUPPORTED : C value 'device' : no Go type for 'Device'
	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'
}

type EventConfigure struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	X         int32
	Y         int32
	Width     int32
	Height    int32
}

type EventCrossing struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	// UNSUPPORTED : C value 'subwindow' : no Go type for 'Window'
	Time uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'mode' : no Go type for 'CrossingMode'
	// UNSUPPORTED : C value 'detail' : no Go type for 'NotifyType'
	// UNSUPPORTED : C value 'focus' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'
}

type EventDND struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	// UNSUPPORTED : C value 'context' : no Go type for 'DragContext'
	Time  uint32
	XRoot int16
	YRoot int16
}

type EventExpose struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Area      *Rectangle
	// UNSUPPORTED : C value 'region' : no Go type for 'cairo.Region'
	Count int32
}

type EventFocus struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	In        int16
}

type EventGrabBroken struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	// UNSUPPORTED : C value 'keyboard' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'implicit' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'grab_window' : no Go type for 'Window'
}

type EventKey struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'
	Keyval          uint32
	Length          int32
	String          string
	HardwareKeycode uint16
	Group           uint8
	IsModifier      uint32
}

type EventMotion struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'axes' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'
	IsHint int16
	// UNSUPPORTED : C value 'device' : no Go type for 'Device'
	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'
}

type EventOwnerChange struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	// UNSUPPORTED : C value 'owner' : no Go type for 'Window'
	// UNSUPPORTED : C value 'reason' : no Go type for 'OwnerChange'
	Selection     *Atom
	Time          uint32
	SelectionTime uint32
}

type EventPadAxis struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Time      uint32
	Group     uint32
	Index     uint32
	Mode      uint32
	// UNSUPPORTED : C value 'value' : no Go type for 'gdouble'
}

type EventPadButton struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Time      uint32
	Group     uint32
	Button    uint32
	Mode      uint32
}

type EventPadGroupMode struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Time      uint32
	Group     uint32
	Mode      uint32
}

type EventProperty struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Atom      *Atom
	Time      uint32
	// UNSUPPORTED : C value 'state' : no Go type for 'PropertyState'
}

type EventProximity struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'device' : no Go type for 'Device'
}

type EventScroll struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'
	// UNSUPPORTED : C value 'direction' : no Go type for 'ScrollDirection'
	// UNSUPPORTED : C value 'device' : no Go type for 'Device'
	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'delta_x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'delta_y' : no Go type for 'gdouble'
	IsStop uint32
}

type EventSelection struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Selection *Atom
	Target    *Atom
	Property  *Atom
	Time      uint32
	// UNSUPPORTED : C value 'requestor' : no Go type for 'Window'
}

type EventSequence struct {
	native uintptr
}

type EventSetting struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	// UNSUPPORTED : C value 'action' : no Go type for 'SettingAction'
	Name string
}

type EventTouch struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'axes' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'
	Sequence *EventSequence
	// UNSUPPORTED : C value 'emulating_pointer' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'device' : no Go type for 'Device'
	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'
}

type EventTouchpadPinch struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Phase     int8
	NFingers  int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'dx' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'dy' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'angle_delta' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'scale' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'
}

type EventTouchpadSwipe struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Phase     int8
	NFingers  int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'dx' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'dy' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'
}

type EventVisibility struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	// UNSUPPORTED : C value 'state' : no Go type for 'VisibilityState'
}

type EventWindowState struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	// UNSUPPORTED : C value 'changed_mask' : no Go type for 'WindowState'
	// UNSUPPORTED : C value 'new_window_state' : no Go type for 'WindowState'
}

type FrameClockClass struct {
	native uintptr
}

type FrameClockPrivate struct {
	native uintptr
}

type FrameTimings struct {
	native uintptr
}

// UNSUPPORTED : C value 'gdk_frame_timings_get_complete' : return type 'gboolean' not supported

var getFrameCounterFrameTimingsInvoker *gi.Function

// GetFrameCounter is a representation of the C type gdk_frame_timings_get_frame_counter.
func (recv *FrameTimings) GetFrameCounter() int64 {
	if getFrameCounterFrameTimingsInvoker == nil {
		getFrameCounterFrameTimingsInvoker = gi.StructFunctionInvokerNew("Gdk", "FrameTimings", "get_frame_counter")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getFrameCounterFrameTimingsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var getFrameTimeFrameTimingsInvoker *gi.Function

// GetFrameTime is a representation of the C type gdk_frame_timings_get_frame_time.
func (recv *FrameTimings) GetFrameTime() int64 {
	if getFrameTimeFrameTimingsInvoker == nil {
		getFrameTimeFrameTimingsInvoker = gi.StructFunctionInvokerNew("Gdk", "FrameTimings", "get_frame_time")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getFrameTimeFrameTimingsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var getPredictedPresentationTimeFrameTimingsInvoker *gi.Function

// GetPredictedPresentationTime is a representation of the C type gdk_frame_timings_get_predicted_presentation_time.
func (recv *FrameTimings) GetPredictedPresentationTime() int64 {
	if getPredictedPresentationTimeFrameTimingsInvoker == nil {
		getPredictedPresentationTimeFrameTimingsInvoker = gi.StructFunctionInvokerNew("Gdk", "FrameTimings", "get_predicted_presentation_time")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPredictedPresentationTimeFrameTimingsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var getPresentationTimeFrameTimingsInvoker *gi.Function

// GetPresentationTime is a representation of the C type gdk_frame_timings_get_presentation_time.
func (recv *FrameTimings) GetPresentationTime() int64 {
	if getPresentationTimeFrameTimingsInvoker == nil {
		getPresentationTimeFrameTimingsInvoker = gi.StructFunctionInvokerNew("Gdk", "FrameTimings", "get_presentation_time")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPresentationTimeFrameTimingsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var getRefreshIntervalFrameTimingsInvoker *gi.Function

// GetRefreshInterval is a representation of the C type gdk_frame_timings_get_refresh_interval.
func (recv *FrameTimings) GetRefreshInterval() int64 {
	if getRefreshIntervalFrameTimingsInvoker == nil {
		getRefreshIntervalFrameTimingsInvoker = gi.StructFunctionInvokerNew("Gdk", "FrameTimings", "get_refresh_interval")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getRefreshIntervalFrameTimingsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var refFrameTimingsInvoker *gi.Function

// Ref is a representation of the C type gdk_frame_timings_ref.
func (recv *FrameTimings) Ref() *FrameTimings {
	if refFrameTimingsInvoker == nil {
		refFrameTimingsInvoker = gi.StructFunctionInvokerNew("Gdk", "FrameTimings", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refFrameTimingsInvoker.Invoke(inArgs[:], nil)

	retGo := &FrameTimings{native: ret.Pointer()}

	return retGo
}

var unrefFrameTimingsInvoker *gi.Function

// Unref is a representation of the C type gdk_frame_timings_unref.
func (recv *FrameTimings) Unref() {
	if unrefFrameTimingsInvoker == nil {
		unrefFrameTimingsInvoker = gi.StructFunctionInvokerNew("Gdk", "FrameTimings", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefFrameTimingsInvoker.Invoke(inArgs[:], nil)

}

type Geometry struct {
	native     uintptr
	MinWidth   int32
	MinHeight  int32
	MaxWidth   int32
	MaxHeight  int32
	BaseWidth  int32
	BaseHeight int32
	WidthInc   int32
	HeightInc  int32
	// UNSUPPORTED : C value 'min_aspect' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'max_aspect' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'win_gravity' : no Go type for 'Gravity'
}

type KeymapKey struct {
	native  uintptr
	Keycode uint32
	Group   int32
	Level   int32
}

type MonitorClass struct {
	native uintptr
}

type Point struct {
	native uintptr
	X      int32
	Y      int32
}

type RGBA struct {
	native uintptr
	// UNSUPPORTED : C value 'red' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'green' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'blue' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'alpha' : no Go type for 'gdouble'
}

var copyRGBAInvoker *gi.Function

// Copy is a representation of the C type gdk_rgba_copy.
func (recv *RGBA) Copy() *RGBA {
	if copyRGBAInvoker == nil {
		copyRGBAInvoker = gi.StructFunctionInvokerNew("Gdk", "RGBA", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyRGBAInvoker.Invoke(inArgs[:], nil)

	retGo := &RGBA{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gdk_rgba_equal' : parameter 'p2' of type 'RGBA' not supported

var freeRGBAInvoker *gi.Function

// Free is a representation of the C type gdk_rgba_free.
func (recv *RGBA) Free() {
	if freeRGBAInvoker == nil {
		freeRGBAInvoker = gi.StructFunctionInvokerNew("Gdk", "RGBA", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeRGBAInvoker.Invoke(inArgs[:], nil)

}

var hashRGBAInvoker *gi.Function

// Hash is a representation of the C type gdk_rgba_hash.
func (recv *RGBA) Hash() uint32 {
	if hashRGBAInvoker == nil {
		hashRGBAInvoker = gi.StructFunctionInvokerNew("Gdk", "RGBA", "hash")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hashRGBAInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_rgba_parse' : return type 'gboolean' not supported

var toStringRGBAInvoker *gi.Function

// ToString is a representation of the C type gdk_rgba_to_string.
func (recv *RGBA) ToString() string {
	if toStringRGBAInvoker == nil {
		toStringRGBAInvoker = gi.StructFunctionInvokerNew("Gdk", "RGBA", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringRGBAInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

type Rectangle struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

// UNSUPPORTED : C value 'gdk_rectangle_equal' : parameter 'rect2' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'gdk_rectangle_intersect' : parameter 'src2' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'gdk_rectangle_union' : parameter 'src2' of type 'Rectangle' not supported

type TimeCoord struct {
	native uintptr
	Time   uint32
	// UNSUPPORTED : C value 'axes' : missing Type
}

type WindowAttr struct {
	native    uintptr
	Title     string
	EventMask int32
	X         int32
	Y         int32
	Width     int32
	Height    int32
	// UNSUPPORTED : C value 'wclass' : no Go type for 'WindowWindowClass'
	// UNSUPPORTED : C value 'visual' : no Go type for 'Visual'
	// UNSUPPORTED : C value 'window_type' : no Go type for 'WindowType'
	// UNSUPPORTED : C value 'cursor' : no Go type for 'Cursor'
	WmclassName  string
	WmclassClass string
	// UNSUPPORTED : C value 'override_redirect' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'type_hint' : no Go type for 'WindowTypeHint'
}

type WindowClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'pick_embedded_child' : missing Type
	// UNSUPPORTED : C value 'to_embedder' : missing Type
	// UNSUPPORTED : C value 'from_embedder' : missing Type
	// UNSUPPORTED : C value 'create_surface' : missing Type
	// UNSUPPORTED : C value '_gdk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gdk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gdk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gdk_reserved4' : missing Type
	// UNSUPPORTED : C value '_gdk_reserved5' : missing Type
	// UNSUPPORTED : C value '_gdk_reserved6' : missing Type
	// UNSUPPORTED : C value '_gdk_reserved7' : missing Type
	// UNSUPPORTED : C value '_gdk_reserved8' : missing Type
}

type WindowRedirect struct {
	native uintptr
}
