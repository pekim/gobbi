// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
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
	native unsafe.Pointer
}

func AtomNewFromNative(native unsafe.Pointer) *Atom {
	instance := &Atom{native: native}

	return instance
}

/*
CastToAtom down casts any arbitrary Object to Atom.
Exercise care, as this is a potentially dangerous function
if the Object is not a Atom.
*/
func CastToAtom(object *gobject.Object) *Atom {
	return AtomNewFromNative(object.Native())
}

// Equals compares this Atom with another Atom, and returns true if they represent the same Object.
func (recv *Atom) Equals(other *Atom) bool {
	return other.Native() == recv.Native()
}

func (recv *Atom) Native() unsafe.Pointer {
	return recv.native
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
	inArgs[0].SetPointer(recv.Native())

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
	atomStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func ColorNewFromNative(native unsafe.Pointer) *Color {
	instance := &Color{native: native}

	return instance
}

/*
CastToColor down casts any arbitrary Object to Color.
Exercise care, as this is a potentially dangerous function
if the Object is not a Color.
*/
func CastToColor(object *gobject.Object) *Color {
	return ColorNewFromNative(object.Native())
}

// Equals compares this Color with another Color, and returns true if they represent the same Object.
func (recv *Color) Equals(other *Color) bool {
	return other.Native() == recv.Native()
}

func (recv *Color) Native() unsafe.Pointer {
	return recv.native
}

// FieldPixel returns the C field 'pixel'.
func (recv *Color) FieldPixel() uint32 {
	argValue := gi.StructFieldGet(colorStruct, recv.Native(), "pixel")
	value := argValue.Uint32()
	return value
}

// SetFieldPixel sets the value of the C field 'pixel'.
func (recv *Color) SetFieldPixel(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(colorStruct, recv.Native(), "pixel", argValue)
}

// FieldRed returns the C field 'red'.
func (recv *Color) FieldRed() uint16 {
	argValue := gi.StructFieldGet(colorStruct, recv.Native(), "red")
	value := argValue.Uint16()
	return value
}

// SetFieldRed sets the value of the C field 'red'.
func (recv *Color) SetFieldRed(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(colorStruct, recv.Native(), "red", argValue)
}

// FieldGreen returns the C field 'green'.
func (recv *Color) FieldGreen() uint16 {
	argValue := gi.StructFieldGet(colorStruct, recv.Native(), "green")
	value := argValue.Uint16()
	return value
}

// SetFieldGreen sets the value of the C field 'green'.
func (recv *Color) SetFieldGreen(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(colorStruct, recv.Native(), "green", argValue)
}

// FieldBlue returns the C field 'blue'.
func (recv *Color) FieldBlue() uint16 {
	argValue := gi.StructFieldGet(colorStruct, recv.Native(), "blue")
	value := argValue.Uint16()
	return value
}

// SetFieldBlue sets the value of the C field 'blue'.
func (recv *Color) SetFieldBlue(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(colorStruct, recv.Native(), "blue", argValue)
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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(colorb.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	colorStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func DevicePadInterfaceNewFromNative(native unsafe.Pointer) *DevicePadInterface {
	instance := &DevicePadInterface{native: native}

	return instance
}

/*
CastToDevicePadInterface down casts any arbitrary Object to DevicePadInterface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DevicePadInterface.
*/
func CastToDevicePadInterface(object *gobject.Object) *DevicePadInterface {
	return DevicePadInterfaceNewFromNative(object.Native())
}

// Equals compares this DevicePadInterface with another DevicePadInterface, and returns true if they represent the same Object.
func (recv *DevicePadInterface) Equals(other *DevicePadInterface) bool {
	return other.Native() == recv.Native()
}

func (recv *DevicePadInterface) Native() unsafe.Pointer {
	return recv.native
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
	devicePadInterfaceStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func DrawingContextClassNewFromNative(native unsafe.Pointer) *DrawingContextClass {
	instance := &DrawingContextClass{native: native}

	return instance
}

/*
CastToDrawingContextClass down casts any arbitrary Object to DrawingContextClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DrawingContextClass.
*/
func CastToDrawingContextClass(object *gobject.Object) *DrawingContextClass {
	return DrawingContextClassNewFromNative(object.Native())
}

// Equals compares this DrawingContextClass with another DrawingContextClass, and returns true if they represent the same Object.
func (recv *DrawingContextClass) Equals(other *DrawingContextClass) bool {
	return other.Native() == recv.Native()
}

func (recv *DrawingContextClass) Native() unsafe.Pointer {
	return recv.native
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
	drawingContextClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventAnyNewFromNative(native unsafe.Pointer) *EventAny {
	instance := &EventAny{native: native}

	return instance
}

/*
CastToEventAny down casts any arbitrary Object to EventAny.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventAny.
*/
func CastToEventAny(object *gobject.Object) *EventAny {
	return EventAnyNewFromNative(object.Native())
}

// Equals compares this EventAny with another EventAny, and returns true if they represent the same Object.
func (recv *EventAny) Equals(other *EventAny) bool {
	return other.Native() == recv.Native()
}

func (recv *EventAny) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventAny) FieldType() EventType {
	argValue := gi.StructFieldGet(eventAnyStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventAny) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventAnyStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventAny) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventAnyStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventAny) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventAnyStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventAny) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventAnyStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventAny) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventAnyStruct, recv.Native(), "send_event", argValue)
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
	eventAnyStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventButtonNewFromNative(native unsafe.Pointer) *EventButton {
	instance := &EventButton{native: native}

	return instance
}

/*
CastToEventButton down casts any arbitrary Object to EventButton.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventButton.
*/
func CastToEventButton(object *gobject.Object) *EventButton {
	return EventButtonNewFromNative(object.Native())
}

// Equals compares this EventButton with another EventButton, and returns true if they represent the same Object.
func (recv *EventButton) Equals(other *EventButton) bool {
	return other.Native() == recv.Native()
}

func (recv *EventButton) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventButton) FieldType() EventType {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventButton) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventButton) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventButton) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventButton) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventButton) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventButton) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventButton) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventButton) FieldX() float64 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventButton) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventButton) FieldY() float64 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventButton) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "y", argValue)
}

// FieldAxes returns the C field 'axes'.
func (recv *EventButton) FieldAxes() float64 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "axes")
	value := argValue.Float64()
	return value
}

// SetFieldAxes sets the value of the C field 'axes'.
func (recv *EventButton) SetFieldAxes(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "axes", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventButton) FieldState() ModifierType {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "state")
	value := ModifierType(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventButton) SetFieldState(value ModifierType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "state", argValue)
}

// FieldButton returns the C field 'button'.
func (recv *EventButton) FieldButton() uint32 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "button")
	value := argValue.Uint32()
	return value
}

// SetFieldButton sets the value of the C field 'button'.
func (recv *EventButton) SetFieldButton(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "button", argValue)
}

// FieldDevice returns the C field 'device'.
func (recv *EventButton) FieldDevice() *Device {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "device")
	value := DeviceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDevice sets the value of the C field 'device'.
func (recv *EventButton) SetFieldDevice(value *Device) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "device", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventButton) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventButton) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventButton) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventButtonStruct, recv.Native(), "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventButton) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventButtonStruct, recv.Native(), "y_root", argValue)
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
	eventButtonStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventConfigureNewFromNative(native unsafe.Pointer) *EventConfigure {
	instance := &EventConfigure{native: native}

	return instance
}

/*
CastToEventConfigure down casts any arbitrary Object to EventConfigure.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventConfigure.
*/
func CastToEventConfigure(object *gobject.Object) *EventConfigure {
	return EventConfigureNewFromNative(object.Native())
}

// Equals compares this EventConfigure with another EventConfigure, and returns true if they represent the same Object.
func (recv *EventConfigure) Equals(other *EventConfigure) bool {
	return other.Native() == recv.Native()
}

func (recv *EventConfigure) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventConfigure) FieldType() EventType {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventConfigure) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventConfigureStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventConfigure) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventConfigure) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventConfigureStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventConfigure) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventConfigure) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventConfigureStruct, recv.Native(), "send_event", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventConfigure) FieldX() int32 {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.Native(), "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventConfigure) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventConfigureStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventConfigure) FieldY() int32 {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.Native(), "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventConfigure) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventConfigureStruct, recv.Native(), "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *EventConfigure) FieldWidth() int32 {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.Native(), "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *EventConfigure) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventConfigureStruct, recv.Native(), "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *EventConfigure) FieldHeight() int32 {
	argValue := gi.StructFieldGet(eventConfigureStruct, recv.Native(), "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *EventConfigure) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventConfigureStruct, recv.Native(), "height", argValue)
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
	eventConfigureStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventCrossingNewFromNative(native unsafe.Pointer) *EventCrossing {
	instance := &EventCrossing{native: native}

	return instance
}

/*
CastToEventCrossing down casts any arbitrary Object to EventCrossing.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventCrossing.
*/
func CastToEventCrossing(object *gobject.Object) *EventCrossing {
	return EventCrossingNewFromNative(object.Native())
}

// Equals compares this EventCrossing with another EventCrossing, and returns true if they represent the same Object.
func (recv *EventCrossing) Equals(other *EventCrossing) bool {
	return other.Native() == recv.Native()
}

func (recv *EventCrossing) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventCrossing) FieldType() EventType {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventCrossing) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventCrossing) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventCrossing) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventCrossing) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventCrossing) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "send_event", argValue)
}

// FieldSubwindow returns the C field 'subwindow'.
func (recv *EventCrossing) FieldSubwindow() *Window {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "subwindow")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldSubwindow sets the value of the C field 'subwindow'.
func (recv *EventCrossing) SetFieldSubwindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "subwindow", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventCrossing) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventCrossing) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventCrossing) FieldX() float64 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventCrossing) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventCrossing) FieldY() float64 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventCrossing) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "y", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventCrossing) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventCrossing) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventCrossing) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventCrossing) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "y_root", argValue)
}

