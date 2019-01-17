// This is a generated file - DO NOT EDIT
// +build gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	gboolean callback_flowboxfilterfuncHandler(GObject *, GtkFlowBoxChild*, gpointer, gpointer);

*/
/*

	void callback_flowboxforeachfuncHandler(GObject *, GtkFlowBox*, GtkFlowBoxChild*, gpointer, gpointer);

*/
/*

	gint callback_flowboxsortfuncHandler(GObject *, GtkFlowBoxChild*, GtkFlowBoxChild*, gpointer, gpointer);

*/
import "C"

var callbackFlowboxfilterfuncId int
var callbackFlowboxfilterfuncMap = make(map[int]FlowboxfilterfuncCallback)
var callbackFlowboxfilterfuncLock sync.RWMutex

// FlowboxfilterfuncCallback is a callback function for a 'FlowBoxFilterFunc' callback.
type FlowboxfilterfuncCallback func(child *FlowBoxChild) bool

var callbackFlowboxforeachfuncId int
var callbackFlowboxforeachfuncMap = make(map[int]FlowboxforeachfuncCallback)
var callbackFlowboxforeachfuncLock sync.RWMutex

// FlowboxforeachfuncCallback is a callback function for a 'FlowBoxForeachFunc' callback.
type FlowboxforeachfuncCallback func(box *FlowBox, child *FlowBoxChild)

var callbackFlowboxsortfuncId int
var callbackFlowboxsortfuncMap = make(map[int]FlowboxsortfuncCallback)
var callbackFlowboxsortfuncLock sync.RWMutex

// FlowboxsortfuncCallback is a callback function for a 'FlowBoxSortFunc' callback.
type FlowboxsortfuncCallback func(child1 *FlowBoxChild, child2 *FlowBoxChild) int32
