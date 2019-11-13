// Code generated - DO NOT EDIT.

package webkit2

// Editortypingattributes is a representation of the C type WebKitEditorTypingAttributes.
//
// since 2.10
type Editortypingattributes int

const (
	// Editortypingattributes_None is a representation of the C type WEBKIT_EDITOR_TYPING_ATTRIBUTE_NONE.
	Editortypingattributes_None Editortypingattributes = 2
	// Editortypingattributes_Bold is a representation of the C type WEBKIT_EDITOR_TYPING_ATTRIBUTE_BOLD.
	Editortypingattributes_Bold Editortypingattributes = 4
	// Editortypingattributes_Italic is a representation of the C type WEBKIT_EDITOR_TYPING_ATTRIBUTE_ITALIC.
	Editortypingattributes_Italic Editortypingattributes = 8
	// Editortypingattributes_Underline is a representation of the C type WEBKIT_EDITOR_TYPING_ATTRIBUTE_UNDERLINE.
	Editortypingattributes_Underline Editortypingattributes = 16
	// Editortypingattributes_Strikethrough is a representation of the C type WEBKIT_EDITOR_TYPING_ATTRIBUTE_STRIKETHROUGH.
	Editortypingattributes_Strikethrough Editortypingattributes = 32
)

// Findoptions is a representation of the C type WebKitFindOptions.
type Findoptions int

const (
	// Findoptions_None is a representation of the C type WEBKIT_FIND_OPTIONS_NONE.
	Findoptions_None Findoptions = 0
	// Findoptions_CaseInsensitive is a representation of the C type WEBKIT_FIND_OPTIONS_CASE_INSENSITIVE.
	Findoptions_CaseInsensitive Findoptions = 1
	// Findoptions_AtWordStarts is a representation of the C type WEBKIT_FIND_OPTIONS_AT_WORD_STARTS.
	Findoptions_AtWordStarts Findoptions = 2
	// Findoptions_TreatMedialCapitalAsWordStart is a representation of the C type WEBKIT_FIND_OPTIONS_TREAT_MEDIAL_CAPITAL_AS_WORD_START.
	Findoptions_TreatMedialCapitalAsWordStart Findoptions = 4
	// Findoptions_Backwards is a representation of the C type WEBKIT_FIND_OPTIONS_BACKWARDS.
	Findoptions_Backwards Findoptions = 8
	// Findoptions_WrapAround is a representation of the C type WEBKIT_FIND_OPTIONS_WRAP_AROUND.
	Findoptions_WrapAround Findoptions = 16
)

// Hittestresultcontext is a representation of the C type WebKitHitTestResultContext.
type Hittestresultcontext int

const (
	// Hittestresultcontext_Document is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_DOCUMENT.
	Hittestresultcontext_Document Hittestresultcontext = 2
	// Hittestresultcontext_Link is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_LINK.
	Hittestresultcontext_Link Hittestresultcontext = 4
	// Hittestresultcontext_Image is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_IMAGE.
	Hittestresultcontext_Image Hittestresultcontext = 8
	// Hittestresultcontext_Media is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_MEDIA.
	Hittestresultcontext_Media Hittestresultcontext = 16
	// Hittestresultcontext_Editable is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_EDITABLE.
	Hittestresultcontext_Editable Hittestresultcontext = 32
	// Hittestresultcontext_Scrollbar is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_SCROLLBAR.
	Hittestresultcontext_Scrollbar Hittestresultcontext = 64
	// Hittestresultcontext_Selection is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_SELECTION.
	Hittestresultcontext_Selection Hittestresultcontext = 128
)

// Snapshotoptions is a representation of the C type WebKitSnapshotOptions.
type Snapshotoptions int

const (
	// Snapshotoptions_None is a representation of the C type WEBKIT_SNAPSHOT_OPTIONS_NONE.
	Snapshotoptions_None Snapshotoptions = 0
	// Snapshotoptions_IncludeSelectionHighlighting is a representation of the C type WEBKIT_SNAPSHOT_OPTIONS_INCLUDE_SELECTION_HIGHLIGHTING.
	Snapshotoptions_IncludeSelectionHighlighting Snapshotoptions = 1
	// Snapshotoptions_TransparentBackground is a representation of the C type WEBKIT_SNAPSHOT_OPTIONS_TRANSPARENT_BACKGROUND.
	Snapshotoptions_TransparentBackground Snapshotoptions = 2
)

// Websitedatatypes is a representation of the C type WebKitWebsiteDataTypes.
//
// since 2.16
type Websitedatatypes int

const (
	// Websitedatatypes_MemoryCache is a representation of the C type WEBKIT_WEBSITE_DATA_MEMORY_CACHE.
	Websitedatatypes_MemoryCache Websitedatatypes = 1
	// Websitedatatypes_DiskCache is a representation of the C type WEBKIT_WEBSITE_DATA_DISK_CACHE.
	Websitedatatypes_DiskCache Websitedatatypes = 2
	// Websitedatatypes_OfflineApplicationCache is a representation of the C type WEBKIT_WEBSITE_DATA_OFFLINE_APPLICATION_CACHE.
	Websitedatatypes_OfflineApplicationCache Websitedatatypes = 4
	// Websitedatatypes_SessionStorage is a representation of the C type WEBKIT_WEBSITE_DATA_SESSION_STORAGE.
	Websitedatatypes_SessionStorage Websitedatatypes = 8
	// Websitedatatypes_LocalStorage is a representation of the C type WEBKIT_WEBSITE_DATA_LOCAL_STORAGE.
	Websitedatatypes_LocalStorage Websitedatatypes = 16
	// Websitedatatypes_WebsqlDatabases is a representation of the C type WEBKIT_WEBSITE_DATA_WEBSQL_DATABASES.
	Websitedatatypes_WebsqlDatabases Websitedatatypes = 32
	// Websitedatatypes_IndexeddbDatabases is a representation of the C type WEBKIT_WEBSITE_DATA_INDEXEDDB_DATABASES.
	Websitedatatypes_IndexeddbDatabases Websitedatatypes = 64
	// Websitedatatypes_PluginData is a representation of the C type WEBKIT_WEBSITE_DATA_PLUGIN_DATA.
	Websitedatatypes_PluginData Websitedatatypes = 128
	// Websitedatatypes_Cookies is a representation of the C type WEBKIT_WEBSITE_DATA_COOKIES.
	Websitedatatypes_Cookies Websitedatatypes = 256
	// Websitedatatypes_DeviceIdHashSalt is a representation of the C type WEBKIT_WEBSITE_DATA_DEVICE_ID_HASH_SALT.
	Websitedatatypes_DeviceIdHashSalt Websitedatatypes = 512
	// Websitedatatypes_HstsCache is a representation of the C type WEBKIT_WEBSITE_DATA_HSTS_CACHE.
	Websitedatatypes_HstsCache Websitedatatypes = 1024
	// Websitedatatypes_All is a representation of the C type WEBKIT_WEBSITE_DATA_ALL.
	Websitedatatypes_All Websitedatatypes = 2047
)
