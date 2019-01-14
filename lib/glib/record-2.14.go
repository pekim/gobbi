// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// HashTableGetKeys is a wrapper around the C function g_hash_table_get_keys.
func HashTableGetKeys(hashTable *HashTable) *List {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	retC := C.g_hash_table_get_keys(c_hash_table)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HashTableGetValues is a wrapper around the C function g_hash_table_get_values.
func HashTableGetValues(hashTable *HashTable) *List {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	retC := C.g_hash_table_get_values(c_hash_table)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LoadFromDirs is a wrapper around the C function g_key_file_load_from_dirs.
func (recv *KeyFile) LoadFromDirs(file string, searchDirs []string, flags KeyFileFlags) (bool, string, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_search_dirs_array := make([]*C.gchar, len(searchDirs), len(searchDirs))
	c_search_dirs_arrayPtr := &c_search_dirs_array[0]
	c_search_dirs := (**C.gchar)(unsafe.Pointer(c_search_dirs_arrayPtr))

	var c_full_path *C.gchar

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_dirs((*C.GKeyFile)(recv.native), c_file, c_search_dirs, &c_full_path, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	fullPath := C.GoString(c_full_path)
	defer C.free(unsafe.Pointer(c_full_path))

	return retGo, fullPath, goError
}

// ExpandReferences is a wrapper around the C function g_match_info_expand_references.
func (recv *MatchInfo) ExpandReferences(stringToExpand string) (string, error) {
	c_string_to_expand := C.CString(stringToExpand)
	defer C.free(unsafe.Pointer(c_string_to_expand))

	var cThrowableError *C.GError

	retC := C.g_match_info_expand_references((*C.GMatchInfo)(recv.native), c_string_to_expand, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Fetch is a wrapper around the C function g_match_info_fetch.
func (recv *MatchInfo) Fetch(matchNum int32) string {
	c_match_num := (C.gint)(matchNum)

	retC := C.g_match_info_fetch((*C.GMatchInfo)(recv.native), c_match_num)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FetchAll is a wrapper around the C function g_match_info_fetch_all.
func (recv *MatchInfo) FetchAll() []string {
	retC := C.g_match_info_fetch_all((*C.GMatchInfo)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// g_once_init_enter : unsupported parameter location : no type generator for gpointer (void*) for param location
// g_once_init_leave : unsupported parameter location : no type generator for gpointer (void*) for param location
// GetHelp is a wrapper around the C function g_option_context_get_help.
func (recv *OptionContext) GetHelp(mainHelp bool, group *OptionGroup) string {
	c_main_help :=
		boolToGboolean(mainHelp)

	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	retC := C.g_option_context_get_help((*C.GOptionContext)(recv.native), c_main_help, c_group)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

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

// Equals compares this Regex with another Regex, and returns true if they represent the same GObject.
func (recv *Regex) Equals(other *Regex) bool {
	return other.ToC() == recv.ToC()
}

// RegexNew is a wrapper around the C function g_regex_new.
func RegexNew(pattern string, compileOptions RegexCompileFlags, matchOptions RegexMatchFlags) (*Regex, error) {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_compile_options := (C.GRegexCompileFlags)(compileOptions)

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var cThrowableError *C.GError

	retC := C.g_regex_new(c_pattern, c_compile_options, c_match_options, &cThrowableError)
	var retGo (*Regex)
	if retC == nil {
		retGo = nil
	} else {
		retGo = RegexNewFromC(unsafe.Pointer(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RegexCheckReplacement is a wrapper around the C function g_regex_check_replacement.
func RegexCheckReplacement(replacement string) (bool, bool, error) {
	c_replacement := C.CString(replacement)
	defer C.free(unsafe.Pointer(c_replacement))

	var c_has_references C.gboolean

	var cThrowableError *C.GError

	retC := C.g_regex_check_replacement(c_replacement, &c_has_references, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	hasReferences := c_has_references == C.TRUE

	return retGo, hasReferences, goError
}

// RegexEscapeString is a wrapper around the C function g_regex_escape_string.
func RegexEscapeString(string_ []string) string {
	c_string_array := make([]C.gchar, len(string_), len(string_))
	c_string_arrayPtr := &c_string_array[0]
	c_string := (*C.gchar)(unsafe.Pointer(c_string_arrayPtr))

	c_length := (C.gint)(len(string_))

	retC := C.g_regex_escape_string(c_string, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// RegexMatchSimple is a wrapper around the C function g_regex_match_simple.
func RegexMatchSimple(pattern string, string_ string, compileOptions RegexCompileFlags, matchOptions RegexMatchFlags) bool {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_compile_options := (C.GRegexCompileFlags)(compileOptions)

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	retC := C.g_regex_match_simple(c_pattern, c_string, c_compile_options, c_match_options)
	retGo := retC == C.TRUE

	return retGo
}

// RegexSplitSimple is a wrapper around the C function g_regex_split_simple.
func RegexSplitSimple(pattern string, string_ string, compileOptions RegexCompileFlags, matchOptions RegexMatchFlags) []string {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_compile_options := (C.GRegexCompileFlags)(compileOptions)

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	retC := C.g_regex_split_simple(c_pattern, c_string, c_compile_options, c_match_options)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
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

// Match is a wrapper around the C function g_regex_match.
func (recv *Regex) Match(string_ string, matchOptions RegexMatchFlags) (bool, *MatchInfo) {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var c_match_info *C.GMatchInfo

	retC := C.g_regex_match((*C.GRegex)(recv.native), c_string, c_match_options, &c_match_info)
	retGo := retC == C.TRUE

	matchInfo := MatchInfoNewFromC(unsafe.Pointer(c_match_info))

	return retGo, matchInfo
}

// MatchAll is a wrapper around the C function g_regex_match_all.
func (recv *Regex) MatchAll(string_ string, matchOptions RegexMatchFlags) (bool, *MatchInfo) {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var c_match_info *C.GMatchInfo

	retC := C.g_regex_match_all((*C.GRegex)(recv.native), c_string, c_match_options, &c_match_info)
	retGo := retC == C.TRUE

	matchInfo := MatchInfoNewFromC(unsafe.Pointer(c_match_info))

	return retGo, matchInfo
}

// MatchAllFull is a wrapper around the C function g_regex_match_all_full.
func (recv *Regex) MatchAllFull(string_ []string, startPosition int32, matchOptions RegexMatchFlags) (bool, *MatchInfo, error) {
	c_string_array := make([]C.gchar, len(string_), len(string_))
	c_string_arrayPtr := &c_string_array[0]
	c_string := (*C.gchar)(unsafe.Pointer(c_string_arrayPtr))

	c_string_len := (C.gssize)(len(string_))

	c_start_position := (C.gint)(startPosition)

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var c_match_info *C.GMatchInfo

	var cThrowableError *C.GError

	retC := C.g_regex_match_all_full((*C.GRegex)(recv.native), c_string, c_string_len, c_start_position, c_match_options, &c_match_info, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	matchInfo := MatchInfoNewFromC(unsafe.Pointer(c_match_info))

	return retGo, matchInfo, goError
}

// MatchFull is a wrapper around the C function g_regex_match_full.
func (recv *Regex) MatchFull(string_ []string, startPosition int32, matchOptions RegexMatchFlags) (bool, *MatchInfo, error) {
	c_string_array := make([]C.gchar, len(string_), len(string_))
	c_string_arrayPtr := &c_string_array[0]
	c_string := (*C.gchar)(unsafe.Pointer(c_string_arrayPtr))

	c_string_len := (C.gssize)(len(string_))

	c_start_position := (C.gint)(startPosition)

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var c_match_info *C.GMatchInfo

	var cThrowableError *C.GError

	retC := C.g_regex_match_full((*C.GRegex)(recv.native), c_string, c_string_len, c_start_position, c_match_options, &c_match_info, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	matchInfo := MatchInfoNewFromC(unsafe.Pointer(c_match_info))

	return retGo, matchInfo, goError
}

// Ref is a wrapper around the C function g_regex_ref.
func (recv *Regex) Ref() *Regex {
	retC := C.g_regex_ref((*C.GRegex)(recv.native))
	retGo := RegexNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Replace is a wrapper around the C function g_regex_replace.
func (recv *Regex) Replace(string_ []string, startPosition int32, replacement string, matchOptions RegexMatchFlags) (string, error) {
	c_string_array := make([]C.gchar, len(string_), len(string_))
	c_string_arrayPtr := &c_string_array[0]
	c_string := (*C.gchar)(unsafe.Pointer(c_string_arrayPtr))

	c_string_len := (C.gssize)(len(string_))

	c_start_position := (C.gint)(startPosition)

	c_replacement := C.CString(replacement)
	defer C.free(unsafe.Pointer(c_replacement))

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var cThrowableError *C.GError

	retC := C.g_regex_replace((*C.GRegex)(recv.native), c_string, c_string_len, c_start_position, c_replacement, c_match_options, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_regex_replace_eval : unsupported parameter eval : no type generator for RegexEvalCallback (GRegexEvalCallback) for param eval

// ReplaceLiteral is a wrapper around the C function g_regex_replace_literal.
func (recv *Regex) ReplaceLiteral(string_ []string, startPosition int32, replacement string, matchOptions RegexMatchFlags) (string, error) {
	c_string_array := make([]C.gchar, len(string_), len(string_))
	c_string_arrayPtr := &c_string_array[0]
	c_string := (*C.gchar)(unsafe.Pointer(c_string_arrayPtr))

	c_string_len := (C.gssize)(len(string_))

	c_start_position := (C.gint)(startPosition)

	c_replacement := C.CString(replacement)
	defer C.free(unsafe.Pointer(c_replacement))

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var cThrowableError *C.GError

	retC := C.g_regex_replace_literal((*C.GRegex)(recv.native), c_string, c_string_len, c_start_position, c_replacement, c_match_options, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Split is a wrapper around the C function g_regex_split.
func (recv *Regex) Split(string_ string, matchOptions RegexMatchFlags) []string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	retC := C.g_regex_split((*C.GRegex)(recv.native), c_string, c_match_options)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// SplitFull is a wrapper around the C function g_regex_split_full.
func (recv *Regex) SplitFull(string_ []string, startPosition int32, matchOptions RegexMatchFlags, maxTokens int32) ([]string, error) {
	c_string_array := make([]C.gchar, len(string_), len(string_))
	c_string_arrayPtr := &c_string_array[0]
	c_string := (*C.gchar)(unsafe.Pointer(c_string_arrayPtr))

	c_string_len := (C.gssize)(len(string_))

	c_start_position := (C.gint)(startPosition)

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	c_max_tokens := (C.gint)(maxTokens)

	var cThrowableError *C.GError

	retC := C.g_regex_split_full((*C.GRegex)(recv.native), c_string, c_string_len, c_start_position, c_match_options, c_max_tokens, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unref is a wrapper around the C function g_regex_unref.
func (recv *Regex) Unref() {
	C.g_regex_unref((*C.GRegex)(recv.native))

	return
}

// g_sequence_foreach_range : unsupported parameter func : no type generator for Func (GFunc) for param func
// g_sequence_get : no return generator
// g_sequence_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// SequenceMove is a wrapper around the C function g_sequence_move.
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

// SequenceMoveRange is a wrapper around the C function g_sequence_move_range.
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

// g_sequence_new : unsupported parameter data_destroy : no type generator for DestroyNotify (GDestroyNotify) for param data_destroy
// SequenceRangeGetMidpoint is a wrapper around the C function g_sequence_range_get_midpoint.
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

// SequenceRemove is a wrapper around the C function g_sequence_remove.
func SequenceRemove(iter *SequenceIter) {
	c_iter := (*C.GSequenceIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GSequenceIter)(iter.ToC())
	}

	C.g_sequence_remove(c_iter)

	return
}

// SequenceRemoveRange is a wrapper around the C function g_sequence_remove_range.
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

// g_sequence_set : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_sequence_sort_changed : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func
// g_sequence_sort_changed_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param iter_cmp
// SequenceSwap is a wrapper around the C function g_sequence_swap.
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

// Unsupported : g_sequence_append : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func

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

// Unsupported : g_sequence_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_insert_sorted_iter : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_prepend : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_search : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_search_iter : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_sort : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func

// Unsupported : g_sequence_sort_iter : unsupported parameter cmp_func : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param cmp_func

// Compare is a wrapper around the C function g_sequence_iter_compare.
func (recv *SequenceIter) Compare(b *SequenceIter) int32 {
	c_b := (*C.GSequenceIter)(C.NULL)
	if b != nil {
		c_b = (*C.GSequenceIter)(b.ToC())
	}

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

// Unsupported : g_string_append_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

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

// Unsupported : g_string_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Clear is a wrapper around the C function g_string_chunk_clear.
func (recv *StringChunk) Clear() {
	C.g_string_chunk_clear((*C.GStringChunk)(recv.native))

	return
}
