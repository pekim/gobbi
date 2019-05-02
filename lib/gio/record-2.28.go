// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

// Equals compares this ActionGroupInterface with another ActionGroupInterface, and returns true if they represent the same GObject.
func (recv *ActionGroupInterface) Equals(other *ActionGroupInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ActionInterface with another ActionInterface, and returns true if they represent the same GObject.
func (recv *ActionInterface) Equals(other *ActionInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ApplicationClass with another ApplicationClass, and returns true if they represent the same GObject.
func (recv *ApplicationClass) Equals(other *ApplicationClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ApplicationCommandLineClass with another ApplicationCommandLineClass, and returns true if they represent the same GObject.
func (recv *ApplicationCommandLineClass) Equals(other *ApplicationCommandLineClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this PollableInputStreamInterface with another PollableInputStreamInterface, and returns true if they represent the same GObject.
func (recv *PollableInputStreamInterface) Equals(other *PollableInputStreamInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this PollableOutputStreamInterface with another PollableOutputStreamInterface, and returns true if they represent the same GObject.
func (recv *PollableOutputStreamInterface) Equals(other *PollableOutputStreamInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this TlsBackendInterface with another TlsBackendInterface, and returns true if they represent the same GObject.
func (recv *TlsBackendInterface) Equals(other *TlsBackendInterface) bool {
	return other.ToC() == recv.ToC()
}
