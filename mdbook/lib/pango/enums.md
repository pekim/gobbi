# `pango` enums

## `Alignment`

A #PangoAlignment describes how to align the lines of a #PangoLayout within the
available space. If the #PangoLayout is set to justify
using pango_layout_set_justify(), this only has effect for partial lines.

- PANGO_ALIGN_LEFT = 0
- PANGO_ALIGN_CENTER = 1
- PANGO_ALIGN_RIGHT = 2

C - `PangoAlignment`

## `AttrType`

The #PangoAttrType
distinguishes between different types of attributes. Along with the
predefined values, it is possible to allocate additional values
for custom attributes using pango_attr_type_register(). The predefined
values are given below. The type of structure used to store the
attribute is listed in parentheses after the description.

- PANGO_ATTR_INVALID = 0
- PANGO_ATTR_LANGUAGE = 1
- PANGO_ATTR_FAMILY = 2
- PANGO_ATTR_STYLE = 3
- PANGO_ATTR_WEIGHT = 4
- PANGO_ATTR_VARIANT = 5
- PANGO_ATTR_STRETCH = 6
- PANGO_ATTR_SIZE = 7
- PANGO_ATTR_FONT_DESC = 8
- PANGO_ATTR_FOREGROUND = 9
- PANGO_ATTR_BACKGROUND = 10
- PANGO_ATTR_UNDERLINE = 11
- PANGO_ATTR_STRIKETHROUGH = 12
- PANGO_ATTR_RISE = 13
- PANGO_ATTR_SHAPE = 14
- PANGO_ATTR_SCALE = 15
- PANGO_ATTR_FALLBACK = 16
- PANGO_ATTR_LETTER_SPACING = 17
- PANGO_ATTR_UNDERLINE_COLOR = 18
- PANGO_ATTR_STRIKETHROUGH_COLOR = 19
- PANGO_ATTR_ABSOLUTE_SIZE = 20
- PANGO_ATTR_GRAVITY = 21
- PANGO_ATTR_GRAVITY_HINT = 22
- PANGO_ATTR_FONT_FEATURES = 23
- PANGO_ATTR_FOREGROUND_ALPHA = 24
- PANGO_ATTR_BACKGROUND_ALPHA = 25

C - `PangoAttrType`

## `BidiType`

The #PangoBidiType type represents the bidirectional character
type of a Unicode character as specified by the
<ulink url="http://www.unicode.org/reports/tr9/">Unicode bidirectional algorithm</ulink>.

- PANGO_BIDI_TYPE_L = 0
- PANGO_BIDI_TYPE_LRE = 1
- PANGO_BIDI_TYPE_LRO = 2
- PANGO_BIDI_TYPE_R = 3
- PANGO_BIDI_TYPE_AL = 4
- PANGO_BIDI_TYPE_RLE = 5
- PANGO_BIDI_TYPE_RLO = 6
- PANGO_BIDI_TYPE_PDF = 7
- PANGO_BIDI_TYPE_EN = 8
- PANGO_BIDI_TYPE_ES = 9
- PANGO_BIDI_TYPE_ET = 10
- PANGO_BIDI_TYPE_AN = 11
- PANGO_BIDI_TYPE_CS = 12
- PANGO_BIDI_TYPE_NSM = 13
- PANGO_BIDI_TYPE_BN = 14
- PANGO_BIDI_TYPE_B = 15
- PANGO_BIDI_TYPE_S = 16
- PANGO_BIDI_TYPE_WS = 17
- PANGO_BIDI_TYPE_ON = 18

C - `PangoBidiType`

## `CoverageLevel`

Used to indicate how well a font can represent a particular Unicode
character point for a particular script.

- PANGO_COVERAGE_NONE = 0
- PANGO_COVERAGE_FALLBACK = 1
- PANGO_COVERAGE_APPROXIMATE = 2
- PANGO_COVERAGE_EXACT = 3

C - `PangoCoverageLevel`

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

