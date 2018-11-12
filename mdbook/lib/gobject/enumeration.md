# `gobject` Enums

## `BindingFlags`

Flags to be passed to g_object_bind_property() or
g_object_bind_property_full().

This enumeration can be extended at later date.

C - `GBindingFlags`

## `ConnectFlags`

The connection flags are used to specify the behaviour of a signal's
connection.

C - `GConnectFlags`

## `ParamFlags`

Through the #GParamFlags flag values, certain aspects of parameters
can be configured. See also #G_PARAM_STATIC_STRINGS.

C - `GParamFlags`

## `SignalFlags`

The signal flags are used to specify a signal's behaviour, the overall
signal description outlines how especially the RUN flags control the
stages of a signal emission.

C - `GSignalFlags`

## `SignalMatchType`

The match types specify what g_signal_handlers_block_matched(),
g_signal_handlers_unblock_matched() and g_signal_handlers_disconnect_matched()
match signals by.

C - `GSignalMatchType`

## `TypeDebugFlags`

These flags used to be passed to g_type_init_with_debug_flags() which
is now deprecated.

If you need to enable debugging features, use the GOBJECT_DEBUG
environment variable.

C - `GTypeDebugFlags`

## `TypeFlags`

Bit masks used to check or determine characteristics of a type.

C - `GTypeFlags`

## `TypeFundamentalFlags`

Bit masks used to check or determine specific characteristics of a
fundamental type.

C - `GTypeFundamentalFlags`

