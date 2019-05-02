// Code generated - DO NOT EDIT.
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

type RegexCompileFlags C.GRegexCompileFlags

const (
	REGEX_CASELESS          RegexCompileFlags = 1
	REGEX_MULTILINE         RegexCompileFlags = 2
	REGEX_DOTALL            RegexCompileFlags = 4
	REGEX_EXTENDED          RegexCompileFlags = 8
	REGEX_ANCHORED          RegexCompileFlags = 16
	REGEX_DOLLAR_ENDONLY    RegexCompileFlags = 32
	REGEX_UNGREEDY          RegexCompileFlags = 512
	REGEX_RAW               RegexCompileFlags = 2048
	REGEX_NO_AUTO_CAPTURE   RegexCompileFlags = 4096
	REGEX_OPTIMIZE          RegexCompileFlags = 8192
	REGEX_FIRSTLINE         RegexCompileFlags = 262144
	REGEX_DUPNAMES          RegexCompileFlags = 524288
	REGEX_NEWLINE_CR        RegexCompileFlags = 1048576
	REGEX_NEWLINE_LF        RegexCompileFlags = 2097152
	REGEX_NEWLINE_CRLF      RegexCompileFlags = 3145728
	REGEX_NEWLINE_ANYCRLF   RegexCompileFlags = 5242880
	REGEX_BSR_ANYCRLF       RegexCompileFlags = 8388608
	REGEX_JAVASCRIPT_COMPAT RegexCompileFlags = 33554432
)

type RegexMatchFlags C.GRegexMatchFlags

const (
	REGEX_MATCH_ANCHORED         RegexMatchFlags = 16
	REGEX_MATCH_NOTBOL           RegexMatchFlags = 128
	REGEX_MATCH_NOTEOL           RegexMatchFlags = 256
	REGEX_MATCH_NOTEMPTY         RegexMatchFlags = 1024
	REGEX_MATCH_PARTIAL          RegexMatchFlags = 32768
	REGEX_MATCH_NEWLINE_CR       RegexMatchFlags = 1048576
	REGEX_MATCH_NEWLINE_LF       RegexMatchFlags = 2097152
	REGEX_MATCH_NEWLINE_CRLF     RegexMatchFlags = 3145728
	REGEX_MATCH_NEWLINE_ANY      RegexMatchFlags = 4194304
	REGEX_MATCH_NEWLINE_ANYCRLF  RegexMatchFlags = 5242880
	REGEX_MATCH_BSR_ANYCRLF      RegexMatchFlags = 8388608
	REGEX_MATCH_BSR_ANY          RegexMatchFlags = 16777216
	REGEX_MATCH_PARTIAL_SOFT     RegexMatchFlags = 32768
	REGEX_MATCH_PARTIAL_HARD     RegexMatchFlags = 134217728
	REGEX_MATCH_NOTEMPTY_ATSTART RegexMatchFlags = 268435456
)
const KEY_FILE_DESKTOP_GROUP string = C.G_KEY_FILE_DESKTOP_GROUP
const KEY_FILE_DESKTOP_KEY_CATEGORIES string = C.G_KEY_FILE_DESKTOP_KEY_CATEGORIES
const KEY_FILE_DESKTOP_KEY_COMMENT string = C.G_KEY_FILE_DESKTOP_KEY_COMMENT
const KEY_FILE_DESKTOP_KEY_EXEC string = C.G_KEY_FILE_DESKTOP_KEY_EXEC
const KEY_FILE_DESKTOP_KEY_GENERIC_NAME string = C.G_KEY_FILE_DESKTOP_KEY_GENERIC_NAME
const KEY_FILE_DESKTOP_KEY_HIDDEN string = C.G_KEY_FILE_DESKTOP_KEY_HIDDEN
const KEY_FILE_DESKTOP_KEY_ICON string = C.G_KEY_FILE_DESKTOP_KEY_ICON
const KEY_FILE_DESKTOP_KEY_MIME_TYPE string = C.G_KEY_FILE_DESKTOP_KEY_MIME_TYPE
const KEY_FILE_DESKTOP_KEY_NAME string = C.G_KEY_FILE_DESKTOP_KEY_NAME
const KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN string = C.G_KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN
const KEY_FILE_DESKTOP_KEY_NO_DISPLAY string = C.G_KEY_FILE_DESKTOP_KEY_NO_DISPLAY
const KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN string = C.G_KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN
const KEY_FILE_DESKTOP_KEY_PATH string = C.G_KEY_FILE_DESKTOP_KEY_PATH
const KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY string = C.G_KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY
const KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS string = C.G_KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS
const KEY_FILE_DESKTOP_KEY_TERMINAL string = C.G_KEY_FILE_DESKTOP_KEY_TERMINAL
const KEY_FILE_DESKTOP_KEY_TRY_EXEC string = C.G_KEY_FILE_DESKTOP_KEY_TRY_EXEC
const KEY_FILE_DESKTOP_KEY_TYPE string = C.G_KEY_FILE_DESKTOP_KEY_TYPE
const KEY_FILE_DESKTOP_KEY_URL string = C.G_KEY_FILE_DESKTOP_KEY_URL
const KEY_FILE_DESKTOP_KEY_VERSION string = C.G_KEY_FILE_DESKTOP_KEY_VERSION
const KEY_FILE_DESKTOP_TYPE_APPLICATION string = C.G_KEY_FILE_DESKTOP_TYPE_APPLICATION
const KEY_FILE_DESKTOP_TYPE_DIRECTORY string = C.G_KEY_FILE_DESKTOP_TYPE_DIRECTORY
const KEY_FILE_DESKTOP_TYPE_LINK string = C.G_KEY_FILE_DESKTOP_TYPE_LINK

