## connecting
To connect an object to a signal that it emits,
call one of the object's `Connect...` methods.

The connect methods all accept a handler argument
that is a function with a specific signature.
The first parameters is always the object that
the signal is for.
Some handlers have no additional parameters,
others have more parameters that provide event-specific
information.  

```go
button := gtk.ButtonNewWithLabel("a button")

// enable emitting of motion events
button.Widget().SetEvents(int32(gdk.GDK_POINTER_MOTION_MASK))

button.Widget().ConnectMotionNotifyEvent(
    func(widget *gtk.Widget, event *gdk.EventMotion) bool {
        fmt.Println(event.Y, event.Y)
        return false
    }
)
```

Remember that it may be necessary to
[upcast](casting.html)
to access the desired connect method.

## disconnecting
All of the `Connect...` methods return a handler id.
The id may be used later to disconnect the handler.

```go
    connectId := button.ConnectClicked(func(...) {
        ...
    })
    
    ...
    
    button.DisconnectClicked(connectId)
```

## notify signal for a single property
`ConnectNotifyProperty` connects a callback
to the `notify` signal
for a specific property of an Object.

This is an example of connecting to the notify signal
for a window's
[title property](https://developer.gnome.org/gtk3/stable/GtkWindow.html#GtkWindow--title).

```go
    window := gtk.WindowNew(...)

    // connect to the "notify::title" signal
    window.Object().ConnectNotifyProperty(
        "title", func(object *gobject.Object, pspec *gobject.ParamSpec,
    ) {
        fmt.Println("title property changed")
    })
```

The returned value represents the connection,
and may be passed to `DisconnectNotify` to remove it.

`ConnectNotifyProperty` is a special case,
and has no direct equivalent in the GObject library.
