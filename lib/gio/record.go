// This is a generated file - DO NOT EDIT

package gio

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GActionEntry

// Blacklisted : GAppInfoIface

// Blacklisted : GAppLaunchContextClass

// Blacklisted : GAppLaunchContextPrivate

// Blacklisted : GApplicationCommandLinePrivate

// Blacklisted : GApplicationPrivate

// Blacklisted : GAsyncResultIface

// Blacklisted : GBufferedInputStreamClass

// Blacklisted : GBufferedInputStreamPrivate

// Blacklisted : GBufferedOutputStreamClass

// Blacklisted : GBufferedOutputStreamPrivate

// Blacklisted : GCancellableClass

// Blacklisted : GCancellablePrivate

// Blacklisted : GCharsetConverterClass

// Blacklisted : GConverterInputStreamClass

// Blacklisted : GConverterInputStreamPrivate

// Blacklisted : GConverterOutputStreamClass

// Blacklisted : GConverterOutputStreamPrivate

// Blacklisted : GDBusInterfaceSkeletonPrivate

// Blacklisted : GDBusObjectManagerClientPrivate

// Blacklisted : GDBusObjectManagerServerPrivate

// Blacklisted : GDBusObjectProxyPrivate

// Blacklisted : GDBusObjectSkeletonPrivate

// Blacklisted : GDBusProxyPrivate

// Blacklisted : GDataInputStreamClass

// Blacklisted : GDataInputStreamPrivate

// Blacklisted : GDataOutputStreamClass

// Blacklisted : GDataOutputStreamPrivate

// Blacklisted : GDesktopAppInfoClass

// Blacklisted : GDesktopAppInfoLookupIface

// Blacklisted : GDriveIface

// Blacklisted : GEmblemClass

// Blacklisted : GEmblemedIconClass

// Blacklisted : GEmblemedIconPrivate

// Blacklisted : GFileAttributeInfo

// Blacklisted : GFileAttributeInfoList

// Blacklisted : GFileAttributeMatcher

// Blacklisted : GFileDescriptorBasedIface

// Blacklisted : GFileEnumeratorClass

// Blacklisted : GFileEnumeratorPrivate

// Blacklisted : GFileIOStreamClass

// Blacklisted : GFileIOStreamPrivate

// Blacklisted : GFileIconClass

// Blacklisted : GFileIface

// Blacklisted : GFileInfoClass

// Blacklisted : GFileInputStreamClass

// Blacklisted : GFileInputStreamPrivate

// Blacklisted : GFileMonitorClass

// Blacklisted : GFileMonitorPrivate

// Blacklisted : GFileOutputStreamClass

// Blacklisted : GFileOutputStreamPrivate

// Blacklisted : GFilenameCompleterClass

// Blacklisted : GFilterInputStreamClass

// Blacklisted : GFilterOutputStreamClass

// Blacklisted : GIOExtension

// Blacklisted : GIOExtensionPoint

// Blacklisted : GIOModuleClass

// Blacklisted : GIOSchedulerJob

// Blacklisted : GIOStreamAdapter

// Blacklisted : GIOStreamClass

// Blacklisted : GIOStreamPrivate

// Blacklisted : GIconIface

// Blacklisted : GInetAddressClass

// Blacklisted : GInetAddressMaskClass

// Blacklisted : GInetAddressMaskPrivate

// Blacklisted : GInetAddressPrivate

// Blacklisted : GInetSocketAddressClass

// Blacklisted : GInetSocketAddressPrivate

// Blacklisted : GInputStreamClass

// Blacklisted : GInputStreamPrivate

// Blacklisted : GListStoreClass

// Blacklisted : GLoadableIconIface

// Blacklisted : GMemoryInputStreamClass

// Blacklisted : GMemoryInputStreamPrivate

// Blacklisted : GMemoryOutputStreamClass

// Blacklisted : GMemoryOutputStreamPrivate

// Blacklisted : GMenuAttributeIterClass

// Blacklisted : GMenuAttributeIterPrivate

// Blacklisted : GMenuLinkIterClass

// Blacklisted : GMenuLinkIterPrivate

// Blacklisted : GMenuModelClass

// Blacklisted : GMenuModelPrivate