// FieldMode returns the C field 'mode'.
func (recv *EventCrossing) FieldMode() CrossingMode {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "mode")
	value := CrossingMode(argValue.Int32())
	return value
}

// SetFieldMode sets the value of the C field 'mode'.
func (recv *EventCrossing) SetFieldMode(value CrossingMode) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "mode", argValue)
}

// FieldDetail returns the C field 'detail'.
func (recv *EventCrossing) FieldDetail() NotifyType {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "detail")
	value := NotifyType(argValue.Int32())
	return value
}

// SetFieldDetail sets the value of the C field 'detail'.
func (recv *EventCrossing) SetFieldDetail(value NotifyType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "detail", argValue)
}

// FieldFocus returns the C field 'focus'.
func (recv *EventCrossing) FieldFocus() bool {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "focus")
	value := argValue.Boolean()
	return value
}

// SetFieldFocus sets the value of the C field 'focus'.
func (recv *EventCrossing) SetFieldFocus(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "focus", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventCrossing) FieldState() ModifierType {
	argValue := gi.StructFieldGet(eventCrossingStruct, recv.Native(), "state")
	value := ModifierType(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventCrossing) SetFieldState(value ModifierType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventCrossingStruct, recv.Native(), "state", argValue)
}

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
	eventCrossingStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventDNDNewFromNative(native unsafe.Pointer) *EventDND {
	instance := &EventDND{native: native}

	return instance
}

/*
CastToEventDND down casts any arbitrary Object to EventDND.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventDND.
*/
func CastToEventDND(object *gobject.Object) *EventDND {
	return EventDNDNewFromNative(object.Native())
}

// Equals compares this EventDND with another EventDND, and returns true if they represent the same Object.
func (recv *EventDND) Equals(other *EventDND) bool {
	return other.Native() == recv.Native()
}

func (recv *EventDND) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventDND) FieldType() EventType {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventDND) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventDNDStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventDND) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventDND) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventDNDStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventDND) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventDND) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventDNDStruct, recv.Native(), "send_event", argValue)
}

// FieldContext returns the C field 'context'.
func (recv *EventDND) FieldContext() *DragContext {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.Native(), "context")
	value := DragContextNewFromNative(argValue.Pointer())
	return value
}

// SetFieldContext sets the value of the C field 'context'.
func (recv *EventDND) SetFieldContext(value *DragContext) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventDNDStruct, recv.Native(), "context", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventDND) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventDND) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventDNDStruct, recv.Native(), "time", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventDND) FieldXRoot() int16 {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.Native(), "x_root")
	value := argValue.Int16()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventDND) SetFieldXRoot(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.StructFieldSet(eventDNDStruct, recv.Native(), "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventDND) FieldYRoot() int16 {
	argValue := gi.StructFieldGet(eventDNDStruct, recv.Native(), "y_root")
	value := argValue.Int16()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventDND) SetFieldYRoot(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.StructFieldSet(eventDNDStruct, recv.Native(), "y_root", argValue)
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
	eventDNDStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventExposeNewFromNative(native unsafe.Pointer) *EventExpose {
	instance := &EventExpose{native: native}

	return instance
}

/*
CastToEventExpose down casts any arbitrary Object to EventExpose.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventExpose.
*/
func CastToEventExpose(object *gobject.Object) *EventExpose {
	return EventExposeNewFromNative(object.Native())
}

// Equals compares this EventExpose with another EventExpose, and returns true if they represent the same Object.
func (recv *EventExpose) Equals(other *EventExpose) bool {
	return other.Native() == recv.Native()
}

func (recv *EventExpose) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventExpose) FieldType() EventType {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventExpose) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventExposeStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventExpose) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventExpose) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventExposeStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventExpose) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventExpose) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventExposeStruct, recv.Native(), "send_event", argValue)
}

// FieldArea returns the C field 'area'.
func (recv *EventExpose) FieldArea() *Rectangle {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.Native(), "area")
	value := RectangleNewFromNative(argValue.Pointer())
	return value
}

// SetFieldArea sets the value of the C field 'area'.
func (recv *EventExpose) SetFieldArea(value *Rectangle) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventExposeStruct, recv.Native(), "area", argValue)
}

// FieldRegion returns the C field 'region'.
func (recv *EventExpose) FieldRegion() *cairo.Region {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.Native(), "region")
	value := cairo.RegionNewFromNative(argValue.Pointer())
	return value
}

// SetFieldRegion sets the value of the C field 'region'.
func (recv *EventExpose) SetFieldRegion(value *cairo.Region) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventExposeStruct, recv.Native(), "region", argValue)
}

// FieldCount returns the C field 'count'.
func (recv *EventExpose) FieldCount() int32 {
	argValue := gi.StructFieldGet(eventExposeStruct, recv.Native(), "count")
	value := argValue.Int32()
	return value
}

// SetFieldCount sets the value of the C field 'count'.
func (recv *EventExpose) SetFieldCount(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventExposeStruct, recv.Native(), "count", argValue)
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
	eventExposeStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventFocusNewFromNative(native unsafe.Pointer) *EventFocus {
	instance := &EventFocus{native: native}

	return instance
}

/*
CastToEventFocus down casts any arbitrary Object to EventFocus.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventFocus.
*/
func CastToEventFocus(object *gobject.Object) *EventFocus {
	return EventFocusNewFromNative(object.Native())
}

// Equals compares this EventFocus with another EventFocus, and returns true if they represent the same Object.
func (recv *EventFocus) Equals(other *EventFocus) bool {
	return other.Native() == recv.Native()
}

func (recv *EventFocus) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventFocus) FieldType() EventType {
	argValue := gi.StructFieldGet(eventFocusStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventFocus) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventFocusStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventFocus) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventFocusStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventFocus) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventFocusStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventFocus) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventFocusStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventFocus) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventFocusStruct, recv.Native(), "send_event", argValue)
}

// FieldIn returns the C field 'in'.
func (recv *EventFocus) FieldIn() int16 {
	argValue := gi.StructFieldGet(eventFocusStruct, recv.Native(), "in")
	value := argValue.Int16()
	return value
}

// SetFieldIn sets the value of the C field 'in'.
func (recv *EventFocus) SetFieldIn(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.StructFieldSet(eventFocusStruct, recv.Native(), "in", argValue)
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
	eventFocusStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventGrabBrokenNewFromNative(native unsafe.Pointer) *EventGrabBroken {
	instance := &EventGrabBroken{native: native}

	return instance
}

/*
CastToEventGrabBroken down casts any arbitrary Object to EventGrabBroken.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventGrabBroken.
*/
func CastToEventGrabBroken(object *gobject.Object) *EventGrabBroken {
	return EventGrabBrokenNewFromNative(object.Native())
}

// Equals compares this EventGrabBroken with another EventGrabBroken, and returns true if they represent the same Object.
func (recv *EventGrabBroken) Equals(other *EventGrabBroken) bool {
	return other.Native() == recv.Native()
}

func (recv *EventGrabBroken) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventGrabBroken) FieldType() EventType {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventGrabBroken) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventGrabBrokenStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventGrabBroken) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventGrabBroken) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventGrabBrokenStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventGrabBroken) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventGrabBroken) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventGrabBrokenStruct, recv.Native(), "send_event", argValue)
}

// FieldKeyboard returns the C field 'keyboard'.
func (recv *EventGrabBroken) FieldKeyboard() bool {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.Native(), "keyboard")
	value := argValue.Boolean()
	return value
}

// SetFieldKeyboard sets the value of the C field 'keyboard'.
func (recv *EventGrabBroken) SetFieldKeyboard(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(eventGrabBrokenStruct, recv.Native(), "keyboard", argValue)
}

// FieldImplicit returns the C field 'implicit'.
func (recv *EventGrabBroken) FieldImplicit() bool {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.Native(), "implicit")
	value := argValue.Boolean()
	return value
}

// SetFieldImplicit sets the value of the C field 'implicit'.
func (recv *EventGrabBroken) SetFieldImplicit(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(eventGrabBrokenStruct, recv.Native(), "implicit", argValue)
}

// FieldGrabWindow returns the C field 'grab_window'.
func (recv *EventGrabBroken) FieldGrabWindow() *Window {
	argValue := gi.StructFieldGet(eventGrabBrokenStruct, recv.Native(), "grab_window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldGrabWindow sets the value of the C field 'grab_window'.
func (recv *EventGrabBroken) SetFieldGrabWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventGrabBrokenStruct, recv.Native(), "grab_window", argValue)
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
	eventGrabBrokenStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventKeyNewFromNative(native unsafe.Pointer) *EventKey {
	instance := &EventKey{native: native}

	return instance
}

/*
CastToEventKey down casts any arbitrary Object to EventKey.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventKey.
*/
func CastToEventKey(object *gobject.Object) *EventKey {
	return EventKeyNewFromNative(object.Native())
}

// Equals compares this EventKey with another EventKey, and returns true if they represent the same Object.
func (recv *EventKey) Equals(other *EventKey) bool {
	return other.Native() == recv.Native()
}

func (recv *EventKey) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventKey) FieldType() EventType {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventKey) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventKey) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventKey) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventKey) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventKey) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventKey) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventKey) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "time", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventKey) FieldState() ModifierType {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "state")
	value := ModifierType(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventKey) SetFieldState(value ModifierType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "state", argValue)
}

// FieldKeyval returns the C field 'keyval'.
func (recv *EventKey) FieldKeyval() uint32 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "keyval")
	value := argValue.Uint32()
	return value
}

