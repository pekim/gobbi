// Code generated - DO NOT EDIT.

package gtksource

// BackgroundPatternType is a representation of the C type GtkSourceBackgroundPatternType.
//
// since 3.16
type BackgroundPatternType int

const (
	// BackgroundPatternType_None is a representation of the C type GTK_SOURCE_BACKGROUND_PATTERN_TYPE_NONE.
	BackgroundPatternType_None BackgroundPatternType = 0
	// BackgroundPatternType_Grid is a representation of the C type GTK_SOURCE_BACKGROUND_PATTERN_TYPE_GRID.
	BackgroundPatternType_Grid BackgroundPatternType = 1
)

// BracketMatchType is a representation of the C type GtkSourceBracketMatchType.
type BracketMatchType int

const (
	// BracketMatchType_None is a representation of the C type GTK_SOURCE_BRACKET_MATCH_NONE.
	BracketMatchType_None BracketMatchType = 0
	// BracketMatchType_OutOfRange is a representation of the C type GTK_SOURCE_BRACKET_MATCH_OUT_OF_RANGE.
	BracketMatchType_OutOfRange BracketMatchType = 1
	// BracketMatchType_NotFound is a representation of the C type GTK_SOURCE_BRACKET_MATCH_NOT_FOUND.
	BracketMatchType_NotFound BracketMatchType = 2
	// BracketMatchType_Found is a representation of the C type GTK_SOURCE_BRACKET_MATCH_FOUND.
	BracketMatchType_Found BracketMatchType = 3
)

// ChangeCaseType is a representation of the C type GtkSourceChangeCaseType.
//
// since 3.12
type ChangeCaseType int

const (
	// ChangeCaseType_Lower is a representation of the C type GTK_SOURCE_CHANGE_CASE_LOWER.
	ChangeCaseType_Lower ChangeCaseType = 0
	// ChangeCaseType_Upper is a representation of the C type GTK_SOURCE_CHANGE_CASE_UPPER.
	ChangeCaseType_Upper ChangeCaseType = 1
	// ChangeCaseType_Toggle is a representation of the C type GTK_SOURCE_CHANGE_CASE_TOGGLE.
	ChangeCaseType_Toggle ChangeCaseType = 2
	// ChangeCaseType_Title is a representation of the C type GTK_SOURCE_CHANGE_CASE_TITLE.
	ChangeCaseType_Title ChangeCaseType = 3
)

// CompletionError is a representation of the C type GtkSourceCompletionError.
type CompletionError int

const (
	// CompletionError_AlreadyBound is a representation of the C type GTK_SOURCE_COMPLETION_ERROR_ALREADY_BOUND.
	CompletionError_AlreadyBound CompletionError = 0
	// CompletionError_NotBound is a representation of the C type GTK_SOURCE_COMPLETION_ERROR_NOT_BOUND.
	CompletionError_NotBound CompletionError = 1
)

// CompressionType is a representation of the C type GtkSourceCompressionType.
//
// since 3.14
type CompressionType int

const (
	// CompressionType_None is a representation of the C type GTK_SOURCE_COMPRESSION_TYPE_NONE.
	CompressionType_None CompressionType = 0
	// CompressionType_Gzip is a representation of the C type GTK_SOURCE_COMPRESSION_TYPE_GZIP.
	CompressionType_Gzip CompressionType = 1
)

// FileLoaderError is a representation of the C type GtkSourceFileLoaderError.
type FileLoaderError int

const (
	// FileLoaderError_TooBig is a representation of the C type GTK_SOURCE_FILE_LOADER_ERROR_TOO_BIG.
	FileLoaderError_TooBig FileLoaderError = 0
	// FileLoaderError_EncodingAutoDetectionFailed is a representation of the C type GTK_SOURCE_FILE_LOADER_ERROR_ENCODING_AUTO_DETECTION_FAILED.
	FileLoaderError_EncodingAutoDetectionFailed FileLoaderError = 1
	// FileLoaderError_ConversionFallback is a representation of the C type GTK_SOURCE_FILE_LOADER_ERROR_CONVERSION_FALLBACK.
	FileLoaderError_ConversionFallback FileLoaderError = 2
)

// FileSaverError is a representation of the C type GtkSourceFileSaverError.
//
// since 3.14
type FileSaverError int

const (
	// FileSaverError_InvalidChars is a representation of the C type GTK_SOURCE_FILE_SAVER_ERROR_INVALID_CHARS.
	FileSaverError_InvalidChars FileSaverError = 0
	// FileSaverError_ExternallyModified is a representation of the C type GTK_SOURCE_FILE_SAVER_ERROR_EXTERNALLY_MODIFIED.
	FileSaverError_ExternallyModified FileSaverError = 1
)

// GutterRendererAlignmentMode is a representation of the C type GtkSourceGutterRendererAlignmentMode.
type GutterRendererAlignmentMode int

const (
	// GutterRendererAlignmentMode_Cell is a representation of the C type GTK_SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_CELL.
	GutterRendererAlignmentMode_Cell GutterRendererAlignmentMode = 0
	// GutterRendererAlignmentMode_First is a representation of the C type GTK_SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_FIRST.
	GutterRendererAlignmentMode_First GutterRendererAlignmentMode = 1
	// GutterRendererAlignmentMode_Last is a representation of the C type GTK_SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_LAST.
	GutterRendererAlignmentMode_Last GutterRendererAlignmentMode = 2
)

// NewlineType is a representation of the C type GtkSourceNewlineType.
//
// since 3.14
type NewlineType int

const (
	// NewlineType_Lf is a representation of the C type GTK_SOURCE_NEWLINE_TYPE_LF.
	NewlineType_Lf NewlineType = 0
	// NewlineType_Cr is a representation of the C type GTK_SOURCE_NEWLINE_TYPE_CR.
	NewlineType_Cr NewlineType = 1
	// NewlineType_CrLf is a representation of the C type GTK_SOURCE_NEWLINE_TYPE_CR_LF.
	NewlineType_CrLf NewlineType = 2
)

// SmartHomeEndType is a representation of the C type GtkSourceSmartHomeEndType.
type SmartHomeEndType int

const (
	// SmartHomeEndType_Disabled is a representation of the C type GTK_SOURCE_SMART_HOME_END_DISABLED.
	SmartHomeEndType_Disabled SmartHomeEndType = 0
	// SmartHomeEndType_Before is a representation of the C type GTK_SOURCE_SMART_HOME_END_BEFORE.
	SmartHomeEndType_Before SmartHomeEndType = 1
	// SmartHomeEndType_After is a representation of the C type GTK_SOURCE_SMART_HOME_END_AFTER.
	SmartHomeEndType_After SmartHomeEndType = 2
	// SmartHomeEndType_Always is a representation of the C type GTK_SOURCE_SMART_HOME_END_ALWAYS.
	SmartHomeEndType_Always SmartHomeEndType = 3
)

// ViewGutterPosition is a representation of the C type GtkSourceViewGutterPosition.
type ViewGutterPosition int

const (
	// ViewGutterPosition_Lines is a representation of the C type GTK_SOURCE_VIEW_GUTTER_POSITION_LINES.
	ViewGutterPosition_Lines ViewGutterPosition = -30
	// ViewGutterPosition_Marks is a representation of the C type GTK_SOURCE_VIEW_GUTTER_POSITION_MARKS.
	ViewGutterPosition_Marks ViewGutterPosition = -20
)
