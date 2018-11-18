+++
title = "enums"
+++
<p class="api-heading">BookmarkFileError</p>
<p class="api-doc">Error codes returned by bookmark file parsing.</p>
<div class="api-notes">
  <p class="api-ctype">GBookmarkFileError</p>
<table>
<tr>
<td class="name">BOOKMARK_FILE_ERROR_INVALID_URI</td>
<td class="value">0</td>
<td class="doc">URI was ill-formed</td>
</tr>
<tr>
<td class="name">BOOKMARK_FILE_ERROR_INVALID_VALUE</td>
<td class="value">1</td>
<td class="doc">a requested field was not found</td>
</tr>
<tr>
<td class="name">BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED</td>
<td class="value">2</td>
<td class="doc">a requested application did
    not register a bookmark</td>
</tr>
<tr>
<td class="name">BOOKMARK_FILE_ERROR_URI_NOT_FOUND</td>
<td class="value">3</td>
<td class="doc">a requested URI was not found</td>
</tr>
<tr>
<td class="name">BOOKMARK_FILE_ERROR_READ</td>
<td class="value">4</td>
<td class="doc">document was ill formed</td>
</tr>
<tr>
<td class="name">BOOKMARK_FILE_ERROR_UNKNOWN_ENCODING</td>
<td class="value">5</td>
<td class="doc">the text being parsed was
    in an unknown encoding</td>
</tr>
<tr>
<td class="name">BOOKMARK_FILE_ERROR_WRITE</td>
<td class="value">6</td>
<td class="doc">an error occurred while writing</td>
</tr>
<tr>
<td class="name">BOOKMARK_FILE_ERROR_FILE_NOT_FOUND</td>
<td class="value">7</td>
<td class="doc">requested file was not found</td>
</tr>
</table>
</div>
<p class="api-heading">ChecksumType</p>
<p class="api-doc">The hashing algorithm to be used by #GChecksum when performing the
digest of some data.

Note that the #GChecksumType enumeration may be extended at a later
date to include new hashing algorithm types.</p>
<div class="api-notes">
  <p class="api-ctype">GChecksumType</p>
  <p class="api-since">since 2.16</p>
<table>
<tr>
<td class="name">CHECKSUM_MD5</td>
<td class="value">0</td>
<td class="doc">Use the MD5 hashing algorithm</td>
</tr>
<tr>
<td class="name">CHECKSUM_SHA1</td>
<td class="value">1</td>
<td class="doc">Use the SHA-1 hashing algorithm</td>
</tr>
<tr>
<td class="name">CHECKSUM_SHA256</td>
<td class="value">2</td>
<td class="doc">Use the SHA-256 hashing algorithm</td>
</tr>
<tr>
<td class="name">CHECKSUM_SHA512</td>
<td class="value">3</td>
<td class="doc">Use the SHA-512 hashing algorithm (Since: 2.36)</td>
</tr>
<tr>
<td class="name">CHECKSUM_SHA384</td>
<td class="value">4</td>
<td class="doc">Use the SHA-384 hashing algorithm (Since: 2.51)</td>
</tr>
</table>
</div>
<p class="api-heading">ConvertError</p>
<p class="api-doc">Error codes returned by character set conversion routines.</p>
<div class="api-notes">
  <p class="api-ctype">GConvertError</p>
<table>
<tr>
<td class="name">CONVERT_ERROR_NO_CONVERSION</td>
<td class="value">0</td>
<td class="doc">Conversion between the requested character
    sets is not supported.</td>
</tr>
<tr>
<td class="name">CONVERT_ERROR_ILLEGAL_SEQUENCE</td>
<td class="value">1</td>
<td class="doc">Invalid byte sequence in conversion input;
   or the character sequence could not be represented in the target
   character set.</td>
</tr>
<tr>
<td class="name">CONVERT_ERROR_FAILED</td>
<td class="value">2</td>
<td class="doc">Conversion failed for some reason.</td>
</tr>
<tr>
<td class="name">CONVERT_ERROR_PARTIAL_INPUT</td>
<td class="value">3</td>
<td class="doc">Partial character sequence at end of input.</td>
</tr>
<tr>
<td class="name">CONVERT_ERROR_BAD_URI</td>
<td class="value">4</td>
<td class="doc">URI is invalid.</td>
</tr>
<tr>
<td class="name">CONVERT_ERROR_NOT_ABSOLUTE_PATH</td>
<td class="value">5</td>
<td class="doc">Pathname is not an absolute path.</td>
</tr>
<tr>
<td class="name">CONVERT_ERROR_NO_MEMORY</td>
<td class="value">6</td>
<td class="doc">No memory available. Since: 2.40</td>
</tr>
<tr>
<td class="name">CONVERT_ERROR_EMBEDDED_NUL</td>
<td class="value">7</td>
<td class="doc">An embedded NUL character is present in
    conversion output where a NUL-terminated string is expected.
    Since: 2.56</td>
</tr>
</table>
</div>
<p class="api-heading">DateDMY</p>
<p class="api-doc">This enumeration isn't used in the API, but may be useful if you need
to mark a number as a day, month, or year.</p>
<div class="api-notes">
  <p class="api-ctype">GDateDMY</p>
<table>
<tr>
<td class="name">DATE_DAY</td>
<td class="value">0</td>
<td class="doc">a day</td>
</tr>
<tr>
<td class="name">DATE_MONTH</td>
<td class="value">1</td>
<td class="doc">a month</td>
</tr>
<tr>
<td class="name">DATE_YEAR</td>
<td class="value">2</td>
<td class="doc">a year</td>
</tr>
</table>
</div>
<p class="api-heading">DateMonth</p>
<p class="api-doc">Enumeration representing a month; values are #G_DATE_JANUARY,
#G_DATE_FEBRUARY, etc. #G_DATE_BAD_MONTH is the invalid value.</p>
<div class="api-notes">
  <p class="api-ctype">GDateMonth</p>
<table>
<tr>
<td class="name">DATE_BAD_MONTH</td>
<td class="value">0</td>
<td class="doc">invalid value</td>
</tr>
<tr>
<td class="name">DATE_JANUARY</td>
<td class="value">1</td>
<td class="doc">January</td>
</tr>
<tr>
<td class="name">DATE_FEBRUARY</td>
<td class="value">2</td>
<td class="doc">February</td>
</tr>
<tr>
<td class="name">DATE_MARCH</td>
<td class="value">3</td>
<td class="doc">March</td>
</tr>
<tr>
<td class="name">DATE_APRIL</td>
<td class="value">4</td>
<td class="doc">April</td>
</tr>
<tr>
<td class="name">DATE_MAY</td>
<td class="value">5</td>
<td class="doc">May</td>
</tr>
<tr>
<td class="name">DATE_JUNE</td>
<td class="value">6</td>
<td class="doc">June</td>
</tr>
<tr>
<td class="name">DATE_JULY</td>
<td class="value">7</td>
<td class="doc">July</td>
</tr>
<tr>
<td class="name">DATE_AUGUST</td>
<td class="value">8</td>
<td class="doc">August</td>
</tr>
<tr>
<td class="name">DATE_SEPTEMBER</td>
<td class="value">9</td>
<td class="doc">September</td>
</tr>
<tr>
<td class="name">DATE_OCTOBER</td>
<td class="value">10</td>
<td class="doc">October</td>
</tr>
<tr>
<td class="name">DATE_NOVEMBER</td>
<td class="value">11</td>
<td class="doc">November</td>
</tr>
<tr>
<td class="name">DATE_DECEMBER</td>
<td class="value">12</td>
<td class="doc">December</td>
</tr>
</table>
</div>
<p class="api-heading">DateWeekday</p>
<p class="api-doc">Enumeration representing a day of the week; #G_DATE_MONDAY,
#G_DATE_TUESDAY, etc. #G_DATE_BAD_WEEKDAY is an invalid weekday.</p>
<div class="api-notes">
  <p class="api-ctype">GDateWeekday</p>
<table>
<tr>
<td class="name">DATE_BAD_WEEKDAY</td>
<td class="value">0</td>
<td class="doc">invalid value</td>
</tr>
<tr>
<td class="name">DATE_MONDAY</td>
<td class="value">1</td>
<td class="doc">Monday</td>
</tr>
<tr>
<td class="name">DATE_TUESDAY</td>
<td class="value">2</td>
<td class="doc">Tuesday</td>
</tr>
<tr>
<td class="name">DATE_WEDNESDAY</td>
<td class="value">3</td>
<td class="doc">Wednesday</td>
</tr>
<tr>
<td class="name">DATE_THURSDAY</td>
<td class="value">4</td>
<td class="doc">Thursday</td>
</tr>
<tr>
<td class="name">DATE_FRIDAY</td>
<td class="value">5</td>
<td class="doc">Friday</td>
</tr>
<tr>
<td class="name">DATE_SATURDAY</td>
<td class="value">6</td>
<td class="doc">Saturday</td>
</tr>
<tr>
<td class="name">DATE_SUNDAY</td>
<td class="value">7</td>
<td class="doc">Sunday</td>
</tr>
</table>
</div>
<p class="api-heading">ErrorType</p>
<p class="api-doc">The possible errors, used in the @v_error field
of #GTokenValue, when the token is a %G_TOKEN_ERROR.</p>
<div class="api-notes">
  <p class="api-ctype">GErrorType</p>
<table>
<tr>
<td class="name">ERR_UNKNOWN</td>
<td class="value">0</td>
<td class="doc">unknown error</td>
</tr>
<tr>
<td class="name">ERR_UNEXP_EOF</td>
<td class="value">1</td>
<td class="doc">unexpected end of file</td>
</tr>
<tr>
<td class="name">ERR_UNEXP_EOF_IN_STRING</td>
<td class="value">2</td>
<td class="doc">unterminated string constant</td>
</tr>
<tr>
<td class="name">ERR_UNEXP_EOF_IN_COMMENT</td>
<td class="value">3</td>
<td class="doc">unterminated comment</td>
</tr>
<tr>
<td class="name">ERR_NON_DIGIT_IN_CONST</td>
<td class="value">4</td>
<td class="doc">non-digit character in a number</td>
</tr>
<tr>
<td class="name">ERR_DIGIT_RADIX</td>
<td class="value">5</td>
<td class="doc">digit beyond radix in a number</td>
</tr>
<tr>
<td class="name">ERR_FLOAT_RADIX</td>
<td class="value">6</td>
<td class="doc">non-decimal floating point number</td>
</tr>
<tr>
<td class="name">ERR_FLOAT_MALFORMED</td>
<td class="value">7</td>
<td class="doc">malformed floating point number</td>
</tr>
</table>
</div>
<p class="api-heading">FileError</p>
<p class="api-doc">Values corresponding to @errno codes returned from file operations
on UNIX. Unlike @errno codes, GFileError values are available on
all systems, even Windows. The exact meaning of each code depends
on what sort of file operation you were performing; the UNIX
documentation gives more details. The following error code descriptions
come from the GNU C Library manual, and are under the copyright
of that manual.

