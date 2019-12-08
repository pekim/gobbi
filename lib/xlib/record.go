// Code generated - DO NOT EDIT.

package xlib

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

var displayStruct *gi.Struct
var displayStruct_Once sync.Once

func displayStruct_Set() error {
	var err error
	displayStruct_Once.Do(func() {
		displayStruct, err = gi.StructNew("xlib", "Display")
	})
	return err
}

type Display struct {
	native unsafe.Pointer
}

func DisplayNewFromNative(native unsafe.Pointer) *Display {
	instance := &Display{native: native}

	return instance
}

/*
CastToDisplay down casts any arbitrary Object to Display.
Exercise care, as this is a potentially dangerous function
if the Object is not a Display.
*/
func CastToDisplay(object *gobject.Object) *Display {
	return DisplayNewFromNative(object.Native())
}

func (recv *Display) Native() unsafe.Pointer {
	return recv.native
}

// DisplayStruct creates an uninitialised Display.
func DisplayStruct() *Display {
	err := displayStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DisplayNewFromNative(displayStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDisplay)
	return structGo
}
func finalizeDisplay(obj *Display) {
	displayStruct.Free(obj.Native())
}

var screenStruct *gi.Struct
var screenStruct_Once sync.Once

func screenStruct_Set() error {
	var err error
	screenStruct_Once.Do(func() {
		screenStruct, err = gi.StructNew("xlib", "Screen")
	})
	return err
}

type Screen struct {
	native unsafe.Pointer
}

func ScreenNewFromNative(native unsafe.Pointer) *Screen {
	instance := &Screen{native: native}

	return instance
}

/*
CastToScreen down casts any arbitrary Object to Screen.
Exercise care, as this is a potentially dangerous function
if the Object is not a Screen.
*/
func CastToScreen(object *gobject.Object) *Screen {
	return ScreenNewFromNative(object.Native())
}

func (recv *Screen) Native() unsafe.Pointer {
	return recv.native
}

// ScreenStruct creates an uninitialised Screen.
func ScreenStruct() *Screen {
	err := screenStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ScreenNewFromNative(screenStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeScreen)
	return structGo
}
func finalizeScreen(obj *Screen) {
	screenStruct.Free(obj.Native())
}

var visualStruct *gi.Struct
var visualStruct_Once sync.Once

func visualStruct_Set() error {
	var err error
	visualStruct_Once.Do(func() {
		visualStruct, err = gi.StructNew("xlib", "Visual")
	})
	return err
}

type Visual struct {
	native unsafe.Pointer
}

func VisualNewFromNative(native unsafe.Pointer) *Visual {
	instance := &Visual{native: native}

	return instance
}

/*
CastToVisual down casts any arbitrary Object to Visual.
Exercise care, as this is a potentially dangerous function
if the Object is not a Visual.
*/
func CastToVisual(object *gobject.Object) *Visual {
	return VisualNewFromNative(object.Native())
}

func (recv *Visual) Native() unsafe.Pointer {
	return recv.native
}

// VisualStruct creates an uninitialised Visual.
func VisualStruct() *Visual {
	err := visualStruct_Set()
	if err != nil {
		return nil
	}

	structGo := VisualNewFromNative(visualStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeVisual)
	return structGo
}
func finalizeVisual(obj *Visual) {
	visualStruct.Free(obj.Native())
}

var xConfigureEventStruct *gi.Struct
var xConfigureEventStruct_Once sync.Once

func xConfigureEventStruct_Set() error {
	var err error
	xConfigureEventStruct_Once.Do(func() {
		xConfigureEventStruct, err = gi.StructNew("xlib", "XConfigureEvent")
	})
	return err
}

type XConfigureEvent struct {
	native unsafe.Pointer
}

func XConfigureEventNewFromNative(native unsafe.Pointer) *XConfigureEvent {
	instance := &XConfigureEvent{native: native}

	return instance
}

/*
CastToXConfigureEvent down casts any arbitrary Object to XConfigureEvent.
Exercise care, as this is a potentially dangerous function
if the Object is not a XConfigureEvent.
*/
func CastToXConfigureEvent(object *gobject.Object) *XConfigureEvent {
	return XConfigureEventNewFromNative(object.Native())
}

func (recv *XConfigureEvent) Native() unsafe.Pointer {
	return recv.native
}

// XConfigureEventStruct creates an uninitialised XConfigureEvent.
func XConfigureEventStruct() *XConfigureEvent {
	err := xConfigureEventStruct_Set()
	if err != nil {
		return nil
	}

	structGo := XConfigureEventNewFromNative(xConfigureEventStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeXConfigureEvent)
	return structGo
}
func finalizeXConfigureEvent(obj *XConfigureEvent) {
	xConfigureEventStruct.Free(obj.Native())
}

var xImageStruct *gi.Struct
var xImageStruct_Once sync.Once

func xImageStruct_Set() error {
	var err error
	xImageStruct_Once.Do(func() {
		xImageStruct, err = gi.StructNew("xlib", "XImage")
	})
	return err
}

type XImage struct {
	native unsafe.Pointer
}

func XImageNewFromNative(native unsafe.Pointer) *XImage {
	instance := &XImage{native: native}

	return instance
}

/*
CastToXImage down casts any arbitrary Object to XImage.
Exercise care, as this is a potentially dangerous function
if the Object is not a XImage.
*/
func CastToXImage(object *gobject.Object) *XImage {
	return XImageNewFromNative(object.Native())
}

func (recv *XImage) Native() unsafe.Pointer {
	return recv.native
}

