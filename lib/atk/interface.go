// This is a generated file - DO NOT EDIT

package atk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

	void hypertext_linkSelectedHandler(GObject *, gint, gpointer);

	static gulong Hypertext_signal_connect_link_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "link-selected", G_CALLBACK(hypertext_linkSelectedHandler), data);
	}

*/
/*

	void selection_selectionChangedHandler(GObject *, gpointer);

	static gulong Selection_signal_connect_selection_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-changed", G_CALLBACK(selection_selectionChangedHandler), data);
	}

*/
/*

	void table_columnDeletedHandler(GObject *, gint, gint, gpointer);

	static gulong Table_signal_connect_column_deleted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "column-deleted", G_CALLBACK(table_columnDeletedHandler), data);
	}

*/
/*

	void table_columnInsertedHandler(GObject *, gint, gint, gpointer);

	static gulong Table_signal_connect_column_inserted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "column-inserted", G_CALLBACK(table_columnInsertedHandler), data);
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

	void table_rowDeletedHandler(GObject *, gint, gint, gpointer);

	static gulong Table_signal_connect_row_deleted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-deleted", G_CALLBACK(table_rowDeletedHandler), data);
	}

*/
/*

	void table_rowInsertedHandler(GObject *, gint, gint, gpointer);

	static gulong Table_signal_connect_row_inserted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-inserted", G_CALLBACK(table_rowInsertedHandler), data);
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

	void text_textCaretMovedHandler(GObject *, gint, gpointer);

	static gulong Text_signal_connect_text_caret_moved(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "text-caret-moved", G_CALLBACK(text_textCaretMovedHandler), data);
	}

*/
/*

	void text_textChangedHandler(GObject *, gint, gint, gpointer);

	static gulong Text_signal_connect_text_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "text-changed", G_CALLBACK(text_textChangedHandler), data);
	}

*/
/*

	void text_textInsertHandler(GObject *, gint, gint, gchar*, gpointer);

	static gulong Text_signal_connect_text_insert(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "text-insert", G_CALLBACK(text_textInsertHandler), data);
	}

*/
/*

	void text_textRemoveHandler(GObject *, gint, gint, gchar*, gpointer);

	static gulong Text_signal_connect_text_remove(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "text-remove", G_CALLBACK(text_textRemoveHandler), data);
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

// Equals compares this Action with another Action, and returns true if they represent the same GObject.
func (recv *Action) Equals(other *Action) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : atk_action_do_action

// Blacklisted : atk_action_get_description

// Blacklisted : atk_action_get_keybinding

// Blacklisted : atk_action_get_localized_name

// Blacklisted : atk_action_get_n_actions

// Blacklisted : atk_action_get_name

// Blacklisted : atk_action_set_description

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

// Equals compares this Component with another Component, and returns true if they represent the same GObject.
func (recv *Component) Equals(other *Component) bool {
	return other.ToC() == recv.ToC()
}

type signalComponentBoundsChangedDetail struct {
	callback  ComponentSignalBoundsChangedCallback
	handlerID C.gulong
}

var signalComponentBoundsChangedId int
var signalComponentBoundsChangedMap = make(map[int]signalComponentBoundsChangedDetail)
var signalComponentBoundsChangedLock sync.RWMutex

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
	signalComponentBoundsChangedLock.RLock()
	defer signalComponentBoundsChangedLock.RUnlock()

	arg1 := RectangleNewFromC(unsafe.Pointer(c_arg1))

	index := int(uintptr(data))
	callback := signalComponentBoundsChangedMap[index].callback
	callback(arg1)
}

// Unsupported : atk_component_add_focus_handler : unsupported parameter handler : no type generator for FocusHandler (AtkFocusHandler) for param handler

// Blacklisted : atk_component_contains

// Blacklisted : atk_component_get_extents

// Blacklisted : atk_component_get_layer

// Blacklisted : atk_component_get_mdi_zorder

// Blacklisted : atk_component_get_position

// Blacklisted : atk_component_get_size

// Blacklisted : atk_component_grab_focus

// Blacklisted : atk_component_ref_accessible_at_point

// Blacklisted : atk_component_remove_focus_handler

// Blacklisted : atk_component_set_extents

// Blacklisted : atk_component_set_position

// Blacklisted : atk_component_set_size

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

// Equals compares this Document with another Document, and returns true if they represent the same GObject.
func (recv *Document) Equals(other *Document) bool {
	return other.ToC() == recv.ToC()
}

type signalDocumentLoadCompleteDetail struct {
	callback  DocumentSignalLoadCompleteCallback
	handlerID C.gulong
}

var signalDocumentLoadCompleteId int
var signalDocumentLoadCompleteMap = make(map[int]signalDocumentLoadCompleteDetail)
var signalDocumentLoadCompleteLock sync.RWMutex

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
	signalDocumentLoadCompleteLock.RLock()
	defer signalDocumentLoadCompleteLock.RUnlock()

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
var signalDocumentLoadStoppedLock sync.RWMutex

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
	signalDocumentLoadStoppedLock.RLock()
	defer signalDocumentLoadStoppedLock.RUnlock()

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
var signalDocumentReloadLock sync.RWMutex

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
	signalDocumentReloadLock.RLock()
	defer signalDocumentReloadLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDocumentReloadMap[index].callback
	callback()
}

// Unsupported : atk_document_get_document : no return generator

// Blacklisted : atk_document_get_document_type

// Blacklisted : atk_document_get_locale

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

// Equals compares this EditableText with another EditableText, and returns true if they represent the same GObject.
func (recv *EditableText) Equals(other *EditableText) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : atk_editable_text_copy_text

// Blacklisted : atk_editable_text_cut_text

// Blacklisted : atk_editable_text_delete_text

// Blacklisted : atk_editable_text_insert_text

// Blacklisted : atk_editable_text_paste_text

// Blacklisted : atk_editable_text_set_run_attributes

// Blacklisted : atk_editable_text_set_text_contents

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

// Equals compares this HyperlinkImpl with another HyperlinkImpl, and returns true if they represent the same GObject.
func (recv *HyperlinkImpl) Equals(other *HyperlinkImpl) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this Hypertext with another Hypertext, and returns true if they represent the same GObject.
func (recv *Hypertext) Equals(other *Hypertext) bool {
	return other.ToC() == recv.ToC()
}

type signalHypertextLinkSelectedDetail struct {
	callback  HypertextSignalLinkSelectedCallback
	handlerID C.gulong
}

var signalHypertextLinkSelectedId int
var signalHypertextLinkSelectedMap = make(map[int]signalHypertextLinkSelectedDetail)
var signalHypertextLinkSelectedLock sync.RWMutex

// HypertextSignalLinkSelectedCallback is a callback function for a 'link-selected' signal emitted from a Hypertext.
type HypertextSignalLinkSelectedCallback func(arg1 int32)

/*
ConnectLinkSelected connects the callback to the 'link-selected' signal for the Hypertext.

The returned value represents the connection, and may be passed to DisconnectLinkSelected to remove it.
*/
func (recv *Hypertext) ConnectLinkSelected(callback HypertextSignalLinkSelectedCallback) int {
	signalHypertextLinkSelectedLock.Lock()
	defer signalHypertextLinkSelectedLock.Unlock()

	signalHypertextLinkSelectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Hypertext_signal_connect_link_selected(instance, C.gpointer(uintptr(signalHypertextLinkSelectedId)))

	detail := signalHypertextLinkSelectedDetail{callback, handlerID}
	signalHypertextLinkSelectedMap[signalHypertextLinkSelectedId] = detail

	return signalHypertextLinkSelectedId
}

