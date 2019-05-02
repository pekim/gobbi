// Code generated - DO NOT EDIT.
// +build pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Changed is a wrapper around the C function pango_context_changed.
func (recv *Context) Changed() {
	C.pango_context_changed((*C.PangoContext)(recv.native))

	return
}

// GetSerial is a wrapper around the C function pango_context_get_serial.
func (recv *Context) GetSerial() uint32 {
	retC := C.pango_context_get_serial((*C.PangoContext)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSerial is a wrapper around the C function pango_font_map_get_serial.
func (recv *FontMap) GetSerial() uint32 {
	retC := C.pango_font_map_get_serial((*C.PangoFontMap)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSerial is a wrapper around the C function pango_layout_get_serial.
func (recv *Layout) GetSerial() uint32 {
	retC := C.pango_layout_get_serial((*C.PangoLayout)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}
