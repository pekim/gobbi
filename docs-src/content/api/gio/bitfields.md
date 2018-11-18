+++
title = "bitfields"
+++
<p class="api-heading">AppInfoCreateFlags</p>
<p class="api-doc">Flags used when creating a #GAppInfo.</p>
<div class="api-notes">
  <p class="api-ctype">GAppInfoCreateFlags</p>
<table>
<tr>
<td class="name">APP_INFO_CREATE_NONE</td>
<td class="value">0</td>
<td class="doc">No flags.</td>
</tr>
<tr>
<td class="name">APP_INFO_CREATE_NEEDS_TERMINAL</td>
<td class="value">1</td>
<td class="doc">Application opens in a terminal window.</td>
</tr>
<tr>
<td class="name">APP_INFO_CREATE_SUPPORTS_URIS</td>
<td class="value">2</td>
<td class="doc">Application supports URI arguments.</td>
</tr>
<tr>
<td class="name">APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION</td>
<td class="value">4</td>
<td class="doc">Application supports startup notification. Since 2.26</td>
</tr>
</table>
</div>
<p class="api-heading">ApplicationFlags</p>
<p class="api-doc">Flags used to define the behaviour of a #GApplication.</p>
<div class="api-notes">
  <p class="api-ctype">GApplicationFlags</p>
  <p class="api-since">since 2.28</p>
<table>
<tr>
<td class="name">APPLICATION_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">Default</td>
</tr>
<tr>
<td class="name">APPLICATION_IS_SERVICE</td>
<td class="value">1</td>
<td class="doc">Run as a service. In this mode, registration
     fails if the service is already running, and the application
     will initially wait up to 10 seconds for an initial activation
     message to arrive.</td>
</tr>
<tr>
<td class="name">APPLICATION_IS_LAUNCHER</td>
<td class="value">2</td>
<td class="doc">Don't try to become the primary instance.</td>
</tr>
<tr>
<td class="name">APPLICATION_HANDLES_OPEN</td>
<td class="value">4</td>
<td class="doc">This application handles opening files (in
    the primary instance). Note that this flag only affects the default
    implementation of local_command_line(), and has no effect if
    %G_APPLICATION_HANDLES_COMMAND_LINE is given.
    See g_application_run() for details.</td>
</tr>
<tr>
<td class="name">APPLICATION_HANDLES_COMMAND_LINE</td>
<td class="value">8</td>
<td class="doc">This application handles command line
    arguments (in the primary instance). Note that this flag only affect
    the default implementation of local_command_line().
    See g_application_run() for details.</td>
</tr>
<tr>
<td class="name">APPLICATION_SEND_ENVIRONMENT</td>
<td class="value">16</td>
<td class="doc">Send the environment of the
    launching process to the primary instance. Set this flag if your
    application is expected to behave differently depending on certain
    environment variables. For instance, an editor might be expected
    to use the `GIT_COMMITTER_NAME` environment variable
    when editing a git commit message. The environment is available
    to the #GApplication::command-line signal handler, via
    g_application_command_line_getenv().</td>
</tr>
<tr>
<td class="name">APPLICATION_NON_UNIQUE</td>
<td class="value">32</td>
<td class="doc">Make no attempts to do any of the typical
    single-instance application negotiation, even if the application
    ID is given.  The application neither attempts to become the
    owner of the application ID nor does it check if an existing
    owner already exists.  Everything occurs in the local process.
    Since: 2.30.</td>
</tr>
<tr>
<td class="name">APPLICATION_CAN_OVERRIDE_APP_ID</td>
<td class="value">64</td>
<td class="doc">Allow users to override the
    application ID from the command line with `--gapplication-app-id`.
    Since: 2.48</td>
</tr>
</table>
</div>
<p class="api-heading">AskPasswordFlags</p>
<p class="api-doc">#GAskPasswordFlags are used to request specific information from the
user, or to notify the user of their choices in an authentication
situation.</p>
<div class="api-notes">
  <p class="api-ctype">GAskPasswordFlags</p>
