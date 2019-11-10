// Code generated - DO NOT EDIT.

package pango

type Alignment uint32

const (
	Alignment_Left   Alignment = int64(0)
	Alignment_Center Alignment = int64(1)
	Alignment_Right  Alignment = int64(2)
)

type AttrType uint32

const (
	AttrType_Invalid            AttrType = int64(0)
	AttrType_Language           AttrType = int64(1)
	AttrType_Family             AttrType = int64(2)
	AttrType_Style              AttrType = int64(3)
	AttrType_Weight             AttrType = int64(4)
	AttrType_Variant            AttrType = int64(5)
	AttrType_Stretch            AttrType = int64(6)
	AttrType_Size               AttrType = int64(7)
	AttrType_FontDesc           AttrType = int64(8)
	AttrType_Foreground         AttrType = int64(9)
	AttrType_Background         AttrType = int64(10)
	AttrType_Underline          AttrType = int64(11)
	AttrType_Strikethrough      AttrType = int64(12)
	AttrType_Rise               AttrType = int64(13)
	AttrType_Shape              AttrType = int64(14)
	AttrType_Scale              AttrType = int64(15)
	AttrType_Fallback           AttrType = int64(16)
	AttrType_LetterSpacing      AttrType = int64(17)
	AttrType_UnderlineColor     AttrType = int64(18)
	AttrType_StrikethroughColor AttrType = int64(19)
	AttrType_AbsoluteSize       AttrType = int64(20)
	AttrType_Gravity            AttrType = int64(21)
	AttrType_GravityHint        AttrType = int64(22)
	AttrType_FontFeatures       AttrType = int64(23)
	AttrType_ForegroundAlpha    AttrType = int64(24)
	AttrType_BackgroundAlpha    AttrType = int64(25)
)

type BidiType uint32

const (
	BidiType_L   BidiType = int64(0)
	BidiType_Lre BidiType = int64(1)
	BidiType_Lro BidiType = int64(2)
	BidiType_R   BidiType = int64(3)
	BidiType_Al  BidiType = int64(4)
	BidiType_Rle BidiType = int64(5)
	BidiType_Rlo BidiType = int64(6)
	BidiType_Pdf BidiType = int64(7)
	BidiType_En  BidiType = int64(8)
	BidiType_Es  BidiType = int64(9)
	BidiType_Et  BidiType = int64(10)
	BidiType_An  BidiType = int64(11)
	BidiType_Cs  BidiType = int64(12)
	BidiType_Nsm BidiType = int64(13)
	BidiType_Bn  BidiType = int64(14)
	BidiType_B   BidiType = int64(15)
	BidiType_S   BidiType = int64(16)
	BidiType_Ws  BidiType = int64(17)
	BidiType_On  BidiType = int64(18)
)

type CoverageLevel uint32

const (
	CoverageLevel_None        CoverageLevel = int64(0)
	CoverageLevel_Fallback    CoverageLevel = int64(1)
	CoverageLevel_Approximate CoverageLevel = int64(2)
	CoverageLevel_Exact       CoverageLevel = int64(3)
)

type Direction uint32

const (
	Direction_Ltr     Direction = int64(0)
	Direction_Rtl     Direction = int64(1)
	Direction_TtbLtr  Direction = int64(2)
	Direction_TtbRtl  Direction = int64(3)
	Direction_WeakLtr Direction = int64(4)
	Direction_WeakRtl Direction = int64(5)
	Direction_Neutral Direction = int64(6)
)

type EllipsizeMode uint32

const (
	EllipsizeMode_None   EllipsizeMode = int64(0)
	EllipsizeMode_Start  EllipsizeMode = int64(1)
	EllipsizeMode_Middle EllipsizeMode = int64(2)
	EllipsizeMode_End    EllipsizeMode = int64(3)
)

type Gravity uint32

const (
	Gravity_South Gravity = int64(0)
	Gravity_East  Gravity = int64(1)
	Gravity_North Gravity = int64(2)
	Gravity_West  Gravity = int64(3)
	Gravity_Auto  Gravity = int64(4)
)

type GravityHint uint32

const (
	GravityHint_Natural GravityHint = int64(0)
	GravityHint_Strong  GravityHint = int64(1)
	GravityHint_Line    GravityHint = int64(2)
)

type RenderPart uint32

const (
	RenderPart_Foreground    RenderPart = int64(0)
	RenderPart_Background    RenderPart = int64(1)
	RenderPart_Underline     RenderPart = int64(2)
	RenderPart_Strikethrough RenderPart = int64(3)
)

type Script int32

const (
	Script_InvalidCode          Script = int64(-1)
	Script_Common               Script = int64(0)
	Script_Inherited            Script = int64(1)
	Script_Arabic               Script = int64(2)
	Script_Armenian             Script = int64(3)
	Script_Bengali              Script = int64(4)
	Script_Bopomofo             Script = int64(5)
	Script_Cherokee             Script = int64(6)
	Script_Coptic               Script = int64(7)
	Script_Cyrillic             Script = int64(8)
	Script_Deseret              Script = int64(9)
	Script_Devanagari           Script = int64(10)
	Script_Ethiopic             Script = int64(11)
	Script_Georgian             Script = int64(12)
	Script_Gothic               Script = int64(13)
	Script_Greek                Script = int64(14)
	Script_Gujarati             Script = int64(15)
	Script_Gurmukhi             Script = int64(16)
	Script_Han                  Script = int64(17)
	Script_Hangul               Script = int64(18)
	Script_Hebrew               Script = int64(19)
	Script_Hiragana             Script = int64(20)
	Script_Kannada              Script = int64(21)
	Script_Katakana             Script = int64(22)
	Script_Khmer                Script = int64(23)
	Script_Lao                  Script = int64(24)
	Script_Latin                Script = int64(25)
	Script_Malayalam            Script = int64(26)
	Script_Mongolian            Script = int64(27)
	Script_Myanmar              Script = int64(28)
	Script_Ogham                Script = int64(29)
	Script_OldItalic            Script = int64(30)
	Script_Oriya                Script = int64(31)
	Script_Runic                Script = int64(32)
	Script_Sinhala              Script = int64(33)
	Script_Syriac               Script = int64(34)
	Script_Tamil                Script = int64(35)
	Script_Telugu               Script = int64(36)
	Script_Thaana               Script = int64(37)
	Script_Thai                 Script = int64(38)
	Script_Tibetan              Script = int64(39)
	Script_CanadianAboriginal   Script = int64(40)
	Script_Yi                   Script = int64(41)
	Script_Tagalog              Script = int64(42)
	Script_Hanunoo              Script = int64(43)
	Script_Buhid                Script = int64(44)
	Script_Tagbanwa             Script = int64(45)
	Script_Braille              Script = int64(46)
	Script_Cypriot              Script = int64(47)
	Script_Limbu                Script = int64(48)
	Script_Osmanya              Script = int64(49)
	Script_Shavian              Script = int64(50)
	Script_LinearB              Script = int64(51)
	Script_TaiLe                Script = int64(52)
	Script_Ugaritic             Script = int64(53)
	Script_NewTaiLue            Script = int64(54)
	Script_Buginese             Script = int64(55)
	Script_Glagolitic           Script = int64(56)
	Script_Tifinagh             Script = int64(57)
	Script_SylotiNagri          Script = int64(58)
	Script_OldPersian           Script = int64(59)
	Script_Kharoshthi           Script = int64(60)
	Script_Unknown              Script = int64(61)
	Script_Balinese             Script = int64(62)
	Script_Cuneiform            Script = int64(63)
	Script_Phoenician           Script = int64(64)
	Script_PhagsPa              Script = int64(65)
	Script_Nko                  Script = int64(66)
	Script_KayahLi              Script = int64(67)
	Script_Lepcha               Script = int64(68)
	Script_Rejang               Script = int64(69)
	Script_Sundanese            Script = int64(70)
	Script_Saurashtra           Script = int64(71)
	Script_Cham                 Script = int64(72)
	Script_OlChiki              Script = int64(73)
	Script_Vai                  Script = int64(74)
	Script_Carian               Script = int64(75)
	Script_Lycian               Script = int64(76)
	Script_Lydian               Script = int64(77)
	Script_Batak                Script = int64(78)
	Script_Brahmi               Script = int64(79)
	Script_Mandaic              Script = int64(80)
	Script_Chakma               Script = int64(81)
	Script_MeroiticCursive      Script = int64(82)
	Script_MeroiticHieroglyphs  Script = int64(83)
	Script_Miao                 Script = int64(84)
	Script_Sharada              Script = int64(85)
	Script_SoraSompeng          Script = int64(86)
	Script_Takri                Script = int64(87)
	Script_BassaVah             Script = int64(88)
	Script_CaucasianAlbanian    Script = int64(89)
	Script_Duployan             Script = int64(90)
	Script_Elbasan              Script = int64(91)
	Script_Grantha              Script = int64(92)
	Script_Khojki               Script = int64(93)
	Script_Khudawadi            Script = int64(94)
	Script_LinearA              Script = int64(95)
	Script_Mahajani             Script = int64(96)
	Script_Manichaean           Script = int64(97)
	Script_MendeKikakui         Script = int64(98)
	Script_Modi                 Script = int64(99)
	Script_Mro                  Script = int64(100)
	Script_Nabataean            Script = int64(101)
	Script_OldNorthArabian      Script = int64(102)
	Script_OldPermic            Script = int64(103)
	Script_PahawhHmong          Script = int64(104)
	Script_Palmyrene            Script = int64(105)
	Script_PauCinHau            Script = int64(106)
	Script_PsalterPahlavi       Script = int64(107)
	Script_Siddham              Script = int64(108)
	Script_Tirhuta              Script = int64(109)
	Script_WarangCiti           Script = int64(110)
	Script_Ahom                 Script = int64(111)
	Script_AnatolianHieroglyphs Script = int64(112)
	Script_Hatran               Script = int64(113)
	Script_Multani              Script = int64(114)
	Script_OldHungarian         Script = int64(115)
	Script_Signwriting          Script = int64(116)
)

type Stretch uint32

const (
	Stretch_UltraCondensed Stretch = int64(0)
	Stretch_ExtraCondensed Stretch = int64(1)
	Stretch_Condensed      Stretch = int64(2)
	Stretch_SemiCondensed  Stretch = int64(3)
	Stretch_Normal         Stretch = int64(4)
	Stretch_SemiExpanded   Stretch = int64(5)
	Stretch_Expanded       Stretch = int64(6)
	Stretch_ExtraExpanded  Stretch = int64(7)
	Stretch_UltraExpanded  Stretch = int64(8)
)

type Style uint32

const (
	Style_Normal  Style = int64(0)
	Style_Oblique Style = int64(1)
	Style_Italic  Style = int64(2)
)

type TabAlign uint32

const (
	TabAlign_Left TabAlign = int64(0)
)

type Underline uint32

const (
	Underline_None   Underline = int64(0)
	Underline_Single Underline = int64(1)
	Underline_Double Underline = int64(2)
	Underline_Low    Underline = int64(3)
	Underline_Error  Underline = int64(4)
)

type Variant uint32

const (
	Variant_Normal    Variant = int64(0)
	Variant_SmallCaps Variant = int64(1)
)

type Weight uint32

const (
	Weight_Thin       Weight = int64(100)
	Weight_Ultralight Weight = int64(200)
	Weight_Light      Weight = int64(300)
	Weight_Semilight  Weight = int64(350)
	Weight_Book       Weight = int64(380)
	Weight_Normal     Weight = int64(400)
	Weight_Medium     Weight = int64(500)
	Weight_Semibold   Weight = int64(600)
	Weight_Bold       Weight = int64(700)
	Weight_Ultrabold  Weight = int64(800)
	Weight_Heavy      Weight = int64(900)
	Weight_Ultraheavy Weight = int64(1000)
)

type WrapMode uint32

const (
	WrapMode_Word     WrapMode = int64(0)
	WrapMode_Char     WrapMode = int64(1)
	WrapMode_WordChar WrapMode = int64(2)
)
