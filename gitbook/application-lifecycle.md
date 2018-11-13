# Application lifecycle

## main thread
The main function, and therefore the main loop,
should run on the main thread. 

```go
func init() {
    // Ensure that the ui's main thread is
    // locked to the main thread.
    runtime.LockOSThread()
}
```

## initialisation
Call `gtk.Init` (or `gtk.InitChack`) before calling
any other gtk functions.
Typically the program's command-line arguments will
be passed to `gtk.Init`.

```go
gtk.Init(os.Args)
```

## window creation
Most applications will create a window when initialising.

```go
window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
``` 

## termination
If an application has only one window,
then it might terminate when the window is closed.

```go
window.Widget().ConnectDestroy(func() {
    gtk.MainQuit()
})
``` 

This may be written more succinctly, like this.

```go
window.Widget().ConnectDestroy(gtk.MainQuit)
``` 

## main loop
To run the main loop, call `gtk.Main`.
The function will not return from the main loop
until `gtk.MainQuit` is called.

```go
gtk.Main()
```

## minimal application
Putting all of the above together yields a minimal
application.
A window is created,
and the application will exit when the window is closed.

```go
package main

import (
	"github.com/pekim/gobbi/lib/gtk"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	gtk.Init(os.Args)

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
```

## gtk.Application
A preferred approach to gtk application and window
management is to use `gtk.Application`.
The [gtk_application](https://github.com/pekim/gobbi/blob/master/example/gtk_application/main.go)
example illustrates this.
