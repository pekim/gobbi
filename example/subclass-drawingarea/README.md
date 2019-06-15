# subclassing example

This example demonstrates an approach for
subclassing a gtk widget.
A subclass of GtkDrawingArea is registered,
and two virtual functions are implemented.

The `draw` virtual function calls a Go function
that marshals the arguments in to Go objects,
and calls another function to draw in a cairo context.
The same result could have been achieved by
connecting to the `draw` signal instead,
(The [custom-drawing](https://github.com/pekim/gobbi/blob/master/example/custom-drawing/main.go)
example uses that approach.)
however the point of this example is to
illustrate class derivation and
the overriding of virtual functions.

The `adjust_size_request` virtual function's
implementation is trivial,
and is implemented entirely in C.

For more background about subclassing with gobbi see
[subclassing](https://pekim.github.io/gobbi/subclassing)
in the documentation.
