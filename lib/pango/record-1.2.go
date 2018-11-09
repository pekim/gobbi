// This is a generated file - DO NOT EDIT
// +build pango_1.2 pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetAttrs is a wrapper around the C function pango_attr_iterator_get_attrs.
func (recv *AttrIterator) GetAttrs() *glib.SList {
	retC := C.pango_attr_iterator_get_attrs((*C.PangoAttrIterator)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_attr_list_filter : unsupported parameter func : no type generator for AttrFilterFunc (PangoAttrFilterFunc) for param func

// ApplyAttrs is a wrapper around the C function pango_glyph_item_apply_attrs.
func (recv *GlyphItem) ApplyAttrs(text string, list *AttrList) *glib.SList {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_list := (*C.PangoAttrList)(C.NULL)
	if list != nil {
		c_list = (*C.PangoAttrList)(list.ToC())
	}

	retC := C.pango_glyph_item_apply_attrs((*C.PangoGlyphItem)(recv.native), c_text, c_list)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Split is a wrapper around the C function pango_glyph_item_split.
func (recv *GlyphItem) Split(text string, splitIndex int32) *GlyphItem {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_split_index := (C.int)(splitIndex)

	retC := C.pango_glyph_item_split((*C.PangoGlyphItem)(recv.native), c_text, c_split_index)
	retGo := GlyphItemNewFromC(unsafe.Pointer(retC))

	return retGo
}
