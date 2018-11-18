+++
title = "enums"
+++
<p class="api-heading">BusType</p>
<p class="api-doc">An enumeration for well-known message buses.</p>
<div class="api-notes">
  <p class="api-ctype">GBusType</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">BUS_TYPE_STARTER</td>
<td class="value">-1</td>
<td class="doc">An alias for the message bus that activated the process, if any.</td>
</tr>
<tr>
<td class="name">BUS_TYPE_NONE</td>
<td class="value">0</td>
<td class="doc">Not a message bus.</td>
</tr>
<tr>
<td class="name">BUS_TYPE_SYSTEM</td>
<td class="value">1</td>
<td class="doc">The system-wide message bus.</td>
</tr>
<tr>
<td class="name">BUS_TYPE_SESSION</td>
<td class="value">2</td>
<td class="doc">The login session message bus.</td>
</tr>
</table>
</div>
<p class="api-heading">ConverterResult</p>
<p class="api-doc">Results returned from g_converter_convert().</p>
<div class="api-notes">
  <p class="api-ctype">GConverterResult</p>
  <p class="api-since">since 2.24</p>
<table>
<tr>
<td class="name">CONVERTER_ERROR</td>
<td class="value">0</td>
<td class="doc">There was an error during conversion.</td>
</tr>
<tr>
<td class="name">CONVERTER_CONVERTED</td>
<td class="value">1</td>
<td class="doc">Some data was consumed or produced</td>
</tr>
<tr>
<td class="name">CONVERTER_FINISHED</td>
<td class="value">2</td>
<td class="doc">The conversion is finished</td>
</tr>
<tr>
<td class="name">CONVERTER_FLUSHED</td>
<td class="value">3</td>
<td class="doc">Flushing is finished</td>
</tr>
</table>
</div>
<p class="api-heading">CredentialsType</p>
<p class="api-doc">Enumeration describing different kinds of native credential types.</p>
<div class="api-notes">
  <p class="api-ctype">GCredentialsType</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">CREDENTIALS_TYPE_INVALID</td>
<td class="value">0</td>
<td class="doc">Indicates an invalid native credential type.</td>
</tr>
<tr>
<td class="name">CREDENTIALS_TYPE_LINUX_UCRED</td>
<td class="value">1</td>
<td class="doc">The native credentials type is a struct ucred.</td>
</tr>
<tr>
<td class="name">CREDENTIALS_TYPE_FREEBSD_CMSGCRED</td>
<td class="value">2</td>
<td class="doc">The native credentials type is a struct cmsgcred.</td>
</tr>
<tr>
<td class="name">CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED</td>
<td class="value">3</td>
<td class="doc">The native credentials type is a struct sockpeercred. Added in 2.30.</td>
</tr>
<tr>
<td class="name">CREDENTIALS_TYPE_SOLARIS_UCRED</td>
<td class="value">4</td>
<td class="doc">The native credentials type is a ucred_t. Added in 2.40.</td>
</tr>
<tr>
<td class="name">CREDENTIALS_TYPE_NETBSD_UNPCBID</td>
<td class="value">5</td>
<td class="doc">The native credentials type is a struct unpcbid.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusError</p>
<p class="api-doc">Error codes for the %G_DBUS_ERROR error domain.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusError</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_ERROR_FAILED</td>
<td class="value">0</td>
<td class="doc">A generic error; "something went wrong" - see the error message for
more.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_NO_MEMORY</td>
<td class="value">1</td>
<td class="doc">There was not enough memory to complete an operation.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SERVICE_UNKNOWN</td>
<td class="value">2</td>
<td class="doc">The bus doesn't know how to launch a service to supply the bus name
you wanted.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_NAME_HAS_NO_OWNER</td>
<td class="value">3</td>
<td class="doc">The bus name you referenced doesn't exist (i.e. no application owns
it).</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_NO_REPLY</td>
<td class="value">4</td>
<td class="doc">No reply to a message expecting one, usually means a timeout occurred.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_IO_ERROR</td>
<td class="value">5</td>
<td class="doc">Something went wrong reading or writing to a socket, for example.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_BAD_ADDRESS</td>
<td class="value">6</td>
<td class="doc">A D-Bus bus address was malformed.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_NOT_SUPPORTED</td>
<td class="value">7</td>
<td class="doc">Requested operation isn't supported (like ENOSYS on UNIX).</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_LIMITS_EXCEEDED</td>
<td class="value">8</td>
<td class="doc">Some limited resource is exhausted.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_ACCESS_DENIED</td>
<td class="value">9</td>
<td class="doc">Security restrictions don't allow doing what you're trying to do.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_AUTH_FAILED</td>
<td class="value">10</td>
<td class="doc">Authentication didn't work.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_NO_SERVER</td>
<td class="value">11</td>
<td class="doc">Unable to connect to server (probably caused by ECONNREFUSED on a
socket).</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_TIMEOUT</td>
<td class="value">12</td>
<td class="doc">Certain timeout errors, possibly ETIMEDOUT on a socket.  Note that
%G_DBUS_ERROR_NO_REPLY is used for message reply timeouts. Warning:
this is confusingly-named given that %G_DBUS_ERROR_TIMED_OUT also
exists. We can't fix it for compatibility reasons so just be
careful.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_NO_NETWORK</td>
<td class="value">13</td>
<td class="doc">No network access (probably ENETUNREACH on a socket).</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_ADDRESS_IN_USE</td>
<td class="value">14</td>
<td class="doc">Can't bind a socket since its address is in use (i.e. EADDRINUSE).</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_DISCONNECTED</td>
<td class="value">15</td>
<td class="doc">The connection is disconnected and you're trying to use it.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_INVALID_ARGS</td>
<td class="value">16</td>
<td class="doc">Invalid arguments passed to a method call.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_FILE_NOT_FOUND</td>
<td class="value">17</td>
<td class="doc">Missing file.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_FILE_EXISTS</td>
<td class="value">18</td>
<td class="doc">Existing file and the operation you're using does not silently overwrite.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_UNKNOWN_METHOD</td>
<td class="value">19</td>
<td class="doc">Method name you invoked isn't known by the object you invoked it on.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_TIMED_OUT</td>
<td class="value">20</td>
<td class="doc">Certain timeout errors, e.g. while starting a service. Warning: this is
confusingly-named given that %G_DBUS_ERROR_TIMEOUT also exists. We
can't fix it for compatibility reasons so just be careful.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_MATCH_RULE_NOT_FOUND</td>
<td class="value">21</td>
<td class="doc">Tried to remove or modify a match rule that didn't exist.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_MATCH_RULE_INVALID</td>
<td class="value">22</td>
<td class="doc">The match rule isn't syntactically valid.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_EXEC_FAILED</td>
<td class="value">23</td>
<td class="doc">While starting a new process, the exec() call failed.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_FORK_FAILED</td>
<td class="value">24</td>
<td class="doc">While starting a new process, the fork() call failed.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_CHILD_EXITED</td>
<td class="value">25</td>
<td class="doc">While starting a new process, the child exited with a status code.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_CHILD_SIGNALED</td>
<td class="value">26</td>
<td class="doc">While starting a new process, the child exited on a signal.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_FAILED</td>
<td class="value">27</td>
<td class="doc">While starting a new process, something went wrong.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_SETUP_FAILED</td>
<td class="value">28</td>
<td class="doc">We failed to setup the environment correctly.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_CONFIG_INVALID</td>
<td class="value">29</td>
<td class="doc">We failed to setup the config parser correctly.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_SERVICE_INVALID</td>
<td class="value">30</td>
<td class="doc">Bus name was not valid.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND</td>
<td class="value">31</td>
<td class="doc">Service file not found in system-services directory.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_PERMISSIONS_INVALID</td>
<td class="value">32</td>
<td class="doc">Permissions are incorrect on the setuid helper.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_FILE_INVALID</td>
<td class="value">33</td>
<td class="doc">Service file invalid (Name, User or Exec missing).</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SPAWN_NO_MEMORY</td>
<td class="value">34</td>
<td class="doc">Tried to get a UNIX process ID and it wasn't available.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN</td>
<td class="value">35</td>
<td class="doc">Tried to get a UNIX process ID and it wasn't available.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_INVALID_SIGNATURE</td>
<td class="value">36</td>
<td class="doc">A type signature is not valid.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_INVALID_FILE_CONTENT</td>
<td class="value">37</td>
<td class="doc">A file contains invalid syntax or is otherwise broken.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN</td>
<td class="value">38</td>
<td class="doc">Asked for SELinux security context and it wasn't available.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN</td>
<td class="value">39</td>
<td class="doc">Asked for ADT audit data and it wasn't available.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_OBJECT_PATH_IN_USE</td>
<td class="value">40</td>
<td class="doc">There's already an object with the requested object path.</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_UNKNOWN_OBJECT</td>
<td class="value">41</td>
<td class="doc">Object you invoked a method on isn't known. Since 2.42</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_UNKNOWN_INTERFACE</td>
<td class="value">42</td>
<td class="doc">Interface you invoked a method on isn't known by the object. Since 2.42</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_UNKNOWN_PROPERTY</td>
<td class="value">43</td>
<td class="doc">Property you tried to access isn't known by the object. Since 2.42</td>
</tr>
<tr>
<td class="name">DBUS_ERROR_PROPERTY_READ_ONLY</td>
<td class="value">44</td>
<td class="doc">Property you tried to set is read-only. Since 2.42</td>
</tr>
</table>
</div>
<p class="api-heading">DBusMessageByteOrder</p>
<p class="api-doc">Enumeration used to describe the byte order of a D-Bus message.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusMessageByteOrder</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN</td>
<td class="value">66</td>
<td class="doc">The byte order is big endian.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN</td>
<td class="value">108</td>
<td class="doc">The byte order is little endian.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusMessageHeaderField</p>
<p class="api-doc">Header fields used in #GDBusMessage.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusMessageHeaderField</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_MESSAGE_HEADER_FIELD_INVALID</td>
<td class="value">0</td>
<td class="doc">Not a valid header field.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_HEADER_FIELD_PATH</td>
<td class="value">1</td>
<td class="doc">The object path.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_HEADER_FIELD_INTERFACE</td>
<td class="value">2</td>
<td class="doc">The interface name.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_HEADER_FIELD_MEMBER</td>
<td class="value">3</td>
<td class="doc">The method or signal name.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME</td>
<td class="value">4</td>
<td class="doc">The name of the error that occurred.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL</td>
<td class="value">5</td>
<td class="doc">The serial number the message is a reply to.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_HEADER_FIELD_DESTINATION</td>
<td class="value">6</td>
<td class="doc">The name the message is intended for.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_HEADER_FIELD_SENDER</td>
<td class="value">7</td>
<td class="doc">Unique name of the sender of the message (filled in by the bus).</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_HEADER_FIELD_SIGNATURE</td>
<td class="value">8</td>
<td class="doc">The signature of the message body.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS</td>
<td class="value">9</td>
<td class="doc">The number of UNIX file descriptors that accompany the message.</td>
</tr>
</table>
</div>
<p class="api-heading">DBusMessageType</p>
<p class="api-doc">Message types used in #GDBusMessage.</p>
<div class="api-notes">
  <p class="api-ctype">GDBusMessageType</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">DBUS_MESSAGE_TYPE_INVALID</td>
