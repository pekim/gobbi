// This is a generated file - DO NOT EDIT
// +build gdk_3.8 gdk_3.10 gdk_3.16 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

type FullscreenMode C.GdkFullscreenMode

const (
	GDK_FULLSCREEN_ON_CURRENT_MONITOR FullscreenMode = 0
	GDK_FULLSCREEN_ON_ALL_MONITORS    FullscreenMode = 1
)
