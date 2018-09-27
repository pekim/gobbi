// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.16 pango_1.22 pango_1.26 pango_1.31.0 pango_1.32 pango_1.38

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// AttrFallbackNew is a wrapper around the C function pango_attr_fallback_new.
func AttrFallbackNew(enableFallback bool) *Attribute {
	c_enable_fallback :=
		boolToGboolean(enableFallback)

	retC := C.pango_attr_fallback_new(c_enable_fallback)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
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

// ItemizeWithBaseDir is a wrapper around the C function pango_itemize_with_base_dir.
func ItemizeWithBaseDir(context *Context, baseDir Direction, text string, startIndex int32, length int32, attrs *AttrList, cachedIter *AttrIterator) *glib.List {
	c_context := (*C.PangoContext)(context.ToC())

	c_base_dir := (C.PangoDirection)(baseDir)

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_start_index := (C.int)(startIndex)

	c_length := (C.int)(length)

	c_attrs := (*C.PangoAttrList)(attrs.ToC())

	c_cached_iter := (*C.PangoAttrIterator)(cachedIter.ToC())

	retC := C.pango_itemize_with_base_dir(c_context, c_base_dir, c_text, c_start_index, c_length, c_attrs, c_cached_iter)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_log2vis_get_embedding_levels : unsupported parameter pbase_dir : PangoDirection* with indirection level of 1

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
	retGo := LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}