+++
title = "enums"
+++
<p class="api-heading">AsciiType</p>
<div class="api-notes">
  <p class="api-ctype">GAsciiType</p>
<table>
<tr>
<td class="name">ASCII_ALNUM</td>
<td class="value">1</td>
</tr>
<tr>
<td class="name">ASCII_ALPHA</td>
<td class="value">2</td>
</tr>
<tr>
<td class="name">ASCII_CNTRL</td>
<td class="value">4</td>
</tr>
<tr>
<td class="name">ASCII_DIGIT</td>
<td class="value">8</td>
</tr>
<tr>
<td class="name">ASCII_GRAPH</td>
<td class="value">16</td>
</tr>
<tr>
<td class="name">ASCII_LOWER</td>
<td class="value">32</td>
</tr>
<tr>
<td class="name">ASCII_PRINT</td>
<td class="value">64</td>
</tr>
<tr>
<td class="name">ASCII_PUNCT</td>
<td class="value">128</td>
</tr>
<tr>
<td class="name">ASCII_SPACE</td>
<td class="value">256</td>
</tr>
<tr>
<td class="name">ASCII_UPPER</td>
<td class="value">512</td>
</tr>
<tr>
<td class="name">ASCII_XDIGIT</td>
<td class="value">1024</td>
</tr>
</table>
</div>
<p class="api-heading">FileTest</p>
<p class="api-doc">A test to perform on a file using g_file_test().</p>
<div class="api-notes">
  <p class="api-ctype">GFileTest</p>
<table>
<tr>
<td class="name">FILE_TEST_IS_REGULAR</td>
<td class="value">1</td>
<td class="doc">%TRUE if the file is a regular file
    (not a directory). Note that this test will also return %TRUE
    if the tested file is a symlink to a regular file.</td>
</tr>
<tr>
<td class="name">FILE_TEST_IS_SYMLINK</td>
<td class="value">2</td>
<td class="doc">%TRUE if the file is a symlink.</td>
</tr>
<tr>
<td class="name">FILE_TEST_IS_DIR</td>
<td class="value">4</td>
<td class="doc">%TRUE if the file is a directory.</td>
</tr>
<tr>
<td class="name">FILE_TEST_IS_EXECUTABLE</td>
<td class="value">8</td>
<td class="doc">%TRUE if the file is executable.</td>
</tr>
<tr>
<td class="name">FILE_TEST_EXISTS</td>
<td class="value">16</td>
<td class="doc">%TRUE if the file exists. It may or may not
    be a regular file.</td>
</tr>
</table>
</div>
<p class="api-heading">FormatSizeFlags</p>
<p class="api-doc">Flags to modify the format of the string returned by g_format_size_full().</p>
<div class="api-notes">
  <p class="api-ctype">GFormatSizeFlags</p>
<table>
<tr>
<td class="name">FORMAT_SIZE_DEFAULT</td>
<td class="value">0</td>
<td class="doc">behave the same as g_format_size()</td>
</tr>
<tr>
<td class="name">FORMAT_SIZE_LONG_FORMAT</td>
<td class="value">1</td>
<td class="doc">include the exact number of bytes as part
    of the returned string.  For example, "45.6 kB (45,612 bytes)".</td>
</tr>
<tr>
<td class="name">FORMAT_SIZE_IEC_UNITS</td>
<td class="value">2</td>
<td class="doc">use IEC (base 1024) units with "KiB"-style
    suffixes. IEC units should only be used for reporting things with
    a strong "power of 2" basis, like RAM sizes or RAID stripe sizes.
    Network and storage sizes should be reported in the normal SI units.</td>
</tr>
<tr>
<td class="name">FORMAT_SIZE_BITS</td>
<td class="value">4</td>
<td class="doc">set the size as a quantity in bits, rather than
    bytes, and return units in bits. For example, ‘Mb’ rather than ‘MB’.</td>
</tr>
</table>
</div>
<p class="api-heading">HookFlagMask</p>
<p class="api-doc">Flags used internally in the #GHook implementation.</p>
<div class="api-notes">
  <p class="api-ctype">GHookFlagMask</p>
<table>
<tr>
<td class="name">HOOK_FLAG_ACTIVE</td>
<td class="value">1</td>
<td class="doc">set if the hook has not been destroyed</td>
</tr>
<tr>
<td class="name">HOOK_FLAG_IN_CALL</td>
<td class="value">2</td>
<td class="doc">set if the hook is currently being run</td>
</tr>
<tr>
<td class="name">HOOK_FLAG_MASK</td>
<td class="value">15</td>
<td class="doc">A mask covering all bits reserved for
  hook flags; see %G_HOOK_FLAG_USER_SHIFT</td>
