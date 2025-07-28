// This file is auto-generated, don't edit it. Thanks.
package storage_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DentryAppPropertiesValue struct {
	// example:
	//
	// property_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// property_value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// PRIVATE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s DentryAppPropertiesValue) String() string {
	return tea.Prettify(s)
}

func (s DentryAppPropertiesValue) GoString() string {
	return s.String()
}

func (s *DentryAppPropertiesValue) SetName(v string) *DentryAppPropertiesValue {
	s.Name = &v
	return s
}

func (s *DentryAppPropertiesValue) SetValue(v string) *DentryAppPropertiesValue {
	s.Value = &v
	return s
}

func (s *DentryAppPropertiesValue) SetVisibility(v string) *DentryAppPropertiesValue {
	s.Visibility = &v
	return s
}

type RoleMapValue struct {
	// example:
	//
	// MANAGER
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// MANAGER
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RoleMapValue) String() string {
	return tea.Prettify(s)
}

func (s RoleMapValue) GoString() string {
	return s.String()
}

func (s *RoleMapValue) SetId(v string) *RoleMapValue {
	s.Id = &v
	return s
}

func (s *RoleMapValue) SetName(v string) *RoleMapValue {
	s.Name = &v
	return s
}

type ResultItemsDentryAppPropertiesValue struct {
	// example:
	//
	// property_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// property_value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// PRIVATE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s ResultItemsDentryAppPropertiesValue) String() string {
	return tea.Prettify(s)
}

func (s ResultItemsDentryAppPropertiesValue) GoString() string {
	return s.String()
}

func (s *ResultItemsDentryAppPropertiesValue) SetName(v string) *ResultItemsDentryAppPropertiesValue {
	s.Name = &v
	return s
}

func (s *ResultItemsDentryAppPropertiesValue) SetValue(v string) *ResultItemsDentryAppPropertiesValue {
	s.Value = &v
	return s
}

func (s *ResultItemsDentryAppPropertiesValue) SetVisibility(v string) *ResultItemsDentryAppPropertiesValue {
	s.Visibility = &v
	return s
}

type DentriesAppPropertiesValue struct {
	// example:
	//
	// property_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// property_value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// PRIVATE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s DentriesAppPropertiesValue) String() string {
	return tea.Prettify(s)
}

func (s DentriesAppPropertiesValue) GoString() string {
	return s.String()
}

func (s *DentriesAppPropertiesValue) SetName(v string) *DentriesAppPropertiesValue {
	s.Name = &v
	return s
}

func (s *DentriesAppPropertiesValue) SetValue(v string) *DentriesAppPropertiesValue {
	s.Value = &v
	return s
}

func (s *DentriesAppPropertiesValue) SetVisibility(v string) *DentriesAppPropertiesValue {
	s.Visibility = &v
	return s
}

type AddFolderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddFolderHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddFolderHeaders) GoString() string {
	return s.String()
}

func (s *AddFolderHeaders) SetCommonHeaders(v map[string]*string) *AddFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddFolderHeaders) SetXAcsDingtalkAccessToken(v string) *AddFolderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddFolderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dentry_name
	Name   *string                 `json:"name,omitempty" xml:"name,omitempty"`
	Option *AddFolderRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFolderRequest) GoString() string {
	return s.String()
}

func (s *AddFolderRequest) SetName(v string) *AddFolderRequest {
	s.Name = &v
	return s
}

func (s *AddFolderRequest) SetOption(v *AddFolderRequestOption) *AddFolderRequest {
	s.Option = v
	return s
}

func (s *AddFolderRequest) SetUnionId(v string) *AddFolderRequest {
	s.UnionId = &v
	return s
}

type AddFolderRequestOption struct {
	AppProperties []*AddFolderRequestOptionAppProperties `json:"appProperties,omitempty" xml:"appProperties,omitempty" type:"Repeated"`
	// example:
	//
	// AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
}

func (s AddFolderRequestOption) String() string {
	return tea.Prettify(s)
}

func (s AddFolderRequestOption) GoString() string {
	return s.String()
}

func (s *AddFolderRequestOption) SetAppProperties(v []*AddFolderRequestOptionAppProperties) *AddFolderRequestOption {
	s.AppProperties = v
	return s
}

func (s *AddFolderRequestOption) SetConflictStrategy(v string) *AddFolderRequestOption {
	s.ConflictStrategy = &v
	return s
}

type AddFolderRequestOptionAppProperties struct {
	// This parameter is required.
	//
	// example:
	//
	// property_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// property_value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s AddFolderRequestOptionAppProperties) String() string {
	return tea.Prettify(s)
}

func (s AddFolderRequestOptionAppProperties) GoString() string {
	return s.String()
}

func (s *AddFolderRequestOptionAppProperties) SetName(v string) *AddFolderRequestOptionAppProperties {
	s.Name = &v
	return s
}

func (s *AddFolderRequestOptionAppProperties) SetValue(v string) *AddFolderRequestOptionAppProperties {
	s.Value = &v
	return s
}

func (s *AddFolderRequestOptionAppProperties) SetVisibility(v string) *AddFolderRequestOptionAppProperties {
	s.Visibility = &v
	return s
}

type AddFolderResponseBody struct {
	Dentry *AddFolderResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
}

func (s AddFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFolderResponseBody) GoString() string {
	return s.String()
}

func (s *AddFolderResponseBody) SetDentry(v *AddFolderResponseBodyDentry) *AddFolderResponseBody {
	s.Dentry = v
	return s
}

type AddFolderResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                `json:"path,omitempty" xml:"path,omitempty"`
	Properties *AddFolderResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AddFolderResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s AddFolderResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *AddFolderResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *AddFolderResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *AddFolderResponseBodyDentry) SetCreateTime(v string) *AddFolderResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetCreatorId(v string) *AddFolderResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetExtension(v string) *AddFolderResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetId(v string) *AddFolderResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetModifiedTime(v string) *AddFolderResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetModifierId(v string) *AddFolderResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetName(v string) *AddFolderResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetParentId(v string) *AddFolderResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetPartitionType(v string) *AddFolderResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetPath(v string) *AddFolderResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetProperties(v *AddFolderResponseBodyDentryProperties) *AddFolderResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *AddFolderResponseBodyDentry) SetSize(v int64) *AddFolderResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetSpaceId(v string) *AddFolderResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetStatus(v string) *AddFolderResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetStorageDriver(v string) *AddFolderResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetType(v string) *AddFolderResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetUuid(v string) *AddFolderResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetVersion(v int64) *AddFolderResponseBodyDentry {
	s.Version = &v
	return s
}

type AddFolderResponseBodyDentryProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s AddFolderResponseBodyDentryProperties) String() string {
	return tea.Prettify(s)
}

func (s AddFolderResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *AddFolderResponseBodyDentryProperties) SetReadOnly(v bool) *AddFolderResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

type AddFolderResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFolderResponse) GoString() string {
	return s.String()
}

func (s *AddFolderResponse) SetHeaders(v map[string]*string) *AddFolderResponse {
	s.Headers = v
	return s
}

func (s *AddFolderResponse) SetStatusCode(v int32) *AddFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFolderResponse) SetBody(v *AddFolderResponseBody) *AddFolderResponse {
	s.Body = v
	return s
}

type AddPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddPermissionHeaders) GoString() string {
	return s.String()
}

func (s *AddPermissionHeaders) SetCommonHeaders(v map[string]*string) *AddPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *AddPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddPermissionRequest struct {
	// This parameter is required.
	Members []*AddPermissionRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	Option  *AddPermissionRequestOption    `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// MANAGER
	RoleId *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPermissionRequest) GoString() string {
	return s.String()
}

func (s *AddPermissionRequest) SetMembers(v []*AddPermissionRequestMembers) *AddPermissionRequest {
	s.Members = v
	return s
}

func (s *AddPermissionRequest) SetOption(v *AddPermissionRequestOption) *AddPermissionRequest {
	s.Option = v
	return s
}

func (s *AddPermissionRequest) SetRoleId(v string) *AddPermissionRequest {
	s.RoleId = &v
	return s
}

func (s *AddPermissionRequest) SetUnionId(v string) *AddPermissionRequest {
	s.UnionId = &v
	return s
}

type AddPermissionRequestMembers struct {
	// example:
	//
	// corp_id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// member_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// USER
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AddPermissionRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s AddPermissionRequestMembers) GoString() string {
	return s.String()
}

func (s *AddPermissionRequestMembers) SetCorpId(v string) *AddPermissionRequestMembers {
	s.CorpId = &v
	return s
}

func (s *AddPermissionRequestMembers) SetId(v string) *AddPermissionRequestMembers {
	s.Id = &v
	return s
}

func (s *AddPermissionRequestMembers) SetType(v string) *AddPermissionRequestMembers {
	s.Type = &v
	return s
}

type AddPermissionRequestOption struct {
	// example:
	//
	// 3600
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
}

func (s AddPermissionRequestOption) String() string {
	return tea.Prettify(s)
}

func (s AddPermissionRequestOption) GoString() string {
	return s.String()
}

func (s *AddPermissionRequestOption) SetDuration(v int64) *AddPermissionRequestOption {
	s.Duration = &v
	return s
}

type AddPermissionResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *AddPermissionResponseBody) SetSuccess(v bool) *AddPermissionResponseBody {
	s.Success = &v
	return s
}

type AddPermissionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPermissionResponse) GoString() string {
	return s.String()
}

func (s *AddPermissionResponse) SetHeaders(v map[string]*string) *AddPermissionResponse {
	s.Headers = v
	return s
}

func (s *AddPermissionResponse) SetStatusCode(v int32) *AddPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPermissionResponse) SetBody(v *AddPermissionResponseBody) *AddPermissionResponse {
	s.Body = v
	return s
}

type AddSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceHeaders) GoString() string {
	return s.String()
}

func (s *AddSpaceHeaders) SetCommonHeaders(v map[string]*string) *AddSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *AddSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddSpaceRequest struct {
	Option *AddSpaceRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceRequest) GoString() string {
	return s.String()
}

func (s *AddSpaceRequest) SetOption(v *AddSpaceRequestOption) *AddSpaceRequest {
	s.Option = v
	return s
}

func (s *AddSpaceRequest) SetUnionId(v string) *AddSpaceRequest {
	s.UnionId = &v
	return s
}

type AddSpaceRequestOption struct {
	Capabilities *AddSpaceRequestOptionCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// example:
	//
	// space_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// USER
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	Quota     *int64  `json:"quota,omitempty" xml:"quota,omitempty"`
	// example:
	//
	// scene
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// scene_id
	SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
}

func (s AddSpaceRequestOption) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceRequestOption) GoString() string {
	return s.String()
}

func (s *AddSpaceRequestOption) SetCapabilities(v *AddSpaceRequestOptionCapabilities) *AddSpaceRequestOption {
	s.Capabilities = v
	return s
}

func (s *AddSpaceRequestOption) SetName(v string) *AddSpaceRequestOption {
	s.Name = &v
	return s
}

func (s *AddSpaceRequestOption) SetOwnerType(v string) *AddSpaceRequestOption {
	s.OwnerType = &v
	return s
}

func (s *AddSpaceRequestOption) SetQuota(v int64) *AddSpaceRequestOption {
	s.Quota = &v
	return s
}

func (s *AddSpaceRequestOption) SetScene(v string) *AddSpaceRequestOption {
	s.Scene = &v
	return s
}

func (s *AddSpaceRequestOption) SetSceneId(v string) *AddSpaceRequestOption {
	s.SceneId = &v
	return s
}

type AddSpaceRequestOptionCapabilities struct {
	// example:
	//
	// true
	CanRecordRecentFile *bool `json:"canRecordRecentFile,omitempty" xml:"canRecordRecentFile,omitempty"`
	// example:
	//
	// true
	CanRename *bool `json:"canRename,omitempty" xml:"canRename,omitempty"`
	// example:
	//
	// true
	CanSearch *bool `json:"canSearch,omitempty" xml:"canSearch,omitempty"`
}

func (s AddSpaceRequestOptionCapabilities) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceRequestOptionCapabilities) GoString() string {
	return s.String()
}

func (s *AddSpaceRequestOptionCapabilities) SetCanRecordRecentFile(v bool) *AddSpaceRequestOptionCapabilities {
	s.CanRecordRecentFile = &v
	return s
}

func (s *AddSpaceRequestOptionCapabilities) SetCanRename(v bool) *AddSpaceRequestOptionCapabilities {
	s.CanRename = &v
	return s
}

func (s *AddSpaceRequestOptionCapabilities) SetCanSearch(v bool) *AddSpaceRequestOptionCapabilities {
	s.CanSearch = &v
	return s
}

type AddSpaceResponseBody struct {
	Space *AddSpaceResponseBodySpace `json:"space,omitempty" xml:"space,omitempty" type:"Struct"`
}

func (s AddSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddSpaceResponseBody) SetSpace(v *AddSpaceResponseBodySpace) *AddSpaceResponseBody {
	s.Space = v
	return s
}

type AddSpaceResponseBodySpace struct {
	// example:
	//
	// app_id
	AppId        *string                                `json:"appId,omitempty" xml:"appId,omitempty"`
	Capabilities *AddSpaceResponseBodySpaceCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// example:
	//
	// corp_id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// space_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// space_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// owner_id
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// example:
	//
	// USER
	OwnerType  *string                                `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	Partitions []*AddSpaceResponseBodySpacePartitions `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	// example:
	//
	// 1048576
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// example:
	//
	// scene
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// scene_id
	SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1024
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s AddSpaceResponseBodySpace) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceResponseBodySpace) GoString() string {
	return s.String()
}

func (s *AddSpaceResponseBodySpace) SetAppId(v string) *AddSpaceResponseBodySpace {
	s.AppId = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetCapabilities(v *AddSpaceResponseBodySpaceCapabilities) *AddSpaceResponseBodySpace {
	s.Capabilities = v
	return s
}

func (s *AddSpaceResponseBodySpace) SetCorpId(v string) *AddSpaceResponseBodySpace {
	s.CorpId = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetCreateTime(v string) *AddSpaceResponseBodySpace {
	s.CreateTime = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetCreatorId(v string) *AddSpaceResponseBodySpace {
	s.CreatorId = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetId(v string) *AddSpaceResponseBodySpace {
	s.Id = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetModifiedTime(v string) *AddSpaceResponseBodySpace {
	s.ModifiedTime = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetModifierId(v string) *AddSpaceResponseBodySpace {
	s.ModifierId = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetName(v string) *AddSpaceResponseBodySpace {
	s.Name = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetOwnerId(v string) *AddSpaceResponseBodySpace {
	s.OwnerId = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetOwnerType(v string) *AddSpaceResponseBodySpace {
	s.OwnerType = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetPartitions(v []*AddSpaceResponseBodySpacePartitions) *AddSpaceResponseBodySpace {
	s.Partitions = v
	return s
}

func (s *AddSpaceResponseBodySpace) SetQuota(v int64) *AddSpaceResponseBodySpace {
	s.Quota = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetScene(v string) *AddSpaceResponseBodySpace {
	s.Scene = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetSceneId(v string) *AddSpaceResponseBodySpace {
	s.SceneId = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetStatus(v string) *AddSpaceResponseBodySpace {
	s.Status = &v
	return s
}

func (s *AddSpaceResponseBodySpace) SetUsedQuota(v int64) *AddSpaceResponseBodySpace {
	s.UsedQuota = &v
	return s
}

type AddSpaceResponseBodySpaceCapabilities struct {
	// example:
	//
	// true
	CanRecordRecentFile *bool `json:"canRecordRecentFile,omitempty" xml:"canRecordRecentFile,omitempty"`
	// example:
	//
	// true
	CanRename *bool `json:"canRename,omitempty" xml:"canRename,omitempty"`
	// example:
	//
	// true
	CanSearch *bool `json:"canSearch,omitempty" xml:"canSearch,omitempty"`
}

func (s AddSpaceResponseBodySpaceCapabilities) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceResponseBodySpaceCapabilities) GoString() string {
	return s.String()
}

func (s *AddSpaceResponseBodySpaceCapabilities) SetCanRecordRecentFile(v bool) *AddSpaceResponseBodySpaceCapabilities {
	s.CanRecordRecentFile = &v
	return s
}

func (s *AddSpaceResponseBodySpaceCapabilities) SetCanRename(v bool) *AddSpaceResponseBodySpaceCapabilities {
	s.CanRename = &v
	return s
}

func (s *AddSpaceResponseBodySpaceCapabilities) SetCanSearch(v bool) *AddSpaceResponseBodySpaceCapabilities {
	s.CanSearch = &v
	return s
}

type AddSpaceResponseBodySpacePartitions struct {
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string                                   `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	Quota         *AddSpaceResponseBodySpacePartitionsQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s AddSpaceResponseBodySpacePartitions) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceResponseBodySpacePartitions) GoString() string {
	return s.String()
}

func (s *AddSpaceResponseBodySpacePartitions) SetPartitionType(v string) *AddSpaceResponseBodySpacePartitions {
	s.PartitionType = &v
	return s
}

func (s *AddSpaceResponseBodySpacePartitions) SetQuota(v *AddSpaceResponseBodySpacePartitionsQuota) *AddSpaceResponseBodySpacePartitions {
	s.Quota = v
	return s
}

type AddSpaceResponseBodySpacePartitionsQuota struct {
	// example:
	//
	// 10000
	Max *int64 `json:"max,omitempty" xml:"max,omitempty"`
	// example:
	//
	// 1000
	Reserved *int64 `json:"reserved,omitempty" xml:"reserved,omitempty"`
	// example:
	//
	// SHARE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1024
	Used *int64 `json:"used,omitempty" xml:"used,omitempty"`
}

func (s AddSpaceResponseBodySpacePartitionsQuota) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceResponseBodySpacePartitionsQuota) GoString() string {
	return s.String()
}

func (s *AddSpaceResponseBodySpacePartitionsQuota) SetMax(v int64) *AddSpaceResponseBodySpacePartitionsQuota {
	s.Max = &v
	return s
}

func (s *AddSpaceResponseBodySpacePartitionsQuota) SetReserved(v int64) *AddSpaceResponseBodySpacePartitionsQuota {
	s.Reserved = &v
	return s
}

func (s *AddSpaceResponseBodySpacePartitionsQuota) SetType(v string) *AddSpaceResponseBodySpacePartitionsQuota {
	s.Type = &v
	return s
}

func (s *AddSpaceResponseBodySpacePartitionsQuota) SetUsed(v int64) *AddSpaceResponseBodySpacePartitionsQuota {
	s.Used = &v
	return s
}

type AddSpaceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceResponse) GoString() string {
	return s.String()
}

func (s *AddSpaceResponse) SetHeaders(v map[string]*string) *AddSpaceResponse {
	s.Headers = v
	return s
}

func (s *AddSpaceResponse) SetStatusCode(v int32) *AddSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSpaceResponse) SetBody(v *AddSpaceResponseBody) *AddSpaceResponse {
	s.Body = v
	return s
}

type BatchQueryRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryRolesHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryRolesHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryRolesHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryRolesRequest struct {
	// This parameter is required.
	DentryIdList []*string `json:"dentryIdList,omitempty" xml:"dentryIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s BatchQueryRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryRolesRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryRolesRequest) SetDentryIdList(v []*string) *BatchQueryRolesRequest {
	s.DentryIdList = v
	return s
}

func (s *BatchQueryRolesRequest) SetUnionId(v string) *BatchQueryRolesRequest {
	s.UnionId = &v
	return s
}

type BatchQueryRolesShrinkRequest struct {
	// This parameter is required.
	DentryIdListShrink *string `json:"dentryIdList,omitempty" xml:"dentryIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s BatchQueryRolesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryRolesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryRolesShrinkRequest) SetDentryIdListShrink(v string) *BatchQueryRolesShrinkRequest {
	s.DentryIdListShrink = &v
	return s
}

func (s *BatchQueryRolesShrinkRequest) SetUnionId(v string) *BatchQueryRolesShrinkRequest {
	s.UnionId = &v
	return s
}

type BatchQueryRolesResponseBody struct {
	RoleMap map[string][]*RoleMapValue `json:"roleMap,omitempty" xml:"roleMap,omitempty"`
}

func (s BatchQueryRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryRolesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryRolesResponseBody) SetRoleMap(v map[string][]*RoleMapValue) *BatchQueryRolesResponseBody {
	s.RoleMap = v
	return s
}

type BatchQueryRolesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryRolesResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryRolesResponse) SetHeaders(v map[string]*string) *BatchQueryRolesResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryRolesResponse) SetStatusCode(v int32) *BatchQueryRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryRolesResponse) SetBody(v *BatchQueryRolesResponseBody) *BatchQueryRolesResponse {
	s.Body = v
	return s
}

type ClearRecycleBinHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ClearRecycleBinHeaders) String() string {
	return tea.Prettify(s)
}

func (s ClearRecycleBinHeaders) GoString() string {
	return s.String()
}

func (s *ClearRecycleBinHeaders) SetCommonHeaders(v map[string]*string) *ClearRecycleBinHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ClearRecycleBinHeaders) SetXAcsDingtalkAccessToken(v string) *ClearRecycleBinHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ClearRecycleBinRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ClearRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *ClearRecycleBinRequest) SetUnionId(v string) *ClearRecycleBinRequest {
	s.UnionId = &v
	return s
}

type ClearRecycleBinResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ClearRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *ClearRecycleBinResponseBody) SetSuccess(v bool) *ClearRecycleBinResponseBody {
	s.Success = &v
	return s
}

type ClearRecycleBinResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *ClearRecycleBinResponse) SetHeaders(v map[string]*string) *ClearRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *ClearRecycleBinResponse) SetStatusCode(v int32) *ClearRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearRecycleBinResponse) SetBody(v *ClearRecycleBinResponseBody) *ClearRecycleBinResponse {
	s.Body = v
	return s
}

type CommitFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CommitFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s CommitFileHeaders) GoString() string {
	return s.String()
}

func (s *CommitFileHeaders) SetCommonHeaders(v map[string]*string) *CommitFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CommitFileHeaders) SetXAcsDingtalkAccessToken(v string) *CommitFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CommitFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dentry_name
	Name   *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Option *CommitFileRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// example:
	//
	// dentry_id
	OverwriteDentryId *string `json:"overwriteDentryId,omitempty" xml:"overwriteDentryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// upload_key
	UploadKey *string `json:"uploadKey,omitempty" xml:"uploadKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CommitFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CommitFileRequest) GoString() string {
	return s.String()
}

func (s *CommitFileRequest) SetName(v string) *CommitFileRequest {
	s.Name = &v
	return s
}

func (s *CommitFileRequest) SetOption(v *CommitFileRequestOption) *CommitFileRequest {
	s.Option = v
	return s
}

func (s *CommitFileRequest) SetOverwriteDentryId(v string) *CommitFileRequest {
	s.OverwriteDentryId = &v
	return s
}

func (s *CommitFileRequest) SetParentId(v string) *CommitFileRequest {
	s.ParentId = &v
	return s
}

func (s *CommitFileRequest) SetUploadKey(v string) *CommitFileRequest {
	s.UploadKey = &v
	return s
}

func (s *CommitFileRequest) SetUnionId(v string) *CommitFileRequest {
	s.UnionId = &v
	return s
}

type CommitFileRequestOption struct {
	AppProperties []*CommitFileRequestOptionAppProperties `json:"appProperties,omitempty" xml:"appProperties,omitempty" type:"Repeated"`
	// example:
	//
	// AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s CommitFileRequestOption) String() string {
	return tea.Prettify(s)
}

func (s CommitFileRequestOption) GoString() string {
	return s.String()
}

func (s *CommitFileRequestOption) SetAppProperties(v []*CommitFileRequestOptionAppProperties) *CommitFileRequestOption {
	s.AppProperties = v
	return s
}

func (s *CommitFileRequestOption) SetConflictStrategy(v string) *CommitFileRequestOption {
	s.ConflictStrategy = &v
	return s
}

func (s *CommitFileRequestOption) SetSize(v int64) *CommitFileRequestOption {
	s.Size = &v
	return s
}

type CommitFileRequestOptionAppProperties struct {
	// This parameter is required.
	//
	// example:
	//
	// property_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// property_value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s CommitFileRequestOptionAppProperties) String() string {
	return tea.Prettify(s)
}

func (s CommitFileRequestOptionAppProperties) GoString() string {
	return s.String()
}

func (s *CommitFileRequestOptionAppProperties) SetName(v string) *CommitFileRequestOptionAppProperties {
	s.Name = &v
	return s
}

func (s *CommitFileRequestOptionAppProperties) SetValue(v string) *CommitFileRequestOptionAppProperties {
	s.Value = &v
	return s
}

func (s *CommitFileRequestOptionAppProperties) SetVisibility(v string) *CommitFileRequestOptionAppProperties {
	s.Visibility = &v
	return s
}

type CommitFileResponseBody struct {
	Dentry *CommitFileResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
}

