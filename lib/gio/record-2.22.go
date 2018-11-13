// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// Provides an interface for asynchronous initializing object such that
// initialization may fail.
/*

C record/class : GAsyncInitableIface
*/
type AsyncInitableIface struct {
	native *C.GAsyncInitableIface
	// g_iface : record
	// no type for init_async
	// no type for init_finish
}

func AsyncInitableIfaceNewFromC(u unsafe.Pointer) *AsyncInitableIface {
	c := (*C.GAsyncInitableIface)(u)
	if c == nil {
		return nil
	}

	g := &AsyncInitableIface{native: c}

	return g
}

func (recv *AsyncInitableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Provides an interface for initializing object such that initialization
// may fail.
/*

C record/class : GInitableIface
*/
type InitableIface struct {
	native *C.GInitableIface
	// g_iface : record
	// no type for init
}

func InitableIfaceNewFromC(u unsafe.Pointer) *InitableIface {
	c := (*C.GInitableIface)(u)
	if c == nil {
		return nil
	}

	g := &InitableIface{native: c}

	return g
}

func (recv *InitableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Structure used for scatter/gather data input.
// You generally pass in an array of #GInputVectors
// and the operation will store the read data starting in the
// first buffer, switching to the next as needed.
/*

C record/class : GInputVector
*/
type InputVector struct {
	native *C.GInputVector
	Buffer uintptr
	Size   uint64
}

func InputVectorNewFromC(u unsafe.Pointer) *InputVector {
	c := (*C.GInputVector)(u)
	if c == nil {
		return nil
	}

	g := &InputVector{
		Buffer: (uintptr)(c.buffer),
		Size:   (uint64)(c.size),
		native: c,
	}

	return g
}

func (recv *InputVector) ToC() unsafe.Pointer {
	recv.native.buffer =
		(C.gpointer)(recv.Buffer)
	recv.native.size =
		(C.gsize)(recv.Size)

	return (unsafe.Pointer)(recv.native)
}

// Structure used for scatter/gather data output.
// You generally pass in an array of #GOutputVectors
// and the operation will use all the buffers as if they were
// one buffer.
/*

C record/class : GOutputVector
*/
type OutputVector struct {
	native *C.GOutputVector
	Buffer uintptr
	Size   uint64
}

func OutputVectorNewFromC(u unsafe.Pointer) *OutputVector {
	c := (*C.GOutputVector)(u)
	if c == nil {
		return nil
	}

	g := &OutputVector{
		Buffer: (uintptr)(c.buffer),
		Size:   (uint64)(c.size),
		native: c,
	}

	return g
}

func (recv *OutputVector) ToC() unsafe.Pointer {
	recv.native.buffer =
		(C.gconstpointer)(recv.Buffer)
	recv.native.size =
		(C.gsize)(recv.Size)

	return (unsafe.Pointer)(recv.native)
}

// Creates a new #GSrvTarget with the given parameters.
//
// You should not need to use this; normally #GSrvTargets are
// created by #GResolver.
/*

C function : g_srv_target_new
*/
func SrvTargetNew(hostname string, port uint16, priority uint16, weight uint16) *SrvTarget {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_port := (C.guint16)(port)

	c_priority := (C.guint16)(priority)

	c_weight := (C.guint16)(weight)

	retC := C.g_srv_target_new(c_hostname, c_port, c_priority, c_weight)
	retGo := SrvTargetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copies @target
/*

C function : g_srv_target_copy
*/
func (recv *SrvTarget) Copy() *SrvTarget {
	retC := C.g_srv_target_copy((*C.GSrvTarget)(recv.native))
	retGo := SrvTargetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees @target
/*

C function : g_srv_target_free
*/
func (recv *SrvTarget) Free() {
	C.g_srv_target_free((*C.GSrvTarget)(recv.native))

	return
}

// Gets @target's hostname (in ASCII form; if you are going to present
// this to the user, you should use g_hostname_is_ascii_encoded() to
// check if it contains encoded Unicode segments, and use
// g_hostname_to_unicode() to convert it if it does.)
/*

C function : g_srv_target_get_hostname
*/
func (recv *SrvTarget) GetHostname() string {
	retC := C.g_srv_target_get_hostname((*C.GSrvTarget)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets @target's port
/*

C function : g_srv_target_get_port
*/
func (recv *SrvTarget) GetPort() uint16 {
	retC := C.g_srv_target_get_port((*C.GSrvTarget)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Gets @target's priority. You should not need to look at this;
// #GResolver already sorts the targets according to the algorithm in
// RFC 2782.
/*

C function : g_srv_target_get_priority
*/
func (recv *SrvTarget) GetPriority() uint16 {
	retC := C.g_srv_target_get_priority((*C.GSrvTarget)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Gets @target's weight. You should not need to look at this;
// #GResolver already sorts the targets according to the algorithm in
// RFC 2782.
/*

C function : g_srv_target_get_weight
*/
func (recv *SrvTarget) GetWeight() uint16 {
	retC := C.g_srv_target_get_weight((*C.GSrvTarget)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}
