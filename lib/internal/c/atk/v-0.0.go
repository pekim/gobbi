// Code generated - DO NOT EDIT.
// +build !atk_2.7.4,!atk_2.12,!atk_2.14,!atk_2.30

package atk

import (
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
	"unsafe"
)

// #include <atk/atk.h>
import "C"

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

// UNSUPPORTED : add_focus_tracker : has callback

// UNSUPPORTED : add_global_event_listener : has callback

// UNSUPPORTED : add_key_event_listener : has callback

func Fn_attribute_set_free(param0 glib.SList) {}

// UNSUPPORTED : focus_tracker_init : has callback

func Fn_focus_tracker_notify(param0 unsafe.Pointer) {}

func Fn_get_default_registry() {
	C.atk_get_default_registry()
}

func Fn_get_root() {
	C.atk_get_root()
}

func Fn_get_toolkit_name() {
	C.atk_get_toolkit_name()
}

func Fn_get_toolkit_version() {
	C.atk_get_toolkit_version()
}

func Fn_relation_type_for_name(param0 string) {}

func Fn_relation_type_get_name(param0 int) {}

func Fn_relation_type_register(param0 string) {}

func Fn_remove_focus_tracker(param0 uint) {}

func Fn_remove_global_event_listener(param0 uint) {}

func Fn_remove_key_event_listener(param0 uint) {}

func Fn_role_for_name(param0 string) {}

func Fn_role_get_localized_name(param0 int) {}

func Fn_role_get_name(param0 int) {}

func Fn_role_register(param0 string) {}

func Fn_state_type_for_name(param0 string) {}

func Fn_state_type_get_name(param0 int) {}

func Fn_state_type_register(param0 string) {}

func Fn_text_attribute_for_name(param0 string) {}

func Fn_text_attribute_get_name(param0 int) {}

func Fn_text_attribute_get_value(param0 int, param1 int) {}

func Fn_text_attribute_register(param0 string) {}

func Fn_value_type_get_localized_name(param0 int) {}

func Fn_value_type_get_name(param0 int) {}