/*
DisconnectLinkSelected disconnects a callback from the 'link-selected' signal for the Hypertext.

The connectionID should be a value returned from a call to ConnectLinkSelected.
*/
func (recv *Hypertext) DisconnectLinkSelected(connectionID int) {
	signalHypertextLinkSelectedLock.Lock()
	defer signalHypertextLinkSelectedLock.Unlock()

	detail, exists := signalHypertextLinkSelectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalHypertextLinkSelectedMap, connectionID)
}

//export hypertext_linkSelectedHandler
func hypertext_linkSelectedHandler(_ *C.GObject, c_arg1 C.gint, data C.gpointer) {
	signalHypertextLinkSelectedLock.RLock()
	defer signalHypertextLinkSelectedLock.RUnlock()

	arg1 := int32(c_arg1)

	index := int(uintptr(data))
	callback := signalHypertextLinkSelectedMap[index].callback
	callback(arg1)
}

// Blacklisted : atk_hypertext_get_link

// Blacklisted : atk_hypertext_get_link_index

// Blacklisted : atk_hypertext_get_n_links

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

// Equals compares this Image with another Image, and returns true if they represent the same GObject.
func (recv *Image) Equals(other *Image) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : atk_image_get_image_description

// Blacklisted : atk_image_get_image_position

