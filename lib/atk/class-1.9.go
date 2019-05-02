// This is a generated file - DO NOT EDIT
// +build atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

import "C"

// AddTarget is a wrapper around the C function atk_relation_add_target.
func (recv *Relation) AddTarget(target *Object) {
	c_target := (*C.AtkObject)(C.NULL)
	if target != nil {
		c_target = (*C.AtkObject)(target.ToC())
	}

	C.atk_relation_add_target((*C.AtkRelation)(recv.native), c_target)

	return
}

// AddRelationByType is a wrapper around the C function atk_relation_set_add_relation_by_type.
func (recv *RelationSet) AddRelationByType(relationship RelationType, target *Object) {
	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(C.NULL)
	if target != nil {
		c_target = (*C.AtkObject)(target.ToC())
	}

	C.atk_relation_set_add_relation_by_type((*C.AtkRelationSet)(recv.native), c_relationship, c_target)

	return
}
