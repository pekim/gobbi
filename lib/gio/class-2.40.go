// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// AppInfoMonitor is a wrapper around the C record GAppInfoMonitor.
type AppInfoMonitor struct {
	native unsafe.Pointer
}

func AppInfoMonitorNewFromC(u unsafe.Pointer) *AppInfoMonitor {
	if u == nil {
		return nil
	}

	g := &AppInfoMonitor{native: u}

	return g
}

func (recv *AppInfoMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Notification is a wrapper around the C record GNotification.
type Notification struct {
	native unsafe.Pointer
}

func NotificationNewFromC(u unsafe.Pointer) *Notification {
	if u == nil {
		return nil
	}

	g := &Notification{native: u}

	return g
}

func (recv *Notification) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Subprocess is a wrapper around the C record GSubprocess.
type Subprocess struct {
	native unsafe.Pointer
}

func SubprocessNewFromC(u unsafe.Pointer) *Subprocess {
	if u == nil {
		return nil
	}

	g := &Subprocess{native: u}

	return g
}

func (recv *Subprocess) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SubprocessLauncher is a wrapper around the C record GSubprocessLauncher.
type SubprocessLauncher struct {
	native unsafe.Pointer
}

func SubprocessLauncherNewFromC(u unsafe.Pointer) *SubprocessLauncher {
	if u == nil {
		return nil
	}

	g := &SubprocessLauncher{native: u}

	return g
}

func (recv *SubprocessLauncher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
