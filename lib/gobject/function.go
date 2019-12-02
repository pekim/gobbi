// Code generated - DO NOT EDIT.

package gobject

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'g_boxed_copy' : parameter 'boxed_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_boxed_free' : parameter 'boxed_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_boxed_type_register_static' : parameter 'boxed_copy' of type 'BoxedCopyFunc' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_BOOLEAN__BOXED_BOXED' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_BOOLEAN__FLAGS' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_STRING__OBJECT_POINTER' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__BOOLEAN' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__BOXED' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__CHAR' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__DOUBLE' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__ENUM' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__FLAGS' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__FLOAT' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__INT' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__LONG' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__OBJECT' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__PARAM' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__POINTER' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__STRING' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__UCHAR' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__UINT' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__UINT_POINTER' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__ULONG' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__VARIANT' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_VOID__VOID' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_marshal_generic' : parameter 'invocation_hint' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_cclosure_new' : parameter 'callback_func' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_cclosure_new_object' : parameter 'callback_func' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_cclosure_new_object_swap' : parameter 'callback_func' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_cclosure_new_swap' : parameter 'callback_func' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_clear_object' : parameter 'object_ptr' of type 'Object' not supported

// UNSUPPORTED : C value 'g_clear_signal_handler' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_enum_complete_type_info' : parameter 'g_enum_type' of type 'GType' not supported

var enumGetValueFunction *gi.Function
var enumGetValueFunction_Once sync.Once

func enumGetValueFunction_Set() error {
	var err error
	enumGetValueFunction_Once.Do(func() {
		enumGetValueFunction, err = gi.FunctionInvokerNew("GObject", "enum_get_value")
	})
	return err
}

// EnumGetValue is a representation of the C type g_enum_get_value.
func EnumGetValue(enumClass *EnumClass, value int32) *EnumValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(enumClass.native)
	inArgs[1].SetInt32(value)

	var ret gi.Argument

	err := enumGetValueFunction_Set()
	if err == nil {
		ret = enumGetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := &EnumValue{native: ret.Pointer()}

	return retGo
}

var enumGetValueByNameFunction *gi.Function
var enumGetValueByNameFunction_Once sync.Once

func enumGetValueByNameFunction_Set() error {
	var err error
	enumGetValueByNameFunction_Once.Do(func() {
		enumGetValueByNameFunction, err = gi.FunctionInvokerNew("GObject", "enum_get_value_by_name")
	})
	return err
}

