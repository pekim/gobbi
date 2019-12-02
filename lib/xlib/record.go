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
	Native uintptr
}

// DisplayStruct creates an uninitialised Display.
func DisplayStruct() *Display {
	err := displayStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Display{}
	structGo.Native = displayStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDisplay)
	return structGo
}
func finalizeDisplay(obj *Display) {
	displayStruct.Free(obj.Native)
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
	Native uintptr
}

// ScreenStruct creates an uninitialised Screen.
func ScreenStruct() *Screen {
	err := screenStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Screen{}
	structGo.Native = screenStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeScreen)
	return structGo
}
func finalizeScreen(obj *Screen) {
	screenStruct.Free(obj.Native)
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
	Native uintptr
}

// VisualStruct creates an uninitialised Visual.
func VisualStruct() *Visual {
	err := visualStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Visual{}
	structGo.Native = visualStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeVisual)
	return structGo
}
func finalizeVisual(obj *Visual) {
	visualStruct.Free(obj.Native)
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
	Native uintptr
}

// XConfigureEventStruct creates an uninitialised XConfigureEvent.
func XConfigureEventStruct() *XConfigureEvent {
	err := xConfigureEventStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &XConfigureEvent{}
	structGo.Native = xConfigureEventStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeXConfigureEvent)
	return structGo
}
func finalizeXConfigureEvent(obj *XConfigureEvent) {
	xConfigureEventStruct.Free(obj.Native)
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
	Native uintptr
}

// XImageStruct creates an uninitialised XImage.
func XImageStruct() *XImage {
	err := xImageStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &XImage{}
	structGo.Native = xImageStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeXImage)
	return structGo
}
func finalizeXImage(obj *XImage) {
	xImageStruct.Free(obj.Native)
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
	Native uintptr
}

// XFontStructStruct creates an uninitialised XFontStruct.
func XFontStructStruct() *XFontStruct {
	err := xFontStructStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &XFontStruct{}
	structGo.Native = xFontStructStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeXFontStruct)
	return structGo
}
func finalizeXFontStruct(obj *XFontStruct) {
	xFontStructStruct.Free(obj.Native)
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
	Native uintptr
}

// XTrapezoidStruct creates an uninitialised XTrapezoid.
func XTrapezoidStruct() *XTrapezoid {
	err := xTrapezoidStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &XTrapezoid{}
	structGo.Native = xTrapezoidStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeXTrapezoid)
	return structGo
}
func finalizeXTrapezoid(obj *XTrapezoid) {
	xTrapezoidStruct.Free(obj.Native)
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
	Native uintptr
}

// XVisualInfoStruct creates an uninitialised XVisualInfo.
func XVisualInfoStruct() *XVisualInfo {
	err := xVisualInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &XVisualInfo{}
	structGo.Native = xVisualInfoStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeXVisualInfo)
	return structGo
}
func finalizeXVisualInfo(obj *XVisualInfo) {
	xVisualInfoStruct.Free(obj.Native)
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
	Native uintptr
}

// XWindowAttributesStruct creates an uninitialised XWindowAttributes.
func XWindowAttributesStruct() *XWindowAttributes {
	err := xWindowAttributesStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &XWindowAttributes{}
	structGo.Native = xWindowAttributesStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeXWindowAttributes)
	return structGo
}
func finalizeXWindowAttributes(obj *XWindowAttributes) {
	xWindowAttributesStruct.Free(obj.Native)
}
