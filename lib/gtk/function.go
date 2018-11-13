// This is a generated file - DO NOT EDIT

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Finds the first accelerator in any #GtkAccelGroup attached
// to @object that matches @accel_key and @accel_mods, and
// activates that accelerator.
/*

C function

gtk_accel_groups_activate
*/
func AccelGroupsActivate(object *gobject.Object, accelKey uint32, accelMods gdk.ModifierType) bool {
	c_object := (*C.GObject)(C.NULL)
	if object != nil {
		c_object = (*C.GObject)(object.ToC())
	}

	c_accel_key := (C.guint)(accelKey)

	c_accel_mods := (C.GdkModifierType)(accelMods)

	retC := C.gtk_accel_groups_activate(c_object, c_accel_key, c_accel_mods)
	retGo := retC == C.TRUE

	return retGo
}

// Gets a list of all accel groups which are attached to @object.
/*

C function

gtk_accel_groups_from_object
*/
func AccelGroupsFromObject(object *gobject.Object) *glib.SList {
	c_object := (*C.GObject)(C.NULL)
	if object != nil {
		c_object = (*C.GObject)(object.ToC())
	}

	retC := C.gtk_accel_groups_from_object(c_object)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the modifier mask.
//
// The modifier mask determines which modifiers are considered significant
// for keyboard accelerators. See gtk_accelerator_set_default_mod_mask().
/*

C function

gtk_accelerator_get_default_mod_mask
*/
func AcceleratorGetDefaultModMask() gdk.ModifierType {
	retC := C.gtk_accelerator_get_default_mod_mask()
	retGo := (gdk.ModifierType)(retC)

	return retGo
}

// Converts an accelerator keyval and modifier mask into a string
// parseable by gtk_accelerator_parse(). For example, if you pass in
// #GDK_KEY_q and #GDK_CONTROL_MASK, this function returns “<Control>q”.
//
// If you need to display accelerators in the user interface,
// see gtk_accelerator_get_label().
/*

C function

gtk_accelerator_name
*/
func AcceleratorName(acceleratorKey uint32, acceleratorMods gdk.ModifierType) string {
	c_accelerator_key := (C.guint)(acceleratorKey)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	retC := C.gtk_accelerator_name(c_accelerator_key, c_accelerator_mods)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// Sets the modifiers that will be considered significant for keyboard
// accelerators. The default mod mask depends on the GDK backend in use,
// but will typically include #GDK_CONTROL_MASK | #GDK_SHIFT_MASK |
// #GDK_MOD1_MASK | #GDK_SUPER_MASK | #GDK_HYPER_MASK | #GDK_META_MASK.
// In other words, Control, Shift, Alt, Super, Hyper and Meta. Other
// modifiers will by default be ignored by #GtkAccelGroup.
//
// You must include at least the three modifiers Control, Shift
// and Alt in any value you pass to this function.
//
// The default mod mask should be changed on application startup,
// before using any accelerator groups.
/*

C function

gtk_accelerator_set_default_mod_mask
*/
func AcceleratorSetDefaultModMask(defaultModMask gdk.ModifierType) {
	c_default_mod_mask := (C.GdkModifierType)(defaultModMask)

	C.gtk_accelerator_set_default_mod_mask(c_default_mod_mask)

	return
}

// Determines whether a given keyval and modifier mask constitute
// a valid keyboard accelerator. For example, the #GDK_KEY_a keyval
// plus #GDK_CONTROL_MASK is valid - this is a “Ctrl+a” accelerator.
// But, you can't, for instance, use the #GDK_KEY_Control_L keyval
// as an accelerator.
/*

C function

gtk_accelerator_valid
*/
func AcceleratorValid(keyval uint32, modifiers gdk.ModifierType) bool {
	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_accelerator_valid(c_keyval, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}

// Override or install a new key binding for @keyval with @modifiers on
// @binding_set.
/*

C function

gtk_binding_entry_add_signall
*/
func BindingEntryAddSignall(bindingSet *BindingSet, keyval uint32, modifiers gdk.ModifierType, signalName string, bindingArgs *glib.SList) {
	c_binding_set := (*C.GtkBindingSet)(C.NULL)
	if bindingSet != nil {
		c_binding_set = (*C.GtkBindingSet)(bindingSet.ToC())
	}

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_signal_name := C.CString(signalName)
	defer C.free(unsafe.Pointer(c_signal_name))

	c_binding_args := (*C.GSList)(C.NULL)
	if bindingArgs != nil {
		c_binding_args = (*C.GSList)(bindingArgs.ToC())
	}

	C.gtk_binding_entry_add_signall(c_binding_set, c_keyval, c_modifiers, c_signal_name, c_binding_args)

	return
}

// Remove a binding previously installed via
// gtk_binding_entry_add_signal() on @binding_set.
/*

C function

gtk_binding_entry_remove
*/
func BindingEntryRemove(bindingSet *BindingSet, keyval uint32, modifiers gdk.ModifierType) {
	c_binding_set := (*C.GtkBindingSet)(C.NULL)
	if bindingSet != nil {
		c_binding_set = (*C.GtkBindingSet)(bindingSet.ToC())
	}

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gtk_binding_entry_remove(c_binding_set, c_keyval, c_modifiers)

	return
}

// This function returns the binding set named after the type name of
// the passed in class structure. New binding sets are created on
// demand by this function.
/*

C function

gtk_binding_set_by_class
*/
func BindingSetByClass(objectClass uintptr) *BindingSet {
	c_object_class := (C.gpointer)(objectClass)

	retC := C.gtk_binding_set_by_class(c_object_class)
	retGo := BindingSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Find a binding set by its globally unique name.
//
// The @set_name can either be a name used for gtk_binding_set_new()
// or the type name of a class used in gtk_binding_set_by_class().
/*

C function

gtk_binding_set_find
*/
func BindingSetFind(setName string) *BindingSet {
	c_set_name := C.CString(setName)
	defer C.free(unsafe.Pointer(c_set_name))

	retC := C.gtk_binding_set_find(c_set_name)
	var retGo (*BindingSet)
	if retC == nil {
		retGo = nil
	} else {
		retGo = BindingSetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GTK+ maintains a global list of binding sets. Each binding set has
// a unique name which needs to be specified upon creation.
/*

C function

gtk_binding_set_new
*/
func BindingSetNew(setName string) *BindingSet {
	c_set_name := C.CString(setName)
	defer C.free(unsafe.Pointer(c_set_name))

	retC := C.gtk_binding_set_new(c_set_name)
	retGo := BindingSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Find a key binding matching @keyval and @modifiers and activate the
// binding on @object.
/*

C function

gtk_bindings_activate
*/
func BindingsActivate(object *gobject.Object, keyval uint32, modifiers gdk.ModifierType) bool {
	c_object := (*C.GObject)(C.NULL)
	if object != nil {
		c_object = (*C.GObject)(object.ToC())
	}

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_bindings_activate(c_object, c_keyval, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}

/*

C function

gtk_builder_error_quark
*/
func BuilderErrorQuark() glib.Quark {
	retC := C.gtk_builder_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Checks that the GTK+ library in use is compatible with the
// given version. Generally you would pass in the constants
// #GTK_MAJOR_VERSION, #GTK_MINOR_VERSION, #GTK_MICRO_VERSION
// as the three arguments to this function; that produces
// a check that the library in use is compatible with
// the version of GTK+ the application or module was compiled
// against.
//
// Compatibility is defined by two things: first the version
// of the running library is newer than the version
// @required_major.required_minor.@required_micro. Second
// the running library must be binary compatible with the
// version @required_major.required_minor.@required_micro
// (same major version.)
//
// This function is primarily for GTK+ modules; the module
// can call this function to check that it wasn’t loaded
// into an incompatible version of GTK+. However, such a
// check isn’t completely reliable, since the module may be
// linked against an old version of GTK+ and calling the
// old version of gtk_check_version(), but still get loaded
// into an application using a newer version of GTK+.
/*

C function

gtk_check_version
*/
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
	c_required_major := (C.guint)(requiredMajor)

	c_required_minor := (C.guint)(requiredMinor)

	c_required_micro := (C.guint)(requiredMicro)

	retC := C.gtk_check_version(c_required_major, c_required_minor, c_required_micro)
	retGo := C.GoString(retC)

	return retGo
}

/*

C function

gtk_css_provider_error_quark
*/
func CssProviderErrorQuark() glib.Quark {
	retC := C.gtk_css_provider_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Prevents gtk_init(), gtk_init_check(), gtk_init_with_args() and
// gtk_parse_args() from automatically
// calling `setlocale (LC_ALL, "")`. You would
// want to use this function if you wanted to set the locale for
// your program to something other than the user’s locale, or if
// you wanted to set different values for different locale categories.
//
// Most programs should not need to call this function.
/*

C function

gtk_disable_setlocale
*/
func DisableSetlocale() {
	C.gtk_disable_setlocale()

	return
}

// Distributes @extra_space to child @sizes by bringing smaller
// children up to natural size first.
//
// The remaining space will be added to the @minimum_size member of the
// GtkRequestedSize struct. If all sizes reach their natural size then
// the remaining space is returned.
/*

C function

gtk_distribute_natural_allocation
*/
func DistributeNaturalAllocation(extraSpace int32, nRequestedSizes uint32, sizes *RequestedSize) int32 {
	c_extra_space := (C.gint)(extraSpace)

	c_n_requested_sizes := (C.guint)(nRequestedSizes)

	c_sizes := (*C.GtkRequestedSize)(C.NULL)
	if sizes != nil {
		c_sizes = (*C.GtkRequestedSize)(sizes.ToC())
	}

	retC := C.gtk_distribute_natural_allocation(c_extra_space, c_n_requested_sizes, c_sizes)
	retGo := (int32)(retC)

	return retGo
}

// Informs the drag source that the drop is finished, and
// that the data of the drag will no longer be required.
/*

C function

gtk_drag_finish
*/
func DragFinish(context *gdk.DragContext, success bool, del bool, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_success :=
		boolToGboolean(success)

	c_del :=
		boolToGboolean(del)

	c_time_ := (C.guint32)(time)

	C.gtk_drag_finish(c_context, c_success, c_del, c_time_)

	return
}

// Determines the source widget for a drag.
/*

C function

gtk_drag_get_source_widget
*/
func DragGetSourceWidget(context *gdk.DragContext) *Widget {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	retC := C.gtk_drag_get_source_widget(c_context)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets the icon for a particular drag to the default
// icon.
/*

C function

gtk_drag_set_icon_default
*/
func DragSetIconDefault(context *gdk.DragContext) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	C.gtk_drag_set_icon_default(c_context)

	return
}

// Sets @pixbuf as the icon for a given drag.
/*

C function

gtk_drag_set_icon_pixbuf
*/
func DragSetIconPixbuf(context *gdk.DragContext, pixbuf *gdkpixbuf.Pixbuf, hotX int32, hotY int32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	c_hot_x := (C.gint)(hotX)

	c_hot_y := (C.gint)(hotY)

	C.gtk_drag_set_icon_pixbuf(c_context, c_pixbuf, c_hot_x, c_hot_y)

	return
}

// Sets the icon for a given drag from a stock ID.
/*

C function

gtk_drag_set_icon_stock
*/
func DragSetIconStock(context *gdk.DragContext, stockId string, hotX int32, hotY int32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_hot_x := (C.gint)(hotX)

	c_hot_y := (C.gint)(hotY)

	C.gtk_drag_set_icon_stock(c_context, c_stock_id, c_hot_x, c_hot_y)

	return
}

// Sets @surface as the icon for a given drag. GTK+ retains
// references for the arguments, and will release them when
// they are no longer needed.
//
// To position the surface relative to the mouse, use
// cairo_surface_set_device_offset() on @surface. The mouse
// cursor will be positioned at the (0,0) coordinate of the
// surface.
/*

C function

gtk_drag_set_icon_surface
*/
func DragSetIconSurface(context *gdk.DragContext, surface *cairo.Surface) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	C.gtk_drag_set_icon_surface(c_context, c_surface)

	return
}

// Changes the icon for drag operation to a given widget.
// GTK+ will not destroy the widget, so if you don’t want
// it to persist, you should connect to the “drag-end”
// signal and destroy it yourself.
/*

C function

gtk_drag_set_icon_widget
*/
func DragSetIconWidget(context *gdk.DragContext, widget *Widget, hotX int32, hotY int32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_hot_x := (C.gint)(hotX)

	c_hot_y := (C.gint)(hotY)

	C.gtk_drag_set_icon_widget(c_context, c_widget, c_hot_x, c_hot_y)

	return
}

// Checks if any events are pending.
//
// This can be used to update the UI and invoke timeouts etc.
// while doing some time intensive computation.
//
// Updating the UI during a long computation
//
// |[<!-- language="C" -->
// computation going on...
//
// while (gtk_events_pending ())
// gtk_main_iteration ();
//
// ...computation continued
// ]|
/*

C function

gtk_events_pending
*/
func EventsPending() bool {
	retC := C.gtk_events_pending()
	retGo := retC == C.TRUE

	return retGo
}

// Analogical to gtk_true(), this function does nothing
// but always returns %FALSE.
/*

C function

gtk_false
*/
func False() bool {
	retC := C.gtk_false()
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_get_current_event : no return generator

// If there is a current event and it has a device, return that
// device, otherwise return %NULL.
/*

C function

gtk_get_current_event_device
*/
func GetCurrentEventDevice() *gdk.Device {
	retC := C.gtk_get_current_event_device()
	var retGo (*gdk.Device)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdk.DeviceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_get_current_event_state : unsupported parameter state : GdkModifierType* with indirection level of 1

// If there is a current event and it has a timestamp,
// return that timestamp, otherwise return %GDK_CURRENT_TIME.
/*

C function

gtk_get_current_event_time
*/
func GetCurrentEventTime() uint32 {
	retC := C.gtk_get_current_event_time()
	retGo := (uint32)(retC)

	return retGo
}

// Returns the GTK+ debug flags.
//
// This function is intended for GTK+ modules that want
// to adjust their debug output based on GTK+ debug flags.
/*

C function

gtk_get_debug_flags
*/
func GetDebugFlags() uint32 {
	retC := C.gtk_get_debug_flags()
	retGo := (uint32)(retC)

	return retGo
}

// Returns the #PangoLanguage for the default language currently in
// effect. (Note that this can change over the life of an
// application.) The default language is derived from the current
// locale. It determines, for example, whether GTK+ uses the
// right-to-left or left-to-right text direction.
//
// This function is equivalent to pango_language_get_default().
// See that function for details.
/*

C function

gtk_get_default_language
*/
func GetDefaultLanguage() *pango.Language {
	retC := C.gtk_get_default_language()
	retGo := pango.LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Queries the current grab of the default window group.
/*

C function

gtk_grab_get_current
*/
func GrabGetCurrent() *Widget {
	retC := C.gtk_grab_get_current()
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Looks up the icon size associated with @name.
/*

C function

gtk_icon_size_from_name
*/
func IconSizeFromName(name string) IconSize {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_icon_size_from_name(c_name)
	retGo := (IconSize)(retC)

	return retGo
}

// Gets the canonical name of the given icon size. The returned string
// is statically allocated and should not be freed.
/*

C function

gtk_icon_size_get_name
*/
func IconSizeGetName(size IconSize) string {
	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_icon_size_get_name(c_size)
	retGo := C.GoString(retC)

	return retGo
}

// Obtains the pixel size of a semantic icon size @size:
// #GTK_ICON_SIZE_MENU, #GTK_ICON_SIZE_BUTTON, etc.  This function
// isn’t normally needed, gtk_icon_theme_load_icon() is the usual
// way to get an icon for rendering, then just look at the size of
// the rendered pixbuf. The rendered pixbuf may not even correspond to
// the width/height returned by gtk_icon_size_lookup(), because themes
// are free to render the pixbuf however they like, including changing
// the usual size.
/*

C function

gtk_icon_size_lookup
*/
func IconSizeLookup(size IconSize) (bool, int32, int32) {
	c_size := (C.GtkIconSize)(size)

	var c_width C.gint

	var c_height C.gint

	retC := C.gtk_icon_size_lookup(c_size, &c_width, &c_height)
	retGo := retC == C.TRUE

	width := (int32)(c_width)

	height := (int32)(c_height)

	return retGo, width, height
}

// Registers a new icon size, along the same lines as #GTK_ICON_SIZE_MENU,
// etc. Returns the integer value for the size.
/*

C function

gtk_icon_size_register
*/
func IconSizeRegister(name string, width int32, height int32) IconSize {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.gtk_icon_size_register(c_name, c_width, c_height)
	retGo := (IconSize)(retC)

	return retGo
}

// Registers @alias as another name for @target.
// So calling gtk_icon_size_from_name() with @alias as argument
// will return @target.
/*

C function

gtk_icon_size_register_alias
*/
func IconSizeRegisterAlias(alias string, target IconSize) {
	c_alias := C.CString(alias)
	defer C.free(unsafe.Pointer(c_alias))

	c_target := (C.GtkIconSize)(target)

	C.gtk_icon_size_register_alias(c_alias, c_target)

	return
}

/*

C function

gtk_icon_theme_error_quark
*/
func IconThemeErrorQuark() glib.Quark {
	retC := C.gtk_icon_theme_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Call this function before using any other GTK+ functions in your GUI
// applications.  It will initialize everything needed to operate the
// toolkit and parses some standard command line options.
//
// Although you are expected to pass the @argc, @argv parameters from main() to
// this function, it is possible to pass %NULL if @argv is not available or
// commandline handling is not required.
//
// @argc and @argv are adjusted accordingly so your own code will
// never see those standard arguments.
//
// Note that there are some alternative ways to initialize GTK+:
// if you are calling gtk_parse_args(), gtk_init_check(),
// gtk_init_with_args() or g_option_context_parse() with
// the option group returned by gtk_get_option_group(),
// you don’t have to call gtk_init().
//
// And if you are using #GtkApplication, you don't have to call any of the
// initialization functions either; the #GtkApplication::startup handler
// does it for you.
//
// This function will terminate your program if it was unable to
// initialize the windowing system for some reason. If you want
// your program to fall back to a textual interface you want to
// call gtk_init_check() instead.
//
// Since 2.18, GTK+ calls `signal (SIGPIPE, SIG_IGN)`
// during initialization, to ignore SIGPIPE signals, since these are
// almost never wanted in graphical applications. If you do need to
// handle SIGPIPE for some reason, reset the handler after gtk_init(),
// but notice that other libraries (e.g. libdbus or gvfs) might do
// similar things.
/*

C function

gtk_init
*/
func Init(args []string) []string {
	cArgc, cArgv := argsIn(args)

	C.gtk_init(&cArgc, &cArgv)

	args = argsOut(cArgc, cArgv)

	return args
}

// This function does the same work as gtk_init() with only a single
// change: It does not terminate the program if the commandline
// arguments couldn’t be parsed or the windowing system can’t be
// initialized. Instead it returns %FALSE on failure.
//
// This way the application can fall back to some other means of
// communication with the user - for example a curses or command line
// interface.
/*

C function

gtk_init_check
*/
func InitCheck(args []string) (bool, []string) {
	cArgc, cArgv := argsIn(args)

	retC := C.gtk_init_check(&cArgc, &cArgv)
	retGo := retC == C.TRUE

	args = argsOut(cArgc, cArgv)

	return retGo, args
}

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc (GtkKeySnoopFunc) for param snooper

// Removes the key snooper function with the given id.
/*

C function

gtk_key_snooper_remove
*/
func KeySnooperRemove(snooperHandlerId uint32) {
	c_snooper_handler_id := (C.guint)(snooperHandlerId)

	C.gtk_key_snooper_remove(c_snooper_handler_id)

	return
}

// Runs the main loop until gtk_main_quit() is called.
//
// You can nest calls to gtk_main(). In that case gtk_main_quit()
// will make the innermost invocation of the main loop return.
/*

C function

gtk_main
*/
func Main() {
	C.gtk_main()

	return
}

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Runs a single iteration of the mainloop.
//
// If no events are waiting to be processed GTK+ will block
// until the next event is noticed. If you don’t want to block
// look at gtk_main_iteration_do() or check if any events are
// pending with gtk_events_pending() first.
/*

C function

gtk_main_iteration
*/
func MainIteration() bool {
	retC := C.gtk_main_iteration()
	retGo := retC == C.TRUE

	return retGo
}

// Runs a single iteration of the mainloop.
// If no events are available either return or block depending on
// the value of @blocking.
/*

C function

gtk_main_iteration_do
*/
func MainIterationDo(blocking bool) bool {
	c_blocking :=
		boolToGboolean(blocking)

	retC := C.gtk_main_iteration_do(c_blocking)
	retGo := retC == C.TRUE

	return retGo
}

// Asks for the current nesting level of the main loop.
/*

C function

gtk_main_level
*/
func MainLevel() uint32 {
	retC := C.gtk_main_level()
	retGo := (uint32)(retC)

	return retGo
}

// Makes the innermost invocation of the main loop return
// when it regains control.
/*

C function

gtk_main_quit
*/
func MainQuit() {
	C.gtk_main_quit()

	return
}

// Draws an arrow in the given rectangle on @cr using the given
// parameters. @arrow_type determines the direction of the arrow.
/*

C function

gtk_paint_arrow
*/
func PaintArrow(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, arrowType ArrowType, fill bool, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_arrow_type := (C.GtkArrowType)(arrowType)

	c_fill :=
		boolToGboolean(fill)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_arrow(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_arrow_type, c_fill, c_x, c_y, c_width, c_height)

	return
}

// Draws a box on @cr with the given parameters.
/*

C function

gtk_paint_box
*/
func PaintBox(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_box(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// Draws a box in @cr using the given style and state and shadow type,
// leaving a gap in one side.
/*

C function

gtk_paint_box_gap
*/
func PaintBoxGap(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32, gapSide PositionType, gapX int32, gapWidth int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	c_gap_x := (C.gint)(gapX)

	c_gap_width := (C.gint)(gapWidth)

	C.gtk_paint_box_gap(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height, c_gap_side, c_gap_x, c_gap_width)

	return
}

// Draws a check button indicator in the given rectangle on @cr with
// the given parameters.
/*

C function

gtk_paint_check
*/
func PaintCheck(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_check(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// Draws a diamond in the given rectangle on @window using the given
// parameters.
/*

C function

gtk_paint_diamond
*/
func PaintDiamond(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_diamond(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// Draws an expander as used in #GtkTreeView. @x and @y specify the
// center the expander. The size of the expander is determined by the
// “expander-size” style property of @widget.  (If widget is not
// specified or doesn’t have an “expander-size” property, an
// unspecified default size will be used, since the caller doesn't
// have sufficient information to position the expander, this is
// likely not useful.) The expander is expander_size pixels tall
// in the collapsed position and expander_size pixels wide in the
// expanded position.
/*

C function

gtk_paint_expander
*/
func PaintExpander(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, x int32, y int32, expanderStyle ExpanderStyle) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_expander_style := (C.GtkExpanderStyle)(expanderStyle)

	C.gtk_paint_expander(c_style, c_cr, c_state_type, c_widget, c_detail, c_x, c_y, c_expander_style)

	return
}

// Draws an extension, i.e. a notebook tab.
/*

C function

gtk_paint_extension
*/
func PaintExtension(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32, gapSide PositionType) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	C.gtk_paint_extension(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height, c_gap_side)

	return
}

// Draws a flat box on @cr with the given parameters.
/*

C function

gtk_paint_flat_box
*/
func PaintFlatBox(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_flat_box(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// Draws a focus indicator around the given rectangle on @cr using the
// given style.
/*

C function

gtk_paint_focus
*/
func PaintFocus(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_focus(c_style, c_cr, c_state_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// Draws a handle as used in #GtkHandleBox and #GtkPaned.
/*

C function

gtk_paint_handle
*/
func PaintHandle(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32, orientation Orientation) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_paint_handle(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height, c_orientation)

	return
}

// Draws a horizontal line from (@x1, @y) to (@x2, @y) in @cr
// using the given style and state.
/*

C function

gtk_paint_hline
*/
func PaintHline(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, x1 int32, x2 int32, y int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x1 := (C.gint)(x1)

	c_x2 := (C.gint)(x2)

	c_y := (C.gint)(y)

	C.gtk_paint_hline(c_style, c_cr, c_state_type, c_widget, c_detail, c_x1, c_x2, c_y)

	return
}

// Draws a layout on @cr using the given parameters.
/*

C function

gtk_paint_layout
*/
func PaintLayout(style *Style, cr *cairo.Context, stateType StateType, useText bool, widget *Widget, detail string, x int32, y int32, layout *pango.Layout) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_use_text :=
		boolToGboolean(useText)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	C.gtk_paint_layout(c_style, c_cr, c_state_type, c_use_text, c_widget, c_detail, c_x, c_y, c_layout)

	return
}

// Draws a radio button indicator in the given rectangle on @cr with
// the given parameters.
/*

C function

gtk_paint_option
*/
func PaintOption(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_option(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// Draws a resize grip in the given rectangle on @cr using the given
// parameters.
/*

C function

gtk_paint_resize_grip
*/
func PaintResizeGrip(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, edge gdk.WindowEdge, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_edge := (C.GdkWindowEdge)(edge)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_resize_grip(c_style, c_cr, c_state_type, c_widget, c_detail, c_edge, c_x, c_y, c_width, c_height)

	return
}

// Draws a shadow around the given rectangle in @cr
// using the given style and state and shadow type.
/*

C function

gtk_paint_shadow
*/
func PaintShadow(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_shadow(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// Draws a shadow around the given rectangle in @cr
// using the given style and state and shadow type, leaving a
// gap in one side.
/*

C function

gtk_paint_shadow_gap
*/
func PaintShadowGap(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32, gapSide PositionType, gapX int32, gapWidth int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	c_gap_x := (C.gint)(gapX)

	c_gap_width := (C.gint)(gapWidth)

	C.gtk_paint_shadow_gap(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height, c_gap_side, c_gap_x, c_gap_width)

	return
}

// Draws a slider in the given rectangle on @cr using the
// given style and orientation.
/*

C function

gtk_paint_slider
*/
func PaintSlider(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32, orientation Orientation) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_paint_slider(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height, c_orientation)

	return
}

// Draws a spinner on @window using the given parameters.
/*

C function

gtk_paint_spinner
*/
func PaintSpinner(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, step uint32, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_step := (C.guint)(step)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_spinner(c_style, c_cr, c_state_type, c_widget, c_detail, c_step, c_x, c_y, c_width, c_height)

	return
}

// Draws an option menu tab (i.e. the up and down pointing arrows)
// in the given rectangle on @cr using the given parameters.
/*

C function

gtk_paint_tab
*/
func PaintTab(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_tab(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// Draws a vertical line from (@x, @y1_) to (@x, @y2_) in @cr
// using the given style and state.
/*

C function

gtk_paint_vline
*/
func PaintVline(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, y1 int32, y2 int32, x int32) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_y1_ := (C.gint)(y1)

	c_y2_ := (C.gint)(y2)

	c_x := (C.gint)(x)

	C.gtk_paint_vline(c_style, c_cr, c_state_type, c_widget, c_detail, c_y1_, c_y2_, c_x)

	return
}

// Parses command line arguments, and initializes global
// attributes of GTK+, but does not actually open a connection
// to a display. (See gdk_display_open(), gdk_get_display_arg_name())
//
// Any arguments used by GTK+ or GDK are removed from the array and
// @argc and @argv are updated accordingly.
//
// There is no need to call this function explicitly if you are using
// gtk_init(), or gtk_init_check().
//
// Note that many aspects of GTK+ require a display connection to
// function, so this way of initializing GTK+ is really only useful
// for specialized use cases.
/*

C function

gtk_parse_args
*/
func ParseArgs(args []string) (bool, []string) {
	cArgc, cArgv := argsIn(args)

	retC := C.gtk_parse_args(&cArgc, &cArgv)
	retGo := retC == C.TRUE

	args = argsOut(cArgc, cArgv)

	return retGo, args
}

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Adds a file to the list of files to be parsed at the
// end of gtk_init().
/*

C function

gtk_rc_add_default_file
*/
func RcAddDefaultFile(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_rc_add_default_file(c_filename)

	return
}

// Searches for a theme engine in the GTK+ search path. This function
// is not useful for applications and should not be used.
/*

C function

gtk_rc_find_module_in_path
*/
func RcFindModuleInPath(moduleFile string) string {
	c_module_file := C.CString(moduleFile)
	defer C.free(unsafe.Pointer(c_module_file))

	retC := C.gtk_rc_find_module_in_path(c_module_file)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Looks up a file in pixmap path for the specified #GtkSettings.
// If the file is not found, it outputs a warning message using
// g_warning() and returns %NULL.
/*

C function

gtk_rc_find_pixmap_in_path
*/
func RcFindPixmapInPath(settings *Settings, scanner *glib.Scanner, pixmapFile string) string {
	c_settings := (*C.GtkSettings)(C.NULL)
	if settings != nil {
		c_settings = (*C.GtkSettings)(settings.ToC())
	}

	c_scanner := (*C.GScanner)(C.NULL)
	if scanner != nil {
		c_scanner = (*C.GScanner)(scanner.ToC())
	}

	c_pixmap_file := C.CString(pixmapFile)
	defer C.free(unsafe.Pointer(c_pixmap_file))

	retC := C.gtk_rc_find_pixmap_in_path(c_settings, c_scanner, c_pixmap_file)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_rc_get_default_files : no return type

// Obtains the path to the IM modules file. See the documentation
// of the `GTK_IM_MODULE_FILE`
// environment variable for more details.
/*

C function

gtk_rc_get_im_module_file
*/
func RcGetImModuleFile() string {
	retC := C.gtk_rc_get_im_module_file()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Obtains the path in which to look for IM modules. See the documentation
// of the `GTK_PATH`
// environment variable for more details about looking up modules. This
// function is useful solely for utilities supplied with GTK+ and should
// not be used by applications under normal circumstances.
/*

C function

gtk_rc_get_im_module_path
*/
func RcGetImModulePath() string {
	retC := C.gtk_rc_get_im_module_path()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns a directory in which GTK+ looks for theme engines.
// For full information about the search for theme engines,
// see the docs for `GTK_PATH` in [Running GTK+ Applications][gtk-running].
/*

C function

gtk_rc_get_module_dir
*/
func RcGetModuleDir() string {
	retC := C.gtk_rc_get_module_dir()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Finds all matching RC styles for a given widget,
// composites them together, and then creates a
// #GtkStyle representing the composite appearance.
// (GTK+ actually keeps a cache of previously
// created styles, so a new style may not be
// created.)
/*

C function

gtk_rc_get_style
*/
func RcGetStyle(widget *Widget) *Style {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_rc_get_style(c_widget)
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates up a #GtkStyle from styles defined in a RC file by providing
// the raw components used in matching. This function may be useful
// when creating pseudo-widgets that should be themed like widgets but
// don’t actually have corresponding GTK+ widgets. An example of this
// would be items inside a GNOME canvas widget.
//
// The action of gtk_rc_get_style() is similar to:
// |[<!-- language="C" -->
// gtk_widget_path (widget, NULL, &path, NULL);
// gtk_widget_class_path (widget, NULL, &class_path, NULL);
// gtk_rc_get_style_by_paths (gtk_widget_get_settings (widget),
// path, class_path,
// G_OBJECT_TYPE (widget));
// ]|
/*

C function

gtk_rc_get_style_by_paths
*/
func RcGetStyleByPaths(settings *Settings, widgetPath string, classPath string, type_ gobject.Type) *Style {
	c_settings := (*C.GtkSettings)(C.NULL)
	if settings != nil {
		c_settings = (*C.GtkSettings)(settings.ToC())
	}

	c_widget_path := C.CString(widgetPath)
	defer C.free(unsafe.Pointer(c_widget_path))

	c_class_path := C.CString(classPath)
	defer C.free(unsafe.Pointer(c_class_path))

	c_type := (C.GType)(type_)

	retC := C.gtk_rc_get_style_by_paths(c_settings, c_widget_path, c_class_path, c_type)
	var retGo (*Style)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StyleNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the standard directory in which themes should
// be installed. (GTK+ does not actually use this directory
// itself.)
/*

C function

gtk_rc_get_theme_dir
*/
func RcGetThemeDir() string {
	retC := C.gtk_rc_get_theme_dir()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Parses a given resource file.
/*

C function

gtk_rc_parse
*/
func RcParse(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_rc_parse(c_filename)

	return
}

// Parses a color in the format expected
// in a RC file.
//
// Note that theme engines should use gtk_rc_parse_color_full() in
// order to support symbolic colors.
/*

C function

gtk_rc_parse_color
*/
func RcParseColor(scanner *glib.Scanner) (uint32, *gdk.Color) {
	c_scanner := (*C.GScanner)(C.NULL)
	if scanner != nil {
		c_scanner = (*C.GScanner)(scanner.ToC())
	}

	var c_color C.GdkColor

	retC := C.gtk_rc_parse_color(c_scanner, &c_color)
	retGo := (uint32)(retC)

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// Unsupported : gtk_rc_parse_priority : unsupported parameter priority : GtkPathPriorityType* with indirection level of 1

// Unsupported : gtk_rc_parse_state : unsupported parameter state : GtkStateType* with indirection level of 1

// Parses resource information directly from a string.
/*

C function

gtk_rc_parse_string
*/
func RcParseString(rcString string) {
	c_rc_string := C.CString(rcString)
	defer C.free(unsafe.Pointer(c_rc_string))

	C.gtk_rc_parse_string(c_rc_string)

	return
}

// Unsupported : gtk_rc_property_parse_border : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_color : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_enum : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_flags : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_requisition : unsupported parameter pspec : Blacklisted record : GParamSpec

// If the modification time on any previously read file for the
// default #GtkSettings has changed, discard all style information
// and then reread all previously read RC files.
/*

C function

gtk_rc_reparse_all
*/
func RcReparseAll() bool {
	retC := C.gtk_rc_reparse_all()
	retGo := retC == C.TRUE

	return retGo
}

// If the modification time on any previously read file
// for the given #GtkSettings has changed, discard all style information
// and then reread all previously read RC files.
/*

C function

gtk_rc_reparse_all_for_settings
*/
func RcReparseAllForSettings(settings *Settings, forceLoad bool) bool {
	c_settings := (*C.GtkSettings)(C.NULL)
	if settings != nil {
		c_settings = (*C.GtkSettings)(settings.ToC())
	}

	c_force_load :=
		boolToGboolean(forceLoad)

	retC := C.gtk_rc_reparse_all_for_settings(c_settings, c_force_load)
	retGo := retC == C.TRUE

	return retGo
}

/*

C function

gtk_rc_scanner_new
*/
func RcScannerNew() *glib.Scanner {
	retC := C.gtk_rc_scanner_new()
	retGo := glib.ScannerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_rc_set_default_files : unsupported parameter filenames :

/*

C function

gtk_recent_chooser_error_quark
*/
func RecentChooserErrorQuark() glib.Quark {
	retC := C.gtk_recent_chooser_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

/*

C function

gtk_recent_manager_error_quark
*/
func RecentManagerErrorQuark() glib.Quark {
	retC := C.gtk_recent_manager_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : gtk_selection_add_target : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_add_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_clear_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Removes all handlers and unsets ownership of all
// selections for a widget. Called when widget is being
// destroyed. This function will not generally be
// called by applications.
/*

C function

gtk_selection_remove_all
*/
func SelectionRemoveAll(widget *Widget) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	C.gtk_selection_remove_all(c_widget)

	return
}

// Sets the GTK+ debug flags.
/*

C function

gtk_set_debug_flags
*/
func SetDebugFlags(flags uint32) {
	c_flags := (C.guint)(flags)

	C.gtk_set_debug_flags(c_flags)

	return
}

// Unsupported : gtk_stock_add : unsupported parameter items :

// Unsupported : gtk_stock_add_static : unsupported parameter items :

// Retrieves a list of all known stock IDs added to a #GtkIconFactory
// or registered with gtk_stock_add(). The list must be freed with g_slist_free(),
// and each string in the list must be freed with g_free().
/*

C function

gtk_stock_list_ids
*/
func StockListIds() *glib.SList {
	retC := C.gtk_stock_list_ids()
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Fills @item with the registered values for @stock_id, returning %TRUE
// if @stock_id was known.
/*

C function

gtk_stock_lookup
*/
func StockLookup(stockId string) (bool, *StockItem) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	var c_item C.GtkStockItem

	retC := C.gtk_stock_lookup(c_stock_id, &c_item)
	retGo := retC == C.TRUE

	item := StockItemNewFromC(unsafe.Pointer(&c_item))

	return retGo, item
}

// Obtains a @tree_model and @path from selection data of target type
// %GTK_TREE_MODEL_ROW. Normally called from a drag_data_received handler.
// This function can only be used if @selection_data originates from the same
// process that’s calling this function, because a pointer to the tree model
// is being passed around. If you aren’t in the same process, then you'll
// get memory corruption. In the #GtkTreeDragDest drag_data_received handler,
// you can assume that selection data of type %GTK_TREE_MODEL_ROW is
// in from the current process. The returned path must be freed with
// gtk_tree_path_free().
/*

C function

gtk_tree_get_row_drag_data
*/
func TreeGetRowDragData(selectionData *SelectionData) (bool, *TreeModel, *TreePath) {
	c_selection_data := (*C.GtkSelectionData)(C.NULL)
	if selectionData != nil {
		c_selection_data = (*C.GtkSelectionData)(selectionData.ToC())
	}

	var c_tree_model *C.GtkTreeModel

	var c_path *C.GtkTreePath

	retC := C.gtk_tree_get_row_drag_data(c_selection_data, &c_tree_model, &c_path)
	retGo := retC == C.TRUE

	treeModel := TreeModelNewFromC(unsafe.Pointer(c_tree_model))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	return retGo, treeModel, path
}

// Lets a set of row reference created by
// gtk_tree_row_reference_new_proxy() know that the
// model emitted the #GtkTreeModel::row-deleted signal.
/*

C function

gtk_tree_row_reference_deleted
*/
func TreeRowReferenceDeleted(proxy *gobject.Object, path *TreePath) {
	c_proxy := (*C.GObject)(C.NULL)
	if proxy != nil {
		c_proxy = (*C.GObject)(proxy.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_tree_row_reference_deleted(c_proxy, c_path)

	return
}

// Lets a set of row reference created by
// gtk_tree_row_reference_new_proxy() know that the
// model emitted the #GtkTreeModel::row-inserted signal.
/*

C function

gtk_tree_row_reference_inserted
*/
func TreeRowReferenceInserted(proxy *gobject.Object, path *TreePath) {
	c_proxy := (*C.GObject)(C.NULL)
	if proxy != nil {
		c_proxy = (*C.GObject)(proxy.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_tree_row_reference_inserted(c_proxy, c_path)

	return
}

// Lets a set of row reference created by
// gtk_tree_row_reference_new_proxy() know that the
// model emitted the #GtkTreeModel::rows-reordered signal.
/*

C function

gtk_tree_row_reference_reordered
*/
func TreeRowReferenceReordered(proxy *gobject.Object, path *TreePath, iter *TreeIter, newOrder []int32) {
	c_proxy := (*C.GObject)(C.NULL)
	if proxy != nil {
		c_proxy = (*C.GObject)(proxy.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_new_order := &newOrder[0]

	C.gtk_tree_row_reference_reordered(c_proxy, c_path, c_iter, (*C.gint)(unsafe.Pointer(c_new_order)))

	return
}

// Sets selection data of target type %GTK_TREE_MODEL_ROW. Normally used
// in a drag_data_get handler.
/*

C function

gtk_tree_set_row_drag_data
*/
func TreeSetRowDragData(selectionData *SelectionData, treeModel *TreeModel, path *TreePath) bool {
	c_selection_data := (*C.GtkSelectionData)(C.NULL)
	if selectionData != nil {
		c_selection_data = (*C.GtkSelectionData)(selectionData.ToC())
	}

	c_tree_model := (*C.GtkTreeModel)(treeModel.ToC())

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_tree_set_row_drag_data(c_selection_data, c_tree_model, c_path)
	retGo := retC == C.TRUE

	return retGo
}

// All this function does it to return %TRUE.
//
// This can be useful for example if you want to inhibit the deletion
// of a window. Of course you should not do this as the user expects
// a reaction from clicking the close icon of the window...
//
// A persistent window
//
// |[<!-- language="C" -->
// #include <gtk/gtk.h>
//
// int
// main (int argc, char **argv)
// {
// GtkWidget *win, *but;
// const char *text = "Close yourself. I mean it!";
//
// gtk_init (&argc, &argv);
//
// win = gtk_window_new (GTK_WINDOW_TOPLEVEL);
// g_signal_connect (win,
// "delete-event",
// G_CALLBACK (gtk_true),
// NULL);
// g_signal_connect (win, "destroy",
// G_CALLBACK (gtk_main_quit),
// NULL);
//
// but = gtk_button_new_with_label (text);
// g_signal_connect_swapped (but, "clicked",
// G_CALLBACK (gtk_object_destroy),
// win);
// gtk_container_add (GTK_CONTAINER (win), but);
//
// gtk_widget_show_all (win);
//
// gtk_main ();
//
// return 0;
// }
// ]|
/*

C function

gtk_true
*/
func True() bool {
	retC := C.gtk_true()
	retGo := retC == C.TRUE

	return retGo
}
