/*

Package c, and packages in child directories, provide low level bindings
to libraries' C apis.
The exported functions do little more than wrap C api functions in Go functions.

This package is intended for internal use within gobbi.

The exported Go functions have the same number of parameters and return value as
the C functions that they wrap.
The parameters and reurn values as transformed only as much as necessary
to map between C types and Go Types.

simple types - Integer like types and floats type simply use Go conversion to transform.

strings - Strings are converted back and forth between Go strings and null terminated strings.

records, classes, interfaces, unions - All of these types are represented in Go as unsafe.Pointer.

arrays - Arrays are converted between Go slices and C array-like representations (pointers to pointers).

*/
package c
