// This is a generated file - DO NOT EDIT

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// TypePlugin is a wrapper around the C record GTypePlugin.
type TypePlugin struct {
	native *C.GTypePlugin
}

func TypePluginNewFromC(u unsafe.Pointer) *TypePlugin {
	c := (*C.GTypePlugin)(u)
	if c == nil {
		return nil
	}

	g := &TypePlugin{native: c}

	return g
}

func (recv *TypePlugin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_type_plugin_complete_interface_info : unsupported parameter instance_type : no type generator for GType, GType

// Unsupported : g_type_plugin_complete_type_info : unsupported parameter g_type : no type generator for GType, GType

// Unuse is a wrapper around the C function g_type_plugin_unuse.
func (recv *TypePlugin) Unuse() {
	C.g_type_plugin_unuse((*C.GTypePlugin)(recv.native))

	return
}

// Use is a wrapper around the C function g_type_plugin_use.
func (recv *TypePlugin) Use() {
	C.g_type_plugin_use((*C.GTypePlugin)(recv.native))

	return
}
