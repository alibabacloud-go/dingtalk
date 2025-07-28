// This file is auto-generated, don't edit it. Thanks.
package service_group_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddContactMemberToGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddContactMemberToGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddContactMemberToGroupHeaders) GoString() string {
	return s.String()
}

func (s *AddContactMemberToGroupHeaders) SetCommonHeaders(v map[string]*string) *AddContactMemberToGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddContactMemberToGroupHeaders) SetXAcsDingtalkAccessToken(v string) *AddContactMemberToGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddContactMemberToGroupRequest struct {
	// example:
	//
	// 不裂变：STANDARD；裂变：FISSION
	FissionType *string `json:"fissionType,omitempty" xml:"fissionType,omitempty"`
	// example:
	//
	// 888
	MemberUnionId *string `json:"memberUnionId,omitempty" xml:"memberUnionId,omitempty"`
	// example:
	//
	// 1
	MemberUserId *string `json:"memberUserId,omitempty" xml:"memberUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid***
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 888
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s AddContactMemberToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddContactMemberToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddContactMemberToGroupRequest) SetFissionType(v string) *AddContactMemberToGroupRequest {
	s.FissionType = &v
	return s
}

func (s *AddContactMemberToGroupRequest) SetMemberUnionId(v string) *AddContactMemberToGroupRequest {
	s.MemberUnionId = &v
	return s
}

func (s *AddContactMemberToGroupRequest) SetMemberUserId(v string) *AddContactMemberToGroupRequest {
	s.MemberUserId = &v
	return s
}

func (s *AddContactMemberToGroupRequest) SetOpenConversationId(v string) *AddContactMemberToGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *AddContactMemberToGroupRequest) SetOpenTeamId(v string) *AddContactMemberToGroupRequest {
	s.OpenTeamId = &v
	return s
}

type AddContactMemberToGroupResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddContactMemberToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddContactMemberToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddContactMemberToGroupResponseBody) SetResult(v bool) *AddContactMemberToGroupResponseBody {
	s.Result = &v
	return s
}

type AddContactMemberToGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddContactMemberToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddContactMemberToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddContactMemberToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddContactMemberToGroupResponse) SetHeaders(v map[string]*string) *AddContactMemberToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddContactMemberToGroupResponse) SetStatusCode(v int32) *AddContactMemberToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddContactMemberToGroupResponse) SetBody(v *AddContactMemberToGroupResponseBody) *AddContactMemberToGroupResponse {
	s.Body = v
	return s
}

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
	AttachmentList []*AddKnowledgeRequestAttachmentList `json:"attachmentList,omitempty" xml:"attachmentList,omitempty" type:"Repeated"`
	// example:
	//
	// 测试内容
	Content         *string `json:"content,omitempty" xml:"content,omitempty"`
	EffectTimeend   *int64  `json:"effectTimeend,omitempty" xml:"effectTimeend,omitempty"`
	EffectTimestart *int64  `json:"effectTimestart,omitempty" xml:"effectTimestart,omitempty"`
	ExtTitle        *string `json:"extTitle,omitempty" xml:"extTitle,omitempty"`
	Keyword         *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// xuvw1245
	LibraryKey *string `json:"libraryKey,omitempty" xml:"libraryKey,omitempty"`
	// example:
	//
	// http://www.test.com/xxxxx
	LinkUrl *string `json:"linkUrl,omitempty" xml:"linkUrl,omitempty"`
	// example:
	//
	// Jxi12wo3qxoa
	OpenTeamId  *string  `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	QuestionIds []*int64 `json:"questionIds,omitempty" xml:"questionIds,omitempty" type:"Repeated"`
	// example:
	//
	// CCM
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// CCM-123
	SourcePrimaryKey *string `json:"sourcePrimaryKey,omitempty" xml:"sourcePrimaryKey,omitempty"`
	// example:
	//
	// 测试
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// CONDITION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// V0193859102
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

func (s *AddKnowledgeRequest) SetEffectTimeend(v int64) *AddKnowledgeRequest {
	s.EffectTimeend = &v
	return s
}

func (s *AddKnowledgeRequest) SetEffectTimestart(v int64) *AddKnowledgeRequest {
	s.EffectTimestart = &v
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

func (s *AddKnowledgeRequest) SetQuestionIds(v []*int64) *AddKnowledgeRequest {
	s.QuestionIds = v
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
	// example:
	//
	// doc
	MimeType *string `json:"mime_type,omitempty" xml:"mime_type,omitempty"`
	Path     *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 655
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// pdf
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
	// example:
	//
	// 测试附件
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
	// example:
	//
	// J23suw1irs
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddKnowledgeResponse) SetStatusCode(v int32) *AddKnowledgeResponse {
	s.StatusCode = &v
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
	// example:
	//
	// 测试库描述
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	OpenTeamIds []*string `json:"openTeamIds,omitempty" xml:"openTeamIds,omitempty" type:"Repeated"`
	// example:
	//
	// CCM
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// CCM-123
	SourcePrimaryKey *string `json:"sourcePrimaryKey,omitempty" xml:"sourcePrimaryKey,omitempty"`
	// example:
	//
	// 测试库
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// EXTERNAL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// manager123
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddLibraryResponse) SetStatusCode(v int32) *AddLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLibraryResponse) SetBody(v *AddLibraryResponseBody) *AddLibraryResponse {
	s.Body = v
	return s
}

type AddMemberToServiceGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddMemberToServiceGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToServiceGroupHeaders) GoString() string {
	return s.String()
}

func (s *AddMemberToServiceGroupHeaders) SetCommonHeaders(v map[string]*string) *AddMemberToServiceGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddMemberToServiceGroupHeaders) SetXAcsDingtalkAccessToken(v string) *AddMemberToServiceGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddMemberToServiceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxxx==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jciwnfw
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddMemberToServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *AddMemberToServiceGroupRequest) SetOpenConversationId(v string) *AddMemberToServiceGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *AddMemberToServiceGroupRequest) SetOpenTeamId(v string) *AddMemberToServiceGroupRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AddMemberToServiceGroupRequest) SetUserIds(v []*string) *AddMemberToServiceGroupRequest {
	s.UserIds = v
	return s
}

type AddMemberToServiceGroupResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddMemberToServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemberToServiceGroupResponseBody) SetSuccess(v bool) *AddMemberToServiceGroupResponseBody {
	s.Success = &v
	return s
}

type AddMemberToServiceGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMemberToServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMemberToServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *AddMemberToServiceGroupResponse) SetHeaders(v map[string]*string) *AddMemberToServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *AddMemberToServiceGroupResponse) SetStatusCode(v int32) *AddMemberToServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMemberToServiceGroupResponse) SetBody(v *AddMemberToServiceGroupResponseBody) *AddMemberToServiceGroupResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 5555
	LibraryId *int64 `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jxi12wo3qxoa
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 0
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试类目
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0159003451667222
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉三多
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
	// example:
	//
	// 2
	Result *AddOpenCategoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
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
	// example:
	//
	// 111
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// title不能为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOpenCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddOpenCategoryResponse) SetStatusCode(v int32) *AddOpenCategoryResponse {
	s.StatusCode = &v
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
	// example:
	//
	// 1
	Attachments []*AddOpenKnowledgeRequestAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 44555
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 这是服务群的介绍
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2100-01-01 23:59:59
	EffectTimeend *string `json:"effectTimeend,omitempty" xml:"effectTimeend,omitempty"`
	// example:
	//
	// 1980-01-01 00:00:00
	EffectTimestart *string `json:"effectTimestart,omitempty" xml:"effectTimestart,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 这是问法1,这是问法2
	ExtTitle *string `json:"extTitle,omitempty" xml:"extTitle,omitempty"`
	// example:
	//
	// 服务群,智能场景群
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	LibraryId *int64 `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jxi12wo3qxoa
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XMD
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// 服务群,智能场景群
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 服务群是什么
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EXTERNAL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0159003451667222
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉三多
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
	// example:
	//
	// PDF
	MimeType *string `json:"mimeType,omitempty" xml:"mimeType,omitempty"`
	// example:
	//
	// https://dtapp-pub.dingtalk.com/dingtalkdesktop/test.pdf
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 444556
	Size *float64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// pdf
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
	// example:
	//
	// 这是一个附件文档
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
	// This parameter is required.
	//
	// example:
	//
	// 22
	Result *AddOpenKnowledgeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
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
	// example:
	//
	// 111
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 知识标问不能为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOpenKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddOpenKnowledgeResponse) SetStatusCode(v int32) *AddOpenKnowledgeResponse {
	s.StatusCode = &v
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
	// if can be null:
	// true
	//
	// example:
	//
	// 这个是业务知识库
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// Jxi12wo3qxoa
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XMD
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试库
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EXTERNAL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0159003451667222
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉三多
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
	// This parameter is required.
	Result *AddOpenLibraryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
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
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id      *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOpenLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddOpenLibraryResponse) SetStatusCode(v int32) *AddOpenLibraryResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a8iS4X94TgtgiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dq9hP8Sk2v6vQ6l05nCe5wiEiE
	ProcessorUnionId *string                         `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	TicketMemo       *AddTicketMemoRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
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
	Attachments []*AddTicketMemoRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 备注
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
	// example:
	//
	// ticket/image/44708069/43003/e27204b382c04832aec4243e940a1367_1625831640499.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// wahaha.txt
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *AddTicketMemoResponse) SetStatusCode(v int32) *AddTicketMemoResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hNiPO2OVktNMiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// This parameter is required.
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	// This parameter is required.
	ProcessorUnionIds []*string                      `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	TicketMemo        *AssignTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
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
	// example:
	//
	// true
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
	Attachments []*AssignTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 备注
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
	// example:
	//
	// ticket/image/44708069/43003/e27204b382c04832aec4243e940a1367_1625831640499.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// wahaha.txt
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *AssignTicketResponse) SetStatusCode(v int32) *AssignTicketResponse {
	s.StatusCode = &v
	return s
}

type BatchBindingGroupBizIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchBindingGroupBizIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchBindingGroupBizIdsHeaders) GoString() string {
	return s.String()
}

func (s *BatchBindingGroupBizIdsHeaders) SetCommonHeaders(v map[string]*string) *BatchBindingGroupBizIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchBindingGroupBizIdsHeaders) SetXAcsDingtalkAccessToken(v string) *BatchBindingGroupBizIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchBindingGroupBizIdsRequest struct {
	// This parameter is required.
	BindingGroupBizIds []*BatchBindingGroupBizIdsRequestBindingGroupBizIds `json:"bindingGroupBizIds,omitempty" xml:"bindingGroupBizIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Jciwnfw
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s BatchBindingGroupBizIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchBindingGroupBizIdsRequest) GoString() string {
	return s.String()
}

func (s *BatchBindingGroupBizIdsRequest) SetBindingGroupBizIds(v []*BatchBindingGroupBizIdsRequestBindingGroupBizIds) *BatchBindingGroupBizIdsRequest {
	s.BindingGroupBizIds = v
	return s
}

func (s *BatchBindingGroupBizIdsRequest) SetOpenTeamId(v string) *BatchBindingGroupBizIdsRequest {
	s.OpenTeamId = &v
	return s
}

type BatchBindingGroupBizIdsRequestBindingGroupBizIds struct {
	// This parameter is required.
	//
	// example:
	//
	// hghhghghhg
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid123
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s BatchBindingGroupBizIdsRequestBindingGroupBizIds) String() string {
	return tea.Prettify(s)
}

func (s BatchBindingGroupBizIdsRequestBindingGroupBizIds) GoString() string {
	return s.String()
}

func (s *BatchBindingGroupBizIdsRequestBindingGroupBizIds) SetBizId(v string) *BatchBindingGroupBizIdsRequestBindingGroupBizIds {
	s.BizId = &v
	return s
}

func (s *BatchBindingGroupBizIdsRequestBindingGroupBizIds) SetOpenConversationId(v string) *BatchBindingGroupBizIdsRequestBindingGroupBizIds {
	s.OpenConversationId = &v
	return s
}

type BatchBindingGroupBizIdsResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BatchBindingGroupBizIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchBindingGroupBizIdsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindingGroupBizIdsResponseBody) SetResult(v bool) *BatchBindingGroupBizIdsResponseBody {
	s.Result = &v
	return s
}

type BatchBindingGroupBizIdsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchBindingGroupBizIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchBindingGroupBizIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchBindingGroupBizIdsResponse) GoString() string {
	return s.String()
}

func (s *BatchBindingGroupBizIdsResponse) SetHeaders(v map[string]*string) *BatchBindingGroupBizIdsResponse {
	s.Headers = v
	return s
}

func (s *BatchBindingGroupBizIdsResponse) SetStatusCode(v int32) *BatchBindingGroupBizIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchBindingGroupBizIdsResponse) SetBody(v *BatchBindingGroupBizIdsResponseBody) *BatchBindingGroupBizIdsResponse {
	s.Body = v
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
	// This parameter is required.
	ConfigKeys []*string `json:"configKeys,omitempty" xml:"configKeys,omitempty" type:"Repeated"`
	// This parameter is required.
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// ROBOT_SWITCH
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetGroupSetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *BatchGetGroupSetConfigResponse) SetStatusCode(v int32) *BatchGetGroupSetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetGroupSetConfigResponse) SetBody(v *BatchGetGroupSetConfigResponseBody) *BatchGetGroupSetConfigResponse {
	s.Body = v
	return s
}

type BatchQueryCustomerSendTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryCustomerSendTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryCustomerSendTaskHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryCustomerSendTaskHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryCustomerSendTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryCustomerSendTaskHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryCustomerSendTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryCustomerSendTaskRequest struct {
	// example:
	//
	// 2023-06-02 00:00:00
	GmtCreateEnd *string `json:"gmtCreateEnd,omitempty" xml:"gmtCreateEnd,omitempty"`
	// example:
	//
	// 2023-06-01 00:00:00
	GmtCreateStart *string `json:"gmtCreateStart,omitempty" xml:"gmtCreateStart,omitempty"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	NeedRichStatistics *bool `json:"needRichStatistics,omitempty" xml:"needRichStatistics,omitempty"`
	// example:
	//
	// 1
	NextToken        *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenBatchTaskIds []*string `json:"openBatchTaskIds,omitempty" xml:"openBatchTaskIds,omitempty" type:"Repeated"`
	// This parameter is required.
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 哈哈哈
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s BatchQueryCustomerSendTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryCustomerSendTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryCustomerSendTaskRequest) SetGmtCreateEnd(v string) *BatchQueryCustomerSendTaskRequest {
	s.GmtCreateEnd = &v
	return s
}

func (s *BatchQueryCustomerSendTaskRequest) SetGmtCreateStart(v string) *BatchQueryCustomerSendTaskRequest {
	s.GmtCreateStart = &v
	return s
}

func (s *BatchQueryCustomerSendTaskRequest) SetMaxResults(v int64) *BatchQueryCustomerSendTaskRequest {
	s.MaxResults = &v
	return s
}

func (s *BatchQueryCustomerSendTaskRequest) SetNeedRichStatistics(v bool) *BatchQueryCustomerSendTaskRequest {
	s.NeedRichStatistics = &v
	return s
}

func (s *BatchQueryCustomerSendTaskRequest) SetNextToken(v string) *BatchQueryCustomerSendTaskRequest {
	s.NextToken = &v
	return s
}

func (s *BatchQueryCustomerSendTaskRequest) SetOpenBatchTaskIds(v []*string) *BatchQueryCustomerSendTaskRequest {
	s.OpenBatchTaskIds = v
	return s
}

func (s *BatchQueryCustomerSendTaskRequest) SetOpenTeamId(v string) *BatchQueryCustomerSendTaskRequest {
	s.OpenTeamId = &v
	return s
}

func (s *BatchQueryCustomerSendTaskRequest) SetTaskName(v string) *BatchQueryCustomerSendTaskRequest {
	s.TaskName = &v
	return s
}

type BatchQueryCustomerSendTaskResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8888
	NextToken *string                                          `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Records   []*BatchQueryCustomerSendTaskResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s BatchQueryCustomerSendTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryCustomerSendTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryCustomerSendTaskResponseBody) SetMaxResults(v int64) *BatchQueryCustomerSendTaskResponseBody {
	s.MaxResults = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBody) SetNextToken(v string) *BatchQueryCustomerSendTaskResponseBody {
	s.NextToken = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBody) SetRecords(v []*BatchQueryCustomerSendTaskResponseBodyRecords) *BatchQueryCustomerSendTaskResponseBody {
	s.Records = v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBody) SetTotalCount(v int64) *BatchQueryCustomerSendTaskResponseBody {
	s.TotalCount = &v
	return s
}

type BatchQueryCustomerSendTaskResponseBodyRecords struct {
	// example:
	//
	// 张三
	CreateName *string `json:"createName,omitempty" xml:"createName,omitempty"`
	// example:
	//
	// 2023-07-14 10:00:00
	CreateTimeStr *string `json:"createTimeStr,omitempty" xml:"createTimeStr,omitempty"`
	// example:
	//
	// 88888
	CreateUnionId *string `json:"createUnionId,omitempty" xml:"createUnionId,omitempty"`
	// example:
	//
	// 88888
	OpenBatchTaskId *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
	// example:
	//
	// 90
	ReadCustomerInc *int64 `json:"readCustomerInc,omitempty" xml:"readCustomerInc,omitempty"`
	// example:
	//
	// 100
	ReadUserInc *int64 `json:"readUserInc,omitempty" xml:"readUserInc,omitempty"`
	// example:
	//
	// 100
	SendCustomerInc *int64 `json:"sendCustomerInc,omitempty" xml:"sendCustomerInc,omitempty"`
	// example:
	//
	// UNFINISH 未完成 FINISHED 已发送 CANCELED 已取消 CREATE_TASK_FAILED 创建任务失败  SENDING 发送中
	SendMessageStatus *string `json:"sendMessageStatus,omitempty" xml:"sendMessageStatus,omitempty"`
	// example:
	//
	// 2023-07-14 11:00:00
	SendTaskTimeStr *string `json:"sendTaskTimeStr,omitempty" xml:"sendTaskTimeStr,omitempty"`
	// example:
	//
	// 200
	SendUserInc *int64 `json:"sendUserInc,omitempty" xml:"sendUserInc,omitempty"`
	// example:
	//
	// 任务名称
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s BatchQueryCustomerSendTaskResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryCustomerSendTaskResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetCreateName(v string) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.CreateName = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetCreateTimeStr(v string) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.CreateTimeStr = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetCreateUnionId(v string) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.CreateUnionId = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetOpenBatchTaskId(v string) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.OpenBatchTaskId = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetReadCustomerInc(v int64) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.ReadCustomerInc = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetReadUserInc(v int64) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.ReadUserInc = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetSendCustomerInc(v int64) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.SendCustomerInc = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetSendMessageStatus(v string) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.SendMessageStatus = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetSendTaskTimeStr(v string) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.SendTaskTimeStr = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetSendUserInc(v int64) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.SendUserInc = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponseBodyRecords) SetTaskName(v string) *BatchQueryCustomerSendTaskResponseBodyRecords {
	s.TaskName = &v
	return s
}

type BatchQueryCustomerSendTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryCustomerSendTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryCustomerSendTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryCustomerSendTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryCustomerSendTaskResponse) SetHeaders(v map[string]*string) *BatchQueryCustomerSendTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryCustomerSendTaskResponse) SetStatusCode(v int32) *BatchQueryCustomerSendTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryCustomerSendTaskResponse) SetBody(v *BatchQueryCustomerSendTaskResponseBody) *BatchQueryCustomerSendTaskResponse {
	s.Body = v
	return s
}

type BatchQueryGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryGroupMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// cid***
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 888
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s BatchQueryGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberRequest) SetMaxResults(v int64) *BatchQueryGroupMemberRequest {
	s.MaxResults = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetNextToken(v string) *BatchQueryGroupMemberRequest {
	s.NextToken = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetOpenConversationId(v string) *BatchQueryGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetOpenTeamId(v string) *BatchQueryGroupMemberRequest {
	s.OpenTeamId = &v
	return s
}

type BatchQueryGroupMemberResponseBody struct {
	HasMore            *bool                                       `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken          *string                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenConversationId *string                                     `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Records            []*BatchQueryGroupMemberResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s BatchQueryGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberResponseBody) SetHasMore(v bool) *BatchQueryGroupMemberResponseBody {
	s.HasMore = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetNextToken(v string) *BatchQueryGroupMemberResponseBody {
	s.NextToken = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetOpenConversationId(v string) *BatchQueryGroupMemberResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetRecords(v []*BatchQueryGroupMemberResponseBodyRecords) *BatchQueryGroupMemberResponseBody {
	s.Records = v
	return s
}

type BatchQueryGroupMemberResponseBodyRecords struct {
	InnerStaff *bool   `json:"innerStaff,omitempty" xml:"innerStaff,omitempty"`
	NickName   *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Owner      *bool   `json:"owner,omitempty" xml:"owner,omitempty"`
	UnionId    *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchQueryGroupMemberResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberResponseBodyRecords) SetInnerStaff(v bool) *BatchQueryGroupMemberResponseBodyRecords {
	s.InnerStaff = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBodyRecords) SetNickName(v string) *BatchQueryGroupMemberResponseBodyRecords {
	s.NickName = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBodyRecords) SetOwner(v bool) *BatchQueryGroupMemberResponseBodyRecords {
	s.Owner = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBodyRecords) SetUnionId(v string) *BatchQueryGroupMemberResponseBodyRecords {
	s.UnionId = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBodyRecords) SetUserId(v string) *BatchQueryGroupMemberResponseBodyRecords {
	s.UserId = &v
	return s
}

type BatchQueryGroupMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberResponse) SetHeaders(v map[string]*string) *BatchQueryGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryGroupMemberResponse) SetStatusCode(v int32) *BatchQueryGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryGroupMemberResponse) SetBody(v *BatchQueryGroupMemberResponseBody) *BatchQueryGroupMemberResponse {
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
	// example:
	//
	// false
	GetReadCount *bool `json:"getReadCount,omitempty" xml:"getReadCount,omitempty"`
	// example:
	//
	// 2022-04-02 00:00:00
	GmtCreateEnd *string `json:"gmtCreateEnd,omitempty" xml:"gmtCreateEnd,omitempty"`
	// example:
	//
	// 2022-04-01 00:00:00
	GmtCreateStart *string `json:"gmtCreateStart,omitempty" xml:"gmtCreateStart,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 首页传递空
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// fwPuycdHiiI
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jciwnfw
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 群发任务双11
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
	MaxResults *int64                                          `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
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
	// This parameter is required.
	SendTaskTimeStr *string `json:"sendTaskTimeStr,omitempty" xml:"sendTaskTimeStr,omitempty"`
	TaskName        *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQuerySendMessageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *BatchQuerySendMessageTaskResponse) SetStatusCode(v int32) *BatchQuerySendMessageTaskResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// btkoYsadwyQiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"robotCode":"123ITJovyMHtmi216233798228941001","robotName":"服务小钉"}]
	RobotConfig  *string `json:"robotConfig,omitempty" xml:"robotConfig,omitempty"`
	TemplateDesc *string `json:"templateDesc,omitempty" xml:"templateDesc,omitempty"`
	// This parameter is required.
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0普通群模板，1内部群模板
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BoundTemplateToTeamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *BoundTemplateToTeamResponse) SetStatusCode(v int32) *BoundTemplateToTeamResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a8iS4X94TgtgiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// example:
	//
	// Dq9hP8Sk2v6vQ6l05nCe5wiEiE
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
	// example:
	//
	// true
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
	// example:
	//
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
	// example:
	//
	// wahaha.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// ticket/image/44708069/43003/e27204b382c04832aec4243e940a1367_1625831640499.txt
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *CancelTicketResponse) SetStatusCode(v int32) *CancelTicketResponse {
	s.StatusCode = &v
	return s
}

type CategoryStatisticsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CategoryStatisticsHeaders) String() string {
	return tea.Prettify(s)
}

func (s CategoryStatisticsHeaders) GoString() string {
	return s.String()
}

func (s *CategoryStatisticsHeaders) SetCommonHeaders(v map[string]*string) *CategoryStatisticsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CategoryStatisticsHeaders) SetXAcsDingtalkAccessToken(v string) *CategoryStatisticsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CategoryStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MaxDt *string `json:"maxDt,omitempty" xml:"maxDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MinDt *string `json:"minDt,omitempty" xml:"minDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KxisoOk
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s CategoryStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s CategoryStatisticsRequest) GoString() string {
	return s.String()
}

func (s *CategoryStatisticsRequest) SetMaxDt(v string) *CategoryStatisticsRequest {
	s.MaxDt = &v
	return s
}

func (s *CategoryStatisticsRequest) SetMinDt(v string) *CategoryStatisticsRequest {
	s.MinDt = &v
	return s
}

func (s *CategoryStatisticsRequest) SetOpenTeamId(v string) *CategoryStatisticsRequest {
	s.OpenTeamId = &v
	return s
}

type CategoryStatisticsResponseBody struct {
	// This parameter is required.
	CategoryStatisticsRecords []*CategoryStatisticsResponseBodyCategoryStatisticsRecords `json:"categoryStatisticsRecords,omitempty" xml:"categoryStatisticsRecords,omitempty" type:"Repeated"`
	// This parameter is required.
	CategoryTrend []*CategoryStatisticsResponseBodyCategoryTrend `json:"categoryTrend,omitempty" xml:"categoryTrend,omitempty" type:"Repeated"`
}

func (s CategoryStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CategoryStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *CategoryStatisticsResponseBody) SetCategoryStatisticsRecords(v []*CategoryStatisticsResponseBodyCategoryStatisticsRecords) *CategoryStatisticsResponseBody {
	s.CategoryStatisticsRecords = v
	return s
}

func (s *CategoryStatisticsResponseBody) SetCategoryTrend(v []*CategoryStatisticsResponseBodyCategoryTrend) *CategoryStatisticsResponseBody {
	s.CategoryTrend = v
	return s
}

type CategoryStatisticsResponseBodyCategoryStatisticsRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9
	LastCount *int64 `json:"lastCount,omitempty" xml:"lastCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 工单类
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CategoryStatisticsResponseBodyCategoryStatisticsRecords) String() string {
	return tea.Prettify(s)
}

func (s CategoryStatisticsResponseBodyCategoryStatisticsRecords) GoString() string {
	return s.String()
}

func (s *CategoryStatisticsResponseBodyCategoryStatisticsRecords) SetCount(v int64) *CategoryStatisticsResponseBodyCategoryStatisticsRecords {
	s.Count = &v
	return s
}

func (s *CategoryStatisticsResponseBodyCategoryStatisticsRecords) SetLastCount(v int64) *CategoryStatisticsResponseBodyCategoryStatisticsRecords {
	s.LastCount = &v
	return s
}

func (s *CategoryStatisticsResponseBodyCategoryStatisticsRecords) SetName(v string) *CategoryStatisticsResponseBodyCategoryStatisticsRecords {
	s.Name = &v
	return s
}

type CategoryStatisticsResponseBodyCategoryTrend struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	Dt *string `json:"dt,omitempty" xml:"dt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 工单类
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CategoryStatisticsResponseBodyCategoryTrend) String() string {
	return tea.Prettify(s)
}

func (s CategoryStatisticsResponseBodyCategoryTrend) GoString() string {
	return s.String()
}

