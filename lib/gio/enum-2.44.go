// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// The host's network connectivity state, as reported by #GNetworkMonitor.
type NetworkConnectivity C.GNetworkConnectivity

const (
	/*
	   The host is not configured with a
	     route to the Internet; it may or may not be connected to a local
	     network.
	*/
	NETWORK_CONNECTIVITY_LOCAL NetworkConnectivity = 1
	/*
	   The host is connected to a network, but
	     does not appear to be able to reach the full Internet, perhaps
	     due to upstream network problems.
	*/
	NETWORK_CONNECTIVITY_LIMITED NetworkConnectivity = 2
	/*
	   The host is behind a captive portal and
	     cannot reach the full Internet.
	*/
	NETWORK_CONNECTIVITY_PORTAL NetworkConnectivity = 3
	/*
	   The host is connected to a network, and
	     appears to be able to reach the full Internet.
	*/
	NETWORK_CONNECTIVITY_FULL NetworkConnectivity = 4
)
