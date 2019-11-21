// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var atomStruct *gi.Struct
var atomStructOnce sync.Once

func atomStructSet() {
	atomStructOnce.Do(func() {
		atomStruct = gi.StructNew("Gdk", "Atom")
	})
}

type Atom struct {
	native uintptr
}

var atomNameFunction *gi.Function
var atomNameFunction_Once sync.Once

func atomNameFunction_Set() {
	atomNameFunction_Once.Do(func() {
		atomNameFunction = gi.FunctionInvokerNew("Gdk", "name")
	})
}

var nameAtomInvoker *gi.Function

// Name is a representation of the C type gdk_atom_name.
func (recv *Atom) Name() string {
	atomNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := atomNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var colorStruct *gi.Struct
var colorStructOnce sync.Once

func colorStructSet() {
	colorStructOnce.Do(func() {
		colorStruct = gi.StructNew("Gdk", "Color")
	})
}

type Color struct {
	native uintptr
	Pixel  uint32
	Red    uint16
	Green  uint16
	Blue   uint16
}

var colorCopyFunction *gi.Function
var colorCopyFunction_Once sync.Once

func colorCopyFunction_Set() {
	colorCopyFunction_Once.Do(func() {
		colorCopyFunction = gi.FunctionInvokerNew("Gdk", "copy")
	})
}

var copyColorInvoker *gi.Function

// Copy is a representation of the C type gdk_color_copy.
func (recv *Color) Copy() *Color {
	colorCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := colorCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Color{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gdk_color_equal' : parameter 'colorb' of type 'Color' not supported

var colorFreeFunction *gi.Function
var colorFreeFunction_Once sync.Once

func colorFreeFunction_Set() {
	colorFreeFunction_Once.Do(func() {
		colorFreeFunction = gi.FunctionInvokerNew("Gdk", "free")
	})
}

var freeColorInvoker *gi.Function

// Free is a representation of the C type gdk_color_free.
func (recv *Color) Free() {
	colorFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	colorFreeFunction.Invoke(inArgs[:], nil)

}

var colorHashFunction *gi.Function
var colorHashFunction_Once sync.Once

func colorHashFunction_Set() {
	colorHashFunction_Once.Do(func() {
		colorHashFunction = gi.FunctionInvokerNew("Gdk", "hash")
	})
}

var hashColorInvoker *gi.Function

// Hash is a representation of the C type gdk_color_hash.
func (recv *Color) Hash() uint32 {
	colorHashFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := colorHashFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var colorToStringFunction *gi.Function
var colorToStringFunction_Once sync.Once

func colorToStringFunction_Set() {
	colorToStringFunction_Once.Do(func() {
		colorToStringFunction = gi.FunctionInvokerNew("Gdk", "to_string")
	})
}

var toStringColorInvoker *gi.Function

// ToString is a representation of the C type gdk_color_to_string.
func (recv *Color) ToString() string {
	colorToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := colorToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var devicePadInterfaceStruct *gi.Struct
var devicePadInterfaceStructOnce sync.Once

func devicePadInterfaceStructSet() {
	devicePadInterfaceStructOnce.Do(func() {
		devicePadInterfaceStruct = gi.StructNew("Gdk", "DevicePadInterface")
	})
}

type DevicePadInterface struct {
	native uintptr
}

var drawingContextClassStruct *gi.Struct
var drawingContextClassStructOnce sync.Once

func drawingContextClassStructSet() {
	drawingContextClassStructOnce.Do(func() {
		drawingContextClassStruct = gi.StructNew("Gdk", "DrawingContextClass")
	})
}

type DrawingContextClass struct {
	native uintptr
}

var eventAnyStruct *gi.Struct
var eventAnyStructOnce sync.Once

func eventAnyStructSet() {
	eventAnyStructOnce.Do(func() {
		eventAnyStruct = gi.StructNew("Gdk", "EventAny")
	})
}

type EventAny struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
}

var eventButtonStruct *gi.Struct
var eventButtonStructOnce sync.Once

func eventButtonStructSet() {
	eventButtonStructOnce.Do(func() {
		eventButtonStruct = gi.StructNew("Gdk", "EventButton")
	})
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

var eventConfigureStruct *gi.Struct
var eventConfigureStructOnce sync.Once

func eventConfigureStructSet() {
	eventConfigureStructOnce.Do(func() {
		eventConfigureStruct = gi.StructNew("Gdk", "EventConfigure")
	})
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

var eventCrossingStruct *gi.Struct
var eventCrossingStructOnce sync.Once

func eventCrossingStructSet() {
	eventCrossingStructOnce.Do(func() {
		eventCrossingStruct = gi.StructNew("Gdk", "EventCrossing")
	})
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

var eventDNDStruct *gi.Struct
var eventDNDStructOnce sync.Once

func eventDNDStructSet() {
	eventDNDStructOnce.Do(func() {
		eventDNDStruct = gi.StructNew("Gdk", "EventDND")
	})
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

var eventExposeStruct *gi.Struct
var eventExposeStructOnce sync.Once

func eventExposeStructSet() {
	eventExposeStructOnce.Do(func() {
		eventExposeStruct = gi.StructNew("Gdk", "EventExpose")
	})
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

var eventFocusStruct *gi.Struct
var eventFocusStructOnce sync.Once

func eventFocusStructSet() {
	eventFocusStructOnce.Do(func() {
		eventFocusStruct = gi.StructNew("Gdk", "EventFocus")
	})
}

type EventFocus struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	In        int16
}

var eventGrabBrokenStruct *gi.Struct
var eventGrabBrokenStructOnce sync.Once

func eventGrabBrokenStructSet() {
	eventGrabBrokenStructOnce.Do(func() {
		eventGrabBrokenStruct = gi.StructNew("Gdk", "EventGrabBroken")
	})
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

var eventKeyStruct *gi.Struct
var eventKeyStructOnce sync.Once

func eventKeyStructSet() {
	eventKeyStructOnce.Do(func() {
		eventKeyStruct = gi.StructNew("Gdk", "EventKey")
	})
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

var eventMotionStruct *gi.Struct
var eventMotionStructOnce sync.Once

func eventMotionStructSet() {
	eventMotionStructOnce.Do(func() {
		eventMotionStruct = gi.StructNew("Gdk", "EventMotion")
	})
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

var eventOwnerChangeStruct *gi.Struct
var eventOwnerChangeStructOnce sync.Once

func eventOwnerChangeStructSet() {
	eventOwnerChangeStructOnce.Do(func() {
		eventOwnerChangeStruct = gi.StructNew("Gdk", "EventOwnerChange")
	})
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

var eventPadAxisStruct *gi.Struct
var eventPadAxisStructOnce sync.Once

func eventPadAxisStructSet() {
	eventPadAxisStructOnce.Do(func() {
		eventPadAxisStruct = gi.StructNew("Gdk", "EventPadAxis")
	})
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

var eventPadButtonStruct *gi.Struct
var eventPadButtonStructOnce sync.Once

func eventPadButtonStructSet() {
	eventPadButtonStructOnce.Do(func() {
		eventPadButtonStruct = gi.StructNew("Gdk", "EventPadButton")
	})
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

var eventPadGroupModeStruct *gi.Struct
var eventPadGroupModeStructOnce sync.Once

func eventPadGroupModeStructSet() {
	eventPadGroupModeStructOnce.Do(func() {
		eventPadGroupModeStruct = gi.StructNew("Gdk", "EventPadGroupMode")
	})
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

var eventPropertyStruct *gi.Struct
var eventPropertyStructOnce sync.Once

func eventPropertyStructSet() {
	eventPropertyStructOnce.Do(func() {
		eventPropertyStruct = gi.StructNew("Gdk", "EventProperty")
	})
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

var eventProximityStruct *gi.Struct
var eventProximityStructOnce sync.Once

func eventProximityStructSet() {
	eventProximityStructOnce.Do(func() {
		eventProximityStruct = gi.StructNew("Gdk", "EventProximity")
	})
}

type EventProximity struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'device' : no Go type for 'Device'
}

var eventScrollStruct *gi.Struct
var eventScrollStructOnce sync.Once

func eventScrollStructSet() {
	eventScrollStructOnce.Do(func() {
		eventScrollStruct = gi.StructNew("Gdk", "EventScroll")
	})
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

var eventSelectionStruct *gi.Struct
var eventSelectionStructOnce sync.Once

func eventSelectionStructSet() {
	eventSelectionStructOnce.Do(func() {
		eventSelectionStruct = gi.StructNew("Gdk", "EventSelection")
	})
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

var eventSequenceStruct *gi.Struct
var eventSequenceStructOnce sync.Once

func eventSequenceStructSet() {
	eventSequenceStructOnce.Do(func() {
		eventSequenceStruct = gi.StructNew("Gdk", "EventSequence")
	})
}

type EventSequence struct {
	native uintptr
}

var eventSettingStruct *gi.Struct
var eventSettingStructOnce sync.Once

func eventSettingStructSet() {
	eventSettingStructOnce.Do(func() {
		eventSettingStruct = gi.StructNew("Gdk", "EventSetting")
	})
}

type EventSetting struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	// UNSUPPORTED : C value 'action' : no Go type for 'SettingAction'
	Name string
}

var eventTouchStruct *gi.Struct
var eventTouchStructOnce sync.Once

func eventTouchStructSet() {
	eventTouchStructOnce.Do(func() {
		eventTouchStruct = gi.StructNew("Gdk", "EventTouch")
	})
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

var eventTouchpadPinchStruct *gi.Struct
var eventTouchpadPinchStructOnce sync.Once

func eventTouchpadPinchStructSet() {
	eventTouchpadPinchStructOnce.Do(func() {
		eventTouchpadPinchStruct = gi.StructNew("Gdk", "EventTouchpadPinch")
	})
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

var eventTouchpadSwipeStruct *gi.Struct
var eventTouchpadSwipeStructOnce sync.Once

func eventTouchpadSwipeStructSet() {
	eventTouchpadSwipeStructOnce.Do(func() {
		eventTouchpadSwipeStruct = gi.StructNew("Gdk", "EventTouchpadSwipe")
	})
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

var eventVisibilityStruct *gi.Struct
var eventVisibilityStructOnce sync.Once

func eventVisibilityStructSet() {
	eventVisibilityStructOnce.Do(func() {
		eventVisibilityStruct = gi.StructNew("Gdk", "EventVisibility")
	})
}

type EventVisibility struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	// UNSUPPORTED : C value 'state' : no Go type for 'VisibilityState'
}

var eventWindowStateStruct *gi.Struct
var eventWindowStateStructOnce sync.Once

func eventWindowStateStructSet() {
	eventWindowStateStructOnce.Do(func() {
		eventWindowStateStruct = gi.StructNew("Gdk", "EventWindowState")
	})
}

type EventWindowState struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	// UNSUPPORTED : C value 'changed_mask' : no Go type for 'WindowState'
	// UNSUPPORTED : C value 'new_window_state' : no Go type for 'WindowState'
}

var frameClockClassStruct *gi.Struct
var frameClockClassStructOnce sync.Once

func frameClockClassStructSet() {
	frameClockClassStructOnce.Do(func() {
		frameClockClassStruct = gi.StructNew("Gdk", "FrameClockClass")
	})
}

type FrameClockClass struct {
	native uintptr
}

var frameClockPrivateStruct *gi.Struct
var frameClockPrivateStructOnce sync.Once

func frameClockPrivateStructSet() {
	frameClockPrivateStructOnce.Do(func() {
		frameClockPrivateStruct = gi.StructNew("Gdk", "FrameClockPrivate")
	})
}

type FrameClockPrivate struct {
	native uintptr
}

var frameTimingsStruct *gi.Struct
var frameTimingsStructOnce sync.Once

func frameTimingsStructSet() {
	frameTimingsStructOnce.Do(func() {
		frameTimingsStruct = gi.StructNew("Gdk", "FrameTimings")
	})
}

type FrameTimings struct {
	native uintptr
}

// UNSUPPORTED : C value 'gdk_frame_timings_get_complete' : return type 'gboolean' not supported

var frameTimingsGetFrameCounterFunction *gi.Function
var frameTimingsGetFrameCounterFunction_Once sync.Once

func frameTimingsGetFrameCounterFunction_Set() {
	frameTimingsGetFrameCounterFunction_Once.Do(func() {
		frameTimingsGetFrameCounterFunction = gi.FunctionInvokerNew("Gdk", "get_frame_counter")
	})
}

var getFrameCounterFrameTimingsInvoker *gi.Function

// GetFrameCounter is a representation of the C type gdk_frame_timings_get_frame_counter.
func (recv *FrameTimings) GetFrameCounter() int64 {
	frameTimingsGetFrameCounterFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := frameTimingsGetFrameCounterFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var frameTimingsGetFrameTimeFunction *gi.Function
var frameTimingsGetFrameTimeFunction_Once sync.Once

func frameTimingsGetFrameTimeFunction_Set() {
	frameTimingsGetFrameTimeFunction_Once.Do(func() {
		frameTimingsGetFrameTimeFunction = gi.FunctionInvokerNew("Gdk", "get_frame_time")
	})
}

var getFrameTimeFrameTimingsInvoker *gi.Function

// GetFrameTime is a representation of the C type gdk_frame_timings_get_frame_time.
func (recv *FrameTimings) GetFrameTime() int64 {
	frameTimingsGetFrameTimeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := frameTimingsGetFrameTimeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var frameTimingsGetPredictedPresentationTimeFunction *gi.Function
var frameTimingsGetPredictedPresentationTimeFunction_Once sync.Once

func frameTimingsGetPredictedPresentationTimeFunction_Set() {
	frameTimingsGetPredictedPresentationTimeFunction_Once.Do(func() {
		frameTimingsGetPredictedPresentationTimeFunction = gi.FunctionInvokerNew("Gdk", "get_predicted_presentation_time")
	})
}

var getPredictedPresentationTimeFrameTimingsInvoker *gi.Function

// GetPredictedPresentationTime is a representation of the C type gdk_frame_timings_get_predicted_presentation_time.
func (recv *FrameTimings) GetPredictedPresentationTime() int64 {
	frameTimingsGetPredictedPresentationTimeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := frameTimingsGetPredictedPresentationTimeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var frameTimingsGetPresentationTimeFunction *gi.Function
var frameTimingsGetPresentationTimeFunction_Once sync.Once

func frameTimingsGetPresentationTimeFunction_Set() {
	frameTimingsGetPresentationTimeFunction_Once.Do(func() {
		frameTimingsGetPresentationTimeFunction = gi.FunctionInvokerNew("Gdk", "get_presentation_time")
	})
}

var getPresentationTimeFrameTimingsInvoker *gi.Function

// GetPresentationTime is a representation of the C type gdk_frame_timings_get_presentation_time.
func (recv *FrameTimings) GetPresentationTime() int64 {
	frameTimingsGetPresentationTimeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := frameTimingsGetPresentationTimeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var frameTimingsGetRefreshIntervalFunction *gi.Function
var frameTimingsGetRefreshIntervalFunction_Once sync.Once

func frameTimingsGetRefreshIntervalFunction_Set() {
	frameTimingsGetRefreshIntervalFunction_Once.Do(func() {
		frameTimingsGetRefreshIntervalFunction = gi.FunctionInvokerNew("Gdk", "get_refresh_interval")
	})
}

var getRefreshIntervalFrameTimingsInvoker *gi.Function

// GetRefreshInterval is a representation of the C type gdk_frame_timings_get_refresh_interval.
func (recv *FrameTimings) GetRefreshInterval() int64 {
	frameTimingsGetRefreshIntervalFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := frameTimingsGetRefreshIntervalFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var frameTimingsRefFunction *gi.Function
var frameTimingsRefFunction_Once sync.Once

func frameTimingsRefFunction_Set() {
	frameTimingsRefFunction_Once.Do(func() {
		frameTimingsRefFunction = gi.FunctionInvokerNew("Gdk", "ref")
	})
}

var refFrameTimingsInvoker *gi.Function

// Ref is a representation of the C type gdk_frame_timings_ref.
func (recv *FrameTimings) Ref() *FrameTimings {
	frameTimingsRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := frameTimingsRefFunction.Invoke(inArgs[:], nil)

	retGo := &FrameTimings{native: ret.Pointer()}

	return retGo
}

var frameTimingsUnrefFunction *gi.Function
var frameTimingsUnrefFunction_Once sync.Once

func frameTimingsUnrefFunction_Set() {
	frameTimingsUnrefFunction_Once.Do(func() {
		frameTimingsUnrefFunction = gi.FunctionInvokerNew("Gdk", "unref")
	})
}

var unrefFrameTimingsInvoker *gi.Function

// Unref is a representation of the C type gdk_frame_timings_unref.
func (recv *FrameTimings) Unref() {
	frameTimingsUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	frameTimingsUnrefFunction.Invoke(inArgs[:], nil)

}

var geometryStruct *gi.Struct
var geometryStructOnce sync.Once

func geometryStructSet() {
	geometryStructOnce.Do(func() {
		geometryStruct = gi.StructNew("Gdk", "Geometry")
	})
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

var keymapKeyStruct *gi.Struct
var keymapKeyStructOnce sync.Once

func keymapKeyStructSet() {
	keymapKeyStructOnce.Do(func() {
		keymapKeyStruct = gi.StructNew("Gdk", "KeymapKey")
	})
}

type KeymapKey struct {
	native  uintptr
	Keycode uint32
	Group   int32
	Level   int32
}

var monitorClassStruct *gi.Struct
var monitorClassStructOnce sync.Once

func monitorClassStructSet() {
	monitorClassStructOnce.Do(func() {
		monitorClassStruct = gi.StructNew("Gdk", "MonitorClass")
	})
}

type MonitorClass struct {
	native uintptr
}

var pointStruct *gi.Struct
var pointStructOnce sync.Once

func pointStructSet() {
	pointStructOnce.Do(func() {
		pointStruct = gi.StructNew("Gdk", "Point")
	})
}

type Point struct {
	native uintptr
	X      int32
	Y      int32
}

var rGBAStruct *gi.Struct
var rGBAStructOnce sync.Once

func rGBAStructSet() {
	rGBAStructOnce.Do(func() {
		rGBAStruct = gi.StructNew("Gdk", "RGBA")
	})
}

type RGBA struct {
	native uintptr
	// UNSUPPORTED : C value 'red' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'green' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'blue' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'alpha' : no Go type for 'gdouble'
}

var rGBACopyFunction *gi.Function
var rGBACopyFunction_Once sync.Once

func rGBACopyFunction_Set() {
	rGBACopyFunction_Once.Do(func() {
		rGBACopyFunction = gi.FunctionInvokerNew("Gdk", "copy")
	})
}

var copyRGBAInvoker *gi.Function

// Copy is a representation of the C type gdk_rgba_copy.
func (recv *RGBA) Copy() *RGBA {
	rGBACopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := rGBACopyFunction.Invoke(inArgs[:], nil)

	retGo := &RGBA{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gdk_rgba_equal' : parameter 'p2' of type 'RGBA' not supported

var rGBAFreeFunction *gi.Function
var rGBAFreeFunction_Once sync.Once

func rGBAFreeFunction_Set() {
	rGBAFreeFunction_Once.Do(func() {
		rGBAFreeFunction = gi.FunctionInvokerNew("Gdk", "free")
	})
}

var freeRGBAInvoker *gi.Function

// Free is a representation of the C type gdk_rgba_free.
func (recv *RGBA) Free() {
	rGBAFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	rGBAFreeFunction.Invoke(inArgs[:], nil)

}

var rGBAHashFunction *gi.Function
var rGBAHashFunction_Once sync.Once

func rGBAHashFunction_Set() {
	rGBAHashFunction_Once.Do(func() {
		rGBAHashFunction = gi.FunctionInvokerNew("Gdk", "hash")
	})
}

var hashRGBAInvoker *gi.Function

// Hash is a representation of the C type gdk_rgba_hash.
func (recv *RGBA) Hash() uint32 {
	rGBAHashFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := rGBAHashFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_rgba_parse' : return type 'gboolean' not supported

var rGBAToStringFunction *gi.Function
var rGBAToStringFunction_Once sync.Once

func rGBAToStringFunction_Set() {
	rGBAToStringFunction_Once.Do(func() {
		rGBAToStringFunction = gi.FunctionInvokerNew("Gdk", "to_string")
	})
}

var toStringRGBAInvoker *gi.Function

// ToString is a representation of the C type gdk_rgba_to_string.
func (recv *RGBA) ToString() string {
	rGBAToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := rGBAToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var rectangleStruct *gi.Struct
var rectangleStructOnce sync.Once

func rectangleStructSet() {
	rectangleStructOnce.Do(func() {
		rectangleStruct = gi.StructNew("Gdk", "Rectangle")
	})
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

var timeCoordStruct *gi.Struct
var timeCoordStructOnce sync.Once

func timeCoordStructSet() {
	timeCoordStructOnce.Do(func() {
		timeCoordStruct = gi.StructNew("Gdk", "TimeCoord")
	})
}

type TimeCoord struct {
	native uintptr
	Time   uint32
	// UNSUPPORTED : C value 'axes' : missing Type
}

var windowAttrStruct *gi.Struct
var windowAttrStructOnce sync.Once

func windowAttrStructSet() {
	windowAttrStructOnce.Do(func() {
		windowAttrStruct = gi.StructNew("Gdk", "WindowAttr")
	})
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

var windowClassStruct *gi.Struct
var windowClassStructOnce sync.Once

func windowClassStructSet() {
	windowClassStructOnce.Do(func() {
		windowClassStruct = gi.StructNew("Gdk", "WindowClass")
	})
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

var windowRedirectStruct *gi.Struct
var windowRedirectStructOnce sync.Once

func windowRedirectStructSet() {
	windowRedirectStructOnce.Do(func() {
		windowRedirectStruct = gi.StructNew("Gdk", "WindowRedirect")
	})
}

type WindowRedirect struct {
	native uintptr
}
