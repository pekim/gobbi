There is no direct support in gobbi for subclassing
GObject derived classes
or for implementing interfaces.
Following exploration and the production of a
proof of concept for class derivation,
it became apparent that an awful lot of work
would have to be put it for a fairly small reward.
Lots more generation code would have to be written,
and many tens of thousands of new lines of code
would be generated.
And even then there would likely be many
cases not covered.

Deriving classes, implementing interfaces,
and implementing virtual functions are
unlikely to be particularly common activities
in gobbi based applications.
So for now at least, adding support to make this
easy has been put to one side.

## example
Instead of providing direct support in gobbi, a
[subclassing](https://github.com/pekim/gobbi/blob/master/example/subclass-drawingarea)
example is provided.
This illustrates how the DrawingArea widget can be subclassed,
and some virtual functions implemented.

## pre-requisites
For the most part using gobbi does not require
a detailed knowledge C or gobject.
With some familiarity with Go, Gtk and perhaps
a passing knowledge of cgo,
it should be possible to write an application
with gobbi.

However subclassing and the implementation of
virtual functions will require a bit
more knowledge.

- comfort with Go
- familiarity with C
- familiarity with cgo
- an understanding of the [gobject base class](https://developer.gnome.org/gobject/stable/chapter-gobject.html)

