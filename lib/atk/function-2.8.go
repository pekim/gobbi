// This is a generated file - DO NOT EDIT
// +build atk_2.8

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// GetBinaryAge is a wrapper around the C function atk_get_binary_age.
func GetBinaryAge() uint32 {
	retC := C.atk_get_binary_age()
	retGo := (uint32)(retC)

	return retGo
}

// GetInterfaceAge is a wrapper around the C function atk_get_interface_age.
func GetInterfaceAge() uint32 {
	retC := C.atk_get_interface_age()
	retGo := (uint32)(retC)

	return retGo
}

// GetMajorVersion is a wrapper around the C function atk_get_major_version.
func GetMajorVersion() uint32 {
	retC := C.atk_get_major_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMicroVersion is a wrapper around the C function atk_get_micro_version.
func GetMicroVersion() uint32 {
	retC := C.atk_get_micro_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMinorVersion is a wrapper around the C function atk_get_minor_version.
func GetMinorVersion() uint32 {
	retC := C.atk_get_minor_version()
	retGo := (uint32)(retC)

	return retGo
}