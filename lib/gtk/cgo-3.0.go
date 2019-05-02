// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
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

// AppChooserButtonNew is a wrapper around the C function gtk_app_chooser_button_new.
func AppChooserButtonNew(contentType string) *AppChooserButton {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_button_new(c_content_type)
	retGo := AppChooserButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppChooserDialogNew is a wrapper around the C function gtk_app_chooser_dialog_new.
func AppChooserDialogNew(parent *Window, flags DialogFlags, file *gio.File) *AppChooserDialog {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	c_flags := (C.GtkDialogFlags)(flags)

	c_file := (*C.GFile)(file.ToC())

	retC := C.gtk_app_chooser_dialog_new(c_parent, c_flags, c_file)
	retGo := AppChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppChooserDialogNewForContentType is a wrapper around the C function gtk_app_chooser_dialog_new_for_content_type.
func AppChooserDialogNewForContentType(parent *Window, flags DialogFlags, contentType string) *AppChooserDialog {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	c_flags := (C.GtkDialogFlags)(flags)

	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_dialog_new_for_content_type(c_parent, c_flags, c_content_type)
	retGo := AppChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppChooserWidgetNew is a wrapper around the C function gtk_app_chooser_widget_new.
func AppChooserWidgetNew(contentType string) *AppChooserWidget {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_widget_new(c_content_type)
	retGo := AppChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ApplicationNew is a wrapper around the C function gtk_application_new.
func ApplicationNew(applicationId string, flags gio.ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.gtk_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BoxNew is a wrapper around the C function gtk_box_new.
func BoxNew(orientation Orientation, spacing int32) *Box {
	c_orientation := (C.GtkOrientation)(orientation)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_box_new(c_orientation, c_spacing)
	retGo := BoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ButtonBoxNew is a wrapper around the C function gtk_button_box_new.
func ButtonBoxNew(orientation Orientation) *ButtonBox {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_button_box_new(c_orientation)
	retGo := ButtonBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellAreaBoxNew is a wrapper around the C function gtk_cell_area_box_new.
func CellAreaBoxNew() *CellAreaBox {
	retC := C.gtk_cell_area_box_new()
	retGo := CellAreaBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ColorButtonNewWithRgba is a wrapper around the C function gtk_color_button_new_with_rgba.
func ColorButtonNewWithRgba(rgba *gdk.RGBA) *ColorButton {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	retC := C.gtk_color_button_new_with_rgba(c_rgba)
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EntryCompletionNewWithArea is a wrapper around the C function gtk_entry_completion_new_with_area.
func EntryCompletionNewWithArea(area *CellArea) *EntryCompletion {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_entry_completion_new_with_area(c_area)
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// IconViewNewWithArea is a wrapper around the C function gtk_icon_view_new_with_area.
func IconViewNewWithArea(area *CellArea) *IconView {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_icon_view_new_with_area(c_area)
	retGo := IconViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PanedNew is a wrapper around the C function gtk_paned_new.
func PanedNew(orientation Orientation) *Paned {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_paned_new(c_orientation)
	retGo := PanedNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ScaleNew is a wrapper around the C function gtk_scale_new.
func ScaleNew(orientation Orientation, adjustment *Adjustment) *Scale {
	c_orientation := (C.GtkOrientation)(orientation)

	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	retC := C.gtk_scale_new(c_orientation, c_adjustment)
	retGo := ScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ScaleNewWithRange is a wrapper around the C function gtk_scale_new_with_range.
func ScaleNewWithRange(orientation Orientation, min float64, max float64, step float64) *Scale {
	c_orientation := (C.GtkOrientation)(orientation)

	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_scale_new_with_range(c_orientation, c_min, c_max, c_step)
	retGo := ScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ScrollbarNew is a wrapper around the C function gtk_scrollbar_new.
func ScrollbarNew(orientation Orientation, adjustment *Adjustment) *Scrollbar {
	c_orientation := (C.GtkOrientation)(orientation)

	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	retC := C.gtk_scrollbar_new(c_orientation, c_adjustment)
	retGo := ScrollbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SeparatorNew is a wrapper around the C function gtk_separator_new.
func SeparatorNew(orientation Orientation) *Separator {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_separator_new(c_orientation)
	retGo := SeparatorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SwitchNew is a wrapper around the C function gtk_switch_new.
func SwitchNew() *Switch {
	retC := C.gtk_switch_new()
	retGo := SwitchNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TreeViewColumnNewWithArea is a wrapper around the C function gtk_tree_view_column_new_with_area.
func TreeViewColumnNewWithArea(area *CellArea) *TreeViewColumn {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_tree_view_column_new_with_area(c_area)
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GradientNewLinear is a wrapper around the C function gtk_gradient_new_linear.
func GradientNewLinear(x0 float64, y0 float64, x1 float64, y1 float64) *Gradient {
	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	retC := C.gtk_gradient_new_linear(c_x0, c_y0, c_x1, c_y1)
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GradientNewRadial is a wrapper around the C function gtk_gradient_new_radial.
func GradientNewRadial(x0 float64, y0 float64, radius0 float64, x1 float64, y1 float64, radius1 float64) *Gradient {
	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_radius0 := (C.gdouble)(radius0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	c_radius1 := (C.gdouble)(radius1)

	retC := C.gtk_gradient_new_radial(c_x0, c_y0, c_radius0, c_x1, c_y1, c_radius1)
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RequisitionNew is a wrapper around the C function gtk_requisition_new.
func RequisitionNew() *Requisition {
	retC := C.gtk_requisition_new()
	retGo := RequisitionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewAlpha is a wrapper around the C function gtk_symbolic_color_new_alpha.
func SymbolicColorNewAlpha(color *SymbolicColor, factor float64) *SymbolicColor {
	c_color := (*C.GtkSymbolicColor)(C.NULL)
	if color != nil {
		c_color = (*C.GtkSymbolicColor)(color.ToC())
	}

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_alpha(c_color, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewLiteral is a wrapper around the C function gtk_symbolic_color_new_literal.
func SymbolicColorNewLiteral(color *gdk.RGBA) *SymbolicColor {
	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	retC := C.gtk_symbolic_color_new_literal(c_color)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewMix is a wrapper around the C function gtk_symbolic_color_new_mix.
func SymbolicColorNewMix(color1 *SymbolicColor, color2 *SymbolicColor, factor float64) *SymbolicColor {
	c_color1 := (*C.GtkSymbolicColor)(C.NULL)
	if color1 != nil {
		c_color1 = (*C.GtkSymbolicColor)(color1.ToC())
	}

	c_color2 := (*C.GtkSymbolicColor)(C.NULL)
	if color2 != nil {
		c_color2 = (*C.GtkSymbolicColor)(color2.ToC())
	}

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_mix(c_color1, c_color2, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewName is a wrapper around the C function gtk_symbolic_color_new_name.
func SymbolicColorNewName(name string) *SymbolicColor {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_symbolic_color_new_name(c_name)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewShade is a wrapper around the C function gtk_symbolic_color_new_shade.
func SymbolicColorNewShade(color *SymbolicColor, factor float64) *SymbolicColor {
	c_color := (*C.GtkSymbolicColor)(C.NULL)
	if color != nil {
		c_color = (*C.GtkSymbolicColor)(color.ToC())
	}

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_shade(c_color, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// WidgetPathNew is a wrapper around the C function gtk_widget_path_new.
func WidgetPathNew() *WidgetPath {
	retC := C.gtk_widget_path_new()
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}
