// Code generated - DO NOT EDIT.

package pango

// Alignment is a representation of the C type Alignment.
type Alignment int

const (
	Alignment_Left   Alignment = 0
	Alignment_Center Alignment = 1
	Alignment_Right  Alignment = 2
)

// Attrtype is a representation of the C type AttrType.
type Attrtype int

const (
	Attrtype_Invalid            Attrtype = 0
	Attrtype_Language           Attrtype = 1
	Attrtype_Family             Attrtype = 2
	Attrtype_Style              Attrtype = 3
	Attrtype_Weight             Attrtype = 4
	Attrtype_Variant            Attrtype = 5
	Attrtype_Stretch            Attrtype = 6
	Attrtype_Size               Attrtype = 7
	Attrtype_FontDesc           Attrtype = 8
	Attrtype_Foreground         Attrtype = 9
	Attrtype_Background         Attrtype = 10
	Attrtype_Underline          Attrtype = 11
	Attrtype_Strikethrough      Attrtype = 12
	Attrtype_Rise               Attrtype = 13
	Attrtype_Shape              Attrtype = 14
	Attrtype_Scale              Attrtype = 15
	Attrtype_Fallback           Attrtype = 16
	Attrtype_LetterSpacing      Attrtype = 17
	Attrtype_UnderlineColor     Attrtype = 18
	Attrtype_StrikethroughColor Attrtype = 19
	Attrtype_AbsoluteSize       Attrtype = 20
	Attrtype_Gravity            Attrtype = 21
	Attrtype_GravityHint        Attrtype = 22
	Attrtype_FontFeatures       Attrtype = 23
	Attrtype_ForegroundAlpha    Attrtype = 24
	Attrtype_BackgroundAlpha    Attrtype = 25
)

// Biditype is a representation of the C type BidiType.
//
// since 1.22
type Biditype int

const (
	Biditype_L   Biditype = 0
	Biditype_Lre Biditype = 1
	Biditype_Lro Biditype = 2
	Biditype_R   Biditype = 3
	Biditype_Al  Biditype = 4
	Biditype_Rle Biditype = 5
	Biditype_Rlo Biditype = 6
	Biditype_Pdf Biditype = 7
	Biditype_En  Biditype = 8
	Biditype_Es  Biditype = 9
	Biditype_Et  Biditype = 10
	Biditype_An  Biditype = 11
	Biditype_Cs  Biditype = 12
	Biditype_Nsm Biditype = 13
	Biditype_Bn  Biditype = 14
	Biditype_B   Biditype = 15
	Biditype_S   Biditype = 16
	Biditype_Ws  Biditype = 17
	Biditype_On  Biditype = 18
)

// Coveragelevel is a representation of the C type CoverageLevel.
type Coveragelevel int

const (
	Coveragelevel_None        Coveragelevel = 0
	Coveragelevel_Fallback    Coveragelevel = 1
	Coveragelevel_Approximate Coveragelevel = 2
	Coveragelevel_Exact       Coveragelevel = 3
)

// Direction is a representation of the C type Direction.
type Direction int

const (
	Direction_Ltr     Direction = 0
	Direction_Rtl     Direction = 1
	Direction_TtbLtr  Direction = 2
	Direction_TtbRtl  Direction = 3
	Direction_WeakLtr Direction = 4
	Direction_WeakRtl Direction = 5
	Direction_Neutral Direction = 6
)

// Ellipsizemode is a representation of the C type EllipsizeMode.
type Ellipsizemode int

const (
	Ellipsizemode_None   Ellipsizemode = 0
	Ellipsizemode_Start  Ellipsizemode = 1
	Ellipsizemode_Middle Ellipsizemode = 2
	Ellipsizemode_End    Ellipsizemode = 3
)

// Gravity is a representation of the C type Gravity.
//
// since 1.16
type Gravity int

const (
	Gravity_South Gravity = 0
	Gravity_East  Gravity = 1
	Gravity_North Gravity = 2
	Gravity_West  Gravity = 3
	Gravity_Auto  Gravity = 4
)

// Gravityhint is a representation of the C type GravityHint.
//
// since 1.16
type Gravityhint int

const (
	Gravityhint_Natural Gravityhint = 0
	Gravityhint_Strong  Gravityhint = 1
	Gravityhint_Line    Gravityhint = 2
)

// Renderpart is a representation of the C type RenderPart.
//
// since 1.8
type Renderpart int

const (
	Renderpart_Foreground    Renderpart = 0
	Renderpart_Background    Renderpart = 1
	Renderpart_Underline     Renderpart = 2
	Renderpart_Strikethrough Renderpart = 3
)

// Script is a representation of the C type Script.
type Script int

