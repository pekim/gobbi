// This is a generated file - DO NOT EDIT

package atk

import (
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

// Blacklisted : atk_gobject_accessible_for_object

// Blacklisted : atk_gobject_accessible_get_object

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

// Blacklisted : atk_hyperlink_get_end_index

// Blacklisted : atk_hyperlink_get_n_anchors

// Blacklisted : atk_hyperlink_get_object

// Blacklisted : atk_hyperlink_get_start_index

// Blacklisted : atk_hyperlink_get_uri

// Blacklisted : atk_hyperlink_is_inline

// Blacklisted : atk_hyperlink_is_valid

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

// Blacklisted : atk_no_op_object_new

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

// Blacklisted : atk_no_op_object_factory_new

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

// Unsupported signal 'active-descendant-changed' for Object : unsupported parameter arg1 : no type generator for gpointer, gpointer

// Unsupported signal 'children-changed' for Object : unsupported parameter arg2 : no type generator for gpointer, gpointer

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

// Unsupported signal 'property-change' for Object : unsupported parameter arg1 : no type generator for gpointer, gpointer

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

// Blacklisted : atk_object_add_relationship

// Unsupported : atk_object_connect_property_change_handler : unsupported parameter handler : no type generator for PropertyChangeHandler (AtkPropertyChangeHandler*) for param handler

// Blacklisted : atk_object_get_description

// Blacklisted : atk_object_get_index_in_parent

// Blacklisted : atk_object_get_layer

// Blacklisted : atk_object_get_mdi_zorder

// Blacklisted : atk_object_get_n_accessible_children

// Blacklisted : atk_object_get_name

// Blacklisted : atk_object_get_parent

// Blacklisted : atk_object_get_role

// Unsupported : atk_object_initialize : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : atk_object_notify_state_change

// Blacklisted : atk_object_peek_parent

// Blacklisted : atk_object_ref_accessible_child

// Blacklisted : atk_object_ref_relation_set

// Blacklisted : atk_object_ref_state_set

// Blacklisted : atk_object_remove_property_change_handler

// Blacklisted : atk_object_remove_relationship

// Blacklisted : atk_object_set_description

// Blacklisted : atk_object_set_name

// Blacklisted : atk_object_set_parent

// Blacklisted : atk_object_set_role

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

// Blacklisted : atk_object_factory_create_accessible

// Blacklisted : atk_object_factory_get_accessible_type

// Blacklisted : atk_object_factory_invalidate

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

// Blacklisted : atk_plug_new

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

// Blacklisted : atk_registry_get_factory

// Blacklisted : atk_registry_get_factory_type

// Blacklisted : atk_registry_set_factory_type

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

// Blacklisted : atk_relation_get_relation_type

// Unsupported : atk_relation_get_target : array return type :

// Blacklisted : atk_relation_remove_target

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

// Blacklisted : atk_relation_set_new

// Blacklisted : atk_relation_set_add

// Blacklisted : atk_relation_set_contains

// Blacklisted : atk_relation_set_contains_target

// Blacklisted : atk_relation_set_get_n_relations

// Blacklisted : atk_relation_set_get_relation

// Blacklisted : atk_relation_set_get_relation_by_type

// Blacklisted : atk_relation_set_remove

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

// Blacklisted : atk_socket_new

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

// Blacklisted : atk_state_set_new

// Blacklisted : atk_state_set_add_state

// Blacklisted : atk_state_set_add_states

// Blacklisted : atk_state_set_and_sets

// Blacklisted : atk_state_set_clear_states

// Blacklisted : atk_state_set_contains_state

// Blacklisted : atk_state_set_contains_states

// Blacklisted : atk_state_set_is_empty

// Blacklisted : atk_state_set_or_sets

// Blacklisted : atk_state_set_remove_state

// Blacklisted : atk_state_set_xor_sets

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
