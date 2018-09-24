// This is a generated file - DO NOT EDIT

package gmodule

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <gmodule.h>
// #include <stdlib.h>
import "C"

// Module is a wrapper around the C record GModule.
type Module struct {
	native *C.GModule
}

func ModuleNewFromC(u unsafe.Pointer) *Module {
	c := (*C.GModule)(u)
	if c == nil {
		return nil
	}

	g := &Module{native: c}

	return g
}

func (recv *Module) toC() *C.GModule {

	return recv.native
}

// Unsupported : g_module_close : no return generator

// Unsupported : g_module_make_resident : no return generator

// Name is a wrapper around the C function g_module_name.
func (recv *Module) Name() string {
	retC := C.g_module_name((*C.GModule)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_module_symbol : unsupported parameter symbol : no type generator for gpointer, gpointer*