// Blacklisted : atk_image_get_image_size

// Blacklisted : atk_image_set_image_description

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

// Equals compares this ImplementorIface with another ImplementorIface, and returns true if they represent the same GObject.
func (recv *ImplementorIface) Equals(other *ImplementorIface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this Selection with another Selection, and returns true if they represent the same GObject.
func (recv *Selection) Equals(other *Selection) bool {
	return other.ToC() == recv.ToC()
}

type signalSelectionSelectionChangedDetail struct {
	callback  SelectionSignalSelectionChangedCallback
	handlerID C.gulong
}

var signalSelectionSelectionChangedId int
var signalSelectionSelectionChangedMap = make(map[int]signalSelectionSelectionChangedDetail)
var signalSelectionSelectionChangedLock sync.RWMutex

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
	signalSelectionSelectionChangedLock.RLock()
	defer signalSelectionSelectionChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalSelectionSelectionChangedMap[index].callback
	callback()
}

// Blacklisted : atk_selection_add_selection

// Blacklisted : atk_selection_clear_selection

// Blacklisted : atk_selection_get_selection_count

// Blacklisted : atk_selection_is_child_selected

// Blacklisted : atk_selection_ref_selection

// Blacklisted : atk_selection_remove_selection

// Blacklisted : atk_selection_select_all_selection

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

// Equals compares this StreamableContent with another StreamableContent, and returns true if they represent the same GObject.
func (recv *StreamableContent) Equals(other *StreamableContent) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : atk_streamable_content_get_mime_type

// Blacklisted : atk_streamable_content_get_n_mime_types

// Blacklisted : atk_streamable_content_get_stream

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

// Equals compares this Table with another Table, and returns true if they represent the same GObject.
func (recv *Table) Equals(other *Table) bool {
	return other.ToC() == recv.ToC()
}

type signalTableColumnDeletedDetail struct {
	callback  TableSignalColumnDeletedCallback
	handlerID C.gulong
}

var signalTableColumnDeletedId int
var signalTableColumnDeletedMap = make(map[int]signalTableColumnDeletedDetail)
var signalTableColumnDeletedLock sync.RWMutex

// TableSignalColumnDeletedCallback is a callback function for a 'column-deleted' signal emitted from a Table.
type TableSignalColumnDeletedCallback func(arg1 int32, arg2 int32)

/*
ConnectColumnDeleted connects the callback to the 'column-deleted' signal for the Table.

The returned value represents the connection, and may be passed to DisconnectColumnDeleted to remove it.
*/
func (recv *Table) ConnectColumnDeleted(callback TableSignalColumnDeletedCallback) int {
	signalTableColumnDeletedLock.Lock()
	defer signalTableColumnDeletedLock.Unlock()

	signalTableColumnDeletedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Table_signal_connect_column_deleted(instance, C.gpointer(uintptr(signalTableColumnDeletedId)))

	detail := signalTableColumnDeletedDetail{callback, handlerID}
	signalTableColumnDeletedMap[signalTableColumnDeletedId] = detail

	return signalTableColumnDeletedId
}

/*
DisconnectColumnDeleted disconnects a callback from the 'column-deleted' signal for the Table.

The connectionID should be a value returned from a call to ConnectColumnDeleted.
*/
func (recv *Table) DisconnectColumnDeleted(connectionID int) {
	signalTableColumnDeletedLock.Lock()
	defer signalTableColumnDeletedLock.Unlock()

	detail, exists := signalTableColumnDeletedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTableColumnDeletedMap, connectionID)
}

