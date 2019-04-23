// This is a generated file - DO NOT EDIT

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Blacklisted : AtkActionIface

// Blacklisted : AtkAttribute

// Blacklisted : AtkComponentIface

// Blacklisted : AtkDocumentIface

// Blacklisted : AtkEditableTextIface

// Blacklisted : AtkGObjectAccessibleClass

// Blacklisted : AtkHyperlinkClass

// Blacklisted : AtkHyperlinkImplIface

// Blacklisted : AtkHypertextIface

// Blacklisted : AtkImageIface

// Blacklisted : AtkImplementor

// Blacklisted : AtkKeyEventStruct

// Blacklisted : AtkMiscClass

// Blacklisted : AtkNoOpObjectClass

// Blacklisted : AtkNoOpObjectFactoryClass

// Blacklisted : AtkObjectClass

// Blacklisted : AtkObjectFactoryClass

// Blacklisted : AtkPlugClass

// Blacklisted : AtkPropertyValues

// Blacklisted : AtkRange

// Rectangle is a wrapper around the C record AtkRectangle.
type Rectangle struct {
	native *C.AtkRectangle
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func RectangleNewFromC(u unsafe.Pointer) *Rectangle {
	c := (*C.AtkRectangle)(u)
	if c == nil {
		return nil
	}

	g := &Rectangle{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *Rectangle) ToC() unsafe.Pointer {
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)
	recv.native.width =
		(C.gint)(recv.Width)
	recv.native.height =
		(C.gint)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Rectangle with another Rectangle, and returns true if they represent the same GObject.
func (recv *Rectangle) Equals(other *Rectangle) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : AtkRegistryClass

// Blacklisted : AtkRelationClass

// Blacklisted : AtkRelationSetClass

// Blacklisted : AtkSelectionIface

// Blacklisted : AtkSocketClass

// Blacklisted : AtkStateSetClass

// Blacklisted : AtkStreamableContentIface

// Blacklisted : AtkTableCellIface

// Blacklisted : AtkTableIface

// Blacklisted : AtkTextIface

// Blacklisted : AtkTextRange

// Blacklisted : AtkTextRectangle

// Blacklisted : AtkUtilClass

// Blacklisted : AtkValueIface

// Blacklisted : AtkWindowIface
