// This is a generated file - DO NOT EDIT
// +build atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

import gobject "github.com/pekim/gobbi/lib/gobject"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

func (recv *GObjectAccessible) Object() *Object {}

func (recv *Hyperlink) Object() *gobject.Object {}

func (recv *Misc) Object() *gobject.Object {}

func (recv *NoOpObject) Object() *Object {}

func (recv *NoOpObjectFactory) ObjectFactory() *ObjectFactory {}

func (recv *Object) Object() *gobject.Object {}

func (recv *ObjectFactory) Object() *gobject.Object {}

func (recv *Plug) Object() *Object {}

func (recv *Registry) Object() *gobject.Object {}

// Unsupported : atk_relation_new : unsupported parameter targets : no param type

// AddTarget is a wrapper around the C function atk_relation_add_target.
func (recv *Relation) AddTarget(target *Object) {
	c_target := (*C.AtkObject)(target.ToC())

	C.atk_relation_add_target((*C.AtkRelation)(recv.native), c_target)

	return
}

func (recv *Relation) Object() *gobject.Object {}

// AddRelationByType is a wrapper around the C function atk_relation_set_add_relation_by_type.
func (recv *RelationSet) AddRelationByType(relationship RelationType, target *Object) {
	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(target.ToC())

	C.atk_relation_set_add_relation_by_type((*C.AtkRelationSet)(recv.native), c_relationship, c_target)

	return
}

func (recv *RelationSet) Object() *gobject.Object {}

func (recv *Socket) Object() *Object {}

func (recv *StateSet) Object() *gobject.Object {}

func (recv *Util) Object() *gobject.Object {}