const (
	Script_InvalidCode          Script = -1
	Script_Common               Script = 0
	Script_Inherited            Script = 1
	Script_Arabic               Script = 2
	Script_Armenian             Script = 3
	Script_Bengali              Script = 4
	Script_Bopomofo             Script = 5
	Script_Cherokee             Script = 6
	Script_Coptic               Script = 7
	Script_Cyrillic             Script = 8
	Script_Deseret              Script = 9
	Script_Devanagari           Script = 10
	Script_Ethiopic             Script = 11
	Script_Georgian             Script = 12
	Script_Gothic               Script = 13
	Script_Greek                Script = 14
	Script_Gujarati             Script = 15
	Script_Gurmukhi             Script = 16
	Script_Han                  Script = 17
	Script_Hangul               Script = 18
	Script_Hebrew               Script = 19
	Script_Hiragana             Script = 20
	Script_Kannada              Script = 21
	Script_Katakana             Script = 22
	Script_Khmer                Script = 23
	Script_Lao                  Script = 24
	Script_Latin                Script = 25
	Script_Malayalam            Script = 26
	Script_Mongolian            Script = 27
	Script_Myanmar              Script = 28
	Script_Ogham                Script = 29
	Script_OldItalic            Script = 30
	Script_Oriya                Script = 31
	Script_Runic                Script = 32
	Script_Sinhala              Script = 33
	Script_Syriac               Script = 34
	Script_Tamil                Script = 35
	Script_Telugu               Script = 36
	Script_Thaana               Script = 37
	Script_Thai                 Script = 38
	Script_Tibetan              Script = 39
	Script_CanadianAboriginal   Script = 40
	Script_Yi                   Script = 41
	Script_Tagalog              Script = 42
	Script_Hanunoo              Script = 43
	Script_Buhid                Script = 44
	Script_Tagbanwa             Script = 45
	Script_Braille              Script = 46
	Script_Cypriot              Script = 47
	Script_Limbu                Script = 48
	Script_Osmanya              Script = 49
	Script_Shavian              Script = 50
	Script_LinearB              Script = 51
	Script_TaiLe                Script = 52
	Script_Ugaritic             Script = 53
	Script_NewTaiLue            Script = 54
	Script_Buginese             Script = 55
	Script_Glagolitic           Script = 56
	Script_Tifinagh             Script = 57
	Script_SylotiNagri          Script = 58
	Script_OldPersian           Script = 59
	Script_Kharoshthi           Script = 60
	Script_Unknown              Script = 61
	Script_Balinese             Script = 62
	Script_Cuneiform            Script = 63
	Script_Phoenician           Script = 64
	Script_PhagsPa              Script = 65
	Script_Nko                  Script = 66
	Script_KayahLi              Script = 67
	Script_Lepcha               Script = 68
	Script_Rejang               Script = 69
	Script_Sundanese            Script = 70
	Script_Saurashtra           Script = 71
	Script_Cham                 Script = 72
	Script_OlChiki              Script = 73
	Script_Vai                  Script = 74
	Script_Carian               Script = 75
	Script_Lycian               Script = 76
	Script_Lydian               Script = 77
	Script_Batak                Script = 78
	Script_Brahmi               Script = 79
	Script_Mandaic              Script = 80
	Script_Chakma               Script = 81
	Script_MeroiticCursive      Script = 82
	Script_MeroiticHieroglyphs  Script = 83
	Script_Miao                 Script = 84
	Script_Sharada              Script = 85
	Script_SoraSompeng          Script = 86
	Script_Takri                Script = 87
	Script_BassaVah             Script = 88
	Script_CaucasianAlbanian    Script = 89
	Script_Duployan             Script = 90
	Script_Elbasan              Script = 91
	Script_Grantha              Script = 92
	Script_Khojki               Script = 93
	Script_Khudawadi            Script = 94
	Script_LinearA              Script = 95
	Script_Mahajani             Script = 96
	Script_Manichaean           Script = 97
	Script_MendeKikakui         Script = 98
	Script_Modi                 Script = 99
	Script_Mro                  Script = 100
	Script_Nabataean            Script = 101
	Script_OldNorthArabian      Script = 102
	Script_OldPermic            Script = 103
	Script_PahawhHmong          Script = 104
	Script_Palmyrene            Script = 105
	Script_PauCinHau            Script = 106
	Script_PsalterPahlavi       Script = 107
	Script_Siddham              Script = 108
	Script_Tirhuta              Script = 109
	Script_WarangCiti           Script = 110
	Script_Ahom                 Script = 111
	Script_AnatolianHieroglyphs Script = 112
	Script_Hatran               Script = 113
	Script_Multani              Script = 114
	Script_OldHungarian         Script = 115
	Script_Signwriting          Script = 116
)

// Stretch is a representation of the C type Stretch.
type Stretch int

const (
	Stretch_UltraCondensed Stretch = 0
	Stretch_ExtraCondensed Stretch = 1
	Stretch_Condensed      Stretch = 2
	Stretch_SemiCondensed  Stretch = 3
	Stretch_Normal         Stretch = 4
	Stretch_SemiExpanded   Stretch = 5
	Stretch_Expanded       Stretch = 6
	Stretch_ExtraExpanded  Stretch = 7
	Stretch_UltraExpanded  Stretch = 8
)

// Style is a representation of the C type Style.
type Style int

const (
	Style_Normal  Style = 0
	Style_Oblique Style = 1
	Style_Italic  Style = 2
)

// Tabalign is a representation of the C type TabAlign.
type Tabalign int

const (
	Tabalign_Left Tabalign = 0
)

// Underline is a representation of the C type Underline.
type Underline int

const (
	Underline_None   Underline = 0
	Underline_Single Underline = 1
	Underline_Double Underline = 2
	Underline_Low    Underline = 3
	Underline_Error  Underline = 4
)

// Variant is a representation of the C type Variant.
type Variant int

const (
	Variant_Normal    Variant = 0
	Variant_SmallCaps Variant = 1
)

// Weight is a representation of the C type Weight.
type Weight int

const (
	Weight_Thin       Weight = 100
	Weight_Ultralight Weight = 200
	Weight_Light      Weight = 300
	Weight_Semilight  Weight = 350
	Weight_Book       Weight = 380
	Weight_Normal     Weight = 400
	Weight_Medium     Weight = 500
	Weight_Semibold   Weight = 600
	Weight_Bold       Weight = 700
	Weight_Ultrabold  Weight = 800
	Weight_Heavy      Weight = 900
	Weight_Ultraheavy Weight = 1000
)

// Wrapmode is a representation of the C type WrapMode.
type Wrapmode int

const (
	Wrapmode_Word     Wrapmode = 0
	Wrapmode_Char     Wrapmode = 1
	Wrapmode_WordChar Wrapmode = 2
)
