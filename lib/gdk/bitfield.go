// This is a generated file - DO NOT EDIT

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GdkDragAction

// Blacklisted : GdkEventMask

type ModifierType C.guint

const (
	GDK_SHIFT_MASK                ModifierType = 1
	GDK_LOCK_MASK                 ModifierType = 2
	GDK_CONTROL_MASK              ModifierType = 4
	GDK_MOD1_MASK                 ModifierType = 8
	GDK_MOD2_MASK                 ModifierType = 16
	GDK_MOD3_MASK                 ModifierType = 32
	GDK_MOD4_MASK                 ModifierType = 64
	GDK_MOD5_MASK                 ModifierType = 128
	GDK_BUTTON1_MASK              ModifierType = 256
	GDK_BUTTON2_MASK              ModifierType = 512
	GDK_BUTTON3_MASK              ModifierType = 1024
	GDK_BUTTON4_MASK              ModifierType = 2048
	GDK_BUTTON5_MASK              ModifierType = 4096
	GDK_MODIFIER_RESERVED_13_MASK ModifierType = 8192
	GDK_MODIFIER_RESERVED_14_MASK ModifierType = 16384
	GDK_MODIFIER_RESERVED_15_MASK ModifierType = 32768
	GDK_MODIFIER_RESERVED_16_MASK ModifierType = 65536
	GDK_MODIFIER_RESERVED_17_MASK ModifierType = 131072
	GDK_MODIFIER_RESERVED_18_MASK ModifierType = 262144
	GDK_MODIFIER_RESERVED_19_MASK ModifierType = 524288
	GDK_MODIFIER_RESERVED_20_MASK ModifierType = 1048576
	GDK_MODIFIER_RESERVED_21_MASK ModifierType = 2097152
	GDK_MODIFIER_RESERVED_22_MASK ModifierType = 4194304
	GDK_MODIFIER_RESERVED_23_MASK ModifierType = 8388608
	GDK_MODIFIER_RESERVED_24_MASK ModifierType = 16777216
	GDK_MODIFIER_RESERVED_25_MASK ModifierType = 33554432
	GDK_SUPER_MASK                ModifierType = 67108864
	GDK_HYPER_MASK                ModifierType = 134217728
	GDK_META_MASK                 ModifierType = 268435456
	GDK_MODIFIER_RESERVED_29_MASK ModifierType = 536870912
	GDK_RELEASE_MASK              ModifierType = 1073741824
	GDK_MODIFIER_MASK             ModifierType = 1543512063
)

// Blacklisted : GdkWMDecoration

// Blacklisted : GdkWMFunction

type WindowAttributesType C.GdkWindowAttributesType

const (
	GDK_WA_TITLE     WindowAttributesType = 2
	GDK_WA_X         WindowAttributesType = 4
	GDK_WA_Y         WindowAttributesType = 8
	GDK_WA_CURSOR    WindowAttributesType = 16
	GDK_WA_VISUAL    WindowAttributesType = 32
	GDK_WA_WMCLASS   WindowAttributesType = 64
	GDK_WA_NOREDIR   WindowAttributesType = 128
	GDK_WA_TYPE_HINT WindowAttributesType = 256
)

// Blacklisted : GdkWindowHints

// Blacklisted : GdkWindowState
