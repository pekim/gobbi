// This is a generated file - DO NOT EDIT
// +build glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Performs an atomic bitwise 'and' of the value of @atomic and @val,
// storing the result back in @atomic.
//
// This call acts as a full compiler and hardware memory barrier.
//
// Think of this operation as an atomic version of
// `{ tmp = *atomic; *atomic &= val; return tmp; }`.
/*

C function : g_atomic_int_and
*/
func AtomicIntAnd(atomic uint32, val uint32) uint32 {
	c_atomic := (C.guint)(atomic)

	c_val := (C.guint)(val)

	retC := C.g_atomic_int_and(&c_atomic, c_val)
	retGo := (uint32)(retC)

	return retGo
}

// Performs an atomic bitwise 'or' of the value of @atomic and @val,
// storing the result back in @atomic.
//
// Think of this operation as an atomic version of
// `{ tmp = *atomic; *atomic |= val; return tmp; }`.
//
// This call acts as a full compiler and hardware memory barrier.
/*

C function : g_atomic_int_or
*/
func AtomicIntOr(atomic uint32, val uint32) uint32 {
	c_atomic := (C.guint)(atomic)

	c_val := (C.guint)(val)

	retC := C.g_atomic_int_or(&c_atomic, c_val)
	retGo := (uint32)(retC)

	return retGo
}

