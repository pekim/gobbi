+++
title="Goroutines"
menu="main"
weight="6"
+++

# goroutines

All calls to gtk functions should be performed on
the main thread.
To perform such a call from a goroutine,
use `glib.IdleAdd` to schedule invocation of a
callback function on the main thread. 

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
