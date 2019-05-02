// This is a generated file - DO NOT EDIT
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

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

// Equals compares this DatagramBasedInterface with another DatagramBasedInterface, and returns true if they represent the same GObject.
func (recv *DatagramBasedInterface) Equals(other *DatagramBasedInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DtlsClientConnectionInterface with another DtlsClientConnectionInterface, and returns true if they represent the same GObject.
func (recv *DtlsClientConnectionInterface) Equals(other *DtlsClientConnectionInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DtlsConnectionInterface with another DtlsConnectionInterface, and returns true if they represent the same GObject.
func (recv *DtlsConnectionInterface) Equals(other *DtlsConnectionInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DtlsServerConnectionInterface with another DtlsServerConnectionInterface, and returns true if they represent the same GObject.
func (recv *DtlsServerConnectionInterface) Equals(other *DtlsServerConnectionInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this InputMessage with another InputMessage, and returns true if they represent the same GObject.
func (recv *InputMessage) Equals(other *InputMessage) bool {
	return other.ToC() == recv.ToC()
}
