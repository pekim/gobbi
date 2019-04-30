// This is a generated file - DO NOT EDIT

package atk

// Action returns the Action interface implemented by Hyperlink
func (recv *Hyperlink) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// Unsupported : atk_no_op_object_new : return type :

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

// Unsupported : atk_no_op_object_factory_new : return type :

// Unsupported : atk_plug_new : return type :

// Component returns the Component interface implemented by Plug
func (recv *Plug) Component() *Component {
	return ComponentNewFromC(recv.ToC())
}

// Unsupported : atk_relation_new : unsupported parameter targets :

// Unsupported : atk_relation_set_new : return type :

// Unsupported : atk_socket_new : return type :

// Component returns the Component interface implemented by Socket
func (recv *Socket) Component() *Component {
	return ComponentNewFromC(recv.ToC())
}

// Unsupported : atk_state_set_new : return type :
