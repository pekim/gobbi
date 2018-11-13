// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Flags specifying compile-time options.
type RegexCompileFlags C.GRegexCompileFlags

const (
	/*
	   Letters in the pattern match both upper- and
	       lowercase letters. This option can be changed within a pattern
	       by a "(?i)" option setting.
	*/
	REGEX_CASELESS RegexCompileFlags = 1
	/*
	   By default, GRegex treats the strings as consisting
	       of a single line of characters (even if it actually contains
	       newlines). The "start of line" metacharacter ("^") matches only
	       at the start of the string, while the "end of line" metacharacter
	       ("$") matches only at the end of the string, or before a terminating
	       newline (unless #G_REGEX_DOLLAR_ENDONLY is set). When
	       #G_REGEX_MULTILINE is set, the "start of line" and "end of line"
	       constructs match immediately following or immediately before any
	       newline in the string, respectively, as well as at the very start
	       and end. This can be changed within a pattern by a "(?m)" option
	       setting.
	*/
	REGEX_MULTILINE RegexCompileFlags = 2
	/*
	   A dot metacharacter (".") in the pattern matches all
	       characters, including newlines. Without it, newlines are excluded.
	       This option can be changed within a pattern by a ("?s") option setting.
	*/
	REGEX_DOTALL RegexCompileFlags = 4
	/*
	   Whitespace data characters in the pattern are
	       totally ignored except when escaped or inside a character class.
	       Whitespace does not include the VT character (code 11). In addition,
	       characters between an unescaped "#" outside a character class and
	       the next newline character, inclusive, are also ignored. This can
	       be changed within a pattern by a "(?x)" option setting.
	*/
	REGEX_EXTENDED RegexCompileFlags = 8
	/*
	   The pattern is forced to be "anchored", that is,
	       it is constrained to match only at the first matching point in the
	       string that is being searched. This effect can also be achieved by
	       appropriate constructs in the pattern itself such as the "^"
	       metacharacter.
	*/
	REGEX_ANCHORED RegexCompileFlags = 16
	/*
	   A dollar metacharacter ("$") in the pattern
	       matches only at the end of the string. Without this option, a
	       dollar also matches immediately before the final character if
	       it is a newline (but not before any other newlines). This option
	       is ignored if #G_REGEX_MULTILINE is set.
	*/
	REGEX_DOLLAR_ENDONLY RegexCompileFlags = 32
	/*
	   Inverts the "greediness" of the quantifiers so that
	       they are not greedy by default, but become greedy if followed by "?".
	       It can also be set by a "(?U)" option setting within the pattern.
	*/
	REGEX_UNGREEDY RegexCompileFlags = 512
	/*
	   Usually strings must be valid UTF-8 strings, using this
	       flag they are considered as a raw sequence of bytes.
	*/
	REGEX_RAW RegexCompileFlags = 2048
	/*
	   Disables the use of numbered capturing
	       parentheses in the pattern. Any opening parenthesis that is not
	       followed by "?" behaves as if it were followed by "?:" but named
	       parentheses can still be used for capturing (and they acquire numbers
	       in the usual way).
	*/
	REGEX_NO_AUTO_CAPTURE RegexCompileFlags = 4096
	/*
	   Optimize the regular expression. If the pattern will
	       be used many times, then it may be worth the effort to optimize it
	       to improve the speed of matches.
	*/
	REGEX_OPTIMIZE RegexCompileFlags = 8192
	/*
	   Limits an unanchored pattern to match before (or at) the
	       first newline. Since: 2.34
	*/
	REGEX_FIRSTLINE RegexCompileFlags = 262144
	/*
	   Names used to identify capturing subpatterns need not
	       be unique. This can be helpful for certain types of pattern when it
	       is known that only one instance of the named subpattern can ever be
	       matched.
	*/
	REGEX_DUPNAMES RegexCompileFlags = 524288
	/*
	   Usually any newline character or character sequence is
	       recognized. If this option is set, the only recognized newline character
	       is '\r'.
	*/
	REGEX_NEWLINE_CR RegexCompileFlags = 1048576
	/*
	   Usually any newline character or character sequence is
	       recognized. If this option is set, the only recognized newline character
	       is '\n'.
	*/
	REGEX_NEWLINE_LF RegexCompileFlags = 2097152
	/*
	   Usually any newline character or character sequence is
	       recognized. If this option is set, the only recognized newline character
	       sequence is '\r\n'.
	*/
	REGEX_NEWLINE_CRLF RegexCompileFlags = 3145728
	/*
	   Usually any newline character or character sequence
	       is recognized. If this option is set, the only recognized newline character
	       sequences are '\r', '\n', and '\r\n'. Since: 2.34
	*/
	REGEX_NEWLINE_ANYCRLF RegexCompileFlags = 5242880
	/*
	   Usually any newline character or character sequence
	       is recognised. If this option is set, then "\R" only recognizes the newline
	      characters '\r', '\n' and '\r\n'. Since: 2.34
	*/
	REGEX_BSR_ANYCRLF RegexCompileFlags = 8388608
	/*
	   Changes behaviour so that it is compatible with
	       JavaScript rather than PCRE. Since: 2.34
	*/
	REGEX_JAVASCRIPT_COMPAT RegexCompileFlags = 33554432
)

