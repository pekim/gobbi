// Code generated - DO NOT EDIT.
// +build pango_1.8

package pango

import (
	pango "github.com/pekim/gobbi/lib/internal/c/pango"
	"unsafe"
)

// RenderPart is a representation of the C enumeration PangoRenderPart.
type RenderPart int

// RenderPart_foreground is a representation of the C enumeration member PANGO_RENDER_PART_FOREGROUND.
const RenderPart_foreground = RenderPart(0)

// RenderPart_background is a representation of the C enumeration member PANGO_RENDER_PART_BACKGROUND.
const RenderPart_background = RenderPart(1)

// RenderPart_underline is a representation of the C enumeration member PANGO_RENDER_PART_UNDERLINE.
const RenderPart_underline = RenderPart(2)

// RenderPart_strikethrough is a representation of the C enumeration member PANGO_RENDER_PART_STRIKETHROUGH.
const RenderPart_strikethrough = RenderPart(3)

// AttrStrikethroughColorNew wraps the C function pango_attr_strikethrough_color_new.
//
// since 1.8
func AttrStrikethroughColorNew(red uint16, green uint16, blue uint16) *Attribute {
	sys_red := red
	sys_green := green
	sys_blue := blue
	retSys := pango.Fn_pango_attr_strikethrough_color_new(sys_red, sys_green, sys_blue)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrUnderlineColorNew wraps the C function pango_attr_underline_color_new.
//
// since 1.8
func AttrUnderlineColorNew(red uint16, green uint16, blue uint16) *Attribute {
	sys_red := red
	sys_green := green
	sys_blue := blue
	retSys := pango.Fn_pango_attr_underline_color_new(sys_red, sys_green, sys_blue)
	ret := AttributeNewFromC(retSys)

	return ret
}

// UNSUPPORTED : pango_break : has array param, attrs

// UNSUPPORTED : pango_config_key_get : blacklisted

// UNSUPPORTED : pango_config_key_get_system : blacklisted

// UNSUPPORTED : pango_default_break : blacklisted

// UNSUPPORTED : pango_find_map : blacklisted

// UNSUPPORTED : pango_find_paragraph_boundary : has [in]out param, paragraph_delimiter_index

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted

// UNSUPPORTED : pango_get_log_attrs : has array param, log_attrs

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted

// UNSUPPORTED : pango_lookup_aliases : blacklisted

// UNSUPPORTED : pango_markup_parser_finish : throws

// UNSUPPORTED : pango_module_register : blacklisted

// UNSUPPORTED : pango_parse_enum : has [in]out param, value

// UNSUPPORTED : pango_parse_markup : throws

// UNSUPPORTED : pango_parse_stretch : has [in]out param, stretch

// UNSUPPORTED : pango_parse_style : has [in]out param, style

// UNSUPPORTED : pango_parse_variant : has [in]out param, variant

// UNSUPPORTED : pango_parse_weight : has [in]out param, weight

// UNSUPPORTED : pango_read_line : has [in]out param, str

// UNSUPPORTED : pango_scan_int : has [in]out param, out

// UNSUPPORTED : pango_scan_string : has [in]out param, out

// UNSUPPORTED : pango_scan_word : has [in]out param, out

// UNSUPPORTED : pango_split_file_list : no array length

// UNSUPPORTED : EngineClass : blacklisted

// UNSUPPORTED : EngineInfo : blacklisted

// UNSUPPORTED : EngineLangClass : blacklisted

// UNSUPPORTED : EngineScriptInfo : blacklisted

// UNSUPPORTED : EngineShapeClass : blacklisted

// UNSUPPORTED : FontClass : blacklisted

// UNSUPPORTED : FontFaceClass : blacklisted

// UNSUPPORTED : FontFamilyClass : blacklisted

// UNSUPPORTED : FontMapClass : blacklisted

// UNSUPPORTED : FontsetClass : blacklisted

// UNSUPPORTED : FontsetSimpleClass : blacklisted

// UNSUPPORTED : IncludedModule : blacklisted

// UNSUPPORTED : Map : blacklisted

// UNSUPPORTED : MapEntry : blacklisted

// RendererClass is a representation of the C record PangoRendererClass.
//
// since 1.8
type RendererClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoRendererClass that represents the RendererClass.
func (recv *RendererClass) ToC() unsafe.Pointer {
	return recv.native
}

// RendererClassNewFromC creates a new RendererClass from a pointer to the C PangoRendererClass that represents the RendererClass.
func RendererClassNewFromC(native unsafe.Pointer) *RendererClass {
	return &RendererClass{native: native}
}

// UNSUPPORTED : Engine : blacklisted

// UNSUPPORTED : FontsetSimple : blacklisted

// Renderer is a representation of the C record PangoRenderer.
//
// since 1.8
type Renderer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoRenderer that represents the Renderer.
func (recv *Renderer) ToC() unsafe.Pointer {
	return recv.native
}

// RendererNewFromC creates a new Renderer from a pointer to the C PangoRenderer that represents the Renderer.
func RendererNewFromC(native unsafe.Pointer) *Renderer {
	return &Renderer{native: native}
}