</tr>
</table>
</div>
<p class="api-heading">IOCondition</p>
<p class="api-doc">A bitwise combination representing a condition to watch for on an
event source.</p>
<div class="api-notes">
  <p class="api-ctype">GIOCondition</p>
<table>
<tr>
<td class="name">IO_IN</td>
<td class="value">1</td>
<td class="doc">There is data to read.</td>
</tr>
<tr>
<td class="name">IO_OUT</td>
<td class="value">4</td>
<td class="doc">Data can be written (without blocking).</td>
</tr>
<tr>
<td class="name">IO_PRI</td>
<td class="value">2</td>
<td class="doc">There is urgent data to read.</td>
</tr>
<tr>
<td class="name">IO_ERR</td>
<td class="value">8</td>
<td class="doc">Error condition.</td>
</tr>
<tr>
<td class="name">IO_HUP</td>
<td class="value">16</td>
<td class="doc">Hung up (the connection has been broken, usually for
           pipes and sockets).</td>
</tr>
<tr>
<td class="name">IO_NVAL</td>
<td class="value">32</td>
<td class="doc">Invalid request. The file descriptor is not open.</td>
</tr>
</table>
</div>
<p class="api-heading">IOFlags</p>
<p class="api-doc">Specifies properties of a #GIOChannel. Some of the flags can only be
read with g_io_channel_get_flags(), but not changed with
g_io_channel_set_flags().</p>
<div class="api-notes">
  <p class="api-ctype">GIOFlags</p>
<table>
<tr>
<td class="name">IO_FLAG_APPEND</td>
<td class="value">1</td>
<td class="doc">turns on append mode, corresponds to %O_APPEND
    (see the documentation of the UNIX open() syscall)</td>
</tr>
<tr>
<td class="name">IO_FLAG_NONBLOCK</td>
<td class="value">2</td>
<td class="doc">turns on nonblocking mode, corresponds to
    %O_NONBLOCK/%O_NDELAY (see the documentation of the UNIX open()
    syscall)</td>
</tr>
<tr>
<td class="name">IO_FLAG_IS_READABLE</td>
<td class="value">4</td>
<td class="doc">indicates that the io channel is readable.
    This flag cannot be changed.</td>
</tr>
<tr>
<td class="name">IO_FLAG_IS_WRITABLE</td>
<td class="value">8</td>
<td class="doc">indicates that the io channel is writable.
    This flag cannot be changed.</td>
</tr>
<tr>
<td class="name">IO_FLAG_IS_WRITEABLE</td>
<td class="value">8</td>
<td class="doc">a misspelled version of @G_IO_FLAG_IS_WRITABLE
    that existed before the spelling was fixed in GLib 2.30. It is kept
    here for compatibility reasons. Deprecated since 2.30</td>
</tr>
<tr>
<td class="name">IO_FLAG_IS_SEEKABLE</td>
<td class="value">16</td>
<td class="doc">indicates that the io channel is seekable,
    i.e. that g_io_channel_seek_position() can be used on it.
    This flag cannot be changed.</td>
</tr>
<tr>
<td class="name">IO_FLAG_MASK</td>
<td class="value">31</td>
<td class="doc">the mask that specifies all the valid flags.</td>
</tr>
<tr>
<td class="name">IO_FLAG_GET_MASK</td>
<td class="value">31</td>
<td class="doc">the mask of the flags that are returned from
    g_io_channel_get_flags()</td>
</tr>
<tr>
<td class="name">IO_FLAG_SET_MASK</td>
<td class="value">3</td>
<td class="doc">the mask of the flags that the user can modify
    with g_io_channel_set_flags()</td>
</tr>
</table>
</div>
<p class="api-heading">KeyFileFlags</p>
<p class="api-doc">Flags which influence the parsing.</p>
<div class="api-notes">
  <p class="api-ctype">GKeyFileFlags</p>
<table>
<tr>
<td class="name">KEY_FILE_NONE</td>
<td class="value">0</td>
<td class="doc">No flags, default behaviour</td>
</tr>
<tr>
<td class="name">KEY_FILE_KEEP_COMMENTS</td>
<td class="value">1</td>
<td class="doc">Use this flag if you plan to write the
    (possibly modified) contents of the key file back to a file;
    otherwise all comments will be lost when the key file is
    written back.</td>
