// +build glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// GetEnviron is a wrapper around the C function g_get_environ.
func GetEnviron() {}

// GetLocaleVariants is a wrapper around the C function g_get_locale_variants.
func GetLocaleVariants(locale string) {}

// GetMonotonicTime is a wrapper around the C function g_get_monotonic_time.
func GetMonotonicTime() {}

// GetRealTime is a wrapper around the C function g_get_real_time.
func GetRealTime() {}

// GetUserRuntimeDir is a wrapper around the C function g_get_user_runtime_dir.
func GetUserRuntimeDir() {}
