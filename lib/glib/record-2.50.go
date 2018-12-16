// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// LoadFromBytes is a wrapper around the C function g_key_file_load_from_bytes.
func (recv *KeyFile) LoadFromBytes(bytes *Bytes, flags KeyFileFlags) (bool, error) {
	c_bytes := (*C.GBytes)(C.NULL)
	if bytes != nil {
		c_bytes = (*C.GBytes)(bytes.ToC())
	}

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_bytes((*C.GKeyFile)(recv.native), c_bytes, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LogField is a wrapper around the C record GLogField.
type LogField struct {
	native *C.GLogField
	Key    string
	Value  uintptr
	Length int64
}

func LogFieldNewFromC(u unsafe.Pointer) *LogField {
	c := (*C.GLogField)(u)
	if c == nil {
		return nil
	}

	g := &LogField{
		Key:    C.GoString(c.key),
		Length: (int64)(c.length),
		Value:  (uintptr)(c.value),
		native: c,
	}

	return g
}

func (recv *LogField) ToC() unsafe.Pointer {
	recv.native.key =
		C.CString(recv.Key)
	recv.native.value =
		(C.gconstpointer)(recv.Value)
	recv.native.length =
		(C.gssize)(recv.Length)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LogField with another LogField, and returns true if they represent the same GObject.
func (recv *LogField) Equals(other *LogField) bool {
	return other.ToC() == recv.ToC()
}
