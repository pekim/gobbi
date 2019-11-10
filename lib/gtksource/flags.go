// Code generated - DO NOT EDIT.

package gtksource

type CompletionActivation uint32

const (
	CompletionActivation_None          CompletionActivation = int64(0)
	CompletionActivation_Interactive   CompletionActivation = int64(1)
	CompletionActivation_UserRequested CompletionActivation = int64(2)
)

type DrawSpacesFlags uint32

const (
	DrawSpacesFlags_Space    DrawSpacesFlags = int64(1)
	DrawSpacesFlags_Tab      DrawSpacesFlags = int64(2)
	DrawSpacesFlags_Newline  DrawSpacesFlags = int64(4)
	DrawSpacesFlags_Nbsp     DrawSpacesFlags = int64(8)
	DrawSpacesFlags_Leading  DrawSpacesFlags = int64(16)
	DrawSpacesFlags_Text     DrawSpacesFlags = int64(32)
	DrawSpacesFlags_Trailing DrawSpacesFlags = int64(64)
	DrawSpacesFlags_All      DrawSpacesFlags = int64(127)
)

type FileSaverFlags uint32

const (
	FileSaverFlags_None                   FileSaverFlags = int64(0)
	FileSaverFlags_IgnoreInvalidChars     FileSaverFlags = int64(1)
	FileSaverFlags_IgnoreModificationTime FileSaverFlags = int64(2)
	FileSaverFlags_CreateBackup           FileSaverFlags = int64(4)
)

type GutterRendererState uint32

const (
	GutterRendererState_Normal   GutterRendererState = int64(0)
	GutterRendererState_Cursor   GutterRendererState = int64(1)
	GutterRendererState_Prelit   GutterRendererState = int64(2)
	GutterRendererState_Selected GutterRendererState = int64(4)
)

type SortFlags uint32

const (
	SortFlags_None             SortFlags = int64(0)
	SortFlags_CaseSensitive    SortFlags = int64(1)
	SortFlags_ReverseOrder     SortFlags = int64(2)
	SortFlags_RemoveDuplicates SortFlags = int64(4)
)

type SpaceLocationFlags uint32

const (
	SpaceLocationFlags_None       SpaceLocationFlags = int64(0)
	SpaceLocationFlags_Leading    SpaceLocationFlags = int64(1)
	SpaceLocationFlags_InsideText SpaceLocationFlags = int64(2)
	SpaceLocationFlags_Trailing   SpaceLocationFlags = int64(4)
	SpaceLocationFlags_All        SpaceLocationFlags = int64(7)
)

type SpaceTypeFlags uint32

const (
	SpaceTypeFlags_None    SpaceTypeFlags = int64(0)
	SpaceTypeFlags_Space   SpaceTypeFlags = int64(1)
	SpaceTypeFlags_Tab     SpaceTypeFlags = int64(2)
	SpaceTypeFlags_Newline SpaceTypeFlags = int64(4)
	SpaceTypeFlags_Nbsp    SpaceTypeFlags = int64(8)
	SpaceTypeFlags_All     SpaceTypeFlags = int64(15)
)
