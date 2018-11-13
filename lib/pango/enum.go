// This is a generated file - DO NOT EDIT

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// A #PangoAlignment describes how to align the lines of a #PangoLayout within the
// available space. If the #PangoLayout is set to justify
// using pango_layout_set_justify(), this only has effect for partial lines.
type Alignment C.PangoAlignment

const (
	// Put all available space on the right
	PANGO_ALIGN_LEFT Alignment = 0

	// Center the line within the available space
	PANGO_ALIGN_CENTER Alignment = 1

	// Put all available space on the left
	PANGO_ALIGN_RIGHT Alignment = 2
)

// The #PangoAttrType
// distinguishes between different types of attributes. Along with the
// predefined values, it is possible to allocate additional values
// for custom attributes using pango_attr_type_register(). The predefined
// values are given below. The type of structure used to store the
// attribute is listed in parentheses after the description.
type AttrType C.PangoAttrType

const (
	// does not happen
	PANGO_ATTR_INVALID AttrType = 0

	// language (#PangoAttrLanguage)
	PANGO_ATTR_LANGUAGE AttrType = 1

	// font family name list (#PangoAttrString)
	PANGO_ATTR_FAMILY AttrType = 2

	// font slant style (#PangoAttrInt)
	PANGO_ATTR_STYLE AttrType = 3

	// font weight (#PangoAttrInt)
	PANGO_ATTR_WEIGHT AttrType = 4

	// font variant (normal or small caps) (#PangoAttrInt)
	PANGO_ATTR_VARIANT AttrType = 5

	// font stretch (#PangoAttrInt)
	PANGO_ATTR_STRETCH AttrType = 6

	// font size in points scaled by %PANGO_SCALE (#PangoAttrInt)
	PANGO_ATTR_SIZE AttrType = 7

	// font description (#PangoAttrFontDesc)
	PANGO_ATTR_FONT_DESC AttrType = 8

	// foreground color (#PangoAttrColor)
	PANGO_ATTR_FOREGROUND AttrType = 9

	// background color (#PangoAttrColor)
	PANGO_ATTR_BACKGROUND AttrType = 10

	// whether the text has an underline (#PangoAttrInt)
	PANGO_ATTR_UNDERLINE AttrType = 11

	// whether the text is struck-through (#PangoAttrInt)
	PANGO_ATTR_STRIKETHROUGH AttrType = 12

	// baseline displacement (#PangoAttrInt)
	PANGO_ATTR_RISE AttrType = 13

	// shape (#PangoAttrShape)
	PANGO_ATTR_SHAPE AttrType = 14

	// font size scale factor (#PangoAttrFloat)
	PANGO_ATTR_SCALE AttrType = 15

	// whether fallback is enabled (#PangoAttrInt)
	PANGO_ATTR_FALLBACK AttrType = 16

	// letter spacing (#PangoAttrInt)
	PANGO_ATTR_LETTER_SPACING AttrType = 17

	// underline color (#PangoAttrColor)
	PANGO_ATTR_UNDERLINE_COLOR AttrType = 18

	// strikethrough color (#PangoAttrColor)
	PANGO_ATTR_STRIKETHROUGH_COLOR AttrType = 19

	// font size in pixels scaled by %PANGO_SCALE (#PangoAttrInt)
	PANGO_ATTR_ABSOLUTE_SIZE AttrType = 20

	// base text gravity (#PangoAttrInt)
	PANGO_ATTR_GRAVITY AttrType = 21

	// gravity hint (#PangoAttrInt)
	PANGO_ATTR_GRAVITY_HINT AttrType = 22

	// OpenType font features (#PangoAttrString). Since 1.38
	PANGO_ATTR_FONT_FEATURES AttrType = 23

	// foreground alpha (#PangoAttrInt). Since 1.38
	PANGO_ATTR_FOREGROUND_ALPHA AttrType = 24

	// background alpha (#PangoAttrInt). Since 1.38
	PANGO_ATTR_BACKGROUND_ALPHA AttrType = 25
)

