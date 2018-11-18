+++
title = "enums"
+++
<p class="api-heading">Alignment</p>
<p class="api-doc">A #PangoAlignment describes how to align the lines of a #PangoLayout within the
available space. If the #PangoLayout is set to justify
using pango_layout_set_justify(), this only has effect for partial lines.</p>
<div class="api-notes">
  <p class="api-ctype">PangoAlignment</p>
<table>
<tr>
<td class="name">PANGO_ALIGN_LEFT</td>
<td class="value">0</td>
<td class="doc">Put all available space on the right</td>
</tr>
<tr>
<td class="name">PANGO_ALIGN_CENTER</td>
<td class="value">1</td>
<td class="doc">Center the line within the available space</td>
</tr>
<tr>
<td class="name">PANGO_ALIGN_RIGHT</td>
<td class="value">2</td>
<td class="doc">Put all available space on the left</td>
</tr>
</table>
</div>
<p class="api-heading">AttrType</p>
<p class="api-doc">The #PangoAttrType
distinguishes between different types of attributes. Along with the
predefined values, it is possible to allocate additional values
for custom attributes using pango_attr_type_register(). The predefined
values are given below. The type of structure used to store the
attribute is listed in parentheses after the description.</p>
<div class="api-notes">
  <p class="api-ctype">PangoAttrType</p>
