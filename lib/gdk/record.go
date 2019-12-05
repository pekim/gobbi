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

func AtomNewFromNative(native uintptr) *Atom {
	return &Atom{native: native}
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

	structGo := AtomNewFromNative(atomStruct.Alloc())
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

func ColorNewFromNative(native uintptr) *Color {
	return &Color{native: native}
}

// FieldPixel returns the C field 'pixel'.
func (recv *Color) FieldPixel() uint32 {
	argValue := gi.StructFieldGet(colorStruct, recv.native, "pixel")
	value := argValue.Uint32()
	return value
}

// SetFieldPixel sets the value of the C field 'pixel'.
func (recv *Color) SetFieldPixel(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(colorStruct, recv.native, "pixel", argValue)
}

// FieldRed returns the C field 'red'.
func (recv *Color) FieldRed() uint16 {
	argValue := gi.StructFieldGet(colorStruct, recv.native, "red")
	value := argValue.Uint16()
	return value
}

// SetFieldRed sets the value of the C field 'red'.
func (recv *Color) SetFieldRed(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(colorStruct, recv.native, "red", argValue)
}

// FieldGreen returns the C field 'green'.
func (recv *Color) FieldGreen() uint16 {
	argValue := gi.StructFieldGet(colorStruct, recv.native, "green")
	value := argValue.Uint16()
	return value
}

// SetFieldGreen sets the value of the C field 'green'.
func (recv *Color) SetFieldGreen(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(colorStruct, recv.native, "green", argValue)
}

// FieldBlue returns the C field 'blue'.
func (recv *Color) FieldBlue() uint16 {
	argValue := gi.StructFieldGet(colorStruct, recv.native, "blue")
	value := argValue.Uint16()
	return value
}

// SetFieldBlue sets the value of the C field 'blue'.
func (recv *Color) SetFieldBlue(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(colorStruct, recv.native, "blue", argValue)
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

	retGo := ColorNewFromNative(ret.Pointer())

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

	structGo := ColorNewFromNative(colorStruct.Alloc())
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

func DevicePadInterfaceNewFromNative(native uintptr) *DevicePadInterface {
	return &DevicePadInterface{native: native}
}

// DevicePadInterfaceStruct creates an uninitialised DevicePadInterface.
func DevicePadInterfaceStruct() *DevicePadInterface {
	err := devicePadInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DevicePadInterfaceNewFromNative(devicePadInterfaceStruct.Alloc())
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

func DrawingContextClassNewFromNative(native uintptr) *DrawingContextClass {
	return &DrawingContextClass{native: native}
}

// DrawingContextClassStruct creates an uninitialised DrawingContextClass.
func DrawingContextClassStruct() *DrawingContextClass {
	err := drawingContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DrawingContextClassNewFromNative(drawingContextClassStruct.Alloc())
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

func EventAnyNewFromNative(native uintptr) *EventAny {
	return &EventAny{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventAny) FieldType() EventType {
	argValue := gi.StructFieldGet(eventAnyStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventAny) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventAnyStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventAny) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventAnyStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventAny) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventAnyStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventAny) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventAnyStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventAny) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventAnyStruct, recv.native, "send_event", argValue)
}

// EventAnyStruct creates an uninitialised EventAny.
func EventAnyStruct() *EventAny {
	err := eventAnyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventAnyNewFromNative(eventAnyStruct.Alloc())
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

func EventButtonNewFromNative(native uintptr) *EventButton {
	return &EventButton{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventButton) FieldType() EventType {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventButton) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventButtonStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventButton) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventButton) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventButtonStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventButton) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventButton) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventButtonStruct, recv.native, "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventButton) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventButton) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventButtonStruct, recv.native, "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventButton) FieldX() float64 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventButton) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventButtonStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventButton) FieldY() float64 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventButton) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventButtonStruct, recv.native, "y", argValue)
}

// FieldAxes returns the C field 'axes'.
func (recv *EventButton) FieldAxes() float64 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "axes")
	value := argValue.Float64()
	return value
}

// SetFieldAxes sets the value of the C field 'axes'.
func (recv *EventButton) SetFieldAxes(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventButtonStruct, recv.native, "axes", argValue)
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// UNSUPPORTED : C value 'state' : for field setter : no Go type for 'ModifierType'

// FieldButton returns the C field 'button'.
func (recv *EventButton) FieldButton() uint32 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "button")
	value := argValue.Uint32()
	return value
}

// SetFieldButton sets the value of the C field 'button'.
func (recv *EventButton) SetFieldButton(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventButtonStruct, recv.native, "button", argValue)
}

// FieldDevice returns the C field 'device'.
func (recv *EventButton) FieldDevice() *Device {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "device")
	value := DeviceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDevice sets the value of the C field 'device'.
func (recv *EventButton) SetFieldDevice(value *Device) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventButtonStruct, recv.native, "device", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventButton) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventButton) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventButtonStruct, recv.native, "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventButton) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventButton) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventButtonStruct, recv.native, "y_root", argValue)
}

// EventButtonStruct creates an uninitialised EventButton.
func EventButtonStruct() *EventButton {
	err := eventButtonStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventButtonNewFromNative(eventButtonStruct.Alloc())
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

func EventConfigureNewFromNative(native uintptr) *EventConfigure {
	return &EventConfigure{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventConfigure) FieldType() EventType {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventConfigure) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventConfigureStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventConfigure) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventConfigure) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventConfigureStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventConfigure) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventConfigure) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventConfigureStruct, recv.native, "send_event", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventConfigure) FieldX() int32 {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventConfigure) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventConfigureStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventConfigure) FieldY() int32 {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventConfigure) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventConfigureStruct, recv.native, "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *EventConfigure) FieldWidth() int32 {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *EventConfigure) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventConfigureStruct, recv.native, "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *EventConfigure) FieldHeight() int32 {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *EventConfigure) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventConfigureStruct, recv.native, "height", argValue)
}

// EventConfigureStruct creates an uninitialised EventConfigure.
func EventConfigureStruct() *EventConfigure {
	err := eventConfigureStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventConfigureNewFromNative(eventConfigureStruct.Alloc())
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

func EventCrossingNewFromNative(native uintptr) *EventCrossing {
	return &EventCrossing{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventCrossing) FieldType() EventType {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventCrossing) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventCrossingStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventCrossing) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventCrossing) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventCrossingStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventCrossing) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventCrossing) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventCrossingStruct, recv.native, "send_event", argValue)
}

// FieldSubwindow returns the C field 'subwindow'.
func (recv *EventCrossing) FieldSubwindow() *Window {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "subwindow")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldSubwindow sets the value of the C field 'subwindow'.
func (recv *EventCrossing) SetFieldSubwindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventCrossingStruct, recv.native, "subwindow", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventCrossing) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventCrossing) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventCrossingStruct, recv.native, "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventCrossing) FieldX() float64 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventCrossing) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventCrossingStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventCrossing) FieldY() float64 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventCrossing) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventCrossingStruct, recv.native, "y", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventCrossing) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventCrossing) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventCrossingStruct, recv.native, "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventCrossing) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventCrossing) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventCrossingStruct, recv.native, "y_root", argValue)
}

// FieldMode returns the C field 'mode'.
func (recv *EventCrossing) FieldMode() CrossingMode {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "mode")
	value := CrossingMode(argValue.Int32())
	return value
}

// SetFieldMode sets the value of the C field 'mode'.
func (recv *EventCrossing) SetFieldMode(value CrossingMode) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventCrossingStruct, recv.native, "mode", argValue)
}

