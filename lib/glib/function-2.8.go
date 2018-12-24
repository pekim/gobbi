// This is a generated file - DO NOT EDIT
// +build glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Access is a wrapper around the C function g_access.
func Access(filename string, mode int32) int32 {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_mode := (C.int)(mode)

	retC := C.g_access(c_filename, c_mode)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_build_filenamev : unsupported parameter args :

// Unsupported : g_build_pathv : unsupported parameter args :

// Chdir is a wrapper around the C function g_chdir.
func Chdir(path string) int32 {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_chdir(c_path)
	retGo := (int32)(retC)

	return retGo
}

// DatalistGetFlags is a wrapper around the C function g_datalist_get_flags.
func DatalistGetFlags(datalist *Data) uint32 {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	retC := C.g_datalist_get_flags(c_datalist)
	retGo := (uint32)(retC)

	return retGo
}

// DatalistSetFlags is a wrapper around the C function g_datalist_set_flags.
func DatalistSetFlags(datalist *Data, flags uint32) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_flags := (C.guint)(flags)

	C.g_datalist_set_flags(c_datalist, c_flags)

	return
}

// DatalistUnsetFlags is a wrapper around the C function g_datalist_unset_flags.
func DatalistUnsetFlags(datalist *Data, flags uint32) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_flags := (C.guint)(flags)

	C.g_datalist_unset_flags(c_datalist, c_flags)

	return
}

// FileSetContents is a wrapper around the C function g_file_set_contents.
func FileSetContents(filename string, contents []uint8) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_contents := &contents[0]

	c_length := (C.gssize)(len(contents))

	var cThrowableError *C.GError

	retC := C.g_file_set_contents(c_filename, (*C.gchar)(unsafe.Pointer(c_contents)), c_length, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetHostName is a wrapper around the C function g_get_host_name.
func GetHostName() string {
	retC := C.g_get_host_name()
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_listenv : no return type

// MkdirWithParents is a wrapper around the C function g_mkdir_with_parents.
func MkdirWithParents(pathname string, mode int32) int32 {
	c_pathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(c_pathname))

	c_mode := (C.gint)(mode)

	retC := C.g_mkdir_with_parents(c_pathname, c_mode)
	retGo := (int32)(retC)

	return retGo
}

// TryMalloc0 is a wrapper around the C function g_try_malloc0.
func TryMalloc0(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_malloc0(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Utf8CollateKeyForFilename is a wrapper around the C function g_utf8_collate_key_for_filename.
func Utf8CollateKeyForFilename(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_collate_key_for_filename(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
