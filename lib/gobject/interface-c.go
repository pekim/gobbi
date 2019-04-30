// This is a generated file - DO NOT EDIT

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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
