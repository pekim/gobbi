// This is a generated file - DO NOT EDIT
// +build gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Error enumeration for #GdkGLContext.
type GLError C.GdkGLError

const (
	// OpenGL support is not available
	GDK_GL_ERROR_NOT_AVAILABLE GLError = 0

	// The requested visual format is not supported
	GDK_GL_ERROR_UNSUPPORTED_FORMAT GLError = 1

	// The requested profile is not supported
	GDK_GL_ERROR_UNSUPPORTED_PROFILE GLError = 2
)
