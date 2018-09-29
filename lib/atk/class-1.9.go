// This is a generated file - DO NOT EDIT
// +build atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Unsupported : atk_object_connect_property_change_handler : unsupported parameter handler : no type generator for PropertyChangeHandler, AtkPropertyChangeHandler*

// Unsupported : atk_object_factory_get_accessible_type : no return generator

// Unsupported : atk_registry_get_factory : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_registry_get_factory_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_registry_set_factory_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_relation_new : unsupported parameter targets : no param type

// AddTarget is a wrapper around the C function atk_relation_add_target.
func (recv *Relation) AddTarget(target *Object) {
	c_target := (*C.AtkObject)(target.ToC())

	C.atk_relation_add_target((*C.AtkRelation)(recv.native), c_target)

	return
}

// Unsupported : atk_relation_get_target : no return type

// AddRelationByType is a wrapper around the C function atk_relation_set_add_relation_by_type.
func (recv *RelationSet) AddRelationByType(relationship RelationType, target *Object) {
	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(target.ToC())

	C.atk_relation_set_add_relation_by_type((*C.AtkRelationSet)(recv.native), c_relationship, c_target)

	return
}

// Unsupported : atk_state_set_add_states : unsupported parameter types : no param type

// Unsupported : atk_state_set_contains_states : unsupported parameter types : no param type
