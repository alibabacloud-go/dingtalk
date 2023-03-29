// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package storage_2_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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

func (s *CommitFileRequest) SetUploadKey(v string) *CommitFileRequest {
	s.UploadKey = &v
	return s
}

func (s *CommitFileRequest) SetUnionId(v string) *CommitFileRequest {
	s.UnionId = &v
	return s
}

type CommitFileRequestOption struct {
	// 文件在应用上的属性, 一个应用最多只能设置3个属性
	// 最大size:
	// 	3
	AppProperties []*CommitFileRequestOptionAppProperties `json:"appProperties,omitempty" xml:"appProperties,omitempty" type:"Repeated"`
	// 文件名称冲突策略
	// 枚举值:
	// 	AUTO_RENAME: 自动重命名
	// 	OVERWRITE: 覆盖
	// 	RETURN_DENTRY_IF_EXISTS: 返回已存在文件
	// 	RETURN_ERROR_IF_EXISTS: 文件已存在时报错
	// 默认值:
	// 	AUTO_RENAME
	ConflictStrategy *string `json:"conflictStrategy,omitempty" xml:"conflictStrategy,omitempty"`
	// 是否转换成在线文档
	// 默认值:
	// 	false
	ConvertToOnlineDoc *bool `json:"convertToOnlineDoc,omitempty" xml:"convertToOnlineDoc,omitempty"`
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

func (s *CommitFileRequestOption) SetSize(v int64) *CommitFileRequestOption {
	s.Size = &v
	return s
}

type CommitFileRequestOptionAppProperties struct {
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
	// 在特定应用上的属性。key是微应用Id, value是属性列表。
	// 可以通过修改DentryAppProperty里的scope来设置属性的可见性
	// 最大size:
	// 	10
	AppProperties map[string][]*DentryAppPropertiesValue `json:"appProperties,omitempty" xml:"appProperties,omitempty"`
	// 类别, 如图片、视频、音频等
	// 枚举值:
	// 	IMAGE: 图片
	// 	VIDEO: 视频
	// 	AUDIO: 音频
	// 	ARCHIVE: 压缩包
	// 	SHORTCUT: 快捷方式
	// 	DOCUMENT: 文档
	// 	ALI_DOC: 钉钉文档
	// 	OTHER: 其它
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
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
	// 缩略图信息
	Thumbnail *CommitFileResponseBodyDentryThumbnail `json:"thumbnail,omitempty" xml:"thumbnail,omitempty" type:"Struct"`
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

type CommitFileResponseBodyDentryThumbnail struct {
	// 缩略图高度
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// 缩略图url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 缩略图宽度
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
	// 优先使用内网传输
	// 前提: 配置了专属存储内网传输
	// 默认值:
	// 	true
	PreferIntranet *bool `json:"preferIntranet,omitempty" xml:"preferIntranet,omitempty"`
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
	// 文件名称, 文件名称合法性和文件名称冲突校验
	// 规则：
	// 1. 头尾不能包含空格，否则会自动去除
	// 2. 不能包含特殊字符，包括：制表符、*、"、<、>、|
	// 3. 不能以"."结尾
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文件大小, 做容量相关校验。不传则不做校验。
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

func (client *Client) CommitFileWithOptions(parentDentryUuid *string, request *CommitFileRequest, headers *CommitFileHeaders, runtime *util.RuntimeOptions) (_result *CommitFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	parentDentryUuid = openapiutil.GetEncodeParam(parentDentryUuid)
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
	_result = &CommitFileResponse{}
	_body, _err := client.DoROARequest(tea.String("CommitFile"), tea.String("storage_2.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v2.0/storage/spaces/files/"+tea.StringValue(parentDentryUuid)+"/commit"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetFileUploadInfoWithOptions(parentDentryUuid *string, request *GetFileUploadInfoRequest, headers *GetFileUploadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	parentDentryUuid = openapiutil.GetEncodeParam(parentDentryUuid)
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
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFileUploadInfo"), tea.String("storage_2.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v2.0/storage/spaces/files/"+tea.StringValue(parentDentryUuid)+"/uploadInfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
