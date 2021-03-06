Gtk classes derive in a hierarchy from the GObject class.

In gobbi's implementation of GObject's classes,
each class does not directly extend
(or embed in Go terms)
its ancestor classes.
Instead, for an class instance it is possible to get
a reference to an ancestor or descendant class instance
for the same gtk object.  

# upcasting
In a gtk application there will frequently be a need to
access members of an object's ancestor class.

For example when adding a Label to a Container,
this will not work.

```go
window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
label := gtk.LabelNew("a label")
window.Add(label)   // <-- ERROR
``` 

`gtk.Window` does not directly have an `Add` method.
The `Add` method is a member of `gtk.Container`,
which is an ancestor of `gtk.Window`.

The `gtk.Window`'s `Container` may be obtained by calling
the `Container()` method.

And as the `Add` method expects a `gtk.Widget`,
the `gtk.Label`'s `Widget` method is called
to cast to a `gtk.Widget`.

```go
window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
label := gtk.LabelNew("a label")
window.Container().Add(label.Widget())
``` 

All derived classes have receiver functions for directly
casting to all ancestor types.
```go
// This will work, but is unnecessary.
widget1 := gtk.WindowNew().Bin().Container().Widget()

// This will achieve the same end result.
widget2 := gtk.WindowNew().Widget()
``` 

Classes also have receiver functions for casting to
any interfaces that they implement.
```go
// gtk.Buildable is an interface implemented by all Widgets.
widget := gtk.LabelNew("a label")
buildable := widget.Buildable()
``` 

# downcasting
Downcasting operates in the other direction,
from a class to a derived class.

Downcasting is less common that upcasting.
It is potentially quite dangerous, in that it makes
it very easy to assert that a widget is something
that it is not.

To downcast an instance,
call one of the `CastTo...` methods.

A typically use is when obtaining an object from a builder.

```go
builder := gtk.BuilderNewFromString(...)

buttonObject := builder.GetObject("some-button")
button := gtk.CastToButton(buttonObject)
button.ConnectClicked(...)
```

See the
[builder](https://github.com/pekim/gobbi/blob/master/example/builder/main.go)
example to see this in an application. 
