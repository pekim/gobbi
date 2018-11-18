+++
title = "bitfields"
+++
<p class="api-heading">BindingFlags</p>
<p class="api-doc">Flags to be passed to g_object_bind_property() or
g_object_bind_property_full().

This enumeration can be extended at later date.</p>
<div class="api-notes">
  <p class="api-ctype">GBindingFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">BINDING_DEFAULT</td>
<td class="value">0</td>
<td class="doc">The default binding; if the source property
  changes, the target property is updated with its value.</td>
</tr>
<tr>
<td class="name">BINDING_BIDIRECTIONAL</td>
<td class="value">1</td>
<td class="doc">Bidirectional binding; if either the
  property of the source or the property of the target changes,
  the other is updated.</td>
</tr>
<tr>
<td class="name">BINDING_SYNC_CREATE</td>
<td class="value">2</td>
<td class="doc">Synchronize the values of the source and
  target properties when creating the binding; the direction of
  the synchronization is always from the source to the target.</td>
</tr>
<tr>
<td class="name">BINDING_INVERT_BOOLEAN</td>
<td class="value">4</td>
<td class="doc">If the two properties being bound are
  booleans, setting one to %TRUE will result in the other being
  set to %FALSE and vice versa. This flag will only work for
  boolean properties, and cannot be used when passing custom
  transformation functions to g_object_bind_property_full().</td>
</tr>
</table>
</div>
<p class="api-heading">ConnectFlags</p>
<p class="api-doc">The connection flags are used to specify the behaviour of a signal's
connection.</p>
<div class="api-notes">
  <p class="api-ctype">GConnectFlags</p>
<table>
<tr>
<td class="name">CONNECT_AFTER</td>
<td class="value">1</td>
<td class="doc">whether the handler should be called before or after the
 default handler of the signal.</td>
</tr>
<tr>
<td class="name">CONNECT_SWAPPED</td>
<td class="value">2</td>
<td class="doc">whether the instance and data should be swapped when
 calling the handler; see g_signal_connect_swapped() for an example.</td>
</tr>
</table>
</div>
<p class="api-heading">ParamFlags</p>
<p class="api-doc">Through the #GParamFlags flag values, certain aspects of parameters
can be configured. See also #G_PARAM_STATIC_STRINGS.</p>
<div class="api-notes">
  <p class="api-ctype">GParamFlags</p>
<table>
<tr>
<td class="name">PARAM_READABLE</td>
<td class="value">1</td>
<td class="doc">the parameter is readable</td>
</tr>
<tr>
<td class="name">PARAM_WRITABLE</td>
<td class="value">2</td>
<td class="doc">the parameter is writable</td>
</tr>
<tr>
<td class="name">PARAM_READWRITE</td>
<td class="value">3</td>
<td class="doc">alias for %G_PARAM_READABLE | %G_PARAM_WRITABLE</td>
</tr>
<tr>
<td class="name">PARAM_CONSTRUCT</td>
<td class="value">4</td>
<td class="doc">the parameter will be set upon object construction</td>
</tr>
<tr>
<td class="name">PARAM_CONSTRUCT_ONLY</td>
<td class="value">8</td>
<td class="doc">the parameter can only be set upon object construction</td>
</tr>
<tr>
<td class="name">PARAM_LAX_VALIDATION</td>
<td class="value">16</td>
<td class="doc">upon parameter conversion (see g_param_value_convert())
 strict validation is not required</td>
</tr>
<tr>
<td class="name">PARAM_STATIC_NAME</td>
<td class="value">32</td>
<td class="doc">the string used as name when constructing the
 parameter is guaranteed to remain valid and
 unmodified for the lifetime of the parameter.
 Since 2.8</td>
</tr>
<tr>
<td class="name">PARAM_PRIVATE</td>
<td class="value">32</td>
<td class="doc">internal</td>
</tr>
<tr>
<td class="name">PARAM_STATIC_NICK</td>
<td class="value">64</td>
<td class="doc">the string used as nick when constructing the
 parameter is guaranteed to remain valid and
 unmmodified for the lifetime of the parameter.
 Since 2.8</td>
