// Code generated - DO NOT EDIT.
// +build !atk_2.7.4,!atk_2.12,!atk_2.14,!atk_2.30

package atk

import (
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
	"unsafe"
)

// #include <atk/atk.h>
import "C"

type ActionIface C.AtkActionIface
type Attribute C.AtkAttribute
type ComponentIface C.AtkComponentIface
type DocumentIface C.AtkDocumentIface
type EditableTextIface C.AtkEditableTextIface
type GObjectAccessibleClass C.AtkGObjectAccessibleClass
type HyperlinkClass C.AtkHyperlinkClass
type HyperlinkImplIface C.AtkHyperlinkImplIface
type HypertextIface C.AtkHypertextIface
type ImageIface C.AtkImageIface
type Implementor C.AtkImplementor
type KeyEventStruct C.AtkKeyEventStruct
type MiscClass C.AtkMiscClass
type NoOpObjectClass C.AtkNoOpObjectClass
type NoOpObjectFactoryClass C.AtkNoOpObjectFactoryClass
type ObjectClass C.AtkObjectClass
type ObjectFactoryClass C.AtkObjectFactoryClass
type PlugClass C.AtkPlugClass
type PropertyValues C.AtkPropertyValues
type Range C.AtkRange
type Rectangle C.AtkRectangle
type RegistryClass C.AtkRegistryClass
type RelationClass C.AtkRelationClass
type RelationSetClass C.AtkRelationSetClass
type SelectionIface C.AtkSelectionIface
type SocketClass C.AtkSocketClass
type StateSetClass C.AtkStateSetClass
type StreamableContentIface C.AtkStreamableContentIface
type TableIface C.AtkTableIface
type TextIface C.AtkTextIface
type TextRange C.AtkTextRange
type TextRectangle C.AtkTextRectangle
type UtilClass C.AtkUtilClass
type ValueIface C.AtkValueIface
type WindowIface C.AtkWindowIface

// UNSUPPORTED : add_focus_tracker : has callback

// UNSUPPORTED : add_global_event_listener : has callback

// UNSUPPORTED : add_key_event_listener : has callback

func Fn_atk_attribute_set_free(param0 glib.SList) {}

// UNSUPPORTED : focus_tracker_init : has callback

func Fn_atk_focus_tracker_notify(param0 unsafe.Pointer) {}

func Fn_atk_get_default_registry() {
	C.atk_get_default_registry()
}

func Fn_atk_get_root() {
	C.atk_get_root()
}

func Fn_atk_get_toolkit_name() {
	C.atk_get_toolkit_name()
}

func Fn_atk_get_toolkit_version() {
	C.atk_get_toolkit_version()
}

func Fn_atk_relation_type_for_name(param0 string) {}

func Fn_atk_relation_type_get_name(param0 int) {}

func Fn_atk_relation_type_register(param0 string) {}

func Fn_atk_remove_focus_tracker(param0 uint) {}

func Fn_atk_remove_global_event_listener(param0 uint) {}

func Fn_atk_remove_key_event_listener(param0 uint) {}

func Fn_atk_role_for_name(param0 string) {}

func Fn_atk_role_get_localized_name(param0 int) {}

func Fn_atk_role_get_name(param0 int) {}

func Fn_atk_role_register(param0 string) {}

func Fn_atk_state_type_for_name(param0 string) {}

func Fn_atk_state_type_get_name(param0 int) {}

func Fn_atk_state_type_register(param0 string) {}

func Fn_atk_text_attribute_for_name(param0 string) {}

func Fn_atk_text_attribute_get_name(param0 int) {}

func Fn_atk_text_attribute_get_value(param0 int, param1 int) {}

func Fn_atk_text_attribute_register(param0 string) {}

func Fn_atk_value_type_get_localized_name(param0 int) {}

func Fn_atk_value_type_get_name(param0 int) {}

func Fn_atk_gobject_accessible_get_object() {}

func Fn_atk_gobject_accessible_for_object(param0 unsafe.Pointer) {}

func Fn_atk_hyperlink_get_end_index() {}

