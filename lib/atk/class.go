// This is a generated file - DO NOT EDIT

package atk

import "unsafe"

// GObjectAccessible is a wrapper around the C record AtkGObjectAccessible.
type GObjectAccessible struct {
	native unsafe.Pointer
	// parent : record
}

// Hyperlink is a wrapper around the C record AtkHyperlink.
type Hyperlink struct {
	native unsafe.Pointer
	// parent : record
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

// NoOpObject is a wrapper around the C record AtkNoOpObject.
type NoOpObject struct {
	native unsafe.Pointer
	// parent : record
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

// ObjectFactory is a wrapper around the C record AtkObjectFactory.
type ObjectFactory struct {
	native unsafe.Pointer
	// parent : record
}

// Plug is a wrapper around the C record AtkPlug.
type Plug struct {
	native unsafe.Pointer
	// parent : record
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

// Relation is a wrapper around the C record AtkRelation.
type Relation struct {
	native unsafe.Pointer
	// parent : record
	// no type for target
	Relationship RelationType
}

// RelationSet is a wrapper around the C record AtkRelationSet.
type RelationSet struct {
	native unsafe.Pointer
	// parent : record
	// no type for relations
}

// Socket is a wrapper around the C record AtkSocket.
type Socket struct {
	native unsafe.Pointer
	// parent : record
	// Private : embedded_plug_id
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

// Util is a wrapper around the C record AtkUtil.
type Util struct {
	native unsafe.Pointer
	// parent : record
}
