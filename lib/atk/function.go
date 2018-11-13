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

// Cause the focus tracker functions which have been specified to be
// executed for the object.
/*

C function : atk_focus_tracker_notify
*/
func FocusTrackerNotify(object *Object) {
	c_object := (*C.AtkObject)(C.NULL)
	if object != nil {
		c_object = (*C.AtkObject)(object.ToC())
	}

	C.atk_focus_tracker_notify(c_object)

	return
}

// Gets a default implementation of the #AtkObjectFactory/type
// registry.
// Note: For most toolkit maintainers, this will be the correct
// registry for registering new #AtkObject factories. Following
// a call to this function, maintainers may call atk_registry_set_factory_type()
// to associate an #AtkObjectFactory subclass with the GType of objects
// for whom accessibility information will be provided.
/*

C function : atk_get_default_registry
*/
func GetDefaultRegistry() *Registry {
	retC := C.atk_get_default_registry()
	retGo := RegistryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the root accessible container for the current application.
/*

C function : atk_get_root
*/
func GetRoot() *Object {
	retC := C.atk_get_root()
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets name string for the GUI toolkit implementing ATK for this application.
/*

C function : atk_get_toolkit_name
*/
func GetToolkitName() string {
	retC := C.atk_get_toolkit_name()
	retGo := C.GoString(retC)

	return retGo
}

// Gets version string for the GUI toolkit implementing ATK for this application.
/*

C function : atk_get_toolkit_version
*/
func GetToolkitVersion() string {
	retC := C.atk_get_toolkit_version()
	retGo := C.GoString(retC)

	return retGo
}

// Get the #AtkRelationType type corresponding to a relation name.
/*

C function : atk_relation_type_for_name
*/
func RelationTypeForName(name string) RelationType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_relation_type_for_name(c_name)
	retGo := (RelationType)(retC)

	return retGo
}

// Gets the description string describing the #AtkRelationType @type.
/*

C function : atk_relation_type_get_name
*/
func RelationTypeGetName(type_ RelationType) string {
	c_type := (C.AtkRelationType)(type_)

	retC := C.atk_relation_type_get_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// Associate @name with a new #AtkRelationType
/*

C function : atk_relation_type_register
*/
func RelationTypeRegister(name string) RelationType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_relation_type_register(c_name)
	retGo := (RelationType)(retC)

	return retGo
}

/*

C function : atk_remove_focus_tracker
*/
func RemoveFocusTracker(trackerId uint32) {
	c_tracker_id := (C.guint)(trackerId)

	C.atk_remove_focus_tracker(c_tracker_id)

	return
}

// @listener_id is the value returned by #atk_add_global_event_listener
// when you registered that event listener.
//
// Toolkit implementor note: ATK provides a default implementation for
// this virtual method. ATK implementors are discouraged from
// reimplementing this method.
//
// Toolkit implementor note: this method is not intended to be used by
// ATK implementors but by ATK consumers.
//
// Removes the specified event listener
/*

C function : atk_remove_global_event_listener
*/
func RemoveGlobalEventListener(listenerId uint32) {
	c_listener_id := (C.guint)(listenerId)

	C.atk_remove_global_event_listener(c_listener_id)

	return
}

// @listener_id is the value returned by #atk_add_key_event_listener
// when you registered that event listener.
//
// Removes the specified event listener.
/*

C function : atk_remove_key_event_listener
*/
func RemoveKeyEventListener(listenerId uint32) {
	c_listener_id := (C.guint)(listenerId)

	C.atk_remove_key_event_listener(c_listener_id)

	return
}

// Get the #AtkRole type corresponding to a rolew name.
/*

C function : atk_role_for_name
*/
func RoleForName(name string) Role {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_role_for_name(c_name)
	retGo := (Role)(retC)

	return retGo
}

// Gets the localized description string describing the #AtkRole @role.
/*

C function : atk_role_get_localized_name
*/
func RoleGetLocalizedName(role Role) string {
	c_role := (C.AtkRole)(role)

	retC := C.atk_role_get_localized_name(c_role)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the description string describing the #AtkRole @role.
/*

C function : atk_role_get_name
*/
func RoleGetName(role Role) string {
	c_role := (C.AtkRole)(role)

	retC := C.atk_role_get_name(c_role)
	retGo := C.GoString(retC)

	return retGo
}

// Registers the role specified by @name. @name must be a meaningful
// name. So it should not be empty, or consisting on whitespaces.
/*

C function : atk_role_register
*/
func RoleRegister(name string) Role {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_role_register(c_name)
	retGo := (Role)(retC)

	return retGo
}

// Gets the #AtkStateType corresponding to the description string @name.
/*

C function : atk_state_type_for_name
*/
func StateTypeForName(name string) StateType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_state_type_for_name(c_name)
	retGo := (StateType)(retC)

	return retGo
}

// Gets the description string describing the #AtkStateType @type.
/*

C function : atk_state_type_get_name
*/
func StateTypeGetName(type_ StateType) string {
	c_type := (C.AtkStateType)(type_)

	retC := C.atk_state_type_get_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// Register a new object state.
/*

C function : atk_state_type_register
*/
func StateTypeRegister(name string) StateType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_state_type_register(c_name)
	retGo := (StateType)(retC)

	return retGo
}

// Get the #AtkTextAttribute type corresponding to a text attribute name.
/*

C function : atk_text_attribute_for_name
*/
func TextAttributeForName(name string) TextAttribute {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_text_attribute_for_name(c_name)
	retGo := (TextAttribute)(retC)

	return retGo
}

// Gets the name corresponding to the #AtkTextAttribute
/*

C function : atk_text_attribute_get_name
*/
func TextAttributeGetName(attr TextAttribute) string {
	c_attr := (C.AtkTextAttribute)(attr)

	retC := C.atk_text_attribute_get_name(c_attr)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the value for the index of the #AtkTextAttribute
/*

C function : atk_text_attribute_get_value
*/
func TextAttributeGetValue(attr TextAttribute, index int32) string {
	c_attr := (C.AtkTextAttribute)(attr)

	c_index_ := (C.gint)(index)

	retC := C.atk_text_attribute_get_value(c_attr, c_index_)
	retGo := C.GoString(retC)

	return retGo
}

// Associate @name with a new #AtkTextAttribute
/*

C function : atk_text_attribute_register
*/
func TextAttributeRegister(name string) TextAttribute {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_text_attribute_register(c_name)
	retGo := (TextAttribute)(retC)

	return retGo
}

// Gets the localized description string describing the #AtkValueType @value_type.
/*

C function : atk_value_type_get_localized_name
*/
func ValueTypeGetLocalizedName(valueType ValueType) string {
	c_value_type := (C.AtkValueType)(valueType)

	retC := C.atk_value_type_get_localized_name(c_value_type)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the description string describing the #AtkValueType @value_type.
/*

C function : atk_value_type_get_name
*/
func ValueTypeGetName(valueType ValueType) string {
	c_value_type := (C.AtkValueType)(valueType)

	retC := C.atk_value_type_get_name(c_value_type)
	retGo := C.GoString(retC)

	return retGo
}
