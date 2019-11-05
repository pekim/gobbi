// Code generated - DO NOT EDIT.
// +build !atk_1.3,!atk_1.4,!atk_1.6,!atk_1.9,!atk_1.12,!atk_1.13,!atk_1.20,!atk_1.30,!atk_2.8,!atk_2.12

package atk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
/*

	void hyperlink_linkActivatedHandler(GObject *, gpointer);

	static gulong Hyperlink_signal_connect_link_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "link-activated", G_CALLBACK(hyperlink_linkActivatedHandler), data);
	}

*/
/*

	void object_focusEventHandler(GObject *, gboolean, gpointer);

	static gulong Object_signal_connect_focus_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-event", G_CALLBACK(object_focusEventHandler), data);
	}

*/
/*

	void object_stateChangeHandler(GObject *, gchar*, gboolean, gpointer);

	static gulong Object_signal_connect_state_change(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-change", G_CALLBACK(object_stateChangeHandler), data);
	}

*/
/*

	void object_visibleDataChangedHandler(GObject *, gpointer);

	static gulong Object_signal_connect_visible_data_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "visible-data-changed", G_CALLBACK(object_visibleDataChangedHandler), data);
	}

*/
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

// AttributeSet is a representation of the C alias AtkAttributeSet.
type AttributeSet *glib.SList

// State is a representation of the C alias AtkState.
type State uint64

type HyperlinkStateFlags C.AtkHyperlinkStateFlags

const (
	ATK_HYPERLINK_IS_INLINE HyperlinkStateFlags = 1
)

// GObjectAccessible is a wrapper around the C record AtkGObjectAccessible.
type GObjectAccessible struct {
	native *C.AtkGObjectAccessible
	// parent : record
}

func GObjectAccessibleNewFromC(u unsafe.Pointer) *GObjectAccessible {
	c := (*C.AtkGObjectAccessible)(u)
	if c == nil {
		return nil
	}

	g := &GObjectAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GObjectAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GObjectAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GObjectAccessible with another GObjectAccessible, and returns true if they represent the same GObject.
func (recv *GObjectAccessible) Equals(other *GObjectAccessible) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *GObjectAccessible) Object() *Object {
	return ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to GObjectAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a GObjectAccessible.
func CastToGObjectAccessible(object *gobject.Object) *GObjectAccessible {
	return GObjectAccessibleNewFromC(object.ToC())
}

// GObjectAccessibleForObject is a wrapper around the C function atk_gobject_accessible_for_object.
func GObjectAccessibleForObject(obj *gobject.Object) *Object {
	c_obj := (*C.GObject)(C.NULL)
	if obj != nil {
		c_obj = (*C.GObject)(obj.ToC())
	}

	retC := C.atk_gobject_accessible_for_object(c_obj)
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetObject is a wrapper around the C function atk_gobject_accessible_get_object.
func (recv *GObjectAccessible) GetObject() *gobject.Object {
	retC := C.atk_gobject_accessible_get_object((*C.AtkGObjectAccessible)(recv.native))
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Hyperlink is a wrapper around the C record AtkHyperlink.
type Hyperlink struct {
	native *C.AtkHyperlink
	// parent : record
}

func HyperlinkNewFromC(u unsafe.Pointer) *Hyperlink {
	c := (*C.AtkHyperlink)(u)
	if c == nil {
		return nil
	}

	g := &Hyperlink{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Hyperlink) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Hyperlink) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Hyperlink with another Hyperlink, and returns true if they represent the same GObject.
func (recv *Hyperlink) Equals(other *Hyperlink) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Hyperlink) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Hyperlink.
// Exercise care, as this is a potentially dangerous function if the Object is not a Hyperlink.
func CastToHyperlink(object *gobject.Object) *Hyperlink {
	return HyperlinkNewFromC(object.ToC())
}

type signalHyperlinkLinkActivatedDetail struct {
	callback  HyperlinkSignalLinkActivatedCallback
	handlerID C.gulong
}

var signalHyperlinkLinkActivatedId int
var signalHyperlinkLinkActivatedMap = make(map[int]signalHyperlinkLinkActivatedDetail)
var signalHyperlinkLinkActivatedLock sync.RWMutex

// HyperlinkSignalLinkActivatedCallback is a callback function for a 'link-activated' signal emitted from a Hyperlink.
type HyperlinkSignalLinkActivatedCallback func(targetObject *Hyperlink)

/*
ConnectLinkActivated connects the callback to the 'link-activated' signal for the Hyperlink.

The returned value represents the connection, and may be passed to DisconnectLinkActivated to remove it.
*/
func (recv *Hyperlink) ConnectLinkActivated(callback HyperlinkSignalLinkActivatedCallback) int {
	signalHyperlinkLinkActivatedLock.Lock()
	defer signalHyperlinkLinkActivatedLock.Unlock()

	signalHyperlinkLinkActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Hyperlink_signal_connect_link_activated(instance, C.gpointer(uintptr(signalHyperlinkLinkActivatedId)))

	detail := signalHyperlinkLinkActivatedDetail{callback, handlerID}
	signalHyperlinkLinkActivatedMap[signalHyperlinkLinkActivatedId] = detail

	return signalHyperlinkLinkActivatedId
}

/*
DisconnectLinkActivated disconnects a callback from the 'link-activated' signal for the Hyperlink.

The connectionID should be a value returned from a call to ConnectLinkActivated.
*/
func (recv *Hyperlink) DisconnectLinkActivated(connectionID int) {
	signalHyperlinkLinkActivatedLock.Lock()
	defer signalHyperlinkLinkActivatedLock.Unlock()

	detail, exists := signalHyperlinkLinkActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalHyperlinkLinkActivatedMap, connectionID)
}

//export hyperlink_linkActivatedHandler
func hyperlink_linkActivatedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalHyperlinkLinkActivatedLock.RLock()
	defer signalHyperlinkLinkActivatedLock.RUnlock()

	targetObject := HyperlinkNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalHyperlinkLinkActivatedMap[index].callback
	callback(targetObject)
}