// Performs an atomic bitwise 'xor' of the value of @atomic and @val,
// storing the result back in @atomic.
//
// Think of this operation as an atomic version of
// `{ tmp = *atomic; *atomic ^= val; return tmp; }`.
//
// This call acts as a full compiler and hardware memory barrier.
/*

C function : g_atomic_int_xor
*/
func AtomicIntXor(atomic uint32, val uint32) uint32 {
	c_atomic := (C.guint)(atomic)

	c_val := (C.guint)(val)

	retC := C.g_atomic_int_xor(&c_atomic, c_val)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_atomic_pointer_add : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_and : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_or : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_xor : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Computes the HMAC for a binary @data of @length. This is a
// convenience wrapper for g_hmac_new(), g_hmac_get_string()
// and g_hmac_unref().
//
// The hexadecimal string returned will be in lower case.
/*

C function : g_compute_hmac_for_data
*/
func ComputeHmacForData(digestType ChecksumType, key []uint8, data []uint8) string {
	c_digest_type := (C.GChecksumType)(digestType)

	c_key := &key[0]

	c_key_len := (C.gsize)(len(key))

	c_data := &data[0]

	c_length := (C.gsize)(len(data))

	retC := C.g_compute_hmac_for_data(c_digest_type, (*C.guchar)(unsafe.Pointer(c_key)), c_key_len, (*C.guchar)(unsafe.Pointer(c_data)), c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Computes the HMAC for a string.
//
// The hexadecimal string returned will be in lower case.
/*

C function : g_compute_hmac_for_string
*/
func ComputeHmacForString(digestType ChecksumType, key []uint8, str string) string {
	c_digest_type := (C.GChecksumType)(digestType)

	c_key := &key[0]

	c_key_len := (C.gsize)(len(key))

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_length := (C.gssize)(len(str))

	retC := C.g_compute_hmac_for_string(c_digest_type, (*C.guchar)(unsafe.Pointer(c_key)), c_key_len, c_str, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Creates a subdirectory in the preferred directory for temporary
// files (as returned by g_get_tmp_dir()).
//
// @tmpl should be a string in the GLib file name encoding containing
// a sequence of six 'X' characters, as the parameter to g_mkstemp().
// However, unlike these functions, the template should only be a
// basename, no directory components are allowed. If template is
// %NULL, a default template is used.
//
// Note that in contrast to g_mkdtemp() (and mkdtemp()) @tmpl is not
// modified, and might thus be a read-only literal string.
/*

C function : g_dir_make_tmp
*/
func DirMakeTmp(tmpl string) (string, error) {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	var cThrowableError *C.GError

	retC := C.g_dir_make_tmp(c_tmpl, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Formats a size (for example the size of a file) into a human readable
// string.  Sizes are rounded to the nearest size prefix (kB, MB, GB)
// and are displayed rounded to the nearest tenth. E.g. the file size
// 3292528 bytes will be converted into the string "3.2 MB".
//
// The prefix units base is 1000 (i.e. 1 kB is 1000 bytes).
//
// This string should be freed with g_free() when not needed any longer.
//
// See g_format_size_full() for more options about how the size might be
// formatted.
/*

C function : g_format_size
*/
func FormatSize(size uint64) string {
	c_size := (C.guint64)(size)

	retC := C.g_format_size(c_size)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Formats a size.
//
// This function is similar to g_format_size() but allows for flags
// that modify the output. See #GFormatSizeFlags.
/*

C function : g_format_size_full
*/
func FormatSizeFull(size uint64, flags FormatSizeFlags) string {
	c_size := (C.guint64)(size)

	c_flags := (C.GFormatSizeFlags)(flags)

	retC := C.g_format_size_full(c_size, c_flags)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Creates a temporary directory. See the mkdtemp() documentation
// on most UNIX-like systems.
//
// The parameter is a string that should follow the rules for
// mkdtemp() templates, i.e. contain the string "XXXXXX".
// g_mkdtemp() is slightly more flexible than mkdtemp() in that the
// sequence does not have to occur at the very end of the template.
// The X string will be modified to form the name of a directory that
// didn't exist.
// The string should be in the GLib file name encoding. Most importantly,
// on Windows it should be in UTF-8.
//
// If you are going to be creating a temporary directory inside the
// directory returned by g_get_tmp_dir(), you might want to use
// g_dir_make_tmp() instead.
/*

C function : g_mkdtemp
*/
func Mkdtemp(tmpl string) string {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	retC := C.g_mkdtemp(c_tmpl)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Creates a temporary directory. See the mkdtemp() documentation
// on most UNIX-like systems.
//
// The parameter is a string that should follow the rules for
// mkdtemp() templates, i.e. contain the string "XXXXXX".
// g_mkdtemp_full() is slightly more flexible than mkdtemp() in that the
// sequence does not have to occur at the very end of the template
// and you can pass a @mode. The X string will be modified to form
// the name of a directory that didn't exist. The string should be
// in the GLib file name encoding. Most importantly, on Windows it
// should be in UTF-8.
//
// If you are going to be creating a temporary directory inside the
// directory returned by g_get_tmp_dir(), you might want to use
// g_dir_make_tmp() instead.
/*

C function : g_mkdtemp_full
*/
func MkdtempFull(tmpl string, mode int32) string {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	c_mode := (C.gint)(mode)

	retC := C.g_mkdtemp_full(c_tmpl, c_mode)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_pointer_bit_lock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_trylock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_unlock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Escapes the nul characters in @string to "\x00".  It can be used
// to compile a regex with embedded nul characters.
//
// For completeness, @length can be -1 for a nul-terminated string.
// In this case the output string will be of course equal to @string.
/*

C function : g_regex_escape_nul
*/
func RegexEscapeNul(string string, length int32) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gint)(length)

	retC := C.g_regex_escape_nul(c_string, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Indicates that a test failed. This function can be called
// multiple times from the same test. You can use this function
// if your test failed in a recoverable way.
//
// Do not use this function if the failure of a test could cause
// other tests to malfunction.
//
// Calling this function will not stop the test from running, you
// need to return from the test function yourself. So you can
// produce additional diagnostic messages or even continue running
// the test.
//
// If not called from inside a test, this function does nothing.
/*

C function : g_test_fail
*/
func TestFail() {
	C.g_test_fail()

	return
}

// Performs a single composition step of the
// Unicode canonical composition algorithm.
//
// This function includes algorithmic Hangul Jamo composition,
// but it is not exactly the inverse of g_unichar_decompose().
// No composition can have either of @a or @b equal to zero.
// To be precise, this function composes if and only if
// there exists a Primary Composite P which is canonically
// equivalent to the sequence <@a,@b>.  See the Unicode
// Standard for the definition of Primary Composite.
//
// If @a and @b do not compose a new character, @ch is set to zero.
//
// See
// [UAX#15](http://unicode.org/reports/tr15/)
// for details.
/*

C function : g_unichar_compose
*/
func UnicharCompose(a rune, b rune, ch rune) bool {
	c_a := (C.gunichar)(a)

	c_b := (C.gunichar)(b)

	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_compose(c_a, c_b, &c_ch)
	retGo := retC == C.TRUE

	return retGo
}

// Performs a single decomposition step of the
// Unicode canonical decomposition algorithm.
//
// This function does not include compatibility
// decompositions. It does, however, include algorithmic
// Hangul Jamo decomposition, as well as 'singleton'
// decompositions which replace a character by a single
// other character. In the case of singletons *@b will
// be set to zero.
//
// If @ch is not decomposable, *@a is set to @ch and *@b
// is set to zero.
//
// Note that the way Unicode decomposition pairs are
// defined, it is guaranteed that @b would not decompose
// further, but @a may itself decompose.  To get the full
// canonical decomposition for @ch, one would need to
// recursively call this function on @a.  Or use
// g_unichar_fully_decompose().
//
// See
// [UAX#15](http://unicode.org/reports/tr15/)
// for details.
/*

C function : g_unichar_decompose
*/
func UnicharDecompose(ch rune, a rune, b rune) bool {
	c_ch := (C.gunichar)(ch)

	c_a := (C.gunichar)(a)

	c_b := (C.gunichar)(b)

	retC := C.g_unichar_decompose(c_ch, &c_a, &c_b)
	retGo := retC == C.TRUE

	return retGo
}

// Computes the canonical or compatibility decomposition of a
// Unicode character.  For compatibility decomposition,
// pass %TRUE for @compat; for canonical decomposition
// pass %FALSE for @compat.
//
// The decomposed sequence is placed in @result.  Only up to
// @result_len characters are written into @result.  The length
// of the full decomposition (irrespective of @result_len) is
// returned by the function.  For canonical decomposition,
// currently all decompositions are of length at most 4, but
// this may change in the future (very unlikely though).
// At any rate, Unicode does guarantee that a buffer of length
// 18 is always enough for both compatibility and canonical
// decompositions, so that is the size recommended. This is provided
// as %G_UNICHAR_MAX_DECOMPOSITION_LENGTH.
//
// See
// [UAX#15](http://unicode.org/reports/tr15/)
// for details.
/*

C function : g_unichar_fully_decompose
*/
func UnicharFullyDecompose(ch rune, compat bool, result rune, resultLen uint64) uint64 {
	c_ch := (C.gunichar)(ch)

	c_compat :=
		boolToGboolean(compat)

	c_result := (C.gunichar)(result)

	c_result_len := (C.gsize)(resultLen)

	retC := C.g_unichar_fully_decompose(c_ch, c_compat, &c_result, c_result_len)
	retGo := (uint64)(retC)

	return retGo
}

// Looks up the Unicode script for @iso15924.  ISO 15924 assigns four-letter
// codes to scripts.  For example, the code for Arabic is 'Arab'.
// This function accepts four letter codes encoded as a @guint32 in a
// big-endian fashion.  That is, the code expected for Arabic is
// 0x41726162 (0x41 is ASCII code for 'A', 0x72 is ASCII code for 'r', etc).
//
// See
// [Codes for the representation of names of scripts](http://unicode.org/iso15924/codelists.html)
// for details.
/*

C function : g_unicode_script_from_iso15924
*/
func UnicodeScriptFromIso15924(iso15924 uint32) UnicodeScript {
	c_iso15924 := (C.guint32)(iso15924)

	retC := C.g_unicode_script_from_iso15924(c_iso15924)
	retGo := (UnicodeScript)(retC)

	return retGo
}

// Looks up the ISO 15924 code for @script.  ISO 15924 assigns four-letter
// codes to scripts.  For example, the code for Arabic is 'Arab'.  The
// four letter codes are encoded as a @guint32 by this function in a
// big-endian fashion.  That is, the code returned for Arabic is
// 0x41726162 (0x41 is ASCII code for 'A', 0x72 is ASCII code for 'r', etc).
//
// See
// [Codes for the representation of names of scripts](http://unicode.org/iso15924/codelists.html)
// for details.
/*

C function : g_unicode_script_to_iso15924
*/
func UnicodeScriptToIso15924(script UnicodeScript) uint32 {
	c_script := (C.GUnicodeScript)(script)

	retC := C.g_unicode_script_to_iso15924(c_script)
	retGo := (uint32)(retC)

	return retGo
}

// Similar to the UNIX pipe() call, but on modern systems like Linux
// uses the pipe2() system call, which atomically creates a pipe with
// the configured flags. The only supported flag currently is
// %FD_CLOEXEC. If for example you want to configure %O_NONBLOCK, that
// must still be done separately with fcntl().
//
// This function does not take %O_CLOEXEC, it takes %FD_CLOEXEC as if
// for fcntl(); these are different on Linux/glibc.
/*

C function : g_unix_open_pipe
*/
func UnixOpenPipe(fds int32, flags int32) (bool, error) {
	c_fds := (C.gint)(fds)

	c_flags := (C.gint)(flags)

	var cThrowableError *C.GError

	retC := C.g_unix_open_pipe(&c_fds, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Control the non-blocking state of the given file descriptor,
// according to @nonblock. On most systems this uses %O_NONBLOCK, but
// on some older ones may use %O_NDELAY.
/*

C function : g_unix_set_fd_nonblocking
*/
func UnixSetFdNonblocking(fd int32, nonblock bool) (bool, error) {
	c_fd := (C.gint)(fd)

	c_nonblock :=
		boolToGboolean(nonblock)

	var cThrowableError *C.GError

	retC := C.g_unix_set_fd_nonblocking(c_fd, c_nonblock, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_unix_signal_add : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// Unsupported : g_unix_signal_add_full : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// Create a #GSource that will be dispatched upon delivery of the UNIX
// signal @signum.  In GLib versions before 2.36, only `SIGHUP`, `SIGINT`,
// `SIGTERM` can be monitored.  In GLib 2.36, `SIGUSR1` and `SIGUSR2`
// were added. In GLib 2.54, `SIGWINCH` was added.
//
// Note that unlike the UNIX default, all sources which have created a
// watch will be dispatched, regardless of which underlying thread
// invoked g_unix_signal_source_new().
//
// For example, an effective use of this function is to handle `SIGTERM`
// cleanly; flushing any outstanding files, and then calling
// g_main_loop_quit ().  It is not safe to do any of this a regular
// UNIX signal handler; your handler may be invoked while malloc() or
// another library function is running, causing reentrancy if you
// attempt to use it from the handler.  None of the GLib/GObject API
// is safe against this kind of reentrancy.
//
// The interaction of this source when combined with native UNIX
// functions like sigprocmask() is not defined.
//
// The source will not initially be associated with any #GMainContext
// and must be added to one with g_source_attach() before it will be
// executed.
/*

C function : g_unix_signal_source_new
*/
func UnixSignalSourceNew(signum int32) *Source {
	c_signum := (C.gint)(signum)

	retC := C.g_unix_signal_source_new(c_signum)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copies a substring out of a UTF-8 encoded string.
// The substring will contain @end_pos - @start_pos characters.
/*

C function : g_utf8_substring
*/
func Utf8Substring(str string, startPos int64, endPos int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_start_pos := (C.glong)(startPos)

	c_end_pos := (C.glong)(endPos)

	retC := C.g_utf8_substring(c_str, c_start_pos, c_end_pos)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
