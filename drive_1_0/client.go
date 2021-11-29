// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package drive_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RecoverRecycleFilesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RecoverRecycleFilesHeaders) String() string {
	return tea.Prettify(s)
}

func (s RecoverRecycleFilesHeaders) GoString() string {
	return s.String()
}

func (s *RecoverRecycleFilesHeaders) SetCommonHeaders(v map[string]*string) *RecoverRecycleFilesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecoverRecycleFilesHeaders) SetXAcsDingtalkAccessToken(v string) *RecoverRecycleFilesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RecoverRecycleFilesRequest struct {
	// 回收站item id列表
	RecycleItemIdList []*int64 `json:"recycleItemIdList,omitempty" xml:"recycleItemIdList,omitempty" type:"Repeated"`
	// 回收站类型
	RecycleType *string `json:"recycleType,omitempty" xml:"recycleType,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RecoverRecycleFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverRecycleFilesRequest) GoString() string {
	return s.String()
}

func (s *RecoverRecycleFilesRequest) SetRecycleItemIdList(v []*int64) *RecoverRecycleFilesRequest {
	s.RecycleItemIdList = v
	return s
}

func (s *RecoverRecycleFilesRequest) SetRecycleType(v string) *RecoverRecycleFilesRequest {
	s.RecycleType = &v
	return s
}

func (s *RecoverRecycleFilesRequest) SetUnionId(v string) *RecoverRecycleFilesRequest {
	s.UnionId = &v
	return s
}

type RecoverRecycleFilesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RecoverRecycleFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverRecycleFilesResponse) GoString() string {
	return s.String()
}

func (s *RecoverRecycleFilesResponse) SetHeaders(v map[string]*string) *RecoverRecycleFilesResponse {
	s.Headers = v
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
	// 空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceRequest) GoString() string {
	return s.String()
}

func (s *AddSpaceRequest) SetName(v string) *AddSpaceRequest {
	s.Name = &v
	return s
}

func (s *AddSpaceRequest) SetUnionId(v string) *AddSpaceRequest {
	s.UnionId = &v
	return s
}

type AddSpaceResponseBody struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 空间名称
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	// 空间类型
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 空间总额度
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 空间已使用额度
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	// 授权模式
	PermissionMode *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
}

func (s AddSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddSpaceResponseBody) SetSpaceId(v string) *AddSpaceResponseBody {
	s.SpaceId = &v
	return s
}

func (s *AddSpaceResponseBody) SetSpaceName(v string) *AddSpaceResponseBody {
	s.SpaceName = &v
	return s
}

func (s *AddSpaceResponseBody) SetSpaceType(v string) *AddSpaceResponseBody {
	s.SpaceType = &v
	return s
}

func (s *AddSpaceResponseBody) SetQuota(v int64) *AddSpaceResponseBody {
	s.Quota = &v
	return s
}

func (s *AddSpaceResponseBody) SetUsedQuota(v int64) *AddSpaceResponseBody {
	s.UsedQuota = &v
	return s
}

func (s *AddSpaceResponseBody) SetPermissionMode(v string) *AddSpaceResponseBody {
	s.PermissionMode = &v
	return s
}

func (s *AddSpaceResponseBody) SetCreateTime(v string) *AddSpaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *AddSpaceResponseBody) SetModifyTime(v string) *AddSpaceResponseBody {
	s.ModifyTime = &v
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

type DeleteRecycleFilesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteRecycleFilesHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleFilesHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRecycleFilesHeaders) SetCommonHeaders(v map[string]*string) *DeleteRecycleFilesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRecycleFilesHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteRecycleFilesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteRecycleFilesRequest struct {
	// 回收站item id列表
	RecycleItemIdList []*int64 `json:"recycleItemIdList,omitempty" xml:"recycleItemIdList,omitempty" type:"Repeated"`
	// 回收站类型
	RecycleType *string `json:"recycleType,omitempty" xml:"recycleType,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteRecycleFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleFilesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecycleFilesRequest) SetRecycleItemIdList(v []*int64) *DeleteRecycleFilesRequest {
	s.RecycleItemIdList = v
	return s
}

func (s *DeleteRecycleFilesRequest) SetRecycleType(v string) *DeleteRecycleFilesRequest {
	s.RecycleType = &v
	return s
}

func (s *DeleteRecycleFilesRequest) SetUnionId(v string) *DeleteRecycleFilesRequest {
	s.UnionId = &v
	return s
}

type DeleteRecycleFilesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteRecycleFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecycleFilesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecycleFilesResponse) SetHeaders(v map[string]*string) *DeleteRecycleFilesResponse {
	s.Headers = v
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
	// 权限角色
	Role    *string                        `json:"role,omitempty" xml:"role,omitempty"`
	Members []*AddPermissionRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPermissionRequest) GoString() string {
	return s.String()
}

func (s *AddPermissionRequest) SetRole(v string) *AddPermissionRequest {
	s.Role = &v
	return s
}

func (s *AddPermissionRequest) SetMembers(v []*AddPermissionRequestMembers) *AddPermissionRequest {
	s.Members = v
	return s
}

func (s *AddPermissionRequest) SetUnionId(v string) *AddPermissionRequest {
	s.UnionId = &v
	return s
}

type AddPermissionRequestMembers struct {
	// 企业corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 成员类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 成员id
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
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

func (s *AddPermissionRequestMembers) SetMemberType(v string) *AddPermissionRequestMembers {
	s.MemberType = &v
	return s
}

func (s *AddPermissionRequestMembers) SetMemberId(v string) *AddPermissionRequestMembers {
	s.MemberId = &v
	return s
}

type AddPermissionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
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

type GetQuotaInfosHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetQuotaInfosHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaInfosHeaders) GoString() string {
	return s.String()
}

func (s *GetQuotaInfosHeaders) SetCommonHeaders(v map[string]*string) *GetQuotaInfosHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetQuotaInfosHeaders) SetXAcsDingtalkAccessToken(v string) *GetQuotaInfosHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetQuotaInfosRequest struct {
	// 容量类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 容量标识符列表
	Identifiers []*string `json:"identifiers,omitempty" xml:"identifiers,omitempty" type:"Repeated"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetQuotaInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaInfosRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaInfosRequest) SetType(v string) *GetQuotaInfosRequest {
	s.Type = &v
	return s
}

func (s *GetQuotaInfosRequest) SetIdentifiers(v []*string) *GetQuotaInfosRequest {
	s.Identifiers = v
	return s
}

func (s *GetQuotaInfosRequest) SetUnionId(v string) *GetQuotaInfosRequest {
	s.UnionId = &v
	return s
}

type GetQuotaInfosResponseBody struct {
	// 容量信息列表
	Quotas []*GetQuotaInfosResponseBodyQuotas `json:"quotas,omitempty" xml:"quotas,omitempty" type:"Repeated"`
}

func (s GetQuotaInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaInfosResponseBody) SetQuotas(v []*GetQuotaInfosResponseBodyQuotas) *GetQuotaInfosResponseBody {
	s.Quotas = v
	return s
}

type GetQuotaInfosResponseBodyQuotas struct {
	// 容量类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 容量标识符
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 总容量
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 已使用容量
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s GetQuotaInfosResponseBodyQuotas) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaInfosResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *GetQuotaInfosResponseBodyQuotas) SetType(v string) *GetQuotaInfosResponseBodyQuotas {
	s.Type = &v
	return s
}

func (s *GetQuotaInfosResponseBodyQuotas) SetIdentifier(v string) *GetQuotaInfosResponseBodyQuotas {
	s.Identifier = &v
	return s
}

func (s *GetQuotaInfosResponseBodyQuotas) SetQuota(v int64) *GetQuotaInfosResponseBodyQuotas {
	s.Quota = &v
	return s
}

func (s *GetQuotaInfosResponseBodyQuotas) SetUsedQuota(v int64) *GetQuotaInfosResponseBodyQuotas {
	s.UsedQuota = &v
	return s
}

type GetQuotaInfosResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQuotaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQuotaInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaInfosResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaInfosResponse) SetHeaders(v map[string]*string) *GetQuotaInfosResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaInfosResponse) SetBody(v *GetQuotaInfosResponseBody) *GetQuotaInfosResponse {
	s.Body = v
	return s
}

type GetFileInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFileInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFileInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFileInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFileInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetFileInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFileInfoRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetFileInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFileInfoRequest) SetUnionId(v string) *GetFileInfoRequest {
	s.UnionId = &v
	return s
}

type GetFileInfoResponseBody struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 父目录id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件id
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件路径
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 文件类型
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 文件内容类型
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改者
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s GetFileInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileInfoResponseBody) SetSpaceId(v string) *GetFileInfoResponseBody {
	s.SpaceId = &v
	return s
}

func (s *GetFileInfoResponseBody) SetParentId(v string) *GetFileInfoResponseBody {
	s.ParentId = &v
	return s
}

func (s *GetFileInfoResponseBody) SetFileId(v string) *GetFileInfoResponseBody {
	s.FileId = &v
	return s
}

func (s *GetFileInfoResponseBody) SetFileName(v string) *GetFileInfoResponseBody {
	s.FileName = &v
	return s
}

func (s *GetFileInfoResponseBody) SetFilePath(v string) *GetFileInfoResponseBody {
	s.FilePath = &v
	return s
}

func (s *GetFileInfoResponseBody) SetFileType(v string) *GetFileInfoResponseBody {
	s.FileType = &v
	return s
}

func (s *GetFileInfoResponseBody) SetContentType(v string) *GetFileInfoResponseBody {
	s.ContentType = &v
	return s
}

func (s *GetFileInfoResponseBody) SetFileExtension(v string) *GetFileInfoResponseBody {
	s.FileExtension = &v
	return s
}

func (s *GetFileInfoResponseBody) SetFileSize(v int64) *GetFileInfoResponseBody {
	s.FileSize = &v
	return s
}

func (s *GetFileInfoResponseBody) SetCreateTime(v string) *GetFileInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetFileInfoResponseBody) SetModifyTime(v string) *GetFileInfoResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetFileInfoResponseBody) SetCreator(v string) *GetFileInfoResponseBody {
	s.Creator = &v
	return s
}

func (s *GetFileInfoResponseBody) SetModifier(v string) *GetFileInfoResponseBody {
	s.Modifier = &v
	return s
}

type GetFileInfoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFileInfoResponse) SetHeaders(v map[string]*string) *GetFileInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFileInfoResponse) SetBody(v *GetFileInfoResponseBody) *GetFileInfoResponse {
	s.Body = v
	return s
}

type ListRecycleFilesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListRecycleFilesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleFilesHeaders) GoString() string {
	return s.String()
}

func (s *ListRecycleFilesHeaders) SetCommonHeaders(v map[string]*string) *ListRecycleFilesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRecycleFilesHeaders) SetXAcsDingtalkAccessToken(v string) *ListRecycleFilesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListRecycleFilesRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 回收站类型
	RecycleType *string `json:"recycleType,omitempty" xml:"recycleType,omitempty"`
	// 分页加载更多锚点
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页长度
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 文件排序类型
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
}

func (s ListRecycleFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleFilesRequest) GoString() string {
	return s.String()
}

func (s *ListRecycleFilesRequest) SetUnionId(v string) *ListRecycleFilesRequest {
	s.UnionId = &v
	return s
}

func (s *ListRecycleFilesRequest) SetRecycleType(v string) *ListRecycleFilesRequest {
	s.RecycleType = &v
	return s
}

func (s *ListRecycleFilesRequest) SetNextToken(v string) *ListRecycleFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecycleFilesRequest) SetMaxResults(v int32) *ListRecycleFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecycleFilesRequest) SetOrderType(v string) *ListRecycleFilesRequest {
	s.OrderType = &v
	return s
}

type ListRecycleFilesResponseBody struct {
	// 回收站文件列表
	RecycleItems []*ListRecycleFilesResponseBodyRecycleItems `json:"recycleItems,omitempty" xml:"recycleItems,omitempty" type:"Repeated"`
	// 加载更多锚点, nextToken不为空表示有更多数据
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListRecycleFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecycleFilesResponseBody) SetRecycleItems(v []*ListRecycleFilesResponseBodyRecycleItems) *ListRecycleFilesResponseBody {
	s.RecycleItems = v
	return s
}

func (s *ListRecycleFilesResponseBody) SetNextToken(v string) *ListRecycleFilesResponseBody {
	s.NextToken = &v
	return s
}

type ListRecycleFilesResponseBodyRecycleItems struct {
	// 回收站item id
	RecycleItemId *string `json:"recycleItemId,omitempty" xml:"recycleItemId,omitempty"`
	// 删除员工工号
	DeleteStaffId *string `json:"deleteStaffId,omitempty" xml:"deleteStaffId,omitempty"`
	// 删除时间
	DeleteTime *string `json:"deleteTime,omitempty" xml:"deleteTime,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 文件类型
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 文件内容类型
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件路径
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
}