<td class="value">0</td>
<td class="doc">Message is of invalid type.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_TYPE_METHOD_CALL</td>
<td class="value">1</td>
<td class="doc">Method call.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_TYPE_METHOD_RETURN</td>
<td class="value">2</td>
<td class="doc">Method reply.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_TYPE_ERROR</td>
<td class="value">3</td>
<td class="doc">Error reply.</td>
</tr>
<tr>
<td class="name">DBUS_MESSAGE_TYPE_SIGNAL</td>
<td class="value">4</td>
<td class="doc">Signal emission.</td>
</tr>
</table>
</div>
<p class="api-heading">DataStreamByteOrder</p>
<p class="api-doc">#GDataStreamByteOrder is used to ensure proper endianness of streaming data sources
across various machine architectures.</p>
<div class="api-notes">
  <p class="api-ctype">GDataStreamByteOrder</p>
<table>
<tr>
<td class="name">DATA_STREAM_BYTE_ORDER_BIG_ENDIAN</td>
<td class="value">0</td>
<td class="doc">Selects Big Endian byte order.</td>
</tr>
<tr>
<td class="name">DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN</td>
<td class="value">1</td>
<td class="doc">Selects Little Endian byte order.</td>
</tr>
<tr>
<td class="name">DATA_STREAM_BYTE_ORDER_HOST_ENDIAN</td>
<td class="value">2</td>
<td class="doc">Selects endianness based on host machine's architecture.</td>
</tr>
</table>
</div>
<p class="api-heading">DataStreamNewlineType</p>
<p class="api-doc">#GDataStreamNewlineType is used when checking for or setting the line endings for a given file.</p>
<div class="api-notes">
  <p class="api-ctype">GDataStreamNewlineType</p>
<table>
<tr>
<td class="name">DATA_STREAM_NEWLINE_TYPE_LF</td>
<td class="value">0</td>
<td class="doc">Selects "LF" line endings, common on most modern UNIX platforms.</td>
</tr>
<tr>
<td class="name">DATA_STREAM_NEWLINE_TYPE_CR</td>
<td class="value">1</td>
<td class="doc">Selects "CR" line endings.</td>
</tr>
<tr>
<td class="name">DATA_STREAM_NEWLINE_TYPE_CR_LF</td>
<td class="value">2</td>
<td class="doc">Selects "CR, LF" line ending, common on Microsoft Windows.</td>
</tr>
<tr>
<td class="name">DATA_STREAM_NEWLINE_TYPE_ANY</td>
<td class="value">3</td>
<td class="doc">Automatically try to handle any line ending type.</td>
</tr>
</table>
</div>
<p class="api-heading">DriveStartStopType</p>
<p class="api-doc">Enumeration describing how a drive can be started/stopped.</p>
<div class="api-notes">
  <p class="api-ctype">GDriveStartStopType</p>
  <p class="api-since">since 2.22</p>
