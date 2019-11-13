// Code generated - DO NOT EDIT.

package webkit2

// Authenticationscheme is a representation of the C type AuthenticationScheme.
//
// since 2.2
type Authenticationscheme int

const (
	// default
	WebkitAuthenticationSchemeDefault Authenticationscheme = 1
	// http_basic
	WebkitAuthenticationSchemeHttpBasic Authenticationscheme = 2
	// http_digest
	WebkitAuthenticationSchemeHttpDigest Authenticationscheme = 3
	// html_form
	WebkitAuthenticationSchemeHtmlForm Authenticationscheme = 4
	// ntlm
	WebkitAuthenticationSchemeNtlm Authenticationscheme = 5
	// negotiate
	WebkitAuthenticationSchemeNegotiate Authenticationscheme = 6
	// client_certificate_requested
	WebkitAuthenticationSchemeClientCertificateRequested Authenticationscheme = 7
	// server_trust_evaluation_requested
	WebkitAuthenticationSchemeServerTrustEvaluationRequested Authenticationscheme = 8
	// unknown
	WebkitAuthenticationSchemeUnknown Authenticationscheme = 100
)

// Cachemodel is a representation of the C type CacheModel.
type Cachemodel int

const (
	// document_viewer
	WebkitCacheModelDocumentViewer Cachemodel = 0
	// web_browser
	WebkitCacheModelWebBrowser Cachemodel = 1
	// document_browser
	WebkitCacheModelDocumentBrowser Cachemodel = 2
)

// Contextmenuaction is a representation of the C type ContextMenuAction.
type Contextmenuaction int

const (
	// no_action
	WebkitContextMenuActionNoAction Contextmenuaction = 0
	// open_link
	WebkitContextMenuActionOpenLink Contextmenuaction = 1
	// open_link_in_new_window
	WebkitContextMenuActionOpenLinkInNewWindow Contextmenuaction = 2
	// download_link_to_disk
	WebkitContextMenuActionDownloadLinkToDisk Contextmenuaction = 3
	// copy_link_to_clipboard
	WebkitContextMenuActionCopyLinkToClipboard Contextmenuaction = 4
	// open_image_in_new_window
	WebkitContextMenuActionOpenImageInNewWindow Contextmenuaction = 5
	// download_image_to_disk
	WebkitContextMenuActionDownloadImageToDisk Contextmenuaction = 6
	// copy_image_to_clipboard
	WebkitContextMenuActionCopyImageToClipboard Contextmenuaction = 7
	// copy_image_url_to_clipboard
	WebkitContextMenuActionCopyImageUrlToClipboard Contextmenuaction = 8
	// open_frame_in_new_window
	WebkitContextMenuActionOpenFrameInNewWindow Contextmenuaction = 9
	// go_back
	WebkitContextMenuActionGoBack Contextmenuaction = 10
	// go_forward
	WebkitContextMenuActionGoForward Contextmenuaction = 11
	// stop
	WebkitContextMenuActionStop Contextmenuaction = 12
	// reload
	WebkitContextMenuActionReload Contextmenuaction = 13
	// copy
	WebkitContextMenuActionCopy Contextmenuaction = 14
	// cut
	WebkitContextMenuActionCut Contextmenuaction = 15
	// paste
	WebkitContextMenuActionPaste Contextmenuaction = 16
	// delete
	WebkitContextMenuActionDelete Contextmenuaction = 17
	// select_all
	WebkitContextMenuActionSelectAll Contextmenuaction = 18
	// input_methods
	WebkitContextMenuActionInputMethods Contextmenuaction = 19
	// unicode
	WebkitContextMenuActionUnicode Contextmenuaction = 20
	// spelling_guess
	WebkitContextMenuActionSpellingGuess Contextmenuaction = 21
	// no_guesses_found
	WebkitContextMenuActionNoGuessesFound Contextmenuaction = 22
	// ignore_spelling
	WebkitContextMenuActionIgnoreSpelling Contextmenuaction = 23
	// learn_spelling
	WebkitContextMenuActionLearnSpelling Contextmenuaction = 24
	// ignore_grammar
	WebkitContextMenuActionIgnoreGrammar Contextmenuaction = 25
	// font_menu
	WebkitContextMenuActionFontMenu Contextmenuaction = 26
	// bold
	WebkitContextMenuActionBold Contextmenuaction = 27
	// italic
	WebkitContextMenuActionItalic Contextmenuaction = 28
	// underline
	WebkitContextMenuActionUnderline Contextmenuaction = 29
	// outline
	WebkitContextMenuActionOutline Contextmenuaction = 30
	// inspect_element
	WebkitContextMenuActionInspectElement Contextmenuaction = 31
	// open_video_in_new_window
	WebkitContextMenuActionOpenVideoInNewWindow Contextmenuaction = 32
	// open_audio_in_new_window
	WebkitContextMenuActionOpenAudioInNewWindow Contextmenuaction = 33
	// copy_video_link_to_clipboard
	WebkitContextMenuActionCopyVideoLinkToClipboard Contextmenuaction = 34
	// copy_audio_link_to_clipboard
	WebkitContextMenuActionCopyAudioLinkToClipboard Contextmenuaction = 35
	// toggle_media_controls
	WebkitContextMenuActionToggleMediaControls Contextmenuaction = 36
	// toggle_media_loop
	WebkitContextMenuActionToggleMediaLoop Contextmenuaction = 37
	// enter_video_fullscreen
	WebkitContextMenuActionEnterVideoFullscreen Contextmenuaction = 38
	// media_play
	WebkitContextMenuActionMediaPlay Contextmenuaction = 39
	// media_pause
	WebkitContextMenuActionMediaPause Contextmenuaction = 40
	// media_mute
	WebkitContextMenuActionMediaMute Contextmenuaction = 41
	// download_video_to_disk
	WebkitContextMenuActionDownloadVideoToDisk Contextmenuaction = 42
	// download_audio_to_disk
	WebkitContextMenuActionDownloadAudioToDisk Contextmenuaction = 43
	// insert_emoji
	WebkitContextMenuActionInsertEmoji Contextmenuaction = 44
	// custom
	WebkitContextMenuActionCustom Contextmenuaction = 10000
)

