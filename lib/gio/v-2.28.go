// Code generated - DO NOT EDIT.
// +build gio_2.28

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	gio "github.com/pekim/gobbi/lib/internal/c/gio"
	"unsafe"
)

// ApplicationFlags is a representation of the C bitfield GApplicationFlags.
type ApplicationFlags int

// ApplicationFlags_flags_none is a representation of the C bitfield member G_APPLICATION_FLAGS_NONE.
const ApplicationFlags_flags_none = ApplicationFlags(0)

// ApplicationFlags_is_service is a representation of the C bitfield member G_APPLICATION_IS_SERVICE.
const ApplicationFlags_is_service = ApplicationFlags(1)

// ApplicationFlags_is_launcher is a representation of the C bitfield member G_APPLICATION_IS_LAUNCHER.
const ApplicationFlags_is_launcher = ApplicationFlags(2)

// ApplicationFlags_handles_open is a representation of the C bitfield member G_APPLICATION_HANDLES_OPEN.
const ApplicationFlags_handles_open = ApplicationFlags(4)

// ApplicationFlags_handles_command_line is a representation of the C bitfield member G_APPLICATION_HANDLES_COMMAND_LINE.
const ApplicationFlags_handles_command_line = ApplicationFlags(8)

// ApplicationFlags_send_environment is a representation of the C bitfield member G_APPLICATION_SEND_ENVIRONMENT.
const ApplicationFlags_send_environment = ApplicationFlags(16)

// ApplicationFlags_non_unique is a representation of the C bitfield member G_APPLICATION_NON_UNIQUE.
const ApplicationFlags_non_unique = ApplicationFlags(32)

// ApplicationFlags_can_override_app_id is a representation of the C bitfield member G_APPLICATION_CAN_OVERRIDE_APP_ID.
const ApplicationFlags_can_override_app_id = ApplicationFlags(64)

// ApplicationFlags_allow_replacement is a representation of the C bitfield member G_APPLICATION_ALLOW_REPLACEMENT.
const ApplicationFlags_allow_replacement = ApplicationFlags(128)

// ApplicationFlags_replace is a representation of the C bitfield member G_APPLICATION_REPLACE.
const ApplicationFlags_replace = ApplicationFlags(256)

// IOStreamSpliceFlags is a representation of the C bitfield GIOStreamSpliceFlags.
type IOStreamSpliceFlags int

// IOStreamSpliceFlags_none is a representation of the C bitfield member G_IO_STREAM_SPLICE_NONE.
const IOStreamSpliceFlags_none = IOStreamSpliceFlags(0)

// IOStreamSpliceFlags_close_stream1 is a representation of the C bitfield member G_IO_STREAM_SPLICE_CLOSE_STREAM1.
const IOStreamSpliceFlags_close_stream1 = IOStreamSpliceFlags(1)

// IOStreamSpliceFlags_close_stream2 is a representation of the C bitfield member G_IO_STREAM_SPLICE_CLOSE_STREAM2.
const IOStreamSpliceFlags_close_stream2 = IOStreamSpliceFlags(2)

// IOStreamSpliceFlags_wait_for_both is a representation of the C bitfield member G_IO_STREAM_SPLICE_WAIT_FOR_BOTH.
const IOStreamSpliceFlags_wait_for_both = IOStreamSpliceFlags(4)

// TlsCertificateFlags is a representation of the C bitfield GTlsCertificateFlags.
type TlsCertificateFlags int

// TlsCertificateFlags_unknown_ca is a representation of the C bitfield member G_TLS_CERTIFICATE_UNKNOWN_CA.
const TlsCertificateFlags_unknown_ca = TlsCertificateFlags(1)

// TlsCertificateFlags_bad_identity is a representation of the C bitfield member G_TLS_CERTIFICATE_BAD_IDENTITY.
const TlsCertificateFlags_bad_identity = TlsCertificateFlags(2)

// TlsCertificateFlags_not_activated is a representation of the C bitfield member G_TLS_CERTIFICATE_NOT_ACTIVATED.
const TlsCertificateFlags_not_activated = TlsCertificateFlags(4)

// TlsCertificateFlags_expired is a representation of the C bitfield member G_TLS_CERTIFICATE_EXPIRED.
const TlsCertificateFlags_expired = TlsCertificateFlags(8)

// TlsCertificateFlags_revoked is a representation of the C bitfield member G_TLS_CERTIFICATE_REVOKED.
const TlsCertificateFlags_revoked = TlsCertificateFlags(16)

// TlsCertificateFlags_insecure is a representation of the C bitfield member G_TLS_CERTIFICATE_INSECURE.
const TlsCertificateFlags_insecure = TlsCertificateFlags(32)

// TlsCertificateFlags_generic_error is a representation of the C bitfield member G_TLS_CERTIFICATE_GENERIC_ERROR.
const TlsCertificateFlags_generic_error = TlsCertificateFlags(64)

