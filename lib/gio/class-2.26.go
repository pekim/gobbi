// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// The #GCredentials type is a reference-counted wrapper for native
// credentials. This information is typically used for identifying,
// authenticating and authorizing other processes.
//
// Some operating systems supports looking up the credentials of the
// remote peer of a communication endpoint - see e.g.
// g_socket_get_credentials().
//
// Some operating systems supports securely sending and receiving
// credentials over a Unix Domain Socket, see
// #GUnixCredentialsMessage, g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials() for details.
//
// On Linux, the native credential type is a struct ucred - see the
// unix(7) man page for details. This corresponds to
// %G_CREDENTIALS_TYPE_LINUX_UCRED.
//
// On FreeBSD, Debian GNU/kFreeBSD, and GNU/Hurd, the native
// credential type is a struct cmsgcred. This corresponds
// to %G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
//
// On NetBSD, the native credential type is a struct unpcbid.
// This corresponds to %G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
//
// On OpenBSD, the native credential type is a struct sockpeercred.
// This corresponds to %G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
//
// On Solaris (including OpenSolaris and its derivatives), the native
// credential type is a ucred_t. This corresponds to
// %G_CREDENTIALS_TYPE_SOLARIS_UCRED.
/*

C type

GCredentials
*/
type Credentials struct {
	native *C.GCredentials
}

func CredentialsNewFromC(u unsafe.Pointer) *Credentials {
	c := (*C.GCredentials)(u)
	if c == nil {
		return nil
	}

	g := &Credentials{native: c}

	return g
}

func (recv *Credentials) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Credentials) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Credentials.
// Exercise care, as this is a potentially dangerous function if the Object is not a Credentials.
func CastToCredentials(object *gobject.Object) *Credentials {
	return CredentialsNewFromC(object.ToC())
}