// FieldDetail returns the C field 'detail'.
func (recv *EventCrossing) FieldDetail() NotifyType {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "detail")
	value := NotifyType(argValue.Int32())
	return value
}

// SetFieldDetail sets the value of the C field 'detail'.
func (recv *EventCrossing) SetFieldDetail(value NotifyType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventCrossingStruct, recv.native, "detail", argValue)
}

// FieldFocus returns the C field 'focus'.
func (recv *EventCrossing) FieldFocus() bool {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.native, "focus")
	value := argValue.Boolean()
	return value
}

// SetFieldFocus sets the value of the C field 'focus'.
func (recv *EventCrossing) SetFieldFocus(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(eventCrossingStruct, recv.native, "focus", argValue)
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// UNSUPPORTED : C value 'state' : for field setter : no Go type for 'ModifierType'

// EventCrossingStruct creates an uninitialised EventCrossing.
func EventCrossingStruct() *EventCrossing {
	err := eventCrossingStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventCrossingNewFromNative(eventCrossingStruct.Alloc())
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

func EventDNDNewFromNative(native uintptr) *EventDND {
	return &EventDND{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventDND) FieldType() EventType {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventDND) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventDNDStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventDND) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventDND) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventDNDStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventDND) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventDND) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventDNDStruct, recv.native, "send_event", argValue)
}

// FieldContext returns the C field 'context'.
func (recv *EventDND) FieldContext() *DragContext {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.native, "context")
	value := DragContextNewFromNative(argValue.Pointer())
	return value
}

// SetFieldContext sets the value of the C field 'context'.
func (recv *EventDND) SetFieldContext(value *DragContext) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventDNDStruct, recv.native, "context", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventDND) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventDND) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventDNDStruct, recv.native, "time", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventDND) FieldXRoot() int16 {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.native, "x_root")
	value := argValue.Int16()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventDND) SetFieldXRoot(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.StructFieldSet(eventDNDStruct, recv.native, "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventDND) FieldYRoot() int16 {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.native, "y_root")
	value := argValue.Int16()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventDND) SetFieldYRoot(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.StructFieldSet(eventDNDStruct, recv.native, "y_root", argValue)
}

// EventDNDStruct creates an uninitialised EventDND.
func EventDNDStruct() *EventDND {
	err := eventDNDStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventDNDNewFromNative(eventDNDStruct.Alloc())
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

func EventExposeNewFromNative(native uintptr) *EventExpose {
	return &EventExpose{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventExpose) FieldType() EventType {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventExpose) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventExposeStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventExpose) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventExpose) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventExposeStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventExpose) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventExpose) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventExposeStruct, recv.native, "send_event", argValue)
}

// FieldArea returns the C field 'area'.
func (recv *EventExpose) FieldArea() *Rectangle {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.native, "area")
	value := RectangleNewFromNative(argValue.Pointer())
	return value
}

// SetFieldArea sets the value of the C field 'area'.
func (recv *EventExpose) SetFieldArea(value *Rectangle) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventExposeStruct, recv.native, "area", argValue)
}

// UNSUPPORTED : C value 'region' : for field getter : no Go type for 'cairo.Region'

// UNSUPPORTED : C value 'region' : for field setter : no Go type for 'cairo.Region'

// FieldCount returns the C field 'count'.
func (recv *EventExpose) FieldCount() int32 {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.native, "count")
	value := argValue.Int32()
	return value
}

// SetFieldCount sets the value of the C field 'count'.
func (recv *EventExpose) SetFieldCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventExposeStruct, recv.native, "count", argValue)
}

// EventExposeStruct creates an uninitialised EventExpose.
func EventExposeStruct() *EventExpose {
	err := eventExposeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventExposeNewFromNative(eventExposeStruct.Alloc())
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

func EventFocusNewFromNative(native uintptr) *EventFocus {
	return &EventFocus{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventFocus) FieldType() EventType {
	argValue := gi.StructFieldGet(eventFocusStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventFocus) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventFocusStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventFocus) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventFocusStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventFocus) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventFocusStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventFocus) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventFocusStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventFocus) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventFocusStruct, recv.native, "send_event", argValue)
}

// FieldIn returns the C field 'in'.
func (recv *EventFocus) FieldIn() int16 {
	argValue := gi.StructFieldGet(eventFocusStruct, recv.native, "in")
	value := argValue.Int16()
	return value
}

// SetFieldIn sets the value of the C field 'in'.
func (recv *EventFocus) SetFieldIn(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.StructFieldSet(eventFocusStruct, recv.native, "in", argValue)
}

// EventFocusStruct creates an uninitialised EventFocus.
func EventFocusStruct() *EventFocus {
	err := eventFocusStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventFocusNewFromNative(eventFocusStruct.Alloc())
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

func EventGrabBrokenNewFromNative(native uintptr) *EventGrabBroken {
	return &EventGrabBroken{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventGrabBroken) FieldType() EventType {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventGrabBroken) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventGrabBrokenStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventGrabBroken) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventGrabBroken) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventGrabBrokenStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventGrabBroken) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventGrabBroken) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventGrabBrokenStruct, recv.native, "send_event", argValue)
}

// FieldKeyboard returns the C field 'keyboard'.
func (recv *EventGrabBroken) FieldKeyboard() bool {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.native, "keyboard")
	value := argValue.Boolean()
	return value
}

// SetFieldKeyboard sets the value of the C field 'keyboard'.
func (recv *EventGrabBroken) SetFieldKeyboard(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(eventGrabBrokenStruct, recv.native, "keyboard", argValue)
}

// FieldImplicit returns the C field 'implicit'.
func (recv *EventGrabBroken) FieldImplicit() bool {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.native, "implicit")
	value := argValue.Boolean()
	return value
}

// SetFieldImplicit sets the value of the C field 'implicit'.
func (recv *EventGrabBroken) SetFieldImplicit(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(eventGrabBrokenStruct, recv.native, "implicit", argValue)
}

// FieldGrabWindow returns the C field 'grab_window'.
func (recv *EventGrabBroken) FieldGrabWindow() *Window {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.native, "grab_window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldGrabWindow sets the value of the C field 'grab_window'.
func (recv *EventGrabBroken) SetFieldGrabWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventGrabBrokenStruct, recv.native, "grab_window", argValue)
}

// EventGrabBrokenStruct creates an uninitialised EventGrabBroken.
func EventGrabBrokenStruct() *EventGrabBroken {
	err := eventGrabBrokenStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventGrabBrokenNewFromNative(eventGrabBrokenStruct.Alloc())
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

func EventKeyNewFromNative(native uintptr) *EventKey {
	return &EventKey{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventKey) FieldType() EventType {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventKey) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventKeyStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventKey) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventKey) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventKeyStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventKey) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventKey) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventKeyStruct, recv.native, "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventKey) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventKey) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventKeyStruct, recv.native, "time", argValue)
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// UNSUPPORTED : C value 'state' : for field setter : no Go type for 'ModifierType'

// FieldKeyval returns the C field 'keyval'.
func (recv *EventKey) FieldKeyval() uint32 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.native, "keyval")
	value := argValue.Uint32()
	return value
}

// SetFieldKeyval sets the value of the C field 'keyval'.
func (recv *EventKey) SetFieldKeyval(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventKeyStruct, recv.native, "keyval", argValue)
}

// FieldLength returns the C field 'length'.
func (recv *EventKey) FieldLength() int32 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.native, "length")
	value := argValue.Int32()
	return value
}

// SetFieldLength sets the value of the C field 'length'.
func (recv *EventKey) SetFieldLength(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventKeyStruct, recv.native, "length", argValue)
}

// FieldString returns the C field 'string'.
func (recv *EventKey) FieldString() string {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.native, "string")
	value := argValue.String(false)
	return value
}