// SetFieldKeyval sets the value of the C field 'keyval'.
func (recv *EventKey) SetFieldKeyval(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "keyval", argValue)
}

// FieldLength returns the C field 'length'.
func (recv *EventKey) FieldLength() int32 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "length")
	value := argValue.Int32()
	return value
}

// SetFieldLength sets the value of the C field 'length'.
func (recv *EventKey) SetFieldLength(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "length", argValue)
}

// FieldString returns the C field 'string'.
func (recv *EventKey) FieldString() string {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "string")
	value := argValue.String(false)
	return value
}

// SetFieldString sets the value of the C field 'string'.
func (recv *EventKey) SetFieldString(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "string", argValue)
}

// FieldHardwareKeycode returns the C field 'hardware_keycode'.
func (recv *EventKey) FieldHardwareKeycode() uint16 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "hardware_keycode")
	value := argValue.Uint16()
	return value
}

// SetFieldHardwareKeycode sets the value of the C field 'hardware_keycode'.
func (recv *EventKey) SetFieldHardwareKeycode(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "hardware_keycode", argValue)
}

// FieldGroup returns the C field 'group'.
func (recv *EventKey) FieldGroup() uint8 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "group")
	value := argValue.Uint8()
	return value
}

// SetFieldGroup sets the value of the C field 'group'.
func (recv *EventKey) SetFieldGroup(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "group", argValue)
}

// FieldIsModifier returns the C field 'is_modifier'.
func (recv *EventKey) FieldIsModifier() uint32 {
	argValue := gi.StructFieldGet(eventKeyStruct, recv.Native(), "is_modifier")
	value := argValue.Uint32()
	return value
}

// SetFieldIsModifier sets the value of the C field 'is_modifier'.
func (recv *EventKey) SetFieldIsModifier(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventKeyStruct, recv.Native(), "is_modifier", argValue)
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
	eventKeyStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventMotionNewFromNative(native unsafe.Pointer) *EventMotion {
	instance := &EventMotion{native: native}

	return instance
}

/*
CastToEventMotion down casts any arbitrary Object to EventMotion.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventMotion.
*/
func CastToEventMotion(object *gobject.Object) *EventMotion {
	return EventMotionNewFromNative(object.Native())
}

// Equals compares this EventMotion with another EventMotion, and returns true if they represent the same Object.
func (recv *EventMotion) Equals(other *EventMotion) bool {
	return other.Native() == recv.Native()
}

func (recv *EventMotion) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventMotion) FieldType() EventType {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventMotion) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventMotion) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventMotion) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventMotion) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventMotion) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventMotion) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventMotion) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventMotion) FieldX() float64 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventMotion) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventMotion) FieldY() float64 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventMotion) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "y", argValue)
}

// FieldAxes returns the C field 'axes'.
func (recv *EventMotion) FieldAxes() float64 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "axes")
	value := argValue.Float64()
	return value
}

// SetFieldAxes sets the value of the C field 'axes'.
func (recv *EventMotion) SetFieldAxes(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "axes", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventMotion) FieldState() ModifierType {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "state")
	value := ModifierType(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventMotion) SetFieldState(value ModifierType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "state", argValue)
}

// FieldIsHint returns the C field 'is_hint'.
func (recv *EventMotion) FieldIsHint() int16 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "is_hint")
	value := argValue.Int16()
	return value
}

// SetFieldIsHint sets the value of the C field 'is_hint'.
func (recv *EventMotion) SetFieldIsHint(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "is_hint", argValue)
}

// FieldDevice returns the C field 'device'.
func (recv *EventMotion) FieldDevice() *Device {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "device")
	value := DeviceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDevice sets the value of the C field 'device'.
func (recv *EventMotion) SetFieldDevice(value *Device) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "device", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventMotion) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventMotion) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventMotion) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventMotionStruct, recv.Native(), "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventMotion) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventMotionStruct, recv.Native(), "y_root", argValue)
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
	eventMotionStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventOwnerChangeNewFromNative(native unsafe.Pointer) *EventOwnerChange {
	instance := &EventOwnerChange{native: native}

	return instance
}

/*
CastToEventOwnerChange down casts any arbitrary Object to EventOwnerChange.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventOwnerChange.
*/
func CastToEventOwnerChange(object *gobject.Object) *EventOwnerChange {
	return EventOwnerChangeNewFromNative(object.Native())
}

// Equals compares this EventOwnerChange with another EventOwnerChange, and returns true if they represent the same Object.
func (recv *EventOwnerChange) Equals(other *EventOwnerChange) bool {
	return other.Native() == recv.Native()
}

func (recv *EventOwnerChange) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventOwnerChange) FieldType() EventType {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventOwnerChange) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventOwnerChangeStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventOwnerChange) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventOwnerChange) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventOwnerChangeStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventOwnerChange) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventOwnerChange) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventOwnerChangeStruct, recv.Native(), "send_event", argValue)
}

// FieldOwner returns the C field 'owner'.
func (recv *EventOwnerChange) FieldOwner() *Window {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.Native(), "owner")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldOwner sets the value of the C field 'owner'.
func (recv *EventOwnerChange) SetFieldOwner(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventOwnerChangeStruct, recv.Native(), "owner", argValue)
}

// FieldReason returns the C field 'reason'.
func (recv *EventOwnerChange) FieldReason() OwnerChange {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.Native(), "reason")
	value := OwnerChange(argValue.Int32())
	return value
}

// SetFieldReason sets the value of the C field 'reason'.
func (recv *EventOwnerChange) SetFieldReason(value OwnerChange) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventOwnerChangeStruct, recv.Native(), "reason", argValue)
}

// FieldSelection returns the C field 'selection'.
func (recv *EventOwnerChange) FieldSelection() *Atom {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.Native(), "selection")
	value := AtomNewFromNative(argValue.Pointer())
	return value
}

// SetFieldSelection sets the value of the C field 'selection'.
func (recv *EventOwnerChange) SetFieldSelection(value *Atom) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventOwnerChangeStruct, recv.Native(), "selection", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventOwnerChange) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventOwnerChange) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventOwnerChangeStruct, recv.Native(), "time", argValue)
}

// FieldSelectionTime returns the C field 'selection_time'.
func (recv *EventOwnerChange) FieldSelectionTime() uint32 {
	argValue := gi.StructFieldGet(eventOwnerChangeStruct, recv.Native(), "selection_time")
	value := argValue.Uint32()
	return value
}

// SetFieldSelectionTime sets the value of the C field 'selection_time'.
func (recv *EventOwnerChange) SetFieldSelectionTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventOwnerChangeStruct, recv.Native(), "selection_time", argValue)
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
	eventOwnerChangeStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventPadAxisNewFromNative(native unsafe.Pointer) *EventPadAxis {
	instance := &EventPadAxis{native: native}

	return instance
}

/*
CastToEventPadAxis down casts any arbitrary Object to EventPadAxis.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventPadAxis.
*/
func CastToEventPadAxis(object *gobject.Object) *EventPadAxis {
	return EventPadAxisNewFromNative(object.Native())
}

// Equals compares this EventPadAxis with another EventPadAxis, and returns true if they represent the same Object.
func (recv *EventPadAxis) Equals(other *EventPadAxis) bool {
	return other.Native() == recv.Native()
}

func (recv *EventPadAxis) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventPadAxis) FieldType() EventType {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventPadAxis) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventPadAxisStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventPadAxis) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventPadAxis) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventPadAxisStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventPadAxis) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventPadAxis) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.Native(), "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventPadAxis) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventPadAxis) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.Native(), "time", argValue)
}

// FieldGroup returns the C field 'group'.
func (recv *EventPadAxis) FieldGroup() uint32 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.Native(), "group")
	value := argValue.Uint32()
	return value
}

// SetFieldGroup sets the value of the C field 'group'.
func (recv *EventPadAxis) SetFieldGroup(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.Native(), "group", argValue)
}

// FieldIndex returns the C field 'index'.
func (recv *EventPadAxis) FieldIndex() uint32 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.Native(), "index")
	value := argValue.Uint32()
	return value
}

// SetFieldIndex sets the value of the C field 'index'.
func (recv *EventPadAxis) SetFieldIndex(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.Native(), "index", argValue)
}

// FieldMode returns the C field 'mode'.
func (recv *EventPadAxis) FieldMode() uint32 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.Native(), "mode")
	value := argValue.Uint32()
	return value
}

// SetFieldMode sets the value of the C field 'mode'.
func (recv *EventPadAxis) SetFieldMode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.Native(), "mode", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *EventPadAxis) FieldValue() float64 {
	argValue := gi.StructFieldGet(eventPadAxisStruct, recv.Native(), "value")
	value := argValue.Float64()
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *EventPadAxis) SetFieldValue(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventPadAxisStruct, recv.Native(), "value", argValue)
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
	eventPadAxisStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventPadButtonNewFromNative(native unsafe.Pointer) *EventPadButton {
	instance := &EventPadButton{native: native}

	return instance
}

/*
CastToEventPadButton down casts any arbitrary Object to EventPadButton.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventPadButton.
*/
func CastToEventPadButton(object *gobject.Object) *EventPadButton {
	return EventPadButtonNewFromNative(object.Native())
}

// Equals compares this EventPadButton with another EventPadButton, and returns true if they represent the same Object.
func (recv *EventPadButton) Equals(other *EventPadButton) bool {
	return other.Native() == recv.Native()
}

func (recv *EventPadButton) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventPadButton) FieldType() EventType {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventPadButton) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventPadButtonStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventPadButton) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventPadButton) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventPadButtonStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventPadButton) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventPadButton) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventPadButtonStruct, recv.Native(), "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventPadButton) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventPadButton) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadButtonStruct, recv.Native(), "time", argValue)
}

