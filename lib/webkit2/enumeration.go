// Code generated - DO NOT EDIT.

package webkit2

// AuthenticationScheme is a representation of the C type WebKitAuthenticationScheme.
//
// since 2.2
type AuthenticationScheme int

const (
	// AuthenticationScheme_Default is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_DEFAULT.
	AuthenticationScheme_Default AuthenticationScheme = 1
	// AuthenticationScheme_HttpBasic is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_HTTP_BASIC.
	AuthenticationScheme_HttpBasic AuthenticationScheme = 2
	// AuthenticationScheme_HttpDigest is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_HTTP_DIGEST.
	AuthenticationScheme_HttpDigest AuthenticationScheme = 3
	// AuthenticationScheme_HtmlForm is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_HTML_FORM.
	AuthenticationScheme_HtmlForm AuthenticationScheme = 4
	// AuthenticationScheme_Ntlm is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_NTLM.
	AuthenticationScheme_Ntlm AuthenticationScheme = 5
	// AuthenticationScheme_Negotiate is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_NEGOTIATE.
	AuthenticationScheme_Negotiate AuthenticationScheme = 6
	// AuthenticationScheme_ClientCertificateRequested is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_CLIENT_CERTIFICATE_REQUESTED.
	AuthenticationScheme_ClientCertificateRequested AuthenticationScheme = 7
	// AuthenticationScheme_ServerTrustEvaluationRequested is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_SERVER_TRUST_EVALUATION_REQUESTED.
	AuthenticationScheme_ServerTrustEvaluationRequested AuthenticationScheme = 8
	// AuthenticationScheme_Unknown is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_UNKNOWN.
	AuthenticationScheme_Unknown AuthenticationScheme = 100
)

// CacheModel is a representation of the C type WebKitCacheModel.
//
type CacheModel int

const (
	// CacheModel_DocumentViewer is a representation of the C type WEBKIT_CACHE_MODEL_DOCUMENT_VIEWER.
	CacheModel_DocumentViewer CacheModel = 0
	// CacheModel_WebBrowser is a representation of the C type WEBKIT_CACHE_MODEL_WEB_BROWSER.
	CacheModel_WebBrowser CacheModel = 1
	// CacheModel_DocumentBrowser is a representation of the C type WEBKIT_CACHE_MODEL_DOCUMENT_BROWSER.
	CacheModel_DocumentBrowser CacheModel = 2
)

// ContextMenuAction is a representation of the C type WebKitContextMenuAction.
//
type ContextMenuAction int

