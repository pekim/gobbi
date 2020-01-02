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

func (recv *Display) ToC() unsafe.Pointer {
	return recv.native
}

// Screen is a representation of the C record Screen.
type Screen struct {
	native unsafe.Pointer
}

func (recv *Screen) ToC() unsafe.Pointer {
	return recv.native
}

// Visual is a representation of the C record Visual.
type Visual struct {
	native unsafe.Pointer
}

func (recv *Visual) ToC() unsafe.Pointer {
	return recv.native
}

// XConfigureEvent is a representation of the C record XConfigureEvent.
type XConfigureEvent struct {
	native unsafe.Pointer
}

func (recv *XConfigureEvent) ToC() unsafe.Pointer {
	return recv.native
}

// XImage is a representation of the C record XImage.
type XImage struct {
	native unsafe.Pointer
}

func (recv *XImage) ToC() unsafe.Pointer {
	return recv.native
}

// XFontStruct is a representation of the C record XFontStruct.
type XFontStruct struct {
	native unsafe.Pointer
}

func (recv *XFontStruct) ToC() unsafe.Pointer {
	return recv.native
}

// XTrapezoid is a representation of the C record XTrapezoid.
type XTrapezoid struct {
	native unsafe.Pointer
}

func (recv *XTrapezoid) ToC() unsafe.Pointer {
	return recv.native
}

// XVisualInfo is a representation of the C record XVisualInfo.
type XVisualInfo struct {
	native unsafe.Pointer
}

func (recv *XVisualInfo) ToC() unsafe.Pointer {
	return recv.native
}

// XWindowAttributes is a representation of the C record XWindowAttributes.
type XWindowAttributes struct {
	native unsafe.Pointer
}

func (recv *XWindowAttributes) ToC() unsafe.Pointer {
	return recv.native
}

// XEvent is a representation of the C union XEvent.
type XEvent struct{}
