// This file is auto-generated, don't edit it. Thanks.
package storage_2_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	DentryUuidList []*string `json:"dentryUuidList,omitempty" xml:"dentryUuidList,omitempty" type:"Repeated"`
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

func (s *BatchQueryRolesRequest) SetDentryUuidList(v []*string) *BatchQueryRolesRequest {
	s.DentryUuidList = v
	return s
}

func (s *BatchQueryRolesRequest) SetUnionId(v string) *BatchQueryRolesRequest {
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
	// false
	ConvertToOnlineDoc *bool `json:"convertToOnlineDoc,omitempty" xml:"convertToOnlineDoc,omitempty"`
	// example:
	//
	// DOC
	ConvertToOnlineDocTargetDocumentType *string `json:"convertToOnlineDocTargetDocumentType,omitempty" xml:"convertToOnlineDocTargetDocumentType,omitempty"`
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

func (s *CommitFileRequestOption) SetConvertToOnlineDoc(v bool) *CommitFileRequestOption {
	s.ConvertToOnlineDoc = &v
	return s
}

func (s *CommitFileRequestOption) SetConvertToOnlineDocTargetDocumentType(v string) *CommitFileRequestOption {
	s.ConvertToOnlineDocTargetDocumentType = &v
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
	Option *GetFileUploadInfoRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
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
	// dentry_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

func (s *GetFileUploadInfoRequestOptionPreCheckParam) SetName(v string) *GetFileUploadInfoRequestOptionPreCheckParam {
	s.Name = &v
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

type GetPermissionInheritanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPermissionInheritanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionInheritanceHeaders) GoString() string {
	return s.String()
}

func (s *GetPermissionInheritanceHeaders) SetCommonHeaders(v map[string]*string) *GetPermissionInheritanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPermissionInheritanceHeaders) SetXAcsDingtalkAccessToken(v string) *GetPermissionInheritanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPermissionInheritanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetPermissionInheritanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionInheritanceRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionInheritanceRequest) SetUnionId(v string) *GetPermissionInheritanceRequest {
	s.UnionId = &v
	return s
}

type GetPermissionInheritanceResponseBody struct {
	// example:
	//
	// PASS_ON
	Inheritance *string `json:"inheritance,omitempty" xml:"inheritance,omitempty"`
}

func (s GetPermissionInheritanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionInheritanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionInheritanceResponseBody) SetInheritance(v string) *GetPermissionInheritanceResponseBody {
	s.Inheritance = &v
	return s
}

type GetPermissionInheritanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPermissionInheritanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPermissionInheritanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionInheritanceResponse) GoString() string {
	return s.String()
}

func (s *GetPermissionInheritanceResponse) SetHeaders(v map[string]*string) *GetPermissionInheritanceResponse {
	s.Headers = v
	return s
}

func (s *GetPermissionInheritanceResponse) SetStatusCode(v int32) *GetPermissionInheritanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPermissionInheritanceResponse) SetBody(v *GetPermissionInheritanceResponseBody) *GetPermissionInheritanceResponse {
	s.Body = v
	return s
}

type GetPermissionShareScopeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPermissionShareScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionShareScopeHeaders) GoString() string {
	return s.String()
}

func (s *GetPermissionShareScopeHeaders) SetCommonHeaders(v map[string]*string) *GetPermissionShareScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPermissionShareScopeHeaders) SetXAcsDingtalkAccessToken(v string) *GetPermissionShareScopeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPermissionShareScopeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetPermissionShareScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionShareScopeRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionShareScopeRequest) SetUnionId(v string) *GetPermissionShareScopeRequest {
	s.UnionId = &v
	return s
}

type GetPermissionShareScopeResponseBody struct {
	// example:
	//
	// ORG_READ
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s GetPermissionShareScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionShareScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionShareScopeResponseBody) SetScope(v string) *GetPermissionShareScopeResponseBody {
	s.Scope = &v
	return s
}

type GetPermissionShareScopeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPermissionShareScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPermissionShareScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionShareScopeResponse) GoString() string {
	return s.String()
}

func (s *GetPermissionShareScopeResponse) SetHeaders(v map[string]*string) *GetPermissionShareScopeResponse {
	s.Headers = v
	return s
}

func (s *GetPermissionShareScopeResponse) SetStatusCode(v int32) *GetPermissionShareScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPermissionShareScopeResponse) SetBody(v *GetPermissionShareScopeResponseBody) *GetPermissionShareScopeResponse {
	s.Body = v
	return s
}

type ListOperationLogsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListOperationLogsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsHeaders) GoString() string {
	return s.String()
}

func (s *ListOperationLogsHeaders) SetCommonHeaders(v map[string]*string) *ListOperationLogsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOperationLogsHeaders) SetXAcsDingtalkAccessToken(v string) *ListOperationLogsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListOperationLogsRequest struct {
	EndTime   *int64                          `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Option    *ListOperationLogsRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	StartTime *int64                          `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListOperationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationLogsRequest) SetEndTime(v int64) *ListOperationLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListOperationLogsRequest) SetOption(v *ListOperationLogsRequestOption) *ListOperationLogsRequest {
	s.Option = v
	return s
}

func (s *ListOperationLogsRequest) SetStartTime(v int64) *ListOperationLogsRequest {
	s.StartTime = &v
	return s
}

type ListOperationLogsRequestOption struct {
	// example:
	//
	// 下载:download_file
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
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
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// 知识库:org_biz_meta
	Scenes []*string `json:"scenes,omitempty" xml:"scenes,omitempty" type:"Repeated"`
	// example:
	//
	// YndMj49yWj95B3qAfGz5pY9d83pmz5aA
	SubjectId *string `json:"subjectId,omitempty" xml:"subjectId,omitempty"`
}

func (s ListOperationLogsRequestOption) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsRequestOption) GoString() string {
	return s.String()
}

func (s *ListOperationLogsRequestOption) SetActions(v []*string) *ListOperationLogsRequestOption {
	s.Actions = v
	return s
}

func (s *ListOperationLogsRequestOption) SetMaxResults(v int32) *ListOperationLogsRequestOption {
	s.MaxResults = &v
	return s
}

