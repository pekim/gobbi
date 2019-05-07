Some C library functions are 
[variadic](https://en.wikipedia.org/wiki/Variadic_function#Example_in_C).
Unfortunately cgo does not support calling variadic C functions. 

## format functions
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

A much simplified representation of such a
gobbi function might help to make this clearer.
```go
package gtk

// a gobbi function
func SomeFormatFunction(
	a string, b int,
	format string, args ...interface{},
) {
	c_a := C.String(a)
	c_b := C.int(b)

	formattedString := fmt.Sprintf(format, args...)
	c_formattedString := C.String(formattedString)

	C.some_format_function(c_a, c_b, c_formattedString)
}

// ----------------------------- //

// some application code
func aFunction () {
    gtk.SomeFormatFunction(
        "a string", 123,
        "the %s is %d", "answer", 42,
    )
}
```

## unsupported variadic functions
Many variadic functions are not supported by gobbi
because it would be impractical to automatically
generate code that would support them.

### array functions
In some cases there is an alternative function that accepts an array
as the final argument instead.
For example
[gtk_list_store_new](https://developer.gnome.org/gtk3/stable/GtkListStore.html#gtk-list-store-new)
is not supported, but 
[gtk_list_store_newv](https://developer.gnome.org/gtk3/stable/GtkListStore.html#gtk-list-store-newv)
is supported as [gtk.ListStoreNewv](https://godoc.org/github.com/pekim/gobbi/lib/gtk#ListStoreNewv).

### multiple functions 
There are other variadac functions that may not have an array
equivalent version, but their functionality
can be achieved through a combination of functions.

An example of this is the unsupported 
[gtk_text_buffer_create_tag](https://developer.gnome.org/gtk3/stable/GtkTextBuffer.html#gtk-text-buffer-create-tag)
function.
A work around might be code like this.

```go
    tag := gtk.TextTagNew("my_tag_name")
    tagTable := gtk.TextTagTableNew()
    tagTable.Add(tag)
    buffer := gtk.TextBufferNew(tagTable)

    editable := gobject.ValueNew(gobject.TYPE_BOOLEAN)
    editable.SetBoolean(false)
    tag.Object().SetProperty("property-name", editable)

    indent := gobject.ValueNew(gobject.TYPE_INT)
    indent.SetInt(4)
    tag.Object().SetProperty("indent", indent)
```