// Used to indicate how well a font can represent a particular Unicode
// character point for a particular script.
type CoverageLevel C.PangoCoverageLevel

const (
	// The character is not representable with the font.
	PANGO_COVERAGE_NONE CoverageLevel = 0

	// The character is represented in a way that may be
	// comprehensible but is not the correct graphical form.
	// For instance, a Hangul character represented as a
	// a sequence of Jamos, or a Latin transliteration of a Cyrillic word.
	PANGO_COVERAGE_FALLBACK CoverageLevel = 1

	// The character is represented as basically the correct
	// graphical form, but with a stylistic variant inappropriate for
	// the current script.
	PANGO_COVERAGE_APPROXIMATE CoverageLevel = 2

	// The character is represented as the correct graphical form.
	PANGO_COVERAGE_EXACT CoverageLevel = 3
)

// The #PangoDirection type represents a direction in the
// Unicode bidirectional algorithm; not every value in this
// enumeration makes sense for every usage of #PangoDirection;
// for example, the return value of pango_unichar_direction()
// and pango_find_base_dir() cannot be %PANGO_DIRECTION_WEAK_LTR
// or %PANGO_DIRECTION_WEAK_RTL, since every character is either
// neutral or has a strong direction; on the other hand
// %PANGO_DIRECTION_NEUTRAL doesn't make sense to pass
// to pango_itemize_with_base_dir().
//
// The %PANGO_DIRECTION_TTB_LTR, %PANGO_DIRECTION_TTB_RTL
// values come from an earlier interpretation of this
// enumeration as the writing direction of a block of
// text and are no longer used; See #PangoGravity for how
// vertical text is handled in Pango.
type Direction C.PangoDirection

const (
	// A strong left-to-right direction
	PANGO_DIRECTION_LTR Direction = 0

	// A strong right-to-left direction
	PANGO_DIRECTION_RTL Direction = 1

	// Deprecated value; treated the
	// same as %PANGO_DIRECTION_RTL.
	PANGO_DIRECTION_TTB_LTR Direction = 2

	// Deprecated value; treated the
	// same as %PANGO_DIRECTION_LTR
	PANGO_DIRECTION_TTB_RTL Direction = 3

	// A weak left-to-right direction
	PANGO_DIRECTION_WEAK_LTR Direction = 4

	// A weak right-to-left direction
	PANGO_DIRECTION_WEAK_RTL Direction = 5

	// No direction specified
	PANGO_DIRECTION_NEUTRAL Direction = 6
)

// The #PangoEllipsizeMode type describes what sort of (if any)
// ellipsization should be applied to a line of text. In
// the ellipsization process characters are removed from the
// text in order to make it fit to a given width and replaced
// with an ellipsis.
type EllipsizeMode C.PangoEllipsizeMode

const (
	// No ellipsization
	PANGO_ELLIPSIZE_NONE EllipsizeMode = 0

	// Omit characters at the start of the text
	PANGO_ELLIPSIZE_START EllipsizeMode = 1

	// Omit characters in the middle of the text
	PANGO_ELLIPSIZE_MIDDLE EllipsizeMode = 2

	// Omit characters at the end of the text
	PANGO_ELLIPSIZE_END EllipsizeMode = 3
)

// The #PangoScript enumeration identifies different writing
// systems. The values correspond to the names as defined in the
// Unicode standard.
// Note that new types may be added in the future. Applications should be ready
// to handle unknown values.  This enumeration is interchangeable with
// #GUnicodeScript.  See <ulink
// url="http://www.unicode.org/reports/tr24/">Unicode Standard Annex
// #24: Script names</ulink>.
type Script C.PangoScript

