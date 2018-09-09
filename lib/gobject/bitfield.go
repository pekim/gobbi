package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

type ConnectFlags C.GConnectFlags

const (
	CONNECT_AFTER   ConnectFlags = 1
	CONNECT_SWAPPED ConnectFlags = 2
)

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

type SignalMatchType C.GSignalMatchType

const (
	SIGNAL_MATCH_ID        SignalMatchType = 1
	SIGNAL_MATCH_DETAIL    SignalMatchType = 2
	SIGNAL_MATCH_CLOSURE   SignalMatchType = 4
	SIGNAL_MATCH_FUNC      SignalMatchType = 8
	SIGNAL_MATCH_DATA      SignalMatchType = 16
	SIGNAL_MATCH_UNBLOCKED SignalMatchType = 32
)

type TypeDebugFlags C.GTypeDebugFlags

const (
	TYPE_DEBUG_NONE           TypeDebugFlags = 0
	TYPE_DEBUG_OBJECTS        TypeDebugFlags = 1
	TYPE_DEBUG_SIGNALS        TypeDebugFlags = 2
	TYPE_DEBUG_INSTANCE_COUNT TypeDebugFlags = 4
	TYPE_DEBUG_MASK           TypeDebugFlags = 7
)

type TypeFlags C.GTypeFlags

const (
	TYPE_FLAG_ABSTRACT       TypeFlags = 16
	TYPE_FLAG_VALUE_ABSTRACT TypeFlags = 32
)

type TypeFundamentalFlags C.GTypeFundamentalFlags

const (
	TYPE_FLAG_CLASSED        TypeFundamentalFlags = 1
	TYPE_FLAG_INSTANTIATABLE TypeFundamentalFlags = 2
	TYPE_FLAG_DERIVABLE      TypeFundamentalFlags = 4
	TYPE_FLAG_DEEP_DERIVABLE TypeFundamentalFlags = 8
)
