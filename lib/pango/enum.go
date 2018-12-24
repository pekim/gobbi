// This is a generated file - DO NOT EDIT

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

type Alignment C.PangoAlignment

const (
	PANGO_ALIGN_LEFT   Alignment = 0
	PANGO_ALIGN_CENTER Alignment = 1
	PANGO_ALIGN_RIGHT  Alignment = 2
)

type AttrType C.PangoAttrType

const (
	PANGO_ATTR_INVALID             AttrType = 0
	PANGO_ATTR_LANGUAGE            AttrType = 1
	PANGO_ATTR_FAMILY              AttrType = 2
	PANGO_ATTR_STYLE               AttrType = 3
	PANGO_ATTR_WEIGHT              AttrType = 4
	PANGO_ATTR_VARIANT             AttrType = 5
	PANGO_ATTR_STRETCH             AttrType = 6
	PANGO_ATTR_SIZE                AttrType = 7
	PANGO_ATTR_FONT_DESC           AttrType = 8
	PANGO_ATTR_FOREGROUND          AttrType = 9
	PANGO_ATTR_BACKGROUND          AttrType = 10
	PANGO_ATTR_UNDERLINE           AttrType = 11
	PANGO_ATTR_STRIKETHROUGH       AttrType = 12
	PANGO_ATTR_RISE                AttrType = 13
	PANGO_ATTR_SHAPE               AttrType = 14
	PANGO_ATTR_SCALE               AttrType = 15
	PANGO_ATTR_FALLBACK            AttrType = 16
	PANGO_ATTR_LETTER_SPACING      AttrType = 17
	PANGO_ATTR_UNDERLINE_COLOR     AttrType = 18
	PANGO_ATTR_STRIKETHROUGH_COLOR AttrType = 19
	PANGO_ATTR_ABSOLUTE_SIZE       AttrType = 20
	PANGO_ATTR_GRAVITY             AttrType = 21
	PANGO_ATTR_GRAVITY_HINT        AttrType = 22
	PANGO_ATTR_FONT_FEATURES       AttrType = 23
	PANGO_ATTR_FOREGROUND_ALPHA    AttrType = 24
	PANGO_ATTR_BACKGROUND_ALPHA    AttrType = 25
)

// AttrTypeRegister is a wrapper around the C function pango_attr_type_register.
func AttrTypeRegister(name string) AttrType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.pango_attr_type_register(c_name)
	retGo := (AttrType)(retC)

	return retGo
}

type CoverageLevel C.PangoCoverageLevel

const (
	PANGO_COVERAGE_NONE        CoverageLevel = 0
	PANGO_COVERAGE_FALLBACK    CoverageLevel = 1
	PANGO_COVERAGE_APPROXIMATE CoverageLevel = 2
	PANGO_COVERAGE_EXACT       CoverageLevel = 3
)

type Direction C.PangoDirection

const (
	PANGO_DIRECTION_LTR      Direction = 0
	PANGO_DIRECTION_RTL      Direction = 1
	PANGO_DIRECTION_TTB_LTR  Direction = 2
	PANGO_DIRECTION_TTB_RTL  Direction = 3
	PANGO_DIRECTION_WEAK_LTR Direction = 4
	PANGO_DIRECTION_WEAK_RTL Direction = 5
	PANGO_DIRECTION_NEUTRAL  Direction = 6
)

type EllipsizeMode C.PangoEllipsizeMode

const (
	PANGO_ELLIPSIZE_NONE   EllipsizeMode = 0
	PANGO_ELLIPSIZE_START  EllipsizeMode = 1
	PANGO_ELLIPSIZE_MIDDLE EllipsizeMode = 2
	PANGO_ELLIPSIZE_END    EllipsizeMode = 3
)

type Script C.PangoScript

