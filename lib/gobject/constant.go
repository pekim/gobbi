// This is a generated file - DO NOT EDIT

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Mask containing the bits of #GParamSpec.flags which are reserved for GLib.
const PARAM_MASK int = C.G_PARAM_MASK

// #GParamFlags value alias for %G_PARAM_STATIC_NAME | %G_PARAM_STATIC_NICK | %G_PARAM_STATIC_BLURB.
//
// Since 2.13.0
const PARAM_STATIC_STRINGS int = C.G_PARAM_STATIC_STRINGS

// Minimum shift count to be used for user defined flags, to be stored in
// #GParamSpec.flags. The maximum allowed is 10.
const PARAM_USER_SHIFT int = C.G_PARAM_USER_SHIFT

// A mask for all #GSignalFlags bits.
const SIGNAL_FLAGS_MASK int = C.G_SIGNAL_FLAGS_MASK

// A mask for all #GSignalMatchType bits.
const SIGNAL_MATCH_MASK int = C.G_SIGNAL_MATCH_MASK

// Unsupported : type GLib.Type for TYPE_FLAG_RESERVED_ID_BIT

// An integer constant that represents the number of identifiers reserved
// for types that are assigned at compile-time.
const TYPE_FUNDAMENTAL_MAX int = C.G_TYPE_FUNDAMENTAL_MAX

// Shift value used in converting numbers to type IDs.
const TYPE_FUNDAMENTAL_SHIFT int = C.G_TYPE_FUNDAMENTAL_SHIFT

// First fundamental type number to create a new fundamental type id with
// G_TYPE_MAKE_FUNDAMENTAL() reserved for BSE.
const TYPE_RESERVED_BSE_FIRST int = C.G_TYPE_RESERVED_BSE_FIRST

// Last fundamental type number reserved for BSE.
const TYPE_RESERVED_BSE_LAST int = C.G_TYPE_RESERVED_BSE_LAST

// First fundamental type number to create a new fundamental type id with
// G_TYPE_MAKE_FUNDAMENTAL() reserved for GLib.
const TYPE_RESERVED_GLIB_FIRST int = C.G_TYPE_RESERVED_GLIB_FIRST

// Last fundamental type number reserved for GLib.
const TYPE_RESERVED_GLIB_LAST int = C.G_TYPE_RESERVED_GLIB_LAST

// First available fundamental type number to create new fundamental
// type id with G_TYPE_MAKE_FUNDAMENTAL().
const TYPE_RESERVED_USER_FIRST int = C.G_TYPE_RESERVED_USER_FIRST

// Blacklisted : VALUE_COLLECT_FORMAT_MAX_LENGTH

// If passed to G_VALUE_COLLECT(), allocated data won't be copied
// but used verbatim. This does not affect ref-counted types like
// objects.
const VALUE_NOCOPY_CONTENTS int = C.G_VALUE_NOCOPY_CONTENTS
