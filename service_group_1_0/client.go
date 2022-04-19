// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package service_group_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddKnowledgeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddKnowledgeHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddKnowledgeHeaders) GoString() string {
	return s.String()
}

func (s *AddKnowledgeHeaders) SetCommonHeaders(v map[string]*string) *AddKnowledgeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddKnowledgeHeaders) SetXAcsDingtalkAccessToken(v string) *AddKnowledgeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddKnowledgeRequest struct {
	// 附件列表
	AttachmentList []*AddKnowledgeRequestAttachmentList `json:"attachmentList,omitempty" xml:"attachmentList,omitempty" type:"Repeated"`
	// 知识点内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 知识点扩展问(多个用英文逗号隔开)
	ExtTitle *string `json:"extTitle,omitempty" xml:"extTitle,omitempty"`
	// 关键字(多个用英文逗号隔开)
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 知识库的唯一标识
	LibraryKey *string `json:"libraryKey,omitempty" xml:"libraryKey,omitempty"`
	// CCM的知识点外链
	LinkUrl *string `json:"linkUrl,omitempty" xml:"linkUrl,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 知识点来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 知识点唯一标识
	SourcePrimaryKey *string `json:"sourcePrimaryKey,omitempty" xml:"sourcePrimaryKey,omitempty"`
	// 知识点名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 知识点类型 NORMAL：普通型 CARD：卡片 CONDITION：条件
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 知识点版本号
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AddKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *AddKnowledgeRequest) SetAttachmentList(v []*AddKnowledgeRequestAttachmentList) *AddKnowledgeRequest {
	s.AttachmentList = v
	return s
}

func (s *AddKnowledgeRequest) SetContent(v string) *AddKnowledgeRequest {
	s.Content = &v
	return s
}

func (s *AddKnowledgeRequest) SetExtTitle(v string) *AddKnowledgeRequest {
	s.ExtTitle = &v
	return s
}

func (s *AddKnowledgeRequest) SetKeyword(v string) *AddKnowledgeRequest {
	s.Keyword = &v
	return s
}

func (s *AddKnowledgeRequest) SetLibraryKey(v string) *AddKnowledgeRequest {
	s.LibraryKey = &v
	return s
}

func (s *AddKnowledgeRequest) SetLinkUrl(v string) *AddKnowledgeRequest {
	s.LinkUrl = &v
	return s
}

func (s *AddKnowledgeRequest) SetOpenTeamId(v string) *AddKnowledgeRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AddKnowledgeRequest) SetSource(v string) *AddKnowledgeRequest {
	s.Source = &v
	return s
}

func (s *AddKnowledgeRequest) SetSourcePrimaryKey(v string) *AddKnowledgeRequest {
	s.SourcePrimaryKey = &v
	return s
}

func (s *AddKnowledgeRequest) SetTitle(v string) *AddKnowledgeRequest {
	s.Title = &v
	return s
}

func (s *AddKnowledgeRequest) SetType(v string) *AddKnowledgeRequest {
	s.Type = &v
	return s
}

func (s *AddKnowledgeRequest) SetVersion(v string) *AddKnowledgeRequest {
	s.Version = &v
	return s
}

type AddKnowledgeRequestAttachmentList struct {
	// 多媒体类型
	MimeType *string `json:"mime_type,omitempty" xml:"mime_type,omitempty"`
	// 附件URL
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 附件大小
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 附件扩展名
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
	// 附件名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s AddKnowledgeRequestAttachmentList) String() string {
	return tea.Prettify(s)
}

func (s AddKnowledgeRequestAttachmentList) GoString() string {
	return s.String()
}

func (s *AddKnowledgeRequestAttachmentList) SetMimeType(v string) *AddKnowledgeRequestAttachmentList {
	s.MimeType = &v
	return s
}

func (s *AddKnowledgeRequestAttachmentList) SetPath(v string) *AddKnowledgeRequestAttachmentList {
	s.Path = &v
	return s
}

func (s *AddKnowledgeRequestAttachmentList) SetSize(v int64) *AddKnowledgeRequestAttachmentList {
	s.Size = &v
	return s
}

func (s *AddKnowledgeRequestAttachmentList) SetSuffix(v string) *AddKnowledgeRequestAttachmentList {
	s.Suffix = &v
	return s
}

func (s *AddKnowledgeRequestAttachmentList) SetTitle(v string) *AddKnowledgeRequestAttachmentList {
	s.Title = &v
	return s
}

type AddKnowledgeResponseBody struct {
	// 开放知识点ID
	OpenKnowledgeId *string `json:"openKnowledgeId,omitempty" xml:"openKnowledgeId,omitempty"`
}

func (s AddKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *AddKnowledgeResponseBody) SetOpenKnowledgeId(v string) *AddKnowledgeResponseBody {
	s.OpenKnowledgeId = &v
	return s
}

type AddKnowledgeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *AddKnowledgeResponse) SetHeaders(v map[string]*string) *AddKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *AddKnowledgeResponse) SetBody(v *AddKnowledgeResponseBody) *AddKnowledgeResponse {
	s.Body = v
	return s
}

type AddLibraryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddLibraryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddLibraryHeaders) GoString() string {
	return s.String()
}

func (s *AddLibraryHeaders) SetCommonHeaders(v map[string]*string) *AddLibraryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddLibraryHeaders) SetXAcsDingtalkAccessToken(v string) *AddLibraryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddLibraryRequest struct {
	// 知识库描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 团队id列表
	OpenTeamIds []*string `json:"openTeamIds,omitempty" xml:"openTeamIds,omitempty" type:"Repeated"`
	// 知识来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 知识库的唯一性标识
	SourcePrimaryKey *string `json:"sourcePrimaryKey,omitempty" xml:"sourcePrimaryKey,omitempty"`
	// 知识库名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 知识库类型 INTERNAL:内部知识库 EXTERNAL:外部知识库
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 员工ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLibraryRequest) GoString() string {
	return s.String()
}

func (s *AddLibraryRequest) SetDescription(v string) *AddLibraryRequest {
	s.Description = &v
	return s
}

func (s *AddLibraryRequest) SetOpenTeamIds(v []*string) *AddLibraryRequest {
	s.OpenTeamIds = v
	return s
}

func (s *AddLibraryRequest) SetSource(v string) *AddLibraryRequest {
	s.Source = &v
	return s
}

func (s *AddLibraryRequest) SetSourcePrimaryKey(v string) *AddLibraryRequest {
	s.SourcePrimaryKey = &v
	return s
}

func (s *AddLibraryRequest) SetTitle(v string) *AddLibraryRequest {
	s.Title = &v
	return s
}

func (s *AddLibraryRequest) SetType(v string) *AddLibraryRequest {
	s.Type = &v
	return s
}

func (s *AddLibraryRequest) SetUserId(v string) *AddLibraryRequest {
	s.UserId = &v
	return s
}

type AddLibraryResponseBody struct {
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *AddLibraryResponseBody) SetSuccess(v bool) *AddLibraryResponseBody {
	s.Success = &v
	return s
}

type AddLibraryResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLibraryResponse) GoString() string {
	return s.String()
}

func (s *AddLibraryResponse) SetHeaders(v map[string]*string) *AddLibraryResponse {
	s.Headers = v
	return s
}

func (s *AddLibraryResponse) SetBody(v *AddLibraryResponseBody) *AddLibraryResponse {
	s.Body = v
	return s
}

type AddOpenCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddOpenCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddOpenCategoryHeaders) GoString() string {
	return s.String()
}

func (s *AddOpenCategoryHeaders) SetCommonHeaders(v map[string]*string) *AddOpenCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOpenCategoryHeaders) SetXAcsDingtalkAccessToken(v string) *AddOpenCategoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddOpenCategoryRequest struct {
	// 所属知识库ID
	LibraryId *int64 `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 父类目ID(为0代表顶层id)
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 类目标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 员工/用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户昵称或姓名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s AddOpenCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOpenCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddOpenCategoryRequest) SetLibraryId(v int64) *AddOpenCategoryRequest {
	s.LibraryId = &v
	return s
}

func (s *AddOpenCategoryRequest) SetOpenTeamId(v string) *AddOpenCategoryRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AddOpenCategoryRequest) SetParentId(v int64) *AddOpenCategoryRequest {
	s.ParentId = &v
	return s
}

func (s *AddOpenCategoryRequest) SetTitle(v string) *AddOpenCategoryRequest {
	s.Title = &v
	return s
}

func (s *AddOpenCategoryRequest) SetUserId(v string) *AddOpenCategoryRequest {
	s.UserId = &v
	return s
}

func (s *AddOpenCategoryRequest) SetUserName(v string) *AddOpenCategoryRequest {
	s.UserName = &v
	return s
}

type AddOpenCategoryResponseBody struct {
	// 返回结果
	Result *AddOpenCategoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddOpenCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOpenCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddOpenCategoryResponseBody) SetResult(v *AddOpenCategoryResponseBodyResult) *AddOpenCategoryResponseBody {
	s.Result = v
	return s
}

func (s *AddOpenCategoryResponseBody) SetSuccess(v bool) *AddOpenCategoryResponseBody {
	s.Success = &v
	return s
}

type AddOpenCategoryResponseBodyResult struct {
	// 添加成类目ID
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 失败时的错误消息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddOpenCategoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddOpenCategoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddOpenCategoryResponseBodyResult) SetId(v int64) *AddOpenCategoryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddOpenCategoryResponseBodyResult) SetMessage(v string) *AddOpenCategoryResponseBodyResult {
	s.Message = &v
	return s
}

func (s *AddOpenCategoryResponseBodyResult) SetSuccess(v bool) *AddOpenCategoryResponseBodyResult {
	s.Success = &v
	return s
}

type AddOpenCategoryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddOpenCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddOpenCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOpenCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddOpenCategoryResponse) SetHeaders(v map[string]*string) *AddOpenCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddOpenCategoryResponse) SetBody(v *AddOpenCategoryResponseBody) *AddOpenCategoryResponse {
	s.Body = v
	return s
}

type AddOpenKnowledgeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddOpenKnowledgeHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddOpenKnowledgeHeaders) GoString() string {
	return s.String()
}

func (s *AddOpenKnowledgeHeaders) SetCommonHeaders(v map[string]*string) *AddOpenKnowledgeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOpenKnowledgeHeaders) SetXAcsDingtalkAccessToken(v string) *AddOpenKnowledgeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddOpenKnowledgeRequest struct {
	// 附件列表
	Attachments []*AddOpenKnowledgeRequestAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// 知识点所属类目ID
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// 知识点正文
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 生效结束时间(默认2100-01-01 23:59:59)
	EffectTimeend *string `json:"effectTimeend,omitempty" xml:"effectTimeend,omitempty"`
	// 生效开始时间(默认1980-01-01 00:00:00)
	EffectTimestart *string `json:"effectTimestart,omitempty" xml:"effectTimestart,omitempty"`
	// 扩展问法(多个英文逗号隔开)
	ExtTitle *string `json:"extTitle,omitempty" xml:"extTitle,omitempty"`
	// 关键词(多个逗号隔开)
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 所属知识库唯一标识id
	LibraryId *int64 `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 知识点来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 标签(多个可逗号隔开)
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 知识点标准问
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 知识点类型()
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 用户/员工ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户昵称或姓名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s AddOpenKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOpenKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *AddOpenKnowledgeRequest) SetAttachments(v []*AddOpenKnowledgeRequestAttachments) *AddOpenKnowledgeRequest {
	s.Attachments = v
	return s
}

func (s *AddOpenKnowledgeRequest) SetCategoryId(v int64) *AddOpenKnowledgeRequest {
	s.CategoryId = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetContent(v string) *AddOpenKnowledgeRequest {
	s.Content = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetEffectTimeend(v string) *AddOpenKnowledgeRequest {
	s.EffectTimeend = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetEffectTimestart(v string) *AddOpenKnowledgeRequest {
	s.EffectTimestart = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetExtTitle(v string) *AddOpenKnowledgeRequest {
	s.ExtTitle = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetKeyword(v string) *AddOpenKnowledgeRequest {
	s.Keyword = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetLibraryId(v int64) *AddOpenKnowledgeRequest {
	s.LibraryId = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetOpenTeamId(v string) *AddOpenKnowledgeRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetSource(v string) *AddOpenKnowledgeRequest {
	s.Source = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetTags(v string) *AddOpenKnowledgeRequest {
	s.Tags = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetTitle(v string) *AddOpenKnowledgeRequest {
	s.Title = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetType(v string) *AddOpenKnowledgeRequest {
	s.Type = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetUserId(v string) *AddOpenKnowledgeRequest {
	s.UserId = &v
	return s
}

func (s *AddOpenKnowledgeRequest) SetUserName(v string) *AddOpenKnowledgeRequest {
	s.UserName = &v
	return s
}

type AddOpenKnowledgeRequestAttachments struct {
	// 媒体类型(扩展名大写)
	MimeType *string `json:"mimeType,omitempty" xml:"mimeType,omitempty"`
	// 附件URL
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 附件大小
	Size *float64 `json:"size,omitempty" xml:"size,omitempty"`
	// 扩展名
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
	// 附件名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s AddOpenKnowledgeRequestAttachments) String() string {
	return tea.Prettify(s)
}

func (s AddOpenKnowledgeRequestAttachments) GoString() string {
	return s.String()
}

func (s *AddOpenKnowledgeRequestAttachments) SetMimeType(v string) *AddOpenKnowledgeRequestAttachments {
	s.MimeType = &v
	return s
}

func (s *AddOpenKnowledgeRequestAttachments) SetPath(v string) *AddOpenKnowledgeRequestAttachments {
	s.Path = &v
	return s
}

func (s *AddOpenKnowledgeRequestAttachments) SetSize(v float64) *AddOpenKnowledgeRequestAttachments {
	s.Size = &v
	return s
}

func (s *AddOpenKnowledgeRequestAttachments) SetSuffix(v string) *AddOpenKnowledgeRequestAttachments {
	s.Suffix = &v
	return s
}

func (s *AddOpenKnowledgeRequestAttachments) SetTitle(v string) *AddOpenKnowledgeRequestAttachments {
	s.Title = &v
	return s
}

type AddOpenKnowledgeResponseBody struct {
	// 返回结果
	Result *AddOpenKnowledgeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddOpenKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOpenKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *AddOpenKnowledgeResponseBody) SetResult(v *AddOpenKnowledgeResponseBodyResult) *AddOpenKnowledgeResponseBody {
	s.Result = v
	return s
}

func (s *AddOpenKnowledgeResponseBody) SetSuccess(v bool) *AddOpenKnowledgeResponseBody {
	s.Success = &v
	return s
}

type AddOpenKnowledgeResponseBodyResult struct {
	// 知识点ID
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 失败错误消息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作是否成功标识
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddOpenKnowledgeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddOpenKnowledgeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddOpenKnowledgeResponseBodyResult) SetId(v int64) *AddOpenKnowledgeResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddOpenKnowledgeResponseBodyResult) SetMessage(v string) *AddOpenKnowledgeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *AddOpenKnowledgeResponseBodyResult) SetSuccess(v bool) *AddOpenKnowledgeResponseBodyResult {
	s.Success = &v
	return s
}

type AddOpenKnowledgeResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddOpenKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddOpenKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOpenKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *AddOpenKnowledgeResponse) SetHeaders(v map[string]*string) *AddOpenKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *AddOpenKnowledgeResponse) SetBody(v *AddOpenKnowledgeResponseBody) *AddOpenKnowledgeResponse {
	s.Body = v
	return s
}

type AddOpenLibraryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddOpenLibraryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddOpenLibraryHeaders) GoString() string {
	return s.String()
}

func (s *AddOpenLibraryHeaders) SetCommonHeaders(v map[string]*string) *AddOpenLibraryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOpenLibraryHeaders) SetXAcsDingtalkAccessToken(v string) *AddOpenLibraryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddOpenLibraryRequest struct {
	// 知识库描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 知识库来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 知识库名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 知识库类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 用户/员工ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户昵称或姓名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s AddOpenLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOpenLibraryRequest) GoString() string {
	return s.String()
}

func (s *AddOpenLibraryRequest) SetDescription(v string) *AddOpenLibraryRequest {
	s.Description = &v
	return s
}

func (s *AddOpenLibraryRequest) SetOpenTeamId(v string) *AddOpenLibraryRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AddOpenLibraryRequest) SetSource(v string) *AddOpenLibraryRequest {
	s.Source = &v
	return s
}

func (s *AddOpenLibraryRequest) SetTitle(v string) *AddOpenLibraryRequest {
	s.Title = &v
	return s
}

func (s *AddOpenLibraryRequest) SetType(v string) *AddOpenLibraryRequest {
	s.Type = &v
	return s
}

func (s *AddOpenLibraryRequest) SetUserId(v string) *AddOpenLibraryRequest {
	s.UserId = &v
	return s
}

func (s *AddOpenLibraryRequest) SetUserName(v string) *AddOpenLibraryRequest {
	s.UserName = &v
	return s
}

type AddOpenLibraryResponseBody struct {
	// 返回结果
	Result *AddOpenLibraryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddOpenLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOpenLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *AddOpenLibraryResponseBody) SetResult(v *AddOpenLibraryResponseBodyResult) *AddOpenLibraryResponseBody {
	s.Result = v
	return s
}

func (s *AddOpenLibraryResponseBody) SetSuccess(v bool) *AddOpenLibraryResponseBody {
	s.Success = &v
	return s
}

type AddOpenLibraryResponseBodyResult struct {
	// 知识库ID
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 失败时错误消息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 添加/修改知识库是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddOpenLibraryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddOpenLibraryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddOpenLibraryResponseBodyResult) SetId(v int64) *AddOpenLibraryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddOpenLibraryResponseBodyResult) SetMessage(v string) *AddOpenLibraryResponseBodyResult {
	s.Message = &v
	return s
}

func (s *AddOpenLibraryResponseBodyResult) SetSuccess(v bool) *AddOpenLibraryResponseBodyResult {
	s.Success = &v
	return s
}

type AddOpenLibraryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddOpenLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddOpenLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOpenLibraryResponse) GoString() string {
	return s.String()
}

func (s *AddOpenLibraryResponse) SetHeaders(v map[string]*string) *AddOpenLibraryResponse {
	s.Headers = v
	return s
}

func (s *AddOpenLibraryResponse) SetBody(v *AddOpenLibraryResponseBody) *AddOpenLibraryResponse {
	s.Body = v
	return s
}

type AddTicketMemoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddTicketMemoHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddTicketMemoHeaders) GoString() string {
	return s.String()
}

func (s *AddTicketMemoHeaders) SetCommonHeaders(v map[string]*string) *AddTicketMemoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddTicketMemoHeaders) SetXAcsDingtalkAccessToken(v string) *AddTicketMemoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddTicketMemoRequest struct {
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 当前工单处理人
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// 备注
	TicketMemo *AddTicketMemoRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s AddTicketMemoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTicketMemoRequest) GoString() string {
	return s.String()
}

func (s *AddTicketMemoRequest) SetOpenTeamId(v string) *AddTicketMemoRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AddTicketMemoRequest) SetOpenTicketId(v string) *AddTicketMemoRequest {
	s.OpenTicketId = &v
	return s
}

func (s *AddTicketMemoRequest) SetProcessorUnionId(v string) *AddTicketMemoRequest {
	s.ProcessorUnionId = &v
	return s
}

func (s *AddTicketMemoRequest) SetTicketMemo(v *AddTicketMemoRequestTicketMemo) *AddTicketMemoRequest {
	s.TicketMemo = v
	return s
}

type AddTicketMemoRequestTicketMemo struct {
	// 备注相关的附件
	Attachments []*AddTicketMemoRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// 文字备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s AddTicketMemoRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s AddTicketMemoRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *AddTicketMemoRequestTicketMemo) SetAttachments(v []*AddTicketMemoRequestTicketMemoAttachments) *AddTicketMemoRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *AddTicketMemoRequestTicketMemo) SetMemo(v string) *AddTicketMemoRequestTicketMemo {
	s.Memo = &v
	return s
}

type AddTicketMemoRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s AddTicketMemoRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s AddTicketMemoRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *AddTicketMemoRequestTicketMemoAttachments) SetFileName(v string) *AddTicketMemoRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *AddTicketMemoRequestTicketMemoAttachments) SetKey(v string) *AddTicketMemoRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type AddTicketMemoResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddTicketMemoResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTicketMemoResponse) GoString() string {
	return s.String()
}

func (s *AddTicketMemoResponse) SetHeaders(v map[string]*string) *AddTicketMemoResponse {
	s.Headers = v
	return s
}

type AssignTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AssignTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketHeaders) GoString() string {
	return s.String()
}

func (s *AssignTicketHeaders) SetCommonHeaders(v map[string]*string) *AssignTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AssignTicketHeaders) SetXAcsDingtalkAccessToken(v string) *AssignTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AssignTicketRequest struct {
	Notify *AssignTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 操作人unionId（管理员）
	OperatorUnionId   *string   `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	ProcessorUnionIds []*string `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	// 备注
	TicketMemo *AssignTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s AssignTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketRequest) GoString() string {
	return s.String()
}

func (s *AssignTicketRequest) SetNotify(v *AssignTicketRequestNotify) *AssignTicketRequest {
	s.Notify = v
	return s
}

func (s *AssignTicketRequest) SetOpenTeamId(v string) *AssignTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AssignTicketRequest) SetOpenTicketId(v string) *AssignTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *AssignTicketRequest) SetOperatorUnionId(v string) *AssignTicketRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *AssignTicketRequest) SetProcessorUnionIds(v []*string) *AssignTicketRequest {
	s.ProcessorUnionIds = v
	return s
}

func (s *AssignTicketRequest) SetTicketMemo(v *AssignTicketRequestTicketMemo) *AssignTicketRequest {
	s.TicketMemo = v
	return s
}

type AssignTicketRequestNotify struct {
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember       *bool     `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
	WorkNoticeReceiverUnionIds []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
}

func (s AssignTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *AssignTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *AssignTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *AssignTicketRequestNotify) SetNoticeAllGroupMember(v bool) *AssignTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *AssignTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *AssignTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

type AssignTicketRequestTicketMemo struct {
	// 备注相关的附件
	Attachments []*AssignTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// 备注文字
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s AssignTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *AssignTicketRequestTicketMemo) SetAttachments(v []*AssignTicketRequestTicketMemoAttachments) *AssignTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *AssignTicketRequestTicketMemo) SetMemo(v string) *AssignTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

type AssignTicketRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s AssignTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *AssignTicketRequestTicketMemoAttachments) SetFileName(v string) *AssignTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *AssignTicketRequestTicketMemoAttachments) SetKey(v string) *AssignTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type AssignTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AssignTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketResponse) GoString() string {
	return s.String()
}

func (s *AssignTicketResponse) SetHeaders(v map[string]*string) *AssignTicketResponse {
	s.Headers = v
	return s
}

type BatchGetGroupSetConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchGetGroupSetConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchGetGroupSetConfigHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetGroupSetConfigHeaders) SetCommonHeaders(v map[string]*string) *BatchGetGroupSetConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetGroupSetConfigHeaders) SetXAcsDingtalkAccessToken(v string) *BatchGetGroupSetConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchGetGroupSetConfigRequest struct {
	// 配置项key列表
	ConfigKeys []*string `json:"configKeys,omitempty" xml:"configKeys,omitempty" type:"Repeated"`
	// 开放群组id
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 开放团队id
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s BatchGetGroupSetConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetGroupSetConfigRequest) GoString() string {
	return s.String()
}

func (s *BatchGetGroupSetConfigRequest) SetConfigKeys(v []*string) *BatchGetGroupSetConfigRequest {
	s.ConfigKeys = v
	return s
}

func (s *BatchGetGroupSetConfigRequest) SetOpenGroupSetId(v string) *BatchGetGroupSetConfigRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *BatchGetGroupSetConfigRequest) SetOpenTeamId(v string) *BatchGetGroupSetConfigRequest {
	s.OpenTeamId = &v
	return s
}

type BatchGetGroupSetConfigResponseBody struct {
	// 群粗配置列表
	GroupSetConfigs []*BatchGetGroupSetConfigResponseBodyGroupSetConfigs `json:"groupSetConfigs,omitempty" xml:"groupSetConfigs,omitempty" type:"Repeated"`
}

func (s BatchGetGroupSetConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetGroupSetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetGroupSetConfigResponseBody) SetGroupSetConfigs(v []*BatchGetGroupSetConfigResponseBodyGroupSetConfigs) *BatchGetGroupSetConfigResponseBody {
	s.GroupSetConfigs = v
	return s
}

type BatchGetGroupSetConfigResponseBodyGroupSetConfigs struct {
	// 配置项key
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// 配置项值
	ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
}

func (s BatchGetGroupSetConfigResponseBodyGroupSetConfigs) String() string {
	return tea.Prettify(s)
}

func (s BatchGetGroupSetConfigResponseBodyGroupSetConfigs) GoString() string {
	return s.String()
}

func (s *BatchGetGroupSetConfigResponseBodyGroupSetConfigs) SetConfigKey(v string) *BatchGetGroupSetConfigResponseBodyGroupSetConfigs {
	s.ConfigKey = &v
	return s
}

func (s *BatchGetGroupSetConfigResponseBodyGroupSetConfigs) SetConfigValue(v string) *BatchGetGroupSetConfigResponseBodyGroupSetConfigs {
	s.ConfigValue = &v
	return s
}

type BatchGetGroupSetConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchGetGroupSetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetGroupSetConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetGroupSetConfigResponse) GoString() string {
	return s.String()
}

func (s *BatchGetGroupSetConfigResponse) SetHeaders(v map[string]*string) *BatchGetGroupSetConfigResponse {
	s.Headers = v
	return s
}

func (s *BatchGetGroupSetConfigResponse) SetBody(v *BatchGetGroupSetConfigResponseBody) *BatchGetGroupSetConfigResponse {
	s.Body = v
	return s
}

type BatchQuerySendMessageTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQuerySendMessageTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySendMessageTaskHeaders) GoString() string {
	return s.String()
}

func (s *BatchQuerySendMessageTaskHeaders) SetCommonHeaders(v map[string]*string) *BatchQuerySendMessageTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQuerySendMessageTaskHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQuerySendMessageTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQuerySendMessageTaskRequest struct {
	// 是否获取群发任务已读数量，默认false
	GetReadCount *bool `json:"getReadCount,omitempty" xml:"getReadCount,omitempty"`
	// 任务查询结束时间
	GmtCreateEnd *string `json:"gmtCreateEnd,omitempty" xml:"gmtCreateEnd,omitempty"`
	// 任务查询开始时间
	GmtCreateStart *string `json:"gmtCreateStart,omitempty" xml:"gmtCreateStart,omitempty"`
	// 每页条数
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 开放群组ID，在服务群-群组- ID信息中获取
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 任务名称
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s BatchQuerySendMessageTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySendMessageTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchQuerySendMessageTaskRequest) SetGetReadCount(v bool) *BatchQuerySendMessageTaskRequest {
	s.GetReadCount = &v
	return s
}

func (s *BatchQuerySendMessageTaskRequest) SetGmtCreateEnd(v string) *BatchQuerySendMessageTaskRequest {
	s.GmtCreateEnd = &v
	return s
}

func (s *BatchQuerySendMessageTaskRequest) SetGmtCreateStart(v string) *BatchQuerySendMessageTaskRequest {
	s.GmtCreateStart = &v
	return s
}

func (s *BatchQuerySendMessageTaskRequest) SetMaxResults(v int64) *BatchQuerySendMessageTaskRequest {
	s.MaxResults = &v
	return s
}

func (s *BatchQuerySendMessageTaskRequest) SetNextToken(v string) *BatchQuerySendMessageTaskRequest {
	s.NextToken = &v
	return s
}

func (s *BatchQuerySendMessageTaskRequest) SetOpenGroupSetId(v string) *BatchQuerySendMessageTaskRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *BatchQuerySendMessageTaskRequest) SetOpenTeamId(v string) *BatchQuerySendMessageTaskRequest {
	s.OpenTeamId = &v
	return s
}

func (s *BatchQuerySendMessageTaskRequest) SetTaskName(v string) *BatchQuerySendMessageTaskRequest {
	s.TaskName = &v
	return s
}

type BatchQuerySendMessageTaskResponseBody struct {
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Id of the request
	NextToken  *string                                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Records    []*BatchQuerySendMessageTaskResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	TotalCount *float32                                        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s BatchQuerySendMessageTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySendMessageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQuerySendMessageTaskResponseBody) SetMaxResults(v int64) *BatchQuerySendMessageTaskResponseBody {
	s.MaxResults = &v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBody) SetNextToken(v string) *BatchQuerySendMessageTaskResponseBody {
	s.NextToken = &v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBody) SetRecords(v []*BatchQuerySendMessageTaskResponseBodyRecords) *BatchQuerySendMessageTaskResponseBody {
	s.Records = v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBody) SetTotalCount(v float32) *BatchQuerySendMessageTaskResponseBody {
	s.TotalCount = &v
	return s
}

type BatchQuerySendMessageTaskResponseBodyRecords struct {
	CreateName        *string `json:"createName,omitempty" xml:"createName,omitempty"`
	CreateTimeStr     *string `json:"createTimeStr,omitempty" xml:"createTimeStr,omitempty"`
	CreateUnionId     *string `json:"createUnionId,omitempty" xml:"createUnionId,omitempty"`
	OpenBatchTaskId   *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
	ReadGroupInc      *int64  `json:"readGroupInc,omitempty" xml:"readGroupInc,omitempty"`
	SendGroupInc      *int64  `json:"sendGroupInc,omitempty" xml:"sendGroupInc,omitempty"`
	SendMessageStatus *string `json:"sendMessageStatus,omitempty" xml:"sendMessageStatus,omitempty"`
	SendTaskTimeStr   *string `json:"sendTaskTimeStr,omitempty" xml:"sendTaskTimeStr,omitempty"`
	TaskName          *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s BatchQuerySendMessageTaskResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySendMessageTaskResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *BatchQuerySendMessageTaskResponseBodyRecords) SetCreateName(v string) *BatchQuerySendMessageTaskResponseBodyRecords {
	s.CreateName = &v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBodyRecords) SetCreateTimeStr(v string) *BatchQuerySendMessageTaskResponseBodyRecords {
	s.CreateTimeStr = &v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBodyRecords) SetCreateUnionId(v string) *BatchQuerySendMessageTaskResponseBodyRecords {
	s.CreateUnionId = &v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBodyRecords) SetOpenBatchTaskId(v string) *BatchQuerySendMessageTaskResponseBodyRecords {
	s.OpenBatchTaskId = &v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBodyRecords) SetReadGroupInc(v int64) *BatchQuerySendMessageTaskResponseBodyRecords {
	s.ReadGroupInc = &v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBodyRecords) SetSendGroupInc(v int64) *BatchQuerySendMessageTaskResponseBodyRecords {
	s.SendGroupInc = &v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBodyRecords) SetSendMessageStatus(v string) *BatchQuerySendMessageTaskResponseBodyRecords {
	s.SendMessageStatus = &v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBodyRecords) SetSendTaskTimeStr(v string) *BatchQuerySendMessageTaskResponseBodyRecords {
	s.SendTaskTimeStr = &v
	return s
}

func (s *BatchQuerySendMessageTaskResponseBodyRecords) SetTaskName(v string) *BatchQuerySendMessageTaskResponseBodyRecords {
	s.TaskName = &v
	return s
}

type BatchQuerySendMessageTaskResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchQuerySendMessageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchQuerySendMessageTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySendMessageTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchQuerySendMessageTaskResponse) SetHeaders(v map[string]*string) *BatchQuerySendMessageTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchQuerySendMessageTaskResponse) SetBody(v *BatchQuerySendMessageTaskResponseBody) *BatchQuerySendMessageTaskResponse {
	s.Body = v
	return s
}

type BoundTemplateToTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BoundTemplateToTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s BoundTemplateToTeamHeaders) GoString() string {
	return s.String()
}

func (s *BoundTemplateToTeamHeaders) SetCommonHeaders(v map[string]*string) *BoundTemplateToTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BoundTemplateToTeamHeaders) SetXAcsDingtalkAccessToken(v string) *BoundTemplateToTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BoundTemplateToTeamRequest struct {
	// 目标团队id
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 模板中的机器人配置信息
	RobotConfig *string `json:"robotConfig,omitempty" xml:"robotConfig,omitempty"`
	// 模板描述信息
	TemplateDesc *string `json:"templateDesc,omitempty" xml:"templateDesc,omitempty"`
	// 模板id
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 模板名字
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// 模板类型
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s BoundTemplateToTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s BoundTemplateToTeamRequest) GoString() string {
	return s.String()
}

func (s *BoundTemplateToTeamRequest) SetOpenTeamId(v string) *BoundTemplateToTeamRequest {
	s.OpenTeamId = &v
	return s
}

func (s *BoundTemplateToTeamRequest) SetRobotConfig(v string) *BoundTemplateToTeamRequest {
	s.RobotConfig = &v
	return s
}

func (s *BoundTemplateToTeamRequest) SetTemplateDesc(v string) *BoundTemplateToTeamRequest {
	s.TemplateDesc = &v
	return s
}

func (s *BoundTemplateToTeamRequest) SetTemplateId(v string) *BoundTemplateToTeamRequest {
	s.TemplateId = &v
	return s
}

func (s *BoundTemplateToTeamRequest) SetTemplateName(v string) *BoundTemplateToTeamRequest {
	s.TemplateName = &v
	return s
}

func (s *BoundTemplateToTeamRequest) SetTemplateType(v string) *BoundTemplateToTeamRequest {
	s.TemplateType = &v
	return s
}

type BoundTemplateToTeamResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BoundTemplateToTeamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BoundTemplateToTeamResponseBody) GoString() string {
	return s.String()
}

func (s *BoundTemplateToTeamResponseBody) SetResult(v bool) *BoundTemplateToTeamResponseBody {
	s.Result = &v
	return s
}

type BoundTemplateToTeamResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BoundTemplateToTeamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BoundTemplateToTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s BoundTemplateToTeamResponse) GoString() string {
	return s.String()
}

func (s *BoundTemplateToTeamResponse) SetHeaders(v map[string]*string) *BoundTemplateToTeamResponse {
	s.Headers = v
	return s
}

func (s *BoundTemplateToTeamResponse) SetBody(v *BoundTemplateToTeamResponseBody) *BoundTemplateToTeamResponse {
	s.Body = v
	return s
}

type CancelTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CancelTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelTicketHeaders) GoString() string {
	return s.String()
}

func (s *CancelTicketHeaders) SetCommonHeaders(v map[string]*string) *CancelTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelTicketHeaders) SetXAcsDingtalkAccessToken(v string) *CancelTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelTicketRequest struct {
	Notify *CancelTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 操作人unionId
	OperatorUnionId *string                        `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	TicketMemo      *CancelTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s CancelTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelTicketRequest) GoString() string {
	return s.String()
}

func (s *CancelTicketRequest) SetNotify(v *CancelTicketRequestNotify) *CancelTicketRequest {
	s.Notify = v
	return s
}

func (s *CancelTicketRequest) SetOpenTeamId(v string) *CancelTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CancelTicketRequest) SetOpenTicketId(v string) *CancelTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *CancelTicketRequest) SetOperatorUnionId(v string) *CancelTicketRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *CancelTicketRequest) SetTicketMemo(v *CancelTicketRequestTicketMemo) *CancelTicketRequest {
	s.TicketMemo = v
	return s
}

type CancelTicketRequestNotify struct {
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember       *bool     `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
	WorkNoticeReceiverUnionIds []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
}

func (s CancelTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s CancelTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *CancelTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *CancelTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *CancelTicketRequestNotify) SetNoticeAllGroupMember(v bool) *CancelTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *CancelTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *CancelTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

type CancelTicketRequestTicketMemo struct {
	Attachments []*CancelTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s CancelTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s CancelTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *CancelTicketRequestTicketMemo) SetAttachments(v []*CancelTicketRequestTicketMemoAttachments) *CancelTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *CancelTicketRequestTicketMemo) SetMemo(v string) *CancelTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

type CancelTicketRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s CancelTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s CancelTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *CancelTicketRequestTicketMemoAttachments) SetFileName(v string) *CancelTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *CancelTicketRequestTicketMemoAttachments) SetKey(v string) *CancelTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type CancelTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CancelTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelTicketResponse) GoString() string {
	return s.String()
}

func (s *CancelTicketResponse) SetHeaders(v map[string]*string) *CancelTicketResponse {
	s.Headers = v
	return s
}

type CloseHumanSessionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CloseHumanSessionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CloseHumanSessionHeaders) GoString() string {
	return s.String()
}

func (s *CloseHumanSessionHeaders) SetCommonHeaders(v map[string]*string) *CloseHumanSessionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseHumanSessionHeaders) SetXAcsDingtalkAccessToken(v string) *CloseHumanSessionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CloseHumanSessionRequest struct {
	// 开放会话id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 开放团队id
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s CloseHumanSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseHumanSessionRequest) GoString() string {
	return s.String()
}

func (s *CloseHumanSessionRequest) SetOpenConversationId(v string) *CloseHumanSessionRequest {
	s.OpenConversationId = &v
	return s
}

func (s *CloseHumanSessionRequest) SetOpenTeamId(v string) *CloseHumanSessionRequest {
	s.OpenTeamId = &v
	return s
}

type CloseHumanSessionResponseBody struct {
	SessionId *int64 `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CloseHumanSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseHumanSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseHumanSessionResponseBody) SetSessionId(v int64) *CloseHumanSessionResponseBody {
	s.SessionId = &v
	return s
}

type CloseHumanSessionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseHumanSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseHumanSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseHumanSessionResponse) GoString() string {
	return s.String()
}

func (s *CloseHumanSessionResponse) SetHeaders(v map[string]*string) *CloseHumanSessionResponse {
	s.Headers = v
	return s
}

func (s *CloseHumanSessionResponse) SetBody(v *CloseHumanSessionResponseBody) *CloseHumanSessionResponse {
	s.Body = v
	return s
}

type CreateGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupRequest struct {
	// 业务关联id
	GroupBizId *string `json:"groupBizId,omitempty" xml:"groupBizId,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 群标签
	GroupTagNames []*string `json:"groupTagNames,omitempty" xml:"groupTagNames,omitempty" type:"Repeated"`
	// 群成员员工ID列表
	MemberStaffIds []*string `json:"memberStaffIds,omitempty" xml:"memberStaffIds,omitempty" type:"Repeated"`
	// 开放群组ID
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 群主员工ID
	OwnerStaffId *string `json:"ownerStaffId,omitempty" xml:"ownerStaffId,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetGroupBizId(v string) *CreateGroupRequest {
	s.GroupBizId = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) SetGroupTagNames(v []*string) *CreateGroupRequest {
	s.GroupTagNames = v
	return s
}

func (s *CreateGroupRequest) SetMemberStaffIds(v []*string) *CreateGroupRequest {
	s.MemberStaffIds = v
	return s
}

func (s *CreateGroupRequest) SetOpenGroupSetId(v string) *CreateGroupRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *CreateGroupRequest) SetOpenTeamId(v string) *CreateGroupRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CreateGroupRequest) SetOwnerStaffId(v string) *CreateGroupRequest {
	s.OwnerStaffId = &v
	return s
}

type CreateGroupResponseBody struct {
	// 入群url
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetGroupUrl(v string) *CreateGroupResponseBody {
	s.GroupUrl = &v
	return s
}

func (s *CreateGroupResponseBody) SetOpenConversationId(v string) *CreateGroupResponseBody {
	s.OpenConversationId = &v
	return s
}

type CreateGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateGroupSetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupSetHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupSetHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupSetHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupSetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupSetHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupSetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupSetRequest struct {
	// groupSetName
	GroupSetName    *string `json:"groupSetName,omitempty" xml:"groupSetName,omitempty"`
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	// openTeamId
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s CreateGroupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupSetRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupSetRequest) SetGroupSetName(v string) *CreateGroupSetRequest {
	s.GroupSetName = &v
	return s
}

func (s *CreateGroupSetRequest) SetGroupTemplateId(v string) *CreateGroupSetRequest {
	s.GroupTemplateId = &v
	return s
}

func (s *CreateGroupSetRequest) SetOpenTeamId(v string) *CreateGroupSetRequest {
	s.OpenTeamId = &v
	return s
}

type CreateGroupSetResponseBody struct {
	// 群分组id
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateGroupSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupSetResponseBody) SetResult(v string) *CreateGroupSetResponseBody {
	s.Result = &v
	return s
}

type CreateGroupSetResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupSetResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupSetResponse) SetHeaders(v map[string]*string) *CreateGroupSetResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupSetResponse) SetBody(v *CreateGroupSetResponseBody) *CreateGroupSetResponse {
	s.Body = v
	return s
}

type CreateTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTeamHeaders) GoString() string {
	return s.String()
}

func (s *CreateTeamHeaders) SetCommonHeaders(v map[string]*string) *CreateTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTeamHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTeamRequest struct {
	// 团队管理员钉钉unionId
	CreatorDingUnionId *string `json:"creatorDingUnionId,omitempty" xml:"creatorDingUnionId,omitempty"`
	// 团队名字
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
}

func (s CreateTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTeamRequest) GoString() string {
	return s.String()
}

func (s *CreateTeamRequest) SetCreatorDingUnionId(v string) *CreateTeamRequest {
	s.CreatorDingUnionId = &v
	return s
}

func (s *CreateTeamRequest) SetTeamName(v string) *CreateTeamRequest {
	s.TeamName = &v
	return s
}

type CreateTeamResponseBody struct {
	// 团队id
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateTeamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTeamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTeamResponseBody) SetResult(v string) *CreateTeamResponseBody {
	s.Result = &v
	return s
}

type CreateTeamResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTeamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTeamResponse) GoString() string {
	return s.String()
}

func (s *CreateTeamResponse) SetHeaders(v map[string]*string) *CreateTeamResponse {
	s.Headers = v
	return s
}

func (s *CreateTeamResponse) SetBody(v *CreateTeamResponseBody) *CreateTeamResponse {
	s.Body = v
	return s
}

type CreateTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketHeaders) GoString() string {
	return s.String()
}

func (s *CreateTicketHeaders) SetCommonHeaders(v map[string]*string) *CreateTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTicketHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTicketRequest struct {
	// 工单创建人UnionId
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// 自定义组件字段值(JSON格式)
	CustomFields *string `json:"customFields,omitempty" xml:"customFields,omitempty"`
	// 通知接收人配置
	Notify *CreateTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单模板业务ID
	OpenTemplateBizId *string `json:"openTemplateBizId,omitempty" xml:"openTemplateBizId,omitempty"`
	// 工单处理人UnionId列表
	ProcessorUnionIds []*string `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	// 工单场景 SG 或 VOC
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 工单场景信息
	SceneContext *CreateTicketRequestSceneContext `json:"sceneContext,omitempty" xml:"sceneContext,omitempty" type:"Struct"`
	// 工单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetCreatorUnionId(v string) *CreateTicketRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *CreateTicketRequest) SetCustomFields(v string) *CreateTicketRequest {
	s.CustomFields = &v
	return s
}

func (s *CreateTicketRequest) SetNotify(v *CreateTicketRequestNotify) *CreateTicketRequest {
	s.Notify = v
	return s
}

func (s *CreateTicketRequest) SetOpenTeamId(v string) *CreateTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CreateTicketRequest) SetOpenTemplateBizId(v string) *CreateTicketRequest {
	s.OpenTemplateBizId = &v
	return s
}

func (s *CreateTicketRequest) SetProcessorUnionIds(v []*string) *CreateTicketRequest {
	s.ProcessorUnionIds = v
	return s
}

func (s *CreateTicketRequest) SetScene(v string) *CreateTicketRequest {
	s.Scene = &v
	return s
}

func (s *CreateTicketRequest) SetSceneContext(v *CreateTicketRequestSceneContext) *CreateTicketRequest {
	s.SceneContext = v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

type CreateTicketRequestNotify struct {
	// 服务群通知接收人（钉钉UnionId）
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember *bool `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
	// 企业工作通知接收人（钉钉UnionId）
	WorkNoticeReceiverUnionIds []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
}

func (s CreateTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *CreateTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *CreateTicketRequestNotify) SetNoticeAllGroupMember(v bool) *CreateTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *CreateTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *CreateTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

type CreateTicketRequestSceneContext struct {
	// 工单相关的群消息列表
	GroupMsgs []*CreateTicketRequestSceneContextGroupMsgs `json:"groupMsgs,omitempty" xml:"groupMsgs,omitempty" type:"Repeated"`
	// 服务群openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 工单相关人UnionId列表
	RelevantorUnionIds []*string `json:"relevantorUnionIds,omitempty" xml:"relevantorUnionIds,omitempty" type:"Repeated"`
	// VOC类型工单，对应话题ID
	TopicId *string `json:"topicId,omitempty" xml:"topicId,omitempty"`
}

func (s CreateTicketRequestSceneContext) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestSceneContext) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestSceneContext) SetGroupMsgs(v []*CreateTicketRequestSceneContextGroupMsgs) *CreateTicketRequestSceneContext {
	s.GroupMsgs = v
	return s
}

func (s *CreateTicketRequestSceneContext) SetOpenConversationId(v string) *CreateTicketRequestSceneContext {
	s.OpenConversationId = &v
	return s
}

func (s *CreateTicketRequestSceneContext) SetRelevantorUnionIds(v []*string) *CreateTicketRequestSceneContext {
	s.RelevantorUnionIds = v
	return s
}

func (s *CreateTicketRequestSceneContext) SetTopicId(v string) *CreateTicketRequestSceneContext {
	s.TopicId = &v
	return s
}

type CreateTicketRequestSceneContextGroupMsgs struct {
	// 是否为锚点消息
	Anchor *bool `json:"anchor,omitempty" xml:"anchor,omitempty"`
	// 勾选消息openMsgId
	OpenMsgId *string `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
}

func (s CreateTicketRequestSceneContextGroupMsgs) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestSceneContextGroupMsgs) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestSceneContextGroupMsgs) SetAnchor(v bool) *CreateTicketRequestSceneContextGroupMsgs {
	s.Anchor = &v
	return s
}

func (s *CreateTicketRequestSceneContextGroupMsgs) SetOpenMsgId(v string) *CreateTicketRequestSceneContextGroupMsgs {
	s.OpenMsgId = &v
	return s
}

type CreateTicketResponseBody struct {
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetOpenTicketId(v string) *CreateTicketResponseBody {
	s.OpenTicketId = &v
	return s
}

type CreateTicketResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketResponse) SetHeaders(v map[string]*string) *CreateTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
	s.Body = v
	return s
}

type DeleteKnowledgeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteKnowledgeHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeHeaders) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeHeaders) SetCommonHeaders(v map[string]*string) *DeleteKnowledgeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteKnowledgeHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteKnowledgeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteKnowledgeRequest struct {
	// 知识库的唯一标识 比如:天工知识库ID
	LibraryKey *string `json:"libraryKey,omitempty" xml:"libraryKey,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 知识点来源 CCM:天工知识库
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 知识点唯一标识
	SourcePrimaryKey *string `json:"sourcePrimaryKey,omitempty" xml:"sourcePrimaryKey,omitempty"`
}

func (s DeleteKnowledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeRequest) SetLibraryKey(v string) *DeleteKnowledgeRequest {
	s.LibraryKey = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetOpenTeamId(v string) *DeleteKnowledgeRequest {
	s.OpenTeamId = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetSource(v string) *DeleteKnowledgeRequest {
	s.Source = &v
	return s
}

func (s *DeleteKnowledgeRequest) SetSourcePrimaryKey(v string) *DeleteKnowledgeRequest {
	s.SourcePrimaryKey = &v
	return s
}

type DeleteKnowledgeResponseBody struct {
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteKnowledgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeResponseBody) SetSuccess(v bool) *DeleteKnowledgeResponseBody {
	s.Success = &v
	return s
}

type DeleteKnowledgeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteKnowledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeResponse) SetBody(v *DeleteKnowledgeResponseBody) *DeleteKnowledgeResponse {
	s.Body = v
	return s
}

type FinishTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FinishTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketHeaders) GoString() string {
	return s.String()
}

func (s *FinishTicketHeaders) SetCommonHeaders(v map[string]*string) *FinishTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FinishTicketHeaders) SetXAcsDingtalkAccessToken(v string) *FinishTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FinishTicketRequest struct {
	// 工单通知
	Notify     *FinishTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	OpenTeamId *string                    `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单开放id
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 当前工单处理人
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// 备注
	TicketMemo *FinishTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s FinishTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketRequest) GoString() string {
	return s.String()
}

func (s *FinishTicketRequest) SetNotify(v *FinishTicketRequestNotify) *FinishTicketRequest {
	s.Notify = v
	return s
}

func (s *FinishTicketRequest) SetOpenTeamId(v string) *FinishTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *FinishTicketRequest) SetOpenTicketId(v string) *FinishTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *FinishTicketRequest) SetProcessorUnionId(v string) *FinishTicketRequest {
	s.ProcessorUnionId = &v
	return s
}

func (s *FinishTicketRequest) SetTicketMemo(v *FinishTicketRequestTicketMemo) *FinishTicketRequest {
	s.TicketMemo = v
	return s
}

type FinishTicketRequestNotify struct {
	// 群中通知接收人（钉钉UnionId）
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember *bool `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
	// 企业工作通知接收人（钉钉UnionId）
	WorkNoticeReceiverUnionIds []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
}

func (s FinishTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *FinishTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *FinishTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *FinishTicketRequestNotify) SetNoticeAllGroupMember(v bool) *FinishTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *FinishTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *FinishTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

type FinishTicketRequestTicketMemo struct {
	// 备注相关的附件
	Attachments []*FinishTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// 备注文字
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s FinishTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *FinishTicketRequestTicketMemo) SetAttachments(v []*FinishTicketRequestTicketMemoAttachments) *FinishTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *FinishTicketRequestTicketMemo) SetMemo(v string) *FinishTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

type FinishTicketRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s FinishTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *FinishTicketRequestTicketMemoAttachments) SetFileName(v string) *FinishTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *FinishTicketRequestTicketMemoAttachments) SetKey(v string) *FinishTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type FinishTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s FinishTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishTicketResponse) GoString() string {
	return s.String()
}

func (s *FinishTicketResponse) SetHeaders(v map[string]*string) *FinishTicketResponse {
	s.Headers = v
	return s
}

type GetOssTempUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOssTempUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOssTempUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetOssTempUrlHeaders) SetCommonHeaders(v map[string]*string) *GetOssTempUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOssTempUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetOssTempUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOssTempUrlRequest struct {
	// 访问模式 AUTO(自动，例如在浏览器中如果是图片，PDF等可以在线直接查看，不能在线看时自动下载)、DOWNLOAD（直接下载）
	FetchMode *string `json:"fetchMode,omitempty" xml:"fetchMode,omitempty"`
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// oss文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 团队开放ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s GetOssTempUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssTempUrlRequest) GoString() string {
	return s.String()
}

func (s *GetOssTempUrlRequest) SetFetchMode(v string) *GetOssTempUrlRequest {
	s.FetchMode = &v
	return s
}

func (s *GetOssTempUrlRequest) SetFileName(v string) *GetOssTempUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetOssTempUrlRequest) SetKey(v string) *GetOssTempUrlRequest {
	s.Key = &v
	return s
}

func (s *GetOssTempUrlRequest) SetOpenTeamId(v string) *GetOssTempUrlRequest {
	s.OpenTeamId = &v
	return s
}

type GetOssTempUrlResponseBody struct {
	// Id of the request
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetOssTempUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssTempUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssTempUrlResponseBody) SetUrl(v string) *GetOssTempUrlResponseBody {
	s.Url = &v
	return s
}

type GetOssTempUrlResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOssTempUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOssTempUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssTempUrlResponse) GoString() string {
	return s.String()
}

func (s *GetOssTempUrlResponse) SetHeaders(v map[string]*string) *GetOssTempUrlResponse {
	s.Headers = v
	return s
}

func (s *GetOssTempUrlResponse) SetBody(v *GetOssTempUrlResponseBody) *GetOssTempUrlResponse {
	s.Body = v
	return s
}

type GetStoragePolicyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetStoragePolicyHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetStoragePolicyHeaders) GoString() string {
	return s.String()
}

func (s *GetStoragePolicyHeaders) SetCommonHeaders(v map[string]*string) *GetStoragePolicyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetStoragePolicyHeaders) SetXAcsDingtalkAccessToken(v string) *GetStoragePolicyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetStoragePolicyRequest struct {
	// 业务类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 文件名称
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件大小，单位字节
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s GetStoragePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStoragePolicyRequest) GoString() string {
	return s.String()
}

func (s *GetStoragePolicyRequest) SetBizType(v string) *GetStoragePolicyRequest {
	s.BizType = &v
	return s
}

func (s *GetStoragePolicyRequest) SetFileName(v string) *GetStoragePolicyRequest {
	s.FileName = &v
	return s
}

func (s *GetStoragePolicyRequest) SetFileSize(v int64) *GetStoragePolicyRequest {
	s.FileSize = &v
	return s
}

func (s *GetStoragePolicyRequest) SetOpenTeamId(v string) *GetStoragePolicyRequest {
	s.OpenTeamId = &v
	return s
}

type GetStoragePolicyResponseBody struct {
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	Endpoint    *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// Id of the request
	Key       *string `json:"key,omitempty" xml:"key,omitempty"`
	Policy    *string `json:"policy,omitempty" xml:"policy,omitempty"`
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s GetStoragePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStoragePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetStoragePolicyResponseBody) SetAccessKeyId(v string) *GetStoragePolicyResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetStoragePolicyResponseBody) SetEndpoint(v string) *GetStoragePolicyResponseBody {
	s.Endpoint = &v
	return s
}

func (s *GetStoragePolicyResponseBody) SetKey(v string) *GetStoragePolicyResponseBody {
	s.Key = &v
	return s
}

func (s *GetStoragePolicyResponseBody) SetPolicy(v string) *GetStoragePolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetStoragePolicyResponseBody) SetSignature(v string) *GetStoragePolicyResponseBody {
	s.Signature = &v
	return s
}

type GetStoragePolicyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStoragePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStoragePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStoragePolicyResponse) GoString() string {
	return s.String()
}

func (s *GetStoragePolicyResponse) SetHeaders(v map[string]*string) *GetStoragePolicyResponse {
	s.Headers = v
	return s
}

func (s *GetStoragePolicyResponse) SetBody(v *GetStoragePolicyResponseBody) *GetStoragePolicyResponse {
	s.Body = v
	return s
}

type GetTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTicketHeaders) GoString() string {
	return s.String()
}

func (s *GetTicketHeaders) SetCommonHeaders(v map[string]*string) *GetTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTicketHeaders) SetXAcsDingtalkAccessToken(v string) *GetTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTicketRequest struct {
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// hNiPO2OVktNMiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
}

func (s GetTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTicketRequest) GoString() string {
	return s.String()
}

func (s *GetTicketRequest) SetOpenTeamId(v string) *GetTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *GetTicketRequest) SetOpenTicketId(v string) *GetTicketRequest {
	s.OpenTicketId = &v
	return s
}

type GetTicketResponseBody struct {
	CreateTime         *string                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator            *GetTicketResponseBodyCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	CustomFields       *string                       `json:"customFields,omitempty" xml:"customFields,omitempty"`
	OpenConversationId *string                       `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// Id of the request
	OpenTicketId *string                         `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	Processor    *GetTicketResponseBodyProcessor `json:"processor,omitempty" xml:"processor,omitempty" type:"Struct"`
	Scene        *string                         `json:"scene,omitempty" xml:"scene,omitempty"`
	SceneContext *string                         `json:"sceneContext,omitempty" xml:"sceneContext,omitempty"`
	Stage        *string                         `json:"stage,omitempty" xml:"stage,omitempty"`
	Takers       []*GetTicketResponseBodyTakers  `json:"takers,omitempty" xml:"takers,omitempty" type:"Repeated"`
	Template     *GetTicketResponseBodyTemplate  `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	Title        *string                         `json:"title,omitempty" xml:"title,omitempty"`
	UpdateTime   *string                         `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBody) SetCreateTime(v string) *GetTicketResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTicketResponseBody) SetCreator(v *GetTicketResponseBodyCreator) *GetTicketResponseBody {
	s.Creator = v
	return s
}

func (s *GetTicketResponseBody) SetCustomFields(v string) *GetTicketResponseBody {
	s.CustomFields = &v
	return s
}

func (s *GetTicketResponseBody) SetOpenConversationId(v string) *GetTicketResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GetTicketResponseBody) SetOpenTicketId(v string) *GetTicketResponseBody {
	s.OpenTicketId = &v
	return s
}

func (s *GetTicketResponseBody) SetProcessor(v *GetTicketResponseBodyProcessor) *GetTicketResponseBody {
	s.Processor = v
	return s
}

func (s *GetTicketResponseBody) SetScene(v string) *GetTicketResponseBody {
	s.Scene = &v
	return s
}

func (s *GetTicketResponseBody) SetSceneContext(v string) *GetTicketResponseBody {
	s.SceneContext = &v
	return s
}

func (s *GetTicketResponseBody) SetStage(v string) *GetTicketResponseBody {
	s.Stage = &v
	return s
}

func (s *GetTicketResponseBody) SetTakers(v []*GetTicketResponseBodyTakers) *GetTicketResponseBody {
	s.Takers = v
	return s
}

func (s *GetTicketResponseBody) SetTemplate(v *GetTicketResponseBodyTemplate) *GetTicketResponseBody {
	s.Template = v
	return s
}

func (s *GetTicketResponseBody) SetTitle(v string) *GetTicketResponseBody {
	s.Title = &v
	return s
}

func (s *GetTicketResponseBody) SetUpdateTime(v string) *GetTicketResponseBody {
	s.UpdateTime = &v
	return s
}

type GetTicketResponseBodyCreator struct {
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetTicketResponseBodyCreator) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyCreator) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyCreator) SetNickName(v string) *GetTicketResponseBodyCreator {
	s.NickName = &v
	return s
}

func (s *GetTicketResponseBodyCreator) SetUnionId(v string) *GetTicketResponseBodyCreator {
	s.UnionId = &v
	return s
}

type GetTicketResponseBodyProcessor struct {
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetTicketResponseBodyProcessor) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyProcessor) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyProcessor) SetNickName(v string) *GetTicketResponseBodyProcessor {
	s.NickName = &v
	return s
}

func (s *GetTicketResponseBodyProcessor) SetUnionId(v string) *GetTicketResponseBodyProcessor {
	s.UnionId = &v
	return s
}

type GetTicketResponseBodyTakers struct {
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetTicketResponseBodyTakers) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyTakers) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyTakers) SetNickName(v string) *GetTicketResponseBodyTakers {
	s.NickName = &v
	return s
}

func (s *GetTicketResponseBodyTakers) SetUnionId(v string) *GetTicketResponseBodyTakers {
	s.UnionId = &v
	return s
}

type GetTicketResponseBodyTemplate struct {
	// 工单模版业务ID
	OpenTemplateBizId *string `json:"openTemplateBizId,omitempty" xml:"openTemplateBizId,omitempty"`
	// 工单模版ID
	OpenTemplateId *string `json:"openTemplateId,omitempty" xml:"openTemplateId,omitempty"`
	// 工单模版名称
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s GetTicketResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyTemplate) SetOpenTemplateBizId(v string) *GetTicketResponseBodyTemplate {
	s.OpenTemplateBizId = &v
	return s
}

func (s *GetTicketResponseBodyTemplate) SetOpenTemplateId(v string) *GetTicketResponseBodyTemplate {
	s.OpenTemplateId = &v
	return s
}

func (s *GetTicketResponseBodyTemplate) SetTemplateName(v string) *GetTicketResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

type GetTicketResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponse) GoString() string {
	return s.String()
}

func (s *GetTicketResponse) SetHeaders(v map[string]*string) *GetTicketResponse {
	s.Headers = v
	return s
}

func (s *GetTicketResponse) SetBody(v *GetTicketResponseBody) *GetTicketResponse {
	s.Body = v
	return s
}

type ListTicketOperateRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListTicketOperateRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordHeaders) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordHeaders) SetCommonHeaders(v map[string]*string) *ListTicketOperateRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTicketOperateRecordHeaders) SetXAcsDingtalkAccessToken(v string) *ListTicketOperateRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListTicketOperateRecordRequest struct {
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
}

func (s ListTicketOperateRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordRequest) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordRequest) SetOpenTeamId(v string) *ListTicketOperateRecordRequest {
	s.OpenTeamId = &v
	return s
}

func (s *ListTicketOperateRecordRequest) SetOpenTicketId(v string) *ListTicketOperateRecordRequest {
	s.OpenTicketId = &v
	return s
}

type ListTicketOperateRecordResponseBody struct {
	// Id of the request
	Records []*ListTicketOperateRecordResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s ListTicketOperateRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBody) SetRecords(v []*ListTicketOperateRecordResponseBodyRecords) *ListTicketOperateRecordResponseBody {
	s.Records = v
	return s
}

type ListTicketOperateRecordResponseBodyRecords struct {
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	OperateData  *string `json:"operateData,omitempty" xml:"operateData,omitempty"`
	// 操作时间
	OperateTime *string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// 动作
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// 动作展示名
	OperationDisplayName *string                                               `json:"operationDisplayName,omitempty" xml:"operationDisplayName,omitempty"`
	Operator             *ListTicketOperateRecordResponseBodyRecordsOperator   `json:"operator,omitempty" xml:"operator,omitempty" type:"Struct"`
	TicketMemo           *ListTicketOperateRecordResponseBodyRecordsTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s ListTicketOperateRecordResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOpenTicketId(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OpenTicketId = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperateData(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OperateData = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperateTime(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OperateTime = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperation(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.Operation = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperationDisplayName(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OperationDisplayName = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperator(v *ListTicketOperateRecordResponseBodyRecordsOperator) *ListTicketOperateRecordResponseBodyRecords {
	s.Operator = v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetTicketMemo(v *ListTicketOperateRecordResponseBodyRecordsTicketMemo) *ListTicketOperateRecordResponseBodyRecords {
	s.TicketMemo = v
	return s
}

type ListTicketOperateRecordResponseBodyRecordsOperator struct {
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListTicketOperateRecordResponseBodyRecordsOperator) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecordsOperator) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecordsOperator) SetNickName(v string) *ListTicketOperateRecordResponseBodyRecordsOperator {
	s.NickName = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsOperator) SetUnionId(v string) *ListTicketOperateRecordResponseBodyRecordsOperator {
	s.UnionId = &v
	return s
}

type ListTicketOperateRecordResponseBodyRecordsTicketMemo struct {
	Attachments []*ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	Memo        *string                                                            `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemo) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemo) SetAttachments(v []*ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) *ListTicketOperateRecordResponseBodyRecordsTicketMemo {
	s.Attachments = v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemo) SetMemo(v string) *ListTicketOperateRecordResponseBodyRecordsTicketMemo {
	s.Memo = &v
	return s
}

type ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	Key      *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) SetFileName(v string) *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) SetKey(v string) *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments {
	s.Key = &v
	return s
}

type ListTicketOperateRecordResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTicketOperateRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTicketOperateRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTicketOperateRecordResponse) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponse) SetHeaders(v map[string]*string) *ListTicketOperateRecordResponse {
	s.Headers = v
	return s
}

func (s *ListTicketOperateRecordResponse) SetBody(v *ListTicketOperateRecordResponseBody) *ListTicketOperateRecordResponse {
	s.Body = v
	return s
}

type ListUserTeamsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListUserTeamsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListUserTeamsHeaders) GoString() string {
	return s.String()
}

func (s *ListUserTeamsHeaders) SetCommonHeaders(v map[string]*string) *ListUserTeamsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUserTeamsHeaders) SetXAcsDingtalkAccessToken(v string) *ListUserTeamsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListUserTeamsResponseBody struct {
	// teams
	Teams []*ListUserTeamsResponseBodyTeams `json:"teams,omitempty" xml:"teams,omitempty" type:"Repeated"`
}

func (s ListUserTeamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserTeamsResponseBody) SetTeams(v []*ListUserTeamsResponseBodyTeams) *ListUserTeamsResponseBody {
	s.Teams = v
	return s
}

type ListUserTeamsResponseBodyTeams struct {
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 团队名称
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
}

func (s ListUserTeamsResponseBodyTeams) String() string {
	return tea.Prettify(s)
}

func (s ListUserTeamsResponseBodyTeams) GoString() string {
	return s.String()
}

func (s *ListUserTeamsResponseBodyTeams) SetOpenTeamId(v string) *ListUserTeamsResponseBodyTeams {
	s.OpenTeamId = &v
	return s
}

func (s *ListUserTeamsResponseBodyTeams) SetTeamName(v string) *ListUserTeamsResponseBodyTeams {
	s.TeamName = &v
	return s
}

type ListUserTeamsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserTeamsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserTeamsResponse) GoString() string {
	return s.String()
}

func (s *ListUserTeamsResponse) SetHeaders(v map[string]*string) *ListUserTeamsResponse {
	s.Headers = v
	return s
}

func (s *ListUserTeamsResponse) SetBody(v *ListUserTeamsResponseBody) *ListUserTeamsResponse {
	s.Body = v
	return s
}

type QueryActiveUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryActiveUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUsersHeaders) GoString() string {
	return s.String()
}

func (s *QueryActiveUsersHeaders) SetCommonHeaders(v map[string]*string) *QueryActiveUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryActiveUsersHeaders) SetXAcsDingtalkAccessToken(v string) *QueryActiveUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryActiveUsersRequest struct {
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 查询topN的数据
	TopN *int64 `json:"topN,omitempty" xml:"topN,omitempty"`
}

func (s QueryActiveUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUsersRequest) GoString() string {
	return s.String()
}

func (s *QueryActiveUsersRequest) SetOpenConversationId(v string) *QueryActiveUsersRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryActiveUsersRequest) SetOpenTeamId(v string) *QueryActiveUsersRequest {
	s.OpenTeamId = &v
	return s
}

func (s *QueryActiveUsersRequest) SetTopN(v int64) *QueryActiveUsersRequest {
	s.TopN = &v
	return s
}

type QueryActiveUsersResponseBody struct {
	// 活跃用户列表
	ActiveUserInfos []*QueryActiveUsersResponseBodyActiveUserInfos `json:"activeUserInfos,omitempty" xml:"activeUserInfos,omitempty" type:"Repeated"`
}

func (s QueryActiveUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryActiveUsersResponseBody) SetActiveUserInfos(v []*QueryActiveUsersResponseBodyActiveUserInfos) *QueryActiveUsersResponseBody {
	s.ActiveUserInfos = v
	return s
}

type QueryActiveUsersResponseBodyActiveUserInfos struct {
	// 最近二周的行为指数
	ActionIndexL14d *float64 `json:"actionIndexL14d,omitempty" xml:"actionIndexL14d,omitempty"`
	// 最近一个月的行为指数
	ActionIndexL30d *float64 `json:"actionIndexL30d,omitempty" xml:"actionIndexL30d,omitempty"`
	// 最近一周的行为指数
	ActionIndexL7d *float64 `json:"actionIndexL7d,omitempty" xml:"actionIndexL7d,omitempty"`
	// 活跃度
	ActiveScore *float64 `json:"activeScore,omitempty" xml:"activeScore,omitempty"`
	// 昵称
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// 排名
	Ranking *int64 `json:"ranking,omitempty" xml:"ranking,omitempty"`
	// 钉钉用户unionId
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryActiveUsersResponseBodyActiveUserInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUsersResponseBodyActiveUserInfos) GoString() string {
	return s.String()
}

func (s *QueryActiveUsersResponseBodyActiveUserInfos) SetActionIndexL14d(v float64) *QueryActiveUsersResponseBodyActiveUserInfos {
	s.ActionIndexL14d = &v
	return s
}

func (s *QueryActiveUsersResponseBodyActiveUserInfos) SetActionIndexL30d(v float64) *QueryActiveUsersResponseBodyActiveUserInfos {
	s.ActionIndexL30d = &v
	return s
}

func (s *QueryActiveUsersResponseBodyActiveUserInfos) SetActionIndexL7d(v float64) *QueryActiveUsersResponseBodyActiveUserInfos {
	s.ActionIndexL7d = &v
	return s
}

func (s *QueryActiveUsersResponseBodyActiveUserInfos) SetActiveScore(v float64) *QueryActiveUsersResponseBodyActiveUserInfos {
	s.ActiveScore = &v
	return s
}

func (s *QueryActiveUsersResponseBodyActiveUserInfos) SetNickName(v string) *QueryActiveUsersResponseBodyActiveUserInfos {
	s.NickName = &v
	return s
}

func (s *QueryActiveUsersResponseBodyActiveUserInfos) SetRanking(v int64) *QueryActiveUsersResponseBodyActiveUserInfos {
	s.Ranking = &v
	return s
}

func (s *QueryActiveUsersResponseBodyActiveUserInfos) SetUnionId(v string) *QueryActiveUsersResponseBodyActiveUserInfos {
	s.UnionId = &v
	return s
}

type QueryActiveUsersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryActiveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryActiveUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryActiveUsersResponse) GoString() string {
	return s.String()
}

func (s *QueryActiveUsersResponse) SetHeaders(v map[string]*string) *QueryActiveUsersResponse {
	s.Headers = v
	return s
}

func (s *QueryActiveUsersResponse) SetBody(v *QueryActiveUsersResponseBody) *QueryActiveUsersResponse {
	s.Body = v
	return s
}

type QueryGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupRequest struct {
	// 业务关联ID，和开放群ID二选一传
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s QueryGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupRequest) SetBizId(v string) *QueryGroupRequest {
	s.BizId = &v
	return s
}

func (s *QueryGroupRequest) SetOpenConversationId(v string) *QueryGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryGroupRequest) SetOpenTeamId(v string) *QueryGroupRequest {
	s.OpenTeamId = &v
	return s
}

type QueryGroupResponseBody struct {
	// 群bizId
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 入群URL
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 开放群组ID
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 服务群机器人code
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// 服务群机器人名称
	RobotName *string `json:"robotName,omitempty" xml:"robotName,omitempty"`
}

func (s QueryGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupResponseBody) SetBizId(v string) *QueryGroupResponseBody {
	s.BizId = &v
	return s
}

func (s *QueryGroupResponseBody) SetGroupName(v string) *QueryGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *QueryGroupResponseBody) SetGroupUrl(v string) *QueryGroupResponseBody {
	s.GroupUrl = &v
	return s
}

func (s *QueryGroupResponseBody) SetOpenConversationId(v string) *QueryGroupResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *QueryGroupResponseBody) SetOpenGroupSetId(v string) *QueryGroupResponseBody {
	s.OpenGroupSetId = &v
	return s
}

func (s *QueryGroupResponseBody) SetOpenTeamId(v string) *QueryGroupResponseBody {
	s.OpenTeamId = &v
	return s
}

func (s *QueryGroupResponseBody) SetRobotCode(v string) *QueryGroupResponseBody {
	s.RobotCode = &v
	return s
}

func (s *QueryGroupResponseBody) SetRobotName(v string) *QueryGroupResponseBody {
	s.RobotName = &v
	return s
}

type QueryGroupResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupResponse) SetHeaders(v map[string]*string) *QueryGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupResponse) SetBody(v *QueryGroupResponseBody) *QueryGroupResponse {
	s.Body = v
	return s
}

type QueryGroupSetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupSetHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupSetHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupSetHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupSetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupSetHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupSetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupSetRequest struct {
	// openTeamId
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s QueryGroupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupSetRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupSetRequest) SetOpenTeamId(v string) *QueryGroupSetRequest {
	s.OpenTeamId = &v
	return s
}

type QueryGroupSetResponseBody struct {
	// records
	Records []*QueryGroupSetResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s QueryGroupSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupSetResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupSetResponseBody) SetRecords(v []*QueryGroupSetResponseBodyRecords) *QueryGroupSetResponseBody {
	s.Records = v
	return s
}

type QueryGroupSetResponseBodyRecords struct {
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	GroupSetName   *string `json:"groupSetName,omitempty" xml:"groupSetName,omitempty"`
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	TemplateId     *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s QueryGroupSetResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupSetResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *QueryGroupSetResponseBodyRecords) SetGmtCreate(v string) *QueryGroupSetResponseBodyRecords {
	s.GmtCreate = &v
	return s
}

func (s *QueryGroupSetResponseBodyRecords) SetGmtModified(v string) *QueryGroupSetResponseBodyRecords {
	s.GmtModified = &v
	return s
}

func (s *QueryGroupSetResponseBodyRecords) SetGroupSetName(v string) *QueryGroupSetResponseBodyRecords {
	s.GroupSetName = &v
	return s
}

func (s *QueryGroupSetResponseBodyRecords) SetOpenGroupSetId(v string) *QueryGroupSetResponseBodyRecords {
	s.OpenGroupSetId = &v
	return s
}

func (s *QueryGroupSetResponseBodyRecords) SetTemplateId(v string) *QueryGroupSetResponseBodyRecords {
	s.TemplateId = &v
	return s
}

type QueryGroupSetResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupSetResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupSetResponse) SetHeaders(v map[string]*string) *QueryGroupSetResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupSetResponse) SetBody(v *QueryGroupSetResponseBody) *QueryGroupSetResponse {
	s.Body = v
	return s
}

type QuerySendMsgTaskStatisticsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySendMsgTaskStatisticsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsHeaders) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsHeaders) SetCommonHeaders(v map[string]*string) *QuerySendMsgTaskStatisticsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySendMsgTaskStatisticsHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySendMsgTaskStatisticsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySendMsgTaskStatisticsRequest struct {
	// 每页条数
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 游标，首页为空
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 开放群发任务ID
	OpenBatchTaskId *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s QuerySendMsgTaskStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsRequest) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsRequest) SetMaxResults(v int64) *QuerySendMsgTaskStatisticsRequest {
	s.MaxResults = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsRequest) SetNextToken(v string) *QuerySendMsgTaskStatisticsRequest {
	s.NextToken = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsRequest) SetOpenBatchTaskId(v string) *QuerySendMsgTaskStatisticsRequest {
	s.OpenBatchTaskId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsRequest) SetOpenTeamId(v string) *QuerySendMsgTaskStatisticsRequest {
	s.OpenTeamId = &v
	return s
}

type QuerySendMsgTaskStatisticsResponseBody struct {
	MaxResults *int64                                           `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                                          `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Records    []*QuerySendMsgTaskStatisticsResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// Id of the request
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QuerySendMsgTaskStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsResponseBody) SetMaxResults(v int64) *QuerySendMsgTaskStatisticsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBody) SetNextToken(v string) *QuerySendMsgTaskStatisticsResponseBody {
	s.NextToken = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBody) SetRecords(v []*QuerySendMsgTaskStatisticsResponseBodyRecords) *QuerySendMsgTaskStatisticsResponseBody {
	s.Records = v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBody) SetTotalCount(v int64) *QuerySendMsgTaskStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

type QuerySendMsgTaskStatisticsResponseBodyRecords struct {
	ErrorDetail             *string                                                               `json:"errorDetail,omitempty" xml:"errorDetail,omitempty"`
	Group                   *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup                   `json:"group,omitempty" xml:"group,omitempty" type:"Struct"`
	GroupUserReadStatistics *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics `json:"groupUserReadStatistics,omitempty" xml:"groupUserReadStatistics,omitempty" type:"Struct"`
	OpenMsgId               *string                                                               `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
	Status                  *string                                                               `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QuerySendMsgTaskStatisticsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecords) SetErrorDetail(v string) *QuerySendMsgTaskStatisticsResponseBodyRecords {
	s.ErrorDetail = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecords) SetGroup(v *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup) *QuerySendMsgTaskStatisticsResponseBodyRecords {
	s.Group = v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecords) SetGroupUserReadStatistics(v *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics) *QuerySendMsgTaskStatisticsResponseBodyRecords {
	s.GroupUserReadStatistics = v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecords) SetOpenMsgId(v string) *QuerySendMsgTaskStatisticsResponseBodyRecords {
	s.OpenMsgId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecords) SetStatus(v string) *QuerySendMsgTaskStatisticsResponseBodyRecords {
	s.Status = &v
	return s
}

type QuerySendMsgTaskStatisticsResponseBodyRecordsGroup struct {
	BizId              *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	GroupName          *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupSetName       *string `json:"groupSetName,omitempty" xml:"groupSetName,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenGroupSetId     *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	OpenTeamId         *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s QuerySendMsgTaskStatisticsResponseBodyRecordsGroup) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsResponseBodyRecordsGroup) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup) SetBizId(v string) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup {
	s.BizId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup) SetGroupName(v string) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup {
	s.GroupName = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup) SetGroupSetName(v string) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup {
	s.GroupSetName = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup) SetOpenConversationId(v string) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup {
	s.OpenConversationId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup) SetOpenGroupSetId(v string) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup {
	s.OpenGroupSetId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup) SetOpenTeamId(v string) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroup {
	s.OpenTeamId = &v
	return s
}

type QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics struct {
	OpenBatchTaskId    *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ReadUserInc        *int64  `json:"readUserInc,omitempty" xml:"readUserInc,omitempty"`
	SendUserInc        *int64  `json:"sendUserInc,omitempty" xml:"sendUserInc,omitempty"`
	UnReadUserInc      *int64  `json:"unReadUserInc,omitempty" xml:"unReadUserInc,omitempty"`
}

func (s QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics) SetOpenBatchTaskId(v string) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics {
	s.OpenBatchTaskId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics) SetOpenConversationId(v string) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics {
	s.OpenConversationId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics) SetReadUserInc(v int64) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics {
	s.ReadUserInc = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics) SetSendUserInc(v int64) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics {
	s.SendUserInc = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics) SetUnReadUserInc(v int64) *QuerySendMsgTaskStatisticsResponseBodyRecordsGroupUserReadStatistics {
	s.UnReadUserInc = &v
	return s
}

type QuerySendMsgTaskStatisticsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySendMsgTaskStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySendMsgTaskStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsResponse) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsResponse) SetHeaders(v map[string]*string) *QuerySendMsgTaskStatisticsResponse {
	s.Headers = v
	return s
}

func (s *QuerySendMsgTaskStatisticsResponse) SetBody(v *QuerySendMsgTaskStatisticsResponseBody) *QuerySendMsgTaskStatisticsResponse {
	s.Body = v
	return s
}

type QuerySendMsgTaskStatisticsDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySendMsgTaskStatisticsDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsDetailHeaders) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsDetailHeaders) SetCommonHeaders(v map[string]*string) *QuerySendMsgTaskStatisticsDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySendMsgTaskStatisticsDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySendMsgTaskStatisticsDetailRequest struct {
	// 每页条数
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 游标，首页传递空
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 开放群发任务ID
	OpenBatchTaskId *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
	// 开放会话ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s QuerySendMsgTaskStatisticsDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsDetailRequest) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsDetailRequest) SetMaxResults(v int64) *QuerySendMsgTaskStatisticsDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailRequest) SetNextToken(v string) *QuerySendMsgTaskStatisticsDetailRequest {
	s.NextToken = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailRequest) SetOpenBatchTaskId(v string) *QuerySendMsgTaskStatisticsDetailRequest {
	s.OpenBatchTaskId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailRequest) SetOpenConversationId(v string) *QuerySendMsgTaskStatisticsDetailRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailRequest) SetOpenTeamId(v string) *QuerySendMsgTaskStatisticsDetailRequest {
	s.OpenTeamId = &v
	return s
}

type QuerySendMsgTaskStatisticsDetailResponseBody struct {
	MaxResults *int64                                                 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Records    []*QuerySendMsgTaskStatisticsDetailResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// Id of the request
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QuerySendMsgTaskStatisticsDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsDetailResponseBody) SetMaxResults(v int64) *QuerySendMsgTaskStatisticsDetailResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailResponseBody) SetNextToken(v string) *QuerySendMsgTaskStatisticsDetailResponseBody {
	s.NextToken = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailResponseBody) SetRecords(v []*QuerySendMsgTaskStatisticsDetailResponseBodyRecords) *QuerySendMsgTaskStatisticsDetailResponseBody {
	s.Records = v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailResponseBody) SetTotalCount(v int64) *QuerySendMsgTaskStatisticsDetailResponseBody {
	s.TotalCount = &v
	return s
}

type QuerySendMsgTaskStatisticsDetailResponseBodyRecords struct {
	OpenBatchTaskId    *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ReadStatus         *int64  `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
	ReadTimeStr        *string `json:"readTimeStr,omitempty" xml:"readTimeStr,omitempty"`
	ReceiverName       *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	ReceiverUnionId    *string `json:"receiverUnionId,omitempty" xml:"receiverUnionId,omitempty"`
}

func (s QuerySendMsgTaskStatisticsDetailResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsDetailResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsDetailResponseBodyRecords) SetOpenBatchTaskId(v string) *QuerySendMsgTaskStatisticsDetailResponseBodyRecords {
	s.OpenBatchTaskId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailResponseBodyRecords) SetOpenConversationId(v string) *QuerySendMsgTaskStatisticsDetailResponseBodyRecords {
	s.OpenConversationId = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailResponseBodyRecords) SetReadStatus(v int64) *QuerySendMsgTaskStatisticsDetailResponseBodyRecords {
	s.ReadStatus = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailResponseBodyRecords) SetReadTimeStr(v string) *QuerySendMsgTaskStatisticsDetailResponseBodyRecords {
	s.ReadTimeStr = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailResponseBodyRecords) SetReceiverName(v string) *QuerySendMsgTaskStatisticsDetailResponseBodyRecords {
	s.ReceiverName = &v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailResponseBodyRecords) SetReceiverUnionId(v string) *QuerySendMsgTaskStatisticsDetailResponseBodyRecords {
	s.ReceiverUnionId = &v
	return s
}

type QuerySendMsgTaskStatisticsDetailResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySendMsgTaskStatisticsDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySendMsgTaskStatisticsDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySendMsgTaskStatisticsDetailResponse) GoString() string {
	return s.String()
}

func (s *QuerySendMsgTaskStatisticsDetailResponse) SetHeaders(v map[string]*string) *QuerySendMsgTaskStatisticsDetailResponse {
	s.Headers = v
	return s
}

func (s *QuerySendMsgTaskStatisticsDetailResponse) SetBody(v *QuerySendMsgTaskStatisticsDetailResponseBody) *QuerySendMsgTaskStatisticsDetailResponse {
	s.Body = v
	return s
}

type QueryServiceGroupMessageReadStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryServiceGroupMessageReadStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceGroupMessageReadStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryServiceGroupMessageReadStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryServiceGroupMessageReadStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryServiceGroupMessageReadStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryServiceGroupMessageReadStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryServiceGroupMessageReadStatusRequest struct {
	// 本次读取的最大数据记录数量，此参数为可选参数，用户传入为空时，应该有默认值。应设置最大值限制，最大不超过100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 用来标记当前开始读取的位置，置空表示从头开始。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 开放消息ID
	OpenMsgTaskId *string `json:"openMsgTaskId,omitempty" xml:"openMsgTaskId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s QueryServiceGroupMessageReadStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceGroupMessageReadStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetMaxResults(v int32) *QueryServiceGroupMessageReadStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetNextToken(v string) *QueryServiceGroupMessageReadStatusRequest {
	s.NextToken = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetOpenConversationId(v string) *QueryServiceGroupMessageReadStatusRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetOpenMsgTaskId(v string) *QueryServiceGroupMessageReadStatusRequest {
	s.OpenMsgTaskId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusRequest) SetOpenTeamId(v string) *QueryServiceGroupMessageReadStatusRequest {
	s.OpenTeamId = &v
	return s
}

type QueryServiceGroupMessageReadStatusResponseBody struct {
	// 本次请求所返回的最大记录条数。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 已读未读信息列表
	Records []*QueryServiceGroupMessageReadStatusResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// 本次请求条件下的数据总量，此参数为可选参数，默认可不返回。本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryServiceGroupMessageReadStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceGroupMessageReadStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServiceGroupMessageReadStatusResponseBody) SetMaxResults(v int32) *QueryServiceGroupMessageReadStatusResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBody) SetNextToken(v string) *QueryServiceGroupMessageReadStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBody) SetRecords(v []*QueryServiceGroupMessageReadStatusResponseBodyRecords) *QueryServiceGroupMessageReadStatusResponseBody {
	s.Records = v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBody) SetTotalCount(v int32) *QueryServiceGroupMessageReadStatusResponseBody {
	s.TotalCount = &v
	return s
}

type QueryServiceGroupMessageReadStatusResponseBodyRecords struct {
	// 状态：已读1/未读0
	ReadStatus *int32 `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
	// 已读时间
	ReadTimeStr *string `json:"readTimeStr,omitempty" xml:"readTimeStr,omitempty"`
	// 接收者dingtalkId
	ReceiverDingTalkId *string `json:"receiverDingTalkId,omitempty" xml:"receiverDingTalkId,omitempty"`
	// 接收者昵称
	ReceiverName *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	// 已读人员为非企业员工则有值
	ReceiverUnionId *string `json:"receiverUnionId,omitempty" xml:"receiverUnionId,omitempty"`
	// 已读人员为企业员工则有值
	ReceiverUserId *string `json:"receiverUserId,omitempty" xml:"receiverUserId,omitempty"`
	// 发送时间
	SendTimeStr *string `json:"sendTimeStr,omitempty" xml:"sendTimeStr,omitempty"`
}

func (s QueryServiceGroupMessageReadStatusResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceGroupMessageReadStatusResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReadStatus(v int32) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReadStatus = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReadTimeStr(v string) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReadTimeStr = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReceiverDingTalkId(v string) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReceiverDingTalkId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReceiverName(v string) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReceiverName = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReceiverUnionId(v string) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReceiverUnionId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetReceiverUserId(v string) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.ReceiverUserId = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponseBodyRecords) SetSendTimeStr(v string) *QueryServiceGroupMessageReadStatusResponseBodyRecords {
	s.SendTimeStr = &v
	return s
}

type QueryServiceGroupMessageReadStatusResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryServiceGroupMessageReadStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServiceGroupMessageReadStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceGroupMessageReadStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryServiceGroupMessageReadStatusResponse) SetHeaders(v map[string]*string) *QueryServiceGroupMessageReadStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponse) SetBody(v *QueryServiceGroupMessageReadStatusResponseBody) *QueryServiceGroupMessageReadStatusResponse {
	s.Body = v
	return s
}

type ResubmitTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ResubmitTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s ResubmitTicketHeaders) GoString() string {
	return s.String()
}

func (s *ResubmitTicketHeaders) SetCommonHeaders(v map[string]*string) *ResubmitTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ResubmitTicketHeaders) SetXAcsDingtalkAccessToken(v string) *ResubmitTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ResubmitTicketRequest struct {
	// 工单创建人UnionId
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// 自定义组件字段值(JSON格式)
	CustomFields *string                      `json:"customFields,omitempty" xml:"customFields,omitempty"`
	Notify       *ResubmitTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单模板业务ID
	OpenTemplateBizId *string `json:"openTemplateBizId,omitempty" xml:"openTemplateBizId,omitempty"`
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 工单处理人UnionId列表
	ProcessorUnionIds []*string `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	// 工单场景 SG 或 VOC
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 工单场景信息
	SceneContext *ResubmitTicketRequestSceneContext `json:"sceneContext,omitempty" xml:"sceneContext,omitempty" type:"Struct"`
	// 备注
	TicketMemo *ResubmitTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
	// 工单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ResubmitTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s ResubmitTicketRequest) GoString() string {
	return s.String()
}

func (s *ResubmitTicketRequest) SetCreatorUnionId(v string) *ResubmitTicketRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *ResubmitTicketRequest) SetCustomFields(v string) *ResubmitTicketRequest {
	s.CustomFields = &v
	return s
}

func (s *ResubmitTicketRequest) SetNotify(v *ResubmitTicketRequestNotify) *ResubmitTicketRequest {
	s.Notify = v
	return s
}

func (s *ResubmitTicketRequest) SetOpenTeamId(v string) *ResubmitTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *ResubmitTicketRequest) SetOpenTemplateBizId(v string) *ResubmitTicketRequest {
	s.OpenTemplateBizId = &v
	return s
}

func (s *ResubmitTicketRequest) SetOpenTicketId(v string) *ResubmitTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *ResubmitTicketRequest) SetProcessorUnionIds(v []*string) *ResubmitTicketRequest {
	s.ProcessorUnionIds = v
	return s
}

func (s *ResubmitTicketRequest) SetScene(v string) *ResubmitTicketRequest {
	s.Scene = &v
	return s
}

func (s *ResubmitTicketRequest) SetSceneContext(v *ResubmitTicketRequestSceneContext) *ResubmitTicketRequest {
	s.SceneContext = v
	return s
}

func (s *ResubmitTicketRequest) SetTicketMemo(v *ResubmitTicketRequestTicketMemo) *ResubmitTicketRequest {
	s.TicketMemo = v
	return s
}

func (s *ResubmitTicketRequest) SetTitle(v string) *ResubmitTicketRequest {
	s.Title = &v
	return s
}

type ResubmitTicketRequestNotify struct {
	// 服务群通知接收人（钉钉UnionId）
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember *bool `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
	// 企业工作通知接收人（钉钉UnionId）
	WorkNoticeReceiverUnionIds []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
}

func (s ResubmitTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s ResubmitTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *ResubmitTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *ResubmitTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *ResubmitTicketRequestNotify) SetNoticeAllGroupMember(v bool) *ResubmitTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *ResubmitTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *ResubmitTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

type ResubmitTicketRequestSceneContext struct {
	GroupMsgs []*ResubmitTicketRequestSceneContextGroupMsgs `json:"groupMsgs,omitempty" xml:"groupMsgs,omitempty" type:"Repeated"`
	// 服务群openConversationId
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 工单相关人UnionId列表
	RelevantorUnionIds []*string `json:"relevantorUnionIds,omitempty" xml:"relevantorUnionIds,omitempty" type:"Repeated"`
}

func (s ResubmitTicketRequestSceneContext) String() string {
	return tea.Prettify(s)
}

func (s ResubmitTicketRequestSceneContext) GoString() string {
	return s.String()
}

func (s *ResubmitTicketRequestSceneContext) SetGroupMsgs(v []*ResubmitTicketRequestSceneContextGroupMsgs) *ResubmitTicketRequestSceneContext {
	s.GroupMsgs = v
	return s
}

func (s *ResubmitTicketRequestSceneContext) SetOpenConversationId(v string) *ResubmitTicketRequestSceneContext {
	s.OpenConversationId = &v
	return s
}

func (s *ResubmitTicketRequestSceneContext) SetRelevantorUnionIds(v []*string) *ResubmitTicketRequestSceneContext {
	s.RelevantorUnionIds = v
	return s
}

type ResubmitTicketRequestSceneContextGroupMsgs struct {
	Anchor *bool `json:"anchor,omitempty" xml:"anchor,omitempty"`
	// 勾选消息openMsgId
	OpenMsgId *string `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
	TopicId   *string `json:"topicId,omitempty" xml:"topicId,omitempty"`
}

func (s ResubmitTicketRequestSceneContextGroupMsgs) String() string {
	return tea.Prettify(s)
}

func (s ResubmitTicketRequestSceneContextGroupMsgs) GoString() string {
	return s.String()
}

func (s *ResubmitTicketRequestSceneContextGroupMsgs) SetAnchor(v bool) *ResubmitTicketRequestSceneContextGroupMsgs {
	s.Anchor = &v
	return s
}

func (s *ResubmitTicketRequestSceneContextGroupMsgs) SetOpenMsgId(v string) *ResubmitTicketRequestSceneContextGroupMsgs {
	s.OpenMsgId = &v
	return s
}

func (s *ResubmitTicketRequestSceneContextGroupMsgs) SetTopicId(v string) *ResubmitTicketRequestSceneContextGroupMsgs {
	s.TopicId = &v
	return s
}

type ResubmitTicketRequestTicketMemo struct {
	Attachments []*ResubmitTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// 备注文字
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s ResubmitTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s ResubmitTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *ResubmitTicketRequestTicketMemo) SetAttachments(v []*ResubmitTicketRequestTicketMemoAttachments) *ResubmitTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *ResubmitTicketRequestTicketMemo) SetMemo(v string) *ResubmitTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

type ResubmitTicketRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ResubmitTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s ResubmitTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *ResubmitTicketRequestTicketMemoAttachments) SetFileName(v string) *ResubmitTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *ResubmitTicketRequestTicketMemoAttachments) SetKey(v string) *ResubmitTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type ResubmitTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ResubmitTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s ResubmitTicketResponse) GoString() string {
	return s.String()
}

func (s *ResubmitTicketResponse) SetHeaders(v map[string]*string) *ResubmitTicketResponse {
	s.Headers = v
	return s
}

type RetractTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RetractTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s RetractTicketHeaders) GoString() string {
	return s.String()
}

func (s *RetractTicketHeaders) SetCommonHeaders(v map[string]*string) *RetractTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RetractTicketHeaders) SetXAcsDingtalkAccessToken(v string) *RetractTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RetractTicketRequest struct {
	Notify     *RetractTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	OpenTeamId *string                     `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 操作人ID
	OperatorUnionId *string                         `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	TicketMemo      *RetractTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s RetractTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s RetractTicketRequest) GoString() string {
	return s.String()
}

func (s *RetractTicketRequest) SetNotify(v *RetractTicketRequestNotify) *RetractTicketRequest {
	s.Notify = v
	return s
}

func (s *RetractTicketRequest) SetOpenTeamId(v string) *RetractTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *RetractTicketRequest) SetOpenTicketId(v string) *RetractTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *RetractTicketRequest) SetOperatorUnionId(v string) *RetractTicketRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *RetractTicketRequest) SetTicketMemo(v *RetractTicketRequestTicketMemo) *RetractTicketRequest {
	s.TicketMemo = v
	return s
}

type RetractTicketRequestNotify struct {
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember       *bool     `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
	WorkNoticeReceiverUnionIds []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
}

func (s RetractTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s RetractTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *RetractTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *RetractTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *RetractTicketRequestNotify) SetNoticeAllGroupMember(v bool) *RetractTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *RetractTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *RetractTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

type RetractTicketRequestTicketMemo struct {
	Attachments []*RetractTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	Memo        *string                                      `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s RetractTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s RetractTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *RetractTicketRequestTicketMemo) SetAttachments(v []*RetractTicketRequestTicketMemoAttachments) *RetractTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *RetractTicketRequestTicketMemo) SetMemo(v string) *RetractTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

type RetractTicketRequestTicketMemoAttachments struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	Key      *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s RetractTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s RetractTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *RetractTicketRequestTicketMemoAttachments) SetFileName(v string) *RetractTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *RetractTicketRequestTicketMemoAttachments) SetKey(v string) *RetractTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type RetractTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RetractTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s RetractTicketResponse) GoString() string {
	return s.String()
}

func (s *RetractTicketResponse) SetHeaders(v map[string]*string) *RetractTicketResponse {
	s.Headers = v
	return s
}

type SearchGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchGroupHeaders) GoString() string {
	return s.String()
}

func (s *SearchGroupHeaders) SetCommonHeaders(v map[string]*string) *SearchGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchGroupHeaders) SetXAcsDingtalkAccessToken(v string) *SearchGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchGroupRequest struct {
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 本次读取的最大数据记录数量，此参数为可选参数，用户传入为空时，应该有默认值。应设置最大值限制，最大不超过100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 用来标记当前开始读取的位置，置空表示从头开始。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 开群组ID
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 搜索类型
	SearchType *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
}

func (s SearchGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchGroupRequest) GoString() string {
	return s.String()
}

func (s *SearchGroupRequest) SetGroupName(v string) *SearchGroupRequest {
	s.GroupName = &v
	return s
}

func (s *SearchGroupRequest) SetMaxResults(v int32) *SearchGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchGroupRequest) SetNextToken(v string) *SearchGroupRequest {
	s.NextToken = &v
	return s
}

func (s *SearchGroupRequest) SetOpenConversationId(v string) *SearchGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SearchGroupRequest) SetOpenGroupSetId(v string) *SearchGroupRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *SearchGroupRequest) SetOpenTeamId(v string) *SearchGroupRequest {
	s.OpenTeamId = &v
	return s
}

func (s *SearchGroupRequest) SetSearchType(v string) *SearchGroupRequest {
	s.SearchType = &v
	return s
}

type SearchGroupResponseBody struct {
	// 本次请求所返回的最大记录条数。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 已读未读信息列表
	Records []*SearchGroupResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// 本次请求条件下的数据总量，此参数为可选参数，默认可不返回。本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SearchGroupResponseBody) SetMaxResults(v int32) *SearchGroupResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchGroupResponseBody) SetNextToken(v string) *SearchGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchGroupResponseBody) SetRecords(v []*SearchGroupResponseBodyRecords) *SearchGroupResponseBody {
	s.Records = v
	return s
}

func (s *SearchGroupResponseBody) SetTotalCount(v int32) *SearchGroupResponseBody {
	s.TotalCount = &v
	return s
}

type SearchGroupResponseBodyRecords struct {
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 入群链接
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	// 开放群ID
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 开放群组ID
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s SearchGroupResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s SearchGroupResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *SearchGroupResponseBodyRecords) SetGroupName(v string) *SearchGroupResponseBodyRecords {
	s.GroupName = &v
	return s
}

func (s *SearchGroupResponseBodyRecords) SetGroupUrl(v string) *SearchGroupResponseBodyRecords {
	s.GroupUrl = &v
	return s
}

func (s *SearchGroupResponseBodyRecords) SetOpenConversationId(v string) *SearchGroupResponseBodyRecords {
	s.OpenConversationId = &v
	return s
}

func (s *SearchGroupResponseBodyRecords) SetOpenGroupSetId(v string) *SearchGroupResponseBodyRecords {
	s.OpenGroupSetId = &v
	return s
}

func (s *SearchGroupResponseBodyRecords) SetOpenTeamId(v string) *SearchGroupResponseBodyRecords {
	s.OpenTeamId = &v
	return s
}

type SearchGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchGroupResponse) GoString() string {
	return s.String()
}

func (s *SearchGroupResponse) SetHeaders(v map[string]*string) *SearchGroupResponse {
	s.Headers = v
	return s
}

func (s *SearchGroupResponse) SetBody(v *SearchGroupResponseBody) *SearchGroupResponse {
	s.Body = v
	return s
}

type SendMsgByTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendMsgByTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskHeaders) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskHeaders) SetCommonHeaders(v map[string]*string) *SendMsgByTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMsgByTaskHeaders) SetXAcsDingtalkAccessToken(v string) *SendMsgByTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendMsgByTaskRequest struct {
	// 群发内容
	MessageContent *SendMsgByTaskRequestMessageContent `json:"messageContent,omitempty" xml:"messageContent,omitempty" type:"Struct"`
	// 开放团队ID
	OpenTeamId *string                         `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	QueryGroup *SendMsgByTaskRequestQueryGroup `json:"queryGroup,omitempty" xml:"queryGroup,omitempty" type:"Struct"`
	// 发送配置
	SendConfig *SendMsgByTaskRequestSendConfig `json:"sendConfig,omitempty" xml:"sendConfig,omitempty" type:"Struct"`
	// 群发任务名称
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s SendMsgByTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskRequest) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskRequest) SetMessageContent(v *SendMsgByTaskRequestMessageContent) *SendMsgByTaskRequest {
	s.MessageContent = v
	return s
}

func (s *SendMsgByTaskRequest) SetOpenTeamId(v string) *SendMsgByTaskRequest {
	s.OpenTeamId = &v
	return s
}

func (s *SendMsgByTaskRequest) SetQueryGroup(v *SendMsgByTaskRequestQueryGroup) *SendMsgByTaskRequest {
	s.QueryGroup = v
	return s
}

func (s *SendMsgByTaskRequest) SetSendConfig(v *SendMsgByTaskRequestSendConfig) *SendMsgByTaskRequest {
	s.SendConfig = v
	return s
}

func (s *SendMsgByTaskRequest) SetTaskName(v string) *SendMsgByTaskRequest {
	s.TaskName = &v
	return s
}

type SendMsgByTaskRequestMessageContent struct {
	// at活跃成员数量
	AtActiveMemberNum *int64 `json:"atActiveMemberNum,omitempty" xml:"atActiveMemberNum,omitempty"`
	// 是否At活跃成员
	AtActiveUser *bool `json:"atActiveUser,omitempty" xml:"atActiveUser,omitempty"`
	// 是否At全部人员
	AtAll *bool                                     `json:"atAll,omitempty" xml:"atAll,omitempty"`
	Btns  []*SendMsgByTaskRequestMessageContentBtns `json:"btns,omitempty" xml:"btns,omitempty" type:"Repeated"`
	// 内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 图片列表
	Images []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// 消息类型
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendMsgByTaskRequestMessageContent) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskRequestMessageContent) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskRequestMessageContent) SetAtActiveMemberNum(v int64) *SendMsgByTaskRequestMessageContent {
	s.AtActiveMemberNum = &v
	return s
}

func (s *SendMsgByTaskRequestMessageContent) SetAtActiveUser(v bool) *SendMsgByTaskRequestMessageContent {
	s.AtActiveUser = &v
	return s
}

func (s *SendMsgByTaskRequestMessageContent) SetAtAll(v bool) *SendMsgByTaskRequestMessageContent {
	s.AtAll = &v
	return s
}

func (s *SendMsgByTaskRequestMessageContent) SetBtns(v []*SendMsgByTaskRequestMessageContentBtns) *SendMsgByTaskRequestMessageContent {
	s.Btns = v
	return s
}

func (s *SendMsgByTaskRequestMessageContent) SetContent(v string) *SendMsgByTaskRequestMessageContent {
	s.Content = &v
	return s
}

func (s *SendMsgByTaskRequestMessageContent) SetImages(v []*string) *SendMsgByTaskRequestMessageContent {
	s.Images = v
	return s
}

func (s *SendMsgByTaskRequestMessageContent) SetMessageType(v string) *SendMsgByTaskRequestMessageContent {
	s.MessageType = &v
	return s
}

func (s *SendMsgByTaskRequestMessageContent) SetTitle(v string) *SendMsgByTaskRequestMessageContent {
	s.Title = &v
	return s
}

type SendMsgByTaskRequestMessageContentBtns struct {
	ActionURL *string `json:"actionURL,omitempty" xml:"actionURL,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendMsgByTaskRequestMessageContentBtns) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskRequestMessageContentBtns) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskRequestMessageContentBtns) SetActionURL(v string) *SendMsgByTaskRequestMessageContentBtns {
	s.ActionURL = &v
	return s
}

func (s *SendMsgByTaskRequestMessageContentBtns) SetTitle(v string) *SendMsgByTaskRequestMessageContentBtns {
	s.Title = &v
	return s
}

type SendMsgByTaskRequestQueryGroup struct {
	// 群标签
	GroupTagNames []*string `json:"groupTagNames,omitempty" xml:"groupTagNames,omitempty" type:"Repeated"`
	// 活跃日期筛选类型，ACTIVE=活跃      NOTACTIVE=不活跃
	LastActiveDateFilterType *string `json:"lastActiveDateFilterType,omitempty" xml:"lastActiveDateFilterType,omitempty"`
	// 最近活跃时间的结束时间
	LastActiveTimeEnd *string `json:"lastActiveTimeEnd,omitempty" xml:"lastActiveTimeEnd,omitempty"`
	// 最近活跃时间的开始时间
	LastActiveTimeStart *string `json:"lastActiveTimeStart,omitempty" xml:"lastActiveTimeStart,omitempty"`
	// 精准圈选-群ID集合
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	// 开放群组ID
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 群发圈选类型 1. AIMED 精准圈选 2. MULTI_CONDITIONS 多条件圈选
	QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
}

func (s SendMsgByTaskRequestQueryGroup) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskRequestQueryGroup) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskRequestQueryGroup) SetGroupTagNames(v []*string) *SendMsgByTaskRequestQueryGroup {
	s.GroupTagNames = v
	return s
}

func (s *SendMsgByTaskRequestQueryGroup) SetLastActiveDateFilterType(v string) *SendMsgByTaskRequestQueryGroup {
	s.LastActiveDateFilterType = &v
	return s
}

func (s *SendMsgByTaskRequestQueryGroup) SetLastActiveTimeEnd(v string) *SendMsgByTaskRequestQueryGroup {
	s.LastActiveTimeEnd = &v
	return s
}

func (s *SendMsgByTaskRequestQueryGroup) SetLastActiveTimeStart(v string) *SendMsgByTaskRequestQueryGroup {
	s.LastActiveTimeStart = &v
	return s
}

func (s *SendMsgByTaskRequestQueryGroup) SetOpenConversationIds(v []*string) *SendMsgByTaskRequestQueryGroup {
	s.OpenConversationIds = v
	return s
}

func (s *SendMsgByTaskRequestQueryGroup) SetOpenGroupSetId(v string) *SendMsgByTaskRequestQueryGroup {
	s.OpenGroupSetId = &v
	return s
}

func (s *SendMsgByTaskRequestQueryGroup) SetQueryType(v string) *SendMsgByTaskRequestQueryGroup {
	s.QueryType = &v
	return s
}

type SendMsgByTaskRequestSendConfig struct {
	// 是否链接追踪
	NeedUrlTrack *bool `json:"needUrlTrack,omitempty" xml:"needUrlTrack,omitempty"`
	// 执行时间（sendType=TIMING时传入）
	SendTime *string `json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	// 发送类型      * TIMING=定时执行      * INSTANT=立即执行
	SendType *string `json:"sendType,omitempty" xml:"sendType,omitempty"`
	// 链接跟踪配置
	UrlTrackConfig []*SendMsgByTaskRequestSendConfigUrlTrackConfig `json:"urlTrackConfig,omitempty" xml:"urlTrackConfig,omitempty" type:"Repeated"`
}

func (s SendMsgByTaskRequestSendConfig) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskRequestSendConfig) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskRequestSendConfig) SetNeedUrlTrack(v bool) *SendMsgByTaskRequestSendConfig {
	s.NeedUrlTrack = &v
	return s
}

func (s *SendMsgByTaskRequestSendConfig) SetSendTime(v string) *SendMsgByTaskRequestSendConfig {
	s.SendTime = &v
	return s
}

func (s *SendMsgByTaskRequestSendConfig) SetSendType(v string) *SendMsgByTaskRequestSendConfig {
	s.SendType = &v
	return s
}

func (s *SendMsgByTaskRequestSendConfig) SetUrlTrackConfig(v []*SendMsgByTaskRequestSendConfigUrlTrackConfig) *SendMsgByTaskRequestSendConfig {
	s.UrlTrackConfig = v
	return s
}

type SendMsgByTaskRequestSendConfigUrlTrackConfig struct {
	// 跟踪链接的标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 跟踪链接的坑位ID（sg开头）
	TrackId *string `json:"trackId,omitempty" xml:"trackId,omitempty"`
	// 跟踪链接URL
	TrackUrl *string `json:"trackUrl,omitempty" xml:"trackUrl,omitempty"`
}

func (s SendMsgByTaskRequestSendConfigUrlTrackConfig) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskRequestSendConfigUrlTrackConfig) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskRequestSendConfigUrlTrackConfig) SetTitle(v string) *SendMsgByTaskRequestSendConfigUrlTrackConfig {
	s.Title = &v
	return s
}

func (s *SendMsgByTaskRequestSendConfigUrlTrackConfig) SetTrackId(v string) *SendMsgByTaskRequestSendConfigUrlTrackConfig {
	s.TrackId = &v
	return s
}

func (s *SendMsgByTaskRequestSendConfigUrlTrackConfig) SetTrackUrl(v string) *SendMsgByTaskRequestSendConfigUrlTrackConfig {
	s.TrackUrl = &v
	return s
}

type SendMsgByTaskResponseBody struct {
	// Id of the request
	OpenBatchTaskId *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
}

func (s SendMsgByTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskResponseBody) SetOpenBatchTaskId(v string) *SendMsgByTaskResponseBody {
	s.OpenBatchTaskId = &v
	return s
}

type SendMsgByTaskResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMsgByTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMsgByTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskResponse) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskResponse) SetHeaders(v map[string]*string) *SendMsgByTaskResponse {
	s.Headers = v
	return s
}

func (s *SendMsgByTaskResponse) SetBody(v *SendMsgByTaskResponseBody) *SendMsgByTaskResponse {
	s.Body = v
	return s
}

type SendServiceGroupMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendServiceGroupMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendServiceGroupMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendServiceGroupMessageHeaders) SetCommonHeaders(v map[string]*string) *SendServiceGroupMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendServiceGroupMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendServiceGroupMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendServiceGroupMessageRequest struct {
	// at dingtalkId
	AtDingtalkIds []*string `json:"atDingtalkIds,omitempty" xml:"atDingtalkIds,omitempty" type:"Repeated"`
	// at 手机号
	AtMobiles []*string `json:"atMobiles,omitempty" xml:"atMobiles,omitempty" type:"Repeated"`
	// at unionIds
	AtUnionIds []*string `json:"atUnionIds,omitempty" xml:"atUnionIds,omitempty" type:"Repeated"`
	// 排列方式：0-按钮竖直排列，1-按钮横向排列
	BtnOrientation *string `json:"btnOrientation,omitempty" xml:"btnOrientation,omitempty"`
	// actionCard按钮
	Btns []*SendServiceGroupMessageRequestBtns `json:"btns,omitempty" xml:"btns,omitempty" type:"Repeated"`
	// 内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 如果正文内容包含链接，并且按钮链接和文本链接分开跳转，则传递true; 否则传递false
	HasContentLinks *bool `json:"hasContentLinks,omitempty" xml:"hasContentLinks,omitempty"`
	// 是否 at所有人
	IsAtAll *bool `json:"isAtAll,omitempty" xml:"isAtAll,omitempty"`
	// 消息类型：MARKDOWN，ACTIONCARD
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// dingtalkId接收者
	ReceiverDingtalkIds []*string `json:"receiverDingtalkIds,omitempty" xml:"receiverDingtalkIds,omitempty" type:"Repeated"`
	// 手机号接收者
	ReceiverMobiles []*string `json:"receiverMobiles,omitempty" xml:"receiverMobiles,omitempty" type:"Repeated"`
	// unionId接收者
	ReceiverUnionIds []*string `json:"receiverUnionIds,omitempty" xml:"receiverUnionIds,omitempty" type:"Repeated"`
	// 开放群ID
	TargetOpenConversationId *string `json:"targetOpenConversationId,omitempty" xml:"targetOpenConversationId,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendServiceGroupMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendServiceGroupMessageRequest) GoString() string {
	return s.String()
}

func (s *SendServiceGroupMessageRequest) SetAtDingtalkIds(v []*string) *SendServiceGroupMessageRequest {
	s.AtDingtalkIds = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetAtMobiles(v []*string) *SendServiceGroupMessageRequest {
	s.AtMobiles = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetAtUnionIds(v []*string) *SendServiceGroupMessageRequest {
	s.AtUnionIds = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetBtnOrientation(v string) *SendServiceGroupMessageRequest {
	s.BtnOrientation = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetBtns(v []*SendServiceGroupMessageRequestBtns) *SendServiceGroupMessageRequest {
	s.Btns = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetContent(v string) *SendServiceGroupMessageRequest {
	s.Content = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetHasContentLinks(v bool) *SendServiceGroupMessageRequest {
	s.HasContentLinks = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetIsAtAll(v bool) *SendServiceGroupMessageRequest {
	s.IsAtAll = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetMessageType(v string) *SendServiceGroupMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetReceiverDingtalkIds(v []*string) *SendServiceGroupMessageRequest {
	s.ReceiverDingtalkIds = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetReceiverMobiles(v []*string) *SendServiceGroupMessageRequest {
	s.ReceiverMobiles = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetReceiverUnionIds(v []*string) *SendServiceGroupMessageRequest {
	s.ReceiverUnionIds = v
	return s
}

func (s *SendServiceGroupMessageRequest) SetTargetOpenConversationId(v string) *SendServiceGroupMessageRequest {
	s.TargetOpenConversationId = &v
	return s
}

func (s *SendServiceGroupMessageRequest) SetTitle(v string) *SendServiceGroupMessageRequest {
	s.Title = &v
	return s
}

type SendServiceGroupMessageRequestBtns struct {
	// 跳转地址
	ActionURL *string `json:"actionURL,omitempty" xml:"actionURL,omitempty"`
	// 按钮名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendServiceGroupMessageRequestBtns) String() string {
	return tea.Prettify(s)
}

func (s SendServiceGroupMessageRequestBtns) GoString() string {
	return s.String()
}

func (s *SendServiceGroupMessageRequestBtns) SetActionURL(v string) *SendServiceGroupMessageRequestBtns {
	s.ActionURL = &v
	return s
}

func (s *SendServiceGroupMessageRequestBtns) SetTitle(v string) *SendServiceGroupMessageRequestBtns {
	s.Title = &v
	return s
}

type SendServiceGroupMessageResponseBody struct {
	// 开放消息任务ID
	OpenMsgTaskId *string `json:"openMsgTaskId,omitempty" xml:"openMsgTaskId,omitempty"`
}

func (s SendServiceGroupMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendServiceGroupMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendServiceGroupMessageResponseBody) SetOpenMsgTaskId(v string) *SendServiceGroupMessageResponseBody {
	s.OpenMsgTaskId = &v
	return s
}

type SendServiceGroupMessageResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendServiceGroupMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendServiceGroupMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendServiceGroupMessageResponse) GoString() string {
	return s.String()
}

func (s *SendServiceGroupMessageResponse) SetHeaders(v map[string]*string) *SendServiceGroupMessageResponse {
	s.Headers = v
	return s
}

func (s *SendServiceGroupMessageResponse) SetBody(v *SendServiceGroupMessageResponseBody) *SendServiceGroupMessageResponse {
	s.Body = v
	return s
}

type SetRobotConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetRobotConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetRobotConfigHeaders) GoString() string {
	return s.String()
}

func (s *SetRobotConfigHeaders) SetCommonHeaders(v map[string]*string) *SetRobotConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetRobotConfigHeaders) SetXAcsDingtalkAccessToken(v string) *SetRobotConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetRobotConfigRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 群组开放ID
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 设置状态，0代表关闭,1代表开启
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SetRobotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRobotConfigRequest) GoString() string {
	return s.String()
}

func (s *SetRobotConfigRequest) SetDingIsvOrgId(v int64) *SetRobotConfigRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *SetRobotConfigRequest) SetDingOrgId(v int64) *SetRobotConfigRequest {
	s.DingOrgId = &v
	return s
}

func (s *SetRobotConfigRequest) SetDingSuiteKey(v string) *SetRobotConfigRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *SetRobotConfigRequest) SetDingTokenGrantType(v int64) *SetRobotConfigRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *SetRobotConfigRequest) SetOpenGroupSetId(v string) *SetRobotConfigRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *SetRobotConfigRequest) SetOpenTeamId(v string) *SetRobotConfigRequest {
	s.OpenTeamId = &v
	return s
}

func (s *SetRobotConfigRequest) SetStatus(v string) *SetRobotConfigRequest {
	s.Status = &v
	return s
}

type SetRobotConfigResponseBody struct {
	Result *SetRobotConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SetRobotConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRobotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetRobotConfigResponseBody) SetResult(v *SetRobotConfigResponseBodyResult) *SetRobotConfigResponseBody {
	s.Result = v
	return s
}

type SetRobotConfigResponseBodyResult struct {
	// 业务Key
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// 业务value
	ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
}

func (s SetRobotConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SetRobotConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SetRobotConfigResponseBodyResult) SetConfigKey(v string) *SetRobotConfigResponseBodyResult {
	s.ConfigKey = &v
	return s
}

func (s *SetRobotConfigResponseBodyResult) SetConfigValue(v string) *SetRobotConfigResponseBodyResult {
	s.ConfigValue = &v
	return s
}

type SetRobotConfigResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetRobotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRobotConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRobotConfigResponse) GoString() string {
	return s.String()
}

func (s *SetRobotConfigResponse) SetHeaders(v map[string]*string) *SetRobotConfigResponse {
	s.Headers = v
	return s
}

func (s *SetRobotConfigResponse) SetBody(v *SetRobotConfigResponseBody) *SetRobotConfigResponse {
	s.Body = v
	return s
}

type TakeTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TakeTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s TakeTicketHeaders) GoString() string {
	return s.String()
}

func (s *TakeTicketHeaders) SetCommonHeaders(v map[string]*string) *TakeTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TakeTicketHeaders) SetXAcsDingtalkAccessToken(v string) *TakeTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TakeTicketRequest struct {
	OpenTeamId   *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	TakerUnionId *string `json:"takerUnionId,omitempty" xml:"takerUnionId,omitempty"`
}

func (s TakeTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s TakeTicketRequest) GoString() string {
	return s.String()
}

func (s *TakeTicketRequest) SetOpenTeamId(v string) *TakeTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *TakeTicketRequest) SetOpenTicketId(v string) *TakeTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *TakeTicketRequest) SetTakerUnionId(v string) *TakeTicketRequest {
	s.TakerUnionId = &v
	return s
}

type TakeTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TakeTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s TakeTicketResponse) GoString() string {
	return s.String()
}

func (s *TakeTicketResponse) SetHeaders(v map[string]*string) *TakeTicketResponse {
	s.Headers = v
	return s
}

type TransferTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TransferTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketHeaders) GoString() string {
	return s.String()
}

func (s *TransferTicketHeaders) SetCommonHeaders(v map[string]*string) *TransferTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TransferTicketHeaders) SetXAcsDingtalkAccessToken(v string) *TransferTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TransferTicketRequest struct {
	Notify *TransferTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单开放ID
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 工单处理人
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// 被转单人UnionId列表
	ProcessorUnionIds []*string `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	// 工单备注
	TicketMemo *TransferTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s TransferTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketRequest) GoString() string {
	return s.String()
}

func (s *TransferTicketRequest) SetNotify(v *TransferTicketRequestNotify) *TransferTicketRequest {
	s.Notify = v
	return s
}

func (s *TransferTicketRequest) SetOpenTeamId(v string) *TransferTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *TransferTicketRequest) SetOpenTicketId(v string) *TransferTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *TransferTicketRequest) SetProcessorUnionId(v string) *TransferTicketRequest {
	s.ProcessorUnionId = &v
	return s
}

func (s *TransferTicketRequest) SetProcessorUnionIds(v []*string) *TransferTicketRequest {
	s.ProcessorUnionIds = v
	return s
}

func (s *TransferTicketRequest) SetTicketMemo(v *TransferTicketRequestTicketMemo) *TransferTicketRequest {
	s.TicketMemo = v
	return s
}

type TransferTicketRequestNotify struct {
	// 群中通知接收人（钉钉UnionId）
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// 是否向群内推送一个全员可见工单通知卡片
	NoticeAllGroupMember *bool `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
	// 企业工作通知接收人（钉钉UnionId）
	WorkNoticeReceiverUnionIds []*string `json:"workNoticeReceiverUnionIds,omitempty" xml:"workNoticeReceiverUnionIds,omitempty" type:"Repeated"`
}

func (s TransferTicketRequestNotify) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *TransferTicketRequestNotify) SetGroupNoticeReceiverUnionIds(v []*string) *TransferTicketRequestNotify {
	s.GroupNoticeReceiverUnionIds = v
	return s
}

func (s *TransferTicketRequestNotify) SetNoticeAllGroupMember(v bool) *TransferTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *TransferTicketRequestNotify) SetWorkNoticeReceiverUnionIds(v []*string) *TransferTicketRequestNotify {
	s.WorkNoticeReceiverUnionIds = v
	return s
}

type TransferTicketRequestTicketMemo struct {
	// 备注相关的附件
	Attachments []*TransferTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// 文字备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s TransferTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *TransferTicketRequestTicketMemo) SetAttachments(v []*TransferTicketRequestTicketMemoAttachments) *TransferTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *TransferTicketRequestTicketMemo) SetMemo(v string) *TransferTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

type TransferTicketRequestTicketMemoAttachments struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	Key      *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s TransferTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *TransferTicketRequestTicketMemoAttachments) SetFileName(v string) *TransferTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *TransferTicketRequestTicketMemoAttachments) SetKey(v string) *TransferTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type TransferTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TransferTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferTicketResponse) GoString() string {
	return s.String()
}

func (s *TransferTicketResponse) SetHeaders(v map[string]*string) *TransferTicketResponse {
	s.Headers = v
	return s
}

type UpdateGroupTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupTagHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupTagHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupTagHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupTagRequest struct {
	// 群会话ID集合
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	TagNames            []*string `json:"tagNames,omitempty" xml:"tagNames,omitempty" type:"Repeated"`
	// 更新类型，APPEND、NOTAPPEND、DELETE三种类型
	UpdateType *string `json:"updateType,omitempty" xml:"updateType,omitempty"`
}

func (s UpdateGroupTagRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupTagRequest) SetOpenConversationIds(v []*string) *UpdateGroupTagRequest {
	s.OpenConversationIds = v
	return s
}

func (s *UpdateGroupTagRequest) SetTagNames(v []*string) *UpdateGroupTagRequest {
	s.TagNames = v
	return s
}

func (s *UpdateGroupTagRequest) SetUpdateType(v string) *UpdateGroupTagRequest {
	s.UpdateType = &v
	return s
}

type UpdateGroupTagResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateGroupTagResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupTagResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupTagResponse) SetHeaders(v map[string]*string) *UpdateGroupTagResponse {
	s.Headers = v
	return s
}

type UpdateTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTicketHeaders) SetCommonHeaders(v map[string]*string) *UpdateTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTicketHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTicketRequest struct {
	// 自定义字段值JSON格式
	CustomFields *string `json:"customFields,omitempty" xml:"customFields,omitempty"`
	// 团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单开放id
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 工单处理人unionId
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// 备注
	TicketMemo *UpdateTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s UpdateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequest) SetCustomFields(v string) *UpdateTicketRequest {
	s.CustomFields = &v
	return s
}

func (s *UpdateTicketRequest) SetOpenTeamId(v string) *UpdateTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *UpdateTicketRequest) SetOpenTicketId(v string) *UpdateTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *UpdateTicketRequest) SetProcessorUnionId(v string) *UpdateTicketRequest {
	s.ProcessorUnionId = &v
	return s
}

func (s *UpdateTicketRequest) SetTicketMemo(v *UpdateTicketRequestTicketMemo) *UpdateTicketRequest {
	s.TicketMemo = v
	return s
}

type UpdateTicketRequestTicketMemo struct {
	// 备注相关的附件
	Attachments []*UpdateTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// 备注文字
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s UpdateTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequestTicketMemo) SetAttachments(v []*UpdateTicketRequestTicketMemoAttachments) *UpdateTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *UpdateTicketRequestTicketMemo) SetMemo(v string) *UpdateTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

type UpdateTicketRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UpdateTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequestTicketMemoAttachments) SetFileName(v string) *UpdateTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *UpdateTicketRequestTicketMemoAttachments) SetKey(v string) *UpdateTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type UpdateTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketResponse) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponse) SetHeaders(v map[string]*string) *UpdateTicketResponse {
	s.Headers = v
	return s
}

type UpgradeCloudGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpgradeCloudGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpgradeCloudGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpgradeCloudGroupHeaders) SetCommonHeaders(v map[string]*string) *UpgradeCloudGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpgradeCloudGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpgradeCloudGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpgradeCloudGroupRequest struct {
	// 云客服租户id
	CcsInstanceId *string `json:"ccsInstanceId,omitempty" xml:"ccsInstanceId,omitempty"`
	// 钉钉群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 升级的目标群组id
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 升级的目标团队id
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 升级的目标模板id
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s UpgradeCloudGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeCloudGroupRequest) GoString() string {
	return s.String()
}

func (s *UpgradeCloudGroupRequest) SetCcsInstanceId(v string) *UpgradeCloudGroupRequest {
	s.CcsInstanceId = &v
	return s
}

func (s *UpgradeCloudGroupRequest) SetOpenConversationId(v string) *UpgradeCloudGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpgradeCloudGroupRequest) SetOpenGroupSetId(v string) *UpgradeCloudGroupRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *UpgradeCloudGroupRequest) SetOpenTeamId(v string) *UpgradeCloudGroupRequest {
	s.OpenTeamId = &v
	return s
}

func (s *UpgradeCloudGroupRequest) SetTemplateId(v string) *UpgradeCloudGroupRequest {
	s.TemplateId = &v
	return s
}

type UpgradeCloudGroupResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpgradeCloudGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeCloudGroupResponse) GoString() string {
	return s.String()
}

func (s *UpgradeCloudGroupResponse) SetHeaders(v map[string]*string) *UpgradeCloudGroupResponse {
	s.Headers = v
	return s
}

type UpgradeNormalGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpgradeNormalGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpgradeNormalGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpgradeNormalGroupHeaders) SetCommonHeaders(v map[string]*string) *UpgradeNormalGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpgradeNormalGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpgradeNormalGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpgradeNormalGroupRequest struct {
	// 群id
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 升级的目标群组id
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// 升级的目标团队id
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 升级的目标模板id
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s UpgradeNormalGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeNormalGroupRequest) GoString() string {
	return s.String()
}

func (s *UpgradeNormalGroupRequest) SetOpenConversationId(v string) *UpgradeNormalGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpgradeNormalGroupRequest) SetOpenGroupSetId(v string) *UpgradeNormalGroupRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *UpgradeNormalGroupRequest) SetOpenTeamId(v string) *UpgradeNormalGroupRequest {
	s.OpenTeamId = &v
	return s
}

func (s *UpgradeNormalGroupRequest) SetTemplateId(v string) *UpgradeNormalGroupRequest {
	s.TemplateId = &v
	return s
}

type UpgradeNormalGroupResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpgradeNormalGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeNormalGroupResponse) GoString() string {
	return s.String()
}

func (s *UpgradeNormalGroupResponse) SetHeaders(v map[string]*string) *UpgradeNormalGroupResponse {
	s.Headers = v
	return s
}

type UrgeTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UrgeTicketHeaders) String() string {
	return tea.Prettify(s)
}

func (s UrgeTicketHeaders) GoString() string {
	return s.String()
}

func (s *UrgeTicketHeaders) SetCommonHeaders(v map[string]*string) *UrgeTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UrgeTicketHeaders) SetXAcsDingtalkAccessToken(v string) *UrgeTicketHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UrgeTicketRequest struct {
	// 开放团队ID
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// 工单开放id
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// 工单催单操作人UnionId
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	// 备注
	TicketMemo *UrgeTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
}

func (s UrgeTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s UrgeTicketRequest) GoString() string {
	return s.String()
}

func (s *UrgeTicketRequest) SetOpenTeamId(v string) *UrgeTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *UrgeTicketRequest) SetOpenTicketId(v string) *UrgeTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *UrgeTicketRequest) SetOperatorUnionId(v string) *UrgeTicketRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *UrgeTicketRequest) SetTicketMemo(v *UrgeTicketRequestTicketMemo) *UrgeTicketRequest {
	s.TicketMemo = v
	return s
}

type UrgeTicketRequestTicketMemo struct {
	// 备注相关的附件
	Attachments []*UrgeTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// 备注文字
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s UrgeTicketRequestTicketMemo) String() string {
	return tea.Prettify(s)
}

func (s UrgeTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *UrgeTicketRequestTicketMemo) SetAttachments(v []*UrgeTicketRequestTicketMemoAttachments) *UrgeTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *UrgeTicketRequestTicketMemo) SetMemo(v string) *UrgeTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

type UrgeTicketRequestTicketMemoAttachments struct {
	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UrgeTicketRequestTicketMemoAttachments) String() string {
	return tea.Prettify(s)
}

func (s UrgeTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *UrgeTicketRequestTicketMemoAttachments) SetFileName(v string) *UrgeTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *UrgeTicketRequestTicketMemoAttachments) SetKey(v string) *UrgeTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

type UrgeTicketResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UrgeTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s UrgeTicketResponse) GoString() string {
	return s.String()
}

func (s *UrgeTicketResponse) SetHeaders(v map[string]*string) *UrgeTicketResponse {
	s.Headers = v
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

func (client *Client) AddKnowledge(request *AddKnowledgeRequest) (_result *AddKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddKnowledgeHeaders{}
	_result = &AddKnowledgeResponse{}
	_body, _err := client.AddKnowledgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddKnowledgeWithOptions(request *AddKnowledgeRequest, headers *AddKnowledgeHeaders, runtime *util.RuntimeOptions) (_result *AddKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachmentList)) {
		body["attachmentList"] = request.AttachmentList
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ExtTitle)) {
		body["extTitle"] = request.ExtTitle
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryKey)) {
		body["libraryKey"] = request.LibraryKey
	}

	if !tea.BoolValue(util.IsUnset(request.LinkUrl)) {
		body["linkUrl"] = request.LinkUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePrimaryKey)) {
		body["sourcePrimaryKey"] = request.SourcePrimaryKey
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
	_result = &AddKnowledgeResponse{}
	_body, _err := client.DoROARequest(tea.String("AddKnowledge"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/knowledges"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLibrary(request *AddLibraryRequest) (_result *AddLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddLibraryHeaders{}
	_result = &AddLibraryResponse{}
	_body, _err := client.AddLibraryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLibraryWithOptions(request *AddLibraryRequest, headers *AddLibraryHeaders, runtime *util.RuntimeOptions) (_result *AddLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamIds)) {
		body["openTeamIds"] = request.OpenTeamIds
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePrimaryKey)) {
		body["sourcePrimaryKey"] = request.SourcePrimaryKey
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &AddLibraryResponse{}
	_body, _err := client.DoROARequest(tea.String("AddLibrary"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/librarys"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddOpenCategory(request *AddOpenCategoryRequest) (_result *AddOpenCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddOpenCategoryHeaders{}
	_result = &AddOpenCategoryResponse{}
	_body, _err := client.AddOpenCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddOpenCategoryWithOptions(request *AddOpenCategoryRequest, headers *AddOpenCategoryHeaders, runtime *util.RuntimeOptions) (_result *AddOpenCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["parentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["userName"] = request.UserName
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
	_result = &AddOpenCategoryResponse{}
	_body, _err := client.DoROARequest(tea.String("AddOpenCategory"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/openCategories"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddOpenKnowledge(request *AddOpenKnowledgeRequest) (_result *AddOpenKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddOpenKnowledgeHeaders{}
	_result = &AddOpenKnowledgeResponse{}
	_body, _err := client.AddOpenKnowledgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddOpenKnowledgeWithOptions(request *AddOpenKnowledgeRequest, headers *AddOpenKnowledgeHeaders, runtime *util.RuntimeOptions) (_result *AddOpenKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attachments)) {
		body["attachments"] = request.Attachments
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["categoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.EffectTimeend)) {
		body["effectTimeend"] = request.EffectTimeend
	}

	if !tea.BoolValue(util.IsUnset(request.EffectTimestart)) {
		body["effectTimestart"] = request.EffectTimestart
	}

	if !tea.BoolValue(util.IsUnset(request.ExtTitle)) {
		body["extTitle"] = request.ExtTitle
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["userName"] = request.UserName
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
	_result = &AddOpenKnowledgeResponse{}
	_body, _err := client.DoROARequest(tea.String("AddOpenKnowledge"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/openKnowledges"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddOpenLibrary(request *AddOpenLibraryRequest) (_result *AddOpenLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddOpenLibraryHeaders{}
	_result = &AddOpenLibraryResponse{}
	_body, _err := client.AddOpenLibraryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddOpenLibraryWithOptions(request *AddOpenLibraryRequest, headers *AddOpenLibraryHeaders, runtime *util.RuntimeOptions) (_result *AddOpenLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["userName"] = request.UserName
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
	_result = &AddOpenLibraryResponse{}
	_body, _err := client.DoROARequest(tea.String("AddOpenLibrary"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/openLibraries"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTicketMemo(request *AddTicketMemoRequest) (_result *AddTicketMemoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddTicketMemoHeaders{}
	_result = &AddTicketMemoResponse{}
	_body, _err := client.AddTicketMemoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTicketMemoWithOptions(request *AddTicketMemoRequest, headers *AddTicketMemoHeaders, runtime *util.RuntimeOptions) (_result *AddTicketMemoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionId)) {
		body["processorUnionId"] = request.ProcessorUnionId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
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
	_result = &AddTicketMemoResponse{}
	_body, _err := client.DoROARequest(tea.String("AddTicketMemo"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/memos"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssignTicket(request *AssignTicketRequest) (_result *AssignTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AssignTicketHeaders{}
	_result = &AssignTicketResponse{}
	_body, _err := client.AssignTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssignTicketWithOptions(request *AssignTicketRequest, headers *AssignTicketHeaders, runtime *util.RuntimeOptions) (_result *AssignTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionIds)) {
		body["processorUnionIds"] = request.ProcessorUnionIds
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
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
	_result = &AssignTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("AssignTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/assign"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGetGroupSetConfig(request *BatchGetGroupSetConfigRequest) (_result *BatchGetGroupSetConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchGetGroupSetConfigHeaders{}
	_result = &BatchGetGroupSetConfigResponse{}
	_body, _err := client.BatchGetGroupSetConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetGroupSetConfigWithOptions(request *BatchGetGroupSetConfigRequest, headers *BatchGetGroupSetConfigHeaders, runtime *util.RuntimeOptions) (_result *BatchGetGroupSetConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigKeys)) {
		body["configKeys"] = request.ConfigKeys
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &BatchGetGroupSetConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchGetGroupSetConfig"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groupSetConfigs/batchQuery"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchQuerySendMessageTask(request *BatchQuerySendMessageTaskRequest) (_result *BatchQuerySendMessageTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQuerySendMessageTaskHeaders{}
	_result = &BatchQuerySendMessageTaskResponse{}
	_body, _err := client.BatchQuerySendMessageTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchQuerySendMessageTaskWithOptions(request *BatchQuerySendMessageTaskRequest, headers *BatchQuerySendMessageTaskHeaders, runtime *util.RuntimeOptions) (_result *BatchQuerySendMessageTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GetReadCount)) {
		body["getReadCount"] = request.GetReadCount
	}

	if !tea.BoolValue(util.IsUnset(request.GmtCreateEnd)) {
		body["gmtCreateEnd"] = request.GmtCreateEnd
	}

	if !tea.BoolValue(util.IsUnset(request.GmtCreateStart)) {
		body["gmtCreateStart"] = request.GmtCreateStart
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["taskName"] = request.TaskName
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
	_result = &BatchQuerySendMessageTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchQuerySendMessageTask"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tasks/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BoundTemplateToTeam(request *BoundTemplateToTeamRequest) (_result *BoundTemplateToTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BoundTemplateToTeamHeaders{}
	_result = &BoundTemplateToTeamResponse{}
	_body, _err := client.BoundTemplateToTeamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BoundTemplateToTeamWithOptions(request *BoundTemplateToTeamRequest, headers *BoundTemplateToTeamHeaders, runtime *util.RuntimeOptions) (_result *BoundTemplateToTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.RobotConfig)) {
		body["robotConfig"] = request.RobotConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateDesc)) {
		body["templateDesc"] = request.TemplateDesc
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		body["templateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["templateType"] = request.TemplateType
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
	_result = &BoundTemplateToTeamResponse{}
	_body, _err := client.DoROARequest(tea.String("BoundTemplateToTeam"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/teams/templates/bound"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelTicket(request *CancelTicketRequest) (_result *CancelTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelTicketHeaders{}
	_result = &CancelTicketResponse{}
	_body, _err := client.CancelTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelTicketWithOptions(request *CancelTicketRequest, headers *CancelTicketHeaders, runtime *util.RuntimeOptions) (_result *CancelTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
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
	_result = &CancelTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("CancelTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/cancel"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseHumanSession(request *CloseHumanSessionRequest) (_result *CloseHumanSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CloseHumanSessionHeaders{}
	_result = &CloseHumanSessionResponse{}
	_body, _err := client.CloseHumanSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseHumanSessionWithOptions(request *CloseHumanSessionRequest, headers *CloseHumanSessionHeaders, runtime *util.RuntimeOptions) (_result *CloseHumanSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &CloseHumanSessionResponse{}
	_body, _err := client.DoROARequest(tea.String("CloseHumanSession"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/humanSessions/close"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupHeaders{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, headers *CreateGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupBizId)) {
		body["groupBizId"] = request.GroupBizId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTagNames)) {
		body["groupTagNames"] = request.GroupTagNames
	}

	if !tea.BoolValue(util.IsUnset(request.MemberStaffIds)) {
		body["memberStaffIds"] = request.MemberStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerStaffId)) {
		body["ownerStaffId"] = request.OwnerStaffId
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
	_result = &CreateGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateGroup"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupSet(request *CreateGroupSetRequest) (_result *CreateGroupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupSetHeaders{}
	_result = &CreateGroupSetResponse{}
	_body, _err := client.CreateGroupSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupSetWithOptions(request *CreateGroupSetRequest, headers *CreateGroupSetHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupSetName)) {
		body["groupSetName"] = request.GroupSetName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &CreateGroupSetResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateGroupSet"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groupSets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTeam(request *CreateTeamRequest) (_result *CreateTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTeamHeaders{}
	_result = &CreateTeamResponse{}
	_body, _err := client.CreateTeamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTeamWithOptions(request *CreateTeamRequest, headers *CreateTeamHeaders, runtime *util.RuntimeOptions) (_result *CreateTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorDingUnionId)) {
		body["creatorDingUnionId"] = request.CreatorDingUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.TeamName)) {
		body["teamName"] = request.TeamName
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
	_result = &CreateTeamResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTeam"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/teams"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTicketHeaders{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, headers *CreateTicketHeaders, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		body["customFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTemplateBizId)) {
		body["openTemplateBizId"] = request.OpenTemplateBizId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionIds)) {
		body["processorUnionIds"] = request.ProcessorUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SceneContext))) {
		body["sceneContext"] = request.SceneContext
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
	_result = &CreateTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKnowledge(request *DeleteKnowledgeRequest) (_result *DeleteKnowledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteKnowledgeHeaders{}
	_result = &DeleteKnowledgeResponse{}
	_body, _err := client.DeleteKnowledgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKnowledgeWithOptions(request *DeleteKnowledgeRequest, headers *DeleteKnowledgeHeaders, runtime *util.RuntimeOptions) (_result *DeleteKnowledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibraryKey)) {
		body["libraryKey"] = request.LibraryKey
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePrimaryKey)) {
		body["sourcePrimaryKey"] = request.SourcePrimaryKey
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
	_result = &DeleteKnowledgeResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteKnowledge"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/knowledges/batchDelete"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishTicket(request *FinishTicketRequest) (_result *FinishTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FinishTicketHeaders{}
	_result = &FinishTicketResponse{}
	_body, _err := client.FinishTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishTicketWithOptions(request *FinishTicketRequest, headers *FinishTicketHeaders, runtime *util.RuntimeOptions) (_result *FinishTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionId)) {
		body["processorUnionId"] = request.ProcessorUnionId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
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
	_result = &FinishTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("FinishTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/finish"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOssTempUrl(request *GetOssTempUrlRequest) (_result *GetOssTempUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOssTempUrlHeaders{}
	_result = &GetOssTempUrlResponse{}
	_body, _err := client.GetOssTempUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOssTempUrlWithOptions(request *GetOssTempUrlRequest, headers *GetOssTempUrlHeaders, runtime *util.RuntimeOptions) (_result *GetOssTempUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FetchMode)) {
		query["fetchMode"] = request.FetchMode
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		query["openTeamId"] = request.OpenTeamId
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
	_result = &GetOssTempUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOssTempUrl"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/serviceGroup/ossServices/tempUrls"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStoragePolicy(request *GetStoragePolicyRequest) (_result *GetStoragePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetStoragePolicyHeaders{}
	_result = &GetStoragePolicyResponse{}
	_body, _err := client.GetStoragePolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStoragePolicyWithOptions(request *GetStoragePolicyRequest, headers *GetStoragePolicyHeaders, runtime *util.RuntimeOptions) (_result *GetStoragePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		body["fileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &GetStoragePolicyResponse{}
	_body, _err := client.DoROARequest(tea.String("GetStoragePolicy"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/ossServices/policies"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTicket(request *GetTicketRequest) (_result *GetTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTicketHeaders{}
	_result = &GetTicketResponse{}
	_body, _err := client.GetTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTicketWithOptions(request *GetTicketRequest, headers *GetTicketHeaders, runtime *util.RuntimeOptions) (_result *GetTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		query["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		query["openTicketId"] = request.OpenTicketId
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
	_result = &GetTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTicketOperateRecord(request *ListTicketOperateRecordRequest) (_result *ListTicketOperateRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTicketOperateRecordHeaders{}
	_result = &ListTicketOperateRecordResponse{}
	_body, _err := client.ListTicketOperateRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTicketOperateRecordWithOptions(request *ListTicketOperateRecordRequest, headers *ListTicketOperateRecordHeaders, runtime *util.RuntimeOptions) (_result *ListTicketOperateRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		query["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		query["openTicketId"] = request.OpenTicketId
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
	_result = &ListTicketOperateRecordResponse{}
	_body, _err := client.DoROARequest(tea.String("ListTicketOperateRecord"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/operateRecords"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserTeams(userId *string) (_result *ListUserTeamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListUserTeamsHeaders{}
	_result = &ListUserTeamsResponse{}
	_body, _err := client.ListUserTeamsWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserTeamsWithOptions(userId *string, headers *ListUserTeamsHeaders, runtime *util.RuntimeOptions) (_result *ListUserTeamsResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &ListUserTeamsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListUserTeams"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/serviceGroup/users/"+tea.StringValue(userId)+"/teams"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryActiveUsers(request *QueryActiveUsersRequest) (_result *QueryActiveUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryActiveUsersHeaders{}
	_result = &QueryActiveUsersResponse{}
	_body, _err := client.QueryActiveUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryActiveUsersWithOptions(request *QueryActiveUsersRequest, headers *QueryActiveUsersHeaders, runtime *util.RuntimeOptions) (_result *QueryActiveUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		query["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.TopN)) {
		query["topN"] = request.TopN
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
	_result = &QueryActiveUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryActiveUsers"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groups/queryActiveUsers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroup(request *QueryGroupRequest) (_result *QueryGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupHeaders{}
	_result = &QueryGroupResponse{}
	_body, _err := client.QueryGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupWithOptions(request *QueryGroupRequest, headers *QueryGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &QueryGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroup"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groups/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupSet(request *QueryGroupSetRequest) (_result *QueryGroupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupSetHeaders{}
	_result = &QueryGroupSetResponse{}
	_body, _err := client.QueryGroupSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupSetWithOptions(request *QueryGroupSetRequest, headers *QueryGroupSetHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		query["openTeamId"] = request.OpenTeamId
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
	_result = &QueryGroupSetResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupSet"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groupSets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySendMsgTaskStatistics(request *QuerySendMsgTaskStatisticsRequest) (_result *QuerySendMsgTaskStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySendMsgTaskStatisticsHeaders{}
	_result = &QuerySendMsgTaskStatisticsResponse{}
	_body, _err := client.QuerySendMsgTaskStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySendMsgTaskStatisticsWithOptions(request *QuerySendMsgTaskStatisticsRequest, headers *QuerySendMsgTaskStatisticsHeaders, runtime *util.RuntimeOptions) (_result *QuerySendMsgTaskStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenBatchTaskId)) {
		body["openBatchTaskId"] = request.OpenBatchTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &QuerySendMsgTaskStatisticsResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySendMsgTaskStatistics"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tasks/statistics/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySendMsgTaskStatisticsDetail(request *QuerySendMsgTaskStatisticsDetailRequest) (_result *QuerySendMsgTaskStatisticsDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySendMsgTaskStatisticsDetailHeaders{}
	_result = &QuerySendMsgTaskStatisticsDetailResponse{}
	_body, _err := client.QuerySendMsgTaskStatisticsDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySendMsgTaskStatisticsDetailWithOptions(request *QuerySendMsgTaskStatisticsDetailRequest, headers *QuerySendMsgTaskStatisticsDetailHeaders, runtime *util.RuntimeOptions) (_result *QuerySendMsgTaskStatisticsDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenBatchTaskId)) {
		body["openBatchTaskId"] = request.OpenBatchTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &QuerySendMsgTaskStatisticsDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySendMsgTaskStatisticsDetail"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tasks/statistics/details/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServiceGroupMessageReadStatus(request *QueryServiceGroupMessageReadStatusRequest) (_result *QueryServiceGroupMessageReadStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryServiceGroupMessageReadStatusHeaders{}
	_result = &QueryServiceGroupMessageReadStatusResponse{}
	_body, _err := client.QueryServiceGroupMessageReadStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServiceGroupMessageReadStatusWithOptions(request *QueryServiceGroupMessageReadStatusRequest, headers *QueryServiceGroupMessageReadStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryServiceGroupMessageReadStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenMsgTaskId)) {
		body["openMsgTaskId"] = request.OpenMsgTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
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
	_result = &QueryServiceGroupMessageReadStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryServiceGroupMessageReadStatus"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/messages/readStatus/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResubmitTicket(request *ResubmitTicketRequest) (_result *ResubmitTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ResubmitTicketHeaders{}
	_result = &ResubmitTicketResponse{}
	_body, _err := client.ResubmitTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResubmitTicketWithOptions(request *ResubmitTicketRequest, headers *ResubmitTicketHeaders, runtime *util.RuntimeOptions) (_result *ResubmitTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		body["customFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTemplateBizId)) {
		body["openTemplateBizId"] = request.OpenTemplateBizId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionIds)) {
		body["processorUnionIds"] = request.ProcessorUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SceneContext))) {
		body["sceneContext"] = request.SceneContext
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
	_result = &ResubmitTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("ResubmitTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/resubmit"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetractTicket(request *RetractTicketRequest) (_result *RetractTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RetractTicketHeaders{}
	_result = &RetractTicketResponse{}
	_body, _err := client.RetractTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetractTicketWithOptions(request *RetractTicketRequest, headers *RetractTicketHeaders, runtime *util.RuntimeOptions) (_result *RetractTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
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
	_result = &RetractTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("RetractTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/retract"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchGroup(request *SearchGroupRequest) (_result *SearchGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchGroupHeaders{}
	_result = &SearchGroupResponse{}
	_body, _err := client.SearchGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchGroupWithOptions(request *SearchGroupRequest, headers *SearchGroupHeaders, runtime *util.RuntimeOptions) (_result *SearchGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchType)) {
		body["searchType"] = request.SearchType
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
	_result = &SearchGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchGroup"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groups/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMsgByTask(request *SendMsgByTaskRequest) (_result *SendMsgByTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendMsgByTaskHeaders{}
	_result = &SendMsgByTaskResponse{}
	_body, _err := client.SendMsgByTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMsgByTaskWithOptions(request *SendMsgByTaskRequest, headers *SendMsgByTaskHeaders, runtime *util.RuntimeOptions) (_result *SendMsgByTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.MessageContent))) {
		body["messageContent"] = request.MessageContent
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.QueryGroup))) {
		body["queryGroup"] = request.QueryGroup
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SendConfig))) {
		body["sendConfig"] = request.SendConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["taskName"] = request.TaskName
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
	_result = &SendMsgByTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("SendMsgByTask"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/messages/tasks/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendServiceGroupMessage(request *SendServiceGroupMessageRequest) (_result *SendServiceGroupMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendServiceGroupMessageHeaders{}
	_result = &SendServiceGroupMessageResponse{}
	_body, _err := client.SendServiceGroupMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendServiceGroupMessageWithOptions(request *SendServiceGroupMessageRequest, headers *SendServiceGroupMessageHeaders, runtime *util.RuntimeOptions) (_result *SendServiceGroupMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AtDingtalkIds)) {
		body["atDingtalkIds"] = request.AtDingtalkIds
	}

	if !tea.BoolValue(util.IsUnset(request.AtMobiles)) {
		body["atMobiles"] = request.AtMobiles
	}

	if !tea.BoolValue(util.IsUnset(request.AtUnionIds)) {
		body["atUnionIds"] = request.AtUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.BtnOrientation)) {
		body["btnOrientation"] = request.BtnOrientation
	}

	if !tea.BoolValue(util.IsUnset(request.Btns)) {
		body["btns"] = request.Btns
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.HasContentLinks)) {
		body["hasContentLinks"] = request.HasContentLinks
	}

	if !tea.BoolValue(util.IsUnset(request.IsAtAll)) {
		body["isAtAll"] = request.IsAtAll
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverDingtalkIds)) {
		body["receiverDingtalkIds"] = request.ReceiverDingtalkIds
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverMobiles)) {
		body["receiverMobiles"] = request.ReceiverMobiles
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUnionIds)) {
		body["receiverUnionIds"] = request.ReceiverUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOpenConversationId)) {
		body["targetOpenConversationId"] = request.TargetOpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
	_result = &SendServiceGroupMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("SendServiceGroupMessage"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/messages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRobotConfig(request *SetRobotConfigRequest) (_result *SetRobotConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetRobotConfigHeaders{}
	_result = &SetRobotConfigResponse{}
	_body, _err := client.SetRobotConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRobotConfigWithOptions(request *SetRobotConfigRequest, headers *SetRobotConfigHeaders, runtime *util.RuntimeOptions) (_result *SetRobotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
	_result = &SetRobotConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("SetRobotConfig"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/groups/configs/set"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TakeTicket(request *TakeTicketRequest) (_result *TakeTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TakeTicketHeaders{}
	_result = &TakeTicketResponse{}
	_body, _err := client.TakeTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TakeTicketWithOptions(request *TakeTicketRequest, headers *TakeTicketHeaders, runtime *util.RuntimeOptions) (_result *TakeTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.TakerUnionId)) {
		body["takerUnionId"] = request.TakerUnionId
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
	_result = &TakeTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("TakeTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/apply"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferTicket(request *TransferTicketRequest) (_result *TransferTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TransferTicketHeaders{}
	_result = &TransferTicketResponse{}
	_body, _err := client.TransferTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferTicketWithOptions(request *TransferTicketRequest, headers *TransferTicketHeaders, runtime *util.RuntimeOptions) (_result *TransferTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Notify))) {
		body["notify"] = request.Notify
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionId)) {
		body["processorUnionId"] = request.ProcessorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionIds)) {
		body["processorUnionIds"] = request.ProcessorUnionIds
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
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
	_result = &TransferTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("TransferTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/transfer"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupTag(request *UpdateGroupTagRequest) (_result *UpdateGroupTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupTagHeaders{}
	_result = &UpdateGroupTagResponse{}
	_body, _err := client.UpdateGroupTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupTagWithOptions(request *UpdateGroupTagRequest, headers *UpdateGroupTagHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationIds)) {
		body["openConversationIds"] = request.OpenConversationIds
	}

	if !tea.BoolValue(util.IsUnset(request.TagNames)) {
		body["tagNames"] = request.TagNames
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateType)) {
		body["updateType"] = request.UpdateType
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
	_result = &UpdateGroupTagResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateGroupTag"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tags"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTicket(request *UpdateTicketRequest) (_result *UpdateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTicketHeaders{}
	_result = &UpdateTicketResponse{}
	_body, _err := client.UpdateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTicketWithOptions(request *UpdateTicketRequest, headers *UpdateTicketHeaders, runtime *util.RuntimeOptions) (_result *UpdateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		body["customFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessorUnionId)) {
		body["processorUnionId"] = request.ProcessorUnionId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
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
	_result = &UpdateTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeCloudGroup(request *UpgradeCloudGroupRequest) (_result *UpgradeCloudGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpgradeCloudGroupHeaders{}
	_result = &UpgradeCloudGroupResponse{}
	_body, _err := client.UpgradeCloudGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeCloudGroupWithOptions(request *UpgradeCloudGroupRequest, headers *UpgradeCloudGroupHeaders, runtime *util.RuntimeOptions) (_result *UpgradeCloudGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CcsInstanceId)) {
		body["ccsInstanceId"] = request.CcsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
	_result = &UpgradeCloudGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("UpgradeCloudGroup"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/cloudGroups/upgrade"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeNormalGroup(request *UpgradeNormalGroupRequest) (_result *UpgradeNormalGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpgradeNormalGroupHeaders{}
	_result = &UpgradeNormalGroupResponse{}
	_body, _err := client.UpgradeNormalGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeNormalGroupWithOptions(request *UpgradeNormalGroupRequest, headers *UpgradeNormalGroupHeaders, runtime *util.RuntimeOptions) (_result *UpgradeNormalGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
	_result = &UpgradeNormalGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("UpgradeNormalGroup"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/normalGroups/upgrade"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UrgeTicket(request *UrgeTicketRequest) (_result *UrgeTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UrgeTicketHeaders{}
	_result = &UrgeTicketResponse{}
	_body, _err := client.UrgeTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UrgeTicketWithOptions(request *UrgeTicketRequest, headers *UrgeTicketHeaders, runtime *util.RuntimeOptions) (_result *UrgeTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTicketId)) {
		body["openTicketId"] = request.OpenTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TicketMemo))) {
		body["ticketMemo"] = request.TicketMemo
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
	_result = &UrgeTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("UrgeTicket"), tea.String("serviceGroup_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/serviceGroup/tickets/urge"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