func (s *ListOperationLogsRequestOption) SetNextToken(v string) *ListOperationLogsRequestOption {
	s.NextToken = &v
	return s
}

func (s *ListOperationLogsRequestOption) SetOperatorId(v string) *ListOperationLogsRequestOption {
	s.OperatorId = &v
	return s
}

func (s *ListOperationLogsRequestOption) SetScenes(v []*string) *ListOperationLogsRequestOption {
	s.Scenes = v
	return s
}

func (s *ListOperationLogsRequestOption) SetSubjectId(v string) *ListOperationLogsRequestOption {
	s.SubjectId = &v
	return s
}

type ListOperationLogsResponseBody struct {
	Items []*ListOperationLogsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListOperationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponseBody) SetItems(v []*ListOperationLogsResponseBodyItems) *ListOperationLogsResponseBody {
	s.Items = v
	return s
}

func (s *ListOperationLogsResponseBody) SetNextToken(v string) *ListOperationLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOperationLogsResponseBody) SetTotalCount(v int64) *ListOperationLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListOperationLogsResponseBodyItems struct {
	// example:
	//
	// add_permission
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// 「我的文档/无标题文档」，给用户“你好”添加了「可编辑」权限
	Details *string `json:"details,omitempty" xml:"details,omitempty"`
	// example:
	//
	// ovEQ1aYDoUrM8NldI7EPaDEuDfNce#AR4GpnMqJzRQ67PpuNNDQn9dJKe0xjE3&USER:1557011407
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	OperateTime *int64  `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// example:
	//
	// org_biz_meta_my_doc
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// AR4GpnMqJzRQ67PpuNNDQn9dJKe0xjE3
	SubjectId *string `json:"subjectId,omitempty" xml:"subjectId,omitempty"`
	// example:
	//
	// 无标题文档
	SubjectName *string `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	// example:
	//
	// DENTRY
	SubjectType *string `json:"subjectType,omitempty" xml:"subjectType,omitempty"`
	// example:
	//
	// https://alidocs.dingtalk.com/i/nodes/AR4GpnMqJzRQ67PpuNNDQn9dJKe0xjE3
	SubjectUrl *string `json:"subjectUrl,omitempty" xml:"subjectUrl,omitempty"`
}

func (s ListOperationLogsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponseBodyItems) SetAction(v string) *ListOperationLogsResponseBodyItems {
	s.Action = &v
	return s
}

func (s *ListOperationLogsResponseBodyItems) SetDetails(v string) *ListOperationLogsResponseBodyItems {
	s.Details = &v
	return s
}

func (s *ListOperationLogsResponseBodyItems) SetId(v string) *ListOperationLogsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListOperationLogsResponseBodyItems) SetOperateTime(v int64) *ListOperationLogsResponseBodyItems {
	s.OperateTime = &v
	return s
}

func (s *ListOperationLogsResponseBodyItems) SetOperatorId(v string) *ListOperationLogsResponseBodyItems {
	s.OperatorId = &v
	return s
}

func (s *ListOperationLogsResponseBodyItems) SetScene(v string) *ListOperationLogsResponseBodyItems {
	s.Scene = &v
	return s
}

func (s *ListOperationLogsResponseBodyItems) SetSubjectId(v string) *ListOperationLogsResponseBodyItems {
	s.SubjectId = &v
	return s
}

func (s *ListOperationLogsResponseBodyItems) SetSubjectName(v string) *ListOperationLogsResponseBodyItems {
	s.SubjectName = &v
	return s
}

func (s *ListOperationLogsResponseBodyItems) SetSubjectType(v string) *ListOperationLogsResponseBodyItems {
	s.SubjectType = &v
	return s
}

func (s *ListOperationLogsResponseBodyItems) SetSubjectUrl(v string) *ListOperationLogsResponseBodyItems {
	s.SubjectUrl = &v
	return s
}

type ListOperationLogsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponse) SetHeaders(v map[string]*string) *ListOperationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationLogsResponse) SetStatusCode(v int32) *ListOperationLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationLogsResponse) SetBody(v *ListOperationLogsResponseBody) *ListOperationLogsResponse {
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
	// space_id
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// example:
	//
	// 3600
	Duration *int64                                        `json:"duration,omitempty" xml:"duration,omitempty"`
	Member   *ListPermissionsResponseBodyPermissionsMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	Role     *ListPermissionsResponseBodyPermissionsRole   `json:"role,omitempty" xml:"role,omitempty" type:"Struct"`
}

func (s ListPermissionsResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissions) SetDentryUuid(v string) *ListPermissionsResponseBodyPermissions {
	s.DentryUuid = &v
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

func (s *ListPermissionsResponseBodyPermissions) SetRole(v *ListPermissionsResponseBodyPermissionsRole) *ListPermissionsResponseBodyPermissions {
	s.Role = v
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

type ManagerGetDefaultHandOverUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ManagerGetDefaultHandOverUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s ManagerGetDefaultHandOverUserHeaders) GoString() string {
	return s.String()
}

func (s *ManagerGetDefaultHandOverUserHeaders) SetCommonHeaders(v map[string]*string) *ManagerGetDefaultHandOverUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ManagerGetDefaultHandOverUserHeaders) SetXAcsDingtalkAccessToken(v string) *ManagerGetDefaultHandOverUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ManagerGetDefaultHandOverUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ManagerGetDefaultHandOverUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ManagerGetDefaultHandOverUserRequest) GoString() string {
	return s.String()
}

func (s *ManagerGetDefaultHandOverUserRequest) SetOperatorId(v string) *ManagerGetDefaultHandOverUserRequest {
	s.OperatorId = &v
	return s
}

type ManagerGetDefaultHandOverUserResponseBody struct {
	// example:
	//
	// staff_id
	DefaultHandoverUserId *string `json:"defaultHandoverUserId,omitempty" xml:"defaultHandoverUserId,omitempty"`
}

func (s ManagerGetDefaultHandOverUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManagerGetDefaultHandOverUserResponseBody) GoString() string {
	return s.String()
}

