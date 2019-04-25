// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"fmt"
	glib "github.com/pekim/gobbi/lib/glib"
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
/*

	static void _g_dbus_error_set_dbus_error(GError** error, const gchar* dbus_error_name, const gchar* dbus_error_message, const gchar* format) {
		return g_dbus_error_set_dbus_error(error, dbus_error_name, dbus_error_message, format);
    }
*/
import "C"

type BusType int

const (
	BUS_TYPE_STARTER BusType = -1
	BUS_TYPE_NONE    BusType = 0
	BUS_TYPE_SYSTEM  BusType = 1
	BUS_TYPE_SESSION BusType = 2
)

type CredentialsType int

const (
	CREDENTIALS_TYPE_INVALID              CredentialsType = 0
	CREDENTIALS_TYPE_LINUX_UCRED          CredentialsType = 1
	CREDENTIALS_TYPE_FREEBSD_CMSGCRED     CredentialsType = 2
	CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED CredentialsType = 3
	CREDENTIALS_TYPE_SOLARIS_UCRED        CredentialsType = 4
	CREDENTIALS_TYPE_NETBSD_UNPCBID       CredentialsType = 5
)

type DBusError int

const (
	DBUS_ERROR_FAILED                           DBusError = 0
	DBUS_ERROR_NO_MEMORY                        DBusError = 1
	DBUS_ERROR_SERVICE_UNKNOWN                  DBusError = 2
	DBUS_ERROR_NAME_HAS_NO_OWNER                DBusError = 3
	DBUS_ERROR_NO_REPLY                         DBusError = 4
	DBUS_ERROR_IO_ERROR                         DBusError = 5
	DBUS_ERROR_BAD_ADDRESS                      DBusError = 6
	DBUS_ERROR_NOT_SUPPORTED                    DBusError = 7
	DBUS_ERROR_LIMITS_EXCEEDED                  DBusError = 8
	DBUS_ERROR_ACCESS_DENIED                    DBusError = 9
	DBUS_ERROR_AUTH_FAILED                      DBusError = 10
	DBUS_ERROR_NO_SERVER                        DBusError = 11
	DBUS_ERROR_TIMEOUT                          DBusError = 12
	DBUS_ERROR_NO_NETWORK                       DBusError = 13
	DBUS_ERROR_ADDRESS_IN_USE                   DBusError = 14
	DBUS_ERROR_DISCONNECTED                     DBusError = 15
	DBUS_ERROR_INVALID_ARGS                     DBusError = 16
	DBUS_ERROR_FILE_NOT_FOUND                   DBusError = 17
	DBUS_ERROR_FILE_EXISTS                      DBusError = 18
	DBUS_ERROR_UNKNOWN_METHOD                   DBusError = 19
	DBUS_ERROR_TIMED_OUT                        DBusError = 20
	DBUS_ERROR_MATCH_RULE_NOT_FOUND             DBusError = 21
	DBUS_ERROR_MATCH_RULE_INVALID               DBusError = 22
	DBUS_ERROR_SPAWN_EXEC_FAILED                DBusError = 23
	DBUS_ERROR_SPAWN_FORK_FAILED                DBusError = 24
	DBUS_ERROR_SPAWN_CHILD_EXITED               DBusError = 25
	DBUS_ERROR_SPAWN_CHILD_SIGNALED             DBusError = 26
	DBUS_ERROR_SPAWN_FAILED                     DBusError = 27
	DBUS_ERROR_SPAWN_SETUP_FAILED               DBusError = 28
	DBUS_ERROR_SPAWN_CONFIG_INVALID             DBusError = 29
	DBUS_ERROR_SPAWN_SERVICE_INVALID            DBusError = 30
	DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND          DBusError = 31
	DBUS_ERROR_SPAWN_PERMISSIONS_INVALID        DBusError = 32
	DBUS_ERROR_SPAWN_FILE_INVALID               DBusError = 33
	DBUS_ERROR_SPAWN_NO_MEMORY                  DBusError = 34
	DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN          DBusError = 35
	DBUS_ERROR_INVALID_SIGNATURE                DBusError = 36
	DBUS_ERROR_INVALID_FILE_CONTENT             DBusError = 37
	DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN DBusError = 38
	DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN           DBusError = 39
	DBUS_ERROR_OBJECT_PATH_IN_USE               DBusError = 40
	DBUS_ERROR_UNKNOWN_OBJECT                   DBusError = 41
	DBUS_ERROR_UNKNOWN_INTERFACE                DBusError = 42
	DBUS_ERROR_UNKNOWN_PROPERTY                 DBusError = 43
	DBUS_ERROR_PROPERTY_READ_ONLY               DBusError = 44
)