// SetFieldString sets the value of the C field 'string'.
func (recv *EventKey) SetFieldString(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(eventKeyStruct, recv.native, "string", argValue)
}

// FieldHardwareKeycode returns the C field 'hardware_keycode'.
func (recv *EventKey) FieldHardwareKeycode() uint16 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.native, "hardware_keycode")
	value := argValue.Uint16()
	return value
}

// SetFieldHardwareKeycode sets the value of the C field 'hardware_keycode'.
func (recv *EventKey) SetFieldHardwareKeycode(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(eventKeyStruct, recv.native, "hardware_keycode", argValue)
}

// FieldGroup returns the C field 'group'.
func (recv *EventKey) FieldGroup() uint8 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.native, "group")
	value := argValue.Uint8()
	return value
}

// SetFieldGroup sets the value of the C field 'group'.
func (recv *EventKey) SetFieldGroup(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.StructFieldSet(eventKeyStruct, recv.native, "group", argValue)
}

// FieldIsModifier returns the C field 'is_modifier'.
func (recv *EventKey) FieldIsModifier() uint32 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.native, "is_modifier")
	value := argValue.Uint32()
	return value
}

// SetFieldIsModifier sets the value of the C field 'is_modifier'.
func (recv *EventKey) SetFieldIsModifier(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventKeyStruct, recv.native, "is_modifier", argValue)
}

// EventKeyStruct creates an uninitialised EventKey.
func EventKeyStruct() *EventKey {
	err := eventKeyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventKeyNewFromNative(eventKeyStruct.Alloc())
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

func EventMotionNewFromNative(native uintptr) *EventMotion {
	return &EventMotion{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventMotion) FieldType() EventType {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventMotion) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventMotionStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventMotion) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventMotion) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventMotionStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventMotion) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventMotion) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventMotionStruct, recv.native, "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventMotion) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventMotion) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventMotionStruct, recv.native, "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventMotion) FieldX() float64 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventMotion) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventMotionStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventMotion) FieldY() float64 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventMotion) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventMotionStruct, recv.native, "y", argValue)
}

// FieldAxes returns the C field 'axes'.
func (recv *EventMotion) FieldAxes() float64 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "axes")
	value := argValue.Float64()
	return value
}

// SetFieldAxes sets the value of the C field 'axes'.
func (recv *EventMotion) SetFieldAxes(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventMotionStruct, recv.native, "axes", argValue)
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// UNSUPPORTED : C value 'state' : for field setter : no Go type for 'ModifierType'

// FieldIsHint returns the C field 'is_hint'.
func (recv *EventMotion) FieldIsHint() int16 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "is_hint")
	value := argValue.Int16()
	return value
}

// SetFieldIsHint sets the value of the C field 'is_hint'.
func (recv *EventMotion) SetFieldIsHint(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.StructFieldSet(eventMotionStruct, recv.native, "is_hint", argValue)
}

// FieldDevice returns the C field 'device'.
func (recv *EventMotion) FieldDevice() *Device {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "device")
	value := DeviceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDevice sets the value of the C field 'device'.
func (recv *EventMotion) SetFieldDevice(value *Device) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventMotionStruct, recv.native, "device", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventMotion) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventMotion) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventMotionStruct, recv.native, "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventMotion) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventMotion) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventMotionStruct, recv.native, "y_root", argValue)
}

// EventMotionStruct creates an uninitialised EventMotion.
func EventMotionStruct() *EventMotion {
	err := eventMotionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventMotionNewFromNative(eventMotionStruct.Alloc())
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

func EventOwnerChangeNewFromNative(native uintptr) *EventOwnerChange {
	return &EventOwnerChange{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventOwnerChange) FieldType() EventType {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventOwnerChange) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventOwnerChangeStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventOwnerChange) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventOwnerChange) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventOwnerChangeStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventOwnerChange) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventOwnerChange) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventOwnerChangeStruct, recv.native, "send_event", argValue)
}

// FieldOwner returns the C field 'owner'.
func (recv *EventOwnerChange) FieldOwner() *Window {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.native, "owner")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldOwner sets the value of the C field 'owner'.
func (recv *EventOwnerChange) SetFieldOwner(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventOwnerChangeStruct, recv.native, "owner", argValue)
}

// FieldReason returns the C field 'reason'.
func (recv *EventOwnerChange) FieldReason() OwnerChange {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.native, "reason")
	value := OwnerChange(argValue.Int32())
	return value
}

// SetFieldReason sets the value of the C field 'reason'.
func (recv *EventOwnerChange) SetFieldReason(value OwnerChange) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventOwnerChangeStruct, recv.native, "reason", argValue)
}

// FieldSelection returns the C field 'selection'.
func (recv *EventOwnerChange) FieldSelection() *Atom {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.native, "selection")
	value := AtomNewFromNative(argValue.Pointer())
	return value
}

// SetFieldSelection sets the value of the C field 'selection'.
func (recv *EventOwnerChange) SetFieldSelection(value *Atom) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventOwnerChangeStruct, recv.native, "selection", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventOwnerChange) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventOwnerChange) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventOwnerChangeStruct, recv.native, "time", argValue)
}

// FieldSelectionTime returns the C field 'selection_time'.
func (recv *EventOwnerChange) FieldSelectionTime() uint32 {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.native, "selection_time")
	value := argValue.Uint32()
	return value
}

// SetFieldSelectionTime sets the value of the C field 'selection_time'.
func (recv *EventOwnerChange) SetFieldSelectionTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventOwnerChangeStruct, recv.native, "selection_time", argValue)
}

// EventOwnerChangeStruct creates an uninitialised EventOwnerChange.
func EventOwnerChangeStruct() *EventOwnerChange {
	err := eventOwnerChangeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventOwnerChangeNewFromNative(eventOwnerChangeStruct.Alloc())
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

func EventPadAxisNewFromNative(native uintptr) *EventPadAxis {
	return &EventPadAxis{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventPadAxis) FieldType() EventType {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventPadAxis) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventPadAxisStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventPadAxis) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventPadAxis) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventPadAxisStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventPadAxis) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventPadAxis) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.native, "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventPadAxis) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventPadAxis) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.native, "time", argValue)
}

// FieldGroup returns the C field 'group'.
func (recv *EventPadAxis) FieldGroup() uint32 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.native, "group")
	value := argValue.Uint32()
	return value
}

// SetFieldGroup sets the value of the C field 'group'.
func (recv *EventPadAxis) SetFieldGroup(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.native, "group", argValue)
}

// FieldIndex returns the C field 'index'.
func (recv *EventPadAxis) FieldIndex() uint32 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.native, "index")
	value := argValue.Uint32()
	return value
}

// SetFieldIndex sets the value of the C field 'index'.
func (recv *EventPadAxis) SetFieldIndex(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.native, "index", argValue)
}

// FieldMode returns the C field 'mode'.
func (recv *EventPadAxis) FieldMode() uint32 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.native, "mode")
	value := argValue.Uint32()
	return value
}

// SetFieldMode sets the value of the C field 'mode'.
func (recv *EventPadAxis) SetFieldMode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.native, "mode", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *EventPadAxis) FieldValue() float64 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.native, "value")
	value := argValue.Float64()
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *EventPadAxis) SetFieldValue(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.native, "value", argValue)
}

// EventPadAxisStruct creates an uninitialised EventPadAxis.
func EventPadAxisStruct() *EventPadAxis {
	err := eventPadAxisStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventPadAxisNewFromNative(eventPadAxisStruct.Alloc())
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

func EventPadButtonNewFromNative(native uintptr) *EventPadButton {
	return &EventPadButton{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventPadButton) FieldType() EventType {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventPadButton) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventPadButtonStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventPadButton) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventPadButton) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventPadButtonStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventPadButton) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventPadButton) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventPadButtonStruct, recv.native, "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventPadButton) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventPadButton) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadButtonStruct, recv.native, "time", argValue)
}

