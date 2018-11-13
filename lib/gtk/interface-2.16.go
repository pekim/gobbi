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

// This is a utility function for #GtkActivatable implementors.
//
// When implementing #GtkActivatable you must call this when
// handling changes of the #GtkActivatable:related-action, and
// you must also use this to break references in #GObject->dispose().
//
// This function adds a reference to the currently set related
// action for you, it also makes sure the #GtkActivatable->update()
// method is called when the related #GtkAction properties change
// and registers to the action’s proxy list.
//
// > Be careful to call this before setting the local
// > copy of the #GtkAction property, since this function uses
// > gtk_activatable_get_related_action() to retrieve the
// > previous action.
/*

C function : gtk_activatable_do_set_related_action
*/
func (recv *Activatable) DoSetRelatedAction(action *Action) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	C.gtk_activatable_do_set_related_action((*C.GtkActivatable)(recv.native), c_action)

	return
}

// Gets the related #GtkAction for @activatable.
/*

C function : gtk_activatable_get_related_action
*/
func (recv *Activatable) GetRelatedAction() *Action {
	retC := C.gtk_activatable_get_related_action((*C.GtkActivatable)(recv.native))
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets whether this activatable should reset its layout
// and appearance when setting the related action or when
// the action changes appearance.
/*

C function : gtk_activatable_get_use_action_appearance
*/
func (recv *Activatable) GetUseActionAppearance() bool {
	retC := C.gtk_activatable_get_use_action_appearance((*C.GtkActivatable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the related action on the @activatable object.
//
// > #GtkActivatable implementors need to handle the #GtkActivatable:related-action
// > property and call gtk_activatable_do_set_related_action() when it changes.
/*

C function : gtk_activatable_set_related_action
*/
func (recv *Activatable) SetRelatedAction(action *Action) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	C.gtk_activatable_set_related_action((*C.GtkActivatable)(recv.native), c_action)

	return
}

// Sets whether this activatable should reset its layout and appearance
// when setting the related action or when the action changes appearance
//
// > #GtkActivatable implementors need to handle the
// > #GtkActivatable:use-action-appearance property and call
// > gtk_activatable_sync_action_properties() to update @activatable
// > if needed.
/*

C function : gtk_activatable_set_use_action_appearance
*/
func (recv *Activatable) SetUseActionAppearance(useAppearance bool) {
	c_use_appearance :=
		boolToGboolean(useAppearance)

	C.gtk_activatable_set_use_action_appearance((*C.GtkActivatable)(recv.native), c_use_appearance)

	return
}

// This is called to update the activatable completely, this is called
// internally when the #GtkActivatable:related-action property is set
// or unset and by the implementing class when
// #GtkActivatable:use-action-appearance changes.
/*

C function : gtk_activatable_sync_action_properties
*/
func (recv *Activatable) SyncActionProperties(action *Action) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	C.gtk_activatable_sync_action_properties((*C.GtkActivatable)(recv.native), c_action)

	return
}

// Retrieves the orientation of the @orientable.
/*

C function : gtk_orientable_get_orientation
*/
func (recv *Orientable) GetOrientation() Orientation {
	retC := C.gtk_orientable_get_orientation((*C.GtkOrientable)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// Sets the orientation of the @orientable.
/*

C function : gtk_orientable_set_orientation
*/
func (recv *Orientable) SetOrientation(orientation Orientation) {
	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_orientable_set_orientation((*C.GtkOrientable)(recv.native), c_orientation)

	return
}
