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

// A wrapper for the POSIX access() function. This function is used to
// test a pathname for one or several of read, write or execute
// permissions, or just existence.
//
// On Windows, the file protection mechanism is not at all POSIX-like,
// and the underlying function in the C library only checks the
// FAT-style READONLY attribute, and does not look at the ACL of a
// file at all. This function is this in practise almost useless on
// Windows. Software that needs to handle file permissions on Windows
// more exactly should use the Win32 API.
//
// See your C library manual for more details about access().
/*

C function : g_access
*/
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

// A wrapper for the POSIX chdir() function. The function changes the
// current directory of the process to @path.
//
// See your C library manual for more details about chdir().
/*

C function : g_chdir
*/
func Chdir(path string) int32 {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_chdir(c_path)
	retGo := (int32)(retC)

	return retGo
}

// Gets flags values packed in together with the datalist.
// See g_datalist_set_flags().
/*

C function : g_datalist_get_flags
*/
func DatalistGetFlags(datalist *Data) uint32 {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	retC := C.g_datalist_get_flags(c_datalist)
	retGo := (uint32)(retC)

	return retGo
}

// Turns on flag values for a data list. This function is used
// to keep a small number of boolean flags in an object with
// a data list without using any additional space. It is
// not generally useful except in circumstances where space
// is very tight. (It is used in the base #GObject type, for
// example.)
/*

C function : g_datalist_set_flags
*/
func DatalistSetFlags(datalist *Data, flags uint32) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_flags := (C.guint)(flags)

	C.g_datalist_set_flags(c_datalist, c_flags)

	return
}

// Turns off flag values for a data list. See g_datalist_unset_flags()
/*

C function : g_datalist_unset_flags
*/
func DatalistUnsetFlags(datalist *Data, flags uint32) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_flags := (C.guint)(flags)

	C.g_datalist_unset_flags(c_datalist, c_flags)

	return
}

// Writes all of @contents to a file named @filename, with good error checking.
// If a file called @filename already exists it will be overwritten.
//
// This write is atomic in the sense that it is first written to a temporary
// file which is then renamed to the final name. Notes:
//
// - On UNIX, if @filename already exists hard links to @filename will break.
// Also since the file is recreated, existing permissions, access control
// lists, metadata etc. may be lost. If @filename is a symbolic link,
// the link itself will be replaced, not the linked file.
//
// - On Windows renaming a file will not remove an existing file with the
// new name, so on Windows there is a race condition between the existing
// file being removed and the temporary file being renamed.
//
// - On Windows there is no way to remove a file that is open to some
// process, or mapped into memory. Thus, this function will fail if
// @filename already exists and is open.
//
// If the call was successful, it returns %TRUE. If the call was not successful,
// it returns %FALSE and sets @error. The error domain is #G_FILE_ERROR.
// Possible error codes are those in the #GFileError enumeration.
//
// Note that the name for the temporary file is constructed by appending up
// to 7 characters to @filename.
/*

C function : g_file_set_contents
*/
func FileSetContents(filename string, contents []uint8) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_contents := &contents[0]

	c_length := (C.gssize)(len(contents))

	var cThrowableError *C.GError

	retC := C.g_file_set_contents(c_filename, (*C.gchar)(unsafe.Pointer(c_contents)), c_length, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Return a name for the machine.
//
// The returned name is not necessarily a fully-qualified domain name,
// or even present in DNS or some other name service at all. It need
// not even be unique on your local network or site, but usually it
// is. Callers should not rely on the return value having any specific
// properties like uniqueness for security purposes. Even if the name
// of the machine is changed while an application is running, the
// return value from this function does not change. The returned
// string is owned by GLib and should not be modified or freed. If no
// name can be determined, a default fixed string "localhost" is
// returned.
//
// The encoding of the returned string is UTF-8.
/*

C function : g_get_host_name
*/
func GetHostName() string {
	retC := C.g_get_host_name()
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_listenv : no return type

// Create a directory if it doesn't already exist. Create intermediate
// parent directories as needed, too.
/*

C function : g_mkdir_with_parents
*/
func MkdirWithParents(pathname string, mode int32) int32 {
	c_pathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(c_pathname))

	c_mode := (C.gint)(mode)

	retC := C.g_mkdir_with_parents(c_pathname, c_mode)
	retGo := (int32)(retC)

	return retGo
}

// Attempts to allocate @n_bytes, initialized to 0's, and returns %NULL on
// failure. Contrast with g_malloc0(), which aborts the program on failure.
/*

C function : g_try_malloc0
*/
func TryMalloc0(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_malloc0(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Converts a string into a collation key that can be compared
// with other collation keys produced by the same function using strcmp().
//
// In order to sort filenames correctly, this function treats the dot '.'
// as a special case. Most dictionary orderings seem to consider it
// insignificant, thus producing the ordering "event.c" "eventgenerator.c"
// "event.h" instead of "event.c" "event.h" "eventgenerator.c". Also, we
// would like to treat numbers intelligently so that "file1" "file10" "file5"
// is sorted as "file1" "file5" "file10".
//
// Note that this function depends on the [current locale][setlocale].
/*

C function : g_utf8_collate_key_for_filename
*/
func Utf8CollateKeyForFilename(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_collate_key_for_filename(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