const (
	PANGO_SCRIPT_INVALID_CODE          Script = -1
	PANGO_SCRIPT_COMMON                Script = 0
	PANGO_SCRIPT_INHERITED             Script = 1
	PANGO_SCRIPT_ARABIC                Script = 2
	PANGO_SCRIPT_ARMENIAN              Script = 3
	PANGO_SCRIPT_BENGALI               Script = 4
	PANGO_SCRIPT_BOPOMOFO              Script = 5
	PANGO_SCRIPT_CHEROKEE              Script = 6
	PANGO_SCRIPT_COPTIC                Script = 7
	PANGO_SCRIPT_CYRILLIC              Script = 8
	PANGO_SCRIPT_DESERET               Script = 9
	PANGO_SCRIPT_DEVANAGARI            Script = 10
	PANGO_SCRIPT_ETHIOPIC              Script = 11
	PANGO_SCRIPT_GEORGIAN              Script = 12
	PANGO_SCRIPT_GOTHIC                Script = 13
	PANGO_SCRIPT_GREEK                 Script = 14
	PANGO_SCRIPT_GUJARATI              Script = 15
	PANGO_SCRIPT_GURMUKHI              Script = 16
	PANGO_SCRIPT_HAN                   Script = 17
	PANGO_SCRIPT_HANGUL                Script = 18
	PANGO_SCRIPT_HEBREW                Script = 19
	PANGO_SCRIPT_HIRAGANA              Script = 20
	PANGO_SCRIPT_KANNADA               Script = 21
	PANGO_SCRIPT_KATAKANA              Script = 22
	PANGO_SCRIPT_KHMER                 Script = 23
	PANGO_SCRIPT_LAO                   Script = 24
	PANGO_SCRIPT_LATIN                 Script = 25
	PANGO_SCRIPT_MALAYALAM             Script = 26
	PANGO_SCRIPT_MONGOLIAN             Script = 27
	PANGO_SCRIPT_MYANMAR               Script = 28
	PANGO_SCRIPT_OGHAM                 Script = 29
	PANGO_SCRIPT_OLD_ITALIC            Script = 30
	PANGO_SCRIPT_ORIYA                 Script = 31
	PANGO_SCRIPT_RUNIC                 Script = 32
	PANGO_SCRIPT_SINHALA               Script = 33
	PANGO_SCRIPT_SYRIAC                Script = 34
	PANGO_SCRIPT_TAMIL                 Script = 35
	PANGO_SCRIPT_TELUGU                Script = 36
	PANGO_SCRIPT_THAANA                Script = 37
	PANGO_SCRIPT_THAI                  Script = 38
	PANGO_SCRIPT_TIBETAN               Script = 39
	PANGO_SCRIPT_CANADIAN_ABORIGINAL   Script = 40
	PANGO_SCRIPT_YI                    Script = 41
	PANGO_SCRIPT_TAGALOG               Script = 42
	PANGO_SCRIPT_HANUNOO               Script = 43
	PANGO_SCRIPT_BUHID                 Script = 44
	PANGO_SCRIPT_TAGBANWA              Script = 45
	PANGO_SCRIPT_BRAILLE               Script = 46
	PANGO_SCRIPT_CYPRIOT               Script = 47
	PANGO_SCRIPT_LIMBU                 Script = 48
	PANGO_SCRIPT_OSMANYA               Script = 49
	PANGO_SCRIPT_SHAVIAN               Script = 50
	PANGO_SCRIPT_LINEAR_B              Script = 51
	PANGO_SCRIPT_TAI_LE                Script = 52
	PANGO_SCRIPT_UGARITIC              Script = 53
	PANGO_SCRIPT_NEW_TAI_LUE           Script = 54
	PANGO_SCRIPT_BUGINESE              Script = 55
	PANGO_SCRIPT_GLAGOLITIC            Script = 56
	PANGO_SCRIPT_TIFINAGH              Script = 57
	PANGO_SCRIPT_SYLOTI_NAGRI          Script = 58
	PANGO_SCRIPT_OLD_PERSIAN           Script = 59
	PANGO_SCRIPT_KHAROSHTHI            Script = 60
	PANGO_SCRIPT_UNKNOWN               Script = 61
	PANGO_SCRIPT_BALINESE              Script = 62
	PANGO_SCRIPT_CUNEIFORM             Script = 63
	PANGO_SCRIPT_PHOENICIAN            Script = 64
	PANGO_SCRIPT_PHAGS_PA              Script = 65
	PANGO_SCRIPT_NKO                   Script = 66
	PANGO_SCRIPT_KAYAH_LI              Script = 67
	PANGO_SCRIPT_LEPCHA                Script = 68
	PANGO_SCRIPT_REJANG                Script = 69
	PANGO_SCRIPT_SUNDANESE             Script = 70
	PANGO_SCRIPT_SAURASHTRA            Script = 71
	PANGO_SCRIPT_CHAM                  Script = 72
	PANGO_SCRIPT_OL_CHIKI              Script = 73
	PANGO_SCRIPT_VAI                   Script = 74
	PANGO_SCRIPT_CARIAN                Script = 75
	PANGO_SCRIPT_LYCIAN                Script = 76
	PANGO_SCRIPT_LYDIAN                Script = 77
	PANGO_SCRIPT_BATAK                 Script = 78
	PANGO_SCRIPT_BRAHMI                Script = 79
	PANGO_SCRIPT_MANDAIC               Script = 80
	PANGO_SCRIPT_CHAKMA                Script = 81
	PANGO_SCRIPT_MEROITIC_CURSIVE      Script = 82
	PANGO_SCRIPT_MEROITIC_HIEROGLYPHS  Script = 83
	PANGO_SCRIPT_MIAO                  Script = 84
	PANGO_SCRIPT_SHARADA               Script = 85
	PANGO_SCRIPT_SORA_SOMPENG          Script = 86
	PANGO_SCRIPT_TAKRI                 Script = 87
	PANGO_SCRIPT_BASSA_VAH             Script = 88
	PANGO_SCRIPT_CAUCASIAN_ALBANIAN    Script = 89
	PANGO_SCRIPT_DUPLOYAN              Script = 90
	PANGO_SCRIPT_ELBASAN               Script = 91
	PANGO_SCRIPT_GRANTHA               Script = 92
	PANGO_SCRIPT_KHOJKI                Script = 93
	PANGO_SCRIPT_KHUDAWADI             Script = 94
	PANGO_SCRIPT_LINEAR_A              Script = 95
	PANGO_SCRIPT_MAHAJANI              Script = 96
	PANGO_SCRIPT_MANICHAEAN            Script = 97
	PANGO_SCRIPT_MENDE_KIKAKUI         Script = 98
	PANGO_SCRIPT_MODI                  Script = 99
	PANGO_SCRIPT_MRO                   Script = 100
	PANGO_SCRIPT_NABATAEAN             Script = 101
	PANGO_SCRIPT_OLD_NORTH_ARABIAN     Script = 102
	PANGO_SCRIPT_OLD_PERMIC            Script = 103
	PANGO_SCRIPT_PAHAWH_HMONG          Script = 104
	PANGO_SCRIPT_PALMYRENE             Script = 105
	PANGO_SCRIPT_PAU_CIN_HAU           Script = 106
	PANGO_SCRIPT_PSALTER_PAHLAVI       Script = 107
	PANGO_SCRIPT_SIDDHAM               Script = 108
	PANGO_SCRIPT_TIRHUTA               Script = 109
	PANGO_SCRIPT_WARANG_CITI           Script = 110
	PANGO_SCRIPT_AHOM                  Script = 111
	PANGO_SCRIPT_ANATOLIAN_HIEROGLYPHS Script = 112
	PANGO_SCRIPT_HATRAN                Script = 113
	PANGO_SCRIPT_MULTANI               Script = 114
	PANGO_SCRIPT_OLD_HUNGARIAN         Script = 115
	PANGO_SCRIPT_SIGNWRITING           Script = 116
)