// FieldGroup returns the C field 'group'.
func (recv *EventPadButton) FieldGroup() uint32 {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.Native(), "group")
	value := argValue.Uint32()
	return value
}

// SetFieldGroup sets the value of the C field 'group'.
func (recv *EventPadButton) SetFieldGroup(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadButtonStruct, recv.Native(), "group", argValue)
}

// FieldButton returns the C field 'button'.
func (recv *EventPadButton) FieldButton() uint32 {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.Native(), "button")
	value := argValue.Uint32()
	return value
}

// SetFieldButton sets the value of the C field 'button'.
func (recv *EventPadButton) SetFieldButton(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadButtonStruct, recv.Native(), "button", argValue)
}

// FieldMode returns the C field 'mode'.
func (recv *EventPadButton) FieldMode() uint32 {
	argValue := gi.StructFieldGet(eventPadButtonStruct, recv.Native(), "mode")
	value := argValue.Uint32()
	return value
}

// SetFieldMode sets the value of the C field 'mode'.
func (recv *EventPadButton) SetFieldMode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadButtonStruct, recv.Native(), "mode", argValue)
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
	eventPadButtonStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventPadGroupModeNewFromNative(native unsafe.Pointer) *EventPadGroupMode {
	instance := &EventPadGroupMode{native: native}

	return instance
}

/*
CastToEventPadGroupMode down casts any arbitrary Object to EventPadGroupMode.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventPadGroupMode.
*/
func CastToEventPadGroupMode(object *gobject.Object) *EventPadGroupMode {
	return EventPadGroupModeNewFromNative(object.Native())
}

// Equals compares this EventPadGroupMode with another EventPadGroupMode, and returns true if they represent the same Object.
func (recv *EventPadGroupMode) Equals(other *EventPadGroupMode) bool {
	return other.Native() == recv.Native()
}

func (recv *EventPadGroupMode) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventPadGroupMode) FieldType() EventType {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventPadGroupMode) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventPadGroupModeStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventPadGroupMode) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventPadGroupMode) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventPadGroupModeStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventPadGroupMode) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventPadGroupMode) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventPadGroupModeStruct, recv.Native(), "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventPadGroupMode) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventPadGroupMode) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadGroupModeStruct, recv.Native(), "time", argValue)
}

// FieldGroup returns the C field 'group'.
func (recv *EventPadGroupMode) FieldGroup() uint32 {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.Native(), "group")
	value := argValue.Uint32()
	return value
}

// SetFieldGroup sets the value of the C field 'group'.
func (recv *EventPadGroupMode) SetFieldGroup(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadGroupModeStruct, recv.Native(), "group", argValue)
}

// FieldMode returns the C field 'mode'.
func (recv *EventPadGroupMode) FieldMode() uint32 {
	argValue := gi.StructFieldGet(eventPadGroupModeStruct, recv.Native(), "mode")
	value := argValue.Uint32()
	return value
}

// SetFieldMode sets the value of the C field 'mode'.
func (recv *EventPadGroupMode) SetFieldMode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPadGroupModeStruct, recv.Native(), "mode", argValue)
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
	eventPadGroupModeStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventPropertyNewFromNative(native unsafe.Pointer) *EventProperty {
	instance := &EventProperty{native: native}

	return instance
}

/*
CastToEventProperty down casts any arbitrary Object to EventProperty.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventProperty.
*/
func CastToEventProperty(object *gobject.Object) *EventProperty {
	return EventPropertyNewFromNative(object.Native())
}

// Equals compares this EventProperty with another EventProperty, and returns true if they represent the same Object.
func (recv *EventProperty) Equals(other *EventProperty) bool {
	return other.Native() == recv.Native()
}

func (recv *EventProperty) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventProperty) FieldType() EventType {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventProperty) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventPropertyStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventProperty) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventProperty) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventPropertyStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventProperty) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventProperty) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventPropertyStruct, recv.Native(), "send_event", argValue)
}

// FieldAtom returns the C field 'atom'.
func (recv *EventProperty) FieldAtom() *Atom {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.Native(), "atom")
	value := AtomNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAtom sets the value of the C field 'atom'.
func (recv *EventProperty) SetFieldAtom(value *Atom) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventPropertyStruct, recv.Native(), "atom", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventProperty) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventProperty) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventPropertyStruct, recv.Native(), "time", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventProperty) FieldState() PropertyState {
	argValue := gi.StructFieldGet(eventPropertyStruct, recv.Native(), "state")
	value := PropertyState(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventProperty) SetFieldState(value PropertyState) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventPropertyStruct, recv.Native(), "state", argValue)
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
	eventPropertyStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventProximityNewFromNative(native unsafe.Pointer) *EventProximity {
	instance := &EventProximity{native: native}

	return instance
}

/*
CastToEventProximity down casts any arbitrary Object to EventProximity.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventProximity.
*/
func CastToEventProximity(object *gobject.Object) *EventProximity {
	return EventProximityNewFromNative(object.Native())
}

// Equals compares this EventProximity with another EventProximity, and returns true if they represent the same Object.
func (recv *EventProximity) Equals(other *EventProximity) bool {
	return other.Native() == recv.Native()
}

func (recv *EventProximity) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventProximity) FieldType() EventType {
	argValue := gi.StructFieldGet(eventProximityStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventProximity) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventProximityStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventProximity) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventProximityStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventProximity) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventProximityStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventProximity) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventProximityStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventProximity) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventProximityStruct, recv.Native(), "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventProximity) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventProximityStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventProximity) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventProximityStruct, recv.Native(), "time", argValue)
}

// FieldDevice returns the C field 'device'.
func (recv *EventProximity) FieldDevice() *Device {
	argValue := gi.StructFieldGet(eventProximityStruct, recv.Native(), "device")
	value := DeviceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDevice sets the value of the C field 'device'.
func (recv *EventProximity) SetFieldDevice(value *Device) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventProximityStruct, recv.Native(), "device", argValue)
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
	eventProximityStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventScrollNewFromNative(native unsafe.Pointer) *EventScroll {
	instance := &EventScroll{native: native}

	return instance
}

/*
CastToEventScroll down casts any arbitrary Object to EventScroll.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventScroll.
*/
func CastToEventScroll(object *gobject.Object) *EventScroll {
	return EventScrollNewFromNative(object.Native())
}

// Equals compares this EventScroll with another EventScroll, and returns true if they represent the same Object.
func (recv *EventScroll) Equals(other *EventScroll) bool {
	return other.Native() == recv.Native()
}

func (recv *EventScroll) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventScroll) FieldType() EventType {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventScroll) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventScroll) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventScroll) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventScroll) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventScroll) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventScroll) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventScroll) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventScroll) FieldX() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventScroll) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventScroll) FieldY() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventScroll) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "y", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventScroll) FieldState() ModifierType {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "state")
	value := ModifierType(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventScroll) SetFieldState(value ModifierType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "state", argValue)
}

// FieldDirection returns the C field 'direction'.
func (recv *EventScroll) FieldDirection() ScrollDirection {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "direction")
	value := ScrollDirection(argValue.Int32())
	return value
}

// SetFieldDirection sets the value of the C field 'direction'.
func (recv *EventScroll) SetFieldDirection(value ScrollDirection) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "direction", argValue)
}

// FieldDevice returns the C field 'device'.
func (recv *EventScroll) FieldDevice() *Device {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "device")
	value := DeviceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDevice sets the value of the C field 'device'.
func (recv *EventScroll) SetFieldDevice(value *Device) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "device", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventScroll) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventScroll) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventScroll) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventScroll) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "y_root", argValue)
}

// FieldDeltaX returns the C field 'delta_x'.
func (recv *EventScroll) FieldDeltaX() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "delta_x")
	value := argValue.Float64()
	return value
}

// SetFieldDeltaX sets the value of the C field 'delta_x'.
func (recv *EventScroll) SetFieldDeltaX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "delta_x", argValue)
}

// FieldDeltaY returns the C field 'delta_y'.
func (recv *EventScroll) FieldDeltaY() float64 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "delta_y")
	value := argValue.Float64()
	return value
}

// SetFieldDeltaY sets the value of the C field 'delta_y'.
func (recv *EventScroll) SetFieldDeltaY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "delta_y", argValue)
}

// FieldIsStop returns the C field 'is_stop'.
func (recv *EventScroll) FieldIsStop() uint32 {
	argValue := gi.StructFieldGet(eventScrollStruct, recv.Native(), "is_stop")
	value := argValue.Uint32()
	return value
}

// SetFieldIsStop sets the value of the C field 'is_stop'.
func (recv *EventScroll) SetFieldIsStop(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventScrollStruct, recv.Native(), "is_stop", argValue)
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
	eventScrollStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventSelectionNewFromNative(native unsafe.Pointer) *EventSelection {
	instance := &EventSelection{native: native}

	return instance
}

/*
CastToEventSelection down casts any arbitrary Object to EventSelection.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventSelection.
*/
func CastToEventSelection(object *gobject.Object) *EventSelection {
	return EventSelectionNewFromNative(object.Native())
}

// Equals compares this EventSelection with another EventSelection, and returns true if they represent the same Object.
func (recv *EventSelection) Equals(other *EventSelection) bool {
	return other.Native() == recv.Native()
}

func (recv *EventSelection) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventSelection) FieldType() EventType {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventSelection) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventSelectionStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventSelection) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventSelection) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventSelectionStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventSelection) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventSelection) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventSelectionStruct, recv.Native(), "send_event", argValue)
}

// FieldSelection returns the C field 'selection'.
func (recv *EventSelection) FieldSelection() *Atom {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.Native(), "selection")
	value := AtomNewFromNative(argValue.Pointer())
	return value
}

