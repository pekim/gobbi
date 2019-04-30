// This is a generated file - DO NOT EDIT
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// DatagramBased is a wrapper around the C record GDatagramBased.
type DatagramBased struct {
	native unsafe.Pointer
}

func DatagramBasedNewFromC(u unsafe.Pointer) *DatagramBased {
	if u == nil {
		return nil
	}

	g := &DatagramBased{native: u}

	return g
}

func (recv *DatagramBased) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsClientConnection is a wrapper around the C record GDtlsClientConnection.
type DtlsClientConnection struct {
	native unsafe.Pointer
}

func DtlsClientConnectionNewFromC(u unsafe.Pointer) *DtlsClientConnection {
	if u == nil {
		return nil
	}

	g := &DtlsClientConnection{native: u}

	return g
}

func (recv *DtlsClientConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsConnection is a wrapper around the C record GDtlsConnection.
type DtlsConnection struct {
	native unsafe.Pointer
}

func DtlsConnectionNewFromC(u unsafe.Pointer) *DtlsConnection {
	if u == nil {
		return nil
	}

	g := &DtlsConnection{native: u}

	return g
}

func (recv *DtlsConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsServerConnection is a wrapper around the C record GDtlsServerConnection.
type DtlsServerConnection struct {
	native unsafe.Pointer
}

func DtlsServerConnectionNewFromC(u unsafe.Pointer) *DtlsServerConnection {
	if u == nil {
		return nil
	}

	g := &DtlsServerConnection{native: u}

	return g
}

func (recv *DtlsServerConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
