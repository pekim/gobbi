// This is a generated file - DO NOT EDIT

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Unsupported : atk_add_focus_tracker : unsupported parameter focus_tracker : no type generator for EventListener (AtkEventListener) for param focus_tracker

// Unsupported : atk_add_global_event_listener : unsupported parameter listener : no type generator for GObject.SignalEmissionHook (GSignalEmissionHook) for param listener

// Unsupported : atk_add_key_event_listener : unsupported parameter listener : no type generator for KeySnoopFunc (AtkKeySnoopFunc) for param listener

// Blacklisted : atk_attribute_set_free

// Unsupported : atk_focus_tracker_init : unsupported parameter init : no type generator for EventListenerInit (AtkEventListenerInit) for param init

// FocusTrackerNotify is a wrapper around the C function atk_focus_tracker_notify.
func FocusTrackerNotify(object *Object) {
	c_object := (*C.AtkObject)(object.ToC())

	C.atk_focus_tracker_notify(c_object)

	return
}

// GetDefaultRegistry is a wrapper around the C function atk_get_default_registry.
func GetDefaultRegistry() *Registry {
	retC := C.atk_get_default_registry()
	retGo := RegistryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRoot is a wrapper around the C function atk_get_root.
func GetRoot() *Object {
	retC := C.atk_get_root()
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetToolkitName is a wrapper around the C function atk_get_toolkit_name.
func GetToolkitName() string {
	retC := C.atk_get_toolkit_name()
	retGo := C.GoString(retC)

	return retGo
}

// GetToolkitVersion is a wrapper around the C function atk_get_toolkit_version.
func GetToolkitVersion() string {
	retC := C.atk_get_toolkit_version()
	retGo := C.GoString(retC)

	return retGo
}

// RelationTypeForName is a wrapper around the C function atk_relation_type_for_name.
func RelationTypeForName(name string) RelationType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_relation_type_for_name(c_name)
	retGo := (RelationType)(retC)

	return retGo
}

// RelationTypeGetName is a wrapper around the C function atk_relation_type_get_name.
func RelationTypeGetName(type_ RelationType) string {
	c_type := (C.AtkRelationType)(type_)

	retC := C.atk_relation_type_get_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// RelationTypeRegister is a wrapper around the C function atk_relation_type_register.
func RelationTypeRegister(name string) RelationType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_relation_type_register(c_name)
	retGo := (RelationType)(retC)

	return retGo
}

// RemoveFocusTracker is a wrapper around the C function atk_remove_focus_tracker.
func RemoveFocusTracker(trackerId uint32) {
	c_tracker_id := (C.guint)(trackerId)

	C.atk_remove_focus_tracker(c_tracker_id)

	return
}

// RemoveGlobalEventListener is a wrapper around the C function atk_remove_global_event_listener.
func RemoveGlobalEventListener(listenerId uint32) {
	c_listener_id := (C.guint)(listenerId)

	C.atk_remove_global_event_listener(c_listener_id)

	return
}

// RemoveKeyEventListener is a wrapper around the C function atk_remove_key_event_listener.
func RemoveKeyEventListener(listenerId uint32) {
	c_listener_id := (C.guint)(listenerId)

	C.atk_remove_key_event_listener(c_listener_id)

	return
}

// RoleForName is a wrapper around the C function atk_role_for_name.
func RoleForName(name string) Role {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_role_for_name(c_name)
	retGo := (Role)(retC)

	return retGo
}

// RoleGetLocalizedName is a wrapper around the C function atk_role_get_localized_name.
func RoleGetLocalizedName(role Role) string {
	c_role := (C.AtkRole)(role)

	retC := C.atk_role_get_localized_name(c_role)
	retGo := C.GoString(retC)

	return retGo
}

// RoleGetName is a wrapper around the C function atk_role_get_name.
func RoleGetName(role Role) string {
	c_role := (C.AtkRole)(role)

	retC := C.atk_role_get_name(c_role)
	retGo := C.GoString(retC)

	return retGo
}

// RoleRegister is a wrapper around the C function atk_role_register.
func RoleRegister(name string) Role {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_role_register(c_name)
	retGo := (Role)(retC)

	return retGo
}

// StateTypeForName is a wrapper around the C function atk_state_type_for_name.
func StateTypeForName(name string) StateType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_state_type_for_name(c_name)
	retGo := (StateType)(retC)

	return retGo
}

// StateTypeGetName is a wrapper around the C function atk_state_type_get_name.
func StateTypeGetName(type_ StateType) string {
	c_type := (C.AtkStateType)(type_)

	retC := C.atk_state_type_get_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// StateTypeRegister is a wrapper around the C function atk_state_type_register.
func StateTypeRegister(name string) StateType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_state_type_register(c_name)
	retGo := (StateType)(retC)

	return retGo
}

// TextAttributeForName is a wrapper around the C function atk_text_attribute_for_name.
func TextAttributeForName(name string) TextAttribute {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_text_attribute_for_name(c_name)
	retGo := (TextAttribute)(retC)

	return retGo
}

// TextAttributeGetName is a wrapper around the C function atk_text_attribute_get_name.
func TextAttributeGetName(attr TextAttribute) string {
	c_attr := (C.AtkTextAttribute)(attr)

	retC := C.atk_text_attribute_get_name(c_attr)
	retGo := C.GoString(retC)

	return retGo
}

// TextAttributeGetValue is a wrapper around the C function atk_text_attribute_get_value.
func TextAttributeGetValue(attr TextAttribute, index int32) string {
	c_attr := (C.AtkTextAttribute)(attr)

	c_index_ := (C.gint)(index)

	retC := C.atk_text_attribute_get_value(c_attr, c_index_)
	retGo := C.GoString(retC)

	return retGo
}

// TextAttributeRegister is a wrapper around the C function atk_text_attribute_register.
func TextAttributeRegister(name string) TextAttribute {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_text_attribute_register(c_name)
	retGo := (TextAttribute)(retC)

	return retGo
}

// Unsupported : atk_text_free_ranges : unsupported parameter ranges :

// ValueTypeGetLocalizedName is a wrapper around the C function atk_value_type_get_localized_name.
func ValueTypeGetLocalizedName(valueType ValueType) string {
	c_value_type := (C.AtkValueType)(valueType)

	retC := C.atk_value_type_get_localized_name(c_value_type)
	retGo := C.GoString(retC)

	return retGo
}

// ValueTypeGetName is a wrapper around the C function atk_value_type_get_name.
func ValueTypeGetName(valueType ValueType) string {
	c_value_type := (C.AtkValueType)(valueType)

	retC := C.atk_value_type_get_name(c_value_type)
	retGo := C.GoString(retC)

	return retGo
}
