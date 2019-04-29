// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// ListModelInterface is a wrapper around the C record GListModelInterface.
type ListModelInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for get_item_type
	// no type for get_n_items
	// no type for get_item
}

// OutputMessage is a wrapper around the C record GOutputMessage.
type OutputMessage struct {
	native unsafe.Pointer
	// address : record
	// vectors : record
	NumVectors uint32
	BytesSent  uint32
	// no type for control_messages
	NumControlMessages uint32
}
