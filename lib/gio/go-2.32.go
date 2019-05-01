// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

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

type ResourceFlags int

const (
	RESOURCE_FLAGS_NONE       ResourceFlags = 0
	RESOURCE_FLAGS_COMPRESSED ResourceFlags = 1
)

type ResourceLookupFlags int

const (
	RESOURCE_LOOKUP_FLAGS_NONE ResourceLookupFlags = 0
)

// Unsupported : g_inet_address_mask_new : return type :

// Unsupported : g_inet_address_mask_new_from_string : return type :

// Unsupported : g_menu_new : return type :

// Unsupported : g_menu_item_new : return type :

// Unsupported : g_menu_item_new_section : return type :

// Unsupported : g_menu_item_new_submenu : return type :

// Unsupported : g_settings_new_full : return type :

const FILE_ATTRIBUTE_FILESYSTEM_USED string = "filesystem::used"
const MENU_ATTRIBUTE_ACTION string = "action"
const MENU_ATTRIBUTE_LABEL string = "label"
const MENU_ATTRIBUTE_TARGET string = "target"
const MENU_LINK_SECTION string = "section"
const MENU_LINK_SUBMENU string = "submenu"

type ResourceError int

const (
	RESOURCE_ERROR_NOT_FOUND ResourceError = 0
	RESOURCE_ERROR_INTERNAL  ResourceError = 1
)

// g_resource_error_quark : return type :
type SocketClientEvent int

const (
	SOCKET_CLIENT_RESOLVING         SocketClientEvent = 0
	SOCKET_CLIENT_RESOLVED          SocketClientEvent = 1
	SOCKET_CLIENT_CONNECTING        SocketClientEvent = 2
	SOCKET_CLIENT_CONNECTED         SocketClientEvent = 3
	SOCKET_CLIENT_PROXY_NEGOTIATING SocketClientEvent = 4
	SOCKET_CLIENT_PROXY_NEGOTIATED  SocketClientEvent = 5
	SOCKET_CLIENT_TLS_HANDSHAKING   SocketClientEvent = 6
	SOCKET_CLIENT_TLS_HANDSHAKED    SocketClientEvent = 7
	SOCKET_CLIENT_COMPLETE          SocketClientEvent = 8
)

// Unsupported : g_file_new_tmp : return type :

// Unsupported : g_network_monitor_get_default : return type :

// Unsupported : g_resource_error_quark : return type :

// Unsupported : g_resource_load : return type :

// Unsupported : g_resources_enumerate_children : array return type :

// Unsupported : g_resources_get_info : return type :

// Unsupported : g_resources_lookup_data : return type :

// Unsupported : g_resources_open_stream : return type :

// Unsupported : g_settings_schema_source_get_default : return type :

// Unsupported : g_resource_new_from_data : return type :

// Unsupported : g_settings_schema_source_new_from_directory : return type :