// FieldGroup returns the C field 'group'.
func (recv *EventPadButton) FieldGroup() uint32 {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.native, "group")
	value := argValue.Uint32()
	return value
}

// SetFieldGroup sets the value of the C field 'group'.
func (recv *EventPadButton) SetFieldGroup(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadButtonStruct, recv.native, "group", argValue)
}

// FieldButton returns the C field 'button'.
func (recv *EventPadButton) FieldButton() uint32 {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.native, "button")
	value := argValue.Uint32()
	return value
}

// SetFieldButton sets the value of the C field 'button'.
func (recv *EventPadButton) SetFieldButton(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadButtonStruct, recv.native, "button", argValue)
}

// FieldMode returns the C field 'mode'.
func (recv *EventPadButton) FieldMode() uint32 {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.native, "mode")
	value := argValue.Uint32()
	return value
}

// SetFieldMode sets the value of the C field 'mode'.
func (recv *EventPadButton) SetFieldMode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadButtonStruct, recv.native, "mode", argValue)
}

// EventPadButtonStruct creates an uninitialised EventPadButton.
func EventPadButtonStruct() *EventPadButton {
	err := eventPadButtonStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventPadButtonNewFromNative(eventPadButtonStruct.Alloc())
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

func EventPadGroupModeNewFromNative(native uintptr) *EventPadGroupMode {
	return &EventPadGroupMode{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventPadGroupMode) FieldType() EventType {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventPadGroupMode) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventPadGroupModeStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventPadGroupMode) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventPadGroupMode) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventPadGroupModeStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventPadGroupMode) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventPadGroupMode) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventPadGroupModeStruct, recv.native, "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventPadGroupMode) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventPadGroupMode) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadGroupModeStruct, recv.native, "time", argValue)
}

// FieldGroup returns the C field 'group'.
func (recv *EventPadGroupMode) FieldGroup() uint32 {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.native, "group")
	value := argValue.Uint32()
	return value
}

// SetFieldGroup sets the value of the C field 'group'.
func (recv *EventPadGroupMode) SetFieldGroup(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadGroupModeStruct, recv.native, "group", argValue)
}

// FieldMode returns the C field 'mode'.
func (recv *EventPadGroupMode) FieldMode() uint32 {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.native, "mode")
	value := argValue.Uint32()
	return value
}

// SetFieldMode sets the value of the C field 'mode'.
func (recv *EventPadGroupMode) SetFieldMode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadGroupModeStruct, recv.native, "mode", argValue)
}

// EventPadGroupModeStruct creates an uninitialised EventPadGroupMode.
func EventPadGroupModeStruct() *EventPadGroupMode {
	err := eventPadGroupModeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventPadGroupModeNewFromNative(eventPadGroupModeStruct.Alloc())
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

func EventPropertyNewFromNative(native uintptr) *EventProperty {
	return &EventProperty{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventProperty) FieldType() EventType {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventProperty) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventPropertyStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventProperty) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventProperty) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventPropertyStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventProperty) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventProperty) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventPropertyStruct, recv.native, "send_event", argValue)
}

// FieldAtom returns the C field 'atom'.
func (recv *EventProperty) FieldAtom() *Atom {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.native, "atom")
	value := AtomNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAtom sets the value of the C field 'atom'.
func (recv *EventProperty) SetFieldAtom(value *Atom) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventPropertyStruct, recv.native, "atom", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventProperty) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventProperty) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPropertyStruct, recv.native, "time", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventProperty) FieldState() PropertyState {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.native, "state")
	value := PropertyState(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventProperty) SetFieldState(value PropertyState) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventPropertyStruct, recv.native, "state", argValue)
}

// EventPropertyStruct creates an uninitialised EventProperty.
func EventPropertyStruct() *EventProperty {
	err := eventPropertyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventPropertyNewFromNative(eventPropertyStruct.Alloc())
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

func EventProximityNewFromNative(native uintptr) *EventProximity {
	return &EventProximity{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventProximity) FieldType() EventType {
	argValue := gi.StructFieldGet(eventProximityStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventProximity) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventProximityStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventProximity) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventProximityStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventProximity) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventProximityStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventProximity) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventProximityStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventProximity) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventProximityStruct, recv.native, "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventProximity) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventProximityStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventProximity) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventProximityStruct, recv.native, "time", argValue)
}

// FieldDevice returns the C field 'device'.
func (recv *EventProximity) FieldDevice() *Device {
	argValue := gi.StructFieldGet(eventProximityStruct, recv.native, "device")
	value := DeviceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDevice sets the value of the C field 'device'.
func (recv *EventProximity) SetFieldDevice(value *Device) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventProximityStruct, recv.native, "device", argValue)
}

// EventProximityStruct creates an uninitialised EventProximity.
func EventProximityStruct() *EventProximity {
	err := eventProximityStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventProximityNewFromNative(eventProximityStruct.Alloc())
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

func EventScrollNewFromNative(native uintptr) *EventScroll {
	return &EventScroll{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventScroll) FieldType() EventType {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventScroll) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventScrollStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventScroll) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventScroll) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventScrollStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventScroll) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventScroll) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventScrollStruct, recv.native, "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventScroll) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventScroll) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventScrollStruct, recv.native, "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventScroll) FieldX() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventScroll) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventScroll) FieldY() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventScroll) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.native, "y", argValue)
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// UNSUPPORTED : C value 'state' : for field setter : no Go type for 'ModifierType'

// FieldDirection returns the C field 'direction'.
func (recv *EventScroll) FieldDirection() ScrollDirection {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "direction")
	value := ScrollDirection(argValue.Int32())
	return value
}

// SetFieldDirection sets the value of the C field 'direction'.
func (recv *EventScroll) SetFieldDirection(value ScrollDirection) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventScrollStruct, recv.native, "direction", argValue)
}

// FieldDevice returns the C field 'device'.
func (recv *EventScroll) FieldDevice() *Device {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "device")
	value := DeviceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDevice sets the value of the C field 'device'.
func (recv *EventScroll) SetFieldDevice(value *Device) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventScrollStruct, recv.native, "device", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventScroll) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventScroll) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.native, "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventScroll) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventScroll) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.native, "y_root", argValue)
}

// FieldDeltaX returns the C field 'delta_x'.
func (recv *EventScroll) FieldDeltaX() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "delta_x")
	value := argValue.Float64()
	return value
}

// SetFieldDeltaX sets the value of the C field 'delta_x'.
func (recv *EventScroll) SetFieldDeltaX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.native, "delta_x", argValue)
}

// FieldDeltaY returns the C field 'delta_y'.
func (recv *EventScroll) FieldDeltaY() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "delta_y")
	value := argValue.Float64()
	return value
}

// SetFieldDeltaY sets the value of the C field 'delta_y'.
func (recv *EventScroll) SetFieldDeltaY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.native, "delta_y", argValue)
}

// FieldIsStop returns the C field 'is_stop'.
func (recv *EventScroll) FieldIsStop() uint32 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.native, "is_stop")
	value := argValue.Uint32()
	return value
}

// SetFieldIsStop sets the value of the C field 'is_stop'.
func (recv *EventScroll) SetFieldIsStop(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventScrollStruct, recv.native, "is_stop", argValue)
}

