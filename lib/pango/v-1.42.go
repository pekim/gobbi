// Code generated - DO NOT EDIT.
// +build pango_1.42

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	pango "github.com/pekim/gobbi/lib/internal/c/pango"
	"unsafe"
)

// Glyph is a representation of the C alias PangoGlyph.
type Glyph uint32

// GlyphUnit is a representation of the C alias PangoGlyphUnit.
type GlyphUnit int32

// LayoutRun is a representation of the C alias PangoLayoutRun.
type LayoutRun GlyphItem

// ANALYSIS_FLAG_CENTERED_BASELINE is a representation of the C constant PANGO_ANALYSIS_FLAG_CENTERED_BASELINE.
//
// since 1.16
const ANALYSIS_FLAG_CENTERED_BASELINE = 1

// ANALYSIS_FLAG_IS_ELLIPSIS is a representation of the C constant PANGO_ANALYSIS_FLAG_IS_ELLIPSIS.
//
// since 1.36.7
const ANALYSIS_FLAG_IS_ELLIPSIS = 2

// ATTR_INDEX_FROM_TEXT_BEGINNING is a representation of the C constant PANGO_ATTR_INDEX_FROM_TEXT_BEGINNING.
//
// since 1.24
const ATTR_INDEX_FROM_TEXT_BEGINNING = 0

// ENGINE_TYPE_LANG is a representation of the C constant PANGO_ENGINE_TYPE_LANG.
const ENGINE_TYPE_LANG = "PangoEngineLang"

// ENGINE_TYPE_SHAPE is a representation of the C constant PANGO_ENGINE_TYPE_SHAPE.
const ENGINE_TYPE_SHAPE = "PangoEngineShape"

// GLYPH_EMPTY is a representation of the C constant PANGO_GLYPH_EMPTY.
const GLYPH_EMPTY = uint64(0xfffffff)

// GLYPH_INVALID_INPUT is a representation of the C constant PANGO_GLYPH_INVALID_INPUT.
//
// since 1.20
const GLYPH_INVALID_INPUT = uint64(0xffffffff)

// GLYPH_UNKNOWN_FLAG is a representation of the C constant PANGO_GLYPH_UNKNOWN_FLAG.
const GLYPH_UNKNOWN_FLAG = uint64(0x10000000)

// RENDER_TYPE_NONE is a representation of the C constant PANGO_RENDER_TYPE_NONE.
const RENDER_TYPE_NONE = "PangoRenderNone"

// SCALE is a representation of the C constant PANGO_SCALE.
const SCALE = 1024

// UNKNOWN_GLYPH_HEIGHT is a representation of the C constant PANGO_UNKNOWN_GLYPH_HEIGHT.
const UNKNOWN_GLYPH_HEIGHT = 14

// UNKNOWN_GLYPH_WIDTH is a representation of the C constant PANGO_UNKNOWN_GLYPH_WIDTH.
const UNKNOWN_GLYPH_WIDTH = 10

// VERSION_MIN_REQUIRED is a representation of the C constant PANGO_VERSION_MIN_REQUIRED.
//
// since 1.42
const VERSION_MIN_REQUIRED = 2

// FontMask is a representation of the C bitfield PangoFontMask.
type FontMask int

// FontMask_family is a representation of the C bitfield member PANGO_FONT_MASK_FAMILY.
const FontMask_family = FontMask(1)

// FontMask_style is a representation of the C bitfield member PANGO_FONT_MASK_STYLE.
const FontMask_style = FontMask(2)

// FontMask_variant is a representation of the C bitfield member PANGO_FONT_MASK_VARIANT.
const FontMask_variant = FontMask(4)

// FontMask_weight is a representation of the C bitfield member PANGO_FONT_MASK_WEIGHT.
const FontMask_weight = FontMask(8)

// FontMask_stretch is a representation of the C bitfield member PANGO_FONT_MASK_STRETCH.
const FontMask_stretch = FontMask(16)

// FontMask_size is a representation of the C bitfield member PANGO_FONT_MASK_SIZE.
const FontMask_size = FontMask(32)

// FontMask_gravity is a representation of the C bitfield member PANGO_FONT_MASK_GRAVITY.
const FontMask_gravity = FontMask(64)

// FontMask_variations is a representation of the C bitfield member PANGO_FONT_MASK_VARIATIONS.
const FontMask_variations = FontMask(128)

// Alignment is a representation of the C enumeration PangoAlignment.
type Alignment int

// Alignment_left is a representation of the C enumeration member PANGO_ALIGN_LEFT.
const Alignment_left = Alignment(0)

// Alignment_center is a representation of the C enumeration member PANGO_ALIGN_CENTER.
const Alignment_center = Alignment(1)

// Alignment_right is a representation of the C enumeration member PANGO_ALIGN_RIGHT.
const Alignment_right = Alignment(2)

// AttrType is a representation of the C enumeration PangoAttrType.
type AttrType int

// AttrType_invalid is a representation of the C enumeration member PANGO_ATTR_INVALID.
const AttrType_invalid = AttrType(0)

// AttrType_language is a representation of the C enumeration member PANGO_ATTR_LANGUAGE.
const AttrType_language = AttrType(1)

// AttrType_family is a representation of the C enumeration member PANGO_ATTR_FAMILY.
const AttrType_family = AttrType(2)

// AttrType_style is a representation of the C enumeration member PANGO_ATTR_STYLE.
const AttrType_style = AttrType(3)

// AttrType_weight is a representation of the C enumeration member PANGO_ATTR_WEIGHT.
const AttrType_weight = AttrType(4)

// AttrType_variant is a representation of the C enumeration member PANGO_ATTR_VARIANT.
const AttrType_variant = AttrType(5)

// AttrType_stretch is a representation of the C enumeration member PANGO_ATTR_STRETCH.
const AttrType_stretch = AttrType(6)

// AttrType_size is a representation of the C enumeration member PANGO_ATTR_SIZE.
const AttrType_size = AttrType(7)

// AttrType_font_desc is a representation of the C enumeration member PANGO_ATTR_FONT_DESC.
const AttrType_font_desc = AttrType(8)

// AttrType_foreground is a representation of the C enumeration member PANGO_ATTR_FOREGROUND.
const AttrType_foreground = AttrType(9)

// AttrType_background is a representation of the C enumeration member PANGO_ATTR_BACKGROUND.
const AttrType_background = AttrType(10)

// AttrType_underline is a representation of the C enumeration member PANGO_ATTR_UNDERLINE.
const AttrType_underline = AttrType(11)

// AttrType_strikethrough is a representation of the C enumeration member PANGO_ATTR_STRIKETHROUGH.
const AttrType_strikethrough = AttrType(12)

// AttrType_rise is a representation of the C enumeration member PANGO_ATTR_RISE.
const AttrType_rise = AttrType(13)

// AttrType_shape is a representation of the C enumeration member PANGO_ATTR_SHAPE.
const AttrType_shape = AttrType(14)

// AttrType_scale is a representation of the C enumeration member PANGO_ATTR_SCALE.
const AttrType_scale = AttrType(15)

// AttrType_fallback is a representation of the C enumeration member PANGO_ATTR_FALLBACK.
const AttrType_fallback = AttrType(16)

// AttrType_letter_spacing is a representation of the C enumeration member PANGO_ATTR_LETTER_SPACING.
const AttrType_letter_spacing = AttrType(17)

// AttrType_underline_color is a representation of the C enumeration member PANGO_ATTR_UNDERLINE_COLOR.
const AttrType_underline_color = AttrType(18)

// AttrType_strikethrough_color is a representation of the C enumeration member PANGO_ATTR_STRIKETHROUGH_COLOR.
const AttrType_strikethrough_color = AttrType(19)

// AttrType_absolute_size is a representation of the C enumeration member PANGO_ATTR_ABSOLUTE_SIZE.
const AttrType_absolute_size = AttrType(20)

// AttrType_gravity is a representation of the C enumeration member PANGO_ATTR_GRAVITY.
const AttrType_gravity = AttrType(21)

// AttrType_gravity_hint is a representation of the C enumeration member PANGO_ATTR_GRAVITY_HINT.
const AttrType_gravity_hint = AttrType(22)

// AttrType_font_features is a representation of the C enumeration member PANGO_ATTR_FONT_FEATURES.
const AttrType_font_features = AttrType(23)

// AttrType_foreground_alpha is a representation of the C enumeration member PANGO_ATTR_FOREGROUND_ALPHA.
const AttrType_foreground_alpha = AttrType(24)

// AttrType_background_alpha is a representation of the C enumeration member PANGO_ATTR_BACKGROUND_ALPHA.
const AttrType_background_alpha = AttrType(25)

// BidiType is a representation of the C enumeration PangoBidiType.
type BidiType int

// BidiType_l is a representation of the C enumeration member PANGO_BIDI_TYPE_L.
const BidiType_l = BidiType(0)

// BidiType_lre is a representation of the C enumeration member PANGO_BIDI_TYPE_LRE.
const BidiType_lre = BidiType(1)

// BidiType_lro is a representation of the C enumeration member PANGO_BIDI_TYPE_LRO.
const BidiType_lro = BidiType(2)

