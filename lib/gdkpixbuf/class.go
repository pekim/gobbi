// This is a generated file - DO NOT EDIT

package gdkpixbuf

import gio "github.com/pekim/gobbi/lib/gio"

// Unsupported : gdk_pixbuf_new : return type :

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter destroy_fn : no type generator for PixbufDestroyNotify (GdkPixbufDestroyNotify) for param destroy_fn

// Unsupported : gdk_pixbuf_new_from_file : return type :

// Unsupported : gdk_pixbuf_new_from_inline : return type :

// Unsupported : gdk_pixbuf_new_from_xpm_data : return type :

// Icon returns the Icon interface implemented by Pixbuf
func (recv *Pixbuf) Icon() *gio.Icon {
	return gio.IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by Pixbuf
func (recv *Pixbuf) LoadableIcon() *gio.LoadableIcon {
	return gio.LoadableIconNewFromC(recv.ToC())
}

// Unsupported : gdk_pixbuf_animation_new_from_file : return type :

// Unsupported : gdk_pixbuf_loader_new : return type :

// Unsupported : gdk_pixbuf_loader_new_with_type : return type :

// Unsupported : PixbufSimpleAnimIter : no CType