// EventScrollStruct creates an uninitialised EventScroll.
func EventScrollStruct() *EventScroll {
	err := eventScrollStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventScrollNewFromNative(eventScrollStruct.Alloc())
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

func EventSelectionNewFromNative(native uintptr) *EventSelection {
	return &EventSelection{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventSelection) FieldType() EventType {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventSelection) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventSelectionStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventSelection) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventSelection) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventSelectionStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventSelection) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventSelection) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventSelectionStruct, recv.native, "send_event", argValue)
}

// FieldSelection returns the C field 'selection'.
func (recv *EventSelection) FieldSelection() *Atom {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.native, "selection")
	value := AtomNewFromNative(argValue.Pointer())
	return value
}

// SetFieldSelection sets the value of the C field 'selection'.
func (recv *EventSelection) SetFieldSelection(value *Atom) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventSelectionStruct, recv.native, "selection", argValue)
}

// FieldTarget returns the C field 'target'.
func (recv *EventSelection) FieldTarget() *Atom {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.native, "target")
	value := AtomNewFromNative(argValue.Pointer())
	return value
}

// SetFieldTarget sets the value of the C field 'target'.
func (recv *EventSelection) SetFieldTarget(value *Atom) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventSelectionStruct, recv.native, "target", argValue)
}

// FieldProperty returns the C field 'property'.
func (recv *EventSelection) FieldProperty() *Atom {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.native, "property")
	value := AtomNewFromNative(argValue.Pointer())
	return value
}

// SetFieldProperty sets the value of the C field 'property'.
func (recv *EventSelection) SetFieldProperty(value *Atom) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventSelectionStruct, recv.native, "property", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventSelection) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventSelection) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventSelectionStruct, recv.native, "time", argValue)
}

// FieldRequestor returns the C field 'requestor'.
func (recv *EventSelection) FieldRequestor() *Window {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.native, "requestor")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldRequestor sets the value of the C field 'requestor'.
func (recv *EventSelection) SetFieldRequestor(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventSelectionStruct, recv.native, "requestor", argValue)
}

// EventSelectionStruct creates an uninitialised EventSelection.
func EventSelectionStruct() *EventSelection {
	err := eventSelectionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventSelectionNewFromNative(eventSelectionStruct.Alloc())
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

func EventSequenceNewFromNative(native uintptr) *EventSequence {
	return &EventSequence{native: native}
}

// EventSequenceStruct creates an uninitialised EventSequence.
func EventSequenceStruct() *EventSequence {
	err := eventSequenceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventSequenceNewFromNative(eventSequenceStruct.Alloc())
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

func EventSettingNewFromNative(native uintptr) *EventSetting {
	return &EventSetting{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventSetting) FieldType() EventType {
	argValue := gi.StructFieldGet(eventSettingStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventSetting) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventSettingStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventSetting) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventSettingStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventSetting) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventSettingStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventSetting) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventSettingStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventSetting) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventSettingStruct, recv.native, "send_event", argValue)
}

// FieldAction returns the C field 'action'.
func (recv *EventSetting) FieldAction() SettingAction {
	argValue := gi.StructFieldGet(eventSettingStruct, recv.native, "action")
	value := SettingAction(argValue.Int32())
	return value
}

// SetFieldAction sets the value of the C field 'action'.
func (recv *EventSetting) SetFieldAction(value SettingAction) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventSettingStruct, recv.native, "action", argValue)
}

// FieldName returns the C field 'name'.
func (recv *EventSetting) FieldName() string {
	argValue := gi.StructFieldGet(eventSettingStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *EventSetting) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(eventSettingStruct, recv.native, "name", argValue)
}

// EventSettingStruct creates an uninitialised EventSetting.
func EventSettingStruct() *EventSetting {
	err := eventSettingStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventSettingNewFromNative(eventSettingStruct.Alloc())
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

func EventTouchNewFromNative(native uintptr) *EventTouch {
	return &EventTouch{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventTouch) FieldType() EventType {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventTouch) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventTouchStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventTouch) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventTouch) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventTouchStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventTouch) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventTouch) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchStruct, recv.native, "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventTouch) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventTouch) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventTouchStruct, recv.native, "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventTouch) FieldX() float64 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventTouch) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventTouch) FieldY() float64 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventTouch) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchStruct, recv.native, "y", argValue)
}

// FieldAxes returns the C field 'axes'.
func (recv *EventTouch) FieldAxes() float64 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "axes")
	value := argValue.Float64()
	return value
}

// SetFieldAxes sets the value of the C field 'axes'.
func (recv *EventTouch) SetFieldAxes(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchStruct, recv.native, "axes", argValue)
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// UNSUPPORTED : C value 'state' : for field setter : no Go type for 'ModifierType'

// FieldSequence returns the C field 'sequence'.
func (recv *EventTouch) FieldSequence() *EventSequence {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "sequence")
	value := EventSequenceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldSequence sets the value of the C field 'sequence'.
func (recv *EventTouch) SetFieldSequence(value *EventSequence) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventTouchStruct, recv.native, "sequence", argValue)
}

// FieldEmulatingPointer returns the C field 'emulating_pointer'.
func (recv *EventTouch) FieldEmulatingPointer() bool {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "emulating_pointer")
	value := argValue.Boolean()
	return value
}

// SetFieldEmulatingPointer sets the value of the C field 'emulating_pointer'.
func (recv *EventTouch) SetFieldEmulatingPointer(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(eventTouchStruct, recv.native, "emulating_pointer", argValue)
}

// FieldDevice returns the C field 'device'.
func (recv *EventTouch) FieldDevice() *Device {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "device")
	value := DeviceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDevice sets the value of the C field 'device'.
func (recv *EventTouch) SetFieldDevice(value *Device) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventTouchStruct, recv.native, "device", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventTouch) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventTouch) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchStruct, recv.native, "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventTouch) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventTouch) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchStruct, recv.native, "y_root", argValue)
}

// EventTouchStruct creates an uninitialised EventTouch.
func EventTouchStruct() *EventTouch {
	err := eventTouchStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventTouchNewFromNative(eventTouchStruct.Alloc())
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

func EventTouchpadPinchNewFromNative(native uintptr) *EventTouchpadPinch {
	return &EventTouchpadPinch{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventTouchpadPinch) FieldType() EventType {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventTouchpadPinch) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventTouchpadPinch) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventTouchpadPinch) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventTouchpadPinch) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventTouchpadPinch) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "send_event", argValue)
}

// FieldPhase returns the C field 'phase'.
func (recv *EventTouchpadPinch) FieldPhase() int8 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "phase")
	value := argValue.Int8()
	return value
}

// SetFieldPhase sets the value of the C field 'phase'.
func (recv *EventTouchpadPinch) SetFieldPhase(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "phase", argValue)
}

// FieldNFingers returns the C field 'n_fingers'.
func (recv *EventTouchpadPinch) FieldNFingers() int8 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "n_fingers")
	value := argValue.Int8()
	return value
}

// SetFieldNFingers sets the value of the C field 'n_fingers'.
func (recv *EventTouchpadPinch) SetFieldNFingers(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "n_fingers", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventTouchpadPinch) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventTouchpadPinch) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventTouchpadPinch) FieldX() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventTouchpadPinch) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventTouchpadPinch) FieldY() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventTouchpadPinch) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "y", argValue)
}

// FieldDx returns the C field 'dx'.
func (recv *EventTouchpadPinch) FieldDx() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "dx")
	value := argValue.Float64()
	return value
}

// SetFieldDx sets the value of the C field 'dx'.
func (recv *EventTouchpadPinch) SetFieldDx(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "dx", argValue)
}

// FieldDy returns the C field 'dy'.
func (recv *EventTouchpadPinch) FieldDy() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "dy")
	value := argValue.Float64()
	return value
}

