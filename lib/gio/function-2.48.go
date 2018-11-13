// This is a generated file - DO NOT EDIT
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
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
// #include <stdlib.h>
import "C"

// Creates a new #GDtlsClientConnection wrapping @base_socket which is
// assumed to communicate with the server identified by @server_identity.
/*

C function : g_dtls_client_connection_new
*/
func DtlsClientConnectionNew(baseSocket *DatagramBased, serverIdentity *SocketConnectable) (*DtlsClientConnection, error) {
	c_base_socket := (*C.GDatagramBased)(baseSocket.ToC())

	c_server_identity := (*C.GSocketConnectable)(serverIdentity.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_client_connection_new(c_base_socket, c_server_identity, &cThrowableError)
	retGo := DtlsClientConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a new #GDtlsServerConnection wrapping @base_socket.
/*

C function : g_dtls_server_connection_new
*/
func DtlsServerConnectionNew(baseSocket *DatagramBased, certificate *TlsCertificate) (*DtlsServerConnection, error) {
	c_base_socket := (*C.GDatagramBased)(baseSocket.ToC())

	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dtls_server_connection_new(c_base_socket, c_certificate, &cThrowableError)
	retGo := DtlsServerConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