const (
	// ContextMenuAction_NoAction is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_NO_ACTION.
	ContextMenuAction_NoAction ContextMenuAction = 0
	// ContextMenuAction_OpenLink is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_LINK.
	ContextMenuAction_OpenLink ContextMenuAction = 1
	// ContextMenuAction_OpenLinkInNewWindow is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_LINK_IN_NEW_WINDOW.
	ContextMenuAction_OpenLinkInNewWindow ContextMenuAction = 2
	// ContextMenuAction_DownloadLinkToDisk is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_LINK_TO_DISK.
	ContextMenuAction_DownloadLinkToDisk ContextMenuAction = 3
	// ContextMenuAction_CopyLinkToClipboard is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY_LINK_TO_CLIPBOARD.
	ContextMenuAction_CopyLinkToClipboard ContextMenuAction = 4
	// ContextMenuAction_OpenImageInNewWindow is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_IMAGE_IN_NEW_WINDOW.
	ContextMenuAction_OpenImageInNewWindow ContextMenuAction = 5
	// ContextMenuAction_DownloadImageToDisk is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_IMAGE_TO_DISK.
	ContextMenuAction_DownloadImageToDisk ContextMenuAction = 6
	// ContextMenuAction_CopyImageToClipboard is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY_IMAGE_TO_CLIPBOARD.
	ContextMenuAction_CopyImageToClipboard ContextMenuAction = 7
	// ContextMenuAction_CopyImageUrlToClipboard is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY_IMAGE_URL_TO_CLIPBOARD.
	ContextMenuAction_CopyImageUrlToClipboard ContextMenuAction = 8
	// ContextMenuAction_OpenFrameInNewWindow is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_FRAME_IN_NEW_WINDOW.
	ContextMenuAction_OpenFrameInNewWindow ContextMenuAction = 9
	// ContextMenuAction_GoBack is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_GO_BACK.
	ContextMenuAction_GoBack ContextMenuAction = 10
	// ContextMenuAction_GoForward is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_GO_FORWARD.
	ContextMenuAction_GoForward ContextMenuAction = 11
	// ContextMenuAction_Stop is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_STOP.
	ContextMenuAction_Stop ContextMenuAction = 12
	// ContextMenuAction_Reload is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_RELOAD.
	ContextMenuAction_Reload ContextMenuAction = 13
	// ContextMenuAction_Copy is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY.
	ContextMenuAction_Copy ContextMenuAction = 14
	// ContextMenuAction_Cut is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_CUT.
	ContextMenuAction_Cut ContextMenuAction = 15
	// ContextMenuAction_Paste is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_PASTE.
	ContextMenuAction_Paste ContextMenuAction = 16
	// ContextMenuAction_Delete is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_DELETE.
	ContextMenuAction_Delete ContextMenuAction = 17
	// ContextMenuAction_SelectAll is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_SELECT_ALL.
	ContextMenuAction_SelectAll ContextMenuAction = 18
	// ContextMenuAction_InputMethods is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_INPUT_METHODS.
	ContextMenuAction_InputMethods ContextMenuAction = 19
	// ContextMenuAction_Unicode is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_UNICODE.
	ContextMenuAction_Unicode ContextMenuAction = 20
	// ContextMenuAction_SpellingGuess is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_SPELLING_GUESS.
	ContextMenuAction_SpellingGuess ContextMenuAction = 21
	// ContextMenuAction_NoGuessesFound is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_NO_GUESSES_FOUND.
	ContextMenuAction_NoGuessesFound ContextMenuAction = 22
	// ContextMenuAction_IgnoreSpelling is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_IGNORE_SPELLING.
	ContextMenuAction_IgnoreSpelling ContextMenuAction = 23
	// ContextMenuAction_LearnSpelling is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_LEARN_SPELLING.
	ContextMenuAction_LearnSpelling ContextMenuAction = 24
	// ContextMenuAction_IgnoreGrammar is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_IGNORE_GRAMMAR.
	ContextMenuAction_IgnoreGrammar ContextMenuAction = 25
	// ContextMenuAction_FontMenu is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_FONT_MENU.
	ContextMenuAction_FontMenu ContextMenuAction = 26
	// ContextMenuAction_Bold is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_BOLD.
	ContextMenuAction_Bold ContextMenuAction = 27
	// ContextMenuAction_Italic is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_ITALIC.
	ContextMenuAction_Italic ContextMenuAction = 28
	// ContextMenuAction_Underline is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_UNDERLINE.
	ContextMenuAction_Underline ContextMenuAction = 29
	// ContextMenuAction_Outline is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OUTLINE.
	ContextMenuAction_Outline ContextMenuAction = 30
	// ContextMenuAction_InspectElement is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_INSPECT_ELEMENT.
	ContextMenuAction_InspectElement ContextMenuAction = 31
	// ContextMenuAction_OpenVideoInNewWindow is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_VIDEO_IN_NEW_WINDOW.
	ContextMenuAction_OpenVideoInNewWindow ContextMenuAction = 32
	// ContextMenuAction_OpenAudioInNewWindow is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_AUDIO_IN_NEW_WINDOW.
	ContextMenuAction_OpenAudioInNewWindow ContextMenuAction = 33
	// ContextMenuAction_CopyVideoLinkToClipboard is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY_VIDEO_LINK_TO_CLIPBOARD.
	ContextMenuAction_CopyVideoLinkToClipboard ContextMenuAction = 34
	// ContextMenuAction_CopyAudioLinkToClipboard is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY_AUDIO_LINK_TO_CLIPBOARD.
	ContextMenuAction_CopyAudioLinkToClipboard ContextMenuAction = 35
	// ContextMenuAction_ToggleMediaControls is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_TOGGLE_MEDIA_CONTROLS.
	ContextMenuAction_ToggleMediaControls ContextMenuAction = 36
	// ContextMenuAction_ToggleMediaLoop is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_TOGGLE_MEDIA_LOOP.
	ContextMenuAction_ToggleMediaLoop ContextMenuAction = 37
	// ContextMenuAction_EnterVideoFullscreen is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_ENTER_VIDEO_FULLSCREEN.
	ContextMenuAction_EnterVideoFullscreen ContextMenuAction = 38
	// ContextMenuAction_MediaPlay is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_MEDIA_PLAY.
	ContextMenuAction_MediaPlay ContextMenuAction = 39
	// ContextMenuAction_MediaPause is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_MEDIA_PAUSE.
	ContextMenuAction_MediaPause ContextMenuAction = 40
	// ContextMenuAction_MediaMute is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_MEDIA_MUTE.
	ContextMenuAction_MediaMute ContextMenuAction = 41
	// ContextMenuAction_DownloadVideoToDisk is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_VIDEO_TO_DISK.
	ContextMenuAction_DownloadVideoToDisk ContextMenuAction = 42
	// ContextMenuAction_DownloadAudioToDisk is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_AUDIO_TO_DISK.
	ContextMenuAction_DownloadAudioToDisk ContextMenuAction = 43
	// ContextMenuAction_InsertEmoji is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_INSERT_EMOJI.
	ContextMenuAction_InsertEmoji ContextMenuAction = 44
	// ContextMenuAction_Custom is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_CUSTOM.
	ContextMenuAction_Custom ContextMenuAction = 10000
)

