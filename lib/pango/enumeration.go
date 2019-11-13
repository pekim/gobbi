// Code generated - DO NOT EDIT.

package pango

// Alignment is a representation of the C type Alignment.
type Alignment int

const (
	// left
	PangoAlignLeft Alignment = 0
	// center
	PangoAlignCenter Alignment = 1
	// right
	PangoAlignRight Alignment = 2
)

// Attrtype is a representation of the C type AttrType.
type Attrtype int

const (
	// invalid
	PangoAttrInvalid Attrtype = 0
	// language
	PangoAttrLanguage Attrtype = 1
	// family
	PangoAttrFamily Attrtype = 2
	// style
	PangoAttrStyle Attrtype = 3
	// weight
	PangoAttrWeight Attrtype = 4
	// variant
	PangoAttrVariant Attrtype = 5
	// stretch
	PangoAttrStretch Attrtype = 6
	// size
	PangoAttrSize Attrtype = 7
	// font_desc
	PangoAttrFontDesc Attrtype = 8
	// foreground
	PangoAttrForeground Attrtype = 9
	// background
	PangoAttrBackground Attrtype = 10
	// underline
	PangoAttrUnderline Attrtype = 11
	// strikethrough
	PangoAttrStrikethrough Attrtype = 12
	// rise
	PangoAttrRise Attrtype = 13
	// shape
	PangoAttrShape Attrtype = 14
	// scale
	PangoAttrScale Attrtype = 15
	// fallback
	PangoAttrFallback Attrtype = 16
	// letter_spacing
	PangoAttrLetterSpacing Attrtype = 17
	// underline_color
	PangoAttrUnderlineColor Attrtype = 18
	// strikethrough_color
	PangoAttrStrikethroughColor Attrtype = 19
	// absolute_size
	PangoAttrAbsoluteSize Attrtype = 20
	// gravity
	PangoAttrGravity Attrtype = 21
	// gravity_hint
	PangoAttrGravityHint Attrtype = 22
	// font_features
	PangoAttrFontFeatures Attrtype = 23
	// foreground_alpha
	PangoAttrForegroundAlpha Attrtype = 24
	// background_alpha
	PangoAttrBackgroundAlpha Attrtype = 25
)

// Biditype is a representation of the C type BidiType.
//
// since 1.22
type Biditype int

const (
	// l
	PangoBidiTypeL Biditype = 0
	// lre
	PangoBidiTypeLre Biditype = 1
	// lro
	PangoBidiTypeLro Biditype = 2
	// r
	PangoBidiTypeR Biditype = 3
	// al
	PangoBidiTypeAl Biditype = 4
	// rle
	PangoBidiTypeRle Biditype = 5
	// rlo
	PangoBidiTypeRlo Biditype = 6
	// pdf
	PangoBidiTypePdf Biditype = 7
	// en
	PangoBidiTypeEn Biditype = 8
	// es
	PangoBidiTypeEs Biditype = 9
	// et
	PangoBidiTypeEt Biditype = 10
	// an
	PangoBidiTypeAn Biditype = 11
	// cs
	PangoBidiTypeCs Biditype = 12
	// nsm
	PangoBidiTypeNsm Biditype = 13
	// bn
	PangoBidiTypeBn Biditype = 14
	// b
	PangoBidiTypeB Biditype = 15
	// s
	PangoBidiTypeS Biditype = 16
	// ws
	PangoBidiTypeWs Biditype = 17
	// on
	PangoBidiTypeOn Biditype = 18
)

// Coveragelevel is a representation of the C type CoverageLevel.
type Coveragelevel int

const (
	// none
	PangoCoverageNone Coveragelevel = 0
	// fallback
	PangoCoverageFallback Coveragelevel = 1
	// approximate
	PangoCoverageApproximate Coveragelevel = 2
	// exact
	PangoCoverageExact Coveragelevel = 3
)

// Direction is a representation of the C type Direction.
type Direction int