// BidiType_r is a representation of the C enumeration member PANGO_BIDI_TYPE_R.
const BidiType_r = BidiType(3)

// BidiType_al is a representation of the C enumeration member PANGO_BIDI_TYPE_AL.
const BidiType_al = BidiType(4)

// BidiType_rle is a representation of the C enumeration member PANGO_BIDI_TYPE_RLE.
const BidiType_rle = BidiType(5)

// BidiType_rlo is a representation of the C enumeration member PANGO_BIDI_TYPE_RLO.
const BidiType_rlo = BidiType(6)

// BidiType_pdf is a representation of the C enumeration member PANGO_BIDI_TYPE_PDF.
const BidiType_pdf = BidiType(7)

// BidiType_en is a representation of the C enumeration member PANGO_BIDI_TYPE_EN.
const BidiType_en = BidiType(8)

// BidiType_es is a representation of the C enumeration member PANGO_BIDI_TYPE_ES.
const BidiType_es = BidiType(9)

// BidiType_et is a representation of the C enumeration member PANGO_BIDI_TYPE_ET.
const BidiType_et = BidiType(10)

// BidiType_an is a representation of the C enumeration member PANGO_BIDI_TYPE_AN.
const BidiType_an = BidiType(11)

// BidiType_cs is a representation of the C enumeration member PANGO_BIDI_TYPE_CS.
const BidiType_cs = BidiType(12)

// BidiType_nsm is a representation of the C enumeration member PANGO_BIDI_TYPE_NSM.
const BidiType_nsm = BidiType(13)

// BidiType_bn is a representation of the C enumeration member PANGO_BIDI_TYPE_BN.
const BidiType_bn = BidiType(14)

// BidiType_b is a representation of the C enumeration member PANGO_BIDI_TYPE_B.
const BidiType_b = BidiType(15)

// BidiType_s is a representation of the C enumeration member PANGO_BIDI_TYPE_S.
const BidiType_s = BidiType(16)

// BidiType_ws is a representation of the C enumeration member PANGO_BIDI_TYPE_WS.
const BidiType_ws = BidiType(17)

// BidiType_on is a representation of the C enumeration member PANGO_BIDI_TYPE_ON.
const BidiType_on = BidiType(18)

// CoverageLevel is a representation of the C enumeration PangoCoverageLevel.
type CoverageLevel int

// CoverageLevel_none is a representation of the C enumeration member PANGO_COVERAGE_NONE.
const CoverageLevel_none = CoverageLevel(0)

// CoverageLevel_fallback is a representation of the C enumeration member PANGO_COVERAGE_FALLBACK.
const CoverageLevel_fallback = CoverageLevel(1)

// CoverageLevel_approximate is a representation of the C enumeration member PANGO_COVERAGE_APPROXIMATE.
const CoverageLevel_approximate = CoverageLevel(2)

// CoverageLevel_exact is a representation of the C enumeration member PANGO_COVERAGE_EXACT.
const CoverageLevel_exact = CoverageLevel(3)

// Direction is a representation of the C enumeration PangoDirection.
type Direction int

// Direction_ltr is a representation of the C enumeration member PANGO_DIRECTION_LTR.
const Direction_ltr = Direction(0)

// Direction_rtl is a representation of the C enumeration member PANGO_DIRECTION_RTL.
const Direction_rtl = Direction(1)

// Direction_ttb_ltr is a representation of the C enumeration member PANGO_DIRECTION_TTB_LTR.
const Direction_ttb_ltr = Direction(2)

// Direction_ttb_rtl is a representation of the C enumeration member PANGO_DIRECTION_TTB_RTL.
const Direction_ttb_rtl = Direction(3)

// Direction_weak_ltr is a representation of the C enumeration member PANGO_DIRECTION_WEAK_LTR.
const Direction_weak_ltr = Direction(4)

// Direction_weak_rtl is a representation of the C enumeration member PANGO_DIRECTION_WEAK_RTL.
const Direction_weak_rtl = Direction(5)

// Direction_neutral is a representation of the C enumeration member PANGO_DIRECTION_NEUTRAL.
const Direction_neutral = Direction(6)

// EllipsizeMode is a representation of the C enumeration PangoEllipsizeMode.
type EllipsizeMode int

// EllipsizeMode_none is a representation of the C enumeration member PANGO_ELLIPSIZE_NONE.
const EllipsizeMode_none = EllipsizeMode(0)

// EllipsizeMode_start is a representation of the C enumeration member PANGO_ELLIPSIZE_START.
const EllipsizeMode_start = EllipsizeMode(1)

// EllipsizeMode_middle is a representation of the C enumeration member PANGO_ELLIPSIZE_MIDDLE.
const EllipsizeMode_middle = EllipsizeMode(2)

// EllipsizeMode_end is a representation of the C enumeration member PANGO_ELLIPSIZE_END.
const EllipsizeMode_end = EllipsizeMode(3)

// Gravity is a representation of the C enumeration PangoGravity.
type Gravity int

// Gravity_south is a representation of the C enumeration member PANGO_GRAVITY_SOUTH.
const Gravity_south = Gravity(0)

// Gravity_east is a representation of the C enumeration member PANGO_GRAVITY_EAST.
const Gravity_east = Gravity(1)

// Gravity_north is a representation of the C enumeration member PANGO_GRAVITY_NORTH.
const Gravity_north = Gravity(2)

// Gravity_west is a representation of the C enumeration member PANGO_GRAVITY_WEST.
const Gravity_west = Gravity(3)

// Gravity_auto is a representation of the C enumeration member PANGO_GRAVITY_AUTO.
const Gravity_auto = Gravity(4)

// GravityHint is a representation of the C enumeration PangoGravityHint.
type GravityHint int

// GravityHint_natural is a representation of the C enumeration member PANGO_GRAVITY_HINT_NATURAL.
const GravityHint_natural = GravityHint(0)

// GravityHint_strong is a representation of the C enumeration member PANGO_GRAVITY_HINT_STRONG.
const GravityHint_strong = GravityHint(1)

// GravityHint_line is a representation of the C enumeration member PANGO_GRAVITY_HINT_LINE.
const GravityHint_line = GravityHint(2)

// RenderPart is a representation of the C enumeration PangoRenderPart.
type RenderPart int

// RenderPart_foreground is a representation of the C enumeration member PANGO_RENDER_PART_FOREGROUND.
const RenderPart_foreground = RenderPart(0)

// RenderPart_background is a representation of the C enumeration member PANGO_RENDER_PART_BACKGROUND.
const RenderPart_background = RenderPart(1)

// RenderPart_underline is a representation of the C enumeration member PANGO_RENDER_PART_UNDERLINE.
const RenderPart_underline = RenderPart(2)

// RenderPart_strikethrough is a representation of the C enumeration member PANGO_RENDER_PART_STRIKETHROUGH.
const RenderPart_strikethrough = RenderPart(3)

// Script is a representation of the C enumeration PangoScript.
type Script int

// Script_invalid_code is a representation of the C enumeration member PANGO_SCRIPT_INVALID_CODE.
const Script_invalid_code = Script(-1)

// Script_common is a representation of the C enumeration member PANGO_SCRIPT_COMMON.
const Script_common = Script(0)

// Script_inherited is a representation of the C enumeration member PANGO_SCRIPT_INHERITED.
const Script_inherited = Script(1)

// Script_arabic is a representation of the C enumeration member PANGO_SCRIPT_ARABIC.
const Script_arabic = Script(2)

// Script_armenian is a representation of the C enumeration member PANGO_SCRIPT_ARMENIAN.
const Script_armenian = Script(3)

// Script_bengali is a representation of the C enumeration member PANGO_SCRIPT_BENGALI.
const Script_bengali = Script(4)

// Script_bopomofo is a representation of the C enumeration member PANGO_SCRIPT_BOPOMOFO.
const Script_bopomofo = Script(5)

// Script_cherokee is a representation of the C enumeration member PANGO_SCRIPT_CHEROKEE.
const Script_cherokee = Script(6)

// Script_coptic is a representation of the C enumeration member PANGO_SCRIPT_COPTIC.
const Script_coptic = Script(7)

// Script_cyrillic is a representation of the C enumeration member PANGO_SCRIPT_CYRILLIC.
const Script_cyrillic = Script(8)

// Script_deseret is a representation of the C enumeration member PANGO_SCRIPT_DESERET.
const Script_deseret = Script(9)

// Script_devanagari is a representation of the C enumeration member PANGO_SCRIPT_DEVANAGARI.
const Script_devanagari = Script(10)

// Script_ethiopic is a representation of the C enumeration member PANGO_SCRIPT_ETHIOPIC.
const Script_ethiopic = Script(11)

// Script_georgian is a representation of the C enumeration member PANGO_SCRIPT_GEORGIAN.
const Script_georgian = Script(12)

// Script_gothic is a representation of the C enumeration member PANGO_SCRIPT_GOTHIC.
const Script_gothic = Script(13)

// Script_greek is a representation of the C enumeration member PANGO_SCRIPT_GREEK.
const Script_greek = Script(14)

// Script_gujarati is a representation of the C enumeration member PANGO_SCRIPT_GUJARATI.
const Script_gujarati = Script(15)

