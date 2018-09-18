// This is a generated file - DO NOT EDIT
// +build glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Datetime is a wrapper around the C record GDateTime.
type Datetime struct {
	native *C.GDateTime
}

func datetimeNewFromC(c *C.GDateTime) *Datetime {
	if c == nil {
		return nil
	}

	g := &Datetime{native: c}

	return g
}

// Timezone is a wrapper around the C record GTimeZone.
type Timezone struct {
	native *C.GTimeZone
}

func timezoneNewFromC(c *C.GTimeZone) *Timezone {
	if c == nil {
		return nil
	}

	g := &Timezone{native: c}

	return g
}
