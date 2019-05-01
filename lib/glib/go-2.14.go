// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

type RegexCompileFlags int

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

type RegexMatchFlags int

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
const KEY_FILE_DESKTOP_GROUP string = "Desktop Entry"
const KEY_FILE_DESKTOP_KEY_CATEGORIES string = "Categories"
const KEY_FILE_DESKTOP_KEY_COMMENT string = "Comment"
const KEY_FILE_DESKTOP_KEY_EXEC string = "Exec"
const KEY_FILE_DESKTOP_KEY_GENERIC_NAME string = "GenericName"
const KEY_FILE_DESKTOP_KEY_HIDDEN string = "Hidden"
const KEY_FILE_DESKTOP_KEY_ICON string = "Icon"
const KEY_FILE_DESKTOP_KEY_MIME_TYPE string = "MimeType"
const KEY_FILE_DESKTOP_KEY_NAME string = "Name"
const KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN string = "NotShowIn"
const KEY_FILE_DESKTOP_KEY_NO_DISPLAY string = "NoDisplay"
const KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN string = "OnlyShowIn"
const KEY_FILE_DESKTOP_KEY_PATH string = "Path"
const KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY string = "StartupNotify"
const KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS string = "StartupWMClass"
const KEY_FILE_DESKTOP_KEY_TERMINAL string = "Terminal"
const KEY_FILE_DESKTOP_KEY_TRY_EXEC string = "TryExec"
const KEY_FILE_DESKTOP_KEY_TYPE string = "Type"
const KEY_FILE_DESKTOP_KEY_URL string = "URL"
const KEY_FILE_DESKTOP_KEY_VERSION string = "Version"
const KEY_FILE_DESKTOP_TYPE_APPLICATION string = "Application"
const KEY_FILE_DESKTOP_TYPE_DIRECTORY string = "Directory"
const KEY_FILE_DESKTOP_TYPE_LINK string = "Link"

type RegexError int

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

type UserDirectory int

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

// Unsupported : g_get_user_special_dir : return type :

// Unsupported : g_once_init_enter : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_once_init_leave : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_regex_check_replacement : return type :

// Unsupported : g_regex_escape_string : return type :

// Unsupported : g_regex_match_simple : return type :

// Unsupported : g_regex_split_simple : array return type :

// Unsupported : g_sequence_get : no return generator

// Unsupported : g_sequence_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_range_get_midpoint : return type :

// Unsupported : g_sequence_set : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_slice_copy : unsupported parameter mem_block : no type generator for gpointer (gconstpointer) for param mem_block

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_source_new_seconds : return type :

// Unsupported : g_unichar_get_script : return type :

// Unsupported : g_unichar_ismark : return type :

// Unsupported : g_unichar_iszerowidth : return type :

// Blacklisted : GRegex