// CookieAcceptPolicy is a representation of the C type WebKitCookieAcceptPolicy.
//
type CookieAcceptPolicy int

const (
	// CookieAcceptPolicy_Always is a representation of the C type WEBKIT_COOKIE_POLICY_ACCEPT_ALWAYS.
	CookieAcceptPolicy_Always CookieAcceptPolicy = 0
	// CookieAcceptPolicy_Never is a representation of the C type WEBKIT_COOKIE_POLICY_ACCEPT_NEVER.
	CookieAcceptPolicy_Never CookieAcceptPolicy = 1
	// CookieAcceptPolicy_NoThirdParty is a representation of the C type WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY.
	CookieAcceptPolicy_NoThirdParty CookieAcceptPolicy = 2
)

// CookiePersistentStorage is a representation of the C type WebKitCookiePersistentStorage.
//
type CookiePersistentStorage int

const (
	// CookiePersistentStorage_Text is a representation of the C type WEBKIT_COOKIE_PERSISTENT_STORAGE_TEXT.
	CookiePersistentStorage_Text CookiePersistentStorage = 0
	// CookiePersistentStorage_Sqlite is a representation of the C type WEBKIT_COOKIE_PERSISTENT_STORAGE_SQLITE.
	CookiePersistentStorage_Sqlite CookiePersistentStorage = 1
)

// CredentialPersistence is a representation of the C type WebKitCredentialPersistence.
//
// since 2.2
type CredentialPersistence int

const (
	// CredentialPersistence_None is a representation of the C type WEBKIT_CREDENTIAL_PERSISTENCE_NONE.
	CredentialPersistence_None CredentialPersistence = 0
	// CredentialPersistence_ForSession is a representation of the C type WEBKIT_CREDENTIAL_PERSISTENCE_FOR_SESSION.
	CredentialPersistence_ForSession CredentialPersistence = 1
	// CredentialPersistence_Permanent is a representation of the C type WEBKIT_CREDENTIAL_PERSISTENCE_PERMANENT.
	CredentialPersistence_Permanent CredentialPersistence = 2
)

// DownloadError is a representation of the C type WebKitDownloadError.
//
type DownloadError int

const (
	// DownloadError_Network is a representation of the C type WEBKIT_DOWNLOAD_ERROR_NETWORK.
	DownloadError_Network DownloadError = 499
	// DownloadError_CancelledByUser is a representation of the C type WEBKIT_DOWNLOAD_ERROR_CANCELLED_BY_USER.
	DownloadError_CancelledByUser DownloadError = 400
	// DownloadError_Destination is a representation of the C type WEBKIT_DOWNLOAD_ERROR_DESTINATION.
	DownloadError_Destination DownloadError = 401
)

// FaviconDatabaseError is a representation of the C type WebKitFaviconDatabaseError.
//
type FaviconDatabaseError int

