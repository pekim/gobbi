# gobbi

**gob**ject **bi**ndings

[![Build Status](https://dev.azure.com/mikepilsbury/gobbi/_apis/build/status/pekim.gobbi?branchName=master)](https://dev.azure.com/mikepilsbury/gobbi/_build/latest?definitionId=1&branchName=master)
[![builds.sr.ht status](https://builds.sr.ht/~pekim/gobbi.svg)](https://builds.sr.ht/~pekim/gobbi?)

gobbi is a set of generated Go bindings for gobject based libraries.
There are bindings, with varying degrees of completeness,
for gtk, gdk, gdkpixbuf, gio, gobject, and glib.

The intention is to provide the bindings required
for writing gtk applications.
So the focus is on supporting functions in the higher level
libraries such as gtk, gdk, gdkpixbuf and gio.
Much of the functionality provided by the lower level
libraries, such as glib and gobject, is adequately supported
in Go core packages.     

## example use
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
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	window.Widget().ConnectDestroy(func(_ *gtk.Widget) {
		gtk.MainQuit()
	})
	window.Widget().ShowAll()

	gtk.Main()
}
```

```bash
go get github.com/pekim/gobbi
go run -tags "glib_2.48 gdk_3.4 gdkpixbuf_2.32 gio_2.32 gtk_3.18" gobbi-simple.go
```

## documentation
https://pekim.github.io/gobbi/

At a minimum it's probably a good idea to read at least the
[Getting started](https://pekim.github.io/gobbi/getting-started.html)
and
[Application lifecycle](https://pekim.github.io/gobbi/application-lifecycle.html)
pages.

## code generation
The code that implements most of the APIs is generated
from the GObject Introspection (gir) XML files in the
[internal/gir-files](internal/gir-files) directory.
The generation code is in the
[internal/generate](internal/generate) package.

All of the generated API code,
and a very small number of manually maintained files,
is in the subdirectories of the
[lib](lib) directory.
 
## status
The library should be perfectly servicable and complete enough
to build most applications.

Some simple example applications have been created,
and may be found in the [example](./example) directory.
It is possible that for more complex applications
some necessary apis may be missing. 

The API is reasonably stable,
although some changes are certainly possible. 

The generation code is rather messy in some areas,
and needs tidying up.
The generated code is simple, and fairly readable.

## gotk3
[gotk3](https://github.com/gotk3/gotk3)
is a more mature library than gobbi.
It might be a safer choice than gobbi for those who
are cautious.
However gobbi should work just as
well in most cases.

gobbi supports a larger number of apis than gotk3.
However gotk3 supports some apis that gobbi does not,
and vice versa.

gobbi is more type safe in some areas,
notably the callbacks connected to signals.
