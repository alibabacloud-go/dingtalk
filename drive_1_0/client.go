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
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
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

func (s *AddFileResponseBody) SetFileExtension(v string) *AddFileResponseBody {
	s.FileExtension = &v
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
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
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

func (s *GetFileInfoResponseBody) SetFileExtension(v string) *GetFileInfoResponseBody {
	s.FileExtension = &v
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
	// 删除者id
	DeleteUid *int64 `json:"deleteUid,omitempty" xml:"deleteUid,omitempty"`
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

func (s *ListRecycleFilesResponseBodyRecycleItems) SetDeleteUid(v int64) *ListRecycleFilesResponseBodyRecycleItems {
	s.DeleteUid = &v
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
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
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

func (s *RenameFileResponseBody) SetFileExtension(v string) *RenameFileResponseBody {
	s.FileExtension = &v
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
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
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

func (s *ListFilesResponseBodyFiles) SetFileExtension(v string) *ListFilesResponseBodyFiles {
	s.FileExtension = &v
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
	// 文件后缀
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
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

func (s *MoveFileResponseBody) SetFileExtension(v string) *MoveFileResponseBody {
	s.FileExtension = &v
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
	// 删除策略
	DeletePolicy *string `json:"deletePolicy,omitempty" xml:"deletePolicy,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
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

func (client *Client) AddFile(unionId *string, spaceId *string, request *AddFileRequest) (_result *AddFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddFileHeaders{}
	_result = &AddFileResponse{}
	_body, _err := client.AddFileWithOptions(unionId, spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFileWithOptions(unionId *string, spaceId *string, request *AddFileRequest, headers *AddFileHeaders, runtime *util.RuntimeOptions) (_result *AddFileResponse, _err error) {
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
	_body, _err := client.DoROARequest(tea.String("AddFile"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/spaces/"+tea.StringValue(spaceId)+"/files"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecoverRecycleFiles(unionId *string, request *RecoverRecycleFilesRequest) (_result *RecoverRecycleFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecoverRecycleFilesHeaders{}
	_result = &RecoverRecycleFilesResponse{}
	_body, _err := client.RecoverRecycleFilesWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoverRecycleFilesWithOptions(unionId *string, request *RecoverRecycleFilesRequest, headers *RecoverRecycleFilesHeaders, runtime *util.RuntimeOptions) (_result *RecoverRecycleFilesResponse, _err error) {
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
	_body, _err := client.DoROARequest(tea.String("RecoverRecycleFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/recycleItems/recover"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRecycleFiles(unionId *string, request *DeleteRecycleFilesRequest) (_result *DeleteRecycleFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRecycleFilesHeaders{}
	_result = &DeleteRecycleFilesResponse{}
	_body, _err := client.DeleteRecycleFilesWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRecycleFilesWithOptions(unionId *string, request *DeleteRecycleFilesRequest, headers *DeleteRecycleFilesHeaders, runtime *util.RuntimeOptions) (_result *DeleteRecycleFilesResponse, _err error) {
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
	_body, _err := client.DoROARequest(tea.String("DeleteRecycleFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/recycleItems/delete"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileInfo(unionId *string, spaceId *string, fileId *string) (_result *GetFileInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFileInfoHeaders{}
	_result = &GetFileInfoResponse{}
	_body, _err := client.GetFileInfoWithOptions(unionId, spaceId, fileId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileInfoWithOptions(unionId *string, spaceId *string, fileId *string, headers *GetFileInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFileInfoResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetFileInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFileInfo"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecycleFiles(unionId *string, request *ListRecycleFilesRequest) (_result *ListRecycleFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRecycleFilesHeaders{}
	_result = &ListRecycleFilesResponse{}
	_body, _err := client.ListRecycleFilesWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecycleFilesWithOptions(unionId *string, request *ListRecycleFilesRequest, headers *ListRecycleFilesHeaders, runtime *util.RuntimeOptions) (_result *ListRecycleFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	_body, _err := client.DoROARequest(tea.String("ListRecycleFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/recycleItems"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameFile(unionId *string, spaceId *string, fileId *string, request *RenameFileRequest) (_result *RenameFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RenameFileHeaders{}
	_result = &RenameFileResponse{}
	_body, _err := client.RenameFileWithOptions(unionId, spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameFileWithOptions(unionId *string, spaceId *string, fileId *string, request *RenameFileRequest, headers *RenameFileHeaders, runtime *util.RuntimeOptions) (_result *RenameFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewFileName)) {
		body["newFileName"] = request.NewFileName
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
	_body, _err := client.DoROARequest(tea.String("RenameFile"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/rename"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFiles(unionId *string, spaceId *string, request *ListFilesRequest) (_result *ListFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListFilesHeaders{}
	_result = &ListFilesResponse{}
	_body, _err := client.ListFilesWithOptions(unionId, spaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFilesWithOptions(unionId *string, spaceId *string, request *ListFilesRequest, headers *ListFilesHeaders, runtime *util.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	_body, _err := client.DoROARequest(tea.String("ListFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/spaces/"+tea.StringValue(spaceId)+"/files"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveFile(unionId *string, spaceId *string, fileId *string, request *MoveFileRequest) (_result *MoveFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MoveFileHeaders{}
	_result = &MoveFileResponse{}
	_body, _err := client.MoveFileWithOptions(unionId, spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveFileWithOptions(unionId *string, spaceId *string, fileId *string, request *MoveFileRequest, headers *MoveFileHeaders, runtime *util.RuntimeOptions) (_result *MoveFileResponse, _err error) {
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
	_body, _err := client.DoROARequest(tea.String("MoveFile"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/move"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDownloadInfo(unionId *string, spaceId *string, fileId *string) (_result *GetDownloadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDownloadInfoHeaders{}
	_result = &GetDownloadInfoResponse{}
	_body, _err := client.GetDownloadInfoWithOptions(unionId, spaceId, fileId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDownloadInfoWithOptions(unionId *string, spaceId *string, fileId *string, headers *GetDownloadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetDownloadInfoResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetDownloadInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDownloadInfo"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)+"/downloadInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUploadInfo(unionId *string, spaceId *string, parentId *string, request *GetUploadInfoRequest) (_result *GetUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUploadInfoHeaders{}
	_result = &GetUploadInfoResponse{}
	_body, _err := client.GetUploadInfoWithOptions(unionId, spaceId, parentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUploadInfoWithOptions(unionId *string, spaceId *string, parentId *string, request *GetUploadInfoRequest, headers *GetUploadInfoHeaders, runtime *util.RuntimeOptions) (_result *GetUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	_body, _err := client.DoROARequest(tea.String("GetUploadInfo"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(parentId)+"/uploadInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSpaces(unionId *string, request *ListSpacesRequest) (_result *ListSpacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSpacesHeaders{}
	_result = &ListSpacesResponse{}
	_body, _err := client.ListSpacesWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSpacesWithOptions(unionId *string, request *ListSpacesRequest, headers *ListSpacesHeaders, runtime *util.RuntimeOptions) (_result *ListSpacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	_body, _err := client.DoROARequest(tea.String("ListSpaces"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/spaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearRecycleFiles(unionId *string, request *ClearRecycleFilesRequest) (_result *ClearRecycleFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ClearRecycleFilesHeaders{}
	_result = &ClearRecycleFilesResponse{}
	_body, _err := client.ClearRecycleFilesWithOptions(unionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearRecycleFilesWithOptions(unionId *string, request *ClearRecycleFilesRequest, headers *ClearRecycleFilesHeaders, runtime *util.RuntimeOptions) (_result *ClearRecycleFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecycleType)) {
		body["recycleType"] = request.RecycleType
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
	_body, _err := client.DoROARequest(tea.String("ClearRecycleFiles"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/recycleItems/clear"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFile(unionId *string, spaceId *string, fileId *string, request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteFileHeaders{}
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(unionId, spaceId, fileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileWithOptions(unionId *string, spaceId *string, fileId *string, request *DeleteFileRequest, headers *DeleteFileHeaders, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	_body, _err := client.DoROARequest(tea.String("DeleteFile"), tea.String("drive_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/drive/users/"+tea.StringValue(unionId)+"/spaces/"+tea.StringValue(spaceId)+"/files/"+tea.StringValue(fileId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