// SetFieldSelection sets the value of the C field 'selection'.
func (recv *EventSelection) SetFieldSelection(value *Atom) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventSelectionStruct, recv.Native(), "selection", argValue)
}

// FieldTarget returns the C field 'target'.
func (recv *EventSelection) FieldTarget() *Atom {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.Native(), "target")
	value := AtomNewFromNative(argValue.Pointer())
	return value
}

// SetFieldTarget sets the value of the C field 'target'.
func (recv *EventSelection) SetFieldTarget(value *Atom) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventSelectionStruct, recv.Native(), "target", argValue)
}

// FieldProperty returns the C field 'property'.
func (recv *EventSelection) FieldProperty() *Atom {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.Native(), "property")
	value := AtomNewFromNative(argValue.Pointer())
	return value
}

// SetFieldProperty sets the value of the C field 'property'.
func (recv *EventSelection) SetFieldProperty(value *Atom) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventSelectionStruct, recv.Native(), "property", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventSelection) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventSelection) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventSelectionStruct, recv.Native(), "time", argValue)
}

// FieldRequestor returns the C field 'requestor'.
func (recv *EventSelection) FieldRequestor() *Window {
	argValue := gi.StructFieldGet(eventSelectionStruct, recv.Native(), "requestor")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldRequestor sets the value of the C field 'requestor'.
func (recv *EventSelection) SetFieldRequestor(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventSelectionStruct, recv.Native(), "requestor", argValue)
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
	eventSelectionStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventSequenceNewFromNative(native unsafe.Pointer) *EventSequence {
	instance := &EventSequence{native: native}

	return instance
}

/*
CastToEventSequence down casts any arbitrary Object to EventSequence.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventSequence.
*/
func CastToEventSequence(object *gobject.Object) *EventSequence {
	return EventSequenceNewFromNative(object.Native())
}

// Equals compares this EventSequence with another EventSequence, and returns true if they represent the same Object.
func (recv *EventSequence) Equals(other *EventSequence) bool {
	return other.Native() == recv.Native()
}

func (recv *EventSequence) Native() unsafe.Pointer {
	return recv.native
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
	eventSequenceStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventSettingNewFromNative(native unsafe.Pointer) *EventSetting {
	instance := &EventSetting{native: native}

	return instance
}

/*
CastToEventSetting down casts any arbitrary Object to EventSetting.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventSetting.
*/
func CastToEventSetting(object *gobject.Object) *EventSetting {
	return EventSettingNewFromNative(object.Native())
}

// Equals compares this EventSetting with another EventSetting, and returns true if they represent the same Object.
func (recv *EventSetting) Equals(other *EventSetting) bool {
	return other.Native() == recv.Native()
}

func (recv *EventSetting) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventSetting) FieldType() EventType {
	argValue := gi.StructFieldGet(eventSettingStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventSetting) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventSettingStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventSetting) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventSettingStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventSetting) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventSettingStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventSetting) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventSettingStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventSetting) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventSettingStruct, recv.Native(), "send_event", argValue)
}

// FieldAction returns the C field 'action'.
func (recv *EventSetting) FieldAction() SettingAction {
	argValue := gi.StructFieldGet(eventSettingStruct, recv.Native(), "action")
	value := SettingAction(argValue.Int32())
	return value
}

// SetFieldAction sets the value of the C field 'action'.
func (recv *EventSetting) SetFieldAction(value SettingAction) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventSettingStruct, recv.Native(), "action", argValue)
}

// FieldName returns the C field 'name'.
func (recv *EventSetting) FieldName() string {
	argValue := gi.StructFieldGet(eventSettingStruct, recv.Native(), "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *EventSetting) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(eventSettingStruct, recv.Native(), "name", argValue)
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
	eventSettingStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventTouchNewFromNative(native unsafe.Pointer) *EventTouch {
	instance := &EventTouch{native: native}

	return instance
}

/*
CastToEventTouch down casts any arbitrary Object to EventTouch.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventTouch.
*/
func CastToEventTouch(object *gobject.Object) *EventTouch {
	return EventTouchNewFromNative(object.Native())
}

// Equals compares this EventTouch with another EventTouch, and returns true if they represent the same Object.
func (recv *EventTouch) Equals(other *EventTouch) bool {
	return other.Native() == recv.Native()
}

func (recv *EventTouch) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventTouch) FieldType() EventType {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventTouch) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventTouch) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventTouch) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventTouch) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventTouch) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "send_event", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventTouch) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventTouch) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventTouch) FieldX() float64 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventTouch) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventTouch) FieldY() float64 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventTouch) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "y", argValue)
}

// FieldAxes returns the C field 'axes'.
func (recv *EventTouch) FieldAxes() float64 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "axes")
	value := argValue.Float64()
	return value
}

// SetFieldAxes sets the value of the C field 'axes'.
func (recv *EventTouch) SetFieldAxes(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "axes", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventTouch) FieldState() ModifierType {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "state")
	value := ModifierType(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventTouch) SetFieldState(value ModifierType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "state", argValue)
}

// FieldSequence returns the C field 'sequence'.
func (recv *EventTouch) FieldSequence() *EventSequence {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "sequence")
	value := EventSequenceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldSequence sets the value of the C field 'sequence'.
func (recv *EventTouch) SetFieldSequence(value *EventSequence) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "sequence", argValue)
}

// FieldEmulatingPointer returns the C field 'emulating_pointer'.
func (recv *EventTouch) FieldEmulatingPointer() bool {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "emulating_pointer")
	value := argValue.Boolean()
	return value
}

// SetFieldEmulatingPointer sets the value of the C field 'emulating_pointer'.
func (recv *EventTouch) SetFieldEmulatingPointer(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "emulating_pointer", argValue)
}

// FieldDevice returns the C field 'device'.
func (recv *EventTouch) FieldDevice() *Device {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "device")
	value := DeviceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDevice sets the value of the C field 'device'.
func (recv *EventTouch) SetFieldDevice(value *Device) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "device", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventTouch) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventTouch) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventTouch) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchStruct, recv.Native(), "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventTouch) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchStruct, recv.Native(), "y_root", argValue)
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
	eventTouchStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventTouchpadPinchNewFromNative(native unsafe.Pointer) *EventTouchpadPinch {
	instance := &EventTouchpadPinch{native: native}

	return instance
}

/*
CastToEventTouchpadPinch down casts any arbitrary Object to EventTouchpadPinch.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventTouchpadPinch.
*/
func CastToEventTouchpadPinch(object *gobject.Object) *EventTouchpadPinch {
	return EventTouchpadPinchNewFromNative(object.Native())
}

// Equals compares this EventTouchpadPinch with another EventTouchpadPinch, and returns true if they represent the same Object.
func (recv *EventTouchpadPinch) Equals(other *EventTouchpadPinch) bool {
	return other.Native() == recv.Native()
}

func (recv *EventTouchpadPinch) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventTouchpadPinch) FieldType() EventType {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventTouchpadPinch) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventTouchpadPinch) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventTouchpadPinch) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventTouchpadPinch) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventTouchpadPinch) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "send_event", argValue)
}

// FieldPhase returns the C field 'phase'.
func (recv *EventTouchpadPinch) FieldPhase() int8 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "phase")
	value := argValue.Int8()
	return value
}

// SetFieldPhase sets the value of the C field 'phase'.
func (recv *EventTouchpadPinch) SetFieldPhase(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "phase", argValue)
}

// FieldNFingers returns the C field 'n_fingers'.
func (recv *EventTouchpadPinch) FieldNFingers() int8 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "n_fingers")
	value := argValue.Int8()
	return value
}

// SetFieldNFingers sets the value of the C field 'n_fingers'.
func (recv *EventTouchpadPinch) SetFieldNFingers(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "n_fingers", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventTouchpadPinch) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventTouchpadPinch) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventTouchpadPinch) FieldX() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventTouchpadPinch) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventTouchpadPinch) FieldY() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventTouchpadPinch) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "y", argValue)
}

// FieldDx returns the C field 'dx'.
func (recv *EventTouchpadPinch) FieldDx() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "dx")
	value := argValue.Float64()
	return value
}

// SetFieldDx sets the value of the C field 'dx'.
func (recv *EventTouchpadPinch) SetFieldDx(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "dx", argValue)
}

// FieldDy returns the C field 'dy'.
func (recv *EventTouchpadPinch) FieldDy() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "dy")
	value := argValue.Float64()
	return value
}

// SetFieldDy sets the value of the C field 'dy'.
func (recv *EventTouchpadPinch) SetFieldDy(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "dy", argValue)
}

// FieldAngleDelta returns the C field 'angle_delta'.
func (recv *EventTouchpadPinch) FieldAngleDelta() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "angle_delta")
	value := argValue.Float64()
	return value
}

// SetFieldAngleDelta sets the value of the C field 'angle_delta'.
func (recv *EventTouchpadPinch) SetFieldAngleDelta(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "angle_delta", argValue)
}

// FieldScale returns the C field 'scale'.
func (recv *EventTouchpadPinch) FieldScale() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "scale")
	value := argValue.Float64()
	return value
}

// SetFieldScale sets the value of the C field 'scale'.
func (recv *EventTouchpadPinch) SetFieldScale(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "scale", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventTouchpadPinch) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventTouchpadPinch) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventTouchpadPinch) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventTouchpadPinch) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "y_root", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventTouchpadPinch) FieldState() ModifierType {
	argValue := gi.StructFieldGet(eventTouchpadPinchStruct, recv.Native(), "state")
	value := ModifierType(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventTouchpadPinch) SetFieldState(value ModifierType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventTouchpadPinchStruct, recv.Native(), "state", argValue)
}

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
	eventTouchpadPinchStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventTouchpadSwipeNewFromNative(native unsafe.Pointer) *EventTouchpadSwipe {
	instance := &EventTouchpadSwipe{native: native}

	return instance
}

