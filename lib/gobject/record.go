// Code generated - DO NOT EDIT.

package gobject

import gi "github.com/pekim/gobbi/internal/gi"

type CClosure struct {
	native  uintptr
	Closure *Closure
	// UNSUPPORTED : C value 'callback' : no Go type for 'gpointer'

}

type Closure struct {
	native    uintptr
	InMarshal uint32
	IsInvalid uint32
	// UNSUPPORTED : C value 'marshal' : missing Type

}

// UNSUPPORTED : C value 'g_closure_new_object' : parameter 'object' of type 'Object' not supported

// UNSUPPORTED : C value 'g_closure_new_simple' : parameter 'data' of type 'gpointer' not supported

type ClosureNotifyData struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'

	// UNSUPPORTED : C value 'notify' : no Go type for 'ClosureNotify'

}

type EnumClass struct {
	native     uintptr
	GTypeClass *TypeClass
	Minimum    int32
	Maximum    int32
	NValues    uint32
	Values     *EnumValue
}

type EnumValue struct {
	native    uintptr
	Value     int32
	ValueName string
	ValueNick string
}

type FlagsClass struct {
	native     uintptr
	GTypeClass *TypeClass
	Mask       uint32
	NValues    uint32
	Values     *FlagsValue
}

type FlagsValue struct {
	native    uintptr
	Value     uint32
	ValueName string
	ValueNick string
}

type InitiallyUnownedClass struct {
	native     uintptr
	GTypeClass *TypeClass
	// UNSUPPORTED : C value 'constructor' : missing Type

	// UNSUPPORTED : C value 'set_property' : missing Type

	// UNSUPPORTED : C value 'get_property' : missing Type

	// UNSUPPORTED : C value 'dispose' : missing Type

	// UNSUPPORTED : C value 'finalize' : missing Type

	// UNSUPPORTED : C value 'dispatch_properties_changed' : missing Type

	// UNSUPPORTED : C value 'notify' : missing Type

	// UNSUPPORTED : C value 'constructed' : missing Type

}

type InterfaceInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'interface_init' : no Go type for 'InterfaceInitFunc'

	// UNSUPPORTED : C value 'interface_finalize' : no Go type for 'InterfaceFinalizeFunc'

	// UNSUPPORTED : C value 'interface_data' : no Go type for 'gpointer'

}

type ObjectClass struct {
	native     uintptr
	GTypeClass *TypeClass
	// UNSUPPORTED : C value 'constructor' : missing Type

	// UNSUPPORTED : C value 'set_property' : missing Type

	// UNSUPPORTED : C value 'get_property' : missing Type

	// UNSUPPORTED : C value 'dispose' : missing Type

	// UNSUPPORTED : C value 'finalize' : missing Type

	// UNSUPPORTED : C value 'dispatch_properties_changed' : missing Type

	// UNSUPPORTED : C value 'notify' : missing Type

	// UNSUPPORTED : C value 'constructed' : missing Type

}

type ObjectConstructParam struct {
	native uintptr
	// UNSUPPORTED : C value 'pspec' : no Go type for 'ParamSpec'

	Value *Value
}

type ParamSpecClass struct {
	native     uintptr
	GTypeClass *TypeClass
	// UNSUPPORTED : C value 'value_type' : no Go type for 'GType'

	// UNSUPPORTED : C value 'finalize' : missing Type

	// UNSUPPORTED : C value 'value_set_default' : missing Type

	// UNSUPPORTED : C value 'value_validate' : missing Type

	// UNSUPPORTED : C value 'values_cmp' : missing Type

}

type ParamSpecPool struct {
	native uintptr
}

type ParamSpecTypeInfo struct {
	native       uintptr
	InstanceSize uint16
	NPreallocs   uint16
	// UNSUPPORTED : C value 'instance_init' : missing Type

	// UNSUPPORTED : C value 'value_type' : no Go type for 'GType'

	// UNSUPPORTED : C value 'finalize' : missing Type

	// UNSUPPORTED : C value 'value_set_default' : missing Type

	// UNSUPPORTED : C value 'value_validate' : missing Type

	// UNSUPPORTED : C value 'values_cmp' : missing Type

}

