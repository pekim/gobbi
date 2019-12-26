// Code generated - DO NOT EDIT.
// +build !atk_2.7.4,!atk_2.12,!atk_2.14,!atk_2.30

package atk

import (
	"fmt"
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
	"unsafe"
)

// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.TRUE
}

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

func Fn_atk_attribute_set_free(param0 *glib.SList) {
	cValue0 := (*C.AtkAttributeSet)(unsafe.Pointer(param0))

	C.atk_attribute_set_free(cValue0)
}

// UNSUPPORTED : focus_tracker_init : has callback

func Fn_atk_focus_tracker_notify(param0 unsafe.Pointer) {
	cValue0 := (*C.AtkObject)(unsafe.Pointer(param0))

	C.atk_focus_tracker_notify(cValue0)
}

func Fn_atk_get_default_registry() unsafe.Pointer {
	ret := C.atk_get_default_registry()

	fmt.Println(ret)
}

func Fn_atk_get_root() unsafe.Pointer {
	ret := C.atk_get_root()

	fmt.Println(ret)
}

func Fn_atk_get_toolkit_name() string {
	ret := C.atk_get_toolkit_name()

	fmt.Println(ret)
}

func Fn_atk_get_toolkit_version() string {
	ret := C.atk_get_toolkit_version()

	fmt.Println(ret)
}

