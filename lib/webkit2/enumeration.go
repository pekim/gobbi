// Code generated - DO NOT EDIT.

package webkit2

// Authenticationscheme is a representation of the C type AuthenticationScheme.
//
// since 2.2
type Authenticationscheme int

const (
	Authenticationscheme_Default                        Authenticationscheme = 1
	Authenticationscheme_HttpBasic                      Authenticationscheme = 2
	Authenticationscheme_HttpDigest                     Authenticationscheme = 3
	Authenticationscheme_HtmlForm                       Authenticationscheme = 4
	Authenticationscheme_Ntlm                           Authenticationscheme = 5
	Authenticationscheme_Negotiate                      Authenticationscheme = 6
	Authenticationscheme_ClientCertificateRequested     Authenticationscheme = 7
	Authenticationscheme_ServerTrustEvaluationRequested Authenticationscheme = 8
	Authenticationscheme_Unknown                        Authenticationscheme = 100
)

// Cachemodel is a representation of the C type CacheModel.
type Cachemodel int

const (
	Cachemodel_DocumentViewer  Cachemodel = 0
	Cachemodel_WebBrowser      Cachemodel = 1
	Cachemodel_DocumentBrowser Cachemodel = 2
)

// Contextmenuaction is a representation of the C type ContextMenuAction.
type Contextmenuaction int

const (
	Contextmenuaction_NoAction                 Contextmenuaction = 0
	Contextmenuaction_OpenLink                 Contextmenuaction = 1
	Contextmenuaction_OpenLinkInNewWindow      Contextmenuaction = 2
	Contextmenuaction_DownloadLinkToDisk       Contextmenuaction = 3
	Contextmenuaction_CopyLinkToClipboard      Contextmenuaction = 4
	Contextmenuaction_OpenImageInNewWindow     Contextmenuaction = 5
	Contextmenuaction_DownloadImageToDisk      Contextmenuaction = 6
	Contextmenuaction_CopyImageToClipboard     Contextmenuaction = 7
	Contextmenuaction_CopyImageUrlToClipboard  Contextmenuaction = 8
	Contextmenuaction_OpenFrameInNewWindow     Contextmenuaction = 9
	Contextmenuaction_GoBack                   Contextmenuaction = 10
	Contextmenuaction_GoForward                Contextmenuaction = 11
	Contextmenuaction_Stop                     Contextmenuaction = 12
	Contextmenuaction_Reload                   Contextmenuaction = 13
	Contextmenuaction_Copy                     Contextmenuaction = 14
	Contextmenuaction_Cut                      Contextmenuaction = 15
	Contextmenuaction_Paste                    Contextmenuaction = 16
	Contextmenuaction_Delete                   Contextmenuaction = 17
	Contextmenuaction_SelectAll                Contextmenuaction = 18
	Contextmenuaction_InputMethods             Contextmenuaction = 19
	Contextmenuaction_Unicode                  Contextmenuaction = 20
	Contextmenuaction_SpellingGuess            Contextmenuaction = 21
	Contextmenuaction_NoGuessesFound           Contextmenuaction = 22
	Contextmenuaction_IgnoreSpelling           Contextmenuaction = 23
	Contextmenuaction_LearnSpelling            Contextmenuaction = 24
	Contextmenuaction_IgnoreGrammar            Contextmenuaction = 25
	Contextmenuaction_FontMenu                 Contextmenuaction = 26
	Contextmenuaction_Bold                     Contextmenuaction = 27
	Contextmenuaction_Italic                   Contextmenuaction = 28
	Contextmenuaction_Underline                Contextmenuaction = 29
	Contextmenuaction_Outline                  Contextmenuaction = 30
	Contextmenuaction_InspectElement           Contextmenuaction = 31
	Contextmenuaction_OpenVideoInNewWindow     Contextmenuaction = 32
	Contextmenuaction_OpenAudioInNewWindow     Contextmenuaction = 33
	Contextmenuaction_CopyVideoLinkToClipboard Contextmenuaction = 34
	Contextmenuaction_CopyAudioLinkToClipboard Contextmenuaction = 35
	Contextmenuaction_ToggleMediaControls      Contextmenuaction = 36
	Contextmenuaction_ToggleMediaLoop          Contextmenuaction = 37
	Contextmenuaction_EnterVideoFullscreen     Contextmenuaction = 38
	Contextmenuaction_MediaPlay                Contextmenuaction = 39
	Contextmenuaction_MediaPause               Contextmenuaction = 40
	Contextmenuaction_MediaMute                Contextmenuaction = 41
	Contextmenuaction_DownloadVideoToDisk      Contextmenuaction = 42
	Contextmenuaction_DownloadAudioToDisk      Contextmenuaction = 43
	Contextmenuaction_InsertEmoji              Contextmenuaction = 44
	Contextmenuaction_Custom                   Contextmenuaction = 10000
)

