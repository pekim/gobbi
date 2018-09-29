// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_attr_iterator_get_font : unsupported parameter language : record with indirection level of 2

// Unsupported : pango_attr_iterator_range : unsupported parameter start : no type generator for gint, gint*

// Unsupported : pango_attr_list_filter : unsupported parameter func : no type generator for AttrFilterFunc, PangoAttrFilterFunc

// Unsupported : pango_coverage_to_bytes : unsupported parameter bytes : no param type

// Unsupported : pango_glyph_item_get_logical_widths : unsupported parameter logical_widths : no param type

// Unsupported : pango_glyph_item_letter_space : unsupported parameter log_attrs : no param type

// GlyphItemIter is a wrapper around the C record PangoGlyphItemIter.
type GlyphItemIter struct {
	native *C.PangoGlyphItemIter
	// glyph_item : record
	Text       string
	StartGlyph int32
	StartIndex int32
	StartChar  int32
	EndGlyph   int32
	EndIndex   int32
	EndChar    int32
}

func GlyphItemIterNewFromC(u unsafe.Pointer) *GlyphItemIter {
	c := (*C.PangoGlyphItemIter)(u)
	if c == nil {
		return nil
	}

	g := &GlyphItemIter{
		EndChar:    (int32)(c.end_char),
		EndGlyph:   (int32)(c.end_glyph),
		EndIndex:   (int32)(c.end_index),
		StartChar:  (int32)(c.start_char),
		StartGlyph: (int32)(c.start_glyph),
		StartIndex: (int32)(c.start_index),
		Text:       C.GoString(c.text),
		native:     c,
	}

	return g
}

func (recv *GlyphItemIter) ToC() unsafe.Pointer {
	recv.native.text =
		C.CString(recv.Text)
	recv.native.start_glyph =
		(C.int)(recv.StartGlyph)
	recv.native.start_index =
		(C.int)(recv.StartIndex)
	recv.native.start_char =
		(C.int)(recv.StartChar)
	recv.native.end_glyph =
		(C.int)(recv.EndGlyph)
	recv.native.end_index =
		(C.int)(recv.EndIndex)
	recv.native.end_char =
		(C.int)(recv.EndChar)

	return (unsafe.Pointer)(recv.native)
}

// Copy is a wrapper around the C function pango_glyph_item_iter_copy.
func (recv *GlyphItemIter) Copy() *GlyphItemIter {
	retC := C.pango_glyph_item_iter_copy((*C.PangoGlyphItemIter)(recv.native))
	retGo := GlyphItemIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function pango_glyph_item_iter_free.
func (recv *GlyphItemIter) Free() {
	C.pango_glyph_item_iter_free((*C.PangoGlyphItemIter)(recv.native))

	return
}

// InitEnd is a wrapper around the C function pango_glyph_item_iter_init_end.
func (recv *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) bool {
	c_glyph_item := (*C.PangoGlyphItem)(glyphItem.ToC())

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.pango_glyph_item_iter_init_end((*C.PangoGlyphItemIter)(recv.native), c_glyph_item, c_text)
	retGo := retC == C.TRUE

	return retGo
}

// InitStart is a wrapper around the C function pango_glyph_item_iter_init_start.
func (recv *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) bool {
	c_glyph_item := (*C.PangoGlyphItem)(glyphItem.ToC())

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.pango_glyph_item_iter_init_start((*C.PangoGlyphItemIter)(recv.native), c_glyph_item, c_text)
	retGo := retC == C.TRUE

	return retGo
}

// NextCluster is a wrapper around the C function pango_glyph_item_iter_next_cluster.
func (recv *GlyphItemIter) NextCluster() bool {
	retC := C.pango_glyph_item_iter_next_cluster((*C.PangoGlyphItemIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PrevCluster is a wrapper around the C function pango_glyph_item_iter_prev_cluster.
func (recv *GlyphItemIter) PrevCluster() bool {
	retC := C.pango_glyph_item_iter_prev_cluster((*C.PangoGlyphItemIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : pango_glyph_string_get_logical_widths : unsupported parameter logical_widths : no param type

// Unsupported : pango_glyph_string_index_to_x : unsupported parameter x_pos : no type generator for gint, int*

// Unsupported : pango_glyph_string_x_to_index : unsupported parameter index_ : no type generator for gint, int*

// Unsupported : pango_language_get_scripts : unsupported parameter num_scripts : no type generator for gint, int*

// Unsupported : pango_layout_line_get_x_ranges : unsupported parameter ranges : no param type

// Unsupported : pango_layout_line_index_to_x : unsupported parameter x_pos : no type generator for gint, int*

// Unsupported : pango_layout_line_x_to_index : unsupported parameter index_ : no type generator for gint, int*

// Unsupported : pango_matrix_get_font_scale_factors : unsupported parameter xscale : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_distance : unsupported parameter dx : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_point : unsupported parameter x : no type generator for gdouble, double*

// Unsupported : pango_script_iter_get_range : unsupported parameter script : PangoScript* with indirection level of 1

// Unsupported : pango_tab_array_new_with_positions : unsupported parameter ... : varargs

// Unsupported : pango_tab_array_get_tab : unsupported parameter alignment : PangoTabAlign* with indirection level of 1

// Unsupported : pango_tab_array_get_tabs : unsupported parameter alignments : PangoTabAlign** with indirection level of 2
