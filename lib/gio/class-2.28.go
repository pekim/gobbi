// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

	return g
}

func (recv *Application) toC() *C.GApplication {

	return recv.native
}

// Unsupported : g_application_new : no return generator

// Unsupported : g_application_activate : no return generator

// Unsupported : g_application_add_main_option : unsupported parameter short_name : no type generator for gchar, char

// Unsupported : g_application_add_main_option_entries : unsupported parameter entries : no param type

// Unsupported : g_application_add_option_group : no return generator

// Unsupported : g_application_bind_busy_property : no return generator

// GetApplicationId is a wrapper around the C function g_application_get_application_id.
func (recv *Application) GetApplicationId() string {
	retC := C.g_application_get_application_id((*C.GApplication)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_application_get_dbus_connection : no return generator

// GetFlags is a wrapper around the C function g_application_get_flags.
func (recv *Application) GetFlags() ApplicationFlags {
	retC := C.g_application_get_flags((*C.GApplication)(recv.native))
	retGo := (ApplicationFlags)(retC)

	return retGo
}

// GetInactivityTimeout is a wrapper around the C function g_application_get_inactivity_timeout.
func (recv *Application) GetInactivityTimeout() uint32 {
	retC := C.g_application_get_inactivity_timeout((*C.GApplication)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_application_get_is_busy : no return generator

// Unsupported : g_application_get_is_registered : no return generator

// Unsupported : g_application_get_is_remote : no return generator

// Unsupported : g_application_hold : no return generator

// Unsupported : g_application_mark_busy : no return generator

// Unsupported : g_application_open : unsupported parameter files : no param type

// Unsupported : g_application_quit : no return generator

// Unsupported : g_application_register : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_application_release : no return generator

// Unsupported : g_application_run : unsupported parameter argv : no param type

// Unsupported : g_application_send_notification : unsupported parameter notification : no type generator for Notification, GNotification*

// Unsupported : g_application_set_action_group : unsupported parameter action_group : no type generator for ActionGroup, GActionGroup*

// Unsupported : g_application_set_application_id : no return generator

// Unsupported : g_application_set_default : no return generator

// Unsupported : g_application_set_flags : no return generator

// Unsupported : g_application_set_inactivity_timeout : no return generator

// Unsupported : g_application_set_option_context_description : no return generator

// Unsupported : g_application_set_option_context_parameter_string : no return generator

// Unsupported : g_application_set_option_context_summary : no return generator

// Unsupported : g_application_set_resource_base_path : no return generator

// Unsupported : g_application_unbind_busy_property : no return generator

// Unsupported : g_application_unmark_busy : no return generator

// Unsupported : g_application_withdraw_notification : no return generator

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

	return g
}

func (recv *SimpleActionGroup) toC() *C.GSimpleActionGroup {

	return recv.native
}

// Unsupported : g_simple_action_group_new : no return generator

// Unsupported : g_simple_action_group_add_entries : unsupported parameter entries : no param type

// Unsupported : g_simple_action_group_insert : unsupported parameter action : no type generator for Action, GAction*

// Unsupported : g_simple_action_group_lookup : no return generator

// Unsupported : g_simple_action_group_remove : no return generator

// TlsCertificate is a wrapper around the C record GTlsCertificate.
type TlsCertificate struct {
	native *C.GTlsCertificate
	// parent_instance : no type generator for GObject.Object, GObject
	// priv : record
}

func TlsCertificateNewFromC(u unsafe.Pointer) *TlsCertificate {
	c := (*C.GTlsCertificate)(u)
	if c == nil {
		return nil
	}

	g := &TlsCertificate{native: c}

	return g
}

func (recv *TlsCertificate) toC() *C.GTlsCertificate {

	return recv.native
}

// Unsupported : g_tls_certificate_new_from_file : no return generator

// Unsupported : g_tls_certificate_new_from_files : no return generator

// Unsupported : g_tls_certificate_new_from_pem : no return generator

// Unsupported : g_tls_certificate_get_issuer : no return generator

// Unsupported : g_tls_certificate_is_same : unsupported parameter cert_two : no type generator for TlsCertificate, GTlsCertificate*

// Unsupported : g_tls_certificate_verify : unsupported parameter identity : no type generator for SocketConnectable, GSocketConnectable*

// TlsConnection is a wrapper around the C record GTlsConnection.
type TlsConnection struct {
	native *C.GTlsConnection
	// parent_instance : no type generator for IOStream, GIOStream
	// priv : record
}

func TlsConnectionNewFromC(u unsafe.Pointer) *TlsConnection {
	c := (*C.GTlsConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsConnection{native: c}

	return g
}

func (recv *TlsConnection) toC() *C.GTlsConnection {

	return recv.native
}

// Unsupported : g_tls_connection_emit_accept_certificate : unsupported parameter peer_cert : no type generator for TlsCertificate, GTlsCertificate*

// Unsupported : g_tls_connection_get_certificate : no return generator

// Unsupported : g_tls_connection_get_database : no return generator

// Unsupported : g_tls_connection_get_interaction : no return generator

// Unsupported : g_tls_connection_get_peer_certificate : no return generator

// GetPeerCertificateErrors is a wrapper around the C function g_tls_connection_get_peer_certificate_errors.
func (recv *TlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	retC := C.g_tls_connection_get_peer_certificate_errors((*C.GTlsConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// GetRehandshakeMode is a wrapper around the C function g_tls_connection_get_rehandshake_mode.
func (recv *TlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	retC := C.g_tls_connection_get_rehandshake_mode((*C.GTlsConnection)(recv.native))
	retGo := (TlsRehandshakeMode)(retC)

	return retGo
}

// Unsupported : g_tls_connection_get_require_close_notify : no return generator

// Unsupported : g_tls_connection_get_use_system_certdb : no return generator

// Unsupported : g_tls_connection_handshake : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_tls_connection_handshake_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_tls_connection_handshake_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_connection_set_certificate : unsupported parameter certificate : no type generator for TlsCertificate, GTlsCertificate*

// Unsupported : g_tls_connection_set_database : unsupported parameter database : no type generator for TlsDatabase, GTlsDatabase*

// Unsupported : g_tls_connection_set_interaction : unsupported parameter interaction : no type generator for TlsInteraction, GTlsInteraction*

// Unsupported : g_tls_connection_set_rehandshake_mode : no return generator

// Unsupported : g_tls_connection_set_require_close_notify : unsupported parameter require_close_notify : no type generator for gboolean, gboolean

// Unsupported : g_tls_connection_set_use_system_certdb : unsupported parameter use_system_certdb : no type generator for gboolean, gboolean