type RegexError C.GRegexError

const (
	REGEX_ERROR_COMPILE                                      RegexError = 0
	REGEX_ERROR_OPTIMIZE                                     RegexError = 1
	REGEX_ERROR_REPLACE                                      RegexError = 2
	REGEX_ERROR_MATCH                                        RegexError = 3
	REGEX_ERROR_INTERNAL                                     RegexError = 4
	REGEX_ERROR_STRAY_BACKSLASH                              RegexError = 101
	REGEX_ERROR_MISSING_CONTROL_CHAR                         RegexError = 102
	REGEX_ERROR_UNRECOGNIZED_ESCAPE                          RegexError = 103
	REGEX_ERROR_QUANTIFIERS_OUT_OF_ORDER                     RegexError = 104
	REGEX_ERROR_QUANTIFIER_TOO_BIG                           RegexError = 105
	REGEX_ERROR_UNTERMINATED_CHARACTER_CLASS                 RegexError = 106
	REGEX_ERROR_INVALID_ESCAPE_IN_CHARACTER_CLASS            RegexError = 107
	REGEX_ERROR_RANGE_OUT_OF_ORDER                           RegexError = 108
	REGEX_ERROR_NOTHING_TO_REPEAT                            RegexError = 109
	REGEX_ERROR_UNRECOGNIZED_CHARACTER                       RegexError = 112
	REGEX_ERROR_POSIX_NAMED_CLASS_OUTSIDE_CLASS              RegexError = 113
	REGEX_ERROR_UNMATCHED_PARENTHESIS                        RegexError = 114
	REGEX_ERROR_INEXISTENT_SUBPATTERN_REFERENCE              RegexError = 115
	REGEX_ERROR_UNTERMINATED_COMMENT                         RegexError = 118
	REGEX_ERROR_EXPRESSION_TOO_LARGE                         RegexError = 120
	REGEX_ERROR_MEMORY_ERROR                                 RegexError = 121
	REGEX_ERROR_VARIABLE_LENGTH_LOOKBEHIND                   RegexError = 125
	REGEX_ERROR_MALFORMED_CONDITION                          RegexError = 126
	REGEX_ERROR_TOO_MANY_CONDITIONAL_BRANCHES                RegexError = 127
	REGEX_ERROR_ASSERTION_EXPECTED                           RegexError = 128
	REGEX_ERROR_UNKNOWN_POSIX_CLASS_NAME                     RegexError = 130
	REGEX_ERROR_POSIX_COLLATING_ELEMENTS_NOT_SUPPORTED       RegexError = 131
	REGEX_ERROR_HEX_CODE_TOO_LARGE                           RegexError = 134
	REGEX_ERROR_INVALID_CONDITION                            RegexError = 135
	REGEX_ERROR_SINGLE_BYTE_MATCH_IN_LOOKBEHIND              RegexError = 136
	REGEX_ERROR_INFINITE_LOOP                                RegexError = 140
	REGEX_ERROR_MISSING_SUBPATTERN_NAME_TERMINATOR           RegexError = 142
	REGEX_ERROR_DUPLICATE_SUBPATTERN_NAME                    RegexError = 143
	REGEX_ERROR_MALFORMED_PROPERTY                           RegexError = 146
	REGEX_ERROR_UNKNOWN_PROPERTY                             RegexError = 147
	REGEX_ERROR_SUBPATTERN_NAME_TOO_LONG                     RegexError = 148
	REGEX_ERROR_TOO_MANY_SUBPATTERNS                         RegexError = 149
	REGEX_ERROR_INVALID_OCTAL_VALUE                          RegexError = 151
	REGEX_ERROR_TOO_MANY_BRANCHES_IN_DEFINE                  RegexError = 154
	REGEX_ERROR_DEFINE_REPETION                              RegexError = 155
	REGEX_ERROR_INCONSISTENT_NEWLINE_OPTIONS                 RegexError = 156
	REGEX_ERROR_MISSING_BACK_REFERENCE                       RegexError = 157
	REGEX_ERROR_INVALID_RELATIVE_REFERENCE                   RegexError = 158
	REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_FORBIDDEN RegexError = 159
	REGEX_ERROR_UNKNOWN_BACKTRACKING_CONTROL_VERB            RegexError = 160
	REGEX_ERROR_NUMBER_TOO_BIG                               RegexError = 161
	REGEX_ERROR_MISSING_SUBPATTERN_NAME                      RegexError = 162
	REGEX_ERROR_MISSING_DIGIT                                RegexError = 163
	REGEX_ERROR_INVALID_DATA_CHARACTER                       RegexError = 164
	REGEX_ERROR_EXTRA_SUBPATTERN_NAME                        RegexError = 165
	REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_REQUIRED  RegexError = 166
	REGEX_ERROR_INVALID_CONTROL_CHAR                         RegexError = 168
	REGEX_ERROR_MISSING_NAME                                 RegexError = 169
	REGEX_ERROR_NOT_SUPPORTED_IN_CLASS                       RegexError = 171
	REGEX_ERROR_TOO_MANY_FORWARD_REFERENCES                  RegexError = 172
	REGEX_ERROR_NAME_TOO_LONG                                RegexError = 175
	REGEX_ERROR_CHARACTER_VALUE_TOO_LARGE                    RegexError = 176
)

