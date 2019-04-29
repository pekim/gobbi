// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// Application is a wrapper around the C record GApplication.
type Application struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// SimpleActionGroup is a wrapper around the C record GSimpleActionGroup.
type SimpleActionGroup struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// TlsCertificate is a wrapper around the C record GTlsCertificate.
type TlsCertificate struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// TlsConnection is a wrapper around the C record GTlsConnection.
type TlsConnection struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}