<table>
<tr>
<td class="name">PANGO_ATTR_INVALID</td>
<td class="value">0</td>
<td class="doc">does not happen</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_LANGUAGE</td>
<td class="value">1</td>
<td class="doc">language (#PangoAttrLanguage)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_FAMILY</td>
<td class="value">2</td>
<td class="doc">font family name list (#PangoAttrString)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_STYLE</td>
<td class="value">3</td>
<td class="doc">font slant style (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_WEIGHT</td>
<td class="value">4</td>
<td class="doc">font weight (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_VARIANT</td>
<td class="value">5</td>
<td class="doc">font variant (normal or small caps) (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_STRETCH</td>
<td class="value">6</td>
<td class="doc">font stretch (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_SIZE</td>
<td class="value">7</td>
<td class="doc">font size in points scaled by %PANGO_SCALE (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_FONT_DESC</td>
<td class="value">8</td>
<td class="doc">font description (#PangoAttrFontDesc)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_FOREGROUND</td>
<td class="value">9</td>
<td class="doc">foreground color (#PangoAttrColor)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_BACKGROUND</td>
<td class="value">10</td>
<td class="doc">background color (#PangoAttrColor)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_UNDERLINE</td>
<td class="value">11</td>
<td class="doc">whether the text has an underline (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_STRIKETHROUGH</td>
<td class="value">12</td>
<td class="doc">whether the text is struck-through (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_RISE</td>
<td class="value">13</td>
<td class="doc">baseline displacement (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_SHAPE</td>
<td class="value">14</td>
<td class="doc">shape (#PangoAttrShape)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_SCALE</td>
<td class="value">15</td>
<td class="doc">font size scale factor (#PangoAttrFloat)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_FALLBACK</td>
<td class="value">16</td>
<td class="doc">whether fallback is enabled (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_LETTER_SPACING</td>
<td class="value">17</td>
<td class="doc">letter spacing (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_UNDERLINE_COLOR</td>
<td class="value">18</td>
<td class="doc">underline color (#PangoAttrColor)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_STRIKETHROUGH_COLOR</td>
<td class="value">19</td>
<td class="doc">strikethrough color (#PangoAttrColor)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_ABSOLUTE_SIZE</td>
<td class="value">20</td>
<td class="doc">font size in pixels scaled by %PANGO_SCALE (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_GRAVITY</td>
<td class="value">21</td>
<td class="doc">base text gravity (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_GRAVITY_HINT</td>
<td class="value">22</td>
<td class="doc">gravity hint (#PangoAttrInt)</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_FONT_FEATURES</td>
<td class="value">23</td>
<td class="doc">OpenType font features (#PangoAttrString). Since 1.38</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_FOREGROUND_ALPHA</td>
<td class="value">24</td>
<td class="doc">foreground alpha (#PangoAttrInt). Since 1.38</td>
</tr>
<tr>
<td class="name">PANGO_ATTR_BACKGROUND_ALPHA</td>
<td class="value">25</td>
<td class="doc">background alpha (#PangoAttrInt). Since 1.38</td>
</tr>
</table>
</div>
<p class="api-heading">BidiType</p>
<p class="api-doc">The #PangoBidiType type represents the bidirectional character
type of a Unicode character as specified by the
<ulink url="http://www.unicode.org/reports/tr9/">Unicode bidirectional algorithm</ulink>.</p>
<div class="api-notes">
  <p class="api-ctype">PangoBidiType</p>
  <p class="api-since">since 1.22</p>
<table>
<tr>
<td class="name">PANGO_BIDI_TYPE_L</td>
<td class="value">0</td>
<td class="doc">Left-to-Right</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_LRE</td>
<td class="value">1</td>
<td class="doc">Left-to-Right Embedding</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_LRO</td>
<td class="value">2</td>
<td class="doc">Left-to-Right Override</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_R</td>
<td class="value">3</td>
<td class="doc">Right-to-Left</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_AL</td>
<td class="value">4</td>
<td class="doc">Right-to-Left Arabic</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_RLE</td>
<td class="value">5</td>
<td class="doc">Right-to-Left Embedding</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_RLO</td>
<td class="value">6</td>
<td class="doc">Right-to-Left Override</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_PDF</td>
<td class="value">7</td>
<td class="doc">Pop Directional Format</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_EN</td>
<td class="value">8</td>
<td class="doc">European Number</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_ES</td>
<td class="value">9</td>
<td class="doc">European Number Separator</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_ET</td>
<td class="value">10</td>
<td class="doc">European Number Terminator</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_AN</td>
<td class="value">11</td>
<td class="doc">Arabic Number</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_CS</td>
<td class="value">12</td>
<td class="doc">Common Number Separator</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_NSM</td>
<td class="value">13</td>
<td class="doc">Nonspacing Mark</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_BN</td>
<td class="value">14</td>
<td class="doc">Boundary Neutral</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_B</td>
<td class="value">15</td>
<td class="doc">Paragraph Separator</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_S</td>
<td class="value">16</td>
<td class="doc">Segment Separator</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_WS</td>
<td class="value">17</td>
<td class="doc">Whitespace</td>
</tr>
<tr>
<td class="name">PANGO_BIDI_TYPE_ON</td>
<td class="value">18</td>
<td class="doc">Other Neutrals</td>
</tr>
</table>
</div>
<p class="api-heading">CoverageLevel</p>
<p class="api-doc">Used to indicate how well a font can represent a particular Unicode
character point for a particular script.</p>
<div class="api-notes">
  <p class="api-ctype">PangoCoverageLevel</p>
<table>
<tr>
<td class="name">PANGO_COVERAGE_NONE</td>
<td class="value">0</td>
<td class="doc">The character is not representable with the font.</td>
</tr>
<tr>
<td class="name">PANGO_COVERAGE_FALLBACK</td>
<td class="value">1</td>
<td class="doc">The character is represented in a way that may be
comprehensible but is not the correct graphical form.
For instance, a Hangul character represented as a
a sequence of Jamos, or a Latin transliteration of a Cyrillic word.</td>
</tr>
<tr>
<td class="name">PANGO_COVERAGE_APPROXIMATE</td>
<td class="value">2</td>
<td class="doc">The character is represented as basically the correct
graphical form, but with a stylistic variant inappropriate for
the current script.</td>
</tr>
<tr>
<td class="name">PANGO_COVERAGE_EXACT</td>
<td class="value">3</td>
<td class="doc">The character is represented as the correct graphical form.</td>
</tr>
</table>
</div>
<p class="api-heading">Direction</p>
<p class="api-doc">The #PangoDirection type represents a direction in the
Unicode bidirectional algorithm; not every value in this
enumeration makes sense for every usage of #PangoDirection;
for example, the return value of pango_unichar_direction()
and pango_find_base_dir() cannot be %PANGO_DIRECTION_WEAK_LTR
or %PANGO_DIRECTION_WEAK_RTL, since every character is either
neutral or has a strong direction; on the other hand
%PANGO_DIRECTION_NEUTRAL doesn't make sense to pass
to pango_itemize_with_base_dir().

The %PANGO_DIRECTION_TTB_LTR, %PANGO_DIRECTION_TTB_RTL
values come from an earlier interpretation of this
enumeration as the writing direction of a block of
text and are no longer used; See #PangoGravity for how
vertical text is handled in Pango.</p>
<div class="api-notes">
  <p class="api-ctype">PangoDirection</p>
<table>
<tr>
<td class="name">PANGO_DIRECTION_LTR</td>
<td class="value">0</td>
<td class="doc">A strong left-to-right direction</td>
</tr>
<tr>
<td class="name">PANGO_DIRECTION_RTL</td>
<td class="value">1</td>
<td class="doc">A strong right-to-left direction</td>
</tr>
<tr>
<td class="name">PANGO_DIRECTION_TTB_LTR</td>
<td class="value">2</td>
<td class="doc">Deprecated value; treated the
  same as %PANGO_DIRECTION_RTL.</td>
</tr>
<tr>
<td class="name">PANGO_DIRECTION_TTB_RTL</td>
<td class="value">3</td>
<td class="doc">Deprecated value; treated the
  same as %PANGO_DIRECTION_LTR</td>
</tr>
<tr>
<td class="name">PANGO_DIRECTION_WEAK_LTR</td>
<td class="value">4</td>
<td class="doc">A weak left-to-right direction</td>
</tr>
<tr>
<td class="name">PANGO_DIRECTION_WEAK_RTL</td>
<td class="value">5</td>
<td class="doc">A weak right-to-left direction</td>
</tr>
<tr>
<td class="name">PANGO_DIRECTION_NEUTRAL</td>
<td class="value">6</td>
<td class="doc">No direction specified</td>
</tr>
</table>
</div>
<p class="api-heading">EllipsizeMode</p>
<p class="api-doc">The #PangoEllipsizeMode type describes what sort of (if any)
ellipsization should be applied to a line of text. In
the ellipsization process characters are removed from the
text in order to make it fit to a given width and replaced
with an ellipsis.</p>
<div class="api-notes">
  <p class="api-ctype">PangoEllipsizeMode</p>
<table>
<tr>
<td class="name">PANGO_ELLIPSIZE_NONE</td>
<td class="value">0</td>
<td class="doc">No ellipsization</td>
</tr>
<tr>
<td class="name">PANGO_ELLIPSIZE_START</td>
<td class="value">1</td>
<td class="doc">Omit characters at the start of the text</td>
</tr>
<tr>
<td class="name">PANGO_ELLIPSIZE_MIDDLE</td>
<td class="value">2</td>
<td class="doc">Omit characters in the middle of the text</td>
</tr>
<tr>
<td class="name">PANGO_ELLIPSIZE_END</td>
<td class="value">3</td>
<td class="doc">Omit characters at the end of the text</td>
</tr>
</table>
</div>
<p class="api-heading">Gravity</p>
<p class="api-doc">The #PangoGravity type represents the orientation of glyphs in a segment
of text.  This is useful when rendering vertical text layouts.  In
those situations, the layout is rotated using a non-identity PangoMatrix,
and then glyph orientation is controlled using #PangoGravity.
Not every value in this enumeration makes sense for every usage of
#PangoGravity; for example, %PANGO_GRAVITY_AUTO only can be passed to
pango_context_set_base_gravity() and can only be returned by
pango_context_get_base_gravity().

See also: #PangoGravityHint</p>
<div class="api-notes">
  <p class="api-ctype">PangoGravity</p>
  <p class="api-since">since 1.16</p>
<table>
<tr>
<td class="name">PANGO_GRAVITY_SOUTH</td>
<td class="value">0</td>
<td class="doc">Glyphs stand upright (default)</td>
</tr>
<tr>
<td class="name">PANGO_GRAVITY_EAST</td>
<td class="value">1</td>
<td class="doc">Glyphs are rotated 90 degrees clockwise</td>
</tr>
<tr>
<td class="name">PANGO_GRAVITY_NORTH</td>
<td class="value">2</td>
<td class="doc">Glyphs are upside-down</td>
</tr>
<tr>
<td class="name">PANGO_GRAVITY_WEST</td>
<td class="value">3</td>
<td class="doc">Glyphs are rotated 90 degrees counter-clockwise</td>
</tr>
<tr>
<td class="name">PANGO_GRAVITY_AUTO</td>
<td class="value">4</td>
<td class="doc">Gravity is resolved from the context matrix</td>
</tr>
</table>
</div>
<p class="api-heading">GravityHint</p>
<p class="api-doc">The #PangoGravityHint defines how horizontal scripts should behave in a
vertical context.  That is, English excerpt in a vertical paragraph for
example.

See #PangoGravity.</p>
<div class="api-notes">
  <p class="api-ctype">PangoGravityHint</p>
  <p class="api-since">since 1.16</p>
<table>
<tr>
<td class="name">PANGO_GRAVITY_HINT_NATURAL</td>
<td class="value">0</td>
<td class="doc">scripts will take their natural gravity based
on the base gravity and the script.  This is the default.</td>
</tr>
<tr>
<td class="name">PANGO_GRAVITY_HINT_STRONG</td>
<td class="value">1</td>
<td class="doc">always use the base gravity set, regardless of
the script.</td>
</tr>
<tr>
<td class="name">PANGO_GRAVITY_HINT_LINE</td>
<td class="value">2</td>
<td class="doc">for scripts not in their natural direction (eg.
Latin in East gravity), choose per-script gravity such that every script
respects the line progression.  This means, Latin and Arabic will take
opposite gravities and both flow top-to-bottom for example.</td>
</tr>
</table>
</div>
<p class="api-heading">RenderPart</p>
<p class="api-doc">#PangoRenderPart defines different items to render for such
purposes as setting colors.</p>
<div class="api-notes">
  <p class="api-ctype">PangoRenderPart</p>
  <p class="api-since">since 1.8</p>
<table>
<tr>
<td class="name">PANGO_RENDER_PART_FOREGROUND</td>
<td class="value">0</td>
<td class="doc">the text itself</td>
</tr>
<tr>
<td class="name">PANGO_RENDER_PART_BACKGROUND</td>
<td class="value">1</td>
<td class="doc">the area behind the text</td>
</tr>
<tr>
<td class="name">PANGO_RENDER_PART_UNDERLINE</td>
<td class="value">2</td>
<td class="doc">underlines</td>
</tr>
<tr>
<td class="name">PANGO_RENDER_PART_STRIKETHROUGH</td>
<td class="value">3</td>
<td class="doc">strikethrough lines</td>
</tr>
</table>
</div>
<p class="api-heading">Script</p>
<p class="api-doc">The #PangoScript enumeration identifies different writing
systems. The values correspond to the names as defined in the
Unicode standard.
Note that new types may be added in the future. Applications should be ready
to handle unknown values.  This enumeration is interchangeable with
#GUnicodeScript.  See <ulink
url="http://www.unicode.org/reports/tr24/">Unicode Standard Annex
#24: Script names</ulink>.</p>
<div class="api-notes">
  <p class="api-ctype">PangoScript</p>
<table>
<tr>
<td class="name">PANGO_SCRIPT_INVALID_CODE</td>
<td class="value">-1</td>
<td class="doc">a value never returned from pango_script_for_unichar()</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_COMMON</td>
<td class="value">0</td>
<td class="doc">a character used by multiple different scripts</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_INHERITED</td>
<td class="value">1</td>
<td class="doc">a mark glyph that takes its script from the
base glyph to which it is attached</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_ARABIC</td>
<td class="value">2</td>
<td class="doc">Arabic</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_ARMENIAN</td>
<td class="value">3</td>
<td class="doc">Armenian</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_BENGALI</td>
<td class="value">4</td>
<td class="doc">Bengali</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_BOPOMOFO</td>
<td class="value">5</td>
<td class="doc">Bopomofo</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_CHEROKEE</td>
<td class="value">6</td>
<td class="doc">Cherokee</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_COPTIC</td>
<td class="value">7</td>
<td class="doc">Coptic</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_CYRILLIC</td>
<td class="value">8</td>
<td class="doc">Cyrillic</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_DESERET</td>
<td class="value">9</td>
<td class="doc">Deseret</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_DEVANAGARI</td>
<td class="value">10</td>
<td class="doc">Devanagari</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_ETHIOPIC</td>
<td class="value">11</td>
<td class="doc">Ethiopic</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_GEORGIAN</td>
<td class="value">12</td>
<td class="doc">Georgian</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_GOTHIC</td>
<td class="value">13</td>
<td class="doc">Gothic</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_GREEK</td>
<td class="value">14</td>
<td class="doc">Greek</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_GUJARATI</td>
<td class="value">15</td>
<td class="doc">Gujarati</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_GURMUKHI</td>
<td class="value">16</td>
<td class="doc">Gurmukhi</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_HAN</td>
<td class="value">17</td>
<td class="doc">Han</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_HANGUL</td>
<td class="value">18</td>
<td class="doc">Hangul</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_HEBREW</td>
<td class="value">19</td>
<td class="doc">Hebrew</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_HIRAGANA</td>
<td class="value">20</td>
<td class="doc">Hiragana</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_KANNADA</td>
<td class="value">21</td>
<td class="doc">Kannada</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_KATAKANA</td>
<td class="value">22</td>
<td class="doc">Katakana</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_KHMER</td>
<td class="value">23</td>
<td class="doc">Khmer</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_LAO</td>
<td class="value">24</td>
<td class="doc">Lao</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_LATIN</td>
<td class="value">25</td>
<td class="doc">Latin</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MALAYALAM</td>
<td class="value">26</td>
<td class="doc">Malayalam</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MONGOLIAN</td>
<td class="value">27</td>
<td class="doc">Mongolian</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MYANMAR</td>
<td class="value">28</td>
<td class="doc">Myanmar</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_OGHAM</td>
<td class="value">29</td>
<td class="doc">Ogham</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_OLD_ITALIC</td>
<td class="value">30</td>
<td class="doc">Old Italic</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_ORIYA</td>
<td class="value">31</td>
<td class="doc">Oriya</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_RUNIC</td>
<td class="value">32</td>
<td class="doc">Runic</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_SINHALA</td>
<td class="value">33</td>
<td class="doc">Sinhala</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_SYRIAC</td>
<td class="value">34</td>
<td class="doc">Syriac</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_TAMIL</td>
<td class="value">35</td>
<td class="doc">Tamil</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_TELUGU</td>
<td class="value">36</td>
<td class="doc">Telugu</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_THAANA</td>
<td class="value">37</td>
<td class="doc">Thaana</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_THAI</td>
<td class="value">38</td>
<td class="doc">Thai</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_TIBETAN</td>
<td class="value">39</td>
<td class="doc">Tibetan</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_CANADIAN_ABORIGINAL</td>
<td class="value">40</td>
<td class="doc">Canadian Aboriginal</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_YI</td>
<td class="value">41</td>
<td class="doc">Yi</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_TAGALOG</td>
<td class="value">42</td>
<td class="doc">Tagalog</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_HANUNOO</td>
<td class="value">43</td>
<td class="doc">Hanunoo</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_BUHID</td>
<td class="value">44</td>
<td class="doc">Buhid</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_TAGBANWA</td>
<td class="value">45</td>
<td class="doc">Tagbanwa</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_BRAILLE</td>
<td class="value">46</td>
<td class="doc">Braille</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_CYPRIOT</td>
<td class="value">47</td>
<td class="doc">Cypriot</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_LIMBU</td>
<td class="value">48</td>
<td class="doc">Limbu</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_OSMANYA</td>
<td class="value">49</td>
<td class="doc">Osmanya</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_SHAVIAN</td>
<td class="value">50</td>
<td class="doc">Shavian</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_LINEAR_B</td>
<td class="value">51</td>
<td class="doc">Linear B</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_TAI_LE</td>
<td class="value">52</td>
<td class="doc">Tai Le</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_UGARITIC</td>
<td class="value">53</td>
<td class="doc">Ugaritic</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_NEW_TAI_LUE</td>
<td class="value">54</td>
<td class="doc">New Tai Lue. Since 1.10</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_BUGINESE</td>
<td class="value">55</td>
<td class="doc">Buginese. Since 1.10</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_GLAGOLITIC</td>
<td class="value">56</td>
<td class="doc">Glagolitic. Since 1.10</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_TIFINAGH</td>
<td class="value">57</td>
<td class="doc">Tifinagh. Since 1.10</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_SYLOTI_NAGRI</td>
<td class="value">58</td>
<td class="doc">Syloti Nagri. Since 1.10</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_OLD_PERSIAN</td>
<td class="value">59</td>
<td class="doc">Old Persian. Since 1.10</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_KHAROSHTHI</td>
<td class="value">60</td>
<td class="doc">Kharoshthi. Since 1.10</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_UNKNOWN</td>
<td class="value">61</td>
<td class="doc">an unassigned code point. Since 1.14</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_BALINESE</td>
<td class="value">62</td>
<td class="doc">Balinese. Since 1.14</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_CUNEIFORM</td>
<td class="value">63</td>
<td class="doc">Cuneiform. Since 1.14</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_PHOENICIAN</td>
<td class="value">64</td>
<td class="doc">Phoenician. Since 1.14</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_PHAGS_PA</td>
<td class="value">65</td>
<td class="doc">Phags-pa. Since 1.14</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_NKO</td>
<td class="value">66</td>
<td class="doc">N'Ko. Since 1.14</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_KAYAH_LI</td>
<td class="value">67</td>
<td class="doc">Kayah Li. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_LEPCHA</td>
<td class="value">68</td>
<td class="doc">Lepcha. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_REJANG</td>
<td class="value">69</td>
<td class="doc">Rejang. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_SUNDANESE</td>
<td class="value">70</td>
<td class="doc">Sundanese. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_SAURASHTRA</td>
<td class="value">71</td>
<td class="doc">Saurashtra. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_CHAM</td>
<td class="value">72</td>
<td class="doc">Cham. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_OL_CHIKI</td>
<td class="value">73</td>
<td class="doc">Ol Chiki. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_VAI</td>
<td class="value">74</td>
<td class="doc">Vai. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_CARIAN</td>
<td class="value">75</td>
<td class="doc">Carian. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_LYCIAN</td>
<td class="value">76</td>
<td class="doc">Lycian. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_LYDIAN</td>
<td class="value">77</td>
<td class="doc">Lydian. Since 1.20.1</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_BATAK</td>
<td class="value">78</td>
<td class="doc">Batak. Since 1.32</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_BRAHMI</td>
<td class="value">79</td>
<td class="doc">Brahmi. Since 1.32</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MANDAIC</td>
<td class="value">80</td>
<td class="doc">Mandaic. Since 1.32</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_CHAKMA</td>
<td class="value">81</td>
<td class="doc">Chakma. Since: 1.32</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MEROITIC_CURSIVE</td>
<td class="value">82</td>
<td class="doc">Meroitic Cursive. Since: 1.32</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MEROITIC_HIEROGLYPHS</td>
<td class="value">83</td>
<td class="doc">Meroitic Hieroglyphs. Since: 1.32</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MIAO</td>
<td class="value">84</td>
<td class="doc">Miao. Since: 1.32</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_SHARADA</td>
<td class="value">85</td>
<td class="doc">Sharada. Since: 1.32</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_SORA_SOMPENG</td>
<td class="value">86</td>
<td class="doc">Sora Sompeng. Since: 1.32</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_TAKRI</td>
<td class="value">87</td>
<td class="doc">Takri. Since: 1.32</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_BASSA_VAH</td>
<td class="value">88</td>
<td class="doc">Bassa. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_CAUCASIAN_ALBANIAN</td>
<td class="value">89</td>
<td class="doc">Caucasian Albanian. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_DUPLOYAN</td>
<td class="value">90</td>
<td class="doc">Duployan. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_ELBASAN</td>
<td class="value">91</td>
<td class="doc">Elbasan. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_GRANTHA</td>
<td class="value">92</td>
<td class="doc">Grantha. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_KHOJKI</td>
<td class="value">93</td>
<td class="doc">Kjohki. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_KHUDAWADI</td>
<td class="value">94</td>
<td class="doc">Khudawadi, Sindhi. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_LINEAR_A</td>
<td class="value">95</td>
<td class="doc">Linear A. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MAHAJANI</td>
<td class="value">96</td>
<td class="doc">Mahajani. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MANICHAEAN</td>
<td class="value">97</td>
<td class="doc">Manichaean. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MENDE_KIKAKUI</td>
<td class="value">98</td>
<td class="doc">Mende Kikakui. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MODI</td>
<td class="value">99</td>
<td class="doc">Modi. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MRO</td>
<td class="value">100</td>
<td class="doc">Mro. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_NABATAEAN</td>
<td class="value">101</td>
<td class="doc">Nabataean. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_OLD_NORTH_ARABIAN</td>
<td class="value">102</td>
<td class="doc">Old North Arabian. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_OLD_PERMIC</td>
<td class="value">103</td>
<td class="doc">Old Permic. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_PAHAWH_HMONG</td>
<td class="value">104</td>
<td class="doc">Pahawh Hmong. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_PALMYRENE</td>
<td class="value">105</td>
<td class="doc">Palmyrene. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_PAU_CIN_HAU</td>
<td class="value">106</td>
<td class="doc">Pau Cin Hau. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_PSALTER_PAHLAVI</td>
<td class="value">107</td>
<td class="doc">Psalter Pahlavi. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_SIDDHAM</td>
<td class="value">108</td>
<td class="doc">Siddham. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_TIRHUTA</td>
<td class="value">109</td>
<td class="doc">Tirhuta. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_WARANG_CITI</td>
<td class="value">110</td>
<td class="doc">Warang Citi. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_AHOM</td>
<td class="value">111</td>
<td class="doc">Ahom. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_ANATOLIAN_HIEROGLYPHS</td>
<td class="value">112</td>
<td class="doc">Anatolian Hieroglyphs. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_HATRAN</td>
<td class="value">113</td>
<td class="doc">Hatran. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_MULTANI</td>
<td class="value">114</td>
<td class="doc">Multani. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_OLD_HUNGARIAN</td>
<td class="value">115</td>
<td class="doc">Old Hungarian. Since: 1.40</td>
</tr>
<tr>
<td class="name">PANGO_SCRIPT_SIGNWRITING</td>
<td class="value">116</td>
<td class="doc">Signwriting. Since: 1.40</td>
</tr>
</table>
</div>
<p class="api-heading">Stretch</p>
<p class="api-doc">An enumeration specifying the width of the font relative to other designs
within a family.</p>
<div class="api-notes">
  <p class="api-ctype">PangoStretch</p>
<table>
<tr>
<td class="name">PANGO_STRETCH_ULTRA_CONDENSED</td>
<td class="value">0</td>
<td class="doc">ultra condensed width</td>
</tr>
<tr>
<td class="name">PANGO_STRETCH_EXTRA_CONDENSED</td>
<td class="value">1</td>
<td class="doc">extra condensed width</td>
</tr>
<tr>
<td class="name">PANGO_STRETCH_CONDENSED</td>
<td class="value">2</td>
<td class="doc">condensed width</td>
</tr>
<tr>
<td class="name">PANGO_STRETCH_SEMI_CONDENSED</td>
<td class="value">3</td>
<td class="doc">semi condensed width</td>
</tr>
<tr>
<td class="name">PANGO_STRETCH_NORMAL</td>
<td class="value">4</td>
<td class="doc">the normal width</td>
</tr>
<tr>
<td class="name">PANGO_STRETCH_SEMI_EXPANDED</td>
<td class="value">5</td>
<td class="doc">semi expanded width</td>
</tr>
<tr>
<td class="name">PANGO_STRETCH_EXPANDED</td>
<td class="value">6</td>
<td class="doc">expanded width</td>
</tr>
<tr>
<td class="name">PANGO_STRETCH_EXTRA_EXPANDED</td>
<td class="value">7</td>
<td class="doc">extra expanded width</td>
</tr>
<tr>
<td class="name">PANGO_STRETCH_ULTRA_EXPANDED</td>
<td class="value">8</td>
<td class="doc">ultra expanded width</td>
</tr>
</table>
</div>
<p class="api-heading">Style</p>
<p class="api-doc">An enumeration specifying the various slant styles possible for a font.</p>
<div class="api-notes">
  <p class="api-ctype">PangoStyle</p>
<table>
<tr>
<td class="name">PANGO_STYLE_NORMAL</td>
<td class="value">0</td>
<td class="doc">the font is upright.</td>
</tr>
<tr>
<td class="name">PANGO_STYLE_OBLIQUE</td>
<td class="value">1</td>
<td class="doc">the font is slanted, but in a roman style.</td>
</tr>
<tr>
<td class="name">PANGO_STYLE_ITALIC</td>
<td class="value">2</td>
<td class="doc">the font is slanted in an italic style.</td>
</tr>
</table>
</div>
<p class="api-heading">TabAlign</p>
<p class="api-doc">A #PangoTabAlign specifies where a tab stop appears relative to the text.</p>
<div class="api-notes">
  <p class="api-ctype">PangoTabAlign</p>
<table>
<tr>
<td class="name">PANGO_TAB_LEFT</td>
<td class="value">0</td>
<td class="doc">the tab stop appears to the left of the text.</td>
</tr>
</table>
</div>
<p class="api-heading">Underline</p>
<p class="api-doc">The #PangoUnderline enumeration is used to specify
whether text should be underlined, and if so, the type
of underlining.</p>
<div class="api-notes">
  <p class="api-ctype">PangoUnderline</p>
<table>
<tr>
<td class="name">PANGO_UNDERLINE_NONE</td>
<td class="value">0</td>
<td class="doc">no underline should be drawn</td>
</tr>
<tr>
<td class="name">PANGO_UNDERLINE_SINGLE</td>
<td class="value">1</td>
<td class="doc">a single underline should be drawn</td>
</tr>
<tr>
<td class="name">PANGO_UNDERLINE_DOUBLE</td>
<td class="value">2</td>
<td class="doc">a double underline should be drawn</td>
</tr>
<tr>
<td class="name">PANGO_UNDERLINE_LOW</td>
<td class="value">3</td>
<td class="doc">a single underline should be drawn at a position
beneath the ink extents of the text being
underlined. This should be used only for underlining
single characters, such as for keyboard
accelerators. %PANGO_UNDERLINE_SINGLE should
be used for extended portions of text.</td>
</tr>
<tr>
<td class="name">PANGO_UNDERLINE_ERROR</td>
<td class="value">4</td>
<td class="doc">a wavy underline should be drawn below.
This underline is typically used to indicate
an error such as a possilble mispelling; in some
cases a contrasting color may automatically
be used. This type of underlining is available since Pango 1.4.</td>
</tr>
</table>
</div>
<p class="api-heading">Variant</p>
<p class="api-doc">An enumeration specifying capitalization variant of the font.</p>
<div class="api-notes">
  <p class="api-ctype">PangoVariant</p>
<table>
<tr>
<td class="name">PANGO_VARIANT_NORMAL</td>
<td class="value">0</td>
<td class="doc">A normal font.</td>
</tr>
<tr>
<td class="name">PANGO_VARIANT_SMALL_CAPS</td>
<td class="value">1</td>
<td class="doc">A font with the lower case characters
replaced by smaller variants of the capital characters.</td>
</tr>
</table>
</div>
<p class="api-heading">Weight</p>
<p class="api-doc">An enumeration specifying the weight (boldness) of a font. This is a numerical
value ranging from 100 to 1000, but there are some predefined values:</p>
<div class="api-notes">
  <p class="api-ctype">PangoWeight</p>
<table>
<tr>
<td class="name">PANGO_WEIGHT_THIN</td>
<td class="value">100</td>
<td class="doc">the thin weight (= 100; Since: 1.24)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_ULTRALIGHT</td>
<td class="value">200</td>
<td class="doc">the ultralight weight (= 200)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_LIGHT</td>
<td class="value">300</td>
<td class="doc">the light weight (= 300)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_SEMILIGHT</td>
<td class="value">350</td>
<td class="doc">the semilight weight (= 350; Since: 1.36.7)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_BOOK</td>
<td class="value">380</td>
<td class="doc">the book weight (= 380; Since: 1.24)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_NORMAL</td>
<td class="value">400</td>
<td class="doc">the default weight (= 400)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_MEDIUM</td>
<td class="value">500</td>
<td class="doc">the normal weight (= 500; Since: 1.24)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_SEMIBOLD</td>
<td class="value">600</td>
<td class="doc">the semibold weight (= 600)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_BOLD</td>
<td class="value">700</td>
<td class="doc">the bold weight (= 700)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_ULTRABOLD</td>
<td class="value">800</td>
<td class="doc">the ultrabold weight (= 800)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_HEAVY</td>
<td class="value">900</td>
<td class="doc">the heavy weight (= 900)</td>
</tr>
<tr>
<td class="name">PANGO_WEIGHT_ULTRAHEAVY</td>
<td class="value">1000</td>
<td class="doc">the ultraheavy weight (= 1000; Since: 1.24)</td>
</tr>
</table>
</div>
<p class="api-heading">WrapMode</p>
<p class="api-doc">A #PangoWrapMode describes how to wrap the lines of a #PangoLayout to the desired width.</p>
<div class="api-notes">
  <p class="api-ctype">PangoWrapMode</p>
<table>
<tr>
<td class="name">PANGO_WRAP_WORD</td>
<td class="value">0</td>
<td class="doc">wrap lines at word boundaries.</td>
</tr>
<tr>
<td class="name">PANGO_WRAP_CHAR</td>
<td class="value">1</td>
<td class="doc">wrap lines at character boundaries.</td>
</tr>
<tr>
<td class="name">PANGO_WRAP_WORD_CHAR</td>
<td class="value">2</td>
<td class="doc">wrap lines at word boundaries, but fall back to character boundaries if there is not
enough space for a full word.</td>
</tr>
</table>
</div>