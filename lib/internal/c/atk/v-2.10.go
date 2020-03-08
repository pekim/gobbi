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

func Fn_atk_attribute_set_free(attribSet *glib.SList) {
	c_attribSet := (*C.AtkAttributeSet)(unsafe.Pointer(attribSet))

	C.atk_attribute_set_free(c_attribSet)
}

func Fn_atk_implementor_ref_accessible(implementor unsafe.Pointer) unsafe.Pointer {
	c_implementor := (*C.AtkImplementor)(unsafe.Pointer(implementor))

	ret := C.atk_implementor_ref_accessible(c_implementor)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : atk_add_focus_tracker : parameter 'focus_tracker' is callback

// UNSUPPORTED : atk_add_global_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_add_key_event_listener : parameter 'listener' is callback

// UNSUPPORTED : atk_focus_tracker_init : parameter 'init' is callback

func Fn_atk_focus_tracker_notify(object unsafe.Pointer) {
	c_object := (*C.AtkObject)(unsafe.Pointer(object))

	C.atk_focus_tracker_notify(c_object)
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

func Fn_atk_remove_focus_tracker(trackerId uint) {
	c_trackerId := (C.guint)(trackerId)

	C.atk_remove_focus_tracker(c_trackerId)
}

func Fn_atk_remove_global_event_listener(listenerId uint) {
	c_listenerId := (C.guint)(listenerId)

	C.atk_remove_global_event_listener(c_listenerId)
}

func Fn_atk_remove_key_event_listener(listenerId uint) {
	c_listenerId := (C.guint)(listenerId)

	C.atk_remove_key_event_listener(c_listenerId)
}

func Fn_atk_gobject_accessible_get_object(obj unsafe.Pointer) unsafe.Pointer {
	c_obj := (*C.AtkGObjectAccessible)(unsafe.Pointer(obj))

	ret := C.atk_gobject_accessible_get_object(c_obj)

	return unsafe.Pointer(ret)
}

func Fn_atk_gobject_accessible_for_object(obj unsafe.Pointer) unsafe.Pointer {
	c_obj := (*C.GObject)(unsafe.Pointer(obj))

	ret := C.atk_gobject_accessible_for_object(c_obj)

	return unsafe.Pointer(ret)
}

func Fn_atk_hyperlink_get_end_index(link unsafe.Pointer) int {
	c_link := (*C.AtkHyperlink)(unsafe.Pointer(link))

	ret := C.atk_hyperlink_get_end_index(c_link)

	return (int)(ret)
}

func Fn_atk_hyperlink_get_n_anchors(link unsafe.Pointer) int {
	c_link := (*C.AtkHyperlink)(unsafe.Pointer(link))

	ret := C.atk_hyperlink_get_n_anchors(c_link)

	return (int)(ret)
}

func Fn_atk_hyperlink_get_object(link unsafe.Pointer, i int) unsafe.Pointer {
	c_link := (*C.AtkHyperlink)(unsafe.Pointer(link))

	c_i := (C.gint)(i)

	ret := C.atk_hyperlink_get_object(c_link, c_i)

	return unsafe.Pointer(ret)
}

func Fn_atk_hyperlink_get_start_index(link unsafe.Pointer) int {
	c_link := (*C.AtkHyperlink)(unsafe.Pointer(link))

	ret := C.atk_hyperlink_get_start_index(c_link)

	return (int)(ret)
}

func Fn_atk_hyperlink_get_uri(link unsafe.Pointer, i int) string {
	c_link := (*C.AtkHyperlink)(unsafe.Pointer(link))

	c_i := (C.gint)(i)

	ret := C.atk_hyperlink_get_uri(c_link, c_i)

	return C.GoString(ret)
}

func Fn_atk_hyperlink_is_inline(link unsafe.Pointer) bool {
	c_link := (*C.AtkHyperlink)(unsafe.Pointer(link))

	ret := C.atk_hyperlink_is_inline(c_link)

	return toGoBool(ret)
}

func Fn_atk_hyperlink_is_selected_link(link unsafe.Pointer) bool {
	c_link := (*C.AtkHyperlink)(unsafe.Pointer(link))

	ret := C.atk_hyperlink_is_selected_link(c_link)

	return toGoBool(ret)
}

func Fn_atk_hyperlink_is_valid(link unsafe.Pointer) bool {
	c_link := (*C.AtkHyperlink)(unsafe.Pointer(link))

	ret := C.atk_hyperlink_is_valid(c_link)

	return toGoBool(ret)
}

func Fn_atk_misc_threads_enter(misc unsafe.Pointer) {
	c_misc := (*C.AtkMisc)(unsafe.Pointer(misc))

	C.atk_misc_threads_enter(c_misc)
}

func Fn_atk_misc_threads_leave(misc unsafe.Pointer) {
	c_misc := (*C.AtkMisc)(unsafe.Pointer(misc))

	C.atk_misc_threads_leave(c_misc)
}

func Fn_atk_misc_get_instance() unsafe.Pointer {
	ret := C.atk_misc_get_instance()

	return unsafe.Pointer(ret)
}

func Fn_atk_no_op_object_new(obj unsafe.Pointer) unsafe.Pointer {
	c_obj := (*C.GObject)(unsafe.Pointer(obj))

	ret := C.atk_no_op_object_new(c_obj)

	return unsafe.Pointer(ret)
}

func Fn_atk_no_op_object_factory_new() unsafe.Pointer {
	ret := C.atk_no_op_object_factory_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_object_add_relationship(object unsafe.Pointer, relationship int, target unsafe.Pointer) bool {
	c_object := (*C.AtkObject)(unsafe.Pointer(object))

	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(unsafe.Pointer(target))

	ret := C.atk_object_add_relationship(c_object, c_relationship, c_target)

	return toGoBool(ret)
}

// UNSUPPORTED : atk_object_connect_property_change_handler : parameter 'handler' is callback

func Fn_atk_object_get_attributes(accessible unsafe.Pointer) *glib.SList {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_get_attributes(c_accessible)

	return (*glib.SList)(ret)
}

func Fn_atk_object_get_description(accessible unsafe.Pointer) string {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_get_description(c_accessible)

	return C.GoString(ret)
}

func Fn_atk_object_get_index_in_parent(accessible unsafe.Pointer) int {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_get_index_in_parent(c_accessible)

	return (int)(ret)
}

func Fn_atk_object_get_layer(accessible unsafe.Pointer) int {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_get_layer(c_accessible)

	return (int)(ret)
}

func Fn_atk_object_get_mdi_zorder(accessible unsafe.Pointer) int {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_get_mdi_zorder(c_accessible)

	return (int)(ret)
}

func Fn_atk_object_get_n_accessible_children(accessible unsafe.Pointer) int {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_get_n_accessible_children(c_accessible)

	return (int)(ret)
}

func Fn_atk_object_get_name(accessible unsafe.Pointer) string {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_get_name(c_accessible)

	return C.GoString(ret)
}

func Fn_atk_object_get_object_locale(accessible unsafe.Pointer) string {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_get_object_locale(c_accessible)

	return C.GoString(ret)
}

func Fn_atk_object_get_parent(accessible unsafe.Pointer) unsafe.Pointer {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_get_parent(c_accessible)

	return unsafe.Pointer(ret)
}

func Fn_atk_object_get_role(accessible unsafe.Pointer) int {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_get_role(c_accessible)

	return (int)(ret)
}

func Fn_atk_object_initialize(accessible unsafe.Pointer, data unsafe.Pointer) {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	c_data := (C.gpointer)(data)

	C.atk_object_initialize(c_accessible, c_data)
}

func Fn_atk_object_notify_state_change(accessible unsafe.Pointer, state uint64, value bool) {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	c_state := (C.AtkState)(state)

	c_value := toCBool(value)

	C.atk_object_notify_state_change(c_accessible, c_state, c_value)
}

func Fn_atk_object_peek_parent(accessible unsafe.Pointer) unsafe.Pointer {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_peek_parent(c_accessible)

	return unsafe.Pointer(ret)
}

func Fn_atk_object_ref_accessible_child(accessible unsafe.Pointer, i int) unsafe.Pointer {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	c_i := (C.gint)(i)

	ret := C.atk_object_ref_accessible_child(c_accessible, c_i)

	return unsafe.Pointer(ret)
}

func Fn_atk_object_ref_relation_set(accessible unsafe.Pointer) unsafe.Pointer {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_ref_relation_set(c_accessible)

	return unsafe.Pointer(ret)
}

func Fn_atk_object_ref_state_set(accessible unsafe.Pointer) unsafe.Pointer {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	ret := C.atk_object_ref_state_set(c_accessible)

	return unsafe.Pointer(ret)
}

func Fn_atk_object_remove_property_change_handler(accessible unsafe.Pointer, handlerId uint) {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	c_handlerId := (C.guint)(handlerId)

	C.atk_object_remove_property_change_handler(c_accessible, c_handlerId)
}

func Fn_atk_object_remove_relationship(object unsafe.Pointer, relationship int, target unsafe.Pointer) bool {
	c_object := (*C.AtkObject)(unsafe.Pointer(object))

	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(unsafe.Pointer(target))

	ret := C.atk_object_remove_relationship(c_object, c_relationship, c_target)

	return toGoBool(ret)
}

func Fn_atk_object_set_description(accessible unsafe.Pointer, description string) {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	c_description := (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(c_description))

	C.atk_object_set_description(c_accessible, c_description)
}

func Fn_atk_object_set_name(accessible unsafe.Pointer, name string) {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	C.atk_object_set_name(c_accessible, c_name)
}

func Fn_atk_object_set_parent(accessible unsafe.Pointer, parent unsafe.Pointer) {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	c_parent := (*C.AtkObject)(unsafe.Pointer(parent))

	C.atk_object_set_parent(c_accessible, c_parent)
}

func Fn_atk_object_set_role(accessible unsafe.Pointer, role int) {
	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	c_role := (C.AtkRole)(role)

	C.atk_object_set_role(c_accessible, c_role)
}

func Fn_atk_object_factory_create_accessible(factory unsafe.Pointer, obj unsafe.Pointer) unsafe.Pointer {
	c_factory := (*C.AtkObjectFactory)(unsafe.Pointer(factory))

	c_obj := (*C.GObject)(unsafe.Pointer(obj))

	ret := C.atk_object_factory_create_accessible(c_factory, c_obj)

	return unsafe.Pointer(ret)
}

func Fn_atk_object_factory_get_accessible_type(factory unsafe.Pointer) uint64 {
	c_factory := (*C.AtkObjectFactory)(unsafe.Pointer(factory))

	ret := C.atk_object_factory_get_accessible_type(c_factory)

	return (uint64)(ret)
}

func Fn_atk_object_factory_invalidate(factory unsafe.Pointer) {
	c_factory := (*C.AtkObjectFactory)(unsafe.Pointer(factory))

	C.atk_object_factory_invalidate(c_factory)
}

func Fn_atk_plug_new() unsafe.Pointer {
	ret := C.atk_plug_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_plug_get_id(plug unsafe.Pointer) string {
	c_plug := (*C.AtkPlug)(unsafe.Pointer(plug))

	ret := C.atk_plug_get_id(c_plug)

	return C.GoString(ret)
}

func Fn_atk_registry_get_factory(registry unsafe.Pointer, type_ uint64) unsafe.Pointer {
	c_registry := (*C.AtkRegistry)(unsafe.Pointer(registry))

	c_type_ := (C.GType)(type_)

	ret := C.atk_registry_get_factory(c_registry, c_type_)

	return unsafe.Pointer(ret)
}

func Fn_atk_registry_get_factory_type(registry unsafe.Pointer, type_ uint64) uint64 {
	c_registry := (*C.AtkRegistry)(unsafe.Pointer(registry))

	c_type_ := (C.GType)(type_)

	ret := C.atk_registry_get_factory_type(c_registry, c_type_)

	return (uint64)(ret)
}

func Fn_atk_registry_set_factory_type(registry unsafe.Pointer, type_ uint64, factoryType uint64) {
	c_registry := (*C.AtkRegistry)(unsafe.Pointer(registry))

	c_type_ := (C.GType)(type_)

	c_factoryType := (C.GType)(factoryType)

	C.atk_registry_set_factory_type(c_registry, c_type_, c_factoryType)
}

func Fn_atk_relation_new(targets []unsafe.Pointer, nTargets int, relationship int) unsafe.Pointer {
	c_targets := (**C.AtkObject)(unsafe.Pointer(&targets[0]))

	c_nTargets := (C.gint)(nTargets)

	c_relationship := (C.AtkRelationType)(relationship)

	ret := C.atk_relation_new(c_targets, c_nTargets, c_relationship)

	return unsafe.Pointer(ret)
}

func Fn_atk_relation_add_target(relation unsafe.Pointer, target unsafe.Pointer) {
	c_relation := (*C.AtkRelation)(unsafe.Pointer(relation))

	c_target := (*C.AtkObject)(unsafe.Pointer(target))

	C.atk_relation_add_target(c_relation, c_target)
}

func Fn_atk_relation_get_relation_type(relation unsafe.Pointer) int {
	c_relation := (*C.AtkRelation)(unsafe.Pointer(relation))

	ret := C.atk_relation_get_relation_type(c_relation)

	return (int)(ret)
}

// UNSUPPORTED : atk_relation_get_target : no array length

func Fn_atk_relation_remove_target(relation unsafe.Pointer, target unsafe.Pointer) bool {
	c_relation := (*C.AtkRelation)(unsafe.Pointer(relation))

	c_target := (*C.AtkObject)(unsafe.Pointer(target))

	ret := C.atk_relation_remove_target(c_relation, c_target)

	return toGoBool(ret)
}

func Fn_atk_relation_set_new() unsafe.Pointer {
	ret := C.atk_relation_set_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_relation_set_add(set unsafe.Pointer, relation unsafe.Pointer) {
	c_set := (*C.AtkRelationSet)(unsafe.Pointer(set))

	c_relation := (*C.AtkRelation)(unsafe.Pointer(relation))

	C.atk_relation_set_add(c_set, c_relation)
}

func Fn_atk_relation_set_add_relation_by_type(set unsafe.Pointer, relationship int, target unsafe.Pointer) {
	c_set := (*C.AtkRelationSet)(unsafe.Pointer(set))

	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(unsafe.Pointer(target))

	C.atk_relation_set_add_relation_by_type(c_set, c_relationship, c_target)
}

func Fn_atk_relation_set_contains(set unsafe.Pointer, relationship int) bool {
	c_set := (*C.AtkRelationSet)(unsafe.Pointer(set))

	c_relationship := (C.AtkRelationType)(relationship)

	ret := C.atk_relation_set_contains(c_set, c_relationship)

	return toGoBool(ret)
}

func Fn_atk_relation_set_contains_target(set unsafe.Pointer, relationship int, target unsafe.Pointer) bool {
	c_set := (*C.AtkRelationSet)(unsafe.Pointer(set))

	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(unsafe.Pointer(target))

	ret := C.atk_relation_set_contains_target(c_set, c_relationship, c_target)

	return toGoBool(ret)
}

func Fn_atk_relation_set_get_n_relations(set unsafe.Pointer) int {
	c_set := (*C.AtkRelationSet)(unsafe.Pointer(set))

	ret := C.atk_relation_set_get_n_relations(c_set)

	return (int)(ret)
}

func Fn_atk_relation_set_get_relation(set unsafe.Pointer, i int) unsafe.Pointer {
	c_set := (*C.AtkRelationSet)(unsafe.Pointer(set))

	c_i := (C.gint)(i)

	ret := C.atk_relation_set_get_relation(c_set, c_i)

	return unsafe.Pointer(ret)
}

func Fn_atk_relation_set_get_relation_by_type(set unsafe.Pointer, relationship int) unsafe.Pointer {
	c_set := (*C.AtkRelationSet)(unsafe.Pointer(set))

	c_relationship := (C.AtkRelationType)(relationship)

	ret := C.atk_relation_set_get_relation_by_type(c_set, c_relationship)

	return unsafe.Pointer(ret)
}

func Fn_atk_relation_set_remove(set unsafe.Pointer, relation unsafe.Pointer) {
	c_set := (*C.AtkRelationSet)(unsafe.Pointer(set))

	c_relation := (*C.AtkRelation)(unsafe.Pointer(relation))

	C.atk_relation_set_remove(c_set, c_relation)
}

func Fn_atk_socket_new() unsafe.Pointer {
	ret := C.atk_socket_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_socket_embed(obj unsafe.Pointer, plugId string) {
	c_obj := (*C.AtkSocket)(unsafe.Pointer(obj))

	c_plugId := (*C.gchar)(C.CString(plugId))
	defer C.free(unsafe.Pointer(c_plugId))

	C.atk_socket_embed(c_obj, c_plugId)
}

func Fn_atk_socket_is_occupied(obj unsafe.Pointer) bool {
	c_obj := (*C.AtkSocket)(unsafe.Pointer(obj))

	ret := C.atk_socket_is_occupied(c_obj)

	return toGoBool(ret)
}

func Fn_atk_state_set_new() unsafe.Pointer {
	ret := C.atk_state_set_new()

	return unsafe.Pointer(ret)
}

func Fn_atk_state_set_add_state(set unsafe.Pointer, type_ int) bool {
	c_set := (*C.AtkStateSet)(unsafe.Pointer(set))

	c_type_ := (C.AtkStateType)(type_)

	ret := C.atk_state_set_add_state(c_set, c_type_)

	return toGoBool(ret)
}

func Fn_atk_state_set_add_states(set unsafe.Pointer, types []int, nTypes int) {
	c_set := (*C.AtkStateSet)(unsafe.Pointer(set))

	c_types := (*C.AtkStateType)(unsafe.Pointer(&types[0]))

	c_nTypes := (C.gint)(nTypes)

	C.atk_state_set_add_states(c_set, c_types, c_nTypes)
}

func Fn_atk_state_set_and_sets(set unsafe.Pointer, compareSet unsafe.Pointer) unsafe.Pointer {
	c_set := (*C.AtkStateSet)(unsafe.Pointer(set))

	c_compareSet := (*C.AtkStateSet)(unsafe.Pointer(compareSet))

	ret := C.atk_state_set_and_sets(c_set, c_compareSet)

	return unsafe.Pointer(ret)
}

func Fn_atk_state_set_clear_states(set unsafe.Pointer) {
	c_set := (*C.AtkStateSet)(unsafe.Pointer(set))

	C.atk_state_set_clear_states(c_set)
}

func Fn_atk_state_set_contains_state(set unsafe.Pointer, type_ int) bool {
	c_set := (*C.AtkStateSet)(unsafe.Pointer(set))

	c_type_ := (C.AtkStateType)(type_)

	ret := C.atk_state_set_contains_state(c_set, c_type_)

	return toGoBool(ret)
}

func Fn_atk_state_set_contains_states(set unsafe.Pointer, types []int, nTypes int) bool {
	c_set := (*C.AtkStateSet)(unsafe.Pointer(set))

	c_types := (*C.AtkStateType)(unsafe.Pointer(&types[0]))

	c_nTypes := (C.gint)(nTypes)

	ret := C.atk_state_set_contains_states(c_set, c_types, c_nTypes)

	return toGoBool(ret)
}

func Fn_atk_state_set_is_empty(set unsafe.Pointer) bool {
	c_set := (*C.AtkStateSet)(unsafe.Pointer(set))

	ret := C.atk_state_set_is_empty(c_set)

	return toGoBool(ret)
}

func Fn_atk_state_set_or_sets(set unsafe.Pointer, compareSet unsafe.Pointer) unsafe.Pointer {
	c_set := (*C.AtkStateSet)(unsafe.Pointer(set))

	c_compareSet := (*C.AtkStateSet)(unsafe.Pointer(compareSet))

	ret := C.atk_state_set_or_sets(c_set, c_compareSet)

	return unsafe.Pointer(ret)
}

func Fn_atk_state_set_remove_state(set unsafe.Pointer, type_ int) bool {
	c_set := (*C.AtkStateSet)(unsafe.Pointer(set))

	c_type_ := (C.AtkStateType)(type_)

	ret := C.atk_state_set_remove_state(c_set, c_type_)

	return toGoBool(ret)
}

func Fn_atk_state_set_xor_sets(set unsafe.Pointer, compareSet unsafe.Pointer) unsafe.Pointer {
	c_set := (*C.AtkStateSet)(unsafe.Pointer(set))

	c_compareSet := (*C.AtkStateSet)(unsafe.Pointer(compareSet))

	ret := C.atk_state_set_xor_sets(c_set, c_compareSet)

	return unsafe.Pointer(ret)
}

func Fn_atk_action_do_action(action unsafe.Pointer, i int) bool {
	c_action := (*C.AtkAction)(unsafe.Pointer(action))

	c_i := (C.gint)(i)

	ret := C.atk_action_do_action(c_action, c_i)

	return toGoBool(ret)
}

func Fn_atk_action_get_description(action unsafe.Pointer, i int) string {
	c_action := (*C.AtkAction)(unsafe.Pointer(action))

	c_i := (C.gint)(i)

	ret := C.atk_action_get_description(c_action, c_i)

	return C.GoString(ret)
}

func Fn_atk_action_get_keybinding(action unsafe.Pointer, i int) string {
	c_action := (*C.AtkAction)(unsafe.Pointer(action))

	c_i := (C.gint)(i)

	ret := C.atk_action_get_keybinding(c_action, c_i)

	return C.GoString(ret)
}

func Fn_atk_action_get_localized_name(action unsafe.Pointer, i int) string {
	c_action := (*C.AtkAction)(unsafe.Pointer(action))

	c_i := (C.gint)(i)

	ret := C.atk_action_get_localized_name(c_action, c_i)

	return C.GoString(ret)
}

func Fn_atk_action_get_n_actions(action unsafe.Pointer) int {
	c_action := (*C.AtkAction)(unsafe.Pointer(action))

	ret := C.atk_action_get_n_actions(c_action)

	return (int)(ret)
}

func Fn_atk_action_get_name(action unsafe.Pointer, i int) string {
	c_action := (*C.AtkAction)(unsafe.Pointer(action))

	c_i := (C.gint)(i)

	ret := C.atk_action_get_name(c_action, c_i)

	return C.GoString(ret)
}

func Fn_atk_action_set_description(action unsafe.Pointer, i int, desc string) bool {
	c_action := (*C.AtkAction)(unsafe.Pointer(action))

	c_i := (C.gint)(i)

	c_desc := (*C.gchar)(C.CString(desc))
	defer C.free(unsafe.Pointer(c_desc))

	ret := C.atk_action_set_description(c_action, c_i, c_desc)

	return toGoBool(ret)
}

// UNSUPPORTED : atk_component_add_focus_handler : parameter 'handler' is callback

func Fn_atk_component_contains(component unsafe.Pointer, x int, y int, coordType int) bool {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coordType := (C.AtkCoordType)(coordType)

	ret := C.atk_component_contains(c_component, c_x, c_y, c_coordType)

	return toGoBool(ret)
}

func Fn_atk_component_get_alpha(component unsafe.Pointer) float64 {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	ret := C.atk_component_get_alpha(c_component)

	return (float64)(ret)
}

func Fn_atk_component_get_extents(component unsafe.Pointer, x *int, y *int, width *int, height *int, coordType int) {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	c_coordType := (C.AtkCoordType)(coordType)

	C.atk_component_get_extents(c_component, c_x, c_y, c_width, c_height, c_coordType)
}

func Fn_atk_component_get_layer(component unsafe.Pointer) int {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	ret := C.atk_component_get_layer(c_component)

	return (int)(ret)
}

func Fn_atk_component_get_mdi_zorder(component unsafe.Pointer) int {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	ret := C.atk_component_get_mdi_zorder(c_component)

	return (int)(ret)
}

func Fn_atk_component_get_position(component unsafe.Pointer, x *int, y *int, coordType int) {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	c_coordType := (C.AtkCoordType)(coordType)

	C.atk_component_get_position(c_component, c_x, c_y, c_coordType)
}

func Fn_atk_component_get_size(component unsafe.Pointer, width *int, height *int) {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	C.atk_component_get_size(c_component, c_width, c_height)
}

func Fn_atk_component_grab_focus(component unsafe.Pointer) bool {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	ret := C.atk_component_grab_focus(c_component)

	return toGoBool(ret)
}

func Fn_atk_component_ref_accessible_at_point(component unsafe.Pointer, x int, y int, coordType int) unsafe.Pointer {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coordType := (C.AtkCoordType)(coordType)

	ret := C.atk_component_ref_accessible_at_point(c_component, c_x, c_y, c_coordType)

	return unsafe.Pointer(ret)
}

func Fn_atk_component_remove_focus_handler(component unsafe.Pointer, handlerId uint) {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	c_handlerId := (C.guint)(handlerId)

	C.atk_component_remove_focus_handler(c_component, c_handlerId)
}

func Fn_atk_component_set_extents(component unsafe.Pointer, x int, y int, width int, height int, coordType int) bool {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_coordType := (C.AtkCoordType)(coordType)

	ret := C.atk_component_set_extents(c_component, c_x, c_y, c_width, c_height, c_coordType)

	return toGoBool(ret)
}

func Fn_atk_component_set_position(component unsafe.Pointer, x int, y int, coordType int) bool {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coordType := (C.AtkCoordType)(coordType)

	ret := C.atk_component_set_position(c_component, c_x, c_y, c_coordType)

	return toGoBool(ret)
}

func Fn_atk_component_set_size(component unsafe.Pointer, width int, height int) bool {
	c_component := (*C.AtkComponent)(unsafe.Pointer(component))

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	ret := C.atk_component_set_size(c_component, c_width, c_height)

	return toGoBool(ret)
}

func Fn_atk_document_get_attribute_value(document unsafe.Pointer, attributeName string) string {
	c_document := (*C.AtkDocument)(unsafe.Pointer(document))

	c_attributeName := (*C.gchar)(C.CString(attributeName))
	defer C.free(unsafe.Pointer(c_attributeName))

	ret := C.atk_document_get_attribute_value(c_document, c_attributeName)

	return C.GoString(ret)
}

func Fn_atk_document_get_attributes(document unsafe.Pointer) *glib.SList {
	c_document := (*C.AtkDocument)(unsafe.Pointer(document))

	ret := C.atk_document_get_attributes(c_document)

	return (*glib.SList)(ret)
}

func Fn_atk_document_get_document(document unsafe.Pointer) unsafe.Pointer {
	c_document := (*C.AtkDocument)(unsafe.Pointer(document))

	ret := C.atk_document_get_document(c_document)

	return (unsafe.Pointer)(ret)
}

func Fn_atk_document_get_document_type(document unsafe.Pointer) string {
	c_document := (*C.AtkDocument)(unsafe.Pointer(document))

	ret := C.atk_document_get_document_type(c_document)

	return C.GoString(ret)
}

func Fn_atk_document_get_locale(document unsafe.Pointer) string {
	c_document := (*C.AtkDocument)(unsafe.Pointer(document))

	ret := C.atk_document_get_locale(c_document)

	return C.GoString(ret)
}

func Fn_atk_document_set_attribute_value(document unsafe.Pointer, attributeName string, attributeValue string) bool {
	c_document := (*C.AtkDocument)(unsafe.Pointer(document))

	c_attributeName := (*C.gchar)(C.CString(attributeName))
	defer C.free(unsafe.Pointer(c_attributeName))

	c_attributeValue := (*C.gchar)(C.CString(attributeValue))
	defer C.free(unsafe.Pointer(c_attributeValue))

	ret := C.atk_document_set_attribute_value(c_document, c_attributeName, c_attributeValue)

	return toGoBool(ret)
}

func Fn_atk_editable_text_copy_text(text unsafe.Pointer, startPos int, endPos int) {
	c_text := (*C.AtkEditableText)(unsafe.Pointer(text))

	c_startPos := (C.gint)(startPos)

	c_endPos := (C.gint)(endPos)

	C.atk_editable_text_copy_text(c_text, c_startPos, c_endPos)
}

func Fn_atk_editable_text_cut_text(text unsafe.Pointer, startPos int, endPos int) {
	c_text := (*C.AtkEditableText)(unsafe.Pointer(text))

	c_startPos := (C.gint)(startPos)

	c_endPos := (C.gint)(endPos)

	C.atk_editable_text_cut_text(c_text, c_startPos, c_endPos)
}

func Fn_atk_editable_text_delete_text(text unsafe.Pointer, startPos int, endPos int) {
	c_text := (*C.AtkEditableText)(unsafe.Pointer(text))

	c_startPos := (C.gint)(startPos)

	c_endPos := (C.gint)(endPos)

	C.atk_editable_text_delete_text(c_text, c_startPos, c_endPos)
}

func Fn_atk_editable_text_insert_text(text unsafe.Pointer, string_ string, length int, position *int) {
	c_text := (*C.AtkEditableText)(unsafe.Pointer(text))

	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	c_length := (C.gint)(length)

	c_position := (*C.gint)(unsafe.Pointer(position))

	C.atk_editable_text_insert_text(c_text, c_string_, c_length, c_position)
}

func Fn_atk_editable_text_paste_text(text unsafe.Pointer, position int) {
	c_text := (*C.AtkEditableText)(unsafe.Pointer(text))

	c_position := (C.gint)(position)

	C.atk_editable_text_paste_text(c_text, c_position)
}

func Fn_atk_editable_text_set_run_attributes(text unsafe.Pointer, attribSet *glib.SList, startOffset int, endOffset int) bool {
	c_text := (*C.AtkEditableText)(unsafe.Pointer(text))

	c_attribSet := (*C.AtkAttributeSet)(unsafe.Pointer(attribSet))

	c_startOffset := (C.gint)(startOffset)

	c_endOffset := (C.gint)(endOffset)

	ret := C.atk_editable_text_set_run_attributes(c_text, c_attribSet, c_startOffset, c_endOffset)

	return toGoBool(ret)
}

func Fn_atk_editable_text_set_text_contents(text unsafe.Pointer, string_ string) {
	c_text := (*C.AtkEditableText)(unsafe.Pointer(text))

	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	C.atk_editable_text_set_text_contents(c_text, c_string_)
}

func Fn_atk_hyperlink_impl_get_hyperlink(impl unsafe.Pointer) unsafe.Pointer {
	c_impl := (*C.AtkHyperlinkImpl)(unsafe.Pointer(impl))

	ret := C.atk_hyperlink_impl_get_hyperlink(c_impl)

	return unsafe.Pointer(ret)
}

func Fn_atk_hypertext_get_link(hypertext unsafe.Pointer, linkIndex int) unsafe.Pointer {
	c_hypertext := (*C.AtkHypertext)(unsafe.Pointer(hypertext))

	c_linkIndex := (C.gint)(linkIndex)

	ret := C.atk_hypertext_get_link(c_hypertext, c_linkIndex)

	return unsafe.Pointer(ret)
}

func Fn_atk_hypertext_get_link_index(hypertext unsafe.Pointer, charIndex int) int {
	c_hypertext := (*C.AtkHypertext)(unsafe.Pointer(hypertext))

	c_charIndex := (C.gint)(charIndex)

	ret := C.atk_hypertext_get_link_index(c_hypertext, c_charIndex)

	return (int)(ret)
}

func Fn_atk_hypertext_get_n_links(hypertext unsafe.Pointer) int {
	c_hypertext := (*C.AtkHypertext)(unsafe.Pointer(hypertext))

	ret := C.atk_hypertext_get_n_links(c_hypertext)

	return (int)(ret)
}

func Fn_atk_image_get_image_description(image unsafe.Pointer) string {
	c_image := (*C.AtkImage)(unsafe.Pointer(image))

	ret := C.atk_image_get_image_description(c_image)

	return C.GoString(ret)
}

func Fn_atk_image_get_image_locale(image unsafe.Pointer) string {
	c_image := (*C.AtkImage)(unsafe.Pointer(image))

	ret := C.atk_image_get_image_locale(c_image)

	return C.GoString(ret)
}

func Fn_atk_image_get_image_position(image unsafe.Pointer, x *int, y *int, coordType int) {
	c_image := (*C.AtkImage)(unsafe.Pointer(image))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	c_coordType := (C.AtkCoordType)(coordType)

	C.atk_image_get_image_position(c_image, c_x, c_y, c_coordType)
}

func Fn_atk_image_get_image_size(image unsafe.Pointer, width *int, height *int) {
	c_image := (*C.AtkImage)(unsafe.Pointer(image))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	C.atk_image_get_image_size(c_image, c_width, c_height)
}

func Fn_atk_image_set_image_description(image unsafe.Pointer, description string) bool {
	c_image := (*C.AtkImage)(unsafe.Pointer(image))

	c_description := (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(c_description))

	ret := C.atk_image_set_image_description(c_image, c_description)

	return toGoBool(ret)
}

func Fn_atk_selection_add_selection(selection unsafe.Pointer, i int) bool {
	c_selection := (*C.AtkSelection)(unsafe.Pointer(selection))

	c_i := (C.gint)(i)

	ret := C.atk_selection_add_selection(c_selection, c_i)

	return toGoBool(ret)
}

func Fn_atk_selection_clear_selection(selection unsafe.Pointer) bool {
	c_selection := (*C.AtkSelection)(unsafe.Pointer(selection))

	ret := C.atk_selection_clear_selection(c_selection)

	return toGoBool(ret)
}

func Fn_atk_selection_get_selection_count(selection unsafe.Pointer) int {
	c_selection := (*C.AtkSelection)(unsafe.Pointer(selection))

	ret := C.atk_selection_get_selection_count(c_selection)

	return (int)(ret)
}

func Fn_atk_selection_is_child_selected(selection unsafe.Pointer, i int) bool {
	c_selection := (*C.AtkSelection)(unsafe.Pointer(selection))

	c_i := (C.gint)(i)

	ret := C.atk_selection_is_child_selected(c_selection, c_i)

	return toGoBool(ret)
}

func Fn_atk_selection_ref_selection(selection unsafe.Pointer, i int) unsafe.Pointer {
	c_selection := (*C.AtkSelection)(unsafe.Pointer(selection))

	c_i := (C.gint)(i)

	ret := C.atk_selection_ref_selection(c_selection, c_i)

	return unsafe.Pointer(ret)
}

func Fn_atk_selection_remove_selection(selection unsafe.Pointer, i int) bool {
	c_selection := (*C.AtkSelection)(unsafe.Pointer(selection))

	c_i := (C.gint)(i)

	ret := C.atk_selection_remove_selection(c_selection, c_i)

	return toGoBool(ret)
}

func Fn_atk_selection_select_all_selection(selection unsafe.Pointer) bool {
	c_selection := (*C.AtkSelection)(unsafe.Pointer(selection))

	ret := C.atk_selection_select_all_selection(c_selection)

	return toGoBool(ret)
}

func Fn_atk_streamable_content_get_mime_type(streamable unsafe.Pointer, i int) string {
	c_streamable := (*C.AtkStreamableContent)(unsafe.Pointer(streamable))

	c_i := (C.gint)(i)

	ret := C.atk_streamable_content_get_mime_type(c_streamable, c_i)

	return C.GoString(ret)
}

func Fn_atk_streamable_content_get_n_mime_types(streamable unsafe.Pointer) int {
	c_streamable := (*C.AtkStreamableContent)(unsafe.Pointer(streamable))

	ret := C.atk_streamable_content_get_n_mime_types(c_streamable)

	return (int)(ret)
}

func Fn_atk_streamable_content_get_stream(streamable unsafe.Pointer, mimeType string) unsafe.Pointer {
	c_streamable := (*C.AtkStreamableContent)(unsafe.Pointer(streamable))

	c_mimeType := (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(c_mimeType))

	ret := C.atk_streamable_content_get_stream(c_streamable, c_mimeType)

	return unsafe.Pointer(ret)
}

func Fn_atk_streamable_content_get_uri(streamable unsafe.Pointer, mimeType string) string {
	c_streamable := (*C.AtkStreamableContent)(unsafe.Pointer(streamable))

	c_mimeType := (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(c_mimeType))

	ret := C.atk_streamable_content_get_uri(c_streamable, c_mimeType)

	return C.GoString(ret)
}

func Fn_atk_table_add_column_selection(table unsafe.Pointer, column int) bool {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_column := (C.gint)(column)

	ret := C.atk_table_add_column_selection(c_table, c_column)

	return toGoBool(ret)
}

func Fn_atk_table_add_row_selection(table unsafe.Pointer, row int) bool {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	ret := C.atk_table_add_row_selection(c_table, c_row)

	return toGoBool(ret)
}

func Fn_atk_table_get_caption(table unsafe.Pointer) unsafe.Pointer {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	ret := C.atk_table_get_caption(c_table)

	return unsafe.Pointer(ret)
}

func Fn_atk_table_get_column_at_index(table unsafe.Pointer, index int) int {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_index := (C.gint)(index)

	ret := C.atk_table_get_column_at_index(c_table, c_index)

	return (int)(ret)
}

func Fn_atk_table_get_column_description(table unsafe.Pointer, column int) string {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_column := (C.gint)(column)

	ret := C.atk_table_get_column_description(c_table, c_column)

	return C.GoString(ret)
}

func Fn_atk_table_get_column_extent_at(table unsafe.Pointer, row int, column int) int {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	ret := C.atk_table_get_column_extent_at(c_table, c_row, c_column)

	return (int)(ret)
}

func Fn_atk_table_get_column_header(table unsafe.Pointer, column int) unsafe.Pointer {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_column := (C.gint)(column)

	ret := C.atk_table_get_column_header(c_table, c_column)

	return unsafe.Pointer(ret)
}

func Fn_atk_table_get_index_at(table unsafe.Pointer, row int, column int) int {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	ret := C.atk_table_get_index_at(c_table, c_row, c_column)

	return (int)(ret)
}

func Fn_atk_table_get_n_columns(table unsafe.Pointer) int {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	ret := C.atk_table_get_n_columns(c_table)

	return (int)(ret)
}

func Fn_atk_table_get_n_rows(table unsafe.Pointer) int {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	ret := C.atk_table_get_n_rows(c_table)

	return (int)(ret)
}

func Fn_atk_table_get_row_at_index(table unsafe.Pointer, index int) int {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_index := (C.gint)(index)

	ret := C.atk_table_get_row_at_index(c_table, c_index)

	return (int)(ret)
}

func Fn_atk_table_get_row_description(table unsafe.Pointer, row int) string {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	ret := C.atk_table_get_row_description(c_table, c_row)

	return C.GoString(ret)
}

func Fn_atk_table_get_row_extent_at(table unsafe.Pointer, row int, column int) int {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	ret := C.atk_table_get_row_extent_at(c_table, c_row, c_column)

	return (int)(ret)
}

func Fn_atk_table_get_row_header(table unsafe.Pointer, row int) unsafe.Pointer {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	ret := C.atk_table_get_row_header(c_table, c_row)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : atk_table_get_selected_columns : parameter 'selected' is non array with indirect count > 1

// UNSUPPORTED : atk_table_get_selected_rows : parameter 'selected' is non array with indirect count > 1

func Fn_atk_table_get_summary(table unsafe.Pointer) unsafe.Pointer {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	ret := C.atk_table_get_summary(c_table)

	return unsafe.Pointer(ret)
}

func Fn_atk_table_is_column_selected(table unsafe.Pointer, column int) bool {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_column := (C.gint)(column)

	ret := C.atk_table_is_column_selected(c_table, c_column)

	return toGoBool(ret)
}

func Fn_atk_table_is_row_selected(table unsafe.Pointer, row int) bool {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	ret := C.atk_table_is_row_selected(c_table, c_row)

	return toGoBool(ret)
}

func Fn_atk_table_is_selected(table unsafe.Pointer, row int, column int) bool {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	ret := C.atk_table_is_selected(c_table, c_row, c_column)

	return toGoBool(ret)
}

func Fn_atk_table_ref_at(table unsafe.Pointer, row int, column int) unsafe.Pointer {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	ret := C.atk_table_ref_at(c_table, c_row, c_column)

	return unsafe.Pointer(ret)
}

func Fn_atk_table_remove_column_selection(table unsafe.Pointer, column int) bool {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_column := (C.gint)(column)

	ret := C.atk_table_remove_column_selection(c_table, c_column)

	return toGoBool(ret)
}

func Fn_atk_table_remove_row_selection(table unsafe.Pointer, row int) bool {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	ret := C.atk_table_remove_row_selection(c_table, c_row)

	return toGoBool(ret)
}

func Fn_atk_table_set_caption(table unsafe.Pointer, caption unsafe.Pointer) {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_caption := (*C.AtkObject)(unsafe.Pointer(caption))

	C.atk_table_set_caption(c_table, c_caption)
}

func Fn_atk_table_set_column_description(table unsafe.Pointer, column int, description string) {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_column := (C.gint)(column)

	c_description := (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(c_description))

	C.atk_table_set_column_description(c_table, c_column, c_description)
}

func Fn_atk_table_set_column_header(table unsafe.Pointer, column int, header unsafe.Pointer) {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_column := (C.gint)(column)

	c_header := (*C.AtkObject)(unsafe.Pointer(header))

	C.atk_table_set_column_header(c_table, c_column, c_header)
}

func Fn_atk_table_set_row_description(table unsafe.Pointer, row int, description string) {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	c_description := (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(c_description))

	C.atk_table_set_row_description(c_table, c_row, c_description)
}

func Fn_atk_table_set_row_header(table unsafe.Pointer, row int, header unsafe.Pointer) {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_row := (C.gint)(row)

	c_header := (*C.AtkObject)(unsafe.Pointer(header))

	C.atk_table_set_row_header(c_table, c_row, c_header)
}

func Fn_atk_table_set_summary(table unsafe.Pointer, accessible unsafe.Pointer) {
	c_table := (*C.AtkTable)(unsafe.Pointer(table))

	c_accessible := (*C.AtkObject)(unsafe.Pointer(accessible))

	C.atk_table_set_summary(c_table, c_accessible)
}

// UNSUPPORTED : atk_table_cell_get_column_header_cells : no array length

// UNSUPPORTED : atk_table_cell_get_row_header_cells : no array length

func Fn_atk_text_add_selection(text unsafe.Pointer, startOffset int, endOffset int) bool {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_startOffset := (C.gint)(startOffset)

	c_endOffset := (C.gint)(endOffset)

	ret := C.atk_text_add_selection(c_text, c_startOffset, c_endOffset)

	return toGoBool(ret)
}

// UNSUPPORTED : atk_text_get_bounded_ranges : no array length

func Fn_atk_text_get_caret_offset(text unsafe.Pointer) int {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	ret := C.atk_text_get_caret_offset(c_text)

	return (int)(ret)
}

func Fn_atk_text_get_character_at_offset(text unsafe.Pointer, offset int) rune {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_offset := (C.gint)(offset)

	ret := C.atk_text_get_character_at_offset(c_text, c_offset)

	return (rune)(ret)
}

func Fn_atk_text_get_character_count(text unsafe.Pointer) int {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	ret := C.atk_text_get_character_count(c_text)

	return (int)(ret)
}

func Fn_atk_text_get_character_extents(text unsafe.Pointer, offset int, x *int, y *int, width *int, height *int, coords int) {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_offset := (C.gint)(offset)

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	c_coords := (C.AtkCoordType)(coords)

	C.atk_text_get_character_extents(c_text, c_offset, c_x, c_y, c_width, c_height, c_coords)
}

// UNSUPPORTED : atk_text_get_default_attributes : blacklisted

func Fn_atk_text_get_n_selections(text unsafe.Pointer) int {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	ret := C.atk_text_get_n_selections(c_text)

	return (int)(ret)
}

func Fn_atk_text_get_offset_at_point(text unsafe.Pointer, x int, y int, coords int) int {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coords := (C.AtkCoordType)(coords)

	ret := C.atk_text_get_offset_at_point(c_text, c_x, c_y, c_coords)

	return (int)(ret)
}

func Fn_atk_text_get_range_extents(text unsafe.Pointer, startOffset int, endOffset int, coordType int, rect unsafe.Pointer) {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_startOffset := (C.gint)(startOffset)

	c_endOffset := (C.gint)(endOffset)

	c_coordType := (C.AtkCoordType)(coordType)

	c_rect := (*C.AtkTextRectangle)(unsafe.Pointer(rect))

	C.atk_text_get_range_extents(c_text, c_startOffset, c_endOffset, c_coordType, c_rect)
}

// UNSUPPORTED : atk_text_get_run_attributes : blacklisted

func Fn_atk_text_get_selection(text unsafe.Pointer, selectionNum int, startOffset *int, endOffset *int) string {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_selectionNum := (C.gint)(selectionNum)

	c_startOffset := (*C.gint)(unsafe.Pointer(startOffset))

	c_endOffset := (*C.gint)(unsafe.Pointer(endOffset))

	ret := C.atk_text_get_selection(c_text, c_selectionNum, c_startOffset, c_endOffset)

	return C.GoString(ret)
}

func Fn_atk_text_get_string_at_offset(text unsafe.Pointer, offset int, granularity int, startOffset *int, endOffset *int) string {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_offset := (C.gint)(offset)

	c_granularity := (C.AtkTextGranularity)(granularity)

	c_startOffset := (*C.gint)(unsafe.Pointer(startOffset))

	c_endOffset := (*C.gint)(unsafe.Pointer(endOffset))

	ret := C.atk_text_get_string_at_offset(c_text, c_offset, c_granularity, c_startOffset, c_endOffset)

	return C.GoString(ret)
}

func Fn_atk_text_get_text(text unsafe.Pointer, startOffset int, endOffset int) string {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_startOffset := (C.gint)(startOffset)

	c_endOffset := (C.gint)(endOffset)

	ret := C.atk_text_get_text(c_text, c_startOffset, c_endOffset)

	return C.GoString(ret)
}

func Fn_atk_text_get_text_after_offset(text unsafe.Pointer, offset int, boundaryType int, startOffset *int, endOffset *int) string {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_offset := (C.gint)(offset)

	c_boundaryType := (C.AtkTextBoundary)(boundaryType)

	c_startOffset := (*C.gint)(unsafe.Pointer(startOffset))

	c_endOffset := (*C.gint)(unsafe.Pointer(endOffset))

	ret := C.atk_text_get_text_after_offset(c_text, c_offset, c_boundaryType, c_startOffset, c_endOffset)

	return C.GoString(ret)
}

func Fn_atk_text_get_text_at_offset(text unsafe.Pointer, offset int, boundaryType int, startOffset *int, endOffset *int) string {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_offset := (C.gint)(offset)

	c_boundaryType := (C.AtkTextBoundary)(boundaryType)

	c_startOffset := (*C.gint)(unsafe.Pointer(startOffset))

	c_endOffset := (*C.gint)(unsafe.Pointer(endOffset))

	ret := C.atk_text_get_text_at_offset(c_text, c_offset, c_boundaryType, c_startOffset, c_endOffset)

	return C.GoString(ret)
}

func Fn_atk_text_get_text_before_offset(text unsafe.Pointer, offset int, boundaryType int, startOffset *int, endOffset *int) string {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_offset := (C.gint)(offset)

	c_boundaryType := (C.AtkTextBoundary)(boundaryType)

	c_startOffset := (*C.gint)(unsafe.Pointer(startOffset))

	c_endOffset := (*C.gint)(unsafe.Pointer(endOffset))

	ret := C.atk_text_get_text_before_offset(c_text, c_offset, c_boundaryType, c_startOffset, c_endOffset)

	return C.GoString(ret)
}

func Fn_atk_text_remove_selection(text unsafe.Pointer, selectionNum int) bool {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_selectionNum := (C.gint)(selectionNum)

	ret := C.atk_text_remove_selection(c_text, c_selectionNum)

	return toGoBool(ret)
}

func Fn_atk_text_set_caret_offset(text unsafe.Pointer, offset int) bool {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_offset := (C.gint)(offset)

	ret := C.atk_text_set_caret_offset(c_text, c_offset)

	return toGoBool(ret)
}

func Fn_atk_text_set_selection(text unsafe.Pointer, selectionNum int, startOffset int, endOffset int) bool {
	c_text := (*C.AtkText)(unsafe.Pointer(text))

	c_selectionNum := (C.gint)(selectionNum)

	c_startOffset := (C.gint)(startOffset)

	c_endOffset := (C.gint)(endOffset)

	ret := C.atk_text_set_selection(c_text, c_selectionNum, c_startOffset, c_endOffset)

	return toGoBool(ret)
}

func Fn_atk_value_get_current_value(obj unsafe.Pointer, value unsafe.Pointer) {
	c_obj := (*C.AtkValue)(unsafe.Pointer(obj))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.atk_value_get_current_value(c_obj, c_value)
}

func Fn_atk_value_get_maximum_value(obj unsafe.Pointer, value unsafe.Pointer) {
	c_obj := (*C.AtkValue)(unsafe.Pointer(obj))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.atk_value_get_maximum_value(c_obj, c_value)
}

func Fn_atk_value_get_minimum_increment(obj unsafe.Pointer, value unsafe.Pointer) {
	c_obj := (*C.AtkValue)(unsafe.Pointer(obj))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.atk_value_get_minimum_increment(c_obj, c_value)
}

func Fn_atk_value_get_minimum_value(obj unsafe.Pointer, value unsafe.Pointer) {
	c_obj := (*C.AtkValue)(unsafe.Pointer(obj))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.atk_value_get_minimum_value(c_obj, c_value)
}

// UNSUPPORTED : atk_value_get_value_and_text : parameter 'text' is non array with indirect count > 1

func Fn_atk_value_set_current_value(obj unsafe.Pointer, value unsafe.Pointer) bool {
	c_obj := (*C.AtkValue)(unsafe.Pointer(obj))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	ret := C.atk_value_set_current_value(c_obj, c_value)

	return toGoBool(ret)
}