const (
	// a value never returned from pango_script_for_unichar()
	PANGO_SCRIPT_INVALID_CODE Script = -1

	// a character used by multiple different scripts
	PANGO_SCRIPT_COMMON Script = 0

	// a mark glyph that takes its script from the
	// base glyph to which it is attached
	PANGO_SCRIPT_INHERITED Script = 1

	// Arabic
	PANGO_SCRIPT_ARABIC Script = 2

	// Armenian
	PANGO_SCRIPT_ARMENIAN Script = 3

	// Bengali
	PANGO_SCRIPT_BENGALI Script = 4

	// Bopomofo
	PANGO_SCRIPT_BOPOMOFO Script = 5

	// Cherokee
	PANGO_SCRIPT_CHEROKEE Script = 6

	// Coptic
	PANGO_SCRIPT_COPTIC Script = 7

	// Cyrillic
	PANGO_SCRIPT_CYRILLIC Script = 8

	// Deseret
	PANGO_SCRIPT_DESERET Script = 9

	// Devanagari
	PANGO_SCRIPT_DEVANAGARI Script = 10

	// Ethiopic
	PANGO_SCRIPT_ETHIOPIC Script = 11

	// Georgian
	PANGO_SCRIPT_GEORGIAN Script = 12

	// Gothic
	PANGO_SCRIPT_GOTHIC Script = 13

	// Greek
	PANGO_SCRIPT_GREEK Script = 14

	// Gujarati
	PANGO_SCRIPT_GUJARATI Script = 15

	// Gurmukhi
	PANGO_SCRIPT_GURMUKHI Script = 16

	// Han
	PANGO_SCRIPT_HAN Script = 17

	// Hangul
	PANGO_SCRIPT_HANGUL Script = 18

	// Hebrew
	PANGO_SCRIPT_HEBREW Script = 19

	// Hiragana
	PANGO_SCRIPT_HIRAGANA Script = 20

	// Kannada
	PANGO_SCRIPT_KANNADA Script = 21

	// Katakana
	PANGO_SCRIPT_KATAKANA Script = 22

	// Khmer
	PANGO_SCRIPT_KHMER Script = 23

	// Lao
	PANGO_SCRIPT_LAO Script = 24

	// Latin
	PANGO_SCRIPT_LATIN Script = 25

	// Malayalam
	PANGO_SCRIPT_MALAYALAM Script = 26

	// Mongolian
	PANGO_SCRIPT_MONGOLIAN Script = 27

	// Myanmar
	PANGO_SCRIPT_MYANMAR Script = 28

	// Ogham
	PANGO_SCRIPT_OGHAM Script = 29

	// Old Italic
	PANGO_SCRIPT_OLD_ITALIC Script = 30

	// Oriya
	PANGO_SCRIPT_ORIYA Script = 31

	// Runic
	PANGO_SCRIPT_RUNIC Script = 32

	// Sinhala
	PANGO_SCRIPT_SINHALA Script = 33

	// Syriac
	PANGO_SCRIPT_SYRIAC Script = 34

	// Tamil
	PANGO_SCRIPT_TAMIL Script = 35

	// Telugu
	PANGO_SCRIPT_TELUGU Script = 36

	// Thaana
	PANGO_SCRIPT_THAANA Script = 37

	// Thai
	PANGO_SCRIPT_THAI Script = 38

	// Tibetan
	PANGO_SCRIPT_TIBETAN Script = 39

	// Canadian Aboriginal
	PANGO_SCRIPT_CANADIAN_ABORIGINAL Script = 40

	// Yi
	PANGO_SCRIPT_YI Script = 41

	// Tagalog
	PANGO_SCRIPT_TAGALOG Script = 42

	// Hanunoo
	PANGO_SCRIPT_HANUNOO Script = 43

	// Buhid
	PANGO_SCRIPT_BUHID Script = 44

	// Tagbanwa
	PANGO_SCRIPT_TAGBANWA Script = 45

	// Braille
	PANGO_SCRIPT_BRAILLE Script = 46

	// Cypriot
	PANGO_SCRIPT_CYPRIOT Script = 47

	// Limbu
	PANGO_SCRIPT_LIMBU Script = 48

	// Osmanya
	PANGO_SCRIPT_OSMANYA Script = 49

	// Shavian
	PANGO_SCRIPT_SHAVIAN Script = 50

	// Linear B
	PANGO_SCRIPT_LINEAR_B Script = 51

	// Tai Le
	PANGO_SCRIPT_TAI_LE Script = 52

	// Ugaritic
	PANGO_SCRIPT_UGARITIC Script = 53

	// New Tai Lue. Since 1.10
	PANGO_SCRIPT_NEW_TAI_LUE Script = 54

	// Buginese. Since 1.10
	PANGO_SCRIPT_BUGINESE Script = 55

	// Glagolitic. Since 1.10
	PANGO_SCRIPT_GLAGOLITIC Script = 56

	// Tifinagh. Since 1.10
	PANGO_SCRIPT_TIFINAGH Script = 57

	// Syloti Nagri. Since 1.10
	PANGO_SCRIPT_SYLOTI_NAGRI Script = 58

	// Old Persian. Since 1.10
	PANGO_SCRIPT_OLD_PERSIAN Script = 59

	// Kharoshthi. Since 1.10
	PANGO_SCRIPT_KHAROSHTHI Script = 60

	// an unassigned code point. Since 1.14
	PANGO_SCRIPT_UNKNOWN Script = 61

	// Balinese. Since 1.14
	PANGO_SCRIPT_BALINESE Script = 62

	// Cuneiform. Since 1.14
	PANGO_SCRIPT_CUNEIFORM Script = 63

	// Phoenician. Since 1.14
	PANGO_SCRIPT_PHOENICIAN Script = 64

	// Phags-pa. Since 1.14
	PANGO_SCRIPT_PHAGS_PA Script = 65

	// N'Ko. Since 1.14
	PANGO_SCRIPT_NKO Script = 66

	// Kayah Li. Since 1.20.1
	PANGO_SCRIPT_KAYAH_LI Script = 67

	// Lepcha. Since 1.20.1
	PANGO_SCRIPT_LEPCHA Script = 68

	// Rejang. Since 1.20.1
	PANGO_SCRIPT_REJANG Script = 69

	// Sundanese. Since 1.20.1
	PANGO_SCRIPT_SUNDANESE Script = 70

	// Saurashtra. Since 1.20.1
	PANGO_SCRIPT_SAURASHTRA Script = 71

	// Cham. Since 1.20.1
	PANGO_SCRIPT_CHAM Script = 72

	// Ol Chiki. Since 1.20.1
	PANGO_SCRIPT_OL_CHIKI Script = 73

	// Vai. Since 1.20.1
	PANGO_SCRIPT_VAI Script = 74

	// Carian. Since 1.20.1
	PANGO_SCRIPT_CARIAN Script = 75

	// Lycian. Since 1.20.1
	PANGO_SCRIPT_LYCIAN Script = 76

	// Lydian. Since 1.20.1
	PANGO_SCRIPT_LYDIAN Script = 77

	// Batak. Since 1.32
	PANGO_SCRIPT_BATAK Script = 78

	// Brahmi. Since 1.32
	PANGO_SCRIPT_BRAHMI Script = 79

	// Mandaic. Since 1.32
	PANGO_SCRIPT_MANDAIC Script = 80

	// Chakma. Since: 1.32
	PANGO_SCRIPT_CHAKMA Script = 81

	// Meroitic Cursive. Since: 1.32
	PANGO_SCRIPT_MEROITIC_CURSIVE Script = 82

	// Meroitic Hieroglyphs. Since: 1.32
	PANGO_SCRIPT_MEROITIC_HIEROGLYPHS Script = 83

	// Miao. Since: 1.32
	PANGO_SCRIPT_MIAO Script = 84

	// Sharada. Since: 1.32
	PANGO_SCRIPT_SHARADA Script = 85

	// Sora Sompeng. Since: 1.32
	PANGO_SCRIPT_SORA_SOMPENG Script = 86

	// Takri. Since: 1.32
	PANGO_SCRIPT_TAKRI Script = 87

	// Bassa. Since: 1.40
	PANGO_SCRIPT_BASSA_VAH Script = 88

	// Caucasian Albanian. Since: 1.40
	PANGO_SCRIPT_CAUCASIAN_ALBANIAN Script = 89

	// Duployan. Since: 1.40
	PANGO_SCRIPT_DUPLOYAN Script = 90

	// Elbasan. Since: 1.40
	PANGO_SCRIPT_ELBASAN Script = 91

	// Grantha. Since: 1.40
	PANGO_SCRIPT_GRANTHA Script = 92

	// Kjohki. Since: 1.40
	PANGO_SCRIPT_KHOJKI Script = 93

	// Khudawadi, Sindhi. Since: 1.40
	PANGO_SCRIPT_KHUDAWADI Script = 94

	// Linear A. Since: 1.40
	PANGO_SCRIPT_LINEAR_A Script = 95

	// Mahajani. Since: 1.40
	PANGO_SCRIPT_MAHAJANI Script = 96

	// Manichaean. Since: 1.40
	PANGO_SCRIPT_MANICHAEAN Script = 97

	// Mende Kikakui. Since: 1.40
	PANGO_SCRIPT_MENDE_KIKAKUI Script = 98

	// Modi. Since: 1.40
	PANGO_SCRIPT_MODI Script = 99

	// Mro. Since: 1.40
	PANGO_SCRIPT_MRO Script = 100

	// Nabataean. Since: 1.40
	PANGO_SCRIPT_NABATAEAN Script = 101

	// Old North Arabian. Since: 1.40
	PANGO_SCRIPT_OLD_NORTH_ARABIAN Script = 102

	// Old Permic. Since: 1.40
	PANGO_SCRIPT_OLD_PERMIC Script = 103

	// Pahawh Hmong. Since: 1.40
	PANGO_SCRIPT_PAHAWH_HMONG Script = 104

	// Palmyrene. Since: 1.40
	PANGO_SCRIPT_PALMYRENE Script = 105

	// Pau Cin Hau. Since: 1.40
	PANGO_SCRIPT_PAU_CIN_HAU Script = 106

	// Psalter Pahlavi. Since: 1.40
	PANGO_SCRIPT_PSALTER_PAHLAVI Script = 107

	// Siddham. Since: 1.40
	PANGO_SCRIPT_SIDDHAM Script = 108

	// Tirhuta. Since: 1.40
	PANGO_SCRIPT_TIRHUTA Script = 109

	// Warang Citi. Since: 1.40
	PANGO_SCRIPT_WARANG_CITI Script = 110

	// Ahom. Since: 1.40
	PANGO_SCRIPT_AHOM Script = 111

	// Anatolian Hieroglyphs. Since: 1.40
	PANGO_SCRIPT_ANATOLIAN_HIEROGLYPHS Script = 112

	// Hatran. Since: 1.40
	PANGO_SCRIPT_HATRAN Script = 113

	// Multani. Since: 1.40
	PANGO_SCRIPT_MULTANI Script = 114

	// Old Hungarian. Since: 1.40
	PANGO_SCRIPT_OLD_HUNGARIAN Script = 115

	// Signwriting. Since: 1.40
	PANGO_SCRIPT_SIGNWRITING Script = 116
)

