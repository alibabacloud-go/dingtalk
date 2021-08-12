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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
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

func (s *AddFileResponseBody) SetCreateTime(v string) *AddFileResponseBody {
	s.CreateTime = &v
	return s
}

func (s *AddFileResponseBody) SetModifyTime(v string) *AddFileResponseBody {
	s.ModifyTime = &v
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
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

func (s *GetFileInfoResponseBody) SetCreateTime(v string) *GetFileInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetFileInfoResponseBody) SetModifyTime(v string) *GetFileInfoResponseBody {
	s.ModifyTime = &v
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
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

func (s *RenameFileResponseBody) SetCreateTime(v string) *RenameFileResponseBody {
	s.CreateTime = &v
	return s
}

func (s *RenameFileResponseBody) SetModifyTime(v string) *RenameFileResponseBody {
	s.ModifyTime = &v
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
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

func (s *ListFilesResponseBodyFiles) SetCreateTime(v string) *ListFilesResponseBodyFiles {
	s.CreateTime = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetModifyTime(v string) *ListFilesResponseBodyFiles {
	s.ModifyTime = &v
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
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
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

func (s *MoveFileResponseBody) SetCreateTime(v string) *MoveFileResponseBody {
	s.CreateTime = &v
	return s
}

func (s *MoveFileResponseBody) SetModifyTime(v string) *MoveFileResponseBody {
	s.ModifyTime = &v
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
	StsUploadInfo *GetUploadInfoResponseBodyStsUploadInfo `json:"stsUploadInfo,omitempty" xml:"stsUploadInfo,omitempty" type:"Struct"`
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
