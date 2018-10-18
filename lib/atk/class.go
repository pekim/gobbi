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

	void Hyperlink_linkActivatedHandler();

	static gulong Hyperlink_signal_connect_link_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "link-activated", Hyperlink_linkActivatedHandler, data);
	}

*/
/*

	void Object_activeDescendantChangedHandler();

	static gulong Object_signal_connect_active_descendant_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "active-descendant-changed", Object_activeDescendantChangedHandler, data);
	}

*/
/*

	void Object_childrenChangedHandler();

	static gulong Object_signal_connect_children_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "children-changed", Object_childrenChangedHandler, data);
	}

*/
/*

	void Object_focusEventHandler();

	static gulong Object_signal_connect_focus_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-event", Object_focusEventHandler, data);
	}

*/
/*

	void Object_propertyChangeHandler();

	static gulong Object_signal_connect_property_change(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "property-change", Object_propertyChangeHandler, data);
	}

*/
/*

	void Object_stateChangeHandler();

	static gulong Object_signal_connect_state_change(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-change", Object_stateChangeHandler, data);
	}

*/
/*

	void Object_visibleDataChangedHandler();

	static gulong Object_signal_connect_visible_data_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "visible-data-changed", Object_visibleDataChangedHandler, data);
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

// Object upcasts to *Object
func (recv *GObjectAccessible) Object() *Object {
	return ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to GObjectAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a GObjectAccessible.
func CastToGObjectAccessible(object *gobject.Object) *GObjectAccessible {
	return GObjectAccessibleNewFromC(object.ToC())
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

// Object upcasts to *Object
func (recv *Hyperlink) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Hyperlink.
// Exercise care, as this is a potentially dangerous function if the Object is not a Hyperlink.
func CastToHyperlink(object *gobject.Object) *Hyperlink {
	return HyperlinkNewFromC(object.ToC())
}

var signalHyperlinkLinkActivatedId int
var signalHyperlinkLinkActivatedMap = make(map[int]HyperlinkSignalLinkActivatedCallback)
var signalHyperlinkLinkActivatedLock sync.Mutex

// HyperlinkSignalLinkActivatedCallback is a callback function for a 'link-activated' signal emitted from a Hyperlink.
type HyperlinkSignalLinkActivatedCallback func()

func (recv *Hyperlink) ConnectLinkActivated(callback HyperlinkSignalLinkActivatedCallback) {
	signalHyperlinkLinkActivatedLock.Lock()
	defer signalHyperlinkLinkActivatedLock.Unlock()

	signalHyperlinkLinkActivatedId++
	signalHyperlinkLinkActivatedMap[signalHyperlinkLinkActivatedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Hyperlink_signal_connect_link_activated(instance, C.gpointer(uintptr(signalHyperlinkLinkActivatedId)))
}

//export Hyperlink_linkActivatedHandler
func Hyperlink_linkActivatedHandler() {}

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
	c_obj := (*C.GObject)(obj.ToC())

	retC := C.atk_no_op_object_new(c_obj)
	retGo := NoOpObjectNewFromC(unsafe.Pointer(retC))

	return retGo
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

// Object upcasts to *Object
func (recv *Object) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Object.
// Exercise care, as this is a potentially dangerous function if the Object is not a Object.
func CastToObject(object *gobject.Object) *Object {
	return ObjectNewFromC(object.ToC())
}

var signalObjectActiveDescendantChangedId int
var signalObjectActiveDescendantChangedMap = make(map[int]ObjectSignalActiveDescendantChangedCallback)
var signalObjectActiveDescendantChangedLock sync.Mutex

// ObjectSignalActiveDescendantChangedCallback is a callback function for a 'active-descendant-changed' signal emitted from a Object.
type ObjectSignalActiveDescendantChangedCallback func(arg1 uintptr)

func (recv *Object) ConnectActiveDescendantChanged(callback ObjectSignalActiveDescendantChangedCallback) {
	signalObjectActiveDescendantChangedLock.Lock()
	defer signalObjectActiveDescendantChangedLock.Unlock()

	signalObjectActiveDescendantChangedId++
	signalObjectActiveDescendantChangedMap[signalObjectActiveDescendantChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Object_signal_connect_active_descendant_changed(instance, C.gpointer(uintptr(signalObjectActiveDescendantChangedId)))
}

//export Object_activeDescendantChangedHandler
func Object_activeDescendantChangedHandler() {}

var signalObjectChildrenChangedId int
var signalObjectChildrenChangedMap = make(map[int]ObjectSignalChildrenChangedCallback)
var signalObjectChildrenChangedLock sync.Mutex

// ObjectSignalChildrenChangedCallback is a callback function for a 'children-changed' signal emitted from a Object.
type ObjectSignalChildrenChangedCallback func(arg1 uint32, arg2 uintptr)

func (recv *Object) ConnectChildrenChanged(callback ObjectSignalChildrenChangedCallback) {
	signalObjectChildrenChangedLock.Lock()
	defer signalObjectChildrenChangedLock.Unlock()

	signalObjectChildrenChangedId++
	signalObjectChildrenChangedMap[signalObjectChildrenChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Object_signal_connect_children_changed(instance, C.gpointer(uintptr(signalObjectChildrenChangedId)))
}

//export Object_childrenChangedHandler
func Object_childrenChangedHandler() {}

var signalObjectFocusEventId int
var signalObjectFocusEventMap = make(map[int]ObjectSignalFocusEventCallback)
var signalObjectFocusEventLock sync.Mutex

// ObjectSignalFocusEventCallback is a callback function for a 'focus-event' signal emitted from a Object.
type ObjectSignalFocusEventCallback func(arg1 bool)

func (recv *Object) ConnectFocusEvent(callback ObjectSignalFocusEventCallback) {
	signalObjectFocusEventLock.Lock()
	defer signalObjectFocusEventLock.Unlock()

	signalObjectFocusEventId++
	signalObjectFocusEventMap[signalObjectFocusEventId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Object_signal_connect_focus_event(instance, C.gpointer(uintptr(signalObjectFocusEventId)))
}

//export Object_focusEventHandler
func Object_focusEventHandler() {}

var signalObjectPropertyChangeId int
var signalObjectPropertyChangeMap = make(map[int]ObjectSignalPropertyChangeCallback)
var signalObjectPropertyChangeLock sync.Mutex

// ObjectSignalPropertyChangeCallback is a callback function for a 'property-change' signal emitted from a Object.
type ObjectSignalPropertyChangeCallback func(arg1 uintptr)

func (recv *Object) ConnectPropertyChange(callback ObjectSignalPropertyChangeCallback) {
	signalObjectPropertyChangeLock.Lock()
	defer signalObjectPropertyChangeLock.Unlock()

	signalObjectPropertyChangeId++
	signalObjectPropertyChangeMap[signalObjectPropertyChangeId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Object_signal_connect_property_change(instance, C.gpointer(uintptr(signalObjectPropertyChangeId)))
}

//export Object_propertyChangeHandler
func Object_propertyChangeHandler() {}

var signalObjectStateChangeId int
var signalObjectStateChangeMap = make(map[int]ObjectSignalStateChangeCallback)
var signalObjectStateChangeLock sync.Mutex

// ObjectSignalStateChangeCallback is a callback function for a 'state-change' signal emitted from a Object.
type ObjectSignalStateChangeCallback func(arg1 string, arg2 bool)

func (recv *Object) ConnectStateChange(callback ObjectSignalStateChangeCallback) {
	signalObjectStateChangeLock.Lock()
	defer signalObjectStateChangeLock.Unlock()

	signalObjectStateChangeId++
	signalObjectStateChangeMap[signalObjectStateChangeId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Object_signal_connect_state_change(instance, C.gpointer(uintptr(signalObjectStateChangeId)))
}

//export Object_stateChangeHandler
func Object_stateChangeHandler() {}

var signalObjectVisibleDataChangedId int
var signalObjectVisibleDataChangedMap = make(map[int]ObjectSignalVisibleDataChangedCallback)
var signalObjectVisibleDataChangedLock sync.Mutex

// ObjectSignalVisibleDataChangedCallback is a callback function for a 'visible-data-changed' signal emitted from a Object.
type ObjectSignalVisibleDataChangedCallback func()

func (recv *Object) ConnectVisibleDataChanged(callback ObjectSignalVisibleDataChangedCallback) {
	signalObjectVisibleDataChangedLock.Lock()
	defer signalObjectVisibleDataChangedLock.Unlock()

	signalObjectVisibleDataChangedId++
	signalObjectVisibleDataChangedMap[signalObjectVisibleDataChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Object_signal_connect_visible_data_changed(instance, C.gpointer(uintptr(signalObjectVisibleDataChangedId)))
}

//export Object_visibleDataChangedHandler
func Object_visibleDataChangedHandler() {}

// AddRelationship is a wrapper around the C function atk_object_add_relationship.
func (recv *Object) AddRelationship(relationship RelationType, target *Object) bool {
	c_relationship := (C.AtkRelationType)(relationship)

	c_target := (*C.AtkObject)(target.ToC())

	retC := C.atk_object_add_relationship((*C.AtkObject)(recv.native), c_relationship, c_target)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : atk_object_connect_property_change_handler : unsupported parameter handler : no type generator for PropertyChangeHandler, AtkPropertyChangeHandler*

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

	c_target := (*C.AtkObject)(target.ToC())

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
	c_parent := (*C.AtkObject)(parent.ToC())

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
	c_obj := (*C.GObject)(obj.ToC())

	retC := C.atk_object_factory_create_accessible((*C.AtkObjectFactory)(recv.native), c_obj)
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : atk_object_factory_get_accessible_type : no return generator

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
	retGo := PlugNewFromC(unsafe.Pointer(retC))

	return retGo
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

// Object upcasts to *Object
func (recv *Registry) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Registry.
// Exercise care, as this is a potentially dangerous function if the Object is not a Registry.
func CastToRegistry(object *gobject.Object) *Registry {
	return RegistryNewFromC(object.ToC())
}

// Unsupported : atk_registry_get_factory : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_registry_get_factory_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_registry_set_factory_type : unsupported parameter type : no type generator for GType, GType

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

// Object upcasts to *Object
func (recv *Relation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Relation.
// Exercise care, as this is a potentially dangerous function if the Object is not a Relation.
func CastToRelation(object *gobject.Object) *Relation {
	return RelationNewFromC(object.ToC())
}

// Unsupported : atk_relation_new : unsupported parameter targets : no param type

// GetRelationType is a wrapper around the C function atk_relation_get_relation_type.
func (recv *Relation) GetRelationType() RelationType {
	retC := C.atk_relation_get_relation_type((*C.AtkRelation)(recv.native))
	retGo := (RelationType)(retC)

	return retGo
}

// Unsupported : atk_relation_get_target : no return type

// RemoveTarget is a wrapper around the C function atk_relation_remove_target.
func (recv *Relation) RemoveTarget(target *Object) bool {
	c_target := (*C.AtkObject)(target.ToC())

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
	retGo := RelationSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Add is a wrapper around the C function atk_relation_set_add.
func (recv *RelationSet) Add(relation *Relation) {
	c_relation := (*C.AtkRelation)(relation.ToC())

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

	c_target := (*C.AtkObject)(target.ToC())

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
	c_relation := (*C.AtkRelation)(relation.ToC())

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
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	return retGo
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

// Unsupported : atk_state_set_add_states : unsupported parameter types : no param type

// AndSets is a wrapper around the C function atk_state_set_and_sets.
func (recv *StateSet) AndSets(compareSet *StateSet) *StateSet {
	c_compare_set := (*C.AtkStateSet)(compareSet.ToC())

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

// Unsupported : atk_state_set_contains_states : unsupported parameter types : no param type

// IsEmpty is a wrapper around the C function atk_state_set_is_empty.
func (recv *StateSet) IsEmpty() bool {
	retC := C.atk_state_set_is_empty((*C.AtkStateSet)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// OrSets is a wrapper around the C function atk_state_set_or_sets.
func (recv *StateSet) OrSets(compareSet *StateSet) *StateSet {
	c_compare_set := (*C.AtkStateSet)(compareSet.ToC())

	retC := C.atk_state_set_or_sets((*C.AtkStateSet)(recv.native), c_compare_set)
	retGo := StateSetNewFromC(unsafe.Pointer(retC))

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
	c_compare_set := (*C.AtkStateSet)(compareSet.ToC())

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

// Object upcasts to *Object
func (recv *Util) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Util.
// Exercise care, as this is a potentially dangerous function if the Object is not a Util.
func CastToUtil(object *gobject.Object) *Util {
	return UtilNewFromC(object.ToC())
}
