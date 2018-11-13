// This is a generated file - DO NOT EDIT

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void celleditable_editingDoneHandler(GObject *, gpointer);

	static gulong CellEditable_signal_connect_editing_done(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "editing-done", G_CALLBACK(celleditable_editingDoneHandler), data);
	}

*/
/*

	void celleditable_removeWidgetHandler(GObject *, gpointer);

	static gulong CellEditable_signal_connect_remove_widget(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "remove-widget", G_CALLBACK(celleditable_removeWidgetHandler), data);
	}

*/
/*

	void editable_changedHandler(GObject *, gpointer);

	static gulong Editable_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(editable_changedHandler), data);
	}

*/
/*

	void filechooser_currentFolderChangedHandler(GObject *, gpointer);

	static gulong FileChooser_signal_connect_current_folder_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "current-folder-changed", G_CALLBACK(filechooser_currentFolderChangedHandler), data);
	}

*/
/*

	void filechooser_fileActivatedHandler(GObject *, gpointer);

	static gulong FileChooser_signal_connect_file_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "file-activated", G_CALLBACK(filechooser_fileActivatedHandler), data);
	}

*/
/*

	void filechooser_selectionChangedHandler(GObject *, gpointer);

	static gulong FileChooser_signal_connect_selection_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-changed", G_CALLBACK(filechooser_selectionChangedHandler), data);
	}

*/
/*

	void filechooser_updatePreviewHandler(GObject *, gpointer);

	static gulong FileChooser_signal_connect_update_preview(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "update-preview", G_CALLBACK(filechooser_updatePreviewHandler), data);
	}

*/
/*

	void printoperationpreview_gotPageSizeHandler(GObject *, GtkPrintContext *, GtkPageSetup *, gpointer);

	static gulong PrintOperationPreview_signal_connect_got_page_size(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "got-page-size", G_CALLBACK(printoperationpreview_gotPageSizeHandler), data);
	}

*/
/*

	void printoperationpreview_readyHandler(GObject *, GtkPrintContext *, gpointer);

	static gulong PrintOperationPreview_signal_connect_ready(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "ready", G_CALLBACK(printoperationpreview_readyHandler), data);
	}

*/
/*

	void treemodel_rowChangedHandler(GObject *, GtkTreePath *, GtkTreeIter *, gpointer);

	static gulong TreeModel_signal_connect_row_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-changed", G_CALLBACK(treemodel_rowChangedHandler), data);
	}

*/
/*

	void treemodel_rowDeletedHandler(GObject *, GtkTreePath *, gpointer);

	static gulong TreeModel_signal_connect_row_deleted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-deleted", G_CALLBACK(treemodel_rowDeletedHandler), data);
	}

*/
/*

	void treemodel_rowHasChildToggledHandler(GObject *, GtkTreePath *, GtkTreeIter *, gpointer);

	static gulong TreeModel_signal_connect_row_has_child_toggled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-has-child-toggled", G_CALLBACK(treemodel_rowHasChildToggledHandler), data);
	}

*/
/*

	void treemodel_rowInsertedHandler(GObject *, GtkTreePath *, GtkTreeIter *, gpointer);

	static gulong TreeModel_signal_connect_row_inserted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-inserted", G_CALLBACK(treemodel_rowInsertedHandler), data);
	}

*/
/*

	void treesortable_sortColumnChangedHandler(GObject *, gpointer);

	static gulong TreeSortable_signal_connect_sort_column_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "sort-column-changed", G_CALLBACK(treesortable_sortColumnChangedHandler), data);
	}

*/
import "C"

// This interface provides a convenient way of associating widgets with
// actions on a #GtkApplicationWindow or #GtkApplication.
//
// It primarily consists of two properties: #GtkActionable:action-name
// and #GtkActionable:action-target. There are also some convenience APIs
// for setting these properties.
//
// The action will be looked up in action groups that are found among
// the widgets ancestors. Most commonly, these will be the actions with
// the “win.” or “app.” prefix that are associated with the #GtkApplicationWindow
// or #GtkApplication, but other action groups that are added with
// gtk_widget_insert_action_group() will be consulted as well.
/*

C record/class : GtkActionable
*/
type Actionable struct {
	native *C.GtkActionable
}

func ActionableNewFromC(u unsafe.Pointer) *Actionable {
	c := (*C.GtkActionable)(u)
	if c == nil {
		return nil
	}

	g := &Actionable{native: c}

	return g
}

func (recv *Actionable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Activatable widgets can be connected to a #GtkAction and reflects
// the state of its action. A #GtkActivatable can also provide feedback
// through its action, as they are responsible for activating their
// related actions.
//
// # Implementing GtkActivatable
//
// When extending a class that is already #GtkActivatable; it is only
// necessary to implement the #GtkActivatable->sync_action_properties()
// and #GtkActivatable->update() methods and chain up to the parent
// implementation, however when introducing
// a new #GtkActivatable class; the #GtkActivatable:related-action and
// #GtkActivatable:use-action-appearance properties need to be handled by
// the implementor. Handling these properties is mostly a matter of installing
// the action pointer and boolean flag on your instance, and calling
// gtk_activatable_do_set_related_action() and
// gtk_activatable_sync_action_properties() at the appropriate times.
//
// ## A class fragment implementing #GtkActivatable
//
// |[<!-- language="C" -->
//
// enum {
// ...
//
// PROP_ACTIVATABLE_RELATED_ACTION,
// PROP_ACTIVATABLE_USE_ACTION_APPEARANCE
// }
//
// struct _FooBarPrivate
// {
//
// ...
//
// GtkAction      *action;
// gboolean        use_action_appearance;
// };
//
// ...
//
// static void foo_bar_activatable_interface_init         (GtkActivatableIface  *iface);
// static void foo_bar_activatable_update                 (GtkActivatable       *activatable,
// GtkAction            *action,
// const gchar          *property_name);
// static void foo_bar_activatable_sync_action_properties (GtkActivatable       *activatable,
// GtkAction            *action);
// ...
//
//
// static void
// foo_bar_class_init (FooBarClass *klass)
// {
//
// ...
//
// g_object_class_override_property (gobject_class, PROP_ACTIVATABLE_RELATED_ACTION, "related-action");
// g_object_class_override_property (gobject_class, PROP_ACTIVATABLE_USE_ACTION_APPEARANCE, "use-action-appearance");
//
// ...
// }
//
//
// static void
// foo_bar_activatable_interface_init (GtkActivatableIface  *iface)
// {
// iface->update = foo_bar_activatable_update;
// iface->sync_action_properties = foo_bar_activatable_sync_action_properties;
// }
//
// ... Break the reference using gtk_activatable_do_set_related_action()...
//
// static void
// foo_bar_dispose (GObject *object)
// {
// FooBar *bar = FOO_BAR (object);
// FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
// ...
//
// if (priv->action)
// {
// gtk_activatable_do_set_related_action (GTK_ACTIVATABLE (bar), NULL);
// priv->action = NULL;
// }
// G_OBJECT_CLASS (foo_bar_parent_class)->dispose (object);
// }
//
// ... Handle the “related-action” and “use-action-appearance” properties ...
//
// static void
// foo_bar_set_property (GObject         *object,
// guint            prop_id,
// const GValue    *value,
// GParamSpec      *pspec)
// {
// FooBar *bar = FOO_BAR (object);
// FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
// switch (prop_id)
// {
//
// ...
//
// case PROP_ACTIVATABLE_RELATED_ACTION:
// foo_bar_set_related_action (bar, g_value_get_object (value));
// break;
// case PROP_ACTIVATABLE_USE_ACTION_APPEARANCE:
// foo_bar_set_use_action_appearance (bar, g_value_get_boolean (value));
// break;
// default:
// G_OBJECT_WARN_INVALID_PROPERTY_ID (object, prop_id, pspec);
// break;
// }
// }
//
// static void
// foo_bar_get_property (GObject         *object,
// guint            prop_id,
// GValue          *value,
// GParamSpec      *pspec)
// {
// FooBar *bar = FOO_BAR (object);
// FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
// switch (prop_id)
// {
//
// ...
//
// case PROP_ACTIVATABLE_RELATED_ACTION:
// g_value_set_object (value, priv->action);
// break;
// case PROP_ACTIVATABLE_USE_ACTION_APPEARANCE:
// g_value_set_boolean (value, priv->use_action_appearance);
// break;
// default:
// G_OBJECT_WARN_INVALID_PROPERTY_ID (object, prop_id, pspec);
// break;
// }
// }
//
//
// static void
// foo_bar_set_use_action_appearance (FooBar   *bar,
// gboolean  use_appearance)
// {
// FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
// if (priv->use_action_appearance != use_appearance)
// {
// priv->use_action_appearance = use_appearance;
//
// gtk_activatable_sync_action_properties (GTK_ACTIVATABLE (bar), priv->action);
// }
// }
//
// ... call gtk_activatable_do_set_related_action() and then assign the action pointer,
// no need to reference the action here since gtk_activatable_do_set_related_action() already
// holds a reference here for you...
// static void
// foo_bar_set_related_action (FooBar    *bar,
// GtkAction *action)
// {
// FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
// if (priv->action == action)
// return;
//
// gtk_activatable_do_set_related_action (GTK_ACTIVATABLE (bar), action);
//
// priv->action = action;
// }
//
// ... Selectively reset and update activatable depending on the use-action-appearance property ...
// static void
// gtk_button_activatable_sync_action_properties (GtkActivatable       *activatable,
// GtkAction            *action)
// {
// GtkButtonPrivate *priv = GTK_BUTTON_GET_PRIVATE (activatable);
//
// if (!action)
// return;
//
// if (gtk_action_is_visible (action))
// gtk_widget_show (GTK_WIDGET (activatable));
// else
// gtk_widget_hide (GTK_WIDGET (activatable));
//
// gtk_widget_set_sensitive (GTK_WIDGET (activatable), gtk_action_is_sensitive (action));
//
// ...
//
// if (priv->use_action_appearance)
// {
// if (gtk_action_get_stock_id (action))
// foo_bar_set_stock (button, gtk_action_get_stock_id (action));
// else if (gtk_action_get_label (action))
// foo_bar_set_label (button, gtk_action_get_label (action));
//
// ...
//
// }
// }
//
// static void
// foo_bar_activatable_update (GtkActivatable       *activatable,
// GtkAction            *action,
// const gchar          *property_name)
// {
// FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (activatable);
//
// if (strcmp (property_name, "visible") == 0)
// {
// if (gtk_action_is_visible (action))
// gtk_widget_show (GTK_WIDGET (activatable));
// else
// gtk_widget_hide (GTK_WIDGET (activatable));
// }
// else if (strcmp (property_name, "sensitive") == 0)
// gtk_widget_set_sensitive (GTK_WIDGET (activatable), gtk_action_is_sensitive (action));
//
// ...
//
// if (!priv->use_action_appearance)
// return;
//
// if (strcmp (property_name, "stock-id") == 0)
// foo_bar_set_stock (button, gtk_action_get_stock_id (action));
// else if (strcmp (property_name, "label") == 0)
// foo_bar_set_label (button, gtk_action_get_label (action));
//
// ...
// }
// ]|
/*

C record/class : GtkActivatable
*/
type Activatable struct {
	native *C.GtkActivatable
}

func ActivatableNewFromC(u unsafe.Pointer) *Activatable {
	c := (*C.GtkActivatable)(u)
	if c == nil {
		return nil
	}

	g := &Activatable{native: c}

	return g
}

func (recv *Activatable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GtkAppChooser is an interface that can be implemented by widgets which
// allow the user to choose an application (typically for the purpose of
// opening a file). The main objects that implement this interface are
// #GtkAppChooserWidget, #GtkAppChooserDialog and #GtkAppChooserButton.
//
// Applications are represented by GIO #GAppInfo objects here.
// GIO has a concept of recommended and fallback applications for a
// given content type. Recommended applications are those that claim
// to handle the content type itself, while fallback also includes
// applications that handle a more generic content type. GIO also
// knows the default and last-used application for a given content
// type. The #GtkAppChooserWidget provides detailed control over
// whether the shown list of applications should include default,
// recommended or fallback applications.
//
// To obtain the application that has been selected in a #GtkAppChooser,
// use gtk_app_chooser_get_app_info().
/*

C record/class : GtkAppChooser
*/
type AppChooser struct {
	native *C.GtkAppChooser
}

func AppChooserNewFromC(u unsafe.Pointer) *AppChooser {
	c := (*C.GtkAppChooser)(u)
	if c == nil {
		return nil
	}

	g := &AppChooser{native: c}

	return g
}

func (recv *AppChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GtkBuildable allows objects to extend and customize their deserialization
// from [GtkBuilder UI descriptions][BUILDER-UI].
// The interface includes methods for setting names and properties of objects,
// parsing custom tags and constructing child objects.
//
// The GtkBuildable interface is implemented by all widgets and
// many of the non-widget objects that are provided by GTK+. The
// main user of this interface is #GtkBuilder. There should be
// very little need for applications to call any of these functions directly.
//
// An object only needs to implement this interface if it needs to extend the
// #GtkBuilder format or run any extra routines at deserialization time.
/*

C record/class : GtkBuildable
*/
type Buildable struct {
	native *C.GtkBuildable
}

func BuildableNewFromC(u unsafe.Pointer) *Buildable {
	c := (*C.GtkBuildable)(u)
	if c == nil {
		return nil
	}

	g := &Buildable{native: c}

	return g
}

func (recv *Buildable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C record/class : GtkCellAccessibleParent
*/
type CellAccessibleParent struct {
	native *C.GtkCellAccessibleParent
}

func CellAccessibleParentNewFromC(u unsafe.Pointer) *CellAccessibleParent {
	c := (*C.GtkCellAccessibleParent)(u)
	if c == nil {
		return nil
	}

	g := &CellAccessibleParent{native: c}

	return g
}

func (recv *CellAccessibleParent) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C function : gtk_cell_accessible_parent_activate
*/
func (recv *CellAccessibleParent) Activate(cell *CellAccessible) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	C.gtk_cell_accessible_parent_activate((*C.GtkCellAccessibleParent)(recv.native), c_cell)

	return
}

/*

C function : gtk_cell_accessible_parent_edit
*/
func (recv *CellAccessibleParent) Edit(cell *CellAccessible) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	C.gtk_cell_accessible_parent_edit((*C.GtkCellAccessibleParent)(recv.native), c_cell)

	return
}

/*

C function : gtk_cell_accessible_parent_expand_collapse
*/
func (recv *CellAccessibleParent) ExpandCollapse(cell *CellAccessible) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	C.gtk_cell_accessible_parent_expand_collapse((*C.GtkCellAccessibleParent)(recv.native), c_cell)

	return
}

// Unsupported : gtk_cell_accessible_parent_get_cell_area : unsupported parameter cell_rect : Blacklisted record : GdkRectangle

/*

C function : gtk_cell_accessible_parent_get_cell_extents
*/
func (recv *CellAccessibleParent) GetCellExtents(cell *CellAccessible, x int32, y int32, width int32, height int32, coordType atk.CoordType) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_coord_type := (C.AtkCoordType)(coordType)

	C.gtk_cell_accessible_parent_get_cell_extents((*C.GtkCellAccessibleParent)(recv.native), c_cell, &c_x, &c_y, &c_width, &c_height, c_coord_type)

	return
}

