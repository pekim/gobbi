// Code generated - DO NOT EDIT.

package pango

// Alignment is a representation of the C type PangoAlignment.
type Alignment int

const (
	// Alignment_Left is a representation of the C type PANGO_ALIGN_LEFT.
	Alignment_Left Alignment = 0
	// Alignment_Center is a representation of the C type PANGO_ALIGN_CENTER.
	Alignment_Center Alignment = 1
	// Alignment_Right is a representation of the C type PANGO_ALIGN_RIGHT.
	Alignment_Right Alignment = 2
)

// Attrtype is a representation of the C type PangoAttrType.
type Attrtype int

const (
	// Attrtype_Invalid is a representation of the C type PANGO_ATTR_INVALID.
	Attrtype_Invalid Attrtype = 0
	// Attrtype_Language is a representation of the C type PANGO_ATTR_LANGUAGE.
	Attrtype_Language Attrtype = 1
	// Attrtype_Family is a representation of the C type PANGO_ATTR_FAMILY.
	Attrtype_Family Attrtype = 2
	// Attrtype_Style is a representation of the C type PANGO_ATTR_STYLE.
	Attrtype_Style Attrtype = 3
	// Attrtype_Weight is a representation of the C type PANGO_ATTR_WEIGHT.
	Attrtype_Weight Attrtype = 4
	// Attrtype_Variant is a representation of the C type PANGO_ATTR_VARIANT.
	Attrtype_Variant Attrtype = 5
	// Attrtype_Stretch is a representation of the C type PANGO_ATTR_STRETCH.
	Attrtype_Stretch Attrtype = 6
	// Attrtype_Size is a representation of the C type PANGO_ATTR_SIZE.
	Attrtype_Size Attrtype = 7
	// Attrtype_FontDesc is a representation of the C type PANGO_ATTR_FONT_DESC.
	Attrtype_FontDesc Attrtype = 8
	// Attrtype_Foreground is a representation of the C type PANGO_ATTR_FOREGROUND.
	Attrtype_Foreground Attrtype = 9
	// Attrtype_Background is a representation of the C type PANGO_ATTR_BACKGROUND.
	Attrtype_Background Attrtype = 10
	// Attrtype_Underline is a representation of the C type PANGO_ATTR_UNDERLINE.
	Attrtype_Underline Attrtype = 11
	// Attrtype_Strikethrough is a representation of the C type PANGO_ATTR_STRIKETHROUGH.
	Attrtype_Strikethrough Attrtype = 12
	// Attrtype_Rise is a representation of the C type PANGO_ATTR_RISE.
	Attrtype_Rise Attrtype = 13
	// Attrtype_Shape is a representation of the C type PANGO_ATTR_SHAPE.
	Attrtype_Shape Attrtype = 14
	// Attrtype_Scale is a representation of the C type PANGO_ATTR_SCALE.
	Attrtype_Scale Attrtype = 15
	// Attrtype_Fallback is a representation of the C type PANGO_ATTR_FALLBACK.
	Attrtype_Fallback Attrtype = 16
	// Attrtype_LetterSpacing is a representation of the C type PANGO_ATTR_LETTER_SPACING.
	Attrtype_LetterSpacing Attrtype = 17
	// Attrtype_UnderlineColor is a representation of the C type PANGO_ATTR_UNDERLINE_COLOR.
	Attrtype_UnderlineColor Attrtype = 18
	// Attrtype_StrikethroughColor is a representation of the C type PANGO_ATTR_STRIKETHROUGH_COLOR.
	Attrtype_StrikethroughColor Attrtype = 19
	// Attrtype_AbsoluteSize is a representation of the C type PANGO_ATTR_ABSOLUTE_SIZE.
	Attrtype_AbsoluteSize Attrtype = 20
	// Attrtype_Gravity is a representation of the C type PANGO_ATTR_GRAVITY.
	Attrtype_Gravity Attrtype = 21
	// Attrtype_GravityHint is a representation of the C type PANGO_ATTR_GRAVITY_HINT.
	Attrtype_GravityHint Attrtype = 22
	// Attrtype_FontFeatures is a representation of the C type PANGO_ATTR_FONT_FEATURES.
	Attrtype_FontFeatures Attrtype = 23
	// Attrtype_ForegroundAlpha is a representation of the C type PANGO_ATTR_FOREGROUND_ALPHA.
	Attrtype_ForegroundAlpha Attrtype = 24
	// Attrtype_BackgroundAlpha is a representation of the C type PANGO_ATTR_BACKGROUND_ALPHA.
	Attrtype_BackgroundAlpha Attrtype = 25
)