- PANGO_DIRECTION_LTR = 0
- PANGO_DIRECTION_RTL = 1
- PANGO_DIRECTION_TTB_LTR = 2
- PANGO_DIRECTION_TTB_RTL = 3
- PANGO_DIRECTION_WEAK_LTR = 4
- PANGO_DIRECTION_WEAK_RTL = 5
- PANGO_DIRECTION_NEUTRAL = 6

C - `PangoDirection`

## `EllipsizeMode`

The #PangoEllipsizeMode type describes what sort of (if any)
ellipsization should be applied to a line of text. In
the ellipsization process characters are removed from the
text in order to make it fit to a given width and replaced
with an ellipsis.

- PANGO_ELLIPSIZE_NONE = 0
- PANGO_ELLIPSIZE_START = 1
- PANGO_ELLIPSIZE_MIDDLE = 2
- PANGO_ELLIPSIZE_END = 3

C - `PangoEllipsizeMode`

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

- PANGO_GRAVITY_SOUTH = 0
- PANGO_GRAVITY_EAST = 1
- PANGO_GRAVITY_NORTH = 2
- PANGO_GRAVITY_WEST = 3
- PANGO_GRAVITY_AUTO = 4

C - `PangoGravity`

## `GravityHint`

The #PangoGravityHint defines how horizontal scripts should behave in a
vertical context.  That is, English excerpt in a vertical paragraph for
example.

See #PangoGravity.

- PANGO_GRAVITY_HINT_NATURAL = 0
- PANGO_GRAVITY_HINT_STRONG = 1
- PANGO_GRAVITY_HINT_LINE = 2

C - `PangoGravityHint`

## `RenderPart`

#PangoRenderPart defines different items to render for such
purposes as setting colors.

- PANGO_RENDER_PART_FOREGROUND = 0
- PANGO_RENDER_PART_BACKGROUND = 1
- PANGO_RENDER_PART_UNDERLINE = 2
- PANGO_RENDER_PART_STRIKETHROUGH = 3

C - `PangoRenderPart`

## `Script`

The #PangoScript enumeration identifies different writing
systems. The values correspond to the names as defined in the
Unicode standard.
Note that new types may be added in the future. Applications should be ready
to handle unknown values.  This enumeration is interchangeable with
&num;GUnicodeScript.  See <ulink
url="http://www.unicode.org/reports/tr24/">Unicode Standard Annex
&num;24: Script names</ulink>.

