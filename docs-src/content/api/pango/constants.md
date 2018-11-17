+++
title = "constants"
layout = "constants"
type = "api"
+++
<p class="api-heading">ANALYSIS_FLAG_CENTERED_BASELINE</p>
<p class="api-doc">Whether the segment should be shifted to center around the baseline.
Used in vertical writing directions mostly.</p>
<div class="api-notes">
  <p class="api-ctype">PANGO_ANALYSIS_FLAG_CENTERED_BASELINE</p>
  <p class="api-since">since 1.16</p>
</div>
<p class="api-heading">ANALYSIS_FLAG_IS_ELLIPSIS</p>
<p class="api-doc">This flag is used to mark runs that hold ellipsized text,
in an ellipsized layout.</p>
<div class="api-notes">
  <p class="api-ctype">PANGO_ANALYSIS_FLAG_IS_ELLIPSIS</p>
  <p class="api-since">since 1.36.7</p>
</div>
<p class="api-heading">ATTR_INDEX_FROM_TEXT_BEGINNING</p>
<p class="api-doc">This value can be used to set the start_index member of a #PangoAttribute
such that the attribute covers from the beginning of the text.</p>
<div class="api-notes">
  <p class="api-ctype">PANGO_ATTR_INDEX_FROM_TEXT_BEGINNING</p>
  <p class="api-since">since 1.24</p>
</div>
<p class="api-heading">GLYPH_EMPTY</p>
<p class="api-doc">The %PANGO_GLYPH_EMPTY macro represents a #PangoGlyph value that has a
 special meaning, which is a zero-width empty glyph.  This is useful for
example in shaper modules, to use as the glyph for various zero-width
Unicode characters (those passing pango_is_zero_width()).</p>
<div class="api-notes">
  <p class="api-ctype">PANGO_GLYPH_EMPTY</p>
</div>
<p class="api-heading">GLYPH_INVALID_INPUT</p>
<p class="api-doc">The %PANGO_GLYPH_INVALID_INPUT macro represents a #PangoGlyph value that has a
special meaning of invalid input.  #PangoLayout produces one such glyph
per invalid input UTF-8 byte and such a glyph is rendered as a crossed
box.

Note that this value is defined such that it has the %PANGO_GLYPH_UNKNOWN_FLAG
on.</p>
<div class="api-notes">
  <p class="api-ctype">PANGO_GLYPH_INVALID_INPUT</p>
  <p class="api-since">since 1.20</p>
</div>
<p class="api-heading">GLYPH_UNKNOWN_FLAG</p>
<p class="api-doc">The %PANGO_GLYPH_UNKNOWN_FLAG macro is a flag value that can be added to
a #gunichar value of a valid Unicode character, to produce a #PangoGlyph
value, representing an unknown-character glyph for the respective #gunichar.</p>
<div class="api-notes">
  <p class="api-ctype">PANGO_GLYPH_UNKNOWN_FLAG</p>
</div>
<p class="api-heading">SCALE</p>
<p class="api-doc">The %PANGO_SCALE macro represents the scale between dimensions used
for Pango distances and device units. (The definition of device
units is dependent on the output device; it will typically be pixels
for a screen, and points for a printer.) %PANGO_SCALE is currently
1024, but this may be changed in the future.

When setting font sizes, device units are always considered to be
points (as in "12 point font"), rather than pixels.</p>
<div class="api-notes">
  <p class="api-ctype">PANGO_SCALE</p>
</div>
