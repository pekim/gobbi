// Code generated - DO NOT EDIT.

package webkit2

// EditorTypingAttributes is a representation of the C type WebKitEditorTypingAttributes.
//
// since 2.10
type EditorTypingAttributes int

const (
	// EditorTypingAttributes_None is a representation of the C type WEBKIT_EDITOR_TYPING_ATTRIBUTE_NONE.
	EditorTypingAttributes_None EditorTypingAttributes = 2
	// EditorTypingAttributes_Bold is a representation of the C type WEBKIT_EDITOR_TYPING_ATTRIBUTE_BOLD.
	EditorTypingAttributes_Bold EditorTypingAttributes = 4
	// EditorTypingAttributes_Italic is a representation of the C type WEBKIT_EDITOR_TYPING_ATTRIBUTE_ITALIC.
	EditorTypingAttributes_Italic EditorTypingAttributes = 8
	// EditorTypingAttributes_Underline is a representation of the C type WEBKIT_EDITOR_TYPING_ATTRIBUTE_UNDERLINE.
	EditorTypingAttributes_Underline EditorTypingAttributes = 16
	// EditorTypingAttributes_Strikethrough is a representation of the C type WEBKIT_EDITOR_TYPING_ATTRIBUTE_STRIKETHROUGH.
	EditorTypingAttributes_Strikethrough EditorTypingAttributes = 32
)

// FindOptions is a representation of the C type WebKitFindOptions.
type FindOptions int

const (
	// FindOptions_None is a representation of the C type WEBKIT_FIND_OPTIONS_NONE.
	FindOptions_None FindOptions = 0
	// FindOptions_CaseInsensitive is a representation of the C type WEBKIT_FIND_OPTIONS_CASE_INSENSITIVE.
	FindOptions_CaseInsensitive FindOptions = 1
	// FindOptions_AtWordStarts is a representation of the C type WEBKIT_FIND_OPTIONS_AT_WORD_STARTS.
	FindOptions_AtWordStarts FindOptions = 2
	// FindOptions_TreatMedialCapitalAsWordStart is a representation of the C type WEBKIT_FIND_OPTIONS_TREAT_MEDIAL_CAPITAL_AS_WORD_START.
	FindOptions_TreatMedialCapitalAsWordStart FindOptions = 4
	// FindOptions_Backwards is a representation of the C type WEBKIT_FIND_OPTIONS_BACKWARDS.
	FindOptions_Backwards FindOptions = 8
	// FindOptions_WrapAround is a representation of the C type WEBKIT_FIND_OPTIONS_WRAP_AROUND.
	FindOptions_WrapAround FindOptions = 16
)

// HitTestResultContext is a representation of the C type WebKitHitTestResultContext.
type HitTestResultContext int

const (
	// HitTestResultContext_Document is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_DOCUMENT.
	HitTestResultContext_Document HitTestResultContext = 2
	// HitTestResultContext_Link is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_LINK.
	HitTestResultContext_Link HitTestResultContext = 4
	// HitTestResultContext_Image is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_IMAGE.
	HitTestResultContext_Image HitTestResultContext = 8
	// HitTestResultContext_Media is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_MEDIA.
	HitTestResultContext_Media HitTestResultContext = 16
	// HitTestResultContext_Editable is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_EDITABLE.
	HitTestResultContext_Editable HitTestResultContext = 32
	// HitTestResultContext_Scrollbar is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_SCROLLBAR.
	HitTestResultContext_Scrollbar HitTestResultContext = 64
	// HitTestResultContext_Selection is a representation of the C type WEBKIT_HIT_TEST_RESULT_CONTEXT_SELECTION.
	HitTestResultContext_Selection HitTestResultContext = 128
)

// SnapshotOptions is a representation of the C type WebKitSnapshotOptions.
type SnapshotOptions int

const (
	// SnapshotOptions_None is a representation of the C type WEBKIT_SNAPSHOT_OPTIONS_NONE.
	SnapshotOptions_None SnapshotOptions = 0
	// SnapshotOptions_IncludeSelectionHighlighting is a representation of the C type WEBKIT_SNAPSHOT_OPTIONS_INCLUDE_SELECTION_HIGHLIGHTING.
	SnapshotOptions_IncludeSelectionHighlighting SnapshotOptions = 1
	// SnapshotOptions_TransparentBackground is a representation of the C type WEBKIT_SNAPSHOT_OPTIONS_TRANSPARENT_BACKGROUND.
	SnapshotOptions_TransparentBackground SnapshotOptions = 2
)

// WebsiteDataTypes is a representation of the C type WebKitWebsiteDataTypes.
//
// since 2.16
type WebsiteDataTypes int

const (
	// WebsiteDataTypes_MemoryCache is a representation of the C type WEBKIT_WEBSITE_DATA_MEMORY_CACHE.
	WebsiteDataTypes_MemoryCache WebsiteDataTypes = 1
	// WebsiteDataTypes_DiskCache is a representation of the C type WEBKIT_WEBSITE_DATA_DISK_CACHE.
	WebsiteDataTypes_DiskCache WebsiteDataTypes = 2
	// WebsiteDataTypes_OfflineApplicationCache is a representation of the C type WEBKIT_WEBSITE_DATA_OFFLINE_APPLICATION_CACHE.
	WebsiteDataTypes_OfflineApplicationCache WebsiteDataTypes = 4
	// WebsiteDataTypes_SessionStorage is a representation of the C type WEBKIT_WEBSITE_DATA_SESSION_STORAGE.
	WebsiteDataTypes_SessionStorage WebsiteDataTypes = 8
	// WebsiteDataTypes_LocalStorage is a representation of the C type WEBKIT_WEBSITE_DATA_LOCAL_STORAGE.
	WebsiteDataTypes_LocalStorage WebsiteDataTypes = 16
	// WebsiteDataTypes_WebsqlDatabases is a representation of the C type WEBKIT_WEBSITE_DATA_WEBSQL_DATABASES.
	WebsiteDataTypes_WebsqlDatabases WebsiteDataTypes = 32
	// WebsiteDataTypes_IndexeddbDatabases is a representation of the C type WEBKIT_WEBSITE_DATA_INDEXEDDB_DATABASES.
	WebsiteDataTypes_IndexeddbDatabases WebsiteDataTypes = 64
	// WebsiteDataTypes_PluginData is a representation of the C type WEBKIT_WEBSITE_DATA_PLUGIN_DATA.
	WebsiteDataTypes_PluginData WebsiteDataTypes = 128
	// WebsiteDataTypes_Cookies is a representation of the C type WEBKIT_WEBSITE_DATA_COOKIES.
	WebsiteDataTypes_Cookies WebsiteDataTypes = 256
	// WebsiteDataTypes_DeviceIdHashSalt is a representation of the C type WEBKIT_WEBSITE_DATA_DEVICE_ID_HASH_SALT.
	WebsiteDataTypes_DeviceIdHashSalt WebsiteDataTypes = 512
	// WebsiteDataTypes_HstsCache is a representation of the C type WEBKIT_WEBSITE_DATA_HSTS_CACHE.
	WebsiteDataTypes_HstsCache WebsiteDataTypes = 1024
	// WebsiteDataTypes_All is a representation of the C type WEBKIT_WEBSITE_DATA_ALL.
	WebsiteDataTypes_All WebsiteDataTypes = 2047
)
