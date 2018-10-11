// This is a generated file - DO NOT EDIT
// +build pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// DescribeWithAbsoluteSize is a wrapper around the C function pango_font_describe_with_absolute_size.
func (recv *Font) DescribeWithAbsoluteSize() *FontDescription {
	retC := C.pango_font_describe_with_absolute_size((*C.PangoFont)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}