func (s *ManagerGetDefaultHandOverUserResponseBody) SetDefaultHandoverUserId(v string) *ManagerGetDefaultHandOverUserResponseBody {
	s.DefaultHandoverUserId = &v
	return s
}

type ManagerGetDefaultHandOverUserResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManagerGetDefaultHandOverUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManagerGetDefaultHandOverUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ManagerGetDefaultHandOverUserResponse) GoString() string {
	return s.String()
}

func (s *ManagerGetDefaultHandOverUserResponse) SetHeaders(v map[string]*string) *ManagerGetDefaultHandOverUserResponse {
	s.Headers = v
	return s
}

func (s *ManagerGetDefaultHandOverUserResponse) SetStatusCode(v int32) *ManagerGetDefaultHandOverUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ManagerGetDefaultHandOverUserResponse) SetBody(v *ManagerGetDefaultHandOverUserResponseBody) *ManagerGetDefaultHandOverUserResponse {
	s.Body = v
	return s
}

type ManagerSetDefaultHandOverUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ManagerSetDefaultHandOverUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s ManagerSetDefaultHandOverUserHeaders) GoString() string {
	return s.String()
}

func (s *ManagerSetDefaultHandOverUserHeaders) SetCommonHeaders(v map[string]*string) *ManagerSetDefaultHandOverUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ManagerSetDefaultHandOverUserHeaders) SetXAcsDingtalkAccessToken(v string) *ManagerSetDefaultHandOverUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ManagerSetDefaultHandOverUserRequest struct {
	// example:
	//
	// staff_id
	DefaultHandoverUserId *string `json:"defaultHandoverUserId,omitempty" xml:"defaultHandoverUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ManagerSetDefaultHandOverUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ManagerSetDefaultHandOverUserRequest) GoString() string {
	return s.String()
}

func (s *ManagerSetDefaultHandOverUserRequest) SetDefaultHandoverUserId(v string) *ManagerSetDefaultHandOverUserRequest {
	s.DefaultHandoverUserId = &v
	return s
}

func (s *ManagerSetDefaultHandOverUserRequest) SetOperatorId(v string) *ManagerSetDefaultHandOverUserRequest {
	s.OperatorId = &v
	return s
}

type ManagerSetDefaultHandOverUserResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ManagerSetDefaultHandOverUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManagerSetDefaultHandOverUserResponseBody) GoString() string {
	return s.String()
}

func (s *ManagerSetDefaultHandOverUserResponseBody) SetSuccess(v bool) *ManagerSetDefaultHandOverUserResponseBody {
	s.Success = &v
	return s
}

type ManagerSetDefaultHandOverUserResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManagerSetDefaultHandOverUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManagerSetDefaultHandOverUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ManagerSetDefaultHandOverUserResponse) GoString() string {
	return s.String()
}

func (s *ManagerSetDefaultHandOverUserResponse) SetHeaders(v map[string]*string) *ManagerSetDefaultHandOverUserResponse {
	s.Headers = v
	return s
}

func (s *ManagerSetDefaultHandOverUserResponse) SetStatusCode(v int32) *ManagerSetDefaultHandOverUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ManagerSetDefaultHandOverUserResponse) SetBody(v *ManagerSetDefaultHandOverUserResponseBody) *ManagerSetDefaultHandOverUserResponse {
	s.Body = v
	return s
}

type SearchDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesHeaders) GoString() string {
	return s.String()
}

func (s *SearchDentriesHeaders) SetCommonHeaders(v map[string]*string) *SearchDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *SearchDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchDentriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// keyword
	Keyword *string                      `json:"keyword,omitempty" xml:"keyword,omitempty"`
	Option  *SearchDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s SearchDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesRequest) GoString() string {
	return s.String()
}

func (s *SearchDentriesRequest) SetKeyword(v string) *SearchDentriesRequest {
	s.Keyword = &v
	return s
}

func (s *SearchDentriesRequest) SetOption(v *SearchDentriesRequestOption) *SearchDentriesRequest {
	s.Option = v
	return s
}

func (s *SearchDentriesRequest) SetOperatorId(v string) *SearchDentriesRequest {
	s.OperatorId = &v
	return s
}

type SearchDentriesRequestOption struct {
	CreateTimeRange  *SearchDentriesRequestOptionCreateTimeRange `json:"createTimeRange,omitempty" xml:"createTimeRange,omitempty" type:"Struct"`
	CreatorIds       []*string                                   `json:"creatorIds,omitempty" xml:"creatorIds,omitempty" type:"Repeated"`
	DentryCategories []*string                                   `json:"dentryCategories,omitempty" xml:"dentryCategories,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults  *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	ModifierIds []*string `json:"modifierIds,omitempty" xml:"modifierIds,omitempty" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken      *string                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SpaceIds       []*int64                                   `json:"spaceIds,omitempty" xml:"spaceIds,omitempty" type:"Repeated"`
	VisitTimeRange *SearchDentriesRequestOptionVisitTimeRange `json:"visitTimeRange,omitempty" xml:"visitTimeRange,omitempty" type:"Struct"`
}

func (s SearchDentriesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesRequestOption) GoString() string {
	return s.String()
}

func (s *SearchDentriesRequestOption) SetCreateTimeRange(v *SearchDentriesRequestOptionCreateTimeRange) *SearchDentriesRequestOption {
	s.CreateTimeRange = v
	return s
}

func (s *SearchDentriesRequestOption) SetCreatorIds(v []*string) *SearchDentriesRequestOption {
	s.CreatorIds = v
	return s
}

func (s *SearchDentriesRequestOption) SetDentryCategories(v []*string) *SearchDentriesRequestOption {
	s.DentryCategories = v
	return s
}

func (s *SearchDentriesRequestOption) SetMaxResults(v int32) *SearchDentriesRequestOption {
	s.MaxResults = &v
	return s
}

func (s *SearchDentriesRequestOption) SetModifierIds(v []*string) *SearchDentriesRequestOption {
	s.ModifierIds = v
	return s
}

func (s *SearchDentriesRequestOption) SetNextToken(v string) *SearchDentriesRequestOption {
	s.NextToken = &v
	return s
}

