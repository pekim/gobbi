// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// CClosure is a wrapper around the C record GCClosure.
type CClosure struct {
	native *C.GCClosure
	// closure : record
	Callback uintptr
}

func CClosureNewFromC(u unsafe.Pointer) *CClosure {
	c := (*C.GCClosure)(u)
	if c == nil {
		return nil
	}

	g := &CClosure{
		Callback: (uintptr)(c.callback),
		native:   c,
	}

	return g
}

func (recv *CClosure) ToC() unsafe.Pointer {
	recv.native.callback =
		(C.gpointer)(recv.Callback)

	return (unsafe.Pointer)(recv.native)
}

// Closure is a wrapper around the C record GClosure.
type Closure struct {
	native *C.GClosure
	// Private : ref_count
	// Private : meta_marshal_nouse
	// Private : n_guards
	// Private : n_fnotifiers
	// Private : n_inotifiers
	// Private : in_inotify
	// Private : floating
	// Private : derivative_flag
	// Bitfield not supported :  1 in_marshal
	// Bitfield not supported :  1 is_invalid
	// no type for marshal
	// Private : data
	// Private : notifiers
}

func ClosureNewFromC(u unsafe.Pointer) *Closure {
	c := (*C.GClosure)(u)
	if c == nil {
		return nil
	}

	g := &Closure{native: c}

	return g
}

