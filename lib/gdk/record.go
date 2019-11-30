// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
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
func (recv *Atom) Name() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := atomNameFunction_Set()
	if err == nil {
		ret = atomNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// AtomStruct creates an uninitialised Atom.
func AtomStruct() *Atom {
	err := atomStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Atom{native: atomStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAtom)
	return structGo
}
func finalizeAtom(obj *Atom) {
	atomStruct.Free(obj.native)
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
}

// Pixel returns the C field 'pixel'.
func (recv *Color) Pixel() uint32 {
	argValue := gi.FieldGet(colorStruct, recv.native, "pixel")
	value := argValue.Uint32()
	return value
}

// Red returns the C field 'red'.
func (recv *Color) Red() uint16 {
	argValue := gi.FieldGet(colorStruct, recv.native, "red")
	value := argValue.Uint16()
	return value
}

// Green returns the C field 'green'.
func (recv *Color) Green() uint16 {
	argValue := gi.FieldGet(colorStruct, recv.native, "green")
	value := argValue.Uint16()
	return value
}

// Blue returns the C field 'blue'.
func (recv *Color) Blue() uint16 {
	argValue := gi.FieldGet(colorStruct, recv.native, "blue")
	value := argValue.Uint16()
	return value
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
func (recv *Color) Copy() *Color {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorCopyFunction_Set()
	if err == nil {
		ret = colorCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Color{native: ret.Pointer()}

	return retGo
}

var colorEqualFunction *gi.Function
var colorEqualFunction_Once sync.Once

func colorEqualFunction_Set() error {
	var err error
	colorEqualFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorEqualFunction, err = colorStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type gdk_color_equal.
func (recv *Color) Equal(colorb *Color) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(colorb.native)

	var ret gi.Argument

	err := colorEqualFunction_Set()
	if err == nil {
		ret = colorEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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
func (recv *Color) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := colorFreeFunction_Set()
	if err == nil {
		colorFreeFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *Color) Hash() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorHashFunction_Set()
	if err == nil {
		ret = colorHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
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
func (recv *Color) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorToStringFunction_Set()
	if err == nil {
		ret = colorToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// ColorStruct creates an uninitialised Color.
func ColorStruct() *Color {
	err := colorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Color{native: colorStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColor)
	return structGo
}
func finalizeColor(obj *Color) {
	colorStruct.Free(obj.native)
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

// DevicePadInterfaceStruct creates an uninitialised DevicePadInterface.
func DevicePadInterfaceStruct() *DevicePadInterface {
	err := devicePadInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DevicePadInterface{native: devicePadInterfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeDevicePadInterface)
	return structGo
}
func finalizeDevicePadInterface(obj *DevicePadInterface) {
	devicePadInterfaceStruct.Free(obj.native)
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

// DrawingContextClassStruct creates an uninitialised DrawingContextClass.
func DrawingContextClassStruct() *DrawingContextClass {
	err := drawingContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DrawingContextClass{native: drawingContextClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeDrawingContextClass)
	return structGo
}
func finalizeDrawingContextClass(obj *DrawingContextClass) {
	drawingContextClassStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventAny) SendEvent() int8 {
	argValue := gi.FieldGet(eventAnyStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// EventAnyStruct creates an uninitialised EventAny.
func EventAnyStruct() *EventAny {
	err := eventAnyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventAny{native: eventAnyStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventAny)
	return structGo
}
func finalizeEventAny(obj *EventAny) {
	eventAnyStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventButton) SendEvent() int8 {
	argValue := gi.FieldGet(eventButtonStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventButton) Time() uint32 {
	argValue := gi.FieldGet(eventButtonStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// X returns the C field 'x'.
func (recv *EventButton) X() float64 {
	argValue := gi.FieldGet(eventButtonStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// Y returns the C field 'y'.
func (recv *EventButton) Y() float64 {
	argValue := gi.FieldGet(eventButtonStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// Axes returns the C field 'axes'.
func (recv *EventButton) Axes() float64 {
	argValue := gi.FieldGet(eventButtonStruct, recv.native, "axes")
	value := argValue.Float64()
	return value
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// Button returns the C field 'button'.
func (recv *EventButton) Button() uint32 {
	argValue := gi.FieldGet(eventButtonStruct, recv.native, "button")
	value := argValue.Uint32()
	return value
}

// UNSUPPORTED : C value 'device' : for field getter : no Go type for 'Device'

// XRoot returns the C field 'x_root'.
func (recv *EventButton) XRoot() float64 {
	argValue := gi.FieldGet(eventButtonStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// YRoot returns the C field 'y_root'.
func (recv *EventButton) YRoot() float64 {
	argValue := gi.FieldGet(eventButtonStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// EventButtonStruct creates an uninitialised EventButton.
func EventButtonStruct() *EventButton {
	err := eventButtonStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventButton{native: eventButtonStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventButton)
	return structGo
}
func finalizeEventButton(obj *EventButton) {
	eventButtonStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventConfigure) SendEvent() int8 {
	argValue := gi.FieldGet(eventConfigureStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// X returns the C field 'x'.
func (recv *EventConfigure) X() int32 {
	argValue := gi.FieldGet(eventConfigureStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// Y returns the C field 'y'.
func (recv *EventConfigure) Y() int32 {
	argValue := gi.FieldGet(eventConfigureStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// Width returns the C field 'width'.
func (recv *EventConfigure) Width() int32 {
	argValue := gi.FieldGet(eventConfigureStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// Height returns the C field 'height'.
func (recv *EventConfigure) Height() int32 {
	argValue := gi.FieldGet(eventConfigureStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

// EventConfigureStruct creates an uninitialised EventConfigure.
func EventConfigureStruct() *EventConfigure {
	err := eventConfigureStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventConfigure{native: eventConfigureStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventConfigure)
	return structGo
}
func finalizeEventConfigure(obj *EventConfigure) {
	eventConfigureStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventCrossing) SendEvent() int8 {
	argValue := gi.FieldGet(eventCrossingStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// UNSUPPORTED : C value 'subwindow' : for field getter : no Go type for 'Window'

// Time returns the C field 'time'.
func (recv *EventCrossing) Time() uint32 {
	argValue := gi.FieldGet(eventCrossingStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// X returns the C field 'x'.
func (recv *EventCrossing) X() float64 {
	argValue := gi.FieldGet(eventCrossingStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// Y returns the C field 'y'.
func (recv *EventCrossing) Y() float64 {
	argValue := gi.FieldGet(eventCrossingStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// XRoot returns the C field 'x_root'.
func (recv *EventCrossing) XRoot() float64 {
	argValue := gi.FieldGet(eventCrossingStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// YRoot returns the C field 'y_root'.
func (recv *EventCrossing) YRoot() float64 {
	argValue := gi.FieldGet(eventCrossingStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// UNSUPPORTED : C value 'mode' : for field getter : no Go type for 'CrossingMode'

// UNSUPPORTED : C value 'detail' : for field getter : no Go type for 'NotifyType'

// Focus returns the C field 'focus'.
func (recv *EventCrossing) Focus() bool {
	argValue := gi.FieldGet(eventCrossingStruct, recv.native, "focus")
	value := argValue.Boolean()
	return value
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// EventCrossingStruct creates an uninitialised EventCrossing.
func EventCrossingStruct() *EventCrossing {
	err := eventCrossingStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventCrossing{native: eventCrossingStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventCrossing)
	return structGo
}
func finalizeEventCrossing(obj *EventCrossing) {
	eventCrossingStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventDND) SendEvent() int8 {
	argValue := gi.FieldGet(eventDNDStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// UNSUPPORTED : C value 'context' : for field getter : no Go type for 'DragContext'

// Time returns the C field 'time'.
func (recv *EventDND) Time() uint32 {
	argValue := gi.FieldGet(eventDNDStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// XRoot returns the C field 'x_root'.
func (recv *EventDND) XRoot() int16 {
	argValue := gi.FieldGet(eventDNDStruct, recv.native, "x_root")
	value := argValue.Int16()
	return value
}

// YRoot returns the C field 'y_root'.
func (recv *EventDND) YRoot() int16 {
	argValue := gi.FieldGet(eventDNDStruct, recv.native, "y_root")
	value := argValue.Int16()
	return value
}

// EventDNDStruct creates an uninitialised EventDND.
func EventDNDStruct() *EventDND {
	err := eventDNDStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventDND{native: eventDNDStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventDND)
	return structGo
}
func finalizeEventDND(obj *EventDND) {
	eventDNDStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventExpose) SendEvent() int8 {
	argValue := gi.FieldGet(eventExposeStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Area returns the C field 'area'.
func (recv *EventExpose) Area() *Rectangle {
	argValue := gi.FieldGet(eventExposeStruct, recv.native, "area")
	value := &Rectangle{native: argValue.Pointer()}
	return value
}

// UNSUPPORTED : C value 'region' : for field getter : no Go type for 'cairo.Region'

// Count returns the C field 'count'.
func (recv *EventExpose) Count() int32 {
	argValue := gi.FieldGet(eventExposeStruct, recv.native, "count")
	value := argValue.Int32()
	return value
}

// EventExposeStruct creates an uninitialised EventExpose.
func EventExposeStruct() *EventExpose {
	err := eventExposeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventExpose{native: eventExposeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventExpose)
	return structGo
}
func finalizeEventExpose(obj *EventExpose) {
	eventExposeStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventFocus) SendEvent() int8 {
	argValue := gi.FieldGet(eventFocusStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// In returns the C field 'in'.
func (recv *EventFocus) In() int16 {
	argValue := gi.FieldGet(eventFocusStruct, recv.native, "in")
	value := argValue.Int16()
	return value
}

// EventFocusStruct creates an uninitialised EventFocus.
func EventFocusStruct() *EventFocus {
	err := eventFocusStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventFocus{native: eventFocusStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventFocus)
	return structGo
}
func finalizeEventFocus(obj *EventFocus) {
	eventFocusStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventGrabBroken) SendEvent() int8 {
	argValue := gi.FieldGet(eventGrabBrokenStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Keyboard returns the C field 'keyboard'.
func (recv *EventGrabBroken) Keyboard() bool {
	argValue := gi.FieldGet(eventGrabBrokenStruct, recv.native, "keyboard")
	value := argValue.Boolean()
	return value
}

// Implicit returns the C field 'implicit'.
func (recv *EventGrabBroken) Implicit() bool {
	argValue := gi.FieldGet(eventGrabBrokenStruct, recv.native, "implicit")
	value := argValue.Boolean()
	return value
}

// UNSUPPORTED : C value 'grab_window' : for field getter : no Go type for 'Window'

// EventGrabBrokenStruct creates an uninitialised EventGrabBroken.
func EventGrabBrokenStruct() *EventGrabBroken {
	err := eventGrabBrokenStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventGrabBroken{native: eventGrabBrokenStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventGrabBroken)
	return structGo
}
func finalizeEventGrabBroken(obj *EventGrabBroken) {
	eventGrabBrokenStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventKey) SendEvent() int8 {
	argValue := gi.FieldGet(eventKeyStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventKey) Time() uint32 {
	argValue := gi.FieldGet(eventKeyStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// Keyval returns the C field 'keyval'.
func (recv *EventKey) Keyval() uint32 {
	argValue := gi.FieldGet(eventKeyStruct, recv.native, "keyval")
	value := argValue.Uint32()
	return value
}

// Length returns the C field 'length'.
func (recv *EventKey) Length() int32 {
	argValue := gi.FieldGet(eventKeyStruct, recv.native, "length")
	value := argValue.Int32()
	return value
}

// String returns the C field 'string'.
func (recv *EventKey) String() string {
	argValue := gi.FieldGet(eventKeyStruct, recv.native, "string")
	value := argValue.String(false)
	return value
}

// HardwareKeycode returns the C field 'hardware_keycode'.
func (recv *EventKey) HardwareKeycode() uint16 {
	argValue := gi.FieldGet(eventKeyStruct, recv.native, "hardware_keycode")
	value := argValue.Uint16()
	return value
}

// Group returns the C field 'group'.
func (recv *EventKey) Group() uint8 {
	argValue := gi.FieldGet(eventKeyStruct, recv.native, "group")
	value := argValue.Uint8()
	return value
}

// IsModifier returns the C field 'is_modifier'.
func (recv *EventKey) IsModifier() uint32 {
	argValue := gi.FieldGet(eventKeyStruct, recv.native, "is_modifier")
	value := argValue.Uint32()
	return value
}

// EventKeyStruct creates an uninitialised EventKey.
func EventKeyStruct() *EventKey {
	err := eventKeyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventKey{native: eventKeyStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventKey)
	return structGo
}
func finalizeEventKey(obj *EventKey) {
	eventKeyStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventMotion) SendEvent() int8 {
	argValue := gi.FieldGet(eventMotionStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventMotion) Time() uint32 {
	argValue := gi.FieldGet(eventMotionStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// X returns the C field 'x'.
func (recv *EventMotion) X() float64 {
	argValue := gi.FieldGet(eventMotionStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// Y returns the C field 'y'.
func (recv *EventMotion) Y() float64 {
	argValue := gi.FieldGet(eventMotionStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// Axes returns the C field 'axes'.
func (recv *EventMotion) Axes() float64 {
	argValue := gi.FieldGet(eventMotionStruct, recv.native, "axes")
	value := argValue.Float64()
	return value
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// IsHint returns the C field 'is_hint'.
func (recv *EventMotion) IsHint() int16 {
	argValue := gi.FieldGet(eventMotionStruct, recv.native, "is_hint")
	value := argValue.Int16()
	return value
}

// UNSUPPORTED : C value 'device' : for field getter : no Go type for 'Device'

// XRoot returns the C field 'x_root'.
func (recv *EventMotion) XRoot() float64 {
	argValue := gi.FieldGet(eventMotionStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// YRoot returns the C field 'y_root'.
func (recv *EventMotion) YRoot() float64 {
	argValue := gi.FieldGet(eventMotionStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// EventMotionStruct creates an uninitialised EventMotion.
func EventMotionStruct() *EventMotion {
	err := eventMotionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventMotion{native: eventMotionStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventMotion)
	return structGo
}
func finalizeEventMotion(obj *EventMotion) {
	eventMotionStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventOwnerChange) SendEvent() int8 {
	argValue := gi.FieldGet(eventOwnerChangeStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// UNSUPPORTED : C value 'owner' : for field getter : no Go type for 'Window'

// UNSUPPORTED : C value 'reason' : for field getter : no Go type for 'OwnerChange'

// Selection returns the C field 'selection'.
func (recv *EventOwnerChange) Selection() *Atom {
	argValue := gi.FieldGet(eventOwnerChangeStruct, recv.native, "selection")
	value := &Atom{native: argValue.Pointer()}
	return value
}

// Time returns the C field 'time'.
func (recv *EventOwnerChange) Time() uint32 {
	argValue := gi.FieldGet(eventOwnerChangeStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SelectionTime returns the C field 'selection_time'.
func (recv *EventOwnerChange) SelectionTime() uint32 {
	argValue := gi.FieldGet(eventOwnerChangeStruct, recv.native, "selection_time")
	value := argValue.Uint32()
	return value
}

// EventOwnerChangeStruct creates an uninitialised EventOwnerChange.
func EventOwnerChangeStruct() *EventOwnerChange {
	err := eventOwnerChangeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventOwnerChange{native: eventOwnerChangeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventOwnerChange)
	return structGo
}
func finalizeEventOwnerChange(obj *EventOwnerChange) {
	eventOwnerChangeStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventPadAxis) SendEvent() int8 {
	argValue := gi.FieldGet(eventPadAxisStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventPadAxis) Time() uint32 {
	argValue := gi.FieldGet(eventPadAxisStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// Group returns the C field 'group'.
func (recv *EventPadAxis) Group() uint32 {
	argValue := gi.FieldGet(eventPadAxisStruct, recv.native, "group")
	value := argValue.Uint32()
	return value
}

// Index returns the C field 'index'.
func (recv *EventPadAxis) Index() uint32 {
	argValue := gi.FieldGet(eventPadAxisStruct, recv.native, "index")
	value := argValue.Uint32()
	return value
}

// Mode returns the C field 'mode'.
func (recv *EventPadAxis) Mode() uint32 {
	argValue := gi.FieldGet(eventPadAxisStruct, recv.native, "mode")
	value := argValue.Uint32()
	return value
}

// Value returns the C field 'value'.
func (recv *EventPadAxis) Value() float64 {
	argValue := gi.FieldGet(eventPadAxisStruct, recv.native, "value")
	value := argValue.Float64()
	return value
}

// EventPadAxisStruct creates an uninitialised EventPadAxis.
func EventPadAxisStruct() *EventPadAxis {
	err := eventPadAxisStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventPadAxis{native: eventPadAxisStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventPadAxis)
	return structGo
}
func finalizeEventPadAxis(obj *EventPadAxis) {
	eventPadAxisStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventPadButton) SendEvent() int8 {
	argValue := gi.FieldGet(eventPadButtonStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventPadButton) Time() uint32 {
	argValue := gi.FieldGet(eventPadButtonStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// Group returns the C field 'group'.
func (recv *EventPadButton) Group() uint32 {
	argValue := gi.FieldGet(eventPadButtonStruct, recv.native, "group")
	value := argValue.Uint32()
	return value
}

// Button returns the C field 'button'.
func (recv *EventPadButton) Button() uint32 {
	argValue := gi.FieldGet(eventPadButtonStruct, recv.native, "button")
	value := argValue.Uint32()
	return value
}

// Mode returns the C field 'mode'.
func (recv *EventPadButton) Mode() uint32 {
	argValue := gi.FieldGet(eventPadButtonStruct, recv.native, "mode")
	value := argValue.Uint32()
	return value
}

// EventPadButtonStruct creates an uninitialised EventPadButton.
func EventPadButtonStruct() *EventPadButton {
	err := eventPadButtonStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventPadButton{native: eventPadButtonStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventPadButton)
	return structGo
}
func finalizeEventPadButton(obj *EventPadButton) {
	eventPadButtonStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventPadGroupMode) SendEvent() int8 {
	argValue := gi.FieldGet(eventPadGroupModeStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventPadGroupMode) Time() uint32 {
	argValue := gi.FieldGet(eventPadGroupModeStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// Group returns the C field 'group'.
func (recv *EventPadGroupMode) Group() uint32 {
	argValue := gi.FieldGet(eventPadGroupModeStruct, recv.native, "group")
	value := argValue.Uint32()
	return value
}

// Mode returns the C field 'mode'.
func (recv *EventPadGroupMode) Mode() uint32 {
	argValue := gi.FieldGet(eventPadGroupModeStruct, recv.native, "mode")
	value := argValue.Uint32()
	return value
}

// EventPadGroupModeStruct creates an uninitialised EventPadGroupMode.
func EventPadGroupModeStruct() *EventPadGroupMode {
	err := eventPadGroupModeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventPadGroupMode{native: eventPadGroupModeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventPadGroupMode)
	return structGo
}
func finalizeEventPadGroupMode(obj *EventPadGroupMode) {
	eventPadGroupModeStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventProperty) SendEvent() int8 {
	argValue := gi.FieldGet(eventPropertyStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Atom returns the C field 'atom'.
func (recv *EventProperty) Atom() *Atom {
	argValue := gi.FieldGet(eventPropertyStruct, recv.native, "atom")
	value := &Atom{native: argValue.Pointer()}
	return value
}

// Time returns the C field 'time'.
func (recv *EventProperty) Time() uint32 {
	argValue := gi.FieldGet(eventPropertyStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'PropertyState'

// EventPropertyStruct creates an uninitialised EventProperty.
func EventPropertyStruct() *EventProperty {
	err := eventPropertyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventProperty{native: eventPropertyStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventProperty)
	return structGo
}
func finalizeEventProperty(obj *EventProperty) {
	eventPropertyStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventProximity) SendEvent() int8 {
	argValue := gi.FieldGet(eventProximityStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventProximity) Time() uint32 {
	argValue := gi.FieldGet(eventProximityStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// UNSUPPORTED : C value 'device' : for field getter : no Go type for 'Device'

// EventProximityStruct creates an uninitialised EventProximity.
func EventProximityStruct() *EventProximity {
	err := eventProximityStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventProximity{native: eventProximityStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventProximity)
	return structGo
}
func finalizeEventProximity(obj *EventProximity) {
	eventProximityStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventScroll) SendEvent() int8 {
	argValue := gi.FieldGet(eventScrollStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventScroll) Time() uint32 {
	argValue := gi.FieldGet(eventScrollStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// X returns the C field 'x'.
func (recv *EventScroll) X() float64 {
	argValue := gi.FieldGet(eventScrollStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// Y returns the C field 'y'.
func (recv *EventScroll) Y() float64 {
	argValue := gi.FieldGet(eventScrollStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// UNSUPPORTED : C value 'direction' : for field getter : no Go type for 'ScrollDirection'

// UNSUPPORTED : C value 'device' : for field getter : no Go type for 'Device'

// XRoot returns the C field 'x_root'.
func (recv *EventScroll) XRoot() float64 {
	argValue := gi.FieldGet(eventScrollStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// YRoot returns the C field 'y_root'.
func (recv *EventScroll) YRoot() float64 {
	argValue := gi.FieldGet(eventScrollStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// DeltaX returns the C field 'delta_x'.
func (recv *EventScroll) DeltaX() float64 {
	argValue := gi.FieldGet(eventScrollStruct, recv.native, "delta_x")
	value := argValue.Float64()
	return value
}

// DeltaY returns the C field 'delta_y'.
func (recv *EventScroll) DeltaY() float64 {
	argValue := gi.FieldGet(eventScrollStruct, recv.native, "delta_y")
	value := argValue.Float64()
	return value
}

// IsStop returns the C field 'is_stop'.
func (recv *EventScroll) IsStop() uint32 {
	argValue := gi.FieldGet(eventScrollStruct, recv.native, "is_stop")
	value := argValue.Uint32()
	return value
}

// EventScrollStruct creates an uninitialised EventScroll.
func EventScrollStruct() *EventScroll {
	err := eventScrollStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventScroll{native: eventScrollStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventScroll)
	return structGo
}
func finalizeEventScroll(obj *EventScroll) {
	eventScrollStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventSelection) SendEvent() int8 {
	argValue := gi.FieldGet(eventSelectionStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Selection returns the C field 'selection'.
func (recv *EventSelection) Selection() *Atom {
	argValue := gi.FieldGet(eventSelectionStruct, recv.native, "selection")
	value := &Atom{native: argValue.Pointer()}
	return value
}

// Target returns the C field 'target'.
func (recv *EventSelection) Target() *Atom {
	argValue := gi.FieldGet(eventSelectionStruct, recv.native, "target")
	value := &Atom{native: argValue.Pointer()}
	return value
}

// Property returns the C field 'property'.
func (recv *EventSelection) Property() *Atom {
	argValue := gi.FieldGet(eventSelectionStruct, recv.native, "property")
	value := &Atom{native: argValue.Pointer()}
	return value
}

// Time returns the C field 'time'.
func (recv *EventSelection) Time() uint32 {
	argValue := gi.FieldGet(eventSelectionStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// UNSUPPORTED : C value 'requestor' : for field getter : no Go type for 'Window'

// EventSelectionStruct creates an uninitialised EventSelection.
func EventSelectionStruct() *EventSelection {
	err := eventSelectionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventSelection{native: eventSelectionStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventSelection)
	return structGo
}
func finalizeEventSelection(obj *EventSelection) {
	eventSelectionStruct.Free(obj.native)
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

// EventSequenceStruct creates an uninitialised EventSequence.
func EventSequenceStruct() *EventSequence {
	err := eventSequenceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventSequence{native: eventSequenceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventSequence)
	return structGo
}
func finalizeEventSequence(obj *EventSequence) {
	eventSequenceStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventSetting) SendEvent() int8 {
	argValue := gi.FieldGet(eventSettingStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// UNSUPPORTED : C value 'action' : for field getter : no Go type for 'SettingAction'

// Name returns the C field 'name'.
func (recv *EventSetting) Name() string {
	argValue := gi.FieldGet(eventSettingStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// EventSettingStruct creates an uninitialised EventSetting.
func EventSettingStruct() *EventSetting {
	err := eventSettingStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventSetting{native: eventSettingStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventSetting)
	return structGo
}
func finalizeEventSetting(obj *EventSetting) {
	eventSettingStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventTouch) SendEvent() int8 {
	argValue := gi.FieldGet(eventTouchStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventTouch) Time() uint32 {
	argValue := gi.FieldGet(eventTouchStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// X returns the C field 'x'.
func (recv *EventTouch) X() float64 {
	argValue := gi.FieldGet(eventTouchStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// Y returns the C field 'y'.
func (recv *EventTouch) Y() float64 {
	argValue := gi.FieldGet(eventTouchStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// Axes returns the C field 'axes'.
func (recv *EventTouch) Axes() float64 {
	argValue := gi.FieldGet(eventTouchStruct, recv.native, "axes")
	value := argValue.Float64()
	return value
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// Sequence returns the C field 'sequence'.
func (recv *EventTouch) Sequence() *EventSequence {
	argValue := gi.FieldGet(eventTouchStruct, recv.native, "sequence")
	value := &EventSequence{native: argValue.Pointer()}
	return value
}

// EmulatingPointer returns the C field 'emulating_pointer'.
func (recv *EventTouch) EmulatingPointer() bool {
	argValue := gi.FieldGet(eventTouchStruct, recv.native, "emulating_pointer")
	value := argValue.Boolean()
	return value
}

// UNSUPPORTED : C value 'device' : for field getter : no Go type for 'Device'

// XRoot returns the C field 'x_root'.
func (recv *EventTouch) XRoot() float64 {
	argValue := gi.FieldGet(eventTouchStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// YRoot returns the C field 'y_root'.
func (recv *EventTouch) YRoot() float64 {
	argValue := gi.FieldGet(eventTouchStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// EventTouchStruct creates an uninitialised EventTouch.
func EventTouchStruct() *EventTouch {
	err := eventTouchStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventTouch{native: eventTouchStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventTouch)
	return structGo
}
func finalizeEventTouch(obj *EventTouch) {
	eventTouchStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventTouchpadPinch) SendEvent() int8 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Phase returns the C field 'phase'.
func (recv *EventTouchpadPinch) Phase() int8 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "phase")
	value := argValue.Int8()
	return value
}

// NFingers returns the C field 'n_fingers'.
func (recv *EventTouchpadPinch) NFingers() int8 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "n_fingers")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventTouchpadPinch) Time() uint32 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// X returns the C field 'x'.
func (recv *EventTouchpadPinch) X() float64 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// Y returns the C field 'y'.
func (recv *EventTouchpadPinch) Y() float64 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// Dx returns the C field 'dx'.
func (recv *EventTouchpadPinch) Dx() float64 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "dx")
	value := argValue.Float64()
	return value
}

// Dy returns the C field 'dy'.
func (recv *EventTouchpadPinch) Dy() float64 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "dy")
	value := argValue.Float64()
	return value
}

// AngleDelta returns the C field 'angle_delta'.
func (recv *EventTouchpadPinch) AngleDelta() float64 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "angle_delta")
	value := argValue.Float64()
	return value
}

// Scale returns the C field 'scale'.
func (recv *EventTouchpadPinch) Scale() float64 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "scale")
	value := argValue.Float64()
	return value
}

// XRoot returns the C field 'x_root'.
func (recv *EventTouchpadPinch) XRoot() float64 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// YRoot returns the C field 'y_root'.
func (recv *EventTouchpadPinch) YRoot() float64 {
	argValue := gi.FieldGet(eventTouchpadPinchStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// EventTouchpadPinchStruct creates an uninitialised EventTouchpadPinch.
func EventTouchpadPinchStruct() *EventTouchpadPinch {
	err := eventTouchpadPinchStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventTouchpadPinch{native: eventTouchpadPinchStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventTouchpadPinch)
	return structGo
}
func finalizeEventTouchpadPinch(obj *EventTouchpadPinch) {
	eventTouchpadPinchStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventTouchpadSwipe) SendEvent() int8 {
	argValue := gi.FieldGet(eventTouchpadSwipeStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// Phase returns the C field 'phase'.
func (recv *EventTouchpadSwipe) Phase() int8 {
	argValue := gi.FieldGet(eventTouchpadSwipeStruct, recv.native, "phase")
	value := argValue.Int8()
	return value
}

// NFingers returns the C field 'n_fingers'.
func (recv *EventTouchpadSwipe) NFingers() int8 {
	argValue := gi.FieldGet(eventTouchpadSwipeStruct, recv.native, "n_fingers")
	value := argValue.Int8()
	return value
}

// Time returns the C field 'time'.
func (recv *EventTouchpadSwipe) Time() uint32 {
	argValue := gi.FieldGet(eventTouchpadSwipeStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// X returns the C field 'x'.
func (recv *EventTouchpadSwipe) X() float64 {
	argValue := gi.FieldGet(eventTouchpadSwipeStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// Y returns the C field 'y'.
func (recv *EventTouchpadSwipe) Y() float64 {
	argValue := gi.FieldGet(eventTouchpadSwipeStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// Dx returns the C field 'dx'.
func (recv *EventTouchpadSwipe) Dx() float64 {
	argValue := gi.FieldGet(eventTouchpadSwipeStruct, recv.native, "dx")
	value := argValue.Float64()
	return value
}

// Dy returns the C field 'dy'.
func (recv *EventTouchpadSwipe) Dy() float64 {
	argValue := gi.FieldGet(eventTouchpadSwipeStruct, recv.native, "dy")
	value := argValue.Float64()
	return value
}

// XRoot returns the C field 'x_root'.
func (recv *EventTouchpadSwipe) XRoot() float64 {
	argValue := gi.FieldGet(eventTouchpadSwipeStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// YRoot returns the C field 'y_root'.
func (recv *EventTouchpadSwipe) YRoot() float64 {
	argValue := gi.FieldGet(eventTouchpadSwipeStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// EventTouchpadSwipeStruct creates an uninitialised EventTouchpadSwipe.
func EventTouchpadSwipeStruct() *EventTouchpadSwipe {
	err := eventTouchpadSwipeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventTouchpadSwipe{native: eventTouchpadSwipeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventTouchpadSwipe)
	return structGo
}
func finalizeEventTouchpadSwipe(obj *EventTouchpadSwipe) {
	eventTouchpadSwipeStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventVisibility) SendEvent() int8 {
	argValue := gi.FieldGet(eventVisibilityStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'VisibilityState'

// EventVisibilityStruct creates an uninitialised EventVisibility.
func EventVisibilityStruct() *EventVisibility {
	err := eventVisibilityStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventVisibility{native: eventVisibilityStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventVisibility)
	return structGo
}
func finalizeEventVisibility(obj *EventVisibility) {
	eventVisibilityStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'EventType'

// UNSUPPORTED : C value 'window' : for field getter : no Go type for 'Window'

// SendEvent returns the C field 'send_event'.
func (recv *EventWindowState) SendEvent() int8 {
	argValue := gi.FieldGet(eventWindowStateStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// UNSUPPORTED : C value 'changed_mask' : for field getter : no Go type for 'WindowState'

// UNSUPPORTED : C value 'new_window_state' : for field getter : no Go type for 'WindowState'

// EventWindowStateStruct creates an uninitialised EventWindowState.
func EventWindowStateStruct() *EventWindowState {
	err := eventWindowStateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventWindowState{native: eventWindowStateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventWindowState)
	return structGo
}
func finalizeEventWindowState(obj *EventWindowState) {
	eventWindowStateStruct.Free(obj.native)
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

// FrameClockClassStruct creates an uninitialised FrameClockClass.
func FrameClockClassStruct() *FrameClockClass {
	err := frameClockClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FrameClockClass{native: frameClockClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFrameClockClass)
	return structGo
}
func finalizeFrameClockClass(obj *FrameClockClass) {
	frameClockClassStruct.Free(obj.native)
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

// FrameClockPrivateStruct creates an uninitialised FrameClockPrivate.
func FrameClockPrivateStruct() *FrameClockPrivate {
	err := frameClockPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FrameClockPrivate{native: frameClockPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFrameClockPrivate)
	return structGo
}
func finalizeFrameClockPrivate(obj *FrameClockPrivate) {
	frameClockPrivateStruct.Free(obj.native)
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
func (recv *FrameTimings) GetComplete() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetCompleteFunction_Set()
	if err == nil {
		ret = frameTimingsGetCompleteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func (recv *FrameTimings) GetFrameCounter() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetFrameCounterFunction_Set()
	if err == nil {
		ret = frameTimingsGetFrameCounterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
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
func (recv *FrameTimings) GetFrameTime() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetFrameTimeFunction_Set()
	if err == nil {
		ret = frameTimingsGetFrameTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
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
func (recv *FrameTimings) GetPredictedPresentationTime() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetPredictedPresentationTimeFunction_Set()
	if err == nil {
		ret = frameTimingsGetPredictedPresentationTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
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
func (recv *FrameTimings) GetPresentationTime() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetPresentationTimeFunction_Set()
	if err == nil {
		ret = frameTimingsGetPresentationTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
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
func (recv *FrameTimings) GetRefreshInterval() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsGetRefreshIntervalFunction_Set()
	if err == nil {
		ret = frameTimingsGetRefreshIntervalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
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
func (recv *FrameTimings) Ref() *FrameTimings {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := frameTimingsRefFunction_Set()
	if err == nil {
		ret = frameTimingsRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FrameTimings{native: ret.Pointer()}

	return retGo
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
func (recv *FrameTimings) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := frameTimingsUnrefFunction_Set()
	if err == nil {
		frameTimingsUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// FrameTimingsStruct creates an uninitialised FrameTimings.
func FrameTimingsStruct() *FrameTimings {
	err := frameTimingsStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FrameTimings{native: frameTimingsStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFrameTimings)
	return structGo
}
func finalizeFrameTimings(obj *FrameTimings) {
	frameTimingsStruct.Free(obj.native)
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
	native uintptr
}

// MinWidth returns the C field 'min_width'.
func (recv *Geometry) MinWidth() int32 {
	argValue := gi.FieldGet(geometryStruct, recv.native, "min_width")
	value := argValue.Int32()
	return value
}

// MinHeight returns the C field 'min_height'.
func (recv *Geometry) MinHeight() int32 {
	argValue := gi.FieldGet(geometryStruct, recv.native, "min_height")
	value := argValue.Int32()
	return value
}

// MaxWidth returns the C field 'max_width'.
func (recv *Geometry) MaxWidth() int32 {
	argValue := gi.FieldGet(geometryStruct, recv.native, "max_width")
	value := argValue.Int32()
	return value
}

// MaxHeight returns the C field 'max_height'.
func (recv *Geometry) MaxHeight() int32 {
	argValue := gi.FieldGet(geometryStruct, recv.native, "max_height")
	value := argValue.Int32()
	return value
}

// BaseWidth returns the C field 'base_width'.
func (recv *Geometry) BaseWidth() int32 {
	argValue := gi.FieldGet(geometryStruct, recv.native, "base_width")
	value := argValue.Int32()
	return value
}

// BaseHeight returns the C field 'base_height'.
func (recv *Geometry) BaseHeight() int32 {
	argValue := gi.FieldGet(geometryStruct, recv.native, "base_height")
	value := argValue.Int32()
	return value
}

// WidthInc returns the C field 'width_inc'.
func (recv *Geometry) WidthInc() int32 {
	argValue := gi.FieldGet(geometryStruct, recv.native, "width_inc")
	value := argValue.Int32()
	return value
}

// HeightInc returns the C field 'height_inc'.
func (recv *Geometry) HeightInc() int32 {
	argValue := gi.FieldGet(geometryStruct, recv.native, "height_inc")
	value := argValue.Int32()
	return value
}

// MinAspect returns the C field 'min_aspect'.
func (recv *Geometry) MinAspect() float64 {
	argValue := gi.FieldGet(geometryStruct, recv.native, "min_aspect")
	value := argValue.Float64()
	return value
}

// MaxAspect returns the C field 'max_aspect'.
func (recv *Geometry) MaxAspect() float64 {
	argValue := gi.FieldGet(geometryStruct, recv.native, "max_aspect")
	value := argValue.Float64()
	return value
}

// UNSUPPORTED : C value 'win_gravity' : for field getter : no Go type for 'Gravity'

// GeometryStruct creates an uninitialised Geometry.
func GeometryStruct() *Geometry {
	err := geometryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Geometry{native: geometryStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGeometry)
	return structGo
}
func finalizeGeometry(obj *Geometry) {
	geometryStruct.Free(obj.native)
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
	native uintptr
}

// Keycode returns the C field 'keycode'.
func (recv *KeymapKey) Keycode() uint32 {
	argValue := gi.FieldGet(keymapKeyStruct, recv.native, "keycode")
	value := argValue.Uint32()
	return value
}

// Group returns the C field 'group'.
func (recv *KeymapKey) Group() int32 {
	argValue := gi.FieldGet(keymapKeyStruct, recv.native, "group")
	value := argValue.Int32()
	return value
}

// Level returns the C field 'level'.
func (recv *KeymapKey) Level() int32 {
	argValue := gi.FieldGet(keymapKeyStruct, recv.native, "level")
	value := argValue.Int32()
	return value
}

// KeymapKeyStruct creates an uninitialised KeymapKey.
func KeymapKeyStruct() *KeymapKey {
	err := keymapKeyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &KeymapKey{native: keymapKeyStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeKeymapKey)
	return structGo
}
func finalizeKeymapKey(obj *KeymapKey) {
	keymapKeyStruct.Free(obj.native)
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

// MonitorClassStruct creates an uninitialised MonitorClass.
func MonitorClassStruct() *MonitorClass {
	err := monitorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MonitorClass{native: monitorClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMonitorClass)
	return structGo
}
func finalizeMonitorClass(obj *MonitorClass) {
	monitorClassStruct.Free(obj.native)
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
}

// X returns the C field 'x'.
func (recv *Point) X() int32 {
	argValue := gi.FieldGet(pointStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// Y returns the C field 'y'.
func (recv *Point) Y() int32 {
	argValue := gi.FieldGet(pointStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// PointStruct creates an uninitialised Point.
func PointStruct() *Point {
	err := pointStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Point{native: pointStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePoint)
	return structGo
}
func finalizePoint(obj *Point) {
	pointStruct.Free(obj.native)
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
}

// Red returns the C field 'red'.
func (recv *RGBA) Red() float64 {
	argValue := gi.FieldGet(rGBAStruct, recv.native, "red")
	value := argValue.Float64()
	return value
}

// Green returns the C field 'green'.
func (recv *RGBA) Green() float64 {
	argValue := gi.FieldGet(rGBAStruct, recv.native, "green")
	value := argValue.Float64()
	return value
}

// Blue returns the C field 'blue'.
func (recv *RGBA) Blue() float64 {
	argValue := gi.FieldGet(rGBAStruct, recv.native, "blue")
	value := argValue.Float64()
	return value
}

// Alpha returns the C field 'alpha'.
func (recv *RGBA) Alpha() float64 {
	argValue := gi.FieldGet(rGBAStruct, recv.native, "alpha")
	value := argValue.Float64()
	return value
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
func (recv *RGBA) Copy() *RGBA {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rGBACopyFunction_Set()
	if err == nil {
		ret = rGBACopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &RGBA{native: ret.Pointer()}

	return retGo
}

var rGBAEqualFunction *gi.Function
var rGBAEqualFunction_Once sync.Once

func rGBAEqualFunction_Set() error {
	var err error
	rGBAEqualFunction_Once.Do(func() {
		err = rGBAStruct_Set()
		if err != nil {
			return
		}
		rGBAEqualFunction, err = rGBAStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type gdk_rgba_equal.
func (recv *RGBA) Equal(p2 *RGBA) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(p2.native)

	var ret gi.Argument

	err := rGBAEqualFunction_Set()
	if err == nil {
		ret = rGBAEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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
func (recv *RGBA) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := rGBAFreeFunction_Set()
	if err == nil {
		rGBAFreeFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *RGBA) Hash() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rGBAHashFunction_Set()
	if err == nil {
		ret = rGBAHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
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
func (recv *RGBA) Parse(spec string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(spec)

	var ret gi.Argument

	err := rGBAParseFunction_Set()
	if err == nil {
		ret = rGBAParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func (recv *RGBA) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rGBAToStringFunction_Set()
	if err == nil {
		ret = rGBAToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// RGBAStruct creates an uninitialised RGBA.
func RGBAStruct() *RGBA {
	err := rGBAStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RGBA{native: rGBAStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRGBA)
	return structGo
}
func finalizeRGBA(obj *RGBA) {
	rGBAStruct.Free(obj.native)
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
}

// X returns the C field 'x'.
func (recv *Rectangle) X() int32 {
	argValue := gi.FieldGet(rectangleStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// Y returns the C field 'y'.
func (recv *Rectangle) Y() int32 {
	argValue := gi.FieldGet(rectangleStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// Width returns the C field 'width'.
func (recv *Rectangle) Width() int32 {
	argValue := gi.FieldGet(rectangleStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// Height returns the C field 'height'.
func (recv *Rectangle) Height() int32 {
	argValue := gi.FieldGet(rectangleStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

var rectangleEqualFunction *gi.Function
var rectangleEqualFunction_Once sync.Once

func rectangleEqualFunction_Set() error {
	var err error
	rectangleEqualFunction_Once.Do(func() {
		err = rectangleStruct_Set()
		if err != nil {
			return
		}
		rectangleEqualFunction, err = rectangleStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type gdk_rectangle_equal.
func (recv *Rectangle) Equal(rect2 *Rectangle) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(rect2.native)

	var ret gi.Argument

	err := rectangleEqualFunction_Set()
	if err == nil {
		ret = rectangleEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var rectangleIntersectFunction *gi.Function
var rectangleIntersectFunction_Once sync.Once

func rectangleIntersectFunction_Set() error {
	var err error
	rectangleIntersectFunction_Once.Do(func() {
		err = rectangleStruct_Set()
		if err != nil {
			return
		}
		rectangleIntersectFunction, err = rectangleStruct.InvokerNew("intersect")
	})
	return err
}

// Intersect is a representation of the C type gdk_rectangle_intersect.
func (recv *Rectangle) Intersect(src2 *Rectangle) (bool, *Rectangle) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(src2.native)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := rectangleIntersectFunction_Set()
	if err == nil {
		ret = rectangleIntersectFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := &Rectangle{native: outArgs[0].Pointer()}

	return retGo, out0
}

var rectangleUnionFunction *gi.Function
var rectangleUnionFunction_Once sync.Once

func rectangleUnionFunction_Set() error {
	var err error
	rectangleUnionFunction_Once.Do(func() {
		err = rectangleStruct_Set()
		if err != nil {
			return
		}
		rectangleUnionFunction, err = rectangleStruct.InvokerNew("union")
	})
	return err
}

// Union is a representation of the C type gdk_rectangle_union.
func (recv *Rectangle) Union(src2 *Rectangle) *Rectangle {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(src2.native)

	var outArgs [1]gi.Argument

	err := rectangleUnionFunction_Set()
	if err == nil {
		rectangleUnionFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}

	return out0
}

// RectangleStruct creates an uninitialised Rectangle.
func RectangleStruct() *Rectangle {
	err := rectangleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Rectangle{native: rectangleStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRectangle)
	return structGo
}
func finalizeRectangle(obj *Rectangle) {
	rectangleStruct.Free(obj.native)
}

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
}

// Time returns the C field 'time'.
func (recv *TimeCoord) Time() uint32 {
	argValue := gi.FieldGet(timeCoordStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// UNSUPPORTED : C value 'axes' : for field getter : missing Type

// TimeCoordStruct creates an uninitialised TimeCoord.
func TimeCoordStruct() *TimeCoord {
	err := timeCoordStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TimeCoord{native: timeCoordStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTimeCoord)
	return structGo
}
func finalizeTimeCoord(obj *TimeCoord) {
	timeCoordStruct.Free(obj.native)
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
	native uintptr
}

// Title returns the C field 'title'.
func (recv *WindowAttr) Title() string {
	argValue := gi.FieldGet(windowAttrStruct, recv.native, "title")
	value := argValue.String(false)
	return value
}

// EventMask returns the C field 'event_mask'.
func (recv *WindowAttr) EventMask() int32 {
	argValue := gi.FieldGet(windowAttrStruct, recv.native, "event_mask")
	value := argValue.Int32()
	return value
}

// X returns the C field 'x'.
func (recv *WindowAttr) X() int32 {
	argValue := gi.FieldGet(windowAttrStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// Y returns the C field 'y'.
func (recv *WindowAttr) Y() int32 {
	argValue := gi.FieldGet(windowAttrStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// Width returns the C field 'width'.
func (recv *WindowAttr) Width() int32 {
	argValue := gi.FieldGet(windowAttrStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// Height returns the C field 'height'.
func (recv *WindowAttr) Height() int32 {
	argValue := gi.FieldGet(windowAttrStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

// UNSUPPORTED : C value 'wclass' : for field getter : no Go type for 'WindowWindowClass'

// UNSUPPORTED : C value 'visual' : for field getter : no Go type for 'Visual'

// UNSUPPORTED : C value 'window_type' : for field getter : no Go type for 'WindowType'

// UNSUPPORTED : C value 'cursor' : for field getter : no Go type for 'Cursor'

// WmclassName returns the C field 'wmclass_name'.
func (recv *WindowAttr) WmclassName() string {
	argValue := gi.FieldGet(windowAttrStruct, recv.native, "wmclass_name")
	value := argValue.String(false)
	return value
}

// WmclassClass returns the C field 'wmclass_class'.
func (recv *WindowAttr) WmclassClass() string {
	argValue := gi.FieldGet(windowAttrStruct, recv.native, "wmclass_class")
	value := argValue.String(false)
	return value
}

// OverrideRedirect returns the C field 'override_redirect'.
func (recv *WindowAttr) OverrideRedirect() bool {
	argValue := gi.FieldGet(windowAttrStruct, recv.native, "override_redirect")
	value := argValue.Boolean()
	return value
}

// UNSUPPORTED : C value 'type_hint' : for field getter : no Go type for 'WindowTypeHint'

// WindowAttrStruct creates an uninitialised WindowAttr.
func WindowAttrStruct() *WindowAttr {
	err := windowAttrStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowAttr{native: windowAttrStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowAttr)
	return structGo
}
func finalizeWindowAttr(obj *WindowAttr) {
	windowAttrStruct.Free(obj.native)
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
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'pick_embedded_child' : for field getter : missing Type

// UNSUPPORTED : C value 'to_embedder' : for field getter : missing Type

// UNSUPPORTED : C value 'from_embedder' : for field getter : missing Type

// UNSUPPORTED : C value 'create_surface' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved8' : for field getter : missing Type

// WindowClassStruct creates an uninitialised WindowClass.
func WindowClassStruct() *WindowClass {
	err := windowClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowClass{native: windowClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowClass)
	return structGo
}
func finalizeWindowClass(obj *WindowClass) {
	windowClassStruct.Free(obj.native)
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

// WindowRedirectStruct creates an uninitialised WindowRedirect.
func WindowRedirectStruct() *WindowRedirect {
	err := windowRedirectStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowRedirect{native: windowRedirectStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowRedirect)
	return structGo
}
func finalizeWindowRedirect(obj *WindowRedirect) {
	windowRedirectStruct.Free(obj.native)
}