func (s *CategoryStatisticsResponseBodyCategoryTrend) SetCount(v int64) *CategoryStatisticsResponseBodyCategoryTrend {
	s.Count = &v
	return s
}

func (s *CategoryStatisticsResponseBodyCategoryTrend) SetDt(v string) *CategoryStatisticsResponseBodyCategoryTrend {
	s.Dt = &v
	return s
}

func (s *CategoryStatisticsResponseBodyCategoryTrend) SetName(v string) *CategoryStatisticsResponseBodyCategoryTrend {
	s.Name = &v
	return s
}

type CategoryStatisticsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CategoryStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CategoryStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s CategoryStatisticsResponse) GoString() string {
	return s.String()
}

func (s *CategoryStatisticsResponse) SetHeaders(v map[string]*string) *CategoryStatisticsResponse {
	s.Headers = v
	return s
}

func (s *CategoryStatisticsResponse) SetStatusCode(v int32) *CategoryStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *CategoryStatisticsResponse) SetBody(v *CategoryStatisticsResponseBody) *CategoryStatisticsResponse {
	s.Body = v
	return s
}

type CloseConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CloseConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s CloseConversationHeaders) GoString() string {
	return s.String()
}

func (s *CloseConversationHeaders) SetCommonHeaders(v map[string]*string) *CloseConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseConversationHeaders) SetXAcsDingtalkAccessToken(v string) *CloseConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CloseConversationRequest struct {
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	OpenTeamId     *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	ServerTips     *string `json:"serverTips,omitempty" xml:"serverTips,omitempty"`
	ServiceToken   *string `json:"serviceToken,omitempty" xml:"serviceToken,omitempty"`
	TargetChannel  *string `json:"targetChannel,omitempty" xml:"targetChannel,omitempty"`
	VisitorToken   *string `json:"visitorToken,omitempty" xml:"visitorToken,omitempty"`
}

func (s CloseConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseConversationRequest) GoString() string {
	return s.String()
}

func (s *CloseConversationRequest) SetConversationId(v string) *CloseConversationRequest {
	s.ConversationId = &v
	return s
}

func (s *CloseConversationRequest) SetOpenTeamId(v string) *CloseConversationRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CloseConversationRequest) SetServerTips(v string) *CloseConversationRequest {
	s.ServerTips = &v
	return s
}

func (s *CloseConversationRequest) SetServiceToken(v string) *CloseConversationRequest {
	s.ServiceToken = &v
	return s
}

func (s *CloseConversationRequest) SetTargetChannel(v string) *CloseConversationRequest {
	s.TargetChannel = &v
	return s
}

func (s *CloseConversationRequest) SetVisitorToken(v string) *CloseConversationRequest {
	s.VisitorToken = &v
	return s
}

type CloseConversationResponseBody struct {
	DingOpenErrcode *int32  `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success         *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CloseConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CloseConversationResponseBody) SetDingOpenErrcode(v int32) *CloseConversationResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *CloseConversationResponseBody) SetErrorMsg(v string) *CloseConversationResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CloseConversationResponseBody) SetResult(v bool) *CloseConversationResponseBody {
	s.Result = &v
	return s
}

func (s *CloseConversationResponseBody) SetSuccess(v bool) *CloseConversationResponseBody {
	s.Success = &v
	return s
}

type CloseConversationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseConversationResponse) GoString() string {
	return s.String()
}

func (s *CloseConversationResponse) SetHeaders(v map[string]*string) *CloseConversationResponse {
	s.Headers = v
	return s
}

func (s *CloseConversationResponse) SetStatusCode(v int32) *CloseConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseConversationResponse) SetBody(v *CloseConversationResponseBody) *CloseConversationResponse {
	s.Body = v
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
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseHumanSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CloseHumanSessionResponse) SetStatusCode(v int32) *CloseHumanSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseHumanSessionResponse) SetBody(v *CloseHumanSessionResponseBody) *CloseHumanSessionResponse {
	s.Body = v
	return s
}

type ConversationCreatedNotifyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConversationCreatedNotifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConversationCreatedNotifyHeaders) GoString() string {
	return s.String()
}

func (s *ConversationCreatedNotifyHeaders) SetCommonHeaders(v map[string]*string) *ConversationCreatedNotifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConversationCreatedNotifyHeaders) SetXAcsDingtalkAccessToken(v string) *ConversationCreatedNotifyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConversationCreatedNotifyRequest struct {
	AlipayUserId   *string `json:"alipayUserId,omitempty" xml:"alipayUserId,omitempty"`
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	NickName       *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// eWaJSqDcLsoiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty"`
	ServerTips *string `json:"serverTips,omitempty" xml:"serverTips,omitempty"`
	// example:
	//
	// 对应外部渠道的会话ID
	ServiceToken      *string `json:"serviceToken,omitempty" xml:"serviceToken,omitempty"`
	TimeoutRemindTips *string `json:"timeoutRemindTips,omitempty" xml:"timeoutRemindTips,omitempty"`
	UserId            *string `json:"userId,omitempty" xml:"userId,omitempty"`
	VisitorToken      *string `json:"visitorToken,omitempty" xml:"visitorToken,omitempty"`
}

func (s ConversationCreatedNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ConversationCreatedNotifyRequest) GoString() string {
	return s.String()
}

func (s *ConversationCreatedNotifyRequest) SetAlipayUserId(v string) *ConversationCreatedNotifyRequest {
	s.AlipayUserId = &v
	return s
}

func (s *ConversationCreatedNotifyRequest) SetConversationId(v string) *ConversationCreatedNotifyRequest {
	s.ConversationId = &v
	return s
}

func (s *ConversationCreatedNotifyRequest) SetNickName(v string) *ConversationCreatedNotifyRequest {
	s.NickName = &v
	return s
}

func (s *ConversationCreatedNotifyRequest) SetOpenTeamId(v string) *ConversationCreatedNotifyRequest {
	s.OpenTeamId = &v
	return s
}

func (s *ConversationCreatedNotifyRequest) SetServerName(v string) *ConversationCreatedNotifyRequest {
	s.ServerName = &v
	return s
}

func (s *ConversationCreatedNotifyRequest) SetServerTips(v string) *ConversationCreatedNotifyRequest {
	s.ServerTips = &v
	return s
}

func (s *ConversationCreatedNotifyRequest) SetServiceToken(v string) *ConversationCreatedNotifyRequest {
	s.ServiceToken = &v
	return s
}

func (s *ConversationCreatedNotifyRequest) SetTimeoutRemindTips(v string) *ConversationCreatedNotifyRequest {
	s.TimeoutRemindTips = &v
	return s
}

func (s *ConversationCreatedNotifyRequest) SetUserId(v string) *ConversationCreatedNotifyRequest {
	s.UserId = &v
	return s
}

func (s *ConversationCreatedNotifyRequest) SetVisitorToken(v string) *ConversationCreatedNotifyRequest {
	s.VisitorToken = &v
	return s
}

type ConversationCreatedNotifyResponseBody struct {
	DingOpenErrcode *int32  `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConversationCreatedNotifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConversationCreatedNotifyResponseBody) GoString() string {
	return s.String()
}

func (s *ConversationCreatedNotifyResponseBody) SetDingOpenErrcode(v int32) *ConversationCreatedNotifyResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *ConversationCreatedNotifyResponseBody) SetErrorMsg(v string) *ConversationCreatedNotifyResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ConversationCreatedNotifyResponseBody) SetResult(v bool) *ConversationCreatedNotifyResponseBody {
	s.Result = &v
	return s
}

func (s *ConversationCreatedNotifyResponseBody) SetSuccess(v bool) *ConversationCreatedNotifyResponseBody {
	s.Success = &v
	return s
}

type ConversationCreatedNotifyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConversationCreatedNotifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConversationCreatedNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ConversationCreatedNotifyResponse) GoString() string {
	return s.String()
}

func (s *ConversationCreatedNotifyResponse) SetHeaders(v map[string]*string) *ConversationCreatedNotifyResponse {
	s.Headers = v
	return s
}

func (s *ConversationCreatedNotifyResponse) SetStatusCode(v int32) *ConversationCreatedNotifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ConversationCreatedNotifyResponse) SetBody(v *ConversationCreatedNotifyResponseBody) *ConversationCreatedNotifyResponse {
	s.Body = v
	return s
}

type ConversationTransferBeginNotifyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConversationTransferBeginNotifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConversationTransferBeginNotifyHeaders) GoString() string {
	return s.String()
}

func (s *ConversationTransferBeginNotifyHeaders) SetCommonHeaders(v map[string]*string) *ConversationTransferBeginNotifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConversationTransferBeginNotifyHeaders) SetXAcsDingtalkAccessToken(v string) *ConversationTransferBeginNotifyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConversationTransferBeginNotifyRequest struct {
	ConversationId     *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	Memo               *string `json:"memo,omitempty" xml:"memo,omitempty"`
	OpenTeamId         *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	ServiceToken       *string `json:"serviceToken,omitempty" xml:"serviceToken,omitempty"`
	SourceSkillGroupId *string `json:"sourceSkillGroupId,omitempty" xml:"sourceSkillGroupId,omitempty"`
	TargetSkillGroupId *string `json:"targetSkillGroupId,omitempty" xml:"targetSkillGroupId,omitempty"`
}

func (s ConversationTransferBeginNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ConversationTransferBeginNotifyRequest) GoString() string {
	return s.String()
}

func (s *ConversationTransferBeginNotifyRequest) SetConversationId(v string) *ConversationTransferBeginNotifyRequest {
	s.ConversationId = &v
	return s
}

func (s *ConversationTransferBeginNotifyRequest) SetMemo(v string) *ConversationTransferBeginNotifyRequest {
	s.Memo = &v
	return s
}

func (s *ConversationTransferBeginNotifyRequest) SetOpenTeamId(v string) *ConversationTransferBeginNotifyRequest {
	s.OpenTeamId = &v
	return s
}

func (s *ConversationTransferBeginNotifyRequest) SetServiceToken(v string) *ConversationTransferBeginNotifyRequest {
	s.ServiceToken = &v
	return s
}

func (s *ConversationTransferBeginNotifyRequest) SetSourceSkillGroupId(v string) *ConversationTransferBeginNotifyRequest {
	s.SourceSkillGroupId = &v
	return s
}

func (s *ConversationTransferBeginNotifyRequest) SetTargetSkillGroupId(v string) *ConversationTransferBeginNotifyRequest {
	s.TargetSkillGroupId = &v
	return s
}

type ConversationTransferBeginNotifyResponseBody struct {
	DingOpenErrcode *int32  `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success         *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConversationTransferBeginNotifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConversationTransferBeginNotifyResponseBody) GoString() string {
	return s.String()
}

func (s *ConversationTransferBeginNotifyResponseBody) SetDingOpenErrcode(v int32) *ConversationTransferBeginNotifyResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *ConversationTransferBeginNotifyResponseBody) SetErrorMsg(v string) *ConversationTransferBeginNotifyResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ConversationTransferBeginNotifyResponseBody) SetResult(v bool) *ConversationTransferBeginNotifyResponseBody {
	s.Result = &v
	return s
}

func (s *ConversationTransferBeginNotifyResponseBody) SetSuccess(v bool) *ConversationTransferBeginNotifyResponseBody {
	s.Success = &v
	return s
}

type ConversationTransferBeginNotifyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConversationTransferBeginNotifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConversationTransferBeginNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ConversationTransferBeginNotifyResponse) GoString() string {
	return s.String()
}

func (s *ConversationTransferBeginNotifyResponse) SetHeaders(v map[string]*string) *ConversationTransferBeginNotifyResponse {
	s.Headers = v
	return s
}

func (s *ConversationTransferBeginNotifyResponse) SetStatusCode(v int32) *ConversationTransferBeginNotifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ConversationTransferBeginNotifyResponse) SetBody(v *ConversationTransferBeginNotifyResponseBody) *ConversationTransferBeginNotifyResponse {
	s.Body = v
	return s
}

type ConversationTransferCompleteNotifyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConversationTransferCompleteNotifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConversationTransferCompleteNotifyHeaders) GoString() string {
	return s.String()
}

func (s *ConversationTransferCompleteNotifyHeaders) SetCommonHeaders(v map[string]*string) *ConversationTransferCompleteNotifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConversationTransferCompleteNotifyHeaders) SetXAcsDingtalkAccessToken(v string) *ConversationTransferCompleteNotifyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConversationTransferCompleteNotifyRequest struct {
	AlipayUserId   *string `json:"alipayUserId,omitempty" xml:"alipayUserId,omitempty"`
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	NickName       *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	OpenTeamId     *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	ServiceToken   *string `json:"serviceToken,omitempty" xml:"serviceToken,omitempty"`
	UserId         *string `json:"userId,omitempty" xml:"userId,omitempty"`
	VisitorToken   *string `json:"visitorToken,omitempty" xml:"visitorToken,omitempty"`
}

func (s ConversationTransferCompleteNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ConversationTransferCompleteNotifyRequest) GoString() string {
	return s.String()
}

func (s *ConversationTransferCompleteNotifyRequest) SetAlipayUserId(v string) *ConversationTransferCompleteNotifyRequest {
	s.AlipayUserId = &v
	return s
}

func (s *ConversationTransferCompleteNotifyRequest) SetConversationId(v string) *ConversationTransferCompleteNotifyRequest {
	s.ConversationId = &v
	return s
}

func (s *ConversationTransferCompleteNotifyRequest) SetNickName(v string) *ConversationTransferCompleteNotifyRequest {
	s.NickName = &v
	return s
}

func (s *ConversationTransferCompleteNotifyRequest) SetOpenTeamId(v string) *ConversationTransferCompleteNotifyRequest {
	s.OpenTeamId = &v
	return s
}

func (s *ConversationTransferCompleteNotifyRequest) SetServiceToken(v string) *ConversationTransferCompleteNotifyRequest {
	s.ServiceToken = &v
	return s
}

func (s *ConversationTransferCompleteNotifyRequest) SetUserId(v string) *ConversationTransferCompleteNotifyRequest {
	s.UserId = &v
	return s
}

func (s *ConversationTransferCompleteNotifyRequest) SetVisitorToken(v string) *ConversationTransferCompleteNotifyRequest {
	s.VisitorToken = &v
	return s
}

type ConversationTransferCompleteNotifyResponseBody struct {
	DingOpenErrcode *int32  `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success         *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConversationTransferCompleteNotifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConversationTransferCompleteNotifyResponseBody) GoString() string {
	return s.String()
}

func (s *ConversationTransferCompleteNotifyResponseBody) SetDingOpenErrcode(v int32) *ConversationTransferCompleteNotifyResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *ConversationTransferCompleteNotifyResponseBody) SetErrorMsg(v string) *ConversationTransferCompleteNotifyResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ConversationTransferCompleteNotifyResponseBody) SetResult(v bool) *ConversationTransferCompleteNotifyResponseBody {
	s.Result = &v
	return s
}

func (s *ConversationTransferCompleteNotifyResponseBody) SetSuccess(v bool) *ConversationTransferCompleteNotifyResponseBody {
	s.Success = &v
	return s
}

type ConversationTransferCompleteNotifyResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConversationTransferCompleteNotifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConversationTransferCompleteNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ConversationTransferCompleteNotifyResponse) GoString() string {
	return s.String()
}

func (s *ConversationTransferCompleteNotifyResponse) SetHeaders(v map[string]*string) *ConversationTransferCompleteNotifyResponse {
	s.Headers = v
	return s
}

func (s *ConversationTransferCompleteNotifyResponse) SetStatusCode(v int32) *ConversationTransferCompleteNotifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ConversationTransferCompleteNotifyResponse) SetBody(v *ConversationTransferCompleteNotifyResponseBody) *ConversationTransferCompleteNotifyResponse {
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
	// example:
	//
	// PID123cjj2
	GroupBizId *string `json:"groupBizId,omitempty" xml:"groupBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试服务群
	GroupName      *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupTagNames  []*string `json:"groupTagNames,omitempty" xml:"groupTagNames,omitempty" type:"Repeated"`
	MemberStaffIds []*string `json:"memberStaffIds,omitempty" xml:"memberStaffIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Jciwnfw
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jciwnfw
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager123
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
	// This parameter is required.
	//
	// example:
	//
	// http://qr.dingtalk.com/xxxxx
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxxx==
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateGroupConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupConversationHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupConversationRequest struct {
	// example:
	//
	// dingadc88253b4d581bd35c2f4657eb6378f
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// fsfsfadfasfdasdfsaf
	DingGroupId        *string `json:"dingGroupId,omitempty" xml:"dingGroupId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// example:
	//
	// 57675657
	DingUserId *string `json:"dingUserId,omitempty" xml:"dingUserId,omitempty"`
	// example:
	//
	// 张三
	DingUserName *string `json:"dingUserName,omitempty" xml:"dingUserName,omitempty"`
	// example:
	//
	// {"isServerInitiative":"true"}
	ExtValues  *string `json:"extValues,omitempty" xml:"extValues,omitempty"`
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 3434
	ServerGroupId *string `json:"serverGroupId,omitempty" xml:"serverGroupId,omitempty"`
}

func (s CreateGroupConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationRequest) SetCorpId(v string) *CreateGroupConversationRequest {
	s.CorpId = &v
	return s
}

func (s *CreateGroupConversationRequest) SetDingGroupId(v string) *CreateGroupConversationRequest {
	s.DingGroupId = &v
	return s
}

func (s *CreateGroupConversationRequest) SetDingSuiteKey(v string) *CreateGroupConversationRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *CreateGroupConversationRequest) SetDingTokenGrantType(v int64) *CreateGroupConversationRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *CreateGroupConversationRequest) SetDingUserId(v string) *CreateGroupConversationRequest {
	s.DingUserId = &v
	return s
}

func (s *CreateGroupConversationRequest) SetDingUserName(v string) *CreateGroupConversationRequest {
	s.DingUserName = &v
	return s
}

func (s *CreateGroupConversationRequest) SetExtValues(v string) *CreateGroupConversationRequest {
	s.ExtValues = &v
	return s
}

func (s *CreateGroupConversationRequest) SetOpenTeamId(v string) *CreateGroupConversationRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CreateGroupConversationRequest) SetServerGroupId(v string) *CreateGroupConversationRequest {
	s.ServerGroupId = &v
	return s
}

type CreateGroupConversationResponseBody struct {
	// example:
	//
	// 500
	DingOpenErrcode *int32 `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// true
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateGroupConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationResponseBody) SetDingOpenErrcode(v int32) *CreateGroupConversationResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *CreateGroupConversationResponseBody) SetErrorMsg(v string) *CreateGroupConversationResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateGroupConversationResponseBody) SetResult(v string) *CreateGroupConversationResponseBody {
	s.Result = &v
	return s
}

func (s *CreateGroupConversationResponseBody) SetSuccess(v bool) *CreateGroupConversationResponseBody {
	s.Success = &v
	return s
}

type CreateGroupConversationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationResponse) SetHeaders(v map[string]*string) *CreateGroupConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupConversationResponse) SetStatusCode(v int32) *CreateGroupConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupConversationResponse) SetBody(v *CreateGroupConversationResponseBody) *CreateGroupConversationResponse {
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
	// This parameter is required.
	GroupSetName *string `json:"groupSetName,omitempty" xml:"groupSetName,omitempty"`
	// This parameter is required.
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateGroupSetResponse) SetStatusCode(v int32) *CreateGroupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupSetResponse) SetBody(v *CreateGroupSetResponseBody) *CreateGroupSetResponse {
	s.Body = v
	return s
}

type CreateInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceHeaders) GoString() string {
	return s.String()
}

func (s *CreateInstanceHeaders) SetCommonHeaders(v map[string]*string) *CreateInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateInstanceRequest struct {
	// example:
	//
	// DOU_YIN
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// example:
	//
	// 888888
	ExternalBizId *string `json:"externalBizId,omitempty" xml:"externalBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DING_CUSTOMER
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"node_1111":"hahha"}
	FormDataList *string `json:"formDataList,omitempty" xml:"formDataList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 88444***
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 88855
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	// example:
	//
	// 88855
	OwnerUnionId *string `json:"ownerUnionId,omitempty" xml:"ownerUnionId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetChannel(v string) *CreateInstanceRequest {
	s.Channel = &v
	return s
}

func (s *CreateInstanceRequest) SetExternalBizId(v string) *CreateInstanceRequest {
	s.ExternalBizId = &v
	return s
}

func (s *CreateInstanceRequest) SetFormCode(v string) *CreateInstanceRequest {
	s.FormCode = &v
	return s
}

func (s *CreateInstanceRequest) SetFormDataList(v string) *CreateInstanceRequest {
	s.FormDataList = &v
	return s
}

func (s *CreateInstanceRequest) SetOpenTeamId(v string) *CreateInstanceRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CreateInstanceRequest) SetOperatorUnionId(v string) *CreateInstanceRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerUnionId(v string) *CreateInstanceRequest {
	s.OwnerUnionId = &v
	return s
}

type CreateInstanceResponseBody struct {
	// example:
	//
	// 8888***
	OpenDataInstanceId *string `json:"openDataInstanceId,omitempty" xml:"openDataInstanceId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetOpenDataInstanceId(v string) *CreateInstanceResponseBody {
	s.OpenDataInstanceId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
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
	// This parameter is required.
	CreatorDingUnionId *string `json:"creatorDingUnionId,omitempty" xml:"creatorDingUnionId,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTeamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateTeamResponse) SetStatusCode(v int32) *CreateTeamResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// Dq9hP8Sk2v6vQ6l05nCe5wiEiE
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// example:
	//
	// [{\"identifier\":\"input1\",\"value\":\"123\"}]
	CustomFields *string                    `json:"customFields,omitempty" xml:"customFields,omitempty"`
	Notify       *CreateTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bLkvfXKiSngQiE
	OpenTemplateBizId *string `json:"openTemplateBizId,omitempty" xml:"openTemplateBizId,omitempty"`
	// This parameter is required.
	ProcessorUnionIds []*string `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// SG
	Scene        *string                          `json:"scene,omitempty" xml:"scene,omitempty"`
	SceneContext *CreateTicketRequestSceneContext `json:"sceneContext,omitempty" xml:"sceneContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
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
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	NoticeAllGroupMember       *bool     `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
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
	GroupMsgs []*CreateTicketRequestSceneContextGroupMsgs `json:"groupMsgs,omitempty" xml:"groupMsgs,omitempty" type:"Repeated"`
	// example:
	//
	// cidZBSNlUi/Jq9x76PAXUCrAA==
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	RelevantorUnionIds []*string `json:"relevantorUnionIds,omitempty" xml:"relevantorUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// a0ba57d5d29a48b51d0eca48da6b1d09
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
	// example:
	//
	// true
	Anchor *bool `json:"anchor,omitempty" xml:"anchor,omitempty"`
	// example:
	//
	// msgsbY4BzTCNX0/ClUwoTTs7w==
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateTicketResponse) SetStatusCode(v int32) *CreateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
	s.Body = v
	return s
}

type CustomerSendMsgTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomerSendMsgTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomerSendMsgTaskHeaders) GoString() string {
	return s.String()
}

func (s *CustomerSendMsgTaskHeaders) SetCommonHeaders(v map[string]*string) *CustomerSendMsgTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomerSendMsgTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CustomerSendMsgTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomerSendMsgTaskRequest struct {
	// This parameter is required.
	MessageContent *CustomerSendMsgTaskRequestMessageContent `json:"messageContent,omitempty" xml:"messageContent,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 8888
	OpenTeamId    *string                                  `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	QueryCustomer *CustomerSendMsgTaskRequestQueryCustomer `json:"queryCustomer,omitempty" xml:"queryCustomer,omitempty" type:"Struct"`
	SendConfig    *CustomerSendMsgTaskRequestSendConfig    `json:"sendConfig,omitempty" xml:"sendConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 任务名称
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s CustomerSendMsgTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomerSendMsgTaskRequest) GoString() string {
	return s.String()
}

func (s *CustomerSendMsgTaskRequest) SetMessageContent(v *CustomerSendMsgTaskRequestMessageContent) *CustomerSendMsgTaskRequest {
	s.MessageContent = v
	return s
}

func (s *CustomerSendMsgTaskRequest) SetOpenTeamId(v string) *CustomerSendMsgTaskRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CustomerSendMsgTaskRequest) SetQueryCustomer(v *CustomerSendMsgTaskRequestQueryCustomer) *CustomerSendMsgTaskRequest {
	s.QueryCustomer = v
	return s
}

func (s *CustomerSendMsgTaskRequest) SetSendConfig(v *CustomerSendMsgTaskRequestSendConfig) *CustomerSendMsgTaskRequest {
	s.SendConfig = v
	return s
}

func (s *CustomerSendMsgTaskRequest) SetTaskName(v string) *CustomerSendMsgTaskRequest {
	s.TaskName = &v
	return s
}

type CustomerSendMsgTaskRequestMessageContent struct {
	Btns []*CustomerSendMsgTaskRequestMessageContentBtns `json:"btns,omitempty" xml:"btns,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ACTIONCAR：卡片消息
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CustomerSendMsgTaskRequestMessageContent) String() string {
	return tea.Prettify(s)
}

func (s CustomerSendMsgTaskRequestMessageContent) GoString() string {
	return s.String()
}

func (s *CustomerSendMsgTaskRequestMessageContent) SetBtns(v []*CustomerSendMsgTaskRequestMessageContentBtns) *CustomerSendMsgTaskRequestMessageContent {
	s.Btns = v
	return s
}

func (s *CustomerSendMsgTaskRequestMessageContent) SetContent(v string) *CustomerSendMsgTaskRequestMessageContent {
	s.Content = &v
	return s
}

func (s *CustomerSendMsgTaskRequestMessageContent) SetMessageType(v string) *CustomerSendMsgTaskRequestMessageContent {
	s.MessageType = &v
	return s
}

func (s *CustomerSendMsgTaskRequestMessageContent) SetTitle(v string) *CustomerSendMsgTaskRequestMessageContent {
	s.Title = &v
	return s
}

type CustomerSendMsgTaskRequestMessageContentBtns struct {
	// example:
	//
	// http://www.baidu.com
	ActionURL *string `json:"actionURL,omitempty" xml:"actionURL,omitempty"`
	// example:
	//
	// 百度
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CustomerSendMsgTaskRequestMessageContentBtns) String() string {
	return tea.Prettify(s)
}

func (s CustomerSendMsgTaskRequestMessageContentBtns) GoString() string {
	return s.String()
}

func (s *CustomerSendMsgTaskRequestMessageContentBtns) SetActionURL(v string) *CustomerSendMsgTaskRequestMessageContentBtns {
	s.ActionURL = &v
	return s
}

func (s *CustomerSendMsgTaskRequestMessageContentBtns) SetTitle(v string) *CustomerSendMsgTaskRequestMessageContentBtns {
	s.Title = &v
	return s
}

type CustomerSendMsgTaskRequestQueryCustomer struct {
	OpenContactIds []*string `json:"openContactIds,omitempty" xml:"openContactIds,omitempty" type:"Repeated"`
	// example:
	//
	// AIMED
	QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	// example:
	//
	// {}
	SearchContactConditions *string `json:"searchContactConditions,omitempty" xml:"searchContactConditions,omitempty"`
	// example:
	//
	// {}
	SearchCustomerConditions *string `json:"searchCustomerConditions,omitempty" xml:"searchCustomerConditions,omitempty"`
}

func (s CustomerSendMsgTaskRequestQueryCustomer) String() string {
	return tea.Prettify(s)
}

