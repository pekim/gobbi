// This is a generated file - DO NOT EDIT

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

/*
The connection flags are used to specify the behaviour of a signal's
connection.
*/
type ConnectFlags C.GConnectFlags

const (
	/*
	   whether the handler should be called before or after the
	    default handler of the signal.
	*/
	CONNECT_AFTER ConnectFlags = 1
	/*
	   whether the instance and data should be swapped when
	    calling the handler; see g_signal_connect_swapped() for an example.
	*/
	CONNECT_SWAPPED ConnectFlags = 2
)

/*
Through the #GParamFlags flag values, certain aspects of parameters
can be configured. See also #G_PARAM_STATIC_STRINGS.
*/
type ParamFlags C.GParamFlags

const (
	// the parameter is readable
	PARAM_READABLE ParamFlags = 1
	// the parameter is writable
	PARAM_WRITABLE ParamFlags = 2
	// alias for %G_PARAM_READABLE | %G_PARAM_WRITABLE
	PARAM_READWRITE ParamFlags = 3
	// the parameter will be set upon object construction
	PARAM_CONSTRUCT ParamFlags = 4
	// the parameter can only be set upon object construction
	PARAM_CONSTRUCT_ONLY ParamFlags = 8
	/*
	   upon parameter conversion (see g_param_value_convert())
	    strict validation is not required
	*/
	PARAM_LAX_VALIDATION ParamFlags = 16
	/*
	   the string used as name when constructing the
	    parameter is guaranteed to remain valid and
	    unmodified for the lifetime of the parameter.
	    Since 2.8
	*/
	PARAM_STATIC_NAME ParamFlags = 32
	// internal
	PARAM_PRIVATE ParamFlags = 32
	/*
	   the string used as nick when constructing the
	    parameter is guaranteed to remain valid and
	    unmmodified for the lifetime of the parameter.
	    Since 2.8
	*/
	PARAM_STATIC_NICK ParamFlags = 64
	/*
	   the string used as blurb when constructing the
	    parameter is guaranteed to remain valid and
	    unmodified for the lifetime of the parameter.
	    Since 2.8
	*/
	PARAM_STATIC_BLURB ParamFlags = 128
	/*
	   calls to g_object_set_property() for this
	     property will not automatically result in a "notify" signal being
	     emitted: the implementation must call g_object_notify() themselves
	     in case the property actually changes.  Since: 2.42.
	*/
	PARAM_EXPLICIT_NOTIFY ParamFlags = 1073741824

// Blacklisted member : G_PARAM_DEPRECATED
)

/*
The signal flags are used to specify a signal's behaviour, the overall
signal description outlines how especially the RUN flags control the
stages of a signal emission.
*/
type SignalFlags C.GSignalFlags

const (
	// Invoke the object method handler in the first emission stage.
	SIGNAL_RUN_FIRST SignalFlags = 1
	// Invoke the object method handler in the third emission stage.
	SIGNAL_RUN_LAST SignalFlags = 2
	// Invoke the object method handler in the last emission stage.
	SIGNAL_RUN_CLEANUP SignalFlags = 4
	/*
	   Signals being emitted for an object while currently being in
	    emission for this very object will not be emitted recursively,
	    but instead cause the first emission to be restarted.
	*/
	SIGNAL_NO_RECURSE SignalFlags = 8
	/*
	   This signal supports "::detail" appendices to the signal name
	    upon handler connections and emissions.
	*/
	SIGNAL_DETAILED SignalFlags = 16
	/*
	   Action signals are signals that may freely be emitted on alive
	    objects from user code via g_signal_emit() and friends, without
	    the need of being embedded into extra code that performs pre or
	    post emission adjustments on the object. They can also be thought
	    of as object methods which can be called generically by
	    third-party code.
	*/
	SIGNAL_ACTION SignalFlags = 32
	// No emissions hooks are supported for this signal.
	SIGNAL_NO_HOOKS SignalFlags = 64
	/*
	   Varargs signal emission will always collect the
	     arguments, even if there are no signal handlers connected.  Since 2.30.
	*/
	SIGNAL_MUST_COLLECT SignalFlags = 128
	/*
	   The signal is deprecated and will be removed
	     in a future version. A warning will be generated if it is connected while
	     running with G_ENABLE_DIAGNOSTIC=1.  Since 2.32.
	*/
	SIGNAL_DEPRECATED SignalFlags = 256
)

/*
The match types specify what g_signal_handlers_block_matched(),
g_signal_handlers_unblock_matched() and g_signal_handlers_disconnect_matched()
match signals by.
*/
type SignalMatchType C.GSignalMatchType

const (
	// The signal id must be equal.
	SIGNAL_MATCH_ID SignalMatchType = 1
	// The signal detail be equal.
	SIGNAL_MATCH_DETAIL SignalMatchType = 2
	// The closure must be the same.
	SIGNAL_MATCH_CLOSURE SignalMatchType = 4
	// The C closure callback must be the same.
	SIGNAL_MATCH_FUNC SignalMatchType = 8
	// The closure data must be the same.
	SIGNAL_MATCH_DATA SignalMatchType = 16
	// Only unblocked signals may matched.
	SIGNAL_MATCH_UNBLOCKED SignalMatchType = 32
)

/*
These flags used to be passed to g_type_init_with_debug_flags() which
is now deprecated.

If you need to enable debugging features, use the GOBJECT_DEBUG
environment variable.
*/
type TypeDebugFlags C.GTypeDebugFlags

const (
	// Print no messages
	TYPE_DEBUG_NONE TypeDebugFlags = 0
	// Print messages about object bookkeeping
	TYPE_DEBUG_OBJECTS TypeDebugFlags = 1
	// Print messages about signal emissions
	TYPE_DEBUG_SIGNALS TypeDebugFlags = 2
	// Keep a count of instances of each type
	TYPE_DEBUG_INSTANCE_COUNT TypeDebugFlags = 4
	// Mask covering all debug flags
	TYPE_DEBUG_MASK TypeDebugFlags = 7
)

// Bit masks used to check or determine characteristics of a type.
type TypeFlags C.GTypeFlags

const (
	/*
	   Indicates an abstract type. No instances can be
	    created for an abstract type
	*/
	TYPE_FLAG_ABSTRACT TypeFlags = 16
	/*
	   Indicates an abstract value type, i.e. a type
	    that introduces a value table, but can't be used for
	    g_value_init()
	*/
	TYPE_FLAG_VALUE_ABSTRACT TypeFlags = 32
)

/*
Bit masks used to check or determine specific characteristics of a
fundamental type.
*/
type TypeFundamentalFlags C.GTypeFundamentalFlags

const (
	// Indicates a classed type
	TYPE_FLAG_CLASSED TypeFundamentalFlags = 1
	// Indicates an instantiable type (implies classed)
	TYPE_FLAG_INSTANTIATABLE TypeFundamentalFlags = 2
	// Indicates a flat derivable type
	TYPE_FLAG_DERIVABLE TypeFundamentalFlags = 4
	// Indicates a deep derivable type (implies derivable)
	TYPE_FLAG_DEEP_DERIVABLE TypeFundamentalFlags = 8
)
