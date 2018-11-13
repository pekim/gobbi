// This is a generated file - DO NOT EDIT
// +build pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Gets the height of layout used for ellipsization.  See
// pango_layout_set_height() for details.
/*

C function : pango_layout_get_height
*/
func (recv *Layout) GetHeight() int32 {
	retC := C.pango_layout_get_height((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the height to which the #PangoLayout should be ellipsized at.  There
// are two different behaviors, based on whether @height is positive or
// negative.
//
// If @height is positive, it will be the maximum height of the layout.  Only
// lines would be shown that would fit, and if there is any text omitted,
// an ellipsis added.  At least one line is included in each paragraph regardless
// of how small the height value is.  A value of zero will render exactly one
// line for the entire layout.
//
// If @height is negative, it will be the (negative of) maximum number of lines per
// paragraph.  That is, the total number of lines shown may well be more than
// this value if the layout contains multiple paragraphs of text.
// The default value of -1 means that first line of each paragraph is ellipsized.
// This behvaior may be changed in the future to act per layout instead of per
// paragraph.  File a bug against pango at <ulink
// url="http://bugzilla.gnome.org/">http://bugzilla.gnome.org/</ulink> if your
// code relies on this behavior.
//
// Height setting only has effect if a positive width is set on
// @layout and ellipsization mode of @layout is not %PANGO_ELLIPSIZE_NONE.
// The behavior is undefined if a height other than -1 is set and
// ellipsization mode is set to %PANGO_ELLIPSIZE_NONE, and may change in the
// future.
/*

C function : pango_layout_set_height
*/
func (recv *Layout) SetHeight(height int32) {
	c_height := (C.int)(height)

	C.pango_layout_set_height((*C.PangoLayout)(recv.native), c_height)

	return
}

// Gets the layout currently being rendered using @renderer.
// Calling this function only makes sense from inside a subclass's
// methods, like in its draw_shape<!---->() for example.
//
// The returned layout should not be modified while still being
// rendered.
/*

C function : pango_renderer_get_layout
*/
func (recv *Renderer) GetLayout() *Layout {
	retC := C.pango_renderer_get_layout((*C.PangoRenderer)(recv.native))
	var retGo (*Layout)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LayoutNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the layout line currently being rendered using @renderer.
// Calling this function only makes sense from inside a subclass's
// methods, like in its draw_shape<!---->() for example.
//
// The returned layout line should not be modified while still being
// rendered.
/*

C function : pango_renderer_get_layout_line
*/
func (recv *Renderer) GetLayoutLine() *LayoutLine {
	retC := C.pango_renderer_get_layout_line((*C.PangoRenderer)(recv.native))
	var retGo (*LayoutLine)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LayoutLineNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}