const (
	// FaviconDatabaseError_NotInitialized is a representation of the C type WEBKIT_FAVICON_DATABASE_ERROR_NOT_INITIALIZED.
	FaviconDatabaseError_NotInitialized FaviconDatabaseError = 0
	// FaviconDatabaseError_FaviconNotFound is a representation of the C type WEBKIT_FAVICON_DATABASE_ERROR_FAVICON_NOT_FOUND.
	FaviconDatabaseError_FaviconNotFound FaviconDatabaseError = 1
	// FaviconDatabaseError_FaviconUnknown is a representation of the C type WEBKIT_FAVICON_DATABASE_ERROR_FAVICON_UNKNOWN.
	FaviconDatabaseError_FaviconUnknown FaviconDatabaseError = 2
)

// HardwareAccelerationPolicy is a representation of the C type WebKitHardwareAccelerationPolicy.
//
// since 2.16
type HardwareAccelerationPolicy int

const (
	// HardwareAccelerationPolicy_OnDemand is a representation of the C type WEBKIT_HARDWARE_ACCELERATION_POLICY_ON_DEMAND.
	HardwareAccelerationPolicy_OnDemand HardwareAccelerationPolicy = 0
	// HardwareAccelerationPolicy_Always is a representation of the C type WEBKIT_HARDWARE_ACCELERATION_POLICY_ALWAYS.
	HardwareAccelerationPolicy_Always HardwareAccelerationPolicy = 1
	// HardwareAccelerationPolicy_Never is a representation of the C type WEBKIT_HARDWARE_ACCELERATION_POLICY_NEVER.
	HardwareAccelerationPolicy_Never HardwareAccelerationPolicy = 2
)

// InsecureContentEvent is a representation of the C type WebKitInsecureContentEvent.
//
type InsecureContentEvent int

const (
	// InsecureContentEvent_Run is a representation of the C type WEBKIT_INSECURE_CONTENT_RUN.
	InsecureContentEvent_Run InsecureContentEvent = 0
	// InsecureContentEvent_Displayed is a representation of the C type WEBKIT_INSECURE_CONTENT_DISPLAYED.
	InsecureContentEvent_Displayed InsecureContentEvent = 1
)

// JavascriptError is a representation of the C type WebKitJavascriptError.
//
type JavascriptError int

const (
	// JavascriptError_Failed is a representation of the C type WEBKIT_JAVASCRIPT_ERROR_SCRIPT_FAILED.
	JavascriptError_Failed JavascriptError = 699
)

// LoadEvent is a representation of the C type WebKitLoadEvent.
//
type LoadEvent int

const (
	// LoadEvent_Started is a representation of the C type WEBKIT_LOAD_STARTED.
	LoadEvent_Started LoadEvent = 0
	// LoadEvent_Redirected is a representation of the C type WEBKIT_LOAD_REDIRECTED.
	LoadEvent_Redirected LoadEvent = 1
	// LoadEvent_Committed is a representation of the C type WEBKIT_LOAD_COMMITTED.
	LoadEvent_Committed LoadEvent = 2
	// LoadEvent_Finished is a representation of the C type WEBKIT_LOAD_FINISHED.
	LoadEvent_Finished LoadEvent = 3
)

// NavigationType is a representation of the C type WebKitNavigationType.
//
type NavigationType int

const (
	// NavigationType_LinkClicked is a representation of the C type WEBKIT_NAVIGATION_TYPE_LINK_CLICKED.
	NavigationType_LinkClicked NavigationType = 0
	// NavigationType_FormSubmitted is a representation of the C type WEBKIT_NAVIGATION_TYPE_FORM_SUBMITTED.
	NavigationType_FormSubmitted NavigationType = 1
	// NavigationType_BackForward is a representation of the C type WEBKIT_NAVIGATION_TYPE_BACK_FORWARD.
	NavigationType_BackForward NavigationType = 2
	// NavigationType_Reload is a representation of the C type WEBKIT_NAVIGATION_TYPE_RELOAD.
	NavigationType_Reload NavigationType = 3
	// NavigationType_FormResubmitted is a representation of the C type WEBKIT_NAVIGATION_TYPE_FORM_RESUBMITTED.
	NavigationType_FormResubmitted NavigationType = 4
	// NavigationType_Other is a representation of the C type WEBKIT_NAVIGATION_TYPE_OTHER.
	NavigationType_Other NavigationType = 5
)

