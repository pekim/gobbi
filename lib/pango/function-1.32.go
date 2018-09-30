// This is a generated file - DO NOT EDIT
// +build pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_break : unsupported parameter attrs : no param type

// Unsupported : pango_find_map : return type : Blacklisted record : PangoMap

// Unsupported : pango_get_log_attrs : unsupported parameter log_attrs : no param type

// Unsupported : pango_log2vis_get_embedding_levels : unsupported parameter pbase_dir : PangoDirection* with indirection level of 1

// Unsupported : pango_lookup_aliases : unsupported parameter families : no param type

// Unsupported : pango_markup_parser_finish : unsupported parameter attr_list : record with indirection level of 2

// Unsupported : pango_module_register : unsupported parameter module : Blacklisted record : PangoIncludedModule

// Unsupported : pango_parse_enum : unsupported parameter type : no type generator for GType, GType

// Unsupported : pango_parse_markup : unsupported parameter attr_list : record with indirection level of 2

// Unsupported : pango_parse_stretch : unsupported parameter stretch : PangoStretch* with indirection level of 1

// Unsupported : pango_parse_style : unsupported parameter style : PangoStyle* with indirection level of 1

// Unsupported : pango_parse_variant : unsupported parameter variant : PangoVariant* with indirection level of 1

// Unsupported : pango_parse_weight : unsupported parameter weight : PangoWeight* with indirection level of 1

// Unsupported : pango_read_line : unsupported parameter stream : no type generator for gpointer, FILE*

// Unsupported : pango_scan_int : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_string : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_word : unsupported parameter pos : in string with indirection level of 2

// ShapeFull is a wrapper around the C function pango_shape_full.
func ShapeFull(itemText string, itemLength int32, paragraphText string, paragraphLength int32, analysis *Analysis, glyphs *GlyphString) {
	c_item_text := C.CString(itemText)
	defer C.free(unsafe.Pointer(c_item_text))

	c_item_length := (C.gint)(itemLength)

	c_paragraph_text := C.CString(paragraphText)
	defer C.free(unsafe.Pointer(c_paragraph_text))

	c_paragraph_length := (C.gint)(paragraphLength)

	c_analysis := (*C.PangoAnalysis)(analysis.ToC())

	c_glyphs := (*C.PangoGlyphString)(glyphs.ToC())

	C.pango_shape_full(c_item_text, c_item_length, c_paragraph_text, c_paragraph_length, c_analysis, c_glyphs)

	return
}

// Unsupported : pango_skip_space : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_split_file_list : no return type