</tr>
<tr>
<td class="name">PARAM_STATIC_BLURB</td>
<td class="value">128</td>
<td class="doc">the string used as blurb when constructing the
 parameter is guaranteed to remain valid and
 unmodified for the lifetime of the parameter.
 Since 2.8</td>
</tr>
<tr>
<td class="name">PARAM_EXPLICIT_NOTIFY</td>
<td class="value">1073741824</td>
<td class="doc">calls to g_object_set_property() for this
  property will not automatically result in a "notify" signal being
  emitted: the implementation must call g_object_notify() themselves
  in case the property actually changes.  Since: 2.42.</td>
</tr>
<tr>
<td class="name">PARAM_DEPRECATED</td>
<td class="value">2147483648</td>
<td class="doc">the parameter is deprecated and will be removed
 in a future version. A warning will be generated if it is used
 while running with G_ENABLE_DIAGNOSTIC=1.
 Since 2.26</td>
</tr>
</table>
</div>
<p class="api-heading">SignalFlags</p>
<p class="api-doc">The signal flags are used to specify a signal's behaviour, the overall
signal description outlines how especially the RUN flags control the
stages of a signal emission.</p>
<div class="api-notes">
  <p class="api-ctype">GSignalFlags</p>
<table>
<tr>
<td class="name">SIGNAL_RUN_FIRST</td>
<td class="value">1</td>
<td class="doc">Invoke the object method handler in the first emission stage.</td>
</tr>
<tr>
<td class="name">SIGNAL_RUN_LAST</td>
<td class="value">2</td>
<td class="doc">Invoke the object method handler in the third emission stage.</td>
</tr>
<tr>
<td class="name">SIGNAL_RUN_CLEANUP</td>
<td class="value">4</td>
<td class="doc">Invoke the object method handler in the last emission stage.</td>
</tr>
<tr>
<td class="name">SIGNAL_NO_RECURSE</td>
<td class="value">8</td>
<td class="doc">Signals being emitted for an object while currently being in
 emission for this very object will not be emitted recursively,
 but instead cause the first emission to be restarted.</td>
</tr>
<tr>
<td class="name">SIGNAL_DETAILED</td>
<td class="value">16</td>
<td class="doc">This signal supports "::detail" appendices to the signal name
 upon handler connections and emissions.</td>
</tr>
<tr>
<td class="name">SIGNAL_ACTION</td>
<td class="value">32</td>
<td class="doc">Action signals are signals that may freely be emitted on alive
 objects from user code via g_signal_emit() and friends, without
 the need of being embedded into extra code that performs pre or
 post emission adjustments on the object. They can also be thought
 of as object methods which can be called generically by
 third-party code.</td>
</tr>
<tr>
<td class="name">SIGNAL_NO_HOOKS</td>
<td class="value">64</td>
<td class="doc">No emissions hooks are supported for this signal.</td>
</tr>
<tr>
<td class="name">SIGNAL_MUST_COLLECT</td>
<td class="value">128</td>
<td class="doc">Varargs signal emission will always collect the
  arguments, even if there are no signal handlers connected.  Since 2.30.</td>
</tr>
<tr>
<td class="name">SIGNAL_DEPRECATED</td>
<td class="value">256</td>
<td class="doc">The signal is deprecated and will be removed
  in a future version. A warning will be generated if it is connected while
  running with G_ENABLE_DIAGNOSTIC=1.  Since 2.32.</td>
</tr>
</table>
</div>
<p class="api-heading">SignalMatchType</p>
<p class="api-doc">The match types specify what g_signal_handlers_block_matched(),
g_signal_handlers_unblock_matched() and g_signal_handlers_disconnect_matched()
match signals by.</p>
<div class="api-notes">
  <p class="api-ctype">GSignalMatchType</p>
