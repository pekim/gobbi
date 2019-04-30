// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// DBusInterface is a wrapper around the C record GDBusInterface.
type DBusInterface struct {
	native unsafe.Pointer
}

func DBusInterfaceNewFromC(u unsafe.Pointer) *DBusInterface {
	if u == nil {
		return nil
	}

	g := &DBusInterface{native: u}

	return g
}

func (recv *DBusInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsFileDatabase is a wrapper around the C record GTlsFileDatabase.
type TlsFileDatabase struct {
	native unsafe.Pointer
}

func TlsFileDatabaseNewFromC(u unsafe.Pointer) *TlsFileDatabase {
	if u == nil {
		return nil
	}

	g := &TlsFileDatabase{native: u}

	return g
}

func (recv *TlsFileDatabase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
