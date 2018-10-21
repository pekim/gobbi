# gobbi

**gob**ject **bi**ndings

gobbi is a set of Go bindings for gobject based libraries.
There are bindings, with varying degrees of completeness,
for gtk, gdk, gdkpixbuf, gio, gobject, and glib.

## status
**caution**: This library is immature,
but may be useable with caution
if the apis required have been generated.

Some simple example applications have been created.
It is likely that for more complex applications
some necessary apis may be missing. 

The API is reasonably stable,
although some changes are certainly possible. 

The generation code is quite messy in some areas,
and needs tidying up.

## pre-requisites
* go 1.11 or later
* libraries
    * debian/ubuntu packages -
        `libatk1.0-dev libcairo2-dev libglib2.0-dev
        libgtk-3-dev libpango1.0-dev`
    * note that dev libraries are required because pkg-config
        and headers are necessary
* C compiler, and a linker
    * required because of the use of cgo
    * for debian based distributions the `build-essential`
        package should suffice 

## getting started
* `git clone https://github.com/pekim/gobbi.git`
* `cd gobbi`
* `./do example simple_window`
    * Be patient.
        The first `go build` or `go run` will take quite a while,
        perhaps a few minutes.
        Subsequent builds of applications will be a *lot* quicker
        as cached gobbi packages will be used by the Go tool set. 

There are more examples in the `example` directory. 

## documentation
There is currently very little Godoc api documentation for gobbi.
As most apis are relatively direct mappings of the C library
apis, the libraries' own documentation should provide
enough information to gain an understanding of how to
use the Go apis.

Out parameters are mapped as return value in Go functions.
Inout parameters are not yet supported, but when they are
the will be exposed as both parameters and return values
in Go functions.

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
