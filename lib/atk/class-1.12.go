// This is a generated file - DO NOT EDIT
// +build atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

func (recv *GObjectAccessible) Object() *Object {}

func (recv *Hyperlink) Object() *gobject.Object {}

func (recv *Misc) Object() *gobject.Object {}

func (recv *NoOpObject) Object() *Object {}

func (recv *NoOpObjectFactory) ObjectFactory() *ObjectFactory {}

// GetAttributes is a wrapper around the C function atk_object_get_attributes.
func (recv *Object) GetAttributes() *AttributeSet {
	retC := C.atk_object_get_attributes((*C.AtkObject)(recv.native))
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Object) Object() *gobject.Object {}

func (recv *ObjectFactory) Object() *gobject.Object {}

func (recv *Plug) Object() *Object {}

func (recv *Registry) Object() *gobject.Object {}

// Unsupported : atk_relation_new : unsupported parameter targets : no param type

func (recv *Relation) Object() *gobject.Object {}

func (recv *RelationSet) Object() *gobject.Object {}

func (recv *Socket) Object() *Object {}

func (recv *StateSet) Object() *gobject.Object {}

func (recv *Util) Object() *gobject.Object {}