func (s CommitFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponseBody) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBody) SetDentry(v *CommitFileResponseBodyDentry) *CommitFileResponseBody {
	s.Dentry = v
	return s
}

type CommitFileResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// example:
	//
	// DOCUMENT
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                 `json:"path,omitempty" xml:"path,omitempty"`
	Properties *CommitFileResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string                                `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	Thumbnail     *CommitFileResponseBodyDentryThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
	// example:
	//
	// FILE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CommitFileResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *CommitFileResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *CommitFileResponseBodyDentry) SetCategory(v string) *CommitFileResponseBodyDentry {
	s.Category = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetCreateTime(v string) *CommitFileResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetCreatorId(v string) *CommitFileResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetExtension(v string) *CommitFileResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetId(v string) *CommitFileResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetModifiedTime(v string) *CommitFileResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetModifierId(v string) *CommitFileResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetName(v string) *CommitFileResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetParentId(v string) *CommitFileResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetPartitionType(v string) *CommitFileResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetPath(v string) *CommitFileResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetProperties(v *CommitFileResponseBodyDentryProperties) *CommitFileResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *CommitFileResponseBodyDentry) SetSize(v int64) *CommitFileResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetSpaceId(v string) *CommitFileResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetStatus(v string) *CommitFileResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetStorageDriver(v string) *CommitFileResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetThumbnail(v *CommitFileResponseBodyDentryThumbnail) *CommitFileResponseBodyDentry {
	s.Thumbnail = v
	return s
}

func (s *CommitFileResponseBodyDentry) SetType(v string) *CommitFileResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetUuid(v string) *CommitFileResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetVersion(v int64) *CommitFileResponseBodyDentry {
	s.Version = &v
	return s
}

type CommitFileResponseBodyDentryProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s CommitFileResponseBodyDentryProperties) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBodyDentryProperties) SetReadOnly(v bool) *CommitFileResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

type CommitFileResponseBodyDentryThumbnail struct {
	// example:
	//
	// 64
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// 64
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s CommitFileResponseBodyDentryThumbnail) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponseBodyDentryThumbnail) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBodyDentryThumbnail) SetHeight(v int32) *CommitFileResponseBodyDentryThumbnail {
	s.Height = &v
	return s
}

func (s *CommitFileResponseBodyDentryThumbnail) SetUrl(v string) *CommitFileResponseBodyDentryThumbnail {
	s.Url = &v
	return s
}

func (s *CommitFileResponseBodyDentryThumbnail) SetWidth(v int32) *CommitFileResponseBodyDentryThumbnail {
	s.Width = &v
	return s
}

type CommitFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommitFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommitFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponse) GoString() string {
	return s.String()
}

func (s *CommitFileResponse) SetHeaders(v map[string]*string) *CommitFileResponse {
	s.Headers = v
	return s
}

func (s *CommitFileResponse) SetStatusCode(v int32) *CommitFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitFileResponse) SetBody(v *CommitFileResponseBody) *CommitFileResponse {
	s.Body = v
	return s
}

type CopyDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CopyDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s CopyDentriesHeaders) GoString() string {
	return s.String()
}

func (s *CopyDentriesHeaders) SetCommonHeaders(v map[string]*string) *CopyDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *CopyDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CopyDentriesRequest struct {
	// This parameter is required.
	DentryIds []*string                  `json:"dentryIds,omitempty" xml:"dentryIds,omitempty" type:"Repeated"`
	Option    *CopyDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// target_folder_id
	TargetFolderId *string `json:"targetFolderId,omitempty" xml:"targetFolderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// target_space_id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CopyDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyDentriesRequest) GoString() string {
	return s.String()
}

func (s *CopyDentriesRequest) SetDentryIds(v []*string) *CopyDentriesRequest {
	s.DentryIds = v
	return s
}

func (s *CopyDentriesRequest) SetOption(v *CopyDentriesRequestOption) *CopyDentriesRequest {
	s.Option = v
	return s
}

func (s *CopyDentriesRequest) SetTargetFolderId(v string) *CopyDentriesRequest {
	s.TargetFolderId = &v
	return s
}

func (s *CopyDentriesRequest) SetTargetSpaceId(v string) *CopyDentriesRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *CopyDentriesRequest) SetUnionId(v string) *CopyDentriesRequest {
	s.UnionId = &v
	return s
}

type CopyDentriesRequestOption struct {
	// example:
	//
	// AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
}

func (s CopyDentriesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s CopyDentriesRequestOption) GoString() string {
	return s.String()
}

func (s *CopyDentriesRequestOption) SetConflictStrategy(v string) *CopyDentriesRequestOption {
	s.ConflictStrategy = &v
	return s
}

type CopyDentriesResponseBody struct {
	ResultItems []*CopyDentriesResponseBodyResultItems `json:"resultItems,omitempty" xml:"resultItems,omitempty" type:"Repeated"`
}

func (s CopyDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *CopyDentriesResponseBody) SetResultItems(v []*CopyDentriesResponseBodyResultItems) *CopyDentriesResponseBody {
	s.ResultItems = v
	return s
}

type CopyDentriesResponseBodyResultItems struct {
	// example:
	//
	// true
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// example:
	//
	// dentry_id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// permissionDenied
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// target_dentry_id
	TargetDentryId *string `json:"targetDentryId,omitempty" xml:"targetDentryId,omitempty"`
	// example:
	//
	// target_space_id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// example:
	//
	// task_id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CopyDentriesResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s CopyDentriesResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *CopyDentriesResponseBodyResultItems) SetAsync(v bool) *CopyDentriesResponseBodyResultItems {
	s.Async = &v
	return s
}

func (s *CopyDentriesResponseBodyResultItems) SetDentryId(v string) *CopyDentriesResponseBodyResultItems {
	s.DentryId = &v
	return s
}

func (s *CopyDentriesResponseBodyResultItems) SetErrorCode(v string) *CopyDentriesResponseBodyResultItems {
	s.ErrorCode = &v
	return s
}

func (s *CopyDentriesResponseBodyResultItems) SetSpaceId(v string) *CopyDentriesResponseBodyResultItems {
	s.SpaceId = &v
	return s
}

func (s *CopyDentriesResponseBodyResultItems) SetSuccess(v bool) *CopyDentriesResponseBodyResultItems {
	s.Success = &v
	return s
}

func (s *CopyDentriesResponseBodyResultItems) SetTargetDentryId(v string) *CopyDentriesResponseBodyResultItems {
	s.TargetDentryId = &v
	return s
}

func (s *CopyDentriesResponseBodyResultItems) SetTargetSpaceId(v string) *CopyDentriesResponseBodyResultItems {
	s.TargetSpaceId = &v
	return s
}

func (s *CopyDentriesResponseBodyResultItems) SetTaskId(v string) *CopyDentriesResponseBodyResultItems {
	s.TaskId = &v
	return s
}

type CopyDentriesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyDentriesResponse) GoString() string {
	return s.String()
}

func (s *CopyDentriesResponse) SetHeaders(v map[string]*string) *CopyDentriesResponse {
	s.Headers = v
	return s
}

func (s *CopyDentriesResponse) SetStatusCode(v int32) *CopyDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDentriesResponse) SetBody(v *CopyDentriesResponseBody) *CopyDentriesResponse {
	s.Body = v
	return s
}

type CopyDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CopyDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryHeaders) GoString() string {
	return s.String()
}

func (s *CopyDentryHeaders) SetCommonHeaders(v map[string]*string) *CopyDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyDentryHeaders) SetXAcsDingtalkAccessToken(v string) *CopyDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CopyDentryRequest struct {
	Option *CopyDentryRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// target_folder_id
	TargetFolderId *string `json:"targetFolderId,omitempty" xml:"targetFolderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// target_space_id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CopyDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryRequest) GoString() string {
	return s.String()
}

func (s *CopyDentryRequest) SetOption(v *CopyDentryRequestOption) *CopyDentryRequest {
	s.Option = v
	return s
}

func (s *CopyDentryRequest) SetTargetFolderId(v string) *CopyDentryRequest {
	s.TargetFolderId = &v
	return s
}

func (s *CopyDentryRequest) SetTargetSpaceId(v string) *CopyDentryRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *CopyDentryRequest) SetUnionId(v string) *CopyDentryRequest {
	s.UnionId = &v
	return s
}

type CopyDentryRequestOption struct {
	// example:
	//
	// AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
}

func (s CopyDentryRequestOption) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryRequestOption) GoString() string {
	return s.String()
}

func (s *CopyDentryRequestOption) SetConflictStrategy(v string) *CopyDentryRequestOption {
	s.ConflictStrategy = &v
	return s
}

type CopyDentryResponseBody struct {
	// example:
	//
	// true
	Async  *bool                         `json:"async,omitempty" xml:"async,omitempty"`
	Dentry *CopyDentryResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
	// example:
	//
	// task_id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CopyDentryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryResponseBody) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBody) SetAsync(v bool) *CopyDentryResponseBody {
	s.Async = &v
	return s
}

func (s *CopyDentryResponseBody) SetDentry(v *CopyDentryResponseBodyDentry) *CopyDentryResponseBody {
	s.Dentry = v
	return s
}

func (s *CopyDentryResponseBody) SetTaskId(v string) *CopyDentryResponseBody {
	s.TaskId = &v
	return s
}

type CopyDentryResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                 `json:"path,omitempty" xml:"path,omitempty"`
	Properties *CopyDentryResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CopyDentryResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *CopyDentryResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetCreateTime(v string) *CopyDentryResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetCreatorId(v string) *CopyDentryResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetExtension(v string) *CopyDentryResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetId(v string) *CopyDentryResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetModifiedTime(v string) *CopyDentryResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetModifierId(v string) *CopyDentryResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetName(v string) *CopyDentryResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetParentId(v string) *CopyDentryResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetPartitionType(v string) *CopyDentryResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetPath(v string) *CopyDentryResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetProperties(v *CopyDentryResponseBodyDentryProperties) *CopyDentryResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetSize(v int64) *CopyDentryResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetSpaceId(v string) *CopyDentryResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetStatus(v string) *CopyDentryResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetStorageDriver(v string) *CopyDentryResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetType(v string) *CopyDentryResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetUuid(v string) *CopyDentryResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *CopyDentryResponseBodyDentry) SetVersion(v int64) *CopyDentryResponseBodyDentry {
	s.Version = &v
	return s
}

type CopyDentryResponseBodyDentryProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s CopyDentryResponseBodyDentryProperties) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodyDentryProperties) SetReadOnly(v bool) *CopyDentryResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

type CopyDentryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyDentryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryResponse) GoString() string {
	return s.String()
}

func (s *CopyDentryResponse) SetHeaders(v map[string]*string) *CopyDentryResponse {
	s.Headers = v
	return s
}

func (s *CopyDentryResponse) SetStatusCode(v int32) *CopyDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDentryResponse) SetBody(v *CopyDentryResponseBody) *CopyDentryResponse {
	s.Body = v
	return s
}

type DeleteDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentriesHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDentriesHeaders) SetCommonHeaders(v map[string]*string) *DeleteDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDentriesRequest struct {
	// This parameter is required.
	DentryIds []*string                    `json:"dentryIds,omitempty" xml:"dentryIds,omitempty" type:"Repeated"`
	Option    *DeleteDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentriesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDentriesRequest) SetDentryIds(v []*string) *DeleteDentriesRequest {
	s.DentryIds = v
	return s
}

func (s *DeleteDentriesRequest) SetOption(v *DeleteDentriesRequestOption) *DeleteDentriesRequest {
	s.Option = v
	return s
}

func (s *DeleteDentriesRequest) SetUnionId(v string) *DeleteDentriesRequest {
	s.UnionId = &v
	return s
}

type DeleteDentriesRequestOption struct {
	// example:
	//
	// true
	ToRecycleBin *bool `json:"toRecycleBin,omitempty" xml:"toRecycleBin,omitempty"`
}

func (s DeleteDentriesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentriesRequestOption) GoString() string {
	return s.String()
}

func (s *DeleteDentriesRequestOption) SetToRecycleBin(v bool) *DeleteDentriesRequestOption {
	s.ToRecycleBin = &v
	return s
}

type DeleteDentriesResponseBody struct {
	ResultItems []*DeleteDentriesResponseBodyResultItems `json:"resultItems,omitempty" xml:"resultItems,omitempty" type:"Repeated"`
}

func (s DeleteDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDentriesResponseBody) SetResultItems(v []*DeleteDentriesResponseBodyResultItems) *DeleteDentriesResponseBody {
	s.ResultItems = v
	return s
}

type DeleteDentriesResponseBodyResultItems struct {
	// example:
	//
	// true
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// example:
	//
	// dentry_id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// permissionDenied
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// task_id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteDentriesResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentriesResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *DeleteDentriesResponseBodyResultItems) SetAsync(v bool) *DeleteDentriesResponseBodyResultItems {
	s.Async = &v
	return s
}

func (s *DeleteDentriesResponseBodyResultItems) SetDentryId(v string) *DeleteDentriesResponseBodyResultItems {
	s.DentryId = &v
	return s
}

func (s *DeleteDentriesResponseBodyResultItems) SetErrorCode(v string) *DeleteDentriesResponseBodyResultItems {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDentriesResponseBodyResultItems) SetSpaceId(v string) *DeleteDentriesResponseBodyResultItems {
	s.SpaceId = &v
	return s
}

func (s *DeleteDentriesResponseBodyResultItems) SetSuccess(v bool) *DeleteDentriesResponseBodyResultItems {
	s.Success = &v
	return s
}

func (s *DeleteDentriesResponseBodyResultItems) SetTaskId(v string) *DeleteDentriesResponseBodyResultItems {
	s.TaskId = &v
	return s
}

type DeleteDentriesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDentriesResponse) SetHeaders(v map[string]*string) *DeleteDentriesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDentriesResponse) SetStatusCode(v int32) *DeleteDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDentriesResponse) SetBody(v *DeleteDentriesResponseBody) *DeleteDentriesResponse {
	s.Body = v
	return s
}

type DeleteDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentryHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDentryHeaders) SetCommonHeaders(v map[string]*string) *DeleteDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDentryHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDentryRequest struct {
	// example:
	//
	// true
	ToRecycleBin *bool `json:"toRecycleBin,omitempty" xml:"toRecycleBin,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDentryRequest) SetToRecycleBin(v bool) *DeleteDentryRequest {
	s.ToRecycleBin = &v
	return s
}

func (s *DeleteDentryRequest) SetUnionId(v string) *DeleteDentryRequest {
	s.UnionId = &v
	return s
}

type DeleteDentryResponseBody struct {
	// example:
	//
	// true
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// example:
	//
	// task_id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteDentryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDentryResponseBody) SetAsync(v bool) *DeleteDentryResponseBody {
	s.Async = &v
	return s
}

func (s *DeleteDentryResponseBody) SetTaskId(v string) *DeleteDentryResponseBody {
	s.TaskId = &v
	return s
}

type DeleteDentryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDentryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentryResponse) GoString() string {
	return s.String()
}

func (s *DeleteDentryResponse) SetHeaders(v map[string]*string) *DeleteDentryResponse {
	s.Headers = v
	return s
}

func (s *DeleteDentryResponse) SetStatusCode(v int32) *DeleteDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDentryResponse) SetBody(v *DeleteDentryResponseBody) *DeleteDentryResponse {
	s.Body = v
	return s
}

type DeleteDentryAppPropertiesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDentryAppPropertiesHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentryAppPropertiesHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDentryAppPropertiesHeaders) SetCommonHeaders(v map[string]*string) *DeleteDentryAppPropertiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDentryAppPropertiesHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDentryAppPropertiesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDentryAppPropertiesRequest struct {
	// This parameter is required.
	PropertyNames []*string `json:"propertyNames,omitempty" xml:"propertyNames,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteDentryAppPropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentryAppPropertiesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDentryAppPropertiesRequest) SetPropertyNames(v []*string) *DeleteDentryAppPropertiesRequest {
	s.PropertyNames = v
	return s
}

func (s *DeleteDentryAppPropertiesRequest) SetUnionId(v string) *DeleteDentryAppPropertiesRequest {
	s.UnionId = &v
	return s
}

type DeleteDentryAppPropertiesResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDentryAppPropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentryAppPropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDentryAppPropertiesResponseBody) SetSuccess(v bool) *DeleteDentryAppPropertiesResponseBody {
	s.Success = &v
	return s
}

type DeleteDentryAppPropertiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDentryAppPropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDentryAppPropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDentryAppPropertiesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDentryAppPropertiesResponse) SetHeaders(v map[string]*string) *DeleteDentryAppPropertiesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDentryAppPropertiesResponse) SetStatusCode(v int32) *DeleteDentryAppPropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDentryAppPropertiesResponse) SetBody(v *DeleteDentryAppPropertiesResponseBody) *DeleteDentryAppPropertiesResponse {
	s.Body = v
	return s
}

type DeletePermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeletePermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionHeaders) GoString() string {
	return s.String()
}

func (s *DeletePermissionHeaders) SetCommonHeaders(v map[string]*string) *DeletePermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeletePermissionHeaders) SetXAcsDingtalkAccessToken(v string) *DeletePermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeletePermissionRequest struct {
	// This parameter is required.
	Members []*DeletePermissionRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// MANAGER
	RoleId *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeletePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionRequest) GoString() string {
	return s.String()
}

func (s *DeletePermissionRequest) SetMembers(v []*DeletePermissionRequestMembers) *DeletePermissionRequest {
	s.Members = v
	return s
}

func (s *DeletePermissionRequest) SetRoleId(v string) *DeletePermissionRequest {
	s.RoleId = &v
	return s
}

func (s *DeletePermissionRequest) SetUnionId(v string) *DeletePermissionRequest {
	s.UnionId = &v
	return s
}

type DeletePermissionRequestMembers struct {
	// example:
	//
	// corp_id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// member_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// USER
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeletePermissionRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionRequestMembers) GoString() string {
	return s.String()
}

func (s *DeletePermissionRequestMembers) SetCorpId(v string) *DeletePermissionRequestMembers {
	s.CorpId = &v
	return s
}

func (s *DeletePermissionRequestMembers) SetId(v string) *DeletePermissionRequestMembers {
	s.Id = &v
	return s
}

func (s *DeletePermissionRequestMembers) SetType(v string) *DeletePermissionRequestMembers {
	s.Type = &v
	return s
}

type DeletePermissionResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeletePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponseBody) SetSuccess(v bool) *DeletePermissionResponseBody {
	s.Success = &v
	return s
}

type DeletePermissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponse) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponse) SetHeaders(v map[string]*string) *DeletePermissionResponse {
	s.Headers = v
	return s
}

func (s *DeletePermissionResponse) SetStatusCode(v int32) *DeletePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePermissionResponse) SetBody(v *DeletePermissionResponseBody) *DeletePermissionResponse {
	s.Body = v
	return s
}

type DeleteRecycleItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteRecycleItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleItemHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRecycleItemHeaders) SetCommonHeaders(v map[string]*string) *DeleteRecycleItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRecycleItemHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteRecycleItemHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteRecycleItemRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteRecycleItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecycleItemRequest) SetUnionId(v string) *DeleteRecycleItemRequest {
	s.UnionId = &v
	return s
}

type DeleteRecycleItemResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRecycleItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecycleItemResponseBody) SetSuccess(v bool) *DeleteRecycleItemResponseBody {
	s.Success = &v
	return s
}

type DeleteRecycleItemResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecycleItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecycleItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecycleItemResponse) SetHeaders(v map[string]*string) *DeleteRecycleItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecycleItemResponse) SetStatusCode(v int32) *DeleteRecycleItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecycleItemResponse) SetBody(v *DeleteRecycleItemResponseBody) *DeleteRecycleItemResponse {
	s.Body = v
	return s
}

type DeleteRecycleItemsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteRecycleItemsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleItemsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRecycleItemsHeaders) SetCommonHeaders(v map[string]*string) *DeleteRecycleItemsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRecycleItemsHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteRecycleItemsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteRecycleItemsRequest struct {
	// This parameter is required.
	RecycleItemIds []*string `json:"recycleItemIds,omitempty" xml:"recycleItemIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteRecycleItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleItemsRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecycleItemsRequest) SetRecycleItemIds(v []*string) *DeleteRecycleItemsRequest {
	s.RecycleItemIds = v
	return s
}

func (s *DeleteRecycleItemsRequest) SetUnionId(v string) *DeleteRecycleItemsRequest {
	s.UnionId = &v
	return s
}

type DeleteRecycleItemsResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRecycleItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecycleItemsResponseBody) SetSuccess(v bool) *DeleteRecycleItemsResponseBody {
	s.Success = &v
	return s
}

type DeleteRecycleItemsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecycleItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecycleItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleItemsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecycleItemsResponse) SetHeaders(v map[string]*string) *DeleteRecycleItemsResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecycleItemsResponse) SetStatusCode(v int32) *DeleteRecycleItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecycleItemsResponse) SetBody(v *DeleteRecycleItemsResponseBody) *DeleteRecycleItemsResponse {
	s.Body = v
	return s
}

type GetCurrentAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCurrentAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentAppHeaders) GoString() string {
	return s.String()
}

func (s *GetCurrentAppHeaders) SetCommonHeaders(v map[string]*string) *GetCurrentAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCurrentAppHeaders) SetXAcsDingtalkAccessToken(v string) *GetCurrentAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCurrentAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetCurrentAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentAppRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentAppRequest) SetUnionId(v string) *GetCurrentAppRequest {
	s.UnionId = &v
	return s
}

type GetCurrentAppResponseBody struct {
	App *GetCurrentAppResponseBodyApp `json:"app,omitempty" xml:"app,omitempty" type:"Struct"`
}

func (s GetCurrentAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetCurrentAppResponseBody) SetApp(v *GetCurrentAppResponseBodyApp) *GetCurrentAppResponseBody {
	s.App = v
	return s
}

type GetCurrentAppResponseBodyApp struct {
	// example:
	//
	// app_id
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// corp_id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// app_name
	Name       *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	Partitions []*GetCurrentAppResponseBodyAppPartitions `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
}

func (s GetCurrentAppResponseBodyApp) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentAppResponseBodyApp) GoString() string {
	return s.String()
}

func (s *GetCurrentAppResponseBodyApp) SetAppId(v string) *GetCurrentAppResponseBodyApp {
	s.AppId = &v
	return s
}

func (s *GetCurrentAppResponseBodyApp) SetCorpId(v string) *GetCurrentAppResponseBodyApp {
	s.CorpId = &v
	return s
}

func (s *GetCurrentAppResponseBodyApp) SetCreateTime(v string) *GetCurrentAppResponseBodyApp {
	s.CreateTime = &v
	return s
}

func (s *GetCurrentAppResponseBodyApp) SetModifiedTime(v string) *GetCurrentAppResponseBodyApp {
	s.ModifiedTime = &v
	return s
}

func (s *GetCurrentAppResponseBodyApp) SetName(v string) *GetCurrentAppResponseBodyApp {
	s.Name = &v
	return s
}

func (s *GetCurrentAppResponseBodyApp) SetPartitions(v []*GetCurrentAppResponseBodyAppPartitions) *GetCurrentAppResponseBodyApp {
	s.Partitions = v
	return s
}

type GetCurrentAppResponseBodyAppPartitions struct {
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string                                      `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	Quota         *GetCurrentAppResponseBodyAppPartitionsQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetCurrentAppResponseBodyAppPartitions) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentAppResponseBodyAppPartitions) GoString() string {
	return s.String()
}

func (s *GetCurrentAppResponseBodyAppPartitions) SetPartitionType(v string) *GetCurrentAppResponseBodyAppPartitions {
	s.PartitionType = &v
	return s
}

func (s *GetCurrentAppResponseBodyAppPartitions) SetQuota(v *GetCurrentAppResponseBodyAppPartitionsQuota) *GetCurrentAppResponseBodyAppPartitions {
	s.Quota = v
	return s
}

type GetCurrentAppResponseBodyAppPartitionsQuota struct {
	// example:
	//
	// 10000
	Max *int64 `json:"max,omitempty" xml:"max,omitempty"`
	// example:
	//
	// 1000
	Reserved *int64 `json:"reserved,omitempty" xml:"reserved,omitempty"`
	// example:
	//
	// SHARE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1024
	Used *int64 `json:"used,omitempty" xml:"used,omitempty"`
}

func (s GetCurrentAppResponseBodyAppPartitionsQuota) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentAppResponseBodyAppPartitionsQuota) GoString() string {
	return s.String()
}

func (s *GetCurrentAppResponseBodyAppPartitionsQuota) SetMax(v int64) *GetCurrentAppResponseBodyAppPartitionsQuota {
	s.Max = &v
	return s
}

func (s *GetCurrentAppResponseBodyAppPartitionsQuota) SetReserved(v int64) *GetCurrentAppResponseBodyAppPartitionsQuota {
	s.Reserved = &v
	return s
}