// NetworkError is a representation of the C type WebKitNetworkError.
//
type NetworkError int

const (
	// NetworkError_Failed is a representation of the C type WEBKIT_NETWORK_ERROR_FAILED.
	NetworkError_Failed NetworkError = 399
	// NetworkError_Transport is a representation of the C type WEBKIT_NETWORK_ERROR_TRANSPORT.
	NetworkError_Transport NetworkError = 300
	// NetworkError_UnknownProtocol is a representation of the C type WEBKIT_NETWORK_ERROR_UNKNOWN_PROTOCOL.
	NetworkError_UnknownProtocol NetworkError = 301
	// NetworkError_Cancelled is a representation of the C type WEBKIT_NETWORK_ERROR_CANCELLED.
	NetworkError_Cancelled NetworkError = 302
	// NetworkError_FileDoesNotExist is a representation of the C type WEBKIT_NETWORK_ERROR_FILE_DOES_NOT_EXIST.
	NetworkError_FileDoesNotExist NetworkError = 303
)

// NetworkProxyMode is a representation of the C type WebKitNetworkProxyMode.
//
// since 2.16
type NetworkProxyMode int

const (
	// NetworkProxyMode_Default is a representation of the C type WEBKIT_NETWORK_PROXY_MODE_DEFAULT.
	NetworkProxyMode_Default NetworkProxyMode = 0
	// NetworkProxyMode_NoProxy is a representation of the C type WEBKIT_NETWORK_PROXY_MODE_NO_PROXY.
	NetworkProxyMode_NoProxy NetworkProxyMode = 1
	// NetworkProxyMode_Custom is a representation of the C type WEBKIT_NETWORK_PROXY_MODE_CUSTOM.
	NetworkProxyMode_Custom NetworkProxyMode = 2
)

// PluginError is a representation of the C type WebKitPluginError.
//
type PluginError int

const (
	// PluginError_Failed is a representation of the C type WEBKIT_PLUGIN_ERROR_FAILED.
	PluginError_Failed PluginError = 299
	// PluginError_CannotFindPlugin is a representation of the C type WEBKIT_PLUGIN_ERROR_CANNOT_FIND_PLUGIN.
	PluginError_CannotFindPlugin PluginError = 200
	// PluginError_CannotLoadPlugin is a representation of the C type WEBKIT_PLUGIN_ERROR_CANNOT_LOAD_PLUGIN.
	PluginError_CannotLoadPlugin PluginError = 201
	// PluginError_JavaUnavailable is a representation of the C type WEBKIT_PLUGIN_ERROR_JAVA_UNAVAILABLE.
	PluginError_JavaUnavailable PluginError = 202
	// PluginError_ConnectionCancelled is a representation of the C type WEBKIT_PLUGIN_ERROR_CONNECTION_CANCELLED.
	PluginError_ConnectionCancelled PluginError = 203
	// PluginError_WillHandleLoad is a representation of the C type WEBKIT_PLUGIN_ERROR_WILL_HANDLE_LOAD.
	PluginError_WillHandleLoad PluginError = 204
)

// PolicyDecisionType is a representation of the C type WebKitPolicyDecisionType.
//
type PolicyDecisionType int

const (
	// PolicyDecisionType_NavigationAction is a representation of the C type WEBKIT_POLICY_DECISION_TYPE_NAVIGATION_ACTION.
	PolicyDecisionType_NavigationAction PolicyDecisionType = 0
	// PolicyDecisionType_NewWindowAction is a representation of the C type WEBKIT_POLICY_DECISION_TYPE_NEW_WINDOW_ACTION.
	PolicyDecisionType_NewWindowAction PolicyDecisionType = 1
	// PolicyDecisionType_Response is a representation of the C type WEBKIT_POLICY_DECISION_TYPE_RESPONSE.
	PolicyDecisionType_Response PolicyDecisionType = 2
)

// PolicyError is a representation of the C type WebKitPolicyError.
//
type PolicyError int

