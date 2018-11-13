// This is a generated file - DO NOT EDIT
// +build gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// This enum is used with gdk_keymap_get_modifier_mask()
// in order to determine what modifiers the
// currently used windowing system backend uses for particular
// purposes. For example, on X11/Windows, the Control key is used for
// invoking menu shortcuts (accelerators), whereas on Apple computers
// it’s the Command key (which correspond to %GDK_CONTROL_MASK and
// %GDK_MOD2_MASK, respectively).
type ModifierIntent C.GdkModifierIntent

const (
	// the primary modifier used to invoke
	// menu accelerators.
	GDK_MODIFIER_INTENT_PRIMARY_ACCELERATOR ModifierIntent = 0
	// the modifier used to invoke context menus.
	// Note that mouse button 3 always triggers context menus. When this modifier
	// is not 0, it additionally triggers context menus when used with mouse button 1.
	GDK_MODIFIER_INTENT_CONTEXT_MENU ModifierIntent = 1
	// the modifier used to extend selections
	// using `modifier`-click or `modifier`-cursor-key
	GDK_MODIFIER_INTENT_EXTEND_SELECTION ModifierIntent = 2
	// the modifier used to modify selections,
	// which in most cases means toggling the clicked item into or out of the selection.
	GDK_MODIFIER_INTENT_MODIFY_SELECTION ModifierIntent = 3
	// when any of these modifiers is pressed, the
	// key event cannot produce a symbol directly. This is meant to be used for
	// input methods, and for use cases like typeahead search.
	GDK_MODIFIER_INTENT_NO_TEXT_INPUT ModifierIntent = 4
	// the modifier that switches between keyboard
	// groups (AltGr on X11/Windows and Option/Alt on OS X).
	GDK_MODIFIER_INTENT_SHIFT_GROUP ModifierIntent = 5
	// The set of modifier masks accepted
	// as modifiers in accelerators. Needed because Command is mapped to MOD2 on
	// OSX, which is widely used, but on X11 MOD2 is NumLock and using that for a
	// mod key is problematic at best.
	// Ref: https://bugzilla.gnome.org/show_bug.cgi?id=736125.
	GDK_MODIFIER_INTENT_DEFAULT_MOD_MASK ModifierIntent = 6
)
