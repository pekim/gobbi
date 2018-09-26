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

// Context is a wrapper around the C record PangoContext.
type Context struct {
	native *C.PangoContext
}

func ContextNewFromC(u unsafe.Pointer) *Context {
	c := (*C.PangoContext)(u)
	if c == nil {
		return nil
	}

	g := &Context{native: c}

	return g
}

func (recv *Context) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContextNew is a wrapper around the C function pango_context_new.
func ContextNew() *Context {
	retC := C.pango_context_new()
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_context_changed : no return generator

// GetBaseDir is a wrapper around the C function pango_context_get_base_dir.
func (recv *Context) GetBaseDir() Direction {
	retC := C.pango_context_get_base_dir((*C.PangoContext)(recv.native))
	retGo := (Direction)(retC)

	return retGo
}

// GetFontDescription is a wrapper around the C function pango_context_get_font_description.
func (recv *Context) GetFontDescription() *FontDescription {
	retC := C.pango_context_get_font_description((*C.PangoContext)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLanguage is a wrapper around the C function pango_context_get_language.
func (recv *Context) GetLanguage() *Language {
	retC := C.pango_context_get_language((*C.PangoContext)(recv.native))
	retGo := LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_context_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Unsupported : pango_context_list_families : unsupported parameter families : no param type

// LoadFont is a wrapper around the C function pango_context_load_font.
func (recv *Context) LoadFont(desc *FontDescription) *Font {
	c_desc := (*C.PangoFontDescription)(desc.ToC())

	retC := C.pango_context_load_font((*C.PangoContext)(recv.native), c_desc)
	retGo := FontNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LoadFontset is a wrapper around the C function pango_context_load_fontset.
func (recv *Context) LoadFontset(desc *FontDescription, language *Language) *Fontset {
	c_desc := (*C.PangoFontDescription)(desc.ToC())

	c_language := (*C.PangoLanguage)(language.ToC())

	retC := C.pango_context_load_fontset((*C.PangoContext)(recv.native), c_desc, c_language)
	retGo := FontsetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_context_set_base_dir : no return generator

// Unsupported : pango_context_set_base_gravity : no return generator

// Unsupported : pango_context_set_font_description : no return generator

// Unsupported : pango_context_set_font_map : no return generator

// Unsupported : pango_context_set_gravity_hint : no return generator

// Unsupported : pango_context_set_language : no return generator

// Unsupported : pango_context_set_matrix : no return generator

// Blacklisted : PangoEngine

// EngineLang is a wrapper around the C record PangoEngineLang.
type EngineLang struct {
	native *C.PangoEngineLang
	// Private : parent_instance
}

func EngineLangNewFromC(u unsafe.Pointer) *EngineLang {
	c := (*C.PangoEngineLang)(u)
	if c == nil {
		return nil
	}

	g := &EngineLang{native: c}

	return g
}

func (recv *EngineLang) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EngineShape is a wrapper around the C record PangoEngineShape.
type EngineShape struct {
	native *C.PangoEngineShape
	// parent_instance : record
}

func EngineShapeNewFromC(u unsafe.Pointer) *EngineShape {
	c := (*C.PangoEngineShape)(u)
	if c == nil {
		return nil
	}

	g := &EngineShape{native: c}

	return g
}

func (recv *EngineShape) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Font is a wrapper around the C record PangoFont.
type Font struct {
	native *C.PangoFont
	// parent_instance : record
}

func FontNewFromC(u unsafe.Pointer) *Font {
	c := (*C.PangoFont)(u)
	if c == nil {
		return nil
	}

	g := &Font{native: c}

	return g
}

func (recv *Font) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Describe is a wrapper around the C function pango_font_describe.
func (recv *Font) Describe() *FontDescription {
	retC := C.pango_font_describe((*C.PangoFont)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindShaper is a wrapper around the C function pango_font_find_shaper.
func (recv *Font) FindShaper(language *Language, ch uint32) *EngineShape {
	c_language := (*C.PangoLanguage)(language.ToC())

	c_ch := (C.guint32)(ch)

	retC := C.pango_font_find_shaper((*C.PangoFont)(recv.native), c_language, c_ch)
	retGo := EngineShapeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCoverage is a wrapper around the C function pango_font_get_coverage.
func (recv *Font) GetCoverage(language *Language) *Coverage {
	c_language := (*C.PangoLanguage)(language.ToC())

	retC := C.pango_font_get_coverage((*C.PangoFont)(recv.native), c_language)
	retGo := CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_font_get_glyph_extents : no return generator

// Unsupported : pango_font_get_metrics : return type : Blacklisted record : PangoFontMetrics

// FontFace is a wrapper around the C record PangoFontFace.
type FontFace struct {
	native *C.PangoFontFace
	// parent_instance : record
}

func FontFaceNewFromC(u unsafe.Pointer) *FontFace {
	c := (*C.PangoFontFace)(u)
	if c == nil {
		return nil
	}

	g := &FontFace{native: c}

	return g
}

func (recv *FontFace) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Describe is a wrapper around the C function pango_font_face_describe.
func (recv *FontFace) Describe() *FontDescription {
	retC := C.pango_font_face_describe((*C.PangoFontFace)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFaceName is a wrapper around the C function pango_font_face_get_face_name.
func (recv *FontFace) GetFaceName() string {
	retC := C.pango_font_face_get_face_name((*C.PangoFontFace)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : pango_font_face_list_sizes : unsupported parameter sizes : no param type

// FontFamily is a wrapper around the C record PangoFontFamily.
type FontFamily struct {
	native *C.PangoFontFamily
	// parent_instance : record
}

func FontFamilyNewFromC(u unsafe.Pointer) *FontFamily {
	c := (*C.PangoFontFamily)(u)
	if c == nil {
		return nil
	}

	g := &FontFamily{native: c}

	return g
}

func (recv *FontFamily) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetName is a wrapper around the C function pango_font_family_get_name.
func (recv *FontFamily) GetName() string {
	retC := C.pango_font_family_get_name((*C.PangoFontFamily)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : pango_font_family_list_faces : unsupported parameter faces : no param type

// FontMap is a wrapper around the C record PangoFontMap.
type FontMap struct {
	native *C.PangoFontMap
	// parent_instance : record
}

func FontMapNewFromC(u unsafe.Pointer) *FontMap {
	c := (*C.PangoFontMap)(u)
	if c == nil {
		return nil
	}

	g := &FontMap{native: c}

	return g
}

func (recv *FontMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_font_map_changed : no return generator

// Unsupported : pango_font_map_list_families : unsupported parameter families : no param type

// LoadFont is a wrapper around the C function pango_font_map_load_font.
func (recv *FontMap) LoadFont(context *Context, desc *FontDescription) *Font {
	c_context := (*C.PangoContext)(context.ToC())

	c_desc := (*C.PangoFontDescription)(desc.ToC())

	retC := C.pango_font_map_load_font((*C.PangoFontMap)(recv.native), c_context, c_desc)
	retGo := FontNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LoadFontset is a wrapper around the C function pango_font_map_load_fontset.
func (recv *FontMap) LoadFontset(context *Context, desc *FontDescription, language *Language) *Fontset {
	c_context := (*C.PangoContext)(context.ToC())

	c_desc := (*C.PangoFontDescription)(desc.ToC())

	c_language := (*C.PangoLanguage)(language.ToC())

	retC := C.pango_font_map_load_fontset((*C.PangoFontMap)(recv.native), c_context, c_desc, c_language)
	retGo := FontsetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Fontset is a wrapper around the C record PangoFontset.
type Fontset struct {
	native *C.PangoFontset
	// parent_instance : record
}

func FontsetNewFromC(u unsafe.Pointer) *Fontset {
	c := (*C.PangoFontset)(u)
	if c == nil {
		return nil
	}

	g := &Fontset{native: c}

	return g
}

func (recv *Fontset) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc, PangoFontsetForeachFunc

// GetFont is a wrapper around the C function pango_fontset_get_font.
func (recv *Fontset) GetFont(wc uint32) *Font {
	c_wc := (C.guint)(wc)

	retC := C.pango_fontset_get_font((*C.PangoFontset)(recv.native), c_wc)
	retGo := FontNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_fontset_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Blacklisted : PangoFontsetSimple

// Layout is a wrapper around the C record PangoLayout.
type Layout struct {
	native *C.PangoLayout
}

func LayoutNewFromC(u unsafe.Pointer) *Layout {
	c := (*C.PangoLayout)(u)
	if c == nil {
		return nil
	}

	g := &Layout{native: c}

	return g
}

func (recv *Layout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LayoutNew is a wrapper around the C function pango_layout_new.
func LayoutNew(context *Context) *Layout {
	c_context := (*C.PangoContext)(context.ToC())

	retC := C.pango_layout_new(c_context)
	retGo := LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_layout_context_changed : no return generator

// Copy is a wrapper around the C function pango_layout_copy.
func (recv *Layout) Copy() *Layout {
	retC := C.pango_layout_copy((*C.PangoLayout)(recv.native))
	retGo := LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAlignment is a wrapper around the C function pango_layout_get_alignment.
func (recv *Layout) GetAlignment() Alignment {
	retC := C.pango_layout_get_alignment((*C.PangoLayout)(recv.native))
	retGo := (Alignment)(retC)

	return retGo
}

// GetAttributes is a wrapper around the C function pango_layout_get_attributes.
func (recv *Layout) GetAttributes() *AttrList {
	retC := C.pango_layout_get_attributes((*C.PangoLayout)(recv.native))
	retGo := AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContext is a wrapper around the C function pango_layout_get_context.
func (recv *Layout) GetContext() *Context {
	retC := C.pango_layout_get_context((*C.PangoLayout)(recv.native))
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_layout_get_cursor_pos : no return generator

// Unsupported : pango_layout_get_extents : no return generator

// GetIndent is a wrapper around the C function pango_layout_get_indent.
func (recv *Layout) GetIndent() int32 {
	retC := C.pango_layout_get_indent((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : pango_layout_get_iter : return type : Blacklisted record : PangoLayoutIter

// GetJustify is a wrapper around the C function pango_layout_get_justify.
func (recv *Layout) GetJustify() bool {
	retC := C.pango_layout_get_justify((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetLine is a wrapper around the C function pango_layout_get_line.
func (recv *Layout) GetLine(line int32) *LayoutLine {
	c_line := (C.int)(line)

	retC := C.pango_layout_get_line((*C.PangoLayout)(recv.native), c_line)
	retGo := LayoutLineNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLineCount is a wrapper around the C function pango_layout_get_line_count.
func (recv *Layout) GetLineCount() int32 {
	retC := C.pango_layout_get_line_count((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLines is a wrapper around the C function pango_layout_get_lines.
func (recv *Layout) GetLines() *glib.SList {
	retC := C.pango_layout_get_lines((*C.PangoLayout)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_layout_get_log_attrs : unsupported parameter attrs : no param type

// Unsupported : pango_layout_get_log_attrs_readonly : unsupported parameter n_attrs : no type generator for gint, gint*

// Unsupported : pango_layout_get_pixel_extents : no return generator

// Unsupported : pango_layout_get_pixel_size : unsupported parameter width : no type generator for gint, int*

// GetSingleParagraphMode is a wrapper around the C function pango_layout_get_single_paragraph_mode.
func (recv *Layout) GetSingleParagraphMode() bool {
	retC := C.pango_layout_get_single_paragraph_mode((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : pango_layout_get_size : unsupported parameter width : no type generator for gint, int*

// GetSpacing is a wrapper around the C function pango_layout_get_spacing.
func (recv *Layout) GetSpacing() int32 {
	retC := C.pango_layout_get_spacing((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTabs is a wrapper around the C function pango_layout_get_tabs.
func (recv *Layout) GetTabs() *TabArray {
	retC := C.pango_layout_get_tabs((*C.PangoLayout)(recv.native))
	retGo := TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetText is a wrapper around the C function pango_layout_get_text.
func (recv *Layout) GetText() string {
	retC := C.pango_layout_get_text((*C.PangoLayout)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWidth is a wrapper around the C function pango_layout_get_width.
func (recv *Layout) GetWidth() int32 {
	retC := C.pango_layout_get_width((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWrap is a wrapper around the C function pango_layout_get_wrap.
func (recv *Layout) GetWrap() WrapMode {
	retC := C.pango_layout_get_wrap((*C.PangoLayout)(recv.native))
	retGo := (WrapMode)(retC)

	return retGo
}

// Unsupported : pango_layout_index_to_line_x : unsupported parameter line : no type generator for gint, int*

// Unsupported : pango_layout_index_to_pos : no return generator

// Unsupported : pango_layout_move_cursor_visually : unsupported parameter new_index : no type generator for gint, int*

// Unsupported : pango_layout_set_alignment : no return generator

// Unsupported : pango_layout_set_attributes : no return generator

// Unsupported : pango_layout_set_auto_dir : no return generator

// Unsupported : pango_layout_set_ellipsize : no return generator

// Unsupported : pango_layout_set_font_description : no return generator

// Unsupported : pango_layout_set_height : no return generator

// Unsupported : pango_layout_set_indent : no return generator

// Unsupported : pango_layout_set_justify : no return generator

// Unsupported : pango_layout_set_markup : no return generator

// Unsupported : pango_layout_set_markup_with_accel : unsupported parameter accel_char : no type generator for gunichar, gunichar*

// Unsupported : pango_layout_set_single_paragraph_mode : no return generator

// Unsupported : pango_layout_set_spacing : no return generator

// Unsupported : pango_layout_set_tabs : no return generator

// Unsupported : pango_layout_set_text : no return generator

// Unsupported : pango_layout_set_width : no return generator

// Unsupported : pango_layout_set_wrap : no return generator

// Unsupported : pango_layout_xy_to_index : unsupported parameter index_ : no type generator for gint, int*
