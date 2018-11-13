// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// Gets a list of fallback #GAppInfos for a given content type, i.e.
// those applications which claim to support the given content type
// by MIME type subclassing and not directly.
/*

C function

g_app_info_get_fallback_for_type
*/
func AppInfoGetFallbackForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_fallback_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a list of recommended #GAppInfos for a given content type, i.e.
// those applications which claim to support the given content type exactly,
// and not by MIME type subclassing.
// Note that the first application of the list is the last used one, i.e.
// the last one for which g_app_info_set_as_last_used_for_type() has been
// called.
/*

C function

g_app_info_get_recommended_for_type
*/
func AppInfoGetRecommendedForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_recommended_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_memory_settings_backend_new

// Blacklisted : g_null_settings_backend_new

// Utility method for #GPollableInputStream and #GPollableOutputStream
// implementations. Creates a new #GSource that expects a callback of
// type #GPollableSourceFunc. The new source does not actually do
// anything on its own; use g_source_add_child_source() to add other
// sources to it to cause it to trigger.
/*

C function

g_pollable_source_new
*/
func PollableSourceNew(pollableStream *gobject.Object) *glib.Source {
	c_pollable_stream := (*C.GObject)(C.NULL)
	if pollableStream != nil {
		c_pollable_stream = (*C.GObject)(pollableStream.ToC())
	}

	retC := C.g_pollable_source_new(c_pollable_stream)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_simple_async_report_take_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Gets the default #GTlsBackend for the system.
/*

C function

g_tls_backend_get_default
*/
func TlsBackendGetDefault() *TlsBackend {
	retC := C.g_tls_backend_get_default()
	retGo := TlsBackendNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GTlsClientConnection wrapping @base_io_stream (which
// must have pollable input and output streams) which is assumed to
// communicate with the server identified by @server_identity.
//
// See the documentation for #GTlsConnection:base-io-stream for restrictions
// on when application code can run operations on the @base_io_stream after
// this function has returned.
/*

C function

g_tls_client_connection_new
*/
func TlsClientConnectionNew(baseIoStream *IOStream, serverIdentity *SocketConnectable) (*TlsClientConnection, error) {
	c_base_io_stream := (*C.GIOStream)(C.NULL)
	if baseIoStream != nil {
		c_base_io_stream = (*C.GIOStream)(baseIoStream.ToC())
	}

	c_server_identity := (*C.GSocketConnectable)(serverIdentity.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_client_connection_new(c_base_io_stream, c_server_identity, &cThrowableError)
	retGo := TlsClientConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the TLS error quark.
/*

C function

g_tls_error_quark
*/
func TlsErrorQuark() glib.Quark {
	retC := C.g_tls_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Creates a new #GTlsServerConnection wrapping @base_io_stream (which
// must have pollable input and output streams).
//
// See the documentation for #GTlsConnection:base-io-stream for restrictions
// on when application code can run operations on the @base_io_stream after
// this function has returned.
/*

C function

g_tls_server_connection_new
*/
func TlsServerConnectionNew(baseIoStream *IOStream, certificate *TlsCertificate) (*TlsServerConnection, error) {
	c_base_io_stream := (*C.GIOStream)(C.NULL)
	if baseIoStream != nil {
		c_base_io_stream = (*C.GIOStream)(baseIoStream.ToC())
	}

	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_server_connection_new(c_base_io_stream, c_certificate, &cThrowableError)
	retGo := TlsServerConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
