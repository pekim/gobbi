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

	void hyperlink_linkActivatedHandler(GObject *, gpointer);

	static gulong Hyperlink_signal_connect_link_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "link-activated", G_CALLBACK(hyperlink_linkActivatedHandler), data);
	}

*/
/*

	void object_activeDescendantChangedHandler(GObject *, gpointer, gpointer);

	static gulong Object_signal_connect_active_descendant_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "active-descendant-changed", G_CALLBACK(object_activeDescendantChangedHandler), data);
	}

*/
/*

	void object_childrenChangedHandler(GObject *, guint, gpointer, gpointer);

	static gulong Object_signal_connect_children_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "children-changed", G_CALLBACK(object_childrenChangedHandler), data);
	}

*/
/*

	void object_focusEventHandler(GObject *, gboolean, gpointer);

	static gulong Object_signal_connect_focus_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-event", G_CALLBACK(object_focusEventHandler), data);
	}

*/
/*

	void object_propertyChangeHandler(GObject *, gpointer, gpointer);

	static gulong Object_signal_connect_property_change(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "property-change", G_CALLBACK(object_propertyChangeHandler), data);
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
import "C"

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

// CastToWidget down casts any arbitary Object to GObjectAccessible.
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

// CastToWidget down casts any arbitary Object to Hyperlink.
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
type HyperlinkSignalLinkActivatedCallback func()

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
func hyperlink_linkActivatedHandler(_ *C.GObject, data C.gpointer) {
	signalHyperlinkLinkActivatedLock.RLock()
	defer signalHyperlinkLinkActivatedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalHyperlinkLinkActivatedMap[index].callback
	callback()
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

// CastToWidget down casts any arbitary Object to Misc.
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

// CastToWidget down casts any arbitary Object to NoOpObject.
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
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := NoOpObjectNewFromC(unsafe.Pointer(retC))

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

// CastToWidget down casts any arbitary Object to NoOpObjectFactory.
// Exercise care, as this is a potentially dangerous function if the Object is not a NoOpObjectFactory.
func CastToNoOpObjectFactory(object *gobject.Object) *NoOpObjectFactory {
	return NoOpObjectFactoryNewFromC(object.ToC())
}

// NoOpObjectFactoryNew is a wrapper around the C function atk_no_op_object_factory_new.
func NoOpObjectFactoryNew() *NoOpObjectFactory {
	retC := C.atk_no_op_object_factory_new()
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := NoOpObjectFactoryNewFromC(unsafe.Pointer(retC))

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

// CastToWidget down casts any arbitary Object to Object.
// Exercise care, as this is a potentially dangerous function if the Object is not a Object.
func CastToObject(object *gobject.Object) *Object {
	return ObjectNewFromC(object.ToC())
}

type signalObjectActiveDescendantChangedDetail struct {
	callback  ObjectSignalActiveDescendantChangedCallback
	handlerID C.gulong
}

var signalObjectActiveDescendantChangedId int
var signalObjectActiveDescendantChangedMap = make(map[int]signalObjectActiveDescendantChangedDetail)
var signalObjectActiveDescendantChangedLock sync.RWMutex

// ObjectSignalActiveDescendantChangedCallback is a callback function for a 'active-descendant-changed' signal emitted from a Object.
type ObjectSignalActiveDescendantChangedCallback func(arg1 uintptr)

/*
ConnectActiveDescendantChanged connects the callback to the 'active-descendant-changed' signal for the Object.

The returned value represents the connection, and may be passed to DisconnectActiveDescendantChanged to remove it.
*/
func (recv *Object) ConnectActiveDescendantChanged(callback ObjectSignalActiveDescendantChangedCallback) int {
	signalObjectActiveDescendantChangedLock.Lock()
	defer signalObjectActiveDescendantChangedLock.Unlock()

	signalObjectActiveDescendantChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Object_signal_connect_active_descendant_changed(instance, C.gpointer(uintptr(signalObjectActiveDescendantChangedId)))

	detail := signalObjectActiveDescendantChangedDetail{callback, handlerID}
	signalObjectActiveDescendantChangedMap[signalObjectActiveDescendantChangedId] = detail

	return signalObjectActiveDescendantChangedId
}

/*
DisconnectActiveDescendantChanged disconnects a callback from the 'active-descendant-changed' signal for the Object.

The connectionID should be a value returned from a call to ConnectActiveDescendantChanged.
*/
func (recv *Object) DisconnectActiveDescendantChanged(connectionID int) {
	signalObjectActiveDescendantChangedLock.Lock()
	defer signalObjectActiveDescendantChangedLock.Unlock()

	detail, exists := signalObjectActiveDescendantChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalObjectActiveDescendantChangedMap, connectionID)
}

//export object_activeDescendantChangedHandler
func object_activeDescendantChangedHandler(_ *C.GObject, c_arg1 C.gpointer, data C.gpointer) {
	signalObjectActiveDescendantChangedLock.RLock()
	defer signalObjectActiveDescendantChangedLock.RUnlock()

	arg1 := uintptr(c_arg1)

	index := int(uintptr(data))
	callback := signalObjectActiveDescendantChangedMap[index].callback
	callback(arg1)
}

