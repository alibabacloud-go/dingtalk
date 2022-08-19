// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package storage_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// 名称(文件名+后缀), 规则：
	// 1. 头尾不能包含空格，否则会自动去除
	// 2. 不能包含特殊字符，包括：制表符、*、"、<、>、|
	// 3. 不能以"."结尾
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 可选参数
	Option *AddFolderRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 用户id
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
	// 文件夹名称冲突策略
	// 枚举值:
	// 	AUTO_RENAME: 自动重命名
	// 	OVERWRITE: 覆盖
	// 	RETURN_DENTRY_IF_EXISTS: 返回已存在文件
	// 	RETURN_ERROR_IF_EXISTS: 文件已存在时报错
	// 默认值:
	// 	AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
}

func (s AddFolderRequestOption) String() string {
	return tea.Prettify(s)
}

func (s AddFolderRequestOption) GoString() string {
	return s.String()
}

func (s *AddFolderRequestOption) SetConflictStrategy(v string) *AddFolderRequestOption {
	s.ConflictStrategy = &v
	return s
}

type AddFolderResponseBody struct {
	// 文件夹信息
	// dentry.type等于FOLDER表示是文件夹
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 后缀
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父目录id, 根目录id值为0
	// 空值代表根目录的parentId不存在
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 存储分区，目前包括公有云OSS存储分区和专属Mini OSS存储分区
	// 枚举值:
	// 	PUBLIC_OSS_PARTITION: 公有云OSS存储分区
	// 	MINI_OSS_PARTITION: 专属Mini OSS存储分区
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// 路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 属性
	Properties *AddFolderResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// 大小, 单位:Byte
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETED: 已删除
	// 	EXPIRED: 已过期
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 驱动类型
	// 枚举值:
	// 	DINGTALK: 钉钉统一存储驱动
	// 	ALIDOC: 钉钉文档存储驱动
	// 	SHANJI: 闪记存储驱动
	// 	UNKNOWN: 未知驱动
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// 类型，目录或文件
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// uuid，如移动文件，此字段不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AddFolderResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s AddFolderResponseBodyDentry) GoString() string {
	return s.String()
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
	// 文件是否只读
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFolderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 权限成员列表
	Members []*AddPermissionRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 可选参数
	Option *AddPermissionRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 权限角色id
	RoleId *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 用户id
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
	// 权限归属的企业
	// 如果存在企业id, 对应member离职的时候会自动清理权限
	// 如果memberType是dept类型，必须要有企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 权限成员id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 权限成员类型
	// 枚举值:
	// 	ORG: 企业
	// 	DEPT: 部门
	// 	TAG: 自定义tag
	// 	CONVERSATION: 会话
	// 	GG: 通用组
	// 	USER: 用户
	// 	ALL_USERS: 所有用户
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
	// 有效时间(秒)
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
	// 本次操作是否成功
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 可选参数
	Option *AddSpaceRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 用户id
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
	// 空间能力项, 默认表示不设置拓展能力项
	Capabilities *AddSpaceRequestOptionCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// 空间名称，默认无空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// owner类型, 空间Owner可以是用户或应用
	// 如果是应用类型，需要单独授权
	// 枚举值:
	// 	USER: 用户类型
	// 	APP: App类型
	// 默认值:
	// 	USER
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// 空间能使用最大容量, 默认表示无最大容量
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 空间场景，详见 Space.scene 字段. 不指定默认值是default
	// 只能由数字和字母组成
	// 默认值:
	// 	default
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 空间场景Id，详见 Space.sceneId 字段. 不指定默认值是0
	// 只能由数字和字母组成
	// 默认值:
	// 	0
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
	// 是否进最近使用, 默认不支持
	// 默认值:
	// 	false
	CanRecordRecentFile *bool `json:"canRecordRecentFile,omitempty" xml:"canRecordRecentFile,omitempty"`
	// 是否支持重命名空间名称, 默认不支持
	// 默认值:
	// 	false
	CanRename *bool `json:"canRename,omitempty" xml:"canRename,omitempty"`
	// 是否支持搜索, 默认不支持
	// 默认值:
	// 	false
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
	// 空间详情
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
	// 空间能力项
	Capabilities *AddSpaceResponseBodySpaceCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// 空间归属企业的id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 空间id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 所有者id, 根据ownerType定义, 确定值的所属类型
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// owner类型
	// 枚举值:
	// 	USER: 用户类型
	// 	APP: App类型
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// 总容量
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 业务场景，可以自定义，表示多个不同空间的聚合，可以提供对特定场景做能力设计、容量管理，如根据场景来做搜索或查询。
	// 创建空间时，不指定scene, 默认值是default
	// 默认值:
	// 	default
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 关联业务id, 配合scene一起使用。创建空间时，不指定sceneId， 默认值是0
	// 默认值:
	// 	0
	SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
	// 空间状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETE: 已删除
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 已使用容量
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s AddSpaceResponseBodySpace) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceResponseBodySpace) GoString() string {
	return s.String()
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
	// 是否进最近使用, 默认不支持
	// 默认值:
	// 	false
	CanRecordRecentFile *bool `json:"canRecordRecentFile,omitempty" xml:"canRecordRecentFile,omitempty"`
	// 是否支持重命名空间名称, 默认不支持
	// 默认值:
	// 	false
	CanRename *bool `json:"canRename,omitempty" xml:"canRename,omitempty"`
	// 是否支持搜索, 默认不支持
	// 默认值:
	// 	false
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

type AddSpaceResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddSpaceResponse) SetBody(v *AddSpaceResponseBody) *AddSpaceResponse {
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
	// 用户id
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
	// 本次操作是否成功
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 名称(文件名+后缀), 规则：
	// 1. 头尾不能包含空格，否则会自动去除
	// 2. 不能包含特殊字符，包括：制表符、*、"、<、>、|
	// 3. 不能以"."结尾
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 可选参数
	Option *CommitFileRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 父目录id, 根目录id值为0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 添加文件唯一标识，可通过DentryService.getUploadInfo来生成
	UploadKey *string `json:"uploadKey,omitempty" xml:"uploadKey,omitempty"`
	// 用户id
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
	// 文件名称冲突策略
	// 枚举值:
	// 	AUTO_RENAME: 自动重命名
	// 	OVERWRITE: 覆盖
	// 	RETURN_DENTRY_IF_EXISTS: 返回已存在文件
	// 	RETURN_ERROR_IF_EXISTS: 文件已存在时报错
	// 默认值:
	// 	AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
	// 默认文件大小, 单位:Byte
	// 如果此字段不为空，企业存储系统会校验文件实际大小是否和此字段是否一致，不一致会报错
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s CommitFileRequestOption) String() string {
	return tea.Prettify(s)
}

