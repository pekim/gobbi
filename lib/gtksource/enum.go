// Code generated - DO NOT EDIT.

package gtksource

type BackgroundPatternType uint32

const (
	BackgroundPatternType_None BackgroundPatternType = int64(0)
	BackgroundPatternType_Grid BackgroundPatternType = int64(1)
)

type BracketMatchType uint32

const (
	BracketMatchType_None       BracketMatchType = int64(0)
	BracketMatchType_OutOfRange BracketMatchType = int64(1)
	BracketMatchType_NotFound   BracketMatchType = int64(2)
	BracketMatchType_Found      BracketMatchType = int64(3)
)

type ChangeCaseType uint32

const (
	ChangeCaseType_Lower  ChangeCaseType = int64(0)
	ChangeCaseType_Upper  ChangeCaseType = int64(1)
	ChangeCaseType_Toggle ChangeCaseType = int64(2)
	ChangeCaseType_Title  ChangeCaseType = int64(3)
)

type CompletionError uint32

const (
	CompletionError_AlreadyBound CompletionError = int64(0)
	CompletionError_NotBound     CompletionError = int64(1)
)

type CompressionType uint32

const (
	CompressionType_None CompressionType = int64(0)
	CompressionType_Gzip CompressionType = int64(1)
)

type FileLoaderError uint32

const (
	FileLoaderError_TooBig                      FileLoaderError = int64(0)
	FileLoaderError_EncodingAutoDetectionFailed FileLoaderError = int64(1)
	FileLoaderError_ConversionFallback          FileLoaderError = int64(2)
)

type FileSaverError uint32

const (
	FileSaverError_InvalidChars       FileSaverError = int64(0)
	FileSaverError_ExternallyModified FileSaverError = int64(1)
)

type GutterRendererAlignmentMode uint32

const (
	GutterRendererAlignmentMode_Cell  GutterRendererAlignmentMode = int64(0)
	GutterRendererAlignmentMode_First GutterRendererAlignmentMode = int64(1)
	GutterRendererAlignmentMode_Last  GutterRendererAlignmentMode = int64(2)
)

type NewlineType uint32

const (
	NewlineType_Lf   NewlineType = int64(0)
	NewlineType_Cr   NewlineType = int64(1)
	NewlineType_CrLf NewlineType = int64(2)
)

type SmartHomeEndType uint32

const (
	SmartHomeEndType_Disabled SmartHomeEndType = int64(0)
	SmartHomeEndType_Before   SmartHomeEndType = int64(1)
	SmartHomeEndType_After    SmartHomeEndType = int64(2)
	SmartHomeEndType_Always   SmartHomeEndType = int64(3)
)

type ViewGutterPosition int32

const (
	ViewGutterPosition_Lines ViewGutterPosition = int64(-30)
	ViewGutterPosition_Marks ViewGutterPosition = int64(-20)
)