// Biditype is a representation of the C type PangoBidiType.
//
// since 1.22
type Biditype int

const (
	// Biditype_L is a representation of the C type PANGO_BIDI_TYPE_L.
	Biditype_L Biditype = 0
	// Biditype_Lre is a representation of the C type PANGO_BIDI_TYPE_LRE.
	Biditype_Lre Biditype = 1
	// Biditype_Lro is a representation of the C type PANGO_BIDI_TYPE_LRO.
	Biditype_Lro Biditype = 2
	// Biditype_R is a representation of the C type PANGO_BIDI_TYPE_R.
	Biditype_R Biditype = 3
	// Biditype_Al is a representation of the C type PANGO_BIDI_TYPE_AL.
	Biditype_Al Biditype = 4
	// Biditype_Rle is a representation of the C type PANGO_BIDI_TYPE_RLE.
	Biditype_Rle Biditype = 5
	// Biditype_Rlo is a representation of the C type PANGO_BIDI_TYPE_RLO.
	Biditype_Rlo Biditype = 6
	// Biditype_Pdf is a representation of the C type PANGO_BIDI_TYPE_PDF.
	Biditype_Pdf Biditype = 7
	// Biditype_En is a representation of the C type PANGO_BIDI_TYPE_EN.
	Biditype_En Biditype = 8
	// Biditype_Es is a representation of the C type PANGO_BIDI_TYPE_ES.
	Biditype_Es Biditype = 9
	// Biditype_Et is a representation of the C type PANGO_BIDI_TYPE_ET.
	Biditype_Et Biditype = 10
	// Biditype_An is a representation of the C type PANGO_BIDI_TYPE_AN.
	Biditype_An Biditype = 11
	// Biditype_Cs is a representation of the C type PANGO_BIDI_TYPE_CS.
	Biditype_Cs Biditype = 12
	// Biditype_Nsm is a representation of the C type PANGO_BIDI_TYPE_NSM.
	Biditype_Nsm Biditype = 13
	// Biditype_Bn is a representation of the C type PANGO_BIDI_TYPE_BN.
	Biditype_Bn Biditype = 14
	// Biditype_B is a representation of the C type PANGO_BIDI_TYPE_B.
	Biditype_B Biditype = 15
	// Biditype_S is a representation of the C type PANGO_BIDI_TYPE_S.
	Biditype_S Biditype = 16
	// Biditype_Ws is a representation of the C type PANGO_BIDI_TYPE_WS.
	Biditype_Ws Biditype = 17
	// Biditype_On is a representation of the C type PANGO_BIDI_TYPE_ON.
	Biditype_On Biditype = 18
)

// Coveragelevel is a representation of the C type PangoCoverageLevel.
type Coveragelevel int

const (
	// Coveragelevel_None is a representation of the C type PANGO_COVERAGE_NONE.
	Coveragelevel_None Coveragelevel = 0
	// Coveragelevel_Fallback is a representation of the C type PANGO_COVERAGE_FALLBACK.
	Coveragelevel_Fallback Coveragelevel = 1
	// Coveragelevel_Approximate is a representation of the C type PANGO_COVERAGE_APPROXIMATE.
	Coveragelevel_Approximate Coveragelevel = 2
	// Coveragelevel_Exact is a representation of the C type PANGO_COVERAGE_EXACT.
	Coveragelevel_Exact Coveragelevel = 3
)

// Direction is a representation of the C type PangoDirection.
type Direction int