- PANGO_SCRIPT_INVALID_CODE = -1
- PANGO_SCRIPT_COMMON = 0
- PANGO_SCRIPT_INHERITED = 1
- PANGO_SCRIPT_ARABIC = 2
- PANGO_SCRIPT_ARMENIAN = 3
- PANGO_SCRIPT_BENGALI = 4
- PANGO_SCRIPT_BOPOMOFO = 5
- PANGO_SCRIPT_CHEROKEE = 6
- PANGO_SCRIPT_COPTIC = 7
- PANGO_SCRIPT_CYRILLIC = 8
- PANGO_SCRIPT_DESERET = 9
- PANGO_SCRIPT_DEVANAGARI = 10
- PANGO_SCRIPT_ETHIOPIC = 11
- PANGO_SCRIPT_GEORGIAN = 12
- PANGO_SCRIPT_GOTHIC = 13
- PANGO_SCRIPT_GREEK = 14
- PANGO_SCRIPT_GUJARATI = 15
- PANGO_SCRIPT_GURMUKHI = 16
- PANGO_SCRIPT_HAN = 17
- PANGO_SCRIPT_HANGUL = 18
- PANGO_SCRIPT_HEBREW = 19
- PANGO_SCRIPT_HIRAGANA = 20
- PANGO_SCRIPT_KANNADA = 21
- PANGO_SCRIPT_KATAKANA = 22
- PANGO_SCRIPT_KHMER = 23
- PANGO_SCRIPT_LAO = 24
- PANGO_SCRIPT_LATIN = 25
- PANGO_SCRIPT_MALAYALAM = 26
- PANGO_SCRIPT_MONGOLIAN = 27
- PANGO_SCRIPT_MYANMAR = 28
- PANGO_SCRIPT_OGHAM = 29
- PANGO_SCRIPT_OLD_ITALIC = 30
- PANGO_SCRIPT_ORIYA = 31
- PANGO_SCRIPT_RUNIC = 32
- PANGO_SCRIPT_SINHALA = 33
- PANGO_SCRIPT_SYRIAC = 34
- PANGO_SCRIPT_TAMIL = 35
- PANGO_SCRIPT_TELUGU = 36
- PANGO_SCRIPT_THAANA = 37
- PANGO_SCRIPT_THAI = 38
- PANGO_SCRIPT_TIBETAN = 39
- PANGO_SCRIPT_CANADIAN_ABORIGINAL = 40
- PANGO_SCRIPT_YI = 41
- PANGO_SCRIPT_TAGALOG = 42
- PANGO_SCRIPT_HANUNOO = 43
- PANGO_SCRIPT_BUHID = 44
- PANGO_SCRIPT_TAGBANWA = 45
- PANGO_SCRIPT_BRAILLE = 46
- PANGO_SCRIPT_CYPRIOT = 47
- PANGO_SCRIPT_LIMBU = 48
- PANGO_SCRIPT_OSMANYA = 49
- PANGO_SCRIPT_SHAVIAN = 50
- PANGO_SCRIPT_LINEAR_B = 51
- PANGO_SCRIPT_TAI_LE = 52
- PANGO_SCRIPT_UGARITIC = 53
- PANGO_SCRIPT_NEW_TAI_LUE = 54
- PANGO_SCRIPT_BUGINESE = 55
- PANGO_SCRIPT_GLAGOLITIC = 56
- PANGO_SCRIPT_TIFINAGH = 57
- PANGO_SCRIPT_SYLOTI_NAGRI = 58
- PANGO_SCRIPT_OLD_PERSIAN = 59
- PANGO_SCRIPT_KHAROSHTHI = 60
- PANGO_SCRIPT_UNKNOWN = 61
- PANGO_SCRIPT_BALINESE = 62
- PANGO_SCRIPT_CUNEIFORM = 63
- PANGO_SCRIPT_PHOENICIAN = 64
- PANGO_SCRIPT_PHAGS_PA = 65
- PANGO_SCRIPT_NKO = 66
- PANGO_SCRIPT_KAYAH_LI = 67
- PANGO_SCRIPT_LEPCHA = 68
- PANGO_SCRIPT_REJANG = 69
- PANGO_SCRIPT_SUNDANESE = 70
- PANGO_SCRIPT_SAURASHTRA = 71
- PANGO_SCRIPT_CHAM = 72
- PANGO_SCRIPT_OL_CHIKI = 73
- PANGO_SCRIPT_VAI = 74
- PANGO_SCRIPT_CARIAN = 75
- PANGO_SCRIPT_LYCIAN = 76
- PANGO_SCRIPT_LYDIAN = 77
- PANGO_SCRIPT_BATAK = 78
- PANGO_SCRIPT_BRAHMI = 79
- PANGO_SCRIPT_MANDAIC = 80
- PANGO_SCRIPT_CHAKMA = 81
- PANGO_SCRIPT_MEROITIC_CURSIVE = 82
- PANGO_SCRIPT_MEROITIC_HIEROGLYPHS = 83
- PANGO_SCRIPT_MIAO = 84
- PANGO_SCRIPT_SHARADA = 85
- PANGO_SCRIPT_SORA_SOMPENG = 86
- PANGO_SCRIPT_TAKRI = 87
- PANGO_SCRIPT_BASSA_VAH = 88
- PANGO_SCRIPT_CAUCASIAN_ALBANIAN = 89
- PANGO_SCRIPT_DUPLOYAN = 90
- PANGO_SCRIPT_ELBASAN = 91
- PANGO_SCRIPT_GRANTHA = 92
- PANGO_SCRIPT_KHOJKI = 93
- PANGO_SCRIPT_KHUDAWADI = 94
- PANGO_SCRIPT_LINEAR_A = 95
- PANGO_SCRIPT_MAHAJANI = 96
- PANGO_SCRIPT_MANICHAEAN = 97
- PANGO_SCRIPT_MENDE_KIKAKUI = 98
- PANGO_SCRIPT_MODI = 99
- PANGO_SCRIPT_MRO = 100
- PANGO_SCRIPT_NABATAEAN = 101
- PANGO_SCRIPT_OLD_NORTH_ARABIAN = 102
- PANGO_SCRIPT_OLD_PERMIC = 103
- PANGO_SCRIPT_PAHAWH_HMONG = 104
- PANGO_SCRIPT_PALMYRENE = 105
- PANGO_SCRIPT_PAU_CIN_HAU = 106
- PANGO_SCRIPT_PSALTER_PAHLAVI = 107
- PANGO_SCRIPT_SIDDHAM = 108
- PANGO_SCRIPT_TIRHUTA = 109
- PANGO_SCRIPT_WARANG_CITI = 110
- PANGO_SCRIPT_AHOM = 111
- PANGO_SCRIPT_ANATOLIAN_HIEROGLYPHS = 112
- PANGO_SCRIPT_HATRAN = 113
- PANGO_SCRIPT_MULTANI = 114
- PANGO_SCRIPT_OLD_HUNGARIAN = 115
- PANGO_SCRIPT_SIGNWRITING = 116