func (s CustomerSendMsgTaskRequestQueryCustomer) GoString() string {
	return s.String()
}

func (s *CustomerSendMsgTaskRequestQueryCustomer) SetOpenContactIds(v []*string) *CustomerSendMsgTaskRequestQueryCustomer {
	s.OpenContactIds = v
	return s
}

func (s *CustomerSendMsgTaskRequestQueryCustomer) SetQueryType(v string) *CustomerSendMsgTaskRequestQueryCustomer {
	s.QueryType = &v
	return s
}

func (s *CustomerSendMsgTaskRequestQueryCustomer) SetSearchContactConditions(v string) *CustomerSendMsgTaskRequestQueryCustomer {
	s.SearchContactConditions = &v
	return s
}

func (s *CustomerSendMsgTaskRequestQueryCustomer) SetSearchCustomerConditions(v string) *CustomerSendMsgTaskRequestQueryCustomer {
	s.SearchCustomerConditions = &v
	return s
}

type CustomerSendMsgTaskRequestSendConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	NeedUrlTrack *bool `json:"needUrlTrack,omitempty" xml:"needUrlTrack,omitempty"`
	// example:
	//
	// 2023-06-01 00:00:00
	SendTime *string `json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANT
	SendType       *string                                               `json:"sendType,omitempty" xml:"sendType,omitempty"`
	UrlTrackConfig []*CustomerSendMsgTaskRequestSendConfigUrlTrackConfig `json:"urlTrackConfig,omitempty" xml:"urlTrackConfig,omitempty" type:"Repeated"`
}

func (s CustomerSendMsgTaskRequestSendConfig) String() string {
	return tea.Prettify(s)
}

func (s CustomerSendMsgTaskRequestSendConfig) GoString() string {
	return s.String()
}

func (s *CustomerSendMsgTaskRequestSendConfig) SetNeedUrlTrack(v bool) *CustomerSendMsgTaskRequestSendConfig {
	s.NeedUrlTrack = &v
	return s
}

func (s *CustomerSendMsgTaskRequestSendConfig) SetSendTime(v string) *CustomerSendMsgTaskRequestSendConfig {
	s.SendTime = &v
	return s
}

func (s *CustomerSendMsgTaskRequestSendConfig) SetSendType(v string) *CustomerSendMsgTaskRequestSendConfig {
	s.SendType = &v
	return s
}

func (s *CustomerSendMsgTaskRequestSendConfig) SetUrlTrackConfig(v []*CustomerSendMsgTaskRequestSendConfigUrlTrackConfig) *CustomerSendMsgTaskRequestSendConfig {
	s.UrlTrackConfig = v
	return s
}

type CustomerSendMsgTaskRequestSendConfigUrlTrackConfig struct {
	// example:
	//
	// 百度
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 88888
	TrackId *string `json:"trackId,omitempty" xml:"trackId,omitempty"`
	// example:
	//
	// http://www.baidu.com
	TrackUrl *string `json:"trackUrl,omitempty" xml:"trackUrl,omitempty"`
}

func (s CustomerSendMsgTaskRequestSendConfigUrlTrackConfig) String() string {
	return tea.Prettify(s)
}

func (s CustomerSendMsgTaskRequestSendConfigUrlTrackConfig) GoString() string {
	return s.String()
}

func (s *CustomerSendMsgTaskRequestSendConfigUrlTrackConfig) SetTitle(v string) *CustomerSendMsgTaskRequestSendConfigUrlTrackConfig {
	s.Title = &v
	return s
}

func (s *CustomerSendMsgTaskRequestSendConfigUrlTrackConfig) SetTrackId(v string) *CustomerSendMsgTaskRequestSendConfigUrlTrackConfig {
	s.TrackId = &v
	return s
}

func (s *CustomerSendMsgTaskRequestSendConfigUrlTrackConfig) SetTrackUrl(v string) *CustomerSendMsgTaskRequestSendConfigUrlTrackConfig {
	s.TrackUrl = &v
	return s
}

type CustomerSendMsgTaskResponseBody struct {
	// example:
	//
	// 88888
	OpenBatchTaskId *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
}

func (s CustomerSendMsgTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomerSendMsgTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CustomerSendMsgTaskResponseBody) SetOpenBatchTaskId(v string) *CustomerSendMsgTaskResponseBody {
	s.OpenBatchTaskId = &v
	return s
}

type CustomerSendMsgTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomerSendMsgTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CustomerSendMsgTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomerSendMsgTaskResponse) GoString() string {
	return s.String()
}

func (s *CustomerSendMsgTaskResponse) SetHeaders(v map[string]*string) *CustomerSendMsgTaskResponse {
	s.Headers = v
	return s
}

func (s *CustomerSendMsgTaskResponse) SetStatusCode(v int32) *CustomerSendMsgTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomerSendMsgTaskResponse) SetBody(v *CustomerSendMsgTaskResponseBody) *CustomerSendMsgTaskResponse {
	s.Body = v
	return s
}

type DeleteGroupMembersFromGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteGroupMembersFromGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMembersFromGroupHeaders) GoString() string {
	return s.String()
}

func (s *DeleteGroupMembersFromGroupHeaders) SetCommonHeaders(v map[string]*string) *DeleteGroupMembersFromGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteGroupMembersFromGroupHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteGroupMembersFromGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteGroupMembersFromGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// GROUP：从群中删除；GROUP_SET：从群组中删除
	DeleteGroupType *string `json:"deleteGroupType,omitempty" xml:"deleteGroupType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8888
	MemberUnionId *string `json:"memberUnionId,omitempty" xml:"memberUnionId,omitempty"`
	// example:
	//
	// cid**
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 8888
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8888
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s DeleteGroupMembersFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMembersFromGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupMembersFromGroupRequest) SetDeleteGroupType(v string) *DeleteGroupMembersFromGroupRequest {
	s.DeleteGroupType = &v
	return s
}

func (s *DeleteGroupMembersFromGroupRequest) SetMemberUnionId(v string) *DeleteGroupMembersFromGroupRequest {
	s.MemberUnionId = &v
	return s
}

func (s *DeleteGroupMembersFromGroupRequest) SetOpenConversationId(v string) *DeleteGroupMembersFromGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *DeleteGroupMembersFromGroupRequest) SetOpenGroupSetId(v string) *DeleteGroupMembersFromGroupRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *DeleteGroupMembersFromGroupRequest) SetOpenTeamId(v string) *DeleteGroupMembersFromGroupRequest {
	s.OpenTeamId = &v
	return s
}

type DeleteGroupMembersFromGroupResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteGroupMembersFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMembersFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupMembersFromGroupResponseBody) SetResult(v bool) *DeleteGroupMembersFromGroupResponseBody {
	s.Result = &v
	return s
}

type DeleteGroupMembersFromGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupMembersFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupMembersFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMembersFromGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupMembersFromGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupMembersFromGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupMembersFromGroupResponse) SetStatusCode(v int32) *DeleteGroupMembersFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupMembersFromGroupResponse) SetBody(v *DeleteGroupMembersFromGroupResponseBody) *DeleteGroupMembersFromGroupResponse {
	s.Body = v
	return s
}

type DeleteInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteInstanceHeaders) SetCommonHeaders(v map[string]*string) *DeleteInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DING_CUSTOMER
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 888***
	OpenDataInstanceId *string `json:"openDataInstanceId,omitempty" xml:"openDataInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 888**
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 8889999
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetFormCode(v string) *DeleteInstanceRequest {
	s.FormCode = &v
	return s
}

func (s *DeleteInstanceRequest) SetOpenDataInstanceId(v string) *DeleteInstanceRequest {
	s.OpenDataInstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetOpenTeamId(v string) *DeleteInstanceRequest {
	s.OpenTeamId = &v
	return s
}

func (s *DeleteInstanceRequest) SetOperatorUnionId(v string) *DeleteInstanceRequest {
	s.OperatorUnionId = &v
	return s
}

type DeleteInstanceResponseBody struct {
	OpenDataInstanceId *string `json:"openDataInstanceId,omitempty" xml:"openDataInstanceId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetOpenDataInstanceId(v string) *DeleteInstanceResponseBody {
	s.OpenDataInstanceId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
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
	// example:
	//
	// xuvw1245
	LibraryKey *string `json:"libraryKey,omitempty" xml:"libraryKey,omitempty"`
	// example:
	//
	// Js1i0w3k
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// CCM
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// CCM-123
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteKnowledgeResponse) SetStatusCode(v int32) *DeleteKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKnowledgeResponse) SetBody(v *DeleteKnowledgeResponseBody) *DeleteKnowledgeResponse {
	s.Body = v
	return s
}

type EmotionStatisticsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EmotionStatisticsHeaders) String() string {
	return tea.Prettify(s)
}

func (s EmotionStatisticsHeaders) GoString() string {
	return s.String()
}

func (s *EmotionStatisticsHeaders) SetCommonHeaders(v map[string]*string) *EmotionStatisticsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EmotionStatisticsHeaders) SetXAcsDingtalkAccessToken(v string) *EmotionStatisticsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EmotionStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MaxDt *string `json:"maxDt,omitempty" xml:"maxDt,omitempty"`
	// example:
	//
	// 0.8
	MaxEmotion *float64 `json:"maxEmotion,omitempty" xml:"maxEmotion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MinDt *string `json:"minDt,omitempty" xml:"minDt,omitempty"`
	// example:
	//
	// 0
	MinEmotion *float64 `json:"minEmotion,omitempty" xml:"minEmotion,omitempty"`
	// example:
	//
	// cidXX,cidYY
	OpenConversationIds *string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty"`
	// example:
	//
	// ksdfosd
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KxisoOk
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s EmotionStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s EmotionStatisticsRequest) GoString() string {
	return s.String()
}

func (s *EmotionStatisticsRequest) SetMaxDt(v string) *EmotionStatisticsRequest {
	s.MaxDt = &v
	return s
}

func (s *EmotionStatisticsRequest) SetMaxEmotion(v float64) *EmotionStatisticsRequest {
	s.MaxEmotion = &v
	return s
}

func (s *EmotionStatisticsRequest) SetMinDt(v string) *EmotionStatisticsRequest {
	s.MinDt = &v
	return s
}

func (s *EmotionStatisticsRequest) SetMinEmotion(v float64) *EmotionStatisticsRequest {
	s.MinEmotion = &v
	return s
}

func (s *EmotionStatisticsRequest) SetOpenConversationIds(v string) *EmotionStatisticsRequest {
	s.OpenConversationIds = &v
	return s
}

func (s *EmotionStatisticsRequest) SetOpenGroupSetId(v string) *EmotionStatisticsRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *EmotionStatisticsRequest) SetOpenTeamId(v string) *EmotionStatisticsRequest {
	s.OpenTeamId = &v
	return s
}

type EmotionStatisticsResponseBody struct {
	// This parameter is required.
	EmotionStatisticsRecords []*EmotionStatisticsResponseBodyEmotionStatisticsRecords `json:"emotionStatisticsRecords,omitempty" xml:"emotionStatisticsRecords,omitempty" type:"Repeated"`
}

func (s EmotionStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EmotionStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *EmotionStatisticsResponseBody) SetEmotionStatisticsRecords(v []*EmotionStatisticsResponseBodyEmotionStatisticsRecords) *EmotionStatisticsResponseBody {
	s.EmotionStatisticsRecords = v
	return s
}

type EmotionStatisticsResponseBodyEmotionStatisticsRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	Dt *string `json:"dt,omitempty" xml:"dt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.3
	EmotionScore *float64 `json:"emotionScore,omitempty" xml:"emotionScore,omitempty"`
}

func (s EmotionStatisticsResponseBodyEmotionStatisticsRecords) String() string {
	return tea.Prettify(s)
}

func (s EmotionStatisticsResponseBodyEmotionStatisticsRecords) GoString() string {
	return s.String()
}

func (s *EmotionStatisticsResponseBodyEmotionStatisticsRecords) SetCount(v int64) *EmotionStatisticsResponseBodyEmotionStatisticsRecords {
	s.Count = &v
	return s
}

func (s *EmotionStatisticsResponseBodyEmotionStatisticsRecords) SetDt(v string) *EmotionStatisticsResponseBodyEmotionStatisticsRecords {
	s.Dt = &v
	return s
}

func (s *EmotionStatisticsResponseBodyEmotionStatisticsRecords) SetEmotionScore(v float64) *EmotionStatisticsResponseBodyEmotionStatisticsRecords {
	s.EmotionScore = &v
	return s
}

type EmotionStatisticsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EmotionStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EmotionStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s EmotionStatisticsResponse) GoString() string {
	return s.String()
}

func (s *EmotionStatisticsResponse) SetHeaders(v map[string]*string) *EmotionStatisticsResponse {
	s.Headers = v
	return s
}

func (s *EmotionStatisticsResponse) SetStatusCode(v int32) *EmotionStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *EmotionStatisticsResponse) SetBody(v *EmotionStatisticsResponseBody) *EmotionStatisticsResponse {
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
	Notify *FinishTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a8iS4X94TgtgiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dq9hP8Sk2v6vQ6l05nCe5wiEiE
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// This parameter is required.
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
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	NoticeAllGroupMember       *bool     `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
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
	Attachments []*FinishTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 备注
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
	// example:
	//
	// wahaha.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// ticket/image/44708069/43003/e27204b382c04832aec4243e940a1367_1625831640499.txt
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *FinishTicketResponse) SetStatusCode(v int32) *FinishTicketResponse {
	s.StatusCode = &v
	return s
}

type GetAuthTokenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAuthTokenHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenHeaders) GoString() string {
	return s.String()
}

func (s *GetAuthTokenHeaders) SetCommonHeaders(v map[string]*string) *GetAuthTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAuthTokenHeaders) SetXAcsDingtalkAccessToken(v string) *GetAuthTokenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAuthTokenRequest struct {
	Channel       *string `json:"channel,omitempty" xml:"channel,omitempty"`
	EffectiveTime *int64  `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	OpenTeamId    *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	ServerId      *string `json:"serverId,omitempty" xml:"serverId,omitempty"`
	ServerName    *string `json:"serverName,omitempty" xml:"serverName,omitempty"`
}

func (s GetAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAuthTokenRequest) SetChannel(v string) *GetAuthTokenRequest {
	s.Channel = &v
	return s
}

func (s *GetAuthTokenRequest) SetEffectiveTime(v int64) *GetAuthTokenRequest {
	s.EffectiveTime = &v
	return s
}

func (s *GetAuthTokenRequest) SetOpenTeamId(v string) *GetAuthTokenRequest {
	s.OpenTeamId = &v
	return s
}

func (s *GetAuthTokenRequest) SetServerId(v string) *GetAuthTokenRequest {
	s.ServerId = &v
	return s
}

func (s *GetAuthTokenRequest) SetServerName(v string) *GetAuthTokenRequest {
	s.ServerName = &v
	return s
}

type GetAuthTokenResponseBody struct {
	DingOpenErrcode *int32                          `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string                         `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *GetAuthTokenResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success         *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBody) SetDingOpenErrcode(v int32) *GetAuthTokenResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetErrorMsg(v string) *GetAuthTokenResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetResult(v *GetAuthTokenResponseBodyResult) *GetAuthTokenResponseBody {
	s.Result = v
	return s
}

func (s *GetAuthTokenResponseBody) SetSuccess(v bool) *GetAuthTokenResponseBody {
	s.Success = &v
	return s
}

type GetAuthTokenResponseBodyResult struct {
	AuthToken     *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	Channel       *string `json:"channel,omitempty" xml:"channel,omitempty"`
	EffectiveTime *int64  `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	ServerId      *string `json:"serverId,omitempty" xml:"serverId,omitempty"`
	ServerName    *string `json:"serverName,omitempty" xml:"serverName,omitempty"`
}

func (s GetAuthTokenResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBodyResult) SetAuthToken(v string) *GetAuthTokenResponseBodyResult {
	s.AuthToken = &v
	return s
}

func (s *GetAuthTokenResponseBodyResult) SetChannel(v string) *GetAuthTokenResponseBodyResult {
	s.Channel = &v
	return s
}

func (s *GetAuthTokenResponseBodyResult) SetEffectiveTime(v int64) *GetAuthTokenResponseBodyResult {
	s.EffectiveTime = &v
	return s
}

func (s *GetAuthTokenResponseBodyResult) SetServerId(v string) *GetAuthTokenResponseBodyResult {
	s.ServerId = &v
	return s
}

func (s *GetAuthTokenResponseBodyResult) SetServerName(v string) *GetAuthTokenResponseBodyResult {
	s.ServerName = &v
	return s
}

type GetAuthTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponse) SetHeaders(v map[string]*string) *GetAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAuthTokenResponse) SetStatusCode(v int32) *GetAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthTokenResponse) SetBody(v *GetAuthTokenResponseBody) *GetAuthTokenResponse {
	s.Body = v
	return s
}

type GetInstancesByIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstancesByIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdsHeaders) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdsHeaders) SetCommonHeaders(v map[string]*string) *GetInstancesByIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstancesByIdsHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstancesByIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstancesByIdsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DING_CUSTOMER
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	OpenDataInstanceIdList []*string `json:"openDataInstanceIdList,omitempty" xml:"openDataInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 888***
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s GetInstancesByIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdsRequest) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdsRequest) SetFormCode(v string) *GetInstancesByIdsRequest {
	s.FormCode = &v
	return s
}

func (s *GetInstancesByIdsRequest) SetOpenDataInstanceIdList(v []*string) *GetInstancesByIdsRequest {
	s.OpenDataInstanceIdList = v
	return s
}

func (s *GetInstancesByIdsRequest) SetOpenTeamId(v string) *GetInstancesByIdsRequest {
	s.OpenTeamId = &v
	return s
}

type GetInstancesByIdsResponseBody struct {
	CustomFormInstanceResponseList []*GetInstancesByIdsResponseBodyCustomFormInstanceResponseList `json:"customFormInstanceResponseList,omitempty" xml:"customFormInstanceResponseList,omitempty" type:"Repeated"`
}

func (s GetInstancesByIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdsResponseBody) SetCustomFormInstanceResponseList(v []*GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) *GetInstancesByIdsResponseBody {
	s.CustomFormInstanceResponseList = v
	return s
}

type GetInstancesByIdsResponseBodyCustomFormInstanceResponseList struct {
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Fields         *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// This parameter is required.
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	GmtModified        *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	ModifiedUnionId    *string `json:"modifiedUnionId,omitempty" xml:"modifiedUnionId,omitempty"`
	OpenDataInstanceId *string `json:"openDataInstanceId,omitempty" xml:"openDataInstanceId,omitempty"`
	// This parameter is required.
	OpenTeamId   *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	OwnerUnionId *string `json:"ownerUnionId,omitempty" xml:"ownerUnionId,omitempty"`
}

func (s GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) SetCreatorUnionId(v string) *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList {
	s.CreatorUnionId = &v
	return s
}

func (s *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) SetFields(v string) *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList {
	s.Fields = &v
	return s
}

func (s *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) SetFormCode(v string) *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList {
	s.FormCode = &v
	return s
}

func (s *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) SetGmtCreate(v string) *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList {
	s.GmtCreate = &v
	return s
}

func (s *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) SetGmtModified(v string) *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList {
	s.GmtModified = &v
	return s
}

func (s *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) SetModifiedUnionId(v string) *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList {
	s.ModifiedUnionId = &v
	return s
}

func (s *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) SetOpenDataInstanceId(v string) *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList {
	s.OpenDataInstanceId = &v
	return s
}

func (s *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) SetOpenTeamId(v string) *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList {
	s.OpenTeamId = &v
	return s
}

func (s *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList) SetOwnerUnionId(v string) *GetInstancesByIdsResponseBodyCustomFormInstanceResponseList {
	s.OwnerUnionId = &v
	return s
}

type GetInstancesByIdsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstancesByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstancesByIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdsResponse) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdsResponse) SetHeaders(v map[string]*string) *GetInstancesByIdsResponse {
	s.Headers = v
	return s
}

func (s *GetInstancesByIdsResponse) SetStatusCode(v int32) *GetInstancesByIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstancesByIdsResponse) SetBody(v *GetInstancesByIdsResponseBody) *GetInstancesByIdsResponse {
	s.Body = v
	return s
}

type GetNegativeWordCloudHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetNegativeWordCloudHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetNegativeWordCloudHeaders) GoString() string {
	return s.String()
}

func (s *GetNegativeWordCloudHeaders) SetCommonHeaders(v map[string]*string) *GetNegativeWordCloudHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNegativeWordCloudHeaders) SetXAcsDingtalkAccessToken(v string) *GetNegativeWordCloudHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetNegativeWordCloudRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// KxisoOk
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s GetNegativeWordCloudRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNegativeWordCloudRequest) GoString() string {
	return s.String()
}

func (s *GetNegativeWordCloudRequest) SetOpenTeamId(v string) *GetNegativeWordCloudRequest {
	s.OpenTeamId = &v
	return s
}

type GetNegativeWordCloudResponseBody struct {
	// This parameter is required.
	Words []*GetNegativeWordCloudResponseBodyWords `json:"words,omitempty" xml:"words,omitempty" type:"Repeated"`
}

func (s GetNegativeWordCloudResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNegativeWordCloudResponseBody) GoString() string {
	return s.String()
}

func (s *GetNegativeWordCloudResponseBody) SetWords(v []*GetNegativeWordCloudResponseBodyWords) *GetNegativeWordCloudResponseBody {
	s.Words = v
	return s
}

type GetNegativeWordCloudResponseBodyWords struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 销售
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s GetNegativeWordCloudResponseBodyWords) String() string {
	return tea.Prettify(s)
}

func (s GetNegativeWordCloudResponseBodyWords) GoString() string {
	return s.String()
}

func (s *GetNegativeWordCloudResponseBodyWords) SetCount(v int64) *GetNegativeWordCloudResponseBodyWords {
	s.Count = &v
	return s
}

func (s *GetNegativeWordCloudResponseBodyWords) SetWord(v string) *GetNegativeWordCloudResponseBodyWords {
	s.Word = &v
	return s
}

type GetNegativeWordCloudResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNegativeWordCloudResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNegativeWordCloudResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNegativeWordCloudResponse) GoString() string {
	return s.String()
}

func (s *GetNegativeWordCloudResponse) SetHeaders(v map[string]*string) *GetNegativeWordCloudResponse {
	s.Headers = v
	return s
}

func (s *GetNegativeWordCloudResponse) SetStatusCode(v int32) *GetNegativeWordCloudResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNegativeWordCloudResponse) SetBody(v *GetNegativeWordCloudResponseBody) *GetNegativeWordCloudResponse {
	s.Body = v
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
	// This parameter is required.
	FetchMode *string `json:"fetchMode,omitempty" xml:"fetchMode,omitempty"`
	// This parameter is required.
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssTempUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetOssTempUrlResponse) SetStatusCode(v int32) *GetOssTempUrlResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// TICKET_IMAGE
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// wahah.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// This parameter is required.
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
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Policy      *string `json:"policy,omitempty" xml:"policy,omitempty"`
	Signature   *string `json:"signature,omitempty" xml:"signature,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStoragePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetStoragePolicyResponse) SetStatusCode(v int32) *GetStoragePolicyResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
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
	CreateTime         *string                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator            *GetTicketResponseBodyCreator   `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	CustomFields       *string                         `json:"customFields,omitempty" xml:"customFields,omitempty"`
	OpenConversationId *string                         `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenTicketId       *string                         `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	Processor          *GetTicketResponseBodyProcessor `json:"processor,omitempty" xml:"processor,omitempty" type:"Struct"`
	Scene              *string                         `json:"scene,omitempty" xml:"scene,omitempty"`
	SceneContext       *string                         `json:"sceneContext,omitempty" xml:"sceneContext,omitempty"`
	Stage              *string                         `json:"stage,omitempty" xml:"stage,omitempty"`
	Takers             []*GetTicketResponseBodyTakers  `json:"takers,omitempty" xml:"takers,omitempty" type:"Repeated"`
	Template           *GetTicketResponseBodyTemplate  `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	Title              *string                         `json:"title,omitempty" xml:"title,omitempty"`
	UpdateTime         *string                         `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
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
	OpenTemplateBizId *string `json:"openTemplateBizId,omitempty" xml:"openTemplateBizId,omitempty"`
	OpenTemplateId    *string `json:"openTemplateId,omitempty" xml:"openTemplateId,omitempty"`
	TemplateName      *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetTicketResponse) SetStatusCode(v int32) *GetTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTicketResponse) SetBody(v *GetTicketResponseBody) *GetTicketResponse {
	s.Body = v
	return s
}

type GetWordCloudHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetWordCloudHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWordCloudHeaders) GoString() string {
	return s.String()
}

func (s *GetWordCloudHeaders) SetCommonHeaders(v map[string]*string) *GetWordCloudHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWordCloudHeaders) SetXAcsDingtalkAccessToken(v string) *GetWordCloudHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetWordCloudRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// KxisoOk
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s GetWordCloudRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWordCloudRequest) GoString() string {
	return s.String()
}

func (s *GetWordCloudRequest) SetOpenTeamId(v string) *GetWordCloudRequest {
	s.OpenTeamId = &v
	return s
}

type GetWordCloudResponseBody struct {
	// This parameter is required.
	Words []*GetWordCloudResponseBodyWords `json:"words,omitempty" xml:"words,omitempty" type:"Repeated"`
}

func (s GetWordCloudResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWordCloudResponseBody) GoString() string {
	return s.String()
}

func (s *GetWordCloudResponseBody) SetWords(v []*GetWordCloudResponseBodyWords) *GetWordCloudResponseBody {
	s.Words = v
	return s
}

type GetWordCloudResponseBodyWords struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 销售
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s GetWordCloudResponseBodyWords) String() string {
	return tea.Prettify(s)
}

func (s GetWordCloudResponseBodyWords) GoString() string {
	return s.String()
}

func (s *GetWordCloudResponseBodyWords) SetCount(v int64) *GetWordCloudResponseBodyWords {
	s.Count = &v
	return s
}

func (s *GetWordCloudResponseBodyWords) SetWord(v string) *GetWordCloudResponseBodyWords {
	s.Word = &v
	return s
}

type GetWordCloudResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWordCloudResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWordCloudResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWordCloudResponse) GoString() string {
	return s.String()
}

func (s *GetWordCloudResponse) SetHeaders(v map[string]*string) *GetWordCloudResponse {
	s.Headers = v
	return s
}

func (s *GetWordCloudResponse) SetStatusCode(v int32) *GetWordCloudResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWordCloudResponse) SetBody(v *GetWordCloudResponseBody) *GetWordCloudResponse {
	s.Body = v
	return s
}

type GroupStatisticsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupStatisticsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupStatisticsHeaders) GoString() string {
	return s.String()
}

func (s *GroupStatisticsHeaders) SetCommonHeaders(v map[string]*string) *GroupStatisticsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupStatisticsHeaders) SetXAcsDingtalkAccessToken(v string) *GroupStatisticsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MaxDt *string `json:"maxDt,omitempty" xml:"maxDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MinDt *string `json:"minDt,omitempty" xml:"minDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KxisoOk
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s GroupStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GroupStatisticsRequest) SetMaxDt(v string) *GroupStatisticsRequest {
	s.MaxDt = &v
	return s
}

func (s *GroupStatisticsRequest) SetMinDt(v string) *GroupStatisticsRequest {
	s.MinDt = &v
	return s
}

func (s *GroupStatisticsRequest) SetOpenTeamId(v string) *GroupStatisticsRequest {
	s.OpenTeamId = &v
	return s
}

type GroupStatisticsResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	GroupCount *int64 `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	// This parameter is required.
	GroupTrend []*GroupStatisticsResponseBodyGroupTrend `json:"groupTrend,omitempty" xml:"groupTrend,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	IncreaseGroupCount *int64 `json:"increaseGroupCount,omitempty" xml:"increaseGroupCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.1
	IncreaseRate *string `json:"increaseRate,omitempty" xml:"increaseRate,omitempty"`
}

func (s GroupStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GroupStatisticsResponseBody) SetGroupCount(v int64) *GroupStatisticsResponseBody {
	s.GroupCount = &v
	return s
}

func (s *GroupStatisticsResponseBody) SetGroupTrend(v []*GroupStatisticsResponseBodyGroupTrend) *GroupStatisticsResponseBody {
	s.GroupTrend = v
	return s
}

func (s *GroupStatisticsResponseBody) SetIncreaseGroupCount(v int64) *GroupStatisticsResponseBody {
	s.IncreaseGroupCount = &v
	return s
}

func (s *GroupStatisticsResponseBody) SetIncreaseRate(v string) *GroupStatisticsResponseBody {
	s.IncreaseRate = &v
	return s
}

type GroupStatisticsResponseBodyGroupTrend struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	Dt *string `json:"dt,omitempty" xml:"dt,omitempty"`
}

func (s GroupStatisticsResponseBodyGroupTrend) String() string {
	return tea.Prettify(s)
}

func (s GroupStatisticsResponseBodyGroupTrend) GoString() string {
	return s.String()
}

func (s *GroupStatisticsResponseBodyGroupTrend) SetCount(v int64) *GroupStatisticsResponseBodyGroupTrend {
	s.Count = &v
	return s
}

func (s *GroupStatisticsResponseBodyGroupTrend) SetDt(v string) *GroupStatisticsResponseBodyGroupTrend {
	s.Dt = &v
	return s
}

type GroupStatisticsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GroupStatisticsResponse) SetHeaders(v map[string]*string) *GroupStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GroupStatisticsResponse) SetStatusCode(v int32) *GroupStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupStatisticsResponse) SetBody(v *GroupStatisticsResponseBody) *GroupStatisticsResponse {
	s.Body = v
	return s
}

type IntentionCategoryStatisticsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IntentionCategoryStatisticsHeaders) String() string {
	return tea.Prettify(s)
}

func (s IntentionCategoryStatisticsHeaders) GoString() string {
	return s.String()
}

func (s *IntentionCategoryStatisticsHeaders) SetCommonHeaders(v map[string]*string) *IntentionCategoryStatisticsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntentionCategoryStatisticsHeaders) SetXAcsDingtalkAccessToken(v string) *IntentionCategoryStatisticsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IntentionCategoryStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MaxDt *string `json:"maxDt,omitempty" xml:"maxDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MinDt *string `json:"minDt,omitempty" xml:"minDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KxisoOk
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s IntentionCategoryStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s IntentionCategoryStatisticsRequest) GoString() string {
	return s.String()
}

func (s *IntentionCategoryStatisticsRequest) SetMaxDt(v string) *IntentionCategoryStatisticsRequest {
	s.MaxDt = &v
	return s
}

func (s *IntentionCategoryStatisticsRequest) SetMinDt(v string) *IntentionCategoryStatisticsRequest {
	s.MinDt = &v
	return s
}

func (s *IntentionCategoryStatisticsRequest) SetOpenTeamId(v string) *IntentionCategoryStatisticsRequest {
	s.OpenTeamId = &v
	return s
}

type IntentionCategoryStatisticsResponseBody struct {
	// This parameter is required.
	IntentionCategoryRecords []*IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords `json:"intentionCategoryRecords,omitempty" xml:"intentionCategoryRecords,omitempty" type:"Repeated"`
}

func (s IntentionCategoryStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IntentionCategoryStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *IntentionCategoryStatisticsResponseBody) SetIntentionCategoryRecords(v []*IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords) *IntentionCategoryStatisticsResponseBody {
	s.IntentionCategoryRecords = v
	return s
}

type IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	AskCount *int64 `json:"askCount,omitempty" xml:"askCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 工单类
	CategoryName *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	DissatisfiedCount *int64 `json:"dissatisfiedCount,omitempty" xml:"dissatisfiedCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	ErrorCount *int64 `json:"errorCount,omitempty" xml:"errorCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PraiseCount *int64 `json:"praiseCount,omitempty" xml:"praiseCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	SuggestCount *int64 `json:"suggestCount,omitempty" xml:"suggestCount,omitempty"`
}

func (s IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords) String() string {
	return tea.Prettify(s)
}

func (s IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords) GoString() string {
	return s.String()
}

func (s *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords) SetAskCount(v int64) *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords {
	s.AskCount = &v
	return s
}

func (s *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords) SetCategoryName(v string) *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords {
	s.CategoryName = &v
	return s
}

func (s *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords) SetDissatisfiedCount(v int64) *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords {
	s.DissatisfiedCount = &v
	return s
}

func (s *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords) SetErrorCount(v int64) *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords {
	s.ErrorCount = &v
	return s
}

func (s *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords) SetPraiseCount(v int64) *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords {
	s.PraiseCount = &v
	return s
}

func (s *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords) SetSuggestCount(v int64) *IntentionCategoryStatisticsResponseBodyIntentionCategoryRecords {
	s.SuggestCount = &v
	return s
}

type IntentionCategoryStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntentionCategoryStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntentionCategoryStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s IntentionCategoryStatisticsResponse) GoString() string {
	return s.String()
}

func (s *IntentionCategoryStatisticsResponse) SetHeaders(v map[string]*string) *IntentionCategoryStatisticsResponse {
	s.Headers = v
	return s
}

func (s *IntentionCategoryStatisticsResponse) SetStatusCode(v int32) *IntentionCategoryStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *IntentionCategoryStatisticsResponse) SetBody(v *IntentionCategoryStatisticsResponseBody) *IntentionCategoryStatisticsResponse {
	s.Body = v
	return s
}

type IntentionStatisticsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IntentionStatisticsHeaders) String() string {
	return tea.Prettify(s)
}

func (s IntentionStatisticsHeaders) GoString() string {
	return s.String()
}

func (s *IntentionStatisticsHeaders) SetCommonHeaders(v map[string]*string) *IntentionStatisticsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntentionStatisticsHeaders) SetXAcsDingtalkAccessToken(v string) *IntentionStatisticsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IntentionStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MaxDt *string `json:"maxDt,omitempty" xml:"maxDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MinDt *string `json:"minDt,omitempty" xml:"minDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KxisoOk
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s IntentionStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s IntentionStatisticsRequest) GoString() string {
	return s.String()
}

func (s *IntentionStatisticsRequest) SetMaxDt(v string) *IntentionStatisticsRequest {
	s.MaxDt = &v
	return s
}

func (s *IntentionStatisticsRequest) SetMinDt(v string) *IntentionStatisticsRequest {
	s.MinDt = &v
	return s
}

func (s *IntentionStatisticsRequest) SetOpenTeamId(v string) *IntentionStatisticsRequest {
	s.OpenTeamId = &v
	return s
}

type IntentionStatisticsResponseBody struct {
	// This parameter is required.
	IntentionStatisticsRecords []*IntentionStatisticsResponseBodyIntentionStatisticsRecords `json:"intentionStatisticsRecords,omitempty" xml:"intentionStatisticsRecords,omitempty" type:"Repeated"`
	// This parameter is required.
	IntentionTrend []*IntentionStatisticsResponseBodyIntentionTrend `json:"intentionTrend,omitempty" xml:"intentionTrend,omitempty" type:"Repeated"`
}

func (s IntentionStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IntentionStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *IntentionStatisticsResponseBody) SetIntentionStatisticsRecords(v []*IntentionStatisticsResponseBodyIntentionStatisticsRecords) *IntentionStatisticsResponseBody {
	s.IntentionStatisticsRecords = v
	return s
}

func (s *IntentionStatisticsResponseBody) SetIntentionTrend(v []*IntentionStatisticsResponseBodyIntentionTrend) *IntentionStatisticsResponseBody {
	s.IntentionTrend = v
	return s
}

type IntentionStatisticsResponseBodyIntentionStatisticsRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 产品异常类
	Intention *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9
	LastCount *int64 `json:"lastCount,omitempty" xml:"lastCount,omitempty"`
}

func (s IntentionStatisticsResponseBodyIntentionStatisticsRecords) String() string {
	return tea.Prettify(s)
}

func (s IntentionStatisticsResponseBodyIntentionStatisticsRecords) GoString() string {
	return s.String()
}

func (s *IntentionStatisticsResponseBodyIntentionStatisticsRecords) SetCount(v int64) *IntentionStatisticsResponseBodyIntentionStatisticsRecords {
	s.Count = &v
	return s
}

func (s *IntentionStatisticsResponseBodyIntentionStatisticsRecords) SetIntention(v string) *IntentionStatisticsResponseBodyIntentionStatisticsRecords {
	s.Intention = &v
	return s
}

func (s *IntentionStatisticsResponseBodyIntentionStatisticsRecords) SetLastCount(v int64) *IntentionStatisticsResponseBodyIntentionStatisticsRecords {
	s.LastCount = &v
	return s
}

type IntentionStatisticsResponseBodyIntentionTrend struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	Dt *string `json:"dt,omitempty" xml:"dt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 产品异常类
	Intention *string `json:"intention,omitempty" xml:"intention,omitempty"`
}

func (s IntentionStatisticsResponseBodyIntentionTrend) String() string {
	return tea.Prettify(s)
}

func (s IntentionStatisticsResponseBodyIntentionTrend) GoString() string {
	return s.String()
}

func (s *IntentionStatisticsResponseBodyIntentionTrend) SetCount(v int64) *IntentionStatisticsResponseBodyIntentionTrend {
	s.Count = &v
	return s
}

func (s *IntentionStatisticsResponseBodyIntentionTrend) SetDt(v string) *IntentionStatisticsResponseBodyIntentionTrend {
	s.Dt = &v
	return s
}

func (s *IntentionStatisticsResponseBodyIntentionTrend) SetIntention(v string) *IntentionStatisticsResponseBodyIntentionTrend {
	s.Intention = &v
	return s
}

type IntentionStatisticsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntentionStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntentionStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s IntentionStatisticsResponse) GoString() string {
	return s.String()
}

func (s *IntentionStatisticsResponse) SetHeaders(v map[string]*string) *IntentionStatisticsResponse {
	s.Headers = v
	return s
}

func (s *IntentionStatisticsResponse) SetStatusCode(v int32) *IntentionStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *IntentionStatisticsResponse) SetBody(v *IntentionStatisticsResponseBody) *IntentionStatisticsResponse {
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
	// This parameter is required.
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
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
	OpenTicketId         *string                                               `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	OperateData          *string                                               `json:"operateData,omitempty" xml:"operateData,omitempty"`
	OperateTime          *string                                               `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	Operation            *string                                               `json:"operation,omitempty" xml:"operation,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTicketOperateRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTicketOperateRecordResponse) SetStatusCode(v int32) *ListTicketOperateRecordResponse {
	s.StatusCode = &v
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
	// example:
	//
	// Jxi12wo3qxoa
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 测试团队
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListUserTeamsResponse) SetStatusCode(v int32) *ListUserTeamsResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxxx==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// KxisoOk
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 5
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
	// This parameter is required.
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
	// This parameter is required.
	ActionIndexL14d *float64 `json:"actionIndexL14d,omitempty" xml:"actionIndexL14d,omitempty"`
	// This parameter is required.
	ActionIndexL30d *float64 `json:"actionIndexL30d,omitempty" xml:"actionIndexL30d,omitempty"`
	// This parameter is required.
	ActionIndexL7d *float64 `json:"actionIndexL7d,omitempty" xml:"actionIndexL7d,omitempty"`
	// This parameter is required.
	ActiveScore *float64 `json:"activeScore,omitempty" xml:"activeScore,omitempty"`
	// This parameter is required.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// This parameter is required.
	Ranking *int64 `json:"ranking,omitempty" xml:"ranking,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryActiveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryActiveUsersResponse) SetStatusCode(v int32) *QueryActiveUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryActiveUsersResponse) SetBody(v *QueryActiveUsersResponseBody) *QueryActiveUsersResponse {
	s.Body = v
	return s
}

type QueryCrmGroupContactHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCrmGroupContactHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmGroupContactHeaders) GoString() string {
	return s.String()
}

func (s *QueryCrmGroupContactHeaders) SetCommonHeaders(v map[string]*string) *QueryCrmGroupContactHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCrmGroupContactHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCrmGroupContactHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCrmGroupContactRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	MinResult *int64 `json:"minResult,omitempty" xml:"minResult,omitempty"`
	// example:
	//
	// 8888
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid888
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 888
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	SearchFields *string `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s QueryCrmGroupContactRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmGroupContactRequest) GoString() string {
	return s.String()
}

func (s *QueryCrmGroupContactRequest) SetMinResult(v int64) *QueryCrmGroupContactRequest {
	s.MinResult = &v
	return s
}

func (s *QueryCrmGroupContactRequest) SetNextToken(v string) *QueryCrmGroupContactRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCrmGroupContactRequest) SetOpenConversationId(v string) *QueryCrmGroupContactRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryCrmGroupContactRequest) SetOpenTeamId(v string) *QueryCrmGroupContactRequest {
	s.OpenTeamId = &v
	return s
}

func (s *QueryCrmGroupContactRequest) SetSearchFields(v string) *QueryCrmGroupContactRequest {
	s.SearchFields = &v
	return s
}

type QueryCrmGroupContactResponseBody struct {
	// example:
	//
	// token****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// cid****
	OpenConversationId *string                                    `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Records            []*QueryCrmGroupContactResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s QueryCrmGroupContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmGroupContactResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCrmGroupContactResponseBody) SetNextToken(v string) *QueryCrmGroupContactResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCrmGroupContactResponseBody) SetOpenConversationId(v string) *QueryCrmGroupContactResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *QueryCrmGroupContactResponseBody) SetRecords(v []*QueryCrmGroupContactResponseBodyRecords) *QueryCrmGroupContactResponseBody {
	s.Records = v
	return s
}

type QueryCrmGroupContactResponseBodyRecords struct {
	// example:
	//
	// {} ,具体字段取决于客户管理-字段管理-联系人字段设置
	ContactData *string `json:"contactData,omitempty" xml:"contactData,omitempty"`
	// example:
	//
	// ahghgg
	MemberUnionId *string `json:"memberUnionId,omitempty" xml:"memberUnionId,omitempty"`
	// example:
	//
	// 张三
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// 88888
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryCrmGroupContactResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmGroupContactResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *QueryCrmGroupContactResponseBodyRecords) SetContactData(v string) *QueryCrmGroupContactResponseBodyRecords {
	s.ContactData = &v
	return s
}

func (s *QueryCrmGroupContactResponseBodyRecords) SetMemberUnionId(v string) *QueryCrmGroupContactResponseBodyRecords {
	s.MemberUnionId = &v
	return s
}

func (s *QueryCrmGroupContactResponseBodyRecords) SetNickName(v string) *QueryCrmGroupContactResponseBodyRecords {
	s.NickName = &v
	return s
}

func (s *QueryCrmGroupContactResponseBodyRecords) SetUserId(v string) *QueryCrmGroupContactResponseBodyRecords {
	s.UserId = &v
	return s
}

type QueryCrmGroupContactResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCrmGroupContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCrmGroupContactResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmGroupContactResponse) GoString() string {
	return s.String()
}

func (s *QueryCrmGroupContactResponse) SetHeaders(v map[string]*string) *QueryCrmGroupContactResponse {
	s.Headers = v
	return s
}

func (s *QueryCrmGroupContactResponse) SetStatusCode(v int32) *QueryCrmGroupContactResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCrmGroupContactResponse) SetBody(v *QueryCrmGroupContactResponseBody) *QueryCrmGroupContactResponse {
	s.Body = v
	return s
}

type QueryCustomerCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCustomerCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerCardHeaders) GoString() string {
	return s.String()
}

func (s *QueryCustomerCardHeaders) SetCommonHeaders(v map[string]*string) *QueryCustomerCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCustomerCardHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCustomerCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCustomerCardRequest struct {
	JsonParams *string `json:"jsonParams,omitempty" xml:"jsonParams,omitempty"`
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s QueryCustomerCardRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerCardRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomerCardRequest) SetJsonParams(v string) *QueryCustomerCardRequest {
	s.JsonParams = &v
	return s
}

func (s *QueryCustomerCardRequest) SetOpenTeamId(v string) *QueryCustomerCardRequest {
	s.OpenTeamId = &v
	return s
}

type QueryCustomerCardResponseBody struct {
	DingOpenErrcode *int32  `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success         *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryCustomerCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerCardResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomerCardResponseBody) SetDingOpenErrcode(v int32) *QueryCustomerCardResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *QueryCustomerCardResponseBody) SetErrorMsg(v string) *QueryCustomerCardResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryCustomerCardResponseBody) SetResult(v bool) *QueryCustomerCardResponseBody {
	s.Result = &v
	return s
}

func (s *QueryCustomerCardResponseBody) SetSuccess(v bool) *QueryCustomerCardResponseBody {
	s.Success = &v
	return s
}

type QueryCustomerCardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomerCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomerCardResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerCardResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomerCardResponse) SetHeaders(v map[string]*string) *QueryCustomerCardResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomerCardResponse) SetStatusCode(v int32) *QueryCustomerCardResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomerCardResponse) SetBody(v *QueryCustomerCardResponseBody) *QueryCustomerCardResponse {
	s.Body = v
	return s
}

type QueryCustomerTaskUserDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCustomerTaskUserDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerTaskUserDetailHeaders) GoString() string {
	return s.String()
}

func (s *QueryCustomerTaskUserDetailHeaders) SetCommonHeaders(v map[string]*string) *QueryCustomerTaskUserDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCustomerTaskUserDetailHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCustomerTaskUserDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCustomerTaskUserDetailRequest struct {
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 8888
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8888
	OpenBatchTaskId *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8888
	OpenTeamId       *string   `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	ReceiverUnionIds []*string `json:"receiverUnionIds,omitempty" xml:"receiverUnionIds,omitempty" type:"Repeated"`
}

func (s QueryCustomerTaskUserDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerTaskUserDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomerTaskUserDetailRequest) SetMaxResults(v int64) *QueryCustomerTaskUserDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCustomerTaskUserDetailRequest) SetNextToken(v string) *QueryCustomerTaskUserDetailRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCustomerTaskUserDetailRequest) SetOpenBatchTaskId(v string) *QueryCustomerTaskUserDetailRequest {
	s.OpenBatchTaskId = &v
	return s
}

func (s *QueryCustomerTaskUserDetailRequest) SetOpenTeamId(v string) *QueryCustomerTaskUserDetailRequest {
	s.OpenTeamId = &v
	return s
}

func (s *QueryCustomerTaskUserDetailRequest) SetReceiverUnionIds(v []*string) *QueryCustomerTaskUserDetailRequest {
	s.ReceiverUnionIds = v
	return s
}

type QueryCustomerTaskUserDetailResponseBody struct {
	// This parameter is required.
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken  *string                                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Records    []*QueryCustomerTaskUserDetailResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	TotalCount *int64                                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryCustomerTaskUserDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerTaskUserDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomerTaskUserDetailResponseBody) SetMaxResults(v int64) *QueryCustomerTaskUserDetailResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBody) SetNextToken(v string) *QueryCustomerTaskUserDetailResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBody) SetRecords(v []*QueryCustomerTaskUserDetailResponseBodyRecords) *QueryCustomerTaskUserDetailResponseBody {
	s.Records = v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBody) SetTotalCount(v int64) *QueryCustomerTaskUserDetailResponseBody {
	s.TotalCount = &v
	return s
}

type QueryCustomerTaskUserDetailResponseBodyRecords struct {
	// example:
	//
	// 客户名称
	CustomerNames *string `json:"customerNames,omitempty" xml:"customerNames,omitempty"`
	// example:
	//
	// 11111
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrorDetail         *string                                                              `json:"errorDetail,omitempty" xml:"errorDetail,omitempty"`
	EventTrackResponses []*QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses `json:"eventTrackResponses,omitempty" xml:"eventTrackResponses,omitempty" type:"Repeated"`
	// example:
	//
	// 8888
	OpenBatchTaskId *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
	// example:
	//
	// 1
	ReadStatus *int64 `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
	// example:
	//
	// 2023-07-14 00:00:00
	ReadTime *string `json:"readTime,omitempty" xml:"readTime,omitempty"`
	// example:
	//
	// 接收人姓名
	ReceiverName *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	// example:
	//
	// 接收人ID
	ReceiverUnionId *string `json:"receiverUnionId,omitempty" xml:"receiverUnionId,omitempty"`
	// example:
	//
	// 2023-07-14 00:00:00
	SendTime *string `json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	// example:
	//
	// UNSEND未发；SEND_SUCCESS成功；SEND_FAILED失败；EXCEED_LIMIT被限流
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryCustomerTaskUserDetailResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerTaskUserDetailResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetCustomerNames(v string) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.CustomerNames = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetErrorCode(v string) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.ErrorCode = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetErrorDetail(v string) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.ErrorDetail = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetEventTrackResponses(v []*QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.EventTrackResponses = v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetOpenBatchTaskId(v string) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.OpenBatchTaskId = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetReadStatus(v int64) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.ReadStatus = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetReadTime(v string) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.ReadTime = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetReceiverName(v string) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.ReceiverName = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetReceiverUnionId(v string) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.ReceiverUnionId = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetSendTime(v string) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.SendTime = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecords) SetStatus(v string) *QueryCustomerTaskUserDetailResponseBodyRecords {
	s.Status = &v
	return s
}

type QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses struct {
	// example:
	//
	// 2023-07-14 00:00:00
	ClickTime *string `json:"clickTime,omitempty" xml:"clickTime,omitempty"`
	// example:
	//
	// 88888
	EventTrackId *string `json:"eventTrackId,omitempty" xml:"eventTrackId,omitempty"`
	// example:
	//
	// true
	OnClick *bool `json:"onClick,omitempty" xml:"onClick,omitempty"`
	// example:
	//
	// 标题名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses) GoString() string {
	return s.String()
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses) SetClickTime(v string) *QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses {
	s.ClickTime = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses) SetEventTrackId(v string) *QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses {
	s.EventTrackId = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses) SetOnClick(v bool) *QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses {
	s.OnClick = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses) SetTitle(v string) *QueryCustomerTaskUserDetailResponseBodyRecordsEventTrackResponses {
	s.Title = &v
	return s
}

type QueryCustomerTaskUserDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomerTaskUserDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomerTaskUserDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerTaskUserDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomerTaskUserDetailResponse) SetHeaders(v map[string]*string) *QueryCustomerTaskUserDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomerTaskUserDetailResponse) SetStatusCode(v int32) *QueryCustomerTaskUserDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomerTaskUserDetailResponse) SetBody(v *QueryCustomerTaskUserDetailResponseBody) *QueryCustomerTaskUserDetailResponse {
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
	// example:
	//
	// 101813123123
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// example:
	//
	// cidxxxxxx==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// KxisoOk
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
	// This parameter is required.
	//
	// example:
	//
	// 234
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉钉专属服务群
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://qr.dingtalk.com/xxxxxxx
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxx==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xjfjdsiw
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xkjhfker
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// jikwrjcowa
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 服务小钉
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryGroupResponse) SetStatusCode(v int32) *QueryGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupResponse) SetBody(v *QueryGroupResponseBody) *QueryGroupResponse {
	s.Body = v
	return s
}

type QueryGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupMemberRequest struct {
	// example:
	//
	// cidxxxxxx==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// KxisoOk
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s QueryGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberRequest) SetOpenConversationId(v string) *QueryGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryGroupMemberRequest) SetOpenTeamId(v string) *QueryGroupMemberRequest {
	s.OpenTeamId = &v
	return s
}

func (s *QueryGroupMemberRequest) SetTargetCorpId(v string) *QueryGroupMemberRequest {
	s.TargetCorpId = &v
	return s
}

type QueryGroupMemberResponseBody struct {
	Result  *QueryGroupMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberResponseBody) SetResult(v *QueryGroupMemberResponseBodyResult) *QueryGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *QueryGroupMemberResponseBody) SetSuccess(v bool) *QueryGroupMemberResponseBody {
	s.Success = &v
	return s
}

type QueryGroupMemberResponseBodyResult struct {
	GroupMemberList    []*QueryGroupMemberResponseBodyResultGroupMemberList `json:"groupMemberList,omitempty" xml:"groupMemberList,omitempty" type:"Repeated"`
	OpenConversationId *string                                              `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryGroupMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberResponseBodyResult) SetGroupMemberList(v []*QueryGroupMemberResponseBodyResultGroupMemberList) *QueryGroupMemberResponseBodyResult {
	s.GroupMemberList = v
	return s
}

func (s *QueryGroupMemberResponseBodyResult) SetOpenConversationId(v string) *QueryGroupMemberResponseBodyResult {
	s.OpenConversationId = &v
	return s
}

type QueryGroupMemberResponseBodyResultGroupMemberList struct {
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	IsUser        *bool   `json:"isUser,omitempty" xml:"isUser,omitempty"`
	NickName      *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Owner         *bool   `json:"owner,omitempty" xml:"owner,omitempty"`
	UnionId       *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryGroupMemberResponseBodyResultGroupMemberList) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberResponseBodyResultGroupMemberList) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberResponseBodyResultGroupMemberList) SetAvatarMediaId(v string) *QueryGroupMemberResponseBodyResultGroupMemberList {
	s.AvatarMediaId = &v
	return s
}

func (s *QueryGroupMemberResponseBodyResultGroupMemberList) SetIsUser(v bool) *QueryGroupMemberResponseBodyResultGroupMemberList {
	s.IsUser = &v
	return s
}

func (s *QueryGroupMemberResponseBodyResultGroupMemberList) SetNickName(v string) *QueryGroupMemberResponseBodyResultGroupMemberList {
	s.NickName = &v
	return s
}

func (s *QueryGroupMemberResponseBodyResultGroupMemberList) SetOwner(v bool) *QueryGroupMemberResponseBodyResultGroupMemberList {
	s.Owner = &v
	return s
}

func (s *QueryGroupMemberResponseBodyResultGroupMemberList) SetUnionId(v string) *QueryGroupMemberResponseBodyResultGroupMemberList {
	s.UnionId = &v
	return s
}

func (s *QueryGroupMemberResponseBodyResultGroupMemberList) SetUserId(v string) *QueryGroupMemberResponseBodyResultGroupMemberList {
	s.UserId = &v
	return s
}

type QueryGroupMemberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberResponse) SetHeaders(v map[string]*string) *QueryGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupMemberResponse) SetStatusCode(v int32) *QueryGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupMemberResponse) SetBody(v *QueryGroupMemberResponseBody) *QueryGroupMemberResponse {
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
	// This parameter is required.
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
	// This parameter is required.
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
	// This parameter is required.
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	GroupSetName *string `json:"groupSetName,omitempty" xml:"groupSetName,omitempty"`
	// This parameter is required.
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryGroupSetResponse) SetStatusCode(v int32) *QueryGroupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupSetResponse) SetBody(v *QueryGroupSetResponseBody) *QueryGroupSetResponse {
	s.Body = v
	return s
}

type QueryInstancesByMultiConditionsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryInstancesByMultiConditionsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancesByMultiConditionsHeaders) GoString() string {
	return s.String()
}

func (s *QueryInstancesByMultiConditionsHeaders) SetCommonHeaders(v map[string]*string) *QueryInstancesByMultiConditionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryInstancesByMultiConditionsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryInstancesByMultiConditionsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryInstancesByMultiConditionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DING_CUSTOMER
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 888**
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// [     {         "fieldCode":"contact_name",         "fieldOperatorType":"like",         "value":"测试api"     } ]
	SearchFields *string                                             `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
	SortFields   []*QueryInstancesByMultiConditionsRequestSortFields `json:"sortFields,omitempty" xml:"sortFields,omitempty" type:"Repeated"`
}

func (s QueryInstancesByMultiConditionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancesByMultiConditionsRequest) GoString() string {
	return s.String()
}

func (s *QueryInstancesByMultiConditionsRequest) SetFormCode(v string) *QueryInstancesByMultiConditionsRequest {
	s.FormCode = &v
	return s
}

func (s *QueryInstancesByMultiConditionsRequest) SetMaxResults(v int64) *QueryInstancesByMultiConditionsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryInstancesByMultiConditionsRequest) SetNextToken(v string) *QueryInstancesByMultiConditionsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryInstancesByMultiConditionsRequest) SetOpenTeamId(v string) *QueryInstancesByMultiConditionsRequest {
	s.OpenTeamId = &v
	return s
}

func (s *QueryInstancesByMultiConditionsRequest) SetSearchFields(v string) *QueryInstancesByMultiConditionsRequest {
	s.SearchFields = &v
	return s
}

func (s *QueryInstancesByMultiConditionsRequest) SetSortFields(v []*QueryInstancesByMultiConditionsRequestSortFields) *QueryInstancesByMultiConditionsRequest {
	s.SortFields = v
	return s
}

type QueryInstancesByMultiConditionsRequestSortFields struct {
	// example:
	//
	// gmt_create
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// example:
	//
	// asc升序；desc降序
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s QueryInstancesByMultiConditionsRequestSortFields) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancesByMultiConditionsRequestSortFields) GoString() string {
	return s.String()
}

func (s *QueryInstancesByMultiConditionsRequestSortFields) SetFieldCode(v string) *QueryInstancesByMultiConditionsRequestSortFields {
	s.FieldCode = &v
	return s
}

func (s *QueryInstancesByMultiConditionsRequestSortFields) SetSortBy(v string) *QueryInstancesByMultiConditionsRequestSortFields {
	s.SortBy = &v
	return s
}

type QueryInstancesByMultiConditionsResponseBody struct {
	// This parameter is required.
	MaxResults *int64                                                `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Records    []*QueryInstancesByMultiConditionsResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryInstancesByMultiConditionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancesByMultiConditionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstancesByMultiConditionsResponseBody) SetMaxResults(v int64) *QueryInstancesByMultiConditionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBody) SetNextToken(v string) *QueryInstancesByMultiConditionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBody) SetRecords(v []*QueryInstancesByMultiConditionsResponseBodyRecords) *QueryInstancesByMultiConditionsResponseBody {
	s.Records = v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBody) SetTotalCount(v int64) *QueryInstancesByMultiConditionsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryInstancesByMultiConditionsResponseBodyRecords struct {
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Fields         *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// This parameter is required.
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	GmtModified     *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	ModifiedUnionId *string `json:"modifiedUnionId,omitempty" xml:"modifiedUnionId,omitempty"`
	// This parameter is required.
	OpenDataInstanceId *string `json:"openDataInstanceId,omitempty" xml:"openDataInstanceId,omitempty"`
	// This parameter is required.
	OpenTeamId   *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	OwnerUnionId *string `json:"ownerUnionId,omitempty" xml:"ownerUnionId,omitempty"`
}

func (s QueryInstancesByMultiConditionsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancesByMultiConditionsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *QueryInstancesByMultiConditionsResponseBodyRecords) SetCreatorUnionId(v string) *QueryInstancesByMultiConditionsResponseBodyRecords {
	s.CreatorUnionId = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBodyRecords) SetFields(v string) *QueryInstancesByMultiConditionsResponseBodyRecords {
	s.Fields = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBodyRecords) SetFormCode(v string) *QueryInstancesByMultiConditionsResponseBodyRecords {
	s.FormCode = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBodyRecords) SetGmtCreate(v string) *QueryInstancesByMultiConditionsResponseBodyRecords {
	s.GmtCreate = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBodyRecords) SetGmtModified(v string) *QueryInstancesByMultiConditionsResponseBodyRecords {
	s.GmtModified = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBodyRecords) SetModifiedUnionId(v string) *QueryInstancesByMultiConditionsResponseBodyRecords {
	s.ModifiedUnionId = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBodyRecords) SetOpenDataInstanceId(v string) *QueryInstancesByMultiConditionsResponseBodyRecords {
	s.OpenDataInstanceId = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBodyRecords) SetOpenTeamId(v string) *QueryInstancesByMultiConditionsResponseBodyRecords {
	s.OpenTeamId = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponseBodyRecords) SetOwnerUnionId(v string) *QueryInstancesByMultiConditionsResponseBodyRecords {
	s.OwnerUnionId = &v
	return s
}

type QueryInstancesByMultiConditionsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstancesByMultiConditionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstancesByMultiConditionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancesByMultiConditionsResponse) GoString() string {
	return s.String()
}

func (s *QueryInstancesByMultiConditionsResponse) SetHeaders(v map[string]*string) *QueryInstancesByMultiConditionsResponse {
	s.Headers = v
	return s
}

func (s *QueryInstancesByMultiConditionsResponse) SetStatusCode(v int32) *QueryInstancesByMultiConditionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstancesByMultiConditionsResponse) SetBody(v *QueryInstancesByMultiConditionsResponseBody) *QueryInstancesByMultiConditionsResponse {
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
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jci333
	OpenBatchTaskId *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jciwnfw
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
	TotalCount *int64                                           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySendMsgTaskStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QuerySendMsgTaskStatisticsResponse) SetStatusCode(v int32) *QuerySendMsgTaskStatisticsResponse {
	s.StatusCode = &v
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
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// J22222
	OpenBatchTaskId *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid1111
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jciwnfw
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
	TotalCount *int64                                                 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySendMsgTaskStatisticsDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QuerySendMsgTaskStatisticsDetailResponse) SetStatusCode(v int32) *QuerySendMsgTaskStatisticsDetailResponse {
	s.StatusCode = &v
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
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// CXiw
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxxx==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// msgxxxxxx==
	OpenMsgTaskId *string `json:"openMsgTaskId,omitempty" xml:"openMsgTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EifWwis
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
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken  *string                                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Records    []*QueryServiceGroupMessageReadStatusResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	TotalCount *int32                                                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	// example:
	//
	// 1
	ReadStatus *int32 `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
	// example:
	//
	// 2021-09-01 00:00:00
	ReadTimeStr *string `json:"readTimeStr,omitempty" xml:"readTimeStr,omitempty"`
	// example:
	//
	// $:LWCP_v1:xxxx==
	ReceiverDingTalkId *string `json:"receiverDingTalkId,omitempty" xml:"receiverDingTalkId,omitempty"`
	// example:
	//
	// 张三
	ReceiverName *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	// example:
	//
	// Kxiwk2
	ReceiverUnionId *string `json:"receiverUnionId,omitempty" xml:"receiverUnionId,omitempty"`
	// example:
	//
	// manager123
	ReceiverUserId *string `json:"receiverUserId,omitempty" xml:"receiverUserId,omitempty"`
	// example:
	//
	// 2021-09-01 00:00:00
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryServiceGroupMessageReadStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryServiceGroupMessageReadStatusResponse) SetStatusCode(v int32) *QueryServiceGroupMessageReadStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryServiceGroupMessageReadStatusResponse) SetBody(v *QueryServiceGroupMessageReadStatusResponseBody) *QueryServiceGroupMessageReadStatusResponse {
	s.Body = v
	return s
}

type QueueNotifyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueueNotifyHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueueNotifyHeaders) GoString() string {
	return s.String()
}

func (s *QueueNotifyHeaders) SetCommonHeaders(v map[string]*string) *QueueNotifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueueNotifyHeaders) SetXAcsDingtalkAccessToken(v string) *QueueNotifyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueueNotifyRequest struct {
	// example:
	//
	// 5
	EstimateWaitMin *int64 `json:"estimateWaitMin,omitempty" xml:"estimateWaitMin,omitempty"`
	// example:
	//
	// eWaJSqDcLsoiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 11
	QueuePlace *int64 `json:"queuePlace,omitempty" xml:"queuePlace,omitempty"`
	// example:
	//
	// 3333333333
	ServiceToken *string `json:"serviceToken,omitempty" xml:"serviceToken,omitempty"`
	// example:
	//
	// SourceTypeEnum
	TargetChannel *string `json:"targetChannel,omitempty" xml:"targetChannel,omitempty"`
	// example:
	//
	// 你好，欢迎来到这里
	Tips *string `json:"tips,omitempty" xml:"tips,omitempty"`
	// example:
	//
	// eeeeeeeeerrrrr
	VisitorToken *string `json:"visitorToken,omitempty" xml:"visitorToken,omitempty"`
}

func (s QueueNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s QueueNotifyRequest) GoString() string {
	return s.String()
}

func (s *QueueNotifyRequest) SetEstimateWaitMin(v int64) *QueueNotifyRequest {
	s.EstimateWaitMin = &v
	return s
}

func (s *QueueNotifyRequest) SetOpenTeamId(v string) *QueueNotifyRequest {
	s.OpenTeamId = &v
	return s
}

func (s *QueueNotifyRequest) SetQueuePlace(v int64) *QueueNotifyRequest {
	s.QueuePlace = &v
	return s
}

func (s *QueueNotifyRequest) SetServiceToken(v string) *QueueNotifyRequest {
	s.ServiceToken = &v
	return s
}

func (s *QueueNotifyRequest) SetTargetChannel(v string) *QueueNotifyRequest {
	s.TargetChannel = &v
	return s
}

func (s *QueueNotifyRequest) SetTips(v string) *QueueNotifyRequest {
	s.Tips = &v
	return s
}

func (s *QueueNotifyRequest) SetVisitorToken(v string) *QueueNotifyRequest {
	s.VisitorToken = &v
	return s
}

type QueueNotifyResponseBody struct {
	DingOpenErrcode *int32  `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueueNotifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueueNotifyResponseBody) GoString() string {
	return s.String()
}

func (s *QueueNotifyResponseBody) SetDingOpenErrcode(v int32) *QueueNotifyResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *QueueNotifyResponseBody) SetErrorMsg(v string) *QueueNotifyResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueueNotifyResponseBody) SetResult(v bool) *QueueNotifyResponseBody {
	s.Result = &v
	return s
}

func (s *QueueNotifyResponseBody) SetSuccess(v bool) *QueueNotifyResponseBody {
	s.Success = &v
	return s
}

type QueueNotifyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueueNotifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueueNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s QueueNotifyResponse) GoString() string {
	return s.String()
}

func (s *QueueNotifyResponse) SetHeaders(v map[string]*string) *QueueNotifyResponse {
	s.Headers = v
	return s
}

func (s *QueueNotifyResponse) SetStatusCode(v int32) *QueueNotifyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueueNotifyResponse) SetBody(v *QueueNotifyResponseBody) *QueueNotifyResponse {
	s.Body = v
	return s
}

type RemoveContactFromOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveContactFromOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveContactFromOrgHeaders) GoString() string {
	return s.String()
}

func (s *RemoveContactFromOrgHeaders) SetCommonHeaders(v map[string]*string) *RemoveContactFromOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveContactFromOrgHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveContactFromOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveContactFromOrgRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 88888
	ContactUnionId *string `json:"contactUnionId,omitempty" xml:"contactUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8888
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s RemoveContactFromOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveContactFromOrgRequest) GoString() string {
	return s.String()
}

func (s *RemoveContactFromOrgRequest) SetContactUnionId(v string) *RemoveContactFromOrgRequest {
	s.ContactUnionId = &v
	return s
}

func (s *RemoveContactFromOrgRequest) SetOpenTeamId(v string) *RemoveContactFromOrgRequest {
	s.OpenTeamId = &v
	return s
}

type RemoveContactFromOrgResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveContactFromOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveContactFromOrgResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveContactFromOrgResponseBody) SetResult(v bool) *RemoveContactFromOrgResponseBody {
	s.Result = &v
	return s
}

type RemoveContactFromOrgResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveContactFromOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveContactFromOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveContactFromOrgResponse) GoString() string {
	return s.String()
}

func (s *RemoveContactFromOrgResponse) SetHeaders(v map[string]*string) *RemoveContactFromOrgResponse {
	s.Headers = v
	return s
}

func (s *RemoveContactFromOrgResponse) SetStatusCode(v int32) *RemoveContactFromOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveContactFromOrgResponse) SetBody(v *RemoveContactFromOrgResponseBody) *RemoveContactFromOrgResponse {
	s.Body = v
	return s
}

type ReportCustomerDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReportCustomerDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReportCustomerDetailHeaders) GoString() string {
	return s.String()
}

func (s *ReportCustomerDetailHeaders) SetCommonHeaders(v map[string]*string) *ReportCustomerDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReportCustomerDetailHeaders) SetXAcsDingtalkAccessToken(v string) *ReportCustomerDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReportCustomerDetailRequest struct {
	// example:
	//
	// true
	HasLogin *bool `json:"hasLogin,omitempty" xml:"hasLogin,omitempty"`
	// example:
	//
	// true
	HasOpenConv *bool `json:"hasOpenConv,omitempty" xml:"hasOpenConv,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220102
	MaxDt *string `json:"maxDt,omitempty" xml:"maxDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MinDt *string `json:"minDt,omitempty" xml:"minDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iSoqrhLQDtK
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ReportCustomerDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportCustomerDetailRequest) GoString() string {
	return s.String()
}

func (s *ReportCustomerDetailRequest) SetHasLogin(v bool) *ReportCustomerDetailRequest {
	s.HasLogin = &v
	return s
}

func (s *ReportCustomerDetailRequest) SetHasOpenConv(v bool) *ReportCustomerDetailRequest {
	s.HasOpenConv = &v
	return s
}

func (s *ReportCustomerDetailRequest) SetMaxDt(v string) *ReportCustomerDetailRequest {
	s.MaxDt = &v
	return s
}

func (s *ReportCustomerDetailRequest) SetMinDt(v string) *ReportCustomerDetailRequest {
	s.MinDt = &v
	return s
}

func (s *ReportCustomerDetailRequest) SetOpenConversationId(v string) *ReportCustomerDetailRequest {
	s.OpenConversationId = &v
	return s
}

func (s *ReportCustomerDetailRequest) SetOpenTeamId(v string) *ReportCustomerDetailRequest {
	s.OpenTeamId = &v
	return s
}

func (s *ReportCustomerDetailRequest) SetPageNumber(v int64) *ReportCustomerDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *ReportCustomerDetailRequest) SetPageSize(v int64) *ReportCustomerDetailRequest {
	s.PageSize = &v
	return s
}

type ReportCustomerDetailResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	Records []*ReportCustomerDetailResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ReportCustomerDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportCustomerDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ReportCustomerDetailResponseBody) SetCurrentPage(v int64) *ReportCustomerDetailResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ReportCustomerDetailResponseBody) SetPageSize(v int64) *ReportCustomerDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *ReportCustomerDetailResponseBody) SetRecords(v []*ReportCustomerDetailResponseBodyRecords) *ReportCustomerDetailResponseBody {
	s.Records = v
	return s
}

func (s *ReportCustomerDetailResponseBody) SetTotalCount(v int64) *ReportCustomerDetailResponseBody {
	s.TotalCount = &v
	return s
}

type ReportCustomerDetailResponseBodyRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AtRobotCnt *int64 `json:"atRobotCnt,omitempty" xml:"atRobotCnt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	CustomerName *string `json:"customerName,omitempty" xml:"customerName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试群
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	HasLogin *bool `json:"hasLogin,omitempty" xml:"hasLogin,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	HasOpenConv *bool `json:"hasOpenConv,omitempty" xml:"hasOpenConv,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	SendMsgCnt *int64 `json:"sendMsgCnt,omitempty" xml:"sendMsgCnt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xcjlsdf
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56789
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ReportCustomerDetailResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ReportCustomerDetailResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ReportCustomerDetailResponseBodyRecords) SetAtRobotCnt(v int64) *ReportCustomerDetailResponseBodyRecords {
	s.AtRobotCnt = &v
	return s
}

func (s *ReportCustomerDetailResponseBodyRecords) SetCustomerName(v string) *ReportCustomerDetailResponseBodyRecords {
	s.CustomerName = &v
	return s
}

func (s *ReportCustomerDetailResponseBodyRecords) SetGroupName(v string) *ReportCustomerDetailResponseBodyRecords {
	s.GroupName = &v
	return s
}

func (s *ReportCustomerDetailResponseBodyRecords) SetHasLogin(v bool) *ReportCustomerDetailResponseBodyRecords {
	s.HasLogin = &v
	return s
}

func (s *ReportCustomerDetailResponseBodyRecords) SetHasOpenConv(v bool) *ReportCustomerDetailResponseBodyRecords {
	s.HasOpenConv = &v
	return s
}

func (s *ReportCustomerDetailResponseBodyRecords) SetSendMsgCnt(v int64) *ReportCustomerDetailResponseBodyRecords {
	s.SendMsgCnt = &v
	return s
}

func (s *ReportCustomerDetailResponseBodyRecords) SetUnionId(v string) *ReportCustomerDetailResponseBodyRecords {
	s.UnionId = &v
	return s
}

func (s *ReportCustomerDetailResponseBodyRecords) SetUserId(v string) *ReportCustomerDetailResponseBodyRecords {
	s.UserId = &v
	return s
}

type ReportCustomerDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportCustomerDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportCustomerDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportCustomerDetailResponse) GoString() string {
	return s.String()
}

func (s *ReportCustomerDetailResponse) SetHeaders(v map[string]*string) *ReportCustomerDetailResponse {
	s.Headers = v
	return s
}

func (s *ReportCustomerDetailResponse) SetStatusCode(v int32) *ReportCustomerDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportCustomerDetailResponse) SetBody(v *ReportCustomerDetailResponseBody) *ReportCustomerDetailResponse {
	s.Body = v
	return s
}

type ReportCustomerStatisticsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReportCustomerStatisticsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReportCustomerStatisticsHeaders) GoString() string {
	return s.String()
}

func (s *ReportCustomerStatisticsHeaders) SetCommonHeaders(v map[string]*string) *ReportCustomerStatisticsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReportCustomerStatisticsHeaders) SetXAcsDingtalkAccessToken(v string) *ReportCustomerStatisticsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReportCustomerStatisticsRequest struct {
	// if can be null:
	// true
	GroupOwnerUserIds []*string `json:"groupOwnerUserIds,omitempty" xml:"groupOwnerUserIds,omitempty" type:"Repeated"`
	// if can be null:
	// true
	GroupTags []*string `json:"groupTags,omitempty" xml:"groupTags,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 20220102
	MaxDt *string `json:"maxDt,omitempty" xml:"maxDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MinDt *string `json:"minDt,omitempty" xml:"minDt,omitempty"`
	// if can be null:
	// true
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	// example:
	//
	// iFoqrhLQDtK
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iSoqrhLQDtK
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ReportCustomerStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportCustomerStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ReportCustomerStatisticsRequest) SetGroupOwnerUserIds(v []*string) *ReportCustomerStatisticsRequest {
	s.GroupOwnerUserIds = v
	return s
}

func (s *ReportCustomerStatisticsRequest) SetGroupTags(v []*string) *ReportCustomerStatisticsRequest {
	s.GroupTags = v
	return s
}

func (s *ReportCustomerStatisticsRequest) SetMaxDt(v string) *ReportCustomerStatisticsRequest {
	s.MaxDt = &v
	return s
}

func (s *ReportCustomerStatisticsRequest) SetMinDt(v string) *ReportCustomerStatisticsRequest {
	s.MinDt = &v
	return s
}

func (s *ReportCustomerStatisticsRequest) SetOpenConversationIds(v []*string) *ReportCustomerStatisticsRequest {
	s.OpenConversationIds = v
	return s
}

func (s *ReportCustomerStatisticsRequest) SetOpenGroupSetId(v string) *ReportCustomerStatisticsRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *ReportCustomerStatisticsRequest) SetOpenTeamId(v string) *ReportCustomerStatisticsRequest {
	s.OpenTeamId = &v
	return s
}

func (s *ReportCustomerStatisticsRequest) SetPageNumber(v int64) *ReportCustomerStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *ReportCustomerStatisticsRequest) SetPageSize(v int64) *ReportCustomerStatisticsRequest {
	s.PageSize = &v
	return s
}

type ReportCustomerStatisticsResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	Records []*ReportCustomerStatisticsResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ReportCustomerStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportCustomerStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ReportCustomerStatisticsResponseBody) SetCurrentPage(v int64) *ReportCustomerStatisticsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBody) SetPageSize(v int64) *ReportCustomerStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBody) SetRecords(v []*ReportCustomerStatisticsResponseBodyRecords) *ReportCustomerStatisticsResponseBody {
	s.Records = v
	return s
}

func (s *ReportCustomerStatisticsResponseBody) SetTotalCount(v int64) *ReportCustomerStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

type ReportCustomerStatisticsResponseBodyRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AtRobotCnt *int64 `json:"atRobotCnt,omitempty" xml:"atRobotCnt,omitempty"`
	// example:
	//
	// bizXX
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	CustomerCnt *int64 `json:"customerCnt,omitempty" xml:"customerCnt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试群
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试群分组
	GroupSetName *string `json:"groupSetName,omitempty" xml:"groupSetName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	LoginCnt *int64 `json:"loginCnt,omitempty" xml:"loginCnt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	OpenConvCnt *int64 `json:"openConvCnt,omitempty" xml:"openConvCnt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iSoqrhLQDtK
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	SendMsgCnt *int64 `json:"sendMsgCnt,omitempty" xml:"sendMsgCnt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	SenderCnt *int64 `json:"senderCnt,omitempty" xml:"senderCnt,omitempty"`
}

func (s ReportCustomerStatisticsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ReportCustomerStatisticsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetAtRobotCnt(v int64) *ReportCustomerStatisticsResponseBodyRecords {
	s.AtRobotCnt = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetBizId(v string) *ReportCustomerStatisticsResponseBodyRecords {
	s.BizId = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetCustomerCnt(v int64) *ReportCustomerStatisticsResponseBodyRecords {
	s.CustomerCnt = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetGroupName(v string) *ReportCustomerStatisticsResponseBodyRecords {
	s.GroupName = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetGroupSetName(v string) *ReportCustomerStatisticsResponseBodyRecords {
	s.GroupSetName = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetLoginCnt(v int64) *ReportCustomerStatisticsResponseBodyRecords {
	s.LoginCnt = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetOpenConvCnt(v int64) *ReportCustomerStatisticsResponseBodyRecords {
	s.OpenConvCnt = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetOpenConversationId(v string) *ReportCustomerStatisticsResponseBodyRecords {
	s.OpenConversationId = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetOpenGroupSetId(v string) *ReportCustomerStatisticsResponseBodyRecords {
	s.OpenGroupSetId = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetSendMsgCnt(v int64) *ReportCustomerStatisticsResponseBodyRecords {
	s.SendMsgCnt = &v
	return s
}

func (s *ReportCustomerStatisticsResponseBodyRecords) SetSenderCnt(v int64) *ReportCustomerStatisticsResponseBodyRecords {
	s.SenderCnt = &v
	return s
}

type ReportCustomerStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportCustomerStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportCustomerStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportCustomerStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ReportCustomerStatisticsResponse) SetHeaders(v map[string]*string) *ReportCustomerStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ReportCustomerStatisticsResponse) SetStatusCode(v int32) *ReportCustomerStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportCustomerStatisticsResponse) SetBody(v *ReportCustomerStatisticsResponseBody) *ReportCustomerStatisticsResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// Dq9hP8Sk2v6vQ6l05nCe5wiEiE
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// example:
	//
	// [{\"identifier\":\"input1\",\"value\":\"123\"}]
	CustomFields *string                      `json:"customFields,omitempty" xml:"customFields,omitempty"`
	Notify       *ResubmitTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bLkvfXKiSngQiE
	OpenTemplateBizId *string `json:"openTemplateBizId,omitempty" xml:"openTemplateBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iPbrfXjdNjRoiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// This parameter is required.
	ProcessorUnionIds []*string `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// SG
	Scene        *string                            `json:"scene,omitempty" xml:"scene,omitempty"`
	SceneContext *ResubmitTicketRequestSceneContext `json:"sceneContext,omitempty" xml:"sceneContext,omitempty" type:"Struct"`
	TicketMemo   *ResubmitTicketRequestTicketMemo   `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
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
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	NoticeAllGroupMember       *bool     `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
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
	// example:
	//
	// cidZBSNlUi/Jq9x76PAXUCrAA==
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
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
	// example:
	//
	// msgsbY4BzTCNX0/ClUwoTTs7w==
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
	// example:
	//
	// 备注
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
	// example:
	//
	// wahaha.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// ticket/image/44708069/43003/e27204b382c04832aec4243e940a1367_1625831640499.txt
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *ResubmitTicketResponse) SetStatusCode(v int32) *ResubmitTicketResponse {
	s.StatusCode = &v
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
	Notify *RetractTicketRequestNotify `json:"notify,omitempty" xml:"notify,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a8iS4X94TgtgiE
	OpenTicketId    *string                         `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
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
	// example:
	//
	// true
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
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *RetractTicketResponse) SetStatusCode(v int32) *RetractTicketResponse {
	s.StatusCode = &v
	return s
}

type RobotMessageRecallHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RobotMessageRecallHeaders) String() string {
	return tea.Prettify(s)
}

func (s RobotMessageRecallHeaders) GoString() string {
	return s.String()
}

func (s *RobotMessageRecallHeaders) SetCommonHeaders(v map[string]*string) *RobotMessageRecallHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RobotMessageRecallHeaders) SetXAcsDingtalkAccessToken(v string) *RobotMessageRecallHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RobotMessageRecallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// msgU87r5gnMP43JTDAZg/ETyQ==
	OpenMsgId *string `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iSoqrhLQDtK
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s RobotMessageRecallRequest) String() string {
	return tea.Prettify(s)
}

func (s RobotMessageRecallRequest) GoString() string {
	return s.String()
}

func (s *RobotMessageRecallRequest) SetOpenConversationId(v string) *RobotMessageRecallRequest {
	s.OpenConversationId = &v
	return s
}

func (s *RobotMessageRecallRequest) SetOpenMsgId(v string) *RobotMessageRecallRequest {
	s.OpenMsgId = &v
	return s
}

func (s *RobotMessageRecallRequest) SetOpenTeamId(v string) *RobotMessageRecallRequest {
	s.OpenTeamId = &v
	return s
}

type RobotMessageRecallResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RobotMessageRecallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RobotMessageRecallResponseBody) GoString() string {
	return s.String()
}

func (s *RobotMessageRecallResponseBody) SetResult(v string) *RobotMessageRecallResponseBody {
	s.Result = &v
	return s
}

type RobotMessageRecallResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RobotMessageRecallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RobotMessageRecallResponse) String() string {
	return tea.Prettify(s)
}

func (s RobotMessageRecallResponse) GoString() string {
	return s.String()
}

func (s *RobotMessageRecallResponse) SetHeaders(v map[string]*string) *RobotMessageRecallResponse {
	s.Headers = v
	return s
}

func (s *RobotMessageRecallResponse) SetStatusCode(v int32) *RobotMessageRecallResponse {
	s.StatusCode = &v
	return s
}

func (s *RobotMessageRecallResponse) SetBody(v *RobotMessageRecallResponseBody) *RobotMessageRecallResponse {
	s.Body = v
	return s
}

type SaveFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *SaveFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *SaveFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *SaveFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveFormInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {}
	FormDataList *string `json:"formDataList,omitempty" xml:"formDataList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8888
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 88888
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	// example:
	//
	// 8888
	OwnerUnionId *string `json:"ownerUnionId,omitempty" xml:"ownerUnionId,omitempty"`
}

func (s SaveFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *SaveFormInstanceRequest) SetFormDataList(v string) *SaveFormInstanceRequest {
	s.FormDataList = &v
	return s
}

func (s *SaveFormInstanceRequest) SetOpenTeamId(v string) *SaveFormInstanceRequest {
	s.OpenTeamId = &v
	return s
}

func (s *SaveFormInstanceRequest) SetOperatorUnionId(v string) *SaveFormInstanceRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *SaveFormInstanceRequest) SetOwnerUnionId(v string) *SaveFormInstanceRequest {
	s.OwnerUnionId = &v
	return s
}

type SaveFormInstanceResponseBody struct {
	// example:
	//
	// 99999
	OpenContactId *string `json:"openContactId,omitempty" xml:"openContactId,omitempty"`
	// example:
	//
	// 88888
	OpenCustomerId *string `json:"openCustomerId,omitempty" xml:"openCustomerId,omitempty"`
}

func (s SaveFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveFormInstanceResponseBody) SetOpenContactId(v string) *SaveFormInstanceResponseBody {
	s.OpenContactId = &v
	return s
}

func (s *SaveFormInstanceResponseBody) SetOpenCustomerId(v string) *SaveFormInstanceResponseBody {
	s.OpenCustomerId = &v
	return s
}

type SaveFormInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *SaveFormInstanceResponse) SetHeaders(v map[string]*string) *SaveFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *SaveFormInstanceResponse) SetStatusCode(v int32) *SaveFormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveFormInstanceResponse) SetBody(v *SaveFormInstanceResponseBody) *SaveFormInstanceResponse {
	s.Body = v
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
	// example:
	//
	// 钉钉专属服务群
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// CXiw
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// cidxxxx==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// sjfuwid
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// example:
	//
	// jfuwida
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 目前支持PAGE 和 SCROLL，默认PAGE类型
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
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken  *string                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Records    []*SearchGroupResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	TotalCount *int32                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 钉钉专属服务群
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingtalk:xxx
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxx==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xjfjdsiw
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xkjhfker
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SearchGroupResponse) SetStatusCode(v int32) *SearchGroupResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	MessageContent *SendMsgByTaskRequestMessageContent `json:"messageContent,omitempty" xml:"messageContent,omitempty" type:"Struct"`
	// This parameter is required.
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	QueryGroup *SendMsgByTaskRequestQueryGroup `json:"queryGroup,omitempty" xml:"queryGroup,omitempty" type:"Struct"`
	// This parameter is required.
	SendConfig *SendMsgByTaskRequestSendConfig `json:"sendConfig,omitempty" xml:"sendConfig,omitempty" type:"Struct"`
	// This parameter is required.
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
	AtActiveMemberNum *int64                                    `json:"atActiveMemberNum,omitempty" xml:"atActiveMemberNum,omitempty"`
	AtActiveUser      *bool                                     `json:"atActiveUser,omitempty" xml:"atActiveUser,omitempty"`
	AtAll             *bool                                     `json:"atAll,omitempty" xml:"atAll,omitempty"`
	Btns              []*SendMsgByTaskRequestMessageContentBtns `json:"btns,omitempty" xml:"btns,omitempty" type:"Repeated"`
	Content           *string                                   `json:"content,omitempty" xml:"content,omitempty"`
	Images            []*string                                 `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// This parameter is required.
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	Remind      *bool   `json:"remind,omitempty" xml:"remind,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	Top         *bool   `json:"top,omitempty" xml:"top,omitempty"`
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

func (s *SendMsgByTaskRequestMessageContent) SetRemind(v bool) *SendMsgByTaskRequestMessageContent {
	s.Remind = &v
	return s
}

func (s *SendMsgByTaskRequestMessageContent) SetTitle(v string) *SendMsgByTaskRequestMessageContent {
	s.Title = &v
	return s
}

func (s *SendMsgByTaskRequestMessageContent) SetTop(v bool) *SendMsgByTaskRequestMessageContent {
	s.Top = &v
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
	GroupTagNames            []*string `json:"groupTagNames,omitempty" xml:"groupTagNames,omitempty" type:"Repeated"`
	LastActiveDateFilterType *string   `json:"lastActiveDateFilterType,omitempty" xml:"lastActiveDateFilterType,omitempty"`
	LastActiveTimeEnd        *string   `json:"lastActiveTimeEnd,omitempty" xml:"lastActiveTimeEnd,omitempty"`
	LastActiveTimeStart      *string   `json:"lastActiveTimeStart,omitempty" xml:"lastActiveTimeStart,omitempty"`
	OpenConversationIds      []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	OpenGroupSetId           *string   `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// This parameter is required.
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
	NeedUrlTrack *bool   `json:"needUrlTrack,omitempty" xml:"needUrlTrack,omitempty"`
	SendTime     *string `json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	// This parameter is required.
	SendType       *string                                         `json:"sendType,omitempty" xml:"sendType,omitempty"`
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
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
	TrackId  *string `json:"trackId,omitempty" xml:"trackId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMsgByTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SendMsgByTaskResponse) SetStatusCode(v int32) *SendMsgByTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMsgByTaskResponse) SetBody(v *SendMsgByTaskResponseBody) *SendMsgByTaskResponse {
	s.Body = v
	return s
}

type SendMsgByTaskSupportInviteJoinOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendMsgByTaskSupportInviteJoinOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskSupportInviteJoinOrgHeaders) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskSupportInviteJoinOrgHeaders) SetCommonHeaders(v map[string]*string) *SendMsgByTaskSupportInviteJoinOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgHeaders) SetXAcsDingtalkAccessToken(v string) *SendMsgByTaskSupportInviteJoinOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendMsgByTaskSupportInviteJoinOrgRequest struct {
	// This parameter is required.
	MessageContent *SendMsgByTaskSupportInviteJoinOrgRequestMessageContent `json:"messageContent,omitempty" xml:"messageContent,omitempty" type:"Struct"`
	// This parameter is required.
	MobilePhones []*string `json:"mobilePhones,omitempty" xml:"mobilePhones,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true
	NeedUrlTrack *bool `json:"needUrlTrack,omitempty" xml:"needUrlTrack,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 88888
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 发送渠道      	- 工作通知：WORK_NOTICE      	- 机器人：SINGLE_ROBOT
	SendChannel *string `json:"sendChannel,omitempty" xml:"sendChannel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 群发任务
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s SendMsgByTaskSupportInviteJoinOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskSupportInviteJoinOrgRequest) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequest) SetMessageContent(v *SendMsgByTaskSupportInviteJoinOrgRequestMessageContent) *SendMsgByTaskSupportInviteJoinOrgRequest {
	s.MessageContent = v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequest) SetMobilePhones(v []*string) *SendMsgByTaskSupportInviteJoinOrgRequest {
	s.MobilePhones = v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequest) SetNeedUrlTrack(v bool) *SendMsgByTaskSupportInviteJoinOrgRequest {
	s.NeedUrlTrack = &v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequest) SetOpenTeamId(v string) *SendMsgByTaskSupportInviteJoinOrgRequest {
	s.OpenTeamId = &v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequest) SetSendChannel(v string) *SendMsgByTaskSupportInviteJoinOrgRequest {
	s.SendChannel = &v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequest) SetTaskName(v string) *SendMsgByTaskSupportInviteJoinOrgRequest {
	s.TaskName = &v
	return s
}

type SendMsgByTaskSupportInviteJoinOrgRequestMessageContent struct {
	Btns []*SendMsgByTaskSupportInviteJoinOrgRequestMessageContentBtns `json:"btns,omitempty" xml:"btns,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ACTIONCARD：卡片消息
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 标题内容
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendMsgByTaskSupportInviteJoinOrgRequestMessageContent) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskSupportInviteJoinOrgRequestMessageContent) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequestMessageContent) SetBtns(v []*SendMsgByTaskSupportInviteJoinOrgRequestMessageContentBtns) *SendMsgByTaskSupportInviteJoinOrgRequestMessageContent {
	s.Btns = v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequestMessageContent) SetContent(v string) *SendMsgByTaskSupportInviteJoinOrgRequestMessageContent {
	s.Content = &v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequestMessageContent) SetMessageType(v string) *SendMsgByTaskSupportInviteJoinOrgRequestMessageContent {
	s.MessageType = &v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequestMessageContent) SetTitle(v string) *SendMsgByTaskSupportInviteJoinOrgRequestMessageContent {
	s.Title = &v
	return s
}

type SendMsgByTaskSupportInviteJoinOrgRequestMessageContentBtns struct {
	// example:
	//
	// http://www.baidu.com
	ActionURL *string `json:"actionURL,omitempty" xml:"actionURL,omitempty"`
	// example:
	//
	// 按钮标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendMsgByTaskSupportInviteJoinOrgRequestMessageContentBtns) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskSupportInviteJoinOrgRequestMessageContentBtns) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequestMessageContentBtns) SetActionURL(v string) *SendMsgByTaskSupportInviteJoinOrgRequestMessageContentBtns {
	s.ActionURL = &v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgRequestMessageContentBtns) SetTitle(v string) *SendMsgByTaskSupportInviteJoinOrgRequestMessageContentBtns {
	s.Title = &v
	return s
}

type SendMsgByTaskSupportInviteJoinOrgResponseBody struct {
	OpenBatchTaskId *string `json:"openBatchTaskId,omitempty" xml:"openBatchTaskId,omitempty"`
}

func (s SendMsgByTaskSupportInviteJoinOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskSupportInviteJoinOrgResponseBody) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskSupportInviteJoinOrgResponseBody) SetOpenBatchTaskId(v string) *SendMsgByTaskSupportInviteJoinOrgResponseBody {
	s.OpenBatchTaskId = &v
	return s
}

type SendMsgByTaskSupportInviteJoinOrgResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMsgByTaskSupportInviteJoinOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMsgByTaskSupportInviteJoinOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMsgByTaskSupportInviteJoinOrgResponse) GoString() string {
	return s.String()
}

func (s *SendMsgByTaskSupportInviteJoinOrgResponse) SetHeaders(v map[string]*string) *SendMsgByTaskSupportInviteJoinOrgResponse {
	s.Headers = v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgResponse) SetStatusCode(v int32) *SendMsgByTaskSupportInviteJoinOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMsgByTaskSupportInviteJoinOrgResponse) SetBody(v *SendMsgByTaskSupportInviteJoinOrgResponseBody) *SendMsgByTaskSupportInviteJoinOrgResponse {
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
	AtDingtalkIds []*string `json:"atDingtalkIds,omitempty" xml:"atDingtalkIds,omitempty" type:"Repeated"`
	AtMobiles     []*string `json:"atMobiles,omitempty" xml:"atMobiles,omitempty" type:"Repeated"`
	AtUnionIds    []*string `json:"atUnionIds,omitempty" xml:"atUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	BtnOrientation *string                               `json:"btnOrientation,omitempty" xml:"btnOrientation,omitempty"`
	Btns           []*SendServiceGroupMessageRequestBtns `json:"btns,omitempty" xml:"btns,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 你有新的任务待审批
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// false
	HasContentLinks *bool `json:"hasContentLinks,omitempty" xml:"hasContentLinks,omitempty"`
	// example:
	//
	// false
	IsAtAll *bool `json:"isAtAll,omitempty" xml:"isAtAll,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MARKDOWN
	MessageType         *string   `json:"messageType,omitempty" xml:"messageType,omitempty"`
	ReceiverDingtalkIds []*string `json:"receiverDingtalkIds,omitempty" xml:"receiverDingtalkIds,omitempty" type:"Repeated"`
	ReceiverMobiles     []*string `json:"receiverMobiles,omitempty" xml:"receiverMobiles,omitempty" type:"Repeated"`
	ReceiverUnionIds    []*string `json:"receiverUnionIds,omitempty" xml:"receiverUnionIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxx==
	TargetOpenConversationId *string `json:"targetOpenConversationId,omitempty" xml:"targetOpenConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 服务提醒
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
	// example:
	//
	// http://www.dingtalk.com
	ActionURL *string `json:"actionURL,omitempty" xml:"actionURL,omitempty"`
	// example:
	//
	// 测试按钮
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
	// This parameter is required.
	//
	// example:
	//
	// msgxxxxxx==
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendServiceGroupMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SendServiceGroupMessageResponse) SetStatusCode(v int32) *SendServiceGroupMessageResponse {
	s.StatusCode = &v
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
	OpenGroupSetId     *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	OpenTeamId         *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 0
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
	ConfigKey   *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRobotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SetRobotConfigResponse) SetStatusCode(v int32) *SetRobotConfigResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a8iS4X94TgtgiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dq9hP8Sk2v6vQ6l05nCe5wiEiE
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *TakeTicketResponse) SetStatusCode(v int32) *TakeTicketResponse {
	s.StatusCode = &v
	return s
}

type TopicStatisticsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TopicStatisticsHeaders) String() string {
	return tea.Prettify(s)
}

func (s TopicStatisticsHeaders) GoString() string {
	return s.String()
}

func (s *TopicStatisticsHeaders) SetCommonHeaders(v map[string]*string) *TopicStatisticsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TopicStatisticsHeaders) SetXAcsDingtalkAccessToken(v string) *TopicStatisticsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TopicStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MaxDt *string `json:"maxDt,omitempty" xml:"maxDt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	MinDt *string `json:"minDt,omitempty" xml:"minDt,omitempty"`
	// example:
	//
	// cidXX,cidYY
	OpenConversationIds *string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KxisoOk
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 工单
	SearchContent *string `json:"searchContent,omitempty" xml:"searchContent,omitempty"`
}

func (s TopicStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s TopicStatisticsRequest) GoString() string {
	return s.String()
}

func (s *TopicStatisticsRequest) SetMaxDt(v string) *TopicStatisticsRequest {
	s.MaxDt = &v
	return s
}

func (s *TopicStatisticsRequest) SetMinDt(v string) *TopicStatisticsRequest {
	s.MinDt = &v
	return s
}

func (s *TopicStatisticsRequest) SetOpenConversationIds(v string) *TopicStatisticsRequest {
	s.OpenConversationIds = &v
	return s
}

func (s *TopicStatisticsRequest) SetOpenTeamId(v string) *TopicStatisticsRequest {
	s.OpenTeamId = &v
	return s
}

func (s *TopicStatisticsRequest) SetSearchContent(v string) *TopicStatisticsRequest {
	s.SearchContent = &v
	return s
}

type TopicStatisticsResponseBody struct {
	// This parameter is required.
	TopicStatisticsRecords []*TopicStatisticsResponseBodyTopicStatisticsRecords `json:"topicStatisticsRecords,omitempty" xml:"topicStatisticsRecords,omitempty" type:"Repeated"`
}

func (s TopicStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TopicStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *TopicStatisticsResponseBody) SetTopicStatisticsRecords(v []*TopicStatisticsResponseBodyTopicStatisticsRecords) *TopicStatisticsResponseBody {
	s.TopicStatisticsRecords = v
	return s
}

type TopicStatisticsResponseBodyTopicStatisticsRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	Dt *string `json:"dt,omitempty" xml:"dt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	MsgCount *int64 `json:"msgCount,omitempty" xml:"msgCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ParticipantsNum *int64 `json:"participantsNum,omitempty" xml:"participantsNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	TopicNum *int64 `json:"topicNum,omitempty" xml:"topicNum,omitempty"`
}

func (s TopicStatisticsResponseBodyTopicStatisticsRecords) String() string {
	return tea.Prettify(s)
}

func (s TopicStatisticsResponseBodyTopicStatisticsRecords) GoString() string {
	return s.String()
}

func (s *TopicStatisticsResponseBodyTopicStatisticsRecords) SetDt(v string) *TopicStatisticsResponseBodyTopicStatisticsRecords {
	s.Dt = &v
	return s
}

func (s *TopicStatisticsResponseBodyTopicStatisticsRecords) SetMsgCount(v int64) *TopicStatisticsResponseBodyTopicStatisticsRecords {
	s.MsgCount = &v
	return s
}

func (s *TopicStatisticsResponseBodyTopicStatisticsRecords) SetParticipantsNum(v int64) *TopicStatisticsResponseBodyTopicStatisticsRecords {
	s.ParticipantsNum = &v
	return s
}

func (s *TopicStatisticsResponseBodyTopicStatisticsRecords) SetTopicNum(v int64) *TopicStatisticsResponseBodyTopicStatisticsRecords {
	s.TopicNum = &v
	return s
}

type TopicStatisticsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TopicStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TopicStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s TopicStatisticsResponse) GoString() string {
	return s.String()
}

func (s *TopicStatisticsResponse) SetHeaders(v map[string]*string) *TopicStatisticsResponse {
	s.Headers = v
	return s
}

func (s *TopicStatisticsResponse) SetStatusCode(v int32) *TopicStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *TopicStatisticsResponse) SetBody(v *TopicStatisticsResponseBody) *TopicStatisticsResponse {
	s.Body = v
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
	// This parameter is required.
	//
	// example:
	//
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iPFWCyMGWPiiIiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dq9hP8Sk2v6vQ6l05nCe5wiEiE
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// This parameter is required.
	ProcessorUnionIds []*string                        `json:"processorUnionIds,omitempty" xml:"processorUnionIds,omitempty" type:"Repeated"`
	TicketMemo        *TransferTicketRequestTicketMemo `json:"ticketMemo,omitempty" xml:"ticketMemo,omitempty" type:"Struct"`
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
	GroupNoticeReceiverUnionIds []*string `json:"groupNoticeReceiverUnionIds,omitempty" xml:"groupNoticeReceiverUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	NoticeAllGroupMember       *bool     `json:"noticeAllGroupMember,omitempty" xml:"noticeAllGroupMember,omitempty"`
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
	Attachments []*TransferTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// example:
	//
	// 备注
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
	// example:
	//
	// wahaha.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// ticket/image/44708069/43003/e27204b382c04832aec4243e940a1367_1625831640499.txt
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *TransferTicketResponse) SetStatusCode(v int32) *TransferTicketResponse {
	s.StatusCode = &v
	return s
}

type UpdateGroupSetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupSetHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSetHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupSetHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupSetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupSetHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupSetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupSetRequest struct {
	// example:
	//
	// cidkeQXxEC9VaGQ2M6nTlhNWQ==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// iPnLAZk8oV4AiE
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// example:
	//
	// u9iSGISL3bqIiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
}

func (s UpdateGroupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupSetRequest) SetOpenConversationId(v string) *UpdateGroupSetRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateGroupSetRequest) SetOpenGroupSetId(v string) *UpdateGroupSetRequest {
	s.OpenGroupSetId = &v
	return s
}

func (s *UpdateGroupSetRequest) SetOpenTeamId(v string) *UpdateGroupSetRequest {
	s.OpenTeamId = &v
	return s
}

type UpdateGroupSetResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateGroupSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupSetResponseBody) SetSuccess(v bool) *UpdateGroupSetResponseBody {
	s.Success = &v
	return s
}

type UpdateGroupSetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupSetResponse) SetHeaders(v map[string]*string) *UpdateGroupSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupSetResponse) SetStatusCode(v int32) *UpdateGroupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupSetResponse) SetBody(v *UpdateGroupSetResponseBody) *UpdateGroupSetResponse {
	s.Body = v
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
	// This parameter is required.
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	TagNames            []*string `json:"tagNames,omitempty" xml:"tagNames,omitempty" type:"Repeated"`
	// This parameter is required.
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *UpdateGroupTagResponse) SetStatusCode(v int32) *UpdateGroupTagResponse {
	s.StatusCode = &v
	return s
}

type UpdateInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInstanceHeaders) SetCommonHeaders(v map[string]*string) *UpdateInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInstanceRequest struct {
	// example:
	//
	// hhdhg
	ExternalBizId *string `json:"externalBizId,omitempty" xml:"externalBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DING_CUSTOMER
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"node_888":"hhhh"}
	FormDataList *string `json:"formDataList,omitempty" xml:"formDataList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 888555
	OpenDataInstanceId *string `json:"openDataInstanceId,omitempty" xml:"openDataInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 888***
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// example:
	//
	// 8888
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	// example:
	//
	// 88888
	OwnerUnionId *string `json:"ownerUnionId,omitempty" xml:"ownerUnionId,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetExternalBizId(v string) *UpdateInstanceRequest {
	s.ExternalBizId = &v
	return s
}

func (s *UpdateInstanceRequest) SetFormCode(v string) *UpdateInstanceRequest {
	s.FormCode = &v
	return s
}

func (s *UpdateInstanceRequest) SetFormDataList(v string) *UpdateInstanceRequest {
	s.FormDataList = &v
	return s
}

func (s *UpdateInstanceRequest) SetOpenDataInstanceId(v string) *UpdateInstanceRequest {
	s.OpenDataInstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetOpenTeamId(v string) *UpdateInstanceRequest {
	s.OpenTeamId = &v
	return s
}

func (s *UpdateInstanceRequest) SetOperatorUnionId(v string) *UpdateInstanceRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *UpdateInstanceRequest) SetOwnerUnionId(v string) *UpdateInstanceRequest {
	s.OwnerUnionId = &v
	return s
}

type UpdateInstanceResponseBody struct {
	OpenDataInstanceId *string `json:"openDataInstanceId,omitempty" xml:"openDataInstanceId,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetOpenDataInstanceId(v string) *UpdateInstanceResponseBody {
	s.OpenDataInstanceId = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
	s.Body = v
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
	// example:
	//
	// [{\"identifier\":\"input1\",\"value\":\"openAPI更新了值\"}]
	CustomFields *string `json:"customFields,omitempty" xml:"customFields,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eKWh3GBwsKEiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iPFWCyMGWPiiIiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// p8VdSjm884SvQ6l05nCe5wiEiE
	ProcessorUnionId *string `json:"processorUnionId,omitempty" xml:"processorUnionId,omitempty"`
	// This parameter is required.
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
	Attachments []*UpdateTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 备注
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
	// example:
	//
	// wahaha.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// ticket/image/44708069/43003/e27204b382c04832aec4243e940a1367_1625831640499.txt
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *UpdateTicketResponse) SetStatusCode(v int32) *UpdateTicketResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// sdfdfser
	CcsInstanceId *string `json:"ccsInstanceId,omitempty" xml:"ccsInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidrQnTVXH/X+ERaVqGaH+asw==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// oPnDlfVYYIUia
	OpenGroupSetId *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	// example:
	//
	// btkoYsadwyQiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *UpgradeCloudGroupResponse) SetStatusCode(v int32) *UpgradeCloudGroupResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenGroupSetId     *string `json:"openGroupSetId,omitempty" xml:"openGroupSetId,omitempty"`
	OpenTeamId         *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	TemplateId         *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *UpgradeNormalGroupResponse) SetStatusCode(v int32) *UpgradeNormalGroupResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// bLkvfXKiSngQiE
	OpenTeamId *string `json:"openTeamId,omitempty" xml:"openTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iPbrfXjdNjRoiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dq9hP8Sk2v6vQ6l05nCe5wiEiE
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	// This parameter is required.
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
	Attachments []*UrgeTicketRequestTicketMemoAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 备注
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
	// example:
	//
	// wahaha.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// ticket/image/44708069/43003/e27204b382c04832aec4243e940a1367_1625831640499.txt
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *UrgeTicketResponse) SetStatusCode(v int32) *UrgeTicketResponse {
	s.StatusCode = &v
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
// 添加联系人到群里
//
// @param request - AddContactMemberToGroupRequest
//
// @param headers - AddContactMemberToGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddContactMemberToGroupResponse
func (client *Client) AddContactMemberToGroupWithOptions(request *AddContactMemberToGroupRequest, headers *AddContactMemberToGroupHeaders, runtime *util.RuntimeOptions) (_result *AddContactMemberToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FissionType)) {
		body["fissionType"] = request.FissionType
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUnionId)) {
		body["memberUnionId"] = request.MemberUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUserId)) {
		body["memberUserId"] = request.MemberUserId
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
	params := &openapi.Params{
		Action:      tea.String("AddContactMemberToGroup"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/contacts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddContactMemberToGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加联系人到群里
//
// @param request - AddContactMemberToGroupRequest
//
// @return AddContactMemberToGroupResponse
func (client *Client) AddContactMemberToGroup(request *AddContactMemberToGroupRequest) (_result *AddContactMemberToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddContactMemberToGroupHeaders{}
	_result = &AddContactMemberToGroupResponse{}
	_body, _err := client.AddContactMemberToGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加知识点
//
// @param request - AddKnowledgeRequest
//
// @param headers - AddKnowledgeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddKnowledgeResponse
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

	if !tea.BoolValue(util.IsUnset(request.LibraryKey)) {
		body["libraryKey"] = request.LibraryKey
	}

	if !tea.BoolValue(util.IsUnset(request.LinkUrl)) {
		body["linkUrl"] = request.LinkUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.QuestionIds)) {
		body["questionIds"] = request.QuestionIds
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
	params := &openapi.Params{
		Action:      tea.String("AddKnowledge"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/knowledges"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddKnowledgeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加知识点
//
// @param request - AddKnowledgeRequest
//
// @return AddKnowledgeResponse
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

// Summary:
//
// 添加服务群知识库
//
// @param request - AddLibraryRequest
//
// @param headers - AddLibraryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLibraryResponse
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
	params := &openapi.Params{
		Action:      tea.String("AddLibrary"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/librarys"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddLibraryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加服务群知识库
//
// @param request - AddLibraryRequest
//
// @return AddLibraryResponse
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

// Summary:
//
// 添加服务群群成员
//
// @param request - AddMemberToServiceGroupRequest
//
// @param headers - AddMemberToServiceGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMemberToServiceGroupResponse
func (client *Client) AddMemberToServiceGroupWithOptions(request *AddMemberToServiceGroupRequest, headers *AddMemberToServiceGroupHeaders, runtime *util.RuntimeOptions) (_result *AddMemberToServiceGroupResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("AddMemberToServiceGroup"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMemberToServiceGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加服务群群成员
//
// @param request - AddMemberToServiceGroupRequest
//
// @return AddMemberToServiceGroupResponse
func (client *Client) AddMemberToServiceGroup(request *AddMemberToServiceGroupRequest) (_result *AddMemberToServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddMemberToServiceGroupHeaders{}
	_result = &AddMemberToServiceGroupResponse{}
	_body, _err := client.AddMemberToServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # OpenApi添加知识点类目
//
// @param request - AddOpenCategoryRequest
//
// @param headers - AddOpenCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOpenCategoryResponse
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
	params := &openapi.Params{
		Action:      tea.String("AddOpenCategory"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/openCategories"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddOpenCategoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OpenApi添加知识点类目
//
// @param request - AddOpenCategoryRequest
//
// @return AddOpenCategoryResponse
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

// Summary:
//
// # OpenApi添加知识入库
//
// @param request - AddOpenKnowledgeRequest
//
// @param headers - AddOpenKnowledgeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOpenKnowledgeResponse
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
	params := &openapi.Params{
		Action:      tea.String("AddOpenKnowledge"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/openKnowledges"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddOpenKnowledgeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OpenApi添加知识入库
//
// @param request - AddOpenKnowledgeRequest
//
// @return AddOpenKnowledgeResponse
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

// Summary:
//
// 智能服务群知识库创建
//
// @param request - AddOpenLibraryRequest
//
// @param headers - AddOpenLibraryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOpenLibraryResponse
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
	params := &openapi.Params{
		Action:      tea.String("AddOpenLibrary"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/openLibraries"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddOpenLibraryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能服务群知识库创建
//
// @param request - AddOpenLibraryRequest
//
// @return AddOpenLibraryResponse
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

// Summary:
//
// 添加工单备注
//
// @param request - AddTicketMemoRequest
//
// @param headers - AddTicketMemoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTicketMemoResponse
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

	if !tea.BoolValue(util.IsUnset(request.TicketMemo)) {
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
	params := &openapi.Params{
		Action:      tea.String("AddTicketMemo"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets/memos"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTicketMemoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加工单备注
//
// @param request - AddTicketMemoRequest
//
// @return AddTicketMemoResponse
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

// Summary:
//
// 工单指派
//
// @param request - AssignTicketRequest
//
// @param headers - AssignTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignTicketResponse
func (client *Client) AssignTicketWithOptions(request *AssignTicketRequest, headers *AssignTicketHeaders, runtime *util.RuntimeOptions) (_result *AssignTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Notify)) {
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

	if !tea.BoolValue(util.IsUnset(request.TicketMemo)) {
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
	params := &openapi.Params{
		Action:      tea.String("AssignTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets/assign"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工单指派
//
// @param request - AssignTicketRequest
//
// @return AssignTicketResponse
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

// Summary:
//
// 批量绑定服务群业务ID
//
// @param request - BatchBindingGroupBizIdsRequest
//
// @param headers - BatchBindingGroupBizIdsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindingGroupBizIdsResponse
func (client *Client) BatchBindingGroupBizIdsWithOptions(request *BatchBindingGroupBizIdsRequest, headers *BatchBindingGroupBizIdsHeaders, runtime *util.RuntimeOptions) (_result *BatchBindingGroupBizIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingGroupBizIds)) {
		body["bindingGroupBizIds"] = request.BindingGroupBizIds
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
	params := &openapi.Params{
		Action:      tea.String("BatchBindingGroupBizIds"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/bind"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchBindingGroupBizIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量绑定服务群业务ID
//
// @param request - BatchBindingGroupBizIdsRequest
//
// @return BatchBindingGroupBizIdsResponse
func (client *Client) BatchBindingGroupBizIds(request *BatchBindingGroupBizIdsRequest) (_result *BatchBindingGroupBizIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchBindingGroupBizIdsHeaders{}
	_result = &BatchBindingGroupBizIdsResponse{}
	_body, _err := client.BatchBindingGroupBizIdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询群组配置
//
// @param request - BatchGetGroupSetConfigRequest
//
// @param headers - BatchGetGroupSetConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetGroupSetConfigResponse
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
	params := &openapi.Params{
		Action:      tea.String("BatchGetGroupSetConfig"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groupSetConfigs/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetGroupSetConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询群组配置
//
// @param request - BatchGetGroupSetConfigRequest
//
// @return BatchGetGroupSetConfigResponse
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

// Summary:
//
// 批量查询客户群发任务
//
// @param request - BatchQueryCustomerSendTaskRequest
//
// @param headers - BatchQueryCustomerSendTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryCustomerSendTaskResponse
func (client *Client) BatchQueryCustomerSendTaskWithOptions(request *BatchQueryCustomerSendTaskRequest, headers *BatchQueryCustomerSendTaskHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryCustomerSendTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GmtCreateEnd)) {
		body["gmtCreateEnd"] = request.GmtCreateEnd
	}

	if !tea.BoolValue(util.IsUnset(request.GmtCreateStart)) {
		body["gmtCreateStart"] = request.GmtCreateStart
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NeedRichStatistics)) {
		body["needRichStatistics"] = request.NeedRichStatistics
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenBatchTaskIds)) {
		body["openBatchTaskIds"] = request.OpenBatchTaskIds
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
	params := &openapi.Params{
		Action:      tea.String("BatchQueryCustomerSendTask"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customers/tasks/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryCustomerSendTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询客户群发任务
//
// @param request - BatchQueryCustomerSendTaskRequest
//
// @return BatchQueryCustomerSendTaskResponse
func (client *Client) BatchQueryCustomerSendTask(request *BatchQueryCustomerSendTaskRequest) (_result *BatchQueryCustomerSendTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryCustomerSendTaskHeaders{}
	_result = &BatchQueryCustomerSendTaskResponse{}
	_body, _err := client.BatchQueryCustomerSendTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询群成员
//
// @param request - BatchQueryGroupMemberRequest
//
// @param headers - BatchQueryGroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryGroupMemberResponse
func (client *Client) BatchQueryGroupMemberWithOptions(request *BatchQueryGroupMemberRequest, headers *BatchQueryGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryGroupMemberResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("BatchQueryGroupMember"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/members/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询群成员
//
// @param request - BatchQueryGroupMemberRequest
//
// @return BatchQueryGroupMemberResponse
func (client *Client) BatchQueryGroupMember(request *BatchQueryGroupMemberRequest) (_result *BatchQueryGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryGroupMemberHeaders{}
	_result = &BatchQueryGroupMemberResponse{}
	_body, _err := client.BatchQueryGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 群发任务批量查询
//
// @param request - BatchQuerySendMessageTaskRequest
//
// @param headers - BatchQuerySendMessageTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQuerySendMessageTaskResponse
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
	params := &openapi.Params{
		Action:      tea.String("BatchQuerySendMessageTask"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tasks/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQuerySendMessageTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群发任务批量查询
//
// @param request - BatchQuerySendMessageTaskRequest
//
// @return BatchQuerySendMessageTaskResponse
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

// Summary:
//
// 绑定服务群模板到团队
//
// @param request - BoundTemplateToTeamRequest
//
// @param headers - BoundTemplateToTeamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BoundTemplateToTeamResponse
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
	params := &openapi.Params{
		Action:      tea.String("BoundTemplateToTeam"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/teams/templates/bound"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BoundTemplateToTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定服务群模板到团队
//
// @param request - BoundTemplateToTeamRequest
//
// @return BoundTemplateToTeamResponse
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

// Summary:
//
// 撤销工单
//
// @param request - CancelTicketRequest
//
// @param headers - CancelTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTicketResponse
func (client *Client) CancelTicketWithOptions(request *CancelTicketRequest, headers *CancelTicketHeaders, runtime *util.RuntimeOptions) (_result *CancelTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Notify)) {
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

	if !tea.BoolValue(util.IsUnset(request.TicketMemo)) {
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
	params := &openapi.Params{
		Action:      tea.String("CancelTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销工单
//
// @param request - CancelTicketRequest
//
// @return CancelTicketResponse
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

// Summary:
//
// 心声总览自定义分类统计
//
// @param request - CategoryStatisticsRequest
//
// @param headers - CategoryStatisticsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CategoryStatisticsResponse
func (client *Client) CategoryStatisticsWithOptions(request *CategoryStatisticsRequest, headers *CategoryStatisticsHeaders, runtime *util.RuntimeOptions) (_result *CategoryStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxDt)) {
		query["maxDt"] = request.MaxDt
	}

	if !tea.BoolValue(util.IsUnset(request.MinDt)) {
		query["minDt"] = request.MinDt
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
	params := &openapi.Params{
		Action:      tea.String("CategoryStatistics"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/voices/dashboards/categories/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CategoryStatisticsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 心声总览自定义分类统计
//
// @param request - CategoryStatisticsRequest
//
// @return CategoryStatisticsResponse
func (client *Client) CategoryStatistics(request *CategoryStatisticsRequest) (_result *CategoryStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CategoryStatisticsHeaders{}
	_result = &CategoryStatisticsResponse{}
	_body, _err := client.CategoryStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭会话回调
//
// @param request - CloseConversationRequest
//
// @param headers - CloseConversationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseConversationResponse
func (client *Client) CloseConversationWithOptions(request *CloseConversationRequest, headers *CloseConversationHeaders, runtime *util.RuntimeOptions) (_result *CloseConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerTips)) {
		body["serverTips"] = request.ServerTips
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceToken)) {
		body["serviceToken"] = request.ServiceToken
	}

	if !tea.BoolValue(util.IsUnset(request.TargetChannel)) {
		body["targetChannel"] = request.TargetChannel
	}

	if !tea.BoolValue(util.IsUnset(request.VisitorToken)) {
		body["visitorToken"] = request.VisitorToken
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
		Action:      tea.String("CloseConversation"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/conversions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseConversationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭会话回调
//
// @param request - CloseConversationRequest
//
// @return CloseConversationResponse
func (client *Client) CloseConversation(request *CloseConversationRequest) (_result *CloseConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CloseConversationHeaders{}
	_result = &CloseConversationResponse{}
	_body, _err := client.CloseConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭人工会话
//
// @param request - CloseHumanSessionRequest
//
// @param headers - CloseHumanSessionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseHumanSessionResponse
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
	params := &openapi.Params{
		Action:      tea.String("CloseHumanSession"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/humanSessions/close"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseHumanSessionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭人工会话
//
// @param request - CloseHumanSessionRequest
//
// @return CloseHumanSessionResponse
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

// Summary:
//
// 客服分配成功通知回调
//
// @param request - ConversationCreatedNotifyRequest
//
// @param headers - ConversationCreatedNotifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConversationCreatedNotifyResponse
func (client *Client) ConversationCreatedNotifyWithOptions(request *ConversationCreatedNotifyRequest, headers *ConversationCreatedNotifyHeaders, runtime *util.RuntimeOptions) (_result *ConversationCreatedNotifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlipayUserId)) {
		body["alipayUserId"] = request.AlipayUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		body["nickName"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerName)) {
		body["serverName"] = request.ServerName
	}

	if !tea.BoolValue(util.IsUnset(request.ServerTips)) {
		body["serverTips"] = request.ServerTips
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceToken)) {
		body["serviceToken"] = request.ServiceToken
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutRemindTips)) {
		body["timeoutRemindTips"] = request.TimeoutRemindTips
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.VisitorToken)) {
		body["visitorToken"] = request.VisitorToken
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
		Action:      tea.String("ConversationCreatedNotify"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConversationCreatedNotifyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 客服分配成功通知回调
//
// @param request - ConversationCreatedNotifyRequest
//
// @return ConversationCreatedNotifyResponse
func (client *Client) ConversationCreatedNotify(request *ConversationCreatedNotifyRequest) (_result *ConversationCreatedNotifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConversationCreatedNotifyHeaders{}
	_result = &ConversationCreatedNotifyResponse{}
	_body, _err := client.ConversationCreatedNotifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会话转接开始通知回调
//
// @param request - ConversationTransferBeginNotifyRequest
//
// @param headers - ConversationTransferBeginNotifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConversationTransferBeginNotifyResponse
func (client *Client) ConversationTransferBeginNotifyWithOptions(request *ConversationTransferBeginNotifyRequest, headers *ConversationTransferBeginNotifyHeaders, runtime *util.RuntimeOptions) (_result *ConversationTransferBeginNotifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Memo)) {
		body["memo"] = request.Memo
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceToken)) {
		body["serviceToken"] = request.ServiceToken
	}

	if !tea.BoolValue(util.IsUnset(request.SourceSkillGroupId)) {
		body["sourceSkillGroupId"] = request.SourceSkillGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSkillGroupId)) {
		body["targetSkillGroupId"] = request.TargetSkillGroupId
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
		Action:      tea.String("ConversationTransferBeginNotify"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/transfers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConversationTransferBeginNotifyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会话转接开始通知回调
//
// @param request - ConversationTransferBeginNotifyRequest
//
// @return ConversationTransferBeginNotifyResponse
func (client *Client) ConversationTransferBeginNotify(request *ConversationTransferBeginNotifyRequest) (_result *ConversationTransferBeginNotifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConversationTransferBeginNotifyHeaders{}
	_result = &ConversationTransferBeginNotifyResponse{}
	_body, _err := client.ConversationTransferBeginNotifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会话转接完成通知回调
//
// @param request - ConversationTransferCompleteNotifyRequest
//
// @param headers - ConversationTransferCompleteNotifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConversationTransferCompleteNotifyResponse
func (client *Client) ConversationTransferCompleteNotifyWithOptions(request *ConversationTransferCompleteNotifyRequest, headers *ConversationTransferCompleteNotifyHeaders, runtime *util.RuntimeOptions) (_result *ConversationTransferCompleteNotifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlipayUserId)) {
		body["alipayUserId"] = request.AlipayUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["conversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		body["nickName"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceToken)) {
		body["serviceToken"] = request.ServiceToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.VisitorToken)) {
		body["visitorToken"] = request.VisitorToken
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
		Action:      tea.String("ConversationTransferCompleteNotify"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/completes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConversationTransferCompleteNotifyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会话转接完成通知回调
//
// @param request - ConversationTransferCompleteNotifyRequest
//
// @return ConversationTransferCompleteNotifyResponse
func (client *Client) ConversationTransferCompleteNotify(request *ConversationTransferCompleteNotifyRequest) (_result *ConversationTransferCompleteNotifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConversationTransferCompleteNotifyHeaders{}
	_result = &ConversationTransferCompleteNotifyResponse{}
	_body, _err := client.ConversationTransferCompleteNotifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建服务群
//
// @param request - CreateGroupRequest
//
// @param headers - CreateGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
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
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务群
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
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

// Summary:
//
// 创建主动会话接口
//
// @param request - CreateGroupConversationRequest
//
// @param headers - CreateGroupConversationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupConversationResponse
func (client *Client) CreateGroupConversationWithOptions(request *CreateGroupConversationRequest, headers *CreateGroupConversationHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingGroupId)) {
		body["dingGroupId"] = request.DingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingUserId)) {
		body["dingUserId"] = request.DingUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DingUserName)) {
		body["dingUserName"] = request.DingUserName
	}

	if !tea.BoolValue(util.IsUnset(request.ExtValues)) {
		body["extValues"] = request.ExtValues
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["serverGroupId"] = request.ServerGroupId
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
		Action:      tea.String("CreateGroupConversation"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/create/conversations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupConversationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建主动会话接口
//
// @param request - CreateGroupConversationRequest
//
// @return CreateGroupConversationResponse
func (client *Client) CreateGroupConversation(request *CreateGroupConversationRequest) (_result *CreateGroupConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupConversationHeaders{}
	_result = &CreateGroupConversationResponse{}
	_body, _err := client.CreateGroupConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建服务群群分组
//
// @param request - CreateGroupSetRequest
//
// @param headers - CreateGroupSetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupSetResponse
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
	params := &openapi.Params{
		Action:      tea.String("CreateGroupSet"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groupSets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务群群分组
//
// @param request - CreateGroupSetRequest
//
// @return CreateGroupSetResponse
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

// Summary:
//
// 服务群新增表单实例
//
// @param request - CreateInstanceRequest
//
// @param headers - CreateInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers *CreateInstanceHeaders, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalBizId)) {
		body["externalBizId"] = request.ExternalBizId
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		body["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataList)) {
		body["formDataList"] = request.FormDataList
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUnionId)) {
		body["ownerUnionId"] = request.OwnerUnionId
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
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customForms/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务群新增表单实例
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateInstanceHeaders{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建服务群团队
//
// @param request - CreateTeamRequest
//
// @param headers - CreateTeamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTeamResponse
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
	params := &openapi.Params{
		Action:      tea.String("CreateTeam"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/teams"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务群团队
//
// @param request - CreateTeamRequest
//
// @return CreateTeamResponse
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

// Summary:
//
// 创建工单
//
// @param request - CreateTicketRequest
//
// @param headers - CreateTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
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

	if !tea.BoolValue(util.IsUnset(request.Notify)) {
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

	if !tea.BoolValue(util.IsUnset(request.SceneContext)) {
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
	params := &openapi.Params{
		Action:      tea.String("CreateTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工单
//
// @param request - CreateTicketRequest
//
// @return CreateTicketResponse
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

// Summary:
//
// 客户群发任务
//
// @param request - CustomerSendMsgTaskRequest
//
// @param headers - CustomerSendMsgTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomerSendMsgTaskResponse
func (client *Client) CustomerSendMsgTaskWithOptions(request *CustomerSendMsgTaskRequest, headers *CustomerSendMsgTaskHeaders, runtime *util.RuntimeOptions) (_result *CustomerSendMsgTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageContent)) {
		body["messageContent"] = request.MessageContent
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCustomer)) {
		body["queryCustomer"] = request.QueryCustomer
	}

	if !tea.BoolValue(util.IsUnset(request.SendConfig)) {
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
	params := &openapi.Params{
		Action:      tea.String("CustomerSendMsgTask"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customers/tasks/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomerSendMsgTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 客户群发任务
//
// @param request - CustomerSendMsgTaskRequest
//
// @return CustomerSendMsgTaskResponse
func (client *Client) CustomerSendMsgTask(request *CustomerSendMsgTaskRequest) (_result *CustomerSendMsgTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomerSendMsgTaskHeaders{}
	_result = &CustomerSendMsgTaskResponse{}
	_body, _err := client.CustomerSendMsgTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从群或者群组剔除成员
//
// @param request - DeleteGroupMembersFromGroupRequest
//
// @param headers - DeleteGroupMembersFromGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupMembersFromGroupResponse
func (client *Client) DeleteGroupMembersFromGroupWithOptions(request *DeleteGroupMembersFromGroupRequest, headers *DeleteGroupMembersFromGroupHeaders, runtime *util.RuntimeOptions) (_result *DeleteGroupMembersFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteGroupType)) {
		body["deleteGroupType"] = request.DeleteGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUnionId)) {
		body["memberUnionId"] = request.MemberUnionId
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
		Action:      tea.String("DeleteGroupMembersFromGroup"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/members/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupMembersFromGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从群或者群组剔除成员
//
// @param request - DeleteGroupMembersFromGroupRequest
//
// @return DeleteGroupMembersFromGroupResponse
func (client *Client) DeleteGroupMembersFromGroup(request *DeleteGroupMembersFromGroupRequest) (_result *DeleteGroupMembersFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteGroupMembersFromGroupHeaders{}
	_result = &DeleteGroupMembersFromGroupResponse{}
	_body, _err := client.DeleteGroupMembersFromGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务群删除表单实例
//
// @param request - DeleteInstanceRequest
//
// @param headers - DeleteInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, headers *DeleteInstanceHeaders, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		body["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenDataInstanceId)) {
		body["openDataInstanceId"] = request.OpenDataInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
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
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customForms/instances/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务群删除表单实例
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteInstanceHeaders{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务群删除知识点
//
// @param request - DeleteKnowledgeRequest
//
// @param headers - DeleteKnowledgeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKnowledgeResponse
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
	params := &openapi.Params{
		Action:      tea.String("DeleteKnowledge"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/knowledges/batchDelete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteKnowledgeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务群删除知识点
//
// @param request - DeleteKnowledgeRequest
//
// @return DeleteKnowledgeResponse
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

// Summary:
//
// 客户心声负面情绪统计
//
// @param request - EmotionStatisticsRequest
//
// @param headers - EmotionStatisticsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EmotionStatisticsResponse
func (client *Client) EmotionStatisticsWithOptions(request *EmotionStatisticsRequest, headers *EmotionStatisticsHeaders, runtime *util.RuntimeOptions) (_result *EmotionStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxDt)) {
		query["maxDt"] = request.MaxDt
	}

	if !tea.BoolValue(util.IsUnset(request.MaxEmotion)) {
		query["maxEmotion"] = request.MaxEmotion
	}

	if !tea.BoolValue(util.IsUnset(request.MinDt)) {
		query["minDt"] = request.MinDt
	}

	if !tea.BoolValue(util.IsUnset(request.MinEmotion)) {
		query["minEmotion"] = request.MinEmotion
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationIds)) {
		query["openConversationIds"] = request.OpenConversationIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		query["openGroupSetId"] = request.OpenGroupSetId
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
	params := &openapi.Params{
		Action:      tea.String("EmotionStatistics"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/voices/emotions/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EmotionStatisticsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 客户心声负面情绪统计
//
// @param request - EmotionStatisticsRequest
//
// @return EmotionStatisticsResponse
func (client *Client) EmotionStatistics(request *EmotionStatisticsRequest) (_result *EmotionStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EmotionStatisticsHeaders{}
	_result = &EmotionStatisticsResponse{}
	_body, _err := client.EmotionStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 结单
//
// @param request - FinishTicketRequest
//
// @param headers - FinishTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishTicketResponse
func (client *Client) FinishTicketWithOptions(request *FinishTicketRequest, headers *FinishTicketHeaders, runtime *util.RuntimeOptions) (_result *FinishTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Notify)) {
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

	if !tea.BoolValue(util.IsUnset(request.TicketMemo)) {
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
	params := &openapi.Params{
		Action:      tea.String("FinishTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets/finish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 结单
//
// @param request - FinishTicketRequest
//
// @return FinishTicketResponse
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

// Summary:
//
// 获取签权Token
//
// @param request - GetAuthTokenRequest
//
// @param headers - GetAuthTokenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthTokenResponse
func (client *Client) GetAuthTokenWithOptions(request *GetAuthTokenRequest, headers *GetAuthTokenHeaders, runtime *util.RuntimeOptions) (_result *GetAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTime)) {
		body["effectiveTime"] = request.EffectiveTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerId)) {
		body["serverId"] = request.ServerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerName)) {
		body["serverName"] = request.ServerName
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
		Action:      tea.String("GetAuthToken"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/get/tokens"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取签权Token
//
// @param request - GetAuthTokenRequest
//
// @return GetAuthTokenResponse
func (client *Client) GetAuthToken(request *GetAuthTokenRequest) (_result *GetAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAuthTokenHeaders{}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.GetAuthTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务群通过实例ID集合批量查询表单实例数据
//
// @param request - GetInstancesByIdsRequest
//
// @param headers - GetInstancesByIdsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstancesByIdsResponse
func (client *Client) GetInstancesByIdsWithOptions(request *GetInstancesByIdsRequest, headers *GetInstancesByIdsHeaders, runtime *util.RuntimeOptions) (_result *GetInstancesByIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		body["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenDataInstanceIdList)) {
		body["openDataInstanceIdList"] = request.OpenDataInstanceIdList
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
	params := &openapi.Params{
		Action:      tea.String("GetInstancesByIds"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customForms/instances/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstancesByIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务群通过实例ID集合批量查询表单实例数据
//
// @param request - GetInstancesByIdsRequest
//
// @return GetInstancesByIdsResponse
func (client *Client) GetInstancesByIds(request *GetInstancesByIdsRequest) (_result *GetInstancesByIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstancesByIdsHeaders{}
	_result = &GetInstancesByIdsResponse{}
	_body, _err := client.GetInstancesByIdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取负面心声词云
//
// @param request - GetNegativeWordCloudRequest
//
// @param headers - GetNegativeWordCloudHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNegativeWordCloudResponse
func (client *Client) GetNegativeWordCloudWithOptions(request *GetNegativeWordCloudRequest, headers *GetNegativeWordCloudHeaders, runtime *util.RuntimeOptions) (_result *GetNegativeWordCloudResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetNegativeWordCloud"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/voices/negatives/wordClouds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNegativeWordCloudResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取负面心声词云
//
// @param request - GetNegativeWordCloudRequest
//
// @return GetNegativeWordCloudResponse
func (client *Client) GetNegativeWordCloud(request *GetNegativeWordCloudRequest) (_result *GetNegativeWordCloudResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetNegativeWordCloudHeaders{}
	_result = &GetNegativeWordCloudResponse{}
	_body, _err := client.GetNegativeWordCloudWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取临时访问链接
//
// @param request - GetOssTempUrlRequest
//
// @param headers - GetOssTempUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssTempUrlResponse
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
	params := &openapi.Params{
		Action:      tea.String("GetOssTempUrl"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/ossServices/tempUrls"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssTempUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取临时访问链接
//
// @param request - GetOssTempUrlRequest
//
// @return GetOssTempUrlResponse
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

// Summary:
//
// 获取表单上传凭证
//
// @param request - GetStoragePolicyRequest
//
// @param headers - GetStoragePolicyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStoragePolicyResponse
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
	params := &openapi.Params{
		Action:      tea.String("GetStoragePolicy"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/ossServices/policies"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStoragePolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表单上传凭证
//
// @param request - GetStoragePolicyRequest
//
// @return GetStoragePolicyResponse
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

// Summary:
//
// 查询工单详情
//
// @param request - GetTicketRequest
//
// @param headers - GetTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTicketResponse
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
	params := &openapi.Params{
		Action:      tea.String("GetTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工单详情
//
// @param request - GetTicketRequest
//
// @return GetTicketResponse
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

// Summary:
//
// 获取心声词云
//
// @param request - GetWordCloudRequest
//
// @param headers - GetWordCloudHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWordCloudResponse
func (client *Client) GetWordCloudWithOptions(request *GetWordCloudRequest, headers *GetWordCloudHeaders, runtime *util.RuntimeOptions) (_result *GetWordCloudResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetWordCloud"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/voices/wordClouds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWordCloudResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取心声词云
//
// @param request - GetWordCloudRequest
//
// @return GetWordCloudResponse
func (client *Client) GetWordCloud(request *GetWordCloudRequest) (_result *GetWordCloudResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWordCloudHeaders{}
	_result = &GetWordCloudResponse{}
	_body, _err := client.GetWordCloudWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 心声总览群统计
//
// @param request - GroupStatisticsRequest
//
// @param headers - GroupStatisticsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupStatisticsResponse
func (client *Client) GroupStatisticsWithOptions(request *GroupStatisticsRequest, headers *GroupStatisticsHeaders, runtime *util.RuntimeOptions) (_result *GroupStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxDt)) {
		query["maxDt"] = request.MaxDt
	}

	if !tea.BoolValue(util.IsUnset(request.MinDt)) {
		query["minDt"] = request.MinDt
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
	params := &openapi.Params{
		Action:      tea.String("GroupStatistics"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/voices/dashboards/groups/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupStatisticsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 心声总览群统计
//
// @param request - GroupStatisticsRequest
//
// @return GroupStatisticsResponse
func (client *Client) GroupStatistics(request *GroupStatisticsRequest) (_result *GroupStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupStatisticsHeaders{}
	_result = &GroupStatisticsResponse{}
	_body, _err := client.GroupStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 心声总览意图&自定义分类统计
//
// @param request - IntentionCategoryStatisticsRequest
//
// @param headers - IntentionCategoryStatisticsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntentionCategoryStatisticsResponse
func (client *Client) IntentionCategoryStatisticsWithOptions(request *IntentionCategoryStatisticsRequest, headers *IntentionCategoryStatisticsHeaders, runtime *util.RuntimeOptions) (_result *IntentionCategoryStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxDt)) {
		query["maxDt"] = request.MaxDt
	}

	if !tea.BoolValue(util.IsUnset(request.MinDt)) {
		query["minDt"] = request.MinDt
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
	params := &openapi.Params{
		Action:      tea.String("IntentionCategoryStatistics"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/voices/dashboards/intentionCategories/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IntentionCategoryStatisticsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 心声总览意图&自定义分类统计
//
// @param request - IntentionCategoryStatisticsRequest
//
// @return IntentionCategoryStatisticsResponse
func (client *Client) IntentionCategoryStatistics(request *IntentionCategoryStatisticsRequest) (_result *IntentionCategoryStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IntentionCategoryStatisticsHeaders{}
	_result = &IntentionCategoryStatisticsResponse{}
	_body, _err := client.IntentionCategoryStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 心声总览意图统计
//
// @param request - IntentionStatisticsRequest
//
// @param headers - IntentionStatisticsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IntentionStatisticsResponse
func (client *Client) IntentionStatisticsWithOptions(request *IntentionStatisticsRequest, headers *IntentionStatisticsHeaders, runtime *util.RuntimeOptions) (_result *IntentionStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxDt)) {
		query["maxDt"] = request.MaxDt
	}

	if !tea.BoolValue(util.IsUnset(request.MinDt)) {
		query["minDt"] = request.MinDt
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
	params := &openapi.Params{
		Action:      tea.String("IntentionStatistics"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/voices/dashboards/intentions/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IntentionStatisticsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 心声总览意图统计
//
// @param request - IntentionStatisticsRequest
//
// @return IntentionStatisticsResponse
func (client *Client) IntentionStatistics(request *IntentionStatisticsRequest) (_result *IntentionStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IntentionStatisticsHeaders{}
	_result = &IntentionStatisticsResponse{}
	_body, _err := client.IntentionStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工单操作记录
//
// @param request - ListTicketOperateRecordRequest
//
// @param headers - ListTicketOperateRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketOperateRecordResponse
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
	params := &openapi.Params{
		Action:      tea.String("ListTicketOperateRecord"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets/operateRecords"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTicketOperateRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工单操作记录
//
// @param request - ListTicketOperateRecordRequest
//
// @return ListTicketOperateRecordResponse
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

// Summary:
//
// 查询用户所在团队
//
// @param headers - ListUserTeamsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserTeamsResponse
func (client *Client) ListUserTeamsWithOptions(userId *string, headers *ListUserTeamsHeaders, runtime *util.RuntimeOptions) (_result *ListUserTeamsResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("ListUserTeams"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/users/" + tea.StringValue(userId) + "/teams"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserTeamsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户所在团队
//
// @return ListUserTeamsResponse
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

// Summary:
//
// 查询服务群活跃成员
//
// @param request - QueryActiveUsersRequest
//
// @param headers - QueryActiveUsersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryActiveUsersResponse
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
	params := &openapi.Params{
		Action:      tea.String("QueryActiveUsers"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/queryActiveUsers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryActiveUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务群活跃成员
//
// @param request - QueryActiveUsersRequest
//
// @return QueryActiveUsersResponse
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

// Summary:
//
// 群联系人画像检索
//
// @param request - QueryCrmGroupContactRequest
//
// @param headers - QueryCrmGroupContactHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCrmGroupContactResponse
func (client *Client) QueryCrmGroupContactWithOptions(request *QueryCrmGroupContactRequest, headers *QueryCrmGroupContactHeaders, runtime *util.RuntimeOptions) (_result *QueryCrmGroupContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MinResult)) {
		body["minResult"] = request.MinResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFields)) {
		body["searchFields"] = request.SearchFields
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
		Action:      tea.String("QueryCrmGroupContact"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/contacts/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCrmGroupContactResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群联系人画像检索
//
// @param request - QueryCrmGroupContactRequest
//
// @return QueryCrmGroupContactResponse
func (client *Client) QueryCrmGroupContact(request *QueryCrmGroupContactRequest) (_result *QueryCrmGroupContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCrmGroupContactHeaders{}
	_result = &QueryCrmGroupContactResponse{}
	_body, _err := client.QueryCrmGroupContactWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询客户信息
//
// @param request - QueryCustomerCardRequest
//
// @param headers - QueryCustomerCardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerCardResponse
func (client *Client) QueryCustomerCardWithOptions(request *QueryCustomerCardRequest, headers *QueryCustomerCardHeaders, runtime *util.RuntimeOptions) (_result *QueryCustomerCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonParams)) {
		body["jsonParams"] = request.JsonParams
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
	params := &openapi.Params{
		Action:      tea.String("QueryCustomerCard"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/userDetials"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCustomerCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户信息
//
// @param request - QueryCustomerCardRequest
//
// @return QueryCustomerCardResponse
func (client *Client) QueryCustomerCard(request *QueryCustomerCardRequest) (_result *QueryCustomerCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCustomerCardHeaders{}
	_result = &QueryCustomerCardResponse{}
	_body, _err := client.QueryCustomerCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询客户群发任务客户触达详情
//
// @param request - QueryCustomerTaskUserDetailRequest
//
// @param headers - QueryCustomerTaskUserDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerTaskUserDetailResponse
func (client *Client) QueryCustomerTaskUserDetailWithOptions(request *QueryCustomerTaskUserDetailRequest, headers *QueryCustomerTaskUserDetailHeaders, runtime *util.RuntimeOptions) (_result *QueryCustomerTaskUserDetailResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ReceiverUnionIds)) {
		body["receiverUnionIds"] = request.ReceiverUnionIds
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
		Action:      tea.String("QueryCustomerTaskUserDetail"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customers/tasks/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCustomerTaskUserDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户群发任务客户触达详情
//
// @param request - QueryCustomerTaskUserDetailRequest
//
// @return QueryCustomerTaskUserDetailResponse
func (client *Client) QueryCustomerTaskUserDetail(request *QueryCustomerTaskUserDetailRequest) (_result *QueryCustomerTaskUserDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCustomerTaskUserDetailHeaders{}
	_result = &QueryCustomerTaskUserDetailResponse{}
	_body, _err := client.QueryCustomerTaskUserDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个服务群信息
//
// @param request - QueryGroupRequest
//
// @param headers - QueryGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupResponse
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
	params := &openapi.Params{
		Action:      tea.String("QueryGroup"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个服务群信息
//
// @param request - QueryGroupRequest
//
// @return QueryGroupResponse
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

// Summary:
//
// 查询指定群成员
//
// @param request - QueryGroupMemberRequest
//
// @param headers - QueryGroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupMemberResponse
func (client *Client) QueryGroupMemberWithOptions(request *QueryGroupMemberRequest, headers *QueryGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupMemberResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("QueryGroupMember"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/members/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定群成员
//
// @param request - QueryGroupMemberRequest
//
// @return QueryGroupMemberResponse
func (client *Client) QueryGroupMember(request *QueryGroupMemberRequest) (_result *QueryGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupMemberHeaders{}
	_result = &QueryGroupMemberResponse{}
	_body, _err := client.QueryGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务群群组信息
//
// @param request - QueryGroupSetRequest
//
// @param headers - QueryGroupSetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupSetResponse
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
	params := &openapi.Params{
		Action:      tea.String("QueryGroupSet"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groupSets"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务群群组信息
//
// @param request - QueryGroupSetRequest
//
// @return QueryGroupSetResponse
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

// Summary:
//
// 服务群通过多添件进行组合检索表单数据实例集合
//
// @param request - QueryInstancesByMultiConditionsRequest
//
// @param headers - QueryInstancesByMultiConditionsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstancesByMultiConditionsResponse
func (client *Client) QueryInstancesByMultiConditionsWithOptions(request *QueryInstancesByMultiConditionsRequest, headers *QueryInstancesByMultiConditionsHeaders, runtime *util.RuntimeOptions) (_result *QueryInstancesByMultiConditionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		body["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFields)) {
		body["searchFields"] = request.SearchFields
	}

	if !tea.BoolValue(util.IsUnset(request.SortFields)) {
		body["sortFields"] = request.SortFields
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
		Action:      tea.String("QueryInstancesByMultiConditions"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customForms/instances/multiConditions/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInstancesByMultiConditionsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务群通过多添件进行组合检索表单数据实例集合
//
// @param request - QueryInstancesByMultiConditionsRequest
//
// @return QueryInstancesByMultiConditionsResponse
func (client *Client) QueryInstancesByMultiConditions(request *QueryInstancesByMultiConditionsRequest) (_result *QueryInstancesByMultiConditionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryInstancesByMultiConditionsHeaders{}
	_result = &QueryInstancesByMultiConditionsResponse{}
	_body, _err := client.QueryInstancesByMultiConditionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 群发任务统计查询
//
// @param request - QuerySendMsgTaskStatisticsRequest
//
// @param headers - QuerySendMsgTaskStatisticsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySendMsgTaskStatisticsResponse
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
	params := &openapi.Params{
		Action:      tea.String("QuerySendMsgTaskStatistics"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tasks/statistics/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySendMsgTaskStatisticsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群发任务统计查询
//
// @param request - QuerySendMsgTaskStatisticsRequest
//
// @return QuerySendMsgTaskStatisticsResponse
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

// Summary:
//
// 群发任务群维度的统计数据
//
// @param request - QuerySendMsgTaskStatisticsDetailRequest
//
// @param headers - QuerySendMsgTaskStatisticsDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySendMsgTaskStatisticsDetailResponse
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
	params := &openapi.Params{
		Action:      tea.String("QuerySendMsgTaskStatisticsDetail"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tasks/statistics/details/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySendMsgTaskStatisticsDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群发任务群维度的统计数据
//
// @param request - QuerySendMsgTaskStatisticsDetailRequest
//
// @return QuerySendMsgTaskStatisticsDetailResponse
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

// Summary:
//
// 查消息的已读/未读列表
//
// @param request - QueryServiceGroupMessageReadStatusRequest
//
// @param headers - QueryServiceGroupMessageReadStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryServiceGroupMessageReadStatusResponse
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
	params := &openapi.Params{
		Action:      tea.String("QueryServiceGroupMessageReadStatus"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/messages/readStatus/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryServiceGroupMessageReadStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查消息的已读/未读列表
//
// @param request - QueryServiceGroupMessageReadStatusRequest
//
// @return QueryServiceGroupMessageReadStatusResponse
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

// Summary:
//
// 外部DT工作台排队通知回调
//
// @param request - QueueNotifyRequest
//
// @param headers - QueueNotifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueueNotifyResponse
func (client *Client) QueueNotifyWithOptions(request *QueueNotifyRequest, headers *QueueNotifyHeaders, runtime *util.RuntimeOptions) (_result *QueueNotifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EstimateWaitMin)) {
		body["estimateWaitMin"] = request.EstimateWaitMin
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.QueuePlace)) {
		body["queuePlace"] = request.QueuePlace
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceToken)) {
		body["serviceToken"] = request.ServiceToken
	}

	if !tea.BoolValue(util.IsUnset(request.TargetChannel)) {
		body["targetChannel"] = request.TargetChannel
	}

	if !tea.BoolValue(util.IsUnset(request.Tips)) {
		body["tips"] = request.Tips
	}

	if !tea.BoolValue(util.IsUnset(request.VisitorToken)) {
		body["visitorToken"] = request.VisitorToken
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
		Action:      tea.String("QueueNotify"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/dts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueueNotifyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 外部DT工作台排队通知回调
//
// @param request - QueueNotifyRequest
//
// @return QueueNotifyResponse
func (client *Client) QueueNotify(request *QueueNotifyRequest) (_result *QueueNotifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueueNotifyHeaders{}
	_result = &QueueNotifyResponse{}
	_body, _err := client.QueueNotifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 组织剔除联系人
//
// @param request - RemoveContactFromOrgRequest
//
// @param headers - RemoveContactFromOrgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveContactFromOrgResponse
func (client *Client) RemoveContactFromOrgWithOptions(request *RemoveContactFromOrgRequest, headers *RemoveContactFromOrgHeaders, runtime *util.RuntimeOptions) (_result *RemoveContactFromOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactUnionId)) {
		body["contactUnionId"] = request.ContactUnionId
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
	params := &openapi.Params{
		Action:      tea.String("RemoveContactFromOrg"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/organizations/contacts/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveContactFromOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 组织剔除联系人
//
// @param request - RemoveContactFromOrgRequest
//
// @return RemoveContactFromOrgResponse
func (client *Client) RemoveContactFromOrg(request *RemoveContactFromOrgRequest) (_result *RemoveContactFromOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveContactFromOrgHeaders{}
	_result = &RemoveContactFromOrgResponse{}
	_body, _err := client.RemoveContactFromOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定群的客户活跃明细查询
//
// @param request - ReportCustomerDetailRequest
//
// @param headers - ReportCustomerDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportCustomerDetailResponse
func (client *Client) ReportCustomerDetailWithOptions(request *ReportCustomerDetailRequest, headers *ReportCustomerDetailHeaders, runtime *util.RuntimeOptions) (_result *ReportCustomerDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HasLogin)) {
		body["hasLogin"] = request.HasLogin
	}

	if !tea.BoolValue(util.IsUnset(request.HasOpenConv)) {
		body["hasOpenConv"] = request.HasOpenConv
	}

	if !tea.BoolValue(util.IsUnset(request.MaxDt)) {
		body["maxDt"] = request.MaxDt
	}

	if !tea.BoolValue(util.IsUnset(request.MinDt)) {
		body["minDt"] = request.MinDt
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
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
		Action:      tea.String("ReportCustomerDetail"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customers/activities/details/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportCustomerDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定群的客户活跃明细查询
//
// @param request - ReportCustomerDetailRequest
//
// @return ReportCustomerDetailResponse
func (client *Client) ReportCustomerDetail(request *ReportCustomerDetailRequest) (_result *ReportCustomerDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReportCustomerDetailHeaders{}
	_result = &ReportCustomerDetailResponse{}
	_body, _err := client.ReportCustomerDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 客户活跃明细指标查询
//
// @param request - ReportCustomerStatisticsRequest
//
// @param headers - ReportCustomerStatisticsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportCustomerStatisticsResponse
func (client *Client) ReportCustomerStatisticsWithOptions(request *ReportCustomerStatisticsRequest, headers *ReportCustomerStatisticsHeaders, runtime *util.RuntimeOptions) (_result *ReportCustomerStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupOwnerUserIds)) {
		body["groupOwnerUserIds"] = request.GroupOwnerUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTags)) {
		body["groupTags"] = request.GroupTags
	}

	if !tea.BoolValue(util.IsUnset(request.MaxDt)) {
		body["maxDt"] = request.MaxDt
	}

	if !tea.BoolValue(util.IsUnset(request.MinDt)) {
		body["minDt"] = request.MinDt
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationIds)) {
		body["openConversationIds"] = request.OpenConversationIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpenGroupSetId)) {
		body["openGroupSetId"] = request.OpenGroupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
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
		Action:      tea.String("ReportCustomerStatistics"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customers/activities/statistics/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportCustomerStatisticsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 客户活跃明细指标查询
//
// @param request - ReportCustomerStatisticsRequest
//
// @return ReportCustomerStatisticsResponse
func (client *Client) ReportCustomerStatistics(request *ReportCustomerStatisticsRequest) (_result *ReportCustomerStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReportCustomerStatisticsHeaders{}
	_result = &ReportCustomerStatisticsResponse{}
	_body, _err := client.ReportCustomerStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新提交工单
//
// @param request - ResubmitTicketRequest
//
// @param headers - ResubmitTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResubmitTicketResponse
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

	if !tea.BoolValue(util.IsUnset(request.Notify)) {
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

	if !tea.BoolValue(util.IsUnset(request.SceneContext)) {
		body["sceneContext"] = request.SceneContext
	}

	if !tea.BoolValue(util.IsUnset(request.TicketMemo)) {
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
	params := &openapi.Params{
		Action:      tea.String("ResubmitTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets/resubmit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResubmitTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新提交工单
//
// @param request - ResubmitTicketRequest
//
// @return ResubmitTicketResponse
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

// Summary:
//
// 撤回工单
//
// @param request - RetractTicketRequest
//
// @param headers - RetractTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetractTicketResponse
func (client *Client) RetractTicketWithOptions(request *RetractTicketRequest, headers *RetractTicketHeaders, runtime *util.RuntimeOptions) (_result *RetractTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Notify)) {
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

	if !tea.BoolValue(util.IsUnset(request.TicketMemo)) {
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
	params := &openapi.Params{
		Action:      tea.String("RetractTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets/retract"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RetractTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤回工单
//
// @param request - RetractTicketRequest
//
// @return RetractTicketResponse
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

// Summary:
//
// 指定群的机器人消息撤回
//
// @param request - RobotMessageRecallRequest
//
// @param headers - RobotMessageRecallHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RobotMessageRecallResponse
func (client *Client) RobotMessageRecallWithOptions(request *RobotMessageRecallRequest, headers *RobotMessageRecallHeaders, runtime *util.RuntimeOptions) (_result *RobotMessageRecallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenMsgId)) {
		body["openMsgId"] = request.OpenMsgId
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
	params := &openapi.Params{
		Action:      tea.String("RobotMessageRecall"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/robots/messages/recall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RobotMessageRecallResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定群的机器人消息撤回
//
// @param request - RobotMessageRecallRequest
//
// @return RobotMessageRecallResponse
func (client *Client) RobotMessageRecall(request *RobotMessageRecallRequest) (_result *RobotMessageRecallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RobotMessageRecallHeaders{}
	_result = &RobotMessageRecallResponse{}
	_body, _err := client.RobotMessageRecallWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务群新增表单实例
//
// @param request - SaveFormInstanceRequest
//
// @param headers - SaveFormInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveFormInstanceResponse
func (client *Client) SaveFormInstanceWithOptions(request *SaveFormInstanceRequest, headers *SaveFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *SaveFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormDataList)) {
		body["formDataList"] = request.FormDataList
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUnionId)) {
		body["ownerUnionId"] = request.OwnerUnionId
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
		Action:      tea.String("SaveFormInstance"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/forms/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveFormInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务群新增表单实例
//
// @param request - SaveFormInstanceRequest
//
// @return SaveFormInstanceResponse
func (client *Client) SaveFormInstance(request *SaveFormInstanceRequest) (_result *SaveFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveFormInstanceHeaders{}
	_result = &SaveFormInstanceResponse{}
	_body, _err := client.SaveFormInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 搜索服务群
//
// @param request - SearchGroupRequest
//
// @param headers - SearchGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchGroupResponse
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
	params := &openapi.Params{
		Action:      tea.String("SearchGroup"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索服务群
//
// @param request - SearchGroupRequest
//
// @return SearchGroupResponse
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

// Summary:
//
// 服务群发任务
//
// @param request - SendMsgByTaskRequest
//
// @param headers - SendMsgByTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMsgByTaskResponse
func (client *Client) SendMsgByTaskWithOptions(request *SendMsgByTaskRequest, headers *SendMsgByTaskHeaders, runtime *util.RuntimeOptions) (_result *SendMsgByTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageContent)) {
		body["messageContent"] = request.MessageContent
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryGroup)) {
		body["queryGroup"] = request.QueryGroup
	}

	if !tea.BoolValue(util.IsUnset(request.SendConfig)) {
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
	params := &openapi.Params{
		Action:      tea.String("SendMsgByTask"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/messages/tasks/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMsgByTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务群发任务
//
// @param request - SendMsgByTaskRequest
//
// @return SendMsgByTaskResponse
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

// Summary:
//
// 增强版客户群发
//
// @param request - SendMsgByTaskSupportInviteJoinOrgRequest
//
// @param headers - SendMsgByTaskSupportInviteJoinOrgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMsgByTaskSupportInviteJoinOrgResponse
func (client *Client) SendMsgByTaskSupportInviteJoinOrgWithOptions(request *SendMsgByTaskSupportInviteJoinOrgRequest, headers *SendMsgByTaskSupportInviteJoinOrgHeaders, runtime *util.RuntimeOptions) (_result *SendMsgByTaskSupportInviteJoinOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageContent)) {
		body["messageContent"] = request.MessageContent
	}

	if !tea.BoolValue(util.IsUnset(request.MobilePhones)) {
		body["mobilePhones"] = request.MobilePhones
	}

	if !tea.BoolValue(util.IsUnset(request.NeedUrlTrack)) {
		body["needUrlTrack"] = request.NeedUrlTrack
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.SendChannel)) {
		body["sendChannel"] = request.SendChannel
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
	params := &openapi.Params{
		Action:      tea.String("SendMsgByTaskSupportInviteJoinOrg"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customers/tasks/groupSend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMsgByTaskSupportInviteJoinOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增强版客户群发
//
// @param request - SendMsgByTaskSupportInviteJoinOrgRequest
//
// @return SendMsgByTaskSupportInviteJoinOrgResponse
func (client *Client) SendMsgByTaskSupportInviteJoinOrg(request *SendMsgByTaskSupportInviteJoinOrgRequest) (_result *SendMsgByTaskSupportInviteJoinOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendMsgByTaskSupportInviteJoinOrgHeaders{}
	_result = &SendMsgByTaskSupportInviteJoinOrgResponse{}
	_body, _err := client.SendMsgByTaskSupportInviteJoinOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务群发消息
//
// @param request - SendServiceGroupMessageRequest
//
// @param headers - SendServiceGroupMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendServiceGroupMessageResponse
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
	params := &openapi.Params{
		Action:      tea.String("SendServiceGroupMessage"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/messages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendServiceGroupMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务群发消息
//
// @param request - SendServiceGroupMessageRequest
//
// @return SendServiceGroupMessageResponse
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

// Summary:
//
// 执行阿里内部用户群组机器人服务开关
//
// @param request - SetRobotConfigRequest
//
// @param headers - SetRobotConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRobotConfigResponse
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
	params := &openapi.Params{
		Action:      tea.String("SetRobotConfig"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/configs/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetRobotConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行阿里内部用户群组机器人服务开关
//
// @param request - SetRobotConfigRequest
//
// @return SetRobotConfigResponse
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

// Summary:
//
// 申领工单
//
// @param request - TakeTicketRequest
//
// @param headers - TakeTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TakeTicketResponse
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
	params := &openapi.Params{
		Action:      tea.String("TakeTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets/apply"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TakeTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申领工单
//
// @param request - TakeTicketRequest
//
// @return TakeTicketResponse
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

// Summary:
//
// 客户心声热门话题统计
//
// @param request - TopicStatisticsRequest
//
// @param headers - TopicStatisticsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TopicStatisticsResponse
func (client *Client) TopicStatisticsWithOptions(request *TopicStatisticsRequest, headers *TopicStatisticsHeaders, runtime *util.RuntimeOptions) (_result *TopicStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxDt)) {
		query["maxDt"] = request.MaxDt
	}

	if !tea.BoolValue(util.IsUnset(request.MinDt)) {
		query["minDt"] = request.MinDt
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationIds)) {
		query["openConversationIds"] = request.OpenConversationIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		query["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchContent)) {
		query["searchContent"] = request.SearchContent
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
		Action:      tea.String("TopicStatistics"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/voices/topics/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TopicStatisticsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 客户心声热门话题统计
//
// @param request - TopicStatisticsRequest
//
// @return TopicStatisticsResponse
func (client *Client) TopicStatistics(request *TopicStatisticsRequest) (_result *TopicStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TopicStatisticsHeaders{}
	_result = &TopicStatisticsResponse{}
	_body, _err := client.TopicStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 转交工单
//
// @param request - TransferTicketRequest
//
// @param headers - TransferTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferTicketResponse
func (client *Client) TransferTicketWithOptions(request *TransferTicketRequest, headers *TransferTicketHeaders, runtime *util.RuntimeOptions) (_result *TransferTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Notify)) {
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

	if !tea.BoolValue(util.IsUnset(request.TicketMemo)) {
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
	params := &openapi.Params{
		Action:      tea.String("TransferTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets/transfer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 转交工单
//
// @param request - TransferTicketRequest
//
// @return TransferTicketResponse
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

// Summary:
//
// 变更群的群组配置信息
//
// @param request - UpdateGroupSetRequest
//
// @param headers - UpdateGroupSetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupSetResponse
func (client *Client) UpdateGroupSetWithOptions(request *UpdateGroupSetRequest, headers *UpdateGroupSetHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupSetResponse, _err error) {
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
		Action:      tea.String("UpdateGroupSet"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/groups/configurations"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更群的群组配置信息
//
// @param request - UpdateGroupSetRequest
//
// @return UpdateGroupSetResponse
func (client *Client) UpdateGroupSet(request *UpdateGroupSetRequest) (_result *UpdateGroupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupSetHeaders{}
	_result = &UpdateGroupSetResponse{}
	_body, _err := client.UpdateGroupSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新服务群标签
//
// @param request - UpdateGroupTagRequest
//
// @param headers - UpdateGroupTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupTagResponse
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
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupTag"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tags"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务群标签
//
// @param request - UpdateGroupTagRequest
//
// @return UpdateGroupTagResponse
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

// Summary:
//
// 服务群更新表单实例
//
// @param request - UpdateInstanceRequest
//
// @param headers - UpdateInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, headers *UpdateInstanceHeaders, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalBizId)) {
		body["externalBizId"] = request.ExternalBizId
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		body["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataList)) {
		body["formDataList"] = request.FormDataList
	}

	if !tea.BoolValue(util.IsUnset(request.OpenDataInstanceId)) {
		body["openDataInstanceId"] = request.OpenDataInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenTeamId)) {
		body["openTeamId"] = request.OpenTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUnionId)) {
		body["ownerUnionId"] = request.OwnerUnionId
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
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/customForms/instances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务群更新表单实例
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInstanceHeaders{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新工单自定义字段
//
// @param request - UpdateTicketRequest
//
// @param headers - UpdateTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTicketResponse
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

	if !tea.BoolValue(util.IsUnset(request.TicketMemo)) {
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
	params := &openapi.Params{
		Action:      tea.String("UpdateTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新工单自定义字段
//
// @param request - UpdateTicketRequest
//
// @return UpdateTicketResponse
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

// Summary:
//
// 将智能云客服下的旧版服务群升级为钉钉智能服务群
//
// @param request - UpgradeCloudGroupRequest
//
// @param headers - UpgradeCloudGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeCloudGroupResponse
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
	params := &openapi.Params{
		Action:      tea.String("UpgradeCloudGroup"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/cloudGroups/upgrade"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeCloudGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将智能云客服下的旧版服务群升级为钉钉智能服务群
//
// @param request - UpgradeCloudGroupRequest
//
// @return UpgradeCloudGroupResponse
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

// Summary:
//
// 将常规钉群升级为钉钉智能服务群
//
// @param request - UpgradeNormalGroupRequest
//
// @param headers - UpgradeNormalGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeNormalGroupResponse
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
	params := &openapi.Params{
		Action:      tea.String("UpgradeNormalGroup"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/normalGroups/upgrade"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeNormalGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将常规钉群升级为钉钉智能服务群
//
// @param request - UpgradeNormalGroupRequest
//
// @return UpgradeNormalGroupResponse
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

// Summary:
//
// 工单催办
//
// @param request - UrgeTicketRequest
//
// @param headers - UrgeTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UrgeTicketResponse
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

	if !tea.BoolValue(util.IsUnset(request.TicketMemo)) {
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
	params := &openapi.Params{
		Action:      tea.String("UrgeTicket"),
		Version:     tea.String("serviceGroup_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/serviceGroup/tickets/urge"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UrgeTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工单催办
//
// @param request - UrgeTicketRequest
//
// @return UrgeTicketResponse
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
