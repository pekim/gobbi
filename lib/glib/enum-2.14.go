// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

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