// GetEndIndex is a wrapper around the C function atk_hyperlink_get_end_index.
func (recv *Hyperlink) GetEndIndex() int32 {
	retC := C.atk_hyperlink_get_end_index((*C.AtkHyperlink)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNAnchors is a wrapper around the C function atk_hyperlink_get_n_anchors.
func (recv *Hyperlink) GetNAnchors() int32 {
	retC := C.atk_hyperlink_get_n_anchors((*C.AtkHyperlink)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetObject is a wrapper around the C function atk_hyperlink_get_object.
func (recv *Hyperlink) GetObject(i int32) *Object {
	c_i := (C.gint)(i)

	retC := C.atk_hyperlink_get_object((*C.AtkHyperlink)(recv.native), c_i)
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStartIndex is a wrapper around the C function atk_hyperlink_get_start_index.
func (recv *Hyperlink) GetStartIndex() int32 {
	retC := C.atk_hyperlink_get_start_index((*C.AtkHyperlink)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetUri is a wrapper around the C function atk_hyperlink_get_uri.
func (recv *Hyperlink) GetUri(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_hyperlink_get_uri((*C.AtkHyperlink)(recv.native), c_i)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// IsInline is a wrapper around the C function atk_hyperlink_is_inline.
func (recv *Hyperlink) IsInline() bool {
	retC := C.atk_hyperlink_is_inline((*C.AtkHyperlink)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsValid is a wrapper around the C function atk_hyperlink_is_valid.
func (recv *Hyperlink) IsValid() bool {
	retC := C.atk_hyperlink_is_valid((*C.AtkHyperlink)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Action returns the Action interface implemented by Hyperlink
func (recv *Hyperlink) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// Misc is a wrapper around the C record AtkMisc.
type Misc struct {
	native *C.AtkMisc
	// parent : record
}

func MiscNewFromC(u unsafe.Pointer) *Misc {
	c := (*C.AtkMisc)(u)
	if c == nil {
		return nil
	}

	g := &Misc{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Misc) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Misc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Misc with another Misc, and returns true if they represent the same GObject.
func (recv *Misc) Equals(other *Misc) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Misc) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Misc.
// Exercise care, as this is a potentially dangerous function if the Object is not a Misc.
func CastToMisc(object *gobject.Object) *Misc {
	return MiscNewFromC(object.ToC())
}

// NoOpObject is a wrapper around the C record AtkNoOpObject.
type NoOpObject struct {
	native *C.AtkNoOpObject
	// parent : record
}

func NoOpObjectNewFromC(u unsafe.Pointer) *NoOpObject {
	c := (*C.AtkNoOpObject)(u)
	if c == nil {
		return nil
	}

	g := &NoOpObject{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NoOpObject) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NoOpObject) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NoOpObject with another NoOpObject, and returns true if they represent the same GObject.
func (recv *NoOpObject) Equals(other *NoOpObject) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *NoOpObject) Object() *Object {
	return ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to NoOpObject.
// Exercise care, as this is a potentially dangerous function if the Object is not a NoOpObject.
func CastToNoOpObject(object *gobject.Object) *NoOpObject {
	return NoOpObjectNewFromC(object.ToC())
}

// NoOpObjectNew is a wrapper around the C function atk_no_op_object_new.
func NoOpObjectNew(obj *gobject.Object) *NoOpObject {
	c_obj := (*C.GObject)(C.NULL)
	if obj != nil {
		c_obj = (*C.GObject)(obj.ToC())
	}

	retC := C.atk_no_op_object_new(c_obj)
	retGo := NoOpObjectNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Action returns the Action interface implemented by NoOpObject
func (recv *NoOpObject) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by NoOpObject
func (recv *NoOpObject) Component() *Component {
	return ComponentNewFromC(recv.ToC())
}

// Document returns the Document interface implemented by NoOpObject
func (recv *NoOpObject) Document() *Document {
	return DocumentNewFromC(recv.ToC())
}

// EditableText returns the EditableText interface implemented by NoOpObject
func (recv *NoOpObject) EditableText() *EditableText {
	return EditableTextNewFromC(recv.ToC())
}

// Hypertext returns the Hypertext interface implemented by NoOpObject
func (recv *NoOpObject) Hypertext() *Hypertext {
	return HypertextNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by NoOpObject
func (recv *NoOpObject) Image() *Image {
	return ImageNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by NoOpObject
func (recv *NoOpObject) Selection() *Selection {
	return SelectionNewFromC(recv.ToC())
}

// Table returns the Table interface implemented by NoOpObject
func (recv *NoOpObject) Table() *Table {
	return TableNewFromC(recv.ToC())
}

// TableCell returns the TableCell interface implemented by NoOpObject
func (recv *NoOpObject) TableCell() *TableCell {
	return TableCellNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by NoOpObject
func (recv *NoOpObject) Text() *Text {
	return TextNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by NoOpObject
func (recv *NoOpObject) Value() *Value {
	return ValueNewFromC(recv.ToC())
}

// Window returns the Window interface implemented by NoOpObject
func (recv *NoOpObject) Window() *Window {
	return WindowNewFromC(recv.ToC())
}

// NoOpObjectFactory is a wrapper around the C record AtkNoOpObjectFactory.
type NoOpObjectFactory struct {
	native *C.AtkNoOpObjectFactory
	// parent : record
}

func NoOpObjectFactoryNewFromC(u unsafe.Pointer) *NoOpObjectFactory {
	c := (*C.AtkNoOpObjectFactory)(u)
	if c == nil {
		return nil
	}

	g := &NoOpObjectFactory{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NoOpObjectFactory) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NoOpObjectFactory) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NoOpObjectFactory with another NoOpObjectFactory, and returns true if they represent the same GObject.
func (recv *NoOpObjectFactory) Equals(other *NoOpObjectFactory) bool {
	return other.ToC() == recv.ToC()
}

// ObjectFactory upcasts to *ObjectFactory
func (recv *NoOpObjectFactory) ObjectFactory() *ObjectFactory {
	return ObjectFactoryNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *NoOpObjectFactory) Object() *gobject.Object {
	return recv.ObjectFactory().Object()
}

// CastToWidget down casts any arbitrary Object to NoOpObjectFactory.
// Exercise care, as this is a potentially dangerous function if the Object is not a NoOpObjectFactory.
func CastToNoOpObjectFactory(object *gobject.Object) *NoOpObjectFactory {
	return NoOpObjectFactoryNewFromC(object.ToC())
}

// NoOpObjectFactoryNew is a wrapper around the C function atk_no_op_object_factory_new.
func NoOpObjectFactoryNew() *NoOpObjectFactory {
	retC := C.atk_no_op_object_factory_new()
	retGo := NoOpObjectFactoryNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Object is a wrapper around the C record AtkObject.
type Object struct {
	native *C.AtkObject
	// parent : record
	Description string
	Name        string
	// accessible_parent : record
	Role Role
	// relation_set : record
	Layer Layer
}

func ObjectNewFromC(u unsafe.Pointer) *Object {
	c := (*C.AtkObject)(u)
	if c == nil {
		return nil
	}

	g := &Object{
		Description: C.GoString(c.description),
		Layer:       (Layer)(c.layer),
		Name:        C.GoString(c.name),
		Role:        (Role)(c.role),
		native:      c,
	}

	return g
}

func (recv *Object) ToC() unsafe.Pointer {
	recv.native.description =
		C.CString(recv.Description)
	recv.native.name =
		C.CString(recv.Name)
	recv.native.role =
		(C.AtkRole)(recv.Role)
	recv.native.layer =
		(C.AtkLayer)(recv.Layer)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Object with another Object, and returns true if they represent the same GObject.
func (recv *Object) Equals(other *Object) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Object) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Object.
// Exercise care, as this is a potentially dangerous function if the Object is not a Object.
func CastToObject(object *gobject.Object) *Object {
	return ObjectNewFromC(object.ToC())
}

// Unsupported signal 'active-descendant-changed' for Object : param arg1 : gpointer

// Unsupported signal 'children-changed' for Object : param arg2 : gpointer

type signalObjectFocusEventDetail struct {
	callback  ObjectSignalFocusEventCallback
	handlerID C.gulong
}

var signalObjectFocusEventId int
var signalObjectFocusEventMap = make(map[int]signalObjectFocusEventDetail)
var signalObjectFocusEventLock sync.RWMutex

// ObjectSignalFocusEventCallback is a callback function for a 'focus-event' signal emitted from a Object.
type ObjectSignalFocusEventCallback func(targetObject *Object, arg1 bool)

/*
ConnectFocusEvent connects the callback to the 'focus-event' signal for the Object.

The returned value represents the connection, and may be passed to DisconnectFocusEvent to remove it.
*/
func (recv *Object) ConnectFocusEvent(callback ObjectSignalFocusEventCallback) int {
	signalObjectFocusEventLock.Lock()
	defer signalObjectFocusEventLock.Unlock()

	signalObjectFocusEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Object_signal_connect_focus_event(instance, C.gpointer(uintptr(signalObjectFocusEventId)))

	detail := signalObjectFocusEventDetail{callback, handlerID}
	signalObjectFocusEventMap[signalObjectFocusEventId] = detail

	return signalObjectFocusEventId
}

/*
DisconnectFocusEvent disconnects a callback from the 'focus-event' signal for the Object.

The connectionID should be a value returned from a call to ConnectFocusEvent.
*/
func (recv *Object) DisconnectFocusEvent(connectionID int) {
	signalObjectFocusEventLock.Lock()
	defer signalObjectFocusEventLock.Unlock()

	detail, exists := signalObjectFocusEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalObjectFocusEventMap, connectionID)
}

//export object_focusEventHandler
func object_focusEventHandler(c_targetObject *C.GObject, c_arg1 C.gboolean, data C.gpointer) {
	signalObjectFocusEventLock.RLock()
	defer signalObjectFocusEventLock.RUnlock()

	arg1 := c_arg1 == C.TRUE

	targetObject := ObjectNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalObjectFocusEventMap[index].callback
	callback(targetObject, arg1)
}

// Unsupported signal 'property-change' for Object : param arg1 : gpointer

type signalObjectStateChangeDetail struct {
	callback  ObjectSignalStateChangeCallback
	handlerID C.gulong
}

var signalObjectStateChangeId int
var signalObjectStateChangeMap = make(map[int]signalObjectStateChangeDetail)
var signalObjectStateChangeLock sync.RWMutex

// ObjectSignalStateChangeCallback is a callback function for a 'state-change' signal emitted from a Object.
type ObjectSignalStateChangeCallback func(targetObject *Object, arg1 string, arg2 bool)

/*
ConnectStateChange connects the callback to the 'state-change' signal for the Object.

The returned value represents the connection, and may be passed to DisconnectStateChange to remove it.
*/
func (recv *Object) ConnectStateChange(callback ObjectSignalStateChangeCallback) int {
	signalObjectStateChangeLock.Lock()
	defer signalObjectStateChangeLock.Unlock()

	signalObjectStateChangeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Object_signal_connect_state_change(instance, C.gpointer(uintptr(signalObjectStateChangeId)))

	detail := signalObjectStateChangeDetail{callback, handlerID}
	signalObjectStateChangeMap[signalObjectStateChangeId] = detail

	return signalObjectStateChangeId
}

/*
DisconnectStateChange disconnects a callback from the 'state-change' signal for the Object.

The connectionID should be a value returned from a call to ConnectStateChange.
*/
func (recv *Object) DisconnectStateChange(connectionID int) {
	signalObjectStateChangeLock.Lock()
	defer signalObjectStateChangeLock.Unlock()

	detail, exists := signalObjectStateChangeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalObjectStateChangeMap, connectionID)
}

//export object_stateChangeHandler
func object_stateChangeHandler(c_targetObject *C.GObject, c_arg1 *C.gchar, c_arg2 C.gboolean, data C.gpointer) {
	signalObjectStateChangeLock.RLock()
	defer signalObjectStateChangeLock.RUnlock()

	arg1 := C.GoString(c_arg1)

	arg2 := c_arg2 == C.TRUE

	targetObject := ObjectNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalObjectStateChangeMap[index].callback
	callback(targetObject, arg1, arg2)
}

type signalObjectVisibleDataChangedDetail struct {
	callback  ObjectSignalVisibleDataChangedCallback
	handlerID C.gulong
}

var signalObjectVisibleDataChangedId int
var signalObjectVisibleDataChangedMap = make(map[int]signalObjectVisibleDataChangedDetail)
var signalObjectVisibleDataChangedLock sync.RWMutex

// ObjectSignalVisibleDataChangedCallback is a callback function for a 'visible-data-changed' signal emitted from a Object.
type ObjectSignalVisibleDataChangedCallback func(targetObject *Object)

/*
ConnectVisibleDataChanged connects the callback to the 'visible-data-changed' signal for the Object.

The returned value represents the connection, and may be passed to DisconnectVisibleDataChanged to remove it.
*/
func (recv *Object) ConnectVisibleDataChanged(callback ObjectSignalVisibleDataChangedCallback) int {
	signalObjectVisibleDataChangedLock.Lock()
	defer signalObjectVisibleDataChangedLock.Unlock()

	signalObjectVisibleDataChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Object_signal_connect_visible_data_changed(instance, C.gpointer(uintptr(signalObjectVisibleDataChangedId)))

	detail := signalObjectVisibleDataChangedDetail{callback, handlerID}
	signalObjectVisibleDataChangedMap[signalObjectVisibleDataChangedId] = detail

	return signalObjectVisibleDataChangedId
}

/*
DisconnectVisibleDataChanged disconnects a callback from the 'visible-data-changed' signal for the Object.

The connectionID should be a value returned from a call to ConnectVisibleDataChanged.
*/
func (recv *Object) DisconnectVisibleDataChanged(connectionID int) {
	signalObjectVisibleDataChangedLock.Lock()
	defer signalObjectVisibleDataChangedLock.Unlock()

	detail, exists := signalObjectVisibleDataChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalObjectVisibleDataChangedMap, connectionID)
}

//export object_visibleDataChangedHandler
func object_visibleDataChangedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalObjectVisibleDataChangedLock.RLock()
	defer signalObjectVisibleDataChangedLock.RUnlock()

	targetObject := ObjectNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalObjectVisibleDataChangedMap[index].callback
	callback(targetObject)
}

// AddRelationship is a wrapper around the C function atk_object_add_relationship.
func (recv *Object) AddRelationship(relationship RelationType, target *Object) bool {
	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(C.NULL)
	if target != nil {
		c_target = (*C.AtkObject)(target.ToC())
	}

	retC := C.atk_object_add_relationship((*C.AtkObject)(recv.native), c_relationship, c_target)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : atk_object_connect_property_change_handler : unsupported parameter handler : no type generator for PropertyChangeHandler (AtkPropertyChangeHandler*) for param handler

// GetDescription is a wrapper around the C function atk_object_get_description.
func (recv *Object) GetDescription() string {
	retC := C.atk_object_get_description((*C.AtkObject)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetIndexInParent is a wrapper around the C function atk_object_get_index_in_parent.
func (recv *Object) GetIndexInParent() int32 {
	retC := C.atk_object_get_index_in_parent((*C.AtkObject)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLayer is a wrapper around the C function atk_object_get_layer.
func (recv *Object) GetLayer() Layer {
	retC := C.atk_object_get_layer((*C.AtkObject)(recv.native))
	retGo := (Layer)(retC)

	return retGo
}

// GetMdiZorder is a wrapper around the C function atk_object_get_mdi_zorder.
func (recv *Object) GetMdiZorder() int32 {
	retC := C.atk_object_get_mdi_zorder((*C.AtkObject)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNAccessibleChildren is a wrapper around the C function atk_object_get_n_accessible_children.
func (recv *Object) GetNAccessibleChildren() int32 {
	retC := C.atk_object_get_n_accessible_children((*C.AtkObject)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetName is a wrapper around the C function atk_object_get_name.
func (recv *Object) GetName() string {
	retC := C.atk_object_get_name((*C.AtkObject)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetParent is a wrapper around the C function atk_object_get_parent.
func (recv *Object) GetParent() *Object {
	retC := C.atk_object_get_parent((*C.AtkObject)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRole is a wrapper around the C function atk_object_get_role.
func (recv *Object) GetRole() Role {
	retC := C.atk_object_get_role((*C.AtkObject)(recv.native))
	retGo := (Role)(retC)

	return retGo
}

// Initialize is a wrapper around the C function atk_object_initialize.
func (recv *Object) Initialize(data uintptr) {
	c_data := (C.gpointer)(data)

	C.atk_object_initialize((*C.AtkObject)(recv.native), c_data)

	return
}

// NotifyStateChange is a wrapper around the C function atk_object_notify_state_change.
func (recv *Object) NotifyStateChange(state State, value bool) {
	c_state := (C.AtkState)(state)

	c_value :=
		boolToGboolean(value)

	C.atk_object_notify_state_change((*C.AtkObject)(recv.native), c_state, c_value)

	return
}

// PeekParent is a wrapper around the C function atk_object_peek_parent.
func (recv *Object) PeekParent() *Object {
	retC := C.atk_object_peek_parent((*C.AtkObject)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RefAccessibleChild is a wrapper around the C function atk_object_ref_accessible_child.
func (recv *Object) RefAccessibleChild(i int32) *Object {
	c_i := (C.gint)(i)

	retC := C.atk_object_ref_accessible_child((*C.AtkObject)(recv.native), c_i)
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RefRelationSet is a wrapper around the C function atk_object_ref_relation_set.
func (recv *Object) RefRelationSet() *RelationSet {
	retC := C.atk_object_ref_relation_set((*C.AtkObject)(recv.native))
	retGo := RelationSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RefStateSet is a wrapper around the C function atk_object_ref_state_set.
func (recv *Object) RefStateSet() *StateSet {
	retC := C.atk_object_ref_state_set((*C.AtkObject)(recv.native))
	retGo := StateSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemovePropertyChangeHandler is a wrapper around the C function atk_object_remove_property_change_handler.
func (recv *Object) RemovePropertyChangeHandler(handlerId uint32) {
	c_handler_id := (C.guint)(handlerId)

	C.atk_object_remove_property_change_handler((*C.AtkObject)(recv.native), c_handler_id)

	return
}

// RemoveRelationship is a wrapper around the C function atk_object_remove_relationship.
func (recv *Object) RemoveRelationship(relationship RelationType, target *Object) bool {
	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(C.NULL)
	if target != nil {
		c_target = (*C.AtkObject)(target.ToC())
	}

	retC := C.atk_object_remove_relationship((*C.AtkObject)(recv.native), c_relationship, c_target)
	retGo := retC == C.TRUE

	return retGo
}

// SetDescription is a wrapper around the C function atk_object_set_description.
func (recv *Object) SetDescription(description string) {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.atk_object_set_description((*C.AtkObject)(recv.native), c_description)

	return
}

// SetName is a wrapper around the C function atk_object_set_name.
func (recv *Object) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.atk_object_set_name((*C.AtkObject)(recv.native), c_name)

	return
}

// SetParent is a wrapper around the C function atk_object_set_parent.
func (recv *Object) SetParent(parent *Object) {
	c_parent := (*C.AtkObject)(C.NULL)
	if parent != nil {
		c_parent = (*C.AtkObject)(parent.ToC())
	}

	C.atk_object_set_parent((*C.AtkObject)(recv.native), c_parent)

	return
}

// SetRole is a wrapper around the C function atk_object_set_role.
func (recv *Object) SetRole(role Role) {
	c_role := (C.AtkRole)(role)

	C.atk_object_set_role((*C.AtkObject)(recv.native), c_role)

	return
}

// ObjectFactory is a wrapper around the C record AtkObjectFactory.
type ObjectFactory struct {
	native *C.AtkObjectFactory
	// parent : record
}

func ObjectFactoryNewFromC(u unsafe.Pointer) *ObjectFactory {
	c := (*C.AtkObjectFactory)(u)
	if c == nil {
		return nil
	}

	g := &ObjectFactory{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ObjectFactory) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ObjectFactory) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ObjectFactory with another ObjectFactory, and returns true if they represent the same GObject.
func (recv *ObjectFactory) Equals(other *ObjectFactory) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ObjectFactory) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ObjectFactory.
// Exercise care, as this is a potentially dangerous function if the Object is not a ObjectFactory.
func CastToObjectFactory(object *gobject.Object) *ObjectFactory {
	return ObjectFactoryNewFromC(object.ToC())
}

// CreateAccessible is a wrapper around the C function atk_object_factory_create_accessible.
func (recv *ObjectFactory) CreateAccessible(obj *gobject.Object) *Object {
	c_obj := (*C.GObject)(C.NULL)
	if obj != nil {
		c_obj = (*C.GObject)(obj.ToC())
	}

	retC := C.atk_object_factory_create_accessible((*C.AtkObjectFactory)(recv.native), c_obj)
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAccessibleType is a wrapper around the C function atk_object_factory_get_accessible_type.
func (recv *ObjectFactory) GetAccessibleType() gobject.Type {
	retC := C.atk_object_factory_get_accessible_type((*C.AtkObjectFactory)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// Invalidate is a wrapper around the C function atk_object_factory_invalidate.
func (recv *ObjectFactory) Invalidate() {
	C.atk_object_factory_invalidate((*C.AtkObjectFactory)(recv.native))

	return
}

// Plug is a wrapper around the C record AtkPlug.
type Plug struct {
	native *C.AtkPlug
	// parent : record
}

func PlugNewFromC(u unsafe.Pointer) *Plug {
	c := (*C.AtkPlug)(u)
	if c == nil {
		return nil
	}

	g := &Plug{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Plug) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Plug) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Plug with another Plug, and returns true if they represent the same GObject.
func (recv *Plug) Equals(other *Plug) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Plug) Object() *Object {
	return ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Plug.
// Exercise care, as this is a potentially dangerous function if the Object is not a Plug.
func CastToPlug(object *gobject.Object) *Plug {
	return PlugNewFromC(object.ToC())
}

// PlugNew is a wrapper around the C function atk_plug_new.
func PlugNew() *Plug {
	retC := C.atk_plug_new()
	retGo := PlugNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Component returns the Component interface implemented by Plug
func (recv *Plug) Component() *Component {
	return ComponentNewFromC(recv.ToC())
}

// Registry is a wrapper around the C record AtkRegistry.
type Registry struct {
	native *C.AtkRegistry
	// parent : record
	// factory_type_registry : record
	// factory_singleton_cache : record
}

func RegistryNewFromC(u unsafe.Pointer) *Registry {
	c := (*C.AtkRegistry)(u)
	if c == nil {
		return nil
	}

	g := &Registry{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Registry) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Registry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Registry with another Registry, and returns true if they represent the same GObject.
func (recv *Registry) Equals(other *Registry) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Registry) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Registry.
// Exercise care, as this is a potentially dangerous function if the Object is not a Registry.
func CastToRegistry(object *gobject.Object) *Registry {
	return RegistryNewFromC(object.ToC())
}

// GetFactory is a wrapper around the C function atk_registry_get_factory.
func (recv *Registry) GetFactory(type_ gobject.Type) *ObjectFactory {
	c_type := (C.GType)(type_)

	retC := C.atk_registry_get_factory((*C.AtkRegistry)(recv.native), c_type)
	retGo := ObjectFactoryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFactoryType is a wrapper around the C function atk_registry_get_factory_type.
func (recv *Registry) GetFactoryType(type_ gobject.Type) gobject.Type {
	c_type := (C.GType)(type_)

	retC := C.atk_registry_get_factory_type((*C.AtkRegistry)(recv.native), c_type)
	retGo := (gobject.Type)(retC)

	return retGo
}

// SetFactoryType is a wrapper around the C function atk_registry_set_factory_type.
func (recv *Registry) SetFactoryType(type_ gobject.Type, factoryType gobject.Type) {
	c_type := (C.GType)(type_)

	c_factory_type := (C.GType)(factoryType)

	C.atk_registry_set_factory_type((*C.AtkRegistry)(recv.native), c_type, c_factory_type)

	return
}

// Relation is a wrapper around the C record AtkRelation.
type Relation struct {
	native *C.AtkRelation
	// parent : record
	// no type for target
	Relationship RelationType
}

func RelationNewFromC(u unsafe.Pointer) *Relation {
	c := (*C.AtkRelation)(u)
	if c == nil {
		return nil
	}

	g := &Relation{
		Relationship: (RelationType)(c.relationship),
		native:       c,
	}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Relation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Relation) ToC() unsafe.Pointer {
	recv.native.relationship =
		(C.AtkRelationType)(recv.Relationship)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Relation with another Relation, and returns true if they represent the same GObject.
func (recv *Relation) Equals(other *Relation) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Relation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Relation.
// Exercise care, as this is a potentially dangerous function if the Object is not a Relation.
func CastToRelation(object *gobject.Object) *Relation {
	return RelationNewFromC(object.ToC())
}

// Unsupported : atk_relation_new : unsupported parameter targets :

// GetRelationType is a wrapper around the C function atk_relation_get_relation_type.
func (recv *Relation) GetRelationType() RelationType {
	retC := C.atk_relation_get_relation_type((*C.AtkRelation)(recv.native))
	retGo := (RelationType)(retC)

	return retGo
}

// Unsupported : atk_relation_get_target : array return type :

// RemoveTarget is a wrapper around the C function atk_relation_remove_target.
func (recv *Relation) RemoveTarget(target *Object) bool {
	c_target := (*C.AtkObject)(C.NULL)
	if target != nil {
		c_target = (*C.AtkObject)(target.ToC())
	}

	retC := C.atk_relation_remove_target((*C.AtkRelation)(recv.native), c_target)
	retGo := retC == C.TRUE

	return retGo
}

// RelationSet is a wrapper around the C record AtkRelationSet.
type RelationSet struct {
	native *C.AtkRelationSet
	// parent : record
	// no type for relations
}

func RelationSetNewFromC(u unsafe.Pointer) *RelationSet {
	c := (*C.AtkRelationSet)(u)
	if c == nil {
		return nil
	}

	g := &RelationSet{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RelationSet) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RelationSet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RelationSet with another RelationSet, and returns true if they represent the same GObject.
func (recv *RelationSet) Equals(other *RelationSet) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *RelationSet) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to RelationSet.
// Exercise care, as this is a potentially dangerous function if the Object is not a RelationSet.
func CastToRelationSet(object *gobject.Object) *RelationSet {
	return RelationSetNewFromC(object.ToC())
}

// RelationSetNew is a wrapper around the C function atk_relation_set_new.
func RelationSetNew() *RelationSet {
	retC := C.atk_relation_set_new()
	retGo := RelationSetNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Add is a wrapper around the C function atk_relation_set_add.
func (recv *RelationSet) Add(relation *Relation) {
	c_relation := (*C.AtkRelation)(C.NULL)
	if relation != nil {
		c_relation = (*C.AtkRelation)(relation.ToC())
	}

	C.atk_relation_set_add((*C.AtkRelationSet)(recv.native), c_relation)

	return
}

// Contains is a wrapper around the C function atk_relation_set_contains.
func (recv *RelationSet) Contains(relationship RelationType) bool {
	c_relationship := (C.AtkRelationType)(relationship)

	retC := C.atk_relation_set_contains((*C.AtkRelationSet)(recv.native), c_relationship)
	retGo := retC == C.TRUE

	return retGo
}

// ContainsTarget is a wrapper around the C function atk_relation_set_contains_target.
func (recv *RelationSet) ContainsTarget(relationship RelationType, target *Object) bool {
	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(C.NULL)
	if target != nil {
		c_target = (*C.AtkObject)(target.ToC())
	}

	retC := C.atk_relation_set_contains_target((*C.AtkRelationSet)(recv.native), c_relationship, c_target)
	retGo := retC == C.TRUE

	return retGo
}

// GetNRelations is a wrapper around the C function atk_relation_set_get_n_relations.
func (recv *RelationSet) GetNRelations() int32 {
	retC := C.atk_relation_set_get_n_relations((*C.AtkRelationSet)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRelation is a wrapper around the C function atk_relation_set_get_relation.
func (recv *RelationSet) GetRelation(i int32) *Relation {
	c_i := (C.gint)(i)

	retC := C.atk_relation_set_get_relation((*C.AtkRelationSet)(recv.native), c_i)
	retGo := RelationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRelationByType is a wrapper around the C function atk_relation_set_get_relation_by_type.
func (recv *RelationSet) GetRelationByType(relationship RelationType) *Relation {
	c_relationship := (C.AtkRelationType)(relationship)

	retC := C.atk_relation_set_get_relation_by_type((*C.AtkRelationSet)(recv.native), c_relationship)
	retGo := RelationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Remove is a wrapper around the C function atk_relation_set_remove.
func (recv *RelationSet) Remove(relation *Relation) {
	c_relation := (*C.AtkRelation)(C.NULL)
	if relation != nil {
		c_relation = (*C.AtkRelation)(relation.ToC())
	}

	C.atk_relation_set_remove((*C.AtkRelationSet)(recv.native), c_relation)

	return
}

// Socket is a wrapper around the C record AtkSocket.
type Socket struct {
	native *C.AtkSocket
	// parent : record
	// Private : embedded_plug_id
}

func SocketNewFromC(u unsafe.Pointer) *Socket {
	c := (*C.AtkSocket)(u)
	if c == nil {
		return nil
	}

	g := &Socket{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Socket) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Socket) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Socket with another Socket, and returns true if they represent the same GObject.
func (recv *Socket) Equals(other *Socket) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Socket) Object() *Object {
	return ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Socket.
// Exercise care, as this is a potentially dangerous function if the Object is not a Socket.
func CastToSocket(object *gobject.Object) *Socket {
	return SocketNewFromC(object.ToC())
}

// SocketNew is a wrapper around the C function atk_socket_new.
func SocketNew() *Socket {
	retC := C.atk_socket_new()
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Component returns the Component interface implemented by Socket
func (recv *Socket) Component() *Component {
	return ComponentNewFromC(recv.ToC())
}

// StateSet is a wrapper around the C record AtkStateSet.
type StateSet struct {
	native *C.AtkStateSet
	// parent : record
}

func StateSetNewFromC(u unsafe.Pointer) *StateSet {
	c := (*C.AtkStateSet)(u)
	if c == nil {
		return nil
	}

	g := &StateSet{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StateSet) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StateSet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StateSet with another StateSet, and returns true if they represent the same GObject.
func (recv *StateSet) Equals(other *StateSet) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *StateSet) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to StateSet.
// Exercise care, as this is a potentially dangerous function if the Object is not a StateSet.
func CastToStateSet(object *gobject.Object) *StateSet {
	return StateSetNewFromC(object.ToC())
}

// StateSetNew is a wrapper around the C function atk_state_set_new.
func StateSetNew() *StateSet {
	retC := C.atk_state_set_new()
	retGo := StateSetNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddState is a wrapper around the C function atk_state_set_add_state.
func (recv *StateSet) AddState(type_ StateType) bool {
	c_type := (C.AtkStateType)(type_)

	retC := C.atk_state_set_add_state((*C.AtkStateSet)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// AddStates is a wrapper around the C function atk_state_set_add_states.
func (recv *StateSet) AddStates(types []StateType) {
	c_types_array := make([]C.AtkStateType, len(types)+1, len(types)+1)
	for i, item := range types {
		c := (C.AtkStateType)(item)
		c_types_array[i] = c
	}
	c_types_array[len(types)] = 0
	c_types_arrayPtr := &c_types_array[0]
	c_types := (*C.AtkStateType)(unsafe.Pointer(c_types_arrayPtr))

	c_n_types := (C.gint)(len(types))

	C.atk_state_set_add_states((*C.AtkStateSet)(recv.native), c_types, c_n_types)

	return
}

// AndSets is a wrapper around the C function atk_state_set_and_sets.
func (recv *StateSet) AndSets(compareSet *StateSet) *StateSet {
	c_compare_set := (*C.AtkStateSet)(C.NULL)
	if compareSet != nil {
		c_compare_set = (*C.AtkStateSet)(compareSet.ToC())
	}

	retC := C.atk_state_set_and_sets((*C.AtkStateSet)(recv.native), c_compare_set)
	retGo := StateSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ClearStates is a wrapper around the C function atk_state_set_clear_states.
func (recv *StateSet) ClearStates() {
	C.atk_state_set_clear_states((*C.AtkStateSet)(recv.native))

	return
}

// ContainsState is a wrapper around the C function atk_state_set_contains_state.
func (recv *StateSet) ContainsState(type_ StateType) bool {
	c_type := (C.AtkStateType)(type_)

	retC := C.atk_state_set_contains_state((*C.AtkStateSet)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// ContainsStates is a wrapper around the C function atk_state_set_contains_states.
func (recv *StateSet) ContainsStates(types []StateType) bool {
	c_types_array := make([]C.AtkStateType, len(types)+1, len(types)+1)
	for i, item := range types {
		c := (C.AtkStateType)(item)
		c_types_array[i] = c
	}
	c_types_array[len(types)] = 0
	c_types_arrayPtr := &c_types_array[0]
	c_types := (*C.AtkStateType)(unsafe.Pointer(c_types_arrayPtr))

	c_n_types := (C.gint)(len(types))

	retC := C.atk_state_set_contains_states((*C.AtkStateSet)(recv.native), c_types, c_n_types)
	retGo := retC == C.TRUE

	return retGo
}

// IsEmpty is a wrapper around the C function atk_state_set_is_empty.
func (recv *StateSet) IsEmpty() bool {
	retC := C.atk_state_set_is_empty((*C.AtkStateSet)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// OrSets is a wrapper around the C function atk_state_set_or_sets.
func (recv *StateSet) OrSets(compareSet *StateSet) *StateSet {
	c_compare_set := (*C.AtkStateSet)(C.NULL)
	if compareSet != nil {
		c_compare_set = (*C.AtkStateSet)(compareSet.ToC())
	}

	retC := C.atk_state_set_or_sets((*C.AtkStateSet)(recv.native), c_compare_set)
	var retGo (*StateSet)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StateSetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// RemoveState is a wrapper around the C function atk_state_set_remove_state.
func (recv *StateSet) RemoveState(type_ StateType) bool {
	c_type := (C.AtkStateType)(type_)

	retC := C.atk_state_set_remove_state((*C.AtkStateSet)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// XorSets is a wrapper around the C function atk_state_set_xor_sets.
func (recv *StateSet) XorSets(compareSet *StateSet) *StateSet {
	c_compare_set := (*C.AtkStateSet)(C.NULL)
	if compareSet != nil {
		c_compare_set = (*C.AtkStateSet)(compareSet.ToC())
	}

	retC := C.atk_state_set_xor_sets((*C.AtkStateSet)(recv.native), c_compare_set)
	retGo := StateSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Util is a wrapper around the C record AtkUtil.
type Util struct {
	native *C.AtkUtil
	// parent : record
}

func UtilNewFromC(u unsafe.Pointer) *Util {
	c := (*C.AtkUtil)(u)
	if c == nil {
		return nil
	}

	g := &Util{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Util) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Util) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Util with another Util, and returns true if they represent the same GObject.
func (recv *Util) Equals(other *Util) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Util) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Util.
// Exercise care, as this is a potentially dangerous function if the Object is not a Util.
func CastToUtil(object *gobject.Object) *Util {
	return UtilNewFromC(object.ToC())
}

type CoordType C.AtkCoordType

const (
	ATK_XY_SCREEN CoordType = 0
	ATK_XY_WINDOW CoordType = 1
)

type KeyEventType C.AtkKeyEventType

const (
	ATK_KEY_EVENT_PRESS        KeyEventType = 0
	ATK_KEY_EVENT_RELEASE      KeyEventType = 1
	ATK_KEY_EVENT_LAST_DEFINED KeyEventType = 2
)

type Layer C.AtkLayer

const (
	ATK_LAYER_INVALID    Layer = 0
	ATK_LAYER_BACKGROUND Layer = 1
	ATK_LAYER_CANVAS     Layer = 2
	ATK_LAYER_WIDGET     Layer = 3
	ATK_LAYER_MDI        Layer = 4
	ATK_LAYER_POPUP      Layer = 5
	ATK_LAYER_OVERLAY    Layer = 6
	ATK_LAYER_WINDOW     Layer = 7
)

type RelationType C.AtkRelationType

const (
	ATK_RELATION_NULL             RelationType = 0
	ATK_RELATION_CONTROLLED_BY    RelationType = 1
	ATK_RELATION_CONTROLLER_FOR   RelationType = 2
	ATK_RELATION_LABEL_FOR        RelationType = 3
	ATK_RELATION_LABELLED_BY      RelationType = 4
	ATK_RELATION_MEMBER_OF        RelationType = 5
	ATK_RELATION_NODE_CHILD_OF    RelationType = 6
	ATK_RELATION_FLOWS_TO         RelationType = 7
	ATK_RELATION_FLOWS_FROM       RelationType = 8
	ATK_RELATION_SUBWINDOW_OF     RelationType = 9
	ATK_RELATION_EMBEDS           RelationType = 10
	ATK_RELATION_EMBEDDED_BY      RelationType = 11
	ATK_RELATION_POPUP_FOR        RelationType = 12
	ATK_RELATION_PARENT_WINDOW_OF RelationType = 13
	ATK_RELATION_DESCRIBED_BY     RelationType = 14
	ATK_RELATION_DESCRIPTION_FOR  RelationType = 15
	ATK_RELATION_NODE_PARENT_OF   RelationType = 16
	ATK_RELATION_DETAILS          RelationType = 17
	ATK_RELATION_DETAILS_FOR      RelationType = 18
	ATK_RELATION_ERROR_MESSAGE    RelationType = 19
	ATK_RELATION_ERROR_FOR        RelationType = 20
	ATK_RELATION_LAST_DEFINED     RelationType = 21
)

// RelationTypeForName is a wrapper around the C function atk_relation_type_for_name.
func RelationTypeForName(name string) RelationType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_relation_type_for_name(c_name)
	retGo := (RelationType)(retC)

	return retGo
}

// RelationTypeGetName is a wrapper around the C function atk_relation_type_get_name.
func RelationTypeGetName(type_ RelationType) string {
	c_type := (C.AtkRelationType)(type_)

	retC := C.atk_relation_type_get_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// RelationTypeRegister is a wrapper around the C function atk_relation_type_register.
func RelationTypeRegister(name string) RelationType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_relation_type_register(c_name)
	retGo := (RelationType)(retC)

	return retGo
}

type Role C.AtkRole

const (
	ATK_ROLE_INVALID               Role = 0
	ATK_ROLE_ACCEL_LABEL           Role = 1
	ATK_ROLE_ALERT                 Role = 2
	ATK_ROLE_ANIMATION             Role = 3
	ATK_ROLE_ARROW                 Role = 4
	ATK_ROLE_CALENDAR              Role = 5
	ATK_ROLE_CANVAS                Role = 6
	ATK_ROLE_CHECK_BOX             Role = 7
	ATK_ROLE_CHECK_MENU_ITEM       Role = 8
	ATK_ROLE_COLOR_CHOOSER         Role = 9
	ATK_ROLE_COLUMN_HEADER         Role = 10
	ATK_ROLE_COMBO_BOX             Role = 11
	ATK_ROLE_DATE_EDITOR           Role = 12
	ATK_ROLE_DESKTOP_ICON          Role = 13
	ATK_ROLE_DESKTOP_FRAME         Role = 14
	ATK_ROLE_DIAL                  Role = 15
	ATK_ROLE_DIALOG                Role = 16
	ATK_ROLE_DIRECTORY_PANE        Role = 17
	ATK_ROLE_DRAWING_AREA          Role = 18
	ATK_ROLE_FILE_CHOOSER          Role = 19
	ATK_ROLE_FILLER                Role = 20
	ATK_ROLE_FONT_CHOOSER          Role = 21
	ATK_ROLE_FRAME                 Role = 22
	ATK_ROLE_GLASS_PANE            Role = 23
	ATK_ROLE_HTML_CONTAINER        Role = 24
	ATK_ROLE_ICON                  Role = 25
	ATK_ROLE_IMAGE                 Role = 26
	ATK_ROLE_INTERNAL_FRAME        Role = 27
	ATK_ROLE_LABEL                 Role = 28
	ATK_ROLE_LAYERED_PANE          Role = 29
	ATK_ROLE_LIST                  Role = 30
	ATK_ROLE_LIST_ITEM             Role = 31
	ATK_ROLE_MENU                  Role = 32
	ATK_ROLE_MENU_BAR              Role = 33
	ATK_ROLE_MENU_ITEM             Role = 34
	ATK_ROLE_OPTION_PANE           Role = 35
	ATK_ROLE_PAGE_TAB              Role = 36
	ATK_ROLE_PAGE_TAB_LIST         Role = 37
	ATK_ROLE_PANEL                 Role = 38
	ATK_ROLE_PASSWORD_TEXT         Role = 39
	ATK_ROLE_POPUP_MENU            Role = 40
	ATK_ROLE_PROGRESS_BAR          Role = 41
	ATK_ROLE_PUSH_BUTTON           Role = 42
	ATK_ROLE_RADIO_BUTTON          Role = 43
	ATK_ROLE_RADIO_MENU_ITEM       Role = 44
	ATK_ROLE_ROOT_PANE             Role = 45
	ATK_ROLE_ROW_HEADER            Role = 46
	ATK_ROLE_SCROLL_BAR            Role = 47
	ATK_ROLE_SCROLL_PANE           Role = 48
	ATK_ROLE_SEPARATOR             Role = 49
	ATK_ROLE_SLIDER                Role = 50
	ATK_ROLE_SPLIT_PANE            Role = 51
	ATK_ROLE_SPIN_BUTTON           Role = 52
	ATK_ROLE_STATUSBAR             Role = 53
	ATK_ROLE_TABLE                 Role = 54
	ATK_ROLE_TABLE_CELL            Role = 55
	ATK_ROLE_TABLE_COLUMN_HEADER   Role = 56
	ATK_ROLE_TABLE_ROW_HEADER      Role = 57
	ATK_ROLE_TEAR_OFF_MENU_ITEM    Role = 58
	ATK_ROLE_TERMINAL              Role = 59
	ATK_ROLE_TEXT                  Role = 60
	ATK_ROLE_TOGGLE_BUTTON         Role = 61
	ATK_ROLE_TOOL_BAR              Role = 62
	ATK_ROLE_TOOL_TIP              Role = 63
	ATK_ROLE_TREE                  Role = 64
	ATK_ROLE_TREE_TABLE            Role = 65
	ATK_ROLE_UNKNOWN               Role = 66
	ATK_ROLE_VIEWPORT              Role = 67
	ATK_ROLE_WINDOW                Role = 68
	ATK_ROLE_HEADER                Role = 69
	ATK_ROLE_FOOTER                Role = 70
	ATK_ROLE_PARAGRAPH             Role = 71
	ATK_ROLE_RULER                 Role = 72
	ATK_ROLE_APPLICATION           Role = 73
	ATK_ROLE_AUTOCOMPLETE          Role = 74
	ATK_ROLE_EDITBAR               Role = 75
	ATK_ROLE_EMBEDDED              Role = 76
	ATK_ROLE_ENTRY                 Role = 77
	ATK_ROLE_CHART                 Role = 78
	ATK_ROLE_CAPTION               Role = 79
	ATK_ROLE_DOCUMENT_FRAME        Role = 80
	ATK_ROLE_HEADING               Role = 81
	ATK_ROLE_PAGE                  Role = 82
	ATK_ROLE_SECTION               Role = 83
	ATK_ROLE_REDUNDANT_OBJECT      Role = 84
	ATK_ROLE_FORM                  Role = 85
	ATK_ROLE_LINK                  Role = 86
	ATK_ROLE_INPUT_METHOD_WINDOW   Role = 87
	ATK_ROLE_TABLE_ROW             Role = 88
	ATK_ROLE_TREE_ITEM             Role = 89
	ATK_ROLE_DOCUMENT_SPREADSHEET  Role = 90
	ATK_ROLE_DOCUMENT_PRESENTATION Role = 91
	ATK_ROLE_DOCUMENT_TEXT         Role = 92
	ATK_ROLE_DOCUMENT_WEB          Role = 93
	ATK_ROLE_DOCUMENT_EMAIL        Role = 94
	ATK_ROLE_COMMENT               Role = 95
	ATK_ROLE_LIST_BOX              Role = 96
	ATK_ROLE_GROUPING              Role = 97
	ATK_ROLE_IMAGE_MAP             Role = 98
	ATK_ROLE_NOTIFICATION          Role = 99
	ATK_ROLE_INFO_BAR              Role = 100
	ATK_ROLE_LEVEL_BAR             Role = 101
	ATK_ROLE_TITLE_BAR             Role = 102
	ATK_ROLE_BLOCK_QUOTE           Role = 103
	ATK_ROLE_AUDIO                 Role = 104
	ATK_ROLE_VIDEO                 Role = 105
	ATK_ROLE_DEFINITION            Role = 106
	ATK_ROLE_ARTICLE               Role = 107
	ATK_ROLE_LANDMARK              Role = 108
	ATK_ROLE_LOG                   Role = 109
	ATK_ROLE_MARQUEE               Role = 110
	ATK_ROLE_MATH                  Role = 111
	ATK_ROLE_RATING                Role = 112
	ATK_ROLE_TIMER                 Role = 113
	ATK_ROLE_DESCRIPTION_LIST      Role = 114
	ATK_ROLE_DESCRIPTION_TERM      Role = 115
	ATK_ROLE_DESCRIPTION_VALUE     Role = 116
	ATK_ROLE_STATIC                Role = 117
	ATK_ROLE_MATH_FRACTION         Role = 118
	ATK_ROLE_MATH_ROOT             Role = 119
	ATK_ROLE_SUBSCRIPT             Role = 120
	ATK_ROLE_SUPERSCRIPT           Role = 121
	ATK_ROLE_FOOTNOTE              Role = 122
	ATK_ROLE_LAST_DEFINED          Role = 123
)

// RoleForName is a wrapper around the C function atk_role_for_name.
func RoleForName(name string) Role {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_role_for_name(c_name)
	retGo := (Role)(retC)

	return retGo
}

// RoleGetLocalizedName is a wrapper around the C function atk_role_get_localized_name.
func RoleGetLocalizedName(role Role) string {
	c_role := (C.AtkRole)(role)

	retC := C.atk_role_get_localized_name(c_role)
	retGo := C.GoString(retC)

	return retGo
}

// RoleGetName is a wrapper around the C function atk_role_get_name.
func RoleGetName(role Role) string {
	c_role := (C.AtkRole)(role)

	retC := C.atk_role_get_name(c_role)
	retGo := C.GoString(retC)

	return retGo
}

// RoleRegister is a wrapper around the C function atk_role_register.
func RoleRegister(name string) Role {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_role_register(c_name)
	retGo := (Role)(retC)

	return retGo
}

type StateType C.AtkStateType

const (
	ATK_STATE_INVALID                 StateType = 0
	ATK_STATE_ACTIVE                  StateType = 1
	ATK_STATE_ARMED                   StateType = 2
	ATK_STATE_BUSY                    StateType = 3
	ATK_STATE_CHECKED                 StateType = 4
	ATK_STATE_DEFUNCT                 StateType = 5
	ATK_STATE_EDITABLE                StateType = 6
	ATK_STATE_ENABLED                 StateType = 7
	ATK_STATE_EXPANDABLE              StateType = 8
	ATK_STATE_EXPANDED                StateType = 9
	ATK_STATE_FOCUSABLE               StateType = 10
	ATK_STATE_FOCUSED                 StateType = 11
	ATK_STATE_HORIZONTAL              StateType = 12
	ATK_STATE_ICONIFIED               StateType = 13
	ATK_STATE_MODAL                   StateType = 14
	ATK_STATE_MULTI_LINE              StateType = 15
	ATK_STATE_MULTISELECTABLE         StateType = 16
	ATK_STATE_OPAQUE                  StateType = 17
	ATK_STATE_PRESSED                 StateType = 18
	ATK_STATE_RESIZABLE               StateType = 19
	ATK_STATE_SELECTABLE              StateType = 20
	ATK_STATE_SELECTED                StateType = 21
	ATK_STATE_SENSITIVE               StateType = 22
	ATK_STATE_SHOWING                 StateType = 23
	ATK_STATE_SINGLE_LINE             StateType = 24
	ATK_STATE_STALE                   StateType = 25
	ATK_STATE_TRANSIENT               StateType = 26
	ATK_STATE_VERTICAL                StateType = 27
	ATK_STATE_VISIBLE                 StateType = 28
	ATK_STATE_MANAGES_DESCENDANTS     StateType = 29
	ATK_STATE_INDETERMINATE           StateType = 30
	ATK_STATE_TRUNCATED               StateType = 31
	ATK_STATE_REQUIRED                StateType = 32
	ATK_STATE_INVALID_ENTRY           StateType = 33
	ATK_STATE_SUPPORTS_AUTOCOMPLETION StateType = 34
	ATK_STATE_SELECTABLE_TEXT         StateType = 35
	ATK_STATE_DEFAULT                 StateType = 36
	ATK_STATE_ANIMATED                StateType = 37
	ATK_STATE_VISITED                 StateType = 38
	ATK_STATE_CHECKABLE               StateType = 39
	ATK_STATE_HAS_POPUP               StateType = 40
	ATK_STATE_HAS_TOOLTIP             StateType = 41
	ATK_STATE_READ_ONLY               StateType = 42
	ATK_STATE_LAST_DEFINED            StateType = 43
)

// StateTypeForName is a wrapper around the C function atk_state_type_for_name.
func StateTypeForName(name string) StateType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_state_type_for_name(c_name)
	retGo := (StateType)(retC)

	return retGo
}

// StateTypeGetName is a wrapper around the C function atk_state_type_get_name.
func StateTypeGetName(type_ StateType) string {
	c_type := (C.AtkStateType)(type_)

	retC := C.atk_state_type_get_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// StateTypeRegister is a wrapper around the C function atk_state_type_register.
func StateTypeRegister(name string) StateType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_state_type_register(c_name)
	retGo := (StateType)(retC)

	return retGo
}

type TextAttribute C.AtkTextAttribute

const (
	ATK_TEXT_ATTR_INVALID            TextAttribute = 0
	ATK_TEXT_ATTR_LEFT_MARGIN        TextAttribute = 1
	ATK_TEXT_ATTR_RIGHT_MARGIN       TextAttribute = 2
	ATK_TEXT_ATTR_INDENT             TextAttribute = 3
	ATK_TEXT_ATTR_INVISIBLE          TextAttribute = 4
	ATK_TEXT_ATTR_EDITABLE           TextAttribute = 5
	ATK_TEXT_ATTR_PIXELS_ABOVE_LINES TextAttribute = 6
	ATK_TEXT_ATTR_PIXELS_BELOW_LINES TextAttribute = 7
	ATK_TEXT_ATTR_PIXELS_INSIDE_WRAP TextAttribute = 8
	ATK_TEXT_ATTR_BG_FULL_HEIGHT     TextAttribute = 9
	ATK_TEXT_ATTR_RISE               TextAttribute = 10
	ATK_TEXT_ATTR_UNDERLINE          TextAttribute = 11
	ATK_TEXT_ATTR_STRIKETHROUGH      TextAttribute = 12
	ATK_TEXT_ATTR_SIZE               TextAttribute = 13
	ATK_TEXT_ATTR_SCALE              TextAttribute = 14
	ATK_TEXT_ATTR_WEIGHT             TextAttribute = 15
	ATK_TEXT_ATTR_LANGUAGE           TextAttribute = 16
	ATK_TEXT_ATTR_FAMILY_NAME        TextAttribute = 17
	ATK_TEXT_ATTR_BG_COLOR           TextAttribute = 18
	ATK_TEXT_ATTR_FG_COLOR           TextAttribute = 19
	ATK_TEXT_ATTR_BG_STIPPLE         TextAttribute = 20
	ATK_TEXT_ATTR_FG_STIPPLE         TextAttribute = 21
	ATK_TEXT_ATTR_WRAP_MODE          TextAttribute = 22
	ATK_TEXT_ATTR_DIRECTION          TextAttribute = 23
	ATK_TEXT_ATTR_JUSTIFICATION      TextAttribute = 24
	ATK_TEXT_ATTR_STRETCH            TextAttribute = 25
	ATK_TEXT_ATTR_VARIANT            TextAttribute = 26
	ATK_TEXT_ATTR_STYLE              TextAttribute = 27
	ATK_TEXT_ATTR_LAST_DEFINED       TextAttribute = 28
)

// TextAttributeForName is a wrapper around the C function atk_text_attribute_for_name.
func TextAttributeForName(name string) TextAttribute {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_text_attribute_for_name(c_name)
	retGo := (TextAttribute)(retC)

	return retGo
}

// TextAttributeGetName is a wrapper around the C function atk_text_attribute_get_name.
func TextAttributeGetName(attr TextAttribute) string {
	c_attr := (C.AtkTextAttribute)(attr)

	retC := C.atk_text_attribute_get_name(c_attr)
	retGo := C.GoString(retC)

	return retGo
}

// TextAttributeGetValue is a wrapper around the C function atk_text_attribute_get_value.
func TextAttributeGetValue(attr TextAttribute, index int32) string {
	c_attr := (C.AtkTextAttribute)(attr)

	c_index_ := (C.gint)(index)

	retC := C.atk_text_attribute_get_value(c_attr, c_index_)
	retGo := C.GoString(retC)

	return retGo
}

// TextAttributeRegister is a wrapper around the C function atk_text_attribute_register.
func TextAttributeRegister(name string) TextAttribute {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.atk_text_attribute_register(c_name)
	retGo := (TextAttribute)(retC)

	return retGo
}

type TextBoundary C.AtkTextBoundary

const (
	ATK_TEXT_BOUNDARY_CHAR           TextBoundary = 0
	ATK_TEXT_BOUNDARY_WORD_START     TextBoundary = 1
	ATK_TEXT_BOUNDARY_WORD_END       TextBoundary = 2
	ATK_TEXT_BOUNDARY_SENTENCE_START TextBoundary = 3
	ATK_TEXT_BOUNDARY_SENTENCE_END   TextBoundary = 4
	ATK_TEXT_BOUNDARY_LINE_START     TextBoundary = 5
	ATK_TEXT_BOUNDARY_LINE_END       TextBoundary = 6
)

type TextClipType C.AtkTextClipType

const (
	ATK_TEXT_CLIP_NONE TextClipType = 0
	ATK_TEXT_CLIP_MIN  TextClipType = 1
	ATK_TEXT_CLIP_MAX  TextClipType = 2
	ATK_TEXT_CLIP_BOTH TextClipType = 3
)

type TextGranularity C.AtkTextGranularity

const (
	ATK_TEXT_GRANULARITY_CHAR      TextGranularity = 0
	ATK_TEXT_GRANULARITY_WORD      TextGranularity = 1
	ATK_TEXT_GRANULARITY_SENTENCE  TextGranularity = 2
	ATK_TEXT_GRANULARITY_LINE      TextGranularity = 3
	ATK_TEXT_GRANULARITY_PARAGRAPH TextGranularity = 4
)

type ValueType C.AtkValueType

const (
	ATK_VALUE_VERY_WEAK    ValueType = 0
	ATK_VALUE_WEAK         ValueType = 1
	ATK_VALUE_ACCEPTABLE   ValueType = 2
	ATK_VALUE_STRONG       ValueType = 3
	ATK_VALUE_VERY_STRONG  ValueType = 4
	ATK_VALUE_VERY_LOW     ValueType = 5
	ATK_VALUE_LOW          ValueType = 6
	ATK_VALUE_MEDIUM       ValueType = 7
	ATK_VALUE_HIGH         ValueType = 8
	ATK_VALUE_VERY_HIGH    ValueType = 9
	ATK_VALUE_VERY_BAD     ValueType = 10
	ATK_VALUE_BAD          ValueType = 11
	ATK_VALUE_GOOD         ValueType = 12
	ATK_VALUE_VERY_GOOD    ValueType = 13
	ATK_VALUE_BEST         ValueType = 14
	ATK_VALUE_LAST_DEFINED ValueType = 15
)

// ValueTypeGetLocalizedName is a wrapper around the C function atk_value_type_get_localized_name.
func ValueTypeGetLocalizedName(valueType ValueType) string {
	c_value_type := (C.AtkValueType)(valueType)

	retC := C.atk_value_type_get_localized_name(c_value_type)
	retGo := C.GoString(retC)

	return retGo
}

// ValueTypeGetName is a wrapper around the C function atk_value_type_get_name.
func ValueTypeGetName(valueType ValueType) string {
	c_value_type := (C.AtkValueType)(valueType)

	retC := C.atk_value_type_get_name(c_value_type)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : atk_add_focus_tracker : unsupported parameter focus_tracker : no type generator for EventListener (AtkEventListener) for param focus_tracker

// Unsupported : atk_add_global_event_listener : unsupported parameter listener : no type generator for GObject.SignalEmissionHook (GSignalEmissionHook) for param listener

// Unsupported : atk_add_key_event_listener : unsupported parameter listener : no type generator for KeySnoopFunc (AtkKeySnoopFunc) for param listener

// Blacklisted : atk_attribute_set_free

// Unsupported : atk_focus_tracker_init : unsupported parameter init : no type generator for EventListenerInit (AtkEventListenerInit) for param init

// FocusTrackerNotify is a wrapper around the C function atk_focus_tracker_notify.
func FocusTrackerNotify(object *Object) {
	c_object := (*C.AtkObject)(C.NULL)
	if object != nil {
		c_object = (*C.AtkObject)(object.ToC())
	}

	C.atk_focus_tracker_notify(c_object)

	return
}

// GetDefaultRegistry is a wrapper around the C function atk_get_default_registry.
func GetDefaultRegistry() *Registry {
	retC := C.atk_get_default_registry()
	retGo := RegistryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRoot is a wrapper around the C function atk_get_root.
func GetRoot() *Object {
	retC := C.atk_get_root()
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetToolkitName is a wrapper around the C function atk_get_toolkit_name.
func GetToolkitName() string {
	retC := C.atk_get_toolkit_name()
	retGo := C.GoString(retC)

	return retGo
}

// GetToolkitVersion is a wrapper around the C function atk_get_toolkit_version.
func GetToolkitVersion() string {
	retC := C.atk_get_toolkit_version()
	retGo := C.GoString(retC)

	return retGo
}

// RemoveFocusTracker is a wrapper around the C function atk_remove_focus_tracker.
func RemoveFocusTracker(trackerId uint32) {
	c_tracker_id := (C.guint)(trackerId)

	C.atk_remove_focus_tracker(c_tracker_id)

	return
}

// RemoveGlobalEventListener is a wrapper around the C function atk_remove_global_event_listener.
func RemoveGlobalEventListener(listenerId uint32) {
	c_listener_id := (C.guint)(listenerId)

	C.atk_remove_global_event_listener(c_listener_id)

	return
}

// RemoveKeyEventListener is a wrapper around the C function atk_remove_key_event_listener.
func RemoveKeyEventListener(listenerId uint32) {
	c_listener_id := (C.guint)(listenerId)

	C.atk_remove_key_event_listener(c_listener_id)

	return
}

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

// DoAction is a wrapper around the C function atk_action_do_action.
func (recv *Action) DoAction(i int32) bool {
	c_i := (C.gint)(i)

	retC := C.atk_action_do_action((*C.AtkAction)(recv.native), c_i)
	retGo := retC == C.TRUE

	return retGo
}

// GetDescription is a wrapper around the C function atk_action_get_description.
func (recv *Action) GetDescription(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_action_get_description((*C.AtkAction)(recv.native), c_i)
	retGo := C.GoString(retC)

	return retGo
}

// GetKeybinding is a wrapper around the C function atk_action_get_keybinding.
func (recv *Action) GetKeybinding(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_action_get_keybinding((*C.AtkAction)(recv.native), c_i)
	retGo := C.GoString(retC)

	return retGo
}

// GetLocalizedName is a wrapper around the C function atk_action_get_localized_name.
func (recv *Action) GetLocalizedName(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_action_get_localized_name((*C.AtkAction)(recv.native), c_i)
	retGo := C.GoString(retC)

	return retGo
}

// GetNActions is a wrapper around the C function atk_action_get_n_actions.
func (recv *Action) GetNActions() int32 {
	retC := C.atk_action_get_n_actions((*C.AtkAction)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetName is a wrapper around the C function atk_action_get_name.
func (recv *Action) GetName(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_action_get_name((*C.AtkAction)(recv.native), c_i)
	retGo := C.GoString(retC)

	return retGo
}

// SetDescription is a wrapper around the C function atk_action_set_description.
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
type ComponentSignalBoundsChangedCallback func(targetObject *Component, arg1 *Rectangle)

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
func component_boundsChangedHandler(c_targetObject *C.GObject, c_arg1 *C.AtkRectangle, data C.gpointer) {
	signalComponentBoundsChangedLock.RLock()
	defer signalComponentBoundsChangedLock.RUnlock()

	arg1 := RectangleNewFromC(unsafe.Pointer(c_arg1))

	targetObject := ComponentNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalComponentBoundsChangedMap[index].callback
	callback(targetObject, arg1)
}

// Unsupported : atk_component_add_focus_handler : unsupported parameter handler : no type generator for FocusHandler (AtkFocusHandler) for param handler

// Contains is a wrapper around the C function atk_component_contains.
func (recv *Component) Contains(x int32, y int32, coordType CoordType) bool {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coord_type := (C.AtkCoordType)(coordType)

	retC := C.atk_component_contains((*C.AtkComponent)(recv.native), c_x, c_y, c_coord_type)
	retGo := retC == C.TRUE

	return retGo
}

// GetExtents is a wrapper around the C function atk_component_get_extents.
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

// GetLayer is a wrapper around the C function atk_component_get_layer.
func (recv *Component) GetLayer() Layer {
	retC := C.atk_component_get_layer((*C.AtkComponent)(recv.native))
	retGo := (Layer)(retC)

	return retGo
}

// GetMdiZorder is a wrapper around the C function atk_component_get_mdi_zorder.
func (recv *Component) GetMdiZorder() int32 {
	retC := C.atk_component_get_mdi_zorder((*C.AtkComponent)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPosition is a wrapper around the C function atk_component_get_position.
func (recv *Component) GetPosition(coordType CoordType) (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	c_coord_type := (C.AtkCoordType)(coordType)

	C.atk_component_get_position((*C.AtkComponent)(recv.native), &c_x, &c_y, c_coord_type)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// GetSize is a wrapper around the C function atk_component_get_size.
func (recv *Component) GetSize() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.atk_component_get_size((*C.AtkComponent)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// GrabFocus is a wrapper around the C function atk_component_grab_focus.
func (recv *Component) GrabFocus() bool {
	retC := C.atk_component_grab_focus((*C.AtkComponent)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// RefAccessibleAtPoint is a wrapper around the C function atk_component_ref_accessible_at_point.
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

// RemoveFocusHandler is a wrapper around the C function atk_component_remove_focus_handler.
func (recv *Component) RemoveFocusHandler(handlerId uint32) {
	c_handler_id := (C.guint)(handlerId)

	C.atk_component_remove_focus_handler((*C.AtkComponent)(recv.native), c_handler_id)

	return
}

// SetExtents is a wrapper around the C function atk_component_set_extents.
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

// SetPosition is a wrapper around the C function atk_component_set_position.
func (recv *Component) SetPosition(x int32, y int32, coordType CoordType) bool {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coord_type := (C.AtkCoordType)(coordType)

	retC := C.atk_component_set_position((*C.AtkComponent)(recv.native), c_x, c_y, c_coord_type)
	retGo := retC == C.TRUE

	return retGo
}

// SetSize is a wrapper around the C function atk_component_set_size.
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
type DocumentSignalLoadCompleteCallback func(targetObject *Document)

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
func document_loadCompleteHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalDocumentLoadCompleteLock.RLock()
	defer signalDocumentLoadCompleteLock.RUnlock()

	targetObject := DocumentNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalDocumentLoadCompleteMap[index].callback
	callback(targetObject)
}

type signalDocumentLoadStoppedDetail struct {
	callback  DocumentSignalLoadStoppedCallback
	handlerID C.gulong
}

var signalDocumentLoadStoppedId int
var signalDocumentLoadStoppedMap = make(map[int]signalDocumentLoadStoppedDetail)
var signalDocumentLoadStoppedLock sync.RWMutex

// DocumentSignalLoadStoppedCallback is a callback function for a 'load-stopped' signal emitted from a Document.
type DocumentSignalLoadStoppedCallback func(targetObject *Document)

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
func document_loadStoppedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalDocumentLoadStoppedLock.RLock()
	defer signalDocumentLoadStoppedLock.RUnlock()

	targetObject := DocumentNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalDocumentLoadStoppedMap[index].callback
	callback(targetObject)
}

type signalDocumentReloadDetail struct {
	callback  DocumentSignalReloadCallback
	handlerID C.gulong
}

var signalDocumentReloadId int
var signalDocumentReloadMap = make(map[int]signalDocumentReloadDetail)
var signalDocumentReloadLock sync.RWMutex

// DocumentSignalReloadCallback is a callback function for a 'reload' signal emitted from a Document.
type DocumentSignalReloadCallback func(targetObject *Document)

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
func document_reloadHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalDocumentReloadLock.RLock()
	defer signalDocumentReloadLock.RUnlock()

	targetObject := DocumentNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalDocumentReloadMap[index].callback
	callback(targetObject)
}

// GetDocument is a wrapper around the C function atk_document_get_document.
func (recv *Document) GetDocument() uintptr {
	retC := C.atk_document_get_document((*C.AtkDocument)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetDocumentType is a wrapper around the C function atk_document_get_document_type.
func (recv *Document) GetDocumentType() string {
	retC := C.atk_document_get_document_type((*C.AtkDocument)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLocale is a wrapper around the C function atk_document_get_locale.
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

// Equals compares this EditableText with another EditableText, and returns true if they represent the same GObject.
func (recv *EditableText) Equals(other *EditableText) bool {
	return other.ToC() == recv.ToC()
}

// CopyText is a wrapper around the C function atk_editable_text_copy_text.
func (recv *EditableText) CopyText(startPos int32, endPos int32) {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	C.atk_editable_text_copy_text((*C.AtkEditableText)(recv.native), c_start_pos, c_end_pos)

	return
}

// CutText is a wrapper around the C function atk_editable_text_cut_text.
func (recv *EditableText) CutText(startPos int32, endPos int32) {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	C.atk_editable_text_cut_text((*C.AtkEditableText)(recv.native), c_start_pos, c_end_pos)

	return
}

// DeleteText is a wrapper around the C function atk_editable_text_delete_text.
func (recv *EditableText) DeleteText(startPos int32, endPos int32) {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	C.atk_editable_text_delete_text((*C.AtkEditableText)(recv.native), c_start_pos, c_end_pos)

	return
}

// InsertText is a wrapper around the C function atk_editable_text_insert_text.
func (recv *EditableText) InsertText(string_ string, length int32, position int32) {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gint)(length)

	c_position := (C.gint)(position)

	C.atk_editable_text_insert_text((*C.AtkEditableText)(recv.native), c_string, c_length, &c_position)

	return
}

// PasteText is a wrapper around the C function atk_editable_text_paste_text.
func (recv *EditableText) PasteText(position int32) {
	c_position := (C.gint)(position)

	C.atk_editable_text_paste_text((*C.AtkEditableText)(recv.native), c_position)

	return
}

// Blacklisted : atk_editable_text_set_run_attributes

// SetTextContents is a wrapper around the C function atk_editable_text_set_text_contents.
func (recv *EditableText) SetTextContents(string_ string) {
	c_string := C.CString(string_)
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
type HypertextSignalLinkSelectedCallback func(targetObject *Hypertext, arg1 int32)

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
func hypertext_linkSelectedHandler(c_targetObject *C.GObject, c_arg1 C.gint, data C.gpointer) {
	signalHypertextLinkSelectedLock.RLock()
	defer signalHypertextLinkSelectedLock.RUnlock()

	arg1 := int32(c_arg1)

	targetObject := HypertextNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalHypertextLinkSelectedMap[index].callback
	callback(targetObject, arg1)
}

// GetLink is a wrapper around the C function atk_hypertext_get_link.
func (recv *Hypertext) GetLink(linkIndex int32) *Hyperlink {
	c_link_index := (C.gint)(linkIndex)

	retC := C.atk_hypertext_get_link((*C.AtkHypertext)(recv.native), c_link_index)
	retGo := HyperlinkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLinkIndex is a wrapper around the C function atk_hypertext_get_link_index.
func (recv *Hypertext) GetLinkIndex(charIndex int32) int32 {
	c_char_index := (C.gint)(charIndex)

	retC := C.atk_hypertext_get_link_index((*C.AtkHypertext)(recv.native), c_char_index)
	retGo := (int32)(retC)

	return retGo
}

// GetNLinks is a wrapper around the C function atk_hypertext_get_n_links.
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

// Equals compares this Image with another Image, and returns true if they represent the same GObject.
func (recv *Image) Equals(other *Image) bool {
	return other.ToC() == recv.ToC()
}

// GetImageDescription is a wrapper around the C function atk_image_get_image_description.
func (recv *Image) GetImageDescription() string {
	retC := C.atk_image_get_image_description((*C.AtkImage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetImagePosition is a wrapper around the C function atk_image_get_image_position.
func (recv *Image) GetImagePosition(coordType CoordType) (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	c_coord_type := (C.AtkCoordType)(coordType)

	C.atk_image_get_image_position((*C.AtkImage)(recv.native), &c_x, &c_y, c_coord_type)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// GetImageSize is a wrapper around the C function atk_image_get_image_size.
func (recv *Image) GetImageSize() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.atk_image_get_image_size((*C.AtkImage)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// SetImageDescription is a wrapper around the C function atk_image_set_image_description.
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
type SelectionSignalSelectionChangedCallback func(targetObject *Selection)

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
func selection_selectionChangedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalSelectionSelectionChangedLock.RLock()
	defer signalSelectionSelectionChangedLock.RUnlock()

	targetObject := SelectionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSelectionSelectionChangedMap[index].callback
	callback(targetObject)
}

// AddSelection is a wrapper around the C function atk_selection_add_selection.
func (recv *Selection) AddSelection(i int32) bool {
	c_i := (C.gint)(i)

	retC := C.atk_selection_add_selection((*C.AtkSelection)(recv.native), c_i)
	retGo := retC == C.TRUE

	return retGo
}

// ClearSelection is a wrapper around the C function atk_selection_clear_selection.
func (recv *Selection) ClearSelection() bool {
	retC := C.atk_selection_clear_selection((*C.AtkSelection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSelectionCount is a wrapper around the C function atk_selection_get_selection_count.
func (recv *Selection) GetSelectionCount() int32 {
	retC := C.atk_selection_get_selection_count((*C.AtkSelection)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// IsChildSelected is a wrapper around the C function atk_selection_is_child_selected.
func (recv *Selection) IsChildSelected(i int32) bool {
	c_i := (C.gint)(i)

	retC := C.atk_selection_is_child_selected((*C.AtkSelection)(recv.native), c_i)
	retGo := retC == C.TRUE

	return retGo
}

// RefSelection is a wrapper around the C function atk_selection_ref_selection.
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

// RemoveSelection is a wrapper around the C function atk_selection_remove_selection.
func (recv *Selection) RemoveSelection(i int32) bool {
	c_i := (C.gint)(i)

	retC := C.atk_selection_remove_selection((*C.AtkSelection)(recv.native), c_i)
	retGo := retC == C.TRUE

	return retGo
}

// SelectAllSelection is a wrapper around the C function atk_selection_select_all_selection.
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

// Equals compares this StreamableContent with another StreamableContent, and returns true if they represent the same GObject.
func (recv *StreamableContent) Equals(other *StreamableContent) bool {
	return other.ToC() == recv.ToC()
}

// GetMimeType is a wrapper around the C function atk_streamable_content_get_mime_type.
func (recv *StreamableContent) GetMimeType(i int32) string {
	c_i := (C.gint)(i)

	retC := C.atk_streamable_content_get_mime_type((*C.AtkStreamableContent)(recv.native), c_i)
	retGo := C.GoString(retC)

	return retGo
}

// GetNMimeTypes is a wrapper around the C function atk_streamable_content_get_n_mime_types.
func (recv *StreamableContent) GetNMimeTypes() int32 {
	retC := C.atk_streamable_content_get_n_mime_types((*C.AtkStreamableContent)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetStream is a wrapper around the C function atk_streamable_content_get_stream.
func (recv *StreamableContent) GetStream(mimeType string) *glib.IOChannel {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	retC := C.atk_streamable_content_get_stream((*C.AtkStreamableContent)(recv.native), c_mime_type)
	retGo := glib.IOChannelNewFromC(unsafe.Pointer(retC))

	return retGo
}

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
type TableSignalColumnDeletedCallback func(targetObject *Table, arg1 int32, arg2 int32)

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
func table_columnDeletedHandler(c_targetObject *C.GObject, c_arg1 C.gint, c_arg2 C.gint, data C.gpointer) {
	signalTableColumnDeletedLock.RLock()
	defer signalTableColumnDeletedLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	targetObject := TableNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTableColumnDeletedMap[index].callback
	callback(targetObject, arg1, arg2)
}

type signalTableColumnInsertedDetail struct {
	callback  TableSignalColumnInsertedCallback
	handlerID C.gulong
}

var signalTableColumnInsertedId int
var signalTableColumnInsertedMap = make(map[int]signalTableColumnInsertedDetail)
var signalTableColumnInsertedLock sync.RWMutex

// TableSignalColumnInsertedCallback is a callback function for a 'column-inserted' signal emitted from a Table.
type TableSignalColumnInsertedCallback func(targetObject *Table, arg1 int32, arg2 int32)

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
func table_columnInsertedHandler(c_targetObject *C.GObject, c_arg1 C.gint, c_arg2 C.gint, data C.gpointer) {
	signalTableColumnInsertedLock.RLock()
	defer signalTableColumnInsertedLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	targetObject := TableNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTableColumnInsertedMap[index].callback
	callback(targetObject, arg1, arg2)
}

type signalTableColumnReorderedDetail struct {
	callback  TableSignalColumnReorderedCallback
	handlerID C.gulong
}

var signalTableColumnReorderedId int
var signalTableColumnReorderedMap = make(map[int]signalTableColumnReorderedDetail)
var signalTableColumnReorderedLock sync.RWMutex

// TableSignalColumnReorderedCallback is a callback function for a 'column-reordered' signal emitted from a Table.
type TableSignalColumnReorderedCallback func(targetObject *Table)

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
func table_columnReorderedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalTableColumnReorderedLock.RLock()
	defer signalTableColumnReorderedLock.RUnlock()

	targetObject := TableNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTableColumnReorderedMap[index].callback
	callback(targetObject)
}

type signalTableModelChangedDetail struct {
	callback  TableSignalModelChangedCallback
	handlerID C.gulong
}

var signalTableModelChangedId int
var signalTableModelChangedMap = make(map[int]signalTableModelChangedDetail)
var signalTableModelChangedLock sync.RWMutex

// TableSignalModelChangedCallback is a callback function for a 'model-changed' signal emitted from a Table.
type TableSignalModelChangedCallback func(targetObject *Table)

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
func table_modelChangedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalTableModelChangedLock.RLock()
	defer signalTableModelChangedLock.RUnlock()

	targetObject := TableNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTableModelChangedMap[index].callback
	callback(targetObject)
}

type signalTableRowDeletedDetail struct {
	callback  TableSignalRowDeletedCallback
	handlerID C.gulong
}

var signalTableRowDeletedId int
var signalTableRowDeletedMap = make(map[int]signalTableRowDeletedDetail)
var signalTableRowDeletedLock sync.RWMutex

// TableSignalRowDeletedCallback is a callback function for a 'row-deleted' signal emitted from a Table.
type TableSignalRowDeletedCallback func(targetObject *Table, arg1 int32, arg2 int32)

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
func table_rowDeletedHandler(c_targetObject *C.GObject, c_arg1 C.gint, c_arg2 C.gint, data C.gpointer) {
	signalTableRowDeletedLock.RLock()
	defer signalTableRowDeletedLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	targetObject := TableNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTableRowDeletedMap[index].callback
	callback(targetObject, arg1, arg2)
}

type signalTableRowInsertedDetail struct {
	callback  TableSignalRowInsertedCallback
	handlerID C.gulong
}

var signalTableRowInsertedId int
var signalTableRowInsertedMap = make(map[int]signalTableRowInsertedDetail)
var signalTableRowInsertedLock sync.RWMutex

// TableSignalRowInsertedCallback is a callback function for a 'row-inserted' signal emitted from a Table.
type TableSignalRowInsertedCallback func(targetObject *Table, arg1 int32, arg2 int32)

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
func table_rowInsertedHandler(c_targetObject *C.GObject, c_arg1 C.gint, c_arg2 C.gint, data C.gpointer) {
	signalTableRowInsertedLock.RLock()
	defer signalTableRowInsertedLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	targetObject := TableNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTableRowInsertedMap[index].callback
	callback(targetObject, arg1, arg2)
}

type signalTableRowReorderedDetail struct {
	callback  TableSignalRowReorderedCallback
	handlerID C.gulong
}

var signalTableRowReorderedId int
var signalTableRowReorderedMap = make(map[int]signalTableRowReorderedDetail)
var signalTableRowReorderedLock sync.RWMutex

// TableSignalRowReorderedCallback is a callback function for a 'row-reordered' signal emitted from a Table.
type TableSignalRowReorderedCallback func(targetObject *Table)

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
func table_rowReorderedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalTableRowReorderedLock.RLock()
	defer signalTableRowReorderedLock.RUnlock()

	targetObject := TableNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTableRowReorderedMap[index].callback
	callback(targetObject)
}

// AddColumnSelection is a wrapper around the C function atk_table_add_column_selection.
func (recv *Table) AddColumnSelection(column int32) bool {
	c_column := (C.gint)(column)

	retC := C.atk_table_add_column_selection((*C.AtkTable)(recv.native), c_column)
	retGo := retC == C.TRUE

	return retGo
}

// AddRowSelection is a wrapper around the C function atk_table_add_row_selection.
func (recv *Table) AddRowSelection(row int32) bool {
	c_row := (C.gint)(row)

	retC := C.atk_table_add_row_selection((*C.AtkTable)(recv.native), c_row)
	retGo := retC == C.TRUE

	return retGo
}

// GetCaption is a wrapper around the C function atk_table_get_caption.
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

// GetColumnAtIndex is a wrapper around the C function atk_table_get_column_at_index.
func (recv *Table) GetColumnAtIndex(index int32) int32 {
	c_index_ := (C.gint)(index)

	retC := C.atk_table_get_column_at_index((*C.AtkTable)(recv.native), c_index_)
	retGo := (int32)(retC)

	return retGo
}

// GetColumnDescription is a wrapper around the C function atk_table_get_column_description.
func (recv *Table) GetColumnDescription(column int32) string {
	c_column := (C.gint)(column)

	retC := C.atk_table_get_column_description((*C.AtkTable)(recv.native), c_column)
	retGo := C.GoString(retC)

	return retGo
}

// GetColumnExtentAt is a wrapper around the C function atk_table_get_column_extent_at.
func (recv *Table) GetColumnExtentAt(row int32, column int32) int32 {
	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	retC := C.atk_table_get_column_extent_at((*C.AtkTable)(recv.native), c_row, c_column)
	retGo := (int32)(retC)

	return retGo
}

// GetColumnHeader is a wrapper around the C function atk_table_get_column_header.
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

// GetIndexAt is a wrapper around the C function atk_table_get_index_at.
func (recv *Table) GetIndexAt(row int32, column int32) int32 {
	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	retC := C.atk_table_get_index_at((*C.AtkTable)(recv.native), c_row, c_column)
	retGo := (int32)(retC)

	return retGo
}

// GetNColumns is a wrapper around the C function atk_table_get_n_columns.
func (recv *Table) GetNColumns() int32 {
	retC := C.atk_table_get_n_columns((*C.AtkTable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNRows is a wrapper around the C function atk_table_get_n_rows.
func (recv *Table) GetNRows() int32 {
	retC := C.atk_table_get_n_rows((*C.AtkTable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRowAtIndex is a wrapper around the C function atk_table_get_row_at_index.
func (recv *Table) GetRowAtIndex(index int32) int32 {
	c_index_ := (C.gint)(index)

	retC := C.atk_table_get_row_at_index((*C.AtkTable)(recv.native), c_index_)
	retGo := (int32)(retC)

	return retGo
}

// GetRowDescription is a wrapper around the C function atk_table_get_row_description.
func (recv *Table) GetRowDescription(row int32) string {
	c_row := (C.gint)(row)

	retC := C.atk_table_get_row_description((*C.AtkTable)(recv.native), c_row)
	retGo := C.GoString(retC)

	return retGo
}

// GetRowExtentAt is a wrapper around the C function atk_table_get_row_extent_at.
func (recv *Table) GetRowExtentAt(row int32, column int32) int32 {
	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	retC := C.atk_table_get_row_extent_at((*C.AtkTable)(recv.native), c_row, c_column)
	retGo := (int32)(retC)

	return retGo
}

// GetRowHeader is a wrapper around the C function atk_table_get_row_header.
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

// GetSummary is a wrapper around the C function atk_table_get_summary.
func (recv *Table) GetSummary() *Object {
	retC := C.atk_table_get_summary((*C.AtkTable)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsColumnSelected is a wrapper around the C function atk_table_is_column_selected.
func (recv *Table) IsColumnSelected(column int32) bool {
	c_column := (C.gint)(column)

	retC := C.atk_table_is_column_selected((*C.AtkTable)(recv.native), c_column)
	retGo := retC == C.TRUE

	return retGo
}

// IsRowSelected is a wrapper around the C function atk_table_is_row_selected.
func (recv *Table) IsRowSelected(row int32) bool {
	c_row := (C.gint)(row)

	retC := C.atk_table_is_row_selected((*C.AtkTable)(recv.native), c_row)
	retGo := retC == C.TRUE

	return retGo
}

// IsSelected is a wrapper around the C function atk_table_is_selected.
func (recv *Table) IsSelected(row int32, column int32) bool {
	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	retC := C.atk_table_is_selected((*C.AtkTable)(recv.native), c_row, c_column)
	retGo := retC == C.TRUE

	return retGo
}

// RefAt is a wrapper around the C function atk_table_ref_at.
func (recv *Table) RefAt(row int32, column int32) *Object {
	c_row := (C.gint)(row)

	c_column := (C.gint)(column)

	retC := C.atk_table_ref_at((*C.AtkTable)(recv.native), c_row, c_column)
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveColumnSelection is a wrapper around the C function atk_table_remove_column_selection.
func (recv *Table) RemoveColumnSelection(column int32) bool {
	c_column := (C.gint)(column)

	retC := C.atk_table_remove_column_selection((*C.AtkTable)(recv.native), c_column)
	retGo := retC == C.TRUE

	return retGo
}

// RemoveRowSelection is a wrapper around the C function atk_table_remove_row_selection.
func (recv *Table) RemoveRowSelection(row int32) bool {
	c_row := (C.gint)(row)

	retC := C.atk_table_remove_row_selection((*C.AtkTable)(recv.native), c_row)
	retGo := retC == C.TRUE

	return retGo
}

// SetCaption is a wrapper around the C function atk_table_set_caption.
func (recv *Table) SetCaption(caption *Object) {
	c_caption := (*C.AtkObject)(C.NULL)
	if caption != nil {
		c_caption = (*C.AtkObject)(caption.ToC())
	}

	C.atk_table_set_caption((*C.AtkTable)(recv.native), c_caption)

	return
}

// SetColumnDescription is a wrapper around the C function atk_table_set_column_description.
func (recv *Table) SetColumnDescription(column int32, description string) {
	c_column := (C.gint)(column)

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.atk_table_set_column_description((*C.AtkTable)(recv.native), c_column, c_description)

	return
}

// SetColumnHeader is a wrapper around the C function atk_table_set_column_header.
func (recv *Table) SetColumnHeader(column int32, header *Object) {
	c_column := (C.gint)(column)

	c_header := (*C.AtkObject)(C.NULL)
	if header != nil {
		c_header = (*C.AtkObject)(header.ToC())
	}

	C.atk_table_set_column_header((*C.AtkTable)(recv.native), c_column, c_header)

	return
}

// SetRowDescription is a wrapper around the C function atk_table_set_row_description.
func (recv *Table) SetRowDescription(row int32, description string) {
	c_row := (C.gint)(row)

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.atk_table_set_row_description((*C.AtkTable)(recv.native), c_row, c_description)

	return
}

// SetRowHeader is a wrapper around the C function atk_table_set_row_header.
func (recv *Table) SetRowHeader(row int32, header *Object) {
	c_row := (C.gint)(row)

	c_header := (*C.AtkObject)(C.NULL)
	if header != nil {
		c_header = (*C.AtkObject)(header.ToC())
	}

	C.atk_table_set_row_header((*C.AtkTable)(recv.native), c_row, c_header)

	return
}

// SetSummary is a wrapper around the C function atk_table_set_summary.
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
type TextSignalTextAttributesChangedCallback func(targetObject *Text)

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
func text_textAttributesChangedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalTextTextAttributesChangedLock.RLock()
	defer signalTextTextAttributesChangedLock.RUnlock()

	targetObject := TextNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTextTextAttributesChangedMap[index].callback
	callback(targetObject)
}

type signalTextTextCaretMovedDetail struct {
	callback  TextSignalTextCaretMovedCallback
	handlerID C.gulong
}

var signalTextTextCaretMovedId int
var signalTextTextCaretMovedMap = make(map[int]signalTextTextCaretMovedDetail)
var signalTextTextCaretMovedLock sync.RWMutex

// TextSignalTextCaretMovedCallback is a callback function for a 'text-caret-moved' signal emitted from a Text.
type TextSignalTextCaretMovedCallback func(targetObject *Text, arg1 int32)

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
func text_textCaretMovedHandler(c_targetObject *C.GObject, c_arg1 C.gint, data C.gpointer) {
	signalTextTextCaretMovedLock.RLock()
	defer signalTextTextCaretMovedLock.RUnlock()

	arg1 := int32(c_arg1)

	targetObject := TextNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTextTextCaretMovedMap[index].callback
	callback(targetObject, arg1)
}

type signalTextTextChangedDetail struct {
	callback  TextSignalTextChangedCallback
	handlerID C.gulong
}

var signalTextTextChangedId int
var signalTextTextChangedMap = make(map[int]signalTextTextChangedDetail)
var signalTextTextChangedLock sync.RWMutex

// TextSignalTextChangedCallback is a callback function for a 'text-changed' signal emitted from a Text.
type TextSignalTextChangedCallback func(targetObject *Text, arg1 int32, arg2 int32)

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
func text_textChangedHandler(c_targetObject *C.GObject, c_arg1 C.gint, c_arg2 C.gint, data C.gpointer) {
	signalTextTextChangedLock.RLock()
	defer signalTextTextChangedLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	targetObject := TextNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTextTextChangedMap[index].callback
	callback(targetObject, arg1, arg2)
}

type signalTextTextInsertDetail struct {
	callback  TextSignalTextInsertCallback
	handlerID C.gulong
}

var signalTextTextInsertId int
var signalTextTextInsertMap = make(map[int]signalTextTextInsertDetail)
var signalTextTextInsertLock sync.RWMutex

// TextSignalTextInsertCallback is a callback function for a 'text-insert' signal emitted from a Text.
type TextSignalTextInsertCallback func(targetObject *Text, arg1 int32, arg2 int32, arg3 string)

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
func text_textInsertHandler(c_targetObject *C.GObject, c_arg1 C.gint, c_arg2 C.gint, c_arg3 *C.gchar, data C.gpointer) {
	signalTextTextInsertLock.RLock()
	defer signalTextTextInsertLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	arg3 := C.GoString(c_arg3)

	targetObject := TextNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTextTextInsertMap[index].callback
	callback(targetObject, arg1, arg2, arg3)
}

type signalTextTextRemoveDetail struct {
	callback  TextSignalTextRemoveCallback
	handlerID C.gulong
}

var signalTextTextRemoveId int
var signalTextTextRemoveMap = make(map[int]signalTextTextRemoveDetail)
var signalTextTextRemoveLock sync.RWMutex

// TextSignalTextRemoveCallback is a callback function for a 'text-remove' signal emitted from a Text.
type TextSignalTextRemoveCallback func(targetObject *Text, arg1 int32, arg2 int32, arg3 string)

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
func text_textRemoveHandler(c_targetObject *C.GObject, c_arg1 C.gint, c_arg2 C.gint, c_arg3 *C.gchar, data C.gpointer) {
	signalTextTextRemoveLock.RLock()
	defer signalTextTextRemoveLock.RUnlock()

	arg1 := int32(c_arg1)

	arg2 := int32(c_arg2)

	arg3 := C.GoString(c_arg3)

	targetObject := TextNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTextTextRemoveMap[index].callback
	callback(targetObject, arg1, arg2, arg3)
}

type signalTextTextSelectionChangedDetail struct {
	callback  TextSignalTextSelectionChangedCallback
	handlerID C.gulong
}

var signalTextTextSelectionChangedId int
var signalTextTextSelectionChangedMap = make(map[int]signalTextTextSelectionChangedDetail)
var signalTextTextSelectionChangedLock sync.RWMutex

// TextSignalTextSelectionChangedCallback is a callback function for a 'text-selection-changed' signal emitted from a Text.
type TextSignalTextSelectionChangedCallback func(targetObject *Text)

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
func text_textSelectionChangedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalTextTextSelectionChangedLock.RLock()
	defer signalTextTextSelectionChangedLock.RUnlock()

	targetObject := TextNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalTextTextSelectionChangedMap[index].callback
	callback(targetObject)
}

// AddSelection is a wrapper around the C function atk_text_add_selection.
func (recv *Text) AddSelection(startOffset int32, endOffset int32) bool {
	c_start_offset := (C.gint)(startOffset)

	c_end_offset := (C.gint)(endOffset)

	retC := C.atk_text_add_selection((*C.AtkText)(recv.native), c_start_offset, c_end_offset)
	retGo := retC == C.TRUE

	return retGo
}

// GetCaretOffset is a wrapper around the C function atk_text_get_caret_offset.
func (recv *Text) GetCaretOffset() int32 {
	retC := C.atk_text_get_caret_offset((*C.AtkText)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetCharacterAtOffset is a wrapper around the C function atk_text_get_character_at_offset.
func (recv *Text) GetCharacterAtOffset(offset int32) rune {
	c_offset := (C.gint)(offset)

	retC := C.atk_text_get_character_at_offset((*C.AtkText)(recv.native), c_offset)
	retGo := (rune)(retC)

	return retGo
}

// GetCharacterCount is a wrapper around the C function atk_text_get_character_count.
func (recv *Text) GetCharacterCount() int32 {
	retC := C.atk_text_get_character_count((*C.AtkText)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetCharacterExtents is a wrapper around the C function atk_text_get_character_extents.
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

// GetNSelections is a wrapper around the C function atk_text_get_n_selections.
func (recv *Text) GetNSelections() int32 {
	retC := C.atk_text_get_n_selections((*C.AtkText)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetOffsetAtPoint is a wrapper around the C function atk_text_get_offset_at_point.
func (recv *Text) GetOffsetAtPoint(x int32, y int32, coords CoordType) int32 {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_coords := (C.AtkCoordType)(coords)

	retC := C.atk_text_get_offset_at_point((*C.AtkText)(recv.native), c_x, c_y, c_coords)
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : atk_text_get_run_attributes

// GetSelection is a wrapper around the C function atk_text_get_selection.
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

// GetText is a wrapper around the C function atk_text_get_text.
func (recv *Text) GetText(startOffset int32, endOffset int32) string {
	c_start_offset := (C.gint)(startOffset)

	c_end_offset := (C.gint)(endOffset)

	retC := C.atk_text_get_text((*C.AtkText)(recv.native), c_start_offset, c_end_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetTextAfterOffset is a wrapper around the C function atk_text_get_text_after_offset.
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

// GetTextAtOffset is a wrapper around the C function atk_text_get_text_at_offset.
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

// GetTextBeforeOffset is a wrapper around the C function atk_text_get_text_before_offset.
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

// RemoveSelection is a wrapper around the C function atk_text_remove_selection.
func (recv *Text) RemoveSelection(selectionNum int32) bool {
	c_selection_num := (C.gint)(selectionNum)

	retC := C.atk_text_remove_selection((*C.AtkText)(recv.native), c_selection_num)
	retGo := retC == C.TRUE

	return retGo
}

// SetCaretOffset is a wrapper around the C function atk_text_set_caret_offset.
func (recv *Text) SetCaretOffset(offset int32) bool {
	c_offset := (C.gint)(offset)

	retC := C.atk_text_set_caret_offset((*C.AtkText)(recv.native), c_offset)
	retGo := retC == C.TRUE

	return retGo
}

// SetSelection is a wrapper around the C function atk_text_set_selection.
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

// Equals compares this Value with another Value, and returns true if they represent the same GObject.
func (recv *Value) Equals(other *Value) bool {
	return other.ToC() == recv.ToC()
}

// GetCurrentValue is a wrapper around the C function atk_value_get_current_value.
func (recv *Value) GetCurrentValue() *gobject.Value {
	var c_value C.GValue

	C.atk_value_get_current_value((*C.AtkValue)(recv.native), &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// GetMaximumValue is a wrapper around the C function atk_value_get_maximum_value.
func (recv *Value) GetMaximumValue() *gobject.Value {
	var c_value C.GValue

	C.atk_value_get_maximum_value((*C.AtkValue)(recv.native), &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// GetMinimumValue is a wrapper around the C function atk_value_get_minimum_value.
func (recv *Value) GetMinimumValue() *gobject.Value {
	var c_value C.GValue

	C.atk_value_get_minimum_value((*C.AtkValue)(recv.native), &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// SetCurrentValue is a wrapper around the C function atk_value_set_current_value.
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

// Equals compares this Window with another Window, and returns true if they represent the same GObject.
func (recv *Window) Equals(other *Window) bool {
	return other.ToC() == recv.ToC()
}

// ActionIface is a wrapper around the C record AtkActionIface.
type ActionIface struct {
	native *C.AtkActionIface
	// parent : record
	// no type for do_action
	// no type for get_n_actions
	// no type for get_description
	// no type for get_name
	// no type for get_keybinding
	// no type for set_description
	// no type for get_localized_name
}

func ActionIfaceNewFromC(u unsafe.Pointer) *ActionIface {
	c := (*C.AtkActionIface)(u)
	if c == nil {
		return nil
	}

	g := &ActionIface{native: c}

	return g
}

func (recv *ActionIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionIface with another ActionIface, and returns true if they represent the same GObject.
func (recv *ActionIface) Equals(other *ActionIface) bool {
	return other.ToC() == recv.ToC()
}

// Attribute is a wrapper around the C record AtkAttribute.
type Attribute struct {
	native *C.AtkAttribute
	Name   string
	Value  string
}

func AttributeNewFromC(u unsafe.Pointer) *Attribute {
	c := (*C.AtkAttribute)(u)
	if c == nil {
		return nil
	}

	g := &Attribute{
		Name:   C.GoString(c.name),
		Value:  C.GoString(c.value),
		native: c,
	}

	return g
}

func (recv *Attribute) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.value =
		C.CString(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Attribute with another Attribute, and returns true if they represent the same GObject.
func (recv *Attribute) Equals(other *Attribute) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : atk_attribute_set_free

// ComponentIface is a wrapper around the C record AtkComponentIface.
type ComponentIface struct {
	native *C.AtkComponentIface
	// parent : record
	// no type for add_focus_handler
	// no type for contains
	// no type for ref_accessible_at_point
	// no type for get_extents
	// no type for get_position
	// no type for get_size
	// no type for grab_focus
	// no type for remove_focus_handler
	// no type for set_extents
	// no type for set_position
	// no type for set_size
	// no type for get_layer
	// no type for get_mdi_zorder
	// no type for bounds_changed
	// no type for get_alpha
}

func ComponentIfaceNewFromC(u unsafe.Pointer) *ComponentIface {
	c := (*C.AtkComponentIface)(u)
	if c == nil {
		return nil
	}

	g := &ComponentIface{native: c}

	return g
}

func (recv *ComponentIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ComponentIface with another ComponentIface, and returns true if they represent the same GObject.
func (recv *ComponentIface) Equals(other *ComponentIface) bool {
	return other.ToC() == recv.ToC()
}

// DocumentIface is a wrapper around the C record AtkDocumentIface.
type DocumentIface struct {
	native *C.AtkDocumentIface
	// parent : record
	// no type for get_document_type
	// no type for get_document
	// no type for get_document_locale
	// no type for get_document_attributes
	// no type for get_document_attribute_value
	// no type for set_document_attribute
	// no type for get_current_page_number
	// no type for get_page_count
}

func DocumentIfaceNewFromC(u unsafe.Pointer) *DocumentIface {
	c := (*C.AtkDocumentIface)(u)
	if c == nil {
		return nil
	}

	g := &DocumentIface{native: c}

	return g
}

func (recv *DocumentIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DocumentIface with another DocumentIface, and returns true if they represent the same GObject.
func (recv *DocumentIface) Equals(other *DocumentIface) bool {
	return other.ToC() == recv.ToC()
}

// EditableTextIface is a wrapper around the C record AtkEditableTextIface.
type EditableTextIface struct {
	native *C.AtkEditableTextIface
	// parent_interface : record
	// no type for set_run_attributes
	// no type for set_text_contents
	// no type for insert_text
	// no type for copy_text
	// no type for cut_text
	// no type for delete_text
	// no type for paste_text
}

func EditableTextIfaceNewFromC(u unsafe.Pointer) *EditableTextIface {
	c := (*C.AtkEditableTextIface)(u)
	if c == nil {
		return nil
	}

	g := &EditableTextIface{native: c}

	return g
}

func (recv *EditableTextIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EditableTextIface with another EditableTextIface, and returns true if they represent the same GObject.
func (recv *EditableTextIface) Equals(other *EditableTextIface) bool {
	return other.ToC() == recv.ToC()
}

// GObjectAccessibleClass is a wrapper around the C record AtkGObjectAccessibleClass.
type GObjectAccessibleClass struct {
	native *C.AtkGObjectAccessibleClass
	// parent_class : record
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
}

func GObjectAccessibleClassNewFromC(u unsafe.Pointer) *GObjectAccessibleClass {
	c := (*C.AtkGObjectAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &GObjectAccessibleClass{native: c}

	return g
}

func (recv *GObjectAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GObjectAccessibleClass with another GObjectAccessibleClass, and returns true if they represent the same GObject.
func (recv *GObjectAccessibleClass) Equals(other *GObjectAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// HyperlinkClass is a wrapper around the C record AtkHyperlinkClass.
type HyperlinkClass struct {
	native *C.AtkHyperlinkClass
	// parent : record
	// no type for get_uri
	// no type for get_object
	// no type for get_end_index
	// no type for get_start_index
	// no type for is_valid
	// no type for get_n_anchors
	// no type for link_state
	// no type for is_selected_link
	// no type for link_activated
	// pad1 : no type generator for Function, AtkFunction
}

func HyperlinkClassNewFromC(u unsafe.Pointer) *HyperlinkClass {
	c := (*C.AtkHyperlinkClass)(u)
	if c == nil {
		return nil
	}

	g := &HyperlinkClass{native: c}

	return g
}

func (recv *HyperlinkClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HyperlinkClass with another HyperlinkClass, and returns true if they represent the same GObject.
func (recv *HyperlinkClass) Equals(other *HyperlinkClass) bool {
	return other.ToC() == recv.ToC()
}

// HyperlinkImplIface is a wrapper around the C record AtkHyperlinkImplIface.
type HyperlinkImplIface struct {
	native *C.AtkHyperlinkImplIface
	// parent : record
	// no type for get_hyperlink
}

func HyperlinkImplIfaceNewFromC(u unsafe.Pointer) *HyperlinkImplIface {
	c := (*C.AtkHyperlinkImplIface)(u)
	if c == nil {
		return nil
	}

	g := &HyperlinkImplIface{native: c}

	return g
}

func (recv *HyperlinkImplIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HyperlinkImplIface with another HyperlinkImplIface, and returns true if they represent the same GObject.
func (recv *HyperlinkImplIface) Equals(other *HyperlinkImplIface) bool {
	return other.ToC() == recv.ToC()
}

// HypertextIface is a wrapper around the C record AtkHypertextIface.
type HypertextIface struct {
	native *C.AtkHypertextIface
	// parent : record
	// no type for get_link
	// no type for get_n_links
	// no type for get_link_index
	// no type for link_selected
}

func HypertextIfaceNewFromC(u unsafe.Pointer) *HypertextIface {
	c := (*C.AtkHypertextIface)(u)
	if c == nil {
		return nil
	}

	g := &HypertextIface{native: c}

	return g
}

func (recv *HypertextIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HypertextIface with another HypertextIface, and returns true if they represent the same GObject.
func (recv *HypertextIface) Equals(other *HypertextIface) bool {
	return other.ToC() == recv.ToC()
}

// ImageIface is a wrapper around the C record AtkImageIface.
type ImageIface struct {
	native *C.AtkImageIface
	// parent : record
	// no type for get_image_position
	// no type for get_image_description
	// no type for get_image_size
	// no type for set_image_description
	// no type for get_image_locale
}

func ImageIfaceNewFromC(u unsafe.Pointer) *ImageIface {
	c := (*C.AtkImageIface)(u)
	if c == nil {
		return nil
	}

	g := &ImageIface{native: c}

	return g
}

func (recv *ImageIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ImageIface with another ImageIface, and returns true if they represent the same GObject.
func (recv *ImageIface) Equals(other *ImageIface) bool {
	return other.ToC() == recv.ToC()
}

// Implementor is a wrapper around the C record AtkImplementor.
type Implementor struct {
	native *C.AtkImplementor
}

func ImplementorNewFromC(u unsafe.Pointer) *Implementor {
	c := (*C.AtkImplementor)(u)
	if c == nil {
		return nil
	}

	g := &Implementor{native: c}

	return g
}

func (recv *Implementor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Implementor with another Implementor, and returns true if they represent the same GObject.
func (recv *Implementor) Equals(other *Implementor) bool {
	return other.ToC() == recv.ToC()
}

// RefAccessible is a wrapper around the C function atk_implementor_ref_accessible.
func (recv *Implementor) RefAccessible() *Object {
	retC := C.atk_implementor_ref_accessible((*C.AtkImplementor)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// KeyEventStruct is a wrapper around the C record AtkKeyEventStruct.
type KeyEventStruct struct {
	native    *C.AtkKeyEventStruct
	Type      int32
	State     uint32
	Keyval    uint32
	Length    int32
	String    string
	Keycode   uint16
	Timestamp uint32
}

func KeyEventStructNewFromC(u unsafe.Pointer) *KeyEventStruct {
	c := (*C.AtkKeyEventStruct)(u)
	if c == nil {
		return nil
	}

	g := &KeyEventStruct{
		Keycode:   (uint16)(c.keycode),
		Keyval:    (uint32)(c.keyval),
		Length:    (int32)(c.length),
		State:     (uint32)(c.state),
		String:    C.GoString(c.string),
		Timestamp: (uint32)(c.timestamp),
		Type:      (int32)(c._type),
		native:    c,
	}

	return g
}

func (recv *KeyEventStruct) ToC() unsafe.Pointer {
	recv.native._type =
		(C.gint)(recv.Type)
	recv.native.state =
		(C.guint)(recv.State)
	recv.native.keyval =
		(C.guint)(recv.Keyval)
	recv.native.length =
		(C.gint)(recv.Length)
	recv.native.string =
		C.CString(recv.String)
	recv.native.keycode =
		(C.guint16)(recv.Keycode)
	recv.native.timestamp =
		(C.guint32)(recv.Timestamp)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this KeyEventStruct with another KeyEventStruct, and returns true if they represent the same GObject.
func (recv *KeyEventStruct) Equals(other *KeyEventStruct) bool {
	return other.ToC() == recv.ToC()
}

// MiscClass is a wrapper around the C record AtkMiscClass.
type MiscClass struct {
	native *C.AtkMiscClass
	// parent : record
	// no type for threads_enter
	// no type for threads_leave
	// no type for vfuncs
}

func MiscClassNewFromC(u unsafe.Pointer) *MiscClass {
	c := (*C.AtkMiscClass)(u)
	if c == nil {
		return nil
	}

	g := &MiscClass{native: c}

	return g
}

func (recv *MiscClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MiscClass with another MiscClass, and returns true if they represent the same GObject.
func (recv *MiscClass) Equals(other *MiscClass) bool {
	return other.ToC() == recv.ToC()
}

// NoOpObjectClass is a wrapper around the C record AtkNoOpObjectClass.
type NoOpObjectClass struct {
	native *C.AtkNoOpObjectClass
	// parent_class : record
}

func NoOpObjectClassNewFromC(u unsafe.Pointer) *NoOpObjectClass {
	c := (*C.AtkNoOpObjectClass)(u)
	if c == nil {
		return nil
	}

	g := &NoOpObjectClass{native: c}

	return g
}

func (recv *NoOpObjectClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NoOpObjectClass with another NoOpObjectClass, and returns true if they represent the same GObject.
func (recv *NoOpObjectClass) Equals(other *NoOpObjectClass) bool {
	return other.ToC() == recv.ToC()
}

// NoOpObjectFactoryClass is a wrapper around the C record AtkNoOpObjectFactoryClass.
type NoOpObjectFactoryClass struct {
	native *C.AtkNoOpObjectFactoryClass
	// parent_class : record
}

func NoOpObjectFactoryClassNewFromC(u unsafe.Pointer) *NoOpObjectFactoryClass {
	c := (*C.AtkNoOpObjectFactoryClass)(u)
	if c == nil {
		return nil
	}

	g := &NoOpObjectFactoryClass{native: c}

	return g
}

func (recv *NoOpObjectFactoryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NoOpObjectFactoryClass with another NoOpObjectFactoryClass, and returns true if they represent the same GObject.
func (recv *NoOpObjectFactoryClass) Equals(other *NoOpObjectFactoryClass) bool {
	return other.ToC() == recv.ToC()
}

// ObjectClass is a wrapper around the C record AtkObjectClass.
type ObjectClass struct {
	native *C.AtkObjectClass
	// parent : record
	// no type for get_name
	// no type for get_description
	// no type for get_parent
	// no type for get_n_children
	// no type for ref_child
	// no type for get_index_in_parent
	// no type for ref_relation_set
	// no type for get_role
	// no type for get_layer
	// no type for get_mdi_zorder
	// no type for ref_state_set
	// no type for set_name
	// no type for set_description
	// no type for set_parent
	// no type for set_role
	// no type for connect_property_change_handler
	// no type for remove_property_change_handler
	// no type for initialize
	// no type for children_changed
	// no type for focus_event
	// no type for property_change
	// no type for state_change
	// no type for visible_data_changed
	// no type for active_descendant_changed
	// no type for get_attributes
	// no type for get_object_locale
	// pad1 : no type generator for Function, AtkFunction
}

func ObjectClassNewFromC(u unsafe.Pointer) *ObjectClass {
	c := (*C.AtkObjectClass)(u)
	if c == nil {
		return nil
	}

	g := &ObjectClass{native: c}

	return g
}

func (recv *ObjectClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ObjectClass with another ObjectClass, and returns true if they represent the same GObject.
func (recv *ObjectClass) Equals(other *ObjectClass) bool {
	return other.ToC() == recv.ToC()
}

// ObjectFactoryClass is a wrapper around the C record AtkObjectFactoryClass.
type ObjectFactoryClass struct {
	native *C.AtkObjectFactoryClass
	// parent_class : record
	// no type for create_accessible
	// no type for invalidate
	// no type for get_accessible_type
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
}

func ObjectFactoryClassNewFromC(u unsafe.Pointer) *ObjectFactoryClass {
	c := (*C.AtkObjectFactoryClass)(u)
	if c == nil {
		return nil
	}

	g := &ObjectFactoryClass{native: c}

	return g
}

func (recv *ObjectFactoryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ObjectFactoryClass with another ObjectFactoryClass, and returns true if they represent the same GObject.
func (recv *ObjectFactoryClass) Equals(other *ObjectFactoryClass) bool {
	return other.ToC() == recv.ToC()
}

// PlugClass is a wrapper around the C record AtkPlugClass.
type PlugClass struct {
	native *C.AtkPlugClass
	// parent_class : record
	// no type for get_object_id
}

func PlugClassNewFromC(u unsafe.Pointer) *PlugClass {
	c := (*C.AtkPlugClass)(u)
	if c == nil {
		return nil
	}

	g := &PlugClass{native: c}

	return g
}

func (recv *PlugClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PlugClass with another PlugClass, and returns true if they represent the same GObject.
func (recv *PlugClass) Equals(other *PlugClass) bool {
	return other.ToC() == recv.ToC()
}

// PropertyValues is a wrapper around the C record AtkPropertyValues.
type PropertyValues struct {
	native       *C.AtkPropertyValues
	PropertyName string
	// old_value : record
	// new_value : record
}

func PropertyValuesNewFromC(u unsafe.Pointer) *PropertyValues {
	c := (*C.AtkPropertyValues)(u)
	if c == nil {
		return nil
	}

	g := &PropertyValues{
		PropertyName: C.GoString(c.property_name),
		native:       c,
	}

	return g
}

func (recv *PropertyValues) ToC() unsafe.Pointer {
	recv.native.property_name =
		C.CString(recv.PropertyName)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PropertyValues with another PropertyValues, and returns true if they represent the same GObject.
func (recv *PropertyValues) Equals(other *PropertyValues) bool {
	return other.ToC() == recv.ToC()
}

// Range is a wrapper around the C record AtkRange.
type Range struct {
	native *C.AtkRange
}

func RangeNewFromC(u unsafe.Pointer) *Range {
	c := (*C.AtkRange)(u)
	if c == nil {
		return nil
	}

	g := &Range{native: c}

	return g
}

func (recv *Range) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Range with another Range, and returns true if they represent the same GObject.
func (recv *Range) Equals(other *Range) bool {
	return other.ToC() == recv.ToC()
}

// Rectangle is a wrapper around the C record AtkRectangle.
type Rectangle struct {
	native *C.AtkRectangle
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func RectangleNewFromC(u unsafe.Pointer) *Rectangle {
	c := (*C.AtkRectangle)(u)
	if c == nil {
		return nil
	}

	g := &Rectangle{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *Rectangle) ToC() unsafe.Pointer {
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)
	recv.native.width =
		(C.gint)(recv.Width)
	recv.native.height =
		(C.gint)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Rectangle with another Rectangle, and returns true if they represent the same GObject.
func (recv *Rectangle) Equals(other *Rectangle) bool {
	return other.ToC() == recv.ToC()
}

// RegistryClass is a wrapper around the C record AtkRegistryClass.
type RegistryClass struct {
	native *C.AtkRegistryClass
	// parent_class : record
}

func RegistryClassNewFromC(u unsafe.Pointer) *RegistryClass {
	c := (*C.AtkRegistryClass)(u)
	if c == nil {
		return nil
	}

	g := &RegistryClass{native: c}

	return g
}

func (recv *RegistryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RegistryClass with another RegistryClass, and returns true if they represent the same GObject.
func (recv *RegistryClass) Equals(other *RegistryClass) bool {
	return other.ToC() == recv.ToC()
}

// RelationClass is a wrapper around the C record AtkRelationClass.
type RelationClass struct {
	native *C.AtkRelationClass
	// parent : record
}

func RelationClassNewFromC(u unsafe.Pointer) *RelationClass {
	c := (*C.AtkRelationClass)(u)
	if c == nil {
		return nil
	}

	g := &RelationClass{native: c}

	return g
}

func (recv *RelationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RelationClass with another RelationClass, and returns true if they represent the same GObject.
func (recv *RelationClass) Equals(other *RelationClass) bool {
	return other.ToC() == recv.ToC()
}

// RelationSetClass is a wrapper around the C record AtkRelationSetClass.
type RelationSetClass struct {
	native *C.AtkRelationSetClass
	// parent : record
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
}

func RelationSetClassNewFromC(u unsafe.Pointer) *RelationSetClass {
	c := (*C.AtkRelationSetClass)(u)
	if c == nil {
		return nil
	}

	g := &RelationSetClass{native: c}

	return g
}

func (recv *RelationSetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RelationSetClass with another RelationSetClass, and returns true if they represent the same GObject.
func (recv *RelationSetClass) Equals(other *RelationSetClass) bool {
	return other.ToC() == recv.ToC()
}

// SelectionIface is a wrapper around the C record AtkSelectionIface.
type SelectionIface struct {
	native *C.AtkSelectionIface
	// parent : record
	// no type for add_selection
	// no type for clear_selection
	// no type for ref_selection
	// no type for get_selection_count
	// no type for is_child_selected
	// no type for remove_selection
	// no type for select_all_selection
	// no type for selection_changed
}

func SelectionIfaceNewFromC(u unsafe.Pointer) *SelectionIface {
	c := (*C.AtkSelectionIface)(u)
	if c == nil {
		return nil
	}

	g := &SelectionIface{native: c}

	return g
}

func (recv *SelectionIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SelectionIface with another SelectionIface, and returns true if they represent the same GObject.
func (recv *SelectionIface) Equals(other *SelectionIface) bool {
	return other.ToC() == recv.ToC()
}

// SocketClass is a wrapper around the C record AtkSocketClass.
type SocketClass struct {
	native *C.AtkSocketClass
	// parent_class : record
	// no type for embed
}

func SocketClassNewFromC(u unsafe.Pointer) *SocketClass {
	c := (*C.AtkSocketClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketClass{native: c}

	return g
}

func (recv *SocketClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketClass with another SocketClass, and returns true if they represent the same GObject.
func (recv *SocketClass) Equals(other *SocketClass) bool {
	return other.ToC() == recv.ToC()
}

// StateSetClass is a wrapper around the C record AtkStateSetClass.
type StateSetClass struct {
	native *C.AtkStateSetClass
	// parent : record
}

func StateSetClassNewFromC(u unsafe.Pointer) *StateSetClass {
	c := (*C.AtkStateSetClass)(u)
	if c == nil {
		return nil
	}

	g := &StateSetClass{native: c}

	return g
}

func (recv *StateSetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StateSetClass with another StateSetClass, and returns true if they represent the same GObject.
func (recv *StateSetClass) Equals(other *StateSetClass) bool {
	return other.ToC() == recv.ToC()
}

// StreamableContentIface is a wrapper around the C record AtkStreamableContentIface.
type StreamableContentIface struct {
	native *C.AtkStreamableContentIface
	// parent : record
	// no type for get_n_mime_types
	// no type for get_mime_type
	// no type for get_stream
	// no type for get_uri
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
	// pad3 : no type generator for Function, AtkFunction
}

func StreamableContentIfaceNewFromC(u unsafe.Pointer) *StreamableContentIface {
	c := (*C.AtkStreamableContentIface)(u)
	if c == nil {
		return nil
	}

	g := &StreamableContentIface{native: c}

	return g
}

func (recv *StreamableContentIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StreamableContentIface with another StreamableContentIface, and returns true if they represent the same GObject.
func (recv *StreamableContentIface) Equals(other *StreamableContentIface) bool {
	return other.ToC() == recv.ToC()
}

// TableCellIface is a wrapper around the C record AtkTableCellIface.
type TableCellIface struct {
	native *C.AtkTableCellIface
	// parent : record
	// no type for get_column_span
	// no type for get_column_header_cells
	// no type for get_position
	// no type for get_row_span
	// no type for get_row_header_cells
	// no type for get_row_column_span
	// no type for get_table
}

func TableCellIfaceNewFromC(u unsafe.Pointer) *TableCellIface {
	c := (*C.AtkTableCellIface)(u)
	if c == nil {
		return nil
	}

	g := &TableCellIface{native: c}

	return g
}

func (recv *TableCellIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TableCellIface with another TableCellIface, and returns true if they represent the same GObject.
func (recv *TableCellIface) Equals(other *TableCellIface) bool {
	return other.ToC() == recv.ToC()
}

// TableIface is a wrapper around the C record AtkTableIface.
type TableIface struct {
	native *C.AtkTableIface
	// parent : record
	// no type for ref_at
	// no type for get_index_at
	// no type for get_column_at_index
	// no type for get_row_at_index
	// no type for get_n_columns
	// no type for get_n_rows
	// no type for get_column_extent_at
	// no type for get_row_extent_at
	// no type for get_caption
	// no type for get_column_description
	// no type for get_column_header
	// no type for get_row_description
	// no type for get_row_header
	// no type for get_summary
	// no type for set_caption
	// no type for set_column_description
	// no type for set_column_header
	// no type for set_row_description
	// no type for set_row_header
	// no type for set_summary
	// no type for get_selected_columns
	// no type for get_selected_rows
	// no type for is_column_selected
	// no type for is_row_selected
	// no type for is_selected
	// no type for add_row_selection
	// no type for remove_row_selection
	// no type for add_column_selection
	// no type for remove_column_selection
	// no type for row_inserted
	// no type for column_inserted
	// no type for row_deleted
	// no type for column_deleted
	// no type for row_reordered
	// no type for column_reordered
	// no type for model_changed
}

func TableIfaceNewFromC(u unsafe.Pointer) *TableIface {
	c := (*C.AtkTableIface)(u)
	if c == nil {
		return nil
	}

	g := &TableIface{native: c}

	return g
}

func (recv *TableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TableIface with another TableIface, and returns true if they represent the same GObject.
func (recv *TableIface) Equals(other *TableIface) bool {
	return other.ToC() == recv.ToC()
}

// TextIface is a wrapper around the C record AtkTextIface.
type TextIface struct {
	native *C.AtkTextIface
	// parent : record
	// no type for get_text
	// no type for get_text_after_offset
	// no type for get_text_at_offset
	// no type for get_character_at_offset
	// no type for get_text_before_offset
	// no type for get_caret_offset
	// no type for get_run_attributes
	// no type for get_default_attributes
	// no type for get_character_extents
	// no type for get_character_count
	// no type for get_offset_at_point
	// no type for get_n_selections
	// no type for get_selection
	// no type for add_selection
	// no type for remove_selection
	// no type for set_selection
	// no type for set_caret_offset
	// no type for text_changed
	// no type for text_caret_moved
	// no type for text_selection_changed
	// no type for text_attributes_changed
	// no type for get_range_extents
	// no type for get_bounded_ranges
	// no type for get_string_at_offset
}

func TextIfaceNewFromC(u unsafe.Pointer) *TextIface {
	c := (*C.AtkTextIface)(u)
	if c == nil {
		return nil
	}

	g := &TextIface{native: c}

	return g
}

func (recv *TextIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TextIface with another TextIface, and returns true if they represent the same GObject.
func (recv *TextIface) Equals(other *TextIface) bool {
	return other.ToC() == recv.ToC()
}

// TextRange is a wrapper around the C record AtkTextRange.
type TextRange struct {
	native *C.AtkTextRange
	// bounds : record
	StartOffset int32
	EndOffset   int32
	Content     string
}

func TextRangeNewFromC(u unsafe.Pointer) *TextRange {
	c := (*C.AtkTextRange)(u)
	if c == nil {
		return nil
	}

	g := &TextRange{
		Content:     C.GoString(c.content),
		EndOffset:   (int32)(c.end_offset),
		StartOffset: (int32)(c.start_offset),
		native:      c,
	}

	return g
}

func (recv *TextRange) ToC() unsafe.Pointer {
	recv.native.start_offset =
		(C.gint)(recv.StartOffset)
	recv.native.end_offset =
		(C.gint)(recv.EndOffset)
	recv.native.content =
		C.CString(recv.Content)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TextRange with another TextRange, and returns true if they represent the same GObject.
func (recv *TextRange) Equals(other *TextRange) bool {
	return other.ToC() == recv.ToC()
}

// TextRectangle is a wrapper around the C record AtkTextRectangle.
type TextRectangle struct {
	native *C.AtkTextRectangle
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func TextRectangleNewFromC(u unsafe.Pointer) *TextRectangle {
	c := (*C.AtkTextRectangle)(u)
	if c == nil {
		return nil
	}

	g := &TextRectangle{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *TextRectangle) ToC() unsafe.Pointer {
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)
	recv.native.width =
		(C.gint)(recv.Width)
	recv.native.height =
		(C.gint)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TextRectangle with another TextRectangle, and returns true if they represent the same GObject.
func (recv *TextRectangle) Equals(other *TextRectangle) bool {
	return other.ToC() == recv.ToC()
}

// UtilClass is a wrapper around the C record AtkUtilClass.
type UtilClass struct {
	native *C.AtkUtilClass
	// parent : record
	// no type for add_global_event_listener
	// no type for remove_global_event_listener
	// no type for add_key_event_listener
	// no type for remove_key_event_listener
	// no type for get_root
	// no type for get_toolkit_name
	// no type for get_toolkit_version
}

func UtilClassNewFromC(u unsafe.Pointer) *UtilClass {
	c := (*C.AtkUtilClass)(u)
	if c == nil {
		return nil
	}

	g := &UtilClass{native: c}

	return g
}

func (recv *UtilClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UtilClass with another UtilClass, and returns true if they represent the same GObject.
func (recv *UtilClass) Equals(other *UtilClass) bool {
	return other.ToC() == recv.ToC()
}

// ValueIface is a wrapper around the C record AtkValueIface.
type ValueIface struct {
	native *C.AtkValueIface
	// parent : record
	// no type for get_current_value
	// no type for get_maximum_value
	// no type for get_minimum_value
	// no type for set_current_value
	// no type for get_minimum_increment
	// no type for get_value_and_text
	// no type for get_range
	// no type for get_increment
	// no type for get_sub_ranges
	// no type for set_value
}

func ValueIfaceNewFromC(u unsafe.Pointer) *ValueIface {
	c := (*C.AtkValueIface)(u)
	if c == nil {
		return nil
	}

	g := &ValueIface{native: c}

	return g
}

func (recv *ValueIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ValueIface with another ValueIface, and returns true if they represent the same GObject.
func (recv *ValueIface) Equals(other *ValueIface) bool {
	return other.ToC() == recv.ToC()
}

// WindowIface is a wrapper around the C record AtkWindowIface.
type WindowIface struct {
	native *C.AtkWindowIface
	// parent : record
}

func WindowIfaceNewFromC(u unsafe.Pointer) *WindowIface {
	c := (*C.AtkWindowIface)(u)
	if c == nil {
		return nil
	}

	g := &WindowIface{native: c}

	return g
}

func (recv *WindowIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WindowIface with another WindowIface, and returns true if they represent the same GObject.
func (recv *WindowIface) Equals(other *WindowIface) bool {
	return other.ToC() == recv.ToC()
}
