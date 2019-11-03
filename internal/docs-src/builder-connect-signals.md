## signals in UI definitions
User interfaces may be defined by in XML documents.
Such documents are typically edited with Glade.

Signals, along with a named handler, maybe appear in a definition. 

```xml
<interface>
  ...
    <object class="GtkButton" id="ok_button">
      <signal name="clicked" handler="ok_button_clicked"/>
    </object>
  ...
</interface>
```

The `handler` attribute value provides a name
that can be used to associate the signal with a handler function.  


## connecting
To connect signals from a UI definition,
call `Builder`'s `ConnectSignals` method.
The single argument is a map of handler names to handler functions

The map keys should match `handler` attribute values in
the UI definition.

The handler function signatures must match the signal's
equivalent `Connect...` function.  

## example
```go
builder := gtk.BuilderNewFromString(`
    <interface>
      <object class="GtkButton" id="ok_button">
        <signal name="clicked" handler="ok_button_clicked"/>
        <signal name="draw" handler="draw"/>
      </object>
    </interface>
`)

err := builder.ConnectSignals(map[string]interface{}{
    "ok_button_clicked": func(button *gtk.Button) {
        // handle the click
    },
    "draw": func(widget *gtk.Widget, cr *cairo.Context) bool {
        // draw something in the context
        return false
    },
})
```

See the tests in
[gtkbuilder-connect-signals_test.go](https://github.com/pekim/gobbi/blob/master/internal/test/gtk/gtkbuilder-connect-signals_test.go)
for more examples
including some error scenarios. 

## errors
There are a number of errors that may occur while connecting signals.
* A key in the map passed to `ConnectSignals` may not match any
`handler` attribute value in the UI definition.
* A handler function's signature may be incorrect.
* A signal name may be valid, but the signal may not be supported by gobbi.

Any of these will result in an `error` being returned. 
