gobbi uses Go 
[build tags](https://golang.org/pkg/go/build/#hdr-Build_Constraints)
to allow targetting of specific library versions.
This means that when using Go build tools
(such as `go build` and `go run`)
to build a gobbi application,
use of the `-tags` flag will usually be necessary.

See
[How to use conditional compilation with the go build tool](https://dave.cheney.net/2013/10/12/how-to-use-conditional-compilation-with-the-go-build-tool)
for more background about build tags.

## no tags
If a gobbi application is built with no tags specified,
the only apis available will be those present in
all version of the various gnome libraries.

## tags targetting versions
To use an api that is only available from a specific
version onwards,
use a tag constructed from the library name and the version.

For example
[gtk_builder_new_from_file](https://developer.gnome.org/gtk3/stable/GtkBuilder.html#gtk-builder-new-from-file)
is only available from gtk 3.10 onwards.
So to use it, a tag specifying 3.10 (or a later version)
is required.

```bash
go build -tags gtk_3.10 my_app.go
```

## typical tags
If targetting Ubuntu 16.04 (Xenial Xerus) or later,
the following tags would be a good starting point. 

```bash
go build -tags " \
    glib_2.48 \
    gdk_3.4 \
    gdkpixbuf_2.32 \
    gio_2.32 \
    gtk_3.18" \
  my_app.go
```

## editors & IDEs
Editors and IDE can provide better completion
and error feedback if they are aware of the build
tags that an application is to be built with.

### vs code
The relevant setting is `Go: Build Tags`.

### goland
In the `Settings...` dialog, the relevant setting can
be found under
- Go
  - Vendoring & Build Tags
    - Custom tags
