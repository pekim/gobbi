// This is a generated file - DO NOT EDIT
// +build pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_attr_iterator_get_font : unsupported parameter language : record with indirection level of 2

// Unsupported : pango_attr_iterator_range : unsupported parameter start : no type generator for gint, gint*

// Unsupported : pango_attr_list_filter : unsupported parameter func : no type generator for AttrFilterFunc, PangoAttrFilterFunc

// Init is a wrapper around the C function pango_attribute_init.
func (recv *Attribute) Init(klass *AttrClass) {
	c_klass := (*C.PangoAttrClass)(klass.ToC())

	C.pango_attribute_init((*C.PangoAttribute)(recv.native), c_klass)

	return
}

// Unsupported : pango_coverage_to_bytes : unsupported parameter bytes : no param type

// Blacklisted : PangoEngineClass

// Blacklisted : PangoEngineInfo

// Blacklisted : PangoEngineLangClass

// Blacklisted : PangoEngineScriptInfo

// Blacklisted : PangoEngineShapeClass

// Blacklisted : PangoFontClass

// Blacklisted : PangoFontFaceClass

// Blacklisted : PangoFontFamilyClass

// Blacklisted : PangoFontMapClass

// Blacklisted : PangoFontMetrics

// Blacklisted : PangoFontsetClass

// Blacklisted : PangoFontsetSimpleClass

// Copy is a wrapper around the C function pango_glyph_item_copy.
func (recv *GlyphItem) Copy() *GlyphItem {
	retC := C.pango_glyph_item_copy((*C.PangoGlyphItem)(recv.native))
	retGo := GlyphItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_glyph_item_get_logical_widths : unsupported parameter logical_widths : no param type

// Unsupported : pango_glyph_item_letter_space : unsupported parameter log_attrs : no param type

// Unsupported : pango_glyph_string_get_logical_widths : unsupported parameter logical_widths : no param type

// Unsupported : pango_glyph_string_index_to_x : unsupported parameter x_pos : no type generator for gint, int*

// Unsupported : pango_glyph_string_x_to_index : unsupported parameter index_ : no type generator for gint, int*

// Blacklisted : PangoIncludedModule

// Unsupported : pango_language_get_scripts : unsupported parameter num_scripts : no type generator for gint, int*

// Blacklisted : PangoLayoutIter

// Unsupported : pango_layout_line_get_x_ranges : unsupported parameter ranges : no param type

// Unsupported : pango_layout_line_index_to_x : unsupported parameter x_pos : no type generator for gint, int*

// Unsupported : pango_layout_line_x_to_index : unsupported parameter index_ : no type generator for gint, int*

// Blacklisted : PangoMap

// Blacklisted : PangoMapEntry

// Unsupported : pango_matrix_get_font_scale_factors : unsupported parameter xscale : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_distance : unsupported parameter dx : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_point : unsupported parameter x : no type generator for gdouble, double*

// Blacklisted : PangoScriptForLang

// Unsupported : pango_script_iter_get_range : unsupported parameter script : PangoScript* with indirection level of 1

// Unsupported : pango_tab_array_new_with_positions : unsupported parameter ... : varargs

// Unsupported : pango_tab_array_get_tab : unsupported parameter alignment : PangoTabAlign* with indirection level of 1

// Unsupported : pango_tab_array_get_tabs : unsupported parameter alignments : PangoTabAlign** with indirection level of 2