func (s CommitFileRequestOption) GoString() string {
	return s.String()
}

func (s *CommitFileRequestOption) SetConflictStrategy(v string) *CommitFileRequestOption {
	s.ConflictStrategy = &v
	return s
}

func (s *CommitFileRequestOption) SetSize(v int64) *CommitFileRequestOption {
	s.Size = &v
	return s
}

type CommitFileResponseBody struct {
	// 文件信息
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 后缀
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父目录id, 根目录id值为0
	// 空值代表根目录的parentId不存在
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 存储分区，目前包括公有云OSS存储分区和专属Mini OSS存储分区
	// 枚举值:
	// 	PUBLIC_OSS_PARTITION: 公有云OSS存储分区
	// 	MINI_OSS_PARTITION: 专属Mini OSS存储分区
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// 路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 属性
	Properties *CommitFileResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// 大小, 单位:Byte
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETED: 已删除
	// 	EXPIRED: 已过期
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 驱动类型
	// 枚举值:
	// 	DINGTALK: 钉钉统一存储驱动
	// 	ALIDOC: 钉钉文档存储驱动
	// 	SHANJI: 闪记存储驱动
	// 	UNKNOWN: 未知驱动
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// 类型，目录或文件
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// uuid，如移动文件，此字段不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CommitFileResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s CommitFileResponseBodyDentry) GoString() string {
	return s.String()
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
	// 文件是否只读
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

type CommitFileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CommitFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CommitFileResponse) SetBody(v *CommitFileResponseBody) *CommitFileResponse {
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
	// 可选参数
	Option *CopyDentryRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 目标文件夹id, 根目录id值为0
	TargetFolderId *string `json:"targetFolderId,omitempty" xml:"targetFolderId,omitempty"`
	// 目标文件夹空间id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// 用户id
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
	// 文件(夹)名称冲突策略
	// 枚举值:
	// 	AUTO_RENAME: 自动重命名
	// 	OVERWRITE: 覆盖
	// 	RETURN_DENTRY_IF_EXISTS: 返回已存在文件
	// 	RETURN_ERROR_IF_EXISTS: 文件已存在时报错
	// 默认值:
	// 	AUTO_RENAME
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
	// 是否是异步任务
	// 如果操作对象有子节点，则会异步处理
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// 文件信息
	Dentry *CopyDentryResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
	// 任务id，用于查询任务执行状态; 查询接口开发中
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 后缀
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父目录id, 根目录id值为0
	// 空值代表根目录的parentId不存在
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 存储分区，目前包括公有云OSS存储分区和专属Mini OSS存储分区
	// 枚举值:
	// 	PUBLIC_OSS_PARTITION: 公有云OSS存储分区
	// 	MINI_OSS_PARTITION: 专属Mini OSS存储分区
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// 路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 属性
	Properties *CopyDentryResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// 大小, 单位:Byte
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETED: 已删除
	// 	EXPIRED: 已过期
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 驱动类型
	// 枚举值:
	// 	DINGTALK: 钉钉统一存储驱动
	// 	ALIDOC: 钉钉文档存储驱动
	// 	SHANJI: 闪记存储驱动
	// 	UNKNOWN: 未知驱动
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// 类型，目录或文件
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// uuid，如移动文件，此字段不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CopyDentryResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s CopyDentryResponseBodyDentry) GoString() string {
	return s.String()
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
	// 文件是否只读
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CopyDentryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CopyDentryResponse) SetBody(v *CopyDentryResponseBody) *CopyDentryResponse {
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
	// 是否删除到回收站，默认不删除到回收站，直接删除
	// 默认值:
	// 	false
	ToRecycleBin *bool `json:"toRecycleBin,omitempty" xml:"toRecycleBin,omitempty"`
	// 用户id
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
	// 是否是异步任务
	// 如果操作对象有子节点，则会异步处理
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// 任务id，用于查询任务执行状态; 查询接口开发中
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDentryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 文件上App属性名称
	PropertyNames []*string `json:"propertyNames,omitempty" xml:"propertyNames,omitempty" type:"Repeated"`
	// 用户id
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
	// 本次操作是否成功
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDentryAppPropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 权限成员列表
	Members []*DeletePermissionRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 权限角色id
	RoleId *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 用户id
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
	// 权限归属的企业
	// 如果存在企业id, 对应member离职的时候会自动清理权限
	// 如果memberType是dept类型，必须要有企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 权限成员id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 权限成员类型
	// 枚举值:
	// 	ORG: 企业
	// 	DEPT: 部门
	// 	TAG: 自定义tag
	// 	CONVERSATION: 会话
	// 	GG: 通用组
	// 	USER: 用户
	// 	ALL_USERS: 所有用户
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
	// 本次操作是否成功
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 用户id
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
	// 本次操作是否成功
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRecycleItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 回收项id列表
	RecycleItemIds []*string `json:"recycleItemIds,omitempty" xml:"recycleItemIds,omitempty" type:"Repeated"`
	// 用户id
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
	// 本次操作是否成功
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRecycleItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 用户id
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
	// 企业存储应用信息
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
	// 开放平台应用appId
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// 应用归属企业的id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 应用创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 应用修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 应用名称，对应开放平台应用名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 分区容量信息
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
	// 分区类型
	// 枚举值:
	// 	PUBLIC_OSS_PARTITION: 公有云OSS存储分区
	// 	MINI_OSS_PARTITION: 专属Mini OSS存储分区
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// 容量信息
	Quota *GetCurrentAppResponseBodyAppPartitionsQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
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
	// 最大容量, 单位: Byte
	// 当前应用容量被设置为max时，代表当前应用容量设置了上限，used<=max
	// 当前应用容量未设置max时，返回空，此时应用共享该企业剩余可用容量
	Max *int64 `json:"max,omitempty" xml:"max,omitempty"`
	// 容量类型
	// 如果是企业维度容量，此值是PRIVATE, 表示企业独占
	// 枚举值:
	// 	SHARE: 共享容量
	// 此模式下，Quota.max为空，表示共享企业容量
	// 	PRIVATE: 专享容量
	// 当Quota.max设置值后，表示容量独占
	// 使用场景：当需要保证单个应用的可用容量不受其他应用影响时, 可使用共享容量
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 已使用容量, 单位: Byte
	// 表示该应用下所用文件占用容量的总和，文件的上传、复制、删除相关操作会对used的值做相应变更
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

func (s *GetCurrentAppResponseBodyAppPartitionsQuota) SetType(v string) *GetCurrentAppResponseBodyAppPartitionsQuota {
	s.Type = &v
	return s
}

func (s *GetCurrentAppResponseBodyAppPartitionsQuota) SetUsed(v int64) *GetCurrentAppResponseBodyAppPartitionsQuota {
	s.Used = &v
	return s
}

type GetCurrentAppResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCurrentAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetCurrentAppResponse) SetBody(v *GetCurrentAppResponseBody) *GetCurrentAppResponse {
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
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetDentryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDentryRequest) GoString() string {
	return s.String()
}

