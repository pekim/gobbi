// This is a generated file - DO NOT EDIT
// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Blacklisted : g_atomic_int_add

// Blacklisted : g_atomic_int_compare_and_exchange

// Blacklisted : g_atomic_int_dec_and_test

// Blacklisted : g_atomic_int_exchange_and_add

// Blacklisted : g_atomic_int_get

// Blacklisted : g_atomic_int_inc

// Blacklisted : g_atomic_int_set

// Unsupported : g_atomic_pointer_compare_and_exchange : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_get : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_set : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_child_watch_add : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Unsupported : g_child_watch_add_full : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Creates a new child_watch source.
//
// The source will not initially be associated with any #GMainContext
// and must be added to one with g_source_attach() before it will be
// executed.
//
// Note that child watch sources can only be used in conjunction with
// `g_spawn...` when the %G_SPAWN_DO_NOT_REAP_CHILD flag is used.
//
// Note that on platforms where #GPid must be explicitly closed
// (see g_spawn_close_pid()) @pid must not be closed while the
// source is still active. Typically, you will want to call
// g_spawn_close_pid() in the callback function for the source.
//
// On POSIX platforms, the following restrictions apply to this API
// due to limitations in POSIX process interfaces:
//
// * @pid must be a child of this process
// * @pid must be positive
// * the application must not call `waitpid` with a non-positive
// first argument, for instance in another thread
// * the application must not wait for @pid to exit by any other
// mechanism, including `waitpid(pid, ...)` or a second child-watch
// source for the same @pid
// * the application must not ignore SIGCHILD
//
// If any of those conditions are not met, this and related APIs will
// not work correctly. This can often be diagnosed via a GLib warning
// stating that `ECHILD` was received by `waitpid`.
//
// Calling `waitpid` for specific processes other than @pid remains a
// valid thing to do.
/*

C function : g_child_watch_source_new
*/
func ChildWatchSourceNew(pid Pid) *Source {
	c_pid := (C.GPid)(pid)

	retC := C.g_child_watch_source_new(c_pid)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Reads the contents of the symbolic link @filename like the POSIX
// readlink() function.  The returned string is in the encoding used
// for filenames. Use g_filename_to_utf8() to convert it to UTF-8.
/*

C function : g_file_read_link
*/
func FileReadLink(filename string) (string, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_file_read_link(c_filename, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_markup_printf_escaped : unsupported parameter ... : varargs

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no type generator for va_list (va_list) for param args

// Sets an environment variable. On UNIX, both the variable's name and
// value can be arbitrary byte strings, except that the variable's name
// cannot contain '='. On Windows, they should be in UTF-8.
//
// Note that on some systems, when variables are overwritten, the memory
// used for the previous variables and its value isn't reclaimed.
//
// You should be mindful of the fact that environment variable handling
// in UNIX is not thread-safe, and your program may crash if one thread
// calls g_setenv() while another thread is calling getenv(). (And note
// that many functions, such as gettext(), call getenv() internally.)
// This function is only safe to use at the very start of your program,
// before creating any other threads (or creating objects that create
// worker threads of their own).
//
// If you need to set up the environment for a child process, you can
// use g_get_environ() to get an environment array, modify that with
// g_environ_setenv() and g_environ_unsetenv(), and then pass that
// array directly to execvpe(), g_spawn_async(), or the like.
/*

C function : g_setenv
*/
func Setenv(variable string, value string, overwrite bool) bool {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_overwrite :=
		boolToGboolean(overwrite)

	retC := C.g_setenv(c_variable, c_value, c_overwrite)
	retGo := retC == C.TRUE

	return retGo
}

// An auxiliary function for gettext() support (see Q_()).
/*

C function : g_strip_context
*/
func StripContext(msgid string, msgval string) string {
	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_msgval := C.CString(msgval)
	defer C.free(unsafe.Pointer(c_msgval))

	retC := C.g_strip_context(c_msgid, c_msgval)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_strsplit_set : no return type

// In Unicode, some characters are "mirrored". This means that their
// images are mirrored horizontally in text that is laid out from right
// to left. For instance, "(" would become its mirror image, ")", in
// right-to-left text.
//
// If @ch has the Unicode mirrored property and there is another unicode
// character that typically has a glyph that is the mirror image of @ch's
// glyph and @mirrored_ch is set, it puts that character in the address
// pointed to by @mirrored_ch.  Otherwise the original character is put.
/*

C function : g_unichar_get_mirror_char
*/
func UnicharGetMirrorChar(ch rune, mirroredCh rune) bool {
	c_ch := (C.gunichar)(ch)

	c_mirrored_ch := (C.gunichar)(mirroredCh)

	retC := C.g_unichar_get_mirror_char(c_ch, &c_mirrored_ch)
	retGo := retC == C.TRUE

	return retGo
}

// Removes an environment variable from the environment.
//
// Note that on some systems, when variables are overwritten, the
// memory used for the previous variables and its value isn't reclaimed.
//
// You should be mindful of the fact that environment variable handling
// in UNIX is not thread-safe, and your program may crash if one thread
// calls g_unsetenv() while another thread is calling getenv(). (And note
// that many functions, such as gettext(), call getenv() internally.) This
// function is only safe to use at the very start of your program, before
// creating any other threads (or creating objects that create worker
// threads of their own).
//
// If you need to set up the environment for a child process, you can
// use g_get_environ() to get an environment array, modify that with
// g_environ_setenv() and g_environ_unsetenv(), and then pass that
// array directly to execvpe(), g_spawn_async(), or the like.
/*

C function : g_unsetenv
*/
func Unsetenv(variable string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	C.g_unsetenv(c_variable)

	return
}

// Unsupported : g_vasprintf : unsupported parameter string : in string with indirection level of 2