<table>
<tr>
<td class="name">SIGNAL_MATCH_ID</td>
<td class="value">1</td>
<td class="doc">The signal id must be equal.</td>
</tr>
<tr>
<td class="name">SIGNAL_MATCH_DETAIL</td>
<td class="value">2</td>
<td class="doc">The signal detail be equal.</td>
</tr>
<tr>
<td class="name">SIGNAL_MATCH_CLOSURE</td>
<td class="value">4</td>
<td class="doc">The closure must be the same.</td>
</tr>
<tr>
<td class="name">SIGNAL_MATCH_FUNC</td>
<td class="value">8</td>
<td class="doc">The C closure callback must be the same.</td>
</tr>
<tr>
<td class="name">SIGNAL_MATCH_DATA</td>
<td class="value">16</td>
<td class="doc">The closure data must be the same.</td>
</tr>
<tr>
<td class="name">SIGNAL_MATCH_UNBLOCKED</td>
<td class="value">32</td>
<td class="doc">Only unblocked signals may matched.</td>
</tr>
</table>
</div>
<p class="api-heading">TypeDebugFlags</p>
<p class="api-doc">These flags used to be passed to g_type_init_with_debug_flags() which
is now deprecated.

If you need to enable debugging features, use the GOBJECT_DEBUG
environment variable.</p>
<div class="api-notes">
  <p class="api-ctype">GTypeDebugFlags</p>
<table>
<tr>
<td class="name">TYPE_DEBUG_NONE</td>
<td class="value">0</td>
<td class="doc">Print no messages</td>
</tr>
<tr>
<td class="name">TYPE_DEBUG_OBJECTS</td>
<td class="value">1</td>
<td class="doc">Print messages about object bookkeeping</td>
</tr>
<tr>
<td class="name">TYPE_DEBUG_SIGNALS</td>
<td class="value">2</td>
<td class="doc">Print messages about signal emissions</td>
</tr>
<tr>
<td class="name">TYPE_DEBUG_INSTANCE_COUNT</td>
<td class="value">4</td>
<td class="doc">Keep a count of instances of each type</td>
</tr>
<tr>
<td class="name">TYPE_DEBUG_MASK</td>
<td class="value">7</td>
<td class="doc">Mask covering all debug flags</td>
</tr>
</table>
</div>
<p class="api-heading">TypeFlags</p>
<p class="api-doc">Bit masks used to check or determine characteristics of a type.</p>
<div class="api-notes">
  <p class="api-ctype">GTypeFlags</p>
<table>
<tr>
<td class="name">TYPE_FLAG_ABSTRACT</td>
<td class="value">16</td>
<td class="doc">Indicates an abstract type. No instances can be
 created for an abstract type</td>
</tr>
<tr>
<td class="name">TYPE_FLAG_VALUE_ABSTRACT</td>
<td class="value">32</td>
<td class="doc">Indicates an abstract value type, i.e. a type
 that introduces a value table, but can't be used for
 g_value_init()</td>
</tr>
</table>
</div>
<p class="api-heading">TypeFundamentalFlags</p>
<p class="api-doc">Bit masks used to check or determine specific characteristics of a
fundamental type.</p>
<div class="api-notes">
  <p class="api-ctype">GTypeFundamentalFlags</p>
<table>
<tr>
<td class="name">TYPE_FLAG_CLASSED</td>
<td class="value">1</td>
<td class="doc">Indicates a classed type</td>
</tr>
<tr>
<td class="name">TYPE_FLAG_INSTANTIATABLE</td>
<td class="value">2</td>
<td class="doc">Indicates an instantiable type (implies classed)</td>
</tr>
<tr>
<td class="name">TYPE_FLAG_DERIVABLE</td>
<td class="value">4</td>
<td class="doc">Indicates a flat derivable type</td>
</tr>
<tr>
<td class="name">TYPE_FLAG_DEEP_DERIVABLE</td>
<td class="value">8</td>
<td class="doc">Indicates a deep derivable type (implies derivable)</td>
</tr>
</table>
</div>