const (
	// PolicyError_Failed is a representation of the C type WEBKIT_POLICY_ERROR_FAILED.
	PolicyError_Failed PolicyError = 199
	// PolicyError_CannotShowMimeType is a representation of the C type WEBKIT_POLICY_ERROR_CANNOT_SHOW_MIME_TYPE.
	PolicyError_CannotShowMimeType PolicyError = 100
	// PolicyError_CannotShowUri is a representation of the C type WEBKIT_POLICY_ERROR_CANNOT_SHOW_URI.
	PolicyError_CannotShowUri PolicyError = 101
	// PolicyError_FrameLoadInterruptedByPolicyChange is a representation of the C type WEBKIT_POLICY_ERROR_FRAME_LOAD_INTERRUPTED_BY_POLICY_CHANGE.
	PolicyError_FrameLoadInterruptedByPolicyChange PolicyError = 102
	// PolicyError_CannotUseRestrictedPort is a representation of the C type WEBKIT_POLICY_ERROR_CANNOT_USE_RESTRICTED_PORT.
	PolicyError_CannotUseRestrictedPort PolicyError = 103
)

// PrintError is a representation of the C type WebKitPrintError.
//
type PrintError int

const (
	// PrintError_General is a representation of the C type WEBKIT_PRINT_ERROR_GENERAL.
	PrintError_General PrintError = 599
	// PrintError_PrinterNotFound is a representation of the C type WEBKIT_PRINT_ERROR_PRINTER_NOT_FOUND.
	PrintError_PrinterNotFound PrintError = 500
	// PrintError_InvalidPageRange is a representation of the C type WEBKIT_PRINT_ERROR_INVALID_PAGE_RANGE.
	PrintError_InvalidPageRange PrintError = 501
)

// PrintOperationResponse is a representation of the C type WebKitPrintOperationResponse.
//
type PrintOperationResponse int

const (
	// PrintOperationResponse_Print is a representation of the C type WEBKIT_PRINT_OPERATION_RESPONSE_PRINT.
	PrintOperationResponse_Print PrintOperationResponse = 0
	// PrintOperationResponse_Cancel is a representation of the C type WEBKIT_PRINT_OPERATION_RESPONSE_CANCEL.
	PrintOperationResponse_Cancel PrintOperationResponse = 1
)

// ProcessModel is a representation of the C type WebKitProcessModel.
//
// since 2.4
type ProcessModel int

const (
	// ProcessModel_SharedSecondaryProcess is a representation of the C type WEBKIT_PROCESS_MODEL_SHARED_SECONDARY_PROCESS.
	ProcessModel_SharedSecondaryProcess ProcessModel = 0
	// ProcessModel_MultipleSecondaryProcesses is a representation of the C type WEBKIT_PROCESS_MODEL_MULTIPLE_SECONDARY_PROCESSES.
	ProcessModel_MultipleSecondaryProcesses ProcessModel = 1
)

// SaveMode is a representation of the C type WebKitSaveMode.
//
type SaveMode int

const (
	// SaveMode_Mhtml is a representation of the C type WEBKIT_SAVE_MODE_MHTML.
	SaveMode_Mhtml SaveMode = 0
)

// ScriptDialogType is a representation of the C type WebKitScriptDialogType.
//
type ScriptDialogType int

const (
	// ScriptDialogType_Alert is a representation of the C type WEBKIT_SCRIPT_DIALOG_ALERT.
	ScriptDialogType_Alert ScriptDialogType = 0
	// ScriptDialogType_Confirm is a representation of the C type WEBKIT_SCRIPT_DIALOG_CONFIRM.
	ScriptDialogType_Confirm ScriptDialogType = 1
	// ScriptDialogType_Prompt is a representation of the C type WEBKIT_SCRIPT_DIALOG_PROMPT.
	ScriptDialogType_Prompt ScriptDialogType = 2
	// ScriptDialogType_BeforeUnloadConfirm is a representation of the C type WEBKIT_SCRIPT_DIALOG_BEFORE_UNLOAD_CONFIRM.
	ScriptDialogType_BeforeUnloadConfirm ScriptDialogType = 3
)

// SnapshotError is a representation of the C type WebKitSnapshotError.
//
type SnapshotError int

