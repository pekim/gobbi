// This is a generated file - DO NOT EDIT

package gmodule

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <gmodule.h>
// #include <stdlib.h>
import "C"

type ModuleFlags C.GModuleFlags

const (
	MODULE_BIND_LAZY  ModuleFlags = 1
	MODULE_BIND_LOCAL ModuleFlags = 2
	MODULE_BIND_MASK  ModuleFlags = 3
)
