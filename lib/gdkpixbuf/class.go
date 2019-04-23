// This is a generated file - DO NOT EDIT

package gdkpixbuf

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
/*

	void pixbufloader_areaPreparedHandler(GObject *, gpointer);

	static gulong PixbufLoader_signal_connect_area_prepared(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "area-prepared", G_CALLBACK(pixbufloader_areaPreparedHandler), data);
	}

*/
/*

	void pixbufloader_areaUpdatedHandler(GObject *, gint, gint, gint, gint, gpointer);

	static gulong PixbufLoader_signal_connect_area_updated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "area-updated", G_CALLBACK(pixbufloader_areaUpdatedHandler), data);
	}

*/
/*

	void pixbufloader_closedHandler(GObject *, gpointer);

	static gulong PixbufLoader_signal_connect_closed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "closed", G_CALLBACK(pixbufloader_closedHandler), data);
	}

*/
/*

	void pixbufloader_sizePreparedHandler(GObject *, gint, gint, gpointer);

	static gulong PixbufLoader_signal_connect_size_prepared(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "size-prepared", G_CALLBACK(pixbufloader_sizePreparedHandler), data);
	}

*/
import "C"

// Pixbuf is a wrapper around the C record GdkPixbuf.
type Pixbuf struct {
	native *C.GdkPixbuf
}

func PixbufNewFromC(u unsafe.Pointer) *Pixbuf {
	c := (*C.GdkPixbuf)(u)
	if c == nil {
		return nil
	}

	g := &Pixbuf{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Pixbuf) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Pixbuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Pixbuf with another Pixbuf, and returns true if they represent the same GObject.
func (recv *Pixbuf) Equals(other *Pixbuf) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Pixbuf) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Pixbuf.
// Exercise care, as this is a potentially dangerous function if the Object is not a Pixbuf.
func CastToPixbuf(object *gobject.Object) *Pixbuf {
	return PixbufNewFromC(object.ToC())
}

// Blacklisted : gdk_pixbuf_new

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter destroy_fn : no type generator for PixbufDestroyNotify (GdkPixbufDestroyNotify) for param destroy_fn