/*

C function : gtk_cell_accessible_parent_get_child_index
*/
func (recv *CellAccessibleParent) GetChildIndex(cell *CellAccessible) int32 {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	retC := C.gtk_cell_accessible_parent_get_child_index((*C.GtkCellAccessibleParent)(recv.native), c_cell)
	retGo := (int32)(retC)

	return retGo
}

/*

C function : gtk_cell_accessible_parent_get_renderer_state
*/
func (recv *CellAccessibleParent) GetRendererState(cell *CellAccessible) CellRendererState {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	retC := C.gtk_cell_accessible_parent_get_renderer_state((*C.GtkCellAccessibleParent)(recv.native), c_cell)
	retGo := (CellRendererState)(retC)

	return retGo
}

/*

C function : gtk_cell_accessible_parent_grab_focus
*/
func (recv *CellAccessibleParent) GrabFocus(cell *CellAccessible) bool {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	retC := C.gtk_cell_accessible_parent_grab_focus((*C.GtkCellAccessibleParent)(recv.native), c_cell)
	retGo := retC == C.TRUE

	return retGo
}

/*

C function : gtk_cell_accessible_parent_update_relationset
*/
func (recv *CellAccessibleParent) UpdateRelationset(cell *CellAccessible, relationset *atk.RelationSet) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	c_relationset := (*C.AtkRelationSet)(C.NULL)
	if relationset != nil {
		c_relationset = (*C.AtkRelationSet)(relationset.ToC())
	}

	C.gtk_cell_accessible_parent_update_relationset((*C.GtkCellAccessibleParent)(recv.native), c_cell, c_relationset)

	return
}

// The #GtkCellEditable interface must be implemented for widgets to be usable
// when editing the contents of a #GtkTreeView cell.
/*

C record/class : GtkCellEditable
*/
type CellEditable struct {
	native *C.GtkCellEditable
}

func CellEditableNewFromC(u unsafe.Pointer) *CellEditable {
	c := (*C.GtkCellEditable)(u)
	if c == nil {
		return nil
	}

	g := &CellEditable{native: c}

	return g
}

func (recv *CellEditable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalCellEditableEditingDoneDetail struct {
	callback  CellEditableSignalEditingDoneCallback
	handlerID C.gulong
}

var signalCellEditableEditingDoneId int
var signalCellEditableEditingDoneMap = make(map[int]signalCellEditableEditingDoneDetail)
var signalCellEditableEditingDoneLock sync.Mutex

// CellEditableSignalEditingDoneCallback is a callback function for a 'editing-done' signal emitted from a CellEditable.
type CellEditableSignalEditingDoneCallback func()

/*
ConnectEditingDone connects the callback to the 'editing-done' signal for the CellEditable.

The returned value represents the connection, and may be passed to DisconnectEditingDone to remove it.
*/
func (recv *CellEditable) ConnectEditingDone(callback CellEditableSignalEditingDoneCallback) int {
	signalCellEditableEditingDoneLock.Lock()
	defer signalCellEditableEditingDoneLock.Unlock()

	signalCellEditableEditingDoneId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellEditable_signal_connect_editing_done(instance, C.gpointer(uintptr(signalCellEditableEditingDoneId)))

	detail := signalCellEditableEditingDoneDetail{callback, handlerID}
	signalCellEditableEditingDoneMap[signalCellEditableEditingDoneId] = detail

	return signalCellEditableEditingDoneId
}

/*
DisconnectEditingDone disconnects a callback from the 'editing-done' signal for the CellEditable.

The connectionID should be a value returned from a call to ConnectEditingDone.
*/
func (recv *CellEditable) DisconnectEditingDone(connectionID int) {
	signalCellEditableEditingDoneLock.Lock()
	defer signalCellEditableEditingDoneLock.Unlock()

	detail, exists := signalCellEditableEditingDoneMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellEditableEditingDoneMap, connectionID)
}

//export celleditable_editingDoneHandler
func celleditable_editingDoneHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalCellEditableEditingDoneMap[index].callback
	callback()
}

type signalCellEditableRemoveWidgetDetail struct {
	callback  CellEditableSignalRemoveWidgetCallback
	handlerID C.gulong
}

var signalCellEditableRemoveWidgetId int
var signalCellEditableRemoveWidgetMap = make(map[int]signalCellEditableRemoveWidgetDetail)
var signalCellEditableRemoveWidgetLock sync.Mutex

// CellEditableSignalRemoveWidgetCallback is a callback function for a 'remove-widget' signal emitted from a CellEditable.
type CellEditableSignalRemoveWidgetCallback func()

/*
ConnectRemoveWidget connects the callback to the 'remove-widget' signal for the CellEditable.

The returned value represents the connection, and may be passed to DisconnectRemoveWidget to remove it.
*/
func (recv *CellEditable) ConnectRemoveWidget(callback CellEditableSignalRemoveWidgetCallback) int {
	signalCellEditableRemoveWidgetLock.Lock()
	defer signalCellEditableRemoveWidgetLock.Unlock()

	signalCellEditableRemoveWidgetId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellEditable_signal_connect_remove_widget(instance, C.gpointer(uintptr(signalCellEditableRemoveWidgetId)))

	detail := signalCellEditableRemoveWidgetDetail{callback, handlerID}
	signalCellEditableRemoveWidgetMap[signalCellEditableRemoveWidgetId] = detail

	return signalCellEditableRemoveWidgetId
}

/*
DisconnectRemoveWidget disconnects a callback from the 'remove-widget' signal for the CellEditable.

The connectionID should be a value returned from a call to ConnectRemoveWidget.
*/
func (recv *CellEditable) DisconnectRemoveWidget(connectionID int) {
	signalCellEditableRemoveWidgetLock.Lock()
	defer signalCellEditableRemoveWidgetLock.Unlock()

	detail, exists := signalCellEditableRemoveWidgetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellEditableRemoveWidgetMap, connectionID)
}

//export celleditable_removeWidgetHandler
func celleditable_removeWidgetHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalCellEditableRemoveWidgetMap[index].callback
	callback()
}

// Emits the #GtkCellEditable::editing-done signal.
/*

C function : gtk_cell_editable_editing_done
*/
func (recv *CellEditable) EditingDone() {
	C.gtk_cell_editable_editing_done((*C.GtkCellEditable)(recv.native))

	return
}

// Emits the #GtkCellEditable::remove-widget signal.
/*

C function : gtk_cell_editable_remove_widget
*/
func (recv *CellEditable) RemoveWidget() {
	C.gtk_cell_editable_remove_widget((*C.GtkCellEditable)(recv.native))

	return
}

