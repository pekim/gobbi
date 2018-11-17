+++
title = "constants"
layout = "constants"
type = "api"
+++
<p class="api-heading">ANALYSIS_FLAG_CENTERED_BASELINE</p>
<p class="api-doc">Whether the segment should be shifted to center around the baseline.
Used in vertical writing directions mostly.</p>
<p class="api-ctype">PANGO_ANALYSIS_FLAG_CENTERED_BASELINE</p>
<p class="api-heading">ANALYSIS_FLAG_IS_ELLIPSIS</p>
<p class="api-doc">This flag is used to mark runs that hold ellipsized text,
in an ellipsized layout.</p>
<p class="api-ctype">PANGO_ANALYSIS_FLAG_IS_ELLIPSIS</p>
<p class="api-heading">ATTR_INDEX_FROM_TEXT_BEGINNING</p>
<p class="api-doc">This value can be used to set the start_index member of a #PangoAttribute
such that the attribute covers from the beginning of the text.</p>
<p class="api-ctype">PANGO_ATTR_INDEX_FROM_TEXT_BEGINNING</p>
<p class="api-heading">ENGINE_TYPE_LANG</p>
<p class="api-doc">A string constant defining the engine type for language engines.
These engines derive from #PangoEngineLang.</p>
<p class="api-ctype">PANGO_ENGINE_TYPE_LANG</p>
<p class="api-heading">ENGINE_TYPE_SHAPE</p>
<p class="api-doc">A string constant defining the engine type for shaping engines.
These engines derive from #PangoEngineShape.</p>
<p class="api-ctype">PANGO_ENGINE_TYPE_SHAPE</p>
<p class="api-heading">GLYPH_EMPTY</p>
<p class="api-doc">The %PANGO_GLYPH_EMPTY macro represents a #PangoGlyph value that has a
 special meaning, which is a zero-width empty glyph.  This is useful for
example in shaper modules, to use as the glyph for various zero-width
Unicode characters (those passing pango_is_zero_width()).</p>
<p class="api-ctype">PANGO_GLYPH_EMPTY</p>
<p class="api-heading">GLYPH_INVALID_INPUT</p>
<p class="api-doc">The %PANGO_GLYPH_INVALID_INPUT macro represents a #PangoGlyph value that has a
special meaning of invalid input.  #PangoLayout produces one such glyph
per invalid input UTF-8 byte and such a glyph is rendered as a crossed
box.

Note that this value is defined such that it has the %PANGO_GLYPH_UNKNOWN_FLAG
on.</p>
<p class="api-ctype">PANGO_GLYPH_INVALID_INPUT</p>
<p class="api-heading">GLYPH_UNKNOWN_FLAG</p>
<p class="api-doc">The %PANGO_GLYPH_UNKNOWN_FLAG macro is a flag value that can be added to
a #gunichar value of a valid Unicode character, to produce a #PangoGlyph
value, representing an unknown-character glyph for the respective #gunichar.</p>
<p class="api-ctype">PANGO_GLYPH_UNKNOWN_FLAG</p>
<p class="api-heading">RENDER_TYPE_NONE</p>
<p class="api-doc">A string constant defining the render type
for engines that are not rendering-system specific.</p>
<p class="api-ctype">PANGO_RENDER_TYPE_NONE</p>
<p class="api-heading">SCALE</p>
<p class="api-doc">The %PANGO_SCALE macro represents the scale between dimensions used
for Pango distances and device units. (The definition of device
units is dependent on the output device; it will typically be pixels
for a screen, and points for a printer.) %PANGO_SCALE is currently
1024, but this may be changed in the future.

When setting font sizes, device units are always considered to be
points (as in "12 point font"), rather than pixels.</p>
<p class="api-ctype">PANGO_SCALE</p>
<p class="api-heading">UNKNOWN_GLYPH_HEIGHT</p>
<p class="api-ctype">PANGO_UNKNOWN_GLYPH_HEIGHT</p>
<p class="api-heading">UNKNOWN_GLYPH_WIDTH</p>
<p class="api-ctype">PANGO_UNKNOWN_GLYPH_WIDTH</p>
<p class="api-heading">VERSION_MIN_REQUIRED</p>
<p class="api-doc">A macro that should be defined by the user prior to including
the pango.h header.
The definition should be one of the predefined Pango version
macros: %PANGO_VERSION_1_2, %PANGO_VERSION_1_4,...

This macro defines the earliest version of Pango that the package is
required to be able to compile against.

If the compiler is configured to warn about the use of deprecated
functions, then using functions that were deprecated in version
%PANGO_VERSION_MIN_REQUIRED or earlier will cause warnings (but
using functions deprecated in later releases will not).</p>
<p class="api-ctype">PANGO_VERSION_MIN_REQUIRED</p>