// PixbufNewFromFile is a wrapper around the C function gdk_pixbuf_new_from_file.
func PixbufNewFromFile(filename string) (*Pixbuf, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_file(c_filename, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Blacklisted : gdk_pixbuf_new_from_inline

// Blacklisted : gdk_pixbuf_new_from_xpm_data

// gdk_pixbuf_from_pixdata : unsupported parameter pixdata : Blacklisted record : GdkPixdata
// Blacklisted : gdk_pixbuf_add_alpha

// Blacklisted : gdk_pixbuf_composite

// Blacklisted : gdk_pixbuf_composite_color

// Blacklisted : gdk_pixbuf_composite_color_simple

// Blacklisted : gdk_pixbuf_copy

// Blacklisted : gdk_pixbuf_copy_area

// Blacklisted : gdk_pixbuf_fill

// Blacklisted : gdk_pixbuf_get_bits_per_sample

// Blacklisted : gdk_pixbuf_get_colorspace

// Blacklisted : gdk_pixbuf_get_has_alpha

// Blacklisted : gdk_pixbuf_get_height

// Blacklisted : gdk_pixbuf_get_n_channels

// Blacklisted : gdk_pixbuf_get_option

// Unsupported : gdk_pixbuf_get_pixels : array return type :

// Blacklisted : gdk_pixbuf_get_rowstride

// Blacklisted : gdk_pixbuf_get_width

// Blacklisted : gdk_pixbuf_new_subpixbuf

// Blacklisted : gdk_pixbuf_ref

// Blacklisted : gdk_pixbuf_saturate_and_pixelate

// Unsupported : gdk_pixbuf_save : unsupported parameter ... : varargs

// Blacklisted : gdk_pixbuf_savev

// Blacklisted : gdk_pixbuf_scale

// Blacklisted : gdk_pixbuf_scale_simple

// Blacklisted : gdk_pixbuf_unref

// Icon returns the Icon interface implemented by Pixbuf
func (recv *Pixbuf) Icon() *gio.Icon {
	return gio.IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by Pixbuf
func (recv *Pixbuf) LoadableIcon() *gio.LoadableIcon {
	return gio.LoadableIconNewFromC(recv.ToC())
}

// PixbufAnimation is a wrapper around the C record GdkPixbufAnimation.
type PixbufAnimation struct {
	native *C.GdkPixbufAnimation
}

func PixbufAnimationNewFromC(u unsafe.Pointer) *PixbufAnimation {
	c := (*C.GdkPixbufAnimation)(u)
	if c == nil {
		return nil
	}

	g := &PixbufAnimation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufAnimation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufAnimation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufAnimation with another PixbufAnimation, and returns true if they represent the same GObject.
func (recv *PixbufAnimation) Equals(other *PixbufAnimation) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PixbufAnimation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PixbufAnimation.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufAnimation.
func CastToPixbufAnimation(object *gobject.Object) *PixbufAnimation {
	return PixbufAnimationNewFromC(object.ToC())
}

// Blacklisted : gdk_pixbuf_animation_new_from_file

// Blacklisted : gdk_pixbuf_animation_get_height

// Blacklisted : gdk_pixbuf_animation_get_iter

// Blacklisted : gdk_pixbuf_animation_get_static_image

// Blacklisted : gdk_pixbuf_animation_get_width

// Blacklisted : gdk_pixbuf_animation_is_static_image

// Blacklisted : gdk_pixbuf_animation_ref

// Blacklisted : gdk_pixbuf_animation_unref

// PixbufAnimationIter is a wrapper around the C record GdkPixbufAnimationIter.
type PixbufAnimationIter struct {
	native *C.GdkPixbufAnimationIter
}

func PixbufAnimationIterNewFromC(u unsafe.Pointer) *PixbufAnimationIter {
	c := (*C.GdkPixbufAnimationIter)(u)
	if c == nil {
		return nil
	}

	g := &PixbufAnimationIter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufAnimationIter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufAnimationIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufAnimationIter with another PixbufAnimationIter, and returns true if they represent the same GObject.
func (recv *PixbufAnimationIter) Equals(other *PixbufAnimationIter) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PixbufAnimationIter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PixbufAnimationIter.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufAnimationIter.
func CastToPixbufAnimationIter(object *gobject.Object) *PixbufAnimationIter {
	return PixbufAnimationIterNewFromC(object.ToC())
}

// Blacklisted : gdk_pixbuf_animation_iter_advance

// Blacklisted : gdk_pixbuf_animation_iter_get_delay_time

// Blacklisted : gdk_pixbuf_animation_iter_get_pixbuf

// Blacklisted : gdk_pixbuf_animation_iter_on_currently_loading_frame

// PixbufLoader is a wrapper around the C record GdkPixbufLoader.
type PixbufLoader struct {
	native *C.GdkPixbufLoader
	// parent_instance : record
	// Private : priv
}

func PixbufLoaderNewFromC(u unsafe.Pointer) *PixbufLoader {
	c := (*C.GdkPixbufLoader)(u)
	if c == nil {
		return nil
	}

	g := &PixbufLoader{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufLoader) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufLoader) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufLoader with another PixbufLoader, and returns true if they represent the same GObject.
func (recv *PixbufLoader) Equals(other *PixbufLoader) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PixbufLoader) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PixbufLoader.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufLoader.
func CastToPixbufLoader(object *gobject.Object) *PixbufLoader {
	return PixbufLoaderNewFromC(object.ToC())
}

type signalPixbufLoaderAreaPreparedDetail struct {
	callback  PixbufLoaderSignalAreaPreparedCallback
	handlerID C.gulong
}

var signalPixbufLoaderAreaPreparedId int
var signalPixbufLoaderAreaPreparedMap = make(map[int]signalPixbufLoaderAreaPreparedDetail)
var signalPixbufLoaderAreaPreparedLock sync.RWMutex

// PixbufLoaderSignalAreaPreparedCallback is a callback function for a 'area-prepared' signal emitted from a PixbufLoader.
type PixbufLoaderSignalAreaPreparedCallback func()

/*
ConnectAreaPrepared connects the callback to the 'area-prepared' signal for the PixbufLoader.

The returned value represents the connection, and may be passed to DisconnectAreaPrepared to remove it.
*/
func (recv *PixbufLoader) ConnectAreaPrepared(callback PixbufLoaderSignalAreaPreparedCallback) int {
	signalPixbufLoaderAreaPreparedLock.Lock()
	defer signalPixbufLoaderAreaPreparedLock.Unlock()

	signalPixbufLoaderAreaPreparedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PixbufLoader_signal_connect_area_prepared(instance, C.gpointer(uintptr(signalPixbufLoaderAreaPreparedId)))

	detail := signalPixbufLoaderAreaPreparedDetail{callback, handlerID}
	signalPixbufLoaderAreaPreparedMap[signalPixbufLoaderAreaPreparedId] = detail

	return signalPixbufLoaderAreaPreparedId
}

/*
DisconnectAreaPrepared disconnects a callback from the 'area-prepared' signal for the PixbufLoader.

The connectionID should be a value returned from a call to ConnectAreaPrepared.
*/
func (recv *PixbufLoader) DisconnectAreaPrepared(connectionID int) {
	signalPixbufLoaderAreaPreparedLock.Lock()
	defer signalPixbufLoaderAreaPreparedLock.Unlock()

	detail, exists := signalPixbufLoaderAreaPreparedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPixbufLoaderAreaPreparedMap, connectionID)
}

//export pixbufloader_areaPreparedHandler
func pixbufloader_areaPreparedHandler(_ *C.GObject, data C.gpointer) {
	signalPixbufLoaderAreaPreparedLock.RLock()
	defer signalPixbufLoaderAreaPreparedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalPixbufLoaderAreaPreparedMap[index].callback
	callback()
}

type signalPixbufLoaderAreaUpdatedDetail struct {
	callback  PixbufLoaderSignalAreaUpdatedCallback
	handlerID C.gulong
}

var signalPixbufLoaderAreaUpdatedId int
var signalPixbufLoaderAreaUpdatedMap = make(map[int]signalPixbufLoaderAreaUpdatedDetail)
var signalPixbufLoaderAreaUpdatedLock sync.RWMutex

// PixbufLoaderSignalAreaUpdatedCallback is a callback function for a 'area-updated' signal emitted from a PixbufLoader.
type PixbufLoaderSignalAreaUpdatedCallback func(x int32, y int32, width int32, height int32)

/*
ConnectAreaUpdated connects the callback to the 'area-updated' signal for the PixbufLoader.

The returned value represents the connection, and may be passed to DisconnectAreaUpdated to remove it.
*/
func (recv *PixbufLoader) ConnectAreaUpdated(callback PixbufLoaderSignalAreaUpdatedCallback) int {
	signalPixbufLoaderAreaUpdatedLock.Lock()
	defer signalPixbufLoaderAreaUpdatedLock.Unlock()

	signalPixbufLoaderAreaUpdatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PixbufLoader_signal_connect_area_updated(instance, C.gpointer(uintptr(signalPixbufLoaderAreaUpdatedId)))

	detail := signalPixbufLoaderAreaUpdatedDetail{callback, handlerID}
	signalPixbufLoaderAreaUpdatedMap[signalPixbufLoaderAreaUpdatedId] = detail

	return signalPixbufLoaderAreaUpdatedId
}

/*
DisconnectAreaUpdated disconnects a callback from the 'area-updated' signal for the PixbufLoader.

The connectionID should be a value returned from a call to ConnectAreaUpdated.
*/
func (recv *PixbufLoader) DisconnectAreaUpdated(connectionID int) {
	signalPixbufLoaderAreaUpdatedLock.Lock()
	defer signalPixbufLoaderAreaUpdatedLock.Unlock()

	detail, exists := signalPixbufLoaderAreaUpdatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPixbufLoaderAreaUpdatedMap, connectionID)
}

//export pixbufloader_areaUpdatedHandler
func pixbufloader_areaUpdatedHandler(_ *C.GObject, c_x C.gint, c_y C.gint, c_width C.gint, c_height C.gint, data C.gpointer) {
	signalPixbufLoaderAreaUpdatedLock.RLock()
	defer signalPixbufLoaderAreaUpdatedLock.RUnlock()

	x := int32(c_x)

	y := int32(c_y)

	width := int32(c_width)

	height := int32(c_height)

	index := int(uintptr(data))
	callback := signalPixbufLoaderAreaUpdatedMap[index].callback
	callback(x, y, width, height)
}

type signalPixbufLoaderClosedDetail struct {
	callback  PixbufLoaderSignalClosedCallback
	handlerID C.gulong
}

var signalPixbufLoaderClosedId int
var signalPixbufLoaderClosedMap = make(map[int]signalPixbufLoaderClosedDetail)
var signalPixbufLoaderClosedLock sync.RWMutex

// PixbufLoaderSignalClosedCallback is a callback function for a 'closed' signal emitted from a PixbufLoader.
type PixbufLoaderSignalClosedCallback func()

/*
ConnectClosed connects the callback to the 'closed' signal for the PixbufLoader.

The returned value represents the connection, and may be passed to DisconnectClosed to remove it.
*/
func (recv *PixbufLoader) ConnectClosed(callback PixbufLoaderSignalClosedCallback) int {
	signalPixbufLoaderClosedLock.Lock()
	defer signalPixbufLoaderClosedLock.Unlock()

	signalPixbufLoaderClosedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PixbufLoader_signal_connect_closed(instance, C.gpointer(uintptr(signalPixbufLoaderClosedId)))

	detail := signalPixbufLoaderClosedDetail{callback, handlerID}
	signalPixbufLoaderClosedMap[signalPixbufLoaderClosedId] = detail

	return signalPixbufLoaderClosedId
}

/*
DisconnectClosed disconnects a callback from the 'closed' signal for the PixbufLoader.

The connectionID should be a value returned from a call to ConnectClosed.
*/
func (recv *PixbufLoader) DisconnectClosed(connectionID int) {
	signalPixbufLoaderClosedLock.Lock()
	defer signalPixbufLoaderClosedLock.Unlock()

	detail, exists := signalPixbufLoaderClosedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPixbufLoaderClosedMap, connectionID)
}

//export pixbufloader_closedHandler
func pixbufloader_closedHandler(_ *C.GObject, data C.gpointer) {
	signalPixbufLoaderClosedLock.RLock()
	defer signalPixbufLoaderClosedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalPixbufLoaderClosedMap[index].callback
	callback()
}

type signalPixbufLoaderSizePreparedDetail struct {
	callback  PixbufLoaderSignalSizePreparedCallback
	handlerID C.gulong
}

var signalPixbufLoaderSizePreparedId int
var signalPixbufLoaderSizePreparedMap = make(map[int]signalPixbufLoaderSizePreparedDetail)
var signalPixbufLoaderSizePreparedLock sync.RWMutex

// PixbufLoaderSignalSizePreparedCallback is a callback function for a 'size-prepared' signal emitted from a PixbufLoader.
type PixbufLoaderSignalSizePreparedCallback func(width int32, height int32)

/*
ConnectSizePrepared connects the callback to the 'size-prepared' signal for the PixbufLoader.

The returned value represents the connection, and may be passed to DisconnectSizePrepared to remove it.
*/
func (recv *PixbufLoader) ConnectSizePrepared(callback PixbufLoaderSignalSizePreparedCallback) int {
	signalPixbufLoaderSizePreparedLock.Lock()
	defer signalPixbufLoaderSizePreparedLock.Unlock()

	signalPixbufLoaderSizePreparedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PixbufLoader_signal_connect_size_prepared(instance, C.gpointer(uintptr(signalPixbufLoaderSizePreparedId)))

	detail := signalPixbufLoaderSizePreparedDetail{callback, handlerID}
	signalPixbufLoaderSizePreparedMap[signalPixbufLoaderSizePreparedId] = detail

	return signalPixbufLoaderSizePreparedId
}

/*
DisconnectSizePrepared disconnects a callback from the 'size-prepared' signal for the PixbufLoader.

The connectionID should be a value returned from a call to ConnectSizePrepared.
*/
func (recv *PixbufLoader) DisconnectSizePrepared(connectionID int) {
	signalPixbufLoaderSizePreparedLock.Lock()
	defer signalPixbufLoaderSizePreparedLock.Unlock()

	detail, exists := signalPixbufLoaderSizePreparedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPixbufLoaderSizePreparedMap, connectionID)
}