<table>
<tr>
<td class="name">DRIVE_START_STOP_TYPE_UNKNOWN</td>
<td class="value">0</td>
<td class="doc">Unknown or drive doesn't support
   start/stop.</td>
</tr>
<tr>
<td class="name">DRIVE_START_STOP_TYPE_SHUTDOWN</td>
<td class="value">1</td>
<td class="doc">The stop method will physically
   shut down the drive and e.g. power down the port the drive is
   attached to.</td>
</tr>
<tr>
<td class="name">DRIVE_START_STOP_TYPE_NETWORK</td>
<td class="value">2</td>
<td class="doc">The start/stop methods are used
   for connecting/disconnect to the drive over the network.</td>
</tr>
<tr>
<td class="name">DRIVE_START_STOP_TYPE_MULTIDISK</td>
<td class="value">3</td>
<td class="doc">The start/stop methods will
   assemble/disassemble a virtual drive from several physical
   drives.</td>
</tr>
<tr>
<td class="name">DRIVE_START_STOP_TYPE_PASSWORD</td>
<td class="value">4</td>
<td class="doc">The start/stop methods will
   unlock/lock the disk (for example using the ATA <quote>SECURITY
   UNLOCK DEVICE</quote> command)</td>
</tr>
</table>
</div>
<p class="api-heading">EmblemOrigin</p>
<p class="api-doc">GEmblemOrigin is used to add information about the origin of the emblem
to #GEmblem.</p>
<div class="api-notes">
  <p class="api-ctype">GEmblemOrigin</p>
  <p class="api-since">since 2.18</p>
<table>
<tr>
<td class="name">EMBLEM_ORIGIN_UNKNOWN</td>
<td class="value">0</td>
<td class="doc">Emblem of unknown origin</td>
</tr>
<tr>
<td class="name">EMBLEM_ORIGIN_DEVICE</td>
<td class="value">1</td>
<td class="doc">Emblem adds device-specific information</td>
</tr>
<tr>
<td class="name">EMBLEM_ORIGIN_LIVEMETADATA</td>
<td class="value">2</td>
<td class="doc">Emblem depicts live metadata, such as "readonly"</td>
</tr>
<tr>
<td class="name">EMBLEM_ORIGIN_TAG</td>
<td class="value">3</td>
<td class="doc">Emblem comes from a user-defined tag, e.g. set by nautilus (in the future)</td>
</tr>
</table>
</div>
<p class="api-heading">FileAttributeStatus</p>
<p class="api-doc">Used by g_file_set_attributes_from_info() when setting file attributes.</p>
<div class="api-notes">
  <p class="api-ctype">GFileAttributeStatus</p>
<table>
<tr>
<td class="name">FILE_ATTRIBUTE_STATUS_UNSET</td>
<td class="value">0</td>
<td class="doc">Attribute value is unset (empty).</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_STATUS_SET</td>
<td class="value">1</td>
<td class="doc">Attribute value is set.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_STATUS_ERROR_SETTING</td>
<td class="value">2</td>
<td class="doc">Indicates an error in setting the value.</td>
</tr>
</table>
</div>
<p class="api-heading">FileAttributeType</p>
<p class="api-doc">The data types for file attributes.</p>
<div class="api-notes">
  <p class="api-ctype">GFileAttributeType</p>
<table>
<tr>
<td class="name">FILE_ATTRIBUTE_TYPE_INVALID</td>
<td class="value">0</td>
<td class="doc">indicates an invalid or uninitalized type.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_TYPE_STRING</td>
<td class="value">1</td>
<td class="doc">a null terminated UTF8 string.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_TYPE_BYTE_STRING</td>
<td class="value">2</td>
<td class="doc">a zero terminated string of non-zero bytes.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_TYPE_BOOLEAN</td>
<td class="value">3</td>
<td class="doc">a boolean value.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_TYPE_UINT32</td>
<td class="value">4</td>
<td class="doc">an unsigned 4-byte/32-bit integer.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_TYPE_INT32</td>
<td class="value">5</td>
<td class="doc">a signed 4-byte/32-bit integer.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_TYPE_UINT64</td>
<td class="value">6</td>
<td class="doc">an unsigned 8-byte/64-bit integer.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_TYPE_INT64</td>
<td class="value">7</td>
<td class="doc">a signed 8-byte/64-bit integer.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_TYPE_OBJECT</td>
<td class="value">8</td>
<td class="doc">a #GObject.</td>
</tr>
<tr>
<td class="name">FILE_ATTRIBUTE_TYPE_STRINGV</td>
<td class="value">9</td>
<td class="doc">a %NULL terminated char **. Since 2.22</td>
</tr>
</table>
</div>
<p class="api-heading">FileMonitorEvent</p>
<p class="api-doc">Specifies what type of event a monitor event is.</p>
<div class="api-notes">
  <p class="api-ctype">GFileMonitorEvent</p>
<table>
<tr>
<td class="name">FILE_MONITOR_EVENT_CHANGED</td>
<td class="value">0</td>
<td class="doc">a file changed.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_EVENT_CHANGES_DONE_HINT</td>
<td class="value">1</td>
<td class="doc">a hint that this was probably the last change in a set of changes.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_EVENT_DELETED</td>
<td class="value">2</td>
<td class="doc">a file was deleted.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_EVENT_CREATED</td>
<td class="value">3</td>
<td class="doc">a file was created.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED</td>
<td class="value">4</td>
<td class="doc">a file attribute was changed.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_EVENT_PRE_UNMOUNT</td>
<td class="value">5</td>
<td class="doc">the file location will soon be unmounted.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_EVENT_UNMOUNTED</td>
<td class="value">6</td>
<td class="doc">the file location was unmounted.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_EVENT_MOVED</td>
<td class="value">7</td>
<td class="doc">the file was moved -- only sent if the
  (deprecated) %G_FILE_MONITOR_SEND_MOVED flag is set</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_EVENT_RENAMED</td>
<td class="value">8</td>
<td class="doc">the file was renamed within the
  current directory -- only sent if the %G_FILE_MONITOR_WATCH_MOVES
  flag is set.  Since: 2.46.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_EVENT_MOVED_IN</td>
<td class="value">9</td>
<td class="doc">the file was moved into the
  monitored directory from another location -- only sent if the
  %G_FILE_MONITOR_WATCH_MOVES flag is set.  Since: 2.46.</td>
</tr>
<tr>
<td class="name">FILE_MONITOR_EVENT_MOVED_OUT</td>
<td class="value">10</td>
<td class="doc">the file was moved out of the
  monitored directory to another location -- only sent if the
  %G_FILE_MONITOR_WATCH_MOVES flag is set.  Since: 2.46</td>
</tr>
</table>
</div>
<p class="api-heading">FileType</p>
<p class="api-doc">Indicates the file's on-disk type.</p>
<div class="api-notes">
  <p class="api-ctype">GFileType</p>
<table>
<tr>
<td class="name">FILE_TYPE_UNKNOWN</td>
<td class="value">0</td>
<td class="doc">File's type is unknown.</td>
</tr>
<tr>
<td class="name">FILE_TYPE_REGULAR</td>
<td class="value">1</td>
<td class="doc">File handle represents a regular file.</td>
</tr>
<tr>
<td class="name">FILE_TYPE_DIRECTORY</td>
<td class="value">2</td>
<td class="doc">File handle represents a directory.</td>
</tr>
<tr>
<td class="name">FILE_TYPE_SYMBOLIC_LINK</td>
<td class="value">3</td>
<td class="doc">File handle represents a symbolic link
   (Unix systems).</td>