It's not very portable to make detailed assumptions about exactly
which errors will be returned from a given operation. Some errors
don't occur on some systems, etc., sometimes there are subtle
differences in when a system will report a given error, etc.</p>
<div class="api-notes">
  <p class="api-ctype">GFileError</p>
<table>
<tr>
<td class="name">FILE_ERROR_EXIST</td>
<td class="value">0</td>
<td class="doc">Operation not permitted; only the owner of
    the file (or other resource) or processes with special privileges
    can perform the operation.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_ISDIR</td>
<td class="value">1</td>
<td class="doc">File is a directory; you cannot open a directory
    for writing, or create or remove hard links to it.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_ACCES</td>
<td class="value">2</td>
<td class="doc">Permission denied; the file permissions do not
    allow the attempted operation.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_NAMETOOLONG</td>
<td class="value">3</td>
<td class="doc">Filename too long.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_NOENT</td>
<td class="value">4</td>
<td class="doc">No such file or directory. This is a "file
    doesn't exist" error for ordinary files that are referenced in
    contexts where they are expected to already exist.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_NOTDIR</td>
<td class="value">5</td>
<td class="doc">A file that isn't a directory was specified when
    a directory is required.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_NXIO</td>
<td class="value">6</td>
<td class="doc">No such device or address. The system tried to
    use the device represented by a file you specified, and it
    couldn't find the device. This can mean that the device file was
    installed incorrectly, or that the physical device is missing or
    not correctly attached to the computer.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_NODEV</td>
<td class="value">7</td>
<td class="doc">The underlying file system of the specified file
    does not support memory mapping.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_ROFS</td>
<td class="value">8</td>
<td class="doc">The directory containing the new link can't be
    modified because it's on a read-only file system.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_TXTBSY</td>
<td class="value">9</td>
<td class="doc">Text file busy.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_FAULT</td>
<td class="value">10</td>
<td class="doc">You passed in a pointer to bad memory.
    (GLib won't reliably return this, don't pass in pointers to bad
    memory.)</td>
</tr>
<tr>
<td class="name">FILE_ERROR_LOOP</td>
<td class="value">11</td>
<td class="doc">Too many levels of symbolic links were encountered
    in looking up a file name. This often indicates a cycle of symbolic
    links.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_NOSPC</td>
<td class="value">12</td>
<td class="doc">No space left on device; write operation on a
    file failed because the disk is full.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_NOMEM</td>
<td class="value">13</td>
<td class="doc">No memory available. The system cannot allocate
    more virtual memory because its capacity is full.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_MFILE</td>
<td class="value">14</td>
<td class="doc">The current process has too many files open and
    can't open any more. Duplicate descriptors do count toward this
    limit.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_NFILE</td>
<td class="value">15</td>
<td class="doc">There are too many distinct file openings in the
    entire system.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_BADF</td>
<td class="value">16</td>
<td class="doc">Bad file descriptor; for example, I/O on a
    descriptor that has been closed or reading from a descriptor open
    only for writing (or vice versa).</td>
</tr>
<tr>
<td class="name">FILE_ERROR_INVAL</td>
<td class="value">17</td>
<td class="doc">Invalid argument. This is used to indicate
    various kinds of problems with passing the wrong argument to a
    library function.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_PIPE</td>
<td class="value">18</td>
<td class="doc">Broken pipe; there is no process reading from the
    other end of a pipe. Every library function that returns this
    error code also generates a 'SIGPIPE' signal; this signal
    terminates the program if not handled or blocked. Thus, your
    program will never actually see this code unless it has handled
    or blocked 'SIGPIPE'.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_AGAIN</td>
<td class="value">19</td>
<td class="doc">Resource temporarily unavailable; the call might
    work if you try again later.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_INTR</td>
<td class="value">20</td>
<td class="doc">Interrupted function call; an asynchronous signal
    occurred and prevented completion of the call. When this
    happens, you should try the call again.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_IO</td>
<td class="value">21</td>
<td class="doc">Input/output error; usually used for physical read
   or write errors. i.e. the disk or other physical device hardware
   is returning errors.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_PERM</td>
<td class="value">22</td>
<td class="doc">Operation not permitted; only the owner of the
   file (or other resource) or processes with special privileges can
   perform the operation.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_NOSYS</td>
<td class="value">23</td>
<td class="doc">Function not implemented; this indicates that
   the system is missing some functionality.</td>
</tr>
<tr>
<td class="name">FILE_ERROR_FAILED</td>
<td class="value">24</td>
<td class="doc">Does not correspond to a UNIX error code; this
   is the standard "failed for unspecified reason" error code present
   in all #GError error code enumerations. Returned if no specific
   code applies.</td>
</tr>
</table>
</div>
<p class="api-heading">IOChannelError</p>
<p class="api-doc">Error codes returned by #GIOChannel operations.</p>
<div class="api-notes">
  <p class="api-ctype">GIOChannelError</p>
<table>
<tr>
<td class="name">IO_CHANNEL_ERROR_FBIG</td>
<td class="value">0</td>
<td class="doc">File too large.</td>
</tr>
<tr>
<td class="name">IO_CHANNEL_ERROR_INVAL</td>
<td class="value">1</td>
<td class="doc">Invalid argument.</td>
</tr>
<tr>
<td class="name">IO_CHANNEL_ERROR_IO</td>
<td class="value">2</td>
<td class="doc">IO error.</td>
</tr>
<tr>
<td class="name">IO_CHANNEL_ERROR_ISDIR</td>
<td class="value">3</td>
<td class="doc">File is a directory.</td>
</tr>
<tr>
<td class="name">IO_CHANNEL_ERROR_NOSPC</td>
<td class="value">4</td>
<td class="doc">No space left on device.</td>
</tr>
<tr>
<td class="name">IO_CHANNEL_ERROR_NXIO</td>
<td class="value">5</td>
<td class="doc">No such device or address.</td>
</tr>
<tr>
<td class="name">IO_CHANNEL_ERROR_OVERFLOW</td>
<td class="value">6</td>
<td class="doc">Value too large for defined datatype.</td>
</tr>
<tr>
<td class="name">IO_CHANNEL_ERROR_PIPE</td>
<td class="value">7</td>
<td class="doc">Broken pipe.</td>
</tr>
<tr>
<td class="name">IO_CHANNEL_ERROR_FAILED</td>
<td class="value">8</td>
<td class="doc">Some other error.</td>
</tr>
</table>
</div>
<p class="api-heading">IOError</p>
<p class="api-doc">#GIOError is only used by the deprecated functions
g_io_channel_read(), g_io_channel_write(), and g_io_channel_seek().</p>
<div class="api-notes">
  <p class="api-ctype">GIOError</p>
<table>
<tr>
<td class="name">IO_ERROR_NONE</td>
<td class="value">0</td>
<td class="doc">no error</td>
</tr>
<tr>
<td class="name">IO_ERROR_AGAIN</td>
<td class="value">1</td>
<td class="doc">an EAGAIN error occurred</td>
</tr>
<tr>
<td class="name">IO_ERROR_INVAL</td>
<td class="value">2</td>
<td class="doc">an EINVAL error occurred</td>
</tr>
<tr>
<td class="name">IO_ERROR_UNKNOWN</td>
<td class="value">3</td>
<td class="doc">another error occurred</td>
</tr>
</table>
</div>
<p class="api-heading">IOStatus</p>
<p class="api-doc">Stati returned by most of the #GIOFuncs functions.</p>
<div class="api-notes">
  <p class="api-ctype">GIOStatus</p>
<table>
<tr>
<td class="name">IO_STATUS_ERROR</td>
<td class="value">0</td>
<td class="doc">An error occurred.</td>
</tr>
<tr>
<td class="name">IO_STATUS_NORMAL</td>
<td class="value">1</td>
<td class="doc">Success.</td>
</tr>
<tr>
<td class="name">IO_STATUS_EOF</td>
<td class="value">2</td>
<td class="doc">End of file.</td>
</tr>
<tr>
<td class="name">IO_STATUS_AGAIN</td>
<td class="value">3</td>
<td class="doc">Resource temporarily unavailable.</td>
</tr>
</table>
</div>
<p class="api-heading">KeyFileError</p>
<p class="api-doc">Error codes returned by key file parsing.</p>
<div class="api-notes">
  <p class="api-ctype">GKeyFileError</p>
<table>
<tr>
<td class="name">KEY_FILE_ERROR_UNKNOWN_ENCODING</td>
<td class="value">0</td>
<td class="doc">the text being parsed was in
    an unknown encoding</td>
</tr>
<tr>
<td class="name">KEY_FILE_ERROR_PARSE</td>
<td class="value">1</td>
<td class="doc">document was ill-formed</td>
</tr>
<tr>
<td class="name">KEY_FILE_ERROR_NOT_FOUND</td>
<td class="value">2</td>
<td class="doc">the file was not found</td>
</tr>
<tr>
<td class="name">KEY_FILE_ERROR_KEY_NOT_FOUND</td>
<td class="value">3</td>
<td class="doc">a requested key was not found</td>
</tr>
<tr>
<td class="name">KEY_FILE_ERROR_GROUP_NOT_FOUND</td>
<td class="value">4</td>
<td class="doc">a requested group was not found</td>
</tr>
<tr>
<td class="name">KEY_FILE_ERROR_INVALID_VALUE</td>
<td class="value">5</td>
<td class="doc">a value could not be parsed</td>
</tr>
</table>
</div>
<p class="api-heading">LogWriterOutput</p>
<p class="api-doc">Return values from #GLogWriterFuncs to indicate whether the given log entry
was successfully handled by the writer, or whether there was an error in
handling it (and hence a fallback writer should be used).

If a #GLogWriterFunc ignores a log entry, it should return
%G_LOG_WRITER_HANDLED.</p>
<div class="api-notes">
  <p class="api-ctype">GLogWriterOutput</p>
  <p class="api-since">since 2.50</p>
<table>
<tr>
<td class="name">LOG_WRITER_HANDLED</td>
<td class="value">1</td>
<td class="doc">Log writer has handled the log entry.</td>
</tr>
<tr>
<td class="name">LOG_WRITER_UNHANDLED</td>
<td class="value">0</td>
<td class="doc">Log writer could not handle the log entry.</td>
</tr>
</table>
</div>
<p class="api-heading">MarkupError</p>
<p class="api-doc">Error codes returned by markup parsing.</p>
<div class="api-notes">
  <p class="api-ctype">GMarkupError</p>