// DBusErrorEncodeGerror is a wrapper around the C function g_dbus_error_encode_gerror.
func DBusErrorEncodeGerror(error *glib.Error) string {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_encode_gerror(c_error)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// DBusErrorGetRemoteError is a wrapper around the C function g_dbus_error_get_remote_error.
func DBusErrorGetRemoteError(error *glib.Error) string {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_get_remote_error(c_error)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// DBusErrorIsRemoteError is a wrapper around the C function g_dbus_error_is_remote_error.
func DBusErrorIsRemoteError(error *glib.Error) bool {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_is_remote_error(c_error)
	retGo := retC == C.TRUE

	return retGo
}

// DBusErrorNewForDbusError is a wrapper around the C function g_dbus_error_new_for_dbus_error.
func DBusErrorNewForDbusError(dbusErrorName string, dbusErrorMessage string) *glib.Error {
	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	c_dbus_error_message := C.CString(dbusErrorMessage)
	defer C.free(unsafe.Pointer(c_dbus_error_message))

	retC := C.g_dbus_error_new_for_dbus_error(c_dbus_error_name, c_dbus_error_message)
	retGo := glib.ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DBusErrorRegisterError is a wrapper around the C function g_dbus_error_register_error.
func DBusErrorRegisterError(errorDomain glib.Quark, errorCode int32, dbusErrorName string) bool {
	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.gint)(errorCode)

	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	retC := C.g_dbus_error_register_error(c_error_domain, c_error_code, c_dbus_error_name)
	retGo := retC == C.TRUE

	return retGo
}

// g_dbus_error_register_error_domain : unsupported parameter entries :
// DBusErrorSetDbusError is a wrapper around the C function g_dbus_error_set_dbus_error.
func DBusErrorSetDbusError(error *glib.Error, dbusErrorName string, dbusErrorMessage string, format string, args ...interface{}) {
	c_error := (**C.GError)(C.NULL)
	if error != nil {
		c_error = (**C.GError)(error.ToC())
	}

	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	c_dbus_error_message := C.CString(dbusErrorMessage)
	defer C.free(unsafe.Pointer(c_dbus_error_message))

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_dbus_error_set_dbus_error(c_error, c_dbus_error_name, c_dbus_error_message, c_format)

	return
}

// g_dbus_error_set_dbus_error_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args
// DBusErrorStripRemoteError is a wrapper around the C function g_dbus_error_strip_remote_error.
func DBusErrorStripRemoteError(error *glib.Error) bool {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_strip_remote_error(c_error)
	retGo := retC == C.TRUE

	return retGo
}

// DBusErrorUnregisterError is a wrapper around the C function g_dbus_error_unregister_error.
func DBusErrorUnregisterError(errorDomain glib.Quark, errorCode int32, dbusErrorName string) bool {
	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.gint)(errorCode)

	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	retC := C.g_dbus_error_unregister_error(c_error_domain, c_error_code, c_dbus_error_name)
	retGo := retC == C.TRUE

	return retGo
}

type DBusMessageByteOrder int

const (
	DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN    DBusMessageByteOrder = 66
	DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN DBusMessageByteOrder = 108
)

type DBusMessageHeaderField int

const (
	DBUS_MESSAGE_HEADER_FIELD_INVALID      DBusMessageHeaderField = 0
	DBUS_MESSAGE_HEADER_FIELD_PATH         DBusMessageHeaderField = 1
	DBUS_MESSAGE_HEADER_FIELD_INTERFACE    DBusMessageHeaderField = 2
	DBUS_MESSAGE_HEADER_FIELD_MEMBER       DBusMessageHeaderField = 3
	DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME   DBusMessageHeaderField = 4
	DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL DBusMessageHeaderField = 5
	DBUS_MESSAGE_HEADER_FIELD_DESTINATION  DBusMessageHeaderField = 6
	DBUS_MESSAGE_HEADER_FIELD_SENDER       DBusMessageHeaderField = 7
	DBUS_MESSAGE_HEADER_FIELD_SIGNATURE    DBusMessageHeaderField = 8
	DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS DBusMessageHeaderField = 9
)

type DBusMessageType int

const (
	DBUS_MESSAGE_TYPE_INVALID       DBusMessageType = 0
	DBUS_MESSAGE_TYPE_METHOD_CALL   DBusMessageType = 1
	DBUS_MESSAGE_TYPE_METHOD_RETURN DBusMessageType = 2
	DBUS_MESSAGE_TYPE_ERROR         DBusMessageType = 3
	DBUS_MESSAGE_TYPE_SIGNAL        DBusMessageType = 4
)

type UnixSocketAddressType int

const (
	UNIX_SOCKET_ADDRESS_INVALID         UnixSocketAddressType = 0
	UNIX_SOCKET_ADDRESS_ANONYMOUS       UnixSocketAddressType = 1
	UNIX_SOCKET_ADDRESS_PATH            UnixSocketAddressType = 2
	UNIX_SOCKET_ADDRESS_ABSTRACT        UnixSocketAddressType = 3
	UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED UnixSocketAddressType = 4
)