// Cookieacceptpolicy is a representation of the C type CookieAcceptPolicy.
type Cookieacceptpolicy int

const (
	// always
	WebkitCookiePolicyAcceptAlways Cookieacceptpolicy = 0
	// never
	WebkitCookiePolicyAcceptNever Cookieacceptpolicy = 1
	// no_third_party
	WebkitCookiePolicyAcceptNoThirdParty Cookieacceptpolicy = 2
)

// Cookiepersistentstorage is a representation of the C type CookiePersistentStorage.
type Cookiepersistentstorage int

const (
	// text
	WebkitCookiePersistentStorageText Cookiepersistentstorage = 0
	// sqlite
	WebkitCookiePersistentStorageSqlite Cookiepersistentstorage = 1
)

// Credentialpersistence is a representation of the C type CredentialPersistence.
//
// since 2.2
type Credentialpersistence int

const (
	// none
	WebkitCredentialPersistenceNone Credentialpersistence = 0
	// for_session
	WebkitCredentialPersistenceForSession Credentialpersistence = 1
	// permanent
	WebkitCredentialPersistencePermanent Credentialpersistence = 2
)

// Downloaderror is a representation of the C type DownloadError.
type Downloaderror int

const (
	// network
	WebkitDownloadErrorNetwork Downloaderror = 499
	// cancelled_by_user
	WebkitDownloadErrorCancelledByUser Downloaderror = 400
	// destination
	WebkitDownloadErrorDestination Downloaderror = 401
)

// Favicondatabaseerror is a representation of the C type FaviconDatabaseError.
type Favicondatabaseerror int

const (
	// not_initialized
	WebkitFaviconDatabaseErrorNotInitialized Favicondatabaseerror = 0
	// favicon_not_found
	WebkitFaviconDatabaseErrorFaviconNotFound Favicondatabaseerror = 1
	// favicon_unknown
	WebkitFaviconDatabaseErrorFaviconUnknown Favicondatabaseerror = 2
)

// Hardwareaccelerationpolicy is a representation of the C type HardwareAccelerationPolicy.
//
// since 2.16
type Hardwareaccelerationpolicy int

