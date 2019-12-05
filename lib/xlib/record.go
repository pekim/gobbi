// Code generated - DO NOT EDIT.

package xlib

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
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
	native uintptr
}

func DisplayNewFromNative(native uintptr) *Display {
	return &Display{native: native}
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
	displayStruct.Free(obj.native)
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
	native uintptr
}

func ScreenNewFromNative(native uintptr) *Screen {
	return &Screen{native: native}
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
	screenStruct.Free(obj.native)
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
	native uintptr
}

func VisualNewFromNative(native uintptr) *Visual {
	return &Visual{native: native}
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
	visualStruct.Free(obj.native)
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
	native uintptr
}

func XConfigureEventNewFromNative(native uintptr) *XConfigureEvent {
	return &XConfigureEvent{native: native}
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
	xConfigureEventStruct.Free(obj.native)
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
	native uintptr
}

func XImageNewFromNative(native uintptr) *XImage {
	return &XImage{native: native}
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
	xImageStruct.Free(obj.native)
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
	native uintptr
}

func XFontStructNewFromNative(native uintptr) *XFontStruct {
	return &XFontStruct{native: native}
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
	xFontStructStruct.Free(obj.native)
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
	native uintptr
}

func XTrapezoidNewFromNative(native uintptr) *XTrapezoid {
	return &XTrapezoid{native: native}
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
	xTrapezoidStruct.Free(obj.native)
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
	native uintptr
}

func XVisualInfoNewFromNative(native uintptr) *XVisualInfo {
	return &XVisualInfo{native: native}
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
	xVisualInfoStruct.Free(obj.native)
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
	native uintptr
}

func XWindowAttributesNewFromNative(native uintptr) *XWindowAttributes {
	return &XWindowAttributes{native: native}
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
	xWindowAttributesStruct.Free(obj.native)
}