const (
	// Direction_Ltr is a representation of the C type PANGO_DIRECTION_LTR.
	Direction_Ltr Direction = 0
	// Direction_Rtl is a representation of the C type PANGO_DIRECTION_RTL.
	Direction_Rtl Direction = 1
	// Direction_TtbLtr is a representation of the C type PANGO_DIRECTION_TTB_LTR.
	Direction_TtbLtr Direction = 2
	// Direction_TtbRtl is a representation of the C type PANGO_DIRECTION_TTB_RTL.
	Direction_TtbRtl Direction = 3
	// Direction_WeakLtr is a representation of the C type PANGO_DIRECTION_WEAK_LTR.
	Direction_WeakLtr Direction = 4
	// Direction_WeakRtl is a representation of the C type PANGO_DIRECTION_WEAK_RTL.
	Direction_WeakRtl Direction = 5
	// Direction_Neutral is a representation of the C type PANGO_DIRECTION_NEUTRAL.
	Direction_Neutral Direction = 6
)

// Ellipsizemode is a representation of the C type PangoEllipsizeMode.
type Ellipsizemode int

const (
	// Ellipsizemode_None is a representation of the C type PANGO_ELLIPSIZE_NONE.
	Ellipsizemode_None Ellipsizemode = 0
	// Ellipsizemode_Start is a representation of the C type PANGO_ELLIPSIZE_START.
	Ellipsizemode_Start Ellipsizemode = 1
	// Ellipsizemode_Middle is a representation of the C type PANGO_ELLIPSIZE_MIDDLE.
	Ellipsizemode_Middle Ellipsizemode = 2
	// Ellipsizemode_End is a representation of the C type PANGO_ELLIPSIZE_END.
	Ellipsizemode_End Ellipsizemode = 3
)

// Gravity is a representation of the C type PangoGravity.
//
// since 1.16
type Gravity int

const (
	// Gravity_South is a representation of the C type PANGO_GRAVITY_SOUTH.
	Gravity_South Gravity = 0
	// Gravity_East is a representation of the C type PANGO_GRAVITY_EAST.
	Gravity_East Gravity = 1
	// Gravity_North is a representation of the C type PANGO_GRAVITY_NORTH.
	Gravity_North Gravity = 2
	// Gravity_West is a representation of the C type PANGO_GRAVITY_WEST.
	Gravity_West Gravity = 3
	// Gravity_Auto is a representation of the C type PANGO_GRAVITY_AUTO.
	Gravity_Auto Gravity = 4
)

// Gravityhint is a representation of the C type PangoGravityHint.
//
// since 1.16
type Gravityhint int

const (
	// Gravityhint_Natural is a representation of the C type PANGO_GRAVITY_HINT_NATURAL.
	Gravityhint_Natural Gravityhint = 0
	// Gravityhint_Strong is a representation of the C type PANGO_GRAVITY_HINT_STRONG.
	Gravityhint_Strong Gravityhint = 1
	// Gravityhint_Line is a representation of the C type PANGO_GRAVITY_HINT_LINE.
	Gravityhint_Line Gravityhint = 2
)

// Renderpart is a representation of the C type PangoRenderPart.
//
// since 1.8
type Renderpart int

const (
	// Renderpart_Foreground is a representation of the C type PANGO_RENDER_PART_FOREGROUND.
	Renderpart_Foreground Renderpart = 0
	// Renderpart_Background is a representation of the C type PANGO_RENDER_PART_BACKGROUND.
	Renderpart_Background Renderpart = 1
	// Renderpart_Underline is a representation of the C type PANGO_RENDER_PART_UNDERLINE.
	Renderpart_Underline Renderpart = 2
	// Renderpart_Strikethrough is a representation of the C type PANGO_RENDER_PART_STRIKETHROUGH.
	Renderpart_Strikethrough Renderpart = 3
)

// Script is a representation of the C type PangoScript.
type Script int

