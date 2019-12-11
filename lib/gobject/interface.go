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

var typePluginCompleteInterfaceInfoFunction *gi.Function
var typePluginCompleteInterfaceInfoFunction_Once sync.Once

func typePluginCompleteInterfaceInfoFunction_Set() error {
	var err error
	typePluginCompleteInterfaceInfoFunction_Once.Do(func() {
		err = typePluginInterface_Set()
		if err != nil {
			return
		}
		typePluginCompleteInterfaceInfoFunction, err = typePluginInterface.InvokerNew("complete_interface_info")
	})
	return err
}

// CompleteInterfaceInfo is a representation of the C type g_type_plugin_complete_interface_info.
func (recv *TypePlugin) CompleteInterfaceInfo(instanceType int64, interfaceType int64, info *InterfaceInfo) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(instanceType)
	inArgs[2].SetInt64(interfaceType)
	inArgs[3].SetPointer(info.Native())

	err := typePluginCompleteInterfaceInfoFunction_Set()
	if err == nil {
		typePluginCompleteInterfaceInfoFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typePluginCompleteTypeInfoFunction *gi.Function
var typePluginCompleteTypeInfoFunction_Once sync.Once

func typePluginCompleteTypeInfoFunction_Set() error {
	var err error
	typePluginCompleteTypeInfoFunction_Once.Do(func() {
		err = typePluginInterface_Set()
		if err != nil {
			return
		}
		typePluginCompleteTypeInfoFunction, err = typePluginInterface.InvokerNew("complete_type_info")
	})
	return err
}

// CompleteTypeInfo is a representation of the C type g_type_plugin_complete_type_info.
func (recv *TypePlugin) CompleteTypeInfo(gType int64, info *TypeInfo, valueTable *TypeValueTable) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(gType)
	inArgs[2].SetPointer(info.Native())
	inArgs[3].SetPointer(valueTable.Native())

	err := typePluginCompleteTypeInfoFunction_Set()
	if err == nil {
		typePluginCompleteTypeInfoFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
