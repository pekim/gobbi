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