</tr>
<tr>
<td class="name">FILE_TYPE_SPECIAL</td>
<td class="value">4</td>
<td class="doc">File is a "special" file, such as a socket, fifo,
   block device, or character device.</td>
</tr>
<tr>
<td class="name">FILE_TYPE_SHORTCUT</td>
<td class="value">5</td>
<td class="doc">File is a shortcut (Windows systems).</td>
</tr>
<tr>
<td class="name">FILE_TYPE_MOUNTABLE</td>
<td class="value">6</td>
<td class="doc">File is a mountable location.</td>
</tr>
</table>
</div>
<p class="api-heading">FilesystemPreviewType</p>
<p class="api-doc">Indicates a hint from the file system whether files should be
previewed in a file manager. Returned as the value of the key
#G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW.</p>
<div class="api-notes">
  <p class="api-ctype">GFilesystemPreviewType</p>
<table>
<tr>
<td class="name">FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS</td>
<td class="value">0</td>
<td class="doc">Only preview files if user has explicitly requested it.</td>
</tr>
<tr>
<td class="name">FILESYSTEM_PREVIEW_TYPE_IF_LOCAL</td>
<td class="value">1</td>
<td class="doc">Preview files if user has requested preview of "local" files.</td>
</tr>
<tr>
<td class="name">FILESYSTEM_PREVIEW_TYPE_NEVER</td>
<td class="value">2</td>
<td class="doc">Never preview files.</td>
</tr>
</table>
</div>
<p class="api-heading">IOErrorEnum</p>
<p class="api-doc">Error codes returned by GIO functions.

Note that this domain may be extended in future GLib releases. In
general, new error codes either only apply to new APIs, or else
replace %G_IO_ERROR_FAILED in cases that were not explicitly
distinguished before. You should therefore avoid writing code like
|[<!-- language="C" -->
if (g_error_matches (error, G_IO_ERROR, G_IO_ERROR_FAILED))
  {
    // Assume that this is EPRINTERONFIRE
    ...
  }
]|
but should instead treat all unrecognized error codes the same as
#G_IO_ERROR_FAILED.</p>
<div class="api-notes">
  <p class="api-ctype">GIOErrorEnum</p>
<table>
<tr>
<td class="name">IO_ERROR_FAILED</td>
<td class="value">0</td>
<td class="doc">Generic error condition for when an operation fails
    and no more specific #GIOErrorEnum value is defined.</td>
</tr>
<tr>
<td class="name">IO_ERROR_NOT_FOUND</td>
<td class="value">1</td>
<td class="doc">File not found.</td>
</tr>
<tr>
<td class="name">IO_ERROR_EXISTS</td>
<td class="value">2</td>
<td class="doc">File already exists.</td>
</tr>
<tr>
<td class="name">IO_ERROR_IS_DIRECTORY</td>
<td class="value">3</td>
<td class="doc">File is a directory.</td>
</tr>
<tr>
<td class="name">IO_ERROR_NOT_DIRECTORY</td>
<td class="value">4</td>
<td class="doc">File is not a directory.</td>
</tr>
<tr>
<td class="name">IO_ERROR_NOT_EMPTY</td>
<td class="value">5</td>
<td class="doc">File is a directory that isn't empty.</td>
</tr>
<tr>
<td class="name">IO_ERROR_NOT_REGULAR_FILE</td>
<td class="value">6</td>
<td class="doc">File is not a regular file.</td>
</tr>
<tr>
<td class="name">IO_ERROR_NOT_SYMBOLIC_LINK</td>
<td class="value">7</td>
<td class="doc">File is not a symbolic link.</td>
</tr>
<tr>
<td class="name">IO_ERROR_NOT_MOUNTABLE_FILE</td>
<td class="value">8</td>
<td class="doc">File cannot be mounted.</td>
</tr>
<tr>
<td class="name">IO_ERROR_FILENAME_TOO_LONG</td>
<td class="value">9</td>
<td class="doc">Filename is too many characters.</td>
</tr>
<tr>
<td class="name">IO_ERROR_INVALID_FILENAME</td>
<td class="value">10</td>
<td class="doc">Filename is invalid or contains invalid characters.</td>
</tr>
<tr>
<td class="name">IO_ERROR_TOO_MANY_LINKS</td>
<td class="value">11</td>
<td class="doc">File contains too many symbolic links.</td>
</tr>
<tr>
<td class="name">IO_ERROR_NO_SPACE</td>
<td class="value">12</td>
<td class="doc">No space left on drive.</td>
</tr>
<tr>
<td class="name">IO_ERROR_INVALID_ARGUMENT</td>
<td class="value">13</td>
<td class="doc">Invalid argument.</td>
</tr>
<tr>
<td class="name">IO_ERROR_PERMISSION_DENIED</td>
<td class="value">14</td>
<td class="doc">Permission denied.</td>
</tr>
<tr>
<td class="name">IO_ERROR_NOT_SUPPORTED</td>
<td class="value">15</td>
<td class="doc">Operation (or one of its parameters) not supported</td>
</tr>
<tr>
<td class="name">IO_ERROR_NOT_MOUNTED</td>
<td class="value">16</td>
<td class="doc">File isn't mounted.</td>
</tr>
<tr>
<td class="name">IO_ERROR_ALREADY_MOUNTED</td>
<td class="value">17</td>
<td class="doc">File is already mounted.</td>
</tr>
<tr>
<td class="name">IO_ERROR_CLOSED</td>
<td class="value">18</td>
<td class="doc">File was closed.</td>
</tr>
<tr>
<td class="name">IO_ERROR_CANCELLED</td>
<td class="value">19</td>
<td class="doc">Operation was cancelled. See #GCancellable.</td>
</tr>
<tr>
<td class="name">IO_ERROR_PENDING</td>
<td class="value">20</td>
<td class="doc">Operations are still pending.</td>
</tr>
<tr>
<td class="name">IO_ERROR_READ_ONLY</td>
<td class="value">21</td>
<td class="doc">File is read only.</td>
</tr>
<tr>
<td class="name">IO_ERROR_CANT_CREATE_BACKUP</td>
<td class="value">22</td>
<td class="doc">Backup couldn't be created.</td>
</tr>
<tr>
<td class="name">IO_ERROR_WRONG_ETAG</td>
<td class="value">23</td>
<td class="doc">File's Entity Tag was incorrect.</td>
</tr>
<tr>
<td class="name">IO_ERROR_TIMED_OUT</td>
<td class="value">24</td>
<td class="doc">Operation timed out.</td>
</tr>
<tr>
<td class="name">IO_ERROR_WOULD_RECURSE</td>
<td class="value">25</td>
<td class="doc">Operation would be recursive.</td>
</tr>
<tr>
<td class="name">IO_ERROR_BUSY</td>
<td class="value">26</td>
<td class="doc">File is busy.</td>
</tr>
<tr>
<td class="name">IO_ERROR_WOULD_BLOCK</td>
<td class="value">27</td>
<td class="doc">Operation would block.</td>
</tr>
<tr>
<td class="name">IO_ERROR_HOST_NOT_FOUND</td>
<td class="value">28</td>
<td class="doc">Host couldn't be found (remote operations).</td>
</tr>
<tr>
<td class="name">IO_ERROR_WOULD_MERGE</td>
<td class="value">29</td>
<td class="doc">Operation would merge files.</td>
</tr>
<tr>
<td class="name">IO_ERROR_FAILED_HANDLED</td>
<td class="value">30</td>
<td class="doc">Operation failed and a helper program has
    already interacted with the user. Do not display any error dialog.</td>
