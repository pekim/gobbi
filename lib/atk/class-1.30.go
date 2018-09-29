// This is a generated file - DO NOT EDIT
// +build atk_1.30 atk_2.8 atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Unsupported : atk_object_connect_property_change_handler : unsupported parameter handler : no type generator for PropertyChangeHandler, AtkPropertyChangeHandler*

// Unsupported : atk_object_factory_get_accessible_type : no return generator

// GetId is a wrapper around the C function atk_plug_get_id.
func (recv *Plug) GetId() string {
	retC := C.atk_plug_get_id((*C.AtkPlug)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : atk_registry_get_factory : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_registry_get_factory_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_registry_set_factory_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_relation_new : unsupported parameter targets : no param type

// Unsupported : atk_relation_get_target : no return type

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

// Unsupported : atk_state_set_add_states : unsupported parameter types : no param type

// Unsupported : atk_state_set_contains_states : unsupported parameter types : no param type
