// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var atomStruct *gi.Struct
var atomStruct_Once sync.Once

func atomStruct_Set() error {
	var err error
	atomStruct_Once.Do(func() {
		atomStruct, err = gi.StructNew("Gdk", "Atom")
	})
	return err
}

type Atom struct {
	native uintptr
}

var atomNameFunction *gi.Function
var atomNameFunction_Once sync.Once

func atomNameFunction_Set() error {
	var err error
	atomNameFunction_Once.Do(func() {
		err = atomStruct_Set()
		if err != nil {
			return
		}
		atomNameFunction, err = atomStruct.InvokerNew("name")
	})
	return err
}

// Name is a representation of the C type gdk_atom_name.
func (recv *Atom) Name() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := atomNameFunction_Set()
	if err == nil {
		ret = atomNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var colorStruct *gi.Struct
var colorStruct_Once sync.Once

func colorStruct_Set() error {
	var err error
	colorStruct_Once.Do(func() {
		colorStruct, err = gi.StructNew("Gdk", "Color")
	})
	return err
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

func colorCopyFunction_Set() error {
	var err error
	colorCopyFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorCopyFunction, err = colorStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gdk_color_copy.
func (recv *Color) Copy() (*Color, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorCopyFunction_Set()
	if err == nil {
		ret = colorCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Color{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'gdk_color_equal' : parameter 'colorb' of type 'Color' not supported

var colorFreeFunction *gi.Function
var colorFreeFunction_Once sync.Once

func colorFreeFunction_Set() error {
	var err error
	colorFreeFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorFreeFunction, err = colorStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gdk_color_free.
func (recv *Color) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := colorFreeFunction_Set()
	if err == nil {
		colorFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var colorHashFunction *gi.Function
var colorHashFunction_Once sync.Once

func colorHashFunction_Set() error {
	var err error
	colorHashFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorHashFunction, err = colorStruct.InvokerNew("hash")
	})
	return err
}

// Hash is a representation of the C type gdk_color_hash.
func (recv *Color) Hash() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorHashFunction_Set()
	if err == nil {
		ret = colorHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var colorToStringFunction *gi.Function
var colorToStringFunction_Once sync.Once

func colorToStringFunction_Set() error {
	var err error
	colorToStringFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorToStringFunction, err = colorStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type gdk_color_to_string.
func (recv *Color) ToString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorToStringFunction_Set()
	if err == nil {
		ret = colorToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var devicePadInterfaceStruct *gi.Struct
var devicePadInterfaceStruct_Once sync.Once

func devicePadInterfaceStruct_Set() error {
	var err error
	devicePadInterfaceStruct_Once.Do(func() {
		devicePadInterfaceStruct, err = gi.StructNew("Gdk", "DevicePadInterface")
	})
	return err
}

type DevicePadInterface struct {
	native uintptr
}

var drawingContextClassStruct *gi.Struct
var drawingContextClassStruct_Once sync.Once

func drawingContextClassStruct_Set() error {
	var err error
	drawingContextClassStruct_Once.Do(func() {
		drawingContextClassStruct, err = gi.StructNew("Gdk", "DrawingContextClass")
	})
	return err
}

type DrawingContextClass struct {
	native uintptr
}

var eventAnyStruct *gi.Struct
var eventAnyStruct_Once sync.Once

func eventAnyStruct_Set() error {
	var err error
	eventAnyStruct_Once.Do(func() {
		eventAnyStruct, err = gi.StructNew("Gdk", "EventAny")
	})
	return err
}

type EventAny struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
}

var eventButtonStruct *gi.Struct
var eventButtonStruct_Once sync.Once

func eventButtonStruct_Set() error {
	var err error
	eventButtonStruct_Once.Do(func() {
		eventButtonStruct, err = gi.StructNew("Gdk", "EventButton")
	})
	return err
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

func eventConfigureStruct_Set() error {
	var err error
	eventConfigureStruct_Once.Do(func() {
		eventConfigureStruct, err = gi.StructNew("Gdk", "EventConfigure")
	})
	return err
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

func eventCrossingStruct_Set() error {
	var err error
	eventCrossingStruct_Once.Do(func() {
		eventCrossingStruct, err = gi.StructNew("Gdk", "EventCrossing")
	})
	return err
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
	Focus bool
	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'
}

var eventDNDStruct *gi.Struct
var eventDNDStruct_Once sync.Once

func eventDNDStruct_Set() error {
	var err error
	eventDNDStruct_Once.Do(func() {
		eventDNDStruct, err = gi.StructNew("Gdk", "EventDND")
	})
	return err
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

func eventExposeStruct_Set() error {
	var err error
	eventExposeStruct_Once.Do(func() {
		eventExposeStruct, err = gi.StructNew("Gdk", "EventExpose")
	})
	return err
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

func eventFocusStruct_Set() error {
	var err error
	eventFocusStruct_Once.Do(func() {
		eventFocusStruct, err = gi.StructNew("Gdk", "EventFocus")
	})
	return err
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

func eventGrabBrokenStruct_Set() error {
	var err error
	eventGrabBrokenStruct_Once.Do(func() {
		eventGrabBrokenStruct, err = gi.StructNew("Gdk", "EventGrabBroken")
	})
	return err
}

type EventGrabBroken struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'
	// UNSUPPORTED : C value 'window' : no Go type for 'Window'
	SendEvent int8
	Keyboard  bool
	Implicit  bool
	// UNSUPPORTED : C value 'grab_window' : no Go type for 'Window'
}

var eventKeyStruct *gi.Struct
var eventKeyStruct_Once sync.Once

func eventKeyStruct_Set() error {
	var err error
	eventKeyStruct_Once.Do(func() {
		eventKeyStruct, err = gi.StructNew("Gdk", "EventKey")
	})
	return err
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

func eventMotionStruct_Set() error {
	var err error
	eventMotionStruct_Once.Do(func() {
		eventMotionStruct, err = gi.StructNew("Gdk", "EventMotion")
	})
	return err
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

func eventOwnerChangeStruct_Set() error {
	var err error
	eventOwnerChangeStruct_Once.Do(func() {
		eventOwnerChangeStruct, err = gi.StructNew("Gdk", "EventOwnerChange")
	})
	return err
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

func eventPadAxisStruct_Set() error {
	var err error
	eventPadAxisStruct_Once.Do(func() {
		eventPadAxisStruct, err = gi.StructNew("Gdk", "EventPadAxis")
	})
	return err
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

func eventPadButtonStruct_Set() error {
	var err error
	eventPadButtonStruct_Once.Do(func() {
		eventPadButtonStruct, err = gi.StructNew("Gdk", "EventPadButton")
	})
	return err
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

func eventPadGroupModeStruct_Set() error {
	var err error
	eventPadGroupModeStruct_Once.Do(func() {
		eventPadGroupModeStruct, err = gi.StructNew("Gdk", "EventPadGroupMode")
	})
	return err
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

func eventPropertyStruct_Set() error {
	var err error
	eventPropertyStruct_Once.Do(func() {
		eventPropertyStruct, err = gi.StructNew("Gdk", "EventProperty")
	})
	return err
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

func eventProximityStruct_Set() error {
	var err error
	eventProximityStruct_Once.Do(func() {
		eventProximityStruct, err = gi.StructNew("Gdk", "EventProximity")
	})
	return err
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

func eventScrollStruct_Set() error {
	var err error
	eventScrollStruct_Once.Do(func() {
		eventScrollStruct, err = gi.StructNew("Gdk", "EventScroll")
	})
	return err
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

func eventSelectionStruct_Set() error {
	var err error
	eventSelectionStruct_Once.Do(func() {
		eventSelectionStruct, err = gi.StructNew("Gdk", "EventSelection")
	})
	return err
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

func eventSequenceStruct_Set() error {
	var err error
	eventSequenceStruct_Once.Do(func() {
		eventSequenceStruct, err = gi.StructNew("Gdk", "EventSequence")
	})
	return err
}

type EventSequence struct {
	native uintptr
}

var eventSettingStruct *gi.Struct
var eventSettingStruct_Once sync.Once

func eventSettingStruct_Set() error {
	var err error
	eventSettingStruct_Once.Do(func() {
		eventSettingStruct, err = gi.StructNew("Gdk", "EventSetting")
	})
	return err
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

func eventTouchStruct_Set() error {
	var err error
	eventTouchStruct_Once.Do(func() {
		eventTouchStruct, err = gi.StructNew("Gdk", "EventTouch")
	})
	return err
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
	Sequence         *EventSequence
	EmulatingPointer bool
	// UNSUPPORTED : C value 'device' : no Go type for 'Device'
	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'
}

var eventTouchpadPinchStruct *gi.Struct
var eventTouchpadPinchStruct_Once sync.Once

func eventTouchpadPinchStruct_Set() error {
	var err error
	eventTouchpadPinchStruct_Once.Do(func() {
		eventTouchpadPinchStruct, err = gi.StructNew("Gdk", "EventTouchpadPinch")
	})
	return err
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

func eventTouchpadSwipeStruct_Set() error {
	var err error
	eventTouchpadSwipeStruct_Once.Do(func() {
		eventTouchpadSwipeStruct, err = gi.StructNew("Gdk", "EventTouchpadSwipe")
	})
	return err
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

func eventVisibilityStruct_Set() error {
	var err error
	eventVisibilityStruct_Once.Do(func() {
		eventVisibilityStruct, err = gi.StructNew("Gdk", "EventVisibility")
	})
	return err
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

func eventWindowStateStruct_Set() error {
	var err error
	eventWindowStateStruct_Once.Do(func() {
		eventWindowStateStruct, err = gi.StructNew("Gdk", "EventWindowState")
	})
	return err
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

func frameClockClassStruct_Set() error {
	var err error
	frameClockClassStruct_Once.Do(func() {
		frameClockClassStruct, err = gi.StructNew("Gdk", "FrameClockClass")
	})
	return err
}

type FrameClockClass struct {
	native uintptr
}

var frameClockPrivateStruct *gi.Struct
var frameClockPrivateStruct_Once sync.Once

func frameClockPrivateStruct_Set() error {
	var err error
	frameClockPrivateStruct_Once.Do(func() {
		frameClockPrivateStruct, err = gi.StructNew("Gdk", "FrameClockPrivate")
	})
	return err
}

type FrameClockPrivate struct {
	native uintptr
}

var frameTimingsStruct *gi.Struct
var frameTimingsStruct_Once sync.Once

func frameTimingsStruct_Set() error {
	var err error
	frameTimingsStruct_Once.Do(func() {
		frameTimingsStruct, err = gi.StructNew("Gdk", "FrameTimings")
	})
	return err
}

type FrameTimings struct {
	native uintptr
}

var frameTimingsGetCompleteFunction *gi.Function
var frameTimingsGetCompleteFunction_Once sync.Once

func frameTimingsGetCompleteFunction_Set() error {
	var err error
	frameTimingsGetCompleteFunction_Once.Do(func() {
		err = frameTimingsStruct_Set()
		if err != nil {
			return
		}
		frameTimingsGetCompleteFunction, err = frameTimingsStruct.InvokerNew("get_complete")
	})
	return err
}

// GetComplete is a representation of the C type gdk_frame_timings_get_complete.
func (recv *FrameTimings) GetComplete() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetCompleteFunction_Set()
	if err == nil {
		ret = frameTimingsGetCompleteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var frameTimingsGetFrameCounterFunction *gi.Function
var frameTimingsGetFrameCounterFunction_Once sync.Once

func frameTimingsGetFrameCounterFunction_Set() error {
	var err error
	frameTimingsGetFrameCounterFunction_Once.Do(func() {
		err = frameTimingsStruct_Set()
		if err != nil {
			return
		}
		frameTimingsGetFrameCounterFunction, err = frameTimingsStruct.InvokerNew("get_frame_counter")
	})
	return err
}

// GetFrameCounter is a representation of the C type gdk_frame_timings_get_frame_counter.
func (recv *FrameTimings) GetFrameCounter() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetFrameCounterFunction_Set()
	if err == nil {
		ret = frameTimingsGetFrameCounterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var frameTimingsGetFrameTimeFunction *gi.Function
var frameTimingsGetFrameTimeFunction_Once sync.Once

func frameTimingsGetFrameTimeFunction_Set() error {
	var err error
	frameTimingsGetFrameTimeFunction_Once.Do(func() {
		err = frameTimingsStruct_Set()
		if err != nil {
			return
		}
		frameTimingsGetFrameTimeFunction, err = frameTimingsStruct.InvokerNew("get_frame_time")
	})
	return err
}

// GetFrameTime is a representation of the C type gdk_frame_timings_get_frame_time.
func (recv *FrameTimings) GetFrameTime() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetFrameTimeFunction_Set()
	if err == nil {
		ret = frameTimingsGetFrameTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var frameTimingsGetPredictedPresentationTimeFunction *gi.Function
var frameTimingsGetPredictedPresentationTimeFunction_Once sync.Once

func frameTimingsGetPredictedPresentationTimeFunction_Set() error {
	var err error
	frameTimingsGetPredictedPresentationTimeFunction_Once.Do(func() {
		err = frameTimingsStruct_Set()
		if err != nil {
			return
		}
		frameTimingsGetPredictedPresentationTimeFunction, err = frameTimingsStruct.InvokerNew("get_predicted_presentation_time")
	})
	return err
}

// GetPredictedPresentationTime is a representation of the C type gdk_frame_timings_get_predicted_presentation_time.
func (recv *FrameTimings) GetPredictedPresentationTime() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetPredictedPresentationTimeFunction_Set()
	if err == nil {
		ret = frameTimingsGetPredictedPresentationTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var frameTimingsGetPresentationTimeFunction *gi.Function
var frameTimingsGetPresentationTimeFunction_Once sync.Once

func frameTimingsGetPresentationTimeFunction_Set() error {
	var err error
	frameTimingsGetPresentationTimeFunction_Once.Do(func() {
		err = frameTimingsStruct_Set()
		if err != nil {
			return
		}
		frameTimingsGetPresentationTimeFunction, err = frameTimingsStruct.InvokerNew("get_presentation_time")
	})
	return err
}

// GetPresentationTime is a representation of the C type gdk_frame_timings_get_presentation_time.
func (recv *FrameTimings) GetPresentationTime() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetPresentationTimeFunction_Set()
	if err == nil {
		ret = frameTimingsGetPresentationTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var frameTimingsGetRefreshIntervalFunction *gi.Function
var frameTimingsGetRefreshIntervalFunction_Once sync.Once

func frameTimingsGetRefreshIntervalFunction_Set() error {
	var err error
	frameTimingsGetRefreshIntervalFunction_Once.Do(func() {
		err = frameTimingsStruct_Set()
		if err != nil {
			return
		}
		frameTimingsGetRefreshIntervalFunction, err = frameTimingsStruct.InvokerNew("get_refresh_interval")
	})
	return err
}

// GetRefreshInterval is a representation of the C type gdk_frame_timings_get_refresh_interval.
func (recv *FrameTimings) GetRefreshInterval() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetRefreshIntervalFunction_Set()
	if err == nil {
		ret = frameTimingsGetRefreshIntervalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var frameTimingsRefFunction *gi.Function
var frameTimingsRefFunction_Once sync.Once

func frameTimingsRefFunction_Set() error {
	var err error
	frameTimingsRefFunction_Once.Do(func() {
		err = frameTimingsStruct_Set()
		if err != nil {
			return
		}
		frameTimingsRefFunction, err = frameTimingsStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gdk_frame_timings_ref.
func (recv *FrameTimings) Ref() (*FrameTimings, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsRefFunction_Set()
	if err == nil {
		ret = frameTimingsRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FrameTimings{native: ret.Pointer()}

	return retGo, err
}

var frameTimingsUnrefFunction *gi.Function
var frameTimingsUnrefFunction_Once sync.Once

func frameTimingsUnrefFunction_Set() error {
	var err error
	frameTimingsUnrefFunction_Once.Do(func() {
		err = frameTimingsStruct_Set()
		if err != nil {
			return
		}
		frameTimingsUnrefFunction, err = frameTimingsStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gdk_frame_timings_unref.
func (recv *FrameTimings) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := frameTimingsUnrefFunction_Set()
	if err == nil {
		frameTimingsUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var geometryStruct *gi.Struct
var geometryStruct_Once sync.Once

func geometryStruct_Set() error {
	var err error
	geometryStruct_Once.Do(func() {
		geometryStruct, err = gi.StructNew("Gdk", "Geometry")
	})
	return err
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

func keymapKeyStruct_Set() error {
	var err error
	keymapKeyStruct_Once.Do(func() {
		keymapKeyStruct, err = gi.StructNew("Gdk", "KeymapKey")
	})
	return err
}

type KeymapKey struct {
	native  uintptr
	Keycode uint32
	Group   int32
	Level   int32
}

var monitorClassStruct *gi.Struct
var monitorClassStruct_Once sync.Once

func monitorClassStruct_Set() error {
	var err error
	monitorClassStruct_Once.Do(func() {
		monitorClassStruct, err = gi.StructNew("Gdk", "MonitorClass")
	})
	return err
}

type MonitorClass struct {
	native uintptr
}

var pointStruct *gi.Struct
var pointStruct_Once sync.Once

func pointStruct_Set() error {
	var err error
	pointStruct_Once.Do(func() {
		pointStruct, err = gi.StructNew("Gdk", "Point")
	})
	return err
}

type Point struct {
	native uintptr
	X      int32
	Y      int32
}

var rGBAStruct *gi.Struct
var rGBAStruct_Once sync.Once

func rGBAStruct_Set() error {
	var err error
	rGBAStruct_Once.Do(func() {
		rGBAStruct, err = gi.StructNew("Gdk", "RGBA")
	})
	return err
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

func rGBACopyFunction_Set() error {
	var err error
	rGBACopyFunction_Once.Do(func() {
		err = rGBAStruct_Set()
		if err != nil {
			return
		}
		rGBACopyFunction, err = rGBAStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gdk_rgba_copy.
func (recv *RGBA) Copy() (*RGBA, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rGBACopyFunction_Set()
	if err == nil {
		ret = rGBACopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &RGBA{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'gdk_rgba_equal' : parameter 'p2' of type 'RGBA' not supported

var rGBAFreeFunction *gi.Function
var rGBAFreeFunction_Once sync.Once

func rGBAFreeFunction_Set() error {
	var err error
	rGBAFreeFunction_Once.Do(func() {
		err = rGBAStruct_Set()
		if err != nil {
			return
		}
		rGBAFreeFunction, err = rGBAStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gdk_rgba_free.
func (recv *RGBA) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := rGBAFreeFunction_Set()
	if err == nil {
		rGBAFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var rGBAHashFunction *gi.Function
var rGBAHashFunction_Once sync.Once

func rGBAHashFunction_Set() error {
	var err error
	rGBAHashFunction_Once.Do(func() {
		err = rGBAStruct_Set()
		if err != nil {
			return
		}
		rGBAHashFunction, err = rGBAStruct.InvokerNew("hash")
	})
	return err
}

// Hash is a representation of the C type gdk_rgba_hash.
func (recv *RGBA) Hash() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rGBAHashFunction_Set()
	if err == nil {
		ret = rGBAHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var rGBAParseFunction *gi.Function
var rGBAParseFunction_Once sync.Once

func rGBAParseFunction_Set() error {
	var err error
	rGBAParseFunction_Once.Do(func() {
		err = rGBAStruct_Set()
		if err != nil {
			return
		}
		rGBAParseFunction, err = rGBAStruct.InvokerNew("parse")
	})
	return err
}

// Parse is a representation of the C type gdk_rgba_parse.
func (recv *RGBA) Parse(spec string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(spec)

	var ret gi.Argument

	err := rGBAParseFunction_Set()
	if err == nil {
		ret = rGBAParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var rGBAToStringFunction *gi.Function
var rGBAToStringFunction_Once sync.Once

func rGBAToStringFunction_Set() error {
	var err error
	rGBAToStringFunction_Once.Do(func() {
		err = rGBAStruct_Set()
		if err != nil {
			return
		}
		rGBAToStringFunction, err = rGBAStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type gdk_rgba_to_string.
func (recv *RGBA) ToString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rGBAToStringFunction_Set()
	if err == nil {
		ret = rGBAToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var rectangleStruct *gi.Struct
var rectangleStruct_Once sync.Once

func rectangleStruct_Set() error {
	var err error
	rectangleStruct_Once.Do(func() {
		rectangleStruct, err = gi.StructNew("Gdk", "Rectangle")
	})
	return err
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

func timeCoordStruct_Set() error {
	var err error
	timeCoordStruct_Once.Do(func() {
		timeCoordStruct, err = gi.StructNew("Gdk", "TimeCoord")
	})
	return err
}

type TimeCoord struct {
	native uintptr
	Time   uint32
	// UNSUPPORTED : C value 'axes' : missing Type
}

var windowAttrStruct *gi.Struct
var windowAttrStruct_Once sync.Once

func windowAttrStruct_Set() error {
	var err error
	windowAttrStruct_Once.Do(func() {
		windowAttrStruct, err = gi.StructNew("Gdk", "WindowAttr")
	})
	return err
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
	WmclassName      string
	WmclassClass     string
	OverrideRedirect bool
	// UNSUPPORTED : C value 'type_hint' : no Go type for 'WindowTypeHint'
}

var windowClassStruct *gi.Struct
var windowClassStruct_Once sync.Once

func windowClassStruct_Set() error {
	var err error
	windowClassStruct_Once.Do(func() {
		windowClassStruct, err = gi.StructNew("Gdk", "WindowClass")
	})
	return err
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

func windowRedirectStruct_Set() error {
	var err error
	windowRedirectStruct_Once.Do(func() {
		windowRedirectStruct, err = gi.StructNew("Gdk", "WindowRedirect")
	})
	return err
}

type WindowRedirect struct {
	native uintptr
}