// SetFieldDy sets the value of the C field 'dy'.
func (recv *EventTouchpadPinch) SetFieldDy(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "dy", argValue)
}

// FieldAngleDelta returns the C field 'angle_delta'.
func (recv *EventTouchpadPinch) FieldAngleDelta() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "angle_delta")
	value := argValue.Float64()
	return value
}

// SetFieldAngleDelta sets the value of the C field 'angle_delta'.
func (recv *EventTouchpadPinch) SetFieldAngleDelta(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "angle_delta", argValue)
}

// FieldScale returns the C field 'scale'.
func (recv *EventTouchpadPinch) FieldScale() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "scale")
	value := argValue.Float64()
	return value
}

// SetFieldScale sets the value of the C field 'scale'.
func (recv *EventTouchpadPinch) SetFieldScale(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "scale", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventTouchpadPinch) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventTouchpadPinch) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventTouchpadPinch) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventTouchpadPinch) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.native, "y_root", argValue)
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// UNSUPPORTED : C value 'state' : for field setter : no Go type for 'ModifierType'

// EventTouchpadPinchStruct creates an uninitialised EventTouchpadPinch.
func EventTouchpadPinchStruct() *EventTouchpadPinch {
	err := eventTouchpadPinchStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventTouchpadPinchNewFromNative(eventTouchpadPinchStruct.Alloc())
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

func EventTouchpadSwipeNewFromNative(native uintptr) *EventTouchpadSwipe {
	return &EventTouchpadSwipe{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventTouchpadSwipe) FieldType() EventType {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventTouchpadSwipe) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventTouchpadSwipe) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventTouchpadSwipe) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventTouchpadSwipe) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventTouchpadSwipe) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "send_event", argValue)
}

// FieldPhase returns the C field 'phase'.
func (recv *EventTouchpadSwipe) FieldPhase() int8 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "phase")
	value := argValue.Int8()
	return value
}

// SetFieldPhase sets the value of the C field 'phase'.
func (recv *EventTouchpadSwipe) SetFieldPhase(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "phase", argValue)
}

// FieldNFingers returns the C field 'n_fingers'.
func (recv *EventTouchpadSwipe) FieldNFingers() int8 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "n_fingers")
	value := argValue.Int8()
	return value
}

// SetFieldNFingers sets the value of the C field 'n_fingers'.
func (recv *EventTouchpadSwipe) SetFieldNFingers(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "n_fingers", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventTouchpadSwipe) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventTouchpadSwipe) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventTouchpadSwipe) FieldX() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventTouchpadSwipe) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventTouchpadSwipe) FieldY() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventTouchpadSwipe) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "y", argValue)
}

// FieldDx returns the C field 'dx'.
func (recv *EventTouchpadSwipe) FieldDx() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "dx")
	value := argValue.Float64()
	return value
}

// SetFieldDx sets the value of the C field 'dx'.
func (recv *EventTouchpadSwipe) SetFieldDx(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "dx", argValue)
}

// FieldDy returns the C field 'dy'.
func (recv *EventTouchpadSwipe) FieldDy() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "dy")
	value := argValue.Float64()
	return value
}

// SetFieldDy sets the value of the C field 'dy'.
func (recv *EventTouchpadSwipe) SetFieldDy(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "dy", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventTouchpadSwipe) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventTouchpadSwipe) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventTouchpadSwipe) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.native, "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventTouchpadSwipe) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.native, "y_root", argValue)
}

// UNSUPPORTED : C value 'state' : for field getter : no Go type for 'ModifierType'

// UNSUPPORTED : C value 'state' : for field setter : no Go type for 'ModifierType'

// EventTouchpadSwipeStruct creates an uninitialised EventTouchpadSwipe.
func EventTouchpadSwipeStruct() *EventTouchpadSwipe {
	err := eventTouchpadSwipeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventTouchpadSwipeNewFromNative(eventTouchpadSwipeStruct.Alloc())
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

func EventVisibilityNewFromNative(native uintptr) *EventVisibility {
	return &EventVisibility{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventVisibility) FieldType() EventType {
	argValue := gi.StructFieldGet(eventVisibilityStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventVisibility) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventVisibilityStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventVisibility) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventVisibilityStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventVisibility) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventVisibilityStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventVisibility) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventVisibilityStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventVisibility) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventVisibilityStruct, recv.native, "send_event", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventVisibility) FieldState() VisibilityState {
	argValue := gi.StructFieldGet(eventVisibilityStruct, recv.native, "state")
	value := VisibilityState(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventVisibility) SetFieldState(value VisibilityState) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventVisibilityStruct, recv.native, "state", argValue)
}

// EventVisibilityStruct creates an uninitialised EventVisibility.
func EventVisibilityStruct() *EventVisibility {
	err := eventVisibilityStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventVisibilityNewFromNative(eventVisibilityStruct.Alloc())
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

func EventWindowStateNewFromNative(native uintptr) *EventWindowState {
	return &EventWindowState{native: native}
}

// FieldType returns the C field 'type'.
func (recv *EventWindowState) FieldType() EventType {
	argValue := gi.StructFieldGet(eventWindowStateStruct, recv.native, "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventWindowState) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventWindowStateStruct, recv.native, "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventWindowState) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventWindowStateStruct, recv.native, "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventWindowState) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(eventWindowStateStruct, recv.native, "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventWindowState) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventWindowStateStruct, recv.native, "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventWindowState) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventWindowStateStruct, recv.native, "send_event", argValue)
}

// UNSUPPORTED : C value 'changed_mask' : for field getter : no Go type for 'WindowState'

// UNSUPPORTED : C value 'changed_mask' : for field setter : no Go type for 'WindowState'

// UNSUPPORTED : C value 'new_window_state' : for field getter : no Go type for 'WindowState'

// UNSUPPORTED : C value 'new_window_state' : for field setter : no Go type for 'WindowState'

// EventWindowStateStruct creates an uninitialised EventWindowState.
func EventWindowStateStruct() *EventWindowState {
	err := eventWindowStateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EventWindowStateNewFromNative(eventWindowStateStruct.Alloc())
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

func FrameClockClassNewFromNative(native uintptr) *FrameClockClass {
	return &FrameClockClass{native: native}
}

// FrameClockClassStruct creates an uninitialised FrameClockClass.
func FrameClockClassStruct() *FrameClockClass {
	err := frameClockClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FrameClockClassNewFromNative(frameClockClassStruct.Alloc())
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

func FrameClockPrivateNewFromNative(native uintptr) *FrameClockPrivate {
	return &FrameClockPrivate{native: native}
}

// FrameClockPrivateStruct creates an uninitialised FrameClockPrivate.
func FrameClockPrivateStruct() *FrameClockPrivate {
	err := frameClockPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FrameClockPrivateNewFromNative(frameClockPrivateStruct.Alloc())
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

func FrameTimingsNewFromNative(native uintptr) *FrameTimings {
	return &FrameTimings{native: native}
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

	retGo := FrameTimingsNewFromNative(ret.Pointer())

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

	structGo := FrameTimingsNewFromNative(frameTimingsStruct.Alloc())
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

func GeometryNewFromNative(native uintptr) *Geometry {
	return &Geometry{native: native}
}

// FieldMinWidth returns the C field 'min_width'.
func (recv *Geometry) FieldMinWidth() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "min_width")
	value := argValue.Int32()
	return value
}

// SetFieldMinWidth sets the value of the C field 'min_width'.
func (recv *Geometry) SetFieldMinWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.native, "min_width", argValue)
}

// FieldMinHeight returns the C field 'min_height'.
func (recv *Geometry) FieldMinHeight() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "min_height")
	value := argValue.Int32()
	return value
}