</tr>
<tr>
<td class="name">IO_ERROR_TOO_MANY_OPEN_FILES</td>
<td class="value">31</td>
<td class="doc">The current process has too many files
    open and can't open any more. Duplicate descriptors do count toward
    this limit. Since 2.20</td>
</tr>
<tr>
<td class="name">IO_ERROR_NOT_INITIALIZED</td>
<td class="value">32</td>
<td class="doc">The object has not been initialized. Since 2.22</td>
</tr>
<tr>
<td class="name">IO_ERROR_ADDRESS_IN_USE</td>
<td class="value">33</td>
<td class="doc">The requested address is already in use. Since 2.22</td>
</tr>
<tr>
<td class="name">IO_ERROR_PARTIAL_INPUT</td>
<td class="value">34</td>
<td class="doc">Need more input to finish operation. Since 2.24</td>
</tr>
<tr>
<td class="name">IO_ERROR_INVALID_DATA</td>
<td class="value">35</td>
<td class="doc">The input data was invalid. Since 2.24</td>
</tr>
<tr>
<td class="name">IO_ERROR_DBUS_ERROR</td>
<td class="value">36</td>
<td class="doc">A remote object generated an error that
    doesn't correspond to a locally registered #GError error
    domain. Use g_dbus_error_get_remote_error() to extract the D-Bus
    error name and g_dbus_error_strip_remote_error() to fix up the
    message so it matches what was received on the wire. Since 2.26.</td>
</tr>
<tr>
<td class="name">IO_ERROR_HOST_UNREACHABLE</td>
<td class="value">37</td>
<td class="doc">Host unreachable. Since 2.26</td>
</tr>
<tr>
<td class="name">IO_ERROR_NETWORK_UNREACHABLE</td>
<td class="value">38</td>
<td class="doc">Network unreachable. Since 2.26</td>
</tr>
<tr>
<td class="name">IO_ERROR_CONNECTION_REFUSED</td>
<td class="value">39</td>
<td class="doc">Connection refused. Since 2.26</td>
</tr>
<tr>
<td class="name">IO_ERROR_PROXY_FAILED</td>
<td class="value">40</td>
<td class="doc">Connection to proxy server failed. Since 2.26</td>
</tr>
<tr>
<td class="name">IO_ERROR_PROXY_AUTH_FAILED</td>
<td class="value">41</td>
<td class="doc">Proxy authentication failed. Since 2.26</td>
</tr>
<tr>
<td class="name">IO_ERROR_PROXY_NEED_AUTH</td>
<td class="value">42</td>
<td class="doc">Proxy server needs authentication. Since 2.26</td>
</tr>
<tr>
<td class="name">IO_ERROR_PROXY_NOT_ALLOWED</td>
<td class="value">43</td>
<td class="doc">Proxy connection is not allowed by ruleset.
    Since 2.26</td>
</tr>
<tr>
<td class="name">IO_ERROR_BROKEN_PIPE</td>
<td class="value">44</td>
<td class="doc">Broken pipe. Since 2.36</td>
</tr>
<tr>
<td class="name">IO_ERROR_CONNECTION_CLOSED</td>
<td class="value">44</td>
<td class="doc">Connection closed by peer. Note that this
    is the same code as %G_IO_ERROR_BROKEN_PIPE; before 2.44 some
    "connection closed" errors returned %G_IO_ERROR_BROKEN_PIPE, but others
    returned %G_IO_ERROR_FAILED. Now they should all return the same
    value, which has this more logical name. Since 2.44.</td>
</tr>
<tr>
<td class="name">IO_ERROR_NOT_CONNECTED</td>
<td class="value">45</td>
<td class="doc">Transport endpoint is not connected. Since 2.44</td>
</tr>
<tr>
<td class="name">IO_ERROR_MESSAGE_TOO_LARGE</td>
<td class="value">46</td>
<td class="doc">Message too large. Since 2.48.</td>
</tr>
</table>
</div>
<p class="api-heading">IOModuleScopeFlags</p>
<p class="api-doc">Flags for use with g_io_module_scope_new().</p>
<div class="api-notes">
  <p class="api-ctype">GIOModuleScopeFlags</p>
  <p class="api-since">since 2.30</p>
<table>
<tr>
<td class="name">IO_MODULE_SCOPE_NONE</td>
<td class="value">0</td>
<td class="doc">No module scan flags</td>
</tr>
<tr>
<td class="name">IO_MODULE_SCOPE_BLOCK_DUPLICATES</td>
<td class="value">1</td>
<td class="doc">When using this scope to load or
    scan modules, automatically block a modules which has the same base
    basename as previously loaded module.</td>
</tr>
</table>
</div>
<p class="api-heading">MountOperationResult</p>
<p class="api-doc">#GMountOperationResult is returned as a result when a request for
information is send by the mounting operation.</p>
<div class="api-notes">
  <p class="api-ctype">GMountOperationResult</p>
<table>
<tr>
<td class="name">MOUNT_OPERATION_HANDLED</td>
<td class="value">0</td>
<td class="doc">The request was fulfilled and the
    user specified data is now available</td>
</tr>
<tr>
<td class="name">MOUNT_OPERATION_ABORTED</td>
<td class="value">1</td>
<td class="doc">The user requested the mount operation
    to be aborted</td>
</tr>
<tr>
<td class="name">MOUNT_OPERATION_UNHANDLED</td>
<td class="value">2</td>
<td class="doc">The request was unhandled (i.e. not
    implemented)</td>
</tr>
</table>
</div>
<p class="api-heading">NetworkConnectivity</p>
<p class="api-doc">The host's network connectivity state, as reported by #GNetworkMonitor.</p>
<div class="api-notes">
  <p class="api-ctype">GNetworkConnectivity</p>
  <p class="api-since">since 2.44</p>
<table>
<tr>
<td class="name">NETWORK_CONNECTIVITY_LOCAL</td>
<td class="value">1</td>
<td class="doc">The host is not configured with a
  route to the Internet; it may or may not be connected to a local
  network.</td>
</tr>
<tr>
<td class="name">NETWORK_CONNECTIVITY_LIMITED</td>
<td class="value">2</td>
<td class="doc">The host is connected to a network, but
  does not appear to be able to reach the full Internet, perhaps
  due to upstream network problems.</td>
</tr>
<tr>
<td class="name">NETWORK_CONNECTIVITY_PORTAL</td>
<td class="value">3</td>
<td class="doc">The host is behind a captive portal and
  cannot reach the full Internet.</td>