// Blacklisted : GMountIface

// Blacklisted : GMountOperationClass

// Blacklisted : GMountOperationPrivate

// Blacklisted : GNativeSocketAddress

// Blacklisted : GNativeVolumeMonitorClass

// Blacklisted : GNetworkAddressClass

// Blacklisted : GNetworkAddressPrivate

// Blacklisted : GNetworkServiceClass

// Blacklisted : GNetworkServicePrivate

// Blacklisted : GOutputStreamClass

// Blacklisted : GOutputStreamPrivate

// Blacklisted : GPermissionClass

// Blacklisted : GPermissionPrivate

// Blacklisted : GProxyAddressEnumeratorClass

// Blacklisted : GProxyAddressEnumeratorPrivate

// Blacklisted : GProxyAddressPrivate

// Blacklisted : GProxyResolverInterface

// Blacklisted : GResolverClass

// Blacklisted : GResolverPrivate

// Blacklisted : GSeekableIface

// Blacklisted : GSettingsBackendClass

// Blacklisted : GSettingsBackendPrivate

// Blacklisted : GSettingsClass

// Blacklisted : GSettingsPrivate

// Blacklisted : GSettingsSchemaKey

// Blacklisted : GSimpleActionGroupClass

// Blacklisted : GSimpleActionGroupPrivate

// Blacklisted : GSimpleAsyncResultClass

// Blacklisted : GSimpleProxyResolverClass

// Blacklisted : GSimpleProxyResolverPrivate

// Blacklisted : GSocketAddressClass

// Blacklisted : GSocketAddressEnumeratorClass

// Blacklisted : GSocketClass

// Blacklisted : GSocketClientClass

// Blacklisted : GSocketClientPrivate

// Blacklisted : GSocketConnectableIface

// Blacklisted : GSocketConnectionClass

// Blacklisted : GSocketConnectionPrivate

// Blacklisted : GSocketControlMessageClass

// Blacklisted : GSocketControlMessagePrivate

// Blacklisted : GSocketListenerClass

// Blacklisted : GSocketListenerPrivate

// Blacklisted : GSocketPrivate

// Blacklisted : GSocketServiceClass

// Blacklisted : GSocketServicePrivate

// Blacklisted : GSrvTarget

// Blacklisted : GStaticResource

// Blacklisted : GTaskClass

// Blacklisted : GTcpConnectionClass

// Blacklisted : GTcpConnectionPrivate

// Blacklisted : GTcpWrapperConnectionClass

// Blacklisted : GTcpWrapperConnectionPrivate

// Blacklisted : GThemedIconClass

// Blacklisted : GThreadedSocketServiceClass

// Blacklisted : GThreadedSocketServicePrivate

// Blacklisted : GTlsCertificateClass

// Blacklisted : GTlsCertificatePrivate

// Blacklisted : GTlsConnectionClass

// Blacklisted : GTlsConnectionPrivate

// Blacklisted : GTlsDatabasePrivate

// Blacklisted : GTlsFileDatabaseInterface

// Blacklisted : GTlsInteractionPrivate

// Blacklisted : GTlsPasswordClass

// Blacklisted : GTlsPasswordPrivate

// Blacklisted : GUnixConnectionClass

// Blacklisted : GUnixConnectionPrivate

// Blacklisted : GUnixCredentialsMessagePrivate

// Blacklisted : GUnixFDListClass

// Blacklisted : GUnixFDListPrivate

// Blacklisted : GUnixFDMessageClass

// Blacklisted : GUnixFDMessagePrivate

// Blacklisted : GUnixInputStreamClass

// Blacklisted : GUnixInputStreamPrivate

// Blacklisted : GUnixMountEntry

// Blacklisted : GUnixMountMonitorClass

// Blacklisted : GUnixMountPoint

// Blacklisted : GUnixOutputStreamClass

// Blacklisted : GUnixOutputStreamPrivate

// Blacklisted : GUnixSocketAddressClass

// Blacklisted : GUnixSocketAddressPrivate

// Blacklisted : GVfsClass

// Blacklisted : GVolumeIface

// Blacklisted : GVolumeMonitorClass

// Blacklisted : GZlibCompressorClass

// Blacklisted : GZlibDecompressorClass
