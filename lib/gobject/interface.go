// This is a generated file - DO NOT EDIT

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// The GObject type system supports dynamic loading of types.
// The #GTypePlugin interface is used to handle the lifecycle
// of dynamically loaded types. It goes as follows:
//
// 1. The type is initially introduced (usually upon loading the module
// the first time, or by your main application that knows what modules
// introduces what types), like this:
// |[<!-- language="C" -->
// new_type_id = g_type_register_dynamic (parent_type_id,
// "TypeName",
// new_type_plugin,
// type_flags);
// ]|
// where @new_type_plugin is an implementation of the
// #GTypePlugin interface.
//
// 2. The type's implementation is referenced, e.g. through
// g_type_class_ref() or through g_type_create_instance() (this is
// being called by g_object_new()) or through one of the above done on
// a type derived from @new_type_id.
//
// 3. This causes the type system to load the type's implementation by
// calling g_type_plugin_use() and g_type_plugin_complete_type_info()
// on @new_type_plugin.
//
// 4. At some point the type's implementation isn't required anymore,
// e.g. after g_type_class_unref() or g_type_free_instance() (called
// when the reference count of an instance drops to zero).
//
// 5. This causes the type system to throw away the information retrieved
// from g_type_plugin_complete_type_info() and then it calls
// g_type_plugin_unuse() on @new_type_plugin.
//
// 6. Things may repeat from the second step.
//
// So basically, you need to implement a #GTypePlugin type that
// carries a use_count, once use_count goes from zero to one, you need
// to load the implementation to successfully handle the upcoming
// g_type_plugin_complete_type_info() call. Later, maybe after
// succeeding use/unuse calls, once use_count drops to zero, you can
// unload the implementation again. The type system makes sure to call
// g_type_plugin_use() and g_type_plugin_complete_type_info() again
// when the type is needed again.
//
// #GTypeModule is an implementation of #GTypePlugin that already
// implements most of this except for the actual module loading and
// unloading. It even handles multiple registered types per module.
/*

C type

GTypePlugin
*/
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

// Calls the @complete_interface_info function from the
// #GTypePluginClass of @plugin. There should be no need to use this
// function outside of the GObject type system itself.
/*

C function

g_type_plugin_complete_interface_info
*/
func (recv *TypePlugin) CompleteInterfaceInfo(instanceType Type, interfaceType Type, info *InterfaceInfo) {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	c_info := (*C.GInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GInterfaceInfo)(info.ToC())
	}

	C.g_type_plugin_complete_interface_info((*C.GTypePlugin)(recv.native), c_instance_type, c_interface_type, c_info)

	return
}

// Calls the @complete_type_info function from the #GTypePluginClass of @plugin.
// There should be no need to use this function outside of the GObject
// type system itself.
/*

C function

g_type_plugin_complete_type_info
*/
func (recv *TypePlugin) CompleteTypeInfo(gType Type, info *TypeInfo, valueTable *TypeValueTable) {
	c_g_type := (C.GType)(gType)

	c_info := (*C.GTypeInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GTypeInfo)(info.ToC())
	}

	c_value_table := (*C.GTypeValueTable)(C.NULL)
	if valueTable != nil {
		c_value_table = (*C.GTypeValueTable)(valueTable.ToC())
	}

	C.g_type_plugin_complete_type_info((*C.GTypePlugin)(recv.native), c_g_type, c_info, c_value_table)

	return
}

// Calls the @unuse_plugin function from the #GTypePluginClass of
// @plugin.  There should be no need to use this function outside of
// the GObject type system itself.
/*

C function

g_type_plugin_unuse
*/
func (recv *TypePlugin) Unuse() {
	C.g_type_plugin_unuse((*C.GTypePlugin)(recv.native))

	return
}

// Calls the @use_plugin function from the #GTypePluginClass of
// @plugin.  There should be no need to use this function outside of
// the GObject type system itself.
/*

C function

g_type_plugin_use
*/
func (recv *TypePlugin) Use() {
	C.g_type_plugin_use((*C.GTypePlugin)(recv.native))

	return
}
