// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package sns_storage_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// 文件(夹)id列表
	// 最大size:
	// 	30
	DentryIds []*string `json:"dentryIds,omitempty" xml:"dentryIds,omitempty" type:"Repeated"`
	// 可选参数
	Option *GetDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 用户id
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
	// 通过指定应用id, 返回对应的可见属性，即dentry.appProperties，
	// 默认都会返回当前应用的属性，
	// 如不指定appIds, 则默认返回当前应用的appProperties
	// 最大size:
	// 	20
	AppIdsForAppProperties []*string `json:"appIdsForAppProperties,omitempty" xml:"appIdsForAppProperties,omitempty" type:"Repeated"`
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

type GetDentriesResponseBody struct {
	// 批量获取文件(夹)信息结果列表
	// 最大size:
	// 	30
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
	// 文件(夹)信息
	Dentry *GetDentriesResponseBodyResultItemsDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
	// 文件(夹)id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// 错误原因
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 文件(夹)空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 是否成功
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
	// 在特定应用上的属性。key是微应用Id, value是属性列表。
	// 可以通过修改DentryAppProperty里的scope来设置属性的可见性
	// 最大size:
	// 	10
	AppProperties map[string][]*ResultItemsDentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
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
	Properties *GetDentriesResponseBodyResultItemsDentryProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
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
	// 缩略图信息
	Thumbnail *GetDentriesResponseBodyResultItemsDentryThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
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
	// 文件是否只读
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
	// 缩略图高度
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// 缩略图url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 缩略图宽度
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 可选参数
	Option *GetDentryRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 用户id
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
	// 通过指定应用id, 返回对应的可见属性，即dentry.appProperties，
	// 默认都会返回当前应用的属性，
	// 如不指定appIds, 则默认返回当前应用的appProperties
	// 最大size:
	// 	20
	AppIdsForAppProperties []*string `json:"appIdsForAppProperties,omitempty" xml:"appIdsForAppProperties,omitempty" type:"Repeated"`
	// 是否获取文件缩略图临时链接
	// 注: 按需获取, 会影响接口耗时
	// 默认值:
	// 	false
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
	// 在特定应用上的属性。key是微应用Id, value是属性列表。
	// 可以通过修改DentryAppProperty里的scope来设置属性的可见性
	// 最大size:
	// 	10
	AppProperties map[string][]*DentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
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
	// 缩略图信息
	Thumbnail *GetDentryResponseBodyDentryThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
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

type GetDentryResponseBodyDentryThumbnail struct {
	// 缩略图高度
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// 缩略图url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 缩略图宽度
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
	// 文件id列表
	// 最大size:
	// 	30
	DentryIds []*string `json:"dentryIds,omitempty" xml:"dentryIds,omitempty" type:"Repeated"`
	// 用户id
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
	// 缩略图获取结果列表
	// 最大size:
	// 	30
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
	// 源文件(夹)id
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// 错误原因
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 源文件(夹)空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 是否成功获取到缩略图
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 缩略图信息
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
	// 缩略图高度
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// 缩略图url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 缩略图宽度
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDentryThumbnailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 优先使用内网传输
	// 前提: 配置了专属存储内网传输
	// 默认值:
	// 	true
	PreferIntranet *bool `json:"preferIntranet,omitempty" xml:"preferIntranet,omitempty"`
	// 历史版本号
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
	// 最大size:
	// 	20
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// 内网URL, 在网络连通的情况下，使用内网URL可加速服务器间上传
	// 最大size:
	// 	10
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
	// 最大size:
	// 	10
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
	// 会话id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceRequest) SetOpenConversationId(v string) *GetSpaceRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetSpaceRequest) SetUnionId(v string) *GetSpaceRequest {
	s.UnionId = &v
	return s
}

type GetSpaceResponseBody struct {
	// IM会话存储空间信息
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
	// 空间归属企业的id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetSpaceResponseBodySpace) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceResponseBodySpace) GoString() string {
	return s.String()
}

func (s *GetSpaceResponseBodySpace) SetCorpId(v string) *GetSpaceResponseBodySpace {
	s.CorpId = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetCreateTime(v string) *GetSpaceResponseBodySpace {
	s.CreateTime = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetModifiedTime(v string) *GetSpaceResponseBodySpace {
	s.ModifiedTime = &v
	return s
}

func (s *GetSpaceResponseBodySpace) SetSpaceId(v string) *GetSpaceResponseBodySpace {
	s.SpaceId = &v
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
	// 可选参数
	Option *ListAllDentriesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 用户id
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
	// 分页大小
	// 默认值:
	// 	50
	// 最大值:
	// 	50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标, 首次拉取不用传
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 排序规则, 升降或降序
	// 目前仅支持按修改时间排序,
	// 如果是升序排序, 分页拉取过程中, 如果文件发生变化, 可能拉取到重复数据
	// 如果是降序排序, 分页拉取过程中, 如果文件发生变化, 可能会丢失数据
	// 枚举值:
	// 	ASC: 升序
	// 	DESC: 降序
	// 默认值:
	// 	ASC
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// 是否获取文件缩略图临时链接
	// 注: 按需获取, 会影响接口耗时
	// 默认值:
	// 	false
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
	// 文件列表
	// 最大size:
	// 	50
	Dentries []*ListAllDentriesResponseBodyDentries `json:"dentries,omitempty" xml:"dentries,omitempty" type:"Repeated"`
	// 分页游标
	// 不为空表示有更多数据
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
	// 在特定应用上的属性。key是微应用Id, value是属性列表。
	// 可以通过修改DentryAppProperty里的scope来设置属性的可见性
	// 最大size:
	// 	10
	AppProperties map[string][]*DentriesAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
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
	Properties *ListAllDentriesResponseBodyDentriesProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
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
	// 缩略图信息
	Thumbnail *ListAllDentriesResponseBodyDentriesThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
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
	// 文件是否只读
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
	// 缩略图高度
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// 缩略图url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 缩略图宽度
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAllDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 分页大小
	// 默认值:
	// 	50
	// 最大值:
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
	// 是否获取文件缩略图临时链接
	// 注: 按需获取, 会影响接口耗时
	// 默认值:
	// 	false
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
	// 文件列表
	// 最大size:
	// 	50
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
	// 在特定应用上的属性。key是微应用Id, value是属性列表。
	// 可以通过修改DentryAppProperty里的scope来设置属性的可见性
	// 最大size:
	// 	10
	AppProperties map[string][]*DentriesAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
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
	// 缩略图信息
	Thumbnail *ListDentriesResponseBodyDentriesThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
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

func (s *ListDentriesResponseBodyDentries) SetAppProperties(v map[string][]*DentriesAppPropertiesValue) *ListDentriesResponseBodyDentries {
	s.AppProperties = v
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

type ListDentriesResponseBodyDentriesThumbnail struct {
	// 缩略图高度
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// 缩略图url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 缩略图宽度
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

type ListExpiredHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListExpiredHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredHeaders) GoString() string {
	return s.String()
}

func (s *ListExpiredHeaders) SetCommonHeaders(v map[string]*string) *ListExpiredHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListExpiredHeaders) SetXAcsDingtalkAccessToken(v string) *ListExpiredHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListExpiredRequest struct {
	// 会话id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 可选参数
	Option *ListExpiredRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListExpiredRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredRequest) GoString() string {
	return s.String()
}

func (s *ListExpiredRequest) SetOpenConversationId(v string) *ListExpiredRequest {
	s.OpenConversationId = &v
	return s
}

func (s *ListExpiredRequest) SetOption(v *ListExpiredRequestOption) *ListExpiredRequest {
	s.Option = v
	return s
}

func (s *ListExpiredRequest) SetUnionId(v string) *ListExpiredRequest {
	s.UnionId = &v
	return s
}

type ListExpiredRequestOption struct {
	// 分页大小
	// 默认值:
	// 	50
	// 最大值:
	// 	50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标, 首次拉取不用传
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListExpiredRequestOption) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredRequestOption) GoString() string {
	return s.String()
}

func (s *ListExpiredRequestOption) SetMaxResults(v int32) *ListExpiredRequestOption {
	s.MaxResults = &v
	return s
}

func (s *ListExpiredRequestOption) SetNextToken(v string) *ListExpiredRequestOption {
	s.NextToken = &v
	return s
}

type ListExpiredResponseBody struct {
	// 过期文件列表
	// 最大size:
	// 	50
	Files []*ListExpiredResponseBodyFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// 分页游标
	// 不为空表示有更多数据
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListExpiredResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredResponseBody) GoString() string {
	return s.String()
}

func (s *ListExpiredResponseBody) SetFiles(v []*ListExpiredResponseBodyFiles) *ListExpiredResponseBody {
	s.Files = v
	return s
}

func (s *ListExpiredResponseBody) SetNextToken(v string) *ListExpiredResponseBody {
	s.NextToken = &v
	return s
}

type ListExpiredResponseBodyFiles struct {
	// 文件所在会话id
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者id
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 文件后缀
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 文件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// 修改者id
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// 文件(夹)名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文件所在的父目录id, 根目录id值为0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 文件大小, 单位:Byte
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 文件所在空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 文件状态
	// 枚举值:
	// 	NORMAL: 正常
	// 	DELETED: 已删除
	// 	EXPIRED: 已过期
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 文件类型：文件、文件夹
	// 枚举值:
	// 	FILE: 文件
	// 	FOLDER: 文件夹
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// uuid，如移动文件，此字段不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 文件版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListExpiredResponseBodyFiles) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *ListExpiredResponseBodyFiles) SetConversationId(v string) *ListExpiredResponseBodyFiles {
	s.ConversationId = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetCreateTime(v string) *ListExpiredResponseBodyFiles {
	s.CreateTime = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetCreatorId(v string) *ListExpiredResponseBodyFiles {
	s.CreatorId = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetExtension(v string) *ListExpiredResponseBodyFiles {
	s.Extension = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetId(v string) *ListExpiredResponseBodyFiles {
	s.Id = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetModifiedTime(v string) *ListExpiredResponseBodyFiles {
	s.ModifiedTime = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetModifierId(v string) *ListExpiredResponseBodyFiles {
	s.ModifierId = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetName(v string) *ListExpiredResponseBodyFiles {
	s.Name = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetParentId(v string) *ListExpiredResponseBodyFiles {
	s.ParentId = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetPath(v string) *ListExpiredResponseBodyFiles {
	s.Path = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetSize(v int64) *ListExpiredResponseBodyFiles {
	s.Size = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetSpaceId(v string) *ListExpiredResponseBodyFiles {
	s.SpaceId = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetStatus(v string) *ListExpiredResponseBodyFiles {
	s.Status = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetType(v string) *ListExpiredResponseBodyFiles {
	s.Type = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetUuid(v string) *ListExpiredResponseBodyFiles {
	s.Uuid = &v
	return s
}

func (s *ListExpiredResponseBodyFiles) SetVersion(v int64) *ListExpiredResponseBodyFiles {
	s.Version = &v
	return s
}

type ListExpiredResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExpiredResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExpiredResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExpiredResponse) GoString() string {
	return s.String()
}

func (s *ListExpiredResponse) SetHeaders(v map[string]*string) *ListExpiredResponse {
	s.Headers = v
	return s
}

func (s *ListExpiredResponse) SetBody(v *ListExpiredResponseBody) *ListExpiredResponse {
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
	// 订阅范围
	// 枚举值:
	// 	SPACE: 空间
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 订阅范围对应的id
	// ORG时，对应的是企业id
	// APP时，对应的是应用id
	// SPACE时，对应的是空间id
	// 枚举值:
	// 	SPACE: 空间
	ScopeId *string `json:"scopeId,omitempty" xml:"scopeId,omitempty"`
	// 用户id
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
	// 本次操作是否成功
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubscribeEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 订阅范围
	// 枚举值:
	// 	SPACE: 空间
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 订阅范围对应的id
	// ORG时，对应的是企业id
	// APP时，对应的是应用id
	// SPACE时，对应的是空间id
	// 枚举值:
	// 	SPACE: 空间
	ScopeId *string `json:"scopeId,omitempty" xml:"scopeId,omitempty"`
	// 用户id
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
	// 本次操作是否成功
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnsubscribeEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UnsubscribeEventResponse) SetBody(v *UnsubscribeEventResponseBody) *UnsubscribeEventResponse {
	s.Body = v
	return s
}

type ResultItemsDentryAppPropertiesValue struct {
	// 属性名称 该属性名称在当前app下需要保证唯一，不同app间同名属性互不影响
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 属性值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 属性可见范围
	// 枚举值:
	// 	PUBLIC: 该属性所有App可见
	// 	PRIVATE: 该属性仅其归属App可见
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

type DentryAppPropertiesValue struct {
	// 属性名称 该属性名称在当前app下需要保证唯一，不同app间同名属性互不影响
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 属性值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 属性可见范围
	// 枚举值:
	// 	PUBLIC: 该属性所有App可见
	// 	PRIVATE: 该属性仅其归属App可见
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

type DentriesAppPropertiesValue struct {
	// 属性名称 该属性名称在当前app下需要保证唯一，不同app间同名属性互不影响
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 属性值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 属性可见范围
	// 枚举值:
	// 	PUBLIC: 该属性所有App可见
	// 	PRIVATE: 该属性仅其归属App可见
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

func (client *Client) GetDentriesWithOptions(spaceId *string, request *GetDentriesRequest, headers *GetDentriesHeaders, runtime *util.RuntimeOptions) (_result *GetDentriesResponse, _err error) {
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
	_result = &GetDentriesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDentries"), tea.String("snsStorage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/snsStorage/spaces/"+tea.StringValue(spaceId)+"/dentries/batchQuery"), tea.String("json"), req, runtime)
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
	_result = &GetDentryResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDentry"), tea.String("snsStorage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/snsStorage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetDentryThumbnailsWithOptions(spaceId *string, request *GetDentryThumbnailsRequest, headers *GetDentryThumbnailsHeaders, runtime *util.RuntimeOptions) (_result *GetDentryThumbnailsResponse, _err error) {
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
	_result = &GetDentryThumbnailsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDentryThumbnails"), tea.String("snsStorage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/snsStorage/spaces/"+tea.StringValue(spaceId)+"/thumbnails/query"), tea.String("json"), req, runtime)
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
	_result = &GetFileDownloadInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFileDownloadInfo"), tea.String("snsStorage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/snsStorage/spaces/"+tea.StringValue(spaceId)+"/dentries/"+tea.StringValue(dentryId)+"/downloadInfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpace(request *GetSpaceRequest) (_result *GetSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSpaceHeaders{}
	_result = &GetSpaceResponse{}
	_body, _err := client.GetSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpaceWithOptions(request *GetSpaceRequest, headers *GetSpaceHeaders, runtime *util.RuntimeOptions) (_result *GetSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &GetSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSpace"), tea.String("snsStorage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/snsStorage/conversations/spaces/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListAllDentriesWithOptions(spaceId *string, request *ListAllDentriesRequest, headers *ListAllDentriesHeaders, runtime *util.RuntimeOptions) (_result *ListAllDentriesResponse, _err error) {
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
	_result = &ListAllDentriesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListAllDentries"), tea.String("snsStorage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/snsStorage/spaces/"+tea.StringValue(spaceId)+"/dentries/listAll"), tea.String("json"), req, runtime)
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
	_result = &ListDentriesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDentries"), tea.String("snsStorage_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/snsStorage/spaces/"+tea.StringValue(spaceId)+"/dentries"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExpired(request *ListExpiredRequest) (_result *ListExpiredResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListExpiredHeaders{}
	_result = &ListExpiredResponse{}
	_body, _err := client.ListExpiredWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExpiredWithOptions(request *ListExpiredRequest, headers *ListExpiredHeaders, runtime *util.RuntimeOptions) (_result *ListExpiredResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &ListExpiredResponse{}
	_body, _err := client.DoROARequest(tea.String("ListExpired"), tea.String("snsStorage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/snsStorage/conversations/expiredFileLists/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &SubscribeEventResponse{}
	_body, _err := client.DoROARequest(tea.String("SubscribeEvent"), tea.String("snsStorage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/snsStorage/events/subscribe"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UnsubscribeEventResponse{}
	_body, _err := client.DoROARequest(tea.String("UnsubscribeEvent"), tea.String("snsStorage_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/snsStorage/events/unsubscribe"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
