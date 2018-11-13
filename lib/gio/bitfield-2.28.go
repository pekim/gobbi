// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

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

// Flags used to define the behaviour of a #GApplication.
type ApplicationFlags C.GApplicationFlags

const (
	// Default
	APPLICATION_FLAGS_NONE ApplicationFlags = 0
	/*
	   Run as a service. In this mode, registration
	        fails if the service is already running, and the application
	        will initially wait up to 10 seconds for an initial activation
	        message to arrive.
	*/
	APPLICATION_IS_SERVICE ApplicationFlags = 1
	// Don't try to become the primary instance.
	APPLICATION_IS_LAUNCHER ApplicationFlags = 2
	/*
	   This application handles opening files (in
	       the primary instance). Note that this flag only affects the default
	       implementation of local_command_line(), and has no effect if
	       %G_APPLICATION_HANDLES_COMMAND_LINE is given.
	       See g_application_run() for details.
	*/
	APPLICATION_HANDLES_OPEN ApplicationFlags = 4
	/*
	   This application handles command line
	       arguments (in the primary instance). Note that this flag only affect
	       the default implementation of local_command_line().
	       See g_application_run() for details.
	*/
	APPLICATION_HANDLES_COMMAND_LINE ApplicationFlags = 8
	/*
	   Send the environment of the
	       launching process to the primary instance. Set this flag if your
	       application is expected to behave differently depending on certain
	       environment variables. For instance, an editor might be expected
	       to use the `GIT_COMMITTER_NAME` environment variable
	       when editing a git commit message. The environment is available
	       to the #GApplication::command-line signal handler, via
	       g_application_command_line_getenv().
	*/
	APPLICATION_SEND_ENVIRONMENT ApplicationFlags = 16
	/*
	   Make no attempts to do any of the typical
	       single-instance application negotiation, even if the application
	       ID is given.  The application neither attempts to become the
	       owner of the application ID nor does it check if an existing
	       owner already exists.  Everything occurs in the local process.
	       Since: 2.30.
	*/
	APPLICATION_NON_UNIQUE ApplicationFlags = 32
	/*
	   Allow users to override the
	       application ID from the command line with `--gapplication-app-id`.
	       Since: 2.48
	*/
	APPLICATION_CAN_OVERRIDE_APP_ID ApplicationFlags = 64
)

// GIOStreamSpliceFlags determine how streams should be spliced.
type IOStreamSpliceFlags C.GIOStreamSpliceFlags

const (
	// Do not close either stream.
	IO_STREAM_SPLICE_NONE IOStreamSpliceFlags = 0
	/*
	   Close the first stream after
	       the splice.
	*/
	IO_STREAM_SPLICE_CLOSE_STREAM1 IOStreamSpliceFlags = 1
	/*
	   Close the second stream after
	       the splice.
	*/
	IO_STREAM_SPLICE_CLOSE_STREAM2 IOStreamSpliceFlags = 2
	/*
	   Wait for both splice operations to finish
	       before calling the callback.
	*/
	IO_STREAM_SPLICE_WAIT_FOR_BOTH IOStreamSpliceFlags = 4
)

/*
A set of flags describing TLS certification validation. This can be
used to set which validation steps to perform (eg, with
g_tls_client_connection_set_validation_flags()), or to describe why
a particular certificate was rejected (eg, in
#GTlsConnection::accept-certificate).
*/
type TlsCertificateFlags C.GTlsCertificateFlags

const (
	/*
	   The signing certificate authority is
	     not known.
	*/
	TLS_CERTIFICATE_UNKNOWN_CA TlsCertificateFlags = 1
	/*
	   The certificate does not match the
	     expected identity of the site that it was retrieved from.
	*/
	TLS_CERTIFICATE_BAD_IDENTITY TlsCertificateFlags = 2
	/*
	   The certificate's activation time
	     is still in the future
	*/
	TLS_CERTIFICATE_NOT_ACTIVATED TlsCertificateFlags = 4
	// The certificate has expired
	TLS_CERTIFICATE_EXPIRED TlsCertificateFlags = 8
	/*
	   The certificate has been revoked
	     according to the #GTlsConnection's certificate revocation list.
	*/
	TLS_CERTIFICATE_REVOKED TlsCertificateFlags = 16
	/*
	   The certificate's algorithm is
	     considered insecure.
	*/
	TLS_CERTIFICATE_INSECURE TlsCertificateFlags = 32
	/*
	   Some other error occurred validating
	     the certificate
	*/
	TLS_CERTIFICATE_GENERIC_ERROR TlsCertificateFlags = 64
	/*
	   the combination of all of the above
	     flags
	*/
	TLS_CERTIFICATE_VALIDATE_ALL TlsCertificateFlags = 127
)
