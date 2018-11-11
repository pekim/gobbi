# gobbi

**gob**ject **bi**ndings

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

## code generation
The code that implements most of the APIs is generated
from the GObject Introspection (gir) XML files in the
[internal/gir-files](internal/gir-files) directory.
The generation code is in the
[generate](internal/generate) package.

All of the generated API code,
and a very small amount of manually maintained code,
is in the subdirectories of the
[lib](lib) directory.
 
## status
**caution**:
This library is immature,
but may be useable with caution
if the required APIs have been generated.

Some simple example applications have been created,
and may be found in the [example](./example) directory.
It is likely that for more complex applications
some necessary apis may be missing. 

The API is reasonably stable,
although some changes are certainly possible. 

The generation code is quite messy in some areas,
and needs tidying up.

## documentation
https://pekim.github.io/gobbi/

## gotk3
[gotk3](https://github.com/gotk3/gotk3)
is a much more mature library than gobbi.
At this time it is a safer choice than gobbi.

gotk3 supports some apis that gobbi does not, and vice versa.
Over time gobbi will add support for more apis.

gotk3 pays good attention to memory management and gobject
reference counting.
gobbi has work to do in this area. 

gobbi is more type safe in some areas,
notably the callbacks connected to signals.
