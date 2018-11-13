// This is a generated file - DO NOT EDIT
// +build gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Checks if the network is metered.
// See #GNetworkMonitor:network-metered for more details.
/*

C function

g_network_monitor_get_network_metered
*/
func (recv *NetworkMonitor) GetNetworkMetered() bool {
	retC := C.g_network_monitor_get_network_metered((*C.GNetworkMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Copies session state from one connection to another. This is
// not normally needed, but may be used when the same session
// needs to be used between different endpoints as is required
// by some protocols such as FTP over TLS. @source should have
// already completed a handshake, and @conn should not have
// completed a handshake.
/*

C function

g_tls_client_connection_copy_session_state
*/
func (recv *TlsClientConnection) CopySessionState(source *TlsClientConnection) {
	c_source := (*C.GTlsClientConnection)(source.ToC())

	C.g_tls_client_connection_copy_session_state((*C.GTlsClientConnection)(recv.native), c_source)

	return
}
