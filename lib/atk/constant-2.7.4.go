// This is a generated file - DO NOT EDIT
// +build atk_2.8 atk_2.12

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Like atk_get_binary_age(), but from the headers used at
// application compile time, rather than from the library linked
// against at application run time.
const BINARY_AGE int = C.ATK_BINARY_AGE

// Like atk_get_interface_age(), but from the headers used at
// application compile time, rather than from the library linked
// against at application run time.
const INTERFACE_AGE int = C.ATK_INTERFACE_AGE

// Like atk_get_major_version(), but from the headers used at
// application compile time, rather than from the library linked
// against at application run time.
const MAJOR_VERSION int = C.ATK_MAJOR_VERSION

// Like atk_get_micro_version(), but from the headers used at
// application compile time, rather than from the library linked
// against at application run time.
const MICRO_VERSION int = C.ATK_MICRO_VERSION

// Like atk_get_minor_version(), but from the headers used at
// application compile time, rather than from the library linked
// against at application run time.
const MINOR_VERSION int = C.ATK_MINOR_VERSION
