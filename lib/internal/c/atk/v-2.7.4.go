// Code generated - DO NOT EDIT.
// +build atk_2.7.4

package atk

// #include <atk/atk.h>
import "C"

// bitfields
type HyperlinkStateFlags C.AtkHyperlinkStateFlags

// enumerations
type CoordType C.AtkCoordType
type KeyEventType C.AtkKeyEventType
type Layer C.AtkLayer
type RelationType C.AtkRelationType
type Role C.AtkRole
type StateType C.AtkStateType
type TextAttribute C.AtkTextAttribute
type TextBoundary C.AtkTextBoundary
type TextClipType C.AtkTextClipType
type TextGranularity C.AtkTextGranularity
type ValueType C.AtkValueType

// records
type ActionIface C.AtkActionIface
type Attribute C.AtkAttribute
type ComponentIface C.AtkComponentIface
type DocumentIface C.AtkDocumentIface
type EditableTextIface C.AtkEditableTextIface
type GObjectAccessibleClass C.AtkGObjectAccessibleClass
type HyperlinkClass C.AtkHyperlinkClass
type HyperlinkImplIface C.AtkHyperlinkImplIface
type HypertextIface C.AtkHypertextIface
type ImageIface C.AtkImageIface
type Implementor C.AtkImplementor
type KeyEventStruct C.AtkKeyEventStruct
type MiscClass C.AtkMiscClass
type NoOpObjectClass C.AtkNoOpObjectClass
type NoOpObjectFactoryClass C.AtkNoOpObjectFactoryClass
type ObjectClass C.AtkObjectClass
type ObjectFactoryClass C.AtkObjectFactoryClass
type PlugClass C.AtkPlugClass
type PropertyValues C.AtkPropertyValues
type Range C.AtkRange
type Rectangle C.AtkRectangle
type RegistryClass C.AtkRegistryClass
type RelationClass C.AtkRelationClass
type RelationSetClass C.AtkRelationSetClass
type SelectionIface C.AtkSelectionIface
type SocketClass C.AtkSocketClass
type StateSetClass C.AtkStateSetClass
type StreamableContentIface C.AtkStreamableContentIface
type TableIface C.AtkTableIface
type TextIface C.AtkTextIface
type TextRange C.AtkTextRange
type TextRectangle C.AtkTextRectangle
type UtilClass C.AtkUtilClass
type ValueIface C.AtkValueIface
type WindowIface C.AtkWindowIface

// classes
type GObjectAccessible C.AtkGObjectAccessible
type Hyperlink C.AtkHyperlink
type Misc C.AtkMisc
type NoOpObject C.AtkNoOpObject
type NoOpObjectFactory C.AtkNoOpObjectFactory
type Object C.AtkObject
type ObjectFactory C.AtkObjectFactory
type Plug C.AtkPlug
type Registry C.AtkRegistry
type Relation C.AtkRelation
type RelationSet C.AtkRelationSet
type Socket C.AtkSocket
type StateSet C.AtkStateSet
type Util C.AtkUtil

// interfaces
type Action C.AtkAction
type Component C.AtkComponent
type Document C.AtkDocument
type EditableText C.AtkEditableText
type HyperlinkImpl C.AtkHyperlinkImpl
type Hypertext C.AtkHypertext
type Image C.AtkImage
type ImplementorIface C.AtkImplementorIface
type Selection C.AtkSelection
type StreamableContent C.AtkStreamableContent
type Table C.AtkTable
type TableCell C.AtkTableCell
type Text C.AtkText
type Value C.AtkValue
type Window C.AtkWindow

func Fn_add_focus_tracker(focusTracker string) {}

func Fn_add_global_event_listener(listener string, eventType string) {}

func Fn_add_key_event_listener(listener string, data string) {}

func Fn_attribute_set_free(attribSet string) {}

func Fn_focus_tracker_init(init string) {}

func Fn_focus_tracker_notify(object string) {}

func Fn_get_default_registry() {}

func Fn_get_focus_object() {}

func Fn_get_root() {}

func Fn_get_toolkit_name() {}

func Fn_get_toolkit_version() {}

func Fn_get_version() {}

func Fn_relation_type_for_name(name string) {}

func Fn_relation_type_get_name(type_ string) {}

func Fn_relation_type_register(name string) {}

func Fn_remove_focus_tracker(trackerId string) {}

func Fn_remove_global_event_listener(listenerId string) {}

func Fn_remove_key_event_listener(listenerId string) {}

func Fn_role_for_name(name string) {}

func Fn_role_get_localized_name(role string) {}

func Fn_role_get_name(role string) {}

func Fn_role_register(name string) {}

func Fn_state_type_for_name(name string) {}

func Fn_state_type_get_name(type_ string) {}

func Fn_state_type_register(name string) {}

func Fn_text_attribute_for_name(name string) {}

func Fn_text_attribute_get_name(attr string) {}

func Fn_text_attribute_get_value(attr string, index string) {}

func Fn_text_attribute_register(name string) {}

func Fn_text_free_ranges(ranges string) {}

func Fn_value_type_get_localized_name(valueType string) {}

func Fn_value_type_get_name(valueType string) {}
