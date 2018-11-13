// This is a generated file - DO NOT EDIT
// +build atk_2.8 atk_2.12

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Returns the binary age as passed to libtool when building the ATK
// library the process is running against.
/*

C function : atk_get_binary_age
*/
func GetBinaryAge() uint32 {
	retC := C.atk_get_binary_age()
	retGo := (uint32)(retC)

	return retGo
}

// Returns the interface age as passed to libtool when building the
// ATK library the process is running against.
/*

C function : atk_get_interface_age
*/
func GetInterfaceAge() uint32 {
	retC := C.atk_get_interface_age()
	retGo := (uint32)(retC)

	return retGo
}

// Returns the major version number of the ATK library.  (e.g. in ATK
// version 2.7.4 this is 2.)
//
// This function is in the library, so it represents the ATK library
// your code is running against. In contrast, the #ATK_MAJOR_VERSION
// macro represents the major version of the ATK headers you have
// included when compiling your code.
/*

C function : atk_get_major_version
*/
func GetMajorVersion() uint32 {
	retC := C.atk_get_major_version()
	retGo := (uint32)(retC)

	return retGo
}

// Returns the micro version number of the ATK library.  (e.g. in ATK
// version 2.7.4 this is 4.)
//
// This function is in the library, so it represents the ATK library
// your code is are running against. In contrast, the
// #ATK_MICRO_VERSION macro represents the micro version of the ATK
// headers you have included when compiling your code.
/*

C function : atk_get_micro_version
*/
func GetMicroVersion() uint32 {
	retC := C.atk_get_micro_version()
	retGo := (uint32)(retC)

	return retGo
}

// Returns the minor version number of the ATK library.  (e.g. in ATK
// version 2.7.4 this is 7.)
//
// This function is in the library, so it represents the ATK library
// your code is are running against. In contrast, the
// #ATK_MINOR_VERSION macro represents the minor version of the ATK
// headers you have included when compiling your code.
/*

C function : atk_get_minor_version
*/
func GetMinorVersion() uint32 {
	retC := C.atk_get_minor_version()
	retGo := (uint32)(retC)

	return retGo
}
