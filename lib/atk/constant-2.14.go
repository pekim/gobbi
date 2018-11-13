// This is a generated file - DO NOT EDIT

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// A macro that should be defined by the user prior to including
// the atk/atk.h header.
// The definition should be one of the predefined ATK version
// macros: %ATK_VERSION_2_12, %ATK_VERSION_2_14,...
//
// This macro defines the earliest version of ATK that the package is
// required to be able to compile against.
//
// If the compiler is configured to warn about the use of deprecated
// functions, then using functions that were deprecated in version
// %ATK_VERSION_MIN_REQUIRED or earlier will cause warnings (but
// using functions deprecated in later releases will not).
const VERSION_MIN_REQUIRED int = C.ATK_VERSION_MIN_REQUIRED
