// This is a generated file - DO NOT EDIT
// +build gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "C"

type ModifierIntent C.GdkModifierIntent

const (
	GDK_MODIFIER_INTENT_PRIMARY_ACCELERATOR ModifierIntent = 0
	GDK_MODIFIER_INTENT_CONTEXT_MENU        ModifierIntent = 1
	GDK_MODIFIER_INTENT_EXTEND_SELECTION    ModifierIntent = 2
	GDK_MODIFIER_INTENT_MODIFY_SELECTION    ModifierIntent = 3
	GDK_MODIFIER_INTENT_NO_TEXT_INPUT       ModifierIntent = 4
	GDK_MODIFIER_INTENT_SHIFT_GROUP         ModifierIntent = 5
	GDK_MODIFIER_INTENT_DEFAULT_MOD_MASK    ModifierIntent = 6
)
