// This is a generated file - DO NOT EDIT

package gobject

import "unsafe"

// TypePlugin is a wrapper around the C record GTypePlugin.
type TypePlugin struct {
	native unsafe.Pointer
}

func TypePluginNewFromC(u unsafe.Pointer) *TypePlugin {
	if u == nil {
		return nil
	}

	g := &TypePlugin{native: u}

	return g
}

func (recv *TypePlugin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
