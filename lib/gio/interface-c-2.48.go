// This is a generated file - DO NOT EDIT
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <gio/gnetworking.h>
// #include <stdlib.h>
import "C"

// DatagramBased is a wrapper around the C record GDatagramBased.
type DatagramBased struct {
	native *C.GDatagramBased
}

func DatagramBasedNewFromC(u unsafe.Pointer) *DatagramBased {
	c := (*C.GDatagramBased)(u)
	if c == nil {
		return nil
	}

	g := &DatagramBased{native: c}

	return g
}

func (recv *DatagramBased) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsClientConnection is a wrapper around the C record GDtlsClientConnection.
type DtlsClientConnection struct {
	native *C.GDtlsClientConnection
}

func DtlsClientConnectionNewFromC(u unsafe.Pointer) *DtlsClientConnection {
	c := (*C.GDtlsClientConnection)(u)
	if c == nil {
		return nil
	}

	g := &DtlsClientConnection{native: c}

	return g
}

func (recv *DtlsClientConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsConnection is a wrapper around the C record GDtlsConnection.
type DtlsConnection struct {
	native *C.GDtlsConnection
}

func DtlsConnectionNewFromC(u unsafe.Pointer) *DtlsConnection {
	c := (*C.GDtlsConnection)(u)
	if c == nil {
		return nil
	}

	g := &DtlsConnection{native: c}

	return g
}

func (recv *DtlsConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsServerConnection is a wrapper around the C record GDtlsServerConnection.
type DtlsServerConnection struct {
	native *C.GDtlsServerConnection
}

func DtlsServerConnectionNewFromC(u unsafe.Pointer) *DtlsServerConnection {
	c := (*C.GDtlsServerConnection)(u)
	if c == nil {
		return nil
	}

	g := &DtlsServerConnection{native: c}

	return g
}

func (recv *DtlsServerConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