func (s *GetCurrentAppResponseBodyAppPartitionsQuota) SetType(v string) *GetCurrentAppResponseBodyAppPartitionsQuota {
	s.Type = &v
	return s
}

func (s *GetCurrentAppResponseBodyAppPartitionsQuota) SetUsed(v int64) *GetCurrentAppResponseBodyAppPartitionsQuota {
	s.Used = &v
	return s
}

type GetCurrentAppResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCurrentAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCurrentAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCurrentAppResponse) GoString() string {
	return s.String()
}

func (s *GetCurrentAppResponse) SetHeaders(v map[string]*string) *GetCurrentAppResponse {
	s.Headers = v
	return s
}

func (s *GetCurrentAppResponse) SetStatusCode(v int32) *GetCurrentAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCurrentAppResponse) SetBody(v *GetCurrentAppResponseBody) *GetCurrentAppResponse {
	s.Body = v
	return s
}

type GetDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesHeaders) GoString() string {
	return s.String()
}

func (s *GetDentriesHeaders) SetCommonHeaders(v map[string]*string) *GetDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *GetDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDentriesRequest struct {
	// This parameter is required.
	DentryIds []*string                 `json:"dentryIds,omitempty" xml:"dentryIds,omitempty" type:"Repeated"`
	Option    *GetDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesRequest) GoString() string {
	return s.String()
}

func (s *GetDentriesRequest) SetDentryIds(v []*string) *GetDentriesRequest {
	s.DentryIds = v
	return s
}

func (s *GetDentriesRequest) SetOption(v *GetDentriesRequestOption) *GetDentriesRequest {
	s.Option = v
	return s
}

func (s *GetDentriesRequest) SetUnionId(v string) *GetDentriesRequest {
	s.UnionId = &v
	return s
}

type GetDentriesRequestOption struct {
	AppIdsForAppProperties []*string `json:"appIdsForAppProperties,omitempty" xml:"appIdsForAppProperties,omitempty" type:"Repeated"`
	// example:
	//
	// true
	WithThumbnail *bool `json:"withThumbnail,omitempty" xml:"withThumbnail,omitempty"`
}

func (s GetDentriesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesRequestOption) GoString() string {
	return s.String()
}

func (s *GetDentriesRequestOption) SetAppIdsForAppProperties(v []*string) *GetDentriesRequestOption {
	s.AppIdsForAppProperties = v
	return s
}

func (s *GetDentriesRequestOption) SetWithThumbnail(v bool) *GetDentriesRequestOption {
	s.WithThumbnail = &v
	return s
}

type GetDentriesResponseBody struct {
	ResultItems []*GetDentriesResponseBodyResultItems `json:"resultItems,omitempty" xml:"resultItems,omitempty" type:"Repeated"`
}

func (s GetDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDentriesResponseBody) SetResultItems(v []*GetDentriesResponseBodyResultItems) *GetDentriesResponseBody {
	s.ResultItems = v
	return s
}

type GetDentriesResponseBodyResultItems struct {
	Dentry *GetDentriesResponseBodyResultItemsDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
	// example:
	//
	// dentry_id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// permissionDenied
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDentriesResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *GetDentriesResponseBodyResultItems) SetDentry(v *GetDentriesResponseBodyResultItemsDentry) *GetDentriesResponseBodyResultItems {
	s.Dentry = v
	return s
}

func (s *GetDentriesResponseBodyResultItems) SetDentryId(v string) *GetDentriesResponseBodyResultItems {
	s.DentryId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItems) SetErrorCode(v string) *GetDentriesResponseBodyResultItems {
	s.ErrorCode = &v
	return s
}

func (s *GetDentriesResponseBodyResultItems) SetSpaceId(v string) *GetDentriesResponseBodyResultItems {
	s.SpaceId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItems) SetSuccess(v bool) *GetDentriesResponseBodyResultItems {
	s.Success = &v
	return s
}

type GetDentriesResponseBodyResultItemsDentry struct {
	AppProperties map[string][]*ResultItemsDentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                             `json:"path,omitempty" xml:"path,omitempty"`
	Properties *GetDentriesResponseBodyResultItemsDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string                                            `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	Thumbnail     *GetDentriesResponseBodyResultItemsDentryThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDentriesResponseBodyResultItemsDentry) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponseBodyResultItemsDentry) GoString() string {
	return s.String()
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetAppProperties(v map[string][]*ResultItemsDentryAppPropertiesValue) *GetDentriesResponseBodyResultItemsDentry {
	s.AppProperties = v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetCreateTime(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.CreateTime = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetCreatorId(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.CreatorId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetExtension(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Extension = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetId(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Id = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetModifiedTime(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.ModifiedTime = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetModifierId(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.ModifierId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetName(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Name = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetParentId(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.ParentId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetPartitionType(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.PartitionType = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetPath(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Path = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetProperties(v *GetDentriesResponseBodyResultItemsDentryProperties) *GetDentriesResponseBodyResultItemsDentry {
	s.Properties = v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetSize(v int64) *GetDentriesResponseBodyResultItemsDentry {
	s.Size = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetSpaceId(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.SpaceId = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetStatus(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Status = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetStorageDriver(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.StorageDriver = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetThumbnail(v *GetDentriesResponseBodyResultItemsDentryThumbnail) *GetDentriesResponseBodyResultItemsDentry {
	s.Thumbnail = v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetType(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Type = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetUuid(v string) *GetDentriesResponseBodyResultItemsDentry {
	s.Uuid = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentry) SetVersion(v int64) *GetDentriesResponseBodyResultItemsDentry {
	s.Version = &v
	return s
}

type GetDentriesResponseBodyResultItemsDentryProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s GetDentriesResponseBodyResultItemsDentryProperties) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponseBodyResultItemsDentryProperties) GoString() string {
	return s.String()
}

func (s *GetDentriesResponseBodyResultItemsDentryProperties) SetReadOnly(v bool) *GetDentriesResponseBodyResultItemsDentryProperties {
	s.ReadOnly = &v
	return s
}

type GetDentriesResponseBodyResultItemsDentryThumbnail struct {
	// example:
	//
	// 64
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// 64
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s GetDentriesResponseBodyResultItemsDentryThumbnail) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponseBodyResultItemsDentryThumbnail) GoString() string {
	return s.String()
}

func (s *GetDentriesResponseBodyResultItemsDentryThumbnail) SetHeight(v int32) *GetDentriesResponseBodyResultItemsDentryThumbnail {
	s.Height = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentryThumbnail) SetUrl(v string) *GetDentriesResponseBodyResultItemsDentryThumbnail {
	s.Url = &v
	return s
}

func (s *GetDentriesResponseBodyResultItemsDentryThumbnail) SetWidth(v int32) *GetDentriesResponseBodyResultItemsDentryThumbnail {
	s.Width = &v
	return s
}

type GetDentriesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDentriesResponse) GoString() string {
	return s.String()
}

func (s *GetDentriesResponse) SetHeaders(v map[string]*string) *GetDentriesResponse {
	s.Headers = v
	return s
}

func (s *GetDentriesResponse) SetStatusCode(v int32) *GetDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDentriesResponse) SetBody(v *GetDentriesResponseBody) *GetDentriesResponse {
	s.Body = v
	return s
}

type GetDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDentryHeaders) GoString() string {
	return s.String()
}

func (s *GetDentryHeaders) SetCommonHeaders(v map[string]*string) *GetDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDentryHeaders) SetXAcsDingtalkAccessToken(v string) *GetDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDentryRequest struct {
	Option *GetDentryRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDentryRequest) GoString() string {
	return s.String()
}

func (s *GetDentryRequest) SetOption(v *GetDentryRequestOption) *GetDentryRequest {
	s.Option = v
	return s
}

func (s *GetDentryRequest) SetUnionId(v string) *GetDentryRequest {
	s.UnionId = &v
	return s
}

type GetDentryRequestOption struct {
	AppIdsForAppProperties []*string `json:"appIdsForAppProperties,omitempty" xml:"appIdsForAppProperties,omitempty" type:"Repeated"`
	// example:
	//
	// true
	WithThumbnail *bool `json:"withThumbnail,omitempty" xml:"withThumbnail,omitempty"`
}

func (s GetDentryRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetDentryRequestOption) GoString() string {
	return s.String()
}

func (s *GetDentryRequestOption) SetAppIdsForAppProperties(v []*string) *GetDentryRequestOption {
	s.AppIdsForAppProperties = v
	return s
}

func (s *GetDentryRequestOption) SetWithThumbnail(v bool) *GetDentryRequestOption {
	s.WithThumbnail = &v
	return s
}

type GetDentryResponseBody struct {
	Dentry *GetDentryResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
}

func (s GetDentryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDentryResponseBody) SetDentry(v *GetDentryResponseBodyDentry) *GetDentryResponseBody {
	s.Dentry = v
	return s
}

type GetDentryResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                `json:"path,omitempty" xml:"path,omitempty"`
	Properties *GetDentryResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string                               `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	Thumbnail     *GetDentryResponseBodyDentryThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDentryResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *GetDentryResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *GetDentryResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *GetDentryResponseBodyDentry) SetCreateTime(v string) *GetDentryResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetCreatorId(v string) *GetDentryResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetExtension(v string) *GetDentryResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetId(v string) *GetDentryResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetModifiedTime(v string) *GetDentryResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetModifierId(v string) *GetDentryResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetName(v string) *GetDentryResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetParentId(v string) *GetDentryResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetPartitionType(v string) *GetDentryResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetPath(v string) *GetDentryResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetProperties(v *GetDentryResponseBodyDentryProperties) *GetDentryResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *GetDentryResponseBodyDentry) SetSize(v int64) *GetDentryResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetSpaceId(v string) *GetDentryResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetStatus(v string) *GetDentryResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetStorageDriver(v string) *GetDentryResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetThumbnail(v *GetDentryResponseBodyDentryThumbnail) *GetDentryResponseBodyDentry {
	s.Thumbnail = v
	return s
}

func (s *GetDentryResponseBodyDentry) SetType(v string) *GetDentryResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetUuid(v string) *GetDentryResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *GetDentryResponseBodyDentry) SetVersion(v int64) *GetDentryResponseBodyDentry {
	s.Version = &v
	return s
}

type GetDentryResponseBodyDentryProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s GetDentryResponseBodyDentryProperties) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *GetDentryResponseBodyDentryProperties) SetReadOnly(v bool) *GetDentryResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

type GetDentryResponseBodyDentryThumbnail struct {
	// example:
	//
	// 64
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// 64
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s GetDentryResponseBodyDentryThumbnail) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponseBodyDentryThumbnail) GoString() string {
	return s.String()
}

func (s *GetDentryResponseBodyDentryThumbnail) SetHeight(v int32) *GetDentryResponseBodyDentryThumbnail {
	s.Height = &v
	return s
}

func (s *GetDentryResponseBodyDentryThumbnail) SetUrl(v string) *GetDentryResponseBodyDentryThumbnail {
	s.Url = &v
	return s
}

func (s *GetDentryResponseBodyDentryThumbnail) SetWidth(v int32) *GetDentryResponseBodyDentryThumbnail {
	s.Width = &v
	return s
}

type GetDentryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDentryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponse) GoString() string {
	return s.String()
}

func (s *GetDentryResponse) SetHeaders(v map[string]*string) *GetDentryResponse {
	s.Headers = v
	return s
}

func (s *GetDentryResponse) SetStatusCode(v int32) *GetDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDentryResponse) SetBody(v *GetDentryResponseBody) *GetDentryResponse {
	s.Body = v
	return s
}

type GetDentryOpenInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDentryOpenInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDentryOpenInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDentryOpenInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDentryOpenInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDentryOpenInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetDentryOpenInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDentryOpenInfoRequest struct {
	Option *GetDentryOpenInfoRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetDentryOpenInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDentryOpenInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDentryOpenInfoRequest) SetOption(v *GetDentryOpenInfoRequestOption) *GetDentryOpenInfoRequest {
	s.Option = v
	return s
}

func (s *GetDentryOpenInfoRequest) SetUnionId(v string) *GetDentryOpenInfoRequest {
	s.UnionId = &v
	return s
}

type GetDentryOpenInfoRequestOption struct {
	// example:
	//
	// true
	CheckLogin *bool `json:"checkLogin,omitempty" xml:"checkLogin,omitempty"`
	// example:
	//
	// PREVIEW
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// true
	WaterMark *bool `json:"waterMark,omitempty" xml:"waterMark,omitempty"`
}

func (s GetDentryOpenInfoRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetDentryOpenInfoRequestOption) GoString() string {
	return s.String()
}

func (s *GetDentryOpenInfoRequestOption) SetCheckLogin(v bool) *GetDentryOpenInfoRequestOption {
	s.CheckLogin = &v
	return s
}

func (s *GetDentryOpenInfoRequestOption) SetType(v string) *GetDentryOpenInfoRequestOption {
	s.Type = &v
	return s
}

func (s *GetDentryOpenInfoRequestOption) SetVersion(v int64) *GetDentryOpenInfoRequestOption {
	s.Version = &v
	return s
}

func (s *GetDentryOpenInfoRequestOption) SetWaterMark(v bool) *GetDentryOpenInfoRequestOption {
	s.WaterMark = &v
	return s
}

type GetDentryOpenInfoResponseBody struct {
	// example:
	//
	// true
	HasWaterMark *bool `json:"hasWaterMark,omitempty" xml:"hasWaterMark,omitempty"`
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDentryOpenInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDentryOpenInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDentryOpenInfoResponseBody) SetHasWaterMark(v bool) *GetDentryOpenInfoResponseBody {
	s.HasWaterMark = &v
	return s
}

func (s *GetDentryOpenInfoResponseBody) SetUrl(v string) *GetDentryOpenInfoResponseBody {
	s.Url = &v
	return s
}

type GetDentryOpenInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDentryOpenInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDentryOpenInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDentryOpenInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDentryOpenInfoResponse) SetHeaders(v map[string]*string) *GetDentryOpenInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDentryOpenInfoResponse) SetStatusCode(v int32) *GetDentryOpenInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDentryOpenInfoResponse) SetBody(v *GetDentryOpenInfoResponseBody) *GetDentryOpenInfoResponse {
	s.Body = v
	return s
}

type GetDentryThumbnailsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDentryThumbnailsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsHeaders) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsHeaders) SetCommonHeaders(v map[string]*string) *GetDentryThumbnailsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDentryThumbnailsHeaders) SetXAcsDingtalkAccessToken(v string) *GetDentryThumbnailsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDentryThumbnailsRequest struct {
	// This parameter is required.
	DentryIds []*string `json:"dentryIds,omitempty" xml:"dentryIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetDentryThumbnailsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsRequest) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsRequest) SetDentryIds(v []*string) *GetDentryThumbnailsRequest {
	s.DentryIds = v
	return s
}

func (s *GetDentryThumbnailsRequest) SetUnionId(v string) *GetDentryThumbnailsRequest {
	s.UnionId = &v
	return s
}

type GetDentryThumbnailsResponseBody struct {
	ResultItems []*GetDentryThumbnailsResponseBodyResultItems `json:"resultItems,omitempty" xml:"resultItems,omitempty" type:"Repeated"`
}

func (s GetDentryThumbnailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsResponseBody) SetResultItems(v []*GetDentryThumbnailsResponseBodyResultItems) *GetDentryThumbnailsResponseBody {
	s.ResultItems = v
	return s
}

type GetDentryThumbnailsResponseBodyResultItems struct {
	// example:
	//
	// dentry_id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// permissionDenied
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// true
	Success   *bool                                                `json:"success,omitempty" xml:"success,omitempty"`
	Thumbnail *GetDentryThumbnailsResponseBodyResultItemsThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
}

func (s GetDentryThumbnailsResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsResponseBodyResultItems) SetDentryId(v string) *GetDentryThumbnailsResponseBodyResultItems {
	s.DentryId = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItems) SetErrorCode(v string) *GetDentryThumbnailsResponseBodyResultItems {
	s.ErrorCode = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItems) SetSpaceId(v string) *GetDentryThumbnailsResponseBodyResultItems {
	s.SpaceId = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItems) SetSuccess(v bool) *GetDentryThumbnailsResponseBodyResultItems {
	s.Success = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItems) SetThumbnail(v *GetDentryThumbnailsResponseBodyResultItemsThumbnail) *GetDentryThumbnailsResponseBodyResultItems {
	s.Thumbnail = v
	return s
}

type GetDentryThumbnailsResponseBodyResultItemsThumbnail struct {
	// example:
	//
	// 64
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// 64
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s GetDentryThumbnailsResponseBodyResultItemsThumbnail) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsResponseBodyResultItemsThumbnail) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsResponseBodyResultItemsThumbnail) SetHeight(v int32) *GetDentryThumbnailsResponseBodyResultItemsThumbnail {
	s.Height = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItemsThumbnail) SetUrl(v string) *GetDentryThumbnailsResponseBodyResultItemsThumbnail {
	s.Url = &v
	return s
}

func (s *GetDentryThumbnailsResponseBodyResultItemsThumbnail) SetWidth(v int32) *GetDentryThumbnailsResponseBodyResultItemsThumbnail {
	s.Width = &v
	return s
}

type GetDentryThumbnailsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDentryThumbnailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDentryThumbnailsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDentryThumbnailsResponse) GoString() string {
	return s.String()
}

func (s *GetDentryThumbnailsResponse) SetHeaders(v map[string]*string) *GetDentryThumbnailsResponse {
	s.Headers = v
	return s
}

func (s *GetDentryThumbnailsResponse) SetStatusCode(v int32) *GetDentryThumbnailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDentryThumbnailsResponse) SetBody(v *GetDentryThumbnailsResponseBody) *GetDentryThumbnailsResponse {
	s.Body = v
	return s
}

type GetFileDownloadInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFileDownloadInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFileDownloadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileDownloadInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetFileDownloadInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFileDownloadInfoRequest struct {
	Option *GetFileDownloadInfoRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetFileDownloadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoRequest) SetOption(v *GetFileDownloadInfoRequestOption) *GetFileDownloadInfoRequest {
	s.Option = v
	return s
}

func (s *GetFileDownloadInfoRequest) SetUnionId(v string) *GetFileDownloadInfoRequest {
	s.UnionId = &v
	return s
}

type GetFileDownloadInfoRequestOption struct {
	// example:
	//
	// true
	PreferIntranet *bool `json:"preferIntranet,omitempty" xml:"preferIntranet,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetFileDownloadInfoRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoRequestOption) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoRequestOption) SetPreferIntranet(v bool) *GetFileDownloadInfoRequestOption {
	s.PreferIntranet = &v
	return s
}

func (s *GetFileDownloadInfoRequestOption) SetVersion(v int64) *GetFileDownloadInfoRequestOption {
	s.Version = &v
	return s
}

type GetFileDownloadInfoResponseBody struct {
	HeaderSignatureInfo *GetFileDownloadInfoResponseBodyHeaderSignatureInfo `json:"headerSignatureInfo,omitempty" xml:"headerSignatureInfo,omitempty" type:"Struct"`
	// example:
	//
	// HEADER_SIGNATURE
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GetFileDownloadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoResponseBody) SetHeaderSignatureInfo(v *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) *GetFileDownloadInfoResponseBody {
	s.HeaderSignatureInfo = v
	return s
}

func (s *GetFileDownloadInfoResponseBody) SetProtocol(v string) *GetFileDownloadInfoResponseBody {
	s.Protocol = &v
	return s
}

type GetFileDownloadInfoResponseBodyHeaderSignatureInfo struct {
	// example:
	//
	// 900
	ExpirationSeconds    *int32             `json:"expirationSeconds,omitempty" xml:"expirationSeconds,omitempty"`
	Headers              map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	InternalResourceUrls []*string          `json:"internalResourceUrls,omitempty" xml:"internalResourceUrls,omitempty" type:"Repeated"`
	// example:
	//
	// ZHANGJIAKOU
	Region       *string   `json:"region,omitempty" xml:"region,omitempty"`
	ResourceUrls []*string `json:"resourceUrls,omitempty" xml:"resourceUrls,omitempty" type:"Repeated"`
}

func (s GetFileDownloadInfoResponseBodyHeaderSignatureInfo) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoResponseBodyHeaderSignatureInfo) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetExpirationSeconds(v int32) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.ExpirationSeconds = &v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetHeaders(v map[string]*string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.Headers = v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetInternalResourceUrls(v []*string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.InternalResourceUrls = v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetRegion(v string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.Region = &v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetResourceUrls(v []*string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.ResourceUrls = v
	return s
}

type GetFileDownloadInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileDownloadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoResponse) SetHeaders(v map[string]*string) *GetFileDownloadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFileDownloadInfoResponse) SetStatusCode(v int32) *GetFileDownloadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileDownloadInfoResponse) SetBody(v *GetFileDownloadInfoResponseBody) *GetFileDownloadInfoResponse {
	s.Body = v
	return s
}

type GetFileUploadInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFileUploadInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFileUploadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileUploadInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetFileUploadInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFileUploadInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Multipart *bool                           `json:"multipart,omitempty" xml:"multipart,omitempty"`
	Option    *GetFileUploadInfoRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// HEADER_SIGNATURE
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetFileUploadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequest) SetMultipart(v bool) *GetFileUploadInfoRequest {
	s.Multipart = &v
	return s
}

func (s *GetFileUploadInfoRequest) SetOption(v *GetFileUploadInfoRequestOption) *GetFileUploadInfoRequest {
	s.Option = v
	return s
}

func (s *GetFileUploadInfoRequest) SetProtocol(v string) *GetFileUploadInfoRequest {
	s.Protocol = &v
	return s
}

func (s *GetFileUploadInfoRequest) SetUnionId(v string) *GetFileUploadInfoRequest {
	s.UnionId = &v
	return s
}

