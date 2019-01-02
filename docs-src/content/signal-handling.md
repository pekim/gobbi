+++
title="Signal handling"
menu="main"
weight="9"
+++

## connecting
To connect an object to a signal that it emits,
call one of the object's `Connect...` methods.

The connect methods all accept a handler argument
that is a function with a specific signature.
Some handlers have no parameters,
others have one or more parameters that provide event-specific
information.  

```go
button := gtk.ButtonNewWithLabel("a button")
button.Widget().SetEvents(int32(gdk.GDK_POINTER_MOTION_MASK))
button.Widget().ConnectMotionNotifyEvent(
    func(event *gdk.EventMotion) bool {
        fmt.Println(event.Y, event.Y)
        return false
    })
```

Remember that it may be necessary to upcast to access the
desired connect method.

## disconnecting
All of the `Connect...` methods return a handler id.
The id may be used later to disconnect the handler.

```go
    connectId := button.ConnectClicked(func() {
        ...
    })
    
    ...
    
    button.DisconnectClicked(connectId)
```

## notify signal for a single property
`ConnectNotifyProperty` connects the callback
to the `notify` signal
for a specific property of an Object.

For example, to connect to the notify signal
for the "title" property of a window.

```go
    window := gtk.WindowNew(...)

    window.Object().ConnectNotifyProperty(
        "title", func(pspec *gobject.ParamSpec,
    ) {
        fmt.Println("title property changed")
    })
```

In that example it is the `notify::title` signal
that is being connected to.

The returned value represents the connection,
and may be passed to `DisconnectNotify` to remove it.

`ConnectNotifyProperty` is a special case,
and has no direct equivalent in the GObject library.
