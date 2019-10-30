// Code generated - DO NOT EDIT.
// +build pango_1.16

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"reflect"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

var gobjectClassGoTypeMap = make(map[string]reflect.Type)

// Glyph is a representation of the C alias PangoGlyph.
type Glyph uint32

// GlyphUnit is a representation of the C alias PangoGlyphUnit.
type GlyphUnit int32

// LayoutRun is a representation of the C alias PangoLayoutRun.
type LayoutRun *GlyphItem

type FontMask C.PangoFontMask

const (
	PANGO_FONT_MASK_FAMILY  FontMask = 1
	PANGO_FONT_MASK_STYLE   FontMask = 2
	PANGO_FONT_MASK_VARIANT FontMask = 4
	PANGO_FONT_MASK_WEIGHT  FontMask = 8
	PANGO_FONT_MASK_STRETCH FontMask = 16
	PANGO_FONT_MASK_SIZE    FontMask = 32
	PANGO_FONT_MASK_GRAVITY FontMask = 64
)

// AddToGobjectClassGoTypeMap : PangoContext

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

// CastToWidget down casts any arbitrary Object to Context.
// Exercise care, as this is a potentially dangerous function if the Object is not a Context.
func CastToContext(object *gobject.Object) *Context {
	return ContextNewFromC(object.ToC())
}

