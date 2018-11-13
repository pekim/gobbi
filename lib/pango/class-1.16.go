// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Retrieves the base gravity for the context. See
// pango_context_set_base_gravity().
/*

C function

pango_context_get_base_gravity
*/
func (recv *Context) GetBaseGravity() Gravity {
	retC := C.pango_context_get_base_gravity((*C.PangoContext)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// Retrieves the gravity for the context. This is similar to
// pango_context_get_base_gravity(), except for when the base gravity
// is %PANGO_GRAVITY_AUTO for which pango_gravity_get_for_matrix() is used
// to return the gravity from the current context matrix.
/*

C function

pango_context_get_gravity
*/
func (recv *Context) GetGravity() Gravity {
	retC := C.pango_context_get_gravity((*C.PangoContext)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// Retrieves the gravity hint for the context. See
// pango_context_set_gravity_hint() for details.
/*

C function

pango_context_get_gravity_hint
*/
func (recv *Context) GetGravityHint() GravityHint {
	retC := C.pango_context_get_gravity_hint((*C.PangoContext)(recv.native))
	retGo := (GravityHint)(retC)

	return retGo
}

// Sets the base gravity for the context.
//
// The base gravity is used in laying vertical text out.
/*

C function

pango_context_set_base_gravity
*/
func (recv *Context) SetBaseGravity(gravity Gravity) {
	c_gravity := (C.PangoGravity)(gravity)

	C.pango_context_set_base_gravity((*C.PangoContext)(recv.native), c_gravity)

	return
}

// Sets the gravity hint for the context.
//
// The gravity hint is used in laying vertical text out, and is only relevant
// if gravity of the context as returned by pango_context_get_gravity()
// is set %PANGO_GRAVITY_EAST or %PANGO_GRAVITY_WEST.
/*

C function

pango_context_set_gravity_hint
*/
func (recv *Context) SetGravityHint(hint GravityHint) {
	c_hint := (C.PangoGravityHint)(hint)

	C.pango_context_set_gravity_hint((*C.PangoContext)(recv.native), c_hint)

	return
}

// Retrieves a particular line from a #PangoLayout.
//
// This is a faster alternative to pango_layout_get_line(),
// but the user is not expected
// to modify the contents of the line (glyphs, glyph widths, etc.).
/*

C function

pango_layout_get_line_readonly
*/
func (recv *Layout) GetLineReadonly(line int32) *LayoutLine {
	c_line := (C.int)(line)

	retC := C.pango_layout_get_line_readonly((*C.PangoLayout)(recv.native), c_line)
	var retGo (*LayoutLine)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LayoutLineNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the lines of the @layout as a list.
//
// This is a faster alternative to pango_layout_get_lines(),
// but the user is not expected
// to modify the contents of the lines (glyphs, glyph widths, etc.).
/*

C function

pango_layout_get_lines_readonly
*/
func (recv *Layout) GetLinesReadonly() *glib.SList {
	retC := C.pango_layout_get_lines_readonly((*C.PangoLayout)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Counts the number unknown glyphs in @layout.  That is, zero if
// glyphs for all characters in the layout text were found, or more
// than zero otherwise.
//
// This function can be used to determine if there are any fonts
// available to render all characters in a certain string, or when
// used in combination with %PANGO_ATTR_FALLBACK, to check if a
// certain font supports all the characters in the string.
/*

C function

pango_layout_get_unknown_glyphs_count
*/
func (recv *Layout) GetUnknownGlyphsCount() int32 {
	retC := C.pango_layout_get_unknown_glyphs_count((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Queries whether the layout had to ellipsize any paragraphs.
//
// This returns %TRUE if the ellipsization mode for @layout
// is not %PANGO_ELLIPSIZE_NONE, a positive width is set on @layout,
// and there are paragraphs exceeding that width that have to be
// ellipsized.
/*

C function

pango_layout_is_ellipsized
*/
func (recv *Layout) IsEllipsized() bool {
	retC := C.pango_layout_is_ellipsized((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Queries whether the layout had to wrap any paragraphs.
//
// This returns %TRUE if a positive width is set on @layout,
// ellipsization mode of @layout is set to %PANGO_ELLIPSIZE_NONE,
// and there are paragraphs exceeding the layout width that have
// to be wrapped.
/*

C function

pango_layout_is_wrapped
*/
func (recv *Layout) IsWrapped() bool {
	retC := C.pango_layout_is_wrapped((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