<table>
<tr>
<td class="name">MARKUP_ERROR_BAD_UTF8</td>
<td class="value">0</td>
<td class="doc">text being parsed was not valid UTF-8</td>
</tr>
<tr>
<td class="name">MARKUP_ERROR_EMPTY</td>
<td class="value">1</td>
<td class="doc">document contained nothing, or only whitespace</td>
</tr>
<tr>
<td class="name">MARKUP_ERROR_PARSE</td>
<td class="value">2</td>
<td class="doc">document was ill-formed</td>
</tr>
<tr>
<td class="name">MARKUP_ERROR_UNKNOWN_ELEMENT</td>
<td class="value">3</td>
<td class="doc">error should be set by #GMarkupParser
    functions; element wasn't known</td>
</tr>
<tr>
<td class="name">MARKUP_ERROR_UNKNOWN_ATTRIBUTE</td>
<td class="value">4</td>
<td class="doc">error should be set by #GMarkupParser
    functions; attribute wasn't known</td>
</tr>
<tr>
<td class="name">MARKUP_ERROR_INVALID_CONTENT</td>
<td class="value">5</td>
<td class="doc">error should be set by #GMarkupParser
    functions; content was invalid</td>
</tr>
<tr>
<td class="name">MARKUP_ERROR_MISSING_ATTRIBUTE</td>
<td class="value">6</td>
<td class="doc">error should be set by #GMarkupParser
    functions; a required attribute was missing</td>
</tr>
</table>
</div>
<p class="api-heading">NormalizeMode</p>
<p class="api-doc">Defines how a Unicode string is transformed in a canonical
form, standardizing such issues as whether a character with
an accent is represented as a base character and combining
accent or as a single precomposed character. Unicode strings
should generally be normalized before comparing them.</p>
<div class="api-notes">
  <p class="api-ctype">GNormalizeMode</p>
<table>
<tr>
<td class="name">NORMALIZE_DEFAULT</td>
<td class="value">0</td>
<td class="doc">standardize differences that do not affect the
    text content, such as the above-mentioned accent representation</td>
</tr>
<tr>
<td class="name">NORMALIZE_NFD</td>
<td class="value">0</td>
<td class="doc">another name for %G_NORMALIZE_DEFAULT</td>
</tr>
<tr>
<td class="name">NORMALIZE_DEFAULT_COMPOSE</td>
<td class="value">1</td>
<td class="doc">like %G_NORMALIZE_DEFAULT, but with
    composed forms rather than a maximally decomposed form</td>
</tr>
<tr>
<td class="name">NORMALIZE_NFC</td>
<td class="value">1</td>
<td class="doc">another name for %G_NORMALIZE_DEFAULT_COMPOSE</td>
</tr>
<tr>
<td class="name">NORMALIZE_ALL</td>
<td class="value">2</td>
<td class="doc">beyond %G_NORMALIZE_DEFAULT also standardize the
    "compatibility" characters in Unicode, such as SUPERSCRIPT THREE
    to the standard forms (in this case DIGIT THREE). Formatting
    information may be lost but for most text operations such
    characters should be considered the same</td>
</tr>
<tr>
<td class="name">NORMALIZE_NFKD</td>
<td class="value">2</td>
<td class="doc">another name for %G_NORMALIZE_ALL</td>
</tr>
<tr>
<td class="name">NORMALIZE_ALL_COMPOSE</td>
<td class="value">3</td>
<td class="doc">like %G_NORMALIZE_ALL, but with composed
    forms rather than a maximally decomposed form</td>
</tr>
<tr>
<td class="name">NORMALIZE_NFKC</td>
<td class="value">3</td>
<td class="doc">another name for %G_NORMALIZE_ALL_COMPOSE</td>
</tr>
</table>
</div>
<p class="api-heading">NumberParserError</p>
<p class="api-doc">Error codes returned by functions converting a string to a number.</p>
<div class="api-notes">
  <p class="api-ctype">GNumberParserError</p>
  <p class="api-since">since 2.54</p>
<table>
<tr>
<td class="name">NUMBER_PARSER_ERROR_INVALID</td>
<td class="value">0</td>
<td class="doc">String was not a valid number.</td>
</tr>
<tr>
<td class="name">NUMBER_PARSER_ERROR_OUT_OF_BOUNDS</td>
<td class="value">1</td>
<td class="doc">String was a number, but out of bounds.</td>
</tr>
</table>
</div>
<p class="api-heading">OnceStatus</p>
<p class="api-doc">The possible statuses of a one-time initialization function
controlled by a #GOnce struct.</p>
<div class="api-notes">
  <p class="api-ctype">GOnceStatus</p>
  <p class="api-since">since 2.4</p>
<table>
<tr>
<td class="name">ONCE_STATUS_NOTCALLED</td>
<td class="value">0</td>
<td class="doc">the function has not been called yet.</td>
</tr>
<tr>
<td class="name">ONCE_STATUS_PROGRESS</td>
<td class="value">1</td>
<td class="doc">the function call is currently in progress.</td>
</tr>
<tr>
<td class="name">ONCE_STATUS_READY</td>
<td class="value">2</td>
<td class="doc">the function has been called.</td>
</tr>
</table>
</div>
<p class="api-heading">OptionArg</p>
<p class="api-doc">The #GOptionArg enum values determine which type of extra argument the
options expect to find. If an option expects an extra argument, it can
be specified in several ways; with a short option: `-x arg`, with a long
option: `--name arg` or combined in a single argument: `--name=arg`.</p>
<div class="api-notes">
  <p class="api-ctype">GOptionArg</p>
