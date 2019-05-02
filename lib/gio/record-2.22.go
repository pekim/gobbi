// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"C"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// Equals compares this AsyncInitableIface with another AsyncInitableIface, and returns true if they represent the same GObject.
func (recv *AsyncInitableIface) Equals(other *AsyncInitableIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this InitableIface with another InitableIface, and returns true if they represent the same GObject.
func (recv *InitableIface) Equals(other *InitableIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this InputVector with another InputVector, and returns true if they represent the same GObject.
func (recv *InputVector) Equals(other *InputVector) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this OutputVector with another OutputVector, and returns true if they represent the same GObject.
func (recv *OutputVector) Equals(other *OutputVector) bool {
	return other.ToC() == recv.ToC()
}

// SrvTargetListSort is a wrapper around the C function g_srv_target_list_sort.
func SrvTargetListSort(targets *glib.List) *glib.List {
	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	retC := C.g_srv_target_list_sort(c_targets)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function g_srv_target_copy.
func (recv *SrvTarget) Copy() *SrvTarget {
	retC := C.g_srv_target_copy((*C.GSrvTarget)(recv.native))
	retGo := SrvTargetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_srv_target_free.
func (recv *SrvTarget) Free() {
	C.g_srv_target_free((*C.GSrvTarget)(recv.native))

	return
}

// GetHostname is a wrapper around the C function g_srv_target_get_hostname.
func (recv *SrvTarget) GetHostname() string {
	retC := C.g_srv_target_get_hostname((*C.GSrvTarget)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPort is a wrapper around the C function g_srv_target_get_port.
func (recv *SrvTarget) GetPort() uint16 {
	retC := C.g_srv_target_get_port((*C.GSrvTarget)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetPriority is a wrapper around the C function g_srv_target_get_priority.
func (recv *SrvTarget) GetPriority() uint16 {
	retC := C.g_srv_target_get_priority((*C.GSrvTarget)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetWeight is a wrapper around the C function g_srv_target_get_weight.
func (recv *SrvTarget) GetWeight() uint16 {
	retC := C.g_srv_target_get_weight((*C.GSrvTarget)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}
