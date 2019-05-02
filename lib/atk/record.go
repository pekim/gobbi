// This is a generated file - DO NOT EDIT

package atk

import (
	"C"
	"unsafe"
)

// Equals compares this ActionIface with another ActionIface, and returns true if they represent the same GObject.
func (recv *ActionIface) Equals(other *ActionIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Attribute with another Attribute, and returns true if they represent the same GObject.
func (recv *Attribute) Equals(other *Attribute) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : atk_attribute_set_free

// Equals compares this ComponentIface with another ComponentIface, and returns true if they represent the same GObject.
func (recv *ComponentIface) Equals(other *ComponentIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DocumentIface with another DocumentIface, and returns true if they represent the same GObject.
func (recv *DocumentIface) Equals(other *DocumentIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EditableTextIface with another EditableTextIface, and returns true if they represent the same GObject.
func (recv *EditableTextIface) Equals(other *EditableTextIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this GObjectAccessibleClass with another GObjectAccessibleClass, and returns true if they represent the same GObject.
func (recv *GObjectAccessibleClass) Equals(other *GObjectAccessibleClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this HyperlinkClass with another HyperlinkClass, and returns true if they represent the same GObject.
func (recv *HyperlinkClass) Equals(other *HyperlinkClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this HyperlinkImplIface with another HyperlinkImplIface, and returns true if they represent the same GObject.
func (recv *HyperlinkImplIface) Equals(other *HyperlinkImplIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this HypertextIface with another HypertextIface, and returns true if they represent the same GObject.
func (recv *HypertextIface) Equals(other *HypertextIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ImageIface with another ImageIface, and returns true if they represent the same GObject.
func (recv *ImageIface) Equals(other *ImageIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Implementor with another Implementor, and returns true if they represent the same GObject.
func (recv *Implementor) Equals(other *Implementor) bool {
	return other.ToC() == recv.ToC()
}

// RefAccessible is a wrapper around the C function atk_implementor_ref_accessible.
func (recv *Implementor) RefAccessible() *Object {
	retC := C.atk_implementor_ref_accessible((*C.AtkImplementor)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equals compares this KeyEventStruct with another KeyEventStruct, and returns true if they represent the same GObject.
func (recv *KeyEventStruct) Equals(other *KeyEventStruct) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this MiscClass with another MiscClass, and returns true if they represent the same GObject.
func (recv *MiscClass) Equals(other *MiscClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this NoOpObjectClass with another NoOpObjectClass, and returns true if they represent the same GObject.
func (recv *NoOpObjectClass) Equals(other *NoOpObjectClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this NoOpObjectFactoryClass with another NoOpObjectFactoryClass, and returns true if they represent the same GObject.
func (recv *NoOpObjectFactoryClass) Equals(other *NoOpObjectFactoryClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ObjectClass with another ObjectClass, and returns true if they represent the same GObject.
func (recv *ObjectClass) Equals(other *ObjectClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ObjectFactoryClass with another ObjectFactoryClass, and returns true if they represent the same GObject.
func (recv *ObjectFactoryClass) Equals(other *ObjectFactoryClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this PlugClass with another PlugClass, and returns true if they represent the same GObject.
func (recv *PlugClass) Equals(other *PlugClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this PropertyValues with another PropertyValues, and returns true if they represent the same GObject.
func (recv *PropertyValues) Equals(other *PropertyValues) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Range with another Range, and returns true if they represent the same GObject.
func (recv *Range) Equals(other *Range) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Rectangle with another Rectangle, and returns true if they represent the same GObject.
func (recv *Rectangle) Equals(other *Rectangle) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this RegistryClass with another RegistryClass, and returns true if they represent the same GObject.
func (recv *RegistryClass) Equals(other *RegistryClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this RelationClass with another RelationClass, and returns true if they represent the same GObject.
func (recv *RelationClass) Equals(other *RelationClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this RelationSetClass with another RelationSetClass, and returns true if they represent the same GObject.
func (recv *RelationSetClass) Equals(other *RelationSetClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this SelectionIface with another SelectionIface, and returns true if they represent the same GObject.
func (recv *SelectionIface) Equals(other *SelectionIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this SocketClass with another SocketClass, and returns true if they represent the same GObject.
func (recv *SocketClass) Equals(other *SocketClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this StateSetClass with another StateSetClass, and returns true if they represent the same GObject.
func (recv *StateSetClass) Equals(other *StateSetClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this StreamableContentIface with another StreamableContentIface, and returns true if they represent the same GObject.
func (recv *StreamableContentIface) Equals(other *StreamableContentIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this TableCellIface with another TableCellIface, and returns true if they represent the same GObject.
func (recv *TableCellIface) Equals(other *TableCellIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this TableIface with another TableIface, and returns true if they represent the same GObject.
func (recv *TableIface) Equals(other *TableIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this TextIface with another TextIface, and returns true if they represent the same GObject.
func (recv *TextIface) Equals(other *TextIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this TextRange with another TextRange, and returns true if they represent the same GObject.
func (recv *TextRange) Equals(other *TextRange) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this TextRectangle with another TextRectangle, and returns true if they represent the same GObject.
func (recv *TextRectangle) Equals(other *TextRectangle) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this UtilClass with another UtilClass, and returns true if they represent the same GObject.
func (recv *UtilClass) Equals(other *UtilClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ValueIface with another ValueIface, and returns true if they represent the same GObject.
func (recv *ValueIface) Equals(other *ValueIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this WindowIface with another WindowIface, and returns true if they represent the same GObject.
func (recv *WindowIface) Equals(other *WindowIface) bool {
	return other.ToC() == recv.ToC()
}
