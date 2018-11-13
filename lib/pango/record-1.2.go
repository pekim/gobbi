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

// Gets a list of all attributes at the current position of the
// iterator.
/*

C function : pango_attr_iterator_get_attrs
*/
func (recv *AttrIterator) GetAttrs() *glib.SList {
	retC := C.pango_attr_iterator_get_attrs((*C.PangoAttrIterator)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_attr_list_filter : unsupported parameter func : no type generator for AttrFilterFunc (PangoAttrFilterFunc) for param func

// Splits a shaped item (PangoGlyphItem) into multiple items based
// on an attribute list. The idea is that if you have attributes
// that don't affect shaping, such as color or underline, to avoid
// affecting shaping, you filter them out (pango_attr_list_filter()),
// apply the shaping process and then reapply them to the result using
// this function.
//
// All attributes that start or end inside a cluster are applied
// to that cluster; for instance, if half of a cluster is underlined
// and the other-half strikethrough, then the cluster will end
// up with both underline and strikethrough attributes. In these
// cases, it may happen that item->extra_attrs for some of the
// result items can have multiple attributes of the same type.
//
// This function takes ownership of @glyph_item; it will be reused
// as one of the elements in the list.
/*

C function : pango_glyph_item_apply_attrs
*/
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

// Modifies @orig to cover only the text after @split_index, and
// returns a new item that covers the text before @split_index that
// used to be in @orig. You can think of @split_index as the length of
// the returned item. @split_index may not be 0, and it may not be
// greater than or equal to the length of @orig (that is, there must
// be at least one byte assigned to each item, you can't create a
// zero-length item).
//
// This function is similar in function to pango_item_split() (and uses
// it internally.)
/*

C function : pango_glyph_item_split
*/
func (recv *GlyphItem) Split(text string, splitIndex int32) *GlyphItem {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_split_index := (C.int)(splitIndex)

	retC := C.pango_glyph_item_split((*C.PangoGlyphItem)(recv.native), c_text, c_split_index)
	retGo := GlyphItemNewFromC(unsafe.Pointer(retC))

	return retGo
}