func (s *GetDentryRequest) SetUnionId(v string) *GetDentryRequest {
	s.UnionId = &v
	return s
}

type GetDentryResponseBody struct {
	// 文件(夹)信息
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 后缀
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父目录id, 根目录id值为0
	// 空值代表根目录的parentId不存在
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 存储分区，目前包括公有云OSS存储分区和专属Mini OSS存储分区
	// 枚举值:
	// 	PUBLIC_OSS_PARTITION: 公有云OSS存储分区
	// 	MINI_OSS_PARTITION: 专属Mini OSS存储分区
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// 路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 属性
	Properties *GetDentryResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// 大小, 单位:Byte
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETED: 已删除
	// 	EXPIRED: 已过期
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 驱动类型
	// 枚举值:
	// 	DINGTALK: 钉钉统一存储驱动
	// 	ALIDOC: 钉钉文档存储驱动
	// 	SHANJI: 闪记存储驱动
	// 	UNKNOWN: 未知驱动
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// 类型，目录或文件
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// uuid，如移动文件，此字段不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDentryResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s GetDentryResponseBodyDentry) GoString() string {
	return s.String()
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
	// 文件是否只读
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

type GetDentryResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDentryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDentryResponse) SetBody(v *GetDentryResponseBody) *GetDentryResponse {
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
	// 可选参数
	Option *GetFileDownloadInfoRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 用户id
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
	// 历史版本号
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetFileDownloadInfoRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetFileDownloadInfoRequestOption) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoRequestOption) SetVersion(v int64) *GetFileDownloadInfoRequestOption {
	s.Version = &v
	return s
}

type GetFileDownloadInfoResponseBody struct {
	// Header加签信息, 当protocol等于HEADER_SIGNATURE时，此字段生效
	HeaderSignatureInfo *GetFileDownloadInfoResponseBodyHeaderSignatureInfo `json:"headerSignatureInfo,omitempty" xml:"headerSignatureInfo,omitempty" type:"Struct"`
	// 文件下载协议
	// 枚举值:
	// 	HEADER_SIGNATURE: Header加签
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
	// 过期时间，单位秒
	ExpirationSeconds *int32 `json:"expirationSeconds,omitempty" xml:"expirationSeconds,omitempty"`
	// 请求头
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// 内网URL, 在网络连通的情况下，使用内网URL可加速服务器间上传
	InternalResourceUrls []*string `json:"internalResourceUrls,omitempty" xml:"internalResourceUrls,omitempty" type:"Repeated"`
	// 地域
	// 枚举值:
	// 	ZHANGJIAKOU: 张家口
	// 	SHENZHEN: 深圳
	// 	SHANGHAI: 上海
	// 	SINGAPORE: 新加坡
	// 	UNKNOWN: 未知
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// 多个上传下载URL, 前面url优先
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFileDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 是否需要分片上传
	// 5G以下文件，建议设为false，简化上传步骤
	// 5G以上文件，必须设为true, 否则上传会失败
	// 具体参考文档: https://help.aliyun.com/document_detail/84778.html
	Multipart *bool `json:"multipart,omitempty" xml:"multipart,omitempty"`
	// 可选参数
	Option *GetFileUploadInfoRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 通过指定上传协议返回不同协议上传所需要的信息
	// 对于部分企业开启了专属存储，必须实现HEADER加签，否则无法支持专属存储组织文件上传。
	// 如果指定上传协议不支持，则会返回错误Errors.DENTRY_UPLOAD_PROTOCOL_NOTSUPPORT, 请尝试用其它协议上传。
	// 枚举值:
	// 	HEADER_SIGNATURE: Header加签
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// 用户id
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
	// 预检查的字段。可实现对文件名称，文件完整性，容量的校验
	PreCheckParam *GetFileUploadInfoRequestOptionPreCheckParam `json:"preCheckParam,omitempty" xml:"preCheckParam,omitempty" type:"Struct"`
	// 优先地域, 倾向于将资源存到哪个地域，可实现就近上传等功能
	// 枚举值:
	// 	ZHANGJIAKOU: 张家口
	// 	SHENZHEN: 深圳
	// 	SHANGHAI: 上海
	// 	SINGAPORE: 新加坡
	// 	UNKNOWN: 未知
	PreferRegion *string `json:"preferRegion,omitempty" xml:"preferRegion,omitempty"`
	// 文件存储驱动类型, 当前只支持DINGTALK
	// 枚举值:
	// 	DINGTALK: 钉钉统一存储驱动
	// 	ALIDOC: 钉钉文档存储驱动
	// 	SHANJI: 闪记存储驱动
	// 	UNKNOWN: 未知驱动
	// 默认值:
	// 	DINGTALK
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

func (s *GetFileUploadInfoRequestOption) SetPreferRegion(v string) *GetFileUploadInfoRequestOption {
	s.PreferRegion = &v
	return s
}

func (s *GetFileUploadInfoRequestOption) SetStorageDriver(v string) *GetFileUploadInfoRequestOption {
	s.StorageDriver = &v
	return s
}

type GetFileUploadInfoRequestOptionPreCheckParam struct {
	// 文件md5值, 做文件完整性校验。不传则不做校验。
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 文件名称, 文件名称合法性和文件名称冲突校验
	// 规则：
	// 1. 头尾不能包含空格，否则会自动去除
	// 2. 不能包含特殊字符，包括：制表符、*、"、<、>、|
	// 3. 不能以"."结尾
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父目录id
	// 根目录id值为0
	// 用于同目录文件名冲突校验
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件大小, 做容量相关校验。不传则不做校验。
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
	// Header加签上传信息, 当protocol等于HEADER_SIGNATURE时，此字段生效
	HeaderSignatureInfo *GetFileUploadInfoResponseBodyHeaderSignatureInfo `json:"headerSignatureInfo,omitempty" xml:"headerSignatureInfo,omitempty" type:"Struct"`
	// 上传协议，根据不同上传类型返回对应的信息.
	// 枚举值:
	// 	HEADER_SIGNATURE: Header加签
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// 文件存储类型
	// 枚举值:
	// 	DINGTALK: 钉钉统一存储驱动
	// 	ALIDOC: 钉钉文档存储驱动
	// 	SHANJI: 闪记存储驱动
	// 	UNKNOWN: 未知驱动
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// 上传唯一标识
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
	// 过期时间，单位秒
	ExpirationSeconds *int32 `json:"expirationSeconds,omitempty" xml:"expirationSeconds,omitempty"`
	// 请求头
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// 内网URL, 在网络连通的情况下，使用内网URL可加速服务器间上传
	InternalResourceUrls []*string `json:"internalResourceUrls,omitempty" xml:"internalResourceUrls,omitempty" type:"Repeated"`
	// 地域
	// 枚举值:
	// 	ZHANGJIAKOU: 张家口
	// 	SHENZHEN: 深圳
	// 	SHANGHAI: 上海
	// 	SINGAPORE: 新加坡
	// 	UNKNOWN: 未知
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// 多个上传下载URL, 前面url优先
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFileUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetFileUploadInfoResponse) SetBody(v *GetFileUploadInfoResponseBody) *GetFileUploadInfoResponse {
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
	// 回收站范围类型
	// 枚举值:
	// 	ORG: 企业
	// 	APP: 应用
	// 	SPACE: 空间
	RecycleBinScope *string `json:"recycleBinScope,omitempty" xml:"recycleBinScope,omitempty"`
	// 回收站范围id
	// 根据recycleBinScope传入对应的企业、应用、空间ID
	ScopeId *string `json:"scopeId,omitempty" xml:"scopeId,omitempty"`
	// 用户id
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
	// 回收站信息
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
	// 回收站id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 回收站范围类型
	// 枚举值:
	// 	ORG: 企业
	// 	APP: 应用
	// 	SPACE: 空间
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 回收站范围id
	// 根据recycleBinScope传入对应的企业、应用、空间ID
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 用户id
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
	// 回收项信息
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
	// 原文件(夹)id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// 回收项id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 删除时间
	OperatorTime *string `json:"operatorTime,omitempty" xml:"operatorTime,omitempty"`
	// 原文件(夹)名称
	OriginalName *string `json:"originalName,omitempty" xml:"originalName,omitempty"`
	// 原文件(夹)路径
	OriginalPath *string `json:"originalPath,omitempty" xml:"originalPath,omitempty"`
	// 原文件(夹)大小
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 原文件(夹)所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 类型，目录或文件
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRecycleItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 用户id
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
	// 空间详情
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
	// 空间能力项
	Capabilities *GetSpaceResponseBodySpaceCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// 空间归属企业的id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 空间id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 所有者id, 根据ownerType定义, 确定值的所属类型
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// owner类型
	// 枚举值:
	// 	USER: 用户类型
	// 	APP: App类型
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// 总容量
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 业务场景，可以自定义，表示多个不同空间的聚合，可以提供对特定场景做能力设计、容量管理，如根据场景来做搜索或查询。
	// 创建空间时，不指定scene, 默认值是default
	// 默认值:
	// 	default
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 关联业务id, 配合scene一起使用。创建空间时，不指定sceneId， 默认值是0
	// 默认值:
	// 	0
	SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
	// 空间状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETE: 已删除
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 已使用容量
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s GetSpaceResponseBodySpace) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBodySpace) GoString() string {
	return s.String()
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
	// 是否进最近使用, 默认不支持
	// 默认值:
	// 	false
	CanRecordRecentFile *bool `json:"canRecordRecentFile,omitempty" xml:"canRecordRecentFile,omitempty"`
	// 是否支持重命名空间名称, 默认不支持
	// 默认值:
	// 	false
	CanRename *bool `json:"canRename,omitempty" xml:"canRename,omitempty"`
	// 是否支持搜索, 默认不支持
	// 默认值:
	// 	false
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

type GetSpaceResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSpaceResponse) SetBody(v *GetSpaceResponseBody) *GetSpaceResponse {
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
	// 分页大小
	// 默认值:
	// 	50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标, 首次拉取不用传
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 排序规则, 升降或降序
	// 枚举值:
	// 	ASC: 升序
	// 	DESC: 降序
	// 默认值:
	// 	DESC
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// 排序字段
	// 枚举值:
	// 	NAME: 名称
	// 	SIZE: 大小
	// 	MODIFIED_TIME: 最后修改时间
	// 	CREATE_TIME: 创建时间
	// 默认值:
	// 	MODIFIED_TIME
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// 父目录id, 根目录id值为0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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

type ListDentriesResponseBody struct {
	// 文件列表
	Dentries []*ListDentriesResponseBodyDentries `json:"dentries,omitempty" xml:"dentries,omitempty" type:"Repeated"`
	// 分页游标
	// 不为空表示有更多数据
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 后缀
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父目录id, 根目录id值为0
	// 空值代表根目录的parentId不存在
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 存储分区，目前包括公有云OSS存储分区和专属Mini OSS存储分区
	// 枚举值:
	// 	PUBLIC_OSS_PARTITION: 公有云OSS存储分区
	// 	MINI_OSS_PARTITION: 专属Mini OSS存储分区
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// 路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 属性
	Properties *ListDentriesResponseBodyDentriesProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// 大小, 单位:Byte
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETED: 已删除
	// 	EXPIRED: 已过期
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 驱动类型
	// 枚举值:
	// 	DINGTALK: 钉钉统一存储驱动
	// 	ALIDOC: 钉钉文档存储驱动
	// 	SHANJI: 闪记存储驱动
	// 	UNKNOWN: 未知驱动
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// 类型，目录或文件
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// uuid，如移动文件，此字段不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListDentriesResponseBodyDentries) String() string {
	return tea.Prettify(s)
}

func (s ListDentriesResponseBodyDentries) GoString() string {
	return s.String()
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
	// 文件是否只读
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

type ListDentriesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 历史版本分页大小，默认100
	// 默认值:
	// 	100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 下一页的游标位置
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 用户id
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
	// 文件版本列表
	Dentries []*ListDentryVersionsResponseBodyDentries `json:"dentries,omitempty" xml:"dentries,omitempty" type:"Repeated"`
	// 分页游标
	// 不为空表示有更多数据
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 后缀
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父目录id, 根目录id值为0
	// 空值代表根目录的parentId不存在
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 存储分区，目前包括公有云OSS存储分区和专属Mini OSS存储分区
	// 枚举值:
	// 	PUBLIC_OSS_PARTITION: 公有云OSS存储分区
	// 	MINI_OSS_PARTITION: 专属Mini OSS存储分区
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// 路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 属性
	Properties *ListDentryVersionsResponseBodyDentriesProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// 大小, 单位:Byte
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETED: 已删除
	// 	EXPIRED: 已过期
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 驱动类型
	// 枚举值:
	// 	DINGTALK: 钉钉统一存储驱动
	// 	ALIDOC: 钉钉文档存储驱动
	// 	SHANJI: 闪记存储驱动
	// 	UNKNOWN: 未知驱动
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// 类型，目录或文件
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// uuid，如移动文件，此字段不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListDentryVersionsResponseBodyDentries) String() string {
	return tea.Prettify(s)
}

func (s ListDentryVersionsResponseBodyDentries) GoString() string {
	return s.String()
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
	// 文件是否只读
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDentryVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 可选参数
	Option *ListPermissionsRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 用户id
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
	// 角色过滤列表
	FilterRoleIds []*string `json:"filterRoleIds,omitempty" xml:"filterRoleIds,omitempty" type:"Repeated"`
	// 分页大小
	// 默认值:
	// 	50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标
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
	// 分页游标, nextToken不为空表示有更多数据
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 权限列表分页数据
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 文件id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// 有效时间
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 权限成员
	Member *ListPermissionsResponseBodyPermissionsMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 权限角色
	Role *ListPermissionsResponseBodyPermissionsRole `json:"role,omitempty" xml:"role,omitempty" type:"Struct"`
	// 空间id
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

func (s *ListPermissionsResponseBodyPermissions) SetRole(v *ListPermissionsResponseBodyPermissionsRole) *ListPermissionsResponseBodyPermissions {
	s.Role = v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetSpaceId(v string) *ListPermissionsResponseBodyPermissions {
	s.SpaceId = &v
	return s
}

type ListPermissionsResponseBodyPermissionsMember struct {
	// 权限归属的企业
	// 如果存在企业id, 对应member离职的时候会自动清理权限
	// 如果memberType是dept类型，必须要有企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 权限成员id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 权限成员类型
	// 枚举值:
	// 	ORG: 企业
	// 	DEPT: 部门
	// 	TAG: 自定义tag
	// 	CONVERSATION: 会话
	// 	GG: 通用组
	// 	USER: 用户
	// 	ALL_USERS: 所有用户
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
	// 角色id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 角色名称
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 分页大小, 不保证全量返回
	// 默认值:
	// 	50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标，首次拉取nextToken传空
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 用户id
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
	// 分页游标
	// 不为空表示有更多数据
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 回收项列表
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
	// 原文件(夹)id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// 回收项id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 删除时间
	OperatorTime *string `json:"operatorTime,omitempty" xml:"operatorTime,omitempty"`
	// 原文件(夹)名称
	OriginalName *string `json:"originalName,omitempty" xml:"originalName,omitempty"`
	// 原文件(夹)路径
	OriginalPath *string `json:"originalPath,omitempty" xml:"originalPath,omitempty"`
	// 原文件(夹)大小
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 原文件(夹)所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 类型，目录或文件
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRecycleItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRecycleItemsResponse) SetBody(v *ListRecycleItemsResponseBody) *ListRecycleItemsResponse {
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
	// 可选参数
	Option *MoveDentryRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 目标文件夹ID
	TargetFolderId *string `json:"targetFolderId,omitempty" xml:"targetFolderId,omitempty"`
	// 目标文件(夹)空间id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// 用户id
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
	// 文件(夹)名称冲突策略
	// 枚举值:
	// 	AUTO_RENAME: 自动重命名
	// 	OVERWRITE: 覆盖
	// 	RETURN_DENTRY_IF_EXISTS: 返回已存在文件
	// 	RETURN_ERROR_IF_EXISTS: 文件已存在时报错
	// 默认值:
	// 	AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
	// 移动后，是否保留权限
	// 默认值:
	// 	false
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
	// 是否是异步任务
	// 如果操作对象有子节点，则会异步处理
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// 文件信息
	Dentry *MoveDentryResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
	// 任务id，用于查询任务执行状态; 查询接口开发中
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 后缀
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父目录id, 根目录id值为0
	// 空值代表根目录的parentId不存在
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 存储分区，目前包括公有云OSS存储分区和专属Mini OSS存储分区
	// 枚举值:
	// 	PUBLIC_OSS_PARTITION: 公有云OSS存储分区
	// 	MINI_OSS_PARTITION: 专属Mini OSS存储分区
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// 路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 属性
	Properties *MoveDentryResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// 大小, 单位:Byte
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETED: 已删除
	// 	EXPIRED: 已过期
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 驱动类型
	// 枚举值:
	// 	DINGTALK: 钉钉统一存储驱动
	// 	ALIDOC: 钉钉文档存储驱动
	// 	SHANJI: 闪记存储驱动
	// 	UNKNOWN: 未知驱动
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// 类型，目录或文件
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// uuid，如移动文件，此字段不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s MoveDentryResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s MoveDentryResponseBodyDentry) GoString() string {
	return s.String()
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
	// 文件是否只读
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveDentryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *MoveDentryResponse) SetBody(v *MoveDentryResponseBody) *MoveDentryResponse {
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
	// 名称(文件名+后缀), 规则：
	// 1. 头尾不能包含空格，否则会自动去除
	// 2. 不能包含特殊字符，包括：制表符、*、"、<、>、|
	// 3. 不能以"."结尾
	NewName *string `json:"newName,omitempty" xml:"newName,omitempty"`
	// 用户id
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
	// 文件信息
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 后缀
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父目录id, 根目录id值为0
	// 空值代表根目录的parentId不存在
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 存储分区，目前包括公有云OSS存储分区和专属Mini OSS存储分区
	// 枚举值:
	// 	PUBLIC_OSS_PARTITION: 公有云OSS存储分区
	// 	MINI_OSS_PARTITION: 专属Mini OSS存储分区
	PartitionType *string `json:"partitionType,omitempty" xml:"partitionType,omitempty"`
	// 路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 属性
	Properties *RenameDentryResponseBodyDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// 大小, 单位:Byte
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETED: 已删除
	// 	EXPIRED: 已过期
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 驱动类型
	// 枚举值:
	// 	DINGTALK: 钉钉统一存储驱动
	// 	ALIDOC: 钉钉文档存储驱动
	// 	SHANJI: 闪记存储驱动
	// 	UNKNOWN: 未知驱动
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// 类型，目录或文件
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// uuid，如移动文件，此字段不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s RenameDentryResponseBodyDentry) String() string {
	return tea.Prettify(s)
}

func (s RenameDentryResponseBodyDentry) GoString() string {
	return s.String()
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
	// 文件是否只读
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenameDentryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 可选参数
	Option *RestoreRecycleItemRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 用户id
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
	// 文件名称冲突策略
	// 还原时原路径可能已经存在同名的文件
	// 枚举值:
	// 	AUTO_RENAME: 自动重命名
	// 	OVERWRITE: 覆盖
	// 	RETURN_DENTRY_IF_EXISTS: 返回已存在文件
	// 	RETURN_ERROR_IF_EXISTS: 文件已存在时报错
	// 默认值:
	// 	AUTO_RENAME
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
	// 本次操作是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RestoreRecycleItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreRecycleItemResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreRecycleItemResponseBody) SetSuccess(v bool) *RestoreRecycleItemResponseBody {
	s.Success = &v
	return s
}

type RestoreRecycleItemResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestoreRecycleItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RestoreRecycleItemResponse) SetBody(v *RestoreRecycleItemResponseBody) *RestoreRecycleItemResponse {
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
	// 用户id
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
	// 本次操作是否成功
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RevertDentryVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RevertDentryVersionResponse) SetBody(v *RevertDentryVersionResponseBody) *RevertDentryVersionResponse {
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
	// App属性列表 属性不存在时则新增，存在则覆盖原值
	AppProperties []*UpdateDentryAppPropertiesRequestAppProperties `json:"appProperties,omitempty" xml:"appProperties,omitempty" type:"Repeated"`
	// 用户id
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
	// 属性名称 该属性名称在当前app下需要保证唯一，不同app间同名属性互不影响
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 属性值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 属性可见范围
	// 枚举值:
	// 	PUBLIC: 该属性所有App可见
	// 	PRIVATE: 该属性仅其归属App可见
	// 默认值:
	// 	PRIVATE
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
	// 本次操作是否成功
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDentryAppPropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 权限成员列表
	Members []*UpdatePermissionRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 可选参数
	Option *UpdatePermissionRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 权限角色id
	RoleId *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 用户id
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
	// 权限归属的企业
	// 如果存在企业id, 对应member离职的时候会自动清理权限
	// 如果memberType是dept类型，必须要有企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 权限成员id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 权限成员类型
	// 枚举值:
	// 	ORG: 企业
	// 	DEPT: 部门
	// 	TAG: 自定义tag
	// 	CONVERSATION: 会话
	// 	GG: 通用组
	// 	USER: 用户
	// 	ALL_USERS: 所有用户
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
	// 有效时间(秒)
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
	// 本次操作是否成功
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

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

func (client *Client) AddFolderWithOptions(spaceId *string, parentId *string, request *AddFolderRequest, headers *AddFolderHeaders, runtime *util.RuntimeOptions) (_result *AddFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	parentId = openapiutil.GetEncodeParam(parentId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
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
	_result = &AddFolderResponse{}
	_body, _err := client.DoROARequest(tea.String("AddFolder"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(parentId)+"/folders"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) AddPermissionWithOptions(spaceId *string, dentryId *string, request *AddPermissionRequest, headers *AddPermissionHeaders, runtime *util.RuntimeOptions) (_result *AddPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
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
	_result = &AddPermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("AddPermission"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/permissions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
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
	_result = &AddSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("AddSpace"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ClearRecycleBinWithOptions(recycleBinId *string, request *ClearRecycleBinRequest, headers *ClearRecycleBinHeaders, runtime *util.RuntimeOptions) (_result *ClearRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	recycleBinId = openapiutil.GetEncodeParam(recycleBinId)
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
	_result = &ClearRecycleBinResponse{}
	_body, _err := client.DoROARequest(tea.String("ClearRecycleBin"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/recycleBins/"+tea.StringValue(recycleBinId)+"/clear"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CommitFileWithOptions(spaceId *string, request *CommitFileRequest, headers *CommitFileHeaders, runtime *util.RuntimeOptions) (_result *CommitFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
		body["option"] = request.Option
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
	_result = &CommitFileResponse{}
	_body, _err := client.DoROARequest(tea.String("CommitFile"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/files/commit"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CopyDentryWithOptions(spaceId *string, dentryId *string, request *CopyDentryRequest, headers *CopyDentryHeaders, runtime *util.RuntimeOptions) (_result *CopyDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
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
	_result = &CopyDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("CopyDentry"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/copy"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteDentryWithOptions(spaceId *string, dentryId *string, request *DeleteDentryRequest, headers *DeleteDentryHeaders, runtime *util.RuntimeOptions) (_result *DeleteDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &DeleteDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteDentry"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteDentryAppPropertiesWithOptions(spaceId *string, dentryId *string, request *DeleteDentryAppPropertiesRequest, headers *DeleteDentryAppPropertiesHeaders, runtime *util.RuntimeOptions) (_result *DeleteDentryAppPropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &DeleteDentryAppPropertiesResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteDentryAppProperties"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/appProperties/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeletePermissionWithOptions(spaceId *string, dentryId *string, request *DeletePermissionRequest, headers *DeletePermissionHeaders, runtime *util.RuntimeOptions) (_result *DeletePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &DeletePermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("DeletePermission"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/permissions/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteRecycleItemWithOptions(recycleBinId *string, recycleItemId *string, request *DeleteRecycleItemRequest, headers *DeleteRecycleItemHeaders, runtime *util.RuntimeOptions) (_result *DeleteRecycleItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	recycleBinId = openapiutil.GetEncodeParam(recycleBinId)
	recycleItemId = openapiutil.GetEncodeParam(recycleItemId)
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
	_result = &DeleteRecycleItemResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRecycleItem"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/storage/recycleBins/"+tea.StringValue(recycleBinId)+"/recycleItems/"+tea.StringValue(recycleItemId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteRecycleItemsWithOptions(recycleBinId *string, request *DeleteRecycleItemsRequest, headers *DeleteRecycleItemsHeaders, runtime *util.RuntimeOptions) (_result *DeleteRecycleItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	recycleBinId = openapiutil.GetEncodeParam(recycleBinId)
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
	_result = &DeleteRecycleItemsResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRecycleItems"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/recycleBins/"+tea.StringValue(recycleBinId)+"/recycleItems/batchRemove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetCurrentAppResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCurrentApp"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/currentApps/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetDentryWithOptions(spaceId *string, dentryId *string, request *GetDentryRequest, headers *GetDentryHeaders, runtime *util.RuntimeOptions) (_result *GetDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &GetDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDentry"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetFileDownloadInfoWithOptions(spaceId *string, dentryId *string, request *GetFileDownloadInfoRequest, headers *GetFileDownloadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileDownloadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
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
	_result = &GetFileDownloadInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFileDownloadInfo"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/downloadInfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetFileUploadInfoWithOptions(spaceId *string, request *GetFileUploadInfoRequest, headers *GetFileUploadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Multipart)) {
		body["multipart"] = request.Multipart
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
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
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFileUploadInfo"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/files/uploadInfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetRecycleBinResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRecycleBin"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/storage/recycleBins"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetRecycleItemWithOptions(recycleBinId *string, recycleItemId *string, request *GetRecycleItemRequest, headers *GetRecycleItemHeaders, runtime *util.RuntimeOptions) (_result *GetRecycleItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	recycleBinId = openapiutil.GetEncodeParam(recycleBinId)
	recycleItemId = openapiutil.GetEncodeParam(recycleItemId)
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
	_result = &GetRecycleItemResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRecycleItem"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/storage/recycleBins/"+tea.StringValue(recycleBinId)+"/recycleItems/"+tea.StringValue(recycleItemId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetSpaceWithOptions(spaceId *string, request *GetSpaceRequest, headers *GetSpaceHeaders, runtime *util.RuntimeOptions) (_result *GetSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
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
	_result = &GetSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSpace"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListDentriesWithOptions(spaceId *string, request *ListDentriesRequest, headers *ListDentriesHeaders, runtime *util.RuntimeOptions) (_result *ListDentriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
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
	_result = &ListDentriesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDentries"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListDentryVersionsWithOptions(spaceId *string, dentryId *string, request *ListDentryVersionsRequest, headers *ListDentryVersionsHeaders, runtime *util.RuntimeOptions) (_result *ListDentryVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &ListDentryVersionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDentryVersions"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/versions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListPermissionsWithOptions(spaceId *string, dentryId *string, request *ListPermissionsRequest, headers *ListPermissionsHeaders, runtime *util.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
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
	_result = &ListPermissionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListPermissions"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/permissions/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListRecycleItemsWithOptions(recycleBinId *string, request *ListRecycleItemsRequest, headers *ListRecycleItemsHeaders, runtime *util.RuntimeOptions) (_result *ListRecycleItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	recycleBinId = openapiutil.GetEncodeParam(recycleBinId)
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
	_result = &ListRecycleItemsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRecycleItems"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/storage/recycleBins/"+tea.StringValue(recycleBinId)+"/recycleItems"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) MoveDentryWithOptions(spaceId *string, dentryId *string, request *MoveDentryRequest, headers *MoveDentryHeaders, runtime *util.RuntimeOptions) (_result *MoveDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
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
	_result = &MoveDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("MoveDentry"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/move"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) RenameDentryWithOptions(spaceId *string, dentryId *string, request *RenameDentryRequest, headers *RenameDentryHeaders, runtime *util.RuntimeOptions) (_result *RenameDentryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &RenameDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("RenameDentry"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/rename"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) RestoreRecycleItemWithOptions(recycleBinId *string, recycleItemId *string, request *RestoreRecycleItemRequest, headers *RestoreRecycleItemHeaders, runtime *util.RuntimeOptions) (_result *RestoreRecycleItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	recycleBinId = openapiutil.GetEncodeParam(recycleBinId)
	recycleItemId = openapiutil.GetEncodeParam(recycleItemId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
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
	_result = &RestoreRecycleItemResponse{}
	_body, _err := client.DoROARequest(tea.String("RestoreRecycleItem"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/recycleBins/"+tea.StringValue(recycleBinId)+"/recycleItems/"+tea.StringValue(recycleItemId)+"/restore"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) RevertDentryVersionWithOptions(spaceId *string, dentryId *string, version *string, request *RevertDentryVersionRequest, headers *RevertDentryVersionHeaders, runtime *util.RuntimeOptions) (_result *RevertDentryVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
	version = openapiutil.GetEncodeParam(version)
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
	_result = &RevertDentryVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("RevertDentryVersion"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/versions/"+tea.StringValue(version)+"/revert"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateDentryAppPropertiesWithOptions(spaceId *string, dentryId *string, request *UpdateDentryAppPropertiesRequest, headers *UpdateDentryAppPropertiesHeaders, runtime *util.RuntimeOptions) (_result *UpdateDentryAppPropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
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
	_result = &UpdateDentryAppPropertiesResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateDentryAppProperties"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/appProperties"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdatePermissionWithOptions(spaceId *string, dentryId *string, request *UpdatePermissionRequest, headers *UpdatePermissionHeaders, runtime *util.RuntimeOptions) (_result *UpdatePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	spaceId = openapiutil.GetEncodeParam(spaceId)
	dentryId = openapiutil.GetEncodeParam(dentryId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Option))) {
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
	_result = &UpdatePermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdatePermission"), tea.String("storage_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/storage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/permissions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
