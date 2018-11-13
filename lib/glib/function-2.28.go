// This is a generated file - DO NOT EDIT
// +build glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_get_environ : no return type

// Unsupported : g_get_locale_variants : no return type

// Queries the system monotonic time.
//
// The monotonic clock will always increase and doesn't suffer
// discontinuities when the user (or NTP) changes the system time.  It
// may or may not continue to tick during times where the machine is
// suspended.
//
// We try to use the clock that corresponds as closely as possible to
// the passage of time as measured by system calls such as poll() but it
// may not always be possible to do this.
/*

C function

g_get_monotonic_time
*/
func GetMonotonicTime() int64 {
	retC := C.g_get_monotonic_time()
	retGo := (int64)(retC)

	return retGo
}

// Queries the system wall-clock time.
//
// This call is functionally equivalent to g_get_current_time() except
// that the return value is often more convenient than dealing with a
// #GTimeVal.
//
// You should only use this call if you are actually interested in the real
// wall-clock time.  g_get_monotonic_time() is probably more useful for
// measuring intervals.
/*

C function

g_get_real_time
*/
func GetRealTime() int64 {
	retC := C.g_get_real_time()
	retGo := (int64)(retC)

	return retGo
}

// Returns a directory that is unique to the current user on the local
// system.
//
// This is determined using the mechanisms described
// in the
// [XDG Base Directory Specification](http://www.freedesktop.org/Standards/basedir-spec).
// This is the directory
// specified in the `XDG_RUNTIME_DIR` environment variable.
// In the case that this variable is not set, we return the value of
// g_get_user_cache_dir(), after verifying that it exists.
/*

C function

g_get_user_runtime_dir
*/
func GetUserRuntimeDir() string {
	retC := C.g_get_user_runtime_dir()
	retGo := C.GoString(retC)

	return retGo
}