const (
	// ltr
	PangoDirectionLtr Direction = 0
	// rtl
	PangoDirectionRtl Direction = 1
	// ttb_ltr
	PangoDirectionTtbLtr Direction = 2
	// ttb_rtl
	PangoDirectionTtbRtl Direction = 3
	// weak_ltr
	PangoDirectionWeakLtr Direction = 4
	// weak_rtl
	PangoDirectionWeakRtl Direction = 5
	// neutral
	PangoDirectionNeutral Direction = 6
)

// Ellipsizemode is a representation of the C type EllipsizeMode.
type Ellipsizemode int

const (
	// none
	PangoEllipsizeNone Ellipsizemode = 0
	// start
	PangoEllipsizeStart Ellipsizemode = 1
	// middle
	PangoEllipsizeMiddle Ellipsizemode = 2
	// end
	PangoEllipsizeEnd Ellipsizemode = 3
)

// Gravity is a representation of the C type Gravity.
//
// since 1.16
type Gravity int

const (
	// south
	PangoGravitySouth Gravity = 0
	// east
	PangoGravityEast Gravity = 1
	// north
	PangoGravityNorth Gravity = 2
	// west
	PangoGravityWest Gravity = 3
	// auto
	PangoGravityAuto Gravity = 4
)

// Gravityhint is a representation of the C type GravityHint.
//
// since 1.16
type Gravityhint int

const (
	// natural
	PangoGravityHintNatural Gravityhint = 0
	// strong
	PangoGravityHintStrong Gravityhint = 1
	// line
	PangoGravityHintLine Gravityhint = 2
)

// Renderpart is a representation of the C type RenderPart.
//
// since 1.8
type Renderpart int

const (
	// foreground
	PangoRenderPartForeground Renderpart = 0
	// background
	PangoRenderPartBackground Renderpart = 1
	// underline
	PangoRenderPartUnderline Renderpart = 2
	// strikethrough
	PangoRenderPartStrikethrough Renderpart = 3
)

// Script is a representation of the C type Script.
type Script int

