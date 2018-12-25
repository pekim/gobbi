+++
title="Format functions"
menu="main"
weight="10"
+++

Some C library variadic functions accept trailing arguments
that include a printf style format string followed by
arguments to insert in to the format string.
[gtk_message_dialog_new](https://developer.gnome.org/gtk3/stable/GtkMessageDialog.html#gtk-message-dialog-new)
is an example of such a function.

The equivalent functions in gobbi accept
[fmt](https://golang.org/pkg/fmt/) package compatible arguments.
This is because `fmt.Sprint` is used to format a string
before passing it as the last argument
to the wrapped C function.

A much simplified representation of such a function might
help to make this clearer.
```go
func SomeRepresentativeFunction(
	a string, b int,
	format string, args ...interface{},
) {
	c_a := C.String(a)
	c_b := C.int(b)

	formattedString := fmt.Sprintf(format, args...)
	c_formattedString := C.String(formattedString)

	C.some_representative_function(c_a, c_b, c_formattedString)
}
```
