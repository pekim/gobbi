The gobject library implements
[Object memory management](https://developer.gnome.org/gobject/stable/gobject-memory.html),
including reference counting,
for GObject derived classes.
When gobbi creates such an object and a Go object referencing it,
it will increment the reference count on the object.
When the Go object is finalized, gobbi will
decrement the GObject's reference count.

The behaviour is slightly different between objects that are
initially owned and those that are initially unowned.
In both cases gobbi attempts to do the right thing
with regard to reference counting.

## Copying gobbi objects
Unless there is a compelling reason,
gobbi created objects should not be copied
or cloned directly by applications,
as this would result in references to GObjects
without the necessary reference count.

If a copy of a gobbi Object derived class is required,
it can be achieved as a side effect
of a suitable cast like this.
As the new object is gobbi managed is will be reference
counted correctly. 

```go
    label1 := gtk.LabelNew("some text")
	label2 := gtk.CastToLabel(label1.Object())
```

However in most cases even that will not be required,
and passing around the pointer
to the gobbi object will be sufficient for an application.
