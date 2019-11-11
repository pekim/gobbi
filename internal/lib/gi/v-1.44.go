// Code generated - DO NOT EDIT.
// +build girepository_1.44

package gi

import (
	glib "github.com/pekim/gobbi/internal/lib/glib"
	gobject "github.com/pekim/gobbi/internal/lib/gobject"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <girepository.h>
// #include <stdlib.h>
import "C"

// ArgInfo is a representation of the C alias GIArgInfo.
type ArgInfo *BaseInfo

// CallableInfo is a representation of the C alias GICallableInfo.
type CallableInfo *BaseInfo

// CallbackInfo is a representation of the C alias GICallbackInfo.
type CallbackInfo *BaseInfo

// ConstantInfo is a representation of the C alias GIConstantInfo.
type ConstantInfo *BaseInfo

// EnumInfo is a representation of the C alias GIEnumInfo.
type EnumInfo *BaseInfo

// FieldInfo is a representation of the C alias GIFieldInfo.
type FieldInfo *BaseInfo

// FunctionInfo is a representation of the C alias GIFunctionInfo.
type FunctionInfo *BaseInfo

// InterfaceInfo is a representation of the C alias GIInterfaceInfo.
type InterfaceInfo *BaseInfo

// ObjectInfo is a representation of the C alias GIObjectInfo.
type ObjectInfo *BaseInfo

// PropertyInfo is a representation of the C alias GIPropertyInfo.
type PropertyInfo *BaseInfo

// RegisteredTypeInfo is a representation of the C alias GIRegisteredTypeInfo.
type RegisteredTypeInfo *BaseInfo

// SignalInfo is a representation of the C alias GISignalInfo.
type SignalInfo *BaseInfo

// StructInfo is a representation of the C alias GIStructInfo.
type StructInfo *BaseInfo

// TypeInfo is a representation of the C alias GITypeInfo.
type TypeInfo *BaseInfo

// UnionInfo is a representation of the C alias GIUnionInfo.
type UnionInfo *BaseInfo

// VFuncInfo is a representation of the C alias GIVFuncInfo.
type VFuncInfo *BaseInfo

// ValueInfo is a representation of the C alias GIValueInfo.
type ValueInfo *BaseInfo

type FieldInfoFlags C.GIFieldInfoFlags

const (
	FIELD_IS_READABLE FieldInfoFlags = 1
	FIELD_IS_WRITABLE FieldInfoFlags = 2
)

type FunctionInfoFlags C.GIFunctionInfoFlags

const (
	FUNCTION_IS_METHOD      FunctionInfoFlags = 1
	FUNCTION_IS_CONSTRUCTOR FunctionInfoFlags = 2
	FUNCTION_IS_GETTER      FunctionInfoFlags = 4
	FUNCTION_IS_SETTER      FunctionInfoFlags = 8
	FUNCTION_WRAPS_VFUNC    FunctionInfoFlags = 16
	FUNCTION_THROWS         FunctionInfoFlags = 32
)

type RepositoryLoadFlags C.GIRepositoryLoadFlags

const (
	G_IREPOSITORY_LOAD_FLAG_LAZY RepositoryLoadFlags = 1
)

type VFuncInfoFlags C.GIVFuncInfoFlags

const (
	VFUNC_MUST_CHAIN_UP     VFuncInfoFlags = 1
	VFUNC_MUST_OVERRIDE     VFuncInfoFlags = 2
	VFUNC_MUST_NOT_OVERRIDE VFuncInfoFlags = 4
	VFUNC_THROWS            VFuncInfoFlags = 8
)

// Repository is a wrapper around the C record GIRepository.
type Repository struct {
	native *C.GIRepository
	// Private : parent
	// Private : priv
}

func RepositoryNewFromC(u unsafe.Pointer) *Repository {
	c := (*C.GIRepository)(u)
	if c == nil {
		return nil
	}

	g := &Repository{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Repository) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Repository) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Repository with another Repository, and returns true if they represent the same GObject.
func (recv *Repository) Equals(other *Repository) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Repository) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Repository.
// Exercise care, as this is a potentially dangerous function if the Object is not a Repository.
func CastToRepository(object *gobject.Object) *Repository {
	return RepositoryNewFromC(object.ToC())
}