//export table_columnDeletedHandler
func table_columnDeletedHandler(_ *C.GObject, c_arg1 C.gint, c_arg2 C.gint, data C.gpointer) {
	signalTableColumnDeletedLock.RLock()
	defer signalTableColumnDeletedLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	index := int(uintptr(data))
	callback := signalTableColumnDeletedMap[index].callback
	callback(arg1, arg2)
}

type signalTableColumnInsertedDetail struct {
	callback  TableSignalColumnInsertedCallback
	handlerID C.gulong
}

var signalTableColumnInsertedId int
var signalTableColumnInsertedMap = make(map[int]signalTableColumnInsertedDetail)
var signalTableColumnInsertedLock sync.RWMutex

// TableSignalColumnInsertedCallback is a callback function for a 'column-inserted' signal emitted from a Table.
type TableSignalColumnInsertedCallback func(arg1 int32, arg2 int32)

/*
ConnectColumnInserted connects the callback to the 'column-inserted' signal for the Table.

The returned value represents the connection, and may be passed to DisconnectColumnInserted to remove it.
*/
func (recv *Table) ConnectColumnInserted(callback TableSignalColumnInsertedCallback) int {
	signalTableColumnInsertedLock.Lock()
	defer signalTableColumnInsertedLock.Unlock()

	signalTableColumnInsertedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Table_signal_connect_column_inserted(instance, C.gpointer(uintptr(signalTableColumnInsertedId)))

	detail := signalTableColumnInsertedDetail{callback, handlerID}
	signalTableColumnInsertedMap[signalTableColumnInsertedId] = detail

	return signalTableColumnInsertedId
}

/*
DisconnectColumnInserted disconnects a callback from the 'column-inserted' signal for the Table.

The connectionID should be a value returned from a call to ConnectColumnInserted.
*/
func (recv *Table) DisconnectColumnInserted(connectionID int) {
	signalTableColumnInsertedLock.Lock()
	defer signalTableColumnInsertedLock.Unlock()

	detail, exists := signalTableColumnInsertedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTableColumnInsertedMap, connectionID)
}

//export table_columnInsertedHandler
func table_columnInsertedHandler(_ *C.GObject, c_arg1 C.gint, c_arg2 C.gint, data C.gpointer) {
	signalTableColumnInsertedLock.RLock()
	defer signalTableColumnInsertedLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	index := int(uintptr(data))
	callback := signalTableColumnInsertedMap[index].callback
	callback(arg1, arg2)
}

type signalTableColumnReorderedDetail struct {
	callback  TableSignalColumnReorderedCallback
	handlerID C.gulong
}

var signalTableColumnReorderedId int
var signalTableColumnReorderedMap = make(map[int]signalTableColumnReorderedDetail)
var signalTableColumnReorderedLock sync.RWMutex

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
	signalTableColumnReorderedLock.RLock()
	defer signalTableColumnReorderedLock.RUnlock()

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
var signalTableModelChangedLock sync.RWMutex

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
	signalTableModelChangedLock.RLock()
	defer signalTableModelChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalTableModelChangedMap[index].callback
	callback()
}

type signalTableRowDeletedDetail struct {
	callback  TableSignalRowDeletedCallback
	handlerID C.gulong
}

var signalTableRowDeletedId int
var signalTableRowDeletedMap = make(map[int]signalTableRowDeletedDetail)
var signalTableRowDeletedLock sync.RWMutex

// TableSignalRowDeletedCallback is a callback function for a 'row-deleted' signal emitted from a Table.
type TableSignalRowDeletedCallback func(arg1 int32, arg2 int32)

/*
ConnectRowDeleted connects the callback to the 'row-deleted' signal for the Table.

The returned value represents the connection, and may be passed to DisconnectRowDeleted to remove it.
*/
func (recv *Table) ConnectRowDeleted(callback TableSignalRowDeletedCallback) int {
	signalTableRowDeletedLock.Lock()
	defer signalTableRowDeletedLock.Unlock()

	signalTableRowDeletedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Table_signal_connect_row_deleted(instance, C.gpointer(uintptr(signalTableRowDeletedId)))

	detail := signalTableRowDeletedDetail{callback, handlerID}
	signalTableRowDeletedMap[signalTableRowDeletedId] = detail

	return signalTableRowDeletedId
}