const (
	// Script_InvalidCode is a representation of the C type PANGO_SCRIPT_INVALID_CODE.
	Script_InvalidCode Script = -1
	// Script_Common is a representation of the C type PANGO_SCRIPT_COMMON.
	Script_Common Script = 0
	// Script_Inherited is a representation of the C type PANGO_SCRIPT_INHERITED.
	Script_Inherited Script = 1
	// Script_Arabic is a representation of the C type PANGO_SCRIPT_ARABIC.
	Script_Arabic Script = 2
	// Script_Armenian is a representation of the C type PANGO_SCRIPT_ARMENIAN.
	Script_Armenian Script = 3
	// Script_Bengali is a representation of the C type PANGO_SCRIPT_BENGALI.
	Script_Bengali Script = 4
	// Script_Bopomofo is a representation of the C type PANGO_SCRIPT_BOPOMOFO.
	Script_Bopomofo Script = 5
	// Script_Cherokee is a representation of the C type PANGO_SCRIPT_CHEROKEE.
	Script_Cherokee Script = 6
	// Script_Coptic is a representation of the C type PANGO_SCRIPT_COPTIC.
	Script_Coptic Script = 7
	// Script_Cyrillic is a representation of the C type PANGO_SCRIPT_CYRILLIC.
	Script_Cyrillic Script = 8
	// Script_Deseret is a representation of the C type PANGO_SCRIPT_DESERET.
	Script_Deseret Script = 9
	// Script_Devanagari is a representation of the C type PANGO_SCRIPT_DEVANAGARI.
	Script_Devanagari Script = 10
	// Script_Ethiopic is a representation of the C type PANGO_SCRIPT_ETHIOPIC.
	Script_Ethiopic Script = 11
	// Script_Georgian is a representation of the C type PANGO_SCRIPT_GEORGIAN.
	Script_Georgian Script = 12
	// Script_Gothic is a representation of the C type PANGO_SCRIPT_GOTHIC.
	Script_Gothic Script = 13
	// Script_Greek is a representation of the C type PANGO_SCRIPT_GREEK.
	Script_Greek Script = 14
	// Script_Gujarati is a representation of the C type PANGO_SCRIPT_GUJARATI.
	Script_Gujarati Script = 15
	// Script_Gurmukhi is a representation of the C type PANGO_SCRIPT_GURMUKHI.
	Script_Gurmukhi Script = 16
	// Script_Han is a representation of the C type PANGO_SCRIPT_HAN.
	Script_Han Script = 17
	// Script_Hangul is a representation of the C type PANGO_SCRIPT_HANGUL.
	Script_Hangul Script = 18
	// Script_Hebrew is a representation of the C type PANGO_SCRIPT_HEBREW.
	Script_Hebrew Script = 19
	// Script_Hiragana is a representation of the C type PANGO_SCRIPT_HIRAGANA.
	Script_Hiragana Script = 20
	// Script_Kannada is a representation of the C type PANGO_SCRIPT_KANNADA.
	Script_Kannada Script = 21
	// Script_Katakana is a representation of the C type PANGO_SCRIPT_KATAKANA.
	Script_Katakana Script = 22
	// Script_Khmer is a representation of the C type PANGO_SCRIPT_KHMER.
	Script_Khmer Script = 23
	// Script_Lao is a representation of the C type PANGO_SCRIPT_LAO.
	Script_Lao Script = 24
	// Script_Latin is a representation of the C type PANGO_SCRIPT_LATIN.
	Script_Latin Script = 25
	// Script_Malayalam is a representation of the C type PANGO_SCRIPT_MALAYALAM.
	Script_Malayalam Script = 26
	// Script_Mongolian is a representation of the C type PANGO_SCRIPT_MONGOLIAN.
	Script_Mongolian Script = 27
	// Script_Myanmar is a representation of the C type PANGO_SCRIPT_MYANMAR.
	Script_Myanmar Script = 28
	// Script_Ogham is a representation of the C type PANGO_SCRIPT_OGHAM.
	Script_Ogham Script = 29
	// Script_OldItalic is a representation of the C type PANGO_SCRIPT_OLD_ITALIC.
	Script_OldItalic Script = 30
	// Script_Oriya is a representation of the C type PANGO_SCRIPT_ORIYA.
	Script_Oriya Script = 31
	// Script_Runic is a representation of the C type PANGO_SCRIPT_RUNIC.
	Script_Runic Script = 32
	// Script_Sinhala is a representation of the C type PANGO_SCRIPT_SINHALA.
	Script_Sinhala Script = 33
	// Script_Syriac is a representation of the C type PANGO_SCRIPT_SYRIAC.
	Script_Syriac Script = 34
	// Script_Tamil is a representation of the C type PANGO_SCRIPT_TAMIL.
	Script_Tamil Script = 35
	// Script_Telugu is a representation of the C type PANGO_SCRIPT_TELUGU.
	Script_Telugu Script = 36
	// Script_Thaana is a representation of the C type PANGO_SCRIPT_THAANA.
	Script_Thaana Script = 37
	// Script_Thai is a representation of the C type PANGO_SCRIPT_THAI.
	Script_Thai Script = 38
	// Script_Tibetan is a representation of the C type PANGO_SCRIPT_TIBETAN.
	Script_Tibetan Script = 39
	// Script_CanadianAboriginal is a representation of the C type PANGO_SCRIPT_CANADIAN_ABORIGINAL.
	Script_CanadianAboriginal Script = 40
	// Script_Yi is a representation of the C type PANGO_SCRIPT_YI.
	Script_Yi Script = 41
	// Script_Tagalog is a representation of the C type PANGO_SCRIPT_TAGALOG.
	Script_Tagalog Script = 42
	// Script_Hanunoo is a representation of the C type PANGO_SCRIPT_HANUNOO.
	Script_Hanunoo Script = 43
	// Script_Buhid is a representation of the C type PANGO_SCRIPT_BUHID.
	Script_Buhid Script = 44
	// Script_Tagbanwa is a representation of the C type PANGO_SCRIPT_TAGBANWA.
	Script_Tagbanwa Script = 45
	// Script_Braille is a representation of the C type PANGO_SCRIPT_BRAILLE.
	Script_Braille Script = 46
	// Script_Cypriot is a representation of the C type PANGO_SCRIPT_CYPRIOT.
	Script_Cypriot Script = 47
	// Script_Limbu is a representation of the C type PANGO_SCRIPT_LIMBU.
	Script_Limbu Script = 48
	// Script_Osmanya is a representation of the C type PANGO_SCRIPT_OSMANYA.
	Script_Osmanya Script = 49
	// Script_Shavian is a representation of the C type PANGO_SCRIPT_SHAVIAN.
	Script_Shavian Script = 50
	// Script_LinearB is a representation of the C type PANGO_SCRIPT_LINEAR_B.
	Script_LinearB Script = 51
	// Script_TaiLe is a representation of the C type PANGO_SCRIPT_TAI_LE.
	Script_TaiLe Script = 52
	// Script_Ugaritic is a representation of the C type PANGO_SCRIPT_UGARITIC.
	Script_Ugaritic Script = 53
	// Script_NewTaiLue is a representation of the C type PANGO_SCRIPT_NEW_TAI_LUE.
	Script_NewTaiLue Script = 54
	// Script_Buginese is a representation of the C type PANGO_SCRIPT_BUGINESE.
	Script_Buginese Script = 55
	// Script_Glagolitic is a representation of the C type PANGO_SCRIPT_GLAGOLITIC.
	Script_Glagolitic Script = 56
	// Script_Tifinagh is a representation of the C type PANGO_SCRIPT_TIFINAGH.
	Script_Tifinagh Script = 57
	// Script_SylotiNagri is a representation of the C type PANGO_SCRIPT_SYLOTI_NAGRI.
	Script_SylotiNagri Script = 58
	// Script_OldPersian is a representation of the C type PANGO_SCRIPT_OLD_PERSIAN.
	Script_OldPersian Script = 59
	// Script_Kharoshthi is a representation of the C type PANGO_SCRIPT_KHAROSHTHI.
	Script_Kharoshthi Script = 60
	// Script_Unknown is a representation of the C type PANGO_SCRIPT_UNKNOWN.
	Script_Unknown Script = 61
	// Script_Balinese is a representation of the C type PANGO_SCRIPT_BALINESE.
	Script_Balinese Script = 62
	// Script_Cuneiform is a representation of the C type PANGO_SCRIPT_CUNEIFORM.
	Script_Cuneiform Script = 63
	// Script_Phoenician is a representation of the C type PANGO_SCRIPT_PHOENICIAN.
	Script_Phoenician Script = 64
	// Script_PhagsPa is a representation of the C type PANGO_SCRIPT_PHAGS_PA.
	Script_PhagsPa Script = 65
	// Script_Nko is a representation of the C type PANGO_SCRIPT_NKO.
	Script_Nko Script = 66
	// Script_KayahLi is a representation of the C type PANGO_SCRIPT_KAYAH_LI.
	Script_KayahLi Script = 67
	// Script_Lepcha is a representation of the C type PANGO_SCRIPT_LEPCHA.
	Script_Lepcha Script = 68
	// Script_Rejang is a representation of the C type PANGO_SCRIPT_REJANG.
	Script_Rejang Script = 69
	// Script_Sundanese is a representation of the C type PANGO_SCRIPT_SUNDANESE.
	Script_Sundanese Script = 70
	// Script_Saurashtra is a representation of the C type PANGO_SCRIPT_SAURASHTRA.
	Script_Saurashtra Script = 71
	// Script_Cham is a representation of the C type PANGO_SCRIPT_CHAM.
	Script_Cham Script = 72
	// Script_OlChiki is a representation of the C type PANGO_SCRIPT_OL_CHIKI.
	Script_OlChiki Script = 73
	// Script_Vai is a representation of the C type PANGO_SCRIPT_VAI.
	Script_Vai Script = 74
	// Script_Carian is a representation of the C type PANGO_SCRIPT_CARIAN.
	Script_Carian Script = 75
	// Script_Lycian is a representation of the C type PANGO_SCRIPT_LYCIAN.
	Script_Lycian Script = 76
	// Script_Lydian is a representation of the C type PANGO_SCRIPT_LYDIAN.
	Script_Lydian Script = 77
	// Script_Batak is a representation of the C type PANGO_SCRIPT_BATAK.
	Script_Batak Script = 78
	// Script_Brahmi is a representation of the C type PANGO_SCRIPT_BRAHMI.
	Script_Brahmi Script = 79
	// Script_Mandaic is a representation of the C type PANGO_SCRIPT_MANDAIC.
	Script_Mandaic Script = 80
	// Script_Chakma is a representation of the C type PANGO_SCRIPT_CHAKMA.
	Script_Chakma Script = 81
	// Script_MeroiticCursive is a representation of the C type PANGO_SCRIPT_MEROITIC_CURSIVE.
	Script_MeroiticCursive Script = 82
	// Script_MeroiticHieroglyphs is a representation of the C type PANGO_SCRIPT_MEROITIC_HIEROGLYPHS.
	Script_MeroiticHieroglyphs Script = 83
	// Script_Miao is a representation of the C type PANGO_SCRIPT_MIAO.
	Script_Miao Script = 84
	// Script_Sharada is a representation of the C type PANGO_SCRIPT_SHARADA.
	Script_Sharada Script = 85
	// Script_SoraSompeng is a representation of the C type PANGO_SCRIPT_SORA_SOMPENG.
	Script_SoraSompeng Script = 86
	// Script_Takri is a representation of the C type PANGO_SCRIPT_TAKRI.
	Script_Takri Script = 87
	// Script_BassaVah is a representation of the C type PANGO_SCRIPT_BASSA_VAH.
	Script_BassaVah Script = 88
	// Script_CaucasianAlbanian is a representation of the C type PANGO_SCRIPT_CAUCASIAN_ALBANIAN.
	Script_CaucasianAlbanian Script = 89
	// Script_Duployan is a representation of the C type PANGO_SCRIPT_DUPLOYAN.
	Script_Duployan Script = 90
	// Script_Elbasan is a representation of the C type PANGO_SCRIPT_ELBASAN.
	Script_Elbasan Script = 91
	// Script_Grantha is a representation of the C type PANGO_SCRIPT_GRANTHA.
	Script_Grantha Script = 92
	// Script_Khojki is a representation of the C type PANGO_SCRIPT_KHOJKI.
	Script_Khojki Script = 93
	// Script_Khudawadi is a representation of the C type PANGO_SCRIPT_KHUDAWADI.
	Script_Khudawadi Script = 94
	// Script_LinearA is a representation of the C type PANGO_SCRIPT_LINEAR_A.
	Script_LinearA Script = 95
	// Script_Mahajani is a representation of the C type PANGO_SCRIPT_MAHAJANI.
	Script_Mahajani Script = 96
	// Script_Manichaean is a representation of the C type PANGO_SCRIPT_MANICHAEAN.
	Script_Manichaean Script = 97
	// Script_MendeKikakui is a representation of the C type PANGO_SCRIPT_MENDE_KIKAKUI.
	Script_MendeKikakui Script = 98
	// Script_Modi is a representation of the C type PANGO_SCRIPT_MODI.
	Script_Modi Script = 99
	// Script_Mro is a representation of the C type PANGO_SCRIPT_MRO.
	Script_Mro Script = 100
	// Script_Nabataean is a representation of the C type PANGO_SCRIPT_NABATAEAN.
	Script_Nabataean Script = 101
	// Script_OldNorthArabian is a representation of the C type PANGO_SCRIPT_OLD_NORTH_ARABIAN.
	Script_OldNorthArabian Script = 102
	// Script_OldPermic is a representation of the C type PANGO_SCRIPT_OLD_PERMIC.
	Script_OldPermic Script = 103
	// Script_PahawhHmong is a representation of the C type PANGO_SCRIPT_PAHAWH_HMONG.
	Script_PahawhHmong Script = 104
	// Script_Palmyrene is a representation of the C type PANGO_SCRIPT_PALMYRENE.
	Script_Palmyrene Script = 105
	// Script_PauCinHau is a representation of the C type PANGO_SCRIPT_PAU_CIN_HAU.
	Script_PauCinHau Script = 106
	// Script_PsalterPahlavi is a representation of the C type PANGO_SCRIPT_PSALTER_PAHLAVI.
	Script_PsalterPahlavi Script = 107
	// Script_Siddham is a representation of the C type PANGO_SCRIPT_SIDDHAM.
	Script_Siddham Script = 108
	// Script_Tirhuta is a representation of the C type PANGO_SCRIPT_TIRHUTA.
	Script_Tirhuta Script = 109
	// Script_WarangCiti is a representation of the C type PANGO_SCRIPT_WARANG_CITI.
	Script_WarangCiti Script = 110
	// Script_Ahom is a representation of the C type PANGO_SCRIPT_AHOM.
	Script_Ahom Script = 111
	// Script_AnatolianHieroglyphs is a representation of the C type PANGO_SCRIPT_ANATOLIAN_HIEROGLYPHS.
	Script_AnatolianHieroglyphs Script = 112
	// Script_Hatran is a representation of the C type PANGO_SCRIPT_HATRAN.
	Script_Hatran Script = 113
	// Script_Multani is a representation of the C type PANGO_SCRIPT_MULTANI.
	Script_Multani Script = 114
	// Script_OldHungarian is a representation of the C type PANGO_SCRIPT_OLD_HUNGARIAN.
	Script_OldHungarian Script = 115
	// Script_Signwriting is a representation of the C type PANGO_SCRIPT_SIGNWRITING.
	Script_Signwriting Script = 116
)

