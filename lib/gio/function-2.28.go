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

// AppInfoGetFallbackForType is a wrapper around the C function g_app_info_get_fallback_for_type.
func AppInfoGetFallbackForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_fallback_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppInfoGetRecommendedForType is a wrapper around the C function g_app_info_get_recommended_for_type.
func AppInfoGetRecommendedForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_recommended_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_memory_settings_backend_new

// Blacklisted : g_null_settings_backend_new

// PollableSourceNew is a wrapper around the C function g_pollable_source_new.
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

// TlsBackendGetDefault is a wrapper around the C function g_tls_backend_get_default.
func TlsBackendGetDefault() *TlsBackend {
	retC := C.g_tls_backend_get_default()
	retGo := TlsBackendNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TlsClientConnectionNew is a wrapper around the C function g_tls_client_connection_new.
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

// TlsErrorQuark is a wrapper around the C function g_tls_error_quark.
func TlsErrorQuark() glib.Quark {
	retC := C.g_tls_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// TlsServerConnectionNew is a wrapper around the C function g_tls_server_connection_new.
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
