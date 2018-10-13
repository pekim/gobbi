// This is a generated file - DO NOT EDIT

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// AttrBackgroundNew is a wrapper around the C function pango_attr_background_new.
func AttrBackgroundNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_background_new(c_red, c_green, c_blue)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrFamilyNew is a wrapper around the C function pango_attr_family_new.
func AttrFamilyNew(family string) *Attribute {
	c_family := C.CString(family)
	defer C.free(unsafe.Pointer(c_family))

	retC := C.pango_attr_family_new(c_family)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrForegroundNew is a wrapper around the C function pango_attr_foreground_new.
func AttrForegroundNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_foreground_new(c_red, c_green, c_blue)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrRiseNew is a wrapper around the C function pango_attr_rise_new.
func AttrRiseNew(rise int32) *Attribute {
	c_rise := (C.int)(rise)

	retC := C.pango_attr_rise_new(c_rise)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrScaleNew is a wrapper around the C function pango_attr_scale_new.
func AttrScaleNew(scaleFactor float64) *Attribute {
	c_scale_factor := (C.double)(scaleFactor)

	retC := C.pango_attr_scale_new(c_scale_factor)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrStretchNew is a wrapper around the C function pango_attr_stretch_new.
func AttrStretchNew(stretch Stretch) *Attribute {
	c_stretch := (C.PangoStretch)(stretch)

	retC := C.pango_attr_stretch_new(c_stretch)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrStrikethroughNew is a wrapper around the C function pango_attr_strikethrough_new.
func AttrStrikethroughNew(strikethrough bool) *Attribute {
	c_strikethrough :=
		boolToGboolean(strikethrough)

	retC := C.pango_attr_strikethrough_new(c_strikethrough)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrStyleNew is a wrapper around the C function pango_attr_style_new.
func AttrStyleNew(style Style) *Attribute {
	c_style := (C.PangoStyle)(style)

	retC := C.pango_attr_style_new(c_style)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrTypeRegister is a wrapper around the C function pango_attr_type_register.
func AttrTypeRegister(name string) AttrType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.pango_attr_type_register(c_name)
	retGo := (AttrType)(retC)

	return retGo
}

// AttrUnderlineNew is a wrapper around the C function pango_attr_underline_new.
func AttrUnderlineNew(underline Underline) *Attribute {
	c_underline := (C.PangoUnderline)(underline)

	retC := C.pango_attr_underline_new(c_underline)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrVariantNew is a wrapper around the C function pango_attr_variant_new.
func AttrVariantNew(variant Variant) *Attribute {
	c_variant := (C.PangoVariant)(variant)

	retC := C.pango_attr_variant_new(c_variant)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrWeightNew is a wrapper around the C function pango_attr_weight_new.
func AttrWeightNew(weight Weight) *Attribute {
	c_weight := (C.PangoWeight)(weight)

	retC := C.pango_attr_weight_new(c_weight)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_break : unsupported parameter attrs : no param type

// Blacklisted : pango_config_key_get

// Blacklisted : pango_config_key_get_system

// Blacklisted : pango_default_break

// Unsupported : pango_find_map : return type : Blacklisted record : PangoMap

// FindParagraphBoundary is a wrapper around the C function pango_find_paragraph_boundary.
func FindParagraphBoundary(text string, length int32) (int32, int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gint)(length)

	var c_paragraph_delimiter_index C.gint

	var c_next_paragraph_start C.gint

	C.pango_find_paragraph_boundary(c_text, c_length, &c_paragraph_delimiter_index, &c_next_paragraph_start)

	paragraphDelimiterIndex := (int32)(c_paragraph_delimiter_index)

	nextParagraphStart := (int32)(c_next_paragraph_start)

	return paragraphDelimiterIndex, nextParagraphStart
}

// FontDescriptionFromString is a wrapper around the C function pango_font_description_from_string.
func FontDescriptionFromString(str string) *FontDescription {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.pango_font_description_from_string(c_str)
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : pango_get_lib_subdirectory

// Unsupported : pango_get_log_attrs : unsupported parameter log_attrs : no param type

// GetMirrorChar is a wrapper around the C function pango_get_mirror_char.
func GetMirrorChar(ch rune, mirroredCh rune) bool {
	c_ch := (C.gunichar)(ch)

	c_mirrored_ch := (C.gunichar)(mirroredCh)

	retC := C.pango_get_mirror_char(c_ch, &c_mirrored_ch)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : pango_get_sysconf_subdirectory

// Itemize is a wrapper around the C function pango_itemize.
func Itemize(context *Context, text string, startIndex int32, length int32, attrs *AttrList, cachedIter *AttrIterator) *glib.List {
	c_context := (*C.PangoContext)(context.ToC())

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_start_index := (C.int)(startIndex)

	c_length := (C.int)(length)

	c_attrs := (*C.PangoAttrList)(attrs.ToC())

	c_cached_iter := (*C.PangoAttrIterator)(cachedIter.ToC())

	retC := C.pango_itemize(c_context, c_text, c_start_index, c_length, c_attrs, c_cached_iter)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LanguageFromString is a wrapper around the C function pango_language_from_string.
func LanguageFromString(language string) *Language {
	c_language := C.CString(language)
	defer C.free(unsafe.Pointer(c_language))

	retC := C.pango_language_from_string(c_language)
	retGo := LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// ReorderItems is a wrapper around the C function pango_reorder_items.
func ReorderItems(logicalItems *glib.List) *glib.List {
	c_logical_items := (*C.GList)(logicalItems.ToC())

	retC := C.pango_reorder_items(c_logical_items)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_scan_int : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_string : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_word : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_shape : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Unsupported : pango_shape_full : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Unsupported : pango_skip_space : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_split_file_list : no return type

// TrimString is a wrapper around the C function pango_trim_string.
func TrimString(str string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.pango_trim_string(c_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UnicharDirection is a wrapper around the C function pango_unichar_direction.
func UnicharDirection(ch rune) Direction {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_unichar_direction(c_ch)
	retGo := (Direction)(retC)

	return retGo
}
