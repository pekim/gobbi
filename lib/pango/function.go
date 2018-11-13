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

// Create a new background color attribute.
/*

C function

pango_attr_background_new
*/
func AttrBackgroundNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_background_new(c_red, c_green, c_blue)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new font family attribute.
/*

C function

pango_attr_family_new
*/
func AttrFamilyNew(family string) *Attribute {
	c_family := C.CString(family)
	defer C.free(unsafe.Pointer(c_family))

	retC := C.pango_attr_family_new(c_family)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new foreground color attribute.
/*

C function

pango_attr_foreground_new
*/
func AttrForegroundNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_foreground_new(c_red, c_green, c_blue)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new baseline displacement attribute.
/*

C function

pango_attr_rise_new
*/
func AttrRiseNew(rise int32) *Attribute {
	c_rise := (C.int)(rise)

	retC := C.pango_attr_rise_new(c_rise)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new font size scale attribute. The base font for the
// affected text will have its size multiplied by @scale_factor.
/*

C function

pango_attr_scale_new
*/
func AttrScaleNew(scaleFactor float64) *Attribute {
	c_scale_factor := (C.double)(scaleFactor)

	retC := C.pango_attr_scale_new(c_scale_factor)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new font stretch attribute
/*

C function

pango_attr_stretch_new
*/
func AttrStretchNew(stretch Stretch) *Attribute {
	c_stretch := (C.PangoStretch)(stretch)

	retC := C.pango_attr_stretch_new(c_stretch)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new strike-through attribute.
/*

C function

pango_attr_strikethrough_new
*/
func AttrStrikethroughNew(strikethrough bool) *Attribute {
	c_strikethrough :=
		boolToGboolean(strikethrough)

	retC := C.pango_attr_strikethrough_new(c_strikethrough)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new font slant style attribute.
/*

C function

pango_attr_style_new
*/
func AttrStyleNew(style Style) *Attribute {
	c_style := (C.PangoStyle)(style)

	retC := C.pango_attr_style_new(c_style)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Allocate a new attribute type ID.  The attribute type name can be accessed
// later by using pango_attr_type_get_name().
/*

C function

pango_attr_type_register
*/
func AttrTypeRegister(name string) AttrType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.pango_attr_type_register(c_name)
	retGo := (AttrType)(retC)

	return retGo
}

// Create a new underline-style attribute.
/*

C function

pango_attr_underline_new
*/
func AttrUnderlineNew(underline Underline) *Attribute {
	c_underline := (C.PangoUnderline)(underline)

	retC := C.pango_attr_underline_new(c_underline)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new font variant attribute (normal or small caps)
/*

C function

pango_attr_variant_new
*/
func AttrVariantNew(variant Variant) *Attribute {
	c_variant := (C.PangoVariant)(variant)

	retC := C.pango_attr_variant_new(c_variant)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new font weight attribute.
/*

C function

pango_attr_weight_new
*/
func AttrWeightNew(weight Weight) *Attribute {
	c_weight := (C.PangoWeight)(weight)

	retC := C.pango_attr_weight_new(c_weight)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_break : unsupported parameter attrs :

// Blacklisted : pango_config_key_get

// Blacklisted : pango_config_key_get_system

// Blacklisted : pango_default_break

// Unsupported : pango_find_map : return type : Blacklisted record : PangoMap

// Locates a paragraph boundary in @text. A boundary is caused by
// delimiter characters, such as a newline, carriage return, carriage
// return-newline pair, or Unicode paragraph separator character.  The
// index of the run of delimiters is returned in
// @paragraph_delimiter_index. The index of the start of the paragraph
// (index after all delimiters) is stored in @next_paragraph_start.
//
// If no delimiters are found, both @paragraph_delimiter_index and
// @next_paragraph_start are filled with the length of @text (an index one
// off the end).
/*

C function

pango_find_paragraph_boundary
*/
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

// Creates a new font description from a string representation in the
// form "[FAMILY-LIST] [STYLE-OPTIONS] [SIZE]", where FAMILY-LIST is a
// comma separated list of families optionally terminated by a comma,
// STYLE_OPTIONS is a whitespace separated list of words where each word
// describes one of style, variant, weight, stretch, or gravity, and SIZE
// is a decimal number (size in points) or optionally followed by the
// unit modifier "px" for absolute size. Any one of the options may
// be absent.  If FAMILY-LIST is absent, then the family_name field of
// the resulting font description will be initialized to %NULL.  If
// STYLE-OPTIONS is missing, then all style options will be set to the
// default values. If SIZE is missing, the size in the resulting font
// description will be set to 0.
/*

C function

pango_font_description_from_string
*/
func FontDescriptionFromString(str string) *FontDescription {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.pango_font_description_from_string(c_str)
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : pango_get_lib_subdirectory

// Unsupported : pango_get_log_attrs : unsupported parameter log_attrs :

// If @ch has the Unicode mirrored property and there is another Unicode
// character that typically has a glyph that is the mirror image of @ch's
// glyph, puts that character in the address pointed to by @mirrored_ch.
//
// Use g_unichar_get_mirror_char() instead; the docs for that function
// provide full details.
/*

C function

pango_get_mirror_char
*/
func GetMirrorChar(ch rune, mirroredCh rune) bool {
	c_ch := (C.gunichar)(ch)

	c_mirrored_ch := (C.gunichar)(mirroredCh)

	retC := C.pango_get_mirror_char(c_ch, &c_mirrored_ch)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : pango_get_sysconf_subdirectory

// Breaks a piece of text into segments with consistent
// directional level and shaping engine. Each byte of @text will
// be contained in exactly one of the items in the returned list;
// the generated list of items will be in logical order (the start
// offsets of the items are ascending).
//
// @cached_iter should be an iterator over @attrs currently positioned at a
// range before or containing @start_index; @cached_iter will be advanced to
// the range covering the position just after @start_index + @length.
// (i.e. if itemizing in a loop, just keep passing in the same @cached_iter).
/*

C function

pango_itemize
*/
func Itemize(context *Context, text string, startIndex int32, length int32, attrs *AttrList, cachedIter *AttrIterator) *glib.List {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_start_index := (C.int)(startIndex)

	c_length := (C.int)(length)

	c_attrs := (*C.PangoAttrList)(C.NULL)
	if attrs != nil {
		c_attrs = (*C.PangoAttrList)(attrs.ToC())
	}

	c_cached_iter := (*C.PangoAttrIterator)(C.NULL)
	if cachedIter != nil {
		c_cached_iter = (*C.PangoAttrIterator)(cachedIter.ToC())
	}

	retC := C.pango_itemize(c_context, c_text, c_start_index, c_length, c_attrs, c_cached_iter)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Take a RFC-3066 format language tag as a string and convert it to a
// #PangoLanguage pointer that can be efficiently copied (copy the
// pointer) and compared with other language tags (compare the
// pointer.)
//
// This function first canonicalizes the string by converting it to
// lowercase, mapping '_' to '-', and stripping all characters other
// than letters and '-'.
//
// Use pango_language_get_default() if you want to get the #PangoLanguage for
// the current locale of the process.
/*

C function

pango_language_from_string
*/
func LanguageFromString(language string) *Language {
	c_language := C.CString(language)
	defer C.free(unsafe.Pointer(c_language))

	retC := C.pango_language_from_string(c_language)
	var retGo (*Language)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LanguageNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : pango_lookup_aliases : unsupported parameter families : output array param families

// Unsupported : pango_module_register : unsupported parameter module : Blacklisted record : PangoIncludedModule

// Parses marked-up text (see
// <link linkend="PangoMarkupFormat">markup format</link>) to create
// a plain-text string and an attribute list.
//
// If @accel_marker is nonzero, the given character will mark the
// character following it as an accelerator. For example, @accel_marker
// might be an ampersand or underscore. All characters marked
// as an accelerator will receive a %PANGO_UNDERLINE_LOW attribute,
// and the first character so marked will be returned in @accel_char.
// Two @accel_marker characters following each other produce a single
// literal @accel_marker character.
//
// To parse a stream of pango markup incrementally, use pango_markup_parser_new().
//
// If any error happens, none of the output arguments are touched except
// for @error.
/*

C function

pango_parse_markup
*/
func ParseMarkup(markupText string, length int32, accelMarker rune) (bool, *AttrList, string, rune, error) {
	c_markup_text := C.CString(markupText)
	defer C.free(unsafe.Pointer(c_markup_text))

	c_length := (C.int)(length)

	c_accel_marker := (C.gunichar)(accelMarker)

	var c_attr_list *C.PangoAttrList

	var c_text *C.char

	var c_accel_char C.gunichar

	var cThrowableError *C.GError

	retC := C.pango_parse_markup(c_markup_text, c_length, c_accel_marker, &c_attr_list, &c_text, &c_accel_char, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	attrList := AttrListNewFromC(unsafe.Pointer(c_attr_list))

	text := C.GoString(c_text)
	defer C.free(unsafe.Pointer(c_text))

	accelChar := (rune)(c_accel_char)

	return retGo, attrList, text, accelChar, goThrowableError
}

// Unsupported : pango_parse_stretch : unsupported parameter stretch : PangoStretch* with indirection level of 1

// Unsupported : pango_parse_style : unsupported parameter style : PangoStyle* with indirection level of 1

// Unsupported : pango_parse_variant : unsupported parameter variant : PangoVariant* with indirection level of 1

// Unsupported : pango_parse_weight : unsupported parameter weight : PangoWeight* with indirection level of 1

// Unsupported : pango_read_line : unsupported parameter stream : no type generator for gpointer (FILE*) for param stream

// From a list of items in logical order and the associated
// directional levels, produce a list in visual order.
// The original list is unmodified.
/*

C function

pango_reorder_items
*/
func ReorderItems(logicalItems *glib.List) *glib.List {
	c_logical_items := (*C.GList)(C.NULL)
	if logicalItems != nil {
		c_logical_items = (*C.GList)(logicalItems.ToC())
	}

	retC := C.pango_reorder_items(c_logical_items)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_scan_int : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_string : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_word : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_shape : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Unsupported : pango_skip_space : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_split_file_list : no return type

// Trims leading and trailing whitespace from a string.
/*

C function

pango_trim_string
*/
func TrimString(str string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.pango_trim_string(c_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Determines the inherent direction of a character; either
// %PANGO_DIRECTION_LTR, %PANGO_DIRECTION_RTL, or
// %PANGO_DIRECTION_NEUTRAL.
//
// This function is useful to categorize characters into left-to-right
// letters, right-to-left letters, and everything else.  If full
// Unicode bidirectional type of a character is needed,
// pango_bidi_type_for_unichar() can be used instead.
/*

C function

pango_unichar_direction
*/
func UnicharDirection(ch rune) Direction {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_unichar_direction(c_ch)
	retGo := (Direction)(retC)

	return retGo
}