//export pixbufloader_sizePreparedHandler
func pixbufloader_sizePreparedHandler(_ *C.GObject, c_width C.gint, c_height C.gint, data C.gpointer) {
	signalPixbufLoaderSizePreparedLock.RLock()
	defer signalPixbufLoaderSizePreparedLock.RUnlock()

	width := int32(c_width)

	height := int32(c_height)

	index := int(uintptr(data))
	callback := signalPixbufLoaderSizePreparedMap[index].callback
	callback(width, height)
}

// Blacklisted : gdk_pixbuf_loader_new

// Blacklisted : gdk_pixbuf_loader_new_with_type

// Blacklisted : gdk_pixbuf_loader_close

// Blacklisted : gdk_pixbuf_loader_get_animation

// Blacklisted : gdk_pixbuf_loader_get_pixbuf

// Blacklisted : gdk_pixbuf_loader_write

// PixbufSimpleAnim is a wrapper around the C record GdkPixbufSimpleAnim.
type PixbufSimpleAnim struct {
	native *C.GdkPixbufSimpleAnim
}

func PixbufSimpleAnimNewFromC(u unsafe.Pointer) *PixbufSimpleAnim {
	c := (*C.GdkPixbufSimpleAnim)(u)
	if c == nil {
		return nil
	}

	g := &PixbufSimpleAnim{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufSimpleAnim) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufSimpleAnim) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufSimpleAnim with another PixbufSimpleAnim, and returns true if they represent the same GObject.
func (recv *PixbufSimpleAnim) Equals(other *PixbufSimpleAnim) bool {
	return other.ToC() == recv.ToC()
}

// PixbufAnimation upcasts to *PixbufAnimation
func (recv *PixbufSimpleAnim) PixbufAnimation() *PixbufAnimation {
	return PixbufAnimationNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *PixbufSimpleAnim) Object() *gobject.Object {
	return recv.PixbufAnimation().Object()
}

// CastToWidget down casts any arbitrary Object to PixbufSimpleAnim.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufSimpleAnim.
func CastToPixbufSimpleAnim(object *gobject.Object) *PixbufSimpleAnim {
	return PixbufSimpleAnimNewFromC(object.ToC())
}

// Unsupported : PixbufSimpleAnimIter : no CType
