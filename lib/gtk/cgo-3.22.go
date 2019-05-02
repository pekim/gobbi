// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// FileFilterNewFromGvariant is a wrapper around the C function gtk_file_filter_new_from_gvariant.
func FileFilterNewFromGvariant(variant *glib.Variant) *FileFilter {
	c_variant := (*C.GVariant)(C.NULL)
	if variant != nil {
		c_variant = (*C.GVariant)(variant.ToC())
	}

	retC := C.gtk_file_filter_new_from_gvariant(c_variant)
	retGo := FileFilterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PadControllerNew is a wrapper around the C function gtk_pad_controller_new.
func PadControllerNew(window *Window, group *gio.ActionGroup, pad *gdk.Device) *PadController {
	c_window := (*C.GtkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GtkWindow)(window.ToC())
	}

	c_group := (*C.GActionGroup)(group.ToC())

	c_pad := (*C.GdkDevice)(C.NULL)
	if pad != nil {
		c_pad = (*C.GdkDevice)(pad.ToC())
	}

	retC := C.gtk_pad_controller_new(c_window, c_group, c_pad)
	retGo := PadControllerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// PageSetupNewFromGvariant is a wrapper around the C function gtk_page_setup_new_from_gvariant.
func PageSetupNewFromGvariant(variant *glib.Variant) *PageSetup {
	c_variant := (*C.GVariant)(C.NULL)
	if variant != nil {
		c_variant = (*C.GVariant)(variant.ToC())
	}

	retC := C.gtk_page_setup_new_from_gvariant(c_variant)
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// PrintSettingsNewFromGvariant is a wrapper around the C function gtk_print_settings_new_from_gvariant.
func PrintSettingsNewFromGvariant(variant *glib.Variant) *PrintSettings {
	c_variant := (*C.GVariant)(C.NULL)
	if variant != nil {
		c_variant = (*C.GVariant)(variant.ToC())
	}

	retC := C.gtk_print_settings_new_from_gvariant(c_variant)
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ShortcutLabelNew is a wrapper around the C function gtk_shortcut_label_new.
func ShortcutLabelNew(accelerator string) *ShortcutLabel {
	c_accelerator := C.CString(accelerator)
	defer C.free(unsafe.Pointer(c_accelerator))

	retC := C.gtk_shortcut_label_new(c_accelerator)
	retGo := ShortcutLabelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PaperSizeNewFromGvariant is a wrapper around the C function gtk_paper_size_new_from_gvariant.
func PaperSizeNewFromGvariant(variant *glib.Variant) *PaperSize {
	c_variant := (*C.GVariant)(C.NULL)
	if variant != nil {
		c_variant = (*C.GVariant)(variant.ToC())
	}

	retC := C.gtk_paper_size_new_from_gvariant(c_variant)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}