func Fn_atk_hyperlink_get_n_anchors() {}

func Fn_atk_hyperlink_get_object(param0 int) {}

func Fn_atk_hyperlink_get_start_index() {}

func Fn_atk_hyperlink_get_uri(param0 int) {}

func Fn_atk_hyperlink_is_inline() {}

func Fn_atk_hyperlink_is_valid() {}

func Fn_atk_no_op_object_new(param0 unsafe.Pointer) {}

func Fn_atk_no_op_object_factory_new() {
	C.atk_no_op_object_factory_new()
}

func Fn_atk_object_add_relationship(param0 int, param1 unsafe.Pointer) {}

// UNSUPPORTED : connect_property_change_handler : has callback

func Fn_atk_object_get_description() {}

func Fn_atk_object_get_index_in_parent() {}

func Fn_atk_object_get_layer() {}

func Fn_atk_object_get_mdi_zorder() {}

func Fn_atk_object_get_n_accessible_children() {}

func Fn_atk_object_get_name() {}

func Fn_atk_object_get_parent() {}

func Fn_atk_object_get_role() {}

func Fn_atk_object_initialize(param0 *unsafe.Pointer) {}

func Fn_atk_object_notify_state_change(param0 uint64, param1 bool) {}

func Fn_atk_object_peek_parent() {}

func Fn_atk_object_ref_accessible_child(param0 int) {}

func Fn_atk_object_ref_relation_set() {}

func Fn_atk_object_ref_state_set() {}

func Fn_atk_object_remove_property_change_handler(param0 uint) {}

func Fn_atk_object_remove_relationship(param0 int, param1 unsafe.Pointer) {}

func Fn_atk_object_set_description(param0 string) {}

func Fn_atk_object_set_name(param0 string) {}

func Fn_atk_object_set_parent(param0 unsafe.Pointer) {}

func Fn_atk_object_set_role(param0 int) {}

func Fn_atk_object_factory_create_accessible(param0 unsafe.Pointer) {}

func Fn_atk_object_factory_get_accessible_type() {}

func Fn_atk_object_factory_invalidate() {}

func Fn_atk_registry_get_factory(param0 uint64) {}

func Fn_atk_registry_get_factory_type(param0 uint64) {}

func Fn_atk_registry_set_factory_type(param0 uint64, param1 uint64) {}

func Fn_atk_relation_new(param0 []unsafe.Pointer, param1 int, param2 int) {}

func Fn_atk_relation_get_relation_type() {}

func Fn_atk_relation_get_target() {}

func Fn_atk_relation_remove_target(param0 unsafe.Pointer) {}

func Fn_atk_relation_set_new() {
	C.atk_relation_set_new()
}

func Fn_atk_relation_set_add(param0 unsafe.Pointer) {}

func Fn_atk_relation_set_contains(param0 int) {}

func Fn_atk_relation_set_contains_target(param0 int, param1 unsafe.Pointer) {}

func Fn_atk_relation_set_get_n_relations() {}

func Fn_atk_relation_set_get_relation(param0 int) {}

func Fn_atk_relation_set_get_relation_by_type(param0 int) {}

func Fn_atk_relation_set_remove(param0 unsafe.Pointer) {}

func Fn_atk_socket_new() {
	C.atk_socket_new()
}

func Fn_atk_state_set_new() {
	C.atk_state_set_new()
}

func Fn_atk_state_set_add_state(param0 int) {}

func Fn_atk_state_set_add_states(param0 []int, param1 int) {}

func Fn_atk_state_set_and_sets(param0 unsafe.Pointer) {}

func Fn_atk_state_set_clear_states() {}

func Fn_atk_state_set_contains_state(param0 int) {}

func Fn_atk_state_set_contains_states(param0 []int, param1 int) {}

func Fn_atk_state_set_is_empty() {}

func Fn_atk_state_set_or_sets(param0 unsafe.Pointer) {}

func Fn_atk_state_set_remove_state(param0 int) {}

func Fn_atk_state_set_xor_sets(param0 unsafe.Pointer) {}