// ContextNew is a wrapper around the C function pango_context_new.
func ContextNew() *Context {
	retC := C.pango_context_new()
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetBaseDir is a wrapper around the C function pango_context_get_base_dir.
func (recv *Context) GetBaseDir() Direction {
	retC := C.pango_context_get_base_dir((*C.PangoContext)(recv.native))
	retGo := (Direction)(retC)

	return retGo
}

// GetBaseGravity is a wrapper around the C function pango_context_get_base_gravity.
func (recv *Context) GetBaseGravity() Gravity {
	retC := C.pango_context_get_base_gravity((*C.PangoContext)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// GetFontDescription is a wrapper around the C function pango_context_get_font_description.
func (recv *Context) GetFontDescription() *FontDescription {
	retC := C.pango_context_get_font_description((*C.PangoContext)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFontMap is a wrapper around the C function pango_context_get_font_map.
func (recv *Context) GetFontMap() *FontMap {
	retC := C.pango_context_get_font_map((*C.PangoContext)(recv.native))
	retGo := FontMapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGravity is a wrapper around the C function pango_context_get_gravity.
func (recv *Context) GetGravity() Gravity {
	retC := C.pango_context_get_gravity((*C.PangoContext)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// GetGravityHint is a wrapper around the C function pango_context_get_gravity_hint.
func (recv *Context) GetGravityHint() GravityHint {
	retC := C.pango_context_get_gravity_hint((*C.PangoContext)(recv.native))
	retGo := (GravityHint)(retC)

	return retGo
}

// GetLanguage is a wrapper around the C function pango_context_get_language.
func (recv *Context) GetLanguage() *Language {
	retC := C.pango_context_get_language((*C.PangoContext)(recv.native))
	retGo := LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMatrix is a wrapper around the C function pango_context_get_matrix.
func (recv *Context) GetMatrix() *Matrix {
	retC := C.pango_context_get_matrix((*C.PangoContext)(recv.native))
	var retGo (*Matrix)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MatrixNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetMetrics is a wrapper around the C function pango_context_get_metrics.
func (recv *Context) GetMetrics(desc *FontDescription, language *Language) *FontMetrics {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	retC := C.pango_context_get_metrics((*C.PangoContext)(recv.native), c_desc, c_language)
	retGo := FontMetricsNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// SetBaseGravity is a wrapper around the C function pango_context_set_base_gravity.
func (recv *Context) SetBaseGravity(gravity Gravity) {
	c_gravity := (C.PangoGravity)(gravity)

	C.pango_context_set_base_gravity((*C.PangoContext)(recv.native), c_gravity)

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

// SetGravityHint is a wrapper around the C function pango_context_set_gravity_hint.
func (recv *Context) SetGravityHint(hint GravityHint) {
	c_hint := (C.PangoGravityHint)(hint)

	C.pango_context_set_gravity_hint((*C.PangoContext)(recv.native), c_hint)

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

// SetMatrix is a wrapper around the C function pango_context_set_matrix.
func (recv *Context) SetMatrix(matrix *Matrix) {
	c_matrix := (*C.PangoMatrix)(C.NULL)
	if matrix != nil {
		c_matrix = (*C.PangoMatrix)(matrix.ToC())
	}

	C.pango_context_set_matrix((*C.PangoContext)(recv.native), c_matrix)

	return
}

// Blacklisted : PangoEngine

// AddToGobjectClassGoTypeMap : PangoEngineLang

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

// CastToWidget down casts any arbitrary Object to EngineLang.
// Exercise care, as this is a potentially dangerous function if the Object is not a EngineLang.
func CastToEngineLang(object *gobject.Object) *EngineLang {
	return EngineLangNewFromC(object.ToC())
}

// AddToGobjectClassGoTypeMap : PangoEngineShape

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

// CastToWidget down casts any arbitrary Object to EngineShape.
// Exercise care, as this is a potentially dangerous function if the Object is not a EngineShape.
func CastToEngineShape(object *gobject.Object) *EngineShape {
	return EngineShapeNewFromC(object.ToC())
}

// AddToGobjectClassGoTypeMap : PangoFont

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

// CastToWidget down casts any arbitrary Object to Font.
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

// DescribeWithAbsoluteSize is a wrapper around the C function pango_font_describe_with_absolute_size.
func (recv *Font) DescribeWithAbsoluteSize() *FontDescription {
	retC := C.pango_font_describe_with_absolute_size((*C.PangoFont)(recv.native))
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

// GetFontMap is a wrapper around the C function pango_font_get_font_map.
func (recv *Font) GetFontMap() *FontMap {
	retC := C.pango_font_get_font_map((*C.PangoFont)(recv.native))
	var retGo (*FontMap)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontMapNewFromC(unsafe.Pointer(retC))
	}

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

// GetMetrics is a wrapper around the C function pango_font_get_metrics.
func (recv *Font) GetMetrics(language *Language) *FontMetrics {
	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	retC := C.pango_font_get_metrics((*C.PangoFont)(recv.native), c_language)
	retGo := FontMetricsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddToGobjectClassGoTypeMap : PangoFontFace

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

// CastToWidget down casts any arbitrary Object to FontFace.
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

// Unsupported : pango_font_face_list_sizes : unsupported parameter sizes : output array param sizes

// AddToGobjectClassGoTypeMap : PangoFontFamily

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

// CastToWidget down casts any arbitrary Object to FontFamily.
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

// IsMonospace is a wrapper around the C function pango_font_family_is_monospace.
func (recv *FontFamily) IsMonospace() bool {
	retC := C.pango_font_family_is_monospace((*C.PangoFontFamily)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : pango_font_family_list_faces : unsupported parameter faces : output array param faces

// AddToGobjectClassGoTypeMap : PangoFontMap

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

// CastToWidget down casts any arbitrary Object to FontMap.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontMap.
func CastToFontMap(object *gobject.Object) *FontMap {
	return FontMapNewFromC(object.ToC())
}

// Blacklisted : pango_font_map_get_shape_engine_type

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

// AddToGobjectClassGoTypeMap : PangoFontset

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

// CastToWidget down casts any arbitrary Object to Fontset.
// Exercise care, as this is a potentially dangerous function if the Object is not a Fontset.
func CastToFontset(object *gobject.Object) *Fontset {
	return FontsetNewFromC(object.ToC())
}

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc (PangoFontsetForeachFunc) for param func

// GetFont is a wrapper around the C function pango_fontset_get_font.
func (recv *Fontset) GetFont(wc uint32) *Font {
	c_wc := (C.guint)(wc)

	retC := C.pango_fontset_get_font((*C.PangoFontset)(recv.native), c_wc)
	retGo := FontNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMetrics is a wrapper around the C function pango_fontset_get_metrics.
func (recv *Fontset) GetMetrics() *FontMetrics {
	retC := C.pango_fontset_get_metrics((*C.PangoFontset)(recv.native))
	retGo := FontMetricsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : PangoFontsetSimple

// AddToGobjectClassGoTypeMap : PangoLayout

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

// CastToWidget down casts any arbitrary Object to Layout.
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

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

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

// GetAutoDir is a wrapper around the C function pango_layout_get_auto_dir.
func (recv *Layout) GetAutoDir() bool {
	retC := C.pango_layout_get_auto_dir((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

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

// GetEllipsize is a wrapper around the C function pango_layout_get_ellipsize.
func (recv *Layout) GetEllipsize() EllipsizeMode {
	retC := C.pango_layout_get_ellipsize((*C.PangoLayout)(recv.native))
	retGo := (EllipsizeMode)(retC)

	return retGo
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

// GetFontDescription is a wrapper around the C function pango_layout_get_font_description.
func (recv *Layout) GetFontDescription() *FontDescription {
	retC := C.pango_layout_get_font_description((*C.PangoLayout)(recv.native))
	var retGo (*FontDescription)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontDescriptionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
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

// GetLineReadonly is a wrapper around the C function pango_layout_get_line_readonly.
func (recv *Layout) GetLineReadonly(line int32) *LayoutLine {
	c_line := (C.int)(line)

	retC := C.pango_layout_get_line_readonly((*C.PangoLayout)(recv.native), c_line)
	var retGo (*LayoutLine)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LayoutLineNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetLines is a wrapper around the C function pango_layout_get_lines.
func (recv *Layout) GetLines() *glib.SList {
	retC := C.pango_layout_get_lines((*C.PangoLayout)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLinesReadonly is a wrapper around the C function pango_layout_get_lines_readonly.
func (recv *Layout) GetLinesReadonly() *glib.SList {
	retC := C.pango_layout_get_lines_readonly((*C.PangoLayout)(recv.native))
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

// GetUnknownGlyphsCount is a wrapper around the C function pango_layout_get_unknown_glyphs_count.
func (recv *Layout) GetUnknownGlyphsCount() int32 {
	retC := C.pango_layout_get_unknown_glyphs_count((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

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

// IsEllipsized is a wrapper around the C function pango_layout_is_ellipsized.
func (recv *Layout) IsEllipsized() bool {
	retC := C.pango_layout_is_ellipsized((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsWrapped is a wrapper around the C function pango_layout_is_wrapped.
func (recv *Layout) IsWrapped() bool {
	retC := C.pango_layout_is_wrapped((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
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

// SetAutoDir is a wrapper around the C function pango_layout_set_auto_dir.
func (recv *Layout) SetAutoDir(autoDir bool) {
	c_auto_dir :=
		boolToGboolean(autoDir)

	C.pango_layout_set_auto_dir((*C.PangoLayout)(recv.native), c_auto_dir)

	return
}

// SetEllipsize is a wrapper around the C function pango_layout_set_ellipsize.
func (recv *Layout) SetEllipsize(ellipsize EllipsizeMode) {
	c_ellipsize := (C.PangoEllipsizeMode)(ellipsize)

	C.pango_layout_set_ellipsize((*C.PangoLayout)(recv.native), c_ellipsize)

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

// AddToGobjectClassGoTypeMap : PangoRenderer

// Renderer is a wrapper around the C record PangoRenderer.
type Renderer struct {
	native *C.PangoRenderer
	// Private : parent_instance
	// Private : underline
	// Private : strikethrough
	// Private : active_count
	// matrix : record
	// Private : priv
}

func RendererNewFromC(u unsafe.Pointer) *Renderer {
	c := (*C.PangoRenderer)(u)
	if c == nil {
		return nil
	}

	g := &Renderer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Renderer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Renderer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Renderer with another Renderer, and returns true if they represent the same GObject.
func (recv *Renderer) Equals(other *Renderer) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Renderer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Renderer.
// Exercise care, as this is a potentially dangerous function if the Object is not a Renderer.
func CastToRenderer(object *gobject.Object) *Renderer {
	return RendererNewFromC(object.ToC())
}

// Activate is a wrapper around the C function pango_renderer_activate.
func (recv *Renderer) Activate() {
	C.pango_renderer_activate((*C.PangoRenderer)(recv.native))

	return
}

// Deactivate is a wrapper around the C function pango_renderer_deactivate.
func (recv *Renderer) Deactivate() {
	C.pango_renderer_deactivate((*C.PangoRenderer)(recv.native))

	return
}

// DrawErrorUnderline is a wrapper around the C function pango_renderer_draw_error_underline.
func (recv *Renderer) DrawErrorUnderline(x int32, y int32, width int32, height int32) {
	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.pango_renderer_draw_error_underline((*C.PangoRenderer)(recv.native), c_x, c_y, c_width, c_height)

	return
}

// DrawGlyph is a wrapper around the C function pango_renderer_draw_glyph.
func (recv *Renderer) DrawGlyph(font *Font, glyph Glyph, x float64, y float64) {
	c_font := (*C.PangoFont)(C.NULL)
	if font != nil {
		c_font = (*C.PangoFont)(font.ToC())
	}

	c_glyph := (C.PangoGlyph)(glyph)

	c_x := (C.double)(x)

	c_y := (C.double)(y)

	C.pango_renderer_draw_glyph((*C.PangoRenderer)(recv.native), c_font, c_glyph, c_x, c_y)

	return
}

// Unsupported : pango_renderer_draw_glyphs : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// DrawLayout is a wrapper around the C function pango_renderer_draw_layout.
func (recv *Renderer) DrawLayout(layout *Layout, x int32, y int32) {
	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_renderer_draw_layout((*C.PangoRenderer)(recv.native), c_layout, c_x, c_y)

	return
}

// DrawLayoutLine is a wrapper around the C function pango_renderer_draw_layout_line.
func (recv *Renderer) DrawLayoutLine(line *LayoutLine, x int32, y int32) {
	c_line := (*C.PangoLayoutLine)(C.NULL)
	if line != nil {
		c_line = (*C.PangoLayoutLine)(line.ToC())
	}

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_renderer_draw_layout_line((*C.PangoRenderer)(recv.native), c_line, c_x, c_y)

	return
}

// DrawRectangle is a wrapper around the C function pango_renderer_draw_rectangle.
func (recv *Renderer) DrawRectangle(part RenderPart, x int32, y int32, width int32, height int32) {
	c_part := (C.PangoRenderPart)(part)

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.pango_renderer_draw_rectangle((*C.PangoRenderer)(recv.native), c_part, c_x, c_y, c_width, c_height)

	return
}

// DrawTrapezoid is a wrapper around the C function pango_renderer_draw_trapezoid.
func (recv *Renderer) DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64) {
	c_part := (C.PangoRenderPart)(part)

	c_y1_ := (C.double)(y1)

	c_x11 := (C.double)(x11)

	c_x21 := (C.double)(x21)

	c_y2 := (C.double)(y2)

	c_x12 := (C.double)(x12)

	c_x22 := (C.double)(x22)

	C.pango_renderer_draw_trapezoid((*C.PangoRenderer)(recv.native), c_part, c_y1_, c_x11, c_x21, c_y2, c_x12, c_x22)

	return
}

// GetColor is a wrapper around the C function pango_renderer_get_color.
func (recv *Renderer) GetColor(part RenderPart) *Color {
	c_part := (C.PangoRenderPart)(part)

	retC := C.pango_renderer_get_color((*C.PangoRenderer)(recv.native), c_part)
	var retGo (*Color)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ColorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetMatrix is a wrapper around the C function pango_renderer_get_matrix.
func (recv *Renderer) GetMatrix() *Matrix {
	retC := C.pango_renderer_get_matrix((*C.PangoRenderer)(recv.native))
	var retGo (*Matrix)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MatrixNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PartChanged is a wrapper around the C function pango_renderer_part_changed.
func (recv *Renderer) PartChanged(part RenderPart) {
	c_part := (C.PangoRenderPart)(part)

	C.pango_renderer_part_changed((*C.PangoRenderer)(recv.native), c_part)

	return
}

// SetColor is a wrapper around the C function pango_renderer_set_color.
func (recv *Renderer) SetColor(part RenderPart, color *Color) {
	c_part := (C.PangoRenderPart)(part)

	c_color := (*C.PangoColor)(C.NULL)
	if color != nil {
		c_color = (*C.PangoColor)(color.ToC())
	}

	C.pango_renderer_set_color((*C.PangoRenderer)(recv.native), c_part, c_color)

	return
}

// SetMatrix is a wrapper around the C function pango_renderer_set_matrix.
func (recv *Renderer) SetMatrix(matrix *Matrix) {
	c_matrix := (*C.PangoMatrix)(C.NULL)
	if matrix != nil {
		c_matrix = (*C.PangoMatrix)(matrix.ToC())
	}

	C.pango_renderer_set_matrix((*C.PangoRenderer)(recv.native), c_matrix)

	return
}

const ANALYSIS_FLAG_CENTERED_BASELINE int32 = C.PANGO_ANALYSIS_FLAG_CENTERED_BASELINE

// Blacklisted : ENGINE_TYPE_LANG

// Blacklisted : ENGINE_TYPE_SHAPE

// Unsupported : type Glyph for GLYPH_EMPTY

// Unsupported : type Glyph for GLYPH_UNKNOWN_FLAG

// Blacklisted : RENDER_TYPE_NONE

const SCALE int32 = C.PANGO_SCALE

// Blacklisted : UNKNOWN_GLYPH_HEIGHT

// Blacklisted : UNKNOWN_GLYPH_WIDTH

type Alignment C.PangoAlignment

const (
	PANGO_ALIGN_LEFT   Alignment = 0
	PANGO_ALIGN_CENTER Alignment = 1
	PANGO_ALIGN_RIGHT  Alignment = 2
)

type AttrType C.PangoAttrType

const (
	PANGO_ATTR_INVALID             AttrType = 0
	PANGO_ATTR_LANGUAGE            AttrType = 1
	PANGO_ATTR_FAMILY              AttrType = 2
	PANGO_ATTR_STYLE               AttrType = 3
	PANGO_ATTR_WEIGHT              AttrType = 4
	PANGO_ATTR_VARIANT             AttrType = 5
	PANGO_ATTR_STRETCH             AttrType = 6
	PANGO_ATTR_SIZE                AttrType = 7
	PANGO_ATTR_FONT_DESC           AttrType = 8
	PANGO_ATTR_FOREGROUND          AttrType = 9
	PANGO_ATTR_BACKGROUND          AttrType = 10
	PANGO_ATTR_UNDERLINE           AttrType = 11
	PANGO_ATTR_STRIKETHROUGH       AttrType = 12
	PANGO_ATTR_RISE                AttrType = 13
	PANGO_ATTR_SHAPE               AttrType = 14
	PANGO_ATTR_SCALE               AttrType = 15
	PANGO_ATTR_FALLBACK            AttrType = 16
	PANGO_ATTR_LETTER_SPACING      AttrType = 17
	PANGO_ATTR_UNDERLINE_COLOR     AttrType = 18
	PANGO_ATTR_STRIKETHROUGH_COLOR AttrType = 19
	PANGO_ATTR_ABSOLUTE_SIZE       AttrType = 20
	PANGO_ATTR_GRAVITY             AttrType = 21
	PANGO_ATTR_GRAVITY_HINT        AttrType = 22
	PANGO_ATTR_FONT_FEATURES       AttrType = 23
	PANGO_ATTR_FOREGROUND_ALPHA    AttrType = 24
	PANGO_ATTR_BACKGROUND_ALPHA    AttrType = 25
)

// AttrTypeRegister is a wrapper around the C function pango_attr_type_register.
func AttrTypeRegister(name string) AttrType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.pango_attr_type_register(c_name)
	retGo := (AttrType)(retC)

	return retGo
}

type CoverageLevel C.PangoCoverageLevel

const (
	PANGO_COVERAGE_NONE        CoverageLevel = 0
	PANGO_COVERAGE_FALLBACK    CoverageLevel = 1
	PANGO_COVERAGE_APPROXIMATE CoverageLevel = 2
	PANGO_COVERAGE_EXACT       CoverageLevel = 3
)

type Direction C.PangoDirection

const (
	PANGO_DIRECTION_LTR      Direction = 0
	PANGO_DIRECTION_RTL      Direction = 1
	PANGO_DIRECTION_TTB_LTR  Direction = 2
	PANGO_DIRECTION_TTB_RTL  Direction = 3
	PANGO_DIRECTION_WEAK_LTR Direction = 4
	PANGO_DIRECTION_WEAK_RTL Direction = 5
	PANGO_DIRECTION_NEUTRAL  Direction = 6
)

type EllipsizeMode C.PangoEllipsizeMode

const (
	PANGO_ELLIPSIZE_NONE   EllipsizeMode = 0
	PANGO_ELLIPSIZE_START  EllipsizeMode = 1
	PANGO_ELLIPSIZE_MIDDLE EllipsizeMode = 2
	PANGO_ELLIPSIZE_END    EllipsizeMode = 3
)

type Gravity C.PangoGravity

const (
	PANGO_GRAVITY_SOUTH Gravity = 0
	PANGO_GRAVITY_EAST  Gravity = 1
	PANGO_GRAVITY_NORTH Gravity = 2
	PANGO_GRAVITY_WEST  Gravity = 3
	PANGO_GRAVITY_AUTO  Gravity = 4
)

// GravityGetForMatrix is a wrapper around the C function pango_gravity_get_for_matrix.
func GravityGetForMatrix(matrix *Matrix) Gravity {
	c_matrix := (*C.PangoMatrix)(C.NULL)
	if matrix != nil {
		c_matrix = (*C.PangoMatrix)(matrix.ToC())
	}

	retC := C.pango_gravity_get_for_matrix(c_matrix)
	retGo := (Gravity)(retC)

	return retGo
}

// GravityGetForScript is a wrapper around the C function pango_gravity_get_for_script.
func GravityGetForScript(script Script, baseGravity Gravity, hint GravityHint) Gravity {
	c_script := (C.PangoScript)(script)

	c_base_gravity := (C.PangoGravity)(baseGravity)

	c_hint := (C.PangoGravityHint)(hint)

	retC := C.pango_gravity_get_for_script(c_script, c_base_gravity, c_hint)
	retGo := (Gravity)(retC)

	return retGo
}

// GravityToRotation is a wrapper around the C function pango_gravity_to_rotation.
func GravityToRotation(gravity Gravity) float64 {
	c_gravity := (C.PangoGravity)(gravity)

	retC := C.pango_gravity_to_rotation(c_gravity)
	retGo := (float64)(retC)

	return retGo
}

type GravityHint C.PangoGravityHint

const (
	PANGO_GRAVITY_HINT_NATURAL GravityHint = 0
	PANGO_GRAVITY_HINT_STRONG  GravityHint = 1
	PANGO_GRAVITY_HINT_LINE    GravityHint = 2
)

type RenderPart C.PangoRenderPart

const (
	PANGO_RENDER_PART_FOREGROUND    RenderPart = 0
	PANGO_RENDER_PART_BACKGROUND    RenderPart = 1
	PANGO_RENDER_PART_UNDERLINE     RenderPart = 2
	PANGO_RENDER_PART_STRIKETHROUGH RenderPart = 3
)

type Script C.PangoScript

const (
	PANGO_SCRIPT_INVALID_CODE          Script = -1
	PANGO_SCRIPT_COMMON                Script = 0
	PANGO_SCRIPT_INHERITED             Script = 1
	PANGO_SCRIPT_ARABIC                Script = 2
	PANGO_SCRIPT_ARMENIAN              Script = 3
	PANGO_SCRIPT_BENGALI               Script = 4
	PANGO_SCRIPT_BOPOMOFO              Script = 5
	PANGO_SCRIPT_CHEROKEE              Script = 6
	PANGO_SCRIPT_COPTIC                Script = 7
	PANGO_SCRIPT_CYRILLIC              Script = 8
	PANGO_SCRIPT_DESERET               Script = 9
	PANGO_SCRIPT_DEVANAGARI            Script = 10
	PANGO_SCRIPT_ETHIOPIC              Script = 11
	PANGO_SCRIPT_GEORGIAN              Script = 12
	PANGO_SCRIPT_GOTHIC                Script = 13
	PANGO_SCRIPT_GREEK                 Script = 14
	PANGO_SCRIPT_GUJARATI              Script = 15
	PANGO_SCRIPT_GURMUKHI              Script = 16
	PANGO_SCRIPT_HAN                   Script = 17
	PANGO_SCRIPT_HANGUL                Script = 18
	PANGO_SCRIPT_HEBREW                Script = 19
	PANGO_SCRIPT_HIRAGANA              Script = 20
	PANGO_SCRIPT_KANNADA               Script = 21
	PANGO_SCRIPT_KATAKANA              Script = 22
	PANGO_SCRIPT_KHMER                 Script = 23
	PANGO_SCRIPT_LAO                   Script = 24
	PANGO_SCRIPT_LATIN                 Script = 25
	PANGO_SCRIPT_MALAYALAM             Script = 26
	PANGO_SCRIPT_MONGOLIAN             Script = 27
	PANGO_SCRIPT_MYANMAR               Script = 28
	PANGO_SCRIPT_OGHAM                 Script = 29
	PANGO_SCRIPT_OLD_ITALIC            Script = 30
	PANGO_SCRIPT_ORIYA                 Script = 31
	PANGO_SCRIPT_RUNIC                 Script = 32
	PANGO_SCRIPT_SINHALA               Script = 33
	PANGO_SCRIPT_SYRIAC                Script = 34
	PANGO_SCRIPT_TAMIL                 Script = 35
	PANGO_SCRIPT_TELUGU                Script = 36
	PANGO_SCRIPT_THAANA                Script = 37
	PANGO_SCRIPT_THAI                  Script = 38
	PANGO_SCRIPT_TIBETAN               Script = 39
	PANGO_SCRIPT_CANADIAN_ABORIGINAL   Script = 40
	PANGO_SCRIPT_YI                    Script = 41
	PANGO_SCRIPT_TAGALOG               Script = 42
	PANGO_SCRIPT_HANUNOO               Script = 43
	PANGO_SCRIPT_BUHID                 Script = 44
	PANGO_SCRIPT_TAGBANWA              Script = 45
	PANGO_SCRIPT_BRAILLE               Script = 46
	PANGO_SCRIPT_CYPRIOT               Script = 47
	PANGO_SCRIPT_LIMBU                 Script = 48
	PANGO_SCRIPT_OSMANYA               Script = 49
	PANGO_SCRIPT_SHAVIAN               Script = 50
	PANGO_SCRIPT_LINEAR_B              Script = 51
	PANGO_SCRIPT_TAI_LE                Script = 52
	PANGO_SCRIPT_UGARITIC              Script = 53
	PANGO_SCRIPT_NEW_TAI_LUE           Script = 54
	PANGO_SCRIPT_BUGINESE              Script = 55
	PANGO_SCRIPT_GLAGOLITIC            Script = 56
	PANGO_SCRIPT_TIFINAGH              Script = 57
	PANGO_SCRIPT_SYLOTI_NAGRI          Script = 58
	PANGO_SCRIPT_OLD_PERSIAN           Script = 59
	PANGO_SCRIPT_KHAROSHTHI            Script = 60
	PANGO_SCRIPT_UNKNOWN               Script = 61
	PANGO_SCRIPT_BALINESE              Script = 62
	PANGO_SCRIPT_CUNEIFORM             Script = 63
	PANGO_SCRIPT_PHOENICIAN            Script = 64
	PANGO_SCRIPT_PHAGS_PA              Script = 65
	PANGO_SCRIPT_NKO                   Script = 66
	PANGO_SCRIPT_KAYAH_LI              Script = 67
	PANGO_SCRIPT_LEPCHA                Script = 68
	PANGO_SCRIPT_REJANG                Script = 69
	PANGO_SCRIPT_SUNDANESE             Script = 70
	PANGO_SCRIPT_SAURASHTRA            Script = 71
	PANGO_SCRIPT_CHAM                  Script = 72
	PANGO_SCRIPT_OL_CHIKI              Script = 73
	PANGO_SCRIPT_VAI                   Script = 74
	PANGO_SCRIPT_CARIAN                Script = 75
	PANGO_SCRIPT_LYCIAN                Script = 76
	PANGO_SCRIPT_LYDIAN                Script = 77
	PANGO_SCRIPT_BATAK                 Script = 78
	PANGO_SCRIPT_BRAHMI                Script = 79
	PANGO_SCRIPT_MANDAIC               Script = 80
	PANGO_SCRIPT_CHAKMA                Script = 81
	PANGO_SCRIPT_MEROITIC_CURSIVE      Script = 82
	PANGO_SCRIPT_MEROITIC_HIEROGLYPHS  Script = 83
	PANGO_SCRIPT_MIAO                  Script = 84
	PANGO_SCRIPT_SHARADA               Script = 85
	PANGO_SCRIPT_SORA_SOMPENG          Script = 86
	PANGO_SCRIPT_TAKRI                 Script = 87
	PANGO_SCRIPT_BASSA_VAH             Script = 88
	PANGO_SCRIPT_CAUCASIAN_ALBANIAN    Script = 89
	PANGO_SCRIPT_DUPLOYAN              Script = 90
	PANGO_SCRIPT_ELBASAN               Script = 91
	PANGO_SCRIPT_GRANTHA               Script = 92
	PANGO_SCRIPT_KHOJKI                Script = 93
	PANGO_SCRIPT_KHUDAWADI             Script = 94
	PANGO_SCRIPT_LINEAR_A              Script = 95
	PANGO_SCRIPT_MAHAJANI              Script = 96
	PANGO_SCRIPT_MANICHAEAN            Script = 97
	PANGO_SCRIPT_MENDE_KIKAKUI         Script = 98
	PANGO_SCRIPT_MODI                  Script = 99
	PANGO_SCRIPT_MRO                   Script = 100
	PANGO_SCRIPT_NABATAEAN             Script = 101
	PANGO_SCRIPT_OLD_NORTH_ARABIAN     Script = 102
	PANGO_SCRIPT_OLD_PERMIC            Script = 103
	PANGO_SCRIPT_PAHAWH_HMONG          Script = 104
	PANGO_SCRIPT_PALMYRENE             Script = 105
	PANGO_SCRIPT_PAU_CIN_HAU           Script = 106
	PANGO_SCRIPT_PSALTER_PAHLAVI       Script = 107
	PANGO_SCRIPT_SIDDHAM               Script = 108
	PANGO_SCRIPT_TIRHUTA               Script = 109
	PANGO_SCRIPT_WARANG_CITI           Script = 110
	PANGO_SCRIPT_AHOM                  Script = 111
	PANGO_SCRIPT_ANATOLIAN_HIEROGLYPHS Script = 112
	PANGO_SCRIPT_HATRAN                Script = 113
	PANGO_SCRIPT_MULTANI               Script = 114
	PANGO_SCRIPT_OLD_HUNGARIAN         Script = 115
	PANGO_SCRIPT_SIGNWRITING           Script = 116
)

// ScriptForUnichar is a wrapper around the C function pango_script_for_unichar.
func ScriptForUnichar(ch rune) Script {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_script_for_unichar(c_ch)
	retGo := (Script)(retC)

	return retGo
}

// ScriptGetSampleLanguage is a wrapper around the C function pango_script_get_sample_language.
func ScriptGetSampleLanguage(script Script) *Language {
	c_script := (C.PangoScript)(script)

	retC := C.pango_script_get_sample_language(c_script)
	var retGo (*Language)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LanguageNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

type Stretch C.PangoStretch

const (
	PANGO_STRETCH_ULTRA_CONDENSED Stretch = 0
	PANGO_STRETCH_EXTRA_CONDENSED Stretch = 1
	PANGO_STRETCH_CONDENSED       Stretch = 2
	PANGO_STRETCH_SEMI_CONDENSED  Stretch = 3
	PANGO_STRETCH_NORMAL          Stretch = 4
	PANGO_STRETCH_SEMI_EXPANDED   Stretch = 5
	PANGO_STRETCH_EXPANDED        Stretch = 6
	PANGO_STRETCH_EXTRA_EXPANDED  Stretch = 7
	PANGO_STRETCH_ULTRA_EXPANDED  Stretch = 8
)

type Style C.PangoStyle

const (
	PANGO_STYLE_NORMAL  Style = 0
	PANGO_STYLE_OBLIQUE Style = 1
	PANGO_STYLE_ITALIC  Style = 2
)

type TabAlign C.PangoTabAlign

const (
	PANGO_TAB_LEFT TabAlign = 0
)

type Underline C.PangoUnderline

const (
	PANGO_UNDERLINE_NONE   Underline = 0
	PANGO_UNDERLINE_SINGLE Underline = 1
	PANGO_UNDERLINE_DOUBLE Underline = 2
	PANGO_UNDERLINE_LOW    Underline = 3
	PANGO_UNDERLINE_ERROR  Underline = 4
)

type Variant C.PangoVariant

const (
	PANGO_VARIANT_NORMAL     Variant = 0
	PANGO_VARIANT_SMALL_CAPS Variant = 1
)

type Weight C.PangoWeight

const (
	PANGO_WEIGHT_THIN       Weight = 100
	PANGO_WEIGHT_ULTRALIGHT Weight = 200
	PANGO_WEIGHT_LIGHT      Weight = 300
	PANGO_WEIGHT_SEMILIGHT  Weight = 350
	PANGO_WEIGHT_BOOK       Weight = 380
	PANGO_WEIGHT_NORMAL     Weight = 400
	PANGO_WEIGHT_MEDIUM     Weight = 500
	PANGO_WEIGHT_SEMIBOLD   Weight = 600
	PANGO_WEIGHT_BOLD       Weight = 700
	PANGO_WEIGHT_ULTRABOLD  Weight = 800
	PANGO_WEIGHT_HEAVY      Weight = 900
	PANGO_WEIGHT_ULTRAHEAVY Weight = 1000
)

type WrapMode C.PangoWrapMode

const (
	PANGO_WRAP_WORD      WrapMode = 0
	PANGO_WRAP_CHAR      WrapMode = 1
	PANGO_WRAP_WORD_CHAR WrapMode = 2
)

// AttrBackgroundNew is a wrapper around the C function pango_attr_background_new.
func AttrBackgroundNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_background_new(c_red, c_green, c_blue)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrFallbackNew is a wrapper around the C function pango_attr_fallback_new.
func AttrFallbackNew(enableFallback bool) *Attribute {
	c_enable_fallback :=
		boolToGboolean(enableFallback)

	retC := C.pango_attr_fallback_new(c_enable_fallback)
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

// AttrGravityHintNew is a wrapper around the C function pango_attr_gravity_hint_new.
func AttrGravityHintNew(hint GravityHint) *Attribute {
	c_hint := (C.PangoGravityHint)(hint)

	retC := C.pango_attr_gravity_hint_new(c_hint)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrGravityNew is a wrapper around the C function pango_attr_gravity_new.
func AttrGravityNew(gravity Gravity) *Attribute {
	c_gravity := (C.PangoGravity)(gravity)

	retC := C.pango_attr_gravity_new(c_gravity)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrLetterSpacingNew is a wrapper around the C function pango_attr_letter_spacing_new.
func AttrLetterSpacingNew(letterSpacing int32) *Attribute {
	c_letter_spacing := (C.int)(letterSpacing)

	retC := C.pango_attr_letter_spacing_new(c_letter_spacing)
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

// AttrStrikethroughColorNew is a wrapper around the C function pango_attr_strikethrough_color_new.
func AttrStrikethroughColorNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_strikethrough_color_new(c_red, c_green, c_blue)
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

// AttrUnderlineColorNew is a wrapper around the C function pango_attr_underline_color_new.
func AttrUnderlineColorNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_underline_color_new(c_red, c_green, c_blue)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

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

// Unsupported : pango_break : unsupported parameter attrs :

// Blacklisted : pango_config_key_get

// Blacklisted : pango_config_key_get_system

// Blacklisted : pango_default_break

// ExtentsToPixels is a wrapper around the C function pango_extents_to_pixels.
func ExtentsToPixels(inclusive *Rectangle, nearest *Rectangle) {
	c_inclusive := (*C.PangoRectangle)(C.NULL)
	if inclusive != nil {
		c_inclusive = (*C.PangoRectangle)(inclusive.ToC())
	}

	c_nearest := (*C.PangoRectangle)(C.NULL)
	if nearest != nil {
		c_nearest = (*C.PangoRectangle)(nearest.ToC())
	}

	C.pango_extents_to_pixels(c_inclusive, c_nearest)

	return
}

// FindBaseDir is a wrapper around the C function pango_find_base_dir.
func FindBaseDir(text string, length int32) Direction {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gint)(length)

	retC := C.pango_find_base_dir(c_text, c_length)
	retGo := (Direction)(retC)

	return retGo
}

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

// Blacklisted : pango_get_lib_subdirectory

// Unsupported : pango_get_log_attrs : unsupported parameter log_attrs :

// GetMirrorChar is a wrapper around the C function pango_get_mirror_char.
func GetMirrorChar(ch rune, mirroredCh rune) bool {
	c_ch := (C.gunichar)(ch)

	c_mirrored_ch := (C.gunichar)(mirroredCh)

	retC := C.pango_get_mirror_char(c_ch, &c_mirrored_ch)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : pango_get_sysconf_subdirectory

// IsZeroWidth is a wrapper around the C function pango_is_zero_width.
func IsZeroWidth(ch rune) bool {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_is_zero_width(c_ch)
	retGo := retC == C.TRUE

	return retGo
}

// Itemize is a wrapper around the C function pango_itemize.
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

// ItemizeWithBaseDir is a wrapper around the C function pango_itemize_with_base_dir.
func ItemizeWithBaseDir(context *Context, baseDir Direction, text string, startIndex int32, length int32, attrs *AttrList, cachedIter *AttrIterator) *glib.List {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	c_base_dir := (C.PangoDirection)(baseDir)

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

	retC := C.pango_itemize_with_base_dir(c_context, c_base_dir, c_text, c_start_index, c_length, c_attrs, c_cached_iter)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : pango_log2vis_get_embedding_levels

// Unsupported : pango_lookup_aliases : unsupported parameter families : output array param families

// Unsupported : pango_module_register : unsupported parameter module : Blacklisted record : PangoIncludedModule

// ParseEnum is a wrapper around the C function pango_parse_enum.
func ParseEnum(type_ gobject.Type, str string, warn bool) (bool, int32, string) {
	c_type := (C.GType)(type_)

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	var c_value C.int

	c_warn :=
		boolToGboolean(warn)

	var c_possible_values *C.char

	retC := C.pango_parse_enum(c_type, c_str, &c_value, c_warn, &c_possible_values)
	retGo := retC == C.TRUE

	value := (int32)(c_value)

	possibleValues := C.GoString(c_possible_values)
	defer C.free(unsafe.Pointer(c_possible_values))

	return retGo, value, possibleValues
}

// ParseMarkup is a wrapper around the C function pango_parse_markup.
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	attrList := AttrListNewFromC(unsafe.Pointer(c_attr_list))

	text := C.GoString(c_text)
	defer C.free(unsafe.Pointer(c_text))

	accelChar := (rune)(c_accel_char)

	return retGo, attrList, text, accelChar, goError
}

// ParseStretch is a wrapper around the C function pango_parse_stretch.
func ParseStretch(str string, warn bool) (bool, Stretch) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	var c_stretch C.PangoStretch

	c_warn :=
		boolToGboolean(warn)

	retC := C.pango_parse_stretch(c_str, &c_stretch, c_warn)
	retGo := retC == C.TRUE

	stretch := (Stretch)(c_stretch)

	return retGo, stretch
}

// ParseStyle is a wrapper around the C function pango_parse_style.
func ParseStyle(str string, warn bool) (bool, Style) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	var c_style C.PangoStyle

	c_warn :=
		boolToGboolean(warn)

	retC := C.pango_parse_style(c_str, &c_style, c_warn)
	retGo := retC == C.TRUE

	style := (Style)(c_style)

	return retGo, style
}

// ParseVariant is a wrapper around the C function pango_parse_variant.
func ParseVariant(str string, warn bool) (bool, Variant) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	var c_variant C.PangoVariant

	c_warn :=
		boolToGboolean(warn)

	retC := C.pango_parse_variant(c_str, &c_variant, c_warn)
	retGo := retC == C.TRUE

	variant := (Variant)(c_variant)

	return retGo, variant
}

// ParseWeight is a wrapper around the C function pango_parse_weight.
func ParseWeight(str string, warn bool) (bool, Weight) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	var c_weight C.PangoWeight

	c_warn :=
		boolToGboolean(warn)

	retC := C.pango_parse_weight(c_str, &c_weight, c_warn)
	retGo := retC == C.TRUE

	weight := (Weight)(c_weight)

	return retGo, weight
}

// QuantizeLineGeometry is a wrapper around the C function pango_quantize_line_geometry.
func QuantizeLineGeometry(thickness int32, position int32) {
	c_thickness := (C.int)(thickness)

	c_position := (C.int)(position)

	C.pango_quantize_line_geometry(&c_thickness, &c_position)

	return
}

// Blacklisted : pango_read_line

// ReorderItems is a wrapper around the C function pango_reorder_items.
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

// SplitFileList is a wrapper around the C function pango_split_file_list.
func SplitFileList(str string) []string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.pango_split_file_list(c_str)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

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

// UnitsFromDouble is a wrapper around the C function pango_units_from_double.
func UnitsFromDouble(d float64) int32 {
	c_d := (C.double)(d)

	retC := C.pango_units_from_double(c_d)
	retGo := (int32)(retC)

	return retGo
}

// UnitsToDouble is a wrapper around the C function pango_units_to_double.
func UnitsToDouble(i int32) float64 {
	c_i := (C.int)(i)

	retC := C.pango_units_to_double(c_i)
	retGo := (float64)(retC)

	return retGo
}

// Version is a wrapper around the C function pango_version.
func Version() int32 {
	retC := C.pango_version()
	retGo := (int32)(retC)

	return retGo
}

// VersionCheck is a wrapper around the C function pango_version_check.
func VersionCheck(requiredMajor int32, requiredMinor int32, requiredMicro int32) string {
	c_required_major := (C.int)(requiredMajor)

	c_required_minor := (C.int)(requiredMinor)

	c_required_micro := (C.int)(requiredMicro)

	retC := C.pango_version_check(c_required_major, c_required_minor, c_required_micro)
	retGo := C.GoString(retC)

	return retGo
}

// VersionString is a wrapper around the C function pango_version_string.
func VersionString() string {
	retC := C.pango_version_string()
	retGo := C.GoString(retC)

	return retGo
}

// Analysis is a wrapper around the C record PangoAnalysis.
type Analysis struct {
	native *C.PangoAnalysis
	// shape_engine : record
	// lang_engine : record
	// font : record
	Level   uint8
	Gravity uint8
	Flags   uint8
	Script  uint8
	// language : record
	// extra_attrs : record
}

func AnalysisNewFromC(u unsafe.Pointer) *Analysis {
	c := (*C.PangoAnalysis)(u)
	if c == nil {
		return nil
	}

	g := &Analysis{
		Flags:   (uint8)(c.flags),
		Gravity: (uint8)(c.gravity),
		Level:   (uint8)(c.level),
		Script:  (uint8)(c.script),
		native:  c,
	}

	return g
}

func (recv *Analysis) ToC() unsafe.Pointer {
	recv.native.level =
		(C.guint8)(recv.Level)
	recv.native.gravity =
		(C.guint8)(recv.Gravity)
	recv.native.flags =
		(C.guint8)(recv.Flags)
	recv.native.script =
		(C.guint8)(recv.Script)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Analysis with another Analysis, and returns true if they represent the same GObject.
func (recv *Analysis) Equals(other *Analysis) bool {
	return other.ToC() == recv.ToC()
}

// AttrClass is a wrapper around the C record PangoAttrClass.
type AttrClass struct {
	native *C.PangoAttrClass
	Type   AttrType
	// no type for copy
	// no type for destroy
	// no type for equal
}

func AttrClassNewFromC(u unsafe.Pointer) *AttrClass {
	c := (*C.PangoAttrClass)(u)
	if c == nil {
		return nil
	}

	g := &AttrClass{
		Type:   (AttrType)(c._type),
		native: c,
	}

	return g
}

func (recv *AttrClass) ToC() unsafe.Pointer {
	recv.native._type =
		(C.PangoAttrType)(recv.Type)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrClass with another AttrClass, and returns true if they represent the same GObject.
func (recv *AttrClass) Equals(other *AttrClass) bool {
	return other.ToC() == recv.ToC()
}

// AttrColor is a wrapper around the C record PangoAttrColor.
type AttrColor struct {
	native *C.PangoAttrColor
	// attr : record
	// color : record
}

func AttrColorNewFromC(u unsafe.Pointer) *AttrColor {
	c := (*C.PangoAttrColor)(u)
	if c == nil {
		return nil
	}

	g := &AttrColor{native: c}

	return g
}

func (recv *AttrColor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrColor with another AttrColor, and returns true if they represent the same GObject.
func (recv *AttrColor) Equals(other *AttrColor) bool {
	return other.ToC() == recv.ToC()
}

// AttrFloat is a wrapper around the C record PangoAttrFloat.
type AttrFloat struct {
	native *C.PangoAttrFloat
	// attr : record
	Value float64
}

func AttrFloatNewFromC(u unsafe.Pointer) *AttrFloat {
	c := (*C.PangoAttrFloat)(u)
	if c == nil {
		return nil
	}

	g := &AttrFloat{
		Value:  (float64)(c.value),
		native: c,
	}

	return g
}

func (recv *AttrFloat) ToC() unsafe.Pointer {
	recv.native.value =
		(C.double)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrFloat with another AttrFloat, and returns true if they represent the same GObject.
func (recv *AttrFloat) Equals(other *AttrFloat) bool {
	return other.ToC() == recv.ToC()
}

// AttrFontDesc is a wrapper around the C record PangoAttrFontDesc.
type AttrFontDesc struct {
	native *C.PangoAttrFontDesc
	// attr : record
	// desc : record
}

func AttrFontDescNewFromC(u unsafe.Pointer) *AttrFontDesc {
	c := (*C.PangoAttrFontDesc)(u)
	if c == nil {
		return nil
	}

	g := &AttrFontDesc{native: c}

	return g
}

func (recv *AttrFontDesc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrFontDesc with another AttrFontDesc, and returns true if they represent the same GObject.
func (recv *AttrFontDesc) Equals(other *AttrFontDesc) bool {
	return other.ToC() == recv.ToC()
}

// AttrFontDescNew is a wrapper around the C function pango_attr_font_desc_new.
func AttrFontDescNew(desc *FontDescription) *Attribute {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	retC := C.pango_attr_font_desc_new(c_desc)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrInt is a wrapper around the C record PangoAttrInt.
type AttrInt struct {
	native *C.PangoAttrInt
	// attr : record
	Value int32
}

func AttrIntNewFromC(u unsafe.Pointer) *AttrInt {
	c := (*C.PangoAttrInt)(u)
	if c == nil {
		return nil
	}

	g := &AttrInt{
		Value:  (int32)(c.value),
		native: c,
	}

	return g
}

func (recv *AttrInt) ToC() unsafe.Pointer {
	recv.native.value =
		(C.int)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrInt with another AttrInt, and returns true if they represent the same GObject.
func (recv *AttrInt) Equals(other *AttrInt) bool {
	return other.ToC() == recv.ToC()
}

// AttrIterator is a wrapper around the C record PangoAttrIterator.
type AttrIterator struct {
	native *C.PangoAttrIterator
}

func AttrIteratorNewFromC(u unsafe.Pointer) *AttrIterator {
	c := (*C.PangoAttrIterator)(u)
	if c == nil {
		return nil
	}

	g := &AttrIterator{native: c}

	return g
}

func (recv *AttrIterator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrIterator with another AttrIterator, and returns true if they represent the same GObject.
func (recv *AttrIterator) Equals(other *AttrIterator) bool {
	return other.ToC() == recv.ToC()
}

// Copy is a wrapper around the C function pango_attr_iterator_copy.
func (recv *AttrIterator) Copy() *AttrIterator {
	retC := C.pango_attr_iterator_copy((*C.PangoAttrIterator)(recv.native))
	retGo := AttrIteratorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Destroy is a wrapper around the C function pango_attr_iterator_destroy.
func (recv *AttrIterator) Destroy() {
	C.pango_attr_iterator_destroy((*C.PangoAttrIterator)(recv.native))

	return
}

// Get is a wrapper around the C function pango_attr_iterator_get.
func (recv *AttrIterator) Get(type_ AttrType) *Attribute {
	c_type := (C.PangoAttrType)(type_)

	retC := C.pango_attr_iterator_get((*C.PangoAttrIterator)(recv.native), c_type)
	var retGo (*Attribute)
	if retC == nil {
		retGo = nil
	} else {
		retGo = AttributeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetAttrs is a wrapper around the C function pango_attr_iterator_get_attrs.
func (recv *AttrIterator) GetAttrs() *glib.SList {
	retC := C.pango_attr_iterator_get_attrs((*C.PangoAttrIterator)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFont is a wrapper around the C function pango_attr_iterator_get_font.
func (recv *AttrIterator) GetFont(desc *FontDescription, language *Language, extraAttrs *glib.SList) {
	c_desc := (*C.PangoFontDescription)(C.NULL)
	if desc != nil {
		c_desc = (*C.PangoFontDescription)(desc.ToC())
	}

	c_language := (**C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (**C.PangoLanguage)(language.ToC())
	}

	c_extra_attrs := (**C.GSList)(C.NULL)
	if extraAttrs != nil {
		c_extra_attrs = (**C.GSList)(extraAttrs.ToC())
	}

	C.pango_attr_iterator_get_font((*C.PangoAttrIterator)(recv.native), c_desc, c_language, c_extra_attrs)

	return
}

// Next is a wrapper around the C function pango_attr_iterator_next.
func (recv *AttrIterator) Next() bool {
	retC := C.pango_attr_iterator_next((*C.PangoAttrIterator)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Range is a wrapper around the C function pango_attr_iterator_range.
func (recv *AttrIterator) Range() (int32, int32) {
	var c_start C.gint

	var c_end C.gint

	C.pango_attr_iterator_range((*C.PangoAttrIterator)(recv.native), &c_start, &c_end)

	start := (int32)(c_start)

	end := (int32)(c_end)

	return start, end
}

// AttrLanguage is a wrapper around the C record PangoAttrLanguage.
type AttrLanguage struct {
	native *C.PangoAttrLanguage
	// attr : record
	// value : record
}

func AttrLanguageNewFromC(u unsafe.Pointer) *AttrLanguage {
	c := (*C.PangoAttrLanguage)(u)
	if c == nil {
		return nil
	}

	g := &AttrLanguage{native: c}

	return g
}

func (recv *AttrLanguage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrLanguage with another AttrLanguage, and returns true if they represent the same GObject.
func (recv *AttrLanguage) Equals(other *AttrLanguage) bool {
	return other.ToC() == recv.ToC()
}

// AttrLanguageNew is a wrapper around the C function pango_attr_language_new.
func AttrLanguageNew(language *Language) *Attribute {
	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	retC := C.pango_attr_language_new(c_language)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrList is a wrapper around the C record PangoAttrList.
type AttrList struct {
	native *C.PangoAttrList
}

func AttrListNewFromC(u unsafe.Pointer) *AttrList {
	c := (*C.PangoAttrList)(u)
	if c == nil {
		return nil
	}

	g := &AttrList{native: c}

	return g
}

func (recv *AttrList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrList with another AttrList, and returns true if they represent the same GObject.
func (recv *AttrList) Equals(other *AttrList) bool {
	return other.ToC() == recv.ToC()
}

// AttrListNew is a wrapper around the C function pango_attr_list_new.
func AttrListNew() *AttrList {
	retC := C.pango_attr_list_new()
	retGo := AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Change is a wrapper around the C function pango_attr_list_change.
func (recv *AttrList) Change(attr *Attribute) {
	c_attr := (*C.PangoAttribute)(C.NULL)
	if attr != nil {
		c_attr = (*C.PangoAttribute)(attr.ToC())
	}

	C.pango_attr_list_change((*C.PangoAttrList)(recv.native), c_attr)

	return
}

// Copy is a wrapper around the C function pango_attr_list_copy.
func (recv *AttrList) Copy() *AttrList {
	retC := C.pango_attr_list_copy((*C.PangoAttrList)(recv.native))
	var retGo (*AttrList)
	if retC == nil {
		retGo = nil
	} else {
		retGo = AttrListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : pango_attr_list_filter : unsupported parameter func : no type generator for AttrFilterFunc (PangoAttrFilterFunc) for param func

// GetIterator is a wrapper around the C function pango_attr_list_get_iterator.
func (recv *AttrList) GetIterator() *AttrIterator {
	retC := C.pango_attr_list_get_iterator((*C.PangoAttrList)(recv.native))
	retGo := AttrIteratorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function pango_attr_list_insert.
func (recv *AttrList) Insert(attr *Attribute) {
	c_attr := (*C.PangoAttribute)(C.NULL)
	if attr != nil {
		c_attr = (*C.PangoAttribute)(attr.ToC())
	}

	C.pango_attr_list_insert((*C.PangoAttrList)(recv.native), c_attr)

	return
}

// InsertBefore is a wrapper around the C function pango_attr_list_insert_before.
func (recv *AttrList) InsertBefore(attr *Attribute) {
	c_attr := (*C.PangoAttribute)(C.NULL)
	if attr != nil {
		c_attr = (*C.PangoAttribute)(attr.ToC())
	}

	C.pango_attr_list_insert_before((*C.PangoAttrList)(recv.native), c_attr)

	return
}

// Ref is a wrapper around the C function pango_attr_list_ref.
func (recv *AttrList) Ref() *AttrList {
	retC := C.pango_attr_list_ref((*C.PangoAttrList)(recv.native))
	retGo := AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Splice is a wrapper around the C function pango_attr_list_splice.
func (recv *AttrList) Splice(other *AttrList, pos int32, len int32) {
	c_other := (*C.PangoAttrList)(C.NULL)
	if other != nil {
		c_other = (*C.PangoAttrList)(other.ToC())
	}

	c_pos := (C.gint)(pos)

	c_len := (C.gint)(len)

	C.pango_attr_list_splice((*C.PangoAttrList)(recv.native), c_other, c_pos, c_len)

	return
}

// Unref is a wrapper around the C function pango_attr_list_unref.
func (recv *AttrList) Unref() {
	C.pango_attr_list_unref((*C.PangoAttrList)(recv.native))

	return
}

// AttrShape is a wrapper around the C record PangoAttrShape.
type AttrShape struct {
	native *C.PangoAttrShape
	// attr : record
	// ink_rect : record
	// logical_rect : record
	Data uintptr
	// copy_func : no type generator for AttrDataCopyFunc, PangoAttrDataCopyFunc
	// destroy_func : no type generator for GLib.DestroyNotify, GDestroyNotify
}

func AttrShapeNewFromC(u unsafe.Pointer) *AttrShape {
	c := (*C.PangoAttrShape)(u)
	if c == nil {
		return nil
	}

	g := &AttrShape{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *AttrShape) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrShape with another AttrShape, and returns true if they represent the same GObject.
func (recv *AttrShape) Equals(other *AttrShape) bool {
	return other.ToC() == recv.ToC()
}

// AttrShapeNew is a wrapper around the C function pango_attr_shape_new.
func AttrShapeNew(inkRect *Rectangle, logicalRect *Rectangle) *Attribute {
	c_ink_rect := (*C.PangoRectangle)(C.NULL)
	if inkRect != nil {
		c_ink_rect = (*C.PangoRectangle)(inkRect.ToC())
	}

	c_logical_rect := (*C.PangoRectangle)(C.NULL)
	if logicalRect != nil {
		c_logical_rect = (*C.PangoRectangle)(logicalRect.ToC())
	}

	retC := C.pango_attr_shape_new(c_ink_rect, c_logical_rect)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// pango_attr_shape_new_with_data : unsupported parameter copy_func : no type generator for AttrDataCopyFunc (PangoAttrDataCopyFunc) for param copy_func
// AttrSize is a wrapper around the C record PangoAttrSize.
type AttrSize struct {
	native *C.PangoAttrSize
	// attr : record
	Size int32
	// Bitfield not supported :  1 absolute
}

func AttrSizeNewFromC(u unsafe.Pointer) *AttrSize {
	c := (*C.PangoAttrSize)(u)
	if c == nil {
		return nil
	}

	g := &AttrSize{
		Size:   (int32)(c.size),
		native: c,
	}

	return g
}

func (recv *AttrSize) ToC() unsafe.Pointer {
	recv.native.size =
		(C.int)(recv.Size)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrSize with another AttrSize, and returns true if they represent the same GObject.
func (recv *AttrSize) Equals(other *AttrSize) bool {
	return other.ToC() == recv.ToC()
}

// AttrSizeNew is a wrapper around the C function pango_attr_size_new.
func AttrSizeNew(size int32) *Attribute {
	c_size := (C.int)(size)

	retC := C.pango_attr_size_new(c_size)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrSizeNewAbsolute is a wrapper around the C function pango_attr_size_new_absolute.
func AttrSizeNewAbsolute(size int32) *Attribute {
	c_size := (C.int)(size)

	retC := C.pango_attr_size_new_absolute(c_size)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrString is a wrapper around the C record PangoAttrString.
type AttrString struct {
	native *C.PangoAttrString
	// attr : record
	Value string
}

func AttrStringNewFromC(u unsafe.Pointer) *AttrString {
	c := (*C.PangoAttrString)(u)
	if c == nil {
		return nil
	}

	g := &AttrString{
		Value:  C.GoString(c.value),
		native: c,
	}

	return g
}

func (recv *AttrString) ToC() unsafe.Pointer {
	recv.native.value =
		C.CString(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttrString with another AttrString, and returns true if they represent the same GObject.
func (recv *AttrString) Equals(other *AttrString) bool {
	return other.ToC() == recv.ToC()
}

// Attribute is a wrapper around the C record PangoAttribute.
type Attribute struct {
	native *C.PangoAttribute
	// klass : record
	StartIndex uint32
	EndIndex   uint32
}

func AttributeNewFromC(u unsafe.Pointer) *Attribute {
	c := (*C.PangoAttribute)(u)
	if c == nil {
		return nil
	}

	g := &Attribute{
		EndIndex:   (uint32)(c.end_index),
		StartIndex: (uint32)(c.start_index),
		native:     c,
	}

	return g
}

func (recv *Attribute) ToC() unsafe.Pointer {
	recv.native.start_index =
		(C.guint)(recv.StartIndex)
	recv.native.end_index =
		(C.guint)(recv.EndIndex)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Attribute with another Attribute, and returns true if they represent the same GObject.
func (recv *Attribute) Equals(other *Attribute) bool {
	return other.ToC() == recv.ToC()
}

// Copy is a wrapper around the C function pango_attribute_copy.
func (recv *Attribute) Copy() *Attribute {
	retC := C.pango_attribute_copy((*C.PangoAttribute)(recv.native))
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Destroy is a wrapper around the C function pango_attribute_destroy.
func (recv *Attribute) Destroy() {
	C.pango_attribute_destroy((*C.PangoAttribute)(recv.native))

	return
}

// Equal is a wrapper around the C function pango_attribute_equal.
func (recv *Attribute) Equal(attr2 *Attribute) bool {
	c_attr2 := (*C.PangoAttribute)(C.NULL)
	if attr2 != nil {
		c_attr2 = (*C.PangoAttribute)(attr2.ToC())
	}

	retC := C.pango_attribute_equal((*C.PangoAttribute)(recv.native), c_attr2)
	retGo := retC == C.TRUE

	return retGo
}

// Color is a wrapper around the C record PangoColor.
type Color struct {
	native *C.PangoColor
	Red    uint16
	Green  uint16
	Blue   uint16
}

func ColorNewFromC(u unsafe.Pointer) *Color {
	c := (*C.PangoColor)(u)
	if c == nil {
		return nil
	}

	g := &Color{
		Blue:   (uint16)(c.blue),
		Green:  (uint16)(c.green),
		Red:    (uint16)(c.red),
		native: c,
	}

	return g
}

func (recv *Color) ToC() unsafe.Pointer {
	recv.native.red =
		(C.guint16)(recv.Red)
	recv.native.green =
		(C.guint16)(recv.Green)
	recv.native.blue =
		(C.guint16)(recv.Blue)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Color with another Color, and returns true if they represent the same GObject.
func (recv *Color) Equals(other *Color) bool {
	return other.ToC() == recv.ToC()
}

// Copy is a wrapper around the C function pango_color_copy.
func (recv *Color) Copy() *Color {
	retC := C.pango_color_copy((*C.PangoColor)(recv.native))
	var retGo (*Color)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ColorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Free is a wrapper around the C function pango_color_free.
func (recv *Color) Free() {
	C.pango_color_free((*C.PangoColor)(recv.native))

	return
}

// Parse is a wrapper around the C function pango_color_parse.
func (recv *Color) Parse(spec string) bool {
	c_spec := C.CString(spec)
	defer C.free(unsafe.Pointer(c_spec))

	retC := C.pango_color_parse((*C.PangoColor)(recv.native), c_spec)
	retGo := retC == C.TRUE

	return retGo
}

// ToString is a wrapper around the C function pango_color_to_string.
func (recv *Color) ToString() string {
	retC := C.pango_color_to_string((*C.PangoColor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ContextClass is a wrapper around the C record PangoContextClass.
type ContextClass struct {
	native *C.PangoContextClass
}

func ContextClassNewFromC(u unsafe.Pointer) *ContextClass {
	c := (*C.PangoContextClass)(u)
	if c == nil {
		return nil
	}

	g := &ContextClass{native: c}

	return g
}

func (recv *ContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContextClass with another ContextClass, and returns true if they represent the same GObject.
func (recv *ContextClass) Equals(other *ContextClass) bool {
	return other.ToC() == recv.ToC()
}

// Coverage is a wrapper around the C record PangoCoverage.
type Coverage struct {
	native *C.PangoCoverage
}

func CoverageNewFromC(u unsafe.Pointer) *Coverage {
	c := (*C.PangoCoverage)(u)
	if c == nil {
		return nil
	}

	g := &Coverage{native: c}

	return g
}

func (recv *Coverage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Coverage with another Coverage, and returns true if they represent the same GObject.
func (recv *Coverage) Equals(other *Coverage) bool {
	return other.ToC() == recv.ToC()
}

// CoverageFromBytes is a wrapper around the C function pango_coverage_from_bytes.
func CoverageFromBytes(bytes []uint8) *Coverage {
	c_bytes_array := make([]C.guint8, len(bytes)+1, len(bytes)+1)
	for i, item := range bytes {
		c := (C.guint8)(item)
		c_bytes_array[i] = c
	}
	c_bytes_array[len(bytes)] = 0
	c_bytes_arrayPtr := &c_bytes_array[0]
	c_bytes := (*C.guchar)(unsafe.Pointer(c_bytes_arrayPtr))

	c_n_bytes := (C.int)(len(bytes))

	retC := C.pango_coverage_from_bytes(c_bytes, c_n_bytes)
	var retGo (*Coverage)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CoverageNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// CoverageNew is a wrapper around the C function pango_coverage_new.
func CoverageNew() *Coverage {
	retC := C.pango_coverage_new()
	retGo := CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function pango_coverage_copy.
func (recv *Coverage) Copy() *Coverage {
	retC := C.pango_coverage_copy((*C.PangoCoverage)(recv.native))
	retGo := CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Get is a wrapper around the C function pango_coverage_get.
func (recv *Coverage) Get(index int32) CoverageLevel {
	c_index_ := (C.int)(index)

	retC := C.pango_coverage_get((*C.PangoCoverage)(recv.native), c_index_)
	retGo := (CoverageLevel)(retC)

	return retGo
}

// Max is a wrapper around the C function pango_coverage_max.
func (recv *Coverage) Max(other *Coverage) {
	c_other := (*C.PangoCoverage)(C.NULL)
	if other != nil {
		c_other = (*C.PangoCoverage)(other.ToC())
	}

	C.pango_coverage_max((*C.PangoCoverage)(recv.native), c_other)

	return
}

// Ref is a wrapper around the C function pango_coverage_ref.
func (recv *Coverage) Ref() *Coverage {
	retC := C.pango_coverage_ref((*C.PangoCoverage)(recv.native))
	retGo := CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Set is a wrapper around the C function pango_coverage_set.
func (recv *Coverage) Set(index int32, level CoverageLevel) {
	c_index_ := (C.int)(index)

	c_level := (C.PangoCoverageLevel)(level)

	C.pango_coverage_set((*C.PangoCoverage)(recv.native), c_index_, c_level)

	return
}

// Unsupported : pango_coverage_to_bytes : unsupported parameter bytes : output array param bytes

// Unref is a wrapper around the C function pango_coverage_unref.
func (recv *Coverage) Unref() {
	C.pango_coverage_unref((*C.PangoCoverage)(recv.native))

	return
}

// Blacklisted : PangoEngineClass

// Blacklisted : PangoEngineInfo

// Blacklisted : PangoEngineLangClass

// Blacklisted : PangoEngineScriptInfo

// Blacklisted : PangoEngineShapeClass

// Blacklisted : PangoFontClass

// FontDescription is a wrapper around the C record PangoFontDescription.
type FontDescription struct {
	native *C.PangoFontDescription
}

func FontDescriptionNewFromC(u unsafe.Pointer) *FontDescription {
	c := (*C.PangoFontDescription)(u)
	if c == nil {
		return nil
	}

	g := &FontDescription{native: c}

	return g
}

func (recv *FontDescription) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FontDescription with another FontDescription, and returns true if they represent the same GObject.
func (recv *FontDescription) Equals(other *FontDescription) bool {
	return other.ToC() == recv.ToC()
}

// FontDescriptionNew is a wrapper around the C function pango_font_description_new.
func FontDescriptionNew() *FontDescription {
	retC := C.pango_font_description_new()
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontDescriptionFromString is a wrapper around the C function pango_font_description_from_string.
func FontDescriptionFromString(str string) *FontDescription {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.pango_font_description_from_string(c_str)
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BetterMatch is a wrapper around the C function pango_font_description_better_match.
func (recv *FontDescription) BetterMatch(oldMatch *FontDescription, newMatch *FontDescription) bool {
	c_old_match := (*C.PangoFontDescription)(C.NULL)
	if oldMatch != nil {
		c_old_match = (*C.PangoFontDescription)(oldMatch.ToC())
	}

	c_new_match := (*C.PangoFontDescription)(C.NULL)
	if newMatch != nil {
		c_new_match = (*C.PangoFontDescription)(newMatch.ToC())
	}

	retC := C.pango_font_description_better_match((*C.PangoFontDescription)(recv.native), c_old_match, c_new_match)
	retGo := retC == C.TRUE

	return retGo
}

// Copy is a wrapper around the C function pango_font_description_copy.
func (recv *FontDescription) Copy() *FontDescription {
	retC := C.pango_font_description_copy((*C.PangoFontDescription)(recv.native))
	var retGo (*FontDescription)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontDescriptionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// CopyStatic is a wrapper around the C function pango_font_description_copy_static.
func (recv *FontDescription) CopyStatic() *FontDescription {
	retC := C.pango_font_description_copy_static((*C.PangoFontDescription)(recv.native))
	var retGo (*FontDescription)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontDescriptionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Equal is a wrapper around the C function pango_font_description_equal.
func (recv *FontDescription) Equal(desc2 *FontDescription) bool {
	c_desc2 := (*C.PangoFontDescription)(C.NULL)
	if desc2 != nil {
		c_desc2 = (*C.PangoFontDescription)(desc2.ToC())
	}

	retC := C.pango_font_description_equal((*C.PangoFontDescription)(recv.native), c_desc2)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function pango_font_description_free.
func (recv *FontDescription) Free() {
	C.pango_font_description_free((*C.PangoFontDescription)(recv.native))

	return
}

// GetFamily is a wrapper around the C function pango_font_description_get_family.
func (recv *FontDescription) GetFamily() string {
	retC := C.pango_font_description_get_family((*C.PangoFontDescription)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetGravity is a wrapper around the C function pango_font_description_get_gravity.
func (recv *FontDescription) GetGravity() Gravity {
	retC := C.pango_font_description_get_gravity((*C.PangoFontDescription)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// GetSetFields is a wrapper around the C function pango_font_description_get_set_fields.
func (recv *FontDescription) GetSetFields() FontMask {
	retC := C.pango_font_description_get_set_fields((*C.PangoFontDescription)(recv.native))
	retGo := (FontMask)(retC)

	return retGo
}

// GetSize is a wrapper around the C function pango_font_description_get_size.
func (recv *FontDescription) GetSize() int32 {
	retC := C.pango_font_description_get_size((*C.PangoFontDescription)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSizeIsAbsolute is a wrapper around the C function pango_font_description_get_size_is_absolute.
func (recv *FontDescription) GetSizeIsAbsolute() bool {
	retC := C.pango_font_description_get_size_is_absolute((*C.PangoFontDescription)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetStretch is a wrapper around the C function pango_font_description_get_stretch.
func (recv *FontDescription) GetStretch() Stretch {
	retC := C.pango_font_description_get_stretch((*C.PangoFontDescription)(recv.native))
	retGo := (Stretch)(retC)

	return retGo
}

// GetStyle is a wrapper around the C function pango_font_description_get_style.
func (recv *FontDescription) GetStyle() Style {
	retC := C.pango_font_description_get_style((*C.PangoFontDescription)(recv.native))
	retGo := (Style)(retC)

	return retGo
}

// GetVariant is a wrapper around the C function pango_font_description_get_variant.
func (recv *FontDescription) GetVariant() Variant {
	retC := C.pango_font_description_get_variant((*C.PangoFontDescription)(recv.native))
	retGo := (Variant)(retC)

	return retGo
}

// GetWeight is a wrapper around the C function pango_font_description_get_weight.
func (recv *FontDescription) GetWeight() Weight {
	retC := C.pango_font_description_get_weight((*C.PangoFontDescription)(recv.native))
	retGo := (Weight)(retC)

	return retGo
}

// Hash is a wrapper around the C function pango_font_description_hash.
func (recv *FontDescription) Hash() uint32 {
	retC := C.pango_font_description_hash((*C.PangoFontDescription)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Merge is a wrapper around the C function pango_font_description_merge.
func (recv *FontDescription) Merge(descToMerge *FontDescription, replaceExisting bool) {
	c_desc_to_merge := (*C.PangoFontDescription)(C.NULL)
	if descToMerge != nil {
		c_desc_to_merge = (*C.PangoFontDescription)(descToMerge.ToC())
	}

	c_replace_existing :=
		boolToGboolean(replaceExisting)

	C.pango_font_description_merge((*C.PangoFontDescription)(recv.native), c_desc_to_merge, c_replace_existing)

	return
}

// MergeStatic is a wrapper around the C function pango_font_description_merge_static.
func (recv *FontDescription) MergeStatic(descToMerge *FontDescription, replaceExisting bool) {
	c_desc_to_merge := (*C.PangoFontDescription)(C.NULL)
	if descToMerge != nil {
		c_desc_to_merge = (*C.PangoFontDescription)(descToMerge.ToC())
	}

	c_replace_existing :=
		boolToGboolean(replaceExisting)

	C.pango_font_description_merge_static((*C.PangoFontDescription)(recv.native), c_desc_to_merge, c_replace_existing)

	return
}

// SetAbsoluteSize is a wrapper around the C function pango_font_description_set_absolute_size.
func (recv *FontDescription) SetAbsoluteSize(size float64) {
	c_size := (C.double)(size)

	C.pango_font_description_set_absolute_size((*C.PangoFontDescription)(recv.native), c_size)

	return
}

// SetFamily is a wrapper around the C function pango_font_description_set_family.
func (recv *FontDescription) SetFamily(family string) {
	c_family := C.CString(family)
	defer C.free(unsafe.Pointer(c_family))

	C.pango_font_description_set_family((*C.PangoFontDescription)(recv.native), c_family)

	return
}

// SetFamilyStatic is a wrapper around the C function pango_font_description_set_family_static.
func (recv *FontDescription) SetFamilyStatic(family string) {
	c_family := C.CString(family)
	defer C.free(unsafe.Pointer(c_family))

	C.pango_font_description_set_family_static((*C.PangoFontDescription)(recv.native), c_family)

	return
}

// SetGravity is a wrapper around the C function pango_font_description_set_gravity.
func (recv *FontDescription) SetGravity(gravity Gravity) {
	c_gravity := (C.PangoGravity)(gravity)

	C.pango_font_description_set_gravity((*C.PangoFontDescription)(recv.native), c_gravity)

	return
}

// SetSize is a wrapper around the C function pango_font_description_set_size.
func (recv *FontDescription) SetSize(size int32) {
	c_size := (C.gint)(size)

	C.pango_font_description_set_size((*C.PangoFontDescription)(recv.native), c_size)

	return
}

// SetStretch is a wrapper around the C function pango_font_description_set_stretch.
func (recv *FontDescription) SetStretch(stretch Stretch) {
	c_stretch := (C.PangoStretch)(stretch)

	C.pango_font_description_set_stretch((*C.PangoFontDescription)(recv.native), c_stretch)

	return
}

// SetStyle is a wrapper around the C function pango_font_description_set_style.
func (recv *FontDescription) SetStyle(style Style) {
	c_style := (C.PangoStyle)(style)

	C.pango_font_description_set_style((*C.PangoFontDescription)(recv.native), c_style)

	return
}

// SetVariant is a wrapper around the C function pango_font_description_set_variant.
func (recv *FontDescription) SetVariant(variant Variant) {
	c_variant := (C.PangoVariant)(variant)

	C.pango_font_description_set_variant((*C.PangoFontDescription)(recv.native), c_variant)

	return
}

// SetWeight is a wrapper around the C function pango_font_description_set_weight.
func (recv *FontDescription) SetWeight(weight Weight) {
	c_weight := (C.PangoWeight)(weight)

	C.pango_font_description_set_weight((*C.PangoFontDescription)(recv.native), c_weight)

	return
}

// ToFilename is a wrapper around the C function pango_font_description_to_filename.
func (recv *FontDescription) ToFilename() string {
	retC := C.pango_font_description_to_filename((*C.PangoFontDescription)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ToString is a wrapper around the C function pango_font_description_to_string.
func (recv *FontDescription) ToString() string {
	retC := C.pango_font_description_to_string((*C.PangoFontDescription)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UnsetFields is a wrapper around the C function pango_font_description_unset_fields.
func (recv *FontDescription) UnsetFields(toUnset FontMask) {
	c_to_unset := (C.PangoFontMask)(toUnset)

	C.pango_font_description_unset_fields((*C.PangoFontDescription)(recv.native), c_to_unset)

	return
}

// Blacklisted : PangoFontFaceClass

// Blacklisted : PangoFontFamilyClass

// Blacklisted : PangoFontMapClass

// FontMetrics is a wrapper around the C record PangoFontMetrics.
type FontMetrics struct {
	native *C.PangoFontMetrics
	// Private : ref_count
	// Private : ascent
	// Private : descent
	// Private : approximate_char_width
	// Private : approximate_digit_width
	// Private : underline_position
	// Private : underline_thickness
	// Private : strikethrough_position
	// Private : strikethrough_thickness
}

func FontMetricsNewFromC(u unsafe.Pointer) *FontMetrics {
	c := (*C.PangoFontMetrics)(u)
	if c == nil {
		return nil
	}

	g := &FontMetrics{native: c}

	return g
}

func (recv *FontMetrics) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FontMetrics with another FontMetrics, and returns true if they represent the same GObject.
func (recv *FontMetrics) Equals(other *FontMetrics) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : pango_font_metrics_new

// GetApproximateCharWidth is a wrapper around the C function pango_font_metrics_get_approximate_char_width.
func (recv *FontMetrics) GetApproximateCharWidth() int32 {
	retC := C.pango_font_metrics_get_approximate_char_width((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetApproximateDigitWidth is a wrapper around the C function pango_font_metrics_get_approximate_digit_width.
func (recv *FontMetrics) GetApproximateDigitWidth() int32 {
	retC := C.pango_font_metrics_get_approximate_digit_width((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetAscent is a wrapper around the C function pango_font_metrics_get_ascent.
func (recv *FontMetrics) GetAscent() int32 {
	retC := C.pango_font_metrics_get_ascent((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDescent is a wrapper around the C function pango_font_metrics_get_descent.
func (recv *FontMetrics) GetDescent() int32 {
	retC := C.pango_font_metrics_get_descent((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetStrikethroughPosition is a wrapper around the C function pango_font_metrics_get_strikethrough_position.
func (recv *FontMetrics) GetStrikethroughPosition() int32 {
	retC := C.pango_font_metrics_get_strikethrough_position((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetStrikethroughThickness is a wrapper around the C function pango_font_metrics_get_strikethrough_thickness.
func (recv *FontMetrics) GetStrikethroughThickness() int32 {
	retC := C.pango_font_metrics_get_strikethrough_thickness((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetUnderlinePosition is a wrapper around the C function pango_font_metrics_get_underline_position.
func (recv *FontMetrics) GetUnderlinePosition() int32 {
	retC := C.pango_font_metrics_get_underline_position((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetUnderlineThickness is a wrapper around the C function pango_font_metrics_get_underline_thickness.
func (recv *FontMetrics) GetUnderlineThickness() int32 {
	retC := C.pango_font_metrics_get_underline_thickness((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Ref is a wrapper around the C function pango_font_metrics_ref.
func (recv *FontMetrics) Ref() *FontMetrics {
	retC := C.pango_font_metrics_ref((*C.PangoFontMetrics)(recv.native))
	var retGo (*FontMetrics)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontMetricsNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unref is a wrapper around the C function pango_font_metrics_unref.
func (recv *FontMetrics) Unref() {
	C.pango_font_metrics_unref((*C.PangoFontMetrics)(recv.native))

	return
}

// Blacklisted : PangoFontsetClass

// Blacklisted : PangoFontsetSimpleClass

// GlyphGeometry is a wrapper around the C record PangoGlyphGeometry.
type GlyphGeometry struct {
	native  *C.PangoGlyphGeometry
	Width   GlyphUnit
	XOffset GlyphUnit
	YOffset GlyphUnit
}

func GlyphGeometryNewFromC(u unsafe.Pointer) *GlyphGeometry {
	c := (*C.PangoGlyphGeometry)(u)
	if c == nil {
		return nil
	}

	g := &GlyphGeometry{
		Width:   (GlyphUnit)(c.width),
		XOffset: (GlyphUnit)(c.x_offset),
		YOffset: (GlyphUnit)(c.y_offset),
		native:  c,
	}

	return g
}

func (recv *GlyphGeometry) ToC() unsafe.Pointer {
	recv.native.width =
		(C.PangoGlyphUnit)(recv.Width)
	recv.native.x_offset =
		(C.PangoGlyphUnit)(recv.XOffset)
	recv.native.y_offset =
		(C.PangoGlyphUnit)(recv.YOffset)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GlyphGeometry with another GlyphGeometry, and returns true if they represent the same GObject.
func (recv *GlyphGeometry) Equals(other *GlyphGeometry) bool {
	return other.ToC() == recv.ToC()
}

// GlyphInfo is a wrapper around the C record PangoGlyphInfo.
type GlyphInfo struct {
	native *C.PangoGlyphInfo
	Glyph  Glyph
	// geometry : record
	// attr : record
}

func GlyphInfoNewFromC(u unsafe.Pointer) *GlyphInfo {
	c := (*C.PangoGlyphInfo)(u)
	if c == nil {
		return nil
	}

	g := &GlyphInfo{
		Glyph:  (Glyph)(c.glyph),
		native: c,
	}

	return g
}

func (recv *GlyphInfo) ToC() unsafe.Pointer {
	recv.native.glyph =
		(C.PangoGlyph)(recv.Glyph)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GlyphInfo with another GlyphInfo, and returns true if they represent the same GObject.
func (recv *GlyphInfo) Equals(other *GlyphInfo) bool {
	return other.ToC() == recv.ToC()
}

// GlyphItem is a wrapper around the C record PangoGlyphItem.
type GlyphItem struct {
	native *C.PangoGlyphItem
	// item : record
	// glyphs : record
}

func GlyphItemNewFromC(u unsafe.Pointer) *GlyphItem {
	c := (*C.PangoGlyphItem)(u)
	if c == nil {
		return nil
	}

	g := &GlyphItem{native: c}

	return g
}

func (recv *GlyphItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GlyphItem with another GlyphItem, and returns true if they represent the same GObject.
func (recv *GlyphItem) Equals(other *GlyphItem) bool {
	return other.ToC() == recv.ToC()
}

// ApplyAttrs is a wrapper around the C function pango_glyph_item_apply_attrs.
func (recv *GlyphItem) ApplyAttrs(text string, list *AttrList) *glib.SList {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_list := (*C.PangoAttrList)(C.NULL)
	if list != nil {
		c_list = (*C.PangoAttrList)(list.ToC())
	}

	retC := C.pango_glyph_item_apply_attrs((*C.PangoGlyphItem)(recv.native), c_text, c_list)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function pango_glyph_item_free.
func (recv *GlyphItem) Free() {
	C.pango_glyph_item_free((*C.PangoGlyphItem)(recv.native))

	return
}

// Unsupported : pango_glyph_item_letter_space : unsupported parameter log_attrs :

// Split is a wrapper around the C function pango_glyph_item_split.
func (recv *GlyphItem) Split(text string, splitIndex int32) *GlyphItem {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_split_index := (C.int)(splitIndex)

	retC := C.pango_glyph_item_split((*C.PangoGlyphItem)(recv.native), c_text, c_split_index)
	retGo := GlyphItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : PangoGlyphString

// GlyphVisAttr is a wrapper around the C record PangoGlyphVisAttr.
type GlyphVisAttr struct {
	native *C.PangoGlyphVisAttr
	// Bitfield not supported :  1 is_cluster_start
}

func GlyphVisAttrNewFromC(u unsafe.Pointer) *GlyphVisAttr {
	c := (*C.PangoGlyphVisAttr)(u)
	if c == nil {
		return nil
	}

	g := &GlyphVisAttr{native: c}

	return g
}

func (recv *GlyphVisAttr) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GlyphVisAttr with another GlyphVisAttr, and returns true if they represent the same GObject.
func (recv *GlyphVisAttr) Equals(other *GlyphVisAttr) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : PangoIncludedModule

// Item is a wrapper around the C record PangoItem.
type Item struct {
	native   *C.PangoItem
	Offset   int32
	Length   int32
	NumChars int32
	// analysis : record
}

func ItemNewFromC(u unsafe.Pointer) *Item {
	c := (*C.PangoItem)(u)
	if c == nil {
		return nil
	}

	g := &Item{
		Length:   (int32)(c.length),
		NumChars: (int32)(c.num_chars),
		Offset:   (int32)(c.offset),
		native:   c,
	}

	return g
}

func (recv *Item) ToC() unsafe.Pointer {
	recv.native.offset =
		(C.gint)(recv.Offset)
	recv.native.length =
		(C.gint)(recv.Length)
	recv.native.num_chars =
		(C.gint)(recv.NumChars)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Item with another Item, and returns true if they represent the same GObject.
func (recv *Item) Equals(other *Item) bool {
	return other.ToC() == recv.ToC()
}

// ItemNew is a wrapper around the C function pango_item_new.
func ItemNew() *Item {
	retC := C.pango_item_new()
	retGo := ItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function pango_item_copy.
func (recv *Item) Copy() *Item {
	retC := C.pango_item_copy((*C.PangoItem)(recv.native))
	var retGo (*Item)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Free is a wrapper around the C function pango_item_free.
func (recv *Item) Free() {
	C.pango_item_free((*C.PangoItem)(recv.native))

	return
}

// Split is a wrapper around the C function pango_item_split.
func (recv *Item) Split(splitIndex int32, splitOffset int32) *Item {
	c_split_index := (C.int)(splitIndex)

	c_split_offset := (C.int)(splitOffset)

	retC := C.pango_item_split((*C.PangoItem)(recv.native), c_split_index, c_split_offset)
	retGo := ItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Language is a wrapper around the C record PangoLanguage.
type Language struct {
	native *C.PangoLanguage
}

func LanguageNewFromC(u unsafe.Pointer) *Language {
	c := (*C.PangoLanguage)(u)
	if c == nil {
		return nil
	}

	g := &Language{native: c}

	return g
}

func (recv *Language) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Language with another Language, and returns true if they represent the same GObject.
func (recv *Language) Equals(other *Language) bool {
	return other.ToC() == recv.ToC()
}

// LanguageFromString is a wrapper around the C function pango_language_from_string.
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

// LanguageGetDefault is a wrapper around the C function pango_language_get_default.
func LanguageGetDefault() *Language {
	retC := C.pango_language_get_default()
	retGo := LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSampleString is a wrapper around the C function pango_language_get_sample_string.
func (recv *Language) GetSampleString() string {
	retC := C.pango_language_get_sample_string((*C.PangoLanguage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// IncludesScript is a wrapper around the C function pango_language_includes_script.
func (recv *Language) IncludesScript(script Script) bool {
	c_script := (C.PangoScript)(script)

	retC := C.pango_language_includes_script((*C.PangoLanguage)(recv.native), c_script)
	retGo := retC == C.TRUE

	return retGo
}

// Matches is a wrapper around the C function pango_language_matches.
func (recv *Language) Matches(rangeList string) bool {
	c_range_list := C.CString(rangeList)
	defer C.free(unsafe.Pointer(c_range_list))

	retC := C.pango_language_matches((*C.PangoLanguage)(recv.native), c_range_list)
	retGo := retC == C.TRUE

	return retGo
}

// ToString is a wrapper around the C function pango_language_to_string.
func (recv *Language) ToString() string {
	retC := C.pango_language_to_string((*C.PangoLanguage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// LayoutClass is a wrapper around the C record PangoLayoutClass.
type LayoutClass struct {
	native *C.PangoLayoutClass
}

func LayoutClassNewFromC(u unsafe.Pointer) *LayoutClass {
	c := (*C.PangoLayoutClass)(u)
	if c == nil {
		return nil
	}

	g := &LayoutClass{native: c}

	return g
}

func (recv *LayoutClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LayoutClass with another LayoutClass, and returns true if they represent the same GObject.
func (recv *LayoutClass) Equals(other *LayoutClass) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : PangoLayoutIter

// LayoutLine is a wrapper around the C record PangoLayoutLine.
type LayoutLine struct {
	native *C.PangoLayoutLine
	// layout : record
	StartIndex int32
	Length     int32
	// runs : record
	// Bitfield not supported :  1 is_paragraph_start
	// Bitfield not supported :  3 resolved_dir
}

func LayoutLineNewFromC(u unsafe.Pointer) *LayoutLine {
	c := (*C.PangoLayoutLine)(u)
	if c == nil {
		return nil
	}

	g := &LayoutLine{
		Length:     (int32)(c.length),
		StartIndex: (int32)(c.start_index),
		native:     c,
	}

	return g
}

func (recv *LayoutLine) ToC() unsafe.Pointer {
	recv.native.start_index =
		(C.gint)(recv.StartIndex)
	recv.native.length =
		(C.gint)(recv.Length)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LayoutLine with another LayoutLine, and returns true if they represent the same GObject.
func (recv *LayoutLine) Equals(other *LayoutLine) bool {
	return other.ToC() == recv.ToC()
}

// GetExtents is a wrapper around the C function pango_layout_line_get_extents.
func (recv *LayoutLine) GetExtents() (*Rectangle, *Rectangle) {
	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_layout_line_get_extents((*C.PangoLayoutLine)(recv.native), &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// GetPixelExtents is a wrapper around the C function pango_layout_line_get_pixel_extents.
func (recv *LayoutLine) GetPixelExtents() (*Rectangle, *Rectangle) {
	var c_ink_rect C.PangoRectangle

	var c_logical_rect C.PangoRectangle

	C.pango_layout_line_get_pixel_extents((*C.PangoLayoutLine)(recv.native), &c_ink_rect, &c_logical_rect)

	inkRect := RectangleNewFromC(unsafe.Pointer(&c_ink_rect))

	logicalRect := RectangleNewFromC(unsafe.Pointer(&c_logical_rect))

	return inkRect, logicalRect
}

// Unsupported : pango_layout_line_get_x_ranges : unsupported parameter ranges : output array param ranges

// IndexToX is a wrapper around the C function pango_layout_line_index_to_x.
func (recv *LayoutLine) IndexToX(index int32, trailing bool) int32 {
	c_index_ := (C.int)(index)

	c_trailing :=
		boolToGboolean(trailing)

	var c_x_pos C.int

	C.pango_layout_line_index_to_x((*C.PangoLayoutLine)(recv.native), c_index_, c_trailing, &c_x_pos)

	xPos := (int32)(c_x_pos)

	return xPos
}

// Ref is a wrapper around the C function pango_layout_line_ref.
func (recv *LayoutLine) Ref() *LayoutLine {
	retC := C.pango_layout_line_ref((*C.PangoLayoutLine)(recv.native))
	retGo := LayoutLineNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function pango_layout_line_unref.
func (recv *LayoutLine) Unref() {
	C.pango_layout_line_unref((*C.PangoLayoutLine)(recv.native))

	return
}

// XToIndex is a wrapper around the C function pango_layout_line_x_to_index.
func (recv *LayoutLine) XToIndex(xPos int32) (bool, int32, int32) {
	c_x_pos := (C.int)(xPos)

	var c_index_ C.int

	var c_trailing C.int

	retC := C.pango_layout_line_x_to_index((*C.PangoLayoutLine)(recv.native), c_x_pos, &c_index_, &c_trailing)
	retGo := retC == C.TRUE

	index := (int32)(c_index_)

	trailing := (int32)(c_trailing)

	return retGo, index, trailing
}

// LogAttr is a wrapper around the C record PangoLogAttr.
type LogAttr struct {
	native *C.PangoLogAttr
	// Bitfield not supported :  1 is_line_break
	// Bitfield not supported :  1 is_mandatory_break
	// Bitfield not supported :  1 is_char_break
	// Bitfield not supported :  1 is_white
	// Bitfield not supported :  1 is_cursor_position
	// Bitfield not supported :  1 is_word_start
	// Bitfield not supported :  1 is_word_end
	// Bitfield not supported :  1 is_sentence_boundary
	// Bitfield not supported :  1 is_sentence_start
	// Bitfield not supported :  1 is_sentence_end
	// Bitfield not supported :  1 backspace_deletes_character
	// Bitfield not supported :  1 is_expandable_space
	// Bitfield not supported :  1 is_word_boundary
}

func LogAttrNewFromC(u unsafe.Pointer) *LogAttr {
	c := (*C.PangoLogAttr)(u)
	if c == nil {
		return nil
	}

	g := &LogAttr{native: c}

	return g
}

func (recv *LogAttr) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LogAttr with another LogAttr, and returns true if they represent the same GObject.
func (recv *LogAttr) Equals(other *LogAttr) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : PangoMap

// Blacklisted : PangoMapEntry

// Matrix is a wrapper around the C record PangoMatrix.
type Matrix struct {
	native *C.PangoMatrix
	Xx     float64
	Xy     float64
	Yx     float64
	Yy     float64
	X0     float64
	Y0     float64
}

func MatrixNewFromC(u unsafe.Pointer) *Matrix {
	c := (*C.PangoMatrix)(u)
	if c == nil {
		return nil
	}

	g := &Matrix{
		X0:     (float64)(c.x0),
		Xx:     (float64)(c.xx),
		Xy:     (float64)(c.xy),
		Y0:     (float64)(c.y0),
		Yx:     (float64)(c.yx),
		Yy:     (float64)(c.yy),
		native: c,
	}

	return g
}

func (recv *Matrix) ToC() unsafe.Pointer {
	recv.native.xx =
		(C.double)(recv.Xx)
	recv.native.xy =
		(C.double)(recv.Xy)
	recv.native.yx =
		(C.double)(recv.Yx)
	recv.native.yy =
		(C.double)(recv.Yy)
	recv.native.x0 =
		(C.double)(recv.X0)
	recv.native.y0 =
		(C.double)(recv.Y0)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Matrix with another Matrix, and returns true if they represent the same GObject.
func (recv *Matrix) Equals(other *Matrix) bool {
	return other.ToC() == recv.ToC()
}

// Concat is a wrapper around the C function pango_matrix_concat.
func (recv *Matrix) Concat(newMatrix *Matrix) {
	c_new_matrix := (*C.PangoMatrix)(C.NULL)
	if newMatrix != nil {
		c_new_matrix = (*C.PangoMatrix)(newMatrix.ToC())
	}

	C.pango_matrix_concat((*C.PangoMatrix)(recv.native), c_new_matrix)

	return
}

// Copy is a wrapper around the C function pango_matrix_copy.
func (recv *Matrix) Copy() *Matrix {
	retC := C.pango_matrix_copy((*C.PangoMatrix)(recv.native))
	var retGo (*Matrix)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MatrixNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Free is a wrapper around the C function pango_matrix_free.
func (recv *Matrix) Free() {
	C.pango_matrix_free((*C.PangoMatrix)(recv.native))

	return
}

// GetFontScaleFactor is a wrapper around the C function pango_matrix_get_font_scale_factor.
func (recv *Matrix) GetFontScaleFactor() float64 {
	retC := C.pango_matrix_get_font_scale_factor((*C.PangoMatrix)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Rotate is a wrapper around the C function pango_matrix_rotate.
func (recv *Matrix) Rotate(degrees float64) {
	c_degrees := (C.double)(degrees)

	C.pango_matrix_rotate((*C.PangoMatrix)(recv.native), c_degrees)

	return
}

// Scale is a wrapper around the C function pango_matrix_scale.
func (recv *Matrix) Scale(scaleX float64, scaleY float64) {
	c_scale_x := (C.double)(scaleX)

	c_scale_y := (C.double)(scaleY)

	C.pango_matrix_scale((*C.PangoMatrix)(recv.native), c_scale_x, c_scale_y)

	return
}

// TransformDistance is a wrapper around the C function pango_matrix_transform_distance.
func (recv *Matrix) TransformDistance(dx float64, dy float64) {
	c_dx := (C.double)(dx)

	c_dy := (C.double)(dy)

	C.pango_matrix_transform_distance((*C.PangoMatrix)(recv.native), &c_dx, &c_dy)

	return
}

// TransformPixelRectangle is a wrapper around the C function pango_matrix_transform_pixel_rectangle.
func (recv *Matrix) TransformPixelRectangle(rect *Rectangle) {
	c_rect := (*C.PangoRectangle)(C.NULL)
	if rect != nil {
		c_rect = (*C.PangoRectangle)(rect.ToC())
	}

	C.pango_matrix_transform_pixel_rectangle((*C.PangoMatrix)(recv.native), c_rect)

	return
}

// TransformPoint is a wrapper around the C function pango_matrix_transform_point.
func (recv *Matrix) TransformPoint(x float64, y float64) {
	c_x := (C.double)(x)

	c_y := (C.double)(y)

	C.pango_matrix_transform_point((*C.PangoMatrix)(recv.native), &c_x, &c_y)

	return
}

// TransformRectangle is a wrapper around the C function pango_matrix_transform_rectangle.
func (recv *Matrix) TransformRectangle(rect *Rectangle) {
	c_rect := (*C.PangoRectangle)(C.NULL)
	if rect != nil {
		c_rect = (*C.PangoRectangle)(rect.ToC())
	}

	C.pango_matrix_transform_rectangle((*C.PangoMatrix)(recv.native), c_rect)

	return
}

// Translate is a wrapper around the C function pango_matrix_translate.
func (recv *Matrix) Translate(tx float64, ty float64) {
	c_tx := (C.double)(tx)

	c_ty := (C.double)(ty)

	C.pango_matrix_translate((*C.PangoMatrix)(recv.native), c_tx, c_ty)

	return
}

// Rectangle is a wrapper around the C record PangoRectangle.
type Rectangle struct {
	native *C.PangoRectangle
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func RectangleNewFromC(u unsafe.Pointer) *Rectangle {
	c := (*C.PangoRectangle)(u)
	if c == nil {
		return nil
	}

	g := &Rectangle{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *Rectangle) ToC() unsafe.Pointer {
	recv.native.x =
		(C.int)(recv.X)
	recv.native.y =
		(C.int)(recv.Y)
	recv.native.width =
		(C.int)(recv.Width)
	recv.native.height =
		(C.int)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Rectangle with another Rectangle, and returns true if they represent the same GObject.
func (recv *Rectangle) Equals(other *Rectangle) bool {
	return other.ToC() == recv.ToC()
}

// RendererClass is a wrapper around the C record PangoRendererClass.
type RendererClass struct {
	native *C.PangoRendererClass
	// Private : parent_class
	// no type for draw_glyphs
	// no type for draw_rectangle
	// no type for draw_error_underline
	// no type for draw_shape
	// no type for draw_trapezoid
	// no type for draw_glyph
	// no type for part_changed
	// no type for begin
	// no type for end
	// no type for prepare_run
	// no type for draw_glyph_item
	// no type for _pango_reserved2
	// no type for _pango_reserved3
	// no type for _pango_reserved4
}

func RendererClassNewFromC(u unsafe.Pointer) *RendererClass {
	c := (*C.PangoRendererClass)(u)
	if c == nil {
		return nil
	}

	g := &RendererClass{native: c}

	return g
}

func (recv *RendererClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RendererClass with another RendererClass, and returns true if they represent the same GObject.
func (recv *RendererClass) Equals(other *RendererClass) bool {
	return other.ToC() == recv.ToC()
}

// RendererPrivate is a wrapper around the C record PangoRendererPrivate.
type RendererPrivate struct {
	native *C.PangoRendererPrivate
}

func RendererPrivateNewFromC(u unsafe.Pointer) *RendererPrivate {
	c := (*C.PangoRendererPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RendererPrivate{native: c}

	return g
}

func (recv *RendererPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RendererPrivate with another RendererPrivate, and returns true if they represent the same GObject.
func (recv *RendererPrivate) Equals(other *RendererPrivate) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : PangoScriptForLang

// ScriptIter is a wrapper around the C record PangoScriptIter.
type ScriptIter struct {
	native *C.PangoScriptIter
}

func ScriptIterNewFromC(u unsafe.Pointer) *ScriptIter {
	c := (*C.PangoScriptIter)(u)
	if c == nil {
		return nil
	}

	g := &ScriptIter{native: c}

	return g
}

func (recv *ScriptIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ScriptIter with another ScriptIter, and returns true if they represent the same GObject.
func (recv *ScriptIter) Equals(other *ScriptIter) bool {
	return other.ToC() == recv.ToC()
}

// ScriptIterNew is a wrapper around the C function pango_script_iter_new.
func ScriptIterNew(text string, length int32) *ScriptIter {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.int)(length)

	retC := C.pango_script_iter_new(c_text, c_length)
	retGo := ScriptIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function pango_script_iter_free.
func (recv *ScriptIter) Free() {
	C.pango_script_iter_free((*C.PangoScriptIter)(recv.native))

	return
}

// GetRange is a wrapper around the C function pango_script_iter_get_range.
func (recv *ScriptIter) GetRange() (string, string, Script) {
	var c_start *C.char

	var c_end *C.char

	var c_script C.PangoScript

	C.pango_script_iter_get_range((*C.PangoScriptIter)(recv.native), &c_start, &c_end, &c_script)

	start := C.GoString(c_start)
	defer C.free(unsafe.Pointer(c_start))

	end := C.GoString(c_end)
	defer C.free(unsafe.Pointer(c_end))

	script := (Script)(c_script)

	return start, end, script
}

// Next is a wrapper around the C function pango_script_iter_next.
func (recv *ScriptIter) Next() bool {
	retC := C.pango_script_iter_next((*C.PangoScriptIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// TabArray is a wrapper around the C record PangoTabArray.
type TabArray struct {
	native *C.PangoTabArray
}

func TabArrayNewFromC(u unsafe.Pointer) *TabArray {
	c := (*C.PangoTabArray)(u)
	if c == nil {
		return nil
	}

	g := &TabArray{native: c}

	return g
}

func (recv *TabArray) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TabArray with another TabArray, and returns true if they represent the same GObject.
func (recv *TabArray) Equals(other *TabArray) bool {
	return other.ToC() == recv.ToC()
}

// TabArrayNew is a wrapper around the C function pango_tab_array_new.
func TabArrayNew(initialSize int32, positionsInPixels bool) *TabArray {
	c_initial_size := (C.gint)(initialSize)

	c_positions_in_pixels :=
		boolToGboolean(positionsInPixels)

	retC := C.pango_tab_array_new(c_initial_size, c_positions_in_pixels)
	retGo := TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_tab_array_new_with_positions : unsupported parameter ... : varargs

// Copy is a wrapper around the C function pango_tab_array_copy.
func (recv *TabArray) Copy() *TabArray {
	retC := C.pango_tab_array_copy((*C.PangoTabArray)(recv.native))
	retGo := TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function pango_tab_array_free.
func (recv *TabArray) Free() {
	C.pango_tab_array_free((*C.PangoTabArray)(recv.native))

	return
}

// GetPositionsInPixels is a wrapper around the C function pango_tab_array_get_positions_in_pixels.
func (recv *TabArray) GetPositionsInPixels() bool {
	retC := C.pango_tab_array_get_positions_in_pixels((*C.PangoTabArray)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSize is a wrapper around the C function pango_tab_array_get_size.
func (recv *TabArray) GetSize() int32 {
	retC := C.pango_tab_array_get_size((*C.PangoTabArray)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTab is a wrapper around the C function pango_tab_array_get_tab.
func (recv *TabArray) GetTab(tabIndex int32) (TabAlign, int32) {
	c_tab_index := (C.gint)(tabIndex)

	var c_alignment C.PangoTabAlign

	var c_location C.gint

	C.pango_tab_array_get_tab((*C.PangoTabArray)(recv.native), c_tab_index, &c_alignment, &c_location)

	alignment := (TabAlign)(c_alignment)

	location := (int32)(c_location)

	return alignment, location
}

// Unsupported : pango_tab_array_get_tabs : unsupported parameter alignments : PangoTabAlign** with indirection level of 2

// Resize is a wrapper around the C function pango_tab_array_resize.
func (recv *TabArray) Resize(newSize int32) {
	c_new_size := (C.gint)(newSize)

	C.pango_tab_array_resize((*C.PangoTabArray)(recv.native), c_new_size)

	return
}

// SetTab is a wrapper around the C function pango_tab_array_set_tab.
func (recv *TabArray) SetTab(tabIndex int32, alignment TabAlign, location int32) {
	c_tab_index := (C.gint)(tabIndex)

	c_alignment := (C.PangoTabAlign)(alignment)

	c_location := (C.gint)(location)

	C.pango_tab_array_set_tab((*C.PangoTabArray)(recv.native), c_tab_index, c_alignment, c_location)

	return
}
