// Code generated - DO NOT EDIT.

package atk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'atk_add_focus_tracker' : parameter 'focus_tracker' of type 'EventListener' not supported

// UNSUPPORTED : C value 'atk_add_global_event_listener' : parameter 'listener' of type 'GObject.SignalEmissionHook' not supported

// UNSUPPORTED : C value 'atk_add_key_event_listener' : parameter 'listener' of type 'KeySnoopFunc' not supported

// UNSUPPORTED : C value 'atk_attribute_set_free' : parameter 'attrib_set' of type 'AttributeSet' not supported

// UNSUPPORTED : C value 'atk_focus_tracker_init' : parameter 'init' of type 'EventListenerInit' not supported

var focusTrackerNotifyFunction *gi.Function
var focusTrackerNotifyFunction_Once sync.Once

func focusTrackerNotifyFunction_Set() error {
	var err error
	focusTrackerNotifyFunction_Once.Do(func() {
		focusTrackerNotifyFunction, err = gi.FunctionInvokerNew("Atk", "focus_tracker_notify")
	})
	return err
}

// FocusTrackerNotify is a representation of the C type atk_focus_tracker_notify.
func FocusTrackerNotify(object *Object) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(object.native)

	err := focusTrackerNotifyFunction_Set()
	if err == nil {
		focusTrackerNotifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var getBinaryAgeFunction *gi.Function
var getBinaryAgeFunction_Once sync.Once

func getBinaryAgeFunction_Set() error {
	var err error
	getBinaryAgeFunction_Once.Do(func() {
		getBinaryAgeFunction, err = gi.FunctionInvokerNew("Atk", "get_binary_age")
	})
	return err
}

// GetBinaryAge is a representation of the C type atk_get_binary_age.
func GetBinaryAge() uint32 {

	var ret gi.Argument

	err := getBinaryAgeFunction_Set()
	if err == nil {
		ret = getBinaryAgeFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getDefaultRegistryFunction *gi.Function
var getDefaultRegistryFunction_Once sync.Once

func getDefaultRegistryFunction_Set() error {
	var err error
	getDefaultRegistryFunction_Once.Do(func() {
		getDefaultRegistryFunction, err = gi.FunctionInvokerNew("Atk", "get_default_registry")
	})
	return err
}

// GetDefaultRegistry is a representation of the C type atk_get_default_registry.
func GetDefaultRegistry() *Registry {

	var ret gi.Argument

	err := getDefaultRegistryFunction_Set()
	if err == nil {
		ret = getDefaultRegistryFunction.Invoke(nil, nil)
	}

	retGo := RegistryNewFromNative(ret.Pointer())

	return retGo
}

var getFocusObjectFunction *gi.Function
var getFocusObjectFunction_Once sync.Once

func getFocusObjectFunction_Set() error {
	var err error
	getFocusObjectFunction_Once.Do(func() {
		getFocusObjectFunction, err = gi.FunctionInvokerNew("Atk", "get_focus_object")
	})
	return err
}

// GetFocusObject is a representation of the C type atk_get_focus_object.
func GetFocusObject() *Object {

	var ret gi.Argument

	err := getFocusObjectFunction_Set()
	if err == nil {
		ret = getFocusObjectFunction.Invoke(nil, nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var getInterfaceAgeFunction *gi.Function
var getInterfaceAgeFunction_Once sync.Once

func getInterfaceAgeFunction_Set() error {
	var err error
	getInterfaceAgeFunction_Once.Do(func() {
		getInterfaceAgeFunction, err = gi.FunctionInvokerNew("Atk", "get_interface_age")
	})
	return err
}

// GetInterfaceAge is a representation of the C type atk_get_interface_age.
func GetInterfaceAge() uint32 {

	var ret gi.Argument

	err := getInterfaceAgeFunction_Set()
	if err == nil {
		ret = getInterfaceAgeFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() error {
	var err error
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction, err = gi.FunctionInvokerNew("Atk", "get_major_version")
	})
	return err
}

// GetMajorVersion is a representation of the C type atk_get_major_version.
func GetMajorVersion() uint32 {

	var ret gi.Argument

	err := getMajorVersionFunction_Set()
	if err == nil {
		ret = getMajorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMicroVersionFunction *gi.Function
var getMicroVersionFunction_Once sync.Once

func getMicroVersionFunction_Set() error {
	var err error
	getMicroVersionFunction_Once.Do(func() {
		getMicroVersionFunction, err = gi.FunctionInvokerNew("Atk", "get_micro_version")
	})
	return err
}

// GetMicroVersion is a representation of the C type atk_get_micro_version.
func GetMicroVersion() uint32 {

	var ret gi.Argument

	err := getMicroVersionFunction_Set()
	if err == nil {
		ret = getMicroVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMinorVersionFunction *gi.Function
var getMinorVersionFunction_Once sync.Once

func getMinorVersionFunction_Set() error {
	var err error
	getMinorVersionFunction_Once.Do(func() {
		getMinorVersionFunction, err = gi.FunctionInvokerNew("Atk", "get_minor_version")
	})
	return err
}

// GetMinorVersion is a representation of the C type atk_get_minor_version.
func GetMinorVersion() uint32 {

	var ret gi.Argument

	err := getMinorVersionFunction_Set()
	if err == nil {
		ret = getMinorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getRootFunction *gi.Function
var getRootFunction_Once sync.Once

func getRootFunction_Set() error {
	var err error
	getRootFunction_Once.Do(func() {
		getRootFunction, err = gi.FunctionInvokerNew("Atk", "get_root")
	})
	return err
}

// GetRoot is a representation of the C type atk_get_root.
func GetRoot() *Object {

	var ret gi.Argument

	err := getRootFunction_Set()
	if err == nil {
		ret = getRootFunction.Invoke(nil, nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var getToolkitNameFunction *gi.Function
var getToolkitNameFunction_Once sync.Once

func getToolkitNameFunction_Set() error {
	var err error
	getToolkitNameFunction_Once.Do(func() {
		getToolkitNameFunction, err = gi.FunctionInvokerNew("Atk", "get_toolkit_name")
	})
	return err
}

// GetToolkitName is a representation of the C type atk_get_toolkit_name.
func GetToolkitName() string {

	var ret gi.Argument

	err := getToolkitNameFunction_Set()
	if err == nil {
		ret = getToolkitNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getToolkitVersionFunction *gi.Function
var getToolkitVersionFunction_Once sync.Once

func getToolkitVersionFunction_Set() error {
	var err error
	getToolkitVersionFunction_Once.Do(func() {
		getToolkitVersionFunction, err = gi.FunctionInvokerNew("Atk", "get_toolkit_version")
	})
	return err
}

// GetToolkitVersion is a representation of the C type atk_get_toolkit_version.
func GetToolkitVersion() string {

	var ret gi.Argument

	err := getToolkitVersionFunction_Set()
	if err == nil {
		ret = getToolkitVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getVersionFunction *gi.Function
var getVersionFunction_Once sync.Once

func getVersionFunction_Set() error {
	var err error
	getVersionFunction_Once.Do(func() {
		getVersionFunction, err = gi.FunctionInvokerNew("Atk", "get_version")
	})
	return err
}

// GetVersion is a representation of the C type atk_get_version.
func GetVersion() string {

	var ret gi.Argument

	err := getVersionFunction_Set()
	if err == nil {
		ret = getVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var relationTypeForNameFunction *gi.Function
var relationTypeForNameFunction_Once sync.Once

func relationTypeForNameFunction_Set() error {
	var err error
	relationTypeForNameFunction_Once.Do(func() {
		relationTypeForNameFunction, err = gi.FunctionInvokerNew("Atk", "relation_type_for_name")
	})
	return err
}

// RelationTypeForName is a representation of the C type atk_relation_type_for_name.
func RelationTypeForName(name string) RelationType {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := relationTypeForNameFunction_Set()
	if err == nil {
		ret = relationTypeForNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := RelationType(ret.Int32())

	return retGo
}

var relationTypeGetNameFunction *gi.Function
var relationTypeGetNameFunction_Once sync.Once

func relationTypeGetNameFunction_Set() error {
	var err error
	relationTypeGetNameFunction_Once.Do(func() {
		relationTypeGetNameFunction, err = gi.FunctionInvokerNew("Atk", "relation_type_get_name")
	})
	return err
}

// RelationTypeGetName is a representation of the C type atk_relation_type_get_name.
func RelationTypeGetName(type_ RelationType) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(type_))

	var ret gi.Argument

	err := relationTypeGetNameFunction_Set()
	if err == nil {
		ret = relationTypeGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var relationTypeRegisterFunction *gi.Function
var relationTypeRegisterFunction_Once sync.Once

func relationTypeRegisterFunction_Set() error {
	var err error
	relationTypeRegisterFunction_Once.Do(func() {
		relationTypeRegisterFunction, err = gi.FunctionInvokerNew("Atk", "relation_type_register")
	})
	return err
}

// RelationTypeRegister is a representation of the C type atk_relation_type_register.
func RelationTypeRegister(name string) RelationType {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := relationTypeRegisterFunction_Set()
	if err == nil {
		ret = relationTypeRegisterFunction.Invoke(inArgs[:], nil)
	}

	retGo := RelationType(ret.Int32())

	return retGo
}

var removeFocusTrackerFunction *gi.Function
var removeFocusTrackerFunction_Once sync.Once

func removeFocusTrackerFunction_Set() error {
	var err error
	removeFocusTrackerFunction_Once.Do(func() {
		removeFocusTrackerFunction, err = gi.FunctionInvokerNew("Atk", "remove_focus_tracker")
	})
	return err
}

// RemoveFocusTracker is a representation of the C type atk_remove_focus_tracker.
func RemoveFocusTracker(trackerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(trackerId)

	err := removeFocusTrackerFunction_Set()
	if err == nil {
		removeFocusTrackerFunction.Invoke(inArgs[:], nil)
	}

	return
}

var removeGlobalEventListenerFunction *gi.Function
var removeGlobalEventListenerFunction_Once sync.Once

func removeGlobalEventListenerFunction_Set() error {
	var err error
	removeGlobalEventListenerFunction_Once.Do(func() {
		removeGlobalEventListenerFunction, err = gi.FunctionInvokerNew("Atk", "remove_global_event_listener")
	})
	return err
}

// RemoveGlobalEventListener is a representation of the C type atk_remove_global_event_listener.
func RemoveGlobalEventListener(listenerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(listenerId)

	err := removeGlobalEventListenerFunction_Set()
	if err == nil {
		removeGlobalEventListenerFunction.Invoke(inArgs[:], nil)
	}

	return
}

var removeKeyEventListenerFunction *gi.Function
var removeKeyEventListenerFunction_Once sync.Once

func removeKeyEventListenerFunction_Set() error {
	var err error
	removeKeyEventListenerFunction_Once.Do(func() {
		removeKeyEventListenerFunction, err = gi.FunctionInvokerNew("Atk", "remove_key_event_listener")
	})
	return err
}

// RemoveKeyEventListener is a representation of the C type atk_remove_key_event_listener.
func RemoveKeyEventListener(listenerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(listenerId)

	err := removeKeyEventListenerFunction_Set()
	if err == nil {
		removeKeyEventListenerFunction.Invoke(inArgs[:], nil)
	}

	return
}

var roleForNameFunction *gi.Function
var roleForNameFunction_Once sync.Once

func roleForNameFunction_Set() error {
	var err error
	roleForNameFunction_Once.Do(func() {
		roleForNameFunction, err = gi.FunctionInvokerNew("Atk", "role_for_name")
	})
	return err
}

// RoleForName is a representation of the C type atk_role_for_name.
func RoleForName(name string) Role {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := roleForNameFunction_Set()
	if err == nil {
		ret = roleForNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := Role(ret.Int32())

	return retGo
}

var roleGetLocalizedNameFunction *gi.Function
var roleGetLocalizedNameFunction_Once sync.Once

func roleGetLocalizedNameFunction_Set() error {
	var err error
	roleGetLocalizedNameFunction_Once.Do(func() {
		roleGetLocalizedNameFunction, err = gi.FunctionInvokerNew("Atk", "role_get_localized_name")
	})
	return err
}

// RoleGetLocalizedName is a representation of the C type atk_role_get_localized_name.
func RoleGetLocalizedName(role Role) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(role))

	var ret gi.Argument

	err := roleGetLocalizedNameFunction_Set()
	if err == nil {
		ret = roleGetLocalizedNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var roleGetNameFunction *gi.Function
var roleGetNameFunction_Once sync.Once

func roleGetNameFunction_Set() error {
	var err error
	roleGetNameFunction_Once.Do(func() {
		roleGetNameFunction, err = gi.FunctionInvokerNew("Atk", "role_get_name")
	})
	return err
}

// RoleGetName is a representation of the C type atk_role_get_name.
func RoleGetName(role Role) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(role))

	var ret gi.Argument

	err := roleGetNameFunction_Set()
	if err == nil {
		ret = roleGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var roleRegisterFunction *gi.Function
var roleRegisterFunction_Once sync.Once

func roleRegisterFunction_Set() error {
	var err error
	roleRegisterFunction_Once.Do(func() {
		roleRegisterFunction, err = gi.FunctionInvokerNew("Atk", "role_register")
	})
	return err
}

// RoleRegister is a representation of the C type atk_role_register.
func RoleRegister(name string) Role {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := roleRegisterFunction_Set()
	if err == nil {
		ret = roleRegisterFunction.Invoke(inArgs[:], nil)
	}

	retGo := Role(ret.Int32())

	return retGo
}

var stateTypeForNameFunction *gi.Function
var stateTypeForNameFunction_Once sync.Once

func stateTypeForNameFunction_Set() error {
	var err error
	stateTypeForNameFunction_Once.Do(func() {
		stateTypeForNameFunction, err = gi.FunctionInvokerNew("Atk", "state_type_for_name")
	})
	return err
}

// StateTypeForName is a representation of the C type atk_state_type_for_name.
func StateTypeForName(name string) StateType {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := stateTypeForNameFunction_Set()
	if err == nil {
		ret = stateTypeForNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := StateType(ret.Int32())

	return retGo
}

var stateTypeGetNameFunction *gi.Function
var stateTypeGetNameFunction_Once sync.Once

func stateTypeGetNameFunction_Set() error {
	var err error
	stateTypeGetNameFunction_Once.Do(func() {
		stateTypeGetNameFunction, err = gi.FunctionInvokerNew("Atk", "state_type_get_name")
	})
	return err
}

// StateTypeGetName is a representation of the C type atk_state_type_get_name.
func StateTypeGetName(type_ StateType) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(type_))

	var ret gi.Argument

	err := stateTypeGetNameFunction_Set()
	if err == nil {
		ret = stateTypeGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var stateTypeRegisterFunction *gi.Function
var stateTypeRegisterFunction_Once sync.Once

func stateTypeRegisterFunction_Set() error {
	var err error
	stateTypeRegisterFunction_Once.Do(func() {
		stateTypeRegisterFunction, err = gi.FunctionInvokerNew("Atk", "state_type_register")
	})
	return err
}

// StateTypeRegister is a representation of the C type atk_state_type_register.
func StateTypeRegister(name string) StateType {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := stateTypeRegisterFunction_Set()
	if err == nil {
		ret = stateTypeRegisterFunction.Invoke(inArgs[:], nil)
	}

	retGo := StateType(ret.Int32())

	return retGo
}

var textAttributeForNameFunction *gi.Function
var textAttributeForNameFunction_Once sync.Once

func textAttributeForNameFunction_Set() error {
	var err error
	textAttributeForNameFunction_Once.Do(func() {
		textAttributeForNameFunction, err = gi.FunctionInvokerNew("Atk", "text_attribute_for_name")
	})
	return err
}

// TextAttributeForName is a representation of the C type atk_text_attribute_for_name.
func TextAttributeForName(name string) TextAttribute {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := textAttributeForNameFunction_Set()
	if err == nil {
		ret = textAttributeForNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := TextAttribute(ret.Int32())

	return retGo
}

var textAttributeGetNameFunction *gi.Function
var textAttributeGetNameFunction_Once sync.Once

func textAttributeGetNameFunction_Set() error {
	var err error
	textAttributeGetNameFunction_Once.Do(func() {
		textAttributeGetNameFunction, err = gi.FunctionInvokerNew("Atk", "text_attribute_get_name")
	})
	return err
}

// TextAttributeGetName is a representation of the C type atk_text_attribute_get_name.
func TextAttributeGetName(attr TextAttribute) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(attr))

	var ret gi.Argument

	err := textAttributeGetNameFunction_Set()
	if err == nil {
		ret = textAttributeGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var textAttributeGetValueFunction *gi.Function
var textAttributeGetValueFunction_Once sync.Once

func textAttributeGetValueFunction_Set() error {
	var err error
	textAttributeGetValueFunction_Once.Do(func() {
		textAttributeGetValueFunction, err = gi.FunctionInvokerNew("Atk", "text_attribute_get_value")
	})
	return err
}

// TextAttributeGetValue is a representation of the C type atk_text_attribute_get_value.
func TextAttributeGetValue(attr TextAttribute, index int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(int32(attr))
	inArgs[1].SetInt32(index)

	var ret gi.Argument

	err := textAttributeGetValueFunction_Set()
	if err == nil {
		ret = textAttributeGetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var textAttributeRegisterFunction *gi.Function
var textAttributeRegisterFunction_Once sync.Once

func textAttributeRegisterFunction_Set() error {
	var err error
	textAttributeRegisterFunction_Once.Do(func() {
		textAttributeRegisterFunction, err = gi.FunctionInvokerNew("Atk", "text_attribute_register")
	})
	return err
}

// TextAttributeRegister is a representation of the C type atk_text_attribute_register.
func TextAttributeRegister(name string) TextAttribute {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := textAttributeRegisterFunction_Set()
	if err == nil {
		ret = textAttributeRegisterFunction.Invoke(inArgs[:], nil)
	}

	retGo := TextAttribute(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'atk_text_free_ranges' : parameter 'ranges' of type 'nil' not supported

var valueTypeGetLocalizedNameFunction *gi.Function
var valueTypeGetLocalizedNameFunction_Once sync.Once

func valueTypeGetLocalizedNameFunction_Set() error {
	var err error
	valueTypeGetLocalizedNameFunction_Once.Do(func() {
		valueTypeGetLocalizedNameFunction, err = gi.FunctionInvokerNew("Atk", "value_type_get_localized_name")
	})
	return err
}

// ValueTypeGetLocalizedName is a representation of the C type atk_value_type_get_localized_name.
func ValueTypeGetLocalizedName(valueType ValueType) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(valueType))

	var ret gi.Argument

	err := valueTypeGetLocalizedNameFunction_Set()
	if err == nil {
		ret = valueTypeGetLocalizedNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var valueTypeGetNameFunction *gi.Function
var valueTypeGetNameFunction_Once sync.Once

func valueTypeGetNameFunction_Set() error {
	var err error
	valueTypeGetNameFunction_Once.Do(func() {
		valueTypeGetNameFunction, err = gi.FunctionInvokerNew("Atk", "value_type_get_name")
	})
	return err
}

// ValueTypeGetName is a representation of the C type atk_value_type_get_name.
func ValueTypeGetName(valueType ValueType) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(valueType))

	var ret gi.Argument

	err := valueTypeGetNameFunction_Set()
	if err == nil {
		ret = valueTypeGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}
