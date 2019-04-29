// This is a generated file - DO NOT EDIT

package atk

import "unsafe"

// GObjectAccessible is a wrapper around the C record AtkGObjectAccessible.
type GObjectAccessible struct {
	native unsafe.Pointer
	// parent : record
}

func GObjectAccessibleNewFromC(u unsafe.Pointer) *GObjectAccessible {
	if u == nil {
		return nil
	}

	g := &GObjectAccessible{native: u}

	return g
}

func (recv *GObjectAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Hyperlink is a wrapper around the C record AtkHyperlink.
type Hyperlink struct {
	native unsafe.Pointer
	// parent : record
}

func HyperlinkNewFromC(u unsafe.Pointer) *Hyperlink {
	if u == nil {
		return nil
	}

	g := &Hyperlink{native: u}

	return g
}

func (recv *Hyperlink) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Action returns the Action interface implemented by Hyperlink
func (recv *Hyperlink) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// Misc is a wrapper around the C record AtkMisc.
type Misc struct {
	native unsafe.Pointer
	// parent : record
}

func MiscNewFromC(u unsafe.Pointer) *Misc {
	if u == nil {
		return nil
	}

	g := &Misc{native: u}

	return g
}

func (recv *Misc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NoOpObject is a wrapper around the C record AtkNoOpObject.
type NoOpObject struct {
	native unsafe.Pointer
	// parent : record
}

func NoOpObjectNewFromC(u unsafe.Pointer) *NoOpObject {
	if u == nil {
		return nil
	}

	g := &NoOpObject{native: u}

	return g
}

func (recv *NoOpObject) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Action returns the Action interface implemented by NoOpObject
func (recv *NoOpObject) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by NoOpObject
func (recv *NoOpObject) Component() *Component {
	return ComponentNewFromC(recv.ToC())
}

// Document returns the Document interface implemented by NoOpObject
func (recv *NoOpObject) Document() *Document {
	return DocumentNewFromC(recv.ToC())
}

// EditableText returns the EditableText interface implemented by NoOpObject
func (recv *NoOpObject) EditableText() *EditableText {
	return EditableTextNewFromC(recv.ToC())
}

// Hypertext returns the Hypertext interface implemented by NoOpObject
func (recv *NoOpObject) Hypertext() *Hypertext {
	return HypertextNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by NoOpObject
func (recv *NoOpObject) Image() *Image {
	return ImageNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by NoOpObject
func (recv *NoOpObject) Selection() *Selection {
	return SelectionNewFromC(recv.ToC())
}

// Table returns the Table interface implemented by NoOpObject
func (recv *NoOpObject) Table() *Table {
	return TableNewFromC(recv.ToC())
}

// TableCell returns the TableCell interface implemented by NoOpObject
func (recv *NoOpObject) TableCell() *TableCell {
	return TableCellNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by NoOpObject
func (recv *NoOpObject) Text() *Text {
	return TextNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by NoOpObject
func (recv *NoOpObject) Value() *Value {
	return ValueNewFromC(recv.ToC())
}

// Window returns the Window interface implemented by NoOpObject
func (recv *NoOpObject) Window() *Window {
	return WindowNewFromC(recv.ToC())
}

// NoOpObjectFactory is a wrapper around the C record AtkNoOpObjectFactory.
type NoOpObjectFactory struct {
	native unsafe.Pointer
	// parent : record
}

func NoOpObjectFactoryNewFromC(u unsafe.Pointer) *NoOpObjectFactory {
	if u == nil {
		return nil
	}

	g := &NoOpObjectFactory{native: u}

	return g
}

func (recv *NoOpObjectFactory) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object is a wrapper around the C record AtkObject.
type Object struct {
	native unsafe.Pointer
	// parent : record
	Description string
	Name        string
	// accessible_parent : record
	Role Role
	// relation_set : record
	Layer Layer
}

func ObjectNewFromC(u unsafe.Pointer) *Object {
	if u == nil {
		return nil
	}

	g := &Object{native: u}

	return g
}

func (recv *Object) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ObjectFactory is a wrapper around the C record AtkObjectFactory.
type ObjectFactory struct {
	native unsafe.Pointer
	// parent : record
}

func ObjectFactoryNewFromC(u unsafe.Pointer) *ObjectFactory {
	if u == nil {
		return nil
	}

	g := &ObjectFactory{native: u}

	return g
}

func (recv *ObjectFactory) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Plug is a wrapper around the C record AtkPlug.
type Plug struct {
	native unsafe.Pointer
	// parent : record
}

func PlugNewFromC(u unsafe.Pointer) *Plug {
	if u == nil {
		return nil
	}

	g := &Plug{native: u}

	return g
}

func (recv *Plug) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Component returns the Component interface implemented by Plug
func (recv *Plug) Component() *Component {
	return ComponentNewFromC(recv.ToC())
}

// Registry is a wrapper around the C record AtkRegistry.
type Registry struct {
	native unsafe.Pointer
	// parent : record
	// factory_type_registry : record
	// factory_singleton_cache : record
}

func RegistryNewFromC(u unsafe.Pointer) *Registry {
	if u == nil {
		return nil
	}

	g := &Registry{native: u}

	return g
}

func (recv *Registry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Relation is a wrapper around the C record AtkRelation.
type Relation struct {
	native unsafe.Pointer
	// parent : record
	// no type for target
	Relationship RelationType
}

func RelationNewFromC(u unsafe.Pointer) *Relation {
	if u == nil {
		return nil
	}

	g := &Relation{native: u}

	return g
}

func (recv *Relation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RelationSet is a wrapper around the C record AtkRelationSet.
type RelationSet struct {
	native unsafe.Pointer
	// parent : record
	// no type for relations
}

func RelationSetNewFromC(u unsafe.Pointer) *RelationSet {
	if u == nil {
		return nil
	}

	g := &RelationSet{native: u}

	return g
}

func (recv *RelationSet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Socket is a wrapper around the C record AtkSocket.
type Socket struct {
	native unsafe.Pointer
	// parent : record
	// Private : embedded_plug_id
}

func SocketNewFromC(u unsafe.Pointer) *Socket {
	if u == nil {
		return nil
	}

	g := &Socket{native: u}

	return g
}

func (recv *Socket) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Component returns the Component interface implemented by Socket
func (recv *Socket) Component() *Component {
	return ComponentNewFromC(recv.ToC())
}

// StateSet is a wrapper around the C record AtkStateSet.
type StateSet struct {
	native unsafe.Pointer
	// parent : record
}

func StateSetNewFromC(u unsafe.Pointer) *StateSet {
	if u == nil {
		return nil
	}

	g := &StateSet{native: u}

	return g
}

func (recv *StateSet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Util is a wrapper around the C record AtkUtil.
type Util struct {
	native unsafe.Pointer
	// parent : record
}

func UtilNewFromC(u unsafe.Pointer) *Util {
	if u == nil {
		return nil
	}

	g := &Util{native: u}

	return g
}

func (recv *Util) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
