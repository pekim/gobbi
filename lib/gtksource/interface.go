// Code generated - DO NOT EDIT.

package gtksource

import (
	callback "github.com/pekim/gobbi/internal/cgo/callback"
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

var completionProposalInterface *gi.Interface
var completionProposalInterface_Once sync.Once

func completionProposalInterface_Set() error {
	var err error
	completionProposalInterface_Once.Do(func() {
		completionProposalInterface, err = gi.InterfaceNew("GtkSource", "CompletionProposal")
	})
	return err
}

type CompletionProposal struct {
	native unsafe.Pointer
}

func CompletionProposalNewFromNative(native unsafe.Pointer) *CompletionProposal {
	instance := &CompletionProposal{native: native}

	return instance
}

/*
CastToCompletionProposal down casts any arbitrary Object to CompletionProposal.
Exercise care, as this is a potentially dangerous function
if the Object is not a CompletionProposal.
*/
func CastToCompletionProposal(object *gobject.Object) *CompletionProposal {
	return CompletionProposalNewFromNative(object.Native())
}

// Equals compares this CompletionProposal with another CompletionProposal, and returns true if they represent the same Object.
func (recv *CompletionProposal) Equals(other *CompletionProposal) bool {
	return other.Native() == recv.Native()
}

func (recv *CompletionProposal) Native() unsafe.Pointer {
	return recv.native
}

var completionProposalChangedFunction *gi.Function
var completionProposalChangedFunction_Once sync.Once

func completionProposalChangedFunction_Set() error {
	var err error
	completionProposalChangedFunction_Once.Do(func() {
		err = completionProposalInterface_Set()
		if err != nil {
			return
		}
		completionProposalChangedFunction, err = completionProposalInterface.InvokerNew("changed")
	})
	return err
}

// Changed is a representation of the C type gtk_source_completion_proposal_changed.
func (recv *CompletionProposal) Changed() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := completionProposalChangedFunction_Set()
	if err == nil {
		completionProposalChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_completion_proposal_equal' : parameter 'other' of type 'CompletionProposal' not supported

// UNSUPPORTED : C value 'gtk_source_completion_proposal_get_gicon' : return type 'Gio.Icon' not supported

var completionProposalGetIconFunction *gi.Function
var completionProposalGetIconFunction_Once sync.Once

func completionProposalGetIconFunction_Set() error {
	var err error
	completionProposalGetIconFunction_Once.Do(func() {
		err = completionProposalInterface_Set()
		if err != nil {
			return
		}
		completionProposalGetIconFunction, err = completionProposalInterface.InvokerNew("get_icon")
	})
	return err
}

// GetIcon is a representation of the C type gtk_source_completion_proposal_get_icon.
func (recv *CompletionProposal) GetIcon() *gdkpixbuf.Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProposalGetIconFunction_Set()
	if err == nil {
		ret = completionProposalGetIconFunction.Invoke(inArgs[:], nil)
	}

	retGo := gdkpixbuf.PixbufNewFromNative(ret.Pointer())

	return retGo
}

var completionProposalGetIconNameFunction *gi.Function
var completionProposalGetIconNameFunction_Once sync.Once

func completionProposalGetIconNameFunction_Set() error {
	var err error
	completionProposalGetIconNameFunction_Once.Do(func() {
		err = completionProposalInterface_Set()
		if err != nil {
			return
		}
		completionProposalGetIconNameFunction, err = completionProposalInterface.InvokerNew("get_icon_name")
	})
	return err
}

// GetIconName is a representation of the C type gtk_source_completion_proposal_get_icon_name.
func (recv *CompletionProposal) GetIconName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProposalGetIconNameFunction_Set()
	if err == nil {
		ret = completionProposalGetIconNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var completionProposalGetInfoFunction *gi.Function
var completionProposalGetInfoFunction_Once sync.Once

func completionProposalGetInfoFunction_Set() error {
	var err error
	completionProposalGetInfoFunction_Once.Do(func() {
		err = completionProposalInterface_Set()
		if err != nil {
			return
		}
		completionProposalGetInfoFunction, err = completionProposalInterface.InvokerNew("get_info")
	})
	return err
}

// GetInfo is a representation of the C type gtk_source_completion_proposal_get_info.
func (recv *CompletionProposal) GetInfo() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProposalGetInfoFunction_Set()
	if err == nil {
		ret = completionProposalGetInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var completionProposalGetLabelFunction *gi.Function
var completionProposalGetLabelFunction_Once sync.Once

func completionProposalGetLabelFunction_Set() error {
	var err error
	completionProposalGetLabelFunction_Once.Do(func() {
		err = completionProposalInterface_Set()
		if err != nil {
			return
		}
		completionProposalGetLabelFunction, err = completionProposalInterface.InvokerNew("get_label")
	})
	return err
}

// GetLabel is a representation of the C type gtk_source_completion_proposal_get_label.
func (recv *CompletionProposal) GetLabel() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProposalGetLabelFunction_Set()
	if err == nil {
		ret = completionProposalGetLabelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var completionProposalGetMarkupFunction *gi.Function
var completionProposalGetMarkupFunction_Once sync.Once

func completionProposalGetMarkupFunction_Set() error {
	var err error
	completionProposalGetMarkupFunction_Once.Do(func() {
		err = completionProposalInterface_Set()
		if err != nil {
			return
		}
		completionProposalGetMarkupFunction, err = completionProposalInterface.InvokerNew("get_markup")
	})
	return err
}

// GetMarkup is a representation of the C type gtk_source_completion_proposal_get_markup.
func (recv *CompletionProposal) GetMarkup() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProposalGetMarkupFunction_Set()
	if err == nil {
		ret = completionProposalGetMarkupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var completionProposalGetTextFunction *gi.Function
var completionProposalGetTextFunction_Once sync.Once

func completionProposalGetTextFunction_Set() error {
	var err error
	completionProposalGetTextFunction_Once.Do(func() {
		err = completionProposalInterface_Set()
		if err != nil {
			return
		}
		completionProposalGetTextFunction, err = completionProposalInterface.InvokerNew("get_text")
	})
	return err
}

// GetText is a representation of the C type gtk_source_completion_proposal_get_text.
func (recv *CompletionProposal) GetText() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProposalGetTextFunction_Set()
	if err == nil {
		ret = completionProposalGetTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var completionProposalHashFunction *gi.Function
var completionProposalHashFunction_Once sync.Once

func completionProposalHashFunction_Set() error {
	var err error
	completionProposalHashFunction_Once.Do(func() {
		err = completionProposalInterface_Set()
		if err != nil {
			return
		}
		completionProposalHashFunction, err = completionProposalInterface.InvokerNew("hash")
	})
	return err
}

// Hash is a representation of the C type gtk_source_completion_proposal_hash.
func (recv *CompletionProposal) Hash() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProposalHashFunction_Set()
	if err == nil {
		ret = completionProposalHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *CompletionProposal) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var completionProviderInterface *gi.Interface
var completionProviderInterface_Once sync.Once

func completionProviderInterface_Set() error {
	var err error
	completionProviderInterface_Once.Do(func() {
		completionProviderInterface, err = gi.InterfaceNew("GtkSource", "CompletionProvider")
	})
	return err
}

type CompletionProvider struct {
	native unsafe.Pointer
}

func CompletionProviderNewFromNative(native unsafe.Pointer) *CompletionProvider {
	instance := &CompletionProvider{native: native}

	return instance
}

/*
CastToCompletionProvider down casts any arbitrary Object to CompletionProvider.
Exercise care, as this is a potentially dangerous function
if the Object is not a CompletionProvider.
*/
func CastToCompletionProvider(object *gobject.Object) *CompletionProvider {
	return CompletionProviderNewFromNative(object.Native())
}

// Equals compares this CompletionProvider with another CompletionProvider, and returns true if they represent the same Object.
func (recv *CompletionProvider) Equals(other *CompletionProvider) bool {
	return other.Native() == recv.Native()
}

func (recv *CompletionProvider) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'gtk_source_completion_provider_activate_proposal' : parameter 'proposal' of type 'CompletionProposal' not supported

// UNSUPPORTED : C value 'gtk_source_completion_provider_get_activation' : return type 'CompletionActivation' not supported

// UNSUPPORTED : C value 'gtk_source_completion_provider_get_gicon' : return type 'Gio.Icon' not supported

var completionProviderGetIconFunction *gi.Function
var completionProviderGetIconFunction_Once sync.Once

func completionProviderGetIconFunction_Set() error {
	var err error
	completionProviderGetIconFunction_Once.Do(func() {
		err = completionProviderInterface_Set()
		if err != nil {
			return
		}
		completionProviderGetIconFunction, err = completionProviderInterface.InvokerNew("get_icon")
	})
	return err
}

// GetIcon is a representation of the C type gtk_source_completion_provider_get_icon.
func (recv *CompletionProvider) GetIcon() *gdkpixbuf.Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProviderGetIconFunction_Set()
	if err == nil {
		ret = completionProviderGetIconFunction.Invoke(inArgs[:], nil)
	}

	retGo := gdkpixbuf.PixbufNewFromNative(ret.Pointer())

	return retGo
}

var completionProviderGetIconNameFunction *gi.Function
var completionProviderGetIconNameFunction_Once sync.Once

func completionProviderGetIconNameFunction_Set() error {
	var err error
	completionProviderGetIconNameFunction_Once.Do(func() {
		err = completionProviderInterface_Set()
		if err != nil {
			return
		}
		completionProviderGetIconNameFunction, err = completionProviderInterface.InvokerNew("get_icon_name")
	})
	return err
}

// GetIconName is a representation of the C type gtk_source_completion_provider_get_icon_name.
func (recv *CompletionProvider) GetIconName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProviderGetIconNameFunction_Set()
	if err == nil {
		ret = completionProviderGetIconNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_completion_provider_get_info_widget' : parameter 'proposal' of type 'CompletionProposal' not supported

var completionProviderGetInteractiveDelayFunction *gi.Function
var completionProviderGetInteractiveDelayFunction_Once sync.Once

func completionProviderGetInteractiveDelayFunction_Set() error {
	var err error
	completionProviderGetInteractiveDelayFunction_Once.Do(func() {
		err = completionProviderInterface_Set()
		if err != nil {
			return
		}
		completionProviderGetInteractiveDelayFunction, err = completionProviderInterface.InvokerNew("get_interactive_delay")
	})
	return err
}

// GetInteractiveDelay is a representation of the C type gtk_source_completion_provider_get_interactive_delay.
func (recv *CompletionProvider) GetInteractiveDelay() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProviderGetInteractiveDelayFunction_Set()
	if err == nil {
		ret = completionProviderGetInteractiveDelayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var completionProviderGetNameFunction *gi.Function
var completionProviderGetNameFunction_Once sync.Once

func completionProviderGetNameFunction_Set() error {
	var err error
	completionProviderGetNameFunction_Once.Do(func() {
		err = completionProviderInterface_Set()
		if err != nil {
			return
		}
		completionProviderGetNameFunction, err = completionProviderInterface.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gtk_source_completion_provider_get_name.
func (recv *CompletionProvider) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProviderGetNameFunction_Set()
	if err == nil {
		ret = completionProviderGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var completionProviderGetPriorityFunction *gi.Function
var completionProviderGetPriorityFunction_Once sync.Once

func completionProviderGetPriorityFunction_Set() error {
	var err error
	completionProviderGetPriorityFunction_Once.Do(func() {
		err = completionProviderInterface_Set()
		if err != nil {
			return
		}
		completionProviderGetPriorityFunction, err = completionProviderInterface.InvokerNew("get_priority")
	})
	return err
}

// GetPriority is a representation of the C type gtk_source_completion_provider_get_priority.
func (recv *CompletionProvider) GetPriority() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := completionProviderGetPriorityFunction_Set()
	if err == nil {
		ret = completionProviderGetPriorityFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_source_completion_provider_get_start_iter' : parameter 'proposal' of type 'CompletionProposal' not supported

var completionProviderMatchFunction *gi.Function
var completionProviderMatchFunction_Once sync.Once

func completionProviderMatchFunction_Set() error {
	var err error
	completionProviderMatchFunction_Once.Do(func() {
		err = completionProviderInterface_Set()
		if err != nil {
			return
		}
		completionProviderMatchFunction, err = completionProviderInterface.InvokerNew("match")
	})
	return err
}

// Match is a representation of the C type gtk_source_completion_provider_match.
func (recv *CompletionProvider) Match(context *CompletionContext) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(context.Native())

	var ret gi.Argument

	err := completionProviderMatchFunction_Set()
	if err == nil {
		ret = completionProviderMatchFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var completionProviderPopulateFunction *gi.Function
var completionProviderPopulateFunction_Once sync.Once

func completionProviderPopulateFunction_Set() error {
	var err error
	completionProviderPopulateFunction_Once.Do(func() {
		err = completionProviderInterface_Set()
		if err != nil {
			return
		}
		completionProviderPopulateFunction, err = completionProviderInterface.InvokerNew("populate")
	})
	return err
}

// Populate is a representation of the C type gtk_source_completion_provider_populate.
func (recv *CompletionProvider) Populate(context *CompletionContext) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(context.Native())

	err := completionProviderPopulateFunction_Set()
	if err == nil {
		completionProviderPopulateFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_source_completion_provider_update_info' : parameter 'proposal' of type 'CompletionProposal' not supported

var styleSchemeChooserInterface *gi.Interface
var styleSchemeChooserInterface_Once sync.Once

func styleSchemeChooserInterface_Set() error {
	var err error
	styleSchemeChooserInterface_Once.Do(func() {
		styleSchemeChooserInterface, err = gi.InterfaceNew("GtkSource", "StyleSchemeChooser")
	})
	return err
}

type StyleSchemeChooser struct {
	native unsafe.Pointer
}

func StyleSchemeChooserNewFromNative(native unsafe.Pointer) *StyleSchemeChooser {
	instance := &StyleSchemeChooser{native: native}

	return instance
}

/*
CastToStyleSchemeChooser down casts any arbitrary Object to StyleSchemeChooser.
Exercise care, as this is a potentially dangerous function
if the Object is not a StyleSchemeChooser.
*/
func CastToStyleSchemeChooser(object *gobject.Object) *StyleSchemeChooser {
	return StyleSchemeChooserNewFromNative(object.Native())
}

// Equals compares this StyleSchemeChooser with another StyleSchemeChooser, and returns true if they represent the same Object.
func (recv *StyleSchemeChooser) Equals(other *StyleSchemeChooser) bool {
	return other.Native() == recv.Native()
}

func (recv *StyleSchemeChooser) Native() unsafe.Pointer {
	return recv.native
}

var styleSchemeChooserGetStyleSchemeFunction *gi.Function
var styleSchemeChooserGetStyleSchemeFunction_Once sync.Once

func styleSchemeChooserGetStyleSchemeFunction_Set() error {
	var err error
	styleSchemeChooserGetStyleSchemeFunction_Once.Do(func() {
		err = styleSchemeChooserInterface_Set()
		if err != nil {
			return
		}
		styleSchemeChooserGetStyleSchemeFunction, err = styleSchemeChooserInterface.InvokerNew("get_style_scheme")
	})
	return err
}

// GetStyleScheme is a representation of the C type gtk_source_style_scheme_chooser_get_style_scheme.
func (recv *StyleSchemeChooser) GetStyleScheme() *StyleScheme {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := styleSchemeChooserGetStyleSchemeFunction_Set()
	if err == nil {
		ret = styleSchemeChooserGetStyleSchemeFunction.Invoke(inArgs[:], nil)
	}

	retGo := StyleSchemeNewFromNative(ret.Pointer())

	return retGo
}

var styleSchemeChooserSetStyleSchemeFunction *gi.Function
var styleSchemeChooserSetStyleSchemeFunction_Once sync.Once

func styleSchemeChooserSetStyleSchemeFunction_Set() error {
	var err error
	styleSchemeChooserSetStyleSchemeFunction_Once.Do(func() {
		err = styleSchemeChooserInterface_Set()
		if err != nil {
			return
		}
		styleSchemeChooserSetStyleSchemeFunction, err = styleSchemeChooserInterface.InvokerNew("set_style_scheme")
	})
	return err
}

// SetStyleScheme is a representation of the C type gtk_source_style_scheme_chooser_set_style_scheme.
func (recv *StyleSchemeChooser) SetStyleScheme(scheme *StyleScheme) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(scheme.Native())

	err := styleSchemeChooserSetStyleSchemeFunction_Set()
	if err == nil {
		styleSchemeChooserSetStyleSchemeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var undoManagerInterface *gi.Interface
var undoManagerInterface_Once sync.Once

func undoManagerInterface_Set() error {
	var err error
	undoManagerInterface_Once.Do(func() {
		undoManagerInterface, err = gi.InterfaceNew("GtkSource", "UndoManager")
	})
	return err
}

type UndoManager struct {
	native unsafe.Pointer
}

func UndoManagerNewFromNative(native unsafe.Pointer) *UndoManager {
	instance := &UndoManager{native: native}

	return instance
}

/*
CastToUndoManager down casts any arbitrary Object to UndoManager.
Exercise care, as this is a potentially dangerous function
if the Object is not a UndoManager.
*/
func CastToUndoManager(object *gobject.Object) *UndoManager {
	return UndoManagerNewFromNative(object.Native())
}

// Equals compares this UndoManager with another UndoManager, and returns true if they represent the same Object.
func (recv *UndoManager) Equals(other *UndoManager) bool {
	return other.Native() == recv.Native()
}

func (recv *UndoManager) Native() unsafe.Pointer {
	return recv.native
}

var undoManagerBeginNotUndoableActionFunction *gi.Function
var undoManagerBeginNotUndoableActionFunction_Once sync.Once

func undoManagerBeginNotUndoableActionFunction_Set() error {
	var err error
	undoManagerBeginNotUndoableActionFunction_Once.Do(func() {
		err = undoManagerInterface_Set()
		if err != nil {
			return
		}
		undoManagerBeginNotUndoableActionFunction, err = undoManagerInterface.InvokerNew("begin_not_undoable_action")
	})
	return err
}

// BeginNotUndoableAction is a representation of the C type gtk_source_undo_manager_begin_not_undoable_action.
func (recv *UndoManager) BeginNotUndoableAction() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := undoManagerBeginNotUndoableActionFunction_Set()
	if err == nil {
		undoManagerBeginNotUndoableActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var undoManagerCanRedoFunction *gi.Function
var undoManagerCanRedoFunction_Once sync.Once

func undoManagerCanRedoFunction_Set() error {
	var err error
	undoManagerCanRedoFunction_Once.Do(func() {
		err = undoManagerInterface_Set()
		if err != nil {
			return
		}
		undoManagerCanRedoFunction, err = undoManagerInterface.InvokerNew("can_redo")
	})
	return err
}

// CanRedo is a representation of the C type gtk_source_undo_manager_can_redo.
func (recv *UndoManager) CanRedo() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := undoManagerCanRedoFunction_Set()
	if err == nil {
		ret = undoManagerCanRedoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var undoManagerCanRedoChangedFunction *gi.Function
var undoManagerCanRedoChangedFunction_Once sync.Once

func undoManagerCanRedoChangedFunction_Set() error {
	var err error
	undoManagerCanRedoChangedFunction_Once.Do(func() {
		err = undoManagerInterface_Set()
		if err != nil {
			return
		}
		undoManagerCanRedoChangedFunction, err = undoManagerInterface.InvokerNew("can_redo_changed")
	})
	return err
}

// CanRedoChanged is a representation of the C type gtk_source_undo_manager_can_redo_changed.
func (recv *UndoManager) CanRedoChanged() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := undoManagerCanRedoChangedFunction_Set()
	if err == nil {
		undoManagerCanRedoChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var undoManagerCanUndoFunction *gi.Function
var undoManagerCanUndoFunction_Once sync.Once

func undoManagerCanUndoFunction_Set() error {
	var err error
	undoManagerCanUndoFunction_Once.Do(func() {
		err = undoManagerInterface_Set()
		if err != nil {
			return
		}
		undoManagerCanUndoFunction, err = undoManagerInterface.InvokerNew("can_undo")
	})
	return err
}

// CanUndo is a representation of the C type gtk_source_undo_manager_can_undo.
func (recv *UndoManager) CanUndo() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := undoManagerCanUndoFunction_Set()
	if err == nil {
		ret = undoManagerCanUndoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var undoManagerCanUndoChangedFunction *gi.Function
var undoManagerCanUndoChangedFunction_Once sync.Once

func undoManagerCanUndoChangedFunction_Set() error {
	var err error
	undoManagerCanUndoChangedFunction_Once.Do(func() {
		err = undoManagerInterface_Set()
		if err != nil {
			return
		}
		undoManagerCanUndoChangedFunction, err = undoManagerInterface.InvokerNew("can_undo_changed")
	})
	return err
}

// CanUndoChanged is a representation of the C type gtk_source_undo_manager_can_undo_changed.
func (recv *UndoManager) CanUndoChanged() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := undoManagerCanUndoChangedFunction_Set()
	if err == nil {
		undoManagerCanUndoChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var undoManagerEndNotUndoableActionFunction *gi.Function
var undoManagerEndNotUndoableActionFunction_Once sync.Once

func undoManagerEndNotUndoableActionFunction_Set() error {
	var err error
	undoManagerEndNotUndoableActionFunction_Once.Do(func() {
		err = undoManagerInterface_Set()
		if err != nil {
			return
		}
		undoManagerEndNotUndoableActionFunction, err = undoManagerInterface.InvokerNew("end_not_undoable_action")
	})
	return err
}

// EndNotUndoableAction is a representation of the C type gtk_source_undo_manager_end_not_undoable_action.
func (recv *UndoManager) EndNotUndoableAction() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := undoManagerEndNotUndoableActionFunction_Set()
	if err == nil {
		undoManagerEndNotUndoableActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var undoManagerRedoFunction *gi.Function
var undoManagerRedoFunction_Once sync.Once

func undoManagerRedoFunction_Set() error {
	var err error
	undoManagerRedoFunction_Once.Do(func() {
		err = undoManagerInterface_Set()
		if err != nil {
			return
		}
		undoManagerRedoFunction, err = undoManagerInterface.InvokerNew("redo")
	})
	return err
}

// Redo is a representation of the C type gtk_source_undo_manager_redo.
func (recv *UndoManager) Redo() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := undoManagerRedoFunction_Set()
	if err == nil {
		undoManagerRedoFunction.Invoke(inArgs[:], nil)
	}

	return
}

var undoManagerUndoFunction *gi.Function
var undoManagerUndoFunction_Once sync.Once

func undoManagerUndoFunction_Set() error {
	var err error
	undoManagerUndoFunction_Once.Do(func() {
		err = undoManagerInterface_Set()
		if err != nil {
			return
		}
		undoManagerUndoFunction, err = undoManagerInterface.InvokerNew("undo")
	})
	return err
}

// Undo is a representation of the C type gtk_source_undo_manager_undo.
func (recv *UndoManager) Undo() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := undoManagerUndoFunction_Set()
	if err == nil {
		undoManagerUndoFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *UndoManager) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}
