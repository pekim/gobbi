// This is a generated file - DO NOT EDIT

package pango

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

// Blacklisted : pango_context_get_base_dir

// Blacklisted : pango_context_get_font_description

// Blacklisted : pango_context_get_language

// Blacklisted : pango_context_get_metrics

// Unsupported : pango_context_list_families : unsupported parameter families : output array param families

// Blacklisted : pango_context_load_font

// Blacklisted : pango_context_load_fontset

// Blacklisted : pango_context_set_base_dir

// Blacklisted : pango_context_set_font_description

// Blacklisted : pango_context_set_font_map

// Blacklisted : pango_context_set_language

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

// CastToWidget down casts any arbitrary Object to EngineLang.
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

// CastToWidget down casts any arbitrary Object to EngineShape.
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

// CastToWidget down casts any arbitrary Object to Font.
// Exercise care, as this is a potentially dangerous function if the Object is not a Font.
func CastToFont(object *gobject.Object) *Font {
	return FontNewFromC(object.ToC())
}

// pango_font_descriptions_free : unsupported parameter descs :
// Blacklisted : pango_font_describe

// Blacklisted : pango_font_find_shaper

// Blacklisted : pango_font_get_coverage

// Blacklisted : pango_font_get_glyph_extents

// Blacklisted : pango_font_get_metrics

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

// Blacklisted : pango_font_face_describe

// Blacklisted : pango_font_face_get_face_name

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

// Blacklisted : pango_font_family_get_name

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

// CastToWidget down casts any arbitrary Object to FontMap.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontMap.
func CastToFontMap(object *gobject.Object) *FontMap {
	return FontMapNewFromC(object.ToC())
}

// Unsupported : pango_font_map_list_families : unsupported parameter families : output array param families

// Blacklisted : pango_font_map_load_font

// Blacklisted : pango_font_map_load_fontset

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

// Blacklisted : pango_fontset_get_font

// Blacklisted : pango_fontset_get_metrics

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

// CastToWidget down casts any arbitrary Object to Layout.
// Exercise care, as this is a potentially dangerous function if the Object is not a Layout.
func CastToLayout(object *gobject.Object) *Layout {
	return LayoutNewFromC(object.ToC())
}

// Blacklisted : pango_layout_new

// Blacklisted : pango_layout_context_changed

// Blacklisted : pango_layout_copy

// Blacklisted : pango_layout_get_alignment

// Blacklisted : pango_layout_get_attributes

// Blacklisted : pango_layout_get_context

// Blacklisted : pango_layout_get_cursor_pos

// Blacklisted : pango_layout_get_extents

// Blacklisted : pango_layout_get_indent

// Unsupported : pango_layout_get_iter : return type : Blacklisted record : PangoLayoutIter

// Blacklisted : pango_layout_get_justify

// Blacklisted : pango_layout_get_line

// Blacklisted : pango_layout_get_line_count

// Blacklisted : pango_layout_get_lines

// Unsupported : pango_layout_get_log_attrs : unsupported parameter attrs : output array param attrs

// Blacklisted : pango_layout_get_pixel_extents

// Blacklisted : pango_layout_get_pixel_size

// Blacklisted : pango_layout_get_single_paragraph_mode

// Blacklisted : pango_layout_get_size

// Blacklisted : pango_layout_get_spacing

// Blacklisted : pango_layout_get_tabs

// GetText is a wrapper around the C function pango_layout_get_text.
func (recv *Layout) GetText() string {
	retC := C.pango_layout_get_text((*C.PangoLayout)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : pango_layout_get_width

// Blacklisted : pango_layout_get_wrap

// Blacklisted : pango_layout_index_to_line_x

// Blacklisted : pango_layout_index_to_pos

// Blacklisted : pango_layout_move_cursor_visually

// Blacklisted : pango_layout_set_alignment

// Blacklisted : pango_layout_set_attributes

// Blacklisted : pango_layout_set_font_description

// Blacklisted : pango_layout_set_indent

// Blacklisted : pango_layout_set_justify

// Blacklisted : pango_layout_set_markup

// Blacklisted : pango_layout_set_markup_with_accel

// Blacklisted : pango_layout_set_single_paragraph_mode

// Blacklisted : pango_layout_set_spacing

// Blacklisted : pango_layout_set_tabs

// SetText is a wrapper around the C function pango_layout_set_text.
func (recv *Layout) SetText(text string, length int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.int)(length)

	C.pango_layout_set_text((*C.PangoLayout)(recv.native), c_text, c_length)

	return
}

// Blacklisted : pango_layout_set_width

// Blacklisted : pango_layout_set_wrap

// Blacklisted : pango_layout_xy_to_index
