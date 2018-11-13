// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Returns the full path of a special directory using its logical id.
//
// On UNIX this is done using the XDG special user directories.
// For compatibility with existing practise, %G_USER_DIRECTORY_DESKTOP
// falls back to `$HOME/Desktop` when XDG special user directories have
// not been set up.
//
// Depending on the platform, the user might be able to change the path
// of the special directory without requiring the session to restart; GLib
// will not reflect any change once the special directories are loaded.
/*

C function : g_get_user_special_dir
*/
func GetUserSpecialDir(directory UserDirectory) string {
	c_directory := (C.GUserDirectory)(directory)

	retC := C.g_get_user_special_dir(c_directory)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_once_init_enter : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_once_init_leave : unsupported parameter location : no type generator for gpointer (void*) for param location

// Checks whether @replacement is a valid replacement string
// (see g_regex_replace()), i.e. that all escape sequences in
// it are valid.
//
// If @has_references is not %NULL then @replacement is checked
// for pattern references. For instance, replacement text 'foo\n'
// does not contain references and may be evaluated without information
// about actual match, but '\0\1' (whole match followed by first
// subpattern) requires valid #GMatchInfo object.
/*

C function : g_regex_check_replacement
*/
func RegexCheckReplacement(replacement string) (bool, bool, error) {
	c_replacement := C.CString(replacement)
	defer C.free(unsafe.Pointer(c_replacement))

	var c_has_references C.gboolean

	var cThrowableError *C.GError

	retC := C.g_regex_check_replacement(c_replacement, &c_has_references, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	hasReferences := c_has_references == C.TRUE

	return retGo, hasReferences, goThrowableError
}

// Unsupported : g_regex_escape_string : unsupported parameter string :

// Scans for a match in @string for @pattern.
//
// This function is equivalent to g_regex_match() but it does not
// require to compile the pattern with g_regex_new(), avoiding some
// lines of code when you need just to do a match without extracting
// substrings, capture counts, and so on.
//
// If this function is to be called on the same @pattern more than
// once, it's more efficient to compile the pattern once with
// g_regex_new() and then use g_regex_match().
/*

C function : g_regex_match_simple
*/
func RegexMatchSimple(pattern string, string string, compileOptions RegexCompileFlags, matchOptions RegexMatchFlags) bool {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_compile_options := (C.GRegexCompileFlags)(compileOptions)

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	retC := C.g_regex_match_simple(c_pattern, c_string, c_compile_options, c_match_options)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_regex_split_simple : no return type

// Returns the data that @iter points to.
/*

C function : g_sequence_get
*/
func SequenceGet(iter *SequenceIter) uintptr {
	c_iter := (*C.GSequenceIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GSequenceIter)(iter.ToC())
	}

	retC := C.g_sequence_get(c_iter)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Inserts a new item just before the item pointed to by @iter.
/*

C function : g_sequence_insert_before
*/
func SequenceInsertBefore(iter *SequenceIter, data uintptr) *SequenceIter {
	c_iter := (*C.GSequenceIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GSequenceIter)(iter.ToC())
	}

	c_data := (C.gpointer)(data)

	retC := C.g_sequence_insert_before(c_iter, c_data)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Moves the item pointed to by @src to the position indicated by @dest.
// After calling this function @dest will point to the position immediately
// after @src. It is allowed for @src and @dest to point into different
// sequences.
/*

C function : g_sequence_move
*/
func SequenceMove(src *SequenceIter, dest *SequenceIter) {
	c_src := (*C.GSequenceIter)(C.NULL)
	if src != nil {
		c_src = (*C.GSequenceIter)(src.ToC())
	}

	c_dest := (*C.GSequenceIter)(C.NULL)
	if dest != nil {
		c_dest = (*C.GSequenceIter)(dest.ToC())
	}

	C.g_sequence_move(c_src, c_dest)

	return
}

// Inserts the (@begin, @end) range at the destination pointed to by ptr.
// The @begin and @end iters must point into the same sequence. It is
// allowed for @dest to point to a different sequence than the one pointed
// into by @begin and @end.
//
// If @dest is NULL, the range indicated by @begin and @end is
// removed from the sequence. If @dest iter points to a place within
// the (@begin, @end) range, the range does not move.
/*

C function : g_sequence_move_range
*/
func SequenceMoveRange(dest *SequenceIter, begin *SequenceIter, end *SequenceIter) {
	c_dest := (*C.GSequenceIter)(C.NULL)
	if dest != nil {
		c_dest = (*C.GSequenceIter)(dest.ToC())
	}

	c_begin := (*C.GSequenceIter)(C.NULL)
	if begin != nil {
		c_begin = (*C.GSequenceIter)(begin.ToC())
	}

	c_end := (*C.GSequenceIter)(C.NULL)
	if end != nil {
		c_end = (*C.GSequenceIter)(end.ToC())
	}

	C.g_sequence_move_range(c_dest, c_begin, c_end)

	return
}

// Finds an iterator somewhere in the range (@begin, @end). This
// iterator will be close to the middle of the range, but is not
// guaranteed to be exactly in the middle.
//
// The @begin and @end iterators must both point to the same sequence
// and @begin must come before or be equal to @end in the sequence.
/*

C function : g_sequence_range_get_midpoint
*/
func SequenceRangeGetMidpoint(begin *SequenceIter, end *SequenceIter) *SequenceIter {
	c_begin := (*C.GSequenceIter)(C.NULL)
	if begin != nil {
		c_begin = (*C.GSequenceIter)(begin.ToC())
	}

	c_end := (*C.GSequenceIter)(C.NULL)
	if end != nil {
		c_end = (*C.GSequenceIter)(end.ToC())
	}

	retC := C.g_sequence_range_get_midpoint(c_begin, c_end)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes the item pointed to by @iter. It is an error to pass the
// end iterator to this function.
//
// If the sequence has a data destroy function associated with it, this
// function is called on the data for the removed item.
/*

C function : g_sequence_remove
*/
func SequenceRemove(iter *SequenceIter) {
	c_iter := (*C.GSequenceIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GSequenceIter)(iter.ToC())
	}

	C.g_sequence_remove(c_iter)

	return
}

// Removes all items in the (@begin, @end) range.
//
// If the sequence has a data destroy function associated with it, this
// function is called on the data for the removed items.
/*

C function : g_sequence_remove_range
*/
func SequenceRemoveRange(begin *SequenceIter, end *SequenceIter) {
	c_begin := (*C.GSequenceIter)(C.NULL)
	if begin != nil {
		c_begin = (*C.GSequenceIter)(begin.ToC())
	}

	c_end := (*C.GSequenceIter)(C.NULL)
	if end != nil {
		c_end = (*C.GSequenceIter)(end.ToC())
	}

	C.g_sequence_remove_range(c_begin, c_end)

	return
}

// Changes the data for the item pointed to by @iter to be @data. If
// the sequence has a data destroy function associated with it, that
// function is called on the existing data that @iter pointed to.
/*

C function : g_sequence_set
*/
func SequenceSet(iter *SequenceIter, data uintptr) {
	c_iter := (*C.GSequenceIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GSequenceIter)(iter.ToC())
	}

	c_data := (C.gpointer)(data)

	C.g_sequence_set(c_iter, c_data)

	return
}

// Swaps the items pointed to by @a and @b. It is allowed for @a and @b
// to point into difference sequences.
/*

C function : g_sequence_swap
*/
func SequenceSwap(a *SequenceIter, b *SequenceIter) {
	c_a := (*C.GSequenceIter)(C.NULL)
	if a != nil {
		c_a = (*C.GSequenceIter)(a.ToC())
	}

	c_b := (*C.GSequenceIter)(C.NULL)
	if b != nil {
		c_b = (*C.GSequenceIter)(b.ToC())
	}

	C.g_sequence_swap(c_a, c_b)

	return
}

// Allocates a block of memory from the slice allocator
// and copies @block_size bytes into it from @mem_block.
//
// @mem_block must be non-%NULL if @block_size is non-zero.
/*

C function : g_slice_copy
*/
func SliceCopy(blockSize uint64, memBlock uintptr) uintptr {
	c_block_size := (C.gsize)(blockSize)

	c_mem_block := (C.gconstpointer)(memBlock)

	retC := C.g_slice_copy(c_block_size, c_mem_block)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Creates a new timeout source.
//
// The source will not initially be associated with any #GMainContext
// and must be added to one with g_source_attach() before it will be
// executed.
//
// The scheduling granularity/accuracy of this timeout source will be
// in seconds.
//
// The interval given in terms of monotonic time, not wall clock time.
// See g_get_monotonic_time().
/*

C function : g_timeout_source_new_seconds
*/
func TimeoutSourceNewSeconds(interval uint32) *Source {
	c_interval := (C.guint)(interval)

	retC := C.g_timeout_source_new_seconds(c_interval)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Determines the canonical combining class of a Unicode character.
/*

C function : g_unichar_combining_class
*/
func UnicharCombiningClass(uc rune) int32 {
	c_uc := (C.gunichar)(uc)

	retC := C.g_unichar_combining_class(c_uc)
	retGo := (int32)(retC)

	return retGo
}

// Looks up the #GUnicodeScript for a particular character (as defined
// by Unicode Standard Annex \#24). No check is made for @ch being a
// valid Unicode character; if you pass in invalid character, the
// result is undefined.
//
// This function is equivalent to pango_script_for_unichar() and the
// two are interchangeable.
/*

C function : g_unichar_get_script
*/
func UnicharGetScript(ch rune) UnicodeScript {
	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_get_script(c_ch)
	retGo := (UnicodeScript)(retC)

	return retGo
}

// Determines whether a character is a mark (non-spacing mark,
// combining mark, or enclosing mark in Unicode speak).
// Given some UTF-8 text, obtain a character value
// with g_utf8_get_char().
//
// Note: in most cases where isalpha characters are allowed,
// ismark characters should be allowed to as they are essential
// for writing most European languages as well as many non-Latin
// scripts.
/*

C function : g_unichar_ismark
*/
func UnicharIsmark(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_ismark(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Determines if a given character typically takes zero width when rendered.
// The return value is %TRUE for all non-spacing and enclosing marks
// (e.g., combining accents), format characters, zero-width
// space, but not U+00AD SOFT HYPHEN.
//
// A typical use of this function is with one of g_unichar_iswide() or
// g_unichar_iswide_cjk() to determine the number of cells a string occupies
// when displayed on a grid display (terminals).  However, note that not all
// terminals support zero-width rendering of zero-width marks.
/*

C function : g_unichar_iszerowidth
*/
func UnicharIszerowidth(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iszerowidth(c_c)
	retGo := retC == C.TRUE

	return retGo
}