type UserDirectory C.GUserDirectory

const (
	USER_DIRECTORY_DESKTOP      UserDirectory = 0
	USER_DIRECTORY_DOCUMENTS    UserDirectory = 1
	USER_DIRECTORY_DOWNLOAD     UserDirectory = 2
	USER_DIRECTORY_MUSIC        UserDirectory = 3
	USER_DIRECTORY_PICTURES     UserDirectory = 4
	USER_DIRECTORY_PUBLIC_SHARE UserDirectory = 5
	USER_DIRECTORY_TEMPLATES    UserDirectory = 6
	USER_DIRECTORY_VIDEOS       UserDirectory = 7
	USER_N_DIRECTORIES          UserDirectory = 8
)

// GetUserSpecialDir is a wrapper around the C function g_get_user_special_dir.
func GetUserSpecialDir(directory UserDirectory) string {
	c_directory := (C.GUserDirectory)(directory)

	retC := C.g_get_user_special_dir(c_directory)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_once_init_enter : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_once_init_leave : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_sequence_get : no return generator

// Unsupported : g_sequence_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_set : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_slice_copy : unsupported parameter mem_block : no type generator for gpointer (gconstpointer) for param mem_block

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// TimeoutSourceNewSeconds is a wrapper around the C function g_timeout_source_new_seconds.
func TimeoutSourceNewSeconds(interval uint32) *Source {
	c_interval := (C.guint)(interval)

	retC := C.g_timeout_source_new_seconds(c_interval)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnicharCombiningClass is a wrapper around the C function g_unichar_combining_class.
func UnicharCombiningClass(uc rune) int32 {
	c_uc := (C.gunichar)(uc)

	retC := C.g_unichar_combining_class(c_uc)
	retGo := (int32)(retC)

	return retGo
}

// UnicharGetScript is a wrapper around the C function g_unichar_get_script.
func UnicharGetScript(ch rune) UnicodeScript {
	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_get_script(c_ch)
	retGo := (UnicodeScript)(retC)

	return retGo
}

// UnicharIsmark is a wrapper around the C function g_unichar_ismark.
func UnicharIsmark(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_ismark(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIszerowidth is a wrapper around the C function g_unichar_iszerowidth.
func UnicharIszerowidth(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iszerowidth(c_c)
	retGo := retC == C.TRUE

	return retGo
}

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

	c_search_dirs_array := make([]*C.gchar, len(searchDirs)+1, len(searchDirs)+1)
	for i, item := range searchDirs {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_search_dirs_array[i] = c
	}
	c_search_dirs_array[len(searchDirs)] = nil
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

// Unsupported : g_match_info_get_regex : return type : Blacklisted record : GRegex

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

// Blacklisted : GRegex

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