type GetFileUploadInfoRequestOption struct {
	PreCheckParam *GetFileUploadInfoRequestOptionPreCheckParam `json:"preCheckParam,omitempty" xml:"preCheckParam,omitempty" type:"Struct"`
	// example:
	//
	// true
	PreferIntranet *bool `json:"preferIntranet,omitempty" xml:"preferIntranet,omitempty"`
	// example:
	//
	// ZHANGJIAKOU
	PreferRegion *string `json:"preferRegion,omitempty" xml:"preferRegion,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
}

func (s GetFileUploadInfoRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoRequestOption) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequestOption) SetPreCheckParam(v *GetFileUploadInfoRequestOptionPreCheckParam) *GetFileUploadInfoRequestOption {
	s.PreCheckParam = v
	return s
}

func (s *GetFileUploadInfoRequestOption) SetPreferIntranet(v bool) *GetFileUploadInfoRequestOption {
	s.PreferIntranet = &v
	return s
}

func (s *GetFileUploadInfoRequestOption) SetPreferRegion(v string) *GetFileUploadInfoRequestOption {
	s.PreferRegion = &v
	return s
}

func (s *GetFileUploadInfoRequestOption) SetStorageDriver(v string) *GetFileUploadInfoRequestOption {
	s.StorageDriver = &v
	return s
}

type GetFileUploadInfoRequestOptionPreCheckParam struct {
	// example:
	//
	// md5
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s GetFileUploadInfoRequestOptionPreCheckParam) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoRequestOptionPreCheckParam) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) SetMd5(v string) *GetFileUploadInfoRequestOptionPreCheckParam {
	s.Md5 = &v
	return s
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) SetName(v string) *GetFileUploadInfoRequestOptionPreCheckParam {
	s.Name = &v
	return s
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) SetParentId(v string) *GetFileUploadInfoRequestOptionPreCheckParam {
	s.ParentId = &v
	return s
}

func (s *GetFileUploadInfoRequestOptionPreCheckParam) SetSize(v int64) *GetFileUploadInfoRequestOptionPreCheckParam {
	s.Size = &v
	return s
}

type GetFileUploadInfoResponseBody struct {
	HeaderSignatureInfo *GetFileUploadInfoResponseBodyHeaderSignatureInfo `json:"headerSignatureInfo,omitempty" xml:"headerSignatureInfo,omitempty" type:"Struct"`
	// example:
	//
	// HEADER_SIGNATURE
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// example:
	//
	// upload_key
	UploadKey *string `json:"uploadKey,omitempty" xml:"uploadKey,omitempty"`
}

func (s GetFileUploadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponseBody) SetHeaderSignatureInfo(v *GetFileUploadInfoResponseBodyHeaderSignatureInfo) *GetFileUploadInfoResponseBody {
	s.HeaderSignatureInfo = v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetProtocol(v string) *GetFileUploadInfoResponseBody {
	s.Protocol = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetStorageDriver(v string) *GetFileUploadInfoResponseBody {
	s.StorageDriver = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetUploadKey(v string) *GetFileUploadInfoResponseBody {
	s.UploadKey = &v
	return s
}

type GetFileUploadInfoResponseBodyHeaderSignatureInfo struct {
	// example:
	//
	// 900
	ExpirationSeconds    *int32             `json:"expirationSeconds,omitempty" xml:"expirationSeconds,omitempty"`
	Headers              map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	InternalResourceUrls []*string          `json:"internalResourceUrls,omitempty" xml:"internalResourceUrls,omitempty" type:"Repeated"`
	// example:
	//
	// ZHANGJIAKOU
	Region       *string   `json:"region,omitempty" xml:"region,omitempty"`
	ResourceUrls []*string `json:"resourceUrls,omitempty" xml:"resourceUrls,omitempty" type:"Repeated"`
}

func (s GetFileUploadInfoResponseBodyHeaderSignatureInfo) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoResponseBodyHeaderSignatureInfo) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetExpirationSeconds(v int32) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.ExpirationSeconds = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetHeaders(v map[string]*string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.Headers = v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetInternalResourceUrls(v []*string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.InternalResourceUrls = v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetRegion(v string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.Region = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetResourceUrls(v []*string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.ResourceUrls = v
	return s
}

type GetFileUploadInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileUploadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponse) SetHeaders(v map[string]*string) *GetFileUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFileUploadInfoResponse) SetStatusCode(v int32) *GetFileUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileUploadInfoResponse) SetBody(v *GetFileUploadInfoResponseBody) *GetFileUploadInfoResponse {
	s.Body = v
	return s
}

type GetMultipartFileUploadInfosHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMultipartFileUploadInfosHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMultipartFileUploadInfosHeaders) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosHeaders) SetCommonHeaders(v map[string]*string) *GetMultipartFileUploadInfosHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMultipartFileUploadInfosHeaders) SetXAcsDingtalkAccessToken(v string) *GetMultipartFileUploadInfosHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMultipartFileUploadInfosRequest struct {
	Option *GetMultipartFileUploadInfosRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	PartNumbers []*int32 `json:"partNumbers,omitempty" xml:"partNumbers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// upload_key
	UploadKey *string `json:"uploadKey,omitempty" xml:"uploadKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetMultipartFileUploadInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultipartFileUploadInfosRequest) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosRequest) SetOption(v *GetMultipartFileUploadInfosRequestOption) *GetMultipartFileUploadInfosRequest {
	s.Option = v
	return s
}

func (s *GetMultipartFileUploadInfosRequest) SetPartNumbers(v []*int32) *GetMultipartFileUploadInfosRequest {
	s.PartNumbers = v
	return s
}

func (s *GetMultipartFileUploadInfosRequest) SetUploadKey(v string) *GetMultipartFileUploadInfosRequest {
	s.UploadKey = &v
	return s
}

func (s *GetMultipartFileUploadInfosRequest) SetUnionId(v string) *GetMultipartFileUploadInfosRequest {
	s.UnionId = &v
	return s
}

type GetMultipartFileUploadInfosRequestOption struct {
	// example:
	//
	// true
	PreferIntranet *bool `json:"preferIntranet,omitempty" xml:"preferIntranet,omitempty"`
}

func (s GetMultipartFileUploadInfosRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetMultipartFileUploadInfosRequestOption) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosRequestOption) SetPreferIntranet(v bool) *GetMultipartFileUploadInfosRequestOption {
	s.PreferIntranet = &v
	return s
}

type GetMultipartFileUploadInfosResponseBody struct {
	MultipartHeaderSignatureInfos []*GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos `json:"multipartHeaderSignatureInfos,omitempty" xml:"multipartHeaderSignatureInfos,omitempty" type:"Repeated"`
}

func (s GetMultipartFileUploadInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultipartFileUploadInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosResponseBody) SetMultipartHeaderSignatureInfos(v []*GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) *GetMultipartFileUploadInfosResponseBody {
	s.MultipartHeaderSignatureInfos = v
	return s
}

type GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos struct {
	HeaderSignatureInfo *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo `json:"headerSignatureInfo,omitempty" xml:"headerSignatureInfo,omitempty" type:"Struct"`
	PartNumber          *int32                                                                                   `json:"partNumber,omitempty" xml:"partNumber,omitempty"`
}

func (s GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) String() string {
	return tea.Prettify(s)
}

func (s GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) SetHeaderSignatureInfo(v *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos {
	s.HeaderSignatureInfo = v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) SetPartNumber(v int32) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos {
	s.PartNumber = &v
	return s
}

type GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo struct {
	// example:
	//
	// 900
	ExpirationSeconds    *int32             `json:"expirationSeconds,omitempty" xml:"expirationSeconds,omitempty"`
	Headers              map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	InternalResourceUrls []*string          `json:"internalResourceUrls,omitempty" xml:"internalResourceUrls,omitempty" type:"Repeated"`
	// example:
	//
	// ZHANGJIAKOU
	Region       *string   `json:"region,omitempty" xml:"region,omitempty"`
	ResourceUrls []*string `json:"resourceUrls,omitempty" xml:"resourceUrls,omitempty" type:"Repeated"`
}

func (s GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) SetExpirationSeconds(v int32) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	s.ExpirationSeconds = &v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) SetHeaders(v map[string]*string) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	s.Headers = v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) SetInternalResourceUrls(v []*string) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	s.InternalResourceUrls = v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) SetRegion(v string) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	s.Region = &v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) SetResourceUrls(v []*string) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	s.ResourceUrls = v
	return s
}

type GetMultipartFileUploadInfosResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultipartFileUploadInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultipartFileUploadInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultipartFileUploadInfosResponse) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosResponse) SetHeaders(v map[string]*string) *GetMultipartFileUploadInfosResponse {
	s.Headers = v
	return s
}

func (s *GetMultipartFileUploadInfosResponse) SetStatusCode(v int32) *GetMultipartFileUploadInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultipartFileUploadInfosResponse) SetBody(v *GetMultipartFileUploadInfosResponseBody) *GetMultipartFileUploadInfosResponse {
	s.Body = v
	return s
}

type GetOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrgHeaders) GoString() string {
	return s.String()
}

func (s *GetOrgHeaders) SetCommonHeaders(v map[string]*string) *GetOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrgHeaders) SetXAcsDingtalkAccessToken(v string) *GetOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOrgRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrgRequest) GoString() string {
	return s.String()
}

func (s *GetOrgRequest) SetUnionId(v string) *GetOrgRequest {
	s.UnionId = &v
	return s
}

type GetOrgResponseBody struct {
	Org *GetOrgResponseBodyOrg `json:"org,omitempty" xml:"org,omitempty" type:"Struct"`
}

func (s GetOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrgResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgResponseBody) SetOrg(v *GetOrgResponseBodyOrg) *GetOrgResponseBody {
	s.Org = v
	return s
}

type GetOrgResponseBodyOrg struct {
	// example:
	//
	// corp_id
	CorpId     *string                            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Partitions []*GetOrgResponseBodyOrgPartitions `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
}

func (s GetOrgResponseBodyOrg) String() string {
	return tea.Prettify(s)
}

func (s GetOrgResponseBodyOrg) GoString() string {
	return s.String()
}

func (s *GetOrgResponseBodyOrg) SetCorpId(v string) *GetOrgResponseBodyOrg {
	s.CorpId = &v
	return s
}

func (s *GetOrgResponseBodyOrg) SetPartitions(v []*GetOrgResponseBodyOrgPartitions) *GetOrgResponseBodyOrg {
	s.Partitions = v
	return s
}

type GetOrgResponseBodyOrgPartitions struct {
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string                               `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	Quota         *GetOrgResponseBodyOrgPartitionsQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetOrgResponseBodyOrgPartitions) String() string {
	return tea.Prettify(s)
}

func (s GetOrgResponseBodyOrgPartitions) GoString() string {
	return s.String()
}

func (s *GetOrgResponseBodyOrgPartitions) SetPartitionType(v string) *GetOrgResponseBodyOrgPartitions {
	s.PartitionType = &v
	return s
}

func (s *GetOrgResponseBodyOrgPartitions) SetQuota(v *GetOrgResponseBodyOrgPartitionsQuota) *GetOrgResponseBodyOrgPartitions {
	s.Quota = v
	return s
}

type GetOrgResponseBodyOrgPartitionsQuota struct {
	// example:
	//
	// 10000
	Max *int64 `json:"max,omitempty" xml:"max,omitempty"`
	// example:
	//
	// 1000
	Reserved *int64 `json:"reserved,omitempty" xml:"reserved,omitempty"`
	// example:
	//
	// SHARE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1024
	Used *int64 `json:"used,omitempty" xml:"used,omitempty"`
}

func (s GetOrgResponseBodyOrgPartitionsQuota) String() string {
	return tea.Prettify(s)
}

func (s GetOrgResponseBodyOrgPartitionsQuota) GoString() string {
	return s.String()
}

func (s *GetOrgResponseBodyOrgPartitionsQuota) SetMax(v int64) *GetOrgResponseBodyOrgPartitionsQuota {
	s.Max = &v
	return s
}

func (s *GetOrgResponseBodyOrgPartitionsQuota) SetReserved(v int64) *GetOrgResponseBodyOrgPartitionsQuota {
	s.Reserved = &v
	return s
}

func (s *GetOrgResponseBodyOrgPartitionsQuota) SetType(v string) *GetOrgResponseBodyOrgPartitionsQuota {
	s.Type = &v
	return s
}

func (s *GetOrgResponseBodyOrgPartitionsQuota) SetUsed(v int64) *GetOrgResponseBodyOrgPartitionsQuota {
	s.Used = &v
	return s
}

type GetOrgResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrgResponse) GoString() string {
	return s.String()
}

func (s *GetOrgResponse) SetHeaders(v map[string]*string) *GetOrgResponse {
	s.Headers = v
	return s
}

func (s *GetOrgResponse) SetStatusCode(v int32) *GetOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgResponse) SetBody(v *GetOrgResponseBody) *GetOrgResponse {
	s.Body = v
	return s
}

type GetRecycleBinHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRecycleBinHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinHeaders) GoString() string {
	return s.String()
}

func (s *GetRecycleBinHeaders) SetCommonHeaders(v map[string]*string) *GetRecycleBinHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRecycleBinHeaders) SetXAcsDingtalkAccessToken(v string) *GetRecycleBinHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRecycleBinRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SPACE
	RecycleBinScope *string `json:"recycleBinScope,omitempty" xml:"recycleBinScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scope_id
	ScopeId *string `json:"scopeId,omitempty" xml:"scopeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *GetRecycleBinRequest) SetRecycleBinScope(v string) *GetRecycleBinRequest {
	s.RecycleBinScope = &v
	return s
}

func (s *GetRecycleBinRequest) SetScopeId(v string) *GetRecycleBinRequest {
	s.ScopeId = &v
	return s
}

func (s *GetRecycleBinRequest) SetUnionId(v string) *GetRecycleBinRequest {
	s.UnionId = &v
	return s
}

type GetRecycleBinResponseBody struct {
	RecycleBin *GetRecycleBinResponseBodyRecycleBin `json:"recycleBin,omitempty" xml:"recycleBin,omitempty" type:"Struct"`
}

func (s GetRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecycleBinResponseBody) SetRecycleBin(v *GetRecycleBinResponseBodyRecycleBin) *GetRecycleBinResponseBody {
	s.RecycleBin = v
	return s
}

type GetRecycleBinResponseBodyRecycleBin struct {
	// example:
	//
	// recyclebin_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// SPACE
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// scope_id
	ScopeId *string `json:"scopeId,omitempty" xml:"scopeId,omitempty"`
}

func (s GetRecycleBinResponseBodyRecycleBin) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinResponseBodyRecycleBin) GoString() string {
	return s.String()
}

func (s *GetRecycleBinResponseBodyRecycleBin) SetId(v string) *GetRecycleBinResponseBodyRecycleBin {
	s.Id = &v
	return s
}

func (s *GetRecycleBinResponseBodyRecycleBin) SetScope(v string) *GetRecycleBinResponseBodyRecycleBin {
	s.Scope = &v
	return s
}

func (s *GetRecycleBinResponseBodyRecycleBin) SetScopeId(v string) *GetRecycleBinResponseBodyRecycleBin {
	s.ScopeId = &v
	return s
}

type GetRecycleBinResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *GetRecycleBinResponse) SetHeaders(v map[string]*string) *GetRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *GetRecycleBinResponse) SetStatusCode(v int32) *GetRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecycleBinResponse) SetBody(v *GetRecycleBinResponseBody) *GetRecycleBinResponse {
	s.Body = v
	return s
}

type GetRecycleItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRecycleItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleItemHeaders) GoString() string {
	return s.String()
}

func (s *GetRecycleItemHeaders) SetCommonHeaders(v map[string]*string) *GetRecycleItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRecycleItemHeaders) SetXAcsDingtalkAccessToken(v string) *GetRecycleItemHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRecycleItemRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetRecycleItemRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleItemRequest) GoString() string {
	return s.String()
}

func (s *GetRecycleItemRequest) SetUnionId(v string) *GetRecycleItemRequest {
	s.UnionId = &v
	return s
}

type GetRecycleItemResponseBody struct {
	Item *GetRecycleItemResponseBodyItem `json:"item,omitempty" xml:"item,omitempty" type:"Struct"`
}

func (s GetRecycleItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecycleItemResponseBody) SetItem(v *GetRecycleItemResponseBodyItem) *GetRecycleItemResponseBody {
	s.Item = v
	return s
}

type GetRecycleItemResponseBodyItem struct {
	// example:
	//
	// dentry_id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// recycle_item_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// operator_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	OperatorTime *string `json:"operatorTime,omitempty" xml:"operatorTime,omitempty"`
	// example:
	//
	// dentry_name
	OriginalName *string `json:"originalName,omitempty" xml:"originalName,omitempty"`
	// example:
	//
	// dentry_path
	OriginalPath *string `json:"originalPath,omitempty" xml:"originalPath,omitempty"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetRecycleItemResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleItemResponseBodyItem) GoString() string {
	return s.String()
}

func (s *GetRecycleItemResponseBodyItem) SetDentryId(v string) *GetRecycleItemResponseBodyItem {
	s.DentryId = &v
	return s
}

func (s *GetRecycleItemResponseBodyItem) SetId(v string) *GetRecycleItemResponseBodyItem {
	s.Id = &v
	return s
}

func (s *GetRecycleItemResponseBodyItem) SetOperatorId(v string) *GetRecycleItemResponseBodyItem {
	s.OperatorId = &v
	return s
}

func (s *GetRecycleItemResponseBodyItem) SetOperatorTime(v string) *GetRecycleItemResponseBodyItem {
	s.OperatorTime = &v
	return s
}

func (s *GetRecycleItemResponseBodyItem) SetOriginalName(v string) *GetRecycleItemResponseBodyItem {
	s.OriginalName = &v
	return s
}

func (s *GetRecycleItemResponseBodyItem) SetOriginalPath(v string) *GetRecycleItemResponseBodyItem {
	s.OriginalPath = &v
	return s
}

func (s *GetRecycleItemResponseBodyItem) SetSize(v int64) *GetRecycleItemResponseBodyItem {
	s.Size = &v
	return s
}

func (s *GetRecycleItemResponseBodyItem) SetSpaceId(v string) *GetRecycleItemResponseBodyItem {
	s.SpaceId = &v
	return s
}

func (s *GetRecycleItemResponseBodyItem) SetType(v string) *GetRecycleItemResponseBodyItem {
	s.Type = &v
	return s
}

type GetRecycleItemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecycleItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecycleItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleItemResponse) GoString() string {
	return s.String()
}

func (s *GetRecycleItemResponse) SetHeaders(v map[string]*string) *GetRecycleItemResponse {
	s.Headers = v
	return s
}

func (s *GetRecycleItemResponse) SetStatusCode(v int32) *GetRecycleItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecycleItemResponse) SetBody(v *GetRecycleItemResponseBody) *GetRecycleItemResponse {
	s.Body = v
	return s
}

type GetSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceHeaders) GoString() string {
	return s.String()
}

func (s *GetSpaceHeaders) SetCommonHeaders(v map[string]*string) *GetSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *GetSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSpaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceRequest) SetUnionId(v string) *GetSpaceRequest {
	s.UnionId = &v
	return s
}

type GetSpaceResponseBody struct {
	Space *GetSpaceResponseBodySpace `json:"space,omitempty" xml:"space,omitempty" type:"Struct"`
}

func (s GetSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceResponseBody) SetSpace(v *GetSpaceResponseBodySpace) *GetSpaceResponseBody {
	s.Space = v
	return s
}

type GetSpaceResponseBodySpace struct {
	// example:
	//
	// app_id
	AppId        *string                                `json:"appId,omitempty" xml:"appId,omitempty"`
	Capabilities *GetSpaceResponseBodySpaceCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// example:
	//
	// corp_id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// space_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// space_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// owner_id
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// example:
	//
	// USER
	OwnerType  *string                                `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	Partitions []*GetSpaceResponseBodySpacePartitions `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	// example:
	//
	// 1048576
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// example:
	//
	// scene
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// scene_id
	SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1024
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s GetSpaceResponseBodySpace) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBodySpace) GoString() string {
	return s.String()
}

