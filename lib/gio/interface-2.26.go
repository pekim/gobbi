// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// A #GProxy handles connecting to a remote host via a given type of
// proxy server. It is implemented by the 'gio-proxy' extension point.
// The extensions are named after their proxy protocol name. As an
// example, a SOCKS5 proxy implementation can be retrieved with the
// name 'socks5' using the function
// g_io_extension_point_get_extension_by_name().
/*

C record/class : GProxy
*/
type Proxy struct {
	native *C.GProxy
}

func ProxyNewFromC(u unsafe.Pointer) *Proxy {
	c := (*C.GProxy)(u)
	if c == nil {
		return nil
	}

	g := &Proxy{native: c}

	return g
}

func (recv *Proxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Given @connection to communicate with a proxy (eg, a
// #GSocketConnection that is connected to the proxy server), this
// does the necessary handshake to connect to @proxy_address, and if
// required, wraps the #GIOStream to handle proxy payload.
/*

C function : g_proxy_connect
*/
func (recv *Proxy) Connect(connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable) (*IOStream, error) {
	c_connection := (*C.GIOStream)(C.NULL)
	if connection != nil {
		c_connection = (*C.GIOStream)(connection.ToC())
	}

	c_proxy_address := (*C.GProxyAddress)(C.NULL)
	if proxyAddress != nil {
		c_proxy_address = (*C.GProxyAddress)(proxyAddress.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_proxy_connect((*C.GProxy)(recv.native), c_connection, c_proxy_address, c_cancellable, &cThrowableError)
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_proxy_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// See g_proxy_connect().
/*

C function : g_proxy_connect_finish
*/
func (recv *Proxy) ConnectFinish(result *AsyncResult) (*IOStream, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_proxy_connect_finish((*C.GProxy)(recv.native), c_result, &cThrowableError)
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Some proxy protocols expect to be passed a hostname, which they
// will resolve to an IP address themselves. Others, like SOCKS4, do
// not allow this. This function will return %FALSE if @proxy is
// implementing such a protocol. When %FALSE is returned, the caller
// should resolve the destination hostname first, and then pass a
// #GProxyAddress containing the stringified IP address to
// g_proxy_connect() or g_proxy_connect_async().
/*

C function : g_proxy_supports_hostname
*/
func (recv *Proxy) SupportsHostname() bool {
	retC := C.g_proxy_supports_hostname((*C.GProxy)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// #GProxyResolver provides synchronous and asynchronous network proxy
// resolution. #GProxyResolver is used within #GSocketClient through
// the method g_socket_connectable_proxy_enumerate().
//
// Implementations of #GProxyResolver based on libproxy and GNOME settings can
// be found in glib-networking. GIO comes with an implementation for use inside
// Flatpak portals.
/*

C record/class : GProxyResolver
*/
type ProxyResolver struct {
	native *C.GProxyResolver
}

func ProxyResolverNewFromC(u unsafe.Pointer) *ProxyResolver {
	c := (*C.GProxyResolver)(u)
	if c == nil {
		return nil
	}

	g := &ProxyResolver{native: c}

	return g
}

func (recv *ProxyResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Checks if @resolver can be used on this system. (This is used
// internally; g_proxy_resolver_get_default() will only return a proxy
// resolver that returns %TRUE for this method.)
/*

C function : g_proxy_resolver_is_supported
*/
func (recv *ProxyResolver) IsSupported() bool {
	retC := C.g_proxy_resolver_is_supported((*C.GProxyResolver)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_proxy_resolver_lookup : no return type

// Unsupported : g_proxy_resolver_lookup_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_proxy_resolver_lookup_finish : no return type

// Creates a #GSocketAddressEnumerator for @connectable that will
// return #GProxyAddresses for addresses that you must connect
// to via a proxy.
//
// If @connectable does not implement
// g_socket_connectable_proxy_enumerate(), this will fall back to
// calling g_socket_connectable_enumerate().
/*

C function : g_socket_connectable_proxy_enumerate
*/
func (recv *SocketConnectable) ProxyEnumerate() *SocketAddressEnumerator {
	retC := C.g_socket_connectable_proxy_enumerate((*C.GSocketConnectable)(recv.native))
	retGo := SocketAddressEnumeratorNewFromC(unsafe.Pointer(retC))

	return retGo
}
