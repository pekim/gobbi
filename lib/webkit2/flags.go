// Code generated - DO NOT EDIT.

package webkit2

type EditorTypingAttributes uint32

const (
	EditorTypingAttributes_None          EditorTypingAttributes = int64(2)
	EditorTypingAttributes_Bold          EditorTypingAttributes = int64(4)
	EditorTypingAttributes_Italic        EditorTypingAttributes = int64(8)
	EditorTypingAttributes_Underline     EditorTypingAttributes = int64(16)
	EditorTypingAttributes_Strikethrough EditorTypingAttributes = int64(32)
)

type FindOptions uint32

const (
	FindOptions_None                          FindOptions = int64(0)
	FindOptions_CaseInsensitive               FindOptions = int64(1)
	FindOptions_AtWordStarts                  FindOptions = int64(2)
	FindOptions_TreatMedialCapitalAsWordStart FindOptions = int64(4)
	FindOptions_Backwards                     FindOptions = int64(8)
	FindOptions_WrapAround                    FindOptions = int64(16)
)

type HitTestResultContext uint32

const (
	HitTestResultContext_Document  HitTestResultContext = int64(2)
	HitTestResultContext_Link      HitTestResultContext = int64(4)
	HitTestResultContext_Image     HitTestResultContext = int64(8)
	HitTestResultContext_Media     HitTestResultContext = int64(16)
	HitTestResultContext_Editable  HitTestResultContext = int64(32)
	HitTestResultContext_Scrollbar HitTestResultContext = int64(64)
	HitTestResultContext_Selection HitTestResultContext = int64(128)
)

type SnapshotOptions uint32

const (
	SnapshotOptions_None                         SnapshotOptions = int64(0)
	SnapshotOptions_IncludeSelectionHighlighting SnapshotOptions = int64(1)
	SnapshotOptions_TransparentBackground        SnapshotOptions = int64(2)
)

type WebsiteDataTypes uint32

const (
	WebsiteDataTypes_MemoryCache             WebsiteDataTypes = int64(1)
	WebsiteDataTypes_DiskCache               WebsiteDataTypes = int64(2)
	WebsiteDataTypes_OfflineApplicationCache WebsiteDataTypes = int64(4)
	WebsiteDataTypes_SessionStorage          WebsiteDataTypes = int64(8)
	WebsiteDataTypes_LocalStorage            WebsiteDataTypes = int64(16)
	WebsiteDataTypes_WebsqlDatabases         WebsiteDataTypes = int64(32)
	WebsiteDataTypes_IndexeddbDatabases      WebsiteDataTypes = int64(64)
	WebsiteDataTypes_PluginData              WebsiteDataTypes = int64(128)
	WebsiteDataTypes_Cookies                 WebsiteDataTypes = int64(256)
	WebsiteDataTypes_DeviceIdHashSalt        WebsiteDataTypes = int64(512)
	WebsiteDataTypes_HstsCache               WebsiteDataTypes = int64(1024)
	WebsiteDataTypes_All                     WebsiteDataTypes = int64(2047)
)