func (s *GetSpaceResponseBodySpace) SetAppId(v string) *GetSpaceResponseBodySpace {
	s.AppId = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetCapabilities(v *GetSpaceResponseBodySpaceCapabilities) *GetSpaceResponseBodySpace {
	s.Capabilities = v
	return s
}

func (s *GetSpaceResponseBodySpace) SetCorpId(v string) *GetSpaceResponseBodySpace {
	s.CorpId = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetCreateTime(v string) *GetSpaceResponseBodySpace {
	s.CreateTime = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetCreatorId(v string) *GetSpaceResponseBodySpace {
	s.CreatorId = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetId(v string) *GetSpaceResponseBodySpace {
	s.Id = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetModifiedTime(v string) *GetSpaceResponseBodySpace {
	s.ModifiedTime = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetModifierId(v string) *GetSpaceResponseBodySpace {
	s.ModifierId = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetName(v string) *GetSpaceResponseBodySpace {
	s.Name = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetOwnerId(v string) *GetSpaceResponseBodySpace {
	s.OwnerId = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetOwnerType(v string) *GetSpaceResponseBodySpace {
	s.OwnerType = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetPartitions(v []*GetSpaceResponseBodySpacePartitions) *GetSpaceResponseBodySpace {
	s.Partitions = v
	return s
}

func (s *GetSpaceResponseBodySpace) SetQuota(v int64) *GetSpaceResponseBodySpace {
	s.Quota = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetScene(v string) *GetSpaceResponseBodySpace {
	s.Scene = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetSceneId(v string) *GetSpaceResponseBodySpace {
	s.SceneId = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetStatus(v string) *GetSpaceResponseBodySpace {
	s.Status = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetUsedQuota(v int64) *GetSpaceResponseBodySpace {
	s.UsedQuota = &v
	return s
}

type GetSpaceResponseBodySpaceCapabilities struct {
	// example:
	//
	// true
	CanRecordRecentFile *bool `json:"canRecordRecentFile,omitempty" xml:"canRecordRecentFile,omitempty"`
	// example:
	//
	// true
	CanRename *bool `json:"canRename,omitempty" xml:"canRename,omitempty"`
	// example:
	//
	// true
	CanSearch *bool `json:"canSearch,omitempty" xml:"canSearch,omitempty"`
}

func (s GetSpaceResponseBodySpaceCapabilities) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBodySpaceCapabilities) GoString() string {
	return s.String()
}

func (s *GetSpaceResponseBodySpaceCapabilities) SetCanRecordRecentFile(v bool) *GetSpaceResponseBodySpaceCapabilities {
	s.CanRecordRecentFile = &v
	return s
}

func (s *GetSpaceResponseBodySpaceCapabilities) SetCanRename(v bool) *GetSpaceResponseBodySpaceCapabilities {
	s.CanRename = &v
	return s
}

func (s *GetSpaceResponseBodySpaceCapabilities) SetCanSearch(v bool) *GetSpaceResponseBodySpaceCapabilities {
	s.CanSearch = &v
	return s
}

type GetSpaceResponseBodySpacePartitions struct {
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string                                   `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	Quota         *GetSpaceResponseBodySpacePartitionsQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetSpaceResponseBodySpacePartitions) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBodySpacePartitions) GoString() string {
	return s.String()
}

func (s *GetSpaceResponseBodySpacePartitions) SetPartitionType(v string) *GetSpaceResponseBodySpacePartitions {
	s.PartitionType = &v
	return s
}

func (s *GetSpaceResponseBodySpacePartitions) SetQuota(v *GetSpaceResponseBodySpacePartitionsQuota) *GetSpaceResponseBodySpacePartitions {
	s.Quota = v
	return s
}

type GetSpaceResponseBodySpacePartitionsQuota struct {
	// example:
	//
	// 10000
	Max *int64 `json:"max,omitempty" xml:"max,omitempty"`
	// example:
	//
	// 1000
	Reserved *int64 `json:"reserved,omitempty" xml:"reserved,omitempty"`
	// example:
	//
	// SHARE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1024
	Used *int64 `json:"used,omitempty" xml:"used,omitempty"`
}

func (s GetSpaceResponseBodySpacePartitionsQuota) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBodySpacePartitionsQuota) GoString() string {
	return s.String()
}

func (s *GetSpaceResponseBodySpacePartitionsQuota) SetMax(v int64) *GetSpaceResponseBodySpacePartitionsQuota {
	s.Max = &v
	return s
}

func (s *GetSpaceResponseBodySpacePartitionsQuota) SetReserved(v int64) *GetSpaceResponseBodySpacePartitionsQuota {
	s.Reserved = &v
	return s
}

func (s *GetSpaceResponseBodySpacePartitionsQuota) SetType(v string) *GetSpaceResponseBodySpacePartitionsQuota {
	s.Type = &v
	return s
}

func (s *GetSpaceResponseBodySpacePartitionsQuota) SetUsed(v int64) *GetSpaceResponseBodySpacePartitionsQuota {
	s.Used = &v
	return s
}

type GetSpaceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceResponse) SetHeaders(v map[string]*string) *GetSpaceResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceResponse) SetStatusCode(v int32) *GetSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpaceResponse) SetBody(v *GetSpaceResponseBody) *GetSpaceResponse {
	s.Body = v
	return s
}

type GetTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTaskHeaders) GoString() string {
	return s.String()
}

func (s *GetTaskHeaders) SetCommonHeaders(v map[string]*string) *GetTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTaskHeaders) SetXAcsDingtalkAccessToken(v string) *GetTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) SetUnionId(v string) *GetTaskRequest {
	s.UnionId = &v
	return s
}

type GetTaskResponseBody struct {
	Task *GetTaskResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody {
	s.Task = v
	return s
}

type GetTaskResponseBodyTask struct {
	// example:
	//
	// 2022-01-01T10:00:00Z
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1
	FailCount *int64 `json:"failCount,omitempty" xml:"failCount,omitempty"`
	// example:
	//
	// permissionDenied
	FailMessage *string `json:"failMessage,omitempty" xml:"failMessage,omitempty"`
	// example:
	//
	// task_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// IN_PROGRESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 3
	SuccessCount *int64 `json:"successCount,omitempty" xml:"successCount,omitempty"`
	// example:
	//
	// 6
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetTaskResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTask) SetBeginTime(v string) *GetTaskResponseBodyTask {
	s.BeginTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetEndTime(v string) *GetTaskResponseBodyTask {
	s.EndTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetFailCount(v int64) *GetTaskResponseBodyTask {
	s.FailCount = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetFailMessage(v string) *GetTaskResponseBodyTask {
	s.FailMessage = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetId(v string) *GetTaskResponseBodyTask {
	s.Id = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetStatus(v string) *GetTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetSuccessCount(v int64) *GetTaskResponseBodyTask {
	s.SuccessCount = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTotalCount(v int64) *GetTaskResponseBodyTask {
	s.TotalCount = &v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetStatusCode(v int32) *GetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type GetWebOfficeUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetWebOfficeUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWebOfficeUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetWebOfficeUrlHeaders) SetCommonHeaders(v map[string]*string) *GetWebOfficeUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWebOfficeUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetWebOfficeUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetWebOfficeUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetWebOfficeUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebOfficeUrlRequest) GoString() string {
	return s.String()
}

func (s *GetWebOfficeUrlRequest) SetUnionId(v string) *GetWebOfficeUrlRequest {
	s.UnionId = &v
	return s
}

type GetWebOfficeUrlResponseBody struct {
	WebOfficeAccessToken  *string `json:"webOfficeAccessToken,omitempty" xml:"webOfficeAccessToken,omitempty"`
	WebOfficeRefreshToken *string `json:"webOfficeRefreshToken,omitempty" xml:"webOfficeRefreshToken,omitempty"`
	WebOfficeUrl          *string `json:"webOfficeUrl,omitempty" xml:"webOfficeUrl,omitempty"`
}

func (s GetWebOfficeUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebOfficeUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebOfficeUrlResponseBody) SetWebOfficeAccessToken(v string) *GetWebOfficeUrlResponseBody {
	s.WebOfficeAccessToken = &v
	return s
}

func (s *GetWebOfficeUrlResponseBody) SetWebOfficeRefreshToken(v string) *GetWebOfficeUrlResponseBody {
	s.WebOfficeRefreshToken = &v
	return s
}

func (s *GetWebOfficeUrlResponseBody) SetWebOfficeUrl(v string) *GetWebOfficeUrlResponseBody {
	s.WebOfficeUrl = &v
	return s
}

type GetWebOfficeUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWebOfficeUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWebOfficeUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebOfficeUrlResponse) GoString() string {
	return s.String()
}

func (s *GetWebOfficeUrlResponse) SetHeaders(v map[string]*string) *GetWebOfficeUrlResponse {
	s.Headers = v
	return s
}

func (s *GetWebOfficeUrlResponse) SetStatusCode(v int32) *GetWebOfficeUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebOfficeUrlResponse) SetBody(v *GetWebOfficeUrlResponseBody) *GetWebOfficeUrlResponse {
	s.Body = v
	return s
}

type InitMultipartFileUploadHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InitMultipartFileUploadHeaders) String() string {
	return tea.Prettify(s)
}

func (s InitMultipartFileUploadHeaders) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadHeaders) SetCommonHeaders(v map[string]*string) *InitMultipartFileUploadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitMultipartFileUploadHeaders) SetXAcsDingtalkAccessToken(v string) *InitMultipartFileUploadHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InitMultipartFileUploadRequest struct {
	Option *InitMultipartFileUploadRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s InitMultipartFileUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s InitMultipartFileUploadRequest) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadRequest) SetOption(v *InitMultipartFileUploadRequestOption) *InitMultipartFileUploadRequest {
	s.Option = v
	return s
}

func (s *InitMultipartFileUploadRequest) SetUnionId(v string) *InitMultipartFileUploadRequest {
	s.UnionId = &v
	return s
}

type InitMultipartFileUploadRequestOption struct {
	PreCheckParam *InitMultipartFileUploadRequestOptionPreCheckParam `json:"preCheckParam,omitempty" xml:"preCheckParam,omitempty" type:"Struct"`
	// example:
	//
	// ZHANGJIAKOU
	PreferRegion *string `json:"preferRegion,omitempty" xml:"preferRegion,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
}

func (s InitMultipartFileUploadRequestOption) String() string {
	return tea.Prettify(s)
}

func (s InitMultipartFileUploadRequestOption) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadRequestOption) SetPreCheckParam(v *InitMultipartFileUploadRequestOptionPreCheckParam) *InitMultipartFileUploadRequestOption {
	s.PreCheckParam = v
	return s
}

func (s *InitMultipartFileUploadRequestOption) SetPreferRegion(v string) *InitMultipartFileUploadRequestOption {
	s.PreferRegion = &v
	return s
}

func (s *InitMultipartFileUploadRequestOption) SetStorageDriver(v string) *InitMultipartFileUploadRequestOption {
	s.StorageDriver = &v
	return s
}

type InitMultipartFileUploadRequestOptionPreCheckParam struct {
	// example:
	//
	// md5
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s InitMultipartFileUploadRequestOptionPreCheckParam) String() string {
	return tea.Prettify(s)
}

func (s InitMultipartFileUploadRequestOptionPreCheckParam) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) SetMd5(v string) *InitMultipartFileUploadRequestOptionPreCheckParam {
	s.Md5 = &v
	return s
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) SetName(v string) *InitMultipartFileUploadRequestOptionPreCheckParam {
	s.Name = &v
	return s
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) SetParentId(v string) *InitMultipartFileUploadRequestOptionPreCheckParam {
	s.ParentId = &v
	return s
}

func (s *InitMultipartFileUploadRequestOptionPreCheckParam) SetSize(v int64) *InitMultipartFileUploadRequestOptionPreCheckParam {
	s.Size = &v
	return s
}

type InitMultipartFileUploadResponseBody struct {
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// example:
	//
	// upload_key
	UploadKey *string `json:"uploadKey,omitempty" xml:"uploadKey,omitempty"`
}

func (s InitMultipartFileUploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitMultipartFileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadResponseBody) SetStorageDriver(v string) *InitMultipartFileUploadResponseBody {
	s.StorageDriver = &v
	return s
}

func (s *InitMultipartFileUploadResponseBody) SetUploadKey(v string) *InitMultipartFileUploadResponseBody {
	s.UploadKey = &v
	return s
}

type InitMultipartFileUploadResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitMultipartFileUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitMultipartFileUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s InitMultipartFileUploadResponse) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadResponse) SetHeaders(v map[string]*string) *InitMultipartFileUploadResponse {
	s.Headers = v
	return s
}

func (s *InitMultipartFileUploadResponse) SetStatusCode(v int32) *InitMultipartFileUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *InitMultipartFileUploadResponse) SetBody(v *InitMultipartFileUploadResponseBody) *InitMultipartFileUploadResponse {
	s.Body = v
	return s
}

type ListAllDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAllDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesHeaders) GoString() string {
	return s.String()
}

func (s *ListAllDentriesHeaders) SetCommonHeaders(v map[string]*string) *ListAllDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAllDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *ListAllDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAllDentriesRequest struct {
	Option *ListAllDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListAllDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesRequest) GoString() string {
	return s.String()
}

func (s *ListAllDentriesRequest) SetOption(v *ListAllDentriesRequestOption) *ListAllDentriesRequest {
	s.Option = v
	return s
}

func (s *ListAllDentriesRequest) SetUnionId(v string) *ListAllDentriesRequest {
	s.UnionId = &v
	return s
}

type ListAllDentriesRequestOption struct {
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// true
	WithThumbnail *bool `json:"withThumbnail,omitempty" xml:"withThumbnail,omitempty"`
}

func (s ListAllDentriesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesRequestOption) GoString() string {
	return s.String()
}

func (s *ListAllDentriesRequestOption) SetMaxResults(v int32) *ListAllDentriesRequestOption {
	s.MaxResults = &v
	return s
}

func (s *ListAllDentriesRequestOption) SetNextToken(v string) *ListAllDentriesRequestOption {
	s.NextToken = &v
	return s
}

func (s *ListAllDentriesRequestOption) SetOrder(v string) *ListAllDentriesRequestOption {
	s.Order = &v
	return s
}

func (s *ListAllDentriesRequestOption) SetWithThumbnail(v bool) *ListAllDentriesRequestOption {
	s.WithThumbnail = &v
	return s
}

type ListAllDentriesResponseBody struct {
	Dentries []*ListAllDentriesResponseBodyDentries `json:"dentries,omitempty" xml:"dentries,omitempty" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAllDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllDentriesResponseBody) SetDentries(v []*ListAllDentriesResponseBodyDentries) *ListAllDentriesResponseBody {
	s.Dentries = v
	return s
}

func (s *ListAllDentriesResponseBody) SetNextToken(v string) *ListAllDentriesResponseBody {
	s.NextToken = &v
	return s
}

type ListAllDentriesResponseBodyDentries struct {
	AppProperties map[string][]*DentriesAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                        `json:"path,omitempty" xml:"path,omitempty"`
	Properties *ListAllDentriesResponseBodyDentriesProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string                                       `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	Thumbnail     *ListAllDentriesResponseBodyDentriesThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListAllDentriesResponseBodyDentries) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesResponseBodyDentries) GoString() string {
	return s.String()
}

func (s *ListAllDentriesResponseBodyDentries) SetAppProperties(v map[string][]*DentriesAppPropertiesValue) *ListAllDentriesResponseBodyDentries {
	s.AppProperties = v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetCreateTime(v string) *ListAllDentriesResponseBodyDentries {
	s.CreateTime = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetCreatorId(v string) *ListAllDentriesResponseBodyDentries {
	s.CreatorId = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetExtension(v string) *ListAllDentriesResponseBodyDentries {
	s.Extension = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetId(v string) *ListAllDentriesResponseBodyDentries {
	s.Id = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetModifiedTime(v string) *ListAllDentriesResponseBodyDentries {
	s.ModifiedTime = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetModifierId(v string) *ListAllDentriesResponseBodyDentries {
	s.ModifierId = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetName(v string) *ListAllDentriesResponseBodyDentries {
	s.Name = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetParentId(v string) *ListAllDentriesResponseBodyDentries {
	s.ParentId = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetPartitionType(v string) *ListAllDentriesResponseBodyDentries {
	s.PartitionType = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetPath(v string) *ListAllDentriesResponseBodyDentries {
	s.Path = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetProperties(v *ListAllDentriesResponseBodyDentriesProperties) *ListAllDentriesResponseBodyDentries {
	s.Properties = v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetSize(v int64) *ListAllDentriesResponseBodyDentries {
	s.Size = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetSpaceId(v string) *ListAllDentriesResponseBodyDentries {
	s.SpaceId = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetStatus(v string) *ListAllDentriesResponseBodyDentries {
	s.Status = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetStorageDriver(v string) *ListAllDentriesResponseBodyDentries {
	s.StorageDriver = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetThumbnail(v *ListAllDentriesResponseBodyDentriesThumbnail) *ListAllDentriesResponseBodyDentries {
	s.Thumbnail = v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetType(v string) *ListAllDentriesResponseBodyDentries {
	s.Type = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetUuid(v string) *ListAllDentriesResponseBodyDentries {
	s.Uuid = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentries) SetVersion(v int64) *ListAllDentriesResponseBodyDentries {
	s.Version = &v
	return s
}

type ListAllDentriesResponseBodyDentriesProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s ListAllDentriesResponseBodyDentriesProperties) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesResponseBodyDentriesProperties) GoString() string {
	return s.String()
}

func (s *ListAllDentriesResponseBodyDentriesProperties) SetReadOnly(v bool) *ListAllDentriesResponseBodyDentriesProperties {
	s.ReadOnly = &v
	return s
}

type ListAllDentriesResponseBodyDentriesThumbnail struct {
	// example:
	//
	// 64
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// 64
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s ListAllDentriesResponseBodyDentriesThumbnail) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesResponseBodyDentriesThumbnail) GoString() string {
	return s.String()
}

func (s *ListAllDentriesResponseBodyDentriesThumbnail) SetHeight(v int32) *ListAllDentriesResponseBodyDentriesThumbnail {
	s.Height = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentriesThumbnail) SetUrl(v string) *ListAllDentriesResponseBodyDentriesThumbnail {
	s.Url = &v
	return s
}

func (s *ListAllDentriesResponseBodyDentriesThumbnail) SetWidth(v int32) *ListAllDentriesResponseBodyDentriesThumbnail {
	s.Width = &v
	return s
}

type ListAllDentriesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllDentriesResponse) GoString() string {
	return s.String()
}

func (s *ListAllDentriesResponse) SetHeaders(v map[string]*string) *ListAllDentriesResponse {
	s.Headers = v
	return s
}

func (s *ListAllDentriesResponse) SetStatusCode(v int32) *ListAllDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllDentriesResponse) SetBody(v *ListAllDentriesResponseBody) *ListAllDentriesResponse {
	s.Body = v
	return s
}

type ListDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesHeaders) GoString() string {
	return s.String()
}

func (s *ListDentriesHeaders) SetCommonHeaders(v map[string]*string) *ListDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *ListDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListDentriesRequest struct {
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// true
	WithThumbnail *bool `json:"withThumbnail,omitempty" xml:"withThumbnail,omitempty"`
}

func (s ListDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesRequest) GoString() string {
	return s.String()
}

func (s *ListDentriesRequest) SetMaxResults(v int32) *ListDentriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDentriesRequest) SetNextToken(v string) *ListDentriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDentriesRequest) SetOrder(v string) *ListDentriesRequest {
	s.Order = &v
	return s
}

func (s *ListDentriesRequest) SetOrderBy(v string) *ListDentriesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDentriesRequest) SetParentId(v string) *ListDentriesRequest {
	s.ParentId = &v
	return s
}

func (s *ListDentriesRequest) SetUnionId(v string) *ListDentriesRequest {
	s.UnionId = &v
	return s
}

func (s *ListDentriesRequest) SetWithThumbnail(v bool) *ListDentriesRequest {
	s.WithThumbnail = &v
	return s
}

type ListDentriesResponseBody struct {
	Dentries []*ListDentriesResponseBodyDentries `json:"dentries,omitempty" xml:"dentries,omitempty" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBody) SetDentries(v []*ListDentriesResponseBodyDentries) *ListDentriesResponseBody {
	s.Dentries = v
	return s
}

func (s *ListDentriesResponseBody) SetNextToken(v string) *ListDentriesResponseBody {
	s.NextToken = &v
	return s
}

type ListDentriesResponseBodyDentries struct {
	AppProperties map[string][]*DentriesAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// example:
	//
	// DOCUMENT
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                     `json:"path,omitempty" xml:"path,omitempty"`
	Properties *ListDentriesResponseBodyDentriesProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string                                    `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	Thumbnail     *ListDentriesResponseBodyDentriesThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
	// example:
	//
	// FILE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListDentriesResponseBodyDentries) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponseBodyDentries) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBodyDentries) SetAppProperties(v map[string][]*DentriesAppPropertiesValue) *ListDentriesResponseBodyDentries {
	s.AppProperties = v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetCategory(v string) *ListDentriesResponseBodyDentries {
	s.Category = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetCreateTime(v string) *ListDentriesResponseBodyDentries {
	s.CreateTime = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetCreatorId(v string) *ListDentriesResponseBodyDentries {
	s.CreatorId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetExtension(v string) *ListDentriesResponseBodyDentries {
	s.Extension = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetId(v string) *ListDentriesResponseBodyDentries {
	s.Id = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetModifiedTime(v string) *ListDentriesResponseBodyDentries {
	s.ModifiedTime = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetModifierId(v string) *ListDentriesResponseBodyDentries {
	s.ModifierId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetName(v string) *ListDentriesResponseBodyDentries {
	s.Name = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetParentId(v string) *ListDentriesResponseBodyDentries {
	s.ParentId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetPartitionType(v string) *ListDentriesResponseBodyDentries {
	s.PartitionType = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetPath(v string) *ListDentriesResponseBodyDentries {
	s.Path = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetProperties(v *ListDentriesResponseBodyDentriesProperties) *ListDentriesResponseBodyDentries {
	s.Properties = v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetSize(v int64) *ListDentriesResponseBodyDentries {
	s.Size = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetSpaceId(v string) *ListDentriesResponseBodyDentries {
	s.SpaceId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetStatus(v string) *ListDentriesResponseBodyDentries {
	s.Status = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetStorageDriver(v string) *ListDentriesResponseBodyDentries {
	s.StorageDriver = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetThumbnail(v *ListDentriesResponseBodyDentriesThumbnail) *ListDentriesResponseBodyDentries {
	s.Thumbnail = v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetType(v string) *ListDentriesResponseBodyDentries {
	s.Type = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetUuid(v string) *ListDentriesResponseBodyDentries {
	s.Uuid = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetVersion(v int64) *ListDentriesResponseBodyDentries {
	s.Version = &v
	return s
}

type ListDentriesResponseBodyDentriesProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s ListDentriesResponseBodyDentriesProperties) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponseBodyDentriesProperties) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBodyDentriesProperties) SetReadOnly(v bool) *ListDentriesResponseBodyDentriesProperties {
	s.ReadOnly = &v
	return s
}

type ListDentriesResponseBodyDentriesThumbnail struct {
	// example:
	//
	// 64
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// 64
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s ListDentriesResponseBodyDentriesThumbnail) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponseBodyDentriesThumbnail) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBodyDentriesThumbnail) SetHeight(v int32) *ListDentriesResponseBodyDentriesThumbnail {
	s.Height = &v
	return s
}

func (s *ListDentriesResponseBodyDentriesThumbnail) SetUrl(v string) *ListDentriesResponseBodyDentriesThumbnail {
	s.Url = &v
	return s
}

func (s *ListDentriesResponseBodyDentriesThumbnail) SetWidth(v int32) *ListDentriesResponseBodyDentriesThumbnail {
	s.Width = &v
	return s
}

type ListDentriesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponse) GoString() string {
	return s.String()
}

func (s *ListDentriesResponse) SetHeaders(v map[string]*string) *ListDentriesResponse {
	s.Headers = v
	return s
}

func (s *ListDentriesResponse) SetStatusCode(v int32) *ListDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDentriesResponse) SetBody(v *ListDentriesResponseBody) *ListDentriesResponse {
	s.Body = v
	return s
}

type ListDentryVersionsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListDentryVersionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDentryVersionsHeaders) GoString() string {
	return s.String()
}

func (s *ListDentryVersionsHeaders) SetCommonHeaders(v map[string]*string) *ListDentryVersionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDentryVersionsHeaders) SetXAcsDingtalkAccessToken(v string) *ListDentryVersionsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListDentryVersionsRequest struct {
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListDentryVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDentryVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListDentryVersionsRequest) SetMaxResults(v int32) *ListDentryVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDentryVersionsRequest) SetNextToken(v string) *ListDentryVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDentryVersionsRequest) SetUnionId(v string) *ListDentryVersionsRequest {
	s.UnionId = &v
	return s
}

type ListDentryVersionsResponseBody struct {
	Dentries []*ListDentryVersionsResponseBodyDentries `json:"dentries,omitempty" xml:"dentries,omitempty" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListDentryVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDentryVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDentryVersionsResponseBody) SetDentries(v []*ListDentryVersionsResponseBodyDentries) *ListDentryVersionsResponseBody {
	s.Dentries = v
	return s
}

func (s *ListDentryVersionsResponseBody) SetNextToken(v string) *ListDentryVersionsResponseBody {
	s.NextToken = &v
	return s
}

type ListDentryVersionsResponseBodyDentries struct {
	AppProperties map[string][]*DentriesAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                           `json:"path,omitempty" xml:"path,omitempty"`
	Properties *ListDentryVersionsResponseBodyDentriesProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListDentryVersionsResponseBodyDentries) String() string {
	return tea.Prettify(s)
}

func (s ListDentryVersionsResponseBodyDentries) GoString() string {
	return s.String()
}

func (s *ListDentryVersionsResponseBodyDentries) SetAppProperties(v map[string][]*DentriesAppPropertiesValue) *ListDentryVersionsResponseBodyDentries {
	s.AppProperties = v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetCreateTime(v string) *ListDentryVersionsResponseBodyDentries {
	s.CreateTime = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetCreatorId(v string) *ListDentryVersionsResponseBodyDentries {
	s.CreatorId = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetExtension(v string) *ListDentryVersionsResponseBodyDentries {
	s.Extension = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetId(v string) *ListDentryVersionsResponseBodyDentries {
	s.Id = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetModifiedTime(v string) *ListDentryVersionsResponseBodyDentries {
	s.ModifiedTime = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetModifierId(v string) *ListDentryVersionsResponseBodyDentries {
	s.ModifierId = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetName(v string) *ListDentryVersionsResponseBodyDentries {
	s.Name = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetParentId(v string) *ListDentryVersionsResponseBodyDentries {
	s.ParentId = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetPartitionType(v string) *ListDentryVersionsResponseBodyDentries {
	s.PartitionType = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetPath(v string) *ListDentryVersionsResponseBodyDentries {
	s.Path = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetProperties(v *ListDentryVersionsResponseBodyDentriesProperties) *ListDentryVersionsResponseBodyDentries {
	s.Properties = v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetSize(v int64) *ListDentryVersionsResponseBodyDentries {
	s.Size = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetSpaceId(v string) *ListDentryVersionsResponseBodyDentries {
	s.SpaceId = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetStatus(v string) *ListDentryVersionsResponseBodyDentries {
	s.Status = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetStorageDriver(v string) *ListDentryVersionsResponseBodyDentries {
	s.StorageDriver = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetType(v string) *ListDentryVersionsResponseBodyDentries {
	s.Type = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetUuid(v string) *ListDentryVersionsResponseBodyDentries {
	s.Uuid = &v
	return s
}

func (s *ListDentryVersionsResponseBodyDentries) SetVersion(v int64) *ListDentryVersionsResponseBodyDentries {
	s.Version = &v
	return s
}

type ListDentryVersionsResponseBodyDentriesProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s ListDentryVersionsResponseBodyDentriesProperties) String() string {
	return tea.Prettify(s)
}

func (s ListDentryVersionsResponseBodyDentriesProperties) GoString() string {
	return s.String()
}

func (s *ListDentryVersionsResponseBodyDentriesProperties) SetReadOnly(v bool) *ListDentryVersionsResponseBodyDentriesProperties {
	s.ReadOnly = &v
	return s
}

type ListDentryVersionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDentryVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDentryVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDentryVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListDentryVersionsResponse) SetHeaders(v map[string]*string) *ListDentryVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListDentryVersionsResponse) SetStatusCode(v int32) *ListDentryVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDentryVersionsResponse) SetBody(v *ListDentryVersionsResponseBody) *ListDentryVersionsResponse {
	s.Body = v
	return s
}

type ListPermissionsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListPermissionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsHeaders) GoString() string {
	return s.String()
}

func (s *ListPermissionsHeaders) SetCommonHeaders(v map[string]*string) *ListPermissionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListPermissionsHeaders) SetXAcsDingtalkAccessToken(v string) *ListPermissionsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListPermissionsRequest struct {
	Option *ListPermissionsRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequest) SetOption(v *ListPermissionsRequestOption) *ListPermissionsRequest {
	s.Option = v
	return s
}

func (s *ListPermissionsRequest) SetUnionId(v string) *ListPermissionsRequest {
	s.UnionId = &v
	return s
}

type ListPermissionsRequestOption struct {
	FilterRoleIds []*string `json:"filterRoleIds,omitempty" xml:"filterRoleIds,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListPermissionsRequestOption) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsRequestOption) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequestOption) SetFilterRoleIds(v []*string) *ListPermissionsRequestOption {
	s.FilterRoleIds = v
	return s
}

func (s *ListPermissionsRequestOption) SetMaxResults(v int32) *ListPermissionsRequestOption {
	s.MaxResults = &v
	return s
}

func (s *ListPermissionsRequestOption) SetNextToken(v string) *ListPermissionsRequestOption {
	s.NextToken = &v
	return s
}

type ListPermissionsResponseBody struct {
	// example:
	//
	// next_token
	NextToken   *string                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Permissions []*ListPermissionsResponseBodyPermissions `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBody) SetNextToken(v string) *ListPermissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPermissionsResponseBody) SetPermissions(v []*ListPermissionsResponseBodyPermissions) *ListPermissionsResponseBody {
	s.Permissions = v
	return s
}

type ListPermissionsResponseBodyPermissions struct {
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// dentry_id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// 3600
	Duration *int64                                        `json:"duration,omitempty" xml:"duration,omitempty"`
	Member   *ListPermissionsResponseBodyPermissionsMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// operator_id
	OperatorId *string                                     `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Role       *ListPermissionsResponseBodyPermissionsRole `json:"role,omitempty" xml:"role,omitempty" type:"Struct"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s ListPermissionsResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissions) SetCreateTime(v string) *ListPermissionsResponseBodyPermissions {
	s.CreateTime = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetDentryId(v string) *ListPermissionsResponseBodyPermissions {
	s.DentryId = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetDuration(v int64) *ListPermissionsResponseBodyPermissions {
	s.Duration = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetMember(v *ListPermissionsResponseBodyPermissionsMember) *ListPermissionsResponseBodyPermissions {
	s.Member = v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetModifiedTime(v string) *ListPermissionsResponseBodyPermissions {
	s.ModifiedTime = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetOperatorId(v string) *ListPermissionsResponseBodyPermissions {
	s.OperatorId = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetRole(v *ListPermissionsResponseBodyPermissionsRole) *ListPermissionsResponseBodyPermissions {
	s.Role = v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetSpaceId(v string) *ListPermissionsResponseBodyPermissions {
	s.SpaceId = &v
	return s
}

type ListPermissionsResponseBodyPermissionsMember struct {
	// example:
	//
	// corp_id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// member_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// USER
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPermissionsResponseBodyPermissionsMember) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissionsMember) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissionsMember) SetCorpId(v string) *ListPermissionsResponseBodyPermissionsMember {
	s.CorpId = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsMember) SetId(v string) *ListPermissionsResponseBodyPermissionsMember {
	s.Id = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsMember) SetType(v string) *ListPermissionsResponseBodyPermissionsMember {
	s.Type = &v
	return s
}

type ListPermissionsResponseBodyPermissionsRole struct {
	// example:
	//
	// MANAGER
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// MANAGER
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPermissionsResponseBodyPermissionsRole) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissionsRole) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissionsRole) SetId(v string) *ListPermissionsResponseBodyPermissionsRole {
	s.Id = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsRole) SetName(v string) *ListPermissionsResponseBodyPermissionsRole {
	s.Name = &v
	return s
}

type ListPermissionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponse) SetHeaders(v map[string]*string) *ListPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionsResponse) SetStatusCode(v int32) *ListPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionsResponse) SetBody(v *ListPermissionsResponseBody) *ListPermissionsResponse {
	s.Body = v
	return s
}

type ListRecycleItemsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListRecycleItemsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleItemsHeaders) GoString() string {
	return s.String()
}

func (s *ListRecycleItemsHeaders) SetCommonHeaders(v map[string]*string) *ListRecycleItemsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRecycleItemsHeaders) SetXAcsDingtalkAccessToken(v string) *ListRecycleItemsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListRecycleItemsRequest struct {
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListRecycleItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleItemsRequest) GoString() string {
	return s.String()
}

func (s *ListRecycleItemsRequest) SetMaxResults(v int32) *ListRecycleItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecycleItemsRequest) SetNextToken(v string) *ListRecycleItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecycleItemsRequest) SetUnionId(v string) *ListRecycleItemsRequest {
	s.UnionId = &v
	return s
}

type ListRecycleItemsResponseBody struct {
	// example:
	//
	// next_token
	NextToken    *string                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RecycleItems []*ListRecycleItemsResponseBodyRecycleItems `json:"recycleItems,omitempty" xml:"recycleItems,omitempty" type:"Repeated"`
}

func (s ListRecycleItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecycleItemsResponseBody) SetNextToken(v string) *ListRecycleItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecycleItemsResponseBody) SetRecycleItems(v []*ListRecycleItemsResponseBodyRecycleItems) *ListRecycleItemsResponseBody {
	s.RecycleItems = v
	return s
}

type ListRecycleItemsResponseBodyRecycleItems struct {
	// example:
	//
	// dentry_id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// recycle_item_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// operator_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	OperatorTime *string `json:"operatorTime,omitempty" xml:"operatorTime,omitempty"`
	// example:
	//
	// dentry_name
	OriginalName *string `json:"originalName,omitempty" xml:"originalName,omitempty"`
	// example:
	//
	// dentry_path
	OriginalPath *string `json:"originalPath,omitempty" xml:"originalPath,omitempty"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRecycleItemsResponseBodyRecycleItems) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleItemsResponseBodyRecycleItems) GoString() string {
	return s.String()
}

func (s *ListRecycleItemsResponseBodyRecycleItems) SetDentryId(v string) *ListRecycleItemsResponseBodyRecycleItems {
	s.DentryId = &v
	return s
}

func (s *ListRecycleItemsResponseBodyRecycleItems) SetId(v string) *ListRecycleItemsResponseBodyRecycleItems {
	s.Id = &v
	return s
}

func (s *ListRecycleItemsResponseBodyRecycleItems) SetOperatorId(v string) *ListRecycleItemsResponseBodyRecycleItems {
	s.OperatorId = &v
	return s
}

func (s *ListRecycleItemsResponseBodyRecycleItems) SetOperatorTime(v string) *ListRecycleItemsResponseBodyRecycleItems {
	s.OperatorTime = &v
	return s
}

func (s *ListRecycleItemsResponseBodyRecycleItems) SetOriginalName(v string) *ListRecycleItemsResponseBodyRecycleItems {
	s.OriginalName = &v
	return s
}

func (s *ListRecycleItemsResponseBodyRecycleItems) SetOriginalPath(v string) *ListRecycleItemsResponseBodyRecycleItems {
	s.OriginalPath = &v
	return s
}

func (s *ListRecycleItemsResponseBodyRecycleItems) SetSize(v int64) *ListRecycleItemsResponseBodyRecycleItems {
	s.Size = &v
	return s
}

func (s *ListRecycleItemsResponseBodyRecycleItems) SetSpaceId(v string) *ListRecycleItemsResponseBodyRecycleItems {
	s.SpaceId = &v
	return s
}

func (s *ListRecycleItemsResponseBodyRecycleItems) SetType(v string) *ListRecycleItemsResponseBodyRecycleItems {
	s.Type = &v
	return s
}

type ListRecycleItemsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecycleItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecycleItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleItemsResponse) GoString() string {
	return s.String()
}

func (s *ListRecycleItemsResponse) SetHeaders(v map[string]*string) *ListRecycleItemsResponse {
	s.Headers = v
	return s
}

func (s *ListRecycleItemsResponse) SetStatusCode(v int32) *ListRecycleItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecycleItemsResponse) SetBody(v *ListRecycleItemsResponseBody) *ListRecycleItemsResponse {
	s.Body = v
	return s
}

type MoveDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MoveDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s MoveDentriesHeaders) GoString() string {
	return s.String()
}

func (s *MoveDentriesHeaders) SetCommonHeaders(v map[string]*string) *MoveDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MoveDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *MoveDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MoveDentriesRequest struct {
	// This parameter is required.
	DentryIds []*string                  `json:"dentryIds,omitempty" xml:"dentryIds,omitempty" type:"Repeated"`
	Option    *MoveDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// target_folder_id
	TargetFolderId *string `json:"targetFolderId,omitempty" xml:"targetFolderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// target_space_id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s MoveDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveDentriesRequest) GoString() string {
	return s.String()
}

func (s *MoveDentriesRequest) SetDentryIds(v []*string) *MoveDentriesRequest {
	s.DentryIds = v
	return s
}

func (s *MoveDentriesRequest) SetOption(v *MoveDentriesRequestOption) *MoveDentriesRequest {
	s.Option = v
	return s
}

func (s *MoveDentriesRequest) SetTargetFolderId(v string) *MoveDentriesRequest {
	s.TargetFolderId = &v
	return s
}

func (s *MoveDentriesRequest) SetTargetSpaceId(v string) *MoveDentriesRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *MoveDentriesRequest) SetUnionId(v string) *MoveDentriesRequest {
	s.UnionId = &v
	return s
}

type MoveDentriesRequestOption struct {
	// example:
	//
	// AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
	// example:
	//
	// true
	PreservePermissions *bool `json:"preservePermissions,omitempty" xml:"preservePermissions,omitempty"`
}

func (s MoveDentriesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s MoveDentriesRequestOption) GoString() string {
	return s.String()
}

func (s *MoveDentriesRequestOption) SetConflictStrategy(v string) *MoveDentriesRequestOption {
	s.ConflictStrategy = &v
	return s
}

func (s *MoveDentriesRequestOption) SetPreservePermissions(v bool) *MoveDentriesRequestOption {
	s.PreservePermissions = &v
	return s
}

type MoveDentriesResponseBody struct {
	ResultItems []*MoveDentriesResponseBodyResultItems `json:"resultItems,omitempty" xml:"resultItems,omitempty" type:"Repeated"`
}

func (s MoveDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *MoveDentriesResponseBody) SetResultItems(v []*MoveDentriesResponseBodyResultItems) *MoveDentriesResponseBody {
	s.ResultItems = v
	return s
}

type MoveDentriesResponseBodyResultItems struct {
	// example:
	//
	// true
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// example:
	//
	// dentry_id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// permissionDenied
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// target_dentry_id
	TargetDentryId *string `json:"targetDentryId,omitempty" xml:"targetDentryId,omitempty"`
	// example:
	//
	// target_space_id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// example:
	//
	// task_id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s MoveDentriesResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s MoveDentriesResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *MoveDentriesResponseBodyResultItems) SetAsync(v bool) *MoveDentriesResponseBodyResultItems {
	s.Async = &v
	return s
}

func (s *MoveDentriesResponseBodyResultItems) SetDentryId(v string) *MoveDentriesResponseBodyResultItems {
	s.DentryId = &v
	return s
}

func (s *MoveDentriesResponseBodyResultItems) SetErrorCode(v string) *MoveDentriesResponseBodyResultItems {
	s.ErrorCode = &v
	return s
}

func (s *MoveDentriesResponseBodyResultItems) SetSpaceId(v string) *MoveDentriesResponseBodyResultItems {
	s.SpaceId = &v
	return s
}

func (s *MoveDentriesResponseBodyResultItems) SetSuccess(v bool) *MoveDentriesResponseBodyResultItems {
	s.Success = &v
	return s
}

func (s *MoveDentriesResponseBodyResultItems) SetTargetDentryId(v string) *MoveDentriesResponseBodyResultItems {
	s.TargetDentryId = &v
	return s
}

func (s *MoveDentriesResponseBodyResultItems) SetTargetSpaceId(v string) *MoveDentriesResponseBodyResultItems {
	s.TargetSpaceId = &v
	return s
}

func (s *MoveDentriesResponseBodyResultItems) SetTaskId(v string) *MoveDentriesResponseBodyResultItems {
	s.TaskId = &v
	return s
}

type MoveDentriesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveDentriesResponse) GoString() string {
	return s.String()
}

func (s *MoveDentriesResponse) SetHeaders(v map[string]*string) *MoveDentriesResponse {
	s.Headers = v
	return s
}

func (s *MoveDentriesResponse) SetStatusCode(v int32) *MoveDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveDentriesResponse) SetBody(v *MoveDentriesResponseBody) *MoveDentriesResponse {
	s.Body = v
	return s
}

type MoveDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MoveDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryHeaders) GoString() string {
	return s.String()
}

func (s *MoveDentryHeaders) SetCommonHeaders(v map[string]*string) *MoveDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MoveDentryHeaders) SetXAcsDingtalkAccessToken(v string) *MoveDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MoveDentryRequest struct {
	Option *MoveDentryRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// target_folder_id
	TargetFolderId *string `json:"targetFolderId,omitempty" xml:"targetFolderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// target_space_id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s MoveDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryRequest) GoString() string {
	return s.String()
}

func (s *MoveDentryRequest) SetOption(v *MoveDentryRequestOption) *MoveDentryRequest {
	s.Option = v
	return s
}

func (s *MoveDentryRequest) SetTargetFolderId(v string) *MoveDentryRequest {
	s.TargetFolderId = &v
	return s
}

func (s *MoveDentryRequest) SetTargetSpaceId(v string) *MoveDentryRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *MoveDentryRequest) SetUnionId(v string) *MoveDentryRequest {
	s.UnionId = &v
	return s
}

type MoveDentryRequestOption struct {
	// example:
	//
	// AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
	// example:
	//
	// true
	PresevePermissions *bool `json:"presevePermissions,omitempty" xml:"presevePermissions,omitempty"`
}

func (s MoveDentryRequestOption) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryRequestOption) GoString() string {
	return s.String()
}

func (s *MoveDentryRequestOption) SetConflictStrategy(v string) *MoveDentryRequestOption {
	s.ConflictStrategy = &v
	return s
}

func (s *MoveDentryRequestOption) SetPresevePermissions(v bool) *MoveDentryRequestOption {
	s.PresevePermissions = &v
	return s
}

type MoveDentryResponseBody struct {
	// example:
	//
	// true
	Async  *bool                         `json:"async,omitempty" xml:"async,omitempty"`
	Dentry *MoveDentryResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
	// example:
	//
	// task_id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s MoveDentryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryResponseBody) GoString() string {
	return s.String()
}

func (s *MoveDentryResponseBody) SetAsync(v bool) *MoveDentryResponseBody {
	s.Async = &v
	return s
}

func (s *MoveDentryResponseBody) SetDentry(v *MoveDentryResponseBodyDentry) *MoveDentryResponseBody {
	s.Dentry = v
	return s
}

func (s *MoveDentryResponseBody) SetTaskId(v string) *MoveDentryResponseBody {
	s.TaskId = &v
	return s
}

type MoveDentryResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                 `json:"path,omitempty" xml:"path,omitempty"`
	Properties *MoveDentryResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s MoveDentryResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *MoveDentryResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *MoveDentryResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetCreateTime(v string) *MoveDentryResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetCreatorId(v string) *MoveDentryResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetExtension(v string) *MoveDentryResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetId(v string) *MoveDentryResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetModifiedTime(v string) *MoveDentryResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetModifierId(v string) *MoveDentryResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetName(v string) *MoveDentryResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetParentId(v string) *MoveDentryResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetPartitionType(v string) *MoveDentryResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetPath(v string) *MoveDentryResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetProperties(v *MoveDentryResponseBodyDentryProperties) *MoveDentryResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetSize(v int64) *MoveDentryResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetSpaceId(v string) *MoveDentryResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetStatus(v string) *MoveDentryResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetStorageDriver(v string) *MoveDentryResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetType(v string) *MoveDentryResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetUuid(v string) *MoveDentryResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *MoveDentryResponseBodyDentry) SetVersion(v int64) *MoveDentryResponseBodyDentry {
	s.Version = &v
	return s
}

type MoveDentryResponseBodyDentryProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s MoveDentryResponseBodyDentryProperties) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *MoveDentryResponseBodyDentryProperties) SetReadOnly(v bool) *MoveDentryResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

type MoveDentryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveDentryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryResponse) GoString() string {
	return s.String()
}

func (s *MoveDentryResponse) SetHeaders(v map[string]*string) *MoveDentryResponse {
	s.Headers = v
	return s
}

func (s *MoveDentryResponse) SetStatusCode(v int32) *MoveDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveDentryResponse) SetBody(v *MoveDentryResponseBody) *MoveDentryResponse {
	s.Body = v
	return s
}

type RefreshWebOfficeTokenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RefreshWebOfficeTokenHeaders) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebOfficeTokenHeaders) GoString() string {
	return s.String()
}

func (s *RefreshWebOfficeTokenHeaders) SetCommonHeaders(v map[string]*string) *RefreshWebOfficeTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RefreshWebOfficeTokenHeaders) SetXAcsDingtalkAccessToken(v string) *RefreshWebOfficeTokenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RefreshWebOfficeTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// This parameter is required.
	WebOfficeAccessToken *string `json:"webOfficeAccessToken,omitempty" xml:"webOfficeAccessToken,omitempty"`
	// This parameter is required.
	WebOfficeRefreshToken *string `json:"webOfficeRefreshToken,omitempty" xml:"webOfficeRefreshToken,omitempty"`
}

func (s RefreshWebOfficeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebOfficeTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshWebOfficeTokenRequest) SetUnionId(v string) *RefreshWebOfficeTokenRequest {
	s.UnionId = &v
	return s
}

func (s *RefreshWebOfficeTokenRequest) SetWebOfficeAccessToken(v string) *RefreshWebOfficeTokenRequest {
	s.WebOfficeAccessToken = &v
	return s
}

func (s *RefreshWebOfficeTokenRequest) SetWebOfficeRefreshToken(v string) *RefreshWebOfficeTokenRequest {
	s.WebOfficeRefreshToken = &v
	return s
}

type RefreshWebOfficeTokenResponseBody struct {
	WebOfficeAccessToken  *string `json:"webOfficeAccessToken,omitempty" xml:"webOfficeAccessToken,omitempty"`
	WebOfficeRefreshToken *string `json:"webOfficeRefreshToken,omitempty" xml:"webOfficeRefreshToken,omitempty"`
}

func (s RefreshWebOfficeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebOfficeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshWebOfficeTokenResponseBody) SetWebOfficeAccessToken(v string) *RefreshWebOfficeTokenResponseBody {
	s.WebOfficeAccessToken = &v
	return s
}

func (s *RefreshWebOfficeTokenResponseBody) SetWebOfficeRefreshToken(v string) *RefreshWebOfficeTokenResponseBody {
	s.WebOfficeRefreshToken = &v
	return s
}

type RefreshWebOfficeTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshWebOfficeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshWebOfficeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebOfficeTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshWebOfficeTokenResponse) SetHeaders(v map[string]*string) *RefreshWebOfficeTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshWebOfficeTokenResponse) SetStatusCode(v int32) *RefreshWebOfficeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshWebOfficeTokenResponse) SetBody(v *RefreshWebOfficeTokenResponseBody) *RefreshWebOfficeTokenResponse {
	s.Body = v
	return s
}

type RegisterOpenInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterOpenInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterOpenInfoHeaders) GoString() string {
	return s.String()
}

func (s *RegisterOpenInfoHeaders) SetCommonHeaders(v map[string]*string) *RegisterOpenInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterOpenInfoHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterOpenInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterOpenInfoRequest struct {
	// This parameter is required.
	OpenInfos []*RegisterOpenInfoRequestOpenInfos `json:"openInfos,omitempty" xml:"openInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// DINGTALK
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RegisterOpenInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterOpenInfoRequest) GoString() string {
	return s.String()
}

func (s *RegisterOpenInfoRequest) SetOpenInfos(v []*RegisterOpenInfoRequestOpenInfos) *RegisterOpenInfoRequest {
	s.OpenInfos = v
	return s
}

func (s *RegisterOpenInfoRequest) SetProvider(v string) *RegisterOpenInfoRequest {
	s.Provider = &v
	return s
}

func (s *RegisterOpenInfoRequest) SetUnionId(v string) *RegisterOpenInfoRequest {
	s.UnionId = &v
	return s
}

type RegisterOpenInfoRequestOpenInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// PREVIEW
	OpenType *string `json:"openType,omitempty" xml:"openType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RegisterOpenInfoRequestOpenInfos) String() string {
	return tea.Prettify(s)
}

func (s RegisterOpenInfoRequestOpenInfos) GoString() string {
	return s.String()
}

func (s *RegisterOpenInfoRequestOpenInfos) SetOpenType(v string) *RegisterOpenInfoRequestOpenInfos {
	s.OpenType = &v
	return s
}

func (s *RegisterOpenInfoRequestOpenInfos) SetUrl(v string) *RegisterOpenInfoRequestOpenInfos {
	s.Url = &v
	return s
}

type RegisterOpenInfoResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RegisterOpenInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterOpenInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterOpenInfoResponseBody) SetSuccess(v bool) *RegisterOpenInfoResponseBody {
	s.Success = &v
	return s
}

type RegisterOpenInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterOpenInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterOpenInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterOpenInfoResponse) GoString() string {
	return s.String()
}

func (s *RegisterOpenInfoResponse) SetHeaders(v map[string]*string) *RegisterOpenInfoResponse {
	s.Headers = v
	return s
}

func (s *RegisterOpenInfoResponse) SetStatusCode(v int32) *RegisterOpenInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterOpenInfoResponse) SetBody(v *RegisterOpenInfoResponseBody) *RegisterOpenInfoResponse {
	s.Body = v
	return s
}

type RenameDentryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RenameDentryHeaders) String() string {
	return tea.Prettify(s)
}

func (s RenameDentryHeaders) GoString() string {
	return s.String()
}

func (s *RenameDentryHeaders) SetCommonHeaders(v map[string]*string) *RenameDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RenameDentryHeaders) SetXAcsDingtalkAccessToken(v string) *RenameDentryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RenameDentryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dentry_name
	NewName *string `json:"newName,omitempty" xml:"newName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RenameDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameDentryRequest) GoString() string {
	return s.String()
}

func (s *RenameDentryRequest) SetNewName(v string) *RenameDentryRequest {
	s.NewName = &v
	return s
}

func (s *RenameDentryRequest) SetUnionId(v string) *RenameDentryRequest {
	s.UnionId = &v
	return s
}

type RenameDentryResponseBody struct {
	Dentry *RenameDentryResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
}

func (s RenameDentryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameDentryResponseBody) GoString() string {
	return s.String()
}

func (s *RenameDentryResponseBody) SetDentry(v *RenameDentryResponseBodyDentry) *RenameDentryResponseBody {
	s.Dentry = v
	return s
}

type RenameDentryResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// parent_id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                   `json:"path,omitempty" xml:"path,omitempty"`
	Properties *RenameDentryResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s RenameDentryResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s RenameDentryResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *RenameDentryResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *RenameDentryResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetCreateTime(v string) *RenameDentryResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetCreatorId(v string) *RenameDentryResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetExtension(v string) *RenameDentryResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetId(v string) *RenameDentryResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetModifiedTime(v string) *RenameDentryResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetModifierId(v string) *RenameDentryResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetName(v string) *RenameDentryResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetParentId(v string) *RenameDentryResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetPartitionType(v string) *RenameDentryResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetPath(v string) *RenameDentryResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetProperties(v *RenameDentryResponseBodyDentryProperties) *RenameDentryResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetSize(v int64) *RenameDentryResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetSpaceId(v string) *RenameDentryResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetStatus(v string) *RenameDentryResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetStorageDriver(v string) *RenameDentryResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetType(v string) *RenameDentryResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetUuid(v string) *RenameDentryResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *RenameDentryResponseBodyDentry) SetVersion(v int64) *RenameDentryResponseBodyDentry {
	s.Version = &v
	return s
}

type RenameDentryResponseBodyDentryProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s RenameDentryResponseBodyDentryProperties) String() string {
	return tea.Prettify(s)
}

func (s RenameDentryResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *RenameDentryResponseBodyDentryProperties) SetReadOnly(v bool) *RenameDentryResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

type RenameDentryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameDentryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameDentryResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameDentryResponse) GoString() string {
	return s.String()
}

func (s *RenameDentryResponse) SetHeaders(v map[string]*string) *RenameDentryResponse {
	s.Headers = v
	return s
}

func (s *RenameDentryResponse) SetStatusCode(v int32) *RenameDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameDentryResponse) SetBody(v *RenameDentryResponseBody) *RenameDentryResponse {
	s.Body = v
	return s
}

type RestoreRecycleItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RestoreRecycleItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemHeaders) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemHeaders) SetCommonHeaders(v map[string]*string) *RestoreRecycleItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RestoreRecycleItemHeaders) SetXAcsDingtalkAccessToken(v string) *RestoreRecycleItemHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RestoreRecycleItemRequest struct {
	Option *RestoreRecycleItemRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RestoreRecycleItemRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemRequest) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemRequest) SetOption(v *RestoreRecycleItemRequestOption) *RestoreRecycleItemRequest {
	s.Option = v
	return s
}

func (s *RestoreRecycleItemRequest) SetUnionId(v string) *RestoreRecycleItemRequest {
	s.UnionId = &v
	return s
}

type RestoreRecycleItemRequestOption struct {
	// example:
	//
	// AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
}

func (s RestoreRecycleItemRequestOption) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemRequestOption) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemRequestOption) SetConflictStrategy(v string) *RestoreRecycleItemRequestOption {
	s.ConflictStrategy = &v
	return s
}

type RestoreRecycleItemResponseBody struct {
	// example:
	//
	// true
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// example:
	//
	// dentry_id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// task_id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RestoreRecycleItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemResponseBody) SetAsync(v bool) *RestoreRecycleItemResponseBody {
	s.Async = &v
	return s
}

func (s *RestoreRecycleItemResponseBody) SetDentryId(v string) *RestoreRecycleItemResponseBody {
	s.DentryId = &v
	return s
}

func (s *RestoreRecycleItemResponseBody) SetSpaceId(v string) *RestoreRecycleItemResponseBody {
	s.SpaceId = &v
	return s
}

func (s *RestoreRecycleItemResponseBody) SetTaskId(v string) *RestoreRecycleItemResponseBody {
	s.TaskId = &v
	return s
}

type RestoreRecycleItemResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreRecycleItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreRecycleItemResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemResponse) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemResponse) SetHeaders(v map[string]*string) *RestoreRecycleItemResponse {
	s.Headers = v
	return s
}

func (s *RestoreRecycleItemResponse) SetStatusCode(v int32) *RestoreRecycleItemResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreRecycleItemResponse) SetBody(v *RestoreRecycleItemResponseBody) *RestoreRecycleItemResponse {
	s.Body = v
	return s
}

type RestoreRecycleItemsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RestoreRecycleItemsHeaders) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemsHeaders) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemsHeaders) SetCommonHeaders(v map[string]*string) *RestoreRecycleItemsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RestoreRecycleItemsHeaders) SetXAcsDingtalkAccessToken(v string) *RestoreRecycleItemsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RestoreRecycleItemsRequest struct {
	Option *RestoreRecycleItemsRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	RecycleItemIds []*string `json:"recycleItemIds,omitempty" xml:"recycleItemIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RestoreRecycleItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemsRequest) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemsRequest) SetOption(v *RestoreRecycleItemsRequestOption) *RestoreRecycleItemsRequest {
	s.Option = v
	return s
}