/*
CastToEventTouchpadSwipe down casts any arbitrary Object to EventTouchpadSwipe.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventTouchpadSwipe.
*/
func CastToEventTouchpadSwipe(object *gobject.Object) *EventTouchpadSwipe {
	return EventTouchpadSwipeNewFromNative(object.Native())
}

// Equals compares this EventTouchpadSwipe with another EventTouchpadSwipe, and returns true if they represent the same Object.
func (recv *EventTouchpadSwipe) Equals(other *EventTouchpadSwipe) bool {
	return other.Native() == recv.Native()
}

func (recv *EventTouchpadSwipe) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventTouchpadSwipe) FieldType() EventType {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventTouchpadSwipe) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventTouchpadSwipe) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventTouchpadSwipe) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventTouchpadSwipe) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventTouchpadSwipe) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "send_event", argValue)
}

// FieldPhase returns the C field 'phase'.
func (recv *EventTouchpadSwipe) FieldPhase() int8 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "phase")
	value := argValue.Int8()
	return value
}

// SetFieldPhase sets the value of the C field 'phase'.
func (recv *EventTouchpadSwipe) SetFieldPhase(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "phase", argValue)
}

// FieldNFingers returns the C field 'n_fingers'.
func (recv *EventTouchpadSwipe) FieldNFingers() int8 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "n_fingers")
	value := argValue.Int8()
	return value
}

// SetFieldNFingers sets the value of the C field 'n_fingers'.
func (recv *EventTouchpadSwipe) SetFieldNFingers(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "n_fingers", argValue)
}

// FieldTime returns the C field 'time'.
func (recv *EventTouchpadSwipe) FieldTime() uint32 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *EventTouchpadSwipe) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "time", argValue)
}

// FieldX returns the C field 'x'.
func (recv *EventTouchpadSwipe) FieldX() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *EventTouchpadSwipe) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *EventTouchpadSwipe) FieldY() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *EventTouchpadSwipe) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "y", argValue)
}

// FieldDx returns the C field 'dx'.
func (recv *EventTouchpadSwipe) FieldDx() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "dx")
	value := argValue.Float64()
	return value
}

// SetFieldDx sets the value of the C field 'dx'.
func (recv *EventTouchpadSwipe) SetFieldDx(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "dx", argValue)
}

// FieldDy returns the C field 'dy'.
func (recv *EventTouchpadSwipe) FieldDy() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "dy")
	value := argValue.Float64()
	return value
}

// SetFieldDy sets the value of the C field 'dy'.
func (recv *EventTouchpadSwipe) SetFieldDy(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "dy", argValue)
}

// FieldXRoot returns the C field 'x_root'.
func (recv *EventTouchpadSwipe) FieldXRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "x_root")
	value := argValue.Float64()
	return value
}

// SetFieldXRoot sets the value of the C field 'x_root'.
func (recv *EventTouchpadSwipe) SetFieldXRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "x_root", argValue)
}

// FieldYRoot returns the C field 'y_root'.
func (recv *EventTouchpadSwipe) FieldYRoot() float64 {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "y_root")
	value := argValue.Float64()
	return value
}

// SetFieldYRoot sets the value of the C field 'y_root'.
func (recv *EventTouchpadSwipe) SetFieldYRoot(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "y_root", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventTouchpadSwipe) FieldState() ModifierType {
	argValue := gi.StructFieldGet(eventTouchpadSwipeStruct, recv.Native(), "state")
	value := ModifierType(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventTouchpadSwipe) SetFieldState(value ModifierType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventTouchpadSwipeStruct, recv.Native(), "state", argValue)
}

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
	eventTouchpadSwipeStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventVisibilityNewFromNative(native unsafe.Pointer) *EventVisibility {
	instance := &EventVisibility{native: native}

	return instance
}

/*
CastToEventVisibility down casts any arbitrary Object to EventVisibility.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventVisibility.
*/
func CastToEventVisibility(object *gobject.Object) *EventVisibility {
	return EventVisibilityNewFromNative(object.Native())
}

// Equals compares this EventVisibility with another EventVisibility, and returns true if they represent the same Object.
func (recv *EventVisibility) Equals(other *EventVisibility) bool {
	return other.Native() == recv.Native()
}

func (recv *EventVisibility) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventVisibility) FieldType() EventType {
	argValue := gi.StructFieldGet(eventVisibilityStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventVisibility) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventVisibilityStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventVisibility) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventVisibilityStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventVisibility) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventVisibilityStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventVisibility) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventVisibilityStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventVisibility) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventVisibilityStruct, recv.Native(), "send_event", argValue)
}

// FieldState returns the C field 'state'.
func (recv *EventVisibility) FieldState() VisibilityState {
	argValue := gi.StructFieldGet(eventVisibilityStruct, recv.Native(), "state")
	value := VisibilityState(argValue.Int32())
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *EventVisibility) SetFieldState(value VisibilityState) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventVisibilityStruct, recv.Native(), "state", argValue)
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
	eventVisibilityStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func EventWindowStateNewFromNative(native unsafe.Pointer) *EventWindowState {
	instance := &EventWindowState{native: native}

	return instance
}

/*
CastToEventWindowState down casts any arbitrary Object to EventWindowState.
Exercise care, as this is a potentially dangerous function
if the Object is not a EventWindowState.
*/
func CastToEventWindowState(object *gobject.Object) *EventWindowState {
	return EventWindowStateNewFromNative(object.Native())
}

// Equals compares this EventWindowState with another EventWindowState, and returns true if they represent the same Object.
func (recv *EventWindowState) Equals(other *EventWindowState) bool {
	return other.Native() == recv.Native()
}

func (recv *EventWindowState) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *EventWindowState) FieldType() EventType {
	argValue := gi.StructFieldGet(eventWindowStateStruct, recv.Native(), "type")
	value := EventType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *EventWindowState) SetFieldType(value EventType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventWindowStateStruct, recv.Native(), "type", argValue)
}

// FieldWindow returns the C field 'window'.
func (recv *EventWindowState) FieldWindow() *Window {
	argValue := gi.StructFieldGet(eventWindowStateStruct, recv.Native(), "window")
	value := WindowNewFromNative(argValue.Pointer())
	return value
}

// SetFieldWindow sets the value of the C field 'window'.
func (recv *EventWindowState) SetFieldWindow(value *Window) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(eventWindowStateStruct, recv.Native(), "window", argValue)
}

// FieldSendEvent returns the C field 'send_event'.
func (recv *EventWindowState) FieldSendEvent() int8 {
	argValue := gi.StructFieldGet(eventWindowStateStruct, recv.Native(), "send_event")
	value := argValue.Int8()
	return value
}

// SetFieldSendEvent sets the value of the C field 'send_event'.
func (recv *EventWindowState) SetFieldSendEvent(value int8) {
	var argValue gi.Argument
	argValue.SetInt8(value)
	gi.StructFieldSet(eventWindowStateStruct, recv.Native(), "send_event", argValue)
}

// FieldChangedMask returns the C field 'changed_mask'.
func (recv *EventWindowState) FieldChangedMask() WindowState {
	argValue := gi.StructFieldGet(eventWindowStateStruct, recv.Native(), "changed_mask")
	value := WindowState(argValue.Int32())
	return value
}

// SetFieldChangedMask sets the value of the C field 'changed_mask'.
func (recv *EventWindowState) SetFieldChangedMask(value WindowState) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventWindowStateStruct, recv.Native(), "changed_mask", argValue)
}

// FieldNewWindowState returns the C field 'new_window_state'.
func (recv *EventWindowState) FieldNewWindowState() WindowState {
	argValue := gi.StructFieldGet(eventWindowStateStruct, recv.Native(), "new_window_state")
	value := WindowState(argValue.Int32())
	return value
}

// SetFieldNewWindowState sets the value of the C field 'new_window_state'.
func (recv *EventWindowState) SetFieldNewWindowState(value WindowState) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(eventWindowStateStruct, recv.Native(), "new_window_state", argValue)
}

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
	eventWindowStateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func FrameClockClassNewFromNative(native unsafe.Pointer) *FrameClockClass {
	instance := &FrameClockClass{native: native}

	return instance
}

/*
CastToFrameClockClass down casts any arbitrary Object to FrameClockClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FrameClockClass.
*/
func CastToFrameClockClass(object *gobject.Object) *FrameClockClass {
	return FrameClockClassNewFromNative(object.Native())
}

// Equals compares this FrameClockClass with another FrameClockClass, and returns true if they represent the same Object.
func (recv *FrameClockClass) Equals(other *FrameClockClass) bool {
	return other.Native() == recv.Native()
}

