// This is a generated file - DO NOT EDIT
// +build atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Adds the specified AtkObject to the target for the relation, if it is
// not already present.  See also atk_object_add_relationship().
/*

C function

atk_relation_add_target
*/
func (recv *Relation) AddTarget(target *Object) {
	c_target := (*C.AtkObject)(C.NULL)
	if target != nil {
		c_target = (*C.AtkObject)(target.ToC())
	}

	C.atk_relation_add_target((*C.AtkRelation)(recv.native), c_target)

	return
}

// Add a new relation of the specified type with the specified target to
// the current relation set if the relation set does not contain a relation
// of that type. If it is does contain a relation of that typea the target
// is added to the relation.
/*

C function

atk_relation_set_add_relation_by_type
*/
func (recv *RelationSet) AddRelationByType(relationship RelationType, target *Object) {
	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(C.NULL)
	if target != nil {
		c_target = (*C.AtkObject)(target.ToC())
	}

	C.atk_relation_set_add_relation_by_type((*C.AtkRelationSet)(recv.native), c_relationship, c_target)

	return
}
