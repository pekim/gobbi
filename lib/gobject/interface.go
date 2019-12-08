// Code generated - DO NOT EDIT.

package gobject

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	"sync"
	"unsafe"
)

var typePluginInterface *gi.Interface
var typePluginInterface_Once sync.Once

func typePluginInterface_Set() error {
	var err error
	typePluginInterface_Once.Do(func() {
		typePluginInterface, err = gi.InterfaceNew("GObject", "TypePlugin")
	})
	return err
}

type TypePlugin struct {
	native unsafe.Pointer
}

func TypePluginNewFromNative(native unsafe.Pointer) *TypePlugin {
	instance := &TypePlugin{native: native}

	return instance
}

// Equals compares this TypePlugin with another TypePlugin, and returns true if they represent the same Object.
func (recv *TypePlugin) Equals(other *TypePlugin) bool {
	return other.Native() == recv.Native()
}

func (recv *TypePlugin) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'g_type_plugin_complete_interface_info' : parameter 'instance_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_plugin_complete_type_info' : parameter 'g_type' of type 'GType' not supported

var typePluginUnuseFunction *gi.Function
var typePluginUnuseFunction_Once sync.Once

func typePluginUnuseFunction_Set() error {
	var err error
	typePluginUnuseFunction_Once.Do(func() {
		err = typePluginInterface_Set()
		if err != nil {
			return
		}
		typePluginUnuseFunction, err = typePluginInterface.InvokerNew("unuse")
	})
	return err
}

// Unuse is a representation of the C type g_type_plugin_unuse.
func (recv *TypePlugin) Unuse() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := typePluginUnuseFunction_Set()
	if err == nil {
		typePluginUnuseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typePluginUseFunction *gi.Function
var typePluginUseFunction_Once sync.Once

func typePluginUseFunction_Set() error {
	var err error
	typePluginUseFunction_Once.Do(func() {
		err = typePluginInterface_Set()
		if err != nil {
			return
		}
		typePluginUseFunction, err = typePluginInterface.InvokerNew("use")
	})
	return err
}

// Use is a representation of the C type g_type_plugin_use.
func (recv *TypePlugin) Use() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := typePluginUseFunction_Set()
	if err == nil {
		typePluginUseFunction.Invoke(inArgs[:], nil)
	}

	return
}