const (
	// on_demand
	WebkitHardwareAccelerationPolicyOnDemand Hardwareaccelerationpolicy = 0
	// always
	WebkitHardwareAccelerationPolicyAlways Hardwareaccelerationpolicy = 1
	// never
	WebkitHardwareAccelerationPolicyNever Hardwareaccelerationpolicy = 2
)

// Insecurecontentevent is a representation of the C type InsecureContentEvent.
type Insecurecontentevent int

const (
	// run
	WebkitInsecureContentRun Insecurecontentevent = 0
	// displayed
	WebkitInsecureContentDisplayed Insecurecontentevent = 1
)

// Javascripterror is a representation of the C type JavascriptError.
type Javascripterror int

const (
	// failed
	WebkitJavascriptErrorScriptFailed Javascripterror = 699
)

// Loadevent is a representation of the C type LoadEvent.
type Loadevent int

const (
	// started
	WebkitLoadStarted Loadevent = 0
	// redirected
	WebkitLoadRedirected Loadevent = 1
	// committed
	WebkitLoadCommitted Loadevent = 2
	// finished
	WebkitLoadFinished Loadevent = 3
)

// Navigationtype is a representation of the C type NavigationType.
type Navigationtype int

const (
	// link_clicked
	WebkitNavigationTypeLinkClicked Navigationtype = 0
	// form_submitted
	WebkitNavigationTypeFormSubmitted Navigationtype = 1
	// back_forward
	WebkitNavigationTypeBackForward Navigationtype = 2
	// reload
	WebkitNavigationTypeReload Navigationtype = 3
	// form_resubmitted
	WebkitNavigationTypeFormResubmitted Navigationtype = 4
	// other
	WebkitNavigationTypeOther Navigationtype = 5
)

// Networkerror is a representation of the C type NetworkError.
type Networkerror int

const (
	// failed
	WebkitNetworkErrorFailed Networkerror = 399
	// transport
	WebkitNetworkErrorTransport Networkerror = 300
	// unknown_protocol
	WebkitNetworkErrorUnknownProtocol Networkerror = 301
	// cancelled
	WebkitNetworkErrorCancelled Networkerror = 302
	// file_does_not_exist
	WebkitNetworkErrorFileDoesNotExist Networkerror = 303
)

// Networkproxymode is a representation of the C type NetworkProxyMode.
//
// since 2.16
type Networkproxymode int

const (
	// default
	WebkitNetworkProxyModeDefault Networkproxymode = 0
	// no_proxy
	WebkitNetworkProxyModeNoProxy Networkproxymode = 1
	// custom
	WebkitNetworkProxyModeCustom Networkproxymode = 2
)

// Pluginerror is a representation of the C type PluginError.
type Pluginerror int

const (
	// failed
	WebkitPluginErrorFailed Pluginerror = 299
	// cannot_find_plugin
	WebkitPluginErrorCannotFindPlugin Pluginerror = 200
	// cannot_load_plugin
	WebkitPluginErrorCannotLoadPlugin Pluginerror = 201
	// java_unavailable
	WebkitPluginErrorJavaUnavailable Pluginerror = 202
	// connection_cancelled
	WebkitPluginErrorConnectionCancelled Pluginerror = 203
	// will_handle_load
	WebkitPluginErrorWillHandleLoad Pluginerror = 204
)

// Policydecisiontype is a representation of the C type PolicyDecisionType.
type Policydecisiontype int

const (
	// navigation_action
	WebkitPolicyDecisionTypeNavigationAction Policydecisiontype = 0
	// new_window_action
	WebkitPolicyDecisionTypeNewWindowAction Policydecisiontype = 1
	// response
	WebkitPolicyDecisionTypeResponse Policydecisiontype = 2
)

// Policyerror is a representation of the C type PolicyError.
type Policyerror int

const (
	// failed
	WebkitPolicyErrorFailed Policyerror = 199
	// cannot_show_mime_type
	WebkitPolicyErrorCannotShowMimeType Policyerror = 100
	// cannot_show_uri
	WebkitPolicyErrorCannotShowUri Policyerror = 101
	// frame_load_interrupted_by_policy_change
	WebkitPolicyErrorFrameLoadInterruptedByPolicyChange Policyerror = 102
	// cannot_use_restricted_port
	WebkitPolicyErrorCannotUseRestrictedPort Policyerror = 103
)