C - `PangoScript`

## `Stretch`

An enumeration specifying the width of the font relative to other designs
within a family.

- PANGO_STRETCH_ULTRA_CONDENSED = 0
- PANGO_STRETCH_EXTRA_CONDENSED = 1
- PANGO_STRETCH_CONDENSED = 2
- PANGO_STRETCH_SEMI_CONDENSED = 3
- PANGO_STRETCH_NORMAL = 4
- PANGO_STRETCH_SEMI_EXPANDED = 5
- PANGO_STRETCH_EXPANDED = 6
- PANGO_STRETCH_EXTRA_EXPANDED = 7
- PANGO_STRETCH_ULTRA_EXPANDED = 8

C - `PangoStretch`

## `Style`

An enumeration specifying the various slant styles possible for a font.

- PANGO_STYLE_NORMAL = 0
- PANGO_STYLE_OBLIQUE = 1
- PANGO_STYLE_ITALIC = 2

C - `PangoStyle`

## `TabAlign`

A #PangoTabAlign specifies where a tab stop appears relative to the text.

- PANGO_TAB_LEFT = 0

C - `PangoTabAlign`

## `Underline`

The #PangoUnderline enumeration is used to specify
whether text should be underlined, and if so, the type
of underlining.

- PANGO_UNDERLINE_NONE = 0
- PANGO_UNDERLINE_SINGLE = 1
- PANGO_UNDERLINE_DOUBLE = 2
- PANGO_UNDERLINE_LOW = 3
- PANGO_UNDERLINE_ERROR = 4

C - `PangoUnderline`

## `Variant`

An enumeration specifying capitalization variant of the font.

- PANGO_VARIANT_NORMAL = 0
- PANGO_VARIANT_SMALL_CAPS = 1

C - `PangoVariant`

## `Weight`

An enumeration specifying the weight (boldness) of a font. This is a numerical
value ranging from 100 to 1000, but there are some predefined values:

- PANGO_WEIGHT_THIN = 100
- PANGO_WEIGHT_ULTRALIGHT = 200
- PANGO_WEIGHT_LIGHT = 300
- PANGO_WEIGHT_SEMILIGHT = 350
- PANGO_WEIGHT_BOOK = 380
- PANGO_WEIGHT_NORMAL = 400
- PANGO_WEIGHT_MEDIUM = 500
- PANGO_WEIGHT_SEMIBOLD = 600
- PANGO_WEIGHT_BOLD = 700
- PANGO_WEIGHT_ULTRABOLD = 800
- PANGO_WEIGHT_HEAVY = 900
- PANGO_WEIGHT_ULTRAHEAVY = 1000

C - `PangoWeight`

## `WrapMode`

A #PangoWrapMode describes how to wrap the lines of a #PangoLayout to the desired width.

- PANGO_WRAP_WORD = 0
- PANGO_WRAP_CHAR = 1
- PANGO_WRAP_WORD_CHAR = 2

C - `PangoWrapMode`

CONDENSED = %!s(int=1)
- PANGO_STRETCH_CONDENSED = %!s(int=2)
- PANGO_STRETCH_SEMI_CONDENSED = %!s(int=3)
- PANGO_STRETCH_NORMAL = %!s(int=4)
- PANGO_STRETCH_SEMI_EXPANDED = %!s(int=5)
- PANGO_STRETCH_EXPANDED = %!s(int=6)
- PANGO_STRETCH_EXTRA_EXPANDED = %!s(int=7)
- PANGO_STRETCH_ULTRA_EXPANDED = %!s(int=8)
## `Style`

An enumeration specifying the various slant styles possible for a font.

C - `PangoStyle`

- PANGO_STYLE_NORMAL = %!s(int=0)
- PANGO_STYLE_OBLIQUE = %!s(int=1)
- PANGO_STYLE_ITALIC = %!s(int=2)
## `TabAlign`

A #PangoTabAlign specifies where a tab stop appears relative to the text.

C - `PangoTabAlign`

- PANGO_TAB_LEFT = %!s(int=0)
## `Underline`

The #PangoUnderline enumeration is used to specify
whether text should be underlined, and if so, the type
of underlining.

C - `PangoUnderline`

- PANGO_UNDERLINE_NONE = %!s(int=0)
- PANGO_UNDERLINE_SINGLE = %!s(int=1)
- PANGO_UNDERLINE_DOUBLE = %!s(int=2)
- PANGO_UNDERLINE_LOW = %!s(int=3)
- PANGO_UNDERLINE_ERROR = %!s(int=4)
## `Variant`

An enumeration specifying capitalization variant of the font.

C - `PangoVariant`

- PANGO_VARIANT_NORMAL = %!s(int=0)
- PANGO_VARIANT_SMALL_CAPS = %!s(int=1)
## `Weight`

An enumeration specifying the weight (boldness) of a font. This is a numerical
value ranging from 100 to 1000, but there are some predefined values:

C - `PangoWeight`

- PANGO_WEIGHT_THIN = %!s(int=100)
- PANGO_WEIGHT_ULTRALIGHT = %!s(int=200)
- PANGO_WEIGHT_LIGHT = %!s(int=300)
- PANGO_WEIGHT_SEMILIGHT = %!s(int=350)
- PANGO_WEIGHT_BOOK = %!s(int=380)
- PANGO_WEIGHT_NORMAL = %!s(int=400)
- PANGO_WEIGHT_MEDIUM = %!s(int=500)
- PANGO_WEIGHT_SEMIBOLD = %!s(int=600)
- PANGO_WEIGHT_BOLD = %!s(int=700)
- PANGO_WEIGHT_ULTRABOLD = %!s(int=800)
- PANGO_WEIGHT_HEAVY = %!s(int=900)
- PANGO_WEIGHT_ULTRAHEAVY = %!s(int=1000)
## `WrapMode`

A #PangoWrapMode describes how to wrap the lines of a #PangoLayout to the desired width.

C - `PangoWrapMode`

- PANGO_WRAP_WORD = %!s(int=0)
- PANGO_WRAP_CHAR = %!s(int=1)
- PANGO_WRAP_WORD_CHAR = %!s(int=2)