</tr>
<tr>
<td class="name">KEY_FILE_KEEP_TRANSLATIONS</td>
<td class="value">2</td>
<td class="doc">Use this flag if you plan to write the
    (possibly modified) contents of the key file back to a file;
    otherwise only the translations for the current language will be
    written back.</td>
</tr>
</table>
</div>
<p class="api-heading">LogLevelFlags</p>
<p class="api-doc">Flags specifying the level of log messages.

It is possible to change how GLib treats messages of the various
levels using g_log_set_handler() and g_log_set_fatal_mask().</p>
<div class="api-notes">
  <p class="api-ctype">GLogLevelFlags</p>
<table>
<tr>
<td class="name">LOG_FLAG_RECURSION</td>
<td class="value">1</td>
<td class="doc">internal flag</td>
</tr>
<tr>
<td class="name">LOG_FLAG_FATAL</td>
<td class="value">2</td>
<td class="doc">internal flag</td>
</tr>
<tr>
<td class="name">LOG_LEVEL_ERROR</td>
<td class="value">4</td>
<td class="doc">log level for errors, see g_error().
    This level is also used for messages produced by g_assert().</td>
</tr>
<tr>
<td class="name">LOG_LEVEL_CRITICAL</td>
<td class="value">8</td>
<td class="doc">log level for critical warning messages, see
    g_critical().
    This level is also used for messages produced by g_return_if_fail()
    and g_return_val_if_fail().</td>
</tr>
<tr>
<td class="name">LOG_LEVEL_WARNING</td>
<td class="value">16</td>
<td class="doc">log level for warnings, see g_warning()</td>
</tr>
<tr>
<td class="name">LOG_LEVEL_MESSAGE</td>
<td class="value">32</td>
<td class="doc">log level for messages, see g_message()</td>
</tr>
<tr>
<td class="name">LOG_LEVEL_INFO</td>
<td class="value">64</td>
<td class="doc">log level for informational messages, see g_info()</td>
</tr>
<tr>
<td class="name">LOG_LEVEL_DEBUG</td>
<td class="value">128</td>
<td class="doc">log level for debug messages, see g_debug()</td>
</tr>
<tr>
<td class="name">LOG_LEVEL_MASK</td>
<td class="value">-4</td>
<td class="doc">a mask including all log levels</td>
</tr>
</table>
</div>
<p class="api-heading">MarkupCollectType</p>
<p class="api-doc">A mixed enumerated type and flags field. You must specify one type
(string, strdup, boolean, tristate).  Additionally, you may  optionally
bitwise OR the type with the flag %G_MARKUP_COLLECT_OPTIONAL.

It is likely that this enum will be extended in the future to
support other types.</p>
<div class="api-notes">
  <p class="api-ctype">GMarkupCollectType</p>
<table>
<tr>
<td class="name">MARKUP_COLLECT_INVALID</td>
<td class="value">0</td>
<td class="doc">used to terminate the list of attributes
    to collect</td>
</tr>
<tr>
<td class="name">MARKUP_COLLECT_STRING</td>
<td class="value">1</td>
<td class="doc">collect the string pointer directly from
    the attribute_values[] array. Expects a parameter of type (const
    char **). If %G_MARKUP_COLLECT_OPTIONAL is specified and the
    attribute isn't present then the pointer will be set to %NULL</td>
</tr>
<tr>
<td class="name">MARKUP_COLLECT_STRDUP</td>
<td class="value">2</td>
<td class="doc">as with %G_MARKUP_COLLECT_STRING, but
    expects a parameter of type (char **) and g_strdup()s the
    returned pointer. The pointer must be freed with g_free()</td>
</tr>
<tr>
<td class="name">MARKUP_COLLECT_BOOLEAN</td>
<td class="value">3</td>
<td class="doc">expects a parameter of type (gboolean *)
    and parses the attribute value as a boolean. Sets %FALSE if the
    attribute isn't present. Valid boolean values consist of
    (case-insensitive) "false", "f", "no", "n", "0" and "true", "t",
    "yes", "y", "1"</td>
</tr>
<tr>
<td class="name">MARKUP_COLLECT_TRISTATE</td>
<td class="value">4</td>
<td class="doc">as with %G_MARKUP_COLLECT_BOOLEAN, but
    in the case of a missing attribute a value is set that compares
    equal to neither %FALSE nor %TRUE G_MARKUP_COLLECT_OPTIONAL is
    implied</td>