// Stretch is a representation of the C type PangoStretch.
type Stretch int

const (
	// Stretch_UltraCondensed is a representation of the C type PANGO_STRETCH_ULTRA_CONDENSED.
	Stretch_UltraCondensed Stretch = 0
	// Stretch_ExtraCondensed is a representation of the C type PANGO_STRETCH_EXTRA_CONDENSED.
	Stretch_ExtraCondensed Stretch = 1
	// Stretch_Condensed is a representation of the C type PANGO_STRETCH_CONDENSED.
	Stretch_Condensed Stretch = 2
	// Stretch_SemiCondensed is a representation of the C type PANGO_STRETCH_SEMI_CONDENSED.
	Stretch_SemiCondensed Stretch = 3
	// Stretch_Normal is a representation of the C type PANGO_STRETCH_NORMAL.
	Stretch_Normal Stretch = 4
	// Stretch_SemiExpanded is a representation of the C type PANGO_STRETCH_SEMI_EXPANDED.
	Stretch_SemiExpanded Stretch = 5
	// Stretch_Expanded is a representation of the C type PANGO_STRETCH_EXPANDED.
	Stretch_Expanded Stretch = 6
	// Stretch_ExtraExpanded is a representation of the C type PANGO_STRETCH_EXTRA_EXPANDED.
	Stretch_ExtraExpanded Stretch = 7
	// Stretch_UltraExpanded is a representation of the C type PANGO_STRETCH_ULTRA_EXPANDED.
	Stretch_UltraExpanded Stretch = 8
)