// EnumGetValueByName is a representation of the C type g_enum_get_value_by_name.
func EnumGetValueByName(enumClass *EnumClass, name string) *EnumValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(enumClass.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := enumGetValueByNameFunction_Set()
	if err == nil {
		ret = enumGetValueByNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := &EnumValue{native: ret.Pointer()}

	return retGo
}

var enumGetValueByNickFunction *gi.Function
var enumGetValueByNickFunction_Once sync.Once

func enumGetValueByNickFunction_Set() error {
	var err error
	enumGetValueByNickFunction_Once.Do(func() {
		enumGetValueByNickFunction, err = gi.FunctionInvokerNew("GObject", "enum_get_value_by_nick")
	})
	return err
}

// EnumGetValueByNick is a representation of the C type g_enum_get_value_by_nick.
func EnumGetValueByNick(enumClass *EnumClass, nick string) *EnumValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(enumClass.native)
	inArgs[1].SetString(nick)

	var ret gi.Argument

	err := enumGetValueByNickFunction_Set()
	if err == nil {
		ret = enumGetValueByNickFunction.Invoke(inArgs[:], nil)
	}

	retGo := &EnumValue{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_enum_register_static' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_enum_to_string' : parameter 'g_enum_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_flags_complete_type_info' : parameter 'g_flags_type' of type 'GType' not supported

var flagsGetFirstValueFunction *gi.Function
var flagsGetFirstValueFunction_Once sync.Once

func flagsGetFirstValueFunction_Set() error {
	var err error
	flagsGetFirstValueFunction_Once.Do(func() {
		flagsGetFirstValueFunction, err = gi.FunctionInvokerNew("GObject", "flags_get_first_value")
	})
	return err
}

// FlagsGetFirstValue is a representation of the C type g_flags_get_first_value.
func FlagsGetFirstValue(flagsClass *FlagsClass, value uint32) *FlagsValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(flagsClass.native)
	inArgs[1].SetUint32(value)

	var ret gi.Argument

	err := flagsGetFirstValueFunction_Set()
	if err == nil {
		ret = flagsGetFirstValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FlagsValue{native: ret.Pointer()}

	return retGo
}

var flagsGetValueByNameFunction *gi.Function
var flagsGetValueByNameFunction_Once sync.Once

func flagsGetValueByNameFunction_Set() error {
	var err error
	flagsGetValueByNameFunction_Once.Do(func() {
		flagsGetValueByNameFunction, err = gi.FunctionInvokerNew("GObject", "flags_get_value_by_name")
	})
	return err
}

// FlagsGetValueByName is a representation of the C type g_flags_get_value_by_name.
func FlagsGetValueByName(flagsClass *FlagsClass, name string) *FlagsValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(flagsClass.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := flagsGetValueByNameFunction_Set()
	if err == nil {
		ret = flagsGetValueByNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FlagsValue{native: ret.Pointer()}

	return retGo
}

var flagsGetValueByNickFunction *gi.Function
var flagsGetValueByNickFunction_Once sync.Once

func flagsGetValueByNickFunction_Set() error {
	var err error
	flagsGetValueByNickFunction_Once.Do(func() {
		flagsGetValueByNickFunction, err = gi.FunctionInvokerNew("GObject", "flags_get_value_by_nick")
	})
	return err
}

// FlagsGetValueByNick is a representation of the C type g_flags_get_value_by_nick.
func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) *FlagsValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(flagsClass.native)
	inArgs[1].SetString(nick)

	var ret gi.Argument

	err := flagsGetValueByNickFunction_Set()
	if err == nil {
		ret = flagsGetValueByNickFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FlagsValue{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_flags_register_static' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_flags_to_string' : parameter 'flags_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_gtype_get_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_boolean' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_boxed' : parameter 'boxed_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_char' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_double' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_enum' : parameter 'enum_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_flags' : parameter 'flags_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_float' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_gtype' : parameter 'is_a_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_int' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_int64' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_long' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_object' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_override' : parameter 'overridden' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_param_spec_param' : parameter 'param_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_param_spec_pointer' : parameter 'flags' of type 'ParamFlags' not supported

var paramSpecPoolNewFunction *gi.Function
var paramSpecPoolNewFunction_Once sync.Once

func paramSpecPoolNewFunction_Set() error {
	var err error
	paramSpecPoolNewFunction_Once.Do(func() {
		paramSpecPoolNewFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_pool_new")
	})
	return err
}

// ParamSpecPoolNew is a representation of the C type g_param_spec_pool_new.
func ParamSpecPoolNew(typePrefixing bool) *ParamSpecPool {
	var inArgs [1]gi.Argument
	inArgs[0].SetBoolean(typePrefixing)

	var ret gi.Argument

	err := paramSpecPoolNewFunction_Set()
	if err == nil {
		ret = paramSpecPoolNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &ParamSpecPool{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_param_spec_string' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_uchar' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_uint' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_uint64' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_ulong' : parameter 'flags' of type 'ParamFlags' not supported

// UNSUPPORTED : C value 'g_param_spec_unichar' : parameter 'default_value' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_param_spec_value_array' : parameter 'element_spec' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_param_spec_variant' : parameter 'type' of type 'GLib.VariantType' not supported

// UNSUPPORTED : C value 'g_param_type_register_static' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_param_value_convert' : parameter 'pspec' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_param_value_defaults' : parameter 'pspec' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_param_value_set_default' : parameter 'pspec' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_param_value_validate' : parameter 'pspec' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_param_values_cmp' : parameter 'pspec' of type 'ParamSpec' not supported

// UNSUPPORTED : C value 'g_pointer_type_register_static' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_signal_accumulator_first_wins' : parameter 'dummy' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_signal_accumulator_true_handled' : parameter 'dummy' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_signal_add_emission_hook' : parameter 'hook_func' of type 'SignalEmissionHook' not supported

// UNSUPPORTED : C value 'g_signal_chain_from_overridden' : parameter 'instance_and_params' of type 'nil' not supported

// UNSUPPORTED : C value 'g_signal_chain_from_overridden_handler' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_signal_connect_closure' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_connect_closure_by_id' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_connect_data' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_connect_object' : parameter 'c_handler' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_signal_emit' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_emit_by_name' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_emit_valist' : parameter 'var_args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_signal_emitv' : parameter 'instance_and_params' of type 'nil' not supported

// UNSUPPORTED : C value 'g_signal_get_invocation_hint' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_handler_block' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_handler_disconnect' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_handler_find' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_handler_is_connected' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_handler_unblock' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_handlers_block_matched' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_handlers_destroy' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_handlers_disconnect_matched' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_handlers_unblock_matched' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_has_handler_pending' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_list_ids' : parameter 'itype' of type 'GType' not supported

// UNSUPPORTED : C value 'g_signal_lookup' : parameter 'itype' of type 'GType' not supported

var signalNameFunction *gi.Function
var signalNameFunction_Once sync.Once

func signalNameFunction_Set() error {
	var err error
	signalNameFunction_Once.Do(func() {
		signalNameFunction, err = gi.FunctionInvokerNew("GObject", "signal_name")
	})
	return err
}

// SignalName is a representation of the C type g_signal_name.
func SignalName(signalId uint32) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(signalId)

	var ret gi.Argument

	err := signalNameFunction_Set()
	if err == nil {
		ret = signalNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_signal_new' : parameter 'itype' of type 'GType' not supported

// UNSUPPORTED : C value 'g_signal_new_class_handler' : parameter 'itype' of type 'GType' not supported

// UNSUPPORTED : C value 'g_signal_new_valist' : parameter 'itype' of type 'GType' not supported

// UNSUPPORTED : C value 'g_signal_newv' : parameter 'itype' of type 'GType' not supported

// UNSUPPORTED : C value 'g_signal_override_class_closure' : parameter 'instance_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_signal_override_class_handler' : parameter 'instance_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_signal_parse_name' : parameter 'itype' of type 'GType' not supported

var signalQueryFunction *gi.Function
var signalQueryFunction_Once sync.Once

func signalQueryFunction_Set() error {
	var err error
	signalQueryFunction_Once.Do(func() {
		signalQueryFunction, err = gi.FunctionInvokerNew("GObject", "signal_query")
	})
	return err
}

// SignalQuery is a representation of the C type g_signal_query.
func SignalQuery(signalId uint32) *SignalQuery_ {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(signalId)

	var outArgs [1]gi.Argument

	err := signalQueryFunction_Set()
	if err == nil {
		signalQueryFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &SignalQuery_{native: outArgs[0].Pointer()}

	return out0
}

var signalRemoveEmissionHookFunction *gi.Function
var signalRemoveEmissionHookFunction_Once sync.Once

func signalRemoveEmissionHookFunction_Set() error {
	var err error
	signalRemoveEmissionHookFunction_Once.Do(func() {
		signalRemoveEmissionHookFunction, err = gi.FunctionInvokerNew("GObject", "signal_remove_emission_hook")
	})
	return err
}

// SignalRemoveEmissionHook is a representation of the C type g_signal_remove_emission_hook.
func SignalRemoveEmissionHook(signalId uint32, hookId uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(signalId)
	inArgs[1].SetUint64(hookId)

	err := signalRemoveEmissionHookFunction_Set()
	if err == nil {
		signalRemoveEmissionHookFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_signal_set_va_marshaller' : parameter 'instance_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_signal_stop_emission' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_stop_emission_by_name' : parameter 'instance' of type 'Object' not supported

// UNSUPPORTED : C value 'g_signal_type_cclosure_new' : parameter 'itype' of type 'GType' not supported

// UNSUPPORTED : C value 'g_source_set_closure' : parameter 'source' of type 'GLib.Source' not supported

// UNSUPPORTED : C value 'g_source_set_dummy_callback' : parameter 'source' of type 'GLib.Source' not supported

var strdupValueContentsFunction *gi.Function
var strdupValueContentsFunction_Once sync.Once

func strdupValueContentsFunction_Set() error {
	var err error
	strdupValueContentsFunction_Once.Do(func() {
		strdupValueContentsFunction, err = gi.FunctionInvokerNew("GObject", "strdup_value_contents")
	})
	return err
}

// StrdupValueContents is a representation of the C type g_strdup_value_contents.
func StrdupValueContents(value *Value) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(value.native)

	var ret gi.Argument

	err := strdupValueContentsFunction_Set()
	if err == nil {
		ret = strdupValueContentsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_type_add_class_cache_func' : parameter 'cache_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_type_add_class_private' : parameter 'class_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_add_instance_private' : parameter 'class_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_add_interface_check' : parameter 'check_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_type_add_interface_dynamic' : parameter 'instance_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_add_interface_static' : parameter 'instance_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_check_class_cast' : parameter 'is_a_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_check_class_is_a' : parameter 'is_a_type' of type 'GType' not supported

var typeCheckInstanceFunction *gi.Function
var typeCheckInstanceFunction_Once sync.Once

func typeCheckInstanceFunction_Set() error {
	var err error
	typeCheckInstanceFunction_Once.Do(func() {
		typeCheckInstanceFunction, err = gi.FunctionInvokerNew("GObject", "type_check_instance")
	})
	return err
}

// TypeCheckInstance is a representation of the C type g_type_check_instance.
func TypeCheckInstance(instance *TypeInstance) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(instance.native)

	var ret gi.Argument

	err := typeCheckInstanceFunction_Set()
	if err == nil {
		ret = typeCheckInstanceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_type_check_instance_cast' : parameter 'iface_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_check_instance_is_a' : parameter 'iface_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_check_instance_is_fundamentally_a' : parameter 'fundamental_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_check_is_value_type' : parameter 'type' of type 'GType' not supported

var typeCheckValueFunction *gi.Function
var typeCheckValueFunction_Once sync.Once

func typeCheckValueFunction_Set() error {
	var err error
	typeCheckValueFunction_Once.Do(func() {
		typeCheckValueFunction, err = gi.FunctionInvokerNew("GObject", "type_check_value")
	})
	return err
}

// TypeCheckValue is a representation of the C type g_type_check_value.
func TypeCheckValue(value *Value) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(value.native)

	var ret gi.Argument

	err := typeCheckValueFunction_Set()
	if err == nil {
		ret = typeCheckValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_type_check_value_holds' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_children' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_class_adjust_private_offset' : parameter 'g_class' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_type_class_peek' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_class_peek_static' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_class_ref' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_create_instance' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_default_interface_peek' : parameter 'g_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_default_interface_ref' : parameter 'g_type' of type 'GType' not supported

var typeDefaultInterfaceUnrefFunction *gi.Function
var typeDefaultInterfaceUnrefFunction_Once sync.Once

func typeDefaultInterfaceUnrefFunction_Set() error {
	var err error
	typeDefaultInterfaceUnrefFunction_Once.Do(func() {
		typeDefaultInterfaceUnrefFunction, err = gi.FunctionInvokerNew("GObject", "type_default_interface_unref")
	})
	return err
}

// TypeDefaultInterfaceUnref is a representation of the C type g_type_default_interface_unref.
func TypeDefaultInterfaceUnref(gIface *TypeInterface) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(gIface.native)

	err := typeDefaultInterfaceUnrefFunction_Set()
	if err == nil {
		typeDefaultInterfaceUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_type_depth' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_ensure' : parameter 'type' of type 'GType' not supported

var typeFreeInstanceFunction *gi.Function
var typeFreeInstanceFunction_Once sync.Once

func typeFreeInstanceFunction_Set() error {
	var err error
	typeFreeInstanceFunction_Once.Do(func() {
		typeFreeInstanceFunction, err = gi.FunctionInvokerNew("GObject", "type_free_instance")
	})
	return err
}

// TypeFreeInstance is a representation of the C type g_type_free_instance.
func TypeFreeInstance(instance *TypeInstance) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(instance.native)

	err := typeFreeInstanceFunction_Set()
	if err == nil {
		typeFreeInstanceFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_type_from_name' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_type_fundamental' : parameter 'type_id' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_fundamental_next' : return type 'GType' not supported

// UNSUPPORTED : C value 'g_type_get_instance_count' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_get_plugin' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_get_qdata' : parameter 'type' of type 'GType' not supported

var typeGetTypeRegistrationSerialFunction *gi.Function
var typeGetTypeRegistrationSerialFunction_Once sync.Once

func typeGetTypeRegistrationSerialFunction_Set() error {
	var err error
	typeGetTypeRegistrationSerialFunction_Once.Do(func() {
		typeGetTypeRegistrationSerialFunction, err = gi.FunctionInvokerNew("GObject", "type_get_type_registration_serial")
	})
	return err
}

// TypeGetTypeRegistrationSerial is a representation of the C type g_type_get_type_registration_serial.
func TypeGetTypeRegistrationSerial() uint32 {

	var ret gi.Argument

	err := typeGetTypeRegistrationSerialFunction_Set()
	if err == nil {
		ret = typeGetTypeRegistrationSerialFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var typeInitFunction *gi.Function
var typeInitFunction_Once sync.Once

func typeInitFunction_Set() error {
	var err error
	typeInitFunction_Once.Do(func() {
		typeInitFunction, err = gi.FunctionInvokerNew("GObject", "type_init")
	})
	return err
}

// TypeInit is a representation of the C type g_type_init.
func TypeInit() {

	err := typeInitFunction_Set()
	if err == nil {
		typeInitFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'g_type_init_with_debug_flags' : parameter 'debug_flags' of type 'TypeDebugFlags' not supported

// UNSUPPORTED : C value 'g_type_interface_add_prerequisite' : parameter 'interface_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_interface_get_plugin' : parameter 'instance_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_interface_peek' : parameter 'iface_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_interface_prerequisites' : parameter 'interface_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_interfaces' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_is_a' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_name' : parameter 'type' of type 'GType' not supported

var typeNameFromClassFunction *gi.Function
var typeNameFromClassFunction_Once sync.Once

func typeNameFromClassFunction_Set() error {
	var err error
	typeNameFromClassFunction_Once.Do(func() {
		typeNameFromClassFunction, err = gi.FunctionInvokerNew("GObject", "type_name_from_class")
	})
	return err
}

// TypeNameFromClass is a representation of the C type g_type_name_from_class.
func TypeNameFromClass(gClass *TypeClass) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(gClass.native)

	var ret gi.Argument

	err := typeNameFromClassFunction_Set()
	if err == nil {
		ret = typeNameFromClassFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var typeNameFromInstanceFunction *gi.Function
var typeNameFromInstanceFunction_Once sync.Once

func typeNameFromInstanceFunction_Set() error {
	var err error
	typeNameFromInstanceFunction_Once.Do(func() {
		typeNameFromInstanceFunction, err = gi.FunctionInvokerNew("GObject", "type_name_from_instance")
	})
	return err
}

// TypeNameFromInstance is a representation of the C type g_type_name_from_instance.
func TypeNameFromInstance(instance *TypeInstance) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(instance.native)

	var ret gi.Argument

	err := typeNameFromInstanceFunction_Set()
	if err == nil {
		ret = typeNameFromInstanceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_type_next_base' : parameter 'leaf_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_parent' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_qname' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_query' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_register_dynamic' : parameter 'parent_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_register_fundamental' : parameter 'type_id' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_register_static' : parameter 'parent_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_register_static_simple' : parameter 'parent_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_remove_class_cache_func' : parameter 'cache_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_type_remove_interface_check' : parameter 'check_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_type_set_qdata' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_test_flags' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_type_value_table_peek' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_value_register_transform_func' : parameter 'src_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_value_type_compatible' : parameter 'src_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_value_type_transformable' : parameter 'src_type' of type 'GType' not supported

var badFunction *gi.Function
var badFunction_Once sync.Once

func badFunction_Set() error {
	var err error
	badFunction_Once.Do(func() {
		badFunction, err = gi.FunctionInvokerNew("GObject", "bad")
	})
	return err
}

// Bad is a representation of the C type g_bad.
func Bad() {

	err := badFunction_Set()
	if err == nil {
		badFunction.Invoke(nil, nil)
	}

	return
}