type signalObjectChildrenChangedDetail struct {
	callback  ObjectSignalChildrenChangedCallback
	handlerID C.gulong
}

var signalObjectChildrenChangedId int
var signalObjectChildrenChangedMap = make(map[int]signalObjectChildrenChangedDetail)
var signalObjectChildrenChangedLock sync.RWMutex

// ObjectSignalChildrenChangedCallback is a callback function for a 'children-changed' signal emitted from a Object.
type ObjectSignalChildrenChangedCallback func(arg1 uint32, arg2 uintptr)

/*
ConnectChildrenChanged connects the callback to the 'children-changed' signal for the Object.

The returned value represents the connection, and may be passed to DisconnectChildrenChanged to remove it.
*/
func (recv *Object) ConnectChildrenChanged(callback ObjectSignalChildrenChangedCallback) int {
	signalObjectChildrenChangedLock.Lock()
	defer signalObjectChildrenChangedLock.Unlock()

	signalObjectChildrenChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Object_signal_connect_children_changed(instance, C.gpointer(uintptr(signalObjectChildrenChangedId)))

	detail := signalObjectChildrenChangedDetail{callback, handlerID}
	signalObjectChildrenChangedMap[signalObjectChildrenChangedId] = detail

	return signalObjectChildrenChangedId
}

/*
DisconnectChildrenChanged disconnects a callback from the 'children-changed' signal for the Object.

The connectionID should be a value returned from a call to ConnectChildrenChanged.
*/
func (recv *Object) DisconnectChildrenChanged(connectionID int) {
	signalObjectChildrenChangedLock.Lock()
	defer signalObjectChildrenChangedLock.Unlock()

	detail, exists := signalObjectChildrenChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalObjectChildrenChangedMap, connectionID)
}

//export object_childrenChangedHandler
func object_childrenChangedHandler(_ *C.GObject, c_arg1 C.guint, c_arg2 C.gpointer, data C.gpointer) {
	signalObjectChildrenChangedLock.RLock()
	defer signalObjectChildrenChangedLock.RUnlock()

	arg1 := uint32(c_arg1)

	arg2 := uintptr(c_arg2)

	index := int(uintptr(data))
	callback := signalObjectChildrenChangedMap[index].callback
	callback(arg1, arg2)
}

type signalObjectFocusEventDetail struct {
	callback  ObjectSignalFocusEventCallback
	handlerID C.gulong
}

var signalObjectFocusEventId int
var signalObjectFocusEventMap = make(map[int]signalObjectFocusEventDetail)
var signalObjectFocusEventLock sync.RWMutex

// ObjectSignalFocusEventCallback is a callback function for a 'focus-event' signal emitted from a Object.
type ObjectSignalFocusEventCallback func(arg1 bool)

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
func object_focusEventHandler(_ *C.GObject, c_arg1 C.gboolean, data C.gpointer) {
	signalObjectFocusEventLock.RLock()
	defer signalObjectFocusEventLock.RUnlock()

	arg1 := c_arg1 == C.TRUE

	index := int(uintptr(data))
	callback := signalObjectFocusEventMap[index].callback
	callback(arg1)
}

type signalObjectPropertyChangeDetail struct {
	callback  ObjectSignalPropertyChangeCallback
	handlerID C.gulong
}

var signalObjectPropertyChangeId int
var signalObjectPropertyChangeMap = make(map[int]signalObjectPropertyChangeDetail)
var signalObjectPropertyChangeLock sync.RWMutex

// ObjectSignalPropertyChangeCallback is a callback function for a 'property-change' signal emitted from a Object.
type ObjectSignalPropertyChangeCallback func(arg1 uintptr)

/*
ConnectPropertyChange connects the callback to the 'property-change' signal for the Object.

The returned value represents the connection, and may be passed to DisconnectPropertyChange to remove it.
*/
func (recv *Object) ConnectPropertyChange(callback ObjectSignalPropertyChangeCallback) int {
	signalObjectPropertyChangeLock.Lock()
	defer signalObjectPropertyChangeLock.Unlock()

	signalObjectPropertyChangeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Object_signal_connect_property_change(instance, C.gpointer(uintptr(signalObjectPropertyChangeId)))

	detail := signalObjectPropertyChangeDetail{callback, handlerID}
	signalObjectPropertyChangeMap[signalObjectPropertyChangeId] = detail

	return signalObjectPropertyChangeId
}

/*
DisconnectPropertyChange disconnects a callback from the 'property-change' signal for the Object.

The connectionID should be a value returned from a call to ConnectPropertyChange.
*/
func (recv *Object) DisconnectPropertyChange(connectionID int) {
	signalObjectPropertyChangeLock.Lock()
	defer signalObjectPropertyChangeLock.Unlock()

	detail, exists := signalObjectPropertyChangeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalObjectPropertyChangeMap, connectionID)
}