func Fn_atk_relation_type_for_name(param0 string) int {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_relation_type_for_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_relation_type_get_name(param0 int) string {
	cValue0 := (C.AtkRelationType)(param0)

	ret := C.atk_relation_type_get_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_relation_type_register(param0 string) int {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_relation_type_register(cValue0)

	fmt.Println(ret)
}

func Fn_atk_remove_focus_tracker(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.atk_remove_focus_tracker(cValue0)
}

func Fn_atk_remove_global_event_listener(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.atk_remove_global_event_listener(cValue0)
}

func Fn_atk_remove_key_event_listener(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.atk_remove_key_event_listener(cValue0)
}

func Fn_atk_role_for_name(param0 string) int {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_role_for_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_role_get_localized_name(param0 int) string {
	cValue0 := (C.AtkRole)(param0)

	ret := C.atk_role_get_localized_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_role_get_name(param0 int) string {
	cValue0 := (C.AtkRole)(param0)

	ret := C.atk_role_get_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_role_register(param0 string) int {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_role_register(cValue0)

	fmt.Println(ret)
}

func Fn_atk_state_type_for_name(param0 string) int {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_state_type_for_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_state_type_get_name(param0 int) string {
	cValue0 := (C.AtkStateType)(param0)

	ret := C.atk_state_type_get_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_state_type_register(param0 string) int {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_state_type_register(cValue0)

	fmt.Println(ret)
}

func Fn_atk_text_attribute_for_name(param0 string) int {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_text_attribute_for_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_text_attribute_get_name(param0 int) string {
	cValue0 := (C.AtkTextAttribute)(param0)

	ret := C.atk_text_attribute_get_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_text_attribute_get_value(param0 int, param1 int) string {
	cValue0 := (C.AtkTextAttribute)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.atk_text_attribute_get_value(cValue0, cValue1)

	fmt.Println(ret)
}

func Fn_atk_text_attribute_register(param0 string) int {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_text_attribute_register(cValue0)

	fmt.Println(ret)
}

func Fn_atk_value_type_get_localized_name(param0 int) string {
	cValue0 := (C.AtkValueType)(param0)

	ret := C.atk_value_type_get_localized_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_value_type_get_name(param0 int) string {
	cValue0 := (C.AtkValueType)(param0)

	ret := C.atk_value_type_get_name(cValue0)

	fmt.Println(ret)
}

func Fn_atk_gobject_accessible_get_object(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkGObjectAccessible)(unsafe.Pointer(paramInstance))

	ret := C.atk_gobject_accessible_get_object(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_gobject_accessible_for_object(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	ret := C.atk_gobject_accessible_for_object(cValue0)

	fmt.Println(ret)
}

func Fn_atk_hyperlink_get_end_index(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_get_end_index(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_hyperlink_get_n_anchors(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_get_n_anchors(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_hyperlink_get_object(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_hyperlink_get_object(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_hyperlink_get_start_index(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_get_start_index(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_hyperlink_get_uri(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_hyperlink_get_uri(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_hyperlink_is_inline(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_is_inline(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_hyperlink_is_valid(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_is_valid(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_no_op_object_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	ret := C.atk_no_op_object_new(cValue0)

	fmt.Println(ret)
}

func Fn_atk_no_op_object_factory_new() unsafe.Pointer {
	ret := C.atk_no_op_object_factory_new()

	fmt.Println(ret)
}

func Fn_atk_object_add_relationship(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	cValue1 := (*C.AtkObject)(unsafe.Pointer(param1))

	ret := C.atk_object_add_relationship(cValueInstance, cValue0, cValue1)

	fmt.Println(ret)
}

// UNSUPPORTED : connect_property_change_handler : has callback

func Fn_atk_object_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_description(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_get_index_in_parent(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_index_in_parent(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_get_layer(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_layer(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_get_mdi_zorder(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_mdi_zorder(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_get_n_accessible_children(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_n_accessible_children(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_name(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_get_parent(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_parent(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_get_role(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_role(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_initialize(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.atk_object_initialize(cValueInstance, cValue0)
}

func Fn_atk_object_notify_state_change(paramInstance unsafe.Pointer, param0 uint64, param1 bool) {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkState)(param0)

	cValue1 := toCBool(param1)

	C.atk_object_notify_state_change(cValueInstance, cValue0, cValue1)
}

func Fn_atk_object_peek_parent(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_peek_parent(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_ref_accessible_child(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_object_ref_accessible_child(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_object_ref_relation_set(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_ref_relation_set(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_ref_state_set(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_ref_state_set(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_remove_property_change_handler(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.atk_object_remove_property_change_handler(cValueInstance, cValue0)
}

func Fn_atk_object_remove_relationship(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	cValue1 := (*C.AtkObject)(unsafe.Pointer(param1))

	ret := C.atk_object_remove_relationship(cValueInstance, cValue0, cValue1)

	fmt.Println(ret)
}

func Fn_atk_object_set_description(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.atk_object_set_description(cValueInstance, cValue0)
}

func Fn_atk_object_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.atk_object_set_name(cValueInstance, cValue0)
}

func Fn_atk_object_set_parent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkObject)(unsafe.Pointer(param0))

	C.atk_object_set_parent(cValueInstance, cValue0)
}

func Fn_atk_object_set_role(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRole)(param0)

	C.atk_object_set_role(cValueInstance, cValue0)
}

func Fn_atk_object_factory_create_accessible(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkObjectFactory)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	ret := C.atk_object_factory_create_accessible(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_object_factory_get_accessible_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.AtkObjectFactory)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_factory_get_accessible_type(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_object_factory_invalidate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.AtkObjectFactory)(unsafe.Pointer(paramInstance))

	C.atk_object_factory_invalidate(cValueInstance)
}

func Fn_atk_registry_get_factory(paramInstance unsafe.Pointer, param0 uint64) unsafe.Pointer {
	cValueInstance := (*C.AtkRegistry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.atk_registry_get_factory(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_registry_get_factory_type(paramInstance unsafe.Pointer, param0 uint64) uint64 {
	cValueInstance := (*C.AtkRegistry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.atk_registry_get_factory_type(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_registry_set_factory_type(paramInstance unsafe.Pointer, param0 uint64, param1 uint64) {
	cValueInstance := (*C.AtkRegistry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	C.atk_registry_set_factory_type(cValueInstance, cValue0, cValue1)
}

func Fn_atk_relation_new(param0 []unsafe.Pointer, param1 int, param2 int) unsafe.Pointer {
	// has non-string array param
}

func Fn_atk_relation_get_relation_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkRelation)(unsafe.Pointer(paramInstance))

	ret := C.atk_relation_get_relation_type(cValueInstance)

	fmt.Println(ret)
}

// UNSUPPORTED : get_target : has array return

func Fn_atk_relation_remove_target(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.AtkRelation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkObject)(unsafe.Pointer(param0))

	ret := C.atk_relation_remove_target(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_relation_set_new() unsafe.Pointer {
	ret := C.atk_relation_set_new()

	fmt.Println(ret)
}

func Fn_atk_relation_set_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkRelation)(unsafe.Pointer(param0))

	C.atk_relation_set_add(cValueInstance, cValue0)
}

func Fn_atk_relation_set_contains(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	ret := C.atk_relation_set_contains(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_relation_set_contains_target(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	cValue1 := (*C.AtkObject)(unsafe.Pointer(param1))

	ret := C.atk_relation_set_contains_target(cValueInstance, cValue0, cValue1)

	fmt.Println(ret)
}

func Fn_atk_relation_set_get_n_relations(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	ret := C.atk_relation_set_get_n_relations(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_relation_set_get_relation(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_relation_set_get_relation(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_relation_set_get_relation_by_type(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	ret := C.atk_relation_set_get_relation_by_type(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_relation_set_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkRelation)(unsafe.Pointer(param0))

	C.atk_relation_set_remove(cValueInstance, cValue0)
}

func Fn_atk_socket_new() unsafe.Pointer {
	ret := C.atk_socket_new()

	fmt.Println(ret)
}

func Fn_atk_state_set_new() unsafe.Pointer {
	ret := C.atk_state_set_new()

	fmt.Println(ret)
}

func Fn_atk_state_set_add_state(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkStateType)(param0)

	ret := C.atk_state_set_add_state(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_state_set_add_states(paramInstance unsafe.Pointer, param0 []int, param1 int) {
	// has non-string array param
}

func Fn_atk_state_set_and_sets(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkStateSet)(unsafe.Pointer(param0))

	ret := C.atk_state_set_and_sets(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_state_set_clear_states(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	C.atk_state_set_clear_states(cValueInstance)
}

func Fn_atk_state_set_contains_state(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkStateType)(param0)

	ret := C.atk_state_set_contains_state(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_state_set_contains_states(paramInstance unsafe.Pointer, param0 []int, param1 int) bool {
	// has non-string array param
}

func Fn_atk_state_set_is_empty(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	ret := C.atk_state_set_is_empty(cValueInstance)

	fmt.Println(ret)
}

func Fn_atk_state_set_or_sets(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkStateSet)(unsafe.Pointer(param0))

	ret := C.atk_state_set_or_sets(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_state_set_remove_state(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkStateType)(param0)

	ret := C.atk_state_set_remove_state(cValueInstance, cValue0)

	fmt.Println(ret)
}

func Fn_atk_state_set_xor_sets(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkStateSet)(unsafe.Pointer(param0))

	ret := C.atk_state_set_xor_sets(cValueInstance, cValue0)

	fmt.Println(ret)
}
