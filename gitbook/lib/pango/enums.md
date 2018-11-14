# `pango` enums

## `Alignment`

A #PangoAlignment describes how to align the lines of a #PangoLayout within the
available space. If the #PangoLayout is set to justify
using pango_layout_set_justify(), this only has effect for partial lines.

C - `PangoAlignment`

### PANGO_ALIGN_LEFT = 0
Put all available space on the right

### PANGO_ALIGN_CENTER = 1
Center the line within the available space

### PANGO_ALIGN_RIGHT = 2
Put all available space on the left


## `AttrType`

The #PangoAttrType
distinguishes between different types of attributes. Along with the
predefined values, it is possible to allocate additional values
for custom attributes using pango_attr_type_register(). The predefined
values are given below. The type of structure used to store the
attribute is listed in parentheses after the description.

C - `PangoAttrType`

### PANGO_ATTR_INVALID = 0
does not happen

### PANGO_ATTR_LANGUAGE = 1
language (#PangoAttrLanguage)

### PANGO_ATTR_FAMILY = 2
font family name list (#PangoAttrString)

### PANGO_ATTR_STYLE = 3
font slant style (#PangoAttrInt)

### PANGO_ATTR_WEIGHT = 4
font weight (#PangoAttrInt)

### PANGO_ATTR_VARIANT = 5
font variant (normal or small caps) (#PangoAttrInt)

### PANGO_ATTR_STRETCH = 6
font stretch (#PangoAttrInt)

### PANGO_ATTR_SIZE = 7
font size in points scaled by %PANGO_SCALE (#PangoAttrInt)

### PANGO_ATTR_FONT_DESC = 8
font description (#PangoAttrFontDesc)

### PANGO_ATTR_FOREGROUND = 9
foreground color (#PangoAttrColor)

### PANGO_ATTR_BACKGROUND = 10
background color (#PangoAttrColor)

### PANGO_ATTR_UNDERLINE = 11
whether the text has an underline (#PangoAttrInt)

### PANGO_ATTR_STRIKETHROUGH = 12
whether the text is struck-through (#PangoAttrInt)

### PANGO_ATTR_RISE = 13
baseline displacement (#PangoAttrInt)

### PANGO_ATTR_SHAPE = 14
shape (#PangoAttrShape)

### PANGO_ATTR_SCALE = 15
font size scale factor (#PangoAttrFloat)

### PANGO_ATTR_FALLBACK = 16
whether fallback is enabled (#PangoAttrInt)

### PANGO_ATTR_LETTER_SPACING = 17
letter spacing (#PangoAttrInt)

### PANGO_ATTR_UNDERLINE_COLOR = 18
underline color (#PangoAttrColor)

### PANGO_ATTR_STRIKETHROUGH_COLOR = 19
strikethrough color (#PangoAttrColor)

### PANGO_ATTR_ABSOLUTE_SIZE = 20
font size in pixels scaled by %PANGO_SCALE (#PangoAttrInt)

### PANGO_ATTR_GRAVITY = 21
base text gravity (#PangoAttrInt)

### PANGO_ATTR_GRAVITY_HINT = 22
gravity hint (#PangoAttrInt)

### PANGO_ATTR_FONT_FEATURES = 23
OpenType font features (#PangoAttrString). Since 1.38

### PANGO_ATTR_FOREGROUND_ALPHA = 24
foreground alpha (#PangoAttrInt). Since 1.38

### PANGO_ATTR_BACKGROUND_ALPHA = 25
background alpha (#PangoAttrInt). Since 1.38


## `BidiType`

The #PangoBidiType type represents the bidirectional character
type of a Unicode character as specified by the
<ulink url="http://www.unicode.org/reports/tr9/">Unicode bidirectional algorithm</ulink>.

C - `PangoBidiType`

### PANGO_BIDI_TYPE_L = 0
Left-to-Right

### PANGO_BIDI_TYPE_LRE = 1
Left-to-Right Embedding

### PANGO_BIDI_TYPE_LRO = 2
Left-to-Right Override

### PANGO_BIDI_TYPE_R = 3
Right-to-Left

### PANGO_BIDI_TYPE_AL = 4
Right-to-Left Arabic

### PANGO_BIDI_TYPE_RLE = 5
Right-to-Left Embedding

### PANGO_BIDI_TYPE_RLO = 6
Right-to-Left Override

### PANGO_BIDI_TYPE_PDF = 7
Pop Directional Format

### PANGO_BIDI_TYPE_EN = 8
European Number

### PANGO_BIDI_TYPE_ES = 9
European Number Separator

### PANGO_BIDI_TYPE_ET = 10
European Number Terminator

### PANGO_BIDI_TYPE_AN = 11
Arabic Number

### PANGO_BIDI_TYPE_CS = 12
Common Number Separator

### PANGO_BIDI_TYPE_NSM = 13
Nonspacing Mark

### PANGO_BIDI_TYPE_BN = 14
Boundary Neutral

### PANGO_BIDI_TYPE_B = 15
Paragraph Separator

### PANGO_BIDI_TYPE_S = 16
Segment Separator

### PANGO_BIDI_TYPE_WS = 17
Whitespace

### PANGO_BIDI_TYPE_ON = 18
Other Neutrals


## `CoverageLevel`

Used to indicate how well a font can represent a particular Unicode
character point for a particular script.

C - `PangoCoverageLevel`

### PANGO_COVERAGE_NONE = 0
The character is not representable with the font.

### PANGO_COVERAGE_FALLBACK = 1
The character is represented in a way that may be
comprehensible but is not the correct graphical form.
For instance, a Hangul character represented as a
a sequence of Jamos, or a Latin transliteration of a Cyrillic word.

### PANGO_COVERAGE_APPROXIMATE = 2
The character is represented as basically the correct
graphical form, but with a stylistic variant inappropriate for
the current script.

### PANGO_COVERAGE_EXACT = 3
The character is represented as the correct graphical form.


## `Direction`

The #PangoDirection type represents a direction in the
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
vertical text is handled in Pango.

C - `PangoDirection`

### PANGO_DIRECTION_LTR = 0
A strong left-to-right direction

### PANGO_DIRECTION_RTL = 1
A strong right-to-left direction

### PANGO_DIRECTION_TTB_LTR = 2
Deprecated value; treated the
  same as %PANGO_DIRECTION_RTL.

### PANGO_DIRECTION_TTB_RTL = 3
Deprecated value; treated the
  same as %PANGO_DIRECTION_LTR

### PANGO_DIRECTION_WEAK_LTR = 4
A weak left-to-right direction

### PANGO_DIRECTION_WEAK_RTL = 5
A weak right-to-left direction

### PANGO_DIRECTION_NEUTRAL = 6
No direction specified


## `EllipsizeMode`

The #PangoEllipsizeMode type describes what sort of (if any)
ellipsization should be applied to a line of text. In
the ellipsization process characters are removed from the
text in order to make it fit to a given width and replaced
with an ellipsis.

C - `PangoEllipsizeMode`

### PANGO_ELLIPSIZE_NONE = 0
No ellipsization

### PANGO_ELLIPSIZE_START = 1
Omit characters at the start of the text

### PANGO_ELLIPSIZE_MIDDLE = 2
Omit characters in the middle of the text

### PANGO_ELLIPSIZE_END = 3
Omit characters at the end of the text


## `Gravity`

The #PangoGravity type represents the orientation of glyphs in a segment
of text.  This is useful when rendering vertical text layouts.  In
those situations, the layout is rotated using a non-identity PangoMatrix,
and then glyph orientation is controlled using #PangoGravity.
Not every value in this enumeration makes sense for every usage of
&num;PangoGravity; for example, %PANGO_GRAVITY_AUTO only can be passed to
pango_context_set_base_gravity() and can only be returned by
pango_context_get_base_gravity().

See also: #PangoGravityHint

C - `PangoGravity`

### PANGO_GRAVITY_SOUTH = 0
Glyphs stand upright (default)

### PANGO_GRAVITY_EAST = 1
Glyphs are rotated 90 degrees clockwise

### PANGO_GRAVITY_NORTH = 2
Glyphs are upside-down

### PANGO_GRAVITY_WEST = 3
Glyphs are rotated 90 degrees counter-clockwise

### PANGO_GRAVITY_AUTO = 4
Gravity is resolved from the context matrix


## `GravityHint`

The #PangoGravityHint defines how horizontal scripts should behave in a
vertical context.  That is, English excerpt in a vertical paragraph for
example.

See #PangoGravity.

C - `PangoGravityHint`

### PANGO_GRAVITY_HINT_NATURAL = 0
scripts will take their natural gravity based
on the base gravity and the script.  This is the default.

### PANGO_GRAVITY_HINT_STRONG = 1
always use the base gravity set, regardless of
the script.

### PANGO_GRAVITY_HINT_LINE = 2
for scripts not in their natural direction (eg.
Latin in East gravity), choose per-script gravity such that every script
respects the line progression.  This means, Latin and Arabic will take
opposite gravities and both flow top-to-bottom for example.


## `RenderPart`

#PangoRenderPart defines different items to render for such
purposes as setting colors.

C - `PangoRenderPart`

### PANGO_RENDER_PART_FOREGROUND = 0
the text itself

### PANGO_RENDER_PART_BACKGROUND = 1
the area behind the text

### PANGO_RENDER_PART_UNDERLINE = 2
underlines

### PANGO_RENDER_PART_STRIKETHROUGH = 3
strikethrough lines


## `Script`

The #PangoScript enumeration identifies different writing
systems. The values correspond to the names as defined in the
Unicode standard.
Note that new types may be added in the future. Applications should be ready
to handle unknown values.  This enumeration is interchangeable with
&num;GUnicodeScript.  See <ulink
url="http://www.unicode.org/reports/tr24/">Unicode Standard Annex
&num;24: Script names</ulink>.

C - `PangoScript`

### PANGO_SCRIPT_INVALID_CODE = -1
a value never returned from pango_script_for_unichar()

### PANGO_SCRIPT_COMMON = 0
a character used by multiple different scripts

### PANGO_SCRIPT_INHERITED = 1
a mark glyph that takes its script from the
base glyph to which it is attached

### PANGO_SCRIPT_ARABIC = 2
Arabic

### PANGO_SCRIPT_ARMENIAN = 3
Armenian

### PANGO_SCRIPT_BENGALI = 4
Bengali

### PANGO_SCRIPT_BOPOMOFO = 5
Bopomofo

### PANGO_SCRIPT_CHEROKEE = 6
Cherokee

### PANGO_SCRIPT_COPTIC = 7
Coptic

### PANGO_SCRIPT_CYRILLIC = 8
Cyrillic

### PANGO_SCRIPT_DESERET = 9
Deseret

### PANGO_SCRIPT_DEVANAGARI = 10
Devanagari

### PANGO_SCRIPT_ETHIOPIC = 11
Ethiopic

### PANGO_SCRIPT_GEORGIAN = 12
Georgian

### PANGO_SCRIPT_GOTHIC = 13
Gothic

### PANGO_SCRIPT_GREEK = 14
Greek

### PANGO_SCRIPT_GUJARATI = 15
Gujarati

### PANGO_SCRIPT_GURMUKHI = 16
Gurmukhi

### PANGO_SCRIPT_HAN = 17
Han

### PANGO_SCRIPT_HANGUL = 18
Hangul

### PANGO_SCRIPT_HEBREW = 19
Hebrew

### PANGO_SCRIPT_HIRAGANA = 20
Hiragana

### PANGO_SCRIPT_KANNADA = 21
Kannada

### PANGO_SCRIPT_KATAKANA = 22
Katakana

### PANGO_SCRIPT_KHMER = 23
Khmer

### PANGO_SCRIPT_LAO = 24
Lao

### PANGO_SCRIPT_LATIN = 25
Latin

### PANGO_SCRIPT_MALAYALAM = 26
Malayalam

### PANGO_SCRIPT_MONGOLIAN = 27
Mongolian

### PANGO_SCRIPT_MYANMAR = 28
Myanmar

### PANGO_SCRIPT_OGHAM = 29
Ogham

### PANGO_SCRIPT_OLD_ITALIC = 30
Old Italic

### PANGO_SCRIPT_ORIYA = 31
Oriya

### PANGO_SCRIPT_RUNIC = 32
Runic

### PANGO_SCRIPT_SINHALA = 33
Sinhala

### PANGO_SCRIPT_SYRIAC = 34
Syriac

### PANGO_SCRIPT_TAMIL = 35
Tamil

### PANGO_SCRIPT_TELUGU = 36
Telugu

### PANGO_SCRIPT_THAANA = 37
Thaana

### PANGO_SCRIPT_THAI = 38
Thai

### PANGO_SCRIPT_TIBETAN = 39
Tibetan

### PANGO_SCRIPT_CANADIAN_ABORIGINAL = 40
Canadian Aboriginal

### PANGO_SCRIPT_YI = 41
Yi

### PANGO_SCRIPT_TAGALOG = 42
Tagalog

### PANGO_SCRIPT_HANUNOO = 43
Hanunoo

### PANGO_SCRIPT_BUHID = 44
Buhid

### PANGO_SCRIPT_TAGBANWA = 45
Tagbanwa

### PANGO_SCRIPT_BRAILLE = 46
Braille

### PANGO_SCRIPT_CYPRIOT = 47
Cypriot

### PANGO_SCRIPT_LIMBU = 48
Limbu

### PANGO_SCRIPT_OSMANYA = 49
Osmanya

### PANGO_SCRIPT_SHAVIAN = 50
Shavian

### PANGO_SCRIPT_LINEAR_B = 51
Linear B

### PANGO_SCRIPT_TAI_LE = 52
Tai Le

### PANGO_SCRIPT_UGARITIC = 53
Ugaritic

### PANGO_SCRIPT_NEW_TAI_LUE = 54
New Tai Lue. Since 1.10

### PANGO_SCRIPT_BUGINESE = 55
Buginese. Since 1.10

### PANGO_SCRIPT_GLAGOLITIC = 56
Glagolitic. Since 1.10

### PANGO_SCRIPT_TIFINAGH = 57
Tifinagh. Since 1.10

### PANGO_SCRIPT_SYLOTI_NAGRI = 58
Syloti Nagri. Since 1.10

### PANGO_SCRIPT_OLD_PERSIAN = 59
Old Persian. Since 1.10

### PANGO_SCRIPT_KHAROSHTHI = 60
Kharoshthi. Since 1.10

### PANGO_SCRIPT_UNKNOWN = 61
an unassigned code point. Since 1.14

### PANGO_SCRIPT_BALINESE = 62
Balinese. Since 1.14

### PANGO_SCRIPT_CUNEIFORM = 63
Cuneiform. Since 1.14

### PANGO_SCRIPT_PHOENICIAN = 64
Phoenician. Since 1.14

### PANGO_SCRIPT_PHAGS_PA = 65
Phags-pa. Since 1.14

### PANGO_SCRIPT_NKO = 66
N'Ko. Since 1.14

### PANGO_SCRIPT_KAYAH_LI = 67
Kayah Li. Since 1.20.1

### PANGO_SCRIPT_LEPCHA = 68
Lepcha. Since 1.20.1

### PANGO_SCRIPT_REJANG = 69
Rejang. Since 1.20.1

### PANGO_SCRIPT_SUNDANESE = 70
Sundanese. Since 1.20.1

### PANGO_SCRIPT_SAURASHTRA = 71
Saurashtra. Since 1.20.1

### PANGO_SCRIPT_CHAM = 72
Cham. Since 1.20.1

### PANGO_SCRIPT_OL_CHIKI = 73
Ol Chiki. Since 1.20.1

### PANGO_SCRIPT_VAI = 74
Vai. Since 1.20.1

### PANGO_SCRIPT_CARIAN = 75
Carian. Since 1.20.1

### PANGO_SCRIPT_LYCIAN = 76
Lycian. Since 1.20.1

### PANGO_SCRIPT_LYDIAN = 77
Lydian. Since 1.20.1

### PANGO_SCRIPT_BATAK = 78
Batak. Since 1.32

### PANGO_SCRIPT_BRAHMI = 79
Brahmi. Since 1.32

### PANGO_SCRIPT_MANDAIC = 80
Mandaic. Since 1.32

### PANGO_SCRIPT_CHAKMA = 81
Chakma. Since: 1.32

### PANGO_SCRIPT_MEROITIC_CURSIVE = 82
Meroitic Cursive. Since: 1.32

### PANGO_SCRIPT_MEROITIC_HIEROGLYPHS = 83
Meroitic Hieroglyphs. Since: 1.32

### PANGO_SCRIPT_MIAO = 84
Miao. Since: 1.32

### PANGO_SCRIPT_SHARADA = 85
Sharada. Since: 1.32

### PANGO_SCRIPT_SORA_SOMPENG = 86
Sora Sompeng. Since: 1.32

### PANGO_SCRIPT_TAKRI = 87
Takri. Since: 1.32

### PANGO_SCRIPT_BASSA_VAH = 88
Bassa. Since: 1.40

### PANGO_SCRIPT_CAUCASIAN_ALBANIAN = 89
Caucasian Albanian. Since: 1.40

### PANGO_SCRIPT_DUPLOYAN = 90
Duployan. Since: 1.40

### PANGO_SCRIPT_ELBASAN = 91
Elbasan. Since: 1.40

### PANGO_SCRIPT_GRANTHA = 92
Grantha. Since: 1.40

### PANGO_SCRIPT_KHOJKI = 93
Kjohki. Since: 1.40

### PANGO_SCRIPT_KHUDAWADI = 94
Khudawadi, Sindhi. Since: 1.40

### PANGO_SCRIPT_LINEAR_A = 95
Linear A. Since: 1.40

### PANGO_SCRIPT_MAHAJANI = 96
Mahajani. Since: 1.40

### PANGO_SCRIPT_MANICHAEAN = 97
Manichaean. Since: 1.40

### PANGO_SCRIPT_MENDE_KIKAKUI = 98
Mende Kikakui. Since: 1.40

### PANGO_SCRIPT_MODI = 99
Modi. Since: 1.40

### PANGO_SCRIPT_MRO = 100
Mro. Since: 1.40

### PANGO_SCRIPT_NABATAEAN = 101
Nabataean. Since: 1.40

### PANGO_SCRIPT_OLD_NORTH_ARABIAN = 102
Old North Arabian. Since: 1.40

### PANGO_SCRIPT_OLD_PERMIC = 103
Old Permic. Since: 1.40

### PANGO_SCRIPT_PAHAWH_HMONG = 104
Pahawh Hmong. Since: 1.40

### PANGO_SCRIPT_PALMYRENE = 105
Palmyrene. Since: 1.40

### PANGO_SCRIPT_PAU_CIN_HAU = 106
Pau Cin Hau. Since: 1.40

### PANGO_SCRIPT_PSALTER_PAHLAVI = 107
Psalter Pahlavi. Since: 1.40

### PANGO_SCRIPT_SIDDHAM = 108
Siddham. Since: 1.40

### PANGO_SCRIPT_TIRHUTA = 109
Tirhuta. Since: 1.40

### PANGO_SCRIPT_WARANG_CITI = 110
Warang Citi. Since: 1.40

### PANGO_SCRIPT_AHOM = 111
Ahom. Since: 1.40

### PANGO_SCRIPT_ANATOLIAN_HIEROGLYPHS = 112
Anatolian Hieroglyphs. Since: 1.40

### PANGO_SCRIPT_HATRAN = 113
Hatran. Since: 1.40

### PANGO_SCRIPT_MULTANI = 114
Multani. Since: 1.40

### PANGO_SCRIPT_OLD_HUNGARIAN = 115
Old Hungarian. Since: 1.40

### PANGO_SCRIPT_SIGNWRITING = 116
Signwriting. Since: 1.40


## `Stretch`

An enumeration specifying the width of the font relative to other designs
within a family.

C - `PangoStretch`

### PANGO_STRETCH_ULTRA_CONDENSED = 0
ultra condensed width

### PANGO_STRETCH_EXTRA_CONDENSED = 1
extra condensed width

### PANGO_STRETCH_CONDENSED = 2
condensed width

### PANGO_STRETCH_SEMI_CONDENSED = 3
semi condensed width

### PANGO_STRETCH_NORMAL = 4
the normal width

### PANGO_STRETCH_SEMI_EXPANDED = 5
semi expanded width

### PANGO_STRETCH_EXPANDED = 6
expanded width

### PANGO_STRETCH_EXTRA_EXPANDED = 7
extra expanded width

### PANGO_STRETCH_ULTRA_EXPANDED = 8
ultra expanded width


## `Style`

An enumeration specifying the various slant styles possible for a font.

C - `PangoStyle`

### PANGO_STYLE_NORMAL = 0
the font is upright.

### PANGO_STYLE_OBLIQUE = 1
the font is slanted, but in a roman style.

### PANGO_STYLE_ITALIC = 2
the font is slanted in an italic style.


## `TabAlign`

A #PangoTabAlign specifies where a tab stop appears relative to the text.

C - `PangoTabAlign`

### PANGO_TAB_LEFT = 0
the tab stop appears to the left of the text.


## `Underline`

The #PangoUnderline enumeration is used to specify
whether text should be underlined, and if so, the type
of underlining.

C - `PangoUnderline`

### PANGO_UNDERLINE_NONE = 0
no underline should be drawn

### PANGO_UNDERLINE_SINGLE = 1
a single underline should be drawn

### PANGO_UNDERLINE_DOUBLE = 2
a double underline should be drawn

### PANGO_UNDERLINE_LOW = 3
a single underline should be drawn at a position
beneath the ink extents of the text being
underlined. This should be used only for underlining
single characters, such as for keyboard
accelerators. %PANGO_UNDERLINE_SINGLE should
be used for extended portions of text.

### PANGO_UNDERLINE_ERROR = 4
a wavy underline should be drawn below.
This underline is typically used to indicate
an error such as a possilble mispelling; in some
cases a contrasting color may automatically
be used. This type of underlining is available since Pango 1.4.


## `Variant`

An enumeration specifying capitalization variant of the font.

C - `PangoVariant`

### PANGO_VARIANT_NORMAL = 0
A normal font.

### PANGO_VARIANT_SMALL_CAPS = 1
A font with the lower case characters
replaced by smaller variants of the capital characters.


## `Weight`

An enumeration specifying the weight (boldness) of a font. This is a numerical
value ranging from 100 to 1000, but there are some predefined values:

C - `PangoWeight`

### PANGO_WEIGHT_THIN = 100
the thin weight (= 100; Since: 1.24)

### PANGO_WEIGHT_ULTRALIGHT = 200
the ultralight weight (= 200)

### PANGO_WEIGHT_LIGHT = 300
the light weight (= 300)

### PANGO_WEIGHT_SEMILIGHT = 350
the semilight weight (= 350; Since: 1.36.7)

### PANGO_WEIGHT_BOOK = 380
the book weight (= 380; Since: 1.24)

### PANGO_WEIGHT_NORMAL = 400
the default weight (= 400)

### PANGO_WEIGHT_MEDIUM = 500
the normal weight (= 500; Since: 1.24)

### PANGO_WEIGHT_SEMIBOLD = 600
the semibold weight (= 600)

### PANGO_WEIGHT_BOLD = 700
the bold weight (= 700)

### PANGO_WEIGHT_ULTRABOLD = 800
the ultrabold weight (= 800)

### PANGO_WEIGHT_HEAVY = 900
the heavy weight (= 900)

### PANGO_WEIGHT_ULTRAHEAVY = 1000
the ultraheavy weight (= 1000; Since: 1.24)


## `WrapMode`

A #PangoWrapMode describes how to wrap the lines of a #PangoLayout to the desired width.

C - `PangoWrapMode`

### PANGO_WRAP_WORD = 0
wrap lines at word boundaries.

### PANGO_WRAP_CHAR = 1
wrap lines at character boundaries.

### PANGO_WRAP_WORD_CHAR = 2
wrap lines at word boundaries, but fall back to character boundaries if there is not
enough space for a full word.


