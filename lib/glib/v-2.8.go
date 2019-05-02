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

// BuildFilenamev is a wrapper around the C function g_build_filenamev.
func BuildFilenamev(args []string) string {
	c_args_array := make([]*C.gchar, len(args)+1, len(args)+1)
	for i, item := range args {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_args_array[i] = c
	}
	c_args_array[len(args)] = nil
	c_args_arrayPtr := &c_args_array[0]
	c_args := (**C.gchar)(unsafe.Pointer(c_args_arrayPtr))

	retC := C.g_build_filenamev(c_args)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// BuildPathv is a wrapper around the C function g_build_pathv.
func BuildPathv(separator string, args []string) string {
	c_separator := C.CString(separator)
	defer C.free(unsafe.Pointer(c_separator))

	c_args_array := make([]*C.gchar, len(args)+1, len(args)+1)
	for i, item := range args {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_args_array[i] = c
	}
	c_args_array[len(args)] = nil
	c_args_arrayPtr := &c_args_array[0]
	c_args := (**C.gchar)(unsafe.Pointer(c_args_arrayPtr))

	retC := C.g_build_pathv(c_separator, c_args)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

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

	c_contents_array := make([]C.guint8, len(contents)+1, len(contents)+1)
	for i, item := range contents {
		c := (C.guint8)(item)
		c_contents_array[i] = c
	}
	c_contents_array[len(contents)] = 0
	c_contents_arrayPtr := &c_contents_array[0]
	c_contents := (*C.gchar)(unsafe.Pointer(c_contents_arrayPtr))

	c_length := (C.gssize)(len(contents))

	var cThrowableError *C.GError

	retC := C.g_file_set_contents(c_filename, c_contents, c_length, &cThrowableError)
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

// Listenv is a wrapper around the C function g_listenv.
func Listenv() []string {
	retC := C.g_listenv()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// MkdirWithParents is a wrapper around the C function g_mkdir_with_parents.
func MkdirWithParents(pathname string, mode int32) int32 {
	c_pathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(c_pathname))

	c_mode := (C.gint)(mode)

	retC := C.g_mkdir_with_parents(c_pathname, c_mode)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_try_malloc0 : no return generator

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

// MappedFileNew is a wrapper around the C function g_mapped_file_new.
func MappedFileNew(filename string, writable bool) (*MappedFile, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_writable :=
		boolToGboolean(writable)

	var cThrowableError *C.GError

	retC := C.g_mapped_file_new(c_filename, c_writable, &cThrowableError)
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Free is a wrapper around the C function g_mapped_file_free.
func (recv *MappedFile) Free() {
	C.g_mapped_file_free((*C.GMappedFile)(recv.native))

	return
}

// GetContents is a wrapper around the C function g_mapped_file_get_contents.
func (recv *MappedFile) GetContents() string {
	retC := C.g_mapped_file_get_contents((*C.GMappedFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetLength is a wrapper around the C function g_mapped_file_get_length.
func (recv *MappedFile) GetLength() uint64 {
	retC := C.g_mapped_file_get_length((*C.GMappedFile)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}
