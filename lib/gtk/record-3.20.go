// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// FileChooserNativeClass is a wrapper around the C record GtkFileChooserNativeClass.
type FileChooserNativeClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// NativeDialogClass is a wrapper around the C record GtkNativeDialogClass.
type NativeDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for response
	// no type for show
	// no type for hide
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// PadActionEntry is a wrapper around the C record GtkPadActionEntry.
type PadActionEntry struct {
	native     unsafe.Pointer
	Type       PadActionType
	Index      int32
	Mode       int32
	Label      string
	ActionName string
}

// PadControllerClass is a wrapper around the C record GtkPadControllerClass.
type PadControllerClass struct {
	native unsafe.Pointer
}

// ShortcutLabelClass is a wrapper around the C record GtkShortcutLabelClass.
type ShortcutLabelClass struct {
	native unsafe.Pointer
}

// ShortcutsGroupClass is a wrapper around the C record GtkShortcutsGroupClass.
type ShortcutsGroupClass struct {
	native unsafe.Pointer
}

// ShortcutsSectionClass is a wrapper around the C record GtkShortcutsSectionClass.
type ShortcutsSectionClass struct {
	native unsafe.Pointer
}

// ShortcutsShortcutClass is a wrapper around the C record GtkShortcutsShortcutClass.
type ShortcutsShortcutClass struct {
	native unsafe.Pointer
}

// ShortcutsWindowClass is a wrapper around the C record GtkShortcutsWindowClass.
type ShortcutsWindowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for close
	// no type for search
}