func (recv *FrameClockClass) Native() unsafe.Pointer {
	return recv.native
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
	frameClockClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func FrameClockPrivateNewFromNative(native unsafe.Pointer) *FrameClockPrivate {
	instance := &FrameClockPrivate{native: native}

	return instance
}

/*
CastToFrameClockPrivate down casts any arbitrary Object to FrameClockPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a FrameClockPrivate.
*/
func CastToFrameClockPrivate(object *gobject.Object) *FrameClockPrivate {
	return FrameClockPrivateNewFromNative(object.Native())
}

// Equals compares this FrameClockPrivate with another FrameClockPrivate, and returns true if they represent the same Object.
func (recv *FrameClockPrivate) Equals(other *FrameClockPrivate) bool {
	return other.Native() == recv.Native()
}

func (recv *FrameClockPrivate) Native() unsafe.Pointer {
	return recv.native
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
	frameClockPrivateStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func FrameTimingsNewFromNative(native unsafe.Pointer) *FrameTimings {
	instance := &FrameTimings{native: native}

	return instance
}

/*
CastToFrameTimings down casts any arbitrary Object to FrameTimings.
Exercise care, as this is a potentially dangerous function
if the Object is not a FrameTimings.
*/
func CastToFrameTimings(object *gobject.Object) *FrameTimings {
	return FrameTimingsNewFromNative(object.Native())
}

// Equals compares this FrameTimings with another FrameTimings, and returns true if they represent the same Object.
func (recv *FrameTimings) Equals(other *FrameTimings) bool {
	return other.Native() == recv.Native()
}

func (recv *FrameTimings) Native() unsafe.Pointer {
	return recv.native
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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	frameTimingsStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func GeometryNewFromNative(native unsafe.Pointer) *Geometry {
	instance := &Geometry{native: native}

	return instance
}

/*
CastToGeometry down casts any arbitrary Object to Geometry.
Exercise care, as this is a potentially dangerous function
if the Object is not a Geometry.
*/
func CastToGeometry(object *gobject.Object) *Geometry {
	return GeometryNewFromNative(object.Native())
}

// Equals compares this Geometry with another Geometry, and returns true if they represent the same Object.
func (recv *Geometry) Equals(other *Geometry) bool {
	return other.Native() == recv.Native()
}

func (recv *Geometry) Native() unsafe.Pointer {
	return recv.native
}

// FieldMinWidth returns the C field 'min_width'.
func (recv *Geometry) FieldMinWidth() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "min_width")
	value := argValue.Int32()
	return value
}

// SetFieldMinWidth sets the value of the C field 'min_width'.
func (recv *Geometry) SetFieldMinWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.Native(), "min_width", argValue)
}

// FieldMinHeight returns the C field 'min_height'.
func (recv *Geometry) FieldMinHeight() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "min_height")
	value := argValue.Int32()
	return value
}

// SetFieldMinHeight sets the value of the C field 'min_height'.
func (recv *Geometry) SetFieldMinHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.Native(), "min_height", argValue)
}

// FieldMaxWidth returns the C field 'max_width'.
func (recv *Geometry) FieldMaxWidth() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "max_width")
	value := argValue.Int32()
	return value
}

// SetFieldMaxWidth sets the value of the C field 'max_width'.
func (recv *Geometry) SetFieldMaxWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.Native(), "max_width", argValue)
}

// FieldMaxHeight returns the C field 'max_height'.
func (recv *Geometry) FieldMaxHeight() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "max_height")
	value := argValue.Int32()
	return value
}

// SetFieldMaxHeight sets the value of the C field 'max_height'.
func (recv *Geometry) SetFieldMaxHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.Native(), "max_height", argValue)
}

// FieldBaseWidth returns the C field 'base_width'.
func (recv *Geometry) FieldBaseWidth() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "base_width")
	value := argValue.Int32()
	return value
}

// SetFieldBaseWidth sets the value of the C field 'base_width'.
func (recv *Geometry) SetFieldBaseWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.Native(), "base_width", argValue)
}

// FieldBaseHeight returns the C field 'base_height'.
func (recv *Geometry) FieldBaseHeight() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "base_height")
	value := argValue.Int32()
	return value
}

// SetFieldBaseHeight sets the value of the C field 'base_height'.
func (recv *Geometry) SetFieldBaseHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.Native(), "base_height", argValue)
}

// FieldWidthInc returns the C field 'width_inc'.
func (recv *Geometry) FieldWidthInc() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "width_inc")
	value := argValue.Int32()
	return value
}

// SetFieldWidthInc sets the value of the C field 'width_inc'.
func (recv *Geometry) SetFieldWidthInc(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.Native(), "width_inc", argValue)
}

// FieldHeightInc returns the C field 'height_inc'.
func (recv *Geometry) FieldHeightInc() int32 {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "height_inc")
	value := argValue.Int32()
	return value
}

// SetFieldHeightInc sets the value of the C field 'height_inc'.
func (recv *Geometry) SetFieldHeightInc(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(geometryStruct, recv.Native(), "height_inc", argValue)
}

// FieldMinAspect returns the C field 'min_aspect'.
func (recv *Geometry) FieldMinAspect() float64 {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "min_aspect")
	value := argValue.Float64()
	return value
}

// SetFieldMinAspect sets the value of the C field 'min_aspect'.
func (recv *Geometry) SetFieldMinAspect(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(geometryStruct, recv.Native(), "min_aspect", argValue)
}

// FieldMaxAspect returns the C field 'max_aspect'.
func (recv *Geometry) FieldMaxAspect() float64 {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "max_aspect")
	value := argValue.Float64()
	return value
}

// SetFieldMaxAspect sets the value of the C field 'max_aspect'.
func (recv *Geometry) SetFieldMaxAspect(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(geometryStruct, recv.Native(), "max_aspect", argValue)
}

// FieldWinGravity returns the C field 'win_gravity'.
func (recv *Geometry) FieldWinGravity() Gravity {
	argValue := gi.StructFieldGet(geometryStruct, recv.Native(), "win_gravity")
	value := Gravity(argValue.Int32())
	return value
}

// SetFieldWinGravity sets the value of the C field 'win_gravity'.
func (recv *Geometry) SetFieldWinGravity(value Gravity) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(geometryStruct, recv.Native(), "win_gravity", argValue)
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
	geometryStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func KeymapKeyNewFromNative(native unsafe.Pointer) *KeymapKey {
	instance := &KeymapKey{native: native}

	return instance
}

/*
CastToKeymapKey down casts any arbitrary Object to KeymapKey.
Exercise care, as this is a potentially dangerous function
if the Object is not a KeymapKey.
*/
func CastToKeymapKey(object *gobject.Object) *KeymapKey {
	return KeymapKeyNewFromNative(object.Native())
}

// Equals compares this KeymapKey with another KeymapKey, and returns true if they represent the same Object.
func (recv *KeymapKey) Equals(other *KeymapKey) bool {
	return other.Native() == recv.Native()
}

func (recv *KeymapKey) Native() unsafe.Pointer {
	return recv.native
}

// FieldKeycode returns the C field 'keycode'.
func (recv *KeymapKey) FieldKeycode() uint32 {
	argValue := gi.StructFieldGet(keymapKeyStruct, recv.Native(), "keycode")
	value := argValue.Uint32()
	return value
}

// SetFieldKeycode sets the value of the C field 'keycode'.
func (recv *KeymapKey) SetFieldKeycode(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(keymapKeyStruct, recv.Native(), "keycode", argValue)
}

// FieldGroup returns the C field 'group'.
func (recv *KeymapKey) FieldGroup() int32 {
	argValue := gi.StructFieldGet(keymapKeyStruct, recv.Native(), "group")
	value := argValue.Int32()
	return value
}

// SetFieldGroup sets the value of the C field 'group'.
func (recv *KeymapKey) SetFieldGroup(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(keymapKeyStruct, recv.Native(), "group", argValue)
}

// FieldLevel returns the C field 'level'.
func (recv *KeymapKey) FieldLevel() int32 {
	argValue := gi.StructFieldGet(keymapKeyStruct, recv.Native(), "level")
	value := argValue.Int32()
	return value
}

// SetFieldLevel sets the value of the C field 'level'.
func (recv *KeymapKey) SetFieldLevel(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(keymapKeyStruct, recv.Native(), "level", argValue)
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
	keymapKeyStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func MonitorClassNewFromNative(native unsafe.Pointer) *MonitorClass {
	instance := &MonitorClass{native: native}

	return instance
}

/*
CastToMonitorClass down casts any arbitrary Object to MonitorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a MonitorClass.
*/
func CastToMonitorClass(object *gobject.Object) *MonitorClass {
	return MonitorClassNewFromNative(object.Native())
}

// Equals compares this MonitorClass with another MonitorClass, and returns true if they represent the same Object.
func (recv *MonitorClass) Equals(other *MonitorClass) bool {
	return other.Native() == recv.Native()
}

func (recv *MonitorClass) Native() unsafe.Pointer {
	return recv.native
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
	monitorClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func PointNewFromNative(native unsafe.Pointer) *Point {
	instance := &Point{native: native}

	return instance
}

/*
CastToPoint down casts any arbitrary Object to Point.
Exercise care, as this is a potentially dangerous function
if the Object is not a Point.
*/
func CastToPoint(object *gobject.Object) *Point {
	return PointNewFromNative(object.Native())
}

// Equals compares this Point with another Point, and returns true if they represent the same Object.
func (recv *Point) Equals(other *Point) bool {
	return other.Native() == recv.Native()
}

func (recv *Point) Native() unsafe.Pointer {
	return recv.native
}

// FieldX returns the C field 'x'.
func (recv *Point) FieldX() int32 {
	argValue := gi.StructFieldGet(pointStruct, recv.Native(), "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *Point) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(pointStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *Point) FieldY() int32 {
	argValue := gi.StructFieldGet(pointStruct, recv.Native(), "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *Point) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(pointStruct, recv.Native(), "y", argValue)
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
	pointStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RGBANewFromNative(native unsafe.Pointer) *RGBA {
	instance := &RGBA{native: native}

	return instance
}

/*
CastToRGBA down casts any arbitrary Object to RGBA.
Exercise care, as this is a potentially dangerous function
if the Object is not a RGBA.
*/
func CastToRGBA(object *gobject.Object) *RGBA {
	return RGBANewFromNative(object.Native())
}

// Equals compares this RGBA with another RGBA, and returns true if they represent the same Object.
func (recv *RGBA) Equals(other *RGBA) bool {
	return other.Native() == recv.Native()
}

func (recv *RGBA) Native() unsafe.Pointer {
	return recv.native
}

// FieldRed returns the C field 'red'.
func (recv *RGBA) FieldRed() float64 {
	argValue := gi.StructFieldGet(rGBAStruct, recv.Native(), "red")
	value := argValue.Float64()
	return value
}

// SetFieldRed sets the value of the C field 'red'.
func (recv *RGBA) SetFieldRed(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rGBAStruct, recv.Native(), "red", argValue)
}

// FieldGreen returns the C field 'green'.
func (recv *RGBA) FieldGreen() float64 {
	argValue := gi.StructFieldGet(rGBAStruct, recv.Native(), "green")
	value := argValue.Float64()
	return value
}

// SetFieldGreen sets the value of the C field 'green'.
func (recv *RGBA) SetFieldGreen(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rGBAStruct, recv.Native(), "green", argValue)
}

// FieldBlue returns the C field 'blue'.
func (recv *RGBA) FieldBlue() float64 {
	argValue := gi.StructFieldGet(rGBAStruct, recv.Native(), "blue")
	value := argValue.Float64()
	return value
}

// SetFieldBlue sets the value of the C field 'blue'.
func (recv *RGBA) SetFieldBlue(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rGBAStruct, recv.Native(), "blue", argValue)
}

// FieldAlpha returns the C field 'alpha'.
func (recv *RGBA) FieldAlpha() float64 {
	argValue := gi.StructFieldGet(rGBAStruct, recv.Native(), "alpha")
	value := argValue.Float64()
	return value
}

// SetFieldAlpha sets the value of the C field 'alpha'.
func (recv *RGBA) SetFieldAlpha(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rGBAStruct, recv.Native(), "alpha", argValue)
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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(p2.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())

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
	inArgs[0].SetPointer(recv.Native())
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
	inArgs[0].SetPointer(recv.Native())

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
	rGBAStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func RectangleNewFromNative(native unsafe.Pointer) *Rectangle {
	instance := &Rectangle{native: native}

	return instance
}

/*
CastToRectangle down casts any arbitrary Object to Rectangle.
Exercise care, as this is a potentially dangerous function
if the Object is not a Rectangle.
*/
func CastToRectangle(object *gobject.Object) *Rectangle {
	return RectangleNewFromNative(object.Native())
}

// Equals compares this Rectangle with another Rectangle, and returns true if they represent the same Object.
func (recv *Rectangle) Equals(other *Rectangle) bool {
	return other.Native() == recv.Native()
}

func (recv *Rectangle) Native() unsafe.Pointer {
	return recv.native
}

// FieldX returns the C field 'x'.
func (recv *Rectangle) FieldX() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.Native(), "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *Rectangle) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *Rectangle) FieldY() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.Native(), "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *Rectangle) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.Native(), "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *Rectangle) FieldWidth() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.Native(), "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *Rectangle) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.Native(), "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *Rectangle) FieldHeight() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.Native(), "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *Rectangle) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.Native(), "height", argValue)
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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(rect2.Native())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(src2.Native())

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
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(src2.Native())

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
	rectangleStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func TimeCoordNewFromNative(native unsafe.Pointer) *TimeCoord {
	instance := &TimeCoord{native: native}

	return instance
}

