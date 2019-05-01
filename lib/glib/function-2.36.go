// This is a generated file - DO NOT EDIT
// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import call "github.com/pekim/gobbi/lib/internal/call"

// Unsupported : g_close : return type :

// GetNumProcessors is a wrapper around the C function g_get_num_processors.
func GetNumProcessors() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(1611, &data)
	ret := data.Return.Uint32()

	return ret
}

// Unsupported : g_unix_fd_add : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Unsupported : g_unix_fd_add_full : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Unsupported : g_unix_fd_source_new : return type :