/*
DisconnectRowDeleted disconnects a callback from the 'row-deleted' signal for the Table.

The connectionID should be a value returned from a call to ConnectRowDeleted.
*/
func (recv *Table) DisconnectRowDeleted(connectionID int) {
	signalTableRowDeletedLock.Lock()
	defer signalTableRowDeletedLock.Unlock()

	detail, exists := signalTableRowDeletedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTableRowDeletedMap, connectionID)
}

//export table_rowDeletedHandler
func table_rowDeletedHandler(_ *C.GObject, c_arg1 C.gint, c_arg2 C.gint, data C.gpointer) {
	signalTableRowDeletedLock.RLock()
	defer signalTableRowDeletedLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	index := int(uintptr(data))
	callback := signalTableRowDeletedMap[index].callback
	callback(arg1, arg2)
}

type signalTableRowInsertedDetail struct {
	callback  TableSignalRowInsertedCallback
	handlerID C.gulong
}

var signalTableRowInsertedId int
var signalTableRowInsertedMap = make(map[int]signalTableRowInsertedDetail)
var signalTableRowInsertedLock sync.RWMutex

// TableSignalRowInsertedCallback is a callback function for a 'row-inserted' signal emitted from a Table.
type TableSignalRowInsertedCallback func(arg1 int32, arg2 int32)

/*
ConnectRowInserted connects the callback to the 'row-inserted' signal for the Table.

The returned value represents the connection, and may be passed to DisconnectRowInserted to remove it.
*/
func (recv *Table) ConnectRowInserted(callback TableSignalRowInsertedCallback) int {
	signalTableRowInsertedLock.Lock()
	defer signalTableRowInsertedLock.Unlock()

	signalTableRowInsertedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Table_signal_connect_row_inserted(instance, C.gpointer(uintptr(signalTableRowInsertedId)))

	detail := signalTableRowInsertedDetail{callback, handlerID}
	signalTableRowInsertedMap[signalTableRowInsertedId] = detail

	return signalTableRowInsertedId
}

/*
DisconnectRowInserted disconnects a callback from the 'row-inserted' signal for the Table.

The connectionID should be a value returned from a call to ConnectRowInserted.
*/
func (recv *Table) DisconnectRowInserted(connectionID int) {
	signalTableRowInsertedLock.Lock()
	defer signalTableRowInsertedLock.Unlock()

	detail, exists := signalTableRowInsertedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTableRowInsertedMap, connectionID)
}

//export table_rowInsertedHandler
func table_rowInsertedHandler(_ *C.GObject, c_arg1 C.gint, c_arg2 C.gint, data C.gpointer) {
	signalTableRowInsertedLock.RLock()
	defer signalTableRowInsertedLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	index := int(uintptr(data))
	callback := signalTableRowInsertedMap[index].callback
	callback(arg1, arg2)
}

type signalTableRowReorderedDetail struct {
	callback  TableSignalRowReorderedCallback
	handlerID C.gulong
}

var signalTableRowReorderedId int
var signalTableRowReorderedMap = make(map[int]signalTableRowReorderedDetail)
var signalTableRowReorderedLock sync.RWMutex

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
	signalTableRowReorderedLock.RLock()
	defer signalTableRowReorderedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalTableRowReorderedMap[index].callback
	callback()
}

// Blacklisted : atk_table_add_column_selection

// Blacklisted : atk_table_add_row_selection

// Blacklisted : atk_table_get_caption

// Blacklisted : atk_table_get_column_at_index

// Blacklisted : atk_table_get_column_description

// Blacklisted : atk_table_get_column_extent_at

// Blacklisted : atk_table_get_column_header

// Blacklisted : atk_table_get_index_at

// Blacklisted : atk_table_get_n_columns

// Blacklisted : atk_table_get_n_rows

// Blacklisted : atk_table_get_row_at_index

// Blacklisted : atk_table_get_row_description

// Blacklisted : atk_table_get_row_extent_at

// Blacklisted : atk_table_get_row_header

// Unsupported : atk_table_get_selected_columns : unsupported parameter selected : gint** with indirection level of 2

// Unsupported : atk_table_get_selected_rows : unsupported parameter selected : gint** with indirection level of 2

// Blacklisted : atk_table_get_summary

// Blacklisted : atk_table_is_column_selected

// Blacklisted : atk_table_is_row_selected

// Blacklisted : atk_table_is_selected

// Blacklisted : atk_table_ref_at

// Blacklisted : atk_table_remove_column_selection

// Blacklisted : atk_table_remove_row_selection

// Blacklisted : atk_table_set_caption

// Blacklisted : atk_table_set_column_description

// Blacklisted : atk_table_set_column_header

// Blacklisted : atk_table_set_row_description

// Blacklisted : atk_table_set_row_header

// Blacklisted : atk_table_set_summary

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

// Equals compares this TableCell with another TableCell, and returns true if they represent the same GObject.
func (recv *TableCell) Equals(other *TableCell) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this Text with another Text, and returns true if they represent the same GObject.
func (recv *Text) Equals(other *Text) bool {
	return other.ToC() == recv.ToC()
}

type signalTextTextAttributesChangedDetail struct {
	callback  TextSignalTextAttributesChangedCallback
	handlerID C.gulong
}

var signalTextTextAttributesChangedId int
var signalTextTextAttributesChangedMap = make(map[int]signalTextTextAttributesChangedDetail)
var signalTextTextAttributesChangedLock sync.RWMutex

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
	signalTextTextAttributesChangedLock.RLock()
	defer signalTextTextAttributesChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalTextTextAttributesChangedMap[index].callback
	callback()
}

type signalTextTextCaretMovedDetail struct {
	callback  TextSignalTextCaretMovedCallback
	handlerID C.gulong
}

var signalTextTextCaretMovedId int
var signalTextTextCaretMovedMap = make(map[int]signalTextTextCaretMovedDetail)
var signalTextTextCaretMovedLock sync.RWMutex

// TextSignalTextCaretMovedCallback is a callback function for a 'text-caret-moved' signal emitted from a Text.
type TextSignalTextCaretMovedCallback func(arg1 int32)

/*
ConnectTextCaretMoved connects the callback to the 'text-caret-moved' signal for the Text.

The returned value represents the connection, and may be passed to DisconnectTextCaretMoved to remove it.
*/
func (recv *Text) ConnectTextCaretMoved(callback TextSignalTextCaretMovedCallback) int {
	signalTextTextCaretMovedLock.Lock()
	defer signalTextTextCaretMovedLock.Unlock()

	signalTextTextCaretMovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Text_signal_connect_text_caret_moved(instance, C.gpointer(uintptr(signalTextTextCaretMovedId)))

	detail := signalTextTextCaretMovedDetail{callback, handlerID}
	signalTextTextCaretMovedMap[signalTextTextCaretMovedId] = detail

	return signalTextTextCaretMovedId
}

/*
DisconnectTextCaretMoved disconnects a callback from the 'text-caret-moved' signal for the Text.

The connectionID should be a value returned from a call to ConnectTextCaretMoved.
*/
func (recv *Text) DisconnectTextCaretMoved(connectionID int) {
	signalTextTextCaretMovedLock.Lock()
	defer signalTextTextCaretMovedLock.Unlock()

	detail, exists := signalTextTextCaretMovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextTextCaretMovedMap, connectionID)
}

//export text_textCaretMovedHandler
func text_textCaretMovedHandler(_ *C.GObject, c_arg1 C.gint, data C.gpointer) {
	signalTextTextCaretMovedLock.RLock()
	defer signalTextTextCaretMovedLock.RUnlock()

	arg1 := int32(c_arg1)

	index := int(uintptr(data))
	callback := signalTextTextCaretMovedMap[index].callback
	callback(arg1)
}

type signalTextTextChangedDetail struct {
	callback  TextSignalTextChangedCallback
	handlerID C.gulong
}