/*
CastToTimeCoord down casts any arbitrary Object to TimeCoord.
Exercise care, as this is a potentially dangerous function
if the Object is not a TimeCoord.
*/
func CastToTimeCoord(object *gobject.Object) *TimeCoord {
	return TimeCoordNewFromNative(object.Native())
}

// Equals compares this TimeCoord with another TimeCoord, and returns true if they represent the same Object.
func (recv *TimeCoord) Equals(other *TimeCoord) bool {
	return other.Native() == recv.Native()
}

func (recv *TimeCoord) Native() unsafe.Pointer {
	return recv.native
}

// FieldTime returns the C field 'time'.
func (recv *TimeCoord) FieldTime() uint32 {
	argValue := gi.StructFieldGet(timeCoordStruct, recv.Native(), "time")
	value := argValue.Uint32()
	return value
}

// SetFieldTime sets the value of the C field 'time'.
func (recv *TimeCoord) SetFieldTime(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(timeCoordStruct, recv.Native(), "time", argValue)
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
	timeCoordStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func WindowAttrNewFromNative(native unsafe.Pointer) *WindowAttr {
	instance := &WindowAttr{native: native}

	return instance
}

/*
CastToWindowAttr down casts any arbitrary Object to WindowAttr.
Exercise care, as this is a potentially dangerous function
if the Object is not a WindowAttr.
*/
func CastToWindowAttr(object *gobject.Object) *WindowAttr {
	return WindowAttrNewFromNative(object.Native())
}

// Equals compares this WindowAttr with another WindowAttr, and returns true if they represent the same Object.
func (recv *WindowAttr) Equals(other *WindowAttr) bool {
	return other.Native() == recv.Native()
}

func (recv *WindowAttr) Native() unsafe.Pointer {
	return recv.native
}

// FieldTitle returns the C field 'title'.
func (recv *WindowAttr) FieldTitle() string {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "title")
	value := argValue.String(false)
	return value
}

// SetFieldTitle sets the value of the C field 'title'.
func (recv *WindowAttr) SetFieldTitle(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "title", argValue)
}

// FieldEventMask returns the C field 'event_mask'.
func (recv *WindowAttr) FieldEventMask() int32 {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "event_mask")
	value := argValue.Int32()
	return value
}

// SetFieldEventMask sets the value of the C field 'event_mask'.
func (recv *WindowAttr) SetFieldEventMask(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "event_mask", argValue)
}

// FieldX returns the C field 'x'.
func (recv *WindowAttr) FieldX() int32 {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *WindowAttr) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *WindowAttr) FieldY() int32 {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *WindowAttr) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *WindowAttr) FieldWidth() int32 {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *WindowAttr) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *WindowAttr) FieldHeight() int32 {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *WindowAttr) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "height", argValue)
}

// FieldWclass returns the C field 'wclass'.
func (recv *WindowAttr) FieldWclass() WindowWindowClass {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "wclass")
	value := WindowWindowClass(argValue.Int32())
	return value
}

// SetFieldWclass sets the value of the C field 'wclass'.
func (recv *WindowAttr) SetFieldWclass(value WindowWindowClass) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "wclass", argValue)
}

// FieldVisual returns the C field 'visual'.
func (recv *WindowAttr) FieldVisual() *Visual {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "visual")
	value := VisualNewFromNative(argValue.Pointer())
	return value
}

// SetFieldVisual sets the value of the C field 'visual'.
func (recv *WindowAttr) SetFieldVisual(value *Visual) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "visual", argValue)
}

// FieldWindowType returns the C field 'window_type'.
func (recv *WindowAttr) FieldWindowType() WindowType {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "window_type")
	value := WindowType(argValue.Int32())
	return value
}

// SetFieldWindowType sets the value of the C field 'window_type'.
func (recv *WindowAttr) SetFieldWindowType(value WindowType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "window_type", argValue)
}

// FieldCursor returns the C field 'cursor'.
func (recv *WindowAttr) FieldCursor() *Cursor {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "cursor")
	value := CursorNewFromNative(argValue.Pointer())
	return value
}

// SetFieldCursor sets the value of the C field 'cursor'.
func (recv *WindowAttr) SetFieldCursor(value *Cursor) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "cursor", argValue)
}

// FieldWmclassName returns the C field 'wmclass_name'.
func (recv *WindowAttr) FieldWmclassName() string {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "wmclass_name")
	value := argValue.String(false)
	return value
}

// SetFieldWmclassName sets the value of the C field 'wmclass_name'.
func (recv *WindowAttr) SetFieldWmclassName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "wmclass_name", argValue)
}

// FieldWmclassClass returns the C field 'wmclass_class'.
func (recv *WindowAttr) FieldWmclassClass() string {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "wmclass_class")
	value := argValue.String(false)
	return value
}

// SetFieldWmclassClass sets the value of the C field 'wmclass_class'.
func (recv *WindowAttr) SetFieldWmclassClass(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "wmclass_class", argValue)
}

// FieldOverrideRedirect returns the C field 'override_redirect'.
func (recv *WindowAttr) FieldOverrideRedirect() bool {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "override_redirect")
	value := argValue.Boolean()
	return value
}

// SetFieldOverrideRedirect sets the value of the C field 'override_redirect'.
func (recv *WindowAttr) SetFieldOverrideRedirect(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "override_redirect", argValue)
}

// FieldTypeHint returns the C field 'type_hint'.
func (recv *WindowAttr) FieldTypeHint() WindowTypeHint {
	argValue := gi.StructFieldGet(windowAttrStruct, recv.Native(), "type_hint")
	value := WindowTypeHint(argValue.Int32())
	return value
}

// SetFieldTypeHint sets the value of the C field 'type_hint'.
func (recv *WindowAttr) SetFieldTypeHint(value WindowTypeHint) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(windowAttrStruct, recv.Native(), "type_hint", argValue)
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
	windowAttrStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func WindowClassNewFromNative(native unsafe.Pointer) *WindowClass {
	instance := &WindowClass{native: native}

	return instance
}

/*
CastToWindowClass down casts any arbitrary Object to WindowClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WindowClass.
*/
func CastToWindowClass(object *gobject.Object) *WindowClass {
	return WindowClassNewFromNative(object.Native())
}

// Equals compares this WindowClass with another WindowClass, and returns true if they represent the same Object.
func (recv *WindowClass) Equals(other *WindowClass) bool {
	return other.Native() == recv.Native()
}

func (recv *WindowClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *WindowClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(windowClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *WindowClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(windowClassStruct, recv.Native(), "parent_class", argValue)
}

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
	windowClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func WindowRedirectNewFromNative(native unsafe.Pointer) *WindowRedirect {
	instance := &WindowRedirect{native: native}

	return instance
}

/*
CastToWindowRedirect down casts any arbitrary Object to WindowRedirect.
Exercise care, as this is a potentially dangerous function
if the Object is not a WindowRedirect.
*/
func CastToWindowRedirect(object *gobject.Object) *WindowRedirect {
	return WindowRedirectNewFromNative(object.Native())
}

// Equals compares this WindowRedirect with another WindowRedirect, and returns true if they represent the same Object.
func (recv *WindowRedirect) Equals(other *WindowRedirect) bool {
	return other.Native() == recv.Native()
}

func (recv *WindowRedirect) Native() unsafe.Pointer {
	return recv.native
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
	windowRedirectStruct.Free(obj.Native())
}