func (s ListRecycleFilesResponseBodyRecycleItems) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleFilesResponseBodyRecycleItems) GoString() string {
	return s.String()
}

func (s *ListRecycleFilesResponseBodyRecycleItems) SetRecycleItemId(v string) *ListRecycleFilesResponseBodyRecycleItems {
	s.RecycleItemId = &v
	return s
}

func (s *ListRecycleFilesResponseBodyRecycleItems) SetDeleteStaffId(v string) *ListRecycleFilesResponseBodyRecycleItems {
	s.DeleteStaffId = &v
	return s
}

func (s *ListRecycleFilesResponseBodyRecycleItems) SetDeleteTime(v string) *ListRecycleFilesResponseBodyRecycleItems {
	s.DeleteTime = &v
	return s
}

func (s *ListRecycleFilesResponseBodyRecycleItems) SetFileSize(v int64) *ListRecycleFilesResponseBodyRecycleItems {
	s.FileSize = &v
	return s
}

func (s *ListRecycleFilesResponseBodyRecycleItems) SetFileType(v string) *ListRecycleFilesResponseBodyRecycleItems {
	s.FileType = &v
	return s
}

func (s *ListRecycleFilesResponseBodyRecycleItems) SetContentType(v string) *ListRecycleFilesResponseBodyRecycleItems {
	s.ContentType = &v
	return s
}

func (s *ListRecycleFilesResponseBodyRecycleItems) SetFileName(v string) *ListRecycleFilesResponseBodyRecycleItems {
	s.FileName = &v
	return s
}

func (s *ListRecycleFilesResponseBodyRecycleItems) SetFilePath(v string) *ListRecycleFilesResponseBodyRecycleItems {
	s.FilePath = &v
	return s
}

type ListRecycleFilesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRecycleFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecycleFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleFilesResponse) GoString() string {
	return s.String()
}

func (s *ListRecycleFilesResponse) SetHeaders(v map[string]*string) *ListRecycleFilesResponse {
	s.Headers = v
	return s
}

func (s *ListRecycleFilesResponse) SetBody(v *ListRecycleFilesResponseBody) *ListRecycleFilesResponse {
	s.Body = v
	return s
}

type DeleteFilesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteFilesHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilesHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFilesHeaders) SetCommonHeaders(v map[string]*string) *DeleteFilesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFilesHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteFilesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteFilesRequest struct {
	// 文件id列表
	FileIds []*string `json:"fileIds,omitempty" xml:"fileIds,omitempty" type:"Repeated"`
	// 删除策略
	DeletePolicy *string `json:"deletePolicy,omitempty" xml:"deletePolicy,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilesRequest) SetFileIds(v []*string) *DeleteFilesRequest {
	s.FileIds = v
	return s
}

func (s *DeleteFilesRequest) SetDeletePolicy(v string) *DeleteFilesRequest {
	s.DeletePolicy = &v
	return s
}

func (s *DeleteFilesRequest) SetUnionId(v string) *DeleteFilesRequest {
	s.UnionId = &v
	return s
}

type DeleteFilesResponseBody struct {
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 异步任务id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFilesResponseBody) SetSuccess(v bool) *DeleteFilesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFilesResponseBody) SetTaskId(v string) *DeleteFilesResponseBody {
	s.TaskId = &v
	return s
}

type DeleteFilesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFilesResponse) SetHeaders(v map[string]*string) *DeleteFilesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFilesResponse) SetBody(v *DeleteFilesResponseBody) *DeleteFilesResponse {
	s.Body = v
	return s
}

type ManagementBuyQuotaHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ManagementBuyQuotaHeaders) String() string {
	return tea.Prettify(s)
}

func (s ManagementBuyQuotaHeaders) GoString() string {
	return s.String()
}

func (s *ManagementBuyQuotaHeaders) SetCommonHeaders(v map[string]*string) *ManagementBuyQuotaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ManagementBuyQuotaHeaders) SetXAcsDingtalkAccessToken(v string) *ManagementBuyQuotaHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ManagementBuyQuotaRequest struct {
	// token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 订单
	Order *ManagementBuyQuotaRequestOrder `json:"order,omitempty" xml:"order,omitempty" type:"Struct"`
}

func (s ManagementBuyQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s ManagementBuyQuotaRequest) GoString() string {
	return s.String()
}

func (s *ManagementBuyQuotaRequest) SetToken(v string) *ManagementBuyQuotaRequest {
	s.Token = &v
	return s
}

func (s *ManagementBuyQuotaRequest) SetUnionId(v string) *ManagementBuyQuotaRequest {
	s.UnionId = &v
	return s
}

func (s *ManagementBuyQuotaRequest) SetOrder(v *ManagementBuyQuotaRequestOrder) *ManagementBuyQuotaRequest {
	s.Order = v
	return s
}

type ManagementBuyQuotaRequestOrder struct {
	// 订单id
	OrderId *int64 `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 业务类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 容量类型
	CapacityType *string `json:"capacityType,omitempty" xml:"capacityType,omitempty"`
	// 待扩容的容量
	Capacity *int64 `json:"capacity,omitempty" xml:"capacity,omitempty"`
	// 时长
	Day *int32 `json:"day,omitempty" xml:"day,omitempty"`
	// 金额
	Money *int64 `json:"money,omitempty" xml:"money,omitempty"`
}

func (s ManagementBuyQuotaRequestOrder) String() string {
	return tea.Prettify(s)
}

func (s ManagementBuyQuotaRequestOrder) GoString() string {
	return s.String()
}

func (s *ManagementBuyQuotaRequestOrder) SetOrderId(v int64) *ManagementBuyQuotaRequestOrder {
	s.OrderId = &v
	return s
}

func (s *ManagementBuyQuotaRequestOrder) SetBizType(v string) *ManagementBuyQuotaRequestOrder {
	s.BizType = &v
	return s
}

func (s *ManagementBuyQuotaRequestOrder) SetCapacityType(v string) *ManagementBuyQuotaRequestOrder {
	s.CapacityType = &v
	return s
}

func (s *ManagementBuyQuotaRequestOrder) SetCapacity(v int64) *ManagementBuyQuotaRequestOrder {
	s.Capacity = &v
	return s
}

func (s *ManagementBuyQuotaRequestOrder) SetDay(v int32) *ManagementBuyQuotaRequestOrder {
	s.Day = &v
	return s
}

func (s *ManagementBuyQuotaRequestOrder) SetMoney(v int64) *ManagementBuyQuotaRequestOrder {
	s.Money = &v
	return s
}

type ManagementBuyQuotaResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ManagementBuyQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s ManagementBuyQuotaResponse) GoString() string {
	return s.String()
}

func (s *ManagementBuyQuotaResponse) SetHeaders(v map[string]*string) *ManagementBuyQuotaResponse {
	s.Headers = v
	return s
}

type RenameFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RenameFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s RenameFileHeaders) GoString() string {
	return s.String()
}