<table>
<tr>
<td class="name">ASK_PASSWORD_NEED_PASSWORD</td>
<td class="value">1</td>
<td class="doc">operation requires a password.</td>
</tr>
<tr>
<td class="name">ASK_PASSWORD_NEED_USERNAME</td>
<td class="value">2</td>
<td class="doc">operation requires a username.</td>
</tr>
<tr>
<td class="name">ASK_PASSWORD_NEED_DOMAIN</td>
<td class="value">4</td>
<td class="doc">operation requires a domain.</td>
</tr>
<tr>
<td class="name">ASK_PASSWORD_SAVING_SUPPORTED</td>
<td class="value">8</td>
<td class="doc">operation supports saving settings.</td>
</tr>
<tr>
<td class="name">ASK_PASSWORD_ANONYMOUS_SUPPORTED</td>
<td class="value">16</td>
<td class="doc">operation supports anonymous users.</td>
</tr>
</table>
</div>
<p class="api-heading">BusNameOwnerFlags</p>
<p class="api-doc">Flags used in g_bus_own_name().</p>
<div class="api-notes">
  <p class="api-ctype">GBusNameOwnerFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">BUS_NAME_OWNER_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT</td>
<td class="value">1</td>
<td class="doc">Allow another message bus connection to claim the name.</td>
</tr>
<tr>
<td class="name">BUS_NAME_OWNER_FLAGS_REPLACE</td>
<td class="value">2</td>
<td class="doc">If another message bus connection owns the name and have
specified #G_BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT, then take the name from the other connection.</td>
</tr>
<tr>
<td class="name">BUS_NAME_OWNER_FLAGS_DO_NOT_QUEUE</td>
<td class="value">4</td>
<td class="doc">If another message bus connection owns the name, immediately
return an error from g_bus_own_name() rather than entering the waiting queue for that name. (Since 2.54)</td>
</tr>
</table>
</div>
<p class="api-heading">BusNameWatcherFlags</p>
<p class="api-doc">Flags used in g_bus_watch_name().</p>
<div class="api-notes">
  <p class="api-ctype">GBusNameWatcherFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">BUS_NAME_WATCHER_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">BUS_NAME_WATCHER_FLAGS_AUTO_START</td>
<td class="value">1</td>
<td class="doc">If no-one owns the name when
beginning to watch the name, ask the bus to launch an owner for the
name.</td>
</tr>
</table>
</div>
<p class="api-heading">ConverterFlags</p>
<p class="api-doc">Flags used when calling a g_converter_convert().</p>
<div class="api-notes">
  <p class="api-ctype">GConverterFlags</p>
  <p class="api-since">since 2.24</p>
<table>
<tr>
<td class="name">CONVERTER_NO_FLAGS</td>
<td class="value">0</td>
<td class="doc">No flags.</td>
</tr>
<tr>
<td class="name">CONVERTER_INPUT_AT_END</td>
<td class="value">1</td>
<td class="doc">At end of input data</td>
</tr>
<tr>
<td class="name">CONVERTER_FLUSH</td>
<td class="value">2</td>
<td class="doc">Flush data</td>
</tr>
</table>
</div>
<p class="api-heading">DBusCallFlags</p>
<p class="api-doc">Flags used in g_dbus_connection_call() and similar APIs.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusCallFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_CALL_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_CALL_FLAGS_NO_AUTO_START</td>
<td class="value">1</td>
<td class="doc">The bus must not launch
an owner for the destination name in response to this method
invocation.</td>
</tr>
<tr>
<td class="name">DBUS_CALL_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION</td>
<td class="value">2</td>
<td class="doc">the caller is prepared to
wait for interactive authorization. Since 2.46.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusCapabilityFlags</p>
<p class="api-doc">Capabilities negotiated with the remote peer.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusCapabilityFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_CAPABILITY_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING</td>
<td class="value">1</td>
<td class="doc">The connection
supports exchanging UNIX file descriptors with the remote peer.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusConnectionFlags</p>
<p class="api-doc">Flags used when creating a new #GDBusConnection.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusConnectionFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_CONNECTION_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT</td>
<td class="value">1</td>
<td class="doc">Perform authentication against server.</td>
</tr>
<tr>
<td class="name">DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER</td>
<td class="value">2</td>
<td class="doc">Perform authentication against client.</td>
</tr>
<tr>
<td class="name">DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS</td>
<td class="value">4</td>
<td class="doc">When
authenticating as a server, allow the anonymous authentication
method.</td>
</tr>
<tr>
<td class="name">DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION</td>
<td class="value">8</td>
<td class="doc">Pass this flag if connecting to a peer that is a
message bus. This means that the Hello() method will be invoked as part of the connection setup.</td>
</tr>
<tr>
<td class="name">DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING</td>
<td class="value">16</td>
<td class="doc">If set, processing of D-Bus messages is
delayed until g_dbus_connection_start_message_processing() is called.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusInterfaceSkeletonFlags</p>
<p class="api-doc">Flags describing the behavior of a #GDBusInterfaceSkeleton instance.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusInterfaceSkeletonFlags</p>
  <p class="api-since">since 2.30</p>
