// This is a generated file - DO NOT EDIT
// +build glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Maps a file into memory. On UNIX, this is using the mmap() function.
//
// If @writable is %TRUE, the mapped buffer may be modified, otherwise
// it is an error to modify the mapped buffer. Modifications to the buffer
// are not visible to other processes mapping the same file, and are not
// written back to the file.
//
// Note that modifications of the underlying file might affect the contents
// of the #GMappedFile. Therefore, mapping should only be used if the file
// will not be modified, or if all modifications of the file are done
// atomically (e.g. using g_file_set_contents()).
//
// If @filename is the name of an empty, regular file, the function
// will successfully return an empty #GMappedFile. In other cases of
// size 0 (e.g. device files such as /dev/null), @error will be set
// to the #GFileError value #G_FILE_ERROR_INVAL.
/*

C function

g_mapped_file_new
*/
func MappedFileNew(filename string, writable bool) (*MappedFile, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_writable :=
		boolToGboolean(writable)

	var cThrowableError *C.GError

	retC := C.g_mapped_file_new(c_filename, c_writable, &cThrowableError)
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This call existed before #GMappedFile had refcounting and is currently
// exactly the same as g_mapped_file_unref().
/*

C function

g_mapped_file_free
*/
func (recv *MappedFile) Free() {
	C.g_mapped_file_free((*C.GMappedFile)(recv.native))

	return
}

// Returns the contents of a #GMappedFile.
//
// Note that the contents may not be zero-terminated,
// even if the #GMappedFile is backed by a text file.
//
// If the file is empty then %NULL is returned.
/*

C function

g_mapped_file_get_contents
*/
func (recv *MappedFile) GetContents() string {
	retC := C.g_mapped_file_get_contents((*C.GMappedFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns the length of the contents of a #GMappedFile.
/*

C function

g_mapped_file_get_length
*/
func (recv *MappedFile) GetLength() uint64 {
	retC := C.g_mapped_file_get_length((*C.GMappedFile)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}
