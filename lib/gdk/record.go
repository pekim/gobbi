// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var atomStruct *gi.Struct
var atomStruct_Once sync.Once

func atomStruct_Set() {
	atomStruct_Once.Do(func() {
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
		atomStruct_Set()
		atomNameFunction = atomStruct.InvokerNew("name")
	})
}

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
var colorStruct_Once sync.Once

func colorStruct_Set() {
	colorStruct_Once.Do(func() {
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
		colorStruct_Set()
		colorCopyFunction = colorStruct.InvokerNew("copy")
	})
}

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
		colorStruct_Set()
		colorFreeFunction = colorStruct.InvokerNew("free")
	})
}

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
		colorStruct_Set()
		colorHashFunction = colorStruct.InvokerNew("hash")
	})
}

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
		colorStruct_Set()
		colorToStringFunction = colorStruct.InvokerNew("to_string")
	})
}

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
var devicePadInterfaceStruct_Once sync.Once

func devicePadInterfaceStruct_Set() {
	devicePadInterfaceStruct_Once.Do(func() {
		devicePadInterfaceStruct = gi.StructNew("Gdk", "DevicePadInterface")
	})
}

type DevicePadInterface struct {
	native uintptr
}

var drawingContextClassStruct *gi.Struct
var drawingContextClassStruct_Once sync.Once

func drawingContextClassStruct_Set() {
	drawingContextClassStruct_Once.Do(func() {
		drawingContextClassStruct = gi.StructNew("Gdk", "DrawingContextClass")
	})
}

type DrawingContextClass struct {
	native uintptr
}

var eventAnyStruct *gi.Struct
var eventAnyStruct_Once sync.Once