func (s *RenameFileHeaders) SetCommonHeaders(v map[string]*string) *RenameFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RenameFileHeaders) SetXAcsDingtalkAccessToken(v string) *RenameFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RenameFileRequest struct {
	// 新文件名称
	NewFileName *string `json:"newFileName,omitempty" xml:"newFileName,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RenameFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameFileRequest) GoString() string {
	return s.String()
}

func (s *RenameFileRequest) SetNewFileName(v string) *RenameFileRequest {
	s.NewFileName = &v
	return s
}

func (s *RenameFileRequest) SetUnionId(v string) *RenameFileRequest {
	s.UnionId = &v
	return s
}

type RenameFileResponseBody struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 父目录id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件id
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件路径
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 文件类型
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 文件内容类型
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改者
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s RenameFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameFileResponseBody) GoString() string {
	return s.String()
}

func (s *RenameFileResponseBody) SetSpaceId(v string) *RenameFileResponseBody {
	s.SpaceId = &v
	return s
}

func (s *RenameFileResponseBody) SetParentId(v string) *RenameFileResponseBody {
	s.ParentId = &v
	return s
}

func (s *RenameFileResponseBody) SetFileId(v string) *RenameFileResponseBody {
	s.FileId = &v
	return s
}

func (s *RenameFileResponseBody) SetFileName(v string) *RenameFileResponseBody {
	s.FileName = &v
	return s
}

func (s *RenameFileResponseBody) SetFilePath(v string) *RenameFileResponseBody {
	s.FilePath = &v
	return s
}

func (s *RenameFileResponseBody) SetFileType(v string) *RenameFileResponseBody {
	s.FileType = &v
	return s
}

func (s *RenameFileResponseBody) SetContentType(v string) *RenameFileResponseBody {
	s.ContentType = &v
	return s
}

func (s *RenameFileResponseBody) SetFileExtension(v string) *RenameFileResponseBody {
	s.FileExtension = &v
	return s
}

func (s *RenameFileResponseBody) SetFileSize(v int64) *RenameFileResponseBody {
	s.FileSize = &v
	return s
}

func (s *RenameFileResponseBody) SetCreateTime(v string) *RenameFileResponseBody {
	s.CreateTime = &v
	return s
}

func (s *RenameFileResponseBody) SetModifyTime(v string) *RenameFileResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *RenameFileResponseBody) SetCreator(v string) *RenameFileResponseBody {
	s.Creator = &v
	return s
}

func (s *RenameFileResponseBody) SetModifier(v string) *RenameFileResponseBody {
	s.Modifier = &v
	return s
}

type RenameFileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenameFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameFileResponse) GoString() string {
	return s.String()
}

func (s *RenameFileResponse) SetHeaders(v map[string]*string) *RenameFileResponse {
	s.Headers = v
	return s
}

func (s *RenameFileResponse) SetBody(v *RenameFileResponseBody) *RenameFileResponse {
	s.Body = v
	return s
}

type GetAsyncTaskInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAsyncTaskInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskInfoHeaders) SetCommonHeaders(v map[string]*string) *GetAsyncTaskInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAsyncTaskInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetAsyncTaskInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAsyncTaskInfoRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetAsyncTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskInfoRequest) SetUnionId(v string) *GetAsyncTaskInfoRequest {
	s.UnionId = &v
	return s
}

type GetAsyncTaskInfoResponseBody struct {
	// 异步任务id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 任务总数
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
	// 完成个数
	Success *int32 `json:"success,omitempty" xml:"success,omitempty"`
	// 失败个数
	Failed *int32 `json:"failed,omitempty" xml:"failed,omitempty"`
	// 任务状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 任务开始时间
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// 任务结束时间
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s GetAsyncTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskInfoResponseBody) SetTaskId(v string) *GetAsyncTaskInfoResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetAsyncTaskInfoResponseBody) SetTotal(v int32) *GetAsyncTaskInfoResponseBody {
	s.Total = &v
	return s
}

func (s *GetAsyncTaskInfoResponseBody) SetSuccess(v int32) *GetAsyncTaskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetAsyncTaskInfoResponseBody) SetFailed(v int32) *GetAsyncTaskInfoResponseBody {
	s.Failed = &v
	return s
}

func (s *GetAsyncTaskInfoResponseBody) SetStatus(v string) *GetAsyncTaskInfoResponseBody {
	s.Status = &v
	return s
}

func (s *GetAsyncTaskInfoResponseBody) SetBeginTime(v string) *GetAsyncTaskInfoResponseBody {
	s.BeginTime = &v
	return s
}

func (s *GetAsyncTaskInfoResponseBody) SetEndTime(v string) *GetAsyncTaskInfoResponseBody {
	s.EndTime = &v
	return s
}

type GetAsyncTaskInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskInfoResponse) SetHeaders(v map[string]*string) *GetAsyncTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncTaskInfoResponse) SetBody(v *GetAsyncTaskInfoResponseBody) *GetAsyncTaskInfoResponse {
	s.Body = v
	return s
}

type ListFilesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListFilesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListFilesHeaders) GoString() string {
	return s.String()
}

func (s *ListFilesHeaders) SetCommonHeaders(v map[string]*string) *ListFilesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFilesHeaders) SetXAcsDingtalkAccessToken(v string) *ListFilesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListFilesRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 父目录id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 分页查询锚点
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页长度
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 排序类型
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
}

func (s ListFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFilesRequest) GoString() string {
	return s.String()
}

func (s *ListFilesRequest) SetUnionId(v string) *ListFilesRequest {
	s.UnionId = &v
	return s
}

func (s *ListFilesRequest) SetParentId(v string) *ListFilesRequest {
	s.ParentId = &v
	return s
}

func (s *ListFilesRequest) SetNextToken(v string) *ListFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListFilesRequest) SetMaxResults(v int32) *ListFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFilesRequest) SetOrderType(v string) *ListFilesRequest {
	s.OrderType = &v
	return s
}

type ListFilesResponseBody struct {
	// 文件列表
	Files []*ListFilesResponseBodyFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// 分页加载锚点, nextToken不为空表示有更多数据
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBody) SetFiles(v []*ListFilesResponseBodyFiles) *ListFilesResponseBody {
	s.Files = v
	return s
}

func (s *ListFilesResponseBody) SetNextToken(v string) *ListFilesResponseBody {
	s.NextToken = &v
	return s
}

type ListFilesResponseBodyFiles struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 父目录id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件id
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件路径
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 文件类型
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 文件内容类型
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改者
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s ListFilesResponseBodyFiles) String() string {
	return tea.Prettify(s)
}

func (s ListFilesResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBodyFiles) SetSpaceId(v string) *ListFilesResponseBodyFiles {
	s.SpaceId = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetParentId(v string) *ListFilesResponseBodyFiles {
	s.ParentId = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetFileId(v string) *ListFilesResponseBodyFiles {
	s.FileId = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetFileName(v string) *ListFilesResponseBodyFiles {
	s.FileName = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetFilePath(v string) *ListFilesResponseBodyFiles {
	s.FilePath = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetFileType(v string) *ListFilesResponseBodyFiles {
	s.FileType = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetContentType(v string) *ListFilesResponseBodyFiles {
	s.ContentType = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetFileExtension(v string) *ListFilesResponseBodyFiles {
	s.FileExtension = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetFileSize(v int64) *ListFilesResponseBodyFiles {
	s.FileSize = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetCreateTime(v string) *ListFilesResponseBodyFiles {
	s.CreateTime = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetModifyTime(v string) *ListFilesResponseBodyFiles {
	s.ModifyTime = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetCreator(v string) *ListFilesResponseBodyFiles {
	s.Creator = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetModifier(v string) *ListFilesResponseBodyFiles {
	s.Modifier = &v
	return s
}

type ListFilesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFilesResponse) GoString() string {
	return s.String()
}

func (s *ListFilesResponse) SetHeaders(v map[string]*string) *ListFilesResponse {
	s.Headers = v
	return s
}

func (s *ListFilesResponse) SetBody(v *ListFilesResponseBody) *ListFilesResponse {
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
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequest) SetUnionId(v string) *ListPermissionsRequest {
	s.UnionId = &v
	return s
}

type ListPermissionsResponseBody struct {
	// 企业内成员权限列表
	Members []*ListPermissionsResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 企业外成员权限列表
	OutMembers []*ListPermissionsResponseBodyOutMembers `json:"outMembers,omitempty" xml:"outMembers,omitempty" type:"Repeated"`
}

func (s ListPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBody) SetMembers(v []*ListPermissionsResponseBodyMembers) *ListPermissionsResponseBody {
	s.Members = v
	return s
}

func (s *ListPermissionsResponseBody) SetOutMembers(v []*ListPermissionsResponseBodyOutMembers) *ListPermissionsResponseBody {
	s.OutMembers = v
	return s
}

type ListPermissionsResponseBodyMembers struct {
	// 权限角色
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 成员信息
	Member *ListPermissionsResponseBodyMembersMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	// 是否是继承的权限
	Extend *bool `json:"extend,omitempty" xml:"extend,omitempty"`
}

func (s ListPermissionsResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyMembers) SetRole(v string) *ListPermissionsResponseBodyMembers {
	s.Role = &v
	return s
}

func (s *ListPermissionsResponseBodyMembers) SetMember(v *ListPermissionsResponseBodyMembersMember) *ListPermissionsResponseBodyMembers {
	s.Member = v
	return s
}

func (s *ListPermissionsResponseBodyMembers) SetExtend(v bool) *ListPermissionsResponseBodyMembers {
	s.Extend = &v
	return s
}

type ListPermissionsResponseBodyMembersMember struct {
	// 企业corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 成员类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 成员id
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// 成员名称
	MemberName *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
}

func (s ListPermissionsResponseBodyMembersMember) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyMembersMember) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyMembersMember) SetCorpId(v string) *ListPermissionsResponseBodyMembersMember {
	s.CorpId = &v
	return s
}

func (s *ListPermissionsResponseBodyMembersMember) SetMemberType(v string) *ListPermissionsResponseBodyMembersMember {
	s.MemberType = &v
	return s
}

func (s *ListPermissionsResponseBodyMembersMember) SetMemberId(v string) *ListPermissionsResponseBodyMembersMember {
	s.MemberId = &v
	return s
}

func (s *ListPermissionsResponseBodyMembersMember) SetMemberName(v string) *ListPermissionsResponseBodyMembersMember {
	s.MemberName = &v
	return s
}

type ListPermissionsResponseBodyOutMembers struct {
	// 权限角色
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 成员信息
	Member *ListPermissionsResponseBodyOutMembersMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	// 是否是继承的权限
	Extend *bool `json:"extend,omitempty" xml:"extend,omitempty"`
}

func (s ListPermissionsResponseBodyOutMembers) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyOutMembers) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyOutMembers) SetRole(v string) *ListPermissionsResponseBodyOutMembers {
	s.Role = &v
	return s
}

func (s *ListPermissionsResponseBodyOutMembers) SetMember(v *ListPermissionsResponseBodyOutMembersMember) *ListPermissionsResponseBodyOutMembers {
	s.Member = v
	return s
}

func (s *ListPermissionsResponseBodyOutMembers) SetExtend(v bool) *ListPermissionsResponseBodyOutMembers {
	s.Extend = &v
	return s
}

type ListPermissionsResponseBodyOutMembersMember struct {
	// 企业corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 成员类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 成员id
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// 成员名称
	MemberName *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
}

func (s ListPermissionsResponseBodyOutMembersMember) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyOutMembersMember) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyOutMembersMember) SetCorpId(v string) *ListPermissionsResponseBodyOutMembersMember {
	s.CorpId = &v
	return s
}

func (s *ListPermissionsResponseBodyOutMembersMember) SetMemberType(v string) *ListPermissionsResponseBodyOutMembersMember {
	s.MemberType = &v
	return s
}

func (s *ListPermissionsResponseBodyOutMembersMember) SetMemberId(v string) *ListPermissionsResponseBodyOutMembersMember {
	s.MemberId = &v
	return s
}

func (s *ListPermissionsResponseBodyOutMembersMember) SetMemberName(v string) *ListPermissionsResponseBodyOutMembersMember {
	s.MemberName = &v
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

type MoveFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MoveFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s MoveFileHeaders) GoString() string {
	return s.String()
}

func (s *MoveFileHeaders) SetCommonHeaders(v map[string]*string) *MoveFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MoveFileHeaders) SetXAcsDingtalkAccessToken(v string) *MoveFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MoveFileRequest struct {
	// 目标空间id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// 目标父目录id
	TargetParentId *string `json:"targetParentId,omitempty" xml:"targetParentId,omitempty"`
	// 文件名冲突策略
	AddConflictPolicy *string `json:"addConflictPolicy,omitempty" xml:"addConflictPolicy,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s MoveFileRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveFileRequest) GoString() string {
	return s.String()
}

func (s *MoveFileRequest) SetTargetSpaceId(v string) *MoveFileRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *MoveFileRequest) SetTargetParentId(v string) *MoveFileRequest {
	s.TargetParentId = &v
	return s
}

func (s *MoveFileRequest) SetAddConflictPolicy(v string) *MoveFileRequest {
	s.AddConflictPolicy = &v
	return s
}

func (s *MoveFileRequest) SetUnionId(v string) *MoveFileRequest {
	s.UnionId = &v
	return s
}

type MoveFileResponseBody struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 父目录id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件id
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件路径
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 文件类型
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 文件内容类型
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改者
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s MoveFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveFileResponseBody) GoString() string {
	return s.String()
}

func (s *MoveFileResponseBody) SetSpaceId(v string) *MoveFileResponseBody {
	s.SpaceId = &v
	return s
}

func (s *MoveFileResponseBody) SetParentId(v string) *MoveFileResponseBody {
	s.ParentId = &v
	return s
}

func (s *MoveFileResponseBody) SetFileId(v string) *MoveFileResponseBody {
	s.FileId = &v
	return s
}

func (s *MoveFileResponseBody) SetFileName(v string) *MoveFileResponseBody {
	s.FileName = &v
	return s
}

func (s *MoveFileResponseBody) SetFilePath(v string) *MoveFileResponseBody {
	s.FilePath = &v
	return s
}

func (s *MoveFileResponseBody) SetFileType(v string) *MoveFileResponseBody {
	s.FileType = &v
	return s
}

func (s *MoveFileResponseBody) SetContentType(v string) *MoveFileResponseBody {
	s.ContentType = &v
	return s
}

func (s *MoveFileResponseBody) SetFileExtension(v string) *MoveFileResponseBody {
	s.FileExtension = &v
	return s
}

func (s *MoveFileResponseBody) SetFileSize(v int64) *MoveFileResponseBody {
	s.FileSize = &v
	return s
}

func (s *MoveFileResponseBody) SetCreateTime(v string) *MoveFileResponseBody {
	s.CreateTime = &v
	return s
}

func (s *MoveFileResponseBody) SetModifyTime(v string) *MoveFileResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *MoveFileResponseBody) SetCreator(v string) *MoveFileResponseBody {
	s.Creator = &v
	return s
}

func (s *MoveFileResponseBody) SetModifier(v string) *MoveFileResponseBody {
	s.Modifier = &v
	return s
}

type MoveFileResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveFileResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveFileResponse) GoString() string {
	return s.String()
}

func (s *MoveFileResponse) SetHeaders(v map[string]*string) *MoveFileResponse {
	s.Headers = v
	return s
}

func (s *MoveFileResponse) SetBody(v *MoveFileResponseBody) *MoveFileResponse {
	s.Body = v
	return s
}

type ListSpacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSpacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSpacesHeaders) GoString() string {
	return s.String()
}

func (s *ListSpacesHeaders) SetCommonHeaders(v map[string]*string) *ListSpacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSpacesHeaders) SetXAcsDingtalkAccessToken(v string) *ListSpacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSpacesRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 空间类型
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 分页加载锚点
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页大小
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s ListSpacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpacesRequest) GoString() string {
	return s.String()
}

func (s *ListSpacesRequest) SetUnionId(v string) *ListSpacesRequest {
	s.UnionId = &v
	return s
}

func (s *ListSpacesRequest) SetSpaceType(v string) *ListSpacesRequest {
	s.SpaceType = &v
	return s
}

func (s *ListSpacesRequest) SetNextToken(v string) *ListSpacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSpacesRequest) SetMaxResults(v int32) *ListSpacesRequest {
	s.MaxResults = &v
	return s
}

type ListSpacesResponseBody struct {
	Spaces []*ListSpacesResponseBodySpaces `json:"spaces,omitempty" xml:"spaces,omitempty" type:"Repeated"`
	// 分页加载更多锚点, nextToken不为空表示有更多数据
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListSpacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpacesResponseBody) SetSpaces(v []*ListSpacesResponseBodySpaces) *ListSpacesResponseBody {
	s.Spaces = v
	return s
}

func (s *ListSpacesResponseBody) SetNextToken(v string) *ListSpacesResponseBody {
	s.NextToken = &v
	return s
}

type ListSpacesResponseBodySpaces struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 空间名称
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	// 空间类型
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 空间总额度
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 空间已使用额度
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	// 授权模式
	PermissionMode *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
}

func (s ListSpacesResponseBodySpaces) String() string {
	return tea.Prettify(s)
}

func (s ListSpacesResponseBodySpaces) GoString() string {
	return s.String()
}

func (s *ListSpacesResponseBodySpaces) SetSpaceId(v string) *ListSpacesResponseBodySpaces {
	s.SpaceId = &v
	return s
}

func (s *ListSpacesResponseBodySpaces) SetSpaceName(v string) *ListSpacesResponseBodySpaces {
	s.SpaceName = &v
	return s
}

func (s *ListSpacesResponseBodySpaces) SetSpaceType(v string) *ListSpacesResponseBodySpaces {
	s.SpaceType = &v
	return s
}

func (s *ListSpacesResponseBodySpaces) SetQuota(v int64) *ListSpacesResponseBodySpaces {
	s.Quota = &v
	return s
}

func (s *ListSpacesResponseBodySpaces) SetUsedQuota(v int64) *ListSpacesResponseBodySpaces {
	s.UsedQuota = &v
	return s
}

func (s *ListSpacesResponseBodySpaces) SetPermissionMode(v string) *ListSpacesResponseBodySpaces {
	s.PermissionMode = &v
	return s
}

func (s *ListSpacesResponseBodySpaces) SetCreateTime(v string) *ListSpacesResponseBodySpaces {
	s.CreateTime = &v
	return s
}

func (s *ListSpacesResponseBodySpaces) SetModifyTime(v string) *ListSpacesResponseBodySpaces {
	s.ModifyTime = &v
	return s
}

type ListSpacesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSpacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpacesResponse) GoString() string {
	return s.String()
}

func (s *ListSpacesResponse) SetHeaders(v map[string]*string) *ListSpacesResponse {
	s.Headers = v
	return s
}

func (s *ListSpacesResponse) SetBody(v *ListSpacesResponseBody) *ListSpacesResponse {
	s.Body = v
	return s
}

type CopyFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CopyFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s CopyFileHeaders) GoString() string {
	return s.String()
}

func (s *CopyFileHeaders) SetCommonHeaders(v map[string]*string) *CopyFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyFileHeaders) SetXAcsDingtalkAccessToken(v string) *CopyFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CopyFileRequest struct {
	// 目标空间id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// 目标父目录id
	TargetParentId *string `json:"targetParentId,omitempty" xml:"targetParentId,omitempty"`
	// 文件名冲突策略
	AddConflictPolicy *string `json:"addConflictPolicy,omitempty" xml:"addConflictPolicy,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CopyFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyFileRequest) GoString() string {
	return s.String()
}

func (s *CopyFileRequest) SetTargetSpaceId(v string) *CopyFileRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *CopyFileRequest) SetTargetParentId(v string) *CopyFileRequest {
	s.TargetParentId = &v
	return s
}

func (s *CopyFileRequest) SetAddConflictPolicy(v string) *CopyFileRequest {
	s.AddConflictPolicy = &v
	return s
}

func (s *CopyFileRequest) SetUnionId(v string) *CopyFileRequest {
	s.UnionId = &v
	return s
}

type CopyFileResponseBody struct {
	// 文件信息
	File *CopyFileResponseBodyFile `json:"file,omitempty" xml:"file,omitempty" type:"Struct"`
	// 异步任务id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CopyFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyFileResponseBody) GoString() string {
	return s.String()
}

func (s *CopyFileResponseBody) SetFile(v *CopyFileResponseBodyFile) *CopyFileResponseBody {
	s.File = v
	return s
}

func (s *CopyFileResponseBody) SetTaskId(v string) *CopyFileResponseBody {
	s.TaskId = &v
	return s
}

type CopyFileResponseBodyFile struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 父目录id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件id
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件路径
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 文件类型
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 文件内容类型
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改者
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s CopyFileResponseBodyFile) String() string {
	return tea.Prettify(s)
}

func (s CopyFileResponseBodyFile) GoString() string {
	return s.String()
}

func (s *CopyFileResponseBodyFile) SetSpaceId(v string) *CopyFileResponseBodyFile {
	s.SpaceId = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetParentId(v string) *CopyFileResponseBodyFile {
	s.ParentId = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetFileId(v string) *CopyFileResponseBodyFile {
	s.FileId = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetFileName(v string) *CopyFileResponseBodyFile {
	s.FileName = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetFilePath(v string) *CopyFileResponseBodyFile {
	s.FilePath = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetFileType(v string) *CopyFileResponseBodyFile {
	s.FileType = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetContentType(v string) *CopyFileResponseBodyFile {
	s.ContentType = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetFileExtension(v string) *CopyFileResponseBodyFile {
	s.FileExtension = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetFileSize(v int64) *CopyFileResponseBodyFile {
	s.FileSize = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetCreateTime(v string) *CopyFileResponseBodyFile {
	s.CreateTime = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetModifyTime(v string) *CopyFileResponseBodyFile {
	s.ModifyTime = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetCreator(v string) *CopyFileResponseBodyFile {
	s.Creator = &v
	return s
}

func (s *CopyFileResponseBodyFile) SetModifier(v string) *CopyFileResponseBodyFile {
	s.Modifier = &v
	return s
}

type CopyFileResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CopyFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CopyFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyFileResponse) GoString() string {
	return s.String()
}

func (s *CopyFileResponse) SetHeaders(v map[string]*string) *CopyFileResponse {
	s.Headers = v
	return s
}

func (s *CopyFileResponse) SetBody(v *CopyFileResponseBody) *CopyFileResponse {
	s.Body = v
	return s
}

type DeleteSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSpaceHeaders) SetCommonHeaders(v map[string]*string) *DeleteSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteSpaceRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpaceRequest) SetUnionId(v string) *DeleteSpaceRequest {
	s.UnionId = &v
	return s
}

type DeleteSpaceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpaceResponse) SetHeaders(v map[string]*string) *DeleteSpaceResponse {
	s.Headers = v
	return s
}

type ClearRecycleFilesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ClearRecycleFilesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ClearRecycleFilesHeaders) GoString() string {
	return s.String()
}

func (s *ClearRecycleFilesHeaders) SetCommonHeaders(v map[string]*string) *ClearRecycleFilesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ClearRecycleFilesHeaders) SetXAcsDingtalkAccessToken(v string) *ClearRecycleFilesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ClearRecycleFilesRequest struct {
	// 回收站类型
	RecycleType *string `json:"recycleType,omitempty" xml:"recycleType,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ClearRecycleFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearRecycleFilesRequest) GoString() string {
	return s.String()
}

func (s *ClearRecycleFilesRequest) SetRecycleType(v string) *ClearRecycleFilesRequest {
	s.RecycleType = &v
	return s
}

func (s *ClearRecycleFilesRequest) SetUnionId(v string) *ClearRecycleFilesRequest {
	s.UnionId = &v
	return s
}

type ClearRecycleFilesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ClearRecycleFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearRecycleFilesResponse) GoString() string {
	return s.String()
}

func (s *ClearRecycleFilesResponse) SetHeaders(v map[string]*string) *ClearRecycleFilesResponse {
	s.Headers = v
	return s
}

type ManagementListSpacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ManagementListSpacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ManagementListSpacesHeaders) GoString() string {
	return s.String()
}

func (s *ManagementListSpacesHeaders) SetCommonHeaders(v map[string]*string) *ManagementListSpacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ManagementListSpacesHeaders) SetXAcsDingtalkAccessToken(v string) *ManagementListSpacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ManagementListSpacesRequest struct {
	// 空间id列表
	SpaceIds []*string `json:"spaceIds,omitempty" xml:"spaceIds,omitempty" type:"Repeated"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ManagementListSpacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ManagementListSpacesRequest) GoString() string {
	return s.String()
}

func (s *ManagementListSpacesRequest) SetSpaceIds(v []*string) *ManagementListSpacesRequest {
	s.SpaceIds = v
	return s
}

func (s *ManagementListSpacesRequest) SetUnionId(v string) *ManagementListSpacesRequest {
	s.UnionId = &v
	return s
}

type ManagementListSpacesResponseBody struct {
	Spaces []*ManagementListSpacesResponseBodySpaces `json:"spaces,omitempty" xml:"spaces,omitempty" type:"Repeated"`
}

func (s ManagementListSpacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManagementListSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *ManagementListSpacesResponseBody) SetSpaces(v []*ManagementListSpacesResponseBodySpaces) *ManagementListSpacesResponseBody {
	s.Spaces = v
	return s
}

type ManagementListSpacesResponseBodySpaces struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 空间名称
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	// 空间类型
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 空间总额度
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 空间已使用额度
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	// 授权模式
	PermissionMode *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
}

func (s ManagementListSpacesResponseBodySpaces) String() string {
	return tea.Prettify(s)
}

func (s ManagementListSpacesResponseBodySpaces) GoString() string {
	return s.String()
}

func (s *ManagementListSpacesResponseBodySpaces) SetSpaceId(v string) *ManagementListSpacesResponseBodySpaces {
	s.SpaceId = &v
	return s
}

func (s *ManagementListSpacesResponseBodySpaces) SetSpaceName(v string) *ManagementListSpacesResponseBodySpaces {
	s.SpaceName = &v
	return s
}

func (s *ManagementListSpacesResponseBodySpaces) SetSpaceType(v string) *ManagementListSpacesResponseBodySpaces {
	s.SpaceType = &v
	return s
}

func (s *ManagementListSpacesResponseBodySpaces) SetQuota(v int64) *ManagementListSpacesResponseBodySpaces {
	s.Quota = &v
	return s
}

func (s *ManagementListSpacesResponseBodySpaces) SetUsedQuota(v int64) *ManagementListSpacesResponseBodySpaces {
	s.UsedQuota = &v
	return s
}

func (s *ManagementListSpacesResponseBodySpaces) SetPermissionMode(v string) *ManagementListSpacesResponseBodySpaces {
	s.PermissionMode = &v
	return s
}

func (s *ManagementListSpacesResponseBodySpaces) SetCreateTime(v string) *ManagementListSpacesResponseBodySpaces {
	s.CreateTime = &v
	return s
}

func (s *ManagementListSpacesResponseBodySpaces) SetModifyTime(v string) *ManagementListSpacesResponseBodySpaces {
	s.ModifyTime = &v
	return s
}

type ManagementListSpacesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ManagementListSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ManagementListSpacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ManagementListSpacesResponse) GoString() string {
	return s.String()
}

func (s *ManagementListSpacesResponse) SetHeaders(v map[string]*string) *ManagementListSpacesResponse {
	s.Headers = v
	return s
}

func (s *ManagementListSpacesResponse) SetBody(v *ManagementListSpacesResponseBody) *ManagementListSpacesResponse {
	s.Body = v
	return s
}

type DeleteFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFileHeaders) SetCommonHeaders(v map[string]*string) *DeleteFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFileHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteFileRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 删除策略
	DeletePolicy *string `json:"deletePolicy,omitempty" xml:"deletePolicy,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) SetUnionId(v string) *DeleteFileRequest {
	s.UnionId = &v
	return s
}

func (s *DeleteFileRequest) SetDeletePolicy(v string) *DeleteFileRequest {
	s.DeletePolicy = &v
	return s
}

type DeleteFileResponseBody struct {
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetSuccess(v bool) *DeleteFileResponseBody {
	s.Success = &v
	return s
}

type DeleteFileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

type AddFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddFileHeaders) GoString() string {
	return s.String()
}

func (s *AddFileHeaders) SetCommonHeaders(v map[string]*string) *AddFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddFileHeaders) SetXAcsDingtalkAccessToken(v string) *AddFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddFileRequest struct {
	// 父目录id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件类型
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// mediaId
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// 文件名冲突策略
	AddConflictPolicy *string `json:"addConflictPolicy,omitempty" xml:"addConflictPolicy,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddFileRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFileRequest) GoString() string {
	return s.String()
}

func (s *AddFileRequest) SetParentId(v string) *AddFileRequest {
	s.ParentId = &v
	return s
}

func (s *AddFileRequest) SetFileType(v string) *AddFileRequest {
	s.FileType = &v
	return s
}

func (s *AddFileRequest) SetFileName(v string) *AddFileRequest {
	s.FileName = &v
	return s
}

func (s *AddFileRequest) SetMediaId(v string) *AddFileRequest {
	s.MediaId = &v
	return s
}

func (s *AddFileRequest) SetAddConflictPolicy(v string) *AddFileRequest {
	s.AddConflictPolicy = &v
	return s
}

func (s *AddFileRequest) SetUnionId(v string) *AddFileRequest {
	s.UnionId = &v
	return s
}

type AddFileResponseBody struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 父目录id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件id
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件路径
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 文件类型
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 文件内容类型
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改者
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s AddFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponseBody) GoString() string {
	return s.String()
}

func (s *AddFileResponseBody) SetSpaceId(v string) *AddFileResponseBody {
	s.SpaceId = &v
	return s
}

func (s *AddFileResponseBody) SetParentId(v string) *AddFileResponseBody {
	s.ParentId = &v
	return s
}

func (s *AddFileResponseBody) SetFileId(v string) *AddFileResponseBody {
	s.FileId = &v
	return s
}

func (s *AddFileResponseBody) SetFileName(v string) *AddFileResponseBody {
	s.FileName = &v
	return s
}

func (s *AddFileResponseBody) SetFilePath(v string) *AddFileResponseBody {
	s.FilePath = &v
	return s
}

func (s *AddFileResponseBody) SetFileType(v string) *AddFileResponseBody {
	s.FileType = &v
	return s
}

func (s *AddFileResponseBody) SetContentType(v string) *AddFileResponseBody {
	s.ContentType = &v
	return s
}

func (s *AddFileResponseBody) SetFileExtension(v string) *AddFileResponseBody {
	s.FileExtension = &v
	return s
}

func (s *AddFileResponseBody) SetFileSize(v int64) *AddFileResponseBody {
	s.FileSize = &v
	return s
}

func (s *AddFileResponseBody) SetCreateTime(v string) *AddFileResponseBody {
	s.CreateTime = &v
	return s
}

func (s *AddFileResponseBody) SetModifyTime(v string) *AddFileResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *AddFileResponseBody) SetCreator(v string) *AddFileResponseBody {
	s.Creator = &v
	return s
}

func (s *AddFileResponseBody) SetModifier(v string) *AddFileResponseBody {
	s.Modifier = &v
	return s
}

type AddFileResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddFileResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponse) GoString() string {
	return s.String()
}

func (s *AddFileResponse) SetHeaders(v map[string]*string) *AddFileResponse {
	s.Headers = v
	return s
}

func (s *AddFileResponse) SetBody(v *AddFileResponseBody) *AddFileResponse {
	s.Body = v
	return s
}

type GetPreviewInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPreviewInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPreviewInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetPreviewInfoHeaders) SetCommonHeaders(v map[string]*string) *GetPreviewInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPreviewInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetPreviewInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPreviewInfoRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetPreviewInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPreviewInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPreviewInfoRequest) SetUnionId(v string) *GetPreviewInfoRequest {
	s.UnionId = &v
	return s
}

type GetPreviewInfoResponseBody struct {
	// 预览信息
	Info *GetPreviewInfoResponseBodyInfo `json:"info,omitempty" xml:"info,omitempty" type:"Struct"`
}

func (s GetPreviewInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPreviewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPreviewInfoResponseBody) SetInfo(v *GetPreviewInfoResponseBodyInfo) *GetPreviewInfoResponseBody {
	s.Info = v
	return s
}

type GetPreviewInfoResponseBodyInfo struct {
	// 预览url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetPreviewInfoResponseBodyInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPreviewInfoResponseBodyInfo) GoString() string {
	return s.String()
}

func (s *GetPreviewInfoResponseBodyInfo) SetUrl(v string) *GetPreviewInfoResponseBodyInfo {
	s.Url = &v
	return s
}

type GetPreviewInfoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPreviewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPreviewInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPreviewInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPreviewInfoResponse) SetHeaders(v map[string]*string) *GetPreviewInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPreviewInfoResponse) SetBody(v *GetPreviewInfoResponseBody) *GetPreviewInfoResponse {
	s.Body = v
	return s
}

type InfoSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InfoSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s InfoSpaceHeaders) GoString() string {
	return s.String()
}

func (s *InfoSpaceHeaders) SetCommonHeaders(v map[string]*string) *InfoSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InfoSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *InfoSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InfoSpaceRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s InfoSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s InfoSpaceRequest) GoString() string {
	return s.String()
}

func (s *InfoSpaceRequest) SetUnionId(v string) *InfoSpaceRequest {
	s.UnionId = &v
	return s
}

type InfoSpaceResponseBody struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 空间名称
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	// 空间类型
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 容量
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 已使用容量
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	// 授权模式
	PermissionMode *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
}

func (s InfoSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InfoSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *InfoSpaceResponseBody) SetSpaceId(v string) *InfoSpaceResponseBody {
	s.SpaceId = &v
	return s
}

func (s *InfoSpaceResponseBody) SetSpaceName(v string) *InfoSpaceResponseBody {
	s.SpaceName = &v
	return s
}

func (s *InfoSpaceResponseBody) SetSpaceType(v string) *InfoSpaceResponseBody {
	s.SpaceType = &v
	return s
}

func (s *InfoSpaceResponseBody) SetQuota(v int64) *InfoSpaceResponseBody {
	s.Quota = &v
	return s
}

func (s *InfoSpaceResponseBody) SetUsedQuota(v int64) *InfoSpaceResponseBody {
	s.UsedQuota = &v
	return s
}

func (s *InfoSpaceResponseBody) SetPermissionMode(v string) *InfoSpaceResponseBody {
	s.PermissionMode = &v
	return s
}

func (s *InfoSpaceResponseBody) SetCreateTime(v string) *InfoSpaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *InfoSpaceResponseBody) SetModifyTime(v string) *InfoSpaceResponseBody {
	s.ModifyTime = &v
	return s
}

type InfoSpaceResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InfoSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InfoSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s InfoSpaceResponse) GoString() string {
	return s.String()
}

func (s *InfoSpaceResponse) SetHeaders(v map[string]*string) *InfoSpaceResponse {
	s.Headers = v
	return s
}

func (s *InfoSpaceResponse) SetBody(v *InfoSpaceResponseBody) *InfoSpaceResponse {
	s.Body = v
	return s
}

type ManagementModifySpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ManagementModifySpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ManagementModifySpaceHeaders) GoString() string {
	return s.String()
}

func (s *ManagementModifySpaceHeaders) SetCommonHeaders(v map[string]*string) *ManagementModifySpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ManagementModifySpaceHeaders) SetXAcsDingtalkAccessToken(v string) *ManagementModifySpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ManagementModifySpaceRequest struct {
	// 空间id列表
	SpaceIds []*string `json:"spaceIds,omitempty" xml:"spaceIds,omitempty" type:"Repeated"`
	// 容量
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ManagementModifySpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s ManagementModifySpaceRequest) GoString() string {
	return s.String()
}

func (s *ManagementModifySpaceRequest) SetSpaceIds(v []*string) *ManagementModifySpaceRequest {
	s.SpaceIds = v
	return s
}

func (s *ManagementModifySpaceRequest) SetQuota(v int64) *ManagementModifySpaceRequest {
	s.Quota = &v
	return s
}

func (s *ManagementModifySpaceRequest) SetUnionId(v string) *ManagementModifySpaceRequest {
	s.UnionId = &v
	return s
}

type ManagementModifySpaceResponseBody struct {
	// 空间列表
	Spaces []*ManagementModifySpaceResponseBodySpaces `json:"spaces,omitempty" xml:"spaces,omitempty" type:"Repeated"`
}

func (s ManagementModifySpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManagementModifySpaceResponseBody) GoString() string {
	return s.String()
}

func (s *ManagementModifySpaceResponseBody) SetSpaces(v []*ManagementModifySpaceResponseBodySpaces) *ManagementModifySpaceResponseBody {
	s.Spaces = v
	return s
}

type ManagementModifySpaceResponseBodySpaces struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 空间名称
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	// 空间类型
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 空间总额度
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 空间已使用额度
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	// 授权模式
	PermissionMode *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
}

func (s ManagementModifySpaceResponseBodySpaces) String() string {
	return tea.Prettify(s)
}

func (s ManagementModifySpaceResponseBodySpaces) GoString() string {
	return s.String()
}

func (s *ManagementModifySpaceResponseBodySpaces) SetSpaceId(v string) *ManagementModifySpaceResponseBodySpaces {
	s.SpaceId = &v
	return s
}

func (s *ManagementModifySpaceResponseBodySpaces) SetSpaceName(v string) *ManagementModifySpaceResponseBodySpaces {
	s.SpaceName = &v
	return s
}

func (s *ManagementModifySpaceResponseBodySpaces) SetSpaceType(v string) *ManagementModifySpaceResponseBodySpaces {
	s.SpaceType = &v
	return s
}

func (s *ManagementModifySpaceResponseBodySpaces) SetQuota(v int64) *ManagementModifySpaceResponseBodySpaces {
	s.Quota = &v
	return s
}

func (s *ManagementModifySpaceResponseBodySpaces) SetUsedQuota(v int64) *ManagementModifySpaceResponseBodySpaces {
	s.UsedQuota = &v
	return s
}

func (s *ManagementModifySpaceResponseBodySpaces) SetPermissionMode(v string) *ManagementModifySpaceResponseBodySpaces {
	s.PermissionMode = &v
	return s
}

func (s *ManagementModifySpaceResponseBodySpaces) SetCreateTime(v string) *ManagementModifySpaceResponseBodySpaces {
	s.CreateTime = &v
	return s
}

func (s *ManagementModifySpaceResponseBodySpaces) SetModifyTime(v string) *ManagementModifySpaceResponseBodySpaces {
	s.ModifyTime = &v
	return s
}

type ManagementModifySpaceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ManagementModifySpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ManagementModifySpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ManagementModifySpaceResponse) GoString() string {
	return s.String()
}

func (s *ManagementModifySpaceResponse) SetHeaders(v map[string]*string) *ManagementModifySpaceResponse {
	s.Headers = v
	return s
}

func (s *ManagementModifySpaceResponse) SetBody(v *ManagementModifySpaceResponseBody) *ManagementModifySpaceResponse {
	s.Body = v
	return s
}

type ModifyPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ModifyPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s ModifyPermissionHeaders) GoString() string {
	return s.String()
}

func (s *ModifyPermissionHeaders) SetCommonHeaders(v map[string]*string) *ModifyPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ModifyPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *ModifyPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ModifyPermissionRequest struct {
	// 权限角色
	Role    *string                           `json:"role,omitempty" xml:"role,omitempty"`
	Members []*ModifyPermissionRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ModifyPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyPermissionRequest) SetRole(v string) *ModifyPermissionRequest {
	s.Role = &v
	return s
}

func (s *ModifyPermissionRequest) SetMembers(v []*ModifyPermissionRequestMembers) *ModifyPermissionRequest {
	s.Members = v
	return s
}

func (s *ModifyPermissionRequest) SetUnionId(v string) *ModifyPermissionRequest {
	s.UnionId = &v
	return s
}

type ModifyPermissionRequestMembers struct {
	// 企业corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 成员类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 成员id
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
}

func (s ModifyPermissionRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s ModifyPermissionRequestMembers) GoString() string {
	return s.String()
}

func (s *ModifyPermissionRequestMembers) SetCorpId(v string) *ModifyPermissionRequestMembers {
	s.CorpId = &v
	return s
}

func (s *ModifyPermissionRequestMembers) SetMemberType(v string) *ModifyPermissionRequestMembers {
	s.MemberType = &v
	return s
}

func (s *ModifyPermissionRequestMembers) SetMemberId(v string) *ModifyPermissionRequestMembers {
	s.MemberId = &v
	return s
}

type ModifyPermissionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ModifyPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPermissionResponse) GoString() string {
	return s.String()
}

func (s *ModifyPermissionResponse) SetHeaders(v map[string]*string) *ModifyPermissionResponse {
	s.Headers = v
	return s
}

type GrantPrivilegeOfCustomSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GrantPrivilegeOfCustomSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GrantPrivilegeOfCustomSpaceHeaders) GoString() string {
	return s.String()
}

func (s *GrantPrivilegeOfCustomSpaceHeaders) SetCommonHeaders(v map[string]*string) *GrantPrivilegeOfCustomSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GrantPrivilegeOfCustomSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *GrantPrivilegeOfCustomSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GrantPrivilegeOfCustomSpaceRequest struct {
	// 权限类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 被授予权限的员工id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 授权访问的文件id列表
	FileIds []*string `json:"fileIds,omitempty" xml:"fileIds,omitempty" type:"Repeated"`
	// 权限有效时间
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GrantPrivilegeOfCustomSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantPrivilegeOfCustomSpaceRequest) GoString() string {
	return s.String()
}

func (s *GrantPrivilegeOfCustomSpaceRequest) SetType(v string) *GrantPrivilegeOfCustomSpaceRequest {
	s.Type = &v
	return s
}

func (s *GrantPrivilegeOfCustomSpaceRequest) SetUserId(v string) *GrantPrivilegeOfCustomSpaceRequest {
	s.UserId = &v
	return s
}

func (s *GrantPrivilegeOfCustomSpaceRequest) SetFileIds(v []*string) *GrantPrivilegeOfCustomSpaceRequest {
	s.FileIds = v
	return s
}

func (s *GrantPrivilegeOfCustomSpaceRequest) SetDuration(v int64) *GrantPrivilegeOfCustomSpaceRequest {
	s.Duration = &v
	return s
}

func (s *GrantPrivilegeOfCustomSpaceRequest) SetUnionId(v string) *GrantPrivilegeOfCustomSpaceRequest {
	s.UnionId = &v
	return s
}

type GrantPrivilegeOfCustomSpaceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GrantPrivilegeOfCustomSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantPrivilegeOfCustomSpaceResponse) GoString() string {
	return s.String()
}

func (s *GrantPrivilegeOfCustomSpaceResponse) SetHeaders(v map[string]*string) *GrantPrivilegeOfCustomSpaceResponse {
	s.Headers = v
	return s
}

type GetDownloadInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDownloadInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDownloadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDownloadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDownloadInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetDownloadInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDownloadInfoRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetDownloadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDownloadInfoRequest) SetUnionId(v string) *GetDownloadInfoRequest {
	s.UnionId = &v
	return s
}

type GetDownloadInfoResponseBody struct {
	// 下载加签url信息
	DownloadInfo *GetDownloadInfoResponseBodyDownloadInfo `json:"downloadInfo,omitempty" xml:"downloadInfo,omitempty" type:"Struct"`
}

func (s GetDownloadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDownloadInfoResponseBody) SetDownloadInfo(v *GetDownloadInfoResponseBodyDownloadInfo) *GetDownloadInfoResponseBody {
	s.DownloadInfo = v
	return s
}

type GetDownloadInfoResponseBodyDownloadInfo struct {
	// 加签url
	ResourceUrl *string `json:"resourceUrl,omitempty" xml:"resourceUrl,omitempty"`
	// 加签url过期时间
	ExpirationSeconds *int32 `json:"expirationSeconds,omitempty" xml:"expirationSeconds,omitempty"`
	// headers
	Headers map[string]interface{} `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s GetDownloadInfoResponseBodyDownloadInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadInfoResponseBodyDownloadInfo) GoString() string {
	return s.String()
}

func (s *GetDownloadInfoResponseBodyDownloadInfo) SetResourceUrl(v string) *GetDownloadInfoResponseBodyDownloadInfo {
	s.ResourceUrl = &v
	return s
}

func (s *GetDownloadInfoResponseBodyDownloadInfo) SetExpirationSeconds(v int32) *GetDownloadInfoResponseBodyDownloadInfo {
	s.ExpirationSeconds = &v
	return s
}

func (s *GetDownloadInfoResponseBodyDownloadInfo) SetHeaders(v map[string]interface{}) *GetDownloadInfoResponseBodyDownloadInfo {
	s.Headers = v
	return s
}

type GetDownloadInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDownloadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDownloadInfoResponse) SetHeaders(v map[string]*string) *GetDownloadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDownloadInfoResponse) SetBody(v *GetDownloadInfoResponseBody) *GetDownloadInfoResponse {
	s.Body = v
	return s
}

type GetUploadInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUploadInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUploadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetUploadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetUploadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUploadInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetUploadInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUploadInfoRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 文件md5
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 文件名称冲突策略
	AddConflictPolicy *string `json:"addConflictPolicy,omitempty" xml:"addConflictPolicy,omitempty"`
	// mediaId
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

func (s GetUploadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUploadInfoRequest) SetUnionId(v string) *GetUploadInfoRequest {
	s.UnionId = &v
	return s
}

func (s *GetUploadInfoRequest) SetFileName(v string) *GetUploadInfoRequest {
	s.FileName = &v
	return s
}

func (s *GetUploadInfoRequest) SetFileSize(v int64) *GetUploadInfoRequest {
	s.FileSize = &v
	return s
}

func (s *GetUploadInfoRequest) SetMd5(v string) *GetUploadInfoRequest {
	s.Md5 = &v
	return s
}

func (s *GetUploadInfoRequest) SetAddConflictPolicy(v string) *GetUploadInfoRequest {
	s.AddConflictPolicy = &v
	return s
}

func (s *GetUploadInfoRequest) SetMediaId(v string) *GetUploadInfoRequest {
	s.MediaId = &v
	return s
}

type GetUploadInfoResponseBody struct {
	StsUploadInfo             *GetUploadInfoResponseBodyStsUploadInfo             `json:"stsUploadInfo,omitempty" xml:"stsUploadInfo,omitempty" type:"Struct"`
	HeaderSignatureUploadInfo *GetUploadInfoResponseBodyHeaderSignatureUploadInfo `json:"headerSignatureUploadInfo,omitempty" xml:"headerSignatureUploadInfo,omitempty" type:"Struct"`
}

func (s GetUploadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadInfoResponseBody) SetStsUploadInfo(v *GetUploadInfoResponseBodyStsUploadInfo) *GetUploadInfoResponseBody {
	s.StsUploadInfo = v
	return s
}

func (s *GetUploadInfoResponseBody) SetHeaderSignatureUploadInfo(v *GetUploadInfoResponseBodyHeaderSignatureUploadInfo) *GetUploadInfoResponseBody {
	s.HeaderSignatureUploadInfo = v
	return s
}

type GetUploadInfoResponseBodyStsUploadInfo struct {
	// bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// endPoint
	EndPoint *string `json:"endPoint,omitempty" xml:"endPoint,omitempty"`
	// accessKeyId
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accessKeySecret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// accessToken
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// accessTokenExpiration
	AccessTokenExpirationMillis *int64 `json:"accessTokenExpirationMillis,omitempty" xml:"accessTokenExpirationMillis,omitempty"`
	// mediaId
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

func (s GetUploadInfoResponseBodyStsUploadInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUploadInfoResponseBodyStsUploadInfo) GoString() string {
	return s.String()
}

func (s *GetUploadInfoResponseBodyStsUploadInfo) SetBucket(v string) *GetUploadInfoResponseBodyStsUploadInfo {
	s.Bucket = &v
	return s
}

func (s *GetUploadInfoResponseBodyStsUploadInfo) SetEndPoint(v string) *GetUploadInfoResponseBodyStsUploadInfo {
	s.EndPoint = &v
	return s
}

func (s *GetUploadInfoResponseBodyStsUploadInfo) SetAccessKeyId(v string) *GetUploadInfoResponseBodyStsUploadInfo {
	s.AccessKeyId = &v
	return s
}

func (s *GetUploadInfoResponseBodyStsUploadInfo) SetAccessKeySecret(v string) *GetUploadInfoResponseBodyStsUploadInfo {
	s.AccessKeySecret = &v
	return s
}

func (s *GetUploadInfoResponseBodyStsUploadInfo) SetAccessToken(v string) *GetUploadInfoResponseBodyStsUploadInfo {
	s.AccessToken = &v
	return s
}

func (s *GetUploadInfoResponseBodyStsUploadInfo) SetAccessTokenExpirationMillis(v int64) *GetUploadInfoResponseBodyStsUploadInfo {
	s.AccessTokenExpirationMillis = &v
	return s
}

func (s *GetUploadInfoResponseBodyStsUploadInfo) SetMediaId(v string) *GetUploadInfoResponseBodyStsUploadInfo {
	s.MediaId = &v
	return s
}

type GetUploadInfoResponseBodyHeaderSignatureUploadInfo struct {
	// 上传地址
	ResourceUrl *string `json:"resourceUrl,omitempty" xml:"resourceUrl,omitempty"`
	// 过期秒数
	ExpirationSeconds *int32 `json:"expirationSeconds,omitempty" xml:"expirationSeconds,omitempty"`
	// header加签信息
	Headers map[string]interface{} `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s GetUploadInfoResponseBodyHeaderSignatureUploadInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUploadInfoResponseBodyHeaderSignatureUploadInfo) GoString() string {
	return s.String()
}

func (s *GetUploadInfoResponseBodyHeaderSignatureUploadInfo) SetResourceUrl(v string) *GetUploadInfoResponseBodyHeaderSignatureUploadInfo {
	s.ResourceUrl = &v
	return s
}

func (s *GetUploadInfoResponseBodyHeaderSignatureUploadInfo) SetExpirationSeconds(v int32) *GetUploadInfoResponseBodyHeaderSignatureUploadInfo {
	s.ExpirationSeconds = &v
	return s
}

func (s *GetUploadInfoResponseBodyHeaderSignatureUploadInfo) SetHeaders(v map[string]interface{}) *GetUploadInfoResponseBodyHeaderSignatureUploadInfo {
	s.Headers = v
	return s
}

type GetUploadInfoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUploadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUploadInfoResponse) SetHeaders(v map[string]*string) *GetUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUploadInfoResponse) SetBody(v *GetUploadInfoResponseBody) *GetUploadInfoResponse {
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
	// 权限角色
	Role    *string                           `json:"role,omitempty" xml:"role,omitempty"`
	Members []*DeletePermissionRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeletePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionRequest) GoString() string {
	return s.String()
}

func (s *DeletePermissionRequest) SetRole(v string) *DeletePermissionRequest {
	s.Role = &v
	return s
}

func (s *DeletePermissionRequest) SetMembers(v []*DeletePermissionRequestMembers) *DeletePermissionRequest {
	s.Members = v
	return s
}

func (s *DeletePermissionRequest) SetUnionId(v string) *DeletePermissionRequest {
	s.UnionId = &v
	return s
}

type DeletePermissionRequestMembers struct {
	// 企业corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 成员类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 成员id
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
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

func (s *DeletePermissionRequestMembers) SetMemberType(v string) *DeletePermissionRequestMembers {
	s.MemberType = &v
	return s
}

func (s *DeletePermissionRequestMembers) SetMemberId(v string) *DeletePermissionRequestMembers {
	s.MemberId = &v
	return s
}

type DeletePermissionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
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

type AddCustomSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddCustomSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCustomSpaceHeaders) GoString() string {
	return s.String()
}

func (s *AddCustomSpaceHeaders) SetCommonHeaders(v map[string]*string) *AddCustomSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCustomSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *AddCustomSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddCustomSpaceRequest struct {
	// 空间标识
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 业务类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 授权模式
	PermissionMode *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddCustomSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomSpaceRequest) GoString() string {
	return s.String()
}

func (s *AddCustomSpaceRequest) SetIdentifier(v string) *AddCustomSpaceRequest {
	s.Identifier = &v
	return s
}

func (s *AddCustomSpaceRequest) SetBizType(v string) *AddCustomSpaceRequest {
	s.BizType = &v
	return s
}

func (s *AddCustomSpaceRequest) SetPermissionMode(v string) *AddCustomSpaceRequest {
	s.PermissionMode = &v
	return s
}

func (s *AddCustomSpaceRequest) SetUnionId(v string) *AddCustomSpaceRequest {
	s.UnionId = &v
	return s
}

type AddCustomSpaceResponseBody struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 空间名称
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	// 空间类型
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 空间总额度
	Quota *int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 空间已使用额度
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	// 授权模式
	PermissionMode *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
}

func (s AddCustomSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCustomSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomSpaceResponseBody) SetSpaceId(v string) *AddCustomSpaceResponseBody {
	s.SpaceId = &v
	return s
}

func (s *AddCustomSpaceResponseBody) SetSpaceName(v string) *AddCustomSpaceResponseBody {
	s.SpaceName = &v
	return s
}

func (s *AddCustomSpaceResponseBody) SetSpaceType(v string) *AddCustomSpaceResponseBody {
	s.SpaceType = &v
	return s
}

func (s *AddCustomSpaceResponseBody) SetQuota(v int64) *AddCustomSpaceResponseBody {
	s.Quota = &v
	return s
}

func (s *AddCustomSpaceResponseBody) SetUsedQuota(v int64) *AddCustomSpaceResponseBody {
	s.UsedQuota = &v
	return s
}

func (s *AddCustomSpaceResponseBody) SetPermissionMode(v string) *AddCustomSpaceResponseBody {
	s.PermissionMode = &v
	return s
}

func (s *AddCustomSpaceResponseBody) SetCreateTime(v string) *AddCustomSpaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *AddCustomSpaceResponseBody) SetModifyTime(v string) *AddCustomSpaceResponseBody {
	s.ModifyTime = &v
	return s
}

type AddCustomSpaceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddCustomSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCustomSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCustomSpaceResponse) GoString() string {
	return s.String()
}

func (s *AddCustomSpaceResponse) SetHeaders(v map[string]*string) *AddCustomSpaceResponse {
	s.Headers = v
	return s
}

func (s *AddCustomSpaceResponse) SetBody(v *AddCustomSpaceResponseBody) *AddCustomSpaceResponse {
	s.Body = v
	return s
}

type MoveFilesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MoveFilesHeaders) String() string {
	return tea.Prettify(s)
}

func (s MoveFilesHeaders) GoString() string {
	return s.String()
}

func (s *MoveFilesHeaders) SetCommonHeaders(v map[string]*string) *MoveFilesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MoveFilesHeaders) SetXAcsDingtalkAccessToken(v string) *MoveFilesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MoveFilesRequest struct {
	// 文件id列表
	FileIds []*string `json:"fileIds,omitempty" xml:"fileIds,omitempty" type:"Repeated"`
	// 目标空间id
	TargetSpaceId *string `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// 目标父目录id
	TargetParentId *string `json:"targetParentId,omitempty" xml:"targetParentId,omitempty"`
	// 文件名冲突策略
	AddConflictPolicy *string `json:"addConflictPolicy,omitempty" xml:"addConflictPolicy,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s MoveFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveFilesRequest) GoString() string {
	return s.String()
}

func (s *MoveFilesRequest) SetFileIds(v []*string) *MoveFilesRequest {
	s.FileIds = v
	return s
}

func (s *MoveFilesRequest) SetTargetSpaceId(v string) *MoveFilesRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *MoveFilesRequest) SetTargetParentId(v string) *MoveFilesRequest {
	s.TargetParentId = &v
	return s
}

func (s *MoveFilesRequest) SetAddConflictPolicy(v string) *MoveFilesRequest {
	s.AddConflictPolicy = &v
	return s
}

func (s *MoveFilesRequest) SetUnionId(v string) *MoveFilesRequest {
	s.UnionId = &v
	return s
}

type MoveFilesResponseBody struct {
	// 文件信息列表
	Files []*MoveFilesResponseBodyFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// 异步任务id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s MoveFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveFilesResponseBody) GoString() string {
	return s.String()
}

func (s *MoveFilesResponseBody) SetFiles(v []*MoveFilesResponseBodyFiles) *MoveFilesResponseBody {
	s.Files = v
	return s
}

func (s *MoveFilesResponseBody) SetTaskId(v string) *MoveFilesResponseBody {
	s.TaskId = &v
	return s
}

type MoveFilesResponseBodyFiles struct {
	// 空间id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 父目录id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件id
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件路径
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 文件类型
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 文件内容类型
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改者
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s MoveFilesResponseBodyFiles) String() string {
	return tea.Prettify(s)
}

func (s MoveFilesResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *MoveFilesResponseBodyFiles) SetSpaceId(v string) *MoveFilesResponseBodyFiles {
	s.SpaceId = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetParentId(v string) *MoveFilesResponseBodyFiles {
	s.ParentId = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetFileId(v string) *MoveFilesResponseBodyFiles {
	s.FileId = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetFileName(v string) *MoveFilesResponseBodyFiles {
	s.FileName = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetFilePath(v string) *MoveFilesResponseBodyFiles {
	s.FilePath = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetFileType(v string) *MoveFilesResponseBodyFiles {
	s.FileType = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetContentType(v string) *MoveFilesResponseBodyFiles {
	s.ContentType = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetFileExtension(v string) *MoveFilesResponseBodyFiles {
	s.FileExtension = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetFileSize(v int64) *MoveFilesResponseBodyFiles {
	s.FileSize = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetCreateTime(v string) *MoveFilesResponseBodyFiles {
	s.CreateTime = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetModifyTime(v string) *MoveFilesResponseBodyFiles {
	s.ModifyTime = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetCreator(v string) *MoveFilesResponseBodyFiles {
	s.Creator = &v
	return s
}

func (s *MoveFilesResponseBodyFiles) SetModifier(v string) *MoveFilesResponseBodyFiles {
	s.Modifier = &v
	return s
}

type MoveFilesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveFilesResponse) GoString() string {
	return s.String()
}

func (s *MoveFilesResponse) SetHeaders(v map[string]*string) *MoveFilesResponse {
	s.Headers = v
	return s
}

func (s *MoveFilesResponse) SetBody(v *MoveFilesResponseBody) *MoveFilesResponse {
	s.Body = v
	return s
}

type GetPrivilegeInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPrivilegeInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPrivilegeInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetPrivilegeInfoHeaders) SetCommonHeaders(v map[string]*string) *GetPrivilegeInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPrivilegeInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetPrivilegeInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPrivilegeInfoRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetPrivilegeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrivilegeInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPrivilegeInfoRequest) SetUnionId(v string) *GetPrivilegeInfoRequest {
	s.UnionId = &v
	return s
}

type GetPrivilegeInfoResponseBody struct {
	// 类型列表
	Types []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s GetPrivilegeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrivilegeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrivilegeInfoResponseBody) SetTypes(v []*string) *GetPrivilegeInfoResponseBody {
	s.Types = v
	return s
}

type GetPrivilegeInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPrivilegeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPrivilegeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrivilegeInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPrivilegeInfoResponse) SetHeaders(v map[string]*string) *GetPrivilegeInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPrivilegeInfoResponse) SetBody(v *GetPrivilegeInfoResponseBody) *GetPrivilegeInfoResponse {
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

func (client *Client) RecoverRecycleFiles(request *RecoverRecycleFilesRequest) (_result *RecoverRecycleFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecoverRecycleFilesHeaders{}
	_result = &RecoverRecycleFilesResponse{}
	_body, _err := client.RecoverRecycleFilesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoverRecycleFilesWithOptions(request *RecoverRecycleFilesRequest, headers *RecoverRecycleFilesHeaders, runtime *util.RuntimeOptions) (_result *RecoverRecycleFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecycleItemIdList)) {
		body["recycleItemIdList"] = request.RecycleItemIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RecycleType)) {
		body["recycleType"] = request.RecycleType
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RecoverRecycleFilesResponse{}
	_body, _err := client.DoROARequest(tea.String("RecoverRecycleFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/recycleItems/recover"), tea.String("none"), req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("AddSpace"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRecycleFiles(request *DeleteRecycleFilesRequest) (_result *DeleteRecycleFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRecycleFilesHeaders{}
	_result = &DeleteRecycleFilesResponse{}
	_body, _err := client.DeleteRecycleFilesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRecycleFilesWithOptions(request *DeleteRecycleFilesRequest, headers *DeleteRecycleFilesHeaders, runtime *util.RuntimeOptions) (_result *DeleteRecycleFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecycleItemIdList)) {
		body["recycleItemIdList"] = request.RecycleItemIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RecycleType)) {
		body["recycleType"] = request.RecycleType
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &DeleteRecycleFilesResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRecycleFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/recycleItems/delete"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPermission(spaceId *string, fileId *string, request *AddPermissionRequest) (_result *AddPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddPermissionHeaders{}
	_result = &AddPermissionResponse{}
	_body, _err := client.AddPermissionWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddPermissionWithOptions(spaceId *string, fileId *string, request *AddPermissionRequest, headers *AddPermissionHeaders, runtime *util.RuntimeOptions) (_result *AddPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddPermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("AddPermission"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/permissions"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQuotaInfos(request *GetQuotaInfosRequest) (_result *GetQuotaInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetQuotaInfosHeaders{}
	_result = &GetQuotaInfosResponse{}
	_body, _err := client.GetQuotaInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQuotaInfosWithOptions(request *GetQuotaInfosRequest, headers *GetQuotaInfosHeaders, runtime *util.RuntimeOptions) (_result *GetQuotaInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Identifiers)) {
		body["identifiers"] = request.Identifiers
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetQuotaInfosResponse{}
	_body, _err := client.DoROARequest(tea.String("GetQuotaInfos"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/quotaInfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileInfo(spaceId *string, fileId *string, request *GetFileInfoRequest) (_result *GetFileInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileInfoHeaders{}
	_result = &GetFileInfoResponse{}
	_body, _err := client.GetFileInfoWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileInfoWithOptions(spaceId *string, fileId *string, request *GetFileInfoRequest, headers *GetFileInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileInfoResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetFileInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFileInfo"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecycleFiles(request *ListRecycleFilesRequest) (_result *ListRecycleFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRecycleFilesHeaders{}
	_result = &ListRecycleFilesResponse{}
	_body, _err := client.ListRecycleFilesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecycleFilesWithOptions(request *ListRecycleFilesRequest, headers *ListRecycleFilesHeaders, runtime *util.RuntimeOptions) (_result *ListRecycleFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.RecycleType)) {
		query["recycleType"] = request.RecycleType
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["orderType"] = request.OrderType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ListRecycleFilesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRecycleFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/recycleItems"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFiles(spaceId *string, request *DeleteFilesRequest) (_result *DeleteFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteFilesHeaders{}
	_result = &DeleteFilesResponse{}
	_body, _err := client.DeleteFilesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFilesWithOptions(spaceId *string, request *DeleteFilesRequest, headers *DeleteFilesHeaders, runtime *util.RuntimeOptions) (_result *DeleteFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileIds)) {
		body["fileIds"] = request.FileIds
	}

	if !tea.BoolValue(util.IsUnset(request.DeletePolicy)) {
		body["deletePolicy"] = request.DeletePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &DeleteFilesResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/batchDelete"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ManagementBuyQuota(request *ManagementBuyQuotaRequest) (_result *ManagementBuyQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ManagementBuyQuotaHeaders{}
	_result = &ManagementBuyQuotaResponse{}
	_body, _err := client.ManagementBuyQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ManagementBuyQuotaWithOptions(request *ManagementBuyQuotaRequest, headers *ManagementBuyQuotaHeaders, runtime *util.RuntimeOptions) (_result *ManagementBuyQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Order))) {
		body["order"] = request.Order
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ManagementBuyQuotaResponse{}
	_body, _err := client.DoROARequest(tea.String("ManagementBuyQuota"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/managements/quotas/buy"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameFile(spaceId *string, fileId *string, request *RenameFileRequest) (_result *RenameFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RenameFileHeaders{}
	_result = &RenameFileResponse{}
	_body, _err := client.RenameFileWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameFileWithOptions(spaceId *string, fileId *string, request *RenameFileRequest, headers *RenameFileHeaders, runtime *util.RuntimeOptions) (_result *RenameFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewFileName)) {
		body["newFileName"] = request.NewFileName
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RenameFileResponse{}
	_body, _err := client.DoROARequest(tea.String("RenameFile"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/rename"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncTaskInfo(taskId *string, request *GetAsyncTaskInfoRequest) (_result *GetAsyncTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAsyncTaskInfoHeaders{}
	_result = &GetAsyncTaskInfoResponse{}
	_body, _err := client.GetAsyncTaskInfoWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncTaskInfoWithOptions(taskId *string, request *GetAsyncTaskInfoRequest, headers *GetAsyncTaskInfoHeaders, runtime *util.RuntimeOptions) (_result *GetAsyncTaskInfoResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetAsyncTaskInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAsyncTaskInfo"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/tasks/"+tea.StringValue(taskId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFiles(spaceId *string, request *ListFilesRequest) (_result *ListFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListFilesHeaders{}
	_result = &ListFilesResponse{}
	_body, _err := client.ListFilesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFilesWithOptions(spaceId *string, request *ListFilesRequest, headers *ListFilesHeaders, runtime *util.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["parentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["orderType"] = request.OrderType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ListFilesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPermissions(spaceId *string, fileId *string, request *ListPermissionsRequest) (_result *ListPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListPermissionsHeaders{}
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPermissionsWithOptions(spaceId *string, fileId *string, request *ListPermissionsRequest, headers *ListPermissionsHeaders, runtime *util.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ListPermissionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListPermissions"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/permissions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveFile(spaceId *string, fileId *string, request *MoveFileRequest) (_result *MoveFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MoveFileHeaders{}
	_result = &MoveFileResponse{}
	_body, _err := client.MoveFileWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveFileWithOptions(spaceId *string, fileId *string, request *MoveFileRequest, headers *MoveFileHeaders, runtime *util.RuntimeOptions) (_result *MoveFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetSpaceId)) {
		body["targetSpaceId"] = request.TargetSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetParentId)) {
		body["targetParentId"] = request.TargetParentId
	}

	if !tea.BoolValue(util.IsUnset(request.AddConflictPolicy)) {
		body["addConflictPolicy"] = request.AddConflictPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &MoveFileResponse{}
	_body, _err := client.DoROARequest(tea.String("MoveFile"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/move"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSpaces(request *ListSpacesRequest) (_result *ListSpacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSpacesHeaders{}
	_result = &ListSpacesResponse{}
	_body, _err := client.ListSpacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSpacesWithOptions(request *ListSpacesRequest, headers *ListSpacesHeaders, runtime *util.RuntimeOptions) (_result *ListSpacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceType)) {
		query["spaceType"] = request.SpaceType
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ListSpacesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSpaces"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/spaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyFile(spaceId *string, fileId *string, request *CopyFileRequest) (_result *CopyFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CopyFileHeaders{}
	_result = &CopyFileResponse{}
	_body, _err := client.CopyFileWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyFileWithOptions(spaceId *string, fileId *string, request *CopyFileRequest, headers *CopyFileHeaders, runtime *util.RuntimeOptions) (_result *CopyFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetSpaceId)) {
		body["targetSpaceId"] = request.TargetSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetParentId)) {
		body["targetParentId"] = request.TargetParentId
	}

	if !tea.BoolValue(util.IsUnset(request.AddConflictPolicy)) {
		body["addConflictPolicy"] = request.AddConflictPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CopyFileResponse{}
	_body, _err := client.DoROARequest(tea.String("CopyFile"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/copy"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSpace(spaceId *string, request *DeleteSpaceRequest) (_result *DeleteSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSpaceHeaders{}
	_result = &DeleteSpaceResponse{}
	_body, _err := client.DeleteSpaceWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSpaceWithOptions(spaceId *string, request *DeleteSpaceRequest, headers *DeleteSpaceHeaders, runtime *util.RuntimeOptions) (_result *DeleteSpaceResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSpace"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearRecycleFiles(request *ClearRecycleFilesRequest) (_result *ClearRecycleFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ClearRecycleFilesHeaders{}
	_result = &ClearRecycleFilesResponse{}
	_body, _err := client.ClearRecycleFilesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearRecycleFilesWithOptions(request *ClearRecycleFilesRequest, headers *ClearRecycleFilesHeaders, runtime *util.RuntimeOptions) (_result *ClearRecycleFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecycleType)) {
		body["recycleType"] = request.RecycleType
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ClearRecycleFilesResponse{}
	_body, _err := client.DoROARequest(tea.String("ClearRecycleFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/recycleItems/clear"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ManagementListSpaces(request *ManagementListSpacesRequest) (_result *ManagementListSpacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ManagementListSpacesHeaders{}
	_result = &ManagementListSpacesResponse{}
	_body, _err := client.ManagementListSpacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ManagementListSpacesWithOptions(request *ManagementListSpacesRequest, headers *ManagementListSpacesHeaders, runtime *util.RuntimeOptions) (_result *ManagementListSpacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceIds)) {
		body["spaceIds"] = request.SpaceIds
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ManagementListSpacesResponse{}
	_body, _err := client.DoROARequest(tea.String("ManagementListSpaces"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/managements/spaces/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFile(spaceId *string, fileId *string, request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteFileHeaders{}
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileWithOptions(spaceId *string, fileId *string, request *DeleteFileRequest, headers *DeleteFileHeaders, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.DeletePolicy)) {
		query["deletePolicy"] = request.DeletePolicy
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteFile"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFile(spaceId *string, request *AddFileRequest) (_result *AddFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddFileHeaders{}
	_result = &AddFileResponse{}
	_body, _err := client.AddFileWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFileWithOptions(spaceId *string, request *AddFileRequest, headers *AddFileHeaders, runtime *util.RuntimeOptions) (_result *AddFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["parentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["fileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.AddConflictPolicy)) {
		body["addConflictPolicy"] = request.AddConflictPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddFileResponse{}
	_body, _err := client.DoROARequest(tea.String("AddFile"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPreviewInfo(spaceId *string, fileId *string, request *GetPreviewInfoRequest) (_result *GetPreviewInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPreviewInfoHeaders{}
	_result = &GetPreviewInfoResponse{}
	_body, _err := client.GetPreviewInfoWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPreviewInfoWithOptions(spaceId *string, fileId *string, request *GetPreviewInfoRequest, headers *GetPreviewInfoHeaders, runtime *util.RuntimeOptions) (_result *GetPreviewInfoResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetPreviewInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPreviewInfo"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/previewInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InfoSpace(spaceId *string, request *InfoSpaceRequest) (_result *InfoSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InfoSpaceHeaders{}
	_result = &InfoSpaceResponse{}
	_body, _err := client.InfoSpaceWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InfoSpaceWithOptions(spaceId *string, request *InfoSpaceRequest, headers *InfoSpaceHeaders, runtime *util.RuntimeOptions) (_result *InfoSpaceResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &InfoSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("InfoSpace"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ManagementModifySpace(request *ManagementModifySpaceRequest) (_result *ManagementModifySpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ManagementModifySpaceHeaders{}
	_result = &ManagementModifySpaceResponse{}
	_body, _err := client.ManagementModifySpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ManagementModifySpaceWithOptions(request *ManagementModifySpaceRequest, headers *ManagementModifySpaceHeaders, runtime *util.RuntimeOptions) (_result *ManagementModifySpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceIds)) {
		body["spaceIds"] = request.SpaceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Quota)) {
		body["quota"] = request.Quota
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ManagementModifySpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("ManagementModifySpace"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/drive/managements/spaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPermission(spaceId *string, fileId *string, request *ModifyPermissionRequest) (_result *ModifyPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ModifyPermissionHeaders{}
	_result = &ModifyPermissionResponse{}
	_body, _err := client.ModifyPermissionWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPermissionWithOptions(spaceId *string, fileId *string, request *ModifyPermissionRequest, headers *ModifyPermissionHeaders, runtime *util.RuntimeOptions) (_result *ModifyPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ModifyPermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyPermission"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/permissions"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantPrivilegeOfCustomSpace(spaceId *string, request *GrantPrivilegeOfCustomSpaceRequest) (_result *GrantPrivilegeOfCustomSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GrantPrivilegeOfCustomSpaceHeaders{}
	_result = &GrantPrivilegeOfCustomSpaceResponse{}
	_body, _err := client.GrantPrivilegeOfCustomSpaceWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantPrivilegeOfCustomSpaceWithOptions(spaceId *string, request *GrantPrivilegeOfCustomSpaceRequest, headers *GrantPrivilegeOfCustomSpaceHeaders, runtime *util.RuntimeOptions) (_result *GrantPrivilegeOfCustomSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.FileIds)) {
		body["fileIds"] = request.FileIds
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GrantPrivilegeOfCustomSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("GrantPrivilegeOfCustomSpace"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/customSpacePrivileges"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDownloadInfo(spaceId *string, fileId *string, request *GetDownloadInfoRequest) (_result *GetDownloadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDownloadInfoHeaders{}
	_result = &GetDownloadInfoResponse{}
	_body, _err := client.GetDownloadInfoWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDownloadInfoWithOptions(spaceId *string, fileId *string, request *GetDownloadInfoRequest, headers *GetDownloadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetDownloadInfoResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetDownloadInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDownloadInfo"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/downloadInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUploadInfo(spaceId *string, parentId *string, request *GetUploadInfoRequest) (_result *GetUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUploadInfoHeaders{}
	_result = &GetUploadInfoResponse{}
	_body, _err := client.GetUploadInfoWithOptions(spaceId, parentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUploadInfoWithOptions(spaceId *string, parentId *string, request *GetUploadInfoRequest, headers *GetUploadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		query["fileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.Md5)) {
		query["md5"] = request.Md5
	}

	if !tea.BoolValue(util.IsUnset(request.AddConflictPolicy)) {
		query["addConflictPolicy"] = request.AddConflictPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		query["mediaId"] = request.MediaId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetUploadInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUploadInfo"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(parentId)+"/uploadInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePermission(spaceId *string, fileId *string, request *DeletePermissionRequest) (_result *DeletePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeletePermissionHeaders{}
	_result = &DeletePermissionResponse{}
	_body, _err := client.DeletePermissionWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePermissionWithOptions(spaceId *string, fileId *string, request *DeletePermissionRequest, headers *DeletePermissionHeaders, runtime *util.RuntimeOptions) (_result *DeletePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &DeletePermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("DeletePermission"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/permissions/delete"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCustomSpace(request *AddCustomSpaceRequest) (_result *AddCustomSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCustomSpaceHeaders{}
	_result = &AddCustomSpaceResponse{}
	_body, _err := client.AddCustomSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCustomSpaceWithOptions(request *AddCustomSpaceRequest, headers *AddCustomSpaceHeaders, runtime *util.RuntimeOptions) (_result *AddCustomSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		body["identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionMode)) {
		body["permissionMode"] = request.PermissionMode
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddCustomSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("AddCustomSpace"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces/customSpaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveFiles(spaceId *string, request *MoveFilesRequest) (_result *MoveFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MoveFilesHeaders{}
	_result = &MoveFilesResponse{}
	_body, _err := client.MoveFilesWithOptions(spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveFilesWithOptions(spaceId *string, request *MoveFilesRequest, headers *MoveFilesHeaders, runtime *util.RuntimeOptions) (_result *MoveFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileIds)) {
		body["fileIds"] = request.FileIds
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSpaceId)) {
		body["targetSpaceId"] = request.TargetSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetParentId)) {
		body["targetParentId"] = request.TargetParentId
	}

	if !tea.BoolValue(util.IsUnset(request.AddConflictPolicy)) {
		body["addConflictPolicy"] = request.AddConflictPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &MoveFilesResponse{}
	_body, _err := client.DoROARequest(tea.String("MoveFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/batchMove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPrivilegeInfo(spaceId *string, fileId *string, request *GetPrivilegeInfoRequest) (_result *GetPrivilegeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPrivilegeInfoHeaders{}
	_result = &GetPrivilegeInfoResponse{}
	_body, _err := client.GetPrivilegeInfoWithOptions(spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPrivilegeInfoWithOptions(spaceId *string, fileId *string, request *GetPrivilegeInfoRequest, headers *GetPrivilegeInfoHeaders, runtime *util.RuntimeOptions) (_result *GetPrivilegeInfoResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetPrivilegeInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPrivilegeInfo"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/privileges"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