</tr>
<tr>
<td class="name">NETWORK_CONNECTIVITY_FULL</td>
<td class="value">4</td>
<td class="doc">The host is connected to a network, and
  appears to be able to reach the full Internet.</td>
</tr>
</table>
</div>
<p class="api-heading">NotificationPriority</p>
<p class="api-doc">Priority levels for #GNotifications.</p>
<div class="api-notes">
  <p class="api-ctype">GNotificationPriority</p>
  <p class="api-since">since 2.42</p>
<table>
<tr>
<td class="name">NOTIFICATION_PRIORITY_NORMAL</td>
<td class="value">0</td>
<td class="doc">the default priority, to be used for the
  majority of notifications (for example email messages, software updates,
  completed download/sync operations)</td>
</tr>
<tr>
<td class="name">NOTIFICATION_PRIORITY_LOW</td>
<td class="value">1</td>
<td class="doc">for notifications that do not require
  immediate attention - typically used for contextual background
  information, such as contact birthdays or local weather</td>
</tr>
<tr>
<td class="name">NOTIFICATION_PRIORITY_HIGH</td>
<td class="value">2</td>
<td class="doc">for events that require more attention,
  usually because responses are time-sensitive (for example chat and SMS
  messages or alarms)</td>
</tr>
<tr>
<td class="name">NOTIFICATION_PRIORITY_URGENT</td>
<td class="value">3</td>
<td class="doc">for urgent notifications, or notifications
  that require a response in a short space of time (for example phone calls
  or emergency warnings)</td>
</tr>
</table>
</div>
<p class="api-heading">PasswordSave</p>
<p class="api-doc">#GPasswordSave is used to indicate the lifespan of a saved password.

#Gvfs stores passwords in the Gnome keyring when this flag allows it
to, and later retrieves it again from there.</p>
<div class="api-notes">
  <p class="api-ctype">GPasswordSave</p>
<table>
<tr>
<td class="name">PASSWORD_SAVE_NEVER</td>
<td class="value">0</td>
<td class="doc">never save a password.</td>
</tr>
<tr>
<td class="name">PASSWORD_SAVE_FOR_SESSION</td>
<td class="value">1</td>
<td class="doc">save a password for the session.</td>
</tr>
<tr>
<td class="name">PASSWORD_SAVE_PERMANENTLY</td>
<td class="value">2</td>
<td class="doc">save a password permanently.</td>
</tr>
</table>
</div>
<p class="api-heading">ResolverError</p>
<p class="api-doc">An error code used with %G_RESOLVER_ERROR in a #GError returned
from a #GResolver routine.</p>
<div class="api-notes">
  <p class="api-ctype">GResolverError</p>
  <p class="api-since">since 2.22</p>
<table>
<tr>
<td class="name">RESOLVER_ERROR_NOT_FOUND</td>
<td class="value">0</td>
<td class="doc">the requested name/address/service was not
    found</td>
</tr>
<tr>
<td class="name">RESOLVER_ERROR_TEMPORARY_FAILURE</td>
<td class="value">1</td>
<td class="doc">the requested information could not
    be looked up due to a network error or similar problem</td>
</tr>
<tr>
<td class="name">RESOLVER_ERROR_INTERNAL</td>
<td class="value">2</td>
<td class="doc">unknown error</td>
</tr>
</table>
</div>
<p class="api-heading">ResolverRecordType</p>
<p class="api-doc">The type of record that g_resolver_lookup_records() or
g_resolver_lookup_records_async() should retrieve. The records are returned
as lists of #GVariant tuples. Each record type has different values in
the variant tuples returned.

%G_RESOLVER_RECORD_SRV records are returned as variants with the signature
'(qqqs)', containing a guint16 with the priority, a guint16 with the
weight, a guint16 with the port, and a string of the hostname.

%G_RESOLVER_RECORD_MX records are returned as variants with the signature
'(qs)', representing a guint16 with the preference, and a string containing
the mail exchanger hostname.

%G_RESOLVER_RECORD_TXT records are returned as variants with the signature
'(as)', representing an array of the strings in the text record.

%G_RESOLVER_RECORD_SOA records are returned as variants with the signature
'(ssuuuuu)', representing a string containing the primary name server, a
string containing the administrator, the serial as a guint32, the refresh
interval as guint32, the retry interval as a guint32, the expire timeout
as a guint32, and the ttl as a guint32.

%G_RESOLVER_RECORD_NS records are returned as variants with the signature
'(s)', representing a string of the hostname of the name server.</p>
<div class="api-notes">
  <p class="api-ctype">GResolverRecordType</p>
  <p class="api-since">since 2.34</p>
<table>
<tr>
<td class="name">RESOLVER_RECORD_SRV</td>
<td class="value">1</td>
<td class="doc">lookup DNS SRV records for a domain</td>
</tr>
<tr>
<td class="name">RESOLVER_RECORD_MX</td>
<td class="value">2</td>
<td class="doc">lookup DNS MX records for a domain</td>
</tr>
<tr>
<td class="name">RESOLVER_RECORD_TXT</td>
<td class="value">3</td>
<td class="doc">lookup DNS TXT records for a name</td>
</tr>
<tr>
<td class="name">RESOLVER_RECORD_SOA</td>
<td class="value">4</td>
<td class="doc">lookup DNS SOA records for a zone</td>
</tr>
<tr>
<td class="name">RESOLVER_RECORD_NS</td>
<td class="value">5</td>
<td class="doc">lookup DNS NS records for a domain</td>
</tr>
</table>
</div>
<p class="api-heading">ResourceError</p>
<p class="api-doc">An error code used with %G_RESOURCE_ERROR in a #GError returned
from a #GResource routine.</p>
<div class="api-notes">
  <p class="api-ctype">GResourceError</p>
  <p class="api-since">since 2.32</p>
<table>
<tr>
<td class="name">RESOURCE_ERROR_NOT_FOUND</td>
<td class="value">0</td>
<td class="doc">no file was found at the requested path</td>
</tr>
<tr>
<td class="name">RESOURCE_ERROR_INTERNAL</td>
<td class="value">1</td>
<td class="doc">unknown error</td>
</tr>
</table>
</div>
<p class="api-heading">SocketClientEvent</p>
<p class="api-doc">Describes an event occurring on a #GSocketClient. See the
#GSocketClient::event signal for more details.

Additional values may be added to this type in the future.</p>
<div class="api-notes">
  <p class="api-ctype">GSocketClientEvent</p>
  <p class="api-since">since 2.32</p>
<table>
<tr>
<td class="name">SOCKET_CLIENT_RESOLVING</td>
<td class="value">0</td>
<td class="doc">The client is doing a DNS lookup.</td>
</tr>
<tr>
<td class="name">SOCKET_CLIENT_RESOLVED</td>
<td class="value">1</td>
<td class="doc">The client has completed a DNS lookup.</td>
</tr>
<tr>
<td class="name">SOCKET_CLIENT_CONNECTING</td>
<td class="value">2</td>
<td class="doc">The client is connecting to a remote
  host (either a proxy or the destination server).</td>
