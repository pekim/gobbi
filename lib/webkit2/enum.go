// Code generated - DO NOT EDIT.

package webkit2

type AuthenticationScheme uint32

const (
	AuthenticationScheme_Default                        AuthenticationScheme = int64(1)
	AuthenticationScheme_HttpBasic                      AuthenticationScheme = int64(2)
	AuthenticationScheme_HttpDigest                     AuthenticationScheme = int64(3)
	AuthenticationScheme_HtmlForm                       AuthenticationScheme = int64(4)
	AuthenticationScheme_Ntlm                           AuthenticationScheme = int64(5)
	AuthenticationScheme_Negotiate                      AuthenticationScheme = int64(6)
	AuthenticationScheme_ClientCertificateRequested     AuthenticationScheme = int64(7)
	AuthenticationScheme_ServerTrustEvaluationRequested AuthenticationScheme = int64(8)
	AuthenticationScheme_Unknown                        AuthenticationScheme = int64(100)
)

type CacheModel uint32

const (
	CacheModel_DocumentViewer  CacheModel = int64(0)
	CacheModel_WebBrowser      CacheModel = int64(1)
	CacheModel_DocumentBrowser CacheModel = int64(2)
)

type ContextMenuAction uint32

const (
	ContextMenuAction_NoAction                 ContextMenuAction = int64(0)
	ContextMenuAction_OpenLink                 ContextMenuAction = int64(1)
	ContextMenuAction_OpenLinkInNewWindow      ContextMenuAction = int64(2)
	ContextMenuAction_DownloadLinkToDisk       ContextMenuAction = int64(3)
	ContextMenuAction_CopyLinkToClipboard      ContextMenuAction = int64(4)
	ContextMenuAction_OpenImageInNewWindow     ContextMenuAction = int64(5)
	ContextMenuAction_DownloadImageToDisk      ContextMenuAction = int64(6)
	ContextMenuAction_CopyImageToClipboard     ContextMenuAction = int64(7)
	ContextMenuAction_CopyImageUrlToClipboard  ContextMenuAction = int64(8)
	ContextMenuAction_OpenFrameInNewWindow     ContextMenuAction = int64(9)
	ContextMenuAction_GoBack                   ContextMenuAction = int64(10)
	ContextMenuAction_GoForward                ContextMenuAction = int64(11)
	ContextMenuAction_Stop                     ContextMenuAction = int64(12)
	ContextMenuAction_Reload                   ContextMenuAction = int64(13)
	ContextMenuAction_Copy                     ContextMenuAction = int64(14)
	ContextMenuAction_Cut                      ContextMenuAction = int64(15)
	ContextMenuAction_Paste                    ContextMenuAction = int64(16)
	ContextMenuAction_Delete                   ContextMenuAction = int64(17)
	ContextMenuAction_SelectAll                ContextMenuAction = int64(18)
	ContextMenuAction_InputMethods             ContextMenuAction = int64(19)
	ContextMenuAction_Unicode                  ContextMenuAction = int64(20)
	ContextMenuAction_SpellingGuess            ContextMenuAction = int64(21)
	ContextMenuAction_NoGuessesFound           ContextMenuAction = int64(22)
	ContextMenuAction_IgnoreSpelling           ContextMenuAction = int64(23)
	ContextMenuAction_LearnSpelling            ContextMenuAction = int64(24)
	ContextMenuAction_IgnoreGrammar            ContextMenuAction = int64(25)
	ContextMenuAction_FontMenu                 ContextMenuAction = int64(26)
	ContextMenuAction_Bold                     ContextMenuAction = int64(27)
	ContextMenuAction_Italic                   ContextMenuAction = int64(28)
	ContextMenuAction_Underline                ContextMenuAction = int64(29)
	ContextMenuAction_Outline                  ContextMenuAction = int64(30)
	ContextMenuAction_InspectElement           ContextMenuAction = int64(31)
	ContextMenuAction_OpenVideoInNewWindow     ContextMenuAction = int64(32)
	ContextMenuAction_OpenAudioInNewWindow     ContextMenuAction = int64(33)
	ContextMenuAction_CopyVideoLinkToClipboard ContextMenuAction = int64(34)
	ContextMenuAction_CopyAudioLinkToClipboard ContextMenuAction = int64(35)
	ContextMenuAction_ToggleMediaControls      ContextMenuAction = int64(36)
	ContextMenuAction_ToggleMediaLoop          ContextMenuAction = int64(37)
	ContextMenuAction_EnterVideoFullscreen     ContextMenuAction = int64(38)
	ContextMenuAction_MediaPlay                ContextMenuAction = int64(39)
	ContextMenuAction_MediaPause               ContextMenuAction = int64(40)
	ContextMenuAction_MediaMute                ContextMenuAction = int64(41)
	ContextMenuAction_DownloadVideoToDisk      ContextMenuAction = int64(42)
	ContextMenuAction_DownloadAudioToDisk      ContextMenuAction = int64(43)
	ContextMenuAction_InsertEmoji              ContextMenuAction = int64(44)
	ContextMenuAction_Custom                   ContextMenuAction = int64(10000)
)

