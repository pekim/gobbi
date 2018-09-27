// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// CssSection is a wrapper around the C record GtkCssSection.
type CssSection struct {
	native *C.GtkCssSection
}

func CssSectionNewFromC(u unsafe.Pointer) *CssSection {
	c := (*C.GtkCssSection)(u)
	if c == nil {
		return nil
	}

	g := &CssSection{native: c}

	return g
}

func (recv *CssSection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetEndLine is a wrapper around the C function gtk_css_section_get_end_line.
func (recv *CssSection) GetEndLine() uint32 {
	retC := C.gtk_css_section_get_end_line((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetEndPosition is a wrapper around the C function gtk_css_section_get_end_position.
func (recv *CssSection) GetEndPosition() uint32 {
	retC := C.gtk_css_section_get_end_position((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_css_section_get_file : no return generator

// GetParent is a wrapper around the C function gtk_css_section_get_parent.
func (recv *CssSection) GetParent() *CssSection {
	retC := C.gtk_css_section_get_parent((*C.GtkCssSection)(recv.native))
	retGo := CssSectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSectionType is a wrapper around the C function gtk_css_section_get_section_type.
func (recv *CssSection) GetSectionType() CssSectionType {
	retC := C.gtk_css_section_get_section_type((*C.GtkCssSection)(recv.native))
	retGo := (CssSectionType)(retC)

	return retGo
}

// GetStartLine is a wrapper around the C function gtk_css_section_get_start_line.
func (recv *CssSection) GetStartLine() uint32 {
	retC := C.gtk_css_section_get_start_line((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetStartPosition is a wrapper around the C function gtk_css_section_get_start_position.
func (recv *CssSection) GetStartPosition() uint32 {
	retC := C.gtk_css_section_get_start_position((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Ref is a wrapper around the C function gtk_css_section_ref.
func (recv *CssSection) Ref() *CssSection {
	retC := C.gtk_css_section_ref((*C.GtkCssSection)(recv.native))
	retGo := CssSectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_css_section_unref : no return generator