// Script_gurmukhi is a representation of the C enumeration member PANGO_SCRIPT_GURMUKHI.
const Script_gurmukhi = Script(16)

// Script_han is a representation of the C enumeration member PANGO_SCRIPT_HAN.
const Script_han = Script(17)

// Script_hangul is a representation of the C enumeration member PANGO_SCRIPT_HANGUL.
const Script_hangul = Script(18)

// Script_hebrew is a representation of the C enumeration member PANGO_SCRIPT_HEBREW.
const Script_hebrew = Script(19)

// Script_hiragana is a representation of the C enumeration member PANGO_SCRIPT_HIRAGANA.
const Script_hiragana = Script(20)

// Script_kannada is a representation of the C enumeration member PANGO_SCRIPT_KANNADA.
const Script_kannada = Script(21)

// Script_katakana is a representation of the C enumeration member PANGO_SCRIPT_KATAKANA.
const Script_katakana = Script(22)

// Script_khmer is a representation of the C enumeration member PANGO_SCRIPT_KHMER.
const Script_khmer = Script(23)

// Script_lao is a representation of the C enumeration member PANGO_SCRIPT_LAO.
const Script_lao = Script(24)

// Script_latin is a representation of the C enumeration member PANGO_SCRIPT_LATIN.
const Script_latin = Script(25)

// Script_malayalam is a representation of the C enumeration member PANGO_SCRIPT_MALAYALAM.
const Script_malayalam = Script(26)

// Script_mongolian is a representation of the C enumeration member PANGO_SCRIPT_MONGOLIAN.
const Script_mongolian = Script(27)

// Script_myanmar is a representation of the C enumeration member PANGO_SCRIPT_MYANMAR.
const Script_myanmar = Script(28)

// Script_ogham is a representation of the C enumeration member PANGO_SCRIPT_OGHAM.
const Script_ogham = Script(29)

// Script_old_italic is a representation of the C enumeration member PANGO_SCRIPT_OLD_ITALIC.
const Script_old_italic = Script(30)

// Script_oriya is a representation of the C enumeration member PANGO_SCRIPT_ORIYA.
const Script_oriya = Script(31)

// Script_runic is a representation of the C enumeration member PANGO_SCRIPT_RUNIC.
const Script_runic = Script(32)

// Script_sinhala is a representation of the C enumeration member PANGO_SCRIPT_SINHALA.
const Script_sinhala = Script(33)

// Script_syriac is a representation of the C enumeration member PANGO_SCRIPT_SYRIAC.
const Script_syriac = Script(34)

// Script_tamil is a representation of the C enumeration member PANGO_SCRIPT_TAMIL.
const Script_tamil = Script(35)

// Script_telugu is a representation of the C enumeration member PANGO_SCRIPT_TELUGU.
const Script_telugu = Script(36)

// Script_thaana is a representation of the C enumeration member PANGO_SCRIPT_THAANA.
const Script_thaana = Script(37)

// Script_thai is a representation of the C enumeration member PANGO_SCRIPT_THAI.
const Script_thai = Script(38)

// Script_tibetan is a representation of the C enumeration member PANGO_SCRIPT_TIBETAN.
const Script_tibetan = Script(39)

// Script_canadian_aboriginal is a representation of the C enumeration member PANGO_SCRIPT_CANADIAN_ABORIGINAL.
const Script_canadian_aboriginal = Script(40)

// Script_yi is a representation of the C enumeration member PANGO_SCRIPT_YI.
const Script_yi = Script(41)

// Script_tagalog is a representation of the C enumeration member PANGO_SCRIPT_TAGALOG.
const Script_tagalog = Script(42)

// Script_hanunoo is a representation of the C enumeration member PANGO_SCRIPT_HANUNOO.
const Script_hanunoo = Script(43)

// Script_buhid is a representation of the C enumeration member PANGO_SCRIPT_BUHID.
const Script_buhid = Script(44)

// Script_tagbanwa is a representation of the C enumeration member PANGO_SCRIPT_TAGBANWA.
const Script_tagbanwa = Script(45)

// Script_braille is a representation of the C enumeration member PANGO_SCRIPT_BRAILLE.
const Script_braille = Script(46)

// Script_cypriot is a representation of the C enumeration member PANGO_SCRIPT_CYPRIOT.
const Script_cypriot = Script(47)

// Script_limbu is a representation of the C enumeration member PANGO_SCRIPT_LIMBU.
const Script_limbu = Script(48)

// Script_osmanya is a representation of the C enumeration member PANGO_SCRIPT_OSMANYA.
const Script_osmanya = Script(49)

// Script_shavian is a representation of the C enumeration member PANGO_SCRIPT_SHAVIAN.
const Script_shavian = Script(50)

// Script_linear_b is a representation of the C enumeration member PANGO_SCRIPT_LINEAR_B.
const Script_linear_b = Script(51)

// Script_tai_le is a representation of the C enumeration member PANGO_SCRIPT_TAI_LE.
const Script_tai_le = Script(52)

// Script_ugaritic is a representation of the C enumeration member PANGO_SCRIPT_UGARITIC.
const Script_ugaritic = Script(53)

// Script_new_tai_lue is a representation of the C enumeration member PANGO_SCRIPT_NEW_TAI_LUE.
const Script_new_tai_lue = Script(54)

// Script_buginese is a representation of the C enumeration member PANGO_SCRIPT_BUGINESE.
const Script_buginese = Script(55)

// Script_glagolitic is a representation of the C enumeration member PANGO_SCRIPT_GLAGOLITIC.
const Script_glagolitic = Script(56)

// Script_tifinagh is a representation of the C enumeration member PANGO_SCRIPT_TIFINAGH.
const Script_tifinagh = Script(57)

// Script_syloti_nagri is a representation of the C enumeration member PANGO_SCRIPT_SYLOTI_NAGRI.
const Script_syloti_nagri = Script(58)

// Script_old_persian is a representation of the C enumeration member PANGO_SCRIPT_OLD_PERSIAN.
const Script_old_persian = Script(59)

// Script_kharoshthi is a representation of the C enumeration member PANGO_SCRIPT_KHAROSHTHI.
const Script_kharoshthi = Script(60)

// Script_unknown is a representation of the C enumeration member PANGO_SCRIPT_UNKNOWN.
const Script_unknown = Script(61)

// Script_balinese is a representation of the C enumeration member PANGO_SCRIPT_BALINESE.
const Script_balinese = Script(62)

// Script_cuneiform is a representation of the C enumeration member PANGO_SCRIPT_CUNEIFORM.
const Script_cuneiform = Script(63)

// Script_phoenician is a representation of the C enumeration member PANGO_SCRIPT_PHOENICIAN.
const Script_phoenician = Script(64)

// Script_phags_pa is a representation of the C enumeration member PANGO_SCRIPT_PHAGS_PA.
const Script_phags_pa = Script(65)

// Script_nko is a representation of the C enumeration member PANGO_SCRIPT_NKO.
const Script_nko = Script(66)

// Script_kayah_li is a representation of the C enumeration member PANGO_SCRIPT_KAYAH_LI.
const Script_kayah_li = Script(67)

// Script_lepcha is a representation of the C enumeration member PANGO_SCRIPT_LEPCHA.
const Script_lepcha = Script(68)

// Script_rejang is a representation of the C enumeration member PANGO_SCRIPT_REJANG.
const Script_rejang = Script(69)

// Script_sundanese is a representation of the C enumeration member PANGO_SCRIPT_SUNDANESE.
const Script_sundanese = Script(70)

// Script_saurashtra is a representation of the C enumeration member PANGO_SCRIPT_SAURASHTRA.
const Script_saurashtra = Script(71)

// Script_cham is a representation of the C enumeration member PANGO_SCRIPT_CHAM.
const Script_cham = Script(72)

// Script_ol_chiki is a representation of the C enumeration member PANGO_SCRIPT_OL_CHIKI.
const Script_ol_chiki = Script(73)

// Script_vai is a representation of the C enumeration member PANGO_SCRIPT_VAI.
const Script_vai = Script(74)

// Script_carian is a representation of the C enumeration member PANGO_SCRIPT_CARIAN.
const Script_carian = Script(75)

// Script_lycian is a representation of the C enumeration member PANGO_SCRIPT_LYCIAN.
const Script_lycian = Script(76)

// Script_lydian is a representation of the C enumeration member PANGO_SCRIPT_LYDIAN.
const Script_lydian = Script(77)

// Script_batak is a representation of the C enumeration member PANGO_SCRIPT_BATAK.
const Script_batak = Script(78)

// Script_brahmi is a representation of the C enumeration member PANGO_SCRIPT_BRAHMI.
const Script_brahmi = Script(79)

// Script_mandaic is a representation of the C enumeration member PANGO_SCRIPT_MANDAIC.
const Script_mandaic = Script(80)

// Script_chakma is a representation of the C enumeration member PANGO_SCRIPT_CHAKMA.
const Script_chakma = Script(81)

// Script_meroitic_cursive is a representation of the C enumeration member PANGO_SCRIPT_MEROITIC_CURSIVE.
const Script_meroitic_cursive = Script(82)