<table>
<tr>
<td class="name">DBUS_INTERFACE_SKELETON_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD</td>
<td class="value">1</td>
<td class="doc">Each method invocation is handled in
  a thread dedicated to the invocation. This means that the method implementation can use blocking IO
  without blocking any other part of the process. It also means that the method implementation must
  use locking to access data structures used by other threads.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusMessageFlags</p>
<p class="api-doc">Message flags used in #GDBusMessage.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusMessageFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_MESSAGE_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED</td>
<td class="value">1</td>
<td class="doc">A reply is not expected.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_FLAGS_NO_AUTO_START</td>
<td class="value">2</td>
<td class="doc">The bus must not launch an
owner for the destination name in response to this message.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION</td>
<td class="value">4</td>
<td class="doc">If set on a method
call, this flag means that the caller is prepared to wait for interactive
authorization. Since 2.46.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusObjectManagerClientFlags</p>
<p class="api-doc">Flags used when constructing a #GDBusObjectManagerClient.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusObjectManagerClientFlags</p>
  <p class="api-since">since 2.30</p>
<table>
<tr>
<td class="name">DBUS_OBJECT_MANAGER_CLIENT_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START</td>
<td class="value">1</td>
<td class="doc">If not set and the
  manager is for a well-known name, then request the bus to launch
  an owner for the name if no-one owns the name. This flag can only
  be used in managers for well-known names.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusPropertyInfoFlags</p>
<p class="api-doc">Flags describing the access control of a D-Bus property.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusPropertyInfoFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_PROPERTY_INFO_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_PROPERTY_INFO_FLAGS_READABLE</td>
<td class="value">1</td>
<td class="doc">Property is readable.</td>
</tr>
<tr>
<td class="name">DBUS_PROPERTY_INFO_FLAGS_WRITABLE</td>
<td class="value">2</td>
<td class="doc">Property is writable.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusProxyFlags</p>
<p class="api-doc">Flags used when constructing an instance of a #GDBusProxy derived class.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusProxyFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_PROXY_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES</td>
<td class="value">1</td>
<td class="doc">Don't load properties.</td>
</tr>
<tr>
<td class="name">DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS</td>
<td class="value">2</td>
<td class="doc">Don't connect to signals on the remote object.</td>
</tr>
<tr>
<td class="name">DBUS_PROXY_FLAGS_DO_NOT_AUTO_START</td>
<td class="value">4</td>
<td class="doc">If the proxy is for a well-known name,
do not ask the bus to launch an owner during proxy initialization or a method call.
This flag is only meaningful in proxies for well-known names.</td>
</tr>
<tr>
<td class="name">DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES</td>
<td class="value">8</td>
<td class="doc">If set, the property value for any __invalidated property__ will be (asynchronously) retrieved upon receiving the [`PropertiesChanged`](http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-properties) D-Bus signal and the property will not cause emission of the #GDBusProxy::g-properties-changed signal. When the value is received the #GDBusProxy::g-properties-changed signal is emitted for the property along with the retrieved value. Since 2.32.</td>
</tr>
<tr>
<td class="name">DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION</td>
<td class="value">16</td>
<td class="doc">If the proxy is for a well-known name,
do not ask the bus to launch an owner during proxy initialization, but allow it to be
autostarted by a method call. This flag is only meaningful in proxies for well-known names,
and only if %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START is not also specified.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusSendMessageFlags</p>
<p class="api-doc">Flags used when sending #GDBusMessages on a #GDBusConnection.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusSendMessageFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_SEND_MESSAGE_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL</td>
<td class="value">1</td>
<td class="doc">Do not automatically
assign a serial number from the #GDBusConnection object when
sending a message.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusServerFlags</p>
<p class="api-doc">Flags used when creating a #GDBusServer.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusServerFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_SERVER_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_SERVER_FLAGS_RUN_IN_THREAD</td>
<td class="value">1</td>
<td class="doc">All #GDBusServer::new-connection
signals will run in separated dedicated threads (see signal for
details).</td>
</tr>
<tr>
<td class="name">DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS</td>
<td class="value">2</td>
<td class="doc">Allow the anonymous
authentication method.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusSignalFlags</p>
<p class="api-doc">Flags used when subscribing to signals via g_dbus_connection_signal_subscribe().</p>
<div class="api-notes">
  <p class="api-ctype">GDBusSignalFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_SIGNAL_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_SIGNAL_FLAGS_NO_MATCH_RULE</td>
