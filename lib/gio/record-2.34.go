// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Gets the description for @key.
//
// If no description has been provided in the schema for @key, returns
// %NULL.
//
// The description can be one sentence to several paragraphs in length.
// Paragraphs are delimited with a double newline.  Descriptions can be
// translated and the value returned from this function is is the
// current locale.
//
// This function is slow.  The summary and description information for
// the schemas is not stored in the compiled schema database so this
// function has to parse all of the source XML files in the schema
// directory.
/*

C function : g_settings_schema_key_get_description
*/
func (recv *SettingsSchemaKey) GetDescription() string {
	retC := C.g_settings_schema_key_get_description((*C.GSettingsSchemaKey)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the summary for @key.
//
// If no summary has been provided in the schema for @key, returns
// %NULL.
//
// The summary is a short description of the purpose of the key; usually
// one short sentence.  Summaries can be translated and the value
// returned from this function is is the current locale.
//
// This function is slow.  The summary and description information for
// the schemas is not stored in the compiled schema database so this
// function has to parse all of the source XML files in the schema
// directory.
/*

C function : g_settings_schema_key_get_summary
*/
func (recv *SettingsSchemaKey) GetSummary() string {
	retC := C.g_settings_schema_key_get_summary((*C.GSettingsSchemaKey)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Guesses the symbolic icon of a Unix mount point.
/*

C function : g_unix_mount_point_guess_symbolic_icon
*/
func (recv *UnixMountPoint) GuessSymbolicIcon() *Icon {
	retC := C.g_unix_mount_point_guess_symbolic_icon((*C.GUnixMountPoint)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}
