All calls to gtk functions should be performed on
the main thread.
To perform such calls from a goroutine,
use `glib.IdleAdd` to schedule invocation of a
callback function on the main thread. 

## `IdleAdd`

```go
glib.IdleAdd(func() bool {
    someLabel.SetText("some text")
    return glib.SOURCE_REMOVE
})
```

Return `glib.SOURCE_REMOVE` to ensure the function is not
called again next time the main loop is idle.
Alternatively return `glib.SOURCE_CONTINUE` to have
the function called again.

## `IdleAddOnce`

Alternatively, `glib.IdleAddOnce`
schedules a _single_ invocation of a
callback function on the main thread.
Unlike `glib.IdleAdd` no value needs to be returned
to avoid subsequent invocations of the callback.


```go
glib.IdleAddOnce(func() {
    someLabel.SetText("some text")
})
```

`IdleAddOnce` is a convenience provided by gobbi.
It has no direct equivalent in the glib C library.
