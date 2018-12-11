// This is a generated file - DO NOT EDIT

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Context) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Context) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Context with another Context, and returns true if they represent the same GObject.
func (recv *Context) Equals(other *Context) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Context) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Context.
// Exercise care, as this is a potentially dangerous function if the Object is not a Context.
func CastToContext(object *gobject.Object) *Context {
	return ContextNewFromC(object.ToC())
}

// ContextNew is a wrapper around the C function pango_context_new.
func ContextNew() *Context {
	retC := C.pango_context_new()
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// Unsupported : pango_context_list_families : unsupported parameter families : output array param families

// LoadFont is a wrapper around the C function pango_context_load_font.
func (recv *Context) LoadFont(desc *FontDescription) *Font {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	retC := C.pango_context_load_font((*C.PangoContext)(recv.native), c_desc)
	var retGo (*Font)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// LoadFontset is a wrapper around the C function pango_context_load_fontset.
func (recv *Context) LoadFontset(desc *FontDescription, language *Language) *Fontset {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	retC := C.pango_context_load_fontset((*C.PangoContext)(recv.native), c_desc, c_language)
	var retGo (*Fontset)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontsetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetBaseDir is a wrapper around the C function pango_context_set_base_dir.
func (recv *Context) SetBaseDir(direction Direction) {
	c_direction := (C.PangoDirection)(direction)

	C.pango_context_set_base_dir((*C.PangoContext)(recv.native), c_direction)

	return
}

// SetFontDescription is a wrapper around the C function pango_context_set_font_description.
func (recv *Context) SetFontDescription(desc *FontDescription) {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	C.pango_context_set_font_description((*C.PangoContext)(recv.native), c_desc)

	return
}

// SetFontMap is a wrapper around the C function pango_context_set_font_map.
func (recv *Context) SetFontMap(fontMap *FontMap) {
	c_font_map := (*C.PangoFontMap)(C.NULL)
	if fontMap != nil {
		c_font_map = (*C.PangoFontMap)(fontMap.ToC())
	}

	C.pango_context_set_font_map((*C.PangoContext)(recv.native), c_font_map)

	return
}

// SetLanguage is a wrapper around the C function pango_context_set_language.
func (recv *Context) SetLanguage(language *Language) {
	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	C.pango_context_set_language((*C.PangoContext)(recv.native), c_language)

	return
}

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EngineLang) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EngineLang) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EngineLang with another EngineLang, and returns true if they represent the same GObject.
func (recv *EngineLang) Equals(other *EngineLang) bool {
	return other.ToC() == recv.ToC()
}

// CastToWidget down casts any arbitary Object to EngineLang.
// Exercise care, as this is a potentially dangerous function if the Object is not a EngineLang.
func CastToEngineLang(object *gobject.Object) *EngineLang {
	return EngineLangNewFromC(object.ToC())
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EngineShape) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EngineShape) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EngineShape with another EngineShape, and returns true if they represent the same GObject.
func (recv *EngineShape) Equals(other *EngineShape) bool {
	return other.ToC() == recv.ToC()
}

// CastToWidget down casts any arbitary Object to EngineShape.
// Exercise care, as this is a potentially dangerous function if the Object is not a EngineShape.
func CastToEngineShape(object *gobject.Object) *EngineShape {
	return EngineShapeNewFromC(object.ToC())
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Font) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Font) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Font with another Font, and returns true if they represent the same GObject.
func (recv *Font) Equals(other *Font) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Font) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Font.
// Exercise care, as this is a potentially dangerous function if the Object is not a Font.
func CastToFont(object *gobject.Object) *Font {
	return FontNewFromC(object.ToC())
}