//export object_propertyChangeHandler
func object_propertyChangeHandler(_ *C.GObject, c_arg1 C.gpointer, data C.gpointer) {
	signalObjectPropertyChangeLock.RLock()
	defer signalObjectPropertyChangeLock.RUnlock()

	arg1 := uintptr(c_arg1)

	index := int(uintptr(data))
	callback := signalObjectPropertyChangeMap[index].callback
	callback(arg1)
}

type signalObjectStateChangeDetail struct {
	callback  ObjectSignalStateChangeCallback
	handlerID C.gulong
}

var signalObjectStateChangeId int
var signalObjectStateChangeMap = make(map[int]signalObjectStateChangeDetail)
var signalObjectStateChangeLock sync.RWMutex

// ObjectSignalStateChangeCallback is a callback function for a 'state-change' signal emitted from a Object.
type ObjectSignalStateChangeCallback func(arg1 string, arg2 bool)

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
func object_stateChangeHandler(_ *C.GObject, c_arg1 *C.gchar, c_arg2 C.gboolean, data C.gpointer) {
	signalObjectStateChangeLock.RLock()
	defer signalObjectStateChangeLock.RUnlock()

	arg1 := C.GoString(c_arg1)

	arg2 := c_arg2 == C.TRUE

	index := int(uintptr(data))
	callback := signalObjectStateChangeMap[index].callback
	callback(arg1, arg2)
}

type signalObjectVisibleDataChangedDetail struct {
	callback  ObjectSignalVisibleDataChangedCallback
	handlerID C.gulong
}

var signalObjectVisibleDataChangedId int
var signalObjectVisibleDataChangedMap = make(map[int]signalObjectVisibleDataChangedDetail)
var signalObjectVisibleDataChangedLock sync.RWMutex

// ObjectSignalVisibleDataChangedCallback is a callback function for a 'visible-data-changed' signal emitted from a Object.
type ObjectSignalVisibleDataChangedCallback func()

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
func object_visibleDataChangedHandler(_ *C.GObject, data C.gpointer) {
	signalObjectVisibleDataChangedLock.RLock()
	defer signalObjectVisibleDataChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalObjectVisibleDataChangedMap[index].callback
	callback()
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

// CastToWidget down casts any arbitary Object to ObjectFactory.
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

// CastToWidget down casts any arbitary Object to Plug.
// Exercise care, as this is a potentially dangerous function if the Object is not a Plug.
func CastToPlug(object *gobject.Object) *Plug {
	return PlugNewFromC(object.ToC())
}

// PlugNew is a wrapper around the C function atk_plug_new.
func PlugNew() *Plug {
	retC := C.atk_plug_new()
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := PlugNewFromC(unsafe.Pointer(retC))

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

// CastToWidget down casts any arbitary Object to Registry.
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

// CastToWidget down casts any arbitary Object to Relation.
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

// Unsupported : atk_relation_get_target : no return type

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

// CastToWidget down casts any arbitary Object to RelationSet.
// Exercise care, as this is a potentially dangerous function if the Object is not a RelationSet.
func CastToRelationSet(object *gobject.Object) *RelationSet {
	return RelationSetNewFromC(object.ToC())
}

// RelationSetNew is a wrapper around the C function atk_relation_set_new.
func RelationSetNew() *RelationSet {
	retC := C.atk_relation_set_new()
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := RelationSetNewFromC(unsafe.Pointer(retC))

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

// CastToWidget down casts any arbitary Object to Socket.
// Exercise care, as this is a potentially dangerous function if the Object is not a Socket.
func CastToSocket(object *gobject.Object) *Socket {
	return SocketNewFromC(object.ToC())
}

// SocketNew is a wrapper around the C function atk_socket_new.
func SocketNew() *Socket {
	retC := C.atk_socket_new()
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := SocketNewFromC(unsafe.Pointer(retC))

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

// CastToWidget down casts any arbitary Object to StateSet.
// Exercise care, as this is a potentially dangerous function if the Object is not a StateSet.
func CastToStateSet(object *gobject.Object) *StateSet {
	return StateSetNewFromC(object.ToC())
}

// StateSetNew is a wrapper around the C function atk_state_set_new.
func StateSetNew() *StateSet {
	retC := C.atk_state_set_new()
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := StateSetNewFromC(unsafe.Pointer(retC))

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
	c_types := &types[0]

	c_n_types := (C.gint)(len(types))

	C.atk_state_set_add_states((*C.AtkStateSet)(recv.native), (*C.AtkStateType)(unsafe.Pointer(c_types)), c_n_types)

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
	c_types := &types[0]

	c_n_types := (C.gint)(len(types))

	retC := C.atk_state_set_contains_states((*C.AtkStateSet)(recv.native), (*C.AtkStateType)(unsafe.Pointer(c_types)), c_n_types)
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

// CastToWidget down casts any arbitary Object to Util.
// Exercise care, as this is a potentially dangerous function if the Object is not a Util.
func CastToUtil(object *gobject.Object) *Util {
	return UtilNewFromC(object.ToC())
}
