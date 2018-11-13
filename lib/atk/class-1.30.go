// This is a generated file - DO NOT EDIT
// +build atk_1.30 atk_2.8 atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Gets the unique ID of an #AtkPlug object, which can be used to
// embed inside of an #AtkSocket using atk_socket_embed().
//
// Internally, this calls a class function that should be registered
// by the IPC layer (usually at-spi2-atk). The implementor of an
// #AtkPlug object should call this function (after atk-bridge is
// loaded) and pass the value to the process implementing the
// #AtkSocket, so it could embed the plug.
/*

C function

atk_plug_get_id
*/
func (recv *Plug) GetId() string {
	retC := C.atk_plug_get_id((*C.AtkPlug)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Embeds the children of an #AtkPlug as the children of the
// #AtkSocket. The plug may be in the same process or in a different
// process.
//
// The class item used by this function should be filled in by the IPC
// layer (usually at-spi2-atk). The implementor of the AtkSocket
// should call this function and pass the id for the plug as returned
// by atk_plug_get_id().  It is the responsibility of the application
// to pass the plug id on to the process implementing the #AtkSocket
// as needed.
/*

C function

atk_socket_embed
*/
func (recv *Socket) Embed(plugId string) {
	c_plug_id := C.CString(plugId)
	defer C.free(unsafe.Pointer(c_plug_id))

	C.atk_socket_embed((*C.AtkSocket)(recv.native), c_plug_id)

	return
}

// Determines whether or not the socket has an embedded plug.
/*

C function

atk_socket_is_occupied
*/
func (recv *Socket) IsOccupied() bool {
	retC := C.atk_socket_is_occupied((*C.AtkSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