func (s *SearchDentriesRequestOption) SetSpaceIds(v []*int64) *SearchDentriesRequestOption {
	s.SpaceIds = v
	return s
}

func (s *SearchDentriesRequestOption) SetVisitTimeRange(v *SearchDentriesRequestOptionVisitTimeRange) *SearchDentriesRequestOption {
	s.VisitTimeRange = v
	return s
}

type SearchDentriesRequestOptionCreateTimeRange struct {
	// example:
	//
	// end_time
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// start_time
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s SearchDentriesRequestOptionCreateTimeRange) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesRequestOptionCreateTimeRange) GoString() string {
	return s.String()
}

func (s *SearchDentriesRequestOptionCreateTimeRange) SetEndTime(v int64) *SearchDentriesRequestOptionCreateTimeRange {
	s.EndTime = &v
	return s
}

func (s *SearchDentriesRequestOptionCreateTimeRange) SetStartTime(v int64) *SearchDentriesRequestOptionCreateTimeRange {
	s.StartTime = &v
	return s
}

type SearchDentriesRequestOptionVisitTimeRange struct {
	// example:
	//
	// end_time
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// start_time
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s SearchDentriesRequestOptionVisitTimeRange) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesRequestOptionVisitTimeRange) GoString() string {
	return s.String()
}

func (s *SearchDentriesRequestOptionVisitTimeRange) SetEndTime(v int64) *SearchDentriesRequestOptionVisitTimeRange {
	s.EndTime = &v
	return s
}

func (s *SearchDentriesRequestOptionVisitTimeRange) SetStartTime(v int64) *SearchDentriesRequestOptionVisitTimeRange {
	s.StartTime = &v
	return s
}

type SearchDentriesResponseBody struct {
	Items []*SearchDentriesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s SearchDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDentriesResponseBody) SetItems(v []*SearchDentriesResponseBodyItems) *SearchDentriesResponseBody {
	s.Items = v
	return s
}

func (s *SearchDentriesResponseBody) SetNextToken(v string) *SearchDentriesResponseBody {
	s.NextToken = &v
	return s
}

type SearchDentriesResponseBodyItems struct {
	Creator *SearchDentriesResponseBodyItemsCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// example:
	//
	// uuid
	DentryUuid     *string                                  `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	LastModifyTime *int64                                   `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Modifier       *SearchDentriesResponseBodyItemsModifier `json:"modifier,omitempty" xml:"modifier,omitempty" type:"Struct"`
	// example:
	//
	// name
	Name *string                              `json:"name,omitempty" xml:"name,omitempty"`
	Path *SearchDentriesResponseBodyItemsPath `json:"path,omitempty" xml:"path,omitempty" type:"Struct"`
}

func (s SearchDentriesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *SearchDentriesResponseBodyItems) SetCreator(v *SearchDentriesResponseBodyItemsCreator) *SearchDentriesResponseBodyItems {
	s.Creator = v
	return s
}

func (s *SearchDentriesResponseBodyItems) SetDentryUuid(v string) *SearchDentriesResponseBodyItems {
	s.DentryUuid = &v
	return s
}

func (s *SearchDentriesResponseBodyItems) SetLastModifyTime(v int64) *SearchDentriesResponseBodyItems {
	s.LastModifyTime = &v
	return s
}

func (s *SearchDentriesResponseBodyItems) SetModifier(v *SearchDentriesResponseBodyItemsModifier) *SearchDentriesResponseBodyItems {
	s.Modifier = v
	return s
}

func (s *SearchDentriesResponseBodyItems) SetName(v string) *SearchDentriesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *SearchDentriesResponseBodyItems) SetPath(v *SearchDentriesResponseBodyItemsPath) *SearchDentriesResponseBodyItems {
	s.Path = v
	return s
}

type SearchDentriesResponseBodyItemsCreator struct {
	// example:
	//
	// user_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// staff_id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchDentriesResponseBodyItemsCreator) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesResponseBodyItemsCreator) GoString() string {
	return s.String()
}

func (s *SearchDentriesResponseBodyItemsCreator) SetName(v string) *SearchDentriesResponseBodyItemsCreator {
	s.Name = &v
	return s
}

func (s *SearchDentriesResponseBodyItemsCreator) SetUserId(v string) *SearchDentriesResponseBodyItemsCreator {
	s.UserId = &v
	return s
}

type SearchDentriesResponseBodyItemsModifier struct {
	// example:
	//
	// user_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// staff_id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchDentriesResponseBodyItemsModifier) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesResponseBodyItemsModifier) GoString() string {
	return s.String()
}

func (s *SearchDentriesResponseBodyItemsModifier) SetName(v string) *SearchDentriesResponseBodyItemsModifier {
	s.Name = &v
	return s
}

func (s *SearchDentriesResponseBodyItemsModifier) SetUserId(v string) *SearchDentriesResponseBodyItemsModifier {
	s.UserId = &v
	return s
}

type SearchDentriesResponseBodyItemsPath struct {
	// example:
	//
	// folderA/folderB
	LongPath *string `json:"longPath,omitempty" xml:"longPath,omitempty"`
	// example:
	//
	// folderA:folderB
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SearchDentriesResponseBodyItemsPath) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesResponseBodyItemsPath) GoString() string {
	return s.String()
}

func (s *SearchDentriesResponseBodyItemsPath) SetLongPath(v string) *SearchDentriesResponseBodyItemsPath {
	s.LongPath = &v
	return s
}

func (s *SearchDentriesResponseBodyItemsPath) SetPath(v string) *SearchDentriesResponseBodyItemsPath {
	s.Path = &v
	return s
}

func (s *SearchDentriesResponseBodyItemsPath) SetUrl(v string) *SearchDentriesResponseBodyItemsPath {
	s.Url = &v
	return s
}

type SearchDentriesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchDentriesResponse) GoString() string {
	return s.String()
}

func (s *SearchDentriesResponse) SetHeaders(v map[string]*string) *SearchDentriesResponse {
	s.Headers = v
	return s
}

func (s *SearchDentriesResponse) SetStatusCode(v int32) *SearchDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDentriesResponse) SetBody(v *SearchDentriesResponseBody) *SearchDentriesResponse {
	s.Body = v
	return s
}

type SearchPublishDentriesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchPublishDentriesHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchPublishDentriesHeaders) GoString() string {
	return s.String()
}

func (s *SearchPublishDentriesHeaders) SetCommonHeaders(v map[string]*string) *SearchPublishDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchPublishDentriesHeaders) SetXAcsDingtalkAccessToken(v string) *SearchPublishDentriesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchPublishDentriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// keyword
	Keyword *string                             `json:"keyword,omitempty" xml:"keyword,omitempty"`
	Option  *SearchPublishDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// workspace_id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s SearchPublishDentriesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchPublishDentriesRequest) GoString() string {
	return s.String()
}

func (s *SearchPublishDentriesRequest) SetKeyword(v string) *SearchPublishDentriesRequest {
	s.Keyword = &v
	return s
}

func (s *SearchPublishDentriesRequest) SetOption(v *SearchPublishDentriesRequestOption) *SearchPublishDentriesRequest {
	s.Option = v
	return s
}

func (s *SearchPublishDentriesRequest) SetWorkspaceId(v string) *SearchPublishDentriesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SearchPublishDentriesRequest) SetOperatorId(v string) *SearchPublishDentriesRequest {
	s.OperatorId = &v
	return s
}

type SearchPublishDentriesRequestOption struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s SearchPublishDentriesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s SearchPublishDentriesRequestOption) GoString() string {
	return s.String()
}

func (s *SearchPublishDentriesRequestOption) SetMaxResults(v int32) *SearchPublishDentriesRequestOption {
	s.MaxResults = &v
	return s
}

func (s *SearchPublishDentriesRequestOption) SetNextToken(v string) *SearchPublishDentriesRequestOption {
	s.NextToken = &v
	return s
}

type SearchPublishDentriesResponseBody struct {
	Items []*SearchPublishDentriesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s SearchPublishDentriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchPublishDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchPublishDentriesResponseBody) SetItems(v []*SearchPublishDentriesResponseBodyItems) *SearchPublishDentriesResponseBody {
	s.Items = v
	return s
}

func (s *SearchPublishDentriesResponseBody) SetNextToken(v string) *SearchPublishDentriesResponseBody {
	s.NextToken = &v
	return s
}

type SearchPublishDentriesResponseBodyItems struct {
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// folderA/folderB
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// name
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SearchPublishDentriesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s SearchPublishDentriesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *SearchPublishDentriesResponseBodyItems) SetName(v string) *SearchPublishDentriesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *SearchPublishDentriesResponseBodyItems) SetPath(v string) *SearchPublishDentriesResponseBodyItems {
	s.Path = &v
	return s
}

func (s *SearchPublishDentriesResponseBodyItems) SetSummary(v string) *SearchPublishDentriesResponseBodyItems {
	s.Summary = &v
	return s
}

func (s *SearchPublishDentriesResponseBodyItems) SetUrl(v string) *SearchPublishDentriesResponseBodyItems {
	s.Url = &v
	return s
}

type SearchPublishDentriesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchPublishDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchPublishDentriesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchPublishDentriesResponse) GoString() string {
	return s.String()
}

func (s *SearchPublishDentriesResponse) SetHeaders(v map[string]*string) *SearchPublishDentriesResponse {
	s.Headers = v
	return s
}

func (s *SearchPublishDentriesResponse) SetStatusCode(v int32) *SearchPublishDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchPublishDentriesResponse) SetBody(v *SearchPublishDentriesResponseBody) *SearchPublishDentriesResponse {
	s.Body = v
	return s
}

type SearchWorkspacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchWorkspacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspacesHeaders) GoString() string {
	return s.String()
}

func (s *SearchWorkspacesHeaders) SetCommonHeaders(v map[string]*string) *SearchWorkspacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchWorkspacesHeaders) SetXAcsDingtalkAccessToken(v string) *SearchWorkspacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchWorkspacesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// keyword
	Keyword *string                        `json:"keyword,omitempty" xml:"keyword,omitempty"`
	Option  *SearchWorkspacesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s SearchWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *SearchWorkspacesRequest) SetKeyword(v string) *SearchWorkspacesRequest {
	s.Keyword = &v
	return s
}

func (s *SearchWorkspacesRequest) SetOption(v *SearchWorkspacesRequestOption) *SearchWorkspacesRequest {
	s.Option = v
	return s
}

func (s *SearchWorkspacesRequest) SetOperatorId(v string) *SearchWorkspacesRequest {
	s.OperatorId = &v
	return s
}

type SearchWorkspacesRequestOption struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s SearchWorkspacesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspacesRequestOption) GoString() string {
	return s.String()
}

func (s *SearchWorkspacesRequestOption) SetMaxResults(v int32) *SearchWorkspacesRequestOption {
	s.MaxResults = &v
	return s
}

func (s *SearchWorkspacesRequestOption) SetNextToken(v string) *SearchWorkspacesRequestOption {
	s.NextToken = &v
	return s
}

type SearchWorkspacesResponseBody struct {
	Items []*SearchWorkspacesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s SearchWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchWorkspacesResponseBody) SetItems(v []*SearchWorkspacesResponseBodyItems) *SearchWorkspacesResponseBody {
	s.Items = v
	return s
}

func (s *SearchWorkspacesResponseBody) SetNextToken(v string) *SearchWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

type SearchWorkspacesResponseBodyItems struct {
	// example:
	//
	// workspace_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// workspace_url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// workspace_id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SearchWorkspacesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspacesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *SearchWorkspacesResponseBodyItems) SetName(v string) *SearchWorkspacesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *SearchWorkspacesResponseBodyItems) SetUrl(v string) *SearchWorkspacesResponseBodyItems {
	s.Url = &v
	return s
}

func (s *SearchWorkspacesResponseBodyItems) SetWorkspaceId(v string) *SearchWorkspacesResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

type SearchWorkspacesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *SearchWorkspacesResponse) SetHeaders(v map[string]*string) *SearchWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *SearchWorkspacesResponse) SetStatusCode(v int32) *SearchWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchWorkspacesResponse) SetBody(v *SearchWorkspacesResponseBody) *SearchWorkspacesResponse {
	s.Body = v
	return s
}

type SetPermissionInheritanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetPermissionInheritanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetPermissionInheritanceHeaders) GoString() string {
	return s.String()
}

func (s *SetPermissionInheritanceHeaders) SetCommonHeaders(v map[string]*string) *SetPermissionInheritanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetPermissionInheritanceHeaders) SetXAcsDingtalkAccessToken(v string) *SetPermissionInheritanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetPermissionInheritanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PASS_ON
	Inheritance *string `json:"inheritance,omitempty" xml:"inheritance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SetPermissionInheritanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPermissionInheritanceRequest) GoString() string {
	return s.String()
}

func (s *SetPermissionInheritanceRequest) SetInheritance(v string) *SetPermissionInheritanceRequest {
	s.Inheritance = &v
	return s
}

func (s *SetPermissionInheritanceRequest) SetUnionId(v string) *SetPermissionInheritanceRequest {
	s.UnionId = &v
	return s
}

type SetPermissionInheritanceResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetPermissionInheritanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPermissionInheritanceResponseBody) GoString() string {
	return s.String()
}

func (s *SetPermissionInheritanceResponseBody) SetSuccess(v bool) *SetPermissionInheritanceResponseBody {
	s.Success = &v
	return s
}

type SetPermissionInheritanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPermissionInheritanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPermissionInheritanceResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPermissionInheritanceResponse) GoString() string {
	return s.String()
}

func (s *SetPermissionInheritanceResponse) SetHeaders(v map[string]*string) *SetPermissionInheritanceResponse {
	s.Headers = v
	return s
}

func (s *SetPermissionInheritanceResponse) SetStatusCode(v int32) *SetPermissionInheritanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPermissionInheritanceResponse) SetBody(v *SetPermissionInheritanceResponseBody) *SetPermissionInheritanceResponse {
	s.Body = v
	return s
}

type SetPermissionShareScopeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetPermissionShareScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetPermissionShareScopeHeaders) GoString() string {
	return s.String()
}

func (s *SetPermissionShareScopeHeaders) SetCommonHeaders(v map[string]*string) *SetPermissionShareScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetPermissionShareScopeHeaders) SetXAcsDingtalkAccessToken(v string) *SetPermissionShareScopeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetPermissionShareScopeRequest struct {
	Option *SetPermissionShareScopeRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// ORG_READ
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// union_id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SetPermissionShareScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPermissionShareScopeRequest) GoString() string {
	return s.String()
}

func (s *SetPermissionShareScopeRequest) SetOption(v *SetPermissionShareScopeRequestOption) *SetPermissionShareScopeRequest {
	s.Option = v
	return s
}

func (s *SetPermissionShareScopeRequest) SetScope(v string) *SetPermissionShareScopeRequest {
	s.Scope = &v
	return s
}

func (s *SetPermissionShareScopeRequest) SetUnionId(v string) *SetPermissionShareScopeRequest {
	s.UnionId = &v
	return s
}

type SetPermissionShareScopeRequestOption struct {
	// example:
	//
	// true
	CanSearch *bool `json:"canSearch,omitempty" xml:"canSearch,omitempty"`
}

func (s SetPermissionShareScopeRequestOption) String() string {
	return tea.Prettify(s)
}

func (s SetPermissionShareScopeRequestOption) GoString() string {
	return s.String()
}

func (s *SetPermissionShareScopeRequestOption) SetCanSearch(v bool) *SetPermissionShareScopeRequestOption {
	s.CanSearch = &v
	return s
}

type SetPermissionShareScopeResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetPermissionShareScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPermissionShareScopeResponseBody) GoString() string {
	return s.String()
}

func (s *SetPermissionShareScopeResponseBody) SetSuccess(v bool) *SetPermissionShareScopeResponseBody {
	s.Success = &v
	return s
}

type SetPermissionShareScopeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPermissionShareScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPermissionShareScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPermissionShareScopeResponse) GoString() string {
	return s.String()
}

func (s *SetPermissionShareScopeResponse) SetHeaders(v map[string]*string) *SetPermissionShareScopeResponse {
	s.Headers = v
	return s
}

func (s *SetPermissionShareScopeResponse) SetStatusCode(v int32) *SetPermissionShareScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPermissionShareScopeResponse) SetBody(v *SetPermissionShareScopeResponseBody) *SetPermissionShareScopeResponse {
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
// 添加权限
//
// @param request - AddPermissionRequest
//
// @param headers - AddPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPermissionResponse
func (client *Client) AddPermissionWithOptions(dentryUuid *string, request *AddPermissionRequest, headers *AddPermissionHeaders, runtime *util.RuntimeOptions) (_result *AddPermissionResponse, _err error) {
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
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/dentries/" + tea.StringValue(dentryUuid) + "/permissions"),
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
func (client *Client) AddPermission(dentryUuid *string, request *AddPermissionRequest) (_result *AddPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddPermissionHeaders{}
	_result = &AddPermissionResponse{}
	_body, _err := client.AddPermissionWithOptions(dentryUuid, request, headers, runtime)
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
// @param request - BatchQueryRolesRequest
//
// @param headers - BatchQueryRolesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryRolesResponse
func (client *Client) BatchQueryRolesWithOptions(request *BatchQueryRolesRequest, headers *BatchQueryRolesHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryUuidList)) {
		body["dentryUuidList"] = request.DentryUuidList
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
		Action:      tea.String("BatchQueryRoles"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/dentries/permissions/roles/batchQuery"),
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
func (client *Client) BatchQueryRoles(request *BatchQueryRolesRequest) (_result *BatchQueryRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryRolesHeaders{}
	_result = &BatchQueryRolesResponse{}
	_body, _err := client.BatchQueryRolesWithOptions(request, headers, runtime)
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
func (client *Client) CommitFileWithOptions(parentDentryUuid *string, request *CommitFileRequest, headers *CommitFileHeaders, runtime *util.RuntimeOptions) (_result *CommitFileResponse, _err error) {
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
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/files/" + tea.StringValue(parentDentryUuid) + "/commit"),
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
func (client *Client) CommitFile(parentDentryUuid *string, request *CommitFileRequest) (_result *CommitFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CommitFileHeaders{}
	_result = &CommitFileResponse{}
	_body, _err := client.CommitFileWithOptions(parentDentryUuid, request, headers, runtime)
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
func (client *Client) DeletePermissionWithOptions(dentryUuid *string, request *DeletePermissionRequest, headers *DeletePermissionHeaders, runtime *util.RuntimeOptions) (_result *DeletePermissionResponse, _err error) {
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
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/dentries/" + tea.StringValue(dentryUuid) + "/permissions/remove"),
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
func (client *Client) DeletePermission(dentryUuid *string, request *DeletePermissionRequest) (_result *DeletePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeletePermissionHeaders{}
	_result = &DeletePermissionResponse{}
	_body, _err := client.DeletePermissionWithOptions(dentryUuid, request, headers, runtime)
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
func (client *Client) GetFileUploadInfoWithOptions(parentDentryUuid *string, request *GetFileUploadInfoRequest, headers *GetFileUploadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileUploadInfoResponse, _err error) {
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
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/files/" + tea.StringValue(parentDentryUuid) + "/uploadInfos/query"),
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
func (client *Client) GetFileUploadInfo(parentDentryUuid *string, request *GetFileUploadInfoRequest) (_result *GetFileUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileUploadInfoHeaders{}
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.GetFileUploadInfoWithOptions(parentDentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取权限继承模式
//
// @param request - GetPermissionInheritanceRequest
//
// @param headers - GetPermissionInheritanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPermissionInheritanceResponse
func (client *Client) GetPermissionInheritanceWithOptions(dentryUuid *string, request *GetPermissionInheritanceRequest, headers *GetPermissionInheritanceHeaders, runtime *util.RuntimeOptions) (_result *GetPermissionInheritanceResponse, _err error) {
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
		Action:      tea.String("GetPermissionInheritance"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/dentries/" + tea.StringValue(dentryUuid) + "/permissions/inheritances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPermissionInheritanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取权限继承模式
//
// @param request - GetPermissionInheritanceRequest
//
// @return GetPermissionInheritanceResponse
func (client *Client) GetPermissionInheritance(dentryUuid *string, request *GetPermissionInheritanceRequest) (_result *GetPermissionInheritanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPermissionInheritanceHeaders{}
	_result = &GetPermissionInheritanceResponse{}
	_body, _err := client.GetPermissionInheritanceWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取分享范围
//
// @param request - GetPermissionShareScopeRequest
//
// @param headers - GetPermissionShareScopeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPermissionShareScopeResponse
func (client *Client) GetPermissionShareScopeWithOptions(dentryUuid *string, request *GetPermissionShareScopeRequest, headers *GetPermissionShareScopeHeaders, runtime *util.RuntimeOptions) (_result *GetPermissionShareScopeResponse, _err error) {
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
		Action:      tea.String("GetPermissionShareScope"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/dentries/" + tea.StringValue(dentryUuid) + "/permissions/scopes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPermissionShareScopeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取分享范围
//
// @param request - GetPermissionShareScopeRequest
//
// @return GetPermissionShareScopeResponse
func (client *Client) GetPermissionShareScope(dentryUuid *string, request *GetPermissionShareScopeRequest) (_result *GetPermissionShareScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPermissionShareScopeHeaders{}
	_result = &GetPermissionShareScopeResponse{}
	_body, _err := client.GetPermissionShareScopeWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询操作日志
//
// @param request - ListOperationLogsRequest
//
// @param headers - ListOperationLogsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationLogsResponse
func (client *Client) ListOperationLogsWithOptions(request *ListOperationLogsRequest, headers *ListOperationLogsHeaders, runtime *util.RuntimeOptions) (_result *ListOperationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOperationLogs"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/managements/operationLogs/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOperationLogsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询操作日志
//
// @param request - ListOperationLogsRequest
//
// @return ListOperationLogsResponse
func (client *Client) ListOperationLogs(request *ListOperationLogsRequest) (_result *ListOperationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListOperationLogsHeaders{}
	_result = &ListOperationLogsResponse{}
	_body, _err := client.ListOperationLogsWithOptions(request, headers, runtime)
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
func (client *Client) ListPermissionsWithOptions(dentryUuid *string, request *ListPermissionsRequest, headers *ListPermissionsHeaders, runtime *util.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
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
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/dentries/" + tea.StringValue(dentryUuid) + "/permissions/query"),
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
func (client *Client) ListPermissions(dentryUuid *string, request *ListPermissionsRequest) (_result *ListPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListPermissionsHeaders{}
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询员工离职时空间默认转交人(管理员)
//
// @param request - ManagerGetDefaultHandOverUserRequest
//
// @param headers - ManagerGetDefaultHandOverUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManagerGetDefaultHandOverUserResponse
func (client *Client) ManagerGetDefaultHandOverUserWithOptions(request *ManagerGetDefaultHandOverUserRequest, headers *ManagerGetDefaultHandOverUserHeaders, runtime *util.RuntimeOptions) (_result *ManagerGetDefaultHandOverUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
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
		Action:      tea.String("ManagerGetDefaultHandOverUser"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/managementSettings/defaultHandOverUsers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ManagerGetDefaultHandOverUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询员工离职时空间默认转交人(管理员)
//
// @param request - ManagerGetDefaultHandOverUserRequest
//
// @return ManagerGetDefaultHandOverUserResponse
func (client *Client) ManagerGetDefaultHandOverUser(request *ManagerGetDefaultHandOverUserRequest) (_result *ManagerGetDefaultHandOverUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ManagerGetDefaultHandOverUserHeaders{}
	_result = &ManagerGetDefaultHandOverUserResponse{}
	_body, _err := client.ManagerGetDefaultHandOverUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置员工离职时空间默认转交人(管理员)
//
// @param request - ManagerSetDefaultHandOverUserRequest
//
// @param headers - ManagerSetDefaultHandOverUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManagerSetDefaultHandOverUserResponse
func (client *Client) ManagerSetDefaultHandOverUserWithOptions(request *ManagerSetDefaultHandOverUserRequest, headers *ManagerSetDefaultHandOverUserHeaders, runtime *util.RuntimeOptions) (_result *ManagerSetDefaultHandOverUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefaultHandoverUserId)) {
		body["defaultHandoverUserId"] = request.DefaultHandoverUserId
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
		Action:      tea.String("ManagerSetDefaultHandOverUser"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/managementSettings/defaultHandOverUsers/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ManagerSetDefaultHandOverUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置员工离职时空间默认转交人(管理员)
//
// @param request - ManagerSetDefaultHandOverUserRequest
//
// @return ManagerSetDefaultHandOverUserResponse
func (client *Client) ManagerSetDefaultHandOverUser(request *ManagerSetDefaultHandOverUserRequest) (_result *ManagerSetDefaultHandOverUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ManagerSetDefaultHandOverUserHeaders{}
	_result = &ManagerSetDefaultHandOverUserResponse{}
	_body, _err := client.ManagerSetDefaultHandOverUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 搜索文件
//
// @param request - SearchDentriesRequest
//
// @param headers - SearchDentriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDentriesResponse
func (client *Client) SearchDentriesWithOptions(request *SearchDentriesRequest, headers *SearchDentriesHeaders, runtime *util.RuntimeOptions) (_result *SearchDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["keyword"] = request.Keyword
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
		Action:      tea.String("SearchDentries"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/dentries/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索文件
//
// @param request - SearchDentriesRequest
//
// @return SearchDentriesResponse
func (client *Client) SearchDentries(request *SearchDentriesRequest) (_result *SearchDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchDentriesHeaders{}
	_result = &SearchDentriesResponse{}
	_body, _err := client.SearchDentriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 搜索公开发布文件
//
// @param request - SearchPublishDentriesRequest
//
// @param headers - SearchPublishDentriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchPublishDentriesResponse
func (client *Client) SearchPublishDentriesWithOptions(request *SearchPublishDentriesRequest, headers *SearchPublishDentriesHeaders, runtime *util.RuntimeOptions) (_result *SearchPublishDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["workspaceId"] = request.WorkspaceId
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
		Action:      tea.String("SearchPublishDentries"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/publishDentries/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchPublishDentriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索公开发布文件
//
// @param request - SearchPublishDentriesRequest
//
// @return SearchPublishDentriesResponse
func (client *Client) SearchPublishDentries(request *SearchPublishDentriesRequest) (_result *SearchPublishDentriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchPublishDentriesHeaders{}
	_result = &SearchPublishDentriesResponse{}
	_body, _err := client.SearchPublishDentriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 搜索知识库
//
// @param request - SearchWorkspacesRequest
//
// @param headers - SearchWorkspacesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchWorkspacesResponse
func (client *Client) SearchWorkspacesWithOptions(request *SearchWorkspacesRequest, headers *SearchWorkspacesHeaders, runtime *util.RuntimeOptions) (_result *SearchWorkspacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["keyword"] = request.Keyword
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
		Action:      tea.String("SearchWorkspaces"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/workspaces/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchWorkspacesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索知识库
//
// @param request - SearchWorkspacesRequest
//
// @return SearchWorkspacesResponse
func (client *Client) SearchWorkspaces(request *SearchWorkspacesRequest) (_result *SearchWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchWorkspacesHeaders{}
	_result = &SearchWorkspacesResponse{}
	_body, _err := client.SearchWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置权限继承模式
//
// @param request - SetPermissionInheritanceRequest
//
// @param headers - SetPermissionInheritanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPermissionInheritanceResponse
func (client *Client) SetPermissionInheritanceWithOptions(dentryUuid *string, request *SetPermissionInheritanceRequest, headers *SetPermissionInheritanceHeaders, runtime *util.RuntimeOptions) (_result *SetPermissionInheritanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Inheritance)) {
		body["inheritance"] = request.Inheritance
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
		Action:      tea.String("SetPermissionInheritance"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/dentries/" + tea.StringValue(dentryUuid) + "/permissions/inheritances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPermissionInheritanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置权限继承模式
//
// @param request - SetPermissionInheritanceRequest
//
// @return SetPermissionInheritanceResponse
func (client *Client) SetPermissionInheritance(dentryUuid *string, request *SetPermissionInheritanceRequest) (_result *SetPermissionInheritanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetPermissionInheritanceHeaders{}
	_result = &SetPermissionInheritanceResponse{}
	_body, _err := client.SetPermissionInheritanceWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置分享范围
//
// @param request - SetPermissionShareScopeRequest
//
// @param headers - SetPermissionShareScopeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPermissionShareScopeResponse
func (client *Client) SetPermissionShareScopeWithOptions(dentryUuid *string, request *SetPermissionShareScopeRequest, headers *SetPermissionShareScopeHeaders, runtime *util.RuntimeOptions) (_result *SetPermissionShareScopeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
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
		Action:      tea.String("SetPermissionShareScope"),
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/dentries/" + tea.StringValue(dentryUuid) + "/permissions/scopes"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPermissionShareScopeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置分享范围
//
// @param request - SetPermissionShareScopeRequest
//
// @return SetPermissionShareScopeResponse
func (client *Client) SetPermissionShareScope(dentryUuid *string, request *SetPermissionShareScopeRequest) (_result *SetPermissionShareScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetPermissionShareScopeHeaders{}
	_result = &SetPermissionShareScopeResponse{}
	_body, _err := client.SetPermissionShareScopeWithOptions(dentryUuid, request, headers, runtime)
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
func (client *Client) UpdatePermissionWithOptions(dentryUuid *string, request *UpdatePermissionRequest, headers *UpdatePermissionHeaders, runtime *util.RuntimeOptions) (_result *UpdatePermissionResponse, _err error) {
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
		Version:     tea.String("storage_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/storage/spaces/dentries/" + tea.StringValue(dentryUuid) + "/permissions"),
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
func (client *Client) UpdatePermission(dentryUuid *string, request *UpdatePermissionRequest) (_result *UpdatePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePermissionHeaders{}
	_result = &UpdatePermissionResponse{}
	_body, _err := client.UpdatePermissionWithOptions(dentryUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
