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

// Blacklisted : PangoEngineLang

// Blacklisted : PangoEngineShape

// Blacklisted : PangoFont

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

// Blacklisted : PangoFontFamily

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

// Blacklisted : PangoFontset

// Blacklisted : PangoFontsetSimple

// Blacklisted : PangoLayout