type Stretch C.PangoStretch

const (
	PANGO_STRETCH_ULTRA_CONDENSED Stretch = 0
	PANGO_STRETCH_EXTRA_CONDENSED Stretch = 1
	PANGO_STRETCH_CONDENSED       Stretch = 2
	PANGO_STRETCH_SEMI_CONDENSED  Stretch = 3
	PANGO_STRETCH_NORMAL          Stretch = 4
	PANGO_STRETCH_SEMI_EXPANDED   Stretch = 5
	PANGO_STRETCH_EXPANDED        Stretch = 6
	PANGO_STRETCH_EXTRA_EXPANDED  Stretch = 7
	PANGO_STRETCH_ULTRA_EXPANDED  Stretch = 8
)

type Style C.PangoStyle

const (
	PANGO_STYLE_NORMAL  Style = 0
	PANGO_STYLE_OBLIQUE Style = 1
	PANGO_STYLE_ITALIC  Style = 2
)

type TabAlign C.PangoTabAlign

const (
	PANGO_TAB_LEFT TabAlign = 0
)

type Underline C.PangoUnderline

const (
	PANGO_UNDERLINE_NONE   Underline = 0
	PANGO_UNDERLINE_SINGLE Underline = 1
	PANGO_UNDERLINE_DOUBLE Underline = 2
	PANGO_UNDERLINE_LOW    Underline = 3
	PANGO_UNDERLINE_ERROR  Underline = 4
)

type Variant C.PangoVariant

const (
	PANGO_VARIANT_NORMAL     Variant = 0
	PANGO_VARIANT_SMALL_CAPS Variant = 1
)

type Weight C.PangoWeight

const (
	PANGO_WEIGHT_THIN       Weight = 100
	PANGO_WEIGHT_ULTRALIGHT Weight = 200
	PANGO_WEIGHT_LIGHT      Weight = 300
	PANGO_WEIGHT_SEMILIGHT  Weight = 350
	PANGO_WEIGHT_BOOK       Weight = 380
	PANGO_WEIGHT_NORMAL     Weight = 400
	PANGO_WEIGHT_MEDIUM     Weight = 500
	PANGO_WEIGHT_SEMIBOLD   Weight = 600
	PANGO_WEIGHT_BOLD       Weight = 700
	PANGO_WEIGHT_ULTRABOLD  Weight = 800
	PANGO_WEIGHT_HEAVY      Weight = 900
	PANGO_WEIGHT_ULTRAHEAVY Weight = 1000
)

type WrapMode C.PangoWrapMode

const (
	PANGO_WRAP_WORD      WrapMode = 0
	PANGO_WRAP_CHAR      WrapMode = 1
	PANGO_WRAP_WORD_CHAR WrapMode = 2
)