// Style is a representation of the C type PangoStyle.
type Style int

const (
	// Style_Normal is a representation of the C type PANGO_STYLE_NORMAL.
	Style_Normal Style = 0
	// Style_Oblique is a representation of the C type PANGO_STYLE_OBLIQUE.
	Style_Oblique Style = 1
	// Style_Italic is a representation of the C type PANGO_STYLE_ITALIC.
	Style_Italic Style = 2
)

// Tabalign is a representation of the C type PangoTabAlign.
type Tabalign int

const (
	// Tabalign_Left is a representation of the C type PANGO_TAB_LEFT.
	Tabalign_Left Tabalign = 0
)

// Underline is a representation of the C type PangoUnderline.
type Underline int

const (
	// Underline_None is a representation of the C type PANGO_UNDERLINE_NONE.
	Underline_None Underline = 0
	// Underline_Single is a representation of the C type PANGO_UNDERLINE_SINGLE.
	Underline_Single Underline = 1
	// Underline_Double is a representation of the C type PANGO_UNDERLINE_DOUBLE.
	Underline_Double Underline = 2
	// Underline_Low is a representation of the C type PANGO_UNDERLINE_LOW.
	Underline_Low Underline = 3
	// Underline_Error is a representation of the C type PANGO_UNDERLINE_ERROR.
	Underline_Error Underline = 4
)

