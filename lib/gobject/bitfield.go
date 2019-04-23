// This is a generated file - DO NOT EDIT

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GConnectFlags

type ParamFlags C.GParamFlags

const (
	PARAM_READABLE        ParamFlags = 1
	PARAM_WRITABLE        ParamFlags = 2
	PARAM_READWRITE       ParamFlags = 3
	PARAM_CONSTRUCT       ParamFlags = 4
	PARAM_CONSTRUCT_ONLY  ParamFlags = 8
	PARAM_LAX_VALIDATION  ParamFlags = 16
	PARAM_STATIC_NAME     ParamFlags = 32
	PARAM_PRIVATE         ParamFlags = 32
	PARAM_STATIC_NICK     ParamFlags = 64
	PARAM_STATIC_BLURB    ParamFlags = 128
	PARAM_EXPLICIT_NOTIFY ParamFlags = 1073741824

// Blacklisted member : G_PARAM_DEPRECATED
)

type SignalFlags C.GSignalFlags

const (
	SIGNAL_RUN_FIRST    SignalFlags = 1
	SIGNAL_RUN_LAST     SignalFlags = 2
	SIGNAL_RUN_CLEANUP  SignalFlags = 4
	SIGNAL_NO_RECURSE   SignalFlags = 8
	SIGNAL_DETAILED     SignalFlags = 16
	SIGNAL_ACTION       SignalFlags = 32
	SIGNAL_NO_HOOKS     SignalFlags = 64
	SIGNAL_MUST_COLLECT SignalFlags = 128
	SIGNAL_DEPRECATED   SignalFlags = 256
)

// Blacklisted : GSignalMatchType

// Blacklisted : GTypeDebugFlags

// Blacklisted : GTypeFlags

// Blacklisted : GTypeFundamentalFlags
