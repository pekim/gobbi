// This is a generated file - DO NOT EDIT
// +build gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// DoSetRelatedAction is a wrapper around the C function gtk_activatable_do_set_related_action.
func (recv *Activatable) DoSetRelatedAction(action *Action) {
	c_action := (*C.GtkAction)(action.ToC())

	C.gtk_activatable_do_set_related_action((*C.GtkActivatable)(recv.native), c_action)

	return
}

// GetRelatedAction is a wrapper around the C function gtk_activatable_get_related_action.
func (recv *Activatable) GetRelatedAction() *Action {
	retC := C.gtk_activatable_get_related_action((*C.GtkActivatable)(recv.native))
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUseActionAppearance is a wrapper around the C function gtk_activatable_get_use_action_appearance.
func (recv *Activatable) GetUseActionAppearance() bool {
	retC := C.gtk_activatable_get_use_action_appearance((*C.GtkActivatable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetRelatedAction is a wrapper around the C function gtk_activatable_set_related_action.
func (recv *Activatable) SetRelatedAction(action *Action) {
	c_action := (*C.GtkAction)(action.ToC())

	C.gtk_activatable_set_related_action((*C.GtkActivatable)(recv.native), c_action)

	return
}

// SetUseActionAppearance is a wrapper around the C function gtk_activatable_set_use_action_appearance.
func (recv *Activatable) SetUseActionAppearance(useAppearance bool) {
	c_use_appearance :=
		boolToGboolean(useAppearance)

	C.gtk_activatable_set_use_action_appearance((*C.GtkActivatable)(recv.native), c_use_appearance)

	return
}

// SyncActionProperties is a wrapper around the C function gtk_activatable_sync_action_properties.
func (recv *Activatable) SyncActionProperties(action *Action) {
	c_action := (*C.GtkAction)(action.ToC())

	C.gtk_activatable_sync_action_properties((*C.GtkActivatable)(recv.native), c_action)

	return
}

// GetOrientation is a wrapper around the C function gtk_orientable_get_orientation.
func (recv *Orientable) GetOrientation() Orientation {
	retC := C.gtk_orientable_get_orientation((*C.GtkOrientable)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// SetOrientation is a wrapper around the C function gtk_orientable_set_orientation.
func (recv *Orientable) SetOrientation(orientation Orientation) {
	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_orientable_set_orientation((*C.GtkOrientable)(recv.native), c_orientation)

	return
}