const (
	// SnapshotError_Create is a representation of the C type WEBKIT_SNAPSHOT_ERROR_FAILED_TO_CREATE.
	SnapshotError_Create SnapshotError = 799
)

// SnapshotRegion is a representation of the C type WebKitSnapshotRegion.
//
type SnapshotRegion int

const (
	// SnapshotRegion_Visible is a representation of the C type WEBKIT_SNAPSHOT_REGION_VISIBLE.
	SnapshotRegion_Visible SnapshotRegion = 0
	// SnapshotRegion_FullDocument is a representation of the C type WEBKIT_SNAPSHOT_REGION_FULL_DOCUMENT.
	SnapshotRegion_FullDocument SnapshotRegion = 1
)

// TLSErrorsPolicy is a representation of the C type WebKitTLSErrorsPolicy.
//
type TLSErrorsPolicy int

const (
	// TLSErrorsPolicy_Ignore is a representation of the C type WEBKIT_TLS_ERRORS_POLICY_IGNORE.
	TLSErrorsPolicy_Ignore TLSErrorsPolicy = 0
	// TLSErrorsPolicy_Fail is a representation of the C type WEBKIT_TLS_ERRORS_POLICY_FAIL.
	TLSErrorsPolicy_Fail TLSErrorsPolicy = 1
)

// UserContentFilterError is a representation of the C type WebKitUserContentFilterError.
//
// since 2.24
type UserContentFilterError int

const (
	// UserContentFilterError_InvalidSource is a representation of the C type WEBKIT_USER_CONTENT_FILTER_ERROR_INVALID_SOURCE.
	UserContentFilterError_InvalidSource UserContentFilterError = 0
	// UserContentFilterError_NotFound is a representation of the C type WEBKIT_USER_CONTENT_FILTER_ERROR_NOT_FOUND.
	UserContentFilterError_NotFound UserContentFilterError = 1
)

// UserContentInjectedFrames is a representation of the C type WebKitUserContentInjectedFrames.
//
// since 2.6
type UserContentInjectedFrames int

const (
	// UserContentInjectedFrames_AllFrames is a representation of the C type WEBKIT_USER_CONTENT_INJECT_ALL_FRAMES.
	UserContentInjectedFrames_AllFrames UserContentInjectedFrames = 0
	// UserContentInjectedFrames_TopFrame is a representation of the C type WEBKIT_USER_CONTENT_INJECT_TOP_FRAME.
	UserContentInjectedFrames_TopFrame UserContentInjectedFrames = 1
)

// UserScriptInjectionTime is a representation of the C type WebKitUserScriptInjectionTime.
//
// since 2.6
type UserScriptInjectionTime int

const (
	// UserScriptInjectionTime_Start is a representation of the C type WEBKIT_USER_SCRIPT_INJECT_AT_DOCUMENT_START.
	UserScriptInjectionTime_Start UserScriptInjectionTime = 0
	// UserScriptInjectionTime_End is a representation of the C type WEBKIT_USER_SCRIPT_INJECT_AT_DOCUMENT_END.
	UserScriptInjectionTime_End UserScriptInjectionTime = 1
)

// UserStyleLevel is a representation of the C type WebKitUserStyleLevel.
//
// since 2.6
type UserStyleLevel int

const (
	// UserStyleLevel_User is a representation of the C type WEBKIT_USER_STYLE_LEVEL_USER.
	UserStyleLevel_User UserStyleLevel = 0
	// UserStyleLevel_Author is a representation of the C type WEBKIT_USER_STYLE_LEVEL_AUTHOR.
	UserStyleLevel_Author UserStyleLevel = 1
)

// WebProcessTerminationReason is a representation of the C type WebKitWebProcessTerminationReason.
//
// since 2.20
type WebProcessTerminationReason int

const (
	// WebProcessTerminationReason_Crashed is a representation of the C type WEBKIT_WEB_PROCESS_CRASHED.
	WebProcessTerminationReason_Crashed WebProcessTerminationReason = 0
	// WebProcessTerminationReason_ExceededMemoryLimit is a representation of the C type WEBKIT_WEB_PROCESS_EXCEEDED_MEMORY_LIMIT.
	WebProcessTerminationReason_ExceededMemoryLimit WebProcessTerminationReason = 1
)