<table>
<tr>
<td class="name">OPTION_ARG_NONE</td>
<td class="value">0</td>
<td class="doc">No extra argument. This is useful for simple flags.</td>
</tr>
<tr>
<td class="name">OPTION_ARG_STRING</td>
<td class="value">1</td>
<td class="doc">The option takes a string argument.</td>
</tr>
<tr>
<td class="name">OPTION_ARG_INT</td>
<td class="value">2</td>
<td class="doc">The option takes an integer argument.</td>
</tr>
<tr>
<td class="name">OPTION_ARG_CALLBACK</td>
<td class="value">3</td>
<td class="doc">The option provides a callback (of type
    #GOptionArgFunc) to parse the extra argument.</td>
</tr>
<tr>
<td class="name">OPTION_ARG_FILENAME</td>
<td class="value">4</td>
<td class="doc">The option takes a filename as argument.</td>
</tr>
<tr>
<td class="name">OPTION_ARG_STRING_ARRAY</td>
<td class="value">5</td>
<td class="doc">The option takes a string argument, multiple
    uses of the option are collected into an array of strings.</td>
</tr>
<tr>
<td class="name">OPTION_ARG_FILENAME_ARRAY</td>
<td class="value">6</td>
<td class="doc">The option takes a filename as argument,
    multiple uses of the option are collected into an array of strings.</td>
</tr>
<tr>
<td class="name">OPTION_ARG_DOUBLE</td>
<td class="value">7</td>
<td class="doc">The option takes a double argument. The argument
    can be formatted either for the user's locale or for the "C" locale.
    Since 2.12</td>
</tr>
<tr>
<td class="name">OPTION_ARG_INT64</td>
<td class="value">8</td>
<td class="doc">The option takes a 64-bit integer. Like
    %G_OPTION_ARG_INT but for larger numbers. The number can be in
    decimal base, or in hexadecimal (when prefixed with `0x`, for
    example, `0xffffffff`). Since 2.12</td>
</tr>
</table>
</div>
<p class="api-heading">OptionError</p>
<p class="api-doc">Error codes returned by option parsing.</p>
<div class="api-notes">
  <p class="api-ctype">GOptionError</p>
<table>
<tr>
<td class="name">OPTION_ERROR_UNKNOWN_OPTION</td>
<td class="value">0</td>
<td class="doc">An option was not known to the parser.
 This error will only be reported, if the parser hasn't been instructed
 to ignore unknown options, see g_option_context_set_ignore_unknown_options().</td>
</tr>
<tr>
<td class="name">OPTION_ERROR_BAD_VALUE</td>
<td class="value">1</td>
<td class="doc">A value couldn't be parsed.</td>
</tr>
<tr>
<td class="name">OPTION_ERROR_FAILED</td>
<td class="value">2</td>
<td class="doc">A #GOptionArgFunc callback failed.</td>
</tr>
</table>
</div>
<p class="api-heading">RegexError</p>
<p class="api-doc">Error codes returned by regular expressions functions.</p>
<div class="api-notes">
  <p class="api-ctype">GRegexError</p>
  <p class="api-since">since 2.14</p>
<table>
<tr>
<td class="name">REGEX_ERROR_COMPILE</td>
<td class="value">0</td>
<td class="doc">Compilation of the regular expression failed.</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_OPTIMIZE</td>
<td class="value">1</td>
<td class="doc">Optimization of the regular expression failed.</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_REPLACE</td>
<td class="value">2</td>
<td class="doc">Replacement failed due to an ill-formed replacement
    string.</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_MATCH</td>
<td class="value">3</td>
<td class="doc">The match process failed.</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_INTERNAL</td>
<td class="value">4</td>
<td class="doc">Internal error of the regular expression engine.
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_STRAY_BACKSLASH</td>
<td class="value">101</td>
<td class="doc">"\\" at end of pattern. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_MISSING_CONTROL_CHAR</td>
<td class="value">102</td>
<td class="doc">"\\c" at end of pattern. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_UNRECOGNIZED_ESCAPE</td>
<td class="value">103</td>
<td class="doc">Unrecognized character follows "\\".
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_QUANTIFIERS_OUT_OF_ORDER</td>
<td class="value">104</td>
<td class="doc">Numbers out of order in "{}"
    quantifier. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_QUANTIFIER_TOO_BIG</td>
<td class="value">105</td>
<td class="doc">Number too big in "{}" quantifier.
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_UNTERMINATED_CHARACTER_CLASS</td>
<td class="value">106</td>
<td class="doc">Missing terminating "]" for
    character class. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_INVALID_ESCAPE_IN_CHARACTER_CLASS</td>
<td class="value">107</td>
<td class="doc">Invalid escape sequence
    in character class. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_RANGE_OUT_OF_ORDER</td>
<td class="value">108</td>
<td class="doc">Range out of order in character class.
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_NOTHING_TO_REPEAT</td>
<td class="value">109</td>
<td class="doc">Nothing to repeat. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_UNRECOGNIZED_CHARACTER</td>
<td class="value">112</td>
<td class="doc">Unrecognized character after "(?",
    "(?<" or "(?P". Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_POSIX_NAMED_CLASS_OUTSIDE_CLASS</td>
<td class="value">113</td>
<td class="doc">POSIX named classes are
    supported only within a class. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_UNMATCHED_PARENTHESIS</td>
<td class="value">114</td>
<td class="doc">Missing terminating ")" or ")"
    without opening "(". Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_INEXISTENT_SUBPATTERN_REFERENCE</td>
<td class="value">115</td>
<td class="doc">Reference to non-existent
    subpattern. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_UNTERMINATED_COMMENT</td>
<td class="value">118</td>
<td class="doc">Missing terminating ")" after comment.
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_EXPRESSION_TOO_LARGE</td>
<td class="value">120</td>
<td class="doc">Regular expression too large.
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_MEMORY_ERROR</td>
<td class="value">121</td>
<td class="doc">Failed to get memory. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_VARIABLE_LENGTH_LOOKBEHIND</td>
<td class="value">125</td>
<td class="doc">Lookbehind assertion is not
    fixed length. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_MALFORMED_CONDITION</td>
<td class="value">126</td>
<td class="doc">Malformed number or name after "(?(".
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_TOO_MANY_CONDITIONAL_BRANCHES</td>
<td class="value">127</td>
<td class="doc">Conditional group contains
    more than two branches. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_ASSERTION_EXPECTED</td>
<td class="value">128</td>
<td class="doc">Assertion expected after "(?(".
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_UNKNOWN_POSIX_CLASS_NAME</td>
<td class="value">130</td>
<td class="doc">Unknown POSIX class name.
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_POSIX_COLLATING_ELEMENTS_NOT_SUPPORTED</td>
<td class="value">131</td>
<td class="doc">POSIX collating
    elements are not supported. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_HEX_CODE_TOO_LARGE</td>
<td class="value">134</td>
<td class="doc">Character value in "\\x{...}" sequence
    is too large. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_INVALID_CONDITION</td>
<td class="value">135</td>
<td class="doc">Invalid condition "(?(0)". Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_SINGLE_BYTE_MATCH_IN_LOOKBEHIND</td>
<td class="value">136</td>
<td class="doc">\\C not allowed in
    lookbehind assertion. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_INFINITE_LOOP</td>
<td class="value">140</td>
<td class="doc">Recursive call could loop indefinitely.
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_MISSING_SUBPATTERN_NAME_TERMINATOR</td>
<td class="value">142</td>
<td class="doc">Missing terminator
    in subpattern name. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_DUPLICATE_SUBPATTERN_NAME</td>
<td class="value">143</td>
<td class="doc">Two named subpatterns have
    the same name. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_MALFORMED_PROPERTY</td>
<td class="value">146</td>
<td class="doc">Malformed "\\P" or "\\p" sequence.
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_UNKNOWN_PROPERTY</td>
<td class="value">147</td>
<td class="doc">Unknown property name after "\\P" or
    "\\p". Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_SUBPATTERN_NAME_TOO_LONG</td>
<td class="value">148</td>
<td class="doc">Subpattern name is too long
    (maximum 32 characters). Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_TOO_MANY_SUBPATTERNS</td>
<td class="value">149</td>
<td class="doc">Too many named subpatterns (maximum
    10,000). Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_INVALID_OCTAL_VALUE</td>
<td class="value">151</td>
<td class="doc">Octal value is greater than "\\377".
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_TOO_MANY_BRANCHES_IN_DEFINE</td>
<td class="value">154</td>
<td class="doc">"DEFINE" group contains more
    than one branch. Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_DEFINE_REPETION</td>
<td class="value">155</td>
<td class="doc">Repeating a "DEFINE" group is not allowed.
    This error is never raised. Since: 2.16 Deprecated: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_INCONSISTENT_NEWLINE_OPTIONS</td>
<td class="value">156</td>
<td class="doc">Inconsistent newline options.
    Since 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_MISSING_BACK_REFERENCE</td>
<td class="value">157</td>
<td class="doc">"\\g" is not followed by a braced,
     angle-bracketed, or quoted name or number, or by a plain number. Since: 2.16</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_INVALID_RELATIVE_REFERENCE</td>
<td class="value">158</td>
<td class="doc">relative reference must not be zero. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_FORBIDDEN</td>
<td class="value">159</td>
<td class="doc">the backtracing
    control verb used does not allow an argument. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_UNKNOWN_BACKTRACKING_CONTROL_VERB</td>
<td class="value">160</td>
<td class="doc">unknown backtracing
    control verb. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_NUMBER_TOO_BIG</td>
<td class="value">161</td>
<td class="doc">number is too big in escape sequence. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_MISSING_SUBPATTERN_NAME</td>
<td class="value">162</td>
<td class="doc">Missing subpattern name. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_MISSING_DIGIT</td>
<td class="value">163</td>
<td class="doc">Missing digit. Since 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_INVALID_DATA_CHARACTER</td>
<td class="value">164</td>
<td class="doc">In JavaScript compatibility mode,
    "[" is an invalid data character. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_EXTRA_SUBPATTERN_NAME</td>
<td class="value">165</td>
<td class="doc">different names for subpatterns of the
    same number are not allowed. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_REQUIRED</td>
<td class="value">166</td>
<td class="doc">the backtracing control
    verb requires an argument. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_INVALID_CONTROL_CHAR</td>
<td class="value">168</td>
<td class="doc">"\\c" must be followed by an ASCII
    character. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_MISSING_NAME</td>
<td class="value">169</td>
<td class="doc">"\\k" is not followed by a braced, angle-bracketed, or
    quoted name. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_NOT_SUPPORTED_IN_CLASS</td>
<td class="value">171</td>
<td class="doc">"\\N" is not supported in a class. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_TOO_MANY_FORWARD_REFERENCES</td>
<td class="value">172</td>
<td class="doc">too many forward references. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_NAME_TOO_LONG</td>
<td class="value">175</td>
<td class="doc">the name is too long in "(*MARK)", "(*PRUNE)",
    "(*SKIP)", or "(*THEN)". Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_ERROR_CHARACTER_VALUE_TOO_LARGE</td>
<td class="value">176</td>
<td class="doc">the character value in the \\u sequence is
    too large. Since: 2.34</td>
</tr>
</table>
</div>
<p class="api-heading">SeekType</p>
<p class="api-doc">An enumeration specifying the base position for a
g_io_channel_seek_position() operation.</p>
<div class="api-notes">
  <p class="api-ctype">GSeekType</p>
<table>
<tr>
<td class="name">SEEK_CUR</td>
<td class="value">0</td>
<td class="doc">the current position in the file.</td>
</tr>
<tr>
<td class="name">SEEK_SET</td>
<td class="value">1</td>
<td class="doc">the start of the file.</td>
</tr>
<tr>
<td class="name">SEEK_END</td>
<td class="value">2</td>
<td class="doc">the end of the file.</td>
</tr>
</table>
</div>
<p class="api-heading">ShellError</p>
<p class="api-doc">Error codes returned by shell functions.</p>
<div class="api-notes">
  <p class="api-ctype">GShellError</p>
<table>
<tr>
<td class="name">SHELL_ERROR_BAD_QUOTING</td>
<td class="value">0</td>
<td class="doc">Mismatched or otherwise mangled quoting.</td>
</tr>
<tr>
<td class="name">SHELL_ERROR_EMPTY_STRING</td>
<td class="value">1</td>
<td class="doc">String to be parsed was empty.</td>
</tr>
<tr>
<td class="name">SHELL_ERROR_FAILED</td>
<td class="value">2</td>
<td class="doc">Some other error.</td>
</tr>
</table>
</div>
<p class="api-heading">SliceConfig</p>
<div class="api-notes">
  <p class="api-ctype">GSliceConfig</p>
<table>
<tr>
<td class="name">SLICE_CONFIG_ALWAYS_MALLOC</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">SLICE_CONFIG_BYPASS_MAGAZINES</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">SLICE_CONFIG_WORKING_SET_MSECS</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">SLICE_CONFIG_COLOR_INCREMENT</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">SLICE_CONFIG_CHUNK_SIZES</td>
<td class="value">5</td>
</tr>
<tr>
<td class="name">SLICE_CONFIG_CONTENTION_COUNTER</td>
<td class="value">6</td>
</tr>
</table>
</div>
<p class="api-heading">SpawnError</p>
<p class="api-doc">Error codes returned by spawning processes.</p>
<div class="api-notes">
  <p class="api-ctype">GSpawnError</p>
<table>
<tr>
<td class="name">SPAWN_ERROR_FORK</td>
<td class="value">0</td>
<td class="doc">Fork failed due to lack of memory.</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_READ</td>
<td class="value">1</td>
<td class="doc">Read or select on pipes failed.</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_CHDIR</td>
<td class="value">2</td>
<td class="doc">Changing to working directory failed.</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_ACCES</td>
<td class="value">3</td>
<td class="doc">execv() returned `EACCES`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_PERM</td>
<td class="value">4</td>
<td class="doc">execv() returned `EPERM`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_TOO_BIG</td>
<td class="value">5</td>
<td class="doc">execv() returned `E2BIG`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_2BIG</td>
<td class="value">5</td>
<td class="doc">deprecated alias for %G_SPAWN_ERROR_TOO_BIG</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_NOEXEC</td>
<td class="value">6</td>
<td class="doc">execv() returned `ENOEXEC`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_NAMETOOLONG</td>
<td class="value">7</td>
<td class="doc">execv() returned `ENAMETOOLONG`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_NOENT</td>
<td class="value">8</td>
<td class="doc">execv() returned `ENOENT`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_NOMEM</td>
<td class="value">9</td>
<td class="doc">execv() returned `ENOMEM`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_NOTDIR</td>
<td class="value">10</td>
<td class="doc">execv() returned `ENOTDIR`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_LOOP</td>
<td class="value">11</td>
<td class="doc">execv() returned `ELOOP`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_TXTBUSY</td>
<td class="value">12</td>
<td class="doc">execv() returned `ETXTBUSY`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_IO</td>
<td class="value">13</td>
<td class="doc">execv() returned `EIO`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_NFILE</td>
<td class="value">14</td>
<td class="doc">execv() returned `ENFILE`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_MFILE</td>
<td class="value">15</td>
<td class="doc">execv() returned `EMFILE`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_INVAL</td>
<td class="value">16</td>
<td class="doc">execv() returned `EINVAL`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_ISDIR</td>
<td class="value">17</td>
<td class="doc">execv() returned `EISDIR`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_LIBBAD</td>
<td class="value">18</td>
<td class="doc">execv() returned `ELIBBAD`</td>
</tr>
<tr>
<td class="name">SPAWN_ERROR_FAILED</td>
<td class="value">19</td>
<td class="doc">Some other fatal failure,
  `error->message` should explain.</td>
</tr>
</table>
</div>
<p class="api-heading">TestFileType</p>
<p class="api-doc">The type of file to return the filename for, when used with
g_test_build_filename().

These two options correspond rather directly to the 'dist' and
'built' terminology that automake uses and are explicitly used to
distinguish between the 'srcdir' and 'builddir' being separate.  All
files in your project should either be dist (in the
`EXTRA_DIST` or `dist_schema_DATA`
sense, in which case they will always be in the srcdir) or built (in
the `BUILT_SOURCES` sense, in which case they will
always be in the builddir).

Note: as a general rule of automake, files that are generated only as
part of the build-from-git process (but then are distributed with the
tarball) always go in srcdir (even if doing a srcdir != builddir
build from git) and are considered as distributed files.</p>
<div class="api-notes">
  <p class="api-ctype">GTestFileType</p>
  <p class="api-since">since 2.38</p>
<table>
<tr>
<td class="name">TEST_DIST</td>
<td class="value">0</td>
<td class="doc">a file that was included in the distribution tarball</td>
</tr>
<tr>
<td class="name">TEST_BUILT</td>
<td class="value">1</td>
<td class="doc">a file that was built on the compiling machine</td>
</tr>
</table>
</div>
<p class="api-heading">TestLogType</p>
<div class="api-notes">
  <p class="api-ctype">GTestLogType</p>
<table>
<tr>
<td class="name">TEST_LOG_NONE</td>
<td class="value">0</td>
</tr>
<tr>
<td class="name">TEST_LOG_ERROR</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">TEST_LOG_START_BINARY</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">TEST_LOG_LIST_CASE</td>
<td class="value">3</td>
</tr>
<tr>
<td class="name">TEST_LOG_SKIP_CASE</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">TEST_LOG_START_CASE</td>
<td class="value">5</td>
</tr>
<tr>
<td class="name">TEST_LOG_STOP_CASE</td>
<td class="value">6</td>
</tr>
<tr>
<td class="name">TEST_LOG_MIN_RESULT</td>
<td class="value">7</td>
</tr>
<tr>
<td class="name">TEST_LOG_MAX_RESULT</td>
<td class="value">8</td>
</tr>
<tr>
<td class="name">TEST_LOG_MESSAGE</td>
<td class="value">9</td>
</tr>
<tr>
<td class="name">TEST_LOG_START_SUITE</td>
<td class="value">10</td>
</tr>
<tr>
<td class="name">TEST_LOG_STOP_SUITE</td>
<td class="value">11</td>
</tr>
</table>
</div>
<p class="api-heading">ThreadError</p>
<p class="api-doc">Possible errors of thread related functions.</p>
<div class="api-notes">
  <p class="api-ctype">GThreadError</p>
<table>
<tr>
<td class="name">THREAD_ERROR_AGAIN</td>
<td class="value">0</td>
<td class="doc">a thread couldn't be created due to resource
                       shortage. Try again later.</td>
</tr>
</table>
</div>
<p class="api-heading">TimeType</p>
<p class="api-doc">Disambiguates a given time in two ways.

First, specifies if the given time is in universal or local time.

Second, if the time is in local time, specifies if it is local
standard time or local daylight time.  This is important for the case
where the same local time occurs twice (during daylight savings time
transitions, for example).</p>
<div class="api-notes">
  <p class="api-ctype">GTimeType</p>
<table>
<tr>
<td class="name">TIME_TYPE_STANDARD</td>
<td class="value">0</td>
<td class="doc">the time is in local standard time</td>
</tr>
<tr>
<td class="name">TIME_TYPE_DAYLIGHT</td>
<td class="value">1</td>
<td class="doc">the time is in local daylight time</td>
</tr>
<tr>
<td class="name">TIME_TYPE_UNIVERSAL</td>
<td class="value">2</td>
<td class="doc">the time is in UTC</td>
</tr>
</table>
</div>
<p class="api-heading">TokenType</p>
<p class="api-doc">The possible types of token returned from each
g_scanner_get_next_token() call.</p>
<div class="api-notes">
  <p class="api-ctype">GTokenType</p>
<table>
<tr>
<td class="name">TOKEN_EOF</td>
<td class="value">0</td>
<td class="doc">the end of the file</td>
</tr>
<tr>
<td class="name">TOKEN_LEFT_PAREN</td>
<td class="value">40</td>
<td class="doc">a '(' character</td>
</tr>
<tr>
<td class="name">TOKEN_RIGHT_PAREN</td>
<td class="value">41</td>
<td class="doc">a ')' character</td>
</tr>
<tr>
<td class="name">TOKEN_LEFT_CURLY</td>
<td class="value">123</td>
<td class="doc">a '{' character</td>
</tr>
<tr>
<td class="name">TOKEN_RIGHT_CURLY</td>
<td class="value">125</td>
<td class="doc">a '}' character</td>
</tr>
<tr>
<td class="name">TOKEN_LEFT_BRACE</td>
<td class="value">91</td>
<td class="doc">a '[' character</td>
</tr>
<tr>
<td class="name">TOKEN_RIGHT_BRACE</td>
<td class="value">93</td>
<td class="doc">a ']' character</td>
</tr>
<tr>
<td class="name">TOKEN_EQUAL_SIGN</td>
<td class="value">61</td>
<td class="doc">a '=' character</td>
</tr>
<tr>
<td class="name">TOKEN_COMMA</td>
<td class="value">44</td>
<td class="doc">a ',' character</td>
</tr>
<tr>
<td class="name">TOKEN_NONE</td>
<td class="value">256</td>
<td class="doc">not a token</td>
</tr>
<tr>
<td class="name">TOKEN_ERROR</td>
<td class="value">257</td>
<td class="doc">an error occurred</td>
</tr>
<tr>
<td class="name">TOKEN_CHAR</td>
<td class="value">258</td>
<td class="doc">a character</td>
</tr>
<tr>
<td class="name">TOKEN_BINARY</td>
<td class="value">259</td>
<td class="doc">a binary integer</td>
</tr>
<tr>
<td class="name">TOKEN_OCTAL</td>
<td class="value">260</td>
<td class="doc">an octal integer</td>
</tr>
<tr>
<td class="name">TOKEN_INT</td>
<td class="value">261</td>
<td class="doc">an integer</td>
</tr>
<tr>
<td class="name">TOKEN_HEX</td>
<td class="value">262</td>
<td class="doc">a hex integer</td>
</tr>
<tr>
<td class="name">TOKEN_FLOAT</td>
<td class="value">263</td>
<td class="doc">a floating point number</td>
</tr>
<tr>
<td class="name">TOKEN_STRING</td>
<td class="value">264</td>
<td class="doc">a string</td>
</tr>
<tr>
<td class="name">TOKEN_SYMBOL</td>
<td class="value">265</td>
<td class="doc">a symbol</td>
</tr>
<tr>
<td class="name">TOKEN_IDENTIFIER</td>
<td class="value">266</td>
<td class="doc">an identifier</td>
</tr>
<tr>
<td class="name">TOKEN_IDENTIFIER_NULL</td>
<td class="value">267</td>
<td class="doc">a null identifier</td>
</tr>
<tr>
<td class="name">TOKEN_COMMENT_SINGLE</td>
<td class="value">268</td>
<td class="doc">one line comment</td>
</tr>
<tr>
<td class="name">TOKEN_COMMENT_MULTI</td>
<td class="value">269</td>
<td class="doc">multi line comment</td>
</tr>
</table>
</div>
<p class="api-heading">TraverseType</p>
<p class="api-doc">Specifies the type of traveral performed by g_tree_traverse(),
g_node_traverse() and g_node_find(). The different orders are
illustrated here:
- In order: A, B, C, D, E, F, G, H, I
  ![](Sorted_binary_tree_inorder.svg)
- Pre order: F, B, A, D, C, E, G, I, H
  ![](Sorted_binary_tree_preorder.svg)
- Post order: A, C, E, D, B, H, I, G, F
  ![](Sorted_binary_tree_postorder.svg)
- Level order: F, B, G, A, D, I, C, E, H
  ![](Sorted_binary_tree_breadth-first_traversal.svg)</p>
<div class="api-notes">
  <p class="api-ctype">GTraverseType</p>
<table>
<tr>
<td class="name">IN_ORDER</td>
<td class="value">0</td>
<td class="doc">vists a node's left child first, then the node itself,
             then its right child. This is the one to use if you
             want the output sorted according to the compare
             function.</td>
</tr>
<tr>
<td class="name">PRE_ORDER</td>
<td class="value">1</td>
<td class="doc">visits a node, then its children.</td>
</tr>
<tr>
<td class="name">POST_ORDER</td>
<td class="value">2</td>
<td class="doc">visits the node's children, then the node itself.</td>
</tr>
<tr>
<td class="name">LEVEL_ORDER</td>
<td class="value">3</td>
<td class="doc">is not implemented for
             [balanced binary trees][glib-Balanced-Binary-Trees].
             For [n-ary trees][glib-N-ary-Trees], it
             vists the root node first, then its children, then
             its grandchildren, and so on. Note that this is less
             efficient than the other orders.</td>
</tr>
</table>
</div>
<p class="api-heading">UnicodeBreakType</p>
<p class="api-doc">These are the possible line break classifications.

Since new unicode versions may add new types here, applications should be ready
to handle unknown values. They may be regarded as %G_UNICODE_BREAK_UNKNOWN.

See [Unicode Line Breaking Algorithm](http://www.unicode.org/unicode/reports/tr14/).</p>
<div class="api-notes">
  <p class="api-ctype">GUnicodeBreakType</p>
<table>
<tr>
<td class="name">UNICODE_BREAK_MANDATORY</td>
<td class="value">0</td>
<td class="doc">Mandatory Break (BK)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_CARRIAGE_RETURN</td>
<td class="value">1</td>
<td class="doc">Carriage Return (CR)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_LINE_FEED</td>
<td class="value">2</td>
<td class="doc">Line Feed (LF)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_COMBINING_MARK</td>
<td class="value">3</td>
<td class="doc">Attached Characters and Combining Marks (CM)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_SURROGATE</td>
<td class="value">4</td>
<td class="doc">Surrogates (SG)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_ZERO_WIDTH_SPACE</td>
<td class="value">5</td>
<td class="doc">Zero Width Space (ZW)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_INSEPARABLE</td>
<td class="value">6</td>
<td class="doc">Inseparable (IN)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_NON_BREAKING_GLUE</td>
<td class="value">7</td>
<td class="doc">Non-breaking ("Glue") (GL)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_CONTINGENT</td>
<td class="value">8</td>
<td class="doc">Contingent Break Opportunity (CB)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_SPACE</td>
<td class="value">9</td>
<td class="doc">Space (SP)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_AFTER</td>
<td class="value">10</td>
<td class="doc">Break Opportunity After (BA)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_BEFORE</td>
<td class="value">11</td>
<td class="doc">Break Opportunity Before (BB)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_BEFORE_AND_AFTER</td>
<td class="value">12</td>
<td class="doc">Break Opportunity Before and After (B2)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_HYPHEN</td>
<td class="value">13</td>
<td class="doc">Hyphen (HY)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_NON_STARTER</td>
<td class="value">14</td>
<td class="doc">Nonstarter (NS)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_OPEN_PUNCTUATION</td>
<td class="value">15</td>
<td class="doc">Opening Punctuation (OP)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_CLOSE_PUNCTUATION</td>
<td class="value">16</td>
<td class="doc">Closing Punctuation (CL)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_QUOTATION</td>
<td class="value">17</td>
<td class="doc">Ambiguous Quotation (QU)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_EXCLAMATION</td>
<td class="value">18</td>
<td class="doc">Exclamation/Interrogation (EX)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_IDEOGRAPHIC</td>
<td class="value">19</td>
<td class="doc">Ideographic (ID)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_NUMERIC</td>
<td class="value">20</td>
<td class="doc">Numeric (NU)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_INFIX_SEPARATOR</td>
<td class="value">21</td>
<td class="doc">Infix Separator (Numeric) (IS)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_SYMBOL</td>
<td class="value">22</td>
<td class="doc">Symbols Allowing Break After (SY)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_ALPHABETIC</td>
<td class="value">23</td>
<td class="doc">Ordinary Alphabetic and Symbol Characters (AL)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_PREFIX</td>
<td class="value">24</td>
<td class="doc">Prefix (Numeric) (PR)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_POSTFIX</td>
<td class="value">25</td>
<td class="doc">Postfix (Numeric) (PO)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_COMPLEX_CONTEXT</td>
<td class="value">26</td>
<td class="doc">Complex Content Dependent (South East Asian) (SA)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_AMBIGUOUS</td>
<td class="value">27</td>
<td class="doc">Ambiguous (Alphabetic or Ideographic) (AI)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_UNKNOWN</td>
<td class="value">28</td>
<td class="doc">Unknown (XX)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_NEXT_LINE</td>
<td class="value">29</td>
<td class="doc">Next Line (NL)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_WORD_JOINER</td>
<td class="value">30</td>
<td class="doc">Word Joiner (WJ)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_HANGUL_L_JAMO</td>
<td class="value">31</td>
<td class="doc">Hangul L Jamo (JL)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_HANGUL_V_JAMO</td>
<td class="value">32</td>
<td class="doc">Hangul V Jamo (JV)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_HANGUL_T_JAMO</td>
<td class="value">33</td>
<td class="doc">Hangul T Jamo (JT)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_HANGUL_LV_SYLLABLE</td>
<td class="value">34</td>
<td class="doc">Hangul LV Syllable (H2)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_HANGUL_LVT_SYLLABLE</td>
<td class="value">35</td>
<td class="doc">Hangul LVT Syllable (H3)</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_CLOSE_PARANTHESIS</td>
<td class="value">36</td>
<td class="doc">Closing Parenthesis (CP). Since 2.28</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_CONDITIONAL_JAPANESE_STARTER</td>
<td class="value">37</td>
<td class="doc">Conditional Japanese Starter (CJ). Since: 2.32</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_HEBREW_LETTER</td>
<td class="value">38</td>
<td class="doc">Hebrew Letter (HL). Since: 2.32</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_REGIONAL_INDICATOR</td>
<td class="value">39</td>
<td class="doc">Regional Indicator (RI). Since: 2.36</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_EMOJI_BASE</td>
<td class="value">40</td>
<td class="doc">Emoji Base (EB). Since: 2.50</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_EMOJI_MODIFIER</td>
<td class="value">41</td>
<td class="doc">Emoji Modifier (EM). Since: 2.50</td>
</tr>
<tr>
<td class="name">UNICODE_BREAK_ZERO_WIDTH_JOINER</td>
<td class="value">42</td>
<td class="doc">Zero Width Joiner (ZWJ). Since: 2.50</td>
</tr>
</table>
</div>
<p class="api-heading">UnicodeScript</p>
<p class="api-doc">The #GUnicodeScript enumeration identifies different writing
systems. The values correspond to the names as defined in the
Unicode standard. The enumeration has been added in GLib 2.14,
and is interchangeable with #PangoScript.

Note that new types may be added in the future. Applications
should be ready to handle unknown values.
See [Unicode Standard Annex #24: Script names](http://www.unicode.org/reports/tr24/).</p>
<div class="api-notes">
  <p class="api-ctype">GUnicodeScript</p>
<table>
<tr>
<td class="name">UNICODE_SCRIPT_INVALID_CODE</td>
<td class="value">-1</td>
<td class="doc">a value never returned from g_unichar_get_script()</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_COMMON</td>
<td class="value">0</td>
<td class="doc">a character used by multiple different scripts</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_INHERITED</td>
<td class="value">1</td>
<td class="doc">a mark glyph that takes its script from the
                              base glyph to which it is attached</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_ARABIC</td>
<td class="value">2</td>
<td class="doc">Arabic</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_ARMENIAN</td>
<td class="value">3</td>
<td class="doc">Armenian</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BENGALI</td>
<td class="value">4</td>
<td class="doc">Bengali</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BOPOMOFO</td>
<td class="value">5</td>
<td class="doc">Bopomofo</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_CHEROKEE</td>
<td class="value">6</td>
<td class="doc">Cherokee</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_COPTIC</td>
<td class="value">7</td>
<td class="doc">Coptic</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_CYRILLIC</td>
<td class="value">8</td>
<td class="doc">Cyrillic</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_DESERET</td>
<td class="value">9</td>
<td class="doc">Deseret</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_DEVANAGARI</td>
<td class="value">10</td>
<td class="doc">Devanagari</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_ETHIOPIC</td>
<td class="value">11</td>
<td class="doc">Ethiopic</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_GEORGIAN</td>
<td class="value">12</td>
<td class="doc">Georgian</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_GOTHIC</td>
<td class="value">13</td>
<td class="doc">Gothic</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_GREEK</td>
<td class="value">14</td>
<td class="doc">Greek</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_GUJARATI</td>
<td class="value">15</td>
<td class="doc">Gujarati</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_GURMUKHI</td>
<td class="value">16</td>
<td class="doc">Gurmukhi</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_HAN</td>
<td class="value">17</td>
<td class="doc">Han</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_HANGUL</td>
<td class="value">18</td>
<td class="doc">Hangul</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_HEBREW</td>
<td class="value">19</td>
<td class="doc">Hebrew</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_HIRAGANA</td>
<td class="value">20</td>
<td class="doc">Hiragana</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_KANNADA</td>
<td class="value">21</td>
<td class="doc">Kannada</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_KATAKANA</td>
<td class="value">22</td>
<td class="doc">Katakana</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_KHMER</td>
<td class="value">23</td>
<td class="doc">Khmer</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_LAO</td>
<td class="value">24</td>
<td class="doc">Lao</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_LATIN</td>
<td class="value">25</td>
<td class="doc">Latin</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MALAYALAM</td>
<td class="value">26</td>
<td class="doc">Malayalam</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MONGOLIAN</td>
<td class="value">27</td>
<td class="doc">Mongolian</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MYANMAR</td>
<td class="value">28</td>
<td class="doc">Myanmar</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OGHAM</td>
<td class="value">29</td>
<td class="doc">Ogham</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OLD_ITALIC</td>
<td class="value">30</td>
<td class="doc">Old Italic</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_ORIYA</td>
<td class="value">31</td>
<td class="doc">Oriya</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_RUNIC</td>
<td class="value">32</td>
<td class="doc">Runic</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SINHALA</td>
<td class="value">33</td>
<td class="doc">Sinhala</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SYRIAC</td>
<td class="value">34</td>
<td class="doc">Syriac</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TAMIL</td>
<td class="value">35</td>
<td class="doc">Tamil</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TELUGU</td>
<td class="value">36</td>
<td class="doc">Telugu</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_THAANA</td>
<td class="value">37</td>
<td class="doc">Thaana</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_THAI</td>
<td class="value">38</td>
<td class="doc">Thai</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TIBETAN</td>
<td class="value">39</td>
<td class="doc">Tibetan</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_CANADIAN_ABORIGINAL</td>
<td class="value">40</td>
<td class="doc">Canadian Aboriginal</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_YI</td>
<td class="value">41</td>
<td class="doc">Yi</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TAGALOG</td>
<td class="value">42</td>
<td class="doc">Tagalog</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_HANUNOO</td>
<td class="value">43</td>
<td class="doc">Hanunoo</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BUHID</td>
<td class="value">44</td>
<td class="doc">Buhid</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TAGBANWA</td>
<td class="value">45</td>
<td class="doc">Tagbanwa</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BRAILLE</td>
<td class="value">46</td>
<td class="doc">Braille</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_CYPRIOT</td>
<td class="value">47</td>
<td class="doc">Cypriot</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_LIMBU</td>
<td class="value">48</td>
<td class="doc">Limbu</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OSMANYA</td>
<td class="value">49</td>
<td class="doc">Osmanya</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SHAVIAN</td>
<td class="value">50</td>
<td class="doc">Shavian</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_LINEAR_B</td>
<td class="value">51</td>
<td class="doc">Linear B</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TAI_LE</td>
<td class="value">52</td>
<td class="doc">Tai Le</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_UGARITIC</td>
<td class="value">53</td>
<td class="doc">Ugaritic</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_NEW_TAI_LUE</td>
<td class="value">54</td>
<td class="doc">New Tai Lue</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BUGINESE</td>
<td class="value">55</td>
<td class="doc">Buginese</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_GLAGOLITIC</td>
<td class="value">56</td>
<td class="doc">Glagolitic</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TIFINAGH</td>
<td class="value">57</td>
<td class="doc">Tifinagh</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SYLOTI_NAGRI</td>
<td class="value">58</td>
<td class="doc">Syloti Nagri</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OLD_PERSIAN</td>
<td class="value">59</td>
<td class="doc">Old Persian</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_KHAROSHTHI</td>
<td class="value">60</td>
<td class="doc">Kharoshthi</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_UNKNOWN</td>
<td class="value">61</td>
<td class="doc">an unassigned code point</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BALINESE</td>
<td class="value">62</td>
<td class="doc">Balinese</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_CUNEIFORM</td>
<td class="value">63</td>
<td class="doc">Cuneiform</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_PHOENICIAN</td>
<td class="value">64</td>
<td class="doc">Phoenician</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_PHAGS_PA</td>
<td class="value">65</td>
<td class="doc">Phags-pa</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_NKO</td>
<td class="value">66</td>
<td class="doc">N'Ko</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_KAYAH_LI</td>
<td class="value">67</td>
<td class="doc">Kayah Li. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_LEPCHA</td>
<td class="value">68</td>
<td class="doc">Lepcha. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_REJANG</td>
<td class="value">69</td>
<td class="doc">Rejang. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SUNDANESE</td>
<td class="value">70</td>
<td class="doc">Sundanese. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SAURASHTRA</td>
<td class="value">71</td>
<td class="doc">Saurashtra. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_CHAM</td>
<td class="value">72</td>
<td class="doc">Cham. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OL_CHIKI</td>
<td class="value">73</td>
<td class="doc">Ol Chiki. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_VAI</td>
<td class="value">74</td>
<td class="doc">Vai. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_CARIAN</td>
<td class="value">75</td>
<td class="doc">Carian. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_LYCIAN</td>
<td class="value">76</td>
<td class="doc">Lycian. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_LYDIAN</td>
<td class="value">77</td>
<td class="doc">Lydian. Since 2.16.3</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_AVESTAN</td>
<td class="value">78</td>
<td class="doc">Avestan. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BAMUM</td>
<td class="value">79</td>
<td class="doc">Bamum. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_EGYPTIAN_HIEROGLYPHS</td>
<td class="value">80</td>
<td class="doc">Egyptian Hieroglpyhs. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_IMPERIAL_ARAMAIC</td>
<td class="value">81</td>
<td class="doc">Imperial Aramaic. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_INSCRIPTIONAL_PAHLAVI</td>
<td class="value">82</td>
<td class="doc">Inscriptional Pahlavi. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_INSCRIPTIONAL_PARTHIAN</td>
<td class="value">83</td>
<td class="doc">Inscriptional Parthian. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_JAVANESE</td>
<td class="value">84</td>
<td class="doc">Javanese. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_KAITHI</td>
<td class="value">85</td>
<td class="doc">Kaithi. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_LISU</td>
<td class="value">86</td>
<td class="doc">Lisu. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MEETEI_MAYEK</td>
<td class="value">87</td>
<td class="doc">Meetei Mayek. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OLD_SOUTH_ARABIAN</td>
<td class="value">88</td>
<td class="doc">Old South Arabian. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OLD_TURKIC</td>
<td class="value">89</td>
<td class="doc">Old Turkic. Since 2.28</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SAMARITAN</td>
<td class="value">90</td>
<td class="doc">Samaritan. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TAI_THAM</td>
<td class="value">91</td>
<td class="doc">Tai Tham. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TAI_VIET</td>
<td class="value">92</td>
<td class="doc">Tai Viet. Since 2.26</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BATAK</td>
<td class="value">93</td>
<td class="doc">Batak. Since 2.28</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BRAHMI</td>
<td class="value">94</td>
<td class="doc">Brahmi. Since 2.28</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MANDAIC</td>
<td class="value">95</td>
<td class="doc">Mandaic. Since 2.28</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_CHAKMA</td>
<td class="value">96</td>
<td class="doc">Chakma. Since: 2.32</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MEROITIC_CURSIVE</td>
<td class="value">97</td>
<td class="doc">Meroitic Cursive. Since: 2.32</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MEROITIC_HIEROGLYPHS</td>
<td class="value">98</td>
<td class="doc">Meroitic Hieroglyphs. Since: 2.32</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MIAO</td>
<td class="value">99</td>
<td class="doc">Miao. Since: 2.32</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SHARADA</td>
<td class="value">100</td>
<td class="doc">Sharada. Since: 2.32</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SORA_SOMPENG</td>
<td class="value">101</td>
<td class="doc">Sora Sompeng. Since: 2.32</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TAKRI</td>
<td class="value">102</td>
<td class="doc">Takri. Since: 2.32</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BASSA_VAH</td>
<td class="value">103</td>
<td class="doc">Bassa. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_CAUCASIAN_ALBANIAN</td>
<td class="value">104</td>
<td class="doc">Caucasian Albanian. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_DUPLOYAN</td>
<td class="value">105</td>
<td class="doc">Duployan. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_ELBASAN</td>
<td class="value">106</td>
<td class="doc">Elbasan. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_GRANTHA</td>
<td class="value">107</td>
<td class="doc">Grantha. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_KHOJKI</td>
<td class="value">108</td>
<td class="doc">Kjohki. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_KHUDAWADI</td>
<td class="value">109</td>
<td class="doc">Khudawadi, Sindhi. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_LINEAR_A</td>
<td class="value">110</td>
<td class="doc">Linear A. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MAHAJANI</td>
<td class="value">111</td>
<td class="doc">Mahajani. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MANICHAEAN</td>
<td class="value">112</td>
<td class="doc">Manichaean. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MENDE_KIKAKUI</td>
<td class="value">113</td>
<td class="doc">Mende Kikakui. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MODI</td>
<td class="value">114</td>
<td class="doc">Modi. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MRO</td>
<td class="value">115</td>
<td class="doc">Mro. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_NABATAEAN</td>
<td class="value">116</td>
<td class="doc">Nabataean. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OLD_NORTH_ARABIAN</td>
<td class="value">117</td>
<td class="doc">Old North Arabian. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OLD_PERMIC</td>
<td class="value">118</td>
<td class="doc">Old Permic. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_PAHAWH_HMONG</td>
<td class="value">119</td>
<td class="doc">Pahawh Hmong. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_PALMYRENE</td>
<td class="value">120</td>
<td class="doc">Palmyrene. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_PAU_CIN_HAU</td>
<td class="value">121</td>
<td class="doc">Pau Cin Hau. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_PSALTER_PAHLAVI</td>
<td class="value">122</td>
<td class="doc">Psalter Pahlavi. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SIDDHAM</td>
<td class="value">123</td>
<td class="doc">Siddham. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TIRHUTA</td>
<td class="value">124</td>
<td class="doc">Tirhuta. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_WARANG_CITI</td>
<td class="value">125</td>
<td class="doc">Warang Citi. Since: 2.42</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_AHOM</td>
<td class="value">126</td>
<td class="doc">Ahom. Since: 2.48</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_ANATOLIAN_HIEROGLYPHS</td>
<td class="value">127</td>
<td class="doc">Anatolian Hieroglyphs. Since: 2.48</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_HATRAN</td>
<td class="value">128</td>
<td class="doc">Hatran. Since: 2.48</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MULTANI</td>
<td class="value">129</td>
<td class="doc">Multani. Since: 2.48</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OLD_HUNGARIAN</td>
<td class="value">130</td>
<td class="doc">Old Hungarian. Since: 2.48</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SIGNWRITING</td>
<td class="value">131</td>
<td class="doc">Signwriting. Since: 2.48</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_ADLAM</td>
<td class="value">132</td>
<td class="doc">Adlam. Since: 2.50</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_BHAIKSUKI</td>
<td class="value">133</td>
<td class="doc">Bhaiksuki. Since: 2.50</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MARCHEN</td>
<td class="value">134</td>
<td class="doc">Marchen. Since: 2.50</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_NEWA</td>
<td class="value">135</td>
<td class="doc">Newa. Since: 2.50</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_OSAGE</td>
<td class="value">136</td>
<td class="doc">Osage. Since: 2.50</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_TANGUT</td>
<td class="value">137</td>
<td class="doc">Tangut. Since: 2.50</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_MASARAM_GONDI</td>
<td class="value">138</td>
<td class="doc">Masaram Gondi. Since: 2.54</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_NUSHU</td>
<td class="value">139</td>
<td class="doc">Nushu. Since: 2.54</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_SOYOMBO</td>
<td class="value">140</td>
<td class="doc">Soyombo. Since: 2.54</td>
</tr>
<tr>
<td class="name">UNICODE_SCRIPT_ZANABAZAR_SQUARE</td>
<td class="value">141</td>
<td class="doc">Zanabazar Square. Since: 2.54</td>
</tr>
</table>
</div>
<p class="api-heading">UnicodeType</p>
<p class="api-doc">These are the possible character classifications from the
Unicode specification.
See [Unicode Character Database](http://www.unicode.org/reports/tr44/#General_Category_Values).</p>
<div class="api-notes">
  <p class="api-ctype">GUnicodeType</p>
<table>
<tr>
<td class="name">UNICODE_CONTROL</td>
<td class="value">0</td>
<td class="doc">General category "Other, Control" (Cc)</td>
</tr>
<tr>
<td class="name">UNICODE_FORMAT</td>
<td class="value">1</td>
<td class="doc">General category "Other, Format" (Cf)</td>
</tr>
<tr>
<td class="name">UNICODE_UNASSIGNED</td>
<td class="value">2</td>
<td class="doc">General category "Other, Not Assigned" (Cn)</td>
</tr>
<tr>
<td class="name">UNICODE_PRIVATE_USE</td>
<td class="value">3</td>
<td class="doc">General category "Other, Private Use" (Co)</td>
</tr>
<tr>
<td class="name">UNICODE_SURROGATE</td>
<td class="value">4</td>
<td class="doc">General category "Other, Surrogate" (Cs)</td>
</tr>
<tr>
<td class="name">UNICODE_LOWERCASE_LETTER</td>
<td class="value">5</td>
<td class="doc">General category "Letter, Lowercase" (Ll)</td>
</tr>
<tr>
<td class="name">UNICODE_MODIFIER_LETTER</td>
<td class="value">6</td>
<td class="doc">General category "Letter, Modifier" (Lm)</td>
</tr>
<tr>
<td class="name">UNICODE_OTHER_LETTER</td>
<td class="value">7</td>
<td class="doc">General category "Letter, Other" (Lo)</td>
</tr>
<tr>
<td class="name">UNICODE_TITLECASE_LETTER</td>
<td class="value">8</td>
<td class="doc">General category "Letter, Titlecase" (Lt)</td>
</tr>
<tr>
<td class="name">UNICODE_UPPERCASE_LETTER</td>
<td class="value">9</td>
<td class="doc">General category "Letter, Uppercase" (Lu)</td>
</tr>
<tr>
<td class="name">UNICODE_SPACING_MARK</td>
<td class="value">10</td>
<td class="doc">General category "Mark, Spacing" (Mc)</td>
</tr>
<tr>
<td class="name">UNICODE_ENCLOSING_MARK</td>
<td class="value">11</td>
<td class="doc">General category "Mark, Enclosing" (Me)</td>
</tr>
<tr>
<td class="name">UNICODE_NON_SPACING_MARK</td>
<td class="value">12</td>
<td class="doc">General category "Mark, Nonspacing" (Mn)</td>
</tr>
<tr>
<td class="name">UNICODE_DECIMAL_NUMBER</td>
<td class="value">13</td>
<td class="doc">General category "Number, Decimal Digit" (Nd)</td>
</tr>
<tr>
<td class="name">UNICODE_LETTER_NUMBER</td>
<td class="value">14</td>
<td class="doc">General category "Number, Letter" (Nl)</td>
</tr>
<tr>
<td class="name">UNICODE_OTHER_NUMBER</td>
<td class="value">15</td>
<td class="doc">General category "Number, Other" (No)</td>
</tr>
<tr>
<td class="name">UNICODE_CONNECT_PUNCTUATION</td>
<td class="value">16</td>
<td class="doc">General category "Punctuation, Connector" (Pc)</td>
</tr>
<tr>
<td class="name">UNICODE_DASH_PUNCTUATION</td>
<td class="value">17</td>
<td class="doc">General category "Punctuation, Dash" (Pd)</td>
</tr>
<tr>
<td class="name">UNICODE_CLOSE_PUNCTUATION</td>
<td class="value">18</td>
<td class="doc">General category "Punctuation, Close" (Pe)</td>
</tr>
<tr>
<td class="name">UNICODE_FINAL_PUNCTUATION</td>
<td class="value">19</td>
<td class="doc">General category "Punctuation, Final quote" (Pf)</td>
</tr>
<tr>
<td class="name">UNICODE_INITIAL_PUNCTUATION</td>
<td class="value">20</td>
<td class="doc">General category "Punctuation, Initial quote" (Pi)</td>
</tr>
<tr>
<td class="name">UNICODE_OTHER_PUNCTUATION</td>
<td class="value">21</td>
<td class="doc">General category "Punctuation, Other" (Po)</td>
</tr>
<tr>
<td class="name">UNICODE_OPEN_PUNCTUATION</td>
<td class="value">22</td>
<td class="doc">General category "Punctuation, Open" (Ps)</td>
</tr>
<tr>
<td class="name">UNICODE_CURRENCY_SYMBOL</td>
<td class="value">23</td>
<td class="doc">General category "Symbol, Currency" (Sc)</td>
</tr>
<tr>
<td class="name">UNICODE_MODIFIER_SYMBOL</td>
<td class="value">24</td>
<td class="doc">General category "Symbol, Modifier" (Sk)</td>
</tr>
<tr>
<td class="name">UNICODE_MATH_SYMBOL</td>
<td class="value">25</td>
<td class="doc">General category "Symbol, Math" (Sm)</td>
</tr>
<tr>
<td class="name">UNICODE_OTHER_SYMBOL</td>
<td class="value">26</td>
<td class="doc">General category "Symbol, Other" (So)</td>
</tr>
<tr>
<td class="name">UNICODE_LINE_SEPARATOR</td>
<td class="value">27</td>
<td class="doc">General category "Separator, Line" (Zl)</td>
</tr>
<tr>
<td class="name">UNICODE_PARAGRAPH_SEPARATOR</td>
<td class="value">28</td>
<td class="doc">General category "Separator, Paragraph" (Zp)</td>
</tr>
<tr>
<td class="name">UNICODE_SPACE_SEPARATOR</td>
<td class="value">29</td>
<td class="doc">General category "Separator, Space" (Zs)</td>
</tr>
</table>
</div>
<p class="api-heading">UserDirectory</p>
<p class="api-doc">These are logical ids for special directories which are defined
depending on the platform used. You should use g_get_user_special_dir()
to retrieve the full path associated to the logical id.

The #GUserDirectory enumeration can be extended at later date. Not
every platform has a directory for every logical id in this
enumeration.</p>
<div class="api-notes">
  <p class="api-ctype">GUserDirectory</p>
  <p class="api-since">since 2.14</p>
<table>
<tr>
<td class="name">USER_DIRECTORY_DESKTOP</td>
<td class="value">0</td>
<td class="doc">the user's Desktop directory</td>
</tr>
<tr>
<td class="name">USER_DIRECTORY_DOCUMENTS</td>
<td class="value">1</td>
<td class="doc">the user's Documents directory</td>
</tr>
<tr>
<td class="name">USER_DIRECTORY_DOWNLOAD</td>
<td class="value">2</td>
<td class="doc">the user's Downloads directory</td>
</tr>
<tr>
<td class="name">USER_DIRECTORY_MUSIC</td>
<td class="value">3</td>
<td class="doc">the user's Music directory</td>
</tr>
<tr>
<td class="name">USER_DIRECTORY_PICTURES</td>
<td class="value">4</td>
<td class="doc">the user's Pictures directory</td>
</tr>
<tr>
<td class="name">USER_DIRECTORY_PUBLIC_SHARE</td>
<td class="value">5</td>
<td class="doc">the user's shared directory</td>
</tr>
<tr>
<td class="name">USER_DIRECTORY_TEMPLATES</td>
<td class="value">6</td>
<td class="doc">the user's Templates directory</td>
</tr>
<tr>
<td class="name">USER_DIRECTORY_VIDEOS</td>
<td class="value">7</td>
<td class="doc">the user's Movies directory</td>
</tr>
<tr>
<td class="name">USER_N_DIRECTORIES</td>
<td class="value">8</td>
<td class="doc">the number of enum values</td>
</tr>
</table>
</div>
<p class="api-heading">VariantClass</p>
<p class="api-doc">The range of possible top-level types of #GVariant instances.</p>
<div class="api-notes">
  <p class="api-ctype">GVariantClass</p>
  <p class="api-since">since 2.24</p>
<table>
<tr>
<td class="name">VARIANT_CLASS_BOOLEAN</td>
<td class="value">98</td>
<td class="doc">The #GVariant is a boolean.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_BYTE</td>
<td class="value">121</td>
<td class="doc">The #GVariant is a byte.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_INT16</td>
<td class="value">110</td>
<td class="doc">The #GVariant is a signed 16 bit integer.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_UINT16</td>
<td class="value">113</td>
<td class="doc">The #GVariant is an unsigned 16 bit integer.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_INT32</td>
<td class="value">105</td>
<td class="doc">The #GVariant is a signed 32 bit integer.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_UINT32</td>
<td class="value">117</td>
<td class="doc">The #GVariant is an unsigned 32 bit integer.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_INT64</td>
<td class="value">120</td>
<td class="doc">The #GVariant is a signed 64 bit integer.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_UINT64</td>
<td class="value">116</td>
<td class="doc">The #GVariant is an unsigned 64 bit integer.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_HANDLE</td>
<td class="value">104</td>
<td class="doc">The #GVariant is a file handle index.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_DOUBLE</td>
<td class="value">100</td>
<td class="doc">The #GVariant is a double precision floating
                         point value.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_STRING</td>
<td class="value">115</td>
<td class="doc">The #GVariant is a normal string.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_OBJECT_PATH</td>
<td class="value">111</td>
<td class="doc">The #GVariant is a D-Bus object path
                              string.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_SIGNATURE</td>
<td class="value">103</td>
<td class="doc">The #GVariant is a D-Bus signature string.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_VARIANT</td>
<td class="value">118</td>
<td class="doc">The #GVariant is a variant.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_MAYBE</td>
<td class="value">109</td>
<td class="doc">The #GVariant is a maybe-typed value.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_ARRAY</td>
<td class="value">97</td>
<td class="doc">The #GVariant is an array.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_TUPLE</td>
<td class="value">40</td>
<td class="doc">The #GVariant is a tuple.</td>
</tr>
<tr>
<td class="name">VARIANT_CLASS_DICT_ENTRY</td>
<td class="value">123</td>
<td class="doc">The #GVariant is a dictionary entry.</td>
</tr>
</table>
</div>
<p class="api-heading">VariantParseError</p>
<p class="api-doc">Error codes returned by parsing text-format GVariants.</p>
<div class="api-notes">
  <p class="api-ctype">GVariantParseError</p>
<table>
<tr>
<td class="name">VARIANT_PARSE_ERROR_FAILED</td>
<td class="value">0</td>
<td class="doc">generic error (unused)</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_BASIC_TYPE_EXPECTED</td>
<td class="value">1</td>
<td class="doc">a non-basic #GVariantType was given where a basic type was expected</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_CANNOT_INFER_TYPE</td>
<td class="value">2</td>
<td class="doc">cannot infer the #GVariantType</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_DEFINITE_TYPE_EXPECTED</td>
<td class="value">3</td>
<td class="doc">an indefinite #GVariantType was given where a definite type was expected</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_INPUT_NOT_AT_END</td>
<td class="value">4</td>
<td class="doc">extra data after parsing finished</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_INVALID_CHARACTER</td>
<td class="value">5</td>
<td class="doc">invalid character in number or unicode escape</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_INVALID_FORMAT_STRING</td>
<td class="value">6</td>
<td class="doc">not a valid #GVariant format string</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_INVALID_OBJECT_PATH</td>
<td class="value">7</td>
<td class="doc">not a valid object path</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_INVALID_SIGNATURE</td>
<td class="value">8</td>
<td class="doc">not a valid type signature</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_INVALID_TYPE_STRING</td>
<td class="value">9</td>
<td class="doc">not a valid #GVariant type string</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_NO_COMMON_TYPE</td>
<td class="value">10</td>
<td class="doc">could not find a common type for array entries</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_NUMBER_OUT_OF_RANGE</td>
<td class="value">11</td>
<td class="doc">the numerical value is out of range of the given type</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_NUMBER_TOO_BIG</td>
<td class="value">12</td>
<td class="doc">the numerical value is out of range for any type</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_TYPE_ERROR</td>
<td class="value">13</td>
<td class="doc">cannot parse as variant of the specified type</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_UNEXPECTED_TOKEN</td>
<td class="value">14</td>
<td class="doc">an unexpected token was encountered</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_UNKNOWN_KEYWORD</td>
<td class="value">15</td>
<td class="doc">an unknown keyword was encountered</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_UNTERMINATED_STRING_CONSTANT</td>
<td class="value">16</td>
<td class="doc">unterminated string constant</td>
</tr>
<tr>
<td class="name">VARIANT_PARSE_ERROR_VALUE_EXPECTED</td>
<td class="value">17</td>
<td class="doc">no value given</td>
</tr>
</table>
</div>