func (recv *Closure) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// A variant of g_closure_new_simple() which stores @object in the
// @data field of the closure and calls g_object_watch_closure() on
// @object and the created closure. This function is mainly useful
// when implementing new types of closures.
/*

C function : g_closure_new_object
*/
func ClosureNewObject(sizeofClosure uint32, object *Object) *Closure {
	c_sizeof_closure := (C.guint)(sizeofClosure)

	c_object := (*C.GObject)(C.NULL)
	if object != nil {
		c_object = (*C.GObject)(object.ToC())
	}

	retC := C.g_closure_new_object(c_sizeof_closure, c_object)
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Allocates a struct of the given size and initializes the initial
// part as a #GClosure. This function is mainly useful when
// implementing new types of closures.
//
// |[<!-- language="C" -->
// typedef struct _MyClosure MyClosure;
// struct _MyClosure
// {
// GClosure closure;
// extra data goes here
// };
//
// static void
// my_closure_finalize (gpointer  notify_data,
// GClosure *closure)
// {
// MyClosure *my_closure = (MyClosure *)closure;
//
// free extra data here
// }
//
// MyClosure *my_closure_new (gpointer data)
// {
// GClosure *closure;
// MyClosure *my_closure;
//
// closure = g_closure_new_simple (sizeof (MyClosure), data);
// my_closure = (MyClosure *) closure;
//
// initialize extra data here
//
// g_closure_add_finalize_notifier (closure, notify_data,
// my_closure_finalize);
// return my_closure;
// }
// ]|
/*

C function : g_closure_new_simple
*/
func ClosureNewSimple(sizeofClosure uint32, data uintptr) *Closure {
	c_sizeof_closure := (C.guint)(sizeofClosure)

	c_data := (C.gpointer)(data)

	retC := C.g_closure_new_simple(c_sizeof_closure, c_data)
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_closure_add_finalize_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_add_invalidate_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_add_marshal_guards : unsupported parameter pre_marshal_notify : no type generator for ClosureNotify (GClosureNotify) for param pre_marshal_notify

// Sets a flag on the closure to indicate that its calling
// environment has become invalid, and thus causes any future
// invocations of g_closure_invoke() on this @closure to be
// ignored. Also, invalidation notifiers installed on the closure will
// be called at this point. Note that unless you are holding a
// reference to the closure yourself, the invalidation notifiers may
// unref the closure and cause it to be destroyed, so if you need to
// access the closure after calling g_closure_invalidate(), make sure
// that you've previously called g_closure_ref().
//
// Note that g_closure_invalidate() will also be called when the
// reference count of a closure drops to zero (unless it has already
// been invalidated before).
/*

C function : g_closure_invalidate
*/
func (recv *Closure) Invalidate() {
	C.g_closure_invalidate((*C.GClosure)(recv.native))

	return
}

// Unsupported : g_closure_invoke : unsupported parameter param_values :

// Increments the reference count on a closure to force it staying
// alive while the caller holds a pointer to it.
/*

C function : g_closure_ref
*/
func (recv *Closure) Ref() *Closure {
	retC := C.g_closure_ref((*C.GClosure)(recv.native))
	retGo := ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_closure_remove_finalize_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_remove_invalidate_notifier : unsupported parameter notify_func : no type generator for ClosureNotify (GClosureNotify) for param notify_func

// Unsupported : g_closure_set_marshal : unsupported parameter marshal : no type generator for ClosureMarshal (GClosureMarshal) for param marshal

// Unsupported : g_closure_set_meta_marshal : unsupported parameter meta_marshal : no type generator for ClosureMarshal (GClosureMarshal) for param meta_marshal

// Takes over the initial ownership of a closure.  Each closure is
// initially created in a "floating" state, which means that the initial
// reference count is not owned by any caller. g_closure_sink() checks
// to see if the object is still floating, and if so, unsets the
// floating state and decreases the reference count. If the closure
// is not floating, g_closure_sink() does nothing. The reason for the
// existence of the floating state is to prevent cumbersome code
// sequences like:
// |[<!-- language="C" -->
// closure = g_cclosure_new (cb_func, cb_data);
// g_source_set_closure (source, closure);
// g_closure_unref (closure); // GObject doesn't really need this
// ]|
// Because g_source_set_closure() (and similar functions) take ownership of the
// initial reference count, if it is unowned, we instead can write:
// |[<!-- language="C" -->
// g_source_set_closure (source, g_cclosure_new (cb_func, cb_data));
// ]|
//
// Generally, this function is used together with g_closure_ref(). Ane example
// of storing a closure for later notification looks like:
// |[<!-- language="C" -->
// static GClosure *notify_closure = NULL;
// void
// foo_notify_set_closure (GClosure *closure)
// {
// if (notify_closure)
// g_closure_unref (notify_closure);
// notify_closure = closure;
// if (notify_closure)
// {
// g_closure_ref (notify_closure);
// g_closure_sink (notify_closure);
// }
// }
// ]|
//
// Because g_closure_sink() may decrement the reference count of a closure
// (if it hasn't been called on @closure yet) just like g_closure_unref(),
// g_closure_ref() should be called prior to this function.
/*

C function : g_closure_sink
*/
func (recv *Closure) Sink() {
	C.g_closure_sink((*C.GClosure)(recv.native))

	return
}

// Decrements the reference count of a closure after it was previously
// incremented by the same caller. If no other callers are using the
// closure, then the closure will be destroyed and freed.
/*

C function : g_closure_unref
*/
func (recv *Closure) Unref() {
	C.g_closure_unref((*C.GClosure)(recv.native))

	return
}

// ClosureNotifyData is a wrapper around the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native *C.GClosureNotifyData
	Data   uintptr
	// notify : no type generator for ClosureNotify, GClosureNotify
}

func ClosureNotifyDataNewFromC(u unsafe.Pointer) *ClosureNotifyData {
	c := (*C.GClosureNotifyData)(u)
	if c == nil {
		return nil
	}

	g := &ClosureNotifyData{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *ClosureNotifyData) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return (unsafe.Pointer)(recv.native)
}

// EnumClass is a wrapper around the C record GEnumClass.
type EnumClass struct {
	native *C.GEnumClass
	// g_type_class : record
	Minimum int32
	Maximum int32
	NValues uint32
	// values : record
}

func EnumClassNewFromC(u unsafe.Pointer) *EnumClass {
	c := (*C.GEnumClass)(u)
	if c == nil {
		return nil
	}

	g := &EnumClass{
		Maximum: (int32)(c.maximum),
		Minimum: (int32)(c.minimum),
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *EnumClass) ToC() unsafe.Pointer {
	recv.native.minimum =
		(C.gint)(recv.Minimum)
	recv.native.maximum =
		(C.gint)(recv.Maximum)
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// EnumValue is a wrapper around the C record GEnumValue.
type EnumValue struct {
	native    *C.GEnumValue
	Value     int32
	ValueName string
	ValueNick string
}

func EnumValueNewFromC(u unsafe.Pointer) *EnumValue {
	c := (*C.GEnumValue)(u)
	if c == nil {
		return nil
	}

	g := &EnumValue{
		Value:     (int32)(c.value),
		ValueName: C.GoString(c.value_name),
		ValueNick: C.GoString(c.value_nick),
		native:    c,
	}

	return g
}

func (recv *EnumValue) ToC() unsafe.Pointer {
	recv.native.value =
		(C.gint)(recv.Value)
	recv.native.value_name =
		C.CString(recv.ValueName)
	recv.native.value_nick =
		C.CString(recv.ValueNick)

	return (unsafe.Pointer)(recv.native)
}

// FlagsClass is a wrapper around the C record GFlagsClass.
type FlagsClass struct {
	native *C.GFlagsClass
	// g_type_class : record
	Mask    uint32
	NValues uint32
	// values : record
}

func FlagsClassNewFromC(u unsafe.Pointer) *FlagsClass {
	c := (*C.GFlagsClass)(u)
	if c == nil {
		return nil
	}

	g := &FlagsClass{
		Mask:    (uint32)(c.mask),
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *FlagsClass) ToC() unsafe.Pointer {
	recv.native.mask =
		(C.guint)(recv.Mask)
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// FlagsValue is a wrapper around the C record GFlagsValue.
type FlagsValue struct {
	native    *C.GFlagsValue
	Value     uint32
	ValueName string
	ValueNick string
}

func FlagsValueNewFromC(u unsafe.Pointer) *FlagsValue {
	c := (*C.GFlagsValue)(u)
	if c == nil {
		return nil
	}

	g := &FlagsValue{
		Value:     (uint32)(c.value),
		ValueName: C.GoString(c.value_name),
		ValueNick: C.GoString(c.value_nick),
		native:    c,
	}

	return g
}

func (recv *FlagsValue) ToC() unsafe.Pointer {
	recv.native.value =
		(C.guint)(recv.Value)
	recv.native.value_name =
		C.CString(recv.ValueName)
	recv.native.value_nick =
		C.CString(recv.ValueNick)

	return (unsafe.Pointer)(recv.native)
}

// InitiallyUnownedClass is a wrapper around the C record GInitiallyUnownedClass.
type InitiallyUnownedClass struct {
	native *C.GInitiallyUnownedClass
	// g_type_class : record
	// Private : construct_properties
	// no type for constructor
	// no type for set_property
	// no type for get_property
	// no type for dispose
	// no type for finalize
	// no type for dispatch_properties_changed
	// no type for notify
	// no type for constructed
	// Private : flags
	// Private : pdummy
}

func InitiallyUnownedClassNewFromC(u unsafe.Pointer) *InitiallyUnownedClass {
	c := (*C.GInitiallyUnownedClass)(u)
	if c == nil {
		return nil
	}

	g := &InitiallyUnownedClass{native: c}

	return g
}

func (recv *InitiallyUnownedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InterfaceInfo is a wrapper around the C record GInterfaceInfo.
type InterfaceInfo struct {
	native *C.GInterfaceInfo
	// interface_init : no type generator for InterfaceInitFunc, GInterfaceInitFunc
	// interface_finalize : no type generator for InterfaceFinalizeFunc, GInterfaceFinalizeFunc
	InterfaceData uintptr
}

func InterfaceInfoNewFromC(u unsafe.Pointer) *InterfaceInfo {
	c := (*C.GInterfaceInfo)(u)
	if c == nil {
		return nil
	}

	g := &InterfaceInfo{
		InterfaceData: (uintptr)(c.interface_data),
		native:        c,
	}

	return g
}

func (recv *InterfaceInfo) ToC() unsafe.Pointer {
	recv.native.interface_data =
		(C.gpointer)(recv.InterfaceData)

	return (unsafe.Pointer)(recv.native)
}

// ObjectClass is a wrapper around the C record GObjectClass.
type ObjectClass struct {
	native *C.GObjectClass
	// g_type_class : record
	// Private : construct_properties
	// no type for constructor
	// no type for set_property
	// no type for get_property
	// no type for dispose
	// no type for finalize
	// no type for dispatch_properties_changed
	// no type for notify
	// no type for constructed
	// Private : flags
	// Private : pdummy
}

func ObjectClassNewFromC(u unsafe.Pointer) *ObjectClass {
	c := (*C.GObjectClass)(u)
	if c == nil {
		return nil
	}

	g := &ObjectClass{native: c}

	return g
}

func (recv *ObjectClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_object_class_find_property : return type : Blacklisted record : GParamSpec

// Unsupported : g_object_class_install_property : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_object_class_list_properties : no return type

// ObjectConstructParam is a wrapper around the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native *C.GObjectConstructParam
	// pspec : record
	// value : record
}

func ObjectConstructParamNewFromC(u unsafe.Pointer) *ObjectConstructParam {
	c := (*C.GObjectConstructParam)(u)
	if c == nil {
		return nil
	}

	g := &ObjectConstructParam{native: c}

	return g
}

func (recv *ObjectConstructParam) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecClass is a wrapper around the C record GParamSpecClass.
type ParamSpecClass struct {
	native *C.GParamSpecClass
	// g_type_class : record
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
	// Private : dummy
}

func ParamSpecClassNewFromC(u unsafe.Pointer) *ParamSpecClass {
	c := (*C.GParamSpecClass)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecClass{
		ValueType: (Type)(c.value_type),
		native:    c,
	}

	return g
}

func (recv *ParamSpecClass) ToC() unsafe.Pointer {
	recv.native.value_type =
		(C.GType)(recv.ValueType)

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecPool is a wrapper around the C record GParamSpecPool.
type ParamSpecPool struct {
	native *C.GParamSpecPool
}

func ParamSpecPoolNewFromC(u unsafe.Pointer) *ParamSpecPool {
	c := (*C.GParamSpecPool)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecPool{native: c}

	return g
}

func (recv *ParamSpecPool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_param_spec_pool_insert : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_pool_list : no return type

// Gets an #GList of all #GParamSpecs owned by @owner_type in
// the pool.
/*

C function : g_param_spec_pool_list_owned
*/
func (recv *ParamSpecPool) ListOwned(ownerType Type) *glib.List {
	c_owner_type := (C.GType)(ownerType)

	retC := C.g_param_spec_pool_list_owned((*C.GParamSpecPool)(recv.native), c_owner_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_param_spec_pool_lookup : return type : Blacklisted record : GParamSpec

// Unsupported : g_param_spec_pool_remove : unsupported parameter pspec : Blacklisted record : GParamSpec

// ParamSpecTypeInfo is a wrapper around the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native       *C.GParamSpecTypeInfo
	InstanceSize uint16
	NPreallocs   uint16
	// no type for instance_init
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
}

func ParamSpecTypeInfoNewFromC(u unsafe.Pointer) *ParamSpecTypeInfo {
	c := (*C.GParamSpecTypeInfo)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecTypeInfo{
		InstanceSize: (uint16)(c.instance_size),
		NPreallocs:   (uint16)(c.n_preallocs),
		ValueType:    (Type)(c.value_type),
		native:       c,
	}

	return g
}

func (recv *ParamSpecTypeInfo) ToC() unsafe.Pointer {
	recv.native.instance_size =
		(C.guint16)(recv.InstanceSize)
	recv.native.n_preallocs =
		(C.guint16)(recv.NPreallocs)
	recv.native.value_type =
		(C.GType)(recv.ValueType)

	return (unsafe.Pointer)(recv.native)
}

// Parameter is a wrapper around the C record GParameter.
type Parameter struct {
	native *C.GParameter
	Name   string
	// value : record
}

func ParameterNewFromC(u unsafe.Pointer) *Parameter {
	c := (*C.GParameter)(u)
	if c == nil {
		return nil
	}

	g := &Parameter{
		Name:   C.GoString(c.name),
		native: c,
	}

	return g
}

func (recv *Parameter) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// SignalInvocationHint is a wrapper around the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native   *C.GSignalInvocationHint
	SignalId uint32
	Detail   glib.Quark
	RunType  SignalFlags
}

func SignalInvocationHintNewFromC(u unsafe.Pointer) *SignalInvocationHint {
	c := (*C.GSignalInvocationHint)(u)
	if c == nil {
		return nil
	}

	g := &SignalInvocationHint{
		Detail:   (glib.Quark)(c.detail),
		RunType:  (SignalFlags)(c.run_type),
		SignalId: (uint32)(c.signal_id),
		native:   c,
	}

	return g
}

func (recv *SignalInvocationHint) ToC() unsafe.Pointer {
	recv.native.signal_id =
		(C.guint)(recv.SignalId)
	recv.native.detail =
		(C.GQuark)(recv.Detail)
	recv.native.run_type =
		(C.GSignalFlags)(recv.RunType)

	return (unsafe.Pointer)(recv.native)
}

// SignalQuery is a wrapper around the C record GSignalQuery.
type SignalQuery struct {
	native      *C.GSignalQuery
	SignalId    uint32
	SignalName  string
	Itype       Type
	SignalFlags SignalFlags
	ReturnType  Type
	NParams     uint32
	// no type for param_types
}

func SignalQueryNewFromC(u unsafe.Pointer) *SignalQuery {
	c := (*C.GSignalQuery)(u)
	if c == nil {
		return nil
	}

	g := &SignalQuery{
		Itype:       (Type)(c.itype),
		NParams:     (uint32)(c.n_params),
		ReturnType:  (Type)(c.return_type),
		SignalFlags: (SignalFlags)(c.signal_flags),
		SignalId:    (uint32)(c.signal_id),
		SignalName:  C.GoString(c.signal_name),
		native:      c,
	}

	return g
}

func (recv *SignalQuery) ToC() unsafe.Pointer {
	recv.native.signal_id =
		(C.guint)(recv.SignalId)
	recv.native.signal_name =
		C.CString(recv.SignalName)
	recv.native.itype =
		(C.GType)(recv.Itype)
	recv.native.signal_flags =
		(C.GSignalFlags)(recv.SignalFlags)
	recv.native.return_type =
		(C.GType)(recv.ReturnType)
	recv.native.n_params =
		(C.guint)(recv.NParams)

	return (unsafe.Pointer)(recv.native)
}

// TypeClass is a wrapper around the C record GTypeClass.
type TypeClass struct {
	native *C.GTypeClass
	// Private : g_type
}

func TypeClassNewFromC(u unsafe.Pointer) *TypeClass {
	c := (*C.GTypeClass)(u)
	if c == nil {
		return nil
	}

	g := &TypeClass{native: c}

	return g
}

func (recv *TypeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C function : g_type_class_get_private
*/
func (recv *TypeClass) GetPrivate(privateType Type) uintptr {
	c_private_type := (C.GType)(privateType)

	retC := C.g_type_class_get_private((*C.GTypeClass)(recv.native), c_private_type)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// This is a convenience function often needed in class initializers.
// It returns the class structure of the immediate parent type of the
// class passed in.  Since derived classes hold a reference count on
// their parent classes as long as they are instantiated, the returned
// class will always exist.
//
// This function is essentially equivalent to:
// g_type_class_peek (g_type_parent (G_TYPE_FROM_CLASS (g_class)))
/*

C function : g_type_class_peek_parent
*/
func (recv *TypeClass) PeekParent() uintptr {
	retC := C.g_type_class_peek_parent((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Decrements the reference count of the class structure being passed in.
// Once the last reference count of a class has been released, classes
// may be finalized by the type system, so further dereferencing of a
// class pointer after g_type_class_unref() are invalid.
/*

C function : g_type_class_unref
*/
func (recv *TypeClass) Unref() {
	C.g_type_class_unref((C.gpointer)(recv.native))

	return
}

// A variant of g_type_class_unref() for use in #GTypeClassCacheFunc
// implementations. It unreferences a class without consulting the chain
// of #GTypeClassCacheFuncs, avoiding the recursion which would occur
// otherwise.
/*

C function : g_type_class_unref_uncached
*/
func (recv *TypeClass) UnrefUncached() {
	C.g_type_class_unref_uncached((C.gpointer)(recv.native))

	return
}

// TypeFundamentalInfo is a wrapper around the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native    *C.GTypeFundamentalInfo
	TypeFlags TypeFundamentalFlags
}

func TypeFundamentalInfoNewFromC(u unsafe.Pointer) *TypeFundamentalInfo {
	c := (*C.GTypeFundamentalInfo)(u)
	if c == nil {
		return nil
	}

	g := &TypeFundamentalInfo{
		TypeFlags: (TypeFundamentalFlags)(c.type_flags),
		native:    c,
	}

	return g
}

func (recv *TypeFundamentalInfo) ToC() unsafe.Pointer {
	recv.native.type_flags =
		(C.GTypeFundamentalFlags)(recv.TypeFlags)

	return (unsafe.Pointer)(recv.native)
}

// TypeInfo is a wrapper around the C record GTypeInfo.
type TypeInfo struct {
	native    *C.GTypeInfo
	ClassSize uint16
	// base_init : no type generator for BaseInitFunc, GBaseInitFunc
	// base_finalize : no type generator for BaseFinalizeFunc, GBaseFinalizeFunc
	// class_init : no type generator for ClassInitFunc, GClassInitFunc
	// class_finalize : no type generator for ClassFinalizeFunc, GClassFinalizeFunc
	ClassData    uintptr
	InstanceSize uint16
	NPreallocs   uint16
	// instance_init : no type generator for InstanceInitFunc, GInstanceInitFunc
	// value_table : record
}

func TypeInfoNewFromC(u unsafe.Pointer) *TypeInfo {
	c := (*C.GTypeInfo)(u)
	if c == nil {
		return nil
	}

	g := &TypeInfo{
		ClassData:    (uintptr)(c.class_data),
		ClassSize:    (uint16)(c.class_size),
		InstanceSize: (uint16)(c.instance_size),
		NPreallocs:   (uint16)(c.n_preallocs),
		native:       c,
	}

	return g
}

func (recv *TypeInfo) ToC() unsafe.Pointer {
	recv.native.class_size =
		(C.guint16)(recv.ClassSize)
	recv.native.class_data =
		(C.gconstpointer)(recv.ClassData)
	recv.native.instance_size =
		(C.guint16)(recv.InstanceSize)
	recv.native.n_preallocs =
		(C.guint16)(recv.NPreallocs)

	return (unsafe.Pointer)(recv.native)
}

// TypeInstance is a wrapper around the C record GTypeInstance.
type TypeInstance struct {
	native *C.GTypeInstance
	// Private : g_class
}

func TypeInstanceNewFromC(u unsafe.Pointer) *TypeInstance {
	c := (*C.GTypeInstance)(u)
	if c == nil {
		return nil
	}

	g := &TypeInstance{native: c}

	return g
}

func (recv *TypeInstance) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C function : g_type_instance_get_private
*/
func (recv *TypeInstance) GetPrivate(privateType Type) uintptr {
	c_private_type := (C.GType)(privateType)

	retC := C.g_type_instance_get_private((*C.GTypeInstance)(recv.native), c_private_type)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TypeInterface is a wrapper around the C record GTypeInterface.
type TypeInterface struct {
	native *C.GTypeInterface
	// Private : g_type
	// Private : g_instance_type
}

func TypeInterfaceNewFromC(u unsafe.Pointer) *TypeInterface {
	c := (*C.GTypeInterface)(u)
	if c == nil {
		return nil
	}

	g := &TypeInterface{native: c}

	return g
}

func (recv *TypeInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Returns the corresponding #GTypeInterface structure of the parent type
// of the instance type to which @g_iface belongs. This is useful when
// deriving the implementation of an interface from the parent type and
// then possibly overriding some methods.
/*

C function : g_type_interface_peek_parent
*/
func (recv *TypeInterface) PeekParent() uintptr {
	retC := C.g_type_interface_peek_parent((C.gpointer)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// TypeModuleClass is a wrapper around the C record GTypeModuleClass.
type TypeModuleClass struct {
	native *C.GTypeModuleClass
	// parent_class : record
	// no type for load
	// no type for unload
	// no type for reserved1
	// no type for reserved2
	// no type for reserved3
	// no type for reserved4
}

func TypeModuleClassNewFromC(u unsafe.Pointer) *TypeModuleClass {
	c := (*C.GTypeModuleClass)(u)
	if c == nil {
		return nil
	}

	g := &TypeModuleClass{native: c}

	return g
}

func (recv *TypeModuleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypePluginClass is a wrapper around the C record GTypePluginClass.
type TypePluginClass struct {
	native *C.GTypePluginClass
	// Private : base_iface
	// use_plugin : no type generator for TypePluginUse, GTypePluginUse
	// unuse_plugin : no type generator for TypePluginUnuse, GTypePluginUnuse
	// complete_type_info : no type generator for TypePluginCompleteTypeInfo, GTypePluginCompleteTypeInfo
	// complete_interface_info : no type generator for TypePluginCompleteInterfaceInfo, GTypePluginCompleteInterfaceInfo
}

func TypePluginClassNewFromC(u unsafe.Pointer) *TypePluginClass {
	c := (*C.GTypePluginClass)(u)
	if c == nil {
		return nil
	}

	g := &TypePluginClass{native: c}

	return g
}

func (recv *TypePluginClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeQuery is a wrapper around the C record GTypeQuery.
type TypeQuery struct {
	native       *C.GTypeQuery
	Type         Type
	TypeName     string
	ClassSize    uint32
	InstanceSize uint32
}

func TypeQueryNewFromC(u unsafe.Pointer) *TypeQuery {
	c := (*C.GTypeQuery)(u)
	if c == nil {
		return nil
	}

	g := &TypeQuery{
		ClassSize:    (uint32)(c.class_size),
		InstanceSize: (uint32)(c.instance_size),
		Type:         (Type)(c._type),
		TypeName:     C.GoString(c.type_name),
		native:       c,
	}

	return g
}

func (recv *TypeQuery) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GType)(recv.Type)
	recv.native.type_name =
		C.CString(recv.TypeName)
	recv.native.class_size =
		(C.guint)(recv.ClassSize)
	recv.native.instance_size =
		(C.guint)(recv.InstanceSize)

	return (unsafe.Pointer)(recv.native)
}

// TypeValueTable is a wrapper around the C record GTypeValueTable.
type TypeValueTable struct {
	native *C.GTypeValueTable
	// no type for value_init
	// no type for value_free
	// no type for value_copy
	// no type for value_peek_pointer
	CollectFormat string
	// no type for collect_value
	LcopyFormat string
	// no type for lcopy_value
}

func TypeValueTableNewFromC(u unsafe.Pointer) *TypeValueTable {
	c := (*C.GTypeValueTable)(u)
	if c == nil {
		return nil
	}

	g := &TypeValueTable{
		CollectFormat: C.GoString(c.collect_format),
		LcopyFormat:   C.GoString(c.lcopy_format),
		native:        c,
	}

	return g
}

func (recv *TypeValueTable) ToC() unsafe.Pointer {
	recv.native.collect_format =
		C.CString(recv.CollectFormat)
	recv.native.lcopy_format =
		C.CString(recv.LcopyFormat)

	return (unsafe.Pointer)(recv.native)
}

// Value is a wrapper around the C record GValue.
type Value struct {
	native *C.GValue
	// Private : g_type
	// no type for data
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	c := (*C.GValue)(u)
	if c == nil {
		return nil
	}

	g := &Value{native: c}

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Copies the value of @src_value into @dest_value.
/*

C function : g_value_copy
*/
func (recv *Value) Copy(destValue *Value) {
	c_dest_value := (*C.GValue)(C.NULL)
	if destValue != nil {
		c_dest_value = (*C.GValue)(destValue.ToC())
	}

	C.g_value_copy((*C.GValue)(recv.native), c_dest_value)

	return
}

// Get the contents of a %G_TYPE_BOXED derived #GValue.  Upon getting,
// the boxed value is duplicated and needs to be later freed with
// g_boxed_free(), e.g. like: g_boxed_free (G_VALUE_TYPE (@value),
// return_value);
/*

C function : g_value_dup_boxed
*/
func (recv *Value) DupBoxed() uintptr {
	retC := C.g_value_dup_boxed((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Get the contents of a %G_TYPE_OBJECT derived #GValue, increasing
// its reference count. If the contents of the #GValue are %NULL, then
// %NULL will be returned.
/*

C function : g_value_dup_object
*/
func (recv *Value) DupObject() uintptr {
	retC := C.g_value_dup_object((*C.GValue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_value_dup_param : return type : Blacklisted record : GParamSpec

// Get a copy the contents of a %G_TYPE_STRING #GValue.
/*

C function : g_value_dup_string
*/
func (recv *Value) DupString() string {
	retC := C.g_value_dup_string((*C.GValue)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Determines if @value will fit inside the size of a pointer value.
// This is an internal function introduced mainly for C marshallers.
/*

C function : g_value_fits_pointer
*/
func (recv *Value) FitsPointer() bool {
	retC := C.g_value_fits_pointer((*C.GValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Get the contents of a %G_TYPE_BOOLEAN #GValue.
/*

C function : g_value_get_boolean
*/
func (recv *Value) GetBoolean() bool {
	retC := C.g_value_get_boolean((*C.GValue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Get the contents of a %G_TYPE_BOXED derived #GValue.
/*

C function : g_value_get_boxed
*/
func (recv *Value) GetBoxed() uintptr {
	retC := C.g_value_get_boxed((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Do not use this function; it is broken on platforms where the %char
// type is unsigned, such as ARM and PowerPC.  See g_value_get_schar().
//
// Get the contents of a %G_TYPE_CHAR #GValue.
/*

C function : g_value_get_char
*/
func (recv *Value) GetChar() rune {
	retC := C.g_value_get_char((*C.GValue)(recv.native))
	retGo := (rune)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_DOUBLE #GValue.
/*

C function : g_value_get_double
*/
func (recv *Value) GetDouble() float64 {
	retC := C.g_value_get_double((*C.GValue)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_ENUM #GValue.
/*

C function : g_value_get_enum
*/
func (recv *Value) GetEnum() int32 {
	retC := C.g_value_get_enum((*C.GValue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_FLAGS #GValue.
/*

C function : g_value_get_flags
*/
func (recv *Value) GetFlags() uint32 {
	retC := C.g_value_get_flags((*C.GValue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_FLOAT #GValue.
/*

C function : g_value_get_float
*/
func (recv *Value) GetFloat() float32 {
	retC := C.g_value_get_float((*C.GValue)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_INT #GValue.
/*

C function : g_value_get_int
*/
func (recv *Value) GetInt() int32 {
	retC := C.g_value_get_int((*C.GValue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_INT64 #GValue.
/*

C function : g_value_get_int64
*/
func (recv *Value) GetInt64() int64 {
	retC := C.g_value_get_int64((*C.GValue)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_LONG #GValue.
/*

C function : g_value_get_long
*/
func (recv *Value) GetLong() int64 {
	retC := C.g_value_get_long((*C.GValue)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_OBJECT derived #GValue.
/*

C function : g_value_get_object
*/
func (recv *Value) GetObject() uintptr {
	retC := C.g_value_get_object((*C.GValue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_value_get_param : return type : Blacklisted record : GParamSpec

// Get the contents of a pointer #GValue.
/*

C function : g_value_get_pointer
*/
func (recv *Value) GetPointer() uintptr {
	retC := C.g_value_get_pointer((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Get the contents of a %G_TYPE_STRING #GValue.
/*

C function : g_value_get_string
*/
func (recv *Value) GetString() string {
	retC := C.g_value_get_string((*C.GValue)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Get the contents of a %G_TYPE_UCHAR #GValue.
/*

C function : g_value_get_uchar
*/
func (recv *Value) GetUchar() uint8 {
	retC := C.g_value_get_uchar((*C.GValue)(recv.native))
	retGo := (uint8)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_UINT #GValue.
/*

C function : g_value_get_uint
*/
func (recv *Value) GetUint() uint32 {
	retC := C.g_value_get_uint((*C.GValue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_UINT64 #GValue.
/*

C function : g_value_get_uint64
*/
func (recv *Value) GetUint64() uint64 {
	retC := C.g_value_get_uint64((*C.GValue)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Get the contents of a %G_TYPE_ULONG #GValue.
/*

C function : g_value_get_ulong
*/
func (recv *Value) GetUlong() uint64 {
	retC := C.g_value_get_ulong((*C.GValue)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Initializes @value with the default value of @type.
/*

C function : g_value_init
*/
func (recv *Value) Init(gType Type) *Value {
	c_g_type := (C.GType)(gType)

	retC := C.g_value_init((*C.GValue)(recv.native), c_g_type)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the value contents as pointer. This function asserts that
// g_value_fits_pointer() returned %TRUE for the passed in value.
// This is an internal function introduced mainly for C marshallers.
/*

C function : g_value_peek_pointer
*/
func (recv *Value) PeekPointer() uintptr {
	retC := C.g_value_peek_pointer((*C.GValue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Clears the current value in @value and resets it to the default value
// (as if the value had just been initialized).
/*

C function : g_value_reset
*/
func (recv *Value) Reset() *Value {
	retC := C.g_value_reset((*C.GValue)(recv.native))
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Set the contents of a %G_TYPE_BOOLEAN #GValue to @v_boolean.
/*

C function : g_value_set_boolean
*/
func (recv *Value) SetBoolean(vBoolean bool) {
	c_v_boolean :=
		boolToGboolean(vBoolean)

	C.g_value_set_boolean((*C.GValue)(recv.native), c_v_boolean)

	return
}

// Set the contents of a %G_TYPE_BOXED derived #GValue to @v_boxed.
/*

C function : g_value_set_boxed
*/
func (recv *Value) SetBoxed(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_set_boxed((*C.GValue)(recv.native), c_v_boxed)

	return
}

// This is an internal function introduced mainly for C marshallers.
/*

C function : g_value_set_boxed_take_ownership
*/
func (recv *Value) SetBoxedTakeOwnership(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_set_boxed_take_ownership((*C.GValue)(recv.native), c_v_boxed)

	return
}

// Set the contents of a %G_TYPE_CHAR #GValue to @v_char.
/*

C function : g_value_set_char
*/
func (recv *Value) SetChar(vChar rune) {
	c_v_char := (C.gchar)(vChar)

	C.g_value_set_char((*C.GValue)(recv.native), c_v_char)

	return
}

// Set the contents of a %G_TYPE_DOUBLE #GValue to @v_double.
/*

C function : g_value_set_double
*/
func (recv *Value) SetDouble(vDouble float64) {
	c_v_double := (C.gdouble)(vDouble)

	C.g_value_set_double((*C.GValue)(recv.native), c_v_double)

	return
}

// Set the contents of a %G_TYPE_ENUM #GValue to @v_enum.
/*

C function : g_value_set_enum
*/
func (recv *Value) SetEnum(vEnum int32) {
	c_v_enum := (C.gint)(vEnum)

	C.g_value_set_enum((*C.GValue)(recv.native), c_v_enum)

	return
}

// Set the contents of a %G_TYPE_FLAGS #GValue to @v_flags.
/*

C function : g_value_set_flags
*/
func (recv *Value) SetFlags(vFlags uint32) {
	c_v_flags := (C.guint)(vFlags)

	C.g_value_set_flags((*C.GValue)(recv.native), c_v_flags)

	return
}

// Set the contents of a %G_TYPE_FLOAT #GValue to @v_float.
/*

C function : g_value_set_float
*/
func (recv *Value) SetFloat(vFloat float32) {
	c_v_float := (C.gfloat)(vFloat)

	C.g_value_set_float((*C.GValue)(recv.native), c_v_float)

	return
}

// Sets @value from an instantiatable type via the
// value_table's collect_value() function.
/*

C function : g_value_set_instance
*/
func (recv *Value) SetInstance(instance uintptr) {
	c_instance := (C.gpointer)(instance)

	C.g_value_set_instance((*C.GValue)(recv.native), c_instance)

	return
}

// Set the contents of a %G_TYPE_INT #GValue to @v_int.
/*

C function : g_value_set_int
*/
func (recv *Value) SetInt(vInt int32) {
	c_v_int := (C.gint)(vInt)

	C.g_value_set_int((*C.GValue)(recv.native), c_v_int)

	return
}

// Set the contents of a %G_TYPE_INT64 #GValue to @v_int64.
/*

C function : g_value_set_int64
*/
func (recv *Value) SetInt64(vInt64 int64) {
	c_v_int64 := (C.gint64)(vInt64)

	C.g_value_set_int64((*C.GValue)(recv.native), c_v_int64)

	return
}

// Set the contents of a %G_TYPE_LONG #GValue to @v_long.
/*

C function : g_value_set_long
*/
func (recv *Value) SetLong(vLong int64) {
	c_v_long := (C.glong)(vLong)

	C.g_value_set_long((*C.GValue)(recv.native), c_v_long)

	return
}

// Set the contents of a %G_TYPE_OBJECT derived #GValue to @v_object.
//
// g_value_set_object() increases the reference count of @v_object
// (the #GValue holds a reference to @v_object).  If you do not wish
// to increase the reference count of the object (i.e. you wish to
// pass your current reference to the #GValue because you no longer
// need it), use g_value_take_object() instead.
//
// It is important that your #GValue holds a reference to @v_object (either its
// own, or one it has taken) to ensure that the object won't be destroyed while
// the #GValue still exists).
/*

C function : g_value_set_object
*/
func (recv *Value) SetObject(vObject uintptr) {
	c_v_object := (C.gpointer)(vObject)

	C.g_value_set_object((*C.GValue)(recv.native), c_v_object)

	return
}

// This is an internal function introduced mainly for C marshallers.
/*

C function : g_value_set_object_take_ownership
*/
func (recv *Value) SetObjectTakeOwnership(vObject uintptr) {
	c_v_object := (C.gpointer)(vObject)

	C.g_value_set_object_take_ownership((*C.GValue)(recv.native), c_v_object)

	return
}

// Unsupported : g_value_set_param : unsupported parameter param : Blacklisted record : GParamSpec

// Unsupported : g_value_set_param_take_ownership : unsupported parameter param : Blacklisted record : GParamSpec

// Set the contents of a pointer #GValue to @v_pointer.
/*

C function : g_value_set_pointer
*/
func (recv *Value) SetPointer(vPointer uintptr) {
	c_v_pointer := (C.gpointer)(vPointer)

	C.g_value_set_pointer((*C.GValue)(recv.native), c_v_pointer)

	return
}

// Set the contents of a %G_TYPE_BOXED derived #GValue to @v_boxed.
// The boxed value is assumed to be static, and is thus not duplicated
// when setting the #GValue.
/*

C function : g_value_set_static_boxed
*/
func (recv *Value) SetStaticBoxed(vBoxed uintptr) {
	c_v_boxed := (C.gconstpointer)(vBoxed)

	C.g_value_set_static_boxed((*C.GValue)(recv.native), c_v_boxed)

	return
}

// Set the contents of a %G_TYPE_STRING #GValue to @v_string.
// The string is assumed to be static, and is thus not duplicated
// when setting the #GValue.
/*

C function : g_value_set_static_string
*/
func (recv *Value) SetStaticString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_set_static_string((*C.GValue)(recv.native), c_v_string)

	return
}

// Set the contents of a %G_TYPE_STRING #GValue to @v_string.
/*

C function : g_value_set_string
*/
func (recv *Value) SetString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_set_string((*C.GValue)(recv.native), c_v_string)

	return
}

// This is an internal function introduced mainly for C marshallers.
/*

C function : g_value_set_string_take_ownership
*/
func (recv *Value) SetStringTakeOwnership(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_set_string_take_ownership((*C.GValue)(recv.native), c_v_string)

	return
}

// Set the contents of a %G_TYPE_UCHAR #GValue to @v_uchar.
/*

C function : g_value_set_uchar
*/
func (recv *Value) SetUchar(vUchar uint8) {
	c_v_uchar := (C.guchar)(vUchar)

	C.g_value_set_uchar((*C.GValue)(recv.native), c_v_uchar)

	return
}

// Set the contents of a %G_TYPE_UINT #GValue to @v_uint.
/*

C function : g_value_set_uint
*/
func (recv *Value) SetUint(vUint uint32) {
	c_v_uint := (C.guint)(vUint)

	C.g_value_set_uint((*C.GValue)(recv.native), c_v_uint)

	return
}

// Set the contents of a %G_TYPE_UINT64 #GValue to @v_uint64.
/*

C function : g_value_set_uint64
*/
func (recv *Value) SetUint64(vUint64 uint64) {
	c_v_uint64 := (C.guint64)(vUint64)

	C.g_value_set_uint64((*C.GValue)(recv.native), c_v_uint64)

	return
}

// Set the contents of a %G_TYPE_ULONG #GValue to @v_ulong.
/*

C function : g_value_set_ulong
*/
func (recv *Value) SetUlong(vUlong uint64) {
	c_v_ulong := (C.gulong)(vUlong)

	C.g_value_set_ulong((*C.GValue)(recv.native), c_v_ulong)

	return
}

// Tries to cast the contents of @src_value into a type appropriate
// to store in @dest_value, e.g. to transform a %G_TYPE_INT value
// into a %G_TYPE_FLOAT value. Performing transformations between
// value types might incur precision lossage. Especially
// transformations into strings might reveal seemingly arbitrary
// results and shouldn't be relied upon for production code (such
// as rcfile value or object property serialization).
/*

C function : g_value_transform
*/
func (recv *Value) Transform(destValue *Value) bool {
	c_dest_value := (*C.GValue)(C.NULL)
	if destValue != nil {
		c_dest_value = (*C.GValue)(destValue.ToC())
	}

	retC := C.g_value_transform((*C.GValue)(recv.native), c_dest_value)
	retGo := retC == C.TRUE

	return retGo
}

// Clears the current value in @value (if any) and "unsets" the type,
// this releases all resources associated with this GValue. An unset
// value is the same as an uninitialized (zero-filled) #GValue
// structure.
/*

C function : g_value_unset
*/
func (recv *Value) Unset() {
	C.g_value_unset((*C.GValue)(recv.native))

	return
}

// ValueArray is a wrapper around the C record GValueArray.
type ValueArray struct {
	native  *C.GValueArray
	NValues uint32
	// values : record
	// Private : n_prealloced
}

func ValueArrayNewFromC(u unsafe.Pointer) *ValueArray {
	c := (*C.GValueArray)(u)
	if c == nil {
		return nil
	}

	g := &ValueArray{
		NValues: (uint32)(c.n_values),
		native:  c,
	}

	return g
}

func (recv *ValueArray) ToC() unsafe.Pointer {
	recv.native.n_values =
		(C.guint)(recv.NValues)

	return (unsafe.Pointer)(recv.native)
}

// Allocate and initialize a new #GValueArray, optionally preserve space
// for @n_prealloced elements. New arrays always contain 0 elements,
// regardless of the value of @n_prealloced.
/*

C function : g_value_array_new
*/
func ValueArrayNew(nPrealloced uint32) *ValueArray {
	c_n_prealloced := (C.guint)(nPrealloced)

	retC := C.g_value_array_new(c_n_prealloced)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert a copy of @value as last element of @value_array. If @value is
// %NULL, an uninitialized value is appended.
/*

C function : g_value_array_append
*/
func (recv *ValueArray) Append(value *Value) *ValueArray {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_value_array_append((*C.GValueArray)(recv.native), c_value)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Construct an exact copy of a #GValueArray by duplicating all its
// contents.
/*

C function : g_value_array_copy
*/
func (recv *ValueArray) Copy() *ValueArray {
	retC := C.g_value_array_copy((*C.GValueArray)(recv.native))
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free a #GValueArray including its contents.
/*

C function : g_value_array_free
*/
func (recv *ValueArray) Free() {
	C.g_value_array_free((*C.GValueArray)(recv.native))

	return
}

// Return a pointer to the value at @index_ containd in @value_array.
/*

C function : g_value_array_get_nth
*/
func (recv *ValueArray) GetNth(index uint32) *Value {
	c_index_ := (C.guint)(index)

	retC := C.g_value_array_get_nth((*C.GValueArray)(recv.native), c_index_)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert a copy of @value at specified position into @value_array. If @value
// is %NULL, an uninitialized value is inserted.
/*

C function : g_value_array_insert
*/
func (recv *ValueArray) Insert(index uint32, value *Value) *ValueArray {
	c_index_ := (C.guint)(index)

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_value_array_insert((*C.GValueArray)(recv.native), c_index_, c_value)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert a copy of @value as first element of @value_array. If @value is
// %NULL, an uninitialized value is prepended.
/*

C function : g_value_array_prepend
*/
func (recv *ValueArray) Prepend(value *Value) *ValueArray {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.g_value_array_prepend((*C.GValueArray)(recv.native), c_value)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Remove the value at position @index_ from @value_array.
/*

C function : g_value_array_remove
*/
func (recv *ValueArray) Remove(index uint32) *ValueArray {
	c_index_ := (C.guint)(index)

	retC := C.g_value_array_remove((*C.GValueArray)(recv.native), c_index_)
	retGo := ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_value_array_sort : unsupported parameter compare_func : no type generator for GLib.CompareFunc (GCompareFunc) for param compare_func

// Unsupported : g_value_array_sort_with_data : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func

// WeakRef is a wrapper around the C record GWeakRef.
type WeakRef struct {
	native *C.GWeakRef
}

func WeakRefNewFromC(u unsafe.Pointer) *WeakRef {
	c := (*C.GWeakRef)(u)
	if c == nil {
		return nil
	}

	g := &WeakRef{native: c}

	return g
}

func (recv *WeakRef) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