const (
	// invalid_code
	PangoScriptInvalidCode Script = -1
	// common
	PangoScriptCommon Script = 0
	// inherited
	PangoScriptInherited Script = 1
	// arabic
	PangoScriptArabic Script = 2
	// armenian
	PangoScriptArmenian Script = 3
	// bengali
	PangoScriptBengali Script = 4
	// bopomofo
	PangoScriptBopomofo Script = 5
	// cherokee
	PangoScriptCherokee Script = 6
	// coptic
	PangoScriptCoptic Script = 7
	// cyrillic
	PangoScriptCyrillic Script = 8
	// deseret
	PangoScriptDeseret Script = 9
	// devanagari
	PangoScriptDevanagari Script = 10
	// ethiopic
	PangoScriptEthiopic Script = 11
	// georgian
	PangoScriptGeorgian Script = 12
	// gothic
	PangoScriptGothic Script = 13
	// greek
	PangoScriptGreek Script = 14
	// gujarati
	PangoScriptGujarati Script = 15
	// gurmukhi
	PangoScriptGurmukhi Script = 16
	// han
	PangoScriptHan Script = 17
	// hangul
	PangoScriptHangul Script = 18
	// hebrew
	PangoScriptHebrew Script = 19
	// hiragana
	PangoScriptHiragana Script = 20
	// kannada
	PangoScriptKannada Script = 21
	// katakana
	PangoScriptKatakana Script = 22
	// khmer
	PangoScriptKhmer Script = 23
	// lao
	PangoScriptLao Script = 24
	// latin
	PangoScriptLatin Script = 25
	// malayalam
	PangoScriptMalayalam Script = 26
	// mongolian
	PangoScriptMongolian Script = 27
	// myanmar
	PangoScriptMyanmar Script = 28
	// ogham
	PangoScriptOgham Script = 29
	// old_italic
	PangoScriptOldItalic Script = 30
	// oriya
	PangoScriptOriya Script = 31
	// runic
	PangoScriptRunic Script = 32
	// sinhala
	PangoScriptSinhala Script = 33
	// syriac
	PangoScriptSyriac Script = 34
	// tamil
	PangoScriptTamil Script = 35
	// telugu
	PangoScriptTelugu Script = 36
	// thaana
	PangoScriptThaana Script = 37
	// thai
	PangoScriptThai Script = 38
	// tibetan
	PangoScriptTibetan Script = 39
	// canadian_aboriginal
	PangoScriptCanadianAboriginal Script = 40
	// yi
	PangoScriptYi Script = 41
	// tagalog
	PangoScriptTagalog Script = 42
	// hanunoo
	PangoScriptHanunoo Script = 43
	// buhid
	PangoScriptBuhid Script = 44
	// tagbanwa
	PangoScriptTagbanwa Script = 45
	// braille
	PangoScriptBraille Script = 46
	// cypriot
	PangoScriptCypriot Script = 47
	// limbu
	PangoScriptLimbu Script = 48
	// osmanya
	PangoScriptOsmanya Script = 49
	// shavian
	PangoScriptShavian Script = 50
	// linear_b
	PangoScriptLinearB Script = 51
	// tai_le
	PangoScriptTaiLe Script = 52
	// ugaritic
	PangoScriptUgaritic Script = 53
	// new_tai_lue
	PangoScriptNewTaiLue Script = 54
	// buginese
	PangoScriptBuginese Script = 55
	// glagolitic
	PangoScriptGlagolitic Script = 56
	// tifinagh
	PangoScriptTifinagh Script = 57
	// syloti_nagri
	PangoScriptSylotiNagri Script = 58
	// old_persian
	PangoScriptOldPersian Script = 59
	// kharoshthi
	PangoScriptKharoshthi Script = 60
	// unknown
	PangoScriptUnknown Script = 61
	// balinese
	PangoScriptBalinese Script = 62
	// cuneiform
	PangoScriptCuneiform Script = 63
	// phoenician
	PangoScriptPhoenician Script = 64
	// phags_pa
	PangoScriptPhagsPa Script = 65
	// nko
	PangoScriptNko Script = 66
	// kayah_li
	PangoScriptKayahLi Script = 67
	// lepcha
	PangoScriptLepcha Script = 68
	// rejang
	PangoScriptRejang Script = 69
	// sundanese
	PangoScriptSundanese Script = 70
	// saurashtra
	PangoScriptSaurashtra Script = 71
	// cham
	PangoScriptCham Script = 72
	// ol_chiki
	PangoScriptOlChiki Script = 73
	// vai
	PangoScriptVai Script = 74
	// carian
	PangoScriptCarian Script = 75
	// lycian
	PangoScriptLycian Script = 76
	// lydian
	PangoScriptLydian Script = 77
	// batak
	PangoScriptBatak Script = 78
	// brahmi
	PangoScriptBrahmi Script = 79
	// mandaic
	PangoScriptMandaic Script = 80
	// chakma
	PangoScriptChakma Script = 81
	// meroitic_cursive
	PangoScriptMeroiticCursive Script = 82
	// meroitic_hieroglyphs
	PangoScriptMeroiticHieroglyphs Script = 83
	// miao
	PangoScriptMiao Script = 84
	// sharada
	PangoScriptSharada Script = 85
	// sora_sompeng
	PangoScriptSoraSompeng Script = 86
	// takri
	PangoScriptTakri Script = 87
	// bassa_vah
	PangoScriptBassaVah Script = 88
	// caucasian_albanian
	PangoScriptCaucasianAlbanian Script = 89
	// duployan
	PangoScriptDuployan Script = 90
	// elbasan
	PangoScriptElbasan Script = 91
	// grantha
	PangoScriptGrantha Script = 92
	// khojki
	PangoScriptKhojki Script = 93
	// khudawadi
	PangoScriptKhudawadi Script = 94
	// linear_a
	PangoScriptLinearA Script = 95
	// mahajani
	PangoScriptMahajani Script = 96
	// manichaean
	PangoScriptManichaean Script = 97
	// mende_kikakui
	PangoScriptMendeKikakui Script = 98
	// modi
	PangoScriptModi Script = 99
	// mro
	PangoScriptMro Script = 100
	// nabataean
	PangoScriptNabataean Script = 101
	// old_north_arabian
	PangoScriptOldNorthArabian Script = 102
	// old_permic
	PangoScriptOldPermic Script = 103
	// pahawh_hmong
	PangoScriptPahawhHmong Script = 104
	// palmyrene
	PangoScriptPalmyrene Script = 105
	// pau_cin_hau
	PangoScriptPauCinHau Script = 106
	// psalter_pahlavi
	PangoScriptPsalterPahlavi Script = 107
	// siddham
	PangoScriptSiddham Script = 108
	// tirhuta
	PangoScriptTirhuta Script = 109
	// warang_citi
	PangoScriptWarangCiti Script = 110
	// ahom
	PangoScriptAhom Script = 111
	// anatolian_hieroglyphs
	PangoScriptAnatolianHieroglyphs Script = 112
	// hatran
	PangoScriptHatran Script = 113
	// multani
	PangoScriptMultani Script = 114
	// old_hungarian
	PangoScriptOldHungarian Script = 115
	// signwriting
	PangoScriptSignwriting Script = 116
)

