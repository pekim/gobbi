// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// Unsupported : g_list_store_new : return type :

// Unsupported : g_network_address_new_loopback : return type :

// SimpleIOStream is a wrapper around the C record GSimpleIOStream.
type SimpleIOStream struct {
	native unsafe.Pointer
}

func SimpleIOStreamNewFromC(u unsafe.Pointer) *SimpleIOStream {
	if u == nil {
		return nil
	}

	g := &SimpleIOStream{native: u}

	return g
}

func (recv *SimpleIOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_simple_io_stream_new : return type :
