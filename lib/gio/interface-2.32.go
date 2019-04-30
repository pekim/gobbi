// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// NetworkMonitor is a wrapper around the C record GNetworkMonitor.
type NetworkMonitor struct {
	native unsafe.Pointer
}

func NetworkMonitorNewFromC(u unsafe.Pointer) *NetworkMonitor {
	if u == nil {
		return nil
	}

	g := &NetworkMonitor{native: u}

	return g
}

func (recv *NetworkMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