// SetFieldMinHeight sets the value of the C field 'min_height'.
func (recv *Geometry) SetFieldMinHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.native, "min_height", argValue)
}

// FieldMaxWidth returns the C field 'max_width'.
func (recv *Geometry) FieldMaxWidth() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "max_width")
	value := argValue.Int32()
	return value
}

// SetFieldMaxWidth sets the value of the C field 'max_width'.
func (recv *Geometry) SetFieldMaxWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.native, "max_width", argValue)
}

// FieldMaxHeight returns the C field 'max_height'.
func (recv *Geometry) FieldMaxHeight() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "max_height")
	value := argValue.Int32()
	return value
}

// SetFieldMaxHeight sets the value of the C field 'max_height'.
func (recv *Geometry) SetFieldMaxHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.native, "max_height", argValue)
}

// FieldBaseWidth returns the C field 'base_width'.
func (recv *Geometry) FieldBaseWidth() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "base_width")
	value := argValue.Int32()
	return value
}

// SetFieldBaseWidth sets the value of the C field 'base_width'.
func (recv *Geometry) SetFieldBaseWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.native, "base_width", argValue)
}

// FieldBaseHeight returns the C field 'base_height'.
func (recv *Geometry) FieldBaseHeight() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "base_height")
	value := argValue.Int32()
	return value
}

// SetFieldBaseHeight sets the value of the C field 'base_height'.
func (recv *Geometry) SetFieldBaseHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.native, "base_height", argValue)
}

// FieldWidthInc returns the C field 'width_inc'.
func (recv *Geometry) FieldWidthInc() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "width_inc")
	value := argValue.Int32()
	return value
}

// SetFieldWidthInc sets the value of the C field 'width_inc'.
func (recv *Geometry) SetFieldWidthInc(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.native, "width_inc", argValue)
}

// FieldHeightInc returns the C field 'height_inc'.
func (recv *Geometry) FieldHeightInc() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "height_inc")
	value := argValue.Int32()
	return value
}

// SetFieldHeightInc sets the value of the C field 'height_inc'.
func (recv *Geometry) SetFieldHeightInc(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.native, "height_inc", argValue)
}

// FieldMinAspect returns the C field 'min_aspect'.
func (recv *Geometry) FieldMinAspect() float64 {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "min_aspect")
	value := argValue.Float64()
	return value
}

// SetFieldMinAspect sets the value of the C field 'min_aspect'.
func (recv *Geometry) SetFieldMinAspect(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(geometryStruct, recv.native, "min_aspect", argValue)
}

// FieldMaxAspect returns the C field 'max_aspect'.
func (recv *Geometry) FieldMaxAspect() float64 {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "max_aspect")
	value := argValue.Float64()
	return value
}

// SetFieldMaxAspect sets the value of the C field 'max_aspect'.
func (recv *Geometry) SetFieldMaxAspect(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(geometryStruct, recv.native, "max_aspect", argValue)
}

// FieldWinGravity returns the C field 'win_gravity'.
func (recv *Geometry) FieldWinGravity() Gravity {
	argValue := gi.StructFieldGet(geometryStruct, recv.native, "win_gravity")
	value := Gravity(argValue.Int32())
	return value
}

// SetFieldWinGravity sets the value of the C field 'win_gravity'.
func (recv *Geometry) SetFieldWinGravity(value Gravity) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(geometryStruct, recv.native, "win_gravity", argValue)
}

// GeometryStruct creates an uninitialised Geometry.
func GeometryStruct() *Geometry {
	err := geometryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GeometryNewFromNative(geometryStruct.Alloc())
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

func KeymapKeyNewFromNative(native uintptr) *KeymapKey {
	return &KeymapKey{native: native}
}

// FieldKeycode returns the C field 'keycode'.
func (recv *KeymapKey) FieldKeycode() uint32 {
	argValue := gi.StructFieldGet(keymapKeyStruct, recv.native, "keycode")
	value := argValue.Uint32()
	return value
}

// SetFieldKeycode sets the value of the C field 'keycode'.
func (recv *KeymapKey) SetFieldKeycode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(keymapKeyStruct, recv.native, "keycode", argValue)
}

// FieldGroup returns the C field 'group'.
func (recv *KeymapKey) FieldGroup() int32 {
	argValue := gi.StructFieldGet(keymapKeyStruct, recv.native, "group")
	value := argValue.Int32()
	return value
}

// SetFieldGroup sets the value of the C field 'group'.
func (recv *KeymapKey) SetFieldGroup(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(keymapKeyStruct, recv.native, "group", argValue)
}

// FieldLevel returns the C field 'level'.
func (recv *KeymapKey) FieldLevel() int32 {
	argValue := gi.StructFieldGet(keymapKeyStruct, recv.native, "level")
	value := argValue.Int32()
	return value
}

// SetFieldLevel sets the value of the C field 'level'.
func (recv *KeymapKey) SetFieldLevel(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(keymapKeyStruct, recv.native, "level", argValue)
}

// KeymapKeyStruct creates an uninitialised KeymapKey.
func KeymapKeyStruct() *KeymapKey {
	err := keymapKeyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := KeymapKeyNewFromNative(keymapKeyStruct.Alloc())
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

func MonitorClassNewFromNative(native uintptr) *MonitorClass {
	return &MonitorClass{native: native}
}

// MonitorClassStruct creates an uninitialised MonitorClass.
func MonitorClassStruct() *MonitorClass {
	err := monitorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MonitorClassNewFromNative(monitorClassStruct.Alloc())
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

func PointNewFromNative(native uintptr) *Point {
	return &Point{native: native}
}

// FieldX returns the C field 'x'.
func (recv *Point) FieldX() int32 {
	argValue := gi.StructFieldGet(pointStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *Point) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(pointStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *Point) FieldY() int32 {
	argValue := gi.StructFieldGet(pointStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *Point) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(pointStruct, recv.native, "y", argValue)
}

// PointStruct creates an uninitialised Point.
func PointStruct() *Point {
	err := pointStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PointNewFromNative(pointStruct.Alloc())
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

func RGBANewFromNative(native uintptr) *RGBA {
	return &RGBA{native: native}
}

// FieldRed returns the C field 'red'.
func (recv *RGBA) FieldRed() float64 {
	argValue := gi.StructFieldGet(rGBAStruct, recv.native, "red")
	value := argValue.Float64()
	return value
}

// SetFieldRed sets the value of the C field 'red'.
func (recv *RGBA) SetFieldRed(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rGBAStruct, recv.native, "red", argValue)
}

// FieldGreen returns the C field 'green'.
func (recv *RGBA) FieldGreen() float64 {
	argValue := gi.StructFieldGet(rGBAStruct, recv.native, "green")
	value := argValue.Float64()
	return value
}

// SetFieldGreen sets the value of the C field 'green'.
func (recv *RGBA) SetFieldGreen(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rGBAStruct, recv.native, "green", argValue)
}

// FieldBlue returns the C field 'blue'.
func (recv *RGBA) FieldBlue() float64 {
	argValue := gi.StructFieldGet(rGBAStruct, recv.native, "blue")
	value := argValue.Float64()
	return value
}

// SetFieldBlue sets the value of the C field 'blue'.
func (recv *RGBA) SetFieldBlue(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rGBAStruct, recv.native, "blue", argValue)
}

// FieldAlpha returns the C field 'alpha'.
func (recv *RGBA) FieldAlpha() float64 {
	argValue := gi.StructFieldGet(rGBAStruct, recv.native, "alpha")
	value := argValue.Float64()
	return value
}

// SetFieldAlpha sets the value of the C field 'alpha'.
func (recv *RGBA) SetFieldAlpha(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rGBAStruct, recv.native, "alpha", argValue)
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

	retGo := RGBANewFromNative(ret.Pointer())

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

	structGo := RGBANewFromNative(rGBAStruct.Alloc())
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

func RectangleNewFromNative(native uintptr) *Rectangle {
	return &Rectangle{native: native}
}

// FieldX returns the C field 'x'.
func (recv *Rectangle) FieldX() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *Rectangle) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *Rectangle) FieldY() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *Rectangle) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *Rectangle) FieldWidth() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *Rectangle) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *Rectangle) FieldHeight() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *Rectangle) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "height", argValue)
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
	out0 := RectangleNewFromNative(outArgs[0].Pointer())

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

	out0 := RectangleNewFromNative(outArgs[0].Pointer())

	return out0
}

