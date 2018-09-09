package gobject

// #include <glib-object.h>
// #include <stdlib.h>
import "C"

const PARAM_MASK int = 255
const PARAM_STATIC_STRINGS int = 224
const PARAM_USER_SHIFT int = 8
const SIGNAL_FLAGS_MASK int = 511
const SIGNAL_MATCH_MASK int = 63
const TYPE_FLAG_RESERVED_ID_BIT int = 1
const TYPE_FUNDAMENTAL_MAX int = 255
const TYPE_FUNDAMENTAL_SHIFT int = 2
const TYPE_RESERVED_BSE_FIRST int = 32
const TYPE_RESERVED_BSE_LAST int = 48
const TYPE_RESERVED_GLIB_FIRST int = 22
const TYPE_RESERVED_GLIB_LAST int = 31
const TYPE_RESERVED_USER_FIRST int = 49
const VALUE_COLLECT_FORMAT_MAX_LENGTH int = 8
const VALUE_NOCOPY_CONTENTS int = 134217728