// Script_meroitic_hieroglyphs is a representation of the C enumeration member PANGO_SCRIPT_MEROITIC_HIEROGLYPHS.
const Script_meroitic_hieroglyphs = Script(83)

// Script_miao is a representation of the C enumeration member PANGO_SCRIPT_MIAO.
const Script_miao = Script(84)

// Script_sharada is a representation of the C enumeration member PANGO_SCRIPT_SHARADA.
const Script_sharada = Script(85)

// Script_sora_sompeng is a representation of the C enumeration member PANGO_SCRIPT_SORA_SOMPENG.
const Script_sora_sompeng = Script(86)

// Script_takri is a representation of the C enumeration member PANGO_SCRIPT_TAKRI.
const Script_takri = Script(87)

// Script_bassa_vah is a representation of the C enumeration member PANGO_SCRIPT_BASSA_VAH.
const Script_bassa_vah = Script(88)

// Script_caucasian_albanian is a representation of the C enumeration member PANGO_SCRIPT_CAUCASIAN_ALBANIAN.
const Script_caucasian_albanian = Script(89)

// Script_duployan is a representation of the C enumeration member PANGO_SCRIPT_DUPLOYAN.
const Script_duployan = Script(90)

// Script_elbasan is a representation of the C enumeration member PANGO_SCRIPT_ELBASAN.
const Script_elbasan = Script(91)

// Script_grantha is a representation of the C enumeration member PANGO_SCRIPT_GRANTHA.
const Script_grantha = Script(92)

// Script_khojki is a representation of the C enumeration member PANGO_SCRIPT_KHOJKI.
const Script_khojki = Script(93)

// Script_khudawadi is a representation of the C enumeration member PANGO_SCRIPT_KHUDAWADI.
const Script_khudawadi = Script(94)

// Script_linear_a is a representation of the C enumeration member PANGO_SCRIPT_LINEAR_A.
const Script_linear_a = Script(95)

// Script_mahajani is a representation of the C enumeration member PANGO_SCRIPT_MAHAJANI.
const Script_mahajani = Script(96)

// Script_manichaean is a representation of the C enumeration member PANGO_SCRIPT_MANICHAEAN.
const Script_manichaean = Script(97)

// Script_mende_kikakui is a representation of the C enumeration member PANGO_SCRIPT_MENDE_KIKAKUI.
const Script_mende_kikakui = Script(98)

// Script_modi is a representation of the C enumeration member PANGO_SCRIPT_MODI.
const Script_modi = Script(99)

// Script_mro is a representation of the C enumeration member PANGO_SCRIPT_MRO.
const Script_mro = Script(100)

// Script_nabataean is a representation of the C enumeration member PANGO_SCRIPT_NABATAEAN.
const Script_nabataean = Script(101)

// Script_old_north_arabian is a representation of the C enumeration member PANGO_SCRIPT_OLD_NORTH_ARABIAN.
const Script_old_north_arabian = Script(102)

// Script_old_permic is a representation of the C enumeration member PANGO_SCRIPT_OLD_PERMIC.
const Script_old_permic = Script(103)

// Script_pahawh_hmong is a representation of the C enumeration member PANGO_SCRIPT_PAHAWH_HMONG.
const Script_pahawh_hmong = Script(104)

// Script_palmyrene is a representation of the C enumeration member PANGO_SCRIPT_PALMYRENE.
const Script_palmyrene = Script(105)

// Script_pau_cin_hau is a representation of the C enumeration member PANGO_SCRIPT_PAU_CIN_HAU.
const Script_pau_cin_hau = Script(106)

// Script_psalter_pahlavi is a representation of the C enumeration member PANGO_SCRIPT_PSALTER_PAHLAVI.
const Script_psalter_pahlavi = Script(107)

// Script_siddham is a representation of the C enumeration member PANGO_SCRIPT_SIDDHAM.
const Script_siddham = Script(108)

// Script_tirhuta is a representation of the C enumeration member PANGO_SCRIPT_TIRHUTA.
const Script_tirhuta = Script(109)

// Script_warang_citi is a representation of the C enumeration member PANGO_SCRIPT_WARANG_CITI.
const Script_warang_citi = Script(110)

// Script_ahom is a representation of the C enumeration member PANGO_SCRIPT_AHOM.
const Script_ahom = Script(111)

// Script_anatolian_hieroglyphs is a representation of the C enumeration member PANGO_SCRIPT_ANATOLIAN_HIEROGLYPHS.
const Script_anatolian_hieroglyphs = Script(112)

// Script_hatran is a representation of the C enumeration member PANGO_SCRIPT_HATRAN.
const Script_hatran = Script(113)

// Script_multani is a representation of the C enumeration member PANGO_SCRIPT_MULTANI.
const Script_multani = Script(114)

// Script_old_hungarian is a representation of the C enumeration member PANGO_SCRIPT_OLD_HUNGARIAN.
const Script_old_hungarian = Script(115)

// Script_signwriting is a representation of the C enumeration member PANGO_SCRIPT_SIGNWRITING.
const Script_signwriting = Script(116)

// Stretch is a representation of the C enumeration PangoStretch.
type Stretch int

// Stretch_ultra_condensed is a representation of the C enumeration member PANGO_STRETCH_ULTRA_CONDENSED.
const Stretch_ultra_condensed = Stretch(0)

// Stretch_extra_condensed is a representation of the C enumeration member PANGO_STRETCH_EXTRA_CONDENSED.
const Stretch_extra_condensed = Stretch(1)

// Stretch_condensed is a representation of the C enumeration member PANGO_STRETCH_CONDENSED.
const Stretch_condensed = Stretch(2)

// Stretch_semi_condensed is a representation of the C enumeration member PANGO_STRETCH_SEMI_CONDENSED.
const Stretch_semi_condensed = Stretch(3)

// Stretch_normal is a representation of the C enumeration member PANGO_STRETCH_NORMAL.
const Stretch_normal = Stretch(4)

// Stretch_semi_expanded is a representation of the C enumeration member PANGO_STRETCH_SEMI_EXPANDED.
const Stretch_semi_expanded = Stretch(5)

// Stretch_expanded is a representation of the C enumeration member PANGO_STRETCH_EXPANDED.
const Stretch_expanded = Stretch(6)

// Stretch_extra_expanded is a representation of the C enumeration member PANGO_STRETCH_EXTRA_EXPANDED.
const Stretch_extra_expanded = Stretch(7)

// Stretch_ultra_expanded is a representation of the C enumeration member PANGO_STRETCH_ULTRA_EXPANDED.
const Stretch_ultra_expanded = Stretch(8)

// Style is a representation of the C enumeration PangoStyle.
type Style int

// Style_normal is a representation of the C enumeration member PANGO_STYLE_NORMAL.
const Style_normal = Style(0)

// Style_oblique is a representation of the C enumeration member PANGO_STYLE_OBLIQUE.
const Style_oblique = Style(1)

// Style_italic is a representation of the C enumeration member PANGO_STYLE_ITALIC.
const Style_italic = Style(2)

// TabAlign is a representation of the C enumeration PangoTabAlign.
type TabAlign int

// TabAlign_left is a representation of the C enumeration member PANGO_TAB_LEFT.
const TabAlign_left = TabAlign(0)

// Underline is a representation of the C enumeration PangoUnderline.
type Underline int

// Underline_none is a representation of the C enumeration member PANGO_UNDERLINE_NONE.
const Underline_none = Underline(0)

// Underline_single is a representation of the C enumeration member PANGO_UNDERLINE_SINGLE.
const Underline_single = Underline(1)

// Underline_double is a representation of the C enumeration member PANGO_UNDERLINE_DOUBLE.
const Underline_double = Underline(2)

// Underline_low is a representation of the C enumeration member PANGO_UNDERLINE_LOW.
const Underline_low = Underline(3)

// Underline_error is a representation of the C enumeration member PANGO_UNDERLINE_ERROR.
const Underline_error = Underline(4)

// Variant is a representation of the C enumeration PangoVariant.
type Variant int

// Variant_normal is a representation of the C enumeration member PANGO_VARIANT_NORMAL.
const Variant_normal = Variant(0)

// Variant_small_caps is a representation of the C enumeration member PANGO_VARIANT_SMALL_CAPS.
const Variant_small_caps = Variant(1)

// Weight is a representation of the C enumeration PangoWeight.
type Weight int

// Weight_thin is a representation of the C enumeration member PANGO_WEIGHT_THIN.
const Weight_thin = Weight(100)

// Weight_ultralight is a representation of the C enumeration member PANGO_WEIGHT_ULTRALIGHT.
const Weight_ultralight = Weight(200)

// Weight_light is a representation of the C enumeration member PANGO_WEIGHT_LIGHT.
const Weight_light = Weight(300)

// Weight_semilight is a representation of the C enumeration member PANGO_WEIGHT_SEMILIGHT.
const Weight_semilight = Weight(350)

// Weight_book is a representation of the C enumeration member PANGO_WEIGHT_BOOK.
const Weight_book = Weight(380)

// Weight_normal is a representation of the C enumeration member PANGO_WEIGHT_NORMAL.
const Weight_normal = Weight(400)