// Cookieacceptpolicy is a representation of the C type CookieAcceptPolicy.
type Cookieacceptpolicy int

const (
	Cookieacceptpolicy_Always       Cookieacceptpolicy = 0
	Cookieacceptpolicy_Never        Cookieacceptpolicy = 1
	Cookieacceptpolicy_NoThirdParty Cookieacceptpolicy = 2
)

// Cookiepersistentstorage is a representation of the C type CookiePersistentStorage.
type Cookiepersistentstorage int

const (
	Cookiepersistentstorage_Text   Cookiepersistentstorage = 0
	Cookiepersistentstorage_Sqlite Cookiepersistentstorage = 1
)

// Credentialpersistence is a representation of the C type CredentialPersistence.
//
// since 2.2
type Credentialpersistence int

const (
	Credentialpersistence_None       Credentialpersistence = 0
	Credentialpersistence_ForSession Credentialpersistence = 1
	Credentialpersistence_Permanent  Credentialpersistence = 2
)

// Downloaderror is a representation of the C type DownloadError.
type Downloaderror int

const (
	Downloaderror_Network         Downloaderror = 499
	Downloaderror_CancelledByUser Downloaderror = 400
	Downloaderror_Destination     Downloaderror = 401
)

// Favicondatabaseerror is a representation of the C type FaviconDatabaseError.
type Favicondatabaseerror int

const (
	Favicondatabaseerror_NotInitialized  Favicondatabaseerror = 0
	Favicondatabaseerror_FaviconNotFound Favicondatabaseerror = 1
	Favicondatabaseerror_FaviconUnknown  Favicondatabaseerror = 2
)

// Hardwareaccelerationpolicy is a representation of the C type HardwareAccelerationPolicy.
//
// since 2.16
type Hardwareaccelerationpolicy int

const (
	Hardwareaccelerationpolicy_OnDemand Hardwareaccelerationpolicy = 0
	Hardwareaccelerationpolicy_Always   Hardwareaccelerationpolicy = 1
	Hardwareaccelerationpolicy_Never    Hardwareaccelerationpolicy = 2
)

// Insecurecontentevent is a representation of the C type InsecureContentEvent.
type Insecurecontentevent int

const (
	Insecurecontentevent_Run       Insecurecontentevent = 0
	Insecurecontentevent_Displayed Insecurecontentevent = 1
)

// Javascripterror is a representation of the C type JavascriptError.
type Javascripterror int

const (
	Javascripterror_Failed Javascripterror = 699
)

// Loadevent is a representation of the C type LoadEvent.
type Loadevent int

const (
	Loadevent_Started    Loadevent = 0
	Loadevent_Redirected Loadevent = 1
	Loadevent_Committed  Loadevent = 2
	Loadevent_Finished   Loadevent = 3
)

// Navigationtype is a representation of the C type NavigationType.
type Navigationtype int

const (
	Navigationtype_LinkClicked     Navigationtype = 0
	Navigationtype_FormSubmitted   Navigationtype = 1
	Navigationtype_BackForward     Navigationtype = 2
	Navigationtype_Reload          Navigationtype = 3
	Navigationtype_FormResubmitted Navigationtype = 4
	Navigationtype_Other           Navigationtype = 5
)

// Networkerror is a representation of the C type NetworkError.
type Networkerror int

const (
	Networkerror_Failed           Networkerror = 399
	Networkerror_Transport        Networkerror = 300
	Networkerror_UnknownProtocol  Networkerror = 301
	Networkerror_Cancelled        Networkerror = 302
	Networkerror_FileDoesNotExist Networkerror = 303
)

// Networkproxymode is a representation of the C type NetworkProxyMode.
//
// since 2.16
type Networkproxymode int

const (
	Networkproxymode_Default Networkproxymode = 0
	Networkproxymode_NoProxy Networkproxymode = 1
	Networkproxymode_Custom  Networkproxymode = 2
)

// Pluginerror is a representation of the C type PluginError.
type Pluginerror int

