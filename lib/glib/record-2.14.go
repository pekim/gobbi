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

// Unsupported : g_key_file_load_from_dirs : unsupported parameter search_dirs :

// Returns a new string containing the text in @string_to_expand with
// references and escape sequences expanded. References refer to the last
// match done with @string against @regex and have the same syntax used by
// g_regex_replace().
//
// The @string_to_expand must be UTF-8 encoded even if #G_REGEX_RAW was
// passed to g_regex_new().
//
// The backreferences are extracted from the string passed to the match
// function, so you cannot call this function after freeing the string.
//
// @match_info may be %NULL in which case @string_to_expand must not
// contain references. For instance "foo\n" does not refer to an actual
// pattern and '\n' merely will be replaced with \n character,
// while to expand "\0" (whole match) one needs the result of a match.
// Use g_regex_check_replacement() to find out whether @string_to_expand
// contains references.
/*

C function : g_match_info_expand_references
*/
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

// Retrieves the text matching the @match_num'th capturing
// parentheses. 0 is the full text of the match, 1 is the first paren
// set, 2 the second, and so on.
//
// If @match_num is a valid sub pattern but it didn't match anything
// (e.g. sub pattern 1, matching "b" against "(a)?b") then an empty
// string is returned.
//
// If the match was obtained using the DFA algorithm, that is using
// g_regex_match_all() or g_regex_match_all_full(), the retrieved
// string is not that of a set of parentheses but that of a matched
// substring. Substrings are matched in reverse order of length, so
// 0 is the longest match.
//
// The string is fetched from the string passed to the match function,
// so you cannot call this function after freeing the string.
/*

C function : g_match_info_fetch
*/
func (recv *MatchInfo) Fetch(matchNum int32) string {
	c_match_num := (C.gint)(matchNum)

	retC := C.g_match_info_fetch((*C.GMatchInfo)(recv.native), c_match_num)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_match_info_fetch_all : no return type

// Retrieves the text matching the capturing parentheses named @name.
//
// If @name is a valid sub pattern name but it didn't match anything
// (e.g. sub pattern "X", matching "b" against "(?P<X>a)?b")
// then an empty string is returned.
//
// The string is fetched from the string passed to the match function,
// so you cannot call this function after freeing the string.
/*

C function : g_match_info_fetch_named
*/
func (recv *MatchInfo) FetchNamed(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_match_info_fetch_named((*C.GMatchInfo)(recv.native), c_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the position in bytes of the capturing parentheses named @name.
//
// If @name is a valid sub pattern name but it didn't match anything
// (e.g. sub pattern "X", matching "b" against "(?P<X>a)?b")
// then @start_pos and @end_pos are set to -1 and %TRUE is returned.
/*

C function : g_match_info_fetch_named_pos
*/
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

// Retrieves the position in bytes of the @match_num'th capturing
// parentheses. 0 is the full text of the match, 1 is the first
// paren set, 2 the second, and so on.
//
// If @match_num is a valid sub pattern but it didn't match anything
// (e.g. sub pattern 1, matching "b" against "(a)?b") then @start_pos
// and @end_pos are set to -1 and %TRUE is returned.
//
// If the match was obtained using the DFA algorithm, that is using
// g_regex_match_all() or g_regex_match_all_full(), the retrieved
// position is not that of a set of parentheses but that of a matched
// substring. Substrings are matched in reverse order of length, so
// 0 is the longest match.
/*

C function : g_match_info_fetch_pos
*/
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

// If @match_info is not %NULL, calls g_match_info_unref(); otherwise does
// nothing.
/*

C function : g_match_info_free
*/
func (recv *MatchInfo) Free() {
	C.g_match_info_free((*C.GMatchInfo)(recv.native))

	return
}

// Retrieves the number of matched substrings (including substring 0,
// that is the whole matched text), so 1 is returned if the pattern
// has no substrings in it and 0 is returned if the match failed.
//
// If the last match was obtained using the DFA algorithm, that is
// using g_regex_match_all() or g_regex_match_all_full(), the retrieved
// count is not that of the number of capturing parentheses but that of
// the number of matched substrings.
/*

C function : g_match_info_get_match_count
*/
func (recv *MatchInfo) GetMatchCount() int32 {
	retC := C.g_match_info_get_match_count((*C.GMatchInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns #GRegex object used in @match_info. It belongs to Glib
// and must not be freed. Use g_regex_ref() if you need to keep it
// after you free @match_info object.
/*

C function : g_match_info_get_regex
*/
func (recv *MatchInfo) GetRegex() *Regex {
	retC := C.g_match_info_get_regex((*C.GMatchInfo)(recv.native))
	retGo := RegexNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the string searched with @match_info. This is the
// string passed to g_regex_match() or g_regex_replace() so
// you may not free it before calling this function.
/*

C function : g_match_info_get_string
*/
func (recv *MatchInfo) GetString() string {
	retC := C.g_match_info_get_string((*C.GMatchInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Usually if the string passed to g_regex_match*() matches as far as
// it goes, but is too short to match the entire pattern, %FALSE is
// returned. There are circumstances where it might be helpful to
// distinguish this case from other cases in which there is no match.
//
// Consider, for example, an application where a human is required to
// type in data for a field with specific formatting requirements. An
// example might be a date in the form ddmmmyy, defined by the pattern
// "^\d?\d(jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)\d\d$".
// If the application sees the user’s keystrokes one by one, and can
// check that what has been typed so far is potentially valid, it is
// able to raise an error as soon as a mistake is made.
//
// GRegex supports the concept of partial matching by means of the
// #G_REGEX_MATCH_PARTIAL_SOFT and #G_REGEX_MATCH_PARTIAL_HARD flags.
// When they are used, the return code for
// g_regex_match() or g_regex_match_full() is, as usual, %TRUE
// for a complete match, %FALSE otherwise. But, when these functions
// return %FALSE, you can check if the match was partial calling
// g_match_info_is_partial_match().
//
// The difference between #G_REGEX_MATCH_PARTIAL_SOFT and
// #G_REGEX_MATCH_PARTIAL_HARD is that when a partial match is encountered
// with #G_REGEX_MATCH_PARTIAL_SOFT, matching continues to search for a
// possible complete match, while with #G_REGEX_MATCH_PARTIAL_HARD matching
// stops at the partial match.
// When both #G_REGEX_MATCH_PARTIAL_SOFT and #G_REGEX_MATCH_PARTIAL_HARD
// are set, the latter takes precedence.
//
// There were formerly some restrictions on the pattern for partial matching.
// The restrictions no longer apply.
//
// See pcrepartial(3) for more information on partial matching.
/*

C function : g_match_info_is_partial_match
*/
func (recv *MatchInfo) IsPartialMatch() bool {
	retC := C.g_match_info_is_partial_match((*C.GMatchInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the previous match operation succeeded.
/*

C function : g_match_info_matches
*/
func (recv *MatchInfo) Matches() bool {
	retC := C.g_match_info_matches((*C.GMatchInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Scans for the next match using the same parameters of the previous
// call to g_regex_match_full() or g_regex_match() that returned
// @match_info.
//
// The match is done on the string passed to the match function, so you
// cannot free it before calling this function.
/*

C function : g_match_info_next
*/
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

// Returns a formatted, translated help text for the given context.
// To obtain the text produced by `--help`, call
// `g_option_context_get_help (context, TRUE, NULL)`.
// To obtain the text produced by `--help-all`, call
// `g_option_context_get_help (context, FALSE, NULL)`.
// To obtain the help text for an option group, call
// `g_option_context_get_help (context, FALSE, group)`.
/*

C function : g_option_context_get_help
*/
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

// Removes all the elements in @queue. If queue elements contain
// dynamically-allocated memory, they should be freed first.
/*

C function : g_queue_clear
*/
func (recv *Queue) Clear() {
	C.g_queue_clear((*C.GQueue)(recv.native))

	return
}

// A statically-allocated #GQueue must be initialized with this function
// before it can be used. Alternatively you can initialize it with
// #G_QUEUE_INIT. It is not necessary to initialize queues created with
// g_queue_new().
/*

C function : g_queue_init
*/
func (recv *Queue) Init() {
	C.g_queue_init((*C.GQueue)(recv.native))

	return
}

// The g_regex_*() functions implement regular
// expression pattern matching using syntax and semantics similar to
// Perl regular expression.
//
// Some functions accept a @start_position argument, setting it differs
// from just passing over a shortened string and setting #G_REGEX_MATCH_NOTBOL
// in the case of a pattern that begins with any kind of lookbehind assertion.
// For example, consider the pattern "\Biss\B" which finds occurrences of "iss"
// in the middle of words. ("\B" matches only if the current position in the
// subject is not a word boundary.) When applied to the string "Mississipi"
// from the fourth byte, namely "issipi", it does not match, because "\B" is
// always false at the start of the subject, which is deemed to be a word
// boundary. However, if the entire string is passed , but with
// @start_position set to 4, it finds the second occurrence of "iss" because
// it is able to look behind the starting point to discover that it is
// preceded by a letter.
//
// Note that, unless you set the #G_REGEX_RAW flag, all the strings passed
// to these functions must be encoded in UTF-8. The lengths and the positions
// inside the strings are in bytes and not in characters, so, for instance,
// "\xc3\xa0" (i.e. "à") is two bytes long but it is treated as a
// single character. If you set #G_REGEX_RAW the strings can be non-valid
// UTF-8 strings and a byte is treated as a character, so "\xc3\xa0" is two
// bytes and two characters long.
//
// When matching a pattern, "\n" matches only against a "\n" character in
// the string, and "\r" matches only a "\r" character. To match any newline
// sequence use "\R". This particular group matches either the two-character
// sequence CR + LF ("\r\n"), or one of the single characters LF (linefeed,
// U+000A, "\n"), VT vertical tab, U+000B, "\v"), FF (formfeed, U+000C, "\f"),
// CR (carriage return, U+000D, "\r"), NEL (next line, U+0085), LS (line
// separator, U+2028), or PS (paragraph separator, U+2029).
//
// The behaviour of the dot, circumflex, and dollar metacharacters are
// affected by newline characters, the default is to recognize any newline
// character (the same characters recognized by "\R"). This can be changed
// with #G_REGEX_NEWLINE_CR, #G_REGEX_NEWLINE_LF and #G_REGEX_NEWLINE_CRLF
// compile options, and with #G_REGEX_MATCH_NEWLINE_ANY,
// #G_REGEX_MATCH_NEWLINE_CR, #G_REGEX_MATCH_NEWLINE_LF and
// #G_REGEX_MATCH_NEWLINE_CRLF match options. These settings are also
// relevant when compiling a pattern if #G_REGEX_EXTENDED is set, and an
// unescaped "#" outside a character class is encountered. This indicates
// a comment that lasts until after the next newline.
//
// When setting the %G_REGEX_JAVASCRIPT_COMPAT flag, pattern syntax and pattern
// matching is changed to be compatible with the way that regular expressions
// work in JavaScript. More precisely, a lonely ']' character in the pattern
// is a syntax error; the '\x' escape only allows 0 to 2 hexadecimal digits, and
// you must use the '\u' escape sequence with 4 hex digits to specify a unicode
// codepoint instead of '\x' or 'x{....}'. If '\x' or '\u' are not followed by
// the specified number of hex digits, they match 'x' and 'u' literally; also
// '\U' always matches 'U' instead of being an error in the pattern. Finally,
// pattern matching is modified so that back references to an unset subpattern
// group produces a match with the empty string instead of an error. See
// pcreapi(3) for more information.
//
// Creating and manipulating the same #GRegex structure from different
// threads is not a problem as #GRegex does not modify its internal
// state between creation and destruction, on the other hand #GMatchInfo
// is not threadsafe.
//
// The regular expressions low-level functionalities are obtained through
// the excellent
// [PCRE](http://www.pcre.org/)
// library written by Philip Hazel.
/*

C record/class : GRegex
*/
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

// Compiles the regular expression to an internal form, and does
// the initial setup of the #GRegex structure.
/*

C function : g_regex_new
*/
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

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Returns the number of capturing subpatterns in the pattern.
/*

C function : g_regex_get_capture_count
*/
func (recv *Regex) GetCaptureCount() int32 {
	retC := C.g_regex_get_capture_count((*C.GRegex)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the number of the highest back reference
// in the pattern, or 0 if the pattern does not contain
// back references.
/*

C function : g_regex_get_max_backref
*/
func (recv *Regex) GetMaxBackref() int32 {
	retC := C.g_regex_get_max_backref((*C.GRegex)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the pattern string associated with @regex, i.e. a copy of
// the string passed to g_regex_new().
/*

C function : g_regex_get_pattern
*/
func (recv *Regex) GetPattern() string {
	retC := C.g_regex_get_pattern((*C.GRegex)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the number of the subexpression named @name.
/*

C function : g_regex_get_string_number
*/
func (recv *Regex) GetStringNumber(name string) int32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_regex_get_string_number((*C.GRegex)(recv.native), c_name)
	retGo := (int32)(retC)

	return retGo
}

// Scans for a match in string for the pattern in @regex.
// The @match_options are combined with the match options specified
// when the @regex structure was created, letting you have more
// flexibility in reusing #GRegex structures.
//
// A #GMatchInfo structure, used to get information on the match,
// is stored in @match_info if not %NULL. Note that if @match_info
// is not %NULL then it is created even if the function returns %FALSE,
// i.e. you must free it regardless if regular expression actually matched.
//
// To retrieve all the non-overlapping matches of the pattern in
// string you can use g_match_info_next().
//
// |[<!-- language="C" -->
// static void
// print_uppercase_words (const gchar *string)
// {
// Print all uppercase-only words.
// GRegex *regex;
// GMatchInfo *match_info;
//
// regex = g_regex_new ("[A-Z]+", 0, 0, NULL);
// g_regex_match (regex, string, 0, &match_info);
// while (g_match_info_matches (match_info))
// {
// gchar *word = g_match_info_fetch (match_info, 0);
// g_print ("Found: %s\n", word);
// g_free (word);
// g_match_info_next (match_info, NULL);
// }
// g_match_info_free (match_info);
// g_regex_unref (regex);
// }
// ]|
//
// @string is not copied and is used in #GMatchInfo internally. If
// you use any #GMatchInfo method (except g_match_info_free()) after
// freeing or modifying @string then the behaviour is undefined.
/*

C function : g_regex_match
*/
func (recv *Regex) Match(string string, matchOptions RegexMatchFlags) (bool, *MatchInfo) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var c_match_info *C.GMatchInfo

	retC := C.g_regex_match((*C.GRegex)(recv.native), c_string, c_match_options, &c_match_info)
	retGo := retC == C.TRUE

	matchInfo := MatchInfoNewFromC(unsafe.Pointer(c_match_info))

	return retGo, matchInfo
}

// Using the standard algorithm for regular expression matching only
// the longest match in the string is retrieved. This function uses
// a different algorithm so it can retrieve all the possible matches.
// For more documentation see g_regex_match_all_full().
//
// A #GMatchInfo structure, used to get information on the match, is
// stored in @match_info if not %NULL. Note that if @match_info is
// not %NULL then it is created even if the function returns %FALSE,
// i.e. you must free it regardless if regular expression actually
// matched.
//
// @string is not copied and is used in #GMatchInfo internally. If
// you use any #GMatchInfo method (except g_match_info_free()) after
// freeing or modifying @string then the behaviour is undefined.
/*

C function : g_regex_match_all
*/
func (recv *Regex) MatchAll(string string, matchOptions RegexMatchFlags) (bool, *MatchInfo) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	var c_match_info *C.GMatchInfo

	retC := C.g_regex_match_all((*C.GRegex)(recv.native), c_string, c_match_options, &c_match_info)
	retGo := retC == C.TRUE

	matchInfo := MatchInfoNewFromC(unsafe.Pointer(c_match_info))

	return retGo, matchInfo
}

// Unsupported : g_regex_match_all_full : unsupported parameter string :

// Unsupported : g_regex_match_full : unsupported parameter string :

// Increases reference count of @regex by 1.
/*

C function : g_regex_ref
*/
func (recv *Regex) Ref() *Regex {
	retC := C.g_regex_ref((*C.GRegex)(recv.native))
	retGo := RegexNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_regex_replace : unsupported parameter string :

// Unsupported : g_regex_replace_eval : unsupported parameter string :

// Unsupported : g_regex_replace_literal : unsupported parameter string :

// Unsupported : g_regex_split : no return type

// Unsupported : g_regex_split_full : unsupported parameter string :

// Decreases reference count of @regex by 1. When reference count drops
// to zero, it frees all the memory associated with the regex structure.
/*

C function : g_regex_unref
*/
func (recv *Regex) Unref() {
	C.g_regex_unref((*C.GRegex)(recv.native))

	return
}

// Adds a new item to the end of @seq.
/*

C function : g_sequence_append
*/
func (recv *Sequence) Append(data uintptr) *SequenceIter {
	c_data := (C.gpointer)(data)

	retC := C.g_sequence_append((*C.GSequence)(recv.native), c_data)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_sequence_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func

// Frees the memory allocated for @seq. If @seq has a data destroy
// function associated with it, that function is called on all items
// in @seq.
/*

C function : g_sequence_free
*/
func (recv *Sequence) Free() {
	C.g_sequence_free((*C.GSequence)(recv.native))

	return
}

// Returns the begin iterator for @seq.
/*

C function : g_sequence_get_begin_iter
*/
func (recv *Sequence) GetBeginIter() *SequenceIter {
	retC := C.g_sequence_get_begin_iter((*C.GSequence)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the end iterator for @seg
/*

C function : g_sequence_get_end_iter
*/
func (recv *Sequence) GetEndIter() *SequenceIter {
	retC := C.g_sequence_get_end_iter((*C.GSequence)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the iterator at position @pos. If @pos is negative or larger
// than the number of items in @seq, the end iterator is returned.
/*

C function : g_sequence_get_iter_at_pos
*/
func (recv *Sequence) GetIterAtPos(pos int32) *SequenceIter {
	c_pos := (C.gint)(pos)

	retC := C.g_sequence_get_iter_at_pos((*C.GSequence)(recv.native), c_pos)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the length of @seq. Note that this method is O(h) where `h' is the
// height of the tree. It is thus more efficient to use g_sequence_is_empty()
// when comparing the length to zero.
/*

C function : g_sequence_get_length
*/
func (recv *Sequence) GetLength() int32 {
	retC := C.g_sequence_get_length((*C.GSequence)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_sequence_insert_sorted : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func

// Unsupported : g_sequence_insert_sorted_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param iter_cmp

// Adds a new item to the front of @seq
/*

C function : g_sequence_prepend
*/
func (recv *Sequence) Prepend(data uintptr) *SequenceIter {
	c_data := (C.gpointer)(data)

	retC := C.g_sequence_prepend((*C.GSequence)(recv.native), c_data)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_sequence_search : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func

// Unsupported : g_sequence_search_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param iter_cmp

// Unsupported : g_sequence_sort : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func

// Unsupported : g_sequence_sort_iter : unsupported parameter cmp_func : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param cmp_func

// Returns a negative number if @a comes before @b, 0 if they are equal,
// and a positive number if @a comes after @b.
//
// The @a and @b iterators must point into the same sequence.
/*

C function : g_sequence_iter_compare
*/
func (recv *SequenceIter) Compare(b *SequenceIter) int32 {
	c_b := (*C.GSequenceIter)(C.NULL)
	if b != nil {
		c_b = (*C.GSequenceIter)(b.ToC())
	}

	retC := C.g_sequence_iter_compare((*C.GSequenceIter)(recv.native), c_b)
	retGo := (int32)(retC)

	return retGo
}

// Returns the position of @iter
/*

C function : g_sequence_iter_get_position
*/
func (recv *SequenceIter) GetPosition() int32 {
	retC := C.g_sequence_iter_get_position((*C.GSequenceIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the #GSequence that @iter points into.
/*

C function : g_sequence_iter_get_sequence
*/
func (recv *SequenceIter) GetSequence() *Sequence {
	retC := C.g_sequence_iter_get_sequence((*C.GSequenceIter)(recv.native))
	retGo := SequenceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether @iter is the begin iterator
/*

C function : g_sequence_iter_is_begin
*/
func (recv *SequenceIter) IsBegin() bool {
	retC := C.g_sequence_iter_is_begin((*C.GSequenceIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether @iter is the end iterator
/*

C function : g_sequence_iter_is_end
*/
func (recv *SequenceIter) IsEnd() bool {
	retC := C.g_sequence_iter_is_end((*C.GSequenceIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the #GSequenceIter which is @delta positions away from @iter.
// If @iter is closer than -@delta positions to the beginning of the sequence,
// the begin iterator is returned. If @iter is closer than @delta positions
// to the end of the sequence, the end iterator is returned.
/*

C function : g_sequence_iter_move
*/
func (recv *SequenceIter) Move(delta int32) *SequenceIter {
	c_delta := (C.gint)(delta)

	retC := C.g_sequence_iter_move((*C.GSequenceIter)(recv.native), c_delta)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns an iterator pointing to the next position after @iter.
// If @iter is the end iterator, the end iterator is returned.
/*

C function : g_sequence_iter_next
*/
func (recv *SequenceIter) Next() *SequenceIter {
	retC := C.g_sequence_iter_next((*C.GSequenceIter)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns an iterator pointing to the previous position before @iter.
// If @iter is the begin iterator, the begin iterator is returned.
/*

C function : g_sequence_iter_prev
*/
func (recv *SequenceIter) Prev() *SequenceIter {
	retC := C.g_sequence_iter_prev((*C.GSequenceIter)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_string_append_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Overwrites part of a string, lengthening it if necessary.
/*

C function : g_string_overwrite
*/
func (recv *String) Overwrite(pos uint64, val string) *String {
	c_pos := (C.gsize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_overwrite((*C.GString)(recv.native), c_pos, c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Overwrites part of a string, lengthening it if necessary.
// This function will work with embedded nuls.
/*

C function : g_string_overwrite_len
*/
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

// Frees all strings contained within the #GStringChunk.
// After calling g_string_chunk_clear() it is not safe to
// access any of the strings which were contained within it.
/*

C function : g_string_chunk_clear
*/
func (recv *StringChunk) Clear() {
	C.g_string_chunk_clear((*C.GStringChunk)(recv.native))

	return
}
