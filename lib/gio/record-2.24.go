// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

// Equals compares this ConverterIface with another ConverterIface, and returns true if they represent the same GObject.
func (recv *ConverterIface) Equals(other *ConverterIface) bool {
	return other.ToC() == recv.ToC()
}