type Parameter struct {
	native uintptr
	Name   string
	Value  *Value
}

type SignalInvocationHint struct {
	native   uintptr
	SignalId uint32
	// UNSUPPORTED : C value 'detail' : no Go type for 'GLib.Quark'

	// UNSUPPORTED : C value 'run_type' : no Go type for 'SignalFlags'

}

type SignalQuery struct {
	native     uintptr
	SignalId   uint32
	SignalName string
	// UNSUPPORTED : C value 'itype' : no Go type for 'GType'

	// UNSUPPORTED : C value 'signal_flags' : no Go type for 'SignalFlags'

	// UNSUPPORTED : C value 'return_type' : no Go type for 'GType'

	NParams uint32
	// UNSUPPORTED : C value 'param_types' : missing Type

}

type TypeClass struct {
	native uintptr
}

type TypeFundamentalInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'type_flags' : no Go type for 'TypeFundamentalFlags'

}

type TypeInfo struct {
	native    uintptr
	ClassSize uint16
	// UNSUPPORTED : C value 'base_init' : no Go type for 'BaseInitFunc'

	// UNSUPPORTED : C value 'base_finalize' : no Go type for 'BaseFinalizeFunc'

	// UNSUPPORTED : C value 'class_init' : no Go type for 'ClassInitFunc'

	// UNSUPPORTED : C value 'class_finalize' : no Go type for 'ClassFinalizeFunc'

	// UNSUPPORTED : C value 'class_data' : no Go type for 'gpointer'

	InstanceSize uint16
	NPreallocs   uint16
	// UNSUPPORTED : C value 'instance_init' : no Go type for 'InstanceInitFunc'

	ValueTable *TypeValueTable
}

type TypeInstance struct {
	native uintptr
}

type TypeInterface struct {
	native uintptr
}

type TypeModuleClass struct {
	native      uintptr
	ParentClass *ObjectClass
	// UNSUPPORTED : C value 'load' : missing Type

	// UNSUPPORTED : C value 'unload' : missing Type

	// UNSUPPORTED : C value 'reserved1' : missing Type

	// UNSUPPORTED : C value 'reserved2' : missing Type

	// UNSUPPORTED : C value 'reserved3' : missing Type

	// UNSUPPORTED : C value 'reserved4' : missing Type

}

type TypePluginClass struct {
	native uintptr
	// UNSUPPORTED : C value 'use_plugin' : no Go type for 'TypePluginUse'

	// UNSUPPORTED : C value 'unuse_plugin' : no Go type for 'TypePluginUnuse'

	// UNSUPPORTED : C value 'complete_type_info' : no Go type for 'TypePluginCompleteTypeInfo'

	// UNSUPPORTED : C value 'complete_interface_info' : no Go type for 'TypePluginCompleteInterfaceInfo'

}

type TypeQuery struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'GType'

	TypeName     string
	ClassSize    uint32
	InstanceSize uint32
}

type TypeValueTable struct {
	native uintptr
	// UNSUPPORTED : C value 'value_init' : missing Type

	// UNSUPPORTED : C value 'value_free' : missing Type

	// UNSUPPORTED : C value 'value_copy' : missing Type

	// UNSUPPORTED : C value 'value_peek_pointer' : missing Type

	CollectFormat string
	// UNSUPPORTED : C value 'collect_value' : missing Type

	LcopyFormat string
	// UNSUPPORTED : C value 'lcopy_value' : missing Type

}

type Value struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : missing Type

}

type ValueArray struct {
	native  uintptr
	NValues uint32
	Values  *Value
}

var newValueArrayInvoker *gi.Function

// ValueArrayNew is a representation of the C type g_value_array_new.
func ValueArrayNew(nPrealloced uint32) *ValueArray {
	if newValueArrayInvoker == nil {
		newValueArrayInvoker = gi.StructFunctionInvokerNew("GObject", "ValueArray", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(nPrealloced)

	ret := newValueArrayInvoker.Invoke(inArgs[:], nil)

	return &ValueArray{native: ret.Pointer()}
}

type WeakRef struct {
	native uintptr
}