// Unsupported : gtk_cell_editable_start_editing : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// #GtkCellLayout is an interface to be implemented by all objects which
// want to provide a #GtkTreeViewColumn like API for packing cells,
// setting attributes and data funcs.
//
// One of the notable features provided by implementations of
// GtkCellLayout are attributes. Attributes let you set the properties
// in flexible ways. They can just be set to constant values like regular
// properties. But they can also be mapped to a column of the underlying
// tree model with gtk_cell_layout_set_attributes(), which means that the value
// of the attribute can change from cell to cell as they are rendered by
// the cell renderer. Finally, it is possible to specify a function with
// gtk_cell_layout_set_cell_data_func() that is called to determine the
// value of the attribute for each cell that is rendered.
//
// # GtkCellLayouts as GtkBuildable
//
// Implementations of GtkCellLayout which also implement the GtkBuildable
// interface (#GtkCellView, #GtkIconView, #GtkComboBox,
// #GtkEntryCompletion, #GtkTreeViewColumn) accept GtkCellRenderer objects
// as <child> elements in UI definitions. They support a custom <attributes>
// element for their children, which can contain multiple <attribute>
// elements. Each <attribute> element has a name attribute which specifies
// a property of the cell renderer; the content of the element is the
// attribute value.
//
// This is an example of a UI definition fragment specifying attributes:
// |[
// <object class="GtkCellView">
// <child>
// <object class="GtkCellRendererText"/>
// <attributes>
// <attribute name="text">0</attribute>
// </attributes>
// </child>"
// </object>
// ]|
//
// Furthermore for implementations of GtkCellLayout that use a #GtkCellArea
// to lay out cells (all GtkCellLayouts in GTK+ use a GtkCellArea)
// [cell properties][cell-properties] can also be defined in the format by
// specifying the custom <cell-packing> attribute which can contain multiple
// <property> elements defined in the normal way.
//
// Here is a UI definition fragment specifying cell properties:
//
// |[
// <object class="GtkTreeViewColumn">
// <child>
// <object class="GtkCellRendererText"/>
// <cell-packing>
// <property name="align">True</property>
// <property name="expand">False</property>
// </cell-packing>
// </child>"
// </object>
// ]|
//
// # Subclassing GtkCellLayout implementations
//
// When subclassing a widget that implements #GtkCellLayout like
// #GtkIconView or #GtkComboBox, there are some considerations related
// to the fact that these widgets internally use a #GtkCellArea.
// The cell area is exposed as a construct-only property by these
// widgets. This means that it is possible to e.g. do
//
// |[<!-- language="C" -->
// combo = g_object_new (GTK_TYPE_COMBO_BOX, "cell-area", my_cell_area, NULL);
// ]|
//
// to use a custom cell area with a combo box. But construct properties
// are only initialized after instance init()
// functions have run, which means that using functions which rely on
// the existence of the cell area in your subclass’ init() function will
// cause the default cell area to be instantiated. In this case, a provided
// construct property value will be ignored (with a warning, to alert
// you to the problem).
//
// |[<!-- language="C" -->
// static void
// my_combo_box_init (MyComboBox *b)
// {
// GtkCellRenderer *cell;
//
// cell = gtk_cell_renderer_pixbuf_new ();
// The following call causes the default cell area for combo boxes,
// a GtkCellAreaBox, to be instantiated
// gtk_cell_layout_pack_start (GTK_CELL_LAYOUT (b), cell, FALSE);
// ...
// }
//
// GtkWidget *
// my_combo_box_new (GtkCellArea *area)
// {
// This call is going to cause a warning about area being ignored
// return g_object_new (MY_TYPE_COMBO_BOX, "cell-area", area, NULL);
// }
// ]|
//
// If supporting alternative cell areas with your derived widget is
// not important, then this does not have to concern you. If you want
// to support alternative cell areas, you can do so by moving the
// problematic calls out of init() and into a constructor()
// for your class.
/*

C record/class : GtkCellLayout
*/
type CellLayout struct {
	native *C.GtkCellLayout
}

func CellLayoutNewFromC(u unsafe.Pointer) *CellLayout {
	c := (*C.GtkCellLayout)(u)
	if c == nil {
		return nil
	}

	g := &CellLayout{native: c}

	return g
}

func (recv *CellLayout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GtkColorChooser is an interface that is implemented by widgets
// for choosing colors. Depending on the situation, colors may be
// allowed to have alpha (translucency).
//
// In GTK+, the main widgets that implement this interface are
// #GtkColorChooserWidget, #GtkColorChooserDialog and #GtkColorButton.
/*

C record/class : GtkColorChooser
*/
type ColorChooser struct {
	native *C.GtkColorChooser
}

func ColorChooserNewFromC(u unsafe.Pointer) *ColorChooser {
	c := (*C.GtkColorChooser)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooser{native: c}

	return g
}

func (recv *ColorChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GtkEditable interface is an interface which should be implemented by
// text editing widgets, such as #GtkEntry and #GtkSpinButton. It contains functions
// for generically manipulating an editable widget, a large number of action
// signals used for key bindings, and several signals that an application can
// connect to to modify the behavior of a widget.
//
// As an example of the latter usage, by connecting
// the following handler to #GtkEditable::insert-text, an application
// can convert all entry into a widget into uppercase.
//
// ## Forcing entry to uppercase.
//
// |[<!-- language="C" -->
// #include <ctype.h>;
//
// void
// insert_text_handler (GtkEditable *editable,
// const gchar *text,
// gint         length,
// gint        *position,
// gpointer     data)
// {
// gchar *result = g_utf8_strup (text, length);
//
// g_signal_handlers_block_by_func (editable,
// (gpointer) insert_text_handler, data);
// gtk_editable_insert_text (editable, result, length, position);
// g_signal_handlers_unblock_by_func (editable,
// (gpointer) insert_text_handler, data);
//
// g_signal_stop_emission_by_name (editable, "insert_text");
//
// g_free (result);
// }
// ]|
/*

C record/class : GtkEditable
*/
type Editable struct {
	native *C.GtkEditable
}

func EditableNewFromC(u unsafe.Pointer) *Editable {
	c := (*C.GtkEditable)(u)
	if c == nil {
		return nil
	}

	g := &Editable{native: c}

	return g
}

func (recv *Editable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalEditableChangedDetail struct {
	callback  EditableSignalChangedCallback
	handlerID C.gulong
}

var signalEditableChangedId int
var signalEditableChangedMap = make(map[int]signalEditableChangedDetail)
var signalEditableChangedLock sync.Mutex

// EditableSignalChangedCallback is a callback function for a 'changed' signal emitted from a Editable.
type EditableSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Editable.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Editable) ConnectChanged(callback EditableSignalChangedCallback) int {
	signalEditableChangedLock.Lock()
	defer signalEditableChangedLock.Unlock()

	signalEditableChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Editable_signal_connect_changed(instance, C.gpointer(uintptr(signalEditableChangedId)))

	detail := signalEditableChangedDetail{callback, handlerID}
	signalEditableChangedMap[signalEditableChangedId] = detail

	return signalEditableChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Editable.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Editable) DisconnectChanged(connectionID int) {
	signalEditableChangedLock.Lock()
	defer signalEditableChangedLock.Unlock()

	detail, exists := signalEditableChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEditableChangedMap, connectionID)
}

//export editable_changedHandler
func editable_changedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalEditableChangedMap[index].callback
	callback()
}

// Unsupported signal 'delete-text' for Editable : unsupported parameter start_pos : type gint :

// Unsupported signal 'insert-text' for Editable : unsupported parameter new_text : type utf8 :

// Copies the contents of the currently selected content in the editable and
// puts it on the clipboard.
/*

C function : gtk_editable_copy_clipboard
*/
func (recv *Editable) CopyClipboard() {
	C.gtk_editable_copy_clipboard((*C.GtkEditable)(recv.native))

	return
}

// Removes the contents of the currently selected content in the editable and
// puts it on the clipboard.
/*

C function : gtk_editable_cut_clipboard
*/
func (recv *Editable) CutClipboard() {
	C.gtk_editable_cut_clipboard((*C.GtkEditable)(recv.native))

	return
}

// Deletes the currently selected text of the editable.
// This call doesn’t do anything if there is no selected text.
/*

C function : gtk_editable_delete_selection
*/
func (recv *Editable) DeleteSelection() {
	C.gtk_editable_delete_selection((*C.GtkEditable)(recv.native))

	return
}

// Deletes a sequence of characters. The characters that are deleted are
// those characters at positions from @start_pos up to, but not including
// @end_pos. If @end_pos is negative, then the characters deleted
// are those from @start_pos to the end of the text.
//
// Note that the positions are specified in characters, not bytes.
/*

C function : gtk_editable_delete_text
*/
func (recv *Editable) DeleteText(startPos int32, endPos int32) {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	C.gtk_editable_delete_text((*C.GtkEditable)(recv.native), c_start_pos, c_end_pos)

	return
}