// An enumeration specifying the width of the font relative to other designs
// within a family.
type Stretch C.PangoStretch

const (
	// ultra condensed width
	PANGO_STRETCH_ULTRA_CONDENSED Stretch = 0

	// extra condensed width
	PANGO_STRETCH_EXTRA_CONDENSED Stretch = 1

	// condensed width
	PANGO_STRETCH_CONDENSED Stretch = 2

	// semi condensed width
	PANGO_STRETCH_SEMI_CONDENSED Stretch = 3

	// the normal width
	PANGO_STRETCH_NORMAL Stretch = 4

	// semi expanded width
	PANGO_STRETCH_SEMI_EXPANDED Stretch = 5

	// expanded width
	PANGO_STRETCH_EXPANDED Stretch = 6

	// extra expanded width
	PANGO_STRETCH_EXTRA_EXPANDED Stretch = 7

	// ultra expanded width
	PANGO_STRETCH_ULTRA_EXPANDED Stretch = 8
)

// An enumeration specifying the various slant styles possible for a font.
type Style C.PangoStyle

const (
	// the font is upright.
	PANGO_STYLE_NORMAL Style = 0

	// the font is slanted, but in a roman style.
	PANGO_STYLE_OBLIQUE Style = 1

	// the font is slanted in an italic style.
	PANGO_STYLE_ITALIC Style = 2
)