// Variant is a representation of the C type PangoVariant.
type Variant int

const (
	// Variant_Normal is a representation of the C type PANGO_VARIANT_NORMAL.
	Variant_Normal Variant = 0
	// Variant_SmallCaps is a representation of the C type PANGO_VARIANT_SMALL_CAPS.
	Variant_SmallCaps Variant = 1
)

// Weight is a representation of the C type PangoWeight.
type Weight int

const (
	// Weight_Thin is a representation of the C type PANGO_WEIGHT_THIN.
	Weight_Thin Weight = 100
	// Weight_Ultralight is a representation of the C type PANGO_WEIGHT_ULTRALIGHT.
	Weight_Ultralight Weight = 200
	// Weight_Light is a representation of the C type PANGO_WEIGHT_LIGHT.
	Weight_Light Weight = 300
	// Weight_Semilight is a representation of the C type PANGO_WEIGHT_SEMILIGHT.
	Weight_Semilight Weight = 350
	// Weight_Book is a representation of the C type PANGO_WEIGHT_BOOK.
	Weight_Book Weight = 380
	// Weight_Normal is a representation of the C type PANGO_WEIGHT_NORMAL.
	Weight_Normal Weight = 400
	// Weight_Medium is a representation of the C type PANGO_WEIGHT_MEDIUM.
	Weight_Medium Weight = 500
	// Weight_Semibold is a representation of the C type PANGO_WEIGHT_SEMIBOLD.
	Weight_Semibold Weight = 600
	// Weight_Bold is a representation of the C type PANGO_WEIGHT_BOLD.
	Weight_Bold Weight = 700
	// Weight_Ultrabold is a representation of the C type PANGO_WEIGHT_ULTRABOLD.
	Weight_Ultrabold Weight = 800
	// Weight_Heavy is a representation of the C type PANGO_WEIGHT_HEAVY.
	Weight_Heavy Weight = 900
	// Weight_Ultraheavy is a representation of the C type PANGO_WEIGHT_ULTRAHEAVY.
	Weight_Ultraheavy Weight = 1000
)

// Wrapmode is a representation of the C type PangoWrapMode.
type Wrapmode int

const (
	// Wrapmode_Word is a representation of the C type PANGO_WRAP_WORD.
	Wrapmode_Word Wrapmode = 0
	// Wrapmode_Char is a representation of the C type PANGO_WRAP_CHAR.
	Wrapmode_Char Wrapmode = 1
	// Wrapmode_WordChar is a representation of the C type PANGO_WRAP_WORD_CHAR.
	Wrapmode_WordChar Wrapmode = 2
)