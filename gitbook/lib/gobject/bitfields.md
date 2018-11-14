# `gobject` bitfields

## `BindingFlags`

Flags to be passed to g_object_bind_property() or
g_object_bind_property_full().

This enumeration can be extended at later date.

C - `GBindingFlags`

### BINDING_DEFAULT = 0
The default binding; if the source property
  changes, the target property is updated with its value.

### BINDING_BIDIRECTIONAL = 1
Bidirectional binding; if either the
  property of the source or the property of the target changes,
  the other is updated.

### BINDING_SYNC_CREATE = 2
Synchronize the values of the source and
  target properties when creating the binding; the direction of
  the synchronization is always from the source to the target.

### BINDING_INVERT_BOOLEAN = 4
If the two properties being bound are
  booleans, setting one to %TRUE will result in the other being
  set to %FALSE and vice versa. This flag will only work for
  boolean properties, and cannot be used when passing custom
  transformation functions to g_object_bind_property_full().


## `ConnectFlags`

The connection flags are used to specify the behaviour of a signal's
connection.

C - `GConnectFlags`

### CONNECT_AFTER = 1
whether the handler should be called before or after the
 default handler of the signal.

### CONNECT_SWAPPED = 2
whether the instance and data should be swapped when
 calling the handler; see g_signal_connect_swapped() for an example.


## `ParamFlags`

Through the #GParamFlags flag values, certain aspects of parameters
can be configured. See also #G_PARAM_STATIC_STRINGS.

C - `GParamFlags`

### PARAM_READABLE = 1
the parameter is readable

### PARAM_WRITABLE = 2
the parameter is writable

### PARAM_READWRITE = 3
alias for %G_PARAM_READABLE | %G_PARAM_WRITABLE

### PARAM_CONSTRUCT = 4
the parameter will be set upon object construction

### PARAM_CONSTRUCT_ONLY = 8
the parameter can only be set upon object construction

### PARAM_LAX_VALIDATION = 16
upon parameter conversion (see g_param_value_convert())
 strict validation is not required

### PARAM_STATIC_NAME = 32
the string used as name when constructing the
 parameter is guaranteed to remain valid and
 unmodified for the lifetime of the parameter.
 Since 2.8

### PARAM_PRIVATE = 32
internal

### PARAM_STATIC_NICK = 64
the string used as nick when constructing the
 parameter is guaranteed to remain valid and
 unmmodified for the lifetime of the parameter.
 Since 2.8

### PARAM_STATIC_BLURB = 128
the string used as blurb when constructing the
 parameter is guaranteed to remain valid and
 unmodified for the lifetime of the parameter.
 Since 2.8

### PARAM_EXPLICIT_NOTIFY = 1073741824
calls to g_object_set_property() for this
  property will not automatically result in a "notify" signal being
  emitted: the implementation must call g_object_notify() themselves
  in case the property actually changes.  Since: 2.42.

###  = 2147483648
the parameter is deprecated and will be removed
 in a future version. A warning will be generated if it is used
 while running with G_ENABLE_DIAGNOSTIC=1.
 Since 2.26


## `SignalFlags`

The signal flags are used to specify a signal's behaviour, the overall
signal description outlines how especially the RUN flags control the
stages of a signal emission.

C - `GSignalFlags`

### SIGNAL_RUN_FIRST = 1
Invoke the object method handler in the first emission stage.

### SIGNAL_RUN_LAST = 2
Invoke the object method handler in the third emission stage.

### SIGNAL_RUN_CLEANUP = 4
Invoke the object method handler in the last emission stage.

### SIGNAL_NO_RECURSE = 8
Signals being emitted for an object while currently being in
 emission for this very object will not be emitted recursively,
 but instead cause the first emission to be restarted.

### SIGNAL_DETAILED = 16
This signal supports "::detail" appendices to the signal name
 upon handler connections and emissions.

### SIGNAL_ACTION = 32
Action signals are signals that may freely be emitted on alive
 objects from user code via g_signal_emit() and friends, without
 the need of being embedded into extra code that performs pre or
 post emission adjustments on the object. They can also be thought
 of as object methods which can be called generically by
 third-party code.

### SIGNAL_NO_HOOKS = 64
No emissions hooks are supported for this signal.

### SIGNAL_MUST_COLLECT = 128
Varargs signal emission will always collect the
  arguments, even if there are no signal handlers connected.  Since 2.30.

### SIGNAL_DEPRECATED = 256
The signal is deprecated and will be removed
  in a future version. A warning will be generated if it is connected while
  running with G_ENABLE_DIAGNOSTIC=1.  Since 2.32.


## `SignalMatchType`

The match types specify what g_signal_handlers_block_matched(),
g_signal_handlers_unblock_matched() and g_signal_handlers_disconnect_matched()
match signals by.

C - `GSignalMatchType`

### SIGNAL_MATCH_ID = 1
The signal id must be equal.

### SIGNAL_MATCH_DETAIL = 2
The signal detail be equal.

### SIGNAL_MATCH_CLOSURE = 4
The closure must be the same.

### SIGNAL_MATCH_FUNC = 8
The C closure callback must be the same.

### SIGNAL_MATCH_DATA = 16
The closure data must be the same.

### SIGNAL_MATCH_UNBLOCKED = 32
Only unblocked signals may matched.


## `TypeDebugFlags`

These flags used to be passed to g_type_init_with_debug_flags() which
is now deprecated.

If you need to enable debugging features, use the GOBJECT_DEBUG
environment variable.

C - `GTypeDebugFlags`

### TYPE_DEBUG_NONE = 0
Print no messages

### TYPE_DEBUG_OBJECTS = 1
Print messages about object bookkeeping

### TYPE_DEBUG_SIGNALS = 2
Print messages about signal emissions

### TYPE_DEBUG_INSTANCE_COUNT = 4
Keep a count of instances of each type

### TYPE_DEBUG_MASK = 7
Mask covering all debug flags


## `TypeFlags`

Bit masks used to check or determine characteristics of a type.

C - `GTypeFlags`

### TYPE_FLAG_ABSTRACT = 16
Indicates an abstract type. No instances can be
 created for an abstract type

### TYPE_FLAG_VALUE_ABSTRACT = 32
Indicates an abstract value type, i.e. a type
 that introduces a value table, but can't be used for
 g_value_init()


## `TypeFundamentalFlags`

Bit masks used to check or determine specific characteristics of a
fundamental type.

C - `GTypeFundamentalFlags`

### TYPE_FLAG_CLASSED = 1
Indicates a classed type

### TYPE_FLAG_INSTANTIATABLE = 2
Indicates an instantiable type (implies classed)

### TYPE_FLAG_DERIVABLE = 4
Indicates a flat derivable type

### TYPE_FLAG_DEEP_DERIVABLE = 8
Indicates a deep derivable type (implies derivable)


