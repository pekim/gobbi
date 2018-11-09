// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// IncludesScript is a wrapper around the C function pango_language_includes_script.
func (recv *Language) IncludesScript(script Script) bool {
	c_script := (C.PangoScript)(script)

	retC := C.pango_language_includes_script((*C.PangoLanguage)(recv.native), c_script)
	retGo := retC == C.TRUE

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
