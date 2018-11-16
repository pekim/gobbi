+++
title="API"
+++

There is currently very little Godoc api documentation for gobbi.
As most apis are relatively direct mappings of the C library
APIs,
the [Gnome libraries documentation](https://developer.gnome.org/references)
should provide enough information to gain some
understanding of how to use the Go apis in gobbi.

## packages
There is one go package for each Gnome library.

## out params
Out parameters are mapped as return values in Go functions.
Inout parameters are not yet supported, but when they are
the will be exposed as both parameters and return values
in Go functions.

## inout params
Inout parameters are not yet supported.

## functions returning GError
Functions that have a GError out parameter, may return
a Go `error`.
The `error` (if not `nil`) will be a `glib.Error`,
which has fields containing details of the error.

## variadic functions
Variadic api functions are not supported.

### array functions
In many cases the library provides an equivalent
array based function that can be used instead.
The names of the array-based functions often
end in a 'v'.

The C arrays are represented as Go slices. 

In the case of such functions that accept two slice,
be certain that they are both of the same length.
The behaviour when they are not is undefined.

### single item functions
In other cases, instead of a variadic function
that would operate on multiple values at once,
there will be a supported function that operates
on a single item at a time. 

## cairo & pango
These libraries currently have very little support,
as the introspection date available for them is inadequate.
Support may be added later, with manually generated
bindings.