</tr>
<tr>
<td class="name">MARKUP_COLLECT_OPTIONAL</td>
<td class="value">65536</td>
<td class="doc">can be bitwise ORed with the other fields.
    If present, allows the attribute not to appear. A default value
    is set depending on what value type is used</td>
</tr>
</table>
</div>
<p class="api-heading">MarkupParseFlags</p>
<p class="api-doc">Flags that affect the behaviour of the parser.</p>
<div class="api-notes">
  <p class="api-ctype">GMarkupParseFlags</p>
<table>
<tr>
<td class="name">MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG</td>
<td class="value">1</td>
<td class="doc">flag you should not use</td>
</tr>
<tr>
<td class="name">MARKUP_TREAT_CDATA_AS_TEXT</td>
<td class="value">2</td>
<td class="doc">When this flag is set, CDATA marked
    sections are not passed literally to the @passthrough function of
    the parser. Instead, the content of the section (without the
    `<![CDATA[` and `]]>`) is
    passed to the @text function. This flag was added in GLib 2.12</td>
</tr>
<tr>
<td class="name">MARKUP_PREFIX_ERROR_POSITION</td>
<td class="value">4</td>
<td class="doc">Normally errors caught by GMarkup
    itself have line/column information prefixed to them to let the
    caller know the location of the error. When this flag is set the
    location information is also prefixed to errors generated by the
    #GMarkupParser implementation functions</td>
