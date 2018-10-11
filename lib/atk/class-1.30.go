// This is a generated file - DO NOT EDIT
// +build atk_1.30 atk_2.8 atk_2.12

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

func (recv *Object) Object() *gobject.Object {}

func (recv *ObjectFactory) Object() *gobject.Object {}

// GetId is a wrapper around the C function atk_plug_get_id.
func (recv *Plug) GetId() string {
	retC := C.atk_plug_get_id((*C.AtkPlug)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

func (recv *Plug) Object() *Object {}

func (recv *Registry) Object() *gobject.Object {}

// Unsupported : atk_relation_new : unsupported parameter targets : no param type

func (recv *Relation) Object() *gobject.Object {}

func (recv *RelationSet) Object() *gobject.Object {}

// Embed is a wrapper around the C function atk_socket_embed.
func (recv *Socket) Embed(plugId string) {
	c_plug_id := C.CString(plugId)
	defer C.free(unsafe.Pointer(c_plug_id))

	C.atk_socket_embed((*C.AtkSocket)(recv.native), c_plug_id)

	return
}

// IsOccupied is a wrapper around the C function atk_socket_is_occupied.
func (recv *Socket) IsOccupied() bool {
	retC := C.atk_socket_is_occupied((*C.AtkSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

func (recv *Socket) Object() *Object {}

func (recv *StateSet) Object() *gobject.Object {}

func (recv *Util) Object() *gobject.Object {}
