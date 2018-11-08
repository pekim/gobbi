// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// AttrTypeGetName is a wrapper around the C function pango_attr_type_get_name.
func AttrTypeGetName(type_ AttrType) string {
	c_type := (C.PangoAttrType)(type_)

	retC := C.pango_attr_type_get_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// BidiTypeForUnichar is a wrapper around the C function pango_bidi_type_for_unichar.
func BidiTypeForUnichar(ch rune) BidiType {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_bidi_type_for_unichar(c_ch)
	retGo := (BidiType)(retC)

	return retGo
}

// Unsupported : pango_break : unsupported parameter attrs : no type generator for LogAttr (PangoLogAttr) for array param attrs

// Unsupported : pango_find_map : return type : Blacklisted record : PangoMap

// Unsupported : pango_get_log_attrs : unsupported parameter log_attrs : no type generator for LogAttr (PangoLogAttr) for array param log_attrs

// Unsupported : pango_log2vis_get_embedding_levels : unsupported parameter pbase_dir : PangoDirection* with indirection level of 1

// Unsupported : pango_lookup_aliases : unsupported parameter families : no type generator for utf8 (char**) for array param families

// Unsupported : pango_markup_parser_finish : unsupported parameter attr_list : record with indirection level of 2

// Unsupported : pango_module_register : unsupported parameter module : Blacklisted record : PangoIncludedModule

// Unsupported : pango_parse_markup : unsupported parameter attr_list : record with indirection level of 2

// Unsupported : pango_parse_stretch : unsupported parameter stretch : PangoStretch* with indirection level of 1

// Unsupported : pango_parse_style : unsupported parameter style : PangoStyle* with indirection level of 1

// Unsupported : pango_parse_variant : unsupported parameter variant : PangoVariant* with indirection level of 1

// Unsupported : pango_parse_weight : unsupported parameter weight : PangoWeight* with indirection level of 1

// Unsupported : pango_read_line : unsupported parameter stream : no type generator for gpointer (FILE*) for param stream

// Unsupported : pango_scan_int : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_string : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_word : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_shape : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Unsupported : pango_shape_full : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Unsupported : pango_skip_space : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_split_file_list : no return type
