// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

type ApplicationFlags int

const (
	APPLICATION_FLAGS_NONE           ApplicationFlags = 0
	APPLICATION_IS_SERVICE           ApplicationFlags = 1
	APPLICATION_IS_LAUNCHER          ApplicationFlags = 2
	APPLICATION_HANDLES_OPEN         ApplicationFlags = 4
	APPLICATION_HANDLES_COMMAND_LINE ApplicationFlags = 8
	APPLICATION_SEND_ENVIRONMENT     ApplicationFlags = 16
	APPLICATION_NON_UNIQUE           ApplicationFlags = 32
	APPLICATION_CAN_OVERRIDE_APP_ID  ApplicationFlags = 64
)

type IOStreamSpliceFlags int

const (
	IO_STREAM_SPLICE_NONE          IOStreamSpliceFlags = 0
	IO_STREAM_SPLICE_CLOSE_STREAM1 IOStreamSpliceFlags = 1
	IO_STREAM_SPLICE_CLOSE_STREAM2 IOStreamSpliceFlags = 2
	IO_STREAM_SPLICE_WAIT_FOR_BOTH IOStreamSpliceFlags = 4
)

type TlsCertificateFlags int

const (
	TLS_CERTIFICATE_UNKNOWN_CA    TlsCertificateFlags = 1
	TLS_CERTIFICATE_BAD_IDENTITY  TlsCertificateFlags = 2
	TLS_CERTIFICATE_NOT_ACTIVATED TlsCertificateFlags = 4
	TLS_CERTIFICATE_EXPIRED       TlsCertificateFlags = 8
	TLS_CERTIFICATE_REVOKED       TlsCertificateFlags = 16
	TLS_CERTIFICATE_INSECURE      TlsCertificateFlags = 32
	TLS_CERTIFICATE_GENERIC_ERROR TlsCertificateFlags = 64
	TLS_CERTIFICATE_VALIDATE_ALL  TlsCertificateFlags = 127
)
