// Code generated - DO NOT EDIT.

package webkit2

// Authenticationscheme is a representation of the C type WebKitAuthenticationScheme.
//
// since 2.2
type Authenticationscheme int

const (
	// Authenticationscheme_Default is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_DEFAULT.
	Authenticationscheme_Default Authenticationscheme = 1
	// Authenticationscheme_HttpBasic is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_HTTP_BASIC.
	Authenticationscheme_HttpBasic Authenticationscheme = 2
	// Authenticationscheme_HttpDigest is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_HTTP_DIGEST.
	Authenticationscheme_HttpDigest Authenticationscheme = 3
	// Authenticationscheme_HtmlForm is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_HTML_FORM.
	Authenticationscheme_HtmlForm Authenticationscheme = 4
	// Authenticationscheme_Ntlm is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_NTLM.
	Authenticationscheme_Ntlm Authenticationscheme = 5
	// Authenticationscheme_Negotiate is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_NEGOTIATE.
	Authenticationscheme_Negotiate Authenticationscheme = 6
	// Authenticationscheme_ClientCertificateRequested is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_CLIENT_CERTIFICATE_REQUESTED.
	Authenticationscheme_ClientCertificateRequested Authenticationscheme = 7
	// Authenticationscheme_ServerTrustEvaluationRequested is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_SERVER_TRUST_EVALUATION_REQUESTED.
	Authenticationscheme_ServerTrustEvaluationRequested Authenticationscheme = 8
	// Authenticationscheme_Unknown is a representation of the C type WEBKIT_AUTHENTICATION_SCHEME_UNKNOWN.
	Authenticationscheme_Unknown Authenticationscheme = 100
)

// Cachemodel is a representation of the C type WebKitCacheModel.
type Cachemodel int

const (
	// Cachemodel_DocumentViewer is a representation of the C type WEBKIT_CACHE_MODEL_DOCUMENT_VIEWER.
	Cachemodel_DocumentViewer Cachemodel = 0
	// Cachemodel_WebBrowser is a representation of the C type WEBKIT_CACHE_MODEL_WEB_BROWSER.
	Cachemodel_WebBrowser Cachemodel = 1
	// Cachemodel_DocumentBrowser is a representation of the C type WEBKIT_CACHE_MODEL_DOCUMENT_BROWSER.
	Cachemodel_DocumentBrowser Cachemodel = 2
)

// Contextmenuaction is a representation of the C type WebKitContextMenuAction.
type Contextmenuaction int

const (
	// Contextmenuaction_NoAction is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_NO_ACTION.
	Contextmenuaction_NoAction Contextmenuaction = 0
	// Contextmenuaction_OpenLink is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_LINK.
	Contextmenuaction_OpenLink Contextmenuaction = 1
	// Contextmenuaction_OpenLinkInNewWindow is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_LINK_IN_NEW_WINDOW.
	Contextmenuaction_OpenLinkInNewWindow Contextmenuaction = 2
	// Contextmenuaction_DownloadLinkToDisk is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_LINK_TO_DISK.
	Contextmenuaction_DownloadLinkToDisk Contextmenuaction = 3
	// Contextmenuaction_CopyLinkToClipboard is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY_LINK_TO_CLIPBOARD.
	Contextmenuaction_CopyLinkToClipboard Contextmenuaction = 4
	// Contextmenuaction_OpenImageInNewWindow is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_IMAGE_IN_NEW_WINDOW.
	Contextmenuaction_OpenImageInNewWindow Contextmenuaction = 5
	// Contextmenuaction_DownloadImageToDisk is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_IMAGE_TO_DISK.
	Contextmenuaction_DownloadImageToDisk Contextmenuaction = 6
	// Contextmenuaction_CopyImageToClipboard is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY_IMAGE_TO_CLIPBOARD.
	Contextmenuaction_CopyImageToClipboard Contextmenuaction = 7
	// Contextmenuaction_CopyImageUrlToClipboard is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY_IMAGE_URL_TO_CLIPBOARD.
	Contextmenuaction_CopyImageUrlToClipboard Contextmenuaction = 8
	// Contextmenuaction_OpenFrameInNewWindow is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_FRAME_IN_NEW_WINDOW.
	Contextmenuaction_OpenFrameInNewWindow Contextmenuaction = 9
	// Contextmenuaction_GoBack is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_GO_BACK.
	Contextmenuaction_GoBack Contextmenuaction = 10
	// Contextmenuaction_GoForward is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_GO_FORWARD.
	Contextmenuaction_GoForward Contextmenuaction = 11
	// Contextmenuaction_Stop is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_STOP.
	Contextmenuaction_Stop Contextmenuaction = 12
	// Contextmenuaction_Reload is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_RELOAD.
	Contextmenuaction_Reload Contextmenuaction = 13
	// Contextmenuaction_Copy is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY.
	Contextmenuaction_Copy Contextmenuaction = 14
	// Contextmenuaction_Cut is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_CUT.
	Contextmenuaction_Cut Contextmenuaction = 15
	// Contextmenuaction_Paste is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_PASTE.
	Contextmenuaction_Paste Contextmenuaction = 16
	// Contextmenuaction_Delete is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_DELETE.
	Contextmenuaction_Delete Contextmenuaction = 17
	// Contextmenuaction_SelectAll is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_SELECT_ALL.
	Contextmenuaction_SelectAll Contextmenuaction = 18
	// Contextmenuaction_InputMethods is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_INPUT_METHODS.
	Contextmenuaction_InputMethods Contextmenuaction = 19
	// Contextmenuaction_Unicode is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_UNICODE.
	Contextmenuaction_Unicode Contextmenuaction = 20
	// Contextmenuaction_SpellingGuess is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_SPELLING_GUESS.
	Contextmenuaction_SpellingGuess Contextmenuaction = 21
	// Contextmenuaction_NoGuessesFound is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_NO_GUESSES_FOUND.
	Contextmenuaction_NoGuessesFound Contextmenuaction = 22
	// Contextmenuaction_IgnoreSpelling is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_IGNORE_SPELLING.
	Contextmenuaction_IgnoreSpelling Contextmenuaction = 23
	// Contextmenuaction_LearnSpelling is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_LEARN_SPELLING.
	Contextmenuaction_LearnSpelling Contextmenuaction = 24
	// Contextmenuaction_IgnoreGrammar is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_IGNORE_GRAMMAR.
	Contextmenuaction_IgnoreGrammar Contextmenuaction = 25
	// Contextmenuaction_FontMenu is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_FONT_MENU.
	Contextmenuaction_FontMenu Contextmenuaction = 26
	// Contextmenuaction_Bold is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_BOLD.
	Contextmenuaction_Bold Contextmenuaction = 27
	// Contextmenuaction_Italic is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_ITALIC.
	Contextmenuaction_Italic Contextmenuaction = 28
	// Contextmenuaction_Underline is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_UNDERLINE.
	Contextmenuaction_Underline Contextmenuaction = 29
	// Contextmenuaction_Outline is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OUTLINE.
	Contextmenuaction_Outline Contextmenuaction = 30
	// Contextmenuaction_InspectElement is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_INSPECT_ELEMENT.
	Contextmenuaction_InspectElement Contextmenuaction = 31
	// Contextmenuaction_OpenVideoInNewWindow is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_VIDEO_IN_NEW_WINDOW.
	Contextmenuaction_OpenVideoInNewWindow Contextmenuaction = 32
	// Contextmenuaction_OpenAudioInNewWindow is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_OPEN_AUDIO_IN_NEW_WINDOW.
	Contextmenuaction_OpenAudioInNewWindow Contextmenuaction = 33
	// Contextmenuaction_CopyVideoLinkToClipboard is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY_VIDEO_LINK_TO_CLIPBOARD.
	Contextmenuaction_CopyVideoLinkToClipboard Contextmenuaction = 34
	// Contextmenuaction_CopyAudioLinkToClipboard is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_COPY_AUDIO_LINK_TO_CLIPBOARD.
	Contextmenuaction_CopyAudioLinkToClipboard Contextmenuaction = 35
	// Contextmenuaction_ToggleMediaControls is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_TOGGLE_MEDIA_CONTROLS.
	Contextmenuaction_ToggleMediaControls Contextmenuaction = 36
	// Contextmenuaction_ToggleMediaLoop is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_TOGGLE_MEDIA_LOOP.
	Contextmenuaction_ToggleMediaLoop Contextmenuaction = 37
	// Contextmenuaction_EnterVideoFullscreen is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_ENTER_VIDEO_FULLSCREEN.
	Contextmenuaction_EnterVideoFullscreen Contextmenuaction = 38
	// Contextmenuaction_MediaPlay is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_MEDIA_PLAY.
	Contextmenuaction_MediaPlay Contextmenuaction = 39
	// Contextmenuaction_MediaPause is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_MEDIA_PAUSE.
	Contextmenuaction_MediaPause Contextmenuaction = 40
	// Contextmenuaction_MediaMute is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_MEDIA_MUTE.
	Contextmenuaction_MediaMute Contextmenuaction = 41
	// Contextmenuaction_DownloadVideoToDisk is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_VIDEO_TO_DISK.
	Contextmenuaction_DownloadVideoToDisk Contextmenuaction = 42
	// Contextmenuaction_DownloadAudioToDisk is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_AUDIO_TO_DISK.
	Contextmenuaction_DownloadAudioToDisk Contextmenuaction = 43
	// Contextmenuaction_InsertEmoji is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_INSERT_EMOJI.
	Contextmenuaction_InsertEmoji Contextmenuaction = 44
	// Contextmenuaction_Custom is a representation of the C type WEBKIT_CONTEXT_MENU_ACTION_CUSTOM.
	Contextmenuaction_Custom Contextmenuaction = 10000
)

// Cookieacceptpolicy is a representation of the C type WebKitCookieAcceptPolicy.
type Cookieacceptpolicy int

const (
	// Cookieacceptpolicy_Always is a representation of the C type WEBKIT_COOKIE_POLICY_ACCEPT_ALWAYS.
	Cookieacceptpolicy_Always Cookieacceptpolicy = 0
	// Cookieacceptpolicy_Never is a representation of the C type WEBKIT_COOKIE_POLICY_ACCEPT_NEVER.
	Cookieacceptpolicy_Never Cookieacceptpolicy = 1
	// Cookieacceptpolicy_NoThirdParty is a representation of the C type WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY.
	Cookieacceptpolicy_NoThirdParty Cookieacceptpolicy = 2
)

// Cookiepersistentstorage is a representation of the C type WebKitCookiePersistentStorage.
type Cookiepersistentstorage int

const (
	// Cookiepersistentstorage_Text is a representation of the C type WEBKIT_COOKIE_PERSISTENT_STORAGE_TEXT.
	Cookiepersistentstorage_Text Cookiepersistentstorage = 0
	// Cookiepersistentstorage_Sqlite is a representation of the C type WEBKIT_COOKIE_PERSISTENT_STORAGE_SQLITE.
	Cookiepersistentstorage_Sqlite Cookiepersistentstorage = 1
)

// Credentialpersistence is a representation of the C type WebKitCredentialPersistence.
//
// since 2.2
type Credentialpersistence int

const (
	// Credentialpersistence_None is a representation of the C type WEBKIT_CREDENTIAL_PERSISTENCE_NONE.
	Credentialpersistence_None Credentialpersistence = 0
	// Credentialpersistence_ForSession is a representation of the C type WEBKIT_CREDENTIAL_PERSISTENCE_FOR_SESSION.
	Credentialpersistence_ForSession Credentialpersistence = 1
	// Credentialpersistence_Permanent is a representation of the C type WEBKIT_CREDENTIAL_PERSISTENCE_PERMANENT.
	Credentialpersistence_Permanent Credentialpersistence = 2
)

// Downloaderror is a representation of the C type WebKitDownloadError.
type Downloaderror int

const (
	// Downloaderror_Network is a representation of the C type WEBKIT_DOWNLOAD_ERROR_NETWORK.
	Downloaderror_Network Downloaderror = 499
	// Downloaderror_CancelledByUser is a representation of the C type WEBKIT_DOWNLOAD_ERROR_CANCELLED_BY_USER.
	Downloaderror_CancelledByUser Downloaderror = 400
	// Downloaderror_Destination is a representation of the C type WEBKIT_DOWNLOAD_ERROR_DESTINATION.
	Downloaderror_Destination Downloaderror = 401
)

// Favicondatabaseerror is a representation of the C type WebKitFaviconDatabaseError.
type Favicondatabaseerror int

const (
	// Favicondatabaseerror_NotInitialized is a representation of the C type WEBKIT_FAVICON_DATABASE_ERROR_NOT_INITIALIZED.
	Favicondatabaseerror_NotInitialized Favicondatabaseerror = 0
	// Favicondatabaseerror_FaviconNotFound is a representation of the C type WEBKIT_FAVICON_DATABASE_ERROR_FAVICON_NOT_FOUND.
	Favicondatabaseerror_FaviconNotFound Favicondatabaseerror = 1
	// Favicondatabaseerror_FaviconUnknown is a representation of the C type WEBKIT_FAVICON_DATABASE_ERROR_FAVICON_UNKNOWN.
	Favicondatabaseerror_FaviconUnknown Favicondatabaseerror = 2
)

// Hardwareaccelerationpolicy is a representation of the C type WebKitHardwareAccelerationPolicy.
//
// since 2.16
type Hardwareaccelerationpolicy int

const (
	// Hardwareaccelerationpolicy_OnDemand is a representation of the C type WEBKIT_HARDWARE_ACCELERATION_POLICY_ON_DEMAND.
	Hardwareaccelerationpolicy_OnDemand Hardwareaccelerationpolicy = 0
	// Hardwareaccelerationpolicy_Always is a representation of the C type WEBKIT_HARDWARE_ACCELERATION_POLICY_ALWAYS.
	Hardwareaccelerationpolicy_Always Hardwareaccelerationpolicy = 1
	// Hardwareaccelerationpolicy_Never is a representation of the C type WEBKIT_HARDWARE_ACCELERATION_POLICY_NEVER.
	Hardwareaccelerationpolicy_Never Hardwareaccelerationpolicy = 2
)

// Insecurecontentevent is a representation of the C type WebKitInsecureContentEvent.
type Insecurecontentevent int

const (
	// Insecurecontentevent_Run is a representation of the C type WEBKIT_INSECURE_CONTENT_RUN.
	Insecurecontentevent_Run Insecurecontentevent = 0
	// Insecurecontentevent_Displayed is a representation of the C type WEBKIT_INSECURE_CONTENT_DISPLAYED.
	Insecurecontentevent_Displayed Insecurecontentevent = 1
)

// Javascripterror is a representation of the C type WebKitJavascriptError.
type Javascripterror int

const (
	// Javascripterror_Failed is a representation of the C type WEBKIT_JAVASCRIPT_ERROR_SCRIPT_FAILED.
	Javascripterror_Failed Javascripterror = 699
)

// Loadevent is a representation of the C type WebKitLoadEvent.
type Loadevent int

const (
	// Loadevent_Started is a representation of the C type WEBKIT_LOAD_STARTED.
	Loadevent_Started Loadevent = 0
	// Loadevent_Redirected is a representation of the C type WEBKIT_LOAD_REDIRECTED.
	Loadevent_Redirected Loadevent = 1
	// Loadevent_Committed is a representation of the C type WEBKIT_LOAD_COMMITTED.
	Loadevent_Committed Loadevent = 2
	// Loadevent_Finished is a representation of the C type WEBKIT_LOAD_FINISHED.
	Loadevent_Finished Loadevent = 3
)

// Navigationtype is a representation of the C type WebKitNavigationType.
type Navigationtype int

const (
	// Navigationtype_LinkClicked is a representation of the C type WEBKIT_NAVIGATION_TYPE_LINK_CLICKED.
	Navigationtype_LinkClicked Navigationtype = 0
	// Navigationtype_FormSubmitted is a representation of the C type WEBKIT_NAVIGATION_TYPE_FORM_SUBMITTED.
	Navigationtype_FormSubmitted Navigationtype = 1
	// Navigationtype_BackForward is a representation of the C type WEBKIT_NAVIGATION_TYPE_BACK_FORWARD.
	Navigationtype_BackForward Navigationtype = 2
	// Navigationtype_Reload is a representation of the C type WEBKIT_NAVIGATION_TYPE_RELOAD.
	Navigationtype_Reload Navigationtype = 3
	// Navigationtype_FormResubmitted is a representation of the C type WEBKIT_NAVIGATION_TYPE_FORM_RESUBMITTED.
	Navigationtype_FormResubmitted Navigationtype = 4
	// Navigationtype_Other is a representation of the C type WEBKIT_NAVIGATION_TYPE_OTHER.
	Navigationtype_Other Navigationtype = 5
)

// Networkerror is a representation of the C type WebKitNetworkError.
type Networkerror int

const (
	// Networkerror_Failed is a representation of the C type WEBKIT_NETWORK_ERROR_FAILED.
	Networkerror_Failed Networkerror = 399
	// Networkerror_Transport is a representation of the C type WEBKIT_NETWORK_ERROR_TRANSPORT.
	Networkerror_Transport Networkerror = 300
	// Networkerror_UnknownProtocol is a representation of the C type WEBKIT_NETWORK_ERROR_UNKNOWN_PROTOCOL.
	Networkerror_UnknownProtocol Networkerror = 301
	// Networkerror_Cancelled is a representation of the C type WEBKIT_NETWORK_ERROR_CANCELLED.
	Networkerror_Cancelled Networkerror = 302
	// Networkerror_FileDoesNotExist is a representation of the C type WEBKIT_NETWORK_ERROR_FILE_DOES_NOT_EXIST.
	Networkerror_FileDoesNotExist Networkerror = 303
)

// Networkproxymode is a representation of the C type WebKitNetworkProxyMode.
//
// since 2.16
type Networkproxymode int

const (
	// Networkproxymode_Default is a representation of the C type WEBKIT_NETWORK_PROXY_MODE_DEFAULT.
	Networkproxymode_Default Networkproxymode = 0
	// Networkproxymode_NoProxy is a representation of the C type WEBKIT_NETWORK_PROXY_MODE_NO_PROXY.
	Networkproxymode_NoProxy Networkproxymode = 1
	// Networkproxymode_Custom is a representation of the C type WEBKIT_NETWORK_PROXY_MODE_CUSTOM.
	Networkproxymode_Custom Networkproxymode = 2
)

// Pluginerror is a representation of the C type WebKitPluginError.
type Pluginerror int

const (
	// Pluginerror_Failed is a representation of the C type WEBKIT_PLUGIN_ERROR_FAILED.
	Pluginerror_Failed Pluginerror = 299
	// Pluginerror_CannotFindPlugin is a representation of the C type WEBKIT_PLUGIN_ERROR_CANNOT_FIND_PLUGIN.
	Pluginerror_CannotFindPlugin Pluginerror = 200
	// Pluginerror_CannotLoadPlugin is a representation of the C type WEBKIT_PLUGIN_ERROR_CANNOT_LOAD_PLUGIN.
	Pluginerror_CannotLoadPlugin Pluginerror = 201
	// Pluginerror_JavaUnavailable is a representation of the C type WEBKIT_PLUGIN_ERROR_JAVA_UNAVAILABLE.
	Pluginerror_JavaUnavailable Pluginerror = 202
	// Pluginerror_ConnectionCancelled is a representation of the C type WEBKIT_PLUGIN_ERROR_CONNECTION_CANCELLED.
	Pluginerror_ConnectionCancelled Pluginerror = 203
	// Pluginerror_WillHandleLoad is a representation of the C type WEBKIT_PLUGIN_ERROR_WILL_HANDLE_LOAD.
	Pluginerror_WillHandleLoad Pluginerror = 204
)

// Policydecisiontype is a representation of the C type WebKitPolicyDecisionType.
type Policydecisiontype int

const (
	// Policydecisiontype_NavigationAction is a representation of the C type WEBKIT_POLICY_DECISION_TYPE_NAVIGATION_ACTION.
	Policydecisiontype_NavigationAction Policydecisiontype = 0
	// Policydecisiontype_NewWindowAction is a representation of the C type WEBKIT_POLICY_DECISION_TYPE_NEW_WINDOW_ACTION.
	Policydecisiontype_NewWindowAction Policydecisiontype = 1
	// Policydecisiontype_Response is a representation of the C type WEBKIT_POLICY_DECISION_TYPE_RESPONSE.
	Policydecisiontype_Response Policydecisiontype = 2
)

// Policyerror is a representation of the C type WebKitPolicyError.
type Policyerror int

const (
	// Policyerror_Failed is a representation of the C type WEBKIT_POLICY_ERROR_FAILED.
	Policyerror_Failed Policyerror = 199
	// Policyerror_CannotShowMimeType is a representation of the C type WEBKIT_POLICY_ERROR_CANNOT_SHOW_MIME_TYPE.
	Policyerror_CannotShowMimeType Policyerror = 100
	// Policyerror_CannotShowUri is a representation of the C type WEBKIT_POLICY_ERROR_CANNOT_SHOW_URI.
	Policyerror_CannotShowUri Policyerror = 101
	// Policyerror_FrameLoadInterruptedByPolicyChange is a representation of the C type WEBKIT_POLICY_ERROR_FRAME_LOAD_INTERRUPTED_BY_POLICY_CHANGE.
	Policyerror_FrameLoadInterruptedByPolicyChange Policyerror = 102
	// Policyerror_CannotUseRestrictedPort is a representation of the C type WEBKIT_POLICY_ERROR_CANNOT_USE_RESTRICTED_PORT.
	Policyerror_CannotUseRestrictedPort Policyerror = 103
)

// Printerror is a representation of the C type WebKitPrintError.
type Printerror int

const (
	// Printerror_General is a representation of the C type WEBKIT_PRINT_ERROR_GENERAL.
	Printerror_General Printerror = 599
	// Printerror_PrinterNotFound is a representation of the C type WEBKIT_PRINT_ERROR_PRINTER_NOT_FOUND.
	Printerror_PrinterNotFound Printerror = 500
	// Printerror_InvalidPageRange is a representation of the C type WEBKIT_PRINT_ERROR_INVALID_PAGE_RANGE.
	Printerror_InvalidPageRange Printerror = 501
)

// Printoperationresponse is a representation of the C type WebKitPrintOperationResponse.
type Printoperationresponse int

const (
	// Printoperationresponse_Print is a representation of the C type WEBKIT_PRINT_OPERATION_RESPONSE_PRINT.
	Printoperationresponse_Print Printoperationresponse = 0
	// Printoperationresponse_Cancel is a representation of the C type WEBKIT_PRINT_OPERATION_RESPONSE_CANCEL.
	Printoperationresponse_Cancel Printoperationresponse = 1
)

// Processmodel is a representation of the C type WebKitProcessModel.
//
// since 2.4
type Processmodel int

const (
	// Processmodel_SharedSecondaryProcess is a representation of the C type WEBKIT_PROCESS_MODEL_SHARED_SECONDARY_PROCESS.
	Processmodel_SharedSecondaryProcess Processmodel = 0
	// Processmodel_MultipleSecondaryProcesses is a representation of the C type WEBKIT_PROCESS_MODEL_MULTIPLE_SECONDARY_PROCESSES.
	Processmodel_MultipleSecondaryProcesses Processmodel = 1
)

// Savemode is a representation of the C type WebKitSaveMode.
type Savemode int

const (
	// Savemode_Mhtml is a representation of the C type WEBKIT_SAVE_MODE_MHTML.
	Savemode_Mhtml Savemode = 0
)

// Scriptdialogtype is a representation of the C type WebKitScriptDialogType.
type Scriptdialogtype int

const (
	// Scriptdialogtype_Alert is a representation of the C type WEBKIT_SCRIPT_DIALOG_ALERT.
	Scriptdialogtype_Alert Scriptdialogtype = 0
	// Scriptdialogtype_Confirm is a representation of the C type WEBKIT_SCRIPT_DIALOG_CONFIRM.
	Scriptdialogtype_Confirm Scriptdialogtype = 1
	// Scriptdialogtype_Prompt is a representation of the C type WEBKIT_SCRIPT_DIALOG_PROMPT.
	Scriptdialogtype_Prompt Scriptdialogtype = 2
	// Scriptdialogtype_BeforeUnloadConfirm is a representation of the C type WEBKIT_SCRIPT_DIALOG_BEFORE_UNLOAD_CONFIRM.
	Scriptdialogtype_BeforeUnloadConfirm Scriptdialogtype = 3
)

// Snapshoterror is a representation of the C type WebKitSnapshotError.
type Snapshoterror int

const (
	// Snapshoterror_Create is a representation of the C type WEBKIT_SNAPSHOT_ERROR_FAILED_TO_CREATE.
	Snapshoterror_Create Snapshoterror = 799
)

// Snapshotregion is a representation of the C type WebKitSnapshotRegion.
type Snapshotregion int

const (
	// Snapshotregion_Visible is a representation of the C type WEBKIT_SNAPSHOT_REGION_VISIBLE.
	Snapshotregion_Visible Snapshotregion = 0
	// Snapshotregion_FullDocument is a representation of the C type WEBKIT_SNAPSHOT_REGION_FULL_DOCUMENT.
	Snapshotregion_FullDocument Snapshotregion = 1
)

// Tlserrorspolicy is a representation of the C type WebKitTLSErrorsPolicy.
type Tlserrorspolicy int

const (
	// Tlserrorspolicy_Ignore is a representation of the C type WEBKIT_TLS_ERRORS_POLICY_IGNORE.
	Tlserrorspolicy_Ignore Tlserrorspolicy = 0
	// Tlserrorspolicy_Fail is a representation of the C type WEBKIT_TLS_ERRORS_POLICY_FAIL.
	Tlserrorspolicy_Fail Tlserrorspolicy = 1
)

// Usercontentfiltererror is a representation of the C type WebKitUserContentFilterError.
//
// since 2.24
type Usercontentfiltererror int

const (
	// Usercontentfiltererror_InvalidSource is a representation of the C type WEBKIT_USER_CONTENT_FILTER_ERROR_INVALID_SOURCE.
	Usercontentfiltererror_InvalidSource Usercontentfiltererror = 0
	// Usercontentfiltererror_NotFound is a representation of the C type WEBKIT_USER_CONTENT_FILTER_ERROR_NOT_FOUND.
	Usercontentfiltererror_NotFound Usercontentfiltererror = 1
)

// Usercontentinjectedframes is a representation of the C type WebKitUserContentInjectedFrames.
//
// since 2.6
type Usercontentinjectedframes int

const (
	// Usercontentinjectedframes_AllFrames is a representation of the C type WEBKIT_USER_CONTENT_INJECT_ALL_FRAMES.
	Usercontentinjectedframes_AllFrames Usercontentinjectedframes = 0
	// Usercontentinjectedframes_TopFrame is a representation of the C type WEBKIT_USER_CONTENT_INJECT_TOP_FRAME.
	Usercontentinjectedframes_TopFrame Usercontentinjectedframes = 1
)

// Userscriptinjectiontime is a representation of the C type WebKitUserScriptInjectionTime.
//
// since 2.6
type Userscriptinjectiontime int

const (
	// Userscriptinjectiontime_Start is a representation of the C type WEBKIT_USER_SCRIPT_INJECT_AT_DOCUMENT_START.
	Userscriptinjectiontime_Start Userscriptinjectiontime = 0
	// Userscriptinjectiontime_End is a representation of the C type WEBKIT_USER_SCRIPT_INJECT_AT_DOCUMENT_END.
	Userscriptinjectiontime_End Userscriptinjectiontime = 1
)

// Userstylelevel is a representation of the C type WebKitUserStyleLevel.
//
// since 2.6
type Userstylelevel int

const (
	// Userstylelevel_User is a representation of the C type WEBKIT_USER_STYLE_LEVEL_USER.
	Userstylelevel_User Userstylelevel = 0
	// Userstylelevel_Author is a representation of the C type WEBKIT_USER_STYLE_LEVEL_AUTHOR.
	Userstylelevel_Author Userstylelevel = 1
)

// Webprocessterminationreason is a representation of the C type WebKitWebProcessTerminationReason.
//
// since 2.20
type Webprocessterminationreason int

const (
	// Webprocessterminationreason_Crashed is a representation of the C type WEBKIT_WEB_PROCESS_CRASHED.
	Webprocessterminationreason_Crashed Webprocessterminationreason = 0
	// Webprocessterminationreason_ExceededMemoryLimit is a representation of the C type WEBKIT_WEB_PROCESS_EXCEEDED_MEMORY_LIMIT.
	Webprocessterminationreason_ExceededMemoryLimit Webprocessterminationreason = 1
)