// A #PangoTabAlign specifies where a tab stop appears relative to the text.
type TabAlign C.PangoTabAlign

const (
	// the tab stop appears to the left of the text.
	PANGO_TAB_LEFT TabAlign = 0
)

// The #PangoUnderline enumeration is used to specify
// whether text should be underlined, and if so, the type
// of underlining.
type Underline C.PangoUnderline

const (
	// no underline should be drawn
	PANGO_UNDERLINE_NONE Underline = 0

	// a single underline should be drawn
	PANGO_UNDERLINE_SINGLE Underline = 1

	// a double underline should be drawn
	PANGO_UNDERLINE_DOUBLE Underline = 2

	// a single underline should be drawn at a position
	// beneath the ink extents of the text being
	// underlined. This should be used only for underlining
	// single characters, such as for keyboard
	// accelerators. %PANGO_UNDERLINE_SINGLE should
	// be used for extended portions of text.
	PANGO_UNDERLINE_LOW Underline = 3

	// a wavy underline should be drawn below.
	// This underline is typically used to indicate
	// an error such as a possilble mispelling; in some
	// cases a contrasting color may automatically
	// be used. This type of underlining is available since Pango 1.4.
	PANGO_UNDERLINE_ERROR Underline = 4
)

// An enumeration specifying capitalization variant of the font.
type Variant C.PangoVariant