// RepositoryDump is a wrapper around the C function g_irepository_dump.
func RepositoryDump(arg string) (bool, error) {
	c_arg := C.CString(arg)
	defer C.free(unsafe.Pointer(c_arg))

	var cThrowableError *C.GError

	retC := C.g_irepository_dump(c_arg, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RepositoryErrorQuark is a wrapper around the C function g_irepository_error_quark.
func RepositoryErrorQuark() glib.Quark {
	retC := C.g_irepository_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// RepositoryGetDefault is a wrapper around the C function g_irepository_get_default.
func RepositoryGetDefault() *Repository {
	retC := C.g_irepository_get_default()
	retGo := RepositoryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RepositoryGetOptionGroup is a wrapper around the C function g_irepository_get_option_group.
func RepositoryGetOptionGroup() *glib.OptionGroup {
	retC := C.g_irepository_get_option_group()
	retGo := glib.OptionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RepositoryGetSearchPath is a wrapper around the C function g_irepository_get_search_path.
func RepositoryGetSearchPath() *glib.SList {
	retC := C.g_irepository_get_search_path()
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RepositoryPrependLibraryPath is a wrapper around the C function g_irepository_prepend_library_path.
func RepositoryPrependLibraryPath(directory string) {
	c_directory := C.CString(directory)
	defer C.free(unsafe.Pointer(c_directory))

	C.g_irepository_prepend_library_path(c_directory)

	return
}

// RepositoryPrependSearchPath is a wrapper around the C function g_irepository_prepend_search_path.
func RepositoryPrependSearchPath(directory string) {
	c_directory := C.CString(directory)
	defer C.free(unsafe.Pointer(c_directory))

	C.g_irepository_prepend_search_path(c_directory)

	return
}

// EnumerateVersions is a wrapper around the C function g_irepository_enumerate_versions.
func (recv *Repository) EnumerateVersions(namespace string) *glib.List {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	retC := C.g_irepository_enumerate_versions((*C.GIRepository)(recv.native), c_namespace_)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindByErrorDomain is a wrapper around the C function g_irepository_find_by_error_domain.
func (recv *Repository) FindByErrorDomain(domain glib.Quark) *EnumInfo {
	c_domain := (C.GQuark)(domain)

	retC := C.g_irepository_find_by_error_domain((*C.GIRepository)(recv.native), c_domain)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindByGtype is a wrapper around the C function g_irepository_find_by_gtype.
func (recv *Repository) FindByGtype(gtype gobject.Type) *BaseInfo {
	c_gtype := (C.GType)(gtype)

	retC := C.g_irepository_find_by_gtype((*C.GIRepository)(recv.native), c_gtype)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindByName is a wrapper around the C function g_irepository_find_by_name.
func (recv *Repository) FindByName(namespace string, name string) *BaseInfo {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_irepository_find_by_name((*C.GIRepository)(recv.native), c_namespace_, c_name)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCPrefix is a wrapper around the C function g_irepository_get_c_prefix.
func (recv *Repository) GetCPrefix(namespace string) string {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	retC := C.g_irepository_get_c_prefix((*C.GIRepository)(recv.native), c_namespace_)
	retGo := C.GoString(retC)

	return retGo
}

// GetDependencies is a wrapper around the C function g_irepository_get_dependencies.
func (recv *Repository) GetDependencies(namespace string) []string {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	retC := C.g_irepository_get_dependencies((*C.GIRepository)(recv.native), c_namespace_)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetImmediateDependencies is a wrapper around the C function g_irepository_get_immediate_dependencies.
func (recv *Repository) GetImmediateDependencies(namespace string) []string {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	retC := C.g_irepository_get_immediate_dependencies((*C.GIRepository)(recv.native), c_namespace_)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetInfo is a wrapper around the C function g_irepository_get_info.
func (recv *Repository) GetInfo(namespace string, index int32) *BaseInfo {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	c_index := (C.gint)(index)

	retC := C.g_irepository_get_info((*C.GIRepository)(recv.native), c_namespace_, c_index)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLoadedNamespaces is a wrapper around the C function g_irepository_get_loaded_namespaces.
func (recv *Repository) GetLoadedNamespaces() []string {
	retC := C.g_irepository_get_loaded_namespaces((*C.GIRepository)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetNInfos is a wrapper around the C function g_irepository_get_n_infos.
func (recv *Repository) GetNInfos(namespace string) int32 {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	retC := C.g_irepository_get_n_infos((*C.GIRepository)(recv.native), c_namespace_)
	retGo := (int32)(retC)

	return retGo
}

// GetSharedLibrary is a wrapper around the C function g_irepository_get_shared_library.
func (recv *Repository) GetSharedLibrary(namespace string) string {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	retC := C.g_irepository_get_shared_library((*C.GIRepository)(recv.native), c_namespace_)
	retGo := C.GoString(retC)

	return retGo
}

// GetTypelibPath is a wrapper around the C function g_irepository_get_typelib_path.
func (recv *Repository) GetTypelibPath(namespace string) string {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	retC := C.g_irepository_get_typelib_path((*C.GIRepository)(recv.native), c_namespace_)
	retGo := C.GoString(retC)

	return retGo
}

// GetVersion is a wrapper around the C function g_irepository_get_version.
func (recv *Repository) GetVersion(namespace string) string {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	retC := C.g_irepository_get_version((*C.GIRepository)(recv.native), c_namespace_)
	retGo := C.GoString(retC)

	return retGo
}

// IsRegistered is a wrapper around the C function g_irepository_is_registered.
func (recv *Repository) IsRegistered(namespace string, version string) bool {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	c_version := C.CString(version)
	defer C.free(unsafe.Pointer(c_version))

	retC := C.g_irepository_is_registered((*C.GIRepository)(recv.native), c_namespace_, c_version)
	retGo := retC == C.TRUE

	return retGo
}

// LoadTypelib is a wrapper around the C function g_irepository_load_typelib.
func (recv *Repository) LoadTypelib(typelib *Typelib, flags RepositoryLoadFlags) (string, error) {
	c_typelib := (*C.GITypelib)(C.NULL)
	if typelib != nil {
		c_typelib = (*C.GITypelib)(typelib.ToC())
	}

	c_flags := (C.GIRepositoryLoadFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_irepository_load_typelib((*C.GIRepository)(recv.native), c_typelib, c_flags, &cThrowableError)
	retGo := C.GoString(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Require is a wrapper around the C function g_irepository_require.
func (recv *Repository) Require(namespace string, version string, flags RepositoryLoadFlags) (*Typelib, error) {
	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	c_version := C.CString(version)
	defer C.free(unsafe.Pointer(c_version))

	c_flags := (C.GIRepositoryLoadFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_irepository_require((*C.GIRepository)(recv.native), c_namespace_, c_version, c_flags, &cThrowableError)
	retGo := TypelibNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RequirePrivate is a wrapper around the C function g_irepository_require_private.
func (recv *Repository) RequirePrivate(typelibDir string, namespace string, version string, flags RepositoryLoadFlags) (*Typelib, error) {
	c_typelib_dir := C.CString(typelibDir)
	defer C.free(unsafe.Pointer(c_typelib_dir))

	c_namespace_ := C.CString(namespace)
	defer C.free(unsafe.Pointer(c_namespace_))

	c_version := C.CString(version)
	defer C.free(unsafe.Pointer(c_version))

	c_flags := (C.GIRepositoryLoadFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_irepository_require_private((*C.GIRepository)(recv.native), c_typelib_dir, c_namespace_, c_version, c_flags, &cThrowableError)
	retGo := TypelibNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

type ArrayType C.GIArrayType

const (
	ARRAY_TYPE_C          ArrayType = 0
	ARRAY_TYPE_ARRAY      ArrayType = 1
	ARRAY_TYPE_PTR_ARRAY  ArrayType = 2
	ARRAY_TYPE_BYTE_ARRAY ArrayType = 3
)

type Direction C.GIDirection

const (
	DIRECTION_IN    Direction = 0
	DIRECTION_OUT   Direction = 1
	DIRECTION_INOUT Direction = 2
)

type InfoType C.GIInfoType

const (
	INFO_TYPE_INVALID    InfoType = 0
	INFO_TYPE_FUNCTION   InfoType = 1
	INFO_TYPE_CALLBACK   InfoType = 2
	INFO_TYPE_STRUCT     InfoType = 3
	INFO_TYPE_BOXED      InfoType = 4
	INFO_TYPE_ENUM       InfoType = 5
	INFO_TYPE_FLAGS      InfoType = 6
	INFO_TYPE_OBJECT     InfoType = 7
	INFO_TYPE_INTERFACE  InfoType = 8
	INFO_TYPE_CONSTANT   InfoType = 9
	INFO_TYPE_INVALID_0  InfoType = 10
	INFO_TYPE_UNION      InfoType = 11
	INFO_TYPE_VALUE      InfoType = 12
	INFO_TYPE_SIGNAL     InfoType = 13
	INFO_TYPE_VFUNC      InfoType = 14
	INFO_TYPE_PROPERTY   InfoType = 15
	INFO_TYPE_FIELD      InfoType = 16
	INFO_TYPE_ARG        InfoType = 17
	INFO_TYPE_TYPE       InfoType = 18
	INFO_TYPE_UNRESOLVED InfoType = 19
)

type RepositoryError C.GIRepositoryError

const (
	G_IREPOSITORY_ERROR_TYPELIB_NOT_FOUND          RepositoryError = 0
	G_IREPOSITORY_ERROR_NAMESPACE_MISMATCH         RepositoryError = 1
	G_IREPOSITORY_ERROR_NAMESPACE_VERSION_CONFLICT RepositoryError = 2
	G_IREPOSITORY_ERROR_LIBRARY_NOT_FOUND          RepositoryError = 3
)

type ScopeType C.GIScopeType

const (
	SCOPE_TYPE_INVALID  ScopeType = 0
	SCOPE_TYPE_CALL     ScopeType = 1
	SCOPE_TYPE_ASYNC    ScopeType = 2
	SCOPE_TYPE_NOTIFIED ScopeType = 3
)

type Transfer C.GITransfer

const (
	TRANSFER_NOTHING    Transfer = 0
	TRANSFER_CONTAINER  Transfer = 1
	TRANSFER_EVERYTHING Transfer = 2
)

type TypeTag C.GITypeTag

const (
	TYPE_TAG_VOID      TypeTag = 0
	TYPE_TAG_BOOLEAN   TypeTag = 1
	TYPE_TAG_INT8      TypeTag = 2
	TYPE_TAG_UINT8     TypeTag = 3
	TYPE_TAG_INT16     TypeTag = 4
	TYPE_TAG_UINT16    TypeTag = 5
	TYPE_TAG_INT32     TypeTag = 6
	TYPE_TAG_UINT32    TypeTag = 7
	TYPE_TAG_INT64     TypeTag = 8
	TYPE_TAG_UINT64    TypeTag = 9
	TYPE_TAG_FLOAT     TypeTag = 10
	TYPE_TAG_DOUBLE    TypeTag = 11
	TYPE_TAG_GTYPE     TypeTag = 12
	TYPE_TAG_UTF8      TypeTag = 13
	TYPE_TAG_FILENAME  TypeTag = 14
	TYPE_TAG_ARRAY     TypeTag = 15
	TYPE_TAG_INTERFACE TypeTag = 16
	TYPE_TAG_GLIST     TypeTag = 17
	TYPE_TAG_GSLIST    TypeTag = 18
	TYPE_TAG_GHASH     TypeTag = 19
	TYPE_TAG_ERROR     TypeTag = 20
	TYPE_TAG_UNICHAR   TypeTag = 21
)

type nvokeError C.GInvokeError

const (
	G_INVOKE_ERROR_FAILED            nvokeError = 0
	G_INVOKE_ERROR_SYMBOL_NOT_FOUND  nvokeError = 1
	G_INVOKE_ERROR_ARGUMENT_MISMATCH nvokeError = 2
)

// ArgInfoGetClosure is a wrapper around the C function g_arg_info_get_closure.
func ArgInfoGetClosure(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_get_closure(c_info)
	retGo := (int32)(retC)

	return retGo
}

// ArgInfoGetDestroy is a wrapper around the C function g_arg_info_get_destroy.
func ArgInfoGetDestroy(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_get_destroy(c_info)
	retGo := (int32)(retC)

	return retGo
}

// ArgInfoGetDirection is a wrapper around the C function g_arg_info_get_direction.
func ArgInfoGetDirection(info *BaseInfo) Direction {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_get_direction(c_info)
	retGo := (Direction)(retC)

	return retGo
}

// ArgInfoGetOwnershipTransfer is a wrapper around the C function g_arg_info_get_ownership_transfer.
func ArgInfoGetOwnershipTransfer(info *BaseInfo) Transfer {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_get_ownership_transfer(c_info)
	retGo := (Transfer)(retC)

	return retGo
}

// ArgInfoGetScope is a wrapper around the C function g_arg_info_get_scope.
func ArgInfoGetScope(info *BaseInfo) ScopeType {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_get_scope(c_info)
	retGo := (ScopeType)(retC)

	return retGo
}

// ArgInfoGetType is a wrapper around the C function g_arg_info_get_type.
func ArgInfoGetType(info *BaseInfo) *TypeInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_get_type(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ArgInfoIsCallerAllocates is a wrapper around the C function g_arg_info_is_caller_allocates.
func ArgInfoIsCallerAllocates(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_is_caller_allocates(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// ArgInfoIsOptional is a wrapper around the C function g_arg_info_is_optional.
func ArgInfoIsOptional(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_is_optional(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// ArgInfoIsReturnValue is a wrapper around the C function g_arg_info_is_return_value.
func ArgInfoIsReturnValue(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_is_return_value(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// ArgInfoIsSkip is a wrapper around the C function g_arg_info_is_skip.
func ArgInfoIsSkip(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_is_skip(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// ArgInfoLoadType is a wrapper around the C function g_arg_info_load_type.
func ArgInfoLoadType(info *BaseInfo) *BaseInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	var c_type C.GIBaseInfo

	C.g_arg_info_load_type(c_info, &c_type)

	type_ := BaseInfoNewFromC(unsafe.Pointer(&c_type))

	return type_
}

// ArgInfoMayBeNull is a wrapper around the C function g_arg_info_may_be_null.
func ArgInfoMayBeNull(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_arg_info_may_be_null(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// CallableInfoCanThrowGerror is a wrapper around the C function g_callable_info_can_throw_gerror.
func CallableInfoCanThrowGerror(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_callable_info_can_throw_gerror(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// CallableInfoGetArg is a wrapper around the C function g_callable_info_get_arg.
func CallableInfoGetArg(info *BaseInfo, n int32) *ArgInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_callable_info_get_arg(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CallableInfoGetCallerOwns is a wrapper around the C function g_callable_info_get_caller_owns.
func CallableInfoGetCallerOwns(info *BaseInfo) Transfer {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_callable_info_get_caller_owns(c_info)
	retGo := (Transfer)(retC)

	return retGo
}

// CallableInfoGetInstanceOwnershipTransfer is a wrapper around the C function g_callable_info_get_instance_ownership_transfer.
func CallableInfoGetInstanceOwnershipTransfer(info *BaseInfo) Transfer {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_callable_info_get_instance_ownership_transfer(c_info)
	retGo := (Transfer)(retC)

	return retGo
}

// CallableInfoGetNArgs is a wrapper around the C function g_callable_info_get_n_args.
func CallableInfoGetNArgs(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_callable_info_get_n_args(c_info)
	retGo := (int32)(retC)

	return retGo
}

// CallableInfoGetReturnAttribute is a wrapper around the C function g_callable_info_get_return_attribute.
func CallableInfoGetReturnAttribute(info *BaseInfo, name string) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_callable_info_get_return_attribute(c_info, c_name)
	retGo := C.GoString(retC)

	return retGo
}

// CallableInfoGetReturnType is a wrapper around the C function g_callable_info_get_return_type.
func CallableInfoGetReturnType(info *BaseInfo) *TypeInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_callable_info_get_return_type(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_callable_info_invoke : unsupported parameter in_args : no type generator for Argument (GIArgument) for array param in_args

// CallableInfoIsMethod is a wrapper around the C function g_callable_info_is_method.
func CallableInfoIsMethod(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_callable_info_is_method(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// CallableInfoIterateReturnAttributes is a wrapper around the C function g_callable_info_iterate_return_attributes.
func CallableInfoIterateReturnAttributes(info *BaseInfo, iterator *AttributeIter) (bool, string, string) {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_iterator := (*C.GIAttributeIter)(C.NULL)
	if iterator != nil {
		c_iterator = (*C.GIAttributeIter)(iterator.ToC())
	}

	var c_name *C.char

	var c_value *C.char

	retC := C.g_callable_info_iterate_return_attributes(c_info, c_iterator, &c_name, &c_value)
	retGo := retC == C.TRUE

	name := C.GoString(c_name)

	value := C.GoString(c_value)

	return retGo, name, value
}

// CallableInfoLoadArg is a wrapper around the C function g_callable_info_load_arg.
func CallableInfoLoadArg(info *BaseInfo, n int32) *BaseInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	var c_arg C.GIBaseInfo

	C.g_callable_info_load_arg(c_info, c_n, &c_arg)

	arg := BaseInfoNewFromC(unsafe.Pointer(&c_arg))

	return arg
}

// CallableInfoLoadReturnType is a wrapper around the C function g_callable_info_load_return_type.
func CallableInfoLoadReturnType(info *BaseInfo) *BaseInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	var c_type C.GIBaseInfo

	C.g_callable_info_load_return_type(c_info, &c_type)

	type_ := BaseInfoNewFromC(unsafe.Pointer(&c_type))

	return type_
}

// CallableInfoMayReturnNull is a wrapper around the C function g_callable_info_may_return_null.
func CallableInfoMayReturnNull(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_callable_info_may_return_null(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// CallableInfoSkipReturn is a wrapper around the C function g_callable_info_skip_return.
func CallableInfoSkipReturn(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_callable_info_skip_return(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_constant_info_free_value : unsupported parameter value : no type generator for Argument (GIArgument*) for param value

// ConstantInfoGetType is a wrapper around the C function g_constant_info_get_type.
func ConstantInfoGetType(info *BaseInfo) *TypeInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_constant_info_get_type(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_constant_info_get_value : unsupported parameter value : no type generator for Argument (GIArgument*) for param value

// EnumInfoGetErrorDomain is a wrapper around the C function g_enum_info_get_error_domain.
func EnumInfoGetErrorDomain(info *BaseInfo) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_enum_info_get_error_domain(c_info)
	retGo := C.GoString(retC)

	return retGo
}

// EnumInfoGetMethod is a wrapper around the C function g_enum_info_get_method.
func EnumInfoGetMethod(info *BaseInfo, n int32) *FunctionInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_enum_info_get_method(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumInfoGetNMethods is a wrapper around the C function g_enum_info_get_n_methods.
func EnumInfoGetNMethods(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_enum_info_get_n_methods(c_info)
	retGo := (int32)(retC)

	return retGo
}

// EnumInfoGetNValues is a wrapper around the C function g_enum_info_get_n_values.
func EnumInfoGetNValues(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_enum_info_get_n_values(c_info)
	retGo := (int32)(retC)

	return retGo
}

// EnumInfoGetStorageType is a wrapper around the C function g_enum_info_get_storage_type.
func EnumInfoGetStorageType(info *BaseInfo) TypeTag {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_enum_info_get_storage_type(c_info)
	retGo := (TypeTag)(retC)

	return retGo
}

// EnumInfoGetValue is a wrapper around the C function g_enum_info_get_value.
func EnumInfoGetValue(info *BaseInfo, n int32) *ValueInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_enum_info_get_value(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_field_info_get_field : unsupported parameter value : no type generator for Argument (GIArgument*) for param value

// FieldInfoGetFlags is a wrapper around the C function g_field_info_get_flags.
func FieldInfoGetFlags(info *BaseInfo) FieldInfoFlags {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_field_info_get_flags(c_info)
	retGo := (FieldInfoFlags)(retC)

	return retGo
}

// FieldInfoGetOffset is a wrapper around the C function g_field_info_get_offset.
func FieldInfoGetOffset(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_field_info_get_offset(c_info)
	retGo := (int32)(retC)

	return retGo
}

// FieldInfoGetSize is a wrapper around the C function g_field_info_get_size.
func FieldInfoGetSize(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_field_info_get_size(c_info)
	retGo := (int32)(retC)

	return retGo
}

// FieldInfoGetType is a wrapper around the C function g_field_info_get_type.
func FieldInfoGetType(info *BaseInfo) *TypeInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_field_info_get_type(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_field_info_set_field : unsupported parameter value : no type generator for Argument (const GIArgument*) for param value

// FunctionInfoGetFlags is a wrapper around the C function g_function_info_get_flags.
func FunctionInfoGetFlags(info *BaseInfo) FunctionInfoFlags {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_function_info_get_flags(c_info)
	retGo := (FunctionInfoFlags)(retC)

	return retGo
}

// FunctionInfoGetProperty is a wrapper around the C function g_function_info_get_property.
func FunctionInfoGetProperty(info *BaseInfo) *PropertyInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_function_info_get_property(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FunctionInfoGetSymbol is a wrapper around the C function g_function_info_get_symbol.
func FunctionInfoGetSymbol(info *BaseInfo) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_function_info_get_symbol(c_info)
	retGo := C.GoString(retC)

	return retGo
}

// FunctionInfoGetVfunc is a wrapper around the C function g_function_info_get_vfunc.
func FunctionInfoGetVfunc(info *BaseInfo) *VFuncInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_function_info_get_vfunc(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_function_info_invoke : unsupported parameter in_args : no type generator for Argument (GIArgument) for array param in_args

// InfoNew is a wrapper around the C function g_info_new.
func InfoNew(type_ InfoType, container *BaseInfo, typelib *Typelib, offset uint32) *BaseInfo {
	c_type := (C.GIInfoType)(type_)

	c_container := (*C.GIBaseInfo)(C.NULL)
	if container != nil {
		c_container = (*C.GIBaseInfo)(container.ToC())
	}

	c_typelib := (*C.GITypelib)(C.NULL)
	if typelib != nil {
		c_typelib = (*C.GITypelib)(typelib.ToC())
	}

	c_offset := (C.guint32)(offset)

	retC := C.g_info_new(c_type, c_container, c_typelib, c_offset)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InfoTypeToString is a wrapper around the C function g_info_type_to_string.
func InfoTypeToString(type_ InfoType) string {
	c_type := (C.GIInfoType)(type_)

	retC := C.g_info_type_to_string(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// InterfaceInfoFindMethod is a wrapper around the C function g_interface_info_find_method.
func InterfaceInfoFindMethod(info *BaseInfo, name string) *FunctionInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_interface_info_find_method(c_info, c_name)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InterfaceInfoFindSignal is a wrapper around the C function g_interface_info_find_signal.
func InterfaceInfoFindSignal(info *BaseInfo, name string) *SignalInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_interface_info_find_signal(c_info, c_name)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InterfaceInfoFindVfunc is a wrapper around the C function g_interface_info_find_vfunc.
func InterfaceInfoFindVfunc(info *BaseInfo, name string) *VFuncInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_interface_info_find_vfunc(c_info, c_name)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InterfaceInfoGetConstant is a wrapper around the C function g_interface_info_get_constant.
func InterfaceInfoGetConstant(info *BaseInfo, n int32) *ConstantInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_interface_info_get_constant(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InterfaceInfoGetIfaceStruct is a wrapper around the C function g_interface_info_get_iface_struct.
func InterfaceInfoGetIfaceStruct(info *BaseInfo) *StructInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_interface_info_get_iface_struct(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InterfaceInfoGetMethod is a wrapper around the C function g_interface_info_get_method.
func InterfaceInfoGetMethod(info *BaseInfo, n int32) *FunctionInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_interface_info_get_method(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InterfaceInfoGetNConstants is a wrapper around the C function g_interface_info_get_n_constants.
func InterfaceInfoGetNConstants(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_interface_info_get_n_constants(c_info)
	retGo := (int32)(retC)

	return retGo
}

// InterfaceInfoGetNMethods is a wrapper around the C function g_interface_info_get_n_methods.
func InterfaceInfoGetNMethods(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_interface_info_get_n_methods(c_info)
	retGo := (int32)(retC)

	return retGo
}

// InterfaceInfoGetNPrerequisites is a wrapper around the C function g_interface_info_get_n_prerequisites.
func InterfaceInfoGetNPrerequisites(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_interface_info_get_n_prerequisites(c_info)
	retGo := (int32)(retC)

	return retGo
}

// InterfaceInfoGetNProperties is a wrapper around the C function g_interface_info_get_n_properties.
func InterfaceInfoGetNProperties(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_interface_info_get_n_properties(c_info)
	retGo := (int32)(retC)

	return retGo
}

// InterfaceInfoGetNSignals is a wrapper around the C function g_interface_info_get_n_signals.
func InterfaceInfoGetNSignals(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_interface_info_get_n_signals(c_info)
	retGo := (int32)(retC)

	return retGo
}

// InterfaceInfoGetNVfuncs is a wrapper around the C function g_interface_info_get_n_vfuncs.
func InterfaceInfoGetNVfuncs(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_interface_info_get_n_vfuncs(c_info)
	retGo := (int32)(retC)

	return retGo
}

// InterfaceInfoGetPrerequisite is a wrapper around the C function g_interface_info_get_prerequisite.
func InterfaceInfoGetPrerequisite(info *BaseInfo, n int32) *BaseInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_interface_info_get_prerequisite(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InterfaceInfoGetProperty is a wrapper around the C function g_interface_info_get_property.
func InterfaceInfoGetProperty(info *BaseInfo, n int32) *PropertyInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_interface_info_get_property(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InterfaceInfoGetSignal is a wrapper around the C function g_interface_info_get_signal.
func InterfaceInfoGetSignal(info *BaseInfo, n int32) *SignalInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_interface_info_get_signal(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InterfaceInfoGetVfunc is a wrapper around the C function g_interface_info_get_vfunc.
func InterfaceInfoGetVfunc(info *BaseInfo, n int32) *VFuncInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_interface_info_get_vfunc(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InvokeErrorQuark is a wrapper around the C function g_invoke_error_quark.
func InvokeErrorQuark() glib.Quark {
	retC := C.g_invoke_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// ObjectInfoFindMethod is a wrapper around the C function g_object_info_find_method.
func ObjectInfoFindMethod(info *BaseInfo, name string) *FunctionInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_object_info_find_method(c_info, c_name)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoFindMethodUsingInterfaces is a wrapper around the C function g_object_info_find_method_using_interfaces.
func ObjectInfoFindMethodUsingInterfaces(info *BaseInfo, name string) (*FunctionInfo, *BaseInfo) {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var c_implementor *C.GIBaseInfo

	retC := C.g_object_info_find_method_using_interfaces(c_info, c_name, &c_implementor)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	implementor := BaseInfoNewFromC(unsafe.Pointer(c_implementor))

	return retGo, implementor
}

// ObjectInfoFindSignal is a wrapper around the C function g_object_info_find_signal.
func ObjectInfoFindSignal(info *BaseInfo, name string) *SignalInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_object_info_find_signal(c_info, c_name)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoFindVfunc is a wrapper around the C function g_object_info_find_vfunc.
func ObjectInfoFindVfunc(info *BaseInfo, name string) *VFuncInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_object_info_find_vfunc(c_info, c_name)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoFindVfuncUsingInterfaces is a wrapper around the C function g_object_info_find_vfunc_using_interfaces.
func ObjectInfoFindVfuncUsingInterfaces(info *BaseInfo, name string) (*VFuncInfo, *BaseInfo) {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var c_implementor *C.GIBaseInfo

	retC := C.g_object_info_find_vfunc_using_interfaces(c_info, c_name, &c_implementor)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	implementor := BaseInfoNewFromC(unsafe.Pointer(c_implementor))

	return retGo, implementor
}

// ObjectInfoGetAbstract is a wrapper around the C function g_object_info_get_abstract.
func ObjectInfoGetAbstract(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_abstract(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// ObjectInfoGetClassStruct is a wrapper around the C function g_object_info_get_class_struct.
func ObjectInfoGetClassStruct(info *BaseInfo) *StructInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_class_struct(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoGetConstant is a wrapper around the C function g_object_info_get_constant.
func ObjectInfoGetConstant(info *BaseInfo, n int32) *ConstantInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_object_info_get_constant(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoGetField is a wrapper around the C function g_object_info_get_field.
func ObjectInfoGetField(info *BaseInfo, n int32) *FieldInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_object_info_get_field(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoGetFundamental is a wrapper around the C function g_object_info_get_fundamental.
func ObjectInfoGetFundamental(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_fundamental(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// ObjectInfoGetGetValueFunction is a wrapper around the C function g_object_info_get_get_value_function.
func ObjectInfoGetGetValueFunction(info *BaseInfo) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_get_value_function(c_info)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_object_info_get_get_value_function_pointer : no return generator

// ObjectInfoGetInterface is a wrapper around the C function g_object_info_get_interface.
func ObjectInfoGetInterface(info *BaseInfo, n int32) *InterfaceInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_object_info_get_interface(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoGetMethod is a wrapper around the C function g_object_info_get_method.
func ObjectInfoGetMethod(info *BaseInfo, n int32) *FunctionInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_object_info_get_method(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoGetNConstants is a wrapper around the C function g_object_info_get_n_constants.
func ObjectInfoGetNConstants(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_n_constants(c_info)
	retGo := (int32)(retC)

	return retGo
}

// ObjectInfoGetNFields is a wrapper around the C function g_object_info_get_n_fields.
func ObjectInfoGetNFields(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_n_fields(c_info)
	retGo := (int32)(retC)

	return retGo
}

// ObjectInfoGetNInterfaces is a wrapper around the C function g_object_info_get_n_interfaces.
func ObjectInfoGetNInterfaces(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_n_interfaces(c_info)
	retGo := (int32)(retC)

	return retGo
}

// ObjectInfoGetNMethods is a wrapper around the C function g_object_info_get_n_methods.
func ObjectInfoGetNMethods(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_n_methods(c_info)
	retGo := (int32)(retC)

	return retGo
}

// ObjectInfoGetNProperties is a wrapper around the C function g_object_info_get_n_properties.
func ObjectInfoGetNProperties(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_n_properties(c_info)
	retGo := (int32)(retC)

	return retGo
}

// ObjectInfoGetNSignals is a wrapper around the C function g_object_info_get_n_signals.
func ObjectInfoGetNSignals(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_n_signals(c_info)
	retGo := (int32)(retC)

	return retGo
}

// ObjectInfoGetNVfuncs is a wrapper around the C function g_object_info_get_n_vfuncs.
func ObjectInfoGetNVfuncs(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_n_vfuncs(c_info)
	retGo := (int32)(retC)

	return retGo
}

// ObjectInfoGetParent is a wrapper around the C function g_object_info_get_parent.
func ObjectInfoGetParent(info *BaseInfo) *ObjectInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_parent(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoGetProperty is a wrapper around the C function g_object_info_get_property.
func ObjectInfoGetProperty(info *BaseInfo, n int32) *PropertyInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_object_info_get_property(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoGetRefFunction is a wrapper around the C function g_object_info_get_ref_function.
func ObjectInfoGetRefFunction(info *BaseInfo) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_ref_function(c_info)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_object_info_get_ref_function_pointer : no return generator

// ObjectInfoGetSetValueFunction is a wrapper around the C function g_object_info_get_set_value_function.
func ObjectInfoGetSetValueFunction(info *BaseInfo) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_set_value_function(c_info)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_object_info_get_set_value_function_pointer : no return generator

// ObjectInfoGetSignal is a wrapper around the C function g_object_info_get_signal.
func ObjectInfoGetSignal(info *BaseInfo, n int32) *SignalInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_object_info_get_signal(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ObjectInfoGetTypeInit is a wrapper around the C function g_object_info_get_type_init.
func ObjectInfoGetTypeInit(info *BaseInfo) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_type_init(c_info)
	retGo := C.GoString(retC)

	return retGo
}

// ObjectInfoGetTypeName is a wrapper around the C function g_object_info_get_type_name.
func ObjectInfoGetTypeName(info *BaseInfo) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_type_name(c_info)
	retGo := C.GoString(retC)

	return retGo
}

// ObjectInfoGetUnrefFunction is a wrapper around the C function g_object_info_get_unref_function.
func ObjectInfoGetUnrefFunction(info *BaseInfo) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_object_info_get_unref_function(c_info)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_object_info_get_unref_function_pointer : no return generator

// ObjectInfoGetVfunc is a wrapper around the C function g_object_info_get_vfunc.
func ObjectInfoGetVfunc(info *BaseInfo, n int32) *VFuncInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_object_info_get_vfunc(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PropertyInfoGetFlags is a wrapper around the C function g_property_info_get_flags.
func PropertyInfoGetFlags(info *BaseInfo) gobject.ParamFlags {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_property_info_get_flags(c_info)
	retGo := (gobject.ParamFlags)(retC)

	return retGo
}

// PropertyInfoGetOwnershipTransfer is a wrapper around the C function g_property_info_get_ownership_transfer.
func PropertyInfoGetOwnershipTransfer(info *BaseInfo) Transfer {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_property_info_get_ownership_transfer(c_info)
	retGo := (Transfer)(retC)

	return retGo
}

// PropertyInfoGetType is a wrapper around the C function g_property_info_get_type.
func PropertyInfoGetType(info *BaseInfo) *TypeInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_property_info_get_type(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RegisteredTypeInfoGetGType is a wrapper around the C function g_registered_type_info_get_g_type.
func RegisteredTypeInfoGetGType(info *BaseInfo) gobject.Type {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_registered_type_info_get_g_type(c_info)
	retGo := (gobject.Type)(retC)

	return retGo
}

// RegisteredTypeInfoGetTypeInit is a wrapper around the C function g_registered_type_info_get_type_init.
func RegisteredTypeInfoGetTypeInit(info *BaseInfo) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_registered_type_info_get_type_init(c_info)
	retGo := C.GoString(retC)

	return retGo
}

// RegisteredTypeInfoGetTypeName is a wrapper around the C function g_registered_type_info_get_type_name.
func RegisteredTypeInfoGetTypeName(info *BaseInfo) string {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_registered_type_info_get_type_name(c_info)
	retGo := C.GoString(retC)

	return retGo
}

// SignalInfoGetClassClosure is a wrapper around the C function g_signal_info_get_class_closure.
func SignalInfoGetClassClosure(info *BaseInfo) *VFuncInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_signal_info_get_class_closure(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SignalInfoGetFlags is a wrapper around the C function g_signal_info_get_flags.
func SignalInfoGetFlags(info *BaseInfo) gobject.SignalFlags {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_signal_info_get_flags(c_info)
	retGo := (gobject.SignalFlags)(retC)

	return retGo
}

// SignalInfoTrueStopsEmit is a wrapper around the C function g_signal_info_true_stops_emit.
func SignalInfoTrueStopsEmit(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_signal_info_true_stops_emit(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// StructInfoFindField is a wrapper around the C function g_struct_info_find_field.
func StructInfoFindField(info *BaseInfo, name string) *FieldInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_struct_info_find_field(c_info, c_name)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StructInfoFindMethod is a wrapper around the C function g_struct_info_find_method.
func StructInfoFindMethod(info *BaseInfo, name string) *FunctionInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_struct_info_find_method(c_info, c_name)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StructInfoGetAlignment is a wrapper around the C function g_struct_info_get_alignment.
func StructInfoGetAlignment(info *BaseInfo) uint64 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_struct_info_get_alignment(c_info)
	retGo := (uint64)(retC)

	return retGo
}

// StructInfoGetField is a wrapper around the C function g_struct_info_get_field.
func StructInfoGetField(info *BaseInfo, n int32) *FieldInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_struct_info_get_field(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StructInfoGetMethod is a wrapper around the C function g_struct_info_get_method.
func StructInfoGetMethod(info *BaseInfo, n int32) *FunctionInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_struct_info_get_method(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StructInfoGetNFields is a wrapper around the C function g_struct_info_get_n_fields.
func StructInfoGetNFields(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_struct_info_get_n_fields(c_info)
	retGo := (int32)(retC)

	return retGo
}

// StructInfoGetNMethods is a wrapper around the C function g_struct_info_get_n_methods.
func StructInfoGetNMethods(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_struct_info_get_n_methods(c_info)
	retGo := (int32)(retC)

	return retGo
}

// StructInfoGetSize is a wrapper around the C function g_struct_info_get_size.
func StructInfoGetSize(info *BaseInfo) uint64 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_struct_info_get_size(c_info)
	retGo := (uint64)(retC)

	return retGo
}

// StructInfoIsForeign is a wrapper around the C function g_struct_info_is_foreign.
func StructInfoIsForeign(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_struct_info_is_foreign(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// StructInfoIsGtypeStruct is a wrapper around the C function g_struct_info_is_gtype_struct.
func StructInfoIsGtypeStruct(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_struct_info_is_gtype_struct(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// TypeInfoGetArrayFixedSize is a wrapper around the C function g_type_info_get_array_fixed_size.
func TypeInfoGetArrayFixedSize(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_type_info_get_array_fixed_size(c_info)
	retGo := (int32)(retC)

	return retGo
}

// TypeInfoGetArrayLength is a wrapper around the C function g_type_info_get_array_length.
func TypeInfoGetArrayLength(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_type_info_get_array_length(c_info)
	retGo := (int32)(retC)

	return retGo
}

// TypeInfoGetArrayType is a wrapper around the C function g_type_info_get_array_type.
func TypeInfoGetArrayType(info *BaseInfo) ArrayType {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_type_info_get_array_type(c_info)
	retGo := (ArrayType)(retC)

	return retGo
}

// TypeInfoGetInterface is a wrapper around the C function g_type_info_get_interface.
func TypeInfoGetInterface(info *BaseInfo) *BaseInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_type_info_get_interface(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeInfoGetParamType is a wrapper around the C function g_type_info_get_param_type.
func TypeInfoGetParamType(info *BaseInfo, n int32) *TypeInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_type_info_get_param_type(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TypeInfoGetTag is a wrapper around the C function g_type_info_get_tag.
func TypeInfoGetTag(info *BaseInfo) TypeTag {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_type_info_get_tag(c_info)
	retGo := (TypeTag)(retC)

	return retGo
}

// TypeInfoIsPointer is a wrapper around the C function g_type_info_is_pointer.
func TypeInfoIsPointer(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_type_info_is_pointer(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// TypeInfoIsZeroTerminated is a wrapper around the C function g_type_info_is_zero_terminated.
func TypeInfoIsZeroTerminated(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_type_info_is_zero_terminated(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// TypeTagToString is a wrapper around the C function g_type_tag_to_string.
func TypeTagToString(type_ TypeTag) string {
	c_type := (C.GITypeTag)(type_)

	retC := C.g_type_tag_to_string(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// UnionInfoFindMethod is a wrapper around the C function g_union_info_find_method.
func UnionInfoFindMethod(info *BaseInfo, name string) *FunctionInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_union_info_find_method(c_info, c_name)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnionInfoGetAlignment is a wrapper around the C function g_union_info_get_alignment.
func UnionInfoGetAlignment(info *BaseInfo) uint64 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_union_info_get_alignment(c_info)
	retGo := (uint64)(retC)

	return retGo
}

// UnionInfoGetDiscriminator is a wrapper around the C function g_union_info_get_discriminator.
func UnionInfoGetDiscriminator(info *BaseInfo, n int32) *ConstantInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_union_info_get_discriminator(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnionInfoGetDiscriminatorOffset is a wrapper around the C function g_union_info_get_discriminator_offset.
func UnionInfoGetDiscriminatorOffset(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_union_info_get_discriminator_offset(c_info)
	retGo := (int32)(retC)

	return retGo
}

// UnionInfoGetDiscriminatorType is a wrapper around the C function g_union_info_get_discriminator_type.
func UnionInfoGetDiscriminatorType(info *BaseInfo) *TypeInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_union_info_get_discriminator_type(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnionInfoGetField is a wrapper around the C function g_union_info_get_field.
func UnionInfoGetField(info *BaseInfo, n int32) *FieldInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_union_info_get_field(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnionInfoGetMethod is a wrapper around the C function g_union_info_get_method.
func UnionInfoGetMethod(info *BaseInfo, n int32) *FunctionInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.g_union_info_get_method(c_info, c_n)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnionInfoGetNFields is a wrapper around the C function g_union_info_get_n_fields.
func UnionInfoGetNFields(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_union_info_get_n_fields(c_info)
	retGo := (int32)(retC)

	return retGo
}

// UnionInfoGetNMethods is a wrapper around the C function g_union_info_get_n_methods.
func UnionInfoGetNMethods(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_union_info_get_n_methods(c_info)
	retGo := (int32)(retC)

	return retGo
}

// UnionInfoGetSize is a wrapper around the C function g_union_info_get_size.
func UnionInfoGetSize(info *BaseInfo) uint64 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_union_info_get_size(c_info)
	retGo := (uint64)(retC)

	return retGo
}

// UnionInfoIsDiscriminated is a wrapper around the C function g_union_info_is_discriminated.
func UnionInfoIsDiscriminated(info *BaseInfo) bool {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_union_info_is_discriminated(c_info)
	retGo := retC == C.TRUE

	return retGo
}

// ValueInfoGetValue is a wrapper around the C function g_value_info_get_value.
func ValueInfoGetValue(info *BaseInfo) int64 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_value_info_get_value(c_info)
	retGo := (int64)(retC)

	return retGo
}

// VfuncInfoGetAddress is a wrapper around the C function g_vfunc_info_get_address.
func VfuncInfoGetAddress(info *BaseInfo, implementorGtype gobject.Type) (uintptr, error) {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	c_implementor_gtype := (C.GType)(implementorGtype)

	var cThrowableError *C.GError

	retC := C.g_vfunc_info_get_address(c_info, c_implementor_gtype, &cThrowableError)
	retGo := (uintptr)(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// VfuncInfoGetFlags is a wrapper around the C function g_vfunc_info_get_flags.
func VfuncInfoGetFlags(info *BaseInfo) VFuncInfoFlags {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_vfunc_info_get_flags(c_info)
	retGo := (VFuncInfoFlags)(retC)

	return retGo
}

// VfuncInfoGetInvoker is a wrapper around the C function g_vfunc_info_get_invoker.
func VfuncInfoGetInvoker(info *BaseInfo) *FunctionInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_vfunc_info_get_invoker(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VfuncInfoGetOffset is a wrapper around the C function g_vfunc_info_get_offset.
func VfuncInfoGetOffset(info *BaseInfo) int32 {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_vfunc_info_get_offset(c_info)
	retGo := (int32)(retC)

	return retGo
}

// VfuncInfoGetSignal is a wrapper around the C function g_vfunc_info_get_signal.
func VfuncInfoGetSignal(info *BaseInfo) *SignalInfo {
	c_info := (*C.GIBaseInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GIBaseInfo)(info.ToC())
	}

	retC := C.g_vfunc_info_get_signal(c_info)
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_vfunc_info_invoke : unsupported parameter in_args : no type generator for Argument (GIArgument) for array param in_args

// AttributeIter is a wrapper around the C record GIAttributeIter.
type AttributeIter struct {
	native *C.GIAttributeIter
	// Private : data
	// Private : data2
	// Private : data3
	// Private : data4
}

func AttributeIterNewFromC(u unsafe.Pointer) *AttributeIter {
	c := (*C.GIAttributeIter)(u)
	if c == nil {
		return nil
	}

	g := &AttributeIter{native: c}

	return g
}

func (recv *AttributeIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AttributeIter with another AttributeIter, and returns true if they represent the same GObject.
func (recv *AttributeIter) Equals(other *AttributeIter) bool {
	return other.ToC() == recv.ToC()
}

// BaseInfo is a wrapper around the C record GIBaseInfo.
type BaseInfo struct {
	native *C.GIBaseInfo
	// Private : dummy1
	// Private : dummy2
	// Private : dummy3
	// Private : dummy4
	// Private : dummy5
	// Private : dummy6
	// Private : dummy7
	// Private : padding
}

func BaseInfoNewFromC(u unsafe.Pointer) *BaseInfo {
	c := (*C.GIBaseInfo)(u)
	if c == nil {
		return nil
	}

	g := &BaseInfo{native: c}

	return g
}

func (recv *BaseInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BaseInfo with another BaseInfo, and returns true if they represent the same GObject.
func (recv *BaseInfo) Equals(other *BaseInfo) bool {
	return other.ToC() == recv.ToC()
}

// Equal is a wrapper around the C function g_base_info_equal.
func (recv *BaseInfo) Equal(info2 *BaseInfo) bool {
	c_info2 := (*C.GIBaseInfo)(C.NULL)
	if info2 != nil {
		c_info2 = (*C.GIBaseInfo)(info2.ToC())
	}

	retC := C.g_base_info_equal((*C.GIBaseInfo)(recv.native), c_info2)
	retGo := retC == C.TRUE

	return retGo
}

// GetAttribute is a wrapper around the C function g_base_info_get_attribute.
func (recv *BaseInfo) GetAttribute(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_base_info_get_attribute((*C.GIBaseInfo)(recv.native), c_name)
	retGo := C.GoString(retC)

	return retGo
}

// GetContainer is a wrapper around the C function g_base_info_get_container.
func (recv *BaseInfo) GetContainer() *BaseInfo {
	retC := C.g_base_info_get_container((*C.GIBaseInfo)(recv.native))
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function g_base_info_get_name.
func (recv *BaseInfo) GetName() string {
	retC := C.g_base_info_get_name((*C.GIBaseInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNamespace is a wrapper around the C function g_base_info_get_namespace.
func (recv *BaseInfo) GetNamespace() string {
	retC := C.g_base_info_get_namespace((*C.GIBaseInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetType is a wrapper around the C function g_base_info_get_type.
func (recv *BaseInfo) GetType() InfoType {
	retC := C.g_base_info_get_type((*C.GIBaseInfo)(recv.native))
	retGo := (InfoType)(retC)

	return retGo
}

// GetTypelib is a wrapper around the C function g_base_info_get_typelib.
func (recv *BaseInfo) GetTypelib() *Typelib {
	retC := C.g_base_info_get_typelib((*C.GIBaseInfo)(recv.native))
	retGo := TypelibNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsDeprecated is a wrapper around the C function g_base_info_is_deprecated.
func (recv *BaseInfo) IsDeprecated() bool {
	retC := C.g_base_info_is_deprecated((*C.GIBaseInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IterateAttributes is a wrapper around the C function g_base_info_iterate_attributes.
func (recv *BaseInfo) IterateAttributes(iterator *AttributeIter) (bool, string, string) {
	c_iterator := (*C.GIAttributeIter)(C.NULL)
	if iterator != nil {
		c_iterator = (*C.GIAttributeIter)(iterator.ToC())
	}

	var c_name *C.char

	var c_value *C.char

	retC := C.g_base_info_iterate_attributes((*C.GIBaseInfo)(recv.native), c_iterator, &c_name, &c_value)
	retGo := retC == C.TRUE

	name := C.GoString(c_name)

	value := C.GoString(c_value)

	return retGo, name, value
}

// Ref is a wrapper around the C function g_base_info_ref.
func (recv *BaseInfo) Ref() *BaseInfo {
	retC := C.g_base_info_ref((*C.GIBaseInfo)(recv.native))
	retGo := BaseInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_base_info_unref.
func (recv *BaseInfo) Unref() {
	C.g_base_info_unref((*C.GIBaseInfo)(recv.native))

	return
}

// RepositoryClass is a wrapper around the C record GIRepositoryClass.
type RepositoryClass struct {
	native *C.GIRepositoryClass
	// Private : parent
}

func RepositoryClassNewFromC(u unsafe.Pointer) *RepositoryClass {
	c := (*C.GIRepositoryClass)(u)
	if c == nil {
		return nil
	}

	g := &RepositoryClass{native: c}

	return g
}

func (recv *RepositoryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RepositoryClass with another RepositoryClass, and returns true if they represent the same GObject.
func (recv *RepositoryClass) Equals(other *RepositoryClass) bool {
	return other.ToC() == recv.ToC()
}

// RepositoryPrivate is a wrapper around the C record GIRepositoryPrivate.
type RepositoryPrivate struct {
	native *C.GIRepositoryPrivate
}

func RepositoryPrivateNewFromC(u unsafe.Pointer) *RepositoryPrivate {
	c := (*C.GIRepositoryPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RepositoryPrivate{native: c}

	return g
}

func (recv *RepositoryPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RepositoryPrivate with another RepositoryPrivate, and returns true if they represent the same GObject.
func (recv *RepositoryPrivate) Equals(other *RepositoryPrivate) bool {
	return other.ToC() == recv.ToC()
}

// Typelib is a wrapper around the C record GITypelib.
type Typelib struct {
	native *C.GITypelib
}

func TypelibNewFromC(u unsafe.Pointer) *Typelib {
	c := (*C.GITypelib)(u)
	if c == nil {
		return nil
	}

	g := &Typelib{native: c}

	return g
}

func (recv *Typelib) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Typelib with another Typelib, and returns true if they represent the same GObject.
func (recv *Typelib) Equals(other *Typelib) bool {
	return other.ToC() == recv.ToC()
}

// TypelibNewFromConstMemory is a wrapper around the C function g_typelib_new_from_const_memory.
func TypelibNewFromConstMemory(memory uint8, len uint64) (*Typelib, error) {
	c_memory := (C.guint8)(memory)

	c_len := (C.gsize)(len)

	var cThrowableError *C.GError

	retC := C.g_typelib_new_from_const_memory(&c_memory, c_len, &cThrowableError)
	retGo := TypelibNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TypelibNewFromMappedFile is a wrapper around the C function g_typelib_new_from_mapped_file.
func TypelibNewFromMappedFile(mfile *glib.MappedFile) (*Typelib, error) {
	c_mfile := (*C.GMappedFile)(C.NULL)
	if mfile != nil {
		c_mfile = (*C.GMappedFile)(mfile.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_typelib_new_from_mapped_file(c_mfile, &cThrowableError)
	retGo := TypelibNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TypelibNewFromMemory is a wrapper around the C function g_typelib_new_from_memory.
func TypelibNewFromMemory(memory uint8, len uint64) (*Typelib, error) {
	c_memory := (C.guint8)(memory)

	c_len := (C.gsize)(len)

	var cThrowableError *C.GError

	retC := C.g_typelib_new_from_memory(&c_memory, c_len, &cThrowableError)
	retGo := TypelibNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Free is a wrapper around the C function g_typelib_free.
func (recv *Typelib) Free() {
	C.g_typelib_free((*C.GITypelib)(recv.native))

	return
}

// GetNamespace is a wrapper around the C function g_typelib_get_namespace.
func (recv *Typelib) GetNamespace() string {
	retC := C.g_typelib_get_namespace((*C.GITypelib)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Symbol is a wrapper around the C function g_typelib_symbol.
func (recv *Typelib) Symbol(symbolName string, symbol uintptr) bool {
	c_symbol_name := C.CString(symbolName)
	defer C.free(unsafe.Pointer(c_symbol_name))

	c_symbol := (C.gpointer)(symbol)

	retC := C.g_typelib_symbol((*C.GITypelib)(recv.native), c_symbol_name, &c_symbol)
	retGo := retC == C.TRUE

	return retGo
}

// UnresolvedInfo is a wrapper around the C record GIUnresolvedInfo.
type UnresolvedInfo struct {
	native *C.GIUnresolvedInfo
}

func UnresolvedInfoNewFromC(u unsafe.Pointer) *UnresolvedInfo {
	c := (*C.GIUnresolvedInfo)(u)
	if c == nil {
		return nil
	}

	g := &UnresolvedInfo{native: c}

	return g
}

func (recv *UnresolvedInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnresolvedInfo with another UnresolvedInfo, and returns true if they represent the same GObject.
func (recv *UnresolvedInfo) Equals(other *UnresolvedInfo) bool {
	return other.ToC() == recv.ToC()
}