// Flags specifying match-time options.
type RegexMatchFlags C.GRegexMatchFlags

const (
	/*
	   The pattern is forced to be "anchored", that is,
	       it is constrained to match only at the first matching point in the
	       string that is being searched. This effect can also be achieved by
	       appropriate constructs in the pattern itself such as the "^"
	       metacharacter.
	*/
	REGEX_MATCH_ANCHORED RegexMatchFlags = 16
	/*
	   Specifies that first character of the string is
	       not the beginning of a line, so the circumflex metacharacter should
	       not match before it. Setting this without #G_REGEX_MULTILINE (at
	       compile time) causes circumflex never to match. This option affects
	       only the behaviour of the circumflex metacharacter, it does not
	       affect "\A".
	*/
	REGEX_MATCH_NOTBOL RegexMatchFlags = 128
	/*
	   Specifies that the end of the subject string is
	       not the end of a line, so the dollar metacharacter should not match
	       it nor (except in multiline mode) a newline immediately before it.
	       Setting this without #G_REGEX_MULTILINE (at compile time) causes
	       dollar never to match. This option affects only the behaviour of
	       the dollar metacharacter, it does not affect "\Z" or "\z".
	*/
	REGEX_MATCH_NOTEOL RegexMatchFlags = 256
	/*
	   An empty string is not considered to be a valid
	       match if this option is set. If there are alternatives in the pattern,
	       they are tried. If all the alternatives match the empty string, the
	       entire match fails. For example, if the pattern "a?b?" is applied to
	       a string not beginning with "a" or "b", it matches the empty string
	       at the start of the string. With this flag set, this match is not
	       valid, so GRegex searches further into the string for occurrences
	       of "a" or "b".
	*/
	REGEX_MATCH_NOTEMPTY RegexMatchFlags = 1024
	/*
	   Turns on the partial matching feature, for more
	       documentation on partial matching see g_match_info_is_partial_match().
	*/
	REGEX_MATCH_PARTIAL RegexMatchFlags = 32768
	/*
	   Overrides the newline definition set when
	       creating a new #GRegex, setting the '\r' character as line terminator.
	*/
	REGEX_MATCH_NEWLINE_CR RegexMatchFlags = 1048576
	/*
	   Overrides the newline definition set when
	       creating a new #GRegex, setting the '\n' character as line terminator.
	*/
	REGEX_MATCH_NEWLINE_LF RegexMatchFlags = 2097152
	/*
	   Overrides the newline definition set when
	       creating a new #GRegex, setting the '\r\n' characters sequence as line terminator.
	*/
	REGEX_MATCH_NEWLINE_CRLF RegexMatchFlags = 3145728
	/*
	   Overrides the newline definition set when
	       creating a new #GRegex, any Unicode newline sequence
	       is recognised as a newline. These are '\r', '\n' and '\rn', and the
	       single characters U+000B LINE TABULATION, U+000C FORM FEED (FF),
	       U+0085 NEXT LINE (NEL), U+2028 LINE SEPARATOR and
	       U+2029 PARAGRAPH SEPARATOR.
	*/
	REGEX_MATCH_NEWLINE_ANY RegexMatchFlags = 4194304
	/*
	   Overrides the newline definition set when
	       creating a new #GRegex; any '\r', '\n', or '\r\n' character sequence
	       is recognized as a newline. Since: 2.34
	*/
	REGEX_MATCH_NEWLINE_ANYCRLF RegexMatchFlags = 5242880
	/*
	   Overrides the newline definition for "\R" set when
	       creating a new #GRegex; only '\r', '\n', or '\r\n' character sequences
	       are recognized as a newline by "\R". Since: 2.34
	*/
	REGEX_MATCH_BSR_ANYCRLF RegexMatchFlags = 8388608
	/*
	   Overrides the newline definition for "\R" set when
	       creating a new #GRegex; any Unicode newline character or character sequence
	       are recognized as a newline by "\R". These are '\r', '\n' and '\rn', and the
	       single characters U+000B LINE TABULATION, U+000C FORM FEED (FF),
	       U+0085 NEXT LINE (NEL), U+2028 LINE SEPARATOR and
	       U+2029 PARAGRAPH SEPARATOR. Since: 2.34
	*/
	REGEX_MATCH_BSR_ANY RegexMatchFlags = 16777216
	// An alias for #G_REGEX_MATCH_PARTIAL. Since: 2.34
	REGEX_MATCH_PARTIAL_SOFT RegexMatchFlags = 32768
	/*
	   Turns on the partial matching feature. In contrast to
	       to #G_REGEX_MATCH_PARTIAL_SOFT, this stops matching as soon as a partial match
	       is found, without continuing to search for a possible complete match. See
	       g_match_info_is_partial_match() for more information. Since: 2.34
	*/
	REGEX_MATCH_PARTIAL_HARD RegexMatchFlags = 134217728
	/*
	   Like #G_REGEX_MATCH_NOTEMPTY, but only applied to
	       the start of the matched string. For anchored
	       patterns this can only happen for pattern containing "\K". Since: 2.34
	*/
	REGEX_MATCH_NOTEMPTY_ATSTART RegexMatchFlags = 268435456
)