</tr>
<tr>
<td class="name">SOCKET_CLIENT_CONNECTED</td>
<td class="value">3</td>
<td class="doc">The client has connected to a remote
  host.</td>
</tr>
<tr>
<td class="name">SOCKET_CLIENT_PROXY_NEGOTIATING</td>
<td class="value">4</td>
<td class="doc">The client is negotiating
  with a proxy to connect to the destination server.</td>
</tr>
<tr>
<td class="name">SOCKET_CLIENT_PROXY_NEGOTIATED</td>
<td class="value">5</td>
<td class="doc">The client has negotiated
  with the proxy server.</td>
</tr>
<tr>
<td class="name">SOCKET_CLIENT_TLS_HANDSHAKING</td>
<td class="value">6</td>
<td class="doc">The client is performing a
  TLS handshake.</td>
</tr>
<tr>
<td class="name">SOCKET_CLIENT_TLS_HANDSHAKED</td>
<td class="value">7</td>
<td class="doc">The client has performed a
  TLS handshake.</td>
</tr>
<tr>
<td class="name">SOCKET_CLIENT_COMPLETE</td>
<td class="value">8</td>
<td class="doc">The client is done with a particular
  #GSocketConnectable.</td>
</tr>
</table>
</div>
<p class="api-heading">SocketFamily</p>
<p class="api-doc">The protocol family of a #GSocketAddress. (These values are
identical to the system defines %AF_INET, %AF_INET6 and %AF_UNIX,
if available.)</p>
<div class="api-notes">
  <p class="api-ctype">GSocketFamily</p>
  <p class="api-since">since 2.22</p>
<table>
<tr>
<td class="name">SOCKET_FAMILY_INVALID</td>
<td class="value">0</td>
<td class="doc">no address family</td>
</tr>
<tr>
<td class="name">SOCKET_FAMILY_UNIX</td>
<td class="value">1</td>
<td class="doc">the UNIX domain family</td>
</tr>
<tr>
<td class="name">SOCKET_FAMILY_IPV4</td>
<td class="value">2</td>
<td class="doc">the IPv4 family</td>
</tr>
<tr>
<td class="name">SOCKET_FAMILY_IPV6</td>
<td class="value">10</td>
<td class="doc">the IPv6 family</td>
</tr>
</table>
</div>
<p class="api-heading">SocketListenerEvent</p>
<p class="api-doc">Describes an event occurring on a #GSocketListener. See the
#GSocketListener::event signal for more details.

Additional values may be added to this type in the future.</p>
<div class="api-notes">
  <p class="api-ctype">GSocketListenerEvent</p>
  <p class="api-since">since 2.46</p>
<table>
<tr>
<td class="name">SOCKET_LISTENER_BINDING</td>
<td class="value">0</td>
<td class="doc">The listener is about to bind a socket.</td>
</tr>
<tr>
<td class="name">SOCKET_LISTENER_BOUND</td>
<td class="value">1</td>
<td class="doc">The listener has bound a socket.</td>
</tr>
<tr>
<td class="name">SOCKET_LISTENER_LISTENING</td>
<td class="value">2</td>
<td class="doc">The listener is about to start
   listening on this socket.</td>
</tr>
<tr>
<td class="name">SOCKET_LISTENER_LISTENED</td>
<td class="value">3</td>
<td class="doc">The listener is now listening on
  this socket.</td>
</tr>
</table>
</div>
<p class="api-heading">SocketProtocol</p>
<p class="api-doc">A protocol identifier is specified when creating a #GSocket, which is a
family/type specific identifier, where 0 means the default protocol for
the particular family/type.

This enum contains a set of commonly available and used protocols. You
can also pass any other identifiers handled by the platform in order to
use protocols not listed here.</p>
<div class="api-notes">
  <p class="api-ctype">GSocketProtocol</p>
  <p class="api-since">since 2.22</p>
<table>
<tr>
<td class="name">SOCKET_PROTOCOL_UNKNOWN</td>
<td class="value">-1</td>
<td class="doc">The protocol type is unknown</td>
</tr>
<tr>
<td class="name">SOCKET_PROTOCOL_DEFAULT</td>
<td class="value">0</td>
<td class="doc">The default protocol for the family/type</td>
</tr>
<tr>
<td class="name">SOCKET_PROTOCOL_TCP</td>
<td class="value">6</td>
<td class="doc">TCP over IP</td>
</tr>
<tr>
<td class="name">SOCKET_PROTOCOL_UDP</td>
<td class="value">17</td>
<td class="doc">UDP over IP</td>
</tr>
<tr>
<td class="name">SOCKET_PROTOCOL_SCTP</td>
<td class="value">132</td>
<td class="doc">SCTP over IP</td>
</tr>
</table>
</div>
<p class="api-heading">SocketType</p>
<p class="api-doc">Flags used when creating a #GSocket. Some protocols may not implement
all the socket types.</p>
<div class="api-notes">
  <p class="api-ctype">GSocketType</p>
  <p class="api-since">since 2.22</p>
<table>
<tr>
<td class="name">SOCKET_TYPE_INVALID</td>
<td class="value">0</td>
<td class="doc">Type unknown or wrong</td>
</tr>
<tr>
<td class="name">SOCKET_TYPE_STREAM</td>
<td class="value">1</td>
<td class="doc">Reliable connection-based byte streams (e.g. TCP).</td>
</tr>
<tr>
<td class="name">SOCKET_TYPE_DATAGRAM</td>
<td class="value">2</td>
<td class="doc">Connectionless, unreliable datagram passing.
    (e.g. UDP)</td>
</tr>
<tr>
<td class="name">SOCKET_TYPE_SEQPACKET</td>
<td class="value">3</td>
<td class="doc">Reliable connection-based passing of datagrams
    of fixed maximum length (e.g. SCTP).</td>
</tr>
</table>
</div>
<p class="api-heading">TlsAuthenticationMode</p>
<p class="api-doc">The client authentication mode for a #GTlsServerConnection.</p>
<div class="api-notes">
  <p class="api-ctype">GTlsAuthenticationMode</p>
  <p class="api-since">since 2.28</p>
<table>
<tr>
<td class="name">TLS_AUTHENTICATION_NONE</td>
<td class="value">0</td>
<td class="doc">client authentication not required</td>
</tr>
<tr>
<td class="name">TLS_AUTHENTICATION_REQUESTED</td>
<td class="value">1</td>
<td class="doc">client authentication is requested</td>
</tr>
<tr>
<td class="name">TLS_AUTHENTICATION_REQUIRED</td>
<td class="value">2</td>
<td class="doc">client authentication is required</td>
</tr>
</table>
</div>
<p class="api-heading">TlsCertificateRequestFlags</p>
<p class="api-doc">Flags for g_tls_interaction_request_certificate(),
g_tls_interaction_request_certificate_async(), and
g_tls_interaction_invoke_request_certificate().</p>
<div class="api-notes">
  <p class="api-ctype">GTlsCertificateRequestFlags</p>
  <p class="api-since">since 2.40</p>
