+++
title = "constants"
+++
<p class="api-heading">BINARY_AGE</p>
<p class="api-doc">Like atk_get_binary_age(), but from the headers used at
application compile time, rather than from the library linked
against at application run time.</p>
<div class="api-notes">
  <p class="api-ctype">C type: ATK_BINARY_AGE</p>
  <p class="api-since">since 2.7.4</p>
</div>
<p class="api-heading">INTERFACE_AGE</p>
<p class="api-doc">Like atk_get_interface_age(), but from the headers used at
application compile time, rather than from the library linked
against at application run time.</p>
<div class="api-notes">
  <p class="api-ctype">C type: ATK_INTERFACE_AGE</p>
  <p class="api-since">since 2.7.4</p>
</div>
<p class="api-heading">MAJOR_VERSION</p>
<p class="api-doc">Like atk_get_major_version(), but from the headers used at
application compile time, rather than from the library linked
against at application run time.</p>
<div class="api-notes">
  <p class="api-ctype">C type: ATK_MAJOR_VERSION</p>
  <p class="api-since">since 2.7.4</p>
</div>
<p class="api-heading">MICRO_VERSION</p>
<p class="api-doc">Like atk_get_micro_version(), but from the headers used at
application compile time, rather than from the library linked
against at application run time.</p>
<div class="api-notes">
  <p class="api-ctype">C type: ATK_MICRO_VERSION</p>
  <p class="api-since">since 2.7.4</p>
</div>
<p class="api-heading">MINOR_VERSION</p>
<p class="api-doc">Like atk_get_minor_version(), but from the headers used at
application compile time, rather than from the library linked
against at application run time.</p>
<div class="api-notes">
  <p class="api-ctype">C type: ATK_MINOR_VERSION</p>
  <p class="api-since">since 2.7.4</p>
</div>
<p class="api-heading">VERSION_MIN_REQUIRED</p>
<p class="api-doc">A macro that should be defined by the user prior to including
the atk/atk.h header.
The definition should be one of the predefined ATK version
macros: %ATK_VERSION_2_12, %ATK_VERSION_2_14,...

This macro defines the earliest version of ATK that the package is
required to be able to compile against.

If the compiler is configured to warn about the use of deprecated
functions, then using functions that were deprecated in version
%ATK_VERSION_MIN_REQUIRED or earlier will cause warnings (but
using functions deprecated in later releases will not).</p>
<div class="api-notes">
  <p class="api-ctype">C type: ATK_VERSION_MIN_REQUIRED</p>
  <p class="api-since">since 2.14</p>
</div>
