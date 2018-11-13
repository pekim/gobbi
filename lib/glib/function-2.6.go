// This is a generated file - DO NOT EDIT
// +build glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Checks that the GLib library in use is compatible with the
// given version. Generally you would pass in the constants
// #GLIB_MAJOR_VERSION, #GLIB_MINOR_VERSION, #GLIB_MICRO_VERSION
// as the three arguments to this function; that produces
// a check that the library in use is compatible with
// the version of GLib the application or module was compiled
// against.
//
// Compatibility is defined by two things: first the version
// of the running library is newer than the version
// @required_major.required_minor.@required_micro. Second
// the running library must be binary compatible with the
// version @required_major.required_minor.@required_micro
// (same major version.)
/*

C function : glib_check_version
*/
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
	c_required_major := (C.guint)(requiredMajor)

	c_required_minor := (C.guint)(requiredMinor)

	c_required_micro := (C.guint)(requiredMicro)

	retC := C.glib_check_version(c_required_major, c_required_minor, c_required_micro)
	retGo := C.GoString(retC)

	return retGo
}

// Returns the display basename for the particular filename, guaranteed
// to be valid UTF-8. The display name might not be identical to the filename,
// for instance there might be problems converting it to UTF-8, and some files
// can be translated in the display.
//
// If GLib cannot make sense of the encoding of @filename, as a last resort it
// replaces unknown characters with U+FFFD, the Unicode replacement character.
// You can search the result for the UTF-8 encoding of this character (which is
// "\357\277\275" in octal notation) to find out if @filename was in an invalid
// encoding.
//
// You must pass the whole absolute pathname to this functions so that
// translation of well known locations can be done.
//
// This function is preferred over g_filename_display_name() if you know the
// whole path, as it allows translation.
/*

C function : g_filename_display_basename
*/
func FilenameDisplayBasename(filename string) string {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_filename_display_basename(c_filename)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Converts a filename into a valid UTF-8 string. The conversion is
// not necessarily reversible, so you should keep the original around
// and use the return value of this function only for display purposes.
// Unlike g_filename_to_utf8(), the result is guaranteed to be non-%NULL
// even if the filename actually isn't in the GLib file name encoding.
//
// If GLib cannot make sense of the encoding of @filename, as a last resort it
// replaces unknown characters with U+FFFD, the Unicode replacement character.
// You can search the result for the UTF-8 encoding of this character (which is
// "\357\277\275" in octal notation) to find out if @filename was in an invalid
// encoding.
//
// If you know the whole pathname of the file you should use
// g_filename_display_basename(), since that allows location-based
// translation of filenames.
/*

C function : g_filename_display_name
*/
func FilenameDisplayName(filename string) string {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_filename_display_name(c_filename)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_get_filename_charsets : unsupported parameter charsets : in string with indirection level of 3

// Unsupported : g_get_language_names : no return type

// Unsupported : g_get_system_config_dirs : no return type

// Unsupported : g_get_system_data_dirs : no return type

// Returns a base directory in which to store non-essential, cached
// data specific to particular user.
//
// On UNIX platforms this is determined using the mechanisms described
// in the
// [XDG Base Directory Specification](http://www.freedesktop.org/Standards/basedir-spec).
// In this case the directory retrieved will be `XDG_CACHE_HOME`.
//
// On Windows it follows XDG Base Directory Specification if `XDG_CACHE_HOME` is defined.
// If `XDG_CACHE_HOME` is undefined, the directory that serves as a common
// repository for temporary Internet files is used instead. A typical path is
// `C:\Documents and Settings\username\Local Settings\Temporary Internet Files`.
// See the [documentation for `CSIDL_INTERNET_CACHE`](https://msdn.microsoft.com/en-us/library/windows/desktop/bb762494%28v=vs.85%29.aspx#csidl_internet_cache).
/*

C function : g_get_user_cache_dir
*/
func GetUserCacheDir() string {
	retC := C.g_get_user_cache_dir()
	retGo := C.GoString(retC)

	return retGo
}

// Returns a base directory in which to store user-specific application
// configuration information such as user preferences and settings.
//
// On UNIX platforms this is determined using the mechanisms described
// in the
// [XDG Base Directory Specification](http://www.freedesktop.org/Standards/basedir-spec).
// In this case the directory retrieved will be `XDG_CONFIG_HOME`.
//
// On Windows it follows XDG Base Directory Specification if `XDG_CONFIG_HOME` is defined.
// If `XDG_CONFIG_HOME` is undefined, the folder to use for local (as opposed
// to roaming) application data is used instead. See the
// [documentation for `CSIDL_LOCAL_APPDATA`](https://msdn.microsoft.com/en-us/library/windows/desktop/bb762494%28v=vs.85%29.aspx#csidl_local_appdata).
// Note that in this case on Windows it will be  the same
// as what g_get_user_data_dir() returns.
/*

C function : g_get_user_config_dir
*/
func GetUserConfigDir() string {
	retC := C.g_get_user_config_dir()
	retGo := C.GoString(retC)

	return retGo
}

// Returns a base directory in which to access application data such
// as icons that is customized for a particular user.
//
// On UNIX platforms this is determined using the mechanisms described
// in the
// [XDG Base Directory Specification](http://www.freedesktop.org/Standards/basedir-spec).
// In this case the directory retrieved will be `XDG_DATA_HOME`.
//
// On Windows it follows XDG Base Directory Specification if `XDG_DATA_HOME`
// is defined. If `XDG_DATA_HOME` is undefined, the folder to use for local (as
// opposed to roaming) application data is used instead. See the
// [documentation for `CSIDL_LOCAL_APPDATA`](https://msdn.microsoft.com/en-us/library/windows/desktop/bb762494%28v=vs.85%29.aspx#csidl_local_appdata).
// Note that in this case on Windows it will be the same
// as what g_get_user_config_dir() returns.
/*

C function : g_get_user_data_dir
*/
func GetUserDataDir() string {
	retC := C.g_get_user_data_dir()
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_log_set_default_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// A wrapper for the POSIX rmdir() function. The rmdir() function
// deletes a directory from the filesystem.
//
// See your C library manual for more details about how rmdir() works
// on your system.
/*

C function : g_rmdir
*/
func Rmdir(filename string) int32 {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_rmdir(c_filename)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_strv_length : unsupported parameter str_array : in string with indirection level of 2

// A wrapper for the POSIX unlink() function. The unlink() function
// deletes a name from the filesystem. If this was the last link to the
// file and no processes have it opened, the diskspace occupied by the
// file is freed.
//
// See your C library manual for more details about unlink(). Note
// that on Windows, it is in general not possible to delete files that
// are open to some process, or mapped into memory.
/*

C function : g_unlink
*/
func Unlink(filename string) int32 {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_unlink(c_filename)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_uri_list_extract_uris : no return type
