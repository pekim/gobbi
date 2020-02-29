// Code generated - DO NOT EDIT.
// +build atk_2.10

package atk

import (
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
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
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

func Fn_atk_attribute_set_free(param0 *glib.SList) {
	cValue0 := (*C.AtkAttributeSet)(unsafe.Pointer(param0))

	C.atk_attribute_set_free(cValue0)
}

func Fn_atk_implementor_ref_accessible(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkImplementor)(unsafe.Pointer(paramInstance))

	ret := C.atk_implementor_ref_accessible(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

func Fn_atk_focus_tracker_notify(param0 unsafe.Pointer) {
	cValue0 := (*C.AtkObject)(unsafe.Pointer(param0))

	C.atk_focus_tracker_notify(cValue0)
}

func Fn_atk_get_binary_age() uint {
	ret := C.atk_get_binary_age()

	return (uint)(ret)
}

func Fn_atk_get_default_registry() unsafe.Pointer {
	ret := C.atk_get_default_registry()

	return unsafe.Pointer(ret)
}

func Fn_atk_get_focus_object() unsafe.Pointer {
	ret := C.atk_get_focus_object()

	return unsafe.Pointer(ret)
}

func Fn_atk_get_interface_age() uint {
	ret := C.atk_get_interface_age()

	return (uint)(ret)
}

func Fn_atk_get_major_version() uint {
	ret := C.atk_get_major_version()

	return (uint)(ret)
}

func Fn_atk_get_micro_version() uint {
	ret := C.atk_get_micro_version()

	return (uint)(ret)
}

func Fn_atk_get_minor_version() uint {
	ret := C.atk_get_minor_version()

	return (uint)(ret)
}

func Fn_atk_get_root() unsafe.Pointer {
	ret := C.atk_get_root()

	return unsafe.Pointer(ret)
}

func Fn_atk_get_toolkit_name() string {
	ret := C.atk_get_toolkit_name()

	return C.GoString(ret)
}

func Fn_atk_get_toolkit_version() string {
	ret := C.atk_get_toolkit_version()

	return C.GoString(ret)
}

func Fn_atk_get_version() string {
	ret := C.atk_get_version()

	return C.GoString(ret)
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

// UNSUPPORTED : atk_text_free_ranges : parameter 'ranges' is array parameter without length parameter

func Fn_atk_gobject_accessible_get_object(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkGObjectAccessible)(unsafe.Pointer(paramInstance))

	ret := C.atk_gobject_accessible_get_object(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_atk_gobject_accessible_for_object(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	ret := C.atk_gobject_accessible_for_object(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_hyperlink_get_end_index(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_get_end_index(cValueInstance)

	return (int)(ret)
}

func Fn_atk_hyperlink_get_n_anchors(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_get_n_anchors(cValueInstance)

	return (int)(ret)
}

func Fn_atk_hyperlink_get_object(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_hyperlink_get_object(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_hyperlink_get_start_index(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_get_start_index(cValueInstance)

	return (int)(ret)
}

func Fn_atk_hyperlink_get_uri(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_hyperlink_get_uri(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_hyperlink_is_inline(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_is_inline(cValueInstance)

	return toGoBool(ret)
}

func Fn_atk_hyperlink_is_selected_link(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_is_selected_link(cValueInstance)

	return toGoBool(ret)
}

func Fn_atk_hyperlink_is_valid(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkHyperlink)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_is_valid(cValueInstance)

	return toGoBool(ret)
}

func Fn_atk_misc_threads_enter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.AtkMisc)(unsafe.Pointer(paramInstance))

	C.atk_misc_threads_enter(cValueInstance)
}

func Fn_atk_misc_threads_leave(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.AtkMisc)(unsafe.Pointer(paramInstance))

	C.atk_misc_threads_leave(cValueInstance)
}

func Fn_atk_misc_get_instance() unsafe.Pointer {
	ret := C.atk_misc_get_instance()

	return unsafe.Pointer(ret)
}

func Fn_atk_no_op_object_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	ret := C.atk_no_op_object_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_no_op_object_factory_new() unsafe.Pointer {
	ret := C.atk_no_op_object_factory_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_object_add_relationship(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	cValue1 := (*C.AtkObject)(unsafe.Pointer(param1))

	ret := C.atk_object_add_relationship(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : atk_object_connect_property_change_handler : parameter 'handler' is callback

func Fn_atk_object_get_attributes(paramInstance unsafe.Pointer) *glib.SList {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_attributes(cValueInstance)

	return (*glib.SList)(ret)
}

func Fn_atk_object_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_description(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_object_get_index_in_parent(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_index_in_parent(cValueInstance)

	return (int)(ret)
}

func Fn_atk_object_get_layer(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_layer(cValueInstance)

	return (int)(ret)
}

func Fn_atk_object_get_mdi_zorder(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_mdi_zorder(cValueInstance)

	return (int)(ret)
}

func Fn_atk_object_get_n_accessible_children(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_n_accessible_children(cValueInstance)

	return (int)(ret)
}

func Fn_atk_object_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_object_get_object_locale(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_object_locale(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_object_get_parent(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_parent(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_atk_object_get_role(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_get_role(cValueInstance)

	return (int)(ret)
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

	return unsafe.Pointer(ret)
}

func Fn_atk_object_ref_accessible_child(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_object_ref_accessible_child(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_object_ref_relation_set(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_ref_relation_set(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_atk_object_ref_state_set(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkObject)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_ref_state_set(cValueInstance)

	return unsafe.Pointer(ret)
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

	return toGoBool(ret)
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

	return unsafe.Pointer(ret)
}

func Fn_atk_object_factory_get_accessible_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.AtkObjectFactory)(unsafe.Pointer(paramInstance))

	ret := C.atk_object_factory_get_accessible_type(cValueInstance)

	return (uint64)(ret)
}

func Fn_atk_object_factory_invalidate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.AtkObjectFactory)(unsafe.Pointer(paramInstance))

	C.atk_object_factory_invalidate(cValueInstance)
}

func Fn_atk_plug_new() unsafe.Pointer {
	ret := C.atk_plug_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_plug_get_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkPlug)(unsafe.Pointer(paramInstance))

	ret := C.atk_plug_get_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_registry_get_factory(paramInstance unsafe.Pointer, param0 uint64) unsafe.Pointer {
	cValueInstance := (*C.AtkRegistry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.atk_registry_get_factory(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_registry_get_factory_type(paramInstance unsafe.Pointer, param0 uint64) uint64 {
	cValueInstance := (*C.AtkRegistry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.atk_registry_get_factory_type(cValueInstance, cValue0)

	return (uint64)(ret)
}

func Fn_atk_registry_set_factory_type(paramInstance unsafe.Pointer, param0 uint64, param1 uint64) {
	cValueInstance := (*C.AtkRegistry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (C.GType)(param1)

	C.atk_registry_set_factory_type(cValueInstance, cValue0, cValue1)
}

func Fn_atk_relation_new(param0 []unsafe.Pointer, param1 int, param2 int) unsafe.Pointer {
	cValue0 := (**C.AtkObject)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.AtkRelationType)(param2)

	ret := C.atk_relation_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_atk_relation_add_target(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkRelation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkObject)(unsafe.Pointer(param0))

	C.atk_relation_add_target(cValueInstance, cValue0)
}

func Fn_atk_relation_get_relation_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkRelation)(unsafe.Pointer(paramInstance))

	ret := C.atk_relation_get_relation_type(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : atk_relation_get_target : no array length

func Fn_atk_relation_remove_target(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.AtkRelation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkObject)(unsafe.Pointer(param0))

	ret := C.atk_relation_remove_target(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_relation_set_new() unsafe.Pointer {
	ret := C.atk_relation_set_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_relation_set_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkRelation)(unsafe.Pointer(param0))

	C.atk_relation_set_add(cValueInstance, cValue0)
}

func Fn_atk_relation_set_add_relation_by_type(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	cValue1 := (*C.AtkObject)(unsafe.Pointer(param1))

	C.atk_relation_set_add_relation_by_type(cValueInstance, cValue0, cValue1)
}

func Fn_atk_relation_set_contains(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	ret := C.atk_relation_set_contains(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_relation_set_contains_target(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	cValue1 := (*C.AtkObject)(unsafe.Pointer(param1))

	ret := C.atk_relation_set_contains_target(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_atk_relation_set_get_n_relations(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	ret := C.atk_relation_set_get_n_relations(cValueInstance)

	return (int)(ret)
}

func Fn_atk_relation_set_get_relation(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_relation_set_get_relation(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_relation_set_get_relation_by_type(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRelationType)(param0)

	ret := C.atk_relation_set_get_relation_by_type(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_relation_set_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkRelationSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkRelation)(unsafe.Pointer(param0))

	C.atk_relation_set_remove(cValueInstance, cValue0)
}

func Fn_atk_socket_new() unsafe.Pointer {
	ret := C.atk_socket_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_socket_embed(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.AtkSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.atk_socket_embed(cValueInstance, cValue0)
}

func Fn_atk_socket_is_occupied(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkSocket)(unsafe.Pointer(paramInstance))

	ret := C.atk_socket_is_occupied(cValueInstance)

	return toGoBool(ret)
}

func Fn_atk_state_set_new() unsafe.Pointer {
	ret := C.atk_state_set_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_state_set_add_state(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkStateType)(param0)

	ret := C.atk_state_set_add_state(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_state_set_add_states(paramInstance unsafe.Pointer, param0 []int, param1 int) {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkStateType)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	C.atk_state_set_add_states(cValueInstance, cValue0, cValue1)
}

func Fn_atk_state_set_and_sets(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkStateSet)(unsafe.Pointer(param0))

	ret := C.atk_state_set_and_sets(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_state_set_clear_states(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	C.atk_state_set_clear_states(cValueInstance)
}

func Fn_atk_state_set_contains_state(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkStateType)(param0)

	ret := C.atk_state_set_contains_state(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_state_set_contains_states(paramInstance unsafe.Pointer, param0 []int, param1 int) bool {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkStateType)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	ret := C.atk_state_set_contains_states(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_atk_state_set_is_empty(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	ret := C.atk_state_set_is_empty(cValueInstance)

	return toGoBool(ret)
}

func Fn_atk_state_set_or_sets(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkStateSet)(unsafe.Pointer(param0))

	ret := C.atk_state_set_or_sets(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_state_set_remove_state(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkStateType)(param0)

	ret := C.atk_state_set_remove_state(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_state_set_xor_sets(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkStateSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkStateSet)(unsafe.Pointer(param0))

	ret := C.atk_state_set_xor_sets(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_action_do_action(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_action_do_action(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_action_get_description(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.AtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_action_get_description(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_action_get_keybinding(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.AtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_action_get_keybinding(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_action_get_localized_name(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.AtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_action_get_localized_name(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_action_get_n_actions(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkAction)(unsafe.Pointer(paramInstance))

	ret := C.atk_action_get_n_actions(cValueInstance)

	return (int)(ret)
}

func Fn_atk_action_get_name(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.AtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_action_get_name(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_action_set_description(paramInstance unsafe.Pointer, param0 int, param1 string) bool {
	cValueInstance := (*C.AtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.atk_action_set_description(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : atk_component_add_focus_handler : parameter 'handler' is callback

func Fn_atk_component_contains(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) bool {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.AtkCoordType)(param2)

	ret := C.atk_component_contains(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_atk_component_get_alpha(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	ret := C.atk_component_get_alpha(cValueInstance)

	return (float64)(ret)
}

func Fn_atk_component_get_extents(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 *int, param3 *int, param4 int) {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (C.AtkCoordType)(param4)

	C.atk_component_get_extents(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_atk_component_get_layer(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	ret := C.atk_component_get_layer(cValueInstance)

	return (int)(ret)
}

func Fn_atk_component_get_mdi_zorder(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	ret := C.atk_component_get_mdi_zorder(cValueInstance)

	return (int)(ret)
}

func Fn_atk_component_get_position(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 int) {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (C.AtkCoordType)(param2)

	C.atk_component_get_position(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_atk_component_get_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.atk_component_get_size(cValueInstance, cValue0, cValue1)
}

func Fn_atk_component_grab_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	ret := C.atk_component_grab_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_atk_component_ref_accessible_at_point(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) unsafe.Pointer {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.AtkCoordType)(param2)

	ret := C.atk_component_ref_accessible_at_point(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_atk_component_remove_focus_handler(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.atk_component_remove_focus_handler(cValueInstance, cValue0)
}

func Fn_atk_component_set_extents(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 int) bool {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (C.AtkCoordType)(param4)

	ret := C.atk_component_set_extents(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_atk_component_set_position(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) bool {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.AtkCoordType)(param2)

	ret := C.atk_component_set_position(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_atk_component_set_size(paramInstance unsafe.Pointer, param0 int, param1 int) bool {
	cValueInstance := (*C.AtkComponent)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.atk_component_set_size(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_atk_document_get_attribute_value(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_document_get_attribute_value(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_document_get_attributes(paramInstance unsafe.Pointer) *glib.SList {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	ret := C.atk_document_get_attributes(cValueInstance)

	return (*glib.SList)(ret)
}

func Fn_atk_document_get_document(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	ret := C.atk_document_get_document(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_atk_document_get_document_type(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	ret := C.atk_document_get_document_type(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_document_get_locale(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	ret := C.atk_document_get_locale(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_document_set_attribute_value(paramInstance unsafe.Pointer, param0 string, param1 string) bool {
	cValueInstance := (*C.AtkDocument)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.atk_document_set_attribute_value(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_atk_editable_text_copy_text(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.AtkEditableText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.atk_editable_text_copy_text(cValueInstance, cValue0, cValue1)
}

func Fn_atk_editable_text_cut_text(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.AtkEditableText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.atk_editable_text_cut_text(cValueInstance, cValue0, cValue1)
}

func Fn_atk_editable_text_delete_text(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.AtkEditableText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.atk_editable_text_delete_text(cValueInstance, cValue0, cValue1)
}

func Fn_atk_editable_text_insert_text(paramInstance unsafe.Pointer, param0 string, param1 int, param2 *int) {
	cValueInstance := (*C.AtkEditableText)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.atk_editable_text_insert_text(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_atk_editable_text_paste_text(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.AtkEditableText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.atk_editable_text_paste_text(cValueInstance, cValue0)
}

func Fn_atk_editable_text_set_run_attributes(paramInstance unsafe.Pointer, param0 *glib.SList, param1 int, param2 int) bool {
	cValueInstance := (*C.AtkEditableText)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkAttributeSet)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	ret := C.atk_editable_text_set_run_attributes(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_atk_editable_text_set_text_contents(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.AtkEditableText)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.atk_editable_text_set_text_contents(cValueInstance, cValue0)
}

func Fn_atk_hyperlink_impl_get_hyperlink(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkHyperlinkImpl)(unsafe.Pointer(paramInstance))

	ret := C.atk_hyperlink_impl_get_hyperlink(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_atk_hypertext_get_link(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkHypertext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_hypertext_get_link(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_hypertext_get_link_index(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.AtkHypertext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_hypertext_get_link_index(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_atk_hypertext_get_n_links(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkHypertext)(unsafe.Pointer(paramInstance))

	ret := C.atk_hypertext_get_n_links(cValueInstance)

	return (int)(ret)
}

func Fn_atk_image_get_image_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkImage)(unsafe.Pointer(paramInstance))

	ret := C.atk_image_get_image_description(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_image_get_image_locale(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.AtkImage)(unsafe.Pointer(paramInstance))

	ret := C.atk_image_get_image_locale(cValueInstance)

	return C.GoString(ret)
}

func Fn_atk_image_get_image_position(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 int) {
	cValueInstance := (*C.AtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (C.AtkCoordType)(param2)

	C.atk_image_get_image_position(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_atk_image_get_image_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.AtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.atk_image_get_image_size(cValueInstance, cValue0, cValue1)
}

func Fn_atk_image_set_image_description(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.AtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_image_set_image_description(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_selection_add_selection(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_selection_add_selection(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_selection_clear_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkSelection)(unsafe.Pointer(paramInstance))

	ret := C.atk_selection_clear_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_atk_selection_get_selection_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkSelection)(unsafe.Pointer(paramInstance))

	ret := C.atk_selection_get_selection_count(cValueInstance)

	return (int)(ret)
}

func Fn_atk_selection_is_child_selected(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_selection_is_child_selected(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_selection_ref_selection(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_selection_ref_selection(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_selection_remove_selection(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_selection_remove_selection(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_selection_select_all_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.AtkSelection)(unsafe.Pointer(paramInstance))

	ret := C.atk_selection_select_all_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_atk_streamable_content_get_mime_type(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.AtkStreamableContent)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_streamable_content_get_mime_type(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_streamable_content_get_n_mime_types(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkStreamableContent)(unsafe.Pointer(paramInstance))

	ret := C.atk_streamable_content_get_n_mime_types(cValueInstance)

	return (int)(ret)
}

func Fn_atk_streamable_content_get_stream(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.AtkStreamableContent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_streamable_content_get_stream(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_streamable_content_get_uri(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.AtkStreamableContent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.atk_streamable_content_get_uri(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_table_add_column_selection(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_add_column_selection(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_table_add_row_selection(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_add_row_selection(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_table_get_caption(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	ret := C.atk_table_get_caption(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_atk_table_get_column_at_index(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_get_column_at_index(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_atk_table_get_column_description(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_get_column_description(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_table_get_column_extent_at(paramInstance unsafe.Pointer, param0 int, param1 int) int {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.atk_table_get_column_extent_at(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_atk_table_get_column_header(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_get_column_header(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_atk_table_get_index_at(paramInstance unsafe.Pointer, param0 int, param1 int) int {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.atk_table_get_index_at(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_atk_table_get_n_columns(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	ret := C.atk_table_get_n_columns(cValueInstance)

	return (int)(ret)
}

func Fn_atk_table_get_n_rows(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	ret := C.atk_table_get_n_rows(cValueInstance)

	return (int)(ret)
}

func Fn_atk_table_get_row_at_index(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_get_row_at_index(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_atk_table_get_row_description(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_get_row_description(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_atk_table_get_row_extent_at(paramInstance unsafe.Pointer, param0 int, param1 int) int {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.atk_table_get_row_extent_at(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_atk_table_get_row_header(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_get_row_header(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : atk_table_get_selected_columns : parameter 'selected' is non array with indirect count > 1

// UNSUPPORTED : atk_table_get_selected_rows : parameter 'selected' is non array with indirect count > 1

func Fn_atk_table_get_summary(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	ret := C.atk_table_get_summary(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_atk_table_is_column_selected(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_is_column_selected(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_table_is_row_selected(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_is_row_selected(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_table_is_selected(paramInstance unsafe.Pointer, param0 int, param1 int) bool {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.atk_table_is_selected(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_atk_table_ref_at(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.atk_table_ref_at(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_atk_table_remove_column_selection(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_remove_column_selection(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_table_remove_row_selection(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_table_remove_row_selection(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_table_set_caption(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkObject)(unsafe.Pointer(param0))

	C.atk_table_set_caption(cValueInstance, cValue0)
}

func Fn_atk_table_set_column_description(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.atk_table_set_column_description(cValueInstance, cValue0, cValue1)
}

func Fn_atk_table_set_column_header(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.AtkObject)(unsafe.Pointer(param1))

	C.atk_table_set_column_header(cValueInstance, cValue0, cValue1)
}

func Fn_atk_table_set_row_description(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.atk_table_set_row_description(cValueInstance, cValue0, cValue1)
}

func Fn_atk_table_set_row_header(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.AtkObject)(unsafe.Pointer(param1))

	C.atk_table_set_row_header(cValueInstance, cValue0, cValue1)
}

func Fn_atk_table_set_summary(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.AtkObject)(unsafe.Pointer(param0))

	C.atk_table_set_summary(cValueInstance, cValue0)
}

// UNSUPPORTED : atk_table_cell_get_column_header_cells : no array length

// UNSUPPORTED : atk_table_cell_get_row_header_cells : no array length

func Fn_atk_text_add_selection(paramInstance unsafe.Pointer, param0 int, param1 int) bool {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.atk_text_add_selection(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : atk_text_get_bounded_ranges : no array length

func Fn_atk_text_get_caret_offset(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	ret := C.atk_text_get_caret_offset(cValueInstance)

	return (int)(ret)
}

func Fn_atk_text_get_character_at_offset(paramInstance unsafe.Pointer, param0 int) rune {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_text_get_character_at_offset(cValueInstance, cValue0)

	return (rune)(ret)
}

func Fn_atk_text_get_character_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	ret := C.atk_text_get_character_count(cValueInstance)

	return (int)(ret)
}

func Fn_atk_text_get_character_extents(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int, param3 *int, param4 *int, param5 int) {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	cValue5 := (C.AtkCoordType)(param5)

	C.atk_text_get_character_extents(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

// UNSUPPORTED : atk_text_get_default_attributes : blacklisted

func Fn_atk_text_get_n_selections(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	ret := C.atk_text_get_n_selections(cValueInstance)

	return (int)(ret)
}

func Fn_atk_text_get_offset_at_point(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) int {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.AtkCoordType)(param2)

	ret := C.atk_text_get_offset_at_point(cValueInstance, cValue0, cValue1, cValue2)

	return (int)(ret)
}

func Fn_atk_text_get_range_extents(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 unsafe.Pointer) {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.AtkCoordType)(param2)

	cValue3 := (*C.AtkTextRectangle)(unsafe.Pointer(param3))

	C.atk_text_get_range_extents(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : atk_text_get_run_attributes : blacklisted

func Fn_atk_text_get_selection(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) string {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	ret := C.atk_text_get_selection(cValueInstance, cValue0, cValue1, cValue2)

	return C.GoString(ret)
}

func Fn_atk_text_get_string_at_offset(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) string {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.AtkTextGranularity)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	ret := C.atk_text_get_string_at_offset(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return C.GoString(ret)
}

func Fn_atk_text_get_text(paramInstance unsafe.Pointer, param0 int, param1 int) string {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.atk_text_get_text(cValueInstance, cValue0, cValue1)

	return C.GoString(ret)
}

func Fn_atk_text_get_text_after_offset(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) string {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.AtkTextBoundary)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	ret := C.atk_text_get_text_after_offset(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return C.GoString(ret)
}

func Fn_atk_text_get_text_at_offset(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) string {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.AtkTextBoundary)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	ret := C.atk_text_get_text_at_offset(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return C.GoString(ret)
}

func Fn_atk_text_get_text_before_offset(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) string {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.AtkTextBoundary)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	ret := C.atk_text_get_text_before_offset(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return C.GoString(ret)
}

func Fn_atk_text_remove_selection(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_text_remove_selection(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_text_set_caret_offset(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.atk_text_set_caret_offset(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_atk_text_set_selection(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) bool {
	cValueInstance := (*C.AtkText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	ret := C.atk_text_set_selection(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_atk_value_get_current_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	C.atk_value_get_current_value(cValueInstance, cValue0)
}

func Fn_atk_value_get_maximum_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	C.atk_value_get_maximum_value(cValueInstance, cValue0)
}

func Fn_atk_value_get_minimum_increment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	C.atk_value_get_minimum_increment(cValueInstance, cValue0)
}

func Fn_atk_value_get_minimum_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	C.atk_value_get_minimum_value(cValueInstance, cValue0)
}

// UNSUPPORTED : atk_value_get_value_and_text : parameter 'text' is non array with indirect count > 1

func Fn_atk_value_set_current_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.AtkValue)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	ret := C.atk_value_set_current_value(cValueInstance, cValue0)

	return toGoBool(ret)
}
