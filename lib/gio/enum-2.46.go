// This is a generated file - DO NOT EDIT
// +build gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "C"

type SocketListenerEvent C.GSocketListenerEvent

const (
	SOCKET_LISTENER_BINDING   SocketListenerEvent = 0
	SOCKET_LISTENER_BOUND     SocketListenerEvent = 1
	SOCKET_LISTENER_LISTENING SocketListenerEvent = 2
	SOCKET_LISTENER_LISTENED  SocketListenerEvent = 3
)
