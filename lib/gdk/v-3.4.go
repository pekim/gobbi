// Code generated - DO NOT EDIT.
// +build gdk_3.4

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// UNSUPPORTED : XEvent : blacklisted

// BUTTON_MIDDLE is a representation of the C constant GDK_BUTTON_MIDDLE.
//
// since 3.4
const BUTTON_MIDDLE = 2

// BUTTON_PRIMARY is a representation of the C constant GDK_BUTTON_PRIMARY.
//
// since 3.4
const BUTTON_PRIMARY = 1

// BUTTON_SECONDARY is a representation of the C constant GDK_BUTTON_SECONDARY.
//
// since 3.4
const BUTTON_SECONDARY = 3

// CURRENT_TIME is a representation of the C constant GDK_CURRENT_TIME.
const CURRENT_TIME = 0

// EVENT_PROPAGATE is a representation of the C constant GDK_EVENT_PROPAGATE.
//
// since 3.4
const EVENT_PROPAGATE = false

// EVENT_STOP is a representation of the C constant GDK_EVENT_STOP.
//
// since 3.4
const EVENT_STOP = true

// KEY_0 is a representation of the C constant GDK_KEY_0.
const KEY_0 = 48

// KEY_1 is a representation of the C constant GDK_KEY_1.
const KEY_1 = 49

// KEY_2 is a representation of the C constant GDK_KEY_2.
const KEY_2 = 50

// KEY_3 is a representation of the C constant GDK_KEY_3.
const KEY_3 = 51

// KEY_3270_AltCursor is a representation of the C constant GDK_KEY_3270_AltCursor.
const KEY_3270_AltCursor = 64784

// KEY_3270_Attn is a representation of the C constant GDK_KEY_3270_Attn.
const KEY_3270_Attn = 64782

// KEY_3270_BackTab is a representation of the C constant GDK_KEY_3270_BackTab.
const KEY_3270_BackTab = 64773

// KEY_3270_ChangeScreen is a representation of the C constant GDK_KEY_3270_ChangeScreen.
const KEY_3270_ChangeScreen = 64793

// KEY_3270_Copy is a representation of the C constant GDK_KEY_3270_Copy.
const KEY_3270_Copy = 64789

// KEY_3270_CursorBlink is a representation of the C constant GDK_KEY_3270_CursorBlink.
const KEY_3270_CursorBlink = 64783

// KEY_3270_CursorSelect is a representation of the C constant GDK_KEY_3270_CursorSelect.
const KEY_3270_CursorSelect = 64796

// KEY_3270_DeleteWord is a representation of the C constant GDK_KEY_3270_DeleteWord.
const KEY_3270_DeleteWord = 64794

// KEY_3270_Duplicate is a representation of the C constant GDK_KEY_3270_Duplicate.
const KEY_3270_Duplicate = 64769

// KEY_3270_Enter is a representation of the C constant GDK_KEY_3270_Enter.
const KEY_3270_Enter = 64798

// KEY_3270_EraseEOF is a representation of the C constant GDK_KEY_3270_EraseEOF.
const KEY_3270_EraseEOF = 64774

// KEY_3270_EraseInput is a representation of the C constant GDK_KEY_3270_EraseInput.
const KEY_3270_EraseInput = 64775

// KEY_3270_ExSelect is a representation of the C constant GDK_KEY_3270_ExSelect.
const KEY_3270_ExSelect = 64795

// KEY_3270_FieldMark is a representation of the C constant GDK_KEY_3270_FieldMark.
const KEY_3270_FieldMark = 64770

// KEY_3270_Ident is a representation of the C constant GDK_KEY_3270_Ident.
const KEY_3270_Ident = 64787

// KEY_3270_Jump is a representation of the C constant GDK_KEY_3270_Jump.
const KEY_3270_Jump = 64786

// KEY_3270_KeyClick is a representation of the C constant GDK_KEY_3270_KeyClick.
const KEY_3270_KeyClick = 64785

// KEY_3270_Left2 is a representation of the C constant GDK_KEY_3270_Left2.
const KEY_3270_Left2 = 64772

// KEY_3270_PA1 is a representation of the C constant GDK_KEY_3270_PA1.
const KEY_3270_PA1 = 64778

// KEY_3270_PA2 is a representation of the C constant GDK_KEY_3270_PA2.
const KEY_3270_PA2 = 64779

// KEY_3270_PA3 is a representation of the C constant GDK_KEY_3270_PA3.
const KEY_3270_PA3 = 64780

// KEY_3270_Play is a representation of the C constant GDK_KEY_3270_Play.
const KEY_3270_Play = 64790

// KEY_3270_PrintScreen is a representation of the C constant GDK_KEY_3270_PrintScreen.
const KEY_3270_PrintScreen = 64797

// KEY_3270_Quit is a representation of the C constant GDK_KEY_3270_Quit.
const KEY_3270_Quit = 64777

// KEY_3270_Record is a representation of the C constant GDK_KEY_3270_Record.
const KEY_3270_Record = 64792

// KEY_3270_Reset is a representation of the C constant GDK_KEY_3270_Reset.
const KEY_3270_Reset = 64776

// KEY_3270_Right2 is a representation of the C constant GDK_KEY_3270_Right2.
const KEY_3270_Right2 = 64771

// KEY_3270_Rule is a representation of the C constant GDK_KEY_3270_Rule.
const KEY_3270_Rule = 64788

// KEY_3270_Setup is a representation of the C constant GDK_KEY_3270_Setup.
const KEY_3270_Setup = 64791

// KEY_3270_Test is a representation of the C constant GDK_KEY_3270_Test.
const KEY_3270_Test = 64781

// KEY_4 is a representation of the C constant GDK_KEY_4.
const KEY_4 = 52

// KEY_5 is a representation of the C constant GDK_KEY_5.
const KEY_5 = 53

// KEY_6 is a representation of the C constant GDK_KEY_6.
const KEY_6 = 54

// KEY_7 is a representation of the C constant GDK_KEY_7.
const KEY_7 = 55

// KEY_8 is a representation of the C constant GDK_KEY_8.
const KEY_8 = 56

// KEY_9 is a representation of the C constant GDK_KEY_9.
const KEY_9 = 57

// KEY_A is a representation of the C constant GDK_KEY_A.
const KEY_A = 65

// KEY_AE is a representation of the C constant GDK_KEY_AE.
const KEY_AE = 198

// KEY_Aacute is a representation of the C constant GDK_KEY_Aacute.
const KEY_Aacute = 193

// KEY_Abelowdot is a representation of the C constant GDK_KEY_Abelowdot.
const KEY_Abelowdot = 16785056

// KEY_Abreve is a representation of the C constant GDK_KEY_Abreve.
const KEY_Abreve = 451

// KEY_Abreveacute is a representation of the C constant GDK_KEY_Abreveacute.
const KEY_Abreveacute = 16785070

// KEY_Abrevebelowdot is a representation of the C constant GDK_KEY_Abrevebelowdot.
const KEY_Abrevebelowdot = 16785078

// KEY_Abrevegrave is a representation of the C constant GDK_KEY_Abrevegrave.
const KEY_Abrevegrave = 16785072

// KEY_Abrevehook is a representation of the C constant GDK_KEY_Abrevehook.
const KEY_Abrevehook = 16785074

// KEY_Abrevetilde is a representation of the C constant GDK_KEY_Abrevetilde.
const KEY_Abrevetilde = 16785076

// KEY_AccessX_Enable is a representation of the C constant GDK_KEY_AccessX_Enable.
const KEY_AccessX_Enable = 65136

// KEY_AccessX_Feedback_Enable is a representation of the C constant GDK_KEY_AccessX_Feedback_Enable.
const KEY_AccessX_Feedback_Enable = 65137

// KEY_Acircumflex is a representation of the C constant GDK_KEY_Acircumflex.
const KEY_Acircumflex = 194

// KEY_Acircumflexacute is a representation of the C constant GDK_KEY_Acircumflexacute.
const KEY_Acircumflexacute = 16785060

// KEY_Acircumflexbelowdot is a representation of the C constant GDK_KEY_Acircumflexbelowdot.
const KEY_Acircumflexbelowdot = 16785068

// KEY_Acircumflexgrave is a representation of the C constant GDK_KEY_Acircumflexgrave.
const KEY_Acircumflexgrave = 16785062

// KEY_Acircumflexhook is a representation of the C constant GDK_KEY_Acircumflexhook.
const KEY_Acircumflexhook = 16785064

// KEY_Acircumflextilde is a representation of the C constant GDK_KEY_Acircumflextilde.
const KEY_Acircumflextilde = 16785066

// KEY_AddFavorite is a representation of the C constant GDK_KEY_AddFavorite.
const KEY_AddFavorite = 269025081

// KEY_Adiaeresis is a representation of the C constant GDK_KEY_Adiaeresis.
const KEY_Adiaeresis = 196

// KEY_Agrave is a representation of the C constant GDK_KEY_Agrave.
const KEY_Agrave = 192

// KEY_Ahook is a representation of the C constant GDK_KEY_Ahook.
const KEY_Ahook = 16785058

// KEY_Alt_L is a representation of the C constant GDK_KEY_Alt_L.
const KEY_Alt_L = 65513

// KEY_Alt_R is a representation of the C constant GDK_KEY_Alt_R.
const KEY_Alt_R = 65514

// KEY_Amacron is a representation of the C constant GDK_KEY_Amacron.
const KEY_Amacron = 960

// KEY_Aogonek is a representation of the C constant GDK_KEY_Aogonek.
const KEY_Aogonek = 417

// KEY_ApplicationLeft is a representation of the C constant GDK_KEY_ApplicationLeft.
const KEY_ApplicationLeft = 269025104

// KEY_ApplicationRight is a representation of the C constant GDK_KEY_ApplicationRight.
const KEY_ApplicationRight = 269025105

// KEY_Arabic_0 is a representation of the C constant GDK_KEY_Arabic_0.
const KEY_Arabic_0 = 16778848

// KEY_Arabic_1 is a representation of the C constant GDK_KEY_Arabic_1.
const KEY_Arabic_1 = 16778849

// KEY_Arabic_2 is a representation of the C constant GDK_KEY_Arabic_2.
const KEY_Arabic_2 = 16778850

// KEY_Arabic_3 is a representation of the C constant GDK_KEY_Arabic_3.
const KEY_Arabic_3 = 16778851

// KEY_Arabic_4 is a representation of the C constant GDK_KEY_Arabic_4.
const KEY_Arabic_4 = 16778852

// KEY_Arabic_5 is a representation of the C constant GDK_KEY_Arabic_5.
const KEY_Arabic_5 = 16778853

// KEY_Arabic_6 is a representation of the C constant GDK_KEY_Arabic_6.
const KEY_Arabic_6 = 16778854

// KEY_Arabic_7 is a representation of the C constant GDK_KEY_Arabic_7.
const KEY_Arabic_7 = 16778855

// KEY_Arabic_8 is a representation of the C constant GDK_KEY_Arabic_8.
const KEY_Arabic_8 = 16778856

// KEY_Arabic_9 is a representation of the C constant GDK_KEY_Arabic_9.
const KEY_Arabic_9 = 16778857

// KEY_Arabic_ain is a representation of the C constant GDK_KEY_Arabic_ain.
const KEY_Arabic_ain = 1497

// KEY_Arabic_alef is a representation of the C constant GDK_KEY_Arabic_alef.
const KEY_Arabic_alef = 1479

// KEY_Arabic_alefmaksura is a representation of the C constant GDK_KEY_Arabic_alefmaksura.
const KEY_Arabic_alefmaksura = 1513

// KEY_Arabic_beh is a representation of the C constant GDK_KEY_Arabic_beh.
const KEY_Arabic_beh = 1480

// KEY_Arabic_comma is a representation of the C constant GDK_KEY_Arabic_comma.
const KEY_Arabic_comma = 1452

// KEY_Arabic_dad is a representation of the C constant GDK_KEY_Arabic_dad.
const KEY_Arabic_dad = 1494

// KEY_Arabic_dal is a representation of the C constant GDK_KEY_Arabic_dal.
const KEY_Arabic_dal = 1487

// KEY_Arabic_damma is a representation of the C constant GDK_KEY_Arabic_damma.
const KEY_Arabic_damma = 1519

// KEY_Arabic_dammatan is a representation of the C constant GDK_KEY_Arabic_dammatan.
const KEY_Arabic_dammatan = 1516

// KEY_Arabic_ddal is a representation of the C constant GDK_KEY_Arabic_ddal.
const KEY_Arabic_ddal = 16778888

// KEY_Arabic_farsi_yeh is a representation of the C constant GDK_KEY_Arabic_farsi_yeh.
const KEY_Arabic_farsi_yeh = 16778956

// KEY_Arabic_fatha is a representation of the C constant GDK_KEY_Arabic_fatha.
const KEY_Arabic_fatha = 1518

// KEY_Arabic_fathatan is a representation of the C constant GDK_KEY_Arabic_fathatan.
const KEY_Arabic_fathatan = 1515

// KEY_Arabic_feh is a representation of the C constant GDK_KEY_Arabic_feh.
const KEY_Arabic_feh = 1505

// KEY_Arabic_fullstop is a representation of the C constant GDK_KEY_Arabic_fullstop.
const KEY_Arabic_fullstop = 16778964

// KEY_Arabic_gaf is a representation of the C constant GDK_KEY_Arabic_gaf.
const KEY_Arabic_gaf = 16778927

// KEY_Arabic_ghain is a representation of the C constant GDK_KEY_Arabic_ghain.
const KEY_Arabic_ghain = 1498

// KEY_Arabic_ha is a representation of the C constant GDK_KEY_Arabic_ha.
const KEY_Arabic_ha = 1511

// KEY_Arabic_hah is a representation of the C constant GDK_KEY_Arabic_hah.
const KEY_Arabic_hah = 1485

// KEY_Arabic_hamza is a representation of the C constant GDK_KEY_Arabic_hamza.
const KEY_Arabic_hamza = 1473

// KEY_Arabic_hamza_above is a representation of the C constant GDK_KEY_Arabic_hamza_above.
const KEY_Arabic_hamza_above = 16778836

// KEY_Arabic_hamza_below is a representation of the C constant GDK_KEY_Arabic_hamza_below.
const KEY_Arabic_hamza_below = 16778837

// KEY_Arabic_hamzaonalef is a representation of the C constant GDK_KEY_Arabic_hamzaonalef.
const KEY_Arabic_hamzaonalef = 1475

// KEY_Arabic_hamzaonwaw is a representation of the C constant GDK_KEY_Arabic_hamzaonwaw.
const KEY_Arabic_hamzaonwaw = 1476

// KEY_Arabic_hamzaonyeh is a representation of the C constant GDK_KEY_Arabic_hamzaonyeh.
const KEY_Arabic_hamzaonyeh = 1478

// KEY_Arabic_hamzaunderalef is a representation of the C constant GDK_KEY_Arabic_hamzaunderalef.
const KEY_Arabic_hamzaunderalef = 1477

// KEY_Arabic_heh is a representation of the C constant GDK_KEY_Arabic_heh.
const KEY_Arabic_heh = 1511

// KEY_Arabic_heh_doachashmee is a representation of the C constant GDK_KEY_Arabic_heh_doachashmee.
const KEY_Arabic_heh_doachashmee = 16778942

// KEY_Arabic_heh_goal is a representation of the C constant GDK_KEY_Arabic_heh_goal.
const KEY_Arabic_heh_goal = 16778945

// KEY_Arabic_jeem is a representation of the C constant GDK_KEY_Arabic_jeem.
const KEY_Arabic_jeem = 1484

// KEY_Arabic_jeh is a representation of the C constant GDK_KEY_Arabic_jeh.
const KEY_Arabic_jeh = 16778904

// KEY_Arabic_kaf is a representation of the C constant GDK_KEY_Arabic_kaf.
const KEY_Arabic_kaf = 1507

// KEY_Arabic_kasra is a representation of the C constant GDK_KEY_Arabic_kasra.
const KEY_Arabic_kasra = 1520

// KEY_Arabic_kasratan is a representation of the C constant GDK_KEY_Arabic_kasratan.
const KEY_Arabic_kasratan = 1517

// KEY_Arabic_keheh is a representation of the C constant GDK_KEY_Arabic_keheh.
const KEY_Arabic_keheh = 16778921

// KEY_Arabic_khah is a representation of the C constant GDK_KEY_Arabic_khah.
const KEY_Arabic_khah = 1486

// KEY_Arabic_lam is a representation of the C constant GDK_KEY_Arabic_lam.
const KEY_Arabic_lam = 1508

// KEY_Arabic_madda_above is a representation of the C constant GDK_KEY_Arabic_madda_above.
const KEY_Arabic_madda_above = 16778835

// KEY_Arabic_maddaonalef is a representation of the C constant GDK_KEY_Arabic_maddaonalef.
const KEY_Arabic_maddaonalef = 1474

// KEY_Arabic_meem is a representation of the C constant GDK_KEY_Arabic_meem.
const KEY_Arabic_meem = 1509

// KEY_Arabic_noon is a representation of the C constant GDK_KEY_Arabic_noon.
const KEY_Arabic_noon = 1510

// KEY_Arabic_noon_ghunna is a representation of the C constant GDK_KEY_Arabic_noon_ghunna.
const KEY_Arabic_noon_ghunna = 16778938

// KEY_Arabic_peh is a representation of the C constant GDK_KEY_Arabic_peh.
const KEY_Arabic_peh = 16778878

// KEY_Arabic_percent is a representation of the C constant GDK_KEY_Arabic_percent.
const KEY_Arabic_percent = 16778858

// KEY_Arabic_qaf is a representation of the C constant GDK_KEY_Arabic_qaf.
const KEY_Arabic_qaf = 1506

// KEY_Arabic_question_mark is a representation of the C constant GDK_KEY_Arabic_question_mark.
const KEY_Arabic_question_mark = 1471

// KEY_Arabic_ra is a representation of the C constant GDK_KEY_Arabic_ra.
const KEY_Arabic_ra = 1489

// KEY_Arabic_rreh is a representation of the C constant GDK_KEY_Arabic_rreh.
const KEY_Arabic_rreh = 16778897

// KEY_Arabic_sad is a representation of the C constant GDK_KEY_Arabic_sad.
const KEY_Arabic_sad = 1493

// KEY_Arabic_seen is a representation of the C constant GDK_KEY_Arabic_seen.
const KEY_Arabic_seen = 1491

// KEY_Arabic_semicolon is a representation of the C constant GDK_KEY_Arabic_semicolon.
const KEY_Arabic_semicolon = 1467

// KEY_Arabic_shadda is a representation of the C constant GDK_KEY_Arabic_shadda.
const KEY_Arabic_shadda = 1521

// KEY_Arabic_sheen is a representation of the C constant GDK_KEY_Arabic_sheen.
const KEY_Arabic_sheen = 1492

// KEY_Arabic_sukun is a representation of the C constant GDK_KEY_Arabic_sukun.
const KEY_Arabic_sukun = 1522

// KEY_Arabic_superscript_alef is a representation of the C constant GDK_KEY_Arabic_superscript_alef.
const KEY_Arabic_superscript_alef = 16778864

// KEY_Arabic_switch is a representation of the C constant GDK_KEY_Arabic_switch.
const KEY_Arabic_switch = 65406

// KEY_Arabic_tah is a representation of the C constant GDK_KEY_Arabic_tah.
const KEY_Arabic_tah = 1495

// KEY_Arabic_tatweel is a representation of the C constant GDK_KEY_Arabic_tatweel.
const KEY_Arabic_tatweel = 1504

// KEY_Arabic_tcheh is a representation of the C constant GDK_KEY_Arabic_tcheh.
const KEY_Arabic_tcheh = 16778886

// KEY_Arabic_teh is a representation of the C constant GDK_KEY_Arabic_teh.
const KEY_Arabic_teh = 1482

// KEY_Arabic_tehmarbuta is a representation of the C constant GDK_KEY_Arabic_tehmarbuta.
const KEY_Arabic_tehmarbuta = 1481

// KEY_Arabic_thal is a representation of the C constant GDK_KEY_Arabic_thal.
const KEY_Arabic_thal = 1488

// KEY_Arabic_theh is a representation of the C constant GDK_KEY_Arabic_theh.
const KEY_Arabic_theh = 1483

// KEY_Arabic_tteh is a representation of the C constant GDK_KEY_Arabic_tteh.
const KEY_Arabic_tteh = 16778873

// KEY_Arabic_veh is a representation of the C constant GDK_KEY_Arabic_veh.
const KEY_Arabic_veh = 16778916

// KEY_Arabic_waw is a representation of the C constant GDK_KEY_Arabic_waw.
const KEY_Arabic_waw = 1512

// KEY_Arabic_yeh is a representation of the C constant GDK_KEY_Arabic_yeh.
const KEY_Arabic_yeh = 1514

// KEY_Arabic_yeh_baree is a representation of the C constant GDK_KEY_Arabic_yeh_baree.
const KEY_Arabic_yeh_baree = 16778962

// KEY_Arabic_zah is a representation of the C constant GDK_KEY_Arabic_zah.
const KEY_Arabic_zah = 1496

// KEY_Arabic_zain is a representation of the C constant GDK_KEY_Arabic_zain.
const KEY_Arabic_zain = 1490

// KEY_Aring is a representation of the C constant GDK_KEY_Aring.
const KEY_Aring = 197

// KEY_Armenian_AT is a representation of the C constant GDK_KEY_Armenian_AT.
const KEY_Armenian_AT = 16778552

// KEY_Armenian_AYB is a representation of the C constant GDK_KEY_Armenian_AYB.
const KEY_Armenian_AYB = 16778545

// KEY_Armenian_BEN is a representation of the C constant GDK_KEY_Armenian_BEN.
const KEY_Armenian_BEN = 16778546

// KEY_Armenian_CHA is a representation of the C constant GDK_KEY_Armenian_CHA.
const KEY_Armenian_CHA = 16778569

// KEY_Armenian_DA is a representation of the C constant GDK_KEY_Armenian_DA.
const KEY_Armenian_DA = 16778548

// KEY_Armenian_DZA is a representation of the C constant GDK_KEY_Armenian_DZA.
const KEY_Armenian_DZA = 16778561

// KEY_Armenian_E is a representation of the C constant GDK_KEY_Armenian_E.
const KEY_Armenian_E = 16778551

// KEY_Armenian_FE is a representation of the C constant GDK_KEY_Armenian_FE.
const KEY_Armenian_FE = 16778582

// KEY_Armenian_GHAT is a representation of the C constant GDK_KEY_Armenian_GHAT.
const KEY_Armenian_GHAT = 16778562

// KEY_Armenian_GIM is a representation of the C constant GDK_KEY_Armenian_GIM.
const KEY_Armenian_GIM = 16778547

// KEY_Armenian_HI is a representation of the C constant GDK_KEY_Armenian_HI.
const KEY_Armenian_HI = 16778565

// KEY_Armenian_HO is a representation of the C constant GDK_KEY_Armenian_HO.
const KEY_Armenian_HO = 16778560

// KEY_Armenian_INI is a representation of the C constant GDK_KEY_Armenian_INI.
const KEY_Armenian_INI = 16778555

// KEY_Armenian_JE is a representation of the C constant GDK_KEY_Armenian_JE.
const KEY_Armenian_JE = 16778571

// KEY_Armenian_KE is a representation of the C constant GDK_KEY_Armenian_KE.
const KEY_Armenian_KE = 16778580

// KEY_Armenian_KEN is a representation of the C constant GDK_KEY_Armenian_KEN.
const KEY_Armenian_KEN = 16778559

// KEY_Armenian_KHE is a representation of the C constant GDK_KEY_Armenian_KHE.
const KEY_Armenian_KHE = 16778557

// KEY_Armenian_LYUN is a representation of the C constant GDK_KEY_Armenian_LYUN.
const KEY_Armenian_LYUN = 16778556

// KEY_Armenian_MEN is a representation of the C constant GDK_KEY_Armenian_MEN.
const KEY_Armenian_MEN = 16778564

// KEY_Armenian_NU is a representation of the C constant GDK_KEY_Armenian_NU.
const KEY_Armenian_NU = 16778566

// KEY_Armenian_O is a representation of the C constant GDK_KEY_Armenian_O.
const KEY_Armenian_O = 16778581

// KEY_Armenian_PE is a representation of the C constant GDK_KEY_Armenian_PE.
const KEY_Armenian_PE = 16778570

// KEY_Armenian_PYUR is a representation of the C constant GDK_KEY_Armenian_PYUR.
const KEY_Armenian_PYUR = 16778579

// KEY_Armenian_RA is a representation of the C constant GDK_KEY_Armenian_RA.
const KEY_Armenian_RA = 16778572

// KEY_Armenian_RE is a representation of the C constant GDK_KEY_Armenian_RE.
const KEY_Armenian_RE = 16778576

// KEY_Armenian_SE is a representation of the C constant GDK_KEY_Armenian_SE.
const KEY_Armenian_SE = 16778573

// KEY_Armenian_SHA is a representation of the C constant GDK_KEY_Armenian_SHA.
const KEY_Armenian_SHA = 16778567

// KEY_Armenian_TCHE is a representation of the C constant GDK_KEY_Armenian_TCHE.
const KEY_Armenian_TCHE = 16778563

// KEY_Armenian_TO is a representation of the C constant GDK_KEY_Armenian_TO.
const KEY_Armenian_TO = 16778553

// KEY_Armenian_TSA is a representation of the C constant GDK_KEY_Armenian_TSA.
const KEY_Armenian_TSA = 16778558

// KEY_Armenian_TSO is a representation of the C constant GDK_KEY_Armenian_TSO.
const KEY_Armenian_TSO = 16778577

// KEY_Armenian_TYUN is a representation of the C constant GDK_KEY_Armenian_TYUN.
const KEY_Armenian_TYUN = 16778575

// KEY_Armenian_VEV is a representation of the C constant GDK_KEY_Armenian_VEV.
const KEY_Armenian_VEV = 16778574

// KEY_Armenian_VO is a representation of the C constant GDK_KEY_Armenian_VO.
const KEY_Armenian_VO = 16778568

// KEY_Armenian_VYUN is a representation of the C constant GDK_KEY_Armenian_VYUN.
const KEY_Armenian_VYUN = 16778578

// KEY_Armenian_YECH is a representation of the C constant GDK_KEY_Armenian_YECH.
const KEY_Armenian_YECH = 16778549

// KEY_Armenian_ZA is a representation of the C constant GDK_KEY_Armenian_ZA.
const KEY_Armenian_ZA = 16778550

// KEY_Armenian_ZHE is a representation of the C constant GDK_KEY_Armenian_ZHE.
const KEY_Armenian_ZHE = 16778554

// KEY_Armenian_accent is a representation of the C constant GDK_KEY_Armenian_accent.
const KEY_Armenian_accent = 16778587

// KEY_Armenian_amanak is a representation of the C constant GDK_KEY_Armenian_amanak.
const KEY_Armenian_amanak = 16778588

// KEY_Armenian_apostrophe is a representation of the C constant GDK_KEY_Armenian_apostrophe.
const KEY_Armenian_apostrophe = 16778586

// KEY_Armenian_at is a representation of the C constant GDK_KEY_Armenian_at.
const KEY_Armenian_at = 16778600

// KEY_Armenian_ayb is a representation of the C constant GDK_KEY_Armenian_ayb.
const KEY_Armenian_ayb = 16778593

// KEY_Armenian_ben is a representation of the C constant GDK_KEY_Armenian_ben.
const KEY_Armenian_ben = 16778594

// KEY_Armenian_but is a representation of the C constant GDK_KEY_Armenian_but.
const KEY_Armenian_but = 16778589

// KEY_Armenian_cha is a representation of the C constant GDK_KEY_Armenian_cha.
const KEY_Armenian_cha = 16778617

// KEY_Armenian_da is a representation of the C constant GDK_KEY_Armenian_da.
const KEY_Armenian_da = 16778596

// KEY_Armenian_dza is a representation of the C constant GDK_KEY_Armenian_dza.
const KEY_Armenian_dza = 16778609

// KEY_Armenian_e is a representation of the C constant GDK_KEY_Armenian_e.
const KEY_Armenian_e = 16778599

// KEY_Armenian_exclam is a representation of the C constant GDK_KEY_Armenian_exclam.
const KEY_Armenian_exclam = 16778588

// KEY_Armenian_fe is a representation of the C constant GDK_KEY_Armenian_fe.
const KEY_Armenian_fe = 16778630

// KEY_Armenian_full_stop is a representation of the C constant GDK_KEY_Armenian_full_stop.
const KEY_Armenian_full_stop = 16778633

// KEY_Armenian_ghat is a representation of the C constant GDK_KEY_Armenian_ghat.
const KEY_Armenian_ghat = 16778610

// KEY_Armenian_gim is a representation of the C constant GDK_KEY_Armenian_gim.
const KEY_Armenian_gim = 16778595

// KEY_Armenian_hi is a representation of the C constant GDK_KEY_Armenian_hi.
const KEY_Armenian_hi = 16778613

// KEY_Armenian_ho is a representation of the C constant GDK_KEY_Armenian_ho.
const KEY_Armenian_ho = 16778608

// KEY_Armenian_hyphen is a representation of the C constant GDK_KEY_Armenian_hyphen.
const KEY_Armenian_hyphen = 16778634

// KEY_Armenian_ini is a representation of the C constant GDK_KEY_Armenian_ini.
const KEY_Armenian_ini = 16778603

// KEY_Armenian_je is a representation of the C constant GDK_KEY_Armenian_je.
const KEY_Armenian_je = 16778619

// KEY_Armenian_ke is a representation of the C constant GDK_KEY_Armenian_ke.
const KEY_Armenian_ke = 16778628

// KEY_Armenian_ken is a representation of the C constant GDK_KEY_Armenian_ken.
const KEY_Armenian_ken = 16778607

// KEY_Armenian_khe is a representation of the C constant GDK_KEY_Armenian_khe.
const KEY_Armenian_khe = 16778605

// KEY_Armenian_ligature_ew is a representation of the C constant GDK_KEY_Armenian_ligature_ew.
const KEY_Armenian_ligature_ew = 16778631

// KEY_Armenian_lyun is a representation of the C constant GDK_KEY_Armenian_lyun.
const KEY_Armenian_lyun = 16778604

// KEY_Armenian_men is a representation of the C constant GDK_KEY_Armenian_men.
const KEY_Armenian_men = 16778612

// KEY_Armenian_nu is a representation of the C constant GDK_KEY_Armenian_nu.
const KEY_Armenian_nu = 16778614

// KEY_Armenian_o is a representation of the C constant GDK_KEY_Armenian_o.
const KEY_Armenian_o = 16778629

// KEY_Armenian_paruyk is a representation of the C constant GDK_KEY_Armenian_paruyk.
const KEY_Armenian_paruyk = 16778590

// KEY_Armenian_pe is a representation of the C constant GDK_KEY_Armenian_pe.
const KEY_Armenian_pe = 16778618

// KEY_Armenian_pyur is a representation of the C constant GDK_KEY_Armenian_pyur.
const KEY_Armenian_pyur = 16778627

// KEY_Armenian_question is a representation of the C constant GDK_KEY_Armenian_question.
const KEY_Armenian_question = 16778590

// KEY_Armenian_ra is a representation of the C constant GDK_KEY_Armenian_ra.
const KEY_Armenian_ra = 16778620

// KEY_Armenian_re is a representation of the C constant GDK_KEY_Armenian_re.
const KEY_Armenian_re = 16778624

// KEY_Armenian_se is a representation of the C constant GDK_KEY_Armenian_se.
const KEY_Armenian_se = 16778621

// KEY_Armenian_separation_mark is a representation of the C constant GDK_KEY_Armenian_separation_mark.
const KEY_Armenian_separation_mark = 16778589

// KEY_Armenian_sha is a representation of the C constant GDK_KEY_Armenian_sha.
const KEY_Armenian_sha = 16778615

// KEY_Armenian_shesht is a representation of the C constant GDK_KEY_Armenian_shesht.
const KEY_Armenian_shesht = 16778587

// KEY_Armenian_tche is a representation of the C constant GDK_KEY_Armenian_tche.
const KEY_Armenian_tche = 16778611

// KEY_Armenian_to is a representation of the C constant GDK_KEY_Armenian_to.
const KEY_Armenian_to = 16778601

// KEY_Armenian_tsa is a representation of the C constant GDK_KEY_Armenian_tsa.
const KEY_Armenian_tsa = 16778606

// KEY_Armenian_tso is a representation of the C constant GDK_KEY_Armenian_tso.
const KEY_Armenian_tso = 16778625

// KEY_Armenian_tyun is a representation of the C constant GDK_KEY_Armenian_tyun.
const KEY_Armenian_tyun = 16778623

// KEY_Armenian_verjaket is a representation of the C constant GDK_KEY_Armenian_verjaket.
const KEY_Armenian_verjaket = 16778633

// KEY_Armenian_vev is a representation of the C constant GDK_KEY_Armenian_vev.
const KEY_Armenian_vev = 16778622

// KEY_Armenian_vo is a representation of the C constant GDK_KEY_Armenian_vo.
const KEY_Armenian_vo = 16778616

// KEY_Armenian_vyun is a representation of the C constant GDK_KEY_Armenian_vyun.
const KEY_Armenian_vyun = 16778626

// KEY_Armenian_yech is a representation of the C constant GDK_KEY_Armenian_yech.
const KEY_Armenian_yech = 16778597

// KEY_Armenian_yentamna is a representation of the C constant GDK_KEY_Armenian_yentamna.
const KEY_Armenian_yentamna = 16778634

// KEY_Armenian_za is a representation of the C constant GDK_KEY_Armenian_za.
const KEY_Armenian_za = 16778598

// KEY_Armenian_zhe is a representation of the C constant GDK_KEY_Armenian_zhe.
const KEY_Armenian_zhe = 16778602

// KEY_Atilde is a representation of the C constant GDK_KEY_Atilde.
const KEY_Atilde = 195

// KEY_AudibleBell_Enable is a representation of the C constant GDK_KEY_AudibleBell_Enable.
const KEY_AudibleBell_Enable = 65146

// KEY_AudioCycleTrack is a representation of the C constant GDK_KEY_AudioCycleTrack.
const KEY_AudioCycleTrack = 269025179

// KEY_AudioForward is a representation of the C constant GDK_KEY_AudioForward.
const KEY_AudioForward = 269025175

// KEY_AudioLowerVolume is a representation of the C constant GDK_KEY_AudioLowerVolume.
const KEY_AudioLowerVolume = 269025041

// KEY_AudioMedia is a representation of the C constant GDK_KEY_AudioMedia.
const KEY_AudioMedia = 269025074

// KEY_AudioMicMute is a representation of the C constant GDK_KEY_AudioMicMute.
const KEY_AudioMicMute = 269025202

// KEY_AudioMute is a representation of the C constant GDK_KEY_AudioMute.
const KEY_AudioMute = 269025042

// KEY_AudioNext is a representation of the C constant GDK_KEY_AudioNext.
const KEY_AudioNext = 269025047

// KEY_AudioPause is a representation of the C constant GDK_KEY_AudioPause.
const KEY_AudioPause = 269025073

// KEY_AudioPlay is a representation of the C constant GDK_KEY_AudioPlay.
const KEY_AudioPlay = 269025044

// KEY_AudioPrev is a representation of the C constant GDK_KEY_AudioPrev.
const KEY_AudioPrev = 269025046

// KEY_AudioRaiseVolume is a representation of the C constant GDK_KEY_AudioRaiseVolume.
const KEY_AudioRaiseVolume = 269025043

// KEY_AudioRandomPlay is a representation of the C constant GDK_KEY_AudioRandomPlay.
const KEY_AudioRandomPlay = 269025177

// KEY_AudioRecord is a representation of the C constant GDK_KEY_AudioRecord.
const KEY_AudioRecord = 269025052

// KEY_AudioRepeat is a representation of the C constant GDK_KEY_AudioRepeat.
const KEY_AudioRepeat = 269025176

// KEY_AudioRewind is a representation of the C constant GDK_KEY_AudioRewind.
const KEY_AudioRewind = 269025086

// KEY_AudioStop is a representation of the C constant GDK_KEY_AudioStop.
const KEY_AudioStop = 269025045

// KEY_Away is a representation of the C constant GDK_KEY_Away.
const KEY_Away = 269025165

// KEY_B is a representation of the C constant GDK_KEY_B.
const KEY_B = 66

// KEY_Babovedot is a representation of the C constant GDK_KEY_Babovedot.
const KEY_Babovedot = 16784898

// KEY_Back is a representation of the C constant GDK_KEY_Back.
const KEY_Back = 269025062

// KEY_BackForward is a representation of the C constant GDK_KEY_BackForward.
const KEY_BackForward = 269025087

// KEY_BackSpace is a representation of the C constant GDK_KEY_BackSpace.
const KEY_BackSpace = 65288

// KEY_Battery is a representation of the C constant GDK_KEY_Battery.
const KEY_Battery = 269025171

// KEY_Begin is a representation of the C constant GDK_KEY_Begin.
const KEY_Begin = 65368

// KEY_Blue is a representation of the C constant GDK_KEY_Blue.
const KEY_Blue = 269025190

// KEY_Bluetooth is a representation of the C constant GDK_KEY_Bluetooth.
const KEY_Bluetooth = 269025172

// KEY_Book is a representation of the C constant GDK_KEY_Book.
const KEY_Book = 269025106

// KEY_BounceKeys_Enable is a representation of the C constant GDK_KEY_BounceKeys_Enable.
const KEY_BounceKeys_Enable = 65140

// KEY_Break is a representation of the C constant GDK_KEY_Break.
const KEY_Break = 65387

// KEY_BrightnessAdjust is a representation of the C constant GDK_KEY_BrightnessAdjust.
const KEY_BrightnessAdjust = 269025083

// KEY_Byelorussian_SHORTU is a representation of the C constant GDK_KEY_Byelorussian_SHORTU.
const KEY_Byelorussian_SHORTU = 1726

// KEY_Byelorussian_shortu is a representation of the C constant GDK_KEY_Byelorussian_shortu.
const KEY_Byelorussian_shortu = 1710

// KEY_C is a representation of the C constant GDK_KEY_C.
const KEY_C = 67

// KEY_CD is a representation of the C constant GDK_KEY_CD.
const KEY_CD = 269025107

// KEY_CH is a representation of the C constant GDK_KEY_CH.
const KEY_CH = 65186

// KEY_C_H is a representation of the C constant GDK_KEY_C_H.
const KEY_C_H = 65189

// KEY_C_h is a representation of the C constant GDK_KEY_C_h.
const KEY_C_h = 65188

// KEY_Cabovedot is a representation of the C constant GDK_KEY_Cabovedot.
const KEY_Cabovedot = 709

// KEY_Cacute is a representation of the C constant GDK_KEY_Cacute.
const KEY_Cacute = 454

// KEY_Calculator is a representation of the C constant GDK_KEY_Calculator.
const KEY_Calculator = 269025053

// KEY_Calendar is a representation of the C constant GDK_KEY_Calendar.
const KEY_Calendar = 269025056

// KEY_Cancel is a representation of the C constant GDK_KEY_Cancel.
const KEY_Cancel = 65385

// KEY_Caps_Lock is a representation of the C constant GDK_KEY_Caps_Lock.
const KEY_Caps_Lock = 65509

// KEY_Ccaron is a representation of the C constant GDK_KEY_Ccaron.
const KEY_Ccaron = 456

// KEY_Ccedilla is a representation of the C constant GDK_KEY_Ccedilla.
const KEY_Ccedilla = 199

// KEY_Ccircumflex is a representation of the C constant GDK_KEY_Ccircumflex.
const KEY_Ccircumflex = 710

// KEY_Ch is a representation of the C constant GDK_KEY_Ch.
const KEY_Ch = 65185

// KEY_Clear is a representation of the C constant GDK_KEY_Clear.
const KEY_Clear = 65291

// KEY_ClearGrab is a representation of the C constant GDK_KEY_ClearGrab.
const KEY_ClearGrab = 269024801

// KEY_Close is a representation of the C constant GDK_KEY_Close.
const KEY_Close = 269025110

// KEY_Codeinput is a representation of the C constant GDK_KEY_Codeinput.
const KEY_Codeinput = 65335

// KEY_ColonSign is a representation of the C constant GDK_KEY_ColonSign.
const KEY_ColonSign = 16785569

// KEY_Community is a representation of the C constant GDK_KEY_Community.
const KEY_Community = 269025085

// KEY_ContrastAdjust is a representation of the C constant GDK_KEY_ContrastAdjust.
const KEY_ContrastAdjust = 269025058

// KEY_Control_L is a representation of the C constant GDK_KEY_Control_L.
const KEY_Control_L = 65507

// KEY_Control_R is a representation of the C constant GDK_KEY_Control_R.
const KEY_Control_R = 65508

// KEY_Copy is a representation of the C constant GDK_KEY_Copy.
const KEY_Copy = 269025111

// KEY_CruzeiroSign is a representation of the C constant GDK_KEY_CruzeiroSign.
const KEY_CruzeiroSign = 16785570

// KEY_Cut is a representation of the C constant GDK_KEY_Cut.
const KEY_Cut = 269025112

// KEY_CycleAngle is a representation of the C constant GDK_KEY_CycleAngle.
const KEY_CycleAngle = 269025180

// KEY_Cyrillic_A is a representation of the C constant GDK_KEY_Cyrillic_A.
const KEY_Cyrillic_A = 1761

// KEY_Cyrillic_BE is a representation of the C constant GDK_KEY_Cyrillic_BE.
const KEY_Cyrillic_BE = 1762

// KEY_Cyrillic_CHE is a representation of the C constant GDK_KEY_Cyrillic_CHE.
const KEY_Cyrillic_CHE = 1790

// KEY_Cyrillic_CHE_descender is a representation of the C constant GDK_KEY_Cyrillic_CHE_descender.
const KEY_Cyrillic_CHE_descender = 16778422

// KEY_Cyrillic_CHE_vertstroke is a representation of the C constant GDK_KEY_Cyrillic_CHE_vertstroke.
const KEY_Cyrillic_CHE_vertstroke = 16778424

// KEY_Cyrillic_DE is a representation of the C constant GDK_KEY_Cyrillic_DE.
const KEY_Cyrillic_DE = 1764

// KEY_Cyrillic_DZHE is a representation of the C constant GDK_KEY_Cyrillic_DZHE.
const KEY_Cyrillic_DZHE = 1727

// KEY_Cyrillic_E is a representation of the C constant GDK_KEY_Cyrillic_E.
const KEY_Cyrillic_E = 1788

// KEY_Cyrillic_EF is a representation of the C constant GDK_KEY_Cyrillic_EF.
const KEY_Cyrillic_EF = 1766

// KEY_Cyrillic_EL is a representation of the C constant GDK_KEY_Cyrillic_EL.
const KEY_Cyrillic_EL = 1772

// KEY_Cyrillic_EM is a representation of the C constant GDK_KEY_Cyrillic_EM.
const KEY_Cyrillic_EM = 1773

// KEY_Cyrillic_EN is a representation of the C constant GDK_KEY_Cyrillic_EN.
const KEY_Cyrillic_EN = 1774

// KEY_Cyrillic_EN_descender is a representation of the C constant GDK_KEY_Cyrillic_EN_descender.
const KEY_Cyrillic_EN_descender = 16778402

// KEY_Cyrillic_ER is a representation of the C constant GDK_KEY_Cyrillic_ER.
const KEY_Cyrillic_ER = 1778

// KEY_Cyrillic_ES is a representation of the C constant GDK_KEY_Cyrillic_ES.
const KEY_Cyrillic_ES = 1779

// KEY_Cyrillic_GHE is a representation of the C constant GDK_KEY_Cyrillic_GHE.
const KEY_Cyrillic_GHE = 1767

// KEY_Cyrillic_GHE_bar is a representation of the C constant GDK_KEY_Cyrillic_GHE_bar.
const KEY_Cyrillic_GHE_bar = 16778386

// KEY_Cyrillic_HA is a representation of the C constant GDK_KEY_Cyrillic_HA.
const KEY_Cyrillic_HA = 1768

// KEY_Cyrillic_HARDSIGN is a representation of the C constant GDK_KEY_Cyrillic_HARDSIGN.
const KEY_Cyrillic_HARDSIGN = 1791

// KEY_Cyrillic_HA_descender is a representation of the C constant GDK_KEY_Cyrillic_HA_descender.
const KEY_Cyrillic_HA_descender = 16778418

// KEY_Cyrillic_I is a representation of the C constant GDK_KEY_Cyrillic_I.
const KEY_Cyrillic_I = 1769

// KEY_Cyrillic_IE is a representation of the C constant GDK_KEY_Cyrillic_IE.
const KEY_Cyrillic_IE = 1765

// KEY_Cyrillic_IO is a representation of the C constant GDK_KEY_Cyrillic_IO.
const KEY_Cyrillic_IO = 1715

// KEY_Cyrillic_I_macron is a representation of the C constant GDK_KEY_Cyrillic_I_macron.
const KEY_Cyrillic_I_macron = 16778466

// KEY_Cyrillic_JE is a representation of the C constant GDK_KEY_Cyrillic_JE.
const KEY_Cyrillic_JE = 1720

// KEY_Cyrillic_KA is a representation of the C constant GDK_KEY_Cyrillic_KA.
const KEY_Cyrillic_KA = 1771

// KEY_Cyrillic_KA_descender is a representation of the C constant GDK_KEY_Cyrillic_KA_descender.
const KEY_Cyrillic_KA_descender = 16778394

// KEY_Cyrillic_KA_vertstroke is a representation of the C constant GDK_KEY_Cyrillic_KA_vertstroke.
const KEY_Cyrillic_KA_vertstroke = 16778396

// KEY_Cyrillic_LJE is a representation of the C constant GDK_KEY_Cyrillic_LJE.
const KEY_Cyrillic_LJE = 1721

// KEY_Cyrillic_NJE is a representation of the C constant GDK_KEY_Cyrillic_NJE.
const KEY_Cyrillic_NJE = 1722

// KEY_Cyrillic_O is a representation of the C constant GDK_KEY_Cyrillic_O.
const KEY_Cyrillic_O = 1775

// KEY_Cyrillic_O_bar is a representation of the C constant GDK_KEY_Cyrillic_O_bar.
const KEY_Cyrillic_O_bar = 16778472

// KEY_Cyrillic_PE is a representation of the C constant GDK_KEY_Cyrillic_PE.
const KEY_Cyrillic_PE = 1776

// KEY_Cyrillic_SCHWA is a representation of the C constant GDK_KEY_Cyrillic_SCHWA.
const KEY_Cyrillic_SCHWA = 16778456

// KEY_Cyrillic_SHA is a representation of the C constant GDK_KEY_Cyrillic_SHA.
const KEY_Cyrillic_SHA = 1787

// KEY_Cyrillic_SHCHA is a representation of the C constant GDK_KEY_Cyrillic_SHCHA.
const KEY_Cyrillic_SHCHA = 1789

// KEY_Cyrillic_SHHA is a representation of the C constant GDK_KEY_Cyrillic_SHHA.
const KEY_Cyrillic_SHHA = 16778426

// KEY_Cyrillic_SHORTI is a representation of the C constant GDK_KEY_Cyrillic_SHORTI.
const KEY_Cyrillic_SHORTI = 1770

// KEY_Cyrillic_SOFTSIGN is a representation of the C constant GDK_KEY_Cyrillic_SOFTSIGN.
const KEY_Cyrillic_SOFTSIGN = 1784

// KEY_Cyrillic_TE is a representation of the C constant GDK_KEY_Cyrillic_TE.
const KEY_Cyrillic_TE = 1780

// KEY_Cyrillic_TSE is a representation of the C constant GDK_KEY_Cyrillic_TSE.
const KEY_Cyrillic_TSE = 1763

// KEY_Cyrillic_U is a representation of the C constant GDK_KEY_Cyrillic_U.
const KEY_Cyrillic_U = 1781

// KEY_Cyrillic_U_macron is a representation of the C constant GDK_KEY_Cyrillic_U_macron.
const KEY_Cyrillic_U_macron = 16778478

// KEY_Cyrillic_U_straight is a representation of the C constant GDK_KEY_Cyrillic_U_straight.
const KEY_Cyrillic_U_straight = 16778414

// KEY_Cyrillic_U_straight_bar is a representation of the C constant GDK_KEY_Cyrillic_U_straight_bar.
const KEY_Cyrillic_U_straight_bar = 16778416

// KEY_Cyrillic_VE is a representation of the C constant GDK_KEY_Cyrillic_VE.
const KEY_Cyrillic_VE = 1783

// KEY_Cyrillic_YA is a representation of the C constant GDK_KEY_Cyrillic_YA.
const KEY_Cyrillic_YA = 1777

// KEY_Cyrillic_YERU is a representation of the C constant GDK_KEY_Cyrillic_YERU.
const KEY_Cyrillic_YERU = 1785

// KEY_Cyrillic_YU is a representation of the C constant GDK_KEY_Cyrillic_YU.
const KEY_Cyrillic_YU = 1760

// KEY_Cyrillic_ZE is a representation of the C constant GDK_KEY_Cyrillic_ZE.
const KEY_Cyrillic_ZE = 1786

// KEY_Cyrillic_ZHE is a representation of the C constant GDK_KEY_Cyrillic_ZHE.
const KEY_Cyrillic_ZHE = 1782

// KEY_Cyrillic_ZHE_descender is a representation of the C constant GDK_KEY_Cyrillic_ZHE_descender.
const KEY_Cyrillic_ZHE_descender = 16778390

// KEY_Cyrillic_a is a representation of the C constant GDK_KEY_Cyrillic_a.
const KEY_Cyrillic_a = 1729

// KEY_Cyrillic_be is a representation of the C constant GDK_KEY_Cyrillic_be.
const KEY_Cyrillic_be = 1730

// KEY_Cyrillic_che is a representation of the C constant GDK_KEY_Cyrillic_che.
const KEY_Cyrillic_che = 1758

// KEY_Cyrillic_che_descender is a representation of the C constant GDK_KEY_Cyrillic_che_descender.
const KEY_Cyrillic_che_descender = 16778423

// KEY_Cyrillic_che_vertstroke is a representation of the C constant GDK_KEY_Cyrillic_che_vertstroke.
const KEY_Cyrillic_che_vertstroke = 16778425

// KEY_Cyrillic_de is a representation of the C constant GDK_KEY_Cyrillic_de.
const KEY_Cyrillic_de = 1732

// KEY_Cyrillic_dzhe is a representation of the C constant GDK_KEY_Cyrillic_dzhe.
const KEY_Cyrillic_dzhe = 1711

// KEY_Cyrillic_e is a representation of the C constant GDK_KEY_Cyrillic_e.
const KEY_Cyrillic_e = 1756

// KEY_Cyrillic_ef is a representation of the C constant GDK_KEY_Cyrillic_ef.
const KEY_Cyrillic_ef = 1734

// KEY_Cyrillic_el is a representation of the C constant GDK_KEY_Cyrillic_el.
const KEY_Cyrillic_el = 1740

// KEY_Cyrillic_em is a representation of the C constant GDK_KEY_Cyrillic_em.
const KEY_Cyrillic_em = 1741

// KEY_Cyrillic_en is a representation of the C constant GDK_KEY_Cyrillic_en.
const KEY_Cyrillic_en = 1742

// KEY_Cyrillic_en_descender is a representation of the C constant GDK_KEY_Cyrillic_en_descender.
const KEY_Cyrillic_en_descender = 16778403

// KEY_Cyrillic_er is a representation of the C constant GDK_KEY_Cyrillic_er.
const KEY_Cyrillic_er = 1746

// KEY_Cyrillic_es is a representation of the C constant GDK_KEY_Cyrillic_es.
const KEY_Cyrillic_es = 1747

// KEY_Cyrillic_ghe is a representation of the C constant GDK_KEY_Cyrillic_ghe.
const KEY_Cyrillic_ghe = 1735

// KEY_Cyrillic_ghe_bar is a representation of the C constant GDK_KEY_Cyrillic_ghe_bar.
const KEY_Cyrillic_ghe_bar = 16778387

// KEY_Cyrillic_ha is a representation of the C constant GDK_KEY_Cyrillic_ha.
const KEY_Cyrillic_ha = 1736

// KEY_Cyrillic_ha_descender is a representation of the C constant GDK_KEY_Cyrillic_ha_descender.
const KEY_Cyrillic_ha_descender = 16778419

// KEY_Cyrillic_hardsign is a representation of the C constant GDK_KEY_Cyrillic_hardsign.
const KEY_Cyrillic_hardsign = 1759

// KEY_Cyrillic_i is a representation of the C constant GDK_KEY_Cyrillic_i.
const KEY_Cyrillic_i = 1737

// KEY_Cyrillic_i_macron is a representation of the C constant GDK_KEY_Cyrillic_i_macron.
const KEY_Cyrillic_i_macron = 16778467

// KEY_Cyrillic_ie is a representation of the C constant GDK_KEY_Cyrillic_ie.
const KEY_Cyrillic_ie = 1733

// KEY_Cyrillic_io is a representation of the C constant GDK_KEY_Cyrillic_io.
const KEY_Cyrillic_io = 1699

// KEY_Cyrillic_je is a representation of the C constant GDK_KEY_Cyrillic_je.
const KEY_Cyrillic_je = 1704

// KEY_Cyrillic_ka is a representation of the C constant GDK_KEY_Cyrillic_ka.
const KEY_Cyrillic_ka = 1739

// KEY_Cyrillic_ka_descender is a representation of the C constant GDK_KEY_Cyrillic_ka_descender.
const KEY_Cyrillic_ka_descender = 16778395

// KEY_Cyrillic_ka_vertstroke is a representation of the C constant GDK_KEY_Cyrillic_ka_vertstroke.
const KEY_Cyrillic_ka_vertstroke = 16778397

// KEY_Cyrillic_lje is a representation of the C constant GDK_KEY_Cyrillic_lje.
const KEY_Cyrillic_lje = 1705

// KEY_Cyrillic_nje is a representation of the C constant GDK_KEY_Cyrillic_nje.
const KEY_Cyrillic_nje = 1706

// KEY_Cyrillic_o is a representation of the C constant GDK_KEY_Cyrillic_o.
const KEY_Cyrillic_o = 1743

// KEY_Cyrillic_o_bar is a representation of the C constant GDK_KEY_Cyrillic_o_bar.
const KEY_Cyrillic_o_bar = 16778473

// KEY_Cyrillic_pe is a representation of the C constant GDK_KEY_Cyrillic_pe.
const KEY_Cyrillic_pe = 1744

// KEY_Cyrillic_schwa is a representation of the C constant GDK_KEY_Cyrillic_schwa.
const KEY_Cyrillic_schwa = 16778457

// KEY_Cyrillic_sha is a representation of the C constant GDK_KEY_Cyrillic_sha.
const KEY_Cyrillic_sha = 1755

// KEY_Cyrillic_shcha is a representation of the C constant GDK_KEY_Cyrillic_shcha.
const KEY_Cyrillic_shcha = 1757

// KEY_Cyrillic_shha is a representation of the C constant GDK_KEY_Cyrillic_shha.
const KEY_Cyrillic_shha = 16778427

// KEY_Cyrillic_shorti is a representation of the C constant GDK_KEY_Cyrillic_shorti.
const KEY_Cyrillic_shorti = 1738

// KEY_Cyrillic_softsign is a representation of the C constant GDK_KEY_Cyrillic_softsign.
const KEY_Cyrillic_softsign = 1752

// KEY_Cyrillic_te is a representation of the C constant GDK_KEY_Cyrillic_te.
const KEY_Cyrillic_te = 1748

// KEY_Cyrillic_tse is a representation of the C constant GDK_KEY_Cyrillic_tse.
const KEY_Cyrillic_tse = 1731

// KEY_Cyrillic_u is a representation of the C constant GDK_KEY_Cyrillic_u.
const KEY_Cyrillic_u = 1749

// KEY_Cyrillic_u_macron is a representation of the C constant GDK_KEY_Cyrillic_u_macron.
const KEY_Cyrillic_u_macron = 16778479

// KEY_Cyrillic_u_straight is a representation of the C constant GDK_KEY_Cyrillic_u_straight.
const KEY_Cyrillic_u_straight = 16778415

// KEY_Cyrillic_u_straight_bar is a representation of the C constant GDK_KEY_Cyrillic_u_straight_bar.
const KEY_Cyrillic_u_straight_bar = 16778417

// KEY_Cyrillic_ve is a representation of the C constant GDK_KEY_Cyrillic_ve.
const KEY_Cyrillic_ve = 1751

// KEY_Cyrillic_ya is a representation of the C constant GDK_KEY_Cyrillic_ya.
const KEY_Cyrillic_ya = 1745

// KEY_Cyrillic_yeru is a representation of the C constant GDK_KEY_Cyrillic_yeru.
const KEY_Cyrillic_yeru = 1753

// KEY_Cyrillic_yu is a representation of the C constant GDK_KEY_Cyrillic_yu.
const KEY_Cyrillic_yu = 1728

// KEY_Cyrillic_ze is a representation of the C constant GDK_KEY_Cyrillic_ze.
const KEY_Cyrillic_ze = 1754

// KEY_Cyrillic_zhe is a representation of the C constant GDK_KEY_Cyrillic_zhe.
const KEY_Cyrillic_zhe = 1750

// KEY_Cyrillic_zhe_descender is a representation of the C constant GDK_KEY_Cyrillic_zhe_descender.
const KEY_Cyrillic_zhe_descender = 16778391

// KEY_D is a representation of the C constant GDK_KEY_D.
const KEY_D = 68

// KEY_DOS is a representation of the C constant GDK_KEY_DOS.
const KEY_DOS = 269025114

// KEY_Dabovedot is a representation of the C constant GDK_KEY_Dabovedot.
const KEY_Dabovedot = 16784906

// KEY_Dcaron is a representation of the C constant GDK_KEY_Dcaron.
const KEY_Dcaron = 463

// KEY_Delete is a representation of the C constant GDK_KEY_Delete.
const KEY_Delete = 65535

// KEY_Display is a representation of the C constant GDK_KEY_Display.
const KEY_Display = 269025113

// KEY_Documents is a representation of the C constant GDK_KEY_Documents.
const KEY_Documents = 269025115

// KEY_DongSign is a representation of the C constant GDK_KEY_DongSign.
const KEY_DongSign = 16785579

// KEY_Down is a representation of the C constant GDK_KEY_Down.
const KEY_Down = 65364

// KEY_Dstroke is a representation of the C constant GDK_KEY_Dstroke.
const KEY_Dstroke = 464

// KEY_E is a representation of the C constant GDK_KEY_E.
const KEY_E = 69

// KEY_ENG is a representation of the C constant GDK_KEY_ENG.
const KEY_ENG = 957

// KEY_ETH is a representation of the C constant GDK_KEY_ETH.
const KEY_ETH = 208

// KEY_EZH is a representation of the C constant GDK_KEY_EZH.
const KEY_EZH = 16777655

// KEY_Eabovedot is a representation of the C constant GDK_KEY_Eabovedot.
const KEY_Eabovedot = 972

// KEY_Eacute is a representation of the C constant GDK_KEY_Eacute.
const KEY_Eacute = 201

// KEY_Ebelowdot is a representation of the C constant GDK_KEY_Ebelowdot.
const KEY_Ebelowdot = 16785080

// KEY_Ecaron is a representation of the C constant GDK_KEY_Ecaron.
const KEY_Ecaron = 460

// KEY_Ecircumflex is a representation of the C constant GDK_KEY_Ecircumflex.
const KEY_Ecircumflex = 202

// KEY_Ecircumflexacute is a representation of the C constant GDK_KEY_Ecircumflexacute.
const KEY_Ecircumflexacute = 16785086

// KEY_Ecircumflexbelowdot is a representation of the C constant GDK_KEY_Ecircumflexbelowdot.
const KEY_Ecircumflexbelowdot = 16785094

// KEY_Ecircumflexgrave is a representation of the C constant GDK_KEY_Ecircumflexgrave.
const KEY_Ecircumflexgrave = 16785088

// KEY_Ecircumflexhook is a representation of the C constant GDK_KEY_Ecircumflexhook.
const KEY_Ecircumflexhook = 16785090

// KEY_Ecircumflextilde is a representation of the C constant GDK_KEY_Ecircumflextilde.
const KEY_Ecircumflextilde = 16785092

// KEY_EcuSign is a representation of the C constant GDK_KEY_EcuSign.
const KEY_EcuSign = 16785568

// KEY_Ediaeresis is a representation of the C constant GDK_KEY_Ediaeresis.
const KEY_Ediaeresis = 203

// KEY_Egrave is a representation of the C constant GDK_KEY_Egrave.
const KEY_Egrave = 200

// KEY_Ehook is a representation of the C constant GDK_KEY_Ehook.
const KEY_Ehook = 16785082

// KEY_Eisu_Shift is a representation of the C constant GDK_KEY_Eisu_Shift.
const KEY_Eisu_Shift = 65327

// KEY_Eisu_toggle is a representation of the C constant GDK_KEY_Eisu_toggle.
const KEY_Eisu_toggle = 65328

// KEY_Eject is a representation of the C constant GDK_KEY_Eject.
const KEY_Eject = 269025068

// KEY_Emacron is a representation of the C constant GDK_KEY_Emacron.
const KEY_Emacron = 938

// KEY_End is a representation of the C constant GDK_KEY_End.
const KEY_End = 65367

// KEY_Eogonek is a representation of the C constant GDK_KEY_Eogonek.
const KEY_Eogonek = 458

// KEY_Escape is a representation of the C constant GDK_KEY_Escape.
const KEY_Escape = 65307

// KEY_Eth is a representation of the C constant GDK_KEY_Eth.
const KEY_Eth = 208

// KEY_Etilde is a representation of the C constant GDK_KEY_Etilde.
const KEY_Etilde = 16785084

// KEY_EuroSign is a representation of the C constant GDK_KEY_EuroSign.
const KEY_EuroSign = 8364

// KEY_Excel is a representation of the C constant GDK_KEY_Excel.
const KEY_Excel = 269025116

// KEY_Execute is a representation of the C constant GDK_KEY_Execute.
const KEY_Execute = 65378

// KEY_Explorer is a representation of the C constant GDK_KEY_Explorer.
const KEY_Explorer = 269025117

// KEY_F is a representation of the C constant GDK_KEY_F.
const KEY_F = 70

// KEY_F1 is a representation of the C constant GDK_KEY_F1.
const KEY_F1 = 65470

// KEY_F10 is a representation of the C constant GDK_KEY_F10.
const KEY_F10 = 65479

// KEY_F11 is a representation of the C constant GDK_KEY_F11.
const KEY_F11 = 65480

// KEY_F12 is a representation of the C constant GDK_KEY_F12.
const KEY_F12 = 65481

// KEY_F13 is a representation of the C constant GDK_KEY_F13.
const KEY_F13 = 65482

// KEY_F14 is a representation of the C constant GDK_KEY_F14.
const KEY_F14 = 65483

// KEY_F15 is a representation of the C constant GDK_KEY_F15.
const KEY_F15 = 65484

// KEY_F16 is a representation of the C constant GDK_KEY_F16.
const KEY_F16 = 65485

// KEY_F17 is a representation of the C constant GDK_KEY_F17.
const KEY_F17 = 65486

// KEY_F18 is a representation of the C constant GDK_KEY_F18.
const KEY_F18 = 65487

// KEY_F19 is a representation of the C constant GDK_KEY_F19.
const KEY_F19 = 65488

// KEY_F2 is a representation of the C constant GDK_KEY_F2.
const KEY_F2 = 65471

// KEY_F20 is a representation of the C constant GDK_KEY_F20.
const KEY_F20 = 65489

// KEY_F21 is a representation of the C constant GDK_KEY_F21.
const KEY_F21 = 65490

// KEY_F22 is a representation of the C constant GDK_KEY_F22.
const KEY_F22 = 65491

// KEY_F23 is a representation of the C constant GDK_KEY_F23.
const KEY_F23 = 65492

// KEY_F24 is a representation of the C constant GDK_KEY_F24.
const KEY_F24 = 65493

// KEY_F25 is a representation of the C constant GDK_KEY_F25.
const KEY_F25 = 65494

// KEY_F26 is a representation of the C constant GDK_KEY_F26.
const KEY_F26 = 65495

// KEY_F27 is a representation of the C constant GDK_KEY_F27.
const KEY_F27 = 65496

// KEY_F28 is a representation of the C constant GDK_KEY_F28.
const KEY_F28 = 65497

// KEY_F29 is a representation of the C constant GDK_KEY_F29.
const KEY_F29 = 65498

// KEY_F3 is a representation of the C constant GDK_KEY_F3.
const KEY_F3 = 65472

// KEY_F30 is a representation of the C constant GDK_KEY_F30.
const KEY_F30 = 65499

// KEY_F31 is a representation of the C constant GDK_KEY_F31.
const KEY_F31 = 65500

// KEY_F32 is a representation of the C constant GDK_KEY_F32.
const KEY_F32 = 65501

// KEY_F33 is a representation of the C constant GDK_KEY_F33.
const KEY_F33 = 65502

// KEY_F34 is a representation of the C constant GDK_KEY_F34.
const KEY_F34 = 65503

// KEY_F35 is a representation of the C constant GDK_KEY_F35.
const KEY_F35 = 65504

// KEY_F4 is a representation of the C constant GDK_KEY_F4.
const KEY_F4 = 65473

// KEY_F5 is a representation of the C constant GDK_KEY_F5.
const KEY_F5 = 65474

// KEY_F6 is a representation of the C constant GDK_KEY_F6.
const KEY_F6 = 65475

// KEY_F7 is a representation of the C constant GDK_KEY_F7.
const KEY_F7 = 65476

// KEY_F8 is a representation of the C constant GDK_KEY_F8.
const KEY_F8 = 65477

// KEY_F9 is a representation of the C constant GDK_KEY_F9.
const KEY_F9 = 65478

// KEY_FFrancSign is a representation of the C constant GDK_KEY_FFrancSign.
const KEY_FFrancSign = 16785571

// KEY_Fabovedot is a representation of the C constant GDK_KEY_Fabovedot.
const KEY_Fabovedot = 16784926

// KEY_Farsi_0 is a representation of the C constant GDK_KEY_Farsi_0.
const KEY_Farsi_0 = 16778992

// KEY_Farsi_1 is a representation of the C constant GDK_KEY_Farsi_1.
const KEY_Farsi_1 = 16778993

// KEY_Farsi_2 is a representation of the C constant GDK_KEY_Farsi_2.
const KEY_Farsi_2 = 16778994

// KEY_Farsi_3 is a representation of the C constant GDK_KEY_Farsi_3.
const KEY_Farsi_3 = 16778995

// KEY_Farsi_4 is a representation of the C constant GDK_KEY_Farsi_4.
const KEY_Farsi_4 = 16778996

// KEY_Farsi_5 is a representation of the C constant GDK_KEY_Farsi_5.
const KEY_Farsi_5 = 16778997

// KEY_Farsi_6 is a representation of the C constant GDK_KEY_Farsi_6.
const KEY_Farsi_6 = 16778998

// KEY_Farsi_7 is a representation of the C constant GDK_KEY_Farsi_7.
const KEY_Farsi_7 = 16778999

// KEY_Farsi_8 is a representation of the C constant GDK_KEY_Farsi_8.
const KEY_Farsi_8 = 16779000

// KEY_Farsi_9 is a representation of the C constant GDK_KEY_Farsi_9.
const KEY_Farsi_9 = 16779001

// KEY_Farsi_yeh is a representation of the C constant GDK_KEY_Farsi_yeh.
const KEY_Farsi_yeh = 16778956

// KEY_Favorites is a representation of the C constant GDK_KEY_Favorites.
const KEY_Favorites = 269025072

// KEY_Finance is a representation of the C constant GDK_KEY_Finance.
const KEY_Finance = 269025084

// KEY_Find is a representation of the C constant GDK_KEY_Find.
const KEY_Find = 65384

// KEY_First_Virtual_Screen is a representation of the C constant GDK_KEY_First_Virtual_Screen.
const KEY_First_Virtual_Screen = 65232

// KEY_Forward is a representation of the C constant GDK_KEY_Forward.
const KEY_Forward = 269025063

// KEY_FrameBack is a representation of the C constant GDK_KEY_FrameBack.
const KEY_FrameBack = 269025181

// KEY_FrameForward is a representation of the C constant GDK_KEY_FrameForward.
const KEY_FrameForward = 269025182

// KEY_G is a representation of the C constant GDK_KEY_G.
const KEY_G = 71

// KEY_Gabovedot is a representation of the C constant GDK_KEY_Gabovedot.
const KEY_Gabovedot = 725

// KEY_Game is a representation of the C constant GDK_KEY_Game.
const KEY_Game = 269025118

// KEY_Gbreve is a representation of the C constant GDK_KEY_Gbreve.
const KEY_Gbreve = 683

// KEY_Gcaron is a representation of the C constant GDK_KEY_Gcaron.
const KEY_Gcaron = 16777702

// KEY_Gcedilla is a representation of the C constant GDK_KEY_Gcedilla.
const KEY_Gcedilla = 939

// KEY_Gcircumflex is a representation of the C constant GDK_KEY_Gcircumflex.
const KEY_Gcircumflex = 728

// KEY_Georgian_an is a representation of the C constant GDK_KEY_Georgian_an.
const KEY_Georgian_an = 16781520

// KEY_Georgian_ban is a representation of the C constant GDK_KEY_Georgian_ban.
const KEY_Georgian_ban = 16781521

// KEY_Georgian_can is a representation of the C constant GDK_KEY_Georgian_can.
const KEY_Georgian_can = 16781546

// KEY_Georgian_char is a representation of the C constant GDK_KEY_Georgian_char.
const KEY_Georgian_char = 16781549

// KEY_Georgian_chin is a representation of the C constant GDK_KEY_Georgian_chin.
const KEY_Georgian_chin = 16781545

// KEY_Georgian_cil is a representation of the C constant GDK_KEY_Georgian_cil.
const KEY_Georgian_cil = 16781548

// KEY_Georgian_don is a representation of the C constant GDK_KEY_Georgian_don.
const KEY_Georgian_don = 16781523

// KEY_Georgian_en is a representation of the C constant GDK_KEY_Georgian_en.
const KEY_Georgian_en = 16781524

// KEY_Georgian_fi is a representation of the C constant GDK_KEY_Georgian_fi.
const KEY_Georgian_fi = 16781558

// KEY_Georgian_gan is a representation of the C constant GDK_KEY_Georgian_gan.
const KEY_Georgian_gan = 16781522

// KEY_Georgian_ghan is a representation of the C constant GDK_KEY_Georgian_ghan.
const KEY_Georgian_ghan = 16781542

// KEY_Georgian_hae is a representation of the C constant GDK_KEY_Georgian_hae.
const KEY_Georgian_hae = 16781552

// KEY_Georgian_har is a representation of the C constant GDK_KEY_Georgian_har.
const KEY_Georgian_har = 16781556

// KEY_Georgian_he is a representation of the C constant GDK_KEY_Georgian_he.
const KEY_Georgian_he = 16781553

// KEY_Georgian_hie is a representation of the C constant GDK_KEY_Georgian_hie.
const KEY_Georgian_hie = 16781554

// KEY_Georgian_hoe is a representation of the C constant GDK_KEY_Georgian_hoe.
const KEY_Georgian_hoe = 16781557

// KEY_Georgian_in is a representation of the C constant GDK_KEY_Georgian_in.
const KEY_Georgian_in = 16781528

// KEY_Georgian_jhan is a representation of the C constant GDK_KEY_Georgian_jhan.
const KEY_Georgian_jhan = 16781551

// KEY_Georgian_jil is a representation of the C constant GDK_KEY_Georgian_jil.
const KEY_Georgian_jil = 16781547

// KEY_Georgian_kan is a representation of the C constant GDK_KEY_Georgian_kan.
const KEY_Georgian_kan = 16781529

// KEY_Georgian_khar is a representation of the C constant GDK_KEY_Georgian_khar.
const KEY_Georgian_khar = 16781541

// KEY_Georgian_las is a representation of the C constant GDK_KEY_Georgian_las.
const KEY_Georgian_las = 16781530

// KEY_Georgian_man is a representation of the C constant GDK_KEY_Georgian_man.
const KEY_Georgian_man = 16781531

// KEY_Georgian_nar is a representation of the C constant GDK_KEY_Georgian_nar.
const KEY_Georgian_nar = 16781532

// KEY_Georgian_on is a representation of the C constant GDK_KEY_Georgian_on.
const KEY_Georgian_on = 16781533

// KEY_Georgian_par is a representation of the C constant GDK_KEY_Georgian_par.
const KEY_Georgian_par = 16781534

// KEY_Georgian_phar is a representation of the C constant GDK_KEY_Georgian_phar.
const KEY_Georgian_phar = 16781540

// KEY_Georgian_qar is a representation of the C constant GDK_KEY_Georgian_qar.
const KEY_Georgian_qar = 16781543

// KEY_Georgian_rae is a representation of the C constant GDK_KEY_Georgian_rae.
const KEY_Georgian_rae = 16781536

// KEY_Georgian_san is a representation of the C constant GDK_KEY_Georgian_san.
const KEY_Georgian_san = 16781537

// KEY_Georgian_shin is a representation of the C constant GDK_KEY_Georgian_shin.
const KEY_Georgian_shin = 16781544

// KEY_Georgian_tan is a representation of the C constant GDK_KEY_Georgian_tan.
const KEY_Georgian_tan = 16781527

// KEY_Georgian_tar is a representation of the C constant GDK_KEY_Georgian_tar.
const KEY_Georgian_tar = 16781538

// KEY_Georgian_un is a representation of the C constant GDK_KEY_Georgian_un.
const KEY_Georgian_un = 16781539

// KEY_Georgian_vin is a representation of the C constant GDK_KEY_Georgian_vin.
const KEY_Georgian_vin = 16781525

// KEY_Georgian_we is a representation of the C constant GDK_KEY_Georgian_we.
const KEY_Georgian_we = 16781555

// KEY_Georgian_xan is a representation of the C constant GDK_KEY_Georgian_xan.
const KEY_Georgian_xan = 16781550

// KEY_Georgian_zen is a representation of the C constant GDK_KEY_Georgian_zen.
const KEY_Georgian_zen = 16781526

// KEY_Georgian_zhar is a representation of the C constant GDK_KEY_Georgian_zhar.
const KEY_Georgian_zhar = 16781535

// KEY_Go is a representation of the C constant GDK_KEY_Go.
const KEY_Go = 269025119

// KEY_Greek_ALPHA is a representation of the C constant GDK_KEY_Greek_ALPHA.
const KEY_Greek_ALPHA = 1985

// KEY_Greek_ALPHAaccent is a representation of the C constant GDK_KEY_Greek_ALPHAaccent.
const KEY_Greek_ALPHAaccent = 1953

// KEY_Greek_BETA is a representation of the C constant GDK_KEY_Greek_BETA.
const KEY_Greek_BETA = 1986

// KEY_Greek_CHI is a representation of the C constant GDK_KEY_Greek_CHI.
const KEY_Greek_CHI = 2007

// KEY_Greek_DELTA is a representation of the C constant GDK_KEY_Greek_DELTA.
const KEY_Greek_DELTA = 1988

// KEY_Greek_EPSILON is a representation of the C constant GDK_KEY_Greek_EPSILON.
const KEY_Greek_EPSILON = 1989

// KEY_Greek_EPSILONaccent is a representation of the C constant GDK_KEY_Greek_EPSILONaccent.
const KEY_Greek_EPSILONaccent = 1954

// KEY_Greek_ETA is a representation of the C constant GDK_KEY_Greek_ETA.
const KEY_Greek_ETA = 1991

// KEY_Greek_ETAaccent is a representation of the C constant GDK_KEY_Greek_ETAaccent.
const KEY_Greek_ETAaccent = 1955

// KEY_Greek_GAMMA is a representation of the C constant GDK_KEY_Greek_GAMMA.
const KEY_Greek_GAMMA = 1987

// KEY_Greek_IOTA is a representation of the C constant GDK_KEY_Greek_IOTA.
const KEY_Greek_IOTA = 1993

// KEY_Greek_IOTAaccent is a representation of the C constant GDK_KEY_Greek_IOTAaccent.
const KEY_Greek_IOTAaccent = 1956

// KEY_Greek_IOTAdiaeresis is a representation of the C constant GDK_KEY_Greek_IOTAdiaeresis.
const KEY_Greek_IOTAdiaeresis = 1957

// KEY_Greek_IOTAdieresis is a representation of the C constant GDK_KEY_Greek_IOTAdieresis.
const KEY_Greek_IOTAdieresis = 1957

// KEY_Greek_KAPPA is a representation of the C constant GDK_KEY_Greek_KAPPA.
const KEY_Greek_KAPPA = 1994

// KEY_Greek_LAMBDA is a representation of the C constant GDK_KEY_Greek_LAMBDA.
const KEY_Greek_LAMBDA = 1995

// KEY_Greek_LAMDA is a representation of the C constant GDK_KEY_Greek_LAMDA.
const KEY_Greek_LAMDA = 1995

// KEY_Greek_MU is a representation of the C constant GDK_KEY_Greek_MU.
const KEY_Greek_MU = 1996

// KEY_Greek_NU is a representation of the C constant GDK_KEY_Greek_NU.
const KEY_Greek_NU = 1997

// KEY_Greek_OMEGA is a representation of the C constant GDK_KEY_Greek_OMEGA.
const KEY_Greek_OMEGA = 2009

// KEY_Greek_OMEGAaccent is a representation of the C constant GDK_KEY_Greek_OMEGAaccent.
const KEY_Greek_OMEGAaccent = 1963

// KEY_Greek_OMICRON is a representation of the C constant GDK_KEY_Greek_OMICRON.
const KEY_Greek_OMICRON = 1999

// KEY_Greek_OMICRONaccent is a representation of the C constant GDK_KEY_Greek_OMICRONaccent.
const KEY_Greek_OMICRONaccent = 1959

// KEY_Greek_PHI is a representation of the C constant GDK_KEY_Greek_PHI.
const KEY_Greek_PHI = 2006

// KEY_Greek_PI is a representation of the C constant GDK_KEY_Greek_PI.
const KEY_Greek_PI = 2000

// KEY_Greek_PSI is a representation of the C constant GDK_KEY_Greek_PSI.
const KEY_Greek_PSI = 2008

// KEY_Greek_RHO is a representation of the C constant GDK_KEY_Greek_RHO.
const KEY_Greek_RHO = 2001

// KEY_Greek_SIGMA is a representation of the C constant GDK_KEY_Greek_SIGMA.
const KEY_Greek_SIGMA = 2002

// KEY_Greek_TAU is a representation of the C constant GDK_KEY_Greek_TAU.
const KEY_Greek_TAU = 2004

// KEY_Greek_THETA is a representation of the C constant GDK_KEY_Greek_THETA.
const KEY_Greek_THETA = 1992

// KEY_Greek_UPSILON is a representation of the C constant GDK_KEY_Greek_UPSILON.
const KEY_Greek_UPSILON = 2005

// KEY_Greek_UPSILONaccent is a representation of the C constant GDK_KEY_Greek_UPSILONaccent.
const KEY_Greek_UPSILONaccent = 1960

// KEY_Greek_UPSILONdieresis is a representation of the C constant GDK_KEY_Greek_UPSILONdieresis.
const KEY_Greek_UPSILONdieresis = 1961

// KEY_Greek_XI is a representation of the C constant GDK_KEY_Greek_XI.
const KEY_Greek_XI = 1998

// KEY_Greek_ZETA is a representation of the C constant GDK_KEY_Greek_ZETA.
const KEY_Greek_ZETA = 1990

// KEY_Greek_accentdieresis is a representation of the C constant GDK_KEY_Greek_accentdieresis.
const KEY_Greek_accentdieresis = 1966

// KEY_Greek_alpha is a representation of the C constant GDK_KEY_Greek_alpha.
const KEY_Greek_alpha = 2017

// KEY_Greek_alphaaccent is a representation of the C constant GDK_KEY_Greek_alphaaccent.
const KEY_Greek_alphaaccent = 1969

// KEY_Greek_beta is a representation of the C constant GDK_KEY_Greek_beta.
const KEY_Greek_beta = 2018

// KEY_Greek_chi is a representation of the C constant GDK_KEY_Greek_chi.
const KEY_Greek_chi = 2039

// KEY_Greek_delta is a representation of the C constant GDK_KEY_Greek_delta.
const KEY_Greek_delta = 2020

// KEY_Greek_epsilon is a representation of the C constant GDK_KEY_Greek_epsilon.
const KEY_Greek_epsilon = 2021

// KEY_Greek_epsilonaccent is a representation of the C constant GDK_KEY_Greek_epsilonaccent.
const KEY_Greek_epsilonaccent = 1970

// KEY_Greek_eta is a representation of the C constant GDK_KEY_Greek_eta.
const KEY_Greek_eta = 2023

// KEY_Greek_etaaccent is a representation of the C constant GDK_KEY_Greek_etaaccent.
const KEY_Greek_etaaccent = 1971

// KEY_Greek_finalsmallsigma is a representation of the C constant GDK_KEY_Greek_finalsmallsigma.
const KEY_Greek_finalsmallsigma = 2035

// KEY_Greek_gamma is a representation of the C constant GDK_KEY_Greek_gamma.
const KEY_Greek_gamma = 2019

// KEY_Greek_horizbar is a representation of the C constant GDK_KEY_Greek_horizbar.
const KEY_Greek_horizbar = 1967

// KEY_Greek_iota is a representation of the C constant GDK_KEY_Greek_iota.
const KEY_Greek_iota = 2025

// KEY_Greek_iotaaccent is a representation of the C constant GDK_KEY_Greek_iotaaccent.
const KEY_Greek_iotaaccent = 1972

// KEY_Greek_iotaaccentdieresis is a representation of the C constant GDK_KEY_Greek_iotaaccentdieresis.
const KEY_Greek_iotaaccentdieresis = 1974

// KEY_Greek_iotadieresis is a representation of the C constant GDK_KEY_Greek_iotadieresis.
const KEY_Greek_iotadieresis = 1973

// KEY_Greek_kappa is a representation of the C constant GDK_KEY_Greek_kappa.
const KEY_Greek_kappa = 2026

// KEY_Greek_lambda is a representation of the C constant GDK_KEY_Greek_lambda.
const KEY_Greek_lambda = 2027

// KEY_Greek_lamda is a representation of the C constant GDK_KEY_Greek_lamda.
const KEY_Greek_lamda = 2027

// KEY_Greek_mu is a representation of the C constant GDK_KEY_Greek_mu.
const KEY_Greek_mu = 2028

// KEY_Greek_nu is a representation of the C constant GDK_KEY_Greek_nu.
const KEY_Greek_nu = 2029

// KEY_Greek_omega is a representation of the C constant GDK_KEY_Greek_omega.
const KEY_Greek_omega = 2041

// KEY_Greek_omegaaccent is a representation of the C constant GDK_KEY_Greek_omegaaccent.
const KEY_Greek_omegaaccent = 1979

// KEY_Greek_omicron is a representation of the C constant GDK_KEY_Greek_omicron.
const KEY_Greek_omicron = 2031

// KEY_Greek_omicronaccent is a representation of the C constant GDK_KEY_Greek_omicronaccent.
const KEY_Greek_omicronaccent = 1975

// KEY_Greek_phi is a representation of the C constant GDK_KEY_Greek_phi.
const KEY_Greek_phi = 2038

// KEY_Greek_pi is a representation of the C constant GDK_KEY_Greek_pi.
const KEY_Greek_pi = 2032

// KEY_Greek_psi is a representation of the C constant GDK_KEY_Greek_psi.
const KEY_Greek_psi = 2040

// KEY_Greek_rho is a representation of the C constant GDK_KEY_Greek_rho.
const KEY_Greek_rho = 2033

// KEY_Greek_sigma is a representation of the C constant GDK_KEY_Greek_sigma.
const KEY_Greek_sigma = 2034

// KEY_Greek_switch is a representation of the C constant GDK_KEY_Greek_switch.
const KEY_Greek_switch = 65406

// KEY_Greek_tau is a representation of the C constant GDK_KEY_Greek_tau.
const KEY_Greek_tau = 2036

// KEY_Greek_theta is a representation of the C constant GDK_KEY_Greek_theta.
const KEY_Greek_theta = 2024

// KEY_Greek_upsilon is a representation of the C constant GDK_KEY_Greek_upsilon.
const KEY_Greek_upsilon = 2037

// KEY_Greek_upsilonaccent is a representation of the C constant GDK_KEY_Greek_upsilonaccent.
const KEY_Greek_upsilonaccent = 1976

// KEY_Greek_upsilonaccentdieresis is a representation of the C constant GDK_KEY_Greek_upsilonaccentdieresis.
const KEY_Greek_upsilonaccentdieresis = 1978

// KEY_Greek_upsilondieresis is a representation of the C constant GDK_KEY_Greek_upsilondieresis.
const KEY_Greek_upsilondieresis = 1977

// KEY_Greek_xi is a representation of the C constant GDK_KEY_Greek_xi.
const KEY_Greek_xi = 2030

// KEY_Greek_zeta is a representation of the C constant GDK_KEY_Greek_zeta.
const KEY_Greek_zeta = 2022

// KEY_Green is a representation of the C constant GDK_KEY_Green.
const KEY_Green = 269025188

// KEY_H is a representation of the C constant GDK_KEY_H.
const KEY_H = 72

// KEY_Hangul is a representation of the C constant GDK_KEY_Hangul.
const KEY_Hangul = 65329

// KEY_Hangul_A is a representation of the C constant GDK_KEY_Hangul_A.
const KEY_Hangul_A = 3775

// KEY_Hangul_AE is a representation of the C constant GDK_KEY_Hangul_AE.
const KEY_Hangul_AE = 3776

// KEY_Hangul_AraeA is a representation of the C constant GDK_KEY_Hangul_AraeA.
const KEY_Hangul_AraeA = 3830

// KEY_Hangul_AraeAE is a representation of the C constant GDK_KEY_Hangul_AraeAE.
const KEY_Hangul_AraeAE = 3831

// KEY_Hangul_Banja is a representation of the C constant GDK_KEY_Hangul_Banja.
const KEY_Hangul_Banja = 65337

// KEY_Hangul_Cieuc is a representation of the C constant GDK_KEY_Hangul_Cieuc.
const KEY_Hangul_Cieuc = 3770

// KEY_Hangul_Codeinput is a representation of the C constant GDK_KEY_Hangul_Codeinput.
const KEY_Hangul_Codeinput = 65335

// KEY_Hangul_Dikeud is a representation of the C constant GDK_KEY_Hangul_Dikeud.
const KEY_Hangul_Dikeud = 3751

// KEY_Hangul_E is a representation of the C constant GDK_KEY_Hangul_E.
const KEY_Hangul_E = 3780

// KEY_Hangul_EO is a representation of the C constant GDK_KEY_Hangul_EO.
const KEY_Hangul_EO = 3779

// KEY_Hangul_EU is a representation of the C constant GDK_KEY_Hangul_EU.
const KEY_Hangul_EU = 3793

// KEY_Hangul_End is a representation of the C constant GDK_KEY_Hangul_End.
const KEY_Hangul_End = 65331

// KEY_Hangul_Hanja is a representation of the C constant GDK_KEY_Hangul_Hanja.
const KEY_Hangul_Hanja = 65332

// KEY_Hangul_Hieuh is a representation of the C constant GDK_KEY_Hangul_Hieuh.
const KEY_Hangul_Hieuh = 3774

// KEY_Hangul_I is a representation of the C constant GDK_KEY_Hangul_I.
const KEY_Hangul_I = 3795

// KEY_Hangul_Ieung is a representation of the C constant GDK_KEY_Hangul_Ieung.
const KEY_Hangul_Ieung = 3767

// KEY_Hangul_J_Cieuc is a representation of the C constant GDK_KEY_Hangul_J_Cieuc.
const KEY_Hangul_J_Cieuc = 3818

// KEY_Hangul_J_Dikeud is a representation of the C constant GDK_KEY_Hangul_J_Dikeud.
const KEY_Hangul_J_Dikeud = 3802

// KEY_Hangul_J_Hieuh is a representation of the C constant GDK_KEY_Hangul_J_Hieuh.
const KEY_Hangul_J_Hieuh = 3822

// KEY_Hangul_J_Ieung is a representation of the C constant GDK_KEY_Hangul_J_Ieung.
const KEY_Hangul_J_Ieung = 3816

// KEY_Hangul_J_Jieuj is a representation of the C constant GDK_KEY_Hangul_J_Jieuj.
const KEY_Hangul_J_Jieuj = 3817

// KEY_Hangul_J_Khieuq is a representation of the C constant GDK_KEY_Hangul_J_Khieuq.
const KEY_Hangul_J_Khieuq = 3819

// KEY_Hangul_J_Kiyeog is a representation of the C constant GDK_KEY_Hangul_J_Kiyeog.
const KEY_Hangul_J_Kiyeog = 3796

// KEY_Hangul_J_KiyeogSios is a representation of the C constant GDK_KEY_Hangul_J_KiyeogSios.
const KEY_Hangul_J_KiyeogSios = 3798

// KEY_Hangul_J_KkogjiDalrinIeung is a representation of the C constant GDK_KEY_Hangul_J_KkogjiDalrinIeung.
const KEY_Hangul_J_KkogjiDalrinIeung = 3833

// KEY_Hangul_J_Mieum is a representation of the C constant GDK_KEY_Hangul_J_Mieum.
const KEY_Hangul_J_Mieum = 3811

// KEY_Hangul_J_Nieun is a representation of the C constant GDK_KEY_Hangul_J_Nieun.
const KEY_Hangul_J_Nieun = 3799

// KEY_Hangul_J_NieunHieuh is a representation of the C constant GDK_KEY_Hangul_J_NieunHieuh.
const KEY_Hangul_J_NieunHieuh = 3801

// KEY_Hangul_J_NieunJieuj is a representation of the C constant GDK_KEY_Hangul_J_NieunJieuj.
const KEY_Hangul_J_NieunJieuj = 3800

// KEY_Hangul_J_PanSios is a representation of the C constant GDK_KEY_Hangul_J_PanSios.
const KEY_Hangul_J_PanSios = 3832

// KEY_Hangul_J_Phieuf is a representation of the C constant GDK_KEY_Hangul_J_Phieuf.
const KEY_Hangul_J_Phieuf = 3821

// KEY_Hangul_J_Pieub is a representation of the C constant GDK_KEY_Hangul_J_Pieub.
const KEY_Hangul_J_Pieub = 3812

// KEY_Hangul_J_PieubSios is a representation of the C constant GDK_KEY_Hangul_J_PieubSios.
const KEY_Hangul_J_PieubSios = 3813

// KEY_Hangul_J_Rieul is a representation of the C constant GDK_KEY_Hangul_J_Rieul.
const KEY_Hangul_J_Rieul = 3803

// KEY_Hangul_J_RieulHieuh is a representation of the C constant GDK_KEY_Hangul_J_RieulHieuh.
const KEY_Hangul_J_RieulHieuh = 3810

// KEY_Hangul_J_RieulKiyeog is a representation of the C constant GDK_KEY_Hangul_J_RieulKiyeog.
const KEY_Hangul_J_RieulKiyeog = 3804

// KEY_Hangul_J_RieulMieum is a representation of the C constant GDK_KEY_Hangul_J_RieulMieum.
const KEY_Hangul_J_RieulMieum = 3805

// KEY_Hangul_J_RieulPhieuf is a representation of the C constant GDK_KEY_Hangul_J_RieulPhieuf.
const KEY_Hangul_J_RieulPhieuf = 3809

// KEY_Hangul_J_RieulPieub is a representation of the C constant GDK_KEY_Hangul_J_RieulPieub.
const KEY_Hangul_J_RieulPieub = 3806

// KEY_Hangul_J_RieulSios is a representation of the C constant GDK_KEY_Hangul_J_RieulSios.
const KEY_Hangul_J_RieulSios = 3807

// KEY_Hangul_J_RieulTieut is a representation of the C constant GDK_KEY_Hangul_J_RieulTieut.
const KEY_Hangul_J_RieulTieut = 3808

// KEY_Hangul_J_Sios is a representation of the C constant GDK_KEY_Hangul_J_Sios.
const KEY_Hangul_J_Sios = 3814

// KEY_Hangul_J_SsangKiyeog is a representation of the C constant GDK_KEY_Hangul_J_SsangKiyeog.
const KEY_Hangul_J_SsangKiyeog = 3797

// KEY_Hangul_J_SsangSios is a representation of the C constant GDK_KEY_Hangul_J_SsangSios.
const KEY_Hangul_J_SsangSios = 3815

// KEY_Hangul_J_Tieut is a representation of the C constant GDK_KEY_Hangul_J_Tieut.
const KEY_Hangul_J_Tieut = 3820

// KEY_Hangul_J_YeorinHieuh is a representation of the C constant GDK_KEY_Hangul_J_YeorinHieuh.
const KEY_Hangul_J_YeorinHieuh = 3834

// KEY_Hangul_Jamo is a representation of the C constant GDK_KEY_Hangul_Jamo.
const KEY_Hangul_Jamo = 65333

// KEY_Hangul_Jeonja is a representation of the C constant GDK_KEY_Hangul_Jeonja.
const KEY_Hangul_Jeonja = 65336

// KEY_Hangul_Jieuj is a representation of the C constant GDK_KEY_Hangul_Jieuj.
const KEY_Hangul_Jieuj = 3768

// KEY_Hangul_Khieuq is a representation of the C constant GDK_KEY_Hangul_Khieuq.
const KEY_Hangul_Khieuq = 3771

// KEY_Hangul_Kiyeog is a representation of the C constant GDK_KEY_Hangul_Kiyeog.
const KEY_Hangul_Kiyeog = 3745

// KEY_Hangul_KiyeogSios is a representation of the C constant GDK_KEY_Hangul_KiyeogSios.
const KEY_Hangul_KiyeogSios = 3747

// KEY_Hangul_KkogjiDalrinIeung is a representation of the C constant GDK_KEY_Hangul_KkogjiDalrinIeung.
const KEY_Hangul_KkogjiDalrinIeung = 3827

// KEY_Hangul_Mieum is a representation of the C constant GDK_KEY_Hangul_Mieum.
const KEY_Hangul_Mieum = 3761

// KEY_Hangul_MultipleCandidate is a representation of the C constant GDK_KEY_Hangul_MultipleCandidate.
const KEY_Hangul_MultipleCandidate = 65341

// KEY_Hangul_Nieun is a representation of the C constant GDK_KEY_Hangul_Nieun.
const KEY_Hangul_Nieun = 3748

// KEY_Hangul_NieunHieuh is a representation of the C constant GDK_KEY_Hangul_NieunHieuh.
const KEY_Hangul_NieunHieuh = 3750

// KEY_Hangul_NieunJieuj is a representation of the C constant GDK_KEY_Hangul_NieunJieuj.
const KEY_Hangul_NieunJieuj = 3749

// KEY_Hangul_O is a representation of the C constant GDK_KEY_Hangul_O.
const KEY_Hangul_O = 3783

// KEY_Hangul_OE is a representation of the C constant GDK_KEY_Hangul_OE.
const KEY_Hangul_OE = 3786

// KEY_Hangul_PanSios is a representation of the C constant GDK_KEY_Hangul_PanSios.
const KEY_Hangul_PanSios = 3826

// KEY_Hangul_Phieuf is a representation of the C constant GDK_KEY_Hangul_Phieuf.
const KEY_Hangul_Phieuf = 3773

// KEY_Hangul_Pieub is a representation of the C constant GDK_KEY_Hangul_Pieub.
const KEY_Hangul_Pieub = 3762

// KEY_Hangul_PieubSios is a representation of the C constant GDK_KEY_Hangul_PieubSios.
const KEY_Hangul_PieubSios = 3764

// KEY_Hangul_PostHanja is a representation of the C constant GDK_KEY_Hangul_PostHanja.
const KEY_Hangul_PostHanja = 65339

// KEY_Hangul_PreHanja is a representation of the C constant GDK_KEY_Hangul_PreHanja.
const KEY_Hangul_PreHanja = 65338

// KEY_Hangul_PreviousCandidate is a representation of the C constant GDK_KEY_Hangul_PreviousCandidate.
const KEY_Hangul_PreviousCandidate = 65342

// KEY_Hangul_Rieul is a representation of the C constant GDK_KEY_Hangul_Rieul.
const KEY_Hangul_Rieul = 3753

// KEY_Hangul_RieulHieuh is a representation of the C constant GDK_KEY_Hangul_RieulHieuh.
const KEY_Hangul_RieulHieuh = 3760

// KEY_Hangul_RieulKiyeog is a representation of the C constant GDK_KEY_Hangul_RieulKiyeog.
const KEY_Hangul_RieulKiyeog = 3754

// KEY_Hangul_RieulMieum is a representation of the C constant GDK_KEY_Hangul_RieulMieum.
const KEY_Hangul_RieulMieum = 3755

// KEY_Hangul_RieulPhieuf is a representation of the C constant GDK_KEY_Hangul_RieulPhieuf.
const KEY_Hangul_RieulPhieuf = 3759

// KEY_Hangul_RieulPieub is a representation of the C constant GDK_KEY_Hangul_RieulPieub.
const KEY_Hangul_RieulPieub = 3756

// KEY_Hangul_RieulSios is a representation of the C constant GDK_KEY_Hangul_RieulSios.
const KEY_Hangul_RieulSios = 3757

// KEY_Hangul_RieulTieut is a representation of the C constant GDK_KEY_Hangul_RieulTieut.
const KEY_Hangul_RieulTieut = 3758

// KEY_Hangul_RieulYeorinHieuh is a representation of the C constant GDK_KEY_Hangul_RieulYeorinHieuh.
const KEY_Hangul_RieulYeorinHieuh = 3823

// KEY_Hangul_Romaja is a representation of the C constant GDK_KEY_Hangul_Romaja.
const KEY_Hangul_Romaja = 65334

// KEY_Hangul_SingleCandidate is a representation of the C constant GDK_KEY_Hangul_SingleCandidate.
const KEY_Hangul_SingleCandidate = 65340

// KEY_Hangul_Sios is a representation of the C constant GDK_KEY_Hangul_Sios.
const KEY_Hangul_Sios = 3765

// KEY_Hangul_Special is a representation of the C constant GDK_KEY_Hangul_Special.
const KEY_Hangul_Special = 65343

// KEY_Hangul_SsangDikeud is a representation of the C constant GDK_KEY_Hangul_SsangDikeud.
const KEY_Hangul_SsangDikeud = 3752

// KEY_Hangul_SsangJieuj is a representation of the C constant GDK_KEY_Hangul_SsangJieuj.
const KEY_Hangul_SsangJieuj = 3769

// KEY_Hangul_SsangKiyeog is a representation of the C constant GDK_KEY_Hangul_SsangKiyeog.
const KEY_Hangul_SsangKiyeog = 3746

// KEY_Hangul_SsangPieub is a representation of the C constant GDK_KEY_Hangul_SsangPieub.
const KEY_Hangul_SsangPieub = 3763

// KEY_Hangul_SsangSios is a representation of the C constant GDK_KEY_Hangul_SsangSios.
const KEY_Hangul_SsangSios = 3766

// KEY_Hangul_Start is a representation of the C constant GDK_KEY_Hangul_Start.
const KEY_Hangul_Start = 65330

// KEY_Hangul_SunkyeongeumMieum is a representation of the C constant GDK_KEY_Hangul_SunkyeongeumMieum.
const KEY_Hangul_SunkyeongeumMieum = 3824

// KEY_Hangul_SunkyeongeumPhieuf is a representation of the C constant GDK_KEY_Hangul_SunkyeongeumPhieuf.
const KEY_Hangul_SunkyeongeumPhieuf = 3828

// KEY_Hangul_SunkyeongeumPieub is a representation of the C constant GDK_KEY_Hangul_SunkyeongeumPieub.
const KEY_Hangul_SunkyeongeumPieub = 3825

// KEY_Hangul_Tieut is a representation of the C constant GDK_KEY_Hangul_Tieut.
const KEY_Hangul_Tieut = 3772

// KEY_Hangul_U is a representation of the C constant GDK_KEY_Hangul_U.
const KEY_Hangul_U = 3788

// KEY_Hangul_WA is a representation of the C constant GDK_KEY_Hangul_WA.
const KEY_Hangul_WA = 3784

// KEY_Hangul_WAE is a representation of the C constant GDK_KEY_Hangul_WAE.
const KEY_Hangul_WAE = 3785

// KEY_Hangul_WE is a representation of the C constant GDK_KEY_Hangul_WE.
const KEY_Hangul_WE = 3790

// KEY_Hangul_WEO is a representation of the C constant GDK_KEY_Hangul_WEO.
const KEY_Hangul_WEO = 3789

// KEY_Hangul_WI is a representation of the C constant GDK_KEY_Hangul_WI.
const KEY_Hangul_WI = 3791

// KEY_Hangul_YA is a representation of the C constant GDK_KEY_Hangul_YA.
const KEY_Hangul_YA = 3777

// KEY_Hangul_YAE is a representation of the C constant GDK_KEY_Hangul_YAE.
const KEY_Hangul_YAE = 3778

// KEY_Hangul_YE is a representation of the C constant GDK_KEY_Hangul_YE.
const KEY_Hangul_YE = 3782

// KEY_Hangul_YEO is a representation of the C constant GDK_KEY_Hangul_YEO.
const KEY_Hangul_YEO = 3781

// KEY_Hangul_YI is a representation of the C constant GDK_KEY_Hangul_YI.
const KEY_Hangul_YI = 3794

// KEY_Hangul_YO is a representation of the C constant GDK_KEY_Hangul_YO.
const KEY_Hangul_YO = 3787

// KEY_Hangul_YU is a representation of the C constant GDK_KEY_Hangul_YU.
const KEY_Hangul_YU = 3792

// KEY_Hangul_YeorinHieuh is a representation of the C constant GDK_KEY_Hangul_YeorinHieuh.
const KEY_Hangul_YeorinHieuh = 3829

// KEY_Hangul_switch is a representation of the C constant GDK_KEY_Hangul_switch.
const KEY_Hangul_switch = 65406

// KEY_Hankaku is a representation of the C constant GDK_KEY_Hankaku.
const KEY_Hankaku = 65321

// KEY_Hcircumflex is a representation of the C constant GDK_KEY_Hcircumflex.
const KEY_Hcircumflex = 678

// KEY_Hebrew_switch is a representation of the C constant GDK_KEY_Hebrew_switch.
const KEY_Hebrew_switch = 65406

// KEY_Help is a representation of the C constant GDK_KEY_Help.
const KEY_Help = 65386

// KEY_Henkan is a representation of the C constant GDK_KEY_Henkan.
const KEY_Henkan = 65315

// KEY_Henkan_Mode is a representation of the C constant GDK_KEY_Henkan_Mode.
const KEY_Henkan_Mode = 65315

// KEY_Hibernate is a representation of the C constant GDK_KEY_Hibernate.
const KEY_Hibernate = 269025192

// KEY_Hiragana is a representation of the C constant GDK_KEY_Hiragana.
const KEY_Hiragana = 65317

// KEY_Hiragana_Katakana is a representation of the C constant GDK_KEY_Hiragana_Katakana.
const KEY_Hiragana_Katakana = 65319

// KEY_History is a representation of the C constant GDK_KEY_History.
const KEY_History = 269025079

// KEY_Home is a representation of the C constant GDK_KEY_Home.
const KEY_Home = 65360

// KEY_HomePage is a representation of the C constant GDK_KEY_HomePage.
const KEY_HomePage = 269025048

// KEY_HotLinks is a representation of the C constant GDK_KEY_HotLinks.
const KEY_HotLinks = 269025082

// KEY_Hstroke is a representation of the C constant GDK_KEY_Hstroke.
const KEY_Hstroke = 673

// KEY_Hyper_L is a representation of the C constant GDK_KEY_Hyper_L.
const KEY_Hyper_L = 65517

// KEY_Hyper_R is a representation of the C constant GDK_KEY_Hyper_R.
const KEY_Hyper_R = 65518

// KEY_I is a representation of the C constant GDK_KEY_I.
const KEY_I = 73

// KEY_ISO_Center_Object is a representation of the C constant GDK_KEY_ISO_Center_Object.
const KEY_ISO_Center_Object = 65075

// KEY_ISO_Continuous_Underline is a representation of the C constant GDK_KEY_ISO_Continuous_Underline.
const KEY_ISO_Continuous_Underline = 65072

// KEY_ISO_Discontinuous_Underline is a representation of the C constant GDK_KEY_ISO_Discontinuous_Underline.
const KEY_ISO_Discontinuous_Underline = 65073

// KEY_ISO_Emphasize is a representation of the C constant GDK_KEY_ISO_Emphasize.
const KEY_ISO_Emphasize = 65074

// KEY_ISO_Enter is a representation of the C constant GDK_KEY_ISO_Enter.
const KEY_ISO_Enter = 65076

// KEY_ISO_Fast_Cursor_Down is a representation of the C constant GDK_KEY_ISO_Fast_Cursor_Down.
const KEY_ISO_Fast_Cursor_Down = 65071

// KEY_ISO_Fast_Cursor_Left is a representation of the C constant GDK_KEY_ISO_Fast_Cursor_Left.
const KEY_ISO_Fast_Cursor_Left = 65068

// KEY_ISO_Fast_Cursor_Right is a representation of the C constant GDK_KEY_ISO_Fast_Cursor_Right.
const KEY_ISO_Fast_Cursor_Right = 65069

// KEY_ISO_Fast_Cursor_Up is a representation of the C constant GDK_KEY_ISO_Fast_Cursor_Up.
const KEY_ISO_Fast_Cursor_Up = 65070

// KEY_ISO_First_Group is a representation of the C constant GDK_KEY_ISO_First_Group.
const KEY_ISO_First_Group = 65036

// KEY_ISO_First_Group_Lock is a representation of the C constant GDK_KEY_ISO_First_Group_Lock.
const KEY_ISO_First_Group_Lock = 65037

// KEY_ISO_Group_Latch is a representation of the C constant GDK_KEY_ISO_Group_Latch.
const KEY_ISO_Group_Latch = 65030

// KEY_ISO_Group_Lock is a representation of the C constant GDK_KEY_ISO_Group_Lock.
const KEY_ISO_Group_Lock = 65031

// KEY_ISO_Group_Shift is a representation of the C constant GDK_KEY_ISO_Group_Shift.
const KEY_ISO_Group_Shift = 65406

// KEY_ISO_Last_Group is a representation of the C constant GDK_KEY_ISO_Last_Group.
const KEY_ISO_Last_Group = 65038

// KEY_ISO_Last_Group_Lock is a representation of the C constant GDK_KEY_ISO_Last_Group_Lock.
const KEY_ISO_Last_Group_Lock = 65039

// KEY_ISO_Left_Tab is a representation of the C constant GDK_KEY_ISO_Left_Tab.
const KEY_ISO_Left_Tab = 65056

// KEY_ISO_Level2_Latch is a representation of the C constant GDK_KEY_ISO_Level2_Latch.
const KEY_ISO_Level2_Latch = 65026

// KEY_ISO_Level3_Latch is a representation of the C constant GDK_KEY_ISO_Level3_Latch.
const KEY_ISO_Level3_Latch = 65028

// KEY_ISO_Level3_Lock is a representation of the C constant GDK_KEY_ISO_Level3_Lock.
const KEY_ISO_Level3_Lock = 65029

// KEY_ISO_Level3_Shift is a representation of the C constant GDK_KEY_ISO_Level3_Shift.
const KEY_ISO_Level3_Shift = 65027

// KEY_ISO_Level5_Latch is a representation of the C constant GDK_KEY_ISO_Level5_Latch.
const KEY_ISO_Level5_Latch = 65042

// KEY_ISO_Level5_Lock is a representation of the C constant GDK_KEY_ISO_Level5_Lock.
const KEY_ISO_Level5_Lock = 65043

// KEY_ISO_Level5_Shift is a representation of the C constant GDK_KEY_ISO_Level5_Shift.
const KEY_ISO_Level5_Shift = 65041

// KEY_ISO_Lock is a representation of the C constant GDK_KEY_ISO_Lock.
const KEY_ISO_Lock = 65025

// KEY_ISO_Move_Line_Down is a representation of the C constant GDK_KEY_ISO_Move_Line_Down.
const KEY_ISO_Move_Line_Down = 65058

// KEY_ISO_Move_Line_Up is a representation of the C constant GDK_KEY_ISO_Move_Line_Up.
const KEY_ISO_Move_Line_Up = 65057

// KEY_ISO_Next_Group is a representation of the C constant GDK_KEY_ISO_Next_Group.
const KEY_ISO_Next_Group = 65032

// KEY_ISO_Next_Group_Lock is a representation of the C constant GDK_KEY_ISO_Next_Group_Lock.
const KEY_ISO_Next_Group_Lock = 65033

// KEY_ISO_Partial_Line_Down is a representation of the C constant GDK_KEY_ISO_Partial_Line_Down.
const KEY_ISO_Partial_Line_Down = 65060

// KEY_ISO_Partial_Line_Up is a representation of the C constant GDK_KEY_ISO_Partial_Line_Up.
const KEY_ISO_Partial_Line_Up = 65059

// KEY_ISO_Partial_Space_Left is a representation of the C constant GDK_KEY_ISO_Partial_Space_Left.
const KEY_ISO_Partial_Space_Left = 65061

// KEY_ISO_Partial_Space_Right is a representation of the C constant GDK_KEY_ISO_Partial_Space_Right.
const KEY_ISO_Partial_Space_Right = 65062

// KEY_ISO_Prev_Group is a representation of the C constant GDK_KEY_ISO_Prev_Group.
const KEY_ISO_Prev_Group = 65034

// KEY_ISO_Prev_Group_Lock is a representation of the C constant GDK_KEY_ISO_Prev_Group_Lock.
const KEY_ISO_Prev_Group_Lock = 65035

// KEY_ISO_Release_Both_Margins is a representation of the C constant GDK_KEY_ISO_Release_Both_Margins.
const KEY_ISO_Release_Both_Margins = 65067

// KEY_ISO_Release_Margin_Left is a representation of the C constant GDK_KEY_ISO_Release_Margin_Left.
const KEY_ISO_Release_Margin_Left = 65065

// KEY_ISO_Release_Margin_Right is a representation of the C constant GDK_KEY_ISO_Release_Margin_Right.
const KEY_ISO_Release_Margin_Right = 65066

// KEY_ISO_Set_Margin_Left is a representation of the C constant GDK_KEY_ISO_Set_Margin_Left.
const KEY_ISO_Set_Margin_Left = 65063

// KEY_ISO_Set_Margin_Right is a representation of the C constant GDK_KEY_ISO_Set_Margin_Right.
const KEY_ISO_Set_Margin_Right = 65064

// KEY_Iabovedot is a representation of the C constant GDK_KEY_Iabovedot.
const KEY_Iabovedot = 681

// KEY_Iacute is a representation of the C constant GDK_KEY_Iacute.
const KEY_Iacute = 205

// KEY_Ibelowdot is a representation of the C constant GDK_KEY_Ibelowdot.
const KEY_Ibelowdot = 16785098

// KEY_Ibreve is a representation of the C constant GDK_KEY_Ibreve.
const KEY_Ibreve = 16777516

// KEY_Icircumflex is a representation of the C constant GDK_KEY_Icircumflex.
const KEY_Icircumflex = 206

// KEY_Idiaeresis is a representation of the C constant GDK_KEY_Idiaeresis.
const KEY_Idiaeresis = 207

// KEY_Igrave is a representation of the C constant GDK_KEY_Igrave.
const KEY_Igrave = 204

// KEY_Ihook is a representation of the C constant GDK_KEY_Ihook.
const KEY_Ihook = 16785096

// KEY_Imacron is a representation of the C constant GDK_KEY_Imacron.
const KEY_Imacron = 975

// KEY_Insert is a representation of the C constant GDK_KEY_Insert.
const KEY_Insert = 65379

// KEY_Iogonek is a representation of the C constant GDK_KEY_Iogonek.
const KEY_Iogonek = 967

// KEY_Itilde is a representation of the C constant GDK_KEY_Itilde.
const KEY_Itilde = 933

// KEY_J is a representation of the C constant GDK_KEY_J.
const KEY_J = 74

// KEY_Jcircumflex is a representation of the C constant GDK_KEY_Jcircumflex.
const KEY_Jcircumflex = 684

// KEY_K is a representation of the C constant GDK_KEY_K.
const KEY_K = 75

// KEY_KP_0 is a representation of the C constant GDK_KEY_KP_0.
const KEY_KP_0 = 65456

// KEY_KP_1 is a representation of the C constant GDK_KEY_KP_1.
const KEY_KP_1 = 65457

// KEY_KP_2 is a representation of the C constant GDK_KEY_KP_2.
const KEY_KP_2 = 65458

// KEY_KP_3 is a representation of the C constant GDK_KEY_KP_3.
const KEY_KP_3 = 65459

// KEY_KP_4 is a representation of the C constant GDK_KEY_KP_4.
const KEY_KP_4 = 65460

// KEY_KP_5 is a representation of the C constant GDK_KEY_KP_5.
const KEY_KP_5 = 65461

// KEY_KP_6 is a representation of the C constant GDK_KEY_KP_6.
const KEY_KP_6 = 65462

// KEY_KP_7 is a representation of the C constant GDK_KEY_KP_7.
const KEY_KP_7 = 65463

// KEY_KP_8 is a representation of the C constant GDK_KEY_KP_8.
const KEY_KP_8 = 65464

// KEY_KP_9 is a representation of the C constant GDK_KEY_KP_9.
const KEY_KP_9 = 65465

// KEY_KP_Add is a representation of the C constant GDK_KEY_KP_Add.
const KEY_KP_Add = 65451

// KEY_KP_Begin is a representation of the C constant GDK_KEY_KP_Begin.
const KEY_KP_Begin = 65437

// KEY_KP_Decimal is a representation of the C constant GDK_KEY_KP_Decimal.
const KEY_KP_Decimal = 65454

// KEY_KP_Delete is a representation of the C constant GDK_KEY_KP_Delete.
const KEY_KP_Delete = 65439

// KEY_KP_Divide is a representation of the C constant GDK_KEY_KP_Divide.
const KEY_KP_Divide = 65455

// KEY_KP_Down is a representation of the C constant GDK_KEY_KP_Down.
const KEY_KP_Down = 65433

// KEY_KP_End is a representation of the C constant GDK_KEY_KP_End.
const KEY_KP_End = 65436

// KEY_KP_Enter is a representation of the C constant GDK_KEY_KP_Enter.
const KEY_KP_Enter = 65421

// KEY_KP_Equal is a representation of the C constant GDK_KEY_KP_Equal.
const KEY_KP_Equal = 65469

// KEY_KP_F1 is a representation of the C constant GDK_KEY_KP_F1.
const KEY_KP_F1 = 65425

// KEY_KP_F2 is a representation of the C constant GDK_KEY_KP_F2.
const KEY_KP_F2 = 65426

// KEY_KP_F3 is a representation of the C constant GDK_KEY_KP_F3.
const KEY_KP_F3 = 65427

// KEY_KP_F4 is a representation of the C constant GDK_KEY_KP_F4.
const KEY_KP_F4 = 65428

// KEY_KP_Home is a representation of the C constant GDK_KEY_KP_Home.
const KEY_KP_Home = 65429

// KEY_KP_Insert is a representation of the C constant GDK_KEY_KP_Insert.
const KEY_KP_Insert = 65438

// KEY_KP_Left is a representation of the C constant GDK_KEY_KP_Left.
const KEY_KP_Left = 65430

// KEY_KP_Multiply is a representation of the C constant GDK_KEY_KP_Multiply.
const KEY_KP_Multiply = 65450

// KEY_KP_Next is a representation of the C constant GDK_KEY_KP_Next.
const KEY_KP_Next = 65435

// KEY_KP_Page_Down is a representation of the C constant GDK_KEY_KP_Page_Down.
const KEY_KP_Page_Down = 65435

// KEY_KP_Page_Up is a representation of the C constant GDK_KEY_KP_Page_Up.
const KEY_KP_Page_Up = 65434

// KEY_KP_Prior is a representation of the C constant GDK_KEY_KP_Prior.
const KEY_KP_Prior = 65434

// KEY_KP_Right is a representation of the C constant GDK_KEY_KP_Right.
const KEY_KP_Right = 65432

// KEY_KP_Separator is a representation of the C constant GDK_KEY_KP_Separator.
const KEY_KP_Separator = 65452

// KEY_KP_Space is a representation of the C constant GDK_KEY_KP_Space.
const KEY_KP_Space = 65408

// KEY_KP_Subtract is a representation of the C constant GDK_KEY_KP_Subtract.
const KEY_KP_Subtract = 65453

// KEY_KP_Tab is a representation of the C constant GDK_KEY_KP_Tab.
const KEY_KP_Tab = 65417

// KEY_KP_Up is a representation of the C constant GDK_KEY_KP_Up.
const KEY_KP_Up = 65431

// KEY_Kana_Lock is a representation of the C constant GDK_KEY_Kana_Lock.
const KEY_Kana_Lock = 65325

// KEY_Kana_Shift is a representation of the C constant GDK_KEY_Kana_Shift.
const KEY_Kana_Shift = 65326

// KEY_Kanji is a representation of the C constant GDK_KEY_Kanji.
const KEY_Kanji = 65313

// KEY_Kanji_Bangou is a representation of the C constant GDK_KEY_Kanji_Bangou.
const KEY_Kanji_Bangou = 65335

// KEY_Katakana is a representation of the C constant GDK_KEY_Katakana.
const KEY_Katakana = 65318

// KEY_KbdBrightnessDown is a representation of the C constant GDK_KEY_KbdBrightnessDown.
const KEY_KbdBrightnessDown = 269025030

// KEY_KbdBrightnessUp is a representation of the C constant GDK_KEY_KbdBrightnessUp.
const KEY_KbdBrightnessUp = 269025029

// KEY_KbdLightOnOff is a representation of the C constant GDK_KEY_KbdLightOnOff.
const KEY_KbdLightOnOff = 269025028

// KEY_Kcedilla is a representation of the C constant GDK_KEY_Kcedilla.
const KEY_Kcedilla = 979

// KEY_Korean_Won is a representation of the C constant GDK_KEY_Korean_Won.
const KEY_Korean_Won = 3839

// KEY_L is a representation of the C constant GDK_KEY_L.
const KEY_L = 76

// KEY_L1 is a representation of the C constant GDK_KEY_L1.
const KEY_L1 = 65480

// KEY_L10 is a representation of the C constant GDK_KEY_L10.
const KEY_L10 = 65489

// KEY_L2 is a representation of the C constant GDK_KEY_L2.
const KEY_L2 = 65481

// KEY_L3 is a representation of the C constant GDK_KEY_L3.
const KEY_L3 = 65482

// KEY_L4 is a representation of the C constant GDK_KEY_L4.
const KEY_L4 = 65483

// KEY_L5 is a representation of the C constant GDK_KEY_L5.
const KEY_L5 = 65484

// KEY_L6 is a representation of the C constant GDK_KEY_L6.
const KEY_L6 = 65485

// KEY_L7 is a representation of the C constant GDK_KEY_L7.
const KEY_L7 = 65486

// KEY_L8 is a representation of the C constant GDK_KEY_L8.
const KEY_L8 = 65487

// KEY_L9 is a representation of the C constant GDK_KEY_L9.
const KEY_L9 = 65488

// KEY_Lacute is a representation of the C constant GDK_KEY_Lacute.
const KEY_Lacute = 453

// KEY_Last_Virtual_Screen is a representation of the C constant GDK_KEY_Last_Virtual_Screen.
const KEY_Last_Virtual_Screen = 65236

// KEY_Launch0 is a representation of the C constant GDK_KEY_Launch0.
const KEY_Launch0 = 269025088

// KEY_Launch1 is a representation of the C constant GDK_KEY_Launch1.
const KEY_Launch1 = 269025089

// KEY_Launch2 is a representation of the C constant GDK_KEY_Launch2.
const KEY_Launch2 = 269025090

// KEY_Launch3 is a representation of the C constant GDK_KEY_Launch3.
const KEY_Launch3 = 269025091

// KEY_Launch4 is a representation of the C constant GDK_KEY_Launch4.
const KEY_Launch4 = 269025092

// KEY_Launch5 is a representation of the C constant GDK_KEY_Launch5.
const KEY_Launch5 = 269025093

// KEY_Launch6 is a representation of the C constant GDK_KEY_Launch6.
const KEY_Launch6 = 269025094

// KEY_Launch7 is a representation of the C constant GDK_KEY_Launch7.
const KEY_Launch7 = 269025095

// KEY_Launch8 is a representation of the C constant GDK_KEY_Launch8.
const KEY_Launch8 = 269025096

// KEY_Launch9 is a representation of the C constant GDK_KEY_Launch9.
const KEY_Launch9 = 269025097

// KEY_LaunchA is a representation of the C constant GDK_KEY_LaunchA.
const KEY_LaunchA = 269025098

// KEY_LaunchB is a representation of the C constant GDK_KEY_LaunchB.
const KEY_LaunchB = 269025099

// KEY_LaunchC is a representation of the C constant GDK_KEY_LaunchC.
const KEY_LaunchC = 269025100

// KEY_LaunchD is a representation of the C constant GDK_KEY_LaunchD.
const KEY_LaunchD = 269025101

// KEY_LaunchE is a representation of the C constant GDK_KEY_LaunchE.
const KEY_LaunchE = 269025102

// KEY_LaunchF is a representation of the C constant GDK_KEY_LaunchF.
const KEY_LaunchF = 269025103

// KEY_Lbelowdot is a representation of the C constant GDK_KEY_Lbelowdot.
const KEY_Lbelowdot = 16784950

// KEY_Lcaron is a representation of the C constant GDK_KEY_Lcaron.
const KEY_Lcaron = 421

// KEY_Lcedilla is a representation of the C constant GDK_KEY_Lcedilla.
const KEY_Lcedilla = 934

// KEY_Left is a representation of the C constant GDK_KEY_Left.
const KEY_Left = 65361

// KEY_LightBulb is a representation of the C constant GDK_KEY_LightBulb.
const KEY_LightBulb = 269025077

// KEY_Linefeed is a representation of the C constant GDK_KEY_Linefeed.
const KEY_Linefeed = 65290

// KEY_LiraSign is a representation of the C constant GDK_KEY_LiraSign.
const KEY_LiraSign = 16785572

// KEY_LogGrabInfo is a representation of the C constant GDK_KEY_LogGrabInfo.
const KEY_LogGrabInfo = 269024805

// KEY_LogOff is a representation of the C constant GDK_KEY_LogOff.
const KEY_LogOff = 269025121

// KEY_LogWindowTree is a representation of the C constant GDK_KEY_LogWindowTree.
const KEY_LogWindowTree = 269024804

// KEY_Lstroke is a representation of the C constant GDK_KEY_Lstroke.
const KEY_Lstroke = 419

// KEY_M is a representation of the C constant GDK_KEY_M.
const KEY_M = 77

// KEY_Mabovedot is a representation of the C constant GDK_KEY_Mabovedot.
const KEY_Mabovedot = 16784960

// KEY_Macedonia_DSE is a representation of the C constant GDK_KEY_Macedonia_DSE.
const KEY_Macedonia_DSE = 1717

// KEY_Macedonia_GJE is a representation of the C constant GDK_KEY_Macedonia_GJE.
const KEY_Macedonia_GJE = 1714

// KEY_Macedonia_KJE is a representation of the C constant GDK_KEY_Macedonia_KJE.
const KEY_Macedonia_KJE = 1724

// KEY_Macedonia_dse is a representation of the C constant GDK_KEY_Macedonia_dse.
const KEY_Macedonia_dse = 1701

// KEY_Macedonia_gje is a representation of the C constant GDK_KEY_Macedonia_gje.
const KEY_Macedonia_gje = 1698

// KEY_Macedonia_kje is a representation of the C constant GDK_KEY_Macedonia_kje.
const KEY_Macedonia_kje = 1708

// KEY_Mae_Koho is a representation of the C constant GDK_KEY_Mae_Koho.
const KEY_Mae_Koho = 65342

// KEY_Mail is a representation of the C constant GDK_KEY_Mail.
const KEY_Mail = 269025049

// KEY_MailForward is a representation of the C constant GDK_KEY_MailForward.
const KEY_MailForward = 269025168

// KEY_Market is a representation of the C constant GDK_KEY_Market.
const KEY_Market = 269025122

// KEY_Massyo is a representation of the C constant GDK_KEY_Massyo.
const KEY_Massyo = 65324

// KEY_Meeting is a representation of the C constant GDK_KEY_Meeting.
const KEY_Meeting = 269025123

// KEY_Memo is a representation of the C constant GDK_KEY_Memo.
const KEY_Memo = 269025054

// KEY_Menu is a representation of the C constant GDK_KEY_Menu.
const KEY_Menu = 65383

// KEY_MenuKB is a representation of the C constant GDK_KEY_MenuKB.
const KEY_MenuKB = 269025125

// KEY_MenuPB is a representation of the C constant GDK_KEY_MenuPB.
const KEY_MenuPB = 269025126

// KEY_Messenger is a representation of the C constant GDK_KEY_Messenger.
const KEY_Messenger = 269025166

// KEY_Meta_L is a representation of the C constant GDK_KEY_Meta_L.
const KEY_Meta_L = 65511

// KEY_Meta_R is a representation of the C constant GDK_KEY_Meta_R.
const KEY_Meta_R = 65512

// KEY_MillSign is a representation of the C constant GDK_KEY_MillSign.
const KEY_MillSign = 16785573

// KEY_ModeLock is a representation of the C constant GDK_KEY_ModeLock.
const KEY_ModeLock = 269025025

// KEY_Mode_switch is a representation of the C constant GDK_KEY_Mode_switch.
const KEY_Mode_switch = 65406

// KEY_MonBrightnessDown is a representation of the C constant GDK_KEY_MonBrightnessDown.
const KEY_MonBrightnessDown = 269025027

// KEY_MonBrightnessUp is a representation of the C constant GDK_KEY_MonBrightnessUp.
const KEY_MonBrightnessUp = 269025026

// KEY_MouseKeys_Accel_Enable is a representation of the C constant GDK_KEY_MouseKeys_Accel_Enable.
const KEY_MouseKeys_Accel_Enable = 65143

// KEY_MouseKeys_Enable is a representation of the C constant GDK_KEY_MouseKeys_Enable.
const KEY_MouseKeys_Enable = 65142

// KEY_Muhenkan is a representation of the C constant GDK_KEY_Muhenkan.
const KEY_Muhenkan = 65314

// KEY_Multi_key is a representation of the C constant GDK_KEY_Multi_key.
const KEY_Multi_key = 65312

// KEY_MultipleCandidate is a representation of the C constant GDK_KEY_MultipleCandidate.
const KEY_MultipleCandidate = 65341

// KEY_Music is a representation of the C constant GDK_KEY_Music.
const KEY_Music = 269025170

// KEY_MyComputer is a representation of the C constant GDK_KEY_MyComputer.
const KEY_MyComputer = 269025075

// KEY_MySites is a representation of the C constant GDK_KEY_MySites.
const KEY_MySites = 269025127

// KEY_N is a representation of the C constant GDK_KEY_N.
const KEY_N = 78

// KEY_Nacute is a representation of the C constant GDK_KEY_Nacute.
const KEY_Nacute = 465

// KEY_NairaSign is a representation of the C constant GDK_KEY_NairaSign.
const KEY_NairaSign = 16785574

// KEY_Ncaron is a representation of the C constant GDK_KEY_Ncaron.
const KEY_Ncaron = 466

// KEY_Ncedilla is a representation of the C constant GDK_KEY_Ncedilla.
const KEY_Ncedilla = 977

// KEY_New is a representation of the C constant GDK_KEY_New.
const KEY_New = 269025128

// KEY_NewSheqelSign is a representation of the C constant GDK_KEY_NewSheqelSign.
const KEY_NewSheqelSign = 16785578

// KEY_News is a representation of the C constant GDK_KEY_News.
const KEY_News = 269025129

// KEY_Next is a representation of the C constant GDK_KEY_Next.
const KEY_Next = 65366

// KEY_Next_VMode is a representation of the C constant GDK_KEY_Next_VMode.
const KEY_Next_VMode = 269024802

// KEY_Next_Virtual_Screen is a representation of the C constant GDK_KEY_Next_Virtual_Screen.
const KEY_Next_Virtual_Screen = 65234

// KEY_Ntilde is a representation of the C constant GDK_KEY_Ntilde.
const KEY_Ntilde = 209

// KEY_Num_Lock is a representation of the C constant GDK_KEY_Num_Lock.
const KEY_Num_Lock = 65407

// KEY_O is a representation of the C constant GDK_KEY_O.
const KEY_O = 79

// KEY_OE is a representation of the C constant GDK_KEY_OE.
const KEY_OE = 5052

// KEY_Oacute is a representation of the C constant GDK_KEY_Oacute.
const KEY_Oacute = 211

// KEY_Obarred is a representation of the C constant GDK_KEY_Obarred.
const KEY_Obarred = 16777631

// KEY_Obelowdot is a representation of the C constant GDK_KEY_Obelowdot.
const KEY_Obelowdot = 16785100

// KEY_Ocaron is a representation of the C constant GDK_KEY_Ocaron.
const KEY_Ocaron = 16777681

// KEY_Ocircumflex is a representation of the C constant GDK_KEY_Ocircumflex.
const KEY_Ocircumflex = 212

// KEY_Ocircumflexacute is a representation of the C constant GDK_KEY_Ocircumflexacute.
const KEY_Ocircumflexacute = 16785104

// KEY_Ocircumflexbelowdot is a representation of the C constant GDK_KEY_Ocircumflexbelowdot.
const KEY_Ocircumflexbelowdot = 16785112

// KEY_Ocircumflexgrave is a representation of the C constant GDK_KEY_Ocircumflexgrave.
const KEY_Ocircumflexgrave = 16785106

// KEY_Ocircumflexhook is a representation of the C constant GDK_KEY_Ocircumflexhook.
const KEY_Ocircumflexhook = 16785108

// KEY_Ocircumflextilde is a representation of the C constant GDK_KEY_Ocircumflextilde.
const KEY_Ocircumflextilde = 16785110

// KEY_Odiaeresis is a representation of the C constant GDK_KEY_Odiaeresis.
const KEY_Odiaeresis = 214

// KEY_Odoubleacute is a representation of the C constant GDK_KEY_Odoubleacute.
const KEY_Odoubleacute = 469

// KEY_OfficeHome is a representation of the C constant GDK_KEY_OfficeHome.
const KEY_OfficeHome = 269025130

// KEY_Ograve is a representation of the C constant GDK_KEY_Ograve.
const KEY_Ograve = 210

// KEY_Ohook is a representation of the C constant GDK_KEY_Ohook.
const KEY_Ohook = 16785102

// KEY_Ohorn is a representation of the C constant GDK_KEY_Ohorn.
const KEY_Ohorn = 16777632

// KEY_Ohornacute is a representation of the C constant GDK_KEY_Ohornacute.
const KEY_Ohornacute = 16785114

// KEY_Ohornbelowdot is a representation of the C constant GDK_KEY_Ohornbelowdot.
const KEY_Ohornbelowdot = 16785122

// KEY_Ohorngrave is a representation of the C constant GDK_KEY_Ohorngrave.
const KEY_Ohorngrave = 16785116

// KEY_Ohornhook is a representation of the C constant GDK_KEY_Ohornhook.
const KEY_Ohornhook = 16785118

// KEY_Ohorntilde is a representation of the C constant GDK_KEY_Ohorntilde.
const KEY_Ohorntilde = 16785120

// KEY_Omacron is a representation of the C constant GDK_KEY_Omacron.
const KEY_Omacron = 978

// KEY_Ooblique is a representation of the C constant GDK_KEY_Ooblique.
const KEY_Ooblique = 216

// KEY_Open is a representation of the C constant GDK_KEY_Open.
const KEY_Open = 269025131

// KEY_OpenURL is a representation of the C constant GDK_KEY_OpenURL.
const KEY_OpenURL = 269025080

// KEY_Option is a representation of the C constant GDK_KEY_Option.
const KEY_Option = 269025132

// KEY_Oslash is a representation of the C constant GDK_KEY_Oslash.
const KEY_Oslash = 216

// KEY_Otilde is a representation of the C constant GDK_KEY_Otilde.
const KEY_Otilde = 213

// KEY_Overlay1_Enable is a representation of the C constant GDK_KEY_Overlay1_Enable.
const KEY_Overlay1_Enable = 65144

// KEY_Overlay2_Enable is a representation of the C constant GDK_KEY_Overlay2_Enable.
const KEY_Overlay2_Enable = 65145

// KEY_P is a representation of the C constant GDK_KEY_P.
const KEY_P = 80

// KEY_Pabovedot is a representation of the C constant GDK_KEY_Pabovedot.
const KEY_Pabovedot = 16784982

// KEY_Page_Down is a representation of the C constant GDK_KEY_Page_Down.
const KEY_Page_Down = 65366

// KEY_Page_Up is a representation of the C constant GDK_KEY_Page_Up.
const KEY_Page_Up = 65365

// KEY_Paste is a representation of the C constant GDK_KEY_Paste.
const KEY_Paste = 269025133

// KEY_Pause is a representation of the C constant GDK_KEY_Pause.
const KEY_Pause = 65299

// KEY_PesetaSign is a representation of the C constant GDK_KEY_PesetaSign.
const KEY_PesetaSign = 16785575

// KEY_Phone is a representation of the C constant GDK_KEY_Phone.
const KEY_Phone = 269025134

// KEY_Pictures is a representation of the C constant GDK_KEY_Pictures.
const KEY_Pictures = 269025169

// KEY_Pointer_Accelerate is a representation of the C constant GDK_KEY_Pointer_Accelerate.
const KEY_Pointer_Accelerate = 65274

// KEY_Pointer_Button1 is a representation of the C constant GDK_KEY_Pointer_Button1.
const KEY_Pointer_Button1 = 65257

// KEY_Pointer_Button2 is a representation of the C constant GDK_KEY_Pointer_Button2.
const KEY_Pointer_Button2 = 65258

// KEY_Pointer_Button3 is a representation of the C constant GDK_KEY_Pointer_Button3.
const KEY_Pointer_Button3 = 65259

// KEY_Pointer_Button4 is a representation of the C constant GDK_KEY_Pointer_Button4.
const KEY_Pointer_Button4 = 65260

// KEY_Pointer_Button5 is a representation of the C constant GDK_KEY_Pointer_Button5.
const KEY_Pointer_Button5 = 65261

// KEY_Pointer_Button_Dflt is a representation of the C constant GDK_KEY_Pointer_Button_Dflt.
const KEY_Pointer_Button_Dflt = 65256

// KEY_Pointer_DblClick1 is a representation of the C constant GDK_KEY_Pointer_DblClick1.
const KEY_Pointer_DblClick1 = 65263

// KEY_Pointer_DblClick2 is a representation of the C constant GDK_KEY_Pointer_DblClick2.
const KEY_Pointer_DblClick2 = 65264

// KEY_Pointer_DblClick3 is a representation of the C constant GDK_KEY_Pointer_DblClick3.
const KEY_Pointer_DblClick3 = 65265

// KEY_Pointer_DblClick4 is a representation of the C constant GDK_KEY_Pointer_DblClick4.
const KEY_Pointer_DblClick4 = 65266

// KEY_Pointer_DblClick5 is a representation of the C constant GDK_KEY_Pointer_DblClick5.
const KEY_Pointer_DblClick5 = 65267

// KEY_Pointer_DblClick_Dflt is a representation of the C constant GDK_KEY_Pointer_DblClick_Dflt.
const KEY_Pointer_DblClick_Dflt = 65262

// KEY_Pointer_DfltBtnNext is a representation of the C constant GDK_KEY_Pointer_DfltBtnNext.
const KEY_Pointer_DfltBtnNext = 65275

// KEY_Pointer_DfltBtnPrev is a representation of the C constant GDK_KEY_Pointer_DfltBtnPrev.
const KEY_Pointer_DfltBtnPrev = 65276

// KEY_Pointer_Down is a representation of the C constant GDK_KEY_Pointer_Down.
const KEY_Pointer_Down = 65251

// KEY_Pointer_DownLeft is a representation of the C constant GDK_KEY_Pointer_DownLeft.
const KEY_Pointer_DownLeft = 65254

// KEY_Pointer_DownRight is a representation of the C constant GDK_KEY_Pointer_DownRight.
const KEY_Pointer_DownRight = 65255

// KEY_Pointer_Drag1 is a representation of the C constant GDK_KEY_Pointer_Drag1.
const KEY_Pointer_Drag1 = 65269

// KEY_Pointer_Drag2 is a representation of the C constant GDK_KEY_Pointer_Drag2.
const KEY_Pointer_Drag2 = 65270

// KEY_Pointer_Drag3 is a representation of the C constant GDK_KEY_Pointer_Drag3.
const KEY_Pointer_Drag3 = 65271

// KEY_Pointer_Drag4 is a representation of the C constant GDK_KEY_Pointer_Drag4.
const KEY_Pointer_Drag4 = 65272

// KEY_Pointer_Drag5 is a representation of the C constant GDK_KEY_Pointer_Drag5.
const KEY_Pointer_Drag5 = 65277

// KEY_Pointer_Drag_Dflt is a representation of the C constant GDK_KEY_Pointer_Drag_Dflt.
const KEY_Pointer_Drag_Dflt = 65268

// KEY_Pointer_EnableKeys is a representation of the C constant GDK_KEY_Pointer_EnableKeys.
const KEY_Pointer_EnableKeys = 65273

// KEY_Pointer_Left is a representation of the C constant GDK_KEY_Pointer_Left.
const KEY_Pointer_Left = 65248

// KEY_Pointer_Right is a representation of the C constant GDK_KEY_Pointer_Right.
const KEY_Pointer_Right = 65249

// KEY_Pointer_Up is a representation of the C constant GDK_KEY_Pointer_Up.
const KEY_Pointer_Up = 65250

// KEY_Pointer_UpLeft is a representation of the C constant GDK_KEY_Pointer_UpLeft.
const KEY_Pointer_UpLeft = 65252

// KEY_Pointer_UpRight is a representation of the C constant GDK_KEY_Pointer_UpRight.
const KEY_Pointer_UpRight = 65253

// KEY_PowerDown is a representation of the C constant GDK_KEY_PowerDown.
const KEY_PowerDown = 269025057

// KEY_PowerOff is a representation of the C constant GDK_KEY_PowerOff.
const KEY_PowerOff = 269025066

// KEY_Prev_VMode is a representation of the C constant GDK_KEY_Prev_VMode.
const KEY_Prev_VMode = 269024803

// KEY_Prev_Virtual_Screen is a representation of the C constant GDK_KEY_Prev_Virtual_Screen.
const KEY_Prev_Virtual_Screen = 65233

// KEY_PreviousCandidate is a representation of the C constant GDK_KEY_PreviousCandidate.
const KEY_PreviousCandidate = 65342

// KEY_Print is a representation of the C constant GDK_KEY_Print.
const KEY_Print = 65377

// KEY_Prior is a representation of the C constant GDK_KEY_Prior.
const KEY_Prior = 65365

// KEY_Q is a representation of the C constant GDK_KEY_Q.
const KEY_Q = 81

// KEY_R is a representation of the C constant GDK_KEY_R.
const KEY_R = 82

// KEY_R1 is a representation of the C constant GDK_KEY_R1.
const KEY_R1 = 65490

// KEY_R10 is a representation of the C constant GDK_KEY_R10.
const KEY_R10 = 65499

// KEY_R11 is a representation of the C constant GDK_KEY_R11.
const KEY_R11 = 65500

// KEY_R12 is a representation of the C constant GDK_KEY_R12.
const KEY_R12 = 65501

// KEY_R13 is a representation of the C constant GDK_KEY_R13.
const KEY_R13 = 65502

// KEY_R14 is a representation of the C constant GDK_KEY_R14.
const KEY_R14 = 65503

// KEY_R15 is a representation of the C constant GDK_KEY_R15.
const KEY_R15 = 65504

// KEY_R2 is a representation of the C constant GDK_KEY_R2.
const KEY_R2 = 65491

// KEY_R3 is a representation of the C constant GDK_KEY_R3.
const KEY_R3 = 65492

// KEY_R4 is a representation of the C constant GDK_KEY_R4.
const KEY_R4 = 65493

// KEY_R5 is a representation of the C constant GDK_KEY_R5.
const KEY_R5 = 65494

// KEY_R6 is a representation of the C constant GDK_KEY_R6.
const KEY_R6 = 65495

// KEY_R7 is a representation of the C constant GDK_KEY_R7.
const KEY_R7 = 65496

// KEY_R8 is a representation of the C constant GDK_KEY_R8.
const KEY_R8 = 65497

// KEY_R9 is a representation of the C constant GDK_KEY_R9.
const KEY_R9 = 65498

// KEY_Racute is a representation of the C constant GDK_KEY_Racute.
const KEY_Racute = 448

// KEY_Rcaron is a representation of the C constant GDK_KEY_Rcaron.
const KEY_Rcaron = 472

// KEY_Rcedilla is a representation of the C constant GDK_KEY_Rcedilla.
const KEY_Rcedilla = 931

// KEY_Red is a representation of the C constant GDK_KEY_Red.
const KEY_Red = 269025187

// KEY_Redo is a representation of the C constant GDK_KEY_Redo.
const KEY_Redo = 65382

// KEY_Refresh is a representation of the C constant GDK_KEY_Refresh.
const KEY_Refresh = 269025065

// KEY_Reload is a representation of the C constant GDK_KEY_Reload.
const KEY_Reload = 269025139

// KEY_RepeatKeys_Enable is a representation of the C constant GDK_KEY_RepeatKeys_Enable.
const KEY_RepeatKeys_Enable = 65138

// KEY_Reply is a representation of the C constant GDK_KEY_Reply.
const KEY_Reply = 269025138

// KEY_Return is a representation of the C constant GDK_KEY_Return.
const KEY_Return = 65293

// KEY_Right is a representation of the C constant GDK_KEY_Right.
const KEY_Right = 65363

// KEY_RockerDown is a representation of the C constant GDK_KEY_RockerDown.
const KEY_RockerDown = 269025060

// KEY_RockerEnter is a representation of the C constant GDK_KEY_RockerEnter.
const KEY_RockerEnter = 269025061

// KEY_RockerUp is a representation of the C constant GDK_KEY_RockerUp.
const KEY_RockerUp = 269025059

// KEY_Romaji is a representation of the C constant GDK_KEY_Romaji.
const KEY_Romaji = 65316

// KEY_RotateWindows is a representation of the C constant GDK_KEY_RotateWindows.
const KEY_RotateWindows = 269025140

// KEY_RotationKB is a representation of the C constant GDK_KEY_RotationKB.
const KEY_RotationKB = 269025142

// KEY_RotationPB is a representation of the C constant GDK_KEY_RotationPB.
const KEY_RotationPB = 269025141

// KEY_RupeeSign is a representation of the C constant GDK_KEY_RupeeSign.
const KEY_RupeeSign = 16785576

// KEY_S is a representation of the C constant GDK_KEY_S.
const KEY_S = 83

// KEY_SCHWA is a representation of the C constant GDK_KEY_SCHWA.
const KEY_SCHWA = 16777615

// KEY_Sabovedot is a representation of the C constant GDK_KEY_Sabovedot.
const KEY_Sabovedot = 16784992

// KEY_Sacute is a representation of the C constant GDK_KEY_Sacute.
const KEY_Sacute = 422

// KEY_Save is a representation of the C constant GDK_KEY_Save.
const KEY_Save = 269025143

// KEY_Scaron is a representation of the C constant GDK_KEY_Scaron.
const KEY_Scaron = 425

// KEY_Scedilla is a representation of the C constant GDK_KEY_Scedilla.
const KEY_Scedilla = 426

// KEY_Scircumflex is a representation of the C constant GDK_KEY_Scircumflex.
const KEY_Scircumflex = 734

// KEY_ScreenSaver is a representation of the C constant GDK_KEY_ScreenSaver.
const KEY_ScreenSaver = 269025069

// KEY_ScrollClick is a representation of the C constant GDK_KEY_ScrollClick.
const KEY_ScrollClick = 269025146

// KEY_ScrollDown is a representation of the C constant GDK_KEY_ScrollDown.
const KEY_ScrollDown = 269025145

// KEY_ScrollUp is a representation of the C constant GDK_KEY_ScrollUp.
const KEY_ScrollUp = 269025144

// KEY_Scroll_Lock is a representation of the C constant GDK_KEY_Scroll_Lock.
const KEY_Scroll_Lock = 65300

// KEY_Search is a representation of the C constant GDK_KEY_Search.
const KEY_Search = 269025051

// KEY_Select is a representation of the C constant GDK_KEY_Select.
const KEY_Select = 65376

// KEY_SelectButton is a representation of the C constant GDK_KEY_SelectButton.
const KEY_SelectButton = 269025184

// KEY_Send is a representation of the C constant GDK_KEY_Send.
const KEY_Send = 269025147

// KEY_Serbian_DJE is a representation of the C constant GDK_KEY_Serbian_DJE.
const KEY_Serbian_DJE = 1713

// KEY_Serbian_DZE is a representation of the C constant GDK_KEY_Serbian_DZE.
const KEY_Serbian_DZE = 1727

// KEY_Serbian_JE is a representation of the C constant GDK_KEY_Serbian_JE.
const KEY_Serbian_JE = 1720

// KEY_Serbian_LJE is a representation of the C constant GDK_KEY_Serbian_LJE.
const KEY_Serbian_LJE = 1721

// KEY_Serbian_NJE is a representation of the C constant GDK_KEY_Serbian_NJE.
const KEY_Serbian_NJE = 1722

// KEY_Serbian_TSHE is a representation of the C constant GDK_KEY_Serbian_TSHE.
const KEY_Serbian_TSHE = 1723

// KEY_Serbian_dje is a representation of the C constant GDK_KEY_Serbian_dje.
const KEY_Serbian_dje = 1697

// KEY_Serbian_dze is a representation of the C constant GDK_KEY_Serbian_dze.
const KEY_Serbian_dze = 1711

// KEY_Serbian_je is a representation of the C constant GDK_KEY_Serbian_je.
const KEY_Serbian_je = 1704

// KEY_Serbian_lje is a representation of the C constant GDK_KEY_Serbian_lje.
const KEY_Serbian_lje = 1705

// KEY_Serbian_nje is a representation of the C constant GDK_KEY_Serbian_nje.
const KEY_Serbian_nje = 1706

// KEY_Serbian_tshe is a representation of the C constant GDK_KEY_Serbian_tshe.
const KEY_Serbian_tshe = 1707

// KEY_Shift_L is a representation of the C constant GDK_KEY_Shift_L.
const KEY_Shift_L = 65505

// KEY_Shift_Lock is a representation of the C constant GDK_KEY_Shift_Lock.
const KEY_Shift_Lock = 65510

// KEY_Shift_R is a representation of the C constant GDK_KEY_Shift_R.
const KEY_Shift_R = 65506

// KEY_Shop is a representation of the C constant GDK_KEY_Shop.
const KEY_Shop = 269025078

// KEY_SingleCandidate is a representation of the C constant GDK_KEY_SingleCandidate.
const KEY_SingleCandidate = 65340

// KEY_Sinh_a is a representation of the C constant GDK_KEY_Sinh_a.
const KEY_Sinh_a = 16780677

// KEY_Sinh_aa is a representation of the C constant GDK_KEY_Sinh_aa.
const KEY_Sinh_aa = 16780678

// KEY_Sinh_aa2 is a representation of the C constant GDK_KEY_Sinh_aa2.
const KEY_Sinh_aa2 = 16780751

// KEY_Sinh_ae is a representation of the C constant GDK_KEY_Sinh_ae.
const KEY_Sinh_ae = 16780679

// KEY_Sinh_ae2 is a representation of the C constant GDK_KEY_Sinh_ae2.
const KEY_Sinh_ae2 = 16780752

// KEY_Sinh_aee is a representation of the C constant GDK_KEY_Sinh_aee.
const KEY_Sinh_aee = 16780680

// KEY_Sinh_aee2 is a representation of the C constant GDK_KEY_Sinh_aee2.
const KEY_Sinh_aee2 = 16780753

// KEY_Sinh_ai is a representation of the C constant GDK_KEY_Sinh_ai.
const KEY_Sinh_ai = 16780691

// KEY_Sinh_ai2 is a representation of the C constant GDK_KEY_Sinh_ai2.
const KEY_Sinh_ai2 = 16780763

// KEY_Sinh_al is a representation of the C constant GDK_KEY_Sinh_al.
const KEY_Sinh_al = 16780746

// KEY_Sinh_au is a representation of the C constant GDK_KEY_Sinh_au.
const KEY_Sinh_au = 16780694

// KEY_Sinh_au2 is a representation of the C constant GDK_KEY_Sinh_au2.
const KEY_Sinh_au2 = 16780766

// KEY_Sinh_ba is a representation of the C constant GDK_KEY_Sinh_ba.
const KEY_Sinh_ba = 16780726

// KEY_Sinh_bha is a representation of the C constant GDK_KEY_Sinh_bha.
const KEY_Sinh_bha = 16780727

// KEY_Sinh_ca is a representation of the C constant GDK_KEY_Sinh_ca.
const KEY_Sinh_ca = 16780704

// KEY_Sinh_cha is a representation of the C constant GDK_KEY_Sinh_cha.
const KEY_Sinh_cha = 16780705

// KEY_Sinh_dda is a representation of the C constant GDK_KEY_Sinh_dda.
const KEY_Sinh_dda = 16780713

// KEY_Sinh_ddha is a representation of the C constant GDK_KEY_Sinh_ddha.
const KEY_Sinh_ddha = 16780714

// KEY_Sinh_dha is a representation of the C constant GDK_KEY_Sinh_dha.
const KEY_Sinh_dha = 16780719

// KEY_Sinh_dhha is a representation of the C constant GDK_KEY_Sinh_dhha.
const KEY_Sinh_dhha = 16780720

// KEY_Sinh_e is a representation of the C constant GDK_KEY_Sinh_e.
const KEY_Sinh_e = 16780689

// KEY_Sinh_e2 is a representation of the C constant GDK_KEY_Sinh_e2.
const KEY_Sinh_e2 = 16780761

// KEY_Sinh_ee is a representation of the C constant GDK_KEY_Sinh_ee.
const KEY_Sinh_ee = 16780690

// KEY_Sinh_ee2 is a representation of the C constant GDK_KEY_Sinh_ee2.
const KEY_Sinh_ee2 = 16780762

// KEY_Sinh_fa is a representation of the C constant GDK_KEY_Sinh_fa.
const KEY_Sinh_fa = 16780742

// KEY_Sinh_ga is a representation of the C constant GDK_KEY_Sinh_ga.
const KEY_Sinh_ga = 16780700

// KEY_Sinh_gha is a representation of the C constant GDK_KEY_Sinh_gha.
const KEY_Sinh_gha = 16780701

// KEY_Sinh_h2 is a representation of the C constant GDK_KEY_Sinh_h2.
const KEY_Sinh_h2 = 16780675

// KEY_Sinh_ha is a representation of the C constant GDK_KEY_Sinh_ha.
const KEY_Sinh_ha = 16780740

// KEY_Sinh_i is a representation of the C constant GDK_KEY_Sinh_i.
const KEY_Sinh_i = 16780681

// KEY_Sinh_i2 is a representation of the C constant GDK_KEY_Sinh_i2.
const KEY_Sinh_i2 = 16780754

// KEY_Sinh_ii is a representation of the C constant GDK_KEY_Sinh_ii.
const KEY_Sinh_ii = 16780682

// KEY_Sinh_ii2 is a representation of the C constant GDK_KEY_Sinh_ii2.
const KEY_Sinh_ii2 = 16780755

// KEY_Sinh_ja is a representation of the C constant GDK_KEY_Sinh_ja.
const KEY_Sinh_ja = 16780706

// KEY_Sinh_jha is a representation of the C constant GDK_KEY_Sinh_jha.
const KEY_Sinh_jha = 16780707

// KEY_Sinh_jnya is a representation of the C constant GDK_KEY_Sinh_jnya.
const KEY_Sinh_jnya = 16780709

// KEY_Sinh_ka is a representation of the C constant GDK_KEY_Sinh_ka.
const KEY_Sinh_ka = 16780698

// KEY_Sinh_kha is a representation of the C constant GDK_KEY_Sinh_kha.
const KEY_Sinh_kha = 16780699

// KEY_Sinh_kunddaliya is a representation of the C constant GDK_KEY_Sinh_kunddaliya.
const KEY_Sinh_kunddaliya = 16780788

// KEY_Sinh_la is a representation of the C constant GDK_KEY_Sinh_la.
const KEY_Sinh_la = 16780733

// KEY_Sinh_lla is a representation of the C constant GDK_KEY_Sinh_lla.
const KEY_Sinh_lla = 16780741

// KEY_Sinh_lu is a representation of the C constant GDK_KEY_Sinh_lu.
const KEY_Sinh_lu = 16780687

// KEY_Sinh_lu2 is a representation of the C constant GDK_KEY_Sinh_lu2.
const KEY_Sinh_lu2 = 16780767

// KEY_Sinh_luu is a representation of the C constant GDK_KEY_Sinh_luu.
const KEY_Sinh_luu = 16780688

// KEY_Sinh_luu2 is a representation of the C constant GDK_KEY_Sinh_luu2.
const KEY_Sinh_luu2 = 16780787

// KEY_Sinh_ma is a representation of the C constant GDK_KEY_Sinh_ma.
const KEY_Sinh_ma = 16780728

// KEY_Sinh_mba is a representation of the C constant GDK_KEY_Sinh_mba.
const KEY_Sinh_mba = 16780729

// KEY_Sinh_na is a representation of the C constant GDK_KEY_Sinh_na.
const KEY_Sinh_na = 16780721

// KEY_Sinh_ndda is a representation of the C constant GDK_KEY_Sinh_ndda.
const KEY_Sinh_ndda = 16780716

// KEY_Sinh_ndha is a representation of the C constant GDK_KEY_Sinh_ndha.
const KEY_Sinh_ndha = 16780723

// KEY_Sinh_ng is a representation of the C constant GDK_KEY_Sinh_ng.
const KEY_Sinh_ng = 16780674

// KEY_Sinh_ng2 is a representation of the C constant GDK_KEY_Sinh_ng2.
const KEY_Sinh_ng2 = 16780702

// KEY_Sinh_nga is a representation of the C constant GDK_KEY_Sinh_nga.
const KEY_Sinh_nga = 16780703

// KEY_Sinh_nja is a representation of the C constant GDK_KEY_Sinh_nja.
const KEY_Sinh_nja = 16780710

// KEY_Sinh_nna is a representation of the C constant GDK_KEY_Sinh_nna.
const KEY_Sinh_nna = 16780715

// KEY_Sinh_nya is a representation of the C constant GDK_KEY_Sinh_nya.
const KEY_Sinh_nya = 16780708

// KEY_Sinh_o is a representation of the C constant GDK_KEY_Sinh_o.
const KEY_Sinh_o = 16780692

// KEY_Sinh_o2 is a representation of the C constant GDK_KEY_Sinh_o2.
const KEY_Sinh_o2 = 16780764

// KEY_Sinh_oo is a representation of the C constant GDK_KEY_Sinh_oo.
const KEY_Sinh_oo = 16780693

// KEY_Sinh_oo2 is a representation of the C constant GDK_KEY_Sinh_oo2.
const KEY_Sinh_oo2 = 16780765

// KEY_Sinh_pa is a representation of the C constant GDK_KEY_Sinh_pa.
const KEY_Sinh_pa = 16780724

// KEY_Sinh_pha is a representation of the C constant GDK_KEY_Sinh_pha.
const KEY_Sinh_pha = 16780725

// KEY_Sinh_ra is a representation of the C constant GDK_KEY_Sinh_ra.
const KEY_Sinh_ra = 16780731

// KEY_Sinh_ri is a representation of the C constant GDK_KEY_Sinh_ri.
const KEY_Sinh_ri = 16780685

// KEY_Sinh_rii is a representation of the C constant GDK_KEY_Sinh_rii.
const KEY_Sinh_rii = 16780686

// KEY_Sinh_ru2 is a representation of the C constant GDK_KEY_Sinh_ru2.
const KEY_Sinh_ru2 = 16780760

// KEY_Sinh_ruu2 is a representation of the C constant GDK_KEY_Sinh_ruu2.
const KEY_Sinh_ruu2 = 16780786

// KEY_Sinh_sa is a representation of the C constant GDK_KEY_Sinh_sa.
const KEY_Sinh_sa = 16780739

// KEY_Sinh_sha is a representation of the C constant GDK_KEY_Sinh_sha.
const KEY_Sinh_sha = 16780737

// KEY_Sinh_ssha is a representation of the C constant GDK_KEY_Sinh_ssha.
const KEY_Sinh_ssha = 16780738

// KEY_Sinh_tha is a representation of the C constant GDK_KEY_Sinh_tha.
const KEY_Sinh_tha = 16780717

// KEY_Sinh_thha is a representation of the C constant GDK_KEY_Sinh_thha.
const KEY_Sinh_thha = 16780718

// KEY_Sinh_tta is a representation of the C constant GDK_KEY_Sinh_tta.
const KEY_Sinh_tta = 16780711

// KEY_Sinh_ttha is a representation of the C constant GDK_KEY_Sinh_ttha.
const KEY_Sinh_ttha = 16780712

// KEY_Sinh_u is a representation of the C constant GDK_KEY_Sinh_u.
const KEY_Sinh_u = 16780683

// KEY_Sinh_u2 is a representation of the C constant GDK_KEY_Sinh_u2.
const KEY_Sinh_u2 = 16780756

// KEY_Sinh_uu is a representation of the C constant GDK_KEY_Sinh_uu.
const KEY_Sinh_uu = 16780684

// KEY_Sinh_uu2 is a representation of the C constant GDK_KEY_Sinh_uu2.
const KEY_Sinh_uu2 = 16780758

// KEY_Sinh_va is a representation of the C constant GDK_KEY_Sinh_va.
const KEY_Sinh_va = 16780736

// KEY_Sinh_ya is a representation of the C constant GDK_KEY_Sinh_ya.
const KEY_Sinh_ya = 16780730

// KEY_Sleep is a representation of the C constant GDK_KEY_Sleep.
const KEY_Sleep = 269025071

// KEY_SlowKeys_Enable is a representation of the C constant GDK_KEY_SlowKeys_Enable.
const KEY_SlowKeys_Enable = 65139

// KEY_Spell is a representation of the C constant GDK_KEY_Spell.
const KEY_Spell = 269025148

// KEY_SplitScreen is a representation of the C constant GDK_KEY_SplitScreen.
const KEY_SplitScreen = 269025149

// KEY_Standby is a representation of the C constant GDK_KEY_Standby.
const KEY_Standby = 269025040

// KEY_Start is a representation of the C constant GDK_KEY_Start.
const KEY_Start = 269025050

// KEY_StickyKeys_Enable is a representation of the C constant GDK_KEY_StickyKeys_Enable.
const KEY_StickyKeys_Enable = 65141

// KEY_Stop is a representation of the C constant GDK_KEY_Stop.
const KEY_Stop = 269025064

// KEY_Subtitle is a representation of the C constant GDK_KEY_Subtitle.
const KEY_Subtitle = 269025178

// KEY_Super_L is a representation of the C constant GDK_KEY_Super_L.
const KEY_Super_L = 65515

// KEY_Super_R is a representation of the C constant GDK_KEY_Super_R.
const KEY_Super_R = 65516

// KEY_Support is a representation of the C constant GDK_KEY_Support.
const KEY_Support = 269025150

// KEY_Suspend is a representation of the C constant GDK_KEY_Suspend.
const KEY_Suspend = 269025191

// KEY_Switch_VT_1 is a representation of the C constant GDK_KEY_Switch_VT_1.
const KEY_Switch_VT_1 = 269024769

// KEY_Switch_VT_10 is a representation of the C constant GDK_KEY_Switch_VT_10.
const KEY_Switch_VT_10 = 269024778

// KEY_Switch_VT_11 is a representation of the C constant GDK_KEY_Switch_VT_11.
const KEY_Switch_VT_11 = 269024779

// KEY_Switch_VT_12 is a representation of the C constant GDK_KEY_Switch_VT_12.
const KEY_Switch_VT_12 = 269024780

// KEY_Switch_VT_2 is a representation of the C constant GDK_KEY_Switch_VT_2.
const KEY_Switch_VT_2 = 269024770

// KEY_Switch_VT_3 is a representation of the C constant GDK_KEY_Switch_VT_3.
const KEY_Switch_VT_3 = 269024771

// KEY_Switch_VT_4 is a representation of the C constant GDK_KEY_Switch_VT_4.
const KEY_Switch_VT_4 = 269024772

// KEY_Switch_VT_5 is a representation of the C constant GDK_KEY_Switch_VT_5.
const KEY_Switch_VT_5 = 269024773

// KEY_Switch_VT_6 is a representation of the C constant GDK_KEY_Switch_VT_6.
const KEY_Switch_VT_6 = 269024774

// KEY_Switch_VT_7 is a representation of the C constant GDK_KEY_Switch_VT_7.
const KEY_Switch_VT_7 = 269024775

// KEY_Switch_VT_8 is a representation of the C constant GDK_KEY_Switch_VT_8.
const KEY_Switch_VT_8 = 269024776

// KEY_Switch_VT_9 is a representation of the C constant GDK_KEY_Switch_VT_9.
const KEY_Switch_VT_9 = 269024777

// KEY_Sys_Req is a representation of the C constant GDK_KEY_Sys_Req.
const KEY_Sys_Req = 65301

// KEY_T is a representation of the C constant GDK_KEY_T.
const KEY_T = 84

// KEY_THORN is a representation of the C constant GDK_KEY_THORN.
const KEY_THORN = 222

// KEY_Tab is a representation of the C constant GDK_KEY_Tab.
const KEY_Tab = 65289

// KEY_Tabovedot is a representation of the C constant GDK_KEY_Tabovedot.
const KEY_Tabovedot = 16785002

// KEY_TaskPane is a representation of the C constant GDK_KEY_TaskPane.
const KEY_TaskPane = 269025151

// KEY_Tcaron is a representation of the C constant GDK_KEY_Tcaron.
const KEY_Tcaron = 427

// KEY_Tcedilla is a representation of the C constant GDK_KEY_Tcedilla.
const KEY_Tcedilla = 478

// KEY_Terminal is a representation of the C constant GDK_KEY_Terminal.
const KEY_Terminal = 269025152

// KEY_Terminate_Server is a representation of the C constant GDK_KEY_Terminate_Server.
const KEY_Terminate_Server = 65237

// KEY_Thai_baht is a representation of the C constant GDK_KEY_Thai_baht.
const KEY_Thai_baht = 3551

// KEY_Thai_bobaimai is a representation of the C constant GDK_KEY_Thai_bobaimai.
const KEY_Thai_bobaimai = 3514

// KEY_Thai_chochan is a representation of the C constant GDK_KEY_Thai_chochan.
const KEY_Thai_chochan = 3496

// KEY_Thai_chochang is a representation of the C constant GDK_KEY_Thai_chochang.
const KEY_Thai_chochang = 3498

// KEY_Thai_choching is a representation of the C constant GDK_KEY_Thai_choching.
const KEY_Thai_choching = 3497

// KEY_Thai_chochoe is a representation of the C constant GDK_KEY_Thai_chochoe.
const KEY_Thai_chochoe = 3500

// KEY_Thai_dochada is a representation of the C constant GDK_KEY_Thai_dochada.
const KEY_Thai_dochada = 3502

// KEY_Thai_dodek is a representation of the C constant GDK_KEY_Thai_dodek.
const KEY_Thai_dodek = 3508

// KEY_Thai_fofa is a representation of the C constant GDK_KEY_Thai_fofa.
const KEY_Thai_fofa = 3517

// KEY_Thai_fofan is a representation of the C constant GDK_KEY_Thai_fofan.
const KEY_Thai_fofan = 3519

// KEY_Thai_hohip is a representation of the C constant GDK_KEY_Thai_hohip.
const KEY_Thai_hohip = 3531

// KEY_Thai_honokhuk is a representation of the C constant GDK_KEY_Thai_honokhuk.
const KEY_Thai_honokhuk = 3534

// KEY_Thai_khokhai is a representation of the C constant GDK_KEY_Thai_khokhai.
const KEY_Thai_khokhai = 3490

// KEY_Thai_khokhon is a representation of the C constant GDK_KEY_Thai_khokhon.
const KEY_Thai_khokhon = 3493

// KEY_Thai_khokhuat is a representation of the C constant GDK_KEY_Thai_khokhuat.
const KEY_Thai_khokhuat = 3491

// KEY_Thai_khokhwai is a representation of the C constant GDK_KEY_Thai_khokhwai.
const KEY_Thai_khokhwai = 3492

// KEY_Thai_khorakhang is a representation of the C constant GDK_KEY_Thai_khorakhang.
const KEY_Thai_khorakhang = 3494

// KEY_Thai_kokai is a representation of the C constant GDK_KEY_Thai_kokai.
const KEY_Thai_kokai = 3489

// KEY_Thai_lakkhangyao is a representation of the C constant GDK_KEY_Thai_lakkhangyao.
const KEY_Thai_lakkhangyao = 3557

// KEY_Thai_lekchet is a representation of the C constant GDK_KEY_Thai_lekchet.
const KEY_Thai_lekchet = 3575

// KEY_Thai_lekha is a representation of the C constant GDK_KEY_Thai_lekha.
const KEY_Thai_lekha = 3573

// KEY_Thai_lekhok is a representation of the C constant GDK_KEY_Thai_lekhok.
const KEY_Thai_lekhok = 3574

// KEY_Thai_lekkao is a representation of the C constant GDK_KEY_Thai_lekkao.
const KEY_Thai_lekkao = 3577

// KEY_Thai_leknung is a representation of the C constant GDK_KEY_Thai_leknung.
const KEY_Thai_leknung = 3569

// KEY_Thai_lekpaet is a representation of the C constant GDK_KEY_Thai_lekpaet.
const KEY_Thai_lekpaet = 3576

// KEY_Thai_leksam is a representation of the C constant GDK_KEY_Thai_leksam.
const KEY_Thai_leksam = 3571

// KEY_Thai_leksi is a representation of the C constant GDK_KEY_Thai_leksi.
const KEY_Thai_leksi = 3572

// KEY_Thai_leksong is a representation of the C constant GDK_KEY_Thai_leksong.
const KEY_Thai_leksong = 3570

// KEY_Thai_leksun is a representation of the C constant GDK_KEY_Thai_leksun.
const KEY_Thai_leksun = 3568

// KEY_Thai_lochula is a representation of the C constant GDK_KEY_Thai_lochula.
const KEY_Thai_lochula = 3532

// KEY_Thai_loling is a representation of the C constant GDK_KEY_Thai_loling.
const KEY_Thai_loling = 3525

// KEY_Thai_lu is a representation of the C constant GDK_KEY_Thai_lu.
const KEY_Thai_lu = 3526

// KEY_Thai_maichattawa is a representation of the C constant GDK_KEY_Thai_maichattawa.
const KEY_Thai_maichattawa = 3563

// KEY_Thai_maiek is a representation of the C constant GDK_KEY_Thai_maiek.
const KEY_Thai_maiek = 3560

// KEY_Thai_maihanakat is a representation of the C constant GDK_KEY_Thai_maihanakat.
const KEY_Thai_maihanakat = 3537

// KEY_Thai_maihanakat_maitho is a representation of the C constant GDK_KEY_Thai_maihanakat_maitho.
const KEY_Thai_maihanakat_maitho = 3550

// KEY_Thai_maitaikhu is a representation of the C constant GDK_KEY_Thai_maitaikhu.
const KEY_Thai_maitaikhu = 3559

// KEY_Thai_maitho is a representation of the C constant GDK_KEY_Thai_maitho.
const KEY_Thai_maitho = 3561

// KEY_Thai_maitri is a representation of the C constant GDK_KEY_Thai_maitri.
const KEY_Thai_maitri = 3562

// KEY_Thai_maiyamok is a representation of the C constant GDK_KEY_Thai_maiyamok.
const KEY_Thai_maiyamok = 3558

// KEY_Thai_moma is a representation of the C constant GDK_KEY_Thai_moma.
const KEY_Thai_moma = 3521

// KEY_Thai_ngongu is a representation of the C constant GDK_KEY_Thai_ngongu.
const KEY_Thai_ngongu = 3495

// KEY_Thai_nikhahit is a representation of the C constant GDK_KEY_Thai_nikhahit.
const KEY_Thai_nikhahit = 3565

// KEY_Thai_nonen is a representation of the C constant GDK_KEY_Thai_nonen.
const KEY_Thai_nonen = 3507

// KEY_Thai_nonu is a representation of the C constant GDK_KEY_Thai_nonu.
const KEY_Thai_nonu = 3513

// KEY_Thai_oang is a representation of the C constant GDK_KEY_Thai_oang.
const KEY_Thai_oang = 3533

// KEY_Thai_paiyannoi is a representation of the C constant GDK_KEY_Thai_paiyannoi.
const KEY_Thai_paiyannoi = 3535

// KEY_Thai_phinthu is a representation of the C constant GDK_KEY_Thai_phinthu.
const KEY_Thai_phinthu = 3546

// KEY_Thai_phophan is a representation of the C constant GDK_KEY_Thai_phophan.
const KEY_Thai_phophan = 3518

// KEY_Thai_phophung is a representation of the C constant GDK_KEY_Thai_phophung.
const KEY_Thai_phophung = 3516

// KEY_Thai_phosamphao is a representation of the C constant GDK_KEY_Thai_phosamphao.
const KEY_Thai_phosamphao = 3520

// KEY_Thai_popla is a representation of the C constant GDK_KEY_Thai_popla.
const KEY_Thai_popla = 3515

// KEY_Thai_rorua is a representation of the C constant GDK_KEY_Thai_rorua.
const KEY_Thai_rorua = 3523

// KEY_Thai_ru is a representation of the C constant GDK_KEY_Thai_ru.
const KEY_Thai_ru = 3524

// KEY_Thai_saraa is a representation of the C constant GDK_KEY_Thai_saraa.
const KEY_Thai_saraa = 3536

// KEY_Thai_saraaa is a representation of the C constant GDK_KEY_Thai_saraaa.
const KEY_Thai_saraaa = 3538

// KEY_Thai_saraae is a representation of the C constant GDK_KEY_Thai_saraae.
const KEY_Thai_saraae = 3553

// KEY_Thai_saraaimaimalai is a representation of the C constant GDK_KEY_Thai_saraaimaimalai.
const KEY_Thai_saraaimaimalai = 3556

// KEY_Thai_saraaimaimuan is a representation of the C constant GDK_KEY_Thai_saraaimaimuan.
const KEY_Thai_saraaimaimuan = 3555

// KEY_Thai_saraam is a representation of the C constant GDK_KEY_Thai_saraam.
const KEY_Thai_saraam = 3539

// KEY_Thai_sarae is a representation of the C constant GDK_KEY_Thai_sarae.
const KEY_Thai_sarae = 3552

// KEY_Thai_sarai is a representation of the C constant GDK_KEY_Thai_sarai.
const KEY_Thai_sarai = 3540

// KEY_Thai_saraii is a representation of the C constant GDK_KEY_Thai_saraii.
const KEY_Thai_saraii = 3541

// KEY_Thai_sarao is a representation of the C constant GDK_KEY_Thai_sarao.
const KEY_Thai_sarao = 3554

// KEY_Thai_sarau is a representation of the C constant GDK_KEY_Thai_sarau.
const KEY_Thai_sarau = 3544

// KEY_Thai_saraue is a representation of the C constant GDK_KEY_Thai_saraue.
const KEY_Thai_saraue = 3542

// KEY_Thai_sarauee is a representation of the C constant GDK_KEY_Thai_sarauee.
const KEY_Thai_sarauee = 3543

// KEY_Thai_sarauu is a representation of the C constant GDK_KEY_Thai_sarauu.
const KEY_Thai_sarauu = 3545

// KEY_Thai_sorusi is a representation of the C constant GDK_KEY_Thai_sorusi.
const KEY_Thai_sorusi = 3529

// KEY_Thai_sosala is a representation of the C constant GDK_KEY_Thai_sosala.
const KEY_Thai_sosala = 3528

// KEY_Thai_soso is a representation of the C constant GDK_KEY_Thai_soso.
const KEY_Thai_soso = 3499

// KEY_Thai_sosua is a representation of the C constant GDK_KEY_Thai_sosua.
const KEY_Thai_sosua = 3530

// KEY_Thai_thanthakhat is a representation of the C constant GDK_KEY_Thai_thanthakhat.
const KEY_Thai_thanthakhat = 3564

// KEY_Thai_thonangmontho is a representation of the C constant GDK_KEY_Thai_thonangmontho.
const KEY_Thai_thonangmontho = 3505

// KEY_Thai_thophuthao is a representation of the C constant GDK_KEY_Thai_thophuthao.
const KEY_Thai_thophuthao = 3506

// KEY_Thai_thothahan is a representation of the C constant GDK_KEY_Thai_thothahan.
const KEY_Thai_thothahan = 3511

// KEY_Thai_thothan is a representation of the C constant GDK_KEY_Thai_thothan.
const KEY_Thai_thothan = 3504

// KEY_Thai_thothong is a representation of the C constant GDK_KEY_Thai_thothong.
const KEY_Thai_thothong = 3512

// KEY_Thai_thothung is a representation of the C constant GDK_KEY_Thai_thothung.
const KEY_Thai_thothung = 3510

// KEY_Thai_topatak is a representation of the C constant GDK_KEY_Thai_topatak.
const KEY_Thai_topatak = 3503

// KEY_Thai_totao is a representation of the C constant GDK_KEY_Thai_totao.
const KEY_Thai_totao = 3509

// KEY_Thai_wowaen is a representation of the C constant GDK_KEY_Thai_wowaen.
const KEY_Thai_wowaen = 3527

// KEY_Thai_yoyak is a representation of the C constant GDK_KEY_Thai_yoyak.
const KEY_Thai_yoyak = 3522

// KEY_Thai_yoying is a representation of the C constant GDK_KEY_Thai_yoying.
const KEY_Thai_yoying = 3501

// KEY_Thorn is a representation of the C constant GDK_KEY_Thorn.
const KEY_Thorn = 222

// KEY_Time is a representation of the C constant GDK_KEY_Time.
const KEY_Time = 269025183

// KEY_ToDoList is a representation of the C constant GDK_KEY_ToDoList.
const KEY_ToDoList = 269025055

// KEY_Tools is a representation of the C constant GDK_KEY_Tools.
const KEY_Tools = 269025153

// KEY_TopMenu is a representation of the C constant GDK_KEY_TopMenu.
const KEY_TopMenu = 269025186

// KEY_TouchpadOff is a representation of the C constant GDK_KEY_TouchpadOff.
const KEY_TouchpadOff = 269025201

// KEY_TouchpadOn is a representation of the C constant GDK_KEY_TouchpadOn.
const KEY_TouchpadOn = 269025200

// KEY_TouchpadToggle is a representation of the C constant GDK_KEY_TouchpadToggle.
const KEY_TouchpadToggle = 269025193

// KEY_Touroku is a representation of the C constant GDK_KEY_Touroku.
const KEY_Touroku = 65323

// KEY_Travel is a representation of the C constant GDK_KEY_Travel.
const KEY_Travel = 269025154

// KEY_Tslash is a representation of the C constant GDK_KEY_Tslash.
const KEY_Tslash = 940

// KEY_U is a representation of the C constant GDK_KEY_U.
const KEY_U = 85

// KEY_UWB is a representation of the C constant GDK_KEY_UWB.
const KEY_UWB = 269025174

// KEY_Uacute is a representation of the C constant GDK_KEY_Uacute.
const KEY_Uacute = 218

// KEY_Ubelowdot is a representation of the C constant GDK_KEY_Ubelowdot.
const KEY_Ubelowdot = 16785124

// KEY_Ubreve is a representation of the C constant GDK_KEY_Ubreve.
const KEY_Ubreve = 733

// KEY_Ucircumflex is a representation of the C constant GDK_KEY_Ucircumflex.
const KEY_Ucircumflex = 219

// KEY_Udiaeresis is a representation of the C constant GDK_KEY_Udiaeresis.
const KEY_Udiaeresis = 220

// KEY_Udoubleacute is a representation of the C constant GDK_KEY_Udoubleacute.
const KEY_Udoubleacute = 475

// KEY_Ugrave is a representation of the C constant GDK_KEY_Ugrave.
const KEY_Ugrave = 217

// KEY_Uhook is a representation of the C constant GDK_KEY_Uhook.
const KEY_Uhook = 16785126

// KEY_Uhorn is a representation of the C constant GDK_KEY_Uhorn.
const KEY_Uhorn = 16777647

// KEY_Uhornacute is a representation of the C constant GDK_KEY_Uhornacute.
const KEY_Uhornacute = 16785128

// KEY_Uhornbelowdot is a representation of the C constant GDK_KEY_Uhornbelowdot.
const KEY_Uhornbelowdot = 16785136

// KEY_Uhorngrave is a representation of the C constant GDK_KEY_Uhorngrave.
const KEY_Uhorngrave = 16785130

// KEY_Uhornhook is a representation of the C constant GDK_KEY_Uhornhook.
const KEY_Uhornhook = 16785132

// KEY_Uhorntilde is a representation of the C constant GDK_KEY_Uhorntilde.
const KEY_Uhorntilde = 16785134

// KEY_Ukrainian_GHE_WITH_UPTURN is a representation of the C constant GDK_KEY_Ukrainian_GHE_WITH_UPTURN.
const KEY_Ukrainian_GHE_WITH_UPTURN = 1725

// KEY_Ukrainian_I is a representation of the C constant GDK_KEY_Ukrainian_I.
const KEY_Ukrainian_I = 1718

// KEY_Ukrainian_IE is a representation of the C constant GDK_KEY_Ukrainian_IE.
const KEY_Ukrainian_IE = 1716

// KEY_Ukrainian_YI is a representation of the C constant GDK_KEY_Ukrainian_YI.
const KEY_Ukrainian_YI = 1719

// KEY_Ukrainian_ghe_with_upturn is a representation of the C constant GDK_KEY_Ukrainian_ghe_with_upturn.
const KEY_Ukrainian_ghe_with_upturn = 1709

// KEY_Ukrainian_i is a representation of the C constant GDK_KEY_Ukrainian_i.
const KEY_Ukrainian_i = 1702

// KEY_Ukrainian_ie is a representation of the C constant GDK_KEY_Ukrainian_ie.
const KEY_Ukrainian_ie = 1700

// KEY_Ukrainian_yi is a representation of the C constant GDK_KEY_Ukrainian_yi.
const KEY_Ukrainian_yi = 1703

// KEY_Ukranian_I is a representation of the C constant GDK_KEY_Ukranian_I.
const KEY_Ukranian_I = 1718

// KEY_Ukranian_JE is a representation of the C constant GDK_KEY_Ukranian_JE.
const KEY_Ukranian_JE = 1716

// KEY_Ukranian_YI is a representation of the C constant GDK_KEY_Ukranian_YI.
const KEY_Ukranian_YI = 1719

// KEY_Ukranian_i is a representation of the C constant GDK_KEY_Ukranian_i.
const KEY_Ukranian_i = 1702

// KEY_Ukranian_je is a representation of the C constant GDK_KEY_Ukranian_je.
const KEY_Ukranian_je = 1700

// KEY_Ukranian_yi is a representation of the C constant GDK_KEY_Ukranian_yi.
const KEY_Ukranian_yi = 1703

// KEY_Umacron is a representation of the C constant GDK_KEY_Umacron.
const KEY_Umacron = 990

// KEY_Undo is a representation of the C constant GDK_KEY_Undo.
const KEY_Undo = 65381

// KEY_Ungrab is a representation of the C constant GDK_KEY_Ungrab.
const KEY_Ungrab = 269024800

// KEY_Uogonek is a representation of the C constant GDK_KEY_Uogonek.
const KEY_Uogonek = 985

// KEY_Up is a representation of the C constant GDK_KEY_Up.
const KEY_Up = 65362

// KEY_Uring is a representation of the C constant GDK_KEY_Uring.
const KEY_Uring = 473

// KEY_User1KB is a representation of the C constant GDK_KEY_User1KB.
const KEY_User1KB = 269025157

// KEY_User2KB is a representation of the C constant GDK_KEY_User2KB.
const KEY_User2KB = 269025158

// KEY_UserPB is a representation of the C constant GDK_KEY_UserPB.
const KEY_UserPB = 269025156

// KEY_Utilde is a representation of the C constant GDK_KEY_Utilde.
const KEY_Utilde = 989

// KEY_V is a representation of the C constant GDK_KEY_V.
const KEY_V = 86

// KEY_VendorHome is a representation of the C constant GDK_KEY_VendorHome.
const KEY_VendorHome = 269025076

// KEY_Video is a representation of the C constant GDK_KEY_Video.
const KEY_Video = 269025159

// KEY_View is a representation of the C constant GDK_KEY_View.
const KEY_View = 269025185

// KEY_VoidSymbol is a representation of the C constant GDK_KEY_VoidSymbol.
const KEY_VoidSymbol = 16777215

// KEY_W is a representation of the C constant GDK_KEY_W.
const KEY_W = 87

// KEY_WLAN is a representation of the C constant GDK_KEY_WLAN.
const KEY_WLAN = 269025173

// KEY_WWW is a representation of the C constant GDK_KEY_WWW.
const KEY_WWW = 269025070

// KEY_Wacute is a representation of the C constant GDK_KEY_Wacute.
const KEY_Wacute = 16785026

// KEY_WakeUp is a representation of the C constant GDK_KEY_WakeUp.
const KEY_WakeUp = 269025067

// KEY_Wcircumflex is a representation of the C constant GDK_KEY_Wcircumflex.
const KEY_Wcircumflex = 16777588

// KEY_Wdiaeresis is a representation of the C constant GDK_KEY_Wdiaeresis.
const KEY_Wdiaeresis = 16785028

// KEY_WebCam is a representation of the C constant GDK_KEY_WebCam.
const KEY_WebCam = 269025167

// KEY_Wgrave is a representation of the C constant GDK_KEY_Wgrave.
const KEY_Wgrave = 16785024

// KEY_WheelButton is a representation of the C constant GDK_KEY_WheelButton.
const KEY_WheelButton = 269025160

// KEY_WindowClear is a representation of the C constant GDK_KEY_WindowClear.
const KEY_WindowClear = 269025109

// KEY_WonSign is a representation of the C constant GDK_KEY_WonSign.
const KEY_WonSign = 16785577

// KEY_Word is a representation of the C constant GDK_KEY_Word.
const KEY_Word = 269025161

// KEY_X is a representation of the C constant GDK_KEY_X.
const KEY_X = 88

// KEY_Xabovedot is a representation of the C constant GDK_KEY_Xabovedot.
const KEY_Xabovedot = 16785034

// KEY_Xfer is a representation of the C constant GDK_KEY_Xfer.
const KEY_Xfer = 269025162

// KEY_Y is a representation of the C constant GDK_KEY_Y.
const KEY_Y = 89

// KEY_Yacute is a representation of the C constant GDK_KEY_Yacute.
const KEY_Yacute = 221

// KEY_Ybelowdot is a representation of the C constant GDK_KEY_Ybelowdot.
const KEY_Ybelowdot = 16785140

// KEY_Ycircumflex is a representation of the C constant GDK_KEY_Ycircumflex.
const KEY_Ycircumflex = 16777590

// KEY_Ydiaeresis is a representation of the C constant GDK_KEY_Ydiaeresis.
const KEY_Ydiaeresis = 5054

// KEY_Yellow is a representation of the C constant GDK_KEY_Yellow.
const KEY_Yellow = 269025189

// KEY_Ygrave is a representation of the C constant GDK_KEY_Ygrave.
const KEY_Ygrave = 16785138

// KEY_Yhook is a representation of the C constant GDK_KEY_Yhook.
const KEY_Yhook = 16785142

// KEY_Ytilde is a representation of the C constant GDK_KEY_Ytilde.
const KEY_Ytilde = 16785144

// KEY_Z is a representation of the C constant GDK_KEY_Z.
const KEY_Z = 90

// KEY_Zabovedot is a representation of the C constant GDK_KEY_Zabovedot.
const KEY_Zabovedot = 431

// KEY_Zacute is a representation of the C constant GDK_KEY_Zacute.
const KEY_Zacute = 428

// KEY_Zcaron is a representation of the C constant GDK_KEY_Zcaron.
const KEY_Zcaron = 430

// KEY_Zen_Koho is a representation of the C constant GDK_KEY_Zen_Koho.
const KEY_Zen_Koho = 65341

// KEY_Zenkaku is a representation of the C constant GDK_KEY_Zenkaku.
const KEY_Zenkaku = 65320

// KEY_Zenkaku_Hankaku is a representation of the C constant GDK_KEY_Zenkaku_Hankaku.
const KEY_Zenkaku_Hankaku = 65322

// KEY_ZoomIn is a representation of the C constant GDK_KEY_ZoomIn.
const KEY_ZoomIn = 269025163

// KEY_ZoomOut is a representation of the C constant GDK_KEY_ZoomOut.
const KEY_ZoomOut = 269025164

// KEY_Zstroke is a representation of the C constant GDK_KEY_Zstroke.
const KEY_Zstroke = 16777653

// KEY_a is a representation of the C constant GDK_KEY_a.
const KEY_a = 97

// KEY_aacute is a representation of the C constant GDK_KEY_aacute.
const KEY_aacute = 225

// KEY_abelowdot is a representation of the C constant GDK_KEY_abelowdot.
const KEY_abelowdot = 16785057

// KEY_abovedot is a representation of the C constant GDK_KEY_abovedot.
const KEY_abovedot = 511

// KEY_abreve is a representation of the C constant GDK_KEY_abreve.
const KEY_abreve = 483

// KEY_abreveacute is a representation of the C constant GDK_KEY_abreveacute.
const KEY_abreveacute = 16785071

// KEY_abrevebelowdot is a representation of the C constant GDK_KEY_abrevebelowdot.
const KEY_abrevebelowdot = 16785079

// KEY_abrevegrave is a representation of the C constant GDK_KEY_abrevegrave.
const KEY_abrevegrave = 16785073

// KEY_abrevehook is a representation of the C constant GDK_KEY_abrevehook.
const KEY_abrevehook = 16785075

// KEY_abrevetilde is a representation of the C constant GDK_KEY_abrevetilde.
const KEY_abrevetilde = 16785077

// KEY_acircumflex is a representation of the C constant GDK_KEY_acircumflex.
const KEY_acircumflex = 226

// KEY_acircumflexacute is a representation of the C constant GDK_KEY_acircumflexacute.
const KEY_acircumflexacute = 16785061

// KEY_acircumflexbelowdot is a representation of the C constant GDK_KEY_acircumflexbelowdot.
const KEY_acircumflexbelowdot = 16785069

// KEY_acircumflexgrave is a representation of the C constant GDK_KEY_acircumflexgrave.
const KEY_acircumflexgrave = 16785063

// KEY_acircumflexhook is a representation of the C constant GDK_KEY_acircumflexhook.
const KEY_acircumflexhook = 16785065

// KEY_acircumflextilde is a representation of the C constant GDK_KEY_acircumflextilde.
const KEY_acircumflextilde = 16785067

// KEY_acute is a representation of the C constant GDK_KEY_acute.
const KEY_acute = 180

// KEY_adiaeresis is a representation of the C constant GDK_KEY_adiaeresis.
const KEY_adiaeresis = 228

// KEY_ae is a representation of the C constant GDK_KEY_ae.
const KEY_ae = 230

// KEY_agrave is a representation of the C constant GDK_KEY_agrave.
const KEY_agrave = 224

// KEY_ahook is a representation of the C constant GDK_KEY_ahook.
const KEY_ahook = 16785059

// KEY_amacron is a representation of the C constant GDK_KEY_amacron.
const KEY_amacron = 992

// KEY_ampersand is a representation of the C constant GDK_KEY_ampersand.
const KEY_ampersand = 38

// KEY_aogonek is a representation of the C constant GDK_KEY_aogonek.
const KEY_aogonek = 433

// KEY_apostrophe is a representation of the C constant GDK_KEY_apostrophe.
const KEY_apostrophe = 39

// KEY_approxeq is a representation of the C constant GDK_KEY_approxeq.
const KEY_approxeq = 16785992

// KEY_approximate is a representation of the C constant GDK_KEY_approximate.
const KEY_approximate = 2248

// KEY_aring is a representation of the C constant GDK_KEY_aring.
const KEY_aring = 229

// KEY_asciicircum is a representation of the C constant GDK_KEY_asciicircum.
const KEY_asciicircum = 94

// KEY_asciitilde is a representation of the C constant GDK_KEY_asciitilde.
const KEY_asciitilde = 126

// KEY_asterisk is a representation of the C constant GDK_KEY_asterisk.
const KEY_asterisk = 42

// KEY_at is a representation of the C constant GDK_KEY_at.
const KEY_at = 64

// KEY_atilde is a representation of the C constant GDK_KEY_atilde.
const KEY_atilde = 227

// KEY_b is a representation of the C constant GDK_KEY_b.
const KEY_b = 98

// KEY_babovedot is a representation of the C constant GDK_KEY_babovedot.
const KEY_babovedot = 16784899

// KEY_backslash is a representation of the C constant GDK_KEY_backslash.
const KEY_backslash = 92

// KEY_ballotcross is a representation of the C constant GDK_KEY_ballotcross.
const KEY_ballotcross = 2804

// KEY_bar is a representation of the C constant GDK_KEY_bar.
const KEY_bar = 124

// KEY_because is a representation of the C constant GDK_KEY_because.
const KEY_because = 16785973

// KEY_blank is a representation of the C constant GDK_KEY_blank.
const KEY_blank = 2527

// KEY_botintegral is a representation of the C constant GDK_KEY_botintegral.
const KEY_botintegral = 2213

// KEY_botleftparens is a representation of the C constant GDK_KEY_botleftparens.
const KEY_botleftparens = 2220

// KEY_botleftsqbracket is a representation of the C constant GDK_KEY_botleftsqbracket.
const KEY_botleftsqbracket = 2216

// KEY_botleftsummation is a representation of the C constant GDK_KEY_botleftsummation.
const KEY_botleftsummation = 2226

// KEY_botrightparens is a representation of the C constant GDK_KEY_botrightparens.
const KEY_botrightparens = 2222

// KEY_botrightsqbracket is a representation of the C constant GDK_KEY_botrightsqbracket.
const KEY_botrightsqbracket = 2218

// KEY_botrightsummation is a representation of the C constant GDK_KEY_botrightsummation.
const KEY_botrightsummation = 2230

// KEY_bott is a representation of the C constant GDK_KEY_bott.
const KEY_bott = 2550

// KEY_botvertsummationconnector is a representation of the C constant GDK_KEY_botvertsummationconnector.
const KEY_botvertsummationconnector = 2228

// KEY_braceleft is a representation of the C constant GDK_KEY_braceleft.
const KEY_braceleft = 123

// KEY_braceright is a representation of the C constant GDK_KEY_braceright.
const KEY_braceright = 125

// KEY_bracketleft is a representation of the C constant GDK_KEY_bracketleft.
const KEY_bracketleft = 91

// KEY_bracketright is a representation of the C constant GDK_KEY_bracketright.
const KEY_bracketright = 93

// KEY_braille_blank is a representation of the C constant GDK_KEY_braille_blank.
const KEY_braille_blank = 16787456

// KEY_braille_dot_1 is a representation of the C constant GDK_KEY_braille_dot_1.
const KEY_braille_dot_1 = 65521

// KEY_braille_dot_10 is a representation of the C constant GDK_KEY_braille_dot_10.
const KEY_braille_dot_10 = 65530

// KEY_braille_dot_2 is a representation of the C constant GDK_KEY_braille_dot_2.
const KEY_braille_dot_2 = 65522

// KEY_braille_dot_3 is a representation of the C constant GDK_KEY_braille_dot_3.
const KEY_braille_dot_3 = 65523

// KEY_braille_dot_4 is a representation of the C constant GDK_KEY_braille_dot_4.
const KEY_braille_dot_4 = 65524

// KEY_braille_dot_5 is a representation of the C constant GDK_KEY_braille_dot_5.
const KEY_braille_dot_5 = 65525

// KEY_braille_dot_6 is a representation of the C constant GDK_KEY_braille_dot_6.
const KEY_braille_dot_6 = 65526

// KEY_braille_dot_7 is a representation of the C constant GDK_KEY_braille_dot_7.
const KEY_braille_dot_7 = 65527

// KEY_braille_dot_8 is a representation of the C constant GDK_KEY_braille_dot_8.
const KEY_braille_dot_8 = 65528

// KEY_braille_dot_9 is a representation of the C constant GDK_KEY_braille_dot_9.
const KEY_braille_dot_9 = 65529

// KEY_braille_dots_1 is a representation of the C constant GDK_KEY_braille_dots_1.
const KEY_braille_dots_1 = 16787457

// KEY_braille_dots_12 is a representation of the C constant GDK_KEY_braille_dots_12.
const KEY_braille_dots_12 = 16787459

// KEY_braille_dots_123 is a representation of the C constant GDK_KEY_braille_dots_123.
const KEY_braille_dots_123 = 16787463

// KEY_braille_dots_1234 is a representation of the C constant GDK_KEY_braille_dots_1234.
const KEY_braille_dots_1234 = 16787471

// KEY_braille_dots_12345 is a representation of the C constant GDK_KEY_braille_dots_12345.
const KEY_braille_dots_12345 = 16787487

// KEY_braille_dots_123456 is a representation of the C constant GDK_KEY_braille_dots_123456.
const KEY_braille_dots_123456 = 16787519

// KEY_braille_dots_1234567 is a representation of the C constant GDK_KEY_braille_dots_1234567.
const KEY_braille_dots_1234567 = 16787583

// KEY_braille_dots_12345678 is a representation of the C constant GDK_KEY_braille_dots_12345678.
const KEY_braille_dots_12345678 = 16787711

// KEY_braille_dots_1234568 is a representation of the C constant GDK_KEY_braille_dots_1234568.
const KEY_braille_dots_1234568 = 16787647

// KEY_braille_dots_123457 is a representation of the C constant GDK_KEY_braille_dots_123457.
const KEY_braille_dots_123457 = 16787551

// KEY_braille_dots_1234578 is a representation of the C constant GDK_KEY_braille_dots_1234578.
const KEY_braille_dots_1234578 = 16787679

// KEY_braille_dots_123458 is a representation of the C constant GDK_KEY_braille_dots_123458.
const KEY_braille_dots_123458 = 16787615

// KEY_braille_dots_12346 is a representation of the C constant GDK_KEY_braille_dots_12346.
const KEY_braille_dots_12346 = 16787503

// KEY_braille_dots_123467 is a representation of the C constant GDK_KEY_braille_dots_123467.
const KEY_braille_dots_123467 = 16787567

// KEY_braille_dots_1234678 is a representation of the C constant GDK_KEY_braille_dots_1234678.
const KEY_braille_dots_1234678 = 16787695

// KEY_braille_dots_123468 is a representation of the C constant GDK_KEY_braille_dots_123468.
const KEY_braille_dots_123468 = 16787631

// KEY_braille_dots_12347 is a representation of the C constant GDK_KEY_braille_dots_12347.
const KEY_braille_dots_12347 = 16787535

// KEY_braille_dots_123478 is a representation of the C constant GDK_KEY_braille_dots_123478.
const KEY_braille_dots_123478 = 16787663

// KEY_braille_dots_12348 is a representation of the C constant GDK_KEY_braille_dots_12348.
const KEY_braille_dots_12348 = 16787599

// KEY_braille_dots_1235 is a representation of the C constant GDK_KEY_braille_dots_1235.
const KEY_braille_dots_1235 = 16787479

// KEY_braille_dots_12356 is a representation of the C constant GDK_KEY_braille_dots_12356.
const KEY_braille_dots_12356 = 16787511

// KEY_braille_dots_123567 is a representation of the C constant GDK_KEY_braille_dots_123567.
const KEY_braille_dots_123567 = 16787575

// KEY_braille_dots_1235678 is a representation of the C constant GDK_KEY_braille_dots_1235678.
const KEY_braille_dots_1235678 = 16787703

// KEY_braille_dots_123568 is a representation of the C constant GDK_KEY_braille_dots_123568.
const KEY_braille_dots_123568 = 16787639

// KEY_braille_dots_12357 is a representation of the C constant GDK_KEY_braille_dots_12357.
const KEY_braille_dots_12357 = 16787543

// KEY_braille_dots_123578 is a representation of the C constant GDK_KEY_braille_dots_123578.
const KEY_braille_dots_123578 = 16787671

// KEY_braille_dots_12358 is a representation of the C constant GDK_KEY_braille_dots_12358.
const KEY_braille_dots_12358 = 16787607

// KEY_braille_dots_1236 is a representation of the C constant GDK_KEY_braille_dots_1236.
const KEY_braille_dots_1236 = 16787495

// KEY_braille_dots_12367 is a representation of the C constant GDK_KEY_braille_dots_12367.
const KEY_braille_dots_12367 = 16787559

// KEY_braille_dots_123678 is a representation of the C constant GDK_KEY_braille_dots_123678.
const KEY_braille_dots_123678 = 16787687

// KEY_braille_dots_12368 is a representation of the C constant GDK_KEY_braille_dots_12368.
const KEY_braille_dots_12368 = 16787623

// KEY_braille_dots_1237 is a representation of the C constant GDK_KEY_braille_dots_1237.
const KEY_braille_dots_1237 = 16787527

// KEY_braille_dots_12378 is a representation of the C constant GDK_KEY_braille_dots_12378.
const KEY_braille_dots_12378 = 16787655

// KEY_braille_dots_1238 is a representation of the C constant GDK_KEY_braille_dots_1238.
const KEY_braille_dots_1238 = 16787591

// KEY_braille_dots_124 is a representation of the C constant GDK_KEY_braille_dots_124.
const KEY_braille_dots_124 = 16787467

// KEY_braille_dots_1245 is a representation of the C constant GDK_KEY_braille_dots_1245.
const KEY_braille_dots_1245 = 16787483

// KEY_braille_dots_12456 is a representation of the C constant GDK_KEY_braille_dots_12456.
const KEY_braille_dots_12456 = 16787515

// KEY_braille_dots_124567 is a representation of the C constant GDK_KEY_braille_dots_124567.
const KEY_braille_dots_124567 = 16787579

// KEY_braille_dots_1245678 is a representation of the C constant GDK_KEY_braille_dots_1245678.
const KEY_braille_dots_1245678 = 16787707

// KEY_braille_dots_124568 is a representation of the C constant GDK_KEY_braille_dots_124568.
const KEY_braille_dots_124568 = 16787643

// KEY_braille_dots_12457 is a representation of the C constant GDK_KEY_braille_dots_12457.
const KEY_braille_dots_12457 = 16787547

// KEY_braille_dots_124578 is a representation of the C constant GDK_KEY_braille_dots_124578.
const KEY_braille_dots_124578 = 16787675

// KEY_braille_dots_12458 is a representation of the C constant GDK_KEY_braille_dots_12458.
const KEY_braille_dots_12458 = 16787611

// KEY_braille_dots_1246 is a representation of the C constant GDK_KEY_braille_dots_1246.
const KEY_braille_dots_1246 = 16787499

// KEY_braille_dots_12467 is a representation of the C constant GDK_KEY_braille_dots_12467.
const KEY_braille_dots_12467 = 16787563

// KEY_braille_dots_124678 is a representation of the C constant GDK_KEY_braille_dots_124678.
const KEY_braille_dots_124678 = 16787691

// KEY_braille_dots_12468 is a representation of the C constant GDK_KEY_braille_dots_12468.
const KEY_braille_dots_12468 = 16787627

// KEY_braille_dots_1247 is a representation of the C constant GDK_KEY_braille_dots_1247.
const KEY_braille_dots_1247 = 16787531

// KEY_braille_dots_12478 is a representation of the C constant GDK_KEY_braille_dots_12478.
const KEY_braille_dots_12478 = 16787659

// KEY_braille_dots_1248 is a representation of the C constant GDK_KEY_braille_dots_1248.
const KEY_braille_dots_1248 = 16787595

// KEY_braille_dots_125 is a representation of the C constant GDK_KEY_braille_dots_125.
const KEY_braille_dots_125 = 16787475

// KEY_braille_dots_1256 is a representation of the C constant GDK_KEY_braille_dots_1256.
const KEY_braille_dots_1256 = 16787507

// KEY_braille_dots_12567 is a representation of the C constant GDK_KEY_braille_dots_12567.
const KEY_braille_dots_12567 = 16787571

// KEY_braille_dots_125678 is a representation of the C constant GDK_KEY_braille_dots_125678.
const KEY_braille_dots_125678 = 16787699

// KEY_braille_dots_12568 is a representation of the C constant GDK_KEY_braille_dots_12568.
const KEY_braille_dots_12568 = 16787635

// KEY_braille_dots_1257 is a representation of the C constant GDK_KEY_braille_dots_1257.
const KEY_braille_dots_1257 = 16787539

// KEY_braille_dots_12578 is a representation of the C constant GDK_KEY_braille_dots_12578.
const KEY_braille_dots_12578 = 16787667

// KEY_braille_dots_1258 is a representation of the C constant GDK_KEY_braille_dots_1258.
const KEY_braille_dots_1258 = 16787603

// KEY_braille_dots_126 is a representation of the C constant GDK_KEY_braille_dots_126.
const KEY_braille_dots_126 = 16787491

// KEY_braille_dots_1267 is a representation of the C constant GDK_KEY_braille_dots_1267.
const KEY_braille_dots_1267 = 16787555

// KEY_braille_dots_12678 is a representation of the C constant GDK_KEY_braille_dots_12678.
const KEY_braille_dots_12678 = 16787683

// KEY_braille_dots_1268 is a representation of the C constant GDK_KEY_braille_dots_1268.
const KEY_braille_dots_1268 = 16787619

// KEY_braille_dots_127 is a representation of the C constant GDK_KEY_braille_dots_127.
const KEY_braille_dots_127 = 16787523

// KEY_braille_dots_1278 is a representation of the C constant GDK_KEY_braille_dots_1278.
const KEY_braille_dots_1278 = 16787651

// KEY_braille_dots_128 is a representation of the C constant GDK_KEY_braille_dots_128.
const KEY_braille_dots_128 = 16787587

// KEY_braille_dots_13 is a representation of the C constant GDK_KEY_braille_dots_13.
const KEY_braille_dots_13 = 16787461

// KEY_braille_dots_134 is a representation of the C constant GDK_KEY_braille_dots_134.
const KEY_braille_dots_134 = 16787469

// KEY_braille_dots_1345 is a representation of the C constant GDK_KEY_braille_dots_1345.
const KEY_braille_dots_1345 = 16787485

// KEY_braille_dots_13456 is a representation of the C constant GDK_KEY_braille_dots_13456.
const KEY_braille_dots_13456 = 16787517

// KEY_braille_dots_134567 is a representation of the C constant GDK_KEY_braille_dots_134567.
const KEY_braille_dots_134567 = 16787581

// KEY_braille_dots_1345678 is a representation of the C constant GDK_KEY_braille_dots_1345678.
const KEY_braille_dots_1345678 = 16787709

// KEY_braille_dots_134568 is a representation of the C constant GDK_KEY_braille_dots_134568.
const KEY_braille_dots_134568 = 16787645

// KEY_braille_dots_13457 is a representation of the C constant GDK_KEY_braille_dots_13457.
const KEY_braille_dots_13457 = 16787549

// KEY_braille_dots_134578 is a representation of the C constant GDK_KEY_braille_dots_134578.
const KEY_braille_dots_134578 = 16787677

// KEY_braille_dots_13458 is a representation of the C constant GDK_KEY_braille_dots_13458.
const KEY_braille_dots_13458 = 16787613

// KEY_braille_dots_1346 is a representation of the C constant GDK_KEY_braille_dots_1346.
const KEY_braille_dots_1346 = 16787501

// KEY_braille_dots_13467 is a representation of the C constant GDK_KEY_braille_dots_13467.
const KEY_braille_dots_13467 = 16787565

// KEY_braille_dots_134678 is a representation of the C constant GDK_KEY_braille_dots_134678.
const KEY_braille_dots_134678 = 16787693

// KEY_braille_dots_13468 is a representation of the C constant GDK_KEY_braille_dots_13468.
const KEY_braille_dots_13468 = 16787629

// KEY_braille_dots_1347 is a representation of the C constant GDK_KEY_braille_dots_1347.
const KEY_braille_dots_1347 = 16787533

// KEY_braille_dots_13478 is a representation of the C constant GDK_KEY_braille_dots_13478.
const KEY_braille_dots_13478 = 16787661

// KEY_braille_dots_1348 is a representation of the C constant GDK_KEY_braille_dots_1348.
const KEY_braille_dots_1348 = 16787597

// KEY_braille_dots_135 is a representation of the C constant GDK_KEY_braille_dots_135.
const KEY_braille_dots_135 = 16787477

// KEY_braille_dots_1356 is a representation of the C constant GDK_KEY_braille_dots_1356.
const KEY_braille_dots_1356 = 16787509

// KEY_braille_dots_13567 is a representation of the C constant GDK_KEY_braille_dots_13567.
const KEY_braille_dots_13567 = 16787573

// KEY_braille_dots_135678 is a representation of the C constant GDK_KEY_braille_dots_135678.
const KEY_braille_dots_135678 = 16787701

// KEY_braille_dots_13568 is a representation of the C constant GDK_KEY_braille_dots_13568.
const KEY_braille_dots_13568 = 16787637

// KEY_braille_dots_1357 is a representation of the C constant GDK_KEY_braille_dots_1357.
const KEY_braille_dots_1357 = 16787541

// KEY_braille_dots_13578 is a representation of the C constant GDK_KEY_braille_dots_13578.
const KEY_braille_dots_13578 = 16787669

// KEY_braille_dots_1358 is a representation of the C constant GDK_KEY_braille_dots_1358.
const KEY_braille_dots_1358 = 16787605

// KEY_braille_dots_136 is a representation of the C constant GDK_KEY_braille_dots_136.
const KEY_braille_dots_136 = 16787493

// KEY_braille_dots_1367 is a representation of the C constant GDK_KEY_braille_dots_1367.
const KEY_braille_dots_1367 = 16787557

// KEY_braille_dots_13678 is a representation of the C constant GDK_KEY_braille_dots_13678.
const KEY_braille_dots_13678 = 16787685

// KEY_braille_dots_1368 is a representation of the C constant GDK_KEY_braille_dots_1368.
const KEY_braille_dots_1368 = 16787621

// KEY_braille_dots_137 is a representation of the C constant GDK_KEY_braille_dots_137.
const KEY_braille_dots_137 = 16787525

// KEY_braille_dots_1378 is a representation of the C constant GDK_KEY_braille_dots_1378.
const KEY_braille_dots_1378 = 16787653

// KEY_braille_dots_138 is a representation of the C constant GDK_KEY_braille_dots_138.
const KEY_braille_dots_138 = 16787589

// KEY_braille_dots_14 is a representation of the C constant GDK_KEY_braille_dots_14.
const KEY_braille_dots_14 = 16787465

// KEY_braille_dots_145 is a representation of the C constant GDK_KEY_braille_dots_145.
const KEY_braille_dots_145 = 16787481

// KEY_braille_dots_1456 is a representation of the C constant GDK_KEY_braille_dots_1456.
const KEY_braille_dots_1456 = 16787513

// KEY_braille_dots_14567 is a representation of the C constant GDK_KEY_braille_dots_14567.
const KEY_braille_dots_14567 = 16787577

// KEY_braille_dots_145678 is a representation of the C constant GDK_KEY_braille_dots_145678.
const KEY_braille_dots_145678 = 16787705

// KEY_braille_dots_14568 is a representation of the C constant GDK_KEY_braille_dots_14568.
const KEY_braille_dots_14568 = 16787641

// KEY_braille_dots_1457 is a representation of the C constant GDK_KEY_braille_dots_1457.
const KEY_braille_dots_1457 = 16787545

// KEY_braille_dots_14578 is a representation of the C constant GDK_KEY_braille_dots_14578.
const KEY_braille_dots_14578 = 16787673

// KEY_braille_dots_1458 is a representation of the C constant GDK_KEY_braille_dots_1458.
const KEY_braille_dots_1458 = 16787609

// KEY_braille_dots_146 is a representation of the C constant GDK_KEY_braille_dots_146.
const KEY_braille_dots_146 = 16787497

// KEY_braille_dots_1467 is a representation of the C constant GDK_KEY_braille_dots_1467.
const KEY_braille_dots_1467 = 16787561

// KEY_braille_dots_14678 is a representation of the C constant GDK_KEY_braille_dots_14678.
const KEY_braille_dots_14678 = 16787689

// KEY_braille_dots_1468 is a representation of the C constant GDK_KEY_braille_dots_1468.
const KEY_braille_dots_1468 = 16787625

// KEY_braille_dots_147 is a representation of the C constant GDK_KEY_braille_dots_147.
const KEY_braille_dots_147 = 16787529

// KEY_braille_dots_1478 is a representation of the C constant GDK_KEY_braille_dots_1478.
const KEY_braille_dots_1478 = 16787657

// KEY_braille_dots_148 is a representation of the C constant GDK_KEY_braille_dots_148.
const KEY_braille_dots_148 = 16787593

// KEY_braille_dots_15 is a representation of the C constant GDK_KEY_braille_dots_15.
const KEY_braille_dots_15 = 16787473

// KEY_braille_dots_156 is a representation of the C constant GDK_KEY_braille_dots_156.
const KEY_braille_dots_156 = 16787505

// KEY_braille_dots_1567 is a representation of the C constant GDK_KEY_braille_dots_1567.
const KEY_braille_dots_1567 = 16787569

// KEY_braille_dots_15678 is a representation of the C constant GDK_KEY_braille_dots_15678.
const KEY_braille_dots_15678 = 16787697

// KEY_braille_dots_1568 is a representation of the C constant GDK_KEY_braille_dots_1568.
const KEY_braille_dots_1568 = 16787633

// KEY_braille_dots_157 is a representation of the C constant GDK_KEY_braille_dots_157.
const KEY_braille_dots_157 = 16787537

// KEY_braille_dots_1578 is a representation of the C constant GDK_KEY_braille_dots_1578.
const KEY_braille_dots_1578 = 16787665

// KEY_braille_dots_158 is a representation of the C constant GDK_KEY_braille_dots_158.
const KEY_braille_dots_158 = 16787601

// KEY_braille_dots_16 is a representation of the C constant GDK_KEY_braille_dots_16.
const KEY_braille_dots_16 = 16787489

// KEY_braille_dots_167 is a representation of the C constant GDK_KEY_braille_dots_167.
const KEY_braille_dots_167 = 16787553

// KEY_braille_dots_1678 is a representation of the C constant GDK_KEY_braille_dots_1678.
const KEY_braille_dots_1678 = 16787681

// KEY_braille_dots_168 is a representation of the C constant GDK_KEY_braille_dots_168.
const KEY_braille_dots_168 = 16787617

// KEY_braille_dots_17 is a representation of the C constant GDK_KEY_braille_dots_17.
const KEY_braille_dots_17 = 16787521

// KEY_braille_dots_178 is a representation of the C constant GDK_KEY_braille_dots_178.
const KEY_braille_dots_178 = 16787649

// KEY_braille_dots_18 is a representation of the C constant GDK_KEY_braille_dots_18.
const KEY_braille_dots_18 = 16787585

// KEY_braille_dots_2 is a representation of the C constant GDK_KEY_braille_dots_2.
const KEY_braille_dots_2 = 16787458

// KEY_braille_dots_23 is a representation of the C constant GDK_KEY_braille_dots_23.
const KEY_braille_dots_23 = 16787462

// KEY_braille_dots_234 is a representation of the C constant GDK_KEY_braille_dots_234.
const KEY_braille_dots_234 = 16787470

// KEY_braille_dots_2345 is a representation of the C constant GDK_KEY_braille_dots_2345.
const KEY_braille_dots_2345 = 16787486

// KEY_braille_dots_23456 is a representation of the C constant GDK_KEY_braille_dots_23456.
const KEY_braille_dots_23456 = 16787518

// KEY_braille_dots_234567 is a representation of the C constant GDK_KEY_braille_dots_234567.
const KEY_braille_dots_234567 = 16787582

// KEY_braille_dots_2345678 is a representation of the C constant GDK_KEY_braille_dots_2345678.
const KEY_braille_dots_2345678 = 16787710

// KEY_braille_dots_234568 is a representation of the C constant GDK_KEY_braille_dots_234568.
const KEY_braille_dots_234568 = 16787646

// KEY_braille_dots_23457 is a representation of the C constant GDK_KEY_braille_dots_23457.
const KEY_braille_dots_23457 = 16787550

// KEY_braille_dots_234578 is a representation of the C constant GDK_KEY_braille_dots_234578.
const KEY_braille_dots_234578 = 16787678

// KEY_braille_dots_23458 is a representation of the C constant GDK_KEY_braille_dots_23458.
const KEY_braille_dots_23458 = 16787614

// KEY_braille_dots_2346 is a representation of the C constant GDK_KEY_braille_dots_2346.
const KEY_braille_dots_2346 = 16787502

// KEY_braille_dots_23467 is a representation of the C constant GDK_KEY_braille_dots_23467.
const KEY_braille_dots_23467 = 16787566

// KEY_braille_dots_234678 is a representation of the C constant GDK_KEY_braille_dots_234678.
const KEY_braille_dots_234678 = 16787694

// KEY_braille_dots_23468 is a representation of the C constant GDK_KEY_braille_dots_23468.
const KEY_braille_dots_23468 = 16787630

// KEY_braille_dots_2347 is a representation of the C constant GDK_KEY_braille_dots_2347.
const KEY_braille_dots_2347 = 16787534

// KEY_braille_dots_23478 is a representation of the C constant GDK_KEY_braille_dots_23478.
const KEY_braille_dots_23478 = 16787662

// KEY_braille_dots_2348 is a representation of the C constant GDK_KEY_braille_dots_2348.
const KEY_braille_dots_2348 = 16787598

// KEY_braille_dots_235 is a representation of the C constant GDK_KEY_braille_dots_235.
const KEY_braille_dots_235 = 16787478

// KEY_braille_dots_2356 is a representation of the C constant GDK_KEY_braille_dots_2356.
const KEY_braille_dots_2356 = 16787510

// KEY_braille_dots_23567 is a representation of the C constant GDK_KEY_braille_dots_23567.
const KEY_braille_dots_23567 = 16787574

// KEY_braille_dots_235678 is a representation of the C constant GDK_KEY_braille_dots_235678.
const KEY_braille_dots_235678 = 16787702

// KEY_braille_dots_23568 is a representation of the C constant GDK_KEY_braille_dots_23568.
const KEY_braille_dots_23568 = 16787638

// KEY_braille_dots_2357 is a representation of the C constant GDK_KEY_braille_dots_2357.
const KEY_braille_dots_2357 = 16787542

// KEY_braille_dots_23578 is a representation of the C constant GDK_KEY_braille_dots_23578.
const KEY_braille_dots_23578 = 16787670

// KEY_braille_dots_2358 is a representation of the C constant GDK_KEY_braille_dots_2358.
const KEY_braille_dots_2358 = 16787606

// KEY_braille_dots_236 is a representation of the C constant GDK_KEY_braille_dots_236.
const KEY_braille_dots_236 = 16787494

// KEY_braille_dots_2367 is a representation of the C constant GDK_KEY_braille_dots_2367.
const KEY_braille_dots_2367 = 16787558

// KEY_braille_dots_23678 is a representation of the C constant GDK_KEY_braille_dots_23678.
const KEY_braille_dots_23678 = 16787686

// KEY_braille_dots_2368 is a representation of the C constant GDK_KEY_braille_dots_2368.
const KEY_braille_dots_2368 = 16787622

// KEY_braille_dots_237 is a representation of the C constant GDK_KEY_braille_dots_237.
const KEY_braille_dots_237 = 16787526

// KEY_braille_dots_2378 is a representation of the C constant GDK_KEY_braille_dots_2378.
const KEY_braille_dots_2378 = 16787654

// KEY_braille_dots_238 is a representation of the C constant GDK_KEY_braille_dots_238.
const KEY_braille_dots_238 = 16787590

// KEY_braille_dots_24 is a representation of the C constant GDK_KEY_braille_dots_24.
const KEY_braille_dots_24 = 16787466

// KEY_braille_dots_245 is a representation of the C constant GDK_KEY_braille_dots_245.
const KEY_braille_dots_245 = 16787482

// KEY_braille_dots_2456 is a representation of the C constant GDK_KEY_braille_dots_2456.
const KEY_braille_dots_2456 = 16787514

// KEY_braille_dots_24567 is a representation of the C constant GDK_KEY_braille_dots_24567.
const KEY_braille_dots_24567 = 16787578

// KEY_braille_dots_245678 is a representation of the C constant GDK_KEY_braille_dots_245678.
const KEY_braille_dots_245678 = 16787706

// KEY_braille_dots_24568 is a representation of the C constant GDK_KEY_braille_dots_24568.
const KEY_braille_dots_24568 = 16787642

// KEY_braille_dots_2457 is a representation of the C constant GDK_KEY_braille_dots_2457.
const KEY_braille_dots_2457 = 16787546

// KEY_braille_dots_24578 is a representation of the C constant GDK_KEY_braille_dots_24578.
const KEY_braille_dots_24578 = 16787674

// KEY_braille_dots_2458 is a representation of the C constant GDK_KEY_braille_dots_2458.
const KEY_braille_dots_2458 = 16787610

// KEY_braille_dots_246 is a representation of the C constant GDK_KEY_braille_dots_246.
const KEY_braille_dots_246 = 16787498

// KEY_braille_dots_2467 is a representation of the C constant GDK_KEY_braille_dots_2467.
const KEY_braille_dots_2467 = 16787562

// KEY_braille_dots_24678 is a representation of the C constant GDK_KEY_braille_dots_24678.
const KEY_braille_dots_24678 = 16787690

// KEY_braille_dots_2468 is a representation of the C constant GDK_KEY_braille_dots_2468.
const KEY_braille_dots_2468 = 16787626

// KEY_braille_dots_247 is a representation of the C constant GDK_KEY_braille_dots_247.
const KEY_braille_dots_247 = 16787530

// KEY_braille_dots_2478 is a representation of the C constant GDK_KEY_braille_dots_2478.
const KEY_braille_dots_2478 = 16787658

// KEY_braille_dots_248 is a representation of the C constant GDK_KEY_braille_dots_248.
const KEY_braille_dots_248 = 16787594

// KEY_braille_dots_25 is a representation of the C constant GDK_KEY_braille_dots_25.
const KEY_braille_dots_25 = 16787474

// KEY_braille_dots_256 is a representation of the C constant GDK_KEY_braille_dots_256.
const KEY_braille_dots_256 = 16787506

// KEY_braille_dots_2567 is a representation of the C constant GDK_KEY_braille_dots_2567.
const KEY_braille_dots_2567 = 16787570

// KEY_braille_dots_25678 is a representation of the C constant GDK_KEY_braille_dots_25678.
const KEY_braille_dots_25678 = 16787698

// KEY_braille_dots_2568 is a representation of the C constant GDK_KEY_braille_dots_2568.
const KEY_braille_dots_2568 = 16787634

// KEY_braille_dots_257 is a representation of the C constant GDK_KEY_braille_dots_257.
const KEY_braille_dots_257 = 16787538

// KEY_braille_dots_2578 is a representation of the C constant GDK_KEY_braille_dots_2578.
const KEY_braille_dots_2578 = 16787666

// KEY_braille_dots_258 is a representation of the C constant GDK_KEY_braille_dots_258.
const KEY_braille_dots_258 = 16787602

// KEY_braille_dots_26 is a representation of the C constant GDK_KEY_braille_dots_26.
const KEY_braille_dots_26 = 16787490

// KEY_braille_dots_267 is a representation of the C constant GDK_KEY_braille_dots_267.
const KEY_braille_dots_267 = 16787554

// KEY_braille_dots_2678 is a representation of the C constant GDK_KEY_braille_dots_2678.
const KEY_braille_dots_2678 = 16787682

// KEY_braille_dots_268 is a representation of the C constant GDK_KEY_braille_dots_268.
const KEY_braille_dots_268 = 16787618

// KEY_braille_dots_27 is a representation of the C constant GDK_KEY_braille_dots_27.
const KEY_braille_dots_27 = 16787522

// KEY_braille_dots_278 is a representation of the C constant GDK_KEY_braille_dots_278.
const KEY_braille_dots_278 = 16787650

// KEY_braille_dots_28 is a representation of the C constant GDK_KEY_braille_dots_28.
const KEY_braille_dots_28 = 16787586

// KEY_braille_dots_3 is a representation of the C constant GDK_KEY_braille_dots_3.
const KEY_braille_dots_3 = 16787460

// KEY_braille_dots_34 is a representation of the C constant GDK_KEY_braille_dots_34.
const KEY_braille_dots_34 = 16787468

// KEY_braille_dots_345 is a representation of the C constant GDK_KEY_braille_dots_345.
const KEY_braille_dots_345 = 16787484

// KEY_braille_dots_3456 is a representation of the C constant GDK_KEY_braille_dots_3456.
const KEY_braille_dots_3456 = 16787516

// KEY_braille_dots_34567 is a representation of the C constant GDK_KEY_braille_dots_34567.
const KEY_braille_dots_34567 = 16787580

// KEY_braille_dots_345678 is a representation of the C constant GDK_KEY_braille_dots_345678.
const KEY_braille_dots_345678 = 16787708

// KEY_braille_dots_34568 is a representation of the C constant GDK_KEY_braille_dots_34568.
const KEY_braille_dots_34568 = 16787644

// KEY_braille_dots_3457 is a representation of the C constant GDK_KEY_braille_dots_3457.
const KEY_braille_dots_3457 = 16787548

// KEY_braille_dots_34578 is a representation of the C constant GDK_KEY_braille_dots_34578.
const KEY_braille_dots_34578 = 16787676

// KEY_braille_dots_3458 is a representation of the C constant GDK_KEY_braille_dots_3458.
const KEY_braille_dots_3458 = 16787612

// KEY_braille_dots_346 is a representation of the C constant GDK_KEY_braille_dots_346.
const KEY_braille_dots_346 = 16787500

// KEY_braille_dots_3467 is a representation of the C constant GDK_KEY_braille_dots_3467.
const KEY_braille_dots_3467 = 16787564

// KEY_braille_dots_34678 is a representation of the C constant GDK_KEY_braille_dots_34678.
const KEY_braille_dots_34678 = 16787692

// KEY_braille_dots_3468 is a representation of the C constant GDK_KEY_braille_dots_3468.
const KEY_braille_dots_3468 = 16787628

// KEY_braille_dots_347 is a representation of the C constant GDK_KEY_braille_dots_347.
const KEY_braille_dots_347 = 16787532

// KEY_braille_dots_3478 is a representation of the C constant GDK_KEY_braille_dots_3478.
const KEY_braille_dots_3478 = 16787660

// KEY_braille_dots_348 is a representation of the C constant GDK_KEY_braille_dots_348.
const KEY_braille_dots_348 = 16787596

// KEY_braille_dots_35 is a representation of the C constant GDK_KEY_braille_dots_35.
const KEY_braille_dots_35 = 16787476

// KEY_braille_dots_356 is a representation of the C constant GDK_KEY_braille_dots_356.
const KEY_braille_dots_356 = 16787508

// KEY_braille_dots_3567 is a representation of the C constant GDK_KEY_braille_dots_3567.
const KEY_braille_dots_3567 = 16787572

// KEY_braille_dots_35678 is a representation of the C constant GDK_KEY_braille_dots_35678.
const KEY_braille_dots_35678 = 16787700

// KEY_braille_dots_3568 is a representation of the C constant GDK_KEY_braille_dots_3568.
const KEY_braille_dots_3568 = 16787636

// KEY_braille_dots_357 is a representation of the C constant GDK_KEY_braille_dots_357.
const KEY_braille_dots_357 = 16787540

// KEY_braille_dots_3578 is a representation of the C constant GDK_KEY_braille_dots_3578.
const KEY_braille_dots_3578 = 16787668

// KEY_braille_dots_358 is a representation of the C constant GDK_KEY_braille_dots_358.
const KEY_braille_dots_358 = 16787604

// KEY_braille_dots_36 is a representation of the C constant GDK_KEY_braille_dots_36.
const KEY_braille_dots_36 = 16787492

// KEY_braille_dots_367 is a representation of the C constant GDK_KEY_braille_dots_367.
const KEY_braille_dots_367 = 16787556

// KEY_braille_dots_3678 is a representation of the C constant GDK_KEY_braille_dots_3678.
const KEY_braille_dots_3678 = 16787684

// KEY_braille_dots_368 is a representation of the C constant GDK_KEY_braille_dots_368.
const KEY_braille_dots_368 = 16787620

// KEY_braille_dots_37 is a representation of the C constant GDK_KEY_braille_dots_37.
const KEY_braille_dots_37 = 16787524

// KEY_braille_dots_378 is a representation of the C constant GDK_KEY_braille_dots_378.
const KEY_braille_dots_378 = 16787652

// KEY_braille_dots_38 is a representation of the C constant GDK_KEY_braille_dots_38.
const KEY_braille_dots_38 = 16787588

// KEY_braille_dots_4 is a representation of the C constant GDK_KEY_braille_dots_4.
const KEY_braille_dots_4 = 16787464

// KEY_braille_dots_45 is a representation of the C constant GDK_KEY_braille_dots_45.
const KEY_braille_dots_45 = 16787480

// KEY_braille_dots_456 is a representation of the C constant GDK_KEY_braille_dots_456.
const KEY_braille_dots_456 = 16787512

// KEY_braille_dots_4567 is a representation of the C constant GDK_KEY_braille_dots_4567.
const KEY_braille_dots_4567 = 16787576

// KEY_braille_dots_45678 is a representation of the C constant GDK_KEY_braille_dots_45678.
const KEY_braille_dots_45678 = 16787704

// KEY_braille_dots_4568 is a representation of the C constant GDK_KEY_braille_dots_4568.
const KEY_braille_dots_4568 = 16787640

// KEY_braille_dots_457 is a representation of the C constant GDK_KEY_braille_dots_457.
const KEY_braille_dots_457 = 16787544

// KEY_braille_dots_4578 is a representation of the C constant GDK_KEY_braille_dots_4578.
const KEY_braille_dots_4578 = 16787672

// KEY_braille_dots_458 is a representation of the C constant GDK_KEY_braille_dots_458.
const KEY_braille_dots_458 = 16787608

// KEY_braille_dots_46 is a representation of the C constant GDK_KEY_braille_dots_46.
const KEY_braille_dots_46 = 16787496

// KEY_braille_dots_467 is a representation of the C constant GDK_KEY_braille_dots_467.
const KEY_braille_dots_467 = 16787560

// KEY_braille_dots_4678 is a representation of the C constant GDK_KEY_braille_dots_4678.
const KEY_braille_dots_4678 = 16787688

// KEY_braille_dots_468 is a representation of the C constant GDK_KEY_braille_dots_468.
const KEY_braille_dots_468 = 16787624

// KEY_braille_dots_47 is a representation of the C constant GDK_KEY_braille_dots_47.
const KEY_braille_dots_47 = 16787528

// KEY_braille_dots_478 is a representation of the C constant GDK_KEY_braille_dots_478.
const KEY_braille_dots_478 = 16787656

// KEY_braille_dots_48 is a representation of the C constant GDK_KEY_braille_dots_48.
const KEY_braille_dots_48 = 16787592

// KEY_braille_dots_5 is a representation of the C constant GDK_KEY_braille_dots_5.
const KEY_braille_dots_5 = 16787472

// KEY_braille_dots_56 is a representation of the C constant GDK_KEY_braille_dots_56.
const KEY_braille_dots_56 = 16787504

// KEY_braille_dots_567 is a representation of the C constant GDK_KEY_braille_dots_567.
const KEY_braille_dots_567 = 16787568

// KEY_braille_dots_5678 is a representation of the C constant GDK_KEY_braille_dots_5678.
const KEY_braille_dots_5678 = 16787696

// KEY_braille_dots_568 is a representation of the C constant GDK_KEY_braille_dots_568.
const KEY_braille_dots_568 = 16787632

// KEY_braille_dots_57 is a representation of the C constant GDK_KEY_braille_dots_57.
const KEY_braille_dots_57 = 16787536

// KEY_braille_dots_578 is a representation of the C constant GDK_KEY_braille_dots_578.
const KEY_braille_dots_578 = 16787664

// KEY_braille_dots_58 is a representation of the C constant GDK_KEY_braille_dots_58.
const KEY_braille_dots_58 = 16787600

// KEY_braille_dots_6 is a representation of the C constant GDK_KEY_braille_dots_6.
const KEY_braille_dots_6 = 16787488

// KEY_braille_dots_67 is a representation of the C constant GDK_KEY_braille_dots_67.
const KEY_braille_dots_67 = 16787552

// KEY_braille_dots_678 is a representation of the C constant GDK_KEY_braille_dots_678.
const KEY_braille_dots_678 = 16787680

// KEY_braille_dots_68 is a representation of the C constant GDK_KEY_braille_dots_68.
const KEY_braille_dots_68 = 16787616

// KEY_braille_dots_7 is a representation of the C constant GDK_KEY_braille_dots_7.
const KEY_braille_dots_7 = 16787520

// KEY_braille_dots_78 is a representation of the C constant GDK_KEY_braille_dots_78.
const KEY_braille_dots_78 = 16787648

// KEY_braille_dots_8 is a representation of the C constant GDK_KEY_braille_dots_8.
const KEY_braille_dots_8 = 16787584

// KEY_breve is a representation of the C constant GDK_KEY_breve.
const KEY_breve = 418

// KEY_brokenbar is a representation of the C constant GDK_KEY_brokenbar.
const KEY_brokenbar = 166

// KEY_c is a representation of the C constant GDK_KEY_c.
const KEY_c = 99

// KEY_c_h is a representation of the C constant GDK_KEY_c_h.
const KEY_c_h = 65187

// KEY_cabovedot is a representation of the C constant GDK_KEY_cabovedot.
const KEY_cabovedot = 741

// KEY_cacute is a representation of the C constant GDK_KEY_cacute.
const KEY_cacute = 486

// KEY_careof is a representation of the C constant GDK_KEY_careof.
const KEY_careof = 2744

// KEY_caret is a representation of the C constant GDK_KEY_caret.
const KEY_caret = 2812

// KEY_caron is a representation of the C constant GDK_KEY_caron.
const KEY_caron = 439

// KEY_ccaron is a representation of the C constant GDK_KEY_ccaron.
const KEY_ccaron = 488

// KEY_ccedilla is a representation of the C constant GDK_KEY_ccedilla.
const KEY_ccedilla = 231

// KEY_ccircumflex is a representation of the C constant GDK_KEY_ccircumflex.
const KEY_ccircumflex = 742

// KEY_cedilla is a representation of the C constant GDK_KEY_cedilla.
const KEY_cedilla = 184

// KEY_cent is a representation of the C constant GDK_KEY_cent.
const KEY_cent = 162

// KEY_ch is a representation of the C constant GDK_KEY_ch.
const KEY_ch = 65184

// KEY_checkerboard is a representation of the C constant GDK_KEY_checkerboard.
const KEY_checkerboard = 2529

// KEY_checkmark is a representation of the C constant GDK_KEY_checkmark.
const KEY_checkmark = 2803

// KEY_circle is a representation of the C constant GDK_KEY_circle.
const KEY_circle = 3023

// KEY_club is a representation of the C constant GDK_KEY_club.
const KEY_club = 2796

// KEY_colon is a representation of the C constant GDK_KEY_colon.
const KEY_colon = 58

// KEY_comma is a representation of the C constant GDK_KEY_comma.
const KEY_comma = 44

// KEY_containsas is a representation of the C constant GDK_KEY_containsas.
const KEY_containsas = 16785931

// KEY_copyright is a representation of the C constant GDK_KEY_copyright.
const KEY_copyright = 169

// KEY_cr is a representation of the C constant GDK_KEY_cr.
const KEY_cr = 2532

// KEY_crossinglines is a representation of the C constant GDK_KEY_crossinglines.
const KEY_crossinglines = 2542

// KEY_cuberoot is a representation of the C constant GDK_KEY_cuberoot.
const KEY_cuberoot = 16785947

// KEY_currency is a representation of the C constant GDK_KEY_currency.
const KEY_currency = 164

// KEY_cursor is a representation of the C constant GDK_KEY_cursor.
const KEY_cursor = 2815

// KEY_d is a representation of the C constant GDK_KEY_d.
const KEY_d = 100

// KEY_dabovedot is a representation of the C constant GDK_KEY_dabovedot.
const KEY_dabovedot = 16784907

// KEY_dagger is a representation of the C constant GDK_KEY_dagger.
const KEY_dagger = 2801

// KEY_dcaron is a representation of the C constant GDK_KEY_dcaron.
const KEY_dcaron = 495

// KEY_dead_A is a representation of the C constant GDK_KEY_dead_A.
const KEY_dead_A = 65153

// KEY_dead_E is a representation of the C constant GDK_KEY_dead_E.
const KEY_dead_E = 65155

// KEY_dead_I is a representation of the C constant GDK_KEY_dead_I.
const KEY_dead_I = 65157

// KEY_dead_O is a representation of the C constant GDK_KEY_dead_O.
const KEY_dead_O = 65159

// KEY_dead_U is a representation of the C constant GDK_KEY_dead_U.
const KEY_dead_U = 65161

// KEY_dead_a is a representation of the C constant GDK_KEY_dead_a.
const KEY_dead_a = 65152

// KEY_dead_abovecomma is a representation of the C constant GDK_KEY_dead_abovecomma.
const KEY_dead_abovecomma = 65124

// KEY_dead_abovedot is a representation of the C constant GDK_KEY_dead_abovedot.
const KEY_dead_abovedot = 65110

// KEY_dead_abovereversedcomma is a representation of the C constant GDK_KEY_dead_abovereversedcomma.
const KEY_dead_abovereversedcomma = 65125

// KEY_dead_abovering is a representation of the C constant GDK_KEY_dead_abovering.
const KEY_dead_abovering = 65112

// KEY_dead_acute is a representation of the C constant GDK_KEY_dead_acute.
const KEY_dead_acute = 65105

// KEY_dead_belowbreve is a representation of the C constant GDK_KEY_dead_belowbreve.
const KEY_dead_belowbreve = 65131

// KEY_dead_belowcircumflex is a representation of the C constant GDK_KEY_dead_belowcircumflex.
const KEY_dead_belowcircumflex = 65129

// KEY_dead_belowcomma is a representation of the C constant GDK_KEY_dead_belowcomma.
const KEY_dead_belowcomma = 65134

// KEY_dead_belowdiaeresis is a representation of the C constant GDK_KEY_dead_belowdiaeresis.
const KEY_dead_belowdiaeresis = 65132

// KEY_dead_belowdot is a representation of the C constant GDK_KEY_dead_belowdot.
const KEY_dead_belowdot = 65120

// KEY_dead_belowmacron is a representation of the C constant GDK_KEY_dead_belowmacron.
const KEY_dead_belowmacron = 65128

// KEY_dead_belowring is a representation of the C constant GDK_KEY_dead_belowring.
const KEY_dead_belowring = 65127

// KEY_dead_belowtilde is a representation of the C constant GDK_KEY_dead_belowtilde.
const KEY_dead_belowtilde = 65130

// KEY_dead_breve is a representation of the C constant GDK_KEY_dead_breve.
const KEY_dead_breve = 65109

// KEY_dead_capital_schwa is a representation of the C constant GDK_KEY_dead_capital_schwa.
const KEY_dead_capital_schwa = 65163

// KEY_dead_caron is a representation of the C constant GDK_KEY_dead_caron.
const KEY_dead_caron = 65114

// KEY_dead_cedilla is a representation of the C constant GDK_KEY_dead_cedilla.
const KEY_dead_cedilla = 65115

// KEY_dead_circumflex is a representation of the C constant GDK_KEY_dead_circumflex.
const KEY_dead_circumflex = 65106

// KEY_dead_currency is a representation of the C constant GDK_KEY_dead_currency.
const KEY_dead_currency = 65135

// KEY_dead_dasia is a representation of the C constant GDK_KEY_dead_dasia.
const KEY_dead_dasia = 65125

// KEY_dead_diaeresis is a representation of the C constant GDK_KEY_dead_diaeresis.
const KEY_dead_diaeresis = 65111

// KEY_dead_doubleacute is a representation of the C constant GDK_KEY_dead_doubleacute.
const KEY_dead_doubleacute = 65113

// KEY_dead_doublegrave is a representation of the C constant GDK_KEY_dead_doublegrave.
const KEY_dead_doublegrave = 65126

// KEY_dead_e is a representation of the C constant GDK_KEY_dead_e.
const KEY_dead_e = 65154

// KEY_dead_grave is a representation of the C constant GDK_KEY_dead_grave.
const KEY_dead_grave = 65104

// KEY_dead_greek is a representation of the C constant GDK_KEY_dead_greek.
const KEY_dead_greek = 65164

// KEY_dead_hook is a representation of the C constant GDK_KEY_dead_hook.
const KEY_dead_hook = 65121

// KEY_dead_horn is a representation of the C constant GDK_KEY_dead_horn.
const KEY_dead_horn = 65122

// KEY_dead_i is a representation of the C constant GDK_KEY_dead_i.
const KEY_dead_i = 65156

// KEY_dead_invertedbreve is a representation of the C constant GDK_KEY_dead_invertedbreve.
const KEY_dead_invertedbreve = 65133

// KEY_dead_iota is a representation of the C constant GDK_KEY_dead_iota.
const KEY_dead_iota = 65117

// KEY_dead_macron is a representation of the C constant GDK_KEY_dead_macron.
const KEY_dead_macron = 65108

// KEY_dead_o is a representation of the C constant GDK_KEY_dead_o.
const KEY_dead_o = 65158

// KEY_dead_ogonek is a representation of the C constant GDK_KEY_dead_ogonek.
const KEY_dead_ogonek = 65116

// KEY_dead_perispomeni is a representation of the C constant GDK_KEY_dead_perispomeni.
const KEY_dead_perispomeni = 65107

// KEY_dead_psili is a representation of the C constant GDK_KEY_dead_psili.
const KEY_dead_psili = 65124

// KEY_dead_semivoiced_sound is a representation of the C constant GDK_KEY_dead_semivoiced_sound.
const KEY_dead_semivoiced_sound = 65119

// KEY_dead_small_schwa is a representation of the C constant GDK_KEY_dead_small_schwa.
const KEY_dead_small_schwa = 65162

// KEY_dead_stroke is a representation of the C constant GDK_KEY_dead_stroke.
const KEY_dead_stroke = 65123

// KEY_dead_tilde is a representation of the C constant GDK_KEY_dead_tilde.
const KEY_dead_tilde = 65107

// KEY_dead_u is a representation of the C constant GDK_KEY_dead_u.
const KEY_dead_u = 65160

// KEY_dead_voiced_sound is a representation of the C constant GDK_KEY_dead_voiced_sound.
const KEY_dead_voiced_sound = 65118

// KEY_decimalpoint is a representation of the C constant GDK_KEY_decimalpoint.
const KEY_decimalpoint = 2749

// KEY_degree is a representation of the C constant GDK_KEY_degree.
const KEY_degree = 176

// KEY_diaeresis is a representation of the C constant GDK_KEY_diaeresis.
const KEY_diaeresis = 168

// KEY_diamond is a representation of the C constant GDK_KEY_diamond.
const KEY_diamond = 2797

// KEY_digitspace is a representation of the C constant GDK_KEY_digitspace.
const KEY_digitspace = 2725

// KEY_dintegral is a representation of the C constant GDK_KEY_dintegral.
const KEY_dintegral = 16785964

// KEY_division is a representation of the C constant GDK_KEY_division.
const KEY_division = 247

// KEY_dollar is a representation of the C constant GDK_KEY_dollar.
const KEY_dollar = 36

// KEY_doubbaselinedot is a representation of the C constant GDK_KEY_doubbaselinedot.
const KEY_doubbaselinedot = 2735

// KEY_doubleacute is a representation of the C constant GDK_KEY_doubleacute.
const KEY_doubleacute = 445

// KEY_doubledagger is a representation of the C constant GDK_KEY_doubledagger.
const KEY_doubledagger = 2802

// KEY_doublelowquotemark is a representation of the C constant GDK_KEY_doublelowquotemark.
const KEY_doublelowquotemark = 2814

// KEY_downarrow is a representation of the C constant GDK_KEY_downarrow.
const KEY_downarrow = 2302

// KEY_downcaret is a representation of the C constant GDK_KEY_downcaret.
const KEY_downcaret = 2984

// KEY_downshoe is a representation of the C constant GDK_KEY_downshoe.
const KEY_downshoe = 3030

// KEY_downstile is a representation of the C constant GDK_KEY_downstile.
const KEY_downstile = 3012

// KEY_downtack is a representation of the C constant GDK_KEY_downtack.
const KEY_downtack = 3010

// KEY_dstroke is a representation of the C constant GDK_KEY_dstroke.
const KEY_dstroke = 496

// KEY_e is a representation of the C constant GDK_KEY_e.
const KEY_e = 101

// KEY_eabovedot is a representation of the C constant GDK_KEY_eabovedot.
const KEY_eabovedot = 1004

// KEY_eacute is a representation of the C constant GDK_KEY_eacute.
const KEY_eacute = 233

// KEY_ebelowdot is a representation of the C constant GDK_KEY_ebelowdot.
const KEY_ebelowdot = 16785081

// KEY_ecaron is a representation of the C constant GDK_KEY_ecaron.
const KEY_ecaron = 492

// KEY_ecircumflex is a representation of the C constant GDK_KEY_ecircumflex.
const KEY_ecircumflex = 234

// KEY_ecircumflexacute is a representation of the C constant GDK_KEY_ecircumflexacute.
const KEY_ecircumflexacute = 16785087

// KEY_ecircumflexbelowdot is a representation of the C constant GDK_KEY_ecircumflexbelowdot.
const KEY_ecircumflexbelowdot = 16785095

// KEY_ecircumflexgrave is a representation of the C constant GDK_KEY_ecircumflexgrave.
const KEY_ecircumflexgrave = 16785089

// KEY_ecircumflexhook is a representation of the C constant GDK_KEY_ecircumflexhook.
const KEY_ecircumflexhook = 16785091

// KEY_ecircumflextilde is a representation of the C constant GDK_KEY_ecircumflextilde.
const KEY_ecircumflextilde = 16785093

// KEY_ediaeresis is a representation of the C constant GDK_KEY_ediaeresis.
const KEY_ediaeresis = 235

// KEY_egrave is a representation of the C constant GDK_KEY_egrave.
const KEY_egrave = 232

// KEY_ehook is a representation of the C constant GDK_KEY_ehook.
const KEY_ehook = 16785083

// KEY_eightsubscript is a representation of the C constant GDK_KEY_eightsubscript.
const KEY_eightsubscript = 16785544

// KEY_eightsuperior is a representation of the C constant GDK_KEY_eightsuperior.
const KEY_eightsuperior = 16785528

// KEY_elementof is a representation of the C constant GDK_KEY_elementof.
const KEY_elementof = 16785928

// KEY_ellipsis is a representation of the C constant GDK_KEY_ellipsis.
const KEY_ellipsis = 2734

// KEY_em3space is a representation of the C constant GDK_KEY_em3space.
const KEY_em3space = 2723

// KEY_em4space is a representation of the C constant GDK_KEY_em4space.
const KEY_em4space = 2724

// KEY_emacron is a representation of the C constant GDK_KEY_emacron.
const KEY_emacron = 954

// KEY_emdash is a representation of the C constant GDK_KEY_emdash.
const KEY_emdash = 2729

// KEY_emfilledcircle is a representation of the C constant GDK_KEY_emfilledcircle.
const KEY_emfilledcircle = 2782

// KEY_emfilledrect is a representation of the C constant GDK_KEY_emfilledrect.
const KEY_emfilledrect = 2783

// KEY_emopencircle is a representation of the C constant GDK_KEY_emopencircle.
const KEY_emopencircle = 2766

// KEY_emopenrectangle is a representation of the C constant GDK_KEY_emopenrectangle.
const KEY_emopenrectangle = 2767

// KEY_emptyset is a representation of the C constant GDK_KEY_emptyset.
const KEY_emptyset = 16785925

// KEY_emspace is a representation of the C constant GDK_KEY_emspace.
const KEY_emspace = 2721

// KEY_endash is a representation of the C constant GDK_KEY_endash.
const KEY_endash = 2730

// KEY_enfilledcircbullet is a representation of the C constant GDK_KEY_enfilledcircbullet.
const KEY_enfilledcircbullet = 2790

// KEY_enfilledsqbullet is a representation of the C constant GDK_KEY_enfilledsqbullet.
const KEY_enfilledsqbullet = 2791

// KEY_eng is a representation of the C constant GDK_KEY_eng.
const KEY_eng = 959

// KEY_enopencircbullet is a representation of the C constant GDK_KEY_enopencircbullet.
const KEY_enopencircbullet = 2784

// KEY_enopensquarebullet is a representation of the C constant GDK_KEY_enopensquarebullet.
const KEY_enopensquarebullet = 2785

// KEY_enspace is a representation of the C constant GDK_KEY_enspace.
const KEY_enspace = 2722

// KEY_eogonek is a representation of the C constant GDK_KEY_eogonek.
const KEY_eogonek = 490

// KEY_equal is a representation of the C constant GDK_KEY_equal.
const KEY_equal = 61

// KEY_eth is a representation of the C constant GDK_KEY_eth.
const KEY_eth = 240

// KEY_etilde is a representation of the C constant GDK_KEY_etilde.
const KEY_etilde = 16785085

// KEY_exclam is a representation of the C constant GDK_KEY_exclam.
const KEY_exclam = 33

// KEY_exclamdown is a representation of the C constant GDK_KEY_exclamdown.
const KEY_exclamdown = 161

// KEY_ezh is a representation of the C constant GDK_KEY_ezh.
const KEY_ezh = 16777874

// KEY_f is a representation of the C constant GDK_KEY_f.
const KEY_f = 102

// KEY_fabovedot is a representation of the C constant GDK_KEY_fabovedot.
const KEY_fabovedot = 16784927

// KEY_femalesymbol is a representation of the C constant GDK_KEY_femalesymbol.
const KEY_femalesymbol = 2808

// KEY_ff is a representation of the C constant GDK_KEY_ff.
const KEY_ff = 2531

// KEY_figdash is a representation of the C constant GDK_KEY_figdash.
const KEY_figdash = 2747

// KEY_filledlefttribullet is a representation of the C constant GDK_KEY_filledlefttribullet.
const KEY_filledlefttribullet = 2780

// KEY_filledrectbullet is a representation of the C constant GDK_KEY_filledrectbullet.
const KEY_filledrectbullet = 2779

// KEY_filledrighttribullet is a representation of the C constant GDK_KEY_filledrighttribullet.
const KEY_filledrighttribullet = 2781

// KEY_filledtribulletdown is a representation of the C constant GDK_KEY_filledtribulletdown.
const KEY_filledtribulletdown = 2793

// KEY_filledtribulletup is a representation of the C constant GDK_KEY_filledtribulletup.
const KEY_filledtribulletup = 2792

// KEY_fiveeighths is a representation of the C constant GDK_KEY_fiveeighths.
const KEY_fiveeighths = 2757

// KEY_fivesixths is a representation of the C constant GDK_KEY_fivesixths.
const KEY_fivesixths = 2743

// KEY_fivesubscript is a representation of the C constant GDK_KEY_fivesubscript.
const KEY_fivesubscript = 16785541

// KEY_fivesuperior is a representation of the C constant GDK_KEY_fivesuperior.
const KEY_fivesuperior = 16785525

// KEY_fourfifths is a representation of the C constant GDK_KEY_fourfifths.
const KEY_fourfifths = 2741

// KEY_foursubscript is a representation of the C constant GDK_KEY_foursubscript.
const KEY_foursubscript = 16785540

// KEY_foursuperior is a representation of the C constant GDK_KEY_foursuperior.
const KEY_foursuperior = 16785524

// KEY_fourthroot is a representation of the C constant GDK_KEY_fourthroot.
const KEY_fourthroot = 16785948

// KEY_function is a representation of the C constant GDK_KEY_function.
const KEY_function = 2294

// KEY_g is a representation of the C constant GDK_KEY_g.
const KEY_g = 103

// KEY_gabovedot is a representation of the C constant GDK_KEY_gabovedot.
const KEY_gabovedot = 757

// KEY_gbreve is a representation of the C constant GDK_KEY_gbreve.
const KEY_gbreve = 699

// KEY_gcaron is a representation of the C constant GDK_KEY_gcaron.
const KEY_gcaron = 16777703

// KEY_gcedilla is a representation of the C constant GDK_KEY_gcedilla.
const KEY_gcedilla = 955

// KEY_gcircumflex is a representation of the C constant GDK_KEY_gcircumflex.
const KEY_gcircumflex = 760

// KEY_grave is a representation of the C constant GDK_KEY_grave.
const KEY_grave = 96

// KEY_greater is a representation of the C constant GDK_KEY_greater.
const KEY_greater = 62

// KEY_greaterthanequal is a representation of the C constant GDK_KEY_greaterthanequal.
const KEY_greaterthanequal = 2238

// KEY_guillemotleft is a representation of the C constant GDK_KEY_guillemotleft.
const KEY_guillemotleft = 171

// KEY_guillemotright is a representation of the C constant GDK_KEY_guillemotright.
const KEY_guillemotright = 187

// KEY_h is a representation of the C constant GDK_KEY_h.
const KEY_h = 104

// KEY_hairspace is a representation of the C constant GDK_KEY_hairspace.
const KEY_hairspace = 2728

// KEY_hcircumflex is a representation of the C constant GDK_KEY_hcircumflex.
const KEY_hcircumflex = 694

// KEY_heart is a representation of the C constant GDK_KEY_heart.
const KEY_heart = 2798

// KEY_hebrew_aleph is a representation of the C constant GDK_KEY_hebrew_aleph.
const KEY_hebrew_aleph = 3296

// KEY_hebrew_ayin is a representation of the C constant GDK_KEY_hebrew_ayin.
const KEY_hebrew_ayin = 3314

// KEY_hebrew_bet is a representation of the C constant GDK_KEY_hebrew_bet.
const KEY_hebrew_bet = 3297

// KEY_hebrew_beth is a representation of the C constant GDK_KEY_hebrew_beth.
const KEY_hebrew_beth = 3297

// KEY_hebrew_chet is a representation of the C constant GDK_KEY_hebrew_chet.
const KEY_hebrew_chet = 3303

// KEY_hebrew_dalet is a representation of the C constant GDK_KEY_hebrew_dalet.
const KEY_hebrew_dalet = 3299

// KEY_hebrew_daleth is a representation of the C constant GDK_KEY_hebrew_daleth.
const KEY_hebrew_daleth = 3299

// KEY_hebrew_doublelowline is a representation of the C constant GDK_KEY_hebrew_doublelowline.
const KEY_hebrew_doublelowline = 3295

// KEY_hebrew_finalkaph is a representation of the C constant GDK_KEY_hebrew_finalkaph.
const KEY_hebrew_finalkaph = 3306

// KEY_hebrew_finalmem is a representation of the C constant GDK_KEY_hebrew_finalmem.
const KEY_hebrew_finalmem = 3309

// KEY_hebrew_finalnun is a representation of the C constant GDK_KEY_hebrew_finalnun.
const KEY_hebrew_finalnun = 3311

// KEY_hebrew_finalpe is a representation of the C constant GDK_KEY_hebrew_finalpe.
const KEY_hebrew_finalpe = 3315

// KEY_hebrew_finalzade is a representation of the C constant GDK_KEY_hebrew_finalzade.
const KEY_hebrew_finalzade = 3317

// KEY_hebrew_finalzadi is a representation of the C constant GDK_KEY_hebrew_finalzadi.
const KEY_hebrew_finalzadi = 3317

// KEY_hebrew_gimel is a representation of the C constant GDK_KEY_hebrew_gimel.
const KEY_hebrew_gimel = 3298

// KEY_hebrew_gimmel is a representation of the C constant GDK_KEY_hebrew_gimmel.
const KEY_hebrew_gimmel = 3298

// KEY_hebrew_he is a representation of the C constant GDK_KEY_hebrew_he.
const KEY_hebrew_he = 3300

// KEY_hebrew_het is a representation of the C constant GDK_KEY_hebrew_het.
const KEY_hebrew_het = 3303

// KEY_hebrew_kaph is a representation of the C constant GDK_KEY_hebrew_kaph.
const KEY_hebrew_kaph = 3307

// KEY_hebrew_kuf is a representation of the C constant GDK_KEY_hebrew_kuf.
const KEY_hebrew_kuf = 3319

// KEY_hebrew_lamed is a representation of the C constant GDK_KEY_hebrew_lamed.
const KEY_hebrew_lamed = 3308

// KEY_hebrew_mem is a representation of the C constant GDK_KEY_hebrew_mem.
const KEY_hebrew_mem = 3310

// KEY_hebrew_nun is a representation of the C constant GDK_KEY_hebrew_nun.
const KEY_hebrew_nun = 3312

// KEY_hebrew_pe is a representation of the C constant GDK_KEY_hebrew_pe.
const KEY_hebrew_pe = 3316

// KEY_hebrew_qoph is a representation of the C constant GDK_KEY_hebrew_qoph.
const KEY_hebrew_qoph = 3319

// KEY_hebrew_resh is a representation of the C constant GDK_KEY_hebrew_resh.
const KEY_hebrew_resh = 3320

// KEY_hebrew_samech is a representation of the C constant GDK_KEY_hebrew_samech.
const KEY_hebrew_samech = 3313

// KEY_hebrew_samekh is a representation of the C constant GDK_KEY_hebrew_samekh.
const KEY_hebrew_samekh = 3313

// KEY_hebrew_shin is a representation of the C constant GDK_KEY_hebrew_shin.
const KEY_hebrew_shin = 3321

// KEY_hebrew_taf is a representation of the C constant GDK_KEY_hebrew_taf.
const KEY_hebrew_taf = 3322

// KEY_hebrew_taw is a representation of the C constant GDK_KEY_hebrew_taw.
const KEY_hebrew_taw = 3322

// KEY_hebrew_tet is a representation of the C constant GDK_KEY_hebrew_tet.
const KEY_hebrew_tet = 3304

// KEY_hebrew_teth is a representation of the C constant GDK_KEY_hebrew_teth.
const KEY_hebrew_teth = 3304

// KEY_hebrew_waw is a representation of the C constant GDK_KEY_hebrew_waw.
const KEY_hebrew_waw = 3301

// KEY_hebrew_yod is a representation of the C constant GDK_KEY_hebrew_yod.
const KEY_hebrew_yod = 3305

// KEY_hebrew_zade is a representation of the C constant GDK_KEY_hebrew_zade.
const KEY_hebrew_zade = 3318

// KEY_hebrew_zadi is a representation of the C constant GDK_KEY_hebrew_zadi.
const KEY_hebrew_zadi = 3318

// KEY_hebrew_zain is a representation of the C constant GDK_KEY_hebrew_zain.
const KEY_hebrew_zain = 3302

// KEY_hebrew_zayin is a representation of the C constant GDK_KEY_hebrew_zayin.
const KEY_hebrew_zayin = 3302

// KEY_hexagram is a representation of the C constant GDK_KEY_hexagram.
const KEY_hexagram = 2778

// KEY_horizconnector is a representation of the C constant GDK_KEY_horizconnector.
const KEY_horizconnector = 2211

// KEY_horizlinescan1 is a representation of the C constant GDK_KEY_horizlinescan1.
const KEY_horizlinescan1 = 2543

// KEY_horizlinescan3 is a representation of the C constant GDK_KEY_horizlinescan3.
const KEY_horizlinescan3 = 2544

// KEY_horizlinescan5 is a representation of the C constant GDK_KEY_horizlinescan5.
const KEY_horizlinescan5 = 2545

// KEY_horizlinescan7 is a representation of the C constant GDK_KEY_horizlinescan7.
const KEY_horizlinescan7 = 2546

// KEY_horizlinescan9 is a representation of the C constant GDK_KEY_horizlinescan9.
const KEY_horizlinescan9 = 2547

// KEY_hstroke is a representation of the C constant GDK_KEY_hstroke.
const KEY_hstroke = 689

// KEY_ht is a representation of the C constant GDK_KEY_ht.
const KEY_ht = 2530

// KEY_hyphen is a representation of the C constant GDK_KEY_hyphen.
const KEY_hyphen = 173

// KEY_i is a representation of the C constant GDK_KEY_i.
const KEY_i = 105

// KEY_iTouch is a representation of the C constant GDK_KEY_iTouch.
const KEY_iTouch = 269025120

// KEY_iacute is a representation of the C constant GDK_KEY_iacute.
const KEY_iacute = 237

// KEY_ibelowdot is a representation of the C constant GDK_KEY_ibelowdot.
const KEY_ibelowdot = 16785099

// KEY_ibreve is a representation of the C constant GDK_KEY_ibreve.
const KEY_ibreve = 16777517

// KEY_icircumflex is a representation of the C constant GDK_KEY_icircumflex.
const KEY_icircumflex = 238

// KEY_identical is a representation of the C constant GDK_KEY_identical.
const KEY_identical = 2255

// KEY_idiaeresis is a representation of the C constant GDK_KEY_idiaeresis.
const KEY_idiaeresis = 239

// KEY_idotless is a representation of the C constant GDK_KEY_idotless.
const KEY_idotless = 697

// KEY_ifonlyif is a representation of the C constant GDK_KEY_ifonlyif.
const KEY_ifonlyif = 2253

// KEY_igrave is a representation of the C constant GDK_KEY_igrave.
const KEY_igrave = 236

// KEY_ihook is a representation of the C constant GDK_KEY_ihook.
const KEY_ihook = 16785097

// KEY_imacron is a representation of the C constant GDK_KEY_imacron.
const KEY_imacron = 1007

// KEY_implies is a representation of the C constant GDK_KEY_implies.
const KEY_implies = 2254

// KEY_includedin is a representation of the C constant GDK_KEY_includedin.
const KEY_includedin = 2266

// KEY_includes is a representation of the C constant GDK_KEY_includes.
const KEY_includes = 2267

// KEY_infinity is a representation of the C constant GDK_KEY_infinity.
const KEY_infinity = 2242

// KEY_integral is a representation of the C constant GDK_KEY_integral.
const KEY_integral = 2239

// KEY_intersection is a representation of the C constant GDK_KEY_intersection.
const KEY_intersection = 2268

// KEY_iogonek is a representation of the C constant GDK_KEY_iogonek.
const KEY_iogonek = 999

// KEY_itilde is a representation of the C constant GDK_KEY_itilde.
const KEY_itilde = 949

// KEY_j is a representation of the C constant GDK_KEY_j.
const KEY_j = 106

// KEY_jcircumflex is a representation of the C constant GDK_KEY_jcircumflex.
const KEY_jcircumflex = 700

// KEY_jot is a representation of the C constant GDK_KEY_jot.
const KEY_jot = 3018

// KEY_k is a representation of the C constant GDK_KEY_k.
const KEY_k = 107

// KEY_kana_A is a representation of the C constant GDK_KEY_kana_A.
const KEY_kana_A = 1201

// KEY_kana_CHI is a representation of the C constant GDK_KEY_kana_CHI.
const KEY_kana_CHI = 1217

// KEY_kana_E is a representation of the C constant GDK_KEY_kana_E.
const KEY_kana_E = 1204

// KEY_kana_FU is a representation of the C constant GDK_KEY_kana_FU.
const KEY_kana_FU = 1228

// KEY_kana_HA is a representation of the C constant GDK_KEY_kana_HA.
const KEY_kana_HA = 1226

// KEY_kana_HE is a representation of the C constant GDK_KEY_kana_HE.
const KEY_kana_HE = 1229

// KEY_kana_HI is a representation of the C constant GDK_KEY_kana_HI.
const KEY_kana_HI = 1227

// KEY_kana_HO is a representation of the C constant GDK_KEY_kana_HO.
const KEY_kana_HO = 1230

// KEY_kana_HU is a representation of the C constant GDK_KEY_kana_HU.
const KEY_kana_HU = 1228

// KEY_kana_I is a representation of the C constant GDK_KEY_kana_I.
const KEY_kana_I = 1202

// KEY_kana_KA is a representation of the C constant GDK_KEY_kana_KA.
const KEY_kana_KA = 1206

// KEY_kana_KE is a representation of the C constant GDK_KEY_kana_KE.
const KEY_kana_KE = 1209

// KEY_kana_KI is a representation of the C constant GDK_KEY_kana_KI.
const KEY_kana_KI = 1207

// KEY_kana_KO is a representation of the C constant GDK_KEY_kana_KO.
const KEY_kana_KO = 1210

// KEY_kana_KU is a representation of the C constant GDK_KEY_kana_KU.
const KEY_kana_KU = 1208

// KEY_kana_MA is a representation of the C constant GDK_KEY_kana_MA.
const KEY_kana_MA = 1231

// KEY_kana_ME is a representation of the C constant GDK_KEY_kana_ME.
const KEY_kana_ME = 1234

// KEY_kana_MI is a representation of the C constant GDK_KEY_kana_MI.
const KEY_kana_MI = 1232

// KEY_kana_MO is a representation of the C constant GDK_KEY_kana_MO.
const KEY_kana_MO = 1235

// KEY_kana_MU is a representation of the C constant GDK_KEY_kana_MU.
const KEY_kana_MU = 1233

// KEY_kana_N is a representation of the C constant GDK_KEY_kana_N.
const KEY_kana_N = 1245

// KEY_kana_NA is a representation of the C constant GDK_KEY_kana_NA.
const KEY_kana_NA = 1221

// KEY_kana_NE is a representation of the C constant GDK_KEY_kana_NE.
const KEY_kana_NE = 1224

// KEY_kana_NI is a representation of the C constant GDK_KEY_kana_NI.
const KEY_kana_NI = 1222

// KEY_kana_NO is a representation of the C constant GDK_KEY_kana_NO.
const KEY_kana_NO = 1225

// KEY_kana_NU is a representation of the C constant GDK_KEY_kana_NU.
const KEY_kana_NU = 1223

// KEY_kana_O is a representation of the C constant GDK_KEY_kana_O.
const KEY_kana_O = 1205

// KEY_kana_RA is a representation of the C constant GDK_KEY_kana_RA.
const KEY_kana_RA = 1239

// KEY_kana_RE is a representation of the C constant GDK_KEY_kana_RE.
const KEY_kana_RE = 1242

// KEY_kana_RI is a representation of the C constant GDK_KEY_kana_RI.
const KEY_kana_RI = 1240

// KEY_kana_RO is a representation of the C constant GDK_KEY_kana_RO.
const KEY_kana_RO = 1243

// KEY_kana_RU is a representation of the C constant GDK_KEY_kana_RU.
const KEY_kana_RU = 1241

// KEY_kana_SA is a representation of the C constant GDK_KEY_kana_SA.
const KEY_kana_SA = 1211

// KEY_kana_SE is a representation of the C constant GDK_KEY_kana_SE.
const KEY_kana_SE = 1214

// KEY_kana_SHI is a representation of the C constant GDK_KEY_kana_SHI.
const KEY_kana_SHI = 1212

// KEY_kana_SO is a representation of the C constant GDK_KEY_kana_SO.
const KEY_kana_SO = 1215

// KEY_kana_SU is a representation of the C constant GDK_KEY_kana_SU.
const KEY_kana_SU = 1213

// KEY_kana_TA is a representation of the C constant GDK_KEY_kana_TA.
const KEY_kana_TA = 1216

// KEY_kana_TE is a representation of the C constant GDK_KEY_kana_TE.
const KEY_kana_TE = 1219

// KEY_kana_TI is a representation of the C constant GDK_KEY_kana_TI.
const KEY_kana_TI = 1217

// KEY_kana_TO is a representation of the C constant GDK_KEY_kana_TO.
const KEY_kana_TO = 1220

// KEY_kana_TSU is a representation of the C constant GDK_KEY_kana_TSU.
const KEY_kana_TSU = 1218

// KEY_kana_TU is a representation of the C constant GDK_KEY_kana_TU.
const KEY_kana_TU = 1218

// KEY_kana_U is a representation of the C constant GDK_KEY_kana_U.
const KEY_kana_U = 1203

// KEY_kana_WA is a representation of the C constant GDK_KEY_kana_WA.
const KEY_kana_WA = 1244

// KEY_kana_WO is a representation of the C constant GDK_KEY_kana_WO.
const KEY_kana_WO = 1190

// KEY_kana_YA is a representation of the C constant GDK_KEY_kana_YA.
const KEY_kana_YA = 1236

// KEY_kana_YO is a representation of the C constant GDK_KEY_kana_YO.
const KEY_kana_YO = 1238

// KEY_kana_YU is a representation of the C constant GDK_KEY_kana_YU.
const KEY_kana_YU = 1237

// KEY_kana_a is a representation of the C constant GDK_KEY_kana_a.
const KEY_kana_a = 1191

// KEY_kana_closingbracket is a representation of the C constant GDK_KEY_kana_closingbracket.
const KEY_kana_closingbracket = 1187

// KEY_kana_comma is a representation of the C constant GDK_KEY_kana_comma.
const KEY_kana_comma = 1188

// KEY_kana_conjunctive is a representation of the C constant GDK_KEY_kana_conjunctive.
const KEY_kana_conjunctive = 1189

// KEY_kana_e is a representation of the C constant GDK_KEY_kana_e.
const KEY_kana_e = 1194

// KEY_kana_fullstop is a representation of the C constant GDK_KEY_kana_fullstop.
const KEY_kana_fullstop = 1185

// KEY_kana_i is a representation of the C constant GDK_KEY_kana_i.
const KEY_kana_i = 1192

// KEY_kana_middledot is a representation of the C constant GDK_KEY_kana_middledot.
const KEY_kana_middledot = 1189

// KEY_kana_o is a representation of the C constant GDK_KEY_kana_o.
const KEY_kana_o = 1195

// KEY_kana_openingbracket is a representation of the C constant GDK_KEY_kana_openingbracket.
const KEY_kana_openingbracket = 1186

// KEY_kana_switch is a representation of the C constant GDK_KEY_kana_switch.
const KEY_kana_switch = 65406

// KEY_kana_tsu is a representation of the C constant GDK_KEY_kana_tsu.
const KEY_kana_tsu = 1199

// KEY_kana_tu is a representation of the C constant GDK_KEY_kana_tu.
const KEY_kana_tu = 1199

// KEY_kana_u is a representation of the C constant GDK_KEY_kana_u.
const KEY_kana_u = 1193

// KEY_kana_ya is a representation of the C constant GDK_KEY_kana_ya.
const KEY_kana_ya = 1196

// KEY_kana_yo is a representation of the C constant GDK_KEY_kana_yo.
const KEY_kana_yo = 1198

// KEY_kana_yu is a representation of the C constant GDK_KEY_kana_yu.
const KEY_kana_yu = 1197

// KEY_kappa is a representation of the C constant GDK_KEY_kappa.
const KEY_kappa = 930

// KEY_kcedilla is a representation of the C constant GDK_KEY_kcedilla.
const KEY_kcedilla = 1011

// KEY_kra is a representation of the C constant GDK_KEY_kra.
const KEY_kra = 930

// KEY_l is a representation of the C constant GDK_KEY_l.
const KEY_l = 108

// KEY_lacute is a representation of the C constant GDK_KEY_lacute.
const KEY_lacute = 485

// KEY_latincross is a representation of the C constant GDK_KEY_latincross.
const KEY_latincross = 2777

// KEY_lbelowdot is a representation of the C constant GDK_KEY_lbelowdot.
const KEY_lbelowdot = 16784951

// KEY_lcaron is a representation of the C constant GDK_KEY_lcaron.
const KEY_lcaron = 437

// KEY_lcedilla is a representation of the C constant GDK_KEY_lcedilla.
const KEY_lcedilla = 950

// KEY_leftanglebracket is a representation of the C constant GDK_KEY_leftanglebracket.
const KEY_leftanglebracket = 2748

// KEY_leftarrow is a representation of the C constant GDK_KEY_leftarrow.
const KEY_leftarrow = 2299

// KEY_leftcaret is a representation of the C constant GDK_KEY_leftcaret.
const KEY_leftcaret = 2979

// KEY_leftdoublequotemark is a representation of the C constant GDK_KEY_leftdoublequotemark.
const KEY_leftdoublequotemark = 2770

// KEY_leftmiddlecurlybrace is a representation of the C constant GDK_KEY_leftmiddlecurlybrace.
const KEY_leftmiddlecurlybrace = 2223

// KEY_leftopentriangle is a representation of the C constant GDK_KEY_leftopentriangle.
const KEY_leftopentriangle = 2764

// KEY_leftpointer is a representation of the C constant GDK_KEY_leftpointer.
const KEY_leftpointer = 2794

// KEY_leftradical is a representation of the C constant GDK_KEY_leftradical.
const KEY_leftradical = 2209

// KEY_leftshoe is a representation of the C constant GDK_KEY_leftshoe.
const KEY_leftshoe = 3034

// KEY_leftsinglequotemark is a representation of the C constant GDK_KEY_leftsinglequotemark.
const KEY_leftsinglequotemark = 2768

// KEY_leftt is a representation of the C constant GDK_KEY_leftt.
const KEY_leftt = 2548

// KEY_lefttack is a representation of the C constant GDK_KEY_lefttack.
const KEY_lefttack = 3036

// KEY_less is a representation of the C constant GDK_KEY_less.
const KEY_less = 60

// KEY_lessthanequal is a representation of the C constant GDK_KEY_lessthanequal.
const KEY_lessthanequal = 2236

// KEY_lf is a representation of the C constant GDK_KEY_lf.
const KEY_lf = 2533

// KEY_logicaland is a representation of the C constant GDK_KEY_logicaland.
const KEY_logicaland = 2270

// KEY_logicalor is a representation of the C constant GDK_KEY_logicalor.
const KEY_logicalor = 2271

// KEY_lowleftcorner is a representation of the C constant GDK_KEY_lowleftcorner.
const KEY_lowleftcorner = 2541

// KEY_lowrightcorner is a representation of the C constant GDK_KEY_lowrightcorner.
const KEY_lowrightcorner = 2538

// KEY_lstroke is a representation of the C constant GDK_KEY_lstroke.
const KEY_lstroke = 435

// KEY_m is a representation of the C constant GDK_KEY_m.
const KEY_m = 109

// KEY_mabovedot is a representation of the C constant GDK_KEY_mabovedot.
const KEY_mabovedot = 16784961

// KEY_macron is a representation of the C constant GDK_KEY_macron.
const KEY_macron = 175

// KEY_malesymbol is a representation of the C constant GDK_KEY_malesymbol.
const KEY_malesymbol = 2807

// KEY_maltesecross is a representation of the C constant GDK_KEY_maltesecross.
const KEY_maltesecross = 2800

// KEY_marker is a representation of the C constant GDK_KEY_marker.
const KEY_marker = 2751

// KEY_masculine is a representation of the C constant GDK_KEY_masculine.
const KEY_masculine = 186

// KEY_minus is a representation of the C constant GDK_KEY_minus.
const KEY_minus = 45

// KEY_minutes is a representation of the C constant GDK_KEY_minutes.
const KEY_minutes = 2774

// KEY_mu is a representation of the C constant GDK_KEY_mu.
const KEY_mu = 181

// KEY_multiply is a representation of the C constant GDK_KEY_multiply.
const KEY_multiply = 215

// KEY_musicalflat is a representation of the C constant GDK_KEY_musicalflat.
const KEY_musicalflat = 2806

// KEY_musicalsharp is a representation of the C constant GDK_KEY_musicalsharp.
const KEY_musicalsharp = 2805

// KEY_n is a representation of the C constant GDK_KEY_n.
const KEY_n = 110

// KEY_nabla is a representation of the C constant GDK_KEY_nabla.
const KEY_nabla = 2245

// KEY_nacute is a representation of the C constant GDK_KEY_nacute.
const KEY_nacute = 497

// KEY_ncaron is a representation of the C constant GDK_KEY_ncaron.
const KEY_ncaron = 498

// KEY_ncedilla is a representation of the C constant GDK_KEY_ncedilla.
const KEY_ncedilla = 1009

// KEY_ninesubscript is a representation of the C constant GDK_KEY_ninesubscript.
const KEY_ninesubscript = 16785545

// KEY_ninesuperior is a representation of the C constant GDK_KEY_ninesuperior.
const KEY_ninesuperior = 16785529

// KEY_nl is a representation of the C constant GDK_KEY_nl.
const KEY_nl = 2536

// KEY_nobreakspace is a representation of the C constant GDK_KEY_nobreakspace.
const KEY_nobreakspace = 160

// KEY_notapproxeq is a representation of the C constant GDK_KEY_notapproxeq.
const KEY_notapproxeq = 16785991

// KEY_notelementof is a representation of the C constant GDK_KEY_notelementof.
const KEY_notelementof = 16785929

// KEY_notequal is a representation of the C constant GDK_KEY_notequal.
const KEY_notequal = 2237

// KEY_notidentical is a representation of the C constant GDK_KEY_notidentical.
const KEY_notidentical = 16786018

// KEY_notsign is a representation of the C constant GDK_KEY_notsign.
const KEY_notsign = 172

// KEY_ntilde is a representation of the C constant GDK_KEY_ntilde.
const KEY_ntilde = 241

// KEY_numbersign is a representation of the C constant GDK_KEY_numbersign.
const KEY_numbersign = 35

// KEY_numerosign is a representation of the C constant GDK_KEY_numerosign.
const KEY_numerosign = 1712

// KEY_o is a representation of the C constant GDK_KEY_o.
const KEY_o = 111

// KEY_oacute is a representation of the C constant GDK_KEY_oacute.
const KEY_oacute = 243

// KEY_obarred is a representation of the C constant GDK_KEY_obarred.
const KEY_obarred = 16777845

// KEY_obelowdot is a representation of the C constant GDK_KEY_obelowdot.
const KEY_obelowdot = 16785101

// KEY_ocaron is a representation of the C constant GDK_KEY_ocaron.
const KEY_ocaron = 16777682

// KEY_ocircumflex is a representation of the C constant GDK_KEY_ocircumflex.
const KEY_ocircumflex = 244

// KEY_ocircumflexacute is a representation of the C constant GDK_KEY_ocircumflexacute.
const KEY_ocircumflexacute = 16785105

// KEY_ocircumflexbelowdot is a representation of the C constant GDK_KEY_ocircumflexbelowdot.
const KEY_ocircumflexbelowdot = 16785113

// KEY_ocircumflexgrave is a representation of the C constant GDK_KEY_ocircumflexgrave.
const KEY_ocircumflexgrave = 16785107

// KEY_ocircumflexhook is a representation of the C constant GDK_KEY_ocircumflexhook.
const KEY_ocircumflexhook = 16785109

// KEY_ocircumflextilde is a representation of the C constant GDK_KEY_ocircumflextilde.
const KEY_ocircumflextilde = 16785111

// KEY_odiaeresis is a representation of the C constant GDK_KEY_odiaeresis.
const KEY_odiaeresis = 246

// KEY_odoubleacute is a representation of the C constant GDK_KEY_odoubleacute.
const KEY_odoubleacute = 501

// KEY_oe is a representation of the C constant GDK_KEY_oe.
const KEY_oe = 5053

// KEY_ogonek is a representation of the C constant GDK_KEY_ogonek.
const KEY_ogonek = 434

// KEY_ograve is a representation of the C constant GDK_KEY_ograve.
const KEY_ograve = 242

// KEY_ohook is a representation of the C constant GDK_KEY_ohook.
const KEY_ohook = 16785103

// KEY_ohorn is a representation of the C constant GDK_KEY_ohorn.
const KEY_ohorn = 16777633

// KEY_ohornacute is a representation of the C constant GDK_KEY_ohornacute.
const KEY_ohornacute = 16785115

// KEY_ohornbelowdot is a representation of the C constant GDK_KEY_ohornbelowdot.
const KEY_ohornbelowdot = 16785123

// KEY_ohorngrave is a representation of the C constant GDK_KEY_ohorngrave.
const KEY_ohorngrave = 16785117

// KEY_ohornhook is a representation of the C constant GDK_KEY_ohornhook.
const KEY_ohornhook = 16785119

// KEY_ohorntilde is a representation of the C constant GDK_KEY_ohorntilde.
const KEY_ohorntilde = 16785121

// KEY_omacron is a representation of the C constant GDK_KEY_omacron.
const KEY_omacron = 1010

// KEY_oneeighth is a representation of the C constant GDK_KEY_oneeighth.
const KEY_oneeighth = 2755

// KEY_onefifth is a representation of the C constant GDK_KEY_onefifth.
const KEY_onefifth = 2738

// KEY_onehalf is a representation of the C constant GDK_KEY_onehalf.
const KEY_onehalf = 189

// KEY_onequarter is a representation of the C constant GDK_KEY_onequarter.
const KEY_onequarter = 188

// KEY_onesixth is a representation of the C constant GDK_KEY_onesixth.
const KEY_onesixth = 2742

// KEY_onesubscript is a representation of the C constant GDK_KEY_onesubscript.
const KEY_onesubscript = 16785537

// KEY_onesuperior is a representation of the C constant GDK_KEY_onesuperior.
const KEY_onesuperior = 185

// KEY_onethird is a representation of the C constant GDK_KEY_onethird.
const KEY_onethird = 2736

// KEY_ooblique is a representation of the C constant GDK_KEY_ooblique.
const KEY_ooblique = 248

// KEY_openrectbullet is a representation of the C constant GDK_KEY_openrectbullet.
const KEY_openrectbullet = 2786

// KEY_openstar is a representation of the C constant GDK_KEY_openstar.
const KEY_openstar = 2789

// KEY_opentribulletdown is a representation of the C constant GDK_KEY_opentribulletdown.
const KEY_opentribulletdown = 2788

// KEY_opentribulletup is a representation of the C constant GDK_KEY_opentribulletup.
const KEY_opentribulletup = 2787

// KEY_ordfeminine is a representation of the C constant GDK_KEY_ordfeminine.
const KEY_ordfeminine = 170

// KEY_oslash is a representation of the C constant GDK_KEY_oslash.
const KEY_oslash = 248

// KEY_otilde is a representation of the C constant GDK_KEY_otilde.
const KEY_otilde = 245

// KEY_overbar is a representation of the C constant GDK_KEY_overbar.
const KEY_overbar = 3008

// KEY_overline is a representation of the C constant GDK_KEY_overline.
const KEY_overline = 1150

// KEY_p is a representation of the C constant GDK_KEY_p.
const KEY_p = 112

// KEY_pabovedot is a representation of the C constant GDK_KEY_pabovedot.
const KEY_pabovedot = 16784983

// KEY_paragraph is a representation of the C constant GDK_KEY_paragraph.
const KEY_paragraph = 182

// KEY_parenleft is a representation of the C constant GDK_KEY_parenleft.
const KEY_parenleft = 40

// KEY_parenright is a representation of the C constant GDK_KEY_parenright.
const KEY_parenright = 41

// KEY_partdifferential is a representation of the C constant GDK_KEY_partdifferential.
const KEY_partdifferential = 16785922

// KEY_partialderivative is a representation of the C constant GDK_KEY_partialderivative.
const KEY_partialderivative = 2287

// KEY_percent is a representation of the C constant GDK_KEY_percent.
const KEY_percent = 37

// KEY_period is a representation of the C constant GDK_KEY_period.
const KEY_period = 46

// KEY_periodcentered is a representation of the C constant GDK_KEY_periodcentered.
const KEY_periodcentered = 183

// KEY_permille is a representation of the C constant GDK_KEY_permille.
const KEY_permille = 2773

// KEY_phonographcopyright is a representation of the C constant GDK_KEY_phonographcopyright.
const KEY_phonographcopyright = 2811

// KEY_plus is a representation of the C constant GDK_KEY_plus.
const KEY_plus = 43

// KEY_plusminus is a representation of the C constant GDK_KEY_plusminus.
const KEY_plusminus = 177

// KEY_prescription is a representation of the C constant GDK_KEY_prescription.
const KEY_prescription = 2772

// KEY_prolongedsound is a representation of the C constant GDK_KEY_prolongedsound.
const KEY_prolongedsound = 1200

// KEY_punctspace is a representation of the C constant GDK_KEY_punctspace.
const KEY_punctspace = 2726

// KEY_q is a representation of the C constant GDK_KEY_q.
const KEY_q = 113

// KEY_quad is a representation of the C constant GDK_KEY_quad.
const KEY_quad = 3020

// KEY_question is a representation of the C constant GDK_KEY_question.
const KEY_question = 63

// KEY_questiondown is a representation of the C constant GDK_KEY_questiondown.
const KEY_questiondown = 191

// KEY_quotedbl is a representation of the C constant GDK_KEY_quotedbl.
const KEY_quotedbl = 34

// KEY_quoteleft is a representation of the C constant GDK_KEY_quoteleft.
const KEY_quoteleft = 96

// KEY_quoteright is a representation of the C constant GDK_KEY_quoteright.
const KEY_quoteright = 39

// KEY_r is a representation of the C constant GDK_KEY_r.
const KEY_r = 114

// KEY_racute is a representation of the C constant GDK_KEY_racute.
const KEY_racute = 480

// KEY_radical is a representation of the C constant GDK_KEY_radical.
const KEY_radical = 2262

// KEY_rcaron is a representation of the C constant GDK_KEY_rcaron.
const KEY_rcaron = 504

// KEY_rcedilla is a representation of the C constant GDK_KEY_rcedilla.
const KEY_rcedilla = 947

// KEY_registered is a representation of the C constant GDK_KEY_registered.
const KEY_registered = 174

// KEY_rightanglebracket is a representation of the C constant GDK_KEY_rightanglebracket.
const KEY_rightanglebracket = 2750

// KEY_rightarrow is a representation of the C constant GDK_KEY_rightarrow.
const KEY_rightarrow = 2301

// KEY_rightcaret is a representation of the C constant GDK_KEY_rightcaret.
const KEY_rightcaret = 2982

// KEY_rightdoublequotemark is a representation of the C constant GDK_KEY_rightdoublequotemark.
const KEY_rightdoublequotemark = 2771

// KEY_rightmiddlecurlybrace is a representation of the C constant GDK_KEY_rightmiddlecurlybrace.
const KEY_rightmiddlecurlybrace = 2224

// KEY_rightmiddlesummation is a representation of the C constant GDK_KEY_rightmiddlesummation.
const KEY_rightmiddlesummation = 2231

// KEY_rightopentriangle is a representation of the C constant GDK_KEY_rightopentriangle.
const KEY_rightopentriangle = 2765

// KEY_rightpointer is a representation of the C constant GDK_KEY_rightpointer.
const KEY_rightpointer = 2795

// KEY_rightshoe is a representation of the C constant GDK_KEY_rightshoe.
const KEY_rightshoe = 3032

// KEY_rightsinglequotemark is a representation of the C constant GDK_KEY_rightsinglequotemark.
const KEY_rightsinglequotemark = 2769

// KEY_rightt is a representation of the C constant GDK_KEY_rightt.
const KEY_rightt = 2549

// KEY_righttack is a representation of the C constant GDK_KEY_righttack.
const KEY_righttack = 3068

// KEY_s is a representation of the C constant GDK_KEY_s.
const KEY_s = 115

// KEY_sabovedot is a representation of the C constant GDK_KEY_sabovedot.
const KEY_sabovedot = 16784993

// KEY_sacute is a representation of the C constant GDK_KEY_sacute.
const KEY_sacute = 438

// KEY_scaron is a representation of the C constant GDK_KEY_scaron.
const KEY_scaron = 441

// KEY_scedilla is a representation of the C constant GDK_KEY_scedilla.
const KEY_scedilla = 442

// KEY_schwa is a representation of the C constant GDK_KEY_schwa.
const KEY_schwa = 16777817

// KEY_scircumflex is a representation of the C constant GDK_KEY_scircumflex.
const KEY_scircumflex = 766

// KEY_script_switch is a representation of the C constant GDK_KEY_script_switch.
const KEY_script_switch = 65406

// KEY_seconds is a representation of the C constant GDK_KEY_seconds.
const KEY_seconds = 2775

// KEY_section is a representation of the C constant GDK_KEY_section.
const KEY_section = 167

// KEY_semicolon is a representation of the C constant GDK_KEY_semicolon.
const KEY_semicolon = 59

// KEY_semivoicedsound is a representation of the C constant GDK_KEY_semivoicedsound.
const KEY_semivoicedsound = 1247

// KEY_seveneighths is a representation of the C constant GDK_KEY_seveneighths.
const KEY_seveneighths = 2758

// KEY_sevensubscript is a representation of the C constant GDK_KEY_sevensubscript.
const KEY_sevensubscript = 16785543

// KEY_sevensuperior is a representation of the C constant GDK_KEY_sevensuperior.
const KEY_sevensuperior = 16785527

// KEY_signaturemark is a representation of the C constant GDK_KEY_signaturemark.
const KEY_signaturemark = 2762

// KEY_signifblank is a representation of the C constant GDK_KEY_signifblank.
const KEY_signifblank = 2732

// KEY_similarequal is a representation of the C constant GDK_KEY_similarequal.
const KEY_similarequal = 2249

// KEY_singlelowquotemark is a representation of the C constant GDK_KEY_singlelowquotemark.
const KEY_singlelowquotemark = 2813

// KEY_sixsubscript is a representation of the C constant GDK_KEY_sixsubscript.
const KEY_sixsubscript = 16785542

// KEY_sixsuperior is a representation of the C constant GDK_KEY_sixsuperior.
const KEY_sixsuperior = 16785526

// KEY_slash is a representation of the C constant GDK_KEY_slash.
const KEY_slash = 47

// KEY_soliddiamond is a representation of the C constant GDK_KEY_soliddiamond.
const KEY_soliddiamond = 2528

// KEY_space is a representation of the C constant GDK_KEY_space.
const KEY_space = 32

// KEY_squareroot is a representation of the C constant GDK_KEY_squareroot.
const KEY_squareroot = 16785946

// KEY_ssharp is a representation of the C constant GDK_KEY_ssharp.
const KEY_ssharp = 223

// KEY_sterling is a representation of the C constant GDK_KEY_sterling.
const KEY_sterling = 163

// KEY_stricteq is a representation of the C constant GDK_KEY_stricteq.
const KEY_stricteq = 16786019

// KEY_t is a representation of the C constant GDK_KEY_t.
const KEY_t = 116

// KEY_tabovedot is a representation of the C constant GDK_KEY_tabovedot.
const KEY_tabovedot = 16785003

// KEY_tcaron is a representation of the C constant GDK_KEY_tcaron.
const KEY_tcaron = 443

// KEY_tcedilla is a representation of the C constant GDK_KEY_tcedilla.
const KEY_tcedilla = 510

// KEY_telephone is a representation of the C constant GDK_KEY_telephone.
const KEY_telephone = 2809

// KEY_telephonerecorder is a representation of the C constant GDK_KEY_telephonerecorder.
const KEY_telephonerecorder = 2810

// KEY_therefore is a representation of the C constant GDK_KEY_therefore.
const KEY_therefore = 2240

// KEY_thinspace is a representation of the C constant GDK_KEY_thinspace.
const KEY_thinspace = 2727

// KEY_thorn is a representation of the C constant GDK_KEY_thorn.
const KEY_thorn = 254

// KEY_threeeighths is a representation of the C constant GDK_KEY_threeeighths.
const KEY_threeeighths = 2756

// KEY_threefifths is a representation of the C constant GDK_KEY_threefifths.
const KEY_threefifths = 2740

// KEY_threequarters is a representation of the C constant GDK_KEY_threequarters.
const KEY_threequarters = 190

// KEY_threesubscript is a representation of the C constant GDK_KEY_threesubscript.
const KEY_threesubscript = 16785539

// KEY_threesuperior is a representation of the C constant GDK_KEY_threesuperior.
const KEY_threesuperior = 179

// KEY_tintegral is a representation of the C constant GDK_KEY_tintegral.
const KEY_tintegral = 16785965

// KEY_topintegral is a representation of the C constant GDK_KEY_topintegral.
const KEY_topintegral = 2212

// KEY_topleftparens is a representation of the C constant GDK_KEY_topleftparens.
const KEY_topleftparens = 2219

// KEY_topleftradical is a representation of the C constant GDK_KEY_topleftradical.
const KEY_topleftradical = 2210

// KEY_topleftsqbracket is a representation of the C constant GDK_KEY_topleftsqbracket.
const KEY_topleftsqbracket = 2215

// KEY_topleftsummation is a representation of the C constant GDK_KEY_topleftsummation.
const KEY_topleftsummation = 2225

// KEY_toprightparens is a representation of the C constant GDK_KEY_toprightparens.
const KEY_toprightparens = 2221

// KEY_toprightsqbracket is a representation of the C constant GDK_KEY_toprightsqbracket.
const KEY_toprightsqbracket = 2217

// KEY_toprightsummation is a representation of the C constant GDK_KEY_toprightsummation.
const KEY_toprightsummation = 2229

// KEY_topt is a representation of the C constant GDK_KEY_topt.
const KEY_topt = 2551

// KEY_topvertsummationconnector is a representation of the C constant GDK_KEY_topvertsummationconnector.
const KEY_topvertsummationconnector = 2227

// KEY_trademark is a representation of the C constant GDK_KEY_trademark.
const KEY_trademark = 2761

// KEY_trademarkincircle is a representation of the C constant GDK_KEY_trademarkincircle.
const KEY_trademarkincircle = 2763

// KEY_tslash is a representation of the C constant GDK_KEY_tslash.
const KEY_tslash = 956

// KEY_twofifths is a representation of the C constant GDK_KEY_twofifths.
const KEY_twofifths = 2739

// KEY_twosubscript is a representation of the C constant GDK_KEY_twosubscript.
const KEY_twosubscript = 16785538

// KEY_twosuperior is a representation of the C constant GDK_KEY_twosuperior.
const KEY_twosuperior = 178

// KEY_twothirds is a representation of the C constant GDK_KEY_twothirds.
const KEY_twothirds = 2737

// KEY_u is a representation of the C constant GDK_KEY_u.
const KEY_u = 117

// KEY_uacute is a representation of the C constant GDK_KEY_uacute.
const KEY_uacute = 250

// KEY_ubelowdot is a representation of the C constant GDK_KEY_ubelowdot.
const KEY_ubelowdot = 16785125

// KEY_ubreve is a representation of the C constant GDK_KEY_ubreve.
const KEY_ubreve = 765

// KEY_ucircumflex is a representation of the C constant GDK_KEY_ucircumflex.
const KEY_ucircumflex = 251

// KEY_udiaeresis is a representation of the C constant GDK_KEY_udiaeresis.
const KEY_udiaeresis = 252

// KEY_udoubleacute is a representation of the C constant GDK_KEY_udoubleacute.
const KEY_udoubleacute = 507

// KEY_ugrave is a representation of the C constant GDK_KEY_ugrave.
const KEY_ugrave = 249

// KEY_uhook is a representation of the C constant GDK_KEY_uhook.
const KEY_uhook = 16785127

// KEY_uhorn is a representation of the C constant GDK_KEY_uhorn.
const KEY_uhorn = 16777648

// KEY_uhornacute is a representation of the C constant GDK_KEY_uhornacute.
const KEY_uhornacute = 16785129

// KEY_uhornbelowdot is a representation of the C constant GDK_KEY_uhornbelowdot.
const KEY_uhornbelowdot = 16785137

// KEY_uhorngrave is a representation of the C constant GDK_KEY_uhorngrave.
const KEY_uhorngrave = 16785131

// KEY_uhornhook is a representation of the C constant GDK_KEY_uhornhook.
const KEY_uhornhook = 16785133

// KEY_uhorntilde is a representation of the C constant GDK_KEY_uhorntilde.
const KEY_uhorntilde = 16785135

// KEY_umacron is a representation of the C constant GDK_KEY_umacron.
const KEY_umacron = 1022

// KEY_underbar is a representation of the C constant GDK_KEY_underbar.
const KEY_underbar = 3014

// KEY_underscore is a representation of the C constant GDK_KEY_underscore.
const KEY_underscore = 95

// KEY_union is a representation of the C constant GDK_KEY_union.
const KEY_union = 2269

// KEY_uogonek is a representation of the C constant GDK_KEY_uogonek.
const KEY_uogonek = 1017

// KEY_uparrow is a representation of the C constant GDK_KEY_uparrow.
const KEY_uparrow = 2300

// KEY_upcaret is a representation of the C constant GDK_KEY_upcaret.
const KEY_upcaret = 2985

// KEY_upleftcorner is a representation of the C constant GDK_KEY_upleftcorner.
const KEY_upleftcorner = 2540

// KEY_uprightcorner is a representation of the C constant GDK_KEY_uprightcorner.
const KEY_uprightcorner = 2539

// KEY_upshoe is a representation of the C constant GDK_KEY_upshoe.
const KEY_upshoe = 3011

// KEY_upstile is a representation of the C constant GDK_KEY_upstile.
const KEY_upstile = 3027

// KEY_uptack is a representation of the C constant GDK_KEY_uptack.
const KEY_uptack = 3022

// KEY_uring is a representation of the C constant GDK_KEY_uring.
const KEY_uring = 505

// KEY_utilde is a representation of the C constant GDK_KEY_utilde.
const KEY_utilde = 1021

// KEY_v is a representation of the C constant GDK_KEY_v.
const KEY_v = 118

// KEY_variation is a representation of the C constant GDK_KEY_variation.
const KEY_variation = 2241

// KEY_vertbar is a representation of the C constant GDK_KEY_vertbar.
const KEY_vertbar = 2552

// KEY_vertconnector is a representation of the C constant GDK_KEY_vertconnector.
const KEY_vertconnector = 2214

// KEY_voicedsound is a representation of the C constant GDK_KEY_voicedsound.
const KEY_voicedsound = 1246

// KEY_vt is a representation of the C constant GDK_KEY_vt.
const KEY_vt = 2537

// KEY_w is a representation of the C constant GDK_KEY_w.
const KEY_w = 119

// KEY_wacute is a representation of the C constant GDK_KEY_wacute.
const KEY_wacute = 16785027

// KEY_wcircumflex is a representation of the C constant GDK_KEY_wcircumflex.
const KEY_wcircumflex = 16777589

// KEY_wdiaeresis is a representation of the C constant GDK_KEY_wdiaeresis.
const KEY_wdiaeresis = 16785029

// KEY_wgrave is a representation of the C constant GDK_KEY_wgrave.
const KEY_wgrave = 16785025

// KEY_x is a representation of the C constant GDK_KEY_x.
const KEY_x = 120

// KEY_xabovedot is a representation of the C constant GDK_KEY_xabovedot.
const KEY_xabovedot = 16785035

// KEY_y is a representation of the C constant GDK_KEY_y.
const KEY_y = 121

// KEY_yacute is a representation of the C constant GDK_KEY_yacute.
const KEY_yacute = 253

// KEY_ybelowdot is a representation of the C constant GDK_KEY_ybelowdot.
const KEY_ybelowdot = 16785141

// KEY_ycircumflex is a representation of the C constant GDK_KEY_ycircumflex.
const KEY_ycircumflex = 16777591

// KEY_ydiaeresis is a representation of the C constant GDK_KEY_ydiaeresis.
const KEY_ydiaeresis = 255

// KEY_yen is a representation of the C constant GDK_KEY_yen.
const KEY_yen = 165

// KEY_ygrave is a representation of the C constant GDK_KEY_ygrave.
const KEY_ygrave = 16785139

// KEY_yhook is a representation of the C constant GDK_KEY_yhook.
const KEY_yhook = 16785143

// KEY_ytilde is a representation of the C constant GDK_KEY_ytilde.
const KEY_ytilde = 16785145

// KEY_z is a representation of the C constant GDK_KEY_z.
const KEY_z = 122

// KEY_zabovedot is a representation of the C constant GDK_KEY_zabovedot.
const KEY_zabovedot = 447

// KEY_zacute is a representation of the C constant GDK_KEY_zacute.
const KEY_zacute = 444

// KEY_zcaron is a representation of the C constant GDK_KEY_zcaron.
const KEY_zcaron = 446

// KEY_zerosubscript is a representation of the C constant GDK_KEY_zerosubscript.
const KEY_zerosubscript = 16785536

// KEY_zerosuperior is a representation of the C constant GDK_KEY_zerosuperior.
const KEY_zerosuperior = 16785520

// KEY_zstroke is a representation of the C constant GDK_KEY_zstroke.
const KEY_zstroke = 16777654

// MAJOR_VERSION is a representation of the C constant GDK_MAJOR_VERSION.
const MAJOR_VERSION = 3

// MAX_TIMECOORD_AXES is a representation of the C constant GDK_MAX_TIMECOORD_AXES.
const MAX_TIMECOORD_AXES = 128

// MICRO_VERSION is a representation of the C constant GDK_MICRO_VERSION.
const MICRO_VERSION = 12

// MINOR_VERSION is a representation of the C constant GDK_MINOR_VERSION.
const MINOR_VERSION = 24

// PARENT_RELATIVE is a representation of the C constant GDK_PARENT_RELATIVE.
const PARENT_RELATIVE = 1

// PRIORITY_REDRAW is a representation of the C constant GDK_PRIORITY_REDRAW.
const PRIORITY_REDRAW = 120

// DragAction is a representation of the C bitfield GdkDragAction.
type DragAction int

// DragAction_default is a representation of the C bitfield member GDK_ACTION_DEFAULT.
const DragAction_default = DragAction(1)

// DragAction_copy is a representation of the C bitfield member GDK_ACTION_COPY.
const DragAction_copy = DragAction(2)

// DragAction_move is a representation of the C bitfield member GDK_ACTION_MOVE.
const DragAction_move = DragAction(4)

// DragAction_link is a representation of the C bitfield member GDK_ACTION_LINK.
const DragAction_link = DragAction(8)

// DragAction_private is a representation of the C bitfield member GDK_ACTION_PRIVATE.
const DragAction_private = DragAction(16)

// DragAction_ask is a representation of the C bitfield member GDK_ACTION_ASK.
const DragAction_ask = DragAction(32)

// EventMask is a representation of the C bitfield GdkEventMask.
type EventMask int

// EventMask_exposure_mask is a representation of the C bitfield member GDK_EXPOSURE_MASK.
const EventMask_exposure_mask = EventMask(2)

// EventMask_pointer_motion_mask is a representation of the C bitfield member GDK_POINTER_MOTION_MASK.
const EventMask_pointer_motion_mask = EventMask(4)

// EventMask_pointer_motion_hint_mask is a representation of the C bitfield member GDK_POINTER_MOTION_HINT_MASK.
const EventMask_pointer_motion_hint_mask = EventMask(8)

// EventMask_button_motion_mask is a representation of the C bitfield member GDK_BUTTON_MOTION_MASK.
const EventMask_button_motion_mask = EventMask(16)

// EventMask_button1_motion_mask is a representation of the C bitfield member GDK_BUTTON1_MOTION_MASK.
const EventMask_button1_motion_mask = EventMask(32)

// EventMask_button2_motion_mask is a representation of the C bitfield member GDK_BUTTON2_MOTION_MASK.
const EventMask_button2_motion_mask = EventMask(64)

// EventMask_button3_motion_mask is a representation of the C bitfield member GDK_BUTTON3_MOTION_MASK.
const EventMask_button3_motion_mask = EventMask(128)

// EventMask_button_press_mask is a representation of the C bitfield member GDK_BUTTON_PRESS_MASK.
const EventMask_button_press_mask = EventMask(256)

// EventMask_button_release_mask is a representation of the C bitfield member GDK_BUTTON_RELEASE_MASK.
const EventMask_button_release_mask = EventMask(512)

// EventMask_key_press_mask is a representation of the C bitfield member GDK_KEY_PRESS_MASK.
const EventMask_key_press_mask = EventMask(1024)

// EventMask_key_release_mask is a representation of the C bitfield member GDK_KEY_RELEASE_MASK.
const EventMask_key_release_mask = EventMask(2048)

// EventMask_enter_notify_mask is a representation of the C bitfield member GDK_ENTER_NOTIFY_MASK.
const EventMask_enter_notify_mask = EventMask(4096)

// EventMask_leave_notify_mask is a representation of the C bitfield member GDK_LEAVE_NOTIFY_MASK.
const EventMask_leave_notify_mask = EventMask(8192)

// EventMask_focus_change_mask is a representation of the C bitfield member GDK_FOCUS_CHANGE_MASK.
const EventMask_focus_change_mask = EventMask(16384)

// EventMask_structure_mask is a representation of the C bitfield member GDK_STRUCTURE_MASK.
const EventMask_structure_mask = EventMask(32768)

// EventMask_property_change_mask is a representation of the C bitfield member GDK_PROPERTY_CHANGE_MASK.
const EventMask_property_change_mask = EventMask(65536)

// EventMask_visibility_notify_mask is a representation of the C bitfield member GDK_VISIBILITY_NOTIFY_MASK.
const EventMask_visibility_notify_mask = EventMask(131072)

// EventMask_proximity_in_mask is a representation of the C bitfield member GDK_PROXIMITY_IN_MASK.
const EventMask_proximity_in_mask = EventMask(262144)

// EventMask_proximity_out_mask is a representation of the C bitfield member GDK_PROXIMITY_OUT_MASK.
const EventMask_proximity_out_mask = EventMask(524288)

// EventMask_substructure_mask is a representation of the C bitfield member GDK_SUBSTRUCTURE_MASK.
const EventMask_substructure_mask = EventMask(1048576)

// EventMask_scroll_mask is a representation of the C bitfield member GDK_SCROLL_MASK.
const EventMask_scroll_mask = EventMask(2097152)

// EventMask_touch_mask is a representation of the C bitfield member GDK_TOUCH_MASK.
const EventMask_touch_mask = EventMask(4194304)

// EventMask_smooth_scroll_mask is a representation of the C bitfield member GDK_SMOOTH_SCROLL_MASK.
const EventMask_smooth_scroll_mask = EventMask(8388608)

// EventMask_touchpad_gesture_mask is a representation of the C bitfield member GDK_TOUCHPAD_GESTURE_MASK.
const EventMask_touchpad_gesture_mask = EventMask(16777216)

// EventMask_tablet_pad_mask is a representation of the C bitfield member GDK_TABLET_PAD_MASK.
const EventMask_tablet_pad_mask = EventMask(33554432)

// EventMask_all_events_mask is a representation of the C bitfield member GDK_ALL_EVENTS_MASK.
const EventMask_all_events_mask = EventMask(67108862)

// ModifierType is a representation of the C bitfield GdkModifierType.
type ModifierType int

// ModifierType_shift_mask is a representation of the C bitfield member GDK_SHIFT_MASK.
const ModifierType_shift_mask = ModifierType(1)

// ModifierType_lock_mask is a representation of the C bitfield member GDK_LOCK_MASK.
const ModifierType_lock_mask = ModifierType(2)

// ModifierType_control_mask is a representation of the C bitfield member GDK_CONTROL_MASK.
const ModifierType_control_mask = ModifierType(4)

// ModifierType_mod1_mask is a representation of the C bitfield member GDK_MOD1_MASK.
const ModifierType_mod1_mask = ModifierType(8)

// ModifierType_mod2_mask is a representation of the C bitfield member GDK_MOD2_MASK.
const ModifierType_mod2_mask = ModifierType(16)

// ModifierType_mod3_mask is a representation of the C bitfield member GDK_MOD3_MASK.
const ModifierType_mod3_mask = ModifierType(32)

// ModifierType_mod4_mask is a representation of the C bitfield member GDK_MOD4_MASK.
const ModifierType_mod4_mask = ModifierType(64)

// ModifierType_mod5_mask is a representation of the C bitfield member GDK_MOD5_MASK.
const ModifierType_mod5_mask = ModifierType(128)

// ModifierType_button1_mask is a representation of the C bitfield member GDK_BUTTON1_MASK.
const ModifierType_button1_mask = ModifierType(256)

// ModifierType_button2_mask is a representation of the C bitfield member GDK_BUTTON2_MASK.
const ModifierType_button2_mask = ModifierType(512)

// ModifierType_button3_mask is a representation of the C bitfield member GDK_BUTTON3_MASK.
const ModifierType_button3_mask = ModifierType(1024)

// ModifierType_button4_mask is a representation of the C bitfield member GDK_BUTTON4_MASK.
const ModifierType_button4_mask = ModifierType(2048)

// ModifierType_button5_mask is a representation of the C bitfield member GDK_BUTTON5_MASK.
const ModifierType_button5_mask = ModifierType(4096)

// ModifierType_modifier_reserved_13_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_13_MASK.
const ModifierType_modifier_reserved_13_mask = ModifierType(8192)

// ModifierType_modifier_reserved_14_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_14_MASK.
const ModifierType_modifier_reserved_14_mask = ModifierType(16384)

// ModifierType_modifier_reserved_15_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_15_MASK.
const ModifierType_modifier_reserved_15_mask = ModifierType(32768)

// ModifierType_modifier_reserved_16_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_16_MASK.
const ModifierType_modifier_reserved_16_mask = ModifierType(65536)

// ModifierType_modifier_reserved_17_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_17_MASK.
const ModifierType_modifier_reserved_17_mask = ModifierType(131072)

// ModifierType_modifier_reserved_18_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_18_MASK.
const ModifierType_modifier_reserved_18_mask = ModifierType(262144)

// ModifierType_modifier_reserved_19_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_19_MASK.
const ModifierType_modifier_reserved_19_mask = ModifierType(524288)

// ModifierType_modifier_reserved_20_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_20_MASK.
const ModifierType_modifier_reserved_20_mask = ModifierType(1048576)

// ModifierType_modifier_reserved_21_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_21_MASK.
const ModifierType_modifier_reserved_21_mask = ModifierType(2097152)

// ModifierType_modifier_reserved_22_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_22_MASK.
const ModifierType_modifier_reserved_22_mask = ModifierType(4194304)

// ModifierType_modifier_reserved_23_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_23_MASK.
const ModifierType_modifier_reserved_23_mask = ModifierType(8388608)

// ModifierType_modifier_reserved_24_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_24_MASK.
const ModifierType_modifier_reserved_24_mask = ModifierType(16777216)

// ModifierType_modifier_reserved_25_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_25_MASK.
const ModifierType_modifier_reserved_25_mask = ModifierType(33554432)

// ModifierType_super_mask is a representation of the C bitfield member GDK_SUPER_MASK.
const ModifierType_super_mask = ModifierType(67108864)

// ModifierType_hyper_mask is a representation of the C bitfield member GDK_HYPER_MASK.
const ModifierType_hyper_mask = ModifierType(134217728)

// ModifierType_meta_mask is a representation of the C bitfield member GDK_META_MASK.
const ModifierType_meta_mask = ModifierType(268435456)

// ModifierType_modifier_reserved_29_mask is a representation of the C bitfield member GDK_MODIFIER_RESERVED_29_MASK.
const ModifierType_modifier_reserved_29_mask = ModifierType(536870912)

// ModifierType_release_mask is a representation of the C bitfield member GDK_RELEASE_MASK.
const ModifierType_release_mask = ModifierType(1073741824)

// ModifierType_modifier_mask is a representation of the C bitfield member GDK_MODIFIER_MASK.
const ModifierType_modifier_mask = ModifierType(1543512063)

// WMDecoration is a representation of the C bitfield GdkWMDecoration.
type WMDecoration int

// WMDecoration_all is a representation of the C bitfield member GDK_DECOR_ALL.
const WMDecoration_all = WMDecoration(1)

// WMDecoration_border is a representation of the C bitfield member GDK_DECOR_BORDER.
const WMDecoration_border = WMDecoration(2)

// WMDecoration_resizeh is a representation of the C bitfield member GDK_DECOR_RESIZEH.
const WMDecoration_resizeh = WMDecoration(4)

// WMDecoration_title is a representation of the C bitfield member GDK_DECOR_TITLE.
const WMDecoration_title = WMDecoration(8)

// WMDecoration_menu is a representation of the C bitfield member GDK_DECOR_MENU.
const WMDecoration_menu = WMDecoration(16)

// WMDecoration_minimize is a representation of the C bitfield member GDK_DECOR_MINIMIZE.
const WMDecoration_minimize = WMDecoration(32)

// WMDecoration_maximize is a representation of the C bitfield member GDK_DECOR_MAXIMIZE.
const WMDecoration_maximize = WMDecoration(64)

// WMFunction is a representation of the C bitfield GdkWMFunction.
type WMFunction int

// WMFunction_all is a representation of the C bitfield member GDK_FUNC_ALL.
const WMFunction_all = WMFunction(1)

// WMFunction_resize is a representation of the C bitfield member GDK_FUNC_RESIZE.
const WMFunction_resize = WMFunction(2)

// WMFunction_move is a representation of the C bitfield member GDK_FUNC_MOVE.
const WMFunction_move = WMFunction(4)

// WMFunction_minimize is a representation of the C bitfield member GDK_FUNC_MINIMIZE.
const WMFunction_minimize = WMFunction(8)

// WMFunction_maximize is a representation of the C bitfield member GDK_FUNC_MAXIMIZE.
const WMFunction_maximize = WMFunction(16)

// WMFunction_close is a representation of the C bitfield member GDK_FUNC_CLOSE.
const WMFunction_close = WMFunction(32)

// WindowAttributesType is a representation of the C bitfield GdkWindowAttributesType.
type WindowAttributesType int

// WindowAttributesType_title is a representation of the C bitfield member GDK_WA_TITLE.
const WindowAttributesType_title = WindowAttributesType(2)

// WindowAttributesType_x is a representation of the C bitfield member GDK_WA_X.
const WindowAttributesType_x = WindowAttributesType(4)

// WindowAttributesType_y is a representation of the C bitfield member GDK_WA_Y.
const WindowAttributesType_y = WindowAttributesType(8)

// WindowAttributesType_cursor is a representation of the C bitfield member GDK_WA_CURSOR.
const WindowAttributesType_cursor = WindowAttributesType(16)

// WindowAttributesType_visual is a representation of the C bitfield member GDK_WA_VISUAL.
const WindowAttributesType_visual = WindowAttributesType(32)

// WindowAttributesType_wmclass is a representation of the C bitfield member GDK_WA_WMCLASS.
const WindowAttributesType_wmclass = WindowAttributesType(64)

// WindowAttributesType_noredir is a representation of the C bitfield member GDK_WA_NOREDIR.
const WindowAttributesType_noredir = WindowAttributesType(128)

// WindowAttributesType_type_hint is a representation of the C bitfield member GDK_WA_TYPE_HINT.
const WindowAttributesType_type_hint = WindowAttributesType(256)

// WindowHints is a representation of the C bitfield GdkWindowHints.
type WindowHints int

// WindowHints_pos is a representation of the C bitfield member GDK_HINT_POS.
const WindowHints_pos = WindowHints(1)

// WindowHints_min_size is a representation of the C bitfield member GDK_HINT_MIN_SIZE.
const WindowHints_min_size = WindowHints(2)

// WindowHints_max_size is a representation of the C bitfield member GDK_HINT_MAX_SIZE.
const WindowHints_max_size = WindowHints(4)

// WindowHints_base_size is a representation of the C bitfield member GDK_HINT_BASE_SIZE.
const WindowHints_base_size = WindowHints(8)

// WindowHints_aspect is a representation of the C bitfield member GDK_HINT_ASPECT.
const WindowHints_aspect = WindowHints(16)

// WindowHints_resize_inc is a representation of the C bitfield member GDK_HINT_RESIZE_INC.
const WindowHints_resize_inc = WindowHints(32)

// WindowHints_win_gravity is a representation of the C bitfield member GDK_HINT_WIN_GRAVITY.
const WindowHints_win_gravity = WindowHints(64)

// WindowHints_user_pos is a representation of the C bitfield member GDK_HINT_USER_POS.
const WindowHints_user_pos = WindowHints(128)

// WindowHints_user_size is a representation of the C bitfield member GDK_HINT_USER_SIZE.
const WindowHints_user_size = WindowHints(256)

// WindowState is a representation of the C bitfield GdkWindowState.
type WindowState int

// WindowState_withdrawn is a representation of the C bitfield member GDK_WINDOW_STATE_WITHDRAWN.
const WindowState_withdrawn = WindowState(1)

// WindowState_iconified is a representation of the C bitfield member GDK_WINDOW_STATE_ICONIFIED.
const WindowState_iconified = WindowState(2)

// WindowState_maximized is a representation of the C bitfield member GDK_WINDOW_STATE_MAXIMIZED.
const WindowState_maximized = WindowState(4)

// WindowState_sticky is a representation of the C bitfield member GDK_WINDOW_STATE_STICKY.
const WindowState_sticky = WindowState(8)

// WindowState_fullscreen is a representation of the C bitfield member GDK_WINDOW_STATE_FULLSCREEN.
const WindowState_fullscreen = WindowState(16)

// WindowState_above is a representation of the C bitfield member GDK_WINDOW_STATE_ABOVE.
const WindowState_above = WindowState(32)

// WindowState_below is a representation of the C bitfield member GDK_WINDOW_STATE_BELOW.
const WindowState_below = WindowState(64)

// WindowState_focused is a representation of the C bitfield member GDK_WINDOW_STATE_FOCUSED.
const WindowState_focused = WindowState(128)

// WindowState_tiled is a representation of the C bitfield member GDK_WINDOW_STATE_TILED.
const WindowState_tiled = WindowState(256)

// WindowState_top_tiled is a representation of the C bitfield member GDK_WINDOW_STATE_TOP_TILED.
const WindowState_top_tiled = WindowState(512)

// WindowState_top_resizable is a representation of the C bitfield member GDK_WINDOW_STATE_TOP_RESIZABLE.
const WindowState_top_resizable = WindowState(1024)

// WindowState_right_tiled is a representation of the C bitfield member GDK_WINDOW_STATE_RIGHT_TILED.
const WindowState_right_tiled = WindowState(2048)

// WindowState_right_resizable is a representation of the C bitfield member GDK_WINDOW_STATE_RIGHT_RESIZABLE.
const WindowState_right_resizable = WindowState(4096)

// WindowState_bottom_tiled is a representation of the C bitfield member GDK_WINDOW_STATE_BOTTOM_TILED.
const WindowState_bottom_tiled = WindowState(8192)

// WindowState_bottom_resizable is a representation of the C bitfield member GDK_WINDOW_STATE_BOTTOM_RESIZABLE.
const WindowState_bottom_resizable = WindowState(16384)

// WindowState_left_tiled is a representation of the C bitfield member GDK_WINDOW_STATE_LEFT_TILED.
const WindowState_left_tiled = WindowState(32768)

// WindowState_left_resizable is a representation of the C bitfield member GDK_WINDOW_STATE_LEFT_RESIZABLE.
const WindowState_left_resizable = WindowState(65536)

// AxisUse is a representation of the C enumeration GdkAxisUse.
type AxisUse int

// AxisUse_ignore is a representation of the C enumeration member GDK_AXIS_IGNORE.
const AxisUse_ignore = AxisUse(0)

// AxisUse_x is a representation of the C enumeration member GDK_AXIS_X.
const AxisUse_x = AxisUse(1)

// AxisUse_y is a representation of the C enumeration member GDK_AXIS_Y.
const AxisUse_y = AxisUse(2)

// AxisUse_pressure is a representation of the C enumeration member GDK_AXIS_PRESSURE.
const AxisUse_pressure = AxisUse(3)

// AxisUse_xtilt is a representation of the C enumeration member GDK_AXIS_XTILT.
const AxisUse_xtilt = AxisUse(4)

// AxisUse_ytilt is a representation of the C enumeration member GDK_AXIS_YTILT.
const AxisUse_ytilt = AxisUse(5)

// AxisUse_wheel is a representation of the C enumeration member GDK_AXIS_WHEEL.
const AxisUse_wheel = AxisUse(6)

// AxisUse_distance is a representation of the C enumeration member GDK_AXIS_DISTANCE.
const AxisUse_distance = AxisUse(7)

// AxisUse_rotation is a representation of the C enumeration member GDK_AXIS_ROTATION.
const AxisUse_rotation = AxisUse(8)

// AxisUse_slider is a representation of the C enumeration member GDK_AXIS_SLIDER.
const AxisUse_slider = AxisUse(9)

// AxisUse_last is a representation of the C enumeration member GDK_AXIS_LAST.
const AxisUse_last = AxisUse(10)

// ByteOrder is a representation of the C enumeration GdkByteOrder.
type ByteOrder int

// ByteOrder_lsb_first is a representation of the C enumeration member GDK_LSB_FIRST.
const ByteOrder_lsb_first = ByteOrder(0)

// ByteOrder_msb_first is a representation of the C enumeration member GDK_MSB_FIRST.
const ByteOrder_msb_first = ByteOrder(1)

// CrossingMode is a representation of the C enumeration GdkCrossingMode.
type CrossingMode int

// CrossingMode_normal is a representation of the C enumeration member GDK_CROSSING_NORMAL.
const CrossingMode_normal = CrossingMode(0)

// CrossingMode_grab is a representation of the C enumeration member GDK_CROSSING_GRAB.
const CrossingMode_grab = CrossingMode(1)

// CrossingMode_ungrab is a representation of the C enumeration member GDK_CROSSING_UNGRAB.
const CrossingMode_ungrab = CrossingMode(2)

// CrossingMode_gtk_grab is a representation of the C enumeration member GDK_CROSSING_GTK_GRAB.
const CrossingMode_gtk_grab = CrossingMode(3)

// CrossingMode_gtk_ungrab is a representation of the C enumeration member GDK_CROSSING_GTK_UNGRAB.
const CrossingMode_gtk_ungrab = CrossingMode(4)

// CrossingMode_state_changed is a representation of the C enumeration member GDK_CROSSING_STATE_CHANGED.
const CrossingMode_state_changed = CrossingMode(5)

// CrossingMode_touch_begin is a representation of the C enumeration member GDK_CROSSING_TOUCH_BEGIN.
const CrossingMode_touch_begin = CrossingMode(6)

// CrossingMode_touch_end is a representation of the C enumeration member GDK_CROSSING_TOUCH_END.
const CrossingMode_touch_end = CrossingMode(7)

// CrossingMode_device_switch is a representation of the C enumeration member GDK_CROSSING_DEVICE_SWITCH.
const CrossingMode_device_switch = CrossingMode(8)

// CursorType is a representation of the C enumeration GdkCursorType.
type CursorType int

// CursorType_x_cursor is a representation of the C enumeration member GDK_X_CURSOR.
const CursorType_x_cursor = CursorType(0)

// CursorType_arrow is a representation of the C enumeration member GDK_ARROW.
const CursorType_arrow = CursorType(2)

// CursorType_based_arrow_down is a representation of the C enumeration member GDK_BASED_ARROW_DOWN.
const CursorType_based_arrow_down = CursorType(4)

// CursorType_based_arrow_up is a representation of the C enumeration member GDK_BASED_ARROW_UP.
const CursorType_based_arrow_up = CursorType(6)

// CursorType_boat is a representation of the C enumeration member GDK_BOAT.
const CursorType_boat = CursorType(8)

// CursorType_bogosity is a representation of the C enumeration member GDK_BOGOSITY.
const CursorType_bogosity = CursorType(10)

// CursorType_bottom_left_corner is a representation of the C enumeration member GDK_BOTTOM_LEFT_CORNER.
const CursorType_bottom_left_corner = CursorType(12)

// CursorType_bottom_right_corner is a representation of the C enumeration member GDK_BOTTOM_RIGHT_CORNER.
const CursorType_bottom_right_corner = CursorType(14)

// CursorType_bottom_side is a representation of the C enumeration member GDK_BOTTOM_SIDE.
const CursorType_bottom_side = CursorType(16)

// CursorType_bottom_tee is a representation of the C enumeration member GDK_BOTTOM_TEE.
const CursorType_bottom_tee = CursorType(18)

// CursorType_box_spiral is a representation of the C enumeration member GDK_BOX_SPIRAL.
const CursorType_box_spiral = CursorType(20)

// CursorType_center_ptr is a representation of the C enumeration member GDK_CENTER_PTR.
const CursorType_center_ptr = CursorType(22)

// CursorType_circle is a representation of the C enumeration member GDK_CIRCLE.
const CursorType_circle = CursorType(24)

// CursorType_clock is a representation of the C enumeration member GDK_CLOCK.
const CursorType_clock = CursorType(26)

// CursorType_coffee_mug is a representation of the C enumeration member GDK_COFFEE_MUG.
const CursorType_coffee_mug = CursorType(28)

// CursorType_cross is a representation of the C enumeration member GDK_CROSS.
const CursorType_cross = CursorType(30)

// CursorType_cross_reverse is a representation of the C enumeration member GDK_CROSS_REVERSE.
const CursorType_cross_reverse = CursorType(32)

// CursorType_crosshair is a representation of the C enumeration member GDK_CROSSHAIR.
const CursorType_crosshair = CursorType(34)

// CursorType_diamond_cross is a representation of the C enumeration member GDK_DIAMOND_CROSS.
const CursorType_diamond_cross = CursorType(36)

// CursorType_dot is a representation of the C enumeration member GDK_DOT.
const CursorType_dot = CursorType(38)

// CursorType_dotbox is a representation of the C enumeration member GDK_DOTBOX.
const CursorType_dotbox = CursorType(40)

// CursorType_double_arrow is a representation of the C enumeration member GDK_DOUBLE_ARROW.
const CursorType_double_arrow = CursorType(42)

// CursorType_draft_large is a representation of the C enumeration member GDK_DRAFT_LARGE.
const CursorType_draft_large = CursorType(44)

// CursorType_draft_small is a representation of the C enumeration member GDK_DRAFT_SMALL.
const CursorType_draft_small = CursorType(46)

// CursorType_draped_box is a representation of the C enumeration member GDK_DRAPED_BOX.
const CursorType_draped_box = CursorType(48)

// CursorType_exchange is a representation of the C enumeration member GDK_EXCHANGE.
const CursorType_exchange = CursorType(50)

// CursorType_fleur is a representation of the C enumeration member GDK_FLEUR.
const CursorType_fleur = CursorType(52)

// CursorType_gobbler is a representation of the C enumeration member GDK_GOBBLER.
const CursorType_gobbler = CursorType(54)

// CursorType_gumby is a representation of the C enumeration member GDK_GUMBY.
const CursorType_gumby = CursorType(56)

// CursorType_hand1 is a representation of the C enumeration member GDK_HAND1.
const CursorType_hand1 = CursorType(58)

// CursorType_hand2 is a representation of the C enumeration member GDK_HAND2.
const CursorType_hand2 = CursorType(60)

// CursorType_heart is a representation of the C enumeration member GDK_HEART.
const CursorType_heart = CursorType(62)

// CursorType_icon is a representation of the C enumeration member GDK_ICON.
const CursorType_icon = CursorType(64)

// CursorType_iron_cross is a representation of the C enumeration member GDK_IRON_CROSS.
const CursorType_iron_cross = CursorType(66)

// CursorType_left_ptr is a representation of the C enumeration member GDK_LEFT_PTR.
const CursorType_left_ptr = CursorType(68)

// CursorType_left_side is a representation of the C enumeration member GDK_LEFT_SIDE.
const CursorType_left_side = CursorType(70)

// CursorType_left_tee is a representation of the C enumeration member GDK_LEFT_TEE.
const CursorType_left_tee = CursorType(72)

// CursorType_leftbutton is a representation of the C enumeration member GDK_LEFTBUTTON.
const CursorType_leftbutton = CursorType(74)

// CursorType_ll_angle is a representation of the C enumeration member GDK_LL_ANGLE.
const CursorType_ll_angle = CursorType(76)

// CursorType_lr_angle is a representation of the C enumeration member GDK_LR_ANGLE.
const CursorType_lr_angle = CursorType(78)

// CursorType_man is a representation of the C enumeration member GDK_MAN.
const CursorType_man = CursorType(80)

// CursorType_middlebutton is a representation of the C enumeration member GDK_MIDDLEBUTTON.
const CursorType_middlebutton = CursorType(82)

// CursorType_mouse is a representation of the C enumeration member GDK_MOUSE.
const CursorType_mouse = CursorType(84)

// CursorType_pencil is a representation of the C enumeration member GDK_PENCIL.
const CursorType_pencil = CursorType(86)

// CursorType_pirate is a representation of the C enumeration member GDK_PIRATE.
const CursorType_pirate = CursorType(88)

// CursorType_plus is a representation of the C enumeration member GDK_PLUS.
const CursorType_plus = CursorType(90)

// CursorType_question_arrow is a representation of the C enumeration member GDK_QUESTION_ARROW.
const CursorType_question_arrow = CursorType(92)

// CursorType_right_ptr is a representation of the C enumeration member GDK_RIGHT_PTR.
const CursorType_right_ptr = CursorType(94)

// CursorType_right_side is a representation of the C enumeration member GDK_RIGHT_SIDE.
const CursorType_right_side = CursorType(96)

// CursorType_right_tee is a representation of the C enumeration member GDK_RIGHT_TEE.
const CursorType_right_tee = CursorType(98)

// CursorType_rightbutton is a representation of the C enumeration member GDK_RIGHTBUTTON.
const CursorType_rightbutton = CursorType(100)

// CursorType_rtl_logo is a representation of the C enumeration member GDK_RTL_LOGO.
const CursorType_rtl_logo = CursorType(102)

// CursorType_sailboat is a representation of the C enumeration member GDK_SAILBOAT.
const CursorType_sailboat = CursorType(104)

// CursorType_sb_down_arrow is a representation of the C enumeration member GDK_SB_DOWN_ARROW.
const CursorType_sb_down_arrow = CursorType(106)

// CursorType_sb_h_double_arrow is a representation of the C enumeration member GDK_SB_H_DOUBLE_ARROW.
const CursorType_sb_h_double_arrow = CursorType(108)

// CursorType_sb_left_arrow is a representation of the C enumeration member GDK_SB_LEFT_ARROW.
const CursorType_sb_left_arrow = CursorType(110)

// CursorType_sb_right_arrow is a representation of the C enumeration member GDK_SB_RIGHT_ARROW.
const CursorType_sb_right_arrow = CursorType(112)

// CursorType_sb_up_arrow is a representation of the C enumeration member GDK_SB_UP_ARROW.
const CursorType_sb_up_arrow = CursorType(114)

// CursorType_sb_v_double_arrow is a representation of the C enumeration member GDK_SB_V_DOUBLE_ARROW.
const CursorType_sb_v_double_arrow = CursorType(116)

// CursorType_shuttle is a representation of the C enumeration member GDK_SHUTTLE.
const CursorType_shuttle = CursorType(118)

// CursorType_sizing is a representation of the C enumeration member GDK_SIZING.
const CursorType_sizing = CursorType(120)

// CursorType_spider is a representation of the C enumeration member GDK_SPIDER.
const CursorType_spider = CursorType(122)

// CursorType_spraycan is a representation of the C enumeration member GDK_SPRAYCAN.
const CursorType_spraycan = CursorType(124)

// CursorType_star is a representation of the C enumeration member GDK_STAR.
const CursorType_star = CursorType(126)

// CursorType_target is a representation of the C enumeration member GDK_TARGET.
const CursorType_target = CursorType(128)

// CursorType_tcross is a representation of the C enumeration member GDK_TCROSS.
const CursorType_tcross = CursorType(130)

// CursorType_top_left_arrow is a representation of the C enumeration member GDK_TOP_LEFT_ARROW.
const CursorType_top_left_arrow = CursorType(132)

// CursorType_top_left_corner is a representation of the C enumeration member GDK_TOP_LEFT_CORNER.
const CursorType_top_left_corner = CursorType(134)

// CursorType_top_right_corner is a representation of the C enumeration member GDK_TOP_RIGHT_CORNER.
const CursorType_top_right_corner = CursorType(136)

// CursorType_top_side is a representation of the C enumeration member GDK_TOP_SIDE.
const CursorType_top_side = CursorType(138)

// CursorType_top_tee is a representation of the C enumeration member GDK_TOP_TEE.
const CursorType_top_tee = CursorType(140)

// CursorType_trek is a representation of the C enumeration member GDK_TREK.
const CursorType_trek = CursorType(142)

// CursorType_ul_angle is a representation of the C enumeration member GDK_UL_ANGLE.
const CursorType_ul_angle = CursorType(144)

// CursorType_umbrella is a representation of the C enumeration member GDK_UMBRELLA.
const CursorType_umbrella = CursorType(146)

// CursorType_ur_angle is a representation of the C enumeration member GDK_UR_ANGLE.
const CursorType_ur_angle = CursorType(148)

// CursorType_watch is a representation of the C enumeration member GDK_WATCH.
const CursorType_watch = CursorType(150)

// CursorType_xterm is a representation of the C enumeration member GDK_XTERM.
const CursorType_xterm = CursorType(152)

// CursorType_last_cursor is a representation of the C enumeration member GDK_LAST_CURSOR.
const CursorType_last_cursor = CursorType(153)

// CursorType_blank_cursor is a representation of the C enumeration member GDK_BLANK_CURSOR.
const CursorType_blank_cursor = CursorType(-2)

// CursorType_cursor_is_pixmap is a representation of the C enumeration member GDK_CURSOR_IS_PIXMAP.
const CursorType_cursor_is_pixmap = CursorType(-1)

// DevicePadFeature is a representation of the C enumeration GdkDevicePadFeature.
type DevicePadFeature int

// DevicePadFeature_button is a representation of the C enumeration member GDK_DEVICE_PAD_FEATURE_BUTTON.
const DevicePadFeature_button = DevicePadFeature(0)

// DevicePadFeature_ring is a representation of the C enumeration member GDK_DEVICE_PAD_FEATURE_RING.
const DevicePadFeature_ring = DevicePadFeature(1)

// DevicePadFeature_strip is a representation of the C enumeration member GDK_DEVICE_PAD_FEATURE_STRIP.
const DevicePadFeature_strip = DevicePadFeature(2)

// DeviceType is a representation of the C enumeration GdkDeviceType.
type DeviceType int

// DeviceType_master is a representation of the C enumeration member GDK_DEVICE_TYPE_MASTER.
const DeviceType_master = DeviceType(0)

// DeviceType_slave is a representation of the C enumeration member GDK_DEVICE_TYPE_SLAVE.
const DeviceType_slave = DeviceType(1)

// DeviceType_floating is a representation of the C enumeration member GDK_DEVICE_TYPE_FLOATING.
const DeviceType_floating = DeviceType(2)

// DragProtocol is a representation of the C enumeration GdkDragProtocol.
type DragProtocol int

// DragProtocol_none is a representation of the C enumeration member GDK_DRAG_PROTO_NONE.
const DragProtocol_none = DragProtocol(0)

// DragProtocol_motif is a representation of the C enumeration member GDK_DRAG_PROTO_MOTIF.
const DragProtocol_motif = DragProtocol(1)

// DragProtocol_xdnd is a representation of the C enumeration member GDK_DRAG_PROTO_XDND.
const DragProtocol_xdnd = DragProtocol(2)

// DragProtocol_rootwin is a representation of the C enumeration member GDK_DRAG_PROTO_ROOTWIN.
const DragProtocol_rootwin = DragProtocol(3)

// DragProtocol_win32_dropfiles is a representation of the C enumeration member GDK_DRAG_PROTO_WIN32_DROPFILES.
const DragProtocol_win32_dropfiles = DragProtocol(4)

// DragProtocol_ole2 is a representation of the C enumeration member GDK_DRAG_PROTO_OLE2.
const DragProtocol_ole2 = DragProtocol(5)

// DragProtocol_local is a representation of the C enumeration member GDK_DRAG_PROTO_LOCAL.
const DragProtocol_local = DragProtocol(6)

// DragProtocol_wayland is a representation of the C enumeration member GDK_DRAG_PROTO_WAYLAND.
const DragProtocol_wayland = DragProtocol(7)

// EventType is a representation of the C enumeration GdkEventType.
type EventType int

// EventType_nothing is a representation of the C enumeration member GDK_NOTHING.
const EventType_nothing = EventType(-1)

// EventType_delete is a representation of the C enumeration member GDK_DELETE.
const EventType_delete = EventType(0)

// EventType_destroy is a representation of the C enumeration member GDK_DESTROY.
const EventType_destroy = EventType(1)

// EventType_expose is a representation of the C enumeration member GDK_EXPOSE.
const EventType_expose = EventType(2)

// EventType_motion_notify is a representation of the C enumeration member GDK_MOTION_NOTIFY.
const EventType_motion_notify = EventType(3)

// EventType_button_press is a representation of the C enumeration member GDK_BUTTON_PRESS.
const EventType_button_press = EventType(4)

// EventType_2button_press is a representation of the C enumeration member GDK_2BUTTON_PRESS.
const EventType_2button_press = EventType(5)

// EventType_double_button_press is a representation of the C enumeration member GDK_DOUBLE_BUTTON_PRESS.
const EventType_double_button_press = EventType(5)

// EventType_3button_press is a representation of the C enumeration member GDK_3BUTTON_PRESS.
const EventType_3button_press = EventType(6)

// EventType_triple_button_press is a representation of the C enumeration member GDK_TRIPLE_BUTTON_PRESS.
const EventType_triple_button_press = EventType(6)

// EventType_button_release is a representation of the C enumeration member GDK_BUTTON_RELEASE.
const EventType_button_release = EventType(7)

// EventType_key_press is a representation of the C enumeration member GDK_KEY_PRESS.
const EventType_key_press = EventType(8)

// EventType_key_release is a representation of the C enumeration member GDK_KEY_RELEASE.
const EventType_key_release = EventType(9)

// EventType_enter_notify is a representation of the C enumeration member GDK_ENTER_NOTIFY.
const EventType_enter_notify = EventType(10)

// EventType_leave_notify is a representation of the C enumeration member GDK_LEAVE_NOTIFY.
const EventType_leave_notify = EventType(11)

// EventType_focus_change is a representation of the C enumeration member GDK_FOCUS_CHANGE.
const EventType_focus_change = EventType(12)

// EventType_configure is a representation of the C enumeration member GDK_CONFIGURE.
const EventType_configure = EventType(13)

// EventType_map is a representation of the C enumeration member GDK_MAP.
const EventType_map = EventType(14)

// EventType_unmap is a representation of the C enumeration member GDK_UNMAP.
const EventType_unmap = EventType(15)

// EventType_property_notify is a representation of the C enumeration member GDK_PROPERTY_NOTIFY.
const EventType_property_notify = EventType(16)

// EventType_selection_clear is a representation of the C enumeration member GDK_SELECTION_CLEAR.
const EventType_selection_clear = EventType(17)

// EventType_selection_request is a representation of the C enumeration member GDK_SELECTION_REQUEST.
const EventType_selection_request = EventType(18)

// EventType_selection_notify is a representation of the C enumeration member GDK_SELECTION_NOTIFY.
const EventType_selection_notify = EventType(19)

// EventType_proximity_in is a representation of the C enumeration member GDK_PROXIMITY_IN.
const EventType_proximity_in = EventType(20)

// EventType_proximity_out is a representation of the C enumeration member GDK_PROXIMITY_OUT.
const EventType_proximity_out = EventType(21)

// EventType_drag_enter is a representation of the C enumeration member GDK_DRAG_ENTER.
const EventType_drag_enter = EventType(22)

// EventType_drag_leave is a representation of the C enumeration member GDK_DRAG_LEAVE.
const EventType_drag_leave = EventType(23)

// EventType_drag_motion is a representation of the C enumeration member GDK_DRAG_MOTION.
const EventType_drag_motion = EventType(24)

// EventType_drag_status is a representation of the C enumeration member GDK_DRAG_STATUS.
const EventType_drag_status = EventType(25)

// EventType_drop_start is a representation of the C enumeration member GDK_DROP_START.
const EventType_drop_start = EventType(26)

// EventType_drop_finished is a representation of the C enumeration member GDK_DROP_FINISHED.
const EventType_drop_finished = EventType(27)

// EventType_client_event is a representation of the C enumeration member GDK_CLIENT_EVENT.
const EventType_client_event = EventType(28)

// EventType_visibility_notify is a representation of the C enumeration member GDK_VISIBILITY_NOTIFY.
const EventType_visibility_notify = EventType(29)

// EventType_scroll is a representation of the C enumeration member GDK_SCROLL.
const EventType_scroll = EventType(31)

// EventType_window_state is a representation of the C enumeration member GDK_WINDOW_STATE.
const EventType_window_state = EventType(32)

// EventType_setting is a representation of the C enumeration member GDK_SETTING.
const EventType_setting = EventType(33)

// EventType_owner_change is a representation of the C enumeration member GDK_OWNER_CHANGE.
const EventType_owner_change = EventType(34)

// EventType_grab_broken is a representation of the C enumeration member GDK_GRAB_BROKEN.
const EventType_grab_broken = EventType(35)

// EventType_damage is a representation of the C enumeration member GDK_DAMAGE.
const EventType_damage = EventType(36)

// EventType_touch_begin is a representation of the C enumeration member GDK_TOUCH_BEGIN.
const EventType_touch_begin = EventType(37)

// EventType_touch_update is a representation of the C enumeration member GDK_TOUCH_UPDATE.
const EventType_touch_update = EventType(38)

// EventType_touch_end is a representation of the C enumeration member GDK_TOUCH_END.
const EventType_touch_end = EventType(39)

// EventType_touch_cancel is a representation of the C enumeration member GDK_TOUCH_CANCEL.
const EventType_touch_cancel = EventType(40)

// EventType_touchpad_swipe is a representation of the C enumeration member GDK_TOUCHPAD_SWIPE.
const EventType_touchpad_swipe = EventType(41)

// EventType_touchpad_pinch is a representation of the C enumeration member GDK_TOUCHPAD_PINCH.
const EventType_touchpad_pinch = EventType(42)

// EventType_pad_button_press is a representation of the C enumeration member GDK_PAD_BUTTON_PRESS.
const EventType_pad_button_press = EventType(43)

// EventType_pad_button_release is a representation of the C enumeration member GDK_PAD_BUTTON_RELEASE.
const EventType_pad_button_release = EventType(44)

// EventType_pad_ring is a representation of the C enumeration member GDK_PAD_RING.
const EventType_pad_ring = EventType(45)

// EventType_pad_strip is a representation of the C enumeration member GDK_PAD_STRIP.
const EventType_pad_strip = EventType(46)

// EventType_pad_group_mode is a representation of the C enumeration member GDK_PAD_GROUP_MODE.
const EventType_pad_group_mode = EventType(47)

// EventType_event_last is a representation of the C enumeration member GDK_EVENT_LAST.
const EventType_event_last = EventType(48)

// FilterReturn is a representation of the C enumeration GdkFilterReturn.
type FilterReturn int

// FilterReturn_continue is a representation of the C enumeration member GDK_FILTER_CONTINUE.
const FilterReturn_continue = FilterReturn(0)

// FilterReturn_translate is a representation of the C enumeration member GDK_FILTER_TRANSLATE.
const FilterReturn_translate = FilterReturn(1)

// FilterReturn_remove is a representation of the C enumeration member GDK_FILTER_REMOVE.
const FilterReturn_remove = FilterReturn(2)

// GrabOwnership is a representation of the C enumeration GdkGrabOwnership.
type GrabOwnership int

// GrabOwnership_none is a representation of the C enumeration member GDK_OWNERSHIP_NONE.
const GrabOwnership_none = GrabOwnership(0)

// GrabOwnership_window is a representation of the C enumeration member GDK_OWNERSHIP_WINDOW.
const GrabOwnership_window = GrabOwnership(1)

// GrabOwnership_application is a representation of the C enumeration member GDK_OWNERSHIP_APPLICATION.
const GrabOwnership_application = GrabOwnership(2)

// GrabStatus is a representation of the C enumeration GdkGrabStatus.
type GrabStatus int

// GrabStatus_success is a representation of the C enumeration member GDK_GRAB_SUCCESS.
const GrabStatus_success = GrabStatus(0)

// GrabStatus_already_grabbed is a representation of the C enumeration member GDK_GRAB_ALREADY_GRABBED.
const GrabStatus_already_grabbed = GrabStatus(1)

// GrabStatus_invalid_time is a representation of the C enumeration member GDK_GRAB_INVALID_TIME.
const GrabStatus_invalid_time = GrabStatus(2)

// GrabStatus_not_viewable is a representation of the C enumeration member GDK_GRAB_NOT_VIEWABLE.
const GrabStatus_not_viewable = GrabStatus(3)

// GrabStatus_frozen is a representation of the C enumeration member GDK_GRAB_FROZEN.
const GrabStatus_frozen = GrabStatus(4)

// GrabStatus_failed is a representation of the C enumeration member GDK_GRAB_FAILED.
const GrabStatus_failed = GrabStatus(5)

// Gravity is a representation of the C enumeration GdkGravity.
type Gravity int

// Gravity_north_west is a representation of the C enumeration member GDK_GRAVITY_NORTH_WEST.
const Gravity_north_west = Gravity(1)

// Gravity_north is a representation of the C enumeration member GDK_GRAVITY_NORTH.
const Gravity_north = Gravity(2)

// Gravity_north_east is a representation of the C enumeration member GDK_GRAVITY_NORTH_EAST.
const Gravity_north_east = Gravity(3)

// Gravity_west is a representation of the C enumeration member GDK_GRAVITY_WEST.
const Gravity_west = Gravity(4)

// Gravity_center is a representation of the C enumeration member GDK_GRAVITY_CENTER.
const Gravity_center = Gravity(5)

// Gravity_east is a representation of the C enumeration member GDK_GRAVITY_EAST.
const Gravity_east = Gravity(6)

// Gravity_south_west is a representation of the C enumeration member GDK_GRAVITY_SOUTH_WEST.
const Gravity_south_west = Gravity(7)

// Gravity_south is a representation of the C enumeration member GDK_GRAVITY_SOUTH.
const Gravity_south = Gravity(8)

// Gravity_south_east is a representation of the C enumeration member GDK_GRAVITY_SOUTH_EAST.
const Gravity_south_east = Gravity(9)

// Gravity_static is a representation of the C enumeration member GDK_GRAVITY_STATIC.
const Gravity_static = Gravity(10)

// InputMode is a representation of the C enumeration GdkInputMode.
type InputMode int

// InputMode_disabled is a representation of the C enumeration member GDK_MODE_DISABLED.
const InputMode_disabled = InputMode(0)

// InputMode_screen is a representation of the C enumeration member GDK_MODE_SCREEN.
const InputMode_screen = InputMode(1)

// InputMode_window is a representation of the C enumeration member GDK_MODE_WINDOW.
const InputMode_window = InputMode(2)

// InputSource is a representation of the C enumeration GdkInputSource.
type InputSource int

// InputSource_mouse is a representation of the C enumeration member GDK_SOURCE_MOUSE.
const InputSource_mouse = InputSource(0)

// InputSource_pen is a representation of the C enumeration member GDK_SOURCE_PEN.
const InputSource_pen = InputSource(1)

// InputSource_eraser is a representation of the C enumeration member GDK_SOURCE_ERASER.
const InputSource_eraser = InputSource(2)

// InputSource_cursor is a representation of the C enumeration member GDK_SOURCE_CURSOR.
const InputSource_cursor = InputSource(3)

// InputSource_keyboard is a representation of the C enumeration member GDK_SOURCE_KEYBOARD.
const InputSource_keyboard = InputSource(4)

// InputSource_touchscreen is a representation of the C enumeration member GDK_SOURCE_TOUCHSCREEN.
const InputSource_touchscreen = InputSource(5)

// InputSource_touchpad is a representation of the C enumeration member GDK_SOURCE_TOUCHPAD.
const InputSource_touchpad = InputSource(6)

// InputSource_trackpoint is a representation of the C enumeration member GDK_SOURCE_TRACKPOINT.
const InputSource_trackpoint = InputSource(7)

// InputSource_tablet_pad is a representation of the C enumeration member GDK_SOURCE_TABLET_PAD.
const InputSource_tablet_pad = InputSource(8)

// ModifierIntent is a representation of the C enumeration GdkModifierIntent.
type ModifierIntent int

// ModifierIntent_primary_accelerator is a representation of the C enumeration member GDK_MODIFIER_INTENT_PRIMARY_ACCELERATOR.
const ModifierIntent_primary_accelerator = ModifierIntent(0)

// ModifierIntent_context_menu is a representation of the C enumeration member GDK_MODIFIER_INTENT_CONTEXT_MENU.
const ModifierIntent_context_menu = ModifierIntent(1)

// ModifierIntent_extend_selection is a representation of the C enumeration member GDK_MODIFIER_INTENT_EXTEND_SELECTION.
const ModifierIntent_extend_selection = ModifierIntent(2)

// ModifierIntent_modify_selection is a representation of the C enumeration member GDK_MODIFIER_INTENT_MODIFY_SELECTION.
const ModifierIntent_modify_selection = ModifierIntent(3)

// ModifierIntent_no_text_input is a representation of the C enumeration member GDK_MODIFIER_INTENT_NO_TEXT_INPUT.
const ModifierIntent_no_text_input = ModifierIntent(4)

// ModifierIntent_shift_group is a representation of the C enumeration member GDK_MODIFIER_INTENT_SHIFT_GROUP.
const ModifierIntent_shift_group = ModifierIntent(5)

// ModifierIntent_default_mod_mask is a representation of the C enumeration member GDK_MODIFIER_INTENT_DEFAULT_MOD_MASK.
const ModifierIntent_default_mod_mask = ModifierIntent(6)

// NotifyType is a representation of the C enumeration GdkNotifyType.
type NotifyType int

// NotifyType_ancestor is a representation of the C enumeration member GDK_NOTIFY_ANCESTOR.
const NotifyType_ancestor = NotifyType(0)

// NotifyType_virtual is a representation of the C enumeration member GDK_NOTIFY_VIRTUAL.
const NotifyType_virtual = NotifyType(1)

// NotifyType_inferior is a representation of the C enumeration member GDK_NOTIFY_INFERIOR.
const NotifyType_inferior = NotifyType(2)

// NotifyType_nonlinear is a representation of the C enumeration member GDK_NOTIFY_NONLINEAR.
const NotifyType_nonlinear = NotifyType(3)

// NotifyType_nonlinear_virtual is a representation of the C enumeration member GDK_NOTIFY_NONLINEAR_VIRTUAL.
const NotifyType_nonlinear_virtual = NotifyType(4)

// NotifyType_unknown is a representation of the C enumeration member GDK_NOTIFY_UNKNOWN.
const NotifyType_unknown = NotifyType(5)

// OwnerChange is a representation of the C enumeration GdkOwnerChange.
type OwnerChange int

// OwnerChange_new_owner is a representation of the C enumeration member GDK_OWNER_CHANGE_NEW_OWNER.
const OwnerChange_new_owner = OwnerChange(0)

// OwnerChange_destroy is a representation of the C enumeration member GDK_OWNER_CHANGE_DESTROY.
const OwnerChange_destroy = OwnerChange(1)

// OwnerChange_close is a representation of the C enumeration member GDK_OWNER_CHANGE_CLOSE.
const OwnerChange_close = OwnerChange(2)

// PropMode is a representation of the C enumeration GdkPropMode.
type PropMode int

// PropMode_replace is a representation of the C enumeration member GDK_PROP_MODE_REPLACE.
const PropMode_replace = PropMode(0)

// PropMode_prepend is a representation of the C enumeration member GDK_PROP_MODE_PREPEND.
const PropMode_prepend = PropMode(1)

// PropMode_append is a representation of the C enumeration member GDK_PROP_MODE_APPEND.
const PropMode_append = PropMode(2)

// PropertyState is a representation of the C enumeration GdkPropertyState.
type PropertyState int

// PropertyState_new_value is a representation of the C enumeration member GDK_PROPERTY_NEW_VALUE.
const PropertyState_new_value = PropertyState(0)

// PropertyState_delete is a representation of the C enumeration member GDK_PROPERTY_DELETE.
const PropertyState_delete = PropertyState(1)

// ScrollDirection is a representation of the C enumeration GdkScrollDirection.
type ScrollDirection int

// ScrollDirection_up is a representation of the C enumeration member GDK_SCROLL_UP.
const ScrollDirection_up = ScrollDirection(0)

// ScrollDirection_down is a representation of the C enumeration member GDK_SCROLL_DOWN.
const ScrollDirection_down = ScrollDirection(1)

// ScrollDirection_left is a representation of the C enumeration member GDK_SCROLL_LEFT.
const ScrollDirection_left = ScrollDirection(2)

// ScrollDirection_right is a representation of the C enumeration member GDK_SCROLL_RIGHT.
const ScrollDirection_right = ScrollDirection(3)

// ScrollDirection_smooth is a representation of the C enumeration member GDK_SCROLL_SMOOTH.
const ScrollDirection_smooth = ScrollDirection(4)

// SettingAction is a representation of the C enumeration GdkSettingAction.
type SettingAction int

// SettingAction_new is a representation of the C enumeration member GDK_SETTING_ACTION_NEW.
const SettingAction_new = SettingAction(0)

// SettingAction_changed is a representation of the C enumeration member GDK_SETTING_ACTION_CHANGED.
const SettingAction_changed = SettingAction(1)

// SettingAction_deleted is a representation of the C enumeration member GDK_SETTING_ACTION_DELETED.
const SettingAction_deleted = SettingAction(2)

// Status is a representation of the C enumeration GdkStatus.
type Status int

// Status_ok is a representation of the C enumeration member GDK_OK.
const Status_ok = Status(0)

// Status_error is a representation of the C enumeration member GDK_ERROR.
const Status_error = Status(-1)

// Status_error_param is a representation of the C enumeration member GDK_ERROR_PARAM.
const Status_error_param = Status(-2)

// Status_error_file is a representation of the C enumeration member GDK_ERROR_FILE.
const Status_error_file = Status(-3)

// Status_error_mem is a representation of the C enumeration member GDK_ERROR_MEM.
const Status_error_mem = Status(-4)

// TouchpadGesturePhase is a representation of the C enumeration GdkTouchpadGesturePhase.
type TouchpadGesturePhase int

// TouchpadGesturePhase_begin is a representation of the C enumeration member GDK_TOUCHPAD_GESTURE_PHASE_BEGIN.
const TouchpadGesturePhase_begin = TouchpadGesturePhase(0)

// TouchpadGesturePhase_update is a representation of the C enumeration member GDK_TOUCHPAD_GESTURE_PHASE_UPDATE.
const TouchpadGesturePhase_update = TouchpadGesturePhase(1)

// TouchpadGesturePhase_end is a representation of the C enumeration member GDK_TOUCHPAD_GESTURE_PHASE_END.
const TouchpadGesturePhase_end = TouchpadGesturePhase(2)

// TouchpadGesturePhase_cancel is a representation of the C enumeration member GDK_TOUCHPAD_GESTURE_PHASE_CANCEL.
const TouchpadGesturePhase_cancel = TouchpadGesturePhase(3)

// VisibilityState is a representation of the C enumeration GdkVisibilityState.
type VisibilityState int

// VisibilityState_unobscured is a representation of the C enumeration member GDK_VISIBILITY_UNOBSCURED.
const VisibilityState_unobscured = VisibilityState(0)

// VisibilityState_partial is a representation of the C enumeration member GDK_VISIBILITY_PARTIAL.
const VisibilityState_partial = VisibilityState(1)

// VisibilityState_fully_obscured is a representation of the C enumeration member GDK_VISIBILITY_FULLY_OBSCURED.
const VisibilityState_fully_obscured = VisibilityState(2)

// VisualType is a representation of the C enumeration GdkVisualType.
type VisualType int

// VisualType_static_gray is a representation of the C enumeration member GDK_VISUAL_STATIC_GRAY.
const VisualType_static_gray = VisualType(0)

// VisualType_grayscale is a representation of the C enumeration member GDK_VISUAL_GRAYSCALE.
const VisualType_grayscale = VisualType(1)

// VisualType_static_color is a representation of the C enumeration member GDK_VISUAL_STATIC_COLOR.
const VisualType_static_color = VisualType(2)

// VisualType_pseudo_color is a representation of the C enumeration member GDK_VISUAL_PSEUDO_COLOR.
const VisualType_pseudo_color = VisualType(3)

// VisualType_true_color is a representation of the C enumeration member GDK_VISUAL_TRUE_COLOR.
const VisualType_true_color = VisualType(4)

// VisualType_direct_color is a representation of the C enumeration member GDK_VISUAL_DIRECT_COLOR.
const VisualType_direct_color = VisualType(5)

// WindowEdge is a representation of the C enumeration GdkWindowEdge.
type WindowEdge int

// WindowEdge_north_west is a representation of the C enumeration member GDK_WINDOW_EDGE_NORTH_WEST.
const WindowEdge_north_west = WindowEdge(0)

// WindowEdge_north is a representation of the C enumeration member GDK_WINDOW_EDGE_NORTH.
const WindowEdge_north = WindowEdge(1)

// WindowEdge_north_east is a representation of the C enumeration member GDK_WINDOW_EDGE_NORTH_EAST.
const WindowEdge_north_east = WindowEdge(2)

// WindowEdge_west is a representation of the C enumeration member GDK_WINDOW_EDGE_WEST.
const WindowEdge_west = WindowEdge(3)

// WindowEdge_east is a representation of the C enumeration member GDK_WINDOW_EDGE_EAST.
const WindowEdge_east = WindowEdge(4)

// WindowEdge_south_west is a representation of the C enumeration member GDK_WINDOW_EDGE_SOUTH_WEST.
const WindowEdge_south_west = WindowEdge(5)

// WindowEdge_south is a representation of the C enumeration member GDK_WINDOW_EDGE_SOUTH.
const WindowEdge_south = WindowEdge(6)

// WindowEdge_south_east is a representation of the C enumeration member GDK_WINDOW_EDGE_SOUTH_EAST.
const WindowEdge_south_east = WindowEdge(7)

// WindowType is a representation of the C enumeration GdkWindowType.
type WindowType int

// WindowType_root is a representation of the C enumeration member GDK_WINDOW_ROOT.
const WindowType_root = WindowType(0)

// WindowType_toplevel is a representation of the C enumeration member GDK_WINDOW_TOPLEVEL.
const WindowType_toplevel = WindowType(1)

// WindowType_child is a representation of the C enumeration member GDK_WINDOW_CHILD.
const WindowType_child = WindowType(2)

// WindowType_temp is a representation of the C enumeration member GDK_WINDOW_TEMP.
const WindowType_temp = WindowType(3)

// WindowType_foreign is a representation of the C enumeration member GDK_WINDOW_FOREIGN.
const WindowType_foreign = WindowType(4)

// WindowType_offscreen is a representation of the C enumeration member GDK_WINDOW_OFFSCREEN.
const WindowType_offscreen = WindowType(5)

// WindowType_subsurface is a representation of the C enumeration member GDK_WINDOW_SUBSURFACE.
const WindowType_subsurface = WindowType(6)

// WindowTypeHint is a representation of the C enumeration GdkWindowTypeHint.
type WindowTypeHint int

// WindowTypeHint_normal is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_NORMAL.
const WindowTypeHint_normal = WindowTypeHint(0)

// WindowTypeHint_dialog is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_DIALOG.
const WindowTypeHint_dialog = WindowTypeHint(1)

// WindowTypeHint_menu is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_MENU.
const WindowTypeHint_menu = WindowTypeHint(2)

// WindowTypeHint_toolbar is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_TOOLBAR.
const WindowTypeHint_toolbar = WindowTypeHint(3)

// WindowTypeHint_splashscreen is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_SPLASHSCREEN.
const WindowTypeHint_splashscreen = WindowTypeHint(4)

// WindowTypeHint_utility is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_UTILITY.
const WindowTypeHint_utility = WindowTypeHint(5)

// WindowTypeHint_dock is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_DOCK.
const WindowTypeHint_dock = WindowTypeHint(6)

// WindowTypeHint_desktop is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_DESKTOP.
const WindowTypeHint_desktop = WindowTypeHint(7)

// WindowTypeHint_dropdown_menu is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU.
const WindowTypeHint_dropdown_menu = WindowTypeHint(8)

// WindowTypeHint_popup_menu is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_POPUP_MENU.
const WindowTypeHint_popup_menu = WindowTypeHint(9)

// WindowTypeHint_tooltip is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_TOOLTIP.
const WindowTypeHint_tooltip = WindowTypeHint(10)

// WindowTypeHint_notification is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_NOTIFICATION.
const WindowTypeHint_notification = WindowTypeHint(11)

// WindowTypeHint_combo is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_COMBO.
const WindowTypeHint_combo = WindowTypeHint(12)

// WindowTypeHint_dnd is a representation of the C enumeration member GDK_WINDOW_TYPE_HINT_DND.
const WindowTypeHint_dnd = WindowTypeHint(13)

// WindowWindowClass is a representation of the C enumeration GdkWindowWindowClass.
type WindowWindowClass int

// WindowWindowClass_input_output is a representation of the C enumeration member GDK_INPUT_OUTPUT.
const WindowWindowClass_input_output = WindowWindowClass(0)

// WindowWindowClass_input_only is a representation of the C enumeration member GDK_INPUT_ONLY.
const WindowWindowClass_input_only = WindowWindowClass(1)

// AddOptionEntriesLibgtkOnly is analogous to the C function gdk_add_option_entries_libgtk_only.
func AddOptionEntriesLibgtkOnly(group *glib.OptionGroup) {
	sys_group := group.ToC()
}

// Beep is analogous to the C function gdk_beep.
func Beep() {}

// CairoCreate is analogous to the C function gdk_cairo_create.
func CairoCreate(window *Window) {
	sys_window := window.ToC()
}

// CairoGetClipRectangle is analogous to the C function gdk_cairo_get_clip_rectangle.
func CairoGetClipRectangle(cr *cairo.Context) {
	sys_cr := cr.ToC()
}

// CairoRectangle is analogous to the C function gdk_cairo_rectangle.
func CairoRectangle(cr *cairo.Context, rectangle *Rectangle) {
	sys_cr := cr.ToC()
	sys_rectangle := rectangle.ToC()
}

// CairoRegion is analogous to the C function gdk_cairo_region.
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	sys_cr := cr.ToC()
	sys_region := region.ToC()
}

// CairoRegionCreateFromSurface is analogous to the C function gdk_cairo_region_create_from_surface.
func CairoRegionCreateFromSurface(surface *cairo.Surface) {
	sys_surface := surface.ToC()
}

// CairoSetSourceColor is analogous to the C function gdk_cairo_set_source_color.
func CairoSetSourceColor(cr *cairo.Context, color *Color) {
	sys_cr := cr.ToC()
	sys_color := color.ToC()
}

// CairoSetSourcePixbuf is analogous to the C function gdk_cairo_set_source_pixbuf.
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, pixbufX float64, pixbufY float64) {
	sys_cr := cr.ToC()
	sys_pixbuf := pixbuf.ToC()
	sys_pixbufX := pixbufX
	sys_pixbufY := pixbufY
}

// CairoSetSourceRgba is analogous to the C function gdk_cairo_set_source_rgba.
func CairoSetSourceRgba(cr *cairo.Context, rgba *RGBA) {
	sys_cr := cr.ToC()
	sys_rgba := rgba.ToC()
}

// CairoSetSourceWindow is analogous to the C function gdk_cairo_set_source_window.
func CairoSetSourceWindow(cr *cairo.Context, window *Window, x float64, y float64) {
	sys_cr := cr.ToC()
	sys_window := window.ToC()
	sys_x := x
	sys_y := y
}

// DisableMultidevice is analogous to the C function gdk_disable_multidevice.
func DisableMultidevice() {}

// DragAbort is analogous to the C function gdk_drag_abort.
func DragAbort(context *DragContext, time uint32) {
	sys_context := context.ToC()
	sys_time := time
}

// DragBegin is analogous to the C function gdk_drag_begin.
func DragBegin(window *Window, targets *glib.List) {
	sys_window := window.ToC()
	sys_targets := targets.ToC()
}

// DragBeginForDevice is analogous to the C function gdk_drag_begin_for_device.
func DragBeginForDevice(window *Window, device *Device, targets *glib.List) {
	sys_window := window.ToC()
	sys_device := device.ToC()
	sys_targets := targets.ToC()
}

// DragDrop is analogous to the C function gdk_drag_drop.
func DragDrop(context *DragContext, time uint32) {
	sys_context := context.ToC()
	sys_time := time
}

// DragDropSucceeded is analogous to the C function gdk_drag_drop_succeeded.
func DragDropSucceeded(context *DragContext) {
	sys_context := context.ToC()
}

// DragFindWindowForScreen is analogous to the C function gdk_drag_find_window_for_screen.
func DragFindWindowForScreen(context *DragContext, dragWindow *Window, screen *Screen, xRoot int, yRoot int) {
	sys_context := context.ToC()
	sys_dragWindow := dragWindow.ToC()
	sys_screen := screen.ToC()
	sys_xRoot := xRoot
	sys_yRoot := yRoot
}

// DragGetSelection is analogous to the C function gdk_drag_get_selection.
func DragGetSelection(context *DragContext) {
	sys_context := context.ToC()
}

// DragMotion is analogous to the C function gdk_drag_motion.
func DragMotion(context *DragContext, destWindow *Window, protocol int, xRoot int, yRoot int, suggestedAction int, possibleActions int, time uint32) {
	sys_context := context.ToC()
	sys_destWindow := destWindow.ToC()
	sys_protocol := protocol
	sys_xRoot := xRoot
	sys_yRoot := yRoot
	sys_suggestedAction := suggestedAction
	sys_possibleActions := possibleActions
	sys_time := time
}

// DragStatus is analogous to the C function gdk_drag_status.
func DragStatus(context *DragContext, action int, time uint32) {
	sys_context := context.ToC()
	sys_action := action
	sys_time := time
}

// DropFinish is analogous to the C function gdk_drop_finish.
func DropFinish(context *DragContext, success bool, time uint32) {
	sys_context := context.ToC()
	sys_success := success
	sys_time := time
}

// DropReply is analogous to the C function gdk_drop_reply.
func DropReply(context *DragContext, accepted bool, time uint32) {
	sys_context := context.ToC()
	sys_accepted := accepted
	sys_time := time
}

// ErrorTrapPop is analogous to the C function gdk_error_trap_pop.
func ErrorTrapPop() {}

// ErrorTrapPopIgnored is analogous to the C function gdk_error_trap_pop_ignored.
func ErrorTrapPopIgnored() {}

// ErrorTrapPush is analogous to the C function gdk_error_trap_push.
func ErrorTrapPush() {}

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

// EventsGetAngle is analogous to the C function gdk_events_get_angle.
func EventsGetAngle(event1 *Event, event2 *Event) {
	sys_event1 := event1.ToC()
	sys_event2 := event2.ToC()
}

// EventsGetCenter is analogous to the C function gdk_events_get_center.
func EventsGetCenter(event1 *Event, event2 *Event) {
	sys_event1 := event1.ToC()
	sys_event2 := event2.ToC()
}

// EventsGetDistance is analogous to the C function gdk_events_get_distance.
func EventsGetDistance(event1 *Event, event2 *Event) {
	sys_event1 := event1.ToC()
	sys_event2 := event2.ToC()
}

// EventsPending is analogous to the C function gdk_events_pending.
func EventsPending() {}

// Flush is analogous to the C function gdk_flush.
func Flush() {}

// GetDefaultRootWindow is analogous to the C function gdk_get_default_root_window.
func GetDefaultRootWindow() {}

// GetDisplay is analogous to the C function gdk_get_display.
func GetDisplay() {}

// GetDisplayArgName is analogous to the C function gdk_get_display_arg_name.
func GetDisplayArgName() {}

// GetProgramClass is analogous to the C function gdk_get_program_class.
func GetProgramClass() {}

// GetShowEvents is analogous to the C function gdk_get_show_events.
func GetShowEvents() {}

// UNSUPPORTED : gdk_init : has array param, argv

// UNSUPPORTED : gdk_init_check : has array param, argv

// KeyboardGrab is analogous to the C function gdk_keyboard_grab.
func KeyboardGrab(window *Window, ownerEvents bool, time uint32) {
	sys_window := window.ToC()
	sys_ownerEvents := ownerEvents
	sys_time := time
}

// KeyboardUngrab is analogous to the C function gdk_keyboard_ungrab.
func KeyboardUngrab(time uint32) {
	sys_time := time
}

// KeyvalConvertCase is analogous to the C function gdk_keyval_convert_case.
func KeyvalConvertCase(symbol uint) {
	sys_symbol := symbol
}

// KeyvalFromName is analogous to the C function gdk_keyval_from_name.
func KeyvalFromName(keyvalName string) {
	sys_keyvalName := keyvalName
}

// KeyvalIsLower is analogous to the C function gdk_keyval_is_lower.
func KeyvalIsLower(keyval uint) {
	sys_keyval := keyval
}

// KeyvalIsUpper is analogous to the C function gdk_keyval_is_upper.
func KeyvalIsUpper(keyval uint) {
	sys_keyval := keyval
}

// KeyvalName is analogous to the C function gdk_keyval_name.
func KeyvalName(keyval uint) {
	sys_keyval := keyval
}

// KeyvalToLower is analogous to the C function gdk_keyval_to_lower.
func KeyvalToLower(keyval uint) {
	sys_keyval := keyval
}

// KeyvalToUnicode is analogous to the C function gdk_keyval_to_unicode.
func KeyvalToUnicode(keyval uint) {
	sys_keyval := keyval
}

// KeyvalToUpper is analogous to the C function gdk_keyval_to_upper.
func KeyvalToUpper(keyval uint) {
	sys_keyval := keyval
}

// ListVisuals is analogous to the C function gdk_list_visuals.
func ListVisuals() {}

// NotifyStartupComplete is analogous to the C function gdk_notify_startup_complete.
func NotifyStartupComplete() {}

// NotifyStartupCompleteWithId is analogous to the C function gdk_notify_startup_complete_with_id.
func NotifyStartupCompleteWithId(startupId string) {
	sys_startupId := startupId
}

// OffscreenWindowGetEmbedder is analogous to the C function gdk_offscreen_window_get_embedder.
func OffscreenWindowGetEmbedder(window *Window) {
	sys_window := window.ToC()
}

// OffscreenWindowGetSurface is analogous to the C function gdk_offscreen_window_get_surface.
func OffscreenWindowGetSurface(window *Window) {
	sys_window := window.ToC()
}

// OffscreenWindowSetEmbedder is analogous to the C function gdk_offscreen_window_set_embedder.
func OffscreenWindowSetEmbedder(window *Window, embedder *Window) {
	sys_window := window.ToC()
	sys_embedder := embedder.ToC()
}

// PangoContextGet is analogous to the C function gdk_pango_context_get.
func PangoContextGet() {}

// PangoContextGetForScreen is analogous to the C function gdk_pango_context_get_for_screen.
func PangoContextGetForScreen(screen *Screen) {
	sys_screen := screen.ToC()
}

// PangoLayoutGetClipRegion is analogous to the C function gdk_pango_layout_get_clip_region.
func PangoLayoutGetClipRegion(layout *pango.Layout, xOrigin int, yOrigin int, indexRanges *int, nRanges int) {
	sys_layout := layout.ToC()
	sys_xOrigin := xOrigin
	sys_yOrigin := yOrigin
	sys_indexRanges := indexRanges
	sys_nRanges := nRanges
}

// UNSUPPORTED : gdk_pango_layout_line_get_clip_region : parameter 'index_ranges' is array parameter without length parameter

// UNSUPPORTED : gdk_parse_args : has array param, argv

// PixbufGetFromSurface is analogous to the C function gdk_pixbuf_get_from_surface.
func PixbufGetFromSurface(surface *cairo.Surface, srcX int, srcY int, width int, height int) {
	sys_surface := surface.ToC()
	sys_srcX := srcX
	sys_srcY := srcY
	sys_width := width
	sys_height := height
}

// PixbufGetFromWindow is analogous to the C function gdk_pixbuf_get_from_window.
func PixbufGetFromWindow(window *Window, srcX int, srcY int, width int, height int) {
	sys_window := window.ToC()
	sys_srcX := srcX
	sys_srcY := srcY
	sys_width := width
	sys_height := height
}

// PointerGrab is analogous to the C function gdk_pointer_grab.
func PointerGrab(window *Window, ownerEvents bool, eventMask int, confineTo *Window, cursor *Cursor, time uint32) {
	sys_window := window.ToC()
	sys_ownerEvents := ownerEvents
	sys_eventMask := eventMask
	sys_confineTo := confineTo.ToC()
	sys_cursor := cursor.ToC()
	sys_time := time
}

// PointerIsGrabbed is analogous to the C function gdk_pointer_is_grabbed.
func PointerIsGrabbed() {}

// PointerUngrab is analogous to the C function gdk_pointer_ungrab.
func PointerUngrab(time uint32) {
	sys_time := time
}

// PreParseLibgtkOnly is analogous to the C function gdk_pre_parse_libgtk_only.
func PreParseLibgtkOnly() {}

// PropertyChange is analogous to the C function gdk_property_change.
func PropertyChange(window *Window, property Atom, type_ Atom, format int, mode int, data *uint8, nelements int) {
	sys_window := window.ToC()
	sys_property := property.ToC()
	sys_type_ := type_.ToC()
	sys_format := format
	sys_mode := mode
	sys_data := data
	sys_nelements := nelements
}

// PropertyDelete is analogous to the C function gdk_property_delete.
func PropertyDelete(window *Window, property Atom) {
	sys_window := window.ToC()
	sys_property := property.ToC()
}

// UNSUPPORTED : gdk_property_get : has array param, data

// UNSUPPORTED : gdk_query_depths : has array param, depths

// UNSUPPORTED : gdk_query_visual_types : has array param, visual_types

// SelectionConvert is analogous to the C function gdk_selection_convert.
func SelectionConvert(requestor *Window, selection Atom, target Atom, time uint32) {
	sys_requestor := requestor.ToC()
	sys_selection := selection.ToC()
	sys_target := target.ToC()
	sys_time := time
}

// SelectionOwnerGet is analogous to the C function gdk_selection_owner_get.
func SelectionOwnerGet(selection Atom) {
	sys_selection := selection.ToC()
}

// SelectionOwnerGetForDisplay is analogous to the C function gdk_selection_owner_get_for_display.
func SelectionOwnerGetForDisplay(display *Display, selection Atom) {
	sys_display := display.ToC()
	sys_selection := selection.ToC()
}

// SelectionOwnerSet is analogous to the C function gdk_selection_owner_set.
func SelectionOwnerSet(owner *Window, selection Atom, time uint32, sendEvent bool) {
	sys_owner := owner.ToC()
	sys_selection := selection.ToC()
	sys_time := time
	sys_sendEvent := sendEvent
}

// SelectionOwnerSetForDisplay is analogous to the C function gdk_selection_owner_set_for_display.
func SelectionOwnerSetForDisplay(display *Display, owner *Window, selection Atom, time uint32, sendEvent bool) {
	sys_display := display.ToC()
	sys_owner := owner.ToC()
	sys_selection := selection.ToC()
	sys_time := time
	sys_sendEvent := sendEvent
}

// SelectionPropertyGet is analogous to the C function gdk_selection_property_get.
func SelectionPropertyGet(requestor *Window, data **uint8, propType *Atom, propFormat *int) {
	sys_requestor := requestor.ToC()
	sys_data := data
	sys_propType := propType.ToC()
	sys_propFormat := propFormat
}

// SelectionSendNotify is analogous to the C function gdk_selection_send_notify.
func SelectionSendNotify(requestor *Window, selection Atom, target Atom, property Atom, time uint32) {
	sys_requestor := requestor.ToC()
	sys_selection := selection.ToC()
	sys_target := target.ToC()
	sys_property := property.ToC()
	sys_time := time
}

// SelectionSendNotifyForDisplay is analogous to the C function gdk_selection_send_notify_for_display.
func SelectionSendNotifyForDisplay(display *Display, requestor *Window, selection Atom, target Atom, property Atom, time uint32) {
	sys_display := display.ToC()
	sys_requestor := requestor.ToC()
	sys_selection := selection.ToC()
	sys_target := target.ToC()
	sys_property := property.ToC()
	sys_time := time
}

// SetDoubleClickTime is analogous to the C function gdk_set_double_click_time.
func SetDoubleClickTime(msec uint) {
	sys_msec := msec
}

// SetProgramClass is analogous to the C function gdk_set_program_class.
func SetProgramClass(programClass string) {
	sys_programClass := programClass
}

// SetShowEvents is analogous to the C function gdk_set_show_events.
func SetShowEvents(showEvents bool) {
	sys_showEvents := showEvents
}

// SettingGet is analogous to the C function gdk_setting_get.
func SettingGet(name string, value *gobject.Value) {
	sys_name := name
	sys_value := value.ToC()
}

// UNSUPPORTED : gdk_synthesize_window_state : blacklisted

// TestRenderSync is analogous to the C function gdk_test_render_sync.
func TestRenderSync(window *Window) {
	sys_window := window.ToC()
}

// TestSimulateButton is analogous to the C function gdk_test_simulate_button.
func TestSimulateButton(window *Window, x int, y int, button uint, modifiers int, buttonPressrelease int) {
	sys_window := window.ToC()
	sys_x := x
	sys_y := y
	sys_button := button
	sys_modifiers := modifiers
	sys_buttonPressrelease := buttonPressrelease
}

// TestSimulateKey is analogous to the C function gdk_test_simulate_key.
func TestSimulateKey(window *Window, x int, y int, keyval uint, modifiers int, keyPressrelease int) {
	sys_window := window.ToC()
	sys_x := x
	sys_y := y
	sys_keyval := keyval
	sys_modifiers := modifiers
	sys_keyPressrelease := keyPressrelease
}

// UNSUPPORTED : gdk_text_property_to_utf8_list_for_display : parameter 'list' is array parameter without length parameter

// UNSUPPORTED : gdk_threads_add_idle : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_idle_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds_full : parameter 'function' is callback

// ThreadsEnter is analogous to the C function gdk_threads_enter.
func ThreadsEnter() {}

// ThreadsInit is analogous to the C function gdk_threads_init.
func ThreadsInit() {}

// ThreadsLeave is analogous to the C function gdk_threads_leave.
func ThreadsLeave() {}

// UNSUPPORTED : gdk_threads_set_lock_functions : parameter 'enter_fn' is callback

// UnicodeToKeyval is analogous to the C function gdk_unicode_to_keyval.
func UnicodeToKeyval(wc uint32) {
	sys_wc := wc
}

// Utf8ToStringTarget is analogous to the C function gdk_utf8_to_string_target.
func Utf8ToStringTarget(str string) {
	sys_str := str
}

// Atom is a representation of the C record GdkAtom.
type Atom struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkAtom that represents the Atom.
func (recv *Atom) ToC() unsafe.Pointer {
	return recv.native
}

// Color is a representation of the C record GdkColor.
type Color struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkColor that represents the Color.
func (recv *Color) ToC() unsafe.Pointer {
	return recv.native
}

// DevicePadInterface is a representation of the C record GdkDevicePadInterface.
type DevicePadInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkDevicePadInterface that represents the DevicePadInterface.
func (recv *DevicePadInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DrawingContextClass is a representation of the C record GdkDrawingContextClass.
type DrawingContextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkDrawingContextClass that represents the DrawingContextClass.
func (recv *DrawingContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// EventAny is a representation of the C record GdkEventAny.
type EventAny struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventAny that represents the EventAny.
func (recv *EventAny) ToC() unsafe.Pointer {
	return recv.native
}

// EventButton is a representation of the C record GdkEventButton.
type EventButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventButton that represents the EventButton.
func (recv *EventButton) ToC() unsafe.Pointer {
	return recv.native
}

// EventConfigure is a representation of the C record GdkEventConfigure.
type EventConfigure struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventConfigure that represents the EventConfigure.
func (recv *EventConfigure) ToC() unsafe.Pointer {
	return recv.native
}

// EventCrossing is a representation of the C record GdkEventCrossing.
type EventCrossing struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventCrossing that represents the EventCrossing.
func (recv *EventCrossing) ToC() unsafe.Pointer {
	return recv.native
}

// EventDND is a representation of the C record GdkEventDND.
type EventDND struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventDND that represents the EventDND.
func (recv *EventDND) ToC() unsafe.Pointer {
	return recv.native
}

// EventExpose is a representation of the C record GdkEventExpose.
type EventExpose struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventExpose that represents the EventExpose.
func (recv *EventExpose) ToC() unsafe.Pointer {
	return recv.native
}

// EventFocus is a representation of the C record GdkEventFocus.
type EventFocus struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventFocus that represents the EventFocus.
func (recv *EventFocus) ToC() unsafe.Pointer {
	return recv.native
}

// EventGrabBroken is a representation of the C record GdkEventGrabBroken.
//
// since 2.8
type EventGrabBroken struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventGrabBroken that represents the EventGrabBroken.
func (recv *EventGrabBroken) ToC() unsafe.Pointer {
	return recv.native
}

// EventKey is a representation of the C record GdkEventKey.
type EventKey struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventKey that represents the EventKey.
func (recv *EventKey) ToC() unsafe.Pointer {
	return recv.native
}

// EventMotion is a representation of the C record GdkEventMotion.
type EventMotion struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventMotion that represents the EventMotion.
func (recv *EventMotion) ToC() unsafe.Pointer {
	return recv.native
}

// EventOwnerChange is a representation of the C record GdkEventOwnerChange.
//
// since 2.6
type EventOwnerChange struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventOwnerChange that represents the EventOwnerChange.
func (recv *EventOwnerChange) ToC() unsafe.Pointer {
	return recv.native
}

// EventProperty is a representation of the C record GdkEventProperty.
type EventProperty struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventProperty that represents the EventProperty.
func (recv *EventProperty) ToC() unsafe.Pointer {
	return recv.native
}

// EventProximity is a representation of the C record GdkEventProximity.
type EventProximity struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventProximity that represents the EventProximity.
func (recv *EventProximity) ToC() unsafe.Pointer {
	return recv.native
}

// EventScroll is a representation of the C record GdkEventScroll.
type EventScroll struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventScroll that represents the EventScroll.
func (recv *EventScroll) ToC() unsafe.Pointer {
	return recv.native
}

// EventSelection is a representation of the C record GdkEventSelection.
type EventSelection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventSelection that represents the EventSelection.
func (recv *EventSelection) ToC() unsafe.Pointer {
	return recv.native
}

// EventSequence is a representation of the C record GdkEventSequence.
type EventSequence struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventSequence that represents the EventSequence.
func (recv *EventSequence) ToC() unsafe.Pointer {
	return recv.native
}

// EventSetting is a representation of the C record GdkEventSetting.
type EventSetting struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventSetting that represents the EventSetting.
func (recv *EventSetting) ToC() unsafe.Pointer {
	return recv.native
}

// EventTouch is a representation of the C record GdkEventTouch.
type EventTouch struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventTouch that represents the EventTouch.
func (recv *EventTouch) ToC() unsafe.Pointer {
	return recv.native
}

// EventTouchpadPinch is a representation of the C record GdkEventTouchpadPinch.
type EventTouchpadPinch struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventTouchpadPinch that represents the EventTouchpadPinch.
func (recv *EventTouchpadPinch) ToC() unsafe.Pointer {
	return recv.native
}

// EventTouchpadSwipe is a representation of the C record GdkEventTouchpadSwipe.
type EventTouchpadSwipe struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventTouchpadSwipe that represents the EventTouchpadSwipe.
func (recv *EventTouchpadSwipe) ToC() unsafe.Pointer {
	return recv.native
}

// EventVisibility is a representation of the C record GdkEventVisibility.
type EventVisibility struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventVisibility that represents the EventVisibility.
func (recv *EventVisibility) ToC() unsafe.Pointer {
	return recv.native
}

// EventWindowState is a representation of the C record GdkEventWindowState.
type EventWindowState struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventWindowState that represents the EventWindowState.
func (recv *EventWindowState) ToC() unsafe.Pointer {
	return recv.native
}

// FrameClockClass is a representation of the C record GdkFrameClockClass.
type FrameClockClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkFrameClockClass that represents the FrameClockClass.
func (recv *FrameClockClass) ToC() unsafe.Pointer {
	return recv.native
}

// FrameClockPrivate is a representation of the C record GdkFrameClockPrivate.
type FrameClockPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkFrameClockPrivate that represents the FrameClockPrivate.
func (recv *FrameClockPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FrameTimings is a representation of the C record GdkFrameTimings.
type FrameTimings struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkFrameTimings that represents the FrameTimings.
func (recv *FrameTimings) ToC() unsafe.Pointer {
	return recv.native
}

// Geometry is a representation of the C record GdkGeometry.
type Geometry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkGeometry that represents the Geometry.
func (recv *Geometry) ToC() unsafe.Pointer {
	return recv.native
}

// KeymapKey is a representation of the C record GdkKeymapKey.
type KeymapKey struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkKeymapKey that represents the KeymapKey.
func (recv *KeymapKey) ToC() unsafe.Pointer {
	return recv.native
}

// MonitorClass is a representation of the C record GdkMonitorClass.
type MonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkMonitorClass that represents the MonitorClass.
func (recv *MonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// Point is a representation of the C record GdkPoint.
type Point struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkPoint that represents the Point.
func (recv *Point) ToC() unsafe.Pointer {
	return recv.native
}

// RGBA is a representation of the C record GdkRGBA.
type RGBA struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkRGBA that represents the RGBA.
func (recv *RGBA) ToC() unsafe.Pointer {
	return recv.native
}

// Rectangle is a representation of the C record GdkRectangle.
type Rectangle struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkRectangle that represents the Rectangle.
func (recv *Rectangle) ToC() unsafe.Pointer {
	return recv.native
}

// TimeCoord is a representation of the C record GdkTimeCoord.
type TimeCoord struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkTimeCoord that represents the TimeCoord.
func (recv *TimeCoord) ToC() unsafe.Pointer {
	return recv.native
}

// WindowAttr is a representation of the C record GdkWindowAttr.
type WindowAttr struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkWindowAttr that represents the WindowAttr.
func (recv *WindowAttr) ToC() unsafe.Pointer {
	return recv.native
}

// WindowClass is a representation of the C record GdkWindowClass.
type WindowClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkWindowClass that represents the WindowClass.
func (recv *WindowClass) ToC() unsafe.Pointer {
	return recv.native
}

// WindowRedirect is a representation of the C record GdkWindowRedirect.
type WindowRedirect struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkWindowRedirect that represents the WindowRedirect.
func (recv *WindowRedirect) ToC() unsafe.Pointer {
	return recv.native
}

// AppLaunchContext is a representation of the C record GdkAppLaunchContext.
type AppLaunchContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkAppLaunchContext that represents the AppLaunchContext.
func (recv *AppLaunchContext) ToC() unsafe.Pointer {
	return recv.native
}

// Cursor is a representation of the C record GdkCursor.
type Cursor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkCursor that represents the Cursor.
func (recv *Cursor) ToC() unsafe.Pointer {
	return recv.native
}

// Device is a representation of the C record GdkDevice.
type Device struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkDevice that represents the Device.
func (recv *Device) ToC() unsafe.Pointer {
	return recv.native
}

// DeviceManager is a representation of the C record GdkDeviceManager.
type DeviceManager struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkDeviceManager that represents the DeviceManager.
func (recv *DeviceManager) ToC() unsafe.Pointer {
	return recv.native
}

// DeviceTool is a representation of the C record GdkDeviceTool.
type DeviceTool struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkDeviceTool that represents the DeviceTool.
func (recv *DeviceTool) ToC() unsafe.Pointer {
	return recv.native
}

// Display is a representation of the C record GdkDisplay.
type Display struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkDisplay that represents the Display.
func (recv *Display) ToC() unsafe.Pointer {
	return recv.native
}

// DisplayManager is a representation of the C record GdkDisplayManager.
type DisplayManager struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkDisplayManager that represents the DisplayManager.
func (recv *DisplayManager) ToC() unsafe.Pointer {
	return recv.native
}

// DragContext is a representation of the C record GdkDragContext.
type DragContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkDragContext that represents the DragContext.
func (recv *DragContext) ToC() unsafe.Pointer {
	return recv.native
}

// DrawingContext is a representation of the C record GdkDrawingContext.
type DrawingContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkDrawingContext that represents the DrawingContext.
func (recv *DrawingContext) ToC() unsafe.Pointer {
	return recv.native
}

// FrameClock is a representation of the C record GdkFrameClock.
type FrameClock struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkFrameClock that represents the FrameClock.
func (recv *FrameClock) ToC() unsafe.Pointer {
	return recv.native
}

// GLContext is a representation of the C record GdkGLContext.
type GLContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkGLContext that represents the GLContext.
func (recv *GLContext) ToC() unsafe.Pointer {
	return recv.native
}

// Keymap is a representation of the C record GdkKeymap.
type Keymap struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkKeymap that represents the Keymap.
func (recv *Keymap) ToC() unsafe.Pointer {
	return recv.native
}

// Monitor is a representation of the C record GdkMonitor.
type Monitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkMonitor that represents the Monitor.
func (recv *Monitor) ToC() unsafe.Pointer {
	return recv.native
}

// Screen is a representation of the C record GdkScreen.
type Screen struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkScreen that represents the Screen.
func (recv *Screen) ToC() unsafe.Pointer {
	return recv.native
}

// Seat is a representation of the C record GdkSeat.
type Seat struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkSeat that represents the Seat.
func (recv *Seat) ToC() unsafe.Pointer {
	return recv.native
}

// Visual is a representation of the C record GdkVisual.
type Visual struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkVisual that represents the Visual.
func (recv *Visual) ToC() unsafe.Pointer {
	return recv.native
}

// Window is a representation of the C record GdkWindow.
type Window struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkWindow that represents the Window.
func (recv *Window) ToC() unsafe.Pointer {
	return recv.native
}

// DevicePad is a representation of the C interface GdkDevicePad.
type DevicePad struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkDevicePad that represents the DevicePad.
func (recv *DevicePad) ToC() unsafe.Pointer {
	return recv.native
}

// Event is a representation of the C union GdkEvent.
type Event struct{}