func (s *RestoreRecycleItemsRequest) SetRecycleItemIds(v []*string) *RestoreRecycleItemsRequest {
	s.RecycleItemIds = v
	return s
}

func (s *RestoreRecycleItemsRequest) SetUnionId(v string) *RestoreRecycleItemsRequest {
	s.UnionId = &v
	return s
}

type RestoreRecycleItemsRequestOption struct {
	// example:
	//
	// AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
}

func (s RestoreRecycleItemsRequestOption) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemsRequestOption) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemsRequestOption) SetConflictStrategy(v string) *RestoreRecycleItemsRequestOption {
	s.ConflictStrategy = &v
	return s
}

type RestoreRecycleItemsResponseBody struct {
	ResultItems []*RestoreRecycleItemsResponseBodyResultItems `json:"resultItems,omitempty" xml:"resultItems,omitempty" type:"Repeated"`
}

func (s RestoreRecycleItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemsResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemsResponseBody) SetResultItems(v []*RestoreRecycleItemsResponseBodyResultItems) *RestoreRecycleItemsResponseBody {
	s.ResultItems = v
	return s
}

type RestoreRecycleItemsResponseBodyResultItems struct {
	// example:
	//
	// true
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// example:
	//
	// dentry_id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// permissionDenied
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// recyclebin_id
	RecycleBinId *string `json:"recycleBinId,omitempty" xml:"recycleBinId,omitempty"`
	// example:
	//
	// recycle_item_id
	RecycleItemId *string `json:"recycleItemId,omitempty" xml:"recycleItemId,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// task_id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RestoreRecycleItemsResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemsResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemsResponseBodyResultItems) SetAsync(v bool) *RestoreRecycleItemsResponseBodyResultItems {
	s.Async = &v
	return s
}

func (s *RestoreRecycleItemsResponseBodyResultItems) SetDentryId(v string) *RestoreRecycleItemsResponseBodyResultItems {
	s.DentryId = &v
	return s
}

func (s *RestoreRecycleItemsResponseBodyResultItems) SetErrorCode(v string) *RestoreRecycleItemsResponseBodyResultItems {
	s.ErrorCode = &v
	return s
}

func (s *RestoreRecycleItemsResponseBodyResultItems) SetRecycleBinId(v string) *RestoreRecycleItemsResponseBodyResultItems {
	s.RecycleBinId = &v
	return s
}

func (s *RestoreRecycleItemsResponseBodyResultItems) SetRecycleItemId(v string) *RestoreRecycleItemsResponseBodyResultItems {
	s.RecycleItemId = &v
	return s
}

func (s *RestoreRecycleItemsResponseBodyResultItems) SetSpaceId(v string) *RestoreRecycleItemsResponseBodyResultItems {
	s.SpaceId = &v
	return s
}

func (s *RestoreRecycleItemsResponseBodyResultItems) SetSuccess(v bool) *RestoreRecycleItemsResponseBodyResultItems {
	s.Success = &v
	return s
}

func (s *RestoreRecycleItemsResponseBodyResultItems) SetTaskId(v string) *RestoreRecycleItemsResponseBodyResultItems {
	s.TaskId = &v
	return s
}

type RestoreRecycleItemsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreRecycleItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreRecycleItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemsResponse) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemsResponse) SetHeaders(v map[string]*string) *RestoreRecycleItemsResponse {
	s.Headers = v
	return s
}

func (s *RestoreRecycleItemsResponse) SetStatusCode(v int32) *RestoreRecycleItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreRecycleItemsResponse) SetBody(v *RestoreRecycleItemsResponseBody) *RestoreRecycleItemsResponse {
	s.Body = v
	return s
}

type RevertDentryVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RevertDentryVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s RevertDentryVersionHeaders) GoString() string {
	return s.String()
}

func (s *RevertDentryVersionHeaders) SetCommonHeaders(v map[string]*string) *RevertDentryVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RevertDentryVersionHeaders) SetXAcsDingtalkAccessToken(v string) *RevertDentryVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RevertDentryVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RevertDentryVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertDentryVersionRequest) GoString() string {
	return s.String()
}

func (s *RevertDentryVersionRequest) SetUnionId(v string) *RevertDentryVersionRequest {
	s.UnionId = &v
	return s
}

type RevertDentryVersionResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RevertDentryVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevertDentryVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RevertDentryVersionResponseBody) SetSuccess(v bool) *RevertDentryVersionResponseBody {
	s.Success = &v
	return s
}

type RevertDentryVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertDentryVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertDentryVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s RevertDentryVersionResponse) GoString() string {
	return s.String()
}

func (s *RevertDentryVersionResponse) SetHeaders(v map[string]*string) *RevertDentryVersionResponse {
	s.Headers = v
	return s
}

func (s *RevertDentryVersionResponse) SetStatusCode(v int32) *RevertDentryVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertDentryVersionResponse) SetBody(v *RevertDentryVersionResponseBody) *RevertDentryVersionResponse {
	s.Body = v
	return s
}

type SubscribeEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SubscribeEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubscribeEventHeaders) GoString() string {
	return s.String()
}

func (s *SubscribeEventHeaders) SetCommonHeaders(v map[string]*string) *SubscribeEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubscribeEventHeaders) SetXAcsDingtalkAccessToken(v string) *SubscribeEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SubscribeEventRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SPACE
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scope_id
	ScopeId *string `json:"scopeId,omitempty" xml:"scopeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SubscribeEventRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeEventRequest) GoString() string {
	return s.String()
}

