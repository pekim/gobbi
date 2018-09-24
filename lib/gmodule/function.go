// This is a generated file - DO NOT EDIT

package gmodule

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <gmodule.h>
// #include <stdlib.h>
import "C"

// ModuleBuildPath is a wrapper around the C function g_module_build_path.
func ModuleBuildPath(directory string, moduleName string) string {
	c_directory := C.CString(directory)
	defer C.free(unsafe.Pointer(c_directory))

	c_module_name := C.CString(moduleName)
	defer C.free(unsafe.Pointer(c_module_name))

	retC := C.g_module_build_path(c_directory, c_module_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ModuleError is a wrapper around the C function g_module_error.
func ModuleError() string {
	retC := C.g_module_error()
	retGo := C.GoString(retC)

	return retGo
}

// ModuleSupported is a wrapper around the C function g_module_supported.
func ModuleSupported() bool {
	retC := C.g_module_supported()
	retGo := (bool)(retC)

	return retGo
}