// Retrieves a sequence of characters. The characters that are retrieved
// are those characters at positions from @start_pos up to, but not
// including @end_pos. If @end_pos is negative, then the characters
// retrieved are those characters from @start_pos to the end of the text.
//
// Note that positions are specified in characters, not bytes.
/*

C function : gtk_editable_get_chars
*/
func (recv *Editable) GetChars(startPos int32, endPos int32) string {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	retC := C.gtk_editable_get_chars((*C.GtkEditable)(recv.native), c_start_pos, c_end_pos)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Retrieves whether @editable is editable. See
// gtk_editable_set_editable().
/*

C function : gtk_editable_get_editable
*/
func (recv *Editable) GetEditable() bool {
	retC := C.gtk_editable_get_editable((*C.GtkEditable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Retrieves the current position of the cursor relative to the start
// of the content of the editable.
//
// Note that this position is in characters, not in bytes.
/*

C function : gtk_editable_get_position
*/
func (recv *Editable) GetPosition() int32 {
	retC := C.gtk_editable_get_position((*C.GtkEditable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the selection bound of the editable. start_pos will be filled
// with the start of the selection and @end_pos with end. If no text was
// selected both will be identical and %FALSE will be returned.
//
// Note that positions are specified in characters, not bytes.
/*

C function : gtk_editable_get_selection_bounds
*/
func (recv *Editable) GetSelectionBounds() (bool, int32, int32) {
	var c_start_pos C.gint

	var c_end_pos C.gint

	retC := C.gtk_editable_get_selection_bounds((*C.GtkEditable)(recv.native), &c_start_pos, &c_end_pos)
	retGo := retC == C.TRUE

	startPos := (int32)(c_start_pos)

	endPos := (int32)(c_end_pos)

	return retGo, startPos, endPos
}

// Inserts @new_text_length bytes of @new_text into the contents of the
// widget, at position @position.
//
// Note that the position is in characters, not in bytes.
// The function updates @position to point after the newly inserted text.
/*

C function : gtk_editable_insert_text
*/
func (recv *Editable) InsertText(newText string, newTextLength int32, position int32) {
	c_new_text := C.CString(newText)
	defer C.free(unsafe.Pointer(c_new_text))

	c_new_text_length := (C.gint)(newTextLength)

	c_position := (C.gint)(position)

	C.gtk_editable_insert_text((*C.GtkEditable)(recv.native), c_new_text, c_new_text_length, &c_position)

	return
}

// Pastes the content of the clipboard to the current position of the
// cursor in the editable.
/*

C function : gtk_editable_paste_clipboard
*/
func (recv *Editable) PasteClipboard() {
	C.gtk_editable_paste_clipboard((*C.GtkEditable)(recv.native))

	return
}

// Selects a region of text. The characters that are selected are
// those characters at positions from @start_pos up to, but not
// including @end_pos. If @end_pos is negative, then the
// characters selected are those characters from @start_pos to
// the end of the text.
//
// Note that positions are specified in characters, not bytes.
/*

C function : gtk_editable_select_region
*/
func (recv *Editable) SelectRegion(startPos int32, endPos int32) {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	C.gtk_editable_select_region((*C.GtkEditable)(recv.native), c_start_pos, c_end_pos)

	return
}

// Determines if the user can edit the text in the editable
// widget or not.
/*

C function : gtk_editable_set_editable
*/
func (recv *Editable) SetEditable(isEditable bool) {
	c_is_editable :=
		boolToGboolean(isEditable)

	C.gtk_editable_set_editable((*C.GtkEditable)(recv.native), c_is_editable)

	return
}

// Sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0)
// index in the contents of the editable. The value must be less than or
// equal to the number of characters in the editable. A value of -1
// indicates that the position should be set after the last character
// of the editable. Note that @position is in characters, not in bytes.
/*

C function : gtk_editable_set_position
*/
func (recv *Editable) SetPosition(position int32) {
	c_position := (C.gint)(position)

	C.gtk_editable_set_position((*C.GtkEditable)(recv.native), c_position)

	return
}

// #GtkFileChooser is an interface that can be implemented by file
// selection widgets.  In GTK+, the main objects that implement this
// interface are #GtkFileChooserWidget, #GtkFileChooserDialog, and
// #GtkFileChooserButton.  You do not need to write an object that
// implements the #GtkFileChooser interface unless you are trying to
// adapt an existing file selector to expose a standard programming
// interface.
//
// #GtkFileChooser allows for shortcuts to various places in the filesystem.
// In the default implementation these are displayed in the left pane. It
// may be a bit confusing at first that these shortcuts come from various
// sources and in various flavours, so lets explain the terminology here:
//
// - Bookmarks: are created by the user, by dragging folders from the
// right pane to the left pane, or by using the “Add”. Bookmarks
// can be renamed and deleted by the user.
//
// - Shortcuts: can be provided by the application. For example, a Paint
// program may want to add a shortcut for a Clipart folder. Shortcuts
// cannot be modified by the user.
//
// - Volumes: are provided by the underlying filesystem abstraction. They are
// the “roots” of the filesystem.
//
// # File Names and Encodings
//
// When the user is finished selecting files in a
// #GtkFileChooser, your program can get the selected names
// either as filenames or as URIs.  For URIs, the normal escaping
// rules are applied if the URI contains non-ASCII characters.
// However, filenames are always returned in
// the character set specified by the
// `G_FILENAME_ENCODING` environment variable.
// Please see the GLib documentation for more details about this
// variable.
//
// This means that while you can pass the result of
// gtk_file_chooser_get_filename() to open() or fopen(),
// you may not be able to directly set it as the text of a
// #GtkLabel widget unless you convert it first to UTF-8,
// which all GTK+ widgets expect. You should use g_filename_to_utf8()
// to convert filenames into strings that can be passed to GTK+
// widgets.
//
// # Adding a Preview Widget
//
// You can add a custom preview widget to a file chooser and then
// get notification about when the preview needs to be updated.
// To install a preview widget, use
// gtk_file_chooser_set_preview_widget().  Then, connect to the
// #GtkFileChooser::update-preview signal to get notified when
// you need to update the contents of the preview.
//
// Your callback should use
// gtk_file_chooser_get_preview_filename() to see what needs
// previewing.  Once you have generated the preview for the
// corresponding file, you must call
// gtk_file_chooser_set_preview_widget_active() with a boolean
// flag that indicates whether your callback could successfully
// generate a preview.
//
// ## Example: Using a Preview Widget ## {#gtkfilechooser-preview}
// |[<!-- language="C" -->
// {
// GtkImage *preview;
//
// ...
//
// preview = gtk_image_new ();
//
// gtk_file_chooser_set_preview_widget (my_file_chooser, preview);
// g_signal_connect (my_file_chooser, "update-preview",
// G_CALLBACK (update_preview_cb), preview);
// }
//
// static void
// update_preview_cb (GtkFileChooser *file_chooser, gpointer data)
// {
// GtkWidget *preview;
// char *filename;
// GdkPixbuf *pixbuf;
// gboolean have_preview;
//
// preview = GTK_WIDGET (data);
// filename = gtk_file_chooser_get_preview_filename (file_chooser);
//
// pixbuf = gdk_pixbuf_new_from_file_at_size (filename, 128, 128, NULL);
// have_preview = (pixbuf != NULL);
// g_free (filename);
//
// gtk_image_set_from_pixbuf (GTK_IMAGE (preview), pixbuf);
// if (pixbuf)
// g_object_unref (pixbuf);
//
// gtk_file_chooser_set_preview_widget_active (file_chooser, have_preview);
// }
// ]|
//
// # Adding Extra Widgets
//
// You can add extra widgets to a file chooser to provide options
// that are not present in the default design.  For example, you
// can add a toggle button to give the user the option to open a
// file in read-only mode.  You can use
// gtk_file_chooser_set_extra_widget() to insert additional
// widgets in a file chooser.
//
// An example for adding extra widgets:
// |[<!-- language="C" -->
//
// GtkWidget *toggle;
//
// ...
//
// toggle = gtk_check_button_new_with_label ("Open file read-only");
// gtk_widget_show (toggle);
// gtk_file_chooser_set_extra_widget (my_file_chooser, toggle);
// }
// ]|
//
// If you want to set more than one extra widget in the file
// chooser, you can a container such as a #GtkBox or a #GtkGrid
// and include your widgets in it.  Then, set the container as
// the whole extra widget.
/*

C record/class : GtkFileChooser
*/
type FileChooser struct {
	native *C.GtkFileChooser
}

func FileChooserNewFromC(u unsafe.Pointer) *FileChooser {
	c := (*C.GtkFileChooser)(u)
	if c == nil {
		return nil
	}

	g := &FileChooser{native: c}

	return g
}

func (recv *FileChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalFileChooserCurrentFolderChangedDetail struct {
	callback  FileChooserSignalCurrentFolderChangedCallback
	handlerID C.gulong
}

var signalFileChooserCurrentFolderChangedId int
var signalFileChooserCurrentFolderChangedMap = make(map[int]signalFileChooserCurrentFolderChangedDetail)
var signalFileChooserCurrentFolderChangedLock sync.Mutex

// FileChooserSignalCurrentFolderChangedCallback is a callback function for a 'current-folder-changed' signal emitted from a FileChooser.
type FileChooserSignalCurrentFolderChangedCallback func()

/*
ConnectCurrentFolderChanged connects the callback to the 'current-folder-changed' signal for the FileChooser.

The returned value represents the connection, and may be passed to DisconnectCurrentFolderChanged to remove it.
*/
func (recv *FileChooser) ConnectCurrentFolderChanged(callback FileChooserSignalCurrentFolderChangedCallback) int {
	signalFileChooserCurrentFolderChangedLock.Lock()
	defer signalFileChooserCurrentFolderChangedLock.Unlock()

	signalFileChooserCurrentFolderChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooser_signal_connect_current_folder_changed(instance, C.gpointer(uintptr(signalFileChooserCurrentFolderChangedId)))

	detail := signalFileChooserCurrentFolderChangedDetail{callback, handlerID}
	signalFileChooserCurrentFolderChangedMap[signalFileChooserCurrentFolderChangedId] = detail

	return signalFileChooserCurrentFolderChangedId
}

/*
DisconnectCurrentFolderChanged disconnects a callback from the 'current-folder-changed' signal for the FileChooser.

The connectionID should be a value returned from a call to ConnectCurrentFolderChanged.
*/
func (recv *FileChooser) DisconnectCurrentFolderChanged(connectionID int) {
	signalFileChooserCurrentFolderChangedLock.Lock()
	defer signalFileChooserCurrentFolderChangedLock.Unlock()

	detail, exists := signalFileChooserCurrentFolderChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserCurrentFolderChangedMap, connectionID)
}

//export filechooser_currentFolderChangedHandler
func filechooser_currentFolderChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalFileChooserCurrentFolderChangedMap[index].callback
	callback()
}

type signalFileChooserFileActivatedDetail struct {
	callback  FileChooserSignalFileActivatedCallback
	handlerID C.gulong
}

var signalFileChooserFileActivatedId int
var signalFileChooserFileActivatedMap = make(map[int]signalFileChooserFileActivatedDetail)
var signalFileChooserFileActivatedLock sync.Mutex

// FileChooserSignalFileActivatedCallback is a callback function for a 'file-activated' signal emitted from a FileChooser.
type FileChooserSignalFileActivatedCallback func()

/*
ConnectFileActivated connects the callback to the 'file-activated' signal for the FileChooser.

The returned value represents the connection, and may be passed to DisconnectFileActivated to remove it.
*/
func (recv *FileChooser) ConnectFileActivated(callback FileChooserSignalFileActivatedCallback) int {
	signalFileChooserFileActivatedLock.Lock()
	defer signalFileChooserFileActivatedLock.Unlock()

	signalFileChooserFileActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooser_signal_connect_file_activated(instance, C.gpointer(uintptr(signalFileChooserFileActivatedId)))

	detail := signalFileChooserFileActivatedDetail{callback, handlerID}
	signalFileChooserFileActivatedMap[signalFileChooserFileActivatedId] = detail

	return signalFileChooserFileActivatedId
}

/*
DisconnectFileActivated disconnects a callback from the 'file-activated' signal for the FileChooser.

The connectionID should be a value returned from a call to ConnectFileActivated.
*/
func (recv *FileChooser) DisconnectFileActivated(connectionID int) {
	signalFileChooserFileActivatedLock.Lock()
	defer signalFileChooserFileActivatedLock.Unlock()

	detail, exists := signalFileChooserFileActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserFileActivatedMap, connectionID)
}

//export filechooser_fileActivatedHandler
func filechooser_fileActivatedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalFileChooserFileActivatedMap[index].callback
	callback()
}

type signalFileChooserSelectionChangedDetail struct {
	callback  FileChooserSignalSelectionChangedCallback
	handlerID C.gulong
}

var signalFileChooserSelectionChangedId int
var signalFileChooserSelectionChangedMap = make(map[int]signalFileChooserSelectionChangedDetail)
var signalFileChooserSelectionChangedLock sync.Mutex

// FileChooserSignalSelectionChangedCallback is a callback function for a 'selection-changed' signal emitted from a FileChooser.
type FileChooserSignalSelectionChangedCallback func()

/*
ConnectSelectionChanged connects the callback to the 'selection-changed' signal for the FileChooser.

The returned value represents the connection, and may be passed to DisconnectSelectionChanged to remove it.
*/
func (recv *FileChooser) ConnectSelectionChanged(callback FileChooserSignalSelectionChangedCallback) int {
	signalFileChooserSelectionChangedLock.Lock()
	defer signalFileChooserSelectionChangedLock.Unlock()

	signalFileChooserSelectionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooser_signal_connect_selection_changed(instance, C.gpointer(uintptr(signalFileChooserSelectionChangedId)))

	detail := signalFileChooserSelectionChangedDetail{callback, handlerID}
	signalFileChooserSelectionChangedMap[signalFileChooserSelectionChangedId] = detail

	return signalFileChooserSelectionChangedId
}

/*
DisconnectSelectionChanged disconnects a callback from the 'selection-changed' signal for the FileChooser.

The connectionID should be a value returned from a call to ConnectSelectionChanged.
*/
func (recv *FileChooser) DisconnectSelectionChanged(connectionID int) {
	signalFileChooserSelectionChangedLock.Lock()
	defer signalFileChooserSelectionChangedLock.Unlock()

	detail, exists := signalFileChooserSelectionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserSelectionChangedMap, connectionID)
}

//export filechooser_selectionChangedHandler
func filechooser_selectionChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalFileChooserSelectionChangedMap[index].callback
	callback()
}

type signalFileChooserUpdatePreviewDetail struct {
	callback  FileChooserSignalUpdatePreviewCallback
	handlerID C.gulong
}

var signalFileChooserUpdatePreviewId int
var signalFileChooserUpdatePreviewMap = make(map[int]signalFileChooserUpdatePreviewDetail)
var signalFileChooserUpdatePreviewLock sync.Mutex

// FileChooserSignalUpdatePreviewCallback is a callback function for a 'update-preview' signal emitted from a FileChooser.
type FileChooserSignalUpdatePreviewCallback func()

/*
ConnectUpdatePreview connects the callback to the 'update-preview' signal for the FileChooser.

The returned value represents the connection, and may be passed to DisconnectUpdatePreview to remove it.
*/
func (recv *FileChooser) ConnectUpdatePreview(callback FileChooserSignalUpdatePreviewCallback) int {
	signalFileChooserUpdatePreviewLock.Lock()
	defer signalFileChooserUpdatePreviewLock.Unlock()

	signalFileChooserUpdatePreviewId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooser_signal_connect_update_preview(instance, C.gpointer(uintptr(signalFileChooserUpdatePreviewId)))

	detail := signalFileChooserUpdatePreviewDetail{callback, handlerID}
	signalFileChooserUpdatePreviewMap[signalFileChooserUpdatePreviewId] = detail

	return signalFileChooserUpdatePreviewId
}

/*
DisconnectUpdatePreview disconnects a callback from the 'update-preview' signal for the FileChooser.

The connectionID should be a value returned from a call to ConnectUpdatePreview.
*/
func (recv *FileChooser) DisconnectUpdatePreview(connectionID int) {
	signalFileChooserUpdatePreviewLock.Lock()
	defer signalFileChooserUpdatePreviewLock.Unlock()

	detail, exists := signalFileChooserUpdatePreviewMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserUpdatePreviewMap, connectionID)
}

//export filechooser_updatePreviewHandler
func filechooser_updatePreviewHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalFileChooserUpdatePreviewMap[index].callback
	callback()
}