const (
	Pluginerror_Failed              Pluginerror = 299
	Pluginerror_CannotFindPlugin    Pluginerror = 200
	Pluginerror_CannotLoadPlugin    Pluginerror = 201
	Pluginerror_JavaUnavailable     Pluginerror = 202
	Pluginerror_ConnectionCancelled Pluginerror = 203
	Pluginerror_WillHandleLoad      Pluginerror = 204
)

// Policydecisiontype is a representation of the C type PolicyDecisionType.
type Policydecisiontype int

const (
	Policydecisiontype_NavigationAction Policydecisiontype = 0
	Policydecisiontype_NewWindowAction  Policydecisiontype = 1
	Policydecisiontype_Response         Policydecisiontype = 2
)

// Policyerror is a representation of the C type PolicyError.
type Policyerror int

const (
	Policyerror_Failed                             Policyerror = 199
	Policyerror_CannotShowMimeType                 Policyerror = 100
	Policyerror_CannotShowUri                      Policyerror = 101
	Policyerror_FrameLoadInterruptedByPolicyChange Policyerror = 102
	Policyerror_CannotUseRestrictedPort            Policyerror = 103
)

// Printerror is a representation of the C type PrintError.
type Printerror int

const (
	Printerror_General          Printerror = 599
	Printerror_PrinterNotFound  Printerror = 500
	Printerror_InvalidPageRange Printerror = 501
)

// Printoperationresponse is a representation of the C type PrintOperationResponse.
type Printoperationresponse int

const (
	Printoperationresponse_Print  Printoperationresponse = 0
	Printoperationresponse_Cancel Printoperationresponse = 1
)

// Processmodel is a representation of the C type ProcessModel.
//
// since 2.4
type Processmodel int

const (
	Processmodel_SharedSecondaryProcess     Processmodel = 0
	Processmodel_MultipleSecondaryProcesses Processmodel = 1
)

// Savemode is a representation of the C type SaveMode.
type Savemode int

const (
	Savemode_Mhtml Savemode = 0
)

// Scriptdialogtype is a representation of the C type ScriptDialogType.
type Scriptdialogtype int

const (
	Scriptdialogtype_Alert               Scriptdialogtype = 0
	Scriptdialogtype_Confirm             Scriptdialogtype = 1
	Scriptdialogtype_Prompt              Scriptdialogtype = 2
	Scriptdialogtype_BeforeUnloadConfirm Scriptdialogtype = 3
)

// Snapshoterror is a representation of the C type SnapshotError.
type Snapshoterror int

const (
	Snapshoterror_Create Snapshoterror = 799
)

// Snapshotregion is a representation of the C type SnapshotRegion.
type Snapshotregion int

const (
	Snapshotregion_Visible      Snapshotregion = 0
	Snapshotregion_FullDocument Snapshotregion = 1
)

// Tlserrorspolicy is a representation of the C type TLSErrorsPolicy.
type Tlserrorspolicy int

const (
	Tlserrorspolicy_Ignore Tlserrorspolicy = 0
	Tlserrorspolicy_Fail   Tlserrorspolicy = 1
)

// Usercontentfiltererror is a representation of the C type UserContentFilterError.
//
// since 2.24
type Usercontentfiltererror int

const (
	Usercontentfiltererror_InvalidSource Usercontentfiltererror = 0
	Usercontentfiltererror_NotFound      Usercontentfiltererror = 1
)

// Usercontentinjectedframes is a representation of the C type UserContentInjectedFrames.
//
// since 2.6
type Usercontentinjectedframes int

const (
	Usercontentinjectedframes_AllFrames Usercontentinjectedframes = 0
	Usercontentinjectedframes_TopFrame  Usercontentinjectedframes = 1
)

// Userscriptinjectiontime is a representation of the C type UserScriptInjectionTime.
//
// since 2.6
type Userscriptinjectiontime int

const (
	Userscriptinjectiontime_Start Userscriptinjectiontime = 0
	Userscriptinjectiontime_End   Userscriptinjectiontime = 1
)

// Userstylelevel is a representation of the C type UserStyleLevel.
//
// since 2.6
type Userstylelevel int

const (
	Userstylelevel_User   Userstylelevel = 0
	Userstylelevel_Author Userstylelevel = 1
)

// Webprocessterminationreason is a representation of the C type WebProcessTerminationReason.
//
// since 2.20
type Webprocessterminationreason int

const (
	Webprocessterminationreason_Crashed             Webprocessterminationreason = 0
	Webprocessterminationreason_ExceededMemoryLimit Webprocessterminationreason = 1
)
