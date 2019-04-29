// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// Credentials is a wrapper around the C record GCredentials.
type Credentials struct {
	native unsafe.Pointer
}

// DBusAuthObserver is a wrapper around the C record GDBusAuthObserver.
type DBusAuthObserver struct {
	native unsafe.Pointer
}

// DBusConnection is a wrapper around the C record GDBusConnection.
type DBusConnection struct {
	native unsafe.Pointer
}

// DBusMessage is a wrapper around the C record GDBusMessage.
type DBusMessage struct {
	native unsafe.Pointer
}

// DBusMethodInvocation is a wrapper around the C record GDBusMethodInvocation.
type DBusMethodInvocation struct {
	native unsafe.Pointer
}

// DBusProxy is a wrapper around the C record GDBusProxy.
type DBusProxy struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// DBusServer is a wrapper around the C record GDBusServer.
type DBusServer struct {
	native unsafe.Pointer
}

// ProxyAddress is a wrapper around the C record GProxyAddress.
type ProxyAddress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// UnixCredentialsMessage is a wrapper around the C record GUnixCredentialsMessage.
type UnixCredentialsMessage struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}
