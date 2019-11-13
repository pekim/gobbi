// Code generated - DO NOT EDIT.

package gtksource

// Backgroundpatterntype is a representation of the C type BackgroundPatternType.
//
// since 3.16
type Backgroundpatterntype int

const (
	Backgroundpatterntype_None Backgroundpatterntype = 0
	Backgroundpatterntype_Grid Backgroundpatterntype = 1
)

// Bracketmatchtype is a representation of the C type BracketMatchType.
type Bracketmatchtype int

const (
	Bracketmatchtype_None       Bracketmatchtype = 0
	Bracketmatchtype_OutOfRange Bracketmatchtype = 1
	Bracketmatchtype_NotFound   Bracketmatchtype = 2
	Bracketmatchtype_Found      Bracketmatchtype = 3
)

// Changecasetype is a representation of the C type ChangeCaseType.
//
// since 3.12
type Changecasetype int

const (
	Changecasetype_Lower  Changecasetype = 0
	Changecasetype_Upper  Changecasetype = 1
	Changecasetype_Toggle Changecasetype = 2
	Changecasetype_Title  Changecasetype = 3
)

// Completionerror is a representation of the C type CompletionError.
type Completionerror int

const (
	Completionerror_AlreadyBound Completionerror = 0
	Completionerror_NotBound     Completionerror = 1
)

// Compressiontype is a representation of the C type CompressionType.
//
// since 3.14
type Compressiontype int

const (
	Compressiontype_None Compressiontype = 0
	Compressiontype_Gzip Compressiontype = 1
)

// Fileloadererror is a representation of the C type FileLoaderError.
type Fileloadererror int

const (
	Fileloadererror_TooBig                      Fileloadererror = 0
	Fileloadererror_EncodingAutoDetectionFailed Fileloadererror = 1
	Fileloadererror_ConversionFallback          Fileloadererror = 2
)

// Filesavererror is a representation of the C type FileSaverError.
//
// since 3.14
type Filesavererror int

const (
	Filesavererror_InvalidChars       Filesavererror = 0
	Filesavererror_ExternallyModified Filesavererror = 1
)

// Gutterrendereralignmentmode is a representation of the C type GutterRendererAlignmentMode.
type Gutterrendereralignmentmode int

const (
	Gutterrendereralignmentmode_Cell  Gutterrendereralignmentmode = 0
	Gutterrendereralignmentmode_First Gutterrendereralignmentmode = 1
	Gutterrendereralignmentmode_Last  Gutterrendereralignmentmode = 2
)

// Newlinetype is a representation of the C type NewlineType.
//
// since 3.14
type Newlinetype int

const (
	Newlinetype_Lf   Newlinetype = 0
	Newlinetype_Cr   Newlinetype = 1
	Newlinetype_CrLf Newlinetype = 2
)

// Smarthomeendtype is a representation of the C type SmartHomeEndType.
type Smarthomeendtype int

const (
	Smarthomeendtype_Disabled Smarthomeendtype = 0
	Smarthomeendtype_Before   Smarthomeendtype = 1
	Smarthomeendtype_After    Smarthomeendtype = 2
	Smarthomeendtype_Always   Smarthomeendtype = 3
)

// Viewgutterposition is a representation of the C type ViewGutterPosition.
type Viewgutterposition int

const (
	Viewgutterposition_Lines Viewgutterposition = -30
	Viewgutterposition_Marks Viewgutterposition = -20
)