// Gets whether a stock label should be drawn with the name of the previewed
// file.  See gtk_file_chooser_set_use_preview_label().
/*

C function : gtk_file_chooser_get_use_preview_label
*/
func (recv *FileChooser) GetUsePreviewLabel() bool {
	retC := C.gtk_file_chooser_get_use_preview_label((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// #GtkFontChooser is an interface that can be implemented by widgets
// displaying the list of fonts. In GTK+, the main objects
// that implement this interface are #GtkFontChooserWidget,
// #GtkFontChooserDialog and #GtkFontButton. The GtkFontChooser interface
// has been introducted in GTK+ 3.2.
/*

C record/class : GtkFontChooser
*/
type FontChooser struct {
	native *C.GtkFontChooser
}

func FontChooserNewFromC(u unsafe.Pointer) *FontChooser {
	c := (*C.GtkFontChooser)(u)
	if c == nil {
		return nil
	}

	g := &FontChooser{native: c}

	return g
}

func (recv *FontChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported signal 'font-activated' for FontChooser : unsupported parameter fontname : type utf8 :

// The #GtkOrientable interface is implemented by all widgets that can be
// oriented horizontally or vertically. Historically, such widgets have been
// realized as subclasses of a common base class (e.g #GtkBox/#GtkHBox/#GtkVBox
// or #GtkScale/#GtkHScale/#GtkVScale). #GtkOrientable is more flexible in that
// it allows the orientation to be changed at runtime, allowing the widgets
// to “flip”.
//
// #GtkOrientable was introduced in GTK+ 2.16.
/*

C record/class : GtkOrientable
*/
type Orientable struct {
	native *C.GtkOrientable
}

func OrientableNewFromC(u unsafe.Pointer) *Orientable {
	c := (*C.GtkOrientable)(u)
	if c == nil {
		return nil
	}

	g := &Orientable{native: c}

	return g
}

func (recv *Orientable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C record/class : GtkPrintOperationPreview
*/
type PrintOperationPreview struct {
	native *C.GtkPrintOperationPreview
}

func PrintOperationPreviewNewFromC(u unsafe.Pointer) *PrintOperationPreview {
	c := (*C.GtkPrintOperationPreview)(u)
	if c == nil {
		return nil
	}

	g := &PrintOperationPreview{native: c}

	return g
}

func (recv *PrintOperationPreview) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalPrintOperationPreviewGotPageSizeDetail struct {
	callback  PrintOperationPreviewSignalGotPageSizeCallback
	handlerID C.gulong
}

var signalPrintOperationPreviewGotPageSizeId int
var signalPrintOperationPreviewGotPageSizeMap = make(map[int]signalPrintOperationPreviewGotPageSizeDetail)
var signalPrintOperationPreviewGotPageSizeLock sync.Mutex

// PrintOperationPreviewSignalGotPageSizeCallback is a callback function for a 'got-page-size' signal emitted from a PrintOperationPreview.
type PrintOperationPreviewSignalGotPageSizeCallback func(context *PrintContext, pageSetup *PageSetup)

/*
ConnectGotPageSize connects the callback to the 'got-page-size' signal for the PrintOperationPreview.

The returned value represents the connection, and may be passed to DisconnectGotPageSize to remove it.
*/
func (recv *PrintOperationPreview) ConnectGotPageSize(callback PrintOperationPreviewSignalGotPageSizeCallback) int {
	signalPrintOperationPreviewGotPageSizeLock.Lock()
	defer signalPrintOperationPreviewGotPageSizeLock.Unlock()

	signalPrintOperationPreviewGotPageSizeId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperationPreview_signal_connect_got_page_size(instance, C.gpointer(uintptr(signalPrintOperationPreviewGotPageSizeId)))

	detail := signalPrintOperationPreviewGotPageSizeDetail{callback, handlerID}
	signalPrintOperationPreviewGotPageSizeMap[signalPrintOperationPreviewGotPageSizeId] = detail

	return signalPrintOperationPreviewGotPageSizeId
}

/*
DisconnectGotPageSize disconnects a callback from the 'got-page-size' signal for the PrintOperationPreview.

The connectionID should be a value returned from a call to ConnectGotPageSize.
*/
func (recv *PrintOperationPreview) DisconnectGotPageSize(connectionID int) {
	signalPrintOperationPreviewGotPageSizeLock.Lock()
	defer signalPrintOperationPreviewGotPageSizeLock.Unlock()

	detail, exists := signalPrintOperationPreviewGotPageSizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationPreviewGotPageSizeMap, connectionID)
}

//export printoperationpreview_gotPageSizeHandler
func printoperationpreview_gotPageSizeHandler(_ *C.GObject, c_context *C.GtkPrintContext, c_page_setup *C.GtkPageSetup, data C.gpointer) {
	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	pageSetup := PageSetupNewFromC(unsafe.Pointer(c_page_setup))

	index := int(uintptr(data))
	callback := signalPrintOperationPreviewGotPageSizeMap[index].callback
	callback(context, pageSetup)
}

type signalPrintOperationPreviewReadyDetail struct {
	callback  PrintOperationPreviewSignalReadyCallback
	handlerID C.gulong
}

var signalPrintOperationPreviewReadyId int
var signalPrintOperationPreviewReadyMap = make(map[int]signalPrintOperationPreviewReadyDetail)
var signalPrintOperationPreviewReadyLock sync.Mutex

// PrintOperationPreviewSignalReadyCallback is a callback function for a 'ready' signal emitted from a PrintOperationPreview.
type PrintOperationPreviewSignalReadyCallback func(context *PrintContext)

/*
ConnectReady connects the callback to the 'ready' signal for the PrintOperationPreview.

The returned value represents the connection, and may be passed to DisconnectReady to remove it.
*/
func (recv *PrintOperationPreview) ConnectReady(callback PrintOperationPreviewSignalReadyCallback) int {
	signalPrintOperationPreviewReadyLock.Lock()
	defer signalPrintOperationPreviewReadyLock.Unlock()

	signalPrintOperationPreviewReadyId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperationPreview_signal_connect_ready(instance, C.gpointer(uintptr(signalPrintOperationPreviewReadyId)))

	detail := signalPrintOperationPreviewReadyDetail{callback, handlerID}
	signalPrintOperationPreviewReadyMap[signalPrintOperationPreviewReadyId] = detail

	return signalPrintOperationPreviewReadyId
}

/*
DisconnectReady disconnects a callback from the 'ready' signal for the PrintOperationPreview.

The connectionID should be a value returned from a call to ConnectReady.
*/
func (recv *PrintOperationPreview) DisconnectReady(connectionID int) {
	signalPrintOperationPreviewReadyLock.Lock()
	defer signalPrintOperationPreviewReadyLock.Unlock()

	detail, exists := signalPrintOperationPreviewReadyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationPreviewReadyMap, connectionID)
}

//export printoperationpreview_readyHandler
func printoperationpreview_readyHandler(_ *C.GObject, c_context *C.GtkPrintContext, data C.gpointer) {
	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	index := int(uintptr(data))
	callback := signalPrintOperationPreviewReadyMap[index].callback
	callback(context)
}

// #GtkRecentChooser is an interface that can be implemented by widgets
// displaying the list of recently used files.  In GTK+, the main objects
// that implement this interface are #GtkRecentChooserWidget,
// #GtkRecentChooserDialog and #GtkRecentChooserMenu.
//
// Recently used files are supported since GTK+ 2.10.
/*

C record/class : GtkRecentChooser
*/
type RecentChooser struct {
	native *C.GtkRecentChooser
}

func RecentChooserNewFromC(u unsafe.Pointer) *RecentChooser {
	c := (*C.GtkRecentChooser)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooser{native: c}

	return g
}

func (recv *RecentChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GtkScrollable is an interface that is implemented by widgets with native
// scrolling ability.
//
// To implement this interface you should override the
// #GtkScrollable:hadjustment and #GtkScrollable:vadjustment properties.
//
// ## Creating a scrollable widget
//
// All scrollable widgets should do the following.
//
// - When a parent widget sets the scrollable child widget’s adjustments,
// the widget should populate the adjustments’
// #GtkAdjustment:lower, #GtkAdjustment:upper,
// #GtkAdjustment:step-increment, #GtkAdjustment:page-increment and
// #GtkAdjustment:page-size properties and connect to the
// #GtkAdjustment::value-changed signal.
//
// - Because its preferred size is the size for a fully expanded widget,
// the scrollable widget must be able to cope with underallocations.
// This means that it must accept any value passed to its
// #GtkWidgetClass.size_allocate() function.
//
// - When the parent allocates space to the scrollable child widget,
// the widget should update the adjustments’ properties with new values.
//
// - When any of the adjustments emits the #GtkAdjustment::value-changed signal,
// the scrollable widget should scroll its contents.
/*

C record/class : GtkScrollable
*/
type Scrollable struct {
	native *C.GtkScrollable
}

func ScrollableNewFromC(u unsafe.Pointer) *Scrollable {
	c := (*C.GtkScrollable)(u)
	if c == nil {
		return nil
	}

	g := &Scrollable{native: c}

	return g
}

func (recv *Scrollable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GtkStyleProvider is an interface used to provide style information to a #GtkStyleContext.
// See gtk_style_context_add_provider() and gtk_style_context_add_provider_for_screen().
/*

C record/class : GtkStyleProvider
*/
type StyleProvider struct {
	native *C.GtkStyleProvider
}

func StyleProviderNewFromC(u unsafe.Pointer) *StyleProvider {
	c := (*C.GtkStyleProvider)(u)
	if c == nil {
		return nil
	}

	g := &StyleProvider{native: c}

	return g
}

func (recv *StyleProvider) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GtkToolShell interface allows container widgets to provide additional
// information when embedding #GtkToolItem widgets.
/*

C record/class : GtkToolShell
*/
type ToolShell struct {
	native *C.GtkToolShell
}

func ToolShellNewFromC(u unsafe.Pointer) *ToolShell {
	c := (*C.GtkToolShell)(u)
	if c == nil {
		return nil
	}

	g := &ToolShell{native: c}

	return g
}

func (recv *ToolShell) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C record/class : GtkTreeDragDest
*/
type TreeDragDest struct {
	native *C.GtkTreeDragDest
}

func TreeDragDestNewFromC(u unsafe.Pointer) *TreeDragDest {
	c := (*C.GtkTreeDragDest)(u)
	if c == nil {
		return nil
	}

	g := &TreeDragDest{native: c}

	return g
}

func (recv *TreeDragDest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Asks the #GtkTreeDragDest to insert a row before the path @dest,
// deriving the contents of the row from @selection_data. If @dest is
// outside the tree so that inserting before it is impossible, %FALSE
// will be returned. Also, %FALSE may be returned if the new row is
// not created for some model-specific reason.  Should robustly handle
// a @dest no longer found in the model!
/*

C function : gtk_tree_drag_dest_drag_data_received
*/
func (recv *TreeDragDest) DragDataReceived(dest *TreePath, selectionData *SelectionData) bool {
	c_dest := (*C.GtkTreePath)(C.NULL)
	if dest != nil {
		c_dest = (*C.GtkTreePath)(dest.ToC())
	}

	c_selection_data := (*C.GtkSelectionData)(C.NULL)
	if selectionData != nil {
		c_selection_data = (*C.GtkSelectionData)(selectionData.ToC())
	}

	retC := C.gtk_tree_drag_dest_drag_data_received((*C.GtkTreeDragDest)(recv.native), c_dest, c_selection_data)
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether a drop is possible before the given @dest_path,
// at the same depth as @dest_path. i.e., can we drop the data in
// @selection_data at that location. @dest_path does not have to
// exist; the return value will almost certainly be %FALSE if the
// parent of @dest_path doesn’t exist, though.
/*

C function : gtk_tree_drag_dest_row_drop_possible
*/
func (recv *TreeDragDest) RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool {
	c_dest_path := (*C.GtkTreePath)(C.NULL)
	if destPath != nil {
		c_dest_path = (*C.GtkTreePath)(destPath.ToC())
	}

	c_selection_data := (*C.GtkSelectionData)(C.NULL)
	if selectionData != nil {
		c_selection_data = (*C.GtkSelectionData)(selectionData.ToC())
	}

	retC := C.gtk_tree_drag_dest_row_drop_possible((*C.GtkTreeDragDest)(recv.native), c_dest_path, c_selection_data)
	retGo := retC == C.TRUE

	return retGo
}

/*

C record/class : GtkTreeDragSource
*/
type TreeDragSource struct {
	native *C.GtkTreeDragSource
}

func TreeDragSourceNewFromC(u unsafe.Pointer) *TreeDragSource {
	c := (*C.GtkTreeDragSource)(u)
	if c == nil {
		return nil
	}

	g := &TreeDragSource{native: c}

	return g
}

func (recv *TreeDragSource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Asks the #GtkTreeDragSource to delete the row at @path, because
// it was moved somewhere else via drag-and-drop. Returns %FALSE
// if the deletion fails because @path no longer exists, or for
// some model-specific reason. Should robustly handle a @path no
// longer found in the model!
/*

C function : gtk_tree_drag_source_drag_data_delete
*/
func (recv *TreeDragSource) DragDataDelete(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_tree_drag_source_drag_data_delete((*C.GtkTreeDragSource)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// Asks the #GtkTreeDragSource to fill in @selection_data with a
// representation of the row at @path. @selection_data->target gives
// the required type of the data.  Should robustly handle a @path no
// longer found in the model!
/*

C function : gtk_tree_drag_source_drag_data_get
*/
func (recv *TreeDragSource) DragDataGet(path *TreePath, selectionData *SelectionData) bool {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_selection_data := (*C.GtkSelectionData)(C.NULL)
	if selectionData != nil {
		c_selection_data = (*C.GtkSelectionData)(selectionData.ToC())
	}

	retC := C.gtk_tree_drag_source_drag_data_get((*C.GtkTreeDragSource)(recv.native), c_path, c_selection_data)
	retGo := retC == C.TRUE

	return retGo
}

// Asks the #GtkTreeDragSource whether a particular row can be used as
// the source of a DND operation. If the source doesn’t implement
// this interface, the row is assumed draggable.
/*

C function : gtk_tree_drag_source_row_draggable
*/
func (recv *TreeDragSource) RowDraggable(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_tree_drag_source_row_draggable((*C.GtkTreeDragSource)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// The #GtkTreeModel interface defines a generic tree interface for
// use by the #GtkTreeView widget. It is an abstract interface, and
// is designed to be usable with any appropriate data structure. The
// programmer just has to implement this interface on their own data
// type for it to be viewable by a #GtkTreeView widget.
//
// The model is represented as a hierarchical tree of strongly-typed,
// columned data. In other words, the model can be seen as a tree where
// every node has different values depending on which column is being
// queried. The type of data found in a column is determined by using
// the GType system (ie. #G_TYPE_INT, #GTK_TYPE_BUTTON, #G_TYPE_POINTER,
// etc). The types are homogeneous per column across all nodes. It is
// important to note that this interface only provides a way of examining
// a model and observing changes. The implementation of each individual
// model decides how and if changes are made.
//
// In order to make life simpler for programmers who do not need to
// write their own specialized model, two generic models are provided
// — the #GtkTreeStore and the #GtkListStore. To use these, the
// developer simply pushes data into these models as necessary. These
// models provide the data structure as well as all appropriate tree
// interfaces. As a result, implementing drag and drop, sorting, and
// storing data is trivial. For the vast majority of trees and lists,
// these two models are sufficient.
//
// Models are accessed on a node/column level of granularity. One can
// query for the value of a model at a certain node and a certain
// column on that node. There are two structures used to reference a
// particular node in a model. They are the #GtkTreePath-struct and
// the #GtkTreeIter-struct (“iter” is short for iterator). Most of the
// interface consists of operations on a #GtkTreeIter-struct.
//
// A path is essentially a potential node. It is a location on a model
// that may or may not actually correspond to a node on a specific
// model. The #GtkTreePath-struct can be converted into either an
// array of unsigned integers or a string. The string form is a list
// of numbers separated by a colon. Each number refers to the offset
// at that level. Thus, the path `0` refers to the root
// node and the path `2:4` refers to the fifth child of
// the third node.
//
// By contrast, a #GtkTreeIter-struct is a reference to a specific node on
// a specific model. It is a generic struct with an integer and three
// generic pointers. These are filled in by the model in a model-specific
// way. One can convert a path to an iterator by calling
// gtk_tree_model_get_iter(). These iterators are the primary way
// of accessing a model and are similar to the iterators used by
// #GtkTextBuffer. They are generally statically allocated on the
// stack and only used for a short time. The model interface defines
// a set of operations using them for navigating the model.
//
// It is expected that models fill in the iterator with private data.
// For example, the #GtkListStore model, which is internally a simple
// linked list, stores a list node in one of the pointers. The
// #GtkTreeModelSort stores an array and an offset in two of the
// pointers. Additionally, there is an integer field. This field is
// generally filled with a unique stamp per model. This stamp is for
// catching errors resulting from using invalid iterators with a model.
//
// The lifecycle of an iterator can be a little confusing at first.
// Iterators are expected to always be valid for as long as the model
// is unchanged (and doesn’t emit a signal). The model is considered
// to own all outstanding iterators and nothing needs to be done to
// free them from the user’s point of view. Additionally, some models
// guarantee that an iterator is valid for as long as the node it refers
// to is valid (most notably the #GtkTreeStore and #GtkListStore).
// Although generally uninteresting, as one always has to allow for
// the case where iterators do not persist beyond a signal, some very
// important performance enhancements were made in the sort model.
// As a result, the #GTK_TREE_MODEL_ITERS_PERSIST flag was added to
// indicate this behavior.
//
// To help show some common operation of a model, some examples are
// provided. The first example shows three ways of getting the iter at
// the location `3:2:5`. While the first method shown is
// easier, the second is much more common, as you often get paths from
// callbacks.
//
// ## Acquiring a #GtkTreeIter-struct
//
// |[<!-- language="C" -->
// Three ways of getting the iter pointing to the location
// GtkTreePath *path;
// GtkTreeIter iter;
// GtkTreeIter parent_iter;
//
// get the iterator from a string
// gtk_tree_model_get_iter_from_string (model,
// &iter,
// "3:2:5");
//
// get the iterator from a path
// path = gtk_tree_path_new_from_string ("3:2:5");
// gtk_tree_model_get_iter (model, &iter, path);
// gtk_tree_path_free (path);
//
// walk the tree to find the iterator
// gtk_tree_model_iter_nth_child (model, &iter,
// NULL, 3);
// parent_iter = iter;
// gtk_tree_model_iter_nth_child (model, &iter,
// &parent_iter, 2);
// parent_iter = iter;
// gtk_tree_model_iter_nth_child (model, &iter,
// &parent_iter, 5);
// ]|
//
// This second example shows a quick way of iterating through a list
// and getting a string and an integer from each row. The
// populate_model() function used below is not
// shown, as it is specific to the #GtkListStore. For information on
// how to write such a function, see the #GtkListStore documentation.
//
// ## Reading data from a #GtkTreeModel
//
// |[<!-- language="C" -->
// enum
// {
// STRING_COLUMN,
// INT_COLUMN,
// N_COLUMNS
// };
//
// ...
//
// GtkTreeModel *list_store;
// GtkTreeIter iter;
// gboolean valid;
// gint row_count = 0;
//
// make a new list_store
// list_store = gtk_list_store_new (N_COLUMNS,
// G_TYPE_STRING,
// G_TYPE_INT);
//
// Fill the list store with data
// populate_model (list_store);
//
// Get the first iter in the list, check it is valid and walk
// through the list, reading each row.
//
// valid = gtk_tree_model_get_iter_first (list_store,
// &iter);
// while (valid)
// {
// gchar *str_data;
// gint   int_data;
//
// Make sure you terminate calls to gtk_tree_model_get() with a “-1” value
// gtk_tree_model_get (list_store, &iter,
// STRING_COLUMN, &str_data,
// INT_COLUMN, &int_data,
// -1);
//
// Do something with the data
// g_print ("Row %d: (%s,%d)\n",
// row_count, str_data, int_data);
// g_free (str_data);
//
// valid = gtk_tree_model_iter_next (list_store,
// &iter);
// row_count++;
// }
// ]|
//
// The #GtkTreeModel interface contains two methods for reference
// counting: gtk_tree_model_ref_node() and gtk_tree_model_unref_node().
// These two methods are optional to implement. The reference counting
// is meant as a way for views to let models know when nodes are being
// displayed. #GtkTreeView will take a reference on a node when it is
// visible, which means the node is either in the toplevel or expanded.
// Being displayed does not mean that the node is currently directly
// visible to the user in the viewport. Based on this reference counting
// scheme a caching model, for example, can decide whether or not to cache
// a node based on the reference count. A file-system based model would
// not want to keep the entire file hierarchy in memory, but just the
// folders that are currently expanded in every current view.
//
// When working with reference counting, the following rules must be taken
// into account:
//
// - Never take a reference on a node without owning a reference on its parent.
// This means that all parent nodes of a referenced node must be referenced
// as well.
//
// - Outstanding references on a deleted node are not released. This is not
// possible because the node has already been deleted by the time the
// row-deleted signal is received.
//
// - Models are not obligated to emit a signal on rows of which none of its
// siblings are referenced. To phrase this differently, signals are only
// required for levels in which nodes are referenced. For the root level
// however, signals must be emitted at all times (however the root level
// is always referenced when any view is attached).
/*

C record/class : GtkTreeModel
*/
type TreeModel struct {
	native *C.GtkTreeModel
}

func TreeModelNewFromC(u unsafe.Pointer) *TreeModel {
	c := (*C.GtkTreeModel)(u)
	if c == nil {
		return nil
	}

	g := &TreeModel{native: c}

	return g
}

func (recv *TreeModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalTreeModelRowChangedDetail struct {
	callback  TreeModelSignalRowChangedCallback
	handlerID C.gulong
}

var signalTreeModelRowChangedId int
var signalTreeModelRowChangedMap = make(map[int]signalTreeModelRowChangedDetail)
var signalTreeModelRowChangedLock sync.Mutex

// TreeModelSignalRowChangedCallback is a callback function for a 'row-changed' signal emitted from a TreeModel.
type TreeModelSignalRowChangedCallback func(path *TreePath, iter *TreeIter)

/*
ConnectRowChanged connects the callback to the 'row-changed' signal for the TreeModel.

The returned value represents the connection, and may be passed to DisconnectRowChanged to remove it.
*/
func (recv *TreeModel) ConnectRowChanged(callback TreeModelSignalRowChangedCallback) int {
	signalTreeModelRowChangedLock.Lock()
	defer signalTreeModelRowChangedLock.Unlock()

	signalTreeModelRowChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.TreeModel_signal_connect_row_changed(instance, C.gpointer(uintptr(signalTreeModelRowChangedId)))

	detail := signalTreeModelRowChangedDetail{callback, handlerID}
	signalTreeModelRowChangedMap[signalTreeModelRowChangedId] = detail

	return signalTreeModelRowChangedId
}

/*
DisconnectRowChanged disconnects a callback from the 'row-changed' signal for the TreeModel.

The connectionID should be a value returned from a call to ConnectRowChanged.
*/
func (recv *TreeModel) DisconnectRowChanged(connectionID int) {
	signalTreeModelRowChangedLock.Lock()
	defer signalTreeModelRowChangedLock.Unlock()

	detail, exists := signalTreeModelRowChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTreeModelRowChangedMap, connectionID)
}

//export treemodel_rowChangedHandler
func treemodel_rowChangedHandler(_ *C.GObject, c_path *C.GtkTreePath, c_iter *C.GtkTreeIter, data C.gpointer) {
	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalTreeModelRowChangedMap[index].callback
	callback(path, iter)
}

type signalTreeModelRowDeletedDetail struct {
	callback  TreeModelSignalRowDeletedCallback
	handlerID C.gulong
}

var signalTreeModelRowDeletedId int
var signalTreeModelRowDeletedMap = make(map[int]signalTreeModelRowDeletedDetail)
var signalTreeModelRowDeletedLock sync.Mutex

// TreeModelSignalRowDeletedCallback is a callback function for a 'row-deleted' signal emitted from a TreeModel.
type TreeModelSignalRowDeletedCallback func(path *TreePath)

/*
ConnectRowDeleted connects the callback to the 'row-deleted' signal for the TreeModel.

The returned value represents the connection, and may be passed to DisconnectRowDeleted to remove it.
*/
func (recv *TreeModel) ConnectRowDeleted(callback TreeModelSignalRowDeletedCallback) int {
	signalTreeModelRowDeletedLock.Lock()
	defer signalTreeModelRowDeletedLock.Unlock()

	signalTreeModelRowDeletedId++
	instance := C.gpointer(recv.native)
	handlerID := C.TreeModel_signal_connect_row_deleted(instance, C.gpointer(uintptr(signalTreeModelRowDeletedId)))

	detail := signalTreeModelRowDeletedDetail{callback, handlerID}
	signalTreeModelRowDeletedMap[signalTreeModelRowDeletedId] = detail

	return signalTreeModelRowDeletedId
}

/*
DisconnectRowDeleted disconnects a callback from the 'row-deleted' signal for the TreeModel.

The connectionID should be a value returned from a call to ConnectRowDeleted.
*/
func (recv *TreeModel) DisconnectRowDeleted(connectionID int) {
	signalTreeModelRowDeletedLock.Lock()
	defer signalTreeModelRowDeletedLock.Unlock()

	detail, exists := signalTreeModelRowDeletedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTreeModelRowDeletedMap, connectionID)
}

//export treemodel_rowDeletedHandler
func treemodel_rowDeletedHandler(_ *C.GObject, c_path *C.GtkTreePath, data C.gpointer) {
	path := TreePathNewFromC(unsafe.Pointer(c_path))

	index := int(uintptr(data))
	callback := signalTreeModelRowDeletedMap[index].callback
	callback(path)
}

type signalTreeModelRowHasChildToggledDetail struct {
	callback  TreeModelSignalRowHasChildToggledCallback
	handlerID C.gulong
}

var signalTreeModelRowHasChildToggledId int
var signalTreeModelRowHasChildToggledMap = make(map[int]signalTreeModelRowHasChildToggledDetail)
var signalTreeModelRowHasChildToggledLock sync.Mutex

// TreeModelSignalRowHasChildToggledCallback is a callback function for a 'row-has-child-toggled' signal emitted from a TreeModel.
type TreeModelSignalRowHasChildToggledCallback func(path *TreePath, iter *TreeIter)

/*
ConnectRowHasChildToggled connects the callback to the 'row-has-child-toggled' signal for the TreeModel.

The returned value represents the connection, and may be passed to DisconnectRowHasChildToggled to remove it.
*/
func (recv *TreeModel) ConnectRowHasChildToggled(callback TreeModelSignalRowHasChildToggledCallback) int {
	signalTreeModelRowHasChildToggledLock.Lock()
	defer signalTreeModelRowHasChildToggledLock.Unlock()

	signalTreeModelRowHasChildToggledId++
	instance := C.gpointer(recv.native)
	handlerID := C.TreeModel_signal_connect_row_has_child_toggled(instance, C.gpointer(uintptr(signalTreeModelRowHasChildToggledId)))

	detail := signalTreeModelRowHasChildToggledDetail{callback, handlerID}
	signalTreeModelRowHasChildToggledMap[signalTreeModelRowHasChildToggledId] = detail

	return signalTreeModelRowHasChildToggledId
}

/*
DisconnectRowHasChildToggled disconnects a callback from the 'row-has-child-toggled' signal for the TreeModel.

The connectionID should be a value returned from a call to ConnectRowHasChildToggled.
*/
func (recv *TreeModel) DisconnectRowHasChildToggled(connectionID int) {
	signalTreeModelRowHasChildToggledLock.Lock()
	defer signalTreeModelRowHasChildToggledLock.Unlock()

	detail, exists := signalTreeModelRowHasChildToggledMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTreeModelRowHasChildToggledMap, connectionID)
}

//export treemodel_rowHasChildToggledHandler
func treemodel_rowHasChildToggledHandler(_ *C.GObject, c_path *C.GtkTreePath, c_iter *C.GtkTreeIter, data C.gpointer) {
	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalTreeModelRowHasChildToggledMap[index].callback
	callback(path, iter)
}

type signalTreeModelRowInsertedDetail struct {
	callback  TreeModelSignalRowInsertedCallback
	handlerID C.gulong
}

var signalTreeModelRowInsertedId int
var signalTreeModelRowInsertedMap = make(map[int]signalTreeModelRowInsertedDetail)
var signalTreeModelRowInsertedLock sync.Mutex

// TreeModelSignalRowInsertedCallback is a callback function for a 'row-inserted' signal emitted from a TreeModel.
type TreeModelSignalRowInsertedCallback func(path *TreePath, iter *TreeIter)

/*
ConnectRowInserted connects the callback to the 'row-inserted' signal for the TreeModel.

The returned value represents the connection, and may be passed to DisconnectRowInserted to remove it.
*/
func (recv *TreeModel) ConnectRowInserted(callback TreeModelSignalRowInsertedCallback) int {
	signalTreeModelRowInsertedLock.Lock()
	defer signalTreeModelRowInsertedLock.Unlock()

	signalTreeModelRowInsertedId++
	instance := C.gpointer(recv.native)
	handlerID := C.TreeModel_signal_connect_row_inserted(instance, C.gpointer(uintptr(signalTreeModelRowInsertedId)))

	detail := signalTreeModelRowInsertedDetail{callback, handlerID}
	signalTreeModelRowInsertedMap[signalTreeModelRowInsertedId] = detail

	return signalTreeModelRowInsertedId
}

/*
DisconnectRowInserted disconnects a callback from the 'row-inserted' signal for the TreeModel.

The connectionID should be a value returned from a call to ConnectRowInserted.
*/
func (recv *TreeModel) DisconnectRowInserted(connectionID int) {
	signalTreeModelRowInsertedLock.Lock()
	defer signalTreeModelRowInsertedLock.Unlock()

	detail, exists := signalTreeModelRowInsertedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTreeModelRowInsertedMap, connectionID)
}

//export treemodel_rowInsertedHandler
func treemodel_rowInsertedHandler(_ *C.GObject, c_path *C.GtkTreePath, c_iter *C.GtkTreeIter, data C.gpointer) {
	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalTreeModelRowInsertedMap[index].callback
	callback(path, iter)
}

// Unsupported signal 'rows-reordered' for TreeModel : unsupported parameter new_order : type gpointer :

// Unsupported : gtk_tree_model_foreach : unsupported parameter func : no type generator for TreeModelForeachFunc (GtkTreeModelForeachFunc) for param func

// Unsupported : gtk_tree_model_get : unsupported parameter ... : varargs

// Returns the type of the column.
/*

C function : gtk_tree_model_get_column_type
*/
func (recv *TreeModel) GetColumnType(index int32) gobject.Type {
	c_index_ := (C.gint)(index)

	retC := C.gtk_tree_model_get_column_type((*C.GtkTreeModel)(recv.native), c_index_)
	retGo := (gobject.Type)(retC)

	return retGo
}

// Returns a set of flags supported by this interface.
//
// The flags are a bitwise combination of #GtkTreeModelFlags.
// The flags supported should not change during the lifetime
// of the @tree_model.
/*

C function : gtk_tree_model_get_flags
*/
func (recv *TreeModel) GetFlags() TreeModelFlags {
	retC := C.gtk_tree_model_get_flags((*C.GtkTreeModel)(recv.native))
	retGo := (TreeModelFlags)(retC)

	return retGo
}

// Sets @iter to a valid iterator pointing to @path.  If @path does
// not exist, @iter is set to an invalid iterator and %FALSE is returned.
/*

C function : gtk_tree_model_get_iter
*/
func (recv *TreeModel) GetIter(path *TreePath) (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_tree_model_get_iter((*C.GtkTreeModel)(recv.native), &c_iter, c_path)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// Initializes @iter with the first iterator in the tree
// (the one at the path "0") and returns %TRUE. Returns
// %FALSE if the tree is empty.
/*

C function : gtk_tree_model_get_iter_first
*/
func (recv *TreeModel) GetIterFirst() (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	retC := C.gtk_tree_model_get_iter_first((*C.GtkTreeModel)(recv.native), &c_iter)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// Sets @iter to a valid iterator pointing to @path_string, if it
// exists. Otherwise, @iter is left invalid and %FALSE is returned.
/*

C function : gtk_tree_model_get_iter_from_string
*/
func (recv *TreeModel) GetIterFromString(pathString string) (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	c_path_string := C.CString(pathString)
	defer C.free(unsafe.Pointer(c_path_string))

	retC := C.gtk_tree_model_get_iter_from_string((*C.GtkTreeModel)(recv.native), &c_iter, c_path_string)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// Returns the number of columns supported by @tree_model.
/*

C function : gtk_tree_model_get_n_columns
*/
func (recv *TreeModel) GetNColumns() int32 {
	retC := C.gtk_tree_model_get_n_columns((*C.GtkTreeModel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns a newly-created #GtkTreePath-struct referenced by @iter.
//
// This path should be freed with gtk_tree_path_free().
/*

C function : gtk_tree_model_get_path
*/
func (recv *TreeModel) GetPath(iter *TreeIter) *TreePath {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_get_path((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_model_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Initializes and sets @value to that at @column.
//
// When done with @value, g_value_unset() needs to be called
// to free any allocated memory.
/*

C function : gtk_tree_model_get_value
*/
func (recv *TreeModel) GetValue(iter *TreeIter, column int32) *gobject.Value {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_column := (C.gint)(column)

	var c_value C.GValue

	C.gtk_tree_model_get_value((*C.GtkTreeModel)(recv.native), c_iter, c_column, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// Sets @iter to point to the first child of @parent.
//
// If @parent has no children, %FALSE is returned and @iter is
// set to be invalid. @parent will remain a valid node after this
// function has been called.
//
// If @parent is %NULL returns the first node, equivalent to
// `gtk_tree_model_get_iter_first (tree_model, iter);`
/*

C function : gtk_tree_model_iter_children
*/
func (recv *TreeModel) IterChildren(parent *TreeIter) (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	c_parent := (*C.GtkTreeIter)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkTreeIter)(parent.ToC())
	}

	retC := C.gtk_tree_model_iter_children((*C.GtkTreeModel)(recv.native), &c_iter, c_parent)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// Returns %TRUE if @iter has children, %FALSE otherwise.
/*

C function : gtk_tree_model_iter_has_child
*/
func (recv *TreeModel) IterHasChild(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_iter_has_child((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Returns the number of children that @iter has.
//
// As a special case, if @iter is %NULL, then the number
// of toplevel nodes is returned.
/*

C function : gtk_tree_model_iter_n_children
*/
func (recv *TreeModel) IterNChildren(iter *TreeIter) int32 {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_iter_n_children((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := (int32)(retC)

	return retGo
}

// Sets @iter to point to the node following it at the current level.
//
// If there is no next @iter, %FALSE is returned and @iter is set
// to be invalid.
/*

C function : gtk_tree_model_iter_next
*/
func (recv *TreeModel) IterNext(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_iter_next((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Sets @iter to be the child of @parent, using the given index.
//
// The first index is 0. If @n is too big, or @parent has no children,
// @iter is set to an invalid iterator and %FALSE is returned. @parent
// will remain a valid node after this function has been called. As a
// special case, if @parent is %NULL, then the @n-th root node
// is set.
/*

C function : gtk_tree_model_iter_nth_child
*/
func (recv *TreeModel) IterNthChild(parent *TreeIter, n int32) (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	c_parent := (*C.GtkTreeIter)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkTreeIter)(parent.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.gtk_tree_model_iter_nth_child((*C.GtkTreeModel)(recv.native), &c_iter, c_parent, c_n)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// Sets @iter to be the parent of @child.
//
// If @child is at the toplevel, and doesn’t have a parent, then
// @iter is set to an invalid iterator and %FALSE is returned.
// @child will remain a valid node after this function has been
// called.
//
// @iter will be initialized before the lookup is performed, so @child
// and @iter cannot point to the same memory location.
/*

C function : gtk_tree_model_iter_parent
*/
func (recv *TreeModel) IterParent(child *TreeIter) (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	c_child := (*C.GtkTreeIter)(C.NULL)
	if child != nil {
		c_child = (*C.GtkTreeIter)(child.ToC())
	}

	retC := C.gtk_tree_model_iter_parent((*C.GtkTreeModel)(recv.native), &c_iter, c_child)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// Lets the tree ref the node.
//
// This is an optional method for models to implement.
// To be more specific, models may ignore this call as it exists
// primarily for performance reasons.
//
// This function is primarily meant as a way for views to let
// caching models know when nodes are being displayed (and hence,
// whether or not to cache that node). Being displayed means a node
// is in an expanded branch, regardless of whether the node is currently
// visible in the viewport. For example, a file-system based model
// would not want to keep the entire file-hierarchy in memory,
// just the sections that are currently being displayed by
// every current view.
//
// A model should be expected to be able to get an iter independent
// of its reffed state.
/*

C function : gtk_tree_model_ref_node
*/
func (recv *TreeModel) RefNode(iter *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_tree_model_ref_node((*C.GtkTreeModel)(recv.native), c_iter)

	return
}

// Emits the #GtkTreeModel::row-changed signal on @tree_model.
/*

C function : gtk_tree_model_row_changed
*/
func (recv *TreeModel) RowChanged(path *TreePath, iter *TreeIter) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_tree_model_row_changed((*C.GtkTreeModel)(recv.native), c_path, c_iter)

	return
}

// Emits the #GtkTreeModel::row-deleted signal on @tree_model.
//
// This should be called by models after a row has been removed.
// The location pointed to by @path should be the location that
// the row previously was at. It may not be a valid location anymore.
//
// Nodes that are deleted are not unreffed, this means that any
// outstanding references on the deleted node should not be released.
/*

C function : gtk_tree_model_row_deleted
*/
func (recv *TreeModel) RowDeleted(path *TreePath) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_tree_model_row_deleted((*C.GtkTreeModel)(recv.native), c_path)

	return
}

// Emits the #GtkTreeModel::row-has-child-toggled signal on
// @tree_model. This should be called by models after the child
// state of a node changes.
/*

C function : gtk_tree_model_row_has_child_toggled
*/
func (recv *TreeModel) RowHasChildToggled(path *TreePath, iter *TreeIter) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_tree_model_row_has_child_toggled((*C.GtkTreeModel)(recv.native), c_path, c_iter)

	return
}

// Emits the #GtkTreeModel::row-inserted signal on @tree_model.
/*

C function : gtk_tree_model_row_inserted
*/
func (recv *TreeModel) RowInserted(path *TreePath, iter *TreeIter) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_tree_model_row_inserted((*C.GtkTreeModel)(recv.native), c_path, c_iter)

	return
}

// Emits the #GtkTreeModel::rows-reordered signal on @tree_model.
//
// This should be called by models when their rows have been
// reordered.
/*

C function : gtk_tree_model_rows_reordered
*/
func (recv *TreeModel) RowsReordered(path *TreePath, iter *TreeIter, newOrder int32) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_new_order := (C.gint)(newOrder)

	C.gtk_tree_model_rows_reordered((*C.GtkTreeModel)(recv.native), c_path, c_iter, &c_new_order)

	return
}

// Creates a new #GtkTreeModel, with @child_model as the child model.
/*

C function : gtk_tree_model_sort_new_with_model
*/
func (recv *TreeModel) SortNewWithModel() *TreeModel {
	retC := C.gtk_tree_model_sort_new_with_model((*C.GtkTreeModel)(recv.native))
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Lets the tree unref the node.
//
// This is an optional method for models to implement.
// To be more specific, models may ignore this call as it exists
// primarily for performance reasons. For more information on what
// this means, see gtk_tree_model_ref_node().
//
// Please note that nodes that are deleted are not unreffed.
/*

C function : gtk_tree_model_unref_node
*/
func (recv *TreeModel) UnrefNode(iter *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_tree_model_unref_node((*C.GtkTreeModel)(recv.native), c_iter)

	return
}

// #GtkTreeSortable is an interface to be implemented by tree models which
// support sorting. The #GtkTreeView uses the methods provided by this interface
// to sort the model.
/*

C record/class : GtkTreeSortable
*/
type TreeSortable struct {
	native *C.GtkTreeSortable
}

func TreeSortableNewFromC(u unsafe.Pointer) *TreeSortable {
	c := (*C.GtkTreeSortable)(u)
	if c == nil {
		return nil
	}

	g := &TreeSortable{native: c}

	return g
}

func (recv *TreeSortable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalTreeSortableSortColumnChangedDetail struct {
	callback  TreeSortableSignalSortColumnChangedCallback
	handlerID C.gulong
}

var signalTreeSortableSortColumnChangedId int
var signalTreeSortableSortColumnChangedMap = make(map[int]signalTreeSortableSortColumnChangedDetail)
var signalTreeSortableSortColumnChangedLock sync.Mutex

// TreeSortableSignalSortColumnChangedCallback is a callback function for a 'sort-column-changed' signal emitted from a TreeSortable.
type TreeSortableSignalSortColumnChangedCallback func()

/*
ConnectSortColumnChanged connects the callback to the 'sort-column-changed' signal for the TreeSortable.

The returned value represents the connection, and may be passed to DisconnectSortColumnChanged to remove it.
*/
func (recv *TreeSortable) ConnectSortColumnChanged(callback TreeSortableSignalSortColumnChangedCallback) int {
	signalTreeSortableSortColumnChangedLock.Lock()
	defer signalTreeSortableSortColumnChangedLock.Unlock()

	signalTreeSortableSortColumnChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.TreeSortable_signal_connect_sort_column_changed(instance, C.gpointer(uintptr(signalTreeSortableSortColumnChangedId)))

	detail := signalTreeSortableSortColumnChangedDetail{callback, handlerID}
	signalTreeSortableSortColumnChangedMap[signalTreeSortableSortColumnChangedId] = detail

	return signalTreeSortableSortColumnChangedId
}

/*
DisconnectSortColumnChanged disconnects a callback from the 'sort-column-changed' signal for the TreeSortable.

The connectionID should be a value returned from a call to ConnectSortColumnChanged.
*/
func (recv *TreeSortable) DisconnectSortColumnChanged(connectionID int) {
	signalTreeSortableSortColumnChangedLock.Lock()
	defer signalTreeSortableSortColumnChangedLock.Unlock()

	detail, exists := signalTreeSortableSortColumnChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTreeSortableSortColumnChangedMap, connectionID)
}

//export treesortable_sortColumnChangedHandler
func treesortable_sortColumnChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalTreeSortableSortColumnChangedMap[index].callback
	callback()
}

// Unsupported : gtk_tree_sortable_get_sort_column_id : unsupported parameter order : GtkSortType* with indirection level of 1

// Returns %TRUE if the model has a default sort function. This is used
// primarily by GtkTreeViewColumns in order to determine if a model can
// go back to the default state, or not.
/*

C function : gtk_tree_sortable_has_default_sort_func
*/
func (recv *TreeSortable) HasDefaultSortFunc() bool {
	retC := C.gtk_tree_sortable_has_default_sort_func((*C.GtkTreeSortable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_sortable_set_default_sort_func : unsupported parameter sort_func : no type generator for TreeIterCompareFunc (GtkTreeIterCompareFunc) for param sort_func

// Sets the current sort column to be @sort_column_id. The @sortable will
// resort itself to reflect this change, after emitting a
// #GtkTreeSortable::sort-column-changed signal. @sort_column_id may either be
// a regular column id, or one of the following special values:
//
// - %GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function
// will be used, if it is set
//
// - %GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur
/*

C function : gtk_tree_sortable_set_sort_column_id
*/
func (recv *TreeSortable) SetSortColumnId(sortColumnId int32, order SortType) {
	c_sort_column_id := (C.gint)(sortColumnId)

	c_order := (C.GtkSortType)(order)

	C.gtk_tree_sortable_set_sort_column_id((*C.GtkTreeSortable)(recv.native), c_sort_column_id, c_order)

	return
}

// Unsupported : gtk_tree_sortable_set_sort_func : unsupported parameter sort_func : no type generator for TreeIterCompareFunc (GtkTreeIterCompareFunc) for param sort_func

// Emits a #GtkTreeSortable::sort-column-changed signal on @sortable.
/*

C function : gtk_tree_sortable_sort_column_changed
*/
func (recv *TreeSortable) SortColumnChanged() {
	C.gtk_tree_sortable_sort_column_changed((*C.GtkTreeSortable)(recv.native))

	return
}