</tr>
<tr>
<td class="name">MARKUP_IGNORE_QUALIFIED</td>
<td class="value">8</td>
<td class="doc">Ignore (don't report) qualified
    attributes and tags, along with their contents.  A qualified
    attribute or tag is one that contains ':' in its name (ie: is in
    another namespace).  Since: 2.40.</td>
</tr>
</table>
</div>
<p class="api-heading">OptionFlags</p>
<p class="api-doc">Flags which modify individual options.</p>
<div class="api-notes">
  <p class="api-ctype">GOptionFlags</p>
<table>
<tr>
<td class="name">OPTION_FLAG_NONE</td>
<td class="value">0</td>
<td class="doc">No flags. Since: 2.42.</td>
</tr>
<tr>
<td class="name">OPTION_FLAG_HIDDEN</td>
<td class="value">1</td>
<td class="doc">The option doesn't appear in `--help` output.</td>
</tr>
<tr>
<td class="name">OPTION_FLAG_IN_MAIN</td>
<td class="value">2</td>
<td class="doc">The option appears in the main section of the
    `--help` output, even if it is defined in a group.</td>
</tr>
<tr>
<td class="name">OPTION_FLAG_REVERSE</td>
<td class="value">4</td>
<td class="doc">For options of the %G_OPTION_ARG_NONE kind, this
    flag indicates that the sense of the option is reversed.</td>
</tr>
<tr>
<td class="name">OPTION_FLAG_NO_ARG</td>
<td class="value">8</td>
<td class="doc">For options of the %G_OPTION_ARG_CALLBACK kind,
    this flag indicates that the callback does not take any argument
    (like a %G_OPTION_ARG_NONE option). Since 2.8</td>
</tr>
<tr>
<td class="name">OPTION_FLAG_FILENAME</td>
<td class="value">16</td>
<td class="doc">For options of the %G_OPTION_ARG_CALLBACK
    kind, this flag indicates that the argument should be passed to the
    callback in the GLib filename encoding rather than UTF-8. Since 2.8</td>
</tr>
<tr>
<td class="name">OPTION_FLAG_OPTIONAL_ARG</td>
<td class="value">32</td>
<td class="doc">For options of the %G_OPTION_ARG_CALLBACK
    kind, this flag indicates that the argument supply is optional.
    If no argument is given then data of %GOptionParseFunc will be
    set to NULL. Since 2.8</td>
</tr>
<tr>
<td class="name">OPTION_FLAG_NOALIAS</td>
<td class="value">64</td>
<td class="doc">This flag turns off the automatic conflict
    resolution which prefixes long option names with `groupname-` if
    there is a conflict. This option should only be used in situations
    where aliasing is necessary to model some legacy commandline interface.
    It is not safe to use this option, unless all option groups are under
    your direct control. Since 2.8.</td>
</tr>
</table>
</div>
<p class="api-heading">RegexCompileFlags</p>
<p class="api-doc">Flags specifying compile-time options.</p>
<div class="api-notes">
  <p class="api-ctype">GRegexCompileFlags</p>
  <p class="api-since">since 2.14</p>
<table>
<tr>
<td class="name">REGEX_CASELESS</td>
<td class="value">1</td>
<td class="doc">Letters in the pattern match both upper- and
    lowercase letters. This option can be changed within a pattern
    by a "(?i)" option setting.</td>
</tr>
<tr>
<td class="name">REGEX_MULTILINE</td>
<td class="value">2</td>
<td class="doc">By default, GRegex treats the strings as consisting
    of a single line of characters (even if it actually contains
    newlines). The "start of line" metacharacter ("^") matches only
    at the start of the string, while the "end of line" metacharacter
    ("$") matches only at the end of the string, or before a terminating
    newline (unless #G_REGEX_DOLLAR_ENDONLY is set). When
    #G_REGEX_MULTILINE is set, the "start of line" and "end of line"
    constructs match immediately following or immediately before any
    newline in the string, respectively, as well as at the very start
    and end. This can be changed within a pattern by a "(?m)" option
    setting.</td>
</tr>
<tr>
<td class="name">REGEX_DOTALL</td>
<td class="value">4</td>
<td class="doc">A dot metacharacter (".") in the pattern matches all
    characters, including newlines. Without it, newlines are excluded.
    This option can be changed within a pattern by a ("?s") option setting.</td>
</tr>
<tr>
<td class="name">REGEX_EXTENDED</td>
<td class="value">8</td>
<td class="doc">Whitespace data characters in the pattern are
    totally ignored except when escaped or inside a character class.
    Whitespace does not include the VT character (code 11). In addition,
    characters between an unescaped "#" outside a character class and
    the next newline character, inclusive, are also ignored. This can
    be changed within a pattern by a "(?x)" option setting.</td>
</tr>
<tr>
<td class="name">REGEX_ANCHORED</td>
<td class="value">16</td>
<td class="doc">The pattern is forced to be "anchored", that is,
    it is constrained to match only at the first matching point in the
    string that is being searched. This effect can also be achieved by
    appropriate constructs in the pattern itself such as the "^"
    metacharacter.</td>
</tr>
<tr>
<td class="name">REGEX_DOLLAR_ENDONLY</td>
<td class="value">32</td>
<td class="doc">A dollar metacharacter ("$") in the pattern
    matches only at the end of the string. Without this option, a
    dollar also matches immediately before the final character if
    it is a newline (but not before any other newlines). This option
    is ignored if #G_REGEX_MULTILINE is set.</td>
</tr>
<tr>
<td class="name">REGEX_UNGREEDY</td>
<td class="value">512</td>
<td class="doc">Inverts the "greediness" of the quantifiers so that
    they are not greedy by default, but become greedy if followed by "?".
    It can also be set by a "(?U)" option setting within the pattern.</td>
</tr>
<tr>
<td class="name">REGEX_RAW</td>
<td class="value">2048</td>
<td class="doc">Usually strings must be valid UTF-8 strings, using this
    flag they are considered as a raw sequence of bytes.</td>
</tr>
<tr>
<td class="name">REGEX_NO_AUTO_CAPTURE</td>
<td class="value">4096</td>
<td class="doc">Disables the use of numbered capturing
    parentheses in the pattern. Any opening parenthesis that is not
    followed by "?" behaves as if it were followed by "?:" but named
    parentheses can still be used for capturing (and they acquire numbers
    in the usual way).</td>
</tr>
<tr>
<td class="name">REGEX_OPTIMIZE</td>
<td class="value">8192</td>
<td class="doc">Optimize the regular expression. If the pattern will
    be used many times, then it may be worth the effort to optimize it
    to improve the speed of matches.</td>
</tr>
<tr>
<td class="name">REGEX_FIRSTLINE</td>
<td class="value">262144</td>
<td class="doc">Limits an unanchored pattern to match before (or at) the
    first newline. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_DUPNAMES</td>
<td class="value">524288</td>
<td class="doc">Names used to identify capturing subpatterns need not
    be unique. This can be helpful for certain types of pattern when it
    is known that only one instance of the named subpattern can ever be
    matched.</td>
</tr>
<tr>
<td class="name">REGEX_NEWLINE_CR</td>
<td class="value">1048576</td>
<td class="doc">Usually any newline character or character sequence is
    recognized. If this option is set, the only recognized newline character
    is '\r'.</td>
</tr>
<tr>
<td class="name">REGEX_NEWLINE_LF</td>
<td class="value">2097152</td>
<td class="doc">Usually any newline character or character sequence is
    recognized. If this option is set, the only recognized newline character
    is '\n'.</td>
</tr>
<tr>
<td class="name">REGEX_NEWLINE_CRLF</td>
<td class="value">3145728</td>
<td class="doc">Usually any newline character or character sequence is
    recognized. If this option is set, the only recognized newline character
    sequence is '\r\n'.</td>
</tr>
<tr>
<td class="name">REGEX_NEWLINE_ANYCRLF</td>
<td class="value">5242880</td>
<td class="doc">Usually any newline character or character sequence
    is recognized. If this option is set, the only recognized newline character
    sequences are '\r', '\n', and '\r\n'. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_BSR_ANYCRLF</td>
<td class="value">8388608</td>
<td class="doc">Usually any newline character or character sequence
    is recognised. If this option is set, then "\R" only recognizes the newline
   characters '\r', '\n' and '\r\n'. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_JAVASCRIPT_COMPAT</td>
<td class="value">33554432</td>
<td class="doc">Changes behaviour so that it is compatible with
    JavaScript rather than PCRE. Since: 2.34</td>
</tr>
</table>
</div>
<p class="api-heading">RegexMatchFlags</p>
<p class="api-doc">Flags specifying match-time options.</p>
<div class="api-notes">
  <p class="api-ctype">GRegexMatchFlags</p>
  <p class="api-since">since 2.14</p>
<table>
<tr>
<td class="name">REGEX_MATCH_ANCHORED</td>
<td class="value">16</td>
<td class="doc">The pattern is forced to be "anchored", that is,
    it is constrained to match only at the first matching point in the
    string that is being searched. This effect can also be achieved by
    appropriate constructs in the pattern itself such as the "^"
    metacharacter.</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_NOTBOL</td>
<td class="value">128</td>
<td class="doc">Specifies that first character of the string is
    not the beginning of a line, so the circumflex metacharacter should
    not match before it. Setting this without #G_REGEX_MULTILINE (at
    compile time) causes circumflex never to match. This option affects
    only the behaviour of the circumflex metacharacter, it does not
    affect "\A".</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_NOTEOL</td>
<td class="value">256</td>
<td class="doc">Specifies that the end of the subject string is
    not the end of a line, so the dollar metacharacter should not match
    it nor (except in multiline mode) a newline immediately before it.
    Setting this without #G_REGEX_MULTILINE (at compile time) causes
    dollar never to match. This option affects only the behaviour of
    the dollar metacharacter, it does not affect "\Z" or "\z".</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_NOTEMPTY</td>
<td class="value">1024</td>
<td class="doc">An empty string is not considered to be a valid
    match if this option is set. If there are alternatives in the pattern,
    they are tried. If all the alternatives match the empty string, the
    entire match fails. For example, if the pattern "a?b?" is applied to
    a string not beginning with "a" or "b", it matches the empty string
    at the start of the string. With this flag set, this match is not
    valid, so GRegex searches further into the string for occurrences
    of "a" or "b".</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_PARTIAL</td>
<td class="value">32768</td>
<td class="doc">Turns on the partial matching feature, for more
    documentation on partial matching see g_match_info_is_partial_match().</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_NEWLINE_CR</td>
<td class="value">1048576</td>
<td class="doc">Overrides the newline definition set when
    creating a new #GRegex, setting the '\r' character as line terminator.</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_NEWLINE_LF</td>
<td class="value">2097152</td>
<td class="doc">Overrides the newline definition set when
    creating a new #GRegex, setting the '\n' character as line terminator.</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_NEWLINE_CRLF</td>
<td class="value">3145728</td>
<td class="doc">Overrides the newline definition set when
    creating a new #GRegex, setting the '\r\n' characters sequence as line terminator.</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_NEWLINE_ANY</td>
<td class="value">4194304</td>
<td class="doc">Overrides the newline definition set when
    creating a new #GRegex, any Unicode newline sequence
    is recognised as a newline. These are '\r', '\n' and '\rn', and the
    single characters U+000B LINE TABULATION, U+000C FORM FEED (FF),
    U+0085 NEXT LINE (NEL), U+2028 LINE SEPARATOR and
    U+2029 PARAGRAPH SEPARATOR.</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_NEWLINE_ANYCRLF</td>
<td class="value">5242880</td>
<td class="doc">Overrides the newline definition set when
    creating a new #GRegex; any '\r', '\n', or '\r\n' character sequence
    is recognized as a newline. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_BSR_ANYCRLF</td>
<td class="value">8388608</td>
<td class="doc">Overrides the newline definition for "\R" set when
    creating a new #GRegex; only '\r', '\n', or '\r\n' character sequences
    are recognized as a newline by "\R". Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_BSR_ANY</td>
<td class="value">16777216</td>
<td class="doc">Overrides the newline definition for "\R" set when
    creating a new #GRegex; any Unicode newline character or character sequence
    are recognized as a newline by "\R". These are '\r', '\n' and '\rn', and the
    single characters U+000B LINE TABULATION, U+000C FORM FEED (FF),
    U+0085 NEXT LINE (NEL), U+2028 LINE SEPARATOR and
    U+2029 PARAGRAPH SEPARATOR. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_PARTIAL_SOFT</td>
<td class="value">32768</td>
<td class="doc">An alias for #G_REGEX_MATCH_PARTIAL. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_PARTIAL_HARD</td>
<td class="value">134217728</td>
<td class="doc">Turns on the partial matching feature. In contrast to
    to #G_REGEX_MATCH_PARTIAL_SOFT, this stops matching as soon as a partial match
    is found, without continuing to search for a possible complete match. See
    g_match_info_is_partial_match() for more information. Since: 2.34</td>
</tr>
<tr>
<td class="name">REGEX_MATCH_NOTEMPTY_ATSTART</td>
<td class="value">268435456</td>
<td class="doc">Like #G_REGEX_MATCH_NOTEMPTY, but only applied to
    the start of the matched string. For anchored
    patterns this can only happen for pattern containing "\K". Since: 2.34</td>
</tr>
</table>
</div>
<p class="api-heading">SpawnFlags</p>
<p class="api-doc">Flags passed to g_spawn_sync(), g_spawn_async() and g_spawn_async_with_pipes().</p>
<div class="api-notes">
  <p class="api-ctype">GSpawnFlags</p>
<table>
<tr>
<td class="name">SPAWN_DEFAULT</td>
<td class="value">0</td>
<td class="doc">no flags, default behaviour</td>
</tr>
<tr>
<td class="name">SPAWN_LEAVE_DESCRIPTORS_OPEN</td>
<td class="value">1</td>
<td class="doc">the parent's open file descriptors will
    be inherited by the child; otherwise all descriptors except stdin,
    stdout and stderr will be closed before calling exec() in the child.</td>
</tr>
<tr>
<td class="name">SPAWN_DO_NOT_REAP_CHILD</td>
<td class="value">2</td>
<td class="doc">the child will not be automatically reaped;
    you must use g_child_watch_add() yourself (or call waitpid() or handle
    `SIGCHLD` yourself), or the child will become a zombie.</td>
</tr>
<tr>
<td class="name">SPAWN_SEARCH_PATH</td>
<td class="value">4</td>
<td class="doc">`argv[0]` need not be an absolute path, it will be
    looked for in the user's `PATH`.</td>
</tr>
<tr>
<td class="name">SPAWN_STDOUT_TO_DEV_NULL</td>
<td class="value">8</td>
<td class="doc">the child's standard output will be discarded,
    instead of going to the same location as the parent's standard output.</td>
</tr>
<tr>
<td class="name">SPAWN_STDERR_TO_DEV_NULL</td>
<td class="value">16</td>
<td class="doc">the child's standard error will be discarded.</td>
</tr>
<tr>
<td class="name">SPAWN_CHILD_INHERITS_STDIN</td>
<td class="value">32</td>
<td class="doc">the child will inherit the parent's standard
    input (by default, the child's standard input is attached to `/dev/null`).</td>
</tr>
<tr>
<td class="name">SPAWN_FILE_AND_ARGV_ZERO</td>
<td class="value">64</td>
<td class="doc">the first element of `argv` is the file to
    execute, while the remaining elements are the actual argument vector
    to pass to the file. Normally g_spawn_async_with_pipes() uses `argv[0]`
    as the file to execute, and passes all of `argv` to the child.</td>
</tr>
<tr>
<td class="name">SPAWN_SEARCH_PATH_FROM_ENVP</td>
<td class="value">128</td>
<td class="doc">if `argv[0]` is not an abolute path,
    it will be looked for in the `PATH` from the passed child environment.
    Since: 2.34</td>
</tr>
<tr>
<td class="name">SPAWN_CLOEXEC_PIPES</td>
<td class="value">256</td>
<td class="doc">create all pipes with the `O_CLOEXEC` flag set.
    Since: 2.40</td>
</tr>
</table>
</div>
<p class="api-heading">TestSubprocessFlags</p>
<p class="api-doc">Flags to pass to g_test_trap_subprocess() to control input and output.

Note that in contrast with g_test_trap_fork(), the default is to
not show stdout and stderr.</p>
<div class="api-notes">
  <p class="api-ctype">GTestSubprocessFlags</p>
<table>
<tr>
<td class="name">TEST_SUBPROCESS_INHERIT_STDIN</td>
<td class="value">1</td>
<td class="doc">If this flag is given, the child
    process will inherit the parent's stdin. Otherwise, the child's
    stdin is redirected to `/dev/null`.</td>
</tr>
<tr>
<td class="name">TEST_SUBPROCESS_INHERIT_STDOUT</td>
<td class="value">2</td>
<td class="doc">If this flag is given, the child
    process will inherit the parent's stdout. Otherwise, the child's
    stdout will not be visible, but it will be captured to allow
    later tests with g_test_trap_assert_stdout().</td>
</tr>
<tr>
<td class="name">TEST_SUBPROCESS_INHERIT_STDERR</td>
<td class="value">4</td>
<td class="doc">If this flag is given, the child
    process will inherit the parent's stderr. Otherwise, the child's
    stderr will not be visible, but it will be captured to allow
    later tests with g_test_trap_assert_stderr().</td>
</tr>
</table>
</div>
<p class="api-heading">TestTrapFlags</p>
<p class="api-doc">Test traps are guards around forked tests.
These flags determine what traps to set.</p>
<div class="api-notes">
  <p class="api-ctype">GTestTrapFlags</p>
<table>
<tr>
<td class="name">TEST_TRAP_SILENCE_STDOUT</td>
<td class="value">128</td>
<td class="doc">Redirect stdout of the test child to
    `/dev/null` so it cannot be observed on the console during test
    runs. The actual output is still captured though to allow later
    tests with g_test_trap_assert_stdout().</td>
</tr>
<tr>
<td class="name">TEST_TRAP_SILENCE_STDERR</td>
<td class="value">256</td>
<td class="doc">Redirect stderr of the test child to
    `/dev/null` so it cannot be observed on the console during test
    runs. The actual output is still captured though to allow later
    tests with g_test_trap_assert_stderr().</td>
</tr>
<tr>
<td class="name">TEST_TRAP_INHERIT_STDIN</td>
<td class="value">512</td>
<td class="doc">If this flag is given, stdin of the
    child process is shared with stdin of its parent process.
    It is redirected to `/dev/null` otherwise.</td>
</tr>
</table>
</div>
<p class="api-heading">TraverseFlags</p>
<p class="api-doc">Specifies which nodes are visited during several of the tree
functions, including g_node_traverse() and g_node_find().</p>
<div class="api-notes">
  <p class="api-ctype">GTraverseFlags</p>
<table>
<tr>
<td class="name">TRAVERSE_LEAVES</td>
<td class="value">1</td>
<td class="doc">only leaf nodes should be visited. This name has
                    been introduced in 2.6, for older version use
                    %G_TRAVERSE_LEAFS.</td>
</tr>
<tr>
<td class="name">TRAVERSE_NON_LEAVES</td>
<td class="value">2</td>
<td class="doc">only non-leaf nodes should be visited. This
                        name has been introduced in 2.6, for older
                        version use %G_TRAVERSE_NON_LEAFS.</td>
</tr>
<tr>
<td class="name">TRAVERSE_ALL</td>
<td class="value">3</td>
<td class="doc">all nodes should be visited.</td>
</tr>
<tr>
<td class="name">TRAVERSE_MASK</td>
<td class="value">3</td>
<td class="doc">a mask of all traverse flags.</td>
</tr>
<tr>
<td class="name">TRAVERSE_LEAFS</td>
<td class="value">1</td>
<td class="doc">identical to %G_TRAVERSE_LEAVES.</td>
</tr>
<tr>
<td class="name">TRAVERSE_NON_LEAFS</td>
<td class="value">2</td>
<td class="doc">identical to %G_TRAVERSE_NON_LEAVES.</td>
</tr>
</table>
</div>