// XImageStruct creates an uninitialised XImage.
func XImageStruct() *XImage {
	err := xImageStruct_Set()
	if err != nil {
		return nil
	}

	structGo := XImageNewFromNative(xImageStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeXImage)
	return structGo
}
func finalizeXImage(obj *XImage) {
	xImageStruct.Free(obj.Native())
}

var xFontStructStruct *gi.Struct
var xFontStructStruct_Once sync.Once

func xFontStructStruct_Set() error {
	var err error
	xFontStructStruct_Once.Do(func() {
		xFontStructStruct, err = gi.StructNew("xlib", "XFontStruct")
	})
	return err
}

type XFontStruct struct {
	native unsafe.Pointer
}

func XFontStructNewFromNative(native unsafe.Pointer) *XFontStruct {
	instance := &XFontStruct{native: native}

	return instance
}

/*
CastToXFontStruct down casts any arbitrary Object to XFontStruct.
Exercise care, as this is a potentially dangerous function
if the Object is not a XFontStruct.
*/
func CastToXFontStruct(object *gobject.Object) *XFontStruct {
	return XFontStructNewFromNative(object.Native())
}

func (recv *XFontStruct) Native() unsafe.Pointer {
	return recv.native
}

// XFontStructStruct creates an uninitialised XFontStruct.
func XFontStructStruct() *XFontStruct {
	err := xFontStructStruct_Set()
	if err != nil {
		return nil
	}

	structGo := XFontStructNewFromNative(xFontStructStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeXFontStruct)
	return structGo
}
func finalizeXFontStruct(obj *XFontStruct) {
	xFontStructStruct.Free(obj.Native())
}

var xTrapezoidStruct *gi.Struct
var xTrapezoidStruct_Once sync.Once

func xTrapezoidStruct_Set() error {
	var err error
	xTrapezoidStruct_Once.Do(func() {
		xTrapezoidStruct, err = gi.StructNew("xlib", "XTrapezoid")
	})
	return err
}

type XTrapezoid struct {
	native unsafe.Pointer
}

func XTrapezoidNewFromNative(native unsafe.Pointer) *XTrapezoid {
	instance := &XTrapezoid{native: native}

	return instance
}

/*
CastToXTrapezoid down casts any arbitrary Object to XTrapezoid.
Exercise care, as this is a potentially dangerous function
if the Object is not a XTrapezoid.
*/
func CastToXTrapezoid(object *gobject.Object) *XTrapezoid {
	return XTrapezoidNewFromNative(object.Native())
}

func (recv *XTrapezoid) Native() unsafe.Pointer {
	return recv.native
}

// XTrapezoidStruct creates an uninitialised XTrapezoid.
func XTrapezoidStruct() *XTrapezoid {
	err := xTrapezoidStruct_Set()
	if err != nil {
		return nil
	}

	structGo := XTrapezoidNewFromNative(xTrapezoidStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeXTrapezoid)
	return structGo
}
func finalizeXTrapezoid(obj *XTrapezoid) {
	xTrapezoidStruct.Free(obj.Native())
}

var xVisualInfoStruct *gi.Struct
var xVisualInfoStruct_Once sync.Once

func xVisualInfoStruct_Set() error {
	var err error
	xVisualInfoStruct_Once.Do(func() {
		xVisualInfoStruct, err = gi.StructNew("xlib", "XVisualInfo")
	})
	return err
}

type XVisualInfo struct {
	native unsafe.Pointer
}

func XVisualInfoNewFromNative(native unsafe.Pointer) *XVisualInfo {
	instance := &XVisualInfo{native: native}

	return instance
}

/*
CastToXVisualInfo down casts any arbitrary Object to XVisualInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a XVisualInfo.
*/
func CastToXVisualInfo(object *gobject.Object) *XVisualInfo {
	return XVisualInfoNewFromNative(object.Native())
}

func (recv *XVisualInfo) Native() unsafe.Pointer {
	return recv.native
}

// XVisualInfoStruct creates an uninitialised XVisualInfo.
func XVisualInfoStruct() *XVisualInfo {
	err := xVisualInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := XVisualInfoNewFromNative(xVisualInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeXVisualInfo)
	return structGo
}
func finalizeXVisualInfo(obj *XVisualInfo) {
	xVisualInfoStruct.Free(obj.Native())
}

var xWindowAttributesStruct *gi.Struct
var xWindowAttributesStruct_Once sync.Once

func xWindowAttributesStruct_Set() error {
	var err error
	xWindowAttributesStruct_Once.Do(func() {
		xWindowAttributesStruct, err = gi.StructNew("xlib", "XWindowAttributes")
	})
	return err
}

type XWindowAttributes struct {
	native unsafe.Pointer
}

func XWindowAttributesNewFromNative(native unsafe.Pointer) *XWindowAttributes {
	instance := &XWindowAttributes{native: native}

	return instance
}

/*
CastToXWindowAttributes down casts any arbitrary Object to XWindowAttributes.
Exercise care, as this is a potentially dangerous function
if the Object is not a XWindowAttributes.
*/
func CastToXWindowAttributes(object *gobject.Object) *XWindowAttributes {
	return XWindowAttributesNewFromNative(object.Native())
}

func (recv *XWindowAttributes) Native() unsafe.Pointer {
	return recv.native
}

// XWindowAttributesStruct creates an uninitialised XWindowAttributes.
func XWindowAttributesStruct() *XWindowAttributes {
	err := xWindowAttributesStruct_Set()
	if err != nil {
		return nil
	}

	structGo := XWindowAttributesNewFromNative(xWindowAttributesStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeXWindowAttributes)
	return structGo
}
func finalizeXWindowAttributes(obj *XWindowAttributes) {
	xWindowAttributesStruct.Free(obj.Native())
}