var signalTextTextChangedId int
var signalTextTextChangedMap = make(map[int]signalTextTextChangedDetail)
var signalTextTextChangedLock sync.RWMutex

// TextSignalTextChangedCallback is a callback function for a 'text-changed' signal emitted from a Text.
type TextSignalTextChangedCallback func(arg1 int32, arg2 int32)

/*
ConnectTextChanged connects the callback to the 'text-changed' signal for the Text.

The returned value represents the connection, and may be passed to DisconnectTextChanged to remove it.
*/
func (recv *Text) ConnectTextChanged(callback TextSignalTextChangedCallback) int {
	signalTextTextChangedLock.Lock()
	defer signalTextTextChangedLock.Unlock()

	signalTextTextChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Text_signal_connect_text_changed(instance, C.gpointer(uintptr(signalTextTextChangedId)))

	detail := signalTextTextChangedDetail{callback, handlerID}
	signalTextTextChangedMap[signalTextTextChangedId] = detail

	return signalTextTextChangedId
}

/*
DisconnectTextChanged disconnects a callback from the 'text-changed' signal for the Text.

The connectionID should be a value returned from a call to ConnectTextChanged.
*/
func (recv *Text) DisconnectTextChanged(connectionID int) {
	signalTextTextChangedLock.Lock()
	defer signalTextTextChangedLock.Unlock()

	detail, exists := signalTextTextChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextTextChangedMap, connectionID)
}

//export text_textChangedHandler
func text_textChangedHandler(_ *C.GObject, c_arg1 C.gint, c_arg2 C.gint, data C.gpointer) {
	signalTextTextChangedLock.RLock()
	defer signalTextTextChangedLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	index := int(uintptr(data))
	callback := signalTextTextChangedMap[index].callback
	callback(arg1, arg2)
}

type signalTextTextInsertDetail struct {
	callback  TextSignalTextInsertCallback
	handlerID C.gulong
}

var signalTextTextInsertId int
var signalTextTextInsertMap = make(map[int]signalTextTextInsertDetail)
var signalTextTextInsertLock sync.RWMutex

// TextSignalTextInsertCallback is a callback function for a 'text-insert' signal emitted from a Text.
type TextSignalTextInsertCallback func(arg1 int32, arg2 int32, arg3 string)

/*
ConnectTextInsert connects the callback to the 'text-insert' signal for the Text.

The returned value represents the connection, and may be passed to DisconnectTextInsert to remove it.
*/
func (recv *Text) ConnectTextInsert(callback TextSignalTextInsertCallback) int {
	signalTextTextInsertLock.Lock()
	defer signalTextTextInsertLock.Unlock()

	signalTextTextInsertId++
	instance := C.gpointer(recv.native)
	handlerID := C.Text_signal_connect_text_insert(instance, C.gpointer(uintptr(signalTextTextInsertId)))

	detail := signalTextTextInsertDetail{callback, handlerID}
	signalTextTextInsertMap[signalTextTextInsertId] = detail

	return signalTextTextInsertId
}

/*
DisconnectTextInsert disconnects a callback from the 'text-insert' signal for the Text.

The connectionID should be a value returned from a call to ConnectTextInsert.
*/
func (recv *Text) DisconnectTextInsert(connectionID int) {
	signalTextTextInsertLock.Lock()
	defer signalTextTextInsertLock.Unlock()

	detail, exists := signalTextTextInsertMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextTextInsertMap, connectionID)
}

//export text_textInsertHandler
func text_textInsertHandler(_ *C.GObject, c_arg1 C.gint, c_arg2 C.gint, c_arg3 *C.gchar, data C.gpointer) {
	signalTextTextInsertLock.RLock()
	defer signalTextTextInsertLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	arg3 := C.GoString(c_arg3)

	index := int(uintptr(data))
	callback := signalTextTextInsertMap[index].callback
	callback(arg1, arg2, arg3)
}

type signalTextTextRemoveDetail struct {
	callback  TextSignalTextRemoveCallback
	handlerID C.gulong
}

