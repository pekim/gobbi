// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

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

// Flags to define the behaviour of a #GSubprocess.
//
// Note that the default for stdin is to redirect from `/dev/null`.  For
// stdout and stderr the default are for them to inherit the
// corresponding descriptor from the calling process.
//
// Note that it is a programmer error to mix 'incompatible' flags.  For
// example, you may not request both %G_SUBPROCESS_FLAGS_STDOUT_PIPE and
// %G_SUBPROCESS_FLAGS_STDOUT_SILENCE.
type SubprocessFlags C.GSubprocessFlags

const (
	// No flags.
	SUBPROCESS_FLAGS_NONE SubprocessFlags = 0
	// create a pipe for the stdin of the
	// spawned process that can be accessed with
	// g_subprocess_get_stdin_pipe().
	SUBPROCESS_FLAGS_STDIN_PIPE SubprocessFlags = 1
	// stdin is inherited from the
	// calling process.
	SUBPROCESS_FLAGS_STDIN_INHERIT SubprocessFlags = 2
	// create a pipe for the stdout of the
	// spawned process that can be accessed with
	// g_subprocess_get_stdout_pipe().
	SUBPROCESS_FLAGS_STDOUT_PIPE SubprocessFlags = 4
	// silence the stdout of the spawned
	// process (ie: redirect to `/dev/null`).
	SUBPROCESS_FLAGS_STDOUT_SILENCE SubprocessFlags = 8
	// create a pipe for the stderr of the
	// spawned process that can be accessed with
	// g_subprocess_get_stderr_pipe().
	SUBPROCESS_FLAGS_STDERR_PIPE SubprocessFlags = 16
	// silence the stderr of the spawned
	// process (ie: redirect to `/dev/null`).
	SUBPROCESS_FLAGS_STDERR_SILENCE SubprocessFlags = 32
	// merge the stderr of the spawned
	// process with whatever the stdout happens to be.  This is a good way
	// of directing both streams to a common log file, for example.
	SUBPROCESS_FLAGS_STDERR_MERGE SubprocessFlags = 64
	// spawned processes will inherit the
	// file descriptors of their parent, unless those descriptors have
	// been explicitly marked as close-on-exec.  This flag has no effect
	// over the "standard" file descriptors (stdin, stdout, stderr).
	SUBPROCESS_FLAGS_INHERIT_FDS SubprocessFlags = 128
)
