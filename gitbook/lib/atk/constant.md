# `atk` Constants

## `BINARY_AGE`

Like atk_get_binary_age(), but from the headers used at
application compile time, rather than from the library linked
against at application run time.

C - `ATK_BINARY_AGE`

## `INTERFACE_AGE`

Like atk_get_interface_age(), but from the headers used at
application compile time, rather than from the library linked
against at application run time.

C - `ATK_INTERFACE_AGE`

## `MAJOR_VERSION`

Like atk_get_major_version(), but from the headers used at
application compile time, rather than from the library linked
against at application run time.

C - `ATK_MAJOR_VERSION`

## `MICRO_VERSION`

Like atk_get_micro_version(), but from the headers used at
application compile time, rather than from the library linked
against at application run time.

C - `ATK_MICRO_VERSION`

## `MINOR_VERSION`

Like atk_get_minor_version(), but from the headers used at
application compile time, rather than from the library linked
against at application run time.

C - `ATK_MINOR_VERSION`

## `VERSION_MIN_REQUIRED`

A macro that should be defined by the user prior to including
the atk/atk.h header.
The definition should be one of the predefined ATK version
macros: %ATK_VERSION_2_12, %ATK_VERSION_2_14,...

This macro defines the earliest version of ATK that the package is
required to be able to compile against.

If the compiler is configured to warn about the use of deprecated
functions, then using functions that were deprecated in version
%ATK_VERSION_MIN_REQUIRED or earlier will cause warnings (but
using functions deprecated in later releases will not).

C - `ATK_VERSION_MIN_REQUIRED`

