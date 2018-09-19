// This is a generated file - DO NOT EDIT
// +build glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// VariantDict is a wrapper around the C record GVariantDict.
type VariantDict struct {
	native *C.GVariantDict
}

func variantDictNewFromC(c *C.GVariantDict) *VariantDict {
	if c == nil {
		return nil
	}

	g := &VariantDict{native: c}

	return g
}

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant
