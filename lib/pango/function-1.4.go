// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Create a new font fallback attribute.
//
// If fallback is disabled, characters will only be used from the
// closest matching font on the system. No fallback will be done to
// other fonts on the system that might contain the characters in the
// text.
/*

C function

pango_attr_fallback_new
*/
func AttrFallbackNew(enableFallback bool) *Attribute {
	c_enable_fallback :=
		boolToGboolean(enableFallback)

	retC := C.pango_attr_fallback_new(c_enable_fallback)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Searches a string the first character that has a strong
// direction, according to the Unicode bidirectional algorithm.
/*

C function

pango_find_base_dir
*/
func FindBaseDir(text string, length int32) Direction {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gint)(length)

	retC := C.pango_find_base_dir(c_text, c_length)
	retGo := (Direction)(retC)

	return retGo
}

// Like pango_itemize(), but the base direction to use when
// computing bidirectional levels (see pango_context_set_base_dir ()),
// is specified explicitly rather than gotten from the #PangoContext.
/*

C function

pango_itemize_with_base_dir
*/
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

// Looks up the #PangoScript for a particular character (as defined by
// Unicode Standard Annex \#24). No check is made for @ch being a
// valid Unicode character; if you pass in invalid character, the
// result is undefined.
//
// As of Pango 1.18, this function simply returns the return value of
// g_unichar_get_script().
/*

C function

pango_script_for_unichar
*/
func ScriptForUnichar(ch rune) Script {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_script_for_unichar(c_ch)
	retGo := (Script)(retC)

	return retGo
}

// Given a script, finds a language tag that is reasonably
// representative of that script. This will usually be the
// most widely spoken or used language written in that script:
// for instance, the sample language for %PANGO_SCRIPT_CYRILLIC
// is <literal>ru</literal> (Russian), the sample language
// for %PANGO_SCRIPT_ARABIC is <literal>ar</literal>.
//
// For some
// scripts, no sample language will be returned because there
// is no language that is sufficiently representative. The best
// example of this is %PANGO_SCRIPT_HAN, where various different
// variants of written Chinese, Japanese, and Korean all use
// significantly different sets of Han characters and forms
// of shared characters. No sample language can be provided
// for many historical scripts as well.
//
// As of 1.18, this function checks the environment variables
// PANGO_LANGUAGE and LANGUAGE (checked in that order) first.
// If one of them is set, it is parsed as a list of language tags
// separated by colons or other separators.  This function
// will return the first language in the parsed list that Pango
// believes may use @script for writing.  This last predicate
// is tested using pango_language_includes_script().  This can
// be used to control Pango's font selection for non-primary
// languages.  For example, a PANGO_LANGUAGE enviroment variable
// set to "en:fa" makes Pango choose fonts suitable for Persian (fa)
// instead of Arabic (ar) when a segment of Arabic text is found
// in an otherwise non-Arabic text.  The same trick can be used to
// choose a default language for %PANGO_SCRIPT_HAN when setting
// context language is not feasible.
/*

C function

pango_script_get_sample_language
*/
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
