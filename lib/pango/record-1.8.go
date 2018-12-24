// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// pango_attr_shape_new_with_data : unsupported parameter copy_func : no type generator for AttrDataCopyFunc (PangoAttrDataCopyFunc) for param copy_func
// AttrSizeNewAbsolute is a wrapper around the C function pango_attr_size_new_absolute.
func AttrSizeNewAbsolute(size int32) *Attribute {
	c_size := (C.int)(size)

	retC := C.pango_attr_size_new_absolute(c_size)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSizeIsAbsolute is a wrapper around the C function pango_font_description_get_size_is_absolute.
func (recv *FontDescription) GetSizeIsAbsolute() bool {
	retC := C.pango_font_description_get_size_is_absolute((*C.PangoFontDescription)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAbsoluteSize is a wrapper around the C function pango_font_description_set_absolute_size.
func (recv *FontDescription) SetAbsoluteSize(size float64) {
	c_size := (C.double)(size)

	C.pango_font_description_set_absolute_size((*C.PangoFontDescription)(recv.native), c_size)

	return
}

// RendererClass is a wrapper around the C record PangoRendererClass.
type RendererClass struct {
	native *C.PangoRendererClass
	// Private : parent_class
	// no type for draw_glyphs
	// no type for draw_rectangle
	// no type for draw_error_underline
	// no type for draw_shape
	// no type for draw_trapezoid
	// no type for draw_glyph
	// no type for part_changed
	// no type for begin
	// no type for end
	// no type for prepare_run
	// no type for draw_glyph_item
	// no type for _pango_reserved2
	// no type for _pango_reserved3
	// no type for _pango_reserved4
}

func RendererClassNewFromC(u unsafe.Pointer) *RendererClass {
	c := (*C.PangoRendererClass)(u)
	if c == nil {
		return nil
	}

	g := &RendererClass{native: c}

	return g
}

func (recv *RendererClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RendererClass with another RendererClass, and returns true if they represent the same GObject.
func (recv *RendererClass) Equals(other *RendererClass) bool {
	return other.ToC() == recv.ToC()
}