// TlsCertificateFlags_validate_all is a representation of the C bitfield member G_TLS_CERTIFICATE_VALIDATE_ALL.
const TlsCertificateFlags_validate_all = TlsCertificateFlags(127)

// TlsAuthenticationMode is a representation of the C enumeration GTlsAuthenticationMode.
type TlsAuthenticationMode int

// TlsAuthenticationMode_none is a representation of the C enumeration member G_TLS_AUTHENTICATION_NONE.
const TlsAuthenticationMode_none = TlsAuthenticationMode(0)

// TlsAuthenticationMode_requested is a representation of the C enumeration member G_TLS_AUTHENTICATION_REQUESTED.
const TlsAuthenticationMode_requested = TlsAuthenticationMode(1)

// TlsAuthenticationMode_required is a representation of the C enumeration member G_TLS_AUTHENTICATION_REQUIRED.
const TlsAuthenticationMode_required = TlsAuthenticationMode(2)

// TlsError is a representation of the C enumeration GTlsError.
type TlsError int

// TlsError_unavailable is a representation of the C enumeration member G_TLS_ERROR_UNAVAILABLE.
const TlsError_unavailable = TlsError(0)

// TlsError_misc is a representation of the C enumeration member G_TLS_ERROR_MISC.
const TlsError_misc = TlsError(1)

// TlsError_bad_certificate is a representation of the C enumeration member G_TLS_ERROR_BAD_CERTIFICATE.
const TlsError_bad_certificate = TlsError(2)

// TlsError_not_tls is a representation of the C enumeration member G_TLS_ERROR_NOT_TLS.
const TlsError_not_tls = TlsError(3)

// TlsError_handshake is a representation of the C enumeration member G_TLS_ERROR_HANDSHAKE.
const TlsError_handshake = TlsError(4)

// TlsError_certificate_required is a representation of the C enumeration member G_TLS_ERROR_CERTIFICATE_REQUIRED.
const TlsError_certificate_required = TlsError(5)

// TlsError_eof is a representation of the C enumeration member G_TLS_ERROR_EOF.
const TlsError_eof = TlsError(6)

// TlsError_inappropriate_fallback is a representation of the C enumeration member G_TLS_ERROR_INAPPROPRIATE_FALLBACK.
const TlsError_inappropriate_fallback = TlsError(7)

// TlsRehandshakeMode is a representation of the C enumeration GTlsRehandshakeMode.
type TlsRehandshakeMode int

// TlsRehandshakeMode_never is a representation of the C enumeration member G_TLS_REHANDSHAKE_NEVER.
const TlsRehandshakeMode_never = TlsRehandshakeMode(0)

// TlsRehandshakeMode_safely is a representation of the C enumeration member G_TLS_REHANDSHAKE_SAFELY.
const TlsRehandshakeMode_safely = TlsRehandshakeMode(1)

// TlsRehandshakeMode_unsafely is a representation of the C enumeration member G_TLS_REHANDSHAKE_UNSAFELY.
const TlsRehandshakeMode_unsafely = TlsRehandshakeMode(2)

// UNSUPPORTED : g_action_parse_detailed_name : throws

// UNSUPPORTED : g_app_info_create_from_commandline : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_app_info_launch_default_for_uri_finish : throws

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get_finish : throws

// UNSUPPORTED : g_bus_get_sync : throws

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

// UNSUPPORTED : g_content_type_guess : has array param, data

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

// UNSUPPORTED : g_content_type_set_mime_dirs : parameter 'dirs' is array parameter without length parameter

// UNSUPPORTED : g_dbus_address_get_for_bus_sync : throws

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_address_get_stream_finish : throws

// UNSUPPORTED : g_dbus_address_get_stream_sync : throws

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

// UNSUPPORTED : g_dbus_error_register_error_domain : has array param, entries

// UNSUPPORTED : g_dbus_gvariant_to_gvalue : has [in]out param, out_gvalue

// UNSUPPORTED : g_dbus_is_supported_address : throws

// UNSUPPORTED : g_dtls_client_connection_new : throws

// UNSUPPORTED : g_dtls_server_connection_new : throws

// UNSUPPORTED : g_file_new_tmp : throws

// UNSUPPORTED : g_icon_new_for_string : throws

// UNSUPPORTED : g_initable_newv : throws

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// PollableSourceNew wraps the C function g_pollable_source_new.
//
// since 2.28
func PollableSourceNew(pollableStream *gobject.Object) *glib.Source {
	sys_pollableStream := pollableStream.ToC()
	retSys := gio.Fn_g_pollable_source_new(sys_pollableStream)
	ret := glib.SourceNewFromC(retSys)

	return ret
}