// pango_font_descriptions_free : unsupported parameter descs :
// Describe is a wrapper around the C function pango_font_describe.
func (recv *Font) Describe() *FontDescription {
	retC := C.pango_font_describe((*C.PangoFont)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindShaper is a wrapper around the C function pango_font_find_shaper.
func (recv *Font) FindShaper(language *Language, ch uint32) *EngineShape {
	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	c_ch := (C.guint32)(ch)

	retC := C.pango_font_find_shaper((*C.PangoFont)(recv.native), c_language, c_ch)
	retGo := EngineShapeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCoverage is a wrapper around the C function pango_font_get_coverage.
func (recv *Font) GetCoverage(language *Language) *Coverage {
	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	retC := C.pango_font_get_coverage((*C.PangoFont)(recv.native), c_language)
	retGo := CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGlyphExtents is a wrapper around the C function pango_font_get_glyph_extents.
func (recv *Font) GetGlyphExtents(glyph Glyph) (*Rectangle, *Rectangle) {
	c_glyph := (C.PangoGlyph)(glyph)

	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_font_get_glyph_extents((*C.PangoFont)(recv.native), c_glyph, &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontFace) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontFace) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FontFace with another FontFace, and returns true if they represent the same GObject.
func (recv *FontFace) Equals(other *FontFace) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FontFace) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FontFace.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontFace.
func CastToFontFace(object *gobject.Object) *FontFace {
	return FontFaceNewFromC(object.ToC())
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontFamily) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontFamily) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FontFamily with another FontFamily, and returns true if they represent the same GObject.
func (recv *FontFamily) Equals(other *FontFamily) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FontFamily) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FontFamily.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontFamily.
func CastToFontFamily(object *gobject.Object) *FontFamily {
	return FontFamilyNewFromC(object.ToC())
}

// GetName is a wrapper around the C function pango_font_family_get_name.
func (recv *FontFamily) GetName() string {
	retC := C.pango_font_family_get_name((*C.PangoFontFamily)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : pango_font_family_list_faces : unsupported parameter faces : output array param faces

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontMap) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FontMap with another FontMap, and returns true if they represent the same GObject.
func (recv *FontMap) Equals(other *FontMap) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FontMap) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FontMap.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontMap.
func CastToFontMap(object *gobject.Object) *FontMap {
	return FontMapNewFromC(object.ToC())
}

// Unsupported : pango_font_map_list_families : unsupported parameter families : output array param families

// LoadFont is a wrapper around the C function pango_font_map_load_font.
func (recv *FontMap) LoadFont(context *Context, desc *FontDescription) *Font {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	retC := C.pango_font_map_load_font((*C.PangoFontMap)(recv.native), c_context, c_desc)
	var retGo (*Font)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// LoadFontset is a wrapper around the C function pango_font_map_load_fontset.
func (recv *FontMap) LoadFontset(context *Context, desc *FontDescription, language *Language) *Fontset {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	retC := C.pango_font_map_load_fontset((*C.PangoFontMap)(recv.native), c_context, c_desc, c_language)
	var retGo (*Fontset)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontsetNewFromC(unsafe.Pointer(retC))
	}

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Fontset) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Fontset) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Fontset with another Fontset, and returns true if they represent the same GObject.
func (recv *Fontset) Equals(other *Fontset) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Fontset) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Fontset.
// Exercise care, as this is a potentially dangerous function if the Object is not a Fontset.
func CastToFontset(object *gobject.Object) *Fontset {
	return FontsetNewFromC(object.ToC())
}

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Layout) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Layout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Layout with another Layout, and returns true if they represent the same GObject.
func (recv *Layout) Equals(other *Layout) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Layout) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Layout.
// Exercise care, as this is a potentially dangerous function if the Object is not a Layout.
func CastToLayout(object *gobject.Object) *Layout {
	return LayoutNewFromC(object.ToC())
}

