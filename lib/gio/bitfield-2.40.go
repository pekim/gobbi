// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

type SubprocessFlags C.GSubprocessFlags

const (
	SUBPROCESS_FLAGS_NONE           SubprocessFlags = 0
	SUBPROCESS_FLAGS_STDIN_PIPE     SubprocessFlags = 1
	SUBPROCESS_FLAGS_STDIN_INHERIT  SubprocessFlags = 2
	SUBPROCESS_FLAGS_STDOUT_PIPE    SubprocessFlags = 4
	SUBPROCESS_FLAGS_STDOUT_SILENCE SubprocessFlags = 8
	SUBPROCESS_FLAGS_STDERR_PIPE    SubprocessFlags = 16
	SUBPROCESS_FLAGS_STDERR_SILENCE SubprocessFlags = 32
	SUBPROCESS_FLAGS_STDERR_MERGE   SubprocessFlags = 64
	SUBPROCESS_FLAGS_INHERIT_FDS    SubprocessFlags = 128
)