// UNSUPPORTED : g_pollable_stream_read : throws

// UNSUPPORTED : g_pollable_stream_write : throws

// UNSUPPORTED : g_pollable_stream_write_all : throws

// UNSUPPORTED : g_resource_load : throws

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_resources_get_info : throws

// UNSUPPORTED : g_resources_lookup_data : throws

// UNSUPPORTED : g_resources_open_stream : throws

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_tls_client_connection_new : throws

// UNSUPPORTED : g_tls_file_database_new : throws

// UNSUPPORTED : g_tls_server_connection_new : throws

// UNSUPPORTED : g_unix_mount_at : has [in]out param, time_read

// UNSUPPORTED : g_unix_mount_for : has [in]out param, time_read

// UNSUPPORTED : g_unix_mount_points_get : has [in]out param, time_read

// UNSUPPORTED : g_unix_mounts_get : has [in]out param, time_read

// ActionGroupInterface is a representation of the C record GActionGroupInterface.
//
// since 2.28
type ActionGroupInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionGroupInterface that represents the ActionGroupInterface.
func (recv *ActionGroupInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroupInterfaceNewFromC creates a new ActionGroupInterface from a pointer to the C GActionGroupInterface that represents the ActionGroupInterface.
func ActionGroupInterfaceNewFromC(native unsafe.Pointer) *ActionGroupInterface {
	return &ActionGroupInterface{native: native}
}

// ActionInterface is a representation of the C record GActionInterface.
//
// since 2.28
type ActionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionInterface that represents the ActionInterface.
func (recv *ActionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ActionInterfaceNewFromC creates a new ActionInterface from a pointer to the C GActionInterface that represents the ActionInterface.
func ActionInterfaceNewFromC(native unsafe.Pointer) *ActionInterface {
	return &ActionInterface{native: native}
}

// ApplicationClass is a representation of the C record GApplicationClass.
//
// since 2.28
type ApplicationClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationClass that represents the ApplicationClass.
func (recv *ApplicationClass) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationClassNewFromC creates a new ApplicationClass from a pointer to the C GApplicationClass that represents the ApplicationClass.
func ApplicationClassNewFromC(native unsafe.Pointer) *ApplicationClass {
	return &ApplicationClass{native: native}
}

// ApplicationCommandLineClass is a representation of the C record GApplicationCommandLineClass.
//
// since 2.28
type ApplicationCommandLineClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationCommandLineClass that represents the ApplicationCommandLineClass.
func (recv *ApplicationCommandLineClass) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationCommandLineClassNewFromC creates a new ApplicationCommandLineClass from a pointer to the C GApplicationCommandLineClass that represents the ApplicationCommandLineClass.
func ApplicationCommandLineClassNewFromC(native unsafe.Pointer) *ApplicationCommandLineClass {
	return &ApplicationCommandLineClass{native: native}
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// PollableInputStreamInterface is a representation of the C record GPollableInputStreamInterface.
//
// since 2.28
type PollableInputStreamInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPollableInputStreamInterface that represents the PollableInputStreamInterface.
func (recv *PollableInputStreamInterface) ToC() unsafe.Pointer {
	return recv.native
}

// PollableInputStreamInterfaceNewFromC creates a new PollableInputStreamInterface from a pointer to the C GPollableInputStreamInterface that represents the PollableInputStreamInterface.
func PollableInputStreamInterfaceNewFromC(native unsafe.Pointer) *PollableInputStreamInterface {
	return &PollableInputStreamInterface{native: native}
}

// PollableOutputStreamInterface is a representation of the C record GPollableOutputStreamInterface.
//
// since 2.28
type PollableOutputStreamInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPollableOutputStreamInterface that represents the PollableOutputStreamInterface.
func (recv *PollableOutputStreamInterface) ToC() unsafe.Pointer {
	return recv.native
}

// PollableOutputStreamInterfaceNewFromC creates a new PollableOutputStreamInterface from a pointer to the C GPollableOutputStreamInterface that represents the PollableOutputStreamInterface.
func PollableOutputStreamInterfaceNewFromC(native unsafe.Pointer) *PollableOutputStreamInterface {
	return &PollableOutputStreamInterface{native: native}
}

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// TlsBackendInterface is a representation of the C record GTlsBackendInterface.
//
// since 2.28
type TlsBackendInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsBackendInterface that represents the TlsBackendInterface.
func (recv *TlsBackendInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TlsBackendInterfaceNewFromC creates a new TlsBackendInterface from a pointer to the C GTlsBackendInterface that represents the TlsBackendInterface.
func TlsBackendInterfaceNewFromC(native unsafe.Pointer) *TlsBackendInterface {
	return &TlsBackendInterface{native: native}
}

// Application is a representation of the C record GApplication.
//
// since 2.28
type Application struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplication that represents the Application.
func (recv *Application) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationNewFromC creates a new Application from a pointer to the C GApplication that represents the Application.
func ApplicationNewFromC(native unsafe.Pointer) *Application {
	return &Application{native: native}
}

// SimpleActionGroup is a representation of the C record GSimpleActionGroup.
//
// since 2.28
type SimpleActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleActionGroup that represents the SimpleActionGroup.
func (recv *SimpleActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleActionGroupNewFromC creates a new SimpleActionGroup from a pointer to the C GSimpleActionGroup that represents the SimpleActionGroup.
func SimpleActionGroupNewFromC(native unsafe.Pointer) *SimpleActionGroup {
	return &SimpleActionGroup{native: native}
}

// TcpWrapperConnection is a representation of the C record GTcpWrapperConnection.
//
// since 2.28
type TcpWrapperConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpWrapperConnection that represents the TcpWrapperConnection.
func (recv *TcpWrapperConnection) ToC() unsafe.Pointer {
	return recv.native
}

// TcpWrapperConnectionNewFromC creates a new TcpWrapperConnection from a pointer to the C GTcpWrapperConnection that represents the TcpWrapperConnection.
func TcpWrapperConnectionNewFromC(native unsafe.Pointer) *TcpWrapperConnection {
	return &TcpWrapperConnection{native: native}
}

// TlsCertificate is a representation of the C record GTlsCertificate.
//
// since 2.28
type TlsCertificate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsCertificate that represents the TlsCertificate.
func (recv *TlsCertificate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsCertificateNewFromC creates a new TlsCertificate from a pointer to the C GTlsCertificate that represents the TlsCertificate.
func TlsCertificateNewFromC(native unsafe.Pointer) *TlsCertificate {
	return &TlsCertificate{native: native}
}

// TlsConnection is a representation of the C record GTlsConnection.
//
// since 2.28
type TlsConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsConnection that represents the TlsConnection.
func (recv *TlsConnection) ToC() unsafe.Pointer {
	return recv.native
}

// TlsConnectionNewFromC creates a new TlsConnection from a pointer to the C GTlsConnection that represents the TlsConnection.
func TlsConnectionNewFromC(native unsafe.Pointer) *TlsConnection {
	return &TlsConnection{native: native}
}

// PollableInputStream is a representation of the C interface GPollableInputStream.
//
// since 2.28
type PollableInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPollableInputStream that represents the PollableInputStream.
func (recv *PollableInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// PollableInputStreamNewFromC creates a new PollableInputStream from a pointer to the C GPollableInputStream that represents the PollableInputStream.
func PollableInputStreamNewFromC(native unsafe.Pointer) *PollableInputStream {
	return &PollableInputStream{native: native}
}

// PollableOutputStream is a representation of the C interface GPollableOutputStream.
//
// since 2.28
type PollableOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPollableOutputStream that represents the PollableOutputStream.
func (recv *PollableOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// PollableOutputStreamNewFromC creates a new PollableOutputStream from a pointer to the C GPollableOutputStream that represents the PollableOutputStream.
func PollableOutputStreamNewFromC(native unsafe.Pointer) *PollableOutputStream {
	return &PollableOutputStream{native: native}
}

// TlsBackend is a representation of the C interface GTlsBackend.
//
// since 2.28
type TlsBackend struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsBackend that represents the TlsBackend.
func (recv *TlsBackend) ToC() unsafe.Pointer {
	return recv.native
}

// TlsBackendNewFromC creates a new TlsBackend from a pointer to the C GTlsBackend that represents the TlsBackend.
func TlsBackendNewFromC(native unsafe.Pointer) *TlsBackend {
	return &TlsBackend{native: native}
}

// TlsClientConnection is a representation of the C interface GTlsClientConnection.
//
// since 2.28
type TlsClientConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsClientConnection that represents the TlsClientConnection.
func (recv *TlsClientConnection) ToC() unsafe.Pointer {
	return recv.native
}

// TlsClientConnectionNewFromC creates a new TlsClientConnection from a pointer to the C GTlsClientConnection that represents the TlsClientConnection.
func TlsClientConnectionNewFromC(native unsafe.Pointer) *TlsClientConnection {
	return &TlsClientConnection{native: native}
}

// TlsServerConnection is a representation of the C interface GTlsServerConnection.
//
// since 2.28
type TlsServerConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsServerConnection that represents the TlsServerConnection.
func (recv *TlsServerConnection) ToC() unsafe.Pointer {
	return recv.native
}

// TlsServerConnectionNewFromC creates a new TlsServerConnection from a pointer to the C GTlsServerConnection that represents the TlsServerConnection.
func TlsServerConnectionNewFromC(native unsafe.Pointer) *TlsServerConnection {
	return &TlsServerConnection{native: native}
}
