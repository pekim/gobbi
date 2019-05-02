// Code generated - DO NOT EDIT.
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_font_face_list_sizes : unsupported parameter sizes : output array param sizes

// IsMonospace is a wrapper around the C function pango_font_family_is_monospace.
func (recv *FontFamily) IsMonospace() bool {
	retC := C.pango_font_family_is_monospace((*C.PangoFontFamily)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : pango_font_map_get_shape_engine_type

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc (PangoFontsetForeachFunc) for param func

// GetAutoDir is a wrapper around the C function pango_layout_get_auto_dir.
func (recv *Layout) GetAutoDir() bool {
	retC := C.pango_layout_get_auto_dir((*C.PangoLayout)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAutoDir is a wrapper around the C function pango_layout_set_auto_dir.
func (recv *Layout) SetAutoDir(autoDir bool) {
	c_auto_dir :=
		boolToGboolean(autoDir)

	C.pango_layout_set_auto_dir((*C.PangoLayout)(recv.native), c_auto_dir)

	return
}

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

// Unsupported : pango_log2vis_get_embedding_levels : unsupported parameter pbase_dir : PangoDirection* with indirection level of 1

// IncludesScript is a wrapper around the C function pango_language_includes_script.
func (recv *Language) IncludesScript(script Script) bool {
	c_script := (C.PangoScript)(script)

	retC := C.pango_language_includes_script((*C.PangoLanguage)(recv.native), c_script)
	retGo := retC == C.TRUE

	return retGo
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

// Unsupported : pango_script_iter_get_range : unsupported parameter script : PangoScript* with indirection level of 1

// Next is a wrapper around the C function pango_script_iter_next.
func (recv *ScriptIter) Next() bool {
	retC := C.pango_script_iter_next((*C.PangoScriptIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