<table>
<tr>
<td class="name">TLS_CERTIFICATE_REQUEST_NONE</td>
<td class="value">0</td>
<td class="doc">No flags</td>
</tr>
</table>
</div>
<p class="api-heading">TlsDatabaseLookupFlags</p>
<p class="api-doc">Flags for g_tls_database_lookup_certificate_for_handle(),
g_tls_database_lookup_certificate_issuer(),
and g_tls_database_lookup_certificates_issued_by().</p>
<div class="api-notes">
  <p class="api-ctype">GTlsDatabaseLookupFlags</p>
  <p class="api-since">since 2.30</p>
<table>
<tr>
<td class="name">TLS_DATABASE_LOOKUP_NONE</td>
<td class="value">0</td>
<td class="doc">No lookup flags</td>
</tr>
<tr>
<td class="name">TLS_DATABASE_LOOKUP_KEYPAIR</td>
<td class="value">1</td>
<td class="doc">Restrict lookup to certificates that have
    a private key.</td>
</tr>
</table>
</div>
<p class="api-heading">TlsError</p>
<p class="api-doc">An error code used with %G_TLS_ERROR in a #GError returned from a
TLS-related routine.</p>
<div class="api-notes">
  <p class="api-ctype">GTlsError</p>
  <p class="api-since">since 2.28</p>
<table>
<tr>
<td class="name">TLS_ERROR_UNAVAILABLE</td>
<td class="value">0</td>
<td class="doc">No TLS provider is available</td>
</tr>
<tr>
<td class="name">TLS_ERROR_MISC</td>
<td class="value">1</td>
<td class="doc">Miscellaneous TLS error</td>
</tr>
<tr>
<td class="name">TLS_ERROR_BAD_CERTIFICATE</td>
<td class="value">2</td>
<td class="doc">A certificate could not be parsed</td>
</tr>
<tr>
<td class="name">TLS_ERROR_NOT_TLS</td>
<td class="value">3</td>
<td class="doc">The TLS handshake failed because the
  peer does not seem to be a TLS server.</td>
</tr>
<tr>
<td class="name">TLS_ERROR_HANDSHAKE</td>
<td class="value">4</td>
<td class="doc">The TLS handshake failed because the
  peer's certificate was not acceptable.</td>
</tr>
<tr>
<td class="name">TLS_ERROR_CERTIFICATE_REQUIRED</td>
<td class="value">5</td>
<td class="doc">The TLS handshake failed because
  the server requested a client-side certificate, but none was
  provided. See g_tls_connection_set_certificate().</td>
</tr>
<tr>
<td class="name">TLS_ERROR_EOF</td>
<td class="value">6</td>
<td class="doc">The TLS connection was closed without proper
  notice, which may indicate an attack. See
  g_tls_connection_set_require_close_notify().</td>
</tr>
</table>
</div>
<p class="api-heading">TlsInteractionResult</p>
<p class="api-doc">#GTlsInteractionResult is returned by various functions in #GTlsInteraction
when finishing an interaction request.</p>
<div class="api-notes">
  <p class="api-ctype">GTlsInteractionResult</p>
  <p class="api-since">since 2.30</p>
<table>
<tr>
<td class="name">TLS_INTERACTION_UNHANDLED</td>
<td class="value">0</td>
<td class="doc">The interaction was unhandled (i.e. not
    implemented).</td>
</tr>
<tr>
<td class="name">TLS_INTERACTION_HANDLED</td>
<td class="value">1</td>
<td class="doc">The interaction completed, and resulting data
    is available.</td>
</tr>
<tr>
<td class="name">TLS_INTERACTION_FAILED</td>
<td class="value">2</td>
<td class="doc">The interaction has failed, or was cancelled.
    and the operation should be aborted.</td>
</tr>
</table>
</div>
<p class="api-heading">TlsRehandshakeMode</p>
<p class="api-doc">When to allow rehandshaking. See
g_tls_connection_set_rehandshake_mode().</p>
<div class="api-notes">
  <p class="api-ctype">GTlsRehandshakeMode</p>
  <p class="api-since">since 2.28</p>
<table>
<tr>
<td class="name">TLS_REHANDSHAKE_NEVER</td>
<td class="value">0</td>
<td class="doc">Never allow rehandshaking</td>
</tr>
<tr>
<td class="name">TLS_REHANDSHAKE_SAFELY</td>
<td class="value">1</td>
<td class="doc">Allow safe rehandshaking only</td>
</tr>
<tr>
<td class="name">TLS_REHANDSHAKE_UNSAFELY</td>
<td class="value">2</td>
<td class="doc">Allow unsafe rehandshaking</td>
</tr>
</table>
</div>
<p class="api-heading">UnixSocketAddressType</p>
<p class="api-doc">The type of name used by a #GUnixSocketAddress.
%G_UNIX_SOCKET_ADDRESS_PATH indicates a traditional unix domain
socket bound to a filesystem path. %G_UNIX_SOCKET_ADDRESS_ANONYMOUS
indicates a socket not bound to any name (eg, a client-side socket,
or a socket created with socketpair()).

For abstract sockets, there are two incompatible ways of naming
them; the man pages suggest using the entire `struct sockaddr_un`
as the name, padding the unused parts of the %sun_path field with
zeroes; this corresponds to %G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED.
However, many programs instead just use a portion of %sun_path, and
pass an appropriate smaller length to bind() or connect(). This is
%G_UNIX_SOCKET_ADDRESS_ABSTRACT.</p>
<div class="api-notes">
  <p class="api-ctype">GUnixSocketAddressType</p>
  <p class="api-since">since 2.26</p>
<table>
<tr>
<td class="name">UNIX_SOCKET_ADDRESS_INVALID</td>
<td class="value">0</td>
<td class="doc">invalid</td>
</tr>
<tr>
<td class="name">UNIX_SOCKET_ADDRESS_ANONYMOUS</td>
<td class="value">1</td>
<td class="doc">anonymous</td>
</tr>
<tr>
<td class="name">UNIX_SOCKET_ADDRESS_PATH</td>
<td class="value">2</td>
<td class="doc">a filesystem path</td>
</tr>
<tr>
<td class="name">UNIX_SOCKET_ADDRESS_ABSTRACT</td>
<td class="value">3</td>
<td class="doc">an abstract name</td>
</tr>
<tr>
<td class="name">UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED</td>
<td class="value">4</td>
<td class="doc">an abstract name, 0-padded
  to the full length of a unix socket name</td>
</tr>
</table>
</div>
<p class="api-heading">ZlibCompressorFormat</p>
<p class="api-doc">Used to select the type of data format to use for #GZlibDecompressor
and #GZlibCompressor.</p>
<div class="api-notes">
  <p class="api-ctype">GZlibCompressorFormat</p>
  <p class="api-since">since 2.24</p>
<table>
<tr>
<td class="name">ZLIB_COMPRESSOR_FORMAT_ZLIB</td>
<td class="value">0</td>
<td class="doc">deflate compression with zlib header</td>
</tr>
<tr>
<td class="name">ZLIB_COMPRESSOR_FORMAT_GZIP</td>
<td class="value">1</td>
<td class="doc">gzip file format</td>
</tr>
<tr>
<td class="name">ZLIB_COMPRESSOR_FORMAT_RAW</td>
<td class="value">2</td>
<td class="doc">deflate compression with no header</td>
</tr>
</table>
</div>