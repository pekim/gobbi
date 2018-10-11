// This is a generated file - DO NOT EDIT
// +build atk_1.4 atk_1.6 atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

import gobject "github.com/pekim/gobbi/lib/gobject"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

func (recv *GObjectAccessible) Object() *Object {}

// IsSelectedLink is a wrapper around the C function atk_hyperlink_is_selected_link.
func (recv *Hyperlink) IsSelectedLink() bool {
	retC := C.atk_hyperlink_is_selected_link((*C.AtkHyperlink)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

func (recv *Hyperlink) Object() *gobject.Object {}

func (recv *Misc) Object() *gobject.Object {}

func (recv *NoOpObject) Object() *Object {}

func (recv *NoOpObjectFactory) ObjectFactory() *ObjectFactory {}

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
