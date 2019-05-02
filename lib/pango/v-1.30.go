// This is a generated file - DO NOT EDIT
// +build pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetCharacterCount is a wrapper around the C function pango_layout_get_character_count.
func (recv *Layout) GetCharacterCount() int32 {
	retC := C.pango_layout_get_character_count((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : pango_layout_get_log_attrs_readonly : array return type :
