// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Determines if @script is one of the scripts used to
// write @language. The returned value is conservative;
// if nothing is known about the language tag @language,
// %TRUE will be returned, since, as far as Pango knows,
// @script might be used to write @language.
//
// This routine is used in Pango's itemization process when
// determining if a supplied language tag is relevant to
// a particular section of text. It probably is not useful for
// applications in most circumstances.
//
// This function uses pango_language_get_scripts() internally.
/*

C function

pango_language_includes_script
*/
func (recv *Language) IncludesScript(script Script) bool {
	c_script := (C.PangoScript)(script)

	retC := C.pango_language_includes_script((*C.PangoLanguage)(recv.native), c_script)
	retGo := retC == C.TRUE

	return retGo
}

// Frees a #PangoScriptIter created with pango_script_iter_new().
/*

C function

pango_script_iter_free
*/
func (recv *ScriptIter) Free() {
	C.pango_script_iter_free((*C.PangoScriptIter)(recv.native))

	return
}

// Unsupported : pango_script_iter_get_range : unsupported parameter script : PangoScript* with indirection level of 1

// Advances a #PangoScriptIter to the next range. If @iter
// is already at the end, it is left unchanged and %FALSE
// is returned.
/*

C function

pango_script_iter_next
*/
func (recv *ScriptIter) Next() bool {
	retC := C.pango_script_iter_next((*C.PangoScriptIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