var signalTextTextRemoveId int
var signalTextTextRemoveMap = make(map[int]signalTextTextRemoveDetail)
var signalTextTextRemoveLock sync.RWMutex

// TextSignalTextRemoveCallback is a callback function for a 'text-remove' signal emitted from a Text.
type TextSignalTextRemoveCallback func(arg1 int32, arg2 int32, arg3 string)

/*
ConnectTextRemove connects the callback to the 'text-remove' signal for the Text.

The returned value represents the connection, and may be passed to DisconnectTextRemove to remove it.
*/
func (recv *Text) ConnectTextRemove(callback TextSignalTextRemoveCallback) int {
	signalTextTextRemoveLock.Lock()
	defer signalTextTextRemoveLock.Unlock()

	signalTextTextRemoveId++
	instance := C.gpointer(recv.native)
	handlerID := C.Text_signal_connect_text_remove(instance, C.gpointer(uintptr(signalTextTextRemoveId)))

	detail := signalTextTextRemoveDetail{callback, handlerID}
	signalTextTextRemoveMap[signalTextTextRemoveId] = detail

	return signalTextTextRemoveId
}

/*
DisconnectTextRemove disconnects a callback from the 'text-remove' signal for the Text.

The connectionID should be a value returned from a call to ConnectTextRemove.
*/
func (recv *Text) DisconnectTextRemove(connectionID int) {
	signalTextTextRemoveLock.Lock()
	defer signalTextTextRemoveLock.Unlock()

	detail, exists := signalTextTextRemoveMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextTextRemoveMap, connectionID)
}

//export text_textRemoveHandler
func text_textRemoveHandler(_ *C.GObject, c_arg1 C.gint, c_arg2 C.gint, c_arg3 *C.gchar, data C.gpointer) {
	signalTextTextRemoveLock.RLock()
	defer signalTextTextRemoveLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	arg3 := C.GoString(c_arg3)

	index := int(uintptr(data))
	callback := signalTextTextRemoveMap[index].callback
	callback(arg1, arg2, arg3)
}

type signalTextTextSelectionChangedDetail struct {
	callback  TextSignalTextSelectionChangedCallback
	handlerID C.gulong
}

var signalTextTextSelectionChangedId int
var signalTextTextSelectionChangedMap = make(map[int]signalTextTextSelectionChangedDetail)
var signalTextTextSelectionChangedLock sync.RWMutex

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
	signalTextTextSelectionChangedLock.RLock()
	defer signalTextTextSelectionChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalTextTextSelectionChangedMap[index].callback
	callback()
}

// Blacklisted : atk_text_add_selection

// Blacklisted : atk_text_get_caret_offset

// Blacklisted : atk_text_get_character_at_offset

// Blacklisted : atk_text_get_character_count

// Blacklisted : atk_text_get_character_extents

// Blacklisted : atk_text_get_default_attributes

// Blacklisted : atk_text_get_n_selections

// Blacklisted : atk_text_get_offset_at_point

// Blacklisted : atk_text_get_run_attributes

// Blacklisted : atk_text_get_selection

// GetText is a wrapper around the C function atk_text_get_text.
func (recv *Text) GetText(startOffset int32, endOffset int32) string {
	c_start_offset := (C.gint)(startOffset)

	c_end_offset := (C.gint)(endOffset)

	retC := C.atk_text_get_text((*C.AtkText)(recv.native), c_start_offset, c_end_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : atk_text_get_text_after_offset

// Blacklisted : atk_text_get_text_at_offset

// Blacklisted : atk_text_get_text_before_offset

// Blacklisted : atk_text_remove_selection

// Blacklisted : atk_text_set_caret_offset

// Blacklisted : atk_text_set_selection

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

// Equals compares this Value with another Value, and returns true if they represent the same GObject.
func (recv *Value) Equals(other *Value) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : atk_value_get_current_value

// Blacklisted : atk_value_get_maximum_value

// Blacklisted : atk_value_get_minimum_value

// Blacklisted : atk_value_set_current_value

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

// Equals compares this Window with another Window, and returns true if they represent the same GObject.
func (recv *Window) Equals(other *Window) bool {
	return other.ToC() == recv.ToC()
}