// LayoutNew is a wrapper around the C function pango_layout_new.
func LayoutNew(context *Context) *Layout {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	retC := C.pango_layout_new(c_context)
	retGo := LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ContextChanged is a wrapper around the C function pango_layout_context_changed.
func (recv *Layout) ContextChanged() {
	C.pango_layout_context_changed((*C.PangoLayout)(recv.native))

	return
}

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

// GetCursorPos is a wrapper around the C function pango_layout_get_cursor_pos.
func (recv *Layout) GetCursorPos(index int32) (*Rectangle, *Rectangle) {
	c_index_ := (C.int)(index)

	var c_strong_pos C.PangoRectangle

	var c_weak_pos C.PangoRectangle

	C.pango_layout_get_cursor_pos((*C.PangoLayout)(recv.native), c_index_, &c_strong_pos, &c_weak_pos)

	strongPos := RectangleNewFromC(unsafe.Pointer(&c_strong_pos))

	weakPos := RectangleNewFromC(unsafe.Pointer(&c_weak_pos))

	return strongPos, weakPos
}

// GetExtents is a wrapper around the C function pango_layout_get_extents.
func (recv *Layout) GetExtents() (*Rectangle, *Rectangle) {
	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_layout_get_extents((*C.PangoLayout)(recv.native), &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

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
	var retGo (*LayoutLine)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LayoutLineNewFromC(unsafe.Pointer(retC))
	}

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

// Unsupported : pango_layout_get_log_attrs : unsupported parameter attrs : output array param attrs

// GetPixelExtents is a wrapper around the C function pango_layout_get_pixel_extents.
func (recv *Layout) GetPixelExtents() (*Rectangle, *Rectangle) {
	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_layout_get_pixel_extents((*C.PangoLayout)(recv.native), &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// GetPixelSize is a wrapper around the C function pango_layout_get_pixel_size.
func (recv *Layout) GetPixelSize() (int32, int32) {
	var c_width C.int

	var c_height C.int

	C.pango_layout_get_pixel_size((*C.PangoLayout)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// GetSingleParagraphMode is a wrapper around the C function pango_layout_get_single_paragraph_mode.
func (recv *Layout) GetSingleParagraphMode() bool {
	retC := C.pango_layout_get_single_paragraph_mode((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSize is a wrapper around the C function pango_layout_get_size.
func (recv *Layout) GetSize() (int32, int32) {
	var c_width C.int

	var c_height C.int

	C.pango_layout_get_size((*C.PangoLayout)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// GetSpacing is a wrapper around the C function pango_layout_get_spacing.
func (recv *Layout) GetSpacing() int32 {
	retC := C.pango_layout_get_spacing((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTabs is a wrapper around the C function pango_layout_get_tabs.
func (recv *Layout) GetTabs() *TabArray {
	retC := C.pango_layout_get_tabs((*C.PangoLayout)(recv.native))
	var retGo (*TabArray)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TabArrayNewFromC(unsafe.Pointer(retC))
	}

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

// IndexToLineX is a wrapper around the C function pango_layout_index_to_line_x.
func (recv *Layout) IndexToLineX(index int32, trailing bool) (int32, int32) {
	c_index_ := (C.int)(index)

	c_trailing :=
		boolToGboolean(trailing)

	var c_line C.int

	var c_x_pos C.int

	C.pango_layout_index_to_line_x((*C.PangoLayout)(recv.native), c_index_, c_trailing, &c_line, &c_x_pos)

	line := (int32)(c_line)

	xPos := (int32)(c_x_pos)

	return line, xPos
}

// IndexToPos is a wrapper around the C function pango_layout_index_to_pos.
func (recv *Layout) IndexToPos(index int32) *Rectangle {
	c_index_ := (C.int)(index)

	var c_pos C.PangoRectangle

	C.pango_layout_index_to_pos((*C.PangoLayout)(recv.native), c_index_, &c_pos)

	pos := RectangleNewFromC(unsafe.Pointer(&c_pos))

	return pos
}

// MoveCursorVisually is a wrapper around the C function pango_layout_move_cursor_visually.
func (recv *Layout) MoveCursorVisually(strong bool, oldIndex int32, oldTrailing int32, direction int32) (int32, int32) {
	c_strong :=
		boolToGboolean(strong)

	c_old_index := (C.int)(oldIndex)

	c_old_trailing := (C.int)(oldTrailing)

	c_direction := (C.int)(direction)

	var c_new_index C.int

	var c_new_trailing C.int

	C.pango_layout_move_cursor_visually((*C.PangoLayout)(recv.native), c_strong, c_old_index, c_old_trailing, c_direction, &c_new_index, &c_new_trailing)

	newIndex := (int32)(c_new_index)

	newTrailing := (int32)(c_new_trailing)

	return newIndex, newTrailing
}

// SetAlignment is a wrapper around the C function pango_layout_set_alignment.
func (recv *Layout) SetAlignment(alignment Alignment) {
	c_alignment := (C.PangoAlignment)(alignment)

	C.pango_layout_set_alignment((*C.PangoLayout)(recv.native), c_alignment)

	return
}

// SetAttributes is a wrapper around the C function pango_layout_set_attributes.
func (recv *Layout) SetAttributes(attrs *AttrList) {
	c_attrs := (*C.PangoAttrList)(C.NULL)
	if attrs != nil {
		c_attrs = (*C.PangoAttrList)(attrs.ToC())
	}

	C.pango_layout_set_attributes((*C.PangoLayout)(recv.native), c_attrs)

	return
}

// SetFontDescription is a wrapper around the C function pango_layout_set_font_description.
func (recv *Layout) SetFontDescription(desc *FontDescription) {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	C.pango_layout_set_font_description((*C.PangoLayout)(recv.native), c_desc)

	return
}

// SetIndent is a wrapper around the C function pango_layout_set_indent.
func (recv *Layout) SetIndent(indent int32) {
	c_indent := (C.int)(indent)

	C.pango_layout_set_indent((*C.PangoLayout)(recv.native), c_indent)

	return
}

// SetJustify is a wrapper around the C function pango_layout_set_justify.
func (recv *Layout) SetJustify(justify bool) {
	c_justify :=
		boolToGboolean(justify)

	C.pango_layout_set_justify((*C.PangoLayout)(recv.native), c_justify)

	return
}

// SetMarkup is a wrapper around the C function pango_layout_set_markup.
func (recv *Layout) SetMarkup(markup string, length int32) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	c_length := (C.int)(length)

	C.pango_layout_set_markup((*C.PangoLayout)(recv.native), c_markup, c_length)

	return
}

// Blacklisted : pango_layout_set_markup_with_accel

// SetSingleParagraphMode is a wrapper around the C function pango_layout_set_single_paragraph_mode.
func (recv *Layout) SetSingleParagraphMode(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.pango_layout_set_single_paragraph_mode((*C.PangoLayout)(recv.native), c_setting)

	return
}

// SetSpacing is a wrapper around the C function pango_layout_set_spacing.
func (recv *Layout) SetSpacing(spacing int32) {
	c_spacing := (C.int)(spacing)

	C.pango_layout_set_spacing((*C.PangoLayout)(recv.native), c_spacing)

	return
}

// SetTabs is a wrapper around the C function pango_layout_set_tabs.
func (recv *Layout) SetTabs(tabs *TabArray) {
	c_tabs := (*C.PangoTabArray)(C.NULL)
	if tabs != nil {
		c_tabs = (*C.PangoTabArray)(tabs.ToC())
	}

	C.pango_layout_set_tabs((*C.PangoLayout)(recv.native), c_tabs)

	return
}

// SetText is a wrapper around the C function pango_layout_set_text.
func (recv *Layout) SetText(text string, length int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.int)(length)

	C.pango_layout_set_text((*C.PangoLayout)(recv.native), c_text, c_length)

	return
}

// SetWidth is a wrapper around the C function pango_layout_set_width.
func (recv *Layout) SetWidth(width int32) {
	c_width := (C.int)(width)

	C.pango_layout_set_width((*C.PangoLayout)(recv.native), c_width)

	return
}

// SetWrap is a wrapper around the C function pango_layout_set_wrap.
func (recv *Layout) SetWrap(wrap WrapMode) {
	c_wrap := (C.PangoWrapMode)(wrap)

	C.pango_layout_set_wrap((*C.PangoLayout)(recv.native), c_wrap)

	return
}

// XyToIndex is a wrapper around the C function pango_layout_xy_to_index.
func (recv *Layout) XyToIndex(x int32, y int32) (bool, int32, int32) {
	c_x := (C.int)(x)

	c_y := (C.int)(y)

	var c_index_ C.int

	var c_trailing C.int

	retC := C.pango_layout_xy_to_index((*C.PangoLayout)(recv.native), c_x, c_y, &c_index_, &c_trailing)
	retGo := retC == C.TRUE

	index := (int32)(c_index_)

	trailing := (int32)(c_trailing)

	return retGo, index, trailing
}
