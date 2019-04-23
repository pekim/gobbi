// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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

// Proxy is a wrapper around the C record GProxy.
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

// Equals compares this Proxy with another Proxy, and returns true if they represent the same GObject.
func (recv *Proxy) Equals(other *Proxy) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_proxy_get_default_for_protocol

// Blacklisted : g_proxy_connect

// Unsupported : g_proxy_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_proxy_connect_finish

// Blacklisted : g_proxy_supports_hostname

// ProxyResolver is a wrapper around the C record GProxyResolver.
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

// Equals compares this ProxyResolver with another ProxyResolver, and returns true if they represent the same GObject.
func (recv *ProxyResolver) Equals(other *ProxyResolver) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_proxy_resolver_get_default

// Blacklisted : g_proxy_resolver_is_supported

// Blacklisted : g_proxy_resolver_lookup

// Unsupported : g_proxy_resolver_lookup_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_proxy_resolver_lookup_finish

// Blacklisted : g_socket_connectable_proxy_enumerate