// Weight_medium is a representation of the C enumeration member PANGO_WEIGHT_MEDIUM.
const Weight_medium = Weight(500)

// Weight_semibold is a representation of the C enumeration member PANGO_WEIGHT_SEMIBOLD.
const Weight_semibold = Weight(600)

// Weight_bold is a representation of the C enumeration member PANGO_WEIGHT_BOLD.
const Weight_bold = Weight(700)

// Weight_ultrabold is a representation of the C enumeration member PANGO_WEIGHT_ULTRABOLD.
const Weight_ultrabold = Weight(800)

// Weight_heavy is a representation of the C enumeration member PANGO_WEIGHT_HEAVY.
const Weight_heavy = Weight(900)

// Weight_ultraheavy is a representation of the C enumeration member PANGO_WEIGHT_ULTRAHEAVY.
const Weight_ultraheavy = Weight(1000)

// WrapMode is a representation of the C enumeration PangoWrapMode.
type WrapMode int

// WrapMode_word is a representation of the C enumeration member PANGO_WRAP_WORD.
const WrapMode_word = WrapMode(0)

// WrapMode_char is a representation of the C enumeration member PANGO_WRAP_CHAR.
const WrapMode_char = WrapMode(1)

// WrapMode_word_char is a representation of the C enumeration member PANGO_WRAP_WORD_CHAR.
const WrapMode_word_char = WrapMode(2)

