// Code generated - DO NOT EDIT.

package xlib

import "unsafe"

// Atom is a representation of the C alias Atom.
type Atom uint64

// Colormap is a representation of the C alias Colormap.
type Colormap uint64

// Cursor is a representation of the C alias Cursor.
type Cursor uint64

// Drawable is a representation of the C alias Drawable.
type Drawable uint64

// GC is a representation of the C alias GC.
type GC unsafe.Pointer

// KeyCode is a representation of the C alias KeyCode.
type KeyCode uint8

// KeySym is a representation of the C alias KeySym.
type KeySym uint64

// Picture is a representation of the C alias Picture.
type Picture uint64

// Time is a representation of the C alias Time.
type Time uint64

// VisualID is a representation of the C alias VisualID.
type VisualID uint64

// Window is a representation of the C alias Window.
type Window uint64

// XID is a representation of the C alias XID.
type XID uint64

// Pixmap is a representation of the C alias Pixmap.
type Pixmap uint64

// UNSUPPORTED : XOpenDisplay : blacklisted

// Display is a representation of the C record Display.
type Display struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C Display that represents the Display.
func (recv *Display) ToC() unsafe.Pointer {
	return recv.native
}

// DisplayNewFromC creates a new Display from a pointer to the C Display that represents the Display.
func DisplayNewFromC(native unsafe.Pointer) *Display {
	return &Display{native: native}
}

// Screen is a representation of the C record Screen.
type Screen struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C Screen that represents the Screen.
func (recv *Screen) ToC() unsafe.Pointer {
	return recv.native
}

// ScreenNewFromC creates a new Screen from a pointer to the C Screen that represents the Screen.
func ScreenNewFromC(native unsafe.Pointer) *Screen {
	return &Screen{native: native}
}

// Visual is a representation of the C record Visual.
type Visual struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C Visual that represents the Visual.
func (recv *Visual) ToC() unsafe.Pointer {
	return recv.native
}

// VisualNewFromC creates a new Visual from a pointer to the C Visual that represents the Visual.
func VisualNewFromC(native unsafe.Pointer) *Visual {
	return &Visual{native: native}
}

// XConfigureEvent is a representation of the C record XConfigureEvent.
type XConfigureEvent struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C XConfigureEvent that represents the XConfigureEvent.
func (recv *XConfigureEvent) ToC() unsafe.Pointer {
	return recv.native
}

// XConfigureEventNewFromC creates a new XConfigureEvent from a pointer to the C XConfigureEvent that represents the XConfigureEvent.
func XConfigureEventNewFromC(native unsafe.Pointer) *XConfigureEvent {
	return &XConfigureEvent{native: native}
}

// XImage is a representation of the C record XImage.
type XImage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C XImage that represents the XImage.
func (recv *XImage) ToC() unsafe.Pointer {
	return recv.native
}

// XImageNewFromC creates a new XImage from a pointer to the C XImage that represents the XImage.
func XImageNewFromC(native unsafe.Pointer) *XImage {
	return &XImage{native: native}
}

// XFontStruct is a representation of the C record XFontStruct.
type XFontStruct struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C XFontStruct that represents the XFontStruct.
func (recv *XFontStruct) ToC() unsafe.Pointer {
	return recv.native
}

// XFontStructNewFromC creates a new XFontStruct from a pointer to the C XFontStruct that represents the XFontStruct.
func XFontStructNewFromC(native unsafe.Pointer) *XFontStruct {
	return &XFontStruct{native: native}
}

// XTrapezoid is a representation of the C record XTrapezoid.
type XTrapezoid struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C XTrapezoid that represents the XTrapezoid.
func (recv *XTrapezoid) ToC() unsafe.Pointer {
	return recv.native
}

// XTrapezoidNewFromC creates a new XTrapezoid from a pointer to the C XTrapezoid that represents the XTrapezoid.
func XTrapezoidNewFromC(native unsafe.Pointer) *XTrapezoid {
	return &XTrapezoid{native: native}
}

// XVisualInfo is a representation of the C record XVisualInfo.
type XVisualInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C XVisualInfo that represents the XVisualInfo.
func (recv *XVisualInfo) ToC() unsafe.Pointer {
	return recv.native
}

// XVisualInfoNewFromC creates a new XVisualInfo from a pointer to the C XVisualInfo that represents the XVisualInfo.
func XVisualInfoNewFromC(native unsafe.Pointer) *XVisualInfo {
	return &XVisualInfo{native: native}
}

// XWindowAttributes is a representation of the C record XWindowAttributes.
type XWindowAttributes struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C XWindowAttributes that represents the XWindowAttributes.
func (recv *XWindowAttributes) ToC() unsafe.Pointer {
	return recv.native
}

// XWindowAttributesNewFromC creates a new XWindowAttributes from a pointer to the C XWindowAttributes that represents the XWindowAttributes.
func XWindowAttributesNewFromC(native unsafe.Pointer) *XWindowAttributes {
	return &XWindowAttributes{native: native}
}

// XEvent is a representation of the C union XEvent.
type XEvent struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C XEvent that represents the XEvent.
func (recv *XEvent) ToC() unsafe.Pointer {
	return recv.native
}

// XEventNewFromC creates a new XEvent from a pointer to the C XEvent that represents the XEvent.
func XEventNewFromC(native unsafe.Pointer) *XEvent {
	return &XEvent{native: native}
}