<td class="value">1</td>
<td class="doc">Don't actually send the AddMatch
D-Bus call for this signal subscription.  This gives you more control
over which match rules you add (but you must add them manually).</td>
</tr>
<tr>
<td class="name">DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE</td>
<td class="value">2</td>
<td class="doc">Match first arguments that
contain a bus or interface name with the given namespace.</td>
</tr>
<tr>
<td class="name">DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH</td>
<td class="value">4</td>
<td class="doc">Match first arguments that
contain an object path that is either equivalent to the given path,
or one of the paths is a subpath of the other.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusSubtreeFlags</p>
<p class="api-doc">Flags passed to g_dbus_connection_register_subtree().</p>
<div class="api-notes">
  <p class="api-ctype">GDBusSubtreeFlags</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_SUBTREE_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES</td>
<td class="value">1</td>
<td class="doc">Method calls to objects not in the enumerated range
                                                      will still be dispatched. This is useful if you want
                                                      to dynamically spawn objects in the subtree.</td>
</tr>
</table>
</div>
<p class="api-heading">DriveStartFlags</p>
<p class="api-doc">Flags used when starting a drive.</p>
<div class="api-notes">
  <p class="api-ctype">GDriveStartFlags</p>
  <p class="api-since">since 2.22</p>
<table>
<tr>
<td class="name">DRIVE_START_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
</table>
</div>
<p class="api-heading">FileAttributeInfoFlags</p>
<p class="api-doc">Flags specifying the behaviour of an attribute.</p>
<div class="api-notes">
  <p class="api-ctype">GFileAttributeInfoFlags</p>
<table>
<tr>
<td class="name">FILE_ATTRIBUTE_INFO_NONE</td>
<td class="value">0</td>
<td class="doc">no flags set.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_INFO_COPY_WITH_FILE</td>
<td class="value">1</td>
<td class="doc">copy the attribute values when the file is copied.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED</td>
<td class="value">2</td>
<td class="doc">copy the attribute values when the file is moved.</td>
</tr>
</table>
</div>
<p class="api-heading">FileCopyFlags</p>
<p class="api-doc">Flags used when copying or moving files.</p>
<div class="api-notes">
  <p class="api-ctype">GFileCopyFlags</p>