const (
	// A normal font.
	PANGO_VARIANT_NORMAL Variant = 0

	// A font with the lower case characters
	// replaced by smaller variants of the capital characters.
	PANGO_VARIANT_SMALL_CAPS Variant = 1
)

// An enumeration specifying the weight (boldness) of a font. This is a numerical
// value ranging from 100 to 1000, but there are some predefined values:
type Weight C.PangoWeight

const (
	// the thin weight (= 100; Since: 1.24)
	PANGO_WEIGHT_THIN Weight = 100

	// the ultralight weight (= 200)
	PANGO_WEIGHT_ULTRALIGHT Weight = 200

	// the light weight (= 300)
	PANGO_WEIGHT_LIGHT Weight = 300

	// the semilight weight (= 350; Since: 1.36.7)
	PANGO_WEIGHT_SEMILIGHT Weight = 350

	// the book weight (= 380; Since: 1.24)
	PANGO_WEIGHT_BOOK Weight = 380

	// the default weight (= 400)
	PANGO_WEIGHT_NORMAL Weight = 400

	// the normal weight (= 500; Since: 1.24)
	PANGO_WEIGHT_MEDIUM Weight = 500

	// the semibold weight (= 600)
	PANGO_WEIGHT_SEMIBOLD Weight = 600

	// the bold weight (= 700)
	PANGO_WEIGHT_BOLD Weight = 700

	// the ultrabold weight (= 800)
	PANGO_WEIGHT_ULTRABOLD Weight = 800

	// the heavy weight (= 900)
	PANGO_WEIGHT_HEAVY Weight = 900

	// the ultraheavy weight (= 1000; Since: 1.24)
	PANGO_WEIGHT_ULTRAHEAVY Weight = 1000
)

// A #PangoWrapMode describes how to wrap the lines of a #PangoLayout to the desired width.
type WrapMode C.PangoWrapMode

const (
	// wrap lines at word boundaries.
	PANGO_WRAP_WORD WrapMode = 0

	// wrap lines at character boundaries.
	PANGO_WRAP_CHAR WrapMode = 1

	// wrap lines at word boundaries, but fall back to character boundaries if there is not
	// enough space for a full word.
	PANGO_WRAP_WORD_CHAR WrapMode = 2
)
