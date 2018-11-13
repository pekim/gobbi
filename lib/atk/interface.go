// This is a generated file - DO NOT EDIT

package atk

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
/*

	void component_boundsChangedHandler(GObject *, AtkRectangle *, gpointer);

	static gulong Component_signal_connect_bounds_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "bounds-changed", G_CALLBACK(component_boundsChangedHandler), data);
	}

*/
/*

	void document_loadCompleteHandler(GObject *, gpointer);

	static gulong Document_signal_connect_load_complete(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "load-complete", G_CALLBACK(document_loadCompleteHandler), data);
	}

*/
/*

	void document_loadStoppedHandler(GObject *, gpointer);

	static gulong Document_signal_connect_load_stopped(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "load-stopped", G_CALLBACK(document_loadStoppedHandler), data);
	}

*/
/*

	void document_reloadHandler(GObject *, gpointer);

	static gulong Document_signal_connect_reload(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "reload", G_CALLBACK(document_reloadHandler), data);
	}

*/
/*

	void selection_selectionChangedHandler(GObject *, gpointer);

	static gulong Selection_signal_connect_selection_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-changed", G_CALLBACK(selection_selectionChangedHandler), data);
	}

*/
/*

	void table_columnReorderedHandler(GObject *, gpointer);

	static gulong Table_signal_connect_column_reordered(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "column-reordered", G_CALLBACK(table_columnReorderedHandler), data);
	}

*/
/*

	void table_modelChangedHandler(GObject *, gpointer);

	static gulong Table_signal_connect_model_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "model-changed", G_CALLBACK(table_modelChangedHandler), data);
	}

*/
/*

	void table_rowReorderedHandler(GObject *, gpointer);

	static gulong Table_signal_connect_row_reordered(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-reordered", G_CALLBACK(table_rowReorderedHandler), data);
	}

*/
/*

	void text_textAttributesChangedHandler(GObject *, gpointer);

	static gulong Text_signal_connect_text_attributes_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "text-attributes-changed", G_CALLBACK(text_textAttributesChangedHandler), data);
	}

*/
/*

	void text_textSelectionChangedHandler(GObject *, gpointer);

	static gulong Text_signal_connect_text_selection_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "text-selection-changed", G_CALLBACK(text_textSelectionChangedHandler), data);
	}

*/
import "C"

