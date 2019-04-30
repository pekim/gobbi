// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// Proxy is a wrapper around the C record GProxy.
type Proxy struct {
	native unsafe.Pointer
}

func ProxyNewFromC(u unsafe.Pointer) *Proxy {
	if u == nil {
		return nil
	}

	g := &Proxy{native: u}

	return g
}

func (recv *Proxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyResolver is a wrapper around the C record GProxyResolver.
type ProxyResolver struct {
	native unsafe.Pointer
}

func ProxyResolverNewFromC(u unsafe.Pointer) *ProxyResolver {
	if u == nil {
		return nil
	}

	g := &ProxyResolver{native: u}

	return g
}

func (recv *ProxyResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
