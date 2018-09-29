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

// GetBaseGravity is a wrapper around the C function pango_context_get_base_gravity.
func (recv *Context) GetBaseGravity() Gravity {
	retC := C.pango_context_get_base_gravity((*C.PangoContext)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// GetGravity is a wrapper around the C function pango_context_get_gravity.
func (recv *Context) GetGravity() Gravity {
	retC := C.pango_context_get_gravity((*C.PangoContext)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// GetGravityHint is a wrapper around the C function pango_context_get_gravity_hint.
func (recv *Context) GetGravityHint() GravityHint {
	retC := C.pango_context_get_gravity_hint((*C.PangoContext)(recv.native))
	retGo := (GravityHint)(retC)

	return retGo
}

// SetBaseGravity is a wrapper around the C function pango_context_set_base_gravity.
func (recv *Context) SetBaseGravity(gravity Gravity) {
	c_gravity := (C.PangoGravity)(gravity)

	C.pango_context_set_base_gravity((*C.PangoContext)(recv.native), c_gravity)

	return
}

// SetGravityHint is a wrapper around the C function pango_context_set_gravity_hint.
func (recv *Context) SetGravityHint(hint GravityHint) {
	c_hint := (C.PangoGravityHint)(hint)

	C.pango_context_set_gravity_hint((*C.PangoContext)(recv.native), c_hint)

	return
}

// GetLineReadonly is a wrapper around the C function pango_layout_get_line_readonly.
func (recv *Layout) GetLineReadonly(line int32) *LayoutLine {
	c_line := (C.int)(line)

	retC := C.pango_layout_get_line_readonly((*C.PangoLayout)(recv.native), c_line)
	retGo := LayoutLineNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLinesReadonly is a wrapper around the C function pango_layout_get_lines_readonly.
func (recv *Layout) GetLinesReadonly() *glib.SList {
	retC := C.pango_layout_get_lines_readonly((*C.PangoLayout)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUnknownGlyphsCount is a wrapper around the C function pango_layout_get_unknown_glyphs_count.
func (recv *Layout) GetUnknownGlyphsCount() int32 {
	retC := C.pango_layout_get_unknown_glyphs_count((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// IsEllipsized is a wrapper around the C function pango_layout_is_ellipsized.
func (recv *Layout) IsEllipsized() bool {
	retC := C.pango_layout_is_ellipsized((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsWrapped is a wrapper around the C function pango_layout_is_wrapped.
func (recv *Layout) IsWrapped() bool {
	retC := C.pango_layout_is_wrapped((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
