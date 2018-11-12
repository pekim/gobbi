# `gobject` bitfields

## `BindingFlags`

Flags to be passed to g_object_bind_property() or
g_object_bind_property_full().

This enumeration can be extended at later date.

- BINDING_DEFAULT = 0
- BINDING_BIDIRECTIONAL = 1
- BINDING_SYNC_CREATE = 2
- BINDING_INVERT_BOOLEAN = 4

C - `GBindingFlags`

## `ConnectFlags`

The connection flags are used to specify the behaviour of a signal's
connection.

- CONNECT_AFTER = 1
- CONNECT_SWAPPED = 2

C - `GConnectFlags`

## `ParamFlags`

Through the #GParamFlags flag values, certain aspects of parameters
can be configured. See also #G_PARAM_STATIC_STRINGS.

- PARAM_READABLE = 1
- PARAM_WRITABLE = 2
- PARAM_READWRITE = 3
- PARAM_CONSTRUCT = 4
- PARAM_CONSTRUCT_ONLY = 8
- PARAM_LAX_VALIDATION = 16
- PARAM_STATIC_NAME = 32
- PARAM_PRIVATE = 32
- PARAM_STATIC_NICK = 64
- PARAM_STATIC_BLURB = 128
- PARAM_EXPLICIT_NOTIFY = 1073741824
-  = 2147483648

C - `GParamFlags`

## `SignalFlags`

The signal flags are used to specify a signal's behaviour, the overall
signal description outlines how especially the RUN flags control the
stages of a signal emission.

- SIGNAL_RUN_FIRST = 1
- SIGNAL_RUN_LAST = 2
- SIGNAL_RUN_CLEANUP = 4
- SIGNAL_NO_RECURSE = 8
- SIGNAL_DETAILED = 16
- SIGNAL_ACTION = 32
- SIGNAL_NO_HOOKS = 64
- SIGNAL_MUST_COLLECT = 128
- SIGNAL_DEPRECATED = 256

C - `GSignalFlags`

## `SignalMatchType`

The match types specify what g_signal_handlers_block_matched(),
g_signal_handlers_unblock_matched() and g_signal_handlers_disconnect_matched()
match signals by.

- SIGNAL_MATCH_ID = 1
- SIGNAL_MATCH_DETAIL = 2
- SIGNAL_MATCH_CLOSURE = 4
- SIGNAL_MATCH_FUNC = 8
- SIGNAL_MATCH_DATA = 16
- SIGNAL_MATCH_UNBLOCKED = 32

C - `GSignalMatchType`

## `TypeDebugFlags`

These flags used to be passed to g_type_init_with_debug_flags() which
is now deprecated.

If you need to enable debugging features, use the GOBJECT_DEBUG
environment variable.

- TYPE_DEBUG_NONE = 0
- TYPE_DEBUG_OBJECTS = 1
- TYPE_DEBUG_SIGNALS = 2
- TYPE_DEBUG_INSTANCE_COUNT = 4
- TYPE_DEBUG_MASK = 7

C - `GTypeDebugFlags`

## `TypeFlags`

Bit masks used to check or determine characteristics of a type.

- TYPE_FLAG_ABSTRACT = 16
- TYPE_FLAG_VALUE_ABSTRACT = 32

C - `GTypeFlags`

## `TypeFundamentalFlags`

Bit masks used to check or determine specific characteristics of a
fundamental type.

- TYPE_FLAG_CLASSED = 1
- TYPE_FLAG_INSTANTIATABLE = 2
- TYPE_FLAG_DERIVABLE = 4
- TYPE_FLAG_DEEP_DERIVABLE = 8

C - `GTypeFundamentalFlags`

e.

C - `GTypeFlags`

- TYPE_FLAG_ABSTRACT = %!s(int=16)
- TYPE_FLAG_VALUE_ABSTRACT = %!s(int=32)
## `TypeFundamentalFlags`

Bit masks used to check or determine specific characteristics of a
fundamental type.

C - `GTypeFundamentalFlags`

- TYPE_FLAG_CLASSED = %!s(int=1)
- TYPE_FLAG_INSTANTIATABLE = %!s(int=2)
- TYPE_FLAG_DERIVABLE = %!s(int=4)
- TYPE_FLAG_DEEP_DERIVABLE = %!s(int=8)
