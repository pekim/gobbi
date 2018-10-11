// This is a generated file - DO NOT EDIT
// +build atk_2.8 atk_2.12

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

// GetObjectLocale is a wrapper around the C function atk_object_get_object_locale.
func (recv *Object) GetObjectLocale() string {
	retC := C.atk_object_get_object_locale((*C.AtkObject)(recv.native))
	retGo := C.GoString(retC)

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