// Printerror is a representation of the C type PrintError.
type Printerror int

const (
	// general
	WebkitPrintErrorGeneral Printerror = 599
	// printer_not_found
	WebkitPrintErrorPrinterNotFound Printerror = 500
	// invalid_page_range
	WebkitPrintErrorInvalidPageRange Printerror = 501
)

// Printoperationresponse is a representation of the C type PrintOperationResponse.
type Printoperationresponse int

const (
	// print
	WebkitPrintOperationResponsePrint Printoperationresponse = 0
	// cancel
	WebkitPrintOperationResponseCancel Printoperationresponse = 1
)

// Processmodel is a representation of the C type ProcessModel.
//
// since 2.4
type Processmodel int

const (
	// shared_secondary_process
	WebkitProcessModelSharedSecondaryProcess Processmodel = 0
	// multiple_secondary_processes
	WebkitProcessModelMultipleSecondaryProcesses Processmodel = 1
)

// Savemode is a representation of the C type SaveMode.
type Savemode int

const (
	// mhtml
	WebkitSaveModeMhtml Savemode = 0
)

// Scriptdialogtype is a representation of the C type ScriptDialogType.
type Scriptdialogtype int

const (
	// alert
	WebkitScriptDialogAlert Scriptdialogtype = 0
	// confirm
	WebkitScriptDialogConfirm Scriptdialogtype = 1
	// prompt
	WebkitScriptDialogPrompt Scriptdialogtype = 2
	// before_unload_confirm
	WebkitScriptDialogBeforeUnloadConfirm Scriptdialogtype = 3
)

// Snapshoterror is a representation of the C type SnapshotError.
type Snapshoterror int

const (
	// create
	WebkitSnapshotErrorFailedToCreate Snapshoterror = 799
)

// Snapshotregion is a representation of the C type SnapshotRegion.
type Snapshotregion int

const (
	// visible
	WebkitSnapshotRegionVisible Snapshotregion = 0
	// full_document
	WebkitSnapshotRegionFullDocument Snapshotregion = 1
)

// Tlserrorspolicy is a representation of the C type TLSErrorsPolicy.
type Tlserrorspolicy int

const (
	// ignore
	WebkitTlsErrorsPolicyIgnore Tlserrorspolicy = 0
	// fail
	WebkitTlsErrorsPolicyFail Tlserrorspolicy = 1
)

// Usercontentfiltererror is a representation of the C type UserContentFilterError.
//
// since 2.24
type Usercontentfiltererror int

const (
	// invalid_source
	WebkitUserContentFilterErrorInvalidSource Usercontentfiltererror = 0
	// not_found
	WebkitUserContentFilterErrorNotFound Usercontentfiltererror = 1
)

// Usercontentinjectedframes is a representation of the C type UserContentInjectedFrames.
//
// since 2.6
type Usercontentinjectedframes int

const (
	// all_frames
	WebkitUserContentInjectAllFrames Usercontentinjectedframes = 0
	// top_frame
	WebkitUserContentInjectTopFrame Usercontentinjectedframes = 1
)

// Userscriptinjectiontime is a representation of the C type UserScriptInjectionTime.
//
// since 2.6
type Userscriptinjectiontime int

const (
	// start
	WebkitUserScriptInjectAtDocumentStart Userscriptinjectiontime = 0
	// end
	WebkitUserScriptInjectAtDocumentEnd Userscriptinjectiontime = 1
)

// Userstylelevel is a representation of the C type UserStyleLevel.
//
// since 2.6
type Userstylelevel int

const (
	// user
	WebkitUserStyleLevelUser Userstylelevel = 0
	// author
	WebkitUserStyleLevelAuthor Userstylelevel = 1
)

// Webprocessterminationreason is a representation of the C type WebProcessTerminationReason.
//
// since 2.20
type Webprocessterminationreason int

const (
	// crashed
	WebkitWebProcessCrashed Webprocessterminationreason = 0
	// exceeded_memory_limit
	WebkitWebProcessExceededMemoryLimit Webprocessterminationreason = 1
)