type CookieAcceptPolicy uint32

const (
	CookieAcceptPolicy_Always       CookieAcceptPolicy = int64(0)
	CookieAcceptPolicy_Never        CookieAcceptPolicy = int64(1)
	CookieAcceptPolicy_NoThirdParty CookieAcceptPolicy = int64(2)
)

type CookiePersistentStorage uint32

const (
	CookiePersistentStorage_Text   CookiePersistentStorage = int64(0)
	CookiePersistentStorage_Sqlite CookiePersistentStorage = int64(1)
)

type CredentialPersistence uint32

const (
	CredentialPersistence_None       CredentialPersistence = int64(0)
	CredentialPersistence_ForSession CredentialPersistence = int64(1)
	CredentialPersistence_Permanent  CredentialPersistence = int64(2)
)

type DownloadError uint32

const (
	DownloadError_Network         DownloadError = int64(499)
	DownloadError_CancelledByUser DownloadError = int64(400)
	DownloadError_Destination     DownloadError = int64(401)
)

type FaviconDatabaseError uint32

const (
	FaviconDatabaseError_NotInitialized  FaviconDatabaseError = int64(0)
	FaviconDatabaseError_FaviconNotFound FaviconDatabaseError = int64(1)
	FaviconDatabaseError_FaviconUnknown  FaviconDatabaseError = int64(2)
)

type HardwareAccelerationPolicy uint32

const (
	HardwareAccelerationPolicy_OnDemand HardwareAccelerationPolicy = int64(0)
	HardwareAccelerationPolicy_Always   HardwareAccelerationPolicy = int64(1)
	HardwareAccelerationPolicy_Never    HardwareAccelerationPolicy = int64(2)
)

type InsecureContentEvent uint32

const (
	InsecureContentEvent_Run       InsecureContentEvent = int64(0)
	InsecureContentEvent_Displayed InsecureContentEvent = int64(1)
)

type JavascriptError uint32

const (
	JavascriptError_Failed JavascriptError = int64(699)
)

type LoadEvent uint32

const (
	LoadEvent_Started    LoadEvent = int64(0)
	LoadEvent_Redirected LoadEvent = int64(1)
	LoadEvent_Committed  LoadEvent = int64(2)
	LoadEvent_Finished   LoadEvent = int64(3)
)

type NavigationType uint32

const (
	NavigationType_LinkClicked     NavigationType = int64(0)
	NavigationType_FormSubmitted   NavigationType = int64(1)
	NavigationType_BackForward     NavigationType = int64(2)
	NavigationType_Reload          NavigationType = int64(3)
	NavigationType_FormResubmitted NavigationType = int64(4)
	NavigationType_Other           NavigationType = int64(5)
)

type NetworkError uint32

const (
	NetworkError_Failed           NetworkError = int64(399)
	NetworkError_Transport        NetworkError = int64(300)
	NetworkError_UnknownProtocol  NetworkError = int64(301)
	NetworkError_Cancelled        NetworkError = int64(302)
	NetworkError_FileDoesNotExist NetworkError = int64(303)
)