// RectangleStruct creates an uninitialised Rectangle.
func RectangleStruct() *Rectangle {
	err := rectangleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RectangleNewFromNative(rectangleStruct.Alloc())
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

func TimeCoordNewFromNative(native uintptr) *TimeCoord {
	return &TimeCoord{native: native}
}

// FieldTime returns the C field 'time'.
func (recv *TimeCoord) FieldTime() uint32 {
	argValue := gi.StructFieldGet(timeCoordStruct, recv.native, "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *TimeCoord) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(timeCoordStruct, recv.native, "time", argValue)
}

// UNSUPPORTED : C value 'axes' : for field getter : missing Type

// UNSUPPORTED : C value 'axes' : for field setter : missing Type

// TimeCoordStruct creates an uninitialised TimeCoord.
func TimeCoordStruct() *TimeCoord {
	err := timeCoordStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TimeCoordNewFromNative(timeCoordStruct.Alloc())
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

func WindowAttrNewFromNative(native uintptr) *WindowAttr {
	return &WindowAttr{native: native}
}

// FieldTitle returns the C field 'title'.
func (recv *WindowAttr) FieldTitle() string {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "title")
	value := argValue.String(false)
	return value
}

// SetFieldTitle sets the value of the C field 'title'.
func (recv *WindowAttr) SetFieldTitle(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(windowAttrStruct, recv.native, "title", argValue)
}

// FieldEventMask returns the C field 'event_mask'.
func (recv *WindowAttr) FieldEventMask() int32 {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "event_mask")
	value := argValue.Int32()
	return value
}

// SetFieldEventMask sets the value of the C field 'event_mask'.
func (recv *WindowAttr) SetFieldEventMask(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(windowAttrStruct, recv.native, "event_mask", argValue)
}

// FieldX returns the C field 'x'.
func (recv *WindowAttr) FieldX() int32 {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *WindowAttr) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(windowAttrStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *WindowAttr) FieldY() int32 {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *WindowAttr) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(windowAttrStruct, recv.native, "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *WindowAttr) FieldWidth() int32 {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *WindowAttr) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(windowAttrStruct, recv.native, "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *WindowAttr) FieldHeight() int32 {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *WindowAttr) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(windowAttrStruct, recv.native, "height", argValue)
}

// FieldWclass returns the C field 'wclass'.
func (recv *WindowAttr) FieldWclass() WindowWindowClass {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "wclass")
	value := WindowWindowClass(argValue.Int32())
	return value
}

// SetFieldWclass sets the value of the C field 'wclass'.
func (recv *WindowAttr) SetFieldWclass(value WindowWindowClass) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(windowAttrStruct, recv.native, "wclass", argValue)
}

// FieldVisual returns the C field 'visual'.
func (recv *WindowAttr) FieldVisual() *Visual {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "visual")
	value := VisualNewFromNative(argValue.Pointer())
	return value
}

// SetFieldVisual sets the value of the C field 'visual'.
func (recv *WindowAttr) SetFieldVisual(value *Visual) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(windowAttrStruct, recv.native, "visual", argValue)
}

// FieldWindowType returns the C field 'window_type'.
func (recv *WindowAttr) FieldWindowType() WindowType {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "window_type")
	value := WindowType(argValue.Int32())
	return value
}

// SetFieldWindowType sets the value of the C field 'window_type'.
func (recv *WindowAttr) SetFieldWindowType(value WindowType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(windowAttrStruct, recv.native, "window_type", argValue)
}

// FieldCursor returns the C field 'cursor'.
func (recv *WindowAttr) FieldCursor() *Cursor {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "cursor")
	value := CursorNewFromNative(argValue.Pointer())
	return value
}

// SetFieldCursor sets the value of the C field 'cursor'.
func (recv *WindowAttr) SetFieldCursor(value *Cursor) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(windowAttrStruct, recv.native, "cursor", argValue)
}

// FieldWmclassName returns the C field 'wmclass_name'.
func (recv *WindowAttr) FieldWmclassName() string {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "wmclass_name")
	value := argValue.String(false)
	return value
}

// SetFieldWmclassName sets the value of the C field 'wmclass_name'.
func (recv *WindowAttr) SetFieldWmclassName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(windowAttrStruct, recv.native, "wmclass_name", argValue)
}

// FieldWmclassClass returns the C field 'wmclass_class'.
func (recv *WindowAttr) FieldWmclassClass() string {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "wmclass_class")
	value := argValue.String(false)
	return value
}

// SetFieldWmclassClass sets the value of the C field 'wmclass_class'.
func (recv *WindowAttr) SetFieldWmclassClass(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(windowAttrStruct, recv.native, "wmclass_class", argValue)
}

// FieldOverrideRedirect returns the C field 'override_redirect'.
func (recv *WindowAttr) FieldOverrideRedirect() bool {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "override_redirect")
	value := argValue.Boolean()
	return value
}

// SetFieldOverrideRedirect sets the value of the C field 'override_redirect'.
func (recv *WindowAttr) SetFieldOverrideRedirect(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(windowAttrStruct, recv.native, "override_redirect", argValue)
}

// FieldTypeHint returns the C field 'type_hint'.
func (recv *WindowAttr) FieldTypeHint() WindowTypeHint {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.native, "type_hint")
	value := WindowTypeHint(argValue.Int32())
	return value
}

// SetFieldTypeHint sets the value of the C field 'type_hint'.
func (recv *WindowAttr) SetFieldTypeHint(value WindowTypeHint) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(windowAttrStruct, recv.native, "type_hint", argValue)
}

// WindowAttrStruct creates an uninitialised WindowAttr.
func WindowAttrStruct() *WindowAttr {
	err := windowAttrStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WindowAttrNewFromNative(windowAttrStruct.Alloc())
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

func WindowClassNewFromNative(native uintptr) *WindowClass {
	return &WindowClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'pick_embedded_child' : for field getter : missing Type

// UNSUPPORTED : C value 'pick_embedded_child' : for field setter : missing Type

// UNSUPPORTED : C value 'to_embedder' : for field getter : missing Type

// UNSUPPORTED : C value 'to_embedder' : for field setter : missing Type

// UNSUPPORTED : C value 'from_embedder' : for field getter : missing Type

// UNSUPPORTED : C value 'from_embedder' : for field setter : missing Type

// UNSUPPORTED : C value 'create_surface' : for field getter : missing Type

// UNSUPPORTED : C value 'create_surface' : for field setter : missing Type

// UNSUPPORTED : C value '_gdk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gdk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gdk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gdk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gdk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gdk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gdk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gdk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gdk_reserved8' : for field setter : missing Type

// WindowClassStruct creates an uninitialised WindowClass.
func WindowClassStruct() *WindowClass {
	err := windowClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WindowClassNewFromNative(windowClassStruct.Alloc())
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

func WindowRedirectNewFromNative(native uintptr) *WindowRedirect {
	return &WindowRedirect{native: native}
}

// WindowRedirectStruct creates an uninitialised WindowRedirect.
func WindowRedirectStruct() *WindowRedirect {
	err := windowRedirectStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WindowRedirectNewFromNative(windowRedirectStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWindowRedirect)
	return structGo
}
func finalizeWindowRedirect(obj *WindowRedirect) {
	windowRedirectStruct.Free(obj.native)
}
