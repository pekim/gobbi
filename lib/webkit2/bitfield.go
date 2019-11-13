// Code generated - DO NOT EDIT.

package webkit2

// Editortypingattributes is a representation of the C type EditorTypingAttributes.
//
// since 2.10
type Editortypingattributes int

const (
	Editortypingattributes_None          Editortypingattributes = 2
	Editortypingattributes_Bold          Editortypingattributes = 4
	Editortypingattributes_Italic        Editortypingattributes = 8
	Editortypingattributes_Underline     Editortypingattributes = 16
	Editortypingattributes_Strikethrough Editortypingattributes = 32
)

// Findoptions is a representation of the C type FindOptions.
type Findoptions int

const (
	Findoptions_None                          Findoptions = 0
	Findoptions_CaseInsensitive               Findoptions = 1
	Findoptions_AtWordStarts                  Findoptions = 2
	Findoptions_TreatMedialCapitalAsWordStart Findoptions = 4
	Findoptions_Backwards                     Findoptions = 8
	Findoptions_WrapAround                    Findoptions = 16
)

// Hittestresultcontext is a representation of the C type HitTestResultContext.
type Hittestresultcontext int

const (
	Hittestresultcontext_Document  Hittestresultcontext = 2
	Hittestresultcontext_Link      Hittestresultcontext = 4
	Hittestresultcontext_Image     Hittestresultcontext = 8
	Hittestresultcontext_Media     Hittestresultcontext = 16
	Hittestresultcontext_Editable  Hittestresultcontext = 32
	Hittestresultcontext_Scrollbar Hittestresultcontext = 64
	Hittestresultcontext_Selection Hittestresultcontext = 128
)

// Snapshotoptions is a representation of the C type SnapshotOptions.
type Snapshotoptions int

const (
	Snapshotoptions_None                         Snapshotoptions = 0
	Snapshotoptions_IncludeSelectionHighlighting Snapshotoptions = 1
	Snapshotoptions_TransparentBackground        Snapshotoptions = 2
)

// Websitedatatypes is a representation of the C type WebsiteDataTypes.
//
// since 2.16
type Websitedatatypes int

const (
	Websitedatatypes_MemoryCache             Websitedatatypes = 1
	Websitedatatypes_DiskCache               Websitedatatypes = 2
	Websitedatatypes_OfflineApplicationCache Websitedatatypes = 4
	Websitedatatypes_SessionStorage          Websitedatatypes = 8
	Websitedatatypes_LocalStorage            Websitedatatypes = 16
	Websitedatatypes_WebsqlDatabases         Websitedatatypes = 32
	Websitedatatypes_IndexeddbDatabases      Websitedatatypes = 64
	Websitedatatypes_PluginData              Websitedatatypes = 128
	Websitedatatypes_Cookies                 Websitedatatypes = 256
	Websitedatatypes_DeviceIdHashSalt        Websitedatatypes = 512
	Websitedatatypes_HstsCache               Websitedatatypes = 1024
	Websitedatatypes_All                     Websitedatatypes = 2047
)
