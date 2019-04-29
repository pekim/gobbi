// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// Credentials is a wrapper around the C record GCredentials.
type Credentials struct {
	native unsafe.Pointer
}

func CredentialsNewFromC(u unsafe.Pointer) *Credentials {
	if u == nil {
		return nil
	}

	g := &Credentials{native: u}

	return g
}

func (recv *Credentials) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusAuthObserver is a wrapper around the C record GDBusAuthObserver.
type DBusAuthObserver struct {
	native unsafe.Pointer
}

func DBusAuthObserverNewFromC(u unsafe.Pointer) *DBusAuthObserver {
	if u == nil {
		return nil
	}

	g := &DBusAuthObserver{native: u}

	return g
}

func (recv *DBusAuthObserver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusConnection is a wrapper around the C record GDBusConnection.
type DBusConnection struct {
	native unsafe.Pointer
}

func DBusConnectionNewFromC(u unsafe.Pointer) *DBusConnection {
	if u == nil {
		return nil
	}

	g := &DBusConnection{native: u}

	return g
}

func (recv *DBusConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusMessage is a wrapper around the C record GDBusMessage.
type DBusMessage struct {
	native unsafe.Pointer
}

func DBusMessageNewFromC(u unsafe.Pointer) *DBusMessage {
	if u == nil {
		return nil
	}

	g := &DBusMessage{native: u}

	return g
}

func (recv *DBusMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusMethodInvocation is a wrapper around the C record GDBusMethodInvocation.
type DBusMethodInvocation struct {
	native unsafe.Pointer
}

func DBusMethodInvocationNewFromC(u unsafe.Pointer) *DBusMethodInvocation {
	if u == nil {
		return nil
	}

	g := &DBusMethodInvocation{native: u}

	return g
}

func (recv *DBusMethodInvocation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusProxy is a wrapper around the C record GDBusProxy.
type DBusProxy struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func DBusProxyNewFromC(u unsafe.Pointer) *DBusProxy {
	if u == nil {
		return nil
	}

	g := &DBusProxy{native: u}

	return g
}

func (recv *DBusProxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusServer is a wrapper around the C record GDBusServer.
type DBusServer struct {
	native unsafe.Pointer
}

func DBusServerNewFromC(u unsafe.Pointer) *DBusServer {
	if u == nil {
		return nil
	}

	g := &DBusServer{native: u}

	return g
}

func (recv *DBusServer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddress is a wrapper around the C record GProxyAddress.
type ProxyAddress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func ProxyAddressNewFromC(u unsafe.Pointer) *ProxyAddress {
	if u == nil {
		return nil
	}

	g := &ProxyAddress{native: u}

	return g
}

func (recv *ProxyAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixCredentialsMessage is a wrapper around the C record GUnixCredentialsMessage.
type UnixCredentialsMessage struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func UnixCredentialsMessageNewFromC(u unsafe.Pointer) *UnixCredentialsMessage {
	if u == nil {
		return nil
	}

	g := &UnixCredentialsMessage{native: u}

	return g
}

func (recv *UnixCredentialsMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