// Stretch is a representation of the C type Stretch.
type Stretch int

const (
	// ultra_condensed
	PangoStretchUltraCondensed Stretch = 0
	// extra_condensed
	PangoStretchExtraCondensed Stretch = 1
	// condensed
	PangoStretchCondensed Stretch = 2
	// semi_condensed
	PangoStretchSemiCondensed Stretch = 3
	// normal
	PangoStretchNormal Stretch = 4
	// semi_expanded
	PangoStretchSemiExpanded Stretch = 5
	// expanded
	PangoStretchExpanded Stretch = 6
	// extra_expanded
	PangoStretchExtraExpanded Stretch = 7
	// ultra_expanded
	PangoStretchUltraExpanded Stretch = 8
)

// Style is a representation of the C type Style.
type Style int

const (
	// normal
	PangoStyleNormal Style = 0
	// oblique
	PangoStyleOblique Style = 1
	// italic
	PangoStyleItalic Style = 2
)

// Tabalign is a representation of the C type TabAlign.
type Tabalign int

const (
	// left
	PangoTabLeft Tabalign = 0
)

// Underline is a representation of the C type Underline.
type Underline int

const (
	// none
	PangoUnderlineNone Underline = 0
	// single
	PangoUnderlineSingle Underline = 1
	// double
	PangoUnderlineDouble Underline = 2
	// low
	PangoUnderlineLow Underline = 3
	// error
	PangoUnderlineError Underline = 4
)

// Variant is a representation of the C type Variant.
type Variant int

const (
	// normal
	PangoVariantNormal Variant = 0
	// small_caps
	PangoVariantSmallCaps Variant = 1
)

// Weight is a representation of the C type Weight.
type Weight int

const (
	// thin
	PangoWeightThin Weight = 100
	// ultralight
	PangoWeightUltralight Weight = 200
	// light
	PangoWeightLight Weight = 300
	// semilight
	PangoWeightSemilight Weight = 350
	// book
	PangoWeightBook Weight = 380
	// normal
	PangoWeightNormal Weight = 400
	// medium
	PangoWeightMedium Weight = 500
	// semibold
	PangoWeightSemibold Weight = 600
	// bold
	PangoWeightBold Weight = 700
	// ultrabold
	PangoWeightUltrabold Weight = 800
	// heavy
	PangoWeightHeavy Weight = 900
	// ultraheavy
	PangoWeightUltraheavy Weight = 1000
)

// Wrapmode is a representation of the C type WrapMode.
type Wrapmode int

const (
	// word
	PangoWrapWord Wrapmode = 0
	// char
	PangoWrapChar Wrapmode = 1
	// word_char
	PangoWrapWordChar Wrapmode = 2
)