// Action is a wrapper around the C record AtkAction.
type Action struct {
	native *C.AtkAction
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.AtkAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Perform the specified action on the object.
/*

C function : atk_action_do_action
*/
func (recv *Action) DoAction(i int32) bool {
	c_i := (C.gint)(i)

	retC := C.atk_action_do_action((*C.AtkAction)(recv.native), c_i)
	retGo := retC == C.TRUE

	return retGo
}

// Returns a description of the specified action of the object.
/*

C function : atk_action_get_description
*/
func (recv *Action) GetDescription(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_action_get_description((*C.AtkAction)(recv.native), c_i)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the keybinding which can be used to activate this action, if one
// exists. The string returned should contain localized, human-readable,
// key sequences as they would appear when displayed on screen. It must
// be in the format "mnemonic;sequence;shortcut".
//
// - The mnemonic key activates the object if it is presently enabled onscreen.
// This typically corresponds to the underlined letter within the widget.
// Example: "n" in a traditional "New..." menu item or the "a" in "Apply" for
// a button.
// - The sequence is the full list of keys which invoke the action even if the
// relevant element is not currently shown on screen. For instance, for a menu
// item the sequence is the keybindings used to open the parent menus before
// invoking. The sequence string is colon-delimited. Example: "Alt+F:N" in a
// traditional "New..." menu item.
// - The shortcut, if it exists, will invoke the same action without showing
// the component or its enclosing menus or dialogs. Example: "Ctrl+N" in a
// traditional "New..." menu item.
//
// Example: For a traditional "New..." menu item, the expected return value
// would be: "N;Alt+F:N;Ctrl+N" for the English locale and "N;Alt+D:N;Strg+N"
// for the German locale. If, hypothetically, this menu item lacked a mnemonic,
// it would be represented by ";;Ctrl+N" and ";;Strg+N" respectively.
/*

C function : atk_action_get_keybinding
*/
func (recv *Action) GetKeybinding(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_action_get_keybinding((*C.AtkAction)(recv.native), c_i)
	retGo := C.GoString(retC)

	return retGo
}

// Returns the localized name of the specified action of the object.
/*

C function : atk_action_get_localized_name
*/
func (recv *Action) GetLocalizedName(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_action_get_localized_name((*C.AtkAction)(recv.native), c_i)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the number of accessible actions available on the object.
// If there are more than one, the first one is considered the
// "default" action of the object.
/*

C function : atk_action_get_n_actions
*/
func (recv *Action) GetNActions() int32 {
	retC := C.atk_action_get_n_actions((*C.AtkAction)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns a non-localized string naming the specified action of the
// object. This name is generally not descriptive of the end result
// of the action, but instead names the 'interaction type' which the
// object supports. By convention, the above strings should be used to
// represent the actions which correspond to the common point-and-click
// interaction techniques of the same name: i.e.
// "click", "press", "release", "drag", "drop", "popup", etc.
// The "popup" action should be used to pop up a context menu for the
// object, if one exists.
//
// For technical reasons, some toolkits cannot guarantee that the
// reported action is actually 'bound' to a nontrivial user event;
// i.e. the result of some actions via atk_action_do_action() may be
// NIL.
/*

C function : atk_action_get_name
*/
func (recv *Action) GetName(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_action_get_name((*C.AtkAction)(recv.native), c_i)
	retGo := C.GoString(retC)

	return retGo
}

// Sets a description of the specified action of the object.
/*

C function : atk_action_set_description
*/
func (recv *Action) SetDescription(i int32, desc string) bool {
	c_i := (C.gint)(i)

	c_desc := C.CString(desc)
	defer C.free(unsafe.Pointer(c_desc))

	retC := C.atk_action_set_description((*C.AtkAction)(recv.native), c_i, c_desc)
	retGo := retC == C.TRUE

	return retGo
}

// Component is a wrapper around the C record AtkComponent.
type Component struct {
	native *C.AtkComponent
}

func ComponentNewFromC(u unsafe.Pointer) *Component {
	c := (*C.AtkComponent)(u)
	if c == nil {
		return nil
	}

	g := &Component{native: c}

	return g
}

func (recv *Component) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalComponentBoundsChangedDetail struct {
	callback  ComponentSignalBoundsChangedCallback
	handlerID C.gulong
}

var signalComponentBoundsChangedId int
var signalComponentBoundsChangedMap = make(map[int]signalComponentBoundsChangedDetail)
var signalComponentBoundsChangedLock sync.Mutex

// ComponentSignalBoundsChangedCallback is a callback function for a 'bounds-changed' signal emitted from a Component.
type ComponentSignalBoundsChangedCallback func(arg1 *Rectangle)

/*
ConnectBoundsChanged connects the callback to the 'bounds-changed' signal for the Component.

The returned value represents the connection, and may be passed to DisconnectBoundsChanged to remove it.
*/
func (recv *Component) ConnectBoundsChanged(callback ComponentSignalBoundsChangedCallback) int {
	signalComponentBoundsChangedLock.Lock()
	defer signalComponentBoundsChangedLock.Unlock()

	signalComponentBoundsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Component_signal_connect_bounds_changed(instance, C.gpointer(uintptr(signalComponentBoundsChangedId)))

	detail := signalComponentBoundsChangedDetail{callback, handlerID}
	signalComponentBoundsChangedMap[signalComponentBoundsChangedId] = detail

	return signalComponentBoundsChangedId
}

/*
DisconnectBoundsChanged disconnects a callback from the 'bounds-changed' signal for the Component.

The connectionID should be a value returned from a call to ConnectBoundsChanged.
*/
func (recv *Component) DisconnectBoundsChanged(connectionID int) {
	signalComponentBoundsChangedLock.Lock()
	defer signalComponentBoundsChangedLock.Unlock()

	detail, exists := signalComponentBoundsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComponentBoundsChangedMap, connectionID)
}

//export component_boundsChangedHandler
func component_boundsChangedHandler(_ *C.GObject, c_arg1 *C.AtkRectangle, data C.gpointer) {
	arg1 := RectangleNewFromC(unsafe.Pointer(c_arg1))

	index := int(uintptr(data))
	callback := signalComponentBoundsChangedMap[index].callback
	callback(arg1)
}

// Unsupported : atk_component_add_focus_handler : unsupported parameter handler : no type generator for FocusHandler (AtkFocusHandler) for param handler

// Checks whether the specified point is within the extent of the @component.
//
// Toolkit implementor note: ATK provides a default implementation for
// this virtual method. In general there are little reason to
// re-implement it.
/*

C function : atk_component_contains
*/
func (recv *Component) Contains(x int32, y int32, coordType CoordType) bool {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coord_type := (C.AtkCoordType)(coordType)

	retC := C.atk_component_contains((*C.AtkComponent)(recv.native), c_x, c_y, c_coord_type)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the rectangle which gives the extent of the @component.
/*

C function : atk_component_get_extents
*/
func (recv *Component) GetExtents(coordType CoordType) (int32, int32, int32, int32) {
	var c_x C.gint

	var c_y C.gint

	var c_width C.gint

	var c_height C.gint

	c_coord_type := (C.AtkCoordType)(coordType)

	C.atk_component_get_extents((*C.AtkComponent)(recv.native), &c_x, &c_y, &c_width, &c_height, c_coord_type)

	x := (int32)(c_x)

	y := (int32)(c_y)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return x, y, width, height
}

// Gets the layer of the component.
/*

C function : atk_component_get_layer
*/
func (recv *Component) GetLayer() Layer {
	retC := C.atk_component_get_layer((*C.AtkComponent)(recv.native))
	retGo := (Layer)(retC)

	return retGo
}

// Gets the zorder of the component. The value G_MININT will be returned
// if the layer of the component is not ATK_LAYER_MDI or ATK_LAYER_WINDOW.
/*

C function : atk_component_get_mdi_zorder
*/
func (recv *Component) GetMdiZorder() int32 {
	retC := C.atk_component_get_mdi_zorder((*C.AtkComponent)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the position of @component in the form of
// a point specifying @component's top-left corner.
/*

C function : atk_component_get_position
*/
func (recv *Component) GetPosition(coordType CoordType) (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	c_coord_type := (C.AtkCoordType)(coordType)

	C.atk_component_get_position((*C.AtkComponent)(recv.native), &c_x, &c_y, c_coord_type)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// Gets the size of the @component in terms of width and height.
/*

C function : atk_component_get_size
*/
func (recv *Component) GetSize() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.atk_component_get_size((*C.AtkComponent)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// Grabs focus for this @component.
/*

C function : atk_component_grab_focus
*/
func (recv *Component) GrabFocus() bool {
	retC := C.atk_component_grab_focus((*C.AtkComponent)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets a reference to the accessible child, if one exists, at the
// coordinate point specified by @x and @y.
/*

C function : atk_component_ref_accessible_at_point
*/
func (recv *Component) RefAccessibleAtPoint(x int32, y int32, coordType CoordType) *Object {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coord_type := (C.AtkCoordType)(coordType)

	retC := C.atk_component_ref_accessible_at_point((*C.AtkComponent)(recv.native), c_x, c_y, c_coord_type)
	var retGo (*Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Remove the handler specified by @handler_id from the list of
// functions to be executed when this object receives focus events
// (in or out).
/*

C function : atk_component_remove_focus_handler
*/
func (recv *Component) RemoveFocusHandler(handlerId uint32) {
	c_handler_id := (C.guint)(handlerId)

	C.atk_component_remove_focus_handler((*C.AtkComponent)(recv.native), c_handler_id)

	return
}

// Sets the extents of @component.
/*

C function : atk_component_set_extents
*/
func (recv *Component) SetExtents(x int32, y int32, width int32, height int32, coordType CoordType) bool {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_coord_type := (C.AtkCoordType)(coordType)

	retC := C.atk_component_set_extents((*C.AtkComponent)(recv.native), c_x, c_y, c_width, c_height, c_coord_type)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the postition of @component.
/*

C function : atk_component_set_position
*/
func (recv *Component) SetPosition(x int32, y int32, coordType CoordType) bool {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coord_type := (C.AtkCoordType)(coordType)

	retC := C.atk_component_set_position((*C.AtkComponent)(recv.native), c_x, c_y, c_coord_type)
	retGo := retC == C.TRUE

	return retGo
}

// Set the size of the @component in terms of width and height.
/*

C function : atk_component_set_size
*/
func (recv *Component) SetSize(width int32, height int32) bool {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.atk_component_set_size((*C.AtkComponent)(recv.native), c_width, c_height)
	retGo := retC == C.TRUE

	return retGo
}

// Document is a wrapper around the C record AtkDocument.
type Document struct {
	native *C.AtkDocument
}

func DocumentNewFromC(u unsafe.Pointer) *Document {
	c := (*C.AtkDocument)(u)
	if c == nil {
		return nil
	}

	g := &Document{native: c}

	return g
}

func (recv *Document) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalDocumentLoadCompleteDetail struct {
	callback  DocumentSignalLoadCompleteCallback
	handlerID C.gulong
}

var signalDocumentLoadCompleteId int
var signalDocumentLoadCompleteMap = make(map[int]signalDocumentLoadCompleteDetail)
var signalDocumentLoadCompleteLock sync.Mutex

// DocumentSignalLoadCompleteCallback is a callback function for a 'load-complete' signal emitted from a Document.
type DocumentSignalLoadCompleteCallback func()

/*
ConnectLoadComplete connects the callback to the 'load-complete' signal for the Document.

The returned value represents the connection, and may be passed to DisconnectLoadComplete to remove it.
*/
func (recv *Document) ConnectLoadComplete(callback DocumentSignalLoadCompleteCallback) int {
	signalDocumentLoadCompleteLock.Lock()
	defer signalDocumentLoadCompleteLock.Unlock()

	signalDocumentLoadCompleteId++
	instance := C.gpointer(recv.native)
	handlerID := C.Document_signal_connect_load_complete(instance, C.gpointer(uintptr(signalDocumentLoadCompleteId)))

	detail := signalDocumentLoadCompleteDetail{callback, handlerID}
	signalDocumentLoadCompleteMap[signalDocumentLoadCompleteId] = detail

	return signalDocumentLoadCompleteId
}

/*
DisconnectLoadComplete disconnects a callback from the 'load-complete' signal for the Document.

The connectionID should be a value returned from a call to ConnectLoadComplete.
*/
func (recv *Document) DisconnectLoadComplete(connectionID int) {
	signalDocumentLoadCompleteLock.Lock()
	defer signalDocumentLoadCompleteLock.Unlock()

	detail, exists := signalDocumentLoadCompleteMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDocumentLoadCompleteMap, connectionID)
}

//export document_loadCompleteHandler
func document_loadCompleteHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalDocumentLoadCompleteMap[index].callback
	callback()
}

type signalDocumentLoadStoppedDetail struct {
	callback  DocumentSignalLoadStoppedCallback
	handlerID C.gulong
}

var signalDocumentLoadStoppedId int
var signalDocumentLoadStoppedMap = make(map[int]signalDocumentLoadStoppedDetail)
var signalDocumentLoadStoppedLock sync.Mutex

// DocumentSignalLoadStoppedCallback is a callback function for a 'load-stopped' signal emitted from a Document.
type DocumentSignalLoadStoppedCallback func()

/*
ConnectLoadStopped connects the callback to the 'load-stopped' signal for the Document.

The returned value represents the connection, and may be passed to DisconnectLoadStopped to remove it.
*/
func (recv *Document) ConnectLoadStopped(callback DocumentSignalLoadStoppedCallback) int {
	signalDocumentLoadStoppedLock.Lock()
	defer signalDocumentLoadStoppedLock.Unlock()

	signalDocumentLoadStoppedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Document_signal_connect_load_stopped(instance, C.gpointer(uintptr(signalDocumentLoadStoppedId)))

	detail := signalDocumentLoadStoppedDetail{callback, handlerID}
	signalDocumentLoadStoppedMap[signalDocumentLoadStoppedId] = detail

	return signalDocumentLoadStoppedId
}

/*
DisconnectLoadStopped disconnects a callback from the 'load-stopped' signal for the Document.

The connectionID should be a value returned from a call to ConnectLoadStopped.
*/
func (recv *Document) DisconnectLoadStopped(connectionID int) {
	signalDocumentLoadStoppedLock.Lock()
	defer signalDocumentLoadStoppedLock.Unlock()

	detail, exists := signalDocumentLoadStoppedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDocumentLoadStoppedMap, connectionID)
}

//export document_loadStoppedHandler
func document_loadStoppedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalDocumentLoadStoppedMap[index].callback
	callback()
}

type signalDocumentReloadDetail struct {
	callback  DocumentSignalReloadCallback
	handlerID C.gulong
}

var signalDocumentReloadId int
var signalDocumentReloadMap = make(map[int]signalDocumentReloadDetail)
var signalDocumentReloadLock sync.Mutex

// DocumentSignalReloadCallback is a callback function for a 'reload' signal emitted from a Document.
type DocumentSignalReloadCallback func()

/*
ConnectReload connects the callback to the 'reload' signal for the Document.

The returned value represents the connection, and may be passed to DisconnectReload to remove it.
*/
func (recv *Document) ConnectReload(callback DocumentSignalReloadCallback) int {
	signalDocumentReloadLock.Lock()
	defer signalDocumentReloadLock.Unlock()

	signalDocumentReloadId++
	instance := C.gpointer(recv.native)
	handlerID := C.Document_signal_connect_reload(instance, C.gpointer(uintptr(signalDocumentReloadId)))

	detail := signalDocumentReloadDetail{callback, handlerID}
	signalDocumentReloadMap[signalDocumentReloadId] = detail

	return signalDocumentReloadId
}

/*
DisconnectReload disconnects a callback from the 'reload' signal for the Document.

The connectionID should be a value returned from a call to ConnectReload.
*/
func (recv *Document) DisconnectReload(connectionID int) {
	signalDocumentReloadLock.Lock()
	defer signalDocumentReloadLock.Unlock()

	detail, exists := signalDocumentReloadMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDocumentReloadMap, connectionID)
}

//export document_reloadHandler
func document_reloadHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalDocumentReloadMap[index].callback
	callback()
}

// Gets a %gpointer that points to an instance of the DOM.  It is
// up to the caller to check atk_document_get_type to determine
// how to cast this pointer.
/*

C function : atk_document_get_document
*/
func (recv *Document) GetDocument() uintptr {
	retC := C.atk_document_get_document((*C.AtkDocument)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Gets a string indicating the document type.
/*

C function : atk_document_get_document_type
*/
func (recv *Document) GetDocumentType() string {
	retC := C.atk_document_get_document_type((*C.AtkDocument)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets a UTF-8 string indicating the POSIX-style LC_MESSAGES locale
// of the content of this document instance.  Individual
// text substrings or images within this document may have
// a different locale, see atk_text_get_attributes and
// atk_image_get_image_locale.
/*

C function : atk_document_get_locale
*/
func (recv *Document) GetLocale() string {
	retC := C.atk_document_get_locale((*C.AtkDocument)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// EditableText is a wrapper around the C record AtkEditableText.
type EditableText struct {
	native *C.AtkEditableText
}

func EditableTextNewFromC(u unsafe.Pointer) *EditableText {
	c := (*C.AtkEditableText)(u)
	if c == nil {
		return nil
	}

	g := &EditableText{native: c}

	return g
}

func (recv *EditableText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Copy text from @start_pos up to, but not including @end_pos
// to the clipboard.
/*

C function : atk_editable_text_copy_text
*/
func (recv *EditableText) CopyText(startPos int32, endPos int32) {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	C.atk_editable_text_copy_text((*C.AtkEditableText)(recv.native), c_start_pos, c_end_pos)

	return
}

// Copy text from @start_pos up to, but not including @end_pos
// to the clipboard and then delete from the widget.
/*

C function : atk_editable_text_cut_text
*/
func (recv *EditableText) CutText(startPos int32, endPos int32) {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	C.atk_editable_text_cut_text((*C.AtkEditableText)(recv.native), c_start_pos, c_end_pos)

	return
}

// Delete text @start_pos up to, but not including @end_pos.
/*

C function : atk_editable_text_delete_text
*/
func (recv *EditableText) DeleteText(startPos int32, endPos int32) {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	C.atk_editable_text_delete_text((*C.AtkEditableText)(recv.native), c_start_pos, c_end_pos)

	return
}

// Insert text at a given position.
/*

C function : atk_editable_text_insert_text
*/
func (recv *EditableText) InsertText(string string, length int32, position int32) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gint)(length)

	c_position := (C.gint)(position)

	C.atk_editable_text_insert_text((*C.AtkEditableText)(recv.native), c_string, c_length, &c_position)

	return
}

// Paste text from clipboard to specified @position.
/*

C function : atk_editable_text_paste_text
*/
func (recv *EditableText) PasteText(position int32) {
	c_position := (C.gint)(position)

	C.atk_editable_text_paste_text((*C.AtkEditableText)(recv.native), c_position)

	return
}

// Blacklisted : atk_editable_text_set_run_attributes

// Set text contents of @text.
/*

C function : atk_editable_text_set_text_contents
*/
func (recv *EditableText) SetTextContents(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.atk_editable_text_set_text_contents((*C.AtkEditableText)(recv.native), c_string)

	return
}

// HyperlinkImpl is a wrapper around the C record AtkHyperlinkImpl.
type HyperlinkImpl struct {
	native *C.AtkHyperlinkImpl
}

func HyperlinkImplNewFromC(u unsafe.Pointer) *HyperlinkImpl {
	c := (*C.AtkHyperlinkImpl)(u)
	if c == nil {
		return nil
	}

	g := &HyperlinkImpl{native: c}

	return g
}

func (recv *HyperlinkImpl) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Hypertext is a wrapper around the C record AtkHypertext.
type Hypertext struct {
	native *C.AtkHypertext
}

func HypertextNewFromC(u unsafe.Pointer) *Hypertext {
	c := (*C.AtkHypertext)(u)
	if c == nil {
		return nil
	}

	g := &Hypertext{native: c}

	return g
}

func (recv *Hypertext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported signal 'link-selected' for Hypertext : unsupported parameter arg1 : type gint :

// Gets the link in this hypertext document at index
// @link_index
/*

C function : atk_hypertext_get_link
*/
func (recv *Hypertext) GetLink(linkIndex int32) *Hyperlink {
	c_link_index := (C.gint)(linkIndex)

	retC := C.atk_hypertext_get_link((*C.AtkHypertext)(recv.native), c_link_index)
	retGo := HyperlinkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the index into the array of hyperlinks that is associated with
// the character specified by @char_index.
/*

C function : atk_hypertext_get_link_index
*/
func (recv *Hypertext) GetLinkIndex(charIndex int32) int32 {
	c_char_index := (C.gint)(charIndex)

	retC := C.atk_hypertext_get_link_index((*C.AtkHypertext)(recv.native), c_char_index)
	retGo := (int32)(retC)

	return retGo
}

// Gets the number of links within this hypertext document.
/*

C function : atk_hypertext_get_n_links
*/
func (recv *Hypertext) GetNLinks() int32 {
	retC := C.atk_hypertext_get_n_links((*C.AtkHypertext)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Image is a wrapper around the C record AtkImage.
type Image struct {
	native *C.AtkImage
}

func ImageNewFromC(u unsafe.Pointer) *Image {
	c := (*C.AtkImage)(u)
	if c == nil {
		return nil
	}

	g := &Image{native: c}

	return g
}

func (recv *Image) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Get a textual description of this image.
/*

C function : atk_image_get_image_description
*/
func (recv *Image) GetImageDescription() string {
	retC := C.atk_image_get_image_description((*C.AtkImage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the position of the image in the form of a point specifying the
// images top-left corner.
/*

C function : atk_image_get_image_position
*/
func (recv *Image) GetImagePosition(coordType CoordType) (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	c_coord_type := (C.AtkCoordType)(coordType)

	C.atk_image_get_image_position((*C.AtkImage)(recv.native), &c_x, &c_y, c_coord_type)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// Get the width and height in pixels for the specified image.
// The values of @width and @height are returned as -1 if the
// values cannot be obtained (for instance, if the object is not onscreen).
/*

C function : atk_image_get_image_size
*/
func (recv *Image) GetImageSize() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.atk_image_get_image_size((*C.AtkImage)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// Sets the textual description for this image.
/*

C function : atk_image_set_image_description
*/
func (recv *Image) SetImageDescription(description string) bool {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	retC := C.atk_image_set_image_description((*C.AtkImage)(recv.native), c_description)
	retGo := retC == C.TRUE

	return retGo
}

// ImplementorIface is a wrapper around the C record AtkImplementorIface.
type ImplementorIface struct {
	native *C.AtkImplementorIface
}

func ImplementorIfaceNewFromC(u unsafe.Pointer) *ImplementorIface {
	c := (*C.AtkImplementorIface)(u)
	if c == nil {
		return nil
	}

	g := &ImplementorIface{native: c}

	return g
}

func (recv *ImplementorIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Selection is a wrapper around the C record AtkSelection.
type Selection struct {
	native *C.AtkSelection
}

func SelectionNewFromC(u unsafe.Pointer) *Selection {
	c := (*C.AtkSelection)(u)
	if c == nil {
		return nil
	}

	g := &Selection{native: c}

	return g
}

func (recv *Selection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalSelectionSelectionChangedDetail struct {
	callback  SelectionSignalSelectionChangedCallback
	handlerID C.gulong
}

var signalSelectionSelectionChangedId int
var signalSelectionSelectionChangedMap = make(map[int]signalSelectionSelectionChangedDetail)
var signalSelectionSelectionChangedLock sync.Mutex

// SelectionSignalSelectionChangedCallback is a callback function for a 'selection-changed' signal emitted from a Selection.
type SelectionSignalSelectionChangedCallback func()

/*
ConnectSelectionChanged connects the callback to the 'selection-changed' signal for the Selection.

The returned value represents the connection, and may be passed to DisconnectSelectionChanged to remove it.
*/
func (recv *Selection) ConnectSelectionChanged(callback SelectionSignalSelectionChangedCallback) int {
	signalSelectionSelectionChangedLock.Lock()
	defer signalSelectionSelectionChangedLock.Unlock()

	signalSelectionSelectionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Selection_signal_connect_selection_changed(instance, C.gpointer(uintptr(signalSelectionSelectionChangedId)))

	detail := signalSelectionSelectionChangedDetail{callback, handlerID}
	signalSelectionSelectionChangedMap[signalSelectionSelectionChangedId] = detail

	return signalSelectionSelectionChangedId
}

/*
DisconnectSelectionChanged disconnects a callback from the 'selection-changed' signal for the Selection.

The connectionID should be a value returned from a call to ConnectSelectionChanged.
*/
func (recv *Selection) DisconnectSelectionChanged(connectionID int) {
	signalSelectionSelectionChangedLock.Lock()
	defer signalSelectionSelectionChangedLock.Unlock()

	detail, exists := signalSelectionSelectionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSelectionSelectionChangedMap, connectionID)
}

//export selection_selectionChangedHandler
func selection_selectionChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalSelectionSelectionChangedMap[index].callback
	callback()
}

// Adds the specified accessible child of the object to the
// object's selection.
/*

C function : atk_selection_add_selection
*/
func (recv *Selection) AddSelection(i int32) bool {
	c_i := (C.gint)(i)

	retC := C.atk_selection_add_selection((*C.AtkSelection)(recv.native), c_i)
	retGo := retC == C.TRUE

	return retGo
}

// Clears the selection in the object so that no children in the object
// are selected.
/*

C function : atk_selection_clear_selection
*/
func (recv *Selection) ClearSelection() bool {
	retC := C.atk_selection_clear_selection((*C.AtkSelection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the number of accessible children currently selected.
// Note: callers should not rely on %NULL or on a zero value for
// indication of whether AtkSelectionIface is implemented, they should
// use type checking/interface checking macros or the
// atk_get_accessible_value() convenience method.
/*

C function : atk_selection_get_selection_count
*/
func (recv *Selection) GetSelectionCount() int32 {
	retC := C.atk_selection_get_selection_count((*C.AtkSelection)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Determines if the current child of this object is selected
// Note: callers should not rely on %NULL or on a zero value for
// indication of whether AtkSelectionIface is implemented, they should
// use type checking/interface checking macros or the
// atk_get_accessible_value() convenience method.
/*

C function : atk_selection_is_child_selected
*/
func (recv *Selection) IsChildSelected(i int32) bool {
	c_i := (C.gint)(i)

	retC := C.atk_selection_is_child_selected((*C.AtkSelection)(recv.native), c_i)
	retGo := retC == C.TRUE

	return retGo
}

// Gets a reference to the accessible object representing the specified
// selected child of the object.
// Note: callers should not rely on %NULL or on a zero value for
// indication of whether AtkSelectionIface is implemented, they should
// use type checking/interface checking macros or the
// atk_get_accessible_value() convenience method.
/*

C function : atk_selection_ref_selection
*/
func (recv *Selection) RefSelection(i int32) *Object {
	c_i := (C.gint)(i)

	retC := C.atk_selection_ref_selection((*C.AtkSelection)(recv.native), c_i)
	var retGo (*Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Removes the specified child of the object from the object's selection.
/*

C function : atk_selection_remove_selection
*/
func (recv *Selection) RemoveSelection(i int32) bool {
	c_i := (C.gint)(i)

	retC := C.atk_selection_remove_selection((*C.AtkSelection)(recv.native), c_i)
	retGo := retC == C.TRUE

	return retGo
}

// Causes every child of the object to be selected if the object
// supports multiple selections.
/*

C function : atk_selection_select_all_selection
*/
func (recv *Selection) SelectAllSelection() bool {
	retC := C.atk_selection_select_all_selection((*C.AtkSelection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// StreamableContent is a wrapper around the C record AtkStreamableContent.
type StreamableContent struct {
	native *C.AtkStreamableContent
}

func StreamableContentNewFromC(u unsafe.Pointer) *StreamableContent {
	c := (*C.AtkStreamableContent)(u)
	if c == nil {
		return nil
	}

	g := &StreamableContent{native: c}

	return g
}

func (recv *StreamableContent) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gets the character string of the specified mime type. The first mime
// type is at position 0, the second at position 1, and so on.
/*

C function : atk_streamable_content_get_mime_type
*/
func (recv *StreamableContent) GetMimeType(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_streamable_content_get_mime_type((*C.AtkStreamableContent)(recv.native), c_i)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the number of mime types supported by this object.
/*

C function : atk_streamable_content_get_n_mime_types
*/
func (recv *StreamableContent) GetNMimeTypes() int32 {
	retC := C.atk_streamable_content_get_n_mime_types((*C.AtkStreamableContent)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : atk_streamable_content_get_stream : return type : Blacklisted record : GIOChannel

// Table is a wrapper around the C record AtkTable.
type Table struct {
	native *C.AtkTable
}

func TableNewFromC(u unsafe.Pointer) *Table {
	c := (*C.AtkTable)(u)
	if c == nil {
		return nil
	}

	g := &Table{native: c}

	return g
}

func (recv *Table) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported signal 'column-deleted' for Table : unsupported parameter arg1 : type gint :

// Unsupported signal 'column-inserted' for Table : unsupported parameter arg1 : type gint :

type signalTableColumnReorderedDetail struct {
	callback  TableSignalColumnReorderedCallback
	handlerID C.gulong
}

var signalTableColumnReorderedId int
var signalTableColumnReorderedMap = make(map[int]signalTableColumnReorderedDetail)
var signalTableColumnReorderedLock sync.Mutex

// TableSignalColumnReorderedCallback is a callback function for a 'column-reordered' signal emitted from a Table.
type TableSignalColumnReorderedCallback func()

/*
ConnectColumnReordered connects the callback to the 'column-reordered' signal for the Table.

The returned value represents the connection, and may be passed to DisconnectColumnReordered to remove it.
*/
func (recv *Table) ConnectColumnReordered(callback TableSignalColumnReorderedCallback) int {
	signalTableColumnReorderedLock.Lock()
	defer signalTableColumnReorderedLock.Unlock()

	signalTableColumnReorderedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Table_signal_connect_column_reordered(instance, C.gpointer(uintptr(signalTableColumnReorderedId)))

	detail := signalTableColumnReorderedDetail{callback, handlerID}
	signalTableColumnReorderedMap[signalTableColumnReorderedId] = detail

	return signalTableColumnReorderedId
}

/*
DisconnectColumnReordered disconnects a callback from the 'column-reordered' signal for the Table.

The connectionID should be a value returned from a call to ConnectColumnReordered.
*/
func (recv *Table) DisconnectColumnReordered(connectionID int) {
	signalTableColumnReorderedLock.Lock()
	defer signalTableColumnReorderedLock.Unlock()

	detail, exists := signalTableColumnReorderedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTableColumnReorderedMap, connectionID)
}

//export table_columnReorderedHandler
func table_columnReorderedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalTableColumnReorderedMap[index].callback
	callback()
}

type signalTableModelChangedDetail struct {
	callback  TableSignalModelChangedCallback
	handlerID C.gulong
}

var signalTableModelChangedId int
var signalTableModelChangedMap = make(map[int]signalTableModelChangedDetail)
var signalTableModelChangedLock sync.Mutex

// TableSignalModelChangedCallback is a callback function for a 'model-changed' signal emitted from a Table.
type TableSignalModelChangedCallback func()

/*
ConnectModelChanged connects the callback to the 'model-changed' signal for the Table.

The returned value represents the connection, and may be passed to DisconnectModelChanged to remove it.
*/
func (recv *Table) ConnectModelChanged(callback TableSignalModelChangedCallback) int {
	signalTableModelChangedLock.Lock()
	defer signalTableModelChangedLock.Unlock()

	signalTableModelChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Table_signal_connect_model_changed(instance, C.gpointer(uintptr(signalTableModelChangedId)))

	detail := signalTableModelChangedDetail{callback, handlerID}
	signalTableModelChangedMap[signalTableModelChangedId] = detail

	return signalTableModelChangedId
}

/*
DisconnectModelChanged disconnects a callback from the 'model-changed' signal for the Table.

The connectionID should be a value returned from a call to ConnectModelChanged.
*/
func (recv *Table) DisconnectModelChanged(connectionID int) {
	signalTableModelChangedLock.Lock()
	defer signalTableModelChangedLock.Unlock()

	detail, exists := signalTableModelChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTableModelChangedMap, connectionID)
}

//export table_modelChangedHandler
func table_modelChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalTableModelChangedMap[index].callback
	callback()
}

// Unsupported signal 'row-deleted' for Table : unsupported parameter arg1 : type gint :

// Unsupported signal 'row-inserted' for Table : unsupported parameter arg1 : type gint :

type signalTableRowReorderedDetail struct {
	callback  TableSignalRowReorderedCallback
	handlerID C.gulong
}

var signalTableRowReorderedId int
var signalTableRowReorderedMap = make(map[int]signalTableRowReorderedDetail)
var signalTableRowReorderedLock sync.Mutex

// TableSignalRowReorderedCallback is a callback function for a 'row-reordered' signal emitted from a Table.
type TableSignalRowReorderedCallback func()

/*
ConnectRowReordered connects the callback to the 'row-reordered' signal for the Table.

The returned value represents the connection, and may be passed to DisconnectRowReordered to remove it.
*/
func (recv *Table) ConnectRowReordered(callback TableSignalRowReorderedCallback) int {
	signalTableRowReorderedLock.Lock()
	defer signalTableRowReorderedLock.Unlock()

	signalTableRowReorderedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Table_signal_connect_row_reordered(instance, C.gpointer(uintptr(signalTableRowReorderedId)))

	detail := signalTableRowReorderedDetail{callback, handlerID}
	signalTableRowReorderedMap[signalTableRowReorderedId] = detail

	return signalTableRowReorderedId
}

/*
DisconnectRowReordered disconnects a callback from the 'row-reordered' signal for the Table.

The connectionID should be a value returned from a call to ConnectRowReordered.
*/
func (recv *Table) DisconnectRowReordered(connectionID int) {
	signalTableRowReorderedLock.Lock()
	defer signalTableRowReorderedLock.Unlock()

	detail, exists := signalTableRowReorderedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTableRowReorderedMap, connectionID)
}

//export table_rowReorderedHandler
func table_rowReorderedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalTableRowReorderedMap[index].callback
	callback()
}

// Adds the specified @column to the selection.
/*

C function : atk_table_add_column_selection
*/
func (recv *Table) AddColumnSelection(column int32) bool {
	c_column := (C.gint)(column)

	retC := C.atk_table_add_column_selection((*C.AtkTable)(recv.native), c_column)
	retGo := retC == C.TRUE

	return retGo
}

// Adds the specified @row to the selection.
/*

C function : atk_table_add_row_selection
*/
func (recv *Table) AddRowSelection(row int32) bool {
	c_row := (C.gint)(row)

	retC := C.atk_table_add_row_selection((*C.AtkTable)(recv.native), c_row)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the caption for the @table.
/*

C function : atk_table_get_caption
*/
func (recv *Table) GetCaption() *Object {
	retC := C.atk_table_get_caption((*C.AtkTable)(recv.native))
	var retGo (*Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets a #gint representing the column at the specified @index_.
/*

C function : atk_table_get_column_at_index
*/
func (recv *Table) GetColumnAtIndex(index int32) int32 {
	c_index_ := (C.gint)(index)

	retC := C.atk_table_get_column_at_index((*C.AtkTable)(recv.native), c_index_)
	retGo := (int32)(retC)

	return retGo
}

// Gets the description text of the specified @column in the table
/*

C function : atk_table_get_column_description
*/
func (recv *Table) GetColumnDescription(column int32) string {
	c_column := (C.gint)(column)

	retC := C.atk_table_get_column_description((*C.AtkTable)(recv.native), c_column)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the number of columns occupied by the accessible object
// at the specified @row and @column in the @table.
/*

C function : atk_table_get_column_extent_at
*/
func (recv *Table) GetColumnExtentAt(row int32, column int32) int32 {
	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	retC := C.atk_table_get_column_extent_at((*C.AtkTable)(recv.native), c_row, c_column)
	retGo := (int32)(retC)

	return retGo
}

// Gets the column header of a specified column in an accessible table.
/*

C function : atk_table_get_column_header
*/
func (recv *Table) GetColumnHeader(column int32) *Object {
	c_column := (C.gint)(column)

	retC := C.atk_table_get_column_header((*C.AtkTable)(recv.native), c_column)
	var retGo (*Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets a #gint representing the index at the specified @row and
// @column.
/*

C function : atk_table_get_index_at
*/
func (recv *Table) GetIndexAt(row int32, column int32) int32 {
	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	retC := C.atk_table_get_index_at((*C.AtkTable)(recv.native), c_row, c_column)
	retGo := (int32)(retC)

	return retGo
}

// Gets the number of columns in the table.
/*

C function : atk_table_get_n_columns
*/
func (recv *Table) GetNColumns() int32 {
	retC := C.atk_table_get_n_columns((*C.AtkTable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the number of rows in the table.
/*

C function : atk_table_get_n_rows
*/
func (recv *Table) GetNRows() int32 {
	retC := C.atk_table_get_n_rows((*C.AtkTable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets a #gint representing the row at the specified @index_.
/*

C function : atk_table_get_row_at_index
*/
func (recv *Table) GetRowAtIndex(index int32) int32 {
	c_index_ := (C.gint)(index)

	retC := C.atk_table_get_row_at_index((*C.AtkTable)(recv.native), c_index_)
	retGo := (int32)(retC)

	return retGo
}

// Gets the description text of the specified row in the table
/*

C function : atk_table_get_row_description
*/
func (recv *Table) GetRowDescription(row int32) string {
	c_row := (C.gint)(row)

	retC := C.atk_table_get_row_description((*C.AtkTable)(recv.native), c_row)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the number of rows occupied by the accessible object
// at a specified @row and @column in the @table.
/*

C function : atk_table_get_row_extent_at
*/
func (recv *Table) GetRowExtentAt(row int32, column int32) int32 {
	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	retC := C.atk_table_get_row_extent_at((*C.AtkTable)(recv.native), c_row, c_column)
	retGo := (int32)(retC)

	return retGo
}

// Gets the row header of a specified row in an accessible table.
/*

C function : atk_table_get_row_header
*/
func (recv *Table) GetRowHeader(row int32) *Object {
	c_row := (C.gint)(row)

	retC := C.atk_table_get_row_header((*C.AtkTable)(recv.native), c_row)
	var retGo (*Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : atk_table_get_selected_columns : unsupported parameter selected : gint** with indirection level of 2

// Unsupported : atk_table_get_selected_rows : unsupported parameter selected : gint** with indirection level of 2

// Gets the summary description of the table.
/*

C function : atk_table_get_summary
*/
func (recv *Table) GetSummary() *Object {
	retC := C.atk_table_get_summary((*C.AtkTable)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a boolean value indicating whether the specified @column
// is selected
/*

C function : atk_table_is_column_selected
*/
func (recv *Table) IsColumnSelected(column int32) bool {
	c_column := (C.gint)(column)

	retC := C.atk_table_is_column_selected((*C.AtkTable)(recv.native), c_column)
	retGo := retC == C.TRUE

	return retGo
}

// Gets a boolean value indicating whether the specified @row
// is selected
/*

C function : atk_table_is_row_selected
*/
func (recv *Table) IsRowSelected(row int32) bool {
	c_row := (C.gint)(row)

	retC := C.atk_table_is_row_selected((*C.AtkTable)(recv.native), c_row)
	retGo := retC == C.TRUE

	return retGo
}

// Gets a boolean value indicating whether the accessible object
// at the specified @row and @column is selected
/*

C function : atk_table_is_selected
*/
func (recv *Table) IsSelected(row int32, column int32) bool {
	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	retC := C.atk_table_is_selected((*C.AtkTable)(recv.native), c_row, c_column)
	retGo := retC == C.TRUE

	return retGo
}

// Get a reference to the table cell at @row, @column. This cell
// should implement the interface #AtkTableCell
/*

C function : atk_table_ref_at
*/
func (recv *Table) RefAt(row int32, column int32) *Object {
	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	retC := C.atk_table_ref_at((*C.AtkTable)(recv.native), c_row, c_column)
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds the specified @column to the selection.
/*

C function : atk_table_remove_column_selection
*/
func (recv *Table) RemoveColumnSelection(column int32) bool {
	c_column := (C.gint)(column)

	retC := C.atk_table_remove_column_selection((*C.AtkTable)(recv.native), c_column)
	retGo := retC == C.TRUE

	return retGo
}

// Removes the specified @row from the selection.
/*

C function : atk_table_remove_row_selection
*/
func (recv *Table) RemoveRowSelection(row int32) bool {
	c_row := (C.gint)(row)

	retC := C.atk_table_remove_row_selection((*C.AtkTable)(recv.native), c_row)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the caption for the table.
/*

C function : atk_table_set_caption
*/
func (recv *Table) SetCaption(caption *Object) {
	c_caption := (*C.AtkObject)(C.NULL)
	if caption != nil {
		c_caption = (*C.AtkObject)(caption.ToC())
	}

	C.atk_table_set_caption((*C.AtkTable)(recv.native), c_caption)

	return
}

// Sets the description text for the specified @column of the @table.
/*

C function : atk_table_set_column_description
*/
func (recv *Table) SetColumnDescription(column int32, description string) {
	c_column := (C.gint)(column)

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.atk_table_set_column_description((*C.AtkTable)(recv.native), c_column, c_description)

	return
}

// Sets the specified column header to @header.
/*

C function : atk_table_set_column_header
*/
func (recv *Table) SetColumnHeader(column int32, header *Object) {
	c_column := (C.gint)(column)

	c_header := (*C.AtkObject)(C.NULL)
	if header != nil {
		c_header = (*C.AtkObject)(header.ToC())
	}

	C.atk_table_set_column_header((*C.AtkTable)(recv.native), c_column, c_header)

	return
}

// Sets the description text for the specified @row of @table.
/*

C function : atk_table_set_row_description
*/
func (recv *Table) SetRowDescription(row int32, description string) {
	c_row := (C.gint)(row)

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.atk_table_set_row_description((*C.AtkTable)(recv.native), c_row, c_description)

	return
}

// Sets the specified row header to @header.
/*

C function : atk_table_set_row_header
*/
func (recv *Table) SetRowHeader(row int32, header *Object) {
	c_row := (C.gint)(row)

	c_header := (*C.AtkObject)(C.NULL)
	if header != nil {
		c_header = (*C.AtkObject)(header.ToC())
	}

	C.atk_table_set_row_header((*C.AtkTable)(recv.native), c_row, c_header)

	return
}

// Sets the summary description of the table.
/*

C function : atk_table_set_summary
*/
func (recv *Table) SetSummary(accessible *Object) {
	c_accessible := (*C.AtkObject)(C.NULL)
	if accessible != nil {
		c_accessible = (*C.AtkObject)(accessible.ToC())
	}

	C.atk_table_set_summary((*C.AtkTable)(recv.native), c_accessible)

	return
}

// TableCell is a wrapper around the C record AtkTableCell.
type TableCell struct {
	native *C.AtkTableCell
}

func TableCellNewFromC(u unsafe.Pointer) *TableCell {
	c := (*C.AtkTableCell)(u)
	if c == nil {
		return nil
	}

	g := &TableCell{native: c}

	return g
}

func (recv *TableCell) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Text is a wrapper around the C record AtkText.
type Text struct {
	native *C.AtkText
}

func TextNewFromC(u unsafe.Pointer) *Text {
	c := (*C.AtkText)(u)
	if c == nil {
		return nil
	}

	g := &Text{native: c}

	return g
}

func (recv *Text) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalTextTextAttributesChangedDetail struct {
	callback  TextSignalTextAttributesChangedCallback
	handlerID C.gulong
}

var signalTextTextAttributesChangedId int
var signalTextTextAttributesChangedMap = make(map[int]signalTextTextAttributesChangedDetail)
var signalTextTextAttributesChangedLock sync.Mutex

// TextSignalTextAttributesChangedCallback is a callback function for a 'text-attributes-changed' signal emitted from a Text.
type TextSignalTextAttributesChangedCallback func()

/*
ConnectTextAttributesChanged connects the callback to the 'text-attributes-changed' signal for the Text.

The returned value represents the connection, and may be passed to DisconnectTextAttributesChanged to remove it.
*/
func (recv *Text) ConnectTextAttributesChanged(callback TextSignalTextAttributesChangedCallback) int {
	signalTextTextAttributesChangedLock.Lock()
	defer signalTextTextAttributesChangedLock.Unlock()

	signalTextTextAttributesChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Text_signal_connect_text_attributes_changed(instance, C.gpointer(uintptr(signalTextTextAttributesChangedId)))

	detail := signalTextTextAttributesChangedDetail{callback, handlerID}
	signalTextTextAttributesChangedMap[signalTextTextAttributesChangedId] = detail

	return signalTextTextAttributesChangedId
}

/*
DisconnectTextAttributesChanged disconnects a callback from the 'text-attributes-changed' signal for the Text.

The connectionID should be a value returned from a call to ConnectTextAttributesChanged.
*/
func (recv *Text) DisconnectTextAttributesChanged(connectionID int) {
	signalTextTextAttributesChangedLock.Lock()
	defer signalTextTextAttributesChangedLock.Unlock()

	detail, exists := signalTextTextAttributesChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextTextAttributesChangedMap, connectionID)
}

//export text_textAttributesChangedHandler
func text_textAttributesChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalTextTextAttributesChangedMap[index].callback
	callback()
}

// Unsupported signal 'text-caret-moved' for Text : unsupported parameter arg1 : type gint :

// Unsupported signal 'text-changed' for Text : unsupported parameter arg1 : type gint :

// Unsupported signal 'text-insert' for Text : unsupported parameter arg1 : type gint :

// Unsupported signal 'text-remove' for Text : unsupported parameter arg1 : type gint :

type signalTextTextSelectionChangedDetail struct {
	callback  TextSignalTextSelectionChangedCallback
	handlerID C.gulong
}

var signalTextTextSelectionChangedId int
var signalTextTextSelectionChangedMap = make(map[int]signalTextTextSelectionChangedDetail)
var signalTextTextSelectionChangedLock sync.Mutex

// TextSignalTextSelectionChangedCallback is a callback function for a 'text-selection-changed' signal emitted from a Text.
type TextSignalTextSelectionChangedCallback func()

/*
ConnectTextSelectionChanged connects the callback to the 'text-selection-changed' signal for the Text.

The returned value represents the connection, and may be passed to DisconnectTextSelectionChanged to remove it.
*/
func (recv *Text) ConnectTextSelectionChanged(callback TextSignalTextSelectionChangedCallback) int {
	signalTextTextSelectionChangedLock.Lock()
	defer signalTextTextSelectionChangedLock.Unlock()

	signalTextTextSelectionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Text_signal_connect_text_selection_changed(instance, C.gpointer(uintptr(signalTextTextSelectionChangedId)))

	detail := signalTextTextSelectionChangedDetail{callback, handlerID}
	signalTextTextSelectionChangedMap[signalTextTextSelectionChangedId] = detail

	return signalTextTextSelectionChangedId
}

/*
DisconnectTextSelectionChanged disconnects a callback from the 'text-selection-changed' signal for the Text.

The connectionID should be a value returned from a call to ConnectTextSelectionChanged.
*/
func (recv *Text) DisconnectTextSelectionChanged(connectionID int) {
	signalTextTextSelectionChangedLock.Lock()
	defer signalTextTextSelectionChangedLock.Unlock()

	detail, exists := signalTextTextSelectionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextTextSelectionChangedMap, connectionID)
}

//export text_textSelectionChangedHandler
func text_textSelectionChangedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalTextTextSelectionChangedMap[index].callback
	callback()
}

// Adds a selection bounded by the specified offsets.
/*

C function : atk_text_add_selection
*/
func (recv *Text) AddSelection(startOffset int32, endOffset int32) bool {
	c_start_offset := (C.gint)(startOffset)

	c_end_offset := (C.gint)(endOffset)

	retC := C.atk_text_add_selection((*C.AtkText)(recv.native), c_start_offset, c_end_offset)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the offset position of the caret (cursor).
/*

C function : atk_text_get_caret_offset
*/
func (recv *Text) GetCaretOffset() int32 {
	retC := C.atk_text_get_caret_offset((*C.AtkText)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the specified text.
/*

C function : atk_text_get_character_at_offset
*/
func (recv *Text) GetCharacterAtOffset(offset int32) rune {
	c_offset := (C.gint)(offset)

	retC := C.atk_text_get_character_at_offset((*C.AtkText)(recv.native), c_offset)
	retGo := (rune)(retC)

	return retGo
}

// Gets the character count.
/*

C function : atk_text_get_character_count
*/
func (recv *Text) GetCharacterCount() int32 {
	retC := C.atk_text_get_character_count((*C.AtkText)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Get the bounding box containing the glyph representing the character at
// a particular text offset.
/*

C function : atk_text_get_character_extents
*/
func (recv *Text) GetCharacterExtents(offset int32, coords CoordType) (int32, int32, int32, int32) {
	c_offset := (C.gint)(offset)

	var c_x C.gint

	var c_y C.gint

	var c_width C.gint

	var c_height C.gint

	c_coords := (C.AtkCoordType)(coords)

	C.atk_text_get_character_extents((*C.AtkText)(recv.native), c_offset, &c_x, &c_y, &c_width, &c_height, c_coords)

	x := (int32)(c_x)

	y := (int32)(c_y)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return x, y, width, height
}

// Blacklisted : atk_text_get_default_attributes

// Gets the number of selected regions.
/*

C function : atk_text_get_n_selections
*/
func (recv *Text) GetNSelections() int32 {
	retC := C.atk_text_get_n_selections((*C.AtkText)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the offset of the character located at coordinates @x and @y. @x and @y
// are interpreted as being relative to the screen or this widget's window
// depending on @coords.
/*

C function : atk_text_get_offset_at_point
*/
func (recv *Text) GetOffsetAtPoint(x int32, y int32, coords CoordType) int32 {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coords := (C.AtkCoordType)(coords)

	retC := C.atk_text_get_offset_at_point((*C.AtkText)(recv.native), c_x, c_y, c_coords)
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : atk_text_get_run_attributes

// Gets the text from the specified selection.
/*

C function : atk_text_get_selection
*/
func (recv *Text) GetSelection(selectionNum int32) (string, int32, int32) {
	c_selection_num := (C.gint)(selectionNum)

	var c_start_offset C.gint

	var c_end_offset C.gint

	retC := C.atk_text_get_selection((*C.AtkText)(recv.native), c_selection_num, &c_start_offset, &c_end_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	startOffset := (int32)(c_start_offset)

	endOffset := (int32)(c_end_offset)

	return retGo, startOffset, endOffset
}

// Gets the specified text.
/*

C function : atk_text_get_text
*/
func (recv *Text) GetText(startOffset int32, endOffset int32) string {
	c_start_offset := (C.gint)(startOffset)

	c_end_offset := (C.gint)(endOffset)

	retC := C.atk_text_get_text((*C.AtkText)(recv.native), c_start_offset, c_end_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the specified text.
/*

C function : atk_text_get_text_after_offset
*/
func (recv *Text) GetTextAfterOffset(offset int32, boundaryType TextBoundary) (string, int32, int32) {
	c_offset := (C.gint)(offset)

	c_boundary_type := (C.AtkTextBoundary)(boundaryType)

	var c_start_offset C.gint

	var c_end_offset C.gint

	retC := C.atk_text_get_text_after_offset((*C.AtkText)(recv.native), c_offset, c_boundary_type, &c_start_offset, &c_end_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	startOffset := (int32)(c_start_offset)

	endOffset := (int32)(c_end_offset)

	return retGo, startOffset, endOffset
}

// Gets the specified text.
//
// If the boundary_type if ATK_TEXT_BOUNDARY_CHAR the character at the
// offset is returned.
//
// If the boundary_type is ATK_TEXT_BOUNDARY_WORD_START the returned string
// is from the word start at or before the offset to the word start after
// the offset.
//
// The returned string will contain the word at the offset if the offset
// is inside a word and will contain the word before the offset if the
// offset is not inside a word.
//
// If the boundary type is ATK_TEXT_BOUNDARY_SENTENCE_START the returned
// string is from the sentence start at or before the offset to the sentence
// start after the offset.
//
// The returned string will contain the sentence at the offset if the offset
// is inside a sentence and will contain the sentence before the offset
// if the offset is not inside a sentence.
//
// If the boundary type is ATK_TEXT_BOUNDARY_LINE_START the returned
// string is from the line start at or before the offset to the line
// start after the offset.
/*

C function : atk_text_get_text_at_offset
*/
func (recv *Text) GetTextAtOffset(offset int32, boundaryType TextBoundary) (string, int32, int32) {
	c_offset := (C.gint)(offset)

	c_boundary_type := (C.AtkTextBoundary)(boundaryType)

	var c_start_offset C.gint

	var c_end_offset C.gint

	retC := C.atk_text_get_text_at_offset((*C.AtkText)(recv.native), c_offset, c_boundary_type, &c_start_offset, &c_end_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	startOffset := (int32)(c_start_offset)

	endOffset := (int32)(c_end_offset)

	return retGo, startOffset, endOffset
}

// Gets the specified text.
/*

C function : atk_text_get_text_before_offset
*/
func (recv *Text) GetTextBeforeOffset(offset int32, boundaryType TextBoundary) (string, int32, int32) {
	c_offset := (C.gint)(offset)

	c_boundary_type := (C.AtkTextBoundary)(boundaryType)

	var c_start_offset C.gint

	var c_end_offset C.gint

	retC := C.atk_text_get_text_before_offset((*C.AtkText)(recv.native), c_offset, c_boundary_type, &c_start_offset, &c_end_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	startOffset := (int32)(c_start_offset)

	endOffset := (int32)(c_end_offset)

	return retGo, startOffset, endOffset
}

// Removes the specified selection.
/*

C function : atk_text_remove_selection
*/
func (recv *Text) RemoveSelection(selectionNum int32) bool {
	c_selection_num := (C.gint)(selectionNum)

	retC := C.atk_text_remove_selection((*C.AtkText)(recv.native), c_selection_num)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the caret (cursor) position to the specified @offset.
/*

C function : atk_text_set_caret_offset
*/
func (recv *Text) SetCaretOffset(offset int32) bool {
	c_offset := (C.gint)(offset)

	retC := C.atk_text_set_caret_offset((*C.AtkText)(recv.native), c_offset)
	retGo := retC == C.TRUE

	return retGo
}

// Changes the start and end offset of the specified selection.
/*

C function : atk_text_set_selection
*/
func (recv *Text) SetSelection(selectionNum int32, startOffset int32, endOffset int32) bool {
	c_selection_num := (C.gint)(selectionNum)

	c_start_offset := (C.gint)(startOffset)

	c_end_offset := (C.gint)(endOffset)

	retC := C.atk_text_set_selection((*C.AtkText)(recv.native), c_selection_num, c_start_offset, c_end_offset)
	retGo := retC == C.TRUE

	return retGo
}

// Value is a wrapper around the C record AtkValue.
type Value struct {
	native *C.AtkValue
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	c := (*C.AtkValue)(u)
	if c == nil {
		return nil
	}

	g := &Value{native: c}

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gets the value of this object.
/*

C function : atk_value_get_current_value
*/
func (recv *Value) GetCurrentValue() *gobject.Value {
	var c_value C.GValue

	C.atk_value_get_current_value((*C.AtkValue)(recv.native), &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// Gets the maximum value of this object.
/*

C function : atk_value_get_maximum_value
*/
func (recv *Value) GetMaximumValue() *gobject.Value {
	var c_value C.GValue

	C.atk_value_get_maximum_value((*C.AtkValue)(recv.native), &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// Gets the minimum value of this object.
/*

C function : atk_value_get_minimum_value
*/
func (recv *Value) GetMinimumValue() *gobject.Value {
	var c_value C.GValue

	C.atk_value_get_minimum_value((*C.AtkValue)(recv.native), &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// Sets the value of this object.
/*

C function : atk_value_set_current_value
*/
func (recv *Value) SetCurrentValue(value *gobject.Value) bool {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.atk_value_set_current_value((*C.AtkValue)(recv.native), c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Window is a wrapper around the C record AtkWindow.
type Window struct {
	native *C.AtkWindow
}

func WindowNewFromC(u unsafe.Pointer) *Window {
	c := (*C.AtkWindow)(u)
	if c == nil {
		return nil
	}

	g := &Window{native: c}

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
