# `gobject` Constants

## `PARAM_MASK`

Mask containing the bits of #GParamSpec.flags which are reserved for GLib.

C - `G_PARAM_MASK`

## `PARAM_STATIC_STRINGS`

#GParamFlags value alias for %G_PARAM_STATIC_NAME | %G_PARAM_STATIC_NICK | %G_PARAM_STATIC_BLURB.

Since 2.13.0

C - `G_PARAM_STATIC_STRINGS`

## `PARAM_USER_SHIFT`

Minimum shift count to be used for user defined flags, to be stored in
&num;GParamSpec.flags. The maximum allowed is 10.

C - `G_PARAM_USER_SHIFT`

## `SIGNAL_FLAGS_MASK`

A mask for all #GSignalFlags bits.

C - `G_SIGNAL_FLAGS_MASK`

## `SIGNAL_MATCH_MASK`

A mask for all #GSignalMatchType bits.

C - `G_SIGNAL_MATCH_MASK`

## `TYPE_FUNDAMENTAL_MAX`

An integer constant that represents the number of identifiers reserved
for types that are assigned at compile-time.

C - `G_TYPE_FUNDAMENTAL_MAX`

## `TYPE_FUNDAMENTAL_SHIFT`

Shift value used in converting numbers to type IDs.

C - `G_TYPE_FUNDAMENTAL_SHIFT`

## `TYPE_RESERVED_BSE_FIRST`

First fundamental type number to create a new fundamental type id with
G_TYPE_MAKE_FUNDAMENTAL() reserved for BSE.

C - `G_TYPE_RESERVED_BSE_FIRST`

## `TYPE_RESERVED_BSE_LAST`

Last fundamental type number reserved for BSE.

C - `G_TYPE_RESERVED_BSE_LAST`

## `TYPE_RESERVED_GLIB_FIRST`

First fundamental type number to create a new fundamental type id with
G_TYPE_MAKE_FUNDAMENTAL() reserved for GLib.

C - `G_TYPE_RESERVED_GLIB_FIRST`

## `TYPE_RESERVED_GLIB_LAST`

Last fundamental type number reserved for GLib.

C - `G_TYPE_RESERVED_GLIB_LAST`

## `TYPE_RESERVED_USER_FIRST`

First available fundamental type number to create new fundamental
type id with G_TYPE_MAKE_FUNDAMENTAL().

C - `G_TYPE_RESERVED_USER_FIRST`

## `VALUE_NOCOPY_CONTENTS`

If passed to G_VALUE_COLLECT(), allocated data won't be copied
but used verbatim. This does not affect ref-counted types like
objects.

C - `G_VALUE_NOCOPY_CONTENTS`