// AttrBackgroundAlphaNew is analogous to the C function pango_attr_background_alpha_new.
func AttrBackgroundAlphaNew(alpha uint16) *Attribute {
	sys_alpha := alpha
	retSys := pango.Fn_pango_attr_background_alpha_new(sys_alpha)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrBackgroundNew is analogous to the C function pango_attr_background_new.
func AttrBackgroundNew(red uint16, green uint16, blue uint16) *Attribute {
	sys_red := red
	sys_green := green
	sys_blue := blue
	retSys := pango.Fn_pango_attr_background_new(sys_red, sys_green, sys_blue)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrFallbackNew is analogous to the C function pango_attr_fallback_new.
func AttrFallbackNew(enableFallback bool) *Attribute {
	sys_enableFallback := enableFallback
	retSys := pango.Fn_pango_attr_fallback_new(sys_enableFallback)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrFamilyNew is analogous to the C function pango_attr_family_new.
func AttrFamilyNew(family string) *Attribute {
	sys_family := family
	retSys := pango.Fn_pango_attr_family_new(sys_family)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrForegroundAlphaNew is analogous to the C function pango_attr_foreground_alpha_new.
func AttrForegroundAlphaNew(alpha uint16) *Attribute {
	sys_alpha := alpha
	retSys := pango.Fn_pango_attr_foreground_alpha_new(sys_alpha)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrForegroundNew is analogous to the C function pango_attr_foreground_new.
func AttrForegroundNew(red uint16, green uint16, blue uint16) *Attribute {
	sys_red := red
	sys_green := green
	sys_blue := blue
	retSys := pango.Fn_pango_attr_foreground_new(sys_red, sys_green, sys_blue)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrGravityHintNew is analogous to the C function pango_attr_gravity_hint_new.
func AttrGravityHintNew(hint GravityHint) *Attribute {
	sys_hint := (int)(hint)
	retSys := pango.Fn_pango_attr_gravity_hint_new(sys_hint)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrGravityNew is analogous to the C function pango_attr_gravity_new.
func AttrGravityNew(gravity Gravity) *Attribute {
	sys_gravity := (int)(gravity)
	retSys := pango.Fn_pango_attr_gravity_new(sys_gravity)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrLetterSpacingNew is analogous to the C function pango_attr_letter_spacing_new.
func AttrLetterSpacingNew(letterSpacing int) *Attribute {
	sys_letterSpacing := letterSpacing
	retSys := pango.Fn_pango_attr_letter_spacing_new(sys_letterSpacing)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrRiseNew is analogous to the C function pango_attr_rise_new.
func AttrRiseNew(rise int) *Attribute {
	sys_rise := rise
	retSys := pango.Fn_pango_attr_rise_new(sys_rise)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrScaleNew is analogous to the C function pango_attr_scale_new.
func AttrScaleNew(scaleFactor float64) *Attribute {
	sys_scaleFactor := scaleFactor
	retSys := pango.Fn_pango_attr_scale_new(sys_scaleFactor)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrStretchNew is analogous to the C function pango_attr_stretch_new.
func AttrStretchNew(stretch Stretch) *Attribute {
	sys_stretch := (int)(stretch)
	retSys := pango.Fn_pango_attr_stretch_new(sys_stretch)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrStrikethroughColorNew is analogous to the C function pango_attr_strikethrough_color_new.
func AttrStrikethroughColorNew(red uint16, green uint16, blue uint16) *Attribute {
	sys_red := red
	sys_green := green
	sys_blue := blue
	retSys := pango.Fn_pango_attr_strikethrough_color_new(sys_red, sys_green, sys_blue)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrStrikethroughNew is analogous to the C function pango_attr_strikethrough_new.
func AttrStrikethroughNew(strikethrough bool) *Attribute {
	sys_strikethrough := strikethrough
	retSys := pango.Fn_pango_attr_strikethrough_new(sys_strikethrough)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrStyleNew is analogous to the C function pango_attr_style_new.
func AttrStyleNew(style Style) *Attribute {
	sys_style := (int)(style)
	retSys := pango.Fn_pango_attr_style_new(sys_style)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrUnderlineColorNew is analogous to the C function pango_attr_underline_color_new.
func AttrUnderlineColorNew(red uint16, green uint16, blue uint16) *Attribute {
	sys_red := red
	sys_green := green
	sys_blue := blue
	retSys := pango.Fn_pango_attr_underline_color_new(sys_red, sys_green, sys_blue)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrUnderlineNew is analogous to the C function pango_attr_underline_new.
func AttrUnderlineNew(underline Underline) *Attribute {
	sys_underline := (int)(underline)
	retSys := pango.Fn_pango_attr_underline_new(sys_underline)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrVariantNew is analogous to the C function pango_attr_variant_new.
func AttrVariantNew(variant Variant) *Attribute {
	sys_variant := (int)(variant)
	retSys := pango.Fn_pango_attr_variant_new(sys_variant)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrWeightNew is analogous to the C function pango_attr_weight_new.
func AttrWeightNew(weight Weight) *Attribute {
	sys_weight := (int)(weight)
	retSys := pango.Fn_pango_attr_weight_new(sys_weight)
	ret := AttributeNewFromC(retSys)

	return ret
}

// UNSUPPORTED : pango_break : has array param, attrs

// UNSUPPORTED : pango_config_key_get : blacklisted

// UNSUPPORTED : pango_config_key_get_system : blacklisted

// UNSUPPORTED : pango_default_break : blacklisted

// ExtentsToPixels is analogous to the C function pango_extents_to_pixels.
func ExtentsToPixels(inclusive *Rectangle, nearest *Rectangle) {
	sys_inclusive := inclusive.ToC()
	sys_nearest := nearest.ToC()
	pango.Fn_pango_extents_to_pixels(sys_inclusive, sys_nearest)
}

// FindBaseDir is analogous to the C function pango_find_base_dir.
func FindBaseDir(text string, length int) int {
	sys_text := text
	sys_length := length
	retSys := pango.Fn_pango_find_base_dir(sys_text, sys_length)
	ret := retSys

	return ret
}

// UNSUPPORTED : pango_find_map : blacklisted

// UNSUPPORTED : pango_find_paragraph_boundary : has [in]out param, paragraph_delimiter_index

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted

// UNSUPPORTED : pango_get_log_attrs : has array param, log_attrs

// GetMirrorChar is analogous to the C function pango_get_mirror_char.
func GetMirrorChar(ch rune, mirroredCh *rune) bool {
	sys_ch := ch
	sys_mirroredCh := mirroredCh
	retSys := pango.Fn_pango_get_mirror_char(sys_ch, sys_mirroredCh)
	ret := retSys

	return ret
}

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted

// IsZeroWidth is analogous to the C function pango_is_zero_width.
func IsZeroWidth(ch rune) bool {
	sys_ch := ch
	retSys := pango.Fn_pango_is_zero_width(sys_ch)
	ret := retSys

	return ret
}

// Itemize is analogous to the C function pango_itemize.
func Itemize(context *Context, text string, startIndex int, length int, attrs *AttrList, cachedIter *AttrIterator) *glib.List {
	sys_context := context.ToC()
	sys_text := text
	sys_startIndex := startIndex
	sys_length := length
	sys_attrs := attrs.ToC()
	sys_cachedIter := cachedIter.ToC()
	retSys := pango.Fn_pango_itemize(sys_context, sys_text, sys_startIndex, sys_length, sys_attrs, sys_cachedIter)
	ret := glib.ListNewFromC(retSys)

	return ret
}

// ItemizeWithBaseDir is analogous to the C function pango_itemize_with_base_dir.
func ItemizeWithBaseDir(context *Context, baseDir Direction, text string, startIndex int, length int, attrs *AttrList, cachedIter *AttrIterator) *glib.List {
	sys_context := context.ToC()
	sys_baseDir := (int)(baseDir)
	sys_text := text
	sys_startIndex := startIndex
	sys_length := length
	sys_attrs := attrs.ToC()
	sys_cachedIter := cachedIter.ToC()
	retSys := pango.Fn_pango_itemize_with_base_dir(sys_context, sys_baseDir, sys_text, sys_startIndex, sys_length, sys_attrs, sys_cachedIter)
	ret := glib.ListNewFromC(retSys)

	return ret
}

// Log2visGetEmbeddingLevels is analogous to the C function pango_log2vis_get_embedding_levels.
func Log2visGetEmbeddingLevels(text string, length int, pbaseDir *Direction) *uint8 {
	sys_text := text
	sys_length := length
	sys_pbaseDir := (*int)(pbaseDir)
	retSys := pango.Fn_pango_log2vis_get_embedding_levels(sys_text, sys_length, sys_pbaseDir)
	ret := retSys

	return ret
}

// UNSUPPORTED : pango_lookup_aliases : blacklisted

// UNSUPPORTED : pango_markup_parser_finish : throws

// MarkupParserNew is analogous to the C function pango_markup_parser_new.
func MarkupParserNew(accelMarker rune) *glib.MarkupParseContext {
	sys_accelMarker := accelMarker
	retSys := pango.Fn_pango_markup_parser_new(sys_accelMarker)
	ret := glib.MarkupParseContextNewFromC(retSys)

	return ret
}

// UNSUPPORTED : pango_module_register : blacklisted

// UNSUPPORTED : pango_parse_enum : has [in]out param, value

// UNSUPPORTED : pango_parse_markup : throws

// UNSUPPORTED : pango_parse_stretch : has [in]out param, stretch

// UNSUPPORTED : pango_parse_style : has [in]out param, style

// UNSUPPORTED : pango_parse_variant : has [in]out param, variant

// UNSUPPORTED : pango_parse_weight : has [in]out param, weight

// QuantizeLineGeometry is analogous to the C function pango_quantize_line_geometry.
func QuantizeLineGeometry(thickness *int, position *int) {
	sys_thickness := thickness
	sys_position := position
	pango.Fn_pango_quantize_line_geometry(sys_thickness, sys_position)
}

// UNSUPPORTED : pango_read_line : has [in]out param, str

// ReorderItems is analogous to the C function pango_reorder_items.
func ReorderItems(logicalItems *glib.List) *glib.List {
	sys_logicalItems := logicalItems.ToC()
	retSys := pango.Fn_pango_reorder_items(sys_logicalItems)
	ret := glib.ListNewFromC(retSys)

	return ret
}

// UNSUPPORTED : pango_scan_int : has [in]out param, out

// UNSUPPORTED : pango_scan_string : has [in]out param, out

// UNSUPPORTED : pango_scan_word : has [in]out param, out

// Shape is analogous to the C function pango_shape.
func Shape(text string, length int, analysis *Analysis, glyphs *GlyphString) {
	sys_text := text
	sys_length := length
	sys_analysis := analysis.ToC()
	sys_glyphs := glyphs.ToC()
	pango.Fn_pango_shape(sys_text, sys_length, sys_analysis, sys_glyphs)
}

// ShapeFull is analogous to the C function pango_shape_full.
func ShapeFull(itemText string, itemLength int, paragraphText string, paragraphLength int, analysis *Analysis, glyphs *GlyphString) {
	sys_itemText := itemText
	sys_itemLength := itemLength
	sys_paragraphText := paragraphText
	sys_paragraphLength := paragraphLength
	sys_analysis := analysis.ToC()
	sys_glyphs := glyphs.ToC()
	pango.Fn_pango_shape_full(sys_itemText, sys_itemLength, sys_paragraphText, sys_paragraphLength, sys_analysis, sys_glyphs)
}

// SkipSpace is analogous to the C function pango_skip_space.
func SkipSpace(pos *string) bool {
	sys_pos := pos
	retSys := pango.Fn_pango_skip_space(sys_pos)
	ret := retSys

	return ret
}

// UNSUPPORTED : pango_split_file_list : no array length

// TrimString is analogous to the C function pango_trim_string.
func TrimString(str string) string {
	sys_str := str
	retSys := pango.Fn_pango_trim_string(sys_str)
	ret := retSys

	return ret
}

// UnicharDirection is analogous to the C function pango_unichar_direction.
func UnicharDirection(ch rune) int {
	sys_ch := ch
	retSys := pango.Fn_pango_unichar_direction(sys_ch)
	ret := retSys

	return ret
}

// UnitsFromDouble is analogous to the C function pango_units_from_double.
func UnitsFromDouble(d float64) int {
	sys_d := d
	retSys := pango.Fn_pango_units_from_double(sys_d)
	ret := retSys

	return ret
}

// UnitsToDouble is analogous to the C function pango_units_to_double.
func UnitsToDouble(i int) float64 {
	sys_i := i
	retSys := pango.Fn_pango_units_to_double(sys_i)
	ret := retSys

	return ret
}

// Version is analogous to the C function pango_version.
func Version() int {
	retSys := pango.Fn_pango_version()
	ret := retSys

	return ret
}

// VersionCheck is analogous to the C function pango_version_check.
func VersionCheck(requiredMajor int, requiredMinor int, requiredMicro int) string {
	sys_requiredMajor := requiredMajor
	sys_requiredMinor := requiredMinor
	sys_requiredMicro := requiredMicro
	retSys := pango.Fn_pango_version_check(sys_requiredMajor, sys_requiredMinor, sys_requiredMicro)
	ret := retSys

	return ret
}

// VersionString is analogous to the C function pango_version_string.
func VersionString() string {
	retSys := pango.Fn_pango_version_string()
	ret := retSys

	return ret
}

// Analysis is a representation of the C record PangoAnalysis.
type Analysis struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAnalysis that represents the Analysis.
func (recv *Analysis) ToC() unsafe.Pointer {
	return recv.native
}

// AnalysisNewFromC creates a new Analysis from a pointer to the C PangoAnalysis that represents the Analysis.
func AnalysisNewFromC(native unsafe.Pointer) *Analysis {
	return &Analysis{native: native}
}

// AttrClass is a representation of the C record PangoAttrClass.
type AttrClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrClass that represents the AttrClass.
func (recv *AttrClass) ToC() unsafe.Pointer {
	return recv.native
}

// AttrClassNewFromC creates a new AttrClass from a pointer to the C PangoAttrClass that represents the AttrClass.
func AttrClassNewFromC(native unsafe.Pointer) *AttrClass {
	return &AttrClass{native: native}
}

// AttrColor is a representation of the C record PangoAttrColor.
type AttrColor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrColor that represents the AttrColor.
func (recv *AttrColor) ToC() unsafe.Pointer {
	return recv.native
}

// AttrColorNewFromC creates a new AttrColor from a pointer to the C PangoAttrColor that represents the AttrColor.
func AttrColorNewFromC(native unsafe.Pointer) *AttrColor {
	return &AttrColor{native: native}
}

// AttrFloat is a representation of the C record PangoAttrFloat.
type AttrFloat struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrFloat that represents the AttrFloat.
func (recv *AttrFloat) ToC() unsafe.Pointer {
	return recv.native
}

// AttrFloatNewFromC creates a new AttrFloat from a pointer to the C PangoAttrFloat that represents the AttrFloat.
func AttrFloatNewFromC(native unsafe.Pointer) *AttrFloat {
	return &AttrFloat{native: native}
}

// AttrFontDesc is a representation of the C record PangoAttrFontDesc.
type AttrFontDesc struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrFontDesc that represents the AttrFontDesc.
func (recv *AttrFontDesc) ToC() unsafe.Pointer {
	return recv.native
}

// AttrFontDescNewFromC creates a new AttrFontDesc from a pointer to the C PangoAttrFontDesc that represents the AttrFontDesc.
func AttrFontDescNewFromC(native unsafe.Pointer) *AttrFontDesc {
	return &AttrFontDesc{native: native}
}

// AttrFontFeatures is a representation of the C record PangoAttrFontFeatures.
//
// since 1.38
type AttrFontFeatures struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrFontFeatures that represents the AttrFontFeatures.
func (recv *AttrFontFeatures) ToC() unsafe.Pointer {
	return recv.native
}

// AttrFontFeaturesNewFromC creates a new AttrFontFeatures from a pointer to the C PangoAttrFontFeatures that represents the AttrFontFeatures.
func AttrFontFeaturesNewFromC(native unsafe.Pointer) *AttrFontFeatures {
	return &AttrFontFeatures{native: native}
}

// AttrInt is a representation of the C record PangoAttrInt.
type AttrInt struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrInt that represents the AttrInt.
func (recv *AttrInt) ToC() unsafe.Pointer {
	return recv.native
}

// AttrIntNewFromC creates a new AttrInt from a pointer to the C PangoAttrInt that represents the AttrInt.
func AttrIntNewFromC(native unsafe.Pointer) *AttrInt {
	return &AttrInt{native: native}
}

// AttrIterator is a representation of the C record PangoAttrIterator.
type AttrIterator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrIterator that represents the AttrIterator.
func (recv *AttrIterator) ToC() unsafe.Pointer {
	return recv.native
}

// AttrIteratorNewFromC creates a new AttrIterator from a pointer to the C PangoAttrIterator that represents the AttrIterator.
func AttrIteratorNewFromC(native unsafe.Pointer) *AttrIterator {
	return &AttrIterator{native: native}
}

// AttrLanguage is a representation of the C record PangoAttrLanguage.
type AttrLanguage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrLanguage that represents the AttrLanguage.
func (recv *AttrLanguage) ToC() unsafe.Pointer {
	return recv.native
}

// AttrLanguageNewFromC creates a new AttrLanguage from a pointer to the C PangoAttrLanguage that represents the AttrLanguage.
func AttrLanguageNewFromC(native unsafe.Pointer) *AttrLanguage {
	return &AttrLanguage{native: native}
}

// AttrList is a representation of the C record PangoAttrList.
type AttrList struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrList that represents the AttrList.
func (recv *AttrList) ToC() unsafe.Pointer {
	return recv.native
}

// AttrListNewFromC creates a new AttrList from a pointer to the C PangoAttrList that represents the AttrList.
func AttrListNewFromC(native unsafe.Pointer) *AttrList {
	return &AttrList{native: native}
}

// AttrShape is a representation of the C record PangoAttrShape.
type AttrShape struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrShape that represents the AttrShape.
func (recv *AttrShape) ToC() unsafe.Pointer {
	return recv.native
}

// AttrShapeNewFromC creates a new AttrShape from a pointer to the C PangoAttrShape that represents the AttrShape.
func AttrShapeNewFromC(native unsafe.Pointer) *AttrShape {
	return &AttrShape{native: native}
}

// AttrSize is a representation of the C record PangoAttrSize.
type AttrSize struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrSize that represents the AttrSize.
func (recv *AttrSize) ToC() unsafe.Pointer {
	return recv.native
}

// AttrSizeNewFromC creates a new AttrSize from a pointer to the C PangoAttrSize that represents the AttrSize.
func AttrSizeNewFromC(native unsafe.Pointer) *AttrSize {
	return &AttrSize{native: native}
}

// AttrString is a representation of the C record PangoAttrString.
type AttrString struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrString that represents the AttrString.
func (recv *AttrString) ToC() unsafe.Pointer {
	return recv.native
}

// AttrStringNewFromC creates a new AttrString from a pointer to the C PangoAttrString that represents the AttrString.
func AttrStringNewFromC(native unsafe.Pointer) *AttrString {
	return &AttrString{native: native}
}

// Attribute is a representation of the C record PangoAttribute.
type Attribute struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttribute that represents the Attribute.
func (recv *Attribute) ToC() unsafe.Pointer {
	return recv.native
}

// AttributeNewFromC creates a new Attribute from a pointer to the C PangoAttribute that represents the Attribute.
func AttributeNewFromC(native unsafe.Pointer) *Attribute {
	return &Attribute{native: native}
}

// Color is a representation of the C record PangoColor.
type Color struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoColor that represents the Color.
func (recv *Color) ToC() unsafe.Pointer {
	return recv.native
}

// ColorNewFromC creates a new Color from a pointer to the C PangoColor that represents the Color.
func ColorNewFromC(native unsafe.Pointer) *Color {
	return &Color{native: native}
}

// ContextClass is a representation of the C record PangoContextClass.
type ContextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoContextClass that represents the ContextClass.
func (recv *ContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// ContextClassNewFromC creates a new ContextClass from a pointer to the C PangoContextClass that represents the ContextClass.
func ContextClassNewFromC(native unsafe.Pointer) *ContextClass {
	return &ContextClass{native: native}
}

// Coverage is a representation of the C record PangoCoverage.
type Coverage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoCoverage that represents the Coverage.
func (recv *Coverage) ToC() unsafe.Pointer {
	return recv.native
}

// CoverageNewFromC creates a new Coverage from a pointer to the C PangoCoverage that represents the Coverage.
func CoverageNewFromC(native unsafe.Pointer) *Coverage {
	return &Coverage{native: native}
}

// UNSUPPORTED : EngineClass : blacklisted

// UNSUPPORTED : EngineInfo : blacklisted

// UNSUPPORTED : EngineLangClass : blacklisted

// UNSUPPORTED : EngineScriptInfo : blacklisted

// UNSUPPORTED : EngineShapeClass : blacklisted

// UNSUPPORTED : FontClass : blacklisted

// FontDescription is a representation of the C record PangoFontDescription.
type FontDescription struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoFontDescription that represents the FontDescription.
func (recv *FontDescription) ToC() unsafe.Pointer {
	return recv.native
}

// FontDescriptionNewFromC creates a new FontDescription from a pointer to the C PangoFontDescription that represents the FontDescription.
func FontDescriptionNewFromC(native unsafe.Pointer) *FontDescription {
	return &FontDescription{native: native}
}

// UNSUPPORTED : FontFaceClass : blacklisted

// UNSUPPORTED : FontFamilyClass : blacklisted

// UNSUPPORTED : FontMapClass : blacklisted

// FontMetrics is a representation of the C record PangoFontMetrics.
type FontMetrics struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoFontMetrics that represents the FontMetrics.
func (recv *FontMetrics) ToC() unsafe.Pointer {
	return recv.native
}

// FontMetricsNewFromC creates a new FontMetrics from a pointer to the C PangoFontMetrics that represents the FontMetrics.
func FontMetricsNewFromC(native unsafe.Pointer) *FontMetrics {
	return &FontMetrics{native: native}
}

// UNSUPPORTED : FontsetClass : blacklisted

// UNSUPPORTED : FontsetSimpleClass : blacklisted

// GlyphGeometry is a representation of the C record PangoGlyphGeometry.
type GlyphGeometry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoGlyphGeometry that represents the GlyphGeometry.
func (recv *GlyphGeometry) ToC() unsafe.Pointer {
	return recv.native
}

// GlyphGeometryNewFromC creates a new GlyphGeometry from a pointer to the C PangoGlyphGeometry that represents the GlyphGeometry.
func GlyphGeometryNewFromC(native unsafe.Pointer) *GlyphGeometry {
	return &GlyphGeometry{native: native}
}

// GlyphInfo is a representation of the C record PangoGlyphInfo.
type GlyphInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoGlyphInfo that represents the GlyphInfo.
func (recv *GlyphInfo) ToC() unsafe.Pointer {
	return recv.native
}

// GlyphInfoNewFromC creates a new GlyphInfo from a pointer to the C PangoGlyphInfo that represents the GlyphInfo.
func GlyphInfoNewFromC(native unsafe.Pointer) *GlyphInfo {
	return &GlyphInfo{native: native}
}

// GlyphItem is a representation of the C record PangoGlyphItem.
type GlyphItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoGlyphItem that represents the GlyphItem.
func (recv *GlyphItem) ToC() unsafe.Pointer {
	return recv.native
}

// GlyphItemNewFromC creates a new GlyphItem from a pointer to the C PangoGlyphItem that represents the GlyphItem.
func GlyphItemNewFromC(native unsafe.Pointer) *GlyphItem {
	return &GlyphItem{native: native}
}

// GlyphItemIter is a representation of the C record PangoGlyphItemIter.
//
// since 1.22
type GlyphItemIter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoGlyphItemIter that represents the GlyphItemIter.
func (recv *GlyphItemIter) ToC() unsafe.Pointer {
	return recv.native
}

// GlyphItemIterNewFromC creates a new GlyphItemIter from a pointer to the C PangoGlyphItemIter that represents the GlyphItemIter.
func GlyphItemIterNewFromC(native unsafe.Pointer) *GlyphItemIter {
	return &GlyphItemIter{native: native}
}

// GlyphString is a representation of the C record PangoGlyphString.
type GlyphString struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoGlyphString that represents the GlyphString.
func (recv *GlyphString) ToC() unsafe.Pointer {
	return recv.native
}

// GlyphStringNewFromC creates a new GlyphString from a pointer to the C PangoGlyphString that represents the GlyphString.
func GlyphStringNewFromC(native unsafe.Pointer) *GlyphString {
	return &GlyphString{native: native}
}

// GlyphVisAttr is a representation of the C record PangoGlyphVisAttr.
type GlyphVisAttr struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoGlyphVisAttr that represents the GlyphVisAttr.
func (recv *GlyphVisAttr) ToC() unsafe.Pointer {
	return recv.native
}

// GlyphVisAttrNewFromC creates a new GlyphVisAttr from a pointer to the C PangoGlyphVisAttr that represents the GlyphVisAttr.
func GlyphVisAttrNewFromC(native unsafe.Pointer) *GlyphVisAttr {
	return &GlyphVisAttr{native: native}
}

// UNSUPPORTED : IncludedModule : blacklisted

// Item is a representation of the C record PangoItem.
type Item struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoItem that represents the Item.
func (recv *Item) ToC() unsafe.Pointer {
	return recv.native
}

// ItemNewFromC creates a new Item from a pointer to the C PangoItem that represents the Item.
func ItemNewFromC(native unsafe.Pointer) *Item {
	return &Item{native: native}
}

// Language is a representation of the C record PangoLanguage.
type Language struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoLanguage that represents the Language.
func (recv *Language) ToC() unsafe.Pointer {
	return recv.native
}

// LanguageNewFromC creates a new Language from a pointer to the C PangoLanguage that represents the Language.
func LanguageNewFromC(native unsafe.Pointer) *Language {
	return &Language{native: native}
}

// LayoutClass is a representation of the C record PangoLayoutClass.
type LayoutClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoLayoutClass that represents the LayoutClass.
func (recv *LayoutClass) ToC() unsafe.Pointer {
	return recv.native
}

// LayoutClassNewFromC creates a new LayoutClass from a pointer to the C PangoLayoutClass that represents the LayoutClass.
func LayoutClassNewFromC(native unsafe.Pointer) *LayoutClass {
	return &LayoutClass{native: native}
}

// LayoutIter is a representation of the C record PangoLayoutIter.
type LayoutIter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoLayoutIter that represents the LayoutIter.
func (recv *LayoutIter) ToC() unsafe.Pointer {
	return recv.native
}

// LayoutIterNewFromC creates a new LayoutIter from a pointer to the C PangoLayoutIter that represents the LayoutIter.
func LayoutIterNewFromC(native unsafe.Pointer) *LayoutIter {
	return &LayoutIter{native: native}
}

// LayoutLine is a representation of the C record PangoLayoutLine.
type LayoutLine struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoLayoutLine that represents the LayoutLine.
func (recv *LayoutLine) ToC() unsafe.Pointer {
	return recv.native
}

// LayoutLineNewFromC creates a new LayoutLine from a pointer to the C PangoLayoutLine that represents the LayoutLine.
func LayoutLineNewFromC(native unsafe.Pointer) *LayoutLine {
	return &LayoutLine{native: native}
}

// LogAttr is a representation of the C record PangoLogAttr.
type LogAttr struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoLogAttr that represents the LogAttr.
func (recv *LogAttr) ToC() unsafe.Pointer {
	return recv.native
}

// LogAttrNewFromC creates a new LogAttr from a pointer to the C PangoLogAttr that represents the LogAttr.
func LogAttrNewFromC(native unsafe.Pointer) *LogAttr {
	return &LogAttr{native: native}
}

// UNSUPPORTED : Map : blacklisted

// UNSUPPORTED : MapEntry : blacklisted

// Matrix is a representation of the C record PangoMatrix.
//
// since 1.6
type Matrix struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoMatrix that represents the Matrix.
func (recv *Matrix) ToC() unsafe.Pointer {
	return recv.native
}

// MatrixNewFromC creates a new Matrix from a pointer to the C PangoMatrix that represents the Matrix.
func MatrixNewFromC(native unsafe.Pointer) *Matrix {
	return &Matrix{native: native}
}

// Rectangle is a representation of the C record PangoRectangle.
type Rectangle struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoRectangle that represents the Rectangle.
func (recv *Rectangle) ToC() unsafe.Pointer {
	return recv.native
}

// RectangleNewFromC creates a new Rectangle from a pointer to the C PangoRectangle that represents the Rectangle.
func RectangleNewFromC(native unsafe.Pointer) *Rectangle {
	return &Rectangle{native: native}
}

// RendererClass is a representation of the C record PangoRendererClass.
//
// since 1.8
type RendererClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoRendererClass that represents the RendererClass.
func (recv *RendererClass) ToC() unsafe.Pointer {
	return recv.native
}

// RendererClassNewFromC creates a new RendererClass from a pointer to the C PangoRendererClass that represents the RendererClass.
func RendererClassNewFromC(native unsafe.Pointer) *RendererClass {
	return &RendererClass{native: native}
}

// RendererPrivate is a representation of the C record PangoRendererPrivate.
type RendererPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoRendererPrivate that represents the RendererPrivate.
func (recv *RendererPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// RendererPrivateNewFromC creates a new RendererPrivate from a pointer to the C PangoRendererPrivate that represents the RendererPrivate.
func RendererPrivateNewFromC(native unsafe.Pointer) *RendererPrivate {
	return &RendererPrivate{native: native}
}

// ScriptIter is a representation of the C record PangoScriptIter.
type ScriptIter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoScriptIter that represents the ScriptIter.
func (recv *ScriptIter) ToC() unsafe.Pointer {
	return recv.native
}

// ScriptIterNewFromC creates a new ScriptIter from a pointer to the C PangoScriptIter that represents the ScriptIter.
func ScriptIterNewFromC(native unsafe.Pointer) *ScriptIter {
	return &ScriptIter{native: native}
}

// TabArray is a representation of the C record PangoTabArray.
type TabArray struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoTabArray that represents the TabArray.
func (recv *TabArray) ToC() unsafe.Pointer {
	return recv.native
}

// TabArrayNewFromC creates a new TabArray from a pointer to the C PangoTabArray that represents the TabArray.
func TabArrayNewFromC(native unsafe.Pointer) *TabArray {
	return &TabArray{native: native}
}

// Context is a representation of the C record PangoContext.
type Context struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoContext that represents the Context.
func (recv *Context) ToC() unsafe.Pointer {
	return recv.native
}

// ContextNewFromC creates a new Context from a pointer to the C PangoContext that represents the Context.
func ContextNewFromC(native unsafe.Pointer) *Context {
	return &Context{native: native}
}

// UNSUPPORTED : Engine : blacklisted

// EngineLang is a representation of the C record PangoEngineLang.
type EngineLang struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoEngineLang that represents the EngineLang.
func (recv *EngineLang) ToC() unsafe.Pointer {
	return recv.native
}

// EngineLangNewFromC creates a new EngineLang from a pointer to the C PangoEngineLang that represents the EngineLang.
func EngineLangNewFromC(native unsafe.Pointer) *EngineLang {
	return &EngineLang{native: native}
}

// EngineShape is a representation of the C record PangoEngineShape.
type EngineShape struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoEngineShape that represents the EngineShape.
func (recv *EngineShape) ToC() unsafe.Pointer {
	return recv.native
}

// EngineShapeNewFromC creates a new EngineShape from a pointer to the C PangoEngineShape that represents the EngineShape.
func EngineShapeNewFromC(native unsafe.Pointer) *EngineShape {
	return &EngineShape{native: native}
}

// Font is a representation of the C record PangoFont.
type Font struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoFont that represents the Font.
func (recv *Font) ToC() unsafe.Pointer {
	return recv.native
}

// FontNewFromC creates a new Font from a pointer to the C PangoFont that represents the Font.
func FontNewFromC(native unsafe.Pointer) *Font {
	return &Font{native: native}
}

// FontFace is a representation of the C record PangoFontFace.
type FontFace struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoFontFace that represents the FontFace.
func (recv *FontFace) ToC() unsafe.Pointer {
	return recv.native
}

// FontFaceNewFromC creates a new FontFace from a pointer to the C PangoFontFace that represents the FontFace.
func FontFaceNewFromC(native unsafe.Pointer) *FontFace {
	return &FontFace{native: native}
}

// FontFamily is a representation of the C record PangoFontFamily.
type FontFamily struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoFontFamily that represents the FontFamily.
func (recv *FontFamily) ToC() unsafe.Pointer {
	return recv.native
}

// FontFamilyNewFromC creates a new FontFamily from a pointer to the C PangoFontFamily that represents the FontFamily.
func FontFamilyNewFromC(native unsafe.Pointer) *FontFamily {
	return &FontFamily{native: native}
}

// FontMap is a representation of the C record PangoFontMap.
type FontMap struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoFontMap that represents the FontMap.
func (recv *FontMap) ToC() unsafe.Pointer {
	return recv.native
}

// FontMapNewFromC creates a new FontMap from a pointer to the C PangoFontMap that represents the FontMap.
func FontMapNewFromC(native unsafe.Pointer) *FontMap {
	return &FontMap{native: native}
}

// Fontset is a representation of the C record PangoFontset.
type Fontset struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoFontset that represents the Fontset.
func (recv *Fontset) ToC() unsafe.Pointer {
	return recv.native
}

// FontsetNewFromC creates a new Fontset from a pointer to the C PangoFontset that represents the Fontset.
func FontsetNewFromC(native unsafe.Pointer) *Fontset {
	return &Fontset{native: native}
}

// UNSUPPORTED : FontsetSimple : blacklisted

// Layout is a representation of the C record PangoLayout.
type Layout struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoLayout that represents the Layout.
func (recv *Layout) ToC() unsafe.Pointer {
	return recv.native
}

// LayoutNewFromC creates a new Layout from a pointer to the C PangoLayout that represents the Layout.
func LayoutNewFromC(native unsafe.Pointer) *Layout {
	return &Layout{native: native}
}

// Renderer is a representation of the C record PangoRenderer.
//
// since 1.8
type Renderer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoRenderer that represents the Renderer.
func (recv *Renderer) ToC() unsafe.Pointer {
	return recv.native
}

// RendererNewFromC creates a new Renderer from a pointer to the C PangoRenderer that represents the Renderer.
func RendererNewFromC(native unsafe.Pointer) *Renderer {
	return &Renderer{native: native}
}