func (s *SubscribeEventRequest) SetScope(v string) *SubscribeEventRequest {
	s.Scope = &v
	return s
}

func (s *SubscribeEventRequest) SetScopeId(v string) *SubscribeEventRequest {
	s.ScopeId = &v
	return s
}

func (s *SubscribeEventRequest) SetUnionId(v string) *SubscribeEventRequest {
	s.UnionId = &v
	return s
}

type SubscribeEventResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubscribeEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscribeEventResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeEventResponseBody) SetSuccess(v bool) *SubscribeEventResponseBody {
	s.Success = &v
	return s
}

type SubscribeEventResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscribeEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscribeEventResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeEventResponse) GoString() string {
	return s.String()
}

func (s *SubscribeEventResponse) SetHeaders(v map[string]*string) *SubscribeEventResponse {
	s.Headers = v
	return s
}

func (s *SubscribeEventResponse) SetStatusCode(v int32) *SubscribeEventResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeEventResponse) SetBody(v *SubscribeEventResponseBody) *SubscribeEventResponse {
	s.Body = v
	return s
}

type UnsubscribeEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UnsubscribeEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeEventHeaders) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventHeaders) SetCommonHeaders(v map[string]*string) *UnsubscribeEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnsubscribeEventHeaders) SetXAcsDingtalkAccessToken(v string) *UnsubscribeEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UnsubscribeEventRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SPACE
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scope_id
	ScopeId *string `json:"scopeId,omitempty" xml:"scopeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UnsubscribeEventRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeEventRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventRequest) SetScope(v string) *UnsubscribeEventRequest {
	s.Scope = &v
	return s
}

func (s *UnsubscribeEventRequest) SetScopeId(v string) *UnsubscribeEventRequest {
	s.ScopeId = &v
	return s
}

func (s *UnsubscribeEventRequest) SetUnionId(v string) *UnsubscribeEventRequest {
	s.UnionId = &v
	return s
}

type UnsubscribeEventResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UnsubscribeEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeEventResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventResponseBody) SetSuccess(v bool) *UnsubscribeEventResponseBody {
	s.Success = &v
	return s
}

type UnsubscribeEventResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnsubscribeEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnsubscribeEventResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeEventResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventResponse) SetHeaders(v map[string]*string) *UnsubscribeEventResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeEventResponse) SetStatusCode(v int32) *UnsubscribeEventResponse {
	s.StatusCode = &v
	return s
}

func (s *UnsubscribeEventResponse) SetBody(v *UnsubscribeEventResponseBody) *UnsubscribeEventResponse {
	s.Body = v
	return s
}

type UpdateDentryAppPropertiesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateDentryAppPropertiesHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateDentryAppPropertiesHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDentryAppPropertiesHeaders) SetCommonHeaders(v map[string]*string) *UpdateDentryAppPropertiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDentryAppPropertiesHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateDentryAppPropertiesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateDentryAppPropertiesRequest struct {
	// This parameter is required.
	AppProperties []*UpdateDentryAppPropertiesRequestAppProperties `json:"appProperties,omitempty" xml:"appProperties,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdateDentryAppPropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDentryAppPropertiesRequest) GoString() string {
	return s.String()
}

func (s *UpdateDentryAppPropertiesRequest) SetAppProperties(v []*UpdateDentryAppPropertiesRequestAppProperties) *UpdateDentryAppPropertiesRequest {
	s.AppProperties = v
	return s
}

func (s *UpdateDentryAppPropertiesRequest) SetUnionId(v string) *UpdateDentryAppPropertiesRequest {
	s.UnionId = &v
	return s
}

type UpdateDentryAppPropertiesRequestAppProperties struct {
	// This parameter is required.
	//
	// example:
	//
	// property_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// property_value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s UpdateDentryAppPropertiesRequestAppProperties) String() string {
	return tea.Prettify(s)
}

func (s UpdateDentryAppPropertiesRequestAppProperties) GoString() string {
	return s.String()
}

func (s *UpdateDentryAppPropertiesRequestAppProperties) SetName(v string) *UpdateDentryAppPropertiesRequestAppProperties {
	s.Name = &v
	return s
}

func (s *UpdateDentryAppPropertiesRequestAppProperties) SetValue(v string) *UpdateDentryAppPropertiesRequestAppProperties {
	s.Value = &v
	return s
}

func (s *UpdateDentryAppPropertiesRequestAppProperties) SetVisibility(v string) *UpdateDentryAppPropertiesRequestAppProperties {
	s.Visibility = &v
	return s
}

type UpdateDentryAppPropertiesResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDentryAppPropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDentryAppPropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDentryAppPropertiesResponseBody) SetSuccess(v bool) *UpdateDentryAppPropertiesResponseBody {
	s.Success = &v
	return s
}

type UpdateDentryAppPropertiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDentryAppPropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDentryAppPropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDentryAppPropertiesResponse) GoString() string {
	return s.String()
}

func (s *UpdateDentryAppPropertiesResponse) SetHeaders(v map[string]*string) *UpdateDentryAppPropertiesResponse {
	s.Headers = v
	return s
}

func (s *UpdateDentryAppPropertiesResponse) SetStatusCode(v int32) *UpdateDentryAppPropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDentryAppPropertiesResponse) SetBody(v *UpdateDentryAppPropertiesResponseBody) *UpdateDentryAppPropertiesResponse {
	s.Body = v
	return s
}

type UpdatePermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdatePermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePermissionHeaders) SetCommonHeaders(v map[string]*string) *UpdatePermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePermissionHeaders) SetXAcsDingtalkAccessToken(v string) *UpdatePermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdatePermissionRequest struct {
	// This parameter is required.
	Members []*UpdatePermissionRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	Option  *UpdatePermissionRequestOption    `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// MANAGER
	RoleId *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdatePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdatePermissionRequest) SetMembers(v []*UpdatePermissionRequestMembers) *UpdatePermissionRequest {
	s.Members = v
	return s
}

func (s *UpdatePermissionRequest) SetOption(v *UpdatePermissionRequestOption) *UpdatePermissionRequest {
	s.Option = v
	return s
}

func (s *UpdatePermissionRequest) SetRoleId(v string) *UpdatePermissionRequest {
	s.RoleId = &v
	return s
}

func (s *UpdatePermissionRequest) SetUnionId(v string) *UpdatePermissionRequest {
	s.UnionId = &v
	return s
}

type UpdatePermissionRequestMembers struct {
	// example:
	//
	// corp_id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// member_id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// USER
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePermissionRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdatePermissionRequestMembers) SetCorpId(v string) *UpdatePermissionRequestMembers {
	s.CorpId = &v
	return s
}

func (s *UpdatePermissionRequestMembers) SetId(v string) *UpdatePermissionRequestMembers {
	s.Id = &v
	return s
}

func (s *UpdatePermissionRequestMembers) SetType(v string) *UpdatePermissionRequestMembers {
	s.Type = &v
	return s
}

type UpdatePermissionRequestOption struct {
	// example:
	//
	// 3600
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
}

func (s UpdatePermissionRequestOption) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionRequestOption) GoString() string {
	return s.String()
}

func (s *UpdatePermissionRequestOption) SetDuration(v int64) *UpdatePermissionRequestOption {
	s.Duration = &v
	return s
}

type UpdatePermissionResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePermissionResponseBody) SetSuccess(v bool) *UpdatePermissionResponseBody {
	s.Success = &v
	return s
}

type UpdatePermissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdatePermissionResponse) SetHeaders(v map[string]*string) *UpdatePermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdatePermissionResponse) SetStatusCode(v int32) *UpdatePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePermissionResponse) SetBody(v *UpdatePermissionResponseBody) *UpdatePermissionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 添加文件夹
//
// @param request - AddFolderRequest
//
// @param headers - AddFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFolderResponse
func (client *Client) AddFolderWithOptions(spaceId *string, parentId *string, request *AddFolderRequest, headers *AddFolderHeaders, runtime *util.RuntimeOptions) (_result *AddFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFolder"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(parentId) + "/folders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFolderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加文件夹
//
// @param request - AddFolderRequest
//
// @return AddFolderResponse
func (client *Client) AddFolder(spaceId *string, parentId *string, request *AddFolderRequest) (_result *AddFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddFolderHeaders{}
	_result = &AddFolderResponse{}
	_body, _err := client.AddFolderWithOptions(spaceId, parentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加权限
//
// @param request - AddPermissionRequest
//
// @param headers - AddPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPermissionResponse
func (client *Client) AddPermissionWithOptions(spaceId *string, dentryId *string, request *AddPermissionRequest, headers *AddPermissionHeaders, runtime *util.RuntimeOptions) (_result *AddPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		body["roleId"] = request.RoleId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddPermission"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/permissions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加权限
//
// @param request - AddPermissionRequest
//
// @return AddPermissionResponse
func (client *Client) AddPermission(spaceId *string, dentryId *string, request *AddPermissionRequest) (_result *AddPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddPermissionHeaders{}
	_result = &AddPermissionResponse{}
	_body, _err := client.AddPermissionWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加空间
//
// @param request - AddSpaceRequest
//
// @param headers - AddSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSpaceResponse
func (client *Client) AddSpaceWithOptions(request *AddSpaceRequest, headers *AddSpaceHeaders, runtime *util.RuntimeOptions) (_result *AddSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSpace"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加空间
//
// @param request - AddSpaceRequest
//
// @return AddSpaceResponse
func (client *Client) AddSpace(request *AddSpaceRequest) (_result *AddSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddSpaceHeaders{}
	_result = &AddSpaceResponse{}
	_body, _err := client.AddSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取权限列表
//
// @param tmpReq - BatchQueryRolesRequest
//
// @param headers - BatchQueryRolesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryRolesResponse
func (client *Client) BatchQueryRolesWithOptions(spaceId *string, tmpReq *BatchQueryRolesRequest, headers *BatchQueryRolesHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryRolesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchQueryRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DentryIdList)) {
		request.DentryIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DentryIdList, tea.String("dentryIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryIdListShrink)) {
		query["dentryIdList"] = request.DentryIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchQueryRoles"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/permissions/roles/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryRolesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取权限列表
//
// @param request - BatchQueryRolesRequest
//
// @return BatchQueryRolesResponse
func (client *Client) BatchQueryRoles(spaceId *string, request *BatchQueryRolesRequest) (_result *BatchQueryRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryRolesHeaders{}
	_result = &BatchQueryRolesResponse{}
	_body, _err := client.BatchQueryRolesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 清空回收站
//
// @param request - ClearRecycleBinRequest
//
// @param headers - ClearRecycleBinHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearRecycleBinResponse
func (client *Client) ClearRecycleBinWithOptions(recycleBinId *string, request *ClearRecycleBinRequest, headers *ClearRecycleBinHeaders, runtime *util.RuntimeOptions) (_result *ClearRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearRecycleBin"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/recycleBins/" + tea.StringValue(recycleBinId) + "/clear"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearRecycleBinResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清空回收站
//
// @param request - ClearRecycleBinRequest
//
// @return ClearRecycleBinResponse
func (client *Client) ClearRecycleBin(recycleBinId *string, request *ClearRecycleBinRequest) (_result *ClearRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ClearRecycleBinHeaders{}
	_result = &ClearRecycleBinResponse{}
	_body, _err := client.ClearRecycleBinWithOptions(recycleBinId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交文件
//
// @param request - CommitFileRequest
//
// @param headers - CommitFileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommitFileResponse
func (client *Client) CommitFileWithOptions(spaceId *string, request *CommitFileRequest, headers *CommitFileHeaders, runtime *util.RuntimeOptions) (_result *CommitFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.OverwriteDentryId)) {
		body["overwriteDentryId"] = request.OverwriteDentryId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["parentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UploadKey)) {
		body["uploadKey"] = request.UploadKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CommitFile"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/files/commit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CommitFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交文件
//
// @param request - CommitFileRequest
//
// @return CommitFileResponse
func (client *Client) CommitFile(spaceId *string, request *CommitFileRequest) (_result *CommitFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CommitFileHeaders{}
	_result = &CommitFileResponse{}
	_body, _err := client.CommitFileWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量拷贝文件或文件夹
//
// @param request - CopyDentriesRequest
//
// @param headers - CopyDentriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyDentriesResponse
func (client *Client) CopyDentriesWithOptions(spaceId *string, request *CopyDentriesRequest, headers *CopyDentriesHeaders, runtime *util.RuntimeOptions) (_result *CopyDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryIds)) {
		body["dentryIds"] = request.DentryIds
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFolderId)) {
		body["targetFolderId"] = request.TargetFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSpaceId)) {
		body["targetSpaceId"] = request.TargetSpaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyDentries"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/batchCopy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量拷贝文件或文件夹
//
// @param request - CopyDentriesRequest
//
// @return CopyDentriesResponse
func (client *Client) CopyDentries(spaceId *string, request *CopyDentriesRequest) (_result *CopyDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CopyDentriesHeaders{}
	_result = &CopyDentriesResponse{}
	_body, _err := client.CopyDentriesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拷贝文件或文件夹
//
// @param request - CopyDentryRequest
//
// @param headers - CopyDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyDentryResponse
func (client *Client) CopyDentryWithOptions(spaceId *string, dentryId *string, request *CopyDentryRequest, headers *CopyDentryHeaders, runtime *util.RuntimeOptions) (_result *CopyDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFolderId)) {
		body["targetFolderId"] = request.TargetFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSpaceId)) {
		body["targetSpaceId"] = request.TargetSpaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyDentry"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/copy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拷贝文件或文件夹
//
// @param request - CopyDentryRequest
//
// @return CopyDentryResponse
func (client *Client) CopyDentry(spaceId *string, dentryId *string, request *CopyDentryRequest) (_result *CopyDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CopyDentryHeaders{}
	_result = &CopyDentryResponse{}
	_body, _err := client.CopyDentryWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除文件或文件夹
//
// @param request - DeleteDentriesRequest
//
// @param headers - DeleteDentriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDentriesResponse
func (client *Client) DeleteDentriesWithOptions(spaceId *string, request *DeleteDentriesRequest, headers *DeleteDentriesHeaders, runtime *util.RuntimeOptions) (_result *DeleteDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryIds)) {
		body["dentryIds"] = request.DentryIds
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDentries"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/batchRemove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除文件或文件夹
//
// @param request - DeleteDentriesRequest
//
// @return DeleteDentriesResponse
func (client *Client) DeleteDentries(spaceId *string, request *DeleteDentriesRequest) (_result *DeleteDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDentriesHeaders{}
	_result = &DeleteDentriesResponse{}
	_body, _err := client.DeleteDentriesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文件或文件夹
//
// @param request - DeleteDentryRequest
//
// @param headers - DeleteDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDentryResponse
func (client *Client) DeleteDentryWithOptions(spaceId *string, dentryId *string, request *DeleteDentryRequest, headers *DeleteDentryHeaders, runtime *util.RuntimeOptions) (_result *DeleteDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ToRecycleBin)) {
		query["toRecycleBin"] = request.ToRecycleBin
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDentry"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文件或文件夹
//
// @param request - DeleteDentryRequest
//
// @return DeleteDentryResponse
func (client *Client) DeleteDentry(spaceId *string, dentryId *string, request *DeleteDentryRequest) (_result *DeleteDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDentryHeaders{}
	_result = &DeleteDentryResponse{}
	_body, _err := client.DeleteDentryWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文件上的App属性值
//
// @param request - DeleteDentryAppPropertiesRequest
//
// @param headers - DeleteDentryAppPropertiesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDentryAppPropertiesResponse
func (client *Client) DeleteDentryAppPropertiesWithOptions(spaceId *string, dentryId *string, request *DeleteDentryAppPropertiesRequest, headers *DeleteDentryAppPropertiesHeaders, runtime *util.RuntimeOptions) (_result *DeleteDentryAppPropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyNames)) {
		body["propertyNames"] = request.PropertyNames
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDentryAppProperties"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/appProperties/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDentryAppPropertiesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文件上的App属性值
//
// @param request - DeleteDentryAppPropertiesRequest
//
// @return DeleteDentryAppPropertiesResponse
func (client *Client) DeleteDentryAppProperties(spaceId *string, dentryId *string, request *DeleteDentryAppPropertiesRequest) (_result *DeleteDentryAppPropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDentryAppPropertiesHeaders{}
	_result = &DeleteDentryAppPropertiesResponse{}
	_body, _err := client.DeleteDentryAppPropertiesWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除权限
//
// @param request - DeletePermissionRequest
//
// @param headers - DeletePermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePermissionResponse
func (client *Client) DeletePermissionWithOptions(spaceId *string, dentryId *string, request *DeletePermissionRequest, headers *DeletePermissionHeaders, runtime *util.RuntimeOptions) (_result *DeletePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		body["roleId"] = request.RoleId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePermission"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/permissions/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除权限
//
// @param request - DeletePermissionRequest
//
// @return DeletePermissionResponse
func (client *Client) DeletePermission(spaceId *string, dentryId *string, request *DeletePermissionRequest) (_result *DeletePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeletePermissionHeaders{}
	_result = &DeletePermissionResponse{}
	_body, _err := client.DeletePermissionWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除回收项, 删除之后该记录从回收站删除, 后续文件就无法恢复了
//
// @param request - DeleteRecycleItemRequest
//
// @param headers - DeleteRecycleItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecycleItemResponse
func (client *Client) DeleteRecycleItemWithOptions(recycleBinId *string, recycleItemId *string, request *DeleteRecycleItemRequest, headers *DeleteRecycleItemHeaders, runtime *util.RuntimeOptions) (_result *DeleteRecycleItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRecycleItem"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/recycleBins/" + tea.StringValue(recycleBinId) + "/recycleItems/" + tea.StringValue(recycleItemId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRecycleItemResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除回收项, 删除之后该记录从回收站删除, 后续文件就无法恢复了
//
// @param request - DeleteRecycleItemRequest
//
// @return DeleteRecycleItemResponse
func (client *Client) DeleteRecycleItem(recycleBinId *string, recycleItemId *string, request *DeleteRecycleItemRequest) (_result *DeleteRecycleItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRecycleItemHeaders{}
	_result = &DeleteRecycleItemResponse{}
	_body, _err := client.DeleteRecycleItemWithOptions(recycleBinId, recycleItemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除回收项, 删除之后该记录从回收站删除, 后续文件就无法恢复了
//
// @param request - DeleteRecycleItemsRequest
//
// @param headers - DeleteRecycleItemsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecycleItemsResponse
func (client *Client) DeleteRecycleItemsWithOptions(recycleBinId *string, request *DeleteRecycleItemsRequest, headers *DeleteRecycleItemsHeaders, runtime *util.RuntimeOptions) (_result *DeleteRecycleItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecycleItemIds)) {
		body["recycleItemIds"] = request.RecycleItemIds
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRecycleItems"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/recycleBins/" + tea.StringValue(recycleBinId) + "/recycleItems/batchRemove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRecycleItemsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除回收项, 删除之后该记录从回收站删除, 后续文件就无法恢复了
//
// @param request - DeleteRecycleItemsRequest
//
// @return DeleteRecycleItemsResponse
func (client *Client) DeleteRecycleItems(recycleBinId *string, request *DeleteRecycleItemsRequest) (_result *DeleteRecycleItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRecycleItemsHeaders{}
	_result = &DeleteRecycleItemsResponse{}
	_body, _err := client.DeleteRecycleItemsWithOptions(recycleBinId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取开放平台应用在企业存储中的相关应用信息
//
// @param request - GetCurrentAppRequest
//
// @param headers - GetCurrentAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCurrentAppResponse
func (client *Client) GetCurrentAppWithOptions(request *GetCurrentAppRequest, headers *GetCurrentAppHeaders, runtime *util.RuntimeOptions) (_result *GetCurrentAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCurrentApp"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/currentApps/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCurrentAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取开放平台应用在企业存储中的相关应用信息
//
// @param request - GetCurrentAppRequest
//
// @return GetCurrentAppResponse
func (client *Client) GetCurrentApp(request *GetCurrentAppRequest) (_result *GetCurrentAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCurrentAppHeaders{}
	_result = &GetCurrentAppResponse{}
	_body, _err := client.GetCurrentAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取文件(夹)信息
//
// @param request - GetDentriesRequest
//
// @param headers - GetDentriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDentriesResponse
func (client *Client) GetDentriesWithOptions(spaceId *string, request *GetDentriesRequest, headers *GetDentriesHeaders, runtime *util.RuntimeOptions) (_result *GetDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryIds)) {
		body["dentryIds"] = request.DentryIds
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDentries"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取文件(夹)信息
//
// @param request - GetDentriesRequest
//
// @return GetDentriesResponse
func (client *Client) GetDentries(spaceId *string, request *GetDentriesRequest) (_result *GetDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDentriesHeaders{}
	_result = &GetDentriesResponse{}
	_body, _err := client.GetDentriesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件(夹)信息
//
// @param request - GetDentryRequest
//
// @param headers - GetDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDentryResponse
func (client *Client) GetDentryWithOptions(spaceId *string, dentryId *string, request *GetDentryRequest, headers *GetDentryHeaders, runtime *util.RuntimeOptions) (_result *GetDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDentry"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件(夹)信息
//
// @param request - GetDentryRequest
//
// @return GetDentryResponse
func (client *Client) GetDentry(spaceId *string, dentryId *string, request *GetDentryRequest) (_result *GetDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDentryHeaders{}
	_result = &GetDentryResponse{}
	_body, _err := client.GetDentryWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件打开链接
//
// @param request - GetDentryOpenInfoRequest
//
// @param headers - GetDentryOpenInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDentryOpenInfoResponse
func (client *Client) GetDentryOpenInfoWithOptions(spaceId *string, dentryId *string, request *GetDentryOpenInfoRequest, headers *GetDentryOpenInfoHeaders, runtime *util.RuntimeOptions) (_result *GetDentryOpenInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDentryOpenInfo"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/openInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDentryOpenInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件打开链接
//
// @param request - GetDentryOpenInfoRequest
//
// @return GetDentryOpenInfoResponse
func (client *Client) GetDentryOpenInfo(spaceId *string, dentryId *string, request *GetDentryOpenInfoRequest) (_result *GetDentryOpenInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDentryOpenInfoHeaders{}
	_result = &GetDentryOpenInfoResponse{}
	_body, _err := client.GetDentryOpenInfoWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取文件缩略图
//
// @param request - GetDentryThumbnailsRequest
//
// @param headers - GetDentryThumbnailsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDentryThumbnailsResponse
func (client *Client) GetDentryThumbnailsWithOptions(spaceId *string, request *GetDentryThumbnailsRequest, headers *GetDentryThumbnailsHeaders, runtime *util.RuntimeOptions) (_result *GetDentryThumbnailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryIds)) {
		body["dentryIds"] = request.DentryIds
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDentryThumbnails"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/thumbnails/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDentryThumbnailsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取文件缩略图
//
// @param request - GetDentryThumbnailsRequest
//
// @return GetDentryThumbnailsResponse
func (client *Client) GetDentryThumbnails(spaceId *string, request *GetDentryThumbnailsRequest) (_result *GetDentryThumbnailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDentryThumbnailsHeaders{}
	_result = &GetDentryThumbnailsResponse{}
	_body, _err := client.GetDentryThumbnailsWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件下载信息
//
// @param request - GetFileDownloadInfoRequest
//
// @param headers - GetFileDownloadInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileDownloadInfoResponse
func (client *Client) GetFileDownloadInfoWithOptions(spaceId *string, dentryId *string, request *GetFileDownloadInfoRequest, headers *GetFileDownloadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileDownloadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileDownloadInfo"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/downloadInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileDownloadInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件下载信息
//
// @param request - GetFileDownloadInfoRequest
//
// @return GetFileDownloadInfoResponse
func (client *Client) GetFileDownloadInfo(spaceId *string, dentryId *string, request *GetFileDownloadInfoRequest) (_result *GetFileDownloadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileDownloadInfoHeaders{}
	_result = &GetFileDownloadInfoResponse{}
	_body, _err := client.GetFileDownloadInfoWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件上传信息
//
// @param request - GetFileUploadInfoRequest
//
// @param headers - GetFileUploadInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileUploadInfoResponse
func (client *Client) GetFileUploadInfoWithOptions(spaceId *string, request *GetFileUploadInfoRequest, headers *GetFileUploadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Multipart)) {
		body["multipart"] = request.Multipart
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["protocol"] = request.Protocol
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileUploadInfo"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/files/uploadInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件上传信息
//
// @param request - GetFileUploadInfoRequest
//
// @return GetFileUploadInfoResponse
func (client *Client) GetFileUploadInfo(spaceId *string, request *GetFileUploadInfoRequest) (_result *GetFileUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileUploadInfoHeaders{}
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.GetFileUploadInfoWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件上传信息(分片上传)
//
// @param request - GetMultipartFileUploadInfosRequest
//
// @param headers - GetMultipartFileUploadInfosHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultipartFileUploadInfosResponse
func (client *Client) GetMultipartFileUploadInfosWithOptions(request *GetMultipartFileUploadInfosRequest, headers *GetMultipartFileUploadInfosHeaders, runtime *util.RuntimeOptions) (_result *GetMultipartFileUploadInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.PartNumbers)) {
		body["partNumbers"] = request.PartNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.UploadKey)) {
		body["uploadKey"] = request.UploadKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultipartFileUploadInfos"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/files/multiPartUploadInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultipartFileUploadInfosResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件上传信息(分片上传)
//
// @param request - GetMultipartFileUploadInfosRequest
//
// @return GetMultipartFileUploadInfosResponse
func (client *Client) GetMultipartFileUploadInfos(request *GetMultipartFileUploadInfosRequest) (_result *GetMultipartFileUploadInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMultipartFileUploadInfosHeaders{}
	_result = &GetMultipartFileUploadInfosResponse{}
	_body, _err := client.GetMultipartFileUploadInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业存储中企业维度的信息
//
// @param request - GetOrgRequest
//
// @param headers - GetOrgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrgResponse
func (client *Client) GetOrgWithOptions(corpId *string, request *GetOrgRequest, headers *GetOrgHeaders, runtime *util.RuntimeOptions) (_result *GetOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrg"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/orgs/" + tea.StringValue(corpId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业存储中企业维度的信息
//
// @param request - GetOrgRequest
//
// @return GetOrgResponse
func (client *Client) GetOrg(corpId *string, request *GetOrgRequest) (_result *GetOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrgHeaders{}
	_result = &GetOrgResponse{}
	_body, _err := client.GetOrgWithOptions(corpId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取回收站信息
//
// @param request - GetRecycleBinRequest
//
// @param headers - GetRecycleBinHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecycleBinResponse
func (client *Client) GetRecycleBinWithOptions(request *GetRecycleBinRequest, headers *GetRecycleBinHeaders, runtime *util.RuntimeOptions) (_result *GetRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecycleBinScope)) {
		query["recycleBinScope"] = request.RecycleBinScope
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeId)) {
		query["scopeId"] = request.ScopeId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecycleBin"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/recycleBins"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecycleBinResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取回收站信息
//
// @param request - GetRecycleBinRequest
//
// @return GetRecycleBinResponse
func (client *Client) GetRecycleBin(request *GetRecycleBinRequest) (_result *GetRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRecycleBinHeaders{}
	_result = &GetRecycleBinResponse{}
	_body, _err := client.GetRecycleBinWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取回收项详情
//
// @param request - GetRecycleItemRequest
//
// @param headers - GetRecycleItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecycleItemResponse
func (client *Client) GetRecycleItemWithOptions(recycleBinId *string, recycleItemId *string, request *GetRecycleItemRequest, headers *GetRecycleItemHeaders, runtime *util.RuntimeOptions) (_result *GetRecycleItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecycleItem"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/recycleBins/" + tea.StringValue(recycleBinId) + "/recycleItems/" + tea.StringValue(recycleItemId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecycleItemResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取回收项详情
//
// @param request - GetRecycleItemRequest
//
// @return GetRecycleItemResponse
func (client *Client) GetRecycleItem(recycleBinId *string, recycleItemId *string, request *GetRecycleItemRequest) (_result *GetRecycleItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRecycleItemHeaders{}
	_result = &GetRecycleItemResponse{}
	_body, _err := client.GetRecycleItemWithOptions(recycleBinId, recycleItemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取空间信息
//
// @param request - GetSpaceRequest
//
// @param headers - GetSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpaceResponse
func (client *Client) GetSpaceWithOptions(spaceId *string, request *GetSpaceRequest, headers *GetSpaceHeaders, runtime *util.RuntimeOptions) (_result *GetSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSpace"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取空间信息
//
// @param request - GetSpaceRequest
//
// @return GetSpaceResponse
func (client *Client) GetSpace(spaceId *string, request *GetSpaceRequest) (_result *GetSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSpaceHeaders{}
	_result = &GetSpaceResponse{}
	_body, _err := client.GetSpaceWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步任务信息
//
// @param request - GetTaskRequest
//
// @param headers - GetTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(taskId *string, request *GetTaskRequest, headers *GetTaskHeaders, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步任务信息
//
// @param request - GetTaskRequest
//
// @return GetTaskResponse
func (client *Client) GetTask(taskId *string, request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTaskHeaders{}
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 WebOfficeUrl 接口
//
// @param request - GetWebOfficeUrlRequest
//
// @param headers - GetWebOfficeUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWebOfficeUrlResponse
func (client *Client) GetWebOfficeUrlWithOptions(spaceId *string, dentryId *string, request *GetWebOfficeUrlRequest, headers *GetWebOfficeUrlHeaders, runtime *util.RuntimeOptions) (_result *GetWebOfficeUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebOfficeUrl"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/webOfficeUrls"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebOfficeUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 WebOfficeUrl 接口
//
// @param request - GetWebOfficeUrlRequest
//
// @return GetWebOfficeUrlResponse
func (client *Client) GetWebOfficeUrl(spaceId *string, dentryId *string, request *GetWebOfficeUrlRequest) (_result *GetWebOfficeUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWebOfficeUrlHeaders{}
	_result = &GetWebOfficeUrlResponse{}
	_body, _err := client.GetWebOfficeUrlWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 初始化文件分片上传
//
// @param request - InitMultipartFileUploadRequest
//
// @param headers - InitMultipartFileUploadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitMultipartFileUploadResponse
func (client *Client) InitMultipartFileUploadWithOptions(spaceId *string, request *InitMultipartFileUploadRequest, headers *InitMultipartFileUploadHeaders, runtime *util.RuntimeOptions) (_result *InitMultipartFileUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitMultipartFileUpload"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/files/multiPartUploadInfos/init"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InitMultipartFileUploadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 初始化文件分片上传
//
// @param request - InitMultipartFileUploadRequest
//
// @return InitMultipartFileUploadResponse
func (client *Client) InitMultipartFileUpload(spaceId *string, request *InitMultipartFileUploadRequest) (_result *InitMultipartFileUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InitMultipartFileUploadHeaders{}
	_result = &InitMultipartFileUploadResponse{}
	_body, _err := client.InitMultipartFileUploadWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件列表
//
// @param request - ListAllDentriesRequest
//
// @param headers - ListAllDentriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllDentriesResponse
func (client *Client) ListAllDentriesWithOptions(spaceId *string, request *ListAllDentriesRequest, headers *ListAllDentriesHeaders, runtime *util.RuntimeOptions) (_result *ListAllDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAllDentries"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/listAll"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAllDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件列表
//
// @param request - ListAllDentriesRequest
//
// @return ListAllDentriesResponse
func (client *Client) ListAllDentries(spaceId *string, request *ListAllDentriesRequest) (_result *ListAllDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAllDentriesHeaders{}
	_result = &ListAllDentriesResponse{}
	_body, _err := client.ListAllDentriesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件列表
//
// @param request - ListDentriesRequest
//
// @param headers - ListDentriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDentriesResponse
func (client *Client) ListDentriesWithOptions(spaceId *string, request *ListDentriesRequest, headers *ListDentriesHeaders, runtime *util.RuntimeOptions) (_result *ListDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["parentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.WithThumbnail)) {
		query["withThumbnail"] = request.WithThumbnail
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDentries"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件列表
//
// @param request - ListDentriesRequest
//
// @return ListDentriesResponse
func (client *Client) ListDentries(spaceId *string, request *ListDentriesRequest) (_result *ListDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDentriesHeaders{}
	_result = &ListDentriesResponse{}
	_body, _err := client.ListDentriesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件历史版本
//
// @param request - ListDentryVersionsRequest
//
// @param headers - ListDentryVersionsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDentryVersionsResponse
func (client *Client) ListDentryVersionsWithOptions(spaceId *string, dentryId *string, request *ListDentryVersionsRequest, headers *ListDentryVersionsHeaders, runtime *util.RuntimeOptions) (_result *ListDentryVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDentryVersions"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDentryVersionsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件历史版本
//
// @param request - ListDentryVersionsRequest
//
// @return ListDentryVersionsResponse
func (client *Client) ListDentryVersions(spaceId *string, dentryId *string, request *ListDentryVersionsRequest) (_result *ListDentryVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDentryVersionsHeaders{}
	_result = &ListDentryVersionsResponse{}
	_body, _err := client.ListDentryVersionsWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取权限列表
//
// @param request - ListPermissionsRequest
//
// @param headers - ListPermissionsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionsResponse
func (client *Client) ListPermissionsWithOptions(spaceId *string, dentryId *string, request *ListPermissionsRequest, headers *ListPermissionsHeaders, runtime *util.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPermissions"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/permissions/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPermissionsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取权限列表
//
// @param request - ListPermissionsRequest
//
// @return ListPermissionsResponse
func (client *Client) ListPermissions(spaceId *string, dentryId *string, request *ListPermissionsRequest) (_result *ListPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListPermissionsHeaders{}
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取回收项列表
//
// @param request - ListRecycleItemsRequest
//
// @param headers - ListRecycleItemsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecycleItemsResponse
func (client *Client) ListRecycleItemsWithOptions(recycleBinId *string, request *ListRecycleItemsRequest, headers *ListRecycleItemsHeaders, runtime *util.RuntimeOptions) (_result *ListRecycleItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecycleItems"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/recycleBins/" + tea.StringValue(recycleBinId) + "/recycleItems"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecycleItemsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取回收项列表
//
// @param request - ListRecycleItemsRequest
//
// @return ListRecycleItemsResponse
func (client *Client) ListRecycleItems(recycleBinId *string, request *ListRecycleItemsRequest) (_result *ListRecycleItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRecycleItemsHeaders{}
	_result = &ListRecycleItemsResponse{}
	_body, _err := client.ListRecycleItemsWithOptions(recycleBinId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量移动文件或文件夹
//
// @param request - MoveDentriesRequest
//
// @param headers - MoveDentriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveDentriesResponse
func (client *Client) MoveDentriesWithOptions(spaceId *string, request *MoveDentriesRequest, headers *MoveDentriesHeaders, runtime *util.RuntimeOptions) (_result *MoveDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryIds)) {
		body["dentryIds"] = request.DentryIds
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFolderId)) {
		body["targetFolderId"] = request.TargetFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSpaceId)) {
		body["targetSpaceId"] = request.TargetSpaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveDentries"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/batchMove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量移动文件或文件夹
//
// @param request - MoveDentriesRequest
//
// @return MoveDentriesResponse
func (client *Client) MoveDentries(spaceId *string, request *MoveDentriesRequest) (_result *MoveDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MoveDentriesHeaders{}
	_result = &MoveDentriesResponse{}
	_body, _err := client.MoveDentriesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动文件或文件夹
//
// @param request - MoveDentryRequest
//
// @param headers - MoveDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveDentryResponse
func (client *Client) MoveDentryWithOptions(spaceId *string, dentryId *string, request *MoveDentryRequest, headers *MoveDentryHeaders, runtime *util.RuntimeOptions) (_result *MoveDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFolderId)) {
		body["targetFolderId"] = request.TargetFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSpaceId)) {
		body["targetSpaceId"] = request.TargetSpaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveDentry"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/move"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动文件或文件夹
//
// @param request - MoveDentryRequest
//
// @return MoveDentryResponse
func (client *Client) MoveDentry(spaceId *string, dentryId *string, request *MoveDentryRequest) (_result *MoveDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MoveDentryHeaders{}
	_result = &MoveDentryResponse{}
	_body, _err := client.MoveDentryWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 accessToken 接口
//
// @param request - RefreshWebOfficeTokenRequest
//
// @param headers - RefreshWebOfficeTokenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshWebOfficeTokenResponse
func (client *Client) RefreshWebOfficeTokenWithOptions(spaceId *string, dentryId *string, request *RefreshWebOfficeTokenRequest, headers *RefreshWebOfficeTokenHeaders, runtime *util.RuntimeOptions) (_result *RefreshWebOfficeTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.WebOfficeAccessToken)) {
		query["webOfficeAccessToken"] = request.WebOfficeAccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.WebOfficeRefreshToken)) {
		query["webOfficeRefreshToken"] = request.WebOfficeRefreshToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshWebOfficeToken"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/refreshWebOfficeToken"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshWebOfficeTokenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 accessToken 接口
//
// @param request - RefreshWebOfficeTokenRequest
//
// @return RefreshWebOfficeTokenResponse
func (client *Client) RefreshWebOfficeToken(spaceId *string, dentryId *string, request *RefreshWebOfficeTokenRequest) (_result *RefreshWebOfficeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RefreshWebOfficeTokenHeaders{}
	_result = &RefreshWebOfficeTokenResponse{}
	_body, _err := client.RefreshWebOfficeTokenWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 注册文件预览或编辑链接
//
// @param request - RegisterOpenInfoRequest
//
// @param headers - RegisterOpenInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterOpenInfoResponse
func (client *Client) RegisterOpenInfoWithOptions(spaceId *string, dentryId *string, request *RegisterOpenInfoRequest, headers *RegisterOpenInfoHeaders, runtime *util.RuntimeOptions) (_result *RegisterOpenInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenInfos)) {
		body["openInfos"] = request.OpenInfos
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		body["provider"] = request.Provider
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterOpenInfo"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/openInfos/register"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterOpenInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册文件预览或编辑链接
//
// @param request - RegisterOpenInfoRequest
//
// @return RegisterOpenInfoResponse
func (client *Client) RegisterOpenInfo(spaceId *string, dentryId *string, request *RegisterOpenInfoRequest) (_result *RegisterOpenInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterOpenInfoHeaders{}
	_result = &RegisterOpenInfoResponse{}
	_body, _err := client.RegisterOpenInfoWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重命名文件或文件夹
//
// @param request - RenameDentryRequest
//
// @param headers - RenameDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameDentryResponse
func (client *Client) RenameDentryWithOptions(spaceId *string, dentryId *string, request *RenameDentryRequest, headers *RenameDentryHeaders, runtime *util.RuntimeOptions) (_result *RenameDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewName)) {
		body["newName"] = request.NewName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameDentry"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/rename"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameDentryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重命名文件或文件夹
//
// @param request - RenameDentryRequest
//
// @return RenameDentryResponse
func (client *Client) RenameDentry(spaceId *string, dentryId *string, request *RenameDentryRequest) (_result *RenameDentryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RenameDentryHeaders{}
	_result = &RenameDentryResponse{}
	_body, _err := client.RenameDentryWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 还原回收站中的回收项
//
// @param request - RestoreRecycleItemRequest
//
// @param headers - RestoreRecycleItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreRecycleItemResponse
func (client *Client) RestoreRecycleItemWithOptions(recycleBinId *string, recycleItemId *string, request *RestoreRecycleItemRequest, headers *RestoreRecycleItemHeaders, runtime *util.RuntimeOptions) (_result *RestoreRecycleItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreRecycleItem"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/recycleBins/" + tea.StringValue(recycleBinId) + "/recycleItems/" + tea.StringValue(recycleItemId) + "/restore"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreRecycleItemResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 还原回收站中的回收项
//
// @param request - RestoreRecycleItemRequest
//
// @return RestoreRecycleItemResponse
func (client *Client) RestoreRecycleItem(recycleBinId *string, recycleItemId *string, request *RestoreRecycleItemRequest) (_result *RestoreRecycleItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RestoreRecycleItemHeaders{}
	_result = &RestoreRecycleItemResponse{}
	_body, _err := client.RestoreRecycleItemWithOptions(recycleBinId, recycleItemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量还原回收站中的回收项
//
// @param request - RestoreRecycleItemsRequest
//
// @param headers - RestoreRecycleItemsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreRecycleItemsResponse
func (client *Client) RestoreRecycleItemsWithOptions(recycleBinId *string, request *RestoreRecycleItemsRequest, headers *RestoreRecycleItemsHeaders, runtime *util.RuntimeOptions) (_result *RestoreRecycleItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.RecycleItemIds)) {
		body["recycleItemIds"] = request.RecycleItemIds
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreRecycleItems"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/recycleBins/" + tea.StringValue(recycleBinId) + "/recycleItems/batchRestore"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreRecycleItemsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量还原回收站中的回收项
//
// @param request - RestoreRecycleItemsRequest
//
// @return RestoreRecycleItemsResponse
func (client *Client) RestoreRecycleItems(recycleBinId *string, request *RestoreRecycleItemsRequest) (_result *RestoreRecycleItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RestoreRecycleItemsHeaders{}
	_result = &RestoreRecycleItemsResponse{}
	_body, _err := client.RestoreRecycleItemsWithOptions(recycleBinId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 恢复文件历史版本
//
// @param request - RevertDentryVersionRequest
//
// @param headers - RevertDentryVersionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertDentryVersionResponse
func (client *Client) RevertDentryVersionWithOptions(spaceId *string, dentryId *string, version *string, request *RevertDentryVersionRequest, headers *RevertDentryVersionHeaders, runtime *util.RuntimeOptions) (_result *RevertDentryVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevertDentryVersion"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/versions/" + tea.StringValue(version) + "/revert"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RevertDentryVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恢复文件历史版本
//
// @param request - RevertDentryVersionRequest
//
// @return RevertDentryVersionResponse
func (client *Client) RevertDentryVersion(spaceId *string, dentryId *string, version *string, request *RevertDentryVersionRequest) (_result *RevertDentryVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RevertDentryVersionHeaders{}
	_result = &RevertDentryVersionResponse{}
	_body, _err := client.RevertDentryVersionWithOptions(spaceId, dentryId, version, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 订阅文件变更事件
//
// @param request - SubscribeEventRequest
//
// @param headers - SubscribeEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribeEventResponse
func (client *Client) SubscribeEventWithOptions(request *SubscribeEventRequest, headers *SubscribeEventHeaders, runtime *util.RuntimeOptions) (_result *SubscribeEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeId)) {
		body["scopeId"] = request.ScopeId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubscribeEvent"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/events/subscribe"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SubscribeEventResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订阅文件变更事件
//
// @param request - SubscribeEventRequest
//
// @return SubscribeEventResponse
func (client *Client) SubscribeEvent(request *SubscribeEventRequest) (_result *SubscribeEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubscribeEventHeaders{}
	_result = &SubscribeEventResponse{}
	_body, _err := client.SubscribeEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消订阅文件变更事件
//
// @param request - UnsubscribeEventRequest
//
// @param headers - UnsubscribeEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnsubscribeEventResponse
func (client *Client) UnsubscribeEventWithOptions(request *UnsubscribeEventRequest, headers *UnsubscribeEventHeaders, runtime *util.RuntimeOptions) (_result *UnsubscribeEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeId)) {
		body["scopeId"] = request.ScopeId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnsubscribeEvent"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/events/unsubscribe"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UnsubscribeEventResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消订阅文件变更事件
//
// @param request - UnsubscribeEventRequest
//
// @return UnsubscribeEventResponse
func (client *Client) UnsubscribeEvent(request *UnsubscribeEventRequest) (_result *UnsubscribeEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnsubscribeEventHeaders{}
	_result = &UnsubscribeEventResponse{}
	_body, _err := client.UnsubscribeEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改文件上的App属性值
//
// @param request - UpdateDentryAppPropertiesRequest
//
// @param headers - UpdateDentryAppPropertiesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDentryAppPropertiesResponse
func (client *Client) UpdateDentryAppPropertiesWithOptions(spaceId *string, dentryId *string, request *UpdateDentryAppPropertiesRequest, headers *UpdateDentryAppPropertiesHeaders, runtime *util.RuntimeOptions) (_result *UpdateDentryAppPropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppProperties)) {
		body["appProperties"] = request.AppProperties
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDentryAppProperties"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/appProperties"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDentryAppPropertiesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改文件上的App属性值
//
// @param request - UpdateDentryAppPropertiesRequest
//
// @return UpdateDentryAppPropertiesResponse
func (client *Client) UpdateDentryAppProperties(spaceId *string, dentryId *string, request *UpdateDentryAppPropertiesRequest) (_result *UpdateDentryAppPropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateDentryAppPropertiesHeaders{}
	_result = &UpdateDentryAppPropertiesResponse{}
	_body, _err := client.UpdateDentryAppPropertiesWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改权限
//
// @param request - UpdatePermissionRequest
//
// @param headers - UpdatePermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePermissionResponse
func (client *Client) UpdatePermissionWithOptions(spaceId *string, dentryId *string, request *UpdatePermissionRequest, headers *UpdatePermissionHeaders, runtime *util.RuntimeOptions) (_result *UpdatePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		body["roleId"] = request.RoleId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePermission"),
		Version:     tea.String("storage_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/storage/spaces/" + tea.StringValue(spaceId) + "/dentries/" + tea.StringValue(dentryId) + "/permissions"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改权限
//
// @param request - UpdatePermissionRequest
//
// @return UpdatePermissionResponse
func (client *Client) UpdatePermission(spaceId *string, dentryId *string, request *UpdatePermissionRequest) (_result *UpdatePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePermissionHeaders{}
	_result = &UpdatePermissionResponse{}
	_body, _err := client.UpdatePermissionWithOptions(spaceId, dentryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