// Creates a new #GCredentials object with credentials matching the
// the current process.
/*

C function

g_credentials_new
*/
func CredentialsNew() *Credentials {
	retC := C.g_credentials_new()
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a pointer to native credentials of type @native_type from
// @credentials.
//
// It is a programming error (which will cause an warning to be
// logged) to use this method if there is no #GCredentials support for
// the OS or if @native_type isn't supported by the OS.
/*

C function

g_credentials_get_native
*/
func (recv *Credentials) GetNative(nativeType CredentialsType) uintptr {
	c_native_type := (C.GCredentialsType)(nativeType)

	retC := C.g_credentials_get_native((*C.GCredentials)(recv.native), c_native_type)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_credentials_get_unix_user : no return generator

// Checks if @credentials and @other_credentials is the same user.
//
// This operation can fail if #GCredentials is not supported on the
// the OS.
/*

C function

g_credentials_is_same_user
*/
func (recv *Credentials) IsSameUser(otherCredentials *Credentials) (bool, error) {
	c_other_credentials := (*C.GCredentials)(C.NULL)
	if otherCredentials != nil {
		c_other_credentials = (*C.GCredentials)(otherCredentials.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_credentials_is_same_user((*C.GCredentials)(recv.native), c_other_credentials, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Copies the native credentials of type @native_type from @native
// into @credentials.
//
// It is a programming error (which will cause an warning to be
// logged) to use this method if there is no #GCredentials support for
// the OS or if @native_type isn't supported by the OS.
/*

C function

g_credentials_set_native
*/
func (recv *Credentials) SetNative(nativeType CredentialsType, native uintptr) {
	c_native_type := (C.GCredentialsType)(nativeType)

	c_native := (C.gpointer)(native)

	C.g_credentials_set_native((*C.GCredentials)(recv.native), c_native_type, c_native)

	return
}

// Unsupported : g_credentials_set_unix_user : unsupported parameter uid : no type generator for guint (uid_t) for param uid

// Creates a human-readable textual representation of @credentials
// that can be used in logging and debug messages. The format of the
// returned string may change in future GLib release.
/*

C function

g_credentials_to_string
*/
func (recv *Credentials) ToString() string {
	retC := C.g_credentials_to_string((*C.GCredentials)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// The #GDBusAuthObserver type provides a mechanism for participating
// in how a #GDBusServer (or a #GDBusConnection) authenticates remote
// peers. Simply instantiate a #GDBusAuthObserver and connect to the
// signals you are interested in. Note that new signals may be added
// in the future
//
// Controlling Authentication # {#auth-observer}
//
// For example, if you only want to allow D-Bus connections from
// processes owned by the same uid as the server, you would use a
// signal handler like the following:
//
// |[<!-- language="C" -->
// static gboolean
// on_authorize_authenticated_peer (GDBusAuthObserver *observer,
// GIOStream         *stream,
// GCredentials      *credentials,
// gpointer           user_data)
// {
// gboolean authorized;
//
// authorized = FALSE;
// if (credentials != NULL)
// {
// GCredentials *own_credentials;
// own_credentials = g_credentials_new ();
// if (g_credentials_is_same_user (credentials, own_credentials, NULL))
// authorized = TRUE;
// g_object_unref (own_credentials);
// }
//
// return authorized;
// }
// ]|
/*

C type

GDBusAuthObserver
*/
type DBusAuthObserver struct {
	native *C.GDBusAuthObserver
}

func DBusAuthObserverNewFromC(u unsafe.Pointer) *DBusAuthObserver {
	c := (*C.GDBusAuthObserver)(u)
	if c == nil {
		return nil
	}

	g := &DBusAuthObserver{native: c}

	return g
}

func (recv *DBusAuthObserver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *DBusAuthObserver) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusAuthObserver.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusAuthObserver.
func CastToDBusAuthObserver(object *gobject.Object) *DBusAuthObserver {
	return DBusAuthObserverNewFromC(object.ToC())
}

// Creates a new #GDBusAuthObserver object.
/*

C function

g_dbus_auth_observer_new
*/
func DBusAuthObserverNew() *DBusAuthObserver {
	retC := C.g_dbus_auth_observer_new()
	retGo := DBusAuthObserverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Emits the #GDBusAuthObserver::authorize-authenticated-peer signal on @observer.
/*

C function

g_dbus_auth_observer_authorize_authenticated_peer
*/
func (recv *DBusAuthObserver) AuthorizeAuthenticatedPeer(stream *IOStream, credentials *Credentials) bool {
	c_stream := (*C.GIOStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GIOStream)(stream.ToC())
	}

	c_credentials := (*C.GCredentials)(C.NULL)
	if credentials != nil {
		c_credentials = (*C.GCredentials)(credentials.ToC())
	}

	retC := C.g_dbus_auth_observer_authorize_authenticated_peer((*C.GDBusAuthObserver)(recv.native), c_stream, c_credentials)
	retGo := retC == C.TRUE

	return retGo
}

// The #GDBusConnection type is used for D-Bus connections to remote
// peers such as a message buses. It is a low-level API that offers a
// lot of flexibility. For instance, it lets you establish a connection
// over any transport that can by represented as an #GIOStream.
//
// This class is rarely used directly in D-Bus clients. If you are writing
// a D-Bus client, it is often easier to use the g_bus_own_name(),
// g_bus_watch_name() or g_dbus_proxy_new_for_bus() APIs.
//
// As an exception to the usual GLib rule that a particular object must not
// be used by two threads at the same time, #GDBusConnection's methods may be
// called from any thread. This is so that g_bus_get() and g_bus_get_sync()
// can safely return the same #GDBusConnection when called from any thread.
//
// Most of the ways to obtain a #GDBusConnection automatically initialize it
// (i.e. connect to D-Bus): for instance, g_dbus_connection_new() and
// g_bus_get(), and the synchronous versions of those methods, give you an
// initialized connection. Language bindings for GIO should use
// g_initable_new() or g_async_initable_new_async(), which also initialize the
// connection.
//
// If you construct an uninitialized #GDBusConnection, such as via
// g_object_new(), you must initialize it via g_initable_init() or
// g_async_initable_init_async() before using its methods or properties.
// Calling methods or accessing properties on a #GDBusConnection that has not
// completed initialization successfully is considered to be invalid, and leads
// to undefined behaviour. In particular, if initialization fails with a
// #GError, the only valid thing you can do with that #GDBusConnection is to
// free it with g_object_unref().
//
// An example D-Bus server # {#gdbus-server}
//
// Here is an example for a D-Bus server:
// [gdbus-example-server.c](https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-server.c)
//
// An example for exporting a subtree # {#gdbus-subtree-server}
//
// Here is an example for exporting a subtree:
// [gdbus-example-subtree.c](https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-subtree.c)
//
// An example for file descriptor passing # {#gdbus-unix-fd-client}
//
// Here is an example for passing UNIX file descriptors:
// [gdbus-unix-fd-client.c](https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-unix-fd-client.c)
//
// An example for exporting a GObject # {#gdbus-export}
//
// Here is an example for exporting a #GObject:
// [gdbus-example-export.c](https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-export.c)
/*

C type

GDBusConnection
*/
type DBusConnection struct {
	native *C.GDBusConnection
}

func DBusConnectionNewFromC(u unsafe.Pointer) *DBusConnection {
	c := (*C.GDBusConnection)(u)
	if c == nil {
		return nil
	}

	g := &DBusConnection{native: c}

	return g
}

func (recv *DBusConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *DBusConnection) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusConnection.
func CastToDBusConnection(object *gobject.Object) *DBusConnection {
	return DBusConnectionNewFromC(object.ToC())
}

// Finishes an operation started with g_dbus_connection_new().
/*

C function

g_dbus_connection_new_finish
*/
func DBusConnectionNewFinish(res *AsyncResult) (*DBusConnection, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_finish(c_res, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Finishes an operation started with g_dbus_connection_new_for_address().
/*

C function

g_dbus_connection_new_for_address_finish
*/
func DBusConnectionNewForAddressFinish(res *AsyncResult) (*DBusConnection, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_for_address_finish(c_res, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Synchronously connects and sets up a D-Bus client connection for
// exchanging D-Bus messages with an endpoint specified by @address
// which must be in the
// [D-Bus address format](https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// This constructor can only be used to initiate client-side
// connections - use g_dbus_connection_new_sync() if you need to act
// as the server. In particular, @flags cannot contain the
// %G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER or
// %G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS flags.
//
// This is a synchronous failable constructor. See
// g_dbus_connection_new_for_address() for the asynchronous version.
//
// If @observer is not %NULL it may be used to control the
// authentication process.
/*

C function

g_dbus_connection_new_for_address_sync
*/
func DBusConnectionNewForAddressSync(address string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusConnection, error) {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_flags := (C.GDBusConnectionFlags)(flags)

	c_observer := (*C.GDBusAuthObserver)(C.NULL)
	if observer != nil {
		c_observer = (*C.GDBusAuthObserver)(observer.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_for_address_sync(c_address, c_flags, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Synchronously sets up a D-Bus connection for exchanging D-Bus messages
// with the end represented by @stream.
//
// If @stream is a #GSocketConnection, then the corresponding #GSocket
// will be put into non-blocking mode.
//
// The D-Bus connection will interact with @stream from a worker thread.
// As a result, the caller should not interact with @stream after this
// method has been called, except by calling g_object_unref() on it.
//
// If @observer is not %NULL it may be used to control the
// authentication process.
//
// This is a synchronous failable constructor. See
// g_dbus_connection_new() for the asynchronous version.
/*

C function

g_dbus_connection_new_sync
*/
func DBusConnectionNewSync(stream *IOStream, guid string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusConnection, error) {
	c_stream := (*C.GIOStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GIOStream)(stream.ToC())
	}

	c_guid := C.CString(guid)
	defer C.free(unsafe.Pointer(c_guid))

	c_flags := (C.GDBusConnectionFlags)(flags)

	c_observer := (*C.GDBusAuthObserver)(C.NULL)
	if observer != nil {
		c_observer = (*C.GDBusAuthObserver)(observer.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_sync(c_stream, c_guid, c_flags, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_connection_add_filter : unsupported parameter filter_function : no type generator for DBusMessageFilterFunction (GDBusMessageFilterFunction) for param filter_function

// Unsupported : g_dbus_connection_call : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_finish : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_close : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an operation started with g_dbus_connection_close().
/*

C function

g_dbus_connection_close_finish
*/
func (recv *DBusConnection) CloseFinish(res *AsyncResult) (bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_close_finish((*C.GDBusConnection)(recv.native), c_res, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Synchronously closees @connection. The calling thread is blocked
// until this is done. See g_dbus_connection_close() for the
// asynchronous version of this method and more details about what it
// does.
/*

C function

g_dbus_connection_close_sync
*/
func (recv *DBusConnection) CloseSync(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_close_sync((*C.GDBusConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_connection_emit_signal : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_flush : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an operation started with g_dbus_connection_flush().
/*

C function

g_dbus_connection_flush_finish
*/
func (recv *DBusConnection) FlushFinish(res *AsyncResult) (bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_flush_finish((*C.GDBusConnection)(recv.native), c_res, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Synchronously flushes @connection. The calling thread is blocked
// until this is done. See g_dbus_connection_flush() for the
// asynchronous version of this method and more details about what it
// does.
/*

C function

g_dbus_connection_flush_sync
*/
func (recv *DBusConnection) FlushSync(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_flush_sync((*C.GDBusConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the capabilities negotiated with the remote peer
/*

C function

g_dbus_connection_get_capabilities
*/
func (recv *DBusConnection) GetCapabilities() DBusCapabilityFlags {
	retC := C.g_dbus_connection_get_capabilities((*C.GDBusConnection)(recv.native))
	retGo := (DBusCapabilityFlags)(retC)

	return retGo
}

// Gets whether the process is terminated when @connection is
// closed by the remote peer. See
// #GDBusConnection:exit-on-close for more details.
/*

C function

g_dbus_connection_get_exit_on_close
*/
func (recv *DBusConnection) GetExitOnClose() bool {
	retC := C.g_dbus_connection_get_exit_on_close((*C.GDBusConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// The GUID of the peer performing the role of server when
// authenticating. See #GDBusConnection:guid for more details.
/*

C function

g_dbus_connection_get_guid
*/
func (recv *DBusConnection) GetGuid() string {
	retC := C.g_dbus_connection_get_guid((*C.GDBusConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the credentials of the authenticated peer. This will always
// return %NULL unless @connection acted as a server
// (e.g. %G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER was passed)
// when set up and the client passed credentials as part of the
// authentication process.
//
// In a message bus setup, the message bus is always the server and
// each application is a client. So this method will always return
// %NULL for message bus clients.
/*

C function

g_dbus_connection_get_peer_credentials
*/
func (recv *DBusConnection) GetPeerCredentials() *Credentials {
	retC := C.g_dbus_connection_get_peer_credentials((*C.GDBusConnection)(recv.native))
	var retGo (*Credentials)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CredentialsNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the underlying stream used for IO.
//
// While the #GDBusConnection is active, it will interact with this
// stream from a worker thread, so it is not safe to interact with
// the stream directly.
/*

C function

g_dbus_connection_get_stream
*/
func (recv *DBusConnection) GetStream() *IOStream {
	retC := C.g_dbus_connection_get_stream((*C.GDBusConnection)(recv.native))
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the unique name of @connection as assigned by the message
// bus. This can also be used to figure out if @connection is a
// message bus connection.
/*

C function

g_dbus_connection_get_unique_name
*/
func (recv *DBusConnection) GetUniqueName() string {
	retC := C.g_dbus_connection_get_unique_name((*C.GDBusConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets whether @connection is closed.
/*

C function

g_dbus_connection_is_closed
*/
func (recv *DBusConnection) IsClosed() bool {
	retC := C.g_dbus_connection_is_closed((*C.GDBusConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_dbus_connection_register_object : unsupported parameter user_data_free_func : no type generator for GLib.DestroyNotify (GDestroyNotify) for param user_data_free_func

// Unsupported : g_dbus_connection_register_subtree : unsupported parameter user_data_free_func : no type generator for GLib.DestroyNotify (GDestroyNotify) for param user_data_free_func

// Removes a filter.
//
// Note that since filters run in a different thread, there is a race
// condition where it is possible that the filter will be running even
// after calling g_dbus_connection_remove_filter(), so you cannot just
// free data that the filter might be using. Instead, you should pass
// a #GDestroyNotify to g_dbus_connection_add_filter(), which will be
// called when it is guaranteed that the data is no longer needed.
/*

C function

g_dbus_connection_remove_filter
*/
func (recv *DBusConnection) RemoveFilter(filterId uint32) {
	c_filter_id := (C.guint)(filterId)

	C.g_dbus_connection_remove_filter((*C.GDBusConnection)(recv.native), c_filter_id)

	return
}

// Asynchronously sends @message to the peer represented by @connection.
//
// Unless @flags contain the
// %G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL flag, the serial number
// will be assigned by @connection and set on @message via
// g_dbus_message_set_serial(). If @out_serial is not %NULL, then the
// serial number used will be written to this location prior to
// submitting the message to the underlying transport.
//
// If @connection is closed then the operation will fail with
// %G_IO_ERROR_CLOSED. If @message is not well-formed,
// the operation fails with %G_IO_ERROR_INVALID_ARGUMENT.
//
// See this [server][gdbus-server] and [client][gdbus-unix-fd-client]
// for an example of how to use this low-level API to send and receive
// UNIX file descriptors.
//
// Note that @message must be unlocked, unless @flags contain the
// %G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL flag.
/*

C function

g_dbus_connection_send_message
*/
func (recv *DBusConnection) SendMessage(message *DBusMessage, flags DBusSendMessageFlags) (bool, uint32, error) {
	c_message := (*C.GDBusMessage)(C.NULL)
	if message != nil {
		c_message = (*C.GDBusMessage)(message.ToC())
	}

	c_flags := (C.GDBusSendMessageFlags)(flags)

	var c_out_serial C.guint32

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_send_message((*C.GDBusConnection)(recv.native), c_message, c_flags, &c_out_serial, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	outSerial := (uint32)(c_out_serial)

	return retGo, outSerial, goThrowableError
}

// Unsupported : g_dbus_connection_send_message_with_reply : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an operation started with g_dbus_connection_send_message_with_reply().
//
// Note that @error is only set if a local in-process error
// occurred. That is to say that the returned #GDBusMessage object may
// be of type %G_DBUS_MESSAGE_TYPE_ERROR. Use
// g_dbus_message_to_gerror() to transcode this to a #GError.
//
// See this [server][gdbus-server] and [client][gdbus-unix-fd-client]
// for an example of how to use this low-level API to send and receive
// UNIX file descriptors.
/*

C function

g_dbus_connection_send_message_with_reply_finish
*/
func (recv *DBusConnection) SendMessageWithReplyFinish(res *AsyncResult) (*DBusMessage, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_send_message_with_reply_finish((*C.GDBusConnection)(recv.native), c_res, &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Synchronously sends @message to the peer represented by @connection
// and blocks the calling thread until a reply is received or the
// timeout is reached. See g_dbus_connection_send_message_with_reply()
// for the asynchronous version of this method.
//
// Unless @flags contain the
// %G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL flag, the serial number
// will be assigned by @connection and set on @message via
// g_dbus_message_set_serial(). If @out_serial is not %NULL, then the
// serial number used will be written to this location prior to
// submitting the message to the underlying transport.
//
// If @connection is closed then the operation will fail with
// %G_IO_ERROR_CLOSED. If @cancellable is canceled, the operation will
// fail with %G_IO_ERROR_CANCELLED. If @message is not well-formed,
// the operation fails with %G_IO_ERROR_INVALID_ARGUMENT.
//
// Note that @error is only set if a local in-process error
// occurred. That is to say that the returned #GDBusMessage object may
// be of type %G_DBUS_MESSAGE_TYPE_ERROR. Use
// g_dbus_message_to_gerror() to transcode this to a #GError.
//
// See this [server][gdbus-server] and [client][gdbus-unix-fd-client]
// for an example of how to use this low-level API to send and receive
// UNIX file descriptors.
//
// Note that @message must be unlocked, unless @flags contain the
// %G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL flag.
/*

C function

g_dbus_connection_send_message_with_reply_sync
*/
func (recv *DBusConnection) SendMessageWithReplySync(message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int32, cancellable *Cancellable) (*DBusMessage, uint32, error) {
	c_message := (*C.GDBusMessage)(C.NULL)
	if message != nil {
		c_message = (*C.GDBusMessage)(message.ToC())
	}

	c_flags := (C.GDBusSendMessageFlags)(flags)

	c_timeout_msec := (C.gint)(timeoutMsec)

	var c_out_serial C.guint32

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_send_message_with_reply_sync((*C.GDBusConnection)(recv.native), c_message, c_flags, c_timeout_msec, &c_out_serial, c_cancellable, &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	outSerial := (uint32)(c_out_serial)

	return retGo, outSerial, goThrowableError
}

// Sets whether the process should be terminated when @connection is
// closed by the remote peer. See #GDBusConnection:exit-on-close for
// more details.
//
// Note that this function should be used with care. Most modern UNIX
// desktops tie the notion of a user session the session bus, and expect
// all of a users applications to quit when their bus connection goes away.
// If you are setting @exit_on_close to %FALSE for the shared session
// bus connection, you should make sure that your application exits
// when the user session ends.
/*

C function

g_dbus_connection_set_exit_on_close
*/
func (recv *DBusConnection) SetExitOnClose(exitOnClose bool) {
	c_exit_on_close :=
		boolToGboolean(exitOnClose)

	C.g_dbus_connection_set_exit_on_close((*C.GDBusConnection)(recv.native), c_exit_on_close)

	return
}

// Unsupported : g_dbus_connection_signal_subscribe : unsupported parameter callback : no type generator for DBusSignalCallback (GDBusSignalCallback) for param callback

// Unsubscribes from signals.
/*

C function

g_dbus_connection_signal_unsubscribe
*/
func (recv *DBusConnection) SignalUnsubscribe(subscriptionId uint32) {
	c_subscription_id := (C.guint)(subscriptionId)

	C.g_dbus_connection_signal_unsubscribe((*C.GDBusConnection)(recv.native), c_subscription_id)

	return
}

// If @connection was created with
// %G_DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING, this method
// starts processing messages. Does nothing on if @connection wasn't
// created with this flag or if the method has already been called.
/*

C function

g_dbus_connection_start_message_processing
*/
func (recv *DBusConnection) StartMessageProcessing() {
	C.g_dbus_connection_start_message_processing((*C.GDBusConnection)(recv.native))

	return
}

// Unregisters an object.
/*

C function

g_dbus_connection_unregister_object
*/
func (recv *DBusConnection) UnregisterObject(registrationId uint32) bool {
	c_registration_id := (C.guint)(registrationId)

	retC := C.g_dbus_connection_unregister_object((*C.GDBusConnection)(recv.native), c_registration_id)
	retGo := retC == C.TRUE

	return retGo
}

// Unregisters a subtree.
/*

C function

g_dbus_connection_unregister_subtree
*/
func (recv *DBusConnection) UnregisterSubtree(registrationId uint32) bool {
	c_registration_id := (C.guint)(registrationId)

	retC := C.g_dbus_connection_unregister_subtree((*C.GDBusConnection)(recv.native), c_registration_id)
	retGo := retC == C.TRUE

	return retGo
}

// A type for representing D-Bus messages that can be sent or received
// on a #GDBusConnection.
/*

C type

GDBusMessage
*/
type DBusMessage struct {
	native *C.GDBusMessage
}

func DBusMessageNewFromC(u unsafe.Pointer) *DBusMessage {
	c := (*C.GDBusMessage)(u)
	if c == nil {
		return nil
	}

	g := &DBusMessage{native: c}

	return g
}

func (recv *DBusMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *DBusMessage) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusMessage.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusMessage.
func CastToDBusMessage(object *gobject.Object) *DBusMessage {
	return DBusMessageNewFromC(object.ToC())
}

// Creates a new empty #GDBusMessage.
/*

C function

g_dbus_message_new
*/
func DBusMessageNew() *DBusMessage {
	retC := C.g_dbus_message_new()
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GDBusMessage from the data stored at @blob. The byte
// order that the message was in can be retrieved using
// g_dbus_message_get_byte_order().
/*

C function

g_dbus_message_new_from_blob
*/
func DBusMessageNewFromBlob(blob []uint8, capabilities DBusCapabilityFlags) (*DBusMessage, error) {
	c_blob := &blob[0]

	c_blob_len := (C.gsize)(len(blob))

	c_capabilities := (C.GDBusCapabilityFlags)(capabilities)

	var cThrowableError *C.GError

	retC := C.g_dbus_message_new_from_blob((*C.guchar)(unsafe.Pointer(c_blob)), c_blob_len, c_capabilities, &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a new #GDBusMessage for a method call.
/*

C function

g_dbus_message_new_method_call
*/
func DBusMessageNewMethodCall(name string, path string, interface_ string, method string) *DBusMessage {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_interface_ := C.CString(interface_)
	defer C.free(unsafe.Pointer(c_interface_))

	c_method := C.CString(method)
	defer C.free(unsafe.Pointer(c_method))

	retC := C.g_dbus_message_new_method_call(c_name, c_path, c_interface_, c_method)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GDBusMessage for a signal emission.
/*

C function

g_dbus_message_new_signal
*/
func DBusMessageNewSignal(path string, interface_ string, signal string) *DBusMessage {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_interface_ := C.CString(interface_)
	defer C.free(unsafe.Pointer(c_interface_))

	c_signal := C.CString(signal)
	defer C.free(unsafe.Pointer(c_signal))

	retC := C.g_dbus_message_new_signal(c_path, c_interface_, c_signal)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copies @message. The copy is a deep copy and the returned
// #GDBusMessage is completely identical except that it is guaranteed
// to not be locked.
//
// This operation can fail if e.g. @message contains file descriptors
// and the per-process or system-wide open files limit is reached.
/*

C function

g_dbus_message_copy
*/
func (recv *DBusMessage) Copy() (*DBusMessage, error) {
	var cThrowableError *C.GError

	retC := C.g_dbus_message_copy((*C.GDBusMessage)(recv.native), &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Convenience to get the first item in the body of @message.
/*

C function

g_dbus_message_get_arg0
*/
func (recv *DBusMessage) GetArg0() string {
	retC := C.g_dbus_message_get_arg0((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_message_get_body : return type : Blacklisted record : GVariant

// Gets the byte order of @message.
/*

C function

g_dbus_message_get_byte_order
*/
func (recv *DBusMessage) GetByteOrder() DBusMessageByteOrder {
	retC := C.g_dbus_message_get_byte_order((*C.GDBusMessage)(recv.native))
	retGo := (DBusMessageByteOrder)(retC)

	return retGo
}

// Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION header field.
/*

C function

g_dbus_message_get_destination
*/
func (recv *DBusMessage) GetDestination() string {
	retC := C.g_dbus_message_get_destination((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME header field.
/*

C function

g_dbus_message_get_error_name
*/
func (recv *DBusMessage) GetErrorName() string {
	retC := C.g_dbus_message_get_error_name((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the flags for @message.
/*

C function

g_dbus_message_get_flags
*/
func (recv *DBusMessage) GetFlags() DBusMessageFlags {
	retC := C.g_dbus_message_get_flags((*C.GDBusMessage)(recv.native))
	retGo := (DBusMessageFlags)(retC)

	return retGo
}

// Unsupported : g_dbus_message_get_header : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_message_get_header_fields : no return type

// Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE header field.
/*

C function

g_dbus_message_get_interface
*/
func (recv *DBusMessage) GetInterface() string {
	retC := C.g_dbus_message_get_interface((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Checks whether @message is locked. To monitor changes to this
// value, conncet to the #GObject::notify signal to listen for changes
// on the #GDBusMessage:locked property.
/*

C function

g_dbus_message_get_locked
*/
func (recv *DBusMessage) GetLocked() bool {
	retC := C.g_dbus_message_get_locked((*C.GDBusMessage)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_MEMBER header field.
/*

C function

g_dbus_message_get_member
*/
func (recv *DBusMessage) GetMember() string {
	retC := C.g_dbus_message_get_member((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the type of @message.
/*

C function

g_dbus_message_get_message_type
*/
func (recv *DBusMessage) GetMessageType() DBusMessageType {
	retC := C.g_dbus_message_get_message_type((*C.GDBusMessage)(recv.native))
	retGo := (DBusMessageType)(retC)

	return retGo
}

// Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS header field.
/*

C function

g_dbus_message_get_num_unix_fds
*/
func (recv *DBusMessage) GetNumUnixFds() uint32 {
	retC := C.g_dbus_message_get_num_unix_fds((*C.GDBusMessage)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_PATH header field.
/*

C function

g_dbus_message_get_path
*/
func (recv *DBusMessage) GetPath() string {
	retC := C.g_dbus_message_get_path((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL header field.
/*

C function

g_dbus_message_get_reply_serial
*/
func (recv *DBusMessage) GetReplySerial() uint32 {
	retC := C.g_dbus_message_get_reply_serial((*C.GDBusMessage)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_SENDER header field.
/*

C function

g_dbus_message_get_sender
*/
func (recv *DBusMessage) GetSender() string {
	retC := C.g_dbus_message_get_sender((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the serial for @message.
/*

C function

g_dbus_message_get_serial
*/
func (recv *DBusMessage) GetSerial() uint32 {
	retC := C.g_dbus_message_get_serial((*C.GDBusMessage)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE header field.
/*

C function

g_dbus_message_get_signature
*/
func (recv *DBusMessage) GetSignature() string {
	retC := C.g_dbus_message_get_signature((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the UNIX file descriptors associated with @message, if any.
//
// This method is only available on UNIX.
/*

C function

g_dbus_message_get_unix_fd_list
*/
func (recv *DBusMessage) GetUnixFdList() *UnixFDList {
	retC := C.g_dbus_message_get_unix_fd_list((*C.GDBusMessage)(recv.native))
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @message is locked, does nothing. Otherwise locks the message.
/*

C function

g_dbus_message_lock
*/
func (recv *DBusMessage) Lock() {
	C.g_dbus_message_lock((*C.GDBusMessage)(recv.native))

	return
}

// Unsupported : g_dbus_message_new_method_error : unsupported parameter ... : varargs

// Creates a new #GDBusMessage that is an error reply to @method_call_message.
/*

C function

g_dbus_message_new_method_error_literal
*/
func (recv *DBusMessage) NewMethodErrorLiteral(errorName string, errorMessage string) *DBusMessage {
	c_error_name := C.CString(errorName)
	defer C.free(unsafe.Pointer(c_error_name))

	c_error_message := C.CString(errorMessage)
	defer C.free(unsafe.Pointer(c_error_message))

	retC := C.g_dbus_message_new_method_error_literal((*C.GDBusMessage)(recv.native), c_error_name, c_error_message)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_message_new_method_error_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Creates a new #GDBusMessage that is a reply to @method_call_message.
/*

C function

g_dbus_message_new_method_reply
*/
func (recv *DBusMessage) NewMethodReply() *DBusMessage {
	retC := C.g_dbus_message_new_method_reply((*C.GDBusMessage)(recv.native))
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Produces a human-readable multi-line description of @message.
//
// The contents of the description has no ABI guarantees, the contents
// and formatting is subject to change at any time. Typical output
// looks something like this:
// |[
// Flags:   none
// Version: 0
// Serial:  4
// Headers:
// path -> objectpath '/org/gtk/GDBus/TestObject'
// interface -> 'org.gtk.GDBus.TestInterface'
// member -> 'GimmeStdout'
// destination -> ':1.146'
// Body: ()
// UNIX File Descriptors:
// (none)
// ]|
// or
// |[
// Flags:   no-reply-expected
// Version: 0
// Serial:  477
// Headers:
// reply-serial -> uint32 4
// destination -> ':1.159'
// sender -> ':1.146'
// num-unix-fds -> uint32 1
// Body: ()
// UNIX File Descriptors:
// fd 12: dev=0:10,mode=020620,ino=5,uid=500,gid=5,rdev=136:2,size=0,atime=1273085037,mtime=1273085851,ctime=1272982635
// ]|
/*

C function

g_dbus_message_print
*/
func (recv *DBusMessage) Print(indent uint32) string {
	c_indent := (C.guint)(indent)

	retC := C.g_dbus_message_print((*C.GDBusMessage)(recv.native), c_indent)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_message_set_body : unsupported parameter body : Blacklisted record : GVariant

// Sets the byte order of @message.
/*

C function

g_dbus_message_set_byte_order
*/
func (recv *DBusMessage) SetByteOrder(byteOrder DBusMessageByteOrder) {
	c_byte_order := (C.GDBusMessageByteOrder)(byteOrder)

	C.g_dbus_message_set_byte_order((*C.GDBusMessage)(recv.native), c_byte_order)

	return
}

// Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION header field.
/*

C function

g_dbus_message_set_destination
*/
func (recv *DBusMessage) SetDestination(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_destination((*C.GDBusMessage)(recv.native), c_value)

	return
}

// Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME header field.
/*

C function

g_dbus_message_set_error_name
*/
func (recv *DBusMessage) SetErrorName(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_error_name((*C.GDBusMessage)(recv.native), c_value)

	return
}

// Sets the flags to set on @message.
/*

C function

g_dbus_message_set_flags
*/
func (recv *DBusMessage) SetFlags(flags DBusMessageFlags) {
	c_flags := (C.GDBusMessageFlags)(flags)

	C.g_dbus_message_set_flags((*C.GDBusMessage)(recv.native), c_flags)

	return
}

// Unsupported : g_dbus_message_set_header : unsupported parameter value : Blacklisted record : GVariant

// Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE header field.
/*

C function

g_dbus_message_set_interface
*/
func (recv *DBusMessage) SetInterface(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_interface((*C.GDBusMessage)(recv.native), c_value)

	return
}

// Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_MEMBER header field.
/*

C function

g_dbus_message_set_member
*/
func (recv *DBusMessage) SetMember(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_member((*C.GDBusMessage)(recv.native), c_value)

	return
}

// Sets @message to be of @type.
/*

C function

g_dbus_message_set_message_type
*/
func (recv *DBusMessage) SetMessageType(type_ DBusMessageType) {
	c_type := (C.GDBusMessageType)(type_)

	C.g_dbus_message_set_message_type((*C.GDBusMessage)(recv.native), c_type)

	return
}

// Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS header field.
/*

C function

g_dbus_message_set_num_unix_fds
*/
func (recv *DBusMessage) SetNumUnixFds(value uint32) {
	c_value := (C.guint32)(value)

	C.g_dbus_message_set_num_unix_fds((*C.GDBusMessage)(recv.native), c_value)

	return
}

// Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_PATH header field.
/*

C function

g_dbus_message_set_path
*/
func (recv *DBusMessage) SetPath(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_path((*C.GDBusMessage)(recv.native), c_value)

	return
}

// Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL header field.
/*

C function

g_dbus_message_set_reply_serial
*/
func (recv *DBusMessage) SetReplySerial(value uint32) {
	c_value := (C.guint32)(value)

	C.g_dbus_message_set_reply_serial((*C.GDBusMessage)(recv.native), c_value)

	return
}

// Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_SENDER header field.
/*

C function

g_dbus_message_set_sender
*/
func (recv *DBusMessage) SetSender(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_sender((*C.GDBusMessage)(recv.native), c_value)

	return
}

// Sets the serial for @message.
/*

C function

g_dbus_message_set_serial
*/
func (recv *DBusMessage) SetSerial(serial uint32) {
	c_serial := (C.guint32)(serial)

	C.g_dbus_message_set_serial((*C.GDBusMessage)(recv.native), c_serial)

	return
}

// Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE header field.
/*

C function

g_dbus_message_set_signature
*/
func (recv *DBusMessage) SetSignature(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_signature((*C.GDBusMessage)(recv.native), c_value)

	return
}

// Sets the UNIX file descriptors associated with @message. As a
// side-effect the %G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS header
// field is set to the number of fds in @fd_list (or cleared if
// @fd_list is %NULL).
//
// This method is only available on UNIX.
/*

C function

g_dbus_message_set_unix_fd_list
*/
func (recv *DBusMessage) SetUnixFdList(fdList *UnixFDList) {
	c_fd_list := (*C.GUnixFDList)(C.NULL)
	if fdList != nil {
		c_fd_list = (*C.GUnixFDList)(fdList.ToC())
	}

	C.g_dbus_message_set_unix_fd_list((*C.GDBusMessage)(recv.native), c_fd_list)

	return
}

// Unsupported : g_dbus_message_to_blob : no return type

// If @message is not of type %G_DBUS_MESSAGE_TYPE_ERROR does
// nothing and returns %FALSE.
//
// Otherwise this method encodes the error in @message as a #GError
// using g_dbus_error_set_dbus_error() using the information in the
// %G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME header field of @message as
// well as the first string item in @message's body.
/*

C function

g_dbus_message_to_gerror
*/
func (recv *DBusMessage) ToGerror() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_dbus_message_to_gerror((*C.GDBusMessage)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Instances of the #GDBusMethodInvocation class are used when
// handling D-Bus method calls. It provides a way to asynchronously
// return results and errors.
//
// The normal way to obtain a #GDBusMethodInvocation object is to receive
// it as an argument to the handle_method_call() function in a
// #GDBusInterfaceVTable that was passed to g_dbus_connection_register_object().
/*

C type

GDBusMethodInvocation
*/
type DBusMethodInvocation struct {
	native *C.GDBusMethodInvocation
}

func DBusMethodInvocationNewFromC(u unsafe.Pointer) *DBusMethodInvocation {
	c := (*C.GDBusMethodInvocation)(u)
	if c == nil {
		return nil
	}

	g := &DBusMethodInvocation{native: c}

	return g
}

func (recv *DBusMethodInvocation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *DBusMethodInvocation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusMethodInvocation.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusMethodInvocation.
func CastToDBusMethodInvocation(object *gobject.Object) *DBusMethodInvocation {
	return DBusMethodInvocationNewFromC(object.ToC())
}

// Gets the #GDBusConnection the method was invoked on.
/*

C function

g_dbus_method_invocation_get_connection
*/
func (recv *DBusMethodInvocation) GetConnection() *DBusConnection {
	retC := C.g_dbus_method_invocation_get_connection((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the name of the D-Bus interface the method was invoked on.
//
// If this method call is a property Get, Set or GetAll call that has
// been redirected to the method call handler then
// "org.freedesktop.DBus.Properties" will be returned.  See
// #GDBusInterfaceVTable for more information.
/*

C function

g_dbus_method_invocation_get_interface_name
*/
func (recv *DBusMethodInvocation) GetInterfaceName() string {
	retC := C.g_dbus_method_invocation_get_interface_name((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the #GDBusMessage for the method invocation. This is useful if
// you need to use low-level protocol features, such as UNIX file
// descriptor passing, that cannot be properly expressed in the
// #GVariant API.
//
// See this [server][gdbus-server] and [client][gdbus-unix-fd-client]
// for an example of how to use this low-level API to send and receive
// UNIX file descriptors.
/*

C function

g_dbus_method_invocation_get_message
*/
func (recv *DBusMethodInvocation) GetMessage() *DBusMessage {
	retC := C.g_dbus_method_invocation_get_message((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets information about the method call, if any.
//
// If this method invocation is a property Get, Set or GetAll call that
// has been redirected to the method call handler then %NULL will be
// returned.  See g_dbus_method_invocation_get_property_info() and
// #GDBusInterfaceVTable for more information.
/*

C function

g_dbus_method_invocation_get_method_info
*/
func (recv *DBusMethodInvocation) GetMethodInfo() *DBusMethodInfo {
	retC := C.g_dbus_method_invocation_get_method_info((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusMethodInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the name of the method that was invoked.
/*

C function

g_dbus_method_invocation_get_method_name
*/
func (recv *DBusMethodInvocation) GetMethodName() string {
	retC := C.g_dbus_method_invocation_get_method_name((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the object path the method was invoked on.
/*

C function

g_dbus_method_invocation_get_object_path
*/
func (recv *DBusMethodInvocation) GetObjectPath() string {
	retC := C.g_dbus_method_invocation_get_object_path((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_method_invocation_get_parameters : return type : Blacklisted record : GVariant

// Gets the bus name that invoked the method.
/*

C function

g_dbus_method_invocation_get_sender
*/
func (recv *DBusMethodInvocation) GetSender() string {
	retC := C.g_dbus_method_invocation_get_sender((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the @user_data #gpointer passed to g_dbus_connection_register_object().
/*

C function

g_dbus_method_invocation_get_user_data
*/
func (recv *DBusMethodInvocation) GetUserData() uintptr {
	retC := C.g_dbus_method_invocation_get_user_data((*C.GDBusMethodInvocation)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Finishes handling a D-Bus method call by returning an error.
//
// This method will take ownership of @invocation. See
// #GDBusInterfaceVTable for more information about the ownership of
// @invocation.
/*

C function

g_dbus_method_invocation_return_dbus_error
*/
func (recv *DBusMethodInvocation) ReturnDbusError(errorName string, errorMessage string) {
	c_error_name := C.CString(errorName)
	defer C.free(unsafe.Pointer(c_error_name))

	c_error_message := C.CString(errorMessage)
	defer C.free(unsafe.Pointer(c_error_message))

	C.g_dbus_method_invocation_return_dbus_error((*C.GDBusMethodInvocation)(recv.native), c_error_name, c_error_message)

	return
}

// Unsupported : g_dbus_method_invocation_return_error : unsupported parameter ... : varargs

// Like g_dbus_method_invocation_return_error() but without printf()-style formatting.
//
// This method will take ownership of @invocation. See
// #GDBusInterfaceVTable for more information about the ownership of
// @invocation.
/*

C function

g_dbus_method_invocation_return_error_literal
*/
func (recv *DBusMethodInvocation) ReturnErrorLiteral(domain glib.Quark, code int32, message string) {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	C.g_dbus_method_invocation_return_error_literal((*C.GDBusMethodInvocation)(recv.native), c_domain, c_code, c_message)

	return
}

// Unsupported : g_dbus_method_invocation_return_error_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Like g_dbus_method_invocation_return_error() but takes a #GError
// instead of the error domain, error code and message.
//
// This method will take ownership of @invocation. See
// #GDBusInterfaceVTable for more information about the ownership of
// @invocation.
/*

C function

g_dbus_method_invocation_return_gerror
*/
func (recv *DBusMethodInvocation) ReturnGerror(error *glib.Error) {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	C.g_dbus_method_invocation_return_gerror((*C.GDBusMethodInvocation)(recv.native), c_error)

	return
}

// Unsupported : g_dbus_method_invocation_return_value : unsupported parameter parameters : Blacklisted record : GVariant

// #GDBusProxy is a base class used for proxies to access a D-Bus
// interface on a remote object. A #GDBusProxy can be constructed for
// both well-known and unique names.
//
// By default, #GDBusProxy will cache all properties (and listen to
// changes) of the remote object, and proxy all signals that get
// emitted. This behaviour can be changed by passing suitable
// #GDBusProxyFlags when the proxy is created. If the proxy is for a
// well-known name, the property cache is flushed when the name owner
// vanishes and reloaded when a name owner appears.
//
// If a #GDBusProxy is used for a well-known name, the owner of the
// name is tracked and can be read from
// #GDBusProxy:g-name-owner. Connect to the #GObject::notify signal to
// get notified of changes. Additionally, only signals and property
// changes emitted from the current name owner are considered and
// calls are always sent to the current name owner. This avoids a
// number of race conditions when the name is lost by one owner and
// claimed by another. However, if no name owner currently exists,
// then calls will be sent to the well-known name which may result in
// the message bus launching an owner (unless
// %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START is set).
//
// The generic #GDBusProxy::g-properties-changed and
// #GDBusProxy::g-signal signals are not very convenient to work with.
// Therefore, the recommended way of working with proxies is to subclass
// #GDBusProxy, and have more natural properties and signals in your derived
// class. This [example][gdbus-example-gdbus-codegen] shows how this can
// easily be done using the [gdbus-codegen][gdbus-codegen] tool.
//
// A #GDBusProxy instance can be used from multiple threads but note
// that all signals (e.g. #GDBusProxy::g-signal, #GDBusProxy::g-properties-changed
// and #GObject::notify) are emitted in the
// [thread-default main context][g-main-context-push-thread-default]
// of the thread where the instance was constructed.
//
// An example using a proxy for a well-known name can be found in
// [gdbus-example-watch-proxy.c](https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-watch-proxy.c)
/*

C type

GDBusProxy
*/
type DBusProxy struct {
	native *C.GDBusProxy
	// Private : parent_instance
	// Private : priv
}

func DBusProxyNewFromC(u unsafe.Pointer) *DBusProxy {
	c := (*C.GDBusProxy)(u)
	if c == nil {
		return nil
	}

	g := &DBusProxy{native: c}

	return g
}

func (recv *DBusProxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *DBusProxy) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusProxy.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusProxy.
func CastToDBusProxy(object *gobject.Object) *DBusProxy {
	return DBusProxyNewFromC(object.ToC())
}

// Finishes creating a #GDBusProxy.
/*

C function

g_dbus_proxy_new_finish
*/
func DBusProxyNewFinish(res *AsyncResult) (*DBusProxy, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_finish(c_res, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Finishes creating a #GDBusProxy.
/*

C function

g_dbus_proxy_new_for_bus_finish
*/
func DBusProxyNewForBusFinish(res *AsyncResult) (*DBusProxy, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_for_bus_finish(c_res, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Like g_dbus_proxy_new_sync() but takes a #GBusType instead of a #GDBusConnection.
//
// #GDBusProxy is used in this [example][gdbus-wellknown-proxy].
/*

C function

g_dbus_proxy_new_for_bus_sync
*/
func DBusProxyNewForBusSync(busType BusType, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable) (*DBusProxy, error) {
	c_bus_type := (C.GBusType)(busType)

	c_flags := (C.GDBusProxyFlags)(flags)

	c_info := (*C.GDBusInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GDBusInterfaceInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_for_bus_sync(c_bus_type, c_flags, c_info, c_name, c_object_path, c_interface_name, c_cancellable, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a proxy for accessing @interface_name on the remote object
// at @object_path owned by @name at @connection and synchronously
// loads D-Bus properties unless the
// %G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES flag is used.
//
// If the %G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS flag is not set, also sets up
// match rules for signals. Connect to the #GDBusProxy::g-signal signal
// to handle signals from the remote object.
//
// If @name is a well-known name and the
// %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START and %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION
// flags aren't set and no name owner currently exists, the message bus
// will be requested to launch a name owner for the name.
//
// This is a synchronous failable constructor. See g_dbus_proxy_new()
// and g_dbus_proxy_new_finish() for the asynchronous version.
//
// #GDBusProxy is used in this [example][gdbus-wellknown-proxy].
/*

C function

g_dbus_proxy_new_sync
*/
func DBusProxyNewSync(connection *DBusConnection, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable) (*DBusProxy, error) {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	c_flags := (C.GDBusProxyFlags)(flags)

	c_info := (*C.GDBusInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GDBusInterfaceInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_sync(c_connection, c_flags, c_info, c_name, c_object_path, c_interface_name, c_cancellable, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_proxy_call : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_finish : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_get_cached_property : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_get_cached_property_names : no return type

// Gets the connection @proxy is for.
/*

C function

g_dbus_proxy_get_connection
*/
func (recv *DBusProxy) GetConnection() *DBusConnection {
	retC := C.g_dbus_proxy_get_connection((*C.GDBusProxy)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the timeout to use if -1 (specifying default timeout) is
// passed as @timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the #GDBusProxy:g-default-timeout property for more details.
/*

C function

g_dbus_proxy_get_default_timeout
*/
func (recv *DBusProxy) GetDefaultTimeout() int32 {
	retC := C.g_dbus_proxy_get_default_timeout((*C.GDBusProxy)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the flags that @proxy was constructed with.
/*

C function

g_dbus_proxy_get_flags
*/
func (recv *DBusProxy) GetFlags() DBusProxyFlags {
	retC := C.g_dbus_proxy_get_flags((*C.GDBusProxy)(recv.native))
	retGo := (DBusProxyFlags)(retC)

	return retGo
}

// Returns the #GDBusInterfaceInfo, if any, specifying the interface
// that @proxy conforms to. See the #GDBusProxy:g-interface-info
// property for more details.
/*

C function

g_dbus_proxy_get_interface_info
*/
func (recv *DBusProxy) GetInterfaceInfo() *DBusInterfaceInfo {
	retC := C.g_dbus_proxy_get_interface_info((*C.GDBusProxy)(recv.native))
	var retGo (*DBusInterfaceInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the D-Bus interface name @proxy is for.
/*

C function

g_dbus_proxy_get_interface_name
*/
func (recv *DBusProxy) GetInterfaceName() string {
	retC := C.g_dbus_proxy_get_interface_name((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the name that @proxy was constructed for.
/*

C function

g_dbus_proxy_get_name
*/
func (recv *DBusProxy) GetName() string {
	retC := C.g_dbus_proxy_get_name((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// The unique name that owns the name that @proxy is for or %NULL if
// no-one currently owns that name. You may connect to the
// #GObject::notify signal to track changes to the
// #GDBusProxy:g-name-owner property.
/*

C function

g_dbus_proxy_get_name_owner
*/
func (recv *DBusProxy) GetNameOwner() string {
	retC := C.g_dbus_proxy_get_name_owner((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the object path @proxy is for.
/*

C function

g_dbus_proxy_get_object_path
*/
func (recv *DBusProxy) GetObjectPath() string {
	retC := C.g_dbus_proxy_get_object_path((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_proxy_set_cached_property : unsupported parameter value : Blacklisted record : GVariant

// Sets the timeout to use if -1 (specifying default timeout) is
// passed as @timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the #GDBusProxy:g-default-timeout property for more details.
/*

C function

g_dbus_proxy_set_default_timeout
*/
func (recv *DBusProxy) SetDefaultTimeout(timeoutMsec int32) {
	c_timeout_msec := (C.gint)(timeoutMsec)

	C.g_dbus_proxy_set_default_timeout((*C.GDBusProxy)(recv.native), c_timeout_msec)

	return
}

// Ensure that interactions with @proxy conform to the given
// interface. See the #GDBusProxy:g-interface-info property for more
// details.
/*

C function

g_dbus_proxy_set_interface_info
*/
func (recv *DBusProxy) SetInterfaceInfo(info *DBusInterfaceInfo) {
	c_info := (*C.GDBusInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GDBusInterfaceInfo)(info.ToC())
	}

	C.g_dbus_proxy_set_interface_info((*C.GDBusProxy)(recv.native), c_info)

	return
}

// #GDBusServer is a helper for listening to and accepting D-Bus
// connections. This can be used to create a new D-Bus server, allowing two
// peers to use the D-Bus protocol for their own specialized communication.
// A server instance provided in this way will not perform message routing or
// implement the org.freedesktop.DBus interface.
//
// To just export an object on a well-known name on a message bus, such as the
// session or system bus, you should instead use g_bus_own_name().
//
// An example of peer-to-peer communication with G-DBus can be found
// in [gdbus-example-peer.c](https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-peer.c).
/*

C type

GDBusServer
*/
type DBusServer struct {
	native *C.GDBusServer
}

func DBusServerNewFromC(u unsafe.Pointer) *DBusServer {
	c := (*C.GDBusServer)(u)
	if c == nil {
		return nil
	}

	g := &DBusServer{native: c}

	return g
}

func (recv *DBusServer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *DBusServer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusServer.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusServer.
func CastToDBusServer(object *gobject.Object) *DBusServer {
	return DBusServerNewFromC(object.ToC())
}

// Creates a new D-Bus server that listens on the first address in
// @address that works.
//
// Once constructed, you can use g_dbus_server_get_client_address() to
// get a D-Bus address string that clients can use to connect.
//
// Connect to the #GDBusServer::new-connection signal to handle
// incoming connections.
//
// The returned #GDBusServer isn't active - you have to start it with
// g_dbus_server_start().
//
// #GDBusServer is used in this [example][gdbus-peer-to-peer].
//
// This is a synchronous failable constructor. See
// g_dbus_server_new() for the asynchronous version.
/*

C function

g_dbus_server_new_sync
*/
func DBusServerNewSync(address string, flags DBusServerFlags, guid string, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusServer, error) {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_flags := (C.GDBusServerFlags)(flags)

	c_guid := C.CString(guid)
	defer C.free(unsafe.Pointer(c_guid))

	c_observer := (*C.GDBusAuthObserver)(C.NULL)
	if observer != nil {
		c_observer = (*C.GDBusAuthObserver)(observer.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_server_new_sync(c_address, c_flags, c_guid, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusServerNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets a
// [D-Bus address](https://dbus.freedesktop.org/doc/dbus-specification.html#addresses)
// string that can be used by clients to connect to @server.
/*

C function

g_dbus_server_get_client_address
*/
func (recv *DBusServer) GetClientAddress() string {
	retC := C.g_dbus_server_get_client_address((*C.GDBusServer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the flags for @server.
/*

C function

g_dbus_server_get_flags
*/
func (recv *DBusServer) GetFlags() DBusServerFlags {
	retC := C.g_dbus_server_get_flags((*C.GDBusServer)(recv.native))
	retGo := (DBusServerFlags)(retC)

	return retGo
}

// Gets the GUID for @server.
/*

C function

g_dbus_server_get_guid
*/
func (recv *DBusServer) GetGuid() string {
	retC := C.g_dbus_server_get_guid((*C.GDBusServer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets whether @server is active.
/*

C function

g_dbus_server_is_active
*/
func (recv *DBusServer) IsActive() bool {
	retC := C.g_dbus_server_is_active((*C.GDBusServer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Starts @server.
/*

C function

g_dbus_server_start
*/
func (recv *DBusServer) Start() {
	C.g_dbus_server_start((*C.GDBusServer)(recv.native))

	return
}

// Stops @server.
/*

C function

g_dbus_server_stop
*/
func (recv *DBusServer) Stop() {
	C.g_dbus_server_stop((*C.GDBusServer)(recv.native))

	return
}

// Reads a string from the data input stream, up to the first
// occurrence of any of the stop characters.
//
// In contrast to g_data_input_stream_read_until(), this function
// does not consume the stop character. You have to use
// g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto() again.
//
// Note that @stop_chars may contain '\0' if @stop_chars_len is
// specified.
//
// The returned string will always be nul-terminated on success.
/*

C function

g_data_input_stream_read_upto
*/
func (recv *DataInputStream) ReadUpto(stopChars string, stopCharsLen int64, cancellable *Cancellable) (string, uint64, error) {
	c_stop_chars := C.CString(stopChars)
	defer C.free(unsafe.Pointer(c_stop_chars))

	c_stop_chars_len := (C.gssize)(stopCharsLen)

	var c_length C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_upto((*C.GDataInputStream)(recv.native), c_stop_chars, c_stop_chars_len, &c_length, c_cancellable, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goThrowableError
}

// Unsupported : g_data_input_stream_read_upto_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Gets any loaded data from the @ostream. Ownership of the data
// is transferred to the caller; when no longer needed it must be
// freed using the free function set in @ostream's
// #GMemoryOutputStream:destroy-function property.
//
// @ostream must be closed before calling this function.
/*

C function

g_memory_output_stream_steal_data
*/
func (recv *MemoryOutputStream) StealData() uintptr {
	retC := C.g_memory_output_stream_steal_data((*C.GMemoryOutputStream)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Gets @addr's scheme
/*

C function

g_network_address_get_scheme
*/
func (recv *NetworkAddress) GetScheme() string {
	retC := C.g_network_address_get_scheme((*C.GNetworkAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Get's the URI scheme used to resolve proxies. By default, the service name
// is used as scheme.
/*

C function

g_network_service_get_scheme
*/
func (recv *NetworkService) GetScheme() string {
	retC := C.g_network_service_get_scheme((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Set's the URI scheme used to resolve proxies. By default, the service name
// is used as scheme.
/*

C function

g_network_service_set_scheme
*/
func (recv *NetworkService) SetScheme(scheme string) {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	C.g_network_service_set_scheme((*C.GNetworkService)(recv.native), c_scheme)

	return
}

// Attempts to acquire the permission represented by @permission.
//
// The precise method by which this happens depends on the permission
// and the underlying authentication mechanism.  A simple example is
// that a dialog may appear asking the user to enter their password.
//
// You should check with g_permission_get_can_acquire() before calling
// this function.
//
// If the permission is acquired then %TRUE is returned.  Otherwise,
// %FALSE is returned and @error is set appropriately.
//
// This call is blocking, likely for a very long time (in the case that
// user interaction is required).  See g_permission_acquire_async() for
// the non-blocking version.
/*

C function

g_permission_acquire
*/
func (recv *Permission) Acquire(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_permission_acquire((*C.GPermission)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_permission_acquire_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Collects the result of attempting to acquire the permission
// represented by @permission.
//
// This is the second half of the asynchronous version of
// g_permission_acquire().
/*

C function

g_permission_acquire_finish
*/
func (recv *Permission) AcquireFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_permission_acquire_finish((*C.GPermission)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the value of the 'allowed' property.  This property is %TRUE if
// the caller currently has permission to perform the action that
// @permission represents the permission to perform.
/*

C function

g_permission_get_allowed
*/
func (recv *Permission) GetAllowed() bool {
	retC := C.g_permission_get_allowed((*C.GPermission)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value of the 'can-acquire' property.  This property is %TRUE
// if it is generally possible to acquire the permission by calling
// g_permission_acquire().
/*

C function

g_permission_get_can_acquire
*/
func (recv *Permission) GetCanAcquire() bool {
	retC := C.g_permission_get_can_acquire((*C.GPermission)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value of the 'can-release' property.  This property is %TRUE
// if it is generally possible to release the permission by calling
// g_permission_release().
/*

C function

g_permission_get_can_release
*/
func (recv *Permission) GetCanRelease() bool {
	retC := C.g_permission_get_can_release((*C.GPermission)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// This function is called by the #GPermission implementation to update
// the properties of the permission.  You should never call this
// function except from a #GPermission implementation.
//
// GObject notify signals are generated, as appropriate.
/*

C function

g_permission_impl_update
*/
func (recv *Permission) ImplUpdate(allowed bool, canAcquire bool, canRelease bool) {
	c_allowed :=
		boolToGboolean(allowed)

	c_can_acquire :=
		boolToGboolean(canAcquire)

	c_can_release :=
		boolToGboolean(canRelease)

	C.g_permission_impl_update((*C.GPermission)(recv.native), c_allowed, c_can_acquire, c_can_release)

	return
}

// Attempts to release the permission represented by @permission.
//
// The precise method by which this happens depends on the permission
// and the underlying authentication mechanism.  In most cases the
// permission will be dropped immediately without further action.
//
// You should check with g_permission_get_can_release() before calling
// this function.
//
// If the permission is released then %TRUE is returned.  Otherwise,
// %FALSE is returned and @error is set appropriately.
//
// This call is blocking, likely for a very long time (in the case that
// user interaction is required).  See g_permission_release_async() for
// the non-blocking version.
/*

C function

g_permission_release
*/
func (recv *Permission) Release(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_permission_release((*C.GPermission)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_permission_release_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Collects the result of attempting to release the permission
// represented by @permission.
//
// This is the second half of the asynchronous version of
// g_permission_release().
/*

C function

g_permission_release_finish
*/
func (recv *Permission) ReleaseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_permission_release_finish((*C.GPermission)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Support for proxied #GInetSocketAddress.
/*

C type

GProxyAddress
*/
type ProxyAddress struct {
	native *C.GProxyAddress
	// parent_instance : record
	// Private : priv
}

func ProxyAddressNewFromC(u unsafe.Pointer) *ProxyAddress {
	c := (*C.GProxyAddress)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddress{native: c}

	return g
}

func (recv *ProxyAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetSocketAddress upcasts to *InetSocketAddress
func (recv *ProxyAddress) InetSocketAddress() *InetSocketAddress {
	return InetSocketAddressNewFromC(unsafe.Pointer(recv.native))
}

// SocketAddress upcasts to *SocketAddress
func (recv *ProxyAddress) SocketAddress() *SocketAddress {
	return recv.InetSocketAddress().SocketAddress()
}

// Object upcasts to *Object
func (recv *ProxyAddress) Object() *gobject.Object {
	return recv.InetSocketAddress().Object()
}

// CastToWidget down casts any arbitary Object to ProxyAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a ProxyAddress.
func CastToProxyAddress(object *gobject.Object) *ProxyAddress {
	return ProxyAddressNewFromC(object.ToC())
}

// Creates a new #GProxyAddress for @inetaddr with @protocol that should
// tunnel through @dest_hostname and @dest_port.
//
// (Note that this method doesn't set the #GProxyAddress:uri or
// #GProxyAddress:destination-protocol fields; use g_object_new()
// directly if you want to set those.)
/*

C function

g_proxy_address_new
*/
func ProxyAddressNew(inetaddr *InetAddress, port uint16, protocol string, destHostname string, destPort uint16, username string, password string) *ProxyAddress {
	c_inetaddr := (*C.GInetAddress)(C.NULL)
	if inetaddr != nil {
		c_inetaddr = (*C.GInetAddress)(inetaddr.ToC())
	}

	c_port := (C.guint16)(port)

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	c_dest_hostname := C.CString(destHostname)
	defer C.free(unsafe.Pointer(c_dest_hostname))

	c_dest_port := (C.guint16)(destPort)

	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	retC := C.g_proxy_address_new(c_inetaddr, c_port, c_protocol, c_dest_hostname, c_dest_port, c_username, c_password)
	retGo := ProxyAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @proxy's destination hostname; that is, the name of the host
// that will be connected to via the proxy, not the name of the proxy
// itself.
/*

C function

g_proxy_address_get_destination_hostname
*/
func (recv *ProxyAddress) GetDestinationHostname() string {
	retC := C.g_proxy_address_get_destination_hostname((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets @proxy's destination port; that is, the port on the
// destination host that will be connected to via the proxy, not the
// port number of the proxy itself.
/*

C function

g_proxy_address_get_destination_port
*/
func (recv *ProxyAddress) GetDestinationPort() uint16 {
	retC := C.g_proxy_address_get_destination_port((*C.GProxyAddress)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Gets @proxy's password.
/*

C function

g_proxy_address_get_password
*/
func (recv *ProxyAddress) GetPassword() string {
	retC := C.g_proxy_address_get_password((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets @proxy's protocol. eg, "socks" or "http"
/*

C function

g_proxy_address_get_protocol
*/
func (recv *ProxyAddress) GetProtocol() string {
	retC := C.g_proxy_address_get_protocol((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets @proxy's username.
/*

C function

g_proxy_address_get_username
*/
func (recv *ProxyAddress) GetUsername() string {
	retC := C.g_proxy_address_get_username((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Creates a new #GSettings object with the schema specified by
// @schema_id.
//
// Signals on the newly created #GSettings object will be dispatched
// via the thread-default #GMainContext in effect at the time of the
// call to g_settings_new().  The new #GSettings will hold a reference
// on the context.  See g_main_context_push_thread_default().
/*

C function

g_settings_new
*/
func SettingsNew(schemaId string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	retC := C.g_settings_new(c_schema_id)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GSettings object with the schema specified by
// @schema_id and a given #GSettingsBackend.
//
// Creating a #GSettings object with a different backend allows accessing
// settings from a database other than the usual one. For example, it may make
// sense to pass a backend corresponding to the "defaults" settings database on
// the system to get a settings object that modifies the system default
// settings instead of the settings for this user.
/*

C function

g_settings_new_with_backend
*/
func SettingsNewWithBackend(schemaId string, backend *SettingsBackend) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_backend := (*C.GSettingsBackend)(C.NULL)
	if backend != nil {
		c_backend = (*C.GSettingsBackend)(backend.ToC())
	}

	retC := C.g_settings_new_with_backend(c_schema_id, c_backend)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GSettings object with the schema specified by
// @schema_id and a given #GSettingsBackend and path.
//
// This is a mix of g_settings_new_with_backend() and
// g_settings_new_with_path().
/*

C function

g_settings_new_with_backend_and_path
*/
func SettingsNewWithBackendAndPath(schemaId string, backend *SettingsBackend, path string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_backend := (*C.GSettingsBackend)(C.NULL)
	if backend != nil {
		c_backend = (*C.GSettingsBackend)(backend.ToC())
	}

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_settings_new_with_backend_and_path(c_schema_id, c_backend, c_path)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GSettings object with the relocatable schema specified
// by @schema_id and a given path.
//
// You only need to do this if you want to directly create a settings
// object with a schema that doesn't have a specified path of its own.
// That's quite rare.
//
// It is a programmer error to call this function for a schema that
// has an explicitly specified path.
//
// It is a programmer error if @path is not a valid path.  A valid path
// begins and ends with '/' and does not contain two consecutive '/'
// characters.
/*

C function

g_settings_new_with_path
*/
func SettingsNewWithPath(schemaId string, path string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_settings_new_with_path(c_schema_id, c_path)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a binding between the @key in the @settings object
// and the property @property of @object.
//
// The binding uses the default GIO mapping functions to map
// between the settings and property values. These functions
// handle booleans, numeric types and string types in a
// straightforward way. Use g_settings_bind_with_mapping() if
// you need a custom mapping, or map between types that are not
// supported by the default mapping functions.
//
// Unless the @flags include %G_SETTINGS_BIND_NO_SENSITIVITY, this
// function also establishes a binding between the writability of
// @key and the "sensitive" property of @object (if @object has
// a boolean property by that name). See g_settings_bind_writable()
// for more details about writable bindings.
//
// Note that the lifecycle of the binding is tied to @object,
// and that you can have only one binding per object property.
// If you bind the same property twice on the same object, the second
// binding overrides the first one.
/*

C function

g_settings_bind
*/
func (recv *Settings) Bind(key string, object uintptr, property string, flags SettingsBindFlags) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_object := (C.gpointer)(object)

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_flags := (C.GSettingsBindFlags)(flags)

	C.g_settings_bind((*C.GSettings)(recv.native), c_key, c_object, c_property, c_flags)

	return
}

// Unsupported : g_settings_bind_with_mapping : unsupported parameter get_mapping : no type generator for SettingsBindGetMapping (GSettingsBindGetMapping) for param get_mapping

// Create a binding between the writability of @key in the
// @settings object and the property @property of @object.
// The property must be boolean; "sensitive" or "visible"
// properties of widgets are the most likely candidates.
//
// Writable bindings are always uni-directional; changes of the
// writability of the setting will be propagated to the object
// property, not the other way.
//
// When the @inverted argument is %TRUE, the binding inverts the
// value as it passes from the setting to the object, i.e. @property
// will be set to %TRUE if the key is not writable.
//
// Note that the lifecycle of the binding is tied to @object,
// and that you can have only one binding per object property.
// If you bind the same property twice on the same object, the second
// binding overrides the first one.
/*

C function

g_settings_bind_writable
*/
func (recv *Settings) BindWritable(key string, object uintptr, property string, inverted bool) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_object := (C.gpointer)(object)

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_inverted :=
		boolToGboolean(inverted)

	C.g_settings_bind_writable((*C.GSettings)(recv.native), c_key, c_object, c_property, c_inverted)

	return
}

// Changes the #GSettings object into 'delay-apply' mode. In this
// mode, changes to @settings are not immediately propagated to the
// backend, but kept locally until g_settings_apply() is called.
/*

C function

g_settings_delay
*/
func (recv *Settings) Delay() {
	C.g_settings_delay((*C.GSettings)(recv.native))

	return
}

// Unsupported : g_settings_get : unsupported parameter ... : varargs

// Gets the value that is stored at @key in @settings.
//
// A convenience variant of g_settings_get() for booleans.
//
// It is a programmer error to give a @key that isn't specified as
// having a boolean type in the schema for @settings.
/*

C function

g_settings_get_boolean
*/
func (recv *Settings) GetBoolean(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_boolean((*C.GSettings)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Creates a child settings object which has a base path of
// `base-path/@name`, where `base-path` is the base path of
// @settings.
//
// The schema for the child settings object must have been declared
// in the schema of @settings using a <child> element.
/*

C function

g_settings_get_child
*/
func (recv *Settings) GetChild(name string) *Settings {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_get_child((*C.GSettings)(recv.native), c_name)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the value that is stored at @key in @settings.
//
// A convenience variant of g_settings_get() for doubles.
//
// It is a programmer error to give a @key that isn't specified as
// having a 'double' type in the schema for @settings.
/*

C function

g_settings_get_double
*/
func (recv *Settings) GetDouble(key string) float64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_double((*C.GSettings)(recv.native), c_key)
	retGo := (float64)(retC)

	return retGo
}

// Gets the value that is stored in @settings for @key and converts it
// to the enum value that it represents.
//
// In order to use this function the type of the value must be a string
// and it must be marked in the schema file as an enumerated type.
//
// It is a programmer error to give a @key that isn't contained in the
// schema for @settings or is not marked as an enumerated type.
//
// If the value stored in the configuration database is not a valid
// value for the enumerated type then this function will return the
// default value.
/*

C function

g_settings_get_enum
*/
func (recv *Settings) GetEnum(key string) int32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_enum((*C.GSettings)(recv.native), c_key)
	retGo := (int32)(retC)

	return retGo
}

// Gets the value that is stored in @settings for @key and converts it
// to the flags value that it represents.
//
// In order to use this function the type of the value must be an array
// of strings and it must be marked in the schema file as an flags type.
//
// It is a programmer error to give a @key that isn't contained in the
// schema for @settings or is not marked as a flags type.
//
// If the value stored in the configuration database is not a valid
// value for the flags type then this function will return the default
// value.
/*

C function

g_settings_get_flags
*/
func (recv *Settings) GetFlags(key string) uint32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_flags((*C.GSettings)(recv.native), c_key)
	retGo := (uint32)(retC)

	return retGo
}

// Returns whether the #GSettings object has any unapplied
// changes.  This can only be the case if it is in 'delayed-apply' mode.
/*

C function

g_settings_get_has_unapplied
*/
func (recv *Settings) GetHasUnapplied() bool {
	retC := C.g_settings_get_has_unapplied((*C.GSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value that is stored at @key in @settings.
//
// A convenience variant of g_settings_get() for 32-bit integers.
//
// It is a programmer error to give a @key that isn't specified as
// having a int32 type in the schema for @settings.
/*

C function

g_settings_get_int
*/
func (recv *Settings) GetInt(key string) int32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_int((*C.GSettings)(recv.native), c_key)
	retGo := (int32)(retC)

	return retGo
}

// Gets the value that is stored at @key in @settings.
//
// A convenience variant of g_settings_get() for strings.
//
// It is a programmer error to give a @key that isn't specified as
// having a string type in the schema for @settings.
/*

C function

g_settings_get_string
*/
func (recv *Settings) GetString(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_string((*C.GSettings)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_settings_get_strv : no return type

// Unsupported : g_settings_get_value : return type : Blacklisted record : GVariant

// Finds out if a key can be written or not
/*

C function

g_settings_is_writable
*/
func (recv *Settings) IsWritable(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_is_writable((*C.GSettings)(recv.native), c_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_settings_set : unsupported parameter ... : varargs

// Sets @key in @settings to @value.
//
// A convenience variant of g_settings_set() for booleans.
//
// It is a programmer error to give a @key that isn't specified as
// having a boolean type in the schema for @settings.
/*

C function

g_settings_set_boolean
*/
func (recv *Settings) SetBoolean(key string, value bool) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value :=
		boolToGboolean(value)

	retC := C.g_settings_set_boolean((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Sets @key in @settings to @value.
//
// A convenience variant of g_settings_set() for doubles.
//
// It is a programmer error to give a @key that isn't specified as
// having a 'double' type in the schema for @settings.
/*

C function

g_settings_set_double
*/
func (recv *Settings) SetDouble(key string, value float64) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gdouble)(value)

	retC := C.g_settings_set_double((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Sets @key in @settings to @value.
//
// A convenience variant of g_settings_set() for 32-bit integers.
//
// It is a programmer error to give a @key that isn't specified as
// having a int32 type in the schema for @settings.
/*

C function

g_settings_set_int
*/
func (recv *Settings) SetInt(key string, value int32) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	retC := C.g_settings_set_int((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Sets @key in @settings to @value.
//
// A convenience variant of g_settings_set() for strings.
//
// It is a programmer error to give a @key that isn't specified as
// having a string type in the schema for @settings.
/*

C function

g_settings_set_string
*/
func (recv *Settings) SetString(key string, value string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	retC := C.g_settings_set_string((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_settings_set_strv : unsupported parameter value :

// Unsupported : g_settings_set_value : unsupported parameter value : Blacklisted record : GVariant

// Blacklisted : g_settings_backend_changed

// Blacklisted : g_settings_backend_changed_tree

// Unsupported : g_settings_backend_keys_changed : unsupported parameter items :

// Blacklisted : g_settings_backend_path_changed

// Blacklisted : g_settings_backend_path_writable_changed

// Blacklisted : g_settings_backend_writable_changed

// Creates a new #GPermission instance that represents an action that is
// either always or never allowed.
/*

C function

g_simple_permission_new
*/
func SimplePermissionNew(allowed bool) *SimplePermission {
	c_allowed :=
		boolToGboolean(allowed)

	retC := C.g_simple_permission_new(c_allowed)
	retGo := SimplePermissionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the credentials of the foreign process connected to this
// socket, if any (e.g. it is only supported for %G_SOCKET_FAMILY_UNIX
// sockets).
//
// If this operation isn't supported on the OS, the method fails with
// the %G_IO_ERROR_NOT_SUPPORTED error. On Linux this is implemented
// by reading the %SO_PEERCRED option on the underlying socket.
//
// Other ways to obtain credentials from a foreign peer includes the
// #GUnixCredentialsMessage type and
// g_unix_connection_send_credentials() /
// g_unix_connection_receive_credentials() functions.
/*

C function

g_socket_get_credentials
*/
func (recv *Socket) GetCredentials() (*Credentials, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_get_credentials((*C.GSocket)(recv.native), &cThrowableError)
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the timeout setting of the socket. For details on this, see
// g_socket_set_timeout().
/*

C function

g_socket_get_timeout
*/
func (recv *Socket) GetTimeout() uint32 {
	retC := C.g_socket_get_timeout((*C.GSocket)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// This behaves exactly the same as g_socket_receive(), except that
// the choice of blocking or non-blocking behavior is determined by
// the @blocking argument rather than by @socket's properties.
/*

C function

g_socket_receive_with_blocking
*/
func (recv *Socket) ReceiveWithBlocking(buffer []uint8, blocking bool, cancellable *Cancellable) (int64, error) {
	c_buffer := &buffer[0]

	c_size := (C.gsize)(len(buffer))

	c_blocking :=
		boolToGboolean(blocking)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_receive_with_blocking((*C.GSocket)(recv.native), (*C.gchar)(unsafe.Pointer(c_buffer)), c_size, c_blocking, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This behaves exactly the same as g_socket_send(), except that
// the choice of blocking or non-blocking behavior is determined by
// the @blocking argument rather than by @socket's properties.
/*

C function

g_socket_send_with_blocking
*/
func (recv *Socket) SendWithBlocking(buffer []uint8, blocking bool, cancellable *Cancellable) (int64, error) {
	c_buffer := &buffer[0]

	c_size := (C.gsize)(len(buffer))

	c_blocking :=
		boolToGboolean(blocking)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_send_with_blocking((*C.GSocket)(recv.native), (*C.gchar)(unsafe.Pointer(c_buffer)), c_size, c_blocking, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the time in seconds after which I/O operations on @socket will
// time out if they have not yet completed.
//
// On a blocking socket, this means that any blocking #GSocket
// operation will time out after @timeout seconds of inactivity,
// returning %G_IO_ERROR_TIMED_OUT.
//
// On a non-blocking socket, calls to g_socket_condition_wait() will
// also fail with %G_IO_ERROR_TIMED_OUT after the given time. Sources
// created with g_socket_create_source() will trigger after
// @timeout seconds of inactivity, with the requested condition
// set, at which point calling g_socket_receive(), g_socket_send(),
// g_socket_check_connect_result(), etc, will fail with
// %G_IO_ERROR_TIMED_OUT.
//
// If @timeout is 0 (the default), operations will never time out
// on their own.
//
// Note that if an I/O operation is interrupted by a signal, this may
// cause the timeout to be reset.
/*

C function

g_socket_set_timeout
*/
func (recv *Socket) SetTimeout(timeout uint32) {
	c_timeout := (C.guint)(timeout)

	C.g_socket_set_timeout((*C.GSocket)(recv.native), c_timeout)

	return
}

// This is a helper function for g_socket_client_connect().
//
// Attempts to create a TCP connection with a network URI.
//
// @uri may be any valid URI containing an "authority" (hostname/port)
// component. If a port is not specified in the URI, @default_port
// will be used. TLS will be negotiated if #GSocketClient:tls is %TRUE.
// (#GSocketClient does not know to automatically assume TLS for
// certain URI schemes.)
//
// Using this rather than g_socket_client_connect() or
// g_socket_client_connect_to_host() allows #GSocketClient to
// determine when to use application-specific proxy protocols.
//
// Upon a successful connection, a new #GSocketConnection is constructed
// and returned.  The caller owns this new object and must drop their
// reference to it when finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) %NULL is returned and @error (if non-%NULL) is set
// accordingly.
/*

C function

g_socket_client_connect_to_uri
*/
func (recv *SocketClient) ConnectToUri(uri string, defaultPort uint16, cancellable *Cancellable) (*SocketConnection, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_default_port := (C.guint16)(defaultPort)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_uri((*C.GSocketClient)(recv.native), c_uri, c_default_port, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_client_connect_to_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an async connect operation. See g_socket_client_connect_to_uri_async()
/*

C function

g_socket_client_connect_to_uri_finish
*/
func (recv *SocketClient) ConnectToUriFinish(result *AsyncResult) (*SocketConnection, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_uri_finish((*C.GSocketClient)(recv.native), c_result, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the proxy enable state; see g_socket_client_set_enable_proxy()
/*

C function

g_socket_client_get_enable_proxy
*/
func (recv *SocketClient) GetEnableProxy() bool {
	retC := C.g_socket_client_get_enable_proxy((*C.GSocketClient)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the I/O timeout time for sockets created by @client.
//
// See g_socket_client_set_timeout() for details.
/*

C function

g_socket_client_get_timeout
*/
func (recv *SocketClient) GetTimeout() uint32 {
	retC := C.g_socket_client_get_timeout((*C.GSocketClient)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Sets whether or not @client attempts to make connections via a
// proxy server. When enabled (the default), #GSocketClient will use a
// #GProxyResolver to determine if a proxy protocol such as SOCKS is
// needed, and automatically do the necessary proxy negotiation.
//
// See also g_socket_client_set_proxy_resolver().
/*

C function

g_socket_client_set_enable_proxy
*/
func (recv *SocketClient) SetEnableProxy(enable bool) {
	c_enable :=
		boolToGboolean(enable)

	C.g_socket_client_set_enable_proxy((*C.GSocketClient)(recv.native), c_enable)

	return
}

// Sets the I/O timeout for sockets created by @client. @timeout is a
// time in seconds, or 0 for no timeout (the default).
//
// The timeout value affects the initial connection attempt as well,
// so setting this may cause calls to g_socket_client_connect(), etc,
// to fail with %G_IO_ERROR_TIMED_OUT.
/*

C function

g_socket_client_set_timeout
*/
func (recv *SocketClient) SetTimeout(timeout uint32) {
	c_timeout := (C.guint)(timeout)

	C.g_socket_client_set_timeout((*C.GSocketClient)(recv.native), c_timeout)

	return
}

// Receives credentials from the sending end of the connection.  The
// sending end has to call g_unix_connection_send_credentials() (or
// similar) for this to work.
//
// As well as reading the credentials this also reads (and discards) a
// single byte from the stream, as this is required for credentials
// passing to work on some implementations.
//
// Other ways to exchange credentials with a foreign peer includes the
// #GUnixCredentialsMessage type and g_socket_get_credentials() function.
/*

C function

g_unix_connection_receive_credentials
*/
func (recv *UnixConnection) ReceiveCredentials(cancellable *Cancellable) (*Credentials, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_unix_connection_receive_credentials((*C.GUnixConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Passes the credentials of the current user the receiving side
// of the connection. The receiving end has to call
// g_unix_connection_receive_credentials() (or similar) to accept the
// credentials.
//
// As well as sending the credentials this also writes a single NUL
// byte to the stream, as this is required for credentials passing to
// work on some implementations.
//
// Other ways to exchange credentials with a foreign peer includes the
// #GUnixCredentialsMessage type and g_socket_get_credentials() function.
/*

C function

g_unix_connection_send_credentials
*/
func (recv *UnixConnection) SendCredentials(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_unix_connection_send_credentials((*C.GUnixConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This #GSocketControlMessage contains a #GCredentials instance.  It
// may be sent using g_socket_send_message() and received using
// g_socket_receive_message() over UNIX sockets (ie: sockets in the
// %G_SOCKET_FAMILY_UNIX family).
//
// For an easier way to send and receive credentials over
// stream-oriented UNIX sockets, see
// g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials(). To receive credentials of
// a foreign process connected to a socket, use
// g_socket_get_credentials().
/*

C type

GUnixCredentialsMessage
*/
type UnixCredentialsMessage struct {
	native *C.GUnixCredentialsMessage
	// parent_instance : record
	// priv : record
}

func UnixCredentialsMessageNewFromC(u unsafe.Pointer) *UnixCredentialsMessage {
	c := (*C.GUnixCredentialsMessage)(u)
	if c == nil {
		return nil
	}

	g := &UnixCredentialsMessage{native: c}

	return g
}

func (recv *UnixCredentialsMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketControlMessage upcasts to *SocketControlMessage
func (recv *UnixCredentialsMessage) SocketControlMessage() *SocketControlMessage {
	return SocketControlMessageNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *UnixCredentialsMessage) Object() *gobject.Object {
	return recv.SocketControlMessage().Object()
}

// CastToWidget down casts any arbitary Object to UnixCredentialsMessage.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixCredentialsMessage.
func CastToUnixCredentialsMessage(object *gobject.Object) *UnixCredentialsMessage {
	return UnixCredentialsMessageNewFromC(object.ToC())
}

// Creates a new #GUnixCredentialsMessage with credentials matching the current processes.
/*

C function

g_unix_credentials_message_new
*/
func UnixCredentialsMessageNew() *UnixCredentialsMessage {
	retC := C.g_unix_credentials_message_new()
	retGo := UnixCredentialsMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GUnixCredentialsMessage holding @credentials.
/*

C function

g_unix_credentials_message_new_with_credentials
*/
func UnixCredentialsMessageNewWithCredentials(credentials *Credentials) *UnixCredentialsMessage {
	c_credentials := (*C.GCredentials)(C.NULL)
	if credentials != nil {
		c_credentials = (*C.GCredentials)(credentials.ToC())
	}

	retC := C.g_unix_credentials_message_new_with_credentials(c_credentials)
	retGo := UnixCredentialsMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the credentials stored in @message.
/*

C function

g_unix_credentials_message_get_credentials
*/
func (recv *UnixCredentialsMessage) GetCredentials() *Credentials {
	retC := C.g_unix_credentials_message_get_credentials((*C.GUnixCredentialsMessage)(recv.native))
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GUnixSocketAddress of type @type with name @path.
//
// If @type is %G_UNIX_SOCKET_ADDRESS_PATH, this is equivalent to
// calling g_unix_socket_address_new().
//
// If @type is %G_UNIX_SOCKET_ADDRESS_ANONYMOUS, @path and @path_len will be
// ignored.
//
// If @path_type is %G_UNIX_SOCKET_ADDRESS_ABSTRACT, then @path_len
// bytes of @path will be copied to the socket's path, and only those
// bytes will be considered part of the name. (If @path_len is -1,
// then @path is assumed to be NUL-terminated.) For example, if @path
// was "test", then calling g_socket_address_get_native_size() on the
// returned socket would return 7 (2 bytes of overhead, 1 byte for the
// abstract-socket indicator byte, and 4 bytes for the name "test").
//
// If @path_type is %G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED, then
// @path_len bytes of @path will be copied to the socket's path, the
// rest of the path will be padded with 0 bytes, and the entire
// zero-padded buffer will be considered the name. (As above, if
// @path_len is -1, then @path is assumed to be NUL-terminated.) In
// this case, g_socket_address_get_native_size() will always return
// the full size of a `struct sockaddr_un`, although
// g_unix_socket_address_get_path_len() will still return just the
// length of @path.
//
// %G_UNIX_SOCKET_ADDRESS_ABSTRACT is preferred over
// %G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED for new programs. Of course,
// when connecting to a server created by another process, you must
// use the appropriate type corresponding to how that process created
// its listening socket.
/*

C function

g_unix_socket_address_new_with_type
*/
func UnixSocketAddressNewWithType(path []rune, type_ UnixSocketAddressType) *UnixSocketAddress {
	c_path := &path[0]

	c_path_len := (C.gint)(len(path))

	c_type := (C.GUnixSocketAddressType)(type_)

	retC := C.g_unix_socket_address_new_with_type((*C.gchar)(unsafe.Pointer(c_path)), c_path_len, c_type)
	retGo := UnixSocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @address's type.
/*

C function

g_unix_socket_address_get_address_type
*/
func (recv *UnixSocketAddress) GetAddressType() UnixSocketAddressType {
	retC := C.g_unix_socket_address_get_address_type((*C.GUnixSocketAddress)(recv.native))
	retGo := (UnixSocketAddressType)(retC)

	return retGo
}

// Returns the #GZlibCompressor:file-info property.
/*

C function

g_zlib_compressor_get_file_info
*/
func (recv *ZlibCompressor) GetFileInfo() *FileInfo {
	retC := C.g_zlib_compressor_get_file_info((*C.GZlibCompressor)(recv.native))
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets @file_info in @compressor. If non-%NULL, and @compressor's
// #GZlibCompressor:format property is %G_ZLIB_COMPRESSOR_FORMAT_GZIP,
// it will be used to set the file name and modification time in
// the GZIP header of the compressed data.
//
// Note: it is an error to call this function while a compression is in
// progress; it may only be called immediately after creation of @compressor,
// or after resetting it with g_converter_reset().
/*

C function

g_zlib_compressor_set_file_info
*/
func (recv *ZlibCompressor) SetFileInfo(fileInfo *FileInfo) {
	c_file_info := (*C.GFileInfo)(C.NULL)
	if fileInfo != nil {
		c_file_info = (*C.GFileInfo)(fileInfo.ToC())
	}

	C.g_zlib_compressor_set_file_info((*C.GZlibCompressor)(recv.native), c_file_info)

	return
}

// Retrieves the #GFileInfo constructed from the GZIP header data
// of compressed data processed by @compressor, or %NULL if @decompressor's
// #GZlibDecompressor:format property is not %G_ZLIB_COMPRESSOR_FORMAT_GZIP,
// or the header data was not fully processed yet, or it not present in the
// data stream at all.
/*

C function

g_zlib_decompressor_get_file_info
*/
func (recv *ZlibDecompressor) GetFileInfo() *FileInfo {
	retC := C.g_zlib_decompressor_get_file_info((*C.GZlibDecompressor)(recv.native))
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}
