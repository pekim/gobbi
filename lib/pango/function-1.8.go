// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// AttrStrikethroughColorNew is a wrapper around the C function pango_attr_strikethrough_color_new.
func AttrStrikethroughColorNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_strikethrough_color_new(c_red, c_green, c_blue)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrUnderlineColorNew is a wrapper around the C function pango_attr_underline_color_new.
func AttrUnderlineColorNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_underline_color_new(c_red, c_green, c_blue)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// Unsupported : pango_skip_space : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_split_file_list : no return type