<table>
<tr>
<td class="name">FILE_COPY_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">FILE_COPY_OVERWRITE</td>
<td class="value">1</td>
<td class="doc">Overwrite any existing files</td>
</tr>
<tr>
<td class="name">FILE_COPY_BACKUP</td>
<td class="value">2</td>
<td class="doc">Make a backup of any existing files.</td>
</tr>
<tr>
<td class="name">FILE_COPY_NOFOLLOW_SYMLINKS</td>
<td class="value">4</td>
<td class="doc">Don't follow symlinks.</td>
</tr>
<tr>
<td class="name">FILE_COPY_ALL_METADATA</td>
<td class="value">8</td>
<td class="doc">Copy all file metadata instead of just default set used for copy (see #GFileInfo).</td>
</tr>
<tr>
<td class="name">FILE_COPY_NO_FALLBACK_FOR_MOVE</td>
<td class="value">16</td>
<td class="doc">Don't use copy and delete fallback if native move not supported.</td>
</tr>
<tr>
<td class="name">FILE_COPY_TARGET_DEFAULT_PERMS</td>
<td class="value">32</td>
<td class="doc">Leaves target file with default perms, instead of setting the source file perms.</td>
</tr>
</table>
</div>
<p class="api-heading">FileCreateFlags</p>
<p class="api-doc">Flags used when an operation may create a file.</p>
<div class="api-notes">
  <p class="api-ctype">GFileCreateFlags</p>
<table>
<tr>
<td class="name">FILE_CREATE_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">FILE_CREATE_PRIVATE</td>
<td class="value">1</td>
<td class="doc">Create a file that can only be
   accessed by the current user.</td>
</tr>
<tr>
<td class="name">FILE_CREATE_REPLACE_DESTINATION</td>
<td class="value">2</td>
<td class="doc">Replace the destination
   as if it didn't exist before. Don't try to keep any old
   permissions, replace instead of following links. This
   is generally useful if you're doing a "copy over"
   rather than a "save new version of" replace operation.
   You can think of it as "unlink destination" before
   writing to it, although the implementation may not
   be exactly like that. Since 2.20</td>
</tr>
</table>
</div>
<p class="api-heading">FileMeasureFlags</p>
<p class="api-doc">Flags that can be used with g_file_measure_disk_usage().</p>
<div class="api-notes">
  <p class="api-ctype">GFileMeasureFlags</p>
  <p class="api-since">since 2.38</p>
<table>
<tr>
<td class="name">FILE_MEASURE_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">FILE_MEASURE_REPORT_ANY_ERROR</td>
<td class="value">2</td>
<td class="doc">Report any error encountered
  while traversing the directory tree.  Normally errors are only
  reported for the toplevel file.</td>
</tr>
<tr>
<td class="name">FILE_MEASURE_APPARENT_SIZE</td>
<td class="value">4</td>
<td class="doc">Tally usage based on apparent file
  sizes.  Normally, the block-size is used, if available, as this is a
  more accurate representation of disk space used.
  Compare with `du --apparent-size`.</td>
</tr>
<tr>
<td class="name">FILE_MEASURE_NO_XDEV</td>
<td class="value">8</td>
<td class="doc">Do not cross mount point boundaries.
  Compare with `du -x`.</td>
</tr>
</table>
</div>
<p class="api-heading">FileMonitorFlags</p>
<p class="api-doc">Flags used to set what a #GFileMonitor will watch for.</p>
<div class="api-notes">
  <p class="api-ctype">GFileMonitorFlags</p>
<table>
<tr>
<td class="name">FILE_MONITOR_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_WATCH_MOUNTS</td>
<td class="value">1</td>
<td class="doc">Watch for mount events.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_SEND_MOVED</td>
<td class="value">2</td>
<td class="doc">Pair DELETED and CREATED events caused
  by file renames (moves) and send a single G_FILE_MONITOR_EVENT_MOVED
  event instead (NB: not supported on all backends; the default
  behaviour -without specifying this flag- is to send single DELETED
  and CREATED events).  Deprecated since 2.46: use
  %G_FILE_MONITOR_WATCH_MOVES instead.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_WATCH_HARD_LINKS</td>
<td class="value">4</td>
<td class="doc">Watch for changes to the file made
  via another hard link. Since 2.36.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_WATCH_MOVES</td>
<td class="value">8</td>
<td class="doc">Watch for rename operations on a
  monitored directory.  This causes %G_FILE_MONITOR_EVENT_RENAMED,
  %G_FILE_MONITOR_EVENT_MOVED_IN and %G_FILE_MONITOR_EVENT_MOVED_OUT
  events to be emitted when possible.  Since: 2.46.</td>
</tr>
</table>
</div>
<p class="api-heading">FileQueryInfoFlags</p>
<p class="api-doc">Flags used when querying a #GFileInfo.</p>
<div class="api-notes">
  <p class="api-ctype">GFileQueryInfoFlags</p>
<table>
<tr>
<td class="name">FILE_QUERY_INFO_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">FILE_QUERY_INFO_NOFOLLOW_SYMLINKS</td>
<td class="value">1</td>
<td class="doc">Don't follow symlinks.</td>
</tr>
</table>
</div>
<p class="api-heading">IOStreamSpliceFlags</p>
<p class="api-doc">GIOStreamSpliceFlags determine how streams should be spliced.</p>
<div class="api-notes">
  <p class="api-ctype">GIOStreamSpliceFlags</p>
  <p class="api-since">since 2.28</p>
<table>
<tr>
<td class="name">IO_STREAM_SPLICE_NONE</td>
<td class="value">0</td>
<td class="doc">Do not close either stream.</td>
</tr>
<tr>
<td class="name">IO_STREAM_SPLICE_CLOSE_STREAM1</td>
<td class="value">1</td>
<td class="doc">Close the first stream after
    the splice.</td>
</tr>
<tr>
<td class="name">IO_STREAM_SPLICE_CLOSE_STREAM2</td>
<td class="value">2</td>
<td class="doc">Close the second stream after
    the splice.</td>
</tr>
<tr>
<td class="name">IO_STREAM_SPLICE_WAIT_FOR_BOTH</td>
<td class="value">4</td>
<td class="doc">Wait for both splice operations to finish
    before calling the callback.</td>
</tr>
</table>
</div>
<p class="api-heading">MountMountFlags</p>
<p class="api-doc">Flags used when mounting a mount.</p>
<div class="api-notes">
  <p class="api-ctype">GMountMountFlags</p>
<table>
<tr>
<td class="name">MOUNT_MOUNT_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
</table>
</div>
<p class="api-heading">MountUnmountFlags</p>
<p class="api-doc">Flags used when an unmounting a mount.</p>
<div class="api-notes">
  <p class="api-ctype">GMountUnmountFlags</p>
<table>
<tr>
<td class="name">MOUNT_UNMOUNT_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">MOUNT_UNMOUNT_FORCE</td>
<td class="value">1</td>
<td class="doc">Unmount even if there are outstanding
 file operations on the mount.</td>
</tr>
</table>
</div>
<p class="api-heading">OutputStreamSpliceFlags</p>
<p class="api-doc">GOutputStreamSpliceFlags determine how streams should be spliced.</p>
<div class="api-notes">
  <p class="api-ctype">GOutputStreamSpliceFlags</p>
<table>
<tr>
<td class="name">OUTPUT_STREAM_SPLICE_NONE</td>
<td class="value">0</td>
<td class="doc">Do not close either stream.</td>
</tr>
<tr>
<td class="name">OUTPUT_STREAM_SPLICE_CLOSE_SOURCE</td>
<td class="value">1</td>
<td class="doc">Close the source stream after
    the splice.</td>
</tr>
<tr>
<td class="name">OUTPUT_STREAM_SPLICE_CLOSE_TARGET</td>
<td class="value">2</td>
<td class="doc">Close the target stream after
    the splice.</td>
</tr>
</table>
</div>
<p class="api-heading">ResourceFlags</p>
<p class="api-doc">GResourceFlags give information about a particular file inside a resource
bundle.</p>
<div class="api-notes">
  <p class="api-ctype">GResourceFlags</p>
  <p class="api-since">since 2.32</p>
<table>
<tr>
<td class="name">RESOURCE_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
<tr>
<td class="name">RESOURCE_FLAGS_COMPRESSED</td>
<td class="value">1</td>
<td class="doc">The file is compressed.</td>
</tr>
</table>
</div>
<p class="api-heading">ResourceLookupFlags</p>
<p class="api-doc">GResourceLookupFlags determine how resource path lookups are handled.</p>
<div class="api-notes">
  <p class="api-ctype">GResourceLookupFlags</p>
  <p class="api-since">since 2.32</p>
<table>
<tr>
<td class="name">RESOURCE_LOOKUP_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags set.</td>
</tr>
</table>
</div>
<p class="api-heading">SettingsBindFlags</p>
<p class="api-doc">Flags used when creating a binding. These flags determine in which
direction the binding works. The default is to synchronize in both
directions.</p>
<div class="api-notes">
  <p class="api-ctype">GSettingsBindFlags</p>
<table>
<tr>
<td class="name">SETTINGS_BIND_DEFAULT</td>
<td class="value">0</td>
<td class="doc">Equivalent to `G_SETTINGS_BIND_GET|G_SETTINGS_BIND_SET`</td>
</tr>
<tr>
<td class="name">SETTINGS_BIND_GET</td>
<td class="value">1</td>
<td class="doc">Update the #GObject property when the setting changes.
    It is an error to use this flag if the property is not writable.</td>
</tr>
<tr>
<td class="name">SETTINGS_BIND_SET</td>
<td class="value">2</td>
<td class="doc">Update the setting when the #GObject property changes.
    It is an error to use this flag if the property is not readable.</td>
</tr>
<tr>
<td class="name">SETTINGS_BIND_NO_SENSITIVITY</td>
<td class="value">4</td>
<td class="doc">Do not try to bind a "sensitivity" property to the writability of the setting</td>
</tr>
<tr>
<td class="name">SETTINGS_BIND_GET_NO_CHANGES</td>
<td class="value">8</td>
<td class="doc">When set in addition to #G_SETTINGS_BIND_GET, set the #GObject property
    value initially from the setting, but do not listen for changes of the setting</td>
</tr>
<tr>
<td class="name">SETTINGS_BIND_INVERT_BOOLEAN</td>
<td class="value">16</td>
<td class="doc">When passed to g_settings_bind(), uses a pair of mapping functions that invert
    the boolean value when mapping between the setting and the property.  The setting and property must both
    be booleans.  You cannot pass this flag to g_settings_bind_with_mapping().</td>
</tr>
</table>
</div>
<p class="api-heading">SocketMsgFlags</p>
<p class="api-doc">Flags used in g_socket_receive_message() and g_socket_send_message().
The flags listed in the enum are some commonly available flags, but the
values used for them are the same as on the platform, and any other flags
are passed in/out as is. So to use a platform specific flag, just include
the right system header and pass in the flag.</p>
<div class="api-notes">
  <p class="api-ctype">GSocketMsgFlags</p>
  <p class="api-since">since 2.22</p>
<table>
<tr>
<td class="name">SOCKET_MSG_NONE</td>
<td class="value">0</td>
<td class="doc">No flags.</td>
</tr>
<tr>
<td class="name">SOCKET_MSG_OOB</td>
<td class="value">1</td>
<td class="doc">Request to send/receive out of band data.</td>
</tr>
<tr>
<td class="name">SOCKET_MSG_PEEK</td>
<td class="value">2</td>
<td class="doc">Read data from the socket without removing it from
    the queue.</td>
</tr>
<tr>
<td class="name">SOCKET_MSG_DONTROUTE</td>
<td class="value">4</td>
<td class="doc">Don't use a gateway to send out the packet,
    only send to hosts on directly connected networks.</td>
</tr>
</table>
</div>
<p class="api-heading">SubprocessFlags</p>
<p class="api-doc">Flags to define the behaviour of a #GSubprocess.

Note that the default for stdin is to redirect from `/dev/null`.  For
stdout and stderr the default are for them to inherit the
corresponding descriptor from the calling process.

Note that it is a programmer error to mix 'incompatible' flags.  For
example, you may not request both %G_SUBPROCESS_FLAGS_STDOUT_PIPE and
%G_SUBPROCESS_FLAGS_STDOUT_SILENCE.</p>
<div class="api-notes">
  <p class="api-ctype">GSubprocessFlags</p>
  <p class="api-since">since 2.40</p>
<table>
<tr>
<td class="name">SUBPROCESS_FLAGS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags.</td>
</tr>
<tr>
<td class="name">SUBPROCESS_FLAGS_STDIN_PIPE</td>
<td class="value">1</td>
<td class="doc">create a pipe for the stdin of the
  spawned process that can be accessed with
  g_subprocess_get_stdin_pipe().</td>
</tr>
<tr>
<td class="name">SUBPROCESS_FLAGS_STDIN_INHERIT</td>
<td class="value">2</td>
<td class="doc">stdin is inherited from the
  calling process.</td>
</tr>
<tr>
<td class="name">SUBPROCESS_FLAGS_STDOUT_PIPE</td>
<td class="value">4</td>
<td class="doc">create a pipe for the stdout of the
  spawned process that can be accessed with
  g_subprocess_get_stdout_pipe().</td>
</tr>
<tr>
<td class="name">SUBPROCESS_FLAGS_STDOUT_SILENCE</td>
<td class="value">8</td>
<td class="doc">silence the stdout of the spawned
  process (ie: redirect to `/dev/null`).</td>
</tr>
<tr>
<td class="name">SUBPROCESS_FLAGS_STDERR_PIPE</td>
<td class="value">16</td>
<td class="doc">create a pipe for the stderr of the
  spawned process that can be accessed with
  g_subprocess_get_stderr_pipe().</td>
</tr>
<tr>
<td class="name">SUBPROCESS_FLAGS_STDERR_SILENCE</td>
<td class="value">32</td>
<td class="doc">silence the stderr of the spawned
  process (ie: redirect to `/dev/null`).</td>
</tr>
<tr>
<td class="name">SUBPROCESS_FLAGS_STDERR_MERGE</td>
<td class="value">64</td>
<td class="doc">merge the stderr of the spawned
  process with whatever the stdout happens to be.  This is a good way
  of directing both streams to a common log file, for example.</td>
</tr>
<tr>
<td class="name">SUBPROCESS_FLAGS_INHERIT_FDS</td>
<td class="value">128</td>
<td class="doc">spawned processes will inherit the
  file descriptors of their parent, unless those descriptors have
  been explicitly marked as close-on-exec.  This flag has no effect
  over the "standard" file descriptors (stdin, stdout, stderr).</td>
</tr>
</table>
</div>
<p class="api-heading">TestDBusFlags</p>
<p class="api-doc">Flags to define future #GTestDBus behaviour.</p>
<div class="api-notes">
  <p class="api-ctype">GTestDBusFlags</p>
  <p class="api-since">since 2.34</p>
<table>
<tr>
<td class="name">TEST_DBUS_NONE</td>
<td class="value">0</td>
<td class="doc">No flags.</td>
</tr>
</table>
</div>
<p class="api-heading">TlsCertificateFlags</p>
<p class="api-doc">A set of flags describing TLS certification validation. This can be
used to set which validation steps to perform (eg, with
g_tls_client_connection_set_validation_flags()), or to describe why
a particular certificate was rejected (eg, in
#GTlsConnection::accept-certificate).</p>
<div class="api-notes">
  <p class="api-ctype">GTlsCertificateFlags</p>
  <p class="api-since">since 2.28</p>
<table>
<tr>
<td class="name">TLS_CERTIFICATE_UNKNOWN_CA</td>
<td class="value">1</td>
<td class="doc">The signing certificate authority is
  not known.</td>
</tr>
<tr>
<td class="name">TLS_CERTIFICATE_BAD_IDENTITY</td>
<td class="value">2</td>
<td class="doc">The certificate does not match the
  expected identity of the site that it was retrieved from.</td>
</tr>
<tr>
<td class="name">TLS_CERTIFICATE_NOT_ACTIVATED</td>
<td class="value">4</td>
<td class="doc">The certificate's activation time
  is still in the future</td>
</tr>
<tr>
<td class="name">TLS_CERTIFICATE_EXPIRED</td>
<td class="value">8</td>
<td class="doc">The certificate has expired</td>
</tr>
<tr>
<td class="name">TLS_CERTIFICATE_REVOKED</td>
<td class="value">16</td>
<td class="doc">The certificate has been revoked
  according to the #GTlsConnection's certificate revocation list.</td>
</tr>
<tr>
<td class="name">TLS_CERTIFICATE_INSECURE</td>
<td class="value">32</td>
<td class="doc">The certificate's algorithm is
  considered insecure.</td>
</tr>
<tr>
<td class="name">TLS_CERTIFICATE_GENERIC_ERROR</td>
<td class="value">64</td>
<td class="doc">Some other error occurred validating
  the certificate</td>
</tr>
<tr>
<td class="name">TLS_CERTIFICATE_VALIDATE_ALL</td>
<td class="value">127</td>
<td class="doc">the combination of all of the above
  flags</td>
</tr>
</table>
</div>
<p class="api-heading">TlsDatabaseVerifyFlags</p>
<p class="api-doc">Flags for g_tls_database_verify_chain().</p>
<div class="api-notes">
  <p class="api-ctype">GTlsDatabaseVerifyFlags</p>
  <p class="api-since">since 2.30</p>
<table>
<tr>
<td class="name">TLS_DATABASE_VERIFY_NONE</td>
<td class="value">0</td>
<td class="doc">No verification flags</td>
</tr>
</table>
</div>
<p class="api-heading">TlsPasswordFlags</p>
<p class="api-doc">Various flags for the password.</p>
<div class="api-notes">
  <p class="api-ctype">GTlsPasswordFlags</p>
  <p class="api-since">since 2.30</p>
<table>
<tr>
<td class="name">TLS_PASSWORD_NONE</td>
<td class="value">0</td>
<td class="doc">No flags</td>
</tr>
<tr>
<td class="name">TLS_PASSWORD_RETRY</td>
<td class="value">2</td>
<td class="doc">The password was wrong, and the user should retry.</td>
</tr>
<tr>
<td class="name">TLS_PASSWORD_MANY_TRIES</td>
<td class="value">4</td>
<td class="doc">Hint to the user that the password has been
   wrong many times, and the user may not have many chances left.</td>
</tr>
<tr>
<td class="name">TLS_PASSWORD_FINAL_TRY</td>
<td class="value">8</td>
<td class="doc">Hint to the user that this is the last try to get
   this password right.</td>
</tr>
</table>
</div>