type NetworkProxyMode uint32

const (
	NetworkProxyMode_Default NetworkProxyMode = int64(0)
	NetworkProxyMode_NoProxy NetworkProxyMode = int64(1)
	NetworkProxyMode_Custom  NetworkProxyMode = int64(2)
)

type PluginError uint32

const (
	PluginError_Failed              PluginError = int64(299)
	PluginError_CannotFindPlugin    PluginError = int64(200)
	PluginError_CannotLoadPlugin    PluginError = int64(201)
	PluginError_JavaUnavailable     PluginError = int64(202)
	PluginError_ConnectionCancelled PluginError = int64(203)
	PluginError_WillHandleLoad      PluginError = int64(204)
)

type PolicyDecisionType uint32

const (
	PolicyDecisionType_NavigationAction PolicyDecisionType = int64(0)
	PolicyDecisionType_NewWindowAction  PolicyDecisionType = int64(1)
	PolicyDecisionType_Response         PolicyDecisionType = int64(2)
)

type PolicyError uint32

const (
	PolicyError_Failed                             PolicyError = int64(199)
	PolicyError_CannotShowMimeType                 PolicyError = int64(100)
	PolicyError_CannotShowUri                      PolicyError = int64(101)
	PolicyError_FrameLoadInterruptedByPolicyChange PolicyError = int64(102)
	PolicyError_CannotUseRestrictedPort            PolicyError = int64(103)
)

type PrintError uint32

const (
	PrintError_General          PrintError = int64(599)
	PrintError_PrinterNotFound  PrintError = int64(500)
	PrintError_InvalidPageRange PrintError = int64(501)
)

type PrintOperationResponse uint32

const (
	PrintOperationResponse_Print  PrintOperationResponse = int64(0)
	PrintOperationResponse_Cancel PrintOperationResponse = int64(1)
)

type ProcessModel uint32

const (
	ProcessModel_SharedSecondaryProcess     ProcessModel = int64(0)
	ProcessModel_MultipleSecondaryProcesses ProcessModel = int64(1)
)

type SaveMode uint32

const (
	SaveMode_Mhtml SaveMode = int64(0)
)

type ScriptDialogType uint32

const (
	ScriptDialogType_Alert               ScriptDialogType = int64(0)
	ScriptDialogType_Confirm             ScriptDialogType = int64(1)
	ScriptDialogType_Prompt              ScriptDialogType = int64(2)
	ScriptDialogType_BeforeUnloadConfirm ScriptDialogType = int64(3)
)

type SnapshotError uint32

const (
	SnapshotError_Create SnapshotError = int64(799)
)

type SnapshotRegion uint32

const (
	SnapshotRegion_Visible      SnapshotRegion = int64(0)
	SnapshotRegion_FullDocument SnapshotRegion = int64(1)
)

type TLSErrorsPolicy uint32

const (
	TLSErrorsPolicy_Ignore TLSErrorsPolicy = int64(0)
	TLSErrorsPolicy_Fail   TLSErrorsPolicy = int64(1)
)

type UserContentFilterError uint32

const (
	UserContentFilterError_InvalidSource UserContentFilterError = int64(0)
	UserContentFilterError_NotFound      UserContentFilterError = int64(1)
)

type UserContentInjectedFrames uint32

const (
	UserContentInjectedFrames_AllFrames UserContentInjectedFrames = int64(0)
	UserContentInjectedFrames_TopFrame  UserContentInjectedFrames = int64(1)
)

type UserScriptInjectionTime uint32

const (
	UserScriptInjectionTime_Start UserScriptInjectionTime = int64(0)
	UserScriptInjectionTime_End   UserScriptInjectionTime = int64(1)
)

type UserStyleLevel uint32

const (
	UserStyleLevel_User   UserStyleLevel = int64(0)
	UserStyleLevel_Author UserStyleLevel = int64(1)
)

type WebProcessTerminationReason uint32

const (
	WebProcessTerminationReason_Crashed             WebProcessTerminationReason = int64(0)
	WebProcessTerminationReason_ExceededMemoryLimit WebProcessTerminationReason = int64(1)
)
