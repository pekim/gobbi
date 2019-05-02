// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"runtime"
	"unsafe"
)

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
// #include <stdlib.h>
import "C"

// Application is a wrapper around the C record GApplication.
type Application struct {
	native *C.GApplication
	// Private : parent_instance
	// Private : priv
}

func ApplicationNewFromC(u unsafe.Pointer) *Application {
	c := (*C.GApplication)(u)
	if c == nil {
		return nil
	}

	g := &Application{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Application) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Application) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationNew is a wrapper around the C function g_application_new.
func ApplicationNew(applicationId string, flags ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.g_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SimpleActionNew is a wrapper around the C function g_simple_action_new.
func SimpleActionNew(name string, parameterType *glib.VariantType) *SimpleAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_parameter_type := (*C.GVariantType)(C.NULL)
	if parameterType != nil {
		c_parameter_type = (*C.GVariantType)(parameterType.ToC())
	}

	retC := C.g_simple_action_new(c_name, c_parameter_type)
	retGo := SimpleActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SimpleActionNewStateful is a wrapper around the C function g_simple_action_new_stateful.
func SimpleActionNewStateful(name string, parameterType *glib.VariantType, state *glib.Variant) *SimpleAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_parameter_type := (*C.GVariantType)(C.NULL)
	if parameterType != nil {
		c_parameter_type = (*C.GVariantType)(parameterType.ToC())
	}

	c_state := (*C.GVariant)(C.NULL)
	if state != nil {
		c_state = (*C.GVariant)(state.ToC())
	}

	retC := C.g_simple_action_new_stateful(c_name, c_parameter_type, c_state)
	retGo := SimpleActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SimpleActionGroup is a wrapper around the C record GSimpleActionGroup.
type SimpleActionGroup struct {
	native *C.GSimpleActionGroup
	// Private : parent_instance
	// Private : priv
}

func SimpleActionGroupNewFromC(u unsafe.Pointer) *SimpleActionGroup {
	c := (*C.GSimpleActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &SimpleActionGroup{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleActionGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleActionGroupNew is a wrapper around the C function g_simple_action_group_new.
func SimpleActionGroupNew() *SimpleActionGroup {
	retC := C.g_simple_action_group_new()
	retGo := SimpleActionGroupNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// TcpWrapperConnectionNew is a wrapper around the C function g_tcp_wrapper_connection_new.
func TcpWrapperConnectionNew(baseIoStream *IOStream, socket *Socket) *TcpWrapperConnection {
	c_base_io_stream := (*C.GIOStream)(C.NULL)
	if baseIoStream != nil {
		c_base_io_stream = (*C.GIOStream)(baseIoStream.ToC())
	}

	c_socket := (*C.GSocket)(C.NULL)
	if socket != nil {
		c_socket = (*C.GSocket)(socket.ToC())
	}

	retC := C.g_tcp_wrapper_connection_new(c_base_io_stream, c_socket)
	retGo := TcpWrapperConnectionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// TlsCertificate is a wrapper around the C record GTlsCertificate.
type TlsCertificate struct {
	native *C.GTlsCertificate
	// parent_instance : record
	// priv : record
}

func TlsCertificateNewFromC(u unsafe.Pointer) *TlsCertificate {
	c := (*C.GTlsCertificate)(u)
	if c == nil {
		return nil
	}

	g := &TlsCertificate{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TlsCertificate) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TlsCertificate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsCertificateNewFromFile is a wrapper around the C function g_tls_certificate_new_from_file.
func TlsCertificateNewFromFile(file string) (*TlsCertificate, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_file(c_file, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TlsCertificateNewFromFiles is a wrapper around the C function g_tls_certificate_new_from_files.
func TlsCertificateNewFromFiles(certFile string, keyFile string) (*TlsCertificate, error) {
	c_cert_file := C.CString(certFile)
	defer C.free(unsafe.Pointer(c_cert_file))

	c_key_file := C.CString(keyFile)
	defer C.free(unsafe.Pointer(c_key_file))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_files(c_cert_file, c_key_file, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TlsCertificateNewFromPem is a wrapper around the C function g_tls_certificate_new_from_pem.
func TlsCertificateNewFromPem(data string) (*TlsCertificate, error) {
	c_data := C.CString(data)
	defer C.free(unsafe.Pointer(c_data))

	c_length := (C.gssize)(len(data))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_pem(c_data, c_length, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TlsConnection is a wrapper around the C record GTlsConnection.
type TlsConnection struct {
	native *C.GTlsConnection
	// parent_instance : record
	// priv : record
}

func TlsConnectionNewFromC(u unsafe.Pointer) *TlsConnection {
	c := (*C.GTlsConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TlsConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TlsConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : g_memory_settings_backend_new

// Blacklisted : g_null_settings_backend_new

// Unsupported : g_simple_async_report_take_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// PollableInputStream is a wrapper around the C record GPollableInputStream.
type PollableInputStream struct {
	native *C.GPollableInputStream
}

func PollableInputStreamNewFromC(u unsafe.Pointer) *PollableInputStream {
	c := (*C.GPollableInputStream)(u)
	if c == nil {
		return nil
	}

	g := &PollableInputStream{native: c}

	return g
}

func (recv *PollableInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PollableOutputStream is a wrapper around the C record GPollableOutputStream.
type PollableOutputStream struct {
	native *C.GPollableOutputStream
}

func PollableOutputStreamNewFromC(u unsafe.Pointer) *PollableOutputStream {
	c := (*C.GPollableOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &PollableOutputStream{native: c}

	return g
}

func (recv *PollableOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsBackend is a wrapper around the C record GTlsBackend.
type TlsBackend struct {
	native *C.GTlsBackend
}

func TlsBackendNewFromC(u unsafe.Pointer) *TlsBackend {
	c := (*C.GTlsBackend)(u)
	if c == nil {
		return nil
	}

	g := &TlsBackend{native: c}

	return g
}

func (recv *TlsBackend) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsClientConnection is a wrapper around the C record GTlsClientConnection.
type TlsClientConnection struct {
	native *C.GTlsClientConnection
}

func TlsClientConnectionNewFromC(u unsafe.Pointer) *TlsClientConnection {
	c := (*C.GTlsClientConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsClientConnection{native: c}

	return g
}

func (recv *TlsClientConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsServerConnection is a wrapper around the C record GTlsServerConnection.
type TlsServerConnection struct {
	native *C.GTlsServerConnection
}

func TlsServerConnectionNewFromC(u unsafe.Pointer) *TlsServerConnection {
	c := (*C.GTlsServerConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsServerConnection{native: c}

	return g
}

func (recv *TlsServerConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionGroupInterface is a wrapper around the C record GActionGroupInterface.
type ActionGroupInterface struct {
	native *C.GActionGroupInterface
	// g_iface : record
	// no type for has_action
	// no type for list_actions
	// no type for get_action_enabled
	// no type for get_action_parameter_type
	// no type for get_action_state_type
	// no type for get_action_state_hint
	// no type for get_action_state
	// no type for change_action_state
	// no type for activate_action
	// no type for action_added
	// no type for action_removed
	// no type for action_enabled_changed
	// no type for action_state_changed
	// no type for query_action
}

func ActionGroupInterfaceNewFromC(u unsafe.Pointer) *ActionGroupInterface {
	c := (*C.GActionGroupInterface)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroupInterface{native: c}

	return g
}

func (recv *ActionGroupInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionInterface is a wrapper around the C record GActionInterface.
type ActionInterface struct {
	native *C.GActionInterface
	// g_iface : record
	// no type for get_name
	// no type for get_parameter_type
	// no type for get_state_type
	// no type for get_state_hint
	// no type for get_enabled
	// no type for get_state
	// no type for change_state
	// no type for activate
}

func ActionInterfaceNewFromC(u unsafe.Pointer) *ActionInterface {
	c := (*C.GActionInterface)(u)
	if c == nil {
		return nil
	}

	g := &ActionInterface{native: c}

	return g
}

func (recv *ActionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationClass is a wrapper around the C record GApplicationClass.
type ApplicationClass struct {
	native *C.GApplicationClass
	// Private : parent_class
	// no type for startup
	// no type for activate
	// no type for open
	// no type for command_line
	// no type for local_command_line
	// no type for before_emit
	// no type for after_emit
	// no type for add_platform_data
	// no type for quit_mainloop
	// no type for run_mainloop
	// no type for shutdown
	// no type for dbus_register
	// no type for dbus_unregister
	// no type for handle_local_options
	// Private : padding
}

func ApplicationClassNewFromC(u unsafe.Pointer) *ApplicationClass {
	c := (*C.GApplicationClass)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationClass{native: c}

	return g
}

func (recv *ApplicationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationCommandLineClass is a wrapper around the C record GApplicationCommandLineClass.
type ApplicationCommandLineClass struct {
	native *C.GApplicationCommandLineClass
	// Private : parent_class
	// no type for print_literal
	// no type for printerr_literal
	// no type for get_stdin
	// Private : padding
}

func ApplicationCommandLineClassNewFromC(u unsafe.Pointer) *ApplicationCommandLineClass {
	c := (*C.GApplicationCommandLineClass)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationCommandLineClass{native: c}

	return g
}

func (recv *ApplicationCommandLineClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PollableInputStreamInterface is a wrapper around the C record GPollableInputStreamInterface.
type PollableInputStreamInterface struct {
	native *C.GPollableInputStreamInterface
	// g_iface : record
	// no type for can_poll
	// no type for is_readable
	// no type for create_source
	// no type for read_nonblocking
}

func PollableInputStreamInterfaceNewFromC(u unsafe.Pointer) *PollableInputStreamInterface {
	c := (*C.GPollableInputStreamInterface)(u)
	if c == nil {
		return nil
	}

	g := &PollableInputStreamInterface{native: c}

	return g
}

func (recv *PollableInputStreamInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PollableOutputStreamInterface is a wrapper around the C record GPollableOutputStreamInterface.
type PollableOutputStreamInterface struct {
	native *C.GPollableOutputStreamInterface
	// g_iface : record
	// no type for can_poll
	// no type for is_writable
	// no type for create_source
	// no type for write_nonblocking
}

func PollableOutputStreamInterfaceNewFromC(u unsafe.Pointer) *PollableOutputStreamInterface {
	c := (*C.GPollableOutputStreamInterface)(u)
	if c == nil {
		return nil
	}

	g := &PollableOutputStreamInterface{native: c}

	return g
}

func (recv *PollableOutputStreamInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsBackendInterface is a wrapper around the C record GTlsBackendInterface.
type TlsBackendInterface struct {
	native *C.GTlsBackendInterface
	// g_iface : record
	// no type for supports_tls
	// no type for get_certificate_type
	// no type for get_client_connection_type
	// no type for get_server_connection_type
	// no type for get_file_database_type
	// no type for get_default_database
	// no type for supports_dtls
	// no type for get_dtls_client_connection_type
	// no type for get_dtls_server_connection_type
}

func TlsBackendInterfaceNewFromC(u unsafe.Pointer) *TlsBackendInterface {
	c := (*C.GTlsBackendInterface)(u)
	if c == nil {
		return nil
	}

	g := &TlsBackendInterface{native: c}

	return g
}

func (recv *TlsBackendInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
