// This is a generated file - DO NOT EDIT
// +build pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Forces a change in the context, which will cause any #PangoLayout
// using this context to re-layout.
//
// This function is only useful when implementing a new backend
// for Pango, something applications won't do. Backends should
// call this function if they have attached extra data to the context
// and such data is changed.
/*

C function : pango_context_changed
*/
func (recv *Context) Changed() {
	C.pango_context_changed((*C.PangoContext)(recv.native))

	return
}

// Returns the current serial number of @context.  The serial number is
// initialized to an small number larger than zero when a new context
// is created and is increased whenever the context is changed using any
// of the setter functions, or the #PangoFontMap it uses to find fonts has
// changed. The serial may wrap, but will never have the value 0. Since it
// can wrap, never compare it with "less than", always use "not equals".
//
// This can be used to automatically detect changes to a #PangoContext, and
// is only useful when implementing objects that need update when their
// #PangoContext changes, like #PangoLayout.
/*

C function : pango_context_get_serial
*/
func (recv *Context) GetSerial() uint32 {
	retC := C.pango_context_get_serial((*C.PangoContext)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the current serial number of @fontmap.  The serial number is
// initialized to an small number larger than zero when a new fontmap
// is created and is increased whenever the fontmap is changed. It may
// wrap, but will never have the value 0. Since it can wrap, never compare
// it with "less than", always use "not equals".
//
// The fontmap can only be changed using backend-specific API, like changing
// fontmap resolution.
//
// This can be used to automatically detect changes to a #PangoFontMap, like
// in #PangoContext.
/*

C function : pango_font_map_get_serial
*/
func (recv *FontMap) GetSerial() uint32 {
	retC := C.pango_font_map_get_serial((*C.PangoFontMap)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the current serial number of @layout.  The serial number is
// initialized to an small number  larger than zero when a new layout
// is created and is increased whenever the layout is changed using any
// of the setter functions, or the #PangoContext it uses has changed.
// The serial may wrap, but will never have the value 0. Since it
// can wrap, never compare it with "less than", always use "not equals".
//
// This can be used to automatically detect changes to a #PangoLayout, and
// is useful for example to decide whether a layout needs redrawing.
// To force the serial to be increased, use pango_layout_context_changed().
/*

C function : pango_layout_get_serial
*/
func (recv *Layout) GetSerial() uint32 {
	retC := C.pango_layout_get_serial((*C.PangoLayout)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}
