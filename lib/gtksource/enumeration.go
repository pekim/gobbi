// Code generated - DO NOT EDIT.

package gtksource

// Backgroundpatterntype is a representation of the C type BackgroundPatternType.
//
// since 3.16
type Backgroundpatterntype int

const (
	// none
	GtkSourceBackgroundPatternTypeNone Backgroundpatterntype = 0
	// grid
	GtkSourceBackgroundPatternTypeGrid Backgroundpatterntype = 1
)

// Bracketmatchtype is a representation of the C type BracketMatchType.
type Bracketmatchtype int

const (
	// none
	GtkSourceBracketMatchNone Bracketmatchtype = 0
	// out_of_range
	GtkSourceBracketMatchOutOfRange Bracketmatchtype = 1
	// not_found
	GtkSourceBracketMatchNotFound Bracketmatchtype = 2
	// found
	GtkSourceBracketMatchFound Bracketmatchtype = 3
)

// Changecasetype is a representation of the C type ChangeCaseType.
//
// since 3.12
type Changecasetype int

const (
	// lower
	GtkSourceChangeCaseLower Changecasetype = 0
	// upper
	GtkSourceChangeCaseUpper Changecasetype = 1
	// toggle
	GtkSourceChangeCaseToggle Changecasetype = 2
	// title
	GtkSourceChangeCaseTitle Changecasetype = 3
)

// Completionerror is a representation of the C type CompletionError.
type Completionerror int

const (
	// already_bound
	GtkSourceCompletionErrorAlreadyBound Completionerror = 0
	// not_bound
	GtkSourceCompletionErrorNotBound Completionerror = 1
)

// Compressiontype is a representation of the C type CompressionType.
//
// since 3.14
type Compressiontype int

const (
	// none
	GtkSourceCompressionTypeNone Compressiontype = 0
	// gzip
	GtkSourceCompressionTypeGzip Compressiontype = 1
)

// Fileloadererror is a representation of the C type FileLoaderError.
type Fileloadererror int

const (
	// too_big
	GtkSourceFileLoaderErrorTooBig Fileloadererror = 0
	// encoding_auto_detection_failed
	GtkSourceFileLoaderErrorEncodingAutoDetectionFailed Fileloadererror = 1
	// conversion_fallback
	GtkSourceFileLoaderErrorConversionFallback Fileloadererror = 2
)

// Filesavererror is a representation of the C type FileSaverError.
//
// since 3.14
type Filesavererror int

const (
	// invalid_chars
	GtkSourceFileSaverErrorInvalidChars Filesavererror = 0
	// externally_modified
	GtkSourceFileSaverErrorExternallyModified Filesavererror = 1
)

// Gutterrendereralignmentmode is a representation of the C type GutterRendererAlignmentMode.
type Gutterrendereralignmentmode int

const (
	// cell
	GtkSourceGutterRendererAlignmentModeCell Gutterrendereralignmentmode = 0
	// first
	GtkSourceGutterRendererAlignmentModeFirst Gutterrendereralignmentmode = 1
	// last
	GtkSourceGutterRendererAlignmentModeLast Gutterrendereralignmentmode = 2
)

// Newlinetype is a representation of the C type NewlineType.
//
// since 3.14
type Newlinetype int

const (
	// lf
	GtkSourceNewlineTypeLf Newlinetype = 0
	// cr
	GtkSourceNewlineTypeCr Newlinetype = 1
	// cr_lf
	GtkSourceNewlineTypeCrLf Newlinetype = 2
)

// Smarthomeendtype is a representation of the C type SmartHomeEndType.
type Smarthomeendtype int

const (
	// disabled
	GtkSourceSmartHomeEndDisabled Smarthomeendtype = 0
	// before
	GtkSourceSmartHomeEndBefore Smarthomeendtype = 1
	// after
	GtkSourceSmartHomeEndAfter Smarthomeendtype = 2
	// always
	GtkSourceSmartHomeEndAlways Smarthomeendtype = 3
)

// Viewgutterposition is a representation of the C type ViewGutterPosition.
type Viewgutterposition int

const (
	// lines
	GtkSourceViewGutterPositionLines Viewgutterposition = -30
	// marks
	GtkSourceViewGutterPositionMarks Viewgutterposition = -20
)