func eventAnyStruct_Set() {
	eventAnyStruct_Once.Do(func() {
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
var eventButtonStruct_Once sync.Once

func eventButtonStruct_Set() {
	eventButtonStruct_Once.Do(func() {
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
var eventConfigureStruct_Once sync.Once

func eventConfigureStruct_Set() {
	eventConfigureStruct_Once.Do(func() {
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
var eventCrossingStruct_Once sync.Once

func eventCrossingStruct_Set() {
	eventCrossingStruct_Once.Do(func() {
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
var eventDNDStruct_Once sync.Once

func eventDNDStruct_Set() {
	eventDNDStruct_Once.Do(func() {
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
var eventExposeStruct_Once sync.Once

func eventExposeStruct_Set() {
	eventExposeStruct_Once.Do(func() {
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
var eventFocusStruct_Once sync.Once

func eventFocusStruct_Set() {
	eventFocusStruct_Once.Do(func() {
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
var eventGrabBrokenStruct_Once sync.Once

func eventGrabBrokenStruct_Set() {
	eventGrabBrokenStruct_Once.Do(func() {
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
var eventKeyStruct_Once sync.Once

func eventKeyStruct_Set() {
	eventKeyStruct_Once.Do(func() {
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
var eventMotionStruct_Once sync.Once

func eventMotionStruct_Set() {
	eventMotionStruct_Once.Do(func() {
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
var eventOwnerChangeStruct_Once sync.Once

func eventOwnerChangeStruct_Set() {
	eventOwnerChangeStruct_Once.Do(func() {
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
var eventPadAxisStruct_Once sync.Once

func eventPadAxisStruct_Set() {
	eventPadAxisStruct_Once.Do(func() {
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
var eventPadButtonStruct_Once sync.Once

func eventPadButtonStruct_Set() {
	eventPadButtonStruct_Once.Do(func() {
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
var eventPadGroupModeStruct_Once sync.Once

func eventPadGroupModeStruct_Set() {
	eventPadGroupModeStruct_Once.Do(func() {
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
var eventPropertyStruct_Once sync.Once

func eventPropertyStruct_Set() {
	eventPropertyStruct_Once.Do(func() {
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
var eventProximityStruct_Once sync.Once

func eventProximityStruct_Set() {
	eventProximityStruct_Once.Do(func() {
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
var eventScrollStruct_Once sync.Once

func eventScrollStruct_Set() {
	eventScrollStruct_Once.Do(func() {
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
var eventSelectionStruct_Once sync.Once

func eventSelectionStruct_Set() {
	eventSelectionStruct_Once.Do(func() {
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
var eventSequenceStruct_Once sync.Once

func eventSequenceStruct_Set() {
	eventSequenceStruct_Once.Do(func() {
		eventSequenceStruct = gi.StructNew("Gdk", "EventSequence")
	})
}

type EventSequence struct {
	native uintptr
}

var eventSettingStruct *gi.Struct
var eventSettingStruct_Once sync.Once

func eventSettingStruct_Set() {
	eventSettingStruct_Once.Do(func() {
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
var eventTouchStruct_Once sync.Once

func eventTouchStruct_Set() {
	eventTouchStruct_Once.Do(func() {
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
var eventTouchpadPinchStruct_Once sync.Once

func eventTouchpadPinchStruct_Set() {
	eventTouchpadPinchStruct_Once.Do(func() {
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
var eventTouchpadSwipeStruct_Once sync.Once

func eventTouchpadSwipeStruct_Set() {
	eventTouchpadSwipeStruct_Once.Do(func() {
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
var eventVisibilityStruct_Once sync.Once

func eventVisibilityStruct_Set() {
	eventVisibilityStruct_Once.Do(func() {
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
var eventWindowStateStruct_Once sync.Once

func eventWindowStateStruct_Set() {
	eventWindowStateStruct_Once.Do(func() {
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
var frameClockClassStruct_Once sync.Once

func frameClockClassStruct_Set() {
	frameClockClassStruct_Once.Do(func() {
		frameClockClassStruct = gi.StructNew("Gdk", "FrameClockClass")
	})
}

type FrameClockClass struct {
	native uintptr
}

var frameClockPrivateStruct *gi.Struct
var frameClockPrivateStruct_Once sync.Once

func frameClockPrivateStruct_Set() {
	frameClockPrivateStruct_Once.Do(func() {
		frameClockPrivateStruct = gi.StructNew("Gdk", "FrameClockPrivate")
	})
}

type FrameClockPrivate struct {
	native uintptr
}

var frameTimingsStruct *gi.Struct
var frameTimingsStruct_Once sync.Once

func frameTimingsStruct_Set() {
	frameTimingsStruct_Once.Do(func() {
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
		frameTimingsStruct_Set()
		frameTimingsGetFrameCounterFunction = frameTimingsStruct.InvokerNew("get_frame_counter")
	})
}

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
		frameTimingsStruct_Set()
		frameTimingsGetFrameTimeFunction = frameTimingsStruct.InvokerNew("get_frame_time")
	})
}

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
		frameTimingsStruct_Set()
		frameTimingsGetPredictedPresentationTimeFunction = frameTimingsStruct.InvokerNew("get_predicted_presentation_time")
	})
}

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
		frameTimingsStruct_Set()
		frameTimingsGetPresentationTimeFunction = frameTimingsStruct.InvokerNew("get_presentation_time")
	})
}

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
		frameTimingsStruct_Set()
		frameTimingsGetRefreshIntervalFunction = frameTimingsStruct.InvokerNew("get_refresh_interval")
	})
}

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
		frameTimingsStruct_Set()
		frameTimingsRefFunction = frameTimingsStruct.InvokerNew("ref")
	})
}

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
		frameTimingsStruct_Set()
		frameTimingsUnrefFunction = frameTimingsStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type gdk_frame_timings_unref.
func (recv *FrameTimings) Unref() {
	frameTimingsUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	frameTimingsUnrefFunction.Invoke(inArgs[:], nil)

}

var geometryStruct *gi.Struct
var geometryStruct_Once sync.Once

func geometryStruct_Set() {
	geometryStruct_Once.Do(func() {
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
var keymapKeyStruct_Once sync.Once

func keymapKeyStruct_Set() {
	keymapKeyStruct_Once.Do(func() {
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
var monitorClassStruct_Once sync.Once

func monitorClassStruct_Set() {
	monitorClassStruct_Once.Do(func() {
		monitorClassStruct = gi.StructNew("Gdk", "MonitorClass")
	})
}

type MonitorClass struct {
	native uintptr
}

var pointStruct *gi.Struct
var pointStruct_Once sync.Once

func pointStruct_Set() {
	pointStruct_Once.Do(func() {
		pointStruct = gi.StructNew("Gdk", "Point")
	})
}

type Point struct {
	native uintptr
	X      int32
	Y      int32
}

var rGBAStruct *gi.Struct
var rGBAStruct_Once sync.Once

func rGBAStruct_Set() {
	rGBAStruct_Once.Do(func() {
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
		rGBAStruct_Set()
		rGBACopyFunction = rGBAStruct.InvokerNew("copy")
	})
}

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
		rGBAStruct_Set()
		rGBAFreeFunction = rGBAStruct.InvokerNew("free")
	})
}

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
		rGBAStruct_Set()
		rGBAHashFunction = rGBAStruct.InvokerNew("hash")
	})
}

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
		rGBAStruct_Set()
		rGBAToStringFunction = rGBAStruct.InvokerNew("to_string")
	})
}

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
var rectangleStruct_Once sync.Once

func rectangleStruct_Set() {
	rectangleStruct_Once.Do(func() {
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
var timeCoordStruct_Once sync.Once

func timeCoordStruct_Set() {
	timeCoordStruct_Once.Do(func() {
		timeCoordStruct = gi.StructNew("Gdk", "TimeCoord")
	})
}

type TimeCoord struct {
	native uintptr
	Time   uint32
	// UNSUPPORTED : C value 'axes' : missing Type
}

var windowAttrStruct *gi.Struct
var windowAttrStruct_Once sync.Once

func windowAttrStruct_Set() {
	windowAttrStruct_Once.Do(func() {
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
var windowClassStruct_Once sync.Once

func windowClassStruct_Set() {
	windowClassStruct_Once.Do(func() {
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
var windowRedirectStruct_Once sync.Once

func windowRedirectStruct_Set() {
	windowRedirectStruct_Once.Do(func() {
		windowRedirectStruct = gi.StructNew("Gdk", "WindowRedirect")
	})
}

type WindowRedirect struct {
	native uintptr
}
