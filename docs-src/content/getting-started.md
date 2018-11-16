+++
title="Getting started"
menu="main"
weight="2"
+++

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

## running examples
* `git clone https://github.com/pekim/gobbi.git`
* `cd gobbi`
* `./gobbi example simple_window`
    * Be patient.
        The first `go build` or `go run` will take quite a while,
        perhaps a few minutes.
        Subsequent builds of applications will be a *lot* quicker
        as cached gobbi packages will be used by the Go commands. 

Run `./gobbi examples` for a list of the available examples.

## use in a application
Using gobbi in an application in no different to using
most other libraries.

To add gobbi to a project that make use of go modules
use `go get`.
```bash
go get github.com/pekim/gobbi
```

Then import a package such as gtk.
```go
import "github.com/pekim/gobbi/lib/gtk"
```

## build tags
It will often be necessary to build gobbi applications
with [build tags](./build-tags.md).
