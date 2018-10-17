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

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_key_file_load_from_dirs : unsupported parameter search_dirs : no param type

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify, GDestroyNotify

// ExpandReferences is a wrapper around the C function g_match_info_expand_references.
func (recv *MatchInfo) ExpandReferences(stringToExpand string) (string, error) {
	c_string_to_expand := C.CString(stringToExpand)
	defer C.free(unsafe.Pointer(c_string_to_expand))

	var cThrowableError *C.GError

	retC := C.g_match_info_expand_references((*C.GMatchInfo)(recv.native), c_string_to_expand, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Fetch is a wrapper around the C function g_match_info_fetch.
func (recv *MatchInfo) Fetch(matchNum int32) string {
	c_match_num := (C.gint)(matchNum)

	retC := C.g_match_info_fetch((*C.GMatchInfo)(recv.native), c_match_num)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_match_info_fetch_all : no return type

// FetchNamed is a wrapper around the C function g_match_info_fetch_named.
func (recv *MatchInfo) FetchNamed(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_match_info_fetch_named((*C.GMatchInfo)(recv.native), c_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FetchNamedPos is a wrapper around the C function g_match_info_fetch_named_pos.
func (recv *MatchInfo) FetchNamedPos(name string) (bool, int32, int32) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var c_start_pos C.gint

	var c_end_pos C.gint

	retC := C.g_match_info_fetch_named_pos((*C.GMatchInfo)(recv.native), c_name, &c_start_pos, &c_end_pos)
	retGo := retC == C.TRUE

	startPos := (int32)(c_start_pos)

	endPos := (int32)(c_end_pos)

	return retGo, startPos, endPos
}

// FetchPos is a wrapper around the C function g_match_info_fetch_pos.
func (recv *MatchInfo) FetchPos(matchNum int32) (bool, int32, int32) {
	c_match_num := (C.gint)(matchNum)

	var c_start_pos C.gint

	var c_end_pos C.gint

	retC := C.g_match_info_fetch_pos((*C.GMatchInfo)(recv.native), c_match_num, &c_start_pos, &c_end_pos)
	retGo := retC == C.TRUE

	startPos := (int32)(c_start_pos)

	endPos := (int32)(c_end_pos)

	return retGo, startPos, endPos
}

// Free is a wrapper around the C function g_match_info_free.
func (recv *MatchInfo) Free() {
	C.g_match_info_free((*C.GMatchInfo)(recv.native))

	return
}

// GetMatchCount is a wrapper around the C function g_match_info_get_match_count.
func (recv *MatchInfo) GetMatchCount() int32 {
	retC := C.g_match_info_get_match_count((*C.GMatchInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRegex is a wrapper around the C function g_match_info_get_regex.
func (recv *MatchInfo) GetRegex() *Regex {
	retC := C.g_match_info_get_regex((*C.GMatchInfo)(recv.native))
	retGo := RegexNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetString is a wrapper around the C function g_match_info_get_string.
func (recv *MatchInfo) GetString() string {
	retC := C.g_match_info_get_string((*C.GMatchInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// IsPartialMatch is a wrapper around the C function g_match_info_is_partial_match.
func (recv *MatchInfo) IsPartialMatch() bool {
	retC := C.g_match_info_is_partial_match((*C.GMatchInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Matches is a wrapper around the C function g_match_info_matches.
func (recv *MatchInfo) Matches() bool {
	retC := C.g_match_info_matches((*C.GMatchInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Next is a wrapper around the C function g_match_info_next.
func (recv *MatchInfo) Next() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_match_info_next((*C.GMatchInfo)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetHelp is a wrapper around the C function g_option_context_get_help.
func (recv *OptionContext) GetHelp(mainHelp bool, group *OptionGroup) string {
	c_main_help :=
		boolToGboolean(mainHelp)

	c_group := (*C.GOptionGroup)(group.ToC())

	retC := C.g_option_context_get_help((*C.GOptionContext)(recv.native), c_main_help, c_group)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify, GDestroyNotify

// Clear is a wrapper around the C function g_queue_clear.
func (recv *Queue) Clear() {
	C.g_queue_clear((*C.GQueue)(recv.native))

	return
}

// Init is a wrapper around the C function g_queue_init.
func (recv *Queue) Init() {
	C.g_queue_init((*C.GQueue)(recv.native))

	return
}

// Regex is a wrapper around the C record GRegex.
type Regex struct {
	native *C.GRegex
}

func RegexNewFromC(u unsafe.Pointer) *Regex {
	c := (*C.GRegex)(u)
	if c == nil {
		return nil
	}

	g := &Regex{native: c}

	return g
}

func (recv *Regex) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RegexNew is a wrapper around the C function g_regex_new.
func RegexNew(pattern string, compileOptions RegexCompileFlags, matchOptions RegexMatchFlags) (*Regex, error) {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_compile_options := (C.GRegexCompileFlags)(compileOptions)

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var cThrowableError *C.GError

	retC := C.g_regex_new(c_pattern, c_compile_options, c_match_options, &cThrowableError)
	retGo := RegexNewFromC(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetCaptureCount is a wrapper around the C function g_regex_get_capture_count.
func (recv *Regex) GetCaptureCount() int32 {
	retC := C.g_regex_get_capture_count((*C.GRegex)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMaxBackref is a wrapper around the C function g_regex_get_max_backref.
func (recv *Regex) GetMaxBackref() int32 {
	retC := C.g_regex_get_max_backref((*C.GRegex)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPattern is a wrapper around the C function g_regex_get_pattern.
func (recv *Regex) GetPattern() string {
	retC := C.g_regex_get_pattern((*C.GRegex)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetStringNumber is a wrapper around the C function g_regex_get_string_number.
func (recv *Regex) GetStringNumber(name string) int32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_regex_get_string_number((*C.GRegex)(recv.native), c_name)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_regex_match : unsupported parameter match_info : record with indirection level of 2

// Unsupported : g_regex_match_all : unsupported parameter match_info : record with indirection level of 2

// Unsupported : g_regex_match_all_full : unsupported parameter string : no param type

// Unsupported : g_regex_match_full : unsupported parameter string : no param type

// Ref is a wrapper around the C function g_regex_ref.
func (recv *Regex) Ref() *Regex {
	retC := C.g_regex_ref((*C.GRegex)(recv.native))
	retGo := RegexNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_regex_replace : unsupported parameter string : no param type

// Unsupported : g_regex_replace_eval : unsupported parameter string : no param type

// Unsupported : g_regex_replace_literal : unsupported parameter string : no param type

// Unsupported : g_regex_split : no return type

// Unsupported : g_regex_split_full : unsupported parameter string : no param type

// Unref is a wrapper around the C function g_regex_unref.
func (recv *Regex) Unref() {
	C.g_regex_unref((*C.GRegex)(recv.native))

	return
}

// Append is a wrapper around the C function g_sequence_append.
func (recv *Sequence) Append(data uintptr) *SequenceIter {
	c_data := (C.gpointer)(data)

	retC := C.g_sequence_append((*C.GSequence)(recv.native), c_data)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_sequence_foreach : unsupported parameter func : no type generator for Func, GFunc

// Free is a wrapper around the C function g_sequence_free.
func (recv *Sequence) Free() {
	C.g_sequence_free((*C.GSequence)(recv.native))

	return
}

// GetBeginIter is a wrapper around the C function g_sequence_get_begin_iter.
func (recv *Sequence) GetBeginIter() *SequenceIter {
	retC := C.g_sequence_get_begin_iter((*C.GSequence)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetEndIter is a wrapper around the C function g_sequence_get_end_iter.
func (recv *Sequence) GetEndIter() *SequenceIter {
	retC := C.g_sequence_get_end_iter((*C.GSequence)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIterAtPos is a wrapper around the C function g_sequence_get_iter_at_pos.
func (recv *Sequence) GetIterAtPos(pos int32) *SequenceIter {
	c_pos := (C.gint)(pos)

	retC := C.g_sequence_get_iter_at_pos((*C.GSequence)(recv.native), c_pos)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLength is a wrapper around the C function g_sequence_get_length.
func (recv *Sequence) GetLength() int32 {
	retC := C.g_sequence_get_length((*C.GSequence)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_sequence_insert_sorted : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_insert_sorted_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Prepend is a wrapper around the C function g_sequence_prepend.
func (recv *Sequence) Prepend(data uintptr) *SequenceIter {
	c_data := (C.gpointer)(data)

	retC := C.g_sequence_prepend((*C.GSequence)(recv.native), c_data)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_sequence_search : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_search_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_sequence_sort : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_sort_iter : unsupported parameter cmp_func : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Compare is a wrapper around the C function g_sequence_iter_compare.
func (recv *SequenceIter) Compare(b *SequenceIter) int32 {
	c_b := (*C.GSequenceIter)(b.ToC())

	retC := C.g_sequence_iter_compare((*C.GSequenceIter)(recv.native), c_b)
	retGo := (int32)(retC)

	return retGo
}

// GetPosition is a wrapper around the C function g_sequence_iter_get_position.
func (recv *SequenceIter) GetPosition() int32 {
	retC := C.g_sequence_iter_get_position((*C.GSequenceIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSequence is a wrapper around the C function g_sequence_iter_get_sequence.
func (recv *SequenceIter) GetSequence() *Sequence {
	retC := C.g_sequence_iter_get_sequence((*C.GSequenceIter)(recv.native))
	retGo := SequenceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsBegin is a wrapper around the C function g_sequence_iter_is_begin.
func (recv *SequenceIter) IsBegin() bool {
	retC := C.g_sequence_iter_is_begin((*C.GSequenceIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsEnd is a wrapper around the C function g_sequence_iter_is_end.
func (recv *SequenceIter) IsEnd() bool {
	retC := C.g_sequence_iter_is_end((*C.GSequenceIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Move is a wrapper around the C function g_sequence_iter_move.
func (recv *SequenceIter) Move(delta int32) *SequenceIter {
	c_delta := (C.gint)(delta)

	retC := C.g_sequence_iter_move((*C.GSequenceIter)(recv.native), c_delta)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Next is a wrapper around the C function g_sequence_iter_next.
func (recv *SequenceIter) Next() *SequenceIter {
	retC := C.g_sequence_iter_next((*C.GSequenceIter)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Prev is a wrapper around the C function g_sequence_iter_prev.
func (recv *SequenceIter) Prev() *SequenceIter {
	retC := C.g_sequence_iter_prev((*C.GSequenceIter)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_string_append_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Overwrite is a wrapper around the C function g_string_overwrite.
func (recv *String) Overwrite(pos uint64, val string) *String {
	c_pos := (C.gsize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_overwrite((*C.GString)(recv.native), c_pos, c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// OverwriteLen is a wrapper around the C function g_string_overwrite_len.
func (recv *String) OverwriteLen(pos uint64, val string, len int64) *String {
	c_pos := (C.gsize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_overwrite_len((*C.GString)(recv.native), c_pos, c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_string_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Clear is a wrapper around the C function g_string_chunk_clear.
func (recv *StringChunk) Clear() {
	C.g_string_chunk_clear((*C.GStringChunk)(recv.native))

	return
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType
