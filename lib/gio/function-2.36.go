// This is a generated file - DO NOT EDIT
// +build gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import call "github.com/pekim/gobbi/lib/internal/call"

// Unsupported : g_dbus_address_escape_value : return type :

// Unsupported : g_file_new_for_commandline_arg_and_cwd : return type :

// NetworkingInit is a wrapper around the C function g_networking_init.
func NetworkingInit() {
	data := call.Data{Return: call.Return{Type: call.RT_VOID}}
	call.Function(2213, &data)
	return
}
