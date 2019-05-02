// